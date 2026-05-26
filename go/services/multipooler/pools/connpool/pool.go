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

package connpool

import (
	"context"
	"errors"
	"log/slog"
	"sync"
	"sync/atomic"
	"time"

	"github.com/multigres/multigres/go/services/multipooler/connstate"
)

var (
	// ErrTimeout is returned if a connection get times out.
	ErrTimeout = errors.New("connection pool timed out")

	// ErrCtxTimeout is returned if a ctx is already expired by the time the connection pool is used
	ErrCtxTimeout = errors.New("connection pool context already expired")

	// ErrPoolClosed is returned when trying to get a connection from a closed pool
	ErrPoolClosed = errors.New("connection pool is closed")

	// PoolCloseTimeout is how long to wait for all connections to be returned to the pool during close
	PoolCloseTimeout = 10 * time.Second
)

// Metrics holds pool metrics for monitoring.
type Metrics struct {
	maxLifetimeClosed atomic.Int64
	getCount          atomic.Int64
	getWithStateCount atomic.Int64
	waitCount         atomic.Int64
	waitTime          atomic.Int64
	idleClosed        atomic.Int64
	diffState         atomic.Int64
	resetState        atomic.Int64
}

func (m *Metrics) MaxLifetimeClosed() int64 { _ = "STUB: not implemented"; return 0 }
func (m *Metrics) GetCount() int64          { _ = "STUB: not implemented"; return 0 }
func (m *Metrics) GetStateCount() int64     { _ = "STUB: not implemented"; return 0 }
func (m *Metrics) WaitCount() int64         { _ = "STUB: not implemented"; return 0 }
func (m *Metrics) WaitTime() time.Duration  { _ = "STUB: not implemented"; return *new(time.Duration) }
func (m *Metrics) IdleClosed() int64        { _ = "STUB: not implemented"; return 0 }
func (m *Metrics) DiffStateCount() int64    { _ = "STUB: not implemented"; return 0 }
func (m *Metrics) ResetStateCount() int64   { _ = "STUB: not implemented"; return 0 }

// Connector is a function that creates a new connection.
// ctx is used for the dial/startup operations.
// poolCtx is the pool's lifecycle context, used to tie the connection's lifetime to the pool.
type Connector[C Connection] func(ctx context.Context, poolCtx context.Context) (C, error)

// RefreshCheck is a callback to check whether the pool needs to be refreshed.
type RefreshCheck func() (bool, error)

// Config holds configuration for the connection pool.
type Config struct {
	// Name is the pool name for logging and metrics (defaults to "" if not set).
	// The name is used in metrics to distinguish between different pools.
	Name string

	Capacity        int64
	MaxIdleCount    int64
	IdleTimeout     time.Duration
	MaxLifetime     time.Duration
	RefreshInterval time.Duration
	LogWait         func(time.Time)
	Logger          *slog.Logger

	// ConnectTimeout bounds background connection attempts (replacement connections
	// created during put, idle cleanup, and capacity increases). When non-zero,
	// a derived context with this timeout is used instead of the unbounded pool.ctx.
	// This prevents a hung dial from permanently consuming an active slot.
	ConnectTimeout time.Duration

	// OTel metrics instruments (optional, noop if not set).
	// These are shared across all pools and created by the owner (e.g., connpoolmanager).
	ConnectionCount ConnectionCount

	// OnBorrow is called after a connection is borrowed from the pool (optional).
	OnBorrow func()

	// OnRecycle is called after a connection is returned to the pool (optional).
	OnRecycle func()
}

// stackMask is the number of connection state stacks minus one;
// the number of stacks must always be a power of two
const stackMask = 7

