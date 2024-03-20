// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: prodvana/desired_state/maestro/maestro.proto

package maestro

import (
	timestamppb1 "github.com/planetscale/vtprotobuf/types/known/timestamppb"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *MaestroConfig) CloneVT() *MaestroConfig {
	if m == nil {
		return (*MaestroConfig)(nil)
	}
	r := new(MaestroConfig)
	r.Strategy = m.Strategy
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *MaestroConfig) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *MaestroReleaseConfig) CloneVT() *MaestroReleaseConfig {
	if m == nil {
		return (*MaestroReleaseConfig)(nil)
	}
	r := new(MaestroReleaseConfig)
	r.EntityId = m.EntityId.CloneVT()
	r.MaestroConfig = m.MaestroConfig.CloneVT()
	r.InputDesiredState = m.InputDesiredState.CloneVT()
	r.CompiledDesiredState = m.CompiledDesiredState.CloneVT()
	r.CreationTimestamp = (*timestamppb.Timestamp)((*timestamppb1.Timestamp)(m.CreationTimestamp).CloneVT())
	r.SetDesiredStateMetadata = m.SetDesiredStateMetadata.CloneVT()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *MaestroReleaseConfig) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *MaestroReleaseState) CloneVT() *MaestroReleaseState {
	if m == nil {
		return (*MaestroReleaseState)(nil)
	}
	r := new(MaestroReleaseState)
	r.Status = m.Status
	r.LastUpdateTimestamp = (*timestamppb.Timestamp)((*timestamppb1.Timestamp)(m.LastUpdateTimestamp).CloneVT())
	if rhs := m.ReleaseChannelStatuses; rhs != nil {
		tmpContainer := make(map[string]Status, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v
		}
		r.ReleaseChannelStatuses = tmpContainer
	}
	if rhs := m.ReleaseChannelStates; rhs != nil {
		tmpContainer := make(map[string]*MaestroReleaseChannelState, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.ReleaseChannelStates = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *MaestroReleaseState) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *MaestroReleaseChannelState) CloneVT() *MaestroReleaseChannelState {
	if m == nil {
		return (*MaestroReleaseChannelState)(nil)
	}
	r := new(MaestroReleaseChannelState)
	r.Status = m.Status
	r.DesiredStateId = m.DesiredStateId
	r.EntityId = m.EntityId
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *MaestroReleaseChannelState) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *MaestroRelease) CloneVT() *MaestroRelease {
	if m == nil {
		return (*MaestroRelease)(nil)
	}
	r := new(MaestroRelease)
	r.Meta = m.Meta.CloneVT()
	r.Config = m.Config.CloneVT()
	r.State = m.State.CloneVT()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *MaestroRelease) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *MaestroConfig) EqualVT(that *MaestroConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Strategy != that.Strategy {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *MaestroConfig) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*MaestroConfig)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *MaestroReleaseConfig) EqualVT(that *MaestroReleaseConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.EntityId.EqualVT(that.EntityId) {
		return false
	}
	if !this.MaestroConfig.EqualVT(that.MaestroConfig) {
		return false
	}
	if !this.CompiledDesiredState.EqualVT(that.CompiledDesiredState) {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.CreationTimestamp).EqualVT((*timestamppb1.Timestamp)(that.CreationTimestamp)) {
		return false
	}
	if !this.InputDesiredState.EqualVT(that.InputDesiredState) {
		return false
	}
	if !this.SetDesiredStateMetadata.EqualVT(that.SetDesiredStateMetadata) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *MaestroReleaseConfig) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*MaestroReleaseConfig)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *MaestroReleaseState) EqualVT(that *MaestroReleaseState) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Status != that.Status {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.LastUpdateTimestamp).EqualVT((*timestamppb1.Timestamp)(that.LastUpdateTimestamp)) {
		return false
	}
	if len(this.ReleaseChannelStatuses) != len(that.ReleaseChannelStatuses) {
		return false
	}
	for i, vx := range this.ReleaseChannelStatuses {
		vy, ok := that.ReleaseChannelStatuses[i]
		if !ok {
			return false
		}
		if vx != vy {
			return false
		}
	}
	if len(this.ReleaseChannelStates) != len(that.ReleaseChannelStates) {
		return false
	}
	for i, vx := range this.ReleaseChannelStates {
		vy, ok := that.ReleaseChannelStates[i]
		if !ok {
			return false
		}
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &MaestroReleaseChannelState{}
			}
			if q == nil {
				q = &MaestroReleaseChannelState{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *MaestroReleaseState) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*MaestroReleaseState)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *MaestroReleaseChannelState) EqualVT(that *MaestroReleaseChannelState) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Status != that.Status {
		return false
	}
	if this.DesiredStateId != that.DesiredStateId {
		return false
	}
	if this.EntityId != that.EntityId {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *MaestroReleaseChannelState) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*MaestroReleaseChannelState)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *MaestroRelease) EqualVT(that *MaestroRelease) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Meta.EqualVT(that.Meta) {
		return false
	}
	if !this.Config.EqualVT(that.Config) {
		return false
	}
	if !this.State.EqualVT(that.State) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *MaestroRelease) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*MaestroRelease)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
