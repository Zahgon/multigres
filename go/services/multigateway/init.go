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

// multigateway is the top-level proxy that masquerades as a PostgreSQL server,
// handling client connections and routing queries to multipooler instances.

// Package multigateway provides multigateway functionality.
package multigateway

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/multigres/multigres/go/common/pgprotocol/server"
	"github.com/multigres/multigres/go/common/rpcclient"
	"github.com/multigres/multigres/go/common/servenv"
	"github.com/multigres/multigres/go/common/servenv/toporeg"
	"github.com/multigres/multigres/go/common/topoclient"
	"github.com/multigres/multigres/go/services/multigateway/buffer"
	"github.com/multigres/multigres/go/services/multigateway/executor"
	"github.com/multigres/multigres/go/services/multigateway/handler"
	"github.com/multigres/multigres/go/services/multigateway/handler/queryregistry"
	"github.com/multigres/multigres/go/services/multigateway/poolergateway"
	"github.com/multigres/multigres/go/services/multigateway/scatterconn"
	"github.com/multigres/multigres/go/tools/viperutil"
)

type MultiGateway struct {
	cell viperutil.Value[string]
	// serviceID string
	serviceID viperutil.Value[string]
	// pgPort is the PostgreSQL protocol listen port
	pgPort viperutil.Value[int]
	// pgBindAddress is the address to bind the PostgreSQL listener to
	pgBindAddress viperutil.Value[string]
	// pgTLSCertFile is the path to the TLS certificate file for PostgreSQL SSL connections.
	pgTLSCertFile viperutil.Value[string]
	// pgTLSKeyFile is the path to the TLS private key file for PostgreSQL SSL connections.
	pgTLSKeyFile viperutil.Value[string]
	// pgRequireSSL rejects plaintext client connections; requires cert + key.
	pgRequireSSL viperutil.Value[bool]
	// poolerDiscovery handles discovery of multipoolers across all cells
	poolerDiscovery *GlobalPoolerDiscovery
	// poolerGateway manages connections to poolers
	poolerGateway *poolergateway.PoolerGateway
	// grpcServer is the grpc server
	grpcServer *servenv.GrpcServer
	// pgListener is the PostgreSQL protocol listener
	pgListener *server.Listener
	// pgHandler is the PostgreSQL protocol handler
	pgHandler *handler.MultiGatewayHandler
	// pgReplicaPort is the optional port for replica-reads connections.
	// When set, a second listener accepts connections that are allowed to read from replicas.
	pgReplicaPort viperutil.Value[int]
	// pgReplicaListener is the optional replica-reads listener
	pgReplicaListener *server.Listener
	// pgReplicaLowLagMs is the preferred replication lag threshold (ms) for replicas.
	// Replicas at or below this lag are considered "healthy" and preferred.
	pgReplicaLowLagMs viperutil.Value[int]
	// pgReplicaHighLagToleranceMs is the absolute maximum replication lag (ms).
	// Replicas above this are never selected. 0 means no upper bound.
	pgReplicaHighLagToleranceMs viperutil.Value[int]
	// cancelManager handles cross-gateway query cancellation
	cancelManager *CancelManager
	// scatterConn coordinates query execution across poolers
	scatterConn *scatterconn.ScatterConn
	// executor handles query execution and routing
	executor *executor.Executor
	// buffer holds requests during PRIMARY failovers
	buffer *buffer.Buffer
	// bufferConfig holds buffer configuration
	bufferConfig *buffer.Config
	// statementTimeout is the default statement execution timeout
	statementTimeout viperutil.Value[time.Duration]
	// authenticationTimeout bounds the PG startup phase (SSL handshake,
	// StartupMessage, SCRAM exchange). Equivalent to PostgreSQL's
	// authentication_timeout GUC.
	authenticationTimeout viperutil.Value[time.Duration]
	// planCacheMemory is the maximum memory (bytes) for the plan cache (0 disables)
	planCacheMemory viperutil.Value[int]
	// queryMetricsMemory is the maximum memory (bytes) for per-query-shape metrics
	// tracking (0 disables fingerprint labeling and the registry RPCs).
	queryMetricsMemory viperutil.Value[int]
	// queryMetricsSQLMaxBytes is the maximum bytes of representative normalized
	// SQL stored per tracked fingerprint.
	queryMetricsSQLMaxBytes viperutil.Value[int]
	// queryLogSampleRate is the 1/N sampling rate for normal-path query logs.
	queryLogSampleRate viperutil.Value[uint64]
	// queryRegistry tracks per-fingerprint query statistics; shared across
	// primary and replica handlers so metrics aggregate to the same bucket.
	queryRegistry *queryregistry.Registry
	// senv is the serving environment
	senv *servenv.ServEnv
	// connConfig holds RPC client configuration (TLS, etc.) for multipooler connections
	connConfig *rpcclient.ConnConfig
	// topoConfig holds topology configuration
	topoConfig   *topoclient.TopoConfig
	ts           topoclient.Store
	tr           *toporeg.TopoReg
	serverStatus Status
	// shutdownCtx is cancelled during Shutdown to propagate cancellation
	// to all long-running goroutines (health streams, discovery, etc.)
	shutdownCtx    context.Context
	shutdownCancel context.CancelFunc
}

