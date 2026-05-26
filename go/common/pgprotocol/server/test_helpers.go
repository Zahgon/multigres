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

package server

import (
	"bytes"
	"net"
	"time"
)

// TestConn is a test-only wrapper around Conn that provides access to the write buffer.
type TestConn struct {
	*Conn
	WriteBuf *bytes.Buffer
}

// exportedTestNetConn is a minimal implementation of net.Conn for testing.
type exportedTestNetConn struct {
	readBuf  *bytes.Buffer
	writeBuf *bytes.Buffer
}

func (m *exportedTestNetConn) Read(b []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *exportedTestNetConn) Write(b []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *exportedTestNetConn) Close() error { _ = "STUB: not implemented"; return nil }
func (m *exportedTestNetConn) LocalAddr() net.Addr {
	_ = "STUB: not implemented"
	return *new(net.Addr)
}
func (m *exportedTestNetConn) RemoteAddr() net.Addr {
	_ = "STUB: not implemented"
	return *new(net.Addr)
}
func (m *exportedTestNetConn) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }
func (m *exportedTestNetConn) SetReadDeadline(t time.Time) error {
	_ = "STUB: not implemented"
	return nil
}
func (m *exportedTestNetConn) SetWriteDeadline(t time.Time) error {
	_ = "STUB: not implemented"

	// TestConnOption configures a TestConn.
	return nil
}

type TestConnOption func(*Conn)

// WithTestHandler sets the handler on a test connection.
func WithTestHandler(h Handler) TestConnOption {
	_ = "STUB: not implemented"
	return *new(TestConnOption)
}

// WithTestDatabase sets the database name on a test connection.
func WithTestDatabase(db string) TestConnOption {
	_ = "STUB: not implemented"
	return *new(TestConnOption)
}

// WithTestReplicationMode sets the replication mode on a test connection.
func WithTestReplicationMode(mode ReplicationMode) TestConnOption {
	_ = "STUB: not implemented"
	return *new(TestConnOption)
}

// NewTestConn creates a Conn suitable for testing.
// readBuf contains data that will be read by the Conn (simulating client input).
// The returned TestConn includes WriteBuf to inspect what was written.
func NewTestConn(readBuf *bytes.Buffer, opts ...TestConnOption) *TestConn {
	_ = "STUB: not implemented"
	return nil
}

// caller can cancel via conn.Close() if needed

// WriteCopyDataMessage writes a CopyData message to the buffer.
// This simulates a client sending COPY data.
func WriteCopyDataMessage(buf *bytes.Buffer, data []byte) { _ = "STUB: not implemented"; return }

// length includes itself

// WriteCopyDoneMessage writes a CopyDone message to the buffer.
// This simulates a client signaling end of COPY data.
func WriteCopyDoneMessage(buf *bytes.Buffer) { _ = "STUB: not implemented"; return }

// length includes itself, no body

// WriteCopyFailMessage writes a CopyFail message to the buffer.
// This simulates a client aborting a COPY operation with an error message.
func WriteCopyFailMessage(buf *bytes.Buffer, errMsg string) { _ = "STUB: not implemented"; return }

// null-terminated
// length includes itself
