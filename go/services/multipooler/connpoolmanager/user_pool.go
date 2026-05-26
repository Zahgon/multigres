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

package connpoolmanager

import (
	"context"
	"log/slog"
	"sync"
	"sync/atomic"
	"time"

	"github.com/multigres/multigres/go/common/pgprotocol/client"
	"github.com/multigres/multigres/go/services/multipooler/connstate"
	"github.com/multigres/multigres/go/services/multipooler/pools/admin"
	"github.com/multigres/multigres/go/services/multipooler/pools/connpool"
	"github.com/multigres/multigres/go/services/multipooler/pools/regular"
	"github.com/multigres/multigres/go/services/multipooler/pools/reserved"
)

// UserPool manages connection pools for a specific user.
// Each user gets their own RegularPool and ReservedPool that connect
// directly as that user (using trust/peer auth), eliminating the need
// for SET ROLE.
type UserPool struct {
	username     string
	regularPool  *regular.Pool
	reservedPool *reserved.Pool
	adminPool    *admin.Pool // Shared reference for kill operations
	logger       *slog.Logger

	// Demand tracking for rebalancer
	regularDemandTracker  *DemandTracker
	reservedDemandTracker *DemandTracker

	// Last activity timestamp (Unix nanos) for garbage collection
	lastActivity atomic.Int64

	mu     sync.Mutex
	closed bool
}

// UserPoolConfig holds configuration for creating a UserPool.
type UserPoolConfig struct {
	// ClientConfig is the PostgreSQL connection configuration.
	// The User field specifies the PostgreSQL user for this pool.
	ClientConfig *client.Config

	// AdminPool is the shared admin pool for kill operations.
	AdminPool *admin.Pool

	// RegularPoolConfig is the configuration for the regular connection pool.
	RegularPoolConfig *connpool.Config

	// ReservedPoolConfig is the configuration for the reserved pool's underlying connection pool.
	ReservedPoolConfig *connpool.Config

	// ReservedInactivityTimeout is how long a reserved connection can be inactive (no client activity)
	// before being killed. This is typically more aggressive (e.g., 30s) than pool idle timeout.
	ReservedInactivityTimeout time.Duration

	// DemandWindow is how far back to consider when calculating peak demand.
	// Set to 0 to disable demand tracking.
	// Example: 30s means "allocate based on peak demand over the last 30 seconds"
	DemandWindow time.Duration

	// RebalanceInterval is how often the rebalancer runs.
	// Number of demand tracking buckets = DemandWindow / RebalanceInterval.
	RebalanceInterval time.Duration

	// Logger for pool operations.
	Logger *slog.Logger

	// OnBorrow is called after a regular connection is borrowed from any pool (optional).
	OnBorrow func()

	// OnRecycle is called after a regular connection is returned to any pool (optional).
	OnRecycle func()

	// OnReserve is called after a new reserved connection is created (optional).
	OnReserve func()

	// OnRelease is called after a reserved connection is released or killed (optional).
	OnRelease func()
}

