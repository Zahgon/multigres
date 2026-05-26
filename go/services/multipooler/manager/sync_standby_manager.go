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

package manager

import (
	"context"
	"log/slog"
	"sync"

	commonconsensus "github.com/multigres/multigres/go/common/consensus"
	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	multipoolermanagerdatapb "github.com/multigres/multigres/go/pb/multipoolermanagerdata"
	"github.com/multigres/multigres/go/services/multipooler/executor"
)

// SyncStandbyManager owns all writes to the synchronous_standby_names GUC.
// Nobody sets the GUC directly; they ask this service to do it. This
// centralisation ensures the ordering of GUC changes relative to WAL writes is
// always correct for the transition type (add vs remove vs promotion).
type SyncStandbyManager interface {
	// SetPolicy computes and applies the Postgres GUC for the given policy.
	// pc.Cohort is the full participant set (leader included). Returns an error if
	// the policy produces no eligible standbys; use Clear instead in that case.
	//
	// ctx must carry either a ruleRowLock token (from lockCurrentRule on a
	// primary) or a prePromote token (from lockCurrentRule on a standby), both
	// of which are set automatically by ruleStore.updateRule.
	SetPolicy(ctx context.Context, pc commonconsensus.PolicyWithCohort) error

	// Clear resets synchronous_standby_names to its default (empty) value and
	// invalidates the in-memory cache. It must only be called after postgres has
	// entered recovery mode (pg_is_in_recovery() = true); calling it on a primary
	// would allow commits to proceed without standby acknowledgment.
	Clear(ctx context.Context) error

	// NeedsApply returns true if the given policy would produce GUC strings that
	// differ from what postgres currently has. It first checks the in-memory cache;
	// when the cache matches the desired state it validates against the live
	// postgres value to catch external GUC changes (e.g. manual ALTER SYSTEM).
	// Safe to call without holding the action lock.
	NeedsApply(ctx context.Context, pc commonconsensus.PolicyWithCohort) (bool, error)
}

// postgresqlSyncStandbyManager implements SyncStandbyManager against a live
// PostgreSQL instance. It is the sole writer of synchronous_commit and
// synchronous_standby_names; the in-memory cache is therefore always consistent
// with what was last applied.
type postgresqlSyncStandbyManager struct {
	logger  *slog.Logger
	qs      executor.InternalQueryService
	localID *clustermetadatapb.ID // identity of the local pooler (always the primary when SetPolicy is called)

	mu               sync.Mutex
	lastSyncCommit   string // serialised GUC string ("on", "remote_apply", …); empty = unknown
	lastStandbyNames string // serialised GUC string ("FIRST 1 (…)"); empty = unknown
}

func newSyncStandbyManager(logger *slog.Logger, qs executor.InternalQueryService, localID *clustermetadatapb.ID) *postgresqlSyncStandbyManager {
	_ = "STUB: not implemented"
	return nil
}

func (s *postgresqlSyncStandbyManager) exec(ctx context.Context, sql string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *postgresqlSyncStandbyManager) setSynchronousCommit(ctx context.Context, level multipoolermanagerdatapb.SynchronousCommitLevel) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *postgresqlSyncStandbyManager) setStandbyNames(ctx context.Context, method multipoolermanagerdatapb.SynchronousMethod, numSync int32, names []poolerID) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *postgresqlSyncStandbyManager) reloadConfig(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// reloadPostgresConfig confirms postmaster has re-read the config, but pooled
// backends pick up SIGHUP asynchronously. Sleep briefly so the next WAL write
// is unlikely to land on a backend still using the old GUC values.

// computedGUC holds the GUC strings and the config needed to apply them.
// A zero wantCommit means the policy produced no eligible standbys.
type computedGUC struct {
	cfg          *commonconsensus.SyncReplicationConfig
	standbyNames []poolerID
	wantCommit   string
	wantStandby  string
}

// computeGUC derives the expected GUC state for the current policy.
// Returns a zero computedGUC (wantCommit == "") when the policy produces no
// eligible standbys; the caller should use Clear in that case.
func (s *postgresqlSyncStandbyManager) computeGUC(pc commonconsensus.PolicyWithCohort) (computedGUC, error) {
	_ = "STUB: not implemented"
	return *new(computedGUC), nil
}

// NeedsApply returns true if the given policy would produce GUC strings that
// differ from what postgres currently has. Safe to call without the action lock.
//
// It queries postgres directly so it can detect GUC changes made outside this
// manager (e.g. manual ALTER SYSTEM). On any query failure it falls back to the
// in-memory cache, which is reliable because this manager is the sole writer of
// these GUCs.
func (s *postgresqlSyncStandbyManager) NeedsApply(ctx context.Context, pc commonconsensus.PolicyWithCohort) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Query postgres for the live GUC values.

// Fall back to cache: unreachable postgres means we can't detect drift,
// but the cache is trustworthy since we own all writes to these GUCs.

// Update the cache to reflect confirmed postgres state so SetPolicy
// skips redundant ALTER SYSTEM calls (e.g. after a process restart).

// SetPolicy computes the Postgres GUC configuration for the given durability policy and applies it.
// Uses an in-memory cache of the last-written GUC strings to skip ALTER SYSTEM calls when the
// desired values haven't changed. This is safe because postgresqlSyncStandbyManager is the sole
// writer of synchronous_commit and synchronous_standby_names.
func (s *postgresqlSyncStandbyManager) SetPolicy(ctx context.Context, pc commonconsensus.PolicyWithCohort) error {
	_ = "STUB: not implemented"
	return nil
}

// Clear resets synchronous_standby_names to its default value and invalidates the
// in-memory cache. Called during demotion so that commits do not block on standbys
// that are no longer connected to this node.
func (s *postgresqlSyncStandbyManager) Clear(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Safety: clearing synchronous_standby_names on a primary would allow commits
// to proceed without standby acknowledgment, violating durability guarantees.

// syncCommitString converts a SynchronousCommitLevel enum to the PostgreSQL GUC string.
func syncCommitString(level multipoolermanagerdatapb.SynchronousCommitLevel) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
