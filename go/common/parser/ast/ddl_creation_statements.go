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

// ==============================================================================
// DDL CREATION STATEMENTS - PostgreSQL parsenodes.h DDL creation implementation
// Ported from postgres/src/include/nodes/parsenodes.h
// ==============================================================================

// CoercionContext represents the context for type coercion operations
// Ported from postgres/src/include/nodes/primnodes.h:712-718
type CoercionContext int

const (
	COERCION_IMPLICIT   CoercionContext = iota // coercion in context of expression
	COERCION_ASSIGNMENT                        // coercion in context of assignment
	COERCION_PLPGSQL                           // if no assignment cast, use CoerceViaIO
	COERCION_EXPLICIT                          // explicit cast operation
)

// String returns string representation of CoercionContext
func (cc CoercionContext) String() string { _ = "STUB: not implemented"; return "" }

// FunctionParameterMode represents parameter passing modes for function parameters
// Ported from postgres/src/include/nodes/parsenodes.h:3439-3449
type FunctionParameterMode int

const (
	FUNC_PARAM_IN       FunctionParameterMode = iota // input only
	FUNC_PARAM_OUT                                   // output only
	FUNC_PARAM_INOUT                                 // both
	FUNC_PARAM_VARIADIC                              // variadic (always input)
	FUNC_PARAM_TABLE                                 // table function output column
	FUNC_PARAM_DEFAULT                               // default; effectively same as IN
)

// String returns string representation of FunctionParameterMode
func (fpm FunctionParameterMode) String() string { _ = "STUB: not implemented"; return "" }

// FetchDirection represents the direction for cursor fetch operations
// Ported from postgres/src/include/nodes/parsenodes.h:3316-3324
type FetchDirection int

const (
	FETCH_FORWARD  FetchDirection = iota // forward direction, howMany is row count
	FETCH_BACKWARD                       // backward direction, howMany is row count
	FETCH_ABSOLUTE                       // absolute position, howMany is position
	FETCH_RELATIVE                       // relative position, howMany is offset
)

// FETCH_ALL constant for fetching all rows
// Ported from postgres/src/include/nodes/parsenodes.h:3326
const FETCH_ALL = 9223372036854775807 // LONG_MAX

// Cursor option constants
// Ported from postgres/src/include/nodes/parsenodes.h:3275-3291
const (
	CURSOR_OPT_BINARY       = 0x0001 // BINARY
	CURSOR_OPT_SCROLL       = 0x0002 // SCROLL explicitly given
	CURSOR_OPT_NO_SCROLL    = 0x0004 // NO SCROLL explicitly given
	CURSOR_OPT_INSENSITIVE  = 0x0008 // INSENSITIVE
	CURSOR_OPT_ASENSITIVE   = 0x0010 // ASENSITIVE
	CURSOR_OPT_HOLD         = 0x0020 // WITH HOLD
	CURSOR_OPT_FAST_PLAN    = 0x0100 // prefer fast-start plan
	CURSOR_OPT_GENERIC_PLAN = 0x0200 // force use of generic plan
	CURSOR_OPT_CUSTOM_PLAN  = 0x0400 // force use of custom plan
	CURSOR_OPT_PARALLEL_OK  = 0x0800 // parallel mode OK
)

// String returns string representation of FetchDirection
func (fd FetchDirection) String() string { _ = "STUB: not implemented"; return "" }

// FunctionParameter represents a function parameter specification
// Ported from postgres/src/include/nodes/parsenodes.h:3451-3458
type FunctionParameter struct {
	BaseNode
	Name    string                // parameter name, or empty string if not given
	ArgType *TypeName             // type name for parameter type
	Mode    FunctionParameterMode // IN/OUT/INOUT/VARIADIC/TABLE/DEFAULT
	DefExpr Node                  // raw default expr, or nil if not given
}

// String returns string representation of FunctionParameter
func (fp *FunctionParameter) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of FunctionParameter
func (fp *FunctionParameter) SqlString() string {
	_ = "STUB: not implemented"

	// Parameter mode
	return ""
}

// IN is default, don't need to specify

// Parameter name

// Parameter type

// Default value

// NewFunctionParameter creates a new FunctionParameter node
func NewFunctionParameter(name string, argType *TypeName, mode FunctionParameterMode, defExpr Node) *FunctionParameter {
	_ = "STUB: not implemented"
	return nil
}

// CreateFunctionStmt represents a CREATE FUNCTION statement
// Ported from postgres/src/include/nodes/parsenodes.h:3427-3437
type CreateFunctionStmt struct {
	IsProcedure bool      // true for CREATE PROCEDURE
	Replace     bool      // true for CREATE OR REPLACE
	FuncName    *NodeList // qualified name of function to create
	Parameters  *NodeList // list of function parameters (includes RETURNS TABLE columns with FUNC_PARAM_TABLE mode)
	ReturnType  *TypeName // the return type
	Options     *NodeList // list of definition elements
	SQLBody     Node      // SQL body for SQL functions
}

