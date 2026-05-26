// Copyright 2023 The Vitess Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Modifications Copyright 2025 Supabase, Inc.

package debug

import (
	"net/http"

	"github.com/multigres/multigres/go/tools/viperutil"
)

// HandlerFunc returns an http.HandlerFunc that renders the combined config
// registry (both static and dynamic) for debugging purposes.
//
// Example requests:
//   - GET /debug/config
//   - GET /debug/config?format=json
func HandlerFunc(reg *viperutil.Registry) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// Collect command-line flags

// Handle default format (debug text)

// should not happen

// Handle JSON format specially to include both cmdline flags and viper config
