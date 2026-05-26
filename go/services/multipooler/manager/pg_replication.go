// Copyright 2025 Supabase, Inc.
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

package manager

import (
	"context"
	"log/slog"

	"github.com/multigres/multigres/go/services/multipooler/executor"

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	multipoolermanagerdatapb "github.com/multigres/multigres/go/pb/multipoolermanagerdata"
)

// ============================================================================
// PostgreSQL Replication Operations
//
// This file contains methods and functions for querying and configuring
// PostgreSQL replication settings. These are low-level operations that
// directly interact with the database.
//
// For high-level orchestration logic (promotion, demotion, etc.), see
// manager.go and rpc_manager.go.
// ============================================================================

// ----------------------------------------------------------------------------
// Application Name Helpers
// ----------------------------------------------------------------------------

// maxApplicationNameLength is the maximum length of a PostgreSQL application_name (NAMEDATALEN - 1).
const maxApplicationNameLength = 63

// poolerID is a validated MULTIPOOLER identity for use in PostgreSQL replication.
// Construct via newPoolerID(). Holds both the canonical cluster ID and its derived
// PostgreSQL application_name. The appName field is the only serialization format
// used internally; other formats (e.g., JSON, gRPC ID) are accessible via the id field.
type poolerID struct {
	id      *clustermetadatapb.ID
	appName string
}

// newPoolerID generates the poolerID for a multipooler from its ID.
// Format: {cell}_{name}
// This is used consistently for:
// - SetPrimaryConnInfo: standby's application_name when connecting to primary
// - ConfigureSynchronousReplication: standby names in synchronous_standby_names
//
// On validation failure an approximate poolerID is returned alongside the error.
// The approximate appName is "{cell}_{name}" with missing fields replaced by
// "<unknown>". For overlong names the full (invalid) string is returned as-is.
// Callers that require a strictly valid appName must check the error. Callers that
// can tolerate an approximate name (e.g. for logging or informational responses)
// may use the returned value regardless.
func newPoolerID(id *clustermetadatapb.ID) (poolerID, error) {
	_ = "STUB: not implemented"
	return *new(poolerID), nil
}

// Underscores are not allowed in Cell or Name because they are used as delimiters
// in the application_name format (cell_name). Allowing underscores would break parsing.

// poolerIDsToAppNames converts a slice of poolerID to their application name strings for use in APIs
// that accept []string (e.g., history records).
func poolerIDsToAppNames(ids []poolerID) []string { _ = "STUB: not implemented"; return nil }

// formatStandbyList formats a list of pooler IDs as a comma-separated list of quoted application names.
func formatStandbyList(ids []poolerID) string { _ = "STUB: not implemented"; return "" }

