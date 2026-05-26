// Copyright 2026 Supabase, Inc.
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

package multigateway

import (
	"context"

	"google.golang.org/grpc"

	"github.com/multigres/multigres/go/common/preparedstatement"
	multigatewaymanagerpb "github.com/multigres/multigres/go/pb/multigatewaymanager"
	multigatewaymanagerdatapb "github.com/multigres/multigres/go/pb/multigatewaymanagerdata"
	"github.com/multigres/multigres/go/services/multigateway/handler"
	"github.com/multigres/multigres/go/services/multigateway/handler/queryregistry"
)

// ManagerServer implements multigatewaymanagerpb.MultiGatewayManagerServer,
// exposing in-process diagnostic snapshots (registry, consolidator) for
// rendering in multiadmin-web.
type ManagerServer struct {
	multigatewaymanagerpb.UnimplementedMultiGatewayManagerServer

	registry *queryregistry.Registry
	handler  *handler.MultiGatewayHandler
}

// NewManagerServer constructs a ManagerServer wired to the gateway's shared
// registry and handler (for prepared-statement consolidator access).
func NewManagerServer(registry *queryregistry.Registry, h *handler.MultiGatewayHandler) *ManagerServer {
	_ = "STUB: not implemented"
	return nil
}

// RegisterWithGRPCServer registers the manager service on the given gRPC server.
func (s *ManagerServer) RegisterWithGRPCServer(grpcServer *grpc.Server) {
	_ = "STUB: not implemented"
	return
}

// GetQueryRegistry returns the per-fingerprint registry snapshot. limit and
// min_calls bound the response so a saturated registry doesn't ship megabytes
// of trend data per poll; tracked_fingerprints in the response always reflects
// the full registry size, independent of any filtering applied.
func (s *ManagerServer) GetQueryRegistry(_ context.Context, req *multigatewaymanagerpb.GetQueryRegistryRequest) (*multigatewaymanagerpb.GetQueryRegistryResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetConsolidatorStats returns a snapshot of the prepared-statement consolidator.
func (s *ManagerServer) GetConsolidatorStats(_ context.Context, _ *multigatewaymanagerpb.GetConsolidatorStatsRequest) (*multigatewaymanagerpb.GetConsolidatorStatsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func snapshotToProto(s *queryregistry.Snapshot) *multigatewaymanagerdatapb.QueryStatSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func consolidatorStatsToProto(s preparedstatement.ConsolidatorStats) *multigatewaymanagerdatapb.ConsolidatorStats {
	_ = "STUB: not implemented"
	return nil
}
