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

// Package executil provides safe subprocess execution with graceful termination,
// explicit environment variable handling, and OpenTelemetry trace propagation.
//
// # Why Use executil Instead of exec.Command?
//
// Go's standard library exec package has several issues for production infrastructure:
//
//  1. Ungraceful termination: exec.CommandContext immediately kills subprocesses
//     on context cancellation, preventing log flushing and telemetry export
//  2. No trace propagation: Manual TRACEPARENT handling is error-prone
//  3. Environment footguns: exec.Cmd.Env has confusing nil semantics
//     (nil = inherit, non-nil = replace entirely)
//  4. Implicit termination: No control over grace periods
//
// By using executil consistently across the codebase, we ensure:
//
//   - Graceful termination (SIGTERM → SIGKILL) reduces telemetry data loss
//   - Automatic trace context propagation to all subprocesses
//   - Explicit environment variable handling (AddEnv/SetEnv)
//   - Explicit grace period control (automatic or explicit)
//
// # Graceful Termination
//
// Commands are terminated gracefully with SIGTERM first, escalating to SIGKILL
// if needed. This allows subprocesses to flush logs, send telemetry, and clean
// up properly.
//
// Termination can happen in two ways:
//   - Automatic: When the parent context passed to Command() is cancelled,
//     the process is terminated with the default grace period (10s).
//   - Explicit: Call Stop(ctx) where ctx's timeout controls the grace period,
//     or call Terminate(ctx)/Kill(ctx) for lower-level control.
//
// # Environment Variables
//
// Environment variables are handled explicitly via AddEnv() and SetEnv() methods,
// avoiding the subtle pitfalls of exec.Cmd.Env (where nil means "inherit" but
// non-nil means "replace entirely").
//
// # Trace Propagation
//
// Trace context is automatically propagated to subprocesses via the TRACEPARENT
// environment variable if the context contains a valid span. This ensures
// distributed tracing works correctly across process boundaries.
//
// # Example
//
//	ctx := context.Background()
//	cmd := executil.Command(ctx, "postgres", "-D", dataDir).
//	    AddEnv("PGPORT=5432").
//	    SetDir(workDir)
//
//	if err := cmd.Start(); err != nil {
//	    return err
//	}
//
//	// Process will be gracefully terminated when ctx is cancelled
//	// or you can explicitly control termination:
//	stopCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	exitErr, stopped := cmd.Stop(stopCtx)
package executil

import (
	"context"
	"os"
	"os/exec"
	"sync"
	"time"

	"go.opentelemetry.io/otel"
)

// DefaultGracePeriod is the time to wait after SIGTERM before escalating to SIGKILL.
// This is used when the parent context is cancelled.
const DefaultGracePeriod = 10 * time.Second

// DefaultKillTimeout is the time to wait for a process to exit after SIGKILL.
// SIGKILL should be nearly instant, but zombies or uninterruptible sleep can delay.
const DefaultKillTimeout = 5 * time.Second

const tracingServiceName = "multigres"

var tracer = otel.Tracer(tracingServiceName)

// Cmd wraps exec.Cmd with a builder pattern for safe configuration.
// Create with Command() or CommandWithGracePeriod().
type Cmd struct {
	*exec.Cmd
	parentCtx          context.Context
	defaultGracePeriod time.Duration
	extraEnv           []string
	clientSpan         bool
	processGroup       bool          // If true, signals target the process group, not just the direct child.
	waitDelay          time.Duration // Max time to wait for I/O after process exit in process group mode.

	// Termination coordination
	terminateOnce sync.Once
	terminated    chan struct{} // Closed when Terminate() is called
	waitDone      chan struct{} // Closed when Wait() completes
	waitErr       error         // Result of Wait() (valid after waitDone closed)
	waitOnce      sync.Once
}

