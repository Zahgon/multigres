// Copyright 2026 Supabase, Inc.
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

package pgctld

import (
	"net"
	"testing"
	"time"

	"github.com/multigres/multigres/go/pb/pgctldservice"
	"github.com/multigres/multigres/go/tools/executil"
)

// TestSetup holds all configuration for pgBackRest server tests
type TestSetup struct {
	TempDir        string
	PoolerDir      string
	CertDir        string
	PgPort         int
	PgBackRestPort int
	BinDir         string
}

// setupPgBackRestTest creates a complete test environment for pgBackRest server tests
// Returns TestSetup with all paths and ports configured
func setupPgBackRestTest(t *testing.T) *TestSetup {
	_ = "STUB: not implemented"

	// Create temp directory
	return nil
}

// Setup mock PostgreSQL binaries

// Set PATH for PostgreSQL binaries

// Generate TLS certificates

// Ensure cert files exist

// Allocate dynamic ports (port 0 = let OS choose)
// PostgreSQL needs a specific port for mock binaries

// Find a free port for pgBackRest

// Close immediately - pgctld will bind to it

// verifyServerRunning polls socket connection until success or timeout
// Returns true if server responds, false if timeout reached
func verifyServerRunning(t *testing.T, port int, timeout time.Duration) bool {
	_ = "STUB: not implemented"
	return false
}

// Continue polling

// verifyServerStopped verifies socket connection fails (server not listening)
// Returns true if port is closed, false if still accepting connections
func verifyServerStopped(t *testing.T, port int) bool { _ = "STUB: not implemented"; return false }

// Connection failed - server is stopped

// Connection succeeded - server still running

// getPgBackRestPID finds the pgbackrest server process ID for a specific config path
// Returns PID or 0 if not found
func getPgBackRestPID(t *testing.T, configPath string) int {
	_ = "STUB: not implemented"

	// Get all pgbackrest server processes
	return 0
}

// Process not found

// Check each PID to find one using our specific config

// Check if this process is using our config by checking environment

// killProcess sends SIGKILL to process by PID
func killProcess(t *testing.T, pid int) { _ = "STUB: not implemented"; return }

// Send SIGKILL

// createTestGRPCServerWithPgBackRest creates and starts a gRPC server with pgBackRest support
// Returns the listener and a cleanup function
func createTestGRPCServerWithPgBackRest(t *testing.T, setup *TestSetup) (net.Listener, func()) {
	_ = "STUB: not implemented"

	// Find a free port
	return *new(net.Listener), nil
}

// Create gRPC server

// Create the pgctld service with pgBackRest configuration

// Start pgBackRest management (runs in background goroutine)

// Register the service

// Start server in background

// Give server time to start

// Return cleanup function that stops server and closes service

// createPgCtldClient creates a gRPC client for pgctld
func createPgCtldClient(t *testing.T, addr string) pgctldservice.PgCtldClient {
	_ = "STUB: not implemented"
	return *new(pgctldservice.PgCtldClient)
}

// initAndStartPostgreSQL initializes and starts PostgreSQL via gRPC
func initAndStartPostgreSQL(t *testing.T, client pgctldservice.PgCtldClient) {
	_ = "STUB: not implemented"
	return
}

// Initialize data directory

// Start PostgreSQL

// Wait for PostgreSQL to be ready

// getPgBackRestStatus gets pgBackRest status from pgctld
func getPgBackRestStatus(t *testing.T, client pgctldservice.PgCtldClient) *pgctldservice.PgBackRestStatus {
	_ = "STUB: not implemented"
	return nil
}

// mockBinEnv returns the base environment for tests that use mock PostgreSQL
// binaries. It includes PATH pointing at binDir, PGDATA, POSTGRES_PASSWORD and
// PGPASSWORD so pgctld init/start/stop subprocesses can satisfy the scram-sha-256
// requirement without needing a real PostgreSQL instance.
func mockBinEnv(binDir, pgDataDir string) []string { _ = "STUB: not implemented"; return nil }

// setupTestEnv sets common environment variables for pgctld subprocesses.
// Includes a default POSTGRES_PASSWORD because pgctld now refuses to init or
// serve without one (pg_hba.conf requires scram-sha-256 for all connections).
// Also sets PGPASSWORD so psql subprocesses can authenticate.
func setupTestEnv(cmd *executil.Cmd, poolerDir string) { _ = "STUB: not implemented"; return }

// Required to avoid "postmaster became multithreaded during startup" on macOS.

// pgCtldServer holds the ports and process handle for a running pgctld server.
type pgCtldServer struct {
	GrpcPort int
	HttpPort int
	PgPort   int
	Cmd      *executil.Cmd
}

// startPgCtldServer starts a pgctld server subprocess and waits until its gRPC
// port is accepting connections. poolerDir is passed as --pooler-dir.
// configFile is optional; if non-empty it is passed as --config-file.
// extraEnv may be used to inject additional environment variables such as
// POSTGRES_PASSWORD or a custom PATH.
//
// The server is stopped automatically when the test ends via t.Cleanup.
// The returned Cmd may also be used directly (e.g. to kill the process
// mid-test for orphan-detection tests).
func startPgCtldServer(t *testing.T, poolerDir, configFile string, extraEnv ...string) *pgCtldServer {
	_ = "STUB: not implemented"
	return nil
}

// Wait for gRPC port to accept connections.
