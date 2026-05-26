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

package recovery

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/multigres/multigres/go/common/rpcclient"
	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	multiorchdatapb "github.com/multigres/multigres/go/pb/multiorchdata"
	multipoolermanagerdatapb "github.com/multigres/multigres/go/pb/multipoolermanagerdata"
	"github.com/multigres/multigres/go/services/multiorch/store"
)

const (
	// streamReconnectInitialBackoff is the initial wait before retrying a
	// failed ManagerHealthStream connection.
	streamReconnectInitialBackoff = 1 * time.Second

	// streamReconnectMaxBackoff caps the exponential backoff between retries.
	streamReconnectMaxBackoff = 30 * time.Second
)

// streamEntry holds the lifecycle handles for one active stream goroutine.
type streamEntry struct {
	cancel context.CancelFunc

	// mu protects stream; set to the live gRPC stream after the start message
	// is sent, and cleared on stream exit.
	mu     sync.Mutex
	stream rpcclient.ManagerHealthStream
}

// HealthStream maintains one ManagerHealthStream stream per pooler. It replaces
// the polling loop: instead of periodically calling the Status RPC, each pooler
// pushes health snapshots to multiorch via a long-lived gRPC stream.
//
// When a snapshot arrives the HealthStream writes the same health fields into
// the pooler store that the old pollPooler function wrote on success. On stream
// disconnect the pooler is marked unreachable and reconnection is attempted
// with exponential backoff (1s → 2s → … → 30s cap).
//
// The orchestrator sends its preferred snapshot_interval and staleness_timeout
// in the start message. The server echoes back the actual values it will use in
// a ManagerHealthStreamStartResponse, which the client uses to arm its staleness
// watchdog.
type HealthStream struct {
	logger    *slog.Logger
	rpcClient rpcclient.MultiPoolerClient
	store     *store.PoolerStore

	// snapshotInterval is requested from the server as the proactive snapshot
	// tick rate. Zero means use the server default (currently 5s).
	snapshotInterval time.Duration

	// stalenessTimeout is sent to the server and used to arm the staleness
	// watchdog (seeded from the start response). Zero means server default
	// (timeouts.DefaultHealthStreamStalenessTimeout).
	stalenessTimeout time.Duration

	// Active stream goroutines, keyed by pooler ID string.
	mu      sync.Mutex
	streams map[string]*streamEntry

	// Parent context; cancelled by Stop().
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

// Option is a functional option for NewHealthStream.
type Option func(*HealthStream)

// WithSnapshotInterval sets the proactive snapshot interval sent to the server
// in the start message.
func WithSnapshotInterval(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithStalenessTimeout sets the staleness timeout sent to the server and used
// to arm the client-side staleness watchdog. Intended for tests.
func WithStalenessTimeout(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// NewHealthStream creates a HealthStream.
//
// Call Start() for each pooler that should be monitored.
func NewHealthStream(
	ctx context.Context,
	rpcClient rpcclient.MultiPoolerClient,
	poolerStore *store.PoolerStore,
	logger *slog.Logger,
	options ...Option,
) *HealthStream {
	_ = "STUB: not implemented"
	return nil
}

// Shutdown cancels all active streams and waits for their goroutines to exit.
func (hs *HealthStream) Shutdown() { _ = "STUB: not implemented"; return }

// Start starts a health stream for id.
// If a stream is already running for this pooler the call is a no-op.
// The pooler's MultiPooler metadata is read from the store on each
// reconnect attempt so topology updates are automatically picked up.
func (hs *HealthStream) Start(id *clustermetadatapb.ID) { _ = "STUB: not implemented"; return }

// Stop the health stream for a pooler.
//
// The stream goroutine will exit and the pooler will be marked unreachable
// until a new stream is started. The pooler's MultiPooler metadata must remain
// in the store for the stream to reconnect if Start() is called again.
//
// If no stream is running for this pooler the call is a no-op.
func (hs *HealthStream) Stop(id *clustermetadatapb.ID) { _ = "STUB: not implemented"; return }

// The goroutine removes itself from hs.streams via its defer so we
// don't need to delete the entry here; just cancel it and let the
// goroutine clean up.

// runStream manages the lifecycle of one stream, reconnecting with backoff on failure.
// It reads the latest MultiPooler metadata from the store on each reconnect attempt
// so hostname/port changes are picked up automatically.
func (hs *HealthStream) runStream(ctx context.Context, poolerID string, entry *streamEntry) {
	_ = "STUB: not implemented"
	return
}

// Read current pooler metadata from store on every attempt.

// Stream was successfully established before failing — reset backoff.

// streamOnce opens one ManagerHealthStream and reads until the stream fails or
// the context is cancelled. Returns (connected, err): connected is true if the
// stream was established before any error occurred.
func (hs *HealthStream) streamOnce(ctx context.Context, poolerID string, poolerHealth *multiorchdatapb.PoolerHealthState, entry *streamEntry) (connected bool, _ error) {
	_ = "STUB: not implemented"
	// Build the start request, sending the orchestrator's preferred timing.
	// Zero values are omitted so the server uses its own defaults.
	return false, nil
}

// Seed the staleness watchdog before any message is received. This is the
// value we sent; the server will confirm (or adjust) it in the start response,
// at which point we reset the watchdog to the echoed value.

// Staleness watchdog: cancel the stream if no message arrives within the
// timeout. This catches the "server goroutine stuck but TCP alive" failure
// mode that gRPC keepalive does not cover.
//
// The watchdog context is passed to ManagerHealthStream so cancelling it
// terminates the gRPC stream and causes stream.Recv() to return an error.

// resetCh carries the new timer duration whenever a message is received.
// Buffered so the recv loop never blocks on the watchdog goroutine.

// Send the start message with the negotiated timing preferences.

// Read the start response — the first server message confirms the actual
// timing values the server will use.

// Determine the effective staleness for the watchdog:
//   - If a local override is set (WithStalenessTimeout), use it directly.
//     This preserves sub-second precision used in tests.
//   - Otherwise, use the server-confirmed value from the start response.

// Expose the live stream so Poll() can send requests.

// Distinguish a staleness-triggered cancellation from an external one
// so the caller can log a useful error message.

// Reset the staleness watchdog. Prefer the local override (which
// preserves sub-second precision); fall back to the server-echoed value.

// A reset is already pending; the watchdog will pick it up.

// Poll sends a poll request on the active stream for poolerID, triggering an
// immediate health snapshot from the pooler. Returns an error if no stream is
// active or the send fails.
func (hs *HealthStream) Poll(id *clustermetadatapb.ID) error { _ = "STUB: not implemented"; return nil }

// applySnapshot writes health fields from a snapshot into the pooler store.
// This mirrors the field writes performed by the old pollPooler function on success.
func (hs *HealthStream) applySnapshot(ctx context.Context, poolerID string, poolerHealth *multiorchdatapb.PoolerHealthState, snapshot *multipoolermanagerdatapb.ManagerHealthSnapshot) {
	_ = "STUB: not implemented"
	return
}

// NOTE: when PostgresReady is false, LastPostgresReadyTime is intentionally
// left at its previous value so callers can reason about "last known good" time.

// markConnected records that the stream is connected in the pooler store.
func (hs *HealthStream) markConnected(poolerID string) { _ = "STUB: not implemented"; return }

// markDisconnected records that the stream is disconnected and the pooler
// should be treated as unreachable.
func (hs *HealthStream) markDisconnected(poolerID string) { _ = "STUB: not implemented"; return }
