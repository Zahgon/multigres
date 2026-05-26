// Copyright 2019 The Vitess Authors.
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
//
// Modifications Copyright 2025 Supabase, Inc.

package topoclient

import (
	"context"
	"sync"
	"time"
)

// This file contains utility methods and definitions to lock resources using topology server.

// Lock describes a long-running lock on a database or a shard.
// It needs to be public as we JSON-serialize it.
type Lock struct {
	// Action and the following fields are set at construction time.
	Action   string
	HostName string
	UserName string
	Time     string
	Options  lockOptions

	// Status is the current status of the Lock.
	Status string
}

// newLock creates a new Lock.
func newLock(action string) *Lock { _ = "STUB: not implemented"; return nil }

// ToJSON returns a JSON representation of the object.
func (l *Lock) ToJSON() (string, error) { _ = "STUB: not implemented"; return "", nil }

// lockInfo is an individual info structure for a lock
type lockInfo struct {
	lockDescriptor LockDescriptor
	actionNode     *Lock
}

// locksInfo is the structure used to remember which locks we took
type locksInfo struct {
	// mu protects the following members of the structure.
	// Safer to be thread safe here, in case multiple go routines
	// lock different things.
	mu sync.Mutex

	// info contains all the locks we took. It is indexed by
	// database (for databases) or database/shard (for shards).
	info map[string]*lockInfo
}

// Context glue
type locksKeyType int

var locksKey locksKeyType

// Support different lock types.
type LockType int

const (
	// Blocking is the default lock type when no other valid type
	// is specified.
	Blocking                LockType = iota
	NonBlocking                      // Uses TryLock
	Named                            // Uses LockName
	NamedNonBlocking                 // Uses LockName semantics with TryLock fail-fast behavior
	NamedNonBlockingWithTTL          // Uses TryLockNameWithTTL for fail-fast with custom TTL
)

func (lt LockType) String() string { _ = "STUB: not implemented"; return "" }

// iTopoLock is the interface for knowing the resource that is being locked.
// It allows for better controlling nuances for different lock types and log messages.
type iTopoLock interface {
	Type() string
	ResourceName() string
	Path() string
}

// perform the topo lock operation
func (l *Lock) lock(ctx context.Context, ts *store, lt iTopoLock, opts ...LockOption) (LockDescriptor, error) {
	_ = "STUB: not implemented"
	return *new(LockDescriptor), nil
}

// Record metrics

// unlock unlocks a previously locked key.
func (l *Lock) unlock(ctx context.Context, ts *store, lt iTopoLock, lockDescriptor LockDescriptor, actionError error) error {
	_ = "STUB: not implemented"
	// Detach from the parent timeout, but preserve the trace span.
	// We need to still release the lock even if the parent context timed out.
	return nil
}

// first update the actionNode

// Record metrics

func (ts *store) internalLock(ctx context.Context, lt iTopoLock, action string, opts ...LockOption) (context.Context, func(*error), error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

// check that we are not already locked

// lock it

// and update our structure

// if we have an error, we log it, but we still want to delete the lock

// both error are set, just log the unlock error

// checkLocked checks that the given resource is locked.
func checkLocked(ctx context.Context, lt iTopoLock) error {
	_ = "STUB: not implemented"
	// extract the locksInfo pointer
	return nil
}

// find the individual entry

// Check the lock server implementation still holds the lock.

// lockOptions configure a Lock call. lockOptions are set by the LockOption
// values passed to the lock functions.
type lockOptions struct {
	lockType LockType
	ttl      time.Duration
}

// LockOption configures how we perform the locking operation.
type LockOption interface {
	apply(*lockOptions)
}

// funcLockOption wraps a function that modifies lockOptions into an
// implementation of the LockOption interface.
type funcLockOption struct {
	f func(*lockOptions)
}

func (flo *funcLockOption) apply(lo *lockOptions) { _ = "STUB: not implemented"; return }

func newFuncLockOption(f func(*lockOptions)) *funcLockOption { _ = "STUB: not implemented"; return nil }

// WithType determines the type of lock we take. The options are defined
// by the LockType type.
func WithType(lt LockType) LockOption { _ = "STUB: not implemented"; return *new(LockOption) }

// WithTTL sets a custom TTL for the lock lease.
// For Named locks, this overrides the default NamedLockTTL (24h).
func WithTTL(ttl time.Duration) LockOption { _ = "STUB: not implemented"; return *new(LockOption) }
