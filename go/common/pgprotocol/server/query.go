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

package server

import (
	"github.com/multigres/multigres/go/common/mterrors"
	"github.com/multigres/multigres/go/common/pgprotocol/protocol"
	"github.com/multigres/multigres/go/common/sqltypes"
	"github.com/multigres/multigres/go/pb/query"
)

// readQueryMessage reads a 'Q' (Query) message from the connection.
// The message format is:
//   - Message type: 'Q' (already read)
//   - Length: int32 (includes length field itself)
//   - Query string: null-terminated string
func (c *Conn) readQueryMessage() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Verify null terminator.

// Convert to string (excluding null terminator). The string()
// conversion copies, so the body buffer can safely be returned
// to the pool by the deferred returnReadBuffer.

// writeParameterDescription writes a 't' (ParameterDescription) message.
// Format:
//   - Type: 't'
//   - Length: int32
//   - Parameter count: int16
//   - For each parameter:
//   - Type OID: int32
func (c *Conn) writeParameterDescription(params []*query.ParameterDescription) error {
	_ = "STUB: not implemented"
	return nil
}

// writeRowDescription writes a 'T' (RowDescription) message.
// Format:
//   - Type: 'T'
//   - Length: int32
//   - Field count: int16
//   - For each field:
//   - Field name: null-terminated string
//   - Table OID: int32
//   - Column number: int16
//   - Type OID: int32
//   - Type size: int16
//   - Type modifier: int32
//   - Format code: int16 (0=text, 1=binary)
func (c *Conn) writeRowDescription(fields []*query.Field) error {
	_ = "STUB: not implemented"
	// field count
	return nil
}

// name + null terminator

// writeDataRow writes a 'D' (DataRow) message.
// Format:
//   - Type: 'D'
//   - Length: int32
//   - Column count: int16
//   - For each column:
//   - Value length: int32 (-1 for NULL)
//   - Value bytes: []byte (if not NULL)
func (c *Conn) writeDataRow(row *sqltypes.Row) error {
	_ = "STUB: not implemented"
	// column count
	return nil
}

// writeCommandComplete writes a 'C' (CommandComplete) message.
// Format:
//   - Type: 'C'
//   - Length: int32
//   - Command tag: null-terminated string (e.g., "SELECT 5", "INSERT 0 1")
func (c *Conn) writeCommandComplete(tag string) error { _ = "STUB: not implemented"; return nil }

// writeReadyForQuery writes a 'Z' (ReadyForQuery) message.
// Format:
//   - Type: 'Z'
//   - Length: int32 (always 5)
//   - Transaction status: byte ('I', 'T', or 'E')
func (c *Conn) writeReadyForQuery() error { _ = "STUB: not implemented"; return nil }

// writeEmptyQueryResponse writes an 'I' (EmptyQueryResponse) message.
// This is sent when the client sends an empty query string.
// Format:
//   - Type: 'I'
//   - Length: int32 (always 4)
func (c *Conn) writeEmptyQueryResponse() error { _ = "STUB: not implemented"; return nil }

// writeNoticeResponse writes an 'N' (NoticeResponse) message.
// Format is identical to ErrorResponse but with different severity levels.
func (c *Conn) writeNoticeResponse(diag *mterrors.PgDiagnostic) error {
	_ = "STUB: not implemented"
	return nil
}

// writeError writes an error response to the client.
// It handles both PostgreSQL errors (preserving all diagnostic fields)
// and generic errors (creating synthetic PgDiagnostic).
func (c *Conn) writeError(err error) error { _ = "STUB: not implemented"; return nil }

// Extract root cause - handles wrapped errors

// Check if root cause is a PostgreSQL error

// Generic error: use outer message for context

// writePgDiagnosticResponse writes a PostgreSQL diagnostic response (error or notice).
// The msgType should be MsgErrorResponse ('E') or MsgNoticeResponse ('N').
// This unified function handles all 14 PostgreSQL diagnostic fields.
func (c *Conn) writePgDiagnosticResponse(msgType byte, diag *mterrors.PgDiagnostic) error {
	_ = "STUB: not implemented"
	return nil
}

// writeErrorOrNotice writes an error or notice message with the given fields.
//
// PostgreSQL field order (S, V, C, M, D, H, P, p, q, W, s, t, c, d, n, F, L, R)
// is preserved across the bodyLen pre-pass and the in-place encoding so the
// buffer is sized to exactly the bytes encoded.
var diagFieldOrder = [...]byte{
	protocol.FieldSeverity,
	protocol.FieldSeverityV,
	protocol.FieldCode,
	protocol.FieldMessage,
	protocol.FieldDetail,
	protocol.FieldHint,
	protocol.FieldPosition,
	protocol.FieldInternalPosition,
	protocol.FieldInternalQuery,
	protocol.FieldWhere,
	protocol.FieldSchema,
	protocol.FieldTable,
	protocol.FieldColumn,
	protocol.FieldDataType,
	protocol.FieldConstraint,
	protocol.FieldFile,
	protocol.FieldLine,
	protocol.FieldRoutine,
}

func (c *Conn) writeErrorOrNotice(msgType byte, fields map[byte]string) error {
	_ = "STUB: not implemented"
	// trailing null terminator
	return nil
}

// type + value + null

// WriteCopyInResponse writes a CopyInResponse ('G') message to the client
// This tells the client that the server is ready to receive COPY data
func (c *Conn) WriteCopyInResponse(format int16, columnFormats []int16) error {
	_ = "STUB: not implemented"
	return nil
}

// ReadCopyDataMessage reads a CopyData ('d') message body
// The message type byte has already been read
// length is the body length (already has 4 subtracted by ReadMessageLength)
func (c *Conn) ReadCopyDataMessage(length int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadCopyDoneMessage reads a CopyDone ('c') message
// The message type byte has already been read
// CopyDone has no body, just validates the length
// length is the body length (already has 4 subtracted by ReadMessageLength)
func (c *Conn) ReadCopyDoneMessage(length int) error {
	_ = "STUB: not implemented"
	// CopyDone has no body, so length should be 0
	return nil
}

// ReadCopyFailMessage reads a CopyFail ('f') message
// The message type byte has already been read
// Returns the error message string from the client
// length is the body length (already has 4 subtracted by ReadMessageLength)
func (c *Conn) ReadCopyFailMessage(length int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Message should be null-terminated
