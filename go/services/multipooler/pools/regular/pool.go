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

package regular

import (
	"context"
	"time"

	"github.com/multigres/multigres/go/common/pgprotocol/client"
	"github.com/multigres/multigres/go/services/multipooler/connstate"
	"github.com/multigres/multigres/go/services/multipooler/pools/admin"
	"github.com/multigres/multigres/go/services/multipooler/pools/connpool"
)

// PooledConn is an alias for a pooled regular connection.
type PooledConn = *connpool.Pooled[*Conn]

// PoolConfig holds configuration for the regular pool.
type PoolConfig struct {
	// ClientConfig is the PostgreSQL connection configuration.
	ClientConfig *client.Config

	// ConnPoolConfig is the connection pool configuration.
	// Use ConnPoolConfig.Name to set a custom pool name for metrics (defaults to "unnamed").
	ConnPoolConfig *connpool.Config

	// AdminPool is used for kill operations on connections.
	AdminPool *admin.Pool
}

// Pool manages regular connections for query execution.
// It wraps connpool.Pool[*Conn] for convenience.
type Pool struct {
	pool   *connpool.Pool[*Conn]
	config *PoolConfig
}

// NewPool creates a new regular connection pool.
// The context is used for background pool operations and OTel tracking.
// The pool must be opened with Open() before use.
func NewPool(ctx context.Context, config *PoolConfig) *Pool {
	_ = "STUB: not implemented"
	// Set default name if not provided in ConnPoolConfig
	return nil
}

// Open opens the pool and starts background workers.
// Must be called before using the pool.
func (p *Pool) Open() { _ = "STUB: not implemented"; return }

// Get returns a connection from the pool.
// The connection is already authenticated as the pool's configured user.
func (p *Pool) Get(ctx context.Context) (PooledConn, error) {
	_ = "STUB: not implemented"
	return *

	// GetWithSettings returns a connection from the pool with the given settings applied.
	// The connection is already authenticated as the pool's configured user.
	//
	// When settings are non-empty, the underlying connpool issues SET commands
	// during acquisition (connpool/pool.go ApplySettings). If that first write
	// hits a stale socket — a connection PostgreSQL closed silently while it
	// sat idle in the pool — the connpool closes the conn and surfaces a
	// connection-class error. Without retry here, callers see that error
	// transparently even though a fresh socket would succeed. We retry with
	// the same regime as the reserved pool (constants.MaxConnPoolRetryAttempts
	// attempts, constants.ConnPoolRetryBackoff between them); each attempt
	// dials a replacement conn since the previous one was closed.
	//
	// Non-connection errors (timeout, malformed SETs, pool closed) propagate
	// unchanged. The settings==nil/empty fast path is unaffected: that case
	// never runs ApplySettings, so there is no acquisition-time stale-socket
	// hazard to retry.
	new(PooledConn), nil
}

func (p *Pool) GetWithSettings(ctx context.Context, settings *connstate.Settings) (PooledConn, error) {
	_ = "STUB: not implemented"
	return *new(PooledConn), nil
}

// Close closes all connections in the pool.
func (p *Pool) Close() {
	_ = "STUB: not implemented"

	// Stats returns current pool statistics.
	return
}

func (p *Pool) Stats() connpool.PoolStats {
	_ = "STUB: not implemented"
	return *

	// SetCapacity changes the pool's maximum capacity.
	// If reducing capacity, may block waiting for borrowed connections to return.
	new(connpool.PoolStats)
}

func (p *Pool) SetCapacity(ctx context.Context, newcap int64) error {
	_ = "STUB: not implemented"
	return nil
}

// InnerPool returns the underlying connpool.Pool.
// Use with caution - prefer the wrapped methods.
func (p *Pool) InnerPool() *connpool.Pool[*Conn] {
	_ = "STUB: not implemented"

	// Requested returns the number of currently requested connections (borrowed + waiters).
	// Used for demand tracking in the rebalancer.
	return nil
}

func (p *Pool) Requested() int64 { _ = "STUB: not implemented"; return 0 }

// PeakRequestedAndReset returns the peak demand since the last reset and resets the peak.
// This captures burst demand that point-in-time sampling might miss.
func (p *Pool) PeakRequestedAndReset() int64 { _ = "STUB: not implemented"; return 0 }

// WaitCount returns the total number of times a client had to wait for a connection.
func (p *Pool) WaitCount() int64 { _ = "STUB: not implemented"; return 0 }

// WaitTime returns the total time clients spent waiting for a connection.
func (p *Pool) WaitTime() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

// GetCount returns the total number of Get() calls (connections borrowed).
func (p *Pool) GetCount() int64 { _ = "STUB: not implemented"; return 0 }
