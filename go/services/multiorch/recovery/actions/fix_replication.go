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

package actions

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/multigres/multigres/go/common/rpcclient"
	"github.com/multigres/multigres/go/common/topoclient"
	"github.com/multigres/multigres/go/services/multiorch/config"
	"github.com/multigres/multigres/go/services/multiorch/recovery/types"
	"github.com/multigres/multigres/go/services/multiorch/store"

	multiorchdatapb "github.com/multigres/multigres/go/pb/multiorchdata"
	multipoolermanagerdatapb "github.com/multigres/multigres/go/pb/multipoolermanagerdata"
)

// Compile-time assertion that FixReplicationAction implements types.RecoveryAction.
var _ types.RecoveryAction = (*FixReplicationAction)(nil)

// errPoolerDrained is returned by tryPgRewind when pg_rewind is not feasible and the
// pooler has been successfully marked DRAINED. The caller should stop attempting to
// verify replication and treat the action as complete.
var errPoolerDrained = errors.New("pooler marked as DRAINED: replication cannot be established")

// FixReplicationAction handles replication configuration and repair for replicas.
//
// This action addresses the following problem codes:
//   - ProblemReplicaNotReplicating: Replication is not configured at all
//
// Future problem codes (TODO):
//   - ProblemReplicaWrongPrimary: Replica is pointing to a stale/wrong primary
//   - ProblemReplicaLagging: Replication is configured but lag is excessive
//
// The action:
//   - Re-verifies the problem still exists (fresh RPC calls)
//   - Identifies the current primary from topology
//   - Configures the replica's primary_conninfo to point to the primary
//   - Verifies replication is streaming
//
// Cohort membership (adding/removing the replica to the primary's
// synchronous standby list) is managed separately by ReconcileCohortAction.
//
// Idempotency:
// This action is fully idempotent. If multiple multiorch instances race to fix
// the same problem, the end result will be identical. The underlying RPC
// operations (SetPrimaryConnInfo) are implemented as idempotent operations
// at the pooler level and serialized by action locks on the poolers, so
// concurrent calls are safe and produce the same final state.

// Default polling parameters for verifyReplicationStarted.
const (
	DefaultVerifyMaxAttempts  = 10
	DefaultVerifyPollInterval = 500 * time.Millisecond
)

type FixReplicationAction struct {
	config      *config.Config
	rpcClient   rpcclient.MultiPoolerClient
	poolerStore *store.PoolerStore
	topoStore   topoclient.Store
	logger      *slog.Logger

	// Polling parameters for verifyReplicationStarted.
	verifyMaxAttempts  int
	verifyPollInterval time.Duration
}

// NewFixReplicationAction creates a new fix replication action.
func NewFixReplicationAction(
	cfg *config.Config,
	rpcClient rpcclient.MultiPoolerClient,
	poolerStore *store.PoolerStore,
	topoStore topoclient.Store,
	logger *slog.Logger,
) *FixReplicationAction {
	_ = "STUB: not implemented"
	return nil
}

// Execute performs replication fix for a replica that is not replicating.
func (a *FixReplicationAction) Execute(ctx context.Context, problem types.Problem) error {
	_ = "STUB: not implemented"
	return nil
}

// Find the affected replica

// Get all poolers in this shard to find the primary

// Find a healthy primary in the shard

// Re-verify the problem still exists

// Dispatch to the appropriate fix based on the problem

// TODO: Future problem codes to handle
// case types.ProblemReplicaWrongPrimary:
//     return a.fixWrongPrimary(ctx, replica, primary, currentStatus)
// case types.ProblemReplicaLagging:
//     return a.fixReplicaLagging(ctx, replica, primary, currentStatus)
// case types.ProblemReplicaMisconfigured:
//     return a.fixMisconfigured(ctx, replica, primary, currentStatus)

// fixNotReplicating handles the case where replication is not set up at all.
// This is the most basic case: the replica has no primary_conninfo configured.
// It checks for timeline divergence first before starting replication to avoid
// the race where PostgreSQL starts, connects to primary, and updates its timeline.
func (a *FixReplicationAction) fixNotReplicating(
	ctx context.Context,
	replica *multiorchdatapb.PoolerHealthState,
	primary *multiorchdatapb.PoolerHealthState,
) (retErr error) {
	_ = "STUB: not implemented"
	return nil
}

// Use the term numbers already carried in the health state rather than
// making extra ConsensusStatus RPCs. Both values come from StatusResponse
// via the health stream, so they reflect the same data we would get from
// a fresh RPC at the time the problem was detected.
//
// We take max(primaryTerm, replicaTerm) because after a failover the
// replica may have accepted a higher term (from BeginTerm) than the
// newly-elected primary has seen yet. validateAndUpdateTerm rejects
// requests whose CurrentTerm is below the local term, so using the
// maximum satisfies both nodes. A higher term is safe: the primary
// accepts it and advances its own term to match.

// Configure primary_conninfo on the replica.

// Verify replication started

// Re-check the primary's latest health-stream state before running pg_rewind.
// pg_rewind stops the replica's postgres before contacting the source; if the
// primary postgres is no longer running the stop will leave two nodes down.
// Return an error for retry — the next cycle will detect PrimaryIsDead.

// pg_rewind was not feasible; pooler marked as DRAINED.
// No point verifying replication — treat as resolved.

