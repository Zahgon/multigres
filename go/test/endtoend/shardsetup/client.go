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

package shardsetup

import (
	"context"
	"testing"

	"google.golang.org/grpc"

	consensuspb "github.com/multigres/multigres/go/pb/consensus"
	multiorchpb "github.com/multigres/multigres/go/pb/multiorch"
	multipoolermanagerpb "github.com/multigres/multigres/go/pb/multipoolermanager"
	pgctldpb "github.com/multigres/multigres/go/pb/pgctldservice"
)

// MultipoolerClient wraps a gRPC connection to a multipooler and provides access to
// manager, consensus, and pooler service clients over the same connection.
// Follows the pattern from multipooler/setup_test.go:multipoolerClient.
type MultipoolerClient struct {
	conn      *grpc.ClientConn
	Manager   multipoolermanagerpb.MultiPoolerManagerClient
	Consensus consensuspb.MultiPoolerConsensusClient
	Pooler    *MultiPoolerTestClient
}

// NewMultipoolerClient creates a new MultipoolerClient connected to the given gRPC port.
// It establishes a single gRPC connection and creates clients for all multipooler services.
// Follows the pattern from multipooler/setup_test.go:newMultipoolerClient.
func NewMultipoolerClient(grpcPort int) (*MultipoolerClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create pooler test client (uses its own connection internally)

// Close closes all underlying connections.
func (c *MultipoolerClient) Close() error { _ = "STUB: not implemented"; return nil }

// WaitForManagerReady waits for the manager to be in ready state.
// Follows the pattern from multipooler/setup_test.go:waitForManagerReady.
func WaitForManagerReady(t *testing.T, manager *ProcessInstance) {
	_ = "STUB: not implemented"

	// Connect to the manager
	return
}

// Use require.Eventually to wait for manager to be ready

// QueryStringValue executes a query and extracts the first column of the first row as a string.
// Returns empty string and error if query fails or returns no rows.
// Follows the pattern from multipooler/setup_test.go:queryStringValue.
func QueryStringValue(ctx context.Context, client *MultiPoolerTestClient, query string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// PgctldClient wraps the pgctld gRPC client.
type PgctldClient struct {
	conn *grpc.ClientConn
	pgctldpb.PgCtldClient
}

// NewPgctldClient creates a new PgctldClient connected to the given gRPC port.
func NewPgctldClient(grpcPort int) (*PgctldClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close closes the underlying connection.
func (c *PgctldClient) Close() error { _ = "STUB: not implemented"; return nil }

// MultiOrchClient wraps the multiorch gRPC client.
type MultiOrchClient struct {
	conn *grpc.ClientConn
	multiorchpb.MultiOrchServiceClient
}

// NewMultiOrchClient creates a new MultiOrchClient connected to the given gRPC port.
func NewMultiOrchClient(grpcPort int) (*MultiOrchClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close closes the underlying connection.
func (c *MultiOrchClient) Close() error { _ = "STUB: not implemented"; return nil }
