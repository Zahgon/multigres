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

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	multiorchdatapb "github.com/multigres/multigres/go/pb/multiorchdata"
)

// ApplyCertifiedRuleChange installs a new shard rule using a fully-populated
// externally certified revocation. Used for both initial leader appointment
// (term 0 → term 1) and stuck-quorum recovery (term N → term N+1).
//
// The coordinator is a pure executor here: every identity and timing field
// (proposed_rule.coordinator_id, .creation_time, .rule_number;
// cert.term_revocation.*) must be populated by the caller. The coordinator
// validates the inputs, then runs the standard Recruit + Propose fan-out
// against the proposed cohort.
func (c *Coordinator) ApplyCertifiedRuleChange(
	ctx context.Context,
	shardKey *clustermetadatapb.ShardKey,
	proposedRule *clustermetadatapb.ShardRule,
	cert *clustermetadatapb.ExternallyCertifiedRevocation,
	reason string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Refresh ConsensusStatus from every pooler registered for the shard so
// the pre-vote check in coordinatorLedRuleChange.Run sees real positions;
// without primed statuses the externally-certified pre-vote sees an empty
// set and fails before recruitment can prove the cert is acceptable.

// Defense in depth: validateCertifiedRuleChange already verified the
// leader appears in cohort_members by ID, and resolveCohort fetched a
// MultiPooler for every cohort member, so this lookup cannot fail in
// practice. The explicit check guards against future refactors.

// The cert's term_revocation is the revocation each pooler will record
// via Recruit. Pre-build the proposal struct that BuildExternallyCertifiedProposal
// will return — it does not change as recruits arrive, since the caller has
// already committed to a specific leader/cohort/durability.

// Externally-certified proposals bypass the outgoing-cohort quorum check:
// the cert attests that the outgoing rule's quorum cannot commit further
// writes, so the receiving pooler should apply the incoming cohort GUC
// directly rather than computing a (incoming ∩ outgoing) transition that
// collapses to empty for bootstrap (where the outgoing cohort is empty)
// or for recovery (where the outgoing quorum is unrecoverable).

// refreshShardConsensusStatuses calls ConsensusStatus on every pooler
// registered for the shard (not just the proposed cohort) in parallel and
// returns the responses keyed by ClusterIDString.
//
// Unreachable poolers are absent from the map: by submitting the cert, the
// operator has already attested that absent nodes will not commit further
// writes, and we can only verify what we can observe.
//
// Topology enumeration errors are fatal — silently skipping a cell we could
// not list might miss a pooler that is ahead of the cert.
func (c *Coordinator) refreshShardConsensusStatuses(
	ctx context.Context,
	shardKey *clustermetadatapb.ShardKey,
) (map[string]*clustermetadatapb.ConsensusStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// resolveCohort looks up each cohort member ID in topology and returns an
// address-by-ID lookup map plus a PoolerHealthState slice for the
// coordinatorLedRuleChange runner. PoolerHealthState entries carry only
// MultiPooler — consensus statuses are gathered fresh during recruit.
func (c *Coordinator) resolveCohort(
	ctx context.Context,
	cohortMembers []*clustermetadatapb.ID,
) (map[string]*clustermetadatapb.PoolerAddress, []*multiorchdatapb.PoolerHealthState, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// validateCertifiedRuleChange enforces the shape contract documented on the
// proto request: every identity and timing field must be set, the proposed
// rule_number must agree with the cert's revoked_below_term, and the leader
// must be in the proposed cohort.
func validateCertifiedRuleChange(
	shardKey *clustermetadatapb.ShardKey,
	proposedRule *clustermetadatapb.ShardRule,
	cert *clustermetadatapb.ExternallyCertifiedRevocation,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Cross-field consistency: the proposed rule's coordinator_term must match
// the cert's revoked_below_term (enforced by validateProposal in
// proposals.go, but we check eagerly here for a clearer error).

// Leader must be in the cohort.
