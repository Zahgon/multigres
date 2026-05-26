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

package multigateway

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/multigres/multigres/go/common/topoclient"

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
)

// CellPoolerDiscovery is a discovery service that watches for multipoolers
// in a single cell using topology watches and maintains a list of available poolers.
type CellPoolerDiscovery struct {
	// Configuration
	topoStore topoclient.Store
	cell      string // The cell this watcher is monitoring
	logger    *slog.Logger

	// Callbacks for notifying about pooler changes (may be nil)
	onPoolerChanged func(pooler *clustermetadatapb.MultiPooler)
	onPoolerRemoved func(pooler *clustermetadatapb.MultiPooler)

	// Control
	ctx        context.Context
	cancelFunc context.CancelFunc
	wg         sync.WaitGroup

	// State (protected by mu)
	// Lock order: acquire this AFTER GlobalPoolerDiscovery.mu (never before)
	mu          sync.Mutex
	poolers     map[string]*topoclient.MultiPoolerInfo // pooler ID -> pooler info
	lastRefresh time.Time
}

// NewCellPoolerDiscovery creates a new pooler discovery service for a single cell.
func NewCellPoolerDiscovery(ctx context.Context, topoStore topoclient.Store, cell string, logger *slog.Logger) *CellPoolerDiscovery {
	_ = "STUB: not implemented"
	return nil
}

// Start begins the discovery process using topology watch.
func (pd *CellPoolerDiscovery) Start() { _ = "STUB: not implemented"; return }

// Context cancelled

// Establish watch and process changes

// Get connection for the cell

// Start watching the poolers directory
// This matches the PoolersPath constant from store.go

// Process initial values

// Reset backoff after watch has been stable for 30s

// Process changes as they come in

// Process the change - this handles both updates and deletions.
// Deletions come as events with Err set to NoNode error.

// Stop stops the discovery service.
func (pd *CellPoolerDiscovery) Stop() { _ = "STUB: not implemented"; return }

// processInitialPoolers processes the initial set of poolers from the watch
func (pd *CellPoolerDiscovery) processInitialPoolers(initial []*topoclient.WatchDataRecursive) {
	_ = "STUB: not implemented"
	return
}

// Save old poolers to detect removals (for watch reconnection)

// Process initial pooler data

// Parse the pooler from watch data

// Notify about removed poolers (existed before but not in new state)

// Notify about all current poolers

// processPoolerChange processes a single pooler change from the watch
func (pd *CellPoolerDiscovery) processPoolerChange(watchData *topoclient.WatchDataRecursive) {
	_ = "STUB: not implemented"
	return
}

// Check if this is an error event

// Check if it's a deletion event (NoNode error) for a pooler path.
// TopoError is returned as a value type, so we need to use errors.Is with a pointer.

// Log other watch errors (non-deletion errors)

// Parse the pooler from watch data

// This is a non-pooler file, skip

// If this is a PRIMARY, evict any other PRIMARY for the same TableGroup/Shard.
// This handles the failover case where a new PRIMARY comes up but the old
// crashed PRIMARY's record is still present.

// Check if this is a new pooler

// GetPoolersForAdmin returns a list of all discovered poolers in this cell.
// This is intended for admin/status pages, not the hot query path.
// Poolers are sorted by name for consistent display order.
func (pd *CellPoolerDiscovery) GetPoolersForAdmin() []*clustermetadatapb.MultiPooler {
	_ = "STUB: not implemented"
	return nil
}

// Collect and sort pooler IDs for consistent ordering

// LastRefresh returns the timestamp of the last successful refresh.
func (pd *CellPoolerDiscovery) LastRefresh() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// PoolerCount returns the current number of discovered poolers.
func (pd *CellPoolerDiscovery) PoolerCount() int { _ = "STUB: not implemented"; return 0 }

// parsePoolerFromWatchData parses a MultiPooler from watch data
func (pd *CellPoolerDiscovery) parsePoolerFromWatchData(watchData *topoclient.WatchDataRecursive) (*topoclient.MultiPoolerInfo, error) {
	_ = "STUB: not implemented"
	// Only process files that end with "Pooler" (the actual pooler data files)
	return nil, nil
}

// Not a pooler file, skip

// If Contents is nil, this might be a deletion

