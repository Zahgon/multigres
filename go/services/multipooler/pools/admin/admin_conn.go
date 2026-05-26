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

// Package admin provides administrative connection management for kill operations.
package admin

import (
	"context"
	"errors"
	"time"

	"github.com/multigres/multigres/go/common/pgprotocol/client"
	"github.com/multigres/multigres/go/common/sqltypes"
	"github.com/multigres/multigres/go/services/multipooler/connstate"
	"github.com/multigres/multigres/go/services/multipooler/pools/connpool"
)

// maxQueryAttempts is the maximum number of attempts for retrying queries.
const maxQueryAttempts = 3

// retryBackoff is the delay between retry attempts. This gives PostgreSQL
// time to finish starting up when the connection error is due to a restart.
const retryBackoff = 100 * time.Millisecond

// DefaultCancelTimeout is the default timeout for cancel/terminate operations.
// This is used when cancelling a backend query due to context cancellation.
const DefaultCancelTimeout = 5 * time.Second

// ErrUserNotFound is returned when a requested PostgreSQL role doesn't exist in pg_authid.
var ErrUserNotFound = errors.New("user not found in pg_authid")

// ErrLoginDisabled is returned when the role exists in pg_authid but has
// rolcanlogin=false. PostgreSQL rejects such logins at ClientAuthentication
// regardless of auth method with SQLSTATE 28000.
var ErrLoginDisabled = errors.New("role is not permitted to log in")

// ErrPasswordExpired is returned when the role's rolvaliduntil is set and
// has elapsed. PostgreSQL treats this as a password-authentication failure
// (SQLSTATE 28P01), indistinguishable from a wrong password at the wire.
var ErrPasswordExpired = errors.New("role password has expired")

// Conn wraps a client.Conn for administrative operations.
// It implements connpool.Connection with no settings support.
//
// AdminConn provides the ability to terminate other backend connections using
// pg_terminate_backend() and pg_cancel_backend().
type Conn struct {
	// conn is the underlying PostgreSQL connection.
	conn *client.Conn
}

// NewConn creates a new AdminConn wrapping the given client connection.
func NewConn(conn *client.Conn) *Conn { _ = "STUB: not implemented"; return nil }

// --- connpool.Connection interface ---

// Settings returns nil because admin connections don't use settings-based routing.
func (c *Conn) Settings() *connstate.Settings {
	_ = "STUB: not implemented"

	// IsClosed returns true if the connection has been closed.
	return nil
}

func (c *Conn) IsClosed() bool { _ = "STUB: not implemented"; return false }

// Close closes the underlying connection.
func (c *Conn) Close() error { _ = "STUB: not implemented"; return nil }

// ApplySettings panics because admin connections don't support settings.
// This should never be called - admin connections are always "clean".
func (c *Conn) ApplySettings(_ context.Context, _ *connstate.Settings) error {
	_ = "STUB: not implemented"
	return nil
}

// ResetAllSettings is a no-op because admin connections don't have settings.
func (c *Conn) ResetAllSettings(_ context.Context) error {
	_ = "STUB: not implemented"

	// --- Admin operations ---
	return nil
}

// RolAuthInfo captures the pg_authid columns multigateway needs to make
// authentication decisions for a given role.
type RolAuthInfo struct {
	// ScramHash is the stored SCRAM-SHA-256 password hash, or empty if the
	// role has no password set.
	ScramHash string

	// IsReplicationRole mirrors pg_authid.rolreplication. PostgreSQL requires
	// this attribute (or rolsuper) for any connection started with the
	// replication=true / replication=database startup parameter.
	IsReplicationRole bool
}

// GetRolAuthInfo retrieves the auth-relevant pg_authid columns for a given username:
// the SCRAM password hash plus rolreplication.
//
// In addition to returning the data, this enforces two native-PG login checks
// so multigateway reaches parity with PostgreSQL's ClientAuthentication:
//   - rolcanlogin=false → ErrLoginDisabled (SQLSTATE 28000 at the gateway)
//   - rolvaliduntil in the past → ErrPasswordExpired (SQLSTATE 28P01)
//
// These checks run at the admin layer because the per-user downstream pool
// connection (local Unix socket, trust auth in pg_hba.conf) bypasses PG's
// own password-based enforcement of rolvaliduntil, so without filtering here
// an expired password would authenticate successfully through multigateway.
//
// Precedence matches PostgreSQL (src/backend/libpq/auth.c CheckPasswordAuth):
// login-disabled is checked before password validity.
func (c *Conn) GetRolAuthInfo(ctx context.Context, username string) (*RolAuthInfo, error) {
	_ = "STUB: not implemented"
	// Query pg_authid for the password hash, the two login-time predicates PG
	// itself evaluates, and rolreplication. Expressing validity in SQL avoids
	// timestamp parsing and matches PG's own `now()` semantics (both are
	// evaluated server-side).
	//
	// Note: PG's own has_rolreplication() short-circuits to true for
	// superusers regardless of rolreplication (miscinit.c:734-750). We do
	// not include rolsuper in the SELECT because the gateway only consults
	// IsReplicationRole after SCRAM has already authenticated the role —
	// reaching that check at all means the role passed login-disabled and
	// password-validity predicates, and any superuser intended for
	// replication is expected to also have rolreplication=true. If a
	// superuser-only-without-rolreplication ever needs walsender access,
	// extend RolAuthInfo with IsSuperuser and OR it in at the gateway.
	return nil, nil
}

