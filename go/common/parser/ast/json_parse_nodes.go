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
package ast

import (
	"strings"
)

// JSON Parse Tree Support Nodes - Phase 1E
// This file contains all JSON-related parse tree nodes from PostgreSQL
// Based on postgres/src/include/nodes/parsenodes.h and primnodes.h

// JsonEncoding represents the JSON ENCODING clause
// Ported from postgres/src/include/nodes/primnodes.h:1636-1643
type JsonEncoding int

const (
	JS_ENC_DEFAULT JsonEncoding = iota // unspecified
	JS_ENC_UTF8
	JS_ENC_UTF16
	JS_ENC_UTF32
)

// JsonFormatType represents the JSON format types
// Ported from postgres/src/include/nodes/primnodes.h:1644-1651
type JsonFormatType int

const (
	JS_FORMAT_DEFAULT JsonFormatType = iota // unspecified
	JS_FORMAT_JSON                          // FORMAT JSON [ENCODING ...]
	JS_FORMAT_JSONB                         // implicit internal format for RETURNING jsonb
)

// JsonFormat represents a JSON FORMAT clause
// Ported from postgres/src/include/nodes/primnodes.h:1648-1655
type JsonFormat struct {
	BaseNode
	FormatType JsonFormatType // format type
	Encoding   JsonEncoding   // JSON encoding
}

// JsonReturning represents a transformed JSON RETURNING clause
// Ported from postgres/src/include/nodes/primnodes.h:1660-1668
type JsonReturning struct {
	BaseNode
	Format *JsonFormat // output JSON format
	Typid  Oid         // target type Oid
	Typmod int32       // target type modifier
}

// JsonValueExpr represents a JSON value expression (expr [FORMAT JsonFormat])
// Ported from postgres/src/include/nodes/primnodes.h:1680-1691
type JsonValueExpr struct {
	BaseNode
	RawExpr       Node        // user-specified expression
	FormattedExpr Expr        // coerced formatted expression
	Format        *JsonFormat // FORMAT clause, if specified
}

// JsonValueType represents JSON item types in IS JSON predicate
// Ported from postgres/src/include/nodes/primnodes.h:1723-1730
type JsonValueType int

const (
	JS_TYPE_ANY    JsonValueType = iota // IS JSON [VALUE]
	JS_TYPE_OBJECT                      // IS JSON OBJECT
	JS_TYPE_ARRAY                       // IS JSON ARRAY
	JS_TYPE_SCALAR                      // IS JSON SCALAR
)

// JsonWrapper represents WRAPPER clause for JSON_QUERY()
// Ported from postgres/src/include/nodes/primnodes.h:1761-1768
type JsonWrapper int

const (
	JSW_UNSPEC JsonWrapper = iota
	JSW_NONE
	JSW_CONDITIONAL
	JSW_UNCONDITIONAL
)

// JsonBehaviorType represents behavior types used in SQL/JSON ON ERROR/EMPTY clauses
// Ported from postgres/src/include/nodes/primnodes.h:1771-1783
type JsonBehaviorType int

const (
	JSON_BEHAVIOR_NULL JsonBehaviorType = iota
	JSON_BEHAVIOR_ERROR
	JSON_BEHAVIOR_EMPTY
	JSON_BEHAVIOR_TRUE
	JSON_BEHAVIOR_FALSE
	JSON_BEHAVIOR_UNKNOWN
	JSON_BEHAVIOR_EMPTY_ARRAY
	JSON_BEHAVIOR_EMPTY_OBJECT
	JSON_BEHAVIOR_DEFAULT
)

// JsonBehavior represents specifications for ON ERROR / ON EMPTY behaviors
// Ported from postgres/src/include/nodes/primnodes.h:1786-1797
type JsonBehavior struct {
	BaseNode
	Btype  JsonBehaviorType
	Expr   Node
	Coerce bool
}

// JsonExprOp represents enumeration of SQL/JSON query function types
// Ported from postgres/src/include/nodes/primnodes.h:1802-1808
type JsonExprOp int

const (
	JSON_EXISTS_OP JsonExprOp = iota // JSON_EXISTS()
	JSON_QUERY_OP                    // JSON_QUERY()
	JSON_VALUE_OP                    // JSON_VALUE()
	JSON_TABLE_OP                    // JSON_TABLE()
)

