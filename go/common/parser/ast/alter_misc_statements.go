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
// ALTER Miscellaneous Statement Nodes - Phase 3J Implementation
// ==============================================================================

// Helper function to convert NodeList of strings to qualified name
func nodeListToQualifiedName(nodeList *NodeList) string { _ = "STUB: not implemented"; return "" }

// AlterTSConfigType represents the type of ALTER TEXT SEARCH CONFIGURATION operation
// Ported from postgres/src/include/nodes/parsenodes.h:3700-3706
type AlterTSConfigType int

const (
	ALTER_TSCONFIG_ADD_MAPPING AlterTSConfigType = iota
	ALTER_TSCONFIG_ALTER_MAPPING_FOR_TOKEN
	ALTER_TSCONFIG_ALTER_MAPPING_REPLACE
	ALTER_TSCONFIG_REPLACE_DICT
	ALTER_TSCONFIG_REPLACE_DICT_FOR_TOKEN
	ALTER_TSCONFIG_DROP_MAPPING
)

func (a AlterTSConfigType) String() string { _ = "STUB: not implemented"; return "" }

// AlterObjectSchemaStmt represents ALTER ... SET SCHEMA statements
// Ported from postgres/src/include/nodes/parsenodes.h:3547-3554
type AlterObjectSchemaStmt struct {
	BaseNode
	ObjectType ObjectType `json:"objectType"` // OBJECT_TABLE, OBJECT_TYPE, etc
	Relation   *RangeVar  `json:"relation"`   // In case it's a table
	Object     Node       `json:"object"`     // In case it's some other object
	NewSchema  string     `json:"newSchema"`  // The new schema
	MissingOk  bool       `json:"missingOk"`  // Skip error if missing?
}

func (n *AlterObjectSchemaStmt) node() { _ = "STUB: not implemented"; return }
func (n *AlterObjectSchemaStmt) stmt() { _ = "STUB: not implemented"; return }

