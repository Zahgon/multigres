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
// ADVANCED ALTER TABLE OPERATIONS
// ==============================================================================

// Using existing AlterTableType and AlterTableCmd from ddl_statements.go
// Enhanced with additional constructor functions for advanced ALTER TABLE operations

// ==============================================================================
// TABLE DEFINITION NODES
// ==============================================================================

// ColumnDef is already defined as placeholder in statements.go
// Will be enhanced separately to replace the placeholder

// TableLikeClause represents LIKE clauses in CREATE TABLE.
// This allows table inheritance and copying of structure.
// Ported from postgres/src/include/nodes/parsenodes.h:751
type TableLikeClause struct {
	BaseNode
	Relation *RangeVar       // Table to copy from - parsenodes.h:752
	Options  TableLikeOption // OR of TableLikeOption flags - parsenodes.h:753
}

// TableLikeOption represents options for LIKE clauses.
// Ported from postgres/src/include/nodes/parsenodes.h:755-763
type TableLikeOption int

const (
	CREATE_TABLE_LIKE_COMMENTS    TableLikeOption = 1 << 0 // INCLUDING COMMENTS - parsenodes.h:756
	CREATE_TABLE_LIKE_COMPRESSION TableLikeOption = 1 << 1 // INCLUDING COMPRESSION - parsenodes.h:757
	CREATE_TABLE_LIKE_CONSTRAINTS TableLikeOption = 1 << 2 // INCLUDING CONSTRAINTS - parsenodes.h:758
	CREATE_TABLE_LIKE_DEFAULTS    TableLikeOption = 1 << 3 // INCLUDING DEFAULTS - parsenodes.h:759
	CREATE_TABLE_LIKE_GENERATED   TableLikeOption = 1 << 4 // INCLUDING GENERATED - parsenodes.h:760
	CREATE_TABLE_LIKE_IDENTITY    TableLikeOption = 1 << 5 // INCLUDING IDENTITY - parsenodes.h:761
	CREATE_TABLE_LIKE_INDEXES     TableLikeOption = 1 << 6 // INCLUDING INDEXES - parsenodes.h:762
	CREATE_TABLE_LIKE_STATISTICS  TableLikeOption = 1 << 7 // INCLUDING STATISTICS - parsenodes.h:763
	CREATE_TABLE_LIKE_STORAGE     TableLikeOption = 1 << 8 // INCLUDING STORAGE - parsenodes.h:764
	CREATE_TABLE_LIKE_ALL         TableLikeOption = 1 << 9 // INCLUDING ALL - parsenodes.h:765
)

// NewTableLikeClause creates a new TableLikeClause node.
func NewTableLikeClause(relation *RangeVar, options TableLikeOption) *TableLikeClause {
	_ = "STUB: not implemented"
	return nil
}

// NewTableLikeAll creates a LIKE clause including all options.
func NewTableLikeAll(relation *RangeVar) *TableLikeClause { _ = "STUB: not implemented"; return nil }

func (tlc *TableLikeClause) String() string { _ = "STUB: not implemented"; return "" }

// SqlString generates the SQL representation of the TableLikeClause
func (tlc *TableLikeClause) SqlString() string { _ = "STUB: not implemented"; return "" }

// Handle INCLUDING options

// ==============================================================================
// PARTITIONING SUPPORT
// ==============================================================================

// PartitionStrategy represents partitioning strategies.
// Ported from postgres/src/include/nodes/parsenodes.h:865-869
type PartitionStrategy string

const (
	PARTITION_STRATEGY_LIST  PartitionStrategy = "list"  // LIST partitioning - parsenodes.h:866
	PARTITION_STRATEGY_RANGE PartitionStrategy = "range" // RANGE partitioning - parsenodes.h:867
	PARTITION_STRATEGY_HASH  PartitionStrategy = "hash"  // HASH partitioning - parsenodes.h:868
)

// PartitionSpec represents table partitioning specifications.
// Modern PostgreSQL feature for table partitioning.
// Ported from postgres/src/include/nodes/parsenodes.h:882
type PartitionSpec struct {
	BaseNode
	Strategy   PartitionStrategy // Partitioning strategy - parsenodes.h:884
	PartParams *NodeList         // List of PartitionElem nodes - parsenodes.h:885
	// Location is provided by BaseNode.Location() method
}

// NewPartitionSpec creates a new PartitionSpec node.
func NewPartitionSpec(strategy PartitionStrategy, partParams *NodeList) *PartitionSpec {
	_ = "STUB: not implemented"
	return nil
}

