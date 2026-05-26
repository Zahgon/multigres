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

package eventlog

import (
	"context"
	"log/slog"
)

// Outcome represents the lifecycle state of an event.
type Outcome string

const (
	Started Outcome = "started"
	Success Outcome = "success"
	Failed  Outcome = "failed"
)

// Event is implemented by all event structs.
type Event interface {
	EventType() string
	LogAttrs() []slog.Attr
}

// Emit logs a lifecycle event through the provided logger.
// Context is required for OpenTelemetry trace correlation.
// Failed outcome logs at ERROR; all others at INFO.
func Emit(ctx context.Context, logger *slog.Logger, outcome Outcome, event Event, extra ...any) {
	_ = "STUB: not implemented"
	return
}

// Convert extra key-value pairs to attrs.
