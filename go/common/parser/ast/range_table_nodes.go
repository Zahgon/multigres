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
// Package ast provides PostgreSQL AST range table and FROM clause node definitions.
// This file contains range table infrastructure nodes essential for FROM clause and JOIN support.
// Ported from postgres/src/include/nodes/parsenodes.h and primnodes.h
package ast

// ==============================================================================
// RANGE TABLE AND FROM CLAUSE INFRASTRUCTURE - Phase 1D Implementation
// Essential range table nodes for FROM clause and JOIN support
// Ported from postgres/src/include/nodes/parsenodes.h and primnodes.h
// ==============================================================================

// Type aliases for PostgreSQL types
type (
	AclMode     uint32  // Access control mode bitmask
	Cardinality float64 // Row count estimates
)

// TableFunc placeholder removed - now implemented in expressions.go

// RTEKind represents the type of a Range Table Entry.
// Ported from postgres/src/include/nodes/parsenodes.h:1022-1033
type RTEKind int

const (
	RTE_RELATION        RTEKind = iota // ordinary relation reference
	RTE_SUBQUERY                       // subquery in FROM
	RTE_JOIN                           // join
	RTE_FUNCTION                       // function in FROM
	RTE_TABLEFUNC                      // TableFunc(.., column list)
	RTE_VALUES                         // VALUES (<exprlist>), (<exprlist>), ...
	RTE_CTE                            // common table expr (WITH list element)
	RTE_NAMEDTUPLESTORE                // tuplestore, e.g. for AFTER triggers
	RTE_RESULT                         // RTE represents an empty FROM clause
)

// String returns the string representation of RTEKind.
func (k RTEKind) String() string { _ = "STUB: not implemented"; return "" }

// RangeTblEntry represents a range table entry which describes a table or subquery in the FROM clause.
// This is a complex structure that supports multiple types of table sources.
// Ported from postgres/src/include/nodes/parsenodes.h:1038-1251
type RangeTblEntry struct {
	BaseNode

	// Fields valid in all RTEs
	Alias    *Alias  // user-written alias clause, if any
	Eref     *Alias  // expanded reference names
	RteKind  RTEKind // see RTEKind enum above
	Lateral  bool    // was LATERAL specified?
	InFromCl bool    // present in FROM clause?

	// Fields valid for a plain relation RTE (RTE_RELATION)
	Relid         Oid                // OID of the relation
	Inh           bool               // inheritance requested?
	RelKind       byte               // relation kind (see pg_class.relkind)
	RelLockMode   int                // lock level that query requires on the rel
	PermInfoIndex int                // index of RTEPermissionInfo entry, or 0
	TableSample   *TableSampleClause // sampling info, or NULL

	// Fields valid for a subquery RTE (RTE_SUBQUERY)
	Subquery        *Query // the sub-query
	SecurityBarrier bool   // is from security_barrier view?

	// Fields valid for a join RTE (RTE_JOIN)
	JoinType       JoinType  // type of join
	JoinMergedCols int       // number of merged (JOIN USING) columns
	JoinAliasVars  *NodeList // list of alias-var expansions
	JoinLeftCols   []int     // left-side input column numbers
	JoinRightCols  []int     // right-side input column numbers
	JoinUsingAlias *Alias    // alias clause attached directly to JOIN/USING

	// Fields valid for a function RTE (RTE_FUNCTION)
	Functions      []*RangeTblFunction // list of RangeTblFunction nodes
	FuncOrdinality bool                // is this called WITH ORDINALITY?

	// Fields valid for a TableFunc RTE (RTE_TABLEFUNC)
	TableFunc *TableFunc // table function specification

	// Fields valid for a values RTE (RTE_VALUES)
	ValuesLists []*NodeList // list of expression lists

	// Fields valid for a CTE RTE (RTE_CTE)
	CteName       string // name of the WITH list item
	CteLevelsUp   Index  // number of query levels up
	SelfReference bool   // is this a recursive self-reference?

	// Fields valid for CTE, VALUES, ENR, and TableFunc RTEs
	ColTypes      []Oid // OID list of column type OIDs
	ColTypMods    []int // integer list of column typmods
	ColCollations []Oid // OID list of column collation OIDs

	// Fields valid for ENR RTEs (RTE_NAMEDTUPLESTORE)
	EnrName   string      // name of ephemeral named relation
	EnrTuples Cardinality // estimated or actual from caller

	// Security-related fields
	SecurityQuals *NodeList // security barrier quals to apply, if any
}

// NewRangeTblEntry creates a new RangeTblEntry node.
func NewRangeTblEntry(rteKind RTEKind, alias *Alias) *RangeTblEntry {
	_ = "STUB: not implemented"
	return nil
}

func (r *RangeTblEntry) String() string { _ = "STUB: not implemented"; return "" }

func (r *RangeTblEntry) StatementType() string { _ = "STUB: not implemented"; return "" }

