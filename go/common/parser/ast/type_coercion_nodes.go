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
// Package ast provides PostgreSQL AST type coercion and advanced expression node definitions.
// These nodes handle PostgreSQL's sophisticated type system including type casting,
// field access, array operations, and various test expressions.
// Ported from postgres/src/include/nodes/primnodes.h
package ast

// ==============================================================================
// TYPE COERCION AND ADVANCED EXPRESSION NODES - PostgreSQL Type System
// ==============================================================================

// RelabelType represents type casting/relabeling operations.
// This is the most common type coercion mechanism in PostgreSQL, used when
// the representation doesn't change but the type label does.
// Ported from postgres/src/include/nodes/primnodes.h:1181
type RelabelType struct {
	BaseExpr
	Arg           Expression   // Input expression - primnodes.h:1184
	Resulttype    Oid          // Output type OID - primnodes.h:1185
	Resulttypmod  int32        // Output typmod (usually -1) - primnodes.h:1187
	Resultcollid  Oid          // OID of collation, or InvalidOid if none - primnodes.h:1189
	Relabelformat CoercionForm // How to display this node - primnodes.h:1191
}

// NewRelabelType creates a new RelabelType node.
func NewRelabelType(arg Expression, resulttype Oid, resulttypmod int32, relabelformat CoercionForm) *RelabelType {
	_ = "STUB: not implemented"
	return nil
}

// NewImplicitRelabelType creates a new RelabelType for implicit casts.
func NewImplicitRelabelType(arg Expression, resulttype Oid) *RelabelType {
	_ = "STUB: not implemented"
	return nil
}

// NewExplicitRelabelType creates a new RelabelType for explicit casts.
func NewExplicitRelabelType(arg Expression, resulttype Oid) *RelabelType {
	_ = "STUB: not implemented"
	return nil
}

func (rt *RelabelType) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (rt *RelabelType) String() string { _ = "STUB: not implemented"; return "" }

// CoerceViaIO represents type coercion through I/O functions.
// This is used when types need to be converted by invoking their I/O functions
// (output function of source type, input function of target type).
// Ported from postgres/src/include/nodes/primnodes.h:1204
type CoerceViaIO struct {
	BaseExpr
	Arg          Expression   // Input expression - primnodes.h:1207
	Resulttype   Oid          // Output type OID - primnodes.h:1208
	Resultcollid Oid          // OID of collation, or InvalidOid if none - primnodes.h:1210
	Coerceformat CoercionForm // How to display this coercion - primnodes.h:1211
}

// NewCoerceViaIO creates a new CoerceViaIO node.
func NewCoerceViaIO(arg Expression, resulttype Oid, coerceformat CoercionForm) *CoerceViaIO {
	_ = "STUB: not implemented"
	return nil
}

// NewExplicitCoerceViaIO creates a new CoerceViaIO for explicit coercion.
func NewExplicitCoerceViaIO(arg Expression, resulttype Oid) *CoerceViaIO {
	_ = "STUB: not implemented"
	return nil
}

func (cvio *CoerceViaIO) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (cvio *CoerceViaIO) String() string { _ = "STUB: not implemented"; return "" }

// ArrayCoerceExpr represents array type coercion.
// This handles coercion of array types, including element-wise coercion.
// Ported from postgres/src/include/nodes/primnodes.h:1230
type ArrayCoerceExpr struct {
	BaseExpr
	Arg          Expression   // Input array expression - primnodes.h:1233
	Elemexpr     Expression   // Expression representing per-element work - primnodes.h:1234
	Resulttype   Oid          // Output type OID (array type) - primnodes.h:1235
	Resulttypmod int32        // Output typmod (usually -1) - primnodes.h:1236
	Resultcollid Oid          // OID of collation, or InvalidOid if none - primnodes.h:1237
	Coerceformat CoercionForm // How to display this coercion - primnodes.h:1238
}

// NewArrayCoerceExpr creates a new ArrayCoerceExpr node.
func NewArrayCoerceExpr(arg, elemexpr Expression, resulttype Oid, coerceformat CoercionForm) *ArrayCoerceExpr {
	_ = "STUB: not implemented"
	return nil
}

