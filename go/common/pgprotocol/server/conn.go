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

package server

import (
	"bufio"
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"log/slog"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/multigres/multigres/go/common/mterrors"
	"github.com/multigres/multigres/go/common/pgprotocol/protocol"
	"github.com/multigres/multigres/go/common/sqltypes"
)

const (
	// connBufferSize is the size of read and write buffers.
	connBufferSize = 16 * 1024

	// defaultFlushDelay is the default delay before auto-flushing buffered writes.
	defaultFlushDelay = 100 * time.Millisecond
)

// errQueryCanceled is the sentinel used as the cancel cause for CancelRequest.
// Must be a single instance: CancelQuery sets it via context.WithCancelCause,
// and queryContextError checks it with errors.Is (pointer equality).
var errQueryCanceled = mterrors.NewQueryCanceled()

// errAuthenticationTimeout is the sentinel returned by serve() when the
// startup-phase deadline (authentication_timeout) fires. handleConnection
// matches on this specifically rather than the generic
// os.ErrDeadlineExceeded so that deadlines added in the future for other
// reasons (idle timeout, per-query socket deadline, etc.) are not silently
// suppressed by the same log-noise filter.
var errAuthenticationTimeout = errors.New("authentication timeout")

// Conn represents the server side connection with a PostgreSQL client.
// It handles the wire protocol encoding/decoding and connection state management.
type Conn struct {
	// conn is the underlying network connection.
	conn net.Conn

	// bufferedReader is used for reading from the connection.
	bufferedReader *bufio.Reader

	// bufferedWriter is used for writing to the connection.
	bufferedWriter *bufio.Writer

	// bufMu protects bufferedReader, bufferedWriter, and the
	// startPacket→writePacket critical section. It is held across
	// in-place packet encoding when startPacket reserves space inside
	// the bufferedWriter, so the body can't be split by an interleaved
	// write from the async notification pusher.
	bufMu sync.Mutex

	// outboundPoolBuf holds the listener-level bufpool buffer that
	// startPacket grabbed for an oversize (slow-path) packet. Non-nil
	// only between startPacket and writePacket, and only on the slow
	// path. writePacket returns it to the pool. Protected by bufMu.
	//
	// Stored as the original *[]byte the pool returned (rather than
	// taking &buf inside writePacket) because passing the address of a
	// stack-local slice to bufPool.Put would force the slice header to
	// escape to the heap on every call, even when the slow path
	// doesn't fire — sync.Pool retains its arguments. Going through
	// this field, which already points at heap memory the pool owns,
	// avoids that escape.
	outboundPoolBuf *[]byte

	// inboundPoolBuf is the read-path equivalent of outboundPoolBuf:
	// holds the listener-level bufpool buffer that readMessageBody
	// grabbed for the most recent message body. Non-nil between
	// readMessageBody and returnReadBuffer. Same escape-avoidance
	// rationale — the parameterless returnReadBuffer reads this field
	// instead of taking &buf, so callers don't pay a 24-byte slice-
	// header heap alloc per message.
	//
	// Reads on a Conn are sequential (one in flight at a time), so a
	// single field is enough.
	inboundPoolBuf *[]byte

	// listener is a reference to the listener that accepted this connection.
	listener *Listener

	// handler processes queries for this connection.
	handler Handler

	// credentialProvider supplies the SCRAM hash and rolreplication flag
	// for the authenticating role. A single lookup feeds both SCRAM and
	// the post-auth replication-role gate; the result is cached on
	// credentials below so the gate doesn't have to round-trip again.
	credentialProvider CredentialProvider

	// credentials is the result of a successful credentialProvider lookup,
	// populated before SCRAM starts. The replication-role gate reads
	// IsReplicationRole from here so neither path has to round-trip a
	// second time.
	credentials *Credentials

	// trustAuthProvider enables trust authentication for testing.
	// When set and AllowTrustAuth() returns true, password auth is skipped.
	trustAuthProvider TrustAuthProvider

	// tlsConfig holds the TLS configuration for SSL connections.
	// When set, the server accepts SSLRequest and upgrades to TLS.
	// When nil, SSLRequest is declined with 'N'.
	tlsConfig *tls.Config

	// requireTLS rejects a plaintext StartupMessage. Copied from the
	// listener at accept time. Cancel requests bypass this check.
	requireTLS bool

	// authMetrics receives auth- and TLS-path metric events for this
	// connection. Never nil — a noop is substituted at listener
	// construction when the caller did not supply one — so startup-phase
	// code can call methods unconditionally.
	authMetrics AuthMetricsRecorder

	// sslDone indicates that an SSLRequest has already been handled
	// (accepted or declined) for this connection. Prevents double negotiation.
	sslDone bool

	// tlsHandshakeComplete is set true once handleSSLRequest has accepted
	// SSL ('S') AND the TLS handshake has finished — i.e., c.conn has been
	// reassigned to the *tls.Conn. Used during the auth-timeout error path
	// to decide whether plaintext writes are still intelligible to the
	// client. Tracking this explicitly (rather than type-asserting c.conn)
	// keeps the check robust if c.conn is ever wrapped in instrumentation.
	tlsHandshakeComplete bool

	// tlsServerCert is the parsed leaf certificate offered to the client
	// during the TLS handshake. Captured once at handshake completion and
	// used to compute the tls-server-end-point channel binding hash for
	// SCRAM-SHA-256-PLUS. Nil for plaintext connections.
	tlsServerCert *x509.Certificate

	// gssDone indicates that a GSSENCRequest has already been handled
	// for this connection. Prevents double negotiation.
	gssDone bool

	// logger for connection-specific logging.
	logger *slog.Logger

	// connectionID is a unique identifier for this connection.
	connectionID uint32

	// backendKeyData is the secret key for this backend, used for cancellation.
	backendKeyData uint32

	// notifPush holds the async notification pusher state.
	notifPush *notifPusher

	// Startup parameters sent by the client.
	user     string
	database string
	params   map[string]string

	// replicationMode reflects the parsed `replication` startup parameter.
	// Default ReplicationOff means a normal SQL session. ReplicationPhysical
	// or ReplicationLogical require pg_authid.rolreplication=true on the
	// authenticated role; the post-auth verifier enforces that.
	replicationMode ReplicationMode

	// SCRAM-SHA-256 keys extracted during the client handshake, used for
	// passthrough authentication to the backing PostgreSQL. Nil for non-SCRAM
	// sessions. Zeroized in Close.
	scramClientKey []byte
	scramServerKey []byte

	// protocolVersion is the negotiated protocol version.
	protocolVersion protocol.ProtocolVersion

	// Current transaction state.
	txnStatus protocol.TransactionStatus

	// state holds handler-specific connection state.
	// Handlers can store their own state here by calling SetConnectionState.
	// This allows different handler implementations to maintain their own state.
	state any

	// queryCancelMu protects queryCancelFunc.
	queryCancelMu sync.Mutex
	// queryCancelFunc cancels the current in-flight query context.
	// Set by BeginQueryCancel, cleared by EndQueryCancel.
	queryCancelFunc context.CancelCauseFunc

	// closed indicates whether the connection has been closed.
	closed atomic.Bool

	// deferredPortalDescribe captures Describe('P') messages so an immediately
	// following Execute on the same portal can fold both into one backend
	// call via HandleExecute(includeDescribe=true). When the deferred state
	// is held and a non-Execute message arrives, resolveDeferredPortalDescribe
	// flushes it through HandleDescribe so the wire stays well-formed.
	//
	// Set by handleDescribe('P'); cleared by handleExecute on a name match,
	// by resolveDeferredPortalDescribe otherwise. Per-Conn state is fine —
	// the protocol guarantees one in-flight request per connection.
	deferredPortalDescribe     bool
	deferredPortalDescribeName string

	// discardingUntilSync is the extended-query error-recovery flag. Set
	// when an ErrorResponse is emitted from any extended-query handler
	// (Parse, Bind, Describe, Execute, Close); cleared when handleMessage
	// observes Sync or Query (both flush boundaries). While set, every
	// Parse/Bind/Describe/Execute/Close inbound message is read off the
	// wire and silently discarded — no handler runs, no reply frame is
	// emitted — matching PostgreSQL's "reads and discards messages until
	// a Sync message is reached" rule. Without this gate, a CloseComplete
	// (or ParseComplete, BindComplete, etc.) emitted after ErrorResponse
	// in the same pipelined batch crashes strict drivers like Postgrex.
	discardingUntilSync bool

	// flushTimer is used for auto-flushing buffered writes.
	flushTimer *time.Timer

	// flushDelay is the delay before auto-flushing.
	flushDelay time.Duration

	// ctx is the context for this connection, cancelled when connection closes.
	ctx    context.Context
	cancel context.CancelFunc
}