// JsonQuotes represents [KEEP|OMIT] QUOTES clause for JSON_QUERY()
// Ported from postgres/src/include/nodes/parsenodes.h:1773-1778
type JsonQuotes int

const (
	JS_QUOTES_UNSPEC JsonQuotes = iota // unspecified
	JS_QUOTES_KEEP                     // KEEP QUOTES
	JS_QUOTES_OMIT                     // OMIT QUOTES
)

// JsonOutput represents a JSON output specification
// Ported from postgres/src/include/nodes/parsenodes.h:1751-1756
type JsonOutput struct {
	BaseNode
	TypeName  *TypeName      // RETURNING type name, if specified
	Returning *JsonReturning // RETURNING FORMAT clause and type Oids
}

// JsonArgument represents an argument from JSON PASSING clause
// Ported from postgres/src/include/nodes/parsenodes.h:1762-1767
type JsonArgument struct {
	BaseNode
	Val  *JsonValueExpr // argument value expression
	Name string         // argument name
}

// JsonFuncExpr represents untransformed function expressions for SQL/JSON query functions
// Ported from postgres/src/include/nodes/parsenodes.h:1785-1800
type JsonFuncExpr struct {
	BaseNode
	Op          JsonExprOp     // expression type
	ColumnName  string         // JSON_TABLE() column name or NULL
	ContextItem *JsonValueExpr // context item expression
	Pathspec    Node           // JSON path specification expression
	Passing     *NodeList      // list of PASSING clause arguments, if any
	Output      *JsonOutput    // output clause, if specified
	OnEmpty     *JsonBehavior  // ON EMPTY behavior
	OnError     *JsonBehavior  // ON ERROR behavior
	Wrapper     JsonWrapper    // array wrapper behavior (JSON_QUERY only)
	Quotes      JsonQuotes     // omit or keep quotes? (JSON_QUERY only)
}

// JsonTablePathSpec represents untransformed specification of JSON path expression with an optional name
// Ported from postgres/src/include/nodes/parsenodes.h:1807-1815
type JsonTablePathSpec struct {
	BaseNode
	StringExpr   Node   // JSON path string expression
	Name         string // optional path name
	NameLocation int    // location of 'name'
}

// JsonTableColumnType represents enumeration of JSON_TABLE column types
// Ported from postgres/src/include/nodes/parsenodes.h:1838-1845
type JsonTableColumnType int

const (
	JTC_FOR_ORDINALITY JsonTableColumnType = iota
	JTC_REGULAR
	JTC_EXISTS
	JTC_FORMATTED
	JTC_NESTED
)

// JsonTable represents untransformed JSON_TABLE
// Ported from postgres/src/include/nodes/parsenodes.h:1821-1832
type JsonTable struct {
	BaseNode
	ContextItem *JsonValueExpr     // context item expression
	Pathspec    *JsonTablePathSpec // JSON path specification
	Passing     *NodeList          // list of PASSING clause arguments, if any
	Columns     *NodeList          // list of JsonTableColumn
	OnError     *JsonBehavior      // ON ERROR behavior
	Alias       *Alias             // table alias in FROM clause
	Lateral     bool               // does it have LATERAL prefix?
}

// JsonTableColumn represents untransformed JSON_TABLE column
// Ported from postgres/src/include/nodes/parsenodes.h:1851-1865
type JsonTableColumn struct {
	BaseNode
	Coltype  JsonTableColumnType // column type
	Name     string              // column name
	TypeName *TypeName           // column type name
	Pathspec *JsonTablePathSpec  // JSON path specification
	Format   *JsonFormat         // JSON format clause, if specified
	Wrapper  JsonWrapper         // WRAPPER behavior for formatted columns
	Quotes   JsonQuotes          // omit or keep quotes on scalar strings?
	Columns  *NodeList           // nested columns
	OnEmpty  *JsonBehavior       // ON EMPTY behavior
	OnError  *JsonBehavior       // ON ERROR behavior
}

