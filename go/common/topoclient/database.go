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

// This file provides the utility methods to save / retrieve Database
// in the topology server.
//

// pathForDatabase returns the path for a database in the topology.
func pathForDatabase(database string) string { _ = "STUB: not implemented"; return "" }

// GetDatabaseNames returns the names of the existing databases. They are
// sorted by name.
func (ts *store) GetDatabaseNames(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/*full*/

// GetDatabase reads a Database from the global Conn.
func (ts *store) GetDatabase(ctx context.Context, database string) (*clustermetadatapb.Database, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Read the file.

// Unpack the contents.

// CreateDatabase creates a new Database with the provided content.
func (ts *store) CreateDatabase(ctx context.Context, database string, db *clustermetadatapb.Database) error {
	_ = "STUB: not implemented"
	return nil
}

// Pack the content.

// Save it.

// UpdateDatabaseFields is a high level helper method to read a Database
// object, update its fields, and then write it back. If the write fails due to
// a version mismatch, it will re-read the record and retry the update.
// If the update method returns ErrNoUpdateNeeded, nothing is written,
// and nil is returned.
func (ts *store) UpdateDatabaseFields(ctx context.Context, database string, update func(*clustermetadatapb.Database) error) error {
	_ = "STUB: not implemented"
	return nil
}

// Read the file, unpack the contents.

// Nothing to do.

// Call update method.

// Pack and save.

// DeleteDatabase deletes the specified Database.
// We first try to make sure no other records reference this database,
// but we'll continue regardless if 'force' is true.
func (ts *store) DeleteDatabase(ctx context.Context, database string, force bool) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Check if this database is being used by any MultiPooler records before deleting it.