// newConn creates a new connection.
func newConn(netConn net.Conn, listener *Listener, connectionID uint32) *Conn {
	_ = "STUB: not implemented"
	return nil
}

// Handler will initialize its own state

// Use pooled readers.

// Close closes the connection and releases resources.
func (c *Conn) Close() error { _ = "STUB: not implemented"; return nil }

// Already closed.

// Clean up handler-specific state (if any).
// The state is set to nil so handlers should handle nil-checking.

// Zeroize SCRAM passthrough keys so a post-mortem or memory dump cannot
// recover credentials for this session after close.

// Return pooled resources.

// Defensive cleanup: if a handler panicked between readMessageBody
// and its returnReadBuffer call, the inbound pool buffer is still
// stashed on the Conn — release it now so the pool can recycle it.

// Same defense for the write side: a panic between startPacket and
// writePacket leaves outboundPoolBuf set (writePacket's defer never
// fires) and bufMu held. We can't re-lock here without deadlocking
// our own goroutine, but Close runs after concurrent access has
// stopped so an unlocked Put is safe.

// End writer buffering (flushes and returns to pool). Flush errors
// during teardown are uninteresting — we're closing the socket
// next anyway — so swallow them here.

// SetTxnStatus sets the protocol-level transaction status indicator.
// This value is sent in ReadyForQuery ('Z') messages to inform the client
// of the current transaction state:
//   - protocol.TxnStatusIdle ('I'): not in a transaction
//   - protocol.TxnStatusInBlock ('T'): in a transaction block
//   - protocol.TxnStatusFailed ('E'): in a failed transaction block
func (c *Conn) SetTxnStatus(status protocol.TransactionStatus) { _ = "STUB: not implemented"; return }

