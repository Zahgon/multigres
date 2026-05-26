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

package local

import (
	"context"
	"time"
)

// waitForServiceReady waits for a service to become ready by checking appropriate endpoints
// The provided parentCtx should contain any active span for distributed tracing
func (p *localProvisioner) waitForServiceReady(parentCtx context.Context, serviceName string, host string, servicePorts map[string]int, timeout time.Duration) error {
	_ = "STUB: not implemented"
	// Create span as child of parent context
	return nil
}

// First check TCP connectivity on all advertised ports

// This port not ready yet

// Not all ports ready yet, will backoff

// For services with HTTP endpoints, check debug/config endpoint

// HTTP endpoint not ready yet, will backoff

// Service is ready

// checkMultigresServiceHealth checks health for all supported service port types
// The context is propagated to all health check calls to maintain trace parent-child relationships
func (p *localProvisioner) checkMultigresServiceHealth(ctx context.Context, serviceName string, host string, servicePorts map[string]int) error {
	_ = "STUB: not implemented"
	// Iterate over service ports and run health checks for supported types
	return nil
}

// Run HTTP health check

// Run gRPC health check for pgctld

// TCP connectivity is already verified above; no HTTP check on the client port.
// /readyz is only available on the metrics listener (etcd_metrics_port).

// Run etcd readiness check via /readyz on the dedicated metrics listener.

// Future: Add other gRPC services

// No health check implemented for this port type, skip

// checkEtcdHealth checks if etcd is ready by querying its /readyz endpoint.
// address must be the host:port of etcd's metrics listener (--listen-metrics-urls),
// not the client listener. /readyz returns 404 on the client port.
func (p *localProvisioner) checkEtcdHealth(ctx context.Context, address string) error {
	_ = "STUB: not implemented"
	return nil
}

// checkDebugConfigEndpoint checks if the debug/config endpoint returns 200 OK
func (p *localProvisioner) checkDebugConfigEndpoint(ctx context.Context, address string) error {
	_ = "STUB: not implemented"
	return nil
}

// checkPgctldGrpcHealth checks if pgctld gRPC server is healthy by calling Status
func (p *localProvisioner) checkPgctldGrpcHealth(ctx context.Context, address string) error {
	_ = "STUB: not implemented"
	return nil
}

// validateProcessRunning checks if a process with the given PID is still running
func (p *localProvisioner) validateProcessRunning(pid int) error {
	_ = "STUB: not implemented"
	return nil
}

// Send signal 0 to check if process exists without actually sending a signal

// checkPortConflict checks if a port is already in use by another process
func (p *localProvisioner) checkPortConflict(port int, serviceName, portName string) error {
	_ = "STUB: not implemented"
	return nil

	// Skip invalid ports
}

// Port is not in use, this is good
//nolint:nilerr // Error means port is free, which is success

// Port is in use by some process
