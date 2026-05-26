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
// Package ast provides PostgreSQL AST statement node definitions.
// Ported from postgres/src/include/nodes/parsenodes.h
package ast

// ==============================================================================
// CORE STATEMENT FRAMEWORK - PostgreSQL parsenodes.h implementation
// Ported from postgres/src/include/nodes/parsenodes.h
// ==============================================================================

// CmdType represents the type of SQL command - ported from postgres/src/include/nodes/nodes.h:262
type CmdType int

const (
	CMD_UNKNOWN CmdType = iota // Unknown command type
	CMD_SELECT                 // SELECT statement
	CMD_UPDATE                 // UPDATE statement
	CMD_INSERT                 // INSERT statement
	CMD_DELETE                 // DELETE statement
	CMD_MERGE                  // MERGE statement
	CMD_UTILITY                // Utility commands (CREATE, DROP, etc.)
	CMD_NOTHING                // Dummy command for INSTEAD NOTHING rules
)

func (c CmdType) String() string { _ = "STUB: not implemented"; return "" }

// QuerySource represents possible sources of a Query - ported from postgres/src/include/nodes/parsenodes.h:34
type QuerySource int

const (
	QSRC_ORIGINAL          QuerySource = iota // Original parse tree (explicit query)
	QSRC_PARSER                               // Added by parse analysis (now unused)
	QSRC_INSTEAD_RULE                         // Added by unconditional INSTEAD rule
	QSRC_QUAL_INSTEAD_RULE                    // Added by conditional INSTEAD rule
	QSRC_NON_INSTEAD_RULE                     // Added by non-INSTEAD rule
)

// Note: DropBehavior and ObjectType are now defined in ddl_statements.go

// SetQuantifier represents set quantifier options for GROUP BY and SELECT DISTINCT.
// Ported from postgres/src/include/nodes/parsenodes.h:61
type SetQuantifier int

const (
	SET_QUANTIFIER_DEFAULT  SetQuantifier = iota // No quantifier specified
	SET_QUANTIFIER_ALL                           // ALL quantifier
	SET_QUANTIFIER_DISTINCT                      // DISTINCT quantifier
)

// LimitOption represents LIMIT clause options.
// Ported from postgres/src/include/nodes/nodes.h:428
type LimitOption int

const (
	LIMIT_OPTION_COUNT     LimitOption = iota // FETCH FIRST... ONLY - nodes.h:430
	LIMIT_OPTION_WITH_TIES                    // FETCH FIRST... WITH TIES - nodes.h:431
)

// OnCommitAction represents actions for temporary tables on transaction commit.
// Ported from postgres/src/include/nodes/primnodes.h:55
type OnCommitAction int

const (
	ONCOMMIT_NOOP          OnCommitAction = iota // No ON COMMIT clause (do nothing) - primnodes.h:57
	ONCOMMIT_PRESERVE_ROWS                       // ON COMMIT PRESERVE ROWS (do nothing) - primnodes.h:58
	ONCOMMIT_DELETE_ROWS                         // ON COMMIT DELETE ROWS - primnodes.h:59
	ONCOMMIT_DROP                                // ON COMMIT DROP - primnodes.h:60
)

// RelPersistence represents table persistence types.
// Ported from postgres/src/include/nodes/primnodes.h:87
const (
	RELPERSISTENCE_PERMANENT rune = 'p' // Regular table - primnodes.h:89
	RELPERSISTENCE_UNLOGGED  rune = 'u' // Unlogged table - primnodes.h:90
	RELPERSISTENCE_TEMP      rune = 't' // Temporary table - primnodes.h:91
)

// ==============================================================================
// SUPPORTING STRUCTURES
// ==============================================================================

// RangeVar represents a table/relation reference.
// Ported from postgres/src/include/nodes/primnodes.h:71
type RangeVar struct {
	BaseNode
	CatalogName    string // Database name, or empty - postgres/src/include/nodes/primnodes.h:76
	SchemaName     string // Schema name, or empty - postgres/src/include/nodes/primnodes.h:79
	RelName        string // Relation/sequence name - postgres/src/include/nodes/primnodes.h:82
	Inh            bool   // Expand relation by inheritance? - postgres/src/include/nodes/primnodes.h:85
	RelPersistence rune   // Persistence type - postgres/src/include/nodes/primnodes.h:87
	Alias          *Alias // Table alias & optional column aliases - postgres/src/include/nodes/primnodes.h:90
}

// SqlString returns the SQL representation of this table reference
func (r *RangeVar) SqlString() string {
	_ = "STUB: not implemented"
	// Use utility function for qualified name formatting
	return ""
}

// Add ONLY prefix if inheritance is disabled

// Add alias if present

