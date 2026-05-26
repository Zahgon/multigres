// PostgreSQL Database Management System
// (also known as Postgres, formerly known as Postgres95)
//
//	Portions Copyright (c) 2025, Supabase, Inc
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
// Package ast provides PostgreSQL AST parse infrastructure node definitions.
// This file contains core parsing infrastructure nodes essential for lexer/parser integration.
// Ported from postgres/src/include/nodes/parsenodes.h
package ast

// ==============================================================================
// CORE PARSE INFRASTRUCTURE - Stage 1A Implementation
// Essential parsing foundation nodes for lexer/parser integration
// Ported from postgres/src/include/nodes/parsenodes.h
// ==============================================================================

// RawStmt represents a raw statement wrapper that contains the original query text
// and location information before semantic analysis.
// Ported from postgres/src/include/nodes/parsenodes.h:2017-2024
type RawStmt struct {
	BaseNode
	Stmt         Stmt // The parsed statement tree
	StmtLocation int  // Start location of stmt in original query string
	StmtLen      int  // Length of stmt in original query string
}

// NewRawStmt creates a new RawStmt node.
func NewRawStmt(stmt Stmt, location int, length int) *RawStmt {
	_ = "STUB: not implemented"
	return nil
}

func (r *RawStmt) String() string { _ = "STUB: not implemented"; return "" }

func (r *RawStmt) StatementType() string {
	_ = "STUB: not implemented"

	// A_Expr represents a generic expression node used during parsing before semantic analysis.
	// This is the primary expression node used in the parse tree.
	// Ported from postgres/src/include/nodes/parsenodes.h:329-339
	return ""
}

type A_Expr struct {
	BaseNode
	Kind  A_Expr_Kind // Expression type (operator, comparison, etc.)
	Name  *NodeList   // Possibly-qualified operator name
	Lexpr Node        // Left operand
	Rexpr Node        // Right operand (or NULL for unary operators)
	// Note: Location is handled by BaseNode.Loc
}

// A_Expr_Kind represents the type of A_Expr.
// Ported from postgres/src/include/nodes/parsenodes.h:331-339
type A_Expr_Kind int

const (
	AEXPR_OP              A_Expr_Kind = iota // Normal operator
	AEXPR_OP_ANY                             // Scalar op ANY (array)
	AEXPR_OP_ALL                             // Scalar op ALL (array)
	AEXPR_DISTINCT                           // IS DISTINCT FROM
	AEXPR_NOT_DISTINCT                       // IS NOT DISTINCT FROM
	AEXPR_NULLIF                             // NULLIF(a, b)
	AEXPR_IN                                 // IN (list)
	AEXPR_LIKE                               // LIKE
	AEXPR_ILIKE                              // ILIKE
	AEXPR_SIMILAR                            // SIMILAR TO
	AEXPR_BETWEEN                            // BETWEEN
	AEXPR_NOT_BETWEEN                        // NOT BETWEEN
	AEXPR_BETWEEN_SYM                        // BETWEEN SYMMETRIC
	AEXPR_NOT_BETWEEN_SYM                    // NOT BETWEEN SYMMETRIC
)

// NewA_Expr creates a new A_Expr node.
func NewA_Expr(kind A_Expr_Kind, name *NodeList, lexpr, rexpr Node, location int) *A_Expr {
	_ = "STUB: not implemented"
	return nil
}

func (a *A_Expr) String() string { _ = "STUB: not implemented"; return "" }

// opName returns the operator string (e.g. "~~", "!~~", "!~") stored in Name.
// Empty when Name is unset.
func (a *A_Expr) opName() string { _ = "STUB: not implemented"; return "" }

// unwrapLikeEscape returns the (pattern, escape) pair from a
// `pg_catalog.<wrapper>(pattern[, escape])` FuncCall, where wrapper is one of
// "like_escape" or "similar_to_escape". The parser folds LIKE/ILIKE/SIMILAR
// TO patterns through these wrappers when an ESCAPE clause is present (and
// always for SIMILAR TO so the engine can compile the SQL-spec pattern to a
// POSIX regex). For deparse we have to peel the wrapper back off; otherwise
// `x LIKE 'p' ESCAPE '#'` round-trips to `x LIKE pg_catalog.like_escape(...)`
// which sends a FuncCall as the pattern instead of the SQL-spec form, and the
// backend sees a different semantics. Returns ("", "") when expr is not the
// expected wrapper so callers can fall back to the raw deparse.
func unwrapLikeEscape(expr Node, wrapper string) (pattern string, escape string) {
	_ = "STUB: not implemented"
	return "", ""
}

// SqlString returns the SQL representation of the A_Expr
func (a *A_Expr) SqlString() string { _ = "STUB: not implemented"; return "" }

// Check if this is a qualified operator (OPERATOR(schema.op) syntax)
// Qualified operators have multiple items in the Name list

// This is a qualified operator - format as OPERATOR(schema.op)

// Join the parts with dots (e.g., "pg_catalog" "+" becomes "pg_catalog.+")

// Format the expression with OPERATOR syntax