// Re-verify replication after rewind. RewindToSource restarts
// PostgreSQL as a standby, and primary_conninfo in
// postgresql.auto.conf is preserved (pg_rewind doesn't touch it).

// Cohort membership (adding the replica to synchronous_standby_names) is
// managed by ReconcileCohortAction separately. By the time this action
// returns, the replica is replicating; the cohort analyzer will pick it up
// on the next cycle and propose adding it to the cohort.

// tryPgRewind attempts to repair a replica using pg_rewind.
// RewindToSource will:
// 1. Stop postgres
// 2. Check if rewind is needed (dry-run)
// 3. Run actual rewind if needed
// 4. Start postgres
// If pg_rewind is not feasible (missing WAL), it marks the pooler as DRAINED.
func (a *FixReplicationAction) tryPgRewind(
	ctx context.Context,
	primary *multiorchdatapb.PoolerHealthState,
	replica *multiorchdatapb.PoolerHealthState,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Call RewindToSource - it handles the entire flow atomically

// RPC failure (e.g. primary postgres unreachable) is transient — do not
// drain the pooler. Return an error so the next recovery cycle retries.

// verifyProblemExists re-checks whether the replication problem still exists.
// Returns true if the problem persists, false if already resolved.
func (a *FixReplicationAction) verifyProblemExists(
	ctx context.Context,
	replica *multiorchdatapb.PoolerHealthState,
	primary *multiorchdatapb.PoolerHealthState,
	problemCode types.ProblemCode,
) (bool, *multipoolermanagerdatapb.StandbyReplicationStatus, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

// verifyReplicaNotReplicating checks if the replica still has no replication configured.
func (a *FixReplicationAction) verifyReplicaNotReplicating(
	ctx context.Context,
	replica *multiorchdatapb.PoolerHealthState,
	primary *multiorchdatapb.PoolerHealthState,
) (bool, *multipoolermanagerdatapb.StandbyReplicationStatus, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

// No status means we can't determine state, assume problem exists

// Check if primary_conninfo is configured

// Check if pointing to the right primary

// TODO: Do we need to verify timeline_id matches the primary's timeline?

// Wrong primary - this would be ProblemReplicaWrongPrimary

// Check if WAL replay is paused (might need to resume)

// getReplicationStatus gets the current replication status from the replica.
func (a *FixReplicationAction) getReplicationStatus(
	ctx context.Context,
	replica *multiorchdatapb.PoolerHealthState,
) (*multipoolermanagerdatapb.StandbyReplicationStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// verifyReplicationStarted checks that replication is actively streaming.
// It polls a few times to allow the WAL receiver to connect.
func (a *FixReplicationAction) verifyReplicationStarted(ctx context.Context, replica *multiorchdatapb.PoolerHealthState) error {
	_ = "STUB: not implemented"
	return nil
}

// Check WAL receiver status first - this is the live connection state

// Also verify we have a receive LSN (sanity check)

// RecoveryAction interface implementation

func (a *FixReplicationAction) RequiresHealthyLeader() bool {
	_ = "STUB: not implemented"
	// Cannot fix replica replication without a healthy primary
	return false
}

func (a *FixReplicationAction) Metadata() types.RecoveryMetadata {
	_ = "STUB: not implemented"
	return *new(types.RecoveryMetadata)
}

func (a *FixReplicationAction) Priority() types.Priority {
	_ = "STUB: not implemented"
	return *new(types.Priority)
}

func (a *FixReplicationAction) GracePeriod() *types.GracePeriodConfig {
	_ = "STUB: not implemented"
	// No grace period needed, execute immediately
	return nil
}

// markPoolerDrained marks a pooler as DRAINED in the topology.
func (a *FixReplicationAction) markPoolerDrained(ctx context.Context, pooler *multiorchdatapb.PoolerHealthState) (retErr error) {
	_ = "STUB: not implemented"
	return nil
}

// =============================================================================
// TODO: Future replication problem handlers
// =============================================================================
//
// The following are replication problems we should handle in future PRs:
//
// ProblemReplicaWrongPrimary
//    - Replica is connected to a stale primary (e.g., after failover)
//    - Fix: Update primary_conninfo to point to new primary, restart streaming
//    - Consider: We need to handle timeline changes.
//
// ProblemReplicaLagging
//    - Replication is working but lag exceeds threshold
//    - Causes to investigate:
//      a) Network congestion between primary and replica
//      b) Replica CPU/IO saturation (can't keep up with replay)
//      c) Long-running queries on replica blocking replay
//      d) Checkpoint/vacuum activity on primary generating excessive WAL
//      e) Synchronous replication bottleneck
//    - Fix: Depends on root cause; short-term we might not fix them, automatically
//           should understand why replication is broken.
//
// ProblemWalReceiverCrashing
//    - WAL receiver process repeatedly crashing
//    - Causes: Bad WAL segment, memory issues, bugs
//    - Fix: May need to skip corrupted WAL or re-clone
//
// ProblemReplicaSlotMissing
//    - NOTE: We are not creating a replication slot right now, so we might need to revisit this.
//    - Replication slot on primary was dropped
//    - Symptoms: Replica can't stream, gets "replication slot does not exist"
//    - Fix: Create new slot, may need to re-clone if WAL recycled
//
// ProblemSynchronousStandbyMisconfigured
//    - synchronous_standby_names doesn't match actual standbys
//    - Symptoms: Primary waiting indefinitely for sync confirmation
//    - Fix: Update synchronous_standby_names to match reality
