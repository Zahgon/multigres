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

// Package grpcpoolerservice implements the gRPC server for MultiPooler
package grpcpoolerservice

import (
	"context"

	"github.com/multigres/multigres/go/common/servenv"
	multipoolerpb "github.com/multigres/multigres/go/pb/multipoolerservice"
	"github.com/multigres/multigres/go/services/multipooler/poolerserver"
	"github.com/multigres/multigres/go/services/multipooler/pubsub"
)

// poolerService is the gRPC wrapper for MultiPooler
type poolerService struct {
	multipoolerpb.UnimplementedMultiPoolerServiceServer
	pooler *poolerserver.QueryPoolerServer
	pubsub *pubsub.Listener
}

func RegisterPoolerServices(senv *servenv.ServEnv, grpc *servenv.GrpcServer) {
	_ = "STUB: not implemented"
	// Register ourselves to be invoked when the pooler starts
	return
}

// StreamExecute executes a SQL query and streams the results back to the client.
// This is the main execution method used by multigateway.
// When req.ReservationOptions has non-zero reasons, creates or extends a reserved connection.
func (s *poolerService) StreamExecute(req *multipoolerpb.StreamExecuteRequest, stream multipoolerpb.MultiPoolerService_StreamExecuteServer) error {
	_ = "STUB: not implemented"
	// Allow during shutdown if using an existing reserved connection.
	// For new reservations (ReservationOptions has reasons but no ReservedConnectionId),
	// block during shutdown since new reservations should not be created.
	return nil
}

// Validate reservation reasons at the gRPC trust boundary.

// Get the executor from the pooler

// Execute the query and stream results

// Send notices first (if any) as separate diagnostic messages

// Send row data (if any)

// Send final message with reserved state if on a reserved connection.
// The send error is intentionally discarded: if the stream is already broken
// the gateway will clean up via ReleaseReservedConnection on client disconnect.

// Convert errors to gRPC format, preserving PostgreSQL error details

