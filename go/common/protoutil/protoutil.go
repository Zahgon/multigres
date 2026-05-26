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

// Package protoutil provides helper functions for creating proto messages.
package protoutil

import (
	"github.com/multigres/multigres/go/pb/query"
)

// NewPreparedStatement creates a new PreparedStatement proto message.
func NewPreparedStatement(name, queryStr string, paramTypes []uint32) *query.PreparedStatement {
	_ = "STUB: not implemented"
	return nil
}

// NewPortal creates a new Portal proto message.
// paramFormats and resultFormats use int16 for compatibility with PostgreSQL wire protocol,
// but are converted to int32 for proto serialization.
// params uses Vitess-style encoding: nil = NULL, []byte{} = empty string.
func NewPortal(name, preparedStatementName string, params [][]byte, paramFormats, resultFormats []int16) *query.Portal {
	_ = "STUB: not implemented"
	// Convert int16 slices to int32 for proto type.
	return nil
}

// Encode params using Vitess-style encoding (lengths + concatenated values).

// TargetEquals checks that the two specified targets are equal or not.
func TargetEquals(t1, t2 *query.Target) bool { _ = "STUB: not implemented"; return false }
