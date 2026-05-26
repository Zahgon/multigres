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

package command

import (
	"context"
	"log/slog"
	"sync"

	"github.com/multigres/multigres/go/common/servenv"
	"github.com/multigres/multigres/go/services/pgctld"
	"github.com/multigres/multigres/go/tools/executil"
	"github.com/multigres/multigres/go/tools/viperutil"

	"github.com/spf13/cobra"

	pb "github.com/multigres/multigres/go/pb/pgctldservice"
)

// intToInt32 safely converts int to int32 for protobuf fields.
// Returns an error if value exceeds int32 range (should never happen for PIDs/ports).
func intToInt32(v int) (int32, error) { _ = "STUB: not implemented"; return 0, nil }

// PgCtldServerCmd holds the server command configuration
type PgCtldServerCmd struct {
	pgCtlCmd          *PgCtlCommand
	grpcServer        *servenv.GrpcServer
	senv              *servenv.ServEnv
	pgbackrestPort    viperutil.Value[int]
	pgbackrestCertDir viperutil.Value[string]
}

// AddServerCommand adds the server subcommand to the root command
func AddServerCommand(root *cobra.Command, pc *PgCtlCommand) { _ = "STUB: not implemented"; return }

// validateServerFlags validates required flags for the server command
func (s *PgCtldServerCmd) validateServerFlags(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *PgCtldServerCmd) createCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Don't register logger and viper config flags since they're already registered
// as persistent flags in the root command and we're sharing those instances

// pgBackRest TLS server flags (server command only)

func (s *PgCtldServerCmd) runServer(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	// TODO(dweitzman): Add ServiceInstanceID and Cell that relate to the multipooler this pgctld serves
	return nil
}

// Get the configured logger

// Start reaping orphaned children to prevent zombie processes.
// Only start when running as PID 1 (container init process). pg_ctl with -W
// forks a child that gets reparented to PID 1 on exit; without this reaper
// those children become zombies.
//
// When pgctld is NOT PID 1 (tests, CI, systemd), orphaned children are
// reparented to the system init — not pgctld — so the reaper is unnecessary.
// Starting it anyway races with cmd.Wait() in RPC handlers (initdb, pg_rewind)
// causing "waitid: no child processes" errors.

// Create and register our service

// Register /ready probe: postgres socket accepting + gRPC accepting.
// Replication health is intentionally excluded (see pgctldReadyHandler).

// Start pgBackRest management

// Register gRPC service with the global GRPCServer

// reapOrphanedChildren handles SIGCHLD signals to reap zombie processes.
// This is necessary because pg_ctl with -W flag creates child processes that get
// reparented to pgctld (when running as PID 1 in a container). Without this reaper,
// these child processes remain in defunct (zombie) state after exit.
//
// The function runs in a goroutine and continuously waits for SIGCHLD signals,
// then reaps all available zombie children using Wait4 with WNOHANG.
func reapOrphanedChildren(logger *slog.Logger) { _ = "STUB: not implemented"; return }

// Reap all zombie children

// No more children to reap

// PgCtldServiceConfig holds the PostgreSQL instance identity and initialization
// parameters. These are the most commonly passed parameters and are grouped
// here to reduce argument lists.
type PgCtldServiceConfig struct {
	Port     int
	User     string
	Database string
	Password string
	// PasswordSource records where Password came from so log lines can report
	// it without leaking the value itself. Optional; defaults to none.
	PasswordSource PasswordSource
	// PasswordFile is the absolute path to the operator-supplied password file
	// when PasswordSource == PasswordSourceFile, otherwise "". initdb is handed
	// this path directly via --pwfile so the plaintext never lands in /tmp.
	PasswordFile         string
	InitdbArgs           string
	InitdbSQLFiles       []string
	InitdbSQLDirs        []string
	InitdbExtraConfFiles []string
}

// PgCtldService implements the pgctld gRPC service
type PgCtldService struct {
	pb.UnimplementedPgCtldServer
	logger     *slog.Logger
	ctldConfig PgCtldServiceConfig
	timeout    int
	poolerDir  string
	pgConfig   *pgctld.PostgresCtlConfig

	// pgBackRest management
	ctx              context.Context
	cancel           context.CancelFunc
	wg               sync.WaitGroup
	pgBackRestCmd    *executil.Cmd
	pgBackRestStatus *pb.PgBackRestStatus
	statusMu         sync.RWMutex
	restartCount     int32
	metrics          *Metrics
}

// pgbackrestServerConfigPath returns the path to the pgbackrest server config file.
func (s *PgCtldService) pgbackrestServerConfigPath() string { _ = "STUB: not implemented"; return "" }

