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
	"log/slog"
	"sync"

	"github.com/multigres/multigres/go/common/rpcclient"
	"github.com/multigres/multigres/go/common/topoclient"
	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	multiorchdatapb "github.com/multigres/multigres/go/pb/multiorchdata"
)

// Coordinator orchestrates consensus-based leader election for shards.
//
// TODO: PoolerStore should be reorganized to be shard-centric rather than pooler-centric.
type Coordinator struct {
	coordinatorID *clustermetadatapb.ID
	topoStore     topoclient.Store
	rpcClient     rpcclient.MultiPoolerClient
	logger        *slog.Logger

	// useNewFlow enables the Recruit/Propose consensus path. Temporary; will be
	// removed once the old BeginTerm/EstablishLeadership path is deleted.
	useNewFlow bool

	// TODO: policyCache will go away when we start reading the policy from nodes instead of etcd.
	// This cache is a temporary way to avoid making failover depend on etcd.
	policyCache sync.Map // database name → *clustermetadatapb.DurabilityPolicy
}

// NewCoordinator creates a new coordinator instance. useNewFlow enables the
// Recruit/Propose consensus path; pass false to use the legacy
// BeginTerm/EstablishLeadership path.
func NewCoordinator(coordinatorID *clustermetadatapb.ID, topoStore topoclient.Store, rpcClient rpcclient.MultiPoolerClient, logger *slog.Logger, useNewFlow bool) *Coordinator {
	_ = "STUB: not implemented"
	return nil
}

