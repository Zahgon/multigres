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

	"github.com/multigres/multigres/go/common/servenv"
	"github.com/multigres/multigres/go/tools/telemetry"
	"github.com/multigres/multigres/go/tools/viperutil"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// PgCtlCommand holds the configuration for pgctld commands.
// This contains all flags and information necessary to run any pgctld command.
type PgCtlCommand struct {
	reg            *viperutil.Registry
	pgDatabase     viperutil.Value[string]
	pgUser         viperutil.Value[string]
	pgPassword     viperutil.Value[string]
	pgPasswordFile viperutil.Value[string]
	// flagSet is the persistent flag set, saved during GetRootCommand so
	// GetPostgresPassword can use pflag.Flag.Changed to distinguish "flag
	// explicitly set" from "flag at default value".
	flagSet            *pflag.FlagSet
	poolerDir          viperutil.Value[string]
	timeout            viperutil.Value[int]
	pgPort             viperutil.Value[int]
	pgListenAddresses  viperutil.Value[string]
	pgHbaTemplate      viperutil.Value[string]
	postgresConfigTmpl viperutil.Value[string]
	pgInitdbArgs       viperutil.Value[string]
	pgInitdbSQLFiles   viperutil.Value[[]string]
	pgInitdbSQLDirs    viperutil.Value[[]string]
	pgInitdbExtraConf  viperutil.Value[[]string]

	vc        *viperutil.ViperConfig
	lg        *servenv.Logger
	telemetry *telemetry.Telemetry
}

// GetRootCommand creates and returns the root command for pgctld with all subcommands
func GetRootCommand() (*cobra.Command, *PgCtlCommand) { _ = "STUB: not implemented"; return nil, nil }

// No FlagName — env var only, no CLI flag

// Flags parsed successfully at this point — suppress usage for any subsequent
// runtime errors so the error message is not buried under the usage text.

// Initialize telemetry for CLI commands (server command will re-initialize via ServEnv.Init)

/* startSpan */

// Shutdown OpenTelemetry to flush all pending spans
// For server command, this runs after the server has shut down

// Backwards-compat alias: --init-db-sql-file → --pg-initdb-sql-files.
// Remove once downstream users have migrated.

// Save the persistent flag set so GetPostgresPassword can use
// pflag.Flag.Changed to distinguish "flag explicitly set" from
// "flag at default value".

// Add all subcommands

// validateGlobalFlags validates required global flags for all pgctld commands
func (pc *PgCtlCommand) validateGlobalFlags(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	// Validate pooler-dir is required and non-empty for all commands
	return nil
}

// If pg-hba-template is specified, read and replace the default template

// If postgres-config-template is specified, read and replace the default template

// GetLogger returns the configured logger instance
func (pc *PgCtlCommand) GetLogger() *slog.Logger { _ = "STUB: not implemented"; return nil }

// PasswordSource describes where GetPostgresPassword resolved the password
// from. Used in log lines so tests and operators can confirm which input was
// consumed without exposing the password itself.
type PasswordSource string

const (
	PasswordSourceNone PasswordSource = "none"
	PasswordSourceFile PasswordSource = "POSTGRES_PASSWORD_FILE" //nolint:gosec // env var name, not a credential
	PasswordSourceEnv  PasswordSource = "POSTGRES_PASSWORD"      //nolint:gosec // env var name, not a credential
)

// buildServiceConfig assembles a PgCtldServiceConfig from this command's
// resolved flags / env vars. Shared between the `server` and `init`
// subcommands so they construct identical configs and stay in sync as
// fields are added. Returns an error from GetPostgresPassword unchanged so
// callers can surface a CLI error.
func (pc *PgCtlCommand) buildServiceConfig() (PgCtldServiceConfig, error) {
	_ = "STUB: not implemented"
	return *new(PgCtldServiceConfig), nil
}

// GetPostgresPassword resolves the postgres superuser password and reports
// which source it came from. Sources are tried in order:
//
//  1. --pg-password-file flag / POSTGRES_PASSWORD_FILE env, if the file exists
//     and has non-empty content. Returns PasswordSourceFile and the file path
//     so downstream code (initdb --pwfile) can read it directly.
//  2. POSTGRES_PASSWORD env, if set to a non-empty value. Returns
//     PasswordSourceEnv.
//
// Two independent inputs are considered, in strict precedence order — once
// a higher-precedence input is "explicitly set" it is authoritative and
// lower-precedence inputs are NOT consulted, even when the higher-precedence
// value turns out to be empty (those empty cases become errors instead of
// fallthroughs):
//
//  1. File path: --pg-password-file flag (Changed), or POSTGRES_PASSWORD_FILE
//     env (LookupEnv).
//     - Path explicitly empty                  → error.
//     - Path set, file content empty           → error.
//     - Path set, file content non-empty       → use it, source=File.
//  2. Env var: POSTGRES_PASSWORD (LookupEnv).
//     - Env set to empty                       → error.
//     - Env set to non-empty                   → use it, source=Env.
//
// Reached the end with no input explicitly set: error "not configured".
// pgctld does not expose a --pg-password flag, so the "option" row from the
// multipooler resolver does not apply here. "Explicitly set" uses
// os.LookupEnv and pflag.Flag.Changed to detect operator intent — viperutil
// collapses unset and empty into the same "" via os.Getenv.
func (pc *PgCtlCommand) GetPostgresPassword() (password string, source PasswordSource, file string, err error) {
	_ = "STUB: not implemented"
	// File path: flag or env.
	return "", *new(PasswordSource), "", nil
}

// Env var: POSTGRES_PASSWORD.

// passwordFileExplicit reports whether the file-path input was explicitly
// set (via --pg-password-file or POSTGRES_PASSWORD_FILE) and whether the
// resulting path is the empty string. It does NOT consider viperutil
// defaults — pflag.Flag.Changed and os.LookupEnv are the source of truth.
func (pc *PgCtlCommand) passwordFileExplicit() (path string, explicit, isEmpty bool) {
	_ = "STUB: not implemented"
	return "", false, false
}

// GetPoolerDir returns the configured pooler directory as an absolute path
func (pc *PgCtlCommand) GetPoolerDir() string { _ = "STUB: not implemented"; return "" }

// If we can't expand the path, return the original to avoid breaking existing behavior
// This should rarely happen in practice

// validateInitialized validates that the PostgreSQL data directory has been initialized
// This should be called by all commands except 'init'
func (pc *PgCtlCommand) validateInitialized(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	// First run the standard global validation
	return nil
}

// Check if data directory is initialized
