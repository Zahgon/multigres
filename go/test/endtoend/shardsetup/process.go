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

// Package shardsetup provides shared test infrastructure for end-to-end tests.
// It sets up the infrastructure for testing a single shard: multipoolers (pgctld + multipooler pairs)
// and optionally multiorch instances.
package shardsetup

import (
	"context"
	"testing"
	"time"

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	"github.com/multigres/multigres/go/provisioner/local"
	"github.com/multigres/multigres/go/tools/executil"
)

// ProcessInstance represents a process instance for testing (pgctld, multipooler, or multiorch).
// This struct is extracted from multipooler/setup_test.go and extended for multiorch support.
type ProcessInstance struct {
	Name        string
	PoolerDir   string // Used by pgctld, multipooler
	ConfigFile  string // Used by pgctld
	LogFile     string
	GrpcPort    int
	PgPort      int    // Used by pgctld
	PgctldAddr  string // Used by multipooler
	EtcdAddr    string // Used by multipooler for topology
	GlobalRoot  string // Topology global root path (used by multipooler, multiorch, multigateway)
	Process     *executil.Cmd
	Binary      string
	Environment []string

	// Multiorch-specific fields
	HttpPort                           int      // HTTP port (used by pgctld and multiorch for health endpoints)
	Cell                               string   // Cell name (used by multipooler and multiorch)
	WatchTargets                       []string // Database/tablegroup/shard targets to watch (multiorch)
	ServiceID                          string   // Service ID (used by multipooler and multiorch)
	LeaderFailoverGracePeriodBase      string   // Grace period base before leader failover (e.g., "0s", "10s")
	LeaderFailoverGracePeriodMaxJitter string   // Max jitter for grace period (e.g., "0s", "5s")

	// Multigateway TLS fields
	TLSCertFile string // TLS certificate file (multigateway)
	TLSKeyFile  string // TLS private key file (multigateway)

	// ReplicaPgPort is the optional PostgreSQL replica-reads listener port (multigateway).
	ReplicaPgPort int

	// ExtraArgs holds additional command-line flags appended to the process args.
	// Used by multigateway for buffer config, etc.
	ExtraArgs []string

	// LogLevel sets --log-level for multipooler/multiorch/multigateway.
	// Defaults to "debug" when empty so existing tests keep verbose logs;
	// benchmarks set "warn" to remove logging from the hot path.
	LogLevel string

	// PgBackRest-specific fields (used by multipooler and pgctld)
	PgBackRestCertPaths *local.PgBackRestCertPaths // pgBackRest TLS certificate paths (multipooler)
	PgBackRestPort      int                        // pgBackRest server port (multipooler, pgctld)
	PgBackRestCertDir   string                     // pgBackRest TLS certificate directory (pgctld)

	// InitdbSQLFiles is a list of SQL files executed after initdb against the
	// target database when InitDataDir runs on this pgctld instance.
	InitdbSQLFiles []string

	// InitdbSQLDirs is a list of role:path entries; each dir's .sql files run
	// under SET SESSION AUTHORIZATION <role> after initdb (pgctld --pg-initdb-sql-dirs).
	InitdbSQLDirs []string

	// PgInitdbExtraConfFiles is a list of postgresql.conf snippets appended to
	// the generated config at init time (pgctld --pg-initdb-extra-conf).
	// Populated by WithMultipoolerPGTLS to enable ssl on the postgres side.
	PgInitdbExtraConfFiles []string

	// PgInitdbArgs is forwarded to pgctld via --pg-initdb-args. Used by the
	// pgregress harness to pass `--no-locale --encoding=SQL_ASCII` so locale-
	// sensitive output (char/varchar sort, to_char 'L' currency, etc.) matches
	// upstream pg_regress fixtures, which initdb the regression cluster in C
	// locale.
	PgInitdbArgs string

	// PgHbaTemplate is an alternate pg_hba.conf template path passed to pgctld
	// via --pg-hba-template. Used by WithMultipoolerPGTLS to relax auth on
	// 127.0.0.1 so the multipooler's per-user pools can dial over TLS without
	// SCRAM passthrough plumbing.
	PgHbaTemplate string

	// SocketFile, PgClientSSLMode, and PgClientSSLRootCert configure the
	// multipooler → postgres dial. SocketFile is set to the default Unix socket
	// path during instance creation; setting it to the empty string forces a
	// TCP dial against the multipooler's hostname so the libpq-style sslmode
	// negotiation actually fires.
	SocketFile          string
	PgClientSSLMode     string
	PgClientSSLRootCert string

	// BackupLocation stores backup configuration from topology (used by pgctld)
	BackupLocation *clustermetadatapb.BackupLocation

	// VpidStampEnabled passes --vpid-stamp-enabled=true to the multipooler so
	// PostgreSQL backends get tagged with `multigres_vpid:<id>` in
	// application_name. Required by the isolation-test harness shim
	// (public.multigres_test_session_is_blocked) to resolve a multigateway
	// virtual PID back to its real backend PID via pg_stat_activity. Default
	// false matches the multipooler's production default; only the pgregress
	// isolation suite flips it on via shardsetup.WithVpidStamping.
	VpidStampEnabled bool
}