// TxnStatus returns the current protocol-level transaction status.
func (c *Conn) TxnStatus() protocol.TransactionStatus {
	_ = "STUB: not implemented"
	return *

	// IsInTransaction returns true if the connection is currently in a transaction.
	new(protocol.TransactionStatus)
}

func (c *Conn) IsInTransaction() bool { _ = "STUB: not implemented"; return false }

// RemoteAddr returns the remote network address.
func (c *Conn) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

// LocalAddr returns the local network address.
func (c *Conn) LocalAddr() net.Addr {
	_ = "STUB: not implemented"
	return *

	// ConnectionID returns the connection ID.
	new(net.Addr)
}

func (c *Conn) ConnectionID() uint32 { _ = "STUB: not implemented"; return 0 }

// Handler returns the protocol handler for this connection.
func (c *Conn) Handler() Handler {
	_ = "STUB: not implemented"

	// User returns the authenticated user.
	return *new(Handler)
}

func (c *Conn) User() string {
	_ = "STUB: not implemented"

	// Database returns the database name.
	return ""
}

func (c *Conn) Database() string {
	_ = "STUB: not implemented"

	// ReplicationMode returns the parsed `replication` startup parameter for this
	// connection. ReplicationOff means the client did not request a replication
	// connection (or sent replication=false).
	return ""
}

func (c *Conn) ReplicationMode() ReplicationMode {
	_ = "STUB: not implemented"
	return *

	// ScramClientKey returns the SCRAM-SHA-256 ClientKey extracted during the
	// client's authentication handshake, or nil if the session did not
	// authenticate via SCRAM. Used for passthrough auth to backend PostgreSQL.
	new(ReplicationMode)
}

func (c *Conn) ScramClientKey() []byte { _ = "STUB: not implemented"; return nil }

// ScramServerKey returns the SCRAM-SHA-256 ServerKey from the user's
// verifier, or nil if the session did not authenticate via SCRAM.
func (c *Conn) ScramServerKey() []byte { _ = "STUB: not implemented"; return nil }

// GetStartupParams returns the startup parameters sent by the client,
// excluding 'user' and 'database' which are handled separately.
func (c *Conn) GetStartupParams() map[string]string { _ = "STUB: not implemented"; return nil }

// Context returns the connection's context.
func (c *Conn) Context() context.Context {
	_ = "STUB: not implemented"

	// GetConnectionState returns the handler-specific connection state.
	// Returns nil if no state has been set.
	return *new(context.Context)
}

func (c *Conn) GetConnectionState() any {
	_ = "STUB: not implemented"

	// SetConnectionState sets the handler-specific connection state.
	// This allows handlers to store their own state per connection.
	return *new(any)
}

func (c *Conn) SetConnectionState(state any) {
	_ = "STUB: not implemented"

	// BackendKeyData returns the secret key for this connection.
	return
}

func (c *Conn) BackendKeyData() uint32 { _ = "STUB: not implemented"; return 0 }

