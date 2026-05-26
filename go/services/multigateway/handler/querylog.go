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

package handler

import (
	"context"
	"log/slog"
	"sync/atomic"
	"time"
)

// queryLogEntry holds the fields emitted in a structured query log record.
type queryLogEntry struct {
	User          string
	Database      string
	OperationName string
	Protocol      string // "simple" or "extended"
	TotalDuration time.Duration
	ParseDuration time.Duration
	PlanDuration  time.Duration
	ExecDuration  time.Duration
	RowCount      int64
	PlanType      string
	TablesUsed    []string
	Error         error
	SQLSTATE      string
	ErrorSource   string
}

// emitQueryLog writes a structured query log entry using slog.LogAttrs for
// minimal allocation on the latency-sensitive query path.
//
// Errored or slow queries always log at WARN. Normal queries log at DEBUG
// (so the default INFO-level handler drops them) after 1/sampleRate sampling
// (sampleRate==0 disables sampling and lets the handler level alone govern).
//
// The slog-OTel bridge (configured in telemetry.go) automatically injects
// trace_id and span_id from the context.
func emitQueryLog(
	ctx context.Context,
	logger *slog.Logger,
	entry queryLogEntry,
	slowThreshold time.Duration,
	sampleRate uint64,
	samplingCursor *atomic.Uint64,
	emitsMetric QueryLogEmits,
) {
	_ = "STUB: not implemented"
	return
}

// Check sampling before logger.Enabled so the disabled-handler path
// also benefits from sampling skipping work.
