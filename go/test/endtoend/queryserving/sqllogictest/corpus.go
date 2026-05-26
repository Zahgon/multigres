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
	"context"
	"regexp"
	"testing"
)

// Upstream sqllogictest corpus. The original is hosted in the SQLite fossil
// repository at https://www.sqlite.org/sqllogictest/ and carries an explicit
// multi-license (GPL / BSD / MIT / CC0 — "use whichever applies best") per
// its COPYRIGHT.md. gregrahn/sqllogictest is an actively-maintained git
// mirror of that fossil repo; pinning against it lets us shallow-clone via
// plain git with SHA-verified reproducibility.
//
// We consume this corpus under the MIT license per the upstream offer.
const (
	// CorpusRepoURL points at the git mirror we fetch from.
	CorpusRepoURL = "https://github.com/gregrahn/sqllogictest"

	// CorpusCommit pins the corpus at a specific revision so pass-rate
	// tracking over time is meaningful (the corpus contents don't drift
	// out from under the recorded baseline).
	CorpusCommit = "c67f97bf3ca7e590d12e073408bcacaf2ff0f3a0"

	// DefaultCorpusGlob matches every .test file in the upstream corpus.
	// Callers override via SLT_CORPUS_GLOB when they want to run a targeted
	// subset (useful for iteration).
	DefaultCorpusGlob = "test/**/*.test"

	defaultCacheRoot = "/tmp/multigres_slt_cache"
)

// resolveCorpusDir returns the directory containing the corpus to run.
//
// Default behaviour shallow-clones the pinned upstream at CorpusCommit into
// $SLT_CACHE_DIR/source/sqllogictest (reusing the cache across runs).
// Override with SLT_CORPUS_DIR to point at any local directory — useful
// for iterating against a subset or against an internal fork.
func resolveCorpusDir(t *testing.T, ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ensureUpstreamCorpus clones the pinned corpus if the cache is missing or
// points at a different SHA. Returns the absolute directory.
func ensureUpstreamCorpus(t *testing.T, ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// listCorpusFiles returns the .test / .slt files in the corpus directory
// matching SLT_CORPUS_GLOB (defaulting to DefaultCorpusGlob). Paths are
// absolute and sorted so per-file ordering is deterministic across runs.
//
// The glob uses doublestar semantics: "**" matches across path components,
// "*" matches within a single component, "?" matches a single non-/ char.
func listCorpusFiles(corpusDir string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// globToRegexp translates a shell glob into an anchored regexp. `**` matches
// across path separators, `*` matches within one segment, `?` matches a
// single non-`/` character. Regex metacharacters are escaped via
// regexp.QuoteMeta so regex syntax in the pattern stays literal.
//
// `a/**/b` also matches `a/b` (zero intermediate segments): a `/` immediately
// after `**` is consumed along with it.
func globToRegexp(pat string) (*regexp.Regexp, error) { _ = "STUB: not implemented"; return nil, nil }
