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

package pubsub

import (
	"context"

	"go.opentelemetry.io/otel/metric"
)

// PubSubMetrics holds OTel metrics for the PubSub notification listener.
// All metrics are recorded from the single-threaded event loop in Listener.run(),
// so no synchronization is needed for recording.
type PubSubMetrics struct {
	channels        metric.Int64UpDownCounter
	subscribers     metric.Int64UpDownCounter
	notifDropped    metric.Int64Counter
	reconnects      metric.Int64Counter
	reconnectGapDur metric.Float64Histogram
}

// ChannelAdd increments the channel gauge when a new PG channel gets its first subscriber.
func (m *PubSubMetrics) ChannelAdd(ctx context.Context) { _ = "STUB: not implemented"; return }

// ChannelRemove decrements the channel gauge when a PG channel loses its last subscriber.
func (m *PubSubMetrics) ChannelRemove(ctx context.Context) { _ = "STUB: not implemented"; return }

// SubscriberAdd increments the subscriber gauge when a new unique subscriber connects.
func (m *PubSubMetrics) SubscriberAdd(ctx context.Context) { _ = "STUB: not implemented"; return }

// SubscriberRemove decrements the subscriber gauge when a subscriber disconnects from all channels.
func (m *PubSubMetrics) SubscriberRemove(ctx context.Context) { _ = "STUB: not implemented"; return }

// NotificationDropped increments the drop counter when a notification cannot be
// delivered because a subscriber's channel buffer is full.
func (m *PubSubMetrics) NotificationDropped(ctx context.Context) { _ = "STUB: not implemented"; return }

// Reconnect increments the reconnect counter when the PG connection is lost
// and a reconnect is scheduled.
func (m *PubSubMetrics) Reconnect(ctx context.Context) { _ = "STUB: not implemented"; return }

// ReconnectGapDuration records the duration between disconnect and successful
// reconnect. This is a proxy for the notification loss window since PG
// LISTEN/NOTIFY is fire-and-forget.
func (m *PubSubMetrics) ReconnectGapDuration(ctx context.Context, durationSec float64) {
	_ = "STUB: not implemented"
	return
}

// NewPubSubMetrics initialises OTel metrics for the PubSub listener.
// Individual metrics that fail to initialise use noop implementations
// and are included in the returned error.
func NewPubSubMetrics() (*PubSubMetrics, error) { _ = "STUB: not implemented"; return nil, nil }
