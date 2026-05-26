// Copyright 2025 Supabase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package executor implements query execution for multipooler.
// It provides the QueryService interface implementation that executes queries
// against PostgreSQL using per-user connection pools.
package executor

import (
	"context"
	"log/slog"

	"github.com/multigres/multigres/go/common/preparedstatement"
	"github.com/multigres/multigres/go/common/queryservice"
	"github.com/multigres/multigres/go/common/sqltypes"
	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	multipoolerpb "github.com/multigres/multigres/go/pb/multipoolerservice"
	"github.com/multigres/multigres/go/pb/query"
	"github.com/multigres/multigres/go/services/multipooler/connpoolmanager"
	"github.com/multigres/multigres/go/services/multipooler/pools/regular"
	"github.com/multigres/multigres/go/services/multipooler/pools/reserved"
)

// Executor implements the QueryService interface for executing queries against PostgreSQL.
// It uses the connpoolmanager for per-user connection pool management and consolidates
// prepared statements across connections to avoid redundant parsing.
type Executor struct {
	logger             *slog.Logger
	poolManager        connpoolmanager.PoolManager
	poolerConsolidator *preparedstatement.PoolerConsolidator
	poolerID           *clustermetadatapb.ID

	// vpidStampEnabled toggles the multigres_vpid:<id> stamping on PostgreSQL
	// backends and the matching application_name filter in
	// sessionSettingsForPool. Both must move together: stamping without
	// filtering lets ApplySettings wipe the stamp via RESET application_name;
	// filtering without stamping silently swallows client-set application_name.
	vpidStampEnabled bool
}

