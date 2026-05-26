// Copyright 2019 The Vitess Authors.
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
//
// Modifications Copyright 2025 Supabase, Inc.

package servenv

import (
	"context"
	"time"

	"github.com/multigres/multigres/go/tools/viperutil"

	"github.com/spf13/pflag"

	"google.golang.org/grpc"
)

// This file handles gRPC server, on its own port.
// Clients register servers, based on service map:
//
// servenv.RegisterGRPCFlags()
//
//	servenv.OnRun(func() {
//	  if servenv.GRPCCheckServiceMap("XXX") {
//	    pb.RegisterXXX(servenv.GRPCServer, XXX)
//	  }
//	}
//
// Note servenv.GRPCServer can only be used in servenv.OnRun,
// and not before, as it is initialized right before calling OnRun.

// GrpcServer holds all gRPC server configuration and the server instance
type GrpcServer struct {
	// auth specifies which auth plugin to use. Currently only "static" and "mtls" are supported.
	auth viperutil.Value[string]

	// port is the port to listen on for gRPC. If zero, don't listen.
	port viperutil.Value[int]

	// bindAddress is the address to bind to for gRPC. If empty, bind to all addresses.
	bindAddress viperutil.Value[string]

	// maxConnectionAge is the maximum age of a client connection, before GoAway is sent.
	maxConnectionAge viperutil.Value[time.Duration]

	// maxConnectionAgeGrace is an additional grace period after maxConnectionAge
	maxConnectionAgeGrace viperutil.Value[time.Duration]

	// initialConnWindowSize sets window size for a connection.
	initialConnWindowSize viperutil.Value[int]

	// initialWindowSize sets window size for stream.
	initialWindowSize viperutil.Value[int]

	// keepAliveEnforcementPolicyMinTime sets the keepalive enforcement policy on the server.
	keepAliveEnforcementPolicyMinTime viperutil.Value[time.Duration]

	// keepAliveEnforcementPolicyPermitWithoutStream allows keepalive pings even when there are no active streams
	keepAliveEnforcementPolicyPermitWithoutStream viperutil.Value[bool]

	// keepaliveTime is the time after which the server pings the client if no activity is seen
	keepaliveTime viperutil.Value[time.Duration]

	// keepaliveTimeout is the wait time after keepalive ping before closing the connection
	keepaliveTimeout viperutil.Value[time.Duration]

	// cert is the certificate file path for TLS
	cert viperutil.Value[string]

	// key is the private key file path for TLS
	key viperutil.Value[string]

	// ca is the CA file path for TLS
	ca viperutil.Value[string]

	// crl is the Certificate Revocation List file path
	crl viperutil.Value[string]

	// enableOptionalTLS enables optional TLS mode
	enableOptionalTLS viperutil.Value[bool]

	// serverCA is the server CA file path
	serverCA viperutil.Value[string]

	// Server is the actual gRPC server instance
	Server *grpc.Server

	// authPlugin is the authenticator plugin
	authPlugin Authenticator

	// socketFile is the named socket for RPCs
	socketFile viperutil.Value[string]
}

// NewGrpcServer creates and initializes a new GrpcServer with viperutil values
func NewGrpcServer(reg *viperutil.Registry) *GrpcServer { _ = "STUB: not implemented"; return nil }

// RegisterFlags registers all gRPC server flags with the given FlagSet
func (g *GrpcServer) RegisterFlags(fs *pflag.FlagSet) { _ = "STUB: not implemented"; return }

// Cert returns the certificate path
func (g *GrpcServer) Cert() string { _ = "STUB: not implemented"; return "" }

// CA returns the CA path
func (g *GrpcServer) CA() string {
	_ = "STUB: not implemented"

	// Key returns the key path
	return ""
}

func (g *GrpcServer) Key() string {
	_ = "STUB: not implemented"

	// Port returns the gRPC port
	return ""
}

func (g *GrpcServer) Port() int { _ = "STUB: not implemented"; return 0 }

// BindAddress returns the bind address
func (g *GrpcServer) BindAddress() string { _ = "STUB: not implemented"; return "" }