// Command creates a new Cmd with graceful termination support.
//
// If the parent context is cancelled, the process receives SIGTERM and is given
// DefaultGracePeriod to exit before SIGKILL. For explicit termination with a
// different grace period, call Terminate(ctx) where ctx's timeout controls the wait.
//
// By default, the command inherits the parent process environment.
// Use AddEnv() to add variables, or SetEnv() to replace the entire environment.
func Command(ctx context.Context, name string, args ...string) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

// CommandWithGracePeriod creates a Cmd with a custom default grace period.
//
// The grace period controls how long to wait after SIGTERM before sending SIGKILL
// when the parent context is cancelled. For explicit Terminate() calls, the caller
// controls the grace period via the context timeout instead.
//
// Use a shorter grace period for commands that should terminate quickly
// (e.g., 100ms for simple queries).
func CommandWithGracePeriod(ctx context.Context, gracePeriod time.Duration, name string, args ...string) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

// AddEnv adds environment variables to the command. Variables are specified
// as "KEY=value" strings. Safe to call multiple times - variables accumulate.
//
// Variables are added on top of the inherited environment (or the explicit
// base if SetEnv was called). The actual environment is finalized when
// Start/Run/Output/CombinedOutput is called.
func (c *Cmd) AddEnv(keyvals ...string) *Cmd { _ = "STUB: not implemented"; return nil }

// SetEnv replaces the entire environment with the provided variables.
// The command will NOT inherit any environment from the parent process.
//
// Call AddEnv() after SetEnv() to add additional variables on top of
// this explicit base.
func (c *Cmd) SetEnv(env []string) *Cmd { _ = "STUB: not implemented"; return nil }

// SetDir sets the working directory for the command.
func (c *Cmd) SetDir(dir string) *Cmd { _ = "STUB: not implemented"; return nil }

// SetStdin sets the stdin for the command.
func (c *Cmd) SetStdin(r *os.File) *Cmd { _ = "STUB: not implemented"; return nil }

// SetStdout sets the stdout for the command.
func (c *Cmd) SetStdout(w *os.File) *Cmd { _ = "STUB: not implemented"; return nil }

// SetStderr sets the stderr for the command.
func (c *Cmd) SetStderr(w *os.File) *Cmd { _ = "STUB: not implemented"; return nil }

// WithClientSpan enables creating an OpenTelemetry client span around
// the command execution. The span is started when Start/Run is called
// and ended when the command completes.
func (c *Cmd) WithClientSpan() *Cmd { _ = "STUB: not implemented"; return nil }

// WithProcessGroup starts the command in its own process group and sends
// signals to the entire group instead of just the direct child. This ensures
// that grandchildren (e.g. make → pg_regress → psql) are terminated when the
// parent context is cancelled.
//
// When enabled:
//   - The child process is started with Setpgid: true
//   - Terminate() sends SIGTERM to the process group (negative PID)
//   - Kill() sends SIGKILL to the process group
//   - Run() uses Start()+Wait() internally to enable context-based termination
//
// Callers should also call SetWaitDelay() to prevent hangs from grandchildren
// holding pipes open after the group leader exits.
func (c *Cmd) WithProcessGroup() *Cmd { _ = "STUB: not implemented"; return nil }

// SetWaitDelay sets the maximum time Run()/Wait() will wait for I/O goroutines
// to finish after the process exits. This prevents hangs when child processes
// inherit stdout/stderr pipes and outlive the parent.
func (c *Cmd) SetWaitDelay(d time.Duration) *Cmd { _ = "STUB: not implemented"; return nil }

// finalizeEnv prepares cmd.Env before execution, including trace propagation.
func (c *Cmd) finalizeEnv() {
	_ = "STUB: not implemented"
	// Add TRACEPARENT if context has a valid span
	return
}

// watchContext starts a background goroutine that monitors the parent context
// for cancellation and gracefully terminates the process if it is cancelled.
//
// Must be called after c.Cmd.Start() succeeds. If called before, a
// cancelled context fires Terminate() while c.Process is still nil, so
// SIGTERM is a no-op and c.terminated is closed — the process then starts
// and runs indefinitely with no watcher.
func (c *Cmd) watchContext() {
	_ = "STUB: not implemented"
	// Watch for parent context cancellation
	return
}

