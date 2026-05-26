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

// Package pubsub implements a shared PG LISTEN/NOTIFY listener
// that fans out notifications to multiple gateway client sessions.
package pubsub

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/multigres/multigres/go/common/sqltypes"
	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	"github.com/multigres/multigres/go/services/multipooler/connpoolmanager"
	"github.com/multigres/multigres/go/services/multipooler/pools/reserved"
)

// request types for the serialized event loop.
type requestType int

const (
	reqSubscribe requestType = iota
	reqUnsubscribe
	reqUnsubscribeAll
)

type request struct {
	typ     requestType
	channel string
	subCh   chan *sqltypes.Notification
	done    chan struct{} // closed when request is processed
}

// Listener manages a single dedicated PG connection for LISTEN/NOTIFY,
// with channel refcounting and fan-out to subscribers.
//
// Architecture: split read/write on a single TCP connection.
// TCP is full-duplex; bufferedReader and bufferedWriter are independent bufio
// instances. The reader goroutine exclusively reads; the event loop exclusively
// writes. Because the reader never holds the connection, the event loop can
// send LISTEN/UNLISTEN at any time without tearing down and re-establishing
// the connection, which would lose notifications for existing channels.
type Listener struct {
	poolManager connpoolmanager.PoolManager
	logger      *slog.Logger
	metrics     *PubSubMetrics

	// requests is the serialized command channel for the event loop.
	requests chan request

	// cancel stops the background goroutines.
	cancel context.CancelFunc

	// wg tracks background goroutines.
	wg sync.WaitGroup

	// stopped is closed when the event loop exits. Public methods select on this
	// to avoid blocking forever when the Listener is not running.
	stopped chan struct{}
}

// NewListener creates a new PubSubListener. Call Start() to begin.
func NewListener(poolManager connpoolmanager.PoolManager, logger *slog.Logger, metrics *PubSubMetrics) *Listener {
	_ = "STUB: not implemented"
	return nil
}

// initially stopped

// Start begins the background event loop. Idempotent — no-op if already running.
func (l *Listener) Start(ctx context.Context) { _ = "STUB: not implemented"; return }

// already running

// Stop shuts down the listener and waits for goroutines to exit. Idempotent.
func (l *Listener) Stop() { _ = "STUB: not implemented"; return }

// not running

// OnStateChange implements manager.StateAware. The listener runs only when the
// multipooler is PRIMARY+SERVING; it is stopped on any other state transition.
func (l *Listener) OnStateChange(_ context.Context, poolerType clustermetadatapb.PoolerType, servingStatus clustermetadatapb.PoolerServingStatus) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gocritic // Long-lived background listener; must not use the transient state-change ctx.

// Unsubscribe removes a subscriber from the given channel.
// Returns immediately if the Listener is stopped.
func (l *Listener) Unsubscribe(channel string, subCh chan *sqltypes.Notification) {
	_ = "STUB: not implemented"
	return
}

// UnsubscribeAll removes a subscriber from all channels.
// Returns immediately if the Listener is stopped.
func (l *Listener) UnsubscribeAll(subCh chan *sqltypes.Notification) {
	_ = "STUB: not implemented"
	return
}

// SubscribeCh registers an existing notification channel to receive notifications
// for the given PG channel name. Unlike Subscribe, this lets multiple PG channels
// feed into the same notification channel.
// Returns immediately if the Listener is stopped.
func (l *Listener) SubscribeCh(channel string, notifCh chan *sqltypes.Notification) {
	_ = "STUB: not implemented"
	return
}

// readerMessage carries messages from the reader goroutine to the event loop.
type readerMessage struct {
	notification *sqltypes.Notification // non-nil for NotificationResponse
	cmdDone      bool                   // true for ReadyForQuery (command completed)
	err          error                  // non-nil for read errors
}

// run is the main event loop. It manages the PG connection, reads notifications,
// and processes subscribe/unsubscribe requests.
//
// Split read/write design:
//   - Reader goroutine: reads ALL messages from the socket (notifications, command
//     completions, errors). Sends parsed results to readerCh.
//   - Event loop (this function): processes subscribe/unsubscribe requests by writing
//     LISTEN/UNLISTEN commands directly on the live connection, then waits for the
//     reader to signal command completion via cmdDone.
//
// This eliminates the need to reconnect when subscribing/unsubscribing, preventing
// notification loss for existing channels.
func (l *Listener) run(ctx context.Context) { _ = "STUB: not implemented"; return }

// Channel refcounts: channel name -> number of subscribers.

// Subscribers: channel name -> list of subscriber channels.

// All subscriber channels (for broadcast warnings).

