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

// go/tools/s3mock/operations.go
package s3mock

import (
	"encoding/xml"
	"net/http"
)

// S3Error represents an S3 error response
type S3Error struct {
	XMLName   xml.Name `xml:"Error"`
	Code      string   `xml:"Code"`
	Message   string   `xml:"Message"`
	RequestId string   `xml:"RequestId"`
}

// writeS3Error writes an S3-compliant error response
func writeS3Error(w http.ResponseWriter, code string, message string, statusCode int) {
	_ = "STUB: not implemented"
	return
}

// parseRange parses HTTP Range header and returns start, end offsets
// Supports "bytes=start-end" and "bytes=-suffix" formats
// Returns start=-1, end=-1 if no range or invalid range
func parseRange(rangeHeader string, contentLength int64) (start int64, end int64) {
	_ = "STUB: not implemented"
	// No range header
	return 0, 0
}

// Must start with "bytes="

// Handle suffix range (bytes=-N means last N bytes)

// Handle normal range (bytes=start-end)

// Parse start

// Parse end (can be empty for "bytes=start-")

// Clamp to content length

// parseBucketAndKey extracts bucket name and object key from URL path
func parseBucketAndKey(urlPath string) (bucket, key string) {
	_ = "STUB: not implemented"
	// Remove leading slash
	return "", ""
}

// Split into bucket and key

// handleHeadBucket handles HEAD /{bucket} requests
func handleHeadBucket(w http.ResponseWriter, r *http.Request, storage *Storage) {
	_ = "STUB: not implemented"
	return
}

// handlePutObject handles PUT /{bucket}/{key} requests
func handlePutObject(w http.ResponseWriter, r *http.Request, storage *Storage) {
	_ = "STUB: not implemented"
	return
}

// Read request body

// Store object

// Get object to retrieve ETag

// handleGetObject handles GET /{bucket}/{key} requests
func handleGetObject(w http.ResponseWriter, r *http.Request, storage *Storage) {
	_ = "STUB: not implemented"
	return
}

// Read full object data

// Check for Range header

// If range is valid, return partial content (206)

// Extract range from data

// Set headers for partial content

// 206

// No range or invalid range - return full object (200)

// handleHeadObject handles HEAD /{bucket}/{key} requests
func handleHeadObject(w http.ResponseWriter, r *http.Request, storage *Storage) {
	_ = "STUB: not implemented"
	return
}

// Set headers but no body

// handleDeleteObject handles DELETE /{bucket}/{key} requests
func handleDeleteObject(w http.ResponseWriter, r *http.Request, storage *Storage) {
	_ = "STUB: not implemented"
	return
}

// S3 returns 204 even if object doesn't exist

// ListBucketResult represents the XML response for ListObjectsV2
type ListBucketResult struct {
	XMLName               xml.Name       `xml:"http://s3.amazonaws.com/doc/2006-03-01/ ListBucketResult"`
	Name                  string         `xml:"Name"`
	Prefix                string         `xml:"Prefix"`
	Delimiter             string         `xml:"Delimiter,omitempty"`
	MaxKeys               int            `xml:"MaxKeys"`
	KeyCount              int            `xml:"KeyCount"`
	IsTruncated           bool           `xml:"IsTruncated"`
	ContinuationToken     string         `xml:"ContinuationToken,omitempty"`
	NextContinuationToken string         `xml:"NextContinuationToken,omitempty"`
	StartAfter            string         `xml:"StartAfter,omitempty"`
	Contents              []ObjectResult `xml:"Contents"`
	CommonPrefixes        []CommonPrefix `xml:"CommonPrefixes,omitempty"`
}

type ObjectResult struct {
	Key          string `xml:"Key"`
	LastModified string `xml:"LastModified"`
	ETag         string `xml:"ETag"`
	Size         int64  `xml:"Size"`
}

type CommonPrefix struct {
	Prefix string `xml:"Prefix"`
}

