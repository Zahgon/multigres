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

// CohortMismatchAnalyzer detects drift between the desired cohort and the
// recorded cohort on the leader.
//
// Two flavors of drift are reported:
//
//   - ProblemPoolerNotInCohort: a healthy, replicating, eligible pooler exists
//     in the shard but is absent from the leader's recorded cohort.
//   - ProblemCohortMemberIneligible: a current cohort member has self-reported
//     INELIGIBLE via AvailabilityStatus.CohortEligibilityStatus and should be
//     removed.
//
// Both surface a single ReconcileCohortAction; the action interprets the
// problem code and applies the appropriate ADD/REMOVE.
//
// TODO: this currently includes every healthy, eligible pooler in the cohort.
// In the future we'll likely want to constrain cohort size (e.g. cap at
// durability_required_count + 1) and choose the best-qualified members from
// the available poolers. That requires a fitness heuristic (LSN proximity,
// cell topology, last-known leadership history, etc.) that hasn't been
// designed yet.
type CohortMismatchAnalyzer struct {
	factory *RecoveryActionFactory
}

func (a *CohortMismatchAnalyzer) Name() types.CheckName {
	_ = "STUB: not implemented"
	return *new(types.CheckName)
}

func (a *CohortMismatchAnalyzer) ProblemCode() types.ProblemCode {
	_ = "STUB: not implemented"
	// This analyzer can produce two problem codes; ProblemCode() returns the
	// primary one. The recovery loop uses this for routing/logging — the
	// per-problem code on each emitted Problem is what matters at execution.
	return *new(types.ProblemCode)
}

func (a *CohortMismatchAnalyzer) RecoveryAction() types.RecoveryAction {
	_ = "STUB: not implemented"
	return *new(types.RecoveryAction)
}

func (a *CohortMismatchAnalyzer) Analyze(sa *ShardAnalysis) ([]types.Problem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We only act when the shard has a reachable, ready leader to receive the
// rule update. Bootstrap and failover paths set up the cohort separately.

// Build a set of current cohort member ID strings for O(1) lookup.

// Removal candidates: current cohort members signaling INELIGIBLE.

// Addition candidates: replicas not currently in the cohort that are
// healthy enough to be added.

// isAdditionCandidate reports whether a non-cohort pooler is healthy enough to
// be added: it must be a non-leader REPLICA, initialized, replicating, and not
// signaling INELIGIBLE.
//
// TODO: long-term, most of these checks should fold into the pooler's
// self-reported cohort eligibility signal. The pooler is in the best position
// to know whether it can durably serve as a cohort member — it knows whether
// it has a working backup, whether it's in a drained state, whether
// replication is healthy, and so on. The analyzer should largely just trust
// the CohortEligibilityStatus signal rather than reconstruct that judgment
// from individual health fields.
//
// The IsLeader gate is also conceptually unnecessary — there's no correctness
// problem an acting primary adding itself to the cohort. This may be useful
// in some propagation scenarios.
func (a *CohortMismatchAnalyzer) isAdditionCandidate(sa *ShardAnalysis, pa *PoolerAnalysis) bool {
	_ = "STUB: not implemented"
	return false
}

// Replication must be configured and not stopped — otherwise the standby
// can't acknowledge writes and adding it would degrade durability.
