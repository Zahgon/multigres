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

package command

import (
	"context"
	"log/slog"
	"regexp"

	"github.com/multigres/multigres/go/tools/retry"
)

// postgresAlreadyRunningPattern matches the postgres error reported when the
// postmaster.pid lock file is held. After a postmaster crash (kill -9 of a single
// PID, OOM, segfault), orphaned worker processes (writer, checkpointer,
// walreceiver, bgwriter) keep the SHM segment attached for ~1-5s while they
// detect parent death via PostmasterIsAlive() and exit. The same error is
// reported during that window even though postgres is not actually running.
var postgresAlreadyRunningPattern = regexp.MustCompile(`lock file ".*" already exists`)

// isPostgresCleanlyStopped checks if PostgreSQL is in a clean shutdown state.
// Returns true if state is "shut down" or "shut down in recovery", false otherwise.
func isPostgresCleanlyStopped(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Clean states: "shut down", "shut down in recovery"
// Anything else means we should try crash recovery

// extractClusterState extracts the cluster state from pg_controldata output
func extractClusterState(output string) string { _ = "STUB: not implemented"; return "" }

// Format: "Database cluster state:               in production"

// runCrashRecovery performs crash recovery in single-user mode.
// This runs postgres --single to complete crash recovery, then exits cleanly.
func runCrashRecovery(ctx context.Context, logger *slog.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

// runCrashRecoveryAttempts retries `postgres --single` while the lock file is held.
// During the orphan-cleanup window after a postmaster crash, the lock will eventually
// release; if it does not within the retry window, postgres is genuinely running and
// we preserve the historical no-op behavior. Extracted for unit-test injection.
func runCrashRecoveryAttempts(
	ctx context.Context,
	logger *slog.Logger,
	run func(context.Context) ([]byte, error),
	r *retry.Retry,
) error {
	_ = "STUB: not implemented"
	return nil
}

// runSingleUserPostgres runs `postgres --single` once and returns its combined
// output and exit error. /dev/null on stdin causes single-user mode to perform
// recovery and exit on EOF.
func runSingleUserPostgres(ctx context.Context) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
