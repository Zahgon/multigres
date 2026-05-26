// Copyright 2026 Supabase, Inc.
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

package protoutil

import (
	multipoolerpb "github.com/multigres/multigres/go/pb/multipoolerservice"
	querypb "github.com/multigres/multigres/go/pb/query"
)

// validReasonsMask is the bitmask of all known reservation reasons.
const validReasonsMask = ReasonTransaction | ReasonTempTable | ReasonPortal | ReasonCopy | ReasonListen | ReasonLogicalReplication

// Reason constants as uint32 for bitmask operations.
// These match the ReservationReason enum values.
const (
	ReasonTransaction = uint32(multipoolerpb.ReservationReason_RESERVATION_REASON_TRANSACTION) // 1
	ReasonTempTable   = uint32(multipoolerpb.ReservationReason_RESERVATION_REASON_TEMP_TABLE)  // 2
	ReasonPortal      = uint32(multipoolerpb.ReservationReason_RESERVATION_REASON_PORTAL)      // 4
	ReasonCopy        = uint32(multipoolerpb.ReservationReason_RESERVATION_REASON_COPY)        // 8
	ReasonListen      = uint32(multipoolerpb.ReservationReason_RESERVATION_REASON_LISTEN)      // 16

	// ReasonLogicalReplication indicates the connection has logical-replication
	// session state (an owned slot or an active replication-protocol stream)
	// and must stay pinned to its current Postgres backend for the session's
	// lifetime.
	//
	// Today this bit is set only at connection-open time, by the
	// reserved.Pool.NewLogicalReplicationConn factory for connections that
	// requested `replication=database` in the startup parameters.
	//
	// TODO: also set this bit mid-session when the planner observes
	// pg_create_logical_replication_slot(...) on a plain SQL connection
	// (polling / CDC-RLS path). This will mirror how ReasonTempTable is set
	// when CREATE TEMP TABLE is observed (see
	// go/services/multigateway/handler/connection_state.go for the
	// PendingTempTableReservation precedent). Until that lands, polling
	// consumers must cooperate by setting an application_name that
	// multigateway recognizes (also future work).
	ReasonLogicalReplication = uint32(multipoolerpb.ReservationReason_RESERVATION_REASON_LOGICAL_REPLICATION) // 32
)

// ValidateReasons returns an error if any unknown bits are set in the reasons bitmask.
func ValidateReasons(reasons uint32) error { _ = "STUB: not implemented"; return nil }

// HasReason returns true if the reasons bitmask contains the specified reason.
func HasReason(reasons uint32, reason uint32) bool { _ = "STUB: not implemented"; return false }

// HasTransactionReason returns true if the reasons bitmask includes transaction.
func HasTransactionReason(reasons uint32) bool { _ = "STUB: not implemented"; return false }

// HasTempTableReason returns true if the reasons bitmask includes temp table.
func HasTempTableReason(reasons uint32) bool { _ = "STUB: not implemented"; return false }

// HasPortalReason returns true if the reasons bitmask includes portal.
func HasPortalReason(reasons uint32) bool { _ = "STUB: not implemented"; return false }

// HasCopyReason returns true if the reasons bitmask includes an active COPY operation.
func HasCopyReason(reasons uint32) bool { _ = "STUB: not implemented"; return false }

// HasListenReason returns true if the reasons bitmask includes LISTEN/NOTIFY.
func HasListenReason(reasons uint32) bool { _ = "STUB: not implemented"; return false }

// HasLogicalReplicationReason returns true if the reasons bitmask includes a logical-replication session.
func HasLogicalReplicationReason(reasons uint32) bool { _ = "STUB: not implemented"; return false }

// AddReason adds a reason to the bitmask and returns the new value.
func AddReason(reasons uint32, reason uint32) uint32 { _ = "STUB: not implemented"; return 0 }

// RemoveReason removes a reason from the bitmask and returns the new value.
func RemoveReason(reasons uint32, reason uint32) uint32 { _ = "STUB: not implemented"; return 0 }

// RequiresBegin returns true if the reasons bitmask includes transaction reason.
// This determines if BEGIN should be executed when reserving.
func RequiresBegin(reasons uint32) bool { _ = "STUB: not implemented"; return false }

// IsEmpty returns true if no reasons are set (connection can be released).
func IsEmpty(reasons uint32) bool { _ = "STUB: not implemented"; return false }

// NewTransactionReservationOptions creates ReservationOptions for a transaction.
func NewTransactionReservationOptions() *querypb.ReservationOptions {
	_ = "STUB: not implemented"
	return nil
}

// NewTempTableReservationOptions creates ReservationOptions for temporary tables.
func NewTempTableReservationOptions() *querypb.ReservationOptions {
	_ = "STUB: not implemented"
	return nil
}

// NewPortalReservationOptions creates ReservationOptions for portal/cursor operations.
func NewPortalReservationOptions() *querypb.ReservationOptions {
	_ = "STUB: not implemented"
	return nil
}

// NewReservationOptions creates ReservationOptions with the given reasons bitmask.
func NewReservationOptions(reasons uint32) *querypb.ReservationOptions {
	_ = "STUB: not implemented"
	return nil
}

// GetReasons extracts the reasons bitmask from ReservationOptions, returning 0 if nil.
func GetReasons(opts *querypb.ReservationOptions) uint32 { _ = "STUB: not implemented"; return 0 }

// ReasonsString returns a human-readable string of the reasons bitmask.
func ReasonsString(reasons uint32) string { _ = "STUB: not implemented"; return "" }
