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

// Package prototest provides test helpers for asserting equality of protobuf messages.
package prototest

import (
	"testing"

	"google.golang.org/protobuf/proto"
)

// AssertEqual asserts that two proto messages are equal using proto.Equal semantics.
// The test continues even if the assertion fails.
func AssertEqual(t *testing.T, expected, actual proto.Message, msgAndArgs ...any) {
	_ = "STUB: not implemented"
	return
}

// RequireEqual asserts that two proto messages are equal using proto.Equal semantics.
// The test stops immediately if the assertion fails.
func RequireEqual(t *testing.T, expected, actual proto.Message, msgAndArgs ...any) {
	_ = "STUB: not implemented"
	return
}

// RequireElementsMatch asserts that two slices of proto messages contain the same elements
// regardless of order, using proto.Equal semantics for comparison.
func RequireElementsMatch[T proto.Message](t *testing.T, expected, actual []T, msgAndArgs ...any) {
	_ = "STUB: not implemented"
	return
}

func formatMsg(msgAndArgs []any) string { _ = "STUB: not implemented"; return "" }
