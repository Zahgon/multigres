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

package memorytopo

import (
	"context"
	"time"

	"github.com/multigres/multigres/go/common/topoclient"
)

// convertError converts a context error into a topo error.
func convertError(err error, nodePath string) error { _ = "STUB: not implemented"; return nil }

// memoryTopoLockDescriptor implements topoclient.LockDescriptor.
type memoryTopoLockDescriptor struct {
	c       *conn
	dirPath string
	lockCh  chan struct{} // the lock channel at the time of acquisition
}

// TryLock is part of the topoclient.Conn interface.
func (c *conn) TryLock(ctx context.Context, dirPath, contents string) (topoclient.LockDescriptor, error) {
	_ = "STUB: not implemented"
	return *new(topoclient.LockDescriptor), nil
}

// checkLockExistence is a private helper method that checks if a lock already exists for the given path.
// It returns nil if no lock exists, or &topoclient.TopoError{Code: topoclient.NodeExists} if a lock already exists.
func (c *conn) checkLockExistence(ctx context.Context, dirPath string, named bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if a lock exists

// No lock exists

// Lock is part of the topoclient.Conn interface.
func (c *conn) Lock(ctx context.Context, dirPath, contents string) (topoclient.LockDescriptor, error) {
	_ = "STUB: not implemented"
	return *new(topoclient.LockDescriptor), nil
}

// LockWithTTL is part of the topoclient.Conn interface.
func (c *conn) LockWithTTL(ctx context.Context, dirPath, contents string, ttl time.Duration) (topoclient.LockDescriptor, error) {
	_ = "STUB: not implemented"
	return *new(topoclient.LockDescriptor), nil
}

// LockName is part of the topoclient.Conn interface.
func (c *conn) LockName(ctx context.Context, dirPath, contents string) (topoclient.LockDescriptor, error) {
	_ = "STUB: not implemented"
	return *new(topoclient.LockDescriptor), nil
}

// LockNameWithTTL is part of the topoclient.Conn interface.
func (c *conn) LockNameWithTTL(ctx context.Context, dirPath, contents string, ttl time.Duration) (topoclient.LockDescriptor, error) {
	_ = "STUB: not implemented"
	return *new(topoclient.LockDescriptor), nil
}

// TryLockWithLease is part of the topoclient.Conn interface.
func (c *conn) TryLockWithLease(ctx context.Context, key, contents string, ttl time.Duration) (topoclient.LockDescriptor, error) {
	_ = "STUB: not implemented"
	return *new(topoclient.LockDescriptor), nil
}

// TryLockNameWithTTL is part of the topoclient.Conn interface.
func (c *conn) TryLockNameWithTTL(ctx context.Context, dirPath, contents string, ttl time.Duration) (topoclient.LockDescriptor, error) {
	_ = "STUB: not implemented"
	return *new(topoclient.LockDescriptor), nil
}

// Check if lock exists, using named=true so the path is created if needed

// TryLockName is part of the topoclient.Conn interface.
func (c *conn) TryLockName(ctx context.Context, dirPath, contents string) (topoclient.LockDescriptor, error) {
	_ = "STUB: not implemented"
	return *new(topoclient.LockDescriptor), nil
}

// Check if lock exists, using named=true so the path is created if needed

// lock acquires a lock without TTL.
func (c *conn) lock(ctx context.Context, dirPath, contents string, named bool) (topoclient.LockDescriptor, error) {
	_ = "STUB: not implemented"
	return *new(topoclient.LockDescriptor), nil
}

// lockWithTTL acquires a lock with an optional TTL. If ttl is 0, the lock does not expire.
func (c *conn) lockWithTTL(ctx context.Context, dirPath, contents string, named bool, ttl time.Duration) (topoclient.LockDescriptor, error) {
	_ = "STUB: not implemented"
	return *new(topoclient.LockDescriptor), nil
}

// Someone else has the lock. Just wait for it.

// Node was unlocked, try again to grab it.

// Done waiting

// No one has the lock, grab it.

// Set up TTL expiration if specified

// Only expire if the lock is still held (not already unlocked)

// Check is part of the topoclient.LockDescriptor interface.
// Returns an error if the lock has been force-unlocked or stolen by another holder.
func (ld *memoryTopoLockDescriptor) Check(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Unlock is part of the topoclient.LockDescriptor interface.
func (ld *memoryTopoLockDescriptor) Unlock(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// RevokeLockWithLease is part of the topoclient.Conn interface.
func (c *conn) RevokeLockWithLease(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

// forceUnlock forcefully removes the lock at the given path regardless of who holds it.
func (c *conn) forceUnlock(ctx context.Context, dirPath string) error {
	_ = "STUB: not implemented"
	return nil
}

// No node, no lock to remove

// No lock held

func (c *conn) unlock(ctx context.Context, dirPath string) error {
	_ = "STUB: not implemented"
	return nil
}

// clearLock stops any TTL timer, closes the lock channel, and clears
// lock-related fields. Must be called with factory.mu held.
func clearLock(n *node) { _ = "STUB: not implemented"; return }
