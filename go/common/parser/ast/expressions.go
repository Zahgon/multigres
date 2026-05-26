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
// Package ast provides PostgreSQL AST expression node definitions.
// Ported from postgres/src/include/nodes/primnodes.h
package ast

// ==============================================================================
// EXPRESSION FRAMEWORK - PostgreSQL primnodes.h implementation
// Ported from postgres/src/include/nodes/primnodes.h
// ==============================================================================

// Supporting types for expressions

// Oid represents an object identifier - ported from postgres/src/include/postgres_ext.h
type Oid uint32

// AttrNumber represents an attribute number - ported from postgres/src/include/access/attnum.h:21
type AttrNumber int16

// Index represents an array index - ported from postgres/src/include/c.h:614
type Index uint32

// Datum represents a PostgreSQL datum - ported from postgres/src/include/postgres.h:64
type Datum uintptr

// CoercionForm represents type coercion forms - ported from postgres/src/include/nodes/primnodes.h:732-737
type CoercionForm int

const (
	COERCE_EXPLICIT_CALL CoercionForm = iota // Explicit function call syntax
	COERCE_EXPLICIT_CAST                     // Explicit cast syntax
	COERCE_IMPLICIT_CAST                     // Implicit cast
	COERCE_SQL_SYNTAX                        // SQL standard syntax
)

// ParamKind represents parameter types - ported from postgres/src/include/nodes/primnodes.h:365-371
type ParamKind int

const (
	PARAM_EXTERN    ParamKind = iota // External parameter
	PARAM_EXEC                       // Executor internal parameter
	PARAM_SUBLINK                    // Sublink output column
	PARAM_MULTIEXPR                  // Multiexpr sublink column
)

// BoolExprType represents boolean expression types - ported from postgres/src/include/nodes/primnodes.h:929-932
type BoolExprType int

const (
	AND_EXPR BoolExprType = iota // AND expression
	OR_EXPR                      // OR expression
	NOT_EXPR                     // NOT expression
)

func (b BoolExprType) String() string { _ = "STUB: not implemented"; return "" }

// ==============================================================================
// BASE EXPRESSION INTERFACE
// ==============================================================================

// Expr is the abstract base type for all expression nodes.
// Ported from postgres/src/include/nodes/primnodes.h:187-192
type Expr interface {
	Node
	ExpressionType() string
	IsExpr() bool
}

// BaseExpr provides common expression functionality.
type BaseExpr struct {
	BaseNode
}

func (e *BaseExpr) IsExpr() bool {
	_ = "STUB: not implemented"

	// ==============================================================================
	// TIER 1 EXPRESSIONS - Foundation Reference and Function Expressions
	// ==============================================================================
	return false
}

// Var represents a reference to a table column.
// Ported from postgres/src/include/nodes/primnodes.h:247
type Var struct {
	BaseExpr
	Varno     int        // Relation index in range table - postgres/src/include/nodes/primnodes.h:249
	Varattno  AttrNumber // Attribute number (0 = whole-row) - postgres/src/include/nodes/primnodes.h:250
	Vartype   Oid        // pg_type OID - postgres/src/include/nodes/primnodes.h:251
	Vartypmod int32      // Type modifier - postgres/src/include/nodes/primnodes.h:252
	Varcollid Oid        // Collation OID - postgres/src/include/nodes/primnodes.h:253
	// TODO: varnullingrels *Bitmapset not yet ported - postgres/src/include/nodes/primnodes.h:274
	Varlevelsup Index      // Subquery nesting level - postgres/src/include/nodes/primnodes.h:265
	Varnosyn    Index      // Syntactic relation index - postgres/src/include/nodes/primnodes.h:276
	Varattnosyn AttrNumber // Syntactic attribute number - postgres/src/include/nodes/primnodes.h:277
}

// NewVar creates a new Var node.
func NewVar(varno int, varattno AttrNumber, vartype Oid) *Var {
	_ = "STUB: not implemented"
	return nil
}

func (v *Var) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (v *Var) String() string { _ = "STUB: not implemented"; return "" }

// Const represents a constant value in an expression.
// Ported from postgres/src/include/nodes/primnodes.h:306
type Const struct {
	BaseExpr
	Consttype   Oid   // Datatype OID - postgres/src/include/nodes/primnodes.h:308
	Consttypmod int32 // Type modifier - postgres/src/include/nodes/primnodes.h:309
	Constcollid Oid   // Collation OID - postgres/src/include/nodes/primnodes.h:310
	Constlen    int   // Type length - postgres/src/include/nodes/primnodes.h:311
	Constvalue  Datum // The actual value - postgres/src/include/nodes/primnodes.h:312
	Constisnull bool  // Whether null - postgres/src/include/nodes/primnodes.h:313
	Constbyval  bool  // Pass by value? - postgres/src/include/nodes/primnodes.h:315
}

// NewConst creates a new Const node.
func NewConst(consttype Oid, constvalue Datum, constisnull bool) *Const {
	_ = "STUB: not implemented"
	return nil
}

func (c *Const) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (c *Const) String() string { _ = "STUB: not implemented"; return "" }

// Param represents a parameter reference in a prepared statement.
// Ported from postgres/src/include/nodes/primnodes.h:373
type Param struct {
	BaseExpr
	Paramkind   ParamKind // Parameter type - postgres/src/include/nodes/primnodes.h:390
	Paramid     int       // Parameter ID - postgres/src/include/nodes/primnodes.h:391
	Paramtype   Oid       // Datatype OID - postgres/src/include/nodes/primnodes.h:392
	Paramtypmod int32     // Type modifier - postgres/src/include/nodes/primnodes.h:393
	Paramcollid Oid       // Collation OID - postgres/src/include/nodes/primnodes.h:394
}

// NewParam creates a new Param node.
func NewParam(paramkind ParamKind, paramid int, paramtype Oid) *Param {
	_ = "STUB: not implemented"
	return nil
}

func (p *Param) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (p *Param) String() string { _ = "STUB: not implemented"; return "" }

