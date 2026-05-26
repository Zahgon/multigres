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

// ListDir is part of the topoclient.Conn interface.
func (s *etcdtopo) ListDir(ctx context.Context, dirPath string, full bool) ([]topoclient.DirEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Special case where s.root is "/", dirPath is empty,
// we would end up with "//". in that case, we want "/".

// No key starts with this prefix, means the directory
// doesn't exist.

// Remove the prefix, base path.

// Keep only the part until the first '/'.

// Remove duplicates, add to list.

// Only locks have a lease associated with them.
