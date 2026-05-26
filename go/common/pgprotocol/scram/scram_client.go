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

package scram

// SCRAMClient implements client-side SCRAM-SHA-256 authentication.
// It supports two modes:
// 1. Password mode: authenticate using a plaintext password
// 2. Passthrough mode: authenticate using pre-extracted ClientKey/ServerKey
//
// Passthrough mode enables SCRAM key passthrough: after a proxy verifies a client,
// it can extract the ClientKey and use it to authenticate to the backend database
// without knowing the plaintext password.
//
// When constructed with EnableChannelBinding, the client speaks
// SCRAM-SHA-256-PLUS instead, embedding the supplied tls-server-end-point
// hash into its messages.
type SCRAMClient struct {
	// Username for authentication.
	username string

	// Mode: either password or keys.
	password  string // For normal mode
	clientKey []byte // For passthrough mode
	serverKey []byte // For passthrough mode (to verify server signature)

	// channelBindingHash is the tls-server-end-point hash to bind to. When
	// non-nil the client uses gs2 flag "p=tls-server-end-point" and includes
	// the hash in the cbind data; when nil it sends "n,,".
	channelBindingHash []byte

	// State from the SCRAM exchange.
	clientNonce            string
	clientFirstMessageBare string
	gs2Header              string
	authMessage            string
}

// NewSCRAMClientWithPassword creates a SCRAM client that authenticates with a password.
// This is the standard mode where the password is used to derive SCRAM keys.
func NewSCRAMClientWithPassword(username, password string) *SCRAMClient {
	_ = "STUB: not implemented"
	return nil
}

// NewSCRAMClientWithKeys creates a SCRAM client that authenticates with extracted SCRAM keys.
// This enables SCRAM passthrough: a proxy can verify a client, extract the ClientKey,
// and use it to authenticate to PostgreSQL without the plaintext password.
//
// The clientKey and serverKey should be extracted from a previous SCRAM authentication
// using ExtractAndVerifyClientProof and the hash's ServerKey.
func NewSCRAMClientWithKeys(username string, clientKey, serverKey []byte) *SCRAMClient {
	_ = "STUB: not implemented"
	return nil
}

// clientNonceLength is the length of the client nonce in bytes.
// 24 bytes provides 192 bits of entropy, base64-encoded to 32 characters.
const clientNonceLength = 24

// EnableChannelBinding configures the client to negotiate SCRAM-SHA-256-PLUS
// using the tls-server-end-point binding type, with the supplied hash as the
// cbind-data. Call before ClientFirstMessage. Pass nil to clear.
func (c *SCRAMClient) EnableChannelBinding(tlsServerEndPointHash []byte) {
	_ = "STUB: not implemented"
	return
}

// Mechanism returns the SASL mechanism name the client will use, based on
// whether channel binding has been enabled.
func (c *SCRAMClient) Mechanism() string { _ = "STUB: not implemented"; return "" }

// ClientFirstMessage generates the client-first-message to send to the server.
// This starts the SCRAM authentication handshake.
// Returns the full message including the GS2 header.
func (c *SCRAMClient) ClientFirstMessage() (string, error) {
	_ = "STUB: not implemented"
	// Generate random client nonce.
	return "", nil
}

// Build client-first-message-bare: n=<username>,r=<nonce>

// GS2 header: "p=tls-server-end-point,," when channel binding is on,
// "n,," otherwise. The authzid slot stays empty in both cases.

// ProcessServerFirst processes the server-first-message and generates the client-final-message.
// The serverFirst parameter is the server's response to the client-first-message.
// Returns the client-final-message to send to the server.
func (c *SCRAMClient) ProcessServerFirst(serverFirst string) (string, error) {
	_ = "STUB: not implemented"
	// Parse server-first-message.
	return "", nil
}

// Verify the combined nonce starts with our client nonce.

// Build cbind data: gs2-header || cbind-data (PLUS) or just gs2-header.

// Build AuthMessage.

// Compute ClientProof.

// Passthrough mode: use pre-extracted ClientKey.

// Password mode: derive keys from password.

// Store derived keys for server signature verification.

// Build client-final-message.

// VerifyServerFinal verifies the server-final-message for mutual authentication.
// The serverFinal parameter is the server's response to the client-final-message.
// Returns nil if the server signature is valid, or an error if verification fails.
func (c *SCRAMClient) VerifyServerFinal(serverFinal string) error {
	_ = "STUB: not implemented"
	// Parse server-final-message: v=<server_signature_b64>
	return nil
}

// Compute expected server signature.

// Use constant-time comparison.
