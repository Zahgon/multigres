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

package local

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/multigres/multigres/go/provisioner"
	"github.com/multigres/multigres/go/tools/pathutil"

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"

	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("github.com/multigres/multigres/go/provisioner/local")

// localProvisioner implements the Provisioner interface for local binary-based provisioning
type localProvisioner struct {
	config              *LocalProvisionerConfig
	pgBackRestCertPaths *PgBackRestCertPaths
}

// Compile-time check to ensure localProvisioner implements Provisioner
var _ provisioner.Provisioner = (*localProvisioner)(nil)

const (
	// StateDir is the directory name where provision state files are stored
	StateDir = "state"
)

// Name returns the name of this provisioner
func (p *localProvisioner) Name() string {
	_ = "STUB: not implemented"

	// initializePgctldDirectories creates all pgctld pooler directories based on the config.
	return ""
}

func (p *localProvisioner) initializePgctldDirectories() error {
	_ = "STUB: not implemented"
	return nil
}

// provisionEtcd provisions etcd using local binary
func (p *localProvisioner) provisionEtcd(ctx context.Context, req *provisioner.ProvisionRequest) (*provisioner.ProvisionResult, error) {
	_ = "STUB: not implemented"
	// Sanity check: ensure this method is called for etcd service
	return nil, nil
}

// Check if etcd is already running by checking state

// Get port from config or use default

// Get peer port from config, or default to port + 1

// Get metrics port from config, or default to port + 2.
// The metrics listener is required for /readyz health checks.

// Find etcd binary (PATH or configured path)

// Check etcd version

// Create data directory

// Generate unique ID for this service instance (needed for log file)

// Create log file path

// Start etcd process

// Validate process is running

// Wait for etcd to be ready

// Create provision state

// Save service state to disk

// findBinary finds a binary by name, checking PATH first, then the executable directory,
// and then the optional configured path
func (p *localProvisioner) findBinary(name string, serviceConfig map[string]any) (string, error) {
	_ = "STUB: not implemented"
	// First try to find in PATH
	return "", nil
}

// Then try configured path if provided

// Check if it's an absolute path or relative path

// Make it relative to current directory

// Check if the binary exists and is executable

// checkEtcdVersion verifies that the etcd binary major version matches expected version
func (p *localProvisioner) checkEtcdVersion(binaryPath, expectedVersion string) error {
	_ = "STUB: not implemented"
	// Run etcd --version to get version info
	return nil
}

// Parse version from output - etcd version output format varies

// Try to extract version number from various etcd output formats

// If we can't parse version, just warn and continue

// ensure v prefix for semver

// Use servenv semver to compare major versions

// readServiceLogs reads the last few lines from a service's log file for debugging
func (p *localProvisioner) readServiceLogs(logFile string, lines int) string {
	_ = "STUB: not implemented"
	return ""
}

// Check if log file exists

// Read the file

// Get the last N lines

// Return last 'lines' lines or all lines if fewer exist

// getRootWorkingDir returns the root working directory from config
func (p *localProvisioner) getRootWorkingDir() string { _ = "STUB: not implemented"; return "" }

// GeneratePoolerDir generates a pooler directory path for a given base directory and service ID
func GeneratePoolerDir(baseDir, serviceID string) string { _ = "STUB: not implemented"; return "" }

// provisionMultigateway provisions multigateway using either binaries or Docker containers
func (p *localProvisioner) provisionMultigateway(ctx context.Context, req *provisioner.ProvisionRequest) (*provisioner.ProvisionResult, error) {
	_ = "STUB: not implemented"
	// Sanity check: ensure this method is called for multigateway service
	return nil, nil
}

// Get cell parameter

// Check if multigateway is already running

// Get parameters from request

// Get cell-specific multigateway config

// Get HTTP port from cell-specific config

// Get gRPC port from cell-specific config

// Get pg port from cell-specific config

// Get log level

// Find multigateway binary

// Generate unique ID for this service instance (needed for log file)

// Create log file path

// Build command arguments

// Start multigateway process

// Validate process is running

// Create provision state

// Save service state to disk

// Wait for multigateway to be ready