// BeginQueryCancel creates a child context for the current query that can be
// independently canceled. Returns the query context that should be passed to
// handler methods. Must be paired with EndQueryCancel.
func (c *Conn) BeginQueryCancel() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// EndQueryCancel clears the query cancel function, signaling that no query
// is in flight. Calls the cancel function with nil to release context resources.
func (c *Conn) EndQueryCancel() { _ = "STUB: not implemented"; return }

// CancelQuery cancels the in-flight query on this connection.
// Returns true if a query was in flight and was canceled.
func (c *Conn) CancelQuery() bool { _ = "STUB: not implemented"; return false }

// queryContextError maps context-related errors to the appropriate PostgreSQL
// protocol error. Cancel requests take priority since they indicate explicit
// user action. If no context-related error is detected, the original error is
// returned unchanged.
func queryContextError(queryCtx context.Context, err error) error {
	_ = "STUB: not implemented"
	return nil
}

// CancelRequest sets errQueryCanceled as the cause on the query context.

// Statement timeout: the handler applies context.WithTimeout for statement
// timeouts, and DeadlineExceeded propagates up through the error chain.

// returnReader returns the buffered reader to the pool.
func (c *Conn) returnReader() { _ = "STUB: not implemented"; return }

// startWriterBuffering begins buffering writes.
func (c *Conn) startWriterBuffering() { _ = "STUB: not implemented"; return }

// flush flushes any buffered writes.
func (c *Conn) flush() error { _ = "STUB: not implemented"; return nil }

// Flush flushes any buffered writes to the client.
// This is exposed for external callers (like gateway handler) that need to
// ensure messages are sent immediately (e.g., CopyInResponse).
func (c *Conn) Flush() error {
	_ = "STUB: not implemented"

	// endWriterBuffering flushes any remaining buffered data and returns
	// the writer to the pool. Returns the Flush error (broken-pipe, write
	// timeout, etc.) so callers in the serve() loop can propagate it for
	// connection-level cleanup; the writer is still detached and pooled
	// even when Flush fails, so the connection state is consistent.
	return nil
}

func (c *Conn) endWriterBuffering() error { _ = "STUB: not implemented"; return nil }

// Return to pool regardless of flush outcome — the writer is in
// a defined state (Reset clears its internal cursors), and
// stranding it on the Conn would leak the pooled buffer.

// canSendPlaintextStartupError reports whether a best-effort plaintext
// ErrorResponse during the startup phase would be intelligible to the
// client. If the client sent SSLRequest and the server already answered
// 'S' but the TLS handshake hasn't completed, the underlying conn is
// still plaintext while the client is waiting for encrypted bytes — so
// any plaintext write would surface as garbage. Every other startup
// state (raw plaintext or fully upgraded TLS) can carry the reply.
func (c *Conn) canSendPlaintextStartupError() bool { _ = "STUB: not implemented"; return false }

// generateBackendKey generates a cryptographically secure random backend key for cancellation.
// The backend key is sent to the client in the BackendKeyData message and is used
// to authenticate query cancellation requests.
func generateBackendKey() uint32 { _ = "STUB: not implemented"; return 0 }

// Fallback to time-based key if crypto/rand fails (should never happen).
// This maintains backward compatibility with the TODO implementation.

// serve is the main command processing loop for the connection.
// It reads messages from the client and processes them until the connection is closed.
//
// Write buffering is managed at this level rather than per-handler. The
// pgwire client pipelines a batch of extended-protocol messages
// (Parse + Bind + Describe + Execute + Sync) and only blocks reading
// after Sync — so coalescing the six reply messages into one syscall
// per batch matches what the client is actually waiting on, and drops
// the syscall.write count by ~5x on sysbench prepared workloads.
//
// The window opens lazily via startWriterBuffering() before each
// handleMessage call (idempotent: no-op if already attached) and is
// released back to the writer pool on flush boundaries — Sync, Flush,
// simple Query, and the error path. That keeps the steady-state for
// pipelined batches at one buffer / one flush, while still returning
// the writer to the pool when the connection genuinely idles between
// batches.
func (c *Conn) serve() error {
	_ = "STUB: not implemented"
	// Bound the startup phase (SSL/GSS negotiation, StartupMessage, SCRAM
	// exchange) — equivalent to PostgreSQL's authentication_timeout. A
	// stalled or malicious client cannot pin this goroutine past the
	// deadline. The deadline is cleared once authentication completes so
	// the main command loop runs without an I/O timeout.
	return nil
}

// First, handle the startup phase.

