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
// Package ast provides PostgreSQL AST node definitions and interfaces.
// Ported from postgres/src/include/nodes/nodes.h and related header files.
package ast

// NodeTag represents the type of an AST node.
// Ported from postgres/src/include/nodes/nodes.h:26-31 (NodeTag enum)
type NodeTag int

// NodeTag constants - ported from postgres/src/include/nodes/nodes.h:28-30
// These represent the fundamental node types in the PostgreSQL AST
const (
	T_Invalid NodeTag = iota // Ported from postgres/src/include/nodes/nodes.h:28

	// Basic node types - will be expanded in full implementation
	T_Node
	T_Query
	T_SelectStmt
	T_InsertStmt
	T_UpdateStmt
	T_DeleteStmt
	T_MergeStmt
	T_CreateStmt
	T_DropStmt
	T_AlterStmt
	T_AlterTableStmt
	T_AlterTableMoveAllStmt
	T_AlterTableCmd
	T_ReplicaIdentityStmt
	T_AlterDomainStmt
	T_IndexStmt
	T_IndexElem
	T_DefElem
	T_Constraint
	T_ViewStmt
	T_CreateSchemaStmt
	T_CreateSeqStmt
	T_AlterSeqStmt
	T_CreateTableAsStmt
	T_CreateAssertionStmt
	T_RefreshMatViewStmt
	T_CreateExtensionStmt
	T_AlterExtensionStmt
	T_AlterExtensionContentsStmt
	T_CreateFdwStmt
	T_AlterFdwStmt
	T_AlterForeignServerStmt
	T_AlterUserMappingStmt
	T_DropUserMappingStmt
	T_CreateEventTrigStmt
	T_AlterEventTrigStmt
	T_CreatedbStmt
	T_DropdbStmt
	T_DropTableSpaceStmt
	T_DropOwnedStmt
	T_ReassignOwnedStmt
	T_ImportForeignSchemaStmt
	T_CreateTableSpaceStmt
	T_AlterTableSpaceStmt
	T_CreateAmStmt
	T_CreateStatsStmt
	T_CreateCastStmt
	T_AlterStatsStmt
	T_AlterFunctionStmt
	T_AlterTypeStmt
	T_CreatePublicationStmt
	T_AlterPublicationStmt
	T_PublicationTable
	T_PublicationObjSpec
	T_CreateSubscriptionStmt
	T_AlterSubscriptionStmt
	T_AlterOpFamilyStmt
	T_CreateOpClassItem
	T_CreateOpClassStmt
	T_CreateOpFamilyStmt
	T_CreateDomainStmt
	T_DefineStmt
	T_CreateEnumStmt
	T_CreateRangeStmt
	T_AlterEnumStmt
	T_CompositeTypeStmt
	T_CreateFunctionStmt
	T_FunctionParameter
	T_CreateConversionStmt
	T_CreateTransformStmt
	T_CreatePLangStmt
	T_RoleSpec
	T_TypeName
	T_CollateClause

	// Utility statement nodes
	T_TransactionStmt
	T_GrantStmt
	T_GrantRoleStmt
	T_AlterDefaultPrivilegesStmt
	T_AccessPriv
	T_CreateRoleStmt
	T_AlterRoleStmt
	T_AlterRoleSetStmt
	T_DropRoleStmt
	T_VariableSetStmt
	T_VariableShowStmt
	T_AlterSystemStmt
	T_ExplainStmt
	T_PrepareStmt
	T_ExecuteStmt
	T_DeallocateStmt
	T_DeclareCursorStmt
	T_FetchStmt
	T_ClosePortalStmt
	T_CopyStmt
	T_VacuumStmt
	T_VacuumRelation
	T_ReindexStmt
	T_ClusterStmt
	T_CheckPointStmt
	T_DiscardStmt
	T_LoadStmt
	T_NotifyStmt
	T_ListenStmt
	T_UnlistenStmt
	T_ConstraintsSetStmt

	// Replication command nodes - ported from postgres/src/include/nodes/replnodes.h
	T_IdentifySystemCmd
	T_CreateReplicationSlotCmd
	T_DropReplicationSlotCmd
	T_AlterReplicationSlotCmd
	T_StartReplicationCmd
	T_ReadReplicationSlotCmd

	// Expression nodes
	T_Expr
	T_Var
	T_Const
	T_Param
	T_Aggref
	T_WindowFunc
	T_FuncExpr
	T_OpExpr
	T_BoolExpr
	T_CaseExpr
	T_ArrayExpr
	T_RowExpr
	T_CoalesceExpr
	T_ScalarArrayOpExpr
	T_SubLink

	// Query execution nodes (essential for SQL functionality)
	T_TargetEntry
	T_FromExpr
	T_JoinExpr
	T_SubPlan
	T_AlternativeSubPlan
	T_CommonTableExpr
	T_WindowClause
	T_SortGroupClause
	T_RowMarkClause
	T_OnConflictExpr

	// Type coercion and advanced expression nodes (essential for PostgreSQL type system)
	T_RelabelType
	T_CoerceViaIO
	T_ArrayCoerceExpr
	T_ConvertRowtypeExpr
	T_CollateExpr
	T_FieldSelect
	T_FieldStore
	T_SubscriptingRef
	T_NullTest
	T_BooleanTest
	T_CoerceToDomain
	T_CoerceToDomainValue
	T_SetToDefault
	T_CurrentOfExpr
	T_NextValueExpr
	T_InferenceElem

	// Administrative and advanced DDL nodes (comprehensive PostgreSQL DDL support)
	T_TableLikeClause
	T_PartitionSpec
	T_PartitionBoundSpec
	T_PartitionRangeDatum
	T_StatsElem
	T_CreateForeignServerStmt
	T_CreateForeignTableStmt
	T_CreateUserMappingStmt
	T_CreateTriggerStmt
	T_CreatePolicyStmt
	T_AlterPolicyStmt
	T_TriggerTransition

	// List and utility nodes
	T_List
	T_ResTarget
	T_RangeVar
	T_ColumnRef
	T_AConst

	// Core parse infrastructure nodes - Stage 1A (25 nodes)
	T_RawStmt
	T_A_Expr
	T_A_Const
	T_ParamRef
	T_TypeCast
	T_ParenExpr
	T_FuncCall
	T_A_Star
	T_A_Indices
	T_A_Indirection
	T_A_ArrayExpr
	T_ColumnDef
	T_WithClause
	T_CTESearchClause
	T_CTECycleClause
	T_MultiAssignRef
	T_WindowDef
	T_SortBy
	T_GroupingSet
	T_LockingClause
	T_XmlSerialize
	T_PartitionElem
	T_TableSampleClause
	T_ObjectWithArgs
	T_SinglePartitionSpec
	T_PartitionCmd

	// Advanced statement nodes - Phase 1B (12 nodes)
	T_SetOperationStmt
	T_ReturnStmt
	T_PLAssignStmt
	T_OnConflictClause
	T_InferClause
	T_WithCheckOption
	T_MergeWhenClause
	T_TruncateStmt
	T_CommentStmt
	T_SecLabelStmt
	T_DoStmt
	T_CallStmt
	T_RenameStmt
	T_AlterOwnerStmt
	T_RuleStmt
	T_LockStmt
	T_AlterObjectSchemaStmt
	T_AlterOperatorStmt
	T_AlterObjectDependsStmt
	T_AlterCollationStmt
	T_AlterDatabaseStmt
	T_AlterDatabaseSetStmt
	T_AlterDatabaseRefreshCollStmt
	T_AlterCompositeTypeStmt
	T_AlterTSConfigurationStmt
	T_AlterTSDictionaryStmt

	// Range table and FROM clause nodes - Phase 1D (9 nodes)
	T_RangeTblEntry
	T_RangeSubselect
	T_RangeFunction
	T_RangeTableFunc
	T_RangeTableFuncCol
	T_RangeTableSample
	T_RangeTblFunction
	T_RTEPermissionInfo
	T_RangeTblRef

	// JSON parse tree nodes - Phase 1E (16 main nodes + 4 supporting types)
	T_JsonFormat
	T_JsonReturning
	T_JsonValueExpr
	T_JsonBehavior
	T_JsonOutput
	T_JsonArgument
	T_JsonFuncExpr
	T_JsonTablePathSpec
	T_JsonTable
	T_JsonTableColumn
	T_JsonKeyValue
	T_JsonParseExpr
	T_JsonScalarExpr
	T_JsonSerializeExpr
	T_JsonObjectConstructor
	T_JsonArrayConstructor
	T_JsonArrayQueryConstructor
	T_JsonAggConstructor
	T_JsonObjectAgg
	T_JsonArrayAgg

	// Phase 1G: JSON Primitive Expressions - primnodes.h JSON expression nodes
	T_JsonConstructorExpr  // JSON constructor expression - primnodes.h:1703
	T_JsonIsPredicate      // JSON IS predicate - primnodes.h:1732
	T_JsonExpr             // JSON expression - primnodes.h:1813
	T_JsonTablePath        // JSON table path - primnodes.h:1867
	T_JsonTablePlan        // JSON table plan - primnodes.h:1882
	T_JsonTablePathScan    // JSON table path scan - primnodes.h:1893
	T_JsonTableSiblingJoin // JSON table sibling join - primnodes.h:1923

	// Phase 1F: Primitive Expression Completion Part 1 - primnodes.h nodes
	T_GroupingFunc           // GROUPING function - primnodes.h:537
	T_WindowFuncRunCondition // Window function run condition - primnodes.h:596
	T_MergeSupportFunc       // Merge support function - primnodes.h:628
	T_NamedArgExpr           // Named argument expression - primnodes.h:787
	T_CaseTestExpr           // CASE test expression - primnodes.h:1352
	T_MinMaxExpr             // MIN/MAX expression - primnodes.h:1506
	T_RowCompareExpr         // Row comparison - primnodes.h:1463
	T_SQLValueFunction       // SQL value function - primnodes.h:1553
	T_XmlExpr                // XML expression - primnodes.h:1596
	T_MergeAction            // MERGE action - primnodes.h:2003
	T_TableFunc              // Table function - primnodes.h:109
	T_IntoClause             // INTO clause - primnodes.h:158

	// Value nodes - ported from postgres/src/include/nodes/value.h
	T_Integer   // Integer literal - postgres/src/include/nodes/value.h:28-34
	T_Float     // Float literal - postgres/src/include/nodes/value.h:47-53
	T_Boolean   // Boolean literal - postgres/src/include/nodes/value.h:55-61
	T_String    // String literal - postgres/src/include/nodes/value.h:63-69
	T_BitString // Bit string literal - postgres/src/include/nodes/value.h:71-77
	T_Null      // NULL literal
)

