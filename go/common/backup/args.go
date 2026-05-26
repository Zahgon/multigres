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

// keyToFlagPrefix converts an override key to a pgBackRest CLI flag prefix.
// Example: "pg2_host" → "--pg2-host="
// Returns empty string if key is invalid (empty or already in flag format).
func keyToFlagPrefix(key string) string { _ = "STUB: not implemented"; return "" }

// Skip if already in flag format

// Convert underscores to hyphens and add pgBackRest flag formatting

// ApplyPgBackRestOverrides applies override values to pgbackrest arguments.
// If an override key matches an existing arg prefix, it replaces that arg.
// Otherwise, it appends the new arg.
func ApplyPgBackRestOverrides(baseArgs []string, overrides map[string]string) []string {
	_ = "STUB: not implemented"
	return nil
}

// Build map of override keys to pgbackrest flags

// Skip invalid keys

// Replace matching args

// Don't add again

// Append any overrides that didn't replace existing args
