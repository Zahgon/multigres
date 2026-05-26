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

package recovery

import (
	"context"
	"log/slog"
	"sync"

	"github.com/multigres/multigres/go/common/topoclient"
	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	"github.com/multigres/multigres/go/services/multiorch/config"
	"github.com/multigres/multigres/go/services/multiorch/store"
	"github.com/multigres/multigres/go/tools/retry"
)

// PoolerWatcher watches etcd for topology changes and keeps the pooler store
// up-to-date. It mirrors the two-tier approach used by multigateway:
//
//  1. A global watcher monitors the cells/ directory to detect cells appearing
//     or disappearing.
//  2. For each cell, a per-cell watcher monitors the poolers/ directory.
//
// When a pooler event arrives, it is filtered in-memory against the engine's
// WatchTargets before the pooler store is updated. Newly discovered poolers
// are reported via the onNewPooler callback so the caller can start monitoring
// them (e.g. open a ManagerHealthStream stream).
type PoolerWatcher struct {
	topoStore   topoclient.Store
	targets     func() []config.WatchTarget // live accessor, same as Engine.shardWatchTargets
	store       *store.PoolerStore
	onNewPooler func(id *clustermetadatapb.ID) // called when a new pooler is first discovered
	logger      *slog.Logger

	// Control
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup

	// State: per-cell watchers (protected by mu)
	mu           sync.Mutex
	cellWatchers map[string]*cellPoolerWatcher
}

// NewPoolerWatcher creates a new PoolerWatcher.
// targets is a function that returns the current WatchTargets (consulted on every event).
// onNewPooler is called when a new pooler is discovered; the pooler is already present
// in the store when the callback fires.
func NewPoolerWatcher(
	ctx context.Context,
	topoStore topoclient.Store,
	targets func() []config.WatchTarget,
	poolerStore *store.PoolerStore,
	onNewPooler func(id *clustermetadatapb.ID),
	logger *slog.Logger,
) *PoolerWatcher {
	_ = "STUB: not implemented"
	return nil
}

// Start launches the global cell-watcher goroutine.
func (pw *PoolerWatcher) Start() { _ = "STUB: not implemented"; return }

// Stop cancels the watcher and waits for all goroutines to finish.
func (pw *PoolerWatcher) Stop() { _ = "STUB: not implemented"; return }

// Sync blocks until all events that were in-flight at the time of the call have
// been processed by every active cell watcher. It is intended for use in tests
// to replace time.Sleep calls after topology mutations.
func (pw *PoolerWatcher) Sync(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// watchCells establishes a WatchRecursive on the global cells/ directory.
// It starts per-cell watchers as cells appear and stops them when cells disappear.
func (pw *PoolerWatcher) watchCells(r *retry.Retry) { _ = "STUB: not implemented"; return }

// Process existing cells

// Reset backoff after 30s of stable watching

func (pw *PoolerWatcher) processInitialCells(initial []*topoclient.WatchDataRecursive) {
	_ = "STUB: not implemented"
	return
}

func (pw *PoolerWatcher) processCellEvent(event *topoclient.WatchDataRecursive) {
	_ = "STUB: not implemented"
	return
}

// startCellWatcher starts a per-cell pooler watcher. Caller must hold pw.mu.
func (pw *PoolerWatcher) startCellWatcher(cell string) { _ = "STUB: not implemented"; return }

// extractCellNameFromPath extracts the cell name from a cells/ watch path.
// Handles both relative paths (memorytopo: "cells/zone1/Cell") and
// absolute paths with root prefix (etcd: "/multigres/global/cells/zone1/Cell").
func extractCellNameFromPath(watchPath string) string { _ = "STUB: not implemented"; return "" }

// ---------------------------------------------------------------------------
// cellPoolerWatcher: per-cell pooler directory watcher
// ---------------------------------------------------------------------------

type cellPoolerWatcher struct {
	topoStore   topoclient.Store
	cell        string
	targets     func() []config.WatchTarget
	store       *store.PoolerStore
	onNewPooler func(id *clustermetadatapb.ID)
	logger      *slog.Logger

	ctx      context.Context
	cancel   context.CancelFunc
	wg       sync.WaitGroup
	syncChan chan chan struct{} // for Sync(); closed by the watchPoolers select loop
}

func newCellPoolerWatcher(
	ctx context.Context,
	topoStore topoclient.Store,
	cell string,
	targets func() []config.WatchTarget,
	poolerStore *store.PoolerStore,
	onNewPooler func(id *clustermetadatapb.ID),
	logger *slog.Logger,
) *cellPoolerWatcher {
	_ = "STUB: not implemented"
	return nil
}

func (cw *cellPoolerWatcher) start() { _ = "STUB: not implemented"; return }

func (cw *cellPoolerWatcher) stop() { _ = "STUB: not implemented"; return }

func (cw *cellPoolerWatcher) watchPoolers(r *retry.Retry) { _ = "STUB: not implemented"; return }

// Process the initial set of poolers

// Reset backoff after 30s of stable watching

// sync blocks until all events enqueued before this call have been processed,
// or until ctx is cancelled. It is intended for use in tests.
//
// It works by sending a sentinel channel into the same select loop that processes
// watch events. Since the select loop is sequential, closing the sentinel is
// guaranteed to happen only after all previously-enqueued events have been handled.
func (cw *cellPoolerWatcher) sync(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// handlePoolerEvent processes a single watch event for a pooler file.
func (cw *cellPoolerWatcher) handlePoolerEvent(wd *topoclient.WatchDataRecursive) {
	_ = "STUB: not implemented"

	// Deletions (NoNode) are intentionally ignored here: the pooler store
	// entry is left in place so that ongoing health checks can continue until
	// bookkeeping removes it after the configured unseen threshold.
	return
}

// Only handle files named "Pooler"

// Apply WatchTarget filter (in-memory, same semantics as old GetMultiPoolersByCell options)

// Use DoUpdate to atomically update only the topology metadata, preserving all
// health-check fields (timestamps, IsUpToDate, IsLastCheckValid, etc.) that may
// be written concurrently by the health check worker. A closure variable tracks
// whether the key existed so we know whether to treat this as a new pooler.

// New pooler — add to store and queue for immediate health check.

// matchesAnyTarget returns true if the pooler matches at least one of the
// configured WatchTargets.
func (cw *cellPoolerWatcher) matchesAnyTarget(pooler *clustermetadatapb.MultiPooler) bool {
	_ = "STUB: not implemented"
	return false
}