// AppointLeader orchestrates the full consensus protocol to appoint a new leader
// for the given shard. It operates on a cohort of nodes (all nodes in the shard).
//
// The process achieves the following goals:
//
//  1. Obtaining a term number: Discover max term from cached health state and increment
//  2. Revocation, Candidacy, Discovery: BeginTerm recruits nodes under the new term,
//     achieving revocation (no old leader can complete requests), candidacy (recruited
//     nodes contain a suitable candidate), and discovery (identify most progressed node)
//  3. Propagation, Establishment: Make the timeline durable under the new term by configuring replication
//     and promoting the candidate. Timeline becomes durable when quorum rules are satisfied.
//
// Returns an error if any stage fails. The operation is idempotent and can be
// retried safely.
func (c *Coordinator) AppointLeader(ctx context.Context, shardID string, cohort []*multiorchdatapb.PoolerHealthState, database string, reason string) (retErr error) {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Apply the policy from the nodes themselves instead of assuming the bootstrap policy is
// still the durable shard policy. To do this, add durability_policy to PoolerHealthState so
// the health check loop populates it, then read from the cohort here for preVote and from the
// BeginTerm response after revocation. Note: GetDurabilityPolicy and CreateDurabilityPolicy
// RPCs in multipoolermanagerdata.proto are currently unimplemented stubs in the pooler.

// Goal 1: Obtaining a term number
// Discover max term from cached health state and increment to get proposed term

// runFailover wires the new-flow failover callbacks for a coordinatorLedRuleChange
// and runs it. Poolers that have signaled REQUESTING_DEMOTION are excluded from
// the cohort entirely: a writing leader's local LSN is always at least as
// advanced as any replica's (async replication), so including a resigned
// leader would let it dominate discovery and dead-end the proposal when health
// check rejects it. The full cohort identity is preserved via OutgoingRule's
// CohortMembers (replicas carry the same rule), so the consensus layer's
// outgoing-quorum check still runs against the original cohort size.
func (c *Coordinator) runFailover(ctx context.Context, cohort []*multiorchdatapb.PoolerHealthState, reason string) error {
	_ = "STUB: not implemented"
	return nil
}

// Failover constructs the revocation via NewTermRevocation: outgoing_rule
// is the highest RuleNumber discovered across cohort statuses.

// appointLeaderWithTerm is the shared core of AppointLeader and AppointInitialLeader.
// Given a resolved policy and proposed term, it runs preVote, BeginTerm, and
// EstablishLeadership.
func (c *Coordinator) appointLeaderWithTerm(ctx context.Context, shardID string, cohort []*multiorchdatapb.PoolerHealthState, policy *clustermetadatapb.DurabilityPolicy, proposedTerm int64, reason string) (retErr error) {
	_ = "STUB: not implemented"
	// Parse the proto policy once into the typed DurabilityPolicy interface so
	// preVote, BeginTerm, and EstablishLeadership can call its quorum,
	// recruitment, and leader-config methods directly.
	return nil
}

// Drop poolers that have self-revoked via REQUESTING_DEMOTION. Their
// BeginTerm RPC would block on the action lock held by their own
// graceful-shutdown sequence (e.g. while pgctld.Stop runs), and the
// legacy recruit fan-out waits for every goroutine — so a single
// resigning leader would stall failover by the full shutdown budget.
// Excluding them before preVote is important: preVote uses the cohort
// for its quorum check, so counting a pooler we're not going to recruit
// would cause preVote to pass against a quorum it can't actually achieve.
// selectCandidate refuses to elect a resigned pooler in any case, so the
// only thing we lose by skipping them is uncommitted WAL position,
// which sync replication makes safe to drop. (The new Recruit/Propose
// flow in rule_change.go does not have this bug — it commits as soon as
// quorum is recruited, so a slow node doesn't stall the path.)

// PreVote — validate that leadership change is likely to succeed.

// Goal 2: Revocation, Candidacy, Discovery
// BeginTerm recruits nodes under the new term, which achieves:
// - Revocation: recruited nodes accept new term, preventing old leader from completing requests
// - Discovery: identify the most progressed node based on WAL position
// - Candidacy: validate recruited nodes satisfy quorum rules for the candidate

// We know the candidate now — emit Started before establishing leadership.

// Reconstruct the recruited list (nodes that accepted the term).
// This is candidate + standbys.
//
// The recruited list may differ from the original cohort in these scenarios:
// - Some nodes in the cohort were unreachable during BeginTerm
// - Some nodes rejected the term (e.g., had a higher term already)
// - Some nodes failed validation (e.g., insufficient LSN)

// Propagation and Establishment

// AppointInitialLeader orchestrates consensus leader election for a freshly bootstrapped
// shard where all poolers start at term 0. It skips term discovery (which would
// return 0 for brand-new nodes) and calls BeginTerm directly with term=1.
//
// Uses GetBootstrapPolicy (not AppointLeader's LoadQuorumRule) because freshly restored
// standbys report UNKNOWN pooler type, which causes LoadQuorumRule to fall back
// to majority quorum instead of the configured durability policy.
func (c *Coordinator) AppointInitialLeader(ctx context.Context, shardID string, cohort []*multiorchdatapb.PoolerHealthState, database string) error {
	_ = "STUB: not implemented"
	return nil
}

// Bootstrap has no outgoing cohort to recruit consent from, so we use
// the externally-certified path. The "external" certification is the
// most-advanced timeline observed across the cohort's cached statuses:
// its rule number caps how far the outgoing cohort could have
// progressed, and its LSN is the frozen point any new leader must
// match. This handles partial bootstraps too — if any cohort member
// already carries a rule, we surface its rule number rather than
// falsely claiming term 0.

// This is the discovery phase of coordinator-led rule changes. For
// externally-certified rule changes, the agent (this method) is
// responsible for choosing the outgoing rule and authoring the
// revocation; common/consensus consumes the cert without re-deriving.

// Freshly bootstrapped shards start at term 0; skip discoverMaxTerm (which
// would return 0 for brand-new nodes) and use term 1 directly.
/* initialTerm */

// GetCoordinatorID returns the coordinator's ID.
func (c *Coordinator) GetCoordinatorID() *clustermetadatapb.ID {
	_ = "STUB: not implemented"
	return nil

	// GetShardNodes retrieves all multipooler nodes for a given shard from the topology.
}

func (c *Coordinator) GetShardNodes(ctx context.Context, cell string, database string, tablegroup string, shardID string) ([]*multiorchdatapb.PoolerHealthState, error) {
	_ = "STUB: not implemented"
	// Get all multipoolers in the cell for this specific shard
	return nil, nil
}

// Convert topology poolers to PoolerHealthState instances

// GetBootstrapPolicy returns the durability policy for the given database by reading
// bootstrap_durability_policy from the topology Database record. The result is cached
// in memory since the policy is assumed not to change for the lifetime of this process.
//
// TODO: Once pooler status updates carry policy information, this should be replaced
// with a live policy loaded from the shard's nodes rather than the bootstrap record.
func (c *Coordinator) GetBootstrapPolicy(ctx context.Context, database string) (*clustermetadatapb.DurabilityPolicy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
