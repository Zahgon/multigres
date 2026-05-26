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

package executor

import (
	"context"
	"log/slog"

	"github.com/multigres/multigres/go/common/parser/ast"
	"github.com/multigres/multigres/go/common/pgprotocol/server"
	"github.com/multigres/multigres/go/common/preparedstatement"
	"github.com/multigres/multigres/go/common/sqltypes"
	"github.com/multigres/multigres/go/pb/query"
	"github.com/multigres/multigres/go/services/multigateway/engine"
	"github.com/multigres/multigres/go/services/multigateway/handler"
	"github.com/multigres/multigres/go/services/multigateway/plancache"
	"github.com/multigres/multigres/go/services/multigateway/planner"
)

const (
	// TODO(GuptaManan100): Remove this and use discovery to find the table group and use that.
	DefaultTableGroup = "default"
)

// Executor is the query execution engine for multigateway.
// It handles query planning, routing to appropriate multipooler instances,
// and result streaming back to clients.
//
// The Executor depends only on the IExecute interface, not on concrete
// implementations like ScatterConn. This makes it easy to test by passing
// mock implementations.
type Executor struct {
	planner   *planner.Planner
	exec      engine.IExecute
	logger    *slog.Logger
	planCache *plancache.PlanCache
}

// NewExecutor creates a new executor instance.
// The IExecute parameter provides the execution backend (typically ScatterConn).
// planCacheMemory controls the maximum memory in bytes for the plan cache (0 disables caching).
func NewExecutor(exec engine.IExecute, logger *slog.Logger, planCacheMemory int) *Executor {
	_ = "STUB: not implemented"
	return nil
}

// StreamExecute executes a query and streams results back via the callback function.
//
// For cacheable statements (SELECT, INSERT, UPDATE, DELETE), the executor
// normalizes the query (replacing literals with $1, $2, ... placeholders)
// and checks the plan cache. On a cache hit, the cached plan is reused with
// the current query's bind variables. On a miss, the query is planned using
// the normalized SQL/AST, and the resulting plan is cached for future reuse.
//
// The callback function is invoked for each chunk of results. For large result sets,
// the callback may be invoked multiple times with partial results.
func (e *Executor) StreamExecute(
	ctx context.Context,
	conn *server.Conn,
	state *handler.MultiGatewayConnectionState,
	queryStr string,
	astStmt ast.Stmt,
	callback func(ctx context.Context, res *sqltypes.Result) error,
) (*handler.ExecuteResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// resolvePlan obtains a query plan, using the plan cache when possible.
// Returns the plan, bind variables extracted during normalization (nil if none),
// whether the plan was a cache hit, the normalized SQL string (empty for
// non-cacheable statements), a stable fingerprint hash of that normalized SQL,
// and any planning error.
func (e *Executor) resolvePlan(
	ctx context.Context,
	queryStr string,
	astStmt ast.Stmt,
	conn *server.Conn,
) (*engine.Plan, []*ast.A_Const, bool, string, string, error) {
	_ = "STUB: not implemented"
	return nil, nil, false, "", "", nil
}

// Normalize: replace literals with $1, $2, ... placeholders.
// If the query has no literals, NormalizedSQL equals the original SQL
// and BindValues is empty — the plan is still cached by its SQL string.

// Cache hit

// Cache miss — plan with normalized SQL/AST and cache the result.

// isCacheable returns true if the statement type is eligible for plan caching.
// Only DML statements that go through planDefault() are cacheable.
func isCacheable(stmt ast.Stmt) bool { _ = "STUB: not implemented"; return false }

// Exclude SELECT INTO — temp-table variants use a different primitive
// (TempTableRoute), and non-temp variants are DDL-like (they create a
// table), so caching their plans is not useful.

// PortalStreamExecute executes a portal and streams results back via the callback function.
func (e *Executor) PortalStreamExecute(
	ctx context.Context,
	conn *server.Conn,
	state *handler.MultiGatewayConnectionState,
	portalInfo *preparedstatement.PortalInfo,
	maxRows int32,
	includeDescribe bool,
	callback func(ctx context.Context, res *sqltypes.Result) error,
) (*handler.ExecuteResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// For cacheable DML (SELECT, INSERT, UPDATE, DELETE), use the plan cache
// to resolve routing. The extended protocol query already has $1, $2, ...
// placeholders — the same form as our normalized cache key — so portal
// queries share cache entries with simple protocol queries.

// Hand off to the plan, which delegates to its root primitive's
// PortalStreamExecute. Each primitive owns its portal-mode behavior:
// Route reissues the portal to the multipooler, Sequence iterates
// children (so any silent ApplySessionState prefix runs before the
// trailing Route forwards), gateway-local primitives ignore
// portalInfo and run their StreamExecute logic.

// Non-cacheable — check if the gateway needs to handle locally (e.g.,
// SET/SHOW gateway-managed variables, LISTEN/NOTIFY, temp table DDL).

// Non-cacheable, non-local — send directly to multipooler with defaults.

// resolvePortalPlan looks up or creates a cached plan for a portal's query.
// The AST's SqlString() is used as the normalized SQL portion of the cache key,
// producing the same canonical form as the simple protocol path. This ensures
// cross-protocol cache sharing regardless of casing or whitespace differences
// in the original query text.
func (e *Executor) resolvePortalPlan(
	ctx context.Context,
	astStmt ast.Stmt,
	conn *server.Conn,
) (*engine.Plan, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// buildCacheKey constructs the plan cache key from the database name and
// normalized SQL. Including the database prevents cross-database plan reuse
// (different databases may have different schemas and routing).
//
// TODO(GuptaManan100): When shard-aware routing is introduced and the planner
// starts resolving table names for shard selection, search_path will need to
// be included in the cache key as well, since it affects table name resolution.
func buildCacheKey(database, normalizedSQL string) string { _ = "STUB: not implemented"; return "" }

// Describe returns metadata about a prepared statement or portal.
func (e *Executor) Describe(
	ctx context.Context,
	conn *server.Conn,
	state *handler.MultiGatewayConnectionState,
	portalInfo *preparedstatement.PortalInfo,
	preparedStatementInfo *preparedstatement.PreparedStatementInfo,
) (*query.StatementDescription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: We will need to plan the query to find whether it can
// be served by a single shard or not. For now, since we only
// support unsharded, we don't have to do much.
// We just send the query to the default table group.

// ReleaseAll releases all reserved connections, regardless of reservation reason.
// Delegates to ReleaseAllReservedConnections which calls ReleaseReservedConnection
// on the multipooler for each reserved connection. The multipooler handles
// rollback, COPY abort, and portal release internally.
// Used for connection cleanup when a client disconnects.
func (e *Executor) ReleaseAll(
	ctx context.Context,
	conn *server.Conn,
	state *handler.MultiGatewayConnectionState,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Close shuts down the executor, releasing resources such as the plan cache.
func (e *Executor) Close() { _ = "STUB: not implemented"; return }

// Ensure Executor implements handler.Executor interface.
var _ handler.Executor = (*Executor)(nil)
