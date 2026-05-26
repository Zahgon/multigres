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

package analysis

import (
	"github.com/multigres/multigres/go/services/multiorch/recovery/types"
)

// LeaderResignedAnalyzer detects when the topology leader has explicitly
// requested its own replacement via LEADERSHIP_SIGNAL_REQUESTING_DEMOTION
// (emitted by EmergencyDemote and graceful shutdown of a leader). Distinct
// from LeaderIsDeadAnalyzer, which infers leader loss from reachability and
// health. Resignation is an unambiguous, intentional signal so failover can
// fire immediately with no follower-connection grace period.
type LeaderResignedAnalyzer struct {
	factory *RecoveryActionFactory
}

func (a *LeaderResignedAnalyzer) Name() types.CheckName {
	_ = "STUB: not implemented"
	return *new(types.CheckName)
}

func (a *LeaderResignedAnalyzer) ProblemCode() types.ProblemCode {
	_ = "STUB: not implemented"
	return *new(types.ProblemCode)
}

func (a *LeaderResignedAnalyzer) RecoveryAction() types.RecoveryAction {
	_ = "STUB: not implemented"
	return *new(types.RecoveryAction)
}

func (a *LeaderResignedAnalyzer) Analyze(sa *ShardAnalysis) ([]types.Problem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
