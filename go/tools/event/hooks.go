// Copyright 2019 The Vitess Authors.
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

package event

import (
	"sync"
)

// Hooks holds a list of parameter-less functions to call whenever the set is
// triggered with Fire().
type Hooks struct {
	funcs []func()
	mu    sync.Mutex
}

// Add appends the given function to the list to be triggered.
func (h *Hooks) Add(f func()) { _ = "STUB: not implemented"; return }

// Fire calls all the functions in a given Hooks list. It launches a goroutine
// for each function and then waits for all of them to finish before returning.
// Concurrent calls to Fire() are serialized.
func (h *Hooks) Fire() { _ = "STUB: not implemented"; return }

// ErrorHooks holds a list of error-returning functions to call whenever the set
// is triggered with Fire().
type ErrorHooks struct {
	funcs []func() error
	mu    sync.Mutex
}

// Add appends the given function to the list to be triggered.
func (h *ErrorHooks) Add(f func() error) { _ = "STUB: not implemented"; return }

// Fire calls all the functions in a given ErrorHooks list in parallel.
// It returns immediately on the first error encountered (fail-fast).
// Concurrent calls to Fire() are serialized.
func (h *ErrorHooks) Fire() error { _ = "STUB: not implemented"; return nil }

// Close channel when all goroutines complete

// Return first error or nil if channel closes without error