// ExecuteQuery executes a SQL query and returns the result
// This should be used sparingly only when we know the result set is small,
// otherwise StreamExecute should be used.
func (s *poolerService) ExecuteQuery(ctx context.Context, req *multipoolerpb.ExecuteQueryRequest) (*multipoolerpb.ExecuteQueryResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get the executor from the pooler

// Execute the query

// Convert errors to gRPC format, preserving PostgreSQL error details

// GetAuthCredentials retrieves authentication credentials (SCRAM hash) for a PostgreSQL user.
// This is used by multigateway to authenticate clients using SCRAM-SHA-256.
//
// This method uses an admin connection directly since normally a non-superuser wouldn't
// have access to password hashes and at the time of this request we wouldn't have authenticated
// that we have permission to run queries under any other user's role.
func (s *poolerService) GetAuthCredentials(ctx context.Context, req *multipoolerpb.GetAuthCredentialsRequest) (*multipoolerpb.GetAuthCredentialsResponse, error) {
	_ = "STUB: not implemented"
	// Validate request.
	return nil, nil
}

// Time the full credential-query path (admin acquire + pg_authid
// lookup + decode) for mg.pooler.auth.credential_query.duration so
// admin-pool contention shows up before it cascades into
// gateway-visible auth latency. The duration is recorded on every
// exit, error or not; the error counter only fires on failures so
// the success rate is implicit.

// An admin connection:
// - has permission to read password hashes
// - also avoids a chicken-egg scenario of needing to create and use a role-specific connection
//   to figure out if the caller should have access to that role-specific connection.

// Get the role auth info (password hash + rolreplication) using the admin
// connection. This queries pg_authid, which requires superuser access.

// Emit as a PgDiagnostic so the SQLSTATE (28000) is the
// distinguishing signal at the gateway, not the gRPC code.
// gRPC auth interceptors use codes.PermissionDenied /
// codes.Unauthenticated for transport failures; keying on code
// alone would misclassify an mTLS or authz error as an app-level
// "role not permitted to log in" rejection to the end user.

// SQLSTATE 28P01 matches PG's opaque "password authentication
// failed" error for expired passwords. PgDiagnostic detail
// survives the gRPC round trip and the gateway matches on it.

// Genuine DB-level failure (SQL error, unexpected result shape,
// connection drop mid-query). Tagged db_error so operators can
// alert on it without false positives from the user_not_found
// baseline.

// Describe returns metadata about a prepared statement or portal.
// Used by multigateway for the Extended Query Protocol.
func (s *poolerService) Describe(ctx context.Context, req *multipoolerpb.DescribeRequest) (*multipoolerpb.DescribeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get the executor from the pooler

// Call the executor's Describe method

// Convert errors to gRPC format, preserving PostgreSQL error details

// PortalStreamExecute executes a portal (bound prepared statement) and streams results.
// Used by multigateway for the Extended Query Protocol.
func (s *poolerService) PortalStreamExecute(req *multipoolerpb.PortalStreamExecuteRequest, stream multipoolerpb.MultiPoolerService_PortalStreamExecuteServer) error {
	_ = "STUB: not implemented"
	return nil
}

// Get the executor from the pooler

// Execute the portal and stream results

// Send notices first (if any) as separate diagnostic messages

// Send row data (if any)

// Note: When PortalStreamExecute returns an error, it also releases any reserved
// connection and returns an empty ReservedState. So we don't need to send a
// reserved connection ID in the error case.
// Convert errors to gRPC format, preserving PostgreSQL error details

// Send final response with reserved connection ID if one was created

// CopyBidiExecute handles bidirectional streaming operations (e.g., COPY commands).
// The gateway sends: INITIATE → DATA (repeated) → DONE/FAIL
// The pooler responds: READY → DATA (for COPY TO) → RESULT/ERROR
func (s *poolerService) CopyBidiExecute(stream multipoolerpb.MultiPoolerService_CopyBidiExecuteServer) error {
	_ = "STUB: not implemented"
	return nil

	// Receive INITIATE message first so we can check reserved connection ID
}

// Get the executor from the pooler

// Phase 1: INITIATE - Send COPY command and get reserved connection

// CopyReady returns a non-nil reservedState when the COPY query was
// rejected by PostgreSQL (e.g., invalid column, conflicting options)
// but the existing reserved connection is still alive and holding
// other reasons. Forward that state to the gateway via an ERROR phase
// response so the gateway keeps tracking the reserved conn; without
// this, the gateway would see only the gRPC status error and the next
// statement would fail with "reserved connection not found".

// Convert columnFormats from []int16 to []int32 for protobuf

// Send READY response with reserved connection info

// Clean up reserved connection on send failure

// Build options with reserved connection ID for subsequent calls

// Capture target from INITIATE for use in error paths where req may be nil.

// Phase 2: Handle DATA/DONE/FAIL messages

// Stream closed or error — abort COPY and send best-effort ERROR response
// so the gateway can update its shard state even if the stream is degraded.
// Note: req may be nil when Recv fails, so we use initiateTarget.

// Phase 2a: DATA - Write data chunk to PostgreSQL

// Send ERROR response with reserved state so gateway can update shard state

// Phase 2b: DONE - Finalize COPY operation

// CopyFinalize has already done its own cleanup:
//   - PG ErrorResponse: ReadyForQuery was drained, the COPY
//     reason was removed, and the conn was either released
//     (no other reasons) or kept with the returned state.
//   - Connection-level failure: conn was released, state is nil.
// Either way, calling CopyAbort here would either be a no-op
// (conn already released) or actively poison a clean conn by
// writing CopyFail on a backend already back in RFQ. Forward
// the state CopyFinalize returned.

// Send RESULT response with final result and reserved state

// Operation completed successfully

// Phase 2c: FAIL - Abort COPY operation

// Send ERROR response with reserved state

// Send ERROR response with reserved state so gateway can update shard state

// ConcludeTransaction concludes a transaction on a reserved connection.
// Executes COMMIT or ROLLBACK based on the conclusion. Returns remaining reasons if connection is still reserved.
func (s *poolerService) ConcludeTransaction(ctx context.Context, req *multipoolerpb.ConcludeTransactionRequest) (*multipoolerpb.ConcludeTransactionResponse, error) {
	_ = "STUB: not implemented"
	// Always on existing reserved connection, allow during shutdown.
	return nil, nil
}

// Get the executor from the pooler

// Conclude the transaction. Forward the per-txn portal-release diff so the
// executor can drop exactly the cursor pins PG closed for this ROLLBACK
// (or fall back to ReleaseAllPortals when release_all_portals is true,
// e.g. for older gateways that don't compute the diff).

// DiscardTempTables sends DISCARD TEMP on a reserved connection and removes the temp table reason.
// Returns remaining reasons if connection is still reserved.
func (s *poolerService) DiscardTempTables(ctx context.Context, req *multipoolerpb.DiscardTempTablesRequest) (*multipoolerpb.DiscardTempTablesResponse, error) {
	_ = "STUB: not implemented"
	// Always on existing reserved connection, allow during shutdown.
	return nil, nil
}

// Get the executor from the pooler

// ReleaseReservedConnection forcefully releases a reserved connection regardless of reason.
func (s *poolerService) ReleaseReservedConnection(ctx context.Context, req *multipoolerpb.ReleaseReservedConnectionRequest) (*multipoolerpb.ReleaseReservedConnectionResponse, error) {
	_ = "STUB: not implemented"
	// Always on existing reserved connection, allow during shutdown.
	return nil, nil
}

// StreamPoolerHealth streams health updates to the client.
// Sends an initial health state immediately, then updates when state changes.
func (s *poolerService) StreamPoolerHealth(req *multipoolerpb.StreamPoolerHealthRequest, stream multipoolerpb.MultiPoolerService_StreamPoolerHealthServer) error {
	_ = "STUB: not implemented"
	return nil

	// Check if pooler is initialized
}

// Get the health provider

// Subscribe to health updates

// Send initial health state

// Stream updates until client disconnects or context is cancelled

// Channel closed, stream ended

// healthStateToProto converts internal health state to proto response.
func healthStateToProto(state *poolerserver.HealthState) *multipoolerpb.StreamPoolerHealthResponse {
	_ = "STUB: not implemented"
	return nil
}

// StreamNotifications streams async notifications for a subscribed channel.
func (s *poolerService) StreamNotifications(
	req *multipoolerpb.StreamNotificationsRequest,
	stream multipoolerpb.MultiPoolerService_StreamNotificationsServer,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Send an empty response as a "ready" signal — all channels are now LISTENed.
