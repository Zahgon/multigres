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

package shardsetup

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"

	"github.com/multigres/multigres/go/tools/executil"

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	multiorchpb "github.com/multigres/multigres/go/pb/multiorch"

	// Register topo plugins
	_ "github.com/multigres/multigres/go/common/plugins/topo"
)

// SetupConfig holds the configuration for creating a ShardSetup.
type SetupConfig struct {
	MultipoolerCount                   int
	MultiOrchCount                     int
	EnableMultigateway                 bool // Enable multigateway (opt-in, default: false)
	EnableMultiadmin                   bool // Enable multiadmin (opt-in, default: false)
	EnableMultigatewayTLS              bool // Enable TLS for multigateway PostgreSQL listener
	EnableMultipoolerPGTLS             bool // Provision postgres with TLS and point multipooler at it via verify-full
	Database                           string
	TableGroup                         string
	Shard                              string
	CellName                           string
	DurabilityPolicy                   string   // Durability policy (e.g., "AT_LEAST_2")
	SkipInitialization                 bool     // Start processes but don't initialize postgres (for bootstrap tests)
	DeferMultipoolerStart              bool     // Start pgctld only; test starts multipooler itself
	LeaderFailoverGracePeriodBase      string   // Grace period base before leader failover (default: "0s" for tests)
	LeaderFailoverGracePeriodMaxJitter string   // Max jitter for grace period (default: "0s" for tests)
	S3BackupBucket                     string   // S3 bucket name (empty = use filesystem)
	S3BackupRegion                     string   // S3 region
	S3BackupEndpoint                   string   // S3 endpoint (empty = use AWS, otherwise s3mock/custom)
	EnableMultigatewayReplicaPort      bool     // Enable replica-reads port on multigateway
	MultigatewayExtraArgs              []string // Extra CLI flags for multigateway (e.g., buffer config)
	OTelCollectorEndpoint              string   // OTLP HTTP endpoint for multigateway span export (empty = disabled)
	EnableMetricsExport                bool     // Enable Prometheus metrics export on all services
	LogLevel                           string   // --log-level for multipooler/multiorch/multigateway (empty = "debug")
	InitdbSQLFiles                     []string // Paths to .sql files executed on each pgctld after initdb against the target database
	InitdbSQLDirs                      []string // role:path entries; each dir's .sql files run under SET SESSION AUTHORIZATION <role> after initdb
	EnableVpidStamping                 bool     // Pass --vpid-stamp-enabled=true to every multipooler (needed by the pgregress isolation harness shim)
	PgInitdbArgs                       string   // Extra args forwarded to pgctld --pg-initdb-args (e.g., "--no-locale --encoding=SQL_ASCII" for pgregress)
	PgInitdbExtraConfFiles             []string // postgresql.conf snippets appended at init time via --pg-initdb-extra-conf (e.g., locale overrides for pgregress)
}

// SetupOption is a function that configures setup creation.
type SetupOption func(*SetupConfig)

// WithMultipoolerCount sets the number of multipooler instances to create.
// Default is 2 (primary + standby).
func WithMultipoolerCount(count int) SetupOption {
	_ = "STUB: not implemented"
	return *new(SetupOption)
}

// WithMultiOrchCount sets the number of multiorch instances to create.
// Default is 0.
func WithMultiOrchCount(count int) SetupOption { _ = "STUB: not implemented"; return *new(SetupOption) }

// WithDatabase sets the database name for the topology.
func WithDatabase(db string) SetupOption { _ = "STUB: not implemented"; return *new(SetupOption) }

// WithCellName sets the cell name for the topology.
func WithCellName(cell string) SetupOption { _ = "STUB: not implemented"; return *new(SetupOption) }

// WithDurabilityPolicy sets the durability policy for the database.
// Default is "AT_LEAST_2".
func WithDurabilityPolicy(policy string) SetupOption {
	_ = "STUB: not implemented"
	return *new(SetupOption)
}

// WithoutInitialization skips postgres initialization and leaves nodes uninitialized.
// Use this for bootstrap tests where multiorch will initialize the shard.
// Processes (pgctld, multipooler) are started but postgres is not initialized.
func WithoutInitialization() SetupOption { _ = "STUB: not implemented"; return *new(SetupOption) }

// WithDeferredMultipoolerStart skips initialization and leaves the multipooler
// unstarted; the test is responsible for starting it.
func WithDeferredMultipoolerStart() SetupOption {
	_ = "STUB: not implemented"
	return *new(SetupOption)
}

