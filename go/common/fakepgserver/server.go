// Copyright 2025 Supabase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package fakepgserver provides a fake PostgreSQL server for testing.
// It speaks the PostgreSQL wire protocol and returns pre-configured results.
// The API mirrors go/tools/fakepgdb for consistency.
package fakepgserver

import (
	"context"
	"regexp"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/multigres/multigres/go/common/pgprotocol/client"
	"github.com/multigres/multigres/go/common/pgprotocol/server"
	"github.com/multigres/multigres/go/common/sqltypes"
)

// Server is a fake PostgreSQL server for testing.
// All methods are thread-safe.
type Server struct {
	// t is our testing.TB instance.
	t testing.TB

	// listener is the PostgreSQL protocol listener.
	listener *server.Listener

	// address is the server's listening address (host:port).
	address string

	// name is the name of this server (for debugging).
	name string

	// orderMatters controls whether query order matters.
	orderMatters atomic.Bool

	// mu protects all the following fields.
	mu sync.Mutex

	// data maps tolower(query) to a result.
	data map[string]*sqltypes.Result

	// rejectedData maps tolower(query) to an error.
	rejectedData map[string]error

	// patternData is a map of regexp queries to results.
	patternData map[string]exprResult

	// patternCalled keeps track of how many times each pattern was matched.
	patternCalled map[string]int

	// queryCalled keeps track of how many times a query was called.
	queryCalled map[string]int

	// querylog keeps track of all called queries.
	querylog []string

	// expectedExecuteFetch is the array of expected queries (for ordered mode).
	expectedExecuteFetch []ExpectedExecuteFetch

	// expectedExecuteFetchIndex is the current index of the query.
	expectedExecuteFetchIndex int

	// neverFail makes unmatched queries return empty results instead of errors.
	neverFail atomic.Bool

	// queryPatternUserCallback stores optional callbacks when a query with a pattern is called.
	queryPatternUserCallback map[*regexp.Regexp]func(string)

	// credentialProvider is the optional inner CredentialProvider used by
	// the listener's auth path. Tests that exercise replication startup
	// (replication=true / replication=database) must install one whose
	// IsReplicationRole is true; without it, verifyReplicationRole fails
	// closed for the trust-auth path.
	credentialProvider server.CredentialProvider

	// lastReplicationMode is the parsed `replication` startup parameter
	// of the most recently established connection. Recorded via the
	// ConnectionEstablishedHandler hook.
	lastReplicationMode server.ReplicationMode

	// rejectNextReplStartup, when armed, causes the next replication-mode
	// startup (replication=true / replication=database) to be rejected with
	// FATAL 42501 as if the role lacked the REPLICATION attribute. One-shot:
	// consumed via CompareAndSwap so it auto-resets after the first matching
	// startup. Non-replication startups are unaffected. Used by tests to
	// assert dial-failure cleanup paths.
	//
	// We piggyback on verifyReplicationRole's GetCredentials path (see the
	// GetCredentials method below) rather than introducing a new
	// startup-rejection hook so the fake's failure modes stay aligned with
	// production rolreplication-rejection plumbing — tests exercise the same
	// post-auth gate real postgres uses, not a parallel error path.
	//
	// Listener-config qualifier: this scoping works because GetCredentials is
	// reached only from verifyReplicationRole under fakepgserver's hardcoded
	// trust auth. If fakepgserver ever grows SCRAM (which calls
	// GetCredentials for every login, not just replication startups), this
	// toggle would broaden to non-replication startups and the logic in
	// GetCredentials would need to gate on replication mode explicitly.
	rejectNextReplStartup atomic.Bool
}

type exprResult struct {
	queryPattern string
	expr         *regexp.Regexp
	result       *sqltypes.Result
	err          string
}

// trustAllProvider implements server.TrustAuthProvider for testing.
// It allows all connections without password authentication,
// simulating Unix socket trust authentication.
type trustAllProvider struct{}

// AllowTrustAuth always returns true, allowing all connections without password.
func (p *trustAllProvider) AllowTrustAuth(_ context.Context, _, _ string) bool {
	_ = "STUB: not implemented"

	// GetCredentials dispatches to the test-installed CredentialProvider when one
	// is set. Trust-auth replication startups invoke this path inside
	// verifyReplicationRole (server/startup.go) when the auth-time SCRAM lookup
	// was skipped. Without an installed provider, replication startups are
	// rejected — which is the right default for a fake server.
	return false
}

