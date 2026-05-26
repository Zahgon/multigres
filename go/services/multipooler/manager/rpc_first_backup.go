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

package manager

import (
	"context"

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
)

// createFirstBackupAndInitializeLocked attempts to create the first pgBackRest backup for this shard.
//
// Each pooler independently runs initdb, starts postgres, creates the multigres schema,
// and initializes the pgBackRest stanza locally. Then it races via the backup lease to
// create the actual backup. If another pooler wins the race, a backup will already exist
// when we acquire the lease and we skip to restore. After the backup exists, the local
// data directory is removed so all poolers restore from the shared backup.
//
// Returns (busy=true, backupFound=false, nil) if the backup lease is held by another pooler —
// the monitor should back off and retry. Returns (false, true, nil) if a backup was found
// (created by another pooler) — the caller should restore immediately.
func (pm *MultiPoolerManager) createFirstBackupAndInitializeLocked(ctx context.Context) (busy bool, backupFound bool, retErr error) {
	_ = "STUB: not implemented"
	return false, false, nil
}

// A sentinel from a prior attempt means we crashed between initdb and the
// final data-directory cleanup. The data directory (if any) is stale and
// safe to remove so we can retry from scratch. The sentinel is recreated
// below before initdb, so we do not remove it here.

// os.RemoveAll (inside removeDataDirectory) is idempotent: a prior crash
// between removeDataDirectory and removeBootstrapSentinel returns nil
// here. Any non-nil error is a real failure worth surfacing.

// Refuse to run on an already-initialized data directory. This node is already initialized

// Read the durability policy from topology before doing any expensive work.
// A misconfigured database (missing durability_policy) should fail fast.
// The policy is also written into the initial row of current_rule so all
// subsequent rule reads carry a non-nil DurabilityPolicy.

// Write the sentinel before initdb so a process crash between here and the
// final cleanup leaves a detectable marker. The sentinel lives in pooler_dir
// (not PGDATA), so it is not captured by pgBackRest backups.

// Cleanup runs on every exit path. On success it tears down the local
// postgres so every pooler restores from the shared backup. On failure it
// removes whatever was partially created. If data-directory removal fails
// we leave the sentinel in place so the next attempt knows to re-clean;
// on an otherwise-successful run we also promote the failure to retErr so
// the caller retries, since leaving stale pg_data would block the next
// create-first-backup tick at the hasDataDirectory() guard.
// Ordering matters: only clear the sentinel after the data directory is
// gone, to preserve the "data dir present ⇒ sentinel present" invariant.

// Initialize a fresh data directory.

// Configure archive_mode before starting PostgreSQL.

// Start PostgreSQL.

// The initial row (term=0, empty cohort) was already inserted by
// createRuleTables via createSidecarSchema above. This signals to multiorch
// that the shard has been initialized but not yet had its cohort established.

// Race for the backup lease. Only the lease holder creates the backup;
// everyone else will find an existing backup and skip to restore.
// Stanza-create runs inside the lease because it is part of the pgBackRest
// setup work that must complete before the backup.

// Re-check inside the lease — another pooler may have created the backup
// between our outer check and acquiring the lease.

// lease held by another pooler, back off

func (pm *MultiPoolerManager) bootstrapSentinelPath() string { _ = "STUB: not implemented"; return "" }

// hasBootstrapSentinel reports whether the sentinel file exists. A non-existent
// file is (false, nil); any other stat failure (e.g. permissions) is surfaced
// as an error so callers don't silently treat it as "not present".
func (pm *MultiPoolerManager) hasBootstrapSentinel() (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (pm *MultiPoolerManager) writeBootstrapSentinel() error { _ = "STUB: not implemented"; return nil }

// removeBootstrapSentinel deletes the sentinel; a missing file is not an error.
func (pm *MultiPoolerManager) removeBootstrapSentinel() error {
	_ = "STUB: not implemented"
	return nil
}

// runStanzaCreate runs pgbackrest stanza-create. The operation is idempotent
// and safe to run concurrently — no backup lease required.
func (pm *MultiPoolerManager) runStanzaCreate(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// loadDurabilityPolicy reads the bootstrap durability policy from the topology database record.
func (pm *MultiPoolerManager) loadDurabilityPolicy(ctx context.Context) (*clustermetadatapb.DurabilityPolicy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
