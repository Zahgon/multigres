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

package cluster

// MultigresConfig represents the structure of the multigres configuration file
type MultigresConfig struct {
	Provisioner       string         `yaml:"provisioner"`
	ProvisionerConfig map[string]any `yaml:"provisioner-config,omitempty"`
}

// LoadConfig loads the multigres configuration from the specified paths
func LoadConfig(configPaths []string) (*MultigresConfig, string, error) {
	_ = "STUB: not implemented"
	// Try to find the config file in the provided paths
	return nil, "", nil
}

// Validate that provisioner is specified
