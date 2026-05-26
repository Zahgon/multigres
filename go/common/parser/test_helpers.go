// PostgreSQL Database Management System
// (also known as Postgres, formerly known as Postgres95)
//
//  Portions Copyright (c) 2025, Supabase, Inc
//
//  Portions Copyright (c) 1996-2025, PostgreSQL Global Development Group
//
//  Portions Copyright (c) 1994, The Regents of the University of California
//
// Permission to use, copy, modify, and distribute this software and its
// documentation for any purpose, without fee, and without a written agreement
// is hereby granted, provided that the above copyright notice and this
// paragraph and the following two paragraphs appear in all copies.
//
// IN NO EVENT SHALL THE UNIVERSITY OF CALIFORNIA BE LIABLE TO ANY PARTY FOR
// DIRECT, INDIRECT, SPECIAL, INCIDENTAL, OR CONSEQUENTIAL DAMAGES, INCLUDING
// LOST PROFITS, ARISING OUT OF THE USE OF THIS SOFTWARE AND ITS
// DOCUMENTATION, EVEN IF THE UNIVERSITY OF CALIFORNIA HAS BEEN ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.
//
// THE UNIVERSITY OF CALIFORNIA SPECIFICALLY DISCLAIMS ANY WARRANTIES,
// INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY
// AND FITNESS FOR A PARTICULAR PURPOSE.  THE SOFTWARE PROVIDED HEREUNDER IS
// ON AN "AS IS" BASIS, AND THE UNIVERSITY OF CALIFORNIA HAS NO OBLIGATIONS TO
// PROVIDE MAINTENANCE, SUPPORT, UPDATES, ENHANCEMENTS, OR MODIFICATIONS.
//
/*
 * Common Test Helpers for PostgreSQL Parser Tests
 *
 * This file contains shared utility functions used across multiple test files
 * to reduce code duplication and improve maintainability.
 */

package parser

import (
	"testing"

	"github.com/multigres/multigres/go/common/parser/ast"
)

// ScanAllTokens scans all tokens from a lexer and returns them as a slice
// This is the consolidated version of the scanAllTokens function that was
// duplicated across multiple test files
func ScanAllTokens(t *testing.T, lexer *Lexer) []Token { _ = "STUB: not implemented"; return nil }

// AssertNoLexerErrors verifies that the lexer has no errors
func AssertNoLexerErrors(t *testing.T, lexer *Lexer) { _ = "STUB: not implemented"; return }

// AssertLexerHasErrors verifies that the lexer has errors
func AssertLexerHasErrors(t *testing.T, lexer *Lexer) { _ = "STUB: not implemented"; return }

// AssertTokenSequence tests that an input produces the expected sequence of token types
func AssertTokenSequence(t *testing.T, input string, expectedTypes []TokenType) {
	_ = "STUB: not implemented"
	return
}

// Last token should be EOF

// AssertTokenTypes verifies that tokens match expected types
func AssertTokenTypes(t *testing.T, tokens []Token, expectedTypes []TokenType) {
	_ = "STUB: not implemented"
	return
}

// AssertTokenValues verifies that tokens match expected string values
func AssertTokenValues(t *testing.T, tokens []Token, expectedValues []string) {
	_ = "STUB: not implemented"
	return
}

// AssertTokenTypesAndValues verifies both types and values
func AssertTokenTypesAndValues(t *testing.T, tokens []Token, expected []struct {
	Type  TokenType
	Value string
},
) {
	_ = "STUB: not implemented"
	return
}

// AssertParseSQL tests that SQL parses without error and returns expected number of statements
func AssertParseSQL(t *testing.T, sql string, expectedStmtCount int) []ast.Stmt {
	_ = "STUB: not implemented"
	return nil
}

// AssertRoundTripSQL tests that SQL can be parsed and deparsed consistently
func AssertRoundTripSQL(t *testing.T, sql string, expectedSQL string) {
	_ = "STUB: not implemented"
	// Parse the input
	return
}

// Deparse the statement

// Determine expected output

// Check if output matches expected

// Test round-trip parsing (re-parse the deparsed SQL)

// Verify statement types match

// Test stability - second deparse should match first

// AssertTokenPosition verifies token position information
func AssertTokenPosition(t *testing.T, token *Token, expectedPosition int) {
	_ = "STUB: not implemented"
	return
}

// AssertLexerContextError verifies specific error types in lexer context
func AssertLexerContextError(t *testing.T, lexer *Lexer, expectedErrorType LexerErrorType) {
	_ = "STUB: not implemented"
	return
}
