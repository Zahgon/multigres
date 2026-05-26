// Copyright 2025 Supabase, Inc.
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

package manager

import (
	"context"
	"sync"

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
)

// ConsensusState manages the in-memory and on-disk consensus state for this node.
// It provides thread-safe access to consensus state and ensures that memory is only
// updated after successful disk writes (pessimistic approach).
type ConsensusState struct {
	poolerDir string
	serviceID *clustermetadatapb.ID

	mu sync.Mutex
	// revocation is persisted to disk; see term_storage.go.
	revocation *clustermetadatapb.TermRevocation
	// replicationPrimary is held in memory only — see HighestKnownRule's proto
	// comment. Populated by RPCs that inform this pooler about the cluster
	// state (SetTermPrimary today; Propose can be wired in via the same RecordTermPrimary
	// entry point). Restarts reset it to nil; coordinators re-inform the
	// pooler.
	replicationPrimary *clustermetadatapb.ReplicationPrimary
}

// NewConsensusState creates a new ConsensusState manager.
// It does not load state from disk - call Load() to initialize.
func NewConsensusState(poolerDir string, serviceID *clustermetadatapb.ID) *ConsensusState {
	_ = "STUB: not implemented"
	return nil
}

// RecordTermPrimary updates the highest-known rule and last-known primary based on
// what this pooler has been told. Either argument may be nil. Bump semantics:
//
//   - rule: updated only when strictly greater than the current value
//     (comparison by RuleNumber: coordinator_term then leader_subterm).
//
//   - primary: updated when the rule advances, OR when the supplied rule
//     equals the current rule but the primary's contact info has changed.
//     The same-rule-different-contact case covers a primary pooler being
//     re-homed (host or port changes) without a new election.
//
// Safe to call on no-op SetTermPrimary paths so the recorded values reflect
// everything the pooler has been told, regardless of whether postgres-side
// changes were applied.
func (cs *ConsensusState) RecordTermPrimary(rule *clustermetadatapb.ShardRule, primary *clustermetadatapb.PoolerAddress) {
	_ = "STUB: not implemented"
	return
}

// Build the next value by starting from the existing fields (via getters,
// which are nil-safe) and overlaying the updates we want to apply.

// Update primary only when one is supplied; an RPC that advances the rule
// without re-stating the primary (or a future WAL-observation path) leaves
// the previously-recorded contact info in place.

// GetReplicationPrimary returns a copy of the primary + rule this pooler has
// been told to use, or nil if it has never been informed. The returned message
// is the same one exposed in ConsensusStatus.
func (cs *ConsensusState) GetReplicationPrimary() *clustermetadatapb.ReplicationPrimary {
	_ = "STUB: not implemented"
	return nil
}

