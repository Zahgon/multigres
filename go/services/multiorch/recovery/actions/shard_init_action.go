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

package actions

import (
	"context"
	"log/slog"

	"github.com/multigres/multigres/go/common/topoclient"
	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	multiorchdatapb "github.com/multigres/multigres/go/pb/multiorchdata"
	"github.com/multigres/multigres/go/services/multiorch/config"
	"github.com/multigres/multigres/go/services/multiorch/recovery/types"
	"github.com/multigres/multigres/go/services/multiorch/store"
)

// shardInitCoordinator is the subset of consensus.Coordinator used by ShardInitAction.
type shardInitCoordinator interface {
	GetBootstrapPolicy(ctx context.Context, database string) (*clustermetadatapb.DurabilityPolicy, error)
	AppointInitialLeader(ctx context.Context, shardID string, cohort []*multiorchdatapb.PoolerHealthState, database string) error
	GetCoordinatorID() *clustermetadatapb.ID
}

// Compile-time assertion that ShardInitAction implements types.RecoveryAction.
var _ types.RecoveryAction = (*ShardInitAction)(nil)

// ShardInitAction handles Phase 2 of shard bootstrap: the first backup has already
// been created and all poolers have restored from it. This action:
//  1. Reads the pooler store to collect initialized poolers and verify no cohort
//     is established yet
//  2. Ensures enough initialized poolers are available to satisfy the durability policy
//  3. Claims the exclusive right to initialize via topoStore.ClaimShardInitialization
//  4. Calls coordinator.AppointInitialLeader with the initialized poolers
type ShardInitAction struct {
	config      *config.Config
	coordinator shardInitCoordinator
	poolerStore *store.PoolerStore
	topoStore   topoclient.Store
	logger      *slog.Logger
}

// NewShardInitAction creates a new ShardInitAction.
func NewShardInitAction(
	cfg *config.Config,
	coordinator shardInitCoordinator,
	poolerStore *store.PoolerStore,
	topoStore topoclient.Store,
	logger *slog.Logger,
) *ShardInitAction {
	_ = "STUB: not implemented"
	return nil
}

// Execute performs the initial cohort establishment for a bootstrapped shard.
func (a *ShardInitAction) Execute(ctx context.Context, problem types.Problem) error {
	_ = "STUB: not implemented"
	return nil
}

// The recovery loop force-polls all poolers before calling Execute, so the pooler
// store holds fresh state. getInitializedPoolers reads that state: it returns nil
// if the cohort is already established, or the list of initialized poolers otherwise.

// Load the bootstrap durability policy from topology (not from nodes) because poolers
// are UNKNOWN type at this point — they have just restored from backup and are in
// hot-standby mode.

// Ensure the initialized poolers we see could ever satisfy the durability policy.

// Claim the exclusive right to initialize this shard and persist the cohort.
// On crash-retry the committed cohort is returned so we reuse the same members.

// Resolve committed IDs to full PoolerHealthState entries from the pooler store.

// getInitializedPoolers reads fresh pooler state from the store (already refreshed by the
// recovery loop before Execute is called). It returns the list of initialized poolers, plus
// a bool indicating whether the cohort is already established (any pooler has CohortMembers).
// If cohortEstablished is true the returned slice is nil and the caller should no-op.
func (a *ShardInitAction) getInitializedPoolers(shardKey *clustermetadatapb.ShardKey) (initialized []*multiorchdatapb.PoolerHealthState, cohortEstablished bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// stop iteration

// buildCohortFromIDs resolves committed cohort IDs to PoolerHealthState entries
// from the local pooler store. Returns only the poolers we can find locally.
func (a *ShardInitAction) buildCohortFromIDs(poolers []*multiorchdatapb.PoolerHealthState, committedIDs []*clustermetadatapb.ID) []*multiorchdatapb.PoolerHealthState {
	_ = "STUB: not implemented"
	return nil
}

// RecoveryAction interface implementation

func (a *ShardInitAction) RequiresHealthyLeader() bool { _ = "STUB: not implemented"; return false }

func (a *ShardInitAction) Metadata() types.RecoveryMetadata {
	_ = "STUB: not implemented"
	return *new(types.RecoveryMetadata)
}

func (a *ShardInitAction) Priority() types.Priority {
	_ = "STUB: not implemented"
	return *new(types.Priority)
}

func (a *ShardInitAction) GracePeriod() *types.GracePeriodConfig {
	_ = "STUB: not implemented"
	return nil
}