// NewExplicitArrayCoerceExpr creates a new explicit ArrayCoerceExpr.
func NewExplicitArrayCoerceExpr(arg, elemexpr Expression, resulttype Oid) *ArrayCoerceExpr {
	_ = "STUB: not implemented"
	return nil
}

func (ace *ArrayCoerceExpr) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (ace *ArrayCoerceExpr) String() string { _ = "STUB: not implemented"; return "" }

// ConvertRowtypeExpr represents row type conversion.
// This converts a whole-row value from one composite type to another.
// Ported from postgres/src/include/nodes/primnodes.h:1258
type ConvertRowtypeExpr struct {
	BaseExpr
	Arg           Expression   // Input expression - primnodes.h:1706
	Resulttype    Oid          // Output type (always a composite type) - primnodes.h:1707
	Convertformat CoercionForm // How to display this node - primnodes.h:1708
}

// NewConvertRowtypeExpr creates a new ConvertRowtypeExpr node.
func NewConvertRowtypeExpr(arg Expression, resulttype Oid, convertformat CoercionForm) *ConvertRowtypeExpr {
	_ = "STUB: not implemented"
	return nil
}

func (crte *ConvertRowtypeExpr) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (crte *ConvertRowtypeExpr) String() string { _ = "STUB: not implemented"; return "" }

// CollateExpr represents a COLLATE expression.
// This specifies a collation to be used for a particular expression.
// Ported from postgres/src/include/nodes/primnodes.h:1276
type CollateExpr struct {
	BaseExpr
	Arg     Expression // Input expression - primnodes.h:1649
	CollOid Oid        // Collation OID - primnodes.h:1650
}

// NewCollateExpr creates a new CollateExpr node.
func NewCollateExpr(arg Expression, collOid Oid) *CollateExpr {
	_ = "STUB: not implemented"
	return nil
}

func (ce *CollateExpr) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (ce *CollateExpr) String() string { _ = "STUB: not implemented"; return "" }

// ==============================================================================
// FIELD AND RECORD OPERATIONS
// ==============================================================================

// FieldSelect represents field selection from a composite value (record.field).
// This extracts a single field from a composite type value.
// Ported from postgres/src/include/nodes/primnodes.h:1125
type FieldSelect struct {
	BaseExpr
	Arg          Expression // Input expression (composite type) - primnodes.h:1409
	Fieldnum     AttrNumber // Attribute number of field to extract - primnodes.h:1410
	Resulttype   Oid        // Type OID of the field - primnodes.h:1411
	Resulttypmod int32      // Output typmod (usually -1) - primnodes.h:1412
}

// NewFieldSelect creates a new FieldSelect node.
func NewFieldSelect(arg Expression, fieldnum AttrNumber, resulttype Oid) *FieldSelect {
	_ = "STUB: not implemented"
	return nil
}

func (fs *FieldSelect) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (fs *FieldSelect) String() string { _ = "STUB: not implemented"; return "" }

// FieldStore represents field assignment to a composite value.
// This is used for UPDATE operations on composite type columns.
// Ported from postgres/src/include/nodes/primnodes.h:1156
type FieldStore struct {
	BaseExpr
	Arg        Expression   // Input expression (composite type) - primnodes.h:1429
	Newvals    []Expression // New value(s) for field(s) - primnodes.h:1430
	Fieldnums  []AttrNumber // Field number(s) to be updated - primnodes.h:1431
	Resulttype Oid          // Type OID of result (same as input type) - primnodes.h:1432
}

// NewFieldStore creates a new FieldStore node.
func NewFieldStore(arg Expression, newvals []Expression, fieldnums []AttrNumber, resulttype Oid) *FieldStore {
	_ = "STUB: not implemented"
	return nil
}

// NewSingleFieldStore creates a FieldStore for updating a single field.
func NewSingleFieldStore(arg Expression, newval Expression, fieldnum AttrNumber, resulttype Oid) *FieldStore {
	_ = "STUB: not implemented"
	return nil
}

