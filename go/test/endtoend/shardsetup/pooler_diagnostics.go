// Copyright 2026 Supabase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package shardsetup

import (
	"testing"
	"time"

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	multipoolermanagerdatapb "github.com/multigres/multigres/go/pb/multipoolermanagerdata"
)

// checkPoolerCondition evaluates condition for each pooler once. Returns nil if all
// poolers satisfy the condition, or a slice of failure lines (one per failing pooler)
// with the reason and FormatPoolerDiagnostics output.
func checkPoolerCondition(t *testing.T, poolers []*MultipoolerInstance, condition func(r PoolerStatusResult) (bool, string)) []string {
	_ = "STUB: not implemented"
	return nil
}

// PoolerStatusResult holds the fetched status for one pooler instance.
// Status is nil and Err is non-nil if the fetch failed.
type PoolerStatusResult struct {
	Name            string
	Status          *multipoolermanagerdatapb.Status
	ConsensusStatus *clustermetadatapb.ConsensusStatus
	Err             error
}

// fetchPoolerStatuses fetches the status of each pooler in order, returning a slice
// that preserves the input ordering (no map iteration randomness).
func fetchPoolerStatuses(t *testing.T, poolers []*MultipoolerInstance) []PoolerStatusResult {
	_ = "STUB: not implemented"
	return nil
}

// EventuallyPoolersCondition is like require.Eventually but fetches all pooler statuses
// on each tick and passes them as a batch to condition. Useful when the condition spans
// multiple poolers (e.g. find which one became primary) or needs to return a value.
//
// On each failed tick, all pooler diagnostics are logged to aid flake investigation.
// Returns the typed value produced by condition once it returns met=true.
func EventuallyPoolersCondition[T any](
	t *testing.T,
	poolers []*MultipoolerInstance,
	timeout, tick time.Duration,
	condition func(statuses []PoolerStatusResult) (value T, met bool, reason string),
	msgAndArgs ...any,
) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// EventuallyPoolerCondition is like require.Eventually but automatically fetches status
// for each pooler on every tick and passes it to condition. When condition returns false
// for any pooler, it logs the reason alongside FormatPoolerDiagnostics to aid flake
// investigation. Returns only when all poolers satisfy the condition.
//
// The condition function returns (met bool, reason string). When met=false, reason is
// logged alongside full pooler diagnostics before the next tick.
func EventuallyPoolerCondition(
	t *testing.T,
	poolers []*MultipoolerInstance,
	timeout, tick time.Duration,
	condition func(r PoolerStatusResult) (bool, string),
	msgAndArgs ...any,
) {
	_ = "STUB: not implemented"
	return
}

// RequirePoolerCondition fetches status for each pooler once and immediately fails the
// test if any pooler does not satisfy condition. Diagnostics for all failing poolers are
// included in the failure message. Use this after RequireRecovery or similar operations
// that guarantee the condition is already met — no polling needed.
func RequirePoolerCondition(
	t *testing.T,
	poolers []*MultipoolerInstance,
	condition func(r PoolerStatusResult) (bool, string),
	msgAndArgs ...any,
) {
	_ = "STUB: not implemented"
	return
}

// WaitForNewPrimary polls all multipoolers in setup until one other than oldPrimaryName
// reports IsInitialized + PoolerType_PRIMARY + PostgresReady, then returns its name.
// Fails the test if no new primary is elected within timeout.
func WaitForNewPrimary(t *testing.T, setup *ShardSetup, oldPrimaryName string, timeout time.Duration) string {
	_ = "STUB: not implemented"
	return ""
}

// FormatPoolerDiagnostics returns a compact diagnostic string for a pooler status,
// useful for appending to "not yet ready" log messages to aid flake investigation.
func FormatPoolerDiagnostics(s *multipoolermanagerdatapb.Status, cs *clustermetadatapb.ConsensusStatus) string {
	_ = "STUB: not implemented"
	return ""
}
