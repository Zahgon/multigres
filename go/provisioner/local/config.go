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

// BackupConfig holds backup configuration (local or S3)
type BackupConfig struct {
	Type  string       `yaml:"type"` // "local", "s3", "azure", etc.
	Local *LocalBackup `yaml:"local,omitempty"`
	S3    *S3Backup    `yaml:"s3,omitempty"`
}

// LocalBackup holds filesystem backup configuration
type LocalBackup struct {
	Path string `yaml:"path"`
}

// S3Backup holds S3 backup configuration
type S3Backup struct {
	Bucket            string `yaml:"bucket"`
	Region            string `yaml:"region"`
	Endpoint          string `yaml:"endpoint,omitempty"`
	KeyPrefix         string `yaml:"key-prefix,omitempty"`
	UseEnvCredentials bool   `yaml:"use-env-credentials,omitempty"`
}

// CellConfig holds the configuration for a single cell
type CellConfig struct {
	Name     string `yaml:"name"`
	RootPath string `yaml:"root-path"`
}

// TopologyConfig holds the configuration for cluster topology
type TopologyConfig struct {
	GlobalRootPath string       `yaml:"global-root-path"`
	Cells          []CellConfig `yaml:"cells"`
}

// CellServicesConfig holds the service configuration for a specific cell
type CellServicesConfig struct {
	Multigateway MultigatewayConfig `yaml:"multigateway"`
	Multipooler  MultipoolerConfig  `yaml:"multipooler"`
	Multiorch    MultiorchConfig    `yaml:"multiorch"`
	Pgctld       PgctldConfig       `yaml:"pgctld"`
}

// LocalProvisionerConfig represents the typed configuration for the local provisioner
type LocalProvisionerConfig struct {
	RootWorkingDir string                        `yaml:"root-working-dir"`
	DefaultDbName  string                        `yaml:"default-db-name"`
	Backup         BackupConfig                  `yaml:"backup"`
	Etcd           EtcdConfig                    `yaml:"etcd"`
	Topology       TopologyConfig                `yaml:"topology"`
	Multiadmin     MultiadminConfig              `yaml:"multiadmin"`
	Cells          map[string]CellServicesConfig `yaml:"cells,omitempty"`
}

// EtcdConfig holds etcd service configuration
type EtcdConfig struct {
	Version  string `yaml:"version"`
	DataDir  string `yaml:"data-dir"`
	Port     int    `yaml:"port"`                // Client port
	PeerPort int    `yaml:"peer-port,omitempty"` // Optional peer port, defaults to Port+1
}

// MultigatewayConfig holds multigateway service configuration
type MultigatewayConfig struct {
	Path     string `yaml:"path"`
	HttpPort int    `yaml:"http-port"`
	GrpcPort int    `yaml:"grpc-port"`
	PgPort   int    `yaml:"pg-port"`
	LogLevel string `yaml:"log-level"`
}

// MultipoolerConfig holds multipooler service configuration
type MultipoolerConfig struct {
	Path           string `yaml:"path"`
	Database       string `yaml:"database"`
	TableGroup     string `yaml:"table-group"`
	Shard          string `yaml:"shard"`
	ServiceID      string `yaml:"service-id"`
	PoolerDir      string `yaml:"pooler-dir"` // Directory path for PostgreSQL socket files
	PgPort         int    `yaml:"pg-port"`    // PostgreSQL port number (same as pgctld)
	HttpPort       int    `yaml:"http-port"`
	GrpcPort       int    `yaml:"grpc-port"`
	GRPCSocketFile string `yaml:"grpc-socket-file"` // Unix socket file path for gRPC
	LogLevel       string `yaml:"log-level"`
}

// MultiorchConfig holds multiorch service configuration
type MultiorchConfig struct {
	Path                      string `yaml:"path"`
	HttpPort                  int    `yaml:"http-port"`
	GrpcPort                  int    `yaml:"grpc-port"`
	LogLevel                  string `yaml:"log-level"`
	PoolerHealthCheckInterval string `yaml:"pooler-health-check-interval,omitempty"`
	RecoveryCycleInterval     string `yaml:"recovery-cycle-interval,omitempty"`
}

// MultiadminConfig holds multiadmin service configuration
type MultiadminConfig struct {
	Path     string `yaml:"path"`
	HttpPort int    `yaml:"http-port"`
	GrpcPort int    `yaml:"grpc-port"`
	LogLevel string `yaml:"log-level"`
}

// PgctldConfig holds pgctld service configuration
type PgctldConfig struct {
	Path              string `yaml:"path"`
	HttpPort          int    `yaml:"http-port"`           // HTTP port for health endpoints
	PoolerDir         string `yaml:"pooler-dir"`          // Base directory for this pgctld instance
	GrpcPort          int    `yaml:"grpc-port"`           // gRPC port for pgctld server
	GRPCSocketFile    string `yaml:"grpc-socket-file"`    // Unix socket file path for gRPC
	PgPort            int    `yaml:"pg-port"`             // PostgreSQL port
	PgDatabase        string `yaml:"pg-database"`         // PostgreSQL database name
	PgUser            string `yaml:"pg-user"`             // PostgreSQL username
	PgPassword        string `yaml:"pg-password"`         // PostgreSQL password (default: "postgres")
	Timeout           int    `yaml:"timeout"`             // Operation timeout in seconds
	LogLevel          string `yaml:"log-level"`           // Log level
	PgBackRestPort    int    `yaml:"pgbackrest-port"`     // pgBackRest TLS server port
	PgBackRestCertDir string `yaml:"pgbackrest-cert-dir"` // pgBackRest TLS certificate directory
}

// LoadConfig loads the provisioner-specific configuration from the given config paths
func (p *localProvisioner) LoadConfig(configPaths []string) error {
	_ = "STUB: not implemented"
	// Try to find the config file in the provided paths
	return nil
}

// Parse the full config file

// Validate that this is for the local provisioner

// Convert the provisioner-config section to our typed config

// buildBackupConfig creates a BackupConfig from flag values
// buildBackupConfig creates a BackupConfig from flag values
func buildBackupConfig(backupConfig map[string]string, baseDir string) BackupConfig {
	_ = "STUB: not implemented"
	return *new(BackupConfig)
}

// default

// DefaultConfig returns the default configuration for the local provisioner
func (p *localProvisioner) DefaultConfig(configPaths []string, backupConfig map[string]string) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

// Generate service IDs for each cell using the same method as topo components

// Create typed configuration with defaults

// Same as pgctld for this zone

// Convert to map[string]any via YAML marshaling to preserve struct ordering

// Fallback to empty config if marshaling fails

// Fallback to empty config if unmarshaling fails

// getServiceConfig gets the configuration for a specific service (global services only)
func (p *localProvisioner) getServiceConfig(service string) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

// Return empty config if not found

// getCellServiceConfig gets the configuration for a specific service in a specific cell
func (p *localProvisioner) getCellServiceConfig(cellName, service string) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
