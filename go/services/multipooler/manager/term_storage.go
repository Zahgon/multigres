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
	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
)

// postgresDataDir returns the PostgreSQL data directory path from PGDATA env var
func postgresDataDir() string { _ = "STUB: not implemented"; return "" }

// multigresDataDir returns the multigres-specific subdirectory within PGDATA
func multigresDataDir() string { _ = "STUB: not implemented"; return "" }

// consensusTermPath returns the path to the consensus term file
func (cs *ConsensusState) consensusTermPath() string { _ = "STUB: not implemented"; return "" }

// getRevocation retrieves the current term revocation from disk.
func (cs *ConsensusState) getRevocation() (*clustermetadatapb.TermRevocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if consensus term file exists

// Return empty term if file doesn't exist

// Read the file

// Unmarshal JSON to protobuf

// setRevocation saves the term revocation to disk atomically.
func (cs *ConsensusState) setRevocation(revocation *clustermetadatapb.TermRevocation) error {
	_ = "STUB: not implemented"
	return nil
}

// Marshal protobuf to JSON

// Write to file atomically using a temporary file

// Rename to final path (atomic operation)

// Clean up temp file on error
