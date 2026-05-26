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

package pooler

import (
	"github.com/spf13/cobra"
)

// AddGetPoolerStatusCommand adds the getpoolerstatus subcommand
func AddGetPoolerStatusCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// runGetPoolerStatus executes the getpoolerstatus command
func runGetPoolerStatus(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Create admin client

// Create context with timeout and call GetPoolerStatus RPC

// Output the response in JSON format using protojson to properly render enums as strings

// Use snake_case field names from proto instead of camelCase
