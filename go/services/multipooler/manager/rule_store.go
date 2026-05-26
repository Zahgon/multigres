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
	"errors"
	"log/slog"
	"sync"
	"time"

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	"github.com/multigres/multigres/go/services/multipooler/executor"
)

// ruleStorer is the interface for reading and writing the current shard rule.
// *ruleStore implements this; tests use fakeRuleStore.
type ruleStorer interface {
	// observePosition reads the current rule and WAL LSN from postgres.
	// Always returns a non-nil position when err is nil (the initial row guarantees a row exists).
	observePosition(ctx context.Context) (*clustermetadatapb.PoolerPosition, error)
	updateRule(ctx context.Context, update *ruleUpdateBuilder) (*clustermetadatapb.PoolerPosition, error)
	// createRuleTables creates multigres.current_rule and multigres.rule_history
	// if they do not already exist, and inserts the initial row for the default
	// shard, populated with the given durability policy. bootstrapID is recorded
	// as the initial row's coordinator_id, analogous to how a pooler is the
	// coordinator for leader-led rule changes. It is idempotent and safe to
	// call multiple times.
	createRuleTables(ctx context.Context, policy *clustermetadatapb.DurabilityPolicy, bootstrapID *clustermetadatapb.ID) error
	// cachedPosition returns the most recently observed or written PoolerPosition
	// from memory, without querying postgres. Returns nil if no position has been
	// cached yet (e.g. before the first observePosition or updateRule call).
	cachedPosition() *clustermetadatapb.PoolerPosition

	// hasInconsistentGUC returns true if the cached rule's policy would produce
	// different GUC strings than what postgres currently has. Safe to call
	// without the action lock.
	hasInconsistentGUC(ctx context.Context) bool

	// reconcileGUC re-reads the current rule (under SELECT FOR UPDATE when
	// inRecovery is false) and re-applies the GUC if needed. Requires the
	// action lock.
	reconcileGUC(ctx context.Context, inRecovery bool) error
}

// ruleStore manages the current shard rule in postgres.
//
// All DB operations that write or read the current rule go through ruleStore,
// ensuring consistent access to rule state.
type ruleStore struct {
	logger       *slog.Logger
	queryService executor.InternalQueryService
	syncStandby  SyncStandbyManager

	mu      sync.Mutex
	lastPos *clustermetadatapb.PoolerPosition // updated on every observePosition / updateRule
}

// newRuleStore creates a ruleStore. ssm must not be nil; tests that do not
// need GUC verification should pass noopSyncStandbyManager{}.
func newRuleStore(
	logger *slog.Logger,
	qs executor.InternalQueryService,
	ssm SyncStandbyManager,
) *ruleStore {
	_ = "STUB: not implemented"
	return nil
}

// cacheRuleObservation updates the in-memory position cache.
func (rs *ruleStore) cacheRuleObservation(pos *clustermetadatapb.PoolerPosition) {
	_ = "STUB: not implemented"
	return
}

// This position observation is stale. Ignore it.

// cachedPosition returns the most recently observed or written PoolerPosition
// from memory. Returns nil if no position has been cached yet.
func (rs *ruleStore) cachedPosition() *clustermetadatapb.PoolerPosition {
	_ = "STUB: not implemented"
	return nil
}

