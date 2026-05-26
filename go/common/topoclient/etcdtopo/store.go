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
//
// Modifications Copyright 2026 Supabase, Inc.

/*
Package etcdtopo implements topoclient.Conn with etcd as the backend.

We expect the following behavior from the etcd client library:

  - Get and Delete return ErrorCodeKeyNotFound if the node doesn't exist.
  - Create returns ErrorCodeNodeExist if the node already exists.
  - Intermediate directories are always created automatically if necessary.
  - Set returns ErrorCodeKeyNotFound if the node doesn't exist already.
  - It returns ErrorCodeTestFailed if the provided version index doesn't match.

We follow these conventions within this package:

  - Call convertError(err) on any errors returned from the etcd client library.
    Functions defined in this package can be assumed to have already converted
    errors as necessary.
*/
package etcdtopo

import (
	"crypto/tls"

	"github.com/spf13/pflag"

	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/multigres/multigres/go/common/servenv"
	"github.com/multigres/multigres/go/common/timeouts"
	"github.com/multigres/multigres/go/common/topoclient"
)

var (
	clientCertPath string
	clientKeyPath  string
	serverCaPath   string
)

var _ topoclient.Conn = (*etcdtopo)(nil)

// remoteOperationTimeout is used for operations where we have to
// call out to etcd for initial data fetches (e.g., watch setup).
const remoteOperationTimeout = timeouts.RemoteOperationTimeout

// Factory is the etcd topoclient.Factory implementation.
type Factory struct{}

// HasGlobalReadOnlyCell is part of the topoclient.Factory interface.
func (f Factory) HasGlobalReadOnlyCell(serverAddr, root string) bool {
	_ = "STUB: not implemented"

	// Create is part of the topoclient.Factory interface.
	return false
}

func (f Factory) Create(cell, root string, serverAddrs []string) (topoclient.Conn, error) {
	_ = "STUB: not implemented"
	return *new(topoclient.Conn), nil
}

// etcdtopo is the implementation of topoclient.Conn for etcd.
type etcdtopo struct {
	// cli is the v3 client.
	cli *clientv3.Client

	// root is the root path for this client.
	root string

	running chan struct{}
}

func init() {
	servenv.OnParse(registerEtcdTopoFlags)
	topoclient.RegisterFactory(topoclient.DefaultTopoImplementation, Factory{})
}

func registerEtcdTopoFlags(fs *pflag.FlagSet) { _ = "STUB: not implemented"; return }

// Close implements topoclient.Conn.Close.
// It will nil out the global and cells fields, so any attempt to
// reuse this server will panic.
func (s *etcdtopo) Close() error { _ = "STUB: not implemented"; return nil }

func newTLSConfig(certPath, keyPath, caPath string) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil,

		// If TLS is enabled, attach TLS config info.
		nil
}

// NewServerWithOpts creates a new server with the provided TLS options
func NewServerWithOpts(serverAddrs []string, root, certPath, keyPath, caPath string) (*etcdtopo, error) {
	_ = "STUB: not implemented"
	// TODO: Rename this to NewServer and change NewServer to a name that signifies it uses the process-wide TLS settings.
	return nil, nil
}

// grpc.WithBlock is deprecated but required by etcd client

// NewEtcdTopo returns a new etcdtopo.Server.
func NewEtcdTopo(serverAddrs []string, root string) (*etcdtopo, error) {
	_ = "STUB: not implemented"
	// TODO: Rename this to a name to signifies this function uses the process-wide TLS settings.
	return nil, nil
}
