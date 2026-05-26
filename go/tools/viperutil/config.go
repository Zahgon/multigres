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

package viperutil

import (
	"context"
	"reflect"
	"slices"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type ViperConfig struct {
	configPaths                  Value[[]string]
	configType                   Value[string]
	configName                   Value[string]
	configFile                   Value[string]
	configFileNotFoundHandling   Value[ConfigFileNotFoundHandling]
	configPersistenceMinInterval Value[time.Duration]
}

func NewViperConfig(reg *Registry) *ViperConfig { _ = "STUB: not implemented"; return nil }

// Use MTDATAROOT environment variable if set, otherwise fall back to pwd/multigres_local

// Need to re-trigger the SetDefault call done during Configure.

// RegisterFlags installs the flags that control viper config-loading behavior.
// It is exported to be called by servenv before parsing flags for all binaries.
//
// It cannot be registered here via servenv.OnParse since this causes an import
// cycle.
func (vc *ViperConfig) RegisterFlags(fs *pflag.FlagSet) { _ = "STUB: not implemented"; return }

// LoadConfig attempts to find, and then load, a config file for viper-backed
// config values to use.
//
// Config searching follows the behavior used by viper [1], namely:
//   - --config-file (full path, including extension) if set will be used to the
//     exclusion of all other flags.
//   - --config-type is required if the config file does not have one of viper's
//     supported extensions (.yaml, .yml, .json, and so on)
//
// An additional --config-file-not-found-handling flag controls how to treat the
// situation where viper cannot find any config files in any of the provided
// paths (for ex, users may want to exit immediately if a config file that
// should exist doesn't for some reason, or may wish to operate with flags and
// environment variables alone, and not use config files at all).
//
// If a config file is successfully loaded, then the dynamic registry will also
// start watching that file for changes. In addition, in-memory changes to the
// config (for example, from a vtgate or vttablet's debugenv) will be persisted
// back to disk, with writes occurring no more frequently than the
// --config-persistence-min-interval flag.
//
// A cancel function is returned to stop the re-persistence background thread,
// if one was started.
//
// [1]: https://github.com/spf13/viper#reading-config-files.
func (vc *ViperConfig) LoadConfig(reg *Registry) (context.CancelFunc, error) {
	_ = "STUB: not implemented"
	return *new(context.CancelFunc), nil
}

// TODO: @rafael (add warning and point to docs)
// after warning, ignore the error

// isConfigFileNotFoundError checks if the error is caused because the file wasn't found.
func isConfigFileNotFoundError(err error) bool { _ = "STUB: not implemented"; return false }

// NotifyConfigReload adds a subscription that the dynamic registry will attempt
// to notify on config changes. The notification fires after the updated config
// has been loaded from disk into the live config.
//
// Analogous to signal.Notify, notifications are sent non-blocking, so users
// should account for this when writing code to consume from the channel.
//
// This function must be called prior to LoadConfig; it will panic if called
// after the dynamic registry has started watching the loaded config.
func NotifyConfigReload(reg *Registry, ch chan<- struct{}) { _ = "STUB: not implemented"; return }

// ConfigFileNotFoundHandling is an enum to control how LoadConfig treats errors
// of type viper.ConfigFileNotFoundError when loading a config.
type ConfigFileNotFoundHandling int

const (
	// IgnoreConfigFileNotFound causes LoadConfig to completely ignore a
	// ConfigFileNotFoundError (i.e. not even logging it).
	IgnoreConfigFileNotFound ConfigFileNotFoundHandling = iota
	// WarnOnConfigFileNotFound causes LoadConfig to log a warning with details
	// about the failed config load, but otherwise proceeds with the given
	// process, which will get config values entirely from defaults,
	// environment variables, and flags.
	WarnOnConfigFileNotFound
	// ErrorOnConfigFileNotFound causes LoadConfig to return the
	// ConfigFileNotFoundError after logging an error.
	ErrorOnConfigFileNotFound
	// ExitOnConfigFileNotFound causes LoadConfig to log.Fatal on a
	// ConfigFileNotFoundError.
	ExitOnConfigFileNotFound
)

var (
	handlingNames         []string
	handlingNamesToValues = map[string]int{
		"ignore": int(IgnoreConfigFileNotFound),
		"warn":   int(WarnOnConfigFileNotFound),
		"error":  int(ErrorOnConfigFileNotFound),
		"exit":   int(ExitOnConfigFileNotFound),
	}
	handlingValuesToNames map[int]string
)

func getHandlingValue(v *viper.Viper) func(key string) ConfigFileNotFoundHandling {
	_ = "STUB: not implemented"
	return nil
}

func decodeHandlingValue(from, to reflect.Type, data any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func init() {
	handlingNames = make([]string, 0, len(handlingNamesToValues))
	handlingValuesToNames = make(map[int]string, len(handlingNamesToValues))

	for name, val := range handlingNamesToValues {
		handlingValuesToNames[val] = name
		handlingNames = append(handlingNames, name)
	}

	slices.Sort(handlingNames)
}

func (h *ConfigFileNotFoundHandling) Set(arg string) error { _ = "STUB: not implemented"; return nil }

func (h *ConfigFileNotFoundHandling) String() string { _ = "STUB: not implemented"; return "" }

func (h *ConfigFileNotFoundHandling) Type() string { _ = "STUB: not implemented"; return "" }