// JsonKeyValue represents untransformed JSON object key-value pair
// Ported from postgres/src/include/nodes/parsenodes.h:1872-1877
type JsonKeyValue struct {
	BaseNode
	Key   Expr           // key expression
	Value *JsonValueExpr // JSON value expression
}

// JsonParseExpr represents untransformed JSON()
// Ported from postgres/src/include/nodes/parsenodes.h:1883-1890
type JsonParseExpr struct {
	BaseNode
	Expr       *JsonValueExpr // string expression
	Output     *JsonOutput    // RETURNING clause, if specified
	UniqueKeys bool           // WITH UNIQUE KEYS?
}

// JsonScalarExpr represents untransformed JSON_SCALAR()
// Ported from postgres/src/include/nodes/parsenodes.h:1896-1902
type JsonScalarExpr struct {
	BaseNode
	Expr   Expr        // scalar expression
	Output *JsonOutput // RETURNING clause, if specified
}

// JsonSerializeExpr represents untransformed JSON_SERIALIZE() function
// Ported from postgres/src/include/nodes/parsenodes.h:1908-1914
type JsonSerializeExpr struct {
	BaseNode
	Expr   *JsonValueExpr // json value expression
	Output *JsonOutput    // RETURNING clause, if specified
}

// JsonObjectConstructor represents untransformed JSON_OBJECT() constructor
// Ported from postgres/src/include/nodes/parsenodes.h:1920-1928
type JsonObjectConstructor struct {
	BaseNode
	Exprs        *NodeList   // list of JsonKeyValue pairs
	Output       *JsonOutput // RETURNING clause, if specified
	AbsentOnNull bool        // skip NULL values?
	Unique       bool        // check key uniqueness?
}

// JsonArrayConstructor represents untransformed JSON_ARRAY(element,...) constructor
// Ported from postgres/src/include/nodes/parsenodes.h:1934-1941
type JsonArrayConstructor struct {
	BaseNode
	Exprs        *NodeList   // list of JsonValueExpr elements
	Output       *JsonOutput // RETURNING clause, if specified
	AbsentOnNull bool        // skip NULL elements?
}

// JsonArrayQueryConstructor represents untransformed JSON_ARRAY(subquery) constructor
// Ported from postgres/src/include/nodes/parsenodes.h:1947-1955
type JsonArrayQueryConstructor struct {
	BaseNode
	Query        Node        // subquery
	Output       *JsonOutput // RETURNING clause, if specified
	Format       *JsonFormat // FORMAT clause for subquery, if specified
	AbsentOnNull bool        // skip NULL elements?
}

// JsonAggConstructor represents common fields of JSON_ARRAYAGG() and JSON_OBJECTAGG()
// Ported from postgres/src/include/nodes/parsenodes.h:1962-1970
type JsonAggConstructor struct {
	BaseNode
	Output    *JsonOutput // RETURNING clause, if any
	AggFilter Node        // FILTER clause, if any
	AggOrder  *NodeList   // ORDER BY clause, if any
	Over      *WindowDef  // OVER clause, if any
}

// JsonObjectAgg represents untransformed JSON_OBJECTAGG()
// Ported from postgres/src/include/nodes/parsenodes.h:1976-1983
type JsonObjectAgg struct {
	BaseNode
	Constructor  *JsonAggConstructor // common fields
	Arg          *JsonKeyValue       // object key-value pair
	AbsentOnNull bool                // skip NULL values?
	Unique       bool                // check key uniqueness?
}

// JsonArrayAgg represents untransformed JSON_ARRAYAGG()
// Ported from postgres/src/include/nodes/parsenodes.h:1989-1995
type JsonArrayAgg struct {
	BaseNode
	Constructor  *JsonAggConstructor // common fields
	Arg          *JsonValueExpr      // array element expression
	AbsentOnNull bool                // skip NULL elements?
}

// Implement Node interface for all JSON types

func (n *JsonFormat) node() { _ = "STUB: not implemented"; return }

func (n *JsonFormat) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of JsonFormat
func (n *JsonFormat) SqlString() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of JsonReturning
func (n *JsonReturning) SqlString() string {
	_ = "STUB: not implemented"
	// For now, just return the type name representation
	// Full implementation would need access to type system for proper type names
	return ""
}

