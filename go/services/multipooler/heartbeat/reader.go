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
	"sync"
	"sync/atomic"
	"time"

	"github.com/multigres/multigres/go/services/multipooler/executor"
	"github.com/multigres/multigres/go/tools/timer"
)

const (
	defaultHeartbeatReadInterval = 1 * time.Second
)

// Reader reads the heartbeat table at a configured interval in order
// to calculate replication lag. It is meant to be run on a replica, and paired
// with a Writer on a primary.
// Lag is calculated by comparing the most recent timestamp in the heartbeat
// table against the current time at read time.
type Reader struct {
	queryService executor.InternalQueryService
	logger       *slog.Logger
	shardID      []byte
	interval     time.Duration
	now          func() time.Time

	runner *timer.PeriodicRunner

	lagMu          sync.Mutex
	lastKnownLag   time.Duration
	lastKnownTime  time.Time
	lastKnownError error

	reads      atomic.Int64
	readErrors atomic.Int64
}

// NewReader returns a new heartbeat reader with the default interval.
func NewReader(queryService executor.InternalQueryService, logger *slog.Logger, shardID []byte) *Reader {
	_ = "STUB: not implemented"
	return nil
}

// newReader creates a heartbeat reader with a configurable interval.
func newReader(queryService executor.InternalQueryService, logger *slog.Logger, shardID []byte, interval time.Duration) *Reader {
	_ = "STUB: not implemented"
	return nil
}

// Open starts the heartbeat ticker.
func (r *Reader) Open() { _ = "STUB: not implemented"; return }

// Close cancels the readHeartbeat periodic ticker. After Close returns,
// no more heartbeat reads will be made and any in-flight read has completed.
func (r *Reader) Close() { _ = "STUB: not implemented"; return }

// IsOpen returns true if the reader is open.
func (r *Reader) IsOpen() bool { _ = "STUB: not implemented"; return false }

// Status returns the most recently recorded lag measurement or error encountered.
func (r *Reader) Status() (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

// Return an error if we didn't receive a heartbeat for more than two intervals

// readHeartbeat reads from the heartbeat table exactly once, updating
// the last known lag and/or error, and incrementing counters.
func (r *Reader) readHeartbeat(ctx context.Context) { _ = "STUB: not implemented"; return }

// fetchMostRecentHeartbeat fetches the most recently recorded heartbeat from the heartbeat table,
// returning the timestamp of the heartbeat in nanoseconds.
func (r *Reader) fetchMostRecentHeartbeat(ctx context.Context) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// recordError keeps track of the lastKnown error for reporting to Status().
func (r *Reader) recordError(err error) { _ = "STUB: not implemented"; return }

// Reads returns the number of successful heartbeat reads.
func (r *Reader) Reads() int64 { _ = "STUB: not implemented"; return 0 }

// ReadErrors returns the number of heartbeat read errors.
func (r *Reader) ReadErrors() int64 { _ = "STUB: not implemented"; return 0 }

// LeadershipView contains the consensus state and replication lag information
type LeadershipView struct {
	LeaderID       string
	LastHeartbeat  time.Time
	ReplicationLag time.Duration
}

// GetLeadershipView returns both replication lag and consensus state
func (r *Reader) GetLeadershipView() (*LeadershipView, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Calculate replication lag
