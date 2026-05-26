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

package internal

import (
	"github.com/dave/jennifer/jen"
	"golang.org/x/tools/go/packages"
)

// Title returns a copy of the string s with all Unicode letters that begin words
// mapped to their Unicode title case.
func Title(s string) string { _ = "STUB: not implemented"; return "" }

// FormatJenFile formats the given *jen.File with goimports and return a slice
// of byte corresponding to the formatted file.
func FormatJenFile(file *jen.File) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// GoImports runs gofmt and goimports on the given file
func GoImports(fullPath string) error {
	_ = "STUB: not implemented"
	// Run gofmt with simplification flag
	return nil
}

// Run goimports

// SaveJenFile saves a jen.File to disk and formats it
func SaveJenFile(fullPath string, file *jen.File) error { _ = "STUB: not implemented"; return nil }

// CheckErrors checks for package loading errors, skipping generated files
func CheckErrors(loaded []*packages.Package) error { _ = "STUB: not implemented"; return nil }

// Skip generated files

// isGeneratedFile returns true if the filename is a generated file
func isGeneratedFile(filename string) bool { _ = "STUB: not implemented"; return false }
