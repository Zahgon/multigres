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

package telemetry

import (
	"context"
)

// getTraceparent extracts the W3C Trace Context traceparent from the context.
func getTraceparent(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

// Extract trace context to W3C Trace Context format

// Get traceparent value (format: version-trace_id-span_id-flags)

// TraceparentEnvVar returns the TRACEPARENT environment variable string
// for propagating trace context to subprocesses, in the format "TRACEPARENT=value".
// Returns an empty string if the context has no valid span.
//
// Note: When using executil.Command(), trace propagation is handled automatically.
// This function is exported for cases where manual trace propagation is needed.
func TraceparentEnvVar(ctx context.Context) string { _ = "STUB: not implemented"; return "" }
