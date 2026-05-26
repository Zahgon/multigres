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
// ADVANCED STATEMENT NODES - PostgreSQL Phase 1B Implementation
// Ported from postgres/src/include/nodes/parsenodes.h
// ==============================================================================

// ==============================================================================
// MERGE Statement Support
// ==============================================================================

// MergeStmt represents a MERGE statement for conditional INSERT/UPDATE/DELETE operations
// Ported from postgres/src/include/nodes/parsenodes.h:2084-2093
type MergeStmt struct {
	BaseNode
	Relation         *RangeVar   `json:"relation"`         // Target relation to merge into
	SourceRelation   Node        `json:"sourceRelation"`   // Source relation
	JoinCondition    Node        `json:"joinCondition"`    // Join condition between source and target
	MergeWhenClauses *NodeList   `json:"mergeWhenClauses"` // List of WHEN clauses
	ReturningList    *NodeList   `json:"returningList"`    // List of expressions to return
	WithClause       *WithClause `json:"withClause"`       // WITH clause
}

func (n *MergeStmt) node() { _ = "STUB: not implemented"; return }
func (n *MergeStmt) stmt() { _ = "STUB: not implemented"; return }

func (n *MergeStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

func (n *MergeStmt) String() string { _ = "STUB: not implemented"; return "" }

// Get proper table names from RangeVar

// SqlString returns the SQL representation of the MergeStmt
func (m *MergeStmt) SqlString() string {
	_ = "STUB: not implemented"

	// WITH clause
	return ""
}

// MERGE INTO target

// USING source

// ON condition

// WHEN clauses

// RETURNING clause

// NewMergeStmt creates a new MergeStmt node
func NewMergeStmt(relation *RangeVar, sourceRelation Node, joinCondition Node) *MergeStmt {
	_ = "STUB: not implemented"
	return nil
}

// MergeMatchKind represents the type of WHEN clause in MERGE statements
// Ported from postgres/src/include/nodes/primnodes.h:1994-1999
type MergeMatchKind int

const (
	MERGE_WHEN_MATCHED               MergeMatchKind = iota // WHEN MATCHED
	MERGE_WHEN_NOT_MATCHED_BY_SOURCE                       // WHEN NOT MATCHED BY SOURCE
	MERGE_WHEN_NOT_MATCHED_BY_TARGET                       // WHEN NOT MATCHED BY TARGET
)

func (m MergeMatchKind) String() string { _ = "STUB: not implemented"; return "" }

func (m MergeMatchKind) SqlString() string { _ = "STUB: not implemented"; return "" }

// BY TARGET is the default, so we omit it

// Note: OverridingKind is already defined in statements.go

// MergeWhenClause represents individual WHEN clauses in MERGE statements
// Ported from postgres/src/include/nodes/parsenodes.h:1717-1727
type MergeWhenClause struct {
	BaseNode
	MatchKind   MergeMatchKind `json:"matchKind"`   // MATCHED/NOT MATCHED BY SOURCE/TARGET
	CommandType CmdType        `json:"commandType"` // INSERT/UPDATE/DELETE/DO NOTHING
	Override    OverridingKind `json:"override"`    // OVERRIDING clause
	Condition   Node           `json:"condition"`   // WHEN conditions
	TargetList  []*ResTarget   `json:"targetList"`  // INSERT/UPDATE targetlist
	Values      *NodeList      `json:"values"`      // VALUES to INSERT, or NULL
}

func (n *MergeWhenClause) node() { _ = "STUB: not implemented"; return }

func (n *MergeWhenClause) String() string { _ = "STUB: not implemented"; return "" }

func (n *MergeWhenClause) SqlString() string { _ = "STUB: not implemented"; return "" }

// If we have columns specified

// OVERRIDING clause comes after column list, before VALUES

// If no explicit values and no target columns, use DEFAULT VALUES

// Reuse the UPDATE SET renderer so multi-column assignments
// `SET (a, b) = <source>` are regrouped instead of emitted as
// duplicated per-column assignments.

// NewMergeWhenClause creates a new MergeWhenClause node
func NewMergeWhenClause(matchKind MergeMatchKind, commandType CmdType) *MergeWhenClause {
	_ = "STUB: not implemented"
	return nil
}

// ==============================================================================
// SET Operations (UNION, INTERSECT, EXCEPT)
// ==============================================================================

// Note: SetOperation is already defined in statements.go

// SetOperationStmt represents set operations like UNION, INTERSECT, EXCEPT
// Ported from postgres/src/include/nodes/parsenodes.h:2185-2204
type SetOperationStmt struct {
	BaseNode
	Op            SetOperation       `json:"op"`            // Type of set operation
	All           bool               `json:"all"`           // ALL specified?
	Larg          Node               `json:"larg"`          // Left child
	Rarg          Node               `json:"rarg"`          // Right child
	ColTypes      []Oid              `json:"colTypes"`      // OID list of output column type OIDs
	ColTypmods    []int32            `json:"colTypmods"`    // Integer list of output column typmods
	ColCollations []Oid              `json:"colCollations"` // OID list of output column collation OIDs
	GroupClauses  []*SortGroupClause `json:"groupClauses"`  // List of SortGroupClauses
}

func (n *SetOperationStmt) node() { _ = "STUB: not implemented"; return }
func (n *SetOperationStmt) stmt() { _ = "STUB: not implemented"; return }

func (n *SetOperationStmt) String() string { _ = "STUB: not implemented"; return "" }

// NewSetOperationStmt creates a new SetOperationStmt node
func NewSetOperationStmt(op SetOperation, all bool, larg Node, rarg Node) *SetOperationStmt {
	_ = "STUB: not implemented"
	return nil
}

// ==============================================================================
// PL/pgSQL Statement Support
// ==============================================================================

// ReturnStmt represents a RETURN statement in stored procedures/functions
// Ported from postgres/src/include/nodes/parsenodes.h:2210-2214
type ReturnStmt struct {
	BaseNode
	ReturnVal Node `json:"returnval"` // Expression to return
}

func (n *ReturnStmt) node() { _ = "STUB: not implemented"; return }
func (n *ReturnStmt) stmt() { _ = "STUB: not implemented"; return }

func (n *ReturnStmt) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of ReturnStmt
func (n *ReturnStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// StatementType returns the statement type for this node
func (n *ReturnStmt) StatementType() string {
	_ = "STUB: not implemented"

	// NewReturnStmt creates a new ReturnStmt node
	return ""
}

func NewReturnStmt(returnVal Node) *ReturnStmt { _ = "STUB: not implemented"; return nil }

// PLAssignStmt represents PL/pgSQL assignment statements
// Ported from postgres/src/include/nodes/parsenodes.h:2224-2233
type PLAssignStmt struct {
	BaseNode
	Name        string      `json:"name"`        // Initial column name
	Indirection *NodeList   `json:"indirection"` // Subscripts and field names, if any
	Nnames      int         `json:"nnames"`      // Number of names to use in ColumnRef
	Val         *SelectStmt `json:"val"`         // The PL/pgSQL expression to assign
}

func (n *PLAssignStmt) node() { _ = "STUB: not implemented"; return }
func (n *PLAssignStmt) stmt() {
	_ = "STUB: not implemented"

	// StatementType returns the statement type
	return
}

func (n *PLAssignStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

func (n *PLAssignStmt) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of the PLAssignStmt
func (n *PLAssignStmt) SqlString() string {
	_ = "STUB: not implemented"

	// NewPLAssignStmt creates a new PLAssignStmt node
	return ""
}

func NewPLAssignStmt(name string, val *SelectStmt) *PLAssignStmt {
	_ = "STUB: not implemented"
	return nil
}

// ==============================================================================
// INSERT ON CONFLICT Support
// ==============================================================================

// Note: OnConflictAction is already defined in query_execution_nodes.go

// Note: OnConflictClause will be added to statements.go to replace the placeholder

// InferClause represents index inference clause for ON CONFLICT
// Ported from postgres/src/include/nodes/parsenodes.h:1606-1613
type InferClause struct {
	BaseNode
	IndexElems  *NodeList `json:"indexElems"`  // IndexElems to infer unique index
	WhereClause Node      `json:"whereClause"` // Qualification (partial-index predicate)
	Conname     string    `json:"conname"`     // Constraint name, or NULL if unnamed
}

func (n *InferClause) node() { _ = "STUB: not implemented"; return }

func (n *InferClause) String() string { _ = "STUB: not implemented"; return "" }

func (n *InferClause) SqlString() string { _ = "STUB: not implemented"; return "" }

// NewInferClause creates a new InferClause node
func NewInferClause() *InferClause { _ = "STUB: not implemented"; return nil }

// ==============================================================================
// WITH CHECK OPTION Support
// ==============================================================================

// WCOKind represents the type of WITH CHECK OPTION
// Ported from postgres/src/include/nodes/parsenodes.h:1358-1366
type WCOKind int

const (
	WCO_VIEW_CHECK             WCOKind = iota // WCO on an auto-updatable view
	WCO_RLS_INSERT_CHECK                      // RLS INSERT WITH CHECK policy
	WCO_RLS_UPDATE_CHECK                      // RLS UPDATE WITH CHECK policy
	WCO_RLS_CONFLICT_CHECK                    // RLS ON CONFLICT DO UPDATE USING policy
	WCO_RLS_MERGE_UPDATE_CHECK                // RLS MERGE UPDATE USING policy
	WCO_RLS_MERGE_DELETE_CHECK                // RLS MERGE DELETE USING policy
)

func (w WCOKind) String() string { _ = "STUB: not implemented"; return "" }

// WithCheckOption represents WITH CHECK OPTION for views and RLS policies
// Ported from postgres/src/include/nodes/parsenodes.h:1368-1376
type WithCheckOption struct {
	BaseNode
	Kind     WCOKind `json:"kind"`     // Kind of WCO
	Relname  string  `json:"relname"`  // Name of relation that specified the WCO
	Polname  string  `json:"polname"`  // Name of RLS policy being checked
	Qual     Node    `json:"qual"`     // Constraint qual to check
	Cascaded bool    `json:"cascaded"` // True for a cascaded WCO on a view
}

func (n *WithCheckOption) node() { _ = "STUB: not implemented"; return }

func (n *WithCheckOption) String() string { _ = "STUB: not implemented"; return "" }

// NewWithCheckOption creates a new WithCheckOption node
func NewWithCheckOption(kind WCOKind, cascaded bool) *WithCheckOption {
	_ = "STUB: not implemented"
	return nil
}

// ==============================================================================
// Additional Statement Types
// ==============================================================================

// TruncateStmt represents TRUNCATE TABLE statements
// Ported from postgres/src/include/nodes/parsenodes.h:3240-3246
type TruncateStmt struct {
	BaseNode
	Relations   *NodeList    `json:"relations"`   // Relations (RangeVars) to be truncated
	RestartSeqs bool         `json:"restartSeqs"` // Restart owned sequences?
	Behavior    DropBehavior `json:"behavior"`    // RESTRICT or CASCADE behavior
}

func (n *TruncateStmt) node() { _ = "STUB: not implemented"; return }
func (n *TruncateStmt) stmt() { _ = "STUB: not implemented"; return }

func (n *TruncateStmt) String() string { _ = "STUB: not implemented"; return "" }

// NewTruncateStmt creates a new TruncateStmt node
func NewTruncateStmt(relations *NodeList) *TruncateStmt { _ = "STUB: not implemented"; return nil }

// StatementType implements the Stmt interface
func (n *TruncateStmt) StatementType() string {
	_ = "STUB: not implemented"

	// SqlString returns the SQL representation of the TRUNCATE statement
	return ""
}

func (n *TruncateStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// CommentStmt represents COMMENT ON statements
// Ported from postgres/src/include/nodes/parsenodes.h:3252-3258
type CommentStmt struct {
	BaseNode
	Objtype ObjectType `json:"objtype"` // Object's type
	Object  Node       `json:"object"`  // Qualified name of the object
	Comment string     `json:"comment"` // Comment to insert, or NULL to remove
}

func (n *CommentStmt) node() { _ = "STUB: not implemented"; return }
func (n *CommentStmt) stmt() { _ = "STUB: not implemented"; return }

func (n *CommentStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

func (n *CommentStmt) String() string { _ = "STUB: not implemented"; return "" }

func (n *CommentStmt) SqlString() string {
	_ = "STUB: not implemented"

	// Handle special object types that need custom formatting
	return ""
}

// Last item is constraint name, preceding items are table name

// First item is TypeName, second is constraint name (as String)

// Get the domain TypeName and format it

// Get the constraint name

// First item is Typename, second is language name

// First item is access method, rest are class name parts

// First item is access method, rest are family name parts

// NumericOnly should format directly

// For RULE, NodeList has [table_name, rule_name] but we want "RULE rule_name ON table_name"

// Fallback to regular formatting

// For TRIGGER with schema, NodeList has [schema, table, trigger_name]
// We want "TRIGGER trigger_name ON schema.table"

// For TRIGGER without schema, NodeList has [table_name, trigger_name]
// We want "TRIGGER trigger_name ON table_name"

// Fallback to regular formatting

// First is source type, second is target type

// Regular format for other object types

// formatObjectName formats a Node (NodeList of String nodes or single String) as a qualified identifier
func formatObjectName(obj Node) string { _ = "STUB: not implemented"; return "" }

// Handle single String nodes (from object_type_name rules)

// Fallback to regular SqlString for other node types (Typename, etc.)

// NewCommentStmt creates a new CommentStmt node
func NewCommentStmt(objtype ObjectType, object Node, comment string) *CommentStmt {
	_ = "STUB: not implemented"
	return nil
}

// SecLabelStmt represents SECURITY LABEL statements
// Ported from postgres/src/include/nodes/parsenodes.h:3267-3274
type SecLabelStmt struct {
	BaseNode
	Objtype  ObjectType `json:"objtype"`  // Object's type
	Object   Node       `json:"object"`   // Qualified name of the object
	Provider string     `json:"provider"` // Label provider (or NULL)
	Label    string     `json:"label"`    // New security label to be assigned
}

func (n *SecLabelStmt) node() { _ = "STUB: not implemented"; return }
func (n *SecLabelStmt) stmt() { _ = "STUB: not implemented"; return }

func (n *SecLabelStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

func (n *SecLabelStmt) String() string { _ = "STUB: not implemented"; return "" }

func (n *SecLabelStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// Format object name as qualified identifier, not string literal

// NewSecLabelStmt creates a new SecLabelStmt node
func NewSecLabelStmt(objtype ObjectType, object Node, provider string, label string) *SecLabelStmt {
	_ = "STUB: not implemented"
	return nil
}

// DoStmt represents DO statements
// Ported from postgres/src/include/nodes/parsenodes.h:3437-3441
type DoStmt struct {
	BaseNode
	Args *NodeList `json:"args"` // List of DefElem nodes
}

func (n *DoStmt) node() { _ = "STUB: not implemented"; return }
func (n *DoStmt) stmt() { _ = "STUB: not implemented"; return }

func (n *DoStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

func (n *DoStmt) String() string { _ = "STUB: not implemented"; return "" }

func (n *DoStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// Emit the options in their original order: the code block and an
// optional LANGUAGE may appear in either order, and re-ordering them
// changes the parsed option list (so the round-trip would differ).

// Use single quotes for the code block

// NewDoStmt creates a new DoStmt node
func NewDoStmt(args *NodeList) *DoStmt { _ = "STUB: not implemented"; return nil }

// CallStmt represents CALL statements
// Ported from postgres/src/include/nodes/parsenodes.h:3451-3458
type CallStmt struct {
	BaseNode
	Funccall *FuncCall `json:"funccall"` // from the parser
	// Note: funcexpr and outargs are used at execution time, not parsing
}

func (n *CallStmt) node() { _ = "STUB: not implemented"; return }
func (n *CallStmt) stmt() { _ = "STUB: not implemented"; return }

func (n *CallStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

func (n *CallStmt) String() string { _ = "STUB: not implemented"; return "" }

func (n *CallStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// NewCallStmt creates a new CallStmt node
func NewCallStmt(funccall *FuncCall) *CallStmt { _ = "STUB: not implemented"; return nil }

// RenameStmt represents ALTER ... RENAME TO statements
// Ported from postgres/src/include/nodes/parsenodes.h:3525-3537
type RenameStmt struct {
	BaseNode
	RenameType   ObjectType   `json:"renameType"`   // OBJECT_TABLE, OBJECT_COLUMN, etc
	RelationType ObjectType   `json:"relationType"` // If column name, associated relation type
	Relation     *RangeVar    `json:"relation"`     // In case it's a table
	Object       Node         `json:"object"`       // In case it's some other object
	Subname      string       `json:"subname"`      // Name of contained object (column, rule, trigger, etc)
	Newname      string       `json:"newname"`      // The new name
	Behavior     DropBehavior `json:"behavior"`     // RESTRICT or CASCADE behavior
	MissingOk    bool         `json:"missingOk"`    // Skip error if missing?
}

func (n *RenameStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

func (n *RenameStmt) String() string { _ = "STUB: not implemented"; return "" }

// formatRenameObjectName formats the object name for a RenameStmt based on the object type
// Uses QuoteIdentifier to properly handle identifier quoting rules
func formatRenameObjectName(renameType ObjectType, object Node) string {
	_ = "STUB: not implemented"
	return ""
}

// Handle NodeList objects - format as qualified identifier

// Special handling for OPERATOR CLASS and OPERATOR FAMILY
// When NodeList has 2 items, first is access method, second is name

// Format as "name USING method"

// Default NodeList handling for other cases

// Use QuoteIdentifier for proper identifier quoting

// Handle String objects - format as identifier

// Default: use object's SqlString

func (n *RenameStmt) SqlString() string {
	_ = "STUB: not implemented"

	// Start with ALTER
	return ""
}

// Add object type

// Column rename requires the relation type

// Use the generic String() method for any unhandled cases

// Add IF EXISTS if specified (for POLICY/RULE/TRIGGER, it goes after the object name)

// Handle different rename patterns

// These have special syntax: ALTER <type> [IF EXISTS] old_name ON tablename RENAME TO new_name

// Add old name

// Add ON tablename

// Column rename: ALTER <type> tablename RENAME [COLUMN] old_name TO new_name

// Constraint rename: ALTER TABLE tablename RENAME CONSTRAINT old_name TO new_name

// Domain constraint rename: ALTER DOMAIN domain_name RENAME CONSTRAINT old_name TO new_name

// Use the helper function to format domain name properly

// Type attribute rename: ALTER TYPE type_name RENAME ATTRIBUTE old_name TO new_name

// Default rename: ALTER <type> old_name RENAME TO new_name

// Use the helper function to format object name based on type

// For DATABASE, SCHEMA, ROLE, etc., the old name might be in Subname

// Add CASCADE behavior if specified

// NewRenameStmt creates a new RenameStmt node
func NewRenameStmt(renameType ObjectType, newname string) *RenameStmt {
	_ = "STUB: not implemented"
	return nil
}

// AlterOwnerStmt represents ALTER ... OWNER TO statements
// Ported from postgres/src/include/nodes/parsenodes.h:3571-3578
type AlterOwnerStmt struct {
	BaseNode
	ObjectType ObjectType `json:"objectType"` // OBJECT_TABLE, OBJECT_TYPE, etc
	Relation   *RangeVar  `json:"relation"`   // In case it's a table
	Object     Node       `json:"object"`     // In case it's some other object
	Newowner   *RoleSpec  `json:"newowner"`   // The new owner
}

func (n *AlterOwnerStmt) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of the AlterOwnerStmt
func (n *AlterOwnerStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// Add object type

// Add object name

// Special handling for NodeList objects

// Format as "name USING method" where first item is method, rest is name

// For other object types, handle as qualified identifier names

// Handle simple string identifiers

// Add OWNER TO clause

// StatementType returns the statement type
func (n *AlterOwnerStmt) StatementType() string {
	_ = "STUB: not implemented"

	// NewAlterOwnerStmt creates a new AlterOwnerStmt node
	return ""
}

func NewAlterOwnerStmt(objectType ObjectType, newowner *RoleSpec) *AlterOwnerStmt {
	_ = "STUB: not implemented"
	return nil
}

// RuleStmt represents CREATE RULE statements
// Ported from postgres/src/include/nodes/parsenodes.h:3606-3616
type RuleStmt struct {
	BaseNode
	Relation    *RangeVar `json:"relation"`    // Relation the rule is for
	Rulename    string    `json:"rulename"`    // Name of the rule
	WhereClause Node      `json:"whereClause"` // Qualifications
	Event       CmdType   `json:"event"`       // SELECT, INSERT, etc
	Instead     bool      `json:"instead"`     // Is a 'do instead'?
	Actions     *NodeList `json:"actions"`     // The action statements
	Replace     bool      `json:"replace"`     // OR REPLACE
}

func (n *RuleStmt) String() string { _ = "STUB: not implemented"; return "" }

// Use SqlString() for proper expression deparsing

func (n *RuleStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

func (n *RuleStmt) Location() int { _ = "STUB: not implemented"; return 0 }

func (n *RuleStmt) NodeTag() NodeTag { _ = "STUB: not implemented"; return *new(NodeTag) }

func (n *RuleStmt) SqlString() string {
	_ = "STUB: not implemented"

	// NewRuleStmt creates a new RuleStmt node
	return ""
}

func NewRuleStmt(relation *RangeVar, rulename string, event CmdType) *RuleStmt {
	_ = "STUB: not implemented"
	return nil
}

// Lock mode constants based on PostgreSQL's LOCKMODE
// From postgres/src/include/storage/lockdefs.h
type LockMode int

const (
	NoLock LockMode = iota
	AccessShareLock
	RowShareLock
	RowExclusiveLock
	ShareUpdateExclusiveLock
	ShareLock
	ShareRowExclusiveLock
	ExclusiveLock
	AccessExclusiveLock
)

// LockStmt represents LOCK TABLE statements
// Ported from postgres/src/include/nodes/parsenodes.h:3942-3948
type LockStmt struct {
	BaseNode
	Relations *NodeList `json:"relations"` // Relations to lock
	Mode      LockMode  `json:"mode"`      // Lock mode
	Nowait    bool      `json:"nowait"`    // No wait mode
}

func (n *LockStmt) String() string { _ = "STUB: not implemented"; return "" }

// Add lock mode

// NewLockStmt creates a new LockStmt node
func NewLockStmt(relations *NodeList, mode LockMode) *LockStmt {
	_ = "STUB: not implemented"
	return nil
}

// StatementType implements the Stmt interface
func (n *LockStmt) StatementType() string {
	_ = "STUB: not implemented"

	// SqlString returns the SQL representation of the LOCK statement
	return ""
}

func (n *LockStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// Add lock mode
