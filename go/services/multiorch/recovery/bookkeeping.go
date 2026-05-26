// Copyright 2025 Supabase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package recovery

import (
	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
)

// runBookkeeping performs periodic bookkeeping tasks.
func (re *Engine) runBookkeeping() { _ = "STUB: not implemented"; return }

// Reload configs first

// Forget instances that haven't been seen in a long time

// TODO: Add more bookkeeping tasks in future PRs
// - Expire old recovery history
// - Clean up stale data

// forgetLongUnseenInstances removes pooler instances that haven't been successfully
// health checked in over 4 hours. This handles three cases:
// 1. Broken entries (nil pointers - should never happen)
// 2. Instances discovered in topology but never successfully health checked
// 3. Instances that were previously healthy but haven't been seen in 4+ hours
func (re *Engine) forgetLongUnseenInstances() { _ = "STUB: not implemented"; return }

// Warn if store gets too large - operator should consider splitting watchers

// Collect entries to delete (can't delete while iterating due to lock)

// Iterate using Range() to hold lock during iteration

// Case 0: Broken entry (should never happen)

// broken entry — no valid ID to stop a stream with

// continue iteration

// Get timestamps as time.Time

// Case 1: Never successfully health checked (LastSeen is zero)

// Check how long since we first saw it (we don't have FirstDiscovered,
// so we use LastCheckAttempted as a proxy, or skip if both are zero)

// No attempts yet, skip for now
// continue iteration

// Case 2: Was previously healthy but not seen in 4+ hours

// continue iteration

// Now delete the entries (outside the iteration)

// audit logs an audit message with consistent formatting.
// This ensures important operations are logged in a structured way for compliance and debugging.
func (re *Engine) audit(auditType, poolerID string, shardKey *clustermetadatapb.ShardKey, message string) {
	_ = "STUB: not implemented"
	return
}