// placeholder - in a real implementation this would resolve n.Typid to a type name

func (n *JsonValueExpr) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of the JsonValueExpr.
func (n *JsonValueExpr) SqlString() string { _ = "STUB: not implemented"; return "" }

func (n *JsonValueExpr) IsExpr() bool           { _ = "STUB: not implemented"; return false }
func (n *JsonValueExpr) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (n *JsonBehavior) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of JsonBehavior
func (n *JsonBehavior) SqlString() string { _ = "STUB: not implemented"; return "" }

func (n *JsonOutput) String() string { _ = "STUB: not implemented"; return "" }

func (n *JsonArgument) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of JsonArgument
func (n *JsonArgument) SqlString() string { _ = "STUB: not implemented"; return "" }

func (n *JsonFuncExpr) String() string { _ = "STUB: not implemented"; return "" }

func (n *JsonFuncExpr) IsExpr() bool           { _ = "STUB: not implemented"; return false }
func (n *JsonFuncExpr) ExpressionType() string { _ = "STUB: not implemented"; return "" }

// appendJsonReturning writes ` RETURNING <type> [FORMAT JSON [ENCODING X]]`
// when an output clause is present. The FORMAT clause is carried on
// Output.Returning.Format and applies to the returning value, distinct from
// any FORMAT clause on the input expression.
func appendJsonReturning(b *strings.Builder, output *JsonOutput) { _ = "STUB: not implemented"; return }

// appendJsonWrapperAndQuotes writes the [WITH/WITHOUT ARRAY WRAPPER]
// and [KEEP/OMIT QUOTES] clauses to the builder. Shared by JsonFuncExpr
// (for JSON_QUERY) and JsonTableColumn (for JTC_REGULAR / JTC_FORMATTED).
func appendJsonWrapperAndQuotes(b *strings.Builder, wrapper JsonWrapper, quotes JsonQuotes) {
	_ = "STUB: not implemented"
	return
}

// SqlString returns the SQL representation of JsonFuncExpr
func (n *JsonFuncExpr) SqlString() string { _ = "STUB: not implemented"; return "" }

// WRAPPER and QUOTES are JSON_QUERY-only.

func (n *JsonTablePathSpec) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of the JsonTablePathSpec.
// The optional `AS <name>` suffix is required for PG to detect duplicate
// column/path-name conflicts inside JSON_TABLE.
func (n *JsonTablePathSpec) SqlString() string { _ = "STUB: not implemented"; return "" }

func (n *JsonTable) String() string {
	_ = "STUB: not implemented"

	// SqlString returns the SQL representation of the JsonTable.
	return ""
}

func (n *JsonTable) SqlString() string { _ = "STUB: not implemented"; return "" }

// Context item expression (usually a JSON document)

// Path specification

// PASSING clause

// COLUMNS clause

// ON ERROR clause

func (n *JsonTableColumn) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of the JsonTableColumn.
func (n *JsonTableColumn) SqlString() string { _ = "STUB: not implemented"; return "" }

// Column name (only for non-NESTED columns)

// Add PATH clause if specified

// Add PATH clause if specified

// Add FORMAT clause if specified

// Add PATH clause if specified

// The `AS <name>` for a NESTED path is stored on the column's Name
// field (the path's own JsonTablePathSpec.Name is left empty by the
// grammar for NESTED paths). Emit it here so PG can validate
// column/path-name uniqueness.

// Add ON EMPTY clause if specified

// Add ON ERROR clause if specified

func (n *JsonKeyValue) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of JsonKeyValue
func (n *JsonKeyValue) SqlString() string { _ = "STUB: not implemented"; return "" }

func (n *JsonParseExpr) node()          { _ = "STUB: not implemented"; return }
func (n *JsonParseExpr) String() string { _ = "STUB: not implemented"; return "" }

func (n *JsonParseExpr) IsExpr() bool           { _ = "STUB: not implemented"; return false }
func (n *JsonParseExpr) ExpressionType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of JsonParseExpr
func (n *JsonParseExpr) SqlString() string { _ = "STUB: not implemented"; return "" }

func (n *JsonScalarExpr) node()          { _ = "STUB: not implemented"; return }
func (n *JsonScalarExpr) String() string { _ = "STUB: not implemented"; return "" }