// String returns the string representation of a NodeTag.
// Used for debugging and error reporting.
func (nt NodeTag) String() string { _ = "STUB: not implemented"; return "" }

// Core parse infrastructure nodes - Stage 1A

// Node is the base interface for all PostgreSQL AST nodes.
// Every node in the parse tree implements this interface.
// Ported from postgres/src/include/nodes/nodes.h:128 (base node concept)
type Node interface {
	// NodeTag returns the type tag for this node
	NodeTag() NodeTag

	// Location returns the byte offset in the source string where this node begins.
	// Returns -1 if location is not available.
	// Ported from postgres/src/include/nodes/parsenodes.h:6-12 location concept
	Location() int

	// SetLocation sets the byte offset in the source string where this node begins.
	SetLocation(location int)

	// String returns a string representation of the node (for debugging)
	String() string

	// SqlString returns the SQL representation of this node for deparsing
	// This enables round-trip parsing: SQL -> AST -> SQL
	SqlString() string
}

// BaseNode provides a basic implementation of the Node interface.
// Other node types should embed this to get default implementations.
// Ported from postgres base node structure concept
type BaseNode struct {
	Tag NodeTag // Node type tag - ported from postgres/src/include/nodes/nodes.h:130
	Loc int     // Source location in bytes - ported from postgres/src/include/nodes/parsenodes.h:6-12
}

