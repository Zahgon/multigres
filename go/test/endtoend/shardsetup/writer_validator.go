// Copyright 2025 Supabase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package shardsetup

import (
	"context"
	"database/sql"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// WriterValidator continuously writes to a test table and tracks successful/failed writes.
// Useful for validating data durability during failover scenarios.
type WriterValidator struct {
	tableName     string
	workerCount   int
	writeInterval time.Duration
	queryTimeout  time.Duration

	db *sql.DB

	nextID atomic.Int64

	mu           sync.Mutex
	successful   []int64
	failed       []int64
	failedErrors []string // error messages for each failed write (parallel to failed slice)

	ctx     context.Context
	cancel  context.CancelFunc
	stop    chan struct{} // signals workers to stop issuing new writes
	wg      sync.WaitGroup
	started bool
}

// WriterValidatorOption configures a WriterValidator.
type WriterValidatorOption func(*WriterValidator)

// WithWorkerCount sets the number of concurrent writer goroutines (default: 4).
func WithWorkerCount(count int) WriterValidatorOption {
	_ = "STUB: not implemented"
	return *new(WriterValidatorOption)
}

// WithWriteInterval sets the interval between writes per worker (default: 10ms).
func WithWriteInterval(interval time.Duration) WriterValidatorOption {
	_ = "STUB: not implemented"
	return *new(WriterValidatorOption)
}

// WithQueryTimeout sets the per-query timeout (default: 5s). When buffering is
// enabled on the gateway, this should be at least as long as the buffer window
// so that the buffer can drain before the client gives up.
func WithQueryTimeout(timeout time.Duration) WriterValidatorOption {
	_ = "STUB: not implemented"
	return *new(WriterValidatorOption)
}

// NewWriterValidator creates a new WriterValidator for the given sql.DB connection.
// It creates the test table immediately and returns a cleanup function that drops it.
func NewWriterValidator(t *testing.T, db *sql.DB, opts ...WriterValidatorOption) (*WriterValidator, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Stop workers if running

// Drop table (best effort, use background context since test context may be done)

// TableName returns the name of the test table.
func (w *WriterValidator) TableName() string {
	_ = "STUB: not implemented"

	// createTable creates the test table.
	return ""
}

func (w *WriterValidator) createTable(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// dropTable drops the test table.
func (w *WriterValidator) dropTable(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Start spawns worker goroutines that continuously write to the table.
func (w *WriterValidator) Start(t *testing.T) { _ = "STUB: not implemented"; return }

// Stop gracefully stops all worker goroutines. It signals workers to stop
// issuing new writes, then waits for any in-flight writes to complete.
// This avoids context cancellation errors on in-flight queries.
func (w *WriterValidator) Stop() { _ = "STUB: not implemented"; return }

// Signal workers to stop, then wait for in-flight writes to finish.

// worker runs the write loop using a ticker.
func (w *WriterValidator) worker() { _ = "STUB: not implemented"; return }

// #nosec G202 -- tableName is a constant from test setup, not user-controlled.

// recordResult records a write attempt result.
func (w *WriterValidator) recordResult(id int64, err error) { _ = "STUB: not implemented"; return }

// SuccessfulWrites returns a copy of all successfully written IDs.
func (w *WriterValidator) SuccessfulWrites() []int64 { _ = "STUB: not implemented"; return nil }

// FailedWrites returns a copy of all failed write IDs.
func (w *WriterValidator) FailedWrites() []int64 { _ = "STUB: not implemented"; return nil }

// FailedErrors returns a summary of all failed write errors. It deduplicates
// error messages and returns a map of error message → count.
func (w *WriterValidator) FailedErrors() map[string]int { _ = "STUB: not implemented"; return nil }

// Stats returns the count of successful and failed writes.
func (w *WriterValidator) Stats() (successful, failed int) { _ = "STUB: not implemented"; return 0, 0 }

// WriteNow performs a single write immediately, records the result, and returns any error.
// Useful for probing whether the connection is routed to a writable primary.
// Safe to call concurrently with running workers.
func (w *WriterValidator) WriteNow(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// #nosec G202 -- tableName is a constant from test setup, not user-controlled.

// Verify checks that all successful writes are present in at least one of the provided poolers.
func (w *WriterValidator) Verify(t *testing.T, poolers []*MultiPoolerTestClient) error {
	_ = "STUB: not implemented"
	return nil
}

// Build set of all IDs found across all pooler connections

// Check all successful writes are present
