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
	"context"

	"github.com/spf13/cobra"
)

// teardownAllServices stops all provisioned services using the provisioner's Teardown method
func teardownAllServices(ctx context.Context, provisionerName string, configPaths []string, clean bool) error {
	_ = "STUB: not implemented"
	// Create provisioner instance
	return nil
}

// Let provisioner load its own configuration

// Use the provisioner's teardown method

// down handles the cluster down command
func down(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

// Get the clean flag

// Get config paths from flags

// Load configuration to determine provisioner type

// Teardown all services using the provisioner

// AddStopCommand adds the stop subcommand to the cluster command
func AddStopCommand(clusterCmd *cobra.Command) { _ = "STUB: not implemented"; return }
