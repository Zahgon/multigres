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
	"errors"
)

// StartupMessage represents a parsed startup message from the client.
type StartupMessage struct {
	ProtocolVersion uint32
	Parameters      map[string]string
}

// ReplicationMode captures the value of the `replication` startup parameter
// per PostgreSQL's replication protocol. The default (parameter omitted or
// "false") is ReplicationOff: a normal SQL connection.
//
// See: https://www.postgresql.org/docs/17/protocol-replication.html
type ReplicationMode int

const (
	// ReplicationOff means the client did not request a replication
	// connection, or sent replication=false. The session uses the standard
	// extended-query / simple-query protocol.
	ReplicationOff ReplicationMode = iota

	// ReplicationPhysical (replication=true / on / 1 / yes) opens a physical
	// walsender stream. The role must have rolreplication=true (or
	// rolsuper=true) in pg_authid.
	ReplicationPhysical

	// ReplicationLogical (replication=database) opens a logical-replication
	// walsender connected to a specific database. Same role requirement as
	// physical replication.
	ReplicationLogical
)

// parseReplicationMode interprets the `replication` startup parameter using
// PostgreSQL's parsing rules (src/backend/utils/misc/guc.c parse_bool_with_len).
// PG accepts case-insensitive on/off, true/false, yes/no, 1/0 and the
// single-character abbreviations t/f/y/n, plus the literal "database" for
// logical-replication connections.
//
// Returns an InvalidParameterValue PgDiagnostic for unrecognized values so
// the gateway can reject them at the same protocol stage PostgreSQL would.
func parseReplicationMode(value string) (ReplicationMode, error) {
	_ = "STUB: not implemented"
	return *new(ReplicationMode), nil
}

// handleStartup handles the initial connection startup phase.
// This includes SSL/GSSAPI encryption negotiation and processing the startup message.
// Returns an error if the startup fails.
func (c *Conn) handleStartup() error { _ = "STUB: not implemented"; return nil }

// readAndDispatchStartup reads a startup packet and dispatches based on protocol code.
// This method is called both for the initial startup and after encryption negotiation
// (SSL or GSSAPI) to handle the fallback ordering defined in the PostgreSQL protocol:
// a client may try GSSENCRequest after SSLRequest is declined, or vice versa.
// See: https://www.postgresql.org/docs/17/protocol-flow.html
func (c *Conn) readAndDispatchStartup() error { _ = "STUB: not implemented"; return nil }

// Distinguish the two ways a client can hit this gate so the
// fleet-validation alert (MUL-420) can tell "client never
// asked for TLS" from "server declined the SSLRequest" — the
// latter would indicate misconfiguration that --pg-require-ssl
// alone won't catch.

// Returning a *PgDiagnostic lets serve() emit the FATAL
// ErrorResponse via its standard startup-error path (which
// also clears the auth deadline and closes the connection).

// handleSSLRequest handles an SSL negotiation request from the client.
// If TLS is configured, accepts with 'S' and upgrades the connection to TLS.
// If TLS is not configured, declines with 'N'.
// After responding, reads the next startup packet which may be a StartupMessage
// or a GSSENCRequest (per PostgreSQL protocol fallback ordering).
func (c *Conn) handleSSLRequest() error { _ = "STUB: not implemented"; return nil }

// No TLS configured, decline SSL.

// Read next packet — could be StartupMessage or GSSENCRequest (fallback).

// Accept SSL and upgrade to TLS.

// Buffer-stuffing attack prevention (CVE-2021-23222):
// If there is buffered data after we sent 'S' but before the TLS handshake,
// a MITM may have injected unencrypted data.

// Wrap the TLS config so dynamic-cert deployments (GetCertificate /
// GetConfigForClient, typical for SNI multi-tenant edges) capture the
// actually-selected leaf cert into c.tlsServerCert during the handshake.
// Static Certificates[0] deployments are handled by the post-handshake
// fallback below and skip the wrapper.

// Perform TLS handshake. Time it for mg.gateway.tls.handshake.duration
// and classify the outcome so the fleet-validation alert can tell a
// crypto failure from a client tear-down. net.ErrClosed / io.EOF map
// to client_aborted; everything else is treated as a handshake_failure.