// Pool is a connection pool for generic connections.
// It uses mutex-protected stacks for connection storage and a waitlist for
// blocking when the pool is at capacity.
type Pool[C Connection] struct {
	// clean is a connections stack for connections with no state applied
	clean connStack[C]
	// states are N connection stacks for connections with a state applied
	// connections are distributed between stacks based on their state hash bucket
	states [stackMask + 1]connStack[C]
	// freshStatesStack is the index in states to the last stack when a connection
	// was pushed, or -1 if no connection with a state has been opened in this pool
	freshStatesStack atomic.Int64
	// wait is the list of clients waiting for a connection to be returned to the pool
	wait waitlist[C]

	// borrowed is the number of connections that the pool has given out to clients
	// and that haven't been returned yet
	borrowed atomic.Int64
	// requested is the number of pending connection requests (Get calls in progress).
	// This includes both waiting requests and borrowed connections.
	// Used for demand tracking: incremented on Get() start, decremented on Get() fail or Recycle().
	requested atomic.Int64
	// peakRequested tracks the highest requested value since last reset.
	// Used for demand tracking: captures burst demand that point-in-time sampling might miss.
	peakRequested atomic.Int64
	// active is the number of connections that the pool has opened; this includes connections
	// in the pool and borrowed by clients
	active atomic.Int64
	// capacity is the maximum number of connections that this pool can open
	capacity atomic.Int64
	// idleCount is the maximum idle connections in the pool
	idleCount atomic.Int64

	// workers is a waitgroup for all the currently running worker goroutines
	workers    sync.WaitGroup
	close      atomic.Pointer[chan struct{}]
	capacityMu sync.Mutex

	// ctx is the context used for background pool operations
	ctx context.Context

	config struct {
		// connect is the callback to create a new connection for the pool
		connect Connector[C]
		// refresh is the callback to check whether the pool needs to be refreshed
		refresh RefreshCheck
		// maxCapacity is the maximum value to which capacity can be set; when the pool
		// is re-opened, it defaults to this capacity
		maxCapacity int64
		// maxIdleCount is the maximum idle connections in the pool
		maxIdleCount int64
		// maxLifetime is the maximum time a connection can be open
		maxLifetime atomic.Int64
		// idleTimeout is the maximum time a connection can remain idle
		idleTimeout atomic.Int64
		// refreshInterval is how often to call the refresh check
		refreshInterval atomic.Int64
		// connectTimeout bounds background connection attempts to prevent pool starvation
		connectTimeout time.Duration
		// logWait is called every time a client must block waiting for a connection
		logWait func(time.Time)
		// onBorrow is called after a connection is borrowed from the pool
		onBorrow func()
		// onRecycle is called after a connection is returned to the pool
		onRecycle func()
	}

	Metrics Metrics
	Name    string
	logger  *slog.Logger

	// otelConnectionCount tracks connection state counts (idle/used).
	// Optional, noop if not set. Provided via Config.ConnectionCount.
	otelConnectionCount ConnectionCount
}

// NewPool creates a new connection pool with the given Config.
// The pool must be Pool.Open before it can start giving out connections.
// The context is used for background pool operations and OTel tracking.
func NewPool[C Connection](ctx context.Context, config *Config) *Pool[C] {
	_ = "STUB: not implemented"
	return nil
}

// Set up OTel idle tracking callbacks on all idle stacks.

func (pool *Pool[C]) runWorker(close <-chan struct{}, interval time.Duration, worker func(now time.Time) bool) {
	_ = "STUB: not implemented"
	return
}

func (pool *Pool[C]) open() { _ = "STUB: not implemented"; return }

// already open

// The expire worker takes care of removing from the waiter list any clients whose
// context has been cancelled.

// Do not allow connections to starve; if there's waiters in the queue
// and connections in the stack, it means we could be starving them.
// Try getting out a connection and handing it over directly

// The idle worker takes care of closing connections that have been idle too long

// The refresh worker periodically checks the refresh callback in this pool
// to decide whether all the connections in the pool need to be cycled

// Open starts the background workers that manage the pool and gets it ready
// to start serving out connections.
func (pool *Pool[C]) Open(connect Connector[C], refresh RefreshCheck) *Pool[C] {
	_ = "STUB: not implemented"
	return nil
}

// Close shuts down the pool. No connections will be returned from Pool.Get after calling this,
// but calling Pool.Put is still allowed. This function will not return until all of the pool's
// connections have been returned or the default PoolCloseTimeout has elapsed.
func (pool *Pool[C]) Close() { _ = "STUB: not implemented"; return }