// Unary qualified operator (rare but possible)

// Simple operator (not qualified)

// For operators, we need the raw string value, not the SQL quoted version

// Fallback for other node types - use their string representation

// Unary operators (NOT, unary +, unary -)

// Unary + or -

// Binary operators

// Determine if this is IN or NOT IN based on the operator

// Check if Rexpr is a SubLink (for subqueries)

// SubLink will handle its own deparsing

// Otherwise, it's a list of values

// Handle scalar op ANY (array) expressions

// Get the operator
// Default operator

// Handle scalar op ALL (array) expressions

// Get the operator
// Default operator

// Handle IS DISTINCT FROM expressions

// Handle IS NOT DISTINCT FROM expressions

// Handle NULLIF(a, b) expressions

// Rexpr should be a NodeList with two elements

// Rexpr should be a NodeList with two elements

// Rexpr should be a NodeList with two elements

// Rexpr should be a NodeList with two elements

func (a *A_Expr) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (a *A_Expr) IsExpr() bool {
	_ = "STUB: not implemented"

	// A_Const represents a constant value in the parse tree.
	// Ported from postgres/src/include/nodes/parsenodes.h:357-365
	return false
}

type A_Const struct {
	BaseNode
	Val    Value // The constant value (Integer, Float, String, BitString, Boolean, or Null)
	Isnull bool  // SQL NULL constant
}

// NewA_Const creates a new A_Const node.
func NewA_Const(val Value, location int) *A_Const { _ = "STUB: not implemented"; return nil }

// NewA_ConstNull creates a new A_Const node representing a NULL value.
func NewA_ConstNull(location int) *A_Const { _ = "STUB: not implemented"; return nil }

func (a *A_Const) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of the A_Const
func (a *A_Const) SqlString() string { _ = "STUB: not implemented"; return "" }

// Use the Value's SqlString() method

func (a *A_Const) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (a *A_Const) IsExpr() bool {
	_ = "STUB: not implemented"

	// ParamRef represents a parameter reference ($1, $2, etc.) in the parse tree.
	// Ported from postgres/src/include/nodes/parsenodes.h:301-309
	return false
}

type ParamRef struct {
	BaseNode
	Number int // Parameter number (1-based)
}

// NewParamRef creates a new ParamRef node.
func NewParamRef(number int, location int) *ParamRef { _ = "STUB: not implemented"; return nil }

func (p *ParamRef) String() string { _ = "STUB: not implemented"; return "" }

func (p *ParamRef) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (p *ParamRef) IsExpr() bool {
	_ = "STUB: not implemented"

	// SqlString returns the SQL representation of the ParamRef
	return false
}

func (p *ParamRef) SqlString() string { _ = "STUB: not implemented"; return "" }

// TypeCast represents a type cast expression (CAST(expr AS type) or expr::type).
// Ported from postgres/src/include/nodes/parsenodes.h:370-380
type TypeCast struct {
	BaseNode
	Arg      Node      // The expression being cast
	TypeName *TypeName // The target type
}

// ParenExpr represents a parenthesized expression to preserve grouping
type ParenExpr struct {
	BaseNode
	Expr Node // The expression inside parentheses
}

// NewTypeCast creates a new TypeCast node.
func NewTypeCast(arg Node, typeName *TypeName, location int) *TypeCast {
	_ = "STUB: not implemented"
	return nil
}

// NewParenExpr creates a new ParenExpr node.
func NewParenExpr(expr Node, location int) *ParenExpr { _ = "STUB: not implemented"; return nil }

func (t *TypeCast) String() string { _ = "STUB: not implemented"; return "" }

func (p *ParenExpr) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of the TypeCast (CAST syntax)
func (t *TypeCast) SqlString() string { _ = "STUB: not implemented"; return "" }

// Special handling for INTERVAL literals - convert back to INTERVAL 'value' UNIT syntax
// if t.TypeName != nil && t.TypeName.Names != nil && t.TypeName.Names.Len() > 0 {
// 	if firstItem, ok := t.TypeName.Names.Items[0].(*String); ok && firstItem.SVal == "interval" {
// 		if t.TypeName.Typmods != nil && t.TypeName.Typmods.Len() > 0 {
// 			if firstMod, ok := t.TypeName.Typmods.Items[0].(*Integer); ok {
// 				intervalUnit := intervalMaskToString(firstMod.IVal)
// 				if intervalUnit == "FULL_RANGE" {
// 					if t.TypeName.Typmods.Len() == 2 {
// 						// INTERVAL(precision) 'value' format for full range with precision
// 						if precision, ok := t.TypeName.Typmods.Items[1].(*Integer); ok {
// 							return fmt.Sprintf("INTERVAL(%d) %s", precision.IVal, argStr)
// 						}
// 					}
// 					// INTERVAL 'value' format for full range without precision
// 					return fmt.Sprintf("INTERVAL %s", argStr)
// 				} else if intervalUnit != "" {
// 					// INTERVAL 'value' UNIT format for specific units
// 					return fmt.Sprintf("INTERVAL %s %s", argStr, intervalUnit)
// 				}
// 			}
// 		}
// 	}
// }