func NewMultiGateway() *MultiGateway { _ = "STUB: not implemented"; return nil }

// 4 MB

// 8 MB; 0 disables per-query tracking

// Executor returns the query executor for this multigateway.
func (mg *MultiGateway) Executor() *executor.Executor {
	_ = "STUB: not implemented"

	// ServEnv returns the serving environment for this multigateway.
	return nil
}

func (mg *MultiGateway) ServEnv() *servenv.ServEnv { _ = "STUB: not implemented"; return nil }

func (mg *MultiGateway) RegisterFlags(fs *pflag.FlagSet) { _ = "STUB: not implemented"; return }

// Init initializes the multigateway. If any services fail to start,
// or if some connections fail, it launches goroutines that retry
// until successful.
func (mg *MultiGateway) Init(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Resolve service ID early for telemetry resource attributes
	return nil
}

// This doesn't change

// Create a service-lifetime context cancelled on shutdown.

// Start pooler discovery (watches all cells)

// Build transport credentials for multipooler gRPC connections.

// Create LoadBalancer and register with discovery for real-time updates

// Create failover buffer if enabled.

// Stop buffering when the streaming health check detects a new primary.
// This is a direct signal from the pooler's health stream — more reliable
// and lower latency than topology-based propagation via etcd.

// Initialize PoolerGateway for managing pooler connections

// Initialize ScatterConn for query coordination

// Initialize the executor for query routing
// Pass ScatterConn as the IExecute implementation

// Initialize gateway-wide OTel metrics up front so the credential
// provider and listener can share the same sink. Failures here are
// non-fatal: the auth and TLS code paths tolerate a nil/noop recorder
// and we don't want metric init to block startup.

// Create the credential provider for SCRAM authentication and the
// replication-role gate. A single GetAuthCredentials RPC feeds both,
// so admitting a replication connection never costs two pooler hops.

// Build TLS config if cert and key files are provided.

// Build the full gateway record. All info (hostname, ports) is available
// after servenv.Init(). PidPrefix is assigned during registration below.

// Reuse existing PID prefix on re-registration.

// Register gateway in topo with a unique PID prefix for cross-gateway
// cancel routing. The register function assigns the prefix, registers the
// full record, and verifies no collision. On collision, RegisterSynchronous
// retries with jitter until two racing gateways converge on different prefixes.

// Reset for next retry.

// Construct the per-query-shape metrics registry. Shared across primary
// and replica handlers so a query hitting either listener aggregates to
// the same stats bucket.
// Start from DefaultConfig so SampleInterval / TrendWindowSamples are
// populated; only the operator-tunable size knobs come from flags.

// Create and start PostgreSQL protocol listener

// Wire LISTEN/NOTIFY notification manager.
// Uses a lazy client getter that resolves the primary pooler connection
// from the load balancer at subscribe time (after pooler discovery).

// Optionally create a second listener for replica-reads connections.

// Register client connection metrics. The gatewayMetrics instance was
// constructed earlier so the credential provider and listeners share it.

// Set up cross-gateway cancel request handling.
// The cancel manager routes to the correct listener based on the connection
// type (primary vs replica) carried in the cancel request / gRPC forward.

// Register gRPC services via OnRun because grpcServer.Server is only
// created in servenv.Run() (after Create()), which runs after Init().

// Start the PostgreSQL listener in a goroutine

// Start the replica listener if configured.

// The gateway is ready only when both conditions are met:
// 1. No init errors (topology registration succeeded)
// 2. At least one pooler has been discovered (can actually serve queries)

func (mg *MultiGateway) RunDefault() error { _ = "STUB: not implemented"; return nil }

func (mg *MultiGateway) CobraPreRunE(cmd *cobra.Command) error {
	_ = "STUB: not implemented"
	return nil
}

func (mg *MultiGateway) Shutdown() { _ = "STUB: not implemented"; return }

// Cancel the service-lifetime context first so health stream goroutines
// stop promptly, before we close the underlying gRPC connections.

// Stop PostgreSQL listener

// Stop replica PostgreSQL listener (if running)

// Close cancel manager's gRPC connections

// Close executor (plan cache cleanup)

// Close per-query-shape metrics registry.

// Stop failover buffer

// Close pooler gateway connections

// Stop pooler discovery

// findUnusedPrefix scans all cells for used PID prefixes and returns a random
// unused one. Randomization reduces the chance of two gateways starting
// simultaneously and picking the same prefix.
func (mg *MultiGateway) findUnusedPrefix(ctx context.Context) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Cell may not have gateways yet.

// Collect all unused prefixes and pick one at random.

// hasPrefixCollision checks if any other gateway in topo has the same PID prefix.
func (mg *MultiGateway) hasPrefixCollision(ctx context.Context, prefix uint32, ownIDStr string) bool {
	_ = "STUB: not implemented"
	return false
}

// buildPGTLSConfig validates TLS flag combinations and loads the certificate.
// Returns nil if neither cert nor key file is configured (plaintext mode).
func buildPGTLSConfig(certFile, keyFile string) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PG 17 ALPN forward compatibility