// CloseWithContext behaves like Close but allows passing in a Context to time out the
// pool closing operation.
func (pool *Pool[C]) CloseWithContext(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// already closed

// Set capacity to 0 and close all idle connections immediately

// Wait for borrowed connections to be returned (with timeout).
// Unlike SetCapacity (which is non-blocking for rebalancer use),
// Close should wait for graceful shutdown.

// waitForDrain waits for all active connections to be closed.
// This is used during graceful shutdown to wait for borrowed connections.
func (pool *Pool[C]) waitForDrain(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (pool *Pool[C]) reopen() { _ = "STUB: not implemented"; return }

// Set capacity to 0 to close all connections, then wait for drain

// Restore original capacity

// IsOpen returns whether the pool is open.
func (pool *Pool[C]) IsOpen() bool { _ = "STUB: not implemented"; return false }

// Capacity returns the maximum amount of connections that this pool can maintain open.
func (pool *Pool[C]) Capacity() int64 { _ = "STUB: not implemented"; return 0 }

// MaxCapacity returns the maximum value to which Capacity can be set via Pool.SetCapacity.
func (pool *Pool[C]) MaxCapacity() int64 { _ = "STUB: not implemented"; return 0 }

func (pool *Pool[C]) setIdleCount() { _ = "STUB: not implemented"; return }

// InUse returns the number of connections that the pool has lent out to clients and that
// haven't been returned yet.
func (pool *Pool[C]) InUse() int64 { _ = "STUB: not implemented"; return 0 }

// Available returns the number of connections that the pool can immediately lend out to
// clients without blocking.
func (pool *Pool[C]) Available() int64 { _ = "STUB: not implemented"; return 0 }

// Active returns the number of connections that the pool has currently open.
func (pool *Pool[C]) Active() int64 { _ = "STUB: not implemented"; return 0 }

func (pool *Pool[D]) IdleTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (pool *Pool[C]) SetIdleTimeout(duration time.Duration) { _ = "STUB: not implemented"; return }

func (pool *Pool[D]) IdleCount() int64 { _ = "STUB: not implemented"; return 0 }

func (pool *Pool[D]) RefreshInterval() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (pool *Pool[C]) recordWait(start time.Time) { _ = "STUB: not implemented"; return }

// Get returns a connection from the pool with no state applied.
// If there are no connections in the pool to be returned, Get blocks until one
// is returned, or until the given ctx is cancelled.
// The connection must be returned to the pool once it's not needed by calling Pooled.Recycle.
func (pool *Pool[C]) Get(ctx context.Context) (*Pooled[C], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetWithSettings returns a connection from the pool with the given settings applied.
// If there are no connections in the pool to be returned, Get blocks until one
// is returned, or until the given ctx is cancelled.
// The connection must be returned to the pool once it's not needed by calling Pooled.Recycle.
func (pool *Pool[C]) GetWithSettings(ctx context.Context, settings *connstate.Settings) (*Pooled[C], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// connectionCtx returns a bounded context for connection operations (dial + startup).
// When connectTimeout is configured, it returns a context with that timeout derived
// from the given ctx. When zero, it returns ctx unchanged (backward compat).
func (pool *Pool[C]) connectionCtx(ctx context.Context) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

// put returns a connection to the pool. This is a private API.
// Return connections to the pool by calling Pooled.Recycle.
func (pool *Pool[C]) put(conn *Pooled[C]) { _ = "STUB: not implemented"; return }

// Track demand: decrement on return

func (pool *Pool[C]) tryReturnConn(conn *Pooled[C]) bool {
	_ = "STUB: not implemented"
	// If we're over capacity, close the connection.
	// This enables non-blocking SetCapacity - excess connections are closed on recycle.
	return false
}

// Direct handoff to waiter: used→used, waiter will do otel used +1

// Connection goes to idle stack

func (pool *Pool[C]) pop(stack *connStack[C]) *Pooled[C] {
	_ = "STUB: not implemented"
	// retry-loop: pop a connection from the stack and atomically check whether
	// its timeout has elapsed. If the timeout has elapsed, the borrow will fail,
	// which means that a background worker has already marked this connection
	// as stale and is in the process of shutting it down. If we successfully mark
	// the timeout as borrowed, we know that background workers will not be able
	// to expire this connection (even if it's still visible to them), so it's
	// safe to return it
	return nil
}

func (pool *Pool[C]) tryReturnAnyConn() bool { _ = "STUB: not implemented"; return false }

// closeOnOverCapacity closes a connection if the number of active connections exceeds capacity.
// This enables non-blocking SetCapacity: capacity is set immediately, and excess connections
// are closed as they are recycled. Returns true if the connection was closed.
func (pool *Pool[C]) closeOnOverCapacity(conn *Pooled[C]) bool {
	_ = "STUB: not implemented"
	return false
}

// closeOnIdleLimitReached closes a connection if the number of idle connections (active - inuse) in the pool
// exceeds the idleCount limit. It returns true if the connection is closed, false otherwise.
func (pool *Pool[C]) closeOnIdleLimitReached(conn *Pooled[C]) bool {
	_ = "STUB: not implemented"
	return false
}

func (pool *Pool[D]) extendedMaxLifetime() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (pool *Pool[C]) connReopen(ctx context.Context, dbconn *Pooled[C], now time.Duration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (pool *Pool[C]) connNew(ctx context.Context) (*Pooled[C], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pool *Pool[C]) getFromSettingsStack(settings *connstate.Settings) *Pooled[C] {
	_ = "STUB: not implemented"
	return nil
}

func (pool *Pool[C]) closedConn() { _ = "STUB: not implemented"; return }

func (pool *Pool[C]) getNew(ctx context.Context) (*Pooled[C], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get returns a pooled connection with no settings applied.
func (pool *Pool[C]) get(ctx context.Context) (*Pooled[C], error) {
	_ = "STUB: not implemented"
	return nil,

		// Track demand: increment at start, decrement on error (success decrements in put on Recycle)
		nil
}

// Update peak demand for accurate demand tracking (captures bursts that sampling might miss)

// best case: if there's a connection in the clean stack, return it right away

// check if we have enough capacity to open a brand-new connection to return

// if we don't have capacity, try popping a connection from any of the settings stacks

// if there are no connections in the settings stacks and we've lent out connections
// to other clients, wait until one of the connections is returned

// no connections available and no connections to wait for (pool is closed)

// if the connection we've acquired has settings applied, we must reset them before returning

// getWithSettings returns a connection from the pool with the given settings applied.
func (pool *Pool[C]) getWithSettings(ctx context.Context, settings *connstate.Settings) (*Pooled[C], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Track demand: increment at start, decrement on error (success decrements in put on Recycle)

// Update peak demand for accurate demand tracking (captures bursts that sampling might miss)

// best case: check if there's a connection in the settings stack where our settings belongs

// if there's no connection with our settings, try popping a clean connection

// otherwise try opening a brand new connection and we'll apply the settings to it

// try on the _other_ settings stacks, even if we have to reset the settings for the returned
// connection

// no connections anywhere in the pool; if we've lent out connections to other clients
// wait for one of them

// no connections available and no connections to wait for (pool is closed)

// ensure that the settings applied to the connection matches the one we want

// if there's other settings applied, reset them before applying our settings

// apply our settings now; if we can't we assume that the conn is broken
// and close it without returning to the pool

// SetCapacity changes the capacity (number of open connections) on the pool.
// This is a non-blocking operation: capacity is set immediately, and idle
// connections are closed aggressively. Any remaining over-capacity connections
// will be closed when they are recycled back to the pool.
//
// This design ensures the rebalancer is never blocked waiting for borrowed
// connections to be returned.
func (pool *Pool[C]) SetCapacity(_ context.Context, newcap int64) error {
	_ = "STUB: not implemented"
	return nil
}

// setCapacity is the internal implementation for SetCapacity; it must be called
// with pool.capacityMu being held.
func (pool *Pool[C]) setCapacity(newcap int64) error { _ = "STUB: not implemented"; return nil }

// Update the idle count to match the new capacity

// Capacity increased: proactively create connections for any waiters.
// This ensures waiters don't have to wait for existing connections to be recycled.

// Capacity decreased: close idle connections to get closer to new capacity.
// Don't wait for borrowed connections - they will be closed on recycle
// via closeOnOverCapacity() in tryReturnConn().

// Try closing from connections which are currently idle in the stacks

// No idle connections available to close.
// Remaining over-capacity connections will be closed when recycled.

// satisfyWaitersOnCapacityIncrease creates new connections for waiting clients
// when capacity has been increased. This is called from setCapacity.
func (pool *Pool[C]) satisfyWaitersOnCapacityIncrease() {
	_ = "STUB: not implemented"
	// Create connections for waiters while we have capacity and waiters
	return
}

// Connection creation failed, stop trying

// No capacity available (active >= capacity), stop

// Try to hand the connection to a waiter

// No more waiters, push connection to idle stack

func (pool *Pool[C]) closeIdleResources(now time.Time) { _ = "STUB: not implemented"; return }

// Do a read-only best effort iteration of all the connections in this
// stack and atomically attempt to mark them as expired.
// Any connections that are marked as expired are _not_ removed from
// the stack; it's generally unsafe to remove nodes from the stack
// besides the head. When clients pop from the stack, they'll immediately
// notice the expired connection and ignore it.
// see: timestamp.expired

// continue iteration

// Create replacement connections AFTER ForEach releases the stack
// mutex. Calling getNew/tryReturnConn inside ForEach would deadlock:
// tryReturnConn may Push to the same stack whose mutex ForEach holds,
// and sync.Mutex is not reentrant.

// Requested returns the current demand (pending connection requests + borrowed connections).
// This is used for demand tracking: it represents how many connections would be needed
// if all current requests were served immediately.
func (pool *Pool[C]) Requested() int64 { _ = "STUB: not implemented"; return 0 }

// PeakRequestedAndReset returns the peak demand since the last reset and resets the peak.
// This captures burst demand that point-in-time sampling might miss. For accurate demand
// tracking, call this method periodically to get the peak demand over an interval.
func (pool *Pool[C]) PeakRequestedAndReset() int64 { _ = "STUB: not implemented"; return 0 }

// Waiting returns the number of clients currently waiting for a connection.
func (pool *Pool[C]) Waiting() int { _ = "STUB: not implemented"; return 0 }

// Stats returns pool statistics.
func (pool *Pool[C]) Stats() PoolStats { _ = "STUB: not implemented"; return *new(PoolStats) }

// PoolStats contains pool statistics.
type PoolStats struct {
	Active    int64 // Total connections
	Borrowed  int64 // Connections borrowed by clients
	Idle      int64 // Connections available in pool
	Capacity  int64 // Maximum connections
	Available int64 // Connections available for immediate use
	Requested int64 // Pending requests + borrowed (demand)
	Waiting   int   // Clients waiting for a connection
}
