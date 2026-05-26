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

package pgbuilder

import (
	"context"
	"testing"

	"github.com/multigres/multigres/go/tools/executil"
)

// Standalone is a single PostgreSQL server process initialized from a
// Builder's installed binaries and running on a dynamically chosen TCP port.
//
// It is intentionally independent of shardsetup, pgctld, and multigateway:
// callers (e.g. the sqllogictest differential harness) need a clean baseline
// PostgreSQL to diff Multigres behaviour against.
type Standalone struct {
	// DataDir is the data directory used by initdb/postgres.
	DataDir string
	// LogPath is the path to the captured server log.
	LogPath string
	// Port is the TCP port the server listens on.
	Port int
	// User is the superuser created by initdb (always "postgres").
	User string
	// Password is the superuser password.
	Password string
	// Database is the default database available on startup.
	Database string

	binDir string
	cmd    *executil.Cmd
}

// StartStandalone runs initdb into a fresh data directory under builder.OutputDir
// and launches a postgres server on a free port. The returned Standalone
// exposes connection parameters and a Stop method.
//
// Callers should defer Stop. The server logs are written to LogPath and left
// on disk after Stop to aid debugging.
func StartStandalone(t *testing.T, ctx context.Context, builder *Builder, password string) (*Standalone, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Loopback-only, md5 auth so psql/pgx connecting with PGPASSWORD works and
// nothing outside the host can reach the server.

func (s *Standalone) runInitdb(ctx context.Context, pwFile string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Standalone) launch() error { _ = "STUB: not implemented"; return nil }

// waitReady polls pg_isready until postgres accepts connections or the deadline
// expires. pg_isready ships with every PG install so this avoids pulling in a
// Go database driver just for the healthcheck.
func (s *Standalone) waitReady(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Stop terminates the running postgres process. Safe to call multiple times.
func (s *Standalone) Stop() error { _ = "STUB: not implemented"; return nil }

// pickFreePort asks the kernel for an unused TCP port by binding :0, then
// closes the listener and returns the port.
func pickFreePort() (int, error) { _ = "STUB: not implemented"; return 0, nil }
