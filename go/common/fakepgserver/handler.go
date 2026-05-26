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

package fakepgserver

import (
	"context"

	"github.com/multigres/multigres/go/common/pgprotocol/server"
	"github.com/multigres/multigres/go/common/preparedstatement"
	"github.com/multigres/multigres/go/common/sqltypes"
	"github.com/multigres/multigres/go/pb/query"
)

// fakeHandler implements server.Handler to return pre-configured query results.
type fakeHandler struct {
	server *Server

	// preparedStatements stores prepared statements by name.
	// Each connection would normally have its own, but for simplicity
	// we use a shared map (tests typically use a single connection).
	preparedStatements map[string]*preparedStmt

	// portals stores bound portals by name.
	portals map[string]*portal
}

type preparedStmt struct {
	name       string
	query      string
	paramTypes []uint32
}

type portal struct {
	name       string
	stmt       *preparedStmt
	params     [][]byte
	resultFmts []int16
}

// HandleQuery handles a simple query protocol message.
func (h *fakeHandler) HandleQuery(ctx context.Context, conn *server.Conn, queryStr string, callback func(context.Context, *sqltypes.Result) error) error {
	_ = "STUB: not implemented"
	return nil
}

// AfterCallbackError: deliver the result first, then return the
// error. This simulates mid-stream failures (rows sent, then
// connection error).

// HandleParse handles a Parse message for the extended query protocol.
func (h *fakeHandler) HandleParse(ctx context.Context, conn *server.Conn, name, queryStr string, paramTypes []uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// HandleBind handles a Bind message for the extended query protocol.
func (h *fakeHandler) HandleBind(ctx context.Context, conn *server.Conn, portalName, stmtName string, params [][]byte, paramFormats, resultFormats []int16) error {
	_ = "STUB: not implemented"
	return nil
}

// HandleExecute handles an Execute message for the extended query protocol.
func (h *fakeHandler) HandleExecute(ctx context.Context, conn *server.Conn, portalName string, maxRows int32, _ bool, callback func(context.Context, *sqltypes.Result) error) error {
	_ = "STUB: not implemented"
	return nil
}

// Execute the query.

// HandleDescribe handles a Describe message.
func (h *fakeHandler) HandleDescribe(ctx context.Context, conn *server.Conn, typ byte, name string) (*query.StatementDescription, error) {
	_ = "STUB: not implemented"
	return nil,

		// Statement
		nil
}

// For the fake server, we return minimal description.
// In a real server, this would describe the statement's parameters and result columns.

// Portal

// Get the expected result to determine fields.

// Use fields directly from the result.

// HandleClose handles a Close message.
func (h *fakeHandler) HandleClose(ctx context.Context, conn *server.Conn, typ byte, name string) error {
	_ = "STUB: not implemented"
	return nil

	// Statement
}

// Portal

// HandleSync handles a Sync message.
func (h *fakeHandler) HandleSync(ctx context.Context, conn *server.Conn) error {
	_ = "STUB: not implemented"
	// Clear unnamed portal after sync (per PostgreSQL protocol).
	return nil
}

// ConnectionEstablished records the replication mode of the freshly
// authenticated connection so tests can assert on it via
// Server.LastReplicationMode.
func (h *fakeHandler) ConnectionEstablished(conn *server.Conn) { _ = "STUB: not implemented"; return }

// ConnectionClosed handles connection cleanup.
func (h *fakeHandler) ConnectionClosed(conn *server.Conn) {
	_ = "STUB: not implemented"

	// GetPreparedStatementInfo returns nil — fakepgserver does not manage
	// gateway-level prepared statement consolidation.
	return
}

func (h *fakeHandler) GetPreparedStatementInfo(connID uint32, name string) *preparedstatement.PreparedStatementInfo {
	_ = "STUB: not implemented"

	// Ensure fakeHandler implements server.Handler.
	return nil
}

var _ server.Handler = (*fakeHandler)(nil)