// SocketFile returns the Unix socket file path, or empty string if not configured.
func (g *GrpcServer) SocketFile() string { _ = "STUB: not implemented"; return "" }

// IsEnabled returns true if gRPC server is enabled
func (g *GrpcServer) IsEnabled() bool { _ = "STUB: not implemented"; return false }

// Create creates the gRPC server instance.
// It has to be called after flags are parsed, but before services register themselves.
func (g *GrpcServer) Create() error {
	_ = "STUB: not implemented"
	// skip if not enabled
	return nil
}

// Fail-fast on flags that aren't implemented yet, before doing any
// further setup work.

// Build TLS config if cert and key files are provided.
// When --grpc-ca is also set, mutual TLS is enabled (client certs required).
// BuildServerTLSConfig validates the cert/key/ca combinations.

// Validate that mTLS auth mode is not used without transport TLS.

// Override the default max message size for both send and receive
// (which is 4 MiB in gRPC 1.0.0).
// Large messages can occur when users try to insert or fetch very big
// rows. If they hit the limit, they'll see the following error:
// grpc: received message length XXXXXXX exceeding the max size 4194304
// Note: For gRPC 1.0.0 it's sufficient to set the limit on the server only
// because it's not enforced on the client side.

// Add OpenTelemetry instrumentation for distributed tracing and metrics
// If no OTEL exporters are configured, noop exporters are used with minimal overhead

// interceptors builds the list of interceptors for the gRPC server
func (g *GrpcServer) interceptors() ([]grpc.ServerOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Serve starts the gRPC server and begins listening for requests
func (g *GrpcServer) Serve(sv *ServEnv) error {
	_ = "STUB: not implemented"
	// skip if not enabled
	return nil
}

// register reflection to support list calls :)

// register the shared ServiceInfo service so every multigres process
// exposes build identity (revision, go version, etc.) on its gRPC port

// register health service to support health checks

// listen on the port

// and serve on it
// NOTE: Before we call Serve(), all services must have registered themselves
//       with the Server. This is the case because go/common/servenv/run.go
//       runs all OnRun() hooks after Create() and before Serve().
//       If this was not the case, the binary would crash with
//       the error "grpc: Server.RegisterService after Server.Serve".

// CheckServiceMap returns if we should register a gRPC service
func (g *GrpcServer) CheckServiceMap(name string, sv *ServEnv) bool {
	_ = "STUB: not implemented"
	// Silently fail individual services if gRPC is not enabled in
	// the first place (either on a grpc port or on the socket file)
	return false
}

// then check ServiceMap

func (g *GrpcServer) authenticatingStreamInterceptor(srv any, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *GrpcServer) authenticatingUnaryInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// WrappedServerStream is based on the service stream wrapper from: https://github.com/grpc-ecosystem/go-grpc-middleware
type WrappedServerStream struct {
	grpc.ServerStream
	WrappedContext context.Context
}

// Context returns the wrapper's WrappedContext, overwriting the nested grpc.ServerStream.Context()
func (w *WrappedServerStream) Context() context.Context {
	_ = "STUB: not implemented"
	return *

	// WrapServerStream returns a ServerStream that has the ability to overwrite context.
	new(context.Context)
}

func WrapServerStream(stream grpc.ServerStream) *WrappedServerStream {
	_ = "STUB: not implemented"
	return nil
}

// serverInterceptorBuilder chains together multiple ServerInterceptors
type serverInterceptorBuilder struct {
	streamInterceptors []grpc.StreamServerInterceptor
	unaryInterceptors  []grpc.UnaryServerInterceptor
}

// Add adds interceptors to the builder
func (collector *serverInterceptorBuilder) Add(s grpc.StreamServerInterceptor, u grpc.UnaryServerInterceptor) {
	_ = "STUB: not implemented"
	return
}

// AddUnary adds a single unary interceptor to the builder
func (collector *serverInterceptorBuilder) AddUnary(u grpc.UnaryServerInterceptor) {
	_ = "STUB: not implemented"
	return
}

// Build returns DialOptions to add to the grpc.Dial call
func (collector *serverInterceptorBuilder) Build() []grpc.ServerOption {
	_ = "STUB: not implemented"
	return nil
}
