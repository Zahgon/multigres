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

package engine

import (
	"context"

	"github.com/multigres/multigres/go/common/parser/ast"
	"github.com/multigres/multigres/go/common/pgprotocol/server"
	"github.com/multigres/multigres/go/common/preparedstatement"
	"github.com/multigres/multigres/go/common/sqltypes"
	"github.com/multigres/multigres/go/pb/query"
	"github.com/multigres/multigres/go/services/multigateway/handler"
)

// TempTableRoute routes a query that creates temporary objects through a
// reserved connection. It sets PendingTempTableReservation on the state
// so that ScatterConn's StreamExecute creates a reserved connection with
// ReasonTempTable, following the same pattern as transactions.
type TempTableRoute struct {
	TableGroup string
	Shard      string
	Query      string

	// PreparedStatement, if set, is a gateway-managed prepared statement to
	// ensure on the backend connection before Query runs. This lets
	// `CREATE TEMP TABLE t AS EXECUTE p` work: the planner rewrites the
	// inner EXECUTE to the canonical name and attaches the metadata here.
	PreparedStatement *query.PreparedStatement
}

// NewTempTableRoute creates a new TempTableRoute primitive.
func NewTempTableRoute(tableGroup, shard, sql string) *TempTableRoute {
	_ = "STUB: not implemented"
	return nil
}

// NewTempTableRouteWithPreparedStatement creates a TempTableRoute that
// carries a gateway-managed prepared statement. See
// TempTableRoute.PreparedStatement for details.
func NewTempTableRouteWithPreparedStatement(tableGroup, shard, sql string, ps *query.PreparedStatement) *TempTableRoute {
	_ = "STUB: not implemented"
	return nil
}

// StreamExecute sets the temp table reservation flag and delegates to
// StreamExecute. ScatterConn will see the flag and create a reserved
// connection with ReasonTempTable.
func (t *TempTableRoute) StreamExecute(
	ctx context.Context,
	exec IExecute,
	conn *server.Conn,
	state *handler.MultiGatewayConnectionState,
	_ []*ast.A_Const,
	callback func(context.Context, *sqltypes.Result) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

// GetTableGroup returns the target tablegroup.
// PortalStreamExecute satisfies the Primitive interface for the
// extended-protocol path. SELECT INTO and CREATE TEMP TABLE are
// excluded from isCacheable, so the executor never lands here via the
// cacheable portal branch — but keeping the implementation real (rather
// than panicking) means a future caller composing TempTableRoute into
// a Sequence won't silently lose work. Delegate.
func (t *TempTableRoute) PortalStreamExecute(
	ctx context.Context,
	exec IExecute,
	conn *server.Conn,
	state *handler.MultiGatewayConnectionState,
	_ *preparedstatement.PortalInfo,
	_ int32,
	_ bool,
	callback func(context.Context, *sqltypes.Result) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *TempTableRoute) GetTableGroup() string {
	_ = "STUB: not implemented"

	// GetQuery returns the SQL query.
	return ""
}

func (t *TempTableRoute) GetQuery() string {
	_ = "STUB: not implemented"

	// String returns a description of the primitive for debugging.
	return ""
}

func (t *TempTableRoute) String() string { _ = "STUB: not implemented"; return "" }

// Ensure TempTableRoute implements Primitive interface.
var _ Primitive = (*TempTableRoute)(nil)
