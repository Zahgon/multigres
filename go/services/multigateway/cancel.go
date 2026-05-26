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
	"log/slog"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/grpc"

	"github.com/multigres/multigres/go/common/topoclient"
	multigatewayservicepb "github.com/multigres/multigres/go/pb/multigatewayservice"
)

const (
	// prefixCacheRefreshInterval is how often the prefix cache is fully
	// rebuilt from topo, regardless of cache hits.
	prefixCacheRefreshInterval = 5 * time.Minute

	// grpcIdleTimeout is how long a gateway-to-gateway gRPC connection sits
	// idle before gRPC automatically closes the underlying transport.
	grpcIdleTimeout = 5 * time.Minute
)

// CancelManager handles cross-gateway query cancellation.
// It implements both server.CancelHandler (for incoming PostgreSQL cancel requests)
// and MultiGatewayServiceServer (for gRPC forwarded cancel requests).
//
// Query cancellation is best-effort: if a forwarded cancel fails, the query
// will eventually complete or time out on its own.
type CancelManager struct {
	multigatewayservicepb.UnimplementedMultiGatewayServiceServer

	// primaryCancelFn attempts to cancel a connection on the primary listener.
	primaryCancelFn func(pid, secret uint32) bool
	// replicaCancelFn attempts to cancel a connection on the replica-reads listener.
	// Nil when no replica listener is configured.
	replicaCancelFn func(pid, secret uint32) bool
	// ownPrefix is this gateway's PID prefix.
	ownPrefix uint32
	// ts is the topology store for discovering other gateways.
	ts topoclient.Store
	// logger for cancel operations.
	logger *slog.Logger

	// prefixCache maps PID prefix to gateway gRPC address. Replaced atomically
	// on cache miss or periodic refresh; reads are lock-free.
	prefixCache atomic.Pointer[map[uint32]string]

	// clientsMu protects clients.
	clientsMu sync.Mutex
	// clients caches gRPC connections by address. gRPC's WithIdleTimeout
	// handles closing unused transports automatically.
	clients map[string]*gatewayConn

	// transportCreds configures transport security for gateway-to-gateway gRPC dials.
	transportCreds grpc.DialOption

	// stop cancels the background prefix cache refresh goroutine.
	stop context.CancelFunc
}

// gatewayConn holds a gRPC client and its underlying connection.
type gatewayConn struct {
	client multigatewayservicepb.MultiGatewayServiceClient
	conn   *grpc.ClientConn
}

// NewCancelManager creates a new CancelManager.
// primaryCancelFn handles cancels for the primary listener.
// replicaCancelFn handles cancels for the replica-reads listener (nil when no
// replica listener is configured).
func NewCancelManager(
	primaryCancelFn func(pid, secret uint32) bool,
	replicaCancelFn func(pid, secret uint32) bool,
	ownPrefix uint32,
	ts topoclient.Store,
	logger *slog.Logger,
	transportCreds grpc.DialOption,
) *CancelManager {
	_ = "STUB: not implemented"
	return nil
}

// RegisterWithGRPCServer registers the CancelManager as a gRPC service.
func (cm *CancelManager) RegisterWithGRPCServer(grpcServer *grpc.Server) {
	_ = "STUB: not implemented"
	return
}

// handleCancel decodes the PID prefix and either cancels locally or forwards
// to the owning gateway. The replica flag indicates whether the cancel arrived
// on a replica-reads listener.
func (cm *CancelManager) handleCancel(ctx context.Context, processID, secretKey uint32, replica bool) {
	_ = "STUB: not implemented"
	return
}

// Forward to the gateway that owns this PID prefix.

// ForListener returns a server.CancelHandler that routes cancel requests
// through this CancelManager with the given connection type baked in.
// This avoids adding replica awareness to the generic server package.
func (cm *CancelManager) ForListener(replica bool) *ListenerCancelHandler {
	_ = "STUB: not implemented"
	return nil
}

// ListenerCancelHandler adapts CancelManager to the server.CancelHandler
// interface for a specific listener type (primary or replica).
type ListenerCancelHandler struct {
	cm      *CancelManager
	replica bool
}

// HandleCancelRequest implements server.CancelHandler.
func (h *ListenerCancelHandler) HandleCancelRequest(ctx context.Context, processID, secretKey uint32) {
	_ = "STUB: not implemented"
	return
}

// cancelLocal dispatches a cancel to the correct local listener based on connection type.
func (cm *CancelManager) cancelLocal(processID, secretKey uint32, replica bool) {
	_ = "STUB: not implemented"
	return
}

// CancelQuery implements MultiGatewayServiceServer.
// This is called by other gateways forwarding cancel requests via gRPC.
func (cm *CancelManager) CancelQuery(ctx context.Context, req *multigatewayservicepb.CancelQueryRequest) (*multigatewayservicepb.CancelQueryResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// forwardCancel finds the gateway with the given PID prefix and forwards the cancel request.
// Cancellation is best-effort — no retries on failure.
func (cm *CancelManager) forwardCancel(ctx context.Context, targetPrefix, processID, secretKey uint32, replica bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Cache miss — rebuild from topo and try again.

// rebuildPrefixCache reads all gateways from topo and atomically replaces the
// prefix cache.
func (cm *CancelManager) rebuildPrefixCache(ctx context.Context) { _ = "STUB: not implemented"; return }

// refreshPrefixCachePeriodically rebuilds the prefix cache on a regular interval
// so that gateway additions/removals are picked up even without a cache miss.
func (cm *CancelManager) refreshPrefixCachePeriodically(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

// getClient returns a cached gRPC client for the given address, creating one if needed.
// Connections use gRPC's built-in idle timeout to close unused transports.
func (cm *CancelManager) getClient(addr string) (multigatewayservicepb.MultiGatewayServiceClient, error) {
	_ = "STUB: not implemented"
	return *new(multigatewayservicepb.MultiGatewayServiceClient), nil
}

// Close stops the background refresh goroutine and closes all cached gRPC connections.
func (cm *CancelManager) Close() { _ = "STUB: not implemented"; return }