// Parse the protobuf data

// extractPoolerIDFromPath extracts the pooler ID from a watch path.
// The path format is: "poolers/{pooler_id}/Pooler"
func (pd *CellPoolerDiscovery) extractPoolerIDFromPath(path string) string {
	_ = "STUB: not implemented"
	// Expected format: "poolers/{pooler_id}/Pooler"
	return ""
}

// The pooler ID is the middle part(s) - everything between "poolers/" and "/Pooler"

// evictConflictingPrimary removes any existing PRIMARY pooler for the same
// TableGroup/Shard that has a different pooler ID. This handles the failover
// case where a new PRIMARY comes up but the old crashed PRIMARY's record is
// still in the discovery cache.
// Caller must hold pd.mu.
func (pd *CellPoolerDiscovery) evictConflictingPrimary(newPoolerID, tableGroup, shard string) {
	_ = "STUB: not implemented"
	return
}

// Skip if it's the same pooler (update case)

// Check if this is a PRIMARY for the same TableGroup/Shard

// Cell returns the cell this discovery is watching.
func (pd *CellPoolerDiscovery) Cell() string {
	_ = "STUB: not implemented"

	// PoolerChangeListener receives notifications about pooler discovery changes.
	// Implementations can use this to maintain connections to discovered poolers.
	return ""
}

type PoolerChangeListener interface {
	// OnPoolerChanged is called when a pooler is added or updated.
	// For new poolers, this creates a connection. For existing poolers,
	// this may recreate the connection if the address changed.
	OnPoolerChanged(pooler *clustermetadatapb.MultiPooler)
	// OnPoolerRemoved is called when a pooler is removed from discovery.
	OnPoolerRemoved(pooler *clustermetadatapb.MultiPooler)
}

// poolerNotification represents a notification to be delivered to listeners.
type poolerNotification struct {
	pooler         *clustermetadatapb.MultiPooler
	isRemoval      bool
	targetListener PoolerChangeListener // nil = broadcast to all

	// Special: listener registration
	isListenerRegistration bool
	newListener            PoolerChangeListener
}

// GlobalPoolerDiscovery orchestrates multiple CellPoolerDiscovery instances,
// one per cell. It watches for cell changes from global etcd and creates/removes
// cell watchers as cells appear/disappear.
type GlobalPoolerDiscovery struct {
	// Configuration
	topoStore topoclient.Store
	localCell string // The cell this multigateway is running in (for cell affinity)
	logger    *slog.Logger

	// Control
	ctx        context.Context
	cancelFunc context.CancelFunc
	wg         sync.WaitGroup

	// State (protected by mu)
	// Lock order: acquire this BEFORE CellPoolerDiscovery.mu
	mu              sync.Mutex
	cellWatchers    map[string]*CellPoolerDiscovery // cell name -> cell watcher
	lastCellRefresh time.Time                       // when cells were last discovered/refreshed

	// Listeners for pooler changes (protected by listenersMu)
	// Lock order: can acquire notificationsMu while holding this
	listenersMu sync.Mutex
	listeners   []PoolerChangeListener

	// Notification queue (protected by notificationsMu)
	// LOCK ORDERING: This is the INNERMOST lock - never acquire other locks while holding it.
	// It CAN be acquired while holding mu, listenersMu, or CellPoolerDiscovery.mu.
	notificationsMu sync.Mutex
	notifications   []poolerNotification
	notifySignal    chan struct{} // buffered(1) to wake up processor goroutine
}

// NewGlobalPoolerDiscovery creates a new global pooler discovery service.
// The localCell parameter indicates which cell this multigateway is running in,
// which will be used for cell affinity when selecting poolers.
func NewGlobalPoolerDiscovery(ctx context.Context, topoStore topoclient.Store, localCell string, logger *slog.Logger) *GlobalPoolerDiscovery {
	_ = "STUB: not implemented"
	return nil
}

// Start begins the discovery process by watching the cells directory in the
// global topology. When cells are created or removed, it starts or stops
// the corresponding CellPoolerDiscovery watchers.
func (gd *GlobalPoolerDiscovery) Start() {
	_ = "STUB: not implemented"
	// Start notification processor goroutine
	return
}

// Context cancelled

// Watch for cells and manage cell watchers

