// Copyright 2025 Supabase, Inc.
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

package ast

import (
	"time"
)

// Format codes for parameter encoding
const (
	TextFormat   = 0
	BinaryFormat = 1
)

// PostgreSQL epoch: 2000-01-01 00:00:00 UTC
var pgEpoch = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

// SubstituteParameters walks through the AST and replaces all ParamRef nodes
// with A_Const nodes containing the actual parameter values.
// This function creates a modified copy of the AST with parameters substituted.
//
// Parameters:
//   - stmt: The AST statement to process
//   - params: The parameter values as byte arrays (nil for NULL)
//   - paramFormats: Format codes for each parameter (0=text, 1=binary)
//   - paramTypes: PostgreSQL type OIDs for each parameter
//
// Returns the modified statement with parameters substituted, or an error if
// parameter parsing fails.
func SubstituteParameters(stmt Stmt, params [][]byte, paramFormats []int16, paramTypes []uint32) (Stmt, error) {
	_ = "STUB: not implemented"
	return *new(Stmt), nil
}

// Create parameter value cache to convert params to A_Const nodes

// Use Rewrite to walk the AST and replace ParamRef nodes

// Parameter numbers are 1-based

// If parameter index is out of bounds, leave ParamRef unchanged
// No need to traverse children of ParamRef

// Continue traversal for other nodes

// createParameterValues converts parameter bytes to A_Const nodes
func createParameterValues(params [][]byte, formats []int16, paramTypes []uint32) ([]*A_Const, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Determine format for each parameter

// Handle NULL

// Get type OID if available

// Parse parameter based on format

// parseTextParameter parses a text format parameter into an A_Const node
func parseTextParameter(data []byte, typeOID uint32) (*A_Const, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate and parse as integer (bitSize 0 = platform int size)

// Validate it's a float

// Parse boolean

// 0 = unknown type, treat as text

// Bytea in text format

// Unknown type, treat as text

// parseBinaryParameter parses a binary format parameter into an A_Const node
func parseBinaryParameter(data []byte, typeOID uint32) (*A_Const, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Text in binary format is just UTF-8 bytes

// PostgreSQL timestamp: microseconds since 2000-01-01 00:00:00 UTC

// Return as string value that will be properly quoted

// Bytea in binary format: raw bytes, encode as hex string

// Unknown binary type - try to interpret as text
