// Copyright 2025 Supabase, Inc.
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

package testutil

import (
	"testing"
)

// TempDir creates a temporary directory for testing and returns a cleanup function
func TempDir(t *testing.T, prefix string) (string, func()) {
	_ = "STUB: not implemented"
	return "", nil
}

// Clean up any leftover PostgreSQL mock processes

// CreateDataDir creates a PostgreSQL-like data directory structure for testing.
// It sets the PGDATA environment variable to baseDir/pg_data for the duration of the test.
func CreateDataDir(t *testing.T, baseDir string, initialized bool) string {
	_ = "STUB: not implemented"
	return ""
}

// Create PG_VERSION file to indicate initialized data directory

// Generate a proper postgresql.conf file using the postgresconfig_gen functionality

// Create other typical PostgreSQL files manually

// Create base directory

// CreatePIDFile creates a postmaster.pid file for testing with a real running process
func CreatePIDFile(t *testing.T, dataDir string, pid int) {
	_ = "STUB: not implemented"

	// Start a background sleep process to get a real PID that will pass the isProcessRunning check
	return
}

// RemovePIDFile removes the postmaster.pid file for testing
func RemovePIDFile(t *testing.T, dataDir string) { _ = "STUB: not implemented"; return }

// CreateDeadPIDFile creates a postmaster.pid file with a PID that does not exist,
// simulating a crashed PostgreSQL process
func CreateDeadPIDFile(t *testing.T, dataDir string, deadPID int) {
	_ = "STUB: not implemented"
	return
}

// cleanupMockProcesses kills any leftover sleep processes created by mock PostgreSQL binaries
func cleanupMockProcesses(t *testing.T, tempDir string) {
	_ = "STUB: not implemented"

	// Look for any postmaster.pid files in the temp directory and kill associated processes
	return
}

//nolint:nilerr // Continue walking even if there's an error with one file

// Read the PID from the file and kill the process

//nolint:nilerr // Continue if we can't read the file

// Try to kill the process (ignore errors since process might already be dead)