func (n *AlterObjectSchemaStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

func (n *AlterObjectSchemaStmt) String() string { _ = "STUB: not implemented"; return "" }

func (n *AlterObjectSchemaStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// Add object type

// Add IF EXISTS if specified

// Add object name

// Special handling for OPERATOR CLASS and OPERATOR FAMILY

// Format as "name USING method" where first item is method, rest is name

// Add SET SCHEMA clause

// NewAlterObjectSchemaStmt creates a new AlterObjectSchemaStmt node
func NewAlterObjectSchemaStmt(objectType ObjectType, newSchema string) *AlterObjectSchemaStmt {
	_ = "STUB: not implemented"
	return nil
}

// AlterOperatorStmt represents ALTER OPERATOR statements
// Ported from postgres/src/include/nodes/parsenodes.h:3585-3589
type AlterOperatorStmt struct {
	BaseNode
	Opername *ObjectWithArgs `json:"opername"` // Operator name and argument types
	Options  *NodeList       `json:"options"`  // List of DefElem nodes
}

func (n *AlterOperatorStmt) node() { _ = "STUB: not implemented"; return }
func (n *AlterOperatorStmt) stmt() { _ = "STUB: not implemented"; return }

func (n *AlterOperatorStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

func (n *AlterOperatorStmt) String() string { _ = "STUB: not implemented"; return "" }

func (n *AlterOperatorStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// COMMUTATOR/NEGATOR carry operator names, which must be emitted
// unquoted rather than as string literals.

// NewAlterOperatorStmt creates a new AlterOperatorStmt node
func NewAlterOperatorStmt(opername *ObjectWithArgs, options *NodeList) *AlterOperatorStmt {
	_ = "STUB: not implemented"
	return nil
}

// AlterObjectDependsStmt represents ALTER ... DEPENDS statements
// Ported from postgres/src/include/nodes/parsenodes.h:3556-3564
type AlterObjectDependsStmt struct {
	BaseNode
	ObjectType ObjectType `json:"objectType"` // OBJECT_FUNCTION, OBJECT_TRIGGER, etc
	Relation   *RangeVar  `json:"relation"`   // In case a table is involved
	Object     Node       `json:"object"`     // Name of the object
	Extname    *String    `json:"extname"`    // Extension name
	Remove     bool       `json:"remove"`     // Set true to remove dep rather than add
}

func (n *AlterObjectDependsStmt) node() { _ = "STUB: not implemented"; return }
func (n *AlterObjectDependsStmt) stmt() { _ = "STUB: not implemented"; return }

func (n *AlterObjectDependsStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

func (n *AlterObjectDependsStmt) String() string { _ = "STUB: not implemented"; return "" }

func (n *AlterObjectDependsStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// Add object type

// Add object name

// For TRIGGER, format as "trigger_name ON table_name"

// Add DEPENDS clause

// NewAlterObjectDependsStmt creates a new AlterObjectDependsStmt node
func NewAlterObjectDependsStmt(objectType ObjectType, extname *String, remove bool) *AlterObjectDependsStmt {
	_ = "STUB: not implemented"
	return nil
}

// AlterCollationStmt represents ALTER COLLATION statements
// Ported from postgres/src/include/nodes/parsenodes.h:2510-2514
type AlterCollationStmt struct {
	BaseNode
	Collname *NodeList `json:"collname"` // Qualified name
}

func (n *AlterCollationStmt) node() { _ = "STUB: not implemented"; return }
func (n *AlterCollationStmt) stmt() { _ = "STUB: not implemented"; return }

func (n *AlterCollationStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

func (n *AlterCollationStmt) String() string { _ = "STUB: not implemented"; return "" }

func (n *AlterCollationStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// NewAlterCollationStmt creates a new AlterCollationStmt node
func NewAlterCollationStmt(collname *NodeList) *AlterCollationStmt {
	_ = "STUB: not implemented"
	return nil
}

// AlterDatabaseStmt represents ALTER DATABASE statements
// Ported from postgres/src/include/nodes/parsenodes.h:3221-3226
type AlterDatabaseStmt struct {
	BaseNode
	Dbname  string    `json:"dbname"`  // Name of database to alter
	Options *NodeList `json:"options"` // List of DefElem nodes
}

func (n *AlterDatabaseStmt) node() { _ = "STUB: not implemented"; return }
func (n *AlterDatabaseStmt) stmt() { _ = "STUB: not implemented"; return }

func (n *AlterDatabaseStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

func (n *AlterDatabaseStmt) String() string { _ = "STUB: not implemented"; return "" }

func (n *AlterDatabaseStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// NewAlterDatabaseStmt creates a new AlterDatabaseStmt node
func NewAlterDatabaseStmt(dbname string, options *NodeList) *AlterDatabaseStmt {
	_ = "STUB: not implemented"
	return nil
}

// AlterDatabaseSetStmt represents ALTER DATABASE SET statements
// Ported from postgres/src/include/nodes/parsenodes.h:3238-3243
type AlterDatabaseSetStmt struct {
	BaseNode
	Dbname  string           `json:"dbname"`  // Database name
	Setstmt *VariableSetStmt `json:"setstmt"` // SET or RESET subcommand
}

func (n *AlterDatabaseSetStmt) node() { _ = "STUB: not implemented"; return }
func (n *AlterDatabaseSetStmt) stmt() { _ = "STUB: not implemented"; return }

func (n *AlterDatabaseSetStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

func (n *AlterDatabaseSetStmt) String() string { _ = "STUB: not implemented"; return "" }

func (n *AlterDatabaseSetStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// NewAlterDatabaseSetStmt creates a new AlterDatabaseSetStmt node
func NewAlterDatabaseSetStmt(dbname string, setstmt *VariableSetStmt) *AlterDatabaseSetStmt {
	_ = "STUB: not implemented"
	return nil
}

// AlterDatabaseRefreshCollStmt represents ALTER DATABASE REFRESH COLLATION VERSION statements
// Ported from postgres/src/include/nodes/parsenodes.h
type AlterDatabaseRefreshCollStmt struct {
	BaseNode
	Dbname string `json:"dbname"` // Database name
}

func (n *AlterDatabaseRefreshCollStmt) node() { _ = "STUB: not implemented"; return }
func (n *AlterDatabaseRefreshCollStmt) stmt() { _ = "STUB: not implemented"; return }

func (n *AlterDatabaseRefreshCollStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

func (n *AlterDatabaseRefreshCollStmt) String() string { _ = "STUB: not implemented"; return "" }

func (n *AlterDatabaseRefreshCollStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// NewAlterDatabaseRefreshCollStmt creates a new AlterDatabaseRefreshCollStmt node
func NewAlterDatabaseRefreshCollStmt(dbname string) *AlterDatabaseRefreshCollStmt {
	_ = "STUB: not implemented"
	return nil
}

// AlterCompositeTypeStmt represents ALTER TYPE (composite) statements
// Note: This is actually handled by AlterTableStmt in PostgreSQL but
// represented here for grammar rule completeness
type AlterCompositeTypeStmt struct {
	BaseNode
	TypeName *NodeList `json:"typeName"` // Qualified type name (list of String)
	Cmds     *NodeList `json:"cmds"`     // List of alter type commands
}

func (n *AlterCompositeTypeStmt) node() { _ = "STUB: not implemented"; return }
func (n *AlterCompositeTypeStmt) stmt() { _ = "STUB: not implemented"; return }

func (n *AlterCompositeTypeStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

func (n *AlterCompositeTypeStmt) String() string { _ = "STUB: not implemented"; return "" }

func (n *AlterCompositeTypeStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// NewAlterCompositeTypeStmt creates a new AlterCompositeTypeStmt node
func NewAlterCompositeTypeStmt(typeName *NodeList, cmds *NodeList) *AlterCompositeTypeStmt {
	_ = "STUB: not implemented"
	return nil
}

// AlterTSConfigurationStmt represents ALTER TEXT SEARCH CONFIGURATION statements
// Ported from postgres/src/include/nodes/parsenodes.h:3708-3721
type AlterTSConfigurationStmt struct {
	BaseNode
	Kind      AlterTSConfigType `json:"kind"`      // ALTER_TSCONFIG_ADD_MAPPING, etc
	Cfgname   *NodeList         `json:"cfgname"`   // Qualified name (list of String)
	Tokentype *NodeList         `json:"tokentype"` // List of String
	Dicts     *NodeList         `json:"dicts"`     // List of String
	Override  bool              `json:"override"`  // If true, replace rather than add
	Replace   bool              `json:"replace"`   // If true, replace rather than modify
	MissingOk bool              `json:"missingOk"` // For ALTER ... IF EXISTS
}

func (n *AlterTSConfigurationStmt) node() { _ = "STUB: not implemented"; return }
func (n *AlterTSConfigurationStmt) stmt() { _ = "STUB: not implemented"; return }

func (n *AlterTSConfigurationStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

func (n *AlterTSConfigurationStmt) String() string { _ = "STUB: not implemented"; return "" }

func (n *AlterTSConfigurationStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// Add operation type with special handling for ALTER MAPPING patterns

// This is likely an ALTER MAPPING REPLACE that was parsed as REPLACE_DICT

// Format: ALTER MAPPING FOR tokens REPLACE

// Add FOR tokens

// Add REPLACE

// Format: ALTER MAPPING REPLACE

// Standard operation types

// Add IF EXISTS if specified (for DROP operations)

// Add FOR token types if specified (for non-REPLACE operations)

// Only add "FOR" if it's not already in the kind string

// Add WITH dictionaries if specified

// Special format for ALTER MAPPING REPLACE: "old_dict WITH new_dict"

// Standard format: "WITH dict1, dict2, ..."

// NewAlterTSConfigurationStmt creates a new AlterTSConfigurationStmt node
func NewAlterTSConfigurationStmt(kind AlterTSConfigType, cfgname *NodeList) *AlterTSConfigurationStmt {
	_ = "STUB: not implemented"
	return nil
}

// AlterTSDictionaryStmt represents ALTER TEXT SEARCH DICTIONARY statements
// Ported from postgres/src/include/nodes/parsenodes.h:3693-3698
type AlterTSDictionaryStmt struct {
	BaseNode
	Dictname *NodeList `json:"dictname"` // Qualified name (list of String)
	Options  *NodeList `json:"options"`  // List of DefElem nodes
}

func (n *AlterTSDictionaryStmt) node() { _ = "STUB: not implemented"; return }
func (n *AlterTSDictionaryStmt) stmt() { _ = "STUB: not implemented"; return }

func (n *AlterTSDictionaryStmt) StatementType() string { _ = "STUB: not implemented"; return "" }

func (n *AlterTSDictionaryStmt) String() string { _ = "STUB: not implemented"; return "" }

func (n *AlterTSDictionaryStmt) SqlString() string { _ = "STUB: not implemented"; return "" }

// NewAlterTSDictionaryStmt creates a new AlterTSDictionaryStmt node
func NewAlterTSDictionaryStmt(dictname *NodeList, options *NodeList) *AlterTSDictionaryStmt {
	_ = "STUB: not implemented"
	return nil
}
