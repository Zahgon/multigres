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

// go/tools/s3mock/server.go

// Package s3mock provides an S3-compatible mock server for testing.
//
// # Logging
//
// By default, HTTP request logging is disabled to reduce noise in CI.
// To enable logging, set the environment variable:
//
//	MULTIGRES_TEST_LOG_S3MOCK=1
//
// Accepted values: 1, true, yes (case-insensitive)
package s3mock

import (
	"context"
	"net"
	"net/http"
)

// PutCallback is called before each PutObject or UploadPart request.
// It can block to pause the upload or return an error to reject it.
// The ctx parameter is the request context — cancellation unblocks automatically.
type PutCallback func(ctx context.Context, bucket, key string) error

// ServerOption configures a Server.
type ServerOption func(*Server)

// WithPutCallback returns a ServerOption that installs cb as the PutCallback.
// cb is called before every PutObject and UploadPart request.
func WithPutCallback(cb PutCallback) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

// Server is an S3-compatible mock server
type Server struct {
	storage     *Storage
	server      *http.Server
	listener    net.Listener
	endpoint    string
	putCallback PutCallback // optional; called before each PutObject/UploadPart
}

// NewServer creates and starts a new S3 mock server on the specified port
func NewServer(port int, opts ...ServerOption) (*Server, error) {
	_ = "STUB: not implemented"
	return nil,

		// Generate TLS config
		nil
}

// Create listener on specified port

// Extract actual port from listener (important when port=0)

// Build endpoint URL with actual port

// Create server

// Create HTTP server with router

// Start server in background

// loggingMiddleware wraps a handler with request logging
func loggingMiddleware(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// isLoggingEnabled checks if S3mock logging should be enabled based on environment variable
func isLoggingEnabled() bool { _ = "STUB: not implemented"; return false }

// Accept: 1, true, True, TRUE, yes, Yes, YES

// router creates the HTTP request router
func (s *Server) router() http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

// Route based on method and query params

// Check for multipart upload operations

// DeleteObjects (batch delete)

// CreateMultipartUpload

// CompleteMultipartUpload

// Create bucket

// Call PutCallback before any upload (PutObject or UploadPart)

// UploadPart

// PutObject

// AbortMultipartUpload

// DeleteObject

// Only apply logging middleware if enabled

// Endpoint returns the HTTPS endpoint URL
func (s *Server) Endpoint() string {
	_ = "STUB: not implemented"

	// CreateBucket creates a bucket in the mock storage
	return ""
}

func (s *Server) CreateBucket(name string) error { _ = "STUB: not implemented"; return nil }

// ListKeys returns the object keys in bucket whose names start with prefix.
// Intended for test assertions; the result is sorted lexicographically.
func (s *Server) ListKeys(bucket, prefix string) []string { _ = "STUB: not implemented"; return nil }

// GetObjectBytes returns the contents of bucket/key, intended for test
// assertions. Returns an error if the object does not exist.
func (s *Server) GetObjectBytes(bucket, key string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Stop gracefully stops the server and cleans up storage
func (s *Server) Stop() error { _ = "STUB: not implemented"; return nil }

// Shutdown HTTP server

// Clean up storage
