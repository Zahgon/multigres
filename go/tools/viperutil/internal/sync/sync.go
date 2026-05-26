// Copyright 2023 The Vitess Authors.
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
//
// Modifications Copyright 2025 Supabase, Inc.

package sync

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/spf13/afero"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Viper is a wrapper around a pair of viper.Viper instances to provide config-
// reloading in a threadsafe manner.
//
// It maintains one viper, called "disk", which does the actual config watch and
// reload (via viper's WatchConfig), and a second viper, called "live", which
// Values (registered via viperutil.Configure with Dynamic=true) access their
// settings from. The "live" config only updates after blocking all values from
// reading in order to swap in the most recently-loaded config from the "disk".
type Viper struct {
	m    sync.Mutex // prevents races between loadFromDisk and AllSettings
	disk *viper.Viper
	live *viper.Viper
	keys map[string]*sync.Mutex

	subscribers    []chan<- struct{}
	watchingConfig bool

	fs afero.Fs

	setCh chan struct{}

	// for testing purposes only
	onConfigWrite func()
}

func (v *Viper) SetFs(fs afero.Fs) { _ = "STUB: not implemented"; return }

// New returns a new synced Viper.
func New() *Viper { _ = "STUB: not implemented"; return nil }

// default Fs used by viper, but we need this set so loadFromDisk doesn't accidentally nil-out the live fs

// Set sets the given key to the given value, in both the disk and live vipers.
func (v *Viper) Set(key string, value any) { _ = "STUB: not implemented"; return }

// We must not update v.disk here; explicit calls to Set will supersede all
// future config reloads.

// Do a non-blocking signal to persist here. Our channel has a buffer of 1,
// so if we've signalled for some other Set call that hasn't been persisted
// yet, this Set will get persisted along with that one and any other
// pending in-memory changes.

// ErrDuplicateWatch is returned when Watch is called on a synced Viper which
// has already started a watch.
var ErrDuplicateWatch = errors.New("duplicate watch")

// Watch starts watching the config used by the passed-in Viper. Before starting
// the watch, the synced viper will perform an initial read and load from disk
// so that the live config is ready for use without requiring an initial config
// change.
//
// If the given static viper did not load a config file (and is instead relying
// purely on defaults, flags, and environment variables), then the settings of
// that viper are merged over, and this synced Viper may be used to set up an
// actual watch later. Additionally, this starts a background goroutine to
// persist changes made in-memory back to disk. It returns a cancel func to stop
// the persist loop, which the caller is responsible for calling during
// shutdown (see package servenv for an example).
//
// This does two things — one which is a nice-to-have, and another which is
// necessary for correctness.
//
// 1. Writing in-memory changes (which usually occur through a request to a
// /debug/env endpoint) ensures they are persisted across process restarts.
// 2. Writing in-memory changes ensures that subsequent modifications to the
// config file do not clobber those changes. Because viper loads the entire
// config on-change, rather than an incremental (diff) load, if a user were to
// edit an unrelated key (keyA) in the file, and we did not persist the
// in-memory change (keyB), then future calls to keyB.Get() would return the
// older value.
//
// If this synced viper is already watching a config file, this function returns
// an ErrDuplicateWatch. Other errors may be returned via underlying viper code
// to ensure the config file can be read in properly.
func (v *Viper) Watch(ctx context.Context, static *viper.Viper, minWaitInterval time.Duration) (cancel context.CancelFunc, err error) {
	_ = "STUB: not implemented"
	return *new(context.CancelFunc), nil
}

// No config file to watch, just merge the settings and return.

// This won't fire until after the config has been updated on v.live.

func (v *Viper) persistChanges(ctx context.Context, minWaitInterval time.Duration) {
	_ = "STUB: not implemented"
	return
}

// If we failed to persist, don't wait the entire interval before
// writing again, instead writing immediately on the next request.

// WriteConfig writes the live viper config back to disk.
func (v *Viper) WriteConfig() error { _ = "STUB: not implemented"; return nil }

// This won't fire until after the config has been written.

// Notify adds a subscription that this synced viper will attempt to notify on
// config changes, after the updated config has been copied over from disk to
// live.
//
// Analogous to signal.Notify, notifications are sent non-blocking, so users
// should account for this when consuming from the channel they've provided.
//
// This function must be called prior to setting up a Watch; it will panic if a
// a watch has already been established on this synced Viper.
func (v *Viper) Notify(ch chan<- struct{}) { _ = "STUB: not implemented"; return }

// AllSettings returns the current live settings.
func (v *Viper) AllSettings() map[string]any { _ = "STUB: not implemented"; return nil }

func (v *Viper) loadFromDisk() { _ = "STUB: not implemented"; return }

// Reset v.live so explicit Set calls don't win over what's just changed on
// disk.

// Fun fact! MergeConfigMap actually only ever returns nil. Maybe in an
// older version of viper it used to actually handle errors, but now it
// decidedly does not. See https://github.com/spf13/viper/blob/v1.8.1/viper.go#L1492-L1499.

// begin implementation of registry.Bindable for sync.Viper

func (v *Viper) BindEnv(vars ...string) error { _ = "STUB: not implemented"; return nil }

func (v *Viper) BindPFlag(key string, flag *pflag.Flag) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Viper) RegisterAlias(alias string, key string) { _ = "STUB: not implemented"; return }

func (v *Viper) SetDefault(key string, value any) { _ = "STUB: not implemented"; return }

// end implementation of registry.Bindable for sync.Viper

// AdaptGetter wraps a get function (matching the signature of
// viperutil.Options.GetFunc) to be threadsafe with the passed-in synced Viper.
//
// It must be called prior to starting a watch on the synced Viper; it will
// panic if a watch has already been established.
//
// This function must be called at most once per key; it will panic if attempting
// to adapt multiple getters for the same key.
func AdaptGetter[T any](key string, getter func(v *viper.Viper) func(key string) T, v *Viper) func(key string) T {
	_ = "STUB: not implemented"
	return nil
}
