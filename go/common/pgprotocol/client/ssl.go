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

package client

import (
	"crypto/tls"
	"crypto/x509"
)

// SSLMode mirrors libpq's sslmode connection parameter.
// Reference: https://www.postgresql.org/docs/17/libpq-ssl.html
type SSLMode string

const (
	// SSLModeDisable never negotiates SSL.
	SSLModeDisable SSLMode = "disable"

	// SSLModeAllow first tries plaintext; falls back to SSL only if the server
	// rejects the plaintext attempt. libpq's least-preferred fallback path.
	SSLModeAllow SSLMode = "allow"

	// SSLModePrefer tries SSL first; on server refusal ('N'), continues plaintext.
	// libpq default.
	SSLModePrefer SSLMode = "prefer"

	// SSLModeRequire negotiates SSL but performs no certificate verification.
	// Encryption only.
	SSLModeRequire SSLMode = "require"

	// SSLModeVerifyCA verifies the server certificate chain against the root CA
	// but does not match the hostname.
	SSLModeVerifyCA SSLMode = "verify-ca"

	// SSLModeVerifyFull verifies the server certificate chain and matches the
	// hostname against SAN (with CN fallback) — full libpq verification.
	SSLModeVerifyFull SSLMode = "verify-full"
)

// ParseSSLMode parses the string form of an sslmode value.
// Empty input returns SSLModePrefer (libpq default).
//
// SSLModeAllow is rejected: libpq's "allow" semantics require a
// plaintext-then-TLS retry loop on a server-side rejection, which is not
// implemented here. Accepting the string would silently behave like "disable",
// breaking the libpq parity an operator would expect.
func ParseSSLMode(s string) (SSLMode, error) { _ = "STUB: not implemented"; return *new(SSLMode), nil }

// AttemptsTLS reports whether the mode wants the SSLRequest negotiation step.
func (m SSLMode) AttemptsTLS() bool { _ = "STUB: not implemented"; return false }

// RequiresTLS reports whether the mode must error if the server declines SSL.
// prefer is the only tolerant mode that reaches the negotiation path.
func (m SSLMode) RequiresTLS() bool { _ = "STUB: not implemented"; return false }

// BuildTLSConfig constructs a *tls.Config matching libpq's sslmode semantics.
//
//   - disable → returns (nil, nil); no TLS attempt.
//   - prefer/require → encryption only, no certificate verification.
//   - verify-ca → chain validated against rootCertPath; hostname not checked.
//   - verify-full → chain validated and ServerName=host (SAN match, with CN
//     fallback for libpq parity since Go's verifier removed CN matching in 1.17).
//
// allow is rejected by ParseSSLMode and never reaches this function. host must
// be non-empty for verify-full; if empty, the function returns an error rather
// than silently building a config whose hostname check is guaranteed to fail.
//
// rootCertPath is required for verify-ca / verify-full.
func BuildTLSConfig(mode SSLMode, rootCertPath, host string) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// libpq parity: require/prefer perform no cert verification — encryption only.

//nolint:gosec // libpq parity: require/prefer perform no cert verification

// verify-ca skips hostname match; verify-full adds libpq's CN fallback
// that Go's default SAN-only verifier won't perform. Both rely on
// VerifyConnection for the actual chain check, so InsecureSkipVerify
// is set to bypass Go's stricter verifier.

func loadCertPool(path string) (*x509.CertPool, error) { _ = "STUB: not implemented"; return nil, nil }

// makeVerifyChain returns a VerifyConnection function that validates the peer
// certificate chain against the supplied root pool but performs no hostname
// match — verify-ca semantics.
func makeVerifyChain(pool *x509.CertPool) func(tls.ConnectionState) error {
	_ = "STUB: not implemented"
	return nil
}

// makeVerifyFull validates the chain and matches the hostname against SAN
// entries first, then falls back to the certificate's Common Name. The CN
// fallback is required for libpq parity; Go's stdlib verifier removed it in 1.17.
func makeVerifyFull(pool *x509.CertPool, host string) func(tls.ConnectionState) error {
	_ = "STUB: not implemented"
	return nil
}

// Try Go's SAN-based hostname match first.

// libpq accepts a match against the certificate Common Name when no SANs
// match. Only fall back to CN if the cert has no SANs (matches libpq's
// behavior — RFC 6125 deprecates CN when SANs are present).
