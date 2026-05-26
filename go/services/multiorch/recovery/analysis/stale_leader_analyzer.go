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

// StaleLeaderAnalyzer detects stale leaders that came back online after failover.
// This happens when an old primary restarts without being properly demoted.
//
// The analyzer operates at the shard level: when multiple leaders are detected,
// it reports all of them except the highest-term leader as stale. Problems are
// sorted most-stale-first with descending priorities so the recovery system addresses
// the most out-of-date primary first.
//
// Note: This is NOT true split-brain. True split-brain means both primaries can accept
// writes. In this scenario, the new primary cannot accept writes because it cannot
// recruit standbys while the stale leader exists.
type StaleLeaderAnalyzer struct {
	factory *RecoveryActionFactory
}

func (a *StaleLeaderAnalyzer) Name() types.CheckName {
	_ = "STUB: not implemented"
	return *new(types.CheckName)
}

func (a *StaleLeaderAnalyzer) ProblemCode() types.ProblemCode {
	_ = "STUB: not implemented"
	return *new(types.ProblemCode)
}

func (a *StaleLeaderAnalyzer) RecoveryAction() types.RecoveryAction {
	_ = "STUB: not implemented"
	return *new(types.RecoveryAction)
}

func (a *StaleLeaderAnalyzer) Analyze(sa *ShardAnalysis) ([]types.Problem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Need multiple leaders to detect staleness.

// A tie in LeaderTerm indicates a consensus bug — skip automatic demotion.

// Collect stale leaders: every topology-PRIMARY pooler that is not the
// highest-term leader is stale. This includes poolers whose own rule has
// caught up (LeaderTerm == 0 because the rule now names a different
// leader) — exactly the post-emergency-demotion state we need to repair.

// Sort most stale first (lowest rule coordinator term first) so the
// recovery system processes the most out-of-date leader at highest
// priority.

// Assign descending priorities so the most stale leader (sorted first)
// gets PriorityEmergency, the next gets PriorityEmergency-1, etc.
