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
//
// Modifications Copyright 2025 Supabase, Inc.

package servenv

// Run starts listening for RPC and HTTP requests,
// and blocks until it the process gets a signal.
func (sv *ServEnv) Run(bindAddress string, port int, grpcServer *GrpcServer) error {
	_ = "STUB: not implemented"
	return nil
}

// Start the HTTP server early so liveness/startup probes respond
// before potentially-blocking run hooks (e.g., waiting for topology
// or manager readiness). This prevents a deadlock on K8s 1.33+ where
// native sidecar startup probes must pass before main containers start.

// If port was 0, log the actual allocated port

// Update the ListeningURL with the actual port

// Wait for signal

// Shutdown telemetry last to ensure all spans from cleanup are captured.
// TODO(dweitzman): Propagate the cobra.Command() context into ServEnv instead of using context.Background()
//nolint:gocritic // shutdown requires fresh context; see TODO above
