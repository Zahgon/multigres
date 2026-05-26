// Copyright 2026 Supabase, Inc.
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

package bf

// doorkeeper is a small bloom-filter-based cache admission policy
type Bloomfilter struct {
	Filter            bitvector // our filter bit vector
	M                 uint32    // size of bit vector in bits
	K                 uint32    // distinct hash functions needed
	FalsePositiveRate float64
	Capacity          int
}

func New(falsePositiveRate float64) *Bloomfilter { _ = "STUB: not implemented"; return nil }

func (d *Bloomfilter) EnsureCapacity(capacity int) { _ = "STUB: not implemented"; return }

// in bits

func (d *Bloomfilter) Exist(h uint64) bool { _ = "STUB: not implemented"; return false }

// insert inserts the byte array b into the bloom filter.  Returns true if the value
// was already considered to be in the bloom filter.
func (d *Bloomfilter) Insert(h uint64) bool { _ = "STUB: not implemented"; return false }

// Reset clears the bloom filter
func (d *Bloomfilter) Reset() { _ = "STUB: not implemented"; return }

// Internal routines for the bit vector
type bitvector []uint64

func newbv(size uint32) bitvector { _ = "STUB: not implemented"; return *new(bitvector) }

func (b bitvector) get(bit uint32) uint { _ = "STUB: not implemented"; return 0 }

// set bit 'bit' in the bitvector d and return previous value
func (b bitvector) getset(bit uint32) uint { _ = "STUB: not implemented"; return 0 }

// return the integer >= i which is a power of two
func nextPowerOfTwo(i uint32) uint32 { _ = "STUB: not implemented"; return 0 }
