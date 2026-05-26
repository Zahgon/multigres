// Copyright 2023 The Vitess Authors.
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

package utils

import (
	"context"
	"testing"
	"time"
)

// LeakCheckContext returns a Context that will be automatically cancelled at the end
// of this test. If the test has finished successfully, it will be checked for goroutine
// leaks after context cancellation.
func LeakCheckContext(t testing.TB) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// LeakCheckContextTimeout behaves like LeakCheckContext but the returned Context will
// be cancelled after `timeout`, or after the test finishes, whichever happens first.
func LeakCheckContextTimeout(t testing.TB, timeout time.Duration) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// EnsureNoLeaks checks for goroutine and socket leaks and fails the test if any are found.
func EnsureNoLeaks(t testing.TB) { _ = "STUB: not implemented"; return }

// GetLeaks checks for goroutine and socket leaks and returns an error if any are found.
// One use case is in TestMain()s to ensure that all tests are cleaned up.
func GetLeaks() error { _ = "STUB: not implemented"; return nil }

func ensureNoLeaks() error { _ = "STUB: not implemented"; return nil }

func ensureNoGoroutines() error { _ = "STUB: not implemented"; return nil }
