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

package multiadmin

import (
	"context"
)

// ServiceInfo represents a discoverable service in the cluster
type ServiceInfo struct {
	Name       string // Service name (e.g., "multigateway")
	Cell       string // Cell name (e.g., "zone1"), empty for global services
	ProxiedURL string // Proxied URL (e.g., "/proxy/gate/zone1/multigateway")
	DirectURL  string // Direct URL if available (e.g., "http://localhost:15001/")
}

// ServiceList holds all discovered services organized by scope
type ServiceList struct {
	GlobalServices []ServiceInfo            // Global services (multiadmin, etc.)
	CellServices   map[string][]ServiceInfo // Cell name -> services in that cell
	Error          string                   // Error message if discovery failed
}

// DiscoverServices queries the topology store and builds a list of all services
// in the cluster. It returns a ServiceList with services organized by scope.
// This function may be slow (topo queries), so it should be called from a
// dedicated endpoint, not the fast-path root handler.
func (ma *MultiAdmin) DiscoverServices(ctx context.Context) (*ServiceList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add multiadmin as a global service
// Multiadmin is always available if this code is running

// Discover cells from topology

// If we can't get cells, still return multiadmin but note the error

// For each cell, discover actual registered services from topology

// Discover multigateway services

// Use the first registered gateway

// Discover multipooler services

// Use the first registered pooler

// Discover multiorch services

// Use the first registered orch

// Sort services within cell alphabetically

// Sort global services alphabetically
