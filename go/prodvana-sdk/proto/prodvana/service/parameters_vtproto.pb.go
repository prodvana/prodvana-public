// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: prodvana/service/parameters.proto

package service

import (
	common_config "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *PerReleaseChannelParameterValues) CloneVT() *PerReleaseChannelParameterValues {
	if m == nil {
		return (*PerReleaseChannelParameterValues)(nil)
	}
	r := new(PerReleaseChannelParameterValues)
	r.ReleaseChannel = m.ReleaseChannel
	if rhs := m.Parameters; rhs != nil {
		tmpContainer := make([]*common_config.ParameterValue, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Parameters = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *PerReleaseChannelParameterValues) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ServiceParameterValues) CloneVT() *ServiceParameterValues {
	if m == nil {
		return (*ServiceParameterValues)(nil)
	}
	r := new(ServiceParameterValues)
	if rhs := m.Parameters; rhs != nil {
		tmpContainer := make([]*common_config.ParameterValue, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Parameters = tmpContainer
	}
	if rhs := m.PerReleaseChannel; rhs != nil {
		tmpContainer := make([]*PerReleaseChannelParameterValues, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.PerReleaseChannel = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ServiceParameterValues) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *PerReleaseChannelParameterValues) StableEqualVT(that *PerReleaseChannelParameterValues) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.ReleaseChannel != that.ReleaseChannel {
		return false
	}
	if len(this.Parameters) != len(that.Parameters) {
		return false
	}
	for i, vx := range this.Parameters {
		vy := that.Parameters[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &common_config.ParameterValue{}
			}
			if q == nil {
				q = &common_config.ParameterValue{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *PerReleaseChannelParameterValues) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*PerReleaseChannelParameterValues)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *PerReleaseChannelParameterValues) EqualVT(that *PerReleaseChannelParameterValues) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.ReleaseChannel != that.ReleaseChannel {
		return false
	}
	if len(this.Parameters) != len(that.Parameters) {
		return false
	}
	for i, vx := range this.Parameters {
		vy := that.Parameters[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &common_config.ParameterValue{}
			}
			if q == nil {
				q = &common_config.ParameterValue{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *PerReleaseChannelParameterValues) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*PerReleaseChannelParameterValues)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ServiceParameterValues) StableEqualVT(that *ServiceParameterValues) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.Parameters) != len(that.Parameters) {
		return false
	}
	for i, vx := range this.Parameters {
		vy := that.Parameters[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &common_config.ParameterValue{}
			}
			if q == nil {
				q = &common_config.ParameterValue{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	if len(this.PerReleaseChannel) != len(that.PerReleaseChannel) {
		return false
	}
	for i, vx := range this.PerReleaseChannel {
		vy := that.PerReleaseChannel[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &PerReleaseChannelParameterValues{}
			}
			if q == nil {
				q = &PerReleaseChannelParameterValues{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ServiceParameterValues) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ServiceParameterValues)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *ServiceParameterValues) EqualVT(that *ServiceParameterValues) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.Parameters) != len(that.Parameters) {
		return false
	}
	for i, vx := range this.Parameters {
		vy := that.Parameters[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &common_config.ParameterValue{}
			}
			if q == nil {
				q = &common_config.ParameterValue{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if len(this.PerReleaseChannel) != len(that.PerReleaseChannel) {
		return false
	}
	for i, vx := range this.PerReleaseChannel {
		vy := that.PerReleaseChannel[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &PerReleaseChannelParameterValues{}
			}
			if q == nil {
				q = &PerReleaseChannelParameterValues{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ServiceParameterValues) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ServiceParameterValues)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
