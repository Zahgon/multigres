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

package s3

import (
	"context"
)

// ValidationConfig contains the configuration for S3 validation
type ValidationConfig struct {
	Bucket       string
	Region       string
	Endpoint     string // Optional: for S3-compatible storage (e.g., s3mock)
	KeyPrefix    string // Optional: prefix for test object
	AccessKey    string
	SecretKey    string
	SessionToken string // Optional: for temporary credentials (AWS STS)
}

// isLocalhostEndpoint checks if an endpoint URL points to localhost
func isLocalhostEndpoint(endpoint string) bool { _ = "STUB: not implemented"; return false }

// ValidateBucketName validates S3 bucket naming rules
//
// https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html
func ValidateBucketName(bucket string) error { _ = "STUB: not implemented"; return nil }

// Basic validation - lowercase, numbers, hyphens, dots

// ValidateAccess tests S3 bucket access by writing and deleting a test object
func ValidateAccess(ctx context.Context, cfg ValidationConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// Create AWS config with explicit credentials

// session token for temporary credentials (AWS STS)

// Set custom endpoint if specified (for s3mock, etc.)

// For localhost endpoints only, skip TLS verification for self-signed certs
// This is needed for local development with s3mock

//nolint:gosec // Required for local s3mock with self-signed certs

// Create S3 client
// When using custom endpoints (s3mock, etc.), use path-style addressing
// This is required for S3-compatible storage that doesn't support virtual-hosted-style

// 1. HeadBucket - verify bucket exists and we have access

// 2. PutObject - write test file to verify write permissions

// 3. DeleteObject - clean up test file

// Don't fail if delete fails - write succeeded which is what matters
