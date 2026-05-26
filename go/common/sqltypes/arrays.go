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

package sqltypes

// ParseTextArray parses a PostgreSQL text-format array literal (e.g. {foo,bar}) into a []string.
// Follows the PostgreSQL specification:
//   - Quoted elements preserve all content; backslash escapes apply inside quotes.
//   - Unquoted elements have leading/trailing whitespace trimmed; backslash escapes apply.
//   - Unquoted NULL (any case, no escapes) is SQL NULL — an error since it cannot be a string.
//   - Backslash-escaped values (e.g. \null) are never treated as NULL.
//
// Multi-dimensional arrays return an error. Malformed input returns an error.
func ParseTextArray(s string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// Skip leading whitespace before the element.

// Reached end after a comma: trailing delimiter.

// Consecutive or trailing delimiters produce an empty unquoted element.

// Unquoted NULL without backslash escapes is SQL NULL, which has no string representation.

// skip comma

// readQuotedArrayElem reads a double-quoted element starting at inner[i] (the opening `"`).
// Returns the element value and the position pointing at `,` or end of inner.
func readQuotedArrayElem(inner string, i int, orig string) (string, int, error) {
	_ = "STUB: not implemented"
	// skip opening `"`
	return "", 0, nil
}

// skip closing `"`
// After the closing quote only whitespace and then `,` or end is valid.

// readUnquotedArrayElem reads an unquoted element starting at inner[i].
// Processes backslash escapes and trims trailing whitespace (leading whitespace is
// consumed by the caller before invoking this function).
// Returns (element, hasEscapes, newPos, error) where newPos points at `,` or end of inner.
func readUnquotedArrayElem(inner string, i int, orig string) (string, bool, int, error) {
	_ = "STUB: not implemented"
	return "", false, 0, nil
}

// length of cur up to and including the last non-whitespace byte

// isArraySpace reports whether c is whitespace per PostgreSQL's scanner_isspace:
// space, tab, newline, carriage return, form feed, vertical tab.
func isArraySpace(c byte) bool { _ = "STUB: not implemented"; return false }
