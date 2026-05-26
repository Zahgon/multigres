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
	"github.com/multigres/multigres/go/tools/viperutil"

	"github.com/spf13/cobra"
)

// initCmd holds the init command configuration
type initCmd struct {
	provisioner viperutil.Value[string]
	backupPath  viperutil.Value[string]
	backupURL   viperutil.Value[string]
	region      viperutil.Value[string]
}

// getConfigPaths returns the list of config paths.
func getConfigPaths(cmd *cobra.Command) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// buildConfigFromFlags creates a MultigresConfig based on command flags
func (icmd *initCmd) buildConfigFromFlags(cmd *cobra.Command, configPaths []string) (*MultigresConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Parse and validate backup configuration

// Get the provisioner instance

// buildBackupConfig reads backup flags and builds config map
func (icmd *initCmd) buildBackupConfig(configPaths []string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If no backup URL, use local backups

// Parse S3 URL

// Validate region requirement

// Generate timestamped prefix

// Validate credentials and S3 access

//nolint:gocritic // CLI entry point, no parent context available

// createConfigFile creates and writes the multigres configuration file
func createConfigFile(config *MultigresConfig, configPaths []string) (string, error) {
	_ = "STUB: not implemented"
	// Validate the configuration before writing it
	return "", nil
}

// Marshal to YAML

// Determine config file path - use the first config path

// Check if config file already exists

// Check if config directory exists

// Print the generated configuration

// Write config file

// runInit handles the initialization of a multigres cluster configuration
func (icmd *initCmd) runInit(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Build configuration from flags

// Create and write config file

// validateConfig validates the configuration using the appropriate provisioner
func validateConfig(config *MultigresConfig) error {
	_ = "STUB: not implemented"
	// Get the provisioner instance
	return nil
}

// Validate the provisioner-specific configuration

// AddInitCommand adds the init subcommand to the cluster command
func AddInitCommand(clusterCmd *cobra.Command) { _ = "STUB: not implemented"; return }
