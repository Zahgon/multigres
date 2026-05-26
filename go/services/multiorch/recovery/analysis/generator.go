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

package analysis

import (
	"time"

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	multiorchdatapb "github.com/multigres/multigres/go/pb/multiorchdata"
	"github.com/multigres/multigres/go/services/multiorch/store"
)

// DefaultReplicaLagThreshold is the threshold above which a replica is considered lagging.
const DefaultReplicaLagThreshold = 10 * time.Second

// replicationHeartbeatStalenessMultiplier is applied to wal_receiver_status_interval
// to compute the heartbeat staleness threshold. The replica sends a status message
// to the primary every wal_receiver_status_interval; the primary echoes a keepalive
// reply. Three missed intervals means the primary has gone silent well before the
// wal_receiver_timeout (60s) would disconnect the WAL receiver.
const replicationHeartbeatStalenessMultiplier = 3

// defaultReplicationHeartbeatStalenessThreshold is the fallback threshold used
// when wal_receiver_status_interval is not available in the replica's health
// state. Equals replicationHeartbeatStalenessMultiplier × the default
// wal_receiver_status_interval (10s).
const defaultReplicationHeartbeatStalenessThreshold = 30 * time.Second

// PoolersByShard is a structured map for efficient lookups.
// Structure: [database][tablegroup][shard][pooler_id] -> PoolerHealthState
type PoolersByShard map[string]map[string]map[string]map[string]*multiorchdatapb.PoolerHealthState

// AnalysisGenerator creates ReplicationAnalysis from the pooler store.
type AnalysisGenerator struct {
	poolerStore    *store.PoolerStore
	poolersByShard PoolersByShard
	// policyLookup returns the bootstrap durability policy for a database name.
	// May be nil; when nil, ShardAnalysis.BootstrapDurabilityPolicy is left nil.
	policyLookup func(database string) *clustermetadatapb.DurabilityPolicy
	now          func() time.Time
}

// NewAnalysisGenerator creates a new analysis generator.
// It eagerly builds the poolersByShard map from the current store state.
// policyLookup is optional; pass nil if the bootstrap policy is unavailable.
func NewAnalysisGenerator(poolerStore *store.PoolerStore, policyLookup func(database string) *clustermetadatapb.DurabilityPolicy) *AnalysisGenerator {
	_ = "STUB: not implemented"
	return nil
}

// GenerateShardAnalyses groups per-pooler analyses into one ShardAnalysis per shard.
func (g *AnalysisGenerator) GenerateShardAnalyses() []*ShardAnalysis {
	_ = "STUB: not implemented"
	return nil
}

