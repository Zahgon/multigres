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

// This file provides the utility methods to save / retrieve Cell
// in the topology server.
//
// Cell records are not meant to be changed while the system is
// running.  In a running system, a Cell can be added, and
// topology server implementations should be able to read them to
// access the cells upon demand. Topology server implementations can
// also read the available Cell at startup to build a list of
// available cells, if necessary. A Cell can only be removed if no
// Shard record references the corresponding cell in its Cells list.

func pathForCell(cell string) string { _ = "STUB: not implemented"; return "" }

// GetCellNames returns the names of the existing cells. They are
// sorted by name.
func (ts *store) GetCellNames(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/*full*/

// GetCell reads a Cell from the global Conn.
func (ts *store) GetCell(ctx context.Context, cell string) (*clustermetadatapb.Cell, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Read the file.

// Unpack the contents.

// CreateCell creates a new Cell with the provided content.
func (ts *store) CreateCell(ctx context.Context, cell string, ci *clustermetadatapb.Cell) error {
	_ = "STUB: not implemented"
	return nil
}

// Pack the content.

// Save it.

// UpdateCellFields is a high level helper method to read a Cell
// object, update its fields, and then write it back.  If the write fails due to
// a version mismatch, it will re-read the record and retry the update.
// If the update method returns ErrNoUpdateNeeded, nothing is written,
// and nil is returned.
func (ts *store) UpdateCellFields(ctx context.Context, cell string, update func(*clustermetadatapb.Cell) error) error {
	_ = "STUB: not implemented"
	return nil
}

// Read the file, unpack the contents.

// Nothing to do.

// Call update method.

// Pack and save.

// This includes the 'err=nil' case.

// DeleteCell deletes the specified Cell.
// We first try to make sure no Shard record points to the cell,
// but we'll continue regardless if 'force' is true.
func (ts *store) DeleteCell(ctx context.Context, cell string, force bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if this cell is being used in any database before deleting it.

// Check if this database references the cell to be deleted
