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

// Package scatterconn handles coordinated query execution across multiple
// multipooler instances. It implements the IExecute interface from the engine
// package and is responsible for:
// - Selecting appropriate poolers for a given tablegroup
// - Executing queries via gRPC
// - Streaming results back
// - Handling failures and retries
package scatterconn

import (
	"context"
	"log/slog"
	"time"

	"go.opentelemetry.io/otel/trace"

	"github.com/multigres/multigres/go/common/pgprotocol/server"
	"github.com/multigres/multigres/go/common/preparedstatement"
	"github.com/multigres/multigres/go/common/sqltypes"
	multipoolerpb "github.com/multigres/multigres/go/pb/multipoolerservice"
	querypb "github.com/multigres/multigres/go/pb/query"
	"github.com/multigres/multigres/go/services/multigateway/engine"
	"github.com/multigres/multigres/go/services/multigateway/handler"
	"github.com/multigres/multigres/go/services/multigateway/poolergateway"
)

// ScatterConn coordinates query execution across multiple multipooler instances.
// It implements the engine.IExecute interface.
type ScatterConn struct {
	logger *slog.Logger

	// gateway is used for executing queries (typically a PoolerGateway)
	gateway poolergateway.Gateway

	metrics *ScatterMetrics
}

// NewScatterConn creates a new ScatterConn instance.
func NewScatterConn(gateway poolergateway.Gateway, logger *slog.Logger) *ScatterConn {
	_ = "STUB: not implemented"
	return nil
}

// userAuthFrom builds the outbound UserAuth payload from the session's captured
// SCRAM passthrough keys. Returns nil for sessions that did not authenticate via
// SCRAM (e.g. trust in tests), so this field stays absent on the wire instead of
// carrying empty slices.
//
// Keys are copied rather than referenced: Conn.Close zeroizes the accessor's
// backing slice in place, and gRPC may marshal lazily (stream init) or re-
// marshal on transient retry. A detached copy guarantees the proto carries
// live bytes for the full lifetime of the RPC.
func userAuthFrom(conn *server.Conn) *querypb.UserAuth { _ = "STUB: not implemented"; return nil }

// buildTarget constructs a routing target from the given tableGroup and shard.
// When the connection arrived on the replica-reads port (state.TargetReplica()),
// the target's PoolerType is set to REPLICA; otherwise PRIMARY.
func (sc *ScatterConn) buildTarget(tableGroup, shard string, state *handler.MultiGatewayConnectionState) *querypb.Target {
	_ = "STUB: not implemented"
	return nil
}

// applyReservedState replaces independent bookkeeping with the authoritative reservation
// state from the multipooler. If the reserved connection ID is zero, the connection was
// destroyed or released — clear the shard state. Otherwise, update the reservation reasons.
//
// When a reserved connection is destroyed while in a transaction, the transaction is marked
// as failed (TxnStatusFailed) so subsequent queries are rejected until ROLLBACK. This is
// defense-in-depth — handler.go also sets TxnStatusFailed on query errors, but setting it
// here ensures coverage for all code paths (COPY, portal, etc.).
func (sc *ScatterConn) applyReservedState(
	conn *server.Conn,
	state *handler.MultiGatewayConnectionState,
	target *querypb.Target,
	rs *querypb.ReservedState,
) {
	_ = "STUB: not implemented"
	return
}

// StreamExecute executes a query on the specified tablegroup and streams results.
// This is the implementation of engine.IExecute.StreamExecute().
//
// Implements 3-case reservation logic for transactions:
//   - Case 1: Has reserved connection → use it
//   - Case 2: In transaction, no reserved conn → call StreamExecute with reservation options
//   - Case 3: Not in transaction → use regular pooled connection
//
// If preparedStatement is non-nil, it is attached to the ExecuteOptions so
// the multipooler can ensurePrepared() on the backend connection before
// running the query. Used for wrapped EXECUTE forms (EXPLAIN EXECUTE,
// CREATE TABLE ... AS EXECUTE) that reference a gateway-managed prepared
// statement by its canonical name.
func (sc *ScatterConn) StreamExecute(
	ctx context.Context,
	conn *server.Conn,
	tableGroup string,
	shard string,
	sql string,
	preparedStatement *querypb.PreparedStatement,
	state *handler.MultiGatewayConnectionState,
	callback func(context.Context, *sqltypes.Result) error,
) (retErr error) {
	_ = "STUB: not implemented"
	return nil
}

// Case 1: Already have reserved connection - use it

// Build reservation options for any reasons we need to add to the
// existing reserved connection.

// For reserved connections with temp tables and a deferred BEGIN:
// set reservation options so the multipooler executes BEGIN on the
// reserved connection before the query.

// If this query creates a temp table, add the reason so the
// multipooler tracks it on the reserved connection.

// If this query declares a `WITH HOLD` cursor, pin the cursor name on
// the reserved backend so the cursor survives COMMIT. Take the
// pending list through the mutex-protected accessor so concurrent
// access (e.g. a future cancellation goroutine touching state) stays
// race-free.

