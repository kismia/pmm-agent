// pmm-agent
// Copyright 2019 Percona LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package perfschema runs built-in QAN Agent for MySQL performance schema.
package perfschema

import (
	"context"
	"database/sql"
	"io"
	"math"
	"time"

	"github.com/AlekSi/pointer"
	_ "github.com/go-sql-driver/mysql" // register SQL driver
	"github.com/percona/pmm/api/agentpb"
	"github.com/percona/pmm/api/inventorypb"
	"github.com/sirupsen/logrus"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/mysql"

	"github.com/percona/pmm-agent/agents"
)

const (
	retainHistory  = 5 * time.Minute
	refreshHistory = 5 * time.Second

	retainSummaries = 25 * time.Hour // make it work for daily queries
	querySummaries  = time.Minute
)

// PerfSchema QAN services connects to MySQL and extracts performance data.
type PerfSchema struct {
	q                    *reform.Querier
	dbCloser             io.Closer
	agentID              string
	disableQueryExamples bool
	l                    *logrus.Entry
	changes              chan agents.Change
	historyCache         *historyCache
	summaryCache         *summaryCache
}

// Params represent Agent parameters.
type Params struct {
	DSN                  string
	AgentID              string
	DisableQueryExamples bool
}

const queryTag = "pmm-agent:perfschema"

// New creates new PerfSchema QAN service.
func New(params *Params, l *logrus.Entry) (*PerfSchema, error) {
	sqlDB, err := sql.Open("mysql", params.DSN)
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetMaxOpenConns(1)
	sqlDB.SetConnMaxLifetime(0)
	q := reform.NewDB(sqlDB, mysql.Dialect, reform.NewPrintfLogger(l.Tracef)).WithTag(queryTag)
	return newPerfSchema(q, sqlDB, params.AgentID, l), nil
}

func newPerfSchema(q *reform.Querier, dbCloser io.Closer, agentID string, l *logrus.Entry) *PerfSchema {
	return &PerfSchema{
		q:            q,
		dbCloser:     dbCloser,
		agentID:      agentID,
		l:            l,
		changes:      make(chan agents.Change, 10),
		historyCache: newHistoryCache(retainHistory),
		summaryCache: newSummaryCache(retainSummaries),
	}
}

// Run extracts performance data and sends it to the channel until ctx is canceled.
func (m *PerfSchema) Run(ctx context.Context) {
	defer func() {
		m.dbCloser.Close() //nolint:errcheck
		m.changes <- agents.Change{Status: inventorypb.AgentStatus_DONE}
		close(m.changes)
	}()

	// add current summaries to cache so they are not send as new on first iteration with incorrect timestamps
	var running bool
	m.changes <- agents.Change{Status: inventorypb.AgentStatus_STARTING}
	if s, err := getSummaries(m.q); err == nil {
		m.summaryCache.refresh(s)
		m.l.Debugf("Got %d initial summaries.", len(s))
		running = true
		m.changes <- agents.Change{Status: inventorypb.AgentStatus_RUNNING}
	} else {
		m.l.Error(err)
		m.changes <- agents.Change{Status: inventorypb.AgentStatus_WAITING}
	}

	go m.runHistoryCacheRefresher(ctx)

	// query events_statements_summary_by_digest every minute at 00 seconds
	start := time.Now()
	wait := start.Truncate(querySummaries).Add(querySummaries).Sub(start)
	m.l.Debugf("Scheduling next collection in %s at %s.", wait, start.Add(wait).Format("15:04:05"))
	t := time.NewTimer(wait)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			m.changes <- agents.Change{Status: inventorypb.AgentStatus_STOPPING}
			m.l.Infof("Context canceled.")
			return

		case <-t.C:
			if !running {
				m.changes <- agents.Change{Status: inventorypb.AgentStatus_STARTING}
			}

			lengthS := uint32(math.Round(wait.Seconds())) // round 59.9s/60.1s to 60s
			buckets, err := m.getNewBuckets(start, lengthS)

			start = time.Now()
			wait = start.Truncate(querySummaries).Add(querySummaries).Sub(start)
			m.l.Debugf("Scheduling next collection in %s at %s.", wait, start.Add(wait).Format("15:04:05"))
			t.Reset(wait)

			if err != nil {
				m.l.Error(err)
				running = false
				m.changes <- agents.Change{Status: inventorypb.AgentStatus_WAITING}
				continue
			}

			if !running {
				running = true
				m.changes <- agents.Change{Status: inventorypb.AgentStatus_RUNNING}
			}

			m.changes <- agents.Change{MetricsBucket: buckets}
		}
	}
}

