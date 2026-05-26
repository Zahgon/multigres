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

package connpoolmanager

import (
	"crypto/tls"
	"log/slog"
	"time"

	"github.com/spf13/pflag"

	"github.com/multigres/multigres/go/common/pgprotocol/client"
	"github.com/multigres/multigres/go/tools/viperutil"
)

// ConnectionConfig holds connection settings passed from the parent multipooler.
// These are not viper-backed since the flags already exist in the multipooler init.
type ConnectionConfig struct {
	// SocketFile is the full path to the PostgreSQL Unix socket file.
	// If set, Unix socket connection is used instead of TCP.
	// Example: /var/run/postgresql/.s.PGSQL.5432
	SocketFile string

	// Host is the PostgreSQL host (for TCP connections).
	// Ignored if SocketFile is set.
	Host string

	// Port is the PostgreSQL port (for TCP connections).
	// Ignored if SocketFile is set.
	Port int

	// Database is the database name to connect to.
	Database string

	// SSLMode controls libpq-style sslmode for the multipooler → PostgreSQL leg.
	// Only honored on TCP connections (SocketFile == "").
	SSLMode client.SSLMode

	// TLSConfig is the *tls.Config built from SSLMode + sslrootcert. Nil for
	// disable/allow; non-nil otherwise. Only honored on TCP connections.
	TLSConfig *tls.Config
}

type pgPasswordSource int

const (
	pwSourceNone   pgPasswordSource = iota
	pwSourceEnv                     // CONNPOOL_ADMIN_PASSWORD / POSTGRES_PASSWORD env var
	pwSourceOption                  // --connpool-admin-password flag
	pwSourceFile                    // password file path (flag or env var)
)

// Config holds viper-backed configuration values for the connection pool manager.
// Create with NewConfig(), register flags with RegisterFlags(), then create the
// manager with NewManager() when ready.
//
// With per-user connection pools, each user gets their own RegularPool and ReservedPool.
// Connections authenticate directly as the user via trust/peer authentication.
type Config struct {
	// --- PostgreSQL superuser credentials ---
	// Used by the admin pool for kill operations and internal system queries
	// (heartbeat, replication tracking).
	// Configured via POSTGRES_USER / POSTGRES_PASSWORD environment variables.
	pgUser         viperutil.Value[string]
	pgPassword     viperutil.Value[string]
	pgPasswordFile viperutil.Value[string]
	// pgPasswordCached caches the file-resolved password so the hot path
	// (PgPassword()) does not touch disk. Populated by ResolvePgPassword at
	// startup; "" until then.
	pgPasswordCached string
	pgPasswordSource pgPasswordSource // file / option / env / none
	// flagSet is saved during RegisterFlags so ResolvePgPassword can use
	// pflag.Flag.Changed to distinguish "flag explicitly set" from "flag
	// at default value" — viperutil's Get() collapses both into the
	// flag's resolved value. nil when RegisterFlags has not run (e.g. in
	// tests that exercise only the env-var or file paths).
	flagSet *pflag.FlagSet

	// --- PostgreSQL TLS (multipooler → postgres leg) ---
	// libpq-style server verification. Mode + optional CA bundle. Client cert
	// auth (sslcert/sslkey) and CRLs are deferred — see MUL-383.
	pgSSLMode     viperutil.Value[string]
	pgSSLRootCert viperutil.Value[string]

	// --- Pool sizing configuration ---

	// Admin pool configuration (shared across all users)
	adminCapacity viperutil.Value[int64]

	// Per-user regular pool configuration (for simple queries without transactions)
	// Note: Capacity is managed by the rebalancer, not configured here.
	// New pools start with initialUserCapacity (10) and the rebalancer adjusts them.
	userRegularIdleTimeout viperutil.Value[time.Duration]
	userRegularMaxLifetime viperutil.Value[time.Duration]

	// Per-user reserved pool configuration (for transactions)
	// Each user's reserved pool has its own underlying connection pool.
	// Note: Capacity is managed by the rebalancer, not configured here.
	userReservedInactivityTimeout viperutil.Value[time.Duration] // For reserved connections (client inactivity)
	userReservedIdleTimeout       viperutil.Value[time.Duration] // For underlying pool connections
	userReservedMaxLifetime       viperutil.Value[time.Duration]

	// Settings cache size (0 = use default)
	settingsCacheSize viperutil.Value[int64]

	// --- Fair share allocation configuration ---

	// Global capacity is the total number of PostgreSQL connections to manage.
	// This is divided between regular and reserved pools based on reservedRatio.
	globalCapacity viperutil.Value[int64]

	// Reserved ratio is the fraction of global capacity allocated to reserved pools (0.0-1.0).
	// Regular pools get (1 - reservedRatio) of the global capacity.
	reservedRatio viperutil.Value[float64]

	// --- Rebalancer configuration ---

	// Rebalance interval is how often the rebalancer runs to adjust pool capacities.
	rebalanceInterval viperutil.Value[time.Duration]

	// Demand window is the sliding window duration for tracking peak demand.
	// The rebalancer considers peak demand over this window when allocating capacity.
	// Number of buckets = DemandWindow / RebalanceInterval.
	// Example: 30s window with 10s rebalance interval = 3 buckets
	demandWindow viperutil.Value[time.Duration]

	// Inactive timeout is how long a user pool can be inactive before being garbage collected.
	inactiveTimeout viperutil.Value[time.Duration]

	// Minimum capacity per user ensures light users always have enough connections
	// for burst demand that point-in-time sampling might miss.
	minCapacityPerUser viperutil.Value[int64]

	// dialTimeout is the timeout for establishing new PostgreSQL connections.
	// Applied to net.Dialer.Timeout for all pool connections (admin, regular, reserved).
	dialTimeout viperutil.Value[time.Duration]

	// drainGracePeriod is how long to wait for in-flight connections to drain
	// during a NOT_SERVING transition before force-closing reserved connections.
	drainGracePeriod viperutil.Value[time.Duration]
}

