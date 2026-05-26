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

//go:build ruleguard

package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
)

func disallowUnderscoreInFlags(m dsl.Matcher) { _ = "STUB: not implemented"; return }

func disallowOtelMeterOutsideMetricsFiles(m dsl.Matcher) { _ = "STUB: not implemented"; return }

func disallowMetricsConstructorArgs(m dsl.Matcher) { _ = "STUB: not implemented"; return }

func requireContextBackgroundJustification(m dsl.Matcher) { _ = "STUB: not implemented"; return }

func requireGrpcCommonNewClient(m dsl.Matcher) { _ = "STUB: not implemented"; return }

// disallowDirectExecCommandContext enforces use of executil.Command() for
// graceful termination support, proper environment variable handling, and
// trace propagation.
func disallowDirectExecCommandContext(m dsl.Matcher) {
	_ = "STUB: not implemented"

	// TODO: Also disallow exec.Command() with no context
	return
}

// disallowDirectProcessTermination enforces use of executil functions
// for consistent graceful SIGTERM -> SIGKILL termination.
func disallowDirectProcessTermination(m dsl.Matcher) { _ = "STUB: not implemented"; return }

// disallowDirectPgctldStopInTests prevents test code from calling pgctld Stop() directly
// in test packages that run multipooler alongside pgctld. The postgres monitor runs
// continuously and will restart postgres immediately after it is stopped, causing races.
// Use ShardSetup.StopPostgres instead — it disables restarts before stopping and returns
// a resume function.
//
// Excluded: pgctld package tests (no multipooler running) and non-test files (e.g.
// ShardSetup.StopPostgres itself calls Stop internally in setup.go).
func disallowDirectPgctldStopInTests(m dsl.Matcher) { _ = "STUB: not implemented"; return }
