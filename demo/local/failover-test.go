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

// failover-test is a tool for testing Multigres failover in local clusters.
//
// Usage:
//
//	go run failover-test.go                # Interactive mode
//	go run failover-test.go --yes          # Automatic mode
//	go run failover-test.go --yes --debug  # With debug logging
package main

import (
	"context"
	"os"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	multiadminpb "github.com/multigres/multigres/go/pb/multiadmin"
	"github.com/multigres/multigres/go/provisioner/local"
)

const (
	checkInterval = 1 * time.Second
)

// ANSI color codes
const (
	colorRed    = "\033[0;31m"
	colorGreen  = "\033[0;32m"
	colorYellow = "\033[1;33m"
	colorBlue   = "\033[0;34m"
	colorReset  = "\033[0m"
)

// PoolerInfo holds information about a pooler instance
type PoolerInfo struct {
	Cell      string
	ServiceID string
	PoolerDir string
	PgUser    string
	PgPort    int
}

// Config holds the loaded cluster configuration
type Config struct {
	Cells       map[string]local.CellServicesConfig
	AdminServer string
	RepoRoot    string
	ConfigPath  string
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1) //nolint:forbidigo // main() is allowed to call os.Exit
	}
}

var rootCmd = &cobra.Command{
	Use:   "failover-test",
	Short: "Multigres failover test script for local clusters",
	Long: `Continuously test failover by stopping the primary pooler and waiting
for a new primary to be elected and the old primary to become a replica.

Examples:
  # Run interactive mode (ask before each failover)
  ./failover-test

  # Run automatic mode (continuous failover without prompts)
  ./failover-test --yes

  # Enable debug logging
  ./failover-test --yes --debug`,
	RunE: runFailoverTest,
}

var (
	autoYes bool
	debug   bool
)

func init() {
	rootCmd.Flags().BoolVarP(&autoYes, "yes", "y", false, "Automatically proceed with failovers without confirmation")
	rootCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Enable debug logging")
}

func runFailoverTest(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gocritic // entry point for CLI tool

// Find repository root

// Load configuration

// Test connectivity

// Disable PostgreSQL monitoring on all poolers

// Start the failover loop

func findRepoRoot() (string, error) {
	_ = "STUB: not implemented"
	// Start from current directory
	return "", nil
}

// Walk up until we find go.mod or reach root

func loadConfig(repoRoot, configPath string) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Convert to LocalProvisionerConfig

func testConnectivity(config *Config) error { _ = "STUB: not implemented"; return nil }

//nolint:gocritic // short-lived connectivity check

func newAdminClient(addr string) (*adminClient, error) { _ = "STUB: not implemented"; return nil, nil }

type adminClient struct {
	multiadminpb.MultiAdminServiceClient
	conn *grpc.ClientConn
}

func (c *adminClient) Close() error { _ = "STUB: not implemented"; return nil }

func disablePostgresMonitoring(ctx context.Context, config *Config) error {
	_ = "STUB: not implemented"
	return nil
}

func getPoolers(ctx context.Context, client *adminClient) ([]*clustermetadatapb.MultiPooler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getPoolerStatus(ctx context.Context, client *adminClient, cell, serviceID string) (*multiadminpb.GetPoolerStatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func findPrimary(ctx context.Context, config *Config) (*PoolerInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check each primary to find a healthy one

func getPoolerInfo(cell, serviceID string, config *Config) *PoolerInfo {
	_ = "STUB: not implemented"
	return nil
}

func stopPooler(poolerInfo *PoolerInfo, config *Config) error {
	_ = "STUB: not implemented"
	return nil
}

func waitForNewPrimary(ctx context.Context, config *Config, oldServiceID string, maxAttempts int) error {
	_ = "STUB: not implemented"
	return nil
}

func waitForReplicaHealth(ctx context.Context, config *Config, cell, serviceID string, maxAttempts int) error {
	_ = "STUB: not implemented"
	return nil
}

// Verify this replica is connected to the primary

// Find the healthy primary

// Check if this replica is in the primary's connected followers

// Verify LSN is advancing

// LSN hasn't advanced, keep waiting

func printReplicationStatus(ctx context.Context, config *Config) { _ = "STUB: not implemented"; return }

// Find the healthy primary

// Print primary info

// Print replica info

// Skip the primary

// Get replication receiver status

// Get checkpoint info

func runSQLQuery(poolerInfo *PoolerInfo, query string) string { _ = "STUB: not implemented"; return "" }

//nolint:gocritic // short-lived query timeout

// For queries that return multiple columns, try scanning them

func failoverLoop(ctx context.Context, config *Config) error { _ = "STUB: not implemented"; return nil }

// Find current primary

// Ask user for confirmation (unless --yes flag is used)

// Stop the primary

// Wait for new primary

// Let the system restart postgres organically

// Wait for the old primary to become a healthy replica

// Print detailed replication status

// Re-disable monitoring for the next iteration

func logInfo(msg string) { _ = "STUB: not implemented"; return }

func logSuccess(msg string) { _ = "STUB: not implemented"; return }

func logWarn(msg string) { _ = "STUB: not implemented"; return }

func logError(msg string) { _ = "STUB: not implemented"; return }
