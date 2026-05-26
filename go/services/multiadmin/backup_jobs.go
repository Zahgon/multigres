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
	"sync"
	"time"

	multiadminpb "github.com/multigres/multigres/go/pb/multiadmin"
)

// DefaultBackupJobExpiration is the default time after which completed/failed backup jobs are removed
const DefaultBackupJobExpiration = 24 * time.Hour

// backupJobMetadata stores additional metadata not in the proto
type backupJobMetadata struct {
	response    multiadminpb.GetBackupJobStatusResponse // Changed from pointer to value
	createdAt   time.Time
	updatedAt   time.Time
	completedAt *time.Time
}

// BackupJobTracker manages async backup/restore jobs.
//
// This is a simple in-memory implementation, so process restarts will lose all jobs.
type BackupJobTracker struct {
	mu         sync.Mutex
	jobs       map[string]*backupJobMetadata
	expiration time.Duration

	// Shutdown-related channels
	stop chan struct{}
	done chan struct{}
}

// NewBackupJobTracker creates a new backup job tracker with default expiration (24 hours)
func NewBackupJobTracker() *BackupJobTracker { _ = "STUB: not implemented"; return nil }

// NewBackupJobTrackerWithExpiration creates a new backup job tracker with custom expiration
func NewBackupJobTrackerWithExpiration(expiration time.Duration) *BackupJobTracker {
	_ = "STUB: not implemented"
	return nil
}

// cleanupLoop runs in the background and removes expired jobs
func (jt *BackupJobTracker) cleanupLoop() { _ = "STUB: not implemented"; return }

// removeExpiredJobs removes jobs that have been completed/failed for longer than the expiration period
func (jt *BackupJobTracker) removeExpiredJobs() { _ = "STUB: not implemented"; return }

// Stop stops the background cleanup goroutine
func (jt *BackupJobTracker) Stop() { _ = "STUB: not implemented"; return }

// CreateJob creates a new backup job with an auto-generated UUID and returns its ID
func (jt *BackupJobTracker) CreateJob(jobType multiadminpb.JobType, database, tableGroup, shard string) string {
	_ = "STUB: not implemented"
	return ""
}

// CreateJobWithID creates a new backup job with the specified ID
func (jt *BackupJobTracker) CreateJobWithID(jobID string, jobType multiadminpb.JobType, database, tableGroup, shard string) string {
	_ = "STUB: not implemented"
	return ""
}

// Copies are ok here

// GetJobStatus retrieves the status of a backup job
func (jt *BackupJobTracker) GetJobStatus(jobID string) (*multiadminpb.GetBackupJobStatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Clone the response to prevent race conditions (protobuf messages contain internal mutexes)

// UpdateJobStatus updates the status of a backup job
func (jt *BackupJobTracker) UpdateJobStatus(jobID string, status multiadminpb.JobStatus) {
	_ = "STUB: not implemented"
	return
}

// CompleteJob marks a backup job as completed with results
func (jt *BackupJobTracker) CompleteJob(jobID string, backupID string) {
	_ = "STUB: not implemented"
	return
}

// FailJob marks a backup job as failed with error message
func (jt *BackupJobTracker) FailJob(jobID string, errorMsg string) {
	_ = "STUB: not implemented"
	return
}