// Parent context cancelled - terminate with default grace period.
// Fresh context needed; parent context is cancelled.

// Fresh context needed for kill timeout

// Already terminated explicitly, nothing to do

// Process exited naturally, nothing to do

// Start starts the command without waiting for it to complete.
//
// If the parent context is cancelled, the process will be terminated with
// SIGTERM followed by SIGKILL after the default grace period.
//
// Note: WithClientSpan() has no effect on Start() since the span cannot be
// ended until Wait() is called. Use Run() for client span support.
func (c *Cmd) Start() error { _ = "STUB: not implemented"; return nil }

// Wait waits for the command to exit and returns its exit status.
// Wait must be called after Start() to release resources.
// Safe to call multiple times or concurrently - returns cached result.
func (c *Cmd) Wait() error {
	_ = "STUB: not implemented"
	// Ensure we only call the underlying Wait() once.
	// sync.Once.Do blocks all callers until the first call completes.
	return nil
}

// Channel close provides happens-before guarantee for waitErr read

// Block until Wait() completes (no-op if already done), then return cached result

// Terminate sends SIGTERM to the process and waits for it to exit gracefully.
//
// Returns (exitErr, true) if the process exited before ctx expired.
// Returns (nil, false) if ctx expired before the process exited.
//
// If false is returned, the process is still running. Call Kill() to force termination.
//
// Safe to call multiple times or concurrently - only the first call sends SIGTERM,
// subsequent calls just wait for the process to exit.
//
// Panics if called on a Cmd not created via Command() or CommandWithGracePeriod().
func (c *Cmd) Terminate(ctx context.Context) (error, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Send SIGTERM only once

// Ensure Wait() is running in background.
// Result is stored in c.waitErr and signaled via c.waitDone.

// Wait for process exit or context timeout

// Kill sends SIGKILL to the process and waits for it to exit.
//
// The context controls how long to wait for the process to actually exit.
// SIGKILL should be nearly instant, but zombies or uninterruptible sleep can delay.
//
// Returns (exitErr, true) if the process exited before ctx expired.
// Returns (ctx.Err(), false) if the wait timed out.
//
// Safe to call after Terminate() times out - reuses the same Wait() call.
func (c *Cmd) Kill(ctx context.Context) (error, bool) {
	_ = "STUB: not implemented"
	// Send SIGKILL
	return nil, false
}

// Ensure Wait() is running in background.
// Result is stored in c.waitErr and signaled via c.waitDone.

// Wait for process exit or context timeout

// Stop gracefully stops the process: SIGTERM first, then SIGKILL if needed.
//
// The context controls how long to wait for graceful shutdown (SIGTERM phase).
// If the process doesn't exit before ctx expires, SIGKILL is sent with a short
// fixed timeout (100ms). SIGKILL should be nearly instant - if it times out,
// the process is likely a zombie or in uninterruptible sleep (rare system issue).
//
// Returns (exitErr, true) if the process stopped.
// Returns (nil, false) if SIGKILL timed out (very rare - indicates system issue).
//
// This is the recommended way to stop a process - always try graceful termination first.
//
// Example:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	exitErr, stopped := cmd.Stop(ctx)
//	// Tries SIGTERM for up to 10s, then SIGKILL with 100ms timeout
func (c *Cmd) Stop(ctx context.Context) (error, bool) { _ = "STUB: not implemented"; return nil, false }

// Run starts the command and waits for it to complete.
// If WithClientSpan() was called, an OpenTelemetry span is created around
// the command execution.
func (c *Cmd) Run() error { _ = "STUB: not implemented"; return nil }

// Output runs the command and returns its stdout.
// If WithClientSpan() was called, an OpenTelemetry span is created around
// the command execution.
func (c *Cmd) Output() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// CombinedOutput runs the command and returns its combined stdout and stderr.
// If WithClientSpan() was called, an OpenTelemetry span is created around
// the command execution.
func (c *Cmd) CombinedOutput() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