// NewRangeVar creates a new RangeVar node.
func NewRangeVar(relName string, schemaName, catalogName string) *RangeVar {
	_ = "STUB: not implemented"
	return nil
}

// Default to inheritance enabled (no ONLY)

func (rv *RangeVar) String() string { _ = "STUB: not implemented"; return "" }

func (rv *RangeVar) StatementType() string {
	_ = "STUB: not implemented"

	// Alias represents table and column aliases.
	// Ported from postgres/src/include/nodes/primnodes.h:47
	return ""
}

type Alias struct {
	BaseNode
	AliasName string    // Alias name - postgres/src/include/nodes/primnodes.h:50
	ColNames  *NodeList // Column aliases - postgres/src/include/nodes/primnodes.h:51
}

// SqlString returns the SQL representation of this alias
func (a *Alias) SqlString() string { _ = "STUB: not implemented"; return "" }

// Add column aliases if present

// NewAlias creates a new Alias node.
func NewAlias(aliasName string, colNames *NodeList) *Alias { _ = "STUB: not implemented"; return nil }

// Use T_String for alias

func (a *Alias) String() string { _ = "STUB: not implemented"; return "" }

// ResTarget represents a target item in a SELECT list or UPDATE SET clause.
// Ported from postgres/src/include/nodes/parsenodes.h:514
type ResTarget struct {
	BaseNode
	Name        string    // Column name or empty - postgres/src/include/nodes/parsenodes.h:518
	Indirection *NodeList // Subscripts, field names, and '*', or nil - postgres/src/include/nodes/parsenodes.h:519
	Val         Node      // Value expression to compute or assign - postgres/src/include/nodes/parsenodes.h:520
}

// NewResTarget creates a new ResTarget node.
func NewResTarget(name string, val Node) *ResTarget { _ = "STUB: not implemented"; return nil }

// NewResTargetWithIndirection creates a new ResTarget node with indirection (for column references with array subscripts).
func NewResTargetWithIndirection(name string, indirection *NodeList) *ResTarget {
	_ = "STUB: not implemented"
	return nil
}

// ColumnNameWithIndirection returns the column name with indirection (array subscripts, field access, etc.)
// This is used for INSERT column lists where we need "column[index]" format
func (r *ResTarget) ColumnNameWithIndirection() string { _ = "STUB: not implemented"; return "" }

// Field selection

// Array index or slice

// Generic indirection

// renderSetClauses renders the assignment list of a SET clause (UPDATE or
// ON CONFLICT DO UPDATE). A multi-column assignment `SET (a, b, ...) = <source>`
// is parsed into one ResTarget per column, each holding a MultiAssignRef that
// points at the shared source; those are regrouped so the source is emitted once
// (rendering each as `a = <source>` would be wrong).
func renderSetClauses(items []Node) []string { _ = "STUB: not implemented"; return nil }

// SetClauseString returns the SQL representation for SET clauses (UPDATE, ON CONFLICT DO UPDATE)
// Format: "column[index] = value" instead of "value AS column"
func (r *ResTarget) SetClauseString() string { _ = "STUB: not implemented"; return "" }

// Build the column name with indirection (e.g., "col[1]", "col.field")

// If no value, just return the column name

func (rt *ResTarget) String() string { _ = "STUB: not implemented"; return "" }

func (rt *ResTarget) ExpressionType() string {
	_ = "STUB: not implemented"

	// SqlString returns the SQL representation of the ResTarget
	return ""
}

func (r *ResTarget) SqlString() string { _ = "STUB: not implemented"; return "" }

// Add indirection if present (e.g., array subscripts, field selection)

// Handle different types of indirection

// Field selection

// Array index or slice

// Generic indirection

// Add alias if present

// ==============================================================================
// CORE QUERY STRUCTURE
// ==============================================================================

