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

package command

import (
	"context"
	"log/slog"
	"time"

	"github.com/multigres/multigres/go/services/pgctld"

	"github.com/spf13/cobra"
)

// PostgreSQL server status values
const (
	statusStopped = "STOPPED"
	statusRunning = "RUNNING"
)

const (
	// pgIsReadyDefaultTimeout is the connection timeout passed to pg_isready when
	// no context deadline is present. pg_isready's built-in default is also 3s,
	// but we pass it explicitly so the value is visible and intentional.
	pgIsReadyDefaultTimeout = 3 * time.Second

	// pgIsReadyDeadlineBuffer is subtracted from the remaining context deadline
	// before passing it to pg_isready via -t. This ensures pg_isready's own
	// timeout fires before the context cancels the subprocess mid-wait, avoiding
	// a race between libpq's connection timeout and executil's SIGTERM.
	// Must be >= 1s because -t only accepts whole seconds and the sub-second
	// remainder is truncated, so a smaller buffer may provide no margin at all.
	pgIsReadyDeadlineBuffer = 1 * time.Second
)

// StatusResult contains the result of checking PostgreSQL status
type StatusResult struct {
	Status        string // statusStopped, statusRunning
	PID           int
	Version       string
	UptimeSeconds int64
	DataDir       string
	Port          int
	Host          string
	Ready         bool
	Message       string
}

// PgCtlStatusCmd holds the status command configuration
type PgCtlStatusCmd struct {
	pgCtlCmd *PgCtlCommand
}

// AddStatusCommand adds the status subcommand to the root command
func AddStatusCommand(root *cobra.Command, pc *PgCtlCommand) { _ = "STUB: not implemented"; return }

func (s *PgCtlStatusCmd) createCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// GetStatusWithResult gets PostgreSQL status with the given configuration and returns detailed result information
func GetStatusWithResult(ctx context.Context, logger *slog.Logger, config *pgctld.PostgresCtlConfig) (*StatusResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if PostgreSQL is running

// Process exists — verify it is actually accepting connections.
// A process that exists but cannot respond (e.g. SIGSTOP, cgroup freeze)
// is treated as not running so that multipooler and multiorch can detect
// the failure and trigger recovery rather than waiting indefinitely.

// Server is running and accepting connections

// Get PID if running

// Get server version if possible

// Get uptime (approximate based on pidfile mtime)

func (s *PgCtlStatusCmd) runStatus(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// No local flag overrides needed - all flags are global now

// Display status for CLI users

// formatUptime formats uptime seconds into human-readable format
func formatUptime(seconds int64) string { _ = "STUB: not implemented"; return "" }

// pgIsReadyTimeoutSecs returns the value to pass to pg_isready's -t flag.
//
// pg_isready -t only accepts whole seconds. The timeout is derived from the
// context deadline so pg_isready's own connection timeout fires before the
// context cancels the subprocess mid-wait. Without this, the subprocess relies
// on libpq's default (3 s), which may race with the gRPC deadline propagated
// down from multiorch (5 s total, shared across two hops).
//
// pgIsReadyDeadlineBuffer is subtracted from the remaining deadline before
// truncating to whole seconds, ensuring the truncation cannot accidentally push
// the timeout above the remaining deadline.
func pgIsReadyTimeoutSecs(ctx context.Context) int { _ = "STUB: not implemented"; return 0 }

func isServerReadyWithConfig(ctx context.Context, config *pgctld.PostgresCtlConfig) bool {
	_ = "STUB: not implemented"
	// Use Unix socket connection for pg_isready
	return false
}

// Need port even for socket connections

func getServerVersionWithConfig(ctx context.Context, config *pgctld.PostgresCtlConfig) string {
	_ = "STUB: not implemented"
	// Use Unix socket connection for psql
	return ""
}

// Need port even for socket connections

func getServerUptime(dataDir string) string { _ = "STUB: not implemented"; return "" }

// Format uptime in human-readable format
