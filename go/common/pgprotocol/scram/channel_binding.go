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

package scram

import (
	"crypto/x509"
	"hash"
)

// ChannelBindingTypeTLSServerEndPoint is the channel binding type defined in
// RFC 5929 §4: a hash of the TLS server's certificate. This is the type
// PostgreSQL advertises for SCRAM-SHA-256-PLUS.
const ChannelBindingTypeTLSServerEndPoint = "tls-server-end-point"

// ChannelBinding carries TLS channel binding context into the SCRAM
// authenticator. When non-nil and TLSServerEndPointHash is populated, the
// authenticator advertises SCRAM-SHA-256-PLUS in addition to SCRAM-SHA-256.
type ChannelBinding struct {
	// TLSServerEndPointHash is the hash of the server's TLS certificate, as
	// defined for the tls-server-end-point channel binding type (RFC 5929).
	// Compute it with ComputeTLSServerEndPointHash.
	TLSServerEndPointHash []byte
}

// ComputeTLSServerEndPointHash computes the tls-server-end-point channel
// binding data for a TLS server certificate per RFC 5929 §4: the hash of the
// DER-encoded certificate, using the certificate signature algorithm's hash
// function — with one PostgreSQL-compatible quirk: certificates signed with
// MD5 or SHA-1 are hashed with SHA-256 instead (mirroring PG's
// be_tls_get_certificate_hash). This is what libpq expects, so any client
// computing the binding the same way will interoperate.
func ComputeTLSServerEndPointHash(cert *x509.Certificate) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// tlsServerEndPointHasher returns the hash function to use for a given cert
// signature algorithm. MD5- and SHA-1-signed certs map to SHA-256 to match
// PostgreSQL (be_tls_get_certificate_hash). Ed25519 is intentionally NOT
// listed: PG's OpenSSL-based path cannot derive a digest for it and refuses
// to advertise PLUS, so accepting Ed25519 here would diverge from libpq's
// expected hash and break the cbind check.
func tlsServerEndPointHasher(sigAlgo x509.SignatureAlgorithm) (hash.Hash, error) {
	_ = "STUB: not implemented"
	return *new(hash.Hash), nil
}
