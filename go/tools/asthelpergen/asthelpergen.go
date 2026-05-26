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

// Package asthelpergen provides code generation for AST (Abstract Syntax Tree) helper methods.
//
// This package automatically generates helper methods for AST nodes including:
//   - Deep cloning (Clone methods)
//   - AST rewriting/transformation (Rewrite methods)
//
// The generator works by discovering all types that implement a root interface and
// then generating the appropriate helper methods for each type using a plugin architecture.
//
// Usage:
//
//	result, err := asthelpergen.GenerateASTHelpers(&asthelpergen.Options{
//	    Packages:      []string{"./mypackage"},
//	    RootInterface: "github.com/multigres/multigres/go/common/parser/ast.Node",
//	})
//
// The generated code follows Go conventions and includes proper error handling,
// nil checks, and type safety.
package asthelpergen

import (
	"go/types"

	"github.com/dave/jennifer/jen"
	"golang.org/x/tools/go/packages"
)

const (
	// License header for generated files (each line will be prefixed with //)
	licenseFileHeader = `Copyright 2025 Supabase, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.`
)

// addLicenseHeader adds the license header as single-line // comments before the package declaration
func addLicenseHeader(file *jen.File) { _ = "STUB: not implemented"; return }

type (
	// generatorSPI provides services to individual generators during code generation.
	// It acts as a service provider interface, giving generators access to type discovery
	// and scope information needed for generating helper methods.
	generatorSPI interface {
		// addType adds a newly discovered type to the processing queue
		addType(t types.Type)
		// scope returns the type scope for finding implementations
		scope() *types.Scope
		// findImplementations finds all types that implement the given interface
		findImplementations(iff *types.Interface, impl func(types.Type) error) error
		// iface returns the root interface that all nodes are expected to implement
		iface() *types.Interface
	}

	// generator defines the interface that all specialized generators must implement.
	// Each generator handles specific types of Go constructs (structs, interfaces, etc.)
	// and produces the appropriate helper methods for those types.
	generator interface {
		// genFile generates the final output file for this generator
		genFile(generatorSPI) (string, *jen.File)
		// interfaceMethod handles interface types with type switching logic
		interfaceMethod(t types.Type, iface *types.Interface, spi generatorSPI) error
		// structMethod handles struct types with field iteration
		structMethod(t types.Type, strct *types.Struct, spi generatorSPI) error
		// ptrToStructMethod handles pointer-to-struct types
		ptrToStructMethod(t types.Type, strct *types.Struct, spi generatorSPI) error
		// ptrToBasicMethod handles pointer-to-basic types (e.g., *int, *string)
		ptrToBasicMethod(t types.Type, basic *types.Basic, spi generatorSPI) error
		// sliceMethod handles slice types with element processing
		sliceMethod(t types.Type, slice *types.Slice, spi generatorSPI) error
		// basicMethod handles basic types (int, string, bool, etc.)
		basicMethod(t types.Type, basic *types.Basic, spi generatorSPI) error
	}

	// astHelperGen is the main orchestrator that coordinates the code generation process.
	// It discovers implementations of a root interface and uses multiple specialized
	// generators to produce helper methods for all discovered types.
	astHelperGen struct {
		// DebugTypes enables debug output for type processing
		DebugTypes bool
		// mod is the Go module information for path resolution
		mod *packages.Module
		// sizes provides platform-specific type size information
		sizes types.Sizes
		// namedIface is the root interface type for which helpers are generated
		namedIface *types.Named
		// _iface is the underlying interface type
		_iface *types.Interface
		// gens is the list of specialized generators (clone, rewrite, etc.)
		gens []generator

		// _scope is the type scope for finding implementations
		_scope *types.Scope
		// todo is the queue of types that need to be processed
		todo []types.Type
	}
)

