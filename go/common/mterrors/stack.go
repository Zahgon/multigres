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

package mterrors

/* This file is copied from https://github.com/pkg/errors/blob/v0.8.0/stack.go */

import (
	"fmt"
)

// Frame represents a program counter inside a stack frame.
type Frame uintptr

// pc returns the program counter for this frame;
// multiple frames may have the same PC value.
func (f Frame) pc() uintptr { _ = "STUB: not implemented"; return 0 }

// file returns the full path to the file that contains the
// function for this Frame's pc.
func (f Frame) file() string { _ = "STUB: not implemented"; return "" }

// line returns the line number of source code of the
// function for this Frame's pc.
func (f Frame) line() int { _ = "STUB: not implemented"; return 0 }

// Format formats the frame according to the fmt.Formatter interface.
//
//	%s    source file
//	%d    source line
//	%n    function name
//	%v    equivalent to %s:%d
//
// Format accepts flags that alter the printing of some verbs, as follows:
//
//	%+s   path of source file relative to the compile time GOPATH
//	%+v   equivalent to %+s:%d
func (f Frame) Format(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }

// StackTrace is stack of Frames from innermost (newest) to outermost (oldest).
type StackTrace []Frame

// Format format the stacktrace according to the fmt.Formatter interface.
//
//	%s    source file
//	%d    source line
//	%n    function name
//	%v    equivalent to %s:%d
//
// Format accepts flags that alter the printing of some verbs, as follows:
//
//	%+s   path of source file relative to the compile time GOPATH
//	%+v   equivalent to %+s:%d
func (st StackTrace) Format(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }

// stack represents a stack of program counters.
type stack []uintptr

func (s *stack) Format(st fmt.State, verb rune) { _ = "STUB: not implemented"; return }

func (s *stack) StackTrace() StackTrace { _ = "STUB: not implemented"; return *new(StackTrace) }

func callers() *stack { _ = "STUB: not implemented"; return nil }

// funcname removes the path prefix component of a function's name reported by func.Name().
func funcname(name string) string { _ = "STUB: not implemented"; return "" }
