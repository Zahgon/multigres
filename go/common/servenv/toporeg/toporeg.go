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

// Package toporeg manages the registration of components to the topoclient.
//
// TODO: Consider adding senv.RegisterWithTopology(register, unregister, alarm) to servenv
// to simplify the current pattern where services manually wire up OnRun/OnClose hooks.
// This would consolidate the registration lifecycle management into servenv.
package toporeg

import (
	"context"
	"log/slog"
	"sync"
)

// TopoReg contains the metadata of the component being registered.
type TopoReg struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
	logger *slog.Logger

	unregister func(ctx context.Context) error
}

// Register registers the component using the register function. If the register function
// returns an error, it will be retried with exponential backoff until successful.
// The alarm will be invoked with the latest error message during retries. If the
// registration succeeds, the alarm will be invoked with an empty string.
func Register(register func(ctx context.Context) error, unregister func(ctx context.Context) error, alarm func(string)) *TopoReg {
	_ = "STUB: not implemented"
	return nil
}

// Use tp's ctx to abort retries if Unregister gets called.

// We've already tried once. Use WithInitialDelay to wait before retrying.

// Context cancelled

// Just call alarm. No need to spam logs.

// RegisterSynchronous registers the component synchronously, retrying with
// exponential backoff and jitter until successful or the context expires.
// Unlike Register, it blocks until registration succeeds and returns an error
// if it cannot complete within the context deadline.
//
// Use this when the caller must know registration succeeded before proceeding
// (e.g., claiming a PID prefix that other components depend on).
func RegisterSynchronous(ctx context.Context, register func(ctx context.Context) error, unregister func(ctx context.Context) error) (*TopoReg, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unregister unregisters the component from topology.
// It will terminate any retry goroutines that are still running.
// It is safe to call Unregister with a nil TopoReg.
func (tp *TopoReg) Unregister() {
	_ = "STUB: not implemented"
	// Safety
	return
}

// Use standalone ctx because tp.ctx is already canceled.
