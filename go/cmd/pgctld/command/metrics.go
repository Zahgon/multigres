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

package command

import (
	"sync/atomic"

	"go.opentelemetry.io/otel/metric"
)

const pgctldMeterName = "github.com/multigres/multigres/go/cmd/pgctld/command"

// Metrics holds OTel gauge instruments for pgBackRest server health.
// Atomic values are used because the OTel callback and setPgBackRestStatus run on different goroutines.
type Metrics struct {
	serverUp     metric.Int64ObservableGauge
	restartCount metric.Int64ObservableGauge
	serverUptime metric.Int64ObservableGauge

	serverUpVal     atomic.Int64
	restartCountVal atomic.Int64
	startedAtNanos  atomic.Int64
}

// NewMetrics creates and registers the pgBackRest health gauges.
// It always returns a non-nil *Metrics; any registration errors are collected and returned.
func NewMetrics() (*Metrics, error) { _ = "STUB: not implemented"; return nil, nil }

// SetServerUp atomically stores 1 if running is true, 0 if false.
// It also records the start timestamp for the uptime gauge.
func (m *Metrics) SetServerUp(running bool) { _ = "STUB: not implemented"; return }

// SetRestartCount atomically stores the restart count as int64.
func (m *Metrics) SetRestartCount(count int32) { _ = "STUB: not implemented"; return }
