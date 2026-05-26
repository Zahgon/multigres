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

// Package memorytopo contains an implementation of the topoclient.Factory /
// topoclient.Conn interfaces based on an in-memory tree of data.
// It is constructed with an immutable set of cells.
package memorytopo

import (
	"context"
	"errors"
	"regexp"
	"sync"
	"sync/atomic"
	"time"

	"github.com/multigres/multigres/go/common/topoclient"
)

const (
	// Path components
	electionsPath = "elections"
)

var ErrConnectionClosed = errors.New("connection closed")

const (
	// UnreachableServerAddr is a sentinel value for CellInfo.ServerAddr.
	// If a memorytopo topoclient.Conn is created with this serverAddr then every
	// method on that conn which takes a context will simply block until the
	// context finishes, and return ctx.Err(), in order to simulate an
	// unreachable local cell for testing.
	UnreachableServerAddr = "unreachable"
)

// Operation is one of the operations defined by topoclient.Conn
type Operation int

// The following is the list of topoclient.Conn operations
const (
	ListDir = Operation(iota)
	Create
	Update
	Get
	List
	Delete
	Lock
	TryLock
	Watch
	WatchRecursive
	NewLeaderParticipation
	Close
)

// Factory is a memory-based implementation of topoclient.Factory.  It
// takes a file-system like approach, with directories at each level
// being an actual directory node. This is meant to be closer to
// file-system like servers, like ZooKeeper or Chubby. etcd or Consul
// implementations would be closer to a node-based implementation.
//
// It contains a single tree of nodes. Each cell topoclient.Conn will use
// a subdirectory in that tree.
type Factory struct {
	// mu protects the following fields.
	mu sync.Mutex
	// cells is the toplevel map that has one entry per cell.
	cells map[string]*node
	// generation is used to generate unique incrementing version
	// numbers.  We want a global counter so when creating a file,
	// then deleting it, then re-creating it, we don't restart the
	// version at 1. It is initialized with a random number,
	// so if we have two implementations, the numbers won't match.
	generation uint64
	// err is used for testing purposes to force queries / watches
	// to return the given error
	err error
	// operationErrors is used for testing purposes to fake errors from
	// operations and paths matching the spec
	operationErrors map[Operation][]errorSpec
	// callstats allows us to keep track of how many topoclient.conn calls
	// we make (Create, Get, Update, Delete, List, ListDir, etc).
	// TODO: Implement stats
	// callstats *stats.CountersWithMultiLabels
}

type errorSpec struct {
	op          Operation
	pathPattern *regexp.Regexp
	err         error
	callCount   int
	maxCalls    int // 0 means unlimited, >0 means fail only this many times
}

// Create is part of the topoclient.Factory interface.
func (f *Factory) Create(cell, root string, serverAddrs []string) (topoclient.Conn, error) {
	_ = "STUB: not implemented"
	return *new(topoclient.Conn), nil
}

//  note (root, doesn't matter for the in memory topo, hence we don't use it).

// SetError forces the given error to be returned from all calls and propagates
// the error to all active watches.
func (f *Factory) SetError(err error) { _ = "STUB: not implemented"; return }

// func (f *Factory) GetCallStats() *stats.CountersWithMultiLabels {
//	return f.callstats
//}

// Lock blocks all requests to the topo and is exposed to allow tests to
// simulate an unresponsive topo server
func (f *Factory) Lock() {
	_ = "STUB: not implemented"

	// Unlock unblocks all requests to the topo and is exposed to allow tests to
	// simulate an unresponsive topo server
	return
}

func (f *Factory) Unlock() {
	_ = "STUB: not implemented"

	// conn implements the topoclient.Conn interface. It remembers the cell and serverAddr,
	// and points at the Factory that has all the data.
	return
}

type conn struct {
	factory     *Factory
	cell        string
	serverAddrs []string
	closed      atomic.Bool
}

var _ topoclient.Conn = (*conn)(nil)