// Check if user exists.

// rolcanlogin takes precedence — PG rejects NOLOGIN roles before looking at
// the password, so the error class differs (28000 vs 28P01).

// rolvaliduntil: a non-NULL value in the past invalidates the password for
// password-based auth. PG returns the same opaque "password authentication
// failed" message as wrong-password in this case.

// Extract the password hash (may be NULL/empty if no password set).

// parsePgBool converts PG's textual boolean representation ("t"/"f") to Go's bool.
// NULL is treated as false defensively — neither rolcanlogin nor the password_valid
// expression should ever return NULL, but we don't want a malformed row to appear
// as "allowed to log in".
func parsePgBool(v sqltypes.Value) bool { _ = "STUB: not implemented"; return false }

// QueryWithRetry executes a query with automatic retry and reconnection on
// connection error. Safe for stateless internal queries (heartbeat, replication tracking).
func (c *Conn) QueryWithRetry(ctx context.Context, sql string) ([]*sqltypes.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// QueryArgsWithRetry executes a parameterized query with automatic retry and
// reconnection on connection error.
func (c *Conn) QueryArgsWithRetry(ctx context.Context, sql string, args ...any) ([]*sqltypes.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// queryWithRetry executes a query with automatic retry on connection error.
// If the first attempt fails with a connection error, it reconnects and retries
// up to maxQueryAttempts total. This handles stale connections that occur when
// PostgreSQL restarts while the pool holds old socket FDs.
func (c *Conn) queryWithRetry(ctx context.Context, sql string) ([]*sqltypes.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Brief backoff before reconnecting to give PostgreSQL time to
// finish starting up if the error is due to a restart.

// Admin connections are bare superuser connections with no session
// state to re-apply (Settings() returns nil, ApplySettings panics).
// A raw client.Conn.Reconnect is sufficient here.

// execQueryWithContextCancel executes a query operation in a goroutine so that
// context cancellation can interrupt a blocking network read.
//
// When the context is cancelled while a query is blocked waiting for PostgreSQL
// (e.g. an INSERT waiting for synchronous standby acknowledgement), the underlying
// connection is force-closed to unblock the goroutine's read. The goroutine is
// waited on before returning to prevent goroutine leaks.
//
// ForceClose is used instead of Close to avoid a concurrent write race: the
// goroutine may hold bufmu while reading from the buffered reader, and Close
// would also write to the same buffered writer.
//
// After a force-close, IsClosed() returns true. The caller's Recycle() will
// then signal the pool to replace the connection rather than returning it.
func execQueryWithContextCancel(ctx context.Context, conn *client.Conn, op func() ([]*sqltypes.Result, error)) ([]*sqltypes.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Force-close the connection to unblock the goroutine's blocking network read.

// Wait for the goroutine to finish (it returns quickly after ForceClose).

// TerminateBackend terminates a backend process using pg_terminate_backend().
// Returns true if the backend was terminated, false if it was not found or
// the caller lacks permission.
//
// If the context is cancelled while waiting for the query, the connection is
// closed to abort the operation, and the context error is returned.
func (c *Conn) TerminateBackend(ctx context.Context, processID uint32) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// CancelBackend cancels the current query on a backend process using pg_cancel_backend().
// This sends SIGINT to the backend, canceling the current query but keeping the connection.
// Returns true if the signal was sent, false if the backend was not found or
// the caller lacks permission.
//
// If the context is cancelled while waiting for the query, the connection is
// closed to abort the operation, and the context error is returned.
func (c *Conn) CancelBackend(ctx context.Context, processID uint32) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// queryResult holds the result of a query execution.
type queryResult struct {
	success bool
	err     error
}

// execBackendFunc executes a pg_*_backend() function with context cancellation support.
// If the context is cancelled while the query is running, the connection is closed
// to abort the operation. The underlying query uses queryWithRetry for automatic
// reconnection on stale connections.
func (c *Conn) execBackendFunc(ctx context.Context, sql, operation string, processID uint32) (bool, error) {
	_ = "STUB: not implemented"
	// Run query in goroutine so we can respect context cancellation.
	return false, nil
}

// pg_*_backend returns a boolean indicating success.

// The result is 't' for true, 'f' for false.

// Context cancelled - close the connection to abort the query.
// This ensures we don't leave a hung query on the server.

// Ensure Conn implements connpool.Connection.
var _ connpool.Connection = (*Conn)(nil)
