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

	_ "github.com/lib/pq"

	"github.com/multigres/multigres/go/common/topoclient"
	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	"github.com/multigres/multigres/go/provisioner/local"
	"github.com/multigres/multigres/go/tools/executil"
)

const (
	// DefaultTestUser is the PostgreSQL user that tests use when connecting to
	// multigateway or multipooler as a regular client.
	DefaultTestUser = "postgres"
)

// MultipoolerInstance represents a multipooler instance, which is a pair of pgctld + multipooler processes.
// In multigres, a multipooler always has both: pgctld manages PostgreSQL, multipooler handles pooling.
type MultipoolerInstance struct {
	Name        string
	Pgctld      *ProcessInstance
	Multipooler *ProcessInstance
}

// ShardSetup holds shared test infrastructure for a single shard.
// MultipoolerInstances are stored in a map by name for flexible access.
type ShardSetup struct {
	TempDir        string
	TempDirCleanup func()
	EtcdClientAddr string
	EtcdCmd        *executil.Cmd
	TopoServer     topoclient.Store
	CellName       string

	// Context for all processes started by this ShardSetup.
	// Cancelled when Cleanup() is called to gracefully terminate all processes.
	runningCtx context.Context
	cancel     context.CancelFunc

	// MultipoolerInstances indexed by name (e.g., "pooler-1", "pooler-2", "pooler-3")
	Multipoolers map[string]*MultipoolerInstance

	// PrimaryName is the name of the node elected as primary after bootstrap.
	// Set by initializeWithMultiOrch. Use GetPrimary() to access.
	PrimaryName string

	// Multiorch instances (can have multiple)
	MultiOrchInstances map[string]*ProcessInstance

	// Multigateway instance (optional, enabled via WithMultigateway)
	Multigateway              *ProcessInstance
	MultigatewayPgPort        int // PostgreSQL protocol port for multigateway
	MultigatewayReplicaPgPort int // PostgreSQL replica-reads port for multigateway (0 = disabled)

	// Multiadmin instance (optional, enabled via WithMultiadmin).
	// The HTTP port is what the Next.js web UI in web/multiadmin/ talks to;
	// point it via MULTIADMIN_API_URL=http://localhost:<MultiadminHttpPort>.
	Multiadmin         *ProcessInstance
	MultiadminHttpPort int
	MultiadminGrpcPort int

	// PgBackRestCertPaths stores the paths to pgBackRest TLS certificates
	PgBackRestCertPaths *local.PgBackRestCertPaths

	// MultigatewayTLSCertPaths stores the paths to multigateway TLS certificates.
	// Set when WithMultigatewayTLS() is used.
	MultigatewayTLSCertPaths *MultigatewayTLSCertPaths

	// MultipoolerPGTLSCertPaths stores the paths used to provision postgres with
	// TLS and the matching CA the multipooler verifies against.
	// Set when WithMultipoolerPGTLS() is used.
	MultipoolerPGTLSCertPaths *MultipoolerPGTLSCertPaths

	// MetricsPorts maps instance name to its Prometheus metrics port.
	// Set when WithMetricsExport() is used. Scrape http://localhost:<port>/metrics.
	MetricsPorts map[string]int

	// BackupLocation stores backup configuration from topology
	BackupLocation *clustermetadatapb.BackupLocation

	// BaselineGucs stores the GUC values captured after bootstrap completes.
	// These are the "clean state" values that ValidateCleanState checks against
	// and that cleanup restores to. Structure: node name → GUC name → value.
	// After bootstrap with replication configured, this includes:
	// - Primary: synchronous_standby_names, synchronous_commit
	// - Replicas: primary_conninfo
	BaselineGucs map[string]map[string]string
}

// Context returns the running context for this setup, which is cancelled when Cleanup() is called.
// Use this when starting processes that should live for the lifetime of the cluster.
func (s *ShardSetup) Context() context.Context {
	_ = "STUB: not implemented"
	return *

	// StopEtcd gracefully stops the etcd process to simulate an etcd outage.
	new(context.Context)
}

func (s *ShardSetup) StopEtcd(t *testing.T) { _ = "STUB: not implemented"; return }

// GetMultipoolerInstance returns a multipooler instance by name, or nil if not found.
func (s *ShardSetup) GetMultipoolerInstance(name string) *MultipoolerInstance {
	_ = "STUB: not implemented"
	return nil
}

// GetMultipooler returns the multipooler process for an instance by name.
// Convenience method for tests that need just the multipooler process.
func (s *ShardSetup) GetMultipooler(name string) *ProcessInstance {
	_ = "STUB: not implemented"
	return nil
}

