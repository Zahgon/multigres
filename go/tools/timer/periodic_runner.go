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

// Package timer provides PeriodicRunner for running callbacks at regular intervals.
package timer

import (
	"context"
	"errors"
	"sync"
	"time"
)

// ErrStopped is sent to notification channels when the runner is stopped
// before a cycle completes.
var ErrStopped = errors.New("runner stopped before cycle completed")

// state represents the lifecycle state of the PeriodicRunner.
type state int

const (
	stopped  state = iota // not running, can be started
	running               // actively running, callbacks may be scheduled
	stopping              // Stop() called, waiting for in-flight callbacks to complete
)

// PeriodicRunner runs a callback at regular intervals with lifecycle management.
//
// Key behaviors:
//   - Callback receives a context derived from the parent context
//   - Stop() cancels the context and waits for in-flight callbacks
//   - Next callback scheduled only after current completes (backpressure)
//   - Supports Start/Stop/Start cycles (reopening)
//
// Example usage:
//
//	runner := timer.NewPeriodicRunner(ctx, 1*time.Second)
//	runner.Start(func(ctx context.Context) {
//	    // periodic work here
//	})
//	// later...
//	runner.Stop() // waits for any in-flight callback
type PeriodicRunner struct {
	parentCtx context.Context
	interval  time.Duration

	mu       sync.Mutex
	cond     *sync.Cond      // for waiting on state transitions
	state    state           // current lifecycle state
	ctx      context.Context // child context, created on Start, cancelled on Stop
	cancel   context.CancelFunc
	timer    *time.Timer
	wg       sync.WaitGroup
	callback func(ctx context.Context)

	// Channels to notify after next cycle completes (for WithAfterNextFullCycle)
	cycleNotifications []chan error
}

// NewPeriodicRunner creates a PeriodicRunner with the given parent context and interval.
// The parent context is used to derive child contexts on each Start() call.
// Callers should typically pass a detached context (e.g., ctxutil.Detach()) to avoid
// the runner being cancelled when request contexts complete.
func NewPeriodicRunner(ctx context.Context, interval time.Duration) *PeriodicRunner {
	_ = "STUB: not implemented"
	return nil
}

// StartOptions configures how the PeriodicRunner starts.
type StartOptions struct {
	fastStart          bool       // If true, schedule callback immediately (0 delay) instead of waiting for first interval
	onStart            func()     // Optional callback invoked when actually starting (not when already running)
	afterNextFullCycle chan error // Optional channel to receive cycle completion status (nil = success, ErrStopped = stopped)
}

// StartOption is a functional option for StartWithOptions.
type StartOption func(*StartOptions)

// WithFastStart configures the runner to schedule the callback with minimal delay (0ms)
// when starting from a stopped state. If the runner is already running, this option
// has no effect. Combine with WithAfterNextFullCycle to wait for completion.
func WithFastStart() StartOption { _ = "STUB: not implemented"; return *new(StartOption) }

// WithOnStart configures a callback to be invoked when the runner actually starts
// (not when it's already running). The callback is invoked before any periodic
// callback can execute.
func WithOnStart(onStart func()) StartOption { _ = "STUB: not implemented"; return *new(StartOption) }

// WithAfterNextFullCycle configures a channel to receive a status after the next full
// (not yet started) cycle completes or the runner is stopped.
//
// The channel will receive:
//   - nil: if the cycle completed successfully
//   - ErrStopped: if the runner was stopped before the cycle completed
//
// Can be combined with WithFastStart: if stopped, fast start schedules immediately and
// the channel receives nil after that cycle completes; if already running, the channel
// receives a value after the next scheduled cycle completes.
//
// Example:
//
//	cycleDone := make(chan error, 1)
//	StartWithOptions(callback, WithFastStart(), WithAfterNextFullCycle(cycleDone))
//	select {
//	case err := <-cycleDone:
//	    if err == nil {
//	        // Cycle completed successfully
//	    } else {
//	        // Runner was stopped (err == ErrStopped)
//	    }
//	case <-ctx.Done():
//	    // Timeout or cancellation
//	}
func WithAfterNextFullCycle(ch chan error) StartOption {
	_ = "STUB: not implemented"
	return *new(StartOption)
}

// StartWithOptions begins running the callback at regular intervals with the given options.
// The callback receives a context that is cancelled when Stop() is called.
// Returns true if the runner was started, false if it was already running.
//
// Options:
//   - WithFastStart(): Schedule immediate execution (0 delay) instead of waiting for first interval
//   - WithOnStart(func()): Invoke callback when starting (not when already running)
//   - WithAfterNextFullCycle(chan): Close channel after next full cycle completes
func (r *PeriodicRunner) StartWithOptions(callback func(ctx context.Context), opts ...StartOption) bool {
	_ = "STUB: not implemented"
	return false
}

// Wait for any in-progress Stop() to complete

// Already running - register notification if requested

// Register notification if requested

// Schedule immediate execution (0 delay)

// Start begins running the callback at regular intervals.
// This is a convenience wrapper around StartWithOptions that preserves the old API.
// The callback receives a context that is cancelled when Stop() is called.
// If onStart is non-nil, it is called exactly once when actually starting
// (not when already running), before any callback can execute.
// Returns true if the runner was started, false if it was already running.
func (r *PeriodicRunner) Start(callback func(ctx context.Context), onStart func()) bool {
	_ = "STUB: not implemented"
	return false
}

// Stop cancels the context and waits for any in-flight callback to complete.
// After Stop returns, no more callbacks will run. Can be restarted with Start().
// Stop is idempotent - calling it when already stopped has no effect.
// Any pending cycle notification channels will receive ErrStopped.
func (r *PeriodicRunner) Stop() { _ = "STUB: not implemented"; return }

// Already stopped or another Stop() is in progress

// Transition to stopping - from this point, this thread owns the stopping->stopped transition

// Cancel context to unblock any in-flight callback

// Stop the timer to prevent new callbacks from being scheduled

// Notify any pending cycle watchers that runner is stopping

// Send ErrStopped to all waiting channels

// Wait for any in-flight callback to complete (outside lock to avoid deadlock)

// Transition to stopped and wake any waiting Start() calls

// Because this goroutine owns the transition from stopping -> stopped,
// nothing should have changed with that while we released the lock to Wait().

// Running returns true if the runner is currently running.
// This includes the stopping state (while waiting for in-flight callbacks).
func (r *PeriodicRunner) Running() bool { _ = "STUB: not implemented"; return false }

// UpdateInterval changes the interval for future executions.
// If the runner is currently running, the next scheduled execution will be cancelled
// and rescheduled with the new interval. If not running, only updates the interval field.
// Returns true if the interval was changed, false if it was already set to newInterval or runner is stopped.
func (r *PeriodicRunner) UpdateInterval(newInterval time.Duration) bool {
	_ = "STUB: not implemented"
	return false
}

// If running, reschedule with new interval

// scheduleNext schedules the next callback execution.
// Must be called while holding r.mu.
func (r *PeriodicRunner) scheduleNext() { _ = "STUB: not implemented"; return }

// execute runs the callback and schedules the next execution.
func (r *PeriodicRunner) execute() { _ = "STUB: not implemented"; return }

// Opportunistically stop if parent context is cancelled

// Track this execution so Stop() can wait for it to complete

// Capture callback, context, and notification channels while holding the lock

// Clear for next cycle

// Release lock during callback execution to avoid blocking Stop()

// Execute callback

// Notify all waiting channels that cycle completed successfully

// Re-acquire lock to schedule next execution

// Schedule next execution only after this one completes (backpressure)
