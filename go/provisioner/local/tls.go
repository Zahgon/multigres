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

package local

import (
	"crypto/rsa"
	"crypto/x509"
)

// GenerateCA generates a self-signed CA certificate and private key.
// The certificate is valid for 10 years with RSA 4096-bit key.
func GenerateCA(certPath, keyPath string) error {
	_ = "STUB: not implemented"
	// Generate RSA private key (4096 bits to match k8s setup)
	return nil
}

// Create CA certificate template

// Valid for 10 years

// Self-sign the CA certificate

// Ensure directories exist

// Write CA certificate

// Write CA private key

// GenerateCert generates a certificate signed by the CA.
// The cn is used as the certificate's Common Name for identification.
// SANs are added as DNS Subject Alternative Names. IP SANs for 127.0.0.1 and ::1 are always included.
func GenerateCert(caCertPath, caKeyPath, certPath, keyPath, cn string, sans []string) error {
	_ = "STUB: not implemented"
	// Load CA certificate and key
	return nil
}

// Generate server private key (2048 bits to match k8s setup)

// Create server certificate template

// Valid for 1 year

// Sign the certificate with CA

// Ensure directories exist

// Write server certificate

// Write server private key

// LoadCA loads a CA certificate and private key from disk.
func LoadCA(certPath, keyPath string) (*x509.Certificate, *rsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	// Load CA certificate
	return nil, nil, nil
}

// Load CA private key
