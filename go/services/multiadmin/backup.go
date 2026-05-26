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

package multiadmin

import (
	"context"

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	multiadminpb "github.com/multigres/multigres/go/pb/multiadmin"
)

// Backup starts an async backup of a specific shard
func (s *MultiAdminServer) Backup(ctx context.Context, req *multiadminpb.BackupRequest) (*multiadminpb.BackupResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Find a pooler synchronously so we can generate a stable job ID.
// The job ID includes the pooler name, which enables recovery after multiadmin restart.

// Generate a job ID that will be stored in pgbackrest and can be queried after restart.
// Format: YYYYMMDD-HHMMSS.microseconds_<pooler_name>

// Create a job to track this backup operation

// Start the backup operation asynchronously with a linked root span

// executeBackup performs the actual backup operation
func (s *MultiAdminServer) executeBackup(ctx context.Context, jobID string, pooler *clustermetadatapb.MultiPooler, req *multiadminpb.BackupRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// Call backup on the pooler using the shared rpcClient.

// Mark job as completed

// findPoolerForBackup finds a pooler for backup operations.
// If forcePrimary is true, finds a PRIMARY pooler; otherwise finds a REPLICA.
func (s *MultiAdminServer) findPoolerForBackup(ctx context.Context, database, tableGroup, shard string, forcePrimary bool) (*clustermetadatapb.MultiPooler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get all cells

// Search across all cells for a pooler of the target type
// TODO: Pick the replica with the least replica lag, as measured by the heartbeat service

// Note: Shard is intentionally left empty to match all shards.
// Multipoolers currently don't set Shard when registering.

// Find a pooler of the target type

// RestoreFromBackup starts an async restore of a specific shard from a backup
func (s *MultiAdminServer) RestoreFromBackup(ctx context.Context, req *multiadminpb.RestoreFromBackupRequest) (*multiadminpb.RestoreFromBackupResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate request

// Note: TableGroup and Shard validation ensures the restore target is fully specified.
// The CLI currently hardcodes these to defaults, but the API is designed to support
// the full use case when multi-shard support is implemented.

// Create job

// TODO: store job_id somewhere persistent, such as the PGDATA directory or topo,
// so that job state is not lost after multiadmin restart

// Start restore in background with a linked root span

// executeRestore performs the actual restore operation
// TODO: Add job_id support for RestoreFromBackup similar to Backup():
//  1. Add JobId field to RestoreFromBackupRequest proto
//  2. Generate backupJobID = backup.GenerateJobID(pooler.Id.Name)
//  3. Store job_id as pgbackrest annotation on restore
//  4. Add GetRestoreByJobId RPC or extend GetBackupByJobId for restore tracking
func (s *MultiAdminServer) executeRestore(ctx context.Context, jobID string, req *multiadminpb.RestoreFromBackupRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// Get pooler info from topology

// Call restore on the pooler

// Mark job as completed

// GetBackupJobStatus checks the status of a backup or restore job
func (s *MultiAdminServer) GetBackupJobStatus(ctx context.Context, req *multiadminpb.GetBackupJobStatusRequest) (*multiadminpb.GetBackupJobStatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate request

// Try the in-memory job tracker first

// Job not in tracker - try fallback to MultiPooler if shard context provided.
// This handles the case where the multiadmin process restarted and lost in-memory job state.

// getBackupJobStatusFromPooler queries a MultiPooler for backup status when job is not in memory.
func (s *MultiAdminServer) getBackupJobStatusFromPooler(ctx context.Context, req *multiadminpb.GetBackupJobStatusRequest) (*multiadminpb.GetBackupJobStatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Find a replica pooler - all poolers for a shard share the same pgbackrest repo

// Query the pooler for backup by job_id

// No backup found with this job_id

// Convert backup metadata to job status

// GetBackups lists backup artifacts with optional filtering
func (s *MultiAdminServer) GetBackups(ctx context.Context, req *multiadminpb.GetBackupsRequest) (*multiadminpb.GetBackupsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate request

// Find a replica pooler - all replicas for a shard share the same pgbackrest repo

// Query backups from the pooler

// Convert multipoolermanagerdata.BackupMetadata to multiadminpb.BackupInfo

// ExpireBackups removes old backups according to retention policy.
// It finds a replica pooler and proxies the request to it.
func (s *MultiAdminServer) ExpireBackups(ctx context.Context, req *multiadminpb.ExpireBackupsRequest) (*multiadminpb.ExpireBackupsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Find a replica pooler — all replicas for a shard share the same pgbackrest repo
