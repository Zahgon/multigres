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
	consensusdatapb "github.com/multigres/multigres/go/pb/consensusdata"
)

// BeginTerm handles coordinator requests during leader appointments.
// It consists of two phases:
//
// 1. Term Acceptance: Accept the new term based on consensus rules
//   - Term must be >= current term
//   - Cannot accept different coordinator for same term
//   - Atomically update term and accept candidate
//
// 2. Action Execution: Execute the specified action after term acceptance
//   - NO_ACTION: Do nothing
//   - REVOKE: Demote primary or pause standby replication to revoke old term
func (pm *MultiPoolerManager) BeginTerm(ctx context.Context, req *consensusdatapb.BeginTermRequest) (_ *consensusdatapb.BeginTermResponse, retErr error) {
	_ = "STUB: not implemented"
	// Acquire the action lock to ensure only one consensus operation runs at a time
	// This prevents split-brain acceptance and ensures term updates are serialized
	return nil, nil
}

// Log the action type for observability

// Validate action

// Valid action

// Valid action

// ========================================================================
// Term Acceptance (Consensus Rules)
// ========================================================================

// Get current term for response

// Atomically update term and accept candidate
// This handles all consensus rules: term validation, duplicate check, etc.

// Term not accepted - return rejection with consensus status so the coordinator
// learns this pooler's current state even from a rejection.

// Determine revoked role before executing any action (needed for event)

// ========================================================================
// Action Execution
// ========================================================================

// Term was already accepted and persisted above, so we must return
// the response with accepted=true AND the error. This tells the coordinator:
// 1. The term was accepted (response.Accepted = true)
// 2. The revoke action failed (error != nil)

// Should never reach here due to validation above

// executeRevoke executes the REVOKE action by demoting primary or pausing standby replication.
// This is called after the term has been accepted.
func (pm *MultiPoolerManager) executeRevoke(ctx context.Context, term int64, response *consensusdatapb.BeginTermResponse) error {
	_ = "STUB: not implemented"
	// CRITICAL: Must be able to reach Postgres to execute revoke
	return nil
}

// Revoke primary: demote
// TODO: Implement graceful (non-emergency) demote for planned failovers.
// This emergency demote path will remain for BeginTerm REVOKE actions.

// Revoke standby: stop receiver and wait for replay to catch up

// Stop WAL receiver and wait for it to fully disconnect

/* wait */

// Wait for replay to finish processing all WAL that is on disk

// Always capture timeline ID after WAL positions are frozen.
// Retained for observability only; does not affect candidate selection.

// Capture the highest consensus term replicated to this node, plus the cohort
// that was active at that point. The coordinator uses leadership_term as
// the primary criterion: a node that has seen a higher term has applied more
// of the agreed WAL history (the history write uses RemoteOperationTimeout,
// so sync standbys are guaranteed to have acknowledged it).
//
// observePosition also warms the ruleStore cache, allowing getCachedConsensusStatus
// below to read the position without an additional postgres round-trip.

// Capture consensus status after WAL positions are frozen (post-revoke snapshot).
// Uses the cached position warmed by observePosition above — no extra DB round-trip.

// buildConsensusStatus constructs a ConsensusStatus from a pre-resolved revocation,
// position, and the highest-known RPC-told (rule, primary). Any argument may be
// nil; the corresponding field is left unset. Never performs I/O.
//
// The published HighestKnownRule reflects best knowledge from any source:
//   - rule: max of the observed position's rule and the rule from the most
//     recent SetTermPrimary/Propose RPC.
//   - primary: the contact info from the most recent SetTermPrimary/Propose, since
//     observePosition cannot carry it.
//
// Result is left nil only when neither source has any information.
func buildConsensusStatus(id *clustermetadatapb.ID, revocation *clustermetadatapb.TermRevocation, pos *clustermetadatapb.PoolerPosition, replicationPrimary *clustermetadatapb.ReplicationPrimary) *clustermetadatapb.ConsensusStatus {
	_ = "STUB: not implemented"
	return nil
}

// buildStatusReplicationPrimary returns the HighestKnownRule to publish given the most
// recent observed position and the most recent rule+primary heard via RPC.
// See buildConsensusStatus for the merge semantics.
func buildStatusReplicationPrimary(pos *clustermetadatapb.PoolerPosition, replicationPrimary *clustermetadatapb.ReplicationPrimary) *clustermetadatapb.ReplicationPrimary {
	_ = "STUB: not implemented"
	return nil
}

