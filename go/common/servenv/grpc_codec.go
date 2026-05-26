// Copyright 2021 The Vitess Authors.
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

package servenv

import (
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/mem"

	// Guarantee that the built-in proto is called registered before this one
	// so that it can be replaced.
	_ "google.golang.org/grpc/encoding/proto" //nolint:revive
)

// Name is the name registered for the proto compressor.
const Name = "proto"

type vtprotoMessage interface {
	MarshalToSizedBufferVT(data []byte) (int, error)
	UnmarshalVT([]byte) error
	SizeVT() int
}

type Codec struct {
	fallback encoding.CodecV2
}

func (Codec) Name() string { _ = "STUB: not implemented"; return "" }

var defaultBufferPool = mem.DefaultBufferPool()

func (c *Codec) Marshal(v any) (mem.BufferSlice, error) {
	_ = "STUB: not implemented"
	return *new(mem.BufferSlice), nil
}

func (c *Codec) Unmarshal(data mem.BufferSlice, v any) error { _ = "STUB: not implemented"; return nil }

func init() {
	encoding.RegisterCodecV2(&Codec{
		fallback: encoding.GetCodecV2("proto"),
	})
}
