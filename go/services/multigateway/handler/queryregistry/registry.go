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

// Package queryregistry tracks per-query-shape statistics using a W-TinyLFU
// cache (via theine) keyed by query fingerprint. The TinyLFU doorkeeper
// ensures one-off queries don't evict popular ones, and its frequency
// comparison naturally promotes newly-popular queries.
//
// The registry is used to:
//   - Bound Prometheus label cardinality by deciding which fingerprints get
//     their own label value vs. fall into the "__other__" bucket.
//   - Back a /debug/queries endpoint that exposes aggregated per-query stats
//     (a pg_stat_statements equivalent).
package queryregistry

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/multigres/multigres/go/common/cache/theine"
)

// OtherLabel is the fingerprint label value used for queries not in the
// tracked top set. Keeping this bucket named lets us still emit metrics for
// the long tail without paying cardinality cost per fingerprint.
const OtherLabel = "__other__"

// UtilityLabel is used for statements we don't fingerprint individually
// (BEGIN/COMMIT/SET/DDL etc.) — their space is tiny and bounded.
const UtilityLabel = "__utility__"

// QueryStats holds aggregated statistics for a single query fingerprint.
// Counter fields use atomics so the hot Record path stays lockless. Trend
// ring buffers and the per-sample bookkeeping fields are protected by
// trendMu — only the sampler goroutine and Snapshot readers touch them.
type QueryStats struct {
	fingerprint   string
	normalizedSQL string

	calls           atomic.Uint64
	errors          atomic.Uint64
	totalDurationNs atomic.Int64
	totalRows       atomic.Uint64
	minDurationNs   atomic.Int64
	maxDurationNs   atomic.Int64
	lastSeenUnixNs  atomic.Int64

	// Per-fingerprint duration histogram. Bucket layout matches
	// durationBucketsNs; index numHistBuckets-1 is the +Inf overflow.
	durationBuckets [numHistBuckets]atomic.Uint64

	// Trend bookkeeping — guarded by trendMu. Last-* fields hold the
	// counter / histogram values at the previous sample so the sampler can
	// compute per-interval deltas. Without this delta-of-histograms the
	// trend percentile would be cumulative-since-admission and would
	// flat-line under sustained load.
	trendMu                   sync.Mutex
	trends                    trendBuffers
	lastSampleCalls           uint64
	lastSampleTotalDurationNs int64
	lastSampleTotalRows       uint64
	lastSampleBuckets         [numHistBuckets]uint64
}

// CachedSize implements theine's cacheval interface so theine can bound
// the registry by bytes, not entries. The fixed overhead covers the struct
// + atomics + histogram buckets; the variable part covers the per-instance
// strings and the trend ring buffers (5 rings × cap × 8 bytes/float64),
// charged eagerly at admit time so theine's admission policy isn't fooled
// into over-admitting once the sampler warms up.
func (s *QueryStats) CachedSize(_ bool) int64 { _ = "STUB: not implemented"; return 0 }

// Snapshot is a point-in-time copy of a QueryStats, safe to hand out to
// HTTP handlers or serialize to JSON. Trend slices are oldest-to-newest
// rolling samples captured by the registry's sampler goroutine — empty
// when sampling is disabled or no samples have been taken yet.
type Snapshot struct {
	Fingerprint     string        `json:"fingerprint"`
	NormalizedSQL   string        `json:"normalized_sql"`
	Calls           uint64        `json:"calls"`
	Errors          uint64        `json:"errors"`
	TotalDuration   time.Duration `json:"total_duration_ns"`
	AverageDuration time.Duration `json:"average_duration_ns"`
	MinDuration     time.Duration `json:"min_duration_ns"`
	MaxDuration     time.Duration `json:"max_duration_ns"`
	P50Duration     time.Duration `json:"p50_duration_ns"`
	P99Duration     time.Duration `json:"p99_duration_ns"`
	TotalRows       uint64        `json:"total_rows"`
	LastSeen        time.Time     `json:"last_seen"`

	// SampleIntervalSeconds is the cadence at which the trend slices
	// below were captured. 0 means trends are disabled.
	SampleIntervalSeconds float64 `json:"sample_interval_seconds"`

	// CallRateTrend is calls/s, oldest sample first.
	CallRateTrend []float64 `json:"call_rate_trend,omitempty"`
	// TotalTimeMsTrend is duration-ms-per-second of wall clock, oldest first.
	TotalTimeMsTrend []float64 `json:"total_time_ms_trend,omitempty"`
	// P50MsTrend is p50 latency in ms over time, oldest first.
	P50MsTrend []float64 `json:"p50_ms_trend,omitempty"`
	// P99MsTrend is p99 latency in ms over time, oldest first.
	P99MsTrend []float64 `json:"p99_ms_trend,omitempty"`
	// RowsRateTrend is rows/s, oldest first.
	RowsRateTrend []float64 `json:"rows_rate_trend,omitempty"`
}

// Registry tracks per-fingerprint query statistics.
// The zero value is not usable — construct via New.
type Registry struct {
	store      *theine.Store[theine.StringKey, *QueryStats]
	maxSQLLen  int
	newEntryMu sync.Mutex

	// Sampler config; trendCapacity == 0 disables trend collection.
	sampleInterval time.Duration
	trendCapacity  int

	// Sampler goroutine lifecycle.
	samplerCancel context.CancelFunc
	samplerDone   chan struct{}
}

