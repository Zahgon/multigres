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

package buffer

import (
	"time"

	"github.com/spf13/pflag"

	"github.com/multigres/multigres/go/tools/viperutil"
)

// Config holds all configuration for failover buffering.
type Config struct {
	Enabled                 viperutil.Value[bool]
	Window                  viperutil.Value[time.Duration]
	Size                    viperutil.Value[int]
	MaxFailoverDuration     viperutil.Value[time.Duration]
	MinTimeBetweenFailovers viperutil.Value[time.Duration]
	DrainConcurrency        viperutil.Value[int]
}

// NewConfig creates a Config with all flags registered in the given registry.
func NewConfig(reg *viperutil.Registry) *Config { _ = "STUB: not implemented"; return nil }

// RegisterFlags registers buffer flags on the given FlagSet.
func (c *Config) RegisterFlags(fs *pflag.FlagSet) { _ = "STUB: not implemented"; return }

// Validate checks that the configuration values are consistent.
func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }
