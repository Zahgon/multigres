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

package grpccommon

import (
	"crypto/tls"
)

// BuildServerTLSConfig creates a TLS configuration for a gRPC server.
// Returns nil if neither certFile nor keyFile is provided (plaintext mode).
//
// When caFile is provided, mutual TLS is enabled: the server requires and
// verifies client certificates against the given CA.
//
// When serverCAFile is provided, the intermediate CA certificate is appended
// to the server's certificate chain so clients receive the full chain.
func BuildServerTLSConfig(certFile, keyFile, caFile, serverCAFile string) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reject ca/serverCA-only configurations: a user that set them
// likely intended TLS/mTLS and would otherwise silently get plaintext.

// Append intermediate CA certificates to the server's certificate chain
// so clients receive the full chain during the TLS handshake.

// Skip non-certificate blocks (e.g., a stray PRIVATE KEY block in a
// mixed PEM file). Appending them would corrupt the chain and only
// surface as a confusing TLS handshake error later.

// Enable mutual TLS if a client CA is provided.

// BuildClientTLSConfig creates a TLS configuration for a gRPC client.
// Returns nil if no TLS parameters are provided (insecure mode).
//
// When caFile is provided, the client verifies the server's certificate
// against the given CA. When certFile and keyFile are provided, the client
// presents a certificate for mutual TLS.
//
// serverName alone is rejected: for internal gRPC that would silently enable
// TLS against the system trust store, which is almost always a
// misconfiguration rather than an intentional "use public CAs" choice.
func BuildClientTLSConfig(certFile, keyFile, caFile, serverName string) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Load CA for server certificate verification.

// Load client certificate for mutual TLS.