// Config configures a Registry.
type Config struct {
	// MaxMemoryBytes caps the memory used by the tracked fingerprint set.
	// A value <= 0 disables the registry (all calls become no-ops).
	MaxMemoryBytes int
	// MaxSQLLength caps the stored representative normalized SQL for each
	// fingerprint. Queries longer than this are truncated in the stored
	// copy; the fingerprint itself is still computed over the full text.
	MaxSQLLength int
	// SampleInterval is the cadence at which the trend sampler captures
	// per-fingerprint deltas (calls/s, p50/p99 ms, etc.) into a rolling
	// ring. 0 disables sampling.
	SampleInterval time.Duration
	// TrendWindowSamples is the number of samples retained per metric per
	// fingerprint. The visible time window is SampleInterval * this value.
	// 0 disables sampling regardless of SampleInterval.
	TrendWindowSamples int
}

// DefaultConfig returns reasonable defaults for the registry.
func DefaultConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

// 8 MB — sized so the per-fingerprint trend rings (5 × TrendWindowSamples × 8 B) don't shrink admission capacity vs the pre-trend layout.

// 60 × 10s = 10-minute trend window

// New constructs a Registry with the given config.
// If cfg.MaxMemoryBytes <= 0 the registry is disabled and all methods become no-ops.
func New(cfg Config) *Registry { _ = "STUB: not implemented"; return nil }

// NewForTest constructs a Registry without the TinyLFU doorkeeper so tests
// can assert deterministic admission behavior.
func NewForTest(cfg Config) *Registry { _ = "STUB: not implemented"; return nil }

func newRegistry(cfg Config, doorkeeper bool) *Registry { _ = "STUB: not implemented"; return nil }

//nolint:gocritic // Long-lived sampler tied to the registry's lifetime, cancelled in Close().

// Record updates stats for the given fingerprint. If the fingerprint is not
// yet tracked, the TinyLFU admission policy decides whether to admit it —
// one-off queries won't enter the tracked set on their first hit.
func (r *Registry) Record(fingerprint, normalizedSQL string, duration time.Duration, rows int, hadError bool) {
	_ = "STUB: not implemented"
	return
}

// Try to admit — may fail if TinyLFU rejects a cold entry.

// Update min/max with CAS loops.

// admit attempts to insert a new QueryStats entry for fingerprint.
// Returns the stats (either newly admitted or racing winner's), or nil if
// TinyLFU rejected the admission.
func (r *Registry) admit(fingerprint, normalizedSQL string) *QueryStats {
	_ = "STUB: not implemented"
	return nil
}

// Re-check after acquiring lock: another goroutine may have admitted it.

// Allocate up-front so CachedSize reflects the real footprint and
// theine's admission policy can use accurate sizes.

// Reload through Get so we return the authoritative stored pointer —
// theine may have rejected via doorkeeper and returned false without
// actually admitting.

// Labelize returns the fingerprint if it is currently tracked in the
// registry, otherwise returns OtherLabel. Use this as the value for
// a Prometheus/OTel label to bound cardinality.
func (r *Registry) Labelize(fingerprint string) string { _ = "STUB: not implemented"; return "" }

// SortKey identifies how Top should sort results.
type SortKey string

const (
	SortByCalls       SortKey = "calls"
	SortByTotalTime   SortKey = "total_time"
	SortByAverageTime SortKey = "avg_time"
	SortByErrors      SortKey = "errors"
	SortByLastSeen    SortKey = "last_seen"
)

// Top returns up to `limit` snapshots of the currently-tracked query
// statistics, sorted by the given key in descending order. Passing
// limit <= 0 returns all tracked entries.
func (r *Registry) Top(limit int, sortBy SortKey) []Snapshot { _ = "STUB: not implemented"; return nil }

// Len returns the number of fingerprints currently tracked in the registry.
func (r *Registry) Len() int { _ = "STUB: not implemented"; return 0 }

// Close stops the background maintenance goroutine.
// Safe to call on a nil or disabled registry.
func (r *Registry) Close() { _ = "STUB: not implemented"; return }

// runSampler periodically captures a per-fingerprint sample (rates +
// histogram percentiles) into each entry's trend rings until ctx is done.
func (r *Registry) runSampler(ctx context.Context) { _ = "STUB: not implemented"; return }

// sampleAll walks the registry and records one trend sample per entry.
func (r *Registry) sampleAll() { _ = "STUB: not implemented"; return }

// takeSample reads the current counters/histogram, computes deltas vs. the
// previous sample, and pushes one new value into each trend ring. The rings
// are pre-allocated by admit() when the registry has sampling enabled, so a
// nil-or-uninitialised state means the entry pre-dates the sampler config
// and trend pushes are skipped.
func (s *QueryStats) takeSample(intervalSec float64) { _ = "STUB: not implemented"; return }

// Percentile sparkline tracks per-interval distribution: snapshot the
// histogram, push the percentile of (current - previous), then save
// the snapshot for the next sample. Without this delta, the sparkline
// would show cumulative-since-admission percentiles and flatten out.

func (s *QueryStats) snapshot(sampleIntervalSec float64) Snapshot {
	_ = "STUB: not implemented"
	return *new(Snapshot)
}

func sortSnapshots(snapshots []Snapshot, sortBy SortKey) { _ = "STUB: not implemented"; return }
