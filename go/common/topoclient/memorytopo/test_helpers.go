// Copyright 2025 The Multigres Authors.
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

package memorytopo

import (
	"context"

	"github.com/multigres/multigres/go/common/topoclient"
)

// CreateTestGateway creates a test gateway in the specified cell for testing
func CreateTestGateway(ctx context.Context, ts topoclient.Store, cellName, gatewayName string) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateTestPooler creates a test pooler in the specified cell for testing
func CreateTestPooler(ctx context.Context, ts topoclient.Store, cellName, poolerName string, database string) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateTestOrchestrator creates a test orchestrator in the specified cell for testing
func CreateTestOrchestrator(ctx context.Context, ts topoclient.Store, cellName, orchName string) error {
	_ = "STUB: not implemented"
	return nil
}

// SetupMultiCellTestData creates test data across multiple cells for testing
func SetupMultiCellTestData(ctx context.Context, ts topoclient.Store) error {
	_ = "STUB: not implemented"
	// Create test gateways in multiple cells
	return nil
}

// Create test poolers in multiple cells

// Create test orchestrators in multiple cells
