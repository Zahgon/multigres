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
	"log/slog"

	"github.com/spf13/cobra"
)

// InitResult contains the result of initializing PostgreSQL data directory
type InitResult struct {
	AlreadyInitialized bool
	Message            string
}

// PgCtldInitCmd holds the init command configuration
type PgCtldInitCmd struct {
	pgCtlCmd *PgCtlCommand
}

// AddInitCommand adds the init subcommand to the root command
func AddInitCommand(root *cobra.Command, pc *PgCtlCommand) { _ = "STUB: not implemented"; return }

func (i *PgCtldInitCmd) createCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// InitDataDirWithResult initializes PostgreSQL data directory and returns detailed result information.
// When pgDatabase differs from the default "postgres" database, it starts PostgreSQL transiently
// and creates the target database — mirroring docker-library/postgres's docker_setup_db behaviour.
func InitDataDirWithResult(logger *slog.Logger, poolerDir string, cfg PgCtldServiceConfig) (*InitResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if data directory is already initialized

// create server config using the pooler directory

// Post-initdb steps that need a running server (custom DB creation, init
// SQL files) share a single transient PostgreSQL instance.

// postInitdbSetup starts a transient PostgreSQL instance and performs any
// post-initdb setup steps: creating a custom target database (if requested)
// and running user-provided init SQL against the target database.
//
// Execution order:
//  1. Create target database (if non-default).
//  2. Run --pg-initdb-sql-dirs: bulk schema/migration directories, each under
//     SET SESSION AUTHORIZATION <role>. Directories establish the base schema.
//  3. Run --pg-initdb-sql-files: individual files applied on top as targeted
//     overrides or patches.
func postInitdbSetup(logger *slog.Logger, cfg PgCtldServiceConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// Dirs first: establish bulk schema/migrations under the specified roles.

// Files second: apply targeted overrides or patches on top of the base schema.

// runInitdbSQLFiles executes each SQL file against database with
// ON_ERROR_STOP=1 so a failing statement aborts its script.
func runInitdbSQLFiles(logger *slog.Logger, pg *pgInstance, database string, files []string) error {
	_ = "STUB: not implemented"
	return nil
}

// runInitdbSQLDirs processes each role:path entry: reads all .sql files from the
// directory in lexicographic order and runs them in a single psql session under
// SET SESSION AUTHORIZATION "<role>" / RESET SESSION AUTHORIZATION.
func runInitdbSQLDirs(logger *slog.Logger, pg *pgInstance, database string, entries []string) error {
	_ = "STUB: not implemented"
	return nil
}

// parseSQLDirEntry splits a "role:path" entry on the first colon.
func parseSQLDirEntry(entry string) (role, dir string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// sqlFilesInDir returns .sql file paths from dir in lexicographic order,
// skipping subdirectories and non-.sql files.
// os.ReadDir already returns entries sorted by name.
func sqlFilesInDir(dir string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (i *PgCtldInitCmd) runInit(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Display appropriate message for CLI users

// createDatabaseOnInstance creates database on an already-running transient
// pgInstance if it does not already exist. This mirrors what the official
// docker-library/postgres image does in its docker_setup_db() entrypoint function.
func createDatabaseOnInstance(logger *slog.Logger, pg *pgInstance, database string) error {
	_ = "STUB: not implemented"
	// Check whether the target database already exists.
	// Use Go string formatting to build the SQL — the database name comes from
	// operator config, not untrusted user input, so simple quoting is safe.
	// Single quotes in the name are escaped as '' per the SQL standard.
	return nil
}

// resolveInitdbPwFile returns a filesystem path suitable for initdb's --pwfile
// flag along with a cleanup function the caller must defer.
//
// The primary path is the operator-supplied password file (POSTGRES_PASSWORD_FILE
// / --pg-password-file). In that case the returned path *is* that file
// and no plaintext is ever copied — initdb reads it directly. The cleanup is a
// no-op.
//
// The fallback path is the legacy POSTGRES_PASSWORD environment variable, which
// is being phased out. There is no file on disk yet, so we have to stage the
// password into a randomly named temp file long enough for initdb to read it.
// The cleanup unlinks that temp file. This branch exists only to keep operators
// who have not yet migrated to file-based secrets working; new deployments
// should always hit the primary path.
//
// Note on first-line semantics: initdb --pwfile reads only the first line of
// the file. pgsecret.ReadPasswordFile returns the whole content (minus
// trailing CR/LF). The two parsers agree for any single-line password (with
// or without a trailing newline) — which is every realistic Kubernetes Secret.
// A multi-line file would silently diverge: the cluster would be initialized
// with line 1, but the admin pool / replication would try to authenticate with
// the whole content. Password files MUST be single-line.
func resolveInitdbPwFile(cfg PgCtldServiceConfig) (path string, cleanup func(), err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Fallback: POSTGRES_PASSWORD env. Stage the plaintext to a randomly named
// file in /tmp. Removed unconditionally on return so the plaintext window
// is bounded by the initdb exec. The filename has no prefix so its purpose
// isn't advertised in /tmp listings.

func initializeDataDir(logger *slog.Logger, cfg PgCtldServiceConfig) error {
	_ = "STUB: not implemented"
	// Derive dataDir from poolerDir using the standard convention
	return nil
}

// Note: initdb will create the data directory itself if it doesn't exist.
// We don't create it beforehand to avoid leaving empty directories if initdb fails.

// Build initdb command
// It's generally a good idea to enable page data checksums. Furthermore,
// pgBackRest will validate checksums for the Postgres cluster it's backing up.
// However, pgBackRest merely logs checksum validation errors but does not fail
// the backup.

// Invariant: production callers (runInit, the InitDataDir gRPC handler)
// populate Password, PasswordSource (and PasswordFile when source==File)
// via PgCtlCommand.GetPostgresPassword, which errors when no source is
// configured. Reaching here with a missing source means a caller — almost
// certainly a test — built PgCtldServiceConfig by hand without going
// through the resolver. Panic rather than guess a source that would
// mislabel the log line below.

// Capture both stdout and stderr to include in error messages

// quoteIdentifier wraps name in double quotes and escapes any embedded double
// quotes as "" per the SQL standard, producing a safe PostgreSQL identifier.
// The value comes from operator config, not from untrusted user input, so a
// simple ReplaceAll is sufficient here. If that assumption ever changes, replace
// this with pq.QuoteIdentifier from github.com/lib/pq, which applies the same
// escaping but is a well-tested, purpose-built function.
func quoteIdentifier(name string) string { _ = "STUB: not implemented"; return "" }
