// Copyright 2026 Supabase, Inc.
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

package handler

import (
	"context"
	"time"

	"github.com/multigres/multigres/go/common/parser/ast"
	"github.com/multigres/multigres/go/common/pgprotocol/server"
)

// handleReplicationCommand parses queryStr as a replication-protocol command
// and emits the corresponding stub response. Returns handled=true if the
// command was claimed (either successfully stubbed or a parse error), and
// handled=false if the query should fall through to the regular SQL path
// (e.g. SHOW, which the SQL grammar also accepts).
func (h *MultiGatewayHandler) handleReplicationCommand(
	ctx context.Context,
	conn *server.Conn,
	queryStr string,
	queryStart time.Time,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// VariableShowStmt — the SQL grammar handles SHOW, so let the
// caller continue down the regular path.

// replicationStubError maps a replication-command AST node to an op name
// and a `feature_not_supported` (SQLSTATE 0A000) error.
//
// Returns (opName, nil) for *ast.VariableShowStmt — SHOW is delegated to the
// normal SQL path by the caller and is not a stub failure case.
func replicationStubError(stmt ast.Stmt) (string, error) { _ = "STUB: not implemented"; return "", nil }

func notSupported(op string) (string, error) { _ = "STUB: not implemented"; return "", nil }