// receives from reader goroutine
// tracks when connection was lost for gap duration metric

// releaseConn releases the reserved connection. The readLoop goroutine
// will get a read error on the closed socket and exit.
// Setting readerCh = nil prevents the event loop from reading stale errors
// from the old readLoop after the connection is released.

// Track disconnect time for reconnect gap duration metric.
// Only set on error-path releases (not idle cleanup).

// connect establishes a new PG connection via the pool manager and starts the
// reader goroutine. Used for initial connection and error recovery (reconnect).
// On reconnect, re-LISTENs all active channels.

// never timeout

// Record reconnect gap duration if this was a reconnect (not initial connect).

// Re-LISTEN all active channels via simple query protocol.
// On initial connect this is a no-op (channels map is empty).

// Start reader goroutine with split read/write pattern.

// Connection is acquired lazily on first subscribe (not eagerly on startup)
// to avoid holding an idle reserved connection in the pool.

// pendingCmds tracks the number of LISTEN/UNLISTEN commands sent
// that haven't received a ReadyForQuery yet. This must be a counter
// (not a boolean) because reqUnsubscribeAll can send multiple
// UNLISTEN commands in sequence.

// Close all subscriber channels.

// First subscriber for this channel — issue LISTEN on the live connection.

// Connection is broken, schedule reconnect.

// No connection — acquire one. connect() will re-LISTEN
// all channels in the map, including the one just added.

// Issue UNLISTEN for channels that hit refcount 0.

// If we sent commands, drain ReadyForQuery signals before
// closing req.done. This ensures all LISTEN/UNLISTEN commands
// have completed on PG before we return to the caller.

// drainCommands may have released the connection on error.
// Nil readerCh to avoid reading stale errors from the old readLoop.

// Connection was released during command send — pending commands
// are lost with the connection, reset the counter.

// Release the connection when no channels are subscribed.
// ReleaseError taints the connection so the pool destroys it (closing the
// socket), which stops the readLoop goroutine. We can't use ReleaseRollback
// because the readLoop is still active on the socket — returning an active
// socket to the pool would corrupt data.
// On next subscribe, connect() will acquire a new reserved connection.

// Only reconnect if we have channels to listen on.

// drainCommands reads from readerCh until all pending command completions
// (ReadyForQuery) have been received. This is called after sending
// LISTEN/UNLISTEN to ensure all commands have completed on PG before returning
// to the caller. Notifications received during this drain are delivered normally.
//
// pendingCmds is decremented for each ReadyForQuery received; the function
// returns when it reaches 0. This correctly handles reqUnsubscribeAll which
// may send multiple UNLISTEN commands in sequence.
func (l *Listener) drainCommands(
	ctx context.Context,
	readerCh chan readerMessage,
	subscribers map[string][]chan *sqltypes.Notification,
	conn **reserved.Conn,
	reconnectTimer *time.Timer,
	pendingCmds *int,
	disconnectedAt *time.Time,
) {
	_ = "STUB: not implemented"
	return
}

// readLoop is the reader goroutine. It reads all messages from the PG connection
// and classifies them into notifications, command completions, or errors.
//
// Message handling:
//   - NotificationResponse ('A'): parsed and sent as notification
//   - CommandComplete ('C'): acknowledged (LISTEN/UNLISTEN response), not forwarded
//   - ReadyForQuery ('Z'): signals command completion to the event loop
//   - ErrorResponse ('E'): treated as connection error
//   - ParameterStatus ('S'), NoticeResponse ('N'): ignored
func (l *Listener) readLoop(conn *reserved.Conn, ch chan<- readerMessage) {
	_ = "STUB: not implemented"
	return
}

// LISTEN/UNLISTEN acknowledgment — skip, wait for ReadyForQuery.

// Command completed (LISTEN/UNLISTEN finished).

// PG returned an error for our LISTEN/UNLISTEN command.
// Treat as connection-level failure and reconnect.

// Ignore.

// removeSubResult holds the outcome of removeSub for metrics recording.
type removeSubResult struct {
	needsUnlisten  bool // channel refcount hit 0 — caller should issue UNLISTEN
	subscriberGone bool // subscriber fully removed from allSubs (no remaining channels)
}

// removeSub removes a single subscriber from a channel.
func (l *Listener) removeSub(
	subscribers map[string][]chan *sqltypes.Notification,
	allSubs map[chan *sqltypes.Notification]struct{},
	channels map[string]int,
	channel string,
	subCh chan *sqltypes.Notification,
) removeSubResult {
	_ = "STUB: not implemented"
	return *new(removeSubResult)
}

// Only remove from allSubs if the subscriber has no remaining subscriptions.