// Replace the underlying connection and reset the buffered reader
// to read from the TLS connection. The buffered writer is nil during
// startup (lazy init via startWriterBuffering), so writeRawByte
// falls back to c.conn directly — after this swap, writes go
// through TLS.

// Static-cert fallback: when the deployment only populates Certificates,
// the dynamic wrapper above is a no-op and the cert was not captured.
// Recover it from Certificates[0] here. A parse failure is non-fatal —
// auth simply falls back to SCRAM-SHA-256 without channel binding.

// Read the actual startup message over the encrypted connection.

// handleGSSENCRequest handles a GSSAPI encryption request.
// We don't support GSSAPI encryption, so we always decline with 'N'.
// After declining, reads the next startup packet which may be a StartupMessage
// or an SSLRequest (per PostgreSQL protocol fallback ordering).
func (c *Conn) handleGSSENCRequest() error { _ = "STUB: not implemented"; return nil }

// Read next packet — could be StartupMessage or SSLRequest (fallback).

// handleCancelRequest handles a query cancellation request.
// This is sent by clients to cancel a running query on another connection.
// Per PostgreSQL protocol, the cancel connection is always closed after processing.
func (c *Conn) handleCancelRequest(reader *MessageReader) error {
	_ = "STUB: not implemented"
	// Read the process ID (connection ID).
	return nil
}

// Read the secret key.

// Delegate to the cancel handler which may forward to a remote gateway.

// No cross-gateway handler (e.g., tests); try to cancel locally.

// splitOptionsTokens splits a PGOPTIONS string on unescaped whitespace.
// Backslash-escaped characters (e.g. `\ ` for a literal space) are preserved
// with the backslash removed.
func splitOptionsTokens(s string) []string { _ = "STUB: not implemented"; return nil }

