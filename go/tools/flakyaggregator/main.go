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

// Command flakyaggregator collects JUnit test artifacts produced by recent
// CI runs on main, counts failure occurrences per test, and posts a Slack
// summary of tests with more than --threshold failures in the last
// --days days.
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/google/go-github/v68/github"
)

const (
	defaultDays       = 7
	defaultThreshold  = 3
	defaultWorkers    = 8 // per-run fetch parallelism; well below GitHub's secondary concurrency limits
	httpTimeoutSec    = 60
	maxRunsToInspect  = 2000 // safety cap (covers main + all PR branches over a week)
	maxRetries        = 3
	maxBackoff        = 5 * time.Minute
	testArtifactGlob1 = "-test-logs"
)

// testWorkflowFiles enumerates the workflow files that run gotestsum with
// --rerun-fails and can therefore produce flake signal. Listing runs scoped
// per-file (instead of across the whole repo) is the largest single rate-
// limit saver: it skips lint, docker-build, pgbench, etc. runs that have
// nothing for us. coverage.yml is intentionally excluded because it
// deliberately omits --rerun-fails (a retry would truncate -coverprofile),
// so its event stream contains no fail-then-pass pairs to detect.
//
// Update this list when a new test workflow with --rerun-fails is added.
var testWorkflowFiles = []string{
	"test-short.yml",
	"test-race.yml",
	"test-integration.yml",
}

func main() {
	days := flag.Int("days", defaultDays, "lookback window in days")
	threshold := flag.Int("threshold", defaultThreshold, "report tests with strictly more failures than this")
	workers := flag.Int("workers", defaultWorkers, "per-run fetch parallelism")
	dryRun := flag.Bool("dry-run", false, "print Slack payload to stdout instead of posting")
	flag.Parse()

	if *workers < 1 {
		log.Fatalf("--workers must be >= 1, got %d", *workers)
	}

	repo := os.Getenv("GITHUB_REPOSITORY")
	token := os.Getenv("GITHUB_TOKEN")
	slackURL := os.Getenv("SLACK_WEBHOOK_URL")

	if repo == "" {
		log.Fatal("GITHUB_REPOSITORY is required")
	}
	if token == "" {
		log.Fatal("GITHUB_TOKEN is required")
	}
	if !*dryRun && slackURL == "" {
		log.Fatal("SLACK_WEBHOOK_URL is required (use --dry-run to skip posting)")
	}
	owner, name, ok := strings.Cut(repo, "/")
	if !ok {
		log.Fatalf("GITHUB_REPOSITORY must be owner/name, got %q", repo)
	}

	//nolint:gocritic // main is the program entry point; context.Background() is the canonical root.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	gh := github.NewClient(nil).WithAuthToken(token)

	since := time.Now().UTC().AddDate(0, 0, -*days)
	log.Printf("scanning runs across all branches since %s for repo %s", since.Format(time.RFC3339), repo)
	logRateLimit(ctx, gh, "start")

	runs, err := listFlakeCandidateRuns(ctx, gh, owner, name, since)
	if err != nil {
		log.Fatalf("list workflow runs: %v", err)
	}
	log.Printf("found %d completed-success runs (flake-candidate scope)", len(runs))

	stats := newAggregator()
	processRunsConcurrently(ctx, gh, owner, name, runs, stats, *workers)

	flaky := stats.filter(*threshold)
	payload := formatSlackPayload(flaky, *days, *threshold)
	logRateLimit(ctx, gh, "end")

	if *dryRun {
		buf, _ := json.MarshalIndent(payload, "", "  ")
		fmt.Println(string(buf))
		return
	}

	if err := postSlack(ctx, slackURL, payload); err != nil {
		log.Fatalf("post slack: %v", err)
	}
	log.Printf("posted Slack message with %d flaky test(s)", len(flaky))
}

// processRunsConcurrently fans out per-run artifact fetch+parse work across
// a small worker pool. The aggregator is shared across workers; recordFailure
// is mutex-guarded so concurrent ingest is safe. Per-run errors are logged
// and skipped — they're independent and should not abort the whole run.
func processRunsConcurrently(ctx context.Context, gh *github.Client, owner, repo string, runs []*github.WorkflowRun, stats *aggregator, workers int) {
	_ = "STUB: not implemented"
	return
}

func processRun(ctx context.Context, gh *github.Client, owner, repo string, run *github.WorkflowRun, stats *aggregator) {
	_ = "STUB: not implemented"
	return
}

// --- Rate-limit handling ---

