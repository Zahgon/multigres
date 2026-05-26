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

package manager

import (
	"regexp"

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	multipoolermanagerdata "github.com/multigres/multigres/go/pb/multipoolermanagerdata"
)

// Regex for parsing synchronous_standby_names format: METHOD NUM (member1, member2, ...)
// Case-insensitive to accept both "FIRST"/"ANY" and "first"/"any"
var syncStandbyNamesRegex = regexp.MustCompile(`(?i)^(FIRST|ANY)\s+(\d+)\s*\(([^)]*)\)$`)

// SyncStandbyConfig represents a parsed synchronous_standby_names configuration
type SyncStandbyConfig struct {
	Method     multipoolermanagerdata.SynchronousMethod // FIRST or ANY
	NumSync    int32                                    // Number of synchronous standbys
	StandbyIDs []*clustermetadatapb.ID                  // List of standby IDs
}

// parseApplicationName parses a postgres replication application_name back
// into a clustermetadata ID. Application names are always set by multipoolers
// (postgres replication clients), so Component is always MULTIPOOLER.
// Format: {cell}_{name}
// Example: "us-west_replica-1" -> ID{Component: MULTIPOOLER, Cell: "us-west", Name: "replica-1"}
//
// For decoding non-pooler IDs that happen to share the cell_name encoding
// (e.g. coordinator_id in rule_history), use topoclient.SplitClusterID directly.
func parseApplicationName(appName string) (*clustermetadatapb.ID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// parseSynchronousStandbyNames parses a PostgreSQL synchronous_standby_names string
// Examples:
//   - "FIRST 2 ("cell_replica1", "cell_replica2", "cell_replica3")"
//   - "ANY 1 ("cell_replica1", "cell_replica2")"
//   - "*" (wildcard - all connected standbys)
//   - "" (empty - no synchronous replication)
func parseSynchronousStandbyNames(value string) (*SyncStandbyConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Handle empty case

// Handle wildcard case - not supported in Multigres context

// Parse format: METHOD NUM (member1, member2, ...)
// Note: this regex assumes standby_names are being controlled by multigres
// and will have the format we expect (i.e cell_name). We are not validating
// for this format here.

// Normalize to uppercase

// Convert string method to enum

// Parse member list

// Split by comma and clean up each member

// Remove surrounding quotes if present

// Parse application name back to ID

// parseAndRedactPrimaryConnInfo parses a PostgreSQL primary_conninfo connection string into structured fields
// Example input: "host=localhost port=5432 user=postgres application_name=cell_name"
// Returns a PrimaryConnInfo message with parsed fields, or an error if parsing fails
// Note: Passwords are redacted in the raw field for security
func parseAndRedactPrimaryConnInfo(connInfoStr string) (*multipoolermanagerdata.PrimaryConnInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Simple space-based parsing of key=value pairs

// Not a key=value pair - parsing failed

// Redact sensitive fields in the raw string

// Parse specific fields we care about

// Set the redacted raw string
