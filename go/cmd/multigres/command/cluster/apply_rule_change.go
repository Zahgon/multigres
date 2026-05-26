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

package cluster

import (
	"time"

	"github.com/spf13/cobra"

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	multiadminpb "github.com/multigres/multigres/go/pb/multiadmin"
	"github.com/multigres/multigres/go/tools/viperutil"
)

type applyRuleChangeCmd struct {
	database              viperutil.Value[string]
	tableGroup            viperutil.Value[string]
	shard                 viperutil.Value[string]
	leader                viperutil.Value[string]
	cohort                viperutil.Value[[]string]
	durability            viperutil.Value[string]
	outgoingRuleTerm      viperutil.Value[int64]
	outgoingLeaderSubterm viperutil.Value[int64]
	frozenLSN             viperutil.Value[string]
	unsafeDeriveCert      viperutil.Value[bool]
	reason                viperutil.Value[string]
	yes                   viperutil.Value[bool]
	timeout               viperutil.Value[time.Duration]
}

// AddApplyRuleChangeCommand registers the apply-rule-change subcommand.
//
// Used for both initial leader appointment (cohort has no committed rule yet
// — pass zero outgoing-rule-term + --frozen-lsn=0/0) and stuck-quorum
// recovery (cohort has a rule but quorum is unreachable — pass the
// outgoing rule's term and frozen LSN, or use --unsafe-derive-cert-from-reachable
// to have multiadmin probe the proposed cohort and derive them).
// newApplyRuleChangeCmd constructs the applyRuleChangeCmd struct with all its
// viperutil.Value flag handles. Shared by the public command constructor and
// by tests that need a struct whose .Set methods feed into buildRequest's
// .Get reads.
func newApplyRuleChangeCmd() *applyRuleChangeCmd { _ = "STUB: not implemented"; return nil }

func AddApplyRuleChangeCommand(clusterCmd *cobra.Command) { _ = "STUB: not implemented"; return }

// buildRequest validates the configured flags and assembles the
// ApplyCertifiedRuleChangeRequest the CLI sends to multiadmin. Extracted from
// run so the flag validation + cert-vs-derive switch can be unit-tested
// without dialing multiadmin or reading stdin.
func (a *applyRuleChangeCmd) buildRequest() (*multiadminpb.ApplyCertifiedRuleChangeRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *applyRuleChangeCmd) run(cmd *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

// parsePoolerID accepts "cell_name" (the canonical encoding used by
// topoclient.ClusterIDString) and returns a fully-qualified MULTIPOOLER ID.
func parsePoolerID(raw string) (*clustermetadatapb.ID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// confirm prints a summary of the operation and prompts the operator to type
// the shard name to proceed. Bypassed by --yes.
func confirm(cmd *cobra.Command, req *multiadminpb.ApplyCertifiedRuleChangeRequest) error {
	_ = "STUB: not implemented"
	return nil
}