// sessionSettingsForPool returns a copy of settings safe to apply to a pooled
// (regular or reserved) PostgreSQL connection.
//
// When vpid stamping is enabled, it excludes application_name. The pool's
// connstate cache must never track a client-supplied application_name: when
// SetApplicationName is later called out-of-band on the same connection (to
// stamp `multigres_vpid:<id>` for lock-detection mapping), connstate is
// unaware of the new value, and a subsequent ApplySettings diff between
// connstate.current (still holding the client's app_name) and the desired
// settings on the next query emits a RESET application_name that wipes the
// stamp before the query runs. Filtering here prevents that ABA on every code
// path that pushes SessionSettings into the pool.
//
// When stamping is disabled, settings pass through unchanged so client-set
// application_name reaches the backend normally.
func (e *Executor) sessionSettingsForPool(settings map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// stampVpidOnReserved tags a reserved connection's PostgreSQL backend with
// `multigres_vpid:<client_connection_id>` in application_name so
// lock-detection functions (e.g. an override of
// pg_isolation_test_session_is_blocked) can map a multigateway virtual PID
// back to the real backend PID via pg_stat_activity. Best-effort: a SET
// failure does not block the actual query — only lock detection through the
// proxy depends on the tag. No-op when vpid stamping is disabled.
func (e *Executor) stampVpidOnReserved(ctx context.Context, conn *reserved.Conn, options *query.ExecuteOptions) {
	_ = "STUB: not implemented"
	return
}

// stampVpidOnRegular tags a pooled regular connection with the same
// `multigres_vpid:<id>` marker. Pooled connections are shared across clients,
// so the next checkout will overwrite this stamp; for the duration of the
// current query the backend is correctly attributed to its client vpid.
// No-op when vpid stamping is disabled.
func (e *Executor) stampVpidOnRegular(ctx context.Context, conn *regular.Conn, options *query.ExecuteOptions) {
	_ = "STUB: not implemented"
	return
}

// NewExecutor creates a new Executor instance.
// vpidStampEnabled controls whether multigres_vpid:<id> is stamped on
// PostgreSQL backends and whether application_name is filtered from pool
// SessionSettings.
func NewExecutor(logger *slog.Logger, poolManager connpoolmanager.PoolManager, poolerID *clustermetadatapb.ID, vpidStampEnabled bool) *Executor {
	_ = "STUB: not implemented"
	return nil
}

// buildReservedState constructs a ReservedState from the current state of a reserved connection.
func (e *Executor) buildReservedState(reservedConn *reserved.Conn) *query.ReservedState {
	_ = "STUB: not implemented"
	return nil
}

// buildReservedStateFromAPI constructs a ReservedState from any value satisfying
// reservedConnAPI. Used by streamExecuteOnReservedConn so the helper can be exercised
// in unit tests with a mock conn instead of a real *reserved.Conn.
func (e *Executor) buildReservedStateFromAPI(rc reservedConnAPI) *query.ReservedState {
	_ = "STUB: not implemented"
	return nil
}

// ExecuteQuery implements queryservice.QueryService.
// It executes a query using a pooled connection for the specified user.
// If ReservedConnectionId is set in options, uses that reserved connection instead.
// Returns ReservedState with the authoritative reservation state from the multipooler.
func (e *Executor) ExecuteQuery(ctx context.Context, target *query.Target, sql string, options *query.ExecuteOptions) (*sqltypes.Result, *query.ReservedState, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Check if we should use an existing reserved connection

// Connection destroyed — return zero state so gateway clears its tracking

// Apply settings if they changed (e.g., SET inside a transaction).
// Reserved connections bypass the pool's normal ApplySettings mechanism,
// so we must explicitly apply settings changes here.

// Stamp multigres_vpid:<id> AFTER ApplySettingsToConn. When the
// filtered desired settings collapse to nil (e.g. the only client
// setting was application_name), ApplySettings issues RESET ALL on
// the reserved conn — which would wipe a stamp set earlier in this
// function. Restamping after the reset ensures the tag is in place
// for the actual query.

// Query failed but connection still exists — return current state

// Get session settings from options

// Get a connection from the pool for this user

// Stamp multigres_vpid on this pooled regular conn too — the next
// client to draw from the pool will overwrite it, but for the duration
// of the current query the backend is correctly tagged.

// Execute the query - the regular.Conn.QueryWithRetry returns []*sqltypes.Result
// with proper field info, rows, and command tags already populated.
// Uses retry variant since this is a stateless pool query.

// Return first result (simple query returns single result)

// StreamExecute executes a query and streams results back via callback.
// This implements the queryservice.QueryService interface.
//
// Handles three cases based on the request:
//   - options.ReservedConnectionId > 0: use existing reserved connection
//   - reservationOptions has non-zero reasons && ReservedConnectionId == 0: create new reserved connection
//   - Neither: use regular pooled connection
//
// When reservationOptions is set on an existing reserved connection, the reasons are
// OR'd into the reservation (e.g., adding ReasonTransaction to a temp-table-reserved conn).
//
// If options.PreparedStatement is set, the statement is parsed on the chosen
// backend connection via ensurePreparedWithName() before `sql` runs. This is
// used for wrapped EXECUTE forms (EXPLAIN EXECUTE, CREATE TABLE ... AS EXECUTE)
// that reference a gateway-managed prepared statement by its canonical name.
//
// Returns ReservedState with the authoritative reservation state from the multipooler.
func (e *Executor) StreamExecute(
	ctx context.Context,
	target *query.Target,
	sql string,
	options *query.ExecuteOptions,
	reservationOptions *query.ReservationOptions,
	callback func(context.Context, *sqltypes.Result) error,
) (*query.ReservedState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Case 1: Use an existing reserved connection

// Connection destroyed — return zero state so gateway clears its tracking

// Apply settings if they changed (e.g., SET inside a transaction).
// Reserved connections bypass the pool's normal ApplySettings mechanism,
// so we must explicitly apply settings changes here. This step depends
// on the underlying *regular.Conn so it stays out of the helper.

// Stamp multigres_vpid:<id> AFTER ApplySettingsToConn so a RESET
// ALL emitted by the empty-desired-settings path doesn't wipe it
// (see the matching ordering in ExecuteQuery).

// If the query references a gateway-managed prepared statement
// (wrapped EXECUTE forms), ensure it is parsed on this backend
// connection before running the query. We do this before delegating
// to streamExecuteOnReservedConn because that helper operates over
// the reservedConnAPI interface which does not expose the underlying
// *regular.Conn needed by ensurePreparedWithName.

// Case 2: Create a new reserved connection

// Case 3: Use regular pooled connection

// Get a connection from the pool for this user

// Stamp multigres_vpid on this pooled regular conn for lock-detection.

// When a PreparedStatement is provided we cannot use the retry-on-connection-error
// variant of QueryStreaming: reconnect wipes per-connection prepared-statement state
// (regular_conn.go Reconnect: "Prepared statements don't survive reconnection"),
// so after a silent reconnect the subsequent query would fail with "prepared
// statement does not exist". Skip the retry for this rare path; the caller can
// reissue the query at the application level on transient failures.

// Use streaming query execution with retry since this is a stateless pool query.

// reserveAndStreamExecute creates a new reserved connection and executes a query.
// Based on reservationOptions.Reasons, it may execute setup commands (e.g., BEGIN for transactions).
func (e *Executor) reserveAndStreamExecute(
	ctx context.Context,
	sql string,
	options *query.ExecuteOptions,
	reservationOptions *query.ReservationOptions,
	callback func(context.Context, *sqltypes.Result) error,
) (*query.ReservedState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get the reasons bitmask and determine if we need to execute BEGIN

// If the query references a gateway-managed prepared statement (wrapped
// EXECUTE forms like CREATE TEMP TABLE ... AS EXECUTE), parse it during
// reserved-connection acquisition. Doing the Parse via the validate
// callback lets the reserved pool transparently swap a stale (silently
// closed) socket for a fresh one before we register the connection — the
// failure mode that flaked TestWrappedPreparedStatementExecution.
//
// Parse is a session-level operation in PostgreSQL, so running it before
// BEGIN is safe; the prepared statement persists into the transaction.

// Create a reserved connection

// Stamp multigres_vpid on the freshly reserved backend so subsequent
// lock-detection probes can map vpid → real pid. Done before BEGIN so
// the value is in place for the entire transaction lifecycle.

// Apply all reservation reasons to the reserved connection.
// BeginWithQuery below adds ReasonTransaction internally, but non-transaction
// reasons (e.g., temp_table) must be added explicitly so that buildReservedState
// returns the correct bitmask and DiscardTempTables can find the shard.

// Register pin-portal entries from the gateway. Each name is a cursor
// declared with WITH HOLD; ReserveForPortal adds the name to the
// per-conn portal set and ORs ReasonPortal into the reservation
// bitmask (idempotent with the AddReservationReason call above).
// This path always releases the connection on QueryStreaming error
// (Release(ReleaseError) below), so a failed DECLARE never leaks a
// pin from the new-reservation branch — even without the explicit
// rollback dance that streamExecuteOnReservedConn performs for the
// existing-reservation case.

// If this is a transaction reservation, execute BEGIN first.
// The BEGIN result is not sent to the callback — it's an internal setup detail.
// The caller (multigateway) handles sending synthetic BEGIN results to the client.
// Use the original BEGIN query if provided to preserve isolation level and access mode.

// Execute the actual query and stream results to the callback as they arrive,
// matching the non-reserved StreamExecute path. This avoids buffering the entire
// result set in memory for large queries inside transactions.

// streamExecuteOnReservedConn executes a query on an existing reserved
// connection. It optionally promotes the reservation by adding new reasons
// (e.g., starting a transaction on a temp-table-reserved connection) before
// running the query.
//
// Defined over reservedConnAPI rather than *reserved.Conn so that unit tests
// can substitute a mock and exercise this path without a live PG connection.
func (e *Executor) streamExecuteOnReservedConn(
	ctx context.Context,
	rc reservedConnAPI,
	sql string,
	reservationOptions *query.ReservationOptions,
	callback func(context.Context, *sqltypes.Result) error,
) (*query.ReservedState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the caller is adding reservation reasons (e.g., promoting a temp-table
// reservation to also hold a transaction), apply them now.

// If the new reasons include transaction and the connection is not
// already in a transaction, execute BEGIN before the query.

// Add all requested non-transaction reasons to the reservation
// (BeginWithQuery already added ReasonTransaction internally).

// Register pin-portal entries for DECLARE … WITH HOLD before running
// the DECLARE itself so the bitmask is consistent during the round
// trip. The gateway only records the cursor in OpenHoldCursors on
// DECLARE success, so we mirror that here: if PG rejects the
// DECLARE (outside-of-transaction, syntax error, table missing,
// duplicate name, etc.) we roll back every pin we just added so
// the multipooler-side bitmask matches what the gateway thinks is
// open. Without this, a failed DECLARE outside an explicit
// transaction block would leak ReasonPortal on the reserved
// connection until session disconnect.

// Roll back every pin we registered for this DECLARE — PG
// rejected the statement, so the gateway will never call
// AddOpenHoldCursor and any matching CLOSE will not arrive.
// ReleasePortal returns true iff the call drained the *last*
// reservation reason on the connection (the bool propagates
// IsEmpty(), not "ReasonPortal cleared"), so a single true
// is sufficient to know the conn should be released.

// Apply portal releases after the query succeeds. CLOSE forwards the
// statement to PG first; only on success do we unpin so the
// gateway's HOLD-cursor bookkeeping matches the server side.
// ReleasePortal returns true iff this call drained the last
// reservation reason on the connection — when that happens, return
// the backend to the pool and surface a zero ReservedState so the
// gateway clears its shard tracking.

// Close closes the executor and releases resources.
// Note: The poolManager is managed by the caller (QueryPoolerServer), not closed here.
func (e *Executor) Close() error {
	_ = "STUB: not implemented"

	// PortalStreamExecute executes a portal (bound prepared statement) and streams results back via callback.
	// If MaxRows > 0, a reserved connection is used since the portal may be suspended and need resumption.
	// Otherwise, a regular connection is used for better pool efficiency.
	//
	// portalOptions carries portal-only knobs (e.g. include_describe). Nil leaves
	// every field at the proto default.
	return nil
}

func (e *Executor) PortalStreamExecute(
	ctx context.Context,
	target *query.Target,
	preparedStatement *query.PreparedStatement,
	portal *query.Portal,
	options *query.ExecuteOptions,
	portalOptions *multipoolerpb.PortalExecuteOptions,
	callback func(context.Context, *sqltypes.Result) error,
) (*query.ReservedState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Convert formats from int32 to int16

// Use reserved connection if:
// 1. ReservedConnectionId is already set (e.g., from transaction or previous portal)
// 2. MaxRows > 0 (portal may be suspended and need resumption)

// Use regular connection for non-suspended execution with no existing reservation

// portalExecuteWithReserved executes a portal using a reserved connection.
func (e *Executor) portalExecuteWithReserved(
	ctx context.Context,
	preparedStatement *query.PreparedStatement,
	portal *query.Portal,
	options *query.ExecuteOptions,
	settings map[string]string,
	user string,
	maxRows int32,
	includeDescribe bool,
	paramFormats, resultFormats []int16,
	callback func(context.Context, *sqltypes.Result) error,
) (*query.ReservedState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if we should use an existing reserved connection

// Create a new reserved connection. Wire ensurePrepared as the
// validate hook so that the Parse — the first user-issued write
// on this freshly acquired socket — triggers a transparent
// retry on a fresh socket if the pooled conn has been silently
// closed by PostgreSQL. The post-acquire ensurePrepared call
// below is a no-op for the new-conn path because connState
// dedupes by canonical name.

// Stamp multigres_vpid:<id> on the (possibly fresh, possibly resumed)
// reserved conn. Done unconditionally — both branches above can return
// a backend without the tag (a freshly created reservation, or an
// existing one whose tag was wiped by a prior ApplySettings RESET ALL).

// Ensure the statement is prepared on this connection (with consolidation).
// For the new-conn branch this is a no-op because the validate hook above
// already parsed it; for the existing-conn branch this is the only call.

// Bind and execute using the portal's own name and the canonical statement name.
// When the protocol layer folded Describe('P') into Execute, fuse the
// backend round trip too so the portal description rides on the Execute
// response.

// If portal is suspended (not completed), keep the reserved connection for continuation

// Portal completed, release this portal's reservation.
// ReleasePortal returns true only when all reservation reasons are gone.

// portalExecuteWithRegular executes a portal using a regular pooled connection.
// clientKey and serverKey are the SCRAM passthrough keys forwarded by the
// caller's session; see connpoolmanager.GetRegularConn.
func (e *Executor) portalExecuteWithRegular(
	ctx context.Context,
	preparedStatement *query.PreparedStatement,
	portal *query.Portal,
	settings map[string]string,
	user string,
	includeDescribe bool,
	clientKey, serverKey []byte,
	paramFormats, resultFormats []int16,
	options *query.ExecuteOptions,
	callback func(context.Context, *sqltypes.Result) error,
) (*query.ReservedState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Stamp multigres_vpid on this pooled regular conn for lock-detection.

// Ensure the statement is prepared on this connection (with consolidation)

// Bind and execute with maxRows=0 (fetch all) using the portal's own name and canonical statement name.
// When the protocol layer folded Describe('P') into Execute, fuse the
// backend round trip too so the portal description rides on the Execute
// response.

// No reserved connection for regular execution

// Describe returns metadata about a prepared statement or portal.
func (e *Executor) Describe(
	ctx context.Context,
	target *query.Target,
	preparedStatement *query.PreparedStatement,
	portal *query.Portal,
	options *query.ExecuteOptions,
) (*query.StatementDescription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Acquire the connection: reserved (transactional) or regular (pooled).

// Ensure the statement is prepared on this connection

// Describe prepared using canonical name

// ensurePreparedWithName ensures that a prepared statement named `name` with
// the given body exists on the backend connection, Parsing it if necessary.
// Unlike ensurePrepared (which derives its own canonical name via the pooler
// consolidator), this variant uses the caller-supplied name directly. It is
// used by the wrapped-EXECUTE StreamExecute path where the gateway has
// already rewritten the SQL to reference a specific name, and the backend
// must have a prepared statement under exactly that name.
//
// The gateway's prepared-statement consolidator assigns globally unique
// monotonic names ("stmt0", "stmt1", ...) that are deduplicated by
// (query, paramTypes), so using them as backend-session prepared statement
// names is safe across multiple gateway client sessions sharing a pool
// connection.
func (e *Executor) ensurePreparedWithName(ctx context.Context, conn *regular.Conn, stmt *query.PreparedStatement) error {
	_ = "STUB: not implemented"
	return nil
}

// If the name is already taken with a different query (e.g. from a
// different gateway instance that reused the same canonical name space),
// close the stale statement first — PostgreSQL rejects Parse for an
// already-existing prepared statement name.

// ensurePrepared ensures the prepared statement is available on the connection.
// It uses the PoolerConsolidator to get a canonical name by (query, paramTypes),
// then checks the connection state to avoid redundant parsing.
// Returns the canonical statement name to use.
func (e *Executor) ensurePrepared(ctx context.Context, conn *regular.Conn, stmt *query.PreparedStatement) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Check if this connection already has the statement prepared

// Statement already prepared on this connection, reuse it

// Parse the statement on this connection

// Store in connection state for future reuse

// CopyReady initiates a COPY FROM STDIN operation and returns format information.
// Uses an existing reserved connection if ReservedConnectionId is set in options,
// otherwise creates a new reserved connection (COPY requires connection affinity).
func (e *Executor) CopyReady(
	ctx context.Context,
	target *query.Target,
	copyQuery string,
	options *query.ExecuteOptions,
	reservationOptions *query.ReservationOptions,
) (int16, []int16, *query.ReservedState, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil, nil
}

// Check if we should use an existing reserved connection

// Existing reserved conns are actively held (never idle in the
// regular pool), so PostgreSQL's idle timeout cannot have closed
// the socket between uses. InitiateCopyFromStdin runs directly
// here without a stale-socket retry hop.
// Stamp the vpid before entering COPY mode: once
// InitiateCopyFromStdin succeeds the backend rejects SET until
// CopyDone/CopyFail, so this is the only window to tag the
// backend for lock-detection during long-running COPYs.

// InitiateCopyFromStdin distinguishes two failure modes:
//   - Connection-level error (broken socket): conn is dead, release it.
//   - PG ErrorResponse + drained ReadyForQuery (e.g., "column does
//     not exist", "conflicting options"): conn is back in a clean
//     'I'/'T'/'E' state and is safe to reuse. The reserved conn may
//     still be holding other reasons (transaction, temp table), so
//     destroying it here would orphan that state and force every
//     subsequent statement to fail with "reserved connection not
//     found". Return the current state alongside the error so the
//     gateway can keep tracking the conn.

// New reserved conn — wire BEGIN-if-needed and InitiateCopyFromStdin
// through the validate hook so that the first writes on this
// freshly acquired socket can be transparently retried on a fresh
// socket if the pooled conn was silently closed by PostgreSQL.
// Capture format / columnFormats via closure for use after acquisition.

// Stamp vpid before InitiateCopyFromStdin: SET is rejected
// once the backend enters COPY mode, and the *regular.Conn
// is the same underlying socket that NewReservedConn will
// promote to a *reserved.Conn, so the tag carries through.

// validate ran BEGIN via the raw *regular.Conn, which bypassed
// reserved.Conn.BeginWithQuery's bookkeeping. Mark the
// reservation reason explicitly so the txn shows up in
// RemainingReasons / RequiresBegin checks.

// Mark the connection as reserved for COPY. If the connection is also
// reserved for a transaction, this adds the COPY reason alongside it.

// CopySendData sends a chunk of data for an active COPY operation.
func (e *Executor) CopySendData(
	ctx context.Context,
	target *query.Target,
	data []byte,
	options *query.ExecuteOptions,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Get the reserved connection

// Get the pooled connection for COPY operations

// Write CopyData to PostgreSQL

// CopyFinalize completes a COPY operation, sending final data and returning the result.
// Returns ReservedState with the authoritative reservation state from the multipooler.
func (e *Executor) CopyFinalize(
	ctx context.Context,
	target *query.Target,
	finalData []byte,
	options *query.ExecuteOptions,
) (*sqltypes.Result, *query.ReservedState, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Get the reserved connection

// Get the pooled connection for COPY operations

// Send any remaining data first

// Connection is in bad state - close it instead of recycling

// Send CopyDone to signal completion

// Connection is in bad state - close it instead of recycling

// Read CommandComplete response from PostgreSQL

// For a PG ErrorResponse (e.g., constraint violation, type mismatch,
// missing column), ReadCopyDoneResponse drains the trailing
// ReadyForQuery so the socket is back in a clean state. The reserved
// conn may still be holding other reasons (transaction, temp table),
// so we mirror CopyAbort here: remove the COPY reason and release
// only if no other reasons remain. A connection-level failure (broken
// socket) still falls through to Release(ReleaseError).

// Build result

// Remove the COPY reason. If other reasons remain (e.g., transaction),
// keep the connection reserved. Otherwise, release it back to the pool.

// CopyAbort aborts a COPY operation.
// Returns ReservedState with the authoritative reservation state from the multipooler.
func (e *Executor) CopyAbort(
	ctx context.Context,
	target *query.Target,
	errorMsg string,
	options *query.ExecuteOptions,
) (*query.ReservedState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Already cleaned up or never initiated

// Get the reserved connection

// Already cleaned up — return zero state

// If the conn is no longer in COPY mode (CopyReady never added the reason,
// or CopyFinalize already removed it), the backend is back at RFQ and a
// CopyFail here would be a protocol violation. This happens on the
// gateway-side deferred abort that fires after CopyFinalize already
// completed its own cleanup. Just return the current state.

// Get the pooled connection for COPY operations

// Send CopyFail to abort the operation

// Continue to try reading response

// Read ErrorResponse + ReadyForQuery from PostgreSQL.
// After CopyFail, PostgreSQL responds with ErrorResponse then ReadyForQuery.
// ReadCopyFailResponse drains both, leaving the connection in a clean state.

// Connection is in a bad protocol state — release it.
// We intentionally return nil error: abort is best-effort cleanup and
// the caller needs a zero ReservedState to know the connection is gone.

//nolint:nilerr // intentional: abort is best-effort

// Clean abort — remove the COPY reason. If other reasons remain
// (e.g., transaction), keep the connection reserved.

// getUserFromOptions extracts the user from ExecuteOptions.
// Returns "postgres" as default if no user is specified.
func (e *Executor) getUserFromOptions(options *query.ExecuteOptions) string {
	_ = "STUB: not implemented"
	return ""
}

// Default to postgres superuser if no user specified

// scramKeysFromOptions returns the SCRAM passthrough keys carried on the
// request, or (nil, nil) if no keys were forwarded. The connpoolmanager
// consults the passthrough flag before consuming them, so it is always safe
// to pass these through.
func scramKeysFromOptions(options *query.ExecuteOptions) (clientKey, serverKey []byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

// int32ToInt16Slice converts a slice of int32 to int16.
func int32ToInt16Slice(in []int32) []int16 { _ = "STUB: not implemented"; return nil }

// wrapQueryError wraps query execution errors with context.
// PostgreSQL errors (*mterrors.PgDiagnostic) are wrapped like any other error;
// the display boundary (writeError) extracts the underlying diagnostic via errors.As.
func wrapQueryError(err error) error { _ = "STUB: not implemented"; return nil }

// ConcludeTransaction concludes a transaction on a reserved connection.
// The connection may remain reserved if there are other reasons to keep it (e.g., temp tables).
//
// On ROLLBACK, the caller controls which portal pins are dropped via
// releasePortalNames + releaseAllPortals. PostgreSQL closes only the
// cursors created inside the rolled-back transaction block; cursors
// declared outside the explicit block (under autocommit, before BEGIN)
// survive the ROLLBACK. The gateway computes the per-txn diff and
// passes the inside-txn cursor names so the multipooler unpins exactly
// those, matching PG. When releaseAllPortals is true (or no diff is
// supplied — e.g. by an older gateway that doesn't yet compute it), the
// historical "drop every pin" semantics are used.
//
// Returns ReservedState with the authoritative reservation state.
func (e *Executor) ConcludeTransaction(
	ctx context.Context,
	target *query.Target,
	options *query.ExecuteOptions,
	conclusion multipoolerpb.TransactionConclusion,
	releasePortalNames []string,
	releaseAllPortals bool,
) (*sqltypes.Result, *query.ReservedState, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Get the reserved connection

// Connection destroyed — return zero state so gateway clears its tracking

// Execute COMMIT or ROLLBACK using the reserved connection's methods,
// which handle both the SQL execution and reason removal.

// PostgreSQL closes the cursors created inside this transaction
// block at ROLLBACK; cursors declared outside the block (under
// autocommit, before BEGIN) survive. Honor the caller's diff so
// the multipooler pin set tracks PG exactly. When the diff isn't
// supplied (releaseAllPortals==true), fall back to the historical
// "drop every pin" behavior — back-compat for callers that
// haven't been updated yet.

// Commit/Rollback already removed the transaction reason.
// If other reasons remain (e.g., temp tables, portals), the connection stays reserved.

// DiscardTempTables sends DISCARD TEMP on a reserved connection and removes the temp table reason.
// The connection may remain reserved if there are other reasons to keep it (e.g., transaction).
// Returns ReservedState with the authoritative reservation state.
func (e *Executor) DiscardTempTables(
	ctx context.Context,
	target *query.Target,
	options *query.ExecuteOptions,
) (*sqltypes.Result, *query.ReservedState, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Get the reserved connection

// Connection destroyed — return zero state so gateway clears its tracking

// Send DISCARD TEMP to PostgreSQL to drop all temp tables on this backend.

// Remove the temp table reason

// If no other reasons remain, release the connection

// ReleaseReservedConnection forcefully releases a reserved connection regardless of reason.
// Used during client disconnect cleanup. Handles transaction rollback, COPY abort,
// and portal release internally. If any cleanup step fails, the connection is
// tainted and closed so the pool creates a fresh one.
func (e *Executor) ReleaseReservedConnection(
	ctx context.Context,
	target *query.Target,
	options *query.ExecuteOptions,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Nothing to release

// Already cleaned up or timed out

// Step 1: If there's a transaction, rollback.

// Step 2: If there's a COPY reason, send CopyFail and read the response.

// Step 3: If there are temp tables, discard them so the backend is clean
// when returned to the pool.

// Step 4: Release all portals (in-memory only, always succeeds).

// Step 5: Release or close the connection.

// Ensure Executor implements queryservice.QueryService
var _ queryservice.QueryService = (*Executor)(nil)
