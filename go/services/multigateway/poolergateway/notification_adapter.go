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
	"errors"
	"log/slog"
	"sync"

	"github.com/multigres/multigres/go/common/sqltypes"
	multipoolerpb "github.com/multigres/multigres/go/pb/multipoolerservice"
)

// GRPCNotificationManager implements sqltypes.NotificationManager by calling
// the pooler's StreamNotifications gRPC. It manages per-channel streams
// and fans out notifications to subscriber channels.
//
// NOTE: Currently opens one gRPC stream per unique PG channel. If the number
// of distinct channels grows large, consider consolidating into a single
// bidirectional stream per gateway-pooler pair with dynamic channel add/remove.
// The StreamNotifications RPC already accepts repeated channels, so the
// multipooler side supports multi-channel subscriptions on a single stream.
// HTTP/2 multiplexes all streams over one TCP connection, but each stream
// still requires a goroutine, so this may need revisiting at scale.
type GRPCNotificationManager struct {
	getClient func() multipoolerpb.MultiPoolerServiceClient
	logger    *slog.Logger
	metrics   *NotificationMetrics

	mu sync.Mutex
	// channels tracks: pgChannel -> list of subscriber notifCh
	channels map[string][]chan *sqltypes.Notification
	// streams tracks: pgChannel -> cancel func for the gRPC stream
	streams map[string]context.CancelFunc
}

// NewGRPCNotificationManager creates a notification manager backed by gRPC.
func NewGRPCNotificationManager(
	getClient func() multipoolerpb.MultiPoolerServiceClient,
	logger *slog.Logger,
	metrics *NotificationMetrics,
) *GRPCNotificationManager {
	_ = "STUB: not implemented"
	return nil
}

// Subscribe registers notifCh to receive notifications for pgChannel.
func (m *GRPCNotificationManager) Subscribe(pgChannel string, notifCh chan *sqltypes.Notification) {
	_ = "STUB: not implemented"
	return
}

// If this is the first subscriber for this channel, start a gRPC stream.

//nolint:gocritic // Long-lived gRPC stream for notification fan-out, not tied to any request.

// Wait for the stream to be established before returning,
// so that a subsequent NOTIFY will be captured.
// This wait happens outside the lock to avoid the unlock/relock race window.

// Unsubscribe removes notifCh from pgChannel subscribers.
func (m *GRPCNotificationManager) Unsubscribe(pgChannel string, notifCh chan *sqltypes.Notification) {
	_ = "STUB: not implemented"
	return
}

// If no more subscribers, cancel the gRPC stream.

//nolint:gocritic // Metric recording at unsubscribe time, no request context available.

// UnsubscribeAll removes notifCh from all channels.
func (m *GRPCNotificationManager) UnsubscribeAll(notifCh chan *sqltypes.Notification) {
	_ = "STUB: not implemented"
	return
}

//nolint:gocritic // Metric recording at unsubscribe time, no request context available.

// streamNotifications opens a StreamNotifications gRPC stream and fans out
// notifications to all subscribers for the given channel. On stream failure,
// it retries with backoff until the context is cancelled (last subscriber left).
func (m *GRPCNotificationManager) streamNotifications(ctx context.Context, pgChannel string, ready chan struct{}) {
	_ = "STUB: not implemented"
	return
}

// cancelled — no more subscribers

// Backoff before retry.

// runStream establishes a single gRPC stream and processes notifications until
// an error occurs or the context is cancelled. On the first attempt, it signals
// the ready channel after the stream is established.
func (m *GRPCNotificationManager) runStream(
	ctx context.Context, pgChannel string, firstAttempt bool, ready chan struct{},
) error {
	_ = "STUB: not implemented"
	return nil
}

// Wait for the ready signal (empty first message) from the pooler.
// This ensures LISTEN is active on PG before the gateway returns LISTEN OK to the client.

// errNoClient is a sentinel error for when no gRPC client is available.
var errNoClient = errors.New("no gRPC client available for notifications")
