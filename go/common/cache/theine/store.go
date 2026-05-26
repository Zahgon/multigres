// Copyright 2026 Supabase, Inc.
// Copyright 2023 The Vitess Authors.
// Copyright 2023 Yiling-J
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

package theine

import (
	"hash/maphash"
	"sync"
	"sync/atomic"

	"github.com/gammazero/deque"

	"github.com/multigres/multigres/go/common/cache/theine/bf"
)

const (
	MaxReadBuffSize  = 64
	MinWriteBuffSize = 4
	MaxWriteBuffSize = 1024
)

type RemoveReason uint8

const (
	REMOVED RemoveReason = iota
	EVICTED
	EXPIRED
)

type Shard[K cachekey, V any] struct {
	hashmap    map[K]*Entry[K, V]
	doorkeeper *bf.Bloomfilter
	deque      *deque.Deque[*Entry[K, V]]
	group      *Group[K, V]
	qsize      uint
	qlen       int
	counter    uint
	mu         sync.RWMutex
}

func NewShard[K cachekey, V any](qsize uint, doorkeeper bool) *Shard[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (s *Shard[K, V]) set(key K, entry *Entry[K, V]) { _ = "STUB: not implemented"; return }

func (s *Shard[K, V]) get(key K) (entry *Entry[K, V], ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (s *Shard[K, V]) delete(entry *Entry[K, V]) bool { _ = "STUB: not implemented"; return false }

func (s *Shard[K, V]) len() int { _ = "STUB: not implemented"; return 0 }

type Metrics struct {
	evicted atomic.Int64
	hits    atomic.Int64
	misses  atomic.Int64
}

func (m *Metrics) Evicted() int64 { _ = "STUB: not implemented"; return 0 }

func (m *Metrics) Hits() int64 { _ = "STUB: not implemented"; return 0 }

func (m *Metrics) Misses() int64 { _ = "STUB: not implemented"; return 0 }

func (m *Metrics) Accesses() int64 { _ = "STUB: not implemented"; return 0 }

type cachekey interface {
	comparable
	Hash() uint64
	Hash2() (uint64, uint64)
}

type HashKey256 [32]byte

func (h HashKey256) Hash() uint64 { _ = "STUB: not implemented"; return 0 }

func (h HashKey256) Hash2() (uint64, uint64) { _ = "STUB: not implemented"; return 0, 0 }

var cacheStringHashSeed = maphash.MakeSeed()

type StringKey string

func (h StringKey) Hash() uint64 { _ = "STUB: not implemented"; return 0 }

func (h StringKey) Hash2() (uint64, uint64) { _ = "STUB: not implemented"; return 0, 0 }

type cacheval interface {
	CachedSize(alloc bool) int64
}

type Store[K cachekey, V cacheval] struct {
	Metrics   Metrics
	OnRemoval func(K, V, RemoveReason)

	entryPool    sync.Pool
	writebuf     chan WriteBufItem[K, V]
	policy       *TinyLfu[K, V]
	readbuf      *Queue[ReadBufItem[K, V]]
	shards       []*Shard[K, V]
	cap          uint
	shardCount   uint
	writebufsize int64
	tailUpdate   bool
	doorkeeper   bool

	mlock       sync.Mutex
	readCounter atomic.Uint32
	open        atomic.Bool
}

func NewStore[K cachekey, V cacheval](maxsize int64, doorkeeper bool) *Store[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store[K, V]) EnsureOpen() { _ = "STUB: not implemented"; return }

func (s *Store[K, V]) getFromShard(key K, hash uint64, shard *Shard[K, V], epoch uint32) (V, bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (s *Store[K, V]) Get(key K, epoch uint32) (V, bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (s *Store[K, V]) GetOrLoad(key K, epoch uint32, load func() (V, error)) (V, bool, error) {
	_ = "STUB: not implemented"
	return *new(V), false, nil
}

func (s *Store[K, V]) setEntry(shard *Shard[K, V], cost int64, epoch uint32, entry *Entry[K, V]) {
	_ = "STUB: not implemented"
	return
}

// cost larger than deque size, send to policy directly

func (s *Store[K, V]) setInternal(key K, value V, cost int64, epoch uint32) (*Shard[K, V], *Entry[K, V], bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}

func (s *Store[K, V]) Set(key K, value V, cost int64, epoch uint32) bool {
	_ = "STUB: not implemented"
	return false
}

type dequeKV[K cachekey, V cacheval] struct {
	k K
	v V
}

func (s *Store[K, V]) processDeque(shard *Shard[K, V], epoch uint32) {
	_ = "STUB: not implemented"
	return
}

// send to slru

// double check because entry maybe removed already by Delete API

func (s *Store[K, V]) Delete(key K) { _ = "STUB: not implemented"; return }

func (s *Store[K, V]) Len() int { _ = "STUB: not implemented"; return 0 }

func (s *Store[K, V]) UsedCapacity() int { _ = "STUB: not implemented"; return 0 }

func (s *Store[K, V]) MaxCapacity() int {
	_ = "STUB: not implemented"

	// spread hash before get index
	return 0
}

func (s *Store[K, V]) index(key K) (uint64, int) { _ = "STUB: not implemented"; return 0, 0 }

func (s *Store[K, V]) postDelete(entry *Entry[K, V]) { _ = "STUB: not implemented"; return }

// remove entry from cache/policy/timingwheel and add back to pool
func (s *Store[K, V]) removeEntry(entry *Entry[K, V], reason RemoveReason) {
	_ = "STUB: not implemented"
	return
}

// already removed from shard map

func (s *Store[K, V]) drainRead() { _ = "STUB: not implemented"; return }

func (s *Store[K, V]) maintenanceItem(item WriteBufItem[K, V]) { _ = "STUB: not implemented"; return }

// lock free because store API never read/modify entry metadata

func (s *Store[K, V]) maintenance() { _ = "STUB: not implemented"; return }

func (s *Store[K, V]) Range(epoch uint32, f func(key K, value V) bool) {
	_ = "STUB: not implemented"
	return
}

func (s *Store[K, V]) Close() { _ = "STUB: not implemented"; return }
