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
	"sync/atomic"
	"time"

	"golang.org/x/sync/semaphore"

	multipoolermanagerdatapb "github.com/multigres/multigres/go/pb/multipoolermanagerdata"
)

// actionLockKey is the context key for storing action lock information
type actionLockKey struct{}

// actionLockValue contains information about a lock acquisition
type actionLockValue struct {
	lockID    uint64
	operation string
	released  *atomic.Bool
}

// ActionLock is a wrapper around a semaphore that tracks ownership via Context.
// It ensures that only one goroutine can hold the lock at a time and provides
// mechanisms to assert that a given context holds the lock.
type ActionLock struct {
	sema      *semaphore.Weighted
	mu        sync.Mutex
	currentID uint64 // ID of current lock holder (0 if unlocked)
	nextID    uint64 // Counter for generating unique IDs

	// activeAction and activeActionStartedAt track the current postgres action
	// in progress. Protected by mu. Cleared automatically on Release.
	activeAction          multipoolermanagerdatapb.PostgresAction
	activeActionStartedAt time.Time
}

// NewActionLock creates a new ActionLock.
func NewActionLock() *ActionLock { _ = "STUB: not implemented"; return nil }

// Start at 1 so 0 can represent "unlocked"

// Acquire acquires the action lock and returns a new context that proves ownership.
// The operation string is used for debugging/tracking purposes.
// Returns an error if the lock cannot be acquired (e.g., context cancelled) or
// if the provided context already holds the lock.
func (al *ActionLock) Acquire(ctx context.Context, operation string) (context.Context, error) {
	_ = "STUB: not implemented"
	// Check if this context already holds the lock
	return *new(context.Context), nil
}

// Try to acquire the semaphore

// Generate a unique ID for this acquisition

// Create the lock value with a released flag

// Return a new context with the lock info

// Release releases the action lock. It validates that the provided context
// holds the lock and panics if it doesn't (which indicates a programming error).
// After releasing, the context is marked as invalid so future assertions will fail.
func (al *ActionLock) Release(ctx context.Context) { _ = "STUB: not implemented"; return }

// Check if already released

// Verify this context holds the current lock

// Mark as released BEFORE actually releasing the semaphore
// This ensures assertions fail immediately

// SetAction records the postgres action currently being performed.
// Must be called while holding the lock (enforced via AssertActionLockHeld).
func (al *ActionLock) SetAction(ctx context.Context, action multipoolermanagerdatapb.PostgresAction) error {
	_ = "STUB: not implemented"
	return nil
}

// ActiveAction returns the current postgres action and how long it has been running.
// Returns UNSPECIFIED and zero duration when no action is in progress.
// Safe to call without holding the action lock.
func (al *ActionLock) ActiveAction() (multipoolermanagerdatapb.PostgresAction, time.Duration) {
	_ = "STUB: not implemented"
	return *new(multipoolermanagerdatapb.PostgresAction), *new(time.Duration)
}

// AssertActionLockHeld returns an error if the provided context does not hold
// an action lock or if the lock has been released. This is a global function
// that doesn't require a reference to the ActionLock.
func AssertActionLockHeld(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
