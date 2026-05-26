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

package admin

import (
	"context"

	"github.com/multigres/multigres/go/common/pgprotocol/client"
	"github.com/multigres/multigres/go/services/multipooler/pools/connpool"
)

// PoolConfig holds configuration for the admin pool.
type PoolConfig struct {
	// ClientConfig is the PostgreSQL connection configuration.
	ClientConfig *client.Config

	// ConnPoolConfig is the connection pool configuration.
	ConnPoolConfig *connpool.Config
}

// PooledConn is an alias for a pooled admin connection.
type PooledConn = *connpool.Pooled[*Conn]

// Pool manages a pool of administrative connections using connpool.Pool.
// These connections are used for terminating/canceling other backend connections
// via pg_terminate_backend() and pg_cancel_backend().
//
// The pool is intentionally small (typically 2-5 connections) and always has
// connections available for kill operations.
type Pool struct {
	pool   *connpool.Pool[*Conn]
	config *PoolConfig
}

// NewPool creates a new admin connection pool.
// The context is used for background pool operations and OTel tracking.
// The pool must be opened with Open() before use.
func NewPool(ctx context.Context, config *PoolConfig) *Pool { _ = "STUB: not implemented"; return nil }

// Open opens the pool and starts background workers.
// Must be called before using the pool.
func (p *Pool) Open() { _ = "STUB: not implemented"; return }

// Get acquires an admin connection from the pool.
// The caller must call Recycle() on the returned PooledConn to return the connection.
func (p *Pool) Get(ctx context.Context) (PooledConn, error) {
	_ = "STUB: not implemented"
	return *

	// TerminateBackend terminates a backend process using pg_terminate_backend().
	// This is a convenience method that handles getting/returning connections.
	// Returns true if the backend was terminated, false if it was not found or
	// the caller lacks permission.
	new(PooledConn), nil
}

func (p *Pool) TerminateBackend(ctx context.Context, processID uint32) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// CancelBackend cancels the current query on a backend process using pg_cancel_backend().
// This is a convenience method that handles getting/returning connections.
// Returns true if the signal was sent, false if the backend was not found or
// the caller lacks permission.
func (p *Pool) CancelBackend(ctx context.Context, processID uint32) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Close closes all connections in the pool.
func (p *Pool) Close() {
	_ = "STUB: not implemented"

	// Stats returns current pool statistics.
	return
}

func (p *Pool) Stats() connpool.PoolStats {
	_ = "STUB: not implemented"
	return *new(connpool.PoolStats)
}
