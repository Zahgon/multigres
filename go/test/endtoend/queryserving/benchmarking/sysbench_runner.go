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
	"regexp"
	"testing"

	"github.com/multigres/multigres/go/test/endtoend/suiteutil"
)

const (
	sysbenchPSModePrepared = "prepared" // sysbench --db-ps-mode=auto
	sysbenchPSModeSimple   = "simple"   // sysbench --db-ps-mode=disable
)

// SysBenchScenario defines a single sysbench oltp_point_select scenario.
type SysBenchScenario struct {
	Name     string // e.g. "sustained_32c_prepared"
	Clients  int
	Duration int    // seconds
	PSMode   string // "prepared" or "simple"
}

// SysBenchRunner wraps the sysbench CLI: detection, prepare, scenario run,
// and stdout parsing.
type SysBenchRunner struct {
	binary    string
	OutputDir string

	// Tables / table-size used by both prepare and run. They MUST match
	// across phases — sysbench's run phase looks up the same sbtestN tables
	// it created in prepare.
	Tables    int
	TableSize int
}

// NewSysBenchRunner locates the sysbench binary, verifies pgsql driver
// support, and creates the output directory.
//
// If sysbench is missing or lacks pgsql, the test is skipped (not failed) —
// this matches how TestPgBench handles a missing pgbench binary.
func NewSysBenchRunner(t *testing.T) *SysBenchRunner { _ = "STUB: not implemented"; return nil }

// checkSysBenchPgsqlDriver returns nil if the sysbench binary supports the
// pgsql driver, otherwise returns an error with an actionable message.
//
// Detection probes by invoking sysbench with --db-driver=pgsql against an
// unreachable host. If the driver isn't compiled in, sysbench fails fast
// with "invalid database driver name". Any other failure (including the
// expected "Connection refused") means the driver is present.
func checkSysBenchPgsqlDriver(binary string) error { _ = "STUB: not implemented"; return nil }

// pgsqlConnArgs returns the --pgsql-* flags that identify a target.
func (r *SysBenchRunner) pgsqlConnArgs(target suiteutil.Target) []string {
	_ = "STUB: not implemented"
	return nil
}

// commonOltpArgs returns the workload args that are constant across phases.
func (r *SysBenchRunner) commonOltpArgs() []string { _ = "STUB: not implemented"; return nil }

// Prepare creates the sbtestN tables on the given target. Idempotent across
// runs only in the sense that we record stdout and surface errors loudly —
// sysbench itself errors if the tables already exist, which the caller can
// inspect in logs/sysbench-prepare.txt.
func (r *SysBenchRunner) Prepare(ctx context.Context, t *testing.T, target suiteutil.Target) error {
	_ = "STUB: not implemented"
	return nil
}

// RunScenario executes one sysbench oltp_point_select scenario and parses
// the summary block from its stdout.
func (r *SysBenchRunner) RunScenario(ctx context.Context, t *testing.T, target suiteutil.Target, sc SysBenchScenario) (*ScenarioResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Force 99th-percentile reporting in the summary block. sysbench
// 1.0.x defaults to 95 (which then ends up labelled "95th
// percentile:" in the output and fails our regex).

// Sysbench summary lines we parse out of the stdout block. Examples (from
// sysbench 1.0.x and 1.1.x — both supported):
//
//	transactions:                        1940425 (32338.55 per sec.)
//	    avg:                                    0.99
//	    99th percentile:                        0.00
var (
	reTransactions = regexp.MustCompile(`transactions:\s+(\d+)\s+\(([\d.]+) per sec\.\)`)
	reAvgLatency   = regexp.MustCompile(`(?m)^\s*avg:\s+([\d.]+)`)
	reP99Latency   = regexp.MustCompile(`(?m)^\s*99th percentile:\s+([\d.]+)`)
)

// parseSysBenchOutput extracts TPS, transactions, avg latency and p99
// latency from sysbench's summary block. Returns an error if any required
// field is missing — we never want to silently record zero metrics.
func parseSysBenchOutput(out []byte) (*ScenarioResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DefaultSysBenchScenarios returns the cross-product of the env-driven
// client list and ps-mode list. Naming convention:
// `sustained_<clients>c_<ps_mode>` — distinct from the pgbench harness's
// `sustained_<clients>c_simple` scenario name to avoid output-dir collisions
// when both tests run in the same package.
func DefaultSysBenchScenarios(duration int, clientCounts []int, psModes []string) []SysBenchScenario {
	_ = "STUB: not implemented"
	return nil
}

// CaptureSysBenchEnvironment fills the EnvironmentInfo block for a sysbench
// run. PlanCacheBytes is the multigateway --plan-cache-memory default; we
// hardcode it because the shared cluster setup doesn't override it.
func CaptureSysBenchEnvironment(t *testing.T, binary string) EnvironmentInfo {
	_ = "STUB: not implemented"
	return *new(EnvironmentInfo)
}

// multigateway default; see go/services/multigateway/init.go

// gitHeadSha returns the short HEAD sha of the repository the test binary
// was built from, or "" if `git` isn't available.
func gitHeadSha() string { _ = "STUB: not implemented"; return "" }

// parseSysBenchInt reads an optional positive int env var with a default.
func parseSysBenchInt(t *testing.T, name string, def int) int { _ = "STUB: not implemented"; return 0 }

// ParseSysBenchClients reads SYSBENCH_CLIENTS as a comma-separated list of
// thread counts. Default: 1,8,32 — same as the handoff sweep.
func ParseSysBenchClients(t *testing.T) []int { _ = "STUB: not implemented"; return nil }

// ParseSysBenchDuration reads SYSBENCH_DURATION (seconds, default 60).
func ParseSysBenchDuration(t *testing.T) int { _ = "STUB: not implemented"; return 0 }

// ParseSysBenchPSModes reads SYSBENCH_PS_MODES; default is both
// {prepared,simple}. Unknown values fail the test loudly.
func ParseSysBenchPSModes(t *testing.T) []string { _ = "STUB: not implemented"; return nil }

// ParseSysBenchTargets reads SYSBENCH_TARGETS; default is multigateway only.
func ParseSysBenchTargets(t *testing.T) (postgres, multigateway, pgbouncer bool) {
	_ = "STUB: not implemented"
	return false, false, false
}

// ParseSysBenchQuiesce reads SYSBENCH_QUIESCE_SECONDS (default 5).
func ParseSysBenchQuiesce(t *testing.T) int { _ = "STUB: not implemented"; return 0 }