func (fs *FieldStore) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (fs *FieldStore) String() string { _ = "STUB: not implemented"; return "" }

// SubscriptingRef represents array/JSON subscripting operations.
// This handles both array indexing (arr[1]) and JSON key access (json['key']).
// Ported from postgres/src/include/nodes/primnodes.h:679
type SubscriptingRef struct {
	BaseExpr
	Refcontainertype Oid          // Type OID of container (array or jsonb) - primnodes.h:682
	Refelemtype      Oid          // The container type's pg_type.typelem - primnodes.h:683
	Refrestype       Oid          // Type OID of the SubscriptingRef's result - primnodes.h:684
	Reftypmod        int32        // Typmod of the result - primnodes.h:685
	Refcollid        Oid          // Collation of result, or InvalidOid if none - primnodes.h:686
	Refupperindexpr  []Expression // Expressions for upper index bounds - primnodes.h:687
	Reflowerindexpr  []Expression // Expressions for lower index bounds - primnodes.h:688
	Refexpr          Expression   // Expression for the container value - primnodes.h:689
	Refassgnexpr     Expression   // Expression for new value in assignment - primnodes.h:690
}

// NewSubscriptingRef creates a new SubscriptingRef node.
func NewSubscriptingRef(containertype, elemtype, restype Oid, refexpr Expression, upperindex []Expression) *SubscriptingRef {
	_ = "STUB: not implemented"
	return nil
}

// NewArraySubscript creates a SubscriptingRef for array indexing (arr[index]).
func NewArraySubscript(arraytype, elemtype Oid, arrayexpr, indexexpr Expression) *SubscriptingRef {
	_ = "STUB: not implemented"
	return nil
}

// Result type is element type for array indexing

// NewArraySlice creates a SubscriptingRef for array slicing (arr[lower:upper]).
func NewArraySlice(arraytype, elemtype Oid, arrayexpr, lowerexpr, upperexpr Expression) *SubscriptingRef {
	_ = "STUB: not implemented"
	return nil
}

// Result type is array type for slicing

// NewArrayAssignment creates a SubscriptingRef for array assignment (arr[index] = value).
func NewArrayAssignment(arraytype, elemtype Oid, arrayexpr, indexexpr, assignexpr Expression) *SubscriptingRef {
	_ = "STUB: not implemented"
	return nil
}

// Result type is array type for assignment

func (sr *SubscriptingRef) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (sr *SubscriptingRef) String() string { _ = "STUB: not implemented"; return "" }

// ==============================================================================
// TEST EXPRESSIONS - NULL, BOOLEAN, AND DOMAIN TESTS
// ==============================================================================

// NullTestType represents the type of NULL test.
// Ported from postgres/src/include/nodes/primnodes.h:1950
type NullTestType int

const (
	IS_NULL     NullTestType = iota // IS NULL - primnodes.h:1952
	IS_NOT_NULL                     // IS NOT NULL - primnodes.h:1952
)

// NullTest represents IS NULL and IS NOT NULL tests.
// This is one of the most fundamental SQL test expressions.
// Ported from postgres/src/include/nodes/primnodes.h:1955
type NullTest struct {
	BaseExpr
	Arg          Expression   // Input expression - primnodes.h:1958
	Nulltesttype NullTestType // IS NULL or IS NOT NULL - primnodes.h:1959
	Argisrow     bool         // True if input is known to be a row value - primnodes.h:1961
}

// NewNullTest creates a new NullTest node.
func NewNullTest(arg Expression, nulltesttype NullTestType) *NullTest {
	_ = "STUB: not implemented"
	return nil
}

// NewIsNullTest creates a new IS NULL test.
func NewIsNullTest(arg Expression) *NullTest { _ = "STUB: not implemented"; return nil }

// NewIsNotNullTest creates a new IS NOT NULL test.
func NewIsNotNullTest(arg Expression) *NullTest { _ = "STUB: not implemented"; return nil }

// NewRowNullTest creates a NullTest for row values.
func NewRowNullTest(arg Expression, nulltesttype NullTestType) *NullTest {
	_ = "STUB: not implemented"
	return nil
}