func (n *JsonScalarExpr) IsExpr() bool           { _ = "STUB: not implemented"; return false }
func (n *JsonScalarExpr) ExpressionType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of JsonScalarExpr
func (n *JsonScalarExpr) SqlString() string { _ = "STUB: not implemented"; return "" }

func (n *JsonSerializeExpr) node()          { _ = "STUB: not implemented"; return }
func (n *JsonSerializeExpr) String() string { _ = "STUB: not implemented"; return "" }

func (n *JsonSerializeExpr) IsExpr() bool           { _ = "STUB: not implemented"; return false }
func (n *JsonSerializeExpr) ExpressionType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of JsonSerializeExpr
func (n *JsonSerializeExpr) SqlString() string { _ = "STUB: not implemented"; return "" }

func (n *JsonObjectConstructor) node()          { _ = "STUB: not implemented"; return }
func (n *JsonObjectConstructor) String() string { _ = "STUB: not implemented"; return "" }

func (n *JsonObjectConstructor) IsExpr() bool           { _ = "STUB: not implemented"; return false }
func (n *JsonObjectConstructor) ExpressionType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of JsonObjectConstructor
func (n *JsonObjectConstructor) SqlString() string { _ = "STUB: not implemented"; return "" }

// JSON_OBJECT defaults to NULL ON NULL; emit only the non-default.

func (n *JsonArrayConstructor) node()          { _ = "STUB: not implemented"; return }
func (n *JsonArrayConstructor) String() string { _ = "STUB: not implemented"; return "" }

func (n *JsonArrayConstructor) IsExpr() bool           { _ = "STUB: not implemented"; return false }
func (n *JsonArrayConstructor) ExpressionType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of JsonArrayConstructor
func (n *JsonArrayConstructor) SqlString() string { _ = "STUB: not implemented"; return "" }

// JSON_ARRAY defaults to ABSENT ON NULL; emit only the non-default.

func (n *JsonArrayQueryConstructor) node()          { _ = "STUB: not implemented"; return }
func (n *JsonArrayQueryConstructor) String() string { _ = "STUB: not implemented"; return "" }

func (n *JsonArrayQueryConstructor) IsExpr() bool           { _ = "STUB: not implemented"; return false }
func (n *JsonArrayQueryConstructor) ExpressionType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of JsonArrayQueryConstructor
func (n *JsonArrayQueryConstructor) SqlString() string { _ = "STUB: not implemented"; return "" }

func (n *JsonAggConstructor) node()          { _ = "STUB: not implemented"; return }
func (n *JsonAggConstructor) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of filter and over clauses
func (n *JsonAggConstructor) SqlString() string { _ = "STUB: not implemented"; return "" }

// Add FILTER clause if present

// Add OVER clause if present

func (n *JsonObjectAgg) node()          { _ = "STUB: not implemented"; return }
func (n *JsonObjectAgg) String() string { _ = "STUB: not implemented"; return "" }

func (n *JsonObjectAgg) IsExpr() bool           { _ = "STUB: not implemented"; return false }
func (n *JsonObjectAgg) ExpressionType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of JsonObjectAgg
func (n *JsonObjectAgg) SqlString() string { _ = "STUB: not implemented"; return "" }

// JSON_OBJECTAGG defaults to NULL ON NULL; emit only the non-default.

// Add RETURNING clause if present

// Add FILTER and OVER clauses if present

func (n *JsonArrayAgg) node()          { _ = "STUB: not implemented"; return }
func (n *JsonArrayAgg) String() string { _ = "STUB: not implemented"; return "" }

func (n *JsonArrayAgg) IsExpr() bool           { _ = "STUB: not implemented"; return false }
func (n *JsonArrayAgg) ExpressionType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of JsonArrayAgg
func (n *JsonArrayAgg) SqlString() string { _ = "STUB: not implemented"; return "" }

// ORDER BY (lives on Constructor.AggOrder, between the arg and the
// ABSENT/RETURNING clauses per the json_aggregate_func grammar).

// JSON_ARRAYAGG defaults to ABSENT ON NULL; emit only the non-default.

