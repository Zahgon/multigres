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

	"github.com/multigres/multigres/go/common/parser/ast"
	"github.com/multigres/multigres/go/common/pgprotocol/server"
	"github.com/multigres/multigres/go/common/sqltypes"
)

// executeWithImplicitTransaction handles multi-statement batches by:
// - Injecting synthetic BEGIN at start and after each COMMIT/ROLLBACK
// - Tracking implicit vs explicit transaction segments
// - Auto-ROLLBACK on failure in implicit segments (matching PostgreSQL behavior)
// - Auto-COMMIT at end of implicit segments
//
// This preserves PostgreSQL's "adoption" semantics where BEGIN mid-batch adopts
// preceding statements into the explicit transaction.
//
// If already in a transaction, it continues without injecting BEGIN but still
// handles mid-batch COMMIT/ROLLBACK that may start new implicit segments.
//
// # CommandComplete deferral for the last statement
//
// When PostgreSQL executes a multi-statement simple query like "S1; S2; S3;",
// it wraps them in an implicit transaction. It commits the transaction *before*
// sending the last statement's CommandComplete. If the commit fails (e.g.,
// deferred constraint violation, serialization failure), the last statement's
// CommandComplete is replaced by an ErrorResponse:
//
//	→ CommandComplete (S1)
//	→ CommandComplete (S2)
//	→ ErrorResponse         ← replaces S3's CommandComplete; commit failed
//	→ ReadyForQuery 'I'     ← back to idle, nothing committed
//
// We replicate this behavior by intercepting the callback for the last statement
// in an implicit transaction. DataRow and RowDescription messages stream normally
// (we can't and shouldn't hold back potentially millions of rows), but the
// CommandComplete (identified by a non-empty CommandTag on the Result) is held
// until the commit succeeds. If the commit fails, the held CommandComplete is
// discarded and the client sees an ErrorResponse instead — exactly matching
// PostgreSQL's wire protocol behavior.
//
// See exec_simple_query() in postgres.c (line ~1290) for PostgreSQL's implementation.
func (h *MultiGatewayHandler) executeWithImplicitTransaction(
	ctx context.Context,
	conn *server.Conn,
	state *MultiGatewayConnectionState,
	queryStr string,
	stmts []ast.Stmt,
	callback func(ctx context.Context, result *sqltypes.Result) error,
) error {
	_ = "STUB: not implemented"
	// recordTableMetrics emits per-table metrics for a single statement's ExecuteResult.
	return nil
}

// silentExecute runs a statement without sending results to the client.
// Used for synthetic BEGIN/COMMIT/ROLLBACK injected by implicit transaction handling.

// If already in a transaction, don't inject BEGIN at start

// heldCommandTag stores the CommandTag (e.g., "SELECT 42", "INSERT 0 1") from the
// last statement's final callback when we're in an implicit transaction. We hold it
// here instead of sending it immediately so we can attempt the commit first. If the
// commit succeeds, we flush it to the client. If the commit fails, we discard it —
// the client sees an ErrorResponse instead of a CommandComplete.

// Inject BEGIN if needed (start of batch or after COMMIT/ROLLBACK)

// User's BEGIN - skip execution (we already started), mark as explicit.
// Still send a result to the client so the response count matches their statements.
//
// If the user's BEGIN has transaction options (e.g., ISOLATION LEVEL SERIALIZABLE),
// we must preserve them. Two cases:
//   - PendingBeginQuery is set (deferred, no backend call yet): replace it with
//     the user's full BEGIN so the options are sent when the connection is created.
//   - PendingBeginQuery is empty (backend already has plain BEGIN + queries ran):
//     error out matching PostgreSQL behavior — SET TRANSACTION ISOLATION LEVEL
//     must be called before any query.

// BEGIN still deferred — adopt the user's isolation level / access mode.

// Backend transaction already has queries executed.
// PostgreSQL rejects this with SQLSTATE 25001.
// Rollback the implicit transaction before returning the error.

// User's COMMIT/ROLLBACK - execute and prepare for next segment.

// Regular statement execution.
//
// For the last statement in an implicit transaction, we use an intercepting
// callback that defers the CommandComplete. The streaming callback is invoked
// multiple times per statement:
//   - Intermediate calls: Fields + batched Rows, CommandTag="" → stream to client
//   - Final call: last batch of Rows + CommandTag="SELECT 42"  → hold CommandTag
//
// We detect the final call by checking CommandTag != "". When we see it, we
// save the CommandTag in heldCommandTag and forward the remaining rows/notices
// without it. This causes the server to write DataRow messages but NOT the
// CommandComplete message — that waits for the commit outcome.

// Hold the CommandTag — we'll send it after a successful commit,
// or discard it if the commit fails.

// The final callback may also carry the last batch of rows and
// any notices. Forward those immediately — only the CommandComplete
// (derived from CommandTag) should be deferred.

// Intermediate callback (rows streaming) — forward as-is.

// Auto-rollback implicit transaction on failure.

// Explicit transaction: enter aborted state.
// The client must issue ROLLBACK to recover.

// Auto-commit if we ended in an implicit transaction.
//
// At this point, the last statement's DataRows have been streamed to the client,
// but its CommandComplete is held in heldCommandTag. We now attempt the commit:
//
//   Commit succeeds → flush heldCommandTag → client sees CommandComplete
//   Commit fails    → rollback, return error → client sees ErrorResponse
//
// This matches PostgreSQL's behavior in exec_simple_query(): the commit is
// attempted before the last statement's CommandComplete is sent to the client.

// Commit failed — rollback to clean up. The held CommandComplete is
// discarded; the caller will send an ErrorResponse instead.

// Commit succeeded — now send the deferred CommandComplete for the last
// statement. The client sees the full expected response sequence:
// DataRow... → CommandComplete → ReadyForQuery.
