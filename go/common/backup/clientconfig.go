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

package backup

// ClientConfigOpts holds options for generating pgbackrest.conf for client operations.
type ClientConfigOpts struct {
	PoolerDir     string // Base directory for pooler data
	Pg1Port       int    // Local PostgreSQL port
	Pg1SocketPath string // Local PostgreSQL socket directory
	Pg1Path       string // Local PostgreSQL data directory
	Pg1User       string // PostgreSQL superuser for pgbackrest connections
}

// WriteClientConfig generates pgbackrest.conf for client operations (backup, restore, info).
// Returns the path to the generated config file.
func WriteClientConfig(opts ClientConfigOpts, backupCfg *Config) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Create directories

// Parse the template

// Generate repo configuration

// Get credentials if using environment credentials

// Prepare template data

// Execute template

// Write config with appropriate permissions
