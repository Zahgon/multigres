// Copyright 2026 Supabase, Inc.
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

package pgparity

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/multigres/multigres/go/test/endtoend/suiteutil"
)

// RecordResult holds the outcome of a single parsed record.
type RecordResult struct {
	LineNo int
	Kind   string // "statement" or "query"
	Pass   bool
	Reason string // failure reason, empty on pass
}

// FileResult aggregates per-record results for one test file on one target.
type FileResult struct {
	Total   int
	Passed  int
	Failed  int
	Records []RecordResult
}

// RunFile executes a parsed test file against a target and reports per-record
// results. A fresh database schema is used (public); callers are responsible
// for ensuring target state is clean before the run.
//
// Connection errors for individual records are reported as failures rather
// than aborting the file: this makes it easy to see how much of a file a
// target can handle. A hard connection failure at the start aborts the run.
func RunFile(ctx context.Context, tf *TestFile, target suiteutil.Target) (*FileResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// runRecord executes one parsed record and returns its result.
func runRecord(ctx context.Context, conn *pgx.Conn, rec Record) RecordResult {
	_ = "STUB: not implemented"
	return *new(RecordResult)
}

// Optional regex match against error message.

// executeQuery runs a query record against the connection and returns the
// result set flattened into one value per cell, normalized by the type string
// specified in the record (I=int, R=real, T=text).
func executeQuery(ctx context.Context, conn *pgx.Conn, rec Record) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// nCols tracks the actual column count from the database, not the
// declared type string length. Using the type string length would
// silently mis-reshape the row grid for rowsort if a test author wrote a
// too-short type string (e.g. `query II` for a 3-column SELECT) — rowsort
// would sort across row boundaries without any error.

// Apply sort mode: rowsort reorders whole rows before comparison, used
// when the SQL has no ORDER BY and the test doesn't care about order.

// formatValue converts a pgx-decoded value into its canonical string form
// based on the declared type char:
//
//	I — integer (NULL becomes "NULL")
//	R — real/float, formatted with three decimal places
//	T — text ("" becomes "(empty)", NULL becomes "NULL")
//
// Unknown type chars fall back to generic string formatting.
func formatValue(v any, typeChar byte) string { _ = "STUB: not implemented"; return "" }

// PostgreSQL NUMERIC type comes back as this struct. Convert via
// Float64Value and truncate to an integer for 'I' columns.

// bitSize=32 so rounding matches the original float32 precision;
// PG `real` / `float4` columns come back as float32 via pgx.

// Division and aggregates like avg() return NUMERIC in pgx. Render
// with the same 3-decimal convention as floats for comparability.

func toFloat64(v any) float64 { _ = "STUB: not implemented"; return 0 }

// rowsort sorts a flat slice of values as if it were a 2D grid of nCols
// columns per row, preserving per-row order within the result.
func rowsort(flat []string, nCols int) []string { _ = "STUB: not implemented"; return nil }

// Should not happen for well-formed queries, but don't panic in tests.

// compareQueryResult returns the empty string when values match the expected
// rows, or a human-readable failure reason otherwise.
//
// Expected rows may themselves be tab-split during parsing, so both sides are
// flat value-per-slot lists.
func compareQueryResult(values []string, rec Record) string {
	_ = "STUB: not implemented"
	// An empty expected section (no `----` block) means the test only
	// required the query to execute without error.
	return ""
}

// truncateList renders a slice for error messages, cutting off after a few
// elements so we don't dump megabytes into test logs.
func truncateList(xs []string) string { _ = "STUB: not implemented"; return "" }
