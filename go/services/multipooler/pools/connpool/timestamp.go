// Copyright 2025 Supabase, Inc.
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

package connpool

import (
	"math"
	"sync/atomic"
	"time"
)

var monotonicRoot = time.Now()

// timestamp is a monotonic point in time, stored as a number of
// nanoseconds since the monotonic root. This type is only 8 bytes
// and hence can always be accessed atomically
type timestamp struct {
	nano atomic.Int64
}

// timestampExpired is a special value that means this timestamp is now past
// an arbitrary expiration point, and hence doesn't need to store
const timestampExpired = math.MaxInt64

// timestampBusy is a special value that means this timestamp no longer
// tracks an expiration point
const timestampBusy = math.MinInt64

// monotonicNow returns the current monotonic time as a time.Duration.
// This is a very efficient operation because time.Since performs direct
// subtraction of monotonic times without considering the wall clock times.
func monotonicNow() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

// monotonicFromTime converts a wall-clock time from time.Now into a
// monotonic timestamp.
// This is a very efficient operation because time.(*Time).Sub performs direct
// subtraction of monotonic times without considering the wall clock times.
func monotonicFromTime(now time.Time) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// set sets this timestamp to the given monotonic value
func (t *timestamp) set(mono time.Duration) { _ = "STUB: not implemented"; return }

// get returns the monotonic time of this timestamp as the number of nanoseconds
// since the monotonic root.
func (t *timestamp) get() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

// elapsed returns the number of nanoseconds that have passed since
// this timestamp was updated
func (t *timestamp) elapsed() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

// update sets this timestamp's value to the current monotonic time
func (t *timestamp) update() { _ = "STUB: not implemented"; return }

// borrow attempts to borrow this timestamp atomically.
// It only succeeds if we can ensure that nobody else has marked
// this timestamp as expired. When succeeded, the timestamp
// is cleared as "busy" as it no longer tracks an expiration point.
func (t *timestamp) borrow() bool { _ = "STUB: not implemented"; return false }

// expired attempts to atomically expire this timestamp.
// It only succeeds if we can ensure the timestamp hasn't been
// concurrently expired or borrowed.
func (t *timestamp) expired(now time.Duration, timeout time.Duration) bool {
	_ = "STUB: not implemented"
	return false
}
