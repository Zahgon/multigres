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
	"time"

	"go.opentelemetry.io/otel/metric"

	"github.com/multigres/multigres/go/services/multipooler/pools/connpool"
)

// Error type labels for mg.pooler.auth.credential_query.errors. Closed set,
// driven by the failure modes in grpcpoolerservice.GetAuthCredentials so
// the cardinality stays bounded. user_not_found / login_disabled /
// password_expired are policy outcomes of a successful pg_authid query
// (baseline traffic); pool_acquire_failed / db_error indicate operator
// issues. All five are tagged on the counter so the success rate equals
// duration_count - errors_total.
const (
	CredentialQueryErrorUserNotFound      = "user_not_found"
	CredentialQueryErrorLoginDisabled     = "login_disabled"
	CredentialQueryErrorPasswordExpired   = "password_expired"
	CredentialQueryErrorPoolAcquireFailed = "pool_acquire_failed"
	CredentialQueryErrorDB                = "db_error"
)

// Metrics holds OpenTelemetry metrics for connection pool management.
type Metrics struct {
	meter metric.Meter

	// registration is the handle returned by RegisterCallback. Stored so the
	// callback can be unregistered when the manager is closed.
	registration metric.Registration

	// regularConnCount tracks PostgreSQL connection states for regular pools
	regularConnCount connpool.ConnectionCount

	// reservedConnCount tracks PostgreSQL connection states for reserved pools
	reservedConnCount connpool.ConnectionCount

	// --- Observable gauges for PgBouncer-equivalent metrics ---

	// poolerUp reports whether the pooler is operational (1 = up, 0 = down).
	poolerUp metric.Int64ObservableGauge

	// poolCount is the number of user connection pools.
	poolCount metric.Int64ObservableGauge

	// userCount is the number of distinct users with pools.
	userCount metric.Int64ObservableGauge

	// databaseCount is the number of databases (always 1 per multipooler instance).
	databaseCount metric.Int64ObservableGauge

	// serverConnections is the number of server connections by state (active/idle).
	serverConnections metric.Int64ObservableGauge

	// clientWaitingConnections is the number of clients waiting for a server connection.
	clientWaitingConnections metric.Int64ObservableGauge

	// reservedActiveConnections is the number of active reserved connections (in-transaction).
	reservedActiveConnections metric.Int64ObservableGauge

	// configMaxServerConnections is the configured maximum server connections (global capacity).
	configMaxServerConnections metric.Int64ObservableGauge

	// poolCapacity is the per-user allocated pool capacity (regular + reserved).
	poolCapacity metric.Int64ObservableGauge

	// poolCurrentConnections is the per-user active server connections (regular + reserved).
	poolCurrentConnections metric.Int64ObservableGauge

	// --- Observable counters for cumulative pool traffic metrics ---

	// clientWaitTimeTotal is the cumulative time clients spent waiting for a server connection.
	clientWaitTimeTotal metric.Float64ObservableCounter

	// queriesPooledTotal is the total number of connections borrowed (Get() calls).
	queriesPooledTotal metric.Int64ObservableCounter

	// --- Auth-path metrics (mg.pooler.auth.*) ---

	// authCredQueryDuration histograms the wall-clock duration of the
	// admin-pool `SELECT rolpassword FROM pg_authid` lookup the gateway
	// triggers via GetAuthCredentials. Covers admin-pool acquire +
	// query + result decode end-to-end, so a tail rise here is the
	// first signal that the admin pool is contended.
	authCredQueryDuration metric.Float64Histogram

	// authCredQueryErrors counts the failure modes of the credential
	// query, labeled by error type (user_not_found / pool_acquire_failed
	// / db_error). user_not_found is the legitimate-traffic baseline;
	// the other two indicate operator issues.
	authCredQueryErrors metric.Int64Counter
}

// NewMetrics initializes OpenTelemetry metrics for connection pool management.
// Individual metrics that fail to initialize will use noop implementations and be included
// in the returned error. The returned Metrics instance is always usable (with noop fallbacks
// for failed metrics), and the error indicates which specific metrics failed to initialize.
func NewMetrics() (*Metrics, error) { _ = "STUB: not implemented"; return nil, nil }

// ConnectionCount for regular pools

// Use zero value (noop) on error

// ConnectionCount for reserved pools

// Use zero value (noop) on error

// Pooler health gauge

// Pool count gauge

// User count gauge

// Database count gauge

// Server connections gauge (with state attribute)

// Client waiting connections gauge

// Reserved active connections gauge

// Config max server connections gauge

// Per-user pool capacity gauge

// Per-user current connections gauge

// Client wait time counter

// Queries pooled counter

// Auth credential-query latency histogram.

// Auth credential-query error counter.

// RecordCredentialQuery records a credential-query latency observation and,
// when errorType is non-empty, bumps the corresponding error counter. Use
// the CredentialQueryError* constants for errorType so the label set stays
// closed. Safe to call on a nil receiver so the gRPC handler can stay
// unconditional even when metric init failed.
func (m *Metrics) RecordCredentialQuery(ctx context.Context, d time.Duration, errorType string) {
	_ = "STUB: not implemented"
	return
}

// RegisterManagerCallbacks registers OTel observable callbacks that read pool statistics.
// The callbacks are invoked by the OTel SDK during metric collection.
//
// Parameters:
//   - statsGetter: returns a snapshot of all pool statistics
//   - poolCountGetter: returns the number of active user pools
//   - globalCapacityGetter: returns the configured global connection capacity
//   - isClosedGetter: returns whether the manager is closed
func (m *Metrics) RegisterManagerCallbacks(
	statsGetter func() ManagerStats,
	poolCountGetter func() int,
	globalCapacityGetter func() int64,
	isClosedGetter func() bool,
) error {
	_ = "STUB: not implemented"
	// Collect all instruments that were successfully initialized.
	return nil
}

// Pooler health.

// Pool/user count.

// Database count (always 1 per multipooler instance).

// Read global capacity config.

// Aggregate stats across all user pools.

// Regular pool: server connections

// Reserved pool: underlying server connections

// Reserved active (clients in transactions)

// Aggregate cumulative metrics

// Per-user pool capacity and current connections.

// Also include admin pool connections.

// Close unregisters the observable callback so the OTel SDK stops invoking it.
// Safe to call multiple times and when no callback was registered.
func (m *Metrics) Close() error { _ = "STUB: not implemented"; return nil }

// RegularConnCount returns the ConnectionCount metric for regular pools.
func (m *Metrics) RegularConnCount() connpool.ConnectionCount {
	_ = "STUB: not implemented"
	return *new(connpool.ConnectionCount)
}

// ReservedConnCount returns the ConnectionCount metric for reserved pools.
func (m *Metrics) ReservedConnCount() connpool.ConnectionCount {
	_ = "STUB: not implemented"
	return *new(connpool.ConnectionCount)
}