// Add RETURNING clause if present

// Add FILTER and OVER clauses if present

// Constructor functions

// NewJsonFormat creates a new JsonFormat node
func NewJsonFormat(formatType JsonFormatType, encoding JsonEncoding, location int) *JsonFormat {
	_ = "STUB: not implemented"
	return nil
}

// NewJsonReturning creates a new JsonReturning node
func NewJsonReturning(format *JsonFormat, typid Oid, typmod int32) *JsonReturning {
	_ = "STUB: not implemented"
	return nil
}

// NewJsonValueExpr creates a new JsonValueExpr node
func NewJsonValueExpr(rawExpr Node, format *JsonFormat) *JsonValueExpr {
	_ = "STUB: not implemented"
	return nil
}

// NewJsonBehavior creates a new JsonBehavior node
func NewJsonBehavior(btype JsonBehaviorType, expr Node, location int) *JsonBehavior {
	_ = "STUB: not implemented"
	return nil
}

// NewJsonOutput creates a new JsonOutput node
func NewJsonOutput(typeName *TypeName, returning *JsonReturning) *JsonOutput {
	_ = "STUB: not implemented"
	return nil
}

// NewJsonArgument creates a new JsonArgument node
func NewJsonArgument(val *JsonValueExpr, name string) *JsonArgument {
	_ = "STUB: not implemented"
	return nil
}

// NewJsonFuncExpr creates a new JsonFuncExpr node
func NewJsonFuncExpr(op JsonExprOp, contextItem *JsonValueExpr, pathspec Node) *JsonFuncExpr {
	_ = "STUB: not implemented"
	return nil
}

// NewJsonTablePathSpec creates a new JsonTablePathSpec node
func NewJsonTablePathSpec(str Node, name string, location int) *JsonTablePathSpec {
	_ = "STUB: not implemented"
	return nil
}

// NewJsonTable creates a new JsonTable node
func NewJsonTable(contextItem *JsonValueExpr, pathspec *JsonTablePathSpec) *JsonTable {
	_ = "STUB: not implemented"
	return nil
}

// NewJsonTableColumn creates a new JsonTableColumn node
func NewJsonTableColumn(coltype JsonTableColumnType, name string) *JsonTableColumn {
	_ = "STUB: not implemented"
	return nil
}

// NewJsonKeyValue creates a new JsonKeyValue node
func NewJsonKeyValue(key Expr, value *JsonValueExpr) *JsonKeyValue {
	_ = "STUB: not implemented"
	return nil
}

// NewJsonParseExpr creates a new JsonParseExpr node
func NewJsonParseExpr(expr *JsonValueExpr, uniqueKeys bool) *JsonParseExpr {
	_ = "STUB: not implemented"
	return nil
}

// NewJsonScalarExpr creates a new JsonScalarExpr node
func NewJsonScalarExpr(expr Expr) *JsonScalarExpr { _ = "STUB: not implemented"; return nil }

// NewJsonSerializeExpr creates a new JsonSerializeExpr node
func NewJsonSerializeExpr(expr *JsonValueExpr) *JsonSerializeExpr {
	_ = "STUB: not implemented"
	return nil
}

// NewJsonObjectConstructor creates a new JsonObjectConstructor node
func NewJsonObjectConstructor(exprs *NodeList, absentOnNull bool, unique bool) *JsonObjectConstructor {
	_ = "STUB: not implemented"
	return nil
}

// NewJsonArrayConstructor creates a new JsonArrayConstructor node
func NewJsonArrayConstructor(exprs *NodeList, absentOnNull bool) *JsonArrayConstructor {
	_ = "STUB: not implemented"
	return nil
}

// NewJsonArrayQueryConstructor creates a new JsonArrayQueryConstructor node
func NewJsonArrayQueryConstructor(query Node, absentOnNull bool) *JsonArrayQueryConstructor {
	_ = "STUB: not implemented"
	return nil
}

// NewJsonAggConstructor creates a new JsonAggConstructor node
func NewJsonAggConstructor(output *JsonOutput) *JsonAggConstructor {
	_ = "STUB: not implemented"
	return nil
}

