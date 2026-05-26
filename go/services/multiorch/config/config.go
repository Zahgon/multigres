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

package config

import (
	"time"

	"github.com/spf13/pflag"

	"github.com/multigres/multigres/go/tools/viperutil"
)

// WatchTarget represents a target to watch in the format:
// - "database" - watch entire database
// - "database/tablegroup" - watch specific tablegroup
// - "database/tablegroup/shard" - watch specific shard
type WatchTarget struct {
	Database   string
	TableGroup string // empty if watching entire database
	Shard      string // empty if watching database or tablegroup level
}

// String returns the string representation of the target.
func (t WatchTarget) String() string { _ = "STUB: not implemented"; return "" }

// MatchesDatabase returns true if this target watches the given database.
func (t WatchTarget) MatchesDatabase(db string) bool { _ = "STUB: not implemented"; return false }

// MatchesTableGroup returns true if this target watches the given database/tablegroup.
// Returns true if watching entire database or specific tablegroup.
func (t WatchTarget) MatchesTableGroup(db, tablegroup string) bool {
	_ = "STUB: not implemented"
	return false
}

// Watching entire database matches all tablegroups

// MatchesShard returns true if this target watches the given database/tablegroup/shard.
// Returns true if watching entire database, entire tablegroup, or specific shard.
func (t WatchTarget) MatchesShard(db, tablegroup, shard string) bool {
	_ = "STUB: not implemented"
	return false
}

// Watching entire database matches all shards

// Watching entire tablegroup matches all shards

// ParseShardWatchTarget parses a shard watch target string.
// Format: "database" or "database/tablegroup" or "database/tablegroup/shard"
//
// Validation rules:
// - Database is always required
// - If shard is provided, tablegroup must be provided
// - Empty parts are not allowed
func ParseShardWatchTarget(s string) (WatchTarget, error) {
	_ = "STUB: not implemented"
	return *new(WatchTarget), nil
}

// Database is always required

// Validate and set tablegroup if provided

// Validate and set shard if provided

// ParseShardWatchTargets parses multiple shard watch target strings.
func ParseShardWatchTargets(targets []string) ([]WatchTarget, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Config encapsulates all multiorch configuration.
// This is passed to the recovery engine and other components.
type Config struct {
	cell                               viperutil.Value[string]
	serviceID                          viperutil.Value[string]
	shardWatchTargets                  viperutil.Value[[]string]
	bookkeepingInterval                viperutil.Value[time.Duration]
	poolerHealthCheckInterval          viperutil.Value[time.Duration]
	healthCheckWorkers                 viperutil.Value[int]
	recoveryCycleInterval              viperutil.Value[time.Duration]
	leaderFailoverGracePeriodBase      viperutil.Value[time.Duration]
	leaderFailoverGracePeriodMaxJitter viperutil.Value[time.Duration]
	verifyReplicationTimeout           viperutil.Value[time.Duration]
	leaderPostgresResponseThreshold    viperutil.Value[time.Duration]
	useNewConsensusFlow                viperutil.Value[bool]
}

// Constants
const (
	// HealthCheckQueueCapacity is the maximum number of poolers that can be queued
	// for health checking before blocking.
	HealthCheckQueueCapacity = 100000
)

// NewConfig creates a new Config with all viperutil values configured.
func NewConfig(reg *viperutil.Registry) *Config { _ = "STUB: not implemented"; return nil }

// Getter methods

func (c *Config) GetCell() string { _ = "STUB: not implemented"; return "" }

func (c *Config) GetServiceID() string { _ = "STUB: not implemented"; return "" }

func (c *Config) GetShardWatchTargets() []string { _ = "STUB: not implemented"; return nil }

func (c *Config) GetBookkeepingInterval() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *Config) GetPoolerHealthCheckInterval() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *Config) GetHealthCheckWorkers() int { _ = "STUB: not implemented"; return 0 }

func (c *Config) GetRecoveryCycleInterval() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *Config) GetLeaderFailoverGracePeriodBase() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *Config) GetLeaderFailoverGracePeriodMaxJitter() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *Config) GetVerifyReplicationTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *Config) GetLeaderPostgresResponseThreshold() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *Config) GetUseNewConsensusFlow() bool { _ = "STUB: not implemented"; return false }

// Defaults for flags (used in RegisterFlags)

func (c *Config) DefaultCell() string { _ = "STUB: not implemented"; return "" }

func (c *Config) DefaultServiceID() string { _ = "STUB: not implemented"; return "" }

func (c *Config) DefaultShardWatchTargets() []string { _ = "STUB: not implemented"; return nil }

func (c *Config) DefaultBookkeepingInterval() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *Config) DefaultPoolerHealthCheckInterval() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *Config) DefaultHealthCheckWorkers() int { _ = "STUB: not implemented"; return 0 }

func (c *Config) DefaultRecoveryCycleInterval() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *Config) DefaultLeaderFailoverGracePeriodBase() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *Config) DefaultLeaderFailoverGracePeriodMaxJitter() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *Config) DefaultVerifyReplicationTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *Config) DefaultLeaderPostgresResponseThreshold() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *Config) DefaultUseNewConsensusFlow() bool { _ = "STUB: not implemented"; return false }

// RegisterFlags registers the config flags with pflag.
func (c *Config) RegisterFlags(fs *pflag.FlagSet) { _ = "STUB: not implemented"; return }

// Test helper functions

// NewTestConfig creates a Config for testing with optional custom values.
// Sets safe defaults for grace period (0 base, 0 jitter) so tests execute immediately.
func NewTestConfig(opts ...func(*Config)) *Config { _ = "STUB: not implemented"; return nil }

// Set safe defaults for tests - no grace period by default

// WithCell sets the cell value for testing.
func WithCell(cell string) func(*Config) { _ = "STUB: not implemented"; return nil }

// WithUseNewConsensusFlow toggles the Recruit/Propose/SetTermPrimary consensus flow.
func WithUseNewConsensusFlow(enabled bool) func(*Config) { _ = "STUB: not implemented"; return nil }

// WithBookkeepingInterval sets the bookkeeping interval for testing.
func WithBookkeepingInterval(d time.Duration) func(*Config) { _ = "STUB: not implemented"; return nil }

// WithPoolerHealthCheckInterval sets the pooler health check interval for testing.
func WithPoolerHealthCheckInterval(d time.Duration) func(*Config) {
	_ = "STUB: not implemented"
	return nil
}

// WithHealthCheckWorkers sets the number of health check workers for testing.
func WithHealthCheckWorkers(n int) func(*Config) { _ = "STUB: not implemented"; return nil }

// WithRecoveryCycleInterval sets the recovery cycle interval for testing.
func WithRecoveryCycleInterval(d time.Duration) func(*Config) {
	_ = "STUB: not implemented"
	return nil
}

// WithLeaderFailoverGracePeriodBase sets the leader failover grace period base for testing.
func WithLeaderFailoverGracePeriodBase(d time.Duration) func(*Config) {
	_ = "STUB: not implemented"
	return nil
}

// WithLeaderFailoverGracePeriodMaxJitter sets the leader failover grace period max jitter for testing.
func WithLeaderFailoverGracePeriodMaxJitter(d time.Duration) func(*Config) {
	_ = "STUB: not implemented"
	return nil
}

// WithLeaderPostgresResponseThreshold sets the primary postgres responded threshold for testing.
func WithLeaderPostgresResponseThreshold(d time.Duration) func(*Config) {
	_ = "STUB: not implemented"
	return nil
}
