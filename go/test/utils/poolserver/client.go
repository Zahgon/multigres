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

package poolserver

import (
	"bufio"
	"net"
	"sync"
)

// Client is a connection to a running pool server.
// A single Client may be shared across goroutines; all operations are
// serialised by an internal mutex.
type Client struct {
	mu      sync.Mutex
	conn    net.Conn
	scanner *bufio.Scanner
	writer  *bufio.Writer
}

// Connect opens a connection to the pool server at socketPath.
func Connect(socketPath string) (*Client, error) { _ = "STUB: not implemented"; return nil, nil }

// maxAllocRetries is the maximum number of collision retries for AllocPort.
const maxAllocRetries = 10

// Alloc requests a new port from the server.
// The port is immediately bindable — the server does not hold a listener on it.
// The server records the port number in its registry to prevent handing the
// same number to another caller.
//
// In the very rare case where the OS returns a port the server is already
// tracking (collision), Alloc retries automatically up to maxAllocRetries
// times before returning an error.
func (c *Client) AllocPort() (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Transient collision; retry.

// Return tells the server the port is no longer in use by any process.
// The server removes the port from its tracking tables.
// Call this from t.Cleanup when a test that used the port completes.
func (c *Client) ReturnPort(port int) error { _ = "STUB: not implemented"; return nil }

// Ping checks that the server is alive.
func (c *Client) Ping() error { _ = "STUB: not implemented"; return nil }

// Close closes the connection to the server.
// The server will automatically clean up any ports still held for this
// connection.
func (c *Client) Close() error { _ = "STUB: not implemented"; return nil }

func (c *Client) send(msg string) (string, error) { _ = "STUB: not implemented"; return "", nil }
