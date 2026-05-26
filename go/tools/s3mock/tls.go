// Copyright 2026 Supabase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// go/tools/s3mock/tls.go
package s3mock

import (
	"crypto/tls"
)

// generateTLSConfig creates an ephemeral self-signed TLS certificate for the mock server.
// The certificate is valid for 1 hour and supports localhost/127.0.0.1.
func generateTLSConfig() (*tls.Config, error) {
	_ = "STUB: not implemented"
	// Generate 2048-bit RSA private key
	return nil, nil
}

// Create certificate template

// Create self-signed certificate

// Build TLS certificate
