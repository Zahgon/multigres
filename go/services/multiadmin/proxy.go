// Copyright 2025 Supabase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package multiadmin

import (
	"net/http"
)

type proxyPathInfo struct {
	serviceType string
	cellName    string
	serviceName string
}

type serviceTarget struct {
	host          string
	port          int
	proxyBasePath string
}

// parseProxyPath extracts routing information from the proxy path
func parseProxyPath(path string) (*proxyPathInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// lookupCellService retrieves service information from topology service
func (ma *MultiAdmin) lookupCellService(r *http.Request, pathInfo proxyPathInfo) (hostname string, httpPort int, err error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

// resolveServiceTarget determines the target host, port, and base path for the proxy
func (ma *MultiAdmin) resolveServiceTarget(r *http.Request, pathInfo proxyPathInfo) (*serviceTarget, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Global service - multiadmin proxying to itself

// Cell services - multigateway, multipooler, multiorch

// handleProxy routes requests to backend services based on path:
// /proxy/admin/{cell}/{name} -> routes to multiadmin (proxying to itself)
// /proxy/gate/{cell}/{name} -> routes to multigateway
// /proxy/pool/{cell}/{name} -> routes to multipooler
// /proxy/orch/{cell}/{name} -> routes to multiorch
func (ma *MultiAdmin) handleProxy(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Create reverse proxy to the target service

// Modify the director to strip the proxy prefix from the request path

// Strip the proxy prefix to get the actual path the backend expects

// Intercept the response to rewrite HTML content

// Only rewrite HTML responses

// Rewrite HTML to fix asset and link paths

// If rewriting fails, return original content

// Update response body

// rewriteHTML injects a <base> tag and rewrites absolute URLs in HTML content
func rewriteHTML(htmlContent []byte, proxyBasePath string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Traverse the document and rewrite URLs

// Inject <base> tag into <head>

// Create <base> element

// Insert as first child of <head>

// Rewrite absolute URLs in href and src attributes

// Skip rewriting if already prefixed with /proxy/ or current proxy base path

// Rewrite absolute path to be relative to proxy base

// Render the modified HTML back to bytes