// FuncExpr represents a function call expression.
// Ported from postgres/src/include/nodes/primnodes.h:746
type FuncExpr struct {
	BaseExpr
	Funcid         Oid          // pg_proc OID - postgres/src/include/nodes/primnodes.h:749
	Funcresulttype Oid          // Result type OID - postgres/src/include/nodes/primnodes.h:750
	Funcretset     bool         // Returns set? - postgres/src/include/nodes/primnodes.h:751
	Funcvariadic   bool         // Variadic arguments? - postgres/src/include/nodes/primnodes.h:752
	Funcformat     CoercionForm // Display format - postgres/src/include/nodes/primnodes.h:753
	Funccollid     Oid          // Result collation - postgres/src/include/nodes/primnodes.h:754
	Inputcollid    Oid          // Input collation - postgres/src/include/nodes/primnodes.h:755
	Args           *NodeList    // Function arguments - postgres/src/include/nodes/primnodes.h:756
}

// NewFuncExpr creates a new FuncExpr node.
func NewFuncExpr(funcid Oid, funcresulttype Oid, args *NodeList) *FuncExpr {
	_ = "STUB: not implemented"
	return nil
}

func (f *FuncExpr) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (f *FuncExpr) String() string { _ = "STUB: not implemented"; return "" }

// OpExpr represents a binary or unary operator expression.
// Ported from postgres/src/include/nodes/primnodes.h:813
type OpExpr struct {
	BaseExpr
	Opno         Oid       // pg_operator OID - postgres/src/include/nodes/primnodes.h:816
	Opfuncid     Oid       // Underlying function OID - postgres/src/include/nodes/primnodes.h:817
	Opresulttype Oid       // Result type - postgres/src/include/nodes/primnodes.h:818
	Opretset     bool      // Returns set? - postgres/src/include/nodes/primnodes.h:819
	Opcollid     Oid       // Result collation - postgres/src/include/nodes/primnodes.h:820
	Inputcollid  Oid       // Input collation - postgres/src/include/nodes/primnodes.h:821
	Args         *NodeList // Operator arguments (1 or 2) - postgres/src/include/nodes/primnodes.h:822
}

// NewOpExpr creates a new OpExpr node.
func NewOpExpr(opno Oid, opfuncid Oid, opresulttype Oid, args *NodeList) *OpExpr {
	_ = "STUB: not implemented"
	return nil
}

func (o *OpExpr) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (o *OpExpr) String() string { _ = "STUB: not implemented"; return "" }

// BoolExpr represents a boolean expression (AND/OR/NOT).
// Ported from postgres/src/include/nodes/primnodes.h:934
type BoolExpr struct {
	BaseExpr
	Boolop BoolExprType // AND/OR/NOT - postgres/src/include/nodes/primnodes.h:947
	Args   *NodeList    // Operand expressions - postgres/src/include/nodes/primnodes.h:948
}

// NewBoolExpr creates a new BoolExpr node.
func NewBoolExpr(boolop BoolExprType, args *NodeList) *BoolExpr {
	_ = "STUB: not implemented"
	return nil
}

func (b *BoolExpr) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (b *BoolExpr) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of BoolExpr
func (b *BoolExpr) SqlString() string { _ = "STUB: not implemented"; return "" }

// Don't add parentheses - let ParenExpr handle explicit parentheses
// and let precedence rules determine when they're needed

// Don't add parentheses - let ParenExpr handle explicit parentheses
// and let precedence rules determine when they're needed

// ==============================================================================
// CONVENIENCE CONSTRUCTORS FOR COMMON PATTERNS
// ==============================================================================

// NewAndExpr creates a new AND boolean expression.
func NewAndExpr(left, right Node) *BoolExpr { _ = "STUB: not implemented"; return nil }

// NewOrExpr creates a new OR boolean expression.
func NewOrExpr(left, right Node) *BoolExpr { _ = "STUB: not implemented"; return nil }

// NewNotExpr creates a new NOT boolean expression.
func NewNotExpr(arg Node) *BoolExpr { _ = "STUB: not implemented"; return nil }

// NewBinaryOp creates a binary operator expression.
func NewBinaryOp(opno Oid, left, right Node) *OpExpr { _ = "STUB: not implemented"; return nil }

// NewUnaryOp creates a unary operator expression.
func NewUnaryOp(opno Oid, arg Node) *OpExpr { _ = "STUB: not implemented"; return nil }

// ==============================================================================
// TIER 2 EXPRESSIONS - Common SQL Features
// ==============================================================================

// CaseExpr represents a CASE expression.
// Ported from postgres/src/include/nodes/primnodes.h:1306
type CaseExpr struct {
	BaseExpr
	Casetype   Oid       // Result type - postgres/src/include/nodes/primnodes.h:1309
	Casecollid Oid       // Result collation - postgres/src/include/nodes/primnodes.h:1310
	Arg        Node      // Implicit comparison argument - postgres/src/include/nodes/primnodes.h:1311
	Args       *NodeList // List of WHEN clauses - postgres/src/include/nodes/primnodes.h:1312
	Defresult  Node      // Default result (ELSE) - postgres/src/include/nodes/primnodes.h:1313
}

// NewCaseExpr creates a new CaseExpr node.
func NewCaseExpr(casetype Oid, arg Node, whens *NodeList, defresult Node) *CaseExpr {
	_ = "STUB: not implemented"
	return nil
}

func (c *CaseExpr) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (c *CaseExpr) String() string { _ = "STUB: not implemented"; return "" }

func (c *CaseExpr) SqlString() string { _ = "STUB: not implemented"; return "" }

// Start with CASE

// Add the case argument if present (for simple CASE expressions)

// Add WHEN clauses

// Add ELSE clause if present

// CaseWhen represents a WHEN clause in a CASE expression.
// Ported from postgres/src/include/nodes/primnodes.h:1322
type CaseWhen struct {
	BaseExpr
	Expr   Node // Condition expression - postgres/src/include/nodes/primnodes.h:1325
	Result Node // Result expression - postgres/src/include/nodes/primnodes.h:1326
}

// NewCaseWhen creates a new CaseWhen node.
func NewCaseWhen(expr Node, result Node) *CaseWhen { _ = "STUB: not implemented"; return nil }

// CaseWhen uses same tag family

func (cw *CaseWhen) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (cw *CaseWhen) String() string { _ = "STUB: not implemented"; return "" }

// CoalesceExpr represents a COALESCE expression.
// Ported from postgres/src/include/nodes/primnodes.h:1484
type CoalesceExpr struct {
	BaseExpr
	Coalescetype   Oid       // Result type - postgres/src/include/nodes/primnodes.h:1487
	Coalescecollid Oid       // Result collation - postgres/src/include/nodes/primnodes.h:1488
	Args           *NodeList // Argument expressions - postgres/src/include/nodes/primnodes.h:1489
}