// NewConfig creates a new Config with all connection pool settings
// registered to the provided registry.
func NewConfig(reg *viperutil.Registry) *Config {
	_ = "STUB: not implemented"
	// Default values for pool configuration.
	return nil
}

// Per-user regular pool defaults (for simple queries without transactions)

// Per-user reserved pool defaults (for transactions)
// Aggressive - kills reserved connections if client inactive
// Less aggressive - for pool size reduction

// Settings cache size

// Fair share allocation defaults

// Rebalancer defaults

// 30s window / 10s rebalance = 3 buckets

// Fair share allocation - minimum per user
// This ensures light users always have enough capacity for burst demand.
// Set equal to initialUserPoolCapacity (10) so capacity isn't reduced
// below the initial value until there's actual resource pressure.

// Dial timeout for establishing new PostgreSQL connections.

// Drain grace period for NOT_SERVING transitions.

// PostgreSQL superuser credentials (also used for internal system queries)

// PostgreSQL TLS — libpq parity. Default "prefer" mirrors libpq.

// Admin pool (shared across all users)

// Per-user regular pool (for simple queries)

// Per-user reserved pool (for transactions)

// Settings cache size

// Fair share allocation

// Rebalancer

// RegisterFlags registers all connection pool flags with the given FlagSet.
func (c *Config) RegisterFlags(fs *pflag.FlagSet) {
	_ = "STUB: not implemented"
	// Save the FlagSet so ResolvePgPassword can use pflag.Flag.Changed to
	// distinguish "flag explicitly set" from "flag at default value".
	return
}

// PostgreSQL superuser credentials

// PostgreSQL TLS (multipooler → postgres). Mirrors libpq sslmode/sslrootcert.

// Admin pool flags (shared across all users)

// Per-user regular pool flags (for simple queries)

// Per-user reserved pool flags (for transactions)

// Settings cache size flag

// Fair share allocation flags

// Rebalancer flags

// --- Getters for individual values ---

// PgUser returns the configured PostgreSQL superuser name.
// Defaults to POSTGRES_USER environment variable.
func (c *Config) PgUser() string { _ = "STUB: not implemented"; return "" }

// PgPassword returns the resolved PostgreSQL superuser password and the
// source it was loaded from ("file" or "env"). Both are empty until
// ResolvePgPassword has run successfully; the source string is the canonical
// "is the password configured?" check — an empty source means no source
// produced a value (or Resolve was never called). Production startup in
// services/multipooler/init.go calls ResolvePgPassword first and surfaces
// its error, so callers reaching this point can treat an empty source as a
// programmer-error invariant violation.
func (c *Config) PgPassword() (string, pgPasswordSource) {
	_ = "STUB: not implemented"
	return "", *new(pgPasswordSource)
}

// ResolvePgPassword chooses the password source and caches the result so
// subsequent PgPassword() calls return it. Three independent inputs are
// considered, in strict precedence order — once a higher-precedence input is
// "explicitly set" it is authoritative and lower-precedence inputs are NOT
// consulted, even when the higher-precedence value turns out to be empty
// (those empty cases become errors instead of fallthroughs):
//
//  1. File path: --connpool-admin-password-file flag, or
//     CONNPOOL_ADMIN_PASSWORD_FILE / POSTGRES_PASSWORD_FILE env.
//     - Path explicitly empty                  → error.
//     - Path set, file content empty           → error.
//     - Path set, file content non-empty       → use it, source=File.
//  2. Flag option: --connpool-admin-password.
//     - Flag set to empty                      → error.
//     - Flag set to non-empty                  → use it, source=Option.
//  3. Env var: CONNPOOL_ADMIN_PASSWORD or POSTGRES_PASSWORD.
//     - Env set to empty                       → error.
//     - Env set to non-empty                   → use it, source=Env.
//
// Reached the end with no input explicitly set: error "not configured".
// "Explicitly set" is distinct from "viperutil resolved to empty" — viperutil
// collapses unset and empty into the same "" via os.Getenv, so we use
// os.LookupEnv and pflag.Flag.Changed to detect operator intent.
func (c *Config) ResolvePgPassword() error {
	_ = "STUB: not implemented"
	// Row 1, 2a, 2b: file path explicitly set.
	return nil
}

