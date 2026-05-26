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

package shardsetup

import (
	"net"
	"net/http"
	"sync"
	"testing"
	"time"

	tracepb "go.opentelemetry.io/proto/otlp/trace/v1"
)

// SpanProto is the protobuf Span type from the OTLP trace package.
// Exported as a convenience so test files don't need to import the proto package directly.
type SpanProto = tracepb.Span

// TestOTLPCollector is a lightweight in-process OTLP HTTP receiver that
// collects spans exported by child processes (e.g., multigateway). It
// listens on a random port and accepts POST /v1/traces with protobuf-encoded
// ExportTraceServiceRequest bodies.
type TestOTLPCollector struct {
	mu       sync.Mutex
	spans    []*tracepb.ResourceSpans
	server   *http.Server
	listener net.Listener
	endpoint string
}

// NewTestOTLPCollector starts a test OTLP HTTP collector on a random port.
// The collector is automatically stopped when the test completes.
func NewTestOTLPCollector(t *testing.T) *TestOTLPCollector { _ = "STUB: not implemented"; return nil }

// Do not use t.Logf here — the goroutine may outlive the test function,
// which causes a panic in Go 1.24+.

// Endpoint returns the collector's HTTP endpoint (e.g., "http://127.0.0.1:12345").
func (c *TestOTLPCollector) Endpoint() string {
	_ = "STUB: not implemented"

	// GetSpans returns all collected ResourceSpans.
	return ""
}

func (c *TestOTLPCollector) GetSpans() []*tracepb.ResourceSpans {
	_ = "STUB: not implemented"
	return nil
}

// WaitForSpans polls until at least count spans with the given name are
// collected, or the timeout expires.
func (c *TestOTLPCollector) WaitForSpans(t *testing.T, spanName string, count int, timeout time.Duration) []*tracepb.Span {
	_ = "STUB: not implemented"
	return nil
}

// findSpans returns all Span protos matching the given name.
func (c *TestOTLPCollector) findSpans(name string) []*tracepb.Span {
	_ = "STUB: not implemented"
	return nil
}

func (c *TestOTLPCollector) handleTraces(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Return success response

// FindSpanAttribute looks up an attribute by key on a span proto.
// Returns the string value and true if found, or ("", false) if not.
func FindSpanAttribute(span *tracepb.Span, key string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// Handle array values by returning a string representation

// FindSpanAttributeArray looks up a string array attribute by key on a span proto.
func FindSpanAttributeArray(span *tracepb.Span, key string) ([]string, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