// Always use explicit CAST(expr AS type) syntax to avoid precedence issues

// SqlString returns the SQL representation of the ParenExpr (preserves parentheses)
func (p *ParenExpr) SqlString() string { _ = "STUB: not implemented"; return "" }

func (t *TypeCast) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (t *TypeCast) IsExpr() bool { _ = "STUB: not implemented"; return false }

func (p *ParenExpr) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (p *ParenExpr) IsExpr() bool {
	_ = "STUB: not implemented"

	// FuncCall represents a function call in the parse tree.
	// Ported from postgres/src/include/nodes/parsenodes.h:423-444
	return false
}

type FuncCall struct {
	BaseNode
	Funcname       *NodeList    // Qualified function name
	Args           *NodeList    // List of arguments
	AggOrder       *NodeList    // ORDER BY list for aggregates
	AggFilter      Node         // FILTER clause for aggregates
	Over           *WindowDef   // OVER clause for window functions
	AggWithinGroup bool         // ORDER BY appeared in WITHIN GROUP
	AggStar        bool         // Function was written as foo(*)
	AggDistinct    bool         // DISTINCT was specified
	FuncVariadic   bool         // VARIADIC was specified
	Funcformat     CoercionForm // How to display this node
}

// NewFuncCall creates a new FuncCall node.
func NewFuncCall(funcname *NodeList, args *NodeList, location int) *FuncCall {
	_ = "STUB: not implemented"
	return nil
}

// Default to explicit call syntax

func (f *FuncCall) String() string { _ = "STUB: not implemented"; return "" }

// isInternalPgCatalogFunction checks if a function name is one where the parser
// adds pg_catalog qualification during syntax transformation. Returns the
// normalized function name and true if it's internal, or empty string and false otherwise.
func isInternalPgCatalogFunction(name string) bool { _ = "STUB: not implemented"; return false }

// AT TIME ZONE

// LIKE ... ESCAPE

// SIMILAR TO

// OVERLAPS

// IS NORMALIZED

// COLLATION FOR

// SYSTEM_USER

// EXTRACT

// NORMALIZE

// OVERLAY

// POSITION

// SUBSTRING

// TRIM(BOTH ...)

// TRIM(LEADING ...)

// TRIM(TRAILING ...)

// XMLEXISTS

// SqlString returns the SQL representation of the FuncCall
func (f *FuncCall) SqlString() string {
	_ = "STUB: not implemented"
	// Build function name (could be qualified like schema.func)
	return ""
}

// Check for pg_catalog.func where the parser added pg_catalog

// Skip pg_catalog for internal functions

// Normalize common function names to uppercase

// Window functions

// Aggregate functions commonly used as window functions

// Special SQL functions that should use their original syntax

// Build argument list

// Prepend DISTINCT qualifier if needed

// Prepend VARIADIC qualifier to last argument if needed

// Add ORDER BY clause inside function parentheses if present (for aggregates that aren't WITHIN GROUP)

// Handle special function syntax

// EXTRACT function uses special syntax: EXTRACT(field FROM source)
// The first argument should be the field name without quotes, the second is the source

// Remove quotes from field name if it's a string literal

// Use lowercase for function name to match PostgreSQL style

// SUBSTRING function with SQL standard syntax: SUBSTRING(string FROM start [FOR length])

// SUBSTRING(string FROM start FOR length)

// SUBSTRING(string FROM start)

// Fallback to regular function call

// POSITION function with SQL standard syntax: POSITION(substring IN string)
// Note: Parser reorders arguments to [string, substring], so we need to swap them back

// POSITION(substring IN string)

// Fallback to regular function call

// OVERLAY function with SQL standard syntax: OVERLAY(string PLACING substring FROM start [FOR length])

// OVERLAY(string PLACING substring FROM start FOR length)

// OVERLAY(string PLACING substring FROM start)

// Fallback to regular function call

// NORMALIZE function: normalize(string [, form])
// The second argument is an A_Const containing the normalization form as a keyword
// Default to the string representation

// The grammar stores the normalization form as an unquoted string

// normalize(string, form, ...)

// normalize(string, form)

// Default to the string representation

// The grammar stores the normalization form from unicode_normal_form rule
// which returns "NFC", "NFD", "NFKC", or "NFKD" (uppercase)

// SYSTEM_USER function call with no arguments should be deparsed as SYSTEM_USER (SQL value function)

// xmlexists function with SQL syntax: xmlexists(xpath PASSING [BY REF] document [BY REF])
// The grammar converts xmlexists(A PASSING [BY REF] B [BY REF]) to xmlexists(A, B, ...)
// We restore the xmlexists syntax using BY REF as separator between arguments

// Fallback for edge cases

// Add WITHIN GROUP clause for ordered-set aggregates

// Add FILTER clause for filtered aggregates

// Add OVER clause for window functions

// Check if this is ONLY a window reference (no additional clauses)

// Pure window reference - no parentheses

// Window specification or reference with additional clauses - with parentheses

