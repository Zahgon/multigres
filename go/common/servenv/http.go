// Copyright 2023 The Vitess Authors.
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

import (
	"net"
	"net/http"
)

// HTTPHandle registers the given handler for the internal servenv mux.
func (sv *ServEnv) HTTPHandle(pattern string, handler http.Handler) {
	_ = "STUB: not implemented"
	return
}

// HTTPHandleFunc registers the given handler func for the internal servenv mux.
func (sv *ServEnv) HTTPHandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	_ = "STUB: not implemented"
	return
}

// corsMiddleware adds CORS headers to allow cross-origin requests.
func corsMiddleware(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// Handle preflight OPTIONS requests

// HTTPServe starts the HTTP server for the internal servenv mux on the listener.
func (sv *ServEnv) HTTPServe(l net.Listener) error { _ = "STUB: not implemented"; return nil }

// Wrap the mux with CORS middleware and OpenTelemetry instrumentation
// If no OTEL exporters are configured, noop exporters are used with minimal overhead

// HTTPRegisterProfile registers the default pprof HTTP endpoints with the internal servenv mux.
func (sv *ServEnv) HTTPRegisterPprofProfile() { _ = "STUB: not implemented"; return }
