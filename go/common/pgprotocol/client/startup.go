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
)

// startup performs the connection startup handshake.
// This includes SSL negotiation (if configured), sending the startup message,
// and handling authentication.
func (c *Conn) startup(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Honor libpq-style sslmode. Unix-socket connections always run plaintext,
	// matching libpq behavior.
	return nil
}

// Send the startup message.

// Process authentication and startup responses.

// negotiateSSL implements the libpq SSLRequest handshake and, on server
// acceptance, upgrades the underlying connection to TLS using c.config.TLSConfig.
//
// Caller guarantees c.config.SSLMode.AttemptsTLS() is true and c.config.SocketFile
// is empty. On a tolerant mode (prefer) where the server returns 'N', this
// returns nil and leaves the plaintext connection in place so startup can
// continue. Strict modes (require, verify-ca, verify-full) error on 'N'.
func (c *Conn) negotiateSSL(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// prefer: continue on plaintext.

// fall through to TLS upgrade.

// PostgreSQL may respond with an ErrorResponse if SSLRequest is
// rejected outright (e.g. unsupported protocol version on the
// server). The body is on the buffered reader following the 'E'
// message-type byte; surface a clear failure rather than trying
// to parse it here.

// Server accepted TLS — upgrade in place. The buffered reader must not
// hold any post-'S' bytes, since they would belong to the TLS record
// stream and not the plaintext connection.

// sendStartupMessage sends the startup message to the server.
//
// The startup packet has no message type byte — only a 4-byte length
// (including itself) followed by the body — so it's encoded inline
// here instead of through startPacket (which always writes a 5-byte
// type+length header).
//
// Called only during single-threaded connection setup, so it does not
// acquire bufMu (no other writer can race on this connection yet).
func (c *Conn) sendStartupMessage() error {
	_ = "STUB: not implemented"
	// Pre-compute body length so we can encode in place.
	return nil
}

// protocol version

// trailing null terminator for the parameter list

// length field includes itself

// Reserve a buffer: AvailableBuffer fast path, plain make slow path.
// We don't borrow from bufPool here because startup messages are
// small and one-shot per connection — pinning a 16 KB pool bucket
// for a few hundred bytes wastes memory.

// processStartupResponses processes all messages until ReadyForQuery.
func (c *Conn) processStartupResponses(ctx context.Context) error {
	_ = "STUB: not implemented"

	// Check context.
	return nil
}

// Read message.

// Process based on message type.

// Startup complete.

// Ignore notices during startup.

// handleAuthenticationRequest handles an AuthenticationRequest message.
func (c *Conn) handleAuthenticationRequest(body []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Authentication successful, nothing more to do.

// Read available SASL mechanisms.

// Check if SCRAM-SHA-256 is supported.

// Prefer SCRAM passthrough when keys are present; otherwise fall back
// to password-based SCRAM. Passthrough lets a proxy authenticate on a
// session's behalf without ever holding the plaintext password.

// handleBackendKeyData handles a BackendKeyData message.
func (c *Conn) handleBackendKeyData(body []byte) error { _ = "STUB: not implemented"; return nil }

// handleParameterStatus handles a ParameterStatus message.
func (c *Conn) handleParameterStatus(body []byte) error { _ = "STUB: not implemented"; return nil }

// handleReadyForQuery handles a ReadyForQuery message.
func (c *Conn) handleReadyForQuery(body []byte) error { _ = "STUB: not implemented"; return nil }

// writeSSLRequest writes an SSL negotiation request.
//
// SSLRequest has no message type byte — just length (8) +
// SSLRequestCode. Encoded as a single 8-byte slice into the
// bufferedWriter's available space so it goes out in one Write.
//
// Called only during single-threaded connection setup, so it does not
// acquire bufMu (no other writer can race on this connection yet).
func (c *Conn) writeSSLRequest() error {
	_ = "STUB: not implemented"
	// 8-byte request: AvailableBuffer fast path, stack-local fallback.
	// bufPool's smallest bucket is 16 KB — vastly oversized for an
	// 8-byte one-shot request — so we use a fixed-size array on the
	// slow path instead of borrowing from the pool.
	return nil
}