func (s *Server) GetCredentials(ctx context.Context, user, database string) (*server.Credentials, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// One-shot rejection toggle: clear IsReplicationRole so the post-auth
// gate in verifyReplicationRole rejects this connection with a FATAL
// 42501, matching production rolreplication-rejection plumbing rather
// than a parallel error path. Trust-auth replication startups are the
// only path that reaches here under fakepgserver's current listener
// config, so non-replication sessions are unaffected. See the
// rejectNextReplStartup field comment for the SCRAM caveat.

// SetRejectNextReplicationStartup arms a one-shot toggle that causes the
// next replication-mode startup (replication=true / replication=database)
// to be rejected at the rolreplication gate. The toggle auto-resets after
// firing, so subsequent replication startups succeed again. Pass false to
// disarm explicitly. Tests use this to exercise dial-failure cleanup
// without dropping the listener.
func (s *Server) SetRejectNextReplicationStartup(reject bool) { _ = "STUB: not implemented"; return }

// SetCredentialProvider installs a server.CredentialProvider that the
// listener's auth path will consult. Tests that exercise replication startup
// (replication=true / replication=database) must call this with a provider
// whose Credentials carry IsReplicationRole=true, otherwise the post-auth
// gate rejects the connection.
func (s *Server) SetCredentialProvider(provider server.CredentialProvider) *Server {
	_ = "STUB: not implemented"
	return nil
}

// LastReplicationMode returns the replication mode of the most recently
// established connection. Returns ReplicationOff if no connection has
// completed startup, or the last connection was a normal (non-replication)
// session.
//
// The `replication` startup parameter is stripped from GetStartupParams()
// after parsing, so tests must use this accessor to verify wire-level
// replication mode.
func (s *Server) LastReplicationMode() server.ReplicationMode {
	_ = "STUB: not implemented"
	return *new(server.ReplicationMode)
}

// recordReplicationMode is invoked by the handler's ConnectionEstablished
// hook to stash the most recent connection's mode for test assertions.
func (s *Server) recordReplicationMode(mode server.ReplicationMode) {
	_ = "STUB: not implemented"
	return
}

// ExpectedExecuteFetch defines for an expected query the to be faked output.
// It is used for ordered expected output.
type ExpectedExecuteFetch struct {
	Query       string
	QueryResult *sqltypes.Result
	Error       error

	// AfterCallbackError, when set alongside QueryResult, causes the handler
	// to first deliver the result via callback and then return this error.
	// This simulates mid-stream failures where partial data is delivered
	// before the connection dies (e.g., rows sent then FATAL 57P01).
	AfterCallbackError error
}

// New creates a new fake PostgreSQL server for testing.
// The server listens on a random available TCP port.
func New(t testing.TB) *Server { _ = "STUB: not implemented"; return nil }

// Create the handler.

// Create listener on random port with trust auth (simulates Unix socket
// trust auth). The server itself is wired as the CredentialProvider so
// tests can install one at runtime via SetCredentialProvider — required
// for replication-mode startups, which post-auth gate on the role's
// IsReplicationRole flag.

// Random available port.

// Get the actual address.

// Start serving in background.

// Don't log errors if the listener was closed intentionally.

// Name returns the name of the server.
func (s *Server) Name() string { _ = "STUB: not implemented"; return "" }

// SetName sets the name of the server.
func (s *Server) SetName(name string) *Server { _ = "STUB: not implemented"; return nil }

// Address returns the server's listening address.
func (s *Server) Address() string {
	_ = "STUB: not implemented"

	// ClientConfig returns a client.Config for connecting to this server.
	// No password is needed since fakepgserver uses trust authentication.
	return ""
}

func (s *Server) ClientConfig() *client.Config { _ = "STUB: not implemented"; return nil }

// Close closes the server and stops accepting connections.
func (s *Server) Close() { _ = "STUB: not implemented"; return }

// CloseListener closes only the TCP listener, preventing new connections
// while keeping existing connections alive. Use this for testing scenarios
// where the initial connection should work but reconnect attempts should fail.
func (s *Server) CloseListener() { _ = "STUB: not implemented"; return }

// OrderMatters sets the orderMatters flag.
func (s *Server) OrderMatters() { _ = "STUB: not implemented"; return }

//
// Methods to add expected queries and results.
//

// AddQuery adds a query and its expected result.
func (s *Server) AddQuery(q string, result *sqltypes.Result) { _ = "STUB: not implemented"; return }

// AddQueryPattern adds an expected result for a set of queries.
// These patterns are checked if no exact matches from AddQuery() are found.
// This function forces the addition of begin/end anchors (^$) and turns on
// case-insensitive matching mode.
func (s *Server) AddQueryPattern(queryPattern string, result *sqltypes.Result) {
	_ = "STUB: not implemented"
	return
}

// RemoveQueryPattern removes a query pattern that was previously added.
func (s *Server) RemoveQueryPattern(queryPattern string) { _ = "STUB: not implemented"; return }

// RejectQueryPattern allows a query pattern to be rejected with an error.
func (s *Server) RejectQueryPattern(queryPattern, errMsg string) { _ = "STUB: not implemented"; return }

// ClearQueryPattern removes all query patterns set up.
func (s *Server) ClearQueryPattern() { _ = "STUB: not implemented"; return }

// AddQueryPatternWithCallback is similar to AddQueryPattern: in addition it calls the provided callback function.
func (s *Server) AddQueryPatternWithCallback(queryPattern string, result *sqltypes.Result, callback func(string)) {
	_ = "STUB: not implemented"
	return
}

// DeleteQuery deletes query from the fake server.
func (s *Server) DeleteQuery(query string) { _ = "STUB: not implemented"; return }

// DeleteAllQueries deletes all expected queries from the fake server.
func (s *Server) DeleteAllQueries() { _ = "STUB: not implemented"; return }

// AddRejectedQuery adds a query which will be rejected at execution time.
func (s *Server) AddRejectedQuery(query string, err error) { _ = "STUB: not implemented"; return }

// DeleteRejectedQuery deletes query from the fake server.
func (s *Server) DeleteRejectedQuery(query string) { _ = "STUB: not implemented"; return }

// GetQueryCalledNum returns how many times the server executed a certain query.
func (s *Server) GetQueryCalledNum(query string) int { _ = "STUB: not implemented"; return 0 }

// QueryLog returns the query log as a semicolon separated string.
func (s *Server) QueryLog() string { _ = "STUB: not implemented"; return "" }

// ResetQueryLog resets the query log.
func (s *Server) ResetQueryLog() { _ = "STUB: not implemented"; return }

//
// Methods for ordered expected queries.
//

// AddExpectedExecuteFetch appends an ExpectedExecuteFetch to the end.
func (s *Server) AddExpectedExecuteFetch(entry ExpectedExecuteFetch) {
	_ = "STUB: not implemented"
	return
}

// AddExpectedQuery adds a single query with no result.
func (s *Server) AddExpectedQuery(q string, err error) { _ = "STUB: not implemented"; return }

// DeleteAllEntries removes all ordered entries.
func (s *Server) DeleteAllEntries() { _ = "STUB: not implemented"; return }

// VerifyAllExecutedOrFail checks that all expected queries were actually executed.
func (s *Server) VerifyAllExecutedOrFail() { _ = "STUB: not implemented"; return }

// GetPatternCalledNum returns how many times a pattern was matched.
func (s *Server) GetPatternCalledNum(pattern string) int { _ = "STUB: not implemented"; return 0 }

// VerifyAllPatternsUsedOrFail checks that all registered patterns were matched at least once.
func (s *Server) VerifyAllPatternsUsedOrFail() { _ = "STUB: not implemented"; return }

// SetNeverFail makes unmatched queries return empty results instead of errors.
func (s *Server) SetNeverFail(neverFail bool) { _ = "STUB: not implemented"; return }

// handleQuery handles a query and returns the result.
// This is called by the handler.
func (s *Server) handleQuery(q string) (*sqltypes.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if we should reject it.

// Check explicit queries from AddQuery().

// Check query patterns from AddQueryPattern().

// Nothing matched.

func (s *Server) handleQueryOrdered(q string) (*sqltypes.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MakeResult creates a simple sqltypes.Result from column names and row values.
// This is a convenience function for tests. All values are converted to text format.
func MakeResult(columns []string, rows [][]any) *sqltypes.Result {
	_ = "STUB: not implemented"
	return nil
}

// TEXT type OID
// Variable length

// NULL
