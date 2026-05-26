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

// Package mock provides mock implementations for testing.
package mock

import (
	"context"
	"regexp"
	"sync"

	"github.com/multigres/multigres/go/common/sqltypes"
	"github.com/multigres/multigres/go/services/multipooler/executor"
)

// QueryService is a mock implementation of executor.InternalQueryService for testing.
type QueryService struct {
	mu       sync.Mutex
	patterns []queryPattern
}

type queryPattern struct {
	pattern     *regexp.Regexp
	result      *sqltypes.Result
	err         error
	callback    func(string)
	ctxCallback func(context.Context, string) // callback that receives context for blocking tests
	consumeOnce bool                          // if true, pattern is removed after first match
}

// NewQueryService creates a new mock query service for testing.
func NewQueryService() *QueryService { _ = "STUB: not implemented"; return nil }

// Compile-time check that QueryService implements InternalQueryService.
var _ executor.InternalQueryService = (*QueryService)(nil)

// AddQueryPattern adds a query pattern with an expected result.
func (m *QueryService) AddQueryPattern(pattern string, result *sqltypes.Result) {
	_ = "STUB: not implemented"
	return
}

// AddQueryPatternWithCallback adds a query pattern with a callback.
func (m *QueryService) AddQueryPatternWithCallback(pattern string, result *sqltypes.Result, callback func(string)) {
	_ = "STUB: not implemented"
	return
}

// AddQueryPatternWithError adds a query pattern that returns an error.
func (m *QueryService) AddQueryPatternWithError(pattern string, err error) {
	_ = "STUB: not implemented"
	return
}

// AddQueryPatternWithContextCallback adds a query pattern with a context-aware callback.
// This is useful for testing blocking queries that should respond to context cancellation.
func (m *QueryService) AddQueryPatternWithContextCallback(pattern string, result *sqltypes.Result, callback func(context.Context, string)) {
	_ = "STUB: not implemented"
	return
}

// AddQueryPatternOnce adds a query pattern that is consumed after the first match.
// This is useful when you need different results for subsequent calls to the same query.
func (m *QueryService) AddQueryPatternOnce(pattern string, result *sqltypes.Result) {
	_ = "STUB: not implemented"
	return
}

// AddQueryPatternOnceWithError adds a query pattern that returns an error and is consumed after the first match.
func (m *QueryService) AddQueryPatternOnceWithError(pattern string, err error) {
	_ = "STUB: not implemented"
	return
}

// ExpectationsWereMet returns an error if any consumeOnce patterns were not matched.
// This is useful for verifying that all expected queries were executed.
func (m *QueryService) ExpectationsWereMet() error { _ = "STUB: not implemented"; return nil }

// Query implements executor.InternalQueryService.
func (m *QueryService) Query(ctx context.Context, queryStr string) (*sqltypes.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Copy the matched pattern's data before potentially modifying the slice

// Remove the pattern if it should only be used once

// QueryMultiStatement implements executor.InternalQueryService.
// For the mock, this delegates to Query (result is discarded).
func (m *QueryService) QueryMultiStatement(ctx context.Context, queryStr string) error {
	_ = "STUB: not implemented"
	return nil
}

// QueryArgs implements executor.InternalQueryService.
// For the mock, arguments are ignored and matching is done solely on the query string.
func (m *QueryService) QueryArgs(ctx context.Context, queryStr string, args ...any) (*sqltypes.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// MakeQueryResult creates a sqltypes.Result from columns and rows.
}

func MakeQueryResult(columns []string, rows [][]any) *sqltypes.Result {
	_ = "STUB: not implemented"
	return nil
}
