// Copyright 2025 Supabase, Inc.
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

package store

import (
	"context"
	"log/slog"

	"github.com/multigres/multigres/go/common/rpcclient"

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	multiorchdatapb "github.com/multigres/multigres/go/pb/multiorchdata"
)

// PoolerStore manages pooler health state and provides RPC-based domain queries.
type PoolerStore struct {
	health    *poolerHealthStore
	rpcClient rpcclient.MultiPoolerClient
	logger    *slog.Logger
}

// NewPoolerStore creates a new PoolerStore.
// rpcClient and logger are used by FindHealthyPrimary; they may be nil in tests
// that do not exercise that method.
func NewPoolerStore(rpcClient rpcclient.MultiPoolerClient, logger *slog.Logger) *PoolerStore {
	_ = "STUB: not implemented"
	return nil
}

// Get retrieves a pooler's health state by its ID string.
// Returns a deep clone safe to mutate, and false if the key does not exist.
func (s *PoolerStore) Get(poolerID string) (*multiorchdatapb.PoolerHealthState, bool) {
	_ = "STUB: not implemented"
	return nil, false

	// Set stores a deep clone of the pooler health state.
}

func (s *PoolerStore) Set(poolerID string, state *multiorchdatapb.PoolerHealthState) {
	_ = "STUB: not implemented"
	return
}

// Delete removes a pooler from the store. Returns true if the pooler existed.
func (s *PoolerStore) Delete(poolerID string) bool { _ = "STUB: not implemented"; return false }

// Len returns the number of poolers in the store.
func (s *PoolerStore) Len() int { _ = "STUB: not implemented"; return 0 }

// Range iterates over all poolers. Each value passed to the callback is a deep
// clone safe to mutate. Iteration stops early if the callback returns false.
func (s *PoolerStore) Range(fn func(key string, value *multiorchdatapb.PoolerHealthState) bool) {
	_ = "STUB: not implemented"
	return
}

// DoUpdate performs an atomic read-modify-write on a pooler's health state.
//
// The provided function receives the current value (or nil if not present) and
// returns the new value to store. This is useful for safely updating state
// based on the existing state without needing to do multiple Get/Set calls.
//
// Note that the function should not do any expensive or blocking calls since it
// is executed while holding the store lock.
func (s *PoolerStore) DoUpdate(key string, fn func(*multiorchdatapb.PoolerHealthState) *multiorchdatapb.PoolerHealthState) {
	_ = "STUB: not implemented"
	return
}

// DoUpdateRange iterates over all poolers while holding the lock and allows
// in-place updates.
//
// Each value passed to the callback is the raw internal value (not a clone).
// Return the updated value to write it back, or nil to leave it unchanged.
// Return false to stop iteration early, consistent with Range. The callback
// must not retain the pointer after it returns, and must not perform any
// expensive or blocking operations since it runs while holding the store lock.
//
// Example:
//
//	store.DoUpdateRange(func(key string, value *PoolerHealthState) (*PoolerHealthState, bool) {
//	    value.LastSeen = timestamppb.Now()
//	    return value, true // write and continue
//	})
func (s *PoolerStore) DoUpdateRange(fn func(key string, value *multiorchdatapb.PoolerHealthState) (*multiorchdatapb.PoolerHealthState, bool)) {
	_ = "STUB: not implemented"
	return
}

// IsInitialized returns true if the pooler has been initialized.
// FindPoolersInShard returns all poolers belonging to the given shard.
func (s *PoolerStore) FindPoolersInShard(shardKey *clustermetadatapb.ShardKey) []*multiorchdatapb.PoolerHealthState {
	_ = "STUB: not implemented"
	return nil
}

// continue

// continue

// FindPoolerByID finds a pooler in the store by its cell and name.
func (s *PoolerStore) FindPoolerByID(id *clustermetadatapb.ID) (*multiorchdatapb.PoolerHealthState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// continue

// stop iteration

// continue

// FindHealthyPrimary finds a healthy, initialized primary in the given pooler slice.
// It verifies health by making an RPC call to each candidate.
// Returns an error if multiple primaries are found (likely a stale primary that needs to be demoted).
//
// Candidate selection uses a union of topology type and live health-stream data because
// topology (from etcd) can be stale when etcd is unavailable after a failover. A pooler
// is considered a candidate if either:
//   - its topology type is PRIMARY (MultiPooler.Type), or
//   - its most recent health-stream snapshot reports it is running as PRIMARY (Status.PoolerType).
//
// Each candidate is then verified via Status RPC; only nodes whose live PoolerType
// is PRIMARY are accepted, so stale topology entries running as standby are skipped.
func (s *PoolerStore) FindHealthyPrimary(
	ctx context.Context,
	poolers []*multiorchdatapb.PoolerHealthState,
) (*multiorchdatapb.PoolerHealthState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Accept candidates indicated as PRIMARY by topology OR live health data.
// Topology can be stale when etcd is unavailable; health data can lag during
// role transitions. Using the union avoids missing the actual primary in either case.

// Verify via Status RPC — check the live PoolerType to skip stale candidates
// (e.g. topology says PRIMARY but postgres is running as standby after a failover).