// errAuthRejected: a FATAL was already sent during the auth flow;
// no AuthenticationOk was emitted and the connection was not
// registered. Skip the command loop entirely so a misbehaving
// client cannot send messages on a half-completed session, and
// don't write a second error frame.

// Map a deadline-exceeded I/O error during startup to a clean
// PG-format FATAL with SQLSTATE 08006 so libpq surfaces it as
// "canceling authentication due to timeout" instead of a raw
// I/O timeout. The startup deadline is still active here — any
// write would inherit the expired deadline and fail — so clear
// it first using a brief grace window for the error reply.

// Brief grace window: long enough to flush the error
// to a well-behaved client, short enough that a wedged
// client can't keep this goroutine pinned.

// Return the sentinel so handleConnection can suppress its
// redundant Error log without also swallowing unrelated
// deadline-exceeded errors that future code paths may surface.

// Set a brief write-deadline grace window so the best-effort
// error reply doesn't race against the still-active auth
// deadline (which may fire mid-write on a slow client).

// Try to send an error response before closing — but only if
// it would be intelligible. Same SSL-accepted-but-handshake-
// pending guard as the timeout path above: a plaintext
// ErrorResponse on a connection where the client expects
// encrypted bytes would surface as garbage.

// If the error is already a PgDiagnostic (e.g., duplicate
// SSLRequest with native SQLSTATE), send it directly.
// Otherwise, wrap with MTE01.

// Cancel requests close the connection inside handleStartup and
// return nil — there is no authenticated session to enter the
// command loop for, and SetDeadline on the closed fd would fail
// and surface as a spurious Error log. Bail out cleanly here.

// Authentication complete — clear the startup deadline before the
// command loop. Subsequent reads must not inherit the auth-phase
// deadline (idle clients are normal, not a protocol violation). If
// the clear fails (rare — typically only when the fd is already
// closed), abort instead of entering the loop with a stale deadline
// that would trip the very next read.

// Main command loop.

// Check if connection is closed.

// Read the message type (1 byte).

// EOF or connection error - close gracefully.

// Open a buffering window if one isn't already attached. Per-
// handler bodies no longer manage this themselves; they just
// write packets and let the loop decide when to flush.

// Process the message based on type.

// Send error response and continue (unless it's a fatal error).

// Best-effort flush of the error reply. Already returning
// the original handler error, so swallow flush errors here.

// For now, close connection on any error.

// Flush at end-of-batch boundaries. For pipelined messages
// (Parse / Bind / Describe / Execute / Close) the buffer
// stays attached so the next inbound message in the same
// batch lands in the same flush. Flush errors here mean the
// socket is broken — propagate so serve() can tear down the
// connection instead of looping until the next ReadMessageType
// fails (which would lose the original write-error context).
// MsgSync and MsgQuery are end-of-batch boundaries: flush and
// release the writer back to the pool so an idle connection
// doesn't pin a 16 KB writer between batches.
//
// MsgFlush itself flushes inside handleMessage (and stays
// buffered, since more pipelined messages typically follow);
// it doesn't need a release here.

// handleMessage processes a single message from the client.
//
// A Describe('P') is held by handleDescribe in case the very next message is
// Execute on the same portal — that pair is fused into one backend call. Any
// other message (including a follow-up Describe of either type) flushes the
// held state here first so the wire reply order matches what an unfused
// Describe(P)+Execute would produce.
//
// Extended-query error recovery: once an ErrorResponse has been emitted in
// the current batch, c.discardingUntilSync is set and every subsequent
// Parse / Bind / Describe / Execute / Close is read off the wire and
// discarded here without dispatching to a handler. Sync clears the flag
// and proceeds to handleSync (which emits the ReadyForQuery the client is
// waiting on). This matches PostgreSQL's documented behavior and prevents
// the post-error frames (notably CloseComplete) that crash strict drivers.
func (c *Conn) handleMessage(msgType byte) error { _ = "STUB: not implemented"; return nil }

// handleExecute consumes the deferred describe directly (it folds
// it into the backend call rather than flushing it separately).
// handleExecute itself re-checks the drain flag after the
// fallback resolveDeferredPortalDescribe path.

// Signal connection should close

// Every other message must run after any pending portal describe is
// flushed so the wire stays well-formed.

// resolveDeferredPortalDescribe may have flushed via HandleDescribe
// and entered drain mode by emitting an ErrorResponse. In that case
// the inbound message that triggered the flush (Parse/Bind/Close/…)
// must also be discarded — otherwise its reply frame would land
// between ErrorResponse and ReadyForQuery and crash strict drivers.

