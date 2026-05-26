// Copyright 2019 The Vitess Authors.
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
// Modifications Copyright 2025 Supabase, Inc.

package mterrors

// This file contains functions to convert errors to and from gRPC codes.
// Use these methods to return an error through gRPC and still
// retain its code.

// truncateError shortens errors because gRPC has a size restriction on them.
func truncateError(err error) string {
	_ = "STUB: not implemented"
	// For more details see: https://github.com/grpc/grpc-go/issues/443
	// The gRPC spec says "Clients may limit the size of Response-Headers,
	// Trailers, and Trailers-Only, with a default of 8 KiB each suggested."
	// Therefore, we assume 8 KiB minus some headroom.
	return ""
}

// ToGRPC returns an error as a gRPC error, with the appropriate error code.
// If the error is a *PgDiagnostic, it includes the PgDiagnostic in the gRPC status details
// so that all PostgreSQL error fields are preserved through the RPC.
func ToGRPC(err error) error { _ = "STUB: not implemented"; return nil }

// Check if this is a PostgreSQL error

// Create gRPC status with RPCError containing the PgDiagnostic

// Attach the RPCError as a detail to the status

// Log a warning with context about the error being lost.
// This can happen if the error details are too large for gRPC limits.

// Fall back to basic error without PgDiagnostic details

// FromGRPC returns a gRPC error as a mterrors error, translating between error codes.
// If the gRPC error contains a PgDiagnostic in its details, it returns a *PgDiagnostic
// to preserve all PostgreSQL error fields.
// However, there are a few errors which are not translated and passed as they
// are. For example, io.EOF since our code base checks for this error to find
// out that a stream has finished.
func FromGRPC(err error) error { _ = "STUB: not implemented"; return nil }

// Do not wrap io.EOF because we compare against it for finished streams.

// Map gRPC context errors to PostgreSQL query_canceled errors.
// gRPC converts context.DeadlineExceeded / context.Canceled into status
// errors that don't wrap the original sentinels, so we return the proper
// PgDiagnostic directly.

// Check for RPCError in status details

// If PgDiagnostic is present, return it directly

// Otherwise use the RPCError message and code

// No RPCError details, fall back to basic conversion