// If this query closes a `WITH HOLD` cursor, unpin it after the CLOSE
// runs on the backend. The multipooler will drop the reservation
// (returning ReservedConnectionId=0) when the last reason clears.

// Case 2: Need a new reserved connection — for transaction, temp table,
// portal pin (DECLARE WITH HOLD), or any combination. Take the pin
// list once up front so we don't double-lock the state mutex via a
// separate Has check.

// If the session already has a temp table reservation on another shard,
// include the temp table reason so the connection survives COMMIT.

// Pass the original BEGIN query (e.g., "BEGIN ISOLATION LEVEL SERIALIZABLE")
// so the multipooler preserves transaction options instead of using plain "BEGIN".

// Case 3: Not in transaction, no temp table — use regular pooled connection

// If it's a PostgreSQL error, don't wrap it - pass through unchanged

// PortalStreamExecute executes a portal (bound prepared statement) and streams results.
// This is the implementation of engine.IExecute.PortalStreamExecute().
func (sc *ScatterConn) PortalStreamExecute(
	ctx context.Context,
	tableGroup string,
	shard string,
	conn *server.Conn,
	state *handler.MultiGatewayConnectionState,
	portalInfo *preparedstatement.PortalInfo,
	maxRows int32,
	includeDescribe bool,
	callback func(context.Context, *sqltypes.Result) error,
) (retErr error) {
	_ = "STUB: not implemented"
	return nil
}

// Create target for routing

// When the protocol layer folded a Describe('P') into this Execute, ask
// the multipooler to fuse Bind+Describe(P)+Execute+Sync into one
// backend round trip. The portal RowDescription rides back through
// the streaming callback's Fields on the first chunk; pgwire-server's
// handleExecute writes it to the wire before any DataRow.

// If we have a reserved connection, we have to ensure
// we are routing the query to the pooler where we got the reserved
// connection from. If a reparent happened, then we will get an error
// back.

// Case 2: Need a new reserved connection — for transaction, temp table, or both.
// We use StreamExecute with reservation options and a no-op "SELECT 1" query
// rather than adding a dedicated ReservePortalStreamExecute RPC.

// Execute portal via QueryService (PoolerGateway) and stream results

// Use the query from the prepared statement

// If it's a PostgreSQL error, don't wrap it - pass through unchanged

// Use authoritative state from multipooler. The multipooler already OR'd in
// the portal reason (if suspended) or removed it (if completed). If no reasons
// remain (ReservedConnectionId == 0) the connection was released.

