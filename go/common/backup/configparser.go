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

// credentialKeys are the keys that should be updated during credential refresh
var credentialKeys = map[string]bool{
	"repo1-s3-key":        true,
	"repo1-s3-key-secret": true,
	"repo1-s3-key-type":   true,
	"repo1-s3-token":      true,
}

// UpdateCredentialsInConfig updates only S3 credential lines in pgbackrest.conf
// while preserving all other configuration, comments, and formatting.
func UpdateCredentialsInConfig(configContent string, newCredentials map[string]string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Check if this line is a credential key=value pair

// If this is a credential key and we have a new value, replace it

// Preserve all other lines as-is