// Process initial cells

// Reset backoff after watch has been stable for 30s

// Watch for cell changes (additions/removals)

// processInitialCells processes the initial set of cells from the watch.
func (gd *GlobalPoolerDiscovery) processInitialCells(initial []*topoclient.WatchDataRecursive) {
	_ = "STUB: not implemented"
	return
}

// Log discovered cells

// processCellChange handles a single cell change event from the watch.
func (gd *GlobalPoolerDiscovery) processCellChange(event *topoclient.WatchDataRecursive) {
	_ = "STUB: not implemented"
	return
}

// Deletion: NoNode error signals the cell was removed

// Addition: start a watcher if we don't have one

// extractCellFromPath extracts the cell name from a cells watch path.
// Handles both relative paths (memorytopo: "cells/zone1/Cell") and
// absolute paths with root prefix (etcd: "/multigres/global/cells/zone1/Cell").
func extractCellFromPath(watchPath string) string { _ = "STUB: not implemented"; return "" }

// startCellWatcher starts a CellPoolerDiscovery for the given cell.
// Caller must hold gd.mu.
func (gd *GlobalPoolerDiscovery) startCellWatcher(cell string) { _ = "STUB: not implemented"; return }

// Stop stops the global discovery service and all cell watchers.
func (gd *GlobalPoolerDiscovery) Stop() {
	_ = "STUB: not implemented"

	// Stop all cell watchers
	return
}

// processNotifications runs in a background goroutine and processes queued notifications.
func (gd *GlobalPoolerDiscovery) processNotifications() { _ = "STUB: not implemented"; return }

// Drain and process all pending notifications

// Buffer empty, go back to waiting

// Process notifications

// Add listener in serial order with notifications

// Targeted replay - only call specific listener

// Broadcast - re-read listeners to include recently-registered ones

// queueNotification adds a notification to the queue and signals the processor.
func (gd *GlobalPoolerDiscovery) queueNotification(notif poolerNotification) {
	_ = "STUB: not implemented"
	return
}

// Wake up processor (non-blocking - buffered channel)

// Already signaled, processor will drain buffer

// RegisterListener adds a listener for pooler change notifications.
// The listener will immediately receive OnPoolerChanged for all currently
// known poolers, then continue to receive updates as poolers change.
func (gd *GlobalPoolerDiscovery) RegisterListener(listener PoolerChangeListener) {
	_ = "STUB: not implemented"
	// Collect current state and queue registration+replay atomically
	// while holding gd.mu to prevent poolers from changing between
	// collection and queuing
	return
}

// Collect replay notifications

// Queue: (1) add listener, (2) replay state
// This must happen while still holding gd.mu to ensure atomicity

// Signal processor outside locks

// notifyPoolerChanged notifies all listeners that a pooler was added or updated.
func (gd *GlobalPoolerDiscovery) notifyPoolerChanged(pooler *clustermetadatapb.MultiPooler) {
	_ = "STUB: not implemented"
	return
}

// Broadcast to all

// notifyPoolerRemoved notifies all listeners that a pooler was removed.
func (gd *GlobalPoolerDiscovery) notifyPoolerRemoved(pooler *clustermetadatapb.MultiPooler) {
	_ = "STUB: not implemented"
	return
}

// Broadcast to all

// PoolerCount returns the total number of discovered poolers across all cells.
// LastCellRefresh returns when cells were last discovered or refreshed.
// Returns zero time if initial cell discovery has not completed yet.
func (gd *GlobalPoolerDiscovery) LastCellRefresh() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (gd *GlobalPoolerDiscovery) PoolerCount() int { _ = "STUB: not implemented"; return 0 }

// CellStatusInfo contains status information for a single cell's discovery.
type CellStatusInfo struct {
	Cell        string
	LastRefresh time.Time
	Poolers     []*clustermetadatapb.MultiPooler
}

// GetCellStatusesForAdmin returns status information for each cell.
// This is intended for admin/status pages, not the hot query path.
// Cells are sorted alphabetically for consistent display order.
func (gd *GlobalPoolerDiscovery) GetCellStatusesForAdmin() []CellStatusInfo {
	_ = "STUB: not implemented"
	return nil
}

// Collect and sort cell names for consistent ordering