func (gen *astHelperGen) iface() *types.Interface {
	_ = "STUB: not implemented"

	// newGenerator creates a new AST helper generator with the specified configuration.
	//
	// Parameters:
	//   - mod: Go module information for path resolution
	//   - sizes: Platform-specific type size information
	//   - named: The root interface type for which helpers will be generated
	//   - generators: Specialized generators for different helper types (clone, rewrite, etc.)
	//
	// Returns a configured astHelperGen ready to process types and generate code.
	return nil
}

func newGenerator(mod *packages.Module, sizes types.Sizes, named *types.Named, generators ...generator) *astHelperGen {
	_ = "STUB: not implemented"
	return nil
}

func findImplementations(scope *types.Scope, iff *types.Interface, impl func(types.Type) error) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if type implements the interface directly

// Check if pointer to type implements the interface

func (gen *astHelperGen) findImplementations(iff *types.Interface, impl func(types.Type) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (gen *astHelperGen) scope() *types.Scope { _ = "STUB: not implemented"; return nil }

func (gen *astHelperGen) addType(t types.Type) { _ = "STUB: not implemented"; return }

// GenerateCode is the main loop where we build up the code per file.
func (gen *astHelperGen) GenerateCode() (map[string]*jen.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// VerifyFilesOnDisk compares the generated results from the codegen against the files that
// currently exist on disk and returns any mismatches
func VerifyFilesOnDisk(result map[string]*jen.File) (errors []error) {
	_ = "STUB: not implemented"
	return nil
}

func (gen *astHelperGen) createFiles() (map[string]*jen.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// processTypeQueue processes all types in the todo queue with all generators
func (gen *astHelperGen) processTypeQueue() error { _ = "STUB: not implemented"; return nil }

// processTypeWithGenerators dispatches a type to all generators based on its underlying type
func (gen *astHelperGen) processTypeWithGenerators(t types.Type) error {
	_ = "STUB: not implemented"
	return nil
}

// handlePointerType handles pointer types by dispatching to the appropriate method
func (gen *astHelperGen) handlePointerType(t types.Type, ptr *types.Pointer, g generator) error {
	_ = "STUB: not implemented"
	return nil
}

// generateOutputFiles collects the generated files from all generators
func (gen *astHelperGen) generateOutputFiles() map[string]*jen.File {
	_ = "STUB: not implemented"
	return nil
}

// noQualifier is used to print types without package qualifiers
var noQualifier = func(*types.Package) string { return "" }

// printableTypeName returns a string that can be used as a valid golang identifier
func printableTypeName(t types.Type) string { _ = "STUB: not implemented"; return "" }

// Options configures the AST helper generation process.
type Options struct {
	// Packages specifies the Go packages to analyze for AST types.
	// Can be package paths like "./mypackage" or import paths like "github.com/example/ast".
	Packages []string

	// RootInterface is the fully qualified name of the root interface that all AST nodes implement.
	// Format: "package.path.InterfaceName" (e.g., "github.com/multigres/multigres/go/common/parser/ast.Node")
	RootInterface string

	// Clone configures the clone generator options
	Clone CloneOptions

	// Rewrite configures the rewrite generator options (currently no options needed)
	Rewrite RewriteOptions
}

// CloneOptions configures the clone generator
type CloneOptions struct {
	// Exclude is a list of type names to exclude from deep cloning
	Exclude []string
}

// RewriteOptions configures the rewrite generator (placeholder for future options)
type RewriteOptions struct{}

// GenerateASTHelpers is the main entry point for generating AST helper methods.
//
// It loads the specified packages, analyzes the types that implement the root interface,
// and generates comprehensive helper methods including clone and rewrite functionality.
//
// The function returns a map where keys are file paths and values are the generated
// Go source files. The caller is responsible for writing these files to disk.
//
// Returns an error if package loading fails, the root interface cannot be found,
// or code generation encounters any issues.
func GenerateASTHelpers(options *Options) (map[string]*jen.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// findTypeObject finds the types.Object for the given interface from the given scopes.
func findTypeObject(interfaceToFind string, scopes map[string]*types.Scope) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

var _ generatorSPI = (*astHelperGen)(nil)
