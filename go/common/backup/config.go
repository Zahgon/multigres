// Copyright 2026 Supabase, Inc.
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

package backup

import (
	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
)

// Config provides a unified interface for backup operations
type Config struct {
	proto *clustermetadatapb.BackupLocation
}

// NewConfig creates a Config from a BackupLocation proto
func NewConfig(loc *clustermetadatapb.BackupLocation) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Type returns the backup location type for logging/metrics
func (c *Config) Type() string { _ = "STUB: not implemented"; return "" }

// FullPath returns the complete backup path for a database/tablegroup/shard
func (c *Config) FullPath(database, tableGroup, shard string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// filesystemFullPath builds a filesystem backup path
func filesystemFullPath(basePath, database, tableGroup, shard string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// s3FullPath builds an S3 backup path
func s3FullPath(s3 *clustermetadatapb.S3Backup, database, tableGroup, shard string) (string, error) {
	_ = "STUB: not implemented"
	// Start with bucket
	return "", nil
}

// Add prefix if set

// Add database/tablegroup/shard

// UsesEnvCredentials returns true if S3 backup uses environment credentials
func (c *Config) UsesEnvCredentials() bool { _ = "STUB: not implemented"; return false }

// PgBackRestCredentials returns credentials for pgBackRest from environment variables.
// Returns nil for non-S3 backups or when UseEnvCredentials is false.
// Returns an error if UseEnvCredentials is true but required env vars are missing.
func (c *Config) PgBackRestCredentials() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Only include session token if it's set

// GetS3Config returns the S3 configuration, or nil if not using S3 backups
func (c *Config) GetS3Config() *clustermetadatapb.S3Backup { _ = "STUB: not implemented"; return nil }

// PgBackRestConfig returns pgBackRest-specific configuration
func (c *Config) PgBackRestConfig(stanzaName string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// S3-specific performance and efficiency settings

// Set credential type based on configuration

// When using env credentials, they should be in a separate file
// Do not include them in the main config

// Use web-id for IRSA, since pgBackRest's "auto" only checks EC2 metadata.

// Fall back to auto for EC2 instance metadata

// AWS S3 - generate standard endpoint

// Use virtual-hosted style (AWS default)

// TLS verification defaults to 'y', no need to set explicitly

// Custom endpoint (s3mock, etc.)

// Repo path includes prefix if set

// DefaultRetentionConfig returns the default pgBackRest retention settings.
// These are placed in the [global] section of pgbackrest.conf and apply to
// all backend types (filesystem and S3).
func DefaultRetentionConfig() map[string]string { _ = "STUB: not implemented"; return nil }

// validate checks that the backup location is properly configured
func validate(loc *clustermetadatapb.BackupLocation) error { _ = "STUB: not implemented"; return nil }
