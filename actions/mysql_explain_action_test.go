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

package actions

import (
	"context"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/percona/pmm/api/agentpb"
	"github.com/stretchr/objx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/percona/pmm-agent/utils/tests"
)

func TestMySQLExplain(t *testing.T) {
	t.Parallel()

	dsn := tests.GetTestMySQLDSN(t)
	db := tests.OpenTestMySQL(t)
	defer db.Close() //nolint:errcheck
	mySQLVersion, mySQLVendor := tests.MySQLVersion(t, db)

	const query = "SELECT * FROM city ORDER BY Population"

	t.Run("Default", func(t *testing.T) {
		params := &agentpb.StartActionRequest_MySQLExplainParams{
			Dsn:          dsn,
			Query:        query,
			OutputFormat: agentpb.MysqlExplainOutputFormat_MYSQL_EXPLAIN_OUTPUT_FORMAT_DEFAULT,
		}
		a := NewMySQLExplainAction("", params)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		b, err := a.Run(ctx)
		require.NoError(t, err)

		var expected string
		switch {
		case mySQLVersion == "5.6" || mySQLVendor == tests.MariaDBMySQL:
			expected = strings.TrimSpace(`
id |select_type |table |type |possible_keys |key  |key_len |ref  |rows |Extra
1  |SIMPLE      |city  |ALL  |NULL          |NULL |NULL    |NULL |4188 |Using filesort
			`)
		default:
			expected = strings.TrimSpace(`
id |select_type |table |partitions |type |possible_keys |key  |key_len |ref  |rows |filtered |Extra
1  |SIMPLE      |city  |NULL       |ALL  |NULL          |NULL |NULL    |NULL |4188 |100.00   |Using filesort
			`)
		}
		actual := strings.TrimSpace(string(b))
		assert.Equal(t, expected, actual)
	})

	t.Run("JSON", func(t *testing.T) {
		params := &agentpb.StartActionRequest_MySQLExplainParams{
			Dsn:          dsn,
			Query:        query,
			OutputFormat: agentpb.MysqlExplainOutputFormat_MYSQL_EXPLAIN_OUTPUT_FORMAT_JSON,
		}
		a := NewMySQLExplainAction("", params)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		b, err := a.Run(ctx)
		require.NoError(t, err)
		t.Logf("Full JSON:\n%s", b)
		m, err := objx.FromJSON(string(b))
		require.NoError(t, err)

		assert.Equal(t, 1.0, m.Get("query_block.select_id").Float64())

		var table map[string]interface{}
		switch mySQLVendor {
		case tests.MariaDBMySQL:
			table = m.Get("query_block.read_sorted_file.filesort.table").MSI()
		default:
			table = m.Get("query_block.ordering_operation.table").MSI()
		}

		require.NotNil(t, table)
		assert.Equal(t, "city", table["table_name"])
		if mySQLVersion != "5.6" && mySQLVendor != tests.MariaDBMySQL {
			assert.Equal(t, []interface{}{"ID", "Name", "CountryCode", "District", "Population"}, table["used_columns"])
		}

		if mySQLVendor != tests.MariaDBMySQL {
			require.Len(t, m.Get("warnings").InterSlice(), 1)
			assert.Equal(t, 1003.0, m.Get("warnings[0].Code").Float64())
			assert.Equal(t, "Note", m.Get("warnings[0].Level").String())
			assert.Contains(t, m.Get("warnings[0].Message").String(), "/* select#1 */")
		}
	})

	t.Run("TraditionalJSON", func(t *testing.T) {
		params := &agentpb.StartActionRequest_MySQLExplainParams{
			Dsn:          dsn,
			Query:        query,
			OutputFormat: agentpb.MysqlExplainOutputFormat_MYSQL_EXPLAIN_OUTPUT_FORMAT_TRADITIONAL_JSON,
		}
		a := NewMySQLExplainAction("", params)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		b, err := a.Run(ctx)
		require.NoError(t, err)
		t.Logf("Full JSON:\n%s", b)

		var actual [][]interface{}
		err = json.Unmarshal(b, &actual)
		require.NoError(t, err)
		require.Len(t, actual, 2)

		switch {
		case mySQLVersion == "5.6" || mySQLVendor == tests.MariaDBMySQL:
			assert.Equal(t, []interface{}{
				"id", "select_type", "table",
				"type", "possible_keys", "key", "key_len", "ref", "rows", "Extra",
			}, actual[0])
			assert.Equal(t, []interface{}{"1", "SIMPLE", "city", "ALL", nil, nil, nil, nil, "4188", "Using filesort"}, actual[1])
		default:
			assert.Equal(t, []interface{}{
				"id", "select_type", "table", "partitions",
				"type", "possible_keys", "key", "key_len", "ref", "rows", "filtered", "Extra",
			}, actual[0])
			assert.Equal(t, []interface{}{"1", "SIMPLE", "city", nil, "ALL", nil, nil, nil, nil, "4188", "100.00", "Using filesort"}, actual[1])
		}
	})

	t.Run("Error", func(t *testing.T) {
		params := &agentpb.StartActionRequest_MySQLExplainParams{
			Dsn:          "pmm-agent:pmm-agent-wrong-password@tcp(127.0.0.1:3306)/world",
			OutputFormat: agentpb.MysqlExplainOutputFormat_MYSQL_EXPLAIN_OUTPUT_FORMAT_DEFAULT,
		}
		a := NewMySQLExplainAction("", params)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		_, err := a.Run(ctx)
		require.Error(t, err)
		assert.Regexp(t, `Error 1045: Access denied for user 'pmm-agent'@'.+' \(using password: YES\)`, err.Error())
	})

	t.Run("LittleBobbyTables", func(t *testing.T) {
		checkCity := func(t *testing.T) {
			var count int
			err := db.QueryRow("SELECT COUNT(*) FROM city").Scan(&count)
			require.NoError(t, err)
			assert.Equal(t, 4079, count)
		}

		t.Run("Drop", func(t *testing.T) {
			params := &agentpb.StartActionRequest_MySQLExplainParams{
				Dsn:          dsn,
				Query:        `SELECT 1; DROP TABLE city; --`,
				OutputFormat: agentpb.MysqlExplainOutputFormat_MYSQL_EXPLAIN_OUTPUT_FORMAT_DEFAULT,
			}
			a := NewMySQLExplainAction("", params)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			_, err := a.Run(ctx)
			expected := "Error 1064: You have an error in your SQL syntax; check the manual that corresponds " +
				"to your (MySQL|MariaDB) server version for the right syntax to use near 'DROP TABLE city; --' at line 1"
			require.Error(t, err)
			assert.Regexp(t, expected, err.Error())
			checkCity(t)
		})

		t.Run("Delete", func(t *testing.T) {
			params := &agentpb.StartActionRequest_MySQLExplainParams{
				Dsn:          dsn,
				Query:        `DELETE FROM city`,
				OutputFormat: agentpb.MysqlExplainOutputFormat_MYSQL_EXPLAIN_OUTPUT_FORMAT_DEFAULT,
			}
			a := NewMySQLExplainAction("", params)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			_, err := a.Run(ctx)
			require.NoError(t, err)
			checkCity(t)
		})
	})
}
