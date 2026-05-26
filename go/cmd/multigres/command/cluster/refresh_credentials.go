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

package cluster

import (
	"github.com/spf13/cobra"
)

// refreshCredentialsCmd holds the refresh-credentials command configuration
type refreshCredentialsCmd struct{}

// AddRefreshCredentialsCommand adds the refresh-credentials subcommand
func AddRefreshCredentialsCommand(clusterCmd *cobra.Command) { _ = "STUB: not implemented"; return }

func (rcmd *refreshCredentialsCmd) runRefreshCredentials(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Get config paths

// Load cluster config

// Get backup config from provisioner-config

// Get S3 config

// Check if using env credentials (can be string or bool)

// Read AWS credentials from environment

// Create backup config to get credentials

// Get new credentials

// Find all pooler directories

// Update pgbackrest.conf in each pooler directory

// Read current config

// Update credentials in config

// Write back atomically

// maskCredential masks all but the first 4 and last 4 characters of a credential
func maskCredential(cred string) string { _ = "STUB: not implemented"; return "" }