// toPoolerIDs converts a slice of IDs to their poolerID representations.
// If any ID fails strict validation the corresponding poolerID contains an
// approximate appName and the first error is returned alongside the full slice.
// Callers that require strict validity must check the error. Callers that can
// tolerate approximate names (e.g. for logging or informational responses) may
// use the returned slice regardless.
func toPoolerIDs(ids []*clustermetadatapb.ID) ([]poolerID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// poolerIDFromAppName parses a PostgreSQL application_name (format: "cell_name") into
// a poolerID. This is the inverse of newPoolerID's appName generation.
// On parse failure an error is returned alongside a best-effort poolerID that
// preserves the raw appName in the Name field so callers can include it rather
// than silently dropping the member.
//
// TODO: once leadership_history stores serialized clustermetadata.ID values directly
// instead of application_name strings, this parsing and its error path can be removed.
func poolerIDFromAppName(appName string) (poolerID, error) {
	_ = "STUB: not implemented"
	return *new(poolerID), nil
}

// ----------------------------------------------------------------------------
// Replication Status Query Methods
// ----------------------------------------------------------------------------

// isPrimary checks if the connected database is a primary (not in recovery)
func (pm *MultiPoolerManager) isPrimary(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// isInRecovery checks if the connected database is in recovery mode (standby).
// Returns true if the database is a standby, false if it's a primary.
func (pm *MultiPoolerManager) isInRecovery(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// getPrimaryLSN gets the current WAL write location (primary only)
func (pm *MultiPoolerManager) getPrimaryLSN(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// getStandbyReplayLSN gets the last replayed WAL location (standby only)
func (pm *MultiPoolerManager) getStandbyReplayLSN(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// getTimelineID gets the current timeline ID from pg_control_checkpoint()
func (pm *MultiPoolerManager) getTimelineID(ctx context.Context) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// querySchemaExists checks if the multigres schema exists in the database
func (pm *MultiPoolerManager) querySchemaExists(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// checkLSNReached checks if the standby has replayed up to or past the target LSN
func (pm *MultiPoolerManager) checkLSNReached(ctx context.Context, targetLsn string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// sqlGetReplicationStatus is the SQL query to retrieve all relevant replication
// status fields from pg_stat_wal_receiver. This is used by
// queryReplicationStatus to get a comprehensive view of the replication state
// in one query. Note that some fields may be NULL depending on the state of the
// standby (e.g., if not in recovery or if no WAL has been received).
//
// Scalar subqueries are used for pg_stat_wal_receiver fields so that NULL is
// returned when the view is empty (e.g., on the primary or when the WAL
// receiver is not running), rather than returning zero rows.
const sqlGetReplicationStatus = `
SELECT	pg_last_wal_replay_lsn(),
		pg_last_wal_receive_lsn(),
		pg_is_wal_replay_paused(),
		pg_get_wal_replay_pause_state(),
		pg_last_xact_replay_timestamp(),
		current_setting('primary_conninfo'),
		(SELECT status FROM pg_stat_wal_receiver LIMIT 1),
		(SELECT last_msg_receipt_time FROM pg_stat_wal_receiver LIMIT 1),
		current_setting('wal_receiver_status_interval'),
		current_setting('wal_receiver_timeout')
`

// queryReplicationStatus queries PostgreSQL for all replication status fields.
// This method handles NULL values properly for LSN fields that may be NULL
// when not in recovery mode or when no WAL has been received/replayed.
func (pm *MultiPoolerManager) queryReplicationStatus(ctx context.Context) (*multipoolermanagerdatapb.StandbyReplicationStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We can use ParseDuration here since PostgreSQL interval settings are
// in a format compatible with Go durations (e.g., "10s", "500ms").

// We can use ParseDuration here since PostgreSQL interval settings are
// in a format compatible with Go durations (e.g., "10s", "500ms").

// Parse primary_conninfo into structured format

// waitForReplicationPause polls until WAL replay is paused and returns the status at that moment.
// This ensures the LSN returned represents the exact point at which replication stopped.
func (pm *MultiPoolerManager) waitForReplicationPause(ctx context.Context) (*multipoolermanagerdatapb.StandbyReplicationStatus, error) {
	_ = "STUB: not implemented"
	// Create a context with timeout for the polling loop
	return nil, nil
}

// Query all replication status fields

// Once paused, we have the exact state at the moment replication stopped

// readPrimaryConnInfo returns the current primary_conninfo setting as a raw string.
// Returns an empty string if primary_conninfo is not set.
func (pm *MultiPoolerManager) readPrimaryConnInfo(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// setPrimaryConnInfo sets the primary_conninfo connection string
func (pm *MultiPoolerManager) setPrimaryConnInfo(ctx context.Context, connInfo string) error {
	_ = "STUB: not implemented"
	return nil
}

// resetPrimaryConnInfo clears primary_conninfo and reloads PostgreSQL configuration.
// This effectively disconnects the replica from the primary.
func (pm *MultiPoolerManager) resetPrimaryConnInfo(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Clear primary_conninfo using ALTER SYSTEM (should be quick)
	return nil
}

// waitForReplayStabilize waits, best effort, for WAL replay to stop making
// observable progress. The intent is to approximate replay is idle given the WAL
// that is currently available to this standby.
//
// WARNING: This function is not perfect and has some theoretical limitations.
// See decision: 2026-02-12-wait-for-replay-stabilize-during-revoke.md for more context.
func (pm *MultiPoolerManager) waitForReplayStabilize(ctx context.Context) (*multipoolermanagerdatapb.StandbyReplicationStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// requiredStablePolls: number of consecutive polls showing the same replay_lsn
// before we declare stability. At 10ms per tick, 3 polls = 30ms of stability.

// queryReplayState returns the current replay LSN and pause state.
// Returns FAILED_PRECONDITION if the server is not in recovery (replay LSN is NULL).
func (pm *MultiPoolerManager) queryReplayState(ctx context.Context) (replayLsn string, isPaused bool, err error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

// waitForReceiverDisconnect waits for the WAL receiver to fully disconnect after clearing primary_conninfo.
// It polls pg_stat_wal_receiver to confirm the receiver has stopped.
func (pm *MultiPoolerManager) waitForReceiverDisconnect(ctx context.Context) (*multipoolermanagerdatapb.StandbyReplicationStatus, error) {
	_ = "STUB: not implemented"
	// Create a context with timeout for the polling loop
	return nil, nil
}

// Track the latest poll snapshot so a timeout has diagnostic detail without
// needing another query after the ctx has already expired.

// -1 = not yet polled

// Re-check waitCtx before issuing a query: when waitCtx expires at the
// same tick boundary, select may pick this branch and the subsequent
// pm.query would surface an opaque "pool ctx expired" instead of the
// real timeout cause.

// Pull the count, the walreceiver status, and the live primary_conninfo
// in a single query so each poll is also a diagnostic snapshot.

// If waitCtx expired between the pre-check and the Pool.Get
// ctx.Err() check, surface the cleaner timeout cause instead of
// an opaque "pool ctx expired" wrapper.

// Done when either the walreceiver slot is gone, OR it's sitting
// in WALRCV_WAITING with primary_conninfo empty. The latter is
// safe because:
//
//   - We hold the action lock, so no other in-process path can
//     write primary_conninfo during this wait.
//   - WAITING → STREAMING requires the startup process to call
//     RequestXLogStreaming, which only fires when primary_conninfo
//     is non-empty.
//   - We can't actively terminate a walreceiver from SQL —
//     pg_terminate_backend only works on regular backends, not
//     auxiliary processes like the walreceiver. The walreceiver
//     only exits when its in-flight libpq call returns, which is
//     bounded by connect_timeout (potentially tens of seconds).
//     Treating WAITING+empty as done lets us proceed without
//     waiting out that timeout.

// Get the final replication status

// pauseReplication pauses replication based on the specified mode.
// If wait is true, it waits for the pause operation to complete before returning.
// Returns the replication status after pausing (if wait is true) or nil (if wait is false).
func (pm *MultiPoolerManager) pauseReplication(ctx context.Context, mode multipoolermanagerdatapb.ReplicationPauseMode, wait bool) (*multipoolermanagerdatapb.StandbyReplicationStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Pause WAL replay on the standby

// Set tight timeout for the pause command itself (should be quick)

// Wait for WAL replay to actually be paused
// pg_wal_replay_pause() is asynchronous, so we need to wait for it to complete

// Stop the WAL receiver by clearing primary_conninfo

// Wait for receiver to fully disconnect

// IMPORTANT: Must stop receiver BEFORE pausing replay
// Reason: When replay is paused, the WAL receiver won't disconnect even if we clear primary_conninfo
// So we must clear primary_conninfo while replay is still running

// First stop receiver (while replay is still running)

// Wait for receiver to disconnect before pausing replay

// Now that receiver is disconnected, pause replay

// Wait for replay pause to complete

// resumeWALReplay resumes WAL replay on a standby server
func (pm *MultiPoolerManager) resumeWALReplay(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// reloadPostgresConfig reloads PostgreSQL configuration to apply changes made via
// ALTER SYSTEM, and waits for postmaster to finish re-reading the config files
// before returning.
//
// pg_reload_conf() returns immediately after sending SIGHUP to postmaster, well
// before any of that work has happened. We use pg_conf_load_time() — the
// timestamp of postmaster's most recent successful config load — as the
// completion signal: once a follow-up query observes it advance past the
// pre-reload value, postmaster has re-read postgresql.auto.conf and signalled
// its child processes.
//
// Caveat: this guarantees postmaster has processed the reload, not that every
// child process has. Backends (the walreceiver, individual query backends)
// each pick up SIGHUP at their own pace — typically within milliseconds, but
// not synchronously. Callers that need to observe a child's reaction (e.g.
// polling pg_stat_wal_receiver for the walreceiver to disconnect after
// clearing primary_conninfo) should still poll, but they can do so knowing
// the new config is loaded server-side rather than racing with postmaster's
// signal handler.
func reloadPostgresConfig(ctx context.Context, logger *slog.Logger, qs executor.InternalQueryService) error {
	_ = "STUB: not implemented"
	return nil
}

// Poll pg_conf_load_time() until it advances. retry.New uses "do work, then
// back off" semantics, so the backoff timer starts after the previous query
// finishes — a slow query under load doesn't cause back-to-back hammering.

// Unreachable: r.Attempts only exits via the ctx-cancelled branch above.

// reloadPostgresConfig is a convenience wrapper around the package-level
// reloadPostgresConfig helper bound to this manager's query service and logger.
func (pm *MultiPoolerManager) reloadPostgresConfig(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// validateExpectedLSN validates that the current replay LSN matches the expected LSN
func (pm *MultiPoolerManager) validateExpectedLSN(ctx context.Context, expectedLSN string) error {
	_ = "STUB: not implemented"
	return nil
}

// No validation requested

// Best practice: WAL replay should be paused before promotion
// The coordinator should have called StopReplication during Discovery stage

// Note: We don't fail here as this is a soft check, but it indicates
// a potential issue in the consensus flow

// ----------------------------------------------------------------------------
// Synchronous Replication Configuration
// ----------------------------------------------------------------------------

// setSynchronousCommit sets the PostgreSQL synchronous_commit level
func (pm *MultiPoolerManager) setSynchronousCommit(ctx context.Context, synchronousCommit multipoolermanagerdatapb.SynchronousCommitLevel) error {
	_ = "STUB: not implemented"
	// Convert enum to PostgreSQL string value
	return nil
}

// buildSynchronousStandbyNamesValue constructs the synchronous_standby_names value string
// This produces values like: FIRST 1 ("standby-1", "standby-2") or ANY 1 ("standby-1", "standby-2")
func buildSynchronousStandbyNamesValue(method multipoolermanagerdatapb.SynchronousMethod, numSync int32, names []poolerID) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// applySynchronousStandbyNames applies the synchronous_standby_names setting to PostgreSQL
func (pm *MultiPoolerManager) applySynchronousStandbyNames(ctx context.Context, value string) error {
	_ = "STUB: not implemented"
	return nil
}

// ALTER SYSTEM SET doesn't support parameterized queries, so we use string formatting

// setSynchronousStandbyNames builds and sets the PostgreSQL synchronous_standby_names configuration
// Format: https://www.postgresql.org/docs/current/runtime-config-replication.html#GUC-SYNCHRONOUS-STANDBY-NAMES
// Examples:
//
//	FIRST 2 (standby1, standby2, standby3)
//	ANY 1 (standby1, standby2)
//
// Note: Use '*' to match all connected standbys, or specify explicit standby application_name values
// Application names are generated from multipooler IDs using the shared newPoolerID helper
func (pm *MultiPoolerManager) setSynchronousStandbyNames(ctx context.Context, synchronousMethod multipoolermanagerdatapb.SynchronousMethod, numSync int32, names []poolerID) error {
	_ = "STUB: not implemented"
	// If standby list is empty, clear synchronous_standby_names
	return nil
}

// If numSync was not provided, default to 1

// Build the synchronous_standby_names value using the shared helper

// Apply the setting

// getSynchronousReplicationConfig retrieves and parses the current synchronous replication configuration
func (pm *MultiPoolerManager) getSynchronousReplicationConfig(ctx context.Context) (*multipoolermanagerdatapb.SynchronousReplicationConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Query synchronous_standby_names

// Only parse standby names if not empty

// Query synchronous_commit

// Map string to enum

// clearSyncReplicationForDemotion clears synchronous replication settings at the start of demotion.
//
// When a stale primary comes back online after failover:
// 1. It still has synchronous_standby_names configured
// 2. No standbys are connected (they're all connected to the new primary)
// 3. Any writes (like heartbeat) block indefinitely waiting for sync acknowledgment
// 4. This blocks the demote flow and causes timeout
//
// ALTER SYSTEM writes to postgresql.auto.conf, not to WAL, so it doesn't need sync
// replication acknowledgment and won't block even with no standbys connected.
func (pm *MultiPoolerManager) clearSyncReplicationForDemotion(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Use a short timeout - if this hangs, the demote will fail anyway

// ALTER SYSTEM writes to postgresql.auto.conf (not WAL), so it doesn't require
// sync replication acknowledgment and won't block.

// resetSynchronousReplication clears the synchronous standby list
// This should be called after the server is read-only to safely clear settings
func (pm *MultiPoolerManager) resetSynchronousReplication(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Clear synchronous_standby_names to remove all standbys

// syncReplicationConfigMatches checks if the current sync replication config matches the requested config
func (pm *MultiPoolerManager) syncReplicationConfigMatches(current *multipoolermanagerdatapb.SynchronousReplicationConfiguration, requested *multipoolermanagerdatapb.ConfigureSynchronousReplicationRequest) bool {
	_ = "STUB: not implemented"
	// Check synchronous commit level
	return false
}

// Check synchronous method

// Check num_sync

// Check standby IDs (must match exactly 1:1, so sort and compare)

// Sort both lists by cell_name for comparison

// Compare sorted lists element by element

// ----------------------------------------------------------------------------
// Validation Helpers
// ----------------------------------------------------------------------------
// validateStandbyIDs validates that the list is non-empty and converts each ID to its poolerID.
func validateStandbyIDs(standbyIDs []*clustermetadatapb.ID) ([]poolerID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// validateSyncReplicationParams validates the parameters for ConfigureSynchronousReplication
func validateSyncReplicationParams(numSync int32, standbyIDs []*clustermetadatapb.ID) ([]poolerID, error) {
	_ = "STUB: not implemented"
	// Validate numSync is non-negative
	return nil, nil
}

// If standbyIDs are provided, validate them

// Validate that numSync doesn't exceed the number of standbys (PostgreSQL requirement)
// Note: numSync=0 is allowed and will be defaulted to 1 in setSynchronousStandbyNames

// Validate each standby ID

// ----------------------------------------------------------------------------
// standbyUpdateOperationName maps a CohortUpdateOperation enum to a short string for logging/history.
func standbyUpdateOperationName(op multipoolermanagerdatapb.CohortUpdateOperation) string {
	_ = "STUB: not implemented"
	return ""
}

// Standby List Operations
// ----------------------------------------------------------------------------

// applyAddOperation adds new standbys to the standby list (idempotent)
func applyAddOperation(currentStandbys, newStandbys []poolerID) []poolerID {
	_ = "STUB: not implemented"
	return nil
}

// applyRemoveOperation removes standby names from the standby list (idempotent)
func applyRemoveOperation(currentStandbys, standbysToRemove []poolerID) []poolerID {
	_ = "STUB: not implemented"
	return nil
}

// poolerIDSetEqual returns true if a and b contain the same set of pooler IDs
// (order-independent comparison using appName as the key).
func poolerIDSetEqual(a, b []poolerID) bool { _ = "STUB: not implemented"; return false }

// ----------------------------------------------------------------------------
// Primary-side Replication Queries
// ----------------------------------------------------------------------------

// getConnectedFollowerIDs queries pg_stat_replication for connected followers and returns their IDs
func (pm *MultiPoolerManager) getConnectedFollowerIDs(ctx context.Context) ([]*clustermetadatapb.ID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Parse application_name back to cluster ID