// GenerateShardAnalysis returns a ShardAnalysis for a specific shard.
// Returns an error if no poolers for that shard are found in the store.
func (g *AnalysisGenerator) GenerateShardAnalysis(shardKey *clustermetadatapb.ShardKey) (*ShardAnalysis, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// buildShardAnalysis constructs a ShardAnalysis for a shard, including shard-level aggregates.
func (g *AnalysisGenerator) buildShardAnalysis(shardKey *clustermetadatapb.ShardKey, poolers map[string]*multiorchdatapb.PoolerHealthState) *ShardAnalysis {
	_ = "STUB: not implemented"
	return nil
}

// buildPoolersByShard creates a structured map by iterating the store once.
// Since ProtoStore.Range() returns clones, we don't need explicit DeepCopy.
func (g *AnalysisGenerator) buildPoolersByShard() PoolersByShard {
	_ = "STUB: not implemented"
	return *new(PoolersByShard)
}

// skip nil entries

// Initialize nested maps if needed

// Store the pooler (already a clone from Range)

// continue

// GetPoolersInShard returns all pooler IDs in the same shard as the given pooler.
// Uses the cached poolersByShard for efficient lookup.
func (g *AnalysisGenerator) GetPoolersInShard(poolerIDStr string) ([]string, error) {
	_ = "STUB: not implemented"
	// Get pooler from store to determine its shard
	return nil, nil
}

// Use cached poolersByShard for efficient lookup

// GenerateAnalysisForPooler generates and returns the ShardAnalysis for the shard containing
// the given pooler ID. Used primarily in tests to inspect shard-level fields like
// ReplicasConnectedToLeader without running the full analysis loop.
func (g *AnalysisGenerator) GenerateAnalysisForPooler(poolerIDStr string) (*ShardAnalysis, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// generateAnalysisForPooler creates a ReplicationAnalysis for a single pooler.
func (g *AnalysisGenerator) generateAnalysisForPooler(
	pooler *multiorchdatapb.PoolerHealthState,
	shardKey *clustermetadatapb.ShardKey,
) *PoolerAnalysis {
	_ = "STUB: not implemented"
	// Determine pooler type from health check (PoolerType).
	// Nodes are never created with topology type PRIMARY, so health check is authoritative.
	// Fall back to topology type only if health check type is UNKNOWN.
	return nil
}

// Compute staleness

// Store consensus status.

// If this is a REPLICA, populate replica-specific fields

// Extract primary connection info

// findHighestTermRawLeader returns the raw PoolerHealthState with the highest LeaderTerm
// among all known leader poolers, regardless of reachability. Returns nil if none found.
//
// A pooler is a candidate if:
//   - its ConsensusStatus names it as leader (IsLeader), OR
//   - its health status reports PoolerType=PRIMARY (fallback when ConsensusStatus is absent,
//     e.g. before the first streaming snapshot populates it, or after a BeginTerm REVOKE
//     where the node still reports PRIMARY while postgres restarts as standby).
//
// Note: we do NOT use MultiPooler.Type (topology type) because topology can be stale when
// etcd is unavailable — topology type reflects the last etcd write, which may be the initial
// assignment rather than the current leader's type.
func findHighestTermRawLeader(poolers map[string]*multiorchdatapb.PoolerHealthState) *multiorchdatapb.PoolerHealthState {
	_ = "STUB: not implemented"
	// TODO: If multiple poolers claim to be leader at the same term, we should surface an error that
	// manual intervention is needed.
	return nil
}

// allReplicasConnectedToLeader checks if ALL postgres standbys in the shard are connected to the leader's postgres.
// A replica is considered connected if:
// 1. Its health check is valid (IsLastCheckValid)
// 2. It has PrimaryConnInfo configured pointing to the leader's postgres
// 3. It has received WAL (LastReceiveLsn is not empty)
//
// Returns true only if all replicas meet these criteria.
// Returns false if there are no replicas or any replica is disconnected.
func (g *AnalysisGenerator) allReplicasConnectedToLeader(
	primary *multiorchdatapb.PoolerHealthState,
	poolers map[string]*multiorchdatapb.PoolerHealthState,
) bool {
	_ = "STUB: not implemented"
	return false
}

// Skip the leader itself

// Skip non-replicas. Note: a node whose postgres crashed into recovery mode
// without going through the normal resign flow could report PoolerType_REPLICA
// while still being the consensus leader. Such a node would be incorrectly
// counted as a follower here, overstating connected-follower count.

// Check if replica is connected to the leader's postgres

// All replicas must be connected (and there must be at least one replica)

// isFollowerConnectedToLeader checks if a single replica is actively connected to the leader's postgres.
// It verifies both that the connection is configured correctly and that the WAL receiver is
// actively exchanging keepalives with the leader's postgres via pg_stat_wal_receiver.
func (g *AnalysisGenerator) isFollowerConnectedToLeader(
	replica *multiorchdatapb.PoolerHealthState,
	primaryHost string,
	primaryPort int32,
) bool {
	_ = "STUB: not implemented"
	// Replica must be reachable
	return false
}

// Replica must have replication status

// Replica must have PrimaryConnInfo pointing to the primary

// Verify the replica is pointing to the correct primary. Note: if this is
// not the case, there is a more fundamental problem (e.g., misconfiguration
// or split-brain). This is not correctly indicated by a simple "false"
// return value, but we still want to return false here to avoid falsely
// triggering failover analyzers that rely on this method.

// Replica must have received WAL (indicates connection was established)

// WAL receiver must be in streaming state

// If last_msg_receive_time is available, verify the leader's postgres is still
// sending keepalives. The threshold is
// replicationHeartbeatStalenessMultiplier × wal_receiver_status_interval,
// falling back to defaultReplicationHeartbeatStalenessThreshold when the
// interval is unknown.
//
// If the last heartbeat is older than WAL receiver timeout, the connection
// is effectively dead even if the replica hasn't noticed yet, so we check
// that as well.

// computeShardLevelFields populates shard-level aggregates on sa after all per-pooler
// analyses have been built. These fields describe the shard as a whole rather than
// any individual pooler, so they are computed once here rather than per-pooler.
func (g *AnalysisGenerator) computeShardLevelFields(sa *ShardAnalysis, poolers map[string]*multiorchdatapb.PoolerHealthState) {
	_ = "STUB: not implemented"
	// Bootstrap durability policy lookup.
	return
}

// Count reachable, initialized poolers for bootstrap analysis.

// Collect all reachable primaries in the shard.

// Determine the highest-term reachable leader (used for stale-leader detection).

// Compute the highest-term leader in the shard regardless of reachability.
// This may differ from HighestTermReachableLeader when the leader pooler is down.

// LeaderHasResigned: AvailabilityStatus and ConsensusTerm are populated from
// StatusResponse on every health stream snapshot, so LeaderNeedsReplacement
// correctly detects REQUESTING_DEMOTION signals without a separate RPC.

// LeaderReachable requires the topology leader to be serving as PRIMARY and
// not have resigned. A resigned leader has voluntarily stepped down;
// treating it as reachable would prevent LeaderIsDead detection even when
// postgres is still running on the demoted node.

// Populate the standby list from the topology primary (used by IsInStandbyList).

// Detect pg_promote transition: multipooler explicitly signals promotion is running.

// HasInitializedReplica: any non-primary, reachable, initialized pooler.

// Determine if all followers are still connected to the leader's postgres.
// Use the highest-term discovered leader (which may be unreachable) so we can detect the
// "pooler down but postgres still running" scenario that ReplicasConnectedToLeader
// is designed to catch.

// findHighestTermLeader returns the leader PoolerAnalysis with the highest LeaderTerm.
// Returns nil if leaders is empty, all have LeaderTerm=0, or there is a tie.
//
// Invariant: In a properly initialized shard, LeaderTerm is always >0 for PRIMARY poolers.
// LeaderTerm is set during promotion and only cleared during demotion. This function
// is defensive and returns nil if all leaders have LeaderTerm=0, but this should
// never happen in a properly initialized shard.
func findHighestTermLeader(primaries []*PoolerAnalysis) *PoolerAnalysis {
	_ = "STUB: not implemented"
	return nil
}

// Defensive: should not happen in initialized shards, but guard against invalid state

// Tie detection: multiple leaders with the same LeaderTerm indicates a consensus bug.
// LeaderTerm should be unique per leader and monotonically increasing. If two leaders
// claim the same LeaderTerm, something went wrong in the consensus protocol (bug in
// promotion logic, data corruption, or split-brain).
//
// TODO: Rather than requiring manual intervention, multiorch could automatically resolve
// this by starting a new term and reappointing one of the leaders, which would update
// its leader_term and make the others stale. For now, we skip automatic demotion to
// avoid making the situation worse without understanding the root cause.