// NodeTag returns the node's type tag.
func (n *BaseNode) NodeTag() NodeTag {
	_ = "STUB: not implemented"

	// Location returns the node's source location.
	return *new(NodeTag)
}

func (n *BaseNode) Location() int {
	_ = "STUB: not implemented"

	// String returns a basic string representation.
	return 0
}

func (n *BaseNode) String() string { _ = "STUB: not implemented"; return "" }

// SqlString provides a default implementation that panics with helpful message.
// Specific node types should override this method to provide actual SQL deparsing.
func (n *BaseNode) SqlString() string {
	_ = "STUB: not implemented"
	// Handle T_Invalid nodes specially - they represent invalid/placeholder nodes
	return ""
}

// SetLocation sets the source location for this node.
// Used during parsing to track where nodes came from in the source.
func (n *BaseNode) SetLocation(location int) {
	_ = "STUB: not implemented"

	// NodeList represents a list of nodes.
	// This is a fundamental PostgreSQL concept used throughout the AST.
	// Ported from postgres List structure concept
	return
}

type NodeList struct {
	BaseNode
	Items []Node // List of nodes
}

// NewNodeList creates a new node list.
func NewNodeList(items ...Node) *NodeList { _ = "STUB: not implemented"; return nil }

// Append adds a node to the list.
func (l *NodeList) Append(node Node) { _ = "STUB: not implemented"; return }

