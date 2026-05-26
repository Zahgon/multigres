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

package utils

// hasPostgreSQLBinaries checks if required PostgreSQL binaries are available
func HasPostgreSQLBinaries() bool { _ = "STUB: not implemented"; return false }

// ShouldSkipRealPostgres returns true if tests should skip real PostgreSQL tests.
// This happens when running short tests AND PostgreSQL binaries are not available.
func ShouldSkipRealPostgres() bool { _ = "STUB: not implemented"; return false }
