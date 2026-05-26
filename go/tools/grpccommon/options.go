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

package grpccommon

import (
	"github.com/spf13/pflag"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/grpc"
)

var (
	// maxMessageSize is the maximum message size which the gRPC server will
	// accept. Larger messages will be rejected.
	// Note: We're using 16 MiB as default value because that's the default in MySQL
	maxMessageSize = 16 * 1024 * 1024
	// enablePrometheus sets a flag to enable grpc client/server grpc monitoring.
	enablePrometheus bool
)

// RegisterFlags installs grpccommon flags on the given FlagSet.
//
// `go/cmd/*` entrypoints should either use servenv.ParseFlags(WithArgs)? which
// calls this function, or call this function directly before parsing
// command-line arguments.
func RegisterFlags(fs *pflag.FlagSet) { _ = "STUB: not implemented"; return }

// EnableGRPCPrometheus returns the value of the --grpc-prometheus flag.
func EnableGRPCPrometheus() bool { _ = "STUB: not implemented"; return false }

// MaxMessageSize returns the value of the --grpc-max-message-size flag.
func MaxMessageSize() int { _ = "STUB: not implemented"; return 0 }

// LocalClientDialOptions returns a slice of grpc.DialOption to be used when creating a gRPC client.
// These options are used for local clients connecting to the gRPC server.
// They are not intended to be used for production environments.
// The WithDisableServiceConfig is a workaround for a known issue
// in MacOS where localhost host takes too long to resolve.
// See the following PR for more details: https://github.com/multigres/multigres/pull/152
func LocalClientDialOptions() []grpc.DialOption { _ = "STUB: not implemented"; return nil }

// ClientDialOptions returns the standard dial options for a gRPC client given
// a caller-supplied transport credentials dial option. It always includes
// WithDisableServiceConfig, a macOS localhost-resolution workaround (see #152),
// so callers can't forget it when wiring up TLS or insecure credentials.
func ClientDialOptions(transportCreds grpc.DialOption) []grpc.DialOption {
	_ = "STUB: not implemented"
	return nil
}

// ClientOption configures OpenTelemetry instrumentation for the gRPC client.
// These options extend the stats handler that NewClient creates.
type ClientOption interface {
	apply(*clientConfig)
}

type clientConfig struct {
	otelOptions []otelgrpc.Option
	dialOptions []grpc.DialOption
}

// WithAttributes adds custom OpenTelemetry attributes to gRPC client spans.
// This is a generic helper that can be used by domain-specific code to add
// custom span attributes without making grpccommon domain-aware.
func WithAttributes(attrs ...attribute.KeyValue) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

// WithDialOptions adds standard gRPC dial options to the client.
func WithDialOptions(opts ...grpc.DialOption) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

type funcOption func(*clientConfig)

func (f funcOption) apply(c *clientConfig) {
	_ = "STUB: not implemented"

	// NewClient creates a gRPC client with OpenTelemetry instrumentation.
	// Use WithPeerService to set the remote service identifier in traces.
	// Use WithDialOptions to pass standard gRPC dial options.
	//
	// All ClientOptions are used to configure a single stats handler, preventing
	// duplication and ensuring consistent telemetry across the application.
	return
}

func NewClient(target string, opts ...ClientOption) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create single stats handler with all OTel options
