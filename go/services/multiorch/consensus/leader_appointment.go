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

package consensus

import (
	"context"

	"github.com/multigres/multigres/go/tools/pgutil"

	commonconsensus "github.com/multigres/multigres/go/common/consensus"

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	consensusdatapb "github.com/multigres/multigres/go/pb/consensusdata"
	multiorchdatapb "github.com/multigres/multigres/go/pb/multiorchdata"
)

// BeginTerm achieves Revocation, Candidacy, and Discovery by recruiting poolers
// under the proposed term:
//
//   - Revocation: Recruited poolers accept the new term, preventing any old leader
//     from completing requests under a previous term
//   - Discovery: Identifies the most progressed pooler based on WAL position to serve
//     as the candidate. From Raft: the log with highest term is most progressed;
//     for identical terms, highest LSN is most progressed
//   - Candidacy: Validates that recruited poolers satisfy the quorum rules, ensuring
//     the candidate has sufficient support to proceed
//
// The proposedTerm parameter is the term number to use (computed as maxTerm + 1).
//
// Returns the candidate pooler, standbys that accepted the term, the term, and any error.
func (c *Coordinator) BeginTerm(ctx context.Context, shardID string, cohort []*multiorchdatapb.PoolerHealthState, policy commonconsensus.DurabilityPolicy, proposedTerm int64) (*multiorchdatapb.PoolerHealthState, []*multiorchdatapb.PoolerHealthState, int64, error) {
	_ = "STUB: not implemented"
	return nil, nil, 0, nil
}

// Recruit Nodes - Send BeginTerm RPC to all poolers in parallel
// This is now FIRST to ensure we only select from nodes that accept the term

// Extract PoolerHealthState list for quorum validation

// Validate recruitment: proposal-agnostic invariants (revocation + majority)
// plus the candidacy check for this failover (does the recruited set satisfy
// the durability policy's quorum?). Candidacy is proposal-specific and
// lives at this layer rather than inside the policy method.

// Resigned poolers are excluded: they may still carry a stale primary rule
// that needs pg_rewind, which the stale-primary analyzer owns.
// TODO: once poolers self-rewind after emergency demotion, this exclusion
// becomes unnecessary — the resigned pooler can rejoin as a standby directly.

// discoverMaxTerm finds the maximum consensus term from cached health state.
// This uses the ConsensusTerm data already populated by health checks, avoiding extra RPCs.
func (c *Coordinator) discoverMaxTerm(cohort []*multiorchdatapb.PoolerHealthState) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Invariant: poolers in the cohort with successful health checks must have TermRevocation populated

// Invariant: at least one pooler in the cohort must have a term > 0

// walPositionLSN extracts and parses the most relevant LSN from a WALPosition.
// For a primary node CurrentLsn is used; for a standby, LastReceiveLsn.
// Returns (0, false) when pos is nil, both LSN fields are empty, or the string
// cannot be parsed.
func walPositionLSN(pos *consensusdatapb.WALPosition) (pgutil.LSN, bool) {
	_ = "STUB: not implemented"
	return *new(pgutil.LSN), false
}

// Get either the CurrentLsn (for primaries) or LastReceiveLsn (for
// standbys). Both field may be empty if the server is not a primary or
// standby, or if the pooler hasn't fully initialized and fetched WAL
// position yet.
// Use the most advanced LSN available: current (primary), receive (streaming standby),
// or replay (standby that has replayed from backup but not yet streaming).

// If all LSN fields are empty, we consider the WAL position invalid for
// selection purposes.

