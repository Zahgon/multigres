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

package manager

import (
	"context"

	"go.opentelemetry.io/otel/metric"
)

const meterName = "github.com/multigres/multigres/go/services/multipooler/manager"

// Metrics holds OTel counter instruments for pgBackRest backup outcomes.
type Metrics struct {
	backupAttempts  metric.Int64Counter
	backupSuccesses metric.Int64Counter
	backupFailures  metric.Int64Counter

	restoreAttempts  metric.Int64Counter
	restoreSuccesses metric.Int64Counter
	restoreFailures  metric.Int64Counter

	backupDuration       metric.Float64Histogram
	backupVerifyDuration metric.Float64Histogram
	restoreDuration      metric.Float64Histogram
	backupLockWait       metric.Float64Histogram
}

// NewMetrics creates and registers the pgBackRest backup counters.
// It always returns a non-nil *Metrics; any registration errors are collected and returned.
func NewMetrics() (*Metrics, error) { _ = "STUB: not implemented"; return nil, nil }

// IncBackupAttempts increments the backup attempts counter.
// Safe to call on a nil receiver or with a nil instrument.
func (m *Metrics) IncBackupAttempts(ctx context.Context) { _ = "STUB: not implemented"; return }

// IncBackupSuccesses increments the backup successes counter.
// Safe to call on a nil receiver or with a nil instrument.
func (m *Metrics) IncBackupSuccesses(ctx context.Context) { _ = "STUB: not implemented"; return }

// IncBackupFailures increments the backup failures counter.
// Safe to call on a nil receiver or with a nil instrument.
func (m *Metrics) IncBackupFailures(ctx context.Context) { _ = "STUB: not implemented"; return }

func (m *Metrics) IncRestoreAttempts(ctx context.Context) { _ = "STUB: not implemented"; return }

func (m *Metrics) IncRestoreSuccesses(ctx context.Context) { _ = "STUB: not implemented"; return }

func (m *Metrics) IncRestoreFailures(ctx context.Context) { _ = "STUB: not implemented"; return }

func (m *Metrics) RecordBackupDuration(ctx context.Context, seconds float64) {
	_ = "STUB: not implemented"
	return
}

func (m *Metrics) RecordBackupVerifyDuration(ctx context.Context, seconds float64) {
	_ = "STUB: not implemented"
	return
}

func (m *Metrics) RecordRestoreDuration(ctx context.Context, seconds float64) {
	_ = "STUB: not implemented"
	return
}

func (m *Metrics) RecordBackupLockWait(ctx context.Context, seconds float64) {
	_ = "STUB: not implemented"
	return
}
