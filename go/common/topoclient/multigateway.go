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

// NewMultiGateway creates a new MultiGateway record with the given name, cell, and hostname.
func NewMultiGateway(name string, cell, host string) *clustermetadatapb.MultiGateway {
	_ = "STUB: not implemented"
	return nil
}

// MultiGatewayInfo is the container for a MultiGateway, read from the topology server.
type MultiGatewayInfo struct {
	version Version // node version - used to prevent stomping concurrent writes
	*clustermetadatapb.MultiGateway
}

// String returns a string describing the multigateway.
func (mgi *MultiGatewayInfo) String() string { _ = "STUB: not implemented"; return "" }

// IDString returns the string representation of the multigateway id
func (mgi *MultiGatewayInfo) IDString() string { _ = "STUB: not implemented"; return "" }

// Addr returns hostname:grpc port.
func (mgi *MultiGatewayInfo) Addr() string { _ = "STUB: not implemented"; return "" }

// Version returns the version of this multigateway from last time it was read or updated.
func (mgi *MultiGatewayInfo) Version() Version {
	_ = "STUB: not implemented"

	// NewMultiGatewayInfo returns a MultiGatewayInfo based on multigateway with the
	// version set. This function should be only used by Server implementations.
	return *new(Version)
}

func NewMultiGatewayInfo(multigateway *clustermetadatapb.MultiGateway, version Version) *MultiGatewayInfo {
	_ = "STUB: not implemented"
	return nil
}

// MultiGatewayIDString returns the string representation of a MultiGateway ID
func MultiGatewayIDString(id *clustermetadatapb.ID) string { _ = "STUB: not implemented"; return "" }

// GetMultiGateway is a high level function to read multigateway data.
func (ts *store) GetMultiGateway(ctx context.Context, id *clustermetadatapb.ID) (*MultiGatewayInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetMultiGatewayIDsByCell returns all the multigateway IDs in a cell.
// It returns ErrNoNode if the cell doesn't exist.
// It returns (nil, nil) if the cell exists, but there are no multigateways in it.
func (ts *store) GetMultiGatewayIDsByCell(ctx context.Context, cell string) ([]*clustermetadatapb.ID, error) {
	_ = "STUB: not implemented"
	// If the cell doesn't exist, this will return ErrNoNode.
	return nil, nil
}

// List the directory, and parse the IDs

// directory doesn't exist, empty list, no error.

// GetMultiGatewaysByCell returns all the multigateways in the cell.
// It returns ErrNoNode if the cell doesn't exist.
// It returns ErrPartialResult if some multigateways couldn't be read. The results in the slice are incomplete.
// It returns (nil, nil) if the cell exists, but there are no multigateways in it.
func (ts *store) GetMultiGatewaysByCell(ctx context.Context, cellName string) ([]*MultiGatewayInfo, error) {
	_ = "STUB: not implemented"
	// If the cell doesn't exist, this will return ErrNoNode.
	return nil, nil
}

// UpdateMultiGateway updates the multigateway data only - not associated replication paths.
func (ts *store) UpdateMultiGateway(ctx context.Context, mgi *MultiGatewayInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateMultiGatewayFields is a high level helper to read a multigateway record, call an
// update function on it, and then write it back. If the write fails due to
// a version mismatch, it will re-read the record and retry the update.
// If the update succeeds, it returns the updated multigateway.
// If the update method returns ErrNoUpdateNeeded, nothing is written,
// and nil,nil is returned.
func (ts *store) UpdateMultiGatewayFields(ctx context.Context, id *clustermetadatapb.ID, update func(*clustermetadatapb.MultiGateway) error) (*clustermetadatapb.MultiGateway, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateMultiGateway creates a new multigateway and all associated paths.
func (ts *store) CreateMultiGateway(ctx context.Context, mtgateway *clustermetadatapb.MultiGateway) error {
	_ = "STUB: not implemented"
	return nil
}

// UnregisterMultiGateway deletes the specified multigateway.
func (ts *store) UnregisterMultiGateway(ctx context.Context, id *clustermetadatapb.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterMultiGateway creates or updates a multigateway. If allowUpdate is true,
// and a multigateway with the same ID exists, just update it.
func (ts *store) RegisterMultiGateway(ctx context.Context, mtgateway *clustermetadatapb.MultiGateway, allowUpdate bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Try to update then
