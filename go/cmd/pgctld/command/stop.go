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

	"github.com/multigres/multigres/go/services/pgctld"
	"github.com/multigres/multigres/go/tools/viperutil"

	"github.com/spf13/cobra"
)

// StopResult contains the result of stopping PostgreSQL
type StopResult struct {
	WasRunning bool
	Message    string
}

// PgCtlStopCmd holds the stop command configuration
type PgCtlStopCmd struct {
	pgCtlCmd *PgCtlCommand
	mode     viperutil.Value[string]
}

// AddStopCommand adds the stop subcommand to the root command
func AddStopCommand(root *cobra.Command, pc *PgCtlCommand) { _ = "STUB: not implemented"; return }

func (s *PgCtlStopCmd) createCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (s *PgCtlStopCmd) runStop(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Stop PostgreSQL and get detailed result information. We are using a
// password here because the CHECKPOINT command requires authentication,
// even if the stop command itself does not.
//
// TODO: Consider removing the CHECKPOINT command in the stop flow since
// it's not strictly necessary and adds complexity (requires password, can
// fail if PostgreSQL is already in a bad state, etc.)

// Display appropriate message for CLI users

// StopPostgreSQLWithResult stops PostgreSQL with the given configuration and returns detailed result information
func StopPostgreSQLWithResult(logger *slog.Logger, config *pgctld.PostgresCtlConfig, mode string) (*StopResult, error) {
	_ = "STUB: not implemented"
	return nil,

		// Default mode to "fast" if not specified
		nil
}

// Check if PostgreSQL is running

// StopPostgreSQLWithConfig stops PostgreSQL with the given configuration and mode
func StopPostgreSQLWithConfig(logger *slog.Logger, config *pgctld.PostgresCtlConfig, mode string) error {
	_ = "STUB: not implemented"
	return nil
}

// For backward compatibility, log the message if PostgreSQL was actually stopped

func stopPostgreSQLWithConfig(logger *slog.Logger, config *pgctld.PostgresCtlConfig, mode string) error {
	_ = "STUB: not implemented"
	// First try using pg_ctl
	return nil
}

func stopWithPgCtlWithConfig(logger *slog.Logger, config *pgctld.PostgresCtlConfig, mode string) error {
	_ = "STUB: not implemented"
	// Take a checkpoint before stopping PostgreSQL for clean shutdown
	return nil
}

// Continue with stop even if checkpoint fails - it's not critical

// takeCheckpoint executes a CHECKPOINT command to ensure all data is written to disk before shutdown
func takeCheckpoint(logger *slog.Logger, config *pgctld.PostgresCtlConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// Use Unix socket connection for psql

// Need port even for socket connections

// quiet mode - suppress messages

// Capture output to avoid cluttering the terminal