func (m *PerfSchema) runHistoryCacheRefresher(ctx context.Context) {
	t := time.NewTicker(refreshHistory)
	defer t.Stop()

	for {
		if err := m.refreshHistoryCache(); err != nil {
			m.l.Error(err)
		}

		select {
		case <-ctx.Done():
			return
		case <-t.C:
			// nothing, continue loop
		}
	}
}

func (m *PerfSchema) refreshHistoryCache() error {
	current, err := getHistory(m.q)
	if err != nil {
		return err
	}
	m.historyCache.refresh(current)
	return nil
}

func (m *PerfSchema) getNewBuckets(periodStart time.Time, periodLengthSecs uint32) ([]*agentpb.MetricsBucket, error) {
	current, err := getSummaries(m.q)
	if err != nil {
		return nil, err
	}
	prev := m.summaryCache.get()

	buckets := makeBuckets(current, prev, m.l)
	startS := uint32(periodStart.Unix())
	m.l.Debugf("Made %d buckets out of %d summaries in %s+%d interval.",
		len(buckets), len(current), periodStart.Format("15:04:05"), periodLengthSecs)

	// merge prev and current in cache
	m.summaryCache.refresh(current)

	// add agent_id, timestamps, and examples from history cache
	history := m.historyCache.get()
	for i, b := range buckets {
		b.Common.AgentId = m.agentID
		b.Common.PeriodStartUnixSecs = startS
		b.Common.PeriodLengthSecs = periodLengthSecs

		if esh := history[b.Common.Queryid]; esh != nil {
			// TODO test if we really need that
			// If we don't need it, we can avoid polling events_statements_history completely
			// if query examples are disabled.
			if b.Common.Schema == "" {
				b.Common.Schema = pointer.GetString(esh.CurrentSchema)
			}

			if !m.disableQueryExamples && esh.SQLText != nil {
				b.Common.Example = *esh.SQLText
				b.Common.ExampleFormat = agentpb.ExampleFormat_EXAMPLE
				b.Common.ExampleType = agentpb.ExampleType_RANDOM
			}
		}

		buckets[i] = b
	}

	return buckets, nil
}

// inc returns increment from prev to current, or 0, if there was a wrap-around.
func inc(current, prev uint64) float32 {
	if current <= prev {
		return 0
	}
	return float32(current - prev)
}

