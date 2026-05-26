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

	"github.com/multigres/multigres/go/tools/viperutil"
)

// checkBackupConfigCmd holds the check-backup-config command configuration
type checkBackupConfigCmd struct {
	backupURL viperutil.Value[string]
	region    viperutil.Value[string]
}

// AddCheckBackupConfigCommand adds the check-backup-config subcommand
func AddCheckBackupConfigCommand(clusterCmd *cobra.Command) { _ = "STUB: not implemented"; return }

func (ccmd *checkBackupConfigCmd) runCheckBackupConfig(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if --backup-url flag was explicitly provided on command line

// Use flags - parse backup URL

// Validate region requirement

// Load from config file

// Get backup config from provisioner-config

// Get S3 config

// Extract S3 config

// Check if using env credentials (can be string or bool)

// Read AWS credentials from environment

// Check credentials

// Validate S3 access

// Provide helpful error messages based on error type
