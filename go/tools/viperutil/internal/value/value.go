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

package value

import (
	"errors"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/multigres/multigres/go/tools/viperutil/internal/registry"
	"github.com/multigres/multigres/go/tools/viperutil/internal/sync"
)

// Registerable is the subset of the interface exposed by Values (which is
// declared in the public viperutil package).
//
// We need a separate interface type because Go generics do not let you define
// a function that takes Value[T] for many, different T's, which we want to do
// for BindFlags.
type Registerable interface {
	Key() string
	Registry() registry.Bindable
	Flag(fs *pflag.FlagSet) (*pflag.Flag, error)
}

// Base is the base functionality shared by Static and Dynamic values. It
// implements viperutil.Value.
type Base[T any] struct {
	KeyName    string
	DefaultVal T

	GetFunc      func(v *viper.Viper) func(key string) T
	BoundGetFunc func(key string) T

	Aliases  []string
	FlagName string
	EnvVars  []string
}

func (val *Base[T]) Key() string { _ = "STUB: not implemented"; return "" }
func (val *Base[T]) Default() T  { _ = "STUB: not implemented"; return *new(T) }
func (val *Base[T]) Get() T      { _ = "STUB: not implemented"; return *new(T) }

// ErrNoFlagDefined is returned when a Value has a FlagName set, but the given
// FlagSet does not define a flag with that name.
var ErrNoFlagDefined = errors.New("flag not defined")

// Flag is part of the Registerable interface. If the given flag set has a flag
// with the name of this value's configured flag, that flag is returned, along
// with a nil error. If no flag exists on the flag set with that name, an error
// is returned.
//
// If the value is not configured to correspond to a flag (FlagName == ""), then
// (nil, nil) is returned.
func (val *Base[T]) Flag(fs *pflag.FlagSet) (*pflag.Flag, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (val *Base[T]) bind(v registry.Bindable) { _ = "STUB: not implemented"; return }

// BindFlags creates bindings between each value's registry and the given flag
// set. This function will panic if any of the values defines a flag that does
// not exist in the flag set.
func BindFlags(fs *pflag.FlagSet, values ...Registerable) { _ = "STUB: not implemented"; return }

// Static is a static value. Static values register to a Static registry, and
// do not respond to changes to config files. Their Get() method will return the
// same value for the lifetime of the process.
type Static[T any] struct {
	*Base[T]
	staticReg *viper.Viper
}

// NewStatic returns a static value derived from the given base value, after
// binding it to the provided static registry.
func NewStatic[T any](staticReg *viper.Viper, base *Base[T]) *Static[T] {
	_ = "STUB: not implemented"
	return nil
}

func (val *Static[T]) Registry() registry.Bindable {
	_ = "STUB: not implemented"
	return *new(registry.Bindable)
}

func (val *Static[T]) Set(v T) { _ = "STUB: not implemented"; return }

// Dynamic is a dynamic value. Dynamic values register to the Dynamic registry,
// and respond to changes to watched config files. Their Get() methods will
// return whatever value is currently live in the config, in a threadsafe
// manner.
type Dynamic[T any] struct {
	*Base[T]
	dynamicReg *sync.Viper
}

// NewDynamic returns a dynamic value derived from the given base value, after
// binding it to the provided dynamic registry and wrapping its GetFunc to be threadsafe
// with respect to config reloading.
func NewDynamic[T any](dynamicReg *sync.Viper, base *Base[T]) *Dynamic[T] {
	_ = "STUB: not implemented"
	return nil
}

func (val *Dynamic[T]) Registry() registry.Bindable {
	_ = "STUB: not implemented"
	return *new(registry.Bindable)
}

func (val *Dynamic[T]) Set(v T) { _ = "STUB: not implemented"; return }