// makeBuckets uses current state of events_statements_summary_by_digest table and accumulated previous state
// to make metrics buckets.
//
// makeBuckets is a pure function for easier testing.
func makeBuckets(current, prev map[string]*eventsStatementsSummaryByDigest, l *logrus.Entry) []*agentpb.MetricsBucket {
	res := make([]*agentpb.MetricsBucket, 0, len(current))

	for digest, currentESS := range current {
		prevESS := prev[digest]
		if prevESS == nil {
			prevESS = new(eventsStatementsSummaryByDigest)
		}

		switch {
		case currentESS.CountStar == prevESS.CountStar:
			// TODO
			// Another way how this is possible is if events_statements_summary_by_digest was truncated,
			// and then the same number of queries were made.
			// Currently, we can't differentiate between those situations.
			// We probably could by using first_seen/last_seen columns.
			l.Debugf("Skipped due to the same number of queries: %s.", currentESS)
			continue
		case currentESS.CountStar < prevESS.CountStar:
			l.Debugf("Truncate detected. Treating as a new query: %s.", currentESS)
			prevESS = new(eventsStatementsSummaryByDigest)
		case prevESS.CountStar == 0:
			l.Debugf("New query: %s.", currentESS)
		default:
			l.Debugf("Normal query: %s.", currentESS)
		}

		count := inc(currentESS.CountStar, prevESS.CountStar)
		mb := &agentpb.MetricsBucket{
			Common: &agentpb.MetricsBucket_Common{
				Schema:                 pointer.GetString(currentESS.SchemaName), // TODO can it be NULL?
				Queryid:                *currentESS.Digest,
				Fingerprint:            *currentESS.DigestText,
				NumQueries:             count,
				NumQueriesWithErrors:   inc(currentESS.SumErrors, prevESS.SumErrors),
				NumQueriesWithWarnings: inc(currentESS.SumWarnings, prevESS.SumWarnings),
				AgentType:              inventorypb.AgentType_QAN_MYSQL_PERFSCHEMA_AGENT,
			},
			Mysql: &agentpb.MetricsBucket_MySQL{},
		}

		for _, p := range []struct {
			value float32  // result value: currentESS.SumXXX-prevESS.SumXXX
			sum   *float32 // MetricsBucket.XXXSum field to write value
			cnt   *float32 // MetricsBucket.XXXCnt field to write count
		}{
			// in order of events_statements_summary_by_digest columns

			// convert picoseconds to seconds
			{inc(currentESS.SumTimerWait, prevESS.SumTimerWait) / 1e12, &mb.Common.MQueryTimeSum, &mb.Common.MQueryTimeCnt},
			{inc(currentESS.SumLockTime, prevESS.SumLockTime) / 1e12, &mb.Mysql.MLockTimeSum, &mb.Mysql.MLockTimeCnt},

			{inc(currentESS.SumRowsAffected, prevESS.SumRowsAffected), &mb.Mysql.MRowsAffectedSum, &mb.Mysql.MRowsAffectedCnt},
			{inc(currentESS.SumRowsSent, prevESS.SumRowsSent), &mb.Mysql.MRowsSentSum, &mb.Mysql.MRowsSentCnt},
			{inc(currentESS.SumRowsExamined, prevESS.SumRowsExamined), &mb.Mysql.MRowsExaminedSum, &mb.Mysql.MRowsExaminedCnt},

			{inc(currentESS.SumCreatedTmpDiskTables, prevESS.SumCreatedTmpDiskTables), &mb.Mysql.MTmpDiskTablesSum, &mb.Mysql.MTmpDiskTablesCnt},
			{inc(currentESS.SumCreatedTmpTables, prevESS.SumCreatedTmpTables), &mb.Mysql.MTmpTablesSum, &mb.Mysql.MTmpTablesCnt},
			{inc(currentESS.SumSelectFullJoin, prevESS.SumSelectFullJoin), &mb.Mysql.MFullJoinSum, &mb.Mysql.MFullJoinCnt},
			{inc(currentESS.SumSelectFullRangeJoin, prevESS.SumSelectFullRangeJoin), &mb.Mysql.MSelectFullRangeJoinSum, &mb.Mysql.MSelectFullRangeJoinCnt},
			{inc(currentESS.SumSelectRange, prevESS.SumSelectRange), &mb.Mysql.MSelectRangeSum, &mb.Mysql.MSelectRangeCnt},
			{inc(currentESS.SumSelectRangeCheck, prevESS.SumSelectRangeCheck), &mb.Mysql.MSelectRangeCheckSum, &mb.Mysql.MSelectRangeCheckCnt},
			{inc(currentESS.SumSelectScan, prevESS.SumSelectScan), &mb.Mysql.MFullScanSum, &mb.Mysql.MFullScanCnt},

			{inc(currentESS.SumSortMergePasses, prevESS.SumSortMergePasses), &mb.Mysql.MMergePassesSum, &mb.Mysql.MMergePassesCnt},
			{inc(currentESS.SumSortRange, prevESS.SumSortRange), &mb.Mysql.MSortRangeSum, &mb.Mysql.MSortRangeCnt},
			{inc(currentESS.SumSortRows, prevESS.SumSortRows), &mb.Mysql.MSortRowsSum, &mb.Mysql.MSortRowsCnt},
			{inc(currentESS.SumSortScan, prevESS.SumSortScan), &mb.Mysql.MSortScanSum, &mb.Mysql.MSortScanCnt},

			{inc(currentESS.SumNoIndexUsed, prevESS.SumNoIndexUsed), &mb.Mysql.MNoIndexUsedSum, &mb.Mysql.MNoIndexUsedCnt},
			{inc(currentESS.SumNoGoodIndexUsed, prevESS.SumNoGoodIndexUsed), &mb.Mysql.MNoGoodIndexUsedSum, &mb.Mysql.MNoGoodIndexUsedCnt},
		} {
			if p.value != 0 {
				*p.sum = p.value
				*p.cnt = count
			}
		}

		res = append(res, mb)
	}

	return res
}

// Changes returns channel that should be read until it is closed.
func (m *PerfSchema) Changes() <-chan agents.Change {
	return m.changes
}