// GetPgctld returns the pgctld process for an instance by name.
// Convenience method for tests that need just the pgctld process.
func (s *ShardSetup) GetPgctld(name string) *ProcessInstance { _ = "STUB: not implemented"; return nil }

// GetMultiOrch returns a multiorch instance by name, or nil if not found.
func (s *ShardSetup) GetMultiOrch(name string) *ProcessInstance {
	_ = "STUB: not implemented"
	return nil
}

// GetPrimary returns the multipooler instance that was elected as primary.
// Fails the test if no primary has been set (e.g., before bootstrap).
func (s *ShardSetup) GetPrimary(t *testing.T) *MultipoolerInstance {
	_ = "STUB: not implemented"
	return nil
}

// RefreshPrimary queries all multipoolers to find the current primary and updates PrimaryName.
func (s *ShardSetup) RefreshPrimary(t *testing.T) *MultipoolerInstance {
	_ = "STUB: not implemented"
	return nil
}

// GetStandbys returns all multipooler instances that are not the primary.
func (s *ShardSetup) GetStandbys() []*MultipoolerInstance { _ = "STUB: not implemented"; return nil }

// PrimaryMultipooler returns the multipooler process for the elected primary.
func (s *ShardSetup) PrimaryMultipooler(t *testing.T) *ProcessInstance {
	_ = "STUB: not implemented"
	return nil
}

// PrimaryPgctld returns the pgctld process for the elected primary.
func (s *ShardSetup) PrimaryPgctld(t *testing.T) *ProcessInstance {
	_ = "STUB: not implemented"
	return nil
}

// CreateMultipoolerInstance creates a new multipooler instance (pgctld + multipooler pair) with the given name.
// The instance is added to the setup's Multipoolers map.
// Follows the patterns from multipooler/setup_test.go.
func (s *ShardSetup) CreateMultipoolerInstance(t *testing.T, name string, grpcPort, pgPort, multipoolerPort int) *MultipoolerInstance {
	_ = "STUB: not implemented"
	return nil
}

// Generate pgBackRest certificates once for the entire setup (shared across all multipoolers)

// Allocate a port for pgBackRest server (one per multipooler)

// Allocate an HTTP port for pgctld health endpoints

// Allocate an HTTP port for multipooler (prevents port collision with dynamic allocation)

// Create pgctld instance

// Create multipooler instance with pgBackRest cert paths and port
// The name (e.g., "primary") is used as the service-id, combined with cell in the topology

// CreatePgctldInstance creates a new pgctld process instance configuration.
// Follows the pattern from multipooler/setup_test.go:createPgctldInstance.
func CreatePgctldInstance(t *testing.T, name, baseDir string, grpcPort, pgPort, httpPort, pgbackrestPort int, pgbackrestCertDir string, backupLocation *clustermetadatapb.BackupLocation) *ProcessInstance {
	_ = "STUB: not implemented"
	return nil
}

// Create data directory

// CreateMultipoolerProcessInstance creates a new multipooler process instance configuration.
// Follows the pattern from multipooler/setup_test.go:createMultipoolerInstance.
func CreateMultipoolerProcessInstance(t *testing.T, name, baseDir string, grpcPort, httpPort int, pgctldAddr string, pgctldDataDir string, pgPort int, etcdAddr string, cell string, certPaths *local.PgBackRestCertPaths, pgbackrestPort int) *ProcessInstance {
	_ = "STUB: not implemented"
	return nil
}

// Create log directory

// Default to the standard Unix socket path under the pgctld data dir.
// Tests that need the TCP path (e.g. PG TLS) clear this after creation.

// Store pgBackRest cert paths struct and port for later use when starting multipooler