// logLevelOrDefault returns p.LogLevel, falling back to "debug" so tests that
// don't opt into a quieter level keep the historical verbose output.
func (p *ProcessInstance) logLevelOrDefault() string { _ = "STUB: not implemented"; return "" }

// multipoolerArgs returns the multipooler command-line arguments derived
// from this ProcessInstance. Extracted from startMultipooler so the arg
// construction is unit-testable without spawning a real process.
//
// p.SocketFile defaults to the standard Unix socket path during instance
// creation; tests that need to exercise the TCP path (e.g. PG TLS) clear
// it before Start to omit --socket-file and force a TCP dial.
func (p *ProcessInstance) multipoolerArgs() []string { _ = "STUB: not implemented"; return nil }

// Required parameter
// Required parameter (MVP only supports "default")
// Required parameter (MVP only supports "0-inf")

// Use the same pooler dir as pgctld

// Allow OnTermSync hooks (notably the graceful-shutdown sequence) to
// run to completion. The default 10s is shorter than the graceful
// shutdown total deadline; without this the hook is cut off mid-flight
// on SIGTERM.

// Start starts the process instance (pgctld, multipooler, multiorch, or multigateway).
// Follows the proven pattern from multipooler/setup_test.go.
func (p *ProcessInstance) Start(ctx context.Context, t *testing.T) error {
	_ = "STUB: not implemented"
	return nil
}

// buildPgctldServerArgs assembles the argv passed to `pgctld server`
// from this instance's configuration. Extracted from startPgctld so the
// flag-forwarding logic (initdb args, extra conf, SQL files/dirs) is
// unit-testable without spawning a real pgctld binary.
func buildPgctldServerArgs(p *ProcessInstance) []string { _ = "STUB: not implemented"; return nil }

// startPgctld starts a pgctld instance (server only, PostgreSQL init/start done separately).
// Copied from multipooler/setup_test.go.
func (p *ProcessInstance) startPgctld(ctx context.Context, t *testing.T) error {
	_ = "STUB: not implemented"
	return nil
}

// Set MULTIGRES_TESTDATA_DIR for directory-deletion triggered cleanup

// startMultipooler starts a multipooler instance.
// Copied from multipooler/setup_test.go.
func (p *ProcessInstance) startMultipooler(ctx context.Context, t *testing.T) error {
	_ = "STUB: not implemented"
	return nil
}

// Start the multipooler server

// Set MULTIGRES_TESTDATA_DIR for directory-deletion triggered cleanup

// startMultiOrch starts a multiorch instance.
// Follows the pattern from multiorch/multiorch_helpers.go:startMultiOrch.
func (p *ProcessInstance) startMultiOrch(ctx context.Context, t *testing.T) error {
	_ = "STUB: not implemented"
	return nil
}

// Add grace period flags if configured (defaults to 0 for fast tests)

// Coverage builds are slower — WAL receiver can take 3-10s to connect.
// So, we Increase the verify-replication timeout to compensate.

// Set up logging like multiorch_helpers.go does

// Start the process with trace context propagation

// Wait for multiorch to be ready (using TCP port check like multiorch_helpers.go)

// startMultigateway starts a multigateway instance.
func (p *ProcessInstance) startMultigateway(ctx context.Context, t *testing.T) error {
	_ = "STUB: not implemented"
	return nil
}

