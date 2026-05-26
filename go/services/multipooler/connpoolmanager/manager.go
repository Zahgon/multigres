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

// Package connpoolmanager provides a unified manager for all connection pool types.
package connpoolmanager

import (
	"context"
	"log/slog"
	"sync"
	"sync/atomic"

	"github.com/multigres/multigres/go/common/pgprotocol/client"
	"github.com/multigres/multigres/go/services/multipooler/connstate"
	"github.com/multigres/multigres/go/services/multipooler/pools/admin"
	"github.com/multigres/multigres/go/services/multipooler/pools/connpool"
	"github.com/multigres/multigres/go/services/multipooler/pools/regular"
	"github.com/multigres/multigres/go/services/multipooler/pools/reserved"
)

const (
	// initialUserPoolCapacity is the initial capacity for new user pools.
	// The rebalancer will adjust this based on demand.
	initialUserPoolCapacity int64 = 10
)

// Manager orchestrates per-user connection pools with a shared admin pool.
// Each user gets their own RegularPool and ReservedPool that connect directly
// as that user via trust/peer authentication.
//
// The manager uses an atomic snapshot pattern for lock-free reads on the hot path.
// Most connection requests (for existing users) complete with just an atomic load
// and map lookup. Only new user pool creation requires acquiring a mutex.
//
// Usage:
//
//	cfg := connpoolmanager.NewConfig(reg)
//	cfg.RegisterFlags(cmd.Flags())
//	// ... parse flags ...
//	mgr := cfg.NewManager(logger)
//	mgr.Open(ctx, connConfig)
//	defer mgr.Close()
//
//	// Get connections as needed
//	adminConn, _ := mgr.GetAdminConn(ctx)
//	regularConn, _ := mgr.GetRegularConn(ctx, user)
//	reservedConn, _ := mgr.NewReservedConn(ctx, settings, user)
type Manager struct {
	config     *Config
	logger     *slog.Logger      // Set by Open()
	connConfig *ConnectionConfig // Stored for lazy pool creation

	ctx           context.Context          // Manager lifecycle context, used for lazy pool creation
	adminPool     *admin.Pool              // Shared admin pool for kill operations
	settingsCache *connstate.SettingsCache // Shared settings cache for all users
	metrics       *Metrics                 // OpenTelemetry metrics

	// userPoolsSnapshot holds an atomic pointer to an immutable map of user pools.
	// This enables lock-free reads on the hot path (existing users).
	// The map is replaced atomically via copy-on-write when new users are added.
	userPoolsSnapshot atomic.Pointer[map[string]*UserPool]

	// createMu serializes user pool creation (cold path only).
	// This mutex is only acquired when a new user pool needs to be created.
	createMu sync.Mutex

	// closed indicates whether the manager has been closed.
	closed atomic.Bool

	// generation is incremented on every Open(). Callers that get
	// ErrPoolClosed from an in-flight pool operation can compare the
	// pre-call generation against the current one: if it advanced, a
	// reopen swapped the pools underneath them and the call is safe to
	// retry against the fresh snapshot.
	generation atomic.Uint64

	// Rebalancer goroutine management
	rebalancerCtx    context.Context
	rebalancerCancel context.CancelFunc
	rebalancerWg     sync.WaitGroup

	// Fair share allocators (created once in Open)
	regularAllocator  *FairShareAllocator
	reservedAllocator *FairShareAllocator

	// Drain tracking: counts connections currently lent out across all pools.
	// Used for graceful drain during state transitions (NOT_SERVING).
	drainMu   sync.Mutex
	lentCount int64
	zeroCh    chan struct{}
}

// CredentialQueryRecorder returns a narrow recorder for the gRPC service's
// credential-query observations. Returns nil when the manager is unopened
// or metric init failed; the *Metrics receiver is nil-safe, so callers
// can treat a nil return as the noop sink.
func (m *Manager) CredentialQueryRecorder() CredentialQueryRecorder {
	_ = "STUB: not implemented"
	return *new(CredentialQueryRecorder)
}

// Open initializes the manager and creates the shared admin pool.
// User pools are created lazily on first connection request.
//
// Parameters:
//   - ctx: Context for pool operations
//   - connConfig: Connection settings (socket file, host, port, database)
func (m *Manager) Open(ctx context.Context, connConfig *ConnectionConfig) {
	_ = "STUB: not implemented"
	return
}

