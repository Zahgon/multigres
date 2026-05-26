// Copyright 2026 Supabase, Inc.
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

package engine

import (
	"context"

	"github.com/multigres/multigres/go/common/parser/ast"
	"github.com/multigres/multigres/go/common/pgprotocol/server"
	"github.com/multigres/multigres/go/common/preparedstatement"
	"github.com/multigres/multigres/go/common/sqltypes"
	"github.com/multigres/multigres/go/services/multigateway/handler"
)

// preparedStmtKind identifies which prepared statement operation to perform.
type preparedStmtKind int

const (
	preparedStmtPrepare       preparedStmtKind = iota // PREPARE name AS query
	preparedStmtExecute                               // EXECUTE name [(params)]
	preparedStmtDeallocate                            // DEALLOCATE name
	preparedStmtDeallocateAll                         // DEALLOCATE ALL
)

// PreparedStatementPrimitive handles PREPARE, EXECUTE, and DEALLOCATE statements
// from the simple query protocol by delegating to the existing extended query
// protocol handler methods (HandleParse, HandleBind, HandleClose) via conn.Handler().
//
// Key behaviors:
//   - PREPARE: Calls HandleParse to register in the consolidator.
//   - EXECUTE: Calls HandleBind to create a portal, then PortalStreamExecute.
//   - DEALLOCATE: Calls HandleClose to remove from the consolidator.
//   - DEALLOCATE ALL: Uses the consolidator directly (no extended protocol equivalent).
type PreparedStatementPrimitive struct {
	kind       preparedStmtKind
	tableGroup string

	// stmtName is the prepared statement name (used by all kinds).
	stmtName string

	// innerQuery is the SQL body of the PREPARE statement.
	innerQuery string

	// paramTypes holds the parameter type OIDs for PREPARE (from SQL type names).
	paramTypes []uint32

	// params holds the converted EXECUTE parameters (text-format byte arrays).
	params [][]byte
}

// NewPreparePrimitive creates a primitive for PREPARE name AS query.
func NewPreparePrimitive(tableGroup, stmtName, innerQuery string, paramTypes []uint32) *PreparedStatementPrimitive {
	_ = "STUB: not implemented"
	return nil
}

// NewExecutePrimitive creates a primitive for EXECUTE name [(params)].
func NewExecutePrimitive(tableGroup, stmtName string, params [][]byte) *PreparedStatementPrimitive {
	_ = "STUB: not implemented"
	return nil
}

// NewDeallocatePrimitive creates a primitive for DEALLOCATE name.
func NewDeallocatePrimitive(tableGroup, stmtName string) *PreparedStatementPrimitive {
	_ = "STUB: not implemented"
	return nil
}

// NewDeallocateAllPrimitive creates a primitive for DEALLOCATE ALL.
func NewDeallocateAllPrimitive(tableGroup string) *PreparedStatementPrimitive {
	_ = "STUB: not implemented"
	return nil
}

func (p *PreparedStatementPrimitive) StreamExecute(
	ctx context.Context,
	exec IExecute,
	conn *server.Conn,
	state *handler.MultiGatewayConnectionState,
	_ []*ast.A_Const,
	callback func(context.Context, *sqltypes.Result) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

// executePrepare delegates to HandleParse to register the statement in the consolidator.
//
// Unlike the extended Parse message, SQL-level PREPARE must reject a name that is
// already in use on this session. HandleParse silently replaces existing entries
// (to tolerate Parse retries after a failed Describe), so we check for the name
// here before delegating.
func (p *PreparedStatementPrimitive) executePrepare(
	ctx context.Context,
	conn *server.Conn,
	callback func(context.Context, *sqltypes.Result) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

// executeExecute delegates to HandleBind to create a portal, then executes it
// via PortalStreamExecute. We avoid calling HandleExecute directly because it
// would double-count metrics and spans (HandleQuery already records those).
//
// Because the extended protocol returns field metadata via Describe (not Execute),
// PortalStreamExecute results lack Fields. We call HandleDescribe to obtain them
// and inject into the first callback so the simple query protocol can emit the
// required RowDescription before DataRows.
func (p *PreparedStatementPrimitive) executeExecute(
	ctx context.Context,
	exec IExecute,
	conn *server.Conn,
	state *handler.MultiGatewayConnectionState,
	callback func(context.Context, *sqltypes.Result) error,
) error {
	_ = "STUB: not implemented"
	// HandleBind looks up the prepared statement, creates a portal, and stores it in state.
	return nil
}

// Retrieve the portal that HandleBind just created.

// Get field descriptions via Describe. In the extended protocol this is a
// separate step; in the simple protocol it must be folded into the result
// stream so the conn writes RowDescription before DataRows.

// Wrap the callback to inject Fields from Describe into the first result.

// executeDeallocate uses HandleClose with typ 'D' which errors on nonexistent
// statements, matching PostgreSQL's DEALLOCATE behavior.
func (p *PreparedStatementPrimitive) executeDeallocate(
	ctx context.Context,
	conn *server.Conn,
	callback func(context.Context, *sqltypes.Result) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *PreparedStatementPrimitive) executeDeallocateAll(
	ctx context.Context,
	conn *server.Conn,
	callback func(context.Context, *sqltypes.Result) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

// PortalStreamExecute satisfies the Primitive interface for the
// extended-protocol path. PREPARE/EXECUTE/DEALLOCATE statements come in
// via the simple protocol (PlanPortal returns nil for them); the
// EXECUTE form has its own internal portal-style flow that already
// reuses HandleBind / PortalStreamExecute on the backend. Delegate.
func (p *PreparedStatementPrimitive) PortalStreamExecute(
	ctx context.Context,
	exec IExecute,
	conn *server.Conn,
	state *handler.MultiGatewayConnectionState,
	_ *preparedstatement.PortalInfo,
	_ int32,
	_ bool,
	callback func(context.Context, *sqltypes.Result) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *PreparedStatementPrimitive) GetTableGroup() string { _ = "STUB: not implemented"; return "" }
func (p *PreparedStatementPrimitive) GetQuery() string      { _ = "STUB: not implemented"; return "" }
func (p *PreparedStatementPrimitive) String() string        { _ = "STUB: not implemented"; return "" }

var _ Primitive = (*PreparedStatementPrimitive)(nil)

// ExtractExecuteParams converts EXECUTE statement parameters from AST expression
// nodes to byte arrays for portal creation. Only literal constant values (A_Const)
// are supported.
func ExtractExecuteParams(stmt *ast.ExecuteStmt) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExtractParamTypeOids resolves the Argtypes from a PREPARE statement to OIDs.
// Returns nil if there are no argument types. Unrecognized type names are passed
// as 0 (unspecified), letting the backend infer them.
func ExtractParamTypeOids(stmt *ast.PrepareStmt) []uint32 { _ = "STUB: not implemented"; return nil }

// Use the last name component (e.g., "pg_catalog"."int4" → "int4").

// ExtractInnerQuery extracts the SQL string of the inner query from a PrepareStmt.
func ExtractInnerQuery(stmt *ast.PrepareStmt) string { _ = "STUB: not implemented"; return "" }

// constToBytes converts an AST constant value to its text-format byte representation.
func constToBytes(val ast.Value) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
