//go:build !windows

/*
Copyright 2023 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Modifications Copyright 2025 Supabase, Inc.
*/

package servenv

// Init is the first phase of the server startup.
// The id parameter provides service identification for telemetry resource attributes.
func (sv *ServEnv) Init(id ServiceIdentity) error { _ = "STUB: not implemented"; return nil }

// Build OTel resource attributes from service identity

// Compute OTel-compliant service.instance.id (cell-qualified for multi-cell uniqueness).
// Per OTel semantic conventions, service.instance.id must be globally unique for each
// instance of the same service.name. For multi-cell deployments, we qualify the instance
// ID with the cell name to achieve global uniqueness.

// Multi-cell: qualify instance ID with cell (e.g., "zone1-0")

// Single-cell: use instance ID directly

// Add multigres-specific resource attributes (multipooler only)

// Tag every metric/span/log with the binary's VCS identity so callers
// can distinguish builds in dashboards and trace mixed-version
// deployments. service.version and vcs.ref.head.revision are the
// canonical OTel semconv attributes for this.

// Initialize OpenTelemetry with service identity attributes

// Continue without telemetry rather than crashing

// Re-wrap logger now that LoggerProvider is initialized

// Ignore SIGPIPE if specified
// The Go runtime catches SIGPIPE for us on all fds except stdout/stderr
// See https://golang.org/pkg/os/signal/#hdr-SIGPIPE

// Add version tag to every info log

// Once you run as root, you pretty much destroy the chances of a
// non-privileged user starting the program correctly.

// We used to set this limit directly, but you pretty much have to
// use a root account to allow increasing a limit reliably. Dropping
// privileges is also tricky. The best strategy is to make a shell
// script set up the limits as root and switch users before starting
// the server.

// Limit the stack size. We don't need huge stacks and smaller limits mean
// any infinite recursion fires earlier and on low memory systems avoids
// out of memory issues in favor of a stack overflow error.

// Get hostname upfront so we can fail early if it fails.

func (sv *ServEnv) populateHostname() error {
	_ = "STUB: not implemented"
	// If hostname was explicitly set via --hostname flag, use that
	return nil
}

// Otherwise, auto-detect hostname

// startOrphanDetection starts a goroutine that monitors for orphan conditions.
// It checks:
// 1. If MULTIGRES_TESTDATA_DIR is set and the directory is deleted
// 2. If MULTIGRES_TEST_PARENT_PID is set and that process no longer exists
// If either condition is true, initiates graceful shutdown.
// This is used in integration tests to ensure child processes don't become
// orphans if the test runner is killed.
func (sv *ServEnv) startOrphanDetection() {
	_ = "STUB: not implemented"
	// Only run if orphan detection environment variables are set
	return
}

// Channel to signal when close hooks have completed

// Check if testdata directory was deleted

// Check if test parent process died

// Trigger graceful shutdown

// Wait for close hooks to complete or 10 second timeout

//nolint:forbidigo // Last resort: this is a test and graceful shutdown already timed out

// Normal shutdown - stop orphan detection
