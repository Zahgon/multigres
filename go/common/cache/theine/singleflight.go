// Copyright 2026 Supabase, Inc.
// Copyright 2023 The Vitess Authors.
// Copyright 2023 Yiling-J
// Copyright 2013 The Go Authors. All rights reserved.
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

// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package theine

import (
	"errors"
	"sync"
	"sync/atomic"
)

// errGoexit indicates the runtime.Goexit was called in
// the user given function.
var errGoexit = errors.New("runtime.Goexit was called")

// A panicError is an arbitrary value recovered from a panic
// with the stack trace during the execution of given function.
type panicError struct {
	value any
	stack []byte
}

// Error implements error interface.
func (p *panicError) Error() string { _ = "STUB: not implemented"; return "" }

func newPanicError(v any) error { _ = "STUB: not implemented"; return nil }

// The first line of the stack trace is of the form "goroutine N [status]:"
// but by the time the panic reaches Do the goroutine may no longer exist
// and its status will have changed. Trim out the misleading line.

// call is an in-flight or completed singleflight.Do call
type call[V any] struct {
	// These fields are written once before the WaitGroup is done
	// and are only read after the WaitGroup is done.
	val V
	err error

	wg sync.WaitGroup

	// These fields are read and written with the singleflight
	// mutex held before the WaitGroup is done, and are read but
	// not written after the WaitGroup is done.
	dups atomic.Int32
}

// Group represents a class of work and forms a namespace in
// which units of work can be executed with duplicate suppression.
type Group[K comparable, V any] struct {
	m        map[K]*call[V] // lazily initialized
	mu       sync.Mutex     // protects m
	callPool sync.Pool
}

func NewGroup[K comparable, V any]() *Group[K, V] { _ = "STUB: not implemented"; return nil }

// Result holds the results of Do, so they can be passed
// on a channel.
type Result struct {
	Val    any
	Err    error
	Shared bool
}

// Do executes and returns the results of the given function, making
// sure that only one execution is in-flight for a given key at a
// time. If a duplicate comes in, the duplicate caller waits for the
// original to complete and receives the same results.
// The return value shared indicates whether v was given to multiple callers.
func (g *Group[K, V]) Do(key K, fn func() (V, error)) (v V, err error, shared bool) {
	_ = "STUB: not implemented"
	return *new(V), nil, false
}

// assign value/err before put back to pool to avoid race

// doCall handles the single call for a key.
func (g *Group[K, V]) doCall(c *call[V], key K, fn func() (V, error)) {
	_ = "STUB: not implemented"
	return
}

// use double-defer to distinguish panic from runtime.Goexit,
// more details see https://golang.org/cl/134395

// the given function invoked runtime.Goexit

// Ideally, we would wait to take a stack trace until we've determined
// whether this is a panic or a runtime.Goexit.
//
// Unfortunately, the only way we can distinguish the two is to see
// whether the recover stopped the goroutine from terminating, and by
// the time we know that, the part of the stack trace relevant to the
// panic has been discarded.