// provisionMultiadmin provisions multiadmin using local binary
func (p *localProvisioner) provisionMultiadmin(ctx context.Context, req *provisioner.ProvisionRequest) (*provisioner.ProvisionResult, error) {
	_ = "STUB: not implemented"
	// Sanity check: ensure this method is called for multiadmin service
	return nil, nil
}

// Check if multiadmin is already running

// Get multiadmin config

// Get HTTP port from config

// Get gRPC port from config

// Get parameters from request

// Get log level

// Find multiadmin binary

// Generate unique ID for this service instance (needed for log file)

// Create log file path

// Build command arguments

// Start multiadmin process

// Validate process is running

// Create provision state

// Save service state to disk

// Wait for multiadmin to be ready (check HTTP port)

// provisionMultipooler provisions multipooler using local binary
func (p *localProvisioner) provisionMultipooler(ctx context.Context, req *provisioner.ProvisionRequest) (*provisioner.ProvisionResult, error) {
	_ = "STUB: not implemented"
	// Sanity check: ensure this method is called for multipooler service
	return nil, nil
}

// Get cell parameter

// Check if multipooler is already running

// Get parameters from request

// Get cell-specific multipooler config

// Get HTTP port from cell-specific config

// Get grpc port from cell-specific config

// Get database from multipooler config, fall back to request if not set

// Get table group from multipooler config, default to "default" if not set

// Get shard from multipooler config, default to "0-inf" if not set

// Get log level

// Get pooler directory

// Get PostgreSQL port from config or use default

// Get gRPC socket file if configured

// Find multipooler binary

// Get service ID from multipooler config - this should always be set

// Create log file path

// Provision pgctld for this multipooler

// Construct the PostgreSQL socket file path: <poolerDir>/pg_sockets/.s.PGSQL.<port>

// Build command arguments with pgctld-addr

// PostgreSQL Unix socket for trust auth

// Add socket file if configured

// Add service map configuration to enable grpc-pooler service

// Get pgbackrest port from pgctld config (pgbackrest is now managed by pgctld)

// Add pgbackrest TLS certificate paths and port

// Start multipooler process

// Point multipooler at the password file pgctld already wrote (and is
// using itself). Both services agree on the credential because they
// resolve it from the same file. Required: multipooler refuses to start
// without it (pg_hba.conf uses scram-sha-256 for every connection).

// Validate process is running

// Wait for multipooler to be ready

// Create provision state

// Save service state to disk

// PgctldProvisionResult contains the result of provisioning pgctld
type PgctldProvisionResult struct {
	Address string
	Port    int
	LogFile string
	// PasswordFile is the path to the postgres password file written under the
	// pooler directory. Both pgctld (already running with POSTGRES_PASSWORD_FILE
	// pointing here) and the multipooler in the same cell read from it.
	PasswordFile string
}

// provisionMultiOrch provisions multi-orchestrator using local binary
func (p *localProvisioner) provisionMultiOrch(ctx context.Context, req *provisioner.ProvisionRequest) (*provisioner.ProvisionResult, error) {
	_ = "STUB: not implemented"
	// Sanity check: ensure this method is called for multiorch service
	return nil, nil
}

// Get cell parameter

// Check if multiorch is already running

// Get parameters from request

// Get cell-specific multiorch config

// Get HTTP port from cell-specific config

// Get grpc port from cell-specific config

// Get log level

// Find multiorch binary

// Generate unique ID for this service instance (needed for log file)

// Create log file path

// Build command arguments

// Add optional interval configs if specified

// Start multiorch process

// Validate process is running

// Wait for multiorch to be ready

// Create provision state

// Save service state to disk

