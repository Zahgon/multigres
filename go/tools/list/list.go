// Copyright 2025 Supabase, Inc.
// Copyright 2009 The Go Authors. All rights reserved.
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

// Package list is the standard library's 'container/list', but using Generics
// for performance.
package list

import "sync/atomic"

// Element is an element of a linked list.
type Element[T any] struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	next, prev *Element[T]

	// The list to which this element belongs.
	list *List[T]

	// The value stored with this element.
	Value T
}

// Next returns the next list element or nil.
func (e *Element[T]) Next() *Element[T] { _ = "STUB: not implemented"; return nil }

// Prev returns the previous list element or nil.
func (e *Element[T]) Prev() *Element[T] { _ = "STUB: not implemented"; return nil }

// List represents a doubly linked list.
// The zero value for List is an empty list ready to use.
type List[T any] struct {
	root Element[T] // sentinel list element, only &root, root.prev, and root.next are used
	len  atomic.Int64
}

// Init initializes or clears list l.
func (l *List[T]) Init() *List[T] { _ = "STUB: not implemented"; return nil }

func (l *List[T]) Len() int { _ = "STUB: not implemented"; return 0 }

// New returns an initialized list.
func New[T any]() *List[T] { _ = "STUB: not implemented"; return nil }

// Front returns the first element of list l or nil if the list is empty.
func (l *List[T]) Front() *Element[T] { _ = "STUB: not implemented"; return nil }

// Back returns the last element of list l or nil if the list is empty.
func (l *List[T]) Back() *Element[T] { _ = "STUB: not implemented"; return nil }

// insert inserts e after at, increments l.len, and returns e.
func (l *List[T]) insert(e, at *Element[T]) *Element[T] { _ = "STUB: not implemented"; return nil }

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *List[T]) insertValue(v T, at *Element[T]) *Element[T] {
	_ = "STUB: not implemented"
	return nil
}

// remove removes e from its list, decrements l.len
func (l *List[T]) remove(e *Element[T]) { _ = "STUB: not implemented"; return }

// avoid memory leaks
// avoid memory leaks

// Remove removes e from l if e is an element of list l.
// It returns the element value e.Value.
// The element must not be nil.
func (l *List[T]) Remove(e *Element[T]) { _ = "STUB: not implemented"; return }

// if e.list == l, l must have been initialized when e was inserted
// in l or l == nil (e is a zero Element) and l.remove will crash

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (l *List[T]) PushFront(v T) *Element[T] { _ = "STUB: not implemented"; return nil }

func (l *List[T]) PushFrontValue(v *Element[T]) { _ = "STUB: not implemented"; return }

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *List[T]) PushBack(v T) *Element[T] { _ = "STUB: not implemented"; return nil }

func (l *List[T]) PushBackValue(v *Element[T]) { _ = "STUB: not implemented"; return }
