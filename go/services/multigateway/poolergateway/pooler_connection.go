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

package poolergateway

import (
	"context"
	"errors"
	"log/slog"
	"sync"
	"sync/atomic"
	"time"

	"github.com/multigres/multigres/go/common/queryservice"
	"github.com/multigres/multigres/go/common/topoclient"
	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	"github.com/multigres/multigres/go/pb/multipoolerservice"
	"github.com/multigres/multigres/go/pb/query"
	"github.com/multigres/multigres/go/tools/retry"

	"google.golang.org/grpc"
)

// errPoolerUninitialized is the initial error before health stream connects.
var errPoolerUninitialized = errors.New("pooler health not initialized")

// PoolerHealth represents the health state received from a multipooler.
// This is a snapshot of health state that can be safely passed around
// without synchronization.
type PoolerHealth struct {
	// Target identifies the tablegroup, shard, and pooler type.
	Target *query.Target

	// PoolerID identifies the multipooler instance.
	PoolerID *clustermetadatapb.ID

	// ServingStatus is the serving state reported by the pooler.
	ServingStatus clustermetadatapb.PoolerServingStatus

	// LeaderObservation contains the pooler's view of who the consensus leader is.
	// Used for term-based leader reconciliation.
	LeaderObservation *multipoolerservice.LeaderObservation

	// ReplicationLagNs is the replication lag in nanoseconds reported by the pooler.
	// Zero on the primary or when not yet measured.
	ReplicationLagNs int64

	// LastError is the most recent error from the health stream.
	LastError error

	// LastResponse is when we last received a health update.
	LastResponse time.Time
}

// IsServing returns true if the pooler is serving traffic.
func (h *PoolerHealth) IsServing() bool { _ = "STUB: not implemented"; return false }

// SimpleCopy returns a shallow copy of the PoolerHealth.
// This is not a deep copy: pointer fields (Target, PoolerID, LeaderObservation)
// reference the same underlying objects. This is safe because these proto objects
// are treated as immutable - they are never modified after creation.
// Returns a shallow copy that is safe to read concurrently.
func (h *PoolerHealth) SimpleCopy() *PoolerHealth { _ = "STUB: not implemented"; return nil }

// PoolerConnection manages a single gRPC connection to a multipooler instance.
// It wraps a QueryService and provides access to pooler metadata.
//
// A PoolerConnection exists if and only if we are actively connected to the pooler.
// The LoadBalancer creates and destroys PoolerConnections based on discovery events.
//
// The connection maintains a health stream to the multipooler and tracks serving state.
// Only connections that are serving should be used for query routing.
type PoolerConnection struct {
	// poolerInfo contains the pooler metadata from discovery.
	// Accessed atomically to avoid data races between UpdatePoolerInfo and readers.
	poolerInfo atomic.Pointer[topoclient.MultiPoolerInfo]

	// conn is the underlying gRPC connection
	conn *grpc.ClientConn

	// client is the gRPC client for health streaming
	client multipoolerservice.MultiPoolerServiceClient

	// queryService handles query execution over gRPC
	queryService queryservice.QueryService

	// logger for debugging
	logger *slog.Logger

	// ctx and cancel control the health stream goroutine lifecycle.
	// cancel must be called before discarding PoolerConnection to ensure
	// the checkConn goroutine terminates.
	ctx    context.Context
	cancel context.CancelFunc

	// healthMu protects the health field
	healthMu sync.Mutex

	// health contains the current health state from the health stream.
	// Updated atomically as a unit when new health responses arrive.
	health *PoolerHealth

	// healthTimedOut indicates if the health stream has timed out.
	// Accessed atomically because there's a race between the timeout
	// goroutine and the stream processing.
	healthTimedOut atomic.Bool

	// onHealthUpdate is called when health state changes.
	// Used by LoadBalancer to update its routing decisions.
	onHealthUpdate func(*PoolerConnection)
}

