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

// timeoutThread is a single goroutine that monitors the head of the global
// queue and evicts entries whose window deadline has passed.
type timeoutThread struct {
	buf      *Buffer
	notifyCh chan struct{} // Signaled when queue changes (new enqueue or eviction)
	stopCh   chan struct{} // Closed to stop the goroutine
}

func newTimeoutThread(buf *Buffer) *timeoutThread { _ = "STUB: not implemented"; return nil }

func (tt *timeoutThread) start() { _ = "STUB: not implemented"; return }

func (tt *timeoutThread) stop() {
	_ = "STUB: not implemented"

	// notify signals the timeout thread to re-check the queue head.
	// Non-blocking: if a notification is already pending, this is a no-op.
	return
}

func (tt *timeoutThread) notify() { _ = "STUB: not implemented"; return }

func (tt *timeoutThread) run() { _ = "STUB: not implemented"; return }

// Check the head of the queue.

// Queue is empty — wait for a notification or stop.

// Calculate time until the head entry's deadline.

// Deadline already passed — evict immediately.

// Wait for the deadline, a notification, or stop.

// Reset the timer. We drain it first to avoid races.

// Head entry's deadline reached — evict it.

// Queue changed — re-check head.

// evictHead removes the first entry from the queue if it's past its deadline.
func (tt *timeoutThread) evictHead() { _ = "STUB: not implemented"; return }

// Not yet expired — the original head may have been removed by
// context cancellation or drain, leaving a newer entry at head.

// Decrement before close so waiters never see a stale gauge.
