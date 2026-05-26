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

package buffer

import (
	"context"
	"log/slog"
	"sync"
	"time"

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
)

// bufferState represents the state of a per-shard buffer.
type bufferState int

const (
	stateIdle      bufferState = iota // Not buffering
	stateBuffering                    // Accepting requests into buffer
	stateDraining                     // Draining buffered requests via retry
)

func (s bufferState) String() string { _ = "STUB: not implemented"; return "" }

// shardBuffer manages the buffering state machine for a single shard.
// State transitions: IDLE -> BUFFERING -> DRAINING -> IDLE
type shardBuffer struct {
	buf      *Buffer
	shardKey *clustermetadatapb.ShardKey
	logger   *slog.Logger

	mu               sync.Mutex
	state            bufferState
	lastStart        time.Time   // When buffering last started
	lastEnd          time.Time   // When buffering last ended
	generation       uint64      // Incremented on each IDLE→BUFFERING transition
	maxDurationTimer *time.Timer // Fires when MaxFailoverDuration is exceeded
	drainWg          sync.WaitGroup
}

func newShardBuffer(buf *Buffer, key *clustermetadatapb.ShardKey) *shardBuffer {
	_ = "STUB: not implemented"
	return nil
}

// waitIfAlreadyBuffering joins an existing buffer if the shard is already
// BUFFERING, but does NOT transition IDLE -> BUFFERING. Used for proactive
// buffering before sending a query.
func (sb *shardBuffer) waitIfAlreadyBuffering(ctx context.Context) (RetryDoneFunc, error) {
	_ = "STUB: not implemented"
	return *new(RetryDoneFunc), nil
}

// waitForFailoverEnd either starts buffering (IDLE -> BUFFERING) or joins
// an existing buffer (already BUFFERING). Returns (nil, nil) if buffering
// is not applicable for this request.
func (sb *shardBuffer) waitForFailoverEnd(ctx context.Context) (RetryDoneFunc, error) {
	_ = "STUB: not implemented"
	// Fast path: if draining or idle with recent failover, skip.
	return *new(RetryDoneFunc), nil
}

// Already draining — the new PRIMARY is available. Signal the caller
// to retry immediately. A recursive retry loop is unlikely because
// the LoadBalancer updates its cached primary before invoking the
// onPrimaryServing callback that triggers StopBuffering, so the new
// PRIMARY is already routable by the time we reach here. It is
// bounded by context timeout in any case.

// Check timing guard: don't start buffering again too soon.

// Transition IDLE -> BUFFERING.

// Start max-duration timer. The generation is captured so that if
// the timer fires after this failover has already ended and a new
// one has started, the stale callback is ignored.

// Already buffering, just enqueue below.

// Enqueue into the global queue.

// waitOnEntry blocks until the entry's done channel is closed or the context is canceled.
func (sb *shardBuffer) waitOnEntry(ctx context.Context, e *entry) (RetryDoneFunc, error) {
	_ = "STUB: not implemented"
	return *new(RetryDoneFunc), nil
}

// Request context canceled (client disconnected, deadline, etc.).

// Signal retry completion so that if drainEntry already extracted
// this entry from the queue, it won't block forever on
// <-e.bufferCtx.Done(). If the entry was still in the queue,
// this is harmless (nobody is watching bufferCtx).

// Entry was evicted (buffer full, window timeout, max duration, shutdown).

// Failover ended successfully — caller should retry.

// stopBuffering transitions from BUFFERING to DRAINING and drains all entries.
// If gen is non-zero, the call is only valid for that specific generation
// (used by maxDurationTimer to avoid killing a subsequent failover's buffering).
// Pass gen=0 to stop unconditionally (used by external callers like StopBuffering).
func (sb *shardBuffer) stopBuffering(reason string, gen uint64) { _ = "STUB: not implemented"; return }

// Extract all entries for this shard from the global queue.

// Drain entries with configured concurrency. Each entry gets its own
// goroutine; the semaphore limits how many run in parallel.

// Acquire drain slot.

// Release drain slot.

// All entries drained, transition back to IDLE.

// drainEntry signals a single entry to retry and waits for its completion.
func (sb *shardBuffer) drainEntry(e *entry) {
	_ = "STUB: not implemented"
	// Decrement before close so waiters never see a stale gauge.
	return
}

// Signal the entry to retry by closing its done channel.

// Wait for the retry to complete (caller invokes RetryDoneFunc which
// calls bufferCancel).

// Release the semaphore slot.
