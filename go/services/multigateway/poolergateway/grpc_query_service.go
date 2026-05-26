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

package poolergateway

import (
	"context"
	"log/slog"
	"sync"

	"github.com/multigres/multigres/go/common/queryservice"
	"github.com/multigres/multigres/go/common/sqltypes"
	"github.com/multigres/multigres/go/pb/multipoolerservice"
	querypb "github.com/multigres/multigres/go/pb/query"

	"google.golang.org/grpc"
)

// grpcQueryService implements queryservice.QueryService using gRPC to communicate with a multipooler instance.
// This is a private implementation used internally by PoolerGateway.
type grpcQueryService struct {
	// conn is the gRPC connection to the multipooler
	conn *grpc.ClientConn

	// client is the generated gRPC client
	client multipoolerservice.MultiPoolerServiceClient

	// logger for debugging
	logger *slog.Logger

	// poolerID for logging
	poolerID string

	// copyStreamsMu protects copyStreams map
	copyStreamsMu sync.Mutex

	// copyStreams maps reserved connection IDs to active bidirectional streams for COPY operations
	copyStreams map[uint64]multipoolerservice.MultiPoolerService_CopyBidiExecuteClient
}

// newGRPCQueryService creates a new QueryService that uses gRPC to communicate
// with a multipooler instance.
func newGRPCQueryService(
	conn *grpc.ClientConn,
	poolerID string,
	logger *slog.Logger,
) queryservice.QueryService {
	_ = "STUB: not implemented"
	return *new(queryservice.QueryService)
}

