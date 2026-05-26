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

package analysis

import (
	"log/slog"

	"github.com/multigres/multigres/go/common/rpcclient"
	"github.com/multigres/multigres/go/common/topoclient"
	"github.com/multigres/multigres/go/services/multiorch/config"
	"github.com/multigres/multigres/go/services/multiorch/consensus"
	"github.com/multigres/multigres/go/services/multiorch/recovery/types"
	"github.com/multigres/multigres/go/services/multiorch/store"
)

// RecoveryActionFactory creates recovery actions with all necessary dependencies.
type RecoveryActionFactory struct {
	config      *config.Config
	poolerStore *store.PoolerStore
	rpcClient   rpcclient.MultiPoolerClient
	topoStore   topoclient.Store
	coordinator *consensus.Coordinator
	logger      *slog.Logger
}

// NewRecoveryActionFactory creates a factory for recovery actions.
func NewRecoveryActionFactory(
	cfg *config.Config,
	poolerStore *store.PoolerStore,
	rpcClient rpcclient.MultiPoolerClient,
	topoStore topoclient.Store,
	coordinator *consensus.Coordinator,
	logger *slog.Logger,
) *RecoveryActionFactory {
	_ = "STUB: not implemented"
	return nil
}

// NewShardInitAction creates a shard initialization action.
func (f *RecoveryActionFactory) NewShardInitAction() types.RecoveryAction {
	_ = "STUB: not implemented"
	return *new(types.RecoveryAction)
}

// NewAppointLeaderAction creates an appoint leader action.
func (f *RecoveryActionFactory) NewAppointLeaderAction() types.RecoveryAction {
	_ = "STUB: not implemented"
	return *new(types.RecoveryAction)
}

// NewFixReplicationAction creates a fix replication action.
func (f *RecoveryActionFactory) NewFixReplicationAction() types.RecoveryAction {
	_ = "STUB: not implemented"
	return *new(types.RecoveryAction)
}

// NewDemoteStaleLeaderAction creates an action to demote a stale primary.
func (f *RecoveryActionFactory) NewDemoteStaleLeaderAction() types.RecoveryAction {
	_ = "STUB: not implemented"
	return *new(types.RecoveryAction)
}

// NewReconcileCohortAction creates an action to add or remove a cohort member.
func (f *RecoveryActionFactory) NewReconcileCohortAction() types.RecoveryAction {
	_ = "STUB: not implemented"
	return *new(types.RecoveryAction)
}

// Logger returns the factory's logger for use by analyzers.
func (f *RecoveryActionFactory) Logger() *slog.Logger {
	_ = "STUB: not implemented"

	// Config returns the factory's config for use by analyzers.
	return nil
}

func (f *RecoveryActionFactory) Config() *config.Config { _ = "STUB: not implemented"; return nil }
