// PostgreSQL Database Management System
// (also known as Postgres, formerly known as Postgres95)
//
//	Portions Copyright (c) 2026, Supabase, Inc
//
//	Portions Copyright (c) 1996-2025, PostgreSQL Global Development Group
//
//	Portions Copyright (c) 1994, The Regents of the University of California
//
// Permission to use, copy, modify, and distribute this software and its
// documentation for any purpose, without fee, and without a written agreement
// is hereby granted, provided that the above copyright notice and this
// paragraph and the following two paragraphs appear in all copies.
//
// Hand-written lexer for the replication-protocol command grammar.
// Ported from postgres/src/backend/replication/repl_scanner.l (PG 17.6).

package replparser

import (
	"errors"

	"github.com/multigres/multigres/go/common/parser/ast"
)

var (
	errUnterminatedString = errors.New("unterminated quoted string")
	errUnterminatedIdent  = errors.New("unterminated quoted identifier")
	errInvalidRecptr      = errors.New("invalid streaming start location")
)

// replKeywords - ported from repl_scanner.l:120-140 (physical-only keywords
// intentionally omitted).
var replKeywords = map[string]int{
	"IDENTIFY_SYSTEM":         K_IDENTIFY_SYSTEM,
	"READ_REPLICATION_SLOT":   K_READ_REPLICATION_SLOT,
	"SHOW":                    K_SHOW,
	"TIMELINE":                K_TIMELINE,
	"START_REPLICATION":       K_START_REPLICATION,
	"CREATE_REPLICATION_SLOT": K_CREATE_REPLICATION_SLOT,
	"DROP_REPLICATION_SLOT":   K_DROP_REPLICATION_SLOT,
	"ALTER_REPLICATION_SLOT":  K_ALTER_REPLICATION_SLOT,
	"LOGICAL":                 K_LOGICAL,
	"SLOT":                    K_SLOT,
	"RESERVE_WAL":             K_RESERVE_WAL,
	"TEMPORARY":               K_TEMPORARY,
	"TWO_PHASE":               K_TWO_PHASE,
	"EXPORT_SNAPSHOT":         K_EXPORT_SNAPSHOT,
	"NOEXPORT_SNAPSHOT":       K_NOEXPORT_SNAPSHOT,
	"USE_SNAPSHOT":            K_USE_SNAPSHOT,
	"WAIT":                    K_WAIT,
}

type replLexer struct {
	input  []byte
	pos    int
	err    error
	result ast.Stmt
}

func newReplLexer(input string) *replLexer { _ = "STUB: not implemented"; return nil }

// Err returns the first lexing/parsing error encountered, or nil.
func (l *replLexer) Err() error {
	_ = "STUB: not implemented"

	// setResult is invoked from the grammar's top-level action.
	return nil
}

func (l *replLexer) setResult(s ast.Stmt) {
	_ = "STUB: not implemented"

	// Lex implements the goyacc lexer interface.
	// Mirrors repl_scanner.l rules at :120-208.
	return
}

func (l *replLexer) Lex(lval *replYySymType) int {
	_ = "STUB: not implemented"
	// Skip whitespace (repl_scanner.l:76,142).
	return 0
}

// Decimal UCONST or %X/%X RECPTR starting with a digit.

// Identifier, keyword, or %X/%X RECPTR starting with a hex letter.

// Error implements the goyacc lexer interface.
func (l *replLexer) Error(s string) { _ = "STUB: not implemented"; return }

// scanSingleQuoted scans a single-quoted string literal.
// Ported from repl_scanner.l:158-176 ('...' with ” as escape).
func (l *replLexer) scanSingleQuoted(lval *replYySymType) int {
	_ = "STUB: not implemented"
	// consume opening '
	return 0
}

// scanDoubleQuoted scans a double-quoted identifier.
// Ported from repl_scanner.l:178-196 ("..." with "" as escape).
func (l *replLexer) scanDoubleQuoted(lval *replYySymType) int {
	_ = "STUB: not implemented"
	// consume opening "
	return 0
}

// scanNumberOrRecptr handles input starting with a decimal digit.
// Emits UCONST for a decimal run, or RECPTR if the run is followed by
// '/<hex>'. Ported from repl_scanner.l:144-156.
func (l *replLexer) scanNumberOrRecptr(lval *replYySymType) int {
	_ = "STUB: not implemented"
	return 0
}

// Not a RECPTR. PG's flex rule for UCONST matches only decimal
// digits, so back up to the longest leading run of [0-9] and emit
// that as UCONST. Anything after will be re-scanned as the next
// token.

// scanIdentOrRecptr handles input starting with a letter or underscore.
// Most inputs become IDENT/keyword, but a run of hex-only characters
// followed by '/<hex>' is a RECPTR (repl_scanner.l flex's longest-match
// behavior picks the RECPTR rule over the identifier rule in that case).
func (l *replLexer) scanIdentOrRecptr(lval *replYySymType) int { _ = "STUB: not implemented"; return 0 }

// Keyword lookup is case-insensitive against uppercase keys
// (repl_scanner.l matches IDENTIFY_SYSTEM case-insensitively via flex).

// Unquoted identifiers are downcased
// (repl_scanner.l:201 calls downcase_truncate_identifier).

// emitRecptr consumes the '/<hex>' tail after the caller has already
// scanned the leading hex run and verified the '/' follows.
func (l *replLexer) emitRecptr(lval *replYySymType, first string) int {
	_ = "STUB: not implemented"
	// consume '/'
	return 0
}

func isAllHexDigits(s string) bool { _ = "STUB: not implemented"; return false }

// isHexDigit reports whether c is a hex digit [0-9A-Fa-f].
// Ported from repl_scanner.l:98.
func isHexDigit(c byte) bool { _ = "STUB: not implemented"; return false }

// isIdentStart reports whether c can start an identifier.
// Ported from repl_scanner.l:100.
func isIdentStart(c byte) bool { _ = "STUB: not implemented"; return false }

// isIdentCont reports whether c can continue an identifier.
// Ported from repl_scanner.l:101.
func isIdentCont(c byte) bool { _ = "STUB: not implemented"; return false }
