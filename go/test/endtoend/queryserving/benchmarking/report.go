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
	"testing"
)

// WriteJSONReport writes the benchmark results as JSON.
func WriteJSONReport(t *testing.T, outputDir string, report *BenchmarkReport) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// WriteMarkdownReport generates a markdown comparison report.
func WriteMarkdownReport(t *testing.T, outputDir string, report *BenchmarkReport) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Group results by scenario type (sustained vs churn) and protocol.

// Determine which targets are present.

// Layout: metrics on rows, client counts on columns. With ~3 client counts
// and 3 targets × 3 metrics + overhead, this is much narrower than the
// transposed view (especially when pgbouncer is present).

// Header

// Per-target rows: TPS, Avg, P99.

// Overhead row: multigateway TPS vs postgres TPS, per client count.

// WriteSysBenchMarkdownReport renders a sysbench-shaped markdown report
// (per-scenario rows with a ps_mode column) under <outputDir>/benchmark-report.md.
//
// Kept separate from WriteMarkdownReport because the sysbench harness has a
// different scenario shape (ps_mode, no churn/protocol axis) and conflating
// the two renderers obscures both.
func WriteSysBenchMarkdownReport(t *testing.T, outputDir string, report *BenchmarkReport) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Render one section per ps_mode, with a row per (target, scenario).

// uniquePSModes returns ps_mode values in the order they first appear.
func uniquePSModes(results []ScenarioResult) []string { _ = "STUB: not implemented"; return nil }

// uniqueTargets returns deduplicated target names in the order they first appear.
func uniqueTargets(results []ScenarioResult) []string { _ = "STUB: not implemented"; return nil }

// uniqueClients returns deduplicated client counts in sorted order from results matching the filter.
func uniqueClients(results []ScenarioResult) []int { _ = "STUB: not implemented"; return nil }

// Results are already generated in order, so this preserves order.

// filterResults returns results matching the given churn mode and protocol.
func filterResults(results []ScenarioResult, churn bool, protocol string) []ScenarioResult {
	_ = "STUB: not implemented"
	return nil
}

// findResult locates a result for the given target and client count.
func findResult(results []ScenarioResult, target string, clients int) *ScenarioResult {
	_ = "STUB: not implemented"
	return nil
}