func (f *FuncCall) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (f *FuncCall) IsExpr() bool {
	_ = "STUB: not implemented"

	// A_Star represents an asterisk (*) in the parse tree, typically used in SELECT *.
	// Ported from postgres/src/include/nodes/parsenodes.h:445-455
	return false
}

type A_Star struct {
	BaseNode
}

// NewA_Star creates a new A_Star node.
func NewA_Star(location int) *A_Star { _ = "STUB: not implemented"; return nil }

func (a *A_Star) String() string { _ = "STUB: not implemented"; return "" }

func (a *A_Star) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (a *A_Star) IsExpr() bool {
	_ = "STUB: not implemented"

	// SqlString returns the SQL representation of A_Star
	return false
}

func (a *A_Star) SqlString() string {
	_ = "STUB: not implemented"

	// A_Indices represents array indices in the parse tree (e.g., array[1:3]).
	// Ported from postgres/src/include/nodes/parsenodes.h:456-462
	return ""
}

type A_Indices struct {
	BaseNode
	IsSlice bool // True for slicing (e.g., array[1:3])
	Lidx    Node // Lower index (NULL if not specified)
	Uidx    Node // Upper index (NULL if not specified)
}

// NewA_Indices creates a new A_Indices node for single index access.
func NewA_Indices(idx Node, location int) *A_Indices { _ = "STUB: not implemented"; return nil }

// NewA_IndicesSlice creates a new A_Indices node for slice access.
func NewA_IndicesSlice(lidx, uidx Node, location int) *A_Indices {
	_ = "STUB: not implemented"
	return nil
}

func (a *A_Indices) String() string { _ = "STUB: not implemented"; return "" }

func (a *A_Indices) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (a *A_Indices) IsExpr() bool {
	_ = "STUB: not implemented"

	// SqlString returns the SQL representation of A_Indices (handles both single index and slice)
	return false
}

func (a *A_Indices) SqlString() string {
	_ = "STUB: not implemented"

	// Slice syntax [lower:upper]
	return ""
}

// Single index syntax [index]

// A_Indirection represents indirection (field access) in the parse tree (e.g., obj.field).
// Ported from postgres/src/include/nodes/parsenodes.h:479-488
type A_Indirection struct {
	BaseNode
	Arg         Node      // The base expression
	Indirection *NodeList // List of A_Indices and/or String nodes
}

// NewA_Indirection creates a new A_Indirection node.
func NewA_Indirection(arg Node, indirection *NodeList, location int) *A_Indirection {
	_ = "STUB: not implemented"
	return nil
}

func (a *A_Indirection) String() string { _ = "STUB: not implemented"; return "" }

func (a *A_Indirection) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (a *A_Indirection) IsExpr() bool {
	_ = "STUB: not implemented"

	// SqlString returns the SQL representation of A_Indirection
	return false
}

func (a *A_Indirection) SqlString() string { _ = "STUB: not implemented"; return "" }

// Write the base expression

// Handle indirections (field access or array subscripts)

// Field access: obj.field

// Array subscript: obj[index] or obj[lower:upper]

// Star expansion: obj.*

// Fallback for unknown indirection types

// A_ArrayExpr represents an array expression in the parse tree (e.g., ARRAY[1,2,3]).
// Ported from postgres/src/include/nodes/parsenodes.h:489-501
type A_ArrayExpr struct {
	BaseNode
	Elements *NodeList // List of expressions
}

// NewA_ArrayExpr creates a new A_ArrayExpr node.
func NewA_ArrayExpr(elements *NodeList, location int) *A_ArrayExpr {
	_ = "STUB: not implemented"
	return nil
}

func (a *A_ArrayExpr) String() string { _ = "STUB: not implemented"; return "" }

func (a *A_ArrayExpr) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (a *A_ArrayExpr) IsExpr() bool {
	_ = "STUB: not implemented"

	// Note: CollateClause and TypeName already exist in ddl_statements.go
	// Using existing implementations to avoid conflicts
	return false
}

// ColumnDef represents a complete column definition in CREATE TABLE.
// Ported from postgres/src/include/nodes/parsenodes.h:723-750
type ColumnDef struct {
	BaseNode
	Colname       string         // Name of the column
	TypeName      *TypeName      // Type of the column
	Compression   string         // Compression method, or NULL
	Inhcount      int            // Number of times column is inherited
	IsLocal       bool           // Column is defined locally
	IsNotNull     bool           // NOT NULL constraint specified
	IsFromType    bool           // Column definition came from table type
	StorageType   char           // Storage type (TOAST)
	StorageName   string         // Storage setting name or NULL for default
	RawDefault    Node           // Default value (untransformed parse tree)
	CookedDefault Node           // Default value (transformed)
	Identity      char           // IDENTITY property
	IdentitySeq   *RangeVar      // To store identity sequence name for ALTER TABLE
	Generated     char           // GENERATED property
	Collclause    *CollateClause // Collation, if any
	CollOid       Oid            // Collation OID (InvalidOid if not set)
	Constraints   *NodeList      // Column constraints
	Fdwoptions    *NodeList      // Foreign-data-wrapper specific options
}

