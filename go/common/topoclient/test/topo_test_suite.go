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

// LocalCellName is the cell name used by this test suite.
const LocalCellName = "test"

// TopoServerTestSuite runs the full topoclient.Server/Conn test suite.
// The factory method should return a topoclient.Server that has a single cell
// called LocalCellName.
func TopoServerTestSuite(t *testing.T, ctx context.Context, factory func() topoclient.Store) {
	_ = "STUB: not implemented"
	return

	// Lock and TryLock are part of the Lock API.
}

// Directory is part of the Directory API.

// Watch and WatchRecursive are part of the Watch API.

// File is part of the File API.

// ShardInitClaim is part of the shard initialization API.
