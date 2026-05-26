// Copyright 2026 Supabase, Inc.
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

package engine

import (
	"context"

	"go.opentelemetry.io/otel/metric"
)

// Transaction outcome constants for metric attribution.
const (
	TxnOutcomeCommit   = "commit"
	TxnOutcomeRollback = "rollback"
)

// TransactionMetrics holds OTel metrics for transaction lifecycle tracking.
// Initialized once in the executor and injected into TransactionPrimitive
// at creation time, following the same pattern as HandlerMetrics.
type TransactionMetrics struct {
	duration TxnDuration
	count    TxnCount
}

// TxnDuration wraps a Float64Histogram for recording transaction durations.
type TxnDuration struct {
	metric.Float64Histogram
}

// Record records a transaction duration with the database and outcome attributes.
func (m TxnDuration) Record(ctx context.Context, durationSec float64, dbNamespace, outcome string) {
	_ = "STUB: not implemented"
	return
}

// TxnCount wraps an Int64Counter for counting completed transactions.
type TxnCount struct {
	metric.Int64Counter
}

// Add increments the transaction counter with the database and outcome attributes.
func (m TxnCount) Add(ctx context.Context, dbNamespace, outcome string) {
	_ = "STUB: not implemented"
	return
}

// RecordCompletion records both duration and count for a completed transaction.
func (m *TransactionMetrics) RecordCompletion(ctx context.Context, durationSec float64, dbNamespace, outcome string) {
	_ = "STUB: not implemented"
	return
}

// NewTransactionMetrics initialises OTel metrics for transaction tracking.
// Individual metrics that fail to initialise use noop implementations
// and are included in the returned error.
func NewTransactionMetrics() (*TransactionMetrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
