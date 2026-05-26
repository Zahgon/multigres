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

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	consensuspb "github.com/multigres/multigres/go/pb/consensus"
	multipoolermanagerpb "github.com/multigres/multigres/go/pb/multipoolermanager"
)

// GetCurrentTerm returns the current consensus term from a node.
// Use this instead of hardcoded term values for test isolation.
func GetCurrentTerm(ctx context.Context, client consensuspb.MultiPoolerConsensusClient) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// MustGetCurrentTerm returns the current term or fails the test.
func MustGetCurrentTerm(t *testing.T, ctx context.Context, client consensuspb.MultiPoolerConsensusClient) int64 {
	_ = "STUB: not implemented"
	return 0
}

// ValidatePoolerType checks that the pooler type in topology matches the expected value.
// Follows the pattern from multipooler/setup_test.go:validatePoolerType.
func ValidatePoolerType(ctx context.Context, client multipoolermanagerpb.MultiPoolerManagerClient, expectedType clustermetadatapb.PoolerType, nodeName string) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateTerm checks that the consensus term matches the expected value.
// Follows the pattern from multipooler/setup_test.go:validateTerm.
func ValidateTerm(ctx context.Context, client consensuspb.MultiPoolerConsensusClient, expectedTerm int64, nodeName string) error {
	_ = "STUB: not implemented"
	return nil
}

// RestorePrimaryAfterDemotion restores the original primary to primary state after it was demoted.
// Uses Force=true to bypass term validation for simplicity in test cleanup.
func RestorePrimaryAfterDemotion(ctx context.Context, t *testing.T, client *MultipoolerClient) error {
	_ = "STUB: not implemented"

	// After EmergencyDemote restarts postgres as standby, the monitor loop reconciles
	// the stored pooler type from PRIMARY to REPLICA asynchronously. StopReplication
	// requires type=REPLICA, so wait for convergence before calling it.
	return nil
}

// Stop replication on primary

// Get current LSN from manager status

// Force promote primary back - term value doesn't matter when Force=true
// (Force bypasses term validation)

// Ignored when Force=true

// SaveGUCs queries multiple GUC values and saves them to a map.
// Returns a map of gucName -> value. Empty values are preserved.
func SaveGUCs(ctx context.Context, client *MultiPoolerTestClient, gucNames []string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// ReloadConfig calls pg_reload_conf() and waits for the reload to complete
// using pg_conf_load_time() as an event-based completion signal.
//
// pg_reload_conf() sends SIGHUP and returns immediately, before postgres has
// processed the signal. This function waits until pg_conf_load_time() advances
// past the pre-reload value, which happens atomically when postgres finishes
// processing the SIGHUP — at which point all GUC values are guaranteed to
// reflect the latest postgresql.auto.conf.
//
// ctx is used only for the pg_reload_conf() call. The reload-completion wait
// uses its own internal context so that a short caller deadline does not cut
// off the wait on a loaded system.
func ReloadConfig(ctx context.Context, t *testing.T, client *MultiPoolerTestClient, instanceName string) {
	_ = "STUB: not implemented"
	return
}

// pg_conf_load_tim() should not normally fail, but if it does, log the
// error so that we can debug the issue.

// Use a fresh context for polling so a short caller ctx does not cut off the wait.

// pg_conf_load_tim() should not normally fail, but if it does, stop
// polling and report it, instead of polling until timeout.

// RestoreGUCs restores GUC values from a saved map using ALTER SYSTEM, then
// calls ReloadConfig to apply the changes and wait for the reload to complete.
// Empty values are treated as RESET (restore to default).
func RestoreGUCs(ctx context.Context, t *testing.T, client *MultiPoolerTestClient, savedGucs map[string]string, instanceName string) {
	_ = "STUB: not implemented"
	return
}

// ValidateGUCValue queries a GUC and returns an error if it doesn't match the expected value.
// Follows the pattern from multipooler/setup_test.go:validateGUCValue.
func ValidateGUCValue(ctx context.Context, client *MultiPoolerTestClient, gucName, expected, instanceName string) error {
	_ = "STUB: not implemented"
	return nil
}
