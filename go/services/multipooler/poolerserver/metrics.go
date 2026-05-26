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

package poolerserver

import (
	"context"

	"go.opentelemetry.io/otel/metric"
)

// drainStats holds OpenTelemetry metrics for graceful-drain observability.
// Operators use these to size --connpool-drain-grace-period and to
// decompose the gateway-perceived failover duration into pooler-time vs.
// new-primary-election time.
type drainStats struct {
	meter metric.Meter

	duration    metric.Float64Histogram
	outcome     metric.Int64Counter
	forceClosed metric.Int64Counter
}

func newDrainStats() *drainStats { _ = "STUB: not implemented"; return nil }

const (
	drainOutcomeGraceful   = "graceful"
	drainOutcomeForceClose = "force_close"
)

// recordDrain records a completed drain event with its wall-clock duration
// and outcome.
func (s *drainStats) recordDrain(ctx context.Context, seconds float64, outcome string) {
	_ = "STUB: not implemented"
	return
}

// recordForceClosed adds to the count of connections force-closed across
// all drain events.
func (s *drainStats) recordForceClosed(ctx context.Context, n int) {
	_ = "STUB: not implemented"
	return
}
