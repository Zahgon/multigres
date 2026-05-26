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

import (
	"context"

	"go.opentelemetry.io/otel/metric"
)

// stats holds OpenTelemetry metrics for failover buffering.
type stats struct {
	meter metric.Meter

	requestsBuffered metric.Int64Counter
	requestsDrained  metric.Int64Counter
	requestsEvicted  metric.Int64Counter
	requestsSkipped  metric.Int64Counter
	failoverCount    metric.Int64Counter
	waitDuration     metric.Float64Histogram
	// queueDepth tracks the number of requests currently held in the buffer.
	// Used to size the buffer-size cap against observed peak depth.
	queueDepth metric.Int64UpDownCounter
	// failoverDuration records the time each failover spent in BUFFERING
	// before the gateway saw a new PRIMARY and began draining. Used to size
	// buffer-window against real failover lengths.
	failoverDuration metric.Float64Histogram
}

func newStats() *stats { _ = "STUB: not implemented"; return nil }

func (s *stats) recordBuffered(ctx context.Context, shardKey string) {
	_ = "STUB: not implemented"
	return
}

func (s *stats) recordDrained(ctx context.Context, shardKey string) {
	_ = "STUB: not implemented"
	return
}

func (s *stats) recordEvicted(ctx context.Context, shardKey string, reason string) {
	_ = "STUB: not implemented"
	return
}

func (s *stats) recordSkipped(ctx context.Context, reason string) {
	_ = "STUB: not implemented"
	return
}

func (s *stats) recordFailover(ctx context.Context, shardKey string) {
	_ = "STUB: not implemented"
	return
}

func (s *stats) recordWaitDuration(ctx context.Context, seconds float64) {
	_ = "STUB: not implemented"
	return
}

func (s *stats) addQueueDepth(ctx context.Context, delta int64) { _ = "STUB: not implemented"; return }

func (s *stats) recordFailoverDuration(ctx context.Context, shardKey string, seconds float64) {
	_ = "STUB: not implemented"
	return
}