// Load loads consensus state from disk into memory.
// If the file doesn't exist, initializes with default values (term 0, no accepted coordinator).
// This method is idempotent - subsequent calls will reload from disk.
func (cs *ConsensusState) Load() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// GetCurrentTermNumber returns the current term.
// Returns 0 if state has not been loaded.
func (cs *ConsensusState) GetCurrentTermNumber(ctx context.Context) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GetInconsistentCurrentTermNumber returns the current term for monitoring.
// It doesn't require the action lock to be held, so the value returned may
// be outdated by the time it's used. Use GetCurrentTermNumber() as part of
// any action workflow to protect against race conditions.
// Returns 0 if state has not been loaded.
func (cs *ConsensusState) GetInconsistentCurrentTermNumber() (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GetInconsistentRevocation returns a copy of the current term revocation for monitoring.
// It doesn't require the action lock to be held, so the value returned may
// be outdated by the time it's used. Use GetRevocation() as part of any action
// workflow to protect against race conditions.
// Returns nil if state has not been loaded.
func (cs *ConsensusState) GetInconsistentRevocation() (*clustermetadatapb.TermRevocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Return a copy to prevent external modifications

// GetAcceptedLeader returns the coordinator ID this pooler accepted the term from.
// Returns empty string if no coordinator was accepted.
func (cs *ConsensusState) GetAcceptedLeader(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetRevocation returns a copy of the current term revocation.
// Returns nil if state has not been loaded.
func (cs *ConsensusState) GetRevocation(ctx context.Context) (*clustermetadatapb.TermRevocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Return a copy to prevent external modifications

// AcceptCandidateAndSave atomically records acceptance of the term from a coordinator.
// This is called when a node accepts the term during BeginTerm.
// Returns error if already accepted from a different coordinator in this term.
// Idempotent: succeeds if already accepted from the same coordinator.
//
// Deprecated: use AcceptRevocation for new callers. This method remains for BeginTerm()
// compatibility until the legacy coordinator protocol is retired.
func (cs *ConsensusState) AcceptCandidateAndSave(ctx context.Context, candidateID *clustermetadatapb.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// If already accepted from this coordinator, idempotent success

// Check if already accepted from someone else in this term

// Prepare acceptance

// Update acceptance - use proto.Clone to ensure deep copy

// Update last acceptance time

// Save and update under lock

// AcceptRevocation validates and persists a TermRevocation in one atomic step.
// It builds the validation status from the observed position in status combined
// with the current in-memory revocation (read under the mutex), so the check
// reflects the actual locked state rather than a potentially stale snapshot.
// This is the preferred path for Recruit(); the legacy Update* methods below
// remain for BeginTerm() until the old coordinator protocol is retired.
func (cs *ConsensusState) AcceptRevocation(ctx context.Context, status *clustermetadatapb.ConsensusStatus, revocation *clustermetadatapb.TermRevocation) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateTermAndAcceptCandidate atomically updates the term and accepts a candidate in one file write.
// This is used by BeginTerm to avoid two separate file writes.
//
// Deprecated: use AcceptRevocation for new callers. This method remains for BeginTerm()
// compatibility until the legacy coordinator protocol is retired.
// If newTerm > currentTerm, updates term and resets acceptance, then sets the candidate.
// If newTerm == currentTerm, just accepts the candidate (same as AcceptCandidateAndSave).
// Returns error if newTerm < currentTerm.
func (cs *ConsensusState) UpdateTermAndAcceptCandidate(ctx context.Context, newTerm int64, candidateID *clustermetadatapb.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// Higher term: create new revocation with the candidate already set

// Same term: just update acceptance (idempotent check first)

// If already accepted from this coordinator, idempotent success

// Check if already accepted from someone else in this term

// Prepare acceptance

// Single file write

// UpdateTermAndSave atomically updates the term number, resetting accepted coordinator.
// This is called when discovering a newer term from another node.
// Returns error if newTerm < currentTerm.
// Idempotent: succeeds without changes if newTerm == currentTerm.
func (cs *ConsensusState) UpdateTermAndSave(ctx context.Context, newTerm int64) error {
	_ = "STUB: not implemented"
	return nil
}

// If same term, nothing to do (idempotent success)

// Only if newTerm > currentTerm: create new revocation with reset acceptance

// Save and update under lock

// saveAndUpdateLocked saves the revocation to disk and updates memory.
// MUST be called with cs.mu held.
// This is the key method that ensures memory never diverges from disk.
// If the save fails, memory remains unchanged and the error is returned.
func (cs *ConsensusState) saveAndUpdateLocked(newRevocation *clustermetadatapb.TermRevocation) error {
	_ = "STUB: not implemented"
	// Save to disk (lock still held)
	return nil
}

// Save failed - don't update memory, propagate error

// Save succeeded - NOW update memory

// DeleteTermFile removes the consensus term file from disk and resets the
// in-memory state to uninitialized (term 0, no accepted coordinator).
// Called after a pgBackRest restore so the node re-joins consensus from
// scratch; the cluster's current term will be propagated by multiorch on
// first contact via BeginTerm.
// If the file does not exist this is a no-op. Returns an error only if
// the file exists but cannot be removed.
func (cs *ConsensusState) DeleteTermFile() error { _ = "STUB: not implemented"; return nil }

// cloneRevocation creates a deep copy of a TermRevocation.
func cloneRevocation(revocation *clustermetadatapb.TermRevocation) *clustermetadatapb.TermRevocation {
	_ = "STUB: not implemented"
	return nil
}
