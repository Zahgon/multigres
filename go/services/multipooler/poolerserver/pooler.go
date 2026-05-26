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

// Package poolerserver implements the multipooler gRPC server
package poolerserver

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/multigres/multigres/go/common/queryservice"
	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	"github.com/multigres/multigres/go/pb/query"
	"github.com/multigres/multigres/go/services/multipooler/connpoolmanager"
	"github.com/multigres/multigres/go/services/multipooler/executor"
	"github.com/multigres/multigres/go/services/multipooler/pubsub"
)

// QueryPoolerServer is the core pooler implementation for query serving.
// It encapsulates the components required to manage query execution
// (e.g. pooling, execution, transactions).
//
// In the future, components like the transaction engine and query engine
// should be added here. The current executor is temporary until those
// engines are introduced.
//
// The lifecycle of the pooler is managed by the MultiPoolerManager.
// The connection pool manager (connpoolmanager) handles all database connections,
// including per-user connection pools with trust/peer authentication.
type QueryPoolerServer struct {
	logger      *slog.Logger
	poolManager connpoolmanager.PoolManager
	executor    *executor.Executor

	// tableGroup and shard identify which tablegroup/shard this pooler serves.
	// Set at construction time and immutable.
	tableGroup string
	shard      string

	mu             sync.Mutex
	poolerType     clustermetadatapb.PoolerType
	servingStatus  clustermetadatapb.PoolerServingStatus
	healthProvider HealthProvider

	// pubsubListener is the shared LISTEN/NOTIFY listener, set by MultiPoolerManager.
	pubsubListener *pubsub.Listener

	// shuttingDown is true during the graceful drain phase (between receiving
	// NOT_SERVING and completing the transition). During this phase, new requests
	// are rejected but existing reserved connections are allowed to finish.
	shuttingDown bool

	// gracePeriod is how long OnStateChange waits for in-flight connections to drain
	// before force-closing reserved connections. Configured via --connpool-drain-grace-period.
	gracePeriod time.Duration

	// stateChanged is closed and re-created on every state transition.
	// AwaitStateChange blocks on this channel to be notified when the
	// pooler's type and serving status are updated.
	stateChanged chan struct{}

	// drainStats holds drain-observability metrics. Never nil.
	drainStats *drainStats
}

// NewQueryPoolerServer creates a new QueryPoolerServer instance with the given pool manager
// and health provider.
// The pool manager must already be opened before calling this function.
// The health provider is used by StreamPoolerHealth to provide health updates to clients.
// gracePeriod controls how long OnStateChange waits for in-flight connections to drain
// during NOT_SERVING transitions before force-closing reserved connections.
func NewQueryPoolerServer(logger *slog.Logger, poolManager connpoolmanager.PoolManager, poolerID *clustermetadatapb.ID, tableGroup, shard string, healthProvider HealthProvider, gracePeriod time.Duration, vpidStampEnabled bool) *QueryPoolerServer {
	_ = "STUB: not implemented"
	return nil
}

// OnStateChange transitions the query service to match the new serving state.
// Implements PoolerController interface.
//
// This method is only called by the StateManager, which serializes calls behind
// a mutex. Concurrent calls are not possible or expected.
//
// For NOT_SERVING transitions, this performs a two-phase graceful drain:
//  1. Set shuttingDown=true to reject new requests (existing reserved connections continue)
//  2. Wait for in-flight connections to drain (up to gracePeriod)
//  3. Set servingStatus=NOT_SERVING
func (s *QueryPoolerServer) OnStateChange(ctx context.Context, poolerType clustermetadatapb.PoolerType, servingStatus clustermetadatapb.PoolerServingStatus) error {
	_ = "STUB: not implemented"
	return nil
}

// NOT_SERVING: begin graceful drain.
// The poolerType is NOT updated yet — in-flight requests on reserved
// connections (e.g., COMMIT after a demotion) must still see the old type
// so that checkTargetLocked allows them to complete.

// Wait for in-flight connections to drain.
// If gracePeriod > 0, the wait is bounded and reserved connections are force-closed on timeout.
// If gracePeriod == 0, the wait is unbounded (drain must complete before transition finishes).

// Force-close all reserved connections to prevent them from being used
// in a non-serving state. This kills backend processes and returns
// connections to the pool.

// Complete the transition. The poolerType is set here (after drain) so that
// in-flight requests saw the old type throughout the drain phase.

// StartRequest checks whether a new request should be admitted.
// Returns nil if the request is allowed, or an error if it should be rejected.
//
// It validates the target against this pooler's identity (tablegroup, shard, type)
// via checkTarget, then checks serving status. A target mismatch returns MTF01,
// which causes the gateway to buffer the request until the correct pooler is available.
//
// During graceful shutdown (shuttingDown=true), requests on existing reserved
// connections (allowOnShutdown=true) are still admitted so that in-flight
// transactions can complete. New reservations and fresh queries are rejected.
func (s *QueryPoolerServer) StartRequest(target *query.Target, allowOnShutdown bool) error {
	_ = "STUB: not implemented"
	return nil
}

