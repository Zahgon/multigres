// Copyright 2026 Supabase, Inc.
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

package benchmarking

import (
	"context"
	"time"
)

// fetchProfile fetches a pprof endpoint at addr and writes the bytes to
// outputPath. Endpoints that block for a window (e.g. CPU profile) are
// supported via the timeout — pass it large enough to cover the
// server-side window plus a small margin.
func fetchProfile(ctx context.Context, addr, urlPath string, timeout time.Duration, outputPath string) error {
	_ = "STUB: not implemented"
	return nil
}

// captureCPUProfile fetches a CPU profile of `seconds` duration from the
// given servenv HTTP address. Blocks for ~seconds.
func captureCPUProfile(ctx context.Context, addr string, seconds int, outputPath string) error {
	_ = "STUB: not implemented"
	return nil
}

// captureHeapProfile fetches a current heap (in-use space) profile snapshot.
// Returns immediately — this is a sample of live allocations.
func captureHeapProfile(ctx context.Context, addr, outputPath string) error {
	_ = "STUB: not implemented"
	return nil
}

// captureAllocsProfile fetches a cumulative allocation profile (all
// allocations since process start). Useful for spotting alloc hotspots.
func captureAllocsProfile(ctx context.Context, addr, outputPath string) error {
	_ = "STUB: not implemented"
	return nil
}

// captureGoroutineProfile fetches the current goroutine stack dump.
// Useful to see what's parked / contending.
func captureGoroutineProfile(ctx context.Context, addr, outputPath string) error {
	_ = "STUB: not implemented"
	return nil
}
