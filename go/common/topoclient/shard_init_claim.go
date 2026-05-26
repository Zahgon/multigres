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

package topoclient

import (
	"context"

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
)

func shardInitClaimPath(shardKey *clustermetadatapb.ShardKey) string {
	_ = "STUB: not implemented"
	return ""
}

// ClaimShardInitialization atomically claims the right to initialize a shard
// using a create-if-not-exists write to the global topology.
//
// The claim persists both the claimer's identity and the proposed cohort so that
// a crash-retry reuses the same cohort for idempotency.
//
// Returns won=true and the committed cohort if this caller created the claim or
// is resuming its own prior claim. Returns won=false if a different coordinator
// already owns the initialization.
func (ts *store) ClaimShardInitialization(ctx context.Context, shardKey *clustermetadatapb.ShardKey, claimerID *clustermetadatapb.ID, proposedCohort []*clustermetadatapb.ID) (bool, []*clustermetadatapb.ID, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

// Claim already exists — read it back and check ownership.
