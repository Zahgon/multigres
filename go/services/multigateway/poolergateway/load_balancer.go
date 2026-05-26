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

package poolergateway

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"google.golang.org/grpc"

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	"github.com/multigres/multigres/go/pb/query"
)

// shardKey identifies a shard for primary caching.
type shardKey struct {
	tableGroup string
	shard      string
}

// cachedPrimary holds the cached primary connection for a shard.
type cachedPrimary struct {
	conn *PoolerConnection
	term int64
}

// LoadBalancer manages PoolerConnections and selects connections for queries.
// It creates connections based on discovery events and destroys them when poolers
// are removed from discovery.
//
// For PRIMARY targets, the LoadBalancer caches the primary connection per-shard.
// The cache is updated via health stream callbacks, so the hot path (GetConnection)
// is a simple map lookup rather than iterating all poolers.
type LoadBalancer struct {
	// localCell is the cell where this gateway is running
	localCell string

	// logger for debugging
	logger *slog.Logger

	// ctx is the service-lifetime context for child goroutines (health streams)
	ctx context.Context

	// mu protects connections and cachedPrimaries
	mu sync.Mutex

	// connections maps pooler ID to PoolerConnection
	connections map[string]*PoolerConnection

	// cachedPrimaries maps shard key to the cached primary connection.
	// Updated by onPoolerHealthUpdate when health streams report LeaderObservation.
	cachedPrimaries map[shardKey]*cachedPrimary

	// onPrimaryServing is called when a new primary is detected via health stream.
	// Used to stop failover buffering for the shard. May be nil.
	onPrimaryServing func(tableGroup, shard string)

	// grpcDialOpt configures transport credentials (TLS or insecure) for pooler connections.
	grpcDialOpt grpc.DialOption

	// lowReplicationLagNs is the preferred replication lag threshold in nanoseconds.
	// Replicas at or below this lag are considered "healthy" and preferred.
	// If all replicas exceed this but are under highReplicationLagToleranceNs,
	// they are still eligible. Zero disables the preferred tier (all replicas equal).
	lowReplicationLagNs int64

	// highReplicationLagToleranceNs is the absolute maximum replication lag in
	// nanoseconds. Replicas above this are never selected. Zero means no upper
	// bound — any replica is eligible regardless of lag.
	highReplicationLagToleranceNs int64
}

// NewLoadBalancer creates a new LoadBalancer.
// The grpcDialOpt configures transport credentials for gRPC connections to poolers.
func NewLoadBalancer(ctx context.Context, localCell string, logger *slog.Logger, grpcDialOpt grpc.DialOption) *LoadBalancer {
	_ = "STUB: not implemented"
	return nil
}

// SetOnPrimaryServing sets a callback invoked when a new primary is detected
// via the streaming health check. This is used to stop failover buffering
// when a new primary becomes available, replacing the topology-based approach.
// Must be called before any poolers are added.
func (lb *LoadBalancer) SetOnPrimaryServing(fn func(tableGroup, shard string)) {
	_ = "STUB: not implemented"
	return
}

// SetReplicationLagThresholds configures two-tier replication lag filtering.
//
//   - lowLag: replicas at or below this lag are preferred ("healthy" tier).
//     If all replicas exceed this but are under highTolerance, they are still used.
//     Zero disables the preferred tier — all replicas are treated equally.
//   - highTolerance: absolute maximum lag. Replicas above this are never selected.
//     Zero means no upper bound.
func (lb *LoadBalancer) SetReplicationLagThresholds(lowLag, highTolerance time.Duration) {
	_ = "STUB: not implemented"
	return
}

// AddPooler creates a new PoolerConnection for the given pooler.
// If a connection already exists for this pooler, it updates the pooler info
// (e.g., when type changes from UNKNOWN to PRIMARY).
func (lb *LoadBalancer) AddPooler(pooler *clustermetadatapb.MultiPooler) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if already exists - update info instead of skipping

// Invalidate seed if type changed away from PRIMARY

// Seed if type changed to PRIMARY

// Seed primary cache from discovery type (term 0 = unconfirmed).
// Health stream callbacks will overwrite with real term data.

// RemovePooler closes and removes the PoolerConnection for the given pooler ID.
// If no connection exists for this pooler, it is a no-op.
func (lb *LoadBalancer) RemovePooler(poolerID string) { _ = "STUB: not implemented"; return }

// Invalidate cached primary if the removed pooler was the cached primary.

// Close outside the lock