// SqlString generates SQL representation of a column definition
func (c *ColumnDef) SqlString() string { _ = "STUB: not implemented"; return "" }

// Add type name

// Add compression clause if specified

// Add storage clause if specified

// Add foreign-data-wrapper OPTIONS (foreign table columns), in PostgreSQL's
// generic-options form: OPTIONS (key 'value', ...).

// Add NOT NULL constraint if specified

// Add DEFAULT clause if specified

// Add collation if specified

// Add constraints if any

// Handle identity constraints specially for column definitions

// Build the identity specification with proper formatting for column definitions

// Add sequence options in parentheses (without SET keywords)

// Use regular SqlString for non-identity constraints

// A bare COLLATE (e.g. a partition/typed-table column
// `a COLLATE "POSIX"`) is stored as a CollateClause in the
// constraint list rather than in Collclause.

// NewColumnDef creates a new ColumnDef node.
func NewColumnDef(colname string, typeName *TypeName, location int) *ColumnDef {
	_ = "STUB: not implemented"
	return nil
}

func (c *ColumnDef) String() string { _ = "STUB: not implemented"; return "" }

func (c *ColumnDef) StatementType() string { _ = "STUB: not implemented"; return "" }

// WithClause represents a complete WITH clause (Common Table Expression clause).
// Ported from postgres/src/include/nodes/parsenodes.h:1592-1605
type WithClause struct {
	BaseNode
	Ctes      *NodeList // List of CommonTableExpr nodes
	Recursive bool      // TRUE for WITH RECURSIVE
}

// NewWithClause creates a new WithClause node.
func NewWithClause(ctes *NodeList, recursive bool, location int) *WithClause {
	_ = "STUB: not implemented"
	return nil
}

func (w *WithClause) String() string { _ = "STUB: not implemented"; return "" }

func (w *WithClause) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (w *WithClause) IsExpr() bool {
	_ = "STUB: not implemented"

	// SqlString returns the SQL representation of the WithClause
	return false
}

func (w *WithClause) SqlString() string { _ = "STUB: not implemented"; return "" }

// MultiAssignRef represents a multi-assignment reference (used in UPDATE (col1, col2) = (val1, val2)).
// Ported from postgres/src/include/nodes/parsenodes.h:532-542
type MultiAssignRef struct {
	BaseNode
	Source   Node // The sub-expression
	Colno    int  // Column number (1-based)
	Ncolumns int  // Number of columns in the multi-assignment
}

// NewMultiAssignRef creates a new MultiAssignRef node.
func NewMultiAssignRef(source Node, colno, ncolumns int, location int) *MultiAssignRef {
	_ = "STUB: not implemented"
	return nil
}

func (m *MultiAssignRef) String() string { _ = "STUB: not implemented"; return "" }

func (m *MultiAssignRef) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (m *MultiAssignRef) IsExpr() bool {
	_ = "STUB: not implemented"

	// SqlString returns the SQL representation of MultiAssignRef
	return false
}

func (m *MultiAssignRef) SqlString() string {
	_ = "STUB: not implemented"
	// MultiAssignRef represents a reference to a specific column in a multi-column assignment
	// In SQL, this appears as the source expression
	return ""
}

// ==============================================================================
// SUPPORT INFRASTRUCTURE NODES
// ==============================================================================

// WindowDef represents a window definition in a WINDOW clause.
// Ported from postgres/src/include/nodes/parsenodes.h:561-583
type WindowDef struct {
	BaseNode
	Name            string    // Window name (NULL for inline windows)
	Refname         string    // Referenced window name, if any
	PartitionClause *NodeList // PARTITION BY expression list
	OrderClause     *NodeList // ORDER BY (list of SortBy)
	FrameOptions    int       // Frame_option flags
	StartOffset     Node      // Expression for start offset
	EndOffset       Node      // Expression for end offset
}

