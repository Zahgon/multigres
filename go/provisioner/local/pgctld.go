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
)

// startPostgreSQLViaPgctld checks PostgreSQL status via pgctld gRPC.
// It does NOT auto-initialize PostgreSQL - that's handled by multiorch's bootstrap process.
// This function only starts PostgreSQL if the data directory is already initialized.
func (p *localProvisioner) startPostgreSQLViaPgctld(ctx context.Context, address string) error {
	_ = "STUB: not implemented"
	return nil
}

// First, check if PostgreSQL is already running

// If already running, we're good

// If not initialized, skip starting PostgreSQL.
// Multiorch will handle initialization through the bootstrap process.

// Data directory exists but PostgreSQL is not running - start it

// Verify PostgreSQL is now running

// stopPostgreSQLViaPgctld stops PostgreSQL via pgctld gRPC
func (p *localProvisioner) stopPostgreSQLViaPgctld(ctx context.Context, address string) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if PostgreSQL is running

// If not running, nothing to stop

// Stop PostgreSQL with fast mode

// Verify PostgreSQL is now stopped

// provisionPgctld provisions a pgctld instance for a multipooler with the new directory structure
func (p *localProvisioner) provisionPgctld(ctx context.Context, dbName, tableGroup, serviceID, cell string) (*PgctldProvisionResult, error) {
	_ = "STUB: not implemented"
	// Create unique pgctld service ID using multipooler's service ID
	return nil, nil
}

// Resolve pgctld config up-front so we can materialize the postgres
// password file before either branch (already-running or fresh start)
// returns. The file is the wire format both pgctld and the multipooler
// in this cell consume via POSTGRES_PASSWORD_FILE.

// Check if pgctld is already running for this service combination

// Check if the existing service matches our specific service ID

// Verify PostgreSQL is running via gRPC health check

// Find pgctld binary

// Get gRPC port from config or use default

// Get HTTP port from config or use default

// Get PostgreSQL port from config or use default

// Get other pgctld configuration values with defaults

// Get gRPC socket file if configured

// Create pgctld log file

// Note: We do NOT run 'pgctld init' here because that would initialize
// the PostgreSQL data directory (initdb) before multiorch can bootstrap
// the cluster. Multiorch needs to control initialization to properly set up
// primary/standby replication across zones.

// Start pgctld server

// Add socket file if configured

// Add pgBackRest configuration if certificates are available

// Get pgbackrest port from config or use default

// On macOS, ensure a valid locale is set for pgctld and its children (initdb, pg_ctl).
// Without LC_ALL or LANG, initdb fails with "invalid locale settings".
// Only inject when neither is set; an existing value in either variable is left untouched.

// Set PGDATA so pgctld knows where the PostgreSQL data directory is.

// Point pgctld at the password file written above. pgctld reads it during
// init (--pwfile) and at server startup; multipooler reads the same file
// for its admin pool.

// Validate process is running

// Wait for pgctld to be ready

// Now that pgctld is healthy, start PostgreSQL

// Create provision state for pgctld

// Save pgctld service state to disk

// deprovisionPgctld stops PostgreSQL via gRPC and then stops the pgctld process
func (p *localProvisioner) deprovisionPgctld(ctx context.Context, service *LocalProvisionedService) error {
	_ = "STUB: not implemented"
	// First, try to gracefully stop PostgreSQL via pgctld gRPC
	return nil
}

// Then stop the pgctld process itself

// Clean up log file

// writePostgresPasswordFile materializes the postgres password into a 0600
// file at <poolerDir>/postgres-password and returns its path. The file is the
// wire format both pgctld and the multipooler in this cell consume via
// POSTGRES_PASSWORD_FILE; writing it idempotently (overwriting on each
// provision) keeps the file in sync with the YAML config when operators rotate
// the value. Cleanup is handled by the provisioner tearing down poolerDir.
func writePostgresPasswordFile(poolerDir, password string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
