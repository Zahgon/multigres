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

package manager

import (
	"context"
	"time"
)

// pgctld.Stop is escalated through these modes in order. Each mode gets its
// own bounded timeout; the total fits inside servenv's --onterm-timeout (20s
// default). Operators who need a longer shutdown can raise --onterm-timeout;
// if both modes fail we log and return and servenv's onterm-timeout
// enforcement eventually forces the process to move on.
//
// "smart" mode is intentionally absent. Smart waits for every postgres
// client connection to disconnect, and our own connection pool keeps
// backends open until process exit — so smart would always time out
// waiting for them. fast sends SIGTERM to postgres which terminates those
// backends directly, achieving the same end state with no wasted wait.
var pgctldStopModes = []struct {
	name    string
	timeout time.Duration
}{
	{"fast", 10 * time.Second},
	{"immediate", 5 * time.Second},
}

// GracefulShutdown publishes REQUESTING_DEMOTION on the health stream (for a
// current leader) and then stops Postgres. Registered as a servenv OnTermSync
// hook so it runs on SIGTERM bounded by --onterm-timeout.
//
// The announcement is sequenced before pgctld.Stop because reading the
// primary term requires querying postgres for the current rule position;
// doing it post-stop would silently no-op and force the coordinator to wait
// for stream EOF + LeaderIsDead grace period instead of firing
// LeaderResignedAnalyzer immediately.
//
// The topology Type=DRAINED transition happens after this returns via the
// existing OnClose -> mp.Shutdown -> tr.Unregister chain registered in
// services/multipooler/init.go.
//
// The action lock is held for the whole sequence: pgctld.Stop is gated behind
// the protectedPgctldClient action-lock check, and announcing the resignation
// involves reading consensus state and writing resignedLeaderAtTerm under the
// same lock — holding it across both serialises against any concurrent
// consensus operation (Promote, Demote, Recruit, BeginTerm REVOKE).
func (pm *MultiPoolerManager) GracefulShutdown(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

// Transition to NOT_SERVING so the gateway sees a clean rejection for new
// queries while in-flight transactions are allowed to complete (bounded by
// --connpool-drain-grace-period). Best-effort: a failure here is logged but
// doesn't block the rest of shutdown.

// If we are the leader, announce REQUESTING_DEMOTION before stopping
// postgres. primaryTermLocked reads the current rule position from
// postgres, so the read must happen while postgres is still alive;
// running it post-stop would fail and the coordinator would have to wait
// for stream EOF + LeaderIsDead grace period instead. No-op for
// non-leaders and partially-initialized managers (consensus not wired).
// Mirrors the EmergencyDemote pattern in rpc_manager.go.

// Signal long-lived subscribers (health-stream gRPC handlers) that the
// manager is shutting down. Their cleanup goroutines close subscriber
// channels, which makes the gRPC handlers return Unavailable and unblocks
// servenv's parallel grpcServer.GracefulStop hook. Without this, the
// handlers sit in `select { <-pollTicker.C }` forever and GracefulStop
// only completes when servenv's --onterm-timeout fires.
//
// Nil guard: some unit tests construct MultiPoolerManager via struct
// literal without going through NewMultiPoolerManager.

// stopPostgresLocked stops postgres via pgctld using fast → immediate.
// Caller must hold the action lock.
func (pm *MultiPoolerManager) stopPostgresLocked(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}
