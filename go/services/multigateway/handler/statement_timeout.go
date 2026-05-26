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

package handler

import (
	"time"

	"github.com/multigres/multigres/go/common/mterrors"
	"github.com/multigres/multigres/go/common/parser/ast"
)

// ResolveStatementTimeout returns the per-query directive if present, otherwise
// the effective timeout from the connection state (session override or default).
// A nil directive means no directive was found; a non-nil directive (including 0,
// which disables timeouts) takes priority over everything else.
func ResolveStatementTimeout(directive *time.Duration, effective time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// ParseStatementTimeoutDirective is a placeholder for per-query directive parsing.
// Per-query directives (e.g., /*mg+ STATEMENT_TIMEOUT_MS=500 */) will be supported
// in the future by parsing them in the SQL grammar (similar to Vitess).
// For now, this always returns nil (no directive found).
func ParseStatementTimeoutDirective(query ast.Stmt) *time.Duration {
	_ = "STUB: not implemented"

	// ParsePostgresInterval parses a PostgreSQL-style interval value for statement_timeout
	// into a time.Duration. Returns PgDiagnostic errors matching PostgreSQL's error format.
	// Supports:
	//   - Plain integers as milliseconds (e.g., "5000" -> 5s) — PostgreSQL's default unit
	//   - Go-compatible duration strings (e.g., "30s", "200ms", "1m")
	return nil
}

func ParsePostgresInterval(paramName, value string) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

// Try parsing as plain integer (milliseconds) first — this is the common PG case.

// Try Go duration format (e.g., "30s", "200ms", "1m").

// invalidParamError returns a PgDiagnostic for an invalid parameter value (SQLSTATE 22023).
func invalidParamError(paramName, value, hint string) *mterrors.PgDiagnostic {
	_ = "STUB: not implemented"
	return nil
}

// outOfRangeParamError returns a PgDiagnostic for an out-of-range parameter value (SQLSTATE 22023).
func outOfRangeParamError(paramName, value string) *mterrors.PgDiagnostic {
	_ = "STUB: not implemented"
	return nil
}

// formatDurationPg formats a time.Duration using PostgreSQL's GUC_UNIT_MS display
// convention. PostgreSQL picks the largest unit that divides evenly into the value:
//
//	0        → "0"
//	500ms    → "500ms"
//	5s       → "5s"
//	90s      → "1min 30s"  (but we use "90s" — PG only splits at clean boundaries)
//	60s      → "1min"
//	3600s    → "1h"
//
// Values that don't divide evenly into the next-larger unit stay in the smaller unit
// (e.g., 1500ms → "1500ms", not "1.5s").
func formatDurationPg(d time.Duration) string { _ = "STUB: not implemented"; return "" }
