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
	"log/slog"
	"time"

	"github.com/multigres/multigres/go/common/rpcclient"
	"github.com/multigres/multigres/go/common/topoclient"
	"github.com/multigres/multigres/go/services/multiorch/config"
	"github.com/multigres/multigres/go/services/multiorch/recovery/types"
	"github.com/multigres/multigres/go/services/multiorch/store"

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	multiorchdatapb "github.com/multigres/multigres/go/pb/multiorchdata"
)

// Compile-time assertion that DemoteStaleLeaderAction implements types.RecoveryAction.
var _ types.RecoveryAction = (*DemoteStaleLeaderAction)(nil)

// StaleLeaderDrainTimeout is a shorter drain timeout for stale leaders.
// Stale leaders that just came back online typically have no active connections,
// so we use a shorter timeout to speed up demotion.
const StaleLeaderDrainTimeout = 5 * time.Second

// DemoteStaleLeaderAction demotes a stale leader that was detected after failover.
// It uses the DemoteStalePrimary RPC with the correct leader's term to force the stale leader
// to accept the term and demote, preventing further writes.
type DemoteStaleLeaderAction struct {
	config      *config.Config
	rpcClient   rpcclient.MultiPoolerClient
	poolerStore *store.PoolerStore
	topoStore   topoclient.Store
	logger      *slog.Logger
}

// NewDemoteStaleLeaderAction creates a new action to demote a stale leader.
func NewDemoteStaleLeaderAction(
	cfg *config.Config,
	rpcClient rpcclient.MultiPoolerClient,
	poolerStore *store.PoolerStore,
	topoStore topoclient.Store,
	logger *slog.Logger,
) *DemoteStaleLeaderAction {
	_ = "STUB: not implemented"
	return nil
}

func (a *DemoteStaleLeaderAction) Metadata() types.RecoveryMetadata {
	_ = "STUB: not implemented"
	return *new(types.RecoveryMetadata)
}

func (a *DemoteStaleLeaderAction) Priority() types.Priority {
	_ = "STUB: not implemented"
	return *new(types.Priority)
}

func (a *DemoteStaleLeaderAction) RequiresHealthyLeader() bool {
	_ = "STUB: not implemented"
	// We're demoting a stale leader, so we can't require a healthy leader
	return false
}

func (a *DemoteStaleLeaderAction) GracePeriod() *types.GracePeriodConfig {
	_ = "STUB: not implemented"
	// Under the new consensus flow the demote goes through SetTermPrimary, which is
	// position-fenced: a leader that's only momentarily-stale would see its
	// own rule >= the incoming rule and no-op without touching postgres.
	// A spurious detection costs an RPC, not a wrongful demote, so the
	// grace period that guarded the old destructive DemoteStalePrimary RPC
	// is no longer needed. Skipping it lets recovery converge much faster
	// across sequential failovers.
	return nil
}

// Execute demotes the stale leader using the DemoteStalePrimary RPC with the correct leader's term.
// This is safer than BeginTerm because:
// 1. We use the correct leader's term (not a new term), avoiding term inconsistency
// 2. The stale leader accepts term >= its current term and demotes
// 3. Both leaders end up with the same term (no term inconsistency)
func (a *DemoteStaleLeaderAction) Execute(ctx context.Context, problem types.Problem) (retErr error) {
	_ = "STUB: not implemented"
	return nil
}

// Get the stale leader from the store

// Check if postgres is running on the stale leader before attempting demote.
// Demote requires postgres to be healthy. If postgres is not running yet,
// we should skip this attempt and let the next recovery cycle retry once
// postgres is ready. This avoids wasting time on RPCs that will fail.
// if !stalePrimary.IsPostgresReady {
// 	return mterrors.New(mtrpcpb.Code_UNAVAILABLE,
// 		fmt.Sprintf("postgres not running on stale leader %s, skipping demote attempt", poolerIDStr))
// }

// Find the correct leader to use as rewind source

// Demote the stale leader. Under the new consensus flow, route through
// SetTermPrimary .
//
// Both RPCs do the same work on the pooler side:
// 1. Stop postgres
// 2. Run pg_rewind to sync with the correct leader's postgres
// 3. Restart as standby
// 4. Clear sync replication config
// 5. Update topology to REPLICA

// findCorrectLeader finds the current leader in the shard and returns it along with its term.
// The correct leader is the one with the highest LeaderTerm.
func (a *DemoteStaleLeaderAction) findCorrectLeader(shardKey *clustermetadatapb.ShardKey, stalePrimaryIDStr string) (*multiorchdatapb.PoolerHealthState, int64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// Iterate through all poolers to find the current leader

// continue

// Only consider poolers in the same shard

// continue

// Skip the stale leader

// continue

// continue

// continue