// hasInconsistentGUC returns true if the cached rule's policy would produce
// different GUC strings than what postgres currently has. Safe to call
// without the action lock.
func (rs *ruleStore) hasInconsistentGUC(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

// reconcileGUC re-reads the current rule under SELECT FOR UPDATE to drain prior
// writers, then re-applies the GUC if the cached values are stale. Requires the
// action lock.
func (rs *ruleStore) reconcileGUC(ctx context.Context, inRecovery bool) error {
	_ = "STUB: not implemented"
	return nil
}

// ----------------------------------------------------------------------------
// Rule Update Builder
// ----------------------------------------------------------------------------

// ruleNumber identifies a specific rule version by coordinator term and subterm.
type ruleNumber struct {
	coordinatorTerm int64
	leaderSubterm   int64
}

// ruleUpdateBuilder constructs the parameters for updateRule.
// coordinatorID, eventType, reason, and createdAt are always required.
// Fields not set via builder methods retain their current value in current_rule.
type ruleUpdateBuilder struct {
	// required
	termNumber    int64
	coordinatorID *clustermetadatapb.ID
	eventType     string
	reason        string
	createdAt     time.Time

	// optional; nil means keep the existing value in current_rule
	leaderID         *clustermetadatapb.ID
	cohortMembers    []*clustermetadatapb.ID
	durabilityPolicy *clustermetadatapb.DurabilityPolicy

	// history-only optional fields
	walPosition     string
	operation       string
	acceptedMembers []*clustermetadatapb.ID

	force              bool
	skipOutgoingQuorum bool        // skip BuildPolicyTransition; apply incoming GUC directly
	previousRule       *ruleNumber // for compare-and-swap; nil means no check
	promotionHook      promotionFn // non-nil iff postgres is known to be in recovery
}

// promotionFn is called by updateRule after the pre-promote GUC is applied and
// before the rule history write. It must call pg_promote() and wait for promotion
// to complete. It is provided iff the caller has already verified that postgres
// is in recovery.
type promotionFn func(ctx context.Context) error

func newRuleUpdate(termNumber int64, coordinatorID *clustermetadatapb.ID, eventType, reason string, createdAt time.Time) *ruleUpdateBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *ruleUpdateBuilder) withLeader(id *clustermetadatapb.ID) *ruleUpdateBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *ruleUpdateBuilder) withCohort(members []*clustermetadatapb.ID) *ruleUpdateBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *ruleUpdateBuilder) withWALPosition(pos string) *ruleUpdateBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *ruleUpdateBuilder) withPromotionHook(fn promotionFn) *ruleUpdateBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *ruleUpdateBuilder) withOperation(op string) *ruleUpdateBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *ruleUpdateBuilder) withAcceptedMembers(members []*clustermetadatapb.ID) *ruleUpdateBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *ruleUpdateBuilder) withDurabilityPolicy(policy *clustermetadatapb.DurabilityPolicy) *ruleUpdateBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *ruleUpdateBuilder) withForce() *ruleUpdateBuilder { _ = "STUB: not implemented"; return nil }

// withSkipOutgoingQuorum instructs updateRule to skip BuildPolicyTransition and apply
// the incoming cohort GUC directly (Both = Incoming). Used for coordinator-directed
// changes where the outgoing cohort is empty (bootstrap) or the coordinator has already
// verified the transition is safe, so no dual-ack window is needed.
func (b *ruleUpdateBuilder) withSkipOutgoingQuorum() *ruleUpdateBuilder {
	_ = "STUB: not implemented"
	return nil
}

// withPreviousRule adds a compare-and-swap check: the update only proceeds if the
// current rule matches the given coordinator term and subterm.
func (b *ruleUpdateBuilder) withPreviousRule(coordinatorTerm, leaderSubterm int64) *ruleUpdateBuilder {
	_ = "STUB: not implemented"
	return nil
}

// ----------------------------------------------------------------------------
// Schema Operations
// ----------------------------------------------------------------------------

