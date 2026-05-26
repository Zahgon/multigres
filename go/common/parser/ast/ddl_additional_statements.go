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

// =============================================================================
// TABLESPACE Statements
// =============================================================================

// CreateTableSpaceStmt represents CREATE TABLESPACE statement
type CreateTableSpaceStmt struct {
	BaseNode
	TablespaceName string    `json:"tablespacename"`
	Owner          *RoleSpec `json:"owner"`
	LocationPath   string    `json:"location"`
	Options        *NodeList `json:"options"`
}

func NewCreateTableSpaceStmt(name string, owner *RoleSpec, location string, options *NodeList) *CreateTableSpaceStmt {
	_ = "STUB: not implemented"
	return nil
}

func (c *CreateTableSpaceStmt) String() string { _ = "STUB: not implemented"; return "" }

func (c *CreateTableSpaceStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of CREATE TABLESPACE statement
func (c *CreateTableSpaceStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// Always include LOCATION clause since it's required in CREATE TABLESPACE syntax

// AlterTableSpaceStmt represents ALTER TABLESPACE statement
type AlterTableSpaceStmt struct {
	BaseNode
	TablespaceName string    `json:"tablespacename"`
	Options        *NodeList `json:"options"`
	IsReset        bool      `json:"isReset"`
}

func NewAlterTableSpaceStmt(name string, options *NodeList, isReset bool) *AlterTableSpaceStmt {
	_ = "STUB: not implemented"
	return nil
}

func (a *AlterTableSpaceStmt) String() string { _ = "STUB: not implemented"; return "" }

func (a *AlterTableSpaceStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of ALTER TABLESPACE statement
func (a *AlterTableSpaceStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// =============================================================================
// ACCESS METHOD Statements
// =============================================================================

// AmType represents access method type
type AmType byte

const (
	AMTYPE_INDEX AmType = 'i' // Index access method
	AMTYPE_TABLE AmType = 't' // Table access method
)

// CreateAmStmt represents CREATE ACCESS METHOD statement
type CreateAmStmt struct {
	BaseNode
	AmName      string    `json:"amname"`
	HandlerName *NodeList `json:"handler_name"`
	AmType      AmType    `json:"amtype"`
}

func NewCreateAmStmt(name string, amType AmType, handler *NodeList) *CreateAmStmt {
	_ = "STUB: not implemented"
	return nil
}

func (c *CreateAmStmt) String() string { _ = "STUB: not implemented"; return "" }

func (c *CreateAmStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of CREATE ACCESS METHOD statement
func (c *CreateAmStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// =============================================================================
// ALTER STATISTICS Statement
// =============================================================================

// AlterStatsStmt represents ALTER STATISTICS statement
type AlterStatsStmt struct {
	BaseNode
	DefNames      *NodeList `json:"defnames"`
	StxStatTarget Node      `json:"stxstattarget"`
	MissingOk     bool      `json:"missing_ok"`
}

func NewAlterStatsStmt(defnames *NodeList, target Node, missingOk bool) *AlterStatsStmt {
	_ = "STUB: not implemented"
	return nil
}

func (a *AlterStatsStmt) String() string { _ = "STUB: not implemented"; return "" }

func (a *AlterStatsStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of ALTER STATISTICS statement
func (a *AlterStatsStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// =============================================================================
// PUBLICATION Statements
// =============================================================================

// AlterPublicationType represents types of ALTER PUBLICATION operations
type AlterPublicationType int

const (
	AP_SetOptions AlterPublicationType = iota
	AP_AddObjects
	AP_SetObjects
	AP_DropObjects
)

// PublicationObjSpecType represents types of publication objects
type PublicationObjSpecType int

const (
	PUBLICATIONOBJ_TABLE PublicationObjSpecType = iota
	PUBLICATIONOBJ_TABLES_IN_SCHEMA
	PUBLICATIONOBJ_TABLES_IN_CUR_SCHEMA
	PUBLICATIONOBJ_CONTINUATION
)

// PublicationTable represents a publication table specification
type PublicationTable struct {
	BaseNode
	Relation    *RangeVar `json:"relation"`    // relation to be published
	WhereClause Node      `json:"whereclause"` // qualifications
	Columns     *NodeList `json:"columns"`     // List of columns in a publication table
}

func NewPublicationTable(relation *RangeVar, whereClause Node, columns *NodeList) *PublicationTable {
	_ = "STUB: not implemented"
	return nil
}

// PublicationObjSpec represents a publication object specification
// Ported from postgres/src/include/nodes/parsenodes.h:4151-4158
type PublicationObjSpec struct {
	BaseNode
	PubObjType PublicationObjSpecType `json:"pubobjtype"` // type of this publication object
	Name       string                 `json:"name"`       // object name
	PubTable   *PublicationTable      `json:"pubtable"`   // publication table details
}

// NewPublicationObjSpecName creates a PublicationObjSpec with just a name
func NewPublicationObjSpecName(objType PublicationObjSpecType, name string) *PublicationObjSpec {
	_ = "STUB: not implemented"
	return nil
}

// NewPublicationObjSpecTable creates a PublicationObjSpec with a PublicationTable
func NewPublicationObjSpecTable(objType PublicationObjSpecType, pubTable *PublicationTable) *PublicationObjSpec {
	_ = "STUB: not implemented"
	return nil
}

// NewPublicationObjSpec creates a basic PublicationObjSpec (for schema-only cases)
func NewPublicationObjSpec(objType PublicationObjSpecType) *PublicationObjSpec {
	_ = "STUB: not implemented"
	return nil
}

// preprocessPubObjList resolves PUBLICATIONOBJ_CONTINUATION entries — emitted
// by the grammar because LALR(1) cannot distinguish "TABLE x, y" from
// "TABLES IN SCHEMA a, b" until a later token — to the type of the preceding
// explicit entry. Mirrors preprocess_pubobj_list in postgres/src/backend/commands/publicationcmds.c.
func preprocessPubObjList(pubObjects *NodeList) { _ = "STUB: not implemented"; return }

func formatPubTable(pt *PublicationTable) string { _ = "STUB: not implemented"; return "" }

// formatPubObjList renders a publication object list, preserving the original
// order. A TABLE / TABLES IN SCHEMA keyword is emitted only when the object type
// changes; consecutive same-type objects share the preceding keyword (the
// inverse of preprocessPubObjList). Grouping the tables and schemas separately
// would reorder the list and change the parsed tree.
func formatPubObjList(pubObjects *NodeList) string { _ = "STUB: not implemented"; return "" }

// sentinel: forces a keyword on the first object

// A continuation entry can carry the name in PubTable instead, along
// with any column list / WHERE clause from the original text. Emit
// those through formatPubTable so they are not silently dropped (the
// combination is rejected by PostgreSQL at analysis time, but it still
// parses, so the round-trip must preserve it).

// A CONTINUATION that survives preprocessing is a first object with no
// preceding TABLE / TABLES IN SCHEMA keyword (e.g. `FOR CURRENT_SCHEMA`
// or a bare name). PostgreSQL rejects these at analysis time, but they
// parse, so emit the bare spelling to preserve the tree.

// CreatePublicationStmt represents CREATE PUBLICATION statement
type CreatePublicationStmt struct {
	BaseNode
	PubName      string    `json:"pubname"`
	PubObjects   *NodeList `json:"pubobjects"`
	ForAllTables bool      `json:"for_all_tables"`
	Options      *NodeList `json:"options"`
}

func NewCreatePublicationStmt(name string, objects *NodeList, forAllTables bool, options *NodeList) *CreatePublicationStmt {
	_ = "STUB: not implemented"
	return nil
}

func (c *CreatePublicationStmt) String() string { _ = "STUB: not implemented"; return "" }

func (c *CreatePublicationStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of CREATE PUBLICATION statement
func (c *CreatePublicationStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// AlterPublicationStmt represents ALTER PUBLICATION statement
type AlterPublicationStmt struct {
	BaseNode
	PubName    string               `json:"pubname"`
	Options    *NodeList            `json:"options"`
	PubObjects *NodeList            `json:"pubobjects"`
	Action     AlterPublicationType `json:"action"`
}

func NewAlterPublicationStmt(name string, options, objects *NodeList, action AlterPublicationType) *AlterPublicationStmt {
	_ = "STUB: not implemented"
	return nil
}

func (a *AlterPublicationStmt) String() string { _ = "STUB: not implemented"; return "" }

func (a *AlterPublicationStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of ALTER PUBLICATION statement
func (a *AlterPublicationStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// =============================================================================
// SUBSCRIPTION Statements
// =============================================================================

// AlterSubscriptionType represents types of ALTER SUBSCRIPTION operations
type AlterSubscriptionType int

const (
	ALTER_SUBSCRIPTION_OPTIONS AlterSubscriptionType = iota
	ALTER_SUBSCRIPTION_CONNECTION
	ALTER_SUBSCRIPTION_REFRESH
	ALTER_SUBSCRIPTION_ADD_PUBLICATION
	ALTER_SUBSCRIPTION_DROP_PUBLICATION
	ALTER_SUBSCRIPTION_SET_PUBLICATION
	ALTER_SUBSCRIPTION_ENABLED
	ALTER_SUBSCRIPTION_SKIP
)

// CreateSubscriptionStmt represents CREATE SUBSCRIPTION statement
type CreateSubscriptionStmt struct {
	BaseNode
	SubName     string    `json:"subname"`
	ConnInfo    string    `json:"conninfo"`
	Publication *NodeList `json:"publication"`
	Options     *NodeList `json:"options"`
}

func NewCreateSubscriptionStmt(name, connInfo string, publication, options *NodeList) *CreateSubscriptionStmt {
	_ = "STUB: not implemented"
	return nil
}

func (c *CreateSubscriptionStmt) String() string { _ = "STUB: not implemented"; return "" }

func (c *CreateSubscriptionStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of CREATE SUBSCRIPTION statement
func (c *CreateSubscriptionStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// Publication names should be identifiers, not quoted strings

// Use special formatting for subscription options

// formatSubscriptionOption formats DefElem for subscription options
func formatSubscriptionOption(d *DefElem) string { _ = "STUB: not implemented"; return "" }

// Special handling for specific subscription options

// For other string options, don't quote

// Fallback to regular SqlString for other types

// AlterSubscriptionStmt represents ALTER SUBSCRIPTION statement
type AlterSubscriptionStmt struct {
	BaseNode
	SubName     string                `json:"subname"`
	Kind        AlterSubscriptionType `json:"kind"`
	ConnInfo    string                `json:"conninfo"`
	Publication *NodeList             `json:"publication"`
	Options     *NodeList             `json:"options"`
}

func NewAlterSubscriptionStmt(name string, kind AlterSubscriptionType, connInfo string, publication, options *NodeList) *AlterSubscriptionStmt {
	_ = "STUB: not implemented"
	return nil
}

func (a *AlterSubscriptionStmt) String() string { _ = "STUB: not implemented"; return "" }

func (a *AlterSubscriptionStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of ALTER SUBSCRIPTION statement
func (a *AlterSubscriptionStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// The "enabled" option carries whether this is ENABLE or DISABLE.

// =============================================================================
// ALTER OPERATOR FAMILY Statement
// =============================================================================

// AlterOpFamilyStmt represents ALTER OPERATOR FAMILY statement
type AlterOpFamilyStmt struct {
	BaseNode
	OpFamilyName *NodeList `json:"opfamilyname"`
	AmName       string    `json:"amname"`
	IsDrop       bool      `json:"isDrop"`
	Items        *NodeList `json:"items"`
}

func NewAlterOpFamilyStmt(name *NodeList, amName string, isDrop bool, items *NodeList) *AlterOpFamilyStmt {
	_ = "STUB: not implemented"
	return nil
}

func (a *AlterOpFamilyStmt) String() string { _ = "STUB: not implemented"; return "" }

func (a *AlterOpFamilyStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of ALTER OPERATOR FAMILY statement
func (a *AlterOpFamilyStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// Add the items

// =============================================================================
// ALTER FUNCTION Statement
// =============================================================================

// AlterFunctionStmt represents ALTER FUNCTION statement
// Ported from postgres/src/include/nodes/parsenodes.h AlterFunctionStmt
type AlterFunctionStmt struct {
	BaseNode
	ObjType ObjectType      `json:"objtype"` // OBJECT_FUNCTION or OBJECT_PROCEDURE
	Func    *ObjectWithArgs `json:"func"`    // name and args of function
	Actions *NodeList       `json:"actions"` // list of DefElem
}

func NewAlterFunctionStmt(objType ObjectType, funcWithArgs *ObjectWithArgs, actions *NodeList) *AlterFunctionStmt {
	_ = "STUB: not implemented"
	return nil
}

func (a *AlterFunctionStmt) String() string { _ = "STUB: not implemented"; return "" }

func (a *AlterFunctionStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of ALTER FUNCTION statement
func (a *AlterFunctionStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// =============================================================================
// ALTER TYPE Statements
// =============================================================================

// AlterTypeStmt represents ALTER TYPE statement (for base types)
// Ported from postgres/src/include/nodes/parsenodes.h AlterTypeStmt
type AlterTypeStmt struct {
	BaseNode
	TypeName *NodeList `json:"typename"` // type name (possibly qualified)
	Options  *NodeList `json:"options"`  // List of DefElem nodes
}

func NewAlterTypeStmt(typeName *NodeList, options *NodeList) *AlterTypeStmt {
	_ = "STUB: not implemented"
	return nil
}

func (a *AlterTypeStmt) String() string { _ = "STUB: not implemented"; return "" }

func (a *AlterTypeStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of ALTER TYPE statement
func (a *AlterTypeStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// When Arg is nil, it represents "= NONE" in PostgreSQL

// =============================================================================
// Supporting Types for OPERATOR CLASS
// =============================================================================

// OpClassItemType represents types of operator class items
type OpClassItemType int

const (
	OPCLASS_ITEM_OPERATOR    OpClassItemType = 1 // #define OPCLASS_ITEM_OPERATOR 1
	OPCLASS_ITEM_FUNCTION    OpClassItemType = 2 // #define OPCLASS_ITEM_FUNCTION 2
	OPCLASS_ITEM_STORAGETYPE OpClassItemType = 3 // #define OPCLASS_ITEM_STORAGETYPE 3
)