// withRetry runs fn, sleeping and retrying on go-github's rate-limit and
// abuse errors up to maxRetries times. RateLimitError carries the reset
// timestamp; AbuseRateLimitError carries an explicit Retry-After. We cap
// each sleep at maxBackoff so a misbehaving response can't stall the whole
// 30-minute run.
func withRetry[T any](ctx context.Context, label string, fn func() (T, *github.Response, error)) (T, *github.Response, error) {
	_ = "STUB: not implemented"
	return *new(T), nil, nil
}

func logRateLimit(ctx context.Context, gh *github.Client, label string) {
	_ = "STUB: not implemented"
	return
}

// --- GitHub API (thin wrappers around go-github) ---

// listFlakeCandidateRuns returns recent successful workflow runs across all
// branches, scoped to the workflow files listed in testWorkflowFiles. The
// per-file scoping is the largest single rate-limit saver: it skips lint,
// docker-build, etc. runs that have no JUnit artifacts. The conclusion=
// success filter is what makes it safe to include PR branches — if a test
// failed and the run still concluded success, gotestsum's --rerun-fails
// recovered, which is the textbook flake signal. Tests that fail every
// retry make the run conclude failure and are correctly excluded as
// real-bug noise rather than counted as flakes.
func listFlakeCandidateRuns(ctx context.Context, gh *github.Client, owner, repo string, since time.Time) ([]*github.WorkflowRun, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func listRunArtifacts(ctx context.Context, gh *github.Client, owner, repo string, runID int64) ([]*github.Artifact, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func downloadArtifact(ctx context.Context, gh *github.Client, owner, repo string, artifactID int64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The redirect URL points to GitHub's blob storage; fetching it does NOT
// count against the REST rate limit, so no withRetry wrapper here.

// --- Aggregation ---

// isTestArtifact returns true for artifact names that contain the gotestsum
// JSONL streams we feed into the flake detector. coverage-* is intentionally
// not matched: those workflows omit --rerun-fails (a retry would truncate
// -coverprofile), so their event streams contain no fail-then-pass pairs.
func isTestArtifact(name string) bool { _ = "STUB: not implemented"; return false }

type testKey struct {
	pkg  string
	name string
}

type testStats struct {
	// runIDs tracks distinct workflow runs in which this test was observed
	// failing — a test that fails twice in one run only counts once.
	runIDs     map[int64]struct{}
	lastSeen   time.Time
	lastRunURL string
}

type aggregator struct {
	mu    sync.Mutex
	tests map[testKey]*testStats
}

func newAggregator() *aggregator { _ = "STUB: not implemented"; return nil }

// recordFailure is safe for concurrent use; ingestArtifactZip is called from
// multiple worker goroutines.
func (a *aggregator) recordFailure(run *github.WorkflowRun, key testKey) {
	_ = "STUB: not implemented"
	return
}

type flakyEntry struct {
	Pkg        string
	Name       string
	Failures   int
	LastSeen   time.Time
	LastRunURL string
}

func (a *aggregator) filter(threshold int) []flakyEntry { _ = "STUB: not implemented"; return nil }

func (a *aggregator) ingestArtifactZip(run *github.WorkflowRun, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Defensive: skip any path traversal / nested files we don't expect.

// --- gotestsum JSONL parsing ---

// testEvent matches the subset of `go test -json` event fields we care about.
// The full event format also carries Time, Elapsed, and Output, but for
// flake detection we only need to know which test moved through which
// terminal state in which order.
type testEvent struct {
	Action  string `json:"Action"`
	Package string `json:"Package"`
	Test    string `json:"Test"`
}

// parseFlakesFromJSONL streams a gotestsum --jsonfile output and returns
// the set of tests that were flaky in this run. A test is flaky iff its
// event stream contains both a "fail" and a "pass" action — gotestsum's
// --rerun-fails appends the retry's events, so a recovered flake is
// recognizable as fail-followed-by-pass on the same (package, test) pair.
//
// Package-level events (Test == "") are ignored: they reflect the package's
// rolled-up status and would double-count tests that are already represented
// by their own events.
func parseFlakesFromJSONL(r io.Reader) ([]testKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// some events embed long Output strings

// tolerate occasional malformed lines rather than abort

// --- Slack ---

type slackPayload struct {
	Text string `json:"text"`
}

func formatSlackPayload(flaky []flakyEntry, days, threshold int) slackPayload {
	_ = "STUB: not implemented"
	return *new(slackPayload)
}

func postSlack(ctx context.Context, webhookURL string, p slackPayload) error {
	_ = "STUB: not implemented"
	return nil
}
