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

// NewMultiOrch creates a new MultiOrch record with the given name, cell, and hostname.
func NewMultiOrch(name string, cell, host string) *clustermetadatapb.MultiOrch {
	_ = "STUB: not implemented"
	return nil
}

// MultiOrchInfo is the container for a MultiOrch, read from the topology server.
type MultiOrchInfo struct {
	version Version // node version - used to prevent stomping concurrent writes
	*clustermetadatapb.MultiOrch
}

// String returns a string describing the multiorch.
func (moi *MultiOrchInfo) String() string { _ = "STUB: not implemented"; return "" }

// IDString returns the string representation of the multiorch id
func (moi *MultiOrchInfo) IDString() string { _ = "STUB: not implemented"; return "" }

// Addr returns hostname:grpc port.
func (moi *MultiOrchInfo) Addr() string { _ = "STUB: not implemented"; return "" }

// Version returns the version of this multiorch from last time it was read or updated.
func (moi *MultiOrchInfo) Version() Version {
	_ = "STUB: not implemented"

	// NewMultiOrchInfo returns a MultiOrchInfo based on multiorch with the
	// version set. This function should be only used by Server implementations.
	return *new(Version)
}

func NewMultiOrchInfo(multiorch *clustermetadatapb.MultiOrch, version Version) *MultiOrchInfo {
	_ = "STUB: not implemented"
	return nil
}

// MultiOrchIDString returns the string representation of a MultiOrch ID
func MultiOrchIDString(id *clustermetadatapb.ID) string { _ = "STUB: not implemented"; return "" }

// GetMultiOrch is a high level function to read multiorch data.
func (ts *store) GetMultiOrch(ctx context.Context, id *clustermetadatapb.ID) (*MultiOrchInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetMultiOrchIDsByCell returns all the multiorch IDs in a cell.
// It returns ErrNoNode if the cell doesn't exist.
// It returns (nil, nil) if the cell exists, but there are no multiorchs in it.
func (ts *store) GetMultiOrchIDsByCell(ctx context.Context, cell string) ([]*clustermetadatapb.ID, error) {
	_ = "STUB: not implemented"
	// If the cell doesn't exist, this will return ErrNoNode.
	return nil, nil
}

// List the directory, and parse the IDs

// directory doesn't exist, empty list, no error.

// GetMultiOrchsByCell returns all the multiorchs in the cell.
// It returns ErrNoNode if the cell doesn't exist.
// It returns ErrPartialResult if some multiorchs couldn't be read. The results in the slice are incomplete.
// It returns (nil, nil) if the cell exists, but there are no multiorchs in it.
func (ts *store) GetMultiOrchsByCell(ctx context.Context, cellName string) ([]*MultiOrchInfo, error) {
	_ = "STUB: not implemented"
	// If the cell doesn't exist, this will return ErrNoNode.
	return nil, nil
}

// UpdateMultiOrch updates the multiorch data only - not associated replication paths.
func (ts *store) UpdateMultiOrch(ctx context.Context, moi *MultiOrchInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateMultiOrchFields is a high level helper to read a multiorch record, call an
// update function on it, and then write it back. If the write fails due to
// a version mismatch, it will re-read the record and retry the update.
// If the update succeeds, it returns the updated multiorch.
// If the update method returns ErrNoUpdateNeeded, nothing is written,
// and nil,nil is returned.
func (ts *store) UpdateMultiOrchFields(ctx context.Context, id *clustermetadatapb.ID, update func(*clustermetadatapb.MultiOrch) error) (*clustermetadatapb.MultiOrch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateMultiOrch creates a new multiorch and all associated paths.
func (ts *store) CreateMultiOrch(ctx context.Context, mtorch *clustermetadatapb.MultiOrch) error {
	_ = "STUB: not implemented"
	return nil
}

// UnregisterMultiOrch deletes the specified multiorch.
func (ts *store) UnregisterMultiOrch(ctx context.Context, id *clustermetadatapb.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterMultiOrch creates or updates a multiorch. If allowUpdate is true,
// and a multiorch with the same ID exists, just update it.
func (ts *store) RegisterMultiOrch(ctx context.Context, mtorch *clustermetadatapb.MultiOrch, allowUpdate bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Try to update then
