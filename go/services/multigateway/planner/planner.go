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

// Package planner handles query planning for multigateway.
// It analyzes SQL queries and creates execution plans with appropriate primitives.
package planner

import (
	"log/slog"

	"github.com/multigres/multigres/go/common/parser/ast"
	"github.com/multigres/multigres/go/common/pgprotocol/server"
	"github.com/multigres/multigres/go/common/preparedstatement"
	"github.com/multigres/multigres/go/services/multigateway/engine"
)

// Planner is responsible for creating query execution plans.
type Planner struct {
	// defaultTableGroup is the tablegroup to use when routing queries.
	// For Phase 1, all queries are routed to this tablegroup.
	defaultTableGroup string

	logger *slog.Logger

	// txnMetrics is injected into TransactionPrimitive at creation time
	// for recording transaction duration and count.
	txnMetrics *engine.TransactionMetrics
}

// NewPlanner creates a new query planner.
func NewPlanner(defaultTableGroup string, logger *slog.Logger, txnMetrics *engine.TransactionMetrics) *Planner {
	_ = "STUB: not implemented"
	return nil
}

// Plan creates an execution plan for the given SQL query and AST.
//
// The planner analyzes the AST to determine query type and creates
// appropriate primitives. Uses PostgreSQL's utility.c dispatch pattern
// with switch on NodeTag for extensibility.
//
// Supported statement types:
// - VariableSetStmt: SET/RESET commands → ApplySessionState
// - Regular queries: Route only
//
// Future phases will add more statement handlers for:
// - TransactionStmt: BEGIN/COMMIT/ROLLBACK
// - SelectStmt: Query optimization and sharding
// - InsertStmt/UpdateStmt/DeleteStmt: Write operations
func (p *Planner) Plan(
	sql string,
	stmt ast.Stmt,
	conn *server.Conn,
) (*engine.Plan, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reject unsupported constructs before dispatch: Tier 2 statement types
// (LOAD, ALTER SYSTEM, CREATE/DROP DATABASE, etc.) plus any blocklisted
// or misplaced FuncCalls in expression trees. Running here (not in the
// executor) means the plan cache short-circuits both checks: a cached
// plan is by construction safe. The normalizer is configured to
// preserve literals inside set_config calls so its args remain A_Const
// at this point.

// Handle wrapped EXECUTE forms (EXPLAIN EXECUTE / CREATE TABLE AS EXECUTE)
// before normal dispatch. The wrapper's inner ExecuteStmt references a
// gateway-managed prepared statement by user-facing name (e.g. "p"); we
// rewrite it to the canonical name (e.g. "stmt42") and attach the
// PreparedStatement metadata so the multipooler can ensurePrepared() on
// the backend connection before running the query. See execute_unwrap.go.

// Dispatch to appropriate planner function based on statement type
// This follows PostgreSQL's utility.c pattern with switch on node tag

// Default: simple route to PostgreSQL

// planTempTableCreation creates a plan that routes through a reserved
// connection with ReasonTempTable. The reservation ensures the temp table
// persists across queries on the same session.
func (p *Planner) planTempTableCreation(sql string, conn *server.Conn) (*engine.Plan, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// planHoldCursorDeclare creates a plan for `DECLARE ... WITH HOLD` cursors.
// WITH HOLD promotes the cursor to session-level state that must survive
// COMMIT; the cursor name is pinned on the reserved backend connection via
// ReasonPortal so the multipooler does not return the backend to the pool
// when the surrounding transaction commits.
func (p *Planner) planHoldCursorDeclare(sql string, stmt *ast.DeclareCursorStmt) (*engine.Plan, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// planClosePortalStmt creates a plan for `CLOSE <name>` / `CLOSE ALL`. The
// CloseCursorRoute forwards the CLOSE to PostgreSQL on the existing reserved
// backend and, when the named cursor was a `WITH HOLD` pin, asks the
// multipooler to drop the corresponding entry from the reserved
// connection's portal set.
func (p *Planner) planClosePortalStmt(sql string, stmt *ast.ClosePortalStmt) (*engine.Plan, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// planDefault creates a simple route plan for queries without special handling.
// This is the fallback for most SQL statements.
func (p *Planner) planDefault(sql string, stmt ast.Stmt, conn *server.Conn) (*engine.Plan, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PlanPortal creates an execution plan for the extended query protocol (portal path).
// Unlike Plan, which handles all statements, PlanPortal only returns a non-nil plan
// for statements that require local handling by the gateway. For all other statements,
// it returns (nil, nil) to indicate they should be sent to PostgreSQL via
// PortalStreamExecute with the portal's bound parameters.
//
// Statements that produce a plan are delegated to Plan to reuse existing planning logic.
// This covers any statement whose semantics cannot be preserved by a plain portal
// execute on a pooled backend connection — for example, gateway-managed session
// variables, LISTEN/UNLISTEN/NOTIFY, DISCARD, temp-table creation, and
// BEGIN/COMMIT/ROLLBACK.
func (p *Planner) PlanPortal(
	portalInfo *preparedstatement.PortalInfo,
	conn *server.Conn,
) (*engine.Plan, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Non-cacheable extended-protocol statements reach PlanPortal directly
// (cacheable ones go through resolvePortalPlan → Plan, which does the
// same checks), so both paths must share the same pre-dispatch rejection.
// We throw away the set_config result here: PlanPortal only routes
// gateway-local statement types, none of which are SELECTs that could
// carry tracked set_configs.

// DISCARD TEMP needs the DiscardTempPrimitive for reservation cleanup.

// BEGIN/COMMIT/ROLLBACK must run through the gateway's transaction
// primitive — executing them as a normal portal on a pooled backend
// connection leaks open (or aborted) transactions across clients when
// the connection is recycled.

// DECLARE … WITH HOLD must go through HoldCursorRoute so the cursor
// name is pinned on the reserved backend (ReasonPortal). Without
// this case, an extended-protocol DECLARE WITH HOLD would land on a
// pooled connection and the cursor would be lost on COMMIT.
// Non-HOLD DECLARE is delegated through Plan too so the parser-driven
// dispatch decides — non-HOLD falls through to planDefault there.

// CLOSE / CLOSE ALL must go through CloseCursorRoute so HOLD-cursor
// pin bookkeeping on the multipooler stays in sync — otherwise the
// reserved backend would leak with a stale ReasonPortal.

// SetDefaultTableGroup updates the default tablegroup for routing.
// This allows dynamic configuration changes.
func (p *Planner) SetDefaultTableGroup(tableGroup string) { _ = "STUB: not implemented"; return }

// GetDefaultTableGroup returns the current default tablegroup.
func (p *Planner) GetDefaultTableGroup() string { _ = "STUB: not implemented"; return "" }

// primitiveName returns a short string identifying the primitive type.
// Used for observability (span attributes and query logs).
func primitiveName(p engine.Primitive) string { _ = "STUB: not implemented"; return "" }

// planListenStmt creates a ListenNotify primitive for LISTEN.
func (p *Planner) planListenStmt(sql string, stmt *ast.ListenStmt) (*engine.Plan, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// planUnlistenStmt creates a ListenNotify primitive for UNLISTEN.
func (p *Planner) planUnlistenStmt(sql string, stmt *ast.UnlistenStmt) (*engine.Plan, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// planNotifyStmt routes NOTIFY to the default table group as a regular query.
func (p *Planner) planNotifyStmt(sql string) (*engine.Plan, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