// NewListPartitionSpec creates a new LIST partition specification.
func NewListPartitionSpec(partParams *NodeList) *PartitionSpec {
	_ = "STUB: not implemented"
	return nil
}

// NewRangePartitionSpec creates a new RANGE partition specification.
func NewRangePartitionSpec(partParams *NodeList) *PartitionSpec {
	_ = "STUB: not implemented"
	return nil
}

// NewHashPartitionSpec creates a new HASH partition specification.
func NewHashPartitionSpec(partParams *NodeList) *PartitionSpec {
	_ = "STUB: not implemented"
	return nil
}

func (ps *PartitionSpec) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of the PartitionSpec
func (ps *PartitionSpec) SqlString() string { _ = "STUB: not implemented"; return "" }

// PartitionBoundSpec represents partition boundary specifications.
// This defines the actual bounds for individual partitions.
// Ported from postgres/src/include/nodes/parsenodes.h:896
type PartitionBoundSpec struct {
	BaseNode
	Strategy   PartitionStrategy // Partitioning strategy - parsenodes.h:898
	IsDefault  bool              // Is this a default partition? - parsenodes.h:899
	Modulus    int               // Hash partition modulus - parsenodes.h:900
	Remainder  int               // Hash partition remainder - parsenodes.h:901
	ListDatums *NodeList         // List of list datums per column - parsenodes.h:902
	LowDatums  *NodeList         // List of lower datums for range bounds - parsenodes.h:903
	HighDatums *NodeList         // List of upper datums for range bounds - parsenodes.h:904
	// Location is provided by BaseNode.Location() method
}

// NewPartitionBoundSpec creates a new PartitionBoundSpec node.
func NewPartitionBoundSpec(strategy PartitionStrategy) *PartitionBoundSpec {
	_ = "STUB: not implemented"
	return nil
}

// NewDefaultPartitionBound creates a new default partition bound.
func NewDefaultPartitionBound() *PartitionBoundSpec { _ = "STUB: not implemented"; return nil }

// NewHashPartitionBound creates a new hash partition bound.
func NewHashPartitionBound(modulus, remainder int) *PartitionBoundSpec {
	_ = "STUB: not implemented"
	return nil
}

// NewListPartitionBound creates a new list partition bound.
func NewListPartitionBound(listDatums *NodeList) *PartitionBoundSpec {
	_ = "STUB: not implemented"
	return nil
}

// NewRangePartitionBound creates a new range partition bound.
func NewRangePartitionBound(lowDatums, highDatums *NodeList) *PartitionBoundSpec {
	_ = "STUB: not implemented"
	return nil
}

func (pbs *PartitionBoundSpec) String() string { _ = "STUB: not implemented"; return "" }

func (pbs *PartitionBoundSpec) SqlString() string { _ = "STUB: not implemented"; return "" }

// Handle single datum (not nested in another NodeList)

// PartitionRangeDatum represents partition range datum values.
// Ported from postgres/src/include/nodes/parsenodes.h:929
type PartitionRangeDatum struct {
	BaseNode
	Kind  PartitionRangeDatumKind // What kind of datum this is - parsenodes.h:931
	Value Node                    // The actual datum value - parsenodes.h:932
	// Location is provided by BaseNode.Location() method
}

// PartitionRangeDatumKind represents types of range datums.
// Ported from postgres/src/include/nodes/parsenodes.h:935-939
type PartitionRangeDatumKind int

const (
	PARTITION_RANGE_DATUM_MINVALUE PartitionRangeDatumKind = iota // MINVALUE - parsenodes.h:936
	PARTITION_RANGE_DATUM_VALUE                                   // Specific value - parsenodes.h:937
	PARTITION_RANGE_DATUM_MAXVALUE                                // MAXVALUE - parsenodes.h:938
)

// NewPartitionRangeDatum creates a new PartitionRangeDatum node.
func NewPartitionRangeDatum(kind PartitionRangeDatumKind, value Node) *PartitionRangeDatum {
	_ = "STUB: not implemented"
	return nil
}

// NewMinValueDatum creates a MINVALUE range datum.
func NewMinValueDatum() *PartitionRangeDatum { _ = "STUB: not implemented"; return nil }

// NewMaxValueDatum creates a MAXVALUE range datum.
func NewMaxValueDatum() *PartitionRangeDatum { _ = "STUB: not implemented"; return nil }

// NewValueDatum creates a specific value range datum.
func NewValueDatum(value Node) *PartitionRangeDatum { _ = "STUB: not implemented"; return nil }

func (prd *PartitionRangeDatum) String() string { _ = "STUB: not implemented"; return "" }

// ==============================================================================
// STATISTICS SUPPORT
// ==============================================================================

