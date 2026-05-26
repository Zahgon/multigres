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

// Package telemetry can help with annotating and exporting metrics, logs, traces, and exemplars.
//
// To start a cluster with the local provisioner configured to export traces and metrics:
//
//	OTEL_EXPORTER_OTLP_PROTOCOL="http/protobuf" \
//	  OTEL_METRICS_EXPORTER=otlp \
//	  OTEL_EXPORTER_OTLP_METRICS_ENDPOINT="http://localhost:9090/api/v1/otlp/v1/metrics" \
//	  OTEL_EXPORTER_OTLP_ENDPOINT="http://localhost:4318" \
//	  OTEL_TRACES_SAMPLER=always_on \
//	  OTEL_TRACES_EXPORTER=otlp \
//	  multigres cluster start --config-path multigres_local
//
// To collect traces locally to view at http://localhost:16686/:
//
//	$ docker run --rm -it --name jaeger-all-in-one \
//	    -e COLLECTOR_OTLP_ENABLED=true \
//	    -e COLLECTOR_OTLP_HTTP_PORT=4318 \
//	    -p 16686:16686 \
//	    -p 4318:4318 \
//	    jaegertracing/all-in-one:latest
//
// To collect metrics locally to view at http://localhost:9090/:
//
//	$ docker run --rm -it \
//	    --name prometheus \
//	    -p 9090:9090 \
//	    prom/prometheus \
//	    --config.file=/etc/prometheus/prometheus.yml \
//	    --web.enable-otlp-receiver \
//	    --enable-feature=exemplar-storage
package telemetry

