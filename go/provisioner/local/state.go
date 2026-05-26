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
	"time"
)

// LocalProvisionedService represents a service instance that has been provisioned
type LocalProvisionedService struct {
	ID         string         `json:"id"`                    // Unique instance ID
	Service    string         `json:"service"`               // Service name (etcd, multigateway, etc.)
	PID        int            `json:"pid,omitempty"`         // For binary processes
	BinaryPath string         `json:"binary-path,omitempty"` // Path to the binary
	DataDir    string         `json:"data-dir,omitempty"`    // Data directory
	LogFile    string         `json:"log-file,omitempty"`    // Path to log file
	Ports      map[string]int `json:"ports"`                 // Port mappings
	FQDN       string         `json:"fqdn"`                  // Hostname/FQDN
	Runtime    string         `json:"runtime"`               // "binary"
	StartedAt  time.Time      `json:"started-at"`            // When it was started
	Metadata   map[string]any `json:"metadata,omitempty"`    // Additional metadata
}

// getStateDir returns the path to the state directory
func (p *localProvisioner) getStateDir() string { _ = "STUB: not implemented"; return "" }

// getLogsDir returns the path to the logs directory
func (p *localProvisioner) getLogsDir() string { _ = "STUB: not implemented"; return "" }

func (p *localProvisioner) getDataDir() string { _ = "STUB: not implemented"; return "" }

// createLogFile creates a log file path and ensures the directory exists
func (p *localProvisioner) createLogFile(serviceName, serviceID, databaseName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// For database services: logs/dbs/dbname/servicename

// For non-database services (like etcd): logs/servicename

// Create the service-specific log directory

// Create the log file path

// cleanupLogFile removes a log file if it exists
func (p *localProvisioner) cleanupLogFile(logFilePath string) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if log file exists

// File doesn't exist, nothing to clean up

// Remove the log file

// saveServiceState saves the provisioned service state to disk
func (p *localProvisioner) saveServiceState(service *LocalProvisionedService, databaseName string) error {
	_ = "STUB: not implemented"
	return nil
}

// For database services: state/dbs/dbname

// For non-database services (like etcd): state/

// File name format: service_id.json (e.g., etcd_abc123.json)

// removeServiceState removes a service state file from disk
func (p *localProvisioner) removeServiceState(serviceID, serviceName, databaseName string) error {
	_ = "STUB: not implemented"
	return nil
}

// For database services: state/dbs/dbname

// For non-database services (like etcd): state/

// loadDbProvisionedServices loads provisioned services for a specific database
func (p *localProvisioner) loadDbProvisionedServices(databaseName string) ([]*LocalProvisionedService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if target directory exists

// No services for this database

// Parse filename: service_id.json

// Skip invalid state files

// Log warning but continue with other services

// loadEtcdServices loads etcd services from the top-level state directory
func (p *localProvisioner) loadEtcdServices() ([]*LocalProvisionedService, error) {
	_ = "STUB: not implemented"
	return nil,

		// Check if state directory exists
		nil
}

// No state directory, no services running

// Parse filename: etcd_serviceID.json

// etcd is a global service

// Log warning but continue with other services

// loadGlobalServices loads all global services (non-database services) from state files
func (p *localProvisioner) loadGlobalServices() ([]*LocalProvisionedService, error) {
	_ = "STUB: not implemented"
	return nil,

		// Check if state directory exists
		nil
}

// No state directory, no services running

// Parse filename: serviceName_serviceID.json

// Load global services (non-etcd services can be included here)

// global services have no database name

// Log warning but continue with other services

// findRunningDbService finds a running service by service name within a specific database and cell
func (p *localProvisioner) findRunningDbService(serviceName, databaseName, cell string) (*LocalProvisionedService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if the service matches the cell

// Check if the service is actually still running

// Check for port conflicts with other processes using cell-specific config

// No running service found

// getExpectedPortsForDbService returns expected ports for a DB-scoped service (per cell)
func (p *localProvisioner) getExpectedPortsForDbService(serviceName, cell string) map[string]int {
	_ = "STUB: not implemented"
	return nil
}

// getExpectedPortsForService returns the expected ports for a service based on its configuration
func (p *localProvisioner) getExpectedPortsForService(serviceName string) map[string]int {
	_ = "STUB: not implemented"
	return nil
}

// findRunningEtcdService finds a running etcd service
func (p *localProvisioner) findRunningEtcdService() (*LocalProvisionedService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if the service is actually still running

// Check for port conflicts with other processes

// No running etcd service found

// findRunningService finds a running service by service name (for global services like multiadmin)
func (p *localProvisioner) findRunningService(serviceName string) (*LocalProvisionedService, error) {
	_ = "STUB: not implemented"
	// Load global services (e.g., multiadmin, etcd)
	return nil, nil
}

// Check if the service is actually still running

// Check for port conflicts with other processes

// No running service found