// NewPoolerConnection creates a new connection to a multipooler instance and
// starts health streaming automatically.
//
// The onHealthUpdate callback is invoked when health state changes. It may be
// nil if health updates don't need to be observed.
//
// Close must be called when the connection is no longer needed to stop the
// health stream goroutine and release resources.
func NewPoolerConnection(
	ctx context.Context,
	pooler *clustermetadatapb.MultiPooler,
	logger *slog.Logger,
	grpcDialOpt grpc.DialOption,
	onHealthUpdate func(*PoolerConnection),
) (*PoolerConnection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create gRPC connection with telemetry attributes

// Derive a cancellable context from the service-lifetime context for the
// health stream goroutine. This ensures proper shutdown propagation.

// Create QueryService wrapper

// Initialize health state to NOT_SERVING until health stream provides data.

// Start health stream goroutine

// ID returns the unique identifier for this pooler connection.
func (pc *PoolerConnection) ID() string { _ = "STUB: not implemented"; return "" }

// Cell returns the cell where this pooler is located.
func (pc *PoolerConnection) Cell() string { _ = "STUB: not implemented"; return "" }

// Type returns the pooler type (PRIMARY or REPLICA).
func (pc *PoolerConnection) Type() clustermetadatapb.PoolerType {
	_ = "STUB: not implemented"
	return *new(clustermetadatapb.PoolerType)
}

// UpdatePoolerInfo updates the pooler metadata (e.g., when type changes from UNKNOWN to PRIMARY).
// This is called when topology watch detects updates to the pooler.
func (pc *PoolerConnection) UpdatePoolerInfo(pooler *clustermetadatapb.MultiPooler) {
	_ = "STUB: not implemented"
	return
}

// PoolerInfo returns the underlying pooler metadata.
func (pc *PoolerConnection) PoolerInfo() *topoclient.MultiPoolerInfo {
	_ = "STUB: not implemented"
	return nil
}

// ServiceClient returns the MultiPoolerServiceClient for admin operations.
// This can be used for authentication, health checks, and other system-level operations.
func (pc *PoolerConnection) ServiceClient() multipoolerservice.MultiPoolerServiceClient {
	_ = "STUB: not implemented"

	// QueryService returns the query execution service for this connection.
	return *new(multipoolerservice.MultiPoolerServiceClient)
}

func (pc *PoolerConnection) QueryService() queryservice.QueryService {
	_ = "STUB: not implemented"
	return *

	// Close stops the health stream goroutine and closes the gRPC connection.
	new(queryservice.QueryService)
}

func (pc *PoolerConnection) Close() error { _ = "STUB: not implemented"; return nil }

// Cancel the health stream context to stop the checkConn goroutine

// Health returns the current health state.
// The returned PoolerHealth is a snapshot that can be safely used without
// synchronization. We don't deep-copy because the PoolerHealth object is
// never modified after creation.
func (pc *PoolerConnection) Health() *PoolerHealth { _ = "STUB: not implemented"; return nil }

// checkConn performs health checking on the pooler connection.
// It continuously attempts to maintain a health stream, retrying with
// exponential backoff on failures.
func (pc *PoolerConnection) checkConn() { _ = "STUB: not implemented"; return }

// Context cancelled - connection is being closed.

// Create a separate context for this stream attempt.
// This allows the staleness timer to cancel the stream independently.

// Stream health responses. This blocks until an error or context cancellation.

// Always cancel the stream context to clean up resources.

// streamHealth opens a health stream and processes responses until an error occurs.
// streamCancel is called by the staleness timer to unblock stream.Recv().
// On each successful message, streamRetrier is reset to use minimum backoff.
func (pc *PoolerConnection) streamHealth(
	streamCtx context.Context,
	streamCancel context.CancelFunc,
	streamRetrier *retry.Retry,
) error {
	_ = "STUB: not implemented"
	return nil

	// Reset healthTimedOut from any previous stream attempt so a shutdown during
	// this attempt isn't misclassified as a staleness timeout.
}

// Open the health stream.

// Set up staleness timer. If no message is received within the timeout,
// the timer cancels the stream context to unblock stream.Recv().

// Process responses from the stream.

// Stream context cancelled (either staleness timeout or shutdown).

// We received a message successfully.

// Reset backoff since we got a successful message.

// Update staleness timeout from server recommendation if provided.

// Process the health response.

// processHealthResponse updates the health state from a StreamPoolerHealthResponse.
// Creates a new immutable PoolerHealth snapshot.
func (pc *PoolerConnection) processHealthResponse(response *multipoolerservice.StreamPoolerHealthResponse) {
	_ = "STUB: not implemented"
	return

	// Build new health snapshot from the response.
}

// Log state changes.

// Notify listener of health update.

// setHealthError updates the health state to reflect an error while preserving
// existing metadata. Uses SimpleCopy to create a new snapshot, then updates
// error-related fields. This ensures forward compatibility: any new fields
// added to PoolerHealth will be automatically preserved.
func (pc *PoolerConnection) setHealthError(err error) { _ = "STUB: not implemented"; return }

// Notify listener that health changed (pooler is now unhealthy).