// StatsElem represents statistics element specifications.
// This is used for extended statistics in PostgreSQL.
// Ported from postgres/src/include/nodes/parsenodes.h:3403
type StatsElem struct {
	BaseNode
	Name string // Name of attribute to compute stats for - parsenodes.h:3405
	Expr Node   // Or expression to compute stats for - parsenodes.h:3406
}

// NewStatsElem creates a new StatsElem node.
func NewStatsElem(name string) *StatsElem { _ = "STUB: not implemented"; return nil }

// NewStatsElemExpr creates a new StatsElem with expression.
func NewStatsElemExpr(expr Node) *StatsElem { _ = "STUB: not implemented"; return nil }

func (se *StatsElem) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of StatsElem
func (se *StatsElem) SqlString() string { _ = "STUB: not implemented"; return "" }

// ==============================================================================
// FOREIGN DATA WRAPPER STATEMENTS
// ==============================================================================

// CreateForeignServerStmt represents CREATE FOREIGN SERVER statements.
// This supports PostgreSQL's foreign data wrapper functionality.
// Ported from postgres/src/include/nodes/parsenodes.h:2870
type CreateForeignServerStmt struct {
	BaseNode
	Servername  string    // Server name - parsenodes.h:2872
	Servertype  string    // Optional server type - parsenodes.h:2873
	Version     string    // Optional server version - parsenodes.h:2874
	Fdwname     string    // FDW name - parsenodes.h:2875
	IfNotExists bool      // IF NOT EXISTS clause - parsenodes.h:2876
	Options     *NodeList // Generic options to FDW - parsenodes.h:2877
}

// NewCreateForeignServerStmt creates a new CreateForeignServerStmt node.
// NewCreateForeignServerStmt creates a new CreateForeignServerStmt node.
func NewCreateForeignServerStmt(servername, servertype, version, fdwname string, options *NodeList, ifNotExists bool) *CreateForeignServerStmt {
	_ = "STUB: not implemented"
	return nil
}