// Read and discard the message length (Int32, always 4: just
// the length field itself). The 'H' type byte was already
// consumed by serve()'s ReadMessageType; if we don't consume
// the length here, the next ReadMessageType picks up 0x00
// and treats it as a bogus type.

// Push any buffered bytes (including a deferred Describe('P')
// that resolveDeferredPortalDescribe flushed into the buffer
// just above) out to the client. We do this here rather than
// in serve()'s post-dispatch so direct callers of
// handleMessage(MsgFlush) — e.g. the unit tests in
// extended_query_test.go — see the same client-visible
// behavior as the production read loop. The buffer stays
// attached: more pipelined messages typically follow.

// handleQuery handles a 'Q' (Query) message - simple query protocol.
// Supports multiple statements in a single query (e.g., "SELECT 1; SELECT 2;").
//
// The buffering window is opened by the serve() loop and torn down
// (flushed + writer returned to the pool) after this returns, since
// MsgQuery is treated as a flush boundary. So we just write packets
// here and let the loop handle the syscall.
func (c *Conn) handleQuery() error {
	_ = "STUB: not implemented"
	// Read the query string.
	return nil
}

// Create a cancelable query context so cancel requests can interrupt this query.

// Track state for current result set.
// This is reset when we complete a result set (when CommandTag is set).

// Execute the query via the handler with streaming callback.
// The callback will be invoked multiple times for:
// 1. Large result sets (streamed in chunks)
// 2. Multiple statements in a single query (each potentially with large result sets)

// Handle empty query (nil result signals empty query).

// Send notices immediately (zero-buffering delivery).
// Notices may arrive as standalone Results (no rows/CommandTag) from the gRPC
// streaming path, or bundled with the final Result that has a CommandTag.

// On first callback with fields for this result set, send RowDescription.
// Use nil check (not len > 0) because zero-column results (e.g., "select union select")
// have a non-nil empty Fields slice and still require a RowDescription message.

// Send all data rows in this chunk.

// If CommandTag is set, this is the last packet of the current result set.

// Reset state for next result set (if any).

// Send ReadyForQuery after all statements have been processed.
// serve() flushes the buffer once we return.

// handleParse handles a 'P' (Parse) message - extended query protocol.
// Parse message format:
// - Statement name (string, null-terminated)
// - Query string (string, null-terminated)
// - Number of parameter data types (int16)
// - Parameter data type OIDs ([]uint32)
//
// The buffering window is opened by serve(); we leave it attached on
// return so the next pipelined message in the same batch lands in
// the same flush. Sync (or an explicit Flush) is what pushes bytes
// to the client.
func (c *Conn) handleParse() error {
	_ = "STUB: not implemented"
	// Read message length.
	return nil
}

// Read message body.

// Parse the message.

// Call the handler to validate and prepare the statement.
// The handler is responsible for storing any state it needs.

// Do NOT send ReadyForQuery here. In the extended query protocol, the client
// pipelines Parse + Describe + Sync (or Parse + Bind + Execute + Sync).
// ReadyForQuery must only be sent in response to Sync. Sending it here would
// cause protocol desynchronization: pgx would read the premature ReadyForQuery
// and think the pipeline is done, but stale responses from subsequent messages
// (Describe, Sync) would corrupt the next query's response stream.
// The error packet stays buffered until Sync flushes the batch — same shape
// as upstream PostgreSQL, which also defers error delivery to Sync/Flush.

// Send ParseComplete message. Stays buffered for the rest of the batch.

// handleBind handles a 'B' (Bind) message - extended query protocol.
// The serve() loop owns the buffering window; this handler just
// writes packets and returns.
func (c *Conn) handleBind() error {
	_ = "STUB: not implemented"
	// Read message length.
	return nil
}

// Read message body.

// Parse the message.

// Read parameter format codes

// Read parameters

// Read result format codes

// Call the handler to create and bind the portal with parameters.

// Do NOT send ReadyForQuery here — same reasoning as handleParse.
// ReadyForQuery is sent only in response to Sync. The error packet
// stays buffered until Sync flushes the batch.

// Send BindComplete message. Stays buffered for the rest of the batch.

// handleExecute handles an 'E' (Execute) message - extended query protocol.
// Execute message format:
// - Portal name (string, null-terminated)
// - Max rows to return (int32, 0 = no limit)
//
// serve() owns the buffering window. The reply (zero-or-more DataRow
// + CommandComplete) stays in the buffer until the trailing Sync
// flushes the batch.
func (c *Conn) handleExecute() error {
	_ = "STUB: not implemented"
	// Read message length.
	return nil
}

