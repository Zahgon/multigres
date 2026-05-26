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

package etcdtopo

import (
	"context"
	"time"

	"github.com/spf13/pflag"

	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/multigres/multigres/go/common/servenv"
	"github.com/multigres/multigres/go/common/topoclient"
)

var leaseTTL = 30 // This is the default used for all non-named locks

func init() {
	servenv.OnParse(registerEtcd2TopoLockFlags)
}

func registerEtcd2TopoLockFlags(fs *pflag.FlagSet) { _ = "STUB: not implemented"; return }

// newUniqueEphemeralKV creates a new file in the provided directory.
// It is linked to the Lease.
// Errors returned are converted to topo errors.
func (s *etcdtopo) newUniqueEphemeralKV(ctx context.Context, cli *clientv3.Client, leaseID clientv3.LeaseID, nodePath string, contents string) (string, int64, error) {
	_ = "STUB: not implemented"
	// Use the lease ID as the file name, so it's guaranteed unique.
	return "", 0, nil
}

// Only create a new file if it doesn't exist already
// (version = 0), to avoid two processes using the
// same file name. Since we use the lease ID, this should never happen.

// Our context was canceled as we were sending
// a creation request. We don't know if it
// succeeded or not. In any case, let's try to
// delete the node, so we don't leave an orphan
// node behind for *leaseTTL time.

// The key already exists, that should not happen.

// The key was created.

// waitOnLastRev waits on all revisions of the files in the provided
// directory that have revisions smaller than the provided revision.
// It returns true only if there is no more other older files.
func (s *etcdtopo) waitOnLastRev(ctx context.Context, cli *clientv3.Client, nodePath string, revision int64) (bool, error) {
	_ = "STUB: not implemented"
	// Get the keys that are blocking us, if any.
	return false, nil
}

// No older key, we're done waiting.

// Wait for release on blocking key. Cancel the watch when we
// exit this function.

// There might still be older keys,
// but not this one.

// The Watch stopped, we're not sure if there are more items.

// etcdLockDescriptor implements topoclient.LockDescriptor.
type etcdLockDescriptor struct {
	s       *etcdtopo
	leaseID clientv3.LeaseID
}

// TryLock is part of the topoclient.Conn interface.
func (s *etcdtopo) TryLock(ctx context.Context, dirPath, contents string) (topoclient.LockDescriptor, error) {
	_ = "STUB: not implemented"
	// We list all the entries under dirPath
	return *new(topoclient.LockDescriptor), nil
}

// We need to return the right error codes, like
// topoclient.ErrNoNode and topoclient.ErrInterrupted, and the
// easiest way to do this is to return convertError(err).
// It may lose some of the context, if this is an issue,
// maybe logging the error would work here.

// If there is a folder '/locks' with some entries in it then we can assume that someone else already has a lock.
// Throw error in this case

// everything is good let's acquire the lock.

// Lock is part of the topoclient.Conn interface.
func (s *etcdtopo) Lock(ctx context.Context, dirPath, contents string) (topoclient.LockDescriptor, error) {
	_ = "STUB: not implemented"
	// We list the directory first to make sure it exists.
	return *new(topoclient.LockDescriptor), nil
}

/*full*/
// We need to return the right error codes, like
// topoclient.ErrNoNode and topoclient.ErrInterrupted, and the
// easiest way to do this is to return convertError(err).
// It may lose some of the context, if this is an issue,
// maybe logging the error would work here.

// LockWithTTL is part of the topoclient.Conn interface.
func (s *etcdtopo) LockWithTTL(ctx context.Context, dirPath, contents string, ttl time.Duration) (topoclient.LockDescriptor, error) {
	_ = "STUB: not implemented"
	// We list the directory first to make sure it exists.
	return *new(topoclient.LockDescriptor), nil
}

/*full*/
// We need to return the right error codes, like
// topoclient.ErrNoNode and topoclient.ErrInterrupted, and the
// easiest way to do this is to return convertError(err).
// It may lose some of the context, if this is an issue,
// maybe logging the error would work here.

// LockName is part of the topoclient.Conn interface.
func (s *etcdtopo) LockName(ctx context.Context, dirPath, contents string) (topoclient.LockDescriptor, error) {
	_ = "STUB: not implemented"
	return *new(topoclient.LockDescriptor), nil
}

// LockNameWithTTL is part of the topoclient.Conn interface.
func (s *etcdtopo) LockNameWithTTL(ctx context.Context, dirPath, contents string, ttl time.Duration) (topoclient.LockDescriptor, error) {
	_ = "STUB: not implemented"
	return *new(topoclient.LockDescriptor), nil
}

// TryLockName is part of the topoclient.Conn interface.
// It combines the fail-fast semantics of TryLock with LockName's ability to
// lock paths that don't exist. It checks if a lock already exists at the
// named lock path, and if so returns an error immediately.
func (s *etcdtopo) TryLockName(ctx context.Context, dirPath, contents string) (topoclient.LockDescriptor, error) {
	_ = "STUB: not implemented"
	// Check if a lock already exists by listing the locks directory
	return *new(topoclient.LockDescriptor), nil
}

// No lock exists, proceed with acquiring the named lock

// lock is used by both Lock() and primary election.
func (s *etcdtopo) lock(ctx context.Context, nodePath, contents string, ttl int) (topoclient.LockDescriptor, error) {
	_ = "STUB: not implemented"
	return *new(topoclient.LockDescriptor), nil
}

// Get a lease, set its KeepAlive.

// Drain the lease keepAlive channel, we're not
// interested in its contents.

// Create an ephemeral node in the locks directory.

// Wait until all older nodes in the locks directory are gone.

// We had an error waiting on the last node.
// Revoke our lease, this will delete the file.

// No more older nodes, we're it!

// Check is part of the topoclient.LockDescriptor interface.
// We use KeepAliveOnce to make sure the lease is still active and well.
func (ld *etcdLockDescriptor) Check(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Unlock is part of the topoclient.LockDescriptor interface.
func (ld *etcdLockDescriptor) Unlock(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// TryLockWithLease is part of the topoclient.Conn interface.
// It uses an atomic compare-version=0 transaction to create a lease-backed key,
// ensuring fail-fast semantics without a TOCTOU race.
func (s *etcdtopo) TryLockWithLease(ctx context.Context, key, contents string, ttl time.Duration) (topoclient.LockDescriptor, error) {
	_ = "STUB: not implemented"
	return *new(topoclient.LockDescriptor), nil
}

// Grant a lease with the requested TTL and start KeepAlive.

// Drain the KeepAlive channel so the etcd client continues sending
// lease renewals. If this channel fills up, renewals stop and the
// lease expires. We don't need the response contents.

// Atomically create the key. The compare-version=0 condition ensures
// that if another holder already created the key, this transaction
// fails immediately — no TOCTOU race.

// Key already exists — another holder has the lock. So, we revoke our
// own lease.

// RevokeLockWithLease is part of the topoclient.Conn interface.
// It forcefully removes the ephemeral lock at the given key by revoking its lease.
func (s *etcdtopo) RevokeLockWithLease(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

// No lock exists

// Key exists but has no lease

// Revoking the lease automatically deletes the ephemeral key, so a
// subsequent TryLockWithLease on the same key will succeed.
