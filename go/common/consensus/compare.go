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

// Package consensus provides utilities for working with consensus types.
package consensus

import (
	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
)

// CompareRuleNumbers compares two RuleNumbers lexicographically.
// Returns -1 if a < b, 0 if a == b, 1 if a > b.
// A nil RuleNumber is treated as zero (the smallest possible value).
func CompareRuleNumbers(a, b *clustermetadatapb.RuleNumber) int {
	_ = "STUB: not implemented"
	return 0
}

// MostAdvancedPosition returns the highest-ranked PoolerPosition among the
// given statuses. Rule number takes precedence; LSN breaks ties within the
// same rule. Returns nil if no status has a parseable LSN.
//
// This is the cached-snapshot analogue of discoverMostAdvancedTimeline, which
// runs over recruited statuses and returns the eligible-leader set. Callers
// that need to derive an ExternallyCertifiedRevocation from cached cohort
// state — e.g. the bootstrap path before recruitment — use this to obtain the
// outgoing rule number and frozen LSN.
func MostAdvancedPosition(statuses []*clustermetadatapb.ConsensusStatus) *clustermetadatapb.PoolerPosition {
	_ = "STUB: not implemented"
	return nil
}

// ReplicationPrimaryMatches reports whether a pooler's published
// ReplicationPrimary already names target as its primary at a rule no older
// than targetRule. Coordinators use this to skip SetTermPrimary RPCs that
// wouldn't change anything on the pooler.
//
// Returns false when:
//   - rp is nil
//   - the published rule is strictly older than targetRule
//   - the published primary is missing
//   - the published primary's (id, host, postgres port) differs from target's
//
// target and targetRule are required; passing nil for either returns false.
func ReplicationPrimaryMatches(rp *clustermetadatapb.ReplicationPrimary, target *clustermetadatapb.PoolerAddress, targetRule *clustermetadatapb.ShardRule) bool {
	_ = "STUB: not implemented"
	return false
}

func idsEqual(a, b *clustermetadatapb.ID) bool { _ = "STUB: not implemented"; return false }

// ComparePosition returns negative, zero, or positive based on whether a is
// behind, equal to, or ahead of b. Rule number takes precedence; LSN breaks
// ties within the same rule. A missing or unparsable LSN is treated as less
// than any valid LSN.
func ComparePosition(a, b *clustermetadatapb.PoolerPosition) int {
	_ = "STUB: not implemented"
	return 0
}
