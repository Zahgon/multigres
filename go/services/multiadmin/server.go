// Copyright 2025 Supabase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package server implements the MultiAdmin gRPC service for multigres cluster administration.
// It provides administrative operations for managing and querying cluster components including
// cells, databases, gateways, poolers, and orchestrators through a unified gRPC interface.
package multiadmin

import (
	"context"
	"log/slog"

	"github.com/multigres/multigres/go/common/rpcclient"
	"github.com/multigres/multigres/go/common/topoclient"
	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	multiadminpb "github.com/multigres/multigres/go/pb/multiadmin"

	"google.golang.org/grpc"
)

// MultiAdminServer implements the MultiAdminService gRPC interface
type MultiAdminServer struct {
	multiadminpb.UnimplementedMultiAdminServiceServer

	// ts is the topology store for querying cluster metadata
	ts topoclient.Store

	// logger for structured logging
	logger *slog.Logger

	// backupJobTracker manages async backup/restore jobs
	backupJobTracker *BackupJobTracker

	// rpcClient is the client for communicating with multipooler nodes
	rpcClient rpcclient.MultiPoolerClient

	// gatewayDialer opens a one-shot gRPC connection to a multigateway by
	// host:port for ad-hoc admin RPCs (registry / consolidator snapshots).
	// Defaults to dialing with the configured transport credentials; tests
	// can swap it for a fake.
	gatewayDialer func(ctx context.Context, target string) (*grpc.ClientConn, error)
}

// NewMultiAdminServer creates a new MultiAdminServer instance.
// The transportCreds dial option configures TLS for connections to multipooler nodes.
func NewMultiAdminServer(ts topoclient.Store, logger *slog.Logger, transportCreds grpc.DialOption) *MultiAdminServer {
	_ = "STUB: not implemented"
	return nil
}

// RegisterWithGRPCServer registers the MultiAdmin service with the provided gRPC server
func (s *MultiAdminServer) RegisterWithGRPCServer(grpcServer *grpc.Server) {
	_ = "STUB: not implemented"
	return
}

// Stop stops background goroutines and releases resources
func (s *MultiAdminServer) Stop() { _ = "STUB: not implemented"; return }

// SetRPCClient sets the RPC client for communicating with multipoolers.
// This is primarily used for testing to inject a fake client.
func (s *MultiAdminServer) SetRPCClient(client rpcclient.MultiPoolerClient) {
	_ = "STUB: not implemented"
	return

	// GetCell retrieves information about a specific cell
}

func (s *MultiAdminServer) GetCell(ctx context.Context, req *multiadminpb.GetCellRequest) (*multiadminpb.GetCellResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate request

// Get cell from topology

// Check if it's a not found error

// Return the response

// GetDatabase retrieves information about a specific database
func (s *MultiAdminServer) GetDatabase(ctx context.Context, req *multiadminpb.GetDatabaseRequest) (*multiadminpb.GetDatabaseResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate request

// Get database from topology

// Check if it's a not found error

// Return the response

// GetCellNames retrieves all cell names in the cluster
func (s *MultiAdminServer) GetCellNames(ctx context.Context, req *multiadminpb.GetCellNamesRequest) (*multiadminpb.GetCellNamesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetDatabaseNames retrieves all database names in the cluster
func (s *MultiAdminServer) GetDatabaseNames(ctx context.Context, req *multiadminpb.GetDatabaseNamesRequest) (*multiadminpb.GetDatabaseNamesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetGateways retrieves gateways filtered by cells
func (s *MultiAdminServer) GetGateways(ctx context.Context, req *multiadminpb.GetGatewaysRequest) (*multiadminpb.GetGatewaysResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Determine which cells to query

// If no cells specified, get all cells

// Query each cell for gateways

// Convert to protobuf

// Return partial results with error if some cells failed

// GetPoolers retrieves poolers filtered by cells and/or database
func (s *MultiAdminServer) GetPoolers(ctx context.Context, req *multiadminpb.GetPoolersRequest) (*multiadminpb.GetPoolersResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Determine which cells to query

// If no cells specified, get all cells

// Query each cell for poolers

// filter by database and shard if specified

// Convert to protobuf

// Return partial results with error if some cells failed

// GetOrchs retrieves orchestrators filtered by cells
func (s *MultiAdminServer) GetOrchs(ctx context.Context, req *multiadminpb.GetOrchsRequest) (*multiadminpb.GetOrchsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Determine which cells to query

// If no cells specified, get all cells

// Query each cell for orchestrators

// Convert to protobuf

// Return partial results with error if some cells failed

// GetPoolerStatus retrieves the unified status of a specific pooler by proxying
// the request to the target pooler's MultiPoolerManager.Status RPC.
func (s *MultiAdminServer) GetPoolerStatus(ctx context.Context, req *multiadminpb.GetPoolerStatusRequest) (*multiadminpb.GetPoolerStatusResponse, error) {
	_ = "STUB: not implemented"
	// Validate request
	return nil, nil
}

// Create a fully-qualified pooler ID for topology lookup

// Get pooler from topology

// Call Status RPC on the pooler

// SetPostgresRestartsEnabled enables or disables automatic PostgreSQL restarts on a specific
// pooler by proxying the request to the target pooler's MultiPoolerManager.SetPostgresRestartsEnabled RPC.
func (s *MultiAdminServer) SetPostgresRestartsEnabled(ctx context.Context, req *multiadminpb.SetPostgresRestartsEnabledRequest) (*multiadminpb.SetPostgresRestartsEnabledResponse, error) {
	_ = "STUB: not implemented"
	// Validate request
	return nil, nil
}

// Create a fully-qualified pooler ID for topology lookup

// Get pooler from topology

// GetGatewayQueries proxies a per-fingerprint query registry snapshot from the
// target multigateway's MultiGatewayManager.GetQueryRegistry RPC.
func (s *MultiAdminServer) GetGatewayQueries(ctx context.Context, req *multiadminpb.GetGatewayQueriesRequest) (*multiadminpb.GetGatewayQueriesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetGatewayConsolidator proxies a prepared-statement consolidator snapshot
// from the target multigateway's MultiGatewayManager.GetConsolidatorStats RPC.
func (s *MultiAdminServer) GetGatewayConsolidator(ctx context.Context, req *multiadminpb.GetGatewayConsolidatorRequest) (*multiadminpb.GetGatewayConsolidatorResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// validateGatewayID returns a gRPC error if the ID is missing required fields.
func validateGatewayID(id *clustermetadatapb.ID) error { _ = "STUB: not implemented"; return nil }

// dialGatewayByID resolves a gateway in topology and returns an open gRPC
// connection to it. The caller is responsible for closing the connection.
func (s *MultiAdminServer) dialGatewayByID(ctx context.Context, id *clustermetadatapb.ID) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
