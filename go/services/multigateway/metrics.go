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

package multigateway

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/metric"

	"github.com/multigres/multigres/go/common/pgprotocol/server"
	"github.com/multigres/multigres/go/services/multigateway/auth"
)

// Compile-time checks that GatewayMetrics implements the recorder
// interfaces it is wired into. Catches drift in interface method
// signatures at build time rather than at the (out-of-test) assignment
// sites in init.go.
var (
	_ server.AuthMetricsRecorder    = (*GatewayMetrics)(nil)
	_ auth.CredentialLookupRecorder = (*GatewayMetrics)(nil)
)

// GatewayMetrics holds OTel metrics for the multigateway service.
type GatewayMetrics struct {
	meter             metric.Meter
	clientConnections metric.Int64ObservableGauge

	// Auth metrics (mg.gateway.auth.*).
	authSCRAMDuration            metric.Float64Histogram
	authAttempts                 metric.Int64Counter
	authCredentialLookupDuration metric.Float64Histogram
	authCredentialLookupRate     metric.Int64Counter

	// TLS metrics (mg.gateway.tls.*).
	tlsHandshakeDuration  metric.Float64Histogram
	tlsConnections        metric.Int64Counter
	tlsPlaintextRejected  metric.Int64Counter
	tlsSSLRequestDeclined metric.Int64Counter
}

// NewGatewayMetrics initializes OTel metrics for the multigateway service.
// Individual metrics that fail to initialize use noop implementations and are
// included in the returned error. The returned instance is always usable.
func NewGatewayMetrics() (*GatewayMetrics, error) { _ = "STUB: not implemented"; return nil, nil }

// RegisterClientConnectionsCallback registers a callback that reports the current
// number of active client connections, broken down by endpoint.
// The replica getter may be nil when the replica-reads port is disabled.
func (m *GatewayMetrics) RegisterClientConnectionsCallback(primaryGetter, replicaGetter func() int) error {
	_ = "STUB: not implemented"
	return nil
}

// RecordSCRAMDuration records the SCRAM handshake duration tagged by outcome.
// Safe to call on a nil receiver so call sites can stay unconditional.
func (m *GatewayMetrics) RecordSCRAMDuration(ctx context.Context, outcome string, d time.Duration) {
	_ = "STUB: not implemented"
	return
}

// RecordAuthAttempt increments the auth-attempts counter tagged by outcome.
func (m *GatewayMetrics) RecordAuthAttempt(ctx context.Context, outcome string) {
	_ = "STUB: not implemented"
	return
}

// RecordCredentialLookup records the GetAuthCredentials RPC latency from the
// gateway side and increments the lookup-rate counter.
func (m *GatewayMetrics) RecordCredentialLookup(ctx context.Context, d time.Duration) {
	_ = "STUB: not implemented"
	return
}

// RecordTLSHandshake records the TLS handshake duration tagged by outcome.
func (m *GatewayMetrics) RecordTLSHandshake(ctx context.Context, outcome string, d time.Duration) {
	_ = "STUB: not implemented"
	return
}

// RecordTLSConnection increments the completed-TLS-connections counter tagged
// with the negotiated tls_version and cipher_suite. Names come from
// crypto/tls so they match Go's published constants (e.g. "TLS 1.3",
// "TLS_AES_128_GCM_SHA256").
func (m *GatewayMetrics) RecordTLSConnection(ctx context.Context, version, cipher uint16) {
	_ = "STUB: not implemented"
	return
}

// RecordPlaintextRejected increments the plaintext-rejection counter tagged
// by reason.
func (m *GatewayMetrics) RecordPlaintextRejected(ctx context.Context, reason string) {
	_ = "STUB: not implemented"
	return
}

// RecordSSLRequestDeclined increments the SSL-declined counter, used before
// enforcement is enabled to size the impact of flipping --pg-require-ssl.
func (m *GatewayMetrics) RecordSSLRequestDeclined(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}