// RangeSubselect represents a subquery in FROM clause.
// Ported from postgres/src/include/nodes/parsenodes.h:615-621
type RangeSubselect struct {
	BaseNode
	Lateral  bool   // does it have LATERAL prefix?
	Subquery Node   // the untransformed sub-select clause
	Alias    *Alias // table alias & optional column aliases
}

// NewRangeSubselect creates a new RangeSubselect node.
func NewRangeSubselect(lateral bool, subquery Node, alias *Alias) *RangeSubselect {
	_ = "STUB: not implemented"
	return nil
}

func (r *RangeSubselect) String() string { _ = "STUB: not implemented"; return "" }

func (r *RangeSubselect) StatementType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of the RangeSubselect.
func (r *RangeSubselect) SqlString() string { _ = "STUB: not implemented"; return "" }

// RangeFunction represents a function call appearing in a FROM clause.
// Supports ROWS FROM() syntax and WITH ORDINALITY.
// Ported from postgres/src/include/nodes/parsenodes.h:637-647
type RangeFunction struct {
	BaseNode
	Lateral    bool      // does it have LATERAL prefix?
	Ordinality bool      // does it have WITH ORDINALITY suffix?
	IsRowsFrom bool      // is result of ROWS FROM() syntax?
	Functions  *NodeList // list of per-function information (each item is a NodeList with function + column definitions)
	Alias      *Alias    // table alias & optional column aliases
	ColDefList *NodeList // list of ColumnDef nodes to describe result of function returning RECORD
}

// NewRangeFunction creates a new RangeFunction node.
func NewRangeFunction(lateral, ordinality, isRowsFrom bool, functions *NodeList, alias *Alias, colDefList *NodeList) *RangeFunction {
	_ = "STUB: not implemented"
	return nil
}

func (r *RangeFunction) String() string { _ = "STUB: not implemented"; return "" }

func (r *RangeFunction) StatementType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of the RangeFunction.
func (r *RangeFunction) SqlString() string { _ = "STUB: not implemented"; return "" }

// Each item is a NodeList containing function + optional column definitions

// Simple function call - Functions contains a single NodeList with one function

// For non-ROWS FROM, Functions.Items[0] is the NodeList containing the function

// A RECORD-returning function carries a column-definition list, written as
// `... AS alias (col type, ...)` or, when unaliased, `... AS (col type, ...)`.

// renderColDefList renders a parenthesized list of column definitions used by
// RECORD-returning functions, e.g. "(a int, b text)".
func renderColDefList(list *NodeList) string { _ = "STUB: not implemented"; return "" }

// RangeTableFunc represents raw form of "table functions" such as XMLTABLE.
// Note: JSON_TABLE uses JsonTable node, not RangeTableFunc.
// Ported from postgres/src/include/nodes/parsenodes.h:655-665
type RangeTableFunc struct {
	BaseNode
	Lateral    bool      // does it have LATERAL prefix?
	DocExpr    Node      // document expression
	RowExpr    Node      // row generator expression
	Namespaces *NodeList // list of namespaces as ResTarget
	Columns    *NodeList // list of RangeTableFuncCol
	Alias      *Alias    // table alias & optional column aliases
}

// NewRangeTableFunc creates a new RangeTableFunc node.
func NewRangeTableFunc(lateral bool, docExpr, rowExpr Node, namespaces, columns *NodeList, alias *Alias, location int) *RangeTableFunc {
	_ = "STUB: not implemented"
	return nil
}

func (r *RangeTableFunc) String() string { _ = "STUB: not implemented"; return "" }

func (r *RangeTableFunc) StatementType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of the RangeTableFunc (XMLTABLE).
func (r *RangeTableFunc) SqlString() string { _ = "STUB: not implemented"; return "" }

// For now, we assume this is XMLTABLE since that's what we're implementing

// Optional XMLNAMESPACES(...) clause comes first. Each namespace is a
// ResTarget: `uri AS name`, or `DEFAULT uri` when unnamed.

// XMLTABLE syntax: XMLTABLE(xpath_expression PASSING document_expression COLUMNS ...)
// RowExpr is the XPath expression, DocExpr is the document

// RangeTableFuncCol represents one column in a RangeTableFunc->columns.
// If ForOrdinality is true (FOR ORDINALITY), then the column is an int4 column
// and the rest of the fields are ignored.
// Ported from postgres/src/include/nodes/parsenodes.h:673-683
type RangeTableFuncCol struct {
	BaseNode
	ColName       string    // name of generated column
	TypeName      *TypeName // type of generated column
	ForOrdinality bool      // does it have FOR ORDINALITY?
	IsNotNull     bool      // does it have NOT NULL?
	ColExpr       Node      // column filter expression
	ColDefExpr    Node      // column default value expression
}

