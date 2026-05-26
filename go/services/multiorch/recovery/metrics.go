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

package recovery

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// Metrics holds all OpenTelemetry metrics for the recovery engine.
// Following OTel conventions, metrics are defined in a separate file with
// wrapper types that hide attribute key names from instrumented code.
//
// This pattern is inspired by:
// https://github.com/open-telemetry/opentelemetry-go/blob/v1.38.0/semconv/v1.37.0/dbconv/metric.go
type Metrics struct {
	meter                  metric.Meter
	poolerStoreSize        PoolerStoreSize
	recoveryActionDuration RecoveryActionDuration
	errorsTotal            ErrorsTotal
	detectedProblems       DetectedProblems
	streamConnected        StreamConnected
	streamSnapshotsTotal   StreamSnapshotsTotal
}

// PoolerStoreSize wraps an Int64ObservableGauge for observing pooler store size.
// Use the Inst() method to get the underlying gauge for callback registration.
type PoolerStoreSize struct {
	metric.Int64ObservableGauge
}

// Inst returns the underlying metric instrument for callback registration.
func (m PoolerStoreSize) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

// RecoveryActionStatus represents the possible status values for a recovery action.
type RecoveryActionStatus string

const (
	RecoveryActionStatusSuccess RecoveryActionStatus = "success"
	RecoveryActionStatusFailure RecoveryActionStatus = "failure"
)

// RecoveryActionDuration wraps a Float64Histogram for recording recovery action durations.
type RecoveryActionDuration struct {
	metric.Float64Histogram
}

// Record records a recovery action duration with proper OTel attributes.
//
// Parameters:
//   - ctx: Context for the metric recording
//   - val: How long the action took (in milliseconds)
//   - actionName: The name of the recovery action (e.g., "FixReplication", "BootstrapShard")
//   - problemCode: The problem code being addressed
//   - status: The action status (success or failure)
//   - dbNamespace: The database name (becomes "db.namespace" attribute per OTel spec)
//   - shard: The shard identifier
func (m RecoveryActionDuration) Record(
	ctx context.Context,
	val float64,
	actionName string,
	problemCode string,
	status RecoveryActionStatus,
	dbNamespace string,
	shard string,
) {
	_ = "STUB: not implemented"
	return
}

// ErrorsTotal wraps an Int64Counter for counting errors in the recovery engine.
// Use the Add method to increment the counter with source information.
type ErrorsTotal struct {
	metric.Int64Counter
}

// Add increments the error counter with the given source component.
//
// Parameters:
//   - ctx: Context for the metric recording
//   - source: The component that produced the error (e.g., "analyzer", "recovery_action")
//   - attrs: Optional additional attributes to include in the metric
func (m ErrorsTotal) Add(ctx context.Context, source string, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

// DetectedProblems wraps an Int64ObservableGauge for observing detected problems by type.
// Use the Inst() method to get the underlying gauge for callback registration.
type DetectedProblems struct {
	metric.Int64ObservableGauge
}

// Inst returns the underlying metric instrument for callback registration.
func (m DetectedProblems) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

// NewMetrics initializes OpenTelemetry metrics for the recovery engine.
// Individual metrics that fail to initialize will use noop implementations and be included
// in the returned error. For the poolerStoreSize observable gauge, use
// RegisterPoolerStoreSizeCallback() to register a callback.
//
// Returns a Metrics instance (with noop fallbacks for failed metrics) and any initialization
// errors that occurred. The caller should log or handle these errors as appropriate.
func NewMetrics() (*Metrics, error) { _ = "STUB: not implemented"; return nil, nil }

// Gauge for current pooler store size

// Histogram for recovery action duration

// Counter for errors

// Gauge for detected problems by type

// Gauge for stream connection status per pooler (1 = connected, 0 = disconnected)

// Gauge for cumulative snapshots received per pooler

// RegisterPoolerStoreSizeCallback registers a callback for the pooler store size observable gauge.
// The poolerStoreGetter function is called periodically to observe the current store size.
// Returns an error if callback registration fails.
func (m *Metrics) RegisterPoolerStoreSizeCallback(poolerStoreGetter func() int) error {
	_ = "STUB: not implemented"
	return nil
}

// StreamConnected wraps an Int64ObservableGauge that reports 1 when a pooler's
// ManagerHealthStream stream is connected, 0 when disconnected.
// Use the Inst() method to get the underlying gauge for callback registration.
type StreamConnected struct {
	metric.Int64ObservableGauge
}

// Inst returns the underlying metric instrument for callback registration.
func (m StreamConnected) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

// StreamSnapshotsTotal wraps an Int64ObservableGauge that reports the cumulative
// number of health snapshots received per pooler since the stream was started.
// Use the Inst() method to get the underlying gauge for callback registration.
type StreamSnapshotsTotal struct {
	metric.Int64ObservableGauge
}

// Inst returns the underlying metric instrument for callback registration.
func (m StreamSnapshotsTotal) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

// StreamHealthData holds per-pooler stream health data for metric observation.
type StreamHealthData struct {
	PoolerID          string
	DBNamespace       string
	Shard             string
	Connected         bool
	SnapshotsReceived int64
}

// DetectedProblemData represents a detected problem with its attributes for metric observation.
// Each problem is tracked per affected entity (pooler ID or shard key).
type DetectedProblemData struct {
	AnalysisType string
	DBNamespace  string
	Shard        string
	EntityID     string
}

// RegisterStreamHealthCallback registers a callback for the stream health gauges.
// The getter is called periodically to observe current stream state for all poolers.
// Both multiorch.recovery.stream.connected and multiorch.recovery.stream.snapshots_received
// are updated in the same callback to keep them consistent.
// Returns an error if callback registration fails.
func (m *Metrics) RegisterStreamHealthCallback(getter func() []StreamHealthData) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterDetectedProblemsCallback registers a callback for the detected problems observable gauge.
// The getter function is called periodically to observe current detected problems.
// Each problem is reported with value=1 per pooler.
// Returns an error if callback registration fails.
func (m *Metrics) RegisterDetectedProblemsCallback(getter func() []DetectedProblemData) error {
	_ = "STUB: not implemented"
	return nil
}
