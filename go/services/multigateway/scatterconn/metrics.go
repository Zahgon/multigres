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

package scatterconn

import (
	"context"

	"go.opentelemetry.io/otel/metric"
)

// ScatterStatus represents the outcome of a shard execution for metric attribution.
type ScatterStatus string

const (
	ScatterStatusOK    ScatterStatus = "ok"
	ScatterStatusError ScatterStatus = "error"
)

// ScatterMetrics holds all OTel metrics for scatterconn shard execution.
// Follows the same wrapper-type pattern as recovery/metrics.go.
type ScatterMetrics struct {
	executeDuration ExecuteDuration
	executeErrors   ExecuteErrors
}

// ExecuteDuration wraps a Float64Histogram for recording shard execution durations.
type ExecuteDuration struct {
	metric.Float64Histogram
}

// Record records a shard execution duration with proper OTel attributes.
func (m ExecuteDuration) Record(
	ctx context.Context,
	val float64,
	dbNamespace string,
	tablegroup string,
	shard string,
	status ScatterStatus,
) {
	_ = "STUB: not implemented"
	return
}

// ExecuteErrors wraps an Int64Counter for counting shard execution errors.
type ExecuteErrors struct {
	metric.Int64Counter
}

// Add increments the shard error counter with proper OTel attributes.
func (m ExecuteErrors) Add(
	ctx context.Context,
	tablegroup string,
	shard string,
	errorType string,
) {
	_ = "STUB: not implemented"
	return
}

// NewScatterMetrics initialises OTel metrics for scatterconn.
// Individual metrics that fail to initialise use noop implementations
// and are included in the returned error.
func NewScatterMetrics() (*ScatterMetrics, error) { _ = "STUB: not implemented"; return nil, nil }
