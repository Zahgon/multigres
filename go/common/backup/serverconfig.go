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

package backup

// ServerConfigOpts holds options for generating pgbackrest-server.conf.
type ServerConfigOpts struct {
	PoolerDir     string // Base directory for pooler data
	CertDir       string // Directory containing TLS certificates
	Port          int    // TLS server port
	Pg1Port       int    // Local PostgreSQL port
	Pg1SocketPath string // Local PostgreSQL socket directory
	Pg1Path       string // Local PostgreSQL data directory
	Pg1User       string // PostgreSQL superuser for pgbackrest connections
}

// WriteServerConfig generates a minimal pgbackrest-server.conf for the TLS server.
// The server config contains only TLS settings, log path, and pg1 stanza.
// Returns the path to the generated config file.
func WriteServerConfig(opts ServerConfigOpts) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// lockPath must be per-pooler. Without it, pgbackrest defaults to
// /tmp/pgbackrest/, which collides across multiple pgctld instances on the
// same host (e.g. parallel endtoend test binaries). The TLS server
// acquires this lock when a remote pgbackrest client (--pg2-host=tls)
// connects to perform a backup, so a default global path turns concurrent
// cross-pooler backups into a flake source.
//
// Distinct from the *client* lock-path on the same pooler (lock/, see
// clientconfig.go). Sharing one path would let same-pooler client and
// server processes serialize on pgbackrest's OS lock and race against
// our lease-stealing logic; keep them isolated until that race is fixed.

// Create directories