// NewCoalesceExpr creates a new CoalesceExpr node.
func NewCoalesceExpr(coalescetype Oid, args *NodeList) *CoalesceExpr {
	_ = "STUB: not implemented"
	return nil
}

func (c *CoalesceExpr) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (c *CoalesceExpr) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of CoalesceExpr
func (c *CoalesceExpr) SqlString() string { _ = "STUB: not implemented"; return "" }

// ArrayExpr represents an ARRAY[] constructor expression.
// Ported from postgres/src/include/nodes/primnodes.h:1370
type ArrayExpr struct {
	BaseExpr
	ArrayTypeid   Oid       // Array type OID - postgres/src/include/nodes/primnodes.h:1373
	ArrayCollid   Oid       // Array collation - postgres/src/include/nodes/primnodes.h:1374
	ElementTypeid Oid       // Element type OID - postgres/src/include/nodes/primnodes.h:1375
	Elements      *NodeList // Array elements - postgres/src/include/nodes/primnodes.h:1376
	Multidims     bool      // Multi-dimensional? - postgres/src/include/nodes/primnodes.h:1377
}

// NewArrayExpr creates a new ArrayExpr node.
func NewArrayExpr(arrayTypeid Oid, elementTypeid Oid, elements *NodeList) *ArrayExpr {
	_ = "STUB: not implemented"
	return nil
}

func (a *ArrayExpr) ExpressionType() string {
	_ = "STUB: not implemented"

	// SqlString formats the ArrayExpr as ARRAY[...] syntax
	return ""
}

func (a *ArrayExpr) SqlString() string { _ = "STUB: not implemented"; return "" }

// Check if all elements are arrays to use compact multidimensional syntax

// Use compact [[...],[...]] syntax for arrays of arrays

// Regular array syntax

func (a *ArrayExpr) String() string { _ = "STUB: not implemented"; return "" }

// ScalarArrayOpExpr represents a scalar op ANY/ALL (array) expression.
// Ported from postgres/src/include/nodes/primnodes.h:893
type ScalarArrayOpExpr struct {
	BaseExpr
	Opno        Oid       // pg_operator OID - postgres/src/include/nodes/primnodes.h:896
	Opfuncid    Oid       // Comparison function OID - postgres/src/include/nodes/primnodes.h:897
	Hashfuncid  Oid       // Hash function OID (optimization) - postgres/src/include/nodes/primnodes.h:898
	Negfuncid   Oid       // Negation function OID - postgres/src/include/nodes/primnodes.h:899
	UseOr       bool      // True for ANY, false for ALL - postgres/src/include/nodes/primnodes.h:900
	Inputcollid Oid       // Input collation - postgres/src/include/nodes/primnodes.h:901
	Args        *NodeList // Scalar and array operands - postgres/src/include/nodes/primnodes.h:902
}

// NewScalarArrayOpExpr creates a new ScalarArrayOpExpr node.
func NewScalarArrayOpExpr(opno Oid, useOr bool, scalar Node, array Node) *ScalarArrayOpExpr {
	_ = "STUB: not implemented"
	return nil
}

func (s *ScalarArrayOpExpr) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (s *ScalarArrayOpExpr) String() string { _ = "STUB: not implemented"; return "" }

// RowExpr represents a ROW() constructor expression.
// Ported from postgres/src/include/nodes/primnodes.h:1408
type RowExpr struct {
	BaseExpr
	Args      *NodeList    // Row field expressions - postgres/src/include/nodes/primnodes.h:1411
	RowTypeid Oid          // Composite type OID - postgres/src/include/nodes/primnodes.h:1412
	RowFormat CoercionForm // Display format - postgres/src/include/nodes/primnodes.h:1413
	Colnames  *NodeList    // Field names (RECORD type only) - postgres/src/include/nodes/primnodes.h:1414
}

// NewRowExpr creates a new RowExpr node.
func NewRowExpr(args *NodeList, rowTypeid Oid) *RowExpr { _ = "STUB: not implemented"; return nil }

func (r *RowExpr) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (r *RowExpr) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of the row expression
func (r *RowExpr) SqlString() string { _ = "STUB: not implemented"; return "" }

// ROW expressions are typically written as just (expr1, expr2, ...)
// unless explicitly using ROW keyword

// ==============================================================================
// TIER 2 CONVENIENCE CONSTRUCTORS
// ==============================================================================

// NewSimpleCase creates a simple CASE expression: CASE expr WHEN val1 THEN result1 ... ELSE def END
func NewSimpleCase(expr Node, whens *NodeList, defresult Node) *CaseExpr {
	_ = "STUB: not implemented"
	return nil
}

// NewSearchedCase creates a searched CASE expression: CASE WHEN condition1 THEN result1 ... ELSE def END
func NewSearchedCase(whens *NodeList, defresult Node) *CaseExpr {
	_ = "STUB: not implemented"
	return nil
}

// NewArrayConstructor creates an ARRAY[...] constructor.
func NewArrayConstructor(elements *NodeList) *ArrayExpr { _ = "STUB: not implemented"; return nil }

// NewRowConstructor creates a ROW(...) constructor.
func NewRowConstructor(fields *NodeList) *RowExpr { _ = "STUB: not implemented"; return nil }

// NewAnyExpr creates a scalar = ANY(array) expression.
func NewAnyExpr(opno Oid, scalar Node, array Node) *ScalarArrayOpExpr {
	_ = "STUB: not implemented"
	return nil
}

// NewAllExpr creates a scalar = ALL(array) expression.
func NewAllExpr(opno Oid, scalar Node, array Node) *ScalarArrayOpExpr {
	_ = "STUB: not implemented"
	return nil
}

// NewInExpr creates a scalar IN (array) expression using = ANY.
func NewInExpr(scalar Node, array Node) *ScalarArrayOpExpr { _ = "STUB: not implemented"; return nil }

// 96 is "=" operator OID

// NewNotInExpr creates a scalar NOT IN (array) expression using <> ALL.
func NewNotInExpr(scalar Node, array Node) *ScalarArrayOpExpr {
	_ = "STUB: not implemented"
	return nil
}

// 518 is "<>" operator OID

// ==============================================================================
// EXPRESSION TYPE MAPPING
// ==============================================================================

// GetExprTag maps expression types to NodeTag constants.
// Expression NodeTag constants are defined in nodes.go
func GetExprTag(exprType string) NodeTag { _ = "STUB: not implemented"; return *new(NodeTag) }

// ==============================================================================
// EXPRESSION UTILITIES
// ==============================================================================