// NewUserPool creates a new UserPool for the given user.
// The pool connects directly as the user using trust/peer authentication.
// Returns an error if demand tracker configuration is invalid.
func NewUserPool(ctx context.Context, config *UserPoolConfig) (*UserPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Wire OnBorrow/OnRecycle callbacks into both regular and reserved pool configs.

// Create regular pool for this user

// Create reserved pool for this user (it manages its own internal regular pool)
// InactivityTimeout kills reserved connections if the client that reserved them
// hasn't used them for this duration (typically aggressive, e.g., 30s).
// The ReservedPoolConfig.IdleTimeout is for the underlying pool (less aggressive, e.g., 5min).

// Create demand trackers for rebalancer (if demand tracking is enabled).
// We use PeakRequestedAndReset instead of Requested to capture burst demand that
// point-in-time sampling would miss (e.g., short-lived queries that complete between samples).

// Username returns the username for this pool.
func (p *UserPool) Username() string {
	_ = "STUB: not implemented"

	// touchActivity updates the last activity timestamp.
	// Called internally when a connection is acquired.
	return ""
}

func (p *UserPool) touchActivity() { _ = "STUB: not implemented"; return }

// LastActivity returns the last activity timestamp (Unix nanos).
// Used by the rebalancer for garbage collection.
func (p *UserPool) LastActivity() int64 { _ = "STUB: not implemented"; return 0 }

// RegularDemand returns the peak demand for regular connections over the sliding window.
// It also rotates the demand tracker to the next bucket, aging out old data.
// This should be called once per rebalance cycle. Returns 0 if demand tracking is disabled.
func (p *UserPool) RegularDemand() int64 { _ = "STUB: not implemented"; return 0 }

// ReservedDemand returns the peak demand for reserved connections over the sliding window.
// It also rotates the demand tracker to the next bucket, aging out old data.
// This should be called once per rebalance cycle. Returns 0 if demand tracking is disabled.
func (p *UserPool) ReservedDemand() int64 { _ = "STUB: not implemented"; return 0 }

// GetRegularConn acquires a regular connection from the pool.
// The connection is already authenticated as the pool's user.
func (p *UserPool) GetRegularConn(ctx context.Context) (regular.PooledConn, error) {
	_ = "STUB: not implemented"
	return *new(regular.PooledConn), nil
}

// GetRegularConnWithSettings acquires a regular connection with the given settings.
// The connection is already authenticated as the pool's user.
func (p *UserPool) GetRegularConnWithSettings(ctx context.Context, settings *connstate.Settings) (regular.PooledConn, error) {
	_ = "STUB: not implemented"
	return *new(regular.PooledConn), nil
}

// NewReservedConn creates a new reserved connection for transactions or portal operations.
// The connection is already authenticated as the pool's user. Optional
// ReservedConnOption values configure validate-with-retry behavior.
func (p *UserPool) NewReservedConn(ctx context.Context, settings *connstate.Settings, opts ...reserved.ReservedConnOption) (*reserved.Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewLogicalReplicationConn returns a Postgres connection opened with
// replication=database in startup parameters and tagged with
// ReasonLogicalReplication. The connection is checked out from this user's
// reserved pool and authenticates as this user — required because the
// replication=database startup parameter is rejected for roles without the
// REPLICATION attribute.
func (p *UserPool) NewLogicalReplicationConn(ctx context.Context) (*reserved.Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetReservedConn retrieves an existing reserved connection by ID.
// Returns nil, false if the connection is not found or has timed out.
func (p *UserPool) GetReservedConn(connID int64) (*reserved.Conn, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// CloseReservedConnections kills all active reserved connections.
// Used during graceful shutdown when the drain grace period has expired.
func (p *UserPool) CloseReservedConnections(ctx context.Context) int {
	_ = "STUB: not implemented"
	return 0
}

// Close closes both regular and reserved pools.
func (p *UserPool) Close() { _ = "STUB: not implemented"; return }

// Close reserved pool first (it has its own internal regular pool)

// Close regular pool

// Stats returns statistics for both pools.
func (p *UserPool) Stats() UserPoolStats { _ = "STUB: not implemented"; return *new(UserPoolStats) }

// SetCapacity updates the capacity of both regular and reserved pools.
// This is a non-blocking operation: capacity is set immediately, idle connections
// are closed aggressively, and any remaining over-capacity connections are closed
// when they are recycled back to the pool.
func (p *UserPool) SetCapacity(ctx context.Context, regularCap, reservedCap int64) error {
	_ = "STUB: not implemented"
	return nil
}

// Set regular pool capacity

// Set reserved pool capacity

// UserPoolStats holds statistics for a user's pools.
type UserPoolStats struct {
	Username       string
	Regular        connpool.PoolStats
	Reserved       reserved.PoolStats
	RegularDemand  int64 // Peak demand from tracker (0 if tracking not enabled)
	ReservedDemand int64 // Peak demand from tracker (0 if tracking not enabled)
	LastActivity   int64 // Unix nanos of last activity

	// Aggregate metrics across regular and reserved pools.
	WaitCount int64         // Total times a client had to wait for a connection
	WaitTime  time.Duration // Total time clients spent waiting for a connection
	GetCount  int64         // Total connections borrowed (Get() calls)
	Waiting   int           // Clients currently waiting for a connection
}
