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

// Package process runs Agent processes.
package process

import (
	"context"
	"os/exec"
	"time"

	"github.com/percona/pmm/api/inventorypb"
	"github.com/percona/pmm/utils/pdeathsig"
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"

	"github.com/percona/pmm-agent/utils/backoff"
)

const (
	runningT = time.Second     // STARTING -> RUNNING delay
	killT    = 5 * time.Second // SIGTERM -> SIGKILL delay

	backoffMinDelay = 1 * time.Second
	backoffMaxDelay = 30 * time.Second

	keepLogLines = 100
)

// Process represents Agent process started by pmm-agent.
//
// Process object should be created with New and started with Run (typically in a separate goroutine).
// It then handles process starting, restarting with backoff, reading its output.
// Process is gracefully stopped when context passed to New is canceled.
// Changes of process status are reported via Changes channel which must be read until it is closed.
//
// Process status is changed by finite state machine (see agent_status.dot).
// Each state logic is encapsulated in toXXX methods. Each method sends a new status to the changes channel,
// implements its own logic, and then switches to then next state via "go toXXX()". "go" statement is used
// only to avoid stack overflow; there are no extra goroutines for states.
type Process struct {
	params  *Params
	l       *logrus.Entry
	pl      *processLogger
	changes chan inventorypb.AgentStatus
	backoff *backoff.Backoff
	ctxDone chan struct{}

	// recreated on each restart
	cmd     *exec.Cmd
	cmdDone chan struct{}
}

// Params represent Agent process parameters: command path, command-line arguments/flags, and process environment.
type Params struct {
	Path string
	Args []string
	Env  []string
}

// New creates new process.
func New(params *Params, l *logrus.Entry) *Process {
	return &Process{
		params:  params,
		l:       l,
		pl:      newProcessLogger(l, keepLogLines),
		changes: make(chan inventorypb.AgentStatus, 10),
		backoff: backoff.New(backoffMinDelay, backoffMaxDelay),
		ctxDone: make(chan struct{}),
	}
}

// Run starts process and runs until ctx is canceled.
func (p *Process) Run(ctx context.Context) {
	go p.toStarting()

	<-ctx.Done()
	p.l.Infof("Process: context canceled.")
	close(p.ctxDone)
}

// STARTING -> RUNNING
// STARTING -> WAITING
func (p *Process) toStarting() {
	p.l.Infof("Process: starting.")
	p.changes <- inventorypb.AgentStatus_STARTING

	p.cmd = exec.Command(p.params.Path, p.params.Args...) //nolint:gosec
	p.cmd.Stdout = p.pl
	p.cmd.Stderr = p.pl

	// restrict process
	p.cmd.Env = p.params.Env
	if p.cmd.Env == nil {
		p.cmd.Env = []string{} // never inherit environment
	}
	p.cmd.Dir = "/"
	pdeathsig.Set(p.cmd, unix.SIGKILL)

	p.cmdDone = make(chan struct{})

	if err := p.cmd.Start(); err != nil {
		p.l.Warnf("Process: failed to start: %s.", err)
		go p.toWaiting()
		return
	}

	go func() {
		// p.cmd.ProcessState is checked once cmdDone is closed, so error there can be ignored
		_ = p.cmd.Wait()
		close(p.cmdDone)
	}()

	t := time.NewTimer(runningT)
	defer t.Stop()
	select {
	case <-t.C:
		go p.toRunning()
	case <-p.cmdDone:
		p.l.Warnf("Process: exited early: %s.", p.cmd.ProcessState)
		go p.toWaiting()
	}
}

// RUNNING -> STOPPING
// RUNNING -> WAITING
func (p *Process) toRunning() {
	p.l.Infof("Process: running.")
	p.changes <- inventorypb.AgentStatus_RUNNING

	p.backoff.Reset()

	select {
	case <-p.ctxDone:
		go p.toStopping()
	case <-p.cmdDone:
		p.l.Warnf("Process: exited: %s.", p.cmd.ProcessState)
		go p.toWaiting()
	}
}

// WAITING -> STARTING
// WAITING -> DONE
func (p *Process) toWaiting() {
	delay := p.backoff.Delay()

	p.l.Infof("Process: waiting %s.", delay)
	p.changes <- inventorypb.AgentStatus_WAITING

	t := time.NewTimer(delay)
	defer t.Stop()
	select {
	case <-t.C:
		go p.toStarting()
	case <-p.ctxDone:
		go p.toDone()
	}
}

// STOPPING -> DONE
func (p *Process) toStopping() {
	p.l.Infof("Process: stopping (sending SIGTERM)...")
	p.changes <- inventorypb.AgentStatus_STOPPING

	if err := p.cmd.Process.Signal(unix.SIGTERM); err != nil {
		p.l.Errorf("Process: failed to send SIGTERM: %s.", err)
	}

	t := time.NewTimer(killT)
	defer t.Stop()
	select {
	case <-p.cmdDone:
		// nothing
	case <-t.C:
		p.l.Warnf("Process: still alive after %s, sending SIGKILL...", killT)
		if err := p.cmd.Process.Signal(unix.SIGKILL); err != nil {
			p.l.Errorf("Process: failed to send SIGKILL: %s.", err)
		}
		<-p.cmdDone
	}

	p.l.Infof("Process: exited: %s.", p.cmd.ProcessState)
	go p.toDone()
}

func (p *Process) toDone() {
	p.l.Info("Process: done.")
	p.changes <- inventorypb.AgentStatus_DONE

	close(p.changes)
}

// Changes returns channel that should be read until it is closed.
func (p *Process) Changes() <-chan inventorypb.AgentStatus {
	return p.changes
}

// Logs returns latest process logs.
func (p *Process) Logs() []string {
	return p.pl.Latest()
}
