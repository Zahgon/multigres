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

package pgctld

import (
	"time"
)

const (
	postgresConfigWaitRetryTime = 100 * time.Millisecond
	// maxIncludeDepth bounds recursion through postgres include directives.
	maxIncludeDepth = 10

	directiveInclude         = "include"
	directiveIncludeIfExists = "include_if_exists"
	directiveIncludeDir      = "include_dir"
)

// PostgresServerConfig is a memory structure that contains PostgreSQL server configuration parameters.
// It can be used to read standard postgresql.conf files and can also be populated from an existing postgresql.conf.
//
// port, listen_addresses, unix_socket_directories, and data_directory are intentionally
// absent: pgctld pins them via postgres command-line args at start, which take precedence
// over postgresql.conf.
type PostgresServerConfig struct {
	// Core connection settings
	MaxConnections int

	// File locations (template fields)
	DataDir   string // matches {{.DataDir}} in template
	HbaFile   string // matches {{.HbaFile}} in template
	IdentFile string // matches {{.IdentFile}} in template

	// Memory settings
	SharedBuffers      string
	MaintenanceWorkMem string
	WorkMem            string

	// Worker and parallel settings
	MaxWorkerProcesses            int
	EffectiveIoConcurrency        int
	MaxParallelWorkers            int
	MaxParallelWorkersPerGather   int
	MaxParallelMaintenanceWorkers int

	// WAL settings
	WalBuffers string
	MinWalSize string
	MaxWalSize string

	// Checkpoint settings
	CheckpointCompletionTarget float64

	// Replication settings
	MaxWalSenders       int
	MaxReplicationSlots int

	// Query planner settings
	EffectiveCacheSize      string
	RandomPageCost          float64
	DefaultStatisticsTarget int

	// Other important settings
	ClusterName string
	User        string // PostgreSQL user name for HBA configuration

	configMap map[string]string
	Path      string // the actual path that represents this postgresql.conf
}

func (cnf *PostgresServerConfig) lookup(key string) string { _ = "STUB: not implemented"; return "" }

func (cnf *PostgresServerConfig) lookupWithDefault(key, defaultVal string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (cnf *PostgresServerConfig) lookupInt(key string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cnf *PostgresServerConfig) lookupFloat(key string) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// parseConfigInto follows postgres include semantics: relative include paths
// resolve against the including file's directory.
func parseConfigInto(path string, configMap map[string]string, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

// parseConfigLine extracts a key/value from a single postgresql.conf line.
// Returns ("", "") for comments, blanks, and malformed lines.
func parseConfigLine(line string) (string, string) { _ = "STUB: not implemented"; return "", "" }

// parseIncludeDir mirrors postgres' include_dir: parse every *.conf in
// lexicographic order, skipping dotfiles and directories.
func parseIncludeDir(dir string, configMap map[string]string, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

func resolveIncludePath(baseDir, p string) string { _ = "STUB: not implemented"; return "" }

// stripQuotes removes surrounding single or double quotes from a string value
// and removes any trailing comments
func stripQuotes(value string) string { _ = "STUB: not implemented"; return "" }

// Remove trailing comments (anything after # with optional whitespace)

// ReadPostgresServerConfig populates pgConfig from postgresql.conf at
// pgConfig.Path. include / include_if_exists / include_dir directives are
// followed recursively (an included file may itself include further files), up
// to maxIncludeDepth, so the in-memory struct matches what postgres will load
// at runtime.
func ReadPostgresServerConfig(pgConfig *PostgresServerConfig, waitTime time.Duration) (*PostgresServerConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Parse and map configuration values to struct fields

// default

// Memory settings - required in our controlled config

// Worker and parallel settings - required in our controlled config

// WAL settings - required in our controlled config

// Checkpoint settings - required in our controlled config

// Replication settings - required in our controlled config

// Query planner settings - required in our controlled config

// Other important settings
