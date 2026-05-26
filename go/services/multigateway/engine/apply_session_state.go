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

package engine

import (
	"context"

	"github.com/multigres/multigres/go/common/parser/ast"
	"github.com/multigres/multigres/go/common/pgprotocol/server"
	"github.com/multigres/multigres/go/common/preparedstatement"
	"github.com/multigres/multigres/go/common/sqltypes"
	"github.com/multigres/multigres/go/services/multigateway/handler"
)

// ApplySessionState handles SET/RESET commands by updating local session state only.
//
// Neither SET nor RESET is routed to PostgreSQL. Both are applied locally and
// return a synthetic CommandComplete response. The pool propagates these settings
// to the backend connection on the next query via ApplySettings.
//
// Behaviour deviation from PostgreSQL:
// SET commands are NOT validated against PostgreSQL. This means a client can SET
// an invalid variable name or value without receiving an immediate error. The error
// will surface on the next query when the pool tries to apply the setting to a
// backend connection. The client must RESET the bad variable to recover. This
// trade-off was chosen intentionally to keep the SET/RESET path simple. It may
// be revisited in the future if stricter validation is needed.
//
// For RESET/RESET ALL:
// The variable is removed from SessionSettings. On the next query, the merged
// settings (SessionSettings overlaid on StartupParams) will fall back to the
// startup parameter value, and the pool will apply the correct SET commands.
type ApplySessionState struct {
	// VariableStmt is the SET/RESET statement from the AST.
	VariableStmt *ast.VariableSetStmt

	// Query is the original SQL string.
	Query string

	// SilentTracking, when true, updates SessionSettings but does NOT invoke
	// the callback. Used inside a Sequence where a sibling primitive (like
	// Route) owns the client-facing result — if both called back, the client
	// would see a stray CommandComplete before the real row data. This is
	// the shape a `SELECT set_config(...)` plan takes: silent tracking step
	// first, then a Route that sends the query to PG and streams the result.
	SilentTracking bool
}

// NewApplySessionState creates a new ApplySessionState primitive.
func NewApplySessionState(sql string, stmt *ast.VariableSetStmt) *ApplySessionState {
	_ = "STUB: not implemented"
	return nil
}

// NewApplySessionStateSilent creates an ApplySessionState that updates the
// tracker without emitting anything to the client. Intended for use inside a
// Sequence where a Route primitive owns the client-facing response — see
// planner.planSelectStmt for the `SELECT set_config(...), * FROM t` case.
func NewApplySessionStateSilent(sql string, stmt *ast.VariableSetStmt) *ApplySessionState {
	_ = "STUB: not implemented"
	return nil
}

// PortalStreamExecute handles SET/RESET on the extended-protocol path. The
// primitive's effect is local to gateway state — it neither reads bind
// values nor talks to a backend — so we delegate to StreamExecute and
// ignore portalInfo entirely. This is the right shape for a silent
// ApplySessionState sitting inside a Sequence whose trailing Route
// reissues the actual portal.
func (s *ApplySessionState) PortalStreamExecute(
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

// StreamExecute handles the SET/RESET command.
func (s *ApplySessionState) StreamExecute(
	ctx context.Context,
	_ IExecute,
	_ *server.Conn,
	state *handler.MultiGatewayConnectionState,
	_ []*ast.A_Const,
	callback func(context.Context, *sqltypes.Result) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

// executeSet handles SET commands: update local state and return a synthetic
// response. The value is NOT validated against PostgreSQL — see the
// ApplySessionState doc comment.
//
// Two modes:
//   - SilentTracking: update state, no callback (a sibling primitive in a
//     Sequence will respond — used for SELECT set_config(...) plans).
//   - default: update state and emit CommandComplete "SET" (real SET stmt).
func (s *ApplySessionState) executeSet(
	ctx context.Context,
	state *handler.MultiGatewayConnectionState,
	callback func(context.Context, *sqltypes.Result) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

// executeReset handles RESET/RESET ALL: update state, return synthetic response.
//
// Two modes (mirrors executeSet):
//   - SilentTracking: update state, no callback. No current planner path
//     produces a silent RESET, but gating defensively prevents a future
//     caller from emitting a stray CommandComplete("RESET") into the
//     protocol stream ahead of a sibling primitive's real response.
//   - default: update state and emit CommandComplete "RESET".
func (s *ApplySessionState) executeReset(
	ctx context.Context,
	state *handler.MultiGatewayConnectionState,
	callback func(context.Context, *sqltypes.Result) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

// RESET variable

// Also reset gateway-managed variables that live outside SessionSettings.

// Return synthetic CommandComplete

// GetTableGroup returns empty string — SET/RESET are local-only and don't target a tablegroup.
func (s *ApplySessionState) GetTableGroup() string {
	_ = "STUB: not implemented"

	// GetQuery returns the original SQL string.
	return ""
}

func (s *ApplySessionState) GetQuery() string {
	_ = "STUB: not implemented"

	// String returns a string representation for debugging.
	return ""
}

func (s *ApplySessionState) String() string { _ = "STUB: not implemented"; return "" }

// extractVariableValue converts AST NodeList arguments to a string value.
func extractVariableValue(args *ast.NodeList) string { _ = "STUB: not implemented"; return "" }

// extractConstValue extracts string value from A_Const node.
func extractConstValue(aConst *ast.A_Const) string { _ = "STUB: not implemented"; return "" }

// Ensure ApplySessionState implements Primitive interface.
var _ Primitive = (*ApplySessionState)(nil)
