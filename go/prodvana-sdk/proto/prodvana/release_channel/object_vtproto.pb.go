// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: prodvana/release_channel/object.proto

package release_channel

import (
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *ReleaseChannelProtectionAttachment) CloneVT() *ReleaseChannelProtectionAttachment {
	if m == nil {
		return (*ReleaseChannelProtectionAttachment)(nil)
	}
	r := new(ReleaseChannelProtectionAttachment)
	r.Protection = m.Protection
	r.Attachment = m.Attachment
	r.DesiredStateId = m.DesiredStateId
	r.AttachmentId = m.AttachmentId
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ReleaseChannelProtectionAttachment) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ReleaseChannelState) CloneVT() *ReleaseChannelState {
	if m == nil {
		return (*ReleaseChannelState)(nil)
	}
	r := new(ReleaseChannelState)
	if rhs := m.ProtectionAttachments; rhs != nil {
		tmpContainer := make([]*ReleaseChannelProtectionAttachment, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.ProtectionAttachments = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ReleaseChannelState) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ReleaseChannel) CloneVT() *ReleaseChannel {
	if m == nil {
		return (*ReleaseChannel)(nil)
	}
	r := new(ReleaseChannel)
	r.Meta = m.Meta.CloneVT()
	r.Config = m.Config.CloneVT()
	r.State = m.State.CloneVT()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ReleaseChannel) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *ReleaseChannelProtectionAttachment) StableEqualVT(that *ReleaseChannelProtectionAttachment) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Protection != that.Protection {
		return false
	}
	if this.Attachment != that.Attachment {
		return false
	}
	if this.DesiredStateId != that.DesiredStateId {
		return false
	}
	if this.AttachmentId != that.AttachmentId {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ReleaseChannelProtectionAttachment) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ReleaseChannelProtectionAttachment)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *ReleaseChannelProtectionAttachment) EqualVT(that *ReleaseChannelProtectionAttachment) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Protection != that.Protection {
		return false
	}
	if this.Attachment != that.Attachment {
		return false
	}
	if this.DesiredStateId != that.DesiredStateId {
		return false
	}
	if this.AttachmentId != that.AttachmentId {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ReleaseChannelProtectionAttachment) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ReleaseChannelProtectionAttachment)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ReleaseChannelState) StableEqualVT(that *ReleaseChannelState) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.ProtectionAttachments) != len(that.ProtectionAttachments) {
		return false
	}
	for i, vx := range this.ProtectionAttachments {
		vy := that.ProtectionAttachments[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &ReleaseChannelProtectionAttachment{}
			}
			if q == nil {
				q = &ReleaseChannelProtectionAttachment{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ReleaseChannelState) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ReleaseChannelState)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *ReleaseChannelState) EqualVT(that *ReleaseChannelState) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.ProtectionAttachments) != len(that.ProtectionAttachments) {
		return false
	}
	for i, vx := range this.ProtectionAttachments {
		vy := that.ProtectionAttachments[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &ReleaseChannelProtectionAttachment{}
			}
			if q == nil {
				q = &ReleaseChannelProtectionAttachment{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ReleaseChannelState) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ReleaseChannelState)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ReleaseChannel) StableEqualVT(that *ReleaseChannel) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Meta.StableEqualVT(that.Meta) {
		return false
	}
	if !this.Config.StableEqualVT(that.Config) {
		return false
	}
	if !this.State.StableEqualVT(that.State) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ReleaseChannel) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ReleaseChannel)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *ReleaseChannel) EqualVT(that *ReleaseChannel) bool {
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

func (this *ReleaseChannel) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ReleaseChannel)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
