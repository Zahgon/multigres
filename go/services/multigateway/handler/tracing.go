// Copyright 2026 Supabase, Inc.
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

package handler

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

// startQuerySpan creates a server-side span for a gateway query operation.
// The span follows OTel database semantic conventions and does NOT include
// db.query.text for security reasons (user queries may contain PII).
func startQuerySpan(
	ctx context.Context,
	operationName string,
	protocol string,
	dbNamespace string,
	user string,
) (context.Context, trace.Span) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(trace.Span)
}

// setSpanPlanAttributes enriches the active span with plan-level metadata.
// Called from recordQueryCompletion after the executor returns.
func setSpanPlanAttributes(ctx context.Context, planType string, tablesUsed []string) {
	_ = "STUB: not implemented"
	return
}

// recordSpanError records an error on a span with its SQLSTATE code.
func recordSpanError(span trace.Span, err error, sqlstate string) {
	_ = "STUB: not implemented"
	return
}
