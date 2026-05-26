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

package memorytopo

import (
	"context"

	"github.com/multigres/multigres/go/common/topoclient"
)

// Below we implement the File methods that are part of the Conn interface:
//

// Create is part of topoclient.Conn interface.
func (c *conn) Create(ctx context.Context, filePath string, contents []byte) (topoclient.Version, error) {
	_ = "STUB: not implemented"
	// c.factory.callstats.Add([]string{"Create"}, 1)
	return *new(topoclient.Version), nil
}

// Get the parent dir.

// Check the file doesn't already exist.

// Create the file.

// Update is part of topoclient.Conn interface.
func (c *conn) Update(ctx context.Context, filePath string, contents []byte, version topoclient.Version) (topoclient.Version, error) {
	_ = "STUB: not implemented"
	//	c.factory.callstats.Add([]string{"Update"}, 1)
	return *new(topoclient.Version), nil
}

// Get the parent dir, we'll need it in case of creation.

// Parent doesn't exist, let's create it if we need to.

// Get the existing file.

// File doesn't exist, see if we need to create it.

// Check if it's a directory.

// Check the version.

// Now we can update.

// Call the watches

// Get is part of topoclient.Conn interface.
func (c *conn) Get(ctx context.Context, filePath string) ([]byte, topoclient.Version, error) {
	_ = "STUB: not implemented"
	// c.factory.callstats.Add([]string{"Get"}, 1)
	return nil, *new(topoclient.Version), nil
}

// Get the node.

// This matches the other topo implementations of returning topoclient.NoNode when calling
// Get() with a key prefix or "directory".

// GetVersion is part of topoclient.Conn interface.
func (c *conn) GetVersion(ctx context.Context, filePath string, version int64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List is part of the topoclient.Conn interface.
func (c *conn) List(ctx context.Context, filePathPrefix string) ([]topoclient.KVInfo, error) {
	_ = "STUB: not implemented"
	// c.factory.callstats.Add([]string{"List"}, 1)
	return nil, nil
}

// Get the node to list.

func gatherChildren(n *node, dirPath string) []topoclient.KVInfo {
	_ = "STUB: not implemented"
	return nil
}

// Delete is part of topoclient.Conn interface.
func (c *conn) Delete(ctx context.Context, filePath string, version topoclient.Version) error {
	_ = "STUB: not implemented"
	// c.factory.callstats.Add([]string{"Delete"}, 1)
	return nil
}

// Get the parent dir.

// Get the existing file.

// Check if it's a directory.

// Check the version.

// Now we can delete.

// Call the watches