// Frame option constants for WindowDef FrameOptions field.
// Ported from postgres/src/include/nodes/parsenodes.h:581-605
const (
	FRAMEOPTION_NONDEFAULT                = 0x00001 // any specified?
	FRAMEOPTION_RANGE                     = 0x00002 // RANGE behavior
	FRAMEOPTION_ROWS                      = 0x00004 // ROWS behavior
	FRAMEOPTION_GROUPS                    = 0x00008 // GROUPS behavior
	FRAMEOPTION_BETWEEN                   = 0x00010 // BETWEEN given?
	FRAMEOPTION_START_UNBOUNDED_PRECEDING = 0x00020 // start is UNBOUNDED PRECEDING
	FRAMEOPTION_END_UNBOUNDED_PRECEDING   = 0x00040 // (disallowed)
	FRAMEOPTION_START_UNBOUNDED_FOLLOWING = 0x00080 // (disallowed)
	FRAMEOPTION_END_UNBOUNDED_FOLLOWING   = 0x00100 // end is UNBOUNDED FOLLOWING
	FRAMEOPTION_START_CURRENT_ROW         = 0x00200 // start is CURRENT ROW
	FRAMEOPTION_END_CURRENT_ROW           = 0x00400 // end is CURRENT ROW
	FRAMEOPTION_START_OFFSET_PRECEDING    = 0x00800 // start is OFFSET PRECEDING
	FRAMEOPTION_END_OFFSET_PRECEDING      = 0x01000 // end is OFFSET PRECEDING
	FRAMEOPTION_START_OFFSET_FOLLOWING    = 0x02000 // start is OFFSET FOLLOWING
	FRAMEOPTION_END_OFFSET_FOLLOWING      = 0x04000 // end is OFFSET FOLLOWING
	FRAMEOPTION_EXCLUDE_CURRENT_ROW       = 0x08000 // omit current row
	FRAMEOPTION_EXCLUDE_GROUP             = 0x10000 // omit current row & peers
	FRAMEOPTION_EXCLUDE_TIES              = 0x20000 // omit current row's peers

	// Compound options - postgres/src/include/nodes/parsenodes.h:600
	FRAMEOPTION_START_OFFSET = FRAMEOPTION_START_OFFSET_PRECEDING | FRAMEOPTION_START_OFFSET_FOLLOWING
	FRAMEOPTION_END_OFFSET   = FRAMEOPTION_END_OFFSET_PRECEDING | FRAMEOPTION_END_OFFSET_FOLLOWING
	FRAMEOPTION_DEFAULTS     = FRAMEOPTION_RANGE | FRAMEOPTION_START_UNBOUNDED_PRECEDING | FRAMEOPTION_END_CURRENT_ROW
)

// NewWindowDef creates a new WindowDef node.
func NewWindowDef(name string, location int) *WindowDef { _ = "STUB: not implemented"; return nil }

func (w *WindowDef) String() string { _ = "STUB: not implemented"; return "" }

func (w *WindowDef) StatementType() string { _ = "STUB: not implemented"; return "" }

func (w *WindowDef) SqlString() string { _ = "STUB: not implemented"; return "" }

// renderFrameOptions converts frame options to SQL string
func (w *WindowDef) renderFrameOptions() string {
	_ = "STUB: not implemented"

	// Frame mode (ROWS, RANGE, or GROUPS)
	return ""
}

// Handle BETWEEN clause

// Start boundary

// End boundary

// Single boundary (no BETWEEN)

// Handle exclusion clause

// renderFrameBoundary renders a single frame boundary (start or end)
func (w *WindowDef) renderFrameBoundary(isStart bool) string { _ = "STUB: not implemented"; return "" }

// Start boundary

// End boundary

func (w *WindowDef) SqlStringForContext(inWindowClause bool) string {
	_ = "STUB: not implemented"

	// Add window reference if present
	return ""
}

// Add PARTITION BY clause

// Add ORDER BY clause

// Add frame specification

// Note: SortBy, SortByDir, and SortByNulls already exist in ddl_statements.go
// However, the existing SortBy is incomplete - let me implement a more complete version

// SortBy represents a sort specification in ORDER BY clauses.
// Ported from postgres/src/include/nodes/parsenodes.h:543-560
type SortBy struct {
	BaseNode
	Node        Node        // Expression to sort on
	SortbyDir   SortByDir   // ASC/DESC/USING/DEFAULT
	SortbyNulls SortByNulls // NULLS FIRST/LAST
	UseOp       *NodeList   // Name of operator to use for comparison
}

// NewSortBy creates a new complete SortBy node.
func NewSortBy(node Node, dir SortByDir, nulls SortByNulls, location int) *SortBy {
	_ = "STUB: not implemented"
	return nil
}

func (s *SortBy) String() string { _ = "STUB: not implemented"; return "" }

func (s *SortBy) StatementType() string {
	_ = "STUB: not implemented"

	// SqlString returns the SQL representation of the SortBy
	return ""
}

func (s *SortBy) SqlString() string { _ = "STUB: not implemented"; return "" }

// Add sort direction

// Use raw string value (without quotes) for operators

// Multiple parts: use OPERATOR(schema.op) syntax

// Single part: use direct operator syntax

// Add null ordering

// GroupingSet represents a grouping set in GROUP BY clauses.
// Ported from postgres/src/include/nodes/parsenodes.h:1506-1517
type GroupingSet struct {
	BaseNode
	Kind    GroupingSetKind // Type of grouping set
	Content *NodeList       // List of expressions
}

// GroupingSetKind represents the type of grouping set.
// Ported from postgres/src/include/nodes/parsenodes.h:1490-1505
type GroupingSetKind int

const (
	GROUPING_SET_EMPTY GroupingSetKind = iota
	GROUPING_SET_SIMPLE
	GROUPING_SET_ROLLUP
	GROUPING_SET_CUBE
	GROUPING_SET_SETS
)

// NewGroupingSet creates a new GroupingSet node.
func NewGroupingSet(kind GroupingSetKind, content *NodeList, location int) *GroupingSet {
	_ = "STUB: not implemented"
	return nil
}

