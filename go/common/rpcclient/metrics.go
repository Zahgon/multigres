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

package rpcclient

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// DialPath represents the possible paths for dialing a connection.
type DialPath string

const (
	DialPathCacheFast DialPath = "cache_fast" // Connection retrieved from cache without blocking
	DialPathSemaFast  DialPath = "sema_fast"  // Semaphore acquired without blocking to dial new connection
	DialPathSemaPoll  DialPath = "sema_poll"  // Polling for evictable connections while waiting for capacity
)

// Metrics holds all OpenTelemetry metrics for the rpcclient connection cache.
type Metrics struct {
	meter        metric.Meter
	connReuse    metric.Int64Counter
	connNew      metric.Int64Counter
	dialTimeouts metric.Int64Counter
	cacheSize    CacheSize
	dialDuration DialDuration
}

// CacheSize wraps an Int64ObservableGauge for observing cache size.
// Use the Inst() method to get the underlying gauge for callback registration.
type CacheSize struct {
	metric.Int64ObservableGauge
}

// Inst returns the underlying metric instrument for callback registration.
func (m CacheSize) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

// DialDuration wraps a Float64Histogram for recording dial operation durations.
// It abstracts away attribute key names so callers don't need to know them.
type DialDuration struct {
	metric.Float64Histogram
}

// Record records a dial operation duration with proper OTel attributes.
//
// Parameters:
//   - ctx: Context for the metric recording
//   - val: How long the dial took (in seconds)
//   - path: The dial path taken (cache_fast, sema_fast, or sema_poll)
//   - attrs: Optional additional attributes to include in the metric
func (m DialDuration) Record(
	ctx context.Context,
	val float64,
	path DialPath,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

// NewMetrics initializes OpenTelemetry metrics for the rpcclient connection cache.
// For the cacheSize observable gauge, use RegisterCacheSizeCallback() to register a callback.
func NewMetrics() *Metrics { _ = "STUB: not implemented"; return nil }

// Counter for connection reuse

// Counter for new connections

// Counter for dial timeouts

// Observable gauge for cache size

// Histogram for dial duration

// RegisterCacheSizeCallback registers a callback for the cache size observable gauge.
// The cacheGetter function is called periodically to observe the current cache size.
// Returns an error if callback registration fails.
func (m *Metrics) RegisterCacheSizeCallback(cacheGetter func() int) error {
	_ = "STUB: not implemented"
	return nil
}

// AddConnReuse increments the connection reuse counter.
func (m *Metrics) AddConnReuse(ctx context.Context) { _ = "STUB: not implemented"; return }

// AddConnNew increments the new connection counter.
func (m *Metrics) AddConnNew(ctx context.Context) { _ = "STUB: not implemented"; return }

// AddDialTimeout increments the dial timeout counter.
func (m *Metrics) AddDialTimeout(ctx context.Context) { _ = "STUB: not implemented"; return }

// RecordDialDuration records a dial operation duration with the specified path.
func (m *Metrics) RecordDialDuration(ctx context.Context, duration time.Duration, path DialPath) {
	_ = "STUB: not implemented"
	return
}