// Len returns the number of items in the list.
func (l *NodeList) Len() int { _ = "STUB: not implemented"; return 0 }

// String returns a string representation of the list.
func (l *NodeList) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of NodeList
func (n *NodeList) SqlString() string { _ = "STUB: not implemented"; return "" }

// Stmt represents the base interface for all SQL statements.
// All top-level SQL constructs implement this interface.
// Ported from postgres statement node concept
type Stmt interface {
	Node
	StatementType() string
}

// Expression represents the base interface for all SQL expressions.
// All expressions in WHERE clauses, SELECT lists, etc. implement this.
// Ported from postgres expression node concept
type Expression interface {
	Node
	ExpressionType() string
}

// Identifier represents a simple identifier (table name, column name, etc.).
// Ported from basic identifier concept used throughout postgres AST
type Identifier struct {
	BaseNode
	Name string // The identifier name
}

// NewIdentifier creates a new identifier node.
func NewIdentifier(name string) *Identifier { _ = "STUB: not implemented"; return nil }

// Use T_String for simple identifiers

// String returns the identifier name.
func (i *Identifier) String() string { _ = "STUB: not implemented"; return "" }

// ExpressionType returns the expression type for Expression interface.
func (i *Identifier) ExpressionType() string { _ = "STUB: not implemented"; return "" }

// ==============================================================================
// VALUE NODE SYSTEM - Complete PostgreSQL value.h implementation
// Ported from postgres/src/include/nodes/value.h
// ==============================================================================

// Integer represents an integer literal value node.
// Ported from postgres/src/include/nodes/value.h:28-34
type Integer struct {
	BaseNode
	IVal int // Integer value - postgres/src/include/nodes/value.h:33
}

// NewInteger creates a new integer literal node.
// Ported from postgres/src/include/nodes/value.h:84 (makeInteger)
func NewInteger(value int) *Integer { _ = "STUB: not implemented"; return nil }

func (i *Integer) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of the Integer
func (i *Integer) SqlString() string { _ = "STUB: not implemented"; return "" }

func (i *Integer) ExpressionType() string {
	_ = "STUB: not implemented"

	// IntVal extracts integer value - ported from postgres/src/include/nodes/value.h:79
	return ""
}

func IntVal(node Node) int { _ = "STUB: not implemented"; return 0 }

// Float represents a floating-point literal value node.
// Stored as string to preserve precision, like PostgreSQL.
// Ported from postgres/src/include/nodes/value.h:47-53
type Float struct {
	BaseNode
	FVal string // Float value as string - postgres/src/include/nodes/value.h:52
}

// NewFloat creates a new float literal node.
// Ported from postgres/src/include/nodes/value.h:85 (makeFloat)
func NewFloat(value string) *Float { _ = "STUB: not implemented"; return nil }

func (f *Float) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of the Float
func (f *Float) SqlString() string { _ = "STUB: not implemented"; return "" }

func (f *Float) ExpressionType() string {
	_ = "STUB: not implemented"

	// FloatVal extracts float value - ported from postgres/src/include/nodes/value.h:80
	return ""
}

func FloatVal(node Node) float64 { _ = "STUB: not implemented"; return 0 }

// Boolean represents a boolean literal value node.
// Ported from postgres/src/include/nodes/value.h:55-61
type Boolean struct {
	BaseNode
	BoolVal bool // Boolean value - postgres/src/include/nodes/value.h:60
}

// NewBoolean creates a new boolean literal node.
// Ported from postgres/src/include/nodes/value.h:86 (makeBoolean)
func NewBoolean(value bool) *Boolean { _ = "STUB: not implemented"; return nil }

func (b *Boolean) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of the Boolean
func (b *Boolean) SqlString() string { _ = "STUB: not implemented"; return "" }

func (b *Boolean) ExpressionType() string {
	_ = "STUB: not implemented"

	// BoolVal extracts boolean value - ported from postgres/src/include/nodes/value.h:81
	return ""
}

func BoolVal(node Node) bool { _ = "STUB: not implemented"; return false }

// String represents a string literal value node.
// Ported from postgres/src/include/nodes/value.h:63-69
type String struct {
	BaseNode
	SVal string // String value - postgres/src/include/nodes/value.h:68
}

// NewString creates a new string literal node.
// Ported from postgres/src/include/nodes/value.h:87 (makeString)
func NewString(value string) *String { _ = "STUB: not implemented"; return nil }