func (g *GroupingSet) String() string { _ = "STUB: not implemented"; return "" }

func (g *GroupingSet) StatementType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of the GroupingSet
func (g *GroupingSet) SqlString() string { _ = "STUB: not implemented"; return "" }

// Simple grouping set - just return the content expressions

// ROLLUP(expr1, expr2, ...)

// CUBE(expr1, expr2, ...)

// GROUPING SETS((expr1), (expr2), ...)

// Each item should be a GroupingSet or expression

// Handle different GroupingSet kinds

// Empty grouping set: ()

// Simple grouping set: (expr1, expr2)

// Other grouping sets (ROLLUP, CUBE, etc)

// LockingClause represents a complete locking clause (FOR UPDATE, FOR SHARE, etc.).
// Ported from postgres/src/include/nodes/parsenodes.h:831-841
type LockingClause struct {
	BaseNode
	LockedRels *NodeList          // For table locking, list of RangeVar nodes
	Strength   LockClauseStrength // Lock strength
	WaitPolicy LockWaitPolicy     // NOWAIT and SKIP LOCKED
}

// LockClauseStrength represents lock strength.
// Ported from postgres/src/include/nodes/parsenodes.h:61-67
type LockClauseStrength int

const (
	LCS_NONE LockClauseStrength = iota
	LCS_FORKEYSHARE
	LCS_FORSHARE
	LCS_FORNOKEYUPDATE
	LCS_FORUPDATE
)

// Note: LockWaitPolicy already exists in query_execution_nodes.go

// NewLockingClause creates a new LockingClause node.
func NewLockingClause(lockedRels *NodeList, strength LockClauseStrength, waitPolicy LockWaitPolicy, location int) *LockingClause {
	_ = "STUB: not implemented"
	return nil
}

func (l *LockingClause) String() string { _ = "STUB: not implemented"; return "" }

func (l *LockingClause) StatementType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of the LockingClause
func (l *LockingClause) SqlString() string {
	_ = "STUB: not implemented"

	// Determine locking strength
	return ""
}

// Add table names if specified

// Add wait policy

// XmlSerialize represents an XML serialization expression.
// Ported from postgres/src/include/nodes/parsenodes.h:842-859
type XmlSerialize struct {
	BaseNode
	XmlOptionType XmlOptionType // DOCUMENT or CONTENT
	Expr          Node          // Expression to serialize
	TypeName      *TypeName     // Target type
	Indent        bool          // INDENT option
}

// XmlOptionType represents XML option types.
// Ported from postgres/src/include/nodes/parsenodes.h:76-80
type XmlOptionType int

const (
	XMLOPTION_DOCUMENT XmlOptionType = iota
	XMLOPTION_CONTENT
)

// XmlStandaloneType represents XML standalone options for XMLROOT
type XmlStandaloneType int

const (
	XML_STANDALONE_YES XmlStandaloneType = iota
	XML_STANDALONE_NO
	XML_STANDALONE_NO_VALUE
	XML_STANDALONE_OMITTED
)

// NewXmlSerialize creates a new XmlSerialize node.
func NewXmlSerialize(xmlOptionType XmlOptionType, expr Node, typeName *TypeName, indent bool, location int) *XmlSerialize {
	_ = "STUB: not implemented"
	return nil
}

func (x *XmlSerialize) String() string { _ = "STUB: not implemented"; return "" }

func (x *XmlSerialize) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (x *XmlSerialize) IsExpr() bool {
	_ = "STUB: not implemented"

	// SqlString returns the SQL representation of XmlSerialize
	return false
}

func (x *XmlSerialize) SqlString() string { _ = "STUB: not implemented"; return "" }

// Add DOCUMENT or CONTENT

// Add the expression

// Add AS TYPE

// Add INDENT if specified

// PartitionElem represents a partition element in partition specifications.
// Ported from postgres/src/include/nodes/parsenodes.h:860-881
type PartitionElem struct {
	BaseNode
	Name      string    // Name of column to partition on
	Expr      Node      // Expression to partition on, or NULL
	Collation *NodeList // Collation name
	Opclass   *NodeList // Operator class name
}

// NewPartitionElem creates a new PartitionElem node.
func NewPartitionElem(name string, expr Node, location int) *PartitionElem {
	_ = "STUB: not implemented"
	return nil
}

func (p *PartitionElem) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of PartitionElem
func (p *PartitionElem) SqlString() string { _ = "STUB: not implemented"; return "" }

// Add COLLATE clause if present

// Collation names should be output as identifiers, not string literals

// Quote as identifier if needed

// Add operator class if present

// Operator class names should be output as identifiers, not string literals

// Quote as identifier if needed

func (p *PartitionElem) StatementType() string { _ = "STUB: not implemented"; return "" }

// TableSampleClause represents a TABLESAMPLE clause.
// Ported from postgres/src/include/nodes/parsenodes.h:1344-1367
type TableSampleClause struct {
	BaseNode
	Tsmhandler Oid       // OID of the tablesample handler function
	Args       *NodeList // List of tablesample arguments
	Repeatable Expr      // REPEATABLE expression, or NULL
}