// Query is the fundamental query structure that all parsed queries transform into.
// Ported from postgres/src/include/nodes/parsenodes.h:117
type Query struct {
	BaseNode
	CommandType    CmdType     // select|insert|update|delete|merge|utility - postgres/src/include/nodes/parsenodes.h:120
	QuerySource    QuerySource // Where did this query come from? - postgres/src/include/nodes/parsenodes.h:121
	QueryId        uint64      // Query identifier - postgres/src/include/nodes/parsenodes.h:122
	CanSetTag      bool        // Do I set the command result tag? - postgres/src/include/nodes/parsenodes.h:123
	UtilityStmt    Node        // Non-null if commandType == CMD_UTILITY - postgres/src/include/nodes/parsenodes.h:124
	ResultRelation int         // rtable index of target relation - postgres/src/include/nodes/parsenodes.h:125

	// Boolean flags for query characteristics - postgres/src/include/nodes/parsenodes.h:127-140
	HasAggs         bool // Has aggregates in tlist or havingQual
	HasWindowFuncs  bool // Has window functions in tlist
	HasTargetSRFs   bool // Has set-returning functions in tlist
	HasSubLinks     bool // Has subquery SubLink
	HasDistinctOn   bool // distinctClause is from DISTINCT ON
	HasRecursive    bool // WITH RECURSIVE was specified
	HasModifyingCTE bool // Has INSERT/UPDATE/DELETE/MERGE in WITH
	HasForUpdate    bool // FOR [KEY] UPDATE/SHARE was specified
	HasRowSecurity  bool // Rewriter has applied some RLS policy
	IsReturn        bool // Is a RETURN statement

	// Query components - postgres/src/include/nodes/parsenodes.h:142-192
	CteList             []*CommonTableExpr // WITH list
	Rtable              []*RangeTblEntry   // Range table entries
	RtePermInfos        *NodeList          // Permission info for rtable entries - parsenodes.h:174
	Jointree            *FromExpr          // Table join tree (FROM and WHERE clauses)
	MergeActionList     *NodeList          // MERGE statement actions - parsenodes.h:178
	MergeTargetRelation int                // MERGE target relation index - parsenodes.h:186
	MergeJoinCondition  Node               // JOIN condition for MERGE - parsenodes.h:189
	TargetList          []*TargetEntry     // Target list
	Override            OverridingKind     // OVERRIDING clause - parsenodes.h:194
	OnConflict          *OnConflictExpr    // ON CONFLICT expression - parsenodes.h:196
	ReturningList       []*TargetEntry     // Return-values list
	GroupClause         []*SortGroupClause // GROUP BY clauses
	GroupDistinct       bool               // Is the GROUP BY clause distinct?
	GroupingSets        *NodeList          // GROUPING SETS if present
	HavingQual          Node               // Qualifications applied to groups
	WindowClause        []*WindowClause    // WINDOW clauses
	DistinctClause      []*SortGroupClause // DISTINCT clauses
	SortClause          []*SortGroupClause // ORDER BY clauses
	LimitOffset         Node               // Number of result tuples to skip
	LimitCount          Node               // Number of result tuples to return
	LimitOption         LimitOption        // Limit type option - parsenodes.h:215
	RowMarks            []*RowMarkClause   // Row mark clauses
	SetOperations       Node               // Set operation tree - parsenodes.h:219
	ConstraintDeps      []Oid              // Constraint dependencies - parsenodes.h:226
	WithCheckOptions    *NodeList          // WITH CHECK OPTIONS - parsenodes.h:228
	StmtLocation        int                // Start location - postgres/src/include/nodes/parsenodes.h:239
	StmtLen             int                // Length in bytes - postgres/src/include/nodes/parsenodes.h:240
}

// NewQuery creates a new Query node.
func NewQuery(cmdType CmdType) *Query { _ = "STUB: not implemented"; return nil }

func (q *Query) String() string { _ = "STUB: not implemented"; return "" }

func (q *Query) StatementType() string { _ = "STUB: not implemented"; return "" }

// ==============================================================================
// DML STATEMENTS
// ==============================================================================

// GroupClause represents a GROUP BY clause with quantifier.
// Private struct for the result of group_clause production
// Ported from postgres/src/backend/parser/gram.y:135
type GroupClause struct {
	Distinct bool      // true if GROUP BY DISTINCT
	List     *NodeList // list of GROUP BY expressions
}

// SelectStmt represents a raw SELECT statement before analysis.
// Ported from postgres/src/include/nodes/parsenodes.h:2116
type SelectStmt struct {
	BaseNode
	// Fields used in "leaf" SelectStmts - postgres/src/include/nodes/parsenodes.h:2120-2130
	DistinctClause *NodeList   // NULL, list of DISTINCT ON exprs, or special marker for ALL
	IntoClause     *IntoClause // Target for SELECT INTO
	TargetList     *NodeList   // Target list
	FromClause     *NodeList   // FROM clause
	WhereClause    Node        // WHERE qualification
	GroupClause    *NodeList   // GROUP BY clauses
	GroupDistinct  bool        // Is this GROUP BY DISTINCT?
	HavingClause   Node        // HAVING conditional-expression
	WindowClause   *NodeList   // WINDOW window_name AS (...), ...
	ValuesLists    *NodeList   // Untransformed list of expression lists

	// Fields used in both "leaf" and upper-level SelectStmts - postgres/src/include/nodes/parsenodes.h:2132-2137
	SortClause    *NodeList   // Sort clause
	LimitOffset   Node        // Number of result tuples to skip
	LimitCount    Node        // Number of result tuples to return
	LimitOption   LimitOption // Limit type option
	LockingClause *NodeList   // FOR UPDATE clauses
	WithClause    *WithClause // WITH clause

	// Fields used only in upper-level SelectStmts - postgres/src/include/nodes/parsenodes.h:2139-2143
	Op   SetOperation // Type of set operation
	All  bool         // ALL specified?
	Larg *SelectStmt  // Left child
	Rarg *SelectStmt  // Right child
}

