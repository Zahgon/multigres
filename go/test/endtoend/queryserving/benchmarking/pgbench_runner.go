// Copyright 2026 Supabase, Inc.
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

package benchmarking

import (
	"context"
	"testing"

	"github.com/multigres/multigres/go/test/endtoend/suiteutil"
)

// ScenarioConfig defines a single pgbench scenario to run.
type ScenarioConfig struct {
	Name      string // Human-readable name, e.g. "sustained_10c_extended"
	Clients   int    // -c flag
	Duration  int    // -T flag (seconds)
	Protocol  string // "simple" or "extended" (-M flag)
	ConnChurn bool   // -C flag (new connection per transaction)
}

// ScenarioResult holds computed metrics from a pgbench or sysbench run.
//
// Protocol/ConnChurn/LatencyP50 are pgbench-specific but kept non-omitempty
// so existing downstream tooling (e.g. the weekly pgbench workflow that
// filters with `select(.connection_churn == false)`) keeps working — their
// zero values for sysbench rows are harmless. PSMode is sysbench-only and
// stays `omitempty` so it doesn't pollute pgbench output.
type ScenarioResult struct {
	Target       string  `json:"target"`
	Scenario     string  `json:"scenario"`
	TPS          float64 `json:"tps"`
	LatencyAvg   float64 `json:"latency_avg_ms"`
	LatencyP50   float64 `json:"latency_p50_ms"`
	LatencyP99   float64 `json:"latency_p99_ms"`
	Clients      int     `json:"clients"`
	Duration     int     `json:"duration_seconds"`
	Protocol     string  `json:"protocol"`
	ConnChurn    bool    `json:"connection_churn"`
	PSMode       string  `json:"ps_mode,omitempty"`
	Transactions int     `json:"transactions"`
}

// BenchmarkReport is the top-level report containing all results.
type BenchmarkReport struct {
	Timestamp   string           `json:"timestamp"`
	LoadTool    string           `json:"load_tool,omitempty"` // "pgbench" or "sysbench"
	Results     []ScenarioResult `json:"results"`
	Environment EnvironmentInfo  `json:"environment"`
}

// EnvironmentInfo captures the test environment for reproducibility.
type EnvironmentInfo struct {
	PostgresVersion  string `json:"postgres_version,omitempty"`
	PgBenchVersion   string `json:"pgbench_version,omitempty"`
	SysBenchVersion  string `json:"sysbench_version,omitempty"`
	MultigresVersion string `json:"multigres_version,omitempty"`
	OS               string `json:"os"`
	Arch             string `json:"arch,omitempty"`
	GOMAXPROCS       int    `json:"gomaxprocs"`
	PlanCacheBytes   int    `json:"plan_cache_bytes,omitempty"`
}

// PgBenchRunner wraps pgbench execution: initialization, scenario runs, and output parsing.
type PgBenchRunner struct {
	pgbenchBinary string
	OutputDir     string
}

// NewPgBenchRunner creates a runner after locating the pgbench binary.
func NewPgBenchRunner(t *testing.T) *PgBenchRunner { _ = "STUB: not implemented"; return nil }

// Initialize runs pgbench -i to create the benchmark tables.
func (r *PgBenchRunner) Initialize(ctx context.Context, t *testing.T, target suiteutil.Target, scaleFactor int) error {
	_ = "STUB: not implemented"
	return nil
}

// RunScenario executes a single pgbench scenario and computes metrics from
// the per-transaction log file (--log), avoiding fragile text output parsing.
//
// pgbench --log writes one line per completed transaction:
//
//	client_id  transaction_no  time_epoch_us  latency_us  schedule_lag_us
//
// We compute TPS, average latency, p50, and p99 directly from this data.
func (r *PgBenchRunner) RunScenario(ctx context.Context, t *testing.T, target suiteutil.Target, scenario ScenarioConfig) (*ScenarioResult, error) {
	_ = "STUB: not implemented"

	// Create a unique log file prefix for this run inside the output directory.
	// Using output dir (not t.TempDir) so files survive for debugging if needed.
	return nil, nil
}

// pgbench creates log files named <prefix>.<pid> (one thread) or
// <prefix>.<pid>.<thread> (multi-thread). Glob for all of them.

// readTransactionLatencies reads per-transaction latency values (in microseconds)
// from one or more pgbench log files.
//
// Each line has the format: client_id  txn_no  latency_us  script_no  epoch_secs  epoch_usecs
// The latency field is at column index 2.
func readTransactionLatencies(t *testing.T, logFiles []string) ([]float64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// computeResult derives TPS, averages, and percentiles from raw latency data.
func computeResult(latencies []float64, scenario ScenarioConfig, targetName string) *ScenarioResult {
	_ = "STUB: not implemented"
	return nil

	// Average latency in microseconds.
}

// Sort for percentile computation.

// us → ms

// percentile returns the p-th percentile of a sorted slice using linear interpolation.
func percentile(sorted []float64, p float64) float64 { _ = "STUB: not implemented"; return 0 }

// CaptureEnvironment collects environment info for the report.
func CaptureEnvironment(t *testing.T, pgbenchBinary string) EnvironmentInfo {
	_ = "STUB: not implemented"
	return *new(EnvironmentInfo)
}

// DefaultScenarios generates the cross-product of benchmark scenarios.
//
// PGBENCH_PROTOCOLS (comma-separated subset of "simple,extended", default
// both) restricts which protocols are exercised. PGBENCH_NO_CHURN=1 skips
// the churn variants — useful for scaling sweeps where churn is noise.
func DefaultScenarios(duration int, clientCounts []int) []ScenarioConfig {
	_ = "STUB: not implemented"
	return nil
}

// Sustained load scenarios (all client counts)

// Connection churn scenarios (limited client counts — churn with many clients is very slow)

// ParseClients parses the PGBENCH_CLIENTS env var into a sorted list of client counts.
func ParseClients(t *testing.T) []int { _ = "STUB: not implemented"; return nil }

// ParseDuration reads the PGBENCH_DURATION env var (default: 30 seconds).
func ParseDuration(t *testing.T) int { _ = "STUB: not implemented"; return 0 }
