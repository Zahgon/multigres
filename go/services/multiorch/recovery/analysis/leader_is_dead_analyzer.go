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

package analysis

import (
	"github.com/multigres/multigres/go/services/multiorch/recovery/types"
)

// LeaderIsDeadAnalyzer detects when a leader exists in the shard but is unhealthy/unreachable.
// It operates at the shard level: if any initialized follower observes the leader as dead,
// one shard-scoped problem is emitted.
type LeaderIsDeadAnalyzer struct {
	factory *RecoveryActionFactory
}

func (a *LeaderIsDeadAnalyzer) Name() types.CheckName {
	_ = "STUB: not implemented"
	return *new(types.CheckName)
}

func (a *LeaderIsDeadAnalyzer) ProblemCode() types.ProblemCode {
	_ = "STUB: not implemented"
	return *new(types.ProblemCode)
}

func (a *LeaderIsDeadAnalyzer) RecoveryAction() types.RecoveryAction {
	_ = "STUB: not implemented"
	return *new(types.RecoveryAction)
}

func (a *LeaderIsDeadAnalyzer) Analyze(sa *ShardAnalysis) ([]types.Problem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// No leader known — nothing to declare dead.

// Leader is fully reachable — no problem.

// No initialized replica to confirm the leader is dead — skip to avoid false positives
// when the shard has no postgres standby that has joined the cluster yet.

// Suppress failover during a known pg_promote() window. The multipooler explicitly
// signals promotion is in progress via PromotingPrimaryID. The conditions:
//   - PromotingPrimaryID != nil: multipooler has flagged pg_promote() is running
//   - LeaderPoolerReachable: stream is live, so the flag is current (not stale)
//   - LeaderPostgresRunning: postgres process is still alive
// If postgres crashes during promotion, LeaderPostgresRunning=false and we fall through.
// If the multipooler crashes, LeaderPoolerReachable=false and we fall through.

// At this point, LeaderReachable is false. This can happen in three cases:
//
// 1. Leader pooler is unreachable (e.g. pooler process crashed).
//    Postgres may still be running; followers can still receive WAL.
//
// 2. Leader pooler is reachable and Postgres process is alive yet
//    unresponsive, pg_isready fails but the process exists (e.g. SIGSTOP or
//    overloaded). Followers remain connected until TCP keepalive times out.
//
// 3. Leader pooler is reachable but Postgres process is dead: this means
//    pg_isready can be assumed to fail and the process is gone (e.g.
//    SIGKILL). Followers may still appear connected for ~30s via TCP
//    keepalive even though Postgres is dead.
//
// For cases 1 and 2, we check if ALL followers are still connected to the
// leader's postgres. If they are, postgres is still running (or recovering)
// and we suppress failover, but only if the leader's postgres responded
// recently enough. This prevents suppressing indefinitely when followers are
// observing stale connections while postgres is unresponsive.
//
// For case 3, we must NOT suppress: the pooler reports the process is dead,
// so followers' apparent connections are stale (TCP keepalive hasn't fired
// yet). Suppressing would delay failover by up to the TCP keepalive
// interval (~30s).

// Cases 1 and 2: followers are connected and the leader pooler is down
// OR the postgres process is alive (but possibly unresponsive). Suppress
// failover if postgres responded recently (within threshold).

// Cases 1 and 2: postgres timestamp expired or unset — suppression window closed, allowing failover.

// Case 3: pooler is reachable but reports postgres process is dead.
// This happens after SIGKILL: the process is gone but followers still show as connected.

// Leader is dead — emit one shard-level problem.
