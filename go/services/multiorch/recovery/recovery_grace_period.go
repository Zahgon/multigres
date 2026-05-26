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

package recovery

import (
	"context"
	"log/slog"
	"math/rand/v2"
	"sync"
	"time"

	"github.com/multigres/multigres/go/services/multiorch/config"
	"github.com/multigres/multigres/go/services/multiorch/recovery/types"
)

const (
	// maxAllowedJitter is the maximum jitter duration we'll allow, regardless of config.
	// This prevents misconfiguration and keeps jitter calculations simple.
	maxAllowedJitter = 1 * time.Minute
)

// gracePeriodKey uniquely identifies a grace period tracking entry.
// entityID is either a pooler ID string (for pooler-scoped problems) or a
// shard key string (for shard-scoped problems).
type gracePeriodKey struct {
	code     types.ProblemCode
	entityID string
}

// RecoveryGracePeriodTracker tracks grace periods for recovery actions.
// It implements a deadline-based model where:
// - While healthy: deadline continuously resets to now + (base + jitter)
// - Problem detected: deadline stops updating, counts down to expiry
// - Action executes only after deadline expires
//
// Thread safety: All methods are safe for concurrent use. The internal rand.Rand
// is protected by the mutex and only accessed while holding a write lock.
type RecoveryGracePeriodTracker struct {
	ctx    context.Context
	config *config.Config
	logger *slog.Logger

	mu        sync.Mutex
	deadlines map[gracePeriodKey]time.Time
	rng       *rand.Rand // Protected by mu - only accessed during Lock()
}

// RecoveryGracePeriodTrackerOption configures the deadline tracker.
type RecoveryGracePeriodTrackerOption func(*RecoveryGracePeriodTracker)

// WithRand sets a custom random generator for jitter generation.
// Useful for deterministic testing with a fixed seed.
func WithRand(rng *rand.Rand) RecoveryGracePeriodTrackerOption {
	_ = "STUB: not implemented"
	return *new(RecoveryGracePeriodTrackerOption)
}

// WithLogger sets a custom logger for the tracker.
func WithLogger(logger *slog.Logger) RecoveryGracePeriodTrackerOption {
	_ = "STUB: not implemented"
	return *new(RecoveryGracePeriodTrackerOption)
}

// NewRecoveryGracePeriodTracker creates a new deadline tracker.
// By default, uses a random seed for jitter generation and slog.Default() for logging.
func NewRecoveryGracePeriodTracker(ctx context.Context, config *config.Config, opts ...RecoveryGracePeriodTrackerOption) *RecoveryGracePeriodTracker {
	_ = "STUB: not implemented"
	return nil
}

// calculateDeadline computes a new deadline with base + jitter for the given grace period config.
// Must be called while holding dt.mu lock.
func (dt *RecoveryGracePeriodTracker) calculateDeadline(cfg types.GracePeriodConfig) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// Clamp to reasonable bounds

// Use [0, maxJitter) range (exclusive upper bound)

// Observe records the health state of a problem type for a specific entity.
// entityID is a pooler ID string for pooler-scoped problems, or a shard key
// string for shard-scoped problems.
//
// This should be called every recovery cycle for each (entity, analyzer) combination.
//
// If isHealthy is true: resets the deadline to now + (base + jitter), with fresh jitter
// If isHealthy is false: freezes the deadline (countdown continues)
//
// If the action doesn't require grace period tracking, this is a noop.
func (dt *RecoveryGracePeriodTracker) Observe(code types.ProblemCode, entityID string, action types.RecoveryAction, isHealthy bool) {
	_ = "STUB: not implemented"
	return
}

// Get grace period config from the action

// Action doesn't require grace period tracking

// Reset deadline with fresh jitter

// First time seeing this problem unhealthy - initialize deadline with base + jitter

// If unhealthy and exists, freeze (do nothing - deadline unchanged)

// ForceExpireAll immediately expires all tracked grace period deadlines.
// After this call, ShouldExecute returns true for all tracked problems regardless
// of their original deadline.
//
// Intended for use in TriggerRecoveryNow so that an explicit operator request
// can bypass the normal grace period wait and act on detected problems immediately.
func (dt *RecoveryGracePeriodTracker) ForceExpireAll() { _ = "STUB: not implemented"; return }

// zero value is before all real timestamps

// ShouldExecute checks if recovery action should execute for this problem.
// Returns true if action should execute (deadline expired or no grace period needed).
// Returns false if still within grace period window (should wait longer).
//
// This assumes Observe() has already been called for the (problem type, entity) combination.
// If the action doesn't require grace period tracking, returns true (execute immediately).
func (dt *RecoveryGracePeriodTracker) ShouldExecute(problem types.Problem) bool {
	_ = "STUB: not implemented"
	return false
}

// Get grace period config from the action

// Action doesn't require grace period tracking - execute immediately

// Problem has grace period but no deadline - this is unexpected
// Observe() should have been called before ShouldExecute()

// Check if deadline has expired

// Deadline not reached yet - log that we're deferring
