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

package admin

import (
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	multiadminpb "github.com/multigres/multigres/go/pb/multiadmin"
)

// Conn wraps a gRPC connection to the multiadmin server.
type Conn struct {
	multiadminpb.MultiAdminServiceClient
	conn *grpc.ClientConn
}

// Close closes the underlying gRPC connection.
func (c *Conn) Close() error { _ = "STUB: not implemented"; return nil }

// NewClient creates a connection to the multiadmin server.
func NewClient(cmd *cobra.Command) (*Conn, error) { _ = "STUB: not implemented"; return nil, nil }

// GetServerAddress resolves the admin server address from flags or config
func GetServerAddress(cmd *cobra.Command) (string, error) {
	_ = "STUB: not implemented"
	// Check if admin-server flag is provided
	return "", nil
}

// Fall back to config file

// Load config and extract multiadmin address

// getServerFromConfig extracts the multiadmin server address from config
func getServerFromConfig(configPaths []string) (string, error) {
	_ = "STUB: not implemented"
	// Find the config file
	return "", nil
}

// Read the config file directly

// Parse the config structure

// Extract multiadmin config for local provisioner

// Convert the map to YAML and then to typed config

// Build admin server address from config
