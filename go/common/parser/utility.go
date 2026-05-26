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
package parser

import (
	"github.com/multigres/multigres/go/common/parser/ast"
)

// linitial returns the first element of a NodeList, equivalent to PostgreSQL's linitial()
func linitial(list *ast.NodeList) ast.Node { _ = "STUB: not implemented"; return *new(ast.Node) }

// lsecond returns the second element of a NodeList, equivalent to PostgreSQL's lsecond()
func lsecond(list *ast.NodeList) ast.Node { _ = "STUB: not implemented"; return *new(ast.Node) }

// lthird returns the third element of a NodeList, equivalent to PostgreSQL's lthird()
func lthird(list *ast.NodeList) ast.Node { _ = "STUB: not implemented"; return *new(ast.Node) }

// llast returns the last element of a NodeList, equivalent to PostgreSQL's llast()
func llast(list *ast.NodeList) ast.Node { _ = "STUB: not implemented"; return *new(ast.Node) }

// makeTypeNameFromNodeList converts *ast.NodeList to *ast.TypeName
func makeTypeNameFromNodeList(list *ast.NodeList) *ast.TypeName {
	_ = "STUB: not implemented"
	return nil
}

func makeTypeNameFromString(str string) *ast.TypeName { _ = "STUB: not implemented"; return nil }

// makeRangeVarFromAnyName converts a list of (dotted) names to a RangeVar.
// The "AnyName" refers to the any_name production in the grammar.
//
// This function should be used when you have an any_name (NodeList of String nodes)
// but need a RangeVar for cases where the grammar has any_name but the AST expects
// a relation reference. This mirrors PostgreSQL's approach in cases where they
// comment "can't use qualified_name, sigh".
//
// Supports:
//   - Single name: "table" -> RangeVar{RelName: "table"}
//   - Schema qualified: "schema.table" -> RangeVar{SchemaName: "schema", RelName: "table"}
//   - Fully qualified: "catalog.schema.table" -> RangeVar{CatalogName: "catalog", SchemaName: "schema", RelName: "table"}
//
// Ported from PostgreSQL's makeRangeVarFromAnyName function.
func makeRangeVarFromAnyName(names *ast.NodeList, position int) (*ast.RangeVar, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Default to inheritance enabled (no ONLY)

// Single name: just the relation name

// Two names: schema.relation

// Three names: catalog.schema.relation

// makeRangeVarFromQualifiedName constructs a RangeVar from a ColId and indirection list.
// This mirrors PostgreSQL's makeRangeVarFromQualifiedName function which is used in
// grammar rules that have "ColId indirection" patterns.
//
// The function handles:
//   - ColId alone: "table" -> RangeVar{RelName: "table"}
//   - ColId + single indirection: "schema" + [".table"] -> RangeVar{SchemaName: "schema", RelName: "table"}
//   - ColId + multiple indirections: "catalog" + [".schema", ".table"] -> RangeVar{CatalogName: "catalog", SchemaName: "schema", RelName: "table"}
//
// Parameters:
//   - name: The initial ColId string
//   - indirection: NodeList containing the indirection elements (can be nil)
//   - position: Source location for error reporting
//
// Returns the constructed RangeVar.
//
// Ported from PostgreSQL's makeRangeVarFromQualifiedName function.
func makeRangeVarFromQualifiedName(name string, indirection *ast.NodeList, position int) *ast.RangeVar {
	_ = "STUB: not implemented"
	return nil
}

// Default to inheritance enabled (no ONLY)

// Start with the base name

// Add indirection elements

// Note: PostgreSQL also handles A_Star nodes for ".*" but we'll focus on String nodes for now

// Build RangeVar based on number of names

// Single name: just the relation name

// Two names: schema.relation

// Three names: catalog.schema.relation

// For more than 3 names, use the last as relation, second-to-last as schema, third-to-last as catalog
// This is a fallback - PostgreSQL would likely error on too many names

// SplitColQualList separates a ColQualList (column qualifier list) into constraints and collate clauses.
// This mirrors PostgreSQL's SplitColQualList function which is used to process column qualifiers
// and domain qualifiers by separating them into appropriate lists.
//
// PostgreSQL reference: src/backend/parser/gram.tab.c:SplitColQualList
//
// Parameters:
//   - qualList: A NodeList containing Constraint and CollateClause nodes
//
// Returns:
//   - constraints: A NodeList of Constraint nodes
//   - collClause: A single CollateClause (PostgreSQL allows only one COLLATE per column/domain)
func SplitColQualList(qualList *ast.NodeList) (*ast.NodeList, *ast.CollateClause) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PostgreSQL allows only one COLLATE clause per column/domain
// If multiple are specified, the last one wins

// Convert constraints slice to NodeList

// makeOrderedSetArgs processes arguments for hypothetical-set aggregates.
// It validates VARIADIC argument consistency and returns a list containing:
// - The concatenated direct and ordered arguments
// - An integer indicating the number of direct arguments
//
// This mirrors PostgreSQL's makeOrderedSetArgs function.
//
// PostgreSQL reference: src/backend/parser/gram.y:makeOrderedSetArgs
func makeOrderedSetArgs(directArgs *ast.NodeList, orderedArgs *ast.NodeList) (*ast.NodeList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if the last direct argument is VARIADIC

// PostgreSQL requires exactly one VARIADIC ordered argument of the same type

// TODO: Check that types are equal when we have proper type comparison
// For now, we skip type checking but drop the duplicate VARIADIC argument

// Store the number of direct arguments

// Concatenate direct and ordered arguments

// Return [concatenated_args, num_direct_args]

// extractAggrArgTypes extracts just the argument types from the output of the aggr_args production.
// This is equivalent to PostgreSQL's extractAggrArgTypes function.
func extractAggrArgTypes(aggrArgs *ast.NodeList) *ast.NodeList {
	_ = "STUB: not implemented"
	return nil
}

// First element contains the actual arguments (or nil for *)

// extractArgTypes extracts just the argument types (TypeNames) from a list of FunctionParameter nodes
// for input parameters only. This is equivalent to PostgreSQL's extractArgTypes function.
func extractArgTypes(parameters *ast.NodeList) *ast.NodeList { _ = "STUB: not implemented"; return nil }

// processConstraintAttributeSpec processes constraint attribute specification bits.
// This is a simplified version of processCASbits from PostgreSQL.
func processConstraintAttributeSpec(casbits int, constraint *ast.Constraint) {
	_ = "STUB: not implemented"
	return
}

// doNegate handles negation of nodes, equivalent to PostgreSQL's doNegate()
// Ported from postgres/src/backend/parser/gram.y:doNegate
func doNegate(n ast.Node, location int) ast.Node { _ = "STUB: not implemented"; return *new(ast.Node) }

// Default: create unary minus expression

// doNegateFloat handles negation of float values
// Ported from postgres/src/backend/parser/gram.y:doNegateFloat
func doNegateFloat(v *ast.Float) { _ = "STUB: not implemented"; return }

// Remove leading +

// Remove leading -

// Add leading -

// makeSetOp creates a set operation (UNION, INTERSECT, EXCEPT) SelectStmt
// Ported from postgres/src/backend/parser/gram.y:makeSetOp
func makeSetOp(op ast.SetOperation, all bool, larg ast.Stmt, rarg ast.Stmt) ast.Stmt {
	_ = "STUB: not implemented"
	return *new(ast.Stmt)
}