func (s *String) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of the String (properly quoted)
func (s *String) SqlString() string { _ = "STUB: not implemented"; return "" }

func (s *String) ExpressionType() string {
	_ = "STUB: not implemented"

	// StrVal extracts string value - ported from postgres/src/include/nodes/value.h:82
	return ""
}

func StrVal(node Node) string { _ = "STUB: not implemented"; return "" }

// BitString represents a bit string literal value node.
// Ported from postgres/src/include/nodes/value.h:71-77
type BitString struct {
	BaseNode
	BSVal string // Bit string value with prefix ('b' or 'x') - postgres/src/include/nodes/value.h:76
}

// NewBitString creates a new bit string literal node.
// Ported from postgres/src/include/nodes/value.h:88 (makeBitString)
// The value should already include the prefix ('b' or 'x') from the lexer,
// matching PostgreSQL's scan.l implementation
func NewBitString(value string) *BitString { _ = "STUB: not implemented"; return nil }

func (bs *BitString) String() string { _ = "STUB: not implemented"; return "" }

func (bs *BitString) ExpressionType() string {
	_ = "STUB: not implemented"

	// SqlString returns the SQL representation of BitString
	return ""
}

func (bs *BitString) SqlString() string {
	_ = "STUB: not implemented"
	// The BSVal contains a prefix character ('b' or 'x') followed by the actual bit string
	// This matches PostgreSQL's scan.l implementation
	return ""
}

// Fallback for unexpected format

// Null represents a NULL literal value node.
type Null struct {
	BaseNode
}

// NewNull creates a new NULL literal node.
func NewNull() *Null { _ = "STUB: not implemented"; return nil }

func (n *Null) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of the Null
func (n *Null) SqlString() string { _ = "STUB: not implemented"; return "" }

func (n *Null) ExpressionType() string {
	_ = "STUB: not implemented"

	// Value is a generic interface for all value types.
	// This provides a common interface for all literal values.
	return ""
}

type Value interface {
	Expression
	IsValue() bool
}

// Implement Value interface for all value types
func (i *Integer) IsValue() bool    { _ = "STUB: not implemented"; return false }
func (f *Float) IsValue() bool      { _ = "STUB: not implemented"; return false }
func (b *Boolean) IsValue() bool    { _ = "STUB: not implemented"; return false }
func (s *String) IsValue() bool     { _ = "STUB: not implemented"; return false }
func (bs *BitString) IsValue() bool { _ = "STUB: not implemented"; return false }
func (n *Null) IsValue() bool {
	_ = "STUB: not implemented"

	// NewValue creates a properly typed value node based on the Go type.
	// This is a convenience function that delegates to the specific typed constructors.
	return false
}

func NewValue(val any) Node { _ = "STUB: not implemented"; return *new(Node) }

// For unknown types, create a string representation

// Helper functions for node creation and manipulation

// IsNode checks if a value implements the Node interface.
func IsNode(v any) bool { _ = "STUB: not implemented"; return false }

// NodeTagOf returns the NodeTag of a node, or T_Invalid if not a node.
func NodeTagOf(v any) NodeTag { _ = "STUB: not implemented"; return *new(NodeTag) }

// LocationOf returns the location of a node, or -1 if not a node or no location.
func LocationOf(v any) int { _ = "STUB: not implemented"; return 0 }

// CastNode safely casts a value to a Node, returning nil if not a node.
func CastNode(v any) Node { _ = "STUB: not implemented"; return *new(Node) }

// NewQualifiedName creates a qualified name node (schema.name)
func NewQualifiedName(schema, name string) Node { _ = "STUB: not implemented"; return *new(Node) }

// NodeWalker is a function type for walking the AST.
// It receives a node and returns whether to continue walking.
type NodeWalker func(Node) bool

// WalkNodes recursively walks all nodes in an AST, calling the walker function.
// This is useful for analysis, transformation, and debugging.
func WalkNodes(node Node, walker NodeWalker) { _ = "STUB: not implemented"; return }

// Handle specific node types that contain other nodes

// Value nodes are leaf nodes - no traversal needed

// Leaf nodes - no children to traverse

// Additional node types will be handled as they're implemented

// For now, we don't traverse into other node types
// This will be expanded as we implement more complex AST structures

// FindNodes finds all nodes of a specific type in an AST.
func FindNodes(root Node, targetTag NodeTag) []Node { _ = "STUB: not implemented"; return nil }

// PrintAST prints a simple representation of an AST for debugging.
func PrintAST(node Node, indent int) { _ = "STUB: not implemented"; return }

// Print children for specific node types
