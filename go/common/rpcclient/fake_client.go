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

package rpcclient

import (
	"context"
	"sync"
	"time"

	clustermetadatapb "github.com/multigres/multigres/go/pb/clustermetadata"
	consensusdatapb "github.com/multigres/multigres/go/pb/consensusdata"
	multipoolermanagerdatapb "github.com/multigres/multigres/go/pb/multipoolermanagerdata"
)

// ResponseWithDelay wraps a response with an optional delay for testing timeouts.
// If Delay is set, the fake client will sleep for that duration before returning.
// If the context is cancelled during the delay, the context error is returned.
type ResponseWithDelay[T any] struct {
	Response T
	Delay    time.Duration
}

// FakeClient implements MultiPoolerClient for testing purposes.
// It provides a simple in-memory implementation that can be configured
// to return specific responses or errors for testing.
//
// This is useful for:
//   - Unit testing multiorch coordinator logic without real gRPC servers
//   - Integration tests that need predictable multipooler responses
//   - Simulating error conditions and edge cases
//
// Usage:
//
//	fake := rpcclient.NewFakeClient()
//	fake.SetStatusResponse("pooler-1", &multipoolermanagerdatapb.StatusResponse{...})
//	// Or with a delay to simulate slow/unresponsive poolers:
//	fake.SetStatusResponseWithDelay("pooler-2", &multipoolermanagerdatapb.StatusResponse{...}, 5*time.Second)
//	resp, err := fake.Status(ctx, pooler, &multipoolermanagerdatapb.StatusRequest{})
type FakeClient struct {
	mu sync.RWMutex

	// Consensus service responses - keyed by pooler ID
	BeginTermResponses       map[string]*consensusdatapb.BeginTermResponse
	RecruitResponses         map[string]*consensusdatapb.RecruitResponse
	ProposeResponses         map[string]*consensusdatapb.ProposeResponse
	SetTermPrimaryResponses  map[string]*consensusdatapb.SetTermPrimaryResponse
	ConsensusStatusResponses map[string]*consensusdatapb.StatusResponse
	EmergencyDemoteResponses map[string]*multipoolermanagerdatapb.EmergencyDemoteResponse
	PromoteResponses         map[string]*multipoolermanagerdatapb.PromoteResponse

	// Manager service responses - keyed by pooler ID
	WaitForLSNResponses          map[string]*multipoolermanagerdatapb.WaitForLSNResponse
	SetPrimaryConnInfoResponses  map[string]*multipoolermanagerdatapb.SetPrimaryConnInfoResponse
	StartReplicationResponses    map[string]*multipoolermanagerdatapb.StartReplicationResponse
	StopReplicationResponses     map[string]*multipoolermanagerdatapb.StopReplicationResponse
	StatusResponses              map[string]*ResponseWithDelay[*multipoolermanagerdatapb.StatusResponse]
	UpdateConsensusRuleResponses map[string]*multipoolermanagerdatapb.UpdateConsensusRuleResponse
	// LastUpdateConsensusRuleRequest captures the most recent UpdateConsensusRule
	// request payload for tests that need to assert on operation/IDs.
	LastUpdateConsensusRuleRequest      *multipoolermanagerdatapb.UpdateConsensusRuleRequest
	BackupResponses                     map[string]*multipoolermanagerdatapb.BackupResponse
	RestoreFromBackupResponses          map[string]*multipoolermanagerdatapb.RestoreFromBackupResponse
	GetBackupsResponses                 map[string]*multipoolermanagerdatapb.GetBackupsResponse
	GetBackupByJobIdResponses           map[string]*multipoolermanagerdatapb.GetBackupByJobIdResponse
	RewindToSourceResponses             map[string]*multipoolermanagerdatapb.RewindToSourceResponse
	SetPostgresRestartsEnabledResponses map[string]*multipoolermanagerdatapb.SetPostgresRestartsEnabledResponse

	// Errors to return - keyed by pooler ID
	Errors map[string]error

	// CallLog tracks which methods were called for verification in tests
	CallLog []string

	// Request tracking for verification in tests
	PromoteRequests        map[string]*multipoolermanagerdatapb.PromoteRequest
	ProposeRequests        map[string]*consensusdatapb.ProposeRequest
	SetTermPrimaryRequests map[string]*consensusdatapb.SetTermPrimaryRequest

	// OnManagerHealthStream, if set, is called after each FakeManagerHealthStream
	// is created. Tests use this to capture the stream and inject snapshots.
	OnManagerHealthStream func(poolerID string, stream *FakeManagerHealthStream)
}

