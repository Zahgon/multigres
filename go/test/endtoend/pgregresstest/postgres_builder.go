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

package pgregresstest

import (
	"context"
	"testing"
	"time"

	"github.com/multigres/multigres/go/test/endtoend/pgbuilder"
	"github.com/multigres/multigres/go/tools/executil"

	// PostgreSQL driver for the diagnostic / shim install path.
	_ "github.com/lib/pq"
)

// Re-export pgbuilder constants so existing callers and scripts that reference
// pgregresstest.PostgresVersion / PostgresGitRepo keep working unchanged.
const (
	PostgresGitRepo  = pgbuilder.PostgresGitRepo
	PostgresVersion  = pgbuilder.PostgresVersion
	PostgresCacheDir = pgbuilder.PostgresCacheDir
)

// PostgresBuilder wraps pgbuilder.Builder with regression/isolation-suite
// specific helpers (run tests, parse TAP output, render reports).
//
// The source/build/install/cleanup logic lives in pgbuilder and is also used
// by the sqllogictest differential harness.
type PostgresBuilder struct {
	*pgbuilder.Builder
}

// TestResults contains the results from running PostgreSQL regression tests.
type TestResults struct {
	TotalTests     int
	PassedTests    int
	FailedTests    int
	SkippedTests   int
	TimedOut       bool // true if the suite was killed by context timeout
	Duration       time.Duration
	FailureDetails []TestFailure
	Tests          []IndividualTestResult // Per-test results in order
}

// IndividualTestResult represents a single test's result.
type IndividualTestResult struct {
	Name     string `json:"name"`
	Status   string `json:"status"` // "pass", "fail", "skip"
	Duration string `json:"duration"`
	// PatchApplied is true when a per-test patch (testdata/pg17/patches/<name>.patch)
	// was applied to the upstream expected output before diffing. This signals
	// that multigres's output intentionally differs from stock PostgreSQL for
	// this test (typically error-message wording) and the difference is
	// documented in the patch file.
	PatchApplied bool `json:"patch_applied,omitempty"`
	// PatchPath is the repo-relative path to the patch file (when applied),
	// so status reports can link back to it.
	PatchPath string `json:"patch_path,omitempty"`
	// FailReason is a short description populated when Status == "fail" and
	// the failure comes from the patch-verification step (e.g. "patch did not
	// apply" or "actual output does not match patched expected").
	FailReason string `json:"fail_reason,omitempty"`
}

// TestFailure represents a single test failure.
type TestFailure struct {
	TestName string
	Error    string
}

// NewPostgresBuilder creates a new PostgresBuilder with unique build directories.
func NewPostgresBuilder(t *testing.T) *PostgresBuilder { _ = "STUB: not implemented"; return nil }

// CheckBuildDependencies verifies that required build tools are available.
func CheckBuildDependencies(t *testing.T) error { _ = "STUB: not implemented"; return nil }

// testSuiteConfig holds configuration for running a PostgreSQL test suite
// via the shared runTestSuite helper.
type testSuiteConfig struct {
	suiteName string // for log messages, e.g. "Regression" or "Isolation"
	outputDir string // where to copy regression.out and regression.diffs
	srcOutDir string // build directory containing regression.out and regression.diffs
}

