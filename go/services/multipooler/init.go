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

// Package multipooler provides multipooler functionality.
package multipooler

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/multigres/multigres/go/common/servenv"
	"github.com/multigres/multigres/go/common/servenv/toporeg"
	"github.com/multigres/multigres/go/common/topoclient"
	"github.com/multigres/multigres/go/services/multipooler/connpoolmanager"
	"github.com/multigres/multigres/go/tools/telemetry"
	"github.com/multigres/multigres/go/tools/viperutil"
)

// MultiPooler represents the main multipooler instance with all configuration and state
type MultiPooler struct {
	pgctldAddr          viperutil.Value[string]
	cell                viperutil.Value[string]
	database            viperutil.Value[string]
	tableGroup          viperutil.Value[string]
	shard               viperutil.Value[string]
	serviceID           viperutil.Value[string]
	socketFilePath      viperutil.Value[string]
	poolerDir           viperutil.Value[string]
	pgPort              viperutil.Value[int]
	heartbeatIntervalMs viperutil.Value[int]
	// pgBackRest TLS certificate paths for client authentication to primary's pgBackRest server
	pgBackRestCertFile viperutil.Value[string]
	pgBackRestKeyFile  viperutil.Value[string]
	pgBackRestCAFile   viperutil.Value[string]
	pgBackRestPort     viperutil.Value[int]
	// vpidStampEnabled controls stamping of multigres_vpid:<id> on PostgreSQL
	// backends so lock-detection can map a multigateway virtual PID to the
	// real backend PID via pg_stat_activity.application_name.
	vpidStampEnabled viperutil.Value[bool]
	// GrpcServer is the grpc server
	grpcServer *servenv.GrpcServer
	// Senv is the serving environment
	senv *servenv.ServEnv
	// TopoConfig holds topology configuration
	topoConfig *topoclient.TopoConfig
	telemetry  *telemetry.Telemetry
	// connPoolConfig holds connection pool configuration (manager created inside MultiPoolerManager)
	connPoolConfig *connpoolmanager.Config

	ts           topoclient.Store
	tr           *toporeg.TopoReg
	serverStatus Status
}

func (mp *MultiPooler) CobraPreRunE(cmd *cobra.Command) error {
	_ = "STUB: not implemented"
	return nil
}

// NewMultiPooler creates a new MultiPooler instance with default configuration
func NewMultiPooler(telemetry *telemetry.Telemetry) *MultiPooler {
	_ = "STUB: not implemented"
	return nil
}

// RegisterFlags registers all multipooler flags with the given FlagSet
func (mp *MultiPooler) RegisterFlags(flags *pflag.FlagSet) { _ = "STUB: not implemented"; return }

// Init initializes the multipooler. If any services fail to start,
// or if some connections fail, it launches goroutines that retry
// until successful.
func (mp *MultiPooler) Init(startCtx context.Context) error { _ = "STUB: not implemented"; return nil }

// Resolve service ID early for telemetry resource attributes

// Get the configured logger

// Ensure we open the topo before we start the context, so that the
// defer that closes the topo runs after cancelling the context.
// This ensures that we've properly closed things like the watchers
// at that point.

// Validate libpq-style sslmode + sslrootcert before any pool opens. A typo
// or missing CA bundle should fail startup rather than silently downgrading
// the multipooler → postgres dials to plaintext.

// Create multipooler record with all fields now that servenv.Init() has set them up

// For now, all poolers start as REPLICA

// pgBackRest TLS certificate paths for connecting to primary's pgBackRest server

// Start the MultiPoolerManager

/* allowUpdate */

// For poolers, we don't un-register them on shutdown (they are persistent component)
// If they are actually deleted, they need to be cleaned up outside the lifecycle of starting / stopping.

// Set PoolerType to DRAINED so administrative views
// (e.g. `multigres getpoolers`) reflect that this pooler
// has gone away rather than continuing to display its
// last live role (often PRIMARY for a node that just
// failed over). This is purely cosmetic: failover keys
// off LeadershipStatus.REQUESTING_DEMOTION on the
// health stream, not off PoolerType. On restart the
// pooler re-registers with PoolerType_REPLICA and the
// orchestrator promotes as usual.

/* alarm */

// Database returns the configured database name.
func (mp *MultiPooler) Database() string { _ = "STUB: not implemented"; return "" }

func (mp *MultiPooler) RunDefault() error { _ = "STUB: not implemented"; return nil }

func (mp *MultiPooler) Shutdown() { _ = "STUB: not implemented"; return }
