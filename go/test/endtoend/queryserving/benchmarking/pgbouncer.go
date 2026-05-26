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

package benchmarking

import (
	"context"
	"testing"
	"time"

	"github.com/multigres/multigres/go/tools/executil"
)

// PgBouncerInstance manages a pgbouncer process for benchmark comparison.
type PgBouncerInstance struct {
	process   *executil.Cmd
	cancel    context.CancelFunc
	configDir string
	port      int
}

// NewPgBouncerInstance starts a pgbouncer instance pointing at the given PostgreSQL backend.
// Returns nil, nil if pgbouncer is not installed (caller should skip pgbouncer benchmarks).
func NewPgBouncerInstance(t *testing.T, backendHost string, backendPort int, user, password string) (*PgBouncerInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:nilnil // nil,nil signals "not available, skip gracefully"

// Write pgbouncer.ini
// Use scram-sha-256 auth to match PostgreSQL's default auth method.

// Write userlist.txt with plaintext password.
// pgbouncer handles the SCRAM-SHA-256 exchange itself when given plaintext.

// Start pgbouncer

// Wait for pgbouncer to accept connections

// Port returns the listen port of this pgbouncer instance.
func (p *PgBouncerInstance) Port() int {
	_ = "STUB: not implemented"

	// Stop terminates the pgbouncer process.
	return 0
}

func (p *PgBouncerInstance) Stop(t *testing.T) { _ = "STUB: not implemented"; return }

// pgbouncerAvailable returns true if the pgbouncer binary is on PATH.
func pgbouncerAvailable() bool { _ = "STUB: not implemented"; return false }

// waitForPort polls a TCP port until it accepts connections or the timeout expires.
func waitForPort(t *testing.T, host string, port int, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}
