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

package queryregistry

// durationBucketsNs are the upper-inclusive bucket boundaries (in ns) for the
// per-fingerprint duration histogram. The layout matches the OTel metric
// `mg.gateway.query.duration` — keeping them aligned means a query landing
// in bucket N here lands in bucket N on the metric, which makes registry
// percentiles and Prometheus percentiles agree.
//
// Boundaries: 0.5ms, 1ms, 5ms, 10ms, 50ms, 100ms, 500ms, 1s, 5s, 10s.
// Values above 10s land in the implicit overflow bucket (index numHistBuckets-1).
var durationBucketsNs = [10]int64{
	500_000,        // 0.5 ms
	1_000_000,      // 1 ms
	5_000_000,      // 5 ms
	10_000_000,     // 10 ms
	50_000_000,     // 50 ms
	100_000_000,    // 100 ms
	500_000_000,    // 500 ms
	1_000_000_000,  // 1 s
	5_000_000_000,  // 5 s
	10_000_000_000, // 10 s
}

// numHistBuckets is the number of histogram buckets including the +Inf overflow.
const numHistBuckets = len(durationBucketsNs) + 1

// bucketIndex returns the index of the bucket that contains durNs, using
// upper-inclusive boundaries. Anything above the last finite boundary lands
// in the overflow bucket at index numHistBuckets-1.
func bucketIndex(durNs int64) int { _ = "STUB: not implemented"; return 0 }

// percentileNs returns the duration (ns) at which the cumulative count crosses
// p (a fraction in [0,1]) of the total, using linear interpolation within the
// crossing bucket. Returns 0 if there are no samples.
//
// counts is laid out as [bucket0, bucket1, ..., overflow].
func percentileNs(counts [numHistBuckets]uint64, p float64) int64 {
	_ = "STUB: not implemented"
	return 0
}

// Interpolate inside this bucket.

// Overflow bucket: no upper bound; pin to the last finite boundary.

// Numerically possible only if total > 0 but we walked past the end.