// dial returns immediately, unless the conn points to the sentinel
// UnreachableServerAddr, in which case it will block until the context expires.
func (c *conn) dial(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Close is part of the topoclient.Conn interface.
func (c *conn) Close() error {
	_ = "STUB: not implemented"
	// c.factory.callstats.Add([]string{"Close"}, 1)
	return nil
}

type watch struct {
	contents  chan *topoclient.WatchData
	recursive chan *topoclient.WatchDataRecursive
	lock      chan string
	cancel    context.CancelFunc // Used to forcibly close the watch
}

// node contains a directory or a file entry.
// Exactly one of contents or children is not nil.
type node struct {
	name     string
	version  uint64
	contents []byte
	children map[string]*node

	// parent is a pointer to the parent node.
	// It is set to nil in toplevel and cell node.
	parent *node

	// watches is a map of all watches for this node.
	watches map[int]watch

	// lock is nil when the node is not locked.
	// otherwise it has a channel that is closed by unlock.
	lock chan struct{}

	// lockContents is the contents of the locks.
	// For regular locks, it has the contents that was passed in.
	// For primary election, it has the id of the election leader.
	lockContents string

	// lockTTLTimer is the timer for auto-expiring locks with TTL.
	// It is nil if no TTL was set or the lock has been unlocked.
	lockTTLTimer *time.Timer
}

func (n *node) isDirectory() bool { _ = "STUB: not implemented"; return false }

// fullPath returns the full path of this node from the cell root.
// It builds the path by walking up the parent chain.
func (n *node) fullPath() string { _ = "STUB: not implemented"; return "" }

func (n *node) recurseContents(callback func(n *node)) { _ = "STUB: not implemented"; return }

func (n *node) propagateRecursiveWatch(ev *topoclient.WatchDataRecursive) {
	_ = "STUB: not implemented"
	return
}

var (
	nextWatchIndex   = 0
	nextWatchIndexMu sync.Mutex
)

func (n *node) addWatch(w watch) int { _ = "STUB: not implemented"; return 0 }

// PropagateWatchError propagates the given error to all watches on this node
// and recursively applies to all children
func (n *node) PropagateWatchError(err error) { _ = "STUB: not implemented"; return }

// CloseWatches closes all watch channels for the given path and its children.
// This simulates what happens when etcd compacts history or a watch is forcibly cancelled.
// It's useful for testing how clients handle watch channel closures.
func (f *Factory) CloseWatches(cell, path string) { _ = "STUB: not implemented"; return }

// Cancel all watches on this node and its children

// AddCell dynamically registers a new cell in the in-memory topology.
// This enables tests to simulate cells appearing after server startup,
// e.g. when multiorch registers cells after multigateway has already started.
func (f *Factory) AddCell(ctx context.Context, ts topoclient.Store, cell string) error {
	_ = "STUB: not implemented"
	return nil
}

// NewServerAndFactory returns a new MemoryTopo and the backing factory for all
// the cells. It will create one cell for each parameter passed in.  It will log.Exit out
// in case of a problem.
func NewServerAndFactory(ctx context.Context, cells ...string) (topoclient.Store, *Factory) {
	_ = "STUB: not implemented"
	return *new(topoclient.Store), nil
}

// NewServerAndFactoryWithConfig is like NewServerAndFactory but allows specifying a custom TopoConfig.
// This is useful for tests that need to customize lock timeouts.
func NewServerAndFactoryWithConfig(ctx context.Context, config *topoclient.TopoConfig, cells ...string) (topoclient.Store, *Factory) {
	_ = "STUB: not implemented"
	return *new(topoclient.Store), nil
}

// callstats:       stats.NewCountersWithMultiLabels("", "", []string{"Call"}),

/*root*/ /*serverAddrs*/

// Create cell with mock server addresses for testing

// NewServer returns the new server
func NewServer(ctx context.Context, cells ...string) topoclient.Store {
	_ = "STUB: not implemented"
	return *new(topoclient.Store)
}

func (f *Factory) getNextVersion() uint64 { _ = "STUB: not implemented"; return 0 }

func (f *Factory) newFile(name string, contents []byte, parent *node) *node {
	_ = "STUB: not implemented"
	return nil
}

func (f *Factory) newDirectory(name string, parent *node) *node {
	_ = "STUB: not implemented"
	return nil
}

func (f *Factory) nodeByPath(cell, filePath string) *node { _ = "STUB: not implemented"; return nil }

// Skip empty parts, usually happens at the end.

// This is a file.

// Path doesn't exist.

func (f *Factory) getOrCreatePath(cell, filePath string) *node {
	_ = "STUB: not implemented"
	return nil
}

// Skip empty parts, usually happens at the end.

// This is a file.

// Path doesn't exist, create it.

// recursiveDelete deletes a node and its parent directory if empty.
func (f *Factory) recursiveDelete(n *node) { _ = "STUB: not implemented"; return }

func (f *Factory) AddOperationError(op Operation, pathPattern string, err error) {
	_ = "STUB: not implemented"
	return
}

// unlimited

// AddOneTimeOperationError adds an error that will only be returned once
func (f *Factory) AddOneTimeOperationError(op Operation, pathPattern string, err error) {
	_ = "STUB: not implemented"
	return
}

// only fail once

// ClearOperationErrors clears all operation errors for testing purposes.
func (f *Factory) ClearOperationErrors() { _ = "STUB: not implemented"; return }

func (f *Factory) getOperationError(op Operation, path string) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if this error should still be returned

// Increment call count
