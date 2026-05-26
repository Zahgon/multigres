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

package ast

// IsBeginStatement returns true if the statement is a BEGIN or START TRANSACTION.
func IsBeginStatement(stmt Stmt) bool { _ = "STUB: not implemented"; return false }

// IsCommitStatement returns true if the statement is a COMMIT.
func IsCommitStatement(stmt Stmt) bool { _ = "STUB: not implemented"; return false }

// IsRollbackStatement returns true if the statement is a ROLLBACK.
// Note: This does not include ROLLBACK TO SAVEPOINT.
func IsRollbackStatement(stmt Stmt) bool { _ = "STUB: not implemented"; return false }

// IsAllowedInAbortedTransaction returns true if the statement may proceed when
// the transaction is in the aborted (failed) state. PostgreSQL allows ROLLBACK,
// ROLLBACK TO SAVEPOINT, and COMMIT (which it converts to ROLLBACK with a
// WARNING) in this state. All other statements are rejected with SQLSTATE 25P02.
func IsAllowedInAbortedTransaction(stmt Stmt) bool { _ = "STUB: not implemented"; return false }

// ExtractTablesUsed walks the AST and returns deduplicated, schema-qualified
// table names from all RangeVar nodes. CTE names are excluded since they are
// virtual tables, not real ones. Returns nil for statements that don't
// reference tables (SET, SHOW, BEGIN, etc.).
func ExtractTablesUsed(stmt Stmt) []string { _ = "STUB: not implemented"; return nil }

// Single pass: collect CTE names and RangeVar references together,
// then filter CTE references from the result.

// Remove CTE references (unqualified names that match a CTE).