func (nt *NullTest) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (nt *NullTest) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of NullTest
func (nt *NullTest) SqlString() string { _ = "STUB: not implemented"; return "" }

// BoolTestType represents the type of boolean test.
// Ported from postgres/src/include/nodes/primnodes.h:1974
type BoolTestType int

const (
	IS_TRUE        BoolTestType = iota // IS TRUE - primnodes.h:1800
	IS_NOT_TRUE                        // IS NOT TRUE - primnodes.h:1801
	IS_FALSE                           // IS FALSE - primnodes.h:1802
	IS_NOT_FALSE                       // IS NOT FALSE - primnodes.h:1803
	IS_UNKNOWN                         // IS UNKNOWN - primnodes.h:1804
	IS_NOT_UNKNOWN                     // IS NOT UNKNOWN - primnodes.h:1805
)

// BooleanTest represents boolean test expressions (IS TRUE, IS FALSE, etc.).
// These tests handle three-valued boolean logic (TRUE/FALSE/UNKNOWN).
// Ported from postgres/src/include/nodes/primnodes.h:1979
type BooleanTest struct {
	BaseExpr
	Arg          Expression   // Input expression - primnodes.h:1806
	Booltesttype BoolTestType // Kind of test - primnodes.h:1807
}

// NewBooleanTest creates a new BooleanTest node.
func NewBooleanTest(arg Expression, booltesttype BoolTestType) *BooleanTest {
	_ = "STUB: not implemented"
	return nil
}

// NewIsTrueTest creates a new IS TRUE test.
func NewIsTrueTest(arg Expression) *BooleanTest { _ = "STUB: not implemented"; return nil }

// NewIsFalseTest creates a new IS FALSE test.
func NewIsFalseTest(arg Expression) *BooleanTest { _ = "STUB: not implemented"; return nil }

// NewIsUnknownTest creates a new IS UNKNOWN test.
func NewIsUnknownTest(arg Expression) *BooleanTest { _ = "STUB: not implemented"; return nil }

func (bt *BooleanTest) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (bt *BooleanTest) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of BooleanTest
func (bt *BooleanTest) SqlString() string { _ = "STUB: not implemented"; return "" }

// CoerceToDomain represents coercion to a domain type.
// Domain types are user-defined types with constraints.
// Ported from postgres/src/include/nodes/primnodes.h:2025
type CoerceToDomain struct {
	BaseExpr
	Arg            Expression   // Input expression - primnodes.h:1716
	Resulttype     Oid          // Domain type OID - primnodes.h:1717
	Resulttypmod   int32        // Output typmod (usually -1) - primnodes.h:1718
	Resultcollid   Oid          // OID of collation, or InvalidOid if none - primnodes.h:1719
	Coercionformat CoercionForm // How to display this coercion - primnodes.h:1720
}

// NewCoerceToDomain creates a new CoerceToDomain node.
func NewCoerceToDomain(arg Expression, resulttype Oid, resulttypmod int32, coercionformat CoercionForm) *CoerceToDomain {
	_ = "STUB: not implemented"
	return nil
}

func (ctd *CoerceToDomain) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (ctd *CoerceToDomain) String() string { _ = "STUB: not implemented"; return "" }

// CoerceToDomainValue represents a value being coerced to a domain type.
// This is used in domain constraint checking.
// Ported from postgres/src/include/nodes/primnodes.h:2048
type CoerceToDomainValue struct {
	BaseExpr
	TypeId    Oid   // Type for substituted value - primnodes.h:2051
	TypeMod   int32 // Typemod for substituted value - primnodes.h:2052
	Collation Oid   // Collation for the substituted value - primnodes.h:2053
}

// NewCoerceToDomainValue creates a new CoerceToDomainValue node.
func NewCoerceToDomainValue(typeId Oid, typeMod int32, collation Oid) *CoerceToDomainValue {
	_ = "STUB: not implemented"
	return nil
}

func (ctdv *CoerceToDomainValue) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (ctdv *CoerceToDomainValue) String() string { _ = "STUB: not implemented"; return "" }

