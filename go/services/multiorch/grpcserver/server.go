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

package grpcserver

import (
	"context"
	"log/slog"

	"google.golang.org/grpc"

	multiorchpb "github.com/multigres/multigres/go/pb/multiorch"
	"github.com/multigres/multigres/go/services/multiorch/consensus"
	"github.com/multigres/multigres/go/services/multiorch/recovery"
)

// MultiOrchServer implements the MultiOrchService gRPC service.
// It provides diagnostic information about the multiorch recovery engine,
// including detected problems and shard health status.
type MultiOrchServer struct {
	multiorchpb.UnimplementedMultiOrchServiceServer
	engine      *recovery.Engine
	coordinator *consensus.Coordinator
	logger      *slog.Logger
}

// NewMultiOrchServer creates a new MultiOrchServer instance.
func NewMultiOrchServer(engine *recovery.Engine, coordinator *consensus.Coordinator, logger *slog.Logger) *MultiOrchServer {
	_ = "STUB: not implemented"
	return nil
}

// RegisterWithGRPCServer registers the MultiOrchService with the provided gRPC server.
func (s *MultiOrchServer) RegisterWithGRPCServer(grpcServer *grpc.Server) {
	_ = "STUB: not implemented"
	return
}

// GetShardStatus returns diagnostic information for a specific shard.
// It includes detected problems, pooler health, and shard summary.
func (s *MultiOrchServer) GetShardStatus(
	ctx context.Context,
	req *multiorchpb.ShardStatusRequest,
) (*multiorchpb.ShardStatusResponse, error) {
	_ = "STUB: not implemented"
	// Validate that this shard is in our watch targets
	return nil, nil
}

// Get all detected problems from the engine

// Filter problems for the requested shard

// DisableRecovery stops the recovery loop and waits for in-flight actions to complete.
func (s *MultiOrchServer) DisableRecovery(_ context.Context, _ *multiorchpb.DisableRecoveryRequest) (*multiorchpb.DisableRecoveryResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EnableRecovery resumes the recovery loop.
func (s *MultiOrchServer) EnableRecovery(_ context.Context, _ *multiorchpb.EnableRecoveryRequest) (*multiorchpb.EnableRecoveryResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetRecoveryStatus returns whether recovery is currently enabled or disabled.
func (s *MultiOrchServer) GetRecoveryStatus(_ context.Context, _ *multiorchpb.GetRecoveryStatusRequest) (*multiorchpb.GetRecoveryStatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TriggerRecoveryNow immediately executes recovery cycles until no problems remain
// or the request context times out. Returns problem codes that remain unresolved.
func (s *MultiOrchServer) TriggerRecoveryNow(ctx context.Context, req *multiorchpb.TriggerRecoveryNowRequest) (*multiorchpb.TriggerRecoveryNowResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Subtract 200ms from deadline to allow time for response overhead.

// ApplyCertifiedRuleChange installs a new shard rule using a fully-populated
// externally certified revocation. See proto/multiorchservice.proto for the
// shape contract — multiorch is a pure executor and the caller must populate
// every identity and timing field.
func (s *MultiOrchServer) ApplyCertifiedRuleChange(
	ctx context.Context,
	req *multiorchpb.ApplyCertifiedRuleChangeRequest,
) (*multiorchpb.ApplyCertifiedRuleChangeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// buildPoolerHealthList creates pooler health snapshots for the requested shard.
func (s *MultiOrchServer) buildPoolerHealthList(req *multiorchpb.ShardStatusRequest) []*multiorchpb.PoolerHealth {
	_ = "STUB: not implemented"
	return nil
}

// Get pooler type string
