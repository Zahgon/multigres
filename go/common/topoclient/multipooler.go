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

package topoclient

import (
	"context"

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
)

// NewMultiPooler creates a new MultiPooler record with the given name, cell, hostname, and tableGroup.
func NewMultiPooler(name string, cell, host, tableGroup string) *clustermetadatapb.MultiPooler {
	_ = "STUB: not implemented"
	return nil
}

// MultiPoolerInfo is the container for a MultiPooler, read from the topology server.
type MultiPoolerInfo struct {
	version Version // node version - used to prevent stomping concurrent writes
	*clustermetadatapb.MultiPooler
}

// String returns a string describing the multipooler.
func (mpi *MultiPoolerInfo) String() string { _ = "STUB: not implemented"; return "" }

// IDString returns the string representation of the multipooler id
func (mpi *MultiPoolerInfo) IDString() string { _ = "STUB: not implemented"; return "" }

// Addr returns hostname:grpc port.
func (mpi *MultiPoolerInfo) Addr() string { _ = "STUB: not implemented"; return "" }

// Version returns the version of this multipooler from last time it was read or updated.
func (mpi *MultiPoolerInfo) Version() Version {
	_ = "STUB: not implemented"

	// NewMultiPoolerInfo returns a MultiPoolerInfo based on multipooler with the
	// version set. This function should be only used by Server implementations.
	return *new(Version)
}

func NewMultiPoolerInfo(multipooler *clustermetadatapb.MultiPooler, version Version) *MultiPoolerInfo {
	_ = "STUB: not implemented"
	return nil
}

// MultiPoolerIDString returns the string representation of a MultiPooler ID
func MultiPoolerIDString(id *clustermetadatapb.ID) string { _ = "STUB: not implemented"; return "" }

// PoolerAddressFor projects a MultiPooler into the contact-info subset the
// consensus RPCs (SetTermPrimary, Propose) take. Returns nil if mp is nil.
func PoolerAddressFor(mp *clustermetadatapb.MultiPooler) *clustermetadatapb.PoolerAddress {
	_ = "STUB: not implemented"
	return nil
}

// GetMultiPooler is a high level function to read multipooler data.
func (ts *store) GetMultiPooler(ctx context.Context, id *clustermetadatapb.ID) (*MultiPoolerInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetMultiPoolerIDsByCell returns all the multipooler IDs in a cell.
// It returns ErrNoNode if the cell doesn't exist.
// It returns (nil, nil) if the cell exists, but there are no multipoolers in it.
func (ts *store) GetMultiPoolerIDsByCell(ctx context.Context, cell string) ([]*clustermetadatapb.ID, error) {
	_ = "STUB: not implemented"
	// If the cell doesn't exist, this will return ErrNoNode.
	return nil, nil
}

// List the directory, and parse the IDs

// directory doesn't exist, empty list, no error.

// GetMultiPoolersByCellOptions controls the behavior of GetMultiPoolersByCell.
type GetMultiPoolersByCellOptions struct {
	// DatabaseShard is the optional database/tablegroup/shard that multipoolers must match.
	// An empty tablegroup value will match all tablegroups in the database.
	// An empty shard value will match all shards in the tablegroup.
	DatabaseShard *DatabaseShard
}

// DatabaseShard represents a database, tablegroup, and shard tuple for filtering.
// Supports hierarchical matching:
// - Database only: matches all tablegroups and shards in database
// - Database + TableGroup: matches all shards in that tablegroup
// - Database + TableGroup + Shard: matches only that specific shard
type DatabaseShard struct {
	Database   string
	TableGroup string // empty = all tablegroups in database
	Shard      string // empty = all shards in tablegroup
}

// GetMultiPoolersByCell returns all the multipoolers in the cell.
// It returns ErrNoNode if the cell doesn't exist.
// It returns ErrPartialResult if some multipoolers couldn't be read. The results in the slice are incomplete.
// It returns (nil, nil) if the cell exists, but there are no multipoolers in it.
func (ts *store) GetMultiPoolersByCell(ctx context.Context, cellName string, opt *GetMultiPoolersByCellOptions) ([]*MultiPoolerInfo, error) {
	_ = "STUB: not implemented"
	// Validate filtering hierarchy: Database -> TableGroup -> Shard
	return nil, nil
}

// If Shard is specified, TableGroup must be specified

// If TableGroup is specified, Database must be specified

// If the cell doesn't exist, this will return ErrNoNode.

// Database must match

// If TableGroup is specified, it must match

// If Shard is specified, it must match

// UpdateMultiPooler updates the multipooler data only - not associated replication paths.
func (ts *store) UpdateMultiPooler(ctx context.Context, mpi *MultiPoolerInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateMultiPoolerFields is a high level helper to read a multipooler record, call an
// update function on it, and then write it back. If the write fails due to
// a version mismatch, it will re-read the record and retry the update.
// If the update succeeds, it returns the updated multipooler.
// If the update method returns ErrNoUpdateNeeded, nothing is written,
// and nil,nil is returned.
func (ts *store) UpdateMultiPoolerFields(ctx context.Context, id *clustermetadatapb.ID, update func(*clustermetadatapb.MultiPooler) error) (*clustermetadatapb.MultiPooler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateMultiPooler creates a new multipooler and all associated paths.
func (ts *store) CreateMultiPooler(ctx context.Context, mtpooler *clustermetadatapb.MultiPooler) error {
	_ = "STUB: not implemented"
	return nil
}

// UnregisterMultiPooler deletes the specified multipooler.
func (ts *store) UnregisterMultiPooler(ctx context.Context, id *clustermetadatapb.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterMultiPooler creates or updates a multipooler. If allowUpdate is true,
// and a multipooler with the same ID exists, just update it.
func (ts *store) RegisterMultiPooler(ctx context.Context, mtpooler *clustermetadatapb.MultiPooler, allowUpdate bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Try to update then