// WithMultigateway enables multigateway in the test setup (default: disabled).
// Multigateway will start after shard bootstrap completes.
func WithMultigateway() SetupOption { _ = "STUB: not implemented"; return *new(SetupOption) }

// WithMultiadmin enables multiadmin in the test setup (default: disabled).
// Multiadmin is started after shard bootstrap and exposes HTTP + gRPC APIs
// against the same etcd topology used by the rest of the cluster. The
// Next.js web UI in web/multiadmin/ can be pointed at the resulting HTTP
// port via MULTIADMIN_API_URL=http://localhost:<port> pnpm dev.
func WithMultiadmin() SetupOption { _ = "STUB: not implemented"; return *new(SetupOption) }

// WithMultigatewayReplicaPort enables the replica-reads listener port on multigateway.
// Connections on this port target replicas. Implies WithMultigateway().
func WithMultigatewayReplicaPort() SetupOption { _ = "STUB: not implemented"; return *new(SetupOption) }

// WithMultigatewayTLS enables TLS for the multigateway PostgreSQL listener.
// Implies WithMultigateway(). TLS certificates are auto-generated during setup.
func WithMultigatewayTLS() SetupOption { _ = "STUB: not implemented"; return *new(SetupOption) }

// WithMultigatewayRequireSSL enables TLS for the multigateway PostgreSQL
// listener AND sets --pg-require-ssl=true, so plaintext StartupMessage is
// rejected. Implies WithMultigatewayTLS(). Exercises the hostssl-equivalent
// posture end-to-end.
func WithMultigatewayRequireSSL() SetupOption { _ = "STUB: not implemented"; return *new(SetupOption) }

// WithMultipoolerPGTLS provisions postgres with TLS via pgctld's
// --pg-initdb-extra-conf hook and configures every multipooler in the setup to
// dial postgres over TCP with sslmode=verify-full and the matching CA bundle.
// Used to exercise the multipooler → postgres TLS path end-to-end (MUL-370).
//
// Switching the multipooler from Unix socket to TCP is necessary because
// libpq (and the pgprotocol/client mirror used here) skips the SSLRequest
// negotiation entirely on socket dials.
func WithMultipoolerPGTLS() SetupOption { _ = "STUB: not implemented"; return *new(SetupOption) }

// WithLeaderFailoverGracePeriod sets the grace period configuration for leader failover.
// Default is "0s" for both base and maxJitter to make tests run fast.
// Use this to test grace period behavior explicitly.
func WithLeaderFailoverGracePeriod(base, maxJitter string) SetupOption {
	_ = "STUB: not implemented"
	return *new(SetupOption)
}

// WithS3Backup configures S3-compatible backup storage instead of filesystem.
// The endpoint parameter should be the s3mock endpoint for testing or empty for AWS S3.
// Environment variables must be set:
//   - AWS_ACCESS_KEY_ID
//   - AWS_SECRET_ACCESS_KEY
func WithS3Backup(bucket, region, endpoint string) SetupOption {
	_ = "STUB: not implemented"
	return *new(SetupOption)
}

// WithMultigatewayBuffering enables failover buffering on the multigateway.
// Implies WithMultigateway(). Configures buffer flags for fast test execution:
// short window, small buffer, low drain concurrency, no min-time-between-failovers guard.
func WithMultigatewayBuffering() SetupOption { _ = "STUB: not implemented"; return *new(SetupOption) }

// WithLogLevel sets the --log-level flag for multipooler, multiorch, and multigateway
// processes. Defaults to "debug" so tests retain verbose logs; pass "warn" or "error"
// when log volume itself perturbs the measurement (e.g. benchmarks).
func WithLogLevel(level string) SetupOption { _ = "STUB: not implemented"; return *new(SetupOption) }

// WithOTelExport configures the multigateway to export traces to the given
// OTLP HTTP endpoint. Use with NewTestOTLPCollector to capture spans in tests.
// Implies WithMultigateway().
func WithOTelExport(endpoint string) SetupOption {
	_ = "STUB: not implemented"
	return *new(SetupOption)
}

// WithMetricsExport enables Prometheus metrics export on multipooler and multigateway.
// Each service gets its own Prometheus port, accessible via ShardSetup.MetricsPorts.
// Implies WithMultigateway().
func WithMetricsExport() SetupOption { _ = "STUB: not implemented"; return *new(SetupOption) }

