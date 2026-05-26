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

package manager

import (
	"context"

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
)

// ============================================================================
// Multigres Schema Operations
//
// This file contains methods for managing the multigres sidecar schema and
// its tables. These are operations that set up and maintain the multigres
// metadata within PostgreSQL.
// ============================================================================

// ----------------------------------------------------------------------------
// Schema Creation
// ----------------------------------------------------------------------------

// createSidecarSchema creates the multigres sidecar schema and all its tables.
//
// MVP Limitation: Currently, we only support the default tablegroup. This function
// validates that the multipooler is configured for the default tablegroup and will
// return an error otherwise.
//
// For the default tablegroup, this function also creates the multischema global
// tables (tablegroup, tablegroup_table, shard).
func (pm *MultiPoolerManager) createSidecarSchema(ctx context.Context, policy *clustermetadatapb.DurabilityPolicy) error {
	_ = "STUB: not implemented"
	return nil
}

// Create multischema global tables for the default tablegroup

// initializeMultischemaData inserts the initial tablegroup and shard records.
//
// MVP Limitation: Currently, we only support the default tablegroup with shard "0-inf".
// This function validates these constraints and returns an error otherwise.
//
// TODO: In the future, tablegroup and shard insertion should be done via a dedicated
// RPC, and the bootstrap code should insert the tablegroup in the default primary
// pooler. For simplicity in the MVP, we do this as part of InitializePrimary since
// we only support a single tablegroup/shard for now.
func (pm *MultiPoolerManager) initializeMultischemaData(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// MVP validation: only default tablegroup with shard 0-inf is supported
// This is an extra guardrail. Multipoolers shouldn't start unless they
// are in the default tablegroup. However, we shouldn't be calling this function
// by the time we support multiple tablegroups/shards.
// This will ensure we make sure to remove this code when we get to that point.

// createSchema creates the multigres schema if it doesn't exist
func (pm *MultiPoolerManager) createSchema(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// ----------------------------------------------------------------------------
// Table Creation
// ----------------------------------------------------------------------------

// createHeartbeatTable creates the heartbeat table for leader election
func (pm *MultiPoolerManager) createHeartbeatTable(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// ----------------------------------------------------------------------------
// Multischema Global Tables (default tablegroup only)
// ----------------------------------------------------------------------------

// createTablegroup creates the tablegroup table for tracking table groups
func (pm *MultiPoolerManager) createTablegroup(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// createTablegroupTable creates the tablegroup_table table for tracking tables within tablegroups
func (pm *MultiPoolerManager) createTablegroupTable(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// createShard creates the shard table for tracking shards within tablegroups
func (pm *MultiPoolerManager) createShard(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// ----------------------------------------------------------------------------
// Data Operations
// ----------------------------------------------------------------------------

// insertTablegroup inserts a tablegroup record into the tablegroup table.
// Uses ON CONFLICT DO NOTHING to handle concurrent insertions gracefully.
// The type is hardcoded to "unsharded" for the MVP.
func (pm *MultiPoolerManager) insertTablegroup(ctx context.Context, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// insertShard inserts a shard record into the shard table.
// Returns an error if the tablegroup doesn't exist.
// Uses ON CONFLICT DO NOTHING on (tablegroup_oid, shard_name) to handle concurrent insertions gracefully.
func (pm *MultiPoolerManager) insertShard(ctx context.Context, tablegroupName string, shardName string) error {
	_ = "STUB: not implemented"
	return nil
}

// First, fetch the tablegroup oid

// Insert the shard
