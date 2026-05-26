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
	"crypto/tls"
)

// captureTLSServerCert records the leaf certificate that crypto/tls
// selected for this connection's handshake. Invoked by the GetCertificate
// / GetConfigForClient wrappers installed by wrapTLSConfigForCertCapture.
//
// Prefers the pre-parsed Leaf to avoid re-parsing on hot paths; falls back
// to parsing Certificate[0] for callers that don't populate Leaf. A parse
// failure is logged at warn and leaves tlsServerCert nil — auth then falls
// back to SCRAM-SHA-256 without channel binding rather than failing the
// handshake.
//
// Called synchronously from the handshake goroutine before tls.Handshake
// returns; the subsequent reader (handleSSLRequest / SCRAM advertisement)
// runs on the same goroutine, so no synchronization is needed.
func (c *Conn) captureTLSServerCert(tlsCert *tls.Certificate) { _ = "STUB: not implemented"; return }

// wrapTLSConfigForCertCapture returns a clone of base whose certificate
// selection paths (GetCertificate and GetConfigForClient) are instrumented
// to invoke capture with the *tls.Certificate that crypto/tls actually
// presented to the peer. The clone may safely be passed to tls.Server.
//
// Why: crypto/tls does not expose the server-presented leaf via
// (*tls.Conn).ConnectionState(), so a deployment using dynamic cert
// selection (typical for SNI multi-tenant TLS) cannot recover the cert
// post-handshake. SCRAM-SHA-256-PLUS channel binding needs that cert
// (RFC 5929 tls-server-end-point). Capturing during selection is the
// only stable way to retrieve it.
//
// TODO(go-stdlib): the wrapper-of-getters pattern below — and this
// whole file — exists only because crypto/tls has no asymmetric
// counterpart to ConnectionState.PeerCertificates for the local
// (server) side. The open proposal is golang/go#24673
// ("crypto/tls: provide a way to access local certificate used to
// set up a connection"). Delete this file when that lands and read
// the leaf directly from (*tls.Conn).ConnectionState().
//
// The static Certificates[0] path is intentionally NOT routed through
// capture here — the caller continues to handle it after the handshake
// to keep behavior unchanged for the common single-cert deployment.
//
// capture is invoked synchronously on the handshake goroutine; the caller
// must serialize subsequent reads (in this package, capture writes to the
// owning *Conn which is only read after Handshake returns on the same
// goroutine).
func wrapTLSConfigForCertCapture(base *tls.Config, capture func(*tls.Certificate)) *tls.Config {
	_ = "STUB: not implemented"
	return nil
}

// Pure static deployment — caller's post-handshake fallback
// already recovers Certificates[0]. Avoid an unnecessary clone.

// crypto/tls falls back to the outer config — already wrapped above.

// wrapInnerCfgForCertCapture handles the per-handshake *tls.Config returned
// by a user-supplied GetConfigForClient. crypto/tls uses this returned
// config for cert selection, so the outer wrapping in
// wrapTLSConfigForCertCapture does not see the eventual cert — we must
// wrap the inner config as well.
func wrapInnerCfgForCertCapture(inner *tls.Config, capture func(*tls.Certificate)) *tls.Config {
	_ = "STUB: not implemented"
	// Skip cloning when neither cert-selection path is populated — the
	// clone would be returned unmodified and crypto/tls would still fall
	// back to the outer (already-wrapped) config. Parallels the static
	// short-circuit in wrapTLSConfigForCertCapture.
	return nil
}

// Inner config has no dynamic getter — crypto/tls will select from
// inner.Certificates. Synthesize a GetCertificate that captures the
// chosen cert. We replicate crypto/tls's selection rules at a minimum:
// the first cert that SupportsCertificate matches, otherwise
// Certificates[0]. NameToCertificate is intentionally skipped (see
// selectCertificate doc-comment).

// selectCertificate mirrors the subset of crypto/tls's internal server cert
// selection we need for inner configs whose user didn't supply a getter.
// Iterates Certificates picking the first SupportsCertificate match; falls
// back to Certificates[0] when no match is found, matching crypto/tls
// behavior. NameToCertificate is intentionally skipped (deprecated in
// crypto/tls; SupportsCertificate covers the same SNI matching).
//
// This is the most fragile piece of the file: it re-implements stdlib
// selection rules, so drift in crypto/tls (new TLS version, new
// SupportsCertificate constraint) would silently break SCRAM-PLUS
// for the GetConfigForClient-with-static-inner deployment shape.
// See the TODO on wrapTLSConfigForCertCapture — when golang/go#24673
// lands, this helper goes away with the rest of the file.
func selectCertificate(cfg *tls.Config, chi *tls.ClientHelloInfo) *tls.Certificate {
	_ = "STUB: not implemented"
	return nil
}

// tlsConfigYieldsServerCert reports whether a TLS config will provide a
// server certificate at handshake time through any supported path:
// static Certificates, GetCertificate, or GetConfigForClient. Listener
// init logs a warning when this returns false on a non-nil config so the
// operator notices the silent SCRAM-SHA-256-PLUS-off state.
func tlsConfigYieldsServerCert(cfg *tls.Config) bool { _ = "STUB: not implemented"; return false }