// IsConstant checks if an expression is a constant value.
func IsConstant(expr Node) bool { _ = "STUB: not implemented"; return false }

// IsVariable checks if an expression is a variable reference.
func IsVariable(expr Node) bool { _ = "STUB: not implemented"; return false }

// IsFunction checks if an expression is a function call.
func IsFunction(expr Node) bool { _ = "STUB: not implemented"; return false }

// GetExpressionArgs returns the arguments of an expression if it has any.
func GetExpressionArgs(expr Node) *NodeList { _ = "STUB: not implemented"; return nil }

// CountArgs returns the number of arguments in an expression.
func CountArgs(expr Node) int { _ = "STUB: not implemented"; return 0 }

// ==============================================================================
// TIER 3 EXPRESSIONS - Advanced Expressions and Aggregations
// ==============================================================================

// SubLinkType represents types of sublinks - ported from postgres/src/include/nodes/primnodes.h:1556-1565
type SubLinkType int

const (
	EXISTS_SUBLINK     SubLinkType = iota // EXISTS(subquery)
	ALL_SUBLINK                           // expr op ALL(subquery)
	ANY_SUBLINK                           // expr op ANY(subquery)
	ROWCOMPARE_SUBLINK                    // (expr list) op (subquery)
	EXPR_SUBLINK                          // expr(subquery)
	MULTIEXPR_SUBLINK                     // multiple expressions
	ARRAY_SUBLINK                         // ARRAY(subquery)
	CTE_SUBLINK                           // for SubPlans only
)

func (s SubLinkType) String() string { _ = "STUB: not implemented"; return "" }

// AggSplit represents aggregate splitting modes - ported from postgres/src/include/nodes/nodes.h:479-487
type AggSplit int

const (
	AGGSPLIT_SIMPLE         AggSplit = 0  // Basic, non-split aggregation
	AGGSPLIT_INITIAL_SERIAL AggSplit = 3  // Initial phase with serialization
	AGGSPLIT_FINAL_DESERIAL AggSplit = 12 // Final phase with deserialization
)

func (a AggSplit) String() string { _ = "STUB: not implemented"; return "" }

// Aggref represents an aggregate function call expression.
// Ported from postgres/src/include/nodes/primnodes.h:439
type Aggref struct {
	BaseExpr
	Aggfnoid      Oid       // pg_proc OID of the aggregate - postgres/src/include/nodes/primnodes.h:481
	Aggtype       Oid       // Result type OID - postgres/src/include/nodes/primnodes.h:483
	Aggcollid     Oid       // Result collation - postgres/src/include/nodes/primnodes.h:485
	Inputcollid   Oid       // Input collation - postgres/src/include/nodes/primnodes.h:487
	Aggtranstype  Oid       // Transition value type - postgres/src/include/nodes/primnodes.h:492
	Aggargtypes   []Oid     // Argument type OIDs - postgres/src/include/nodes/primnodes.h:494
	Aggdirectargs *NodeList // Direct arguments for ordered-set aggs - postgres/src/include/nodes/primnodes.h:495
	Args          *NodeList // Aggregated arguments - postgres/src/include/nodes/primnodes.h:496
	Aggorder      *NodeList // ORDER BY expressions - postgres/src/include/nodes/primnodes.h:497
	Aggdistinct   *NodeList // DISTINCT expressions - postgres/src/include/nodes/primnodes.h:498
	Aggfilter     Node      // FILTER expression - postgres/src/include/nodes/primnodes.h:499
	Aggstar       bool      // True if argument list was '*' - postgres/src/include/nodes/primnodes.h:501
	Aggvariadic   bool      // True if variadic args combined - postgres/src/include/nodes/primnodes.h:505
	Aggkind       byte      // Aggregate kind - postgres/src/include/nodes/primnodes.h:507
	Aggpresorted  bool      // Input already sorted - postgres/src/include/nodes/primnodes.h:509
	Agglevelsup   Index     // > 0 if agg belongs to outer query - postgres/src/include/nodes/primnodes.h:511
	Aggsplit      AggSplit  // Expected splitting mode - postgres/src/include/nodes/primnodes.h:513
	Aggno         int       // Unique ID within Agg node - postgres/src/include/nodes/primnodes.h:515
	Aggtransno    int       // Unique ID of transition state - postgres/src/include/nodes/primnodes.h:517
}

// NewAggref creates a new Aggref node.
func NewAggref(aggfnoid Oid, aggtype Oid, args *NodeList) *Aggref {
	_ = "STUB: not implemented"
	return nil
}

func (a *Aggref) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (a *Aggref) String() string { _ = "STUB: not implemented"; return "" }

// WindowFunc represents a window function call expression.
// Ported from postgres/src/include/nodes/primnodes.h:563
type WindowFunc struct {
	BaseExpr
	Winfnoid     Oid       // pg_proc OID of the function - postgres/src/include/nodes/primnodes.h:593
	Wintype      Oid       // Result type OID - postgres/src/include/nodes/primnodes.h:595
	Wincollid    Oid       // Result collation - postgres/src/include/nodes/primnodes.h:597
	Inputcollid  Oid       // Input collation - postgres/src/include/nodes/primnodes.h:599
	Args         *NodeList // Arguments to window function - postgres/src/include/nodes/primnodes.h:601
	Aggfilter    Node      // FILTER expression - postgres/src/include/nodes/primnodes.h:603
	RunCondition *NodeList // List of run conditions - postgres/src/include/nodes/primnodes.h:605
	Winref       Index     // Index of associated WindowClause - postgres/src/include/nodes/primnodes.h:607
	Winstar      bool      // True if argument list was '*' - postgres/src/include/nodes/primnodes.h:609
	Winagg       bool      // Is function a simple aggregate? - postgres/src/include/nodes/primnodes.h:611
}

// NewWindowFunc creates a new WindowFunc node.
func NewWindowFunc(winfnoid Oid, wintype Oid, args *NodeList, winref Index) *WindowFunc {
	_ = "STUB: not implemented"
	return nil
}

func (w *WindowFunc) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (w *WindowFunc) String() string { _ = "STUB: not implemented"; return "" }

