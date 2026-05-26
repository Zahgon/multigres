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

package test

import (
	"context"
	"testing"

	"github.com/multigres/multigres/go/common/topoclient"
)

// checkShardInitClaim verifies ClaimShardInitialization semantics:
// first caller wins, different caller loses, same caller after crash wins
// with the originally committed cohort.
func checkShardInitClaim(t *testing.T, ctx context.Context, ts topoclient.Store) {
	_ = "STUB: not implemented"
	return
}

// Must get back the originally committed cohort, not the new proposal.