// selectCandidate chooses the best candidate from recruited poolers.
//
// WAL positions are captured after REVOKE (demote/pause), so they reflect the
// final state and won't advance further.
//
// Selection Strategy (per generalized consensus):
//
// Two-level ordering: leadership_term → LSN.
//
// 1. Highest leadership_term (primary criterion).
//
//	Each promotion and replication-config change writes a record to
//	multigres.rule_history with the current consensus term, using
//	RemoteOperationTimeout so synchronous standbys acknowledge the write
//	before the primary returns. The most recent coordinator_term in a standby's
//	local rule_history therefore reflects how far through the agreed
//	consensus history that standby has replicated. A standby with a higher
//	leadership_term has definitively applied more of the cluster's
//	committed WAL history than one with a lower term.
//
//	A stranded standby that diverged before a term-N promotion write was
//	replicated may accumulate a numerically larger LSN (e.g. a large
//	transaction written just before its primary crashed), but its
//	leadership_term will be lower, correctly excluding it.
//
//	Falls back to 0 when rule_history is empty (pre-bootstrap), in
//	which case the secondary criteria (LSN) determine the winner.
//
// 2. Highest LSN (secondary tiebreaker).
//
//	When terms are identical, LSN measures genuine progress
//	within the same history, so the node with the highest LSN is selected.
//
// Example (why LSN-only is wrong):
//
//	mp1: term=1, LSN=0/5000000  ← stranded, high LSN, old term
//	mp2: term=2, LSN=0/3000000  ← current term
//
// LSN-only would incorrectly elect mp1. Term-first correctly elects mp2.
func (c *Coordinator) selectCandidate(ctx context.Context, recruited []recruitmentResult) (*multiorchdatapb.PoolerHealthState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Skip nodes that have voluntarily requested demotion. The resignation
// signal is a deliberate request not to be re-elected; honoring it
// unconditionally avoids confusing re-elections of a node that just
// stepped down. If all candidates are resigned the election is deferred
// until a non-resigned candidate is available.

// Select by: highest leadership_term, then highest LSN.

// Log candidate selection details for observability. timeline_id is printed
// for debugging purposes but does not factor into selection criteria.

// selectCandidateFromRecruited chooses the best candidate from recruited nodes based on WAL position.
// Selection criteria: prefer the node with the highest LSN.
// recruitmentResult captures recruitment outcome and WAL position from BeginTerm response
type recruitmentResult struct {
	pooler      *multiorchdatapb.PoolerHealthState
	walPosition *consensusdatapb.WALPosition
}

// recruitNodes sends BeginTerm RPC to all poolers in parallel and returns those that accepted.
func (c *Coordinator) recruitNodes(ctx context.Context, cohort []*multiorchdatapb.PoolerHealthState, term int64, action consensusdatapb.BeginTermAction) ([]recruitmentResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Wait for all goroutines to complete

// Collect accepted poolers with WAL position data

// Log LSN from WAL position

// EstablishLeadership achieves the Propagation and Establishment goals from the consensus model:
// It makes the candidate's timeline durable under the new term and establishes leadership.
//
// In consensus terminology:
// - Propagation: Ensuring the recruited quorum has the candidate's complete timeline
// - Establishment: Delegating the term to the candidate so it can begin accepting requests
//
// This is accomplished by:
//  1. Configuring standbys to replicate from the candidate (before promotion)
//  2. Promoting the candidate to primary with synchronous replication configured.
//  3. Writing rule history under the new timeline. This write blocks until
//     acknowledged by the quorum, which proves:
//     a) The quorum has replicated the candidate's entire timeline (up to promotion point).
//     b) The quorum has replicated the rule history write itself.
//     c) The timeline is now durable under the new term.
//
// Once the rule history write succeeds, the new leader has successfully
// propagated its timeline and established leadership. The rule_history table serves
// as the canonical source of truth for when the new term began.
//
// Critical ordering: Standbys MUST be configured BEFORE promotion (step 1) to avoid
// deadlock. Promotion configures sync replication and writes rule history, which
// blocks waiting for acknowledgments. If standbys aren't replicating yet, the write
// blocks forever.
//
// If we fail to write rule history, leadership couldn't be established.
// A future coordinator will need to re-discover the most advanced timeline and re-propagate.
func (c *Coordinator) EstablishLeadership(
	ctx context.Context,
	candidate *multiorchdatapb.PoolerHealthState,
	standbys []*multiorchdatapb.PoolerHealthState,
	term int64,
	policy commonconsensus.DurabilityPolicy,
	reason string,
	cohort []*multiorchdatapb.PoolerHealthState,
	recruited []*multiorchdatapb.PoolerHealthState,
) error {
	_ = "STUB: not implemented"
	// Get current WAL position before promotion (for validation)
	return nil
}

// Wait for standby to replay all received WAL before promotion.
// This ensures validateExpectedLSN in Promote will pass.

// Build lists of cohort member IDs and accepted member IDs

// Configure standbys to replicate from the candidate BEFORE promoting.
// This ensures standbys are ready to connect when sync replication is configured.
// Without this, the Promote call can deadlock: it configures sync replication and
// tries to write rule history, but blocks waiting for standby acknowledgment.
// The standbys can't acknowledge because they haven't been told to replicate yet.

// Start streaming immediately so sync replication can proceed

// Check for errors - log but don't fail, standbys can be fixed later

// Build synchronous replication configuration based on quorum policy.
// Pass the full cohort; the policy excludes the candidate internally.

// Contract: the policy method must return a non-nil config on success so
// every promotion explicitly rewires sync replication. A nil config here
// is a bug, not a legitimate "no config needed" signal — refuse to proceed.

// preVote performs a pre-election check to decide whether an election is
// likely to succeed. It prevents disruptive elections that would fail due to:
//  1. Not enough currently reachable poolers to achieve a valid recruitment
//     (candidacy + revocation) under the durability policy.
//  2. Another coordinator recently started an election (within last 10 seconds).
//
// Returns (canProceed, reason) where canProceed indicates if election should proceed.
func (c *Coordinator) preVote(ctx context.Context, cohort []*multiorchdatapb.PoolerHealthState, policy commonconsensus.DurabilityPolicy, proposedTerm int64) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

// Filter cohort to poolers eligible to participate in recruitment right now.
// A pooler is eligible only if we can reach it, it's initialized with
// consensus-term data, and its postgres process is running. Without consensus-term
// info we can't reason about election safety; without postgres the pooler
// can't participate.
//
// We use postgres_running (process alive) rather than postgres_ready (pg_isready
// succeeds) so that standbys that are briefly unresponsive to pg_isready during
// WAL receiver reconnection after primary failure are still counted as eligible.
// A running multipooler process can accept BeginTerm RPCs regardless of whether
// pg_isready is momentarily failing.

// If we attempted recruitment right now with the eligible poolers, would
// the result be sufficient (candidacy + revocation) under the policy?
// If not, abort early rather than disrupt the cluster with a doomed election.

// Check 2: Has another coordinator recently started an election?
// If we detect a recent term acceptance (within the last 10 seconds), back off
// to give the other coordinator a chance to complete their election.

// Check if this pooler recently accepted a term from another coordinator

// If the acceptance was recent (within our window), back off

// poolerIDs extracts the clustermetadata IDs from a slice of PoolerHealthState.
// Used at the boundary where poolers cross into the durability-policy layer,
// which operates on bare *clustermetadatapb.ID values.
func poolerIDs(poolers []*multiorchdatapb.PoolerHealthState) []*clustermetadatapb.ID {
	_ = "STUB: not implemented"
	return nil
}
