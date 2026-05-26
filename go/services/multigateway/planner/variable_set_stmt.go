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

package planner

import (
	"github.com/multigres/multigres/go/common/parser/ast"
	"github.com/multigres/multigres/go/common/pgprotocol/server"
	"github.com/multigres/multigres/go/services/multigateway/engine"
)

// planVariableSetStmt plans SET/RESET commands.
// Creates an ApplySessionState that handles SET and RESET as local state updates
// with synthetic responses. PG validation is deferred to the next query when
// the pool applies settings to the backend connection.
func (p *Planner) planVariableSetStmt(
	sql string,
	stmt *ast.VariableSetStmt,
	conn *server.Conn,
) (*engine.Plan, error) {
	_ = "STUB: not implemented"
	// Gateway-managed variables are handled locally without routing to PostgreSQL,
	// regardless of whether SET or SET LOCAL is used. This check must come before
	// the IsLocal pass-through so SET LOCAL on a gateway-managed variable updates
	// the gateway state instead of the (uninvolved) backend, keeping subsequent
	// SHOW consistent with PostgreSQL semantics. The check also runs before the
	// Kind filter because VAR_SET_DEFAULT needs to be intercepted (treated as RESET).
	return nil, nil
}

// Non-gateway-managed SET LOCAL passes through to PostgreSQL unchanged —
// the backend is authoritative for those variables.

// SET var TO DEFAULT is equivalent to RESET var in PostgreSQL
// (PG's ExecSetVariableStmt falls through from VAR_SET_DEFAULT to VAR_RESET).
// Normalize before the switch so it shares the same tracking path.

// These are tracked locally

// VAR_SET_MULTI: SET TRANSACTION / SET SESSION CHARACTERISTICS — transaction-scoped,
//   must be executed directly on the backend, no session tracking needed.
// VAR_SET_CURRENT: SET var FROM CURRENT — reads current PG value, needs backend execution.

// isGatewayManagedVariable returns true for session variables that are managed
// entirely by the gateway and should NOT be forwarded to PostgreSQL.
// These variables control gateway-level behavior (e.g., timeouts) and sending
// them to PostgreSQL would be redundant or counterproductive for connection pooling.
func isGatewayManagedVariable(name string) bool { _ = "STUB: not implemented"; return false }

// planGatewayManagedVariable creates a GatewaySessionState primitive for a
// gateway-managed variable. All parsing and validation happens here at plan
// time so the primitive's execute path is a simple assignment.
func (p *Planner) planGatewayManagedVariable(
	sql string,
	stmt *ast.VariableSetStmt,
	value string,
) (engine.Primitive, error) {
	_ = "STUB: not implemented"
	return *new(engine.Primitive), nil
}

// RESET and SET ... TO DEFAULT revert to the flag default.
// SET LOCAL var TO DEFAULT (IsLocal && VAR_SET_DEFAULT) is distinct: it
// installs a transaction-scoped override equal to the default, masking
// (not destroying) the session value. The primitive branches on isLocal.
// RESET itself has no LOCAL form in the grammar, so stmt.IsLocal is
// always false for VAR_RESET.
// Note: VAR_RESET_ALL is not listed here because RESET ALL has stmt.Name=""
// which never passes isGatewayManagedVariable. RESET ALL is handled by
// ApplySessionState (after routing to PostgreSQL) which resets both
// PostgreSQL session settings and gateway-managed variables.
// isResetStmt is set only for VAR_RESET so the wire CommandTag is
// "RESET" for `RESET var` and "SET" for `SET [LOCAL] var TO DEFAULT`,
// matching PostgreSQL.

// extractVariableValue converts AST NodeList arguments to a string value.
// Handles: single values, multiple values, integers, strings, etc.
func extractVariableValue(args *ast.NodeList) string { _ = "STUB: not implemented"; return "" }

// Handle multiple args (e.g., search_path = 'schema1', 'schema2')

// A_Const wraps the actual value - unwrap it

// Direct String literal - SVal is already unquoted

// Direct Integer literal

// For complex types, use SqlString() as fallback

// Join multiple values with ", " (PostgreSQL format)

// extractConstValue extracts string value from A_Const node.
func extractConstValue(aConst *ast.A_Const) string { _ = "STUB: not implemented"; return "" }
