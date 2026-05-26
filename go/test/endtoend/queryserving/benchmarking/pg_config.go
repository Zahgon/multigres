// Copyright 2026 Supabase, Inc.
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

package benchmarking

import (
	"context"
	"testing"

	"github.com/multigres/multigres/go/test/endtoend/shardsetup"
)

// bumpPostgresMaxConnections rewrites postgresql.auto.conf on every pgctld
// in the cluster to set max_connections=n, then restarts each postgres via
// pgctld so the new value takes effect.
//
// Required when running pgbench with more clients than the default
// postgres max_connections (60).
func bumpPostgresMaxConnections(ctx context.Context, t *testing.T, setup *shardsetup.ShardSetup, n int) {
	_ = "STUB: not implemented"
	return
}

// Disable the multipooler monitor's automatic postgres restart so it
// doesn't race with our pgctld Restart.

// Re-enable the monitor.

const maxConnectionsMarker = "# multigres_pgbench_override\n"

// writeMaxConnectionsOverride appends (idempotently) the override to
// postgresql.auto.conf inside the postgres data directory.
func writeMaxConnectionsOverride(poolerDir string, n int) error {
	_ = "STUB: not implemented"
	return nil
}

// If we already wrote the override, replace the line; otherwise append.

// stripPriorOverride removes the previous max_connections override block
// (marker + immediate next line) so we can rewrite cleanly. Anything else
// in postgresql.auto.conf is preserved.
func stripPriorOverride(in string) string { _ = "STUB: not implemented"; return "" }

// drop the next line (the actual setting)

func indexOf(s, sub string) int { _ = "STUB: not implemented"; return 0 }