// CreateMultiOrchInstance creates a new multiorch instance and adds it to the setup.
// Returns the instance and a cleanup function that should be deferred or called manually.
// The cleanup function gracefully terminates the process if it's still running.
func (s *ShardSetup) CreateMultiOrchInstance(t *testing.T, name string, watchTargets []string, config *SetupConfig) (*ProcessInstance, func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create data directory

// Use the instance name as the service ID

// Apply defaults if not specified (0s for fast tests)

// CreateMultigatewayInstance creates a multigateway process instance.
// Returns the created ProcessInstance. Does not start the process.
// Call Start() on the returned instance to start it, and waitForMultigatewayQueryServing() after bootstrap.
func (s *ShardSetup) CreateMultigatewayInstance(t *testing.T, name string, pgPort, httpPort, grpcPort int) *ProcessInstance {
	_ = "STUB: not implemented"
	return nil
}

// Add TLS cert paths if multigateway TLS is enabled

// CreateMultiadminInstance creates a multiadmin process instance.
// Returns the created ProcessInstance. Does not start the process.
func (s *ShardSetup) CreateMultiadminInstance(t *testing.T, name string, httpPort, grpcPort int) *ProcessInstance {
	_ = "STUB: not implemented"
	return nil
}

// WaitForMultigatewayQueryServing waits for multigateway to be able to execute queries.
// This verifies that multigateway has discovered poolers from topology and can route queries.
// Should be called AFTER bootstrap completes.
//
// The timeout is generous (30s) because after bootstrap, the multigateway needs time to:
// 1. Receive the topology watch notification that pooler-1 was promoted to PRIMARY
// 2. Update its LoadBalancer with the new PRIMARY pooler
// This typically takes a few seconds but can be longer under load or slow CI environments.
func (s *ShardSetup) WaitForMultigatewayQueryServing(t *testing.T) {
	_ = "STUB: not implemented"

	// When TLS is configured on the gateway, use sslmode=require so this
	// readiness probe still works under --pg-require-ssl=true. The probe
	// only needs an encrypted transport, not certificate verification, so
	// sslmode=require is sufficient regardless of cert chain.
	return
}

// Verify both read and write paths work. SELECT 1 may succeed
// via a REPLICA before multigateway learns about the PRIMARY.
// CREATE TABLE forces routing to PRIMARY, confirming that
// multigateway has discovered the primary pooler.

// Cleanup cleans up the shared test infrastructure.
// If testsFailed is true, preserves the temp directory with logs for debugging.
// Follows the pattern from multipooler/setup_test.go:cleanupSharedTestSetup.
func (s *ShardSetup) Cleanup(testsFailed bool) { _ = "STUB: not implemented"; return }

// Gracefully terminate all processes BEFORE cancelling the context.
// For pgctld this issues pg_ctl stop via gRPC, which releases System V
// shared memory segments. The gRPC call must happen while the context
// is still active so it can reach the running pgctld.

// pgctld needs longer: stopPostgreSQL issues pg_ctl stop
// via gRPC (10s timeout) before SIGTERM is sent.

// Cancel the context to terminate any remaining processes. Processes
// wrapped in run_in_test.sh (etcd) are started with
// context.Background() and use a separate goroutine to detect context
// cancellation and clean up.

// Close topology server (can do this immediately since context cancellation is async)

// Clean up temp directory only if tests passed

// PrintLogLocation prints the temp directory location for debugging.
// If TEST_PRINT_LOGS env var is set, also prints all log contents from the temp directory.
func PrintLogLocation(tempDir string) { _ = "STUB: not implemented"; return }

// Only print log contents if TEST_PRINT_LOGS is set

// Print all .log files found in the temp directory

//nolint:nilerr // Continue walking even if one file fails

// DumpServiceLogs prints the location of service log files to help debug test failures.
// Call this before cleanup so logs are available.
// Always prints the temp directory location. If TEST_PRINT_LOGS env var is set, also prints log contents.
// Follows the pattern from multipooler/setup_test.go:dumpServiceLogs.
func (s *ShardSetup) DumpServiceLogs() { _ = "STUB: not implemented"; return }

// Use the shared utility function which prints location and optionally all logs

// CheckSharedProcesses verifies all shared test processes are still running.
// This catches crashes from previous tests early, before confusing timeout errors.
// Follows the pattern from multipooler/setup_test.go:checkSharedProcesses.
func (s *ShardSetup) CheckSharedProcesses(t *testing.T) { _ = "STUB: not implemented"; return }

// Check multigateway

// Check multiadmin

// Check multipooler instances

// TODO (@rafa): We can check multiorch processes once
// we are able to disable them on a shard basis.

// TestTarget represents a connection target for running tests against.
// Use GetComparisonTargets to obtain targets for both direct PostgreSQL and multigateway,
// enabling the same test logic to verify proxy behavior matches native PostgreSQL.
type TestTarget struct {
	// Name identifies the target (e.g., "postgres", "multigateway").
	Name string
	// Port is the PostgreSQL protocol port to connect to.
	Port int
}

// GetComparisonTargets returns test targets for both the primary PostgreSQL instance and
// the multigateway. Running identical tests against both ensures the proxy behavior
// matches native PostgreSQL exactly.
func (s *ShardSetup) GetComparisonTargets(t *testing.T) []TestTarget {
	_ = "STUB: not implemented"
	return nil
}
