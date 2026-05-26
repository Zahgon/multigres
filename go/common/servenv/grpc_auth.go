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

	"github.com/spf13/pflag"
	"google.golang.org/grpc"
)

var grpcAuthServerFlagHooks []func(*pflag.FlagSet)

// RegisterGRPCServerAuthFlags registers flags required to enable server-side
// authentication in multigres gRPC services.
//
// `go/cmd/*` entrypoints should call this function before
// ParseFlags(WithArgs)? if they wish to expose Authenticator functionality.
func RegisterGRPCServerAuthFlags() { _ = "STUB: not implemented"; return }

// Auth returns the auth mode
func (g *GrpcServer) Auth() string { _ = "STUB: not implemented"; return "" }

// Authenticator provides an interface to implement auth in Multigres in
// grpc server
type Authenticator interface {
	Authenticate(ctx context.Context, fullMethod string) (context.Context, error)
}

// authPlugins is a registry of AuthPlugin initializers.
var authPlugins = make(map[string]func() (Authenticator, error))

// RegisterAuthPlugin registers an implementation of AuthServer.
// Returns an error if a plugin with the same name is already registered.
func RegisterAuthPlugin(name string, authPlugin func() (Authenticator, error)) error {
	_ = "STUB: not implemented"
	return nil
}

// GetAuthenticator returns an AuthPlugin by name.
func GetAuthenticator(name string) (func() (Authenticator, error), error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FakeAuthStreamInterceptor fake interceptor to test plugin
func FakeAuthStreamInterceptor(srv any, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	_ = "STUB: not implemented"
	return nil
}

// FakeAuthUnaryInterceptor fake interceptor to test plugin
func FakeAuthUnaryInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func fakeDummyAuthenticate(ctx context.Context) bool { _ = "STUB: not implemented"; return false }
