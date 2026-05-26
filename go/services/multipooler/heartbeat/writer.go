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

// Package heartbeat is responsible for reading and writing heartbeats
// to the heartbeat table.
package heartbeat

import (
	"context"
	"log/slog"
	"sync/atomic"
	"time"

	"github.com/multigres/multigres/go/services/multipooler/executor"
	"github.com/multigres/multigres/go/tools/timer"
)

// Make these modifiable for testing.
var (
	defaultHeartbeatInterval = 1 * time.Second
)

// Writer runs on primary databases and writes heartbeats to the heartbeat
// table at regular intervals.
type Writer struct {
	queryService executor.InternalQueryService
	logger       *slog.Logger
	shardID      []byte
	poolerID     string
	interval     time.Duration
	now          func() time.Time

	runner *timer.PeriodicRunner

	writes      atomic.Int64
	writeErrors atomic.Int64
}

// NewWriter creates a new heartbeat writer.
//
// We do not support on-demand or disabled heartbeats at this time.
func NewWriter(queryService executor.InternalQueryService, logger *slog.Logger, shardID []byte, poolerID string, intervalMs int) *Writer {
	_ = "STUB: not implemented"
	return nil
}

// Open starts the heartbeat writer.
func (w *Writer) Open() { _ = "STUB: not implemented"; return }

// Close stops the heartbeat writer. After Close returns, no more heartbeat
// writes will be made and any in-flight write has completed.
func (w *Writer) Close() { _ = "STUB: not implemented"; return }

// IsOpen returns true if the writer is open.
func (w *Writer) IsOpen() bool { _ = "STUB: not implemented"; return false }

// writeHeartbeat updates the heartbeat row with the current time in nanoseconds.
func (w *Writer) writeHeartbeat(ctx context.Context) { _ = "STUB: not implemented"; return }

// write writes a single heartbeat update.
func (w *Writer) write(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Writes returns the number of successful heartbeat writes.
func (w *Writer) Writes() int64 { _ = "STUB: not implemented"; return 0 }

// WriteErrors returns the number of heartbeat write errors.
func (w *Writer) WriteErrors() int64 { _ = "STUB: not implemented"; return 0 }
