// Copyright 2026 Supabase, Inc.
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

package eventlog

import "log/slog"

type NodeJoin struct{ NodeName string }

func (NodeJoin) EventType() string       { _ = "STUB: not implemented"; return "" }
func (e NodeJoin) LogAttrs() []slog.Attr { _ = "STUB: not implemented"; return nil }

type PrimaryPromotion struct{ NewPrimary string }

func (PrimaryPromotion) EventType() string       { _ = "STUB: not implemented"; return "" }
func (e PrimaryPromotion) LogAttrs() []slog.Attr { _ = "STUB: not implemented"; return nil }

type BackupAttempt struct{ BackupName string }

func (BackupAttempt) EventType() string       { _ = "STUB: not implemented"; return "" }
func (e BackupAttempt) LogAttrs() []slog.Attr { _ = "STUB: not implemented"; return nil }

type RestoreAttempt struct {
	BackupName string
}

func (RestoreAttempt) EventType() string       { _ = "STUB: not implemented"; return "" }
func (e RestoreAttempt) LogAttrs() []slog.Attr { _ = "STUB: not implemented"; return nil }

type PrimaryDemotion struct {
	NodeName string
	Reason   string // "stale" | "emergency"
}

func (PrimaryDemotion) EventType() string       { _ = "STUB: not implemented"; return "" }
func (e PrimaryDemotion) LogAttrs() []slog.Attr { _ = "STUB: not implemented"; return nil }

type NodeDrain struct {
	NodeName string
	Reason   string // e.g. "rewind_not_feasible"
}

func (NodeDrain) EventType() string       { _ = "STUB: not implemented"; return "" }
func (e NodeDrain) LogAttrs() []slog.Attr { _ = "STUB: not implemented"; return nil }

type BackupLeaseStolen struct {
	Stealer string
}

func (BackupLeaseStolen) EventType() string       { _ = "STUB: not implemented"; return "" }
func (e BackupLeaseStolen) LogAttrs() []slog.Attr { _ = "STUB: not implemented"; return nil }

type BackupLeaseLost struct {
	Holder string
}

func (BackupLeaseLost) EventType() string       { _ = "STUB: not implemented"; return "" }
func (e BackupLeaseLost) LogAttrs() []slog.Attr { _ = "STUB: not implemented"; return nil }

type TermBegin struct {
	NewTerm      int64
	PreviousTerm int64
	RevokedRole  string // "primary" | "standby" | "" (empty = no revoke)
}

func (TermBegin) EventType() string       { _ = "STUB: not implemented"; return "" }
func (e TermBegin) LogAttrs() []slog.Attr { _ = "STUB: not implemented"; return nil }