// SubLink represents a sublink expression (subquery).
// Ported from postgres/src/include/nodes/primnodes.h:1008
type SubLink struct {
	BaseExpr
	SubLinkType SubLinkType // Type of sublink - postgres/src/include/nodes/primnodes.h:1575
	SubLinkId   int         // ID (1..n); 0 if not MULTIEXPR - postgres/src/include/nodes/primnodes.h:1576
	Testexpr    Node        // Outer-query test for ALL/ANY/ROWCOMPARE - postgres/src/include/nodes/primnodes.h:1577
	OperName    *NodeList   // Originally specified operator name - postgres/src/include/nodes/primnodes.h:1579
	Subselect   Node        // Subselect as Query* or raw parsetree - postgres/src/include/nodes/primnodes.h:1581
}

// NewSubLink creates a new SubLink node.
func NewSubLink(subLinkType SubLinkType, subselect Node) *SubLink {
	_ = "STUB: not implemented"
	return nil
}

func (s *SubLink) ExpressionType() string {
	_ = "STUB: not implemented"

	// SqlString returns the SQL representation of the SubLink
	return ""
}

func (s *SubLink) SqlString() string { _ = "STUB: not implemented"; return "" }

// Handle different sublink types

// EXISTS(subquery)

// expr op ALL(subquery)

// expr op ANY(subquery) - includes IN which is = ANY

// Special case: OperName nil means it's IN not = ANY

// Regular ANY with operator

// (expr list) op (subquery)

// Simple scalar subquery: (subquery)

// Multiple expressions - just wrap in parentheses

// ARRAY(subquery)

// For SubPlans only - shouldn't appear in normal SQL

// Fallback to simple subquery

// extractOperatorFromNodeList extracts the operator string from a NodeList
func extractOperatorFromNodeList(list *NodeList) string { _ = "STUB: not implemented"; return "" }

// The operator is typically stored as a String node in the list

// Fallback: try to get SqlString of first item

func (s *SubLink) String() string { _ = "STUB: not implemented"; return "" }

// ==============================================================================
// TIER 3 CONVENIENCE CONSTRUCTORS
// ==============================================================================

// NewCountStar creates a COUNT(*) aggregate.
func NewCountStar() *Aggref { _ = "STUB: not implemented"; return nil }

// COUNT function OID 2147, result type bigint 20

// NewCount creates a COUNT(expr) aggregate.
func NewCount(expr Node) *Aggref { _ = "STUB: not implemented"; return nil }

// COUNT function OID 2147, result type bigint 20

// NewSum creates a SUM(expr) aggregate.
func NewSum(expr Node) *Aggref { _ = "STUB: not implemented"; return nil }

// SUM function OID (varies by type)

// NewAvg creates an AVG(expr) aggregate.
func NewAvg(expr Node) *Aggref { _ = "STUB: not implemented"; return nil }

// AVG function OID (varies by type)

// NewMax creates a MAX(expr) aggregate.
func NewMax(expr Node) *Aggref { _ = "STUB: not implemented"; return nil }

// MAX function OID (varies by type)

// NewMin creates a MIN(expr) aggregate.
func NewMin(expr Node) *Aggref { _ = "STUB: not implemented"; return nil }

// MIN function OID (varies by type)

// NewExistsSublink creates an EXISTS(subquery) expression.
func NewExistsSublink(subquery Node) *SubLink { _ = "STUB: not implemented"; return nil }

// NewInSublink creates an expr IN (subquery) expression.
func NewInSublink(testexpr Node, subquery Node) *SubLink { _ = "STUB: not implemented"; return nil }

// NewNotInSublink creates an expr NOT IN (subquery) expression.
func NewNotInSublink(testexpr Node, subquery Node) *SubLink { _ = "STUB: not implemented"; return nil }

// NewExprSublink creates a scalar subquery expression.
func NewExprSublink(subquery Node) *SubLink { _ = "STUB: not implemented"; return nil }

// NewArraySublink creates an ARRAY(subquery) expression.
func NewArraySublink(subquery Node) *SubLink { _ = "STUB: not implemented"; return nil }

// NewRowNumber creates a ROW_NUMBER() window function.
func NewRowNumber() *WindowFunc { _ = "STUB: not implemented"; return nil }

// ROW_NUMBER function OID 3100, result type bigint 20

// NewRank creates a RANK() window function.
func NewRank() *WindowFunc { _ = "STUB: not implemented"; return nil }

// RANK function OID 3101, result type bigint 20

// NewDenseRank creates a DENSE_RANK() window function.
func NewDenseRank() *WindowFunc { _ = "STUB: not implemented"; return nil }

// DENSE_RANK function OID 3102, result type bigint 20

// NewLag creates a LAG(expr) window function.
func NewLag(expr Node) *WindowFunc { _ = "STUB: not implemented"; return nil }

// LAG function OID (varies by type)

// NewLead creates a LEAD(expr) window function.
func NewLead(expr Node) *WindowFunc { _ = "STUB: not implemented"; return nil }

// LEAD function OID (varies by type)

// ==============================================================================
// TIER 2 EXPRESSION UTILITIES
// ==============================================================================

// IsCaseExpr checks if an expression is a CASE expression.
func IsCaseExpr(expr Node) bool { _ = "STUB: not implemented"; return false }

// IsCoalesceExpr checks if an expression is a COALESCE expression.
func IsCoalesceExpr(expr Node) bool { _ = "STUB: not implemented"; return false }

// IsArrayExpr checks if an expression is an ARRAY constructor.
func IsArrayExpr(expr Node) bool { _ = "STUB: not implemented"; return false }

// IsScalarArrayOpExpr checks if an expression is a scalar array operation.
func IsScalarArrayOpExpr(expr Node) bool { _ = "STUB: not implemented"; return false }

// IsRowExpr checks if an expression is a ROW constructor.
func IsRowExpr(expr Node) bool { _ = "STUB: not implemented"; return false }

// IsInExpr checks if an expression is an IN operation (scalar = ANY(array)).
func IsInExpr(expr Node) bool { _ = "STUB: not implemented"; return false }

// "=" operator with ANY

// IsNotInExpr checks if an expression is a NOT IN operation (scalar <> ALL(array)).
func IsNotInExpr(expr Node) bool { _ = "STUB: not implemented"; return false }

// "<>" operator with ALL

// GetCaseWhenCount returns the number of WHEN clauses in a CASE expression.
func GetCaseWhenCount(expr Node) int { _ = "STUB: not implemented"; return 0 }

// HasCaseElse checks if a CASE expression has an ELSE clause.
func HasCaseElse(expr Node) bool { _ = "STUB: not implemented"; return false }

// GetArrayElements returns the elements of an array expression.
func GetArrayElements(expr Node) *NodeList { _ = "STUB: not implemented"; return nil }

