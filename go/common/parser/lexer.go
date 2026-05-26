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
 * PostgreSQL Parser Lexer - Core Lexer Implementation
 *
 * This file implements the core lexer structure and interface for the
 * PostgreSQL-compatible lexer. This provides the foundation for the
 * complete lexical analysis system.
 * Ported from postgres/src/backend/parser/scan.l and postgres/src/include/parser/scanner.h
 */

package parser

import (
	"github.com/multigres/multigres/go/common/parser/ast"
)

// Lexer represents the main lexer instance
// Equivalent to PostgreSQL's scanner state management
type Lexer struct {
	context   *ParseContext // Thread-safe unified context
	parseTree []ast.Stmt    // Parse tree result from parsing
}

// Import ast package (will be added at the top)
// "github.com/manangupta/multigres/go/common/parser/ast"

// NewLexer creates a new PostgreSQL-compatible lexer instance
// Equivalent to postgres/src/backend/parser/scan.l:1404 (scanner_init function)
func NewLexer(input string) *Lexer { _ = "STUB: not implemented"; return nil }

// NextToken returns the next token from the input stream
// This is the main lexer interface, equivalent to PostgreSQL's base_yylex
// Implements the same token filtering logic as PostgreSQL's base_yylex function
// Equivalent to postgres/src/backend/parser/parser.c:110-324 (base_yylex function)
func (l *Lexer) NextToken() *Token { _ = "STUB: not implemented"; return nil }

// Apply PostgreSQL's base_yylex token filtering
// This handles both keyword lookahead (WITH_LA, etc.) and Unicode token conversion (UIDENT→IDENT, USCONST→SCONST)

// applyTokenFilter implements PostgreSQL's base_yylex token filtering logic
// This function handles both keyword lookahead transformations and Unicode token processing
// Equivalent to postgres/src/backend/parser/parser.c:110-324 (base_yylex function)
func (l *Lexer) applyTokenFilter(token *Token) *Token {
	_ = "STUB: not implemented"
	// First level: Check if this token needs special processing
	// Based on PostgreSQL parser.c lines 138-161 and 252-254
	return nil
}

// Keywords that need lookahead - handle via existing checkLookaheadToken

// Unicode tokens that need conversion - implement PostgreSQL's UIDENT/USCONST processing
// Based on PostgreSQL parser.c lines 252-321

// No special processing needed

// processUnicodeToken handles UIDENT and USCONST token conversion following PostgreSQL's logic
// Implements the same logic as PostgreSQL's base_yylex function for UIDENT/USCONST cases
// Equivalent to postgres/src/backend/parser/parser.c:252-321
func (l *Lexer) processUnicodeToken(token *Token) *Token {
	_ = "STUB: not implemented"
	// Look ahead for UESCAPE token following PostgreSQL's pattern
	// postgres/src/backend/parser/parser.c:255-256
	return nil
}

// Default escape character (postgres/src/backend/parser/parser.c:303)

// Found UESCAPE - need to get the third token which should be SCONST
// postgres/src/backend/parser/parser.c:257-279

// Error getting UESCAPE value - use default escape character

// Invalid UESCAPE value - record error but continue with default

// Apply Unicode decoding using the escape character
// postgres/src/backend/parser/parser.c:284-306

// Fall back to original value

// If we found and processed UESCAPE, consume those tokens

// Convert token type and apply identifier truncation if needed
// postgres/src/backend/parser/parser.c:308-320

// Truncate identifier following PostgreSQL rules
// postgres/src/backend/parser/parser.c:310-314

// USCONST
// Convert to regular string constant
// postgres/src/backend/parser/parser.c:316-319

// peekUescapeValue peeks ahead to get the UESCAPE value (third token)
// Returns the escape string from the SCONST that follows UESCAPE
// Equivalent to postgres/src/backend/parser/parser.c:268-276
func (l *Lexer) peekUescapeValue() (string, error) {
	_ = "STUB: not implemented"
	// Save current position to restore later
	return "", nil
}

// Restore position after lookahead

// Skip whitespace to find UESCAPE keyword

