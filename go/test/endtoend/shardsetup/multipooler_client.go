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

package shardsetup

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"

	"github.com/multigres/multigres/go/common/sqltypes"
	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	multipoolerpb "github.com/multigres/multigres/go/pb/multipoolerservice"
)

// MultiPoolerTestClient wraps the gRPC client for testing
type MultiPoolerTestClient struct {
	conn   *grpc.ClientConn
	client multipoolerpb.MultiPoolerServiceClient
	addr   string
}

// NewMultiPoolerTestClient creates a new test client for multipooler
func NewMultiPoolerTestClient(addr string) (*MultiPoolerTestClient, error) {
	_ = "STUB: not implemented"
	// Validate address format
	return nil, nil
}

// ExecuteQuery executes a SQL query via the multipooler gRPC service.
// Returns sqltypes.Result with properly decoded column values.
func (c *MultiPoolerTestClient) ExecuteQuery(ctx context.Context, query string, maxRows uint64) (*sqltypes.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Convert proto QueryResult to sqltypes.Result for proper value decoding

// Close closes the gRPC connection
func (c *MultiPoolerTestClient) Close() error { _ = "STUB: not implemented"; return nil }

// Address returns the address this client is connected to
func (c *MultiPoolerTestClient) Address() string {
	_ = "STUB: not implemented"

	// IsLeader checks if the multipooler is the consensus leader by calling the Status RPC.
	// Returns true if the pooler is PRIMARY (leader), false otherwise.
	return ""
}

func IsLeader(addr string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// WaitForPoolerTypeAssigned waits for the pooler type to be assigned (either PRIMARY or REPLICA, not UNKNOWN).
// This is useful when the test needs to call RPCs that require the pooler type to be known.
func WaitForPoolerTypeAssigned(t *testing.T, addr string, timeout time.Duration) (clustermetadatapb.PoolerType, error) {
	_ = "STUB: not implemented"
	return *new(clustermetadatapb.PoolerType), nil
}

// Test helper functions

// TestBasicSelect tests a basic SELECT query
func TestBasicSelect(t *testing.T, client *MultiPoolerTestClient) {
	_ = "STUB: not implemented"
	return
}

// Verify structure

// TestCreateTable tests creating a table
func TestCreateTable(t *testing.T, client *MultiPoolerTestClient, tableName string) {
	_ = "STUB: not implemented"
	return
}

// CREATE TABLE is a modification query, so no rows returned

// TestInsertData tests inserting data into a table
func TestInsertData(t *testing.T, client *MultiPoolerTestClient, tableName string, testData []map[string]any) {
	_ = "STUB: not implemented"
	return
}

// INSERT is a modification query

// TestSelectData tests selecting data from a table
func TestSelectData(t *testing.T, client *MultiPoolerTestClient, tableName string, expectedRowCount int) {
	_ = "STUB: not implemented"
	return
}

// Verify structure

// Verify each row has the right number of values

// TestQueryLimits tests the max_rows parameter
func TestQueryLimits(t *testing.T, client *MultiPoolerTestClient, tableName string) {
	_ = "STUB: not implemented"
	return
}

// Test with limit

// TestUpdateData tests updating data in a table
func TestUpdateData(t *testing.T, client *MultiPoolerTestClient, tableName string) {
	_ = "STUB: not implemented"
	return
}

// UPDATE is a modification query

// TestDeleteData tests deleting data from a table
func TestDeleteData(t *testing.T, client *MultiPoolerTestClient, tableName string) {
	_ = "STUB: not implemented"
	return
}

// DELETE is a modification query

// TestDropTable tests dropping a table
func TestDropTable(t *testing.T, client *MultiPoolerTestClient, tableName string) {
	_ = "STUB: not implemented"
	return
}

// DROP TABLE is a modification query

// TestDataTypes tests various PostgreSQL data types
func TestDataTypes(t *testing.T, client *MultiPoolerTestClient) { _ = "STUB: not implemented"; return }

// UNKNOWN (NULL) column types are automatically coerced to TEXT. See:
// https://github.com/postgres/postgres/blob/e849bd551c323a384f2b14d20a1b7bfaa6127ed7/src/backend/parser/parse_coerce.c#L1441

// TestMultigresSchemaExists verifies that the multigres schema exists
func TestMultigresSchemaExists(t *testing.T, client *MultiPoolerTestClient) {
	_ = "STUB: not implemented"
	return
}

// TestHeartbeatTableExists verifies that the heartbeat table exists with expected columns
func TestHeartbeatTableExists(t *testing.T, client *MultiPoolerTestClient) {
	_ = "STUB: not implemented"

	// Check that the table exists
	return
}

// Check the columns

// Verify column details - check that we have the expected columns

// Verify primary key constraint exists on shard_id

// TestPrimaryDetection verifies that pg_is_in_recovery() can detect primary vs standby
func TestPrimaryDetection(t *testing.T, client *MultiPoolerTestClient) {
	_ = "STUB: not implemented"

	// Query pg_is_in_recovery() to check if connected to primary or standby
	return
}

// Note: PostgreSQL wire protocol returns boolean as 't' or 'f' in text format
// inRecovery is "f" for primary, "t" for standby/replica

// printServiceLogs prints the last N lines of a service's log file for debugging.
// If lines is 0, prints all lines.
// serviceName is the name of the service (e.g., "multiorch", "multipooler").
// database is the database name (e.g., "postgres") or empty for global services.
// serviceID is the service identifier used in the log filename (e.g., "zone1-multiorch").
// If serviceID is empty, prints all log files for the service.
func printServiceLogs(t *testing.T, configDir, serviceName, database, serviceID string, lines int) {
	_ = "STUB: not implemented"
	return
}

// If serviceID is empty, print all log files in the directory

// WaitForBootstrap waits for multiorch to bootstrap the cluster by polling until
// the multigres schema exists. Returns an error if bootstrap doesn't complete within timeout.
// This function handles the case where PostgreSQL hasn't been initialized yet by retrying
// until the database becomes available.
// If configDir and database are provided, prints service logs on failure to help debug CI issues.
func WaitForBootstrap(t *testing.T, addr string, timeout time.Duration, configDir, database string) error {
	_ = "STUB: not implemented"

	// Create gRPC connection directly without testing query
	// (PostgreSQL may not be initialized yet during bootstrap)
	return nil
}

// Log progress with error details

// Print logs to help debug CI failures