// IsMultiDimArray checks if an array expression is multi-dimensional.
func IsMultiDimArray(expr Node) bool { _ = "STUB: not implemented"; return false }

// ==============================================================================
// TIER 3 EXPRESSION UTILITIES
// ==============================================================================

// IsAggref checks if an expression is an aggregate function call.
func IsAggref(expr Node) bool { _ = "STUB: not implemented"; return false }

// IsWindowFunc checks if an expression is a window function call.
func IsWindowFunc(expr Node) bool { _ = "STUB: not implemented"; return false }

// IsSubLink checks if an expression is a sublink (subquery).
func IsSubLink(expr Node) bool { _ = "STUB: not implemented"; return false }

// IsAggregate checks if an expression is any kind of aggregate (Aggref or WindowFunc with winagg=true).
func IsAggregate(expr Node) bool { _ = "STUB: not implemented"; return false }

// GetAggregateArgs returns the arguments of an aggregate expression.
func GetAggregateArgs(expr Node) *NodeList { _ = "STUB: not implemented"; return nil }

// HasAggregateFilter checks if an aggregate expression has a FILTER clause.
func HasAggregateFilter(expr Node) bool { _ = "STUB: not implemented"; return false }

// GetAggregateFilter returns the FILTER expression of an aggregate.
func GetAggregateFilter(expr Node) Node { _ = "STUB: not implemented"; return *new(Node) }

// IsDistinctAggregate checks if an aggregate has DISTINCT.
func IsDistinctAggregate(expr Node) bool { _ = "STUB: not implemented"; return false }

// IsStarAggregate checks if an aggregate uses * (like COUNT(*)).
func IsStarAggregate(expr Node) bool { _ = "STUB: not implemented"; return false }

// GetSubLinkType returns the type of a SubLink.
func GetSubLinkType(expr Node) SubLinkType { _ = "STUB: not implemented"; return *new(SubLinkType) }

// Invalid

// IsExistsSublink checks if an expression is an EXISTS sublink.
func IsExistsSublink(expr Node) bool { _ = "STUB: not implemented"; return false }

// IsScalarSublink checks if an expression is a scalar sublink (EXPR_SUBLINK).
func IsScalarSublink(expr Node) bool { _ = "STUB: not implemented"; return false }

// IsArraySublink checks if an expression is an ARRAY sublink.
func IsArraySublink(expr Node) bool { _ = "STUB: not implemented"; return false }

// GetSubquery returns the subquery node from a SubLink.
func GetSubquery(expr Node) Node { _ = "STUB: not implemented"; return *new(Node) }

// HasSubLinkTest checks if a SubLink has a test expression.
func HasSubLinkTest(expr Node) bool { _ = "STUB: not implemented"; return false }

// GetSubLinkTest returns the test expression from a SubLink.
func GetSubLinkTest(expr Node) Node { _ = "STUB: not implemented"; return *new(Node) }

// IsOrderedSetAggregate checks if an aggregate is an ordered-set aggregate.
func IsOrderedSetAggregate(expr Node) bool { _ = "STUB: not implemented"; return false }

// GetAggregateDirectArgs returns the direct arguments of an ordered-set aggregate.
func GetAggregateDirectArgs(expr Node) *NodeList { _ = "STUB: not implemented"; return nil }

// GetAggregateOrderBy returns the ORDER BY expressions of an aggregate.
func GetAggregateOrderBy(expr Node) *NodeList { _ = "STUB: not implemented"; return nil }

// GetWindowFuncRef returns the window reference index from a WindowFunc.
func GetWindowFuncRef(expr Node) Index { _ = "STUB: not implemented"; return *new(Index) }

// IsSimpleWindowAgg checks if a WindowFunc is a simple aggregate.
func IsSimpleWindowAgg(expr Node) bool { _ = "STUB: not implemented"; return false }

// ==============================================================================
// PHASE 1F: PRIMITIVE EXPRESSION COMPLETION PART 1
// Core primitive expressions from PostgreSQL primnodes.h
// ==============================================================================

// RowCompareType represents row comparison types
// Ported from postgres/src/include/nodes/primnodes.h:1445-1452
type RowCompareType int

const (
	ROWCOMPARE_LT RowCompareType = 1 // BTLessStrategyNumber
	ROWCOMPARE_LE RowCompareType = 2 // BTLessEqualStrategyNumber
	ROWCOMPARE_EQ RowCompareType = 3 // BTEqualStrategyNumber
	ROWCOMPARE_GE RowCompareType = 4 // BTGreaterEqualStrategyNumber
	ROWCOMPARE_GT RowCompareType = 5 // BTGreaterStrategyNumber
	ROWCOMPARE_NE RowCompareType = 6 // no such btree strategy
)

// MinMaxOp represents MIN/MAX operation types
// Ported from postgres/src/include/nodes/primnodes.h:1499-1502
type MinMaxOp int

const (
	IS_GREATEST MinMaxOp = iota
	IS_LEAST
)

// SQLValueFunctionOp represents SQL value function operation types
// Ported from postgres/src/include/nodes/primnodes.h:1522-1541
type SQLValueFunctionOp int

const (
	SVFOP_CURRENT_DATE SQLValueFunctionOp = iota
	SVFOP_CURRENT_TIME
	SVFOP_CURRENT_TIME_N
	SVFOP_CURRENT_TIMESTAMP
	SVFOP_CURRENT_TIMESTAMP_N
	SVFOP_LOCALTIME
	SVFOP_LOCALTIME_N
	SVFOP_LOCALTIMESTAMP
	SVFOP_LOCALTIMESTAMP_N
	SVFOP_CURRENT_ROLE
	SVFOP_CURRENT_USER
	SVFOP_USER
	SVFOP_SESSION_USER
	SVFOP_CURRENT_CATALOG
	SVFOP_CURRENT_SCHEMA
)

// XmlExprOp represents XML expression operation types
// Ported from postgres/src/include/nodes/primnodes.h:1577-1586
type XmlExprOp int

const (
	IS_XMLCONCAT    XmlExprOp = iota // XMLCONCAT(args)
	IS_XMLELEMENT                    // XMLELEMENT(name, xml_attributes, args)
	IS_XMLFOREST                     // XMLFOREST(xml_attributes)
	IS_XMLPARSE                      // XMLPARSE(text, is_doc, preserve_ws)
	IS_XMLPI                         // XMLPI(name [, args])
	IS_XMLROOT                       // XMLROOT(xml, version, standalone)
	IS_XMLSERIALIZE                  // XMLSERIALIZE(is_document, xmlval, indent)
	IS_DOCUMENT                      // xmlval IS DOCUMENT
)

