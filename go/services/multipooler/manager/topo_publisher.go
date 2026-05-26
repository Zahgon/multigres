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
	"time"

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
)

const topoPublisherRetryInterval = 30 * time.Second

// topoRegistrar is the subset of topoclient.Store used by topoPublisher.
type topoRegistrar interface {
	RegisterMultiPooler(ctx context.Context, multipooler *clustermetadatapb.MultiPooler, allowUpdate bool) error
}

// topoPublisher eventually-consistently reflects the multipooler's in-memory state
// to etcd. Callers invoke Notify with the current state whenever it changes; a
// background goroutine wakes immediately to write it and also retries periodically
// so that writes missed during an etcd outage are recovered automatically.
type topoPublisher struct {
	logger     *slog.Logger
	topoClient topoRegistrar

	// wakeup is a size-1 buffered channel. A non-blocking send schedules a
	// publish without accumulating multiple pending signals.
	wakeup chan struct{}

	mu            sync.Mutex
	desired       *clustermetadatapb.MultiPooler // latest state that should be in etcd
	lastPublished *clustermetadatapb.MultiPooler // last state successfully written; nil if never written
}

func newTopoPublisher(logger *slog.Logger, topoClient topoRegistrar) *topoPublisher {
	_ = "STUB: not implemented"
	return nil
}

// Notify records mp as the desired topology state and schedules an immediate
// publish attempt. mp is cloned so the caller may reuse the original.
//
// ctx must carry an action lock (see AssertActionLockHeld). The action lock
// serialises state transitions, preventing concurrent calls from racing to
// overwrite each other's desired state.
func (tp *topoPublisher) Notify(ctx context.Context, mp *clustermetadatapb.MultiPooler) error {
	_ = "STUB: not implemented"
	return nil
}

// Non-blocking send: if the channel is already full, a publish is already
// pending and will pick up the latest desired state.

// Run is the background loop. It blocks until ctx is cancelled. Call it in a
// goroutine: go tp.Run(ctx).
func (tp *topoPublisher) Run(ctx context.Context) { _ = "STUB: not implemented"; return }

// run is the internal loop, accepting an injectable ticker channel so tests can
// drive retries without real clock time.
func (tp *topoPublisher) run(ctx context.Context, tickC <-chan time.Time) {
	_ = "STUB: not implemented"
	return
}

// publishIfNeeded writes the desired state to etcd if it differs from the last
// successfully published state. A no-op when state is already current.
func (tp *topoPublisher) publishIfNeeded(ctx context.Context) { _ = "STUB: not implemented"; return }
