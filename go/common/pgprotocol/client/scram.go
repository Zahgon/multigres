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
	"github.com/multigres/multigres/go/common/pgprotocol/scram"
)

// scramClient handles the SCRAM-SHA-256 authentication flow over a connection.
// It wraps scram.SCRAMClient and adds protocol I/O handling.
type scramClient struct {
	conn   *Conn
	client *scram.SCRAMClient
}

// newScramClient creates a new SCRAM client using password-based authentication.
func newScramClient(conn *Conn, username, password string) *scramClient {
	_ = "STUB: not implemented"
	return nil
}

// newScramClientWithKeys creates a SCRAM client using pre-computed keys.
// This enables SCRAM passthrough authentication where keys were extracted during
// client authentication and are reused for backend authentication.
func newScramClientWithKeys(conn *Conn, username string, clientKey, serverKey []byte) *scramClient {
	_ = "STUB: not implemented"
	return nil
}

// authenticate performs the full SCRAM-SHA-256 authentication exchange.
func (s *scramClient) authenticate() error {
	_ = "STUB: not implemented"
	// Step 1: Generate and send client-first message.
	return nil
}

// Step 2: Receive and process server-first message.

// Step 3: Generate and send client-final message.

// Step 4: Receive and verify server-final message.

// sendClientFirst sends the SASLInitialResponse with client-first message.
//
// Called only during single-threaded connection setup, so it does not
// acquire bufMu (no other writer can race on this connection yet).
func (s *scramClient) sendClientFirst() error { _ = "STUB: not implemented"; return nil }

// Send SASLInitialResponse message:
//   - mechanism (null-terminated string)
//   - client-first-message length (int32)
//   - client-first-message bytes

// receiveServerFirst receives and parses the AuthenticationSASLContinue message.
func (s *scramClient) receiveServerFirst() (string, error) {
	_ = "STUB: not implemented"
	// Read message from server.
	return "", nil
}

// Check message type.

// Parse authentication request.

// Get server-first-message (all remaining bytes).

// sendClientFinal computes the proof and sends the client-final message.
//
// Called only during single-threaded connection setup, so it does not
// acquire bufMu (no other writer can race on this connection yet).
func (s *scramClient) sendClientFinal(serverFirst string) error {
	_ = "STUB: not implemented"
	return nil
}

// Send SASLResponse message: just the client-final-message bytes.

// receiveServerFinal receives and verifies the AuthenticationSASLFinal message.
func (s *scramClient) receiveServerFinal() error {
	_ = "STUB: not implemented"
	// Read message from server.
	return nil
}

// Check message type.

// Parse authentication request.

// Get server-final-message (all remaining bytes).

// Verify server signature.
