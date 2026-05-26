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
	"github.com/spf13/cobra"
)

// AddGetDatabaseCommand adds the getdatabase subcommand
func AddGetDatabaseCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Add command-specific flags

// Mark the name flag as required

// runGetDatabase executes the getdatabase command
func runGetDatabase(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	// Get the database name
	return nil
}

// Create admin client

// Create context with timeout and call GetDatabase RPC

// Output the response in JSON format
