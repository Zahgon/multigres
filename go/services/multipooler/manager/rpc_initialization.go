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

package manager

import (
	"context"

	multipoolermanagerdatapb "github.com/multigres/multigres/go/pb/multipoolermanagerdata"
)

// multigresInitMarker is the filename for the initialization marker.
// This file is created after full initialization completes (schema created, backup done).
const multigresInitMarker = "MULTIGRES_INITIALIZED"

// isInitialized checks if the pooler has been initialized (has data directory and multigres schema)
// This should return true even when postgres is not running, as long as the node was previously initialized.
func (pm *MultiPoolerManager) isInitialized(ctx context.Context) bool {
	_ = "STUB: not implemented"
	// Fast path: check cached state first
	return false
}

// Try to check if multigres schema exists via a query

// This marker is created after full initialization completes (schema created, backup done).
// It's more reliable than checking for PG_VERSION/global because those exist after initdb
// but before the full initialization process completes.

// Update cached state if we discovered initialization

// setInitialized marks the pooler as initialized and writes the marker file.
// This should be called after successful initialization (primary init, standby init, or restore).
// Once set, the pooler will skip auto-restore attempts.
func (pm *MultiPoolerManager) setInitialized() error { _ = "STUB: not implemented"; return nil }

// writeInitializationMarker creates the initialization marker file to indicate
// that full initialization, including restore from backup (for standbys) has
// completed. This is called at the end of primary and standby initialization.
//
// The marker file is needed to determine whether a replica pooler is
// initialized, because replica initialization is not done until the restore
// from backup completes. There is no other persistent way to determine this.
func (pm *MultiPoolerManager) writeInitializationMarker() error {
	_ = "STUB: not implemented"
	return nil
}

// hasDataDirectory checks if the PostgreSQL data directory exists
func (pm *MultiPoolerManager) hasDataDirectory() bool {
	_ = "STUB: not implemented"
	// Check if PG_VERSION file exists to confirm the data directory is properly initialized.
	// This prevents treating an empty directory (e.g., left behind by a failed initdb) as initialized.
	return false
}

// isPostgresRunning checks if the PostgreSQL process exists, regardless of
// whether it accepts connections. Returns true if the process is running (even if
// suspended via SIGSTOP). Returns false if the process is dead (e.g. after SIGKILL).
//
// When pgctld is not available, falls back to isPostgresReady (which requires
// both process existence and connection acceptance).
func (pm *MultiPoolerManager) isPostgresRunning(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false

	// No pgctld client — fall back to connection-based check.
	// Without pgctld we can't distinguish a stopped-but-alive process from a dead one.
}

// Only check if the process is running; do NOT require pg_isready (statusResp.Ready).

// isPostgresReady checks if PostgreSQL is currently running and accepting connections.
// Returns true only if the process is running AND pg_isready succeeds.
// Use isPostgresRunning to check only if the process exists.
func (pm *MultiPoolerManager) isPostgresReady(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false

	// No pgctld client, try a simple query to check if PostgreSQL is responding
}

// getServerStatus returns the observed state of the PostgreSQL server process.
// Priority: STARTING (action lock) > PROMOTING (in-flight pg_promote) > PRIMARY/STANDBY (pg_is_in_recovery).
func (pm *MultiPoolerManager) getServerStatus(ctx context.Context) multipoolermanagerdatapb.PostgresStatus {
	_ = "STUB: not implemented"
	return *new(multipoolermanagerdatapb.PostgresStatus)
}

// getWALPosition returns the current WAL position and any error encountered
func (pm *MultiPoolerManager) getWALPosition(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// getShardID returns the shard ID for this pooler.
// Prefers the topology value (pm.multipooler.Shard) but falls back to config
// if topology hasn't loaded yet. These should always be identical since
// the topology value is set from config at registration (init.go).
func (pm *MultiPoolerManager) getShardID() string { _ = "STUB: not implemented"; return "" }

// Fall back to MultiPooler - always available and authoritative

// removeDataDirectory removes the PostgreSQL data directory
func (pm *MultiPoolerManager) removeDataDirectory() error { _ = "STUB: not implemented"; return nil }

// Safety check: ensure we're not deleting root or home directory

// waitForDatabaseConnection waits for the database connection to become available
func (pm *MultiPoolerManager) waitForDatabaseConnection(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Test if database is already reachable
	return nil
}

// Start heartbeat tracker if not already running

// default shard ID

// Wait for connection to become available with retry logic

// Use exponential backoff starting at 500ms, up to 30s max backoff

// Check if context was cancelled or exceeded deadline

// Try to query the database

// Start heartbeat tracker if not already running

// default shard ID

// Don't fail - heartbeat is not critical for initialization

// This should not be reached due to the context check in the loop, but just in case

// removeArchiveConfigFromAutoConf removes archive configuration lines from postgresql.auto.conf
// This is used after restore to remove the primary's archive config before applying the standby's config
func (pm *MultiPoolerManager) removeArchiveConfigFromAutoConf() error {
	_ = "STUB: not implemented"
	return nil
}

// File doesn't exist, nothing to remove

// Skip archive-related lines

// configureArchiveMode configures archive_mode in postgresql.auto.conf for pgbackrest
// This must be called after InitDataDir but BEFORE starting PostgreSQL
func (pm *MultiPoolerManager) configureArchiveMode(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if pgbackrest config file exists before configuring archive mode

// Check if archive_mode is already configured to avoid duplicates

// Configure archive_mode in postgresql.auto.conf
// Following the pattern from test/endtoend/multipooler/setup_test.go:479-498
