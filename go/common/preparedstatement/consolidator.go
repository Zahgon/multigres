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

package preparedstatement

import (
	"sync"

	"github.com/multigres/multigres/go/common/parser/ast"
	querypb "github.com/multigres/multigres/go/pb/query"
)

// Consolidator is used to consolidate prepared statements that
// are preparing the same statement but with different names. The intent is to be able to
// use the same connection for both of them to execute this because underlying, they are using
// the same prepared statement.
type Consolidator struct {
	// Mutex to protect the fields
	mu sync.Mutex

	// Map from (query, paramTypes) dedup key to canonical prepared statement
	stmts map[string]*PreparedStatementInfo
	// Map from connection ID and statement name to prepared statement reference
	incoming map[uint32]map[string]*PreparedStatementInfo
	// Reference count: number of connections using each prepared statement
	usageCount map[*PreparedStatementInfo]int

	// lastUsedID is the last id of the statement name that we used.
	lastUsedID int
}

// ConsolidatorStats contains statistics about the prepared statement consolidator.
type ConsolidatorStats struct {
	// UniqueStatements is the number of unique prepared statements being tracked.
	UniqueStatements int `json:"unique_statements"`
	// TotalReferences is the total number of references across all connections.
	TotalReferences int `json:"total_references"`
	// ConnectionCount is the number of connections that have prepared statements.
	ConnectionCount int `json:"connection_count"`
	// Statements contains details about each unique prepared statement.
	Statements []StatementStats `json:"statements"`
}

// StatementStats contains statistics for a single prepared statement.
type StatementStats struct {
	// Name is the canonical name of the prepared statement.
	Name string `json:"name"`
	// Query is the SQL query of the prepared statement.
	Query string `json:"query"`
	// UsageCount is the number of connections using this prepared statement.
	UsageCount int `json:"usage_count"`
}

type PortalInfo struct {
	*querypb.Portal
	*PreparedStatementInfo
}

type PreparedStatementInfo struct {
	*querypb.PreparedStatement
	astStruct ast.Stmt
}

// AstStmt returns the parsed AST statement for this prepared statement.
func (psi *PreparedStatementInfo) AstStmt() ast.Stmt {
	_ = "STUB: not implemented"
	return *

	// NewPreparedStatementInfo parses the query in the prepared statement and stores it along with the
	// prepared statement information for future use.
	new(ast.Stmt)
}

func NewPreparedStatementInfo(ps *querypb.PreparedStatement) (*PreparedStatementInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewPortalInfo creates the PortalInfo.
func NewPortalInfo(psi *PreparedStatementInfo, portal *querypb.Portal) *PortalInfo {
	_ = "STUB: not implemented"
	return nil
}

// NewConsolidator gets a new prepared statement consolidator
// used to consolidate and reuse the same prepared statements.
func NewConsolidator() *Consolidator { _ = "STUB: not implemented"; return nil }

// AddPreparedStatement adds a prepared statement to the consolidator.
// Returns the PreparedStatementInfo (either existing or newly created) and any error.
func (psc *Consolidator) AddPreparedStatement(connId uint32, name, queryStr string, paramTypes []uint32) (*PreparedStatementInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initialize the map for this connection if it doesn't exist

// If the name is non-empty and a prepared statement for this name already exists
// on the connection, replace it. This matches PostgreSQL behavior where re-parsing
// with an existing name replaces the old statement. This is necessary to handle
// the case where Parse succeeds (adding to consolidator) but the subsequent
// Describe fails — the client retries Parse with the same name.

// Let's check if a prepared statement with this (query, paramTypes) already exists.

// We found an existing prepared statement, we should be using that.

// We didn't find any existing prepared statement with this (query, paramTypes).
// Create a new one in our stmts list tracking unique prepared statements.

// GetPreparedStatementInfo gets the information for a previously added prepared statement to the consolidator.
func (psc *Consolidator) GetPreparedStatementInfo(connId uint32, name string) *PreparedStatementInfo {
	_ = "STUB: not implemented"
	return nil
}

// RemovePreparedStatement removes prepared statement.
func (psc *Consolidator) RemovePreparedStatement(connId uint32, name string) {
	_ = "STUB: not implemented"
	return
}

// RemoveConnection removes all prepared statements associated with a connection.
// This should be called when a client connection is closed.
func (psc *Consolidator) RemoveConnection(connId uint32) { _ = "STUB: not implemented"; return }

// Stats returns statistics about the consolidator's current state.
func (psc *Consolidator) Stats() ConsolidatorStats {
	_ = "STUB: not implemented"
	return *new(ConsolidatorStats)
}