// NewSelectStmt creates a new SelectStmt node.
func NewSelectStmt() *SelectStmt { _ = "STUB: not implemented"; return nil }

func (s *SelectStmt) String() string { _ = "STUB: not implemented"; return "" }

func (s *SelectStmt) StatementType() string {
	_ = "STUB: not implemented"

	// SqlString returns the SQL representation of the SelectStmt
	return ""
}

func (s *SelectStmt) SqlString() string {
	_ = "STUB: not implemented"

	// Handle set operations (UNION, INTERSECT, EXCEPT)
	return ""
}

// Add left operand, with parentheses only if needed for complex queries

// Handle ORDER BY clause for set operations

// Handle LIMIT clause for set operations

// Use FETCH FIRST syntax for WITH TIES (required by SQL standard)

// Check if the limit count is a constant 1 (implicit case)

// Traditional LIMIT syntax (default for LIMIT_OPTION_COUNT)
// Handle LIMIT ALL case specially

// Handle OFFSET clause for set operations

// FOR UPDATE/SHARE clauses also attach to the set-operation node.

// WITH clause attaches to the outer set-operation node and must be
// prepended here. The non-set-op branch handles WithClause below; the
// set-op branch returns early, so without this the CTE is silently
// dropped on round-trip and PG sees `relation "cte" does not exist`.

// Handle VALUES clause

// WITH cte(...) AS (...) VALUES (...) attaches WithClause to the
// outer SelectStmt; same drop-on-round-trip class as the set-op
// branch above.

// Regular SELECT statement

// DISTINCT clause
// Note: DistinctClause == nil means no DISTINCT
//       DistinctClause == &NodeList{Items: []} means plain DISTINCT
//       DistinctClause == &NodeList{Items: [expr1, expr2]} means DISTINCT ON (...)

// Check if there are any expressions (means DISTINCT ON)

// Always add DISTINCT part if DistinctClause is not nil

// Target list (what to select)

// INTO clause

// FROM clause

// WHERE clause

// GROUP BY clause

// HAVING clause

// WINDOW clause

// For WINDOW clause, format as "name AS (specification)"

// ORDER BY clause (from SortClause)

// LIMIT clause

// Use FETCH FIRST syntax for WITH TIES (required by SQL standard)

// Check if the limit count is a constant 1 (implicit case)

// Traditional LIMIT syntax (default for LIMIT_OPTION_COUNT)
// Handle LIMIT ALL case specially

// OFFSET clause

// FOR UPDATE/SHARE clauses

// WITH clause (CTEs)

// WITH clause typically comes first, so we need to prepend it

// needsParenthesesInSetOperation determines if a SelectStmt needs parentheses when used in a set operation
func (s *SelectStmt) needsParenthesesInSetOperation() bool { _ = "STUB: not implemented"; return false }

// A SELECT needs parentheses if it has any of these complex clauses

// WITH clauses need parentheses in set operations
// Nested set operations always need parentheses

// isConstantOne checks if a Node represents the constant integer 1
func isConstantOne(node Node) bool { _ = "STUB: not implemented"; return false }

// Check if it's an A_Const with integer value 1

// InsertStmt represents an INSERT statement.
// Ported from postgres/src/include/nodes/parsenodes.h:2039
type InsertStmt struct {
	BaseNode
	Relation         *RangeVar         // Relation to insert into - postgres/src/include/nodes/parsenodes.h:2042
	Cols             *NodeList         // Optional: names of the target columns - postgres/src/include/nodes/parsenodes.h:2043
	SelectStmt       Node              // Source SELECT/VALUES, or NULL - postgres/src/include/nodes/parsenodes.h:2044
	OnConflictClause *OnConflictClause // ON CONFLICT clause - postgres/src/include/nodes/parsenodes.h:2045
	ReturningList    *NodeList         // List of expressions to return - postgres/src/include/nodes/parsenodes.h:2046
	WithClause       *WithClause       // WITH clause - postgres/src/include/nodes/parsenodes.h:2047
	Override         OverridingKind    // OVERRIDING clause - postgres/src/include/nodes/parsenodes.h:2048
}

// NewInsertStmt creates a new InsertStmt node.
func NewInsertStmt(relation *RangeVar) *InsertStmt { _ = "STUB: not implemented"; return nil }

func (i *InsertStmt) String() string { _ = "STUB: not implemented"; return "" }

