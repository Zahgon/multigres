// Copyright 2019 The Vitess Authors.
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

package etcdtopo

import (
	"context"
	"testing"

	"github.com/multigres/multigres/go/tools/executil"
)

// checkPortAvailable checks if a port is available for binding
func checkPortAvailable(port int) error { _ = "STUB: not implemented"; return nil }

// WaitForReady waits for etcd to be ready by querying its /readyz endpoint in a loop.
// metricsAddr must be the HTTP address of etcd's metrics listener (the address passed
// to --listen-metrics-urls), e.g. "http://localhost:2381". /readyz is not served on
// the client listener, so passing the client address here will always return 404.
// Callers should pass a context with an appropriate timeout.
func WaitForReady(ctx context.Context, metricsAddr string) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if error is "connection refused" and provide helpful diagnostic

// Check for orphan etcd processes

// EtcdOptions contains optional configuration for starting etcd.
type EtcdOptions struct {
	// ClientPort is the client port to listen on.
	// If 0, a port will be automatically assigned.
	ClientPort int

	// PeerPort is the peer port for etcd cluster communication.
	// If 0 and ClientPort is also 0, will be automatically assigned.
	// If 0 and ClientPort is specified, defaults to ClientPort+1 for backwards compatibility.
	PeerPort int

	// MetricsPort is the port for etcd's metrics/health listener (--listen-metrics-urls).
	// /readyz is only served on this listener, not on the client listener.
	// If 0, a port will be automatically assigned.
	MetricsPort int

	// DataDir is the directory for etcd data storage.
	// If empty, a temporary directory will be created and cleaned up after the test.
	DataDir string
}

// StartEtcd starts an etcd subprocess with automatically allocated ports.
// Returns the client address, the metrics address (to pass to WaitForReady),
// and the process handle.
func StartEtcd(t *testing.T) (clientAddr, metricsAddr string, cmd *executil.Cmd) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// StartEtcdWithOptions starts an etcd subprocess with custom options, and waits for it to be ready.
// Returns the client address, the metrics address (to pass to WaitForReady), and the process handle.
func StartEtcdWithOptions(t *testing.T, opts EtcdOptions) (clientAddr, metricsAddr string, cmd *executil.Cmd) {
	_ = "STUB: not implemented"
	// Check if etcd is available in PATH
	return "", "", nil
}

// Create a temporary directory if not specified.

// Get our ports to listen to - client and peer must be specified; metrics is auto-allocated if unset.

// Check if ports are available before starting etcd

// Wrap etcd with run_in_test.sh for orphan protection. Stops gracefully when
// the test context is cancelled so run_in_test.sh can terminate etcd cleanly.

// Set MULTIGRES_TESTDATA_DIR for directory-deletion triggered cleanup

// Wait for etcd to be ready via the metrics listener (/readyz requires --listen-metrics-urls)
