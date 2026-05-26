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

package shardsetup

import (
	"testing"
)

// MultigatewayTLSCertPaths holds the paths to the generated multigateway TLS certificates.
type MultigatewayTLSCertPaths struct {
	CACertFile     string // CA certificate file (for client verify-ca / verify-full)
	ServerCertFile string // Server certificate file
	ServerKeyFile  string // Server private key file
}

// generateMultigatewayTLSCerts creates TLS certificates for the multigateway PostgreSQL listener.
// Generates a CA and server cert with CN=localhost, SANs=[localhost] for test connections.
func (s *ShardSetup) generateMultigatewayTLSCerts(t *testing.T) { _ = "STUB: not implemented"; return }