func (i *InsertStmt) StatementType() string {
	_ = "STUB: not implemented"

	// SqlString returns the SQL representation of the InsertStmt
	return ""
}

func (i *InsertStmt) SqlString() string {
	_ = "STUB: not implemented"

	// WITH clause
	return ""
}

// INSERT INTO table

// Column list (if specified)

// OVERRIDING { SYSTEM | USER } VALUE clause — semantically significant
// when inserting into GENERATED { ALWAYS | BY DEFAULT } AS IDENTITY columns.

// SelectStmt/VALUES clause

// If the INSERT has parentheses around the SELECT, we need to preserve them
// This is determined by checking if the original query had parentheses
// For now, we'll check if this is a simple SELECT vs a subquery by looking at the SelectStmt
// If it has a WHERE clause or other complexity, it's likely a subquery that should be parenthesized

// This appears to be a complex SELECT that was likely parenthesized in the original

// Not a SelectStmt, could be VALUES clause, append as is

// DEFAULT VALUES case (SelectStmt is nil)

// ON CONFLICT clause

// RETURNING clause

// UpdateStmt represents an UPDATE statement.
// Ported from postgres/src/include/nodes/parsenodes.h:2069
type UpdateStmt struct {
	BaseNode
	Relation      *RangeVar   // Relation to update - postgres/src/include/nodes/parsenodes.h:2072
	TargetList    *NodeList   // Target list (of ResTarget) - postgres/src/include/nodes/parsenodes.h:2073
	WhereClause   Node        // Qualifications - postgres/src/include/nodes/parsenodes.h:2074
	FromClause    *NodeList   // Optional from clause for more tables - postgres/src/include/nodes/parsenodes.h:2075
	ReturningList *NodeList   // List of expressions to return - postgres/src/include/nodes/parsenodes.h:2076
	WithClause    *WithClause // WITH clause - postgres/src/include/nodes/parsenodes.h:2077
}

// NewUpdateStmt creates a new UpdateStmt node.
func NewUpdateStmt(relation *RangeVar) *UpdateStmt { _ = "STUB: not implemented"; return nil }

func (u *UpdateStmt) String() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateStmt) StatementType() string {
	_ = "STUB: not implemented"

	// SqlString returns the SQL representation of the UpdateStmt
	return ""
}

func (u *UpdateStmt) SqlString() string {
	_ = "STUB: not implemented"

	// WITH clause
	return ""
}

// UPDATE table

// SET clause

// FROM clause

// WHERE clause

// RETURNING clause

// DeleteStmt represents a DELETE statement.
// Ported from postgres/src/include/nodes/parsenodes.h:2055
type DeleteStmt struct {
	BaseNode
	Relation      *RangeVar   // Relation to delete from - postgres/src/include/nodes/parsenodes.h:2058
	UsingClause   *NodeList   // Optional using clause for more tables - postgres/src/include/nodes/parsenodes.h:2059
	WhereClause   Node        // Qualifications - postgres/src/include/nodes/parsenodes.h:2060
	ReturningList *NodeList   // List of expressions to return - postgres/src/include/nodes/parsenodes.h:2061
	WithClause    *WithClause // WITH clause - postgres/src/include/nodes/parsenodes.h:2062
}

// NewDeleteStmt creates a new DeleteStmt node.
func NewDeleteStmt(relation *RangeVar) *DeleteStmt { _ = "STUB: not implemented"; return nil }

func (d *DeleteStmt) String() string { _ = "STUB: not implemented"; return "" }

func (d *DeleteStmt) StatementType() string {
	_ = "STUB: not implemented"

	// SqlString returns the SQL representation of the DeleteStmt
	return ""
}

func (d *DeleteStmt) SqlString() string {
	_ = "STUB: not implemented"

	// WITH clause
	return ""
}

// DELETE FROM table

// USING clause

// WHERE clause

// RETURNING clause

// ==============================================================================
// DDL STATEMENTS
// ==============================================================================

// CreateStmt represents a CREATE TABLE statement.
// Ported from postgres/src/include/nodes/parsenodes.h:2648
type CreateStmt struct {
	BaseNode
	Relation       *RangeVar           // Relation to create - postgres/src/include/nodes/parsenodes.h:2651
	TableElts      *NodeList           // Column definitions - postgres/src/include/nodes/parsenodes.h:2652
	InhRelations   *NodeList           // Relations to inherit from - postgres/src/include/nodes/parsenodes.h:2653
	PartBound      *PartitionBoundSpec // FOR VALUES clause - parsenodes.h
	PartSpec       *PartitionSpec      // PARTITION BY clause - parsenodes.h
	OfTypename     *TypeName           // OF typename clause - parsenodes.h
	Constraints    []*Constraint       // Constraints - postgres/src/include/nodes/parsenodes.h:2659
	Options        *NodeList           // Options from WITH clause - postgres/src/include/nodes/parsenodes.h:2660
	OnCommit       OnCommitAction      // OnCommitAction for temp tables - parsenodes.h
	TableSpaceName string              // Table space to use, or empty - postgres/src/include/nodes/parsenodes.h:2662
	AccessMethod   string              // Table access method - postgres/src/include/nodes/parsenodes.h:2663
	IfNotExists    bool                // Just do nothing if it already exists? - postgres/src/include/nodes/parsenodes.h:2664
}

