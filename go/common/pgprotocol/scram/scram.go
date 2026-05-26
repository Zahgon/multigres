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

const (
	// ScramSHA256Mechanism is the SASL mechanism name for SCRAM-SHA-256.
	ScramSHA256Mechanism = "SCRAM-SHA-256"

	// ScramSHA256PlusMechanism is the SASL mechanism name for SCRAM-SHA-256
	// with channel binding (RFC 5802 §6 + PostgreSQL convention). Advertised
	// only over TLS, alongside ScramSHA256Mechanism.
	ScramSHA256PlusMechanism = "SCRAM-SHA-256-PLUS"

	// serverNonceLength is the number of random bytes to add to the client nonce.
	serverNonceLength = 18
)

// clientFirstMessage represents a parsed SCRAM client-first-message.
// Format: gs2-header client-first-message-bare
// Where gs2-header = gs2-cbind-flag "," [authzid] ","
// And client-first-message-bare = username "," nonce ["," extensions]
type clientFirstMessage struct {
	// gs2CbindFlag is the raw channel binding flag from the wire:
	//   "n"      — client doesn't support channel binding
	//   "y"      — client supports but believes server doesn't advertise it
	//   "p=<t>"  — client requires channel binding of type <t>
	gs2CbindFlag string

	// channelBindingType is the parsed binding type when gs2CbindFlag has
	// the "p=" form (e.g. "tls-server-end-point"). Empty for "n"/"y".
	channelBindingType string

	// gs2Header is the literal "gs2-cbind-flag,[authzid]," prefix as it
	// appeared on the wire. Required to verify the cbind data the client
	// sends back in client-final-message.
	gs2Header string

	// authzid is the optional authorization identity.
	authzid string

	// username is the authentication identity (saslprep normalized).
	username string

	// clientNonce is the client-generated random nonce.
	clientNonce string

	// clientFirstMessageBare is the message without the GS2 header.
	// This is needed for computing the AuthMessage.
	clientFirstMessageBare string
}

// clientFinalMessage represents a parsed SCRAM client-final-message.
// Format: channel-binding "," nonce "," proof
type clientFinalMessage struct {
	// channelBinding is the base64-encoded channel binding data.
	// For no channel binding, this is "biws" (base64 of "n,,").
	channelBinding string

	// nonce is the combined client+server nonce.
	nonce string

	// proof is the decoded client proof.
	proof []byte

	// clientFinalMessageWithoutProof is the message without the proof.
	// This is needed for computing the AuthMessage.
	clientFinalMessageWithoutProof string
}

// parseClientFirstMessage parses a SCRAM client-first-message.
// The message format is: gs2-cbind-flag "," [authzid] "," saslname "=" username "," "r=" nonce
// Example: "n,,n=user,r=fyko+d2lbbFgONRv9qkxdawL"
func parseClientFirstMessage(msg string) (*clientFirstMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Split by comma to get parts.
// Format: gs2-cbind-flag, [authzid], client-first-message-bare-parts...

// Parse GS2 header. Note the channel binding decision (use/reject/downgrade)
// is the authenticator's job — here we only parse the wire form.

// Client doesn't support channel binding.

// Client supports channel binding but thinks server doesn't.

// Parse optional authzid.

// Preserve the literal gs2-header prefix for later cbind verification.

// The rest is the client-first-message-bare.

// Parse the client-first-message-bare for username and nonce.

// Ignore other attributes (extensions).

// parseClientFinalMessage parses a SCRAM client-final-message.
// The message format is: "c=" channel-binding "," "r=" nonce "," "p=" proof
// Example: "c=biws,r=fyko+d2lbbFgONRv9qkxdawL3rfcNHYJY1ZVvWVs7j,p=v0X8v3Bz2T0CJGbJQyF0X+HI4Ts="
func parseClientFinalMessage(msg string) (*clientFinalMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Parse attributes. The RFC permits each attribute at most once; an
// attacker repeating "c=" with a different value could otherwise slip
// a tampered binding past a last-write-wins parse. Reject duplicates.

// Decode the proof.

// Build client-final-message-without-proof for AuthMessage computation.
// Find where ",p=" starts and take everything before it.

// generateServerFirstMessage generates a SCRAM server-first-message.
// Returns the message string, the combined nonce, and any error.
// The message format is: "r=" nonce "," "s=" salt "," "i=" iteration-count
func generateServerFirstMessage(clientNonce string, salt []byte, iterations int) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// Generate server nonce (random bytes, base64 encoded to be printable).

// Combined nonce = client nonce + server nonce

// Encode salt as base64.

// Build the message.

// generateServerFinalMessage generates a SCRAM server-final-message.
// The message format is: "v=" base64(server-signature)
func generateServerFinalMessage(serverSignature []byte) string {
	_ = "STUB: not implemented"
	return ""
}

// decodeSaslName decodes a SASL-encoded username.
// In SASL names, '=' is encoded as '=3D' and ',' is encoded as '=2C'.
func decodeSaslName(s string) string { _ = "STUB: not implemented"; return "" }

// encodeSaslName encodes a username for SASL.
// '=' must be encoded as '=3D' and ',' must be encoded as '=2C'.
func encodeSaslName(s string) string { _ = "STUB: not implemented"; return "" }

// parseServerFirstMessage parses a SCRAM server-first-message.
// This is useful for client-side implementations and testing.
// The message format is: "r=" nonce "," "s=" salt "," "i=" iteration-count
func parseServerFirstMessage(msg string) (nonce string, salt []byte, iterations int, err error) {
	_ = "STUB: not implemented"
	return "", nil, 0, nil
}
