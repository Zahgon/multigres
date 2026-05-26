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
	"fmt"
	"log/slog"

	"github.com/spf13/pflag"
)

var (
	// clientCertSubstrings list of substrings of at least one of the client certificate names to use during authorization
	clientCertSubstrings string
	// MtlsAuthPlugin implements AuthPlugin interface
	_ Authenticator = (*MtlsAuthPlugin)(nil)
)

func registerGRPCServerAuthMTLSFlags(fs *pflag.FlagSet) { _ = "STUB: not implemented"; return }

// MtlsAuthPlugin  implements static username/password authentication for grpc. It contains an array of username/passwords
// that will be authorized to connect to the grpc server.
type MtlsAuthPlugin struct {
	clientCertSubstrings []string
}

// Authenticate implements Authenticator interface. This method will be used inside a middleware in grpc_server to authenticate
// incoming requests.
func (ma *MtlsAuthPlugin) Authenticate(ctx context.Context, fullMethod string) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func mtlsAuthPluginInitializer() (Authenticator, error) {
	_ = "STUB: not implemented"
	return *new(Authenticator), nil
}

// ClientCertSubstrings returns the value of the
// `--grpc-auth-mtls-allowed-substrings` flag.
func ClientCertSubstrings() string { _ = "STUB: not implemented"; return "" }

func init() {
	if err := RegisterAuthPlugin("mtls", mtlsAuthPluginInitializer); err != nil {
		slog.Error("failed to register mtls auth plugin", "error", err)
		panic(fmt.Sprintf("failed to register mtls auth plugin: %v", err))
	}
	grpcAuthServerFlagHooks = append(grpcAuthServerFlagHooks, registerGRPCServerAuthMTLSFlags)
}
