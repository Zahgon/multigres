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

package utils

import (
	"testing"

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
)

// TestConfigBuilder builds multigres.yaml configurations for testing
type TestConfigBuilder struct {
	provisioner string
	backupType  string
	localPath   string
	s3Config    *S3TestConfig
}

// S3TestConfig holds S3 backup configuration for tests
type S3TestConfig struct {
	Bucket      string
	Region      string
	Endpoint    string
	UseEnvCreds bool
	KeyPrefix   string
}

// NewTestConfig creates a new test config builder
func NewTestConfig() *TestConfigBuilder { _ = "STUB: not implemented"; return nil }

// WithLocalBackup configures local backup storage
func (b *TestConfigBuilder) WithLocalBackup(path string) *TestConfigBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithS3Backup configures S3 backup storage
func (b *TestConfigBuilder) WithS3Backup(bucket, region string) *TestConfigBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithEndpoint sets the S3 endpoint (for s3mock, etc.)
func (b *TestConfigBuilder) WithEndpoint(endpoint string) *TestConfigBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithEnvCredentials configures whether to use environment credentials
func (b *TestConfigBuilder) WithEnvCredentials(use bool) *TestConfigBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithKeyPrefix sets the S3 key prefix
func (b *TestConfigBuilder) WithKeyPrefix(prefix string) *TestConfigBuilder {
	_ = "STUB: not implemented"
	return nil
}

// buildBackupConfig creates the backup config map
func (b *TestConfigBuilder) buildBackupConfig() map[string]any {
	_ = "STUB: not implemented"
	return nil
}

// WriteToDir writes the config to multigres.yaml in the specified directory
func (b *TestConfigBuilder) WriteToDir(t *testing.T, dir string) { _ = "STUB: not implemented"; return }

// PgBackrestOpts holds options for pgbackrest.conf generation
type PgBackrestOpts struct {
	AccessKey    string
	SecretKey    string
	SessionToken string
	Bucket       string
	Region       string
	Endpoint     string
}

// WritePgBackrestConfig creates a pgbackrest.conf file with test credentials
func WritePgBackrestConfig(t *testing.T, dir string, opts PgBackrestOpts) {
	_ = "STUB: not implemented"
	return
}

// buildPgBackrestConfig generates pgbackrest.conf content
func buildPgBackrestConfig(opts PgBackrestOpts) string { _ = "STUB: not implemented"; return "" }

// FilesystemBackupLocation creates a BackupLocation for filesystem backups.
func FilesystemBackupLocation(path string) *clustermetadatapb.BackupLocation {
	_ = "STUB: not implemented"
	return nil
}

// S3BackupLocation creates a BackupLocation for S3 backups with basic configuration.
// Use S3Option functions to customize the configuration.
func S3BackupLocation(bucket, region string, opts ...S3Option) *clustermetadatapb.BackupLocation {
	_ = "STUB: not implemented"
	return nil
}

// S3Option is a functional option for configuring S3 backup locations.
type S3Option func(*clustermetadatapb.S3Backup)

// WithS3KeyPrefix sets the key prefix for S3 backups.
func WithS3KeyPrefix(prefix string) S3Option { _ = "STUB: not implemented"; return *new(S3Option) }

// WithS3Endpoint sets a custom endpoint for S3 backups (e.g., s3mock).
func WithS3Endpoint(endpoint string) S3Option { _ = "STUB: not implemented"; return *new(S3Option) }

// WithS3EnvCredentials configures S3 to use environment credentials.
func WithS3EnvCredentials() S3Option { _ = "STUB: not implemented"; return *new(S3Option) }