// NewCreateStmt creates a new CreateStmt node.
func NewCreateStmt(relation *RangeVar) *CreateStmt { _ = "STUB: not implemented"; return nil }

func (c *CreateStmt) String() string { _ = "STUB: not implemented"; return "" }

func (c *CreateStmt) StatementType() string {
	_ = "STUB: not implemented"

	// DropStmt represents a DROP statement.
	// Ported from postgres/src/include/nodes/parsenodes.h:3226
	return ""
}

type DropStmt struct {
	BaseNode
	Objects    *NodeList    // List of names - postgres/src/include/nodes/parsenodes.h:3229
	RemoveType ObjectType   // Object type - postgres/src/include/nodes/parsenodes.h:3230
	Behavior   DropBehavior // RESTRICT or CASCADE behavior - postgres/src/include/nodes/parsenodes.h:3231
	MissingOk  bool         // Skip error if object is missing? - postgres/src/include/nodes/parsenodes.h:3232
	Concurrent bool         // Drop index concurrently? - postgres/src/include/nodes/parsenodes.h:3233
}

// NewDropStmt creates a new DropStmt node.
func NewDropStmt(objects *NodeList, removeType ObjectType) *DropStmt {
	_ = "STUB: not implemented"
	return nil
}

func (d *DropStmt) String() string { _ = "STUB: not implemented"; return "" }

func (d *DropStmt) StatementType() string {
	_ = "STUB: not implemented"

	// ==============================================================================
	// COLUMN REFERENCES
	// ==============================================================================
	return ""
}

// ColumnRef represents a column reference in expressions.
// Ported from postgres/src/include/nodes/parsenodes.h:291
type ColumnRef struct {
	BaseNode
	Fields *NodeList // List of field names - postgres/src/include/nodes/parsenodes.h:292
}

// NewColumnRef creates a new ColumnRef node.
func NewColumnRef(fields ...Node) *ColumnRef { _ = "STUB: not implemented"; return nil }

func (c *ColumnRef) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of the ColumnRef
func (c *ColumnRef) SqlString() string { _ = "STUB: not implemented"; return "" }

// Handle different field types (String for column names, A_Star for *, A_Indices for array access)

// Array access - format as [index] or [start:end]

// For other field types, all nodes implement SqlString()

// For simple column references, join with dots
// For complex ones with array access, concatenate appropriately

// Array access - no dot separator

// Regular field access - use dot separator

func (c *ColumnRef) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (c *ColumnRef) IsExpr() bool {
	_ = "STUB: not implemented"

	// ==============================================================================
	// PLACEHOLDER TYPES - To be fully implemented later
	// ==============================================================================
	return false
}

// These types are referenced by the main statements above and are now fully implemented.
// Implementation moved to query_execution_nodes.go for better organization.

// CommonTableExpr represents a WITH clause (Common Table Expression / CTE).
// CTEs are increasingly common in modern SQL and enable recursive queries
// and improved query organization.
// Ported from postgres/src/include/nodes/parsenodes.h:1668
type CommonTableExpr struct {
	BaseNode
	Ctename          string           // Query name (never qualified) - parsenodes.h:1676
	Aliascolnames    *NodeList        // Optional list of column names - parsenodes.h:1678
	Ctematerialized  CTEMaterialized  // Is this an optimization fence? - parsenodes.h:1679
	Ctequery         Node             // The CTE's subquery - parsenodes.h:1681
	SearchClause     *CTESearchClause // SEARCH clause, if any - parsenodes.h:1682
	CycleClause      *CTECycleClause  // CYCLE clause, if any - parsenodes.h:1683
	Cterecursive     bool             // Is this a recursive CTE? - parsenodes.h:1687
	Cterefcount      int              // Number of RTEs referencing this CTE - parsenodes.h:1693
	Ctecolnames      *NodeList        // List of output column names - parsenodes.h:1696
	Ctecoltypes      []Oid            // OID list of output column type OIDs - parsenodes.h:1697
	Ctecoltypmods    []int32          // Integer list of output column typmods - parsenodes.h:1698
	Ctecolcollations []Oid            // OID list of column collation OIDs - parsenodes.h:1699
}

