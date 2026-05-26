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

package testutil

import (
	"context"
	"sync"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	pb "github.com/multigres/multigres/go/pb/pgctldservice"
)

// MockPgCtldService implements a mock version of the PgCtld gRPC service for testing
type MockPgCtldService struct {
	pb.UnimplementedPgCtldServer
	mu            sync.Mutex
	StartCalls    []*pb.StartRequest
	StopCalls     []*pb.StopRequest
	RestartCalls  []*pb.RestartRequest
	ReloadCalls   []*pb.ReloadConfigRequest
	StatusCalls   []*pb.StatusRequest
	VersionCalls  []*pb.VersionRequest
	InitDirCalls  []*pb.InitDataDirRequest
	PgRewindCalls []*pb.PgRewindRequest

	// Response configurations
	StartResponse    *pb.StartResponse
	StopResponse     *pb.StopResponse
	RestartResponse  *pb.RestartResponse
	ReloadResponse   *pb.ReloadConfigResponse
	StatusResponse   *pb.StatusResponse
	VersionResponse  *pb.VersionResponse
	InitDirResponse  *pb.InitDataDirResponse
	PgRewindResponse *pb.PgRewindResponse

	// Error configurations
	StartError    error
	StopError     error
	RestartError  error
	ReloadError   error
	StatusError   error
	VersionError  error
	InitDirError  error
	PgRewindError error
}

func (m *MockPgCtldService) Start(ctx context.Context, req *pb.StartRequest) (*pb.StartResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MockPgCtldService) Stop(ctx context.Context, req *pb.StopRequest) (*pb.StopResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MockPgCtldService) Restart(ctx context.Context, req *pb.RestartRequest) (*pb.RestartResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MockPgCtldService) ReloadConfig(ctx context.Context, req *pb.ReloadConfigRequest) (*pb.ReloadConfigResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MockPgCtldService) Status(ctx context.Context, req *pb.StatusRequest) (*pb.StatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MockPgCtldService) Version(ctx context.Context, req *pb.VersionRequest) (*pb.VersionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MockPgCtldService) InitDataDir(ctx context.Context, req *pb.InitDataDirRequest) (*pb.InitDataDirResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MockPgCtldService) PgRewind(ctx context.Context, req *pb.PgRewindRequest) (*pb.PgRewindResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Default response: no divergence for dry-run, success for actual rewind

// Empty output means no divergence

// TestGRPCServer provides utilities for testing gRPC services
type TestGRPCServer struct {
	server   *grpc.Server
	listener *bufconn.Listener
	address  string
}

// NewTestGRPCServer creates a new test gRPC server with bufconn
func NewTestGRPCServer(t *testing.T) *TestGRPCServer { _ = "STUB: not implemented"; return nil }

// RegisterService registers a service with the test server
func (ts *TestGRPCServer) RegisterService(service pb.PgCtldServer) {
	_ = "STUB: not implemented"
	return
}

// Start starts the test gRPC server
func (ts *TestGRPCServer) Start(t *testing.T) { _ = "STUB: not implemented"; return }

// Stop stops the test gRPC server
func (ts *TestGRPCServer) Stop() { _ = "STUB: not implemented"; return }

// NewClient creates a new gRPC client connected to the test server
func (ts *TestGRPCServer) NewClient(t *testing.T) pb.PgCtldClient {
	_ = "STUB: not implemented"
	return *new(pb.PgCtldClient)
}

// StartTestServer starts a real gRPC server on a random port for integration testing
func StartTestServer(t *testing.T, service pb.PgCtldServer) (pb.PgCtldClient, func()) {
	_ = "STUB: not implemented"

	// Find an available port
	return *new(pb.PgCtldClient), nil
}

// Create gRPC server

// Start server in background

// Create client

// Wait for server to be ready

// StartMockPgctldServer starts a mock pgctld server with the provided mock service
// Returns the server address and a cleanup function
func StartMockPgctldServer(t *testing.T, mockService *MockPgCtldService) (string, func()) {
	_ = "STUB: not implemented"

	// Create a listener on a random port
	return "", nil
}

// Create gRPC server with mock service

// Start serving in background
