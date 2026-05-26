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

import (
	"context"
)

// startRebalancer starts the background rebalancer goroutine.
// The rebalancer periodically:
//  1. Collects demand from DemandTrackers
//  2. Computes fair allocations using FairShareAllocator
//  3. Applies new capacities via UserPool.SetCapacity()
//  4. Garbage collects inactive user pools
func (m *Manager) startRebalancer() { _ = "STUB: not implemented"; return }

// rebalanceLoop is the main loop for the rebalancer goroutine.
func (m *Manager) rebalanceLoop() { _ = "STUB: not implemented"; return }

// rebalance performs one rebalance cycle:
// - Collects demand from all user pools
// - Computes fair allocations
// - Applies new capacities
// - Garbage collects inactive pools
func (m *Manager) rebalance(ctx context.Context) { _ = "STUB: not implemented"; return }

// 1. Collect demands from all user pools

// 2. Compute fair allocations

// 3. Apply new capacities to each pool

// 4. Garbage collect inactive pools

// garbageCollectInactivePools removes user pools that have been inactive
// longer than the configured timeout.
func (m *Manager) garbageCollectInactivePools(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

// Find inactive pools

// Remove inactive pools using copy-on-write

// Re-read snapshot with lock held

// Create new map without inactive pools

// Double-check activity timestamp (may have been updated since first check)

// Close and remove the pool
