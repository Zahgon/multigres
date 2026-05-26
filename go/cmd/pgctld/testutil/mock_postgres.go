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
	"os/exec"
	"testing"
)

// MockExecCommand mocks exec.Command for testing
type MockExecCommand struct {
	commands map[string]MockCommandResult
}

// MockCommandResult defines the expected result of a mocked command
type MockCommandResult struct {
	ExitCode int
	Stdout   string
	Stderr   string
	Error    error
}

// NewMockExecCommand creates a new mock command executor
func NewMockExecCommand() *MockExecCommand { _ = "STUB: not implemented"; return nil }

// AddCommand adds a mock command with expected result
func (m *MockExecCommand) AddCommand(cmdLine string, result MockCommandResult) {
	_ = "STUB: not implemented"
	return
}

// MockCommand simulates command execution for testing
func (m *MockExecCommand) MockCommand(name string, args ...string) *exec.Cmd {
	_ = "STUB: not implemented"
	return nil
}

// Create a fake command that will be handled by the test helper

// Store the command line for verification

// VerifyCommand checks if a command was called with expected arguments
func (m *MockExecCommand) VerifyCommand(t *testing.T, expectedCmd string) {
	_ = "STUB: not implemented"
	return
}

// MockBinary creates a mock binary for testing
func MockBinary(t *testing.T, binDir, name, content string) string {
	_ = "STUB: not implemented"
	return ""
}

// CreateMockPostgreSQLBinaries creates mock PostgreSQL binaries for testing
func CreateMockPostgreSQLBinaries(t *testing.T, binDir string) {
	_ = "STUB: not implemented"

	// Mock initdb
	return
}

// Mock pg_controldata

// Mock postgres

// Mock pg_ctl

// Mock pg_isready

// Mock psql
