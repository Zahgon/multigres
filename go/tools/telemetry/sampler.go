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

package telemetry

import (
	"go.opentelemetry.io/otel/attribute"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// SamplingConfig represents the YAML configuration for operation-aware sampling
type SamplingConfig struct {
	Categories map[string]CategoryConfig `yaml:"categories"`
	GRPC       GRPCSpanConfig            `yaml:"grpc"`
	HTTP       HTTPSpanConfig            `yaml:"http"`
	Spans      SpanConfig                `yaml:"spans"`
}

// CategoryConfig defines the sampling probability for a category
type CategoryConfig struct {
	Probability float64 `yaml:"probability"`
}

// GRPCSpanConfig configures sampling for gRPC spans
// Services provides service-level defaults (e.g., "/package.Service")
// Methods provides method-level overrides (e.g., "/package.Service/Method")
type GRPCSpanConfig struct {
	Services map[string]string `yaml:"services"`
	Methods  map[string]string `yaml:"methods"`
}

// HTTPSpanConfig configures sampling for HTTP spans
// Exact provides exact matches (e.g., "GET /live")
// Patterns provides glob patterns (e.g., "GET /api/*")
type HTTPSpanConfig struct {
	Exact    map[string]string `yaml:"exact"`
	Patterns map[string]string `yaml:"patterns"`
}

// SpanConfig configures sampling for manually created spans
// Exact provides exact matches
// Patterns provides glob patterns
type SpanConfig struct {
	Exact    map[string]string `yaml:"exact"`
	Patterns map[string]string `yaml:"patterns"`
}

// loadSamplingConfig loads and validates sampling configuration from a YAML file
func loadSamplingConfig(path string) (*SamplingConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate categories

// Ensure "default" category exists

// Validate probabilities are in [0, 1]

// Validate all span mappings reference defined categories

// ConfigurableSampler wraps base samplers and routes sampling decisions based on
// a file-based configuration that maps span names to categories with different probabilities.
// This enables fine-grained control over sampling rates for different operation types.
type ConfigurableSampler struct {
	config     *SamplingConfig
	samplers   map[string]sdktrace.Sampler // category name -> sampler
	defaultCat string                      // default category name
}

// getAttributeValue finds an attribute value by key from the span attributes
func getAttributeValue(attrs []attribute.KeyValue, key string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// getCategoryForSpan determines the sampling category using span metadata.
// It uses SpanKind and semantic convention attributes to robustly identify span types:
//   - gRPC spans: identified by rpc.system="grpc" attribute
//   - HTTP spans: identified by http.method attribute and SpanKind.Server
//   - Other spans: use span name for matching
//
// Priority order:
//  1. gRPC method (from rpc.service + rpc.method attributes)
//  2. gRPC service (from rpc.service attribute)
//  3. HTTP exact match (from http.method + http.target)
//  4. HTTP pattern match
//  5. Manual span exact match (from span name)
//  6. Manual span pattern match
//  7. Default category
func (s *ConfigurableSampler) getCategoryForSpan(params sdktrace.SamplingParameters) string {
	_ = "STUB: not implemented"
	// Check for gRPC span using semantic conventions
	// otelgrpc sets: rpc.system="grpc", rpc.service, rpc.method
	return ""
}

// Check method-level configuration (e.g., "/package.Service/Method")

// Check service-level configuration (e.g., "/package.Service")

// gRPC span but no config match - use default

// Check for HTTP span using semantic conventions
// otelhttp sets: http.method, http.target (or http.route)

// Prefer http.route (pattern) over http.target (actual path) for better matching

// Check exact match (e.g., "GET /live")

// Check pattern match using filepath.Match

// HTTP span but no config match - use default

// Not gRPC or HTTP - treat as manual span, use span name for matching

// TODO: Consider adding custom semantic conventions for multigres-specific operations
// to make matching more robust and explicit. For example:
//   - multigres.operation.type: "maintenance", "recovery", "background"
//   - multigres.operation.priority: "critical", "normal", "low"
// This would allow matching by attributes instead of span names, making configuration
// safer and less fragile than pattern matching. Example:
//   if opType, ok := getAttributeValue(params.Attributes, "multigres.operation.type"); ok {
//       // Use opType for category lookup
//   }

// Check exact match

// Check pattern match using filepath.Match

// Fall back to default category

func (s *ConfigurableSampler) ShouldSample(params sdktrace.SamplingParameters) sdktrace.SamplingResult {
	_ = "STUB: not implemented"
	// Determine category for this span using full span metadata
	return *new(sdktrace.SamplingResult)
}

// Get sampler for this category

// Should not happen due to validation, but fallback to default

// Delegate to category sampler

func (s *ConfigurableSampler) Description() string { _ = "STUB: not implemented"; return "" }

// maybeCreateCustomSampler creates a custom file-based sampler if configured, otherwise returns nil
// to defer to standard OpenTelemetry environment variable handling.
//
// Configuration:
//
//   - OTEL_TRACES_SAMPLER: sampler type
//
//   - "multigres_custom": use file-based configuration (reads OTEL_TRACES_SAMPLER_CONFIG)
//
//   - other values: defer to standard OTEL handling (always_on, always_off, parentbased_traceidratio, etc.)
//
//   - OTEL_TRACES_SAMPLER_CONFIG: path to YAML config file (only used when OTEL_TRACES_SAMPLER="multigres_custom")
//
// Returns (nil, nil) if standard OTEL sampling should be used, allowing OTEL SDK to handle
// all standard sampler types automatically via environment variables.
// Returns (nil, error) if custom sampling is requested but configuration fails.
func maybeCreateCustomSampler() (sdktrace.Sampler, error) {
	_ = "STUB: not implemented"
	return *new(sdktrace.Sampler), nil
}

// Only handle "multigres_custom" sampler type - defer everything else to OTEL

// Load file-based configuration

// Create samplers for each category

// Wrap with ParentBased to ensure complete distributed traces
