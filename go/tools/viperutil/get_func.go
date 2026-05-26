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
	"github.com/spf13/viper"
)

// GetFuncForType returns the default getter function for a given type T. A
// getter function is a function which takes a viper and returns a function that
// takes a key and (finally!) returns a value of type T.
//
// For example, the default getter for a value of type string is a function that
// takes a viper instance v and calls v.GetString with the provided key.
//
// In most cases, callers of Configure should be able to rely on the defaults
// provided here (and may refer to get_func_test.go for an up-to-date example
// of the provided functionalities), but if more fine-grained control is needed
// (this should be an **exceptional** circumstance), they may provide their own
// GetFunc as an option to Configure.
//
// This function may panic if called for an unsupported type. This is captured
// in the test code as well.
func GetFuncForType[T any]() func(v *viper.Viper) func(key string) T {
	_ = "STUB: not implemented"
	return nil
}

// Unupported, fallthrough to `if f == nil` check below switch.

// Even though the code would be extremely similar to slice types, we
// cannot support arrays because there's no way to write a function that
// returns, say, [N]int, for some value of N which we only know at
// runtime.

func unmarshalFunc[T any]() func(v *viper.Viper) func(key string) T {
	_ = "STUB: not implemented"
	return nil
}

// TODO: panic on this error

func getCastedInt[T int8 | int16]() func(v *viper.Viper) func(key string) T {
	_ = "STUB: not implemented"
	return nil
}

func getCastedUint[T uint8 | uint16]() func(v *viper.Viper) func(key string) T {
	_ = "STUB: not implemented"
	return nil
}

func getComplex[T complex64 | complex128](bitSize int) func(v *viper.Viper) func(key string) T {
	_ = "STUB: not implemented"
	return nil
}

// TODO: wrap with more details (key, type (64 vs 128), etc)