// Read message body.

// Parse the message.

// If a Describe('P') was deferred for this exact portal, fold it into
// this Execute call — the handler/executor will fetch RowDescription
// alongside the data rows in one backend round trip. A mismatched
// portal name forces the deferred describe to flush via the handler
// before this Execute proceeds.

// If the deferred Describe flush errored, we're now in
// drain mode. Don't run HandleExecute — its
// RowDescription / DataRow / CommandComplete frames
// would leak between ErrorResponse and ReadyForQuery.

// Create a cancelable query context so cancel requests can interrupt this execution.

// Track state for streaming results.

// Call the handler to execute the portal with streaming callback.
// The handler is responsible for retrieving the portal and executing it.

// Send notices immediately (zero-buffering delivery).

// On first callback with fields, send RowDescription.
// Use nil check (not len > 0) because zero-column results still require RowDescription.
// When a folded Describe('P') is in flight, the same RowDescription resolves it.

// Send all data rows in this chunk.

// If CommandTag is set, this is the last packet.

// DML folded with a deferred Describe never sends Fields through
// the callback, but the protocol still requires NoData before
// CommandComplete. Emit it here so the wire order matches a
// non-folded Describe(P)+Execute sequence.

// handleDescribe handles a 'D' (Describe) message - extended query protocol.
// Describes either a prepared statement ('S') or a portal ('P').
// serve() owns the buffering window; reply packets stay buffered
// until Sync flushes the batch.
func (c *Conn) handleDescribe() error {
	_ = "STUB: not implemented"
	// Read message length.
	return nil
}

// Read the full body (type byte + null-terminated name) into a
// pooled buffer.

// String() copies, so the body buffer can be returned to the pool.

// Describe('P') is captured and held: if the next message is Execute on
// the same portal, the two fold into a single backend call. Otherwise
// the next non-Execute handler (including a follow-up Describe of any
// type) flushes the held state via resolveDeferredPortalDescribe in
// handleMessage before its own reply runs. No writes happen here, so
// skip buffer acquisition entirely.

// Flush errors at this teardown are best-effort — if the handler
// itself returned an error we already have it in flight, and if
// not, serve()'s post-dispatch will see the next ReadMessageType
// fail and tear down the connection with the right context.

// Statement describe ('S') — call handler synchronously. Any prior
// Describe('P') was already flushed by handleMessage before this call.

// Send ParameterDescription only for statement describes ('S').
// PostgreSQL protocol: Describe('S') returns ParameterDescription + RowDescription/NoData,
// but Describe('P') returns only RowDescription/NoData — no ParameterDescription.
// The 'P' branch above already returned, so this is reachable only for
// 'S' under the contract; the explicit check is defense in depth in case
// a future caller routes a different byte through this path.

// Send RowDescription or NoData based on whether we have field info.
// Use nil check (not len > 0) because zero-column results (e.g., "SELECT FROM foo")
// have a non-nil empty Fields slice and still require a RowDescription message.

// Send NoData when there are no fields (e.g., DML statements).

// resolveDeferredPortalDescribe flushes a pending Describe('P') by calling
// the handler synchronously and writing the wire-correct RowDescription
// (or NoData) before any subsequent message's reply is generated. Called
// from every non-Execute handler that runs while a deferred Describe is
// outstanding.
//
// No-op when nothing is deferred.
func (c *Conn) resolveDeferredPortalDescribe() error { _ = "STUB: not implemented"; return nil }

// handleClose handles a 'C' (Close) message - extended query protocol.
// Closes either a prepared statement ('S') or a portal ('P'). serve()
// owns the buffering window; CloseComplete stays buffered until the
// trailing Sync flushes the batch.
func (c *Conn) handleClose() error {
	_ = "STUB: not implemented"
	// Read message length.
	return nil
}

// Read the full body (type byte + null-terminated name) into a
// pooled buffer.

// String() copies, so the body buffer can be returned to the pool.

// Call the handler.

// Send CloseComplete. Stays buffered for the rest of the batch.

