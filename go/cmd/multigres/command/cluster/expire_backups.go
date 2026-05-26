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
	"time"

	"github.com/spf13/cobra"

	"github.com/multigres/multigres/go/tools/viperutil"
)

type expireBackupsCmd struct {
	database viperutil.Value[string]
	timeout  viperutil.Value[time.Duration]
}

// AddExpireBackupsCommand adds the expire-backups subcommand to the cluster command
func AddExpireBackupsCommand(clusterCmd *cobra.Command) { _ = "STUB: not implemented"; return }

func (ecmd *expireBackupsCmd) runExpireBackups(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}
