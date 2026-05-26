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
// Modifications Copyright 2025 Supabase, Inc.

// Package netutil contains network-related utility functions.
package netutil

import (
	"net"
)

// SplitHostPort is an alternative to net.SplitHostPort that also parses the
// integer port. In addition, it is more tolerant of improperly escaped IPv6
// addresses, such as "::1:456", which should actually be "[::1]:456".
func SplitHostPort(addr string) (string, int, error) { _ = "STUB: not implemented"; return "", 0, nil }

// If the above proper parsing fails, fall back on a naive split.

// JoinHostPort is an extension to net.JoinHostPort that also formats the
// integer port.
func JoinHostPort(host string, port int32) string { _ = "STUB: not implemented"; return "" }

// FullyQualifiedHostname returns the FQDN of the machine.
func FullyQualifiedHostname() (string, error) {
	_ = "STUB: not implemented"
	// The machine hostname (which is also returned by os.Hostname()) may not be
	// set to the FQDN, but only the first part of it e.g. "localhost" instead of
	// "localhost.localdomain".
	// To get the full FQDN, we do the following:
	return "", nil
}

// 1. Get the machine hostname. Example: localhost

// 2. Look up the IP address for that hostname. Example: 127.0.0.1

// Prefer IPv4 addresses over IPv6 link-local addresses to avoid DNS timeout issues

// Prefer IPv4 addresses, or non-link-local IPv6

// This is an IPv4 address - prefer it

// This is a non-link-local IPv6 address - acceptable

// If no preferred IP found, fall back to the first one

// 3. Reverse lookup the IP. Example: localhost.localdomain

// If multiple hostnames are found, we return only the first one.
// If multiple hostnames are listed e.g. in an entry in the /etc/hosts file,
// the current Go implementation returns them in that order.
// Example for an /etc/hosts entry:
//   127.0.0.1	localhost.localdomain localhost
// If the FQDN isn't returned by this function, check the order in the entry
// in your /etc/hosts file.

// FullyQualifiedHostnameOrPanic is the same as FullyQualifiedHostname
// but panics in case of an error.
func FullyQualifiedHostnameOrPanic() string { _ = "STUB: not implemented"; return "" }

func dnsLookup(host string) ([]net.IP, error) { _ = "STUB: not implemented"; return nil, nil }

// DNSTracker is a closure that persists state for
//
//	tracking changes in the DNS resolution of a target dns name
//	returns true if the DNS name resolution has changed
//	If there is a lookup problem, we pretend nothing has changed
func DNSTracker(host string) func() (bool, error) { _ = "STUB: not implemented"; return nil }

// Should not happen, but just in case

// Update the closure variable

func addrEqual(a, b []net.IP) bool { _ = "STUB: not implemented"; return false }

// NormalizeIP normalizes loopback addresses to avoid spurious errors when
// communicating to different representations of the loopback.
//
// Note: this also maps IPv6 localhost to IPv4 localhost, as
// TabletManagerClient.GetReplicas() (the only place this function is used on)
// will return only IPv4 addresses.
func NormalizeIP(s string) string { _ = "STUB: not implemented"; return "" }