// NewFakeClient creates a new FakeClient with empty response maps.
func NewFakeClient() *FakeClient { _ = "STUB: not implemented"; return nil }

// Helper methods

func (f *FakeClient) getPoolerID(pooler *clustermetadatapb.MultiPooler) string {
	_ = "STUB: not implemented"
	return ""
}

func (f *FakeClient) logCall(method string, poolerID string) { _ = "STUB: not implemented"; return }

// GetCallLog returns a copy of the call log in a thread-safe manner.
func (f *FakeClient) GetCallLog() []string { _ = "STUB: not implemented"; return nil }

// ResetCallLog clears the call log in a thread-safe manner.
func (f *FakeClient) ResetCallLog() { _ = "STUB: not implemented"; return }

func (f *FakeClient) checkError(poolerID string) error { _ = "STUB: not implemented"; return nil }

// SetStatusResponse sets a Status response for a pooler with no delay.
func (f *FakeClient) SetStatusResponse(poolerID string, resp *multipoolermanagerdatapb.StatusResponse) {
	_ = "STUB: not implemented"
	return
}

// SetStatusResponseWithDelay sets a Status response for a pooler with a delay.
// The delay simulates a slow or unresponsive pooler for testing timeout behavior.
func (f *FakeClient) SetStatusResponseWithDelay(poolerID string, resp *multipoolermanagerdatapb.StatusResponse, delay time.Duration) {
	_ = "STUB: not implemented"
	return
}

// SetPostgresRestartsEnabledResponse sets a SetPostgresRestartsEnabled response for a pooler.
func (f *FakeClient) SetPostgresRestartsEnabledResponse(poolerID string, resp *multipoolermanagerdatapb.SetPostgresRestartsEnabledResponse) {
	_ = "STUB: not implemented"
	return
}

//
// Consensus Service Methods
//