// NewJsonObjectAgg creates a new JsonObjectAgg node
func NewJsonObjectAgg(constructor *JsonAggConstructor, arg *JsonKeyValue, absentOnNull bool, unique bool) *JsonObjectAgg {
	_ = "STUB: not implemented"
	return nil
}

// NewJsonArrayAgg creates a new JsonArrayAgg node
func NewJsonArrayAgg(constructor *JsonAggConstructor, arg *JsonValueExpr, absentOnNull bool) *JsonArrayAgg {
	_ = "STUB: not implemented"
	return nil
}

// ==============================================================================
// PHASE 1G: JSON PRIMITIVE EXPRESSIONS - Missing JSON expression nodes
// Ported from postgres/src/include/nodes/primnodes.h
// ==============================================================================

// JsonConstructorType represents JSON constructor types
// Ported from postgres/src/include/nodes/primnodes.h:1688-1697
type JsonConstructorType int

const (
	JSCTOR_JSON_OBJECT    JsonConstructorType = iota + 1 // JSON_OBJECT constructor
	JSCTOR_JSON_ARRAY                                    // JSON_ARRAY constructor
	JSCTOR_JSON_OBJECTAGG                                // JSON_OBJECTAGG constructor
	JSCTOR_JSON_ARRAYAGG                                 // JSON_ARRAYAGG constructor
	JSCTOR_JSON_PARSE                                    // JSON_PARSE constructor
	JSCTOR_JSON_SCALAR                                   // JSON_SCALAR constructor
	JSCTOR_JSON_SERIALIZE                                // JSON_SERIALIZE constructor
)

// JsonConstructorExpr represents a wrapper over FuncExpr/Aggref/WindowFunc for SQL/JSON constructors
// Ported from postgres/src/include/nodes/primnodes.h:1703-1714
type JsonConstructorExpr struct {
	BaseExpr
	Type         JsonConstructorType // constructor type
	Args         *NodeList           // arguments list
	Func         Expr                // underlying json[b]_xxx() function call
	Coercion     Expr                // coercion to RETURNING type
	Returning    *JsonReturning      // RETURNING clause
	AbsentOnNull bool                // ABSENT ON NULL?
	Unique       bool                // WITH UNIQUE KEYS? (JSON_OBJECT[AGG] only)
}

func (j *JsonConstructorExpr) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (j *JsonConstructorExpr) String() string { _ = "STUB: not implemented"; return "" }

// NewJsonConstructorExpr creates a new JsonConstructorExpr node
func NewJsonConstructorExpr(constructorType JsonConstructorType, args *NodeList, function, coercion Expr, returning *JsonReturning, absentOnNull, unique bool, location int) *JsonConstructorExpr {
	_ = "STUB: not implemented"
	return nil
}

// JsonIsPredicate represents an IS JSON predicate
// Ported from postgres/src/include/nodes/primnodes.h:1732-1740
type JsonIsPredicate struct {
	BaseNode
	Expr       Node          // subject expression
	Format     *JsonFormat   // FORMAT clause, if specified
	ItemType   JsonValueType // JSON item type
	UniqueKeys bool          // check key uniqueness?
}

func (j *JsonIsPredicate) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of JsonIsPredicate
func (j *JsonIsPredicate) SqlString() string { _ = "STUB: not implemented"; return "" }

// For JS_TYPE_ANY, we don't add any suffix (just "IS JSON")

// NewJsonIsPredicate creates a new JsonIsPredicate node
func NewJsonIsPredicate(expr Node, format *JsonFormat, itemType JsonValueType, uniqueKeys bool, location int) *JsonIsPredicate {
	_ = "STUB: not implemented"
	return nil
}

