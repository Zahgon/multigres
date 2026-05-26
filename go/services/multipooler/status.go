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

// Package multipooler provides multipooler functionality.
package multipooler

import (
	"net/http"
	"sync"
)

// Link represents a link on the status page.
type Link struct {
	Title       string
	Description string
	Link        string
}

// Status represents the response from the temporary status endpoint
type Status struct {
	mu sync.Mutex

	Title string `json:"title"`

	InitError  string            `json:"init_error"`
	TopoStatus map[string]string `json:"topo_status"`

	Cell           string `json:"cell"`
	ServiceID      string `json:"service_id"`
	Database       string `json:"database"`
	TableGroup     string `json:"table_group"`
	PgctldAddr     string `json:"pgctld_addr"`
	SocketFilePath string `json:"socket_file_path"`

	Links []Link `json:"links"`
}

// handleIndex serves the index page
func (mp *MultiPooler) handleIndex(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