func (f *FakeClient) BeginTerm(ctx context.Context, pooler *clustermetadatapb.MultiPooler, request *consensusdatapb.BeginTermRequest) (*consensusdatapb.BeginTermResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FakeClient) Recruit(ctx context.Context, pooler *clustermetadatapb.MultiPooler, request *consensusdatapb.RecruitRequest) (*consensusdatapb.RecruitResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Stamp the request's TermRevocation onto the ConsensusStatus. The real
// multipooler stores the accepted revocation in its state and returns it in
// ConsensusStatus; filterByRevocation in BuildSafeProposal matches on it.

func (f *FakeClient) Propose(ctx context.Context, pooler *clustermetadatapb.MultiPooler, request *consensusdatapb.ProposeRequest) (*consensusdatapb.ProposeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FakeClient) SetTermPrimary(ctx context.Context, pooler *clustermetadatapb.MultiPooler, request *consensusdatapb.SetTermPrimaryRequest) (*consensusdatapb.SetTermPrimaryResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FakeClient) ConsensusStatus(ctx context.Context, pooler *clustermetadatapb.MultiPooler, request *consensusdatapb.StatusRequest) (*consensusdatapb.StatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FakeClient) EmergencyDemote(ctx context.Context, pooler *clustermetadatapb.MultiPooler, request *multipoolermanagerdatapb.EmergencyDemoteRequest) (*multipoolermanagerdatapb.EmergencyDemoteResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FakeClient) DemoteStalePrimary(ctx context.Context, pooler *clustermetadatapb.MultiPooler, request *multipoolermanagerdatapb.DemoteStalePrimaryRequest) (*multipoolermanagerdatapb.DemoteStalePrimaryResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FakeClient) Promote(ctx context.Context, pooler *clustermetadatapb.MultiPooler, request *multipoolermanagerdatapb.PromoteRequest) (*multipoolermanagerdatapb.PromoteResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Record the request for test verification

func (f *FakeClient) UpdateConsensusRule(ctx context.Context, pooler *clustermetadatapb.MultiPooler, request *multipoolermanagerdatapb.UpdateConsensusRuleRequest) (*multipoolermanagerdatapb.UpdateConsensusRuleResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//
// Manager Service Methods - Status and Monitoring
//

func (f *FakeClient) Status(ctx context.Context, pooler *clustermetadatapb.MultiPooler, request *multipoolermanagerdatapb.StatusRequest) (*multipoolermanagerdatapb.StatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//
// Manager Service Methods - Replication
//

func (f *FakeClient) WaitForLSN(ctx context.Context, pooler *clustermetadatapb.MultiPooler, request *multipoolermanagerdatapb.WaitForLSNRequest) (*multipoolermanagerdatapb.WaitForLSNResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FakeClient) SetPrimaryConnInfo(ctx context.Context, pooler *clustermetadatapb.MultiPooler, request *multipoolermanagerdatapb.SetPrimaryConnInfoRequest) (*multipoolermanagerdatapb.SetPrimaryConnInfoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FakeClient) StartReplication(ctx context.Context, pooler *clustermetadatapb.MultiPooler, request *multipoolermanagerdatapb.StartReplicationRequest) (*multipoolermanagerdatapb.StartReplicationResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FakeClient) StopReplication(ctx context.Context, pooler *clustermetadatapb.MultiPooler, request *multipoolermanagerdatapb.StopReplicationRequest) (*multipoolermanagerdatapb.StopReplicationResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//
// Manager Service Methods - Backup and Restore
//

func (f *FakeClient) Backup(ctx context.Context, pooler *clustermetadatapb.MultiPooler, request *multipoolermanagerdatapb.BackupRequest) (*multipoolermanagerdatapb.BackupResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FakeClient) RestoreFromBackup(ctx context.Context, pooler *clustermetadatapb.MultiPooler, request *multipoolermanagerdatapb.RestoreFromBackupRequest) (*multipoolermanagerdatapb.RestoreFromBackupResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FakeClient) GetBackups(ctx context.Context, pooler *clustermetadatapb.MultiPooler, request *multipoolermanagerdatapb.GetBackupsRequest) (*multipoolermanagerdatapb.GetBackupsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FakeClient) GetBackupByJobId(ctx context.Context, pooler *clustermetadatapb.MultiPooler, request *multipoolermanagerdatapb.GetBackupByJobIdRequest) (*multipoolermanagerdatapb.GetBackupByJobIdResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FakeClient) ExpireBackups(ctx context.Context, pooler *clustermetadatapb.MultiPooler, request *multipoolermanagerdatapb.ExpireBackupsRequest) (*multipoolermanagerdatapb.ExpireBackupsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//
// Manager Service Methods - Timeline Repair
//

func (f *FakeClient) RewindToSource(ctx context.Context, pooler *clustermetadatapb.MultiPooler, req *multipoolermanagerdatapb.RewindToSourceRequest) (*multipoolermanagerdatapb.RewindToSourceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//
// Manager Service Methods - PostgreSQL Restart Control
//

func (f *FakeClient) SetPostgresRestartsEnabled(ctx context.Context, pooler *clustermetadatapb.MultiPooler, req *multipoolermanagerdatapb.SetPostgresRestartsEnabledRequest) (*multipoolermanagerdatapb.SetPostgresRestartsEnabledResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//
// Manager Service Methods - Health Streaming
//

// FakeManagerHealthStream implements ManagerHealthStream for testing.
// Recv blocks until the context is cancelled or a response is injected via Ch.
// Sent messages (init, poll) are recorded on the Sent channel.
type FakeManagerHealthStream struct {
	ctx  context.Context
	Ch   chan *multipoolermanagerdatapb.ManagerHealthStreamResponse
	Sent chan *multipoolermanagerdatapb.ManagerHealthStreamClientMessage
}

// Recv blocks until a response is available or the context is cancelled.
func (f *FakeManagerHealthStream) Recv() (*multipoolermanagerdatapb.ManagerHealthStreamResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Send records the outgoing message on the Sent channel.
// Non-blocking: if Sent is full the message is dropped (tests should drain it).
func (f *FakeManagerHealthStream) Send(msg *multipoolermanagerdatapb.ManagerHealthStreamClientMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// ManagerHealthStream returns a FakeManagerHealthStream. Tests inject snapshots
// by sending on stream.Ch or close it to simulate disconnection. Outgoing
// messages (init/poll) are readable from stream.Sent.
func (f *FakeClient) ManagerHealthStream(ctx context.Context, pooler *clustermetadatapb.MultiPooler) (ManagerHealthStream, error) {
	_ = "STUB: not implemented"
	return *new(ManagerHealthStream), nil
}

//
// Connection Management Methods
//

func (f *FakeClient) Close() {
	_ = "STUB: not implemented"
	// No-op for fake client
	return
}

func (f *FakeClient) CloseTablet(pooler *clustermetadatapb.MultiPooler) {
	_ = "STUB: not implemented"
	return
}
