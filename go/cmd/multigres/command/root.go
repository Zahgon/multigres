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

package command

import (
	"github.com/multigres/multigres/go/tools/telemetry"
	"github.com/multigres/multigres/go/tools/viperutil"

	"github.com/spf13/cobra"
)

// MultigresCommand holds the configuration for multigres commands
type MultigresCommand struct {
	reg       *viperutil.Registry
	vc        *viperutil.ViperConfig
	telemetry *telemetry.Telemetry
}

// GetRootCommand creates and returns the root command for multigres with all subcommands
func GetRootCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Silence usage for application errors, but allow it for flag errors
// This gets called after flag parsing, so flag errors will still show usage

// Set multigres-specific config name

// Load config (without the full servenv setup)

/* startSpan */

// Shutdown OpenTelemetry to flush all pending spans
// This is critical for CLI commands to export traces before process exit

// Add any other servenv flags

// Override the default display value for multigres

// Configure output streams explicitly
// Otherwise cobra will output commands to StdErr

// Add all subcommands
