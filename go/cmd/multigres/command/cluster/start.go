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

import (
	"github.com/multigres/multigres/go/provisioner"

	"github.com/spf13/cobra"
)

// ServiceInfo holds information about a provisioned service
type ServiceInfo struct {
	Name    string
	FQDN    string
	Ports   map[string]int
	LogFile string
}

// ServiceSummary holds all provisioned services
type ServiceSummary struct {
	Services []ServiceInfo
}

// AddService adds a service to the summary
func (s *ServiceSummary) AddService(name string, result *provisioner.ProvisionResult) {
	_ = "STUB: not implemented"
	// Extract log file path from metadata if available
	return
}

// PrintSummary prints a formatted summary of all provisioned services
func (s *ServiceSummary) PrintSummary() { _ = "STUB: not implemented"; return }

// Single port format

// Multiple ports format

// Find services with HTTP ports and add direct links

// Find the first multigateway service and show connection command

// Show only the first gateway

// start handles the cluster up command
func start(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

// Get config paths from flags

// Load configuration to determine provisioner type

// Create provisioner instance

// Let provisioner load its own configuration

// Initialize service summary to track all provisioned services

// Use the provisioner's Bootstrap method to provision all services

// Add all returned services to summary dynamically

// Print comprehensive summary

// AddStartCommand adds the start subcommand to the cluster command
func AddStartCommand(clusterCmd *cobra.Command) { _ = "STUB: not implemented"; return }

// No additional flags needed - config-path is provided by viperutil via root command