// TableFuncType represents table function types
// Ported from postgres/src/include/nodes/primnodes.h:98-101
type TableFuncType int

const (
	TFT_XMLTABLE TableFuncType = iota
	TFT_JSON_TABLE
)

// GroupingFunc represents a GROUPING function call
// Ported from postgres/src/include/nodes/primnodes.h:537-548
type GroupingFunc struct {
	BaseExpr
	Args        *NodeList // Arguments, not evaluated but kept for EXPLAIN
	Refs        *NodeList // Resource sort group refs of arguments
	Cols        *NodeList // Actual column positions set by planner
	AggLevelsUp Index     // Same as Aggref.agglevelsup
}

// NewGroupingFunc creates a new GroupingFunc node.
func NewGroupingFunc(args *NodeList, refs, cols *NodeList, aggLevelsUp Index, location int) *GroupingFunc {
	_ = "STUB: not implemented"
	return nil
}

func (g *GroupingFunc) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (g *GroupingFunc) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of GroupingFunc
func (g *GroupingFunc) SqlString() string { _ = "STUB: not implemented"; return "" }

// WindowFuncRunCondition represents a window function run condition
// Ported from postgres/src/include/nodes/primnodes.h:596-609
type WindowFuncRunCondition struct {
	BaseExpr
	Opno        Oid        // PG_OPERATOR OID of the operator
	InputCollid Oid        // OID of collation that operator should use
	WfuncLeft   bool       // True if WindowFunc belongs on the left
	Arg         Expression // The argument expression
}

// NewWindowFuncRunCondition creates a new WindowFuncRunCondition node.
func NewWindowFuncRunCondition(opno, inputCollid Oid, wfuncLeft bool, arg Expression, location int) *WindowFuncRunCondition {
	_ = "STUB: not implemented"
	return nil
}

func (w *WindowFuncRunCondition) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (w *WindowFuncRunCondition) String() string { _ = "STUB: not implemented"; return "" }

// MergeSupportFunc represents a merge support function
// Ported from postgres/src/include/nodes/primnodes.h:628-635
type MergeSupportFunc struct {
	BaseExpr
	MsfType   Oid // Type OID of result
	MsfCollid Oid // OID of collation, or InvalidOid if none
}

// NewMergeSupportFunc creates a new MergeSupportFunc node.
func NewMergeSupportFunc(msfType, msfCollid Oid, location int) *MergeSupportFunc {
	_ = "STUB: not implemented"
	return nil
}

func (m *MergeSupportFunc) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (m *MergeSupportFunc) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of MergeSupportFunc
func (m *MergeSupportFunc) SqlString() string { _ = "STUB: not implemented"; return "" }

// NamedArgExpr represents a named argument expression
// Ported from postgres/src/include/nodes/primnodes.h:787-795
type NamedArgExpr struct {
	BaseExpr
	Arg       Expression // The argument expression
	Name      string     // The name
	ArgNumber int        // Argument's number in positional notation
}

// NewNamedArgExpr creates a new NamedArgExpr node.
func NewNamedArgExpr(arg Expression, name string, argNumber, location int) *NamedArgExpr {
	_ = "STUB: not implemented"
	return nil
}

func (n *NamedArgExpr) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (n *NamedArgExpr) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of NamedArgExpr
func (n *NamedArgExpr) SqlString() string {
	_ = "STUB: not implemented"
	// Format as: name => value
	return ""
}

// CaseTestExpr represents a CASE test expression
// Ported from postgres/src/include/nodes/primnodes.h:1352-1359
type CaseTestExpr struct {
	BaseExpr
	TypeId    Oid // Type for substituted value
	TypeMod   int // Typemod for substituted value
	Collation Oid // Collation for the substituted value
}

// NewCaseTestExpr creates a new CaseTestExpr node.
func NewCaseTestExpr(typeId Oid, typeMod int, collation Oid, location int) *CaseTestExpr {
	_ = "STUB: not implemented"
	return nil
}

func (c *CaseTestExpr) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (c *CaseTestExpr) String() string { _ = "STUB: not implemented"; return "" }

// MinMaxExpr represents a MIN/MAX expression
// Ported from postgres/src/include/nodes/primnodes.h:1506-1517
type MinMaxExpr struct {
	BaseExpr
	MinMaxType   Oid       // Common type of arguments and result
	MinMaxCollid Oid       // OID of collation of result
	InputCollid  Oid       // OID of collation that function should use
	Op           MinMaxOp  // Function to execute
	Args         *NodeList // The arguments
}

// NewMinMaxExpr creates a new MinMaxExpr node.
func NewMinMaxExpr(minMaxType, minMaxCollid, inputCollid Oid, op MinMaxOp, args *NodeList, location int) *MinMaxExpr {
	_ = "STUB: not implemented"
	return nil
}

func (m *MinMaxExpr) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (m *MinMaxExpr) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of MinMaxExpr
func (m *MinMaxExpr) SqlString() string { _ = "STUB: not implemented"; return "" }

// Write the function name

// Write the arguments

// RowCompareExpr represents a row comparison expression
// Ported from postgres/src/include/nodes/primnodes.h:1463-1474
type RowCompareExpr struct {
	BaseExpr
	Rctype       RowCompareType // LT LE GE or GT, never EQ or NE
	Opnos        []Oid          // OID list of pairwise comparison ops
	Opfamilies   []Oid          // OID list of containing operator families
	InputCollids []Oid          // OID list of collations for comparisons
	Largs        []Expression   // The left-hand input arguments
	Rargs        []Expression   // The right-hand input arguments
}

// NewRowCompareExpr creates a new RowCompareExpr node.
func NewRowCompareExpr(rctype RowCompareType, opnos, opfamilies, inputCollids []Oid, largs, rargs []Expression, location int) *RowCompareExpr {
	_ = "STUB: not implemented"
	return nil
}

func (r *RowCompareExpr) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (r *RowCompareExpr) String() string { _ = "STUB: not implemented"; return "" }

// SQLValueFunction represents parameterless functions with special grammar
// Ported from postgres/src/include/nodes/primnodes.h:1553-1563
type SQLValueFunction struct {
	BaseExpr
	Op      SQLValueFunctionOp // Which function this is
	Type    Oid                // Result type
	TypeMod int                // Result typmod
}

