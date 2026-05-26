// Copyright 2025 Supabase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package command

import (
	"log/slog"

	"github.com/spf13/cobra"

	"github.com/multigres/multigres/go/services/pgctld"
)

// StartResult contains the result of starting PostgreSQL
type StartResult struct {
	PID            int
	AlreadyRunning bool
	Message        string
}

// NewPostgresCtlConfigFromDefaults creates a PostgresCtlConfig using
// command-line parameters. Port, listen_addresses, and
// unix_socket_directories come from CLI flags, not from the config file.
//
// Password is intentionally left unset — callers that need it (start, stop)
// resolve it via PgCtlCommand.GetPostgresPassword so the file/env precedence
// stays in one place.
func NewPostgresCtlConfigFromDefaults(poolerDir string, pgPort int, pgListenAddresses string, pgUser string, pgDatabase string, timeout int) (*pgctld.PostgresCtlConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddStartCommand adds the start subcommand to the root command
func AddStartCommand(root *cobra.Command, pc *PgCtlCommand) { _ = "STUB: not implemented"; return }

// PgCtlStartCmd holds the start command configuration
type PgCtlStartCmd struct {
	pgCtlCmd *PgCtlCommand
}

func (s *PgCtlStartCmd) createCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (s *PgCtlStartCmd) runStart(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Display appropriate message for CLI users

// StartPostgreSQLWithResult starts PostgreSQL with the given configuration and returns detailed result information
func StartPostgreSQLWithResult(logger *slog.Logger, config *pgctld.PostgresCtlConfig) (*StartResult, error) {
	_ = "STUB: not implemented"
	return nil,

		// Check if PostgreSQL is already running
		nil
}

// Get PID of running instance

// Ensure Unix socket directory exists before starting PostgreSQL
// This is necessary for restarts after restores, where pgBackRest only restores pg_data
// but not external directories like pg_sockets

// Enforce PGDATA permission invariant before pg_ctl start

// Start PostgreSQL

// Wait for server to be ready

// Get PID of started instance

// StartPostgreSQLWithConfig starts PostgreSQL with the given configuration
func StartPostgreSQLWithConfig(logger *slog.Logger, config *pgctld.PostgresCtlConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// For backward compatibility, log the message if provided

// ensurePGDATAPermissions ensures PGDATA is owned by the effective UID and set to 0700 before pg_ctl start.
// initdb sets this on bootstrap, but restore, rewind, or volume remounts may change it.
func ensurePGDATAPermissions(logger *slog.Logger, dataDir string) error {
	_ = "STUB: not implemented"
	return nil
}

func isPostgreSQLRunning(dataDir string) bool {
	_ = "STUB: not implemented"
	// Check if postmaster.pid file exists and process is running
	return false
}

// Read PID from file and check if process is actually running

func startPostgreSQLWithConfig(logger *slog.Logger, config *pgctld.PostgresCtlConfig) error {
	_ = "STUB: not implemented"
	// Use pg_ctl to start PostgreSQL properly as a daemon
	// Pass port, listen_addresses, and unix_socket_directories as command-line parameters for portability
	return nil
}

// don't wait - we'll check readiness ourselves

// If orphan detection environment variables are set, spawn a watchdog process
// that will stop postgres if the test parent dies or testdata dir is deleted

// Put watchdog in its own process group so SIGINT/SIGTERM to parent doesn't kill it
// The watchdog needs to survive the parent's death to perform cleanup

// Environment variables automatically inherit

// Don't fail the start operation if watchdog fails to start

// readLogTail reads the last N lines from the PostgreSQL log file for diagnostics
func readLogTail(logPath string, lines int) string { _ = "STUB: not implemented"; return "" }

func waitForPostgreSQLWithConfig(logger *slog.Logger, config *pgctld.PostgresCtlConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// On timeout, include diagnostic information

// Check if PostgreSQL process is still running (after first second)

// No PID file means PostgreSQL never started or crashed immediately

// PID file exists but process is gone - crashed

// Log progress every 5 seconds

func readPostmasterPID(dataDir string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// First line contains the PID

func isProcessRunning(pid int) bool { _ = "STUB: not implemented"; return false }

// On Unix, sending signal 0 checks if process exists without actually sending a signal.
// This is the standard way to check if a process is running.
