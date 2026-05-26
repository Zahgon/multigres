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

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
)

// waitForInitialValue waits for the initial value of
// databases/test_database/Database to appear, and match the
// provided database.
func waitForInitialValue(t *testing.T, conn topoclient.Conn, database *clustermetadatapb.Database) (changes <-chan *topoclient.WatchData, cancel context.CancelFunc) {
	_ = "STUB: not implemented"
	return nil, *new(context.CancelFunc)
}

// hasn't appeared yet

// we got a valid result

// waitForInitialValueRecursive waits for the initial value of
// databases/test_database. Any files that appear inside that directory
// will be watched. In this case will be waiting for the database to appear.
func waitForInitialValueRecursive(t *testing.T, conn topoclient.Conn, database *clustermetadatapb.Database) (changes <-chan *topoclient.WatchDataRecursive, cancel context.CancelFunc, err error) {
	_ = "STUB: not implemented"
	return nil, *new(context.CancelFunc), nil
}

// hasn't appeared yet

// If this is not supported, skip the test

// we got a valid result

// checkWatch runs the tests on the Watch part of the Conn API.
// We use a Database object.
func checkWatch(t *testing.T, ctx context.Context, ts topoclient.Store) {
	_ = "STUB: not implemented"
	return
}

// start watching something that doesn't exist -> error

// create some data

// start watching again, it should work

// change the data

// Make sure we get the watch data, maybe not as first notice,
// but eventually. The API specifies it is possible to get duplicate
// notifications.

// extra first value, still good

// watch worked, good

// remove the database

// Make sure we get the ErrNoNode notification eventually.
// The API specifies it is possible to get duplicate
// notifications.

// good

// we got something, better be the right value

// good value

// now the channel should be closed

// checkWatchInterrupt tests we can interrupt a watch.
func checkWatchInterrupt(t *testing.T, ctx context.Context, ts topoclient.Store) {
	_ = "STUB: not implemented"
	return
}

// create some data

// Start watching, it should work.

// Now cancel the watch.

// Make sure we get the topoclient.ErrInterrupted notification eventually.

// good

// we got something, better be the right value

// good value

// Now the channel should be closed.

// And calling cancel() again should just work.

// checkWatchRecursive tests we can setup a recursive watch
func checkWatchRecursive(t *testing.T, ctx context.Context, ts topoclient.Store) {
	_ = "STUB: not implemented"
	return
}

// create some data

// start watching again, it should work

// Skip the rest if there's no implementation

// change the data

// Make sure we get the watch data, maybe not as first notice,
// but eventually. The API specifies it is possible to get duplicate
// notifications.

// extra first value, still good

// watch worked, good

// remove the database

// Make sure we get the ErrNoNode notification eventually.
// The API specifies it is possible to get duplicate
// notifications.

// good

// we got something, better be the right value

// good value

// We now have to stop watching. This doesn't automatically
// happen for recursive watches on a single file since others
// can still be seen.

// Make sure we get the topoclient.ErrInterrupted notification eventually.

// good

// we got something, better be the right value

// good value

// Now the channel should be closed.

// And calling cancel() again should just work.
