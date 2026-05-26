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

// ReadMessageType reads a single byte message type from the connection.
// Returns 0 and io.EOF if the connection is closed gracefully.
//
// Uses bufferedReader.ReadByte (concrete-type method, no interface
// dispatch) so the compiler can prove nothing escapes — vs the
// previous io.ReadFull(io.Reader, []byte) shape, which forced the
// stack-local slice header to the heap on every call.
func (c *Conn) ReadMessageType() (byte, error) { _ = "STUB: not implemented"; return 0, nil }

// ReadMessageLength reads the 4-byte message length from the connection.
// The length includes itself but excludes the message type byte.
// Returns the length of the message body (length - 4).
//
// Uses Peek + Discard to read the 4 bytes directly out of bufio's
// internal buffer with no intermediate slice — same escape-avoidance
// reason as ReadMessageType. Peek returns a slice into bufio's
// storage; we read it before Discarding the bytes so the buffer
// can reuse the space.
func (c *Conn) ReadMessageLength() (int, error) { _ = "STUB: not implemented"; return 0, nil }

// readMessageBody reads the message body of the given length.
// Returns a slice that must be released by calling returnReadBuffer.
//
// On the pool path, the underlying *[]byte is stashed in
// c.inboundPoolBuf so returnReadBuffer can return it without taking
// the address of a stack-local slice (which would force the slice
// header to the heap on every call). On the make-fallback path
// (no listener), there's nothing to recycle and inboundPoolBuf
// stays nil.
//
// If a previous body buffer is still held by inboundPoolBuf (the
// caller's defer hasn't fired, or readAndDispatchStartup recursed
// through SSL/GSS negotiation between calls), it is returned to the
// pool first to prevent a leak.
func (c *Conn) readMessageBody(length int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// returnReadBuffer releases the buffer held by inboundPoolBuf back
// to the listener bufpool. No-op when readMessageBody used the make
// fallback (inboundPoolBuf is nil) or returned a zero-length slice.
func (c *Conn) returnReadBuffer() { _ = "STUB: not implemented"; return }

// returnOutboundBuffer releases the buffer held by outboundPoolBuf
// back to the listener bufpool. Normally writePacket releases it via
// defer; this method exists for defensive cleanup on Conn.Close so a
// panic during body encoding (between startPacket and writePacket)
// doesn't strand the pool buffer. Does NOT touch bufMu — Close runs
// after concurrent access has stopped, and the panicking goroutine
// still holds bufMu unrelinquished, so re-locking would deadlock.
func (c *Conn) returnOutboundBuffer() { _ = "STUB: not implemented"; return }

// readStartupPacket reads a startup packet (no message type byte).
// Startup packets only have a length field followed by the body.
func (c *Conn) readStartupPacket() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// writeMessage writes a complete message with type, length, and body.
// The length is calculated automatically (includes length field, excludes
// type byte).
//
// Routes through startPacket/writePacket so the header + body land in
// the bufferedWriter as a single Write (a self-copy on the fast path).
// Used by the few callers that already have the body materialized as
// a []byte — body-less control messages like ParseComplete /
// BindComplete / NoData / CloseComplete pass nil here.
func (c *Conn) writeMessage(msgType byte, body []byte) error { _ = "STUB: not implemented"; return nil }

// startPacket reserves space for a single pgwire packet of the given
// body length, with the message type and 4-byte length header pre-
// written. The returned slice has length 5+bodyLen and pos points at
// the first body byte. Callers encode the body in-place via writeXxxAt
// then call writePacket exactly once.
//
// startPacket acquires bufMu and writePacket releases it. The lock is
// held across body encoding so an interleaved write from the async
// notification pusher cannot split the packet.
//
// Fast path: when a bufferedWriter is attached and the packet fits in
// its currently-available capacity, the body is written directly into
// the buffered writer's internal byte slice. writePacket then commits
// the bytes in place — no copy through a separate buffer, no buffer
// pool round-trip, no per-message allocation. This mirrors how postgres
// builds messages straight into PqSendBuffer.
//
// Fallback: when the body wouldn't fit (large packets) or no
// bufferedWriter is set yet (pre-startup writes), borrow a buffer from
// the listener's bufpool — the same pool the read path uses to stage
// inbound message bodies. The read and write sides see the same byte
// stream (multipooler's reads from postgres become multigateway's
// writes to the client), so any workload that exercises the read pool
// will exercise the write pool at a similar rate; routing both sides
// through one pool means a workload of large packets allocates once
// and then reuses the buffer indefinitely. writePacket returns the
// buffer to the pool. We do not keep a per-connection scratch buffer:
// the pool is per-listener, so memory amortizes across all
// connections rather than being pinned per-connection.
func (c *Conn) startPacket(msgType byte, bodyLen int) ([]byte, int) {
	_ = "STUB: not implemented"
	return nil, 0
}

// writePacket commits the packet started by startPacket. Releases
// bufMu (which was acquired by startPacket) and returns any pool
// buffer the slow path borrowed.
//
// On the fast path, buf aliases the bufferedWriter's internal storage,
// so bufio.Writer.Write performs a self-copy that is effectively a no-
// op and just advances the internal write cursor. On the slow path,
// this is a normal Write of the pool-backed slice; afterwards the
// pool buffer is returned to listener.bufPool for reuse.
//
// pos is the cursor returned by the final writeXxxAt encoder; it must
// equal len(buf), i.e. the bodyLen passed to startPacket must exactly
// match the bytes encoded. A mismatch panics — sizing bugs surface
// loudly in tests instead of producing truncated/garbage packets on
// the wire.
//
// Cleanup runs via defer so a panic during body encoding still releases
// bufMu and returns the pool buffer.
func (c *Conn) writePacket(buf []byte, pos int) error { _ = "STUB: not implemented"; return nil }

// In-place packet body encoders. Each writes at buf[pos:] and returns the
// new position. Callers must size the buffer (via startPacket) so that
// these never run off the end — out-of-range slice writes will panic.

func writeByteAt(buf []byte, pos int, b byte) int { _ = "STUB: not implemented"; return 0 }

func writeInt16At(buf []byte, pos int, v int16) int { _ = "STUB: not implemented"; return 0 }

func writeInt32At(buf []byte, pos int, v int32) int { _ = "STUB: not implemented"; return 0 }

func writeUint32At(buf []byte, pos int, v uint32) int { _ = "STUB: not implemented"; return 0 }

// writeStringAt writes s followed by a single null terminator.
//
// Caller is responsible for ensuring s contains no embedded NULs.
// pgwire strings are NUL-terminated, so an embedded NUL would
// truncate the field on the receiver and the bodyLen pre-pass would
// silently mis-frame the packet on the wire. Inputs that come from
// untrusted sources (query text, identifiers) should be validated by
// the caller.
func writeStringAt(buf []byte, pos int, s string) int { _ = "STUB: not implemented"; return 0 }

// writeBytesAt writes raw bytes (no terminator).
func writeBytesAt(buf []byte, pos int, b []byte) int { _ = "STUB: not implemented"; return 0 }

// writeRawByte writes a single byte that is NOT a pgwire packet —
// just one raw byte on the wire. Used only for the SSL/GSSENC
// negotiation response ('S' or 'N'), which has no length prefix and
// no message-type framing. Goes through the bufferedWriter if one is
// attached, otherwise straight to the conn.
func (c *Conn) writeRawByte(b byte) error { _ = "STUB: not implemented"; return nil }

// MessageReader provides helper methods for reading message fields.
type MessageReader struct {
	buf []byte
	pos int
}

// NewMessageReader creates a new message reader for the given buffer.
//
// Returns by value (not pointer). Callers do `r := NewMessageReader(body)`
// and the struct lives on their stack — no heap allocation. The
// methods take pointer receivers to mutate `pos`; Go auto-addresses
// the stack-local value when calling them.
//
// IMPORTANT: helpers that walk the SAME MessageReader as their
// caller MUST accept `*MessageReader`, not `MessageReader` by value.
// A by-value parameter copies the struct, and any cursor mutations
// inside the helper vanish on return — silently producing wrong
// reads if the caller continues parsing after the call. Pass-by-
// pointer-of-stack-local doesn't escape (the helper just calls
// methods locally), so safety is free.
func NewMessageReader(buf []byte) MessageReader {
	_ = "STUB: not implemented"
	return *new(MessageReader)
}

// Remaining returns the number of unread bytes.
func (r *MessageReader) Remaining() int { _ = "STUB: not implemented"; return 0 }

// ReadByte reads a single byte.
func (r *MessageReader) ReadByte() (byte, error) { _ = "STUB: not implemented"; return 0, nil }

// ReadUint16 reads a 16-bit unsigned integer in network byte order.
func (r *MessageReader) ReadUint16() (uint16, error) { _ = "STUB: not implemented"; return 0, nil }

// ReadUint32 reads a 32-bit unsigned integer in network byte order.
func (r *MessageReader) ReadUint32() (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

// ReadInt16 reads a 16-bit signed integer in network byte order.
func (r *MessageReader) ReadInt16() (int16, error) { _ = "STUB: not implemented"; return 0, nil }

// ReadInt32 reads a 32-bit signed integer in network byte order.
func (r *MessageReader) ReadInt32() (int32, error) { _ = "STUB: not implemented"; return 0, nil }

// ReadString reads a null-terminated string.
func (r *MessageReader) ReadString() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Skip null terminator.

// ReadBytes reads n bytes.
func (r *MessageReader) ReadBytes(n int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadByteString reads a length-prefixed string (4-byte length + data).
// Returns nil if length is -1 (NULL).
func (r *MessageReader) ReadByteString() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NULL
