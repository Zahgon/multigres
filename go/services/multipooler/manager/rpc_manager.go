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
	"time"

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	multipoolermanagerdatapb "github.com/multigres/multigres/go/pb/multipoolermanagerdata"
)

// broadcastHealth broadcasts the current health state to all subscribers.
//
// This should be called whenever there is a state change that clients should be
// aware of (e.g., PostgreSQL availability, replication status, etc.). Clients
// will receive the latest health snapshot immediately if they are connected, or
// upon their next connection if they are not currently connected.
func (pm *MultiPoolerManager) broadcastHealth() { _ = "STUB: not implemented"; return }

// WaitForLSN waits for PostgreSQL server to reach a specific LSN position
func (pm *MultiPoolerManager) WaitForLSN(ctx context.Context, targetLsn string) error {
	_ = "STUB: not implemented"
	return nil
}

// Check REPLICA guardrails (pooler type and recovery mode)

// Wait for the standby to replay WAL up to the target LSN
// We use a polling approach to check if the replay LSN has reached the target

// Check if the standby has replayed up to the target LSN

// SetPrimaryConnInfo sets the primary connection info for a standby server
func (pm *MultiPoolerManager) SetPrimaryConnInfo(ctx context.Context, primary *clustermetadatapb.MultiPooler, stopReplicationBefore, startReplicationAfter bool, currentTerm int64, force bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Acquire the action lock to ensure only one mutation runs at a time

// Validate and update consensus term following consensus rules

// Extract host and port from the MultiPooler (nil means clear the config)

// Record the (rule, primary) tuple so ReplicationPrimary stays the canonical
// source for "who is the primary now."
//
// TODO: this entire SetPrimaryConnInfo RPC goes away with the legacy
// consensus flow; the new flow's SetTermPrimary already carries a real
// ShardRule. Until then, build a minimal synthetic rule from
// (currentTerm, primary.Id). No consumer of rp.Rule reads cohort_members
// or durability_policy today, so the stub is sufficient for the
// rule-number ordering and self-as-leader checks that downstream code
// performs.

// Call the locked version that assumes action lock is already held

// Push an immediate health snapshot so orchestrators learn about the new
// replication configuration (e.g., cleared primary_conninfo) without waiting
// for the next 30-second heartbeat.

// setPrimaryConnInfoLocked sets the primary connection info for a standby server.
// This function assumes the action lock is already held by the caller.
//
// Refuses with FAILED_PRECONDITION when StopReplication previously cleared
// primary_conninfo and set the manual-stop flag — every conninfo writer
// (orch's FixReplication via SetPrimaryConnInfo, SetTermPrimary's standby
// branch, and the postgres-monitor self-heal) funnels through here, so this
// single check is what keeps the admin pause honored against routine
// reconciliation. Use StartReplication to clear the flag before rewriting
// conninfo. demoteStalePrimaryLocked clears the flag itself before reaching
// this point — a stale-primary detection is an escalated event that
// supersedes an older admin pause.
func (pm *MultiPoolerManager) setPrimaryConnInfoLocked(ctx context.Context, host string, port int32, stopReplicationBefore, startReplicationAfter bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Guardrail: Check if the PostgreSQL instance is in recovery (standby mode)

// Optionally stop replication before making changes

// Build primary_conninfo connection string
// Format: host=<host> port=<port> user=<user> application_name=<name> [passfile=<path>]
// The heartbeat_interval is converted to keepalives_interval/keepalives_idle.
// passfile points libpq at the pgpass file written at manager startup so the
// standby can authenticate to the primary via SCRAM without embedding the
// password in postgresql.auto.conf. It is omitted when pgpassPath is unset
// (early startup or unit tests that bypass loadMultiPoolerFromTopo).

// Set primary_conninfo using ALTER SYSTEM

// Reload PostgreSQL configuration to apply changes

// Optionally start replication after making changes.
// Note: If replication was already running when calling SetPrimaryConnInfo,
// even if we don't set startReplicationAfter to true, replication will be running.

// Wait for database to be available after restart

// StartReplication starts WAL replay on standby (calls pg_wal_replay_resume).
// As the counterpart to StopReplication, it also clears any in-memory
// "manually stopped" signal that StopReplication may have set. Clearing
// flips this pooler's published CohortEligibility back to ELIGIBLE and
// re-enables the postgres-monitor self-heal of primary_conninfo, which in
// turn lets routine reconciliation re-establish the WAL receiver.
//
// StartReplication itself does not rewrite primary_conninfo — that happens
// via the monitor's self-heal (or via orch's FixReplicationAction) once
// eligibility flips.
func (pm *MultiPoolerManager) StartReplication(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Acquire the action lock to ensure only one mutation runs at a time

// Check REPLICA guardrails (pooler type and recovery mode)

// Resume WAL replay on the standby

// StopReplication stops replication based on the specified mode
func (pm *MultiPoolerManager) StopReplication(ctx context.Context, mode multipoolermanagerdatapb.ReplicationPauseMode, wait bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Acquire the action lock to ensure only one mutation runs at a time

// Check REPLICA guardrails (pooler type and recovery mode)

// Modes that clear primary_conninfo are an explicit admin signal that
// replication should stay stopped. Mark this so the postgres monitor
// does not "self-heal" the cleared conninfo back to the recorded primary,
// and so this pooler publishes COHORT_ELIGIBILITY_INELIGIBLE while
// stopped. Cleared the next time something re-establishes the primary
// link (SetTermPrimary / SetPrimaryConnInfo / demoteStalePrimaryLocked).

// StandbyReplicationStatus gets the current replication status of the standby
func (pm *MultiPoolerManager) StandbyReplicationStatus(ctx context.Context) (*multipoolermanagerdatapb.StandbyReplicationStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check REPLICA guardrails (pooler type and recovery mode)

// Query all replication status fields

// Status gets unified status that works for both PRIMARY and REPLICA poolers.
// This RPC works even when the database connection is unavailable - fields that require
// database access will be nil/empty in that case. This allows callers to always get
// initialization status without needing a separate RPC.
func (pm *MultiPoolerManager) Status(ctx context.Context) (*multipoolermanagerdatapb.StatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get WAL position (ignore errors, just return empty string)

// Get cohort members from the current rule (best-effort).

// Try to get detailed status based on PostgreSQL role

// Can't determine role - return what we have

// Populate role-specific status

// Acting as primary - get primary status (skip guardrails since we already checked isPrimary)

// Return partial status instead of error

// Acting as standby - get replication status (skip guardrails since we already checked isPrimary)

// Return partial status instead of error

// ResetReplication resets the standby's connection to its primary by clearing primary_conninfo
// and reloading PostgreSQL configuration. This effectively disconnects the replica from the primary
// and prevents it from acknowledging commits, making it unavailable for synchronous replication
// until reconfigured.
func (pm *MultiPoolerManager) ResetReplication(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Acquire the action lock to ensure only one mutation runs at a time

// Check REPLICA guardrails (pooler type and recovery mode)

// Pause the receiver (clear primary_conninfo) and wait for disconnect
/* wait */

// ConfigureSynchronousReplication configures PostgreSQL synchronous replication settings
func (pm *MultiPoolerManager) ConfigureSynchronousReplication(ctx context.Context, synchronousCommit multipoolermanagerdatapb.SynchronousCommitLevel, synchronousMethod multipoolermanagerdatapb.SynchronousMethod, numSync int32, standbyIDs []*clustermetadatapb.ID, reloadConfig bool, force bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Acquire the action lock to ensure only one mutation runs at a time

// configureSynchronousReplicationLocked configures PostgreSQL synchronous replication settings.
// The caller MUST already hold the action lock.
func (pm *MultiPoolerManager) configureSynchronousReplicationLocked(ctx context.Context, synchronousCommit multipoolermanagerdatapb.SynchronousCommitLevel, synchronousMethod multipoolermanagerdatapb.SynchronousMethod, numSync int32, standbyIDs []*clustermetadatapb.ID, reloadConfig bool, force bool) error {
	_ = "STUB: not implemented"
	// Validate input parameters
	return nil
}

// Check PRIMARY guardrails (pooler type and non-recovery mode)

// Insert history before applying GUCs.
// Rationale: we want to ensure that a new cohort is advertised
// before this primary can accept ACKs from it.
// This is for safe replica joining of the cluster.
// It will ensure multiorch can discover the new cohort during a failure.

// Set synchronous_commit level

// Build and set synchronous_standby_names

// Reload configuration if requested

// UpdateConsensusRule updates PostgreSQL synchronous_standby_names by adding
// or removing members. It is idempotent and only valid when synchronous
// replication is already configured.
//
// expectedOutgoingRule provides compare-and-swap semantics: the operation
// proceeds only if this pooler's current recorded rule matches the given
// RuleNumber. If they differ (the caller's view is stale), the operation
// fails — the caller should re-read state and retry.
func (pm *MultiPoolerManager) UpdateConsensusRule(ctx context.Context, operation multipoolermanagerdatapb.CohortUpdateOperation, standbyIDs []*clustermetadatapb.ID, expectedOutgoingRule *clustermetadatapb.RuleNumber, coordinatorID *clustermetadatapb.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate operation

// Validate standby IDs using the shared validation function

// Pre-compute history fields before acquiring the lock.

// Check PRIMARY guardrails (pooler type and non-recovery mode)

// === Parse Current Configuration ===

// Read current cohort from the rule store (authoritative source of truth).

// Check if synchronous replication is configured

// Convert current cohort IDs to pooler IDs for set operations.

// === Apply Operation ===

// Validate that the final list is not empty

// Check if there are any changes (idempotent).

// Insert history before applying GUCs
// Rationale: we want to ensure that a new cohort is advertised
// before this primary can accept ACKs from it.
// This is for safe replica joining of the cluster.
// It will ensure multiorch can discover the new cohort during a failure.

// The new rule inherits the expected coordinator term — we're not
// changing the leader, just amending its cohort. The rule store assigns
// a fresh leader_subterm.

// Push an immediate health snapshot so orchestrators learn about the changed
// synchronous standby list without waiting for the next 30-second heartbeat.

// getPrimaryStatusInternal gets primary status without guardrail checks.
// Called by Status() which has already verified the PostgreSQL role.
func (pm *MultiPoolerManager) getPrimaryStatusInternal(ctx context.Context) (*multipoolermanagerdatapb.PrimaryStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get current LSN

// Get connected followers from pg_stat_replication

// Get synchronous replication configuration

// getStandbyStatusInternal gets standby replication status without guardrail checks.
// Called by Status() which has already verified the PostgreSQL role.
func (pm *MultiPoolerManager) getStandbyStatusInternal(ctx context.Context) (*multipoolermanagerdatapb.StandbyReplicationStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PrimaryStatus gets the status of the leader server
func (pm *MultiPoolerManager) PrimaryStatus(ctx context.Context) (*multipoolermanagerdatapb.PrimaryStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check PRIMARY guardrails (pooler type and non-recovery mode)

// PrimaryPosition gets the current LSN position of the leader
func (pm *MultiPoolerManager) PrimaryPosition(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Check PRIMARY guardrails (pooler type and non-recovery mode)

// Get current primary LSN position

// StopReplicationAndGetStatus stops PostgreSQL replication (replay and/or receiver based on mode) and returns the status
func (pm *MultiPoolerManager) StopReplicationAndGetStatus(ctx context.Context, mode multipoolermanagerdatapb.ReplicationPauseMode, wait bool) (*multipoolermanagerdatapb.StandbyReplicationStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Acquire the action lock to ensure only one mutation runs at a time

// Check REPLICA guardrails (pooler type and recovery mode)

// changeTypeLocked updates the pooler type without acquiring the action lock.
// The caller MUST already hold the action lock.
func (pm *MultiPoolerManager) changeTypeLocked(ctx context.Context, poolerType clustermetadatapb.PoolerType) error {
	_ = "STUB: not implemented"
	return nil
}

// Use the serving state manager to transition components and update the multipooler record.
// The serving status stays SERVING during type changes (the node remains available).

// Notify the topology publisher of the new state. The write to etcd happens
// asynchronously so that a temporarily unreachable etcd does not block type changes.

// ChangeType changes the pooler type (PRIMARY/REPLICA)
func (pm *MultiPoolerManager) ChangeType(ctx context.Context, poolerType string) error {
	_ = "STUB: not implemented"
	return nil
}

// Acquire the action lock to ensure only one mutation runs at a time

// Validate pooler type

// TODO: For now allow to change type to PRIMARY, this is to make it easier
// to perform tests while we are still developing HA. Once, we have multiorch
// fully implemented, we shouldn't allow to change the type to Primary.
// This would happen organically as part of Promote workflow.

// Call the locked version

// EmergencyDemote demotes the current primary server
// This can be called for any of the following use cases:
// - By orchestrator when fixing a broken shard.
// - When performing a Planned demotion.
// - When receiving a SIGTERM and the pooler needs to shutdown.
func (pm *MultiPoolerManager) EmergencyDemote(ctx context.Context, consensusTerm int64, drainTimeout time.Duration, force bool) (_ *multipoolermanagerdatapb.EmergencyDemoteResponse, retErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Acquire the action lock to ensure only one mutation runs at a time

// Validate the term but DON'T update yet. We only update the term AFTER
// successful demotion to avoid a race where a failed demote (e.g., postgres
// not ready) updates the term, causing subsequent detection to see equal
// terms and skip demotion.

// Perform the actual demotion

// Only update term AFTER successful demotion
// This ensures the stale primary keeps its lower term until it's actually demoted,
// allowing subsequent detection to continue flagging it as stale.

// Log but don't fail - demotion succeeded, term update is secondary

// emergencyDemoteLocked performs the core demotion logic.
// REQUIRES: action lock must already be held by the caller.
// This is used for emergency demote operations.
// We won't try to perform a graceful switchover in this case.
// We will drain this pooler and stop postgres.
// This should only be called during ungraceful shutdown.
// MultiOrch will try to contact all nodes in the cohort.
// In the case that the dead primary received the RPC, it should just
// shut down itself.
func (pm *MultiPoolerManager) emergencyDemoteLocked(ctx context.Context, consensusTerm int64, drainTimeout time.Duration) (*multipoolermanagerdatapb.EmergencyDemoteResponse, error) {
	_ = "STUB: not implemented"
	// Verify action lock is held
	return nil, nil
}

// === Validation & State Check ===

// Guard rail: Demote can only be called on a PRIMARY

// Check current demotion state

// If everything is already complete, return early (fully idempotent)

// Transition to NOT_SERVING — rejects all queries and stops heartbeat.
// This ensures no new writes arrive while we drain existing connections.

// Drain write connections

// Terminate Remaining Write Connections

// Log but don't fail - connections will eventually timeout

// Capture State & Make PostgreSQL Read-Only

// Signal voluntary resignation so the coordinator can trigger an immediate
// election without waiting for a heartbeat timeout. Use this node's own
// primary_term (not the incoming consensusTerm) so the coordinator can
// correlate the signal with the term at which this node was elected.
// setResignedLeaderAtTerm broadcasts internally on a change so multiorch
// sees leadership_status.REQUESTING_DEMOTION before the next periodic
// health stream interval fires.

// Restart PostgreSQL as standby. Unlike the old stop-only path, this keeps
// the node in the cluster as a replication target, avoiding timeline divergence
// in most cases. The coordinator still uses pg_rewind for nodes that diverged.

// Suppress the postgres monitor until a rewind completes; the monitor would
// otherwise restart postgres on this demoted node.

// UndoDemote undoes a demotion
func (pm *MultiPoolerManager) UndoDemote(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Acquire the action lock to ensure only one mutation runs at a time

// DemoteStalePrimary demotes a stale primary that came back online after failover.
// This is a complete operation that:
// 1. Stops postgres if running
// 2. Runs pg_rewind to sync with the correct primary
// 3. Clears sync replication config
// 4. Restarts as standby
// 5. Updates topology to REPLICA
func (pm *MultiPoolerManager) DemoteStalePrimary(
	ctx context.Context,
	source *clustermetadatapb.MultiPooler,
	consensusTerm int64,
	force bool,
) (*multipoolermanagerdatapb.DemoteStalePrimaryResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate source pooler

// Acquire the action lock to ensure only one mutation runs at a time

// Check if already demoted by consulting the current rule. Only short-circuit
// when we can positively confirm this node is no longer the primary (err ==
// nil AND term == 0). On any error — e.g. postgres is down after a crash —
// we fall through and run the idempotent demotion flow, because we cannot
// tell the difference between "already demoted" and "primary with unreachable
// postgres" from local state alone.

// Return success with rewind_performed=false since node is already in correct state

// Validate the term

// Project the legacy MultiPooler source into the PoolerAddress shape the
// helper takes. demoteStalePrimaryLocked only needs id/host/postgres_port.

// TODO: this RPC will be removed in the new consensus flow; the new flow's
// SetTermPrimary passes the real ShardRule it received. Until then, build
// a synthetic rule from (term, leader_id). No consumer of rp.Rule reads
// cohort_members or durability_policy today.

// Bump the local revoked_below_term to match the new primary's term. This
// is the explicit RPC, where the caller (typically multiorch) intends to
// commit this node to the new term. SetTermPrimary deliberately does not do this
// — see demoteStalePrimaryLocked's doc.

// demoteStalePrimaryLocked performs the postgres + topology work to convert a
// stale primary into a standby pointing at source. The action lock must be held
// by the caller. Idempotency checks and term validation are the caller's
// responsibility.
//
// The helper does not touch term_revocation: revocations are authored by
// coordinators via Recruit/AcceptRevocation, not by side effects of demotion.
// Callers that want to record the new term (the explicit DemoteStalePrimary
// RPC) call updateTermIfNewer themselves after the demotion succeeds; SetTermPrimary
// deliberately does not, because an SetTermPrimary is a notification, not a revoke.
//
// Sequence: stop postgres -> pg_rewind -> fix pgbackrest paths -> restart as
// standby -> reset sync replication -> set primary_conninfo -> report leader
// observation -> read final LSN -> flip topology type to REPLICA.
func (pm *MultiPoolerManager) demoteStalePrimaryLocked(
	ctx context.Context,
	source *clustermetadatapb.PoolerAddress,
	rule *clustermetadatapb.ShardRule,
) (rewindPerformed bool, finalLSN string, err error) {
	_ = "STUB: not implemented"
	return false, "", nil
}

// Fix pgbackrest paths in postgresql.auto.conf after pg_rewind
// The config may have wrong paths copied from another pooler during initial setup

// Record the (rule, primary) tuple so ReplicationPrimary stays the canonical
// source for "who is the primary now." The new flow's SetTermPrimary
// passes the real rule; the legacy DemoteStalePrimary RPC synthesises one
// from its term parameter.

// Call the locked version directly since we already hold the action lock
// (calling SetPrimaryConnInfo would deadlock trying to acquire the same lock)

// Report the new primary (source) so the gateway can use this observation.

// Update topology to REPLICA

// Promote promotes a standby to primary
// This is called during the Propagate stage of generalized consensus to safely
// transition a standby to primary and reconfigure replication.
// This operation is fully idempotent - it checks what steps are already complete
// and only executes the missing steps.
func (pm *MultiPoolerManager) Promote(ctx context.Context, consensusTerm int64, expectedLSN string, syncReplicationConfig *multipoolermanagerdatapb.ConfigureSynchronousReplicationRequest, force bool, reason string, coordinatorID *clustermetadatapb.ID, cohortMemberIDs, acceptedMemberIDs []*clustermetadatapb.ID) (*multipoolermanagerdatapb.PromoteResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Acquire the action lock to ensure only one mutation runs at a time

// Validation & Readiness

// Validate term - strict equality, no automatic updates

// Check current promotion state to determine what needs to be done

// Guard rail: Check topology type and validate state consistency
// If topology is PRIMARY, verify everything is in expected state (idempotency check)
// If topology is REPLICA, proceed with promotion

// Topology shows PRIMARY - validate that everything is consistent

// Check if everything is in expected state

// Everything is consistent and complete - idempotent success

// Inconsistent state detected

// Without force flag, require manual intervention

// If PostgreSQL is not promoted yet, validate expected LSN before promotion

// Execute missing steps

// Promote PostgreSQL if needed

// Configure sync replication if needed

// Get final LSN position

// Clear any outstanding resignation signal now that the coordinator has
// explicitly re-promoted us at a new term. A higher primary_term implicitly
// invalidates the old signal, but clearing eagerly avoids a window where
// a stale REQUESTING_DEMOTION is still published in StatusResponse.

// Write rule history record - this validates that sync replication is working.
// If this fails (typically due to timeout waiting for standby acknowledgment), we fail
// the promotion. It's better to have no primary than one that can't satisfy durability.

// Update topology and notify all components (best-effort, don't fail promotion)

// RewindToSource performs pg_rewind to synchronize this server with a source.
// This operation:
// 1. Stops PostgreSQL
// 2. Runs pg_rewind --dry-run to check if rewind is needed
// 3. If needed, runs actual pg_rewind
// 4. Starts PostgreSQL
func (pm *MultiPoolerManager) RewindToSource(ctx context.Context, source *clustermetadatapb.MultiPooler) (*multipoolermanagerdatapb.RewindToSourceResponse, error) {
	_ = "STUB: not implemented"
	// Check if multipooler is ready
	return nil, nil
}

// Validate source pooler

// Acquire the action lock to ensure only one mutation runs at a time

// Check if pgctld client is available

// Pause manager and stop PostgreSQL for pg_rewind
// resume() is called explicitly after PostgreSQL restart, and also via defer for cleanup

// Safety net: ensure manager is resumed even if errors occur

// Run pg_rewind --dry-run to check if rewind is needed

// Check if rewind is needed by parsing output

// Servers have diverged - run actual pg_rewind

// Step 4: Start PostgreSQL as standby
// Use Restart with as_standby=true to create standby.signal and start postgres
// Note: postgres is already stopped, so the stop phase will be a no-op

// Resume manager now that PostgreSQL is running

// Wait for database connection

// Rewind succeeded: allow the monitor to resume normal operation.

// SetPostgresRestartsEnabled enables or disables automatic PostgreSQL restarts by the monitor.
// When disabled, the monitor continues to run and detect problems but will not auto-restart
// a stopped PostgreSQL instance. Used by tests and demos during controlled failovers.
func (pm *MultiPoolerManager) SetPostgresRestartsEnabled(ctx context.Context, req *multipoolermanagerdatapb.SetPostgresRestartsEnabledRequest) (*multipoolermanagerdatapb.SetPostgresRestartsEnabledResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ====================================================================================
// Helper methods for DemoteStalePrimary
// ====================================================================================

// stopPostgresIfRunning stops postgres if it's currently running.
func (pm *MultiPoolerManager) stopPostgresIfRunning(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Close ONLY connection pools to release database connections.
// This allows postgres to stop cleanly without waiting for connections,
// but keeps the manager operational for subsequent operations.

// Stop postgres (no-op if already stopped)

// Treat "already stopped" errors as success to make this truly idempotent.
// This handles race conditions where postgres was stopped between our check and stop call.

// runPgRewind runs pg_rewind to sync with source.
// Returns true if rewind was performed, false if not needed.
func (pm *MultiPoolerManager) runPgRewind(ctx context.Context, sourceHost string, sourcePort int32) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Get application name for replication connection

// Dry-run to check if rewind is needed

// Check if servers diverged

// No divergence: the node is already in sync with the source. The rewind is
// effectively complete; clear the flag so the monitor resumes.

// fixPgBackRestPaths fixes the pgbackrest paths in postgresql.auto.conf
// After pg_rewind, the restore_command and archive_command may have paths from another pooler
// This function updates them to point to the current pooler's directories
func (pm *MultiPoolerManager) fixPgBackRestPaths(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Read the file

// Replace all occurrences of old pooler paths with current pooler paths
// We need to fix: --config, --lock-path, --log-path, --pg1-path
// These paths follow the pattern: /some/path/pooler-X/data/...
// We want to replace them with: /some/path/pooler-current/data/...

// Extract current pooler dir path pattern
// poolerDir is like: /tmp/test_12345/pooler-1/data
// We want to match patterns like: /tmp/test_12345/pooler-X/data
// Go up two levels to get base directory

// Use regex to replace pooler-X paths with current pooler paths
// Pattern matches: /path/to/pooler-<anything>/data

// Write the file back

// restartAsStandbyAfterRewind restarts postgres as standby after rewind.
func (pm *MultiPoolerManager) restartAsStandbyAfterRewind(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Use existing restartPostgresAsStandby with a state that indicates postgres is not running
	return nil
}

// Postgres was stopped, not in standby mode yet
