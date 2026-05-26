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

// go/tools/s3mock/storage.go
package s3mock

import (
	// #nosec G501 - MD5 used for S3 ETag calculation (non-cryptographic use)

	"errors"
	"sync"
	"time"
)

var (
	ErrBucketNotFound      = errors.New("bucket not found")
	ErrBucketAlreadyExists = errors.New("bucket already exists")
	ErrObjectNotFound      = errors.New("object not found")
	ErrUploadNotFound      = errors.New("upload not found")
)

// MultipartUpload tracks an in-progress multipart upload
type MultipartUpload struct {
	uploadID string
	bucket   string
	key      string
	parts    map[int]*UploadPart
	mu       sync.Mutex
}

// UploadPart represents one part of a multipart upload
type UploadPart struct {
	partNumber int
	etag       string
	size       int64
	data       []byte
}

// CompletedPart represents a part in CompleteMultipartUpload request
type CompletedPart struct {
	PartNumber int
	ETag       string
}

// Object represents an S3 object in storage
type Object struct {
	Key          string
	filePath     string // path to file on disk
	ETag         string
	LastModified time.Time
	Size         int64
}

// ReadData reads the object's data from disk
func (o *Object) ReadData() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Bucket represents an S3 bucket
type Bucket struct {
	mu      sync.RWMutex
	objects map[string]*Object
}

// Storage is the disk-based storage backend
type Storage struct {
	mu               sync.RWMutex
	buckets          map[string]*Bucket
	tempDir          string
	multipartUploads map[string]*MultipartUpload
}

// NewStorage creates a new disk-based storage backend
func NewStorage() *Storage { _ = "STUB: not implemented"; return nil }

// Close cleans up the storage, removing all temporary files
func (s *Storage) Close() error { _ = "STUB: not implemented"; return nil }

// CreateBucket creates a new bucket
func (s *Storage) CreateBucket(name string) error { _ = "STUB: not implemented"; return nil }

// BucketExists checks if a bucket exists
func (s *Storage) BucketExists(name string) bool { _ = "STUB: not implemented"; return false }

// PutObject stores an object on disk
func (s *Storage) PutObject(bucket, key string, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Calculate ETag (MD5 of data)
// #nosec G401 - MD5 used for S3 ETag calculation (non-cryptographic use)

// Write data to temp file

// Delete old file if object already exists

// GetObject retrieves an object
func (s *Storage) GetObject(bucket, key string) (*Object, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteObject deletes an object
func (s *Storage) DeleteObject(bucket, key string) error { _ = "STUB: not implemented"; return nil }

// ListObjects returns objects in a bucket with optional prefix and start-after filtering
func (s *Storage) ListObjects(bucket, prefix, startAfter string) []*Object {
	_ = "STUB: not implemented"
	return nil
}

// Filter by prefix

// Filter by start-after (lexicographically greater than)

// Sort by key for consistent ordering

// CreateMultipartUpload initiates a multipart upload
func (s *Storage) CreateMultipartUpload(bucket, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Generate unique upload ID (16 random bytes as hex)

// UploadPart uploads a part of a multipart upload
func (s *Storage) UploadPart(bucket, key, uploadID string, partNumber int, data []byte) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Verify bucket and key match

// Calculate ETag for this part (MD5)
// #nosec G401 - MD5 used for S3 ETag calculation

// CompleteMultipartUpload completes a multipart upload by assembling parts
func (s *Storage) CompleteMultipartUpload(bucket, key, uploadID string, parts []CompletedPart) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Verify bucket and key match

// Assemble parts in order

// Verify ETag matches

// Calculate final ETag (MD5 of assembled data)
// #nosec G401 - MD5 used for S3 ETag calculation

// Write assembled data to disk

// Store object

// Delete old file if object already exists

// Remove multipart upload

// AbortMultipartUpload aborts a multipart upload
func (s *Storage) AbortMultipartUpload(bucket, key, uploadID string) error {
	_ = "STUB: not implemented"
	return nil
}

// Verify bucket and key match

// Remove multipart upload