// JsonExpr represents transformed representation of JSON_VALUE(), JSON_QUERY(), and JSON_EXISTS()
// Ported from postgres/src/include/nodes/primnodes.h:1813-1860
type JsonExpr struct {
	BaseExpr
	Op              JsonExprOp     // JSON expression operation type
	ColumnName      string         // JSON_TABLE() column name or empty if not for JSON_TABLE()
	FormattedExpr   Node           // jsonb-valued expression to query
	Format          *JsonFormat    // Format of the above expression needed by ruleutils.c
	PathSpec        Node           // jsonpath-valued expression containing the query pattern
	Returning       *JsonReturning // Expected type/format of the output
	PassingNames    []string       // PASSING argument names
	PassingValues   *NodeList      // PASSING argument values
	OnEmpty         *JsonBehavior  // User-specified or default ON EMPTY behavior
	OnError         *JsonBehavior  // User-specified or default ON ERROR behavior
	UseIOCoercion   bool           // Information about converting the result to RETURNING type
	UseJsonCoercion bool           // Additional conversion information
	Wrapper         JsonWrapper    // WRAPPER specification for JSON_QUERY
	OmitQuotes      bool           // KEEP or OMIT QUOTES for singleton scalars returned by JSON_QUERY()
	Collation       Oid            // JsonExpr's collation
}

func (j *JsonExpr) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (j *JsonExpr) String() string { _ = "STUB: not implemented"; return "" }

// NewJsonExpr creates a new JsonExpr node
func NewJsonExpr(op JsonExprOp, columnName string, formattedExpr Node, format *JsonFormat, pathSpec Node, returning *JsonReturning, passingNames []string, passingValues *NodeList, onEmpty, onError *JsonBehavior, useIOCoercion, useJsonCoercion bool, wrapper JsonWrapper, omitQuotes bool, collation Oid, location int) *JsonExpr {
	_ = "STUB: not implemented"
	return nil
}

// JsonTablePath represents a JSON path expression to be computed as part of evaluating a JSON_TABLE plan node
// Ported from postgres/src/include/nodes/primnodes.h:1867-1873
type JsonTablePath struct {
	BaseNode
	Value *Const // path value
	Name  string // path name
}

func (j *JsonTablePath) String() string { _ = "STUB: not implemented"; return "" }

// NewJsonTablePath creates a new JsonTablePath node
func NewJsonTablePath(value *Const, name string, location int) *JsonTablePath {
	_ = "STUB: not implemented"
	return nil
}

// JsonTablePlan represents abstract base type for different types of JSON_TABLE "plans"
// Ported from postgres/src/include/nodes/primnodes.h:1882-1887
type JsonTablePlan struct {
	BaseNode
}

func (j *JsonTablePlan) String() string { _ = "STUB: not implemented"; return "" }

// NewJsonTablePlan creates a new JsonTablePlan node
func NewJsonTablePlan(location int) *JsonTablePlan { _ = "STUB: not implemented"; return nil }

// JsonTablePathScan represents a JSON_TABLE plan to evaluate a JSON path expression and NESTED paths
// Ported from postgres/src/include/nodes/primnodes.h:1893-1916
type JsonTablePathScan struct {
	BaseNode
	Path         *JsonTablePath // JSON path to evaluate
	ErrorOnError bool           // ERROR/EMPTY ON ERROR behavior; only significant in the plan for the top-level path
	Child        *JsonTablePlan // Plan(s) for nested columns, if any
	ColMin       int            // 0-based index in TableFunc.colvalexprs of the 1st column covered by this plan
	ColMax       int            // 0-based index in TableFunc.colvalexprs of the last column covered by this plan
}

func (j *JsonTablePathScan) String() string { _ = "STUB: not implemented"; return "" }

// NewJsonTablePathScan creates a new JsonTablePathScan node
func NewJsonTablePathScan(path *JsonTablePath, errorOnError bool, child *JsonTablePlan, colMin, colMax, location int) *JsonTablePathScan {
	_ = "STUB: not implemented"
	return nil
}

// JsonTableSiblingJoin represents a plan to join rows of sibling NESTED COLUMNS clauses in the same parent COLUMNS clause
// Ported from postgres/src/include/nodes/primnodes.h:1923-1929
type JsonTableSiblingJoin struct {
	BaseNode
	Lplan *JsonTablePlan // left plan
	Rplan *JsonTablePlan // right plan
}

func (j *JsonTableSiblingJoin) String() string { _ = "STUB: not implemented"; return "" }

// NewJsonTableSiblingJoin creates a new JsonTableSiblingJoin node
func NewJsonTableSiblingJoin(lplan, rplan *JsonTablePlan, location int) *JsonTableSiblingJoin {
	_ = "STUB: not implemented"
	return nil
}