// Add replica port flag if configured

// Add TLS certificate flags if configured

// Append any extra args (e.g., buffer configuration flags)

// Set MULTIGRES_TESTDATA_DIR for directory-deletion triggered cleanup

// Set up logging

// Start the process with trace context propagation

// Wait for multigateway to be ready (Status RPC check)

// startMultiadmin starts a multiadmin instance pointed at the harness's etcd.
// The HTTP port serves both the JSON API used by the Next.js web UI in
// web/multiadmin/ and the gRPC-gateway endpoints; the gRPC port is used by
// the multigres CLI (admin-server flag).
func (p *ProcessInstance) startMultiadmin(ctx context.Context, t *testing.T) error {
	_ = "STUB: not implemented"
	return nil
}

// waitForStartup handles the common startup and waiting logic.
// Copied from multipooler/setup_test.go.
func (p *ProcessInstance) waitForStartup(ctx context.Context, t *testing.T, timeout time.Duration, logInterval int) error {
	_ = "STUB: not implemented"

	// Start the process in background with trace context propagation
	return nil
}

// Give the process a moment to potentially fail immediately

// Check if process died immediately

// Wait for server to be ready

// Check if process died during startup

// Test gRPC connectivity

// If we timed out, try to get process status

// LogRecentOutput logs recent output from the process log file.
// Copied from multipooler/setup_test.go.
func (p *ProcessInstance) LogRecentOutput(t *testing.T, context string) {
	_ = "STUB: not implemented"
	return
}

// IsRunning checks if the process is still running.
// Returns false if the process has exited or was never started.
// Copied from multipooler/setup_test.go.
func (p *ProcessInstance) IsRunning() bool { _ = "STUB: not implemented"; return false }

// ProcessState is set after Wait() returns, meaning process has exited

// Signal 0 checks if process exists without actually sending a signal

// StopPostgres stops PostgreSQL via pgctld gRPC (best effort, no error handling).
// Uses "fast" mode, which takes a checkpoint before stopping.
// Use this to stop postgres before removing data directories for auto-restore tests.
func (p *ProcessInstance) StopPostgres(t *testing.T) { _ = "STUB: not implemented"; return }

// StopPostgresImmediate stops PostgreSQL via pgctld gRPC with "immediate" mode,
// which sends SIGQUIT and skips the pre-shutdown checkpoint. Use this when the
// data directory is about to be wiped anyway, to avoid long graceful-shutdown
// windows that can leave the postgres listen port in TIME_WAIT and block the
// next postgres from binding it on restart.
func (p *ProcessInstance) StopPostgresImmediate(t *testing.T) { _ = "STUB: not implemented"; return }

// stopPostgreSQL stops PostgreSQL via gRPC (best effort, no error handling).
// mode is passed through to pg_ctl stop -m (smart | fast | immediate).
func (p *ProcessInstance) stopPostgreSQL(mode string) { _ = "STUB: not implemented"; return }

// Can't connect, nothing we can do

// pg_ctl stop -m fast takes a checkpoint before stopping, which can
// take several seconds under load. Immediate mode skips the checkpoint.

// TerminateGracefully gracefully terminates a process by first sending SIGTERM,
// waiting for graceful shutdown, and only using SIGKILL if necessary.
// For pgctld, it first stops PostgreSQL via gRPC so that System V shared memory
// segments are released (macOS kern.sysv.shmmni defaults to 32).
func (p *ProcessInstance) TerminateGracefully(logf func(string, ...any), timeout time.Duration) {
	_ = "STUB: not implemented"
	return
}

// For pgctld, stop PostgreSQL first via gRPC. pg_ctl stop -m fast
// releases SysV shared memory segments that would otherwise leak.

// CleanupFunc returns a cleanup function that gracefully terminates the process.
func (p *ProcessInstance) CleanupFunc(logf func(string, ...any)) func() {
	_ = "STUB: not implemented"
	return nil
}

// WaitForPortReady waits for a process to be ready by checking its gRPC port.
// Follows the pattern from multiorch/multiorch_helpers.go:waitForProcessReady.
func WaitForPortReady(t *testing.T, name string, grpcPort int, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// Test gRPC connectivity