// CTEMaterialized represents CTE materialization settings.
// Ported from postgres/src/include/nodes/parsenodes.h:1636
type CTEMaterialized int

const (
	CTEMaterializeDefault CTEMaterialized = iota // No materialization clause - parsenodes.h:1638
	CTEMaterializeAlways                         // MATERIALIZED - parsenodes.h:1639
	CTEMaterializeNever                          // NOT MATERIALIZED - parsenodes.h:1640
)

// CTESearchClause represents a SEARCH clause in recursive CTEs.
// Ported from postgres/src/include/nodes/parsenodes.h:1643
type CTESearchClause struct {
	BaseNode
	SearchColList      *NodeList // List of columns to search - parsenodes.h:1646
	SearchBreadthFirst bool      // True for BREADTH FIRST, false for DEPTH FIRST - parsenodes.h:1647
	SearchSeqColumn    string    // Name of column to set search sequence - parsenodes.h:1648
}

// CTECycleClause represents a CYCLE clause in recursive CTEs.
// Ported from postgres/src/include/nodes/parsenodes.h:1652
type CTECycleClause struct {
	BaseNode
	CycleColList       *NodeList  // List of columns to check for cycles - parsenodes.h:1655
	CycleMarkColumn    string     // Name of column to mark cycles - parsenodes.h:1656
	CycleMarkValue     Expression // Value to set when cycle detected - parsenodes.h:1657
	CycleMarkDefault   Expression // Value to set when no cycle - parsenodes.h:1658
	CyclePathColumn    string     // Name of column to track path - parsenodes.h:1659
	CycleMarkType      Oid        // Common type of mark_value and mark_default - parsenodes.h:1662
	CycleMarkTypmod    int32      // Type modifier - parsenodes.h:1663
	CycleMarkCollation Oid        // Collation - parsenodes.h:1664
	CycleMarkNeop      Oid        // <> operator for type - parsenodes.h:1665
}

// NewCTESearchClause creates a new CTESearchClause node.
func NewCTESearchClause(searchColList *NodeList, breadthFirst bool, seqColumn string) *CTESearchClause {
	_ = "STUB: not implemented"
	return nil
}

// SqlString returns the SQL representation of the CTESearchClause.
func (sc *CTESearchClause) SqlString() string { _ = "STUB: not implemented"; return "" }

// NewCTECycleClause creates a new CTECycleClause node.
func NewCTECycleClause(cycleColList *NodeList, markColumn string, markValue, markDefault Expression, pathColumn string) *CTECycleClause {
	_ = "STUB: not implemented"
	return nil
}

// SqlString returns the SQL representation of the CTECycleClause.
func (cc *CTECycleClause) SqlString() string { _ = "STUB: not implemented"; return "" }

// NewCommonTableExpr creates a new CommonTableExpr node.
func NewCommonTableExpr(ctename string, ctequery Node) *CommonTableExpr {
	_ = "STUB: not implemented"
	return nil
}

// NewRecursiveCommonTableExpr creates a new recursive CommonTableExpr node.
func NewRecursiveCommonTableExpr(ctename string, ctequery Node) *CommonTableExpr {
	_ = "STUB: not implemented"
	return nil
}

func (cte *CommonTableExpr) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of the CommonTableExpr
func (c *CommonTableExpr) SqlString() string { _ = "STUB: not implemented"; return "" }

// Add column names if specified

// Add the CTE query

// Determine if materialized/not materialized

// Add the actual query (usually in parentheses)

// Add SEARCH clause if present

// Add CYCLE clause if present

// Placeholder structs for other query execution nodes implemented in query_execution_nodes.go
// IntoClause placeholder removed - now implemented in expressions.go
// SetOperation represents the type of set operation
// Ported from postgres/src/include/nodes/parsenodes.h:2108-2114
type SetOperation int

const (
	SETOP_NONE      SetOperation = iota // No set operation
	SETOP_UNION                         // UNION
	SETOP_INTERSECT                     // INTERSECT
	SETOP_EXCEPT                        // EXCEPT
)

func (s SetOperation) String() string { _ = "STUB: not implemented"; return "" }

// OnConflictClause represents ON CONFLICT clause for INSERT statements
// Ported from postgres/src/include/nodes/parsenodes.h:1621-1629
type OnConflictClause struct {
	BaseNode
	Action      OnConflictAction `json:"action"`      // DO NOTHING or UPDATE?
	Infer       *InferClause     `json:"infer"`       // Optional index inference clause
	TargetList  *NodeList        `json:"targetList"`  // The target list (of ResTarget)
	WhereClause Node             `json:"whereClause"` // Qualifications
}

func (n *OnConflictClause) node() { _ = "STUB: not implemented"; return }

func (n *OnConflictClause) String() string { _ = "STUB: not implemented"; return "" }

