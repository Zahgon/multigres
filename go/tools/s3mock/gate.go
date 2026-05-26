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

package s3mock

import (
	"context"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"
)

// Matcher reports whether an S3 PUT operation should trigger the gate.
type Matcher func(bucket, key string) bool

// Hit describes the operation that tripped the gate.
type Hit struct {
	Method, Bucket, Key string
}

// Gate pauses S3 operations matching its Matcher once Arm has been called,
// records the first match in a Hit, and blocks the request goroutine until
// Release is called or the request context is cancelled.
type Gate struct {
	matcher Matcher
	armed   atomic.Bool

	mu      sync.Mutex
	hitCh   chan Hit
	hitOnce sync.Once
	relCh   chan struct{}
}

// NewGate returns a disarmed Gate.
func NewGate(m Matcher) *Gate { _ = "STUB: not implemented"; return nil }

// Arm enables the gate. Before arming, all operations pass through.
func (g *Gate) Arm() {
	_ = "STUB: not implemented"

	// Wait blocks until a matching operation arrives, or ctx is done.
	// Returns the Hit on success, or ctx.Err() on cancellation.
	return
}

func (g *Gate) Wait(ctx context.Context) (Hit, error) {
	_ = "STUB: not implemented"
	return *new(Hit), nil
}

// Release unblocks the held operation. Safe to call multiple times; the
// first call wins, subsequent calls are no-ops.
func (g *Gate) Release() { _ = "STUB: not implemented"; return }

// already closed

// Rearm disarms, then re-enables the gate for the next match. Resets the
// release channel and the one-shot hit. Caller is responsible for not
// racing this with in-flight Wait/Release on the same Gate.
func (g *Gate) Rearm() { _ = "STUB: not implemented"; return }

// callback returns a PutCallback that consults the gate.
func (g *Gate) callback() PutCallback { _ = "STUB: not implemented"; return *new(PutCallback) }

// Single-shot send so multiple matching PUTs do not stall behind a closed Wait.

// WithGate returns a ServerOption that installs g's callback. If a previous
// PutCallback was installed via WithPutCallback, it is replaced.
func WithGate(g *Gate) ServerOption { _ = "STUB: not implemented"; return *new(ServerOption) }

// MatchPut returns a Matcher that fires when the key contains substr.
func MatchPut(substr string) Matcher { _ = "STUB: not implemented"; return *new(Matcher) }

// MatchPutRegex returns a Matcher that fires when the key matches re.
func MatchPutRegex(re *regexp.Regexp) Matcher { _ = "STUB: not implemented"; return *new(Matcher) }

// Phase is a goroutine-safe string label that tests use to drive
// PhaseMatcher. Use descriptive labels ("bootstrap", "first-backup",
// "second-backup") so the s3mock log narrates the test as it runs.
type Phase struct {
	v atomic.Value
}

// Set updates the current phase.
func (p *Phase) Set(s string) {
	_ = "STUB: not implemented"

	// Get returns the current phase, or "" if Set has never been called.
	return
}

func (p *Phase) Get() string { _ = "STUB: not implemented"; return "" }

// PhaseMatcher returns a Matcher that delegates to base only while
// phase.Get() == want. It is the standard way to use a Gate in tests
// that need to discriminate between multiple actors hitting the same
// s3mock — for example, holding the first concurrent backup while
// letting the second pass through. Switch the phase to a different
// value to stop matching new requests; goroutines already parked
// inside gate.callback stay parked until ctx cancellation.
func PhaseMatcher(phase *Phase, want string, base Matcher) Matcher {
	_ = "STUB: not implemented"
	return *new(Matcher)
}

// MatchDataUpload fires on data-file PUTs inside a backup set's pg_data/
// (loose files like backup_label.zst) or bundle/ (pgBackRest's bundled small
// files). Excludes manifest writes and WAL archive PUTs.
//
// This is the only useful pause point for "kill primary postgres" fault
// scenarios: pgBackRest's first S3 write during a backup is to a bundle file,
// and pausing here is early enough that pg_stop_backup has not yet been
// called. Other pause points were considered (manifest writes at the end of
// a backup) but they happen after pg_stop_backup returns, so killing postgres
// at that point is too late to fail the backup.
var MatchDataUpload Matcher = func(_ string, key string) bool {
	if !strings.Contains(key, "/backup/") {
		return false
	}
	return strings.Contains(key, "/pg_data/") || strings.Contains(key, "/bundle/")
}

// MatchWalArchive fires on WAL archive PUTs (pgBackRest archive-push). The
// pgBackRest layout writes archived WAL segments under
// <stanza>/archive/<archive-id>/<timeline>-<wal-prefix>/<segment>, so the
// key contains "/archive/" and not "/backup/".
//
// Use to deterministically force a "WAL on the primary's local pg_wal but
// not yet pushed to the repo" condition for crash/durability tests.
var MatchWalArchive Matcher = func(_ string, key string) bool {
	if strings.Contains(key, "/backup/") {
		return false
	}
	return strings.Contains(key, "/archive/")
}
