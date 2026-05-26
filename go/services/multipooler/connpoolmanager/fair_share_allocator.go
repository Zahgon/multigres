// Copyright 2026 Supabase, Inc.
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

package connpoolmanager

// FairShareAllocator distributes connection capacity among users using max-min fairness.
// It is agnostic of resource type - create separate instances for regular and reserved pools.
//
// The algorithm ensures:
//   - Each user gets at least minPerUser connections (floor to handle burst demand)
//   - No user gets more than their demand (unless demand < minPerUser)
//   - Total allocation does not exceed capacity
//   - Remaining capacity is distributed fairly among unsatisfied users
type FairShareAllocator struct {
	capacity   int64
	minPerUser int64
}

// NewFairShareAllocator creates a new allocator with the given capacity budget.
// minPerUser sets the minimum allocation per user to handle burst demand that
// point-in-time sampling might miss. This prevents capacity from being reduced
// too aggressively for light users who occasionally need concurrent connections.
func NewFairShareAllocator(capacity int64, minPerUser int64) *FairShareAllocator {
	_ = "STUB: not implemented"
	return nil
}

// Capacity returns the total capacity this allocator manages.
func (a *FairShareAllocator) Capacity() int64 {
	_ = "STUB: not implemented"

	// Allocate distributes capacity among users based on their demands using max-min fairness.
	//
	// The algorithm (progressive filling):
	//  1. Start with all allocations at 0
	//  2. Calculate fair share = remaining_capacity / unsatisfied_users
	//  3. Give each unsatisfied user min(their_demand, fair_share)
	//  4. Users whose demand is met become "satisfied"
	//  5. Repeat with remaining capacity among unsatisfied users
	//  6. If capacity remains after all demands are met, split it evenly for burst headroom
	return 0
}

func (a *FairShareAllocator) Allocate(demands map[string]int64) map[string]int64 {
	_ = "STUB: not implemented"
	return nil
}

// Initialize allocations to 0

// Track which users are still unsatisfied (allocation < effective demand)

// Progressive filling: keep distributing until no capacity or all satisfied

// Calculate fair share of remaining capacity

// Not enough capacity for everyone - give 1 to as many as possible

// Track how much we actually allocate this round

// Effective demand is at least minPerUser to handle burst demand

// Already satisfied

// Give min(remaining_demand, fair_share, remaining_capacity)

// Check if now satisfied

// Remove satisfied users from unsatisfied set

// Safety: if we allocated nothing this round, break to avoid infinite loop

// If there's remaining capacity after all demands are met, distribute it evenly
// among all users for burst headroom. This ensures users can handle sudden
// traffic spikes without waiting for the next rebalance cycle.
