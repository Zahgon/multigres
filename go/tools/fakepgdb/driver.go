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

package fakepgdb

import (
	"context"
	"database/sql/driver"
)

// fakeDriver implements driver.Driver
type fakeDriver struct {
	db *DB
}

// Open returns a new connection to the fake database.
func (d *fakeDriver) Open(name string) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

// fakeConn implements driver.Conn
type fakeConn struct {
	db *DB
}

// Prepare returns a prepared statement, bound to this connection.
func (c *fakeConn) Prepare(query string) (driver.Stmt, error) {
	_ = "STUB: not implemented"
	return *new(driver.Stmt), nil
}

// Close closes the connection.
func (c *fakeConn) Close() error {
	_ = "STUB: not implemented"

	// Begin starts and returns a new transaction.
	return nil
}

func (c *fakeConn) Begin() (driver.Tx, error) {
	_ = "STUB: not implemented"
	return *new(driver.Tx), nil
}

// QueryContext executes a query that may return rows.
func (c *fakeConn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	_ = "STUB: not implemented"
	// Convert args to []interface{} for compatibility
	return *new(driver.Rows), nil
}

// ExecContext executes a query that doesn't return rows.
func (c *fakeConn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	_ = "STUB: not implemented"
	// Convert args to []interface{} for compatibility
	return *new(driver.Result), nil
}

// fakeStmt implements driver.Stmt
type fakeStmt struct {
	conn  *fakeConn
	query string
}

// Close closes the statement.
func (s *fakeStmt) Close() error {
	_ = "STUB: not implemented"

	// NumInput returns the number of placeholder parameters.
	return nil
}

func (s *fakeStmt) NumInput() int {
	_ = "STUB: not implemented"
	// -1 means the driver doesn't know
	return 0
}

// Exec executes a query that doesn't return rows.
func (s *fakeStmt) Exec(args []driver.Value) (driver.Result, error) {
	_ = "STUB: not implemented"
	return *new(driver.Result), nil
}

// Query executes a query that may return rows.
func (s *fakeStmt) Query(args []driver.Value) (driver.Rows, error) {
	_ = "STUB: not implemented"
	return *new(driver.Rows), nil
}

// fakeTx implements driver.Tx
type fakeTx struct {
	conn *fakeConn
}

// Commit commits the transaction.
func (tx *fakeTx) Commit() error {
	_ = "STUB: not implemented"

	// Rollback aborts the transaction.
	return nil
}

func (tx *fakeTx) Rollback() error {
	_ = "STUB: not implemented"

	// fakeResult implements driver.Result
	return nil
}

type fakeResult struct {
	lastInsertId int64
	rowsAffected int64
}

// LastInsertId returns the database's auto-generated ID.
func (r *fakeResult) LastInsertId() (int64, error) {
	_ = "STUB: not implemented"
	return 0,

		// RowsAffected returns the number of rows affected by the query.
		nil
}

func (r *fakeResult) RowsAffected() (int64, error) {
	_ = "STUB: not implemented"
	return 0,

		// fakeRows implements driver.Rows
		nil
}

type fakeRows struct {
	columns []string
	rows    [][]any
	index   int
}

// Columns returns the names of the columns.
func (r *fakeRows) Columns() []string {
	_ = "STUB: not implemented"

	// Close closes the rows iterator.
	return nil
}

func (r *fakeRows) Close() error {
	_ = "STUB: not implemented"

	// Next is called to populate the next row of data into the provided slice.
	return nil
}

func (r *fakeRows) Next(dest []driver.Value) error { _ = "STUB: not implemented"; return nil }

// Ensure interfaces are implemented
var (
	_ driver.Driver         = (*fakeDriver)(nil)
	_ driver.Conn           = (*fakeConn)(nil)
	_ driver.QueryerContext = (*fakeConn)(nil)
	_ driver.ExecerContext  = (*fakeConn)(nil)
	_ driver.Stmt           = (*fakeStmt)(nil)
	_ driver.Tx             = (*fakeTx)(nil)
	_ driver.Result         = (*fakeResult)(nil)
	_ driver.Rows           = (*fakeRows)(nil)
)