// StreamExecute executes a query and streams results back via callback.
// Returns ReservedState with the authoritative reservation state from the multipooler.
func (g *grpcQueryService) StreamExecute(
	ctx context.Context,
	target *querypb.Target,
	sql string,
	options *querypb.ExecuteOptions,
	reservationOptions *querypb.ReservationOptions,
	callback func(context.Context, *sqltypes.Result) error,
) (*querypb.ReservedState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create the request

// TODO: Add caller_id when we have authentication

// Call the gRPC StreamExecute

// Stream results back via callback

// Stream completed successfully

// Extract reserved state if present

// Handle the union type payload (if present)

// Row data - convert and send to callback

// Diagnostic (notice or error) - convert to Result with notice for backwards compat

// ExecuteQuery implements queryservice.QueryService.
// This should be used sparingly only when we know the result set is small,
// otherwise StreamExecute should be used.
func (g *grpcQueryService) ExecuteQuery(ctx context.Context, target *querypb.Target, sql string, options *querypb.ExecuteOptions) (*sqltypes.Result, *querypb.ReservedState, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Create the request

// TODO: Add caller_id when we have authentication

// Call the gRPC ExecuteQuery. FromGRPC restores any *PgDiagnostic attached
// by the multipooler so the client sees the underlying PostgreSQL error.

// Convert proto result to sqltypes (preserves NULL vs empty string)

// PortalStreamExecute executes a portal (bound prepared statement) and streams results back via callback.
// Returns ReservedState containing information about the reserved connection used for this execution.
func (g *grpcQueryService) PortalStreamExecute(
	ctx context.Context,
	target *querypb.Target,
	preparedStatement *querypb.PreparedStatement,
	portal *querypb.Portal,
	options *querypb.ExecuteOptions,
	portalOptions *multipoolerservice.PortalExecuteOptions,
	callback func(context.Context, *sqltypes.Result) error,
) (*querypb.ReservedState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create the request

// TODO: Add caller_id when we have authentication

// Call the gRPC PortalStreamExecute

// Stream results back via callback

// Stream completed successfully

// Extract reserved state if present

// Handle the union type payload (if present)

// Row data - convert and send to callback

// Diagnostic (notice or error) - convert to Result with notice for backwards compat

// Describe returns metadata about a prepared statement or portal.
func (g *grpcQueryService) Describe(
	ctx context.Context,
	target *querypb.Target,
	preparedStatement *querypb.PreparedStatement,
	portal *querypb.Portal,
	options *querypb.ExecuteOptions,
) (*querypb.StatementDescription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create the request

// TODO: Add caller_id when we have authentication

// Call the gRPC Describe

// Close closes the gRPC connection.
func (g *grpcQueryService) Close() error { _ = "STUB: not implemented"; return nil }

// CopyReady initiates a COPY FROM STDIN operation and returns format information.
// The stream is stored internally and can be accessed via options.ReservedConnectionId in subsequent calls.
func (g *grpcQueryService) CopyReady(
	ctx context.Context,
	target *querypb.Target,
	copyQuery string,
	options *querypb.ExecuteOptions,
	reservationOptions *querypb.ReservationOptions,
) (int16, []int16, *querypb.ReservedState, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil, nil
}

// Start the bidirectional stream

// Ensure stream is closed if we fail before adding it to copyStreams

// Drain any pending response to allow server-side cleanup

// Send INITIATE message

// Receive READY response

// Check for ERROR response. When the multipooler rejects the COPY at
// initiation (e.g., PG raised "column does not exist") but the underlying
// reserved connection is still alive — typically because it was already
// reserved for an unrelated reason such as a transaction or temp tables —
// it sends an ERROR phase response carrying the surviving ReservedState.
// Propagate that state to the caller so the gateway keeps tracking the
// reserved connection; if no state was attached, the connection is gone
// and the gateway should clear its tracking.

// Validate READY response

// Convert columnFormats from []int32 to []int16

// Store the stream keyed by reserved connection ID

// Commit point: stream is now managed by copyStreams

// CopySendData sends a chunk of data for an active COPY operation.
func (g *grpcQueryService) CopySendData(
	ctx context.Context,
	target *querypb.Target,
	data []byte,
	options *querypb.ExecuteOptions,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Look up the stream by reserved connection ID

// CopyFinalize completes a COPY operation, sending final data and returning the result.
// Returns ReservedState with the authoritative reservation state from the multipooler.
func (g *grpcQueryService) CopyFinalize(
	ctx context.Context,
	target *querypb.Target,
	finalData []byte,
	options *querypb.ExecuteOptions,
) (*sqltypes.Result, *querypb.ReservedState, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Look up and remove the stream by reserved connection ID

// Send DONE message with final data

// Close send direction

// Receive RESULT response

// Check for ERROR response. The multipooler attaches the surviving
// ReservedState when CopyFinalize hit a PG error (e.g., constraint
// violation) but the underlying reserved connection is still alive
// because of another reason such as a transaction. Forward that state
// so the gateway keeps tracking the connection instead of clearing it.

// Validate RESULT response

// Build reserved state from response

// CopyAbort aborts a COPY operation.
// Returns ReservedState with the authoritative reservation state from the multipooler.
func (g *grpcQueryService) CopyAbort(
	ctx context.Context,
	target *querypb.Target,
	errorMsg string,
	options *querypb.ExecuteOptions,
) (*querypb.ReservedState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Look up and remove the stream by reserved connection ID

// Already cleaned up or never existed - that's okay for abort

// Send FAIL message

// Log but don't return error - we're already aborting

// Close send direction

// Try to receive response (may be ERROR) and extract reserved state

// ConcludeTransaction concludes a transaction on a reserved connection.
// Returns the result and the authoritative reservation state from the multipooler.
func (g *grpcQueryService) ConcludeTransaction(
	ctx context.Context,
	target *querypb.Target,
	options *querypb.ExecuteOptions,
	conclusion multipoolerservice.TransactionConclusion,
	releasePortalNames []string,
	releaseAllPortals bool,
) (*sqltypes.Result, *querypb.ReservedState, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Create the request

// Call the gRPC ConcludeTransaction. FromGRPC restores any *PgDiagnostic
// attached by the multipooler so the client sees the underlying PostgreSQL
// error (sqlstate + message); Wrapf adds a debug-context prefix on top
// without breaking the errors.As chain to that diagnostic.

// DiscardTempTables sends DISCARD TEMP on a reserved connection.
// Returns the result and the authoritative reservation state from the multipooler.
func (g *grpcQueryService) DiscardTempTables(
	ctx context.Context,
	target *querypb.Target,
	options *querypb.ExecuteOptions,
) (*sqltypes.Result, *querypb.ReservedState, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Create the request

// Call the gRPC DiscardTempTables. FromGRPC restores any *PgDiagnostic
// attached by the multipooler so the client sees the underlying PostgreSQL
// error; Wrapf adds a debug-context prefix without breaking the
// errors.As chain to that diagnostic.

// ReleaseReservedConnection forcefully releases a reserved connection.
func (g *grpcQueryService) ReleaseReservedConnection(
	ctx context.Context,
	target *querypb.Target,
	options *querypb.ExecuteOptions,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Clean up any stale COPY stream for this connection.

// FromGRPC restores any *PgDiagnostic attached by the multipooler so the
// client sees the underlying PostgreSQL error; Wrapf adds a debug-context
// prefix without breaking the errors.As chain to that diagnostic.

// Ensure grpcQueryService implements queryservice.QueryService
var _ queryservice.QueryService = (*grpcQueryService)(nil)
