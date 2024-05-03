// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: prodvana/desired_state/debug/debug.proto

package debug

import (
	anypb1 "github.com/planetscale/vtprotobuf/types/known/anypb"
	timestamppb1 "github.com/planetscale/vtprotobuf/types/known/timestamppb"
	model "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/model"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *EntityDumpState_IOState) CloneVT() *EntityDumpState_IOState {
	if m == nil {
		return (*EntityDumpState_IOState)(nil)
	}
	r := new(EntityDumpState_IOState)
	r.LastFetchedTimestamp = (*timestamppb.Timestamp)((*timestamppb1.Timestamp)(m.LastFetchedTimestamp).CloneVT())
	r.LastAppliedTimestamp = (*timestamppb.Timestamp)((*timestamppb1.Timestamp)(m.LastAppliedTimestamp).CloneVT())
	r.ExpectedNextApplyTimestamp = (*timestamppb.Timestamp)((*timestamppb1.Timestamp)(m.ExpectedNextApplyTimestamp).CloneVT())
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *EntityDumpState_IOState) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *EntityDumpState_ChildState) CloneVT() *EntityDumpState_ChildState {
	if m == nil {
		return (*EntityDumpState_ChildState)(nil)
	}
	r := new(EntityDumpState_ChildState)
	r.Id = m.Id.CloneVT()
	r.Status = m.Status
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *EntityDumpState_ChildState) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *EntityDumpState) CloneVT() *EntityDumpState {
	if m == nil {
		return (*EntityDumpState)(nil)
	}
	r := new(EntityDumpState)
	r.Id = m.Id.CloneVT()
	r.LastManagerUpdateTimestamp = (*timestamppb.Timestamp)((*timestamppb1.Timestamp)(m.LastManagerUpdateTimestamp).CloneVT())
	r.Annotations = m.Annotations.CloneVT()
	r.Status = m.Status
	r.SimpleStatus = m.SimpleStatus
	r.Control = m.Control.CloneVT()
	r.IoState = m.IoState.CloneVT()
	r.DesiredState = (*anypb.Any)((*anypb1.Any)(m.DesiredState).CloneVT())
	r.FetchedState = (*anypb.Any)((*anypb1.Any)(m.FetchedState).CloneVT())
	r.TargetState = (*anypb.Any)((*anypb1.Any)(m.TargetState).CloneVT())
	r.Absent = m.Absent
	r.Deleted = m.Deleted
	r.Observer = m.Observer
	r.Stale = m.Stale
	if rhs := m.Logs; rhs != nil {
		tmpContainer := make([]*model.DebugLog, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Logs = tmpContainer
	}
	if rhs := m.ChildStates; rhs != nil {
		tmpContainer := make([]*EntityDumpState_ChildState, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.ChildStates = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *EntityDumpState) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *DumpState) CloneVT() *DumpState {
	if m == nil {
		return (*DumpState)(nil)
	}
	r := new(DumpState)
	if rhs := m.Entities; rhs != nil {
		tmpContainer := make([]*EntityDumpState, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Entities = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *DumpState) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *EntityDumpState_IOState) StableEqualVT(that *EntityDumpState_IOState) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.LastFetchedTimestamp).StableEqualVT((*timestamppb1.Timestamp)(that.LastFetchedTimestamp)) {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.LastAppliedTimestamp).StableEqualVT((*timestamppb1.Timestamp)(that.LastAppliedTimestamp)) {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.ExpectedNextApplyTimestamp).StableEqualVT((*timestamppb1.Timestamp)(that.ExpectedNextApplyTimestamp)) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *EntityDumpState_IOState) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*EntityDumpState_IOState)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *EntityDumpState_ChildState) StableEqualVT(that *EntityDumpState_ChildState) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Id.StableEqualVT(that.Id) {
		return false
	}
	if this.Status != that.Status {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *EntityDumpState_ChildState) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*EntityDumpState_ChildState)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *EntityDumpState) StableEqualVT(that *EntityDumpState) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Id.StableEqualVT(that.Id) {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.LastManagerUpdateTimestamp).StableEqualVT((*timestamppb1.Timestamp)(that.LastManagerUpdateTimestamp)) {
		return false
	}
	if !this.Annotations.StableEqualVT(that.Annotations) {
		return false
	}
	if this.Status != that.Status {
		return false
	}
	if !this.Control.StableEqualVT(that.Control) {
		return false
	}
	if !this.IoState.StableEqualVT(that.IoState) {
		return false
	}
	if len(this.Logs) != len(that.Logs) {
		return false
	}
	for i, vx := range this.Logs {
		vy := that.Logs[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &model.DebugLog{}
			}
			if q == nil {
				q = &model.DebugLog{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	if !(*anypb1.Any)(this.DesiredState).StableEqualVT((*anypb1.Any)(that.DesiredState)) {
		return false
	}
	if !(*anypb1.Any)(this.FetchedState).StableEqualVT((*anypb1.Any)(that.FetchedState)) {
		return false
	}
	if len(this.ChildStates) != len(that.ChildStates) {
		return false
	}
	for i, vx := range this.ChildStates {
		vy := that.ChildStates[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &EntityDumpState_ChildState{}
			}
			if q == nil {
				q = &EntityDumpState_ChildState{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	if this.SimpleStatus != that.SimpleStatus {
		return false
	}
	if !(*anypb1.Any)(this.TargetState).StableEqualVT((*anypb1.Any)(that.TargetState)) {
		return false
	}
	if this.Absent != that.Absent {
		return false
	}
	if this.Deleted != that.Deleted {
		return false
	}
	if this.Observer != that.Observer {
		return false
	}
	if this.Stale != that.Stale {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *EntityDumpState) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*EntityDumpState)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *EntityDumpState_IOState) EqualVT(that *EntityDumpState_IOState) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.LastFetchedTimestamp).EqualVT((*timestamppb1.Timestamp)(that.LastFetchedTimestamp)) {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.LastAppliedTimestamp).EqualVT((*timestamppb1.Timestamp)(that.LastAppliedTimestamp)) {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.ExpectedNextApplyTimestamp).EqualVT((*timestamppb1.Timestamp)(that.ExpectedNextApplyTimestamp)) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *EntityDumpState_IOState) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*EntityDumpState_IOState)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *EntityDumpState_ChildState) EqualVT(that *EntityDumpState_ChildState) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Id.EqualVT(that.Id) {
		return false
	}
	if this.Status != that.Status {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *EntityDumpState_ChildState) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*EntityDumpState_ChildState)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *EntityDumpState) EqualVT(that *EntityDumpState) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Id.EqualVT(that.Id) {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.LastManagerUpdateTimestamp).EqualVT((*timestamppb1.Timestamp)(that.LastManagerUpdateTimestamp)) {
		return false
	}
	if !this.Annotations.EqualVT(that.Annotations) {
		return false
	}
	if this.Status != that.Status {
		return false
	}
	if !this.Control.EqualVT(that.Control) {
		return false
	}
	if !this.IoState.EqualVT(that.IoState) {
		return false
	}
	if len(this.Logs) != len(that.Logs) {
		return false
	}
	for i, vx := range this.Logs {
		vy := that.Logs[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &model.DebugLog{}
			}
			if q == nil {
				q = &model.DebugLog{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if !(*anypb1.Any)(this.DesiredState).EqualVT((*anypb1.Any)(that.DesiredState)) {
		return false
	}
	if !(*anypb1.Any)(this.FetchedState).EqualVT((*anypb1.Any)(that.FetchedState)) {
		return false
	}
	if len(this.ChildStates) != len(that.ChildStates) {
		return false
	}
	for i, vx := range this.ChildStates {
		vy := that.ChildStates[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &EntityDumpState_ChildState{}
			}
			if q == nil {
				q = &EntityDumpState_ChildState{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if this.SimpleStatus != that.SimpleStatus {
		return false
	}
	if !(*anypb1.Any)(this.TargetState).EqualVT((*anypb1.Any)(that.TargetState)) {
		return false
	}
	if this.Absent != that.Absent {
		return false
	}
	if this.Deleted != that.Deleted {
		return false
	}
	if this.Observer != that.Observer {
		return false
	}
	if this.Stale != that.Stale {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *EntityDumpState) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*EntityDumpState)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *DumpState) StableEqualVT(that *DumpState) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.Entities) != len(that.Entities) {
		return false
	}
	for i, vx := range this.Entities {
		vy := that.Entities[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &EntityDumpState{}
			}
			if q == nil {
				q = &EntityDumpState{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *DumpState) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*DumpState)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *DumpState) EqualVT(that *DumpState) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.Entities) != len(that.Entities) {
		return false
	}
	for i, vx := range this.Entities {
		vy := that.Entities[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &EntityDumpState{}
			}
			if q == nil {
				q = &EntityDumpState{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *DumpState) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*DumpState)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}