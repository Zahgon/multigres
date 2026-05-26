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

package client

import (
	"context"

	"github.com/multigres/multigres/go/common/mterrors"
	"github.com/multigres/multigres/go/common/sqltypes"
	"github.com/multigres/multigres/go/pb/query"
)

// DefaultStreamingBatchSize is the default size threshold (in bytes) for batching
// rows during streaming. When accumulated row data exceeds this size, the batch
// is flushed via callback. This balances memory efficiency with callback overhead.
const DefaultStreamingBatchSize = 2 * 1024 * 1024 // 2MB

// queryTracingKey is the context key for query tracing configuration.
type queryTracingKey struct{}

// QueryTracingConfig holds optional configuration for query tracing.
// Spans are always created for queries; this config controls optional details.
type QueryTracingConfig struct {
	// OperationName is a semantic name for the operation (e.g., "pg_is_in_recovery").
	// This should describe what the query does, not the SQL itself.
	// If empty, "QUERY" will be used.
	OperationName string

	// IncludeQueryText enables recording the SQL query text in the span.
	//
	// SECURITY WARNING: This should ONLY be enabled for internal system queries where:
	// - The SQL is hardcoded or uses PostgreSQL system functions
	// - No user-provided data appears in the query text
	// - No PII (Personally Identifiable Information) is included
	//
	// Examples of SAFE usage (internal queries):
	// - SELECT pg_is_in_recovery()
	// - SHOW server_version
	// - SELECT setting FROM pg_settings WHERE name = 'max_connections'
	//
	// Examples of UNSAFE usage (NEVER enable for these):
	// - Any query containing user input
	// - SELECT * FROM users WHERE email = 'user@example.com'
	// - Queries with bind parameters that may contain PII
	//
	// Default: false (SQL text is never included in spans)
	IncludeQueryText bool
}

// WithQueryTracing returns a context with query tracing configuration.
// This allows callers to customize the span (e.g., set operation name, include SQL text).
// Spans are created for all queries by default; this just adds configuration.
func WithQueryTracing(ctx context.Context, config QueryTracingConfig) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// getQueryTracingConfig returns the tracing config from context.
// Returns an empty config if none is set (spans are still created).
func getQueryTracingConfig(ctx context.Context) QueryTracingConfig {
	_ = "STUB: not implemented"
	return *new(QueryTracingConfig)
}

// defaultOperationName returns a safe default operation name.
// TODO: In the future, could support SELECT, UPDATE, INSERT, DELETE based on
// parsing the first keyword from a fixed allowlist.
func defaultOperationName() string {
	_ = "STUB: not implemented"

	// Query executes a simple query and returns all results.
	// For large result sets, consider using QueryStreaming instead.
	return ""
}