// getConsensusStatus builds a ConsensusStatus snapshot while holding the action lock.
// Callers must already hold the action lock (i.e. this is called from BeginTerm or
// executeRevoke). Uses a consistent disk read for the term and a fresh postgres query
// for the current position.
//
// Returns an error if postgres is unreachable, since a partial status (term revocation
// without current_position) could mislead callers about this pooler's rule position.
func (pm *MultiPoolerManager) getConsensusStatus(ctx context.Context) (*clustermetadatapb.ConsensusStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getCachedConsensusStatus builds a ConsensusStatus using the in-memory term cache and
// the ruleStore's cached position. Never queries postgres or disk.
//
// The action lock must be held by the caller, which prevents concurrent term updates.
// Returns nil if no position has been cached yet (i.e. observePosition or updateRule
// has never been called).
func (pm *MultiPoolerManager) getCachedConsensusStatus() (*clustermetadatapb.ConsensusStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getInconsistentConsensusStatus builds a ConsensusStatus without holding the action lock.
// Like GetInconsistentTerm, it may observe a partially-updated state during a concurrent
// BeginTerm, so it is suitable for observability (StatusResponse, health monitors) but not
// for decisions that require a consistent view.
//
// Falls back to the ruleStore's cached position when postgres is unreachable, so
// that callers can still derive the last-known primary term (e.g. for stale-
// primary detection) after postgres has crashed.
func (pm *MultiPoolerManager) getInconsistentConsensusStatus(ctx context.Context) (*clustermetadatapb.ConsensusStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Postgres is unreachable — fall back to the last observed position
// cached in memory. May be stale, but preserves visibility into the
// most recent rule across postgres restarts and crashes.

// buildAvailabilityStatus returns the current AvailabilityStatus for this node.
// Leaders that have resigned publish a LeadershipStatus. Every pooler publishes
// its cohort eligibility, so the result is non-nil.
func (pm *MultiPoolerManager) buildAvailabilityStatus() *clustermetadatapb.AvailabilityStatus {
	_ = "STUB: not implemented"
	return nil
}

// buildCohortEligibilityStatus returns the pooler's self-reported willingness
// to be a cohort member. Defaults to ELIGIBLE; downgraded to INELIGIBLE when
// the WAL receiver was manually stopped (StopReplication cleared
// primary_conninfo), so the coordinator does not try to re-include this node
// while the admin signal is in effect. setCohortEligibility (currently
// test-only) sets the base value the dynamic downgrade applies on top of.
func (pm *MultiPoolerManager) buildCohortEligibilityStatus() *clustermetadatapb.CohortEligibilityStatus {
	_ = "STUB: not implemented"
	return nil
}

// setCohortEligibility records this pooler's cohort eligibility. If the
// signal actually changed, an immediate health broadcast is pushed so the
// coordinator sees the new value without waiting for the next heartbeat —
// otherwise a transition INELIGIBLE → cohort removal could be delayed by up
// to a heartbeat interval. Currently test-only — there is no operator/admin
// RPC to flip it yet.
func (pm *MultiPoolerManager) setCohortEligibility(signal clustermetadatapb.CohortEligibilitySignal) {
	_ = "STUB: not implemented"
	return
}

// buildLeadershipStatus returns the LeadershipStatus for this node. Non-nil only
// when resignedLeaderAtTerm is set (i.e. after a BeginTerm REVOKE or graceful
// shutdown of a leader). Nil means this node has not recently held or resigned
// from primary leadership.
func (pm *MultiPoolerManager) buildLeadershipStatus() *clustermetadatapb.LeadershipStatus {
	_ = "STUB: not implemented"
	return nil
}

// setResignedLeaderAtTerm records that this node is requesting demotion as primary
// for the given term. The signal is included in subsequent StatusResponses so the
// coordinator can trigger an immediate election. Broadcasts to health stream
// subscribers when the value actually changes, so the coordinator sees the new
// signal without waiting for the next periodic snapshot.
// Requires the action lock (ctx must be an action-lock context).
func (pm *MultiPoolerManager) setResignedLeaderAtTerm(ctx context.Context, term int64) error {
	_ = "STUB: not implemented"
	return nil
}

// clearResignedLeaderAtTerm clears the leadership demotion request. Called by
// coordinator-driven promotion (Promote) when this node is explicitly
// re-appointed as primary at a new term.
// Requires the action lock (ctx must be an action-lock context).
func (pm *MultiPoolerManager) clearResignedLeaderAtTerm(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Recruit handles a coordinator's request to stop replication participation and
// record a TermRevocation, returning the node's stable position afterward.
//
// Order of operations:
//  1. Sanity-check the current rule position against the revocation term.
//  2. Stop replication participation (primary: full demote + restart as standby;
//     standby: clear primary_conninfo + drain replay).
//  3. Read the stable position and re-check against the revocation term to catch
//     the rare race where a WAL rule entry arrived after the sanity check.
//     On failure: primary re-promotes; standby restores primary_conninfo.
//  4. Persist the TermRevocation only if the position is consistent.
//  5. Return ConsensusStatus with the stable post-revoke position.
func (pm *MultiPoolerManager) Recruit(ctx context.Context, req *consensusdatapb.RecruitRequest) (*consensusdatapb.RecruitResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// State check — reject immediately if the node's committed WAL
// rule or stored revocation already conflicts with this request.
// Fails open on I/O error: a nil status passes ValidateRevocation safely.

// Refuse recruitment if a rewind is still pending from a prior emergency
// demotion. The node's WAL is in an indeterminate state until RewindToSource
// completes; allowing it to be recruited could elect a leader with divergent
// or missing WAL.

// Stop replication participation.
// non-empty if standby; used for recovery on race failure

// Save primary_conninfo so we can restore it if the position check fails.

/* wait */

// Re-check against the stable position and persist atomically.
// AcceptRevocation combines the observed WAL position with the locked stored
// revocation so ValidateRevocation sees authoritative state for both checks.

// Attempt to restore the node to its prior replication role.

// TODO: In theory it should be safe to re-promote the primary if this happens, but to keep things
// simpler for now we just keep publishing the signal that this pooler resigned from its term as
// leader to allow orch to do a failover.

// Step 5: Return ConsensusStatus with the stable post-revoke position.
// Uses the cached position warmed by the getConsensusStatus call in step 3.

// recruitDrainTimeout is the drain window when recruiting a primary.
const recruitDrainTimeout = 5 * time.Second

// setPrimaryConnInfoAndReload sets primary_conninfo and reloads postgres config so the
// WAL receiver reconnects. Used to restore a standby's replication after a recruit failure.
func (pm *MultiPoolerManager) setPrimaryConnInfoAndReload(ctx context.Context, connInfo string) error {
	_ = "STUB: not implemented"
	return nil
}

// Propose handles a coordinator's proposal for a new shard rule. The pooler
// either promotes its postgres to primary (if designated leader) or configures
// replication toward the new primary (if replica).
//
// Propose requires prior recruitment: the stored term_revocation must match the
// proposal's term exactly. There is no implicit recruitment on Propose.
//
// Order of operations:
//  1. Validate fields and check that the stored revocation matches the proposal term.
//  2. Determine role by comparing this pooler's ID to proposal_leader.id.
//     3a. Leader: promote postgres, write the rule to the rule store, enable query service.
//     3b. Replica: configure primary_conninfo toward the new leader's postgres.
//  4. Return ConsensusStatus with the post-propose position.
func (pm *MultiPoolerManager) Propose(ctx context.Context, req *consensusdatapb.ProposeRequest) (*consensusdatapb.ProposeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Identity and timing for the installed rule come from the proposed
// rule itself, not the revocation. The revocation's
// accepted_coordinator_id identifies who ran the recruit round; the
// rule's coordinator_id identifies the coordinator-of-record for this
// rule change. They are usually the same orch but the proposal is the
// authoritative source — falling back to time.Now() or the revocation
// would silently rewrite the caller's intent.

// Step 1: Validate the term revocation.
// ValidateRevocation ensures the WAL position is safe and the coordinator is consistent.
// Fails open on I/O error (nil status passes safely).

// Require an explicit Recruit() for this exact term before accepting a
// Propose. Implicit recruitment (accepting the term here without a prior
// Recruit call) could in principle be made safe, but it would need to
// reproduce everything Recruit does: pausing replication on replicas and
// restarting primaries in standby mode. We keep things simple for now by
// requiring the two-phase protocol. ValidateRevocation already ensures
// storedTerm <= revokedBelowTerm, so a mismatch here always means Recruit
// was never called for this term.

// Verify postgres is in the expected standby state: in recovery with no
// primary_conninfo set. Together these prove that Recruit ran (which clears
// primary_conninfo and goes into recovery mode) and that no prior Propose on
// this node succeeded.

// Propose is only valid for the designated leader. Non-leaders should
// receive the leader's identity via SetTermPrimary, which handles
// replication setup without requiring a prior Recruit.

// Leader path: promote postgres, write rule, enable query service.

// TODO: If this proposal already exists, we're being asked to propagate
// rather than make a new entry. We can make the rule store understand
// propagation for that case.

// IMPORTANT: updateTopologyAfterPromotion must only be called after updateRule
// succeeds. It advertises PRIMARY + SERVING to the gateway, opening write traffic.
// updateRule is the durability gate: it waits for sync-standby acknowledgment.

// Record the (rule, primary) — this pooler IS now the primary. Stamping
// the published ReplicationPrimary lets the health stream advertise the
// new leadership immediately.

// Step 4: Return ConsensusStatus. The cache was warmed by getConsensusStatus in step 1
// and updated by updateRule (leader path); the replica position is unchanged.

// SetTermPrimary updates this pooler's replication settings to point at the supplied
// primary, but only if the supplied position is strictly higher than the
// pooler's own current position. If the supplied position is equal or behind,
// SetTermPrimary is a successful no-op — this makes it safe under retries and under
// out-of-order delivery from stale recovery rounds.
//
// When the receiver is a standby, SetTermPrimary rewrites primary_conninfo (same effect
// as SetPrimaryConnInfo). When the receiver is currently acting as primary,
// the caller knows about a more recent rule with a different leader, so this
// node is a stale primary and gets demoted (same effect as DemoteStalePrimary).
//
// Unlike SetPrimaryConnInfo and DemoteStalePrimary, SetTermPrimary does not perform
// term validation — the rule comparison is the gate.
//
// TODO: when the rule comparison no-ops but WAL replay is paused
// (pg_is_wal_replay_paused), the caller's intent ("ensure this replica is
// pointed at the right primary") would be better served by also resuming
// replay. We don't do that today because StopReplication() is an explicit
// admin/test signal — auto-resuming would silently override it. Implement
// once StopReplication() can leave behind a "do not auto-resume" marker that
// SetTermPrimary can check.
func (pm *MultiPoolerManager) SetTermPrimary(ctx context.Context, req *consensusdatapb.SetTermPrimaryRequest) (*consensusdatapb.SetTermPrimaryResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The rule's leader_id is authoritative; the leader field carries contact
// info for that ID. A mismatch is a caller bug — we'd otherwise route
// replication at an identity that doesn't match the consensus-elected one.

// SetTermPrimary is the follower-side RPC. If the coordinator is telling this
// pooler that it is the new primary, that's a routing mistake — the leader
// path goes through Propose (which carries the full CoordinatorProposal and
// the Recruit-established term revocation needed to safely promote).

// Honor the revocation promise we made via Recruit/BeginTerm. If the
// incoming rule is revoked, ignore it: SetTermPrimary is a best-effort FYI and
// the cohort will reconverge as it makes progress. Returning the cached
// status keeps the response shape consistent with the "incoming rule
// not higher" no-op below.

// Record what we've been told, even if we don't end up applying the change.
// Two consumers:
//   - Health stream / multiorch: reads highest_known_rule to skip redundant
//     SetTermPrimary RPCs during the window after an apply but before streaming
//     replication has caught up enough for current_position to advance.
//   - Pooler-side reconciliation: reads last-known-primary to retry
//     ALTER SYSTEM SET primary_conninfo if this SetTermPrimary arrived while
//     postgres was unavailable.

// Observe the freshest view of our rule. SetTermPrimary is the staleness gate,
// so we want authoritative state — not the cached snapshot.

// Compare by RuleNumber only — LSN is intentionally not part of the gate.
// See SetTermPrimaryRequest's proto comment for the reasoning.

// Decide between "standby update" and "stale-primary demote" based on
// actual postgres recovery state rather than topology — a node mid-promote
// or mid-demote may have a topology label that lags reality.

// A standby with rewindPending=true was emergency-demoted earlier and still
// has divergent WAL relative to the new primary. Routing through
// demoteStalePrimaryLocked runs pg_rewind, which clears rewindPending and
// makes the node recruitable again. Without this, the lightweight standby
// branch sets primary_conninfo but leaves the WAL divergent, and the next
// Recruit refuses with "rewind pending after emergency demotion".

// Reported to the gateway as the new leader's term. Not term validation —
// the rule compare above is the gate. SetTermPrimary does not bump the local
// revocation: revocations are authored by coordinators via Recruit, and
// an SetTermPrimary is a notification, not a revoke.

/* stopReplicationBefore */ /* startReplicationAfter */

// Ensure topology reflects REPLICA. This matters when postgres has
// already been demoted (e.g. by BeginTerm REVOKE or an external
// pg_promote-then-restart) but the pooler's topology entry still
// reads PRIMARY. Without this, the stale PRIMARY label causes the
// stale-leader analyzer to keep firing forever. Propose has the same
// step on its replica branch for the same reason.

// Advertise the new leader to the health stream so the gateway can route
// reads/writes against it. The stale-primary branch gets this for free
// via demoteStalePrimaryLocked; the standby branch must do it explicitly.
// TODO: LeaderObservation is redundant with the (rule, primary) tuple
// already recorded in consensusState.replicationPrimary. Plan to make
// RecordTermPrimary (or its successor) drive the health-stream
// observation directly, so callers don't have to remember to do both.

// ConsensusStatus returns the current status of this node for consensus
func (pm *MultiPoolerManager) ConsensusStatus(ctx context.Context, req *consensusdatapb.StatusRequest) (*consensusdatapb.StatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
