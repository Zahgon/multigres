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
	"time"

	"github.com/spf13/cobra"

	multiadminpb "github.com/multigres/multigres/go/pb/multiadmin"
	"github.com/multigres/multigres/go/tools/viperutil"
)

const (
	backupPollInterval = 2 * time.Second
	backupPollTimeout  = 30 * time.Minute
)

// backupCmd holds the backup command configuration
type backupCmd struct {
	database   viperutil.Value[string]
	backupType viperutil.Value[string]
	primary    viperutil.Value[bool]
	timeout    viperutil.Value[time.Duration]
}

// AddBackupCommand adds the backup subcommand to the cluster command
func AddBackupCommand(clusterCmd *cobra.Command) {
	_ = "STUB: not implemented"
	// Create a viperutil registry for backup command flags
	return
}

func (bcmd *backupCmd) runBackup(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Create admin client

// Start backup

// Create context with timeout for the initial Backup call

// Poll for completion

func pollBackupJobStatus(cmd *cobra.Command, client multiadminpb.MultiAdminServiceClient, jobID string) error {
	_ = "STUB: not implemented"
	return nil
}

// Create a fresh timeout context for each GetBackupJobStatus call
