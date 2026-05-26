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

package executor

import (
	"github.com/multigres/multigres/go/common/sqltypes"
)

// ScanRow scans values from a query result row into the provided destinations.
// Each destination should be a pointer to a supported type (bool, string, int, int32, int64, float64, time.Time).
func ScanRow(row *sqltypes.Row, dests ...any) error { _ = "STUB: not implemented"; return nil }

// ScanSingleRow is a convenience function that scans the first row of a result.
// Returns an error if the result has no rows.
func ScanSingleRow(result *sqltypes.Result, dests ...any) error {
	_ = "STUB: not implemented"
	return nil
}

// scanValue scans a single value into the destination.
func scanValue(val []byte, dest any) error {
	_ = "STUB: not implemented"
	// Handle nullable pointer types first (e.g., **string, **int)
	// These allow detecting NULL values by checking if the pointer is nil
	return nil
}

// Handle non-nullable types

// Handle NULL values - leave the destination unchanged
// (similar to sql.Scanner behavior with default values)

// PostgreSQL returns "t" or "f" for boolean values

// PostgreSQL timestamp formats from the PostgreSQL documentation:
// https://www.postgresql.org/docs/current/datatype-datetime.html#DATATYPE-DATETIME-OUTPUT
// We support the ISO 8601 style formats commonly returned by PostgreSQL.

// getValue is a helper that validates the row and column, then scans the value into dest.
func getValue[T any](row *sqltypes.Row, col int) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// GetString extracts a string value from a row at the given column index.
func GetString(row *sqltypes.Row, col int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetBool extracts a boolean value from a row at the given column index.
func GetBool(row *sqltypes.Row, col int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil

	// GetInt extracts an integer value from a row at the given column index.
}

func GetInt(row *sqltypes.Row, col int) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// GetInt32 extracts an int32 value from a row at the given column index.
func GetInt32(row *sqltypes.Row, col int) (int32, error) { _ = "STUB: not implemented"; return 0, nil }

// GetInt64 extracts an int64 value from a row at the given column index.
func GetInt64(row *sqltypes.Row, col int) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// ParseFloat64 parses a string as a float64.
// This is useful for parsing numeric values that were cast to text in SQL.
func ParseFloat64(s string) (float64, error) { _ = "STUB: not implemented"; return 0, nil }
