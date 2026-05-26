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

const cloneName = "Clone"

// cloneGen creates the deep clone methods for the AST. It works by discovering the types
// that it needs to support, starting from a root interface type. While creating the clone
// method for this root interface, more types that need to be cloned are discovered.
// This continues type by type until all necessary types have been traversed.
type cloneGen struct {
	exclude []string
	file    *jen.File
}

var _ generator = (*cloneGen)(nil)

func newCloneGen(pkgname string, options *CloneOptions) *cloneGen {
	_ = "STUB: not implemented"
	return nil
}

func (c *cloneGen) addFunc(name string, code *jen.Statement) { _ = "STUB: not implemented"; return }

func (c *cloneGen) genFile(_ generatorSPI) (string, *jen.File) {
	_ = "STUB: not implemented"
	return "", nil

	// readValueOfType produces code to read the expression of type `t`, and adds the type to the todo-list
}

func (c *cloneGen) readValueOfType(t types.Type, expr jen.Code, spi generatorSPI) jen.Code {
	_ = "STUB: not implemented"
	return *new(jen.Code)
}

// any/interface{} fields have to be handled manually

// structMethod handles cloning for struct types (value, not pointer)
func (c *cloneGen) structMethod(t types.Type, _ *types.Struct, spi generatorSPI) error {
	_ = "STUB: not implemented"
	return nil
}

// sliceMethod handles cloning for slice types
func (c *cloneGen) sliceMethod(t types.Type, slice *types.Slice, spi generatorSPI) error {
	_ = "STUB: not implemented"
	return nil
}

// copySliceElement generates code to copy slice elements
func (c *cloneGen) copySliceElement(t types.Type, elType types.Type, spi generatorSPI) jen.Code {
	_ = "STUB: not implemented"
	// For unnamed basic slices (like []byte), use built-in copy
	return *new(jen.Code)
}

// For other types, iterate and clone each element

// basicMethod handles cloning for basic types (int, string, etc.)
// Basic types don't need cloning - they're copied by value
func (c *cloneGen) basicMethod(t types.Type, _ *types.Basic, spi generatorSPI) error {
	_ = "STUB: not implemented"

	// interfaceMethod handles cloning for interface types
	return nil
}

func (c *cloneGen) interfaceMethod(t types.Type, iface *types.Interface, spi generatorSPI) error {
	_ = "STUB: not implemented"
	return nil
}

// case Type: return CloneType(in)

// ptrToBasicMethod handles cloning for pointer-to-basic types
func (c *cloneGen) ptrToBasicMethod(t types.Type, _ *types.Basic, spi generatorSPI) error {
	_ = "STUB: not implemented"
	return nil
}

// ptrToOtherMethod handles cloning for pointer types
func (c *cloneGen) ptrToOtherMethod(t types.Type, ptr *types.Pointer, spi generatorSPI) error {
	_ = "STUB: not implemented"
	return nil
}

// ptrToStructMethod handles cloning for pointer-to-struct types
func (c *cloneGen) ptrToStructMethod(t types.Type, strct *types.Struct, spi generatorSPI) error {
	_ = "STUB: not implemented"
	return nil
}

// If this type is in the exclude list, just return the original

// Skip basic types (copied in shallow copy) and private fields

// out.Field = CloneType(n.Field)

// Shallow copy: out := *n

// Deep clone all non-basic fields

// Helper functions

func ifNilReturnNil(id string) *jen.Statement { _ = "STUB: not implemented"; return nil }

func isNamed(t types.Type) bool { _ = "STUB: not implemented"; return false }

func isBasic(t types.Type) bool { _ = "STUB: not implemented"; return false }
