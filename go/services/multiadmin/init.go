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

// Package multiadmin provides multiadmin functionality.
package multiadmin

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/multigres/multigres/go/common/rpcclient"
	"github.com/multigres/multigres/go/common/servenv"
	"github.com/multigres/multigres/go/common/topoclient"
)

type MultiAdmin struct {
	// adminServer holds the gRPC admin server instance
	adminServer *MultiAdminServer

	// grpcServer is the grpc server
	grpcServer *servenv.GrpcServer

	// senv is the serving environment
	senv *servenv.ServEnv

	// connConfig holds RPC client configuration (TLS, etc.)
	connConfig *rpcclient.ConnConfig

	// topoConfig holds topology configuration
	topoConfig   *topoclient.TopoConfig
	ts           topoclient.Store
	serverStatus Status
}

func (ma *MultiAdmin) RunDefault() error { _ = "STUB: not implemented"; return nil }

func (ma *MultiAdmin) CobraPreRunE(cmd *cobra.Command) error { _ = "STUB: not implemented"; return nil }

func NewMultiAdmin() *MultiAdmin { _ = "STUB: not implemented"; return nil }

// RegisterFlags registers flags specific to multiadmin.
func (ma *MultiAdmin) RegisterFlags(fs *pflag.FlagSet) { _ = "STUB: not implemented"; return }

// Init initializes the multiadmin. If any services fail to start,
// or if some connections fail, it launches goroutines that retry
// until successful.
func (ma *MultiAdmin) Init(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Get the configured logger

// Register multiadmin gRPC and HTTP API services if enabled in service map

// Set up grpc-gateway for REST API

// NOTE: The ctx parameter to the generated method here is unused.

func (ma *MultiAdmin) Shutdown() { _ = "STUB: not implemented"; return }
