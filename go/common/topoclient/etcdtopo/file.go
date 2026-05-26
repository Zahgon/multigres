// Copyright 2019 The Vitess Authors.
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

package etcdtopo

import (
	"context"

	"github.com/multigres/multigres/go/common/topoclient"
)

// Create is part of the topoclient.Conn interface.
func (s *etcdtopo) Create(ctx context.Context, filePath string, contents []byte) (topoclient.Version, error) {
	_ = "STUB: not implemented"
	return *new(topoclient.Version), nil
}

// We have to do a transaction, comparing existing version with 0.
// This means: if the file doesn't exist, create it.

// Update is part of the topoclient.Conn interface.
func (s *etcdtopo) Update(ctx context.Context, filePath string, contents []byte, version topoclient.Version) (topoclient.Version, error) {
	_ = "STUB: not implemented"
	return *new(topoclient.Version), nil
}

// We have to do a transaction. This means: if the
// current file revision is what we expect, save it.

// No version specified. We can use a simple unconditional Put.

// Get is part of the topoclient.Conn interface.
func (s *etcdtopo) Get(ctx context.Context, filePath string) ([]byte, topoclient.Version, error) {
	_ = "STUB: not implemented"
	return nil, *new(topoclient.Version), nil
}

// GetVersion is part of the topoclient.Conn interface.
func (s *etcdtopo) GetVersion(ctx context.Context, filePath string, version int64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List is part of the topoclient.Conn interface.
func (s *etcdtopo) List(ctx context.Context, filePathPrefix string) ([]topoclient.KVInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete is part of the topoclient.Conn interface.
func (s *etcdtopo) Delete(ctx context.Context, filePath string, version topoclient.Version) error {
	_ = "STUB: not implemented"
	return nil
}

// We have to do a transaction. This means: if the
// node revision is what we expect, delete it,
// otherwise get the file. If the transaction doesn't
// succeed, we also ask for the value of the
// node. That way we'll know if it failed because it
// didn't exist, or because the version was wrong.

// This is just a regular unconditional Delete here.
