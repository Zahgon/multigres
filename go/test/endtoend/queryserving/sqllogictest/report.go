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

package sqllogictest

import (
	"testing"
	"time"
)

// fileReport is the per-file record in the serialized report. It captures
// the file path, how each target fared, and basic counters callers can
// aggregate. The field layout intentionally mirrors pgregresstest's
// IndividualTestResult so the existing .github/scripts/detect-regressions.sh
// can consume our output with minimal special-casing.
type fileReport struct {
	Name     string `json:"name"`           // relative path, e.g. "test/select1.test"
	Status   string `json:"status"`         // "pass" | "fail"
	Duration string `json:"duration"`       // combined duration across targets
	Postgres perRun `json:"postgres"`       // baseline (direct PG)
	Gateway  perRun `json:"multigateway"`   // candidate (multigres)
	Note     string `json:"note,omitempty"` // human note (timeout, exec err, divergence)
}

// perRun is one target's outcome for a single file.
type perRun struct {
	Passed   bool   `json:"passed"`
	TimedOut bool   `json:"timed_out,omitempty"`
	Duration string `json:"duration"`
	Output   string `json:"output,omitempty"`
	ExecErr  string `json:"exec_err,omitempty"`
}

// suiteReport is what we serialize. Fields line up with pgregresstest so
// downstream tooling can treat the two identically. One suiteReport is
// produced per wire protocol (simple / extended) so regression tracking
// in detect-regressions.sh is scoped per protocol.
type suiteReport struct {
	Name          string       `json:"name"` // "SQLLogicTest-simple" / "SQLLogicTest-extended"
	CorpusDir     string       `json:"corpus_dir"`
	CorpusCommit  string       `json:"corpus_commit"`
	TimedOut      bool         `json:"timed_out"`
	TotalFiles    int          `json:"total_files"`
	PassedBoth    int          `json:"passed_both"`
	PassedPGOnly  int          `json:"passed_pg_only"`
	PassedMGOnly  int          `json:"passed_mg_only"`
	FailedBoth    int          `json:"failed_both"`
	PGPassed      int          `json:"postgres_passed"`
	GatewayPassed int          `json:"multigateway_passed"`
	StartedAt     time.Time    `json:"started_at"`
	Duration      string       `json:"duration"`
	Tests         []fileReport `json:"tests"`
}

// newSuiteReport builds one suite report from paired per-file results.
// pgResults and mgResults must align on File path.
func newSuiteReport(name, corpusRoot string, pgResults, mgResults []*runResult, startedAt time.Time, timedOut bool) *suiteReport {
	_ = "STUB: not implemented"
	return nil
}

func toPerRun(r *runResult) perRun { _ = "STUB: not implemented"; return *new(perRun) }

// Only keep captured output for failing runs — passing files produce a
// one-line success banner we don't need in results.json.

// writeJSON writes the merged suite list to <outputDir>/results.json as a
// pretty-printed array of suite objects. The array shape is exactly what
// .github/scripts/detect-regressions.sh expects, so regression tracking
// is cross-run-portable per protocol.
func writeJSON(outputDir string, reports []*suiteReport) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// writeMarkdownSummary emits one section per suite (i.e. per protocol),
// with shields.io badges, counters, and the proxy-divergence list. The
// summary is mirrored to GITHUB_STEP_SUMMARY when set so the CI job page
// shows pass rates without downloading artifacts.
func writeMarkdownSummary(t *testing.T, outputDir string, reports []*suiteReport) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func firstLine(s string) string { _ = "STUB: not implemented"; return "" }

// logSummary dumps a compact summary of one suite to the test log.
func (r *suiteReport) logSummary(t *testing.T) { _ = "STUB: not implemented"; return }