// node implements the Node interface
func (cfs *CreateFunctionStmt) node() {
	_ = "STUB: not implemented"

	// stmt implements the Stmt interface
	return
}

func (cfs *CreateFunctionStmt) stmt() {
	_ = "STUB: not implemented"

	// Location returns the statement's source location (dummy implementation)
	return
}

func (cfs *CreateFunctionStmt) Location() int {
	_ = "STUB: not implemented"
	// TODO: Implement proper location tracking
	return 0
}

// SetLocation is a no-op; this node does not track source location yet.
func (cfs *CreateFunctionStmt) SetLocation(int) {
	_ = "STUB: not implemented"

	// NodeTag returns the node's type tag
	return
}

func (cfs *CreateFunctionStmt) NodeTag() NodeTag { _ = "STUB: not implemented"; return *new(NodeTag) }

// StatementType returns the statement type for this node
func (cfs *CreateFunctionStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns SQL representation of the CREATE FUNCTION statement
func (cfs *CreateFunctionStmt) SqlString() string {
	_ = "STUB: not implemented"

	// CREATE [OR REPLACE]
	return ""
}

// FUNCTION or PROCEDURE

// Function name

// Parameters - separate regular params from RETURNS TABLE columns

// This is a RETURNS TABLE column

// RETURNS type (for functions, not procedures)

// RETURNS TABLE (col1 type1, col2 type2, ...)

// Function options - process in original order

// Use DollarQuoteString to handle nested dollar quotes properly

// Single element: the function body, dollar-quoted.

// Two elements: a C function's 'objfile', 'symbol' — each a
// plain string literal. Dropping the second is wrong.

// All other options (SECURITY DEFINER, STABLE, STRICT, etc.)

// SQL body

// Check if this is a compound statement (BEGIN ATOMIC ... END)
// Compound statements are stored as a NodeList containing another NodeList

// Regular SQL body

// Regular SQL body

// String returns string representation of CreateFunctionStmt
func (cfs *CreateFunctionStmt) String() string { _ = "STUB: not implemented"; return "" }

// Parameters

// Return type (only for functions, not procedures)

// NewCreateFunctionStmt creates a new CreateFunctionStmt node
func NewCreateFunctionStmt(isProcedure, replace bool, funcName *NodeList, parameters *NodeList, returnType *TypeName, options *NodeList, sqlBody Node) *CreateFunctionStmt {
	_ = "STUB: not implemented"
	return nil
}

// MergeTableFuncParameters merges regular function parameters with RETURNS TABLE columns.
// The table columns are appended to the parameter list with FUNC_PARAM_TABLE mode.
// This matches PostgreSQL's mergeTableFuncParameters function.
func MergeTableFuncParameters(params *NodeList, tableCols *NodeList) *NodeList {
	_ = "STUB: not implemented"
	return nil
}

// TableFuncTypeName creates a TypeName representing a RECORD type for RETURNS TABLE.
// This matches PostgreSQL's TableFuncTypeName function.
func TableFuncTypeName(tableCols *NodeList) *TypeName {
	_ = "STUB: not implemented"
	// Create a RECORD type - PostgreSQL uses this for table functions
	return nil
}

// CreateSeqStmt represents a CREATE SEQUENCE statement
// Ported from postgres/src/include/nodes/parsenodes.h:3117-3125
type CreateSeqStmt struct {
	BaseNode
	Sequence    *RangeVar // the sequence to create
	Options     *NodeList // list of options (DefElem nodes)
	OwnerID     Oid       // ID of owner, or 0 for default (InvalidOid)
	ForIdentity bool      // true if for IDENTITY column
	IfNotExists bool      // true for IF NOT EXISTS
}

// node implements the Node interface
func (css *CreateSeqStmt) node() {
	_ = "STUB: not implemented"

	// stmt implements the Stmt interface
	return
}

func (css *CreateSeqStmt) stmt() {
	_ = "STUB: not implemented"

	// StatementType returns the type of statement
	return
}

func (css *CreateSeqStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

// String returns string representation of CreateSeqStmt
func (css *CreateSeqStmt) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of CreateSeqStmt
func (css *CreateSeqStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// Add sequence options if present

// NewCreateSeqStmt creates a new CreateSeqStmt node
func NewCreateSeqStmt(sequence *RangeVar, options *NodeList, ownerID Oid, forIdentity, ifNotExists bool) *CreateSeqStmt {
	_ = "STUB: not implemented"
	return nil
}

// AlterSeqStmt represents an ALTER SEQUENCE statement
// Ported from postgres/src/include/nodes/parsenodes.h:3470-3477
type AlterSeqStmt struct {
	BaseNode
	Sequence    *RangeVar // the sequence to alter
	Options     *NodeList // list of DefElem options
	ForIdentity bool      // for IDENTITY column
	MissingOk   bool      // skip error if sequence doesn't exist
}

// NewAlterSeqStmt creates a new AlterSeqStmt node
func NewAlterSeqStmt(sequence *RangeVar, options *NodeList, forIdentity, missingOk bool) *AlterSeqStmt {
	_ = "STUB: not implemented"
	return nil
}

// StatementType returns the statement type
func (ass *AlterSeqStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

// String returns string representation of AlterSeqStmt
func (ass *AlterSeqStmt) String() string { _ = "STUB: not implemented"; return "" }

// formatSeqOption formats a sequence option DefElem for SQL output
func formatSeqOption(opt *DefElem) string { _ = "STUB: not implemented"; return "" }

// Special handling for sequence-specific options

// Fall back to default formatting

// SqlString returns the SQL representation of AlterSeqStmt
func (ass *AlterSeqStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// Add sequence options if present

// CreateOpClassItem represents an item in a CREATE OPERATOR CLASS statement
// Ported from postgres/src/include/nodes/parsenodes.h:3184-3195
type CreateOpClassItem struct {
	BaseNode
	ItemType    OpClassItemType // OPCLASS_ITEM_OPERATOR, OPCLASS_ITEM_FUNCTION, or OPCLASS_ITEM_STORAGETYPE
	Name        *ObjectWithArgs // operator or function name and args
	Number      int             // strategy num or support proc num
	OrderFamily *NodeList       // only used for ordering operators
	ClassArgs   *NodeList       // amproclefttype/amprocrighttype or amoplefttype/amoprighttype
	StoredType  *TypeName       // datatype stored in index (for storage type items)
}

// String returns string representation of CreateOpClassItem
func (oci *CreateOpClassItem) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of CreateOpClassItem
func (oci *CreateOpClassItem) SqlString() string { _ = "STUB: not implemented"; return "" }

// For functions, add ClassArgs before the function name

// Preserve original PostgreSQL type names

// Fix spacing and type names for operators

// Fix type normalization in operator arguments (order matters!)

// Add a space between the operator name and the argument list,
// e.g. "=(int4, int4)" -> "= (int4, int4)". If there is no operator
// name (DROP OPERATOR), nameStr starts with "(" and we must not
// emit a leading space — that produces a double space when joined.

// Fix type normalization in function arguments (order matters!)

// Add ClassArgs if present for non-function items (like operators)

// Preserve original PostgreSQL type names

// Add FOR ORDER BY clause if OrderFamily is present

// Handle String nodes specially to avoid quotes for operator family names

func (oci *CreateOpClassItem) StatementType() string { _ = "STUB: not implemented"; return "" }

// NewCreateOpClassItem creates a new CreateOpClassItem node
func NewCreateOpClassItem(itemType OpClassItemType, name *ObjectWithArgs, number int, orderFamily *NodeList, classArgs *NodeList, storedType *TypeName) *CreateOpClassItem {
	_ = "STUB: not implemented"
	return nil
}

// NewOpClassItemOperator creates a new CreateOpClassItem for operators
// Two different signatures to match grammar usage:
// 1. OPERATOR Iconst any_operator opclass_purpose opt_recheck
// 2. OPERATOR Iconst operator_with_argtypes opclass_purpose opt_recheck
func NewOpClassItemOperator(number int, name *ObjectWithArgs, orderFamily *NodeList) *CreateOpClassItem {
	_ = "STUB: not implemented"
	return nil
}

// NewOpClassItemFunction creates a new CreateOpClassItem for functions
// Two different signatures to match grammar usage:
// 1. FUNCTION Iconst function_with_argtypes
// 2. FUNCTION Iconst '(' type_list ')' function_with_argtypes
func NewOpClassItemFunction(number int, name *ObjectWithArgs, classArgs *NodeList) *CreateOpClassItem {
	_ = "STUB: not implemented"
	return nil
}

// NewOpClassItemStorage creates a new CreateOpClassItem for storage type
func NewOpClassItemStorage(storedType *TypeName) *CreateOpClassItem {
	_ = "STUB: not implemented"
	return nil
}

// CreateOpClassStmt represents a CREATE OPERATOR CLASS statement
// Ported from postgres/src/include/nodes/parsenodes.h:3169-3178
type CreateOpClassStmt struct {
	BaseNode
	OpClassName  *NodeList // qualified name (list of String)
	OpFamilyName *NodeList // qualified name (list of String); nil if omitted
	AmName       string    // name of index AM opclass is for
	DataType     *TypeName // datatype of indexed column
	Items        *NodeList // list of CreateOpClassItem nodes
	IsDefault    bool      // should be marked as default for type?
}

// StatementType implements the Stmt interface
func (cocs *CreateOpClassStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

// NodeTag implements the Node interface
func (cocs *CreateOpClassStmt) NodeTag() NodeTag { _ = "STUB: not implemented"; return *new(NodeTag) }

// SqlString implements the Stmt interface
func (cocs *CreateOpClassStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// Preserve original case for types like 'uuid'

// Convert back to lowercase for certain common types

// Properly iterate through Items instead of hardcoding

// Fix type case for storage items too

// String returns string representation of CreateOpClassStmt
func (cocs *CreateOpClassStmt) String() string { _ = "STUB: not implemented"; return "" }

// NewCreateOpClassStmt creates a new CreateOpClassStmt node
func NewCreateOpClassStmt(opClassName, opFamilyName *NodeList, amName string, dataType *TypeName, items *NodeList, isDefault bool) *CreateOpClassStmt {
	_ = "STUB: not implemented"
	return nil
}

// CreateOpFamilyStmt represents a CREATE OPERATOR FAMILY statement
// Ported from postgres/src/include/nodes/parsenodes.h:3201-3206
type CreateOpFamilyStmt struct {
	BaseNode
	OpFamilyName *NodeList // qualified name (list of String)
	AmName       string    // name of index AM opfamily is for
}

// StatementType implements the Stmt interface
func (cofs *CreateOpFamilyStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

// NodeTag implements the Node interface
func (cofs *CreateOpFamilyStmt) NodeTag() NodeTag { _ = "STUB: not implemented"; return *new(NodeTag) }

// SqlString implements the Stmt interface
func (cofs *CreateOpFamilyStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// String returns string representation of CreateOpFamilyStmt
func (cofs *CreateOpFamilyStmt) String() string { _ = "STUB: not implemented"; return "" }

// NewCreateOpFamilyStmt creates a new CreateOpFamilyStmt node
func NewCreateOpFamilyStmt(opFamilyName *NodeList, amName string) *CreateOpFamilyStmt {
	_ = "STUB: not implemented"
	return nil
}

// CreateCastStmt represents a CREATE CAST statement
// Ported from postgres/src/include/nodes/parsenodes.h:4002-4010
type CreateCastStmt struct {
	SourceType *TypeName       // source data type
	TargetType *TypeName       // target data type
	Func       *ObjectWithArgs // conversion function, or nil
	Context    CoercionContext // coercion context
	Inout      bool            // true for INOUT cast
}

// String returns string representation of CreateCastStmt
func (ccs *CreateCastStmt) String() string { _ = "STUB: not implemented"; return "" }

func (ccs *CreateCastStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

func (ccs *CreateCastStmt) Location() int {
	_ = "STUB: not implemented"

	// SetLocation is a no-op; this node does not track source location yet.
	return 0
}

func (ccs *CreateCastStmt) SetLocation(int) { _ = "STUB: not implemented"; return }

func (ccs *CreateCastStmt) NodeTag() NodeTag { _ = "STUB: not implemented"; return *new(NodeTag) }

func (ccs *CreateCastStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// Add the context clause (AS IMPLICIT, AS ASSIGNMENT, or nothing for EXPLICIT)

// COERCION_EXPLICIT doesn't have an AS clause

// NewCreateCastStmt creates a new CreateCastStmt node
func NewCreateCastStmt(sourceType, targetType *TypeName, function *ObjectWithArgs, context CoercionContext, inout bool) *CreateCastStmt {
	_ = "STUB: not implemented"
	return nil
}

// CreateConversionStmt represents a CREATE CONVERSION statement
// Ported from postgres/src/include/nodes/parsenodes.h:3988-3996
type CreateConversionStmt struct {
	BaseNode
	ConversionName  *NodeList // name of the conversion
	ForEncodingName string    // source encoding name
	ToEncodingName  string    // destination encoding name
	FuncName        *NodeList // qualified conversion function name
	Def             bool      // true if this is a default conversion
}

// String returns string representation of CreateConversionStmt
func (ccs *CreateConversionStmt) String() string { _ = "STUB: not implemented"; return "" }

func (ccs *CreateConversionStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

func (ccs *CreateConversionStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// NewCreateConversionStmt creates a new CreateConversionStmt node
func NewCreateConversionStmt(conversionName *NodeList, forEncodingName, toEncodingName string, funcName *NodeList, def bool) *CreateConversionStmt {
	_ = "STUB: not implemented"
	return nil
}

// CreateTransformStmt represents a CREATE TRANSFORM statement
// Ported from postgres/src/include/nodes/parsenodes.h:4016-4024
type CreateTransformStmt struct {
	BaseNode
	Replace  bool            // true for CREATE OR REPLACE
	TypeName *TypeName       // type name
	Lang     string          // language name
	FromSql  *ObjectWithArgs // FROM SQL function
	ToSql    *ObjectWithArgs // TO SQL function
}

// String returns string representation of CreateTransformStmt
func (cts *CreateTransformStmt) String() string { _ = "STUB: not implemented"; return "" }

func (cts *CreateTransformStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

func (cts *CreateTransformStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// NewCreateTransformStmt creates a new CreateTransformStmt node
func NewCreateTransformStmt(replace bool, typeName *TypeName, lang string, fromSql, toSql Node) *CreateTransformStmt {
	_ = "STUB: not implemented"
	return nil
}

// DefineStmt represents a CREATE {AGGREGATE|OPERATOR|TYPE} statement
// Ported from postgres/src/include/nodes/parsenodes.h:3140-3150
type DefineStmt struct {
	BaseNode
	Kind        ObjectType // aggregate, operator, type
	OldStyle    bool       // hack to signal old CREATE AGG syntax
	DefNames    *NodeList  // qualified name (list of String)
	Args        *NodeList  // list of TypeName (if needed)
	Definition  *NodeList  // list of DefElem
	IfNotExists bool       // true for IF NOT EXISTS
	Replace     bool       // true for CREATE OR REPLACE
}

// node implements the Node interface
func (ds *DefineStmt) node() {
	_ = "STUB: not implemented"

	// stmt implements the Stmt interface
	return
}

func (ds *DefineStmt) stmt() {
	_ = "STUB: not implemented"

	// StatementType returns the statement type
	return
}

func (ds *DefineStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

// String returns string representation of DefineStmt
func (ds *DefineStmt) String() string { _ = "STUB: not implemented"; return "" }

// Add arguments for aggregates and operators

// Check for nil which represents * in aggregates like COUNT(*)

// Use the full SqlString to include VARIADIC, parameter names, etc.

// Add definition

// NewDefineStmt creates a new DefineStmt node
func NewDefineStmt(kind ObjectType, oldStyle bool, defNames *NodeList, args *NodeList, definition *NodeList, ifNotExists, replace bool) *DefineStmt {
	_ = "STUB: not implemented"
	return nil
}

// SqlString returns the SQL representation of DefineStmt
// operatorDefElemString renders a single CREATE OPERATOR definition element.
// The COMMUTATOR and NEGATOR options carry an operator name (a list of String
// parts), which must be emitted unquoted (`commutator = ===`), not as a string
// literal the way DefElem.SqlString would render a *NodeList. Other options
// (leftarg, procedure, ...) fall back to the default rendering.
func operatorDefElemString(d *DefElem) string { _ = "STUB: not implemented"; return "" }

// The final part is the operator symbol (unquoted); any leading parts
// are schema qualifiers (quoted as identifiers).

func (ds *DefineStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// Add object type

// Add IF NOT EXISTS if present

// Add name

// Add arguments for aggregates and operators

// Try the proper aggr_args structure first

// Args should be [argList, numDirectArgs] from aggr_args grammar

// Fallback for malformed structure

// Regular aggregate or COUNT(*)

// COUNT(*) case

// Ordered-set aggregate without direct args: (ORDER BY args)

// Hypothetical-set aggregate: (direct_args ORDER BY ordered_args)

// numDirectArgs == total arg count happens only when the
// last direct arg is VARIADIC: makeOrderedSetArgs drops the
// duplicate VARIADIC ordered arg and folds it into the direct
// list. Reconstruct it so this re-parses as an ordered-set
// aggregate rather than a plain one.

// Fallback to old logic for non-aggregate cases or malformed structures

// Check for nil which represents * in aggregates like COUNT(*)

// Add definition

// Check if this is a FROM clause (for COLLATION FROM syntax)

// Check if the first item is a NodeList (which indicates FROM syntax)

// This is FROM syntax - handle qualified names in FROM clause

// Normal definition with DefElem items

// DeclareCursorStmt represents a DECLARE cursor statement
// Ported from postgres/src/include/nodes/parsenodes.h:3293-3299
type DeclareCursorStmt struct {
	BaseNode
	PortalName string // name of the portal (cursor)
	Options    int    // bitmask of options
	Query      Node   // the query (raw parse tree)
}

// String returns string representation of DeclareCursorStmt
func (dcs *DeclareCursorStmt) String() string { _ = "STUB: not implemented"; return "" }

func (dcs *DeclareCursorStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of the DECLARE CURSOR statement
func (dcs *DeclareCursorStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// Add cursor options

// Add the query

// NewDeclareCursorStmt creates a new DeclareCursorStmt node
func NewDeclareCursorStmt(portalName string, options int, query Node) *DeclareCursorStmt {
	_ = "STUB: not implemented"
	return nil
}

// FetchStmt represents a FETCH statement (also MOVE)
// Ported from postgres/src/include/nodes/parsenodes.h:3328-3335
type FetchStmt struct {
	BaseNode
	Direction  FetchDirection // see FetchDirection enum
	HowMany    int64          // number of rows, or position argument
	PortalName string         // name of portal (cursor)
	IsMove     bool           // true if MOVE
}

// node implements the Node interface
func (fs *FetchStmt) node() {
	_ = "STUB: not implemented"

	// stmt implements the Stmt interface
	return
}

func (fs *FetchStmt) stmt() {
	_ = "STUB: not implemented"

	// String returns string representation of FetchStmt
	return
}

func (fs *FetchStmt) String() string { _ = "STUB: not implemented"; return "" }

func (fs *FetchStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of the FETCH/MOVE statement
func (fs *FetchStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// Handle direction and count

// Default case - just FETCH/MOVE cursor_name

// NewFetchStmt creates a new FetchStmt node
func NewFetchStmt(direction FetchDirection, howMany int64, portalName string, isMove bool) *FetchStmt {
	_ = "STUB: not implemented"
	return nil
}

// ClosePortalStmt represents a CLOSE statement
// Ported from postgres/src/include/nodes/parsenodes.h:3305-3310
type ClosePortalStmt struct {
	BaseNode
	PortalName string // name of the portal (cursor); nil means CLOSE ALL
}

// node implements the Node interface
func (cps *ClosePortalStmt) node() {
	_ = "STUB: not implemented"

	// stmt implements the Stmt interface
	return
}

func (cps *ClosePortalStmt) stmt() {
	_ = "STUB: not implemented"

	// String returns string representation of ClosePortalStmt
	return
}

func (cps *ClosePortalStmt) String() string { _ = "STUB: not implemented"; return "" }

func (cps *ClosePortalStmt) StatementType() string {
	_ = "STUB: not implemented"

	// SqlString returns the SQL representation of the CLOSE statement
	return ""
}

func (cps *ClosePortalStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// NewClosePortalStmt creates a new ClosePortalStmt node
func NewClosePortalStmt(portalName string) *ClosePortalStmt { _ = "STUB: not implemented"; return nil }

// CreateEnumStmt represents a CREATE TYPE ... AS ENUM statement
// Ported from postgres/src/include/nodes/parsenodes.h:3696-3701
type CreateEnumStmt struct {
	BaseNode
	TypeName *NodeList // qualified name (list of String)
	Vals     *NodeList // enum values (list of String)
}

// node implements the Node interface
func (ces *CreateEnumStmt) node() {
	_ = "STUB: not implemented"

	// stmt implements the Stmt interface
	return
}

func (ces *CreateEnumStmt) stmt() {
	_ = "STUB: not implemented"

	// StatementType returns the statement type
	return
}

func (ces *CreateEnumStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

// String returns string representation of CreateEnumStmt
func (ces *CreateEnumStmt) String() string { _ = "STUB: not implemented"; return "" }

// NewCreateEnumStmt creates a new CreateEnumStmt node
func NewCreateEnumStmt(typeName *NodeList, vals *NodeList) *CreateEnumStmt {
	_ = "STUB: not implemented"
	return nil
}

// SqlString returns the SQL representation of CreateEnumStmt
func (ces *CreateEnumStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// Add type name

// Add enum values

// CompositeTypeStmt represents a CREATE TYPE ... AS (...) statement
// Ported from postgres/src/include/nodes/parsenodes.h:3684
type CompositeTypeStmt struct {
	BaseNode
	Typevar    *RangeVar // the composite type name
	Coldeflist *NodeList // list of ColumnDef nodes
}

// node implements the Node interface
func (cts *CompositeTypeStmt) node() {
	_ = "STUB: not implemented"

	// stmt implements the Stmt interface
	return
}

func (cts *CompositeTypeStmt) stmt() {
	_ = "STUB: not implemented"

	// StatementType returns the statement type
	return
}

func (cts *CompositeTypeStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

// String returns string representation of CompositeTypeStmt
func (cts *CompositeTypeStmt) String() string { _ = "STUB: not implemented"; return "" }

// Simplified representation

// NewCompositeTypeStmt creates a new CompositeTypeStmt node
func NewCompositeTypeStmt(typevar *RangeVar, coldeflist *NodeList) *CompositeTypeStmt {
	_ = "STUB: not implemented"
	return nil
}

// SqlString returns the SQL representation of CompositeTypeStmt
func (cts *CompositeTypeStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// Add type name

// Add column definitions

// AlterEnumStmt represents an ALTER TYPE ... ADD VALUE statement
// Ported from postgres/src/include/nodes/parsenodes.h:3693
type AlterEnumStmt struct {
	BaseNode
	TypeName           *NodeList // qualified name (list of String)
	OldVal             string    // old enum value's name, if renaming
	NewVal             string    // new enum value's name
	NewValNeighbor     string    // neighboring enum value, if specified
	NewValIsAfter      bool      // place new val after neighbor?
	SkipIfNewValExists bool      // ignore statement if new already exists
}

// node implements the Node interface
func (aes *AlterEnumStmt) node() {
	_ = "STUB: not implemented"

	// stmt implements the Stmt interface
	return
}

func (aes *AlterEnumStmt) stmt() {
	_ = "STUB: not implemented"

	// StatementType returns the statement type
	return
}

func (aes *AlterEnumStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

// String returns string representation of AlterEnumStmt
func (aes *AlterEnumStmt) String() string { _ = "STUB: not implemented"; return "" }

// NewAlterEnumStmt creates a new AlterEnumStmt node
func NewAlterEnumStmt(typeName *NodeList) *AlterEnumStmt { _ = "STUB: not implemented"; return nil }

// SqlString returns the SQL representation of AlterEnumStmt
func (aes *AlterEnumStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// Add type name

// RENAME VALUE

// ADD VALUE

// CreateRangeStmt represents a CREATE TYPE ... AS RANGE statement
// Ported from postgres/src/include/nodes/parsenodes.h:3707-3712
type CreateRangeStmt struct {
	BaseNode
	TypeName *NodeList // qualified name (list of String)
	Params   *NodeList // range parameters (list of DefElem)
}

// node implements the Node interface
func (crs *CreateRangeStmt) node() {
	_ = "STUB: not implemented"

	// stmt implements the Stmt interface
	return
}

func (crs *CreateRangeStmt) stmt() {
	_ = "STUB: not implemented"

	// StatementType returns the statement type
	return
}

func (crs *CreateRangeStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

// String returns string representation of CreateRangeStmt
func (crs *CreateRangeStmt) String() string { _ = "STUB: not implemented"; return "" }

// Simplified representation

// NewCreateRangeStmt creates a new CreateRangeStmt node
func NewCreateRangeStmt(typeName *NodeList, params *NodeList) *CreateRangeStmt {
	_ = "STUB: not implemented"
	return nil
}

// SqlString returns the SQL representation of CreateRangeStmt
func (crs *CreateRangeStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// Add type name

// Add parameters

// CreateStatsStmt represents a CREATE STATISTICS statement
// Ported from postgres/src/include/nodes/parsenodes.h:3384-3394
type CreateStatsStmt struct {
	DefNames    *NodeList // qualified name (list of String)
	StatTypes   *NodeList // stat types (list of String)
	Exprs       *NodeList // expressions to build statistics on
	Relations   *NodeList // rels to build stats on (list of RangeVar)
	StxComment  string    // comment to apply to stats, or empty string
	Transformed bool      // true when transformStatsStmt is finished
	IfNotExists bool      // do nothing if stats name already exists
}

// node implements the Node interface
func (css *CreateStatsStmt) node() {
	_ = "STUB: not implemented"

	// stmt implements the Stmt interface
	return
}

func (css *CreateStatsStmt) stmt() {
	_ = "STUB: not implemented"

	// String returns string representation of CreateStatsStmt
	return
}

func (css *CreateStatsStmt) String() string { _ = "STUB: not implemented"; return "" }

func (css *CreateStatsStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

func (css *CreateStatsStmt) Location() int {
	_ = "STUB: not implemented"

	// SetLocation is a no-op; this node does not track source location yet.
	return 0
}

func (css *CreateStatsStmt) SetLocation(int) { _ = "STUB: not implemented"; return }

func (css *CreateStatsStmt) NodeTag() NodeTag { _ = "STUB: not implemented"; return *new(NodeTag) }

func (css *CreateStatsStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// NewCreateStatsStmt creates a new CreateStatsStmt node
func NewCreateStatsStmt(defNames *NodeList, statTypes *NodeList, exprs *NodeList, relations *NodeList, stxComment string, transformed, ifNotExists bool) *CreateStatsStmt {
	_ = "STUB: not implemented"
	return nil
}

// CreatePLangStmt represents a CREATE LANGUAGE statement
// Ported from postgres/src/include/nodes/parsenodes.h:3054-3063
type CreatePLangStmt struct {
	BaseNode
	Replace     bool      // true => replace if already exists
	PLName      string    // PL name
	PLHandler   *NodeList // PL call handler function (qualified name)
	PLInline    *NodeList // optional inline function (qualified name)
	PLValidator *NodeList // optional validator function (qualified name)
	PLTrusted   bool      // PL is trusted
}

// String returns string representation of CreatePLangStmt
func (cpls *CreatePLangStmt) String() string { _ = "STUB: not implemented"; return "" }

func (cpls *CreatePLangStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

func (cpls *CreatePLangStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// NewCreatePLangStmt creates a new CreatePLangStmt node
func NewCreatePLangStmt(replace bool, plName string, plHandler, plInline, plValidator *NodeList, plTrusted bool) *CreatePLangStmt {
	_ = "STUB: not implemented"
	return nil
}

// CreateTableAsStmt represents CREATE TABLE AS and CREATE MATERIALIZED VIEW statements
// Ported from postgres/src/include/nodes/parsenodes.h:3607-3615
type CreateTableAsStmt struct {
	BaseNode
	Query        Node        // the query (generally a SelectStmt)
	Into         *IntoClause // target relation
	ObjType      ObjectType  // table or materialized view
	IsSelectInto bool        // it's a SELECT INTO, not CREATE TABLE AS
	IfNotExists  bool        // IF NOT EXISTS was specified
}

// Location returns the statement's source location (dummy implementation)
func (ctas *CreateTableAsStmt) Location() int {
	_ = "STUB: not implemented"
	// TODO: Implement proper location tracking
	return 0
}

// SetLocation is a no-op; this node does not track source location yet.
func (ctas *CreateTableAsStmt) SetLocation(int) {
	_ = "STUB: not implemented"

	// NodeTag returns the node's type tag
	return
}

func (ctas *CreateTableAsStmt) NodeTag() NodeTag { _ = "STUB: not implemented"; return *new(NodeTag) }

// StatementType returns the statement type for this node
func (ctas *CreateTableAsStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns SQL representation of the CREATE MATERIALIZED VIEW statement
func (ctas *CreateTableAsStmt) SqlString() string {
	_ = "STUB: not implemented"

	// CREATE [TEMP] [MATERIALIZED]
	return ""
}

// Handle TEMP keyword for regular tables

// Check if UNLOGGED

// Target relation name and column list

// Add USING access method if present

// Add WITH options if present

// ON COMMIT action (temp tables); NOOP means no clause was given.

// TABLESPACE

// AS query

// WITH [NO] DATA

// Note: Only add "WITH DATA" explicitly if it was explicitly specified
// PostgreSQL's default is WITH DATA, so we omit it to match original SQL

// String returns string representation of CreateTableAsStmt
func (ctas *CreateTableAsStmt) String() string { _ = "STUB: not implemented"; return "" }

// Add WITH DATA/NO DATA clause if applicable

// Only add "WITH DATA" if it's explicit (to match PostgreSQL behavior)
// For now, we'll omit it since PostgreSQL's default is WITH DATA

// NewCreateTableAsStmt creates a new CreateTableAsStmt node
func NewCreateTableAsStmt(query Node, into *IntoClause, objType ObjectType, isSelectInto, ifNotExists bool) *CreateTableAsStmt {
	_ = "STUB: not implemented"
	return nil
}

// RefreshMatViewStmt represents a REFRESH MATERIALIZED VIEW statement
// Ported from postgres/src/include/nodes/parsenodes.h:3617-3622
type RefreshMatViewStmt struct {
	BaseNode
	Concurrent bool      // allow concurrent access?
	SkipData   bool      // true for WITH NO DATA
	Relation   *RangeVar // relation to refresh
}

// Location returns the statement's source location (dummy implementation)
func (rmvs *RefreshMatViewStmt) Location() int {
	_ = "STUB: not implemented"
	// TODO: Implement proper location tracking
	return 0
}

// SetLocation is a no-op; this node does not track source location yet.
func (rmvs *RefreshMatViewStmt) SetLocation(int) {
	_ = "STUB: not implemented"

	// NodeTag returns the node's type tag
	return
}

func (rmvs *RefreshMatViewStmt) NodeTag() NodeTag { _ = "STUB: not implemented"; return *new(NodeTag) }

// StatementType returns the statement type for this node
func (rmvs *RefreshMatViewStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns SQL representation of the REFRESH MATERIALIZED VIEW statement
func (rmvs *RefreshMatViewStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// String returns string representation of RefreshMatViewStmt
func (rmvs *RefreshMatViewStmt) String() string { _ = "STUB: not implemented"; return "" }

// NewRefreshMatViewStmt creates a new RefreshMatViewStmt node
func NewRefreshMatViewStmt(concurrent, skipData bool, relation *RangeVar) *RefreshMatViewStmt {
	_ = "STUB: not implemented"
	return nil
}

// CreateAssertionStmt represents CREATE ASSERTION statement
// Note: This is not yet implemented in PostgreSQL but the grammar rule exists
// Ported from postgres/src/backend/parser/gram.y:9847-9855
type CreateAssertionStmt struct {
	BaseNode
	Name               *NodeList // Assertion name (qualified name)
	CheckClause        Node      // Check constraint expression
	ConstraintAttrSpec *NodeList // Constraint attributes (e.g., DEFERRABLE)
}

// String returns string representation of CreateAssertionStmt
func (n *CreateAssertionStmt) String() string { _ = "STUB: not implemented"; return "" }

// Add constraint attributes if any

// SqlString returns SQL representation of CreateAssertionStmt
func (n *CreateAssertionStmt) SqlString() string {
	_ = "STUB: not implemented"

	// StatementType returns the statement type
	return ""
}

func (n *CreateAssertionStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

// Location returns the statement's source location
func (n *CreateAssertionStmt) Location() int {
	_ = "STUB: not implemented"

	// SetLocation sets the statement's source location.
	return 0
}

func (n *CreateAssertionStmt) SetLocation(loc int) {
	_ = "STUB: not implemented"

	// NodeTag returns the node's type tag
	return
}

func (n *CreateAssertionStmt) NodeTag() NodeTag { _ = "STUB: not implemented"; return *new(NodeTag) }

// NewCreateAssertionStmt creates a new CreateAssertionStmt node
func NewCreateAssertionStmt(name *NodeList, check Node, attrs *NodeList) *CreateAssertionStmt {
	_ = "STUB: not implemented"
	return nil
}
