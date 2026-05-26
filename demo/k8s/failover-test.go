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

// failover-test is a tool for testing Multigres failover in Kubernetes clusters.
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

	multiadminpb "github.com/multigres/multigres/go/pb/multiadmin"
)

const (
	checkInterval = 1 * time.Second
)

// Kubernetes configuration
const (
	pgctldBin = "/usr/local/bin/pgctld"
	poolerDir = "/data"
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
	PodName   string
}

// Config holds environment configuration
type Config struct {
	MultiadminURL       string
	MultiadminGRPC      string
	KubectlContext      string
	KubernetesNamespace string
}

// PoolerStatus represents the status returned from the HTTP API
type PoolerStatus struct {
	Status struct {
		PoolerType    string `json:"pooler_type"`
		PostgresReady bool   `json:"postgres_ready"`
		PrimaryStatus *struct {
			Ready              bool `json:"ready"`
			ConnectedFollowers []struct {
				Cell string `json:"cell"`
				Name string `json:"name"`
			} `json:"connected_followers"`
		} `json:"primary_status"`
		ReplicationStatus *struct {
			LastReceiveLSN    string `json:"last_receive_lsn"`
			LastReplayLSN     string `json:"last_replay_lsn"`
			IsWALReplayPaused bool   `json:"is_wal_replay_paused"`
		} `json:"replication_status"`
	} `json:"status"`
}

// PoolersResponse represents the poolers list from HTTP API
type PoolersResponse struct {
	Poolers []struct {
		ID struct {
			Cell string `json:"cell"`
			Name string `json:"name"`
		} `json:"id"`
		Type string `json:"type"`
	} `json:"poolers"`
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1) //nolint:forbidigo // main() is allowed to call os.Exit
	}
}

var rootCmd = &cobra.Command{
	Use:   "failover-test",
	Short: "Multigres failover test script for Kubernetes clusters",
	Long: `Continuously test failover by stopping the primary pooler and waiting
for a new primary to be elected and the old primary to become a replica.

Environment variables:
  MULTIADMIN_URL          MultiAdmin API URL (default: http://localhost:18000)
  MULTIADMIN_GRPC         MultiAdmin gRPC address (default: localhost:18070)
  KUBECTL_CONTEXT         kubectl context to use (default: kind-multidemo)
  KUBERNETES_NAMESPACE    Kubernetes namespace (default: default)

Examples:
  # Run interactive mode
  go run failover-test.go

  # Run automatic mode
  go run failover-test.go --yes

  # Use custom settings
  MULTIADMIN_URL=http://localhost:8000 go run failover-test.go --yes`,
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

// Verify prerequisites

// Disable PostgreSQL monitoring on all poolers

// Start the failover loop

func loadConfig() *Config { _ = "STUB: not implemented"; return nil }

func getEnvOrDefault(key, defaultValue string) string { _ = "STUB: not implemented"; return "" }

func verifyPrerequisites(config *Config) error {
	_ = "STUB: not implemented"
	// Check kubectl connectivity
	return nil
}

// Test API connectivity
//nolint:gocritic // short-lived connectivity check

func disablePostgresMonitoring(ctx context.Context, config *Config) error {
	_ = "STUB: not implemented"
	return nil
}

func newAdminClient(addr string) (*adminClient, error) { _ = "STUB: not implemented"; return nil, nil }

type adminClient struct {
	multiadminpb.MultiAdminServiceClient
	conn *grpc.ClientConn
}

func (c *adminClient) Close() error { _ = "STUB: not implemented"; return nil }

func getPoolers(config *Config) (*PoolersResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getPoolerStatus(config *Config, cell, serviceID string) (*PoolerStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // URL constructed from trusted config and validated inputs

func findPrimary(config *Config) (*PoolerInfo, error) { _ = "STUB: not implemented"; return nil, nil }

// Check each primary to find a healthy one

func getPoolerInfo(cell, serviceID string) *PoolerInfo { _ = "STUB: not implemented"; return nil }

func stopPooler(poolerInfo *PoolerInfo, config *Config) error {
	_ = "STUB: not implemented"
	return nil
}

func waitForNewPrimary(config *Config, oldServiceID string, maxAttempts int) error {
	_ = "STUB: not implemented"
	return nil
}

func waitForReplicaHealth(config *Config, cell, serviceID string, maxAttempts int) error {
	_ = "STUB: not implemented"
	return nil
}

// Verify this replica is connected to the primary

// Find the healthy primary

// Check if this replica is in the primary's connected followers

// Verify LSN is advancing

// LSN hasn't advanced, keep waiting

func printReplicationStatus(config *Config) { _ = "STUB: not implemented"; return }

// Find the healthy primary

// Pod name is same as service ID

// Print primary info

// Print replica info

// Skip the primary

// Get replication receiver status

// Get checkpoint info

func runSQLQueryInPod(config *Config, podName, query string) string {
	_ = "STUB: not implemented"
	return ""
}

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
