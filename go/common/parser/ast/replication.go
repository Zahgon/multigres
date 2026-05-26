// PostgreSQL Database Management System
// (also known as Postgres, formerly known as Postgres95)
//
//	Portions Copyright (c) 2026, Supabase, Inc
//
//	Portions Copyright (c) 1996-2025, PostgreSQL Global Development Group
//
//	Portions Copyright (c) 1994, The Regents of the University of California
//
// Permission to use, copy, modify, and distribute this software and its
// documentation for any purpose, without fee, and without a written agreement
// is hereby granted, provided that the above copyright notice and this
// paragraph and the following two paragraphs appear in all copies.
//
// IN NO EVENT SHALL THE UNIVERSITY OF CALIFORNIA BE LIABLE TO ANY PARTY FOR
// DIRECT, INDIRECT, SPECIAL, INCIDENTAL, OR CONSEQUENTIAL DAMAGES, INCLUDING
// LOST PROFITS, ARISING OUT OF THE USE OF THIS SOFTWARE AND ITS
// DOCUMENTATION, EVEN IF THE UNIVERSITY OF CALIFORNIA HAS BEEN ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.
//
// THE UNIVERSITY OF CALIFORNIA SPECIFICALLY DISCLAIMS ANY WARRANTIES,
// INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY
// AND FITNESS FOR A PARTICULAR PURPOSE.  THE SOFTWARE PROVIDED HEREUNDER IS
// ON AN "AS IS" BASIS, AND THE UNIVERSITY OF CALIFORNIA HAS NO OBLIGATIONS TO
// PROVIDE MAINTENANCE, SUPPORT, UPDATES, ENHANCEMENTS, OR MODIFICATIONS.
//
// Package ast: replication command nodes.
// Ported from postgres/src/include/nodes/replnodes.h (PG 17.6).
package ast

// ReplicationKind - ported from postgres/src/include/nodes/replnodes.h:20-24
type ReplicationKind int

const (
	ReplicationKindPhysical ReplicationKind = iota // replnodes.h:22
	ReplicationKindLogical                         // replnodes.h:23
)

// XLogRecPtr - WAL record pointer.
// Ported from postgres/src/include/access/xlogdefs.h:21
type XLogRecPtr uint64

// TimeLineID - timeline identifier.
// Ported from postgres/src/include/access/xlogdefs.h:59
type TimeLineID uint32

// IdentifySystemCmd - IDENTIFY_SYSTEM command.
// Ported from postgres/src/include/nodes/replnodes.h:31-34
type IdentifySystemCmd struct {
	BaseNode
}

func NewIdentifySystemCmd() *IdentifySystemCmd { _ = "STUB: not implemented"; return nil }

func (n *IdentifySystemCmd) StatementType() string { _ = "STUB: not implemented"; return "" }
func (n *IdentifySystemCmd) SqlString() string     { _ = "STUB: not implemented"; return "" }

// CreateReplicationSlotCmd - CREATE_REPLICATION_SLOT command.
// Ported from postgres/src/include/nodes/replnodes.h:52-60
type CreateReplicationSlotCmd struct {
	BaseNode
	SlotName  string          // replnodes.h:55
	Kind      ReplicationKind // replnodes.h:56
	Plugin    string          // replnodes.h:57 (LOGICAL only)
	Temporary bool            // replnodes.h:58
	Options   []*DefElem      // replnodes.h:59
}

func NewCreateReplicationSlotCmd(slot string, kind ReplicationKind, plugin string, temp bool) *CreateReplicationSlotCmd {
	_ = "STUB: not implemented"
	return nil
}

func (n *CreateReplicationSlotCmd) StatementType() string { _ = "STUB: not implemented"; return "" }
func (n *CreateReplicationSlotCmd) SqlString() string     { _ = "STUB: not implemented"; return "" }

// DropReplicationSlotCmd - DROP_REPLICATION_SLOT command.
// Ported from postgres/src/include/nodes/replnodes.h:67-72
type DropReplicationSlotCmd struct {
	BaseNode
	SlotName string // replnodes.h:70
	Wait     bool   // replnodes.h:71
}

func NewDropReplicationSlotCmd(slot string, wait bool) *DropReplicationSlotCmd {
	_ = "STUB: not implemented"
	return nil
}

func (n *DropReplicationSlotCmd) StatementType() string { _ = "STUB: not implemented"; return "" }
func (n *DropReplicationSlotCmd) SqlString() string     { _ = "STUB: not implemented"; return "" }

// AlterReplicationSlotCmd - ALTER_REPLICATION_SLOT command.
// Ported from postgres/src/include/nodes/replnodes.h:79-84
type AlterReplicationSlotCmd struct {
	BaseNode
	SlotName string     // replnodes.h:82
	Options  []*DefElem // replnodes.h:83
}

func NewAlterReplicationSlotCmd(slot string, options []*DefElem) *AlterReplicationSlotCmd {
	_ = "STUB: not implemented"
	return nil
}

func (n *AlterReplicationSlotCmd) StatementType() string { _ = "STUB: not implemented"; return "" }
func (n *AlterReplicationSlotCmd) SqlString() string     { _ = "STUB: not implemented"; return "" }

// StartReplicationCmd - START_REPLICATION command.
// Ported from postgres/src/include/nodes/replnodes.h:91-99
type StartReplicationCmd struct {
	BaseNode
	Kind       ReplicationKind // replnodes.h:94
	SlotName   string          // replnodes.h:95
	Timeline   TimeLineID      // replnodes.h:96
	StartPoint XLogRecPtr      // replnodes.h:97
	Options    []*DefElem      // replnodes.h:98
}

func NewStartReplicationCmd(kind ReplicationKind, slot string, timeline TimeLineID, start XLogRecPtr, options []*DefElem) *StartReplicationCmd {
	_ = "STUB: not implemented"
	return nil
}

func (n *StartReplicationCmd) StatementType() string { _ = "STUB: not implemented"; return "" }
func (n *StartReplicationCmd) SqlString() string     { _ = "STUB: not implemented"; return "" }

// ReadReplicationSlotCmd - READ_REPLICATION_SLOT command.
// Ported from postgres/src/include/nodes/replnodes.h:106-110
type ReadReplicationSlotCmd struct {
	BaseNode
	SlotName string // replnodes.h:109
}

func NewReadReplicationSlotCmd(slot string) *ReadReplicationSlotCmd {
	_ = "STUB: not implemented"
	return nil
}

func (n *ReadReplicationSlotCmd) StatementType() string { _ = "STUB: not implemented"; return "" }
func (n *ReadReplicationSlotCmd) SqlString() string     { _ = "STUB: not implemented"; return "" }
