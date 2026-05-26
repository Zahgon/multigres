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
	"context"
	"log/slog"
)

type PgRewindResult struct {
	// Status message
	Message string
	Output  string
}

func PgRewindWithResult(ctx context.Context, logger *slog.Logger, sourceServer, password string, dryRun bool, extraArgs []string) (*PgRewindResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set PGPASSWORD environment variable for pg_rewind to use
// pg_rewind doesn't reliably use passwords from connection strings

// Capture both Stdout and Stderr