// createRuleTables creates multigres.current_rule and multigres.rule_history if
// they do not already exist, then inserts the initial row for the default
// shard. It is idempotent and safe to call multiple times.
//
// current_rule holds a single row per shard representing the current cluster rule.
// It is used as a locking target (SELECT FOR UPDATE) to serialise concurrent
// writes; rule_history provides the append-only audit log.
//
// coordinator_term=0 in the initial row means no rule has been applied yet.
// policy is written into the initial row so all subsequent rule reads have a
// non-nil DurabilityPolicy; operations that do not change the policy (e.g.
// Promote) carry it forward via COALESCE in updateRule.
//
// bootstrapID becomes the initial row's coordinator_id. The pooler that
// initializes the schema acts as the coordinator for the initial row —
// analogous to how a pooler is the coordinator for leader-led rule changes.
func (rs *ruleStore) createRuleTables(ctx context.Context, policy *clustermetadatapb.DurabilityPolicy, bootstrapID *clustermetadatapb.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// Each row records a cluster state change (promotion, cohort membership, durability policy).
// The composite primary key (coordinator_term, leader_subterm) uniquely identifies each rule;
// leader_subterm is assigned by the application as MAX(leader_subterm)+1 within a coordinator_term.

// ----------------------------------------------------------------------------
// Read/Write Operations
// ----------------------------------------------------------------------------

// errRuleConflict is returned by updateRule when a compare-and-swap check fails:
// either withPreviousRule's explicit version check did not match, or a concurrent
// write changed the rule between our read and our write.
var errRuleConflict = errors.New("rule conflict: current rule version changed since last read")

// ----------------------------------------------------------------------------
// Shared row reader
// ----------------------------------------------------------------------------

// readCurrentRule reads the current_rule row for the default shard. If forUpdate
// is true, appends FOR UPDATE NOWAIT to acquire a row-level lock; the NOWAIT
// clause causes an immediate error if the row is already locked rather than
// blocking, so callers never wait indefinitely. On a standby this must be false
// since the node is read-only. Returns an error when the sentinel row is missing
// (tables not initialized) or when postgres is unreachable.
//
// The caller is responsible for adding an appropriate context timeout.
func (rs *ruleStore) readCurrentRule(ctx context.Context, forUpdate bool) (*clustermetadatapb.PoolerPosition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// observePosition reads the current rule and WAL LSN from postgres and returns
// the observed position. Always returns a non-nil position when err is nil.
//
// Returns an error if postgres is unreachable or if the current_rule sentinel
// row is missing (which indicates the tables are not initialized).
func (rs *ruleStore) observePosition(ctx context.Context) (*clustermetadatapb.PoolerPosition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// readCurrentRuleLocked reads the current_rule row and returns a lockedCtx that
// carries proof that prior rule writes from any previous action lock holder have
// been drained (withPriorRuleWritesDrained). The timeout is managed internally;
// lockedCtx is derived from ctx (not the internal timeout context) and remains
// valid for subsequent operations after the read completes.
//
// When inRecovery is false (primary path): uses FOR UPDATE NOWAIT, which
// succeeds immediately if no other transaction holds the row lock, or fails
// fast if the row is locked. Callers that receive an error should retry.
// When inRecovery is true (standby/promotion path): omits FOR UPDATE since the
// node is read-only and no concurrent writes to current_rule are possible.
func (rs *ruleStore) readCurrentRuleLocked(ctx context.Context, inRecovery bool) (*clustermetadatapb.PoolerPosition, context.Context, error) {
	_ = "STUB: not implemented"
	return nil, *new(context.Context), nil
}

// updateRule writes a new rule to current_rule and rule_history.
//
// The leader_subterm is assigned as:
//   - 0 if termNumber is greater than the current coordinator_term (new term)
//   - current leader_subterm + 1 if termNumber equals the current coordinator_term
//
// Fields not set via the builder (leaderID, cohortMembers, durabilityPolicy) retain
// their current values from current_rule.
//
// GUC transition: the outgoing ("both") policy is applied before the WAL write so that
// writes issued during the transition satisfy both the old and new replication requirements.
// The incoming (new) policy is applied after the write commits. On a promotion the outgoing
// GUC is applied while still a standby (before pg_promote); on a primary-side rule change
// it is applied immediately before the write CTE.
//
// Returns the node's position (rule + WAL LSN) at the time of the write,
// or nil if force mode skipped the write.
//
// This operation uses the remote-operation-timeout and will fail if it cannot
// complete within that time. A timeout typically indicates that synchronous
// replication is not functioning.
func (rs *ruleStore) updateRule(ctx context.Context, update *ruleUpdateBuilder) (*clustermetadatapb.PoolerPosition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Force mode skips history recording entirely. Force operations are emergency
// operations that must configure replication GUCs regardless. The write would
// block on sync replication with unreachable standbys, consuming the parent
// context's deadline and causing subsequent GUC changes to fail.

// Identity and timing must be supplied by the caller. ClusterIDString(nil)
// silently returns "" and the coordinator_id column is TEXT NOT NULL (not
// rejected by postgres because "" != NULL), so without these checks a nil
// coordinatorID would write a corrupt row instead of failing. createdAt
// has the same property: a zero time.Time inserts as a zero timestamp.
// Failing fast here also avoids leaving partial work in the caller, which
// often touches postgres GUCs around this write.

// Read the current rule to establish the CAS baseline and drain any in-flight
// rule writes from a previous action lock holder.

// Optional explicit CAS: verify the caller's expected version matches what we read.

// Compute the next leader_subterm.

// Resolve values to write: caller-supplied values take priority; nil retains existing.

// Validate that the new cohort can satisfy the new durability policy.

// Compute the GUC transition. The Both policy satisfies the old and new durability
// requirements simultaneously and is applied before the WAL write. The Incoming
// (new) policy is applied after the write commits.

// Skip BuildPolicyTransition and apply the incoming cohort directly.
// Used when the outgoing cohort is empty (bootstrap) or the coordinator
// has already verified the transition is safe.

// Convert values to SQL parameters.

// newDP is always non-nil: updateRule falls back to the current rule's policy when
// the caller omits withDurabilityPolicy(), so these values are always present.

// Apply the transition GUC before writing the rule. The transition (Both) policy
// satisfies both old and new durability requirements simultaneously.
// Promotion path: set GUC while still a standby, then call pg_promote().
// Primary path: set GUC immediately before the write CTE.

// Write the rule. The remote-operation timeout applies because this write must be
// acknowledged by synchronous standbys; a timeout indicates replication is not functioning.

// shard_id
// cas_term
// cas_subterm
// new_term
// new_subterm
// new_leader_id (NULLIF: leader absent on sentinel row)
// new_coordinator_id
// new_cohort
// dp_name
// dp_quorum_type
// dp_required_count
// created_at
// event_type
// wal_position (NULLIF: optional)
// operation    (NULLIF: optional)
// reason
// accepted_members

// Zero rows means either the CAS check failed (concurrent write between our read
// and write) or the shard row is missing (should never happen after initialisation).

// Apply the incoming (new) GUC after the write commits.

// queryRuleHistory returns the most recent rule history records in descending
// order by (coordinator_term, leader_subterm). Returns at most limit records.
func (rs *ruleStore) queryRuleHistory(ctx context.Context, limit int) ([]ruleHistoryRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ----------------------------------------------------------------------------
// Helpers
// ----------------------------------------------------------------------------

// buildPoolerPosition constructs a *clustermetadatapb.PoolerPosition from raw DB column values.
// leaderIDStr and coordinatorIDStr are app-name formatted strings (e.g. "zone1_pooler-name").
// Durability fields are NOT NULL in the DB and are always populated in the returned position.
// createdAt is the coordinator-supplied CreationTime persisted with the rule.
func buildPoolerPosition(
	coordinatorTerm, leaderSubterm int64,
	leaderIDStr *string,
	coordinatorIDStr string,
	cohortNames []string,
	durabilityPolicyName, durabilityQuorumType string,
	durabilityRequiredCount int64,
	createdAt time.Time,
	lsn string,
) (*clustermetadatapb.PoolerPosition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Coordinator IDs are multiorch, not multipooler — parseApplicationName
// is pooler-specific, so decode the cell_name encoding directly.

// appNamesToIDs converts a slice of app-name formatted strings to proto IDs.
func appNamesToIDs(names []string) ([]*clustermetadatapb.ID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ruleHistoryRecord represents a row from multigres.rule_history or multigres.current_rule.
type ruleHistoryRecord struct {
	CoordinatorTerm         int64
	LeaderSubterm           int64
	EventType               string
	LeaderID                *poolerID // nil if not set
	CoordinatorID           *string   // informational only; component type is not stored
	WALPosition             *string
	Operation               *string
	Reason                  string
	CohortMembers           []poolerID
	AcceptedMembers         []poolerID
	DurabilityPolicyName    string
	DurabilityQuorumType    string
	DurabilityRequiredCount int32
	CreatedAt               time.Time
}

// parsePoolerIDStrings converts a slice of "cell_name" app name strings into poolerIDs.
// Returns nil for nil input, preserving the distinction between "not set" and "empty".
func parsePoolerIDStrings(names []string) ([]poolerID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// scanRuleHistoryRow scans string-typed DB columns into a ruleHistoryRecord,
// parsing leader_id, cohort_members, and accepted_members into poolerIDs.
// leaderIDStr, cohortNames, and acceptedNames are intermediary scan targets.
func scanRuleHistoryRow(rec *ruleHistoryRecord, leaderIDStr *string, cohortNames, acceptedNames []string) error {
	_ = "STUB: not implemented"
	return nil
}

// priorRuleWritesDrainedKey is a context key proving that any in-flight rule
// writes from a previous action lock holder have been resolved before
// SyncStandbyManager.SetPolicy is called. This is established in one of two ways:
//
//   - Primary path: a SELECT FOR UPDATE on current_rule blocks until any
//     in-progress transaction from the prior holder commits or rolls back, after
//     which our row lock prevents new writers from interposing.
//   - Recovery path (standby before pg_promote): the node is read-only, so no
//     concurrent writes to current_rule are possible.
//
// The action lock (checked separately via AssertActionLockHeld) ensures no
// concurrent goroutine in this process can also hold this proof.
type priorRuleWritesDrainedKey struct{}

// withPriorRuleWritesDrained returns a derived context carrying proof that any
// in-flight rule writes from the previous action lock holder have been resolved.
// Called by readCurrentRuleLocked; callers must not stamp the context themselves.
func withPriorRuleWritesDrained(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// assertPriorRuleWritesDrained returns an error if the context does not carry
// proof that prior rule writes have been drained. Called automatically via
// readCurrentRuleLocked; callers must not stamp the context themselves.
func assertPriorRuleWritesDrained(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