// Row 3, 4: --connpool-admin-password flag explicitly set.

// Row 5, 6: env vars. CONNPOOL_ADMIN_PASSWORD checked before POSTGRES_PASSWORD.

// Row 7: no source configured.

// passwordFileExplicit reports whether the file-path input was explicitly
// set (via flag or env var) and whether the resulting path is the empty
// string. It does NOT consider viperutil defaults — the flag's
// pflag.Flag.Changed bit and os.LookupEnv are the source of truth.
func (c *Config) passwordFileExplicit() (path string, explicit, isEmpty bool) {
	_ = "STUB: not implemented"
	return "", false, false
}

// PgSSLMode parses and returns the configured libpq-style sslmode.
// An invalid value returns the parser error so the caller can fail startup
// rather than silently downgrading to plaintext.
func (c *Config) PgSSLMode() (client.SSLMode, error) {
	_ = "STUB: not implemented"
	return *new(client.SSLMode), nil
}

// PgSSLRootCert returns the configured CA bundle path used to verify the
// PostgreSQL server certificate. Empty unless the operator set it.
func (c *Config) PgSSLRootCert() string { _ = "STUB: not implemented"; return "" }

// ValidatePGSSL checks the libpq-style sslmode + sslrootcert flags at startup
// so a typo or missing CA bundle aborts the multipooler before the connection
// pool manager opens — preventing a silent downgrade to plaintext.
//
// host is the address the multipooler will dial postgres on (only used to
// validate verify-full). Pass an empty host when the multipooler is configured
// for a Unix socket; this function only runs the SSL validation when host is
// non-empty.
func (c *Config) ValidatePGSSL(host string) error { _ = "STUB: not implemented"; return nil }

// AdminCapacity returns the configured admin pool capacity.
func (c *Config) AdminCapacity() int64 { _ = "STUB: not implemented"; return 0 }

// UserRegularIdleTimeout returns the per-user regular pool idle timeout.
func (c *Config) UserRegularIdleTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// UserRegularMaxLifetime returns the per-user regular pool max lifetime.
func (c *Config) UserRegularMaxLifetime() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// UserReservedInactivityTimeout returns the reserved connection inactivity timeout.
// This is how long a reserved connection can be inactive before being killed.
func (c *Config) UserReservedInactivityTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// UserReservedIdleTimeout returns the idle timeout for connections in the reserved pool.
func (c *Config) UserReservedIdleTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// UserReservedMaxLifetime returns the per-user reserved pool max lifetime.
func (c *Config) UserReservedMaxLifetime() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// SettingsCacheSize returns the settings cache size.
func (c *Config) SettingsCacheSize() int { _ = "STUB: not implemented"; return 0 }

// GlobalCapacity returns the total PostgreSQL connections to manage.
// This is divided between regular and reserved pools based on ReservedRatio.
func (c *Config) GlobalCapacity() int64 { _ = "STUB: not implemented"; return 0 }

// ReservedRatio returns the fraction of global capacity allocated to reserved pools (0.0-1.0).
// Regular pools get (1 - reservedRatio) of the global capacity.
func (c *Config) ReservedRatio() float64 { _ = "STUB: not implemented"; return 0 }

// RebalanceInterval returns how often the rebalancer runs to adjust pool capacities.
func (c *Config) RebalanceInterval() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// DemandWindow returns the sliding window duration for peak demand tracking.
// The rebalancer considers peak demand over this window when allocating capacity.
func (c *Config) DemandWindow() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// InactiveTimeout returns how long a user pool can be inactive before garbage collection.
func (c *Config) InactiveTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// MinCapacityPerUser returns the minimum connections per user.
// This ensures light users always have enough capacity for burst demand.
func (c *Config) MinCapacityPerUser() int64 { _ = "STUB: not implemented"; return 0 }

// DialTimeout returns the timeout for establishing new PostgreSQL connections.
func (c *Config) DialTimeout() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

// DrainGracePeriod returns how long to wait for in-flight connections to drain
// during NOT_SERVING transitions before force-closing reserved connections.
func (c *Config) DrainGracePeriod() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// NewManager creates a new connection pool manager from this config.
// Call this after flags have been parsed and when you're ready to create the manager.
// The manager starts in a closed state; call Open() before using it.
func (c *Config) NewManager(logger *slog.Logger) *Manager { _ = "STUB: not implemented"; return nil }

// Manager is closed until Open() is called
