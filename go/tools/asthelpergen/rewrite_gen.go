// Copyright 2025 Supabase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package asthelpergen

import (
	"go/types"

	"github.com/dave/jennifer/jen"
)

const (
	rewriteName = "rewrite"
)

type rewriteGen struct {
	ifaceName string
	file      *jen.File
}

var _ generator = (*rewriteGen)(nil)

func newRewriterGen(pkgname string, ifaceName string) *rewriteGen {
	_ = "STUB: not implemented"
	return nil
}

// genFile returns the generated rewrite file
func (r *rewriteGen) genFile(_ generatorSPI) (string, *jen.File) {
	_ = "STUB: not implemented"
	return "", nil
}

// interfaceMethod handles rewriting for interface types
func (r *rewriteGen) interfaceMethod(t types.Type, iface *types.Interface, spi generatorSPI) error {
	_ = "STUB: not implemented"
	return nil
}

// nil check

// switch on actual type

// skip if it's itself an interface

// structMethod handles rewriting for struct types
func (r *rewriteGen) structMethod(t types.Type, strct *types.Struct, spi generatorSPI) error {
	_ = "STUB: not implemented"
	return nil
}

// ptrToStructMethod handles rewriting for pointer-to-struct types
func (r *rewriteGen) ptrToStructMethod(t types.Type, strct *types.Struct, spi generatorSPI) error {
	_ = "STUB: not implemented"
	return nil
}

// ptrToBasicMethod handles rewriting for pointer-to-basic types
func (r *rewriteGen) ptrToBasicMethod(t types.Type, _ *types.Basic, spi generatorSPI) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *rewriteGen) basicMethod(t types.Type, _ *types.Basic, spi generatorSPI) error {
	_ = "STUB: not implemented"
	return nil
}

// sliceMethod handles rewriting for slice types
func (r *rewriteGen) sliceMethod(t types.Type, slice *types.Slice, spi generatorSPI) error {
	_ = "STUB: not implemented"
	return nil
}

// run pre if available

// rewrite slice elements if needed

// run post if available

// rewriteStructFields processes a struct's fields, running pre/post if applicable
func (r *rewriteGen) rewriteStructFields(t types.Type, strct *types.Struct, spi generatorSPI, fail bool) []jen.Code {
	_ = "STUB: not implemented"
	return nil
}

// rewriteAllFields checks struct fields and rewrites relevant ones
func (r *rewriteGen) rewriteAllFields(t types.Type, strct *types.Struct, spi generatorSPI, fail bool) []jen.Code {
	_ = "STUB: not implemented"
	return nil
}

// single field that implements the interface

// slice of fields that implement the interface

// rewriteChild handles rewriting a single struct field that implements iface
func (r *rewriteGen) rewriteChild(
	t, field types.Type,
	fieldName string,
	param jen.Code,
	replace jen.Code,
	fail bool,
) []jen.Code {
	_ = "STUB: not implemented"
	return nil
}

// replacerFunc block

// rewriting call

// rewriteChildSlice handles rewriting an element in a slice field
func (r *rewriteGen) rewriteChildSlice(
	t, field types.Type,
	fieldName string,
	param jen.Code,
	replace jen.Code,
	fail bool,
) []jen.Code {
	_ = "STUB: not implemented"
	return nil
}

// define a replacer function

// rewriting call

// rewriteFunc is the top-level generator of the rewrite methods
func (r *rewriteGen) rewriteFunc(t types.Type, stmts []jen.Code, source string) {
	_ = "STUB: not implemented"
	return
}

// setupCursor initializes a.cur with replacer, parent, and node
func setupCursor() []jen.Code { _ = "STUB: not implemented"; return nil }

// executePreRewrite runs the pre step if a.pre is not nil
func (r *rewriteGen) executePreRewrite() jen.Code { _ = "STUB: not implemented"; return *new(jen.Code) }

// executePostRewrite runs the post step if a.post is not nil
func executePostRewrite(seenChildren bool) jen.Code {
	_ = "STUB: not implemented"
	return *new(jen.Code)
}

// re-init cursor fields if children were visited

func returnTrue() jen.Code { _ = "STUB: not implemented"; return *new(jen.Code) }

func returnFalse() jen.Code { _ = "STUB: not implemented"; return *new(jen.Code) }

// shouldAdd checks if a type should be included in code generation
func shouldAdd(t types.Type, i *types.Interface) bool { _ = "STUB: not implemented"; return false }