// maybeDispatchDrain is the extended-query error-drain gate. When
// c.discardingUntilSync is set it routes the current inbound message
// without dispatching to its normal handler — Sync/Query clear the
// flag and signal the caller to fall through, Flush flushes buffered
// bytes without writing a reply, Parse/Bind/Describe/Execute/Close get
// their bodies drained off the wire, and Terminate falls through to
// the normal teardown.
//
// Returns (handled=true, err) when the message was absorbed here and
// the caller should return immediately; (handled=false, nil) when the
// caller should continue to its normal dispatch.
//
// Called twice from handleMessage — once at the top for messages that
// arrive while drain mode is already set, and once again after
// resolveDeferredPortalDescribe runs (which can itself enter drain
// mode by flushing a failing deferred Describe).
func (c *Conn) maybeDispatchDrain(msgType byte) (handled bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Flush boundaries — clear the flag and signal the caller to
// fall through so handleSync emits ReadyForQuery (or
// handleQuery runs a fresh simple-query exchange).

// Connection teardown — fall through to normal dispatch.

// Flush still pushes any buffered ErrorResponse to the client
// but writes no reply of its own. Mirror the length-consume +
// flush dance that handleMessage's MsgFlush branch does.

// drainExtendedQueryMessage reads and discards a single extended-query
// message body while the connection is in error-drain mode. The bytes
// must still be consumed off the wire so the next ReadMessageType aligns
// to the next message header — only the handler call and the reply
// frame are suppressed. readMessageBody handles zero-length bodies as a
// no-op, so no special case is needed here.
func (c *Conn) drainExtendedQueryMessage() error { _ = "STUB: not implemented"; return nil }

// writeExtendedQueryError writes an ErrorResponse and enters error-drain
// mode. PostgreSQL's protocol spec requires that after an extended-query
// message errors, the backend discards every subsequent message until
// Sync and only then emits ReadyForQuery. Strict drivers (Postgrex, JDBC)
// treat any frame between ErrorResponse and ReadyForQuery as protocol
// corruption and drop the connection.
//
// Use this from any handler that processes an extended-query inbound
// message (Parse, Bind, Describe, Execute, Close) when reporting an
// error to the client. Do NOT use it from handleQuery (simple Query is
// its own self-contained ErrorResponse + ReadyForQuery flow) or from
// handleSync (Sync itself emits ReadyForQuery and is the boundary the
// drain ends at).
func (c *Conn) writeExtendedQueryError(err error) error { _ = "STUB: not implemented"; return nil }

// handleSync handles an 'S' (Sync) message - extended query protocol.
// Sync indicates the end of an extended query cycle and transaction boundary.
// Always sends ReadyForQuery in response.
//
// Sync is a flush boundary: serve() will release the buffering window
// (flushing any queued ParseComplete / BindComplete / RowDescription /
// DataRow / CommandComplete from earlier messages in the batch, plus
// the ReadyForQuery we write here) once we return.
func (c *Conn) handleSync() error {
	_ = "STUB: not implemented"
	// Read (and discard) message length.
	return nil
}

// Call the handler.

// Even if handler returns error, we still send ReadyForQuery after Sync.

// Always send ReadyForQuery after Sync.

// notifPusher holds state for async notification delivery.
type notifPusher struct {
	ch     chan *sqltypes.Notification
	cancel context.CancelFunc
}

// EnableAsyncNotifications starts a background goroutine that delivers
// notifications from notifCh to the client socket. Must be called at most once.
// Returns a channel that the caller should send notifications to.
func (c *Conn) EnableAsyncNotifications(ctx context.Context) chan<- *sqltypes.Notification {
	_ = "STUB: not implemented"
	return nil
}

// writeNotificationResponseMsg acquires bufMu through
// startPacket/writePacket; each notification packet is
// committed atomically under that lock, so it can't be
// interleaved with a synchronous handler's packet.

// StopAsyncNotifications stops the background notification pusher.
func (c *Conn) StopAsyncNotifications() { _ = "STUB: not implemented"; return }

// FlushPendingNotifications drains all pending notifications from
// the async pusher channel and writes them to the client socket.
// Called synchronously after each query completes (before
// ReadyForQuery) to deliver notifications that arrived during query
// execution.
//
// Each notification is written via writeNotificationResponseMsg,
// which acquires bufMu inside startPacket/writePacket per packet.
// Multiple notifications may interleave with the synchronous query
// handler's writes between packets, but every individual packet is
// committed atomically under the lock, so packet bodies can never be
// split.
func (c *Conn) FlushPendingNotifications() error { _ = "STUB: not implemented"; return nil }

// writeNotificationResponseMsg writes a NotificationResponse ('A')
// packet: int32 pid, null-terminated channel, null-terminated payload.
// Goes through startPacket/writePacket — caller does NOT need to hold
// bufMu (the helpers manage it).
func (c *Conn) writeNotificationResponseMsg(pid int32, channel, payload string) error {
	_ = "STUB: not implemented"
	return nil
}
