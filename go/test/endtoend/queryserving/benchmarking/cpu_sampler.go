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
	"sync"
	"testing"
	"time"
)

// CPUStats summarizes CPU usage samples for a single process tree
// across a scenario run. Values are in "percent of one core" — i.e.
// 100.0 means 1 core fully busy. With many backends it can exceed 100.
type CPUStats struct {
	Label       string    `json:"label"`
	Samples     int       `json:"samples"`
	MeanPercent float64   `json:"mean_percent"`
	MinPercent  float64   `json:"min_percent"`
	MaxPercent  float64   `json:"max_percent"`
	Series      []float64 `json:"series_percent"`
}

// pidEnumerator returns the set of PIDs to sample at a given moment.
type pidEnumerator func() ([]int, error)

// staticPidEnumerator returns a constant list of PIDs.
func staticPidEnumerator(pids ...int) pidEnumerator {
	_ = "STUB: not implemented"
	return *new(pidEnumerator)
}

// postgresPidEnumerator returns the postmaster pid plus its direct children
// (backends, bgworker, walwriter, etc.) on each call. The postmaster pid is
// read from postmaster.pid in the data dir, so it's resilient to postgres
// restarts mid-scenario.
func postgresPidEnumerator(dataDir string) pidEnumerator {
	_ = "STUB: not implemented"
	return *new(pidEnumerator)
}

// First line is the postmaster PID.

// Children (one level is enough for postgres backends).

// pgrep returns exit 1 when no children; treat as just the postmaster.

// readPSCPU returns the sum of %cpu across the given pids using a single
// `ps` invocation. macOS ps reports "current" CPU% (decayed average).
// On Linux ps reports cumulative-since-start which is less useful for
// short windows, but we sample frequently enough that the per-sample
// granularity matters more than the absolute scale.
func readPSCPU(pids []int) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

// Errors are intentionally swallowed: pids may exit between
// enumeration and sampling, in which case `ps` returns non-zero but
// still emits %cpu lines for the survivors. Sum what we got.

// cpuSampler runs a goroutine sampling the given pidEnumerator at the
// configured interval. Stop() returns the aggregated stats once the
// goroutine has finished collecting.
type cpuSampler struct {
	label    string
	enum     pidEnumerator
	interval time.Duration

	wg     sync.WaitGroup
	cancel context.CancelFunc
	mu     sync.Mutex
	series []float64
}

func newCPUSampler(label string, enum pidEnumerator, interval time.Duration) *cpuSampler {
	_ = "STUB: not implemented"
	return nil
}

// Start begins sampling. Cancel ctx (or call Stop) to finish.
func (s *cpuSampler) Start(parent context.Context) { _ = "STUB: not implemented"; return }

// Stop finalizes the sampler and returns aggregated stats.
func (s *cpuSampler) Stop() CPUStats { _ = "STUB: not implemented"; return *new(CPUStats) }

// writeCPUStatsFile persists per-process CPU stats for one (scenario,target)
// run as JSON under <outputDir>/cpu/<scenario>/<target>.json.
func writeCPUStatsFile(t *testing.T, outputDir, scenario, target string, stats map[string]CPUStats) {
	_ = "STUB: not implemented"
	return
}

// writeSysBenchCPUStatsFile persists per-process CPU stats under
// <outputDir>/cpu/<scenario>/<target>/processes.json. The "processes.json"
// filename matches the vitess sysbench bundle's layout so cross-stack tools
// can read either; the per-target directory keeps multi-target runs (e.g.
// SYSBENCH_TARGETS=postgres,multigateway) from clobbering each other.
func writeSysBenchCPUStatsFile(t *testing.T, outputDir, scenario, target string, stats map[string]CPUStats) {
	_ = "STUB: not implemented"
	return
}
