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

// Package pgbuilder provides reusable helpers for checking out, building, and
// installing a pinned PostgreSQL source tree from git for end-to-end tests.
//
// Both the PostgreSQL regression/isolation harness (pgregresstest) and the
// sqllogictest differential harness (queryserving/sqllogictest) consume this
// package so they run against the same PostgreSQL version.
package pgbuilder

import (
	"context"
	"testing"
)

const (
	// PostgresGitRepo is the official PostgreSQL git repository.
	PostgresGitRepo = "https://github.com/postgres/postgres"

	// PostgresVersion is the git tag to checkout.
	PostgresVersion = "REL_17_6"

	// PostgresCacheDir is the default cache directory for PostgreSQL source and builds.
	PostgresCacheDir = "/tmp/multigres_pg_cache"
)

// Builder manages a PostgreSQL source checkout, an isolated build, and a
// per-test install prefix. It does not run any server processes itself;
// callers either invoke the built binaries directly or compose Builder with
// higher-level helpers (see Standalone in this package).
type Builder struct {
	// SourceDir is the shared source checkout, reused across runs.
	SourceDir string
	// BuildDir is the per-invocation build directory.
	BuildDir string
	// InstallDir is the per-invocation install prefix (contains bin/, lib/, share/).
	InstallDir string
	// OutputDir is a persistent per-invocation directory for caller-written artifacts
	// (reports, diffs, etc.). pgbuilder itself does not write here.
	OutputDir string
}

// New returns a Builder with unique per-invocation build and install directories
// rooted at $MULTIGRES_PG_CACHE_DIR (or /tmp/multigres_pg_cache when unset).
// Multiple concurrent callers get distinct build/install trees but share the
// source checkout.
func New(t *testing.T) *Builder { _ = "STUB: not implemented"; return nil }

// BinDir is the directory containing the built PostgreSQL binaries (postgres,
// initdb, psql, pg_ctl, ...).
func (b *Builder) BinDir() string { _ = "STUB: not implemented"; return "" }

// CheckBuildDependencies verifies that required C toolchain is available on
// the host. Callers that depend on building PostgreSQL from source should
// invoke this early and skip the test on a clear error message.
func CheckBuildDependencies(t *testing.T) error { _ = "STUB: not implemented"; return nil }

// EnsureSource ensures the pinned PostgreSQL source tree is available,
// cloning it if missing or wrong version. The source is shared across
// concurrent builders.
func (b *Builder) EnsureSource(t *testing.T, ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Build runs ./configure + make + make install into b.InstallDir.
// ICU is disabled so the build does not require icu4c headers; that matches
// what the existing pgregresstest suite already does.
func (b *Builder) Build(t *testing.T, ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Cleanup removes per-invocation build and install artifacts but leaves the
// shared source checkout in place so subsequent runs skip the clone.
func (b *Builder) Cleanup() { _ = "STUB: not implemented"; return }
