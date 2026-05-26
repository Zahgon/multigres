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
	"github.com/multigres/multigres/go/services/multigateway/handler"
)

// ListenAction specifies what type of LISTEN/UNLISTEN action to take.
type ListenAction int

const (
	ListenActionListen ListenAction = iota
	ListenActionUnlisten
	ListenActionUnlistenAll
)

// ListenNotifyPrimitive handles LISTEN/UNLISTEN commands.
//
// In autocommit mode (outside a transaction), it updates connection state and
// syncs subscriptions via SubSync before returning success to the client.
// Inside a transaction, changes are buffered as pending and applied at COMMIT
// time by the TransactionPrimitive.
type ListenNotifyPrimitive struct {
	Action  ListenAction
	Channel string
	Query   string
}

func NewListenPrimitive(channel, query string) *ListenNotifyPrimitive {
	_ = "STUB: not implemented"
	return nil
}

func NewUnlistenPrimitive(channel, query string) *ListenNotifyPrimitive {
	_ = "STUB: not implemented"
	return nil
}

func NewUnlistenAllPrimitive(query string) *ListenNotifyPrimitive {
	_ = "STUB: not implemented"
	return nil
}

func (l *ListenNotifyPrimitive) StreamExecute(
	ctx context.Context,
	exec IExecute,
	conn *server.Conn,
	state *handler.MultiGatewayConnectionState,
	_ []*ast.A_Const,
	callback func(context.Context, *sqltypes.Result) error,
) error {
	_ = "STUB: not implemented"
	// PostgreSQL truncates channel names to NAMEDATALEN-1 (63 chars).
	// We must do the same so our internal tracking matches PG's behavior.
	return nil
}

// Update connection state and, for autocommit, sync subscriptions immediately
// via SubSync so they are active before the client is told LISTEN/UNLISTEN
// succeeded. Inside a transaction, changes are buffered as pending and
// applied at COMMIT by the TransactionPrimitive.

func (l *ListenNotifyPrimitive) String() string { _ = "STUB: not implemented"; return "" }

// GetQuery returns the original SQL string.
func (l *ListenNotifyPrimitive) GetQuery() string {
	_ = "STUB: not implemented"

	// GetTableGroup returns empty string — LISTEN/UNLISTEN don't target a tablegroup.
	return ""
}

func (l *ListenNotifyPrimitive) GetTableGroup() string {
	_ = "STUB: not implemented"

	// PortalStreamExecute satisfies the Primitive interface for the
	// extended-protocol path. LISTEN/UNLISTEN take no parameters; the
	// channel name is fixed at plan time. Delegate.
	return ""
}

func (l *ListenNotifyPrimitive) PortalStreamExecute(
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