// WithVpidStamping passes --vpid-stamp-enabled=true to every multipooler in
// the setup. Tags each PostgreSQL backend's application_name with
// `multigres_vpid:<id>` so lock-detection probes can map a multigateway
// virtual PID back to its real backend PID via pg_stat_activity. Required by
// the pgregress isolation harness shim and harmless elsewhere, but kept
// opt-in so tests that probe application_name as a generic GUC aren't
// accidentally shadowed.
func WithVpidStamping() SetupOption { _ = "STUB: not implemented"; return *new(SetupOption) }

// WithInitdbSQLFiles forwards the given SQL file paths to every pgctld in the
// shard via --pg-initdb-sql-files. pgctld runs each file against the target
// database after initdb completes (during the InitDataDir RPC triggered by
// shard bootstrap). Files run in the order provided.
func WithInitdbSQLFiles(files ...string) SetupOption {
	_ = "STUB: not implemented"
	return *new(SetupOption)
}

// WithInitdbSQLDirs forwards role:path entries to every pgctld via --pg-initdb-sql-dirs.
// pgctld runs all .sql files in each directory (lexicographic order) under
// SET SESSION AUTHORIZATION <role> after initdb completes.
func WithInitdbSQLDirs(dirs ...string) SetupOption {
	_ = "STUB: not implemented"
	return *new(SetupOption)
}

// WithPgInitdbArgs forwards the given args verbatim to every pgctld via
// --pg-initdb-args. Used by the pgregress harness to invoke initdb with
// `--no-locale --encoding=UTF8`, matching the locale pg_regress uses
// upstream so locale-sensitive output (char/varchar sort order, to_char
// 'L' currency symbol, etc.) reproduces the expected fixtures.
func WithPgInitdbArgs(args string) SetupOption { _ = "STUB: not implemented"; return *new(SetupOption) }

// WithPgInitdbExtraConfFiles appends the given postgresql.conf snippet paths
// to every pgctld via --pg-initdb-extra-conf. Files are concatenated onto the
// generated postgresql.conf at init time; postgres applies last-write-wins so
// settings here override the template defaults. Used by the pgregress harness
// to force `lc_messages/lc_monetary/lc_numeric/lc_time = 'C'` (the template
// otherwise hard-codes en_US.UTF-8, which makes locale-sensitive output
// diverge from upstream `pg_regress --no-locale` expected fixtures).
func WithPgInitdbExtraConfFiles(paths ...string) SetupOption {
	_ = "STUB: not implemented"
	return *new(SetupOption)
}

// SetupTestConfig holds configuration for SetupTest.
type SetupTestConfig struct {
	NoReplication    bool     // Don't configure replication
	PauseReplication bool     // Configure replication but pause WAL replay
	GucsToReset      []string // GUCs to save before test and restore after
}

// SetupTestOption is a function that configures SetupTest behavior.
type SetupTestOption func(*SetupTestConfig)

// WithoutReplication returns an option that actively breaks replication.
// Clears primary_conninfo and synchronous_standby_names, so tests can set up replication from scratch.
func WithoutReplication() SetupTestOption { _ = "STUB: not implemented"; return *new(SetupTestOption) }

// WithPausedReplication returns an option that pauses WAL replay on standbys.
// Replication is already configured from bootstrap; this just pauses WAL application.
// Use this for tests that need to test pg_wal_replay_resume().
func WithPausedReplication() SetupTestOption {
	_ = "STUB: not implemented"
	return *new(SetupTestOption)
}

// WithResetGuc returns an option that saves and restores specific GUC settings.
func WithResetGuc(gucNames ...string) SetupTestOption {
	_ = "STUB: not implemented"
	return *new(SetupTestOption)
}

// multipoolerName returns the name for a multipooler instance by index.
// Uses generic names like "pooler-1", "pooler-2" since multiorch decides which becomes primary.
func multipoolerName(index int) string { _ = "STUB: not implemented"; return "" }

// multiOrchName returns the name for a multiorch instance by index.
func multiOrchName(index int) string { _ = "STUB: not implemented"; return "" }

