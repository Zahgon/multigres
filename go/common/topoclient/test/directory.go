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

package test

import (
	"context"
	"testing"

	"github.com/multigres/multigres/go/common/topoclient"
)

// checkDirectory tests the directory part of the topoclient.Conn API.
func checkDirectory(t *testing.T, ctx context.Context, ts topoclient.Store) {
	_ = "STUB: not implemented"
	// global topo
	return
}

/*hasCells*/

// local topo

/*hasCells*/

func checkListDir(ctx context.Context, t *testing.T, conn topoclient.Conn, dirPath string, expected []topoclient.DirEntry) {
	_ = "STUB: not implemented"

	// Build the shallow expected list, when full=false.
	return
}

// Test with full=false.
/*full*/

// Test with full=true.
/*full*/

func checkDirectoryInCell(t *testing.T, conn topoclient.Conn, hasCells bool) {
	_ = "STUB: not implemented"

	// ListDir root: nothing
	return
}

// Create a topolevel entry

// ListDir should return it.

// Delete it, it should be gone.

// Create a file 3 layers down.

// Check listing at all levels.

// Add a second file

// Check entries at all levels

// Delete the first file, expect all lists to return the second one.

// Delete the second file, expect all lists to return nothing.
