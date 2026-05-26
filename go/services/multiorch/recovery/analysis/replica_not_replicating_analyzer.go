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
	"github.com/multigres/multigres/go/services/multiorch/recovery/types"
)

// ReplicaNotReplicatingAnalyzer detects when a replica has no replication configured.
// This happens when primary_conninfo is not set or replication is stopped.
type ReplicaNotReplicatingAnalyzer struct {
	factory *RecoveryActionFactory
}

func (a *ReplicaNotReplicatingAnalyzer) Name() types.CheckName {
	_ = "STUB: not implemented"
	return *new(types.CheckName)
}

func (a *ReplicaNotReplicatingAnalyzer) ProblemCode() types.ProblemCode {
	_ = "STUB: not implemented"
	return *new(types.ProblemCode)
}

func (a *ReplicaNotReplicatingAnalyzer) RecoveryAction() types.RecoveryAction {
	_ = "STUB: not implemented"
	return *new(types.RecoveryAction)
}

func (a *ReplicaNotReplicatingAnalyzer) Analyze(sa *ShardAnalysis) ([]types.Problem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *ReplicaNotReplicatingAnalyzer) analyzePooler(sa *ShardAnalysis, poolerAnalysis *PoolerAnalysis) (*types.Problem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Only analyze replicas

// Skip if replica is not initialized (ShardNeedsInitialization handles that)

// Skip if there's no usable primary yet. HighestTermReachableLeader is
// non-nil only when the leader is reachable AND has published its rule
// (findHighestTermLeader filters out unobserved-rule leaders). Without a
// rule the recovery action can't populate SetTermPrimaryRequest.Rule — firing
// the problem now would produce a guaranteed-fail SetTermPrimary on the next
// cycle. PrimaryIsDead handles the unreachable case separately.

// Check if replication is not configured or stopped

// needsReplicationFix returns true if replication is not configured or stopped.
func (a *ReplicaNotReplicatingAnalyzer) needsReplicationFix(analysis *PoolerAnalysis) bool {
	_ = "STUB: not implemented"
	// No primary_conninfo configured
	return false
}

// Replication explicitly stopped
