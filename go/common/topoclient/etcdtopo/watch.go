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

	"github.com/multigres/multigres/go/common/topoclient"
)

// Watch is part of the topoclient.Conn interface.
func (s *etcdtopo) Watch(ctx context.Context, filePath string) (*topoclient.WatchData, <-chan *topoclient.WatchData, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Get the initial version of the file

// Generic error.

// Node doesn't exist.

// ModRevision is used for the topoclient.Version value as we get the new Revision value back
// when updating the file/key within a transaction in file.go and so this is the opaque
// version that we can use to enforce serializabile writes for the file/key.

// Create an outer context that will be canceled on return and will cancel all inner watches.

// Create a context, will be used to cancel the watch on retry.

// Create the Watcher.  We start watching from the response we
// got, not from the file original version, as the server may
// not have that much history.

// Create the notifications channel, send updates to it.

// This includes context cancellation errors.

// Cancel inner context on retry and create new one.

// Final notification.

// Node is gone, send a final notice.

// WatchRecursive is part of the topoclient.Conn interface.
func (s *etcdtopo) WatchRecursive(ctx context.Context, dirpath string) ([]*topoclient.WatchDataRecursive, <-chan *topoclient.WatchDataRecursive, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Get the initial version of the file

// Create an outer context that will be canceled on return and will cancel all inner watches.

// Create a context, will be used to cancel the watch on retry.

// Create the Watcher.  We start watching from the response we
// got, not from the file original version, as the server may
// not have that much history.

// Create the notifications channel, send updates to it.

// This includes context cancellation errors.

// Cancel inner context on retry and create new one.

// Final notification.
