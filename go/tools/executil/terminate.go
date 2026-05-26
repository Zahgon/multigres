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

package executil

import (
	"context"
	"os"
)

// TerminateProcess sends SIGTERM to a process and waits for graceful exit.
//
// Returns true if the process exited before ctx expired.
// Returns false if ctx expired - process may still be running, call KillProcess().
//
// Returns true immediately if process is nil or was already dead.
//
// NOTE: Prefer using Cmd.Stop() if you have an executil.Cmd for this process.
// Only use TerminateProcess() when dealing with a process started by a different
// process than the currently-running one (e.g., orphaned processes, processes
// from PID files, or processes discovered via OS APIs).
func TerminateProcess(ctx context.Context, process *os.Process) bool {
	_ = "STUB: not implemented"
	return false
}

// KillProcess sends SIGKILL to a process and waits for it to exit.
//
// Returns (nil, true) if the process exited before ctx expired.
// Returns (ctx.Err(), false) if the wait timed out (unexpected for SIGKILL).
//
// Returns (nil, true) immediately if process is nil or was already dead.
//
// NOTE: Prefer using Cmd.Stop() if you have an executil.Cmd for this process.
// Only use KillProcess() when dealing with a process started by a different
// process than the currently-running one (e.g., orphaned processes, processes
// from PID files, or processes discovered via OS APIs).
func KillProcess(ctx context.Context, process *os.Process) (error, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// TerminatePID sends SIGTERM to a process by PID and waits for graceful exit.
//
// Returns true if the process exited before ctx expired.
// Returns false if ctx expired - process may still be running, call KillPID().
//
// Returns true immediately if the process was already dead.
//
// NOTE: Prefer using Cmd.Stop() if you have an executil.Cmd for this process.
// Only use TerminatePID() when dealing with a process started by a different
// process than the currently-running one (e.g., orphaned processes, processes
// from PID files, or processes discovered via OS APIs).
func TerminatePID(ctx context.Context, pid int) bool { _ = "STUB: not implemented"; return false }

// Process doesn't exist

// Send SIGTERM

// SIGTERM failed for unexpected reason - process state unknown

// Wait for process to exit or context timeout

// KillPID sends SIGKILL to a process by PID and waits for it to exit.
//
// Returns (nil, true) if the process exited before ctx expired.
// Returns (ctx.Err(), false) if the wait timed out (unexpected for SIGKILL).
//
// Returns (nil, true) immediately if the process was already dead.
//
// NOTE: Prefer using Cmd.Stop() if you have an executil.Cmd for this process.
// Only use KillPID() when dealing with a process started by a different
// process than the currently-running one (e.g., orphaned processes, processes
// from PID files, or processes discovered via OS APIs).
func KillPID(ctx context.Context, pid int) (error, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

//nolint:nilerr // err means process doesn't exist, which is success for kill

// Send SIGKILL

// Wait for process to exit

// StopProcess gracefully stops a process: SIGTERM first, then SIGKILL if needed.
//
// The context controls how long to wait for graceful shutdown (SIGTERM phase).
// If the process doesn't exit before ctx expires, SIGKILL is sent with a short
// fixed timeout (100ms).
//
// Returns (nil, true) if the process stopped.
// Returns (nil, false) if SIGKILL timed out (very rare - indicates system issue).
//
// This is the recommended way to stop a process - always try graceful termination first.
func StopProcess(ctx context.Context, process *os.Process) (error, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// StopPID gracefully stops a process by PID: SIGTERM first, then SIGKILL if needed.
//
// The context controls how long to wait for graceful shutdown (SIGTERM phase).
// If the process doesn't exit before ctx expires, SIGKILL is sent with a short
// fixed timeout (100ms).
//
// Returns (nil, true) if the process stopped.
// Returns (nil, false) if SIGKILL timed out (very rare - indicates system issue).
//
// This is the recommended way to stop a process - always try graceful termination first.
func StopPID(ctx context.Context, pid int) (error, bool) {
	_ = "STUB: not implemented"
	// Try SIGTERM with caller's timeout
	return nil, false
}

// SIGTERM didn't work, escalate to SIGKILL with fixed short timeout
// Fresh context for kill - parent context already expired

// isProcessGone returns true if the error indicates the process doesn't exist.
func isProcessGone(err error) bool { _ = "STUB: not implemented"; return false }

// Fallback to string matching for edge cases

// waitForProcessExit polls until the process exits or context is done.
// Returns true if process exited, false if context was cancelled/timed out.
//
// Uses exponential backoff starting at 1ms, doubling each time up to 100ms.
// This responds quickly for fast exits while reducing CPU usage for slow exits.
func waitForProcessExit(ctx context.Context, process *os.Process) bool {
	_ = "STUB: not implemented"
	// Check immediately before starting the timer
	return false
}
