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

package topo

import (
	"context"

	"github.com/multigres/multigres/go/common/topoclient"
	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"

	"github.com/spf13/cobra"
)

// CreateClusterMetadataCommand adds the createclustermetadataCommand subcommand
func CreateClusterMetadataCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Add command-specific flags

// createClusterMetadata executes the createclustermetadata command
func createClusterMetadata(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Create topology store using configured backend

// Create each cell

// Provision the database

// createCell creates a single cell in the topology
func createCell(ctx context.Context, ts topoclient.Store, cellName, etcdAddress, globalTopoRoot string) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if cell already exists

// Create the cell if it doesn't exist

// Construct cell root path based on global root and cell name

// Some other error occurred

// provisionDatabase registers a database in the global topology
func provisionDatabase(ctx context.Context, ts topoclient.Store, databaseName string, cellNames []string, backupLocation string, bootstrapPolicy *clustermetadatapb.DurabilityPolicy) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if database already exists

// Create the database if it doesn't exist

// Some other error occurred