// NewRangeTableFuncCol creates a new RangeTableFuncCol node.
func NewRangeTableFuncCol(colName string, typeName *TypeName, forOrdinality, isNotNull bool, colExpr, colDefExpr Node, location int) *RangeTableFuncCol {
	_ = "STUB: not implemented"
	return nil
}

func (r *RangeTableFuncCol) String() string { _ = "STUB: not implemented"; return "" }

func (r *RangeTableFuncCol) StatementType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of the RangeTableFuncCol.
func (r *RangeTableFuncCol) SqlString() string { _ = "STUB: not implemented"; return "" }

// Quote column name if it needs quoting (contains special chars, is a keyword, etc.)

// Add PATH clause if ColExpr is present

// Add DEFAULT clause if ColDefExpr is present

// RangeTableSample represents TABLESAMPLE appearing in a raw FROM clause.
// This node represents: <relation> TABLESAMPLE <method> (<params>) REPEATABLE (<num>)
// Ported from postgres/src/include/nodes/parsenodes.h:695-703
type RangeTableSample struct {
	BaseNode
	Relation   Node      // relation to be sampled
	Method     *NodeList // sampling method name (possibly qualified)
	Args       *NodeList // argument(s) for sampling method
	Repeatable Node      // REPEATABLE expression, or NULL if none
}

// NewRangeTableSample creates a new RangeTableSample node.
func NewRangeTableSample(relation Node, method, args *NodeList, repeatable Node, location int) *RangeTableSample {
	_ = "STUB: not implemented"
	return nil
}

func (r *RangeTableSample) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of RangeTableSample
func (r *RangeTableSample) SqlString() string {
	_ = "STUB: not implemented"

	// Add relation
	return ""
}

// Add TABLESAMPLE keyword

// Add method name

// For TABLESAMPLE method names, treat as identifiers not string literals

// Add arguments if any

// Add REPEATABLE clause if present

func (r *RangeTableSample) StatementType() string { _ = "STUB: not implemented"; return "" }

// RangeTblFunction represents RangeTblEntry subsidiary data for one function in a FUNCTION RTE.
// Used when a function had a column definition list for an otherwise-unspecified RECORD result.
// Ported from postgres/src/include/nodes/parsenodes.h:1317-1337
type RangeTblFunction struct {
	BaseNode
	FuncExpr          Node     // expression tree for func call
	FuncColCount      int      // number of columns it contributes to RTE
	FuncColNames      []string // column names (list of String)
	FuncColTypes      []Oid    // OID list of column type OIDs
	FuncColTypMods    []int    // integer list of column typmods
	FuncColCollations []Oid    // OID list of column collation OIDs
	FuncParams        []int    // PARAM_EXEC Param IDs affecting this func (set during planning)
}

// NewRangeTblFunction creates a new RangeTblFunction node.
func NewRangeTblFunction(funcExpr Node, funcColCount int) *RangeTblFunction {
	_ = "STUB: not implemented"
	return nil
}

func (r *RangeTblFunction) String() string { _ = "STUB: not implemented"; return "" }

func (r *RangeTblFunction) StatementType() string { _ = "STUB: not implemented"; return "" }

// RTEPermissionInfo represents per-relation information for permission checking.
// Added to the Query node by the parser when adding the corresponding RTE to the query range table.
// Ported from postgres/src/include/nodes/parsenodes.h:1286-1297
type RTEPermissionInfo struct {
	BaseNode
	Relid         Oid     // relation OID
	Inh           bool    // separately check inheritance children?
	RequiredPerms AclMode // bitmask of required access permissions
	CheckAsUser   Oid     // if valid, check access as this role
	SelectedCols  []int   // columns needing SELECT permission (simplified from Bitmapset)
	InsertedCols  []int   // columns needing INSERT permission (simplified from Bitmapset)
	UpdatedCols   []int   // columns needing UPDATE permission (simplified from Bitmapset)
}

// NewRTEPermissionInfo creates a new RTEPermissionInfo node.
func NewRTEPermissionInfo(relid Oid, inh bool, requiredPerms AclMode, checkAsUser Oid) *RTEPermissionInfo {
	_ = "STUB: not implemented"
	return nil
}

func (r *RTEPermissionInfo) String() string { _ = "STUB: not implemented"; return "" }

func (r *RTEPermissionInfo) StatementType() string { _ = "STUB: not implemented"; return "" }

// RangeTblRef represents a range table reference from primnodes.h.
// This is a simple reference to an entry in the range table by index.
// Ported from postgres/src/include/nodes/primnodes.h:2243-2247
type RangeTblRef struct {
	BaseNode
	RtIndex int // index into the range table
}

// NewRangeTblRef creates a new RangeTblRef node.
func NewRangeTblRef(rtIndex int) *RangeTblRef { _ = "STUB: not implemented"; return nil }

func (r *RangeTblRef) String() string { _ = "STUB: not implemented"; return "" }

func (r *RangeTblRef) ExpressionType() string { _ = "STUB: not implemented"; return "" }
