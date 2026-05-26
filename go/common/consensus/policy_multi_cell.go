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

package consensus

import (
	"log/slog"

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
)

// MultiCellPolicy requires acknowledgement from poolers spanning at least N
// distinct cells.
type MultiCellPolicy struct {
	N int
}

// CheckAchievable returns nil iff the proposed cohort spans at least N
// distinct cells.
func (p MultiCellPolicy) CheckAchievable(proposedCohort []*clustermetadatapb.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// CheckSufficientRecruitment enforces two proposal-agnostic invariants:
//   - Revocation: the un-recruited cohort poolers span fewer than N distinct
//     cells, so they cannot themselves form a commit quorum satisfying the
//     policy. Cell coverage by recruited is not enough when a cell holds
//     multiple poolers: a recruited pooler in a cell does not block an
//     un-recruited pooler in the same cell from participating in a separate
//     quorum elsewhere.
//   - Majority: recruited is a pooler-majority of cohort, so any two
//     concurrent recruitments must share at least one pooler. Cell-level
//     intersection is not sufficient here because two pooler-disjoint
//     recruitments can share cells when a cell has multiple poolers.
//
// Candidacy (whether recruited spans enough cells for the *proposed*
// leadership change) is not checked here — that is a proposal-specific
// concern handled by the leader-appointment layer via CheckAchievable.
func (p MultiCellPolicy) CheckSufficientRecruitment(cohort, recruited []*clustermetadatapb.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// Revocation: if the un-recruited cohort poolers themselves span N or more cells, they could
// form a commit quorum on their own under this policy. We can't allow that.
// Example: MULTI_CELL_AT_LEAST_2 and a cohort of 6 poolers (2 per cell across 3 cells).
// Recruiting one pooler from each cell covers every cohort cell, but the 3 un-recruited
// poolers still span 3 cells — enough to form a separate 2-cell quorum on their own.

// BuildSyncReplicationConfig returns the Postgres-level config the primary
// must apply to satisfy MULTI_CELL_AT_LEAST_N. Standbys in the primary's own
// cell are excluded so synchronous acknowledgement always crosses a cell boundary.
//
// Errors when no eligible different-cell standbys exist or when the eligible
// set is too small to satisfy num_sync.
func (p MultiCellPolicy) BuildSyncReplicationConfig(
	logger *slog.Logger,
	cohort []*clustermetadatapb.ID,
	primary *clustermetadatapb.ID,
) (*SyncReplicationConfig, error) {
	_ = "STUB: not implemented"
	// N==1 means the primary alone satisfies durability — return an explicit
	// "no sync standbys" config so the new primary clears any stale
	// synchronous_standby_names instead of silently inheriting them.
	return nil, nil
}

// Drop cohort members in the primary's own cell so synchronous
// acknowledgement always crosses a cell boundary. The primary itself is
// naturally excluded (it's in its own cell).

// num_sync = required_count - 1: primary itself counts as 1 ack.

// Description returns a human-readable summary of the policy.
func (p MultiCellPolicy) Description() string { _ = "STUB: not implemented"; return "" }

// cellsOf returns the set of distinct cells covered by poolers.
func cellsOf(poolers []*clustermetadatapb.ID) map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}
