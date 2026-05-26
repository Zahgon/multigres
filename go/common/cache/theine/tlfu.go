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
	"sync/atomic"
)

type TinyLfu[K cachekey, V any] struct {
	slru      *Slru[K, V]
	sketch    *CountMinSketch
	size      uint
	counter   uint
	total     atomic.Uint32
	hit       atomic.Uint32
	hr        float32
	threshold atomic.Int32
	lruFactor uint8
	step      int8
}

func NewTinyLfu[K cachekey, V any](size uint) *TinyLfu[K, V] { _ = "STUB: not implemented"; return nil }

// default threshold to -1 so all entries are admitted until cache is full

func (t *TinyLfu[K, V]) climb() { _ = "STUB: not implemented"; return }

// reset

func (t *TinyLfu[K, V]) Set(entry *Entry[K, V]) *Entry[K, V] { _ = "STUB: not implemented"; return nil }

func (t *TinyLfu[K, V]) Access(item ReadBufItem[K, V]) { _ = "STUB: not implemented"; return }

func (t *TinyLfu[K, V]) Remove(entry *Entry[K, V]) { _ = "STUB: not implemented"; return }

func (t *TinyLfu[K, V]) UpdateCost(entry *Entry[K, V], delta int64) {
	_ = "STUB: not implemented"
	return
}

func (t *TinyLfu[K, V]) EvictEntries() []*Entry[K, V] { _ = "STUB: not implemented"; return nil }

func (t *TinyLfu[K, V]) UpdateThreshold() { _ = "STUB: not implemented"; return }

// cache is not full
