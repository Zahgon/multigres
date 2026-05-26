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

// ExpandToAbsolutePath converts a relative path to an absolute path.
// If the path is already absolute, it returns the path unchanged.
func ExpandToAbsolutePath(dir string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// If already absolute, return as-is

// Convert relative path to absolute

// GeneratePostgresServerConfig writes postgresql.conf from the embedded template,
// appends extraConfFiles verbatim (postgres last-write-wins), then reads it back.
func GeneratePostgresServerConfig(poolerDir string, pgUser string, extraConfFiles []string) (*PostgresServerConfig, error) {
	_ = "STUB: not implemented"
	// Create minimal config for template generation
	return nil, nil
}

// Expand relative path to absolute path for consistent path handling

// Set Multigres default values - starting with Pico instance defaults from Supabase
// Reference: https://github.com/supabase/supabase-admin-api/blob/3765a153ef6361cb19a1cbd485cdbf93e0a1820a/optimizations/postgres.go#L38
// These can be changed in the future based on instance size/requirements

// TODO: @rafael - This setting doesn't work for local on macOS environment,
// so it's not matching exactly what we have in Supabase.

// Generate config file from template

// Generate HBA file from template

// Read the generated config back from disk to get all template values

// appendExtraConfFiles concatenates each path onto postgresql.conf. The
// "## <path>" header before each block lets readers attribute lines back to
// their source file.
func (cnf *PostgresServerConfig) appendExtraConfFiles(paths []string) error {
	_ = "STUB: not implemented"
	return nil
}

// generateConfigFile creates the postgresql.conf file using the embedded template
func (cnf *PostgresServerConfig) generateConfigFile() error {
	_ = "STUB: not implemented"
	// Ensure directory exists
	return nil
}

// Generate config content from template

// Write to file

// generateHbaFile creates the pg_hba.conf file using the embedded template
func (cnf *PostgresServerConfig) generateHbaFile() error {
	_ = "STUB: not implemented"
	// Generate HBA content from template
	return nil
}

// Write to file

// PostgresDataDir returns the PostgreSQL data directory from the PGDATA environment variable.
func PostgresDataDir() string { _ = "STUB: not implemented"; return "" }

// PostgresSocketDir returns the default location of the PostgreSQL Unix sockets.
func PostgresSocketDir(poolerDir string) string { _ = "STUB: not implemented"; return "" }

// PostgresConfigFile returns the location of the postgresql.conf file within PGDATA.
func PostgresConfigFile() string { _ = "STUB: not implemented"; return "" }

// MakePostgresConf will substitute values in the template
func (cnf *PostgresServerConfig) MakePostgresConf(templateContent string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// MakeHbaConf will substitute values in the HBA template
func (cnf *PostgresServerConfig) MakeHbaConf(templateContent string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
