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

package consensus

import (
	"context"
	"log/slog"

	commonconsensus "github.com/multigres/multigres/go/common/consensus"

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	consensusdatapb "github.com/multigres/multigres/go/pb/consensusdata"
	multiorchdatapb "github.com/multigres/multigres/go/pb/multiorchdata"
)

// coordinatorLedRuleChange orchestrates the recruit → check → propose workflow
// for a coordinator-initiated rule change. It is parameterized by action-specific
// callbacks so the same workflow serves both normal failover and bootstrap.
type coordinatorLedRuleChange struct {
	coordinator           *Coordinator
	reason                string
	tryBuildProposal      func(*clustermetadatapb.TermRevocation, []*clustermetadatapb.ConsensusStatus) (*consensusdatapb.CoordinatorProposal, error)
	checkProposalPossible func(*clustermetadatapb.TermRevocation, []*clustermetadatapb.ConsensusStatus) error
}

func (c *Coordinator) newRuleChange(
	reason string,
	tryBuildProposal func(*clustermetadatapb.TermRevocation, []*clustermetadatapb.ConsensusStatus) (*consensusdatapb.CoordinatorProposal, error),
	checkProposalPossible func(*clustermetadatapb.TermRevocation, []*clustermetadatapb.ConsensusStatus) error,
) *coordinatorLedRuleChange {
	_ = "STUB: not implemented"
	return nil
}

// Run executes the rule change: pre-validate, recruit all nodes concurrently,
// and propose as soon as a viable proposal can be assembled. Each node
// receives its Propose immediately once it has been recruited and the
// proposal is ready — there is no unnecessary waiting between the two.
//
// revocation is the authoritative term revocation the coordinator is
// proposing. Callers own its construction:
//   - For safe coordinator-led transitions (failover), use
//     commonconsensus.NewTermRevocation to derive it from cohort statuses.
//   - For externally-certified transitions (bootstrap, operator override),
//     construct the cert and pass cert.GetTermRevocation() — the agent
//     defines revoked_below_term and outgoing_rule, not local discovery.
func (r *coordinatorLedRuleChange) Run(
	ctx context.Context,
	cohort []*multiorchdatapb.PoolerHealthState,
	revocation *clustermetadatapb.TermRevocation,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Extract cached consensus statuses for the pre-vote feasibility check.

// Back off if any node recently accepted a revocation — another coordinator
// may be running an election.

// Pre-validate that a proposal would be feasible with current statuses before
// committing to a recruitment round.

// Recruit all nodes concurrently.

// nil if recruit failed

// dispatchProposal sends Propose to the designated leader and SetTermPrimary
// to every other cohort member. r.propose handles the dispatch internally.

// Phase 1: collect recruits until tryBuildProposal succeeds (or we run out).
//
// We commit to the leader at the FIRST successful tryBuildProposal and stream
// Proposes from there — recruits arriving after that point don't change
// the choice. This is consensus-correct: only committed writes are
// guaranteed durable across the quorum, so the chosen leader carries
// every committed write by definition. The tradeoff is that *uncommitted*
// WAL on a late-arriving more-advanced node is discarded when the node
// joins as a follower (pg_rewind on rejoin) instead of being preserved
// by electing that node leader.
//
// TODO: add a configurable grace period after quorum is reached but
// before the leader is locked in, to let slow-but-not-dead nodes
// participate in leader selection. With grace=0 we keep the current
// streaming behavior; with grace>0 we wait up to that long for further
// recruits and re-run tryBuildProposal on the augmented set. The right default
// should come from observed failover latency in production.

// Phase 2: drain remaining recruits, proposing to each as it arrives.

// Wait for every Propose to return, not just the leader's. The rule
// change is committed when the leader's Propose succeeds — which can
// only happen because enough followers endorsed the proposal first to
// satisfy the cohort's durability quorum. Returning early here would
// let Run's context cancel the in-flight non-leader Proposes. The shard
// is already safe at that point, but any pooler whose Propose was
// cancelled wouldn't learn the new primary on this round and would
// need to be re-wired before it could route correctly. Waiting avoids
// that follow-up.
//
// Exception: if the leader's Propose fails, no new primary exists for
// non-leaders to learn about, so there is nothing to gain by waiting
// — return as soon as we see the leader's failure.

// recruit issues a Recruit RPC to a single pooler and returns the resulting
// ConsensusStatus, or nil if the call failed or returned an empty status.
func (r *coordinatorLedRuleChange) recruit(
	ctx context.Context,
	p *multiorchdatapb.PoolerHealthState,
	revocation *clustermetadatapb.TermRevocation,
) *clustermetadatapb.ConsensusStatus {
	_ = "STUB: not implemented"
	return nil
}

// propose dispatches the rule change to a single pooler. The designated leader
// receives a Propose RPC (it promotes postgres and writes the new rule); every
// other cohort member receives SetTermPrimary (it points replication at the
// new leader). Both RPCs read directly from the CoordinatorProposal —
// proposal_leader is a PoolerAddress, exactly what SetTermPrimary's leader
// field wants.
func (r *coordinatorLedRuleChange) propose(
	ctx context.Context,
	p *multiorchdatapb.PoolerHealthState,
	req *consensusdatapb.ProposeRequest,
	isLeader bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

// buildFailoverProposal constructs a CoordinatorProposal for normal failover.
// It picks the first eligible leader from result.EligibleLeaders and derives
// the cohort and durability policy from result.OutgoingRule. Resigned poolers
// are expected to have been filtered out upstream (see Coordinator.runFailover);
// any pooler reaching this point is treated as a valid leader candidate.
func buildFailoverProposal(
	result commonconsensus.RecruitmentResult,
	addressByID map[string]*clustermetadatapb.PoolerAddress,
) (*consensusdatapb.CoordinatorProposal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The coordinator that ran the recruit round (carried in the
// revocation's accepted_coordinator_id) is also the
// coordinator-of-record for the rule it produces.

// buildBootstrapProposal constructs a CoordinatorProposal for initial leader
// appointment on a fresh shard. Unlike buildFailoverProposal, the cohort and
// durability policy come from the caller — there is no recorded rule to
// derive them from — and any eligible leader works since none have prior
// commitments.
func buildBootstrapProposal(
	result commonconsensus.RecruitmentResult,
	cohortIDs []*clustermetadatapb.ID,
	policy *clustermetadatapb.DurabilityPolicy,
	addressByID map[string]*clustermetadatapb.PoolerAddress,
) (*consensusdatapb.CoordinatorProposal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// checkRecentAcceptance returns an error if any node in the cohort recently
// accepted a term revocation, which may indicate another coordinator is making
// a rule change.
func checkRecentAcceptance(ctx context.Context, logger *slog.Logger, cohort []*multiorchdatapb.PoolerHealthState) error {
	_ = "STUB: not implemented"
	return nil
}

// buildCohortMaps returns address-by-ID and health-by-ID lookup maps built
// from a cohort slice. Keys are ClusterIDString values. Consensus paths only
// need the leader's contact info to build proposals, so the cohort map is
// flattened to PoolerAddress here.
func buildCohortMaps(cohort []*multiorchdatapb.PoolerHealthState) (map[string]*clustermetadatapb.PoolerAddress, map[string]*multiorchdatapb.PoolerHealthState) {
	_ = "STUB: not implemented"
	return nil, nil
}