// runTestSuite executes a pre-built test command and handles result parsing,
// artifact copying, and failure reporting. Both RunRegressionTests and
// RunIsolationTests delegate to this after constructing their command.
func (pb *PostgresBuilder) runTestSuite(t *testing.T, ctx context.Context, cmd *executil.Cmd, cfg testSuiteConfig, multigatewayPort int, password string) (*TestResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cap how long Run() waits for I/O after the process exits. Without this,
// grandchildren (e.g. psql) holding pipes open would cause a hang.

// pg_regress expected .out files were captured against vanilla PostgreSQL
// defaults, so the planner picks different shapes (e.g. Hash↔Merge,
// Bitmap↔Index) when pgctld's tuned GUCs are in effect. Override to
// PostgreSQL defaults per session via PGOPTIONS so the harness matches
// upstream fixtures without changing pgctld defaults.

// PG_TEST_TIMEOUT_DEFAULT (in seconds) caps how long isolationtester
// waits per step before cancelling. Upstream default is 180 →
// max_step_wait = 360 s with hard-exit at 720 s, which lets a
// single compat-incompatible spec burn ~12 minutes of the suite
// ctx and starve the rest of the schedule. The cap applies per
// step and a multi-permutation spec where every permutation
// hangs scales linearly. 5 s gives 10 s cancel / 20 s hard-exit
// per step, ~500x headroom over the 10 ms poll interval used to
// detect legitimate blocking.

// Capture stdout for result parsing while still printing to the terminal.
// pg_regress deletes regression.out when all tests pass, so stdout is our
// reliable source for TAP output. We still try the on-disk file first
// because it contains partial results if the process is killed mid-run.

// RunRegressionTests runs PostgreSQL regression tests against multigateway.
//
// Uses make installcheck-tests with TESTS variable to run specific regression tests
// against the existing PostgreSQL server (multigateway).
//
// The installcheck-tests target runs specific tests against an already-running
// PostgreSQL server, unlike installcheck which runs the entire parallel_schedule.
//
// From PostgreSQL's src/test/regress/GNUmakefile:
//
//	installcheck: runs --schedule=parallel_schedule (all tests)
//	installcheck-tests: runs $(TESTS) (specific tests only)
//
// Environment variables that pg_regress reads:
//
//	PGHOST, PGPORT, PGUSER, PGPASSWORD, PGDATABASE - connection params
//
// Reference: https://github.com/postgres/postgres/blob/master/src/test/regress/GNUmakefile
func (pb *PostgresBuilder) RunRegressionTests(t *testing.T, ctx context.Context, multigatewayPort int, password string) (*TestResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// --use-existing: skip pg_regress's automatic DROP + CREATE of the
// "regression" database. Multigateway rejects DROP DATABASE (see the
// unsafe-statement list in go/services/multigateway/planner/unsafe_stmt.go),
// so we can't let pg_regress manage the database. Instead we run against
// the existing "postgres" database for the whole suite.
//
// --dbname=postgres: point pg_regress at that existing database. pg_regress
// will create/drop the expected schema objects per test; cross-test state
// leakage is still possible but has not surfaced in practice.

// ParseTestResults parses pg_regress TAP output to extract test results.
// Returns an error if no TAP-formatted lines are found in the output.
func (pb *PostgresBuilder) ParseTestResults(output string) (*TestResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Parse TAP format output from pg_regress
// Example lines:
//   ok 1         - test_setup                                178 ms  (serial)
//   ok 2         + boolean                                    61 ms  (parallel)
//   not ok 3     + char                                       39 ms  (parallel)

// PatchesDir returns the absolute path to the per-test patch directory for
// the PostgreSQL major version this builder is pinned to. Patches live under
// testdata/pg<major>/patches/ next to this file.
func PatchesDir() string { _ = "STUB: not implemented"; return "" }

// pgMajorDir returns the directory name under testdata/ that holds patches for
// the PostgreSQL version this test targets. Pinned to pg17 today; add a case
// when PostgresVersion advances.
func pgMajorDir() string {
	_ = "STUB: not implemented"
	// "REL_17_6" → "pg17"
	return ""
}

// Unknown version: fall back to pg17 and let the test fail if patches
// are missing.

// findRepoRoot walks up from this source file until it finds a directory
// containing go.mod. Returns empty string if not found — callers degrade to
// absolute paths in that case.
func findRepoRoot() string { _ = "STUB: not implemented"; return "" }

// findExpectedFile locates the expected-output file pg_regress would use for
// the given test name. PostgreSQL ships variant files (name_1.out, name_2.out,
// …) for platform-dependent output. We try the canonical name first, then
// numbered variants in order. Returns the empty string if none exist.
func findExpectedFile(regressDir, name string) string { _ = "STUB: not implemented"; return "" }

// VerifyWithPatches re-evaluates each test's pass/fail status using the
// patch-based pipeline (see patch_verify.go). After pg_regress runs, this
// ignores pg_regress's own pass/fail verdict (which is a strict text diff)
// and replaces it with: does the actual output match the (patched) expected
// output? Results are updated in-place, including aggregate counters.
//
// In generate mode, any residual diffs are absorbed by (re)writing patches.
//
// Expected output lives in the source tree (prep_buildtree does not symlink
// .out files into the build tree). Actual output is written by pg_regress
// into the build tree's results/ directory.
func (pb *PostgresBuilder) VerifyWithPatches(t *testing.T, ctx context.Context, results *TestResults, buildRegressDir, outputDir string) error {
	_ = "STUB: not implemented"
	return nil
}

// Ensure patch dir exists in generate mode so writes don't fail.

// Per-test residual diffs are written under outputDir/diffs/ for inclusion
// in the CI artifact. Concatenated failures.diffs is written at the end.

// Recompute aggregates from the per-test results after verification.
// We intentionally discard pg_regress's TAP-derived aggregates because
// patch-based verification is authoritative.

// Leave skipped tests alone.

// Infrastructure problem (test didn't run, expected missing).
// Preserve TAP verdict, count accordingly.

// Preserve SkippedTests + any pre-existing total if it exceeds ran.

// CountScheduleTests parses a PostgreSQL schedule file and returns the number
// of tests listed. Each line starting with "test:" contains space-separated
// test names (parallel groups have multiple tests per line).
func CountScheduleTests(scheduleFile string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// SuiteResult holds results for one test suite.
type SuiteResult struct {
	Name          string       // e.g. "Regression Tests", "Isolation Tests"
	Results       *TestResults // parsed TAP results
	ExpectedTests int          // total tests from schedule file (0 = unknown)
}

// githubBlobURLPrefix returns an absolute URL prefix of the form
// "https://github.com/<owner>/<repo>/blob/<sha>/" when running inside a
// GitHub Actions job, or an empty string otherwise. Repo-relative paths
// concatenated onto this prefix resolve to the blob view of that file at
// the exact commit the job is executing against.
func githubBlobURLPrefix() string { _ = "STUB: not implemented"; return "" }

// WriteMarkdownSummary generates a unified markdown report covering one or more
// test suites. It writes the report to pb.OutputDir/compatibility-report.md and
// appends it to GITHUB_STEP_SUMMARY when running in CI.
//
// Each suite gets its own badge showing pass rate and timeout status. Full diffs
// are available in the CI artifact (regression.diffs).
func (pb *PostgresBuilder) WriteMarkdownSummary(t *testing.T, suites []SuiteResult) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Build an absolute URL prefix for patch links when running in CI.
// Without this, markdown like [patch](go/test/.../foo.patch) is resolved
// relative to the step-summary page (/actions/runs/<id>/) and produces a
// 404 URL like .../actions/runs/go/test/.../foo.patch. Falls back to a
// relative link when the GitHub Actions env vars aren't set (local runs).

// jsonSuiteResult is the JSON-serializable representation of a single test suite's results.
type jsonSuiteResult struct {
	Name  string                 `json:"name"`
	Tests []IndividualTestResult `json:"tests"`
}

// WriteJSONResults serializes suite results to pb.OutputDir/results.json.
// This file is consumed by CI scripts that compare runs to detect regressions.
func (pb *PostgresBuilder) WriteJSONResults(t *testing.T, suites []SuiteResult) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// patchIsolationtester rewrites two pieces of
// src/test/isolation/isolationtester.c so the harness works against
// multigateway:
//
//  1. Per-session application_name setup. Upstream sends
//     PQexecParams("SELECT set_config('application_name',
//     current_setting('application_name') || '/' || $1, false)", ...),
//     which the multigateway planner rejects (the value must be a literal
//     constant for is_local=false set_config so the pooler can track it).
//     The replacement reads PGAPPNAME (set by isolation_main.c per test),
//     concatenates the session name client-side, escapes via
//     PQescapeLiteral, and sends a simple-protocol PQexec.
//
//  2. Lock-wait probe function name. Upstream prepares
//     `pg_catalog.pg_isolation_test_session_is_blocked(...)`. Replacing
//     that builtin C function with a PL/pgSQL shim via CREATE OR REPLACE
//     proved unreliable (fresh backends were observed to still bind the
//     C entry, returning false for every probe and hanging every
//     blocking spec at max_step_wait). We point the harness at our own
//     function `public.multigres_test_session_is_blocked` (installed by
//     installPIDMappingFunction) with explicit arg casts so PG resolves
//     it under extended-protocol Parse with paramTypes=NULL.
//
// Idempotent: source is reset via `git checkout` before patching, so
// repeat invocations against the cached checkout produce the same
// result.
func (pb *PostgresBuilder) patchIsolationtester(t *testing.T, ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Add explicit type casts on both args. isolationtester PQprepares the
// wait-query with paramTypes=NULL so $1 enters parse as UNKNOWN, and the
// '{...}' literal is also UNKNOWN. PG resolves this for the pg_catalog
// C builtin via implicit catalog priority but fails for a public
// PL/pgSQL function with "function public.X(unknown, unknown) does not
// exist". Casting to int4 / int4[] removes the ambiguity.

// BuildIsolation builds the PostgreSQL isolation test tools (isolationtester and
// pg_isolation_regress). Must be called after Build().
func (pb *PostgresBuilder) BuildIsolation(t *testing.T, ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// truncateForLog clips s to at most n characters (with an ellipsis suffix when
// truncation occurs). Used for compact log/error messages.
func truncateForLog(s string, n int) string { _ = "STUB: not implemented"; return "" }

// installPIDMappingFunction creates public.multigres_test_session_is_blocked
// in the target database so the patched isolationtester (see
// patchIsolationtester) can probe lock waits through the multigateway →
// multipooler → PostgreSQL hop.
//
// Multipooler is configured with --database=postgres and routes every
// query to a pooled connection against the postgres DB regardless of the
// dbname in the client startup packet, so the shim must live in postgres.
// Both isolation invocation paths (selective via PGISOLATION_TESTS and
// full-suite via the make installcheck target) force --dbname=postgres on
// pg_isolation_regress, so postgres is also the dbname the harness opens.
//
// The shim mirrors the upstream builtin: returns true if check_pid is
// waiting on any pid in blocked_by, considering both heavyweight lock
// waits (pg_blocking_pids) and SSI safe-snapshot waits
// (pg_safe_snapshot_blocking_pids — required for SERIALIZABLE READ ONLY
// DEFERRABLE specs such as read-only-anomaly-3). Both inputs are
// multigateway virtual pids; we map them to real PostgreSQL backend pids
// via pg_stat_activity.application_name (the multipooler stamps each
// backend with `multigres_vpid:<id>` per query). A given vpid can map to
// multiple PG backends in flight (a leftover stamp on a pool conn after
// a regular query, plus the live reserved conn) so the wait-check
// aggregates over every matching backend rather than picking one
// non-deterministically.
func (pb *PostgresBuilder) installPIDMappingFunction(t *testing.T, pgPort int, password string) error {
	_ = "STUB: not implemented"
	return nil
}

// Debug table: every shim invocation logs its inputs/outputs and a
// snapshot of every backend in this DB so failures can be diagnosed
// post-hoc by querying isolation_debug_log (see
// dumpIsolationDebugLog).

// Non-destructive add for runs against a pre-existing table from an
// earlier shim version that lacked all_pg_backends.

// Sanity check: the function exists and is plpgsql.

// dumpIsolationDebugLog prints recent entries from
// public.isolation_debug_log so investigators can see the inputs/outputs
// of every shim invocation during the isolation run. Best-effort;
// failures are logged and ignored.
func (pb *PostgresBuilder) dumpIsolationDebugLog(t *testing.T, pgPort int, password string) {
	_ = "STUB: not implemented"
	return
}

// RunIsolationTests runs PostgreSQL isolation tests against multigateway.
// Isolation tests exercise multi-connection concurrency (deadlocks, serialization
// anomalies, lock contention, concurrent DDL) using isolationtester.
//
// directPgPort is the primary's direct PostgreSQL port; it's used to install
// the public.multigres_test_session_is_blocked shim that the patched
// isolationtester binary calls (see patchIsolationtester for the source-side
// rewrite that retargets the wait query at the public function).
//
// The isolation Makefile has no installcheck-tests target, so for selective tests
// we invoke pg_isolation_regress directly with test names as positional args.
func (pb *PostgresBuilder) RunIsolationTests(t *testing.T, ctx context.Context, multigatewayPort, directPgPort int, password string) (*TestResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Install the lock-detection shim on PostgreSQL directly (bypassing
// multigateway). Both the selective (PGISOLATION_TESTS) and full-suite
// paths force --dbname=postgres on pg_isolation_regress (see the cmd
// construction below), and multipooler routes every query to the
// postgres DB anyway, so the shim only needs to live there.

// Post-suite diagnostic: dump the last entries of isolation_debug_log
// so investigators can see what the shim observed (or didn't) for
// hung specs. The table lives in the postgres DB on the primary;
// query it directly to bypass any multigateway routing that a
// failing wait-query would have used.