// NewIsolated creates a new isolated ShardSetup for a single test and returns a cleanup function.
// Use this instead of a shared setup when tests need to kill primaries or perform other
// destructive operations that can't be cleanly restored.
//
// Example:
//
//	setup, cleanup := shardsetup.NewIsolated(t, shardsetup.WithMultipoolerCount(3))
//	defer cleanup()
//	// ... test code that kills primaries, etc.
//
// The cleanup function stops all processes, removes the temp directory, etc.
// If the test failed, it dumps service logs before cleanup to aid debugging.
// Unlike shared setups, this shard is completely isolated and won't affect other tests.
func NewIsolated(t *testing.T, opts ...SetupOption) (*ShardSetup, func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

// New creates a new ShardSetup with the specified configuration.
// This follows the pattern from multipooler/setup_test.go:getSharedTestSetup.
func New(t *testing.T, opts ...SetupOption) *ShardSetup {
	_ = "STUB: not implemented"

	// Get context from testing.T and create root span
	return nil
}

// Default configuration

// primary + standby

// Apply options

// Add configuration attributes to span

// Verify TestMain set up PATH correctly (our binaries should be available)

// Check if PostgreSQL binaries are available

// Create a long-lived context for all processes in this ShardSetup.
// This context is cancelled in Cleanup() to gracefully terminate all processes.
// Derive from context.Background() rather than the span context to avoid premature cancellation.

// Start etcd for topology

// Create topology server and cell

// Create the cell

// Create the database entry in topology with backup_location

// S3/MinIO backend

// Filesystem backend (current behavior)

// Provision postgres-side TLS assets up front so every pgctld + multipooler
// shares the same CA / server cert. Done before the per-pooler loop so the
// snippet path is available when wiring each ProcessInstance.

// Create all multipooler instances (but don't start yet)

// Append SSL config to postgresql.conf at init time.

// Use a permissive pg_hba template that trusts 127.0.0.1 over TLS so
// the multipooler's per-user pools (which dial password="" without
// SCRAM passthrough plumbing) can authenticate over the encrypted
// channel.

// Switch the multipooler off Unix socket onto TCP so SSLRequest is
// actually exchanged, and point it at the same CA the postgres
// server cert was issued from.

// Configure Prometheus metrics export on multipooler if enabled.

// Start all processes (pgctld, multipooler, pgbackrest) for all nodes
// Use setup.ctx for process lifetime, passed ctx only for tracing

// Create multiorch instances (if any requested by the test)

// Start multigateway (if enabled) - MUST be after bootstrap so poolers are in topology

// Generate TLS certificates for multigateway if TLS is enabled

// Allocate ports for multigateway

// Allocate replica port if enabled

// Create multigateway instance (doesn't start it)

// Configure OTel trace export if an endpoint was provided.

// Configure Prometheus metrics export if enabled.

// Start multigateway (waits for Status RPC ready)
// Use setupCtx for process lifetime, passed ctx only for tracing

// Start multiadmin (if enabled). Like multigateway, this is started after
// the multipooler instances exist so it can read them from topology.
// Multiadmin is a passive observer of topology — order vs. bootstrap
// doesn't matter the way it does for multigateway query serving.

// For uninitialized mode (bootstrap tests), we're done - leave nodes uninitialized

// Use multiorch to bootstrap the shard organically

// Verify multigateway can execute queries (if enabled)

// createMultiOrchInstances creates multiorch instances (but doesn't start them).
func (s *ShardSetup) createMultiOrchInstances(t *testing.T, config *SetupConfig) {
	_ = "STUB: not implemented"
	return
}

// StartMultiOrchs starts all multiorch instances.
// Use this for tests that need multiorch running from the get-go.
func (s *ShardSetup) StartMultiOrchs(ctx context.Context, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Register cleanup to ensure recovery is always enabled
// This prevents test failures from leaving recovery disabled
// Capture for closure

// DisableRecovery pauses recovery on the specified multiorch instance.
// Returns a cleanup function that re-enables recovery.
// Recovery is also automatically re-enabled by test cleanup (defense in depth).
func (s *ShardSetup) DisableRecovery(t *testing.T, orchName string) func() {
	_ = "STUB: not implemented"
	return nil
}

// EnableRecovery resumes recovery on the specified multiorch instance.
func (s *ShardSetup) EnableRecovery(t *testing.T, orchName string) {
	_ = "STUB: not implemented"
	return
}

// TriggerRecoveryOnce runs a single immediate recovery cycle and returns any problem codes
// that remain unresolved. Unlike RequireRecovery, it does not keep retrying and does not
// fail the test on problems — the caller decides what to do with the result.
func (s *ShardSetup) TriggerRecoveryOnce(t *testing.T, orchName string, timeout time.Duration) []string {
	_ = "STUB: not implemented"
	return nil
}

// RequireRecovery triggers immediate recovery and blocks until all problems are resolved or
// timeout. Automatically fails the test if any problems remain after timeout.
//
// Logs pooler diagnostics and multiorch status every 5 seconds while waiting, and dumps a
// final cluster state snapshot if recovery times out, to aid flake investigation.
func (s *ShardSetup) RequireRecovery(t *testing.T, orchName string, timeout time.Duration) {
	_ = "STUB: not implemented"
	return
}

// Log cluster state every 5 seconds while the RPC is in flight.

// WaitForHealthStreamsEstablished blocks until the named multiorch instance
// reports `Reachable=true` for every pooler in this shard, indicating it has
// received at least one snapshot from each pooler over the ManagerHealthStream
// — i.e. the stream is dialled, handshaked, and exchanging data.
//
// Tests should call this after StartMultiOrchs (and RequireRecovery, if used)
// but before any test action that depends on the orchestrator observing a
// pooler-side event (notably SIGTERM-triggered failover, which only fires
// quickly when the orchestrator is already subscribed to receive the
// REQUESTING_DEMOTION snapshot). RequireRecovery's "no problems" condition
// can be satisfied by topology state alone — before any health stream is up
// — and that gap is the source of the stream-establishment flake under
// CPU-overhead conditions (subprocess-coverage CI, busy runners).
func (s *ShardSetup) WaitForHealthStreamsEstablished(t *testing.T, orchName string, timeout time.Duration) {
	_ = "STUB: not implemented"
	return
}

// connectToMultiOrch creates a gRPC client connection to the named multiorch instance.
// Fails the test if the instance is not found or the connection cannot be established.
func (s *ShardSetup) connectToMultiOrch(t *testing.T, orchName string) *grpc.ClientConn {
	_ = "STUB: not implemented"
	return nil
}

// ensureRecoveryEnabled makes a best-effort attempt to enable recovery.
// Used in test cleanup to prevent disabled recovery from affecting subsequent tests.
func ensureRecoveryEnabled(t *testing.T, mo *ProcessInstance) { _ = "STUB: not implemented"; return }

// initializeWithMultiOrch uses multiorch to bootstrap the shard organically.
// It starts a single multiorch (temporary if none configured), waits for it to
// initialize the shard, then stops it (clean state = multiorch not running).
func initializeWithMultiOrch(ctx context.Context, t *testing.T, setup *ShardSetup, config *SetupConfig) {
	_ = "STUB: not implemented"
	return
}

// Use existing multiorch or create a temporary one

// Use the first multiorch instance

// Create a temporary multiorch for initialization

// Start multiorch

// Wait for multiorch to bootstrap the shard (elect a primary)

// This before we return the cleanup function, so let's dump the logs if we
// fail to bootstrap the shard

// Stop multiorch (clean state = multiorch not running)

// Remove temporary multiorch from the map

// Save the current GUC values as the baseline "clean state".
// After bootstrap, replication is configured, so the baseline includes:
// - Primary: synchronous_standby_names with standby list, synchronous_commit=on
// - Replicas: primary_conninfo pointing to primary
// ValidateCleanState and cleanup will restore to these values.

// waitForShardBootstrap waits for multiorch to bootstrap the shard by electing a primary
// and initializing all standbys. Returns the name of the elected primary or an error.
func waitForShardBootstrap(ctx context.Context, t *testing.T, setup *ShardSetup) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// checkBootstrapStatus checks if all nodes are initialized and returns the primary name.
// A node is considered initialized only if it can be queried AND has an explicit type (PRIMARY or REPLICA).
// Additionally checks that:
// - PRIMARY has sync replication configured with the full cohort and all replicas connected
// - REPLICA has primary_conn_info configured
func checkBootstrapStatus(ctx context.Context, t *testing.T, setup *ShardSetup) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// Build the set of all multipooler names for exact membership checks.

// Build human-readable status for each pooler

// Try both the Status RPC and a postgres query independently — either can
// succeed without the other, and we want to report as much as possible.

// --- Manager Status RPC (works even when postgres is down) ---

// --- Postgres query ---

// Build a status string from everything we know.
// Start with queryability, then layer in type/replication/action details.

// Format active action suffix (present in every line when an action is running).

// Verify the sync standby list contains every multipooler in the cohort.

// Verify connected_followers contains every replica (all names except this primary).

// Check that primary_conn_info is configured

// UNKNOWN type means not fully initialized yet - don't count

// Get latest backup ID from primary

// Set summary attributes and detailed pooler statuses

// Query multiorch instances for status (best-effort diagnostic logging)

// startMultipoolerInstances starts pgctld and multipooler processes without initializing postgres.
// Use this for bootstrap tests where multiorch will initialize the shard.
//
// TODO: Consider parallelizing Start() calls using a WaitGroup for faster startup.
// Currently processes are started sequentially which adds latency.
func startMultipoolerInstances(ctx context.Context, t *testing.T, instances []*MultipoolerInstance, deferMultipoolerStart bool) {
	_ = "STUB: not implemented"
	return
}

// Create child span for each instance

// Start pgctld (postgres will be initialized later, or by multiorch for bootstrap)

// Start multipooler

// Wait for multipooler to be ready

// startEtcd starts etcd without registering t.Cleanup() handlers
// since cleanup is handled manually by TestMain via Cleanup().
// Follows the pattern from multipooler/setup_test.go:startEtcdForSharedSetup.
func startEtcd(ctx context.Context, t *testing.T, dataDir string) (string, *executil.Cmd, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Check if etcd is available in PATH

// Get ports for etcd (client, peer, and metrics)

// Wrap etcd with run_in_test.sh for orphan protection. Stops gracefully when
// runningCtx is cancelled so run_in_test.sh can terminate etcd cleanly.

// Set MULTIGRES_TESTDATA_DIR for directory-deletion triggered cleanup

// Stop the etcd process if it's not ready

// ValidateCleanState checks that all multipoolers are in the expected clean state.
// Clean state is defined by the baseline GUCs captured after bootstrap:
//   - Primary: not in recovery, GUCs match baseline, type=PRIMARY
//   - Standbys: in recovery, GUCs match baseline, wal_replay not paused, type=REPLICA
//   - MultiOrch: NOT running (multiorch starts in SetupTest and stops in cleanup)
//
// Note: Term is NOT validated. It can increase across tests and there's no safe
// way to reset it. Tests should work with whatever term they start with.
//
// Returns an error if state is not clean.
func (s *ShardSetup) ValidateCleanState() error { _ = "STUB: not implemented"; return nil }

// Require primary to be set (happens after bootstrap)

// Verify multiorch instances are NOT running (clean state = no orchestration)

// Check recovery mode

// Validate pooler type is PRIMARY

// Verify WAL replay not paused

// Validate pooler type is REPLICA

// Validate GUCs match baseline values

// Note: We intentionally don't validate term here.
// Term can increase across tests (e.g., when BeginTerm is called) and
// there's no safe way to reset it without an RPC. Tests should work with
// whatever term they start with and use relative term values.

// ResetToCleanState resets all multipoolers to the baseline clean state.
// This restores GUCs to baseline values, pooler types to PRIMARY/REPLICA,
// resumes WAL replay, and stops multiorch instances.
// Note: Term is NOT reset. It can only increase and tests should handle any starting term.
func (s *ShardSetup) ResetToCleanState(t *testing.T) { _ = "STUB: not implemented"; return }

// Stop multiorch instances first (clean state = not running)

// Check if primary was demoted and restore if needed

// Restore GUCs to baseline values

// Resume WAL replay if paused (for standbys)

// Note: We don't reset term here. Term can only increase and there's no
// safe way to reset it without an RPC. Tests should handle any starting term.

// ReinitializeCluster tears down the running cluster and brings up a fresh one.
// It stops all processes (multigateway, multipooler, pgctld), removes PostgreSQL
// data directories, restarts everything, and re-bootstraps via multiorch.
// Use this between independent test suites that may leave the cluster in a
// degraded state (e.g., PostgreSQL regression tests that crash connections).
func (s *ShardSetup) ReinitializeCluster(t *testing.T) { _ = "STUB: not implemented"; return }

// 1. Stop multigateway (routes to multipoolers, stop first)

// 2. Stop multiorch instances

// 3. Stop multipooler + pgctld and remove PostgreSQL data.
// StopPostgres must be called BEFORE killing pgctld, otherwise
// the postgres process survives and holds the port.
//
// Use immediate shutdown (SIGQUIT, no checkpoint): the data dir is
// about to be wiped, so clean-shutdown state is irrelevant. Fast mode's
// pre-shutdown checkpoint can take >10s on the primary (especially under
// replication load), triggering a graceful-period SIGKILL that leaves
// postgres's listen port in TIME_WAIT on Linux for ~60s. The subsequent
// postgres restart then cannot bind its port and retries until
// WaitForManagerReady times out.

// Remove ALL contents of the data directory so pgctld starts
// completely fresh. This clears pg_data (PostgreSQL data),
// pg_sockets (stale Unix sockets), pgbackrest (backup state),
// and any other state files.

// 3b. Clear the shared backup repository so pgbackrest doesn't
// reference stale backups from the previous cluster.

// 3c. Clear stale topology state. Without this, a pooler that was elected
// PRIMARY in the previous suite reads its stale role from etcd on restart,
// finds an empty data directory (wiped in step 3), and hangs in recovery —
// never reaching PostgresReady. The restart loop below then times out on
// WaitForManagerReady for that pooler.
//
// We wipe:
//   - databases/<db>/<tablegroup>/*   (shard records + ShardInitClaim)
//   - <cell>/poolers|gateways|orchs/* (per-process registration state)
// We keep:
//   - databases/<db>/Database         (preserves BackupLocation + DurabilityPolicy)
//   - cells/<cell>/Cell               (cell config multipoolers need to reconnect)

// 4. Start pgctld + multipooler for all nodes

// 5. Start multigateway

// 6. Bootstrap via temporary multiorch

// 7. Wait for multigateway to serve queries

// wipeTopologyForReinit removes the etcd keys that would otherwise carry
// stale shard-election and per-process registration state across a
// ReinitializeCluster call. See the call site in ReinitializeCluster for
// the full rationale.
//
// It connects to etcd directly (rather than via TopoServer) because the
// public topoclient API only exposes per-record delete helpers and this
// needs a recursive prefix delete.
func (s *ShardSetup) wipeTopologyForReinit(t *testing.T, database, tableGroup string) {
	_ = "STUB: not implemented"
	return
}

// topoclient test root is "/multigres". Global topo is under /multigres/global,
// and each cell is under /multigres/<cell>/.

// SetupTest provides test isolation by validating clean state and automatically
// restoring baseline state at test cleanup.
//
// DEFAULT BEHAVIOR (no options):
//   - Validates clean state before test (GUCs match baseline from bootstrap)
//   - Replication is already configured from bootstrap
//   - Registers cleanup to restore baseline GUCs and reset state after test
//
// WithoutReplication():
//   - Actively breaks replication: clears primary_conninfo and synchronous_standby_names
//   - Use for tests that need to set up replication from scratch
//   - Cleanup restores baseline (re-enables replication)
//
// WithPausedReplication():
//   - Pauses WAL replay on standbys (replication already configured)
//   - Use for tests that need to test pg_wal_replay_resume()
//
// WithResetGuc(gucNames...):
//   - Adds additional GUCs to save/restore beyond baseline
//
// Follows the pattern from multipooler/setup_test.go:setupPoolerTest.
func (s *ShardSetup) SetupTest(t *testing.T, opts ...SetupTestOption) {
	_ = "STUB: not implemented"
	return
}

// Fail fast if shared processes died

// Validate that settings are in the expected clean state (GUCs match baseline)

// If WithoutReplication is set, actively break replication

// If WithPausedReplication is set, pause WAL replay on standbys

// Start multiorch instances
// TODO (@rafa): once we have a way to disable multiorch on a shard, we don't need
// this big hammer of stopping / starting on each test.

// Register cleanup handler to restore to baseline state.
// Note: Processes are still running during t.Cleanup() - they're only stopped later
// by ShardSetup.Cleanup() which cancels s.ctx.

// Stop multiorch instances first (clean state = multiorch not running)
// Use explicit termination here since multiorch should be stopped before restoring state.

// Check if primary was demoted and restore if needed

// Restore GUCs to baseline values

// Always resume WAL replay (must be after GUC restoration)
// This ensures we leave the system in a good state even if tests paused replay.

// Note: We don't reset term here. Term can only increase and there's no
// safe way to reset it without an RPC. Tests should handle any starting term.

// Validate cleanup worked.
// Use a generous timeout: GUC values written by RestoreGUCs are already
// waited on inside that function, so this is a final sanity check that
// should pass quickly. The extra headroom guards against slow CI runners.

// breakReplication clears replication configuration on all nodes.
// Use this for tests that need to set up replication from scratch.
func (s *ShardSetup) breakReplication(t *testing.T, ctx context.Context) {
	_ = "STUB: not implemented"

	// Clear synchronous_standby_names on primary
	return
}

// Clear primary_conninfo on standbys and wait for WAL receiver to stop

// Wait for the WAL receiver to stop. ReloadConfig has already confirmed
// primary_conninfo is cleared; this waits for postgres to act on it.

// Also verify WAL receiver has stopped

// pauseReplicationOnStandbys pauses WAL replay on all standbys.
func (s *ShardSetup) pauseReplicationOnStandbys(t *testing.T, ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

// NewClient returns a new MultipoolerClient for the specified multipooler instance.
// The caller is responsible for closing the client.
func (s *ShardSetup) NewClient(t *testing.T, name string) *MultipoolerClient {
	_ = "STUB: not implemented"
	return nil
}

// unreachable, but needed for linter

// NewPrimaryClient returns a new MultipoolerClient for the primary instance.
// The caller is responsible for closing the client.
func (s *ShardSetup) NewPrimaryClient(t *testing.T) *MultipoolerClient {
	_ = "STUB: not implemented"
	return nil
}

// makeMultipoolerID creates a multipooler ID for testing.
func makeMultipoolerID(cell, name string) *clustermetadatapb.ID {
	_ = "STUB: not implemented"
	return nil
}

// GetMultipoolerID returns the multipooler ID for the named instance.
func (s *ShardSetup) GetMultipoolerID(name string) *clustermetadatapb.ID {
	_ = "STUB: not implemented"
	return nil
}

// KillPostgres kills postgres using SIGKILL on a node (simulates hard database crash).
// This sends SIGKILL directly to the postgres process without clean shutdown.
// The multipooler stays running to report the unhealthy status to multiorch.
func (s *ShardSetup) KillPostgres(t *testing.T, name string) { _ = "STUB: not implemented"; return }

// unreachable, but needed for linter

// Read the PID from postmaster.pid file

// The first line of postmaster.pid contains the PID

// Send SIGKILL to the postgres process

// StopPostgres disables automatic postgres restarts on the named node, then stops postgres
// via the pgctld Stop RPC with the given mode (e.g. "fast", "immediate").
// It returns a resume function that re-enables restarts; the caller should defer it.
//
// This is safer than calling pgctld Stop directly because the postgres monitor runs
// continuously and would otherwise restart postgres immediately after it stops.
func (s *ShardSetup) StopPostgres(t *testing.T, name, mode string) (resume func()) {
	_ = "STUB: not implemented"
	return nil
}

// Disable automatic restarts so the monitor does not restart postgres before we stop it.

// Stop postgres via pgctld.

// ShutdownPostgres gracefully shuts down postgres on the specified node using pgctld Stop RPC.
// This is different from KillPostgres which uses SIGKILL for immediate termination.
// Use this to test scenarios where postgres shuts down cleanly vs crash scenarios.
func (s *ShardSetup) ShutdownPostgres(t *testing.T, name string) (resume func()) {
	_ = "STUB: not implemented"
	return nil
}

// baselineGucNames returns the GUC names to save/restore for baseline state.
var baselineGucNames = []string{
	"synchronous_standby_names",
	"synchronous_commit",
	"primary_conninfo",
}

// saveBaselineGucs captures the current GUC values from all nodes as the baseline "clean state".
// This is called after bootstrap completes, so the baseline includes replication configuration.
func (s *ShardSetup) saveBaselineGucs(t *testing.T) { _ = "STUB: not implemented"; return }

// logMultiOrchStatus queries each running multiorch and logs its view of the shard.
// This includes pooler states and detected problems.
// Best-effort diagnostic logging - failures are logged but do not fail the test.
func logMultiOrchStatus(ctx context.Context, t *testing.T, setup *ShardSetup, label string) {
	_ = "STUB: not implemented"
	return
}

// Query shard status for the default shard
// TODO: Handle multiple shards if needed

// "default" - must match multipooler registration
// "0-inf" - must match multipooler registration

// Format pooler health status

// Format problems

// missingNames returns the sorted list of keys present in expected but absent from actual.
func missingNames(expected, actual map[string]struct{}) []string {
	_ = "STUB: not implemented"
	return nil
}

// formatProblemsCompact creates a one-line summary: [code1@pooler1, code2@pooler2]
func formatProblemsCompact(problems []*multiorchpb.DetectedProblem) string {
	_ = "STUB: not implemented"
	return ""
}

// formatPoolerHealth creates a detailed status: 3/3 reachable (pooler-1:PRIMARY/up, pooler-2:REPLICA/up, pooler-3:REPLICA/up)
func formatPoolerHealth(healthList []*multiorchpb.PoolerHealth) string {
	_ = "STUB: not implemented"
	return ""
}

// Count reachable poolers

// Build individual pooler status strings

// Format as: pooler-1:PRIMARY/up or pooler-1:UNKNOWN/down