// Describe returns metadata about a prepared statement or portal.
// This is the implementation of engine.IExecute.Describe().
func (sc *ScatterConn) Describe(
	ctx context.Context,
	tableGroup string,
	shard string,
	conn *server.Conn,
	state *handler.MultiGatewayConnectionState,
	portalInfo *preparedstatement.PortalInfo,
	preparedStatementInfo *preparedstatement.PreparedStatementInfo,
) (*querypb.StatementDescription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create target for routing

// If we have a reserved connection, use it

// Call Describe on the query service

// If it's a PostgreSQL error, don't wrap it - pass through unchanged

// ConcludeTransaction concludes a transaction on reserved connections that have
// the transaction reason set. Shards reserved for other reasons only (e.g., temp
// tables, portals) are left untouched. Based on the returned remainingReasons from
// the multipooler, shard state entries are either cleared (connection fully released)
// or updated with the new reason bitmask.
func (sc *ScatterConn) ConcludeTransaction(
	ctx context.Context,
	conn *server.Conn,
	state *handler.MultiGatewayConnectionState,
	conclusion multipoolerpb.TransactionConclusion,
	releasePortalNames []string,
	releaseAllPortals bool,
	callback func(context.Context, *sqltypes.Result) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Collect shard state updates to apply after iteration.
// We cannot call ClearReservedConnection during iteration because it
// uses swap-and-truncate which mutates the underlying slice.

// true = remove entry, false = set state
// only used when clear == false

// Count shards with a transaction reason — multi-shard transactions are not
// yet supported (distributed transactions). Log a warning as a sentinel so
// unexpected multi-shard cases are visible before DT is implemented.

// Iterate over all shard states with reserved connections.
// Only conclude on shards that have the transaction reason.
// Continue on errors so all shards get concluded.

// Connection lost — mark for clearing, continue to other shards

// ROLLBACK on a destroyed connection is graceful recovery — don't propagate error

// Connection fully released by multipooler

// Connection still reserved for other reasons — update our local tracking

// Keep the last successful result for the callback

// Apply collected updates outside the iteration loop.

// Return only the first error. PostgreSQL clients expect a single ErrorResponse
// with a SQLSTATE — a joined multi-line error confuses ORMs and connection poolers.

// Send the result to the client (COMMIT/ROLLBACK command tag)

// DiscardTempTables sends DISCARD TEMP on reserved connections that have the
// temp table reason set. Based on the returned state from the multipooler,
// shard state entries are either cleared (connection fully released) or updated.
func (sc *ScatterConn) DiscardTempTables(
	ctx context.Context,
	conn *server.Conn,
	state *handler.MultiGatewayConnectionState,
	callback func(context.Context, *sqltypes.Result) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Collect shard state updates to apply after iteration.

// Iterate over all shard states with reserved connections.
// Only discard on shards that have the temp table reason.

// Connection fully released by multipooler

// Connection still reserved for other reasons — update our local tracking

// Keep the last successful result for the callback

// Apply collected updates outside the iteration loop.

// Send the result to the client

// --- COPY FROM STDIN methods ---

// CopyInitiate initiates a COPY FROM STDIN operation using bidirectional streaming.
// Stores reserved connection info in state.ShardStates for the given tableGroup/shard.
// Returns: format, columnFormats, error
func (sc *ScatterConn) CopyInitiate(
	ctx context.Context,
	conn *server.Conn,
	tableGroup string,
	shard string,
	queryStr string,
	state *handler.MultiGatewayConnectionState,
	callback func(ctx context.Context, result *sqltypes.Result) error,
) (int16, []int16, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

// Create target for routing - COPY always goes to PRIMARY

// Create execute options

// If there's already a reserved connection for this target (e.g., in a transaction),
// pass its ID so CopyReady reuses it instead of creating a new one.
// If we're in a transaction but no reserved connection exists yet (deferred BEGIN),
// pass ReservationOptions so CopyReady creates a connection with the pending BEGIN.

// Deferred BEGIN: pass transaction reservation options so the executor
// executes BEGIN on the new connection before initiating COPY.

// Call CopyReady on gateway to initiate the COPY and get format info

// When init fails, the multipooler may still have a live reserved
// connection (e.g., it was already held for a transaction or temp
// tables and only the COPY query itself was rejected). Apply whatever
// state came back: a non-nil ReservedState keeps the gateway pointed
// at the surviving conn, while a nil/zero state clears the tracking
// so we don't end up sending future statements to a connection that
// no longer exists.

// Use authoritative state from multipooler (reasons already include copy + transaction if applicable)

// CopySendData sends a chunk of COPY data via bidirectional stream.
// Looks up reserved connection from state.ShardStates based on tableGroup/shard.
func (sc *ScatterConn) CopySendData(
	ctx context.Context,
	conn *server.Conn,
	tableGroup string,
	shard string,
	state *handler.MultiGatewayConnectionState,
	data []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Create target for routing

// Get the reserved connection ID from shard state

// Build options with reserved connection ID

// Send data via gateway

// CopyFinalize sends the final chunk and CopyDone via bidirectional stream.
// Looks up reserved connection from state.ShardStates based on tableGroup/shard.
func (sc *ScatterConn) CopyFinalize(
	ctx context.Context,
	conn *server.Conn,
	tableGroup string,
	shard string,
	state *handler.MultiGatewayConnectionState,
	finalData []byte,
	callback func(ctx context.Context, result *sqltypes.Result) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Create target for routing

// Get the reserved connection ID from shard state

// Build options with reserved connection ID

// Finalize the COPY operation via gateway

// CopyFinalize returns a non-nil ReservedState when a PG-level error
// (e.g., constraint violation) left the underlying reserved connection
// alive because it is still holding another reason such as a
// transaction. A nil state means the connection was released. Either
// way, applyReservedState does the right thing: keep tracking if the
// state has a non-zero conn ID, clear it and mark the transaction
// failed if not.

// Call callback with result

// Update shard state with authoritative state from multipooler

// CopyAbort aborts the COPY operation via bidirectional stream.
// Looks up reserved connection from state.ShardStates based on tableGroup/shard.
func (sc *ScatterConn) CopyAbort(
	ctx context.Context,
	conn *server.Conn,
	tableGroup string,
	shard string,
	state *handler.MultiGatewayConnectionState,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Create target for routing

// Get the reserved connection ID from shard state

// Already cleaned up

// Build options with reserved connection ID

// Abort the COPY operation via gateway

// Update shard state with authoritative state from multipooler

// ReleaseAllReservedConnections forcefully releases all reserved connections.
// Iterates all shard states and calls ReleaseReservedConnection on the multipooler
// for each one. Errors are logged and collected but do not stop the iteration
// (best-effort). After the loop, all local shard state is cleared.
func (sc *ScatterConn) ReleaseAllReservedConnections(
	ctx context.Context,
	conn *server.Conn,
	state *handler.MultiGatewayConnectionState,
) error {
	_ = "STUB: not implemented"
	return nil
}

// endAction records shard-level metrics and span status for both success and
// error outcomes. Designed to be called via defer with a pointer to the named
// error return.
func (sc *ScatterConn) endAction(ctx context.Context, span trace.Span, start time.Time, dbNamespace, tableGroup, shard string, err *error) {
	_ = "STUB: not implemented"
	return
}

// TODO: Consider filtering out client-caused errors (e.g. unique constraint violations)
// from the counter to avoid inflating error rates. We currently count all errors and
// rely on the error.type label for dashboard filtering. Revisit if counters get noisy.

// Ensure ScatterConn implements engine.IExecute interface.
// This will be checked at compile time.
var _ engine.IExecute = (*ScatterConn)(nil)
