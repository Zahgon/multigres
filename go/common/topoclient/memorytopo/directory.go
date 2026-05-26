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

// ListDir is part of the topoclient.Conn interface.
func (c *conn) ListDir(ctx context.Context, dirPath string, full bool) ([]topoclient.DirEntry, error) {
	_ = "STUB: not implemented"
	// c.factory.callstats.Add([]string{"ListDir"}, 1)
	return nil, nil
}

// Get the node to list.

// Check it's a directory.
