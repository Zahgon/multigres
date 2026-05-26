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

package poolergateway

import (
	"context"

	"go.opentelemetry.io/otel/metric"
)

// NotificationMetrics holds OTel metrics for the gateway notification system.
type NotificationMetrics struct {
	streams      metric.Int64UpDownCounter
	notifDropped metric.Int64Counter
}

// StreamAdd increments the stream gauge when a new gRPC notification stream is opened.
func (m *NotificationMetrics) StreamAdd(ctx context.Context) { _ = "STUB: not implemented"; return }

// StreamRemove decrements the stream gauge when a gRPC notification stream is closed.
func (m *NotificationMetrics) StreamRemove(ctx context.Context) { _ = "STUB: not implemented"; return }

// NotificationDropped increments the drop counter when a notification cannot be
// delivered because a subscriber's channel buffer is full.
func (m *NotificationMetrics) NotificationDropped(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

// NewNotificationMetrics initialises OTel metrics for gateway notifications.
// Individual metrics that fail to initialise use noop implementations
// and are included in the returned error.
func NewNotificationMetrics() (*NotificationMetrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
