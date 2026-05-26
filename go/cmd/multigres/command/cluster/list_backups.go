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
	"github.com/spf13/cobra"

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	multiadminpb "github.com/multigres/multigres/go/pb/multiadmin"
)

// AddListBackupsCommand adds the list-backups subcommand to the cluster command
func AddListBackupsCommand(clusterCmd *cobra.Command) { _ = "STUB: not implemented"; return }

func runListBackups(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

// Create admin client

// Create context with timeout

// Note: TableGroup and Shard are currently hardcoded to defaults because
// the system only supports a single table group ("default") and shard ("0-inf").
// Once multi-shard support is implemented, we should add --table-group and --shard flags.

// Calculate column widths based on data (start with header widths)

// Scan data to find max widths

// Build format string

// Calculate total width for separator

// 12 for spacing (6 gaps × 2 spaces)

// Print header

// Print each backup

func backupStatusToString(status multiadminpb.BackupStatus) string {
	_ = "STUB: not implemented"
	return ""
}

func poolerTypeToString(pt clustermetadatapb.PoolerType) string {
	_ = "STUB: not implemented"
	return ""
}

func formatBytes(bytes uint64) string { _ = "STUB: not implemented"; return "" }
