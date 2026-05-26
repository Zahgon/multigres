// Copyright 2025 Supabase, Inc.
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

package utils

import (
	"sync"
	"testing"
	"time"

	"github.com/multigres/multigres/go/test/utils/poolserver"
)

// portCache tracks ports currently allocated to tests to prevent duplicates
// within a single test binary. Values are always true.
var portCache sync.Map

// poolOnce initialises the singleton pool client exactly once per process.
var (
	poolOnce   sync.Once
	poolClient *poolserver.Client // nil if pool server is not in use
	warnNoPool sync.Once          // prints the "no pool server" warning at most once
)

const (
	poolConnectRetries = 10
	poolConnectDelay   = 500 * time.Millisecond
)

func getPoolClient() *poolserver.Client { _ = "STUB: not implemented"; return nil }

// GetFreePort returns a port number that was verified free by the OS and is
// not currently allocated to another test in this process.
//
// When MULTIGRES_PORT_POOL_ADDR is set and the pool server is reachable, the
// server records the returned port in a central registry. This prevents
// parallel test binaries from receiving the same port from the OS, eliminating
// a class of flaky-test failures caused by port collisions across concurrent
// "go test ./..." invocations. The port is immediately bindable.
//
// When the pool server is not in use, GetFreePort falls back to the original
// in-process behaviour: the OS assigns a free port via net.Listen(":0") and
// the port is tracked in a process-local cache to prevent duplicates.
//
// Ports are automatically released from the cache (and returned to the pool
// server, if in use) when the test completes via t.Cleanup.
func GetFreePort(t *testing.T) int { _ = "STUB: not implemented"; return 0 }

// getFreePortFromPool allocates a port through the running pool server.
func getFreePortFromPool(t *testing.T, client *poolserver.Client) int {
	_ = "STUB: not implemented"
	return 0
}

// Track in the local cache so GetFreePort never returns it twice within
// this process (belt-and-suspenders; the server already guarantees this).

// getFreePortLocal is the original in-process port allocator used when the
// pool server is not configured.
func getFreePortLocal(t *testing.T) int {
	_ = "STUB: not implemented"

	// Track listeners we need to keep open to prevent OS port reuse
	return 0
}

// Clean up all held listeners when done

// Try to claim this port atomically

// Successfully claimed this port

// Port already in cache - hold this listener open to prevent OS reuse
// and try again. Add to slice so we can close all at once when done.
