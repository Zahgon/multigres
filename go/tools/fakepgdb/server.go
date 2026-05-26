// Copyright 2025 Supabase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package fakepgdb provides a fake PostgreSQL server for tests.
// It is inspired by Vitess's fakesqldb but adapted for PostgreSQL.
package fakepgdb

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"regexp"
	"sync"
	"sync/atomic"
	"testing"
)

const (
	appendEntry = -1
)

// DB is a fake PostgreSQL database. All methods are thread-safe.
// It implements driver.Connector to be used with sql.OpenDB.
type DB struct {
	// t is our testing.TB instance
	t testing.TB

	// name is the name of this DB
	name string

	// orderMatters controls whether query order matters
	orderMatters atomic.Bool

	// mu protects all the following fields
	mu sync.Mutex

	// data maps tolower(query) to a result
	data map[string]*ExpectedResult

	// rejectedData maps tolower(query) to an error
	rejectedData map[string]error

	// patternData is a map of regexp queries to results
	patternData map[string]exprResult

	// queryCalled keeps track of how many times a query was called
	queryCalled map[string]int

	// querylog keeps track of all called queries
	querylog []string

	// expectedExecuteFetch is the array of expected queries (for ordered mode)
	expectedExecuteFetch []ExpectedExecuteFetch

	// expectedExecuteFetchIndex is the current index of the query
	expectedExecuteFetchIndex int

	// neverFail makes unmatched queries return empty results instead of errors
	neverFail atomic.Bool

	// allowAll returns empty result for all queries (for benchmarking)
	allowAll atomic.Bool

	// queryPatternUserCallback stores optional callbacks when a query with a pattern is called
	queryPatternUserCallback map[*regexp.Regexp]func(string)
}

// ExpectedResult holds the data for a matched query.
type ExpectedResult struct {
	Columns []string
	Rows    [][]any
	// BeforeFunc() is synchronously called before the server returns the result.
	BeforeFunc func()
}

type exprResult struct {
	queryPattern string
	expr         *regexp.Regexp
	result       *ExpectedResult
	err          string
}

// ExpectedExecuteFetch defines for an expected query the to be faked output.
// It is used for ordered expected output.
type ExpectedExecuteFetch struct {
	Query       string
	QueryResult *ExpectedResult
	Error       error
	// AfterFunc is a callback which is executed while the query
	// is executed i.e., before the fake responds to the client.
	AfterFunc func()
}

// New creates a new fake PostgreSQL database for testing.
func New(t testing.TB) *DB { _ = "STUB: not implemented"; return nil }

// Name returns the name of the DB.
func (db *DB) Name() string { _ = "STUB: not implemented"; return "" }

// SetName sets the name of the DB.
func (db *DB) SetName(name string) *DB { _ = "STUB: not implemented"; return nil }

// OrderMatters sets the orderMatters flag.
func (db *DB) OrderMatters() { _ = "STUB: not implemented"; return }

// Connect returns a driver.Conn implementation.
func (db *DB) Connect(ctx context.Context) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

// Driver returns a driver.Driver implementation.
func (db *DB) Driver() driver.Driver { _ = "STUB: not implemented"; return *new(driver.Driver) }

// OpenDB returns a *sql.DB connected to this fake database.
func (db *DB) OpenDB() *sql.DB { _ = "STUB: not implemented"; return nil }

//
// Methods to add expected queries and results.
//

// AddQuery adds a query and its expected result.
func (db *DB) AddQuery(query string, expectedResult *ExpectedResult) *ExpectedResult {
	_ = "STUB: not implemented"
	return nil
}

// SetBeforeFunc sets the BeforeFunc field for the previously registered "query".
func (db *DB) SetBeforeFunc(query string, f func()) { _ = "STUB: not implemented"; return }

// AddQueryPattern adds an expected result for a set of queries.
// These patterns are checked if no exact matches from AddQuery() are found.
// This function forces the addition of begin/end anchors (^$) and turns on
// case-insensitive matching mode.
func (db *DB) AddQueryPattern(queryPattern string, expectedResult *ExpectedResult) {
	_ = "STUB: not implemented"
	return
}

// RemoveQueryPattern removes a query pattern that was previously added.
func (db *DB) RemoveQueryPattern(queryPattern string) { _ = "STUB: not implemented"; return }

// RejectQueryPattern allows a query pattern to be rejected with an error
func (db *DB) RejectQueryPattern(queryPattern, error string) { _ = "STUB: not implemented"; return }

// ClearQueryPattern removes all query patterns set up
func (db *DB) ClearQueryPattern() { _ = "STUB: not implemented"; return }

// AddQueryPatternWithCallback is similar to AddQueryPattern: in addition it calls the provided callback function
func (db *DB) AddQueryPatternWithCallback(queryPattern string, expectedResult *ExpectedResult, callback func(string)) {
	_ = "STUB: not implemented"
	return
}

// DeleteQuery deletes query from the fake DB.
func (db *DB) DeleteQuery(query string) { _ = "STUB: not implemented"; return }

// DeleteAllQueries deletes all expected queries from the fake DB.
func (db *DB) DeleteAllQueries() { _ = "STUB: not implemented"; return }

// AddRejectedQuery adds a query which will be rejected at execution time.
func (db *DB) AddRejectedQuery(query string, err error) { _ = "STUB: not implemented"; return }

// DeleteRejectedQuery deletes query from the fake DB.
func (db *DB) DeleteRejectedQuery(query string) { _ = "STUB: not implemented"; return }

// GetQueryCalledNum returns how many times db executes a certain query.
func (db *DB) GetQueryCalledNum(query string) int { _ = "STUB: not implemented"; return 0 }

// QueryLog returns the query log as a semicolon separated string
func (db *DB) QueryLog() string { _ = "STUB: not implemented"; return "" }

// ResetQueryLog resets the query log
func (db *DB) ResetQueryLog() { _ = "STUB: not implemented"; return }

//
// Methods for ordered expected queries.
//

// AddExpectedExecuteFetch adds an ExpectedExecuteFetch directly.
func (db *DB) AddExpectedExecuteFetch(entry ExpectedExecuteFetch) {
	_ = "STUB: not implemented"
	return
}

// AddExpectedExecuteFetchAtIndex inserts a new entry at index.
func (db *DB) AddExpectedExecuteFetchAtIndex(index int, entry ExpectedExecuteFetch) {
	_ = "STUB: not implemented"
	return
}

// Grow the slice by one element

// Use copy to move the upper part of the slice out of the way and open a hole

// Store the new value

// AddExpectedQuery adds a single query with no result.
func (db *DB) AddExpectedQuery(query string, err error) { _ = "STUB: not implemented"; return }

// DeleteAllEntries removes all ordered entries.
func (db *DB) DeleteAllEntries() { _ = "STUB: not implemented"; return }

// VerifyAllExecutedOrFail checks that all expected queries were actually executed.
func (db *DB) VerifyAllExecutedOrFail() { _ = "STUB: not implemented"; return }

// SetAllowAll makes all queries return empty results.
func (db *DB) SetAllowAll(allowAll bool) { _ = "STUB: not implemented"; return }

// SetNeverFail makes unmatched queries return empty results instead of errors.
func (db *DB) SetNeverFail(neverFail bool) { _ = "STUB: not implemented"; return }

// handleQuery handles a query and returns the result.
func (db *DB) handleQuery(query string) (*ExpectedResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if we should reject it

// Check explicit queries from AddQuery()

// Check query patterns from AddQueryPattern()

// Nothing matched

func (db *DB) handleQueryOrdered(query string) (*ExpectedResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
