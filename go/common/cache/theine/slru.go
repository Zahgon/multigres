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

type Slru[K cachekey, V any] struct {
	probation *List[K, V]
	protected *List[K, V]
	maxsize   uint
}

func NewSlru[K cachekey, V any](size uint) *Slru[K, V] { _ = "STUB: not implemented"; return nil }

func (s *Slru[K, V]) insert(entry *Entry[K, V]) *Entry[K, V] { _ = "STUB: not implemented"; return nil }

func (s *Slru[K, V]) victim() *Entry[K, V] { _ = "STUB: not implemented"; return nil }

func (s *Slru[K, V]) access(entry *Entry[K, V]) { _ = "STUB: not implemented"; return }

func (s *Slru[K, V]) remove(entry *Entry[K, V]) { _ = "STUB: not implemented"; return }

func (s *Slru[K, V]) updateCost(entry *Entry[K, V], delta int64) { _ = "STUB: not implemented"; return }