func (n *OnConflictClause) SqlString() string { _ = "STUB: not implemented"; return "" }

// NewOnConflictClause creates a new OnConflictClause node
func NewOnConflictClause(action OnConflictAction) *OnConflictClause {
	_ = "STUB: not implemented"
	return nil
}

// OverridingKind represents OVERRIDING clause options
// Ported from postgres/src/include/nodes/primnodes.h:25-30
type OverridingKind int

const (
	OVERRIDING_NOT_SET      OverridingKind = iota // No OVERRIDING clause
	OVERRIDING_USER_VALUE                         // OVERRIDING USER VALUE
	OVERRIDING_SYSTEM_VALUE                       // OVERRIDING SYSTEM VALUE
)

func (o OverridingKind) String() string { _ = "STUB: not implemented"; return "" }

func (o OverridingKind) SqlString() string { _ = "STUB: not implemented"; return "" }

// Note: Constraint is now defined in ddl_statements.go

// SqlString returns the SQL representation of CREATE TABLE statement
func (c *CreateStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// Add TEMPORARY if specified

// Add IF NOT EXISTS if specified

// Add table name

// Handle different table types with correct ordering

// For partition tables: PARTITION OF parent [constraints] FOR VALUES [PARTITION BY]

// Add constraints for partition tables (but only if there are any)

// Add FOR VALUES clause

// Add PARTITION BY for default partition tables

// For typed tables: OF typename [ ( column_options | table_constraint, ... ) ].
// A typed-table column is a ColumnDef with no type (just constraints); the
// optional WITH OPTIONS keyword is syntactic sugar (identical AST), so the
// plain `colname <constraints>` rendering round-trips.

// Regular table with columns and constraints

// Add table-level constraints

// Add INHERITS clause for regular inheritance

// Add PARTITION BY clause for regular partitioned tables

// Add USING clause if specified (for table access method)

// Add WITH options if specified

// Add ON COMMIT clause if specified

// ONCOMMIT_PRESERVE_ROWS

// ONCOMMIT_DELETE_ROWS

// ONCOMMIT_DROP

// Add tablespace if specified

// SqlString returns the SQL representation of DROP statement
func (d *DropStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// Handle special cases first

// Add object type - always include it as all valid ObjectTypes should have string representations

// Add CONCURRENTLY if specified for indexes

// Add IF EXISTS if specified

// Add object names

// Handle functions, aggregates, operators with arguments

// Handle qualified names (schema.table)

// Handle type names (for DROP TYPE, etc.)

// Add CASCADE/RESTRICT behavior

// Note: We don't output RESTRICT as it's the default behavior in PostgreSQL
// Only CASCADE needs to be explicitly specified

// sqlStringForDropCast handles DROP CAST (source_type AS target_type)
func (d *DropStmt) sqlStringForDropCast() string { _ = "STUB: not implemented"; return "" }

// Objects should contain a single NodeList with [source_type, target_type]

// sqlStringForDropOpClass handles DROP OPERATOR CLASS name USING access_method
func (d *DropStmt) sqlStringForDropOpClass() string { _ = "STUB: not implemented"; return "" }

// Objects should contain a single NodeList with [access_method, ...names]

// First item is access method, rest are qualified name parts

// First item is access method

// Rest are name parts

// Add qualified name

// Add USING access_method

// sqlStringForDropOpFamily handles DROP OPERATOR FAMILY name USING access_method
func (d *DropStmt) sqlStringForDropOpFamily() string { _ = "STUB: not implemented"; return "" }

// Same logic as DROP OPERATOR CLASS

// First item is access method

// Rest are name parts

// Add qualified name

// Add USING access_method

// sqlStringForDropTransform handles DROP TRANSFORM FOR type LANGUAGE lang
func (d *DropStmt) sqlStringForDropTransform() string { _ = "STUB: not implemented"; return "" }

// Objects should contain a single NodeList with [type, language_name]

// sqlStringForDropSubscription handles DROP SUBSCRIPTION name
func (d *DropStmt) sqlStringForDropSubscription() string { _ = "STUB: not implemented"; return "" }

// Objects should contain a single NodeList with subscription name

// sqlStringForDropOnTable handles DROP RULE/TRIGGER/POLICY name ON table
func (d *DropStmt) sqlStringForDropOnTable() string { _ = "STUB: not implemented"; return "" }

// Add object type

// Add IF EXISTS if specified

// For RULE/TRIGGER/POLICY the object is one flat qualified name whose last
// part is the rule/trigger/policy name and whose preceding parts form the
// (optionally schema-qualified) table: DROP TRIGGER <name> ON <table>.

// Add CASCADE/RESTRICT behavior

// Note: We don't output RESTRICT as it's the default behavior in PostgreSQL
