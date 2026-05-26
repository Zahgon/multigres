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

package multiorch

import (
	"database/sql"
	"testing"
	"time"

	_ "github.com/lib/pq"

	"github.com/multigres/multigres/go/test/endtoend/shardsetup"
)

// connectToPostgres establishes a connection to PostgreSQL using Unix socket
func connectToPostgres(t *testing.T, socketDir string, port int) *sql.DB {
	_ = "STUB: not implemented"
	return nil
}

// waitForReplicationBroken polls until the instance's primary_conninfo host is empty,
// indicating replication is no longer configured or streaming.
func waitForReplicationBroken(t *testing.T, inst *shardsetup.MultipoolerInstance, timeout time.Duration) {
	_ = "STUB: not implemented"
	return
}

// waitForShardReady polls until one node is an initialized primary, expectedStandbyCount nodes
// are initialized replicating standbys, and sync replication is configured on the primary.
// Returns the name of the elected primary.
func waitForShardReady(t *testing.T, setup *shardsetup.ShardSetup, expectedStandbyCount int, timeout time.Duration) string {
	_ = "STUB: not implemented"
	return ""
}

// Build lookup sets from the primary's sync config.
// short name (e.g. "pooler-2") -> true
// application name (e.g. "test-cell_pooler-2") -> true
