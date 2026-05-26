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

	"github.com/multigres/multigres/go/common/topoclient"
	"github.com/multigres/multigres/go/services/multiorch/config"
	"github.com/multigres/multigres/go/services/multiorch/consensus"
	"github.com/multigres/multigres/go/services/multiorch/recovery/types"
	"github.com/multigres/multigres/go/services/multiorch/store"

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	multiorchdatapb "github.com/multigres/multigres/go/pb/multiorchdata"
)

// Compile-time assertion that AppointLeaderAction implements types.RecoveryAction.
var _ types.RecoveryAction = (*AppointLeaderAction)(nil)

// AppointLeaderAction handles leader appointment using the coordinator's consensus protocol.
// This action is used for both repair (mixed initialized/empty nodes) and reelect
// (all nodes initialized) scenarios. The consensus.AppointLeader method handles
// both cases by selecting the most advanced node based on WAL position and running
// the full consensus protocol to establish a new primary.
type AppointLeaderAction struct {
	config      *config.Config
	consensus   *consensus.Coordinator
	poolerStore *store.PoolerStore
	topoStore   topoclient.Store
	logger      *slog.Logger
}

// NewAppointLeaderAction creates a new leader appointment action
func NewAppointLeaderAction(
	cfg *config.Config,
	consensus *consensus.Coordinator,
	poolerStore *store.PoolerStore,
	topoStore topoclient.Store,
	logger *slog.Logger,
) *AppointLeaderAction {
	_ = "STUB: not implemented"
	return nil
}

// Execute performs leader appointment by running the coordinator's consensus protocol
func (a *AppointLeaderAction) Execute(ctx context.Context, problem types.Problem) error {
	_ = "STUB: not implemented"
	return nil
}

// Fetch cohort and recheck the problem

// Check if a primary already exists and is healthy (problem resolved).
// We must verify both that the pooler is reachable (IsLastCheckValid) AND that
// PostgreSQL is ready (IsPostgresReady). If the pooler is up but Postgres
// is not ready, or if the primary has signalled it needs replacement, we still
// need to trigger failover.
//
// Note: this relies on the resign flow maintaining PoolerType_PRIMARY until
// DemoteStalePrimary completes. If a node somehow becomes the consensus leader
// while reporting PoolerType_REPLICA (e.g. a crash-restart as standby without
// going through the normal resign → appoint → demote flow), this check would
// miss it and proceed with an appointment unnecessarily.

// Use the coordinator's AppointLeader to handle the election
// It will select the most advanced node based on WAL position
// and run the full consensus protocol (term discovery, candidate selection,
// node recruitment, quorum validation, promotion, and replication setup)
//
// Use the problem code as the reason for the election

// getCohort fetches all poolers in the shard from the pooler store.
func (a *AppointLeaderAction) getCohort(shardKey *clustermetadatapb.ShardKey) []*multiorchdatapb.PoolerHealthState {
	_ = "STUB: not implemented"
	return nil
}

// continue

// continue

// RecoveryAction interface implementation

func (a *AppointLeaderAction) RequiresHealthyLeader() bool {
	_ = "STUB: not implemented"
	// leader appointment doesn't need existing primary
	return false
}

func (a *AppointLeaderAction) Metadata() types.RecoveryMetadata {
	_ = "STUB: not implemented"
	return *new(types.RecoveryMetadata)
}

// can retry if it fails

func (a *AppointLeaderAction) Priority() types.Priority {
	_ = "STUB: not implemented"
	return *new(types.Priority)
}

func (a *AppointLeaderAction) GracePeriod() *types.GracePeriodConfig {
	_ = "STUB: not implemented"
	return nil
}