// parseOptions parses a PGOPTIONS string into individual key-value pairs.
// It supports:
//   - `-c key=value` and `-ckey=value`
//   - `--key=value` (hyphens in key converted to underscores)
//   - Multiple flags in a single string
func parseOptions(options string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// -c key=value (space-separated)

// -ckey=value (no space)

// --key=value

// Convert hyphens to underscores in key.

// handleStartupMessage processes a startup message and extracts connection parameters.
func (c *Conn) handleStartupMessage(protocolVersion uint32, reader *MessageReader) error {
	_ = "STUB: not implemented"
	return nil
}

// Store the protocol version.

// Parse key-value pairs until we hit a null byte.

// Read the key.

// Empty key means we've reached the end.

// Read the value.

// Store the parameter.

// Parse PGOPTIONS if present.

// `replication` is a startup-phase-only protocol parameter; PG does
// not accept it as a `-c replication=...` GUC inside PGOPTIONS.
// Drop it here so the gateway doesn't honor a source PG would
// reject. The auth gate is unaffected either way (rolreplication
// is enforced when the direct startup field is set).

// Extract required parameters.

// Default database to user if not specified.

// Parse the optional `replication` startup parameter. Replication
// connections (physical or logical) follow the same auth path but the
// role must additionally satisfy pg_authid.rolreplication=true.
// On parse failure, return the PgDiagnostic; serve()'s startup-error
// path writes it to the client and closes the connection — matching
// PG's behavior of rejecting unrecognized `replication` values before
// authentication runs.
//
// Strip the key from c.params after parsing: `replication` is a
// protocol-only startup parameter, not a GUC. Leaving it in the map
// would let it flow through GetStartupParams → session settings →
// `SET SESSION "replication" = ...` on the backend, which PG rejects
// as unrecognized. The same reason `options` is deleted just above.

// Now perform authentication.

// errAuthRejected signals that the auth flow rejected the client and a FATAL
// message has already been written. The error propagates up to serve(),
// which recognizes it and closes the connection cleanly without entering
// the command loop or writing a second error frame. Without this propagation
// the connection would proceed to the command loop after a rejection — a
// well-behaved client closes after FATAL and the next read returns EOF, but
// a malicious or buggy client could attempt to send messages on a session
// where AuthenticationOk was never emitted and RegisterConn was never called.
//
// All sendAuthError-style helpers return this on a successful FATAL write so
// the post-auth completion sequence is skipped uniformly.
var errAuthRejected = errors.New("auth rejected; FATAL already sent")

// authenticate performs authentication with the client.
// If a TrustAuthProvider is configured and allows the user, trust auth is used.
// Otherwise, SCRAM-SHA-256 authentication is performed.
//
// On success, this also enforces post-auth role attribute checks (today
// rolreplication for replication startup connections) and emits the
// AuthenticationOk → BackendKeyData → ParameterStatus → ReadyForQuery
// completion sequence. The role-attribute check runs *before*
// AuthenticationOk; native PostgreSQL sequences the same check as
// SASLFinal → AuthenticationOk → (InitPostgres rolreplication check) → FATAL,
// so multigres collapses two server-to-client frames into one for rejected
// replication clients. libpq, pgx, and JDBC all accept ErrorResponse at this
// stage either way — the wire-visible difference is one fewer frame and no
// successful-handshake-then-rejection optic for the client.
func (c *Conn) authenticate() (err error) {
	_ = "STUB: not implemented"
	// Track the outcome label for mg.gateway.auth.attempts. Each rejection
	// site below assigns a specific value before returning so we don't lose
	// the cause when sendAuthError-style helpers fold it into errAuthRejected.
	return nil
}

// Check if trust auth is allowed for this connection. errAuthRejected
// is propagated unchanged so serve() can short-circuit out of the
// startup phase without entering the command loop.

// authenticateSCRAM returns a specific outcome label alongside
// the error so this caller doesn't have to reverse-engineer it
// from errAuthRejected (which folds the cause).

// Replication startup parameter requires rolreplication=true on the role.
// Done post-auth so we don't leak which roles exist for unauthenticated
// clients.

// SCRAM succeeded but the role lacks rolreplication. Tagged as
// its own outcome so MUL-420's fleet validation alert can
// distinguish role-attribute rejections from login_disabled
// without scanning logs.

// classifyAuthError maps an auth-path error to the closed set of outcome
// labels used by mg.gateway.auth.attempts and mg.gateway.auth.scram.duration.
// Order matters: the SCRAM sentinels must be checked before the generic
// errAuthRejected fallback so cause-specific labels survive the wrapping
// done by sendAuthError / sendScramFatal.
func classifyAuthError(err error) string { _ = "STUB: not implemented"; return "" }

// authenticateTrust performs trust authentication (no password required).
// This is used in tests to simulate Unix socket trust authentication.
//
// Trust auth has no over-the-wire negotiation step before AuthenticationOk
// — the caller (authenticate) is responsible for the success sequence.
func (c *Conn) authenticateTrust() error { _ = "STUB: not implemented"; return nil }

// finishAuth emits the post-authentication completion sequence shared by
// both trust and SCRAM paths: AuthenticationOk, BackendKeyData, the
// initial ParameterStatus run, and ReadyForQuery.
func (c *Conn) finishAuth() error { _ = "STUB: not implemented"; return nil }

// Send BackendKeyData for query cancellation.

// Register connection for cancel request lookup now that the client knows the PID.

// Send initial ParameterStatus messages.

// Notify handlers that opted into the established hook *before*
// sending ReadyForQuery. Otherwise the client returns from Connect
// as soon as it sees ReadyForQuery, races ahead, and may observe
// the conn before the hook records per-connection startup state.

// Send ReadyForQuery to indicate we're ready to receive commands.

// verifyReplicationRole enforces pg_authid.rolreplication for clients that
// requested a replication startup connection (replication=true /
// replication=database). Skipped for normal sessions.
//
// The flag was fetched alongside the SCRAM hash during authenticateSCRAM
// and cached on c.credentials, so this gate is a constant-time field
// check rather than a second pooler round-trip.
//
// For the trust-auth path c.credentials is unset, so we fall back to
// fetching from the credential provider here. Trust auth is test-only.
//
// Mismatches produce PG's exact wording with SQLSTATE 42501, matching
// what native PostgreSQL emits in walsender startup. The error is sent
// FATAL so libpq tears down the connection.
func (c *Conn) verifyReplicationRole() error { _ = "STUB: not implemented"; return nil }

// Use cached credentials from SCRAM if available.

// Trust-auth path: no SCRAM lookup happened, so we need to fetch the
// flag here. Without a credential provider configured, fail closed.

// Lookup failure (including ErrUserNotFound / ErrLoginDisabled /
// ErrPasswordExpired): fail closed with a generic FATAL so we
// don't leak which roles exist. Operators see the underlying
// error in the logs.

// sendReplicationRoleError emits PG's exact wording for a replication-role
// rejection. SQLSTATE 42501 (insufficient_privilege) matches the error
// PostgreSQL raises in walsender setup when the role lacks rolreplication
// and is not a superuser. Returns errAuthRejected on success.
func (c *Conn) sendReplicationRoleError() error { _ = "STUB: not implemented"; return nil }

// authenticateSCRAM performs SCRAM-SHA-256 authentication with the client.
//
// Credentials are fetched once up front via the credential provider; the
// SCRAM hash drives the handshake and the IsReplicationRole flag is cached
// on the connection for the later post-auth gate, so one lookup suffices
// for both. Lookup-time sentinels (login-disabled, expired, missing user)
// are mapped to the matching native-PG error here, before any SASL frames
// are emitted, so we don't reveal which case applied.
//
// Returns the outcome label for mg.gateway.auth.attempts alongside the
// error. The label survives the errAuthRejected wrapping that
// sendAuthError-style helpers apply, so the caller does not have to
// reverse-engineer the cause from the returned error. The SCRAM handshake
// duration (client-first to server-final) is recorded only on paths that
// actually exchange SASL frames; credential-lookup failures exit before
// any frame is emitted and are observed via the separate
// mg.gateway.auth.credential_lookup.duration metric.
func (c *Conn) authenticateSCRAM() (outcome string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// rolcanlogin=false: emit PG's exact wording with SQLSTATE 28000
// (invalid_authorization_specification), matching native PG.

// Expired rolvaliduntil and unknown-user both surface as the opaque
// "password authentication failed" message (28P01), matching PG's
// convention of not disclosing why auth failed.

// Generic credential-lookup failure (pooler unreachable, parse
// error, transport error not carrying a PgDiagnostic). Fail
// closed with the same opaque message PG sends for wrong
// passwords so the client does not learn whether the user
// exists; operators see the underlying cause in the logs.

// Start the SCRAM handshake timer here so the histogram captures only
// the SASL exchange (client-first → server-final), not the upstream
// credential lookup which has its own metric.

// Create the SCRAM authenticator with the pre-fetched hash.

// Tell the authenticator whether we're over TLS regardless of whether
// we manage to compute a cbind hash — the downgrade-attempt detection
// (gs2 flag "y") must fire on any TLS connection.

// Attach channel binding context when the connection is over TLS so the
// authenticator advertises SCRAM-SHA-256-PLUS in addition to SCRAM-SHA-256.
// Plaintext sessions skip this and continue to advertise SCRAM-SHA-256 only.

// Don't fail the connection — log and fall back to plain
// SCRAM-SHA-256 so unusual certs (e.g. unsupported signature
// algorithms) still permit auth, matching PG's permissive
// behavior. The downgrade-detection gate on overTLS still
// fires here.

// Send AuthenticationSASL with supported mechanisms.

// Read SASLInitialResponse (chosen mechanism + client-first-message).

// Process client-first-message and generate server-first-message.
// Pass the username from the startup message as fallback for clients that
// send empty username in SCRAM (like pgx).

// Send AuthenticationSASLContinue with server-first-message.

// Read SASLResponse (contains client-final-message).

// Verify client proof and generate server signature.

// Capture keys for SCRAM passthrough to the backing PostgreSQL.

// Send AuthenticationSASLFinal with server signature. SCRAM ends here;
// the caller (authenticate) continues with the post-auth role-attribute
// check and the AuthenticationOk → ReadyForQuery completion sequence.

// sendAuthenticationSASL sends AuthenticationSASL message with supported mechanisms.
func (c *Conn) sendAuthenticationSASL(mechanisms []string) error {
	_ = "STUB: not implemented"
	// AuthSASL int32
	return nil
}

// Trailing null terminator for the mechanism list.

// sendAuthenticationSASLContinue sends AuthenticationSASLContinue with server data.
func (c *Conn) sendAuthenticationSASLContinue(data string) error {
	_ = "STUB: not implemented"
	return nil
}

// sendAuthenticationSASLFinal sends AuthenticationSASLFinal with server signature.
func (c *Conn) sendAuthenticationSASLFinal(data string) error {
	_ = "STUB: not implemented"
	return nil
}

// readSASLInitialResponse reads SASLInitialResponse from the client.
// Returns the SASL mechanism the client picked and the SASL data
// (client-first-message for SCRAM). The mechanism is validated against the
// set the server advertised — anything else is rejected so a malicious or
// confused client can't downgrade past what was offered.
func (c *Conn) readSASLInitialResponse(advertised map[string]struct{}) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// Read mechanism name.

// Read data length.

// Handle case where client sends no initial data (length = -1).
// This shouldn't happen for SCRAM-SHA-256, but some clients may do this.

// Read SASL data.

// readSASLResponse reads SASLResponse from the client.
// Returns the SASL data (client-final-message for SCRAM).
func (c *Conn) readSASLResponse() (string, error) { _ = "STUB: not implemented"; return "", nil }

// The entire body is the SASL data.

// sendAuthError sends a FATAL authentication-failure response. On a successful
// write it returns errAuthRejected so authenticate() short-circuits the
// post-auth completion sequence; an actual write/flush error is propagated.
func (c *Conn) sendAuthError(message string) error { _ = "STUB: not implemented"; return nil }

// sendLoginDisabledError sends the FATAL error PostgreSQL emits when a role
// with rolcanlogin=false attempts to authenticate (SQLSTATE 28000). The
// message format matches native PG verbatim so libpq-compatible clients
// parse it identically. Returns errAuthRejected on success.
func (c *Conn) sendLoginDisabledError() error { _ = "STUB: not implemented"; return nil }

// mapSCRAMProtocolError converts a SCRAM authenticator error into the exact
// PostgreSQL-style FATAL the client expects to see on the wire.
//
// Returns handled=true iff the error matched a SCRAM protocol class and a
// FATAL has been written. The caller MUST then return ferr up the stack so
// authenticate() short-circuits — ferr is errAuthRejected on success, or the
// underlying write error on failure.
//
// Returns handled=false, nil when the error is not a SCRAM protocol error
// — caller should continue with other classifications.
//
// SQLSTATE + message mapping comes straight from PG17 auth-scram.c so libpq
// and any other PG-compatible client see identical diagnostics.
func (c *Conn) mapSCRAMProtocolError(err error) (handled bool, ferr error) {
	_ = "STUB: not implemented"
	return false, nil
}

// sendScramFatal emits a FATAL ErrorResponse with the supplied SQLSTATE,
// errmsg, and (optional) errdetail. Returns errAuthRejected on success so
// authenticate() short-circuits, matching the sendAuthError pattern.
func (c *Conn) sendScramFatal(sqlState, msg, detail string) error {
	_ = "STUB: not implemented"
	return nil
}

// sendAuthenticationOk sends an AuthenticationOk message to the client.
func (c *Conn) sendAuthenticationOk() error { _ = "STUB: not implemented"; return nil }

// sendBackendKeyData sends the BackendKeyData message.
// This contains the process ID (connection ID) and secret key for query cancellation.
func (c *Conn) sendBackendKeyData() error { _ = "STUB: not implemented"; return nil }

// sendParameterStatuses sends initial ParameterStatus messages to the client.
// These inform the client about server settings.
func (c *Conn) sendParameterStatuses() error {
	_ = "STUB: not implemented"
	// Send standard parameters that clients expect.
	return nil
}

// Pretend to be PostgreSQL 17

// sendParameterStatus sends a single ParameterStatus message.
func (c *Conn) sendParameterStatus(name, value string) error { _ = "STUB: not implemented"; return nil }

// isClientAbortError reports whether a TLS handshake error looks like the
// client closed the connection (network teardown, EOF, or use of an already-
// closed socket) rather than a server-side or crypto failure. Used to split
// the tls.handshake.duration outcome label so the fleet-validation alert
// can suppress noise from clients that hang up mid-handshake.
//
// Timeouts are intentionally excluded: a deadline-exceeded error during
// the TLS handshake is ambiguous between a stalled client and the server's
// own authentication_timeout firing on c.conn. Tagging timeouts as
// handshake_failure (the default) is more accurate than blaming the
// client; operators alerting on a sustained handshake_failure rate will
// still spot stalled-client patterns via the timing distribution.
func isClientAbortError(err error) bool { _ = "STUB: not implemented"; return false }

// sendReadyForQuery sends a ReadyForQuery message to indicate the server is ready.
func (c *Conn) sendReadyForQuery() error { _ = "STUB: not implemented"; return nil }

// Flush to ensure the client receives the message immediately.