// Skip the UESCAPE keyword

// Skip whitespace to find the string literal

// Check for string literal (SCONST)

// Parse the string literal to get the escape character

// skipIdentifier skips over an identifier for lookahead purposes
func (l *Lexer) skipIdentifier() {
	_ = "STUB: not implemented"
	// Skip first character
	return
}

// Skip rest of identifier

// consumeUescapeTokens consumes the UESCAPE keyword and string value tokens
// This is called when we've confirmed the UESCAPE sequence is valid
func (l *Lexer) consumeUescapeTokens() {
	_ = "STUB: not implemented"
	// Skip whitespace and consume UESCAPE keyword
	return
}

// Skip whitespace and consume the string literal

// Parse and consume the string literal

// Consume the SCONST token

// isValidUescapeChar checks if a character is valid as a Unicode escape character
// Equivalent to postgres/src/backend/parser/parser.c:351-362 (check_uescapechar)
func (l *Lexer) isValidUescapeChar(escape byte) bool {
	_ = "STUB: not implemented"
	// Invalid characters: hex digits, +, ', ", whitespace
	return false
}

// decodeUnicodeString decodes Unicode escape sequences in a string.
// Mirrors PostgreSQL's str_udeescape (postgres/src/backend/parser/parser.c:371-527)
// including surrogate-pair pairing, code-point range validation, and rejection
// of lone or out-of-order surrogates. Errors emitted here surface as
// `invalid Unicode surrogate pair` / `invalid Unicode escape value` /
// `invalid Unicode escape sequence` matching upstream wording.
func (l *Lexer) decodeUnicodeString(input string, escapeChar byte) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// non-zero while awaiting low surrogate; 0 otherwise

// High surrogate not followed by another \ escape sequence.

// Escape sequence. First handle doubled escape char: literal char.

// isHexSequence checks if a string contains only hexadecimal digits
func (l *Lexer) isHexSequence(s string) bool { _ = "STUB: not implemented"; return false }