// GetConnection returns a PoolerConnection matching the target specification.
// Returns an error immediately if no suitable connection is available.
//
// Selection logic:
// - For PRIMARY: uses cached primary (updated by health stream callbacks)
// - For REPLICA: prefers local cell serving replicas, with randomization
func (lb *LoadBalancer) GetConnection(target *query.Target) (*PoolerConnection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use cached primary (populated by health stream callbacks).
// The health stream's LeaderObservation is the authoritative source
// for leader identity — no type-based fallback.

// For REPLICA: collect only replica-type poolers

// All replicas exceeded the replication lag threshold.

// GetConnectionByID returns a PoolerConnection for a specific pooler ID.
// This is used for reserved connections where queries need to be routed to
// a specific pooler instance (e.g., for session affinity with prepared statements).
// Returns an error immediately if the pooler connection doesn't exist (fail-fast).
func (lb *LoadBalancer) GetConnectionByID(poolerID *clustermetadatapb.ID) (*PoolerConnection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// selectReplicaConnection chooses the best replica connection from candidates
// using two-tier replication lag filtering:
//
//  1. Exclude replicas above highReplicationLagToleranceNs (absolute max).
//  2. Prefer replicas at or below lowReplicationLagNs ("healthy" tier).
//  3. If no healthy replicas, fall back to any that passed the high tolerance.
//  4. If none pass either threshold, return nil (error to client).
//
// Within each tier, selection prefers local-cell serving replicas with
// randomization to distribute load.
func (lb *LoadBalancer) selectReplicaConnection(candidates []*PoolerConnection) *PoolerConnection {
	_ = "STUB: not implemented"
	return nil
}

// Single-pass: lag filtering + locality categorization combined.
// "healthy" = lag within lowThreshold (or unknown). "tolerable" = between
// low and high thresholds. Each bucket is further split by locality.

// Exclude replicas above the absolute maximum.

// Lag unknown (0) or within the low threshold → healthy.

// Between low and high thresholds → tolerable fallback.

// No lag filtering — all candidates go into "healthy".

// Pick the best bucket: prefer healthy, fall back to tolerable.

// All replicas exceed the high tolerance threshold.

// No lag filter and no candidates categorized — shouldn't happen
// because candidates is non-empty, but be safe.

// Select from tiers in preference order, with randomization within each tier.

// onPoolerHealthUpdate is the callback invoked by PoolerConnection when health
// state changes. It updates the cached primary for the connection's shard
// based on LeaderObservation term reconciliation.
//
// This is safe to call concurrently: both processHealthResponse and setHealthError
// release healthMu before invoking this callback.
func (lb *LoadBalancer) onPoolerHealthUpdate(conn *PoolerConnection) {
	_ = "STUB: not implemented"
	return
}

// New primary or higher term — update the cached primary.

// Only stop failover buffering when the primary is confirmed to be
// PRIMARY type and SERVING. The LeaderObservation can arrive before
// the pooler has transitioned its query server to PRIMARY/SERVING
// (e.g., during Promote, UpdateLeaderObservation fires before
// changeTypeLocked). Draining buffered requests too early would send
// them to a pooler that still rejects PRIMARY traffic.

// Same term — the primary is already cached but may not have been
// SERVING when we first saw the observation. Re-check now so that
// StopBuffering fires once the primary transitions to PRIMARY/SERVING.

// Stale term — ignore.

// notifyIfPrimaryServingLocked calls onPrimaryServing if the primary connection
// is confirmed to be PRIMARY type and SERVING. StopBuffering is idempotent, so
// calling this on every health update is safe and ensures buffering stops
// promptly once the primary is ready. Caller must hold lb.mu.
func (lb *LoadBalancer) notifyIfPrimaryServingLocked(key shardKey, primaryConn *PoolerConnection) {
	_ = "STUB: not implemented"
	return
}

// matchesShardTarget checks if a connection matches the tablegroup and shard,
// regardless of pooler type.
func matchesShardTarget(conn *PoolerConnection, target *query.Target) bool {
	_ = "STUB: not implemented"
	return false
}

// Check tablegroup match

// Check shard match (empty target shard matches any)

// matchesTarget checks if a connection matches the target specification.
func matchesTarget(conn *PoolerConnection, target *query.Target) bool {
	_ = "STUB: not implemented"
	return false
}

// Check type match

// poolerIDString returns the string ID for a pooler.
// Uses the same format as PoolerConnection.ID() for consistency.
func poolerIDString(id *clustermetadatapb.ID) string { _ = "STUB: not implemented"; return "" }

// ConnectionCount returns the number of active connections.
func (lb *LoadBalancer) ConnectionCount() int { _ = "STUB: not implemented"; return 0 }

// Close closes all connections.
func (lb *LoadBalancer) Close() error { _ = "STUB: not implemented"; return nil }

// LoadBalancerListener wraps a LoadBalancer to implement multigateway.PoolerChangeListener.
type LoadBalancerListener struct {
	lb *LoadBalancer
}

// NewLoadBalancerListener creates a listener adapter for the given LoadBalancer.
func NewLoadBalancerListener(lb *LoadBalancer) *LoadBalancerListener {
	_ = "STUB: not implemented"
	return nil
}

// OnPoolerChanged implements multigateway.PoolerChangeListener.
func (l *LoadBalancerListener) OnPoolerChanged(pooler *clustermetadatapb.MultiPooler) {
	_ = "STUB: not implemented"
	return
}

// OnPoolerRemoved implements multigateway.PoolerChangeListener.
func (l *LoadBalancerListener) OnPoolerRemoved(pooler *clustermetadatapb.MultiPooler) {
	_ = "STUB: not implemented"
	return
}