// NewSQLValueFunction creates a new SQLValueFunction node.
func NewSQLValueFunction(op SQLValueFunctionOp, typ Oid, typeMod, location int) *SQLValueFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *SQLValueFunction) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (s *SQLValueFunction) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of SQLValueFunction
func (s *SQLValueFunction) SqlString() string { _ = "STUB: not implemented"; return "" }

// XmlExpr represents various SQL/XML functions requiring special grammar
// Ported from postgres/src/include/nodes/primnodes.h:1596-1618
type XmlExpr struct {
	BaseExpr
	Op        XmlExprOp     // XML function ID
	Name      string        // Name in xml(NAME foo ...) syntaxes
	NamedArgs *NodeList     // Non-XML expressions for xml_attributes
	ArgNames  *NodeList     // Parallel list of String values
	Args      *NodeList     // List of expressions
	Xmloption XmlOptionType // DOCUMENT or CONTENT
	Indent    bool          // INDENT option for XMLSERIALIZE
	Type      Oid           // Target type for XMLSERIALIZE
	TypeMod   int           // Target typmod for XMLSERIALIZE
}

// NewXmlExpr creates a new XmlExpr node.
func NewXmlExpr(op XmlExprOp, name string, namedArgs, argNames, args *NodeList, xmloption XmlOptionType, indent bool, typ Oid, typeMod, location int) *XmlExpr {
	_ = "STUB: not implemented"
	return nil
}

func (x *XmlExpr) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (x *XmlExpr) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of XmlExpr
func (x *XmlExpr) SqlString() string { _ = "STUB: not implemented"; return "" }

// NamedArgs hold the XMLATTRIBUTES (each a ResTarget rendering as
// `value [AS name]`), which come right after the element name.

// Args are [xml, version, standalone]. The version is a NULL A_Const for
// VERSION NO VALUE; the standalone is an integer A_Const holding an
// XmlStandaloneType.

// Type information would be in the Type field, but for basic implementation
// we'll just use a placeholder

// TableFunc represents a table function such as XMLTABLE and JSON_TABLE
// Ported from postgres/src/include/nodes/primnodes.h:109-146
type TableFunc struct {
	BaseNode
	Functype        TableFuncType // XMLTABLE or JSON_TABLE
	NsUris          []Expression  // List of namespace URI expressions
	NsNames         []string      // List of namespace names or NULL
	Docexpr         Expression    // Input document expression
	Rowexpr         Expression    // Row filter expression
	Colnames        []string      // Column names (list of String)
	Coltypes        []Oid         // OID list of column type OIDs
	Coltypmods      []int         // Integer list of column typmods
	Colcollations   []Oid         // OID list of column collation OIDs
	Colexprs        []Expression  // List of column filter expressions
	Coldefexprs     []Expression  // List of column default expressions
	Colvalexprs     []Expression  // JSON_TABLE: list of column value expressions
	Passingvalexprs []Expression  // JSON_TABLE: list of PASSING argument expressions
	Notnulls        []bool        // Nullability flag for each output column
	Plan            Node          // JSON_TABLE plan
	Ordinalitycol   int           // Counts from 0; -1 if none specified
}

// NewTableFunc creates a new TableFunc node.
func NewTableFunc(functype TableFuncType, nsUris []Expression, nsNames []string, docexpr, rowexpr Expression, colnames []string, coltypes []Oid, coltypmods []int, colcollations []Oid, colexprs, coldefexprs, colvalexprs, passingvalexprs []Expression, notnulls []bool, plan Node, ordinalitycol, location int) *TableFunc {
	_ = "STUB: not implemented"
	return nil
}

func (t *TableFunc) String() string { _ = "STUB: not implemented"; return "" }

func (t *TableFunc) StatementType() string { _ = "STUB: not implemented"; return "" }

// IntoClause represents target information for SELECT INTO, CREATE TABLE AS, and CREATE MATERIALIZED VIEW
// Ported from postgres/src/include/nodes/primnodes.h:158-171
type IntoClause struct {
	BaseNode
	Rel            *RangeVar      // Target relation name
	ColNames       *NodeList      // Column names to assign, or NIL
	AccessMethod   string         // Table access method
	Options        *NodeList      // Options from WITH clause
	OnCommit       OnCommitAction // What do we do at COMMIT?
	TableSpaceName string         // Table space to use, or NULL
	ViewQuery      Node           // Materialized view's SELECT query
	SkipData       bool           // True for WITH NO DATA
}

// NewIntoClause creates a new IntoClause node.
func NewIntoClause(rel *RangeVar, colNames *NodeList, accessMethod string, options *NodeList, onCommit OnCommitAction, tableSpaceName string, viewQuery Node, skipData bool, location int) *IntoClause {
	_ = "STUB: not implemented"
	return nil
}

func (i *IntoClause) String() string { _ = "STUB: not implemented"; return "" }

func (i *IntoClause) StatementType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of the IntoClause
// TargetString returns just the target relation and column list without "INTO"
func (i *IntoClause) TargetString() string { _ = "STUB: not implemented"; return "" }

// Add the target relation

// Add column names if specified

func (i *IntoClause) SqlString() string { _ = "STUB: not implemented"; return "" }

// Add persistence keywords if present

// Add the target relation

// Add column names if specified

// Add WITH options if present

// Add ON COMMIT clause

// Add TABLESPACE clause

// Add WITH NO DATA if specified

// MergeAction represents a MERGE action
// Ported from postgres/src/include/nodes/primnodes.h:2003-2013
type MergeAction struct {
	BaseNode
	MatchKind    MergeMatchKind // MATCHED/NOT MATCHED BY SOURCE/TARGET
	CommandType  CmdType        // INSERT/UPDATE/DELETE/DO NOTHING
	Override     OverridingKind // OVERRIDING clause
	Qual         Node           // Transformed WHEN conditions
	TargetList   []*TargetEntry // The target list (of TargetEntry)
	UpdateColnos []AttrNumber   // Target attribute numbers of an UPDATE
}

// NewMergeAction creates a new MergeAction node.
func NewMergeAction(matchKind MergeMatchKind, commandType CmdType, override OverridingKind, qual Node, targetList []*TargetEntry, updateColnos []AttrNumber, location int) *MergeAction {
	_ = "STUB: not implemented"
	return nil
}

func (m *MergeAction) String() string { _ = "STUB: not implemented"; return "" }

func (m *MergeAction) StatementType() string { _ = "STUB: not implemented"; return "" }
