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

// ShardNeedsInitializationAnalyzer detects Phase 2 of shard bootstrap: the shard has
// been initialized (first backup created and a 0-member cohort record written) but
// its initial cohort has not yet been established by multiorch.
//
// Fires when the shard has at least one reachable, initialized pooler, no primary,
// and no pooler in the shard has cohort members. This is a shard-level problem:
// ShardInitAction needs to see all initialized poolers to form a quorum.
type ShardNeedsInitializationAnalyzer struct {
	factory *RecoveryActionFactory
}

func (a *ShardNeedsInitializationAnalyzer) Name() types.CheckName {
	_ = "STUB: not implemented"
	return *new(types.CheckName)
}

func (a *ShardNeedsInitializationAnalyzer) ProblemCode() types.ProblemCode {
	_ = "STUB: not implemented"
	return *new(types.ProblemCode)
}

func (a *ShardNeedsInitializationAnalyzer) RecoveryAction() types.RecoveryAction {
	_ = "STUB: not implemented"
	return *new(types.RecoveryAction)
}

func (a *ShardNeedsInitializationAnalyzer) Analyze(sa *ShardAnalysis) ([]types.Problem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Without a policy we can't determine the quorum requirement — skip.

// If any pooler has cohort members, the cohort is already established.

// Not enough initialized poolers to satisfy the durability policy yet.

// Not achievable yet — silently skip; the analyzer re-evaluates as poolers come online.
//nolint:nilerr
