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

package test

import (
	"context"
	"testing"
	"time"

	"github.com/multigres/multigres/go/common/topoclient"
)

// timeUntilLockIsTaken is the time to wait until a lock is taken.
// We haven't found a better simpler way to guarantee a routine is stuck
// waiting for a topo lock than sleeping that amount.
var timeUntilLockIsTaken = 10 * time.Millisecond

// checkLock checks we can lock / unlock as expected. It's using a database
// as the lock target.
func checkLock(t *testing.T, ctx context.Context, ts topoclient.Store) {
	_ = "STUB: not implemented"
	return
}

func checkLockTimeout(ctx context.Context, t *testing.T, conn topoclient.Conn) {
	_ = "STUB: not implemented"
	return
}

// We have the lock, list the database directory.
// It should not contain anything, except Ephemeral files.
/*full*/

// Non-ephemeral entries better have only ephemeral children.

/*full*/

// test we can't take the lock again

// test we can interrupt taking the lock

// test we can't unlock again

// checkLockUnblocks makes sure that a routine waiting on a lock
// is unblocked when another routine frees the lock
func checkLockUnblocks(ctx context.Context, t *testing.T, conn topoclient.Conn) {
	_ = "STUB: not implemented"
	return
}

// As soon as we're unblocked, we try to lock the database.

// Lock the database.

// unblock the go routine so it starts waiting

// sleep for a while so we're sure the go routine is blocking

// checkLockName checks if we can lock / unlock using LockName as expected.
// LockName doesn't require the path to exist and has a static 24-hour TTL.
func checkLockName(t *testing.T, ctx context.Context, ts topoclient.Store) {
	_ = "STUB: not implemented"
	return
}

// Use a non-existent path since LockName doesn't require it to exist

// We should not be able to take the same named lock again

// test we can interrupt taking the lock

// test we can't unlock again

// checkLockNameWithTTL checks if we can lock / unlock using LockNameWithTTL as expected.
// LockNameWithTTL doesn't require the path to exist and allows specifying a custom TTL.
func checkLockNameWithTTL(t *testing.T, ctx context.Context, ts topoclient.Store) {
	_ = "STUB: not implemented"
	return
}

// Use a non-existent path since LockNameWithTTL doesn't require it to exist

// Test with custom TTL (1 hour)

// We should not be able to take the same named lock again

// test we can interrupt taking the lock

// test we can't unlock again

// Test with zero TTL (should default to NamedLockTTL)

// checkTryLockName checks if we can lock / unlock using TryLockName as expected.
// TryLockName doesn't require the path to exist and fails fast if the lock is already held.
func checkTryLockName(t *testing.T, ctx context.Context, ts topoclient.Store) {
	_ = "STUB: not implemented"
	return
}

// Use a non-existent path since TryLockName doesn't require it to exist

// We should not be able to take the same lock again - it should fail fast with NodeExists

// After unlocking, we should be able to acquire the lock again

// checkTryLock checks if we can lock / unlock as expected. It's using a database
// as the lock target.
func checkTryLock(t *testing.T, ctx context.Context, ts topoclient.Store) {
	_ = "STUB: not implemented"
	return
}

// checkTryLockTimeout test the fail-fast nature of TryLock
func checkTryLockTimeout(ctx context.Context, t *testing.T, conn topoclient.Conn) {
	_ = "STUB: not implemented"
	return
}

// We have the lock, list the cell location directory.
// It should not contain anything, except Ephemeral files.
/*full*/

// Non-ephemeral entries better have only ephemeral children.

/*full*/

// We should not be able to take the lock again. It should throw `NodeExists` error.

// test we can interrupt taking the lock

// go routine to cancel the context.

// after attempting the `TryLock` and getting an error `NodeExists`, we will cancel the context deliberately
// and expect `context canceled` error in next iteration of `for` loop.

// we expect context to fail with `context canceled` error

// test we can't unlock again

// unlike 'checkLockUnblocks', checkTryLockUnblocks will not block on other client but instead
// keep retrying until it gets the lock.
func checkTryLockUnblocks(ctx context.Context, t *testing.T, conn topoclient.Conn) {
	_ = "STUB: not implemented"
	return
}

// TryLock will keep getting NodeExists until lockDescriptor2 unlock itself.
// It will not wait but immediately return with NodeExists error.

// Lock the database.

// unblock the go routine so it starts waiting