func (cfss *CreateForeignServerStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

func (cfss *CreateForeignServerStmt) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of the CREATE FOREIGN SERVER statement
func (cfss *CreateForeignServerStmt) SqlString() string {
	_ = "STUB: not implemented"

	// CREATE SERVER [IF NOT EXISTS]
	return ""
}

// server name

// TYPE clause

// VERSION clause

// FOREIGN DATA WRAPPER

// OPTIONS clause

// For generic options, use PostgreSQL format: key 'value' (no =)

// CreateForeignTableStmt represents CREATE FOREIGN TABLE statements.
// This creates tables that reference external data sources.
// Ported from postgres/src/include/nodes/parsenodes.h:2895
type CreateForeignTableStmt struct {
	BaseNode
	Base       *CreateStmt // Base CREATE TABLE statement - parsenodes.h:2897
	Servername string      // Foreign server name - parsenodes.h:2898
	Options    *NodeList   // OPTIONS clause - parsenodes.h:2899
}

// NewCreateForeignTableStmt creates a new CreateForeignTableStmt node.
// NewCreateForeignTableStmt creates a CreateForeignTableStmt from grammar components.
func NewCreateForeignTableStmt(relation *RangeVar, tableElts, inhRelations *NodeList, servername string, options *NodeList, partBound *PartitionBoundSpec, ifNotExists bool) *CreateForeignTableStmt {
	_ = "STUB: not implemented"
	// Create the base CreateStmt
	return nil
}

func (cfts *CreateForeignTableStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

func (cfts *CreateForeignTableStmt) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of the CREATE FOREIGN TABLE statement
func (cfts *CreateForeignTableStmt) SqlString() string {
	_ = "STUB: not implemented"

	// CREATE FOREIGN TABLE [IF NOT EXISTS]
	return ""
}

// table name

// column definitions - add parentheses unless this is a partition table

// Regular foreign table - always include parentheses

// Partition table with column definitions (rare case)

// PARTITION OF or INHERITS clause

// PARTITION OF clause for partition tables

// Add partition bound specification

// Regular INHERITS clause

// SERVER clause

// OPTIONS clause

// For generic options, use PostgreSQL format: key 'value' (no =)

// CreateUserMappingStmt represents CREATE USER MAPPING statements.
// This maps database users to foreign server users.
// Ported from postgres/src/include/nodes/parsenodes.h:2907
type CreateUserMappingStmt struct {
	BaseNode
	User        *RoleSpec // User role - parsenodes.h:2909
	Servername  string    // Foreign server name - parsenodes.h:2910
	IfNotExists bool      // IF NOT EXISTS clause - parsenodes.h:2911
	Options     *NodeList // Generic options to FDW - parsenodes.h:2912
}

// NewCreateUserMappingStmt creates a new CreateUserMappingStmt node.
// NewCreateUserMappingStmt creates a new CreateUserMappingStmt node.
func NewCreateUserMappingStmt(user *RoleSpec, servername string, options *NodeList, ifNotExists bool) *CreateUserMappingStmt {
	_ = "STUB: not implemented"
	return nil
}

func (cums *CreateUserMappingStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

func (cums *CreateUserMappingStmt) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of the CREATE USER MAPPING statement
func (cums *CreateUserMappingStmt) SqlString() string {
	_ = "STUB: not implemented"

	// CREATE USER MAPPING [IF NOT EXISTS]
	return ""
}

// FOR user

// SERVER name

// OPTIONS clause

// For generic options, use PostgreSQL format: key 'value' (no =)

// ==============================================================================
// TRIGGER STATEMENTS
// ==============================================================================

// TriggerTransition represents trigger transition tables.
// Ported from postgres/src/include/nodes/parsenodes.h:1737
type TriggerTransition struct {
	BaseNode
	Name    string // Transition table name - parsenodes.h:1739
	IsNew   bool   // Is this NEW table? (or OLD table?) - parsenodes.h:1740
	IsTable bool   // Is this a table? (or row?) - parsenodes.h:1741
}

// CreateTriggerStmt represents CREATE TRIGGER statements.
// Triggers are essential for PostgreSQL's event system.
// Ported from postgres/src/include/nodes/parsenodes.h:3001
type CreateTriggerStmt struct {
	BaseNode
	Replace      bool      // Replace existing trigger? - parsenodes.h:3003
	IsConstraint bool      // Is this a constraint trigger? - parsenodes.h:3004
	Trigname     string    // Trigger name - parsenodes.h:3005
	Relation     *RangeVar // Relation trigger is on - parsenodes.h:3006
	Funcname     *NodeList // Qual. name of function to call - parsenodes.h:3007
	Args         *NodeList // List of (T_String) Values or NIL - parsenodes.h:3008
	Row          bool      // ROW/STATEMENT - parsenodes.h:3009
	Timing       int16     // BEFORE, AFTER, or INSTEAD - parsenodes.h:3010
	Events       int16     // "OR" of INSERT/UPDATE/DELETE/TRUNCATE - parsenodes.h:3011
	Columns      *NodeList // Column names, or NIL for all columns - parsenodes.h:3012
	WhenClause   Node      // WHEN clause - parsenodes.h:3013
	Constrrel    *RangeVar // Opposite relation, if RI trigger - parsenodes.h:3014
	Deferrable   bool      // DEFERRABLE - parsenodes.h:3015
	Initdeferred bool      // INITIALLY DEFERRED - parsenodes.h:3016
	Transitions  *NodeList // Transition table clauses - parsenodes.h:3017
}

// TriggerType represents trigger event types.
const (
	TRIGGER_TYPE_INSERT   = 1 << 0 // INSERT trigger
	TRIGGER_TYPE_UPDATE   = 1 << 1 // UPDATE trigger
	TRIGGER_TYPE_DELETE   = 1 << 2 // DELETE trigger
	TRIGGER_TYPE_TRUNCATE = 1 << 3 // TRUNCATE trigger
)

// TriggerTiming represents trigger timing.
const (
	TRIGGER_TIMING_BEFORE  = 1 // BEFORE trigger
	TRIGGER_TIMING_AFTER   = 2 // AFTER trigger
	TRIGGER_TIMING_INSTEAD = 3 // INSTEAD OF trigger
)

// NewCreateTriggerStmt creates a new CreateTriggerStmt node.
func NewCreateTriggerStmt(trigname string, relation *RangeVar, funcname *NodeList, timing int16, events int16) *CreateTriggerStmt {
	_ = "STUB: not implemented"
	return nil
}

// Default to ROW trigger

// NewBeforeInsertTrigger creates a BEFORE INSERT trigger.
func NewBeforeInsertTrigger(trigname string, relation *RangeVar, funcname *NodeList) *CreateTriggerStmt {
	_ = "STUB: not implemented"
	return nil
}

// NewAfterUpdateTrigger creates an AFTER UPDATE trigger.
func NewAfterUpdateTrigger(trigname string, relation *RangeVar, funcname *NodeList) *CreateTriggerStmt {
	_ = "STUB: not implemented"
	return nil
}

// NewConstraintTrigger creates a constraint trigger.
func NewConstraintTrigger(trigname string, relation *RangeVar, funcname *NodeList, timing int16, events int16, constrrel *RangeVar) *CreateTriggerStmt {
	_ = "STUB: not implemented"
	return nil
}

// Default to NOT DEFERRABLE
// Default to INITIALLY IMMEDIATE

// NewDeferrableConstraintTrigger creates a deferrable constraint trigger.
func NewDeferrableConstraintTrigger(trigname string, relation *RangeVar, funcname *NodeList, timing int16, events int16, constrrel *RangeVar, initdeferred bool) *CreateTriggerStmt {
	_ = "STUB: not implemented"
	return nil
}

func (cts *CreateTriggerStmt) String() string { _ = "STUB: not implemented"; return "" }

// Build the base string

// Add constraint trigger information

// Add transition information

// Location returns the location for this node
func (cts *CreateTriggerStmt) Location() int {
	_ = "STUB: not implemented"
	// TODO: Implement proper location tracking
	return 0
}

// NodeTag returns the node's type tag
func (cts *CreateTriggerStmt) NodeTag() NodeTag { _ = "STUB: not implemented"; return *new(NodeTag) }

// StatementType returns the statement type for this node
func (cts *CreateTriggerStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns SQL representation of the CREATE TRIGGER statement
func (cts *CreateTriggerStmt) SqlString() string {
	_ = "STUB: not implemented"

	// CREATE [OR REPLACE] [CONSTRAINT] TRIGGER
	return ""
}

// Timing (BEFORE, AFTER, INSTEAD OF)

// Events (INSERT, UPDATE, DELETE, TRUNCATE)

// Add column list for UPDATE OF if present

// ON table

// FROM constraint_table (for constraint triggers)

// Deferrable options (for constraint triggers)

// REFERENCING transitions

// FOR EACH ROW/STATEMENT (only output if ROW is true, STATEMENT is default)

// WHEN clause

// EXECUTE FUNCTION

// Function name with arguments

// Function arguments

// Check if it's a numeric string or needs quotes

// Not a number, add quotes

// It's a number

// Combine function name and arguments without extra space

// ==============================================================================
// POLICY STATEMENTS (ROW LEVEL SECURITY)
// ==============================================================================

// CreatePolicyStmt represents CREATE POLICY statements.
// This implements PostgreSQL's row-level security policies.
// Ported from postgres/src/include/nodes/parsenodes.h:2959
type CreatePolicyStmt struct {
	BaseNode
	PolicyName string    // Policy name - parsenodes.h:2961
	Table      *RangeVar // Table the policy applies to - parsenodes.h:2962
	CmdName    string    // Command name (SELECT, INSERT, etc.) - parsenodes.h:2963
	Permissive bool      // Is this a permissive policy? - parsenodes.h:2964
	Roles      *NodeList // Roles policy applies to - parsenodes.h:2965
	Qual       Node      // USING clause - parsenodes.h:2966
	WithCheck  Node      // WITH CHECK clause - parsenodes.h:2967
}

// NewCreatePolicyStmt creates a new CreatePolicyStmt node.
func NewCreatePolicyStmt(policyName string, table *RangeVar, permissive bool, cmdName string, roles *NodeList, qual Node, withCheck Node) *CreatePolicyStmt {
	_ = "STUB: not implemented"
	return nil
}

func (cps *CreatePolicyStmt) String() string { _ = "STUB: not implemented"; return "" }

func (cps *CreatePolicyStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of CREATE POLICY statement
func (cps *CreatePolicyStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// Add AS clause for PERMISSIVE/RESTRICTIVE

// AlterPolicyStmt represents ALTER POLICY statements.
// Ported from postgres/src/include/nodes/parsenodes.h:2975
type AlterPolicyStmt struct {
	BaseNode
	PolicyName string    // Policy name - parsenodes.h:2977
	Table      *RangeVar // Table the policy applies to - parsenodes.h:2978
	Roles      *NodeList // Roles policy applies to - parsenodes.h:2979
	Qual       Node      // USING clause - parsenodes.h:2980
	WithCheck  Node      // WITH CHECK clause - parsenodes.h:2981
}

// NewAlterPolicyStmt creates a new AlterPolicyStmt node.
func NewAlterPolicyStmt(policyName string, table *RangeVar, roles *NodeList, qual Node, withCheck Node) *AlterPolicyStmt {
	_ = "STUB: not implemented"
	return nil
}

func (aps *AlterPolicyStmt) String() string { _ = "STUB: not implemented"; return "" }

func (aps *AlterPolicyStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of ALTER POLICY statement
func (aps *AlterPolicyStmt) SqlString() string { _ = "STUB: not implemented"; return "" }