import (
	"context"
	"log/slog"
	"sync"

	"github.com/spf13/cobra"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

// TODO(dweitzman): Do we want package-specific tracing services, or is a shared
// one for all of multigres fine?
const tracingServiceName = "github.com/multigres/multigres"

var tracer = otel.Tracer(tracingServiceName)

// Tracer returns a tracer for creating spans named github.com/multigres/multigres
func Tracer() trace.Tracer {
	_ = "STUB: not implemented"

	// Telemetry holds OpenTelemetry configuration and state
	return *new(trace.Tracer)
}

type Telemetry struct {
	// State
	mu             sync.Mutex
	tracerProvider *sdktrace.TracerProvider
	meterProvider  *sdkmetric.MeterProvider
	loggerProvider *sdklog.LoggerProvider
	initialized    bool

	// Test overrides (only used in tests)
	testSpanExporter sdktrace.SpanExporter
	testMetricReader sdkmetric.Reader
	testLogProcessor sdklog.Processor
}

// NewTelemetry creates a new Telemetry instance
func NewTelemetry() *Telemetry { _ = "STUB: not implemented"; return nil }

// WithTestExporters configures the telemetry instance to use test exporters instead of autoexport.
// This allows tests to capture and verify telemetry data while still going through normal initialization.
// Must be called before InitTelemetry().
func (t *Telemetry) WithTestExporters(spanExporter sdktrace.SpanExporter, metricReader sdkmetric.Reader, logProcessor sdklog.Processor) *Telemetry {
	_ = "STUB: not implemented"
	return nil
}

// InitTelemetry initializes OpenTelemetry providers and exporters.
// The serviceName parameter sets the service.name resource attribute (can be overridden by OTEL_SERVICE_NAME env var).
// Additional OTel resource attributes can be passed via the attrs variadic parameter.
//
// Configuration is done via standard OpenTelemetry environment variables.
func (t *Telemetry) InitTelemetry(ctx context.Context, serviceName string, attrs ...attribute.KeyValue) error {
	_ = "STUB: not implemented"
	return nil
}

// Determine service name (env var > parameter)

// Create resource with service name, environment-provided attributes, and any
// additional Multigres service identity attributes. We don't merge with
// resource.Default() to avoid schema version conflicts.

// Instrument the default HTTP client for automatic tracing and metrics of outgoing HTTP requests
// This must happen AFTER both tracing and metrics are initialized so otelhttp can capture
// the correct TracerProvider and MeterProvider

// Set up trace context propagation

// initTracing initializes the TracerProvider using autoexport
// The exporter is automatically configured based on OTEL_TRACES_EXPORTER and OTEL_EXPORTER_OTLP_PROTOCOL
func (t *Telemetry) initTracing(ctx context.Context, res *resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Use test exporter if provided, otherwise use autoexport

// Default to "none" if OTEL_TRACES_EXPORTER is not explicitly set
// This prevents unwanted data export when telemetry is not explicitly configured

// TODO(dweitzman): For "multigres cluster start" we may want different tracing settings for
// the "multigres" command vs for the long-running services it starts. For example, maybe
// the multigres command itself should have tracing at 100% but the services should have tracing
// at a lower sample rate.
//
// Also, different commands may want to export their telemetry data in different ways. They can't
// all use the same port for a Prometheus exporter, for example.

// Create TracerProvider with batch span processor (or syncer for tests)
// Batch processing reduces overhead by grouping spans before export

// Use synchronous export for tests to avoid timing issues

// Create sampler - either custom file-based or defer to OTEL defaults

// Only set custom sampler if one was created
// Otherwise OTEL will use its default sampler based on environment variables

// initMetrics initializes the MeterProvider with dual exporters (autoexport + Prometheus)
func (t *Telemetry) initMetrics(ctx context.Context, res *resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Use test metric reader if provided, otherwise use autoexport

// Default to "none" if OTEL_METRICS_EXPORTER is not explicitly set
// This prevents unwanted data export when telemetry is not explicitly configured

// TODO(dweitzman): Add an additional prometheus exporter that's always at /metrics for debugging

// Configured via env vars or test reader

// Set global meter provider

// initLogs initializes the LoggerProvider using autoexport.
// The exporter is automatically configured based on OTEL_LOGS_EXPORTER and OTEL_EXPORTER_OTLP_PROTOCOL.
func (t *Telemetry) initLogs(ctx context.Context, res *resource.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// Use test processor if provided, otherwise use autoexport

// For tests, use simple processor with sync export

// Default to "none" if OTEL_LOGS_EXPORTER is not explicitly set
// This prevents unwanted data export when telemetry is not explicitly configured

// Check if exporter is "none" (no-op)

// Skip LoggerProvider creation for none exporter
// This avoids overhead when logs export is disabled

// Create LoggerProvider with batch processor
// Batch processing reduces overhead by grouping log records before export

// WithEnvTraceparent parses the TRACEPARENT env variable and returns a context within that
// parent
func (t *Telemetry) WithEnvTraceparent(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// Parse W3C Trace Context format: version-trace_id-span_id-flags
// Example: 00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01

// InitForCommand initializes telemetry for CLI commands with just a service name.
// Use this for one-shot commands that don't need additional resource attributes.
func (t *Telemetry) InitForCommand(cmd *cobra.Command, serviceName string, startSpan bool) (trace.Span, error) {
	_ = "STUB: not implemented"
	return *new(trace.Span), nil
}

// GetTracerProvider returns the configured TracerProvider.
func (t *Telemetry) GetTracerProvider() trace.TracerProvider {
	_ = "STUB: not implemented"
	return *new(trace.TracerProvider)
}

// GetMeterProvider returns the configured MeterProvider.
func (t *Telemetry) GetMeterProvider() metric.MeterProvider {
	_ = "STUB: not implemented"
	return *new(metric.MeterProvider)
}

// ShutdownTelemetry gracefully shuts down all telemetry providers
// This ensures all pending spans and metrics are flushed before the service exits
func (t *Telemetry) ShutdownTelemetry(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Shutdown tracer provider

// Shutdown meter provider

// Shutdown logger provider

// Mark as not initialized so subsequent calls are no-ops

// WrapSlogHandler wraps an slog.Handler to:
// 1. Inject trace context (trace_id, span_id) into log records
// 2. Bridge to OpenTelemetry logs SDK for OTLP export (if configured)
//
// The resulting handler maintains dual output:
// - Local logging via the wrapped handler (stdout/stderr/file)
// - OTLP export via OpenTelemetry LoggerProvider (if configured)
func (t *Telemetry) WrapSlogHandler(handler slog.Handler) slog.Handler {
	_ = "STUB: not implemented"
	// First, wrap with trace context injection
	return *new(slog.Handler)
}

// Then, if LoggerProvider is configured, add OTel bridge

// Create otelslog handler for OTLP export
// This bridges slog records to OpenTelemetry log records

// Compose: local logging + trace context + OTLP export

// If no LoggerProvider, just return trace context injection

// compositeHandler sends log records to both local and OTel handlers
type compositeHandler struct {
	local slog.Handler // Local logging (stdout/stderr/file)
	otel  slog.Handler // OpenTelemetry bridge for OTLP export
}

func (h *compositeHandler) Enabled(ctx context.Context, level slog.Level) bool {
	_ = "STUB: not implemented"
	// Enabled if either handler is enabled
	return false
}

func (h *compositeHandler) Handle(ctx context.Context, r slog.Record) error {
	_ = "STUB: not implemented"
	// Send to handlers that are enabled for this level
	// Check Enabled() before Handle() to avoid overhead for disabled levels
	return nil
}

func (h *compositeHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

func (h *compositeHandler) WithGroup(name string) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

// traceHandler wraps an slog.Handler to inject trace_id and span_id from context
type traceHandler struct {
	wrapped slog.Handler
}

func (h *traceHandler) Enabled(ctx context.Context, level slog.Level) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *traceHandler) Handle(ctx context.Context, r slog.Record) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *traceHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

func (h *traceHandler) WithGroup(name string) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

// WithSpan starts a child span named name, calls fn with the span context,
// records any error on the span, and ends the span. It returns fn's error.
func WithSpan(ctx context.Context, name string, fn func(context.Context) error) error {
	_ = "STUB: not implemented"
	return nil
}
