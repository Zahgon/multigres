// Copyright 2025 Supabase, Inc.
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

/*
Package event provides a reflect-based framework for low-frequency global
dispatching of events, which are values of any arbitrary type, to a set of
listener functions, which are usually registered by plugin packages during init().

Listeners should do work in a separate goroutine if it might block. Dispatch
should be called synchronously to make sure work enters the listener's work
queue before moving on. After Dispatch returns, the listener is responsible for
arranging to flush its work queue before program termination if desired.

For example, any package can define an event type:

	package mypackage

	type MyEvent struct {
		field1, field2 string
	}

Then, any other package (e.g. a plugin) can listen for those events:

	package myplugin

	import (
		"event"
		"mypackage"
	)

	func onMyEvent(ev mypackage.MyEvent) {
		// do something with ev
	}

	func init() {
		event.AddListener(onMyEvent)
	}

Any registered listeners that accept a single argument of type MyEvent will
be called when a value of type MyEvent is dispatched:

	package myotherpackage

	import (
		"event"
		"mypackage"
	)

	func InMediasRes() {
		ev := mypackage.MyEvent{
			field1: "foo",
			field2: "bar",
		}

		event.Dispatch(ev)
	}

In addition, listener functions that accept an interface type will be called
for any dispatched value that implements the specified interface. A listener
that accepts `any` will be called for every event type. Listeners can also
accept pointer types, but they will only be called if the dispatch site calls
Dispatch() on a pointer.
*/
package event

import (
	"reflect"
	"sync"
)

var (
	listenersMutex sync.Mutex // protects listeners and interfaces
	listeners      = make(map[reflect.Type][]any)
	interfaces     = make([]reflect.Type, 0)
)

// BadListenerError is raised via panic() when AddListener is called with an
// invalid listener function.
type BadListenerError string

func (why BadListenerError) Error() string { _ = "STUB: not implemented"; return "" }

// AddListener registers a listener function that will be called when a matching
// event is dispatched. The type of the function's first (and only) argument
// declares the event type (or interface) to listen for.
func AddListener(fn any) { _ = "STUB: not implemented"; return }

// check that the function type is what we think: # of inputs/outputs, etc.
// panic if conditions not met (because it's a programming error to have that happen)

// the first input parameter is the event

// keep a list of listeners for each event type

// if eventType is an interface, store it in a separate list
// so we can check non-interface objects against all interfaces

// Dispatch sends an event to all registered listeners that were declared
// to accept values of the event's type, or interfaces that the value implements.
func Dispatch(ev any) { _ = "STUB: not implemented"; return }

// Build a map of listeners to be called while holding the lock

// also check if the type implements any of the registered interfaces

// Updater is an interface that events can implement to combine updating and
// dispatching into one call.
type Updater interface {
	// Update is called by DispatchUpdate() before the event is dispatched.
	Update(update any)
}

// DispatchUpdate calls Update() on the event and then dispatches it. This is a
// shortcut for combining updates and dispatches into a single call.
func DispatchUpdate(ev Updater, update any) { _ = "STUB: not implemented"; return }