// InitiateMultipartUploadResult represents XML response for CreateMultipartUpload
type InitiateMultipartUploadResult struct {
	XMLName  xml.Name `xml:"InitiateMultipartUploadResult"`
	Bucket   string   `xml:"Bucket"`
	Key      string   `xml:"Key"`
	UploadId string   `xml:"UploadId"`
}

// CompleteMultipartUploadResult represents XML response for CompleteMultipartUpload
type CompleteMultipartUploadResult struct {
	XMLName  xml.Name `xml:"CompleteMultipartUploadResult"`
	Location string   `xml:"Location"`
	Bucket   string   `xml:"Bucket"`
	Key      string   `xml:"Key"`
	ETag     string   `xml:"ETag"`
}

// CompleteMultipartUploadRequest represents XML request for CompleteMultipartUpload
type CompleteMultipartUploadRequest struct {
	XMLName xml.Name       `xml:"CompleteMultipartUpload"`
	Parts   []CompletePart `xml:"Part"`
}

// CompletePart represents a part in CompleteMultipartUpload XML request
type CompletePart struct {
	PartNumber int    `xml:"PartNumber"`
	ETag       string `xml:"ETag"`
}

// DeleteRequest represents the XML request for DeleteObjects
type DeleteRequest struct {
	XMLName xml.Name         `xml:"Delete"`
	Objects []ObjectToDelete `xml:"Object"`
}

// ObjectToDelete represents an object to delete
type ObjectToDelete struct {
	Key string `xml:"Key"`
}

// DeleteResult represents the XML response for DeleteObjects
type DeleteResult struct {
	XMLName xml.Name        `xml:"DeleteResult"`
	Deleted []DeletedObject `xml:"Deleted"`
}

// DeletedObject represents a successfully deleted object
type DeletedObject struct {
	Key string `xml:"Key"`
}

// handleListObjectsV2 handles GET /{bucket}?list-type=2 requests
func handleListObjectsV2(w http.ResponseWriter, r *http.Request, storage *Storage) {
	_ = "STUB: not implemented"
	return
}

// Parse query parameters

// Parse max-keys (default 1000, AWS limit)

// If continuation token is present, it takes precedence over start-after

// Get all matching objects

// Process objects with delimiter

// Find the next delimiter after the prefix

// This object is in a "subdirectory" - add to common prefixes

// This object is at the current "level" - add to contents

// No delimiter - return objects up to maxKeys

// Use the last key we included as the continuation token

// Build response

// Include continuation tokens if paginating

// handleCreateMultipartUpload handles POST /{bucket}/{key}?uploads requests
func handleCreateMultipartUpload(w http.ResponseWriter, r *http.Request, storage *Storage) {
	_ = "STUB: not implemented"
	return
}

// handleUploadPart handles PUT /{bucket}/{key}?partNumber=N&uploadId=ID requests
func handleUploadPart(w http.ResponseWriter, r *http.Request, storage *Storage) {
	_ = "STUB: not implemented"
	return
}

// Read part data

// handleCompleteMultipartUpload handles POST /{bucket}/{key}?uploadId=ID requests
func handleCompleteMultipartUpload(w http.ResponseWriter, r *http.Request, storage *Storage) {
	_ = "STUB: not implemented"
	return
}

// Parse XML request body

// Convert XML parts to storage CompletedPart

// Remove quotes from ETag if present

// handleAbortMultipartUpload handles DELETE /{bucket}/{key}?uploadId=ID requests
func handleAbortMultipartUpload(w http.ResponseWriter, r *http.Request, storage *Storage) {
	_ = "STUB: not implemented"
	return
}

// handleDeleteObjects handles POST /{bucket}?delete requests
func handleDeleteObjects(w http.ResponseWriter, r *http.Request, storage *Storage) {
	_ = "STUB: not implemented"
	return
}

// Parse XML request

// Delete each object

// S3 doesn't return errors for missing objects