// parseHexCodepoint parses a hex string into a Unicode codepoint
func (l *Lexer) parseHexCodepoint(hex string) (rune, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// truncateIdentifier truncates an identifier to PostgreSQL's maximum length
// Equivalent to postgres/src/backend/parser/scansup.c:truncate_identifier
func (l *Lexer) truncateIdentifier(ident string) string { _ = "STUB: not implemented"; return "" }

// PostgreSQL's maximum identifier length + 1

// Truncate and warn (in a real implementation, this would log a warning)

// nextTokenInternal is the internal token scanning function
func (l *Lexer) nextTokenInternal() (*Token, error) {
	_ = "STUB: not implemented"

	// Skip whitespace and comments
	return nil, nil
}

// Check for end of input

// Record start position for this token - equivalent to SET_YYLLOC()
// postgres/src/backend/parser/scan.l:105

// Save position for error reporting (near text extraction)

// State-based dispatch - PostgreSQL uses exclusive states
// postgres/src/backend/parser/scan.l:175-188

// Unicode strings

// StateXEU is now handled within string processing, not as a separate token state
// This should not be reached in normal operation

// scanInitialState handles scanning in the initial (default) state
// Equivalent to postgres/src/backend/parser/scan.l INITIAL state rules
func (l *Lexer) scanInitialState(startPos, startScanPos int) (*Token, error) {
	_ = "STUB: not implemented"
	// Peek at the current character to determine token type
	return nil, nil
}

// Character-based dispatch following PostgreSQL patterns
// postgres/src/backend/parser/scan.l:346-900+ (rules section)
// IMPORTANT: Order matters! Specific character cases must come before general isIdentStart()

// Numbers first - postgres/src/backend/parser/scan.l:430-500+

// Decimal point numbers (.5, .123, etc.) - postgres/src/backend/parser/scan.l:409

// Check for ".." operator FIRST before checking for digits

// Handle as DOT_DOT operator

// Otherwise, handle as operator/punctuation

// String literals - postgres/src/backend/parser/scan.l:286-318

// Extended string literals (E'...') - postgres/src/backend/parser/scan.l:275

// Otherwise, treat as identifier

// Bit string literals (B'...') - postgres/src/backend/parser/scan.l:264

// Otherwise, treat as identifier

// Hex string literals (X'...') - postgres/src/backend/parser/scan.l:268

// Otherwise, treat as identifier

// National character literals (N'...') - postgres/src/backend/parser/scan.l:272

// consume 'N'

// National character strings are treated like regular strings

// Otherwise, treat as identifier

// Unicode identifiers (U&"...") - postgres/src/backend/parser/scan.l:315

// consume 'U&'

// Unicode strings

// Otherwise, treat as identifier

// Delimited identifiers ("...") - postgres/src/backend/parser/scan.l:309

// Dollar quoting ($...$) - postgres/src/backend/parser/scan.l:301

// Comments and operators - postgres/src/backend/parser/scan.l:342-500+

// Otherwise, treat as operator

// Identifiers - MUST come after specific character cases
// postgres/src/backend/parser/scan.l:346-349

// Single-character tokens and operators

// scanIdentifier scans an identifier or keyword following PostgreSQL rules
// Equivalent to postgres/src/backend/parser/scan.l:346-349 (identifier rule)
func (l *Lexer) scanIdentifier(startPos, startScanPos int) (*Token, error) {
	_ = "STUB: not implemented"
	// PostgreSQL identifier rules: postgres/src/backend/parser/scan.l:346-349
	// ident_start: [A-Za-z\200-\377_]
	// ident_cont: [A-Za-z\200-\377_0-9\$]
	// identifier: {ident_start}{ident_cont}*
	return nil, nil
}

// First character already validated by caller

// Continue with identifier continuation characters

// Check if this is a keyword using the keyword lookup

// Handle lookahead tokens (WITH_LA, FORMAT_LA, etc.)

// Return keyword token with the proper keyword token type and keyword value

// Regular identifier - PostgreSQL lowercases unquoted identifiers

// checkLookaheadToken checks if a keyword should be converted to its lookahead variant
// Based on PostgreSQL's base_yylex function in parser.c
// This implements the same extensible nested switch pattern PostgreSQL uses
func (l *Lexer) checkLookaheadToken(keyword *KeywordInfo, text string) TokenType {
	_ = "STUB: not implemented"
	return *new(TokenType)
}

// First level: Check if this token needs lookahead examination
// Based on PostgreSQL parser.c lines 138-161

// No lookahead needed for this token

// Get the next token for lookahead analysis

// Second level: Based on current + next token combination, determine if we need *_LA variant
// Based on PostgreSQL parser.c lines 195-321

// No lookahead transformation needed

// peekNextTokenType peeks at the next token type without consuming it
// This is used for lookahead token analysis
func (l *Lexer) peekNextTokenType() TokenType {
	_ = "STUB: not implemented"
	// Save current position to restore later
	return *new(TokenType)
}

// Restore position after lookahead

// Skip whitespace to find next token

// Try to read the next identifier/keyword

// Check if it's a keyword

// Not a keyword, return IDENT

// peekNextIdentifier peeks at the next identifier without consuming it
func (l *Lexer) peekNextIdentifier() string { _ = "STUB: not implemented"; return "" }

// Read first character

// Read rest of identifier

// skipWhitespaceForLookahead skips whitespace and comments for lookahead
func (l *Lexer) skipWhitespaceForLookahead() { _ = "STUB: not implemented"; return }

// Skip single-line comments

// Skip to end of line

// Skip multi-line comments

// consume '/'
// consume '*'

// Skip until */

// consume '*'
// consume '/'

// Not whitespace or comment

// normalizeIdentifierCase normalizes identifier case following PostgreSQL rules.
// PostgreSQL converts unquoted identifiers to lowercase using ASCII-only conversion.
// Based on postgres/src/backend/parser/scansup.c:downcase_truncate_identifier
func normalizeIdentifierCase(s string) string {
	_ = "STUB: not implemented"
	// PostgreSQL does ASCII-only case conversion for identifiers
	// This is the same logic as normalizeKeywordCase but semantically different
	return ""
}

// scanNumber scans numeric literals (integers and floating-point)
// Implements all PostgreSQL numeric literal formats
// postgres/src/backend/parser/scan.l:395-414, 1018-1077
func (l *Lexer) scanNumber(startPos, startScanPos int) (*Token, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Handle numbers starting with decimal point (.5, .123, etc.)
	// postgres/src/backend/parser/scan.l:409 - \.{decinteger}
}

// consume '.'

// Scan fractional part

// Check for exponent

// Invalid - just a lone dot, shouldn't happen

// Check for special integer formats (hex, octal, binary)
// postgres/src/backend/parser/scan.l:401-403

// Look ahead 4 bytes to handle 0x_X patterns

// Check for hexfail pattern first: 0[xX]_?
// postgres/src/backend/parser/scan.l:405

// Hexadecimal: 0[xX](_?{hexdigit})+

// Check for octfail pattern first: 0[oO]_?
// postgres/src/backend/parser/scan.l:406

// Octal: 0[oO](_?{octdigit})+

// Check for binfail pattern first: 0[bB]_?
// postgres/src/backend/parser/scan.l:407

// Binary: 0[bB](_?{bindigit})+

// Scan decimal part
// Pattern: {decdigit}(_?{decdigit})*
// postgres/src/backend/parser/scan.l:400

// PostgreSQL allows underscores in numeric literals for readability
// Peek ahead to ensure underscore is followed by digit

// consume underscore

// Not a digit or valid underscore - stop scanning

// Check for decimal point (floating point)
// postgres/src/backend/parser/scan.l:409

// Check for ".." operator (numericfail pattern)
// postgres/src/backend/parser/scan.l:410

// This is ".." operator, not a decimal point

// Consume decimal point

// Scan fractional part

// Underscores allowed in fractional part too

// Not a digit or valid underscore in fraction

// If we have digits before or after decimal point, it's a valid numeric

// Check for exponent

// Check for exponent (scientific notation)
// postgres/src/backend/parser/scan.l:412

// Check for integer_junk pattern BEFORE capturing text
// postgres/src/backend/parser/scan.l:435, 1066-1076
// Adds error to context if junk detected

// Get the text including any trailing junk

// Return the integer token (PostgreSQL continues parsing even with junk)

// scanDollarToken scans dollar-related tokens ($1, $tag$, etc.)
// Enhanced in Phase 2C to handle dollar-quoted strings - postgres/src/backend/parser/scan.l:290-320
func (l *Lexer) scanDollarToken(startPos, startScanPos int) (*Token, error) {
	_ = "STUB: not implemented"
	// Check for parameter ($1, $2, etc.)
	return nil, nil
}

// consume '$'

// Check for param_junk: $1abc pattern
// postgres/src/backend/parser/scan.l:438 and :1013-1016

// Special case: if the next character is '$' followed by a digit,
// it's another parameter, not junk

// This is another parameter ($2, $3, etc.), not junk

// Continue scanning identifier characters for real junk

// Return PARAM token despite the error

// Check for potential dollar-quoted string delimiter
// Look for pattern $tag$ where tag is optional and follows PostgreSQL rules

// Single $ character - treat as operator
// consume '$'

// isDollarQuoteStart checks if current position could start a dollar-quoted string
// Equivalent to dolqdelim pattern - postgres/src/backend/parser/scan.l:292
func (l *Lexer) isDollarQuoteStart() bool {
	_ = "STUB: not implemented"
	// Must start with $
	return false
}

// Look ahead to see if this could be a valid delimiter
// Get remaining entire query for dollar-quoted string detection

// Skip optional tag characters - dolq_start followed by dolq_cont

// Must end with $ (this completes the opening delimiter)

// Now we have the opening delimiter, extract it

// For "$$", we need to look ahead to see if there's a closing "$$" to determine
// if this should be treated as a dollar-quoted string or separate operators

// Look for closing $$ in the remaining input

// Found closing delimiter, treat as dollar-quoted string

// No closing delimiter found, treat as separate operators

// We have a valid opening delimiter (e.g., $tag$), so this should be treated
// as a dollar-quoted string regardless of whether there's a closing delimiter.
// If no closing delimiter is found, scanDollarQuotedString will handle the error.
// This matches PostgreSQL's behavior where dolqdelim pattern matches first,
// then the lexer enters xdolq state to scan for content and closing delimiter.

// isDollarQuoteTagStart checks if character can start a dollar-quote tag
func (l *Lexer) isDollarQuoteTagStart(ch rune) bool { _ = "STUB: not implemented"; return false }

// isDollarQuoteTagCont checks if character can continue a dollar-quote tag
func (l *Lexer) isDollarQuoteTagCont(ch rune) bool { _ = "STUB: not implemented"; return false }

// scanOperatorOrPunctuation scans operators and punctuation following PostgreSQL rules
// Equivalent to postgres/src/backend/parser/scan.l:380-382 (self and operator rules)
func (l *Lexer) scanOperatorOrPunctuation(startPos, startScanPos int) (*Token, error) {
	_ = "STUB: not implemented"
	return nil,

		// Special handling for dot sequences (not operator chars in PostgreSQL sense)
		nil
}

// Handle consecutive dots with multi-dot sequence tracking

// We're already in a multi-dot sequence, treat this as individual dot
// Fall through to return single dot

// Count total consecutive dots to determine if this starts a multi-dot sequence
// Current dot

// Exactly two dots - create DOT_DOT
// consume the second '.'

// Start of multi-dot sequence - mark flag and return first dot

// Fall through to return single dot

// No next dot, clear multi-dot sequence flag

// Check if this is a floating point number (.5, .123, etc.)
// postgres/src/backend/parser/scan.l:409 - numeric: (({decinteger}\.{decinteger}?)|(\.{decinteger}))

// Single dot - treat as self character

// Handle special case of colon-based operators (: is not an OpChar)

// Single colon - treat as self character

// For operator characters, check if we need to look for multi-character operators

// Count consecutive operator chars to determine if this could be multi-char

// For 2+ character operators, check for known specific operators

// Two character operators - check for known specific operators

// consume second character

// Unknown 2-char operator or 3 or more characters (like !==, ===) - always use general operator scanning

// Single character operator - prioritize self-character behavior

// Generic single-char operator

// Check if this is a "self" character (non-operator single-char token)
// postgres/src/backend/parser/scan.l:380 - self: [,()\[\].;\:\+\-\*\/\%\^\<\>\=]

// Fallback to generic operator

// countConsecutiveOpChars counts consecutive operator characters from current position
func (l *Lexer) countConsecutiveOpChars() int { _ = "STUB: not implemented"; return 0 }

// Check for comment start sequences and stop before them

// Check for /* comment start

// Check for -- comment start

// getOpText gets the operator text of specified length from startScanPos
func (l *Lexer) getOpText(startScanPos, length int) string { _ = "STUB: not implemented"; return "" }

// scanOperator scans a multi-character operator
// Equivalent to postgres/src/backend/parser/scan.l:868-963 (operator rule)
func (l *Lexer) scanOperator(startPos, startScanPos int) (*Token, error) {
	_ = "STUB: not implemented"
	// Continue scanning operator characters
	return nil, nil
}

// Check for embedded comment starts
// postgres/src/backend/parser/scan.l:869-888

// Put back the extra characters

// Apply PostgreSQL's give-back trailing +/- rule.
// postgres/src/backend/parser/scan.l:890-925

// Re-classify after trimming so single-char self tokens and known multi-char
// tokens are returned with their correct types rather than as generic Op.
// postgres/src/backend/parser/scan.l:933-962

// hasOperatorExtensionChar reports whether op contains any non-SQL operator
// character. PostgreSQL keeps a trailing '+' or '-' on a multi-char operator
// only when one of these chars is present.
// postgres/src/backend/parser/scan.l:907-911
func hasOperatorExtensionChar(op string) bool { _ = "STUB: not implemented"; return false }

// giveBackTrailingPlusMinus implements PostgreSQL's rule that '+' or '-' may
// not be the last character of a multi-char operator unless the operator
// contains a non-SQL operator char. Trailing '+'/'-' chars are stripped so
// that "=-1" lexes as "=" then "-1" rather than as the operator "=-".
// postgres/src/backend/parser/scan.l:890-925
func giveBackTrailingPlusMinus(op string) string { _ = "STUB: not implemented"; return "" }

// skipWhitespace skips whitespace and handles comments following PostgreSQL rules
// Equivalent to postgres/src/backend/parser/scan.l:222-230 (whitespace and comment rules)
func (l *Lexer) skipWhitespace() error { _ = "STUB: not implemented"; return nil }

// Handle regular whitespace - postgres/src/backend/parser/scan.l:222
// space: [ \t\n\r\f\v]

// Update line/column tracking for position reporting

// Handle line comments - postgres/src/backend/parser/scan.l:227
// comment: ("--"{non_newline}*)

// Skip the "--"

// Skip to end of line or EOF

// Consume the newline and update position tracking

// Handle \r\n as single newline

// No more whitespace to skip

// GetContext returns the lexer context (for testing and debugging)
func (l *Lexer) GetContext() *ParseContext {
	_ = "STUB: not implemented"

	// State-based scanner functions - placeholders for future phases
	// These implement the PostgreSQL exclusive state scanning
	return nil
}

// scanQuoteStopState handles scanning in the xqs state (quote stop detection)
// Placeholder for Phase 2C - postgres/src/backend/parser/scan.l:197
func (l *Lexer) scanQuoteStopState(startPos, startScanPos int) (*Token, error) {
	_ = "STUB: not implemented"
	// For now, return to initial state and scan as operator
	return nil, nil
}

// Character classification functions are now provided by charclass.go
// This provides optimized lookup table-based classification instead of function calls

// scanExponentPart handles the exponent part of floating point numbers
// postgres/src/backend/parser/scan.l:412-413
func (l *Lexer) scanExponentPart(startPos, startScanPos int, isFloat bool) (*Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Consume 'e' or 'E'

// Check for optional sign

// Must have at least one digit after E[+-]

// Not a digit or valid underscore in exponent

// Check for realfail pattern
// postgres/src/backend/parser/scan.l:413, 1062-1064

// Return as FCONST even if incomplete

// Check for real_junk pattern BEFORE capturing text
// Adds error to context if junk detected

// No exponent

// Check for numeric_junk pattern BEFORE capturing text
// Adds error to context if junk detected

// Integer - check for trailing junk BEFORE capturing text

// Error already added to context by checkTrailingJunk

// processIntegerLiteral processes an integer literal and returns appropriate token
// Handles decimal, hexadecimal (0x), octal (0o), and binary (0b) integer literals
// with underscore separators and overflow detection like PostgreSQL
// postgres/src/backend/utils/adt/numutils.c:pg_strtoint32_safe
func (l *Lexer) processIntegerLiteral(text string, startPos int) *Token {
	_ = "STUB: not implemented"
	// Try to parse as 32-bit integer first
	return nil
}

// Integer too large, treat as float (like PostgreSQL)

// Return as integer constant

// parseInteger32 parses a string as a 32-bit integer, handling different bases
// Returns the value and whether an overflow occurred
// Based on PostgreSQL's pg_strtoint32_safe implementation
func (l *Lexer) parseInteger32(s string) (int32, bool) { _ = "STUB: not implemented"; return 0, false }

// Handle sign

// Determine base and parse accordingly

// Hexadecimal

// Octal

// Binary

// Regular decimal

// Regular decimal

// Convert to signed and check bounds

// Check if negation would overflow
// -INT32_MIN

// INT32_MAX

// parseDecimalInteger parses decimal digits with optional underscores
func (l *Lexer) parseDecimalInteger(s string, start int) (uint32, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// Check for overflow before multiplying

// Underscore must be followed by more digits

// Invalid syntax

// parseHexInteger parses hexadecimal digits with optional underscores
func (l *Lexer) parseHexInteger(s string, start int) (uint32, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// Underscore must be followed by more hex digits

// Invalid syntax

// Check for overflow before shifting

// parseOctalInteger parses octal digits with optional underscores
func (l *Lexer) parseOctalInteger(s string, start int) (uint32, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// Check for overflow before shifting

// Underscore must be followed by more octal digits

// Invalid syntax

// parseBinaryInteger parses binary digits with optional underscores
func (l *Lexer) parseBinaryInteger(s string, start int) (uint32, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// Use if/else (not switch) so the trailing `break` exits the loop;
// `break` inside a switch only exits the switch and would spin here.
//nolint:staticcheck // QF1003: switch would defeat the point
// Check for overflow before shifting

// Underscore must be followed by more binary digits

// Invalid syntax

// isHexDigit checks if a character is a valid hexadecimal digit
func isHexDigit(c byte) bool { _ = "STUB: not implemented"; return false }

// checkTrailingJunk checks for invalid characters after numeric literals
// Handles PostgreSQL's integer_junk, numeric_junk, and real_junk patterns
// postgres/src/backend/parser/scan.l:435-437, 1066-1076
func (l *Lexer) checkTrailingJunk() error { _ = "STUB: not implemented"; return nil }

// PostgreSQL's junk patterns: {numeric}{identifier}
// Any identifier character (including underscore) immediately following
// a numeric literal is considered trailing junk

// Consume all the junk characters to match PostgreSQL behavior
// PostgreSQL's junk patterns consume the entire invalid token

// checkIntegerFailPattern checks for fail patterns (0x_, 0o_, 0b_) and handles them
// This consolidates the common logic for hexfail, octfail, and binfail patterns
// postgres/src/backend/parser/scan.l:405-407
func (l *Lexer) checkIntegerFailPattern(peekAhead []byte, prefixChar byte, digitChecker func(byte) bool, errorMsg string, errorType LexerErrorType, startPos, startScanPos int) (*Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This is a fail pattern - consume "0" + prefixChar and optional "_"
// consume '0'
// consume prefix char (x/o/b)

// consume '_'

// Not a fail pattern, continue with normal processing

// scanSpecialInteger consolidates hex/octal/binary integer scanning
// This combines scanHexInteger, scanOctalInteger, and scanBinaryInteger with identical logic
// postgres/src/backend/parser/scan.l:401-403
func (l *Lexer) scanSpecialInteger(startPos, startScanPos int, digitChecker func(byte) bool, errorMsg string, errorType LexerErrorType) (*Token, error) {
	_ = "STUB: not implemented"
	// Consume "0" + prefix char (already done by caller)
	return nil, nil
}

// consume '0'
// consume prefix char (x/o/b)

// Check underscore is followed by valid digit
// PostgreSQL pattern: 0[xX](_?{hexdigit})+ allows underscore before any digit

// Not a valid digit or valid underscore

// Must have at least one digit (PostgreSQL pattern requires at least one group)

// Reject identifier-continuation chars immediately after the radix body.
// Upstream emits this via the {hexinteger}{identifier} junk pattern at
// scan.l:1066, so e.g. `0x0o` errors instead of tokenizing as `0x0` AS `o`.

//nolint:nilerr // Error is collected via context, not returned

// String returns a string representation of the lexer for debugging
func (l *Lexer) String() string { _ = "STUB: not implemented"; return "" }

// SetParseTree sets the parse tree result
func (l *Lexer) SetParseTree(tree []ast.Stmt) {
	_ = "STUB: not implemented"

	// GetParseTree returns the parse tree result
	return
}

func (l *Lexer) GetParseTree() []ast.Stmt {
	_ = "STUB: not implemented"

	// GetPosition returns the current position in the input
	return nil
}

func (l *Lexer) GetPosition() int { _ = "STUB: not implemented"; return 0 }

// RecordError records a parsing error
func (l *Lexer) RecordError(err error) { _ = "STUB: not implemented"; return }

// HasErrors returns true if there are any errors
func (l *Lexer) HasErrors() bool { _ = "STUB: not implemented"; return false }

// GetErrors returns all errors encountered
func (l *Lexer) GetErrors() []error { _ = "STUB: not implemented"; return nil }
