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

// checkFile tests the file part of the Conn API.
func checkFile(t *testing.T, ctx context.Context, ts topoclient.Store) {
	_ = "STUB: not implemented"
	// global cell
	return
}

/*hasCells*/

// local cell

/*hasCells*/

func checkFileInCell(t *testing.T, conn topoclient.Conn, hasCells bool) {
	_ = "STUB: not implemented"
	return

	// ListDir root: nothing.
}

// Get with no file -> ErrNoNode.

// Create a file.

// See it in the listing now.

// Get should work, get the right contents and version.

// Update it, make sure version changes.

// Get should work, get the right contents and version.

// Try to update again with wrong version, should fail.

// Try to update again with nil version, should work.

// Get should work, get the right contents and version.

// Try to update again with empty content, should work.

// Try to delete with wrong version, should fail.

// Now delete it.

// ListDir root: nothing.

// Try to delete again, should fail.

// Create again, with unconditional update.

// Check contents.

// See it in the listing now.

// Unconditional delete.

// ListDir root: nothing.

// checkList tests the file part of the Conn API.
func checkList(t *testing.T, ctx context.Context, ts topoclient.Store) {
	_ = "STUB: not implemented"
	// global topo
	return
}

// If this is not supported, skip the test