// Initialize zeroCh as closed: starts at zero lent connections (drained).

// Build admin client config. pwSourceNone signals that ResolvePgPassword
// was never called — production startup in services/multipooler/init.go
// enforces it before reaching Open, so an unset source here is strictly
// a programmer error.

// Build admin pool config

// Create shared admin pool (used by all user pools for kill operations)

// Create fair share allocators based on global capacity and reserved ratio

// Use configurable minCapacityPerUser as the minimum per-user floor.
// This ensures light users always have enough capacity for burst demand.

// Start the rebalancer goroutine

// Register observable metric callbacks for pool statistics.

// buildClientConfig creates a client.Config with the specified user and password.
// Used by the admin pool and by user pools when SCRAM passthrough is disabled.
//
// SSLMode/TLSConfig from the ConnectionConfig propagate to every dial built by
// this manager. They are honored only on TCP connections (libpq parity); the
// client startup code skips SSLRequest when SocketFile is set.
func (m *Manager) buildClientConfig(user, password string) *client.Config {
	_ = "STUB: not implemented"
	return nil
}

// buildUserClientConfig creates a client.Config for a per-user pool dial.
// When the session carried ClientKey / ServerKey (forwarded via
// ExecuteOptions.UserAuth), the returned config authenticates to PostgreSQL
// using those keys. Absent keys fall back to the empty-password path, which
// only succeeds against a pg_hba.conf that still trusts the local socket for
// the dialing user — the template's narrow admin-user trust exception.
func (m *Manager) buildUserClientConfig(user string, clientKey, serverKey []byte) *client.Config {
	_ = "STUB: not implemented"
	// If SCRAM passthrough keys are present, use them — no password needed.
	// Otherwise we need a password to authenticate. When the requested user
	// matches the configured admin user, fall back to the admin password
	// (this covers internal queries like heartbeat reads that don't carry a
	// SCRAM session). For any other user without keys, return an empty
	// password and let the dial fail with a clear auth error.
	return nil
}

// Invariant violation: production startup in
// services/multipooler/init.go calls
// ResolvePgPassword before any user pool is created.