// NewTableSampleClause creates a new TableSampleClause node.
func NewTableSampleClause(tsmhandler Oid, args *NodeList, repeatable Expr, location int) *TableSampleClause {
	_ = "STUB: not implemented"
	return nil
}

func (t *TableSampleClause) String() string { _ = "STUB: not implemented"; return "" }

func (t *TableSampleClause) StatementType() string { _ = "STUB: not implemented"; return "" }

// ObjectWithArgs represents an object name with arguments (used for functions, operators, etc.).
// Ported from postgres/src/include/nodes/parsenodes.h:2524-2539
type ObjectWithArgs struct {
	BaseNode
	Objname         *NodeList // Qualified object name
	Objargs         *NodeList // List of argument types (TypeName nodes)
	ObjfuncArgs     *NodeList // List of function arguments for ALTER FUNCTION
	ArgsUnspecified bool      // Arguments were omitted, so name must be unique
}

// NewObjectWithArgs creates a new ObjectWithArgs node.
func NewObjectWithArgs(objname *NodeList, objargs *NodeList, argsUnspecified bool, location int) *ObjectWithArgs {
	_ = "STUB: not implemented"
	return nil
}

func (o *ObjectWithArgs) String() string { _ = "STUB: not implemented"; return "" }

func (o *ObjectWithArgs) StatementType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of ObjectWithArgs
func (o *ObjectWithArgs) SqlString() string {
	_ = "STUB: not implemented"

	// Add object name
	return ""
}

// Add arguments if specified

// Check if this is an aggregate with ObjfuncArgs from aggr_args
// If so, use the DefineStmt logic to properly format ORDER BY and VARIADIC

// ObjfuncArgs is already a *NodeList, no need for type assertion

// Fall back to regular Objargs formatting

// Regular case - use Objargs
// Special case: if ObjfuncArgs is nil and Objargs is empty, this is a star aggregate

// formatAggrArgsList formats aggregate argument list with proper ORDER BY and VARIADIC syntax
// This is similar to the logic in DefineStmt.SqlString() for aggregates
func formatAggrArgsList(argsList *NodeList) string { _ = "STUB: not implemented"; return "" }

// For aggr_args, we expect the full function parameters list, not the 2-element structure
// But we need to check if it's actually a 2-element structure from aggr_args

// This is the aggr_args [args, numDirectArgs] structure

// Fall back to regular function parameter formatting

// formatAggrArgsStructure formats the [args, numDirectArgs] structure from aggr_args
func formatAggrArgsStructure(argList *NodeList, numDirectArgs int) string {
	_ = "STUB: not implemented"
	return ""
}

// Regular aggregate or COUNT(*)

// COUNT(*) case

// Ordered-set aggregate without direct args: (ORDER BY args)

// Hypothetical-set aggregate: (direct_args ORDER BY ordered_args)

// ExtractArgTypes extracts argument types from function arguments
// This function is used to convert FunctionParameter nodes to TypeName nodes
func ExtractArgTypes(funcArgs *NodeList) *NodeList { _ = "STUB: not implemented"; return nil }

// NewEmptyObjectWithArgs creates a new ObjectWithArgs node with empty constructor
func NewEmptyObjectWithArgs() *ObjectWithArgs { _ = "STUB: not implemented"; return nil }

// SinglePartitionSpec represents a single partition specification.
// Ported from postgres/src/include/nodes/parsenodes.h:945-952
type SinglePartitionSpec struct {
	BaseNode
}

// NewSinglePartitionSpec creates a new SinglePartitionSpec node.
func NewSinglePartitionSpec(location int) *SinglePartitionSpec {
	_ = "STUB: not implemented"
	return nil
}

func (s *SinglePartitionSpec) String() string { _ = "STUB: not implemented"; return "" }

func (s *SinglePartitionSpec) StatementType() string { _ = "STUB: not implemented"; return "" }

// PartitionCmd represents a partition command in ALTER TABLE.
// Ported from postgres/src/include/nodes/parsenodes.h:953-964
type PartitionCmd struct {
	BaseNode
	Name       *RangeVar           // Name of the partition
	Bound      *PartitionBoundSpec // Partition bound specification
	Concurrent bool                // CONCURRENTLY option
}

// NewPartitionCmd creates a new PartitionCmd node.
func NewPartitionCmd(name *RangeVar, bound *PartitionBoundSpec, concurrent bool, location int) *PartitionCmd {
	_ = "STUB: not implemented"
	return nil
}

func (p *PartitionCmd) String() string { _ = "STUB: not implemented"; return "" }

func (p *PartitionCmd) SqlString() string {
	_ = "STUB: not implemented"

	// Add the partition name
	return ""
}

// Add the partition bound specification if present (for ATTACH PARTITION)

func (p *PartitionCmd) StatementType() string { _ = "STUB: not implemented"; return "" }

// ==============================================================================
// SUPPORTING CONSTANTS AND HELPER TYPES
// ==============================================================================

// char represents a single character (PostgreSQL char type).
type char byte
