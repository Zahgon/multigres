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

package topoclient

// ErrorCode is the error code for topo errors.
type ErrorCode int

// The following is the list of error codes.
const (
	NodeExists = ErrorCode(iota)
	NoNode
	NodeNotEmpty
	Timeout
	Interrupted
	BadVersion
	PartialResult
	NoUpdateNeeded
	NoImplementation
	NoReadOnlyImplementation
	ResourceExhausted
	BadInput
)

// TopoError represents a topo error.
type TopoError struct {
	Code    ErrorCode
	Message string
}

// NewError creates a new topo error.
func NewError(code ErrorCode, node string) error { _ = "STUB: not implemented"; return nil }

// Error satisfies error.
func (e TopoError) Error() string { _ = "STUB: not implemented"; return "" }

// Is implements error comparison for errors.Is.
func (e TopoError) Is(target error) bool { _ = "STUB: not implemented"; return false }