// checkTargetLocked validates that the request's target matches this pooler's identity.
//
// Validation rules:
//   - TableGroup and Shard must always match. A mismatch is a routing bug (MTD01).
//   - PoolerType: if the target is PRIMARY but this pooler is a REPLICA, the gateway
//     has stale topology (likely a demotion happened). Returns MTF01 to trigger
//     buffering until the new primary is discovered. All other type combinations
//     are allowed (PRIMARY serves both PRIMARY and REPLICA traffic).
//
// Caller must hold s.mu.
func (s *QueryPoolerServer) checkTargetLocked(target *query.Target) error {
	_ = "STUB: not implemented"
	return nil
}

// A PRIMARY request hitting a REPLICA means the gateway thinks this pooler is
// still the primary, but it was demoted. Return MTF01 to trigger buffering.

// notifyStateChangedLocked closes the stateChanged channel to wake any
// AwaitStateChange callers, then replaces it with a fresh channel.
// Caller must hold s.mu.
func (s *QueryPoolerServer) notifyStateChangedLocked() { _ = "STUB: not implemented"; return }

// AwaitStateChange blocks until the pooler's type and serving status match
// the given targets, or ctx is cancelled. Used by the health streamer to
// ensure the query server is ready before broadcasting the new state.
func (s *QueryPoolerServer) AwaitStateChange(ctx context.Context, poolerType clustermetadatapb.PoolerType, servingStatus clustermetadatapb.PoolerServingStatus) {
	_ = "STUB: not implemented"
	return
}

// IsServing returns true if currently serving queries.
// Implements PoolerController interface.
func (s *QueryPoolerServer) IsServing() bool { _ = "STUB: not implemented"; return false }

// IsHealthy checks if the controller is healthy.
// Implements PoolerController interface.
func (s *QueryPoolerServer) IsHealthy() error { _ = "STUB: not implemented"; return nil }

// The pool manager handles connection health internally.
// If we have the executor and pool manager, we're healthy.

// RegisterGRPCServices registers gRPC services (called by manager during startup).
// Implements PoolerController interface.
func (s *QueryPoolerServer) RegisterGRPCServices() { _ = "STUB: not implemented"; return }

// StartServiceForTests is a convenience method for tests to initialize and start the pooler.
// Following Vitess pattern: "StartService is only used for testing."
// It sets the serving type to PRIMARY + SERVING.
func (s *QueryPoolerServer) StartServiceForTests() error { _ = "STUB: not implemented"; return nil }

// Executor returns the executor instance for use by gRPC service handlers.
// Implements PoolerController interface.
// Returns error if the pooler is not initialized.
//
// Note: Callers must call StartRequest() before Executor() to check serving state.
// StartRequest returns MTF01 when not serving or shutting down, which the gateway
// uses to trigger failover buffering.
func (s *QueryPoolerServer) Executor() (queryservice.QueryService, error) {
	_ = "STUB: not implemented"
	return *new(queryservice.QueryService), nil
}

// SetPubSubListener sets the PubSub listener on the pooler server.
// The listener is created and managed by the MultiPoolerManager.
func (s *QueryPoolerServer) SetPubSubListener(l *pubsub.Listener) {
	_ = "STUB: not implemented"
	return

	// PubSubListener returns the shared PubSub listener (may be nil).
}

func (s *QueryPoolerServer) PubSubListener() *pubsub.Listener {
	_ = "STUB: not implemented"
	return nil

	// PoolManager returns the pool manager instance.
	// This is used by GetAuthCredentials to query pg_authid using an admin connection,
	// which works even before the executor is fully initialized during bootstrap.
}

func (s *QueryPoolerServer) PoolManager() connpoolmanager.PoolManager {
	_ = "STUB: not implemented"
	return *new(connpoolmanager.PoolManager)
}

// InternalQueryService returns the executor as an InternalQueryService for internal queries.
// Implements PoolerController interface.
func (s *QueryPoolerServer) InternalQueryService() executor.InternalQueryService {
	_ = "STUB: not implemented"
	// Explicit nil check required: returning a nil *Executor directly would produce a
	// non-nil interface value wrapping a nil pointer, causing callers' == nil checks to
	// pass but method calls on the interface to panic.
	return *new(executor.InternalQueryService)
}

// HealthProvider returns the health provider for streaming health updates.
func (s *QueryPoolerServer) HealthProvider() HealthProvider {
	_ = "STUB: not implemented"
	return *new(HealthProvider)
}