// Deprovision removes/stops a specific service
func (p *localProvisioner) Deprovision(ctx context.Context, req *provisioner.DeprovisionRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// Stop the service using the service-specific method

// Remove state file on successful stop

// loadServiceState loads a specific service state from disk
func (p *localProvisioner) loadServiceState(req *provisioner.DeprovisionRequest) (*LocalProvisionedService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// For database services: state/dbs/dbname

// For non-database services (like etcd): state/

// Check if state file exists

// Service not found

// Sanity check: ensure this method is called for the expected service type

// stopService stops a specific service based on its type using the internal methods
func (p *localProvisioner) stopService(ctx context.Context, req *provisioner.DeprovisionRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// pgctld requires special handling to stop PostgreSQL first

// deprovisionService(ctx stops a multiorch service instance
func (p *localProvisioner) deprovisionService(ctx context.Context, req *provisioner.DeprovisionRequest) error {
	_ = "STUB: not implemented"
	// Load the specific service state
	return nil
}

// Stop the process if it's running

// Clean up log file if it exists

// Remove state file

// Clean up data directory if requested

// stopProcessByPID stops a process by its PID
func (p *localProvisioner) stopProcessByPID(ctx context.Context, name string, pid int) error {
	_ = "STUB: not implemented"
	return nil
}

// Use executil.StopPID for graceful termination with 2s grace period

// Process exited with error, but it's stopped - that's ok for cleanup

// Bootstrap sets up etcd and creates the default database
func (p *localProvisioner) Bootstrap(ctx context.Context) ([]*provisioner.ProvisionResult, error) {
	_ = "STUB: not implemented"
	// Validate binary paths before starting
	return nil, nil
}

// Validate required system binaries before starting

// Provision etcd

// Setup default cell using the configured cell name

// Initialize pgctld directories and password files

// Get all cells and set them up

// Set up all cells

// Provision multiadmin (global admin service)

// Setup default database

// Teardown shuts down all services (reverse of Bootstrap)
func (p *localProvisioner) Teardown(ctx context.Context, clean bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Get the typed configuration

// Get etcd address (assuming etcd is running locally)

// 1. Deprovision database services first

// 2. Deprovision global services (multiadmin)

// multiadmin is a global service

// 3. Deprovision etcd last

// etcd is a global service

// 4. Clean up logs, state, and data directories if requested

// cleanupLogsDirectory removes the entire logs directory and all its contents
func (p *localProvisioner) cleanupLogsDirectory(logsDir string) error {
	_ = "STUB: not implemented"
	// Check if logs directory exists
	return nil
}

// Directory doesn't exist, nothing to clean up

// Remove the entire logs directory

// cleanupStateDirectory removes the entire state directory and all its contents
func (p *localProvisioner) cleanupStateDirectory(stateDir string) error {
	_ = "STUB: not implemented"
	// Check if state directory exists
	return nil
}

// Directory doesn't exist, nothing to clean up

// Remove the entire state directory

// cleanupDataDirectory removes the entire data directory and all its contents
func (p *localProvisioner) cleanupDataDirectory(dataDir string) error {
	_ = "STUB: not implemented"
	// Check if data directory exists
	return nil
}

// Directory doesn't exist, nothing to clean up

// Remove the entire data directory

// cleanupSocketsDirectory removes the entire sockets directory and all its contents
func (p *localProvisioner) cleanupSocketsDirectory(socketsDir string) error {
	_ = "STUB: not implemented"
	return nil
}

// Directory doesn't exist, nothing to clean up

// cleanupSpoolDirectory removes the entire spool directory and all its contents
func (p *localProvisioner) cleanupSpoolDirectory(spoolDir string) error {
	_ = "STUB: not implemented"
	return nil
}

// Directory doesn't exist, nothing to clean up

// getGRPCSocketFile extracts and prepares the gRPC socket file path from a service config.
// It returns the absolute path to the socket file and ensures the socket directory exists.
// Returns empty string if no socket file is configured.
func getGRPCSocketFile(serviceConfig map[string]any) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// No socket file configured

// Convert to absolute path since the working directory may change

// Ensure socket directory exists

// getDefaultDatabaseName returns the default database name from config
func (p *localProvisioner) getDefaultDatabaseName() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ProvisionDatabase provisions a complete database stack in all cells (assumes etcd is already running and cells are configured)
func (p *localProvisioner) ProvisionDatabase(ctx context.Context, databaseName string, etcdAddress string) ([]*provisioner.ProvisionResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create backup directory if using local backups

// S3 backups don't need local directory creation

// Get topology configuration from provisioner config

// Get all cell information

// Register database in global topology store first

// Check if database already exists

// Create the database if it doesn't exist

// Generate pgBackRest certificates before starting services

// Provision all services in parallel across all cells.
// Multiorch's bootstrap action has a quorum check that will wait for enough
// poolers to be available before attempting bootstrap, so strict ordering
// is not required.

// Calculate total number of services to provision
// multigateway + multipooler + multiorch per cell (pgbackrest now managed by pgctld)

// Start all services in parallel

// capture for goroutine

// Start multigateway

// Start multipooler

// Start multiorch

// Collect all results

// Only append non-nil results (pgbackrest returns nil)

// Report any errors

// buildBackupLocation creates a BackupLocation proto from config
func (p *localProvisioner) buildBackupLocation() (*clustermetadatapb.BackupLocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// No backup type configured - use default filesystem backup location

// generatePgBackRestCertsOnce generates pgBackRest certificates once for all cells
func (p *localProvisioner) generatePgBackRestCertsOnce(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Store the cert paths for later use

// setupDefaultCell initializes the topology cell configuration for a database
func (p *localProvisioner) setupDefaultCell(ctx context.Context, cellName, etcdAddress string) error {
	_ = "STUB: not implemented"
	return nil
}

// Get topology configuration

// Create topology store using configured backend

// Check if cell already exists

// Create the cell if it doesn't exist

// Get the specific cell config for this cell name

// Some other error occurred

// DeprovisionDatabase deprovisions all services for a database
func (p *localProvisioner) DeprovisionDatabase(ctx context.Context, databaseName string, etcdAddress string) error {
	_ = "STUB: not implemented"
	return nil
}

// Find all running services related to this database

// Clean up data when deprovisioning database

// Remove state file

// getTopologyConfig extracts topology configuration from provisioner config
func (p *localProvisioner) getTopologyConfig() (*TopologyConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getAllCells returns all configured cells
func (p *localProvisioner) getAllCells() ([]CellConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getCellNames returns the names of all configured cells
func (p *localProvisioner) getCellNames() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getCellByName returns the cell configuration for a specific cell name
func (p *localProvisioner) getCellByName(cellName string) (*CellConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Find the specific cell by name

// ValidateConfig validates the local provisioner configuration
func (p *localProvisioner) ValidateConfig(config map[string]any) error {
	_ = "STUB: not implemented"
	// Convert to typed configuration for validation
	return nil
}

// Validate required topology fields

// Validate each cell

// Validate Unix socket path length limits

// UnixPathMax returns the maximum Unix socket path length for the current platform.
func UnixPathMax() int { _ = "STUB: not implemented"; return 0 }

// validateUnixSocketPathLength validates that Unix socket paths won't exceed system limits
func (p *localProvisioner) validateUnixSocketPathLength(config *LocalProvisionerConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// Convert root working dir to absolute path for accurate length calculation

// Calculate the maximum possible path length for Unix sockets
// Path structure: <rootWorkingDir>/data/pooler_<serviceID>/pg_sockets/.s.PGSQL.5432
// We use a worst-case service ID length (8 chars) to be safe

// validateBinaryPaths validates that all configured binary paths exist and are executable
func (p *localProvisioner) validateBinaryPaths(config *LocalProvisionerConfig) error {
	_ = "STUB: not implemented"
	return nil

	// Validate global service binaries
}

// Validate cell service binaries

// Validate multigateway

// Validate multipooler

// Validate multiorch

// Validate pgctld

// validateBinaryExists checks if a binary path exists and is executable
func (p *localProvisioner) validateBinaryExists(binaryPath, serviceName string) error {
	_ = "STUB: not implemented"
	// Use exec.LookPath to find and validate the binary
	return nil
}

// validateSystemBinaries validates that required system binaries are available in PATH
func (p *localProvisioner) validateSystemBinaries() error { _ = "STUB: not implemented"; return nil }

// NewLocalProvisioner creates a new local provisioner instance
func NewLocalProvisioner() (provisioner.Provisioner, error) {
	_ = "STUB: not implemented"
	return *new(provisioner.Provisioner), nil
}

func getExecutablePath() (string, error) { _ = "STUB: not implemented"; return "", nil }

func init() {
	// Register the local provisioner
	provisioner.RegisterProvisioner("local", NewLocalProvisioner)

	// Add the executable directory to the PATH. We're expecting
	// to find the other executables in the same directory.
	if binDir, err := getExecutablePath(); err == nil {
		pathutil.PrependPath(binDir)
	} else {
		slog.Error("local provisioner failed to get executable path", "error", err)
		panic(fmt.Sprintf("local provisioner failed to get executable path: %v", err))
	}
}
