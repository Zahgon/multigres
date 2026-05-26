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

// Watch is part of the topoclient.Conn interface.
func (c *conn) Watch(ctx context.Context, filePath string) (*topoclient.WatchData, <-chan *topoclient.WatchData, error) {
	_ = "STUB: not implemented"
	// c.factory.callstats.Add([]string{"Watch"}, 1)
	return nil, nil, nil
}

// it's a directory

// Create a cancellable context so we can forcibly close the watch

// This function can be called at any point, so we first need
// to make sure the watch is still valid.

// WatchRecursive is part of the topoclient.Conn interface.
func (c *conn) WatchRecursive(ctx context.Context, dirpath string) ([]*topoclient.WatchDataRecursive, <-chan *topoclient.WatchDataRecursive, error) {
	_ = "STUB: not implemented"
	// c.factory.callstats.Add([]string{"WatchRecursive"}, 1)
	return nil, nil, nil
}

// Create a cancellable context so we can forcibly close the watch
