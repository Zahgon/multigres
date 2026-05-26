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

package multiadmin

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	multiadminpb "github.com/multigres/multigres/go/pb/multiadmin"
)

// ApplyCertifiedRuleChange installs a new shard rule using an externally
// certified revocation. Multiadmin is a convenience layer over the multiorch
// RPC: it (1) optionally derives outgoing_rule_number and frozen_lsn by
// probing the proposed cohort, (2) picks an orch and fills in any identity
// or timing fields the caller omitted, and (3) forwards the fully-populated
// request to the chosen multiorch.
func (s *MultiAdminServer) ApplyCertifiedRuleChange(ctx context.Context, req *multiadminpb.ApplyCertifiedRuleChangeRequest) (*multiadminpb.ApplyCertifiedRuleChangeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// pickOrch returns a multiorch to forward the rule change to. It prefers an
// orch in the same cell as the proposed leader (cross-cell traffic during
// Recruit/Propose is wasteful) and falls back to any orch in any other cell.
//
// Per-cell listing errors are tolerated as long as we find at least one
// orch somewhere — they only surface if no orch is found in any cell, in
// which case they are included in the error message.
//
// Until orchs publish their watch_targets in topology, multiadmin accepts
// whatever multiorch it finds and lets the orch reject with NotFound if it
// does not actually watch the shard.
//
// TODO(before launch): the chosen orch might be registered but not
// responding, leaving the caller with a non-recoverable error. Options to
// address: (a) gRPC client-side load balancing that tries each orch in turn,
// or (b) expose an optional orch_id parameter on the request so a caller
// who already knows a healthy orch can pin it. Out of scope for this PR.
func (s *MultiAdminServer) pickOrch(ctx context.Context, leaderID *clustermetadatapb.ID) (*clustermetadatapb.MultiOrch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Try the leader's cell first, then the rest. Per-cell errors are
// logged but not fatal; we only fail if no cell yields an orch.

// dialOrch opens a gRPC connection to a multiorch using its grpc port.
func (s *MultiAdminServer) dialOrch(ctx context.Context, orch *clustermetadatapb.MultiOrch) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// buildCert returns the ExternallyCertifiedRevocation to forward. For an
// explicit cert this is a clone of the caller's input. For unsafe_derive_cert
// it probes the proposed cohort and computes term_revocation.outgoing_rule
// and frozen_lsn from the most-advanced response.
func (s *MultiAdminServer) buildCert(
	ctx context.Context,
	req *multiadminpb.ApplyCertifiedRuleChangeRequest,
	proposedRule *clustermetadatapb.ShardRule,
) (*clustermetadatapb.ExternallyCertifiedRevocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// probeMostAdvanced calls MultiPoolerManager.Status on every proposed cohort
// member and returns the highest (rule_number, lsn) pair observed across the
// reachable subset.
//
// Hard failures:
//   - A cohort member ID that does not resolve in topology — we refuse to
//     silently exclude it.
//   - Insufficient responses: per durabilityPolicy.CheckSufficientRecruitment,
//     the reachable subset must be enough to satisfy the new rule's quorum.
//     If not, the cert we'd derive could be missing the most-advanced node
//     in the proposed cohort.
//
// Soft failures: individual RPC errors are logged and skipped. The operator
// chose unsafe_derive_cert; the reachable subset is what we derive from.
func (s *MultiAdminServer) probeMostAdvanced(
	ctx context.Context,
	cohortMembers []*clustermetadatapb.ID,
	durabilityPolicy *clustermetadatapb.DurabilityPolicy,
) (*clustermetadatapb.RuleNumber, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// Resolve every cohort member up front. A missing pooler is a hard
// failure — proceeding would mean deriving the cert from a strict
// subset of the cohort that the caller did not opt into.

// Require a quorum of the proposed cohort to respond. Without that,
// the derived cert could understate the cohort's most-advanced position.

// Fresh bootstrap: no rule recorded anywhere. Bootstrap convention.

// fillIdentityFields populates any identity / timing fields the caller left
// unset on proposedRule or cert. Fields the caller provided are validated
// for consistency with the chosen orch but otherwise preserved.
func fillIdentityFields(
	proposedRule *clustermetadatapb.ShardRule,
	cert *clustermetadatapb.ExternallyCertifiedRevocation,
	orchID *clustermetadatapb.ID,
	now *timestamppb.Timestamp,
) error {
	_ = "STUB: not implemented"
	return nil
}