// NewPgCtldService creates a new PgCtldService with validation
func NewPgCtldService(
	logger *slog.Logger,
	cfg PgCtldServiceConfig,
	timeout int,
	poolerDir string,
	listenAddresses string,
	pgbackrestPort int,
	pgbackrestCertDir string,
) (*PgCtldService, error) {
	_ = "STUB: not implemented"
	// Validate essential parameters for service creation
	// Note: We don't validate postgresDataDir or postgresConfigFile existence here
	// because the server should be able to start even with uninitialized data directory
	return nil, nil
}

// cfg.Password emptiness is not re-checked here: production callers
// (runServer) populate it via PgCtlCommand.GetPostgresPassword, which
// returns an error when no password source is configured.

// Write a pgpass file and set PGPASSFILE so pgbackrest (archive-push runs as a
// postgres subprocess and inherits this process's environment) can authenticate
// against PostgreSQL without exposing the password in the process environment.

// Create the PostgreSQL config once during service initialization

// Generate pgbackrest-server.conf if pgbackrest port and cert dir provided

//nolint:gocritic // Background context for pgBackRest lifecycle management

// setPgBackRestStatus updates the pgBackRest status thread-safely and returns the current restart count
func (s *PgCtldService) setPgBackRestStatus(running bool, errorMessage string, incrementRestart bool) int32 {
	_ = "STUB: not implemented"
	return 0
}

// getPgBackRestStatus returns a copy of the current status thread-safely
func (s *PgCtldService) getPgBackRestStatus() *pb.PgBackRestStatus {
	_ = "STUB: not implemented"
	return nil
}

// Return a copy to avoid race conditions

// Close shuts down the pgctld service gracefully
func (s *PgCtldService) Close() { _ = "STUB: not implemented"; return }

// Signal managePgBackRest goroutine to stop

// Kill pgBackRest process if running

// Wait for goroutines to fully exit

// startPgBackRest starts the pgBackRest TLS server process
// Returns the command on success, or error on failure
func (s *PgCtldService) startPgBackRest(ctx context.Context) (*executil.Cmd, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Verify config exists

// Build command: pgbackrest server
// Note: Config is passed via PGBACKREST_CONFIG environment variable

// Start the process

// managePgBackRest manages the pgBackRest TLS server lifecycle with retry and restart logic
func (s *PgCtldService) managePgBackRest(ctx context.Context) {
	_ = "STUB: not implemented"
	// Check if pgbackrest config exists before attempting to start
	return
}

// Try to start with retry policy (max 5 attempts)

// Context cancelled during startup - clean shutdown

// Success, exit retry loop

// Failed to start after retries

// Wait for exit OR context cancellation

// Ignore error - process exit is expected

// Process exited, restart

// Shutdown requested, stop process

// Ignore error - process may already be dead

// Wait for Wait() to complete

// StartPgBackRestManagement begins pgBackRest management in background
func (s *PgCtldService) StartPgBackRestManagement() { _ = "STUB: not implemented"; return }

func (s *PgCtldService) Start(ctx context.Context, req *pb.StartRequest) (*pb.StartResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if data directory is initialized

// Use the pre-configured PostgreSQL config for start operation

func (s *PgCtldService) Stop(ctx context.Context, req *pb.StopRequest) (*pb.StopResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if data directory is initialized

// Use the pre-configured PostgreSQL config for stop operation

func (s *PgCtldService) Restart(ctx context.Context, req *pb.RestartRequest) (*pb.RestartResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if data directory is initialized

// Use the pre-configured PostgreSQL config for restart operation

func (s *PgCtldService) ReloadConfig(ctx context.Context, req *pb.ReloadConfigRequest) (*pb.ReloadConfigResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if data directory is initialized

// Use the pre-configured PostgreSQL config for reload operation

func (s *PgCtldService) Status(ctx context.Context, req *pb.StatusRequest) (*pb.StatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// First check if data directory is initialized

// Use the pre-configured PostgreSQL config for status operation

// Convert status string to protobuf enum

func (s *PgCtldService) Version(ctx context.Context, req *pb.VersionRequest) (*pb.VersionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *PgCtldService) InitDataDir(ctx context.Context, req *pb.InitDataDirRequest) (*pb.InitDataDirResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use the shared init function with detailed result

func (s *PgCtldService) PgRewind(ctx context.Context, req *pb.PgRewindRequest) (*pb.PgRewindResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if postgres is cleanly stopped before pg_rewind
// If not, try crash recovery - this is needed for rewind dry-run to work
// This check is best effort. It's not harmful to try the pg_rewind if
// crash recovery fails, the dry run is just unlikely to succeed in that case.

// Try to run crash recovery.
// It's not harmful to do this if postgres is already running.

// Construct source server connection string (without password - will use PGPASSWORD env var)
// Include application_name if provided (used for replication identification)
// connect_timeout ensures pg_rewind fails fast if the source is not yet ready to accept
// connections (e.g. newly-promoted primary still in crash recovery), allowing the caller
// to retry rather than blocking for the OS TCP timeout.

// Use the shared rewind function with detailed result, passing password separately