func (c *Conn) Query(ctx context.Context, queryStr string) ([]*sqltypes.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Handle notice-only results (zero-buffering notice delivery).

// Accumulate notices into current result.

// Accumulate rows into the current result.

// Also accumulate any notices that came with row data.

// CommandTag being set signals the end of a result set.

// QueryStreaming executes a simple query and streams results via callback.
// The callback is invoked in a streaming fashion with batched rows:
// - Rows are accumulated until DefaultStreamingBatchSize is exceeded, then flushed with Fields
// - On CommandComplete: remaining rows + CommandTag sent together (signals end of result set)
// For small result sets, this means a single callback with Fields, Rows, and CommandTag.
// For large result sets, multiple callbacks with rows, final one includes CommandTag.
// For multi-statement queries, this pattern repeats for each statement.
//
// A span is always created for query execution with database semantic conventions.
// Use WithQueryTracing to customize the span (operation name, include SQL text).
func (c *Conn) QueryStreaming(ctx context.Context, queryStr string, callback func(ctx context.Context, result *sqltypes.Result) error) error {
	_ = "STUB: not implemented"
	// Create span for query execution
	return nil
}

// Send the Query message.

// Process responses.

// writeQueryMessage writes a 'Q' (Query) message.
func (c *Conn) writeQueryMessage(queryStr string) error { _ = "STUB: not implemented"; return nil }

// null terminator

// processQueryResponses processes all responses to a query until ReadyForQuery.
// The callback is invoked in a streaming fashion with batched rows:
// - Rows are accumulated until DefaultStreamingBatchSize is exceeded, then flushed with Fields
// - On CommandComplete: remaining rows + CommandTag sent together (signals end of result set)
// For small result sets, this means a single callback with Fields, Rows, and CommandTag.
// For large result sets, multiple callbacks with rows, final one includes CommandTag.
//
// IMPORTANT: This function always reads until ReadyForQuery to keep the connection
// in a clean state. Callback errors are captured but do not stop message processing.
// Context cancellation should be handled by the caller (e.g., by killing the query
// on the server side) rather than here, to avoid leaving unread messages on the wire.
func (c *Conn) processQueryResponses(ctx context.Context, callback func(ctx context.Context, result *sqltypes.Result) error) error {
	_ = "STUB: not implemented"
	// Track state for current result set.
	return nil
}

// Track the first error encountered. We continue processing messages to drain
// the connection, then return this error after ReadyForQuery.

// flushBatch sends accumulated rows via callback and resets the batch.
// Does not reset currentFields as they may be needed for subsequent batches.
// Captures errors but does not return them - we continue draining.

// Read message.

// Start of a new result set - parse and store fields.
// Fields will be included in the first batch callback.

// Add row to batch and track size.

// Flush batch if size threshold exceeded.

// Send final batch with CommandTag (signals end of result set).
// This combines any remaining rows with the command completion.

// Reset for next result set.

// Empty query, call callback with empty result.

// Query complete. Return any error that was captured.

// Capture the error but continue draining until ReadyForQuery.

// Stream notice immediately via callback (zero-buffering notice delivery).
// Notices are sent as separate Results with no rows or command tag.

// NotificationResponse ('A') can arrive at any time on connections that
// have issued LISTEN. Parse and deliver via callback as a notification result.

// Handle parameter status updates. Capture error but continue draining.

// Unexpected message type. Capture error but continue draining.

// parseRowDescription parses a RowDescription message.
//
// TODO: Migrate errors in parseRowDescription, parseDataRow, parseCommandComplete, and readMessage
// to use the mterrors package with proper error codes. These errors indicate connection-level failures
// (truncated/incomplete messages) and should be categorized so that isConnectionError() in
// regular_conn.go can detect them using error codes instead of string matching.
func (c *Conn) parseRowDescription(body []byte) ([]*query.Field, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// parseDataRow parses a DataRow message.
// Returns sqltypes.Row where nil values represent NULL.
func (c *Conn) parseDataRow(body []byte) (*sqltypes.Row, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// nil for NULL, []byte{} for empty string - preserved correctly

// parseCommandComplete parses a CommandComplete message.
func (c *Conn) parseCommandComplete(body []byte) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// parseRowsAffected extracts the row count from a command tag.
// Returns 0 for SELECT statements since they don't "affect" rows (they only read).
func parseRowsAffected(tag string) uint64 {
	_ = "STUB: not implemented"
	// Command tags have formats like:
	// - "SELECT 5" (5 rows returned, but not "affected" since SELECT is read-only)
	// - "INSERT 0 1" (1 row inserted)
	// - "UPDATE 10" (10 rows updated)
	// - "DELETE 3" (3 rows deleted)
	return 0
}

// SELECT doesn't affect rows, only reads them

// Find the last space-separated number.

// parseDiagnosticFields parses all 14 PostgreSQL diagnostic fields from the wire format.
// This is a shared helper used by both parseError() and parseNotice() since PostgreSQL
// uses the same field format for ErrorResponse ('E') and NoticeResponse ('N') messages.
// The msgType parameter should be protocol.MsgErrorResponse or protocol.MsgNoticeResponse.
//
// If the parsed diagnostic fails validation (missing required fields), a warning is logged
// but the diagnostic is still returned. This allows lenient handling of malformed messages.
func parseDiagnosticFields(msgType byte, body []byte) *mterrors.PgDiagnostic {
	_ = "STUB: not implemented"
	return nil
}

// End of fields.

// FieldSeverityV ('V') is the non-localized severity.
// Only use it if FieldSeverity ('S') wasn't already set.

// Validate the parsed diagnostic. Log a warning if validation fails,
// but still return the diagnostic to allow lenient handling.

// Convert single byte to string directly (msgType is 'E' or 'N')

// parseError parses an ErrorResponse message into a *mterrors.PgDiagnostic.
// Since mterrors.PgDiagnostic implements the error interface, it can be returned
// directly as an error. This eliminates the need for a separate wrapper type.
// It captures all 14 PostgreSQL error fields defined in the protocol.
func (c *Conn) parseError(body []byte) error { _ = "STUB: not implemented"; return nil }

// parseNotice parses a NoticeResponse message into a mterrors.PgDiagnostic.
func (c *Conn) parseNotice(body []byte) *mterrors.PgDiagnostic {
	_ = "STUB: not implemented"
	return nil
}

// parseNotificationResponse parses a NotificationResponse ('A') message.
// Format: Int32(pid) + String(channel) + String(payload)
func (c *Conn) parseNotificationResponse(body []byte) (*sqltypes.Notification, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SendQuery writes a simple query message without reading the response.
// The caller must arrange for response messages to be read separately.
// This enables split read/write patterns where one goroutine reads messages
// while another writes commands (e.g., PubSubListener's LISTEN/UNLISTEN).
func (c *Conn) SendQuery(sql string) error { _ = "STUB: not implemented"; return nil }

// ParseNotification parses a NotificationResponse ('A') message body.
// This is the exported counterpart of parseNotificationResponse, enabling
// callers that read raw messages to parse notifications themselves.
func (c *Conn) ParseNotification(body []byte) (*sqltypes.Notification, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WaitForNotification blocks until a NotificationResponse message is received
// from PostgreSQL. This is used by the shared PubSubListener to read async
// notifications on a dedicated listener connection.
//
// The context can be used to cancel the wait. Note that cancellation may leave
// the connection in an unusable state if a partial read occurred.
//
// Returns the parsed notification, or an error if the connection is closed,
// context is cancelled, or an unexpected message is received.
func (c *Conn) WaitForNotification(ctx context.Context) (*sqltypes.Notification, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ignore parameter status updates

// Ignore notices
