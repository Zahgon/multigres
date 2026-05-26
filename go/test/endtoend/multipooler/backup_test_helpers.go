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

// This file contains test helpers for backup_test.go to reduce duplication.
// These helpers follow Go testing best practices:
// - All helpers call t.Helper() for proper failure attribution
// - Helpers fail fast with require for critical assertions
// - Cleanup is automatic via t.Cleanup() where applicable

package multipooler

import (
	"database/sql"
	"regexp"
	"testing"
	"time"

	consensuspb "github.com/multigres/multigres/go/pb/consensus"
	multipoolermanagerpb "github.com/multigres/multigres/go/pb/multipoolermanager"
	multipoolermanagerdata "github.com/multigres/multigres/go/pb/multipoolermanagerdata"
)

// Backup ID format patterns
var (
	fullBackupIDPattern = regexp.MustCompile(`^\d{8}-\d{6}F$`)
)

// createBackupClient creates a gRPC client for backup operations.
// The connection is automatically closed via t.Cleanup.
func createBackupClient(t *testing.T, grpcPort int) multipoolermanagerpb.MultiPoolerManagerClient {
	_ = "STUB: not implemented"
	return *new(multipoolermanagerpb.MultiPoolerManagerClient)
}

// createConsensusClient creates a gRPC client for consensus operations.
// The connection is automatically closed via t.Cleanup.
func createConsensusClient(t *testing.T, grpcPort int) consensuspb.MultiPoolerConsensusClient {
	_ = "STUB: not implemented"
	return *new(consensuspb.MultiPoolerConsensusClient)
}

// assertBackupIDFormat verifies the backup ID matches the expected format.
func assertBackupIDFormat(t *testing.T, backupID string, backupType string) {
	_ = "STUB: not implemented"
	return
}

// findBackupInList searches for a backup by ID in the list of backups.
// Fails the test if the backup is not found.
func findBackupInList(t *testing.T, backups []*multipoolermanagerdata.BackupMetadata, backupID string) *multipoolermanagerdata.BackupMetadata {
	_ = "STUB: not implemented"
	return nil
}

// unreachable

// assertBackupComplete verifies a backup has completed successfully.
func assertBackupComplete(t *testing.T, backup *multipoolermanagerdata.BackupMetadata, expectedID string) {
	_ = "STUB: not implemented"
	return
}

// connectToPostgresViaSocket establishes a connection to PostgreSQL using Unix socket.
// The connection is automatically closed via defer in the caller.
func connectToPostgresViaSocket(t *testing.T, socketDir string, port int) *sql.DB {
	_ = "STUB: not implemented"
	return nil
}

// getPostgresSocketPath returns the path to the PostgreSQL Unix socket directory.
func getPostgresSocketPath(pgctldDataDir string) string { _ = "STUB: not implemented"; return "" }

// createAndVerifyBackup creates a backup and verifies it was created successfully.
// Returns the backup ID.
func createAndVerifyBackup(t *testing.T, client multipoolermanagerpb.MultiPoolerManagerClient, backupType string, forcePrimary bool, timeout time.Duration, overrides map[string]string) string {
	_ = "STUB: not implemented"
	return ""
}

// listAndFindBackup lists backups and finds a specific backup by ID.
// Returns the found backup metadata.
func listAndFindBackup(t *testing.T, client multipoolermanagerpb.MultiPoolerManagerClient, backupID string, limit uint32) *multipoolermanagerdata.BackupMetadata {
	_ = "STUB: not implemented"
	return nil
}