// ==============================================================================
// SPECIAL VALUE EXPRESSIONS
// ==============================================================================

// SetToDefault represents a DEFAULT expression.
// This is used in INSERT and UPDATE statements to indicate use of the default value.
// Ported from postgres/src/include/nodes/primnodes.h:2068
type SetToDefault struct {
	BaseExpr
	TypeId    Oid   // Type for substituted value - primnodes.h:2071
	TypeMod   int32 // Typemod for substituted value - primnodes.h:2072
	Collation Oid   // Collation for the substituted value - primnodes.h:2073
}

// NewSetToDefault creates a new SetToDefault node.
func NewSetToDefault(typeId Oid, typeMod int32, collation Oid) *SetToDefault {
	_ = "STUB: not implemented"
	return nil
}

func (std *SetToDefault) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (std *SetToDefault) String() string { _ = "STUB: not implemented"; return "" }

// SqlString returns the SQL representation of SetToDefault
func (std *SetToDefault) SqlString() string {
	_ = "STUB: not implemented"

	// CurrentOfExpr represents CURRENT OF cursor_name expressions.
	// This is used in UPDATE and DELETE statements to refer to the current row of a cursor.
	// Ported from postgres/src/include/nodes/primnodes.h:2094
	return ""
}

type CurrentOfExpr struct {
	BaseExpr
	Cvarno      Index  // RT index of target relation - primnodes.h:2097
	CursorName  string // Name of referenced cursor, or NULL - primnodes.h:2098
	CursorParam int    // Refcursor parameter number, or 0 - primnodes.h:2099
}

// SqlString returns the SQL representation of the CurrentOfExpr.
func (c *CurrentOfExpr) SqlString() string { _ = "STUB: not implemented"; return "" }

// NewCurrentOfExpr creates a new CurrentOfExpr node.
func NewCurrentOfExpr(cvarno Index, cursor_name string) *CurrentOfExpr {
	_ = "STUB: not implemented"
	return nil
}

// NewCurrentOfExprParam creates a new CurrentOfExpr with parameter reference.
func NewCurrentOfExprParam(cvarno Index, cursor_param int) *CurrentOfExpr {
	_ = "STUB: not implemented"
	return nil
}

func (coe *CurrentOfExpr) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (coe *CurrentOfExpr) String() string { _ = "STUB: not implemented"; return "" }

// NextValueExpr represents nextval() and currval() sequence operations.
// This handles sequence value generation and retrieval.
// Ported from postgres/src/include/nodes/primnodes.h:2109
type NextValueExpr struct {
	BaseExpr
	Seqid  Oid // OID of sequence relation - primnodes.h:2112
	TypeId Oid // Type OID of result - primnodes.h:2113
}

// NewNextValueExpr creates a new NextValueExpr node.
func NewNextValueExpr(seqid, typeId Oid) *NextValueExpr { _ = "STUB: not implemented"; return nil }

func (nve *NextValueExpr) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (nve *NextValueExpr) String() string { _ = "STUB: not implemented"; return "" }

// InferenceElem represents an inference element for ON CONFLICT clauses.
// This is used to specify which unique index to use for conflict detection.
// Ported from postgres/src/include/nodes/primnodes.h:2123
type InferenceElem struct {
	BaseExpr
	Expr         Node // Expression to infer from - primnodes.h:2014
	Infercollid  Oid  // OID of collation, or InvalidOid - primnodes.h:2015
	Inferopclass Oid  // OID of operator class, or InvalidOid - primnodes.h:2016
}

// NewInferenceElem creates a new InferenceElem node.
func NewInferenceElem(expr Node) *InferenceElem { _ = "STUB: not implemented"; return nil }

// NewInferenceElemWithCollation creates a new InferenceElem with collation.
func NewInferenceElemWithCollation(expr Node, infercollid, inferopclass Oid) *InferenceElem {
	_ = "STUB: not implemented"
	return nil
}

func (ie *InferenceElem) ExpressionType() string { _ = "STUB: not implemented"; return "" }

func (ie *InferenceElem) String() string { _ = "STUB: not implemented"; return "" }