// getOrCreateUserPool returns the pool for the given user, creating it if needed.
//
// This method uses an atomic snapshot pattern for lock-free reads on the hot path.
// For existing users, this completes with just an atomic load and map lookup.
// Only new user creation acquires the createMu mutex.
//
// clientKey and serverKey are the SCRAM passthrough keys forwarded by the
// gateway on the triggering RPC. They are consumed only when a new pool is
// created (cold path); subsequent RPCs for the same user reuse the existing
// pool regardless of the keys they carry. ClientKey is deterministic per
// (user, password) in PostgreSQL's SCRAM scheme, so cached keys stay valid
// until the user's password rotates; password rotation is surfaced via SCRAM
// auth failure on the next dial.
func (m *Manager) getOrCreateUserPool(user string, clientKey, serverKey []byte) (*UserPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Hot path: atomic load + map lookup (no lock)

// Check if closed before attempting to create

// Cold path: need to create a new user pool.
// Use the manager's lifecycle context so the pool is tied to the manager's
// lifetime, not the caller's request context.

// createUserPoolSlow creates a new user pool. This is the cold path that requires
// acquiring the createMu mutex.
func (m *Manager) createUserPoolSlow(ctx context.Context, user string, clientKey, serverKey []byte) (*UserPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Double-check after acquiring lock

// Check if closed (with lock held)

// Calculate initial capacities proportional to the global split.
// Regular pools get (1 - reservedRatio) of initial capacity, reserved pools get reservedRatio.

// Create drain tracking callbacks for this user pool.
// These are called on every borrow/recycle/reserve/release to track lent connections.

// Create new user pool with per-user pool names for metric cardinality.
// Note: Including username in pool names enables per-user monitoring but increases
// metric cardinality. If this becomes an issue with many users, we can make it configurable.
// Create new user pool with initial capacity. The rebalancer will adjust
// the capacity based on demand within a few seconds.

// Copy-on-write: create new map with the new pool

// Atomic publish

// Close shuts down all connection pools.
func (m *Manager) Close() { _ = "STUB: not implemented"; return }

// Stop the rebalancer goroutine first

// Close all user pools

// Clear the snapshot

// Close shared admin pool last

// Unregister observable metric callbacks so the OTel SDK stops invoking
// them against closed pool state.

// PgUser returns the configured PostgreSQL user for system queries.
func (m *Manager) PgUser() string { _ = "STUB: not implemented"; return "" }

// PgPassword returns the resolved PostgreSQL superuser password and an "ok"
// flag indicating whether ResolvePgPassword ran successfully and produced a
// source. A false ok means no password source was configured (or Resolve was
// never called); production startup in services/multipooler/init.go calls
// Resolve first and surfaces its error, so callers reaching this point can
// treat !ok as a programmer-error invariant violation.
func (m *Manager) PgPassword() (string, bool) { _ = "STUB: not implemented"; return "", false }

// --- Admin Pool Operations ---

// GetAdminConn acquires an admin connection from the shared pool.
// Admin connections are used for control plane operations like killing queries.
// The caller must call Recycle() on the returned connection to return it to the pool.
func (m *Manager) GetAdminConn(ctx context.Context) (admin.PooledConn, error) {
	_ = "STUB: not implemented"
	// Read adminPool under createMu to avoid a nil-pointer panic if Close() is
	// racing with this call. Close() sets m.adminPool = nil while holding
	// createMu, so a snapshot taken here is either the valid pool or nil.
	//
	// We do not use the defer pattern used in other methods because that would
	// mean that we hold the mutex while calling Get() below. If the Get() call
	// block waiting for I/O, no other action will be able to get a connection.
	return *new(admin.PooledConn), nil
}

// --- Regular Pool Operations ---

// GetRegularConn acquires a regular connection for the specified user,
// optionally carrying SCRAM passthrough keys from the caller's session. Keys
// may be nil for admin/internal callers that dial via the local-trust line in
// pg_hba.conf. When non-nil, keys are consumed only when this call triggers
// first-time pool creation for the user; subsequent calls reuse the existing
// pool regardless of the keys they pass. The caller must call Recycle() on
// the returned connection to return it to the pool.
func (m *Manager) GetRegularConn(ctx context.Context, user string, clientKey, serverKey []byte) (regular.PooledConn, error) {
	_ = "STUB: not implemented"
	return *new(regular.PooledConn), nil
}

// GetRegularConnWithSettings is GetRegularConn that additionally applies
// per-session settings. Settings are converted via the shared SettingsCache
// for consistent bucket assignment.
func (m *Manager) GetRegularConnWithSettings(ctx context.Context, settings map[string]string, user string, clientKey, serverKey []byte) (regular.PooledConn, error) {
	_ = "STUB: not implemented"
	return *new(regular.PooledConn), nil
}

// --- Reserved Pool Operations ---

// NewReservedConn creates a new reserved connection for the specified user,
// optionally carrying SCRAM passthrough keys. Settings are converted via the
// shared SettingsCache for consistent bucket assignment. The connection is
// assigned a unique ID for client-side tracking. Optional ReservedConnOption
// values configure validate-with-retry behavior. Key-consumption semantics
// match GetRegularConn. The caller must call Release() when done with the
// connection.
func (m *Manager) NewReservedConn(ctx context.Context, settings map[string]string, user string, clientKey, serverKey []byte, opts ...reserved.ReservedConnOption) (*reserved.Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewLogicalReplicationConn returns a Postgres connection opened in
// replication mode (replication=database) and tagged with
// ReasonLogicalReplication, on the specified user's reserved pool. SCRAM
// passthrough key semantics match NewReservedConn.
func (m *Manager) NewLogicalReplicationConn(ctx context.Context, user string, clientKey, serverKey []byte) (*reserved.Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// evictUserPool removes stale from the snapshot and closes it. Used when a
// pool's cached SCRAM keys no longer match pg_authid — typically because the
// user rotated their PostgreSQL password — so the next getOrCreateUserPool
// call rebuilds from the triggering session's fresh keys.
//
// Eviction is a no-op (returns false) if a racing caller already swapped the
// snapshot to a different pool for this user, or dropped the entry entirely.
// Callers should still retry against whatever the snapshot now holds: the
// racing caller's fresh pool is the right target, and if the entry is gone
// getOrCreateUserPool creates a new one.
func (m *Manager) evictUserPool(user string, stale *UserPool) bool {
	_ = "STUB: not implemented"
	return false
}

// Close outside the createMu critical section would be safer for latency
// but Close is cheap and racing creators already passed the double-check
// by now, so finishing synchronously keeps invariants simple.

// withReopenRetry runs op against the current user pool with two single-shot
// retry paths for transient, self-healable failures:
//
//  1. ErrPoolClosed after a generation bump: reopenConnections() swapped the
//     pools mid-flight (PostgreSQL auto-restart recovery). Retry against the
//     fresh pool. A closed-pool error with no generation bump means the
//     manager is genuinely shutting down — surface it unchanged.
//
//  2. Class-28 SQLSTATE from PostgreSQL: the cached user pool's ClientConfig
//     carries stale SCRAM keys (password rotated in pg_authid). Evict the
//     pool and recreate from the triggering session's keys, which are
//     known-current — they were derived moments ago during the session's
//     SCRAM handshake at MultiGateway against whatever verifier pg_authid
//     holds right now. If the retry also auth-fails (retrier itself used the
//     old password at the gateway), we surface the clean 28xxx error and the
//     client reconnects to re-derive keys against the new verifier.
//
// clientKey and serverKey forward the SCRAM passthrough material to
// getOrCreateUserPool so that a first-time pool creation triggered during
// this call (including after an eviction or reopen swap) gets the session's
// keys.
func withReopenRetry[T any](m *Manager, user string, clientKey, serverKey []byte, op func(*UserPool) (T, error)) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// Manager-restart race: reopenConnections swapped pools mid-flight.

// Stale-key self-heal. evictUserPool is best-effort; even if it returns
// false (racing eviction), getOrCreateUserPool returns whatever is in the
// snapshot now, which is either a freshly-created pool with fresh keys
// or absent (and we create one).

// GetReservedConn retrieves an existing reserved connection by ID for the specified user.
// Returns nil, false if the user pool doesn't exist, the connection is not found, or has timed out.
func (m *Manager) GetReservedConn(connID int64, user string) (*reserved.Conn, bool) {
	_ = "STUB: not implemented"
	// Lock-free read via atomic snapshot
	return nil, false
}

// ApplySettingsToConn ensures the connection's settings match the given session
// settings. ApplySettings handles the diff internally: it resets removed
// variables via individual RESET commands (safe inside transactions, unlike
// RESET ALL) and applies desired variables via SET SESSION.
func (m *Manager) ApplySettingsToConn(ctx context.Context, conn *regular.Conn, settings map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// Pointer equality — same *Settings means same settings (via cache interning)

// --- Stats ---

// Stats returns statistics for all pools.
// This reads from the atomic snapshot, providing a consistent view of all pools.
func (m *Manager) Stats() ManagerStats { _ = "STUB: not implemented"; return *new(ManagerStats) }

// ManagerStats holds statistics for all managed pools.
type ManagerStats struct {
	Admin     connpool.PoolStats       // Shared admin pool stats
	UserPools map[string]UserPoolStats // Per-user pool stats
}

// lentAdd adjusts the lent connection count by n.
// When the count transitions to zero, zeroCh is closed to unblock WaitForDrain.
// When it transitions away from zero, a new zeroCh is created.
func (m *Manager) lentAdd(n int64) { _ = "STUB: not implemented"; return }

// Signal drain complete by closing the channel.

// Already closed, nothing to do.

// Transitioned from 0 to positive: create a new open channel.

// WaitForDrain blocks until all lent connections have been returned or ctx is cancelled.
func (m *Manager) WaitForDrain(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// CloseReservedConnections kills all active reserved connections across all user pools.
// Used after drain grace period expires to prevent reserved connections from being
// used in a non-serving state.
func (m *Manager) CloseReservedConnections(ctx context.Context) int {
	_ = "STUB: not implemented"
	return 0
}

// IsClosed returns whether the manager has been closed.
func (m *Manager) IsClosed() bool { _ = "STUB: not implemented"; return false }

// UserPoolCount returns the number of user pools currently managed.
func (m *Manager) UserPoolCount() int { _ = "STUB: not implemented"; return 0 }

// HasUserPool returns whether a pool exists for the given user.
func (m *Manager) HasUserPool(user string) bool { _ = "STUB: not implemented"; return false }
