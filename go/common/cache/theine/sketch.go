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

type CountMinSketch struct {
	Table      []uint64
	Additions  uint
	SampleSize uint
	BlockMask  uint
}

func NewCountMinSketch() *CountMinSketch { _ = "STUB: not implemented"; return nil }

// indexOf return table index and counter index together
func (s *CountMinSketch) indexOf(h uint64, block uint64, offset uint8) (uint, uint) {
	_ = "STUB: not implemented"
	return 0, 0
}

// max block + 7(8 * 8 bytes), fit 64 bytes cache line

func (s *CountMinSketch) inc(index uint, offset uint) bool { _ = "STUB: not implemented"; return false }

func (s *CountMinSketch) Add(h uint64) bool { _ = "STUB: not implemented"; return false }

func (s *CountMinSketch) reset() { _ = "STUB: not implemented"; return }

func (s *CountMinSketch) count(h uint64, block uint64, offset uint8) uint {
	_ = "STUB: not implemented"
	return 0
}

func (s *CountMinSketch) Estimate(h uint64) uint { _ = "STUB: not implemented"; return 0 }

func next2Power(x uint) uint { _ = "STUB: not implemented"; return 0 }

func (s *CountMinSketch) EnsureCapacity(size uint) { _ = "STUB: not implemented"; return }

func spread(h uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func rehash(h uint64) uint64 { _ = "STUB: not implemented"; return 0 }
