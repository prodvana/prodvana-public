// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: prodvana/application/user_metadata.proto

package application

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

func (m *ApplicationUserMetadata) CloneVT() *ApplicationUserMetadata {
	if m == nil {
		return (*ApplicationUserMetadata)(nil)
	}
	r := new(ApplicationUserMetadata)
	r.Description = m.Description
	if rhs := m.Links; rhs != nil {
		tmpContainer := make([]*common_config.Link, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Links = tmpContainer
	}
	if rhs := m.ServiceLinks; rhs != nil {
		tmpContainer := make([]*common_config.Link, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.ServiceLinks = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ApplicationUserMetadata) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *ApplicationUserMetadata) StableEqualVT(that *ApplicationUserMetadata) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Description != that.Description {
		return false
	}
	if len(this.Links) != len(that.Links) {
		return false
	}
	for i, vx := range this.Links {
		vy := that.Links[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &common_config.Link{}
			}
			if q == nil {
				q = &common_config.Link{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	if len(this.ServiceLinks) != len(that.ServiceLinks) {
		return false
	}
	for i, vx := range this.ServiceLinks {
		vy := that.ServiceLinks[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &common_config.Link{}
			}
			if q == nil {
				q = &common_config.Link{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ApplicationUserMetadata) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ApplicationUserMetadata)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *ApplicationUserMetadata) EqualVT(that *ApplicationUserMetadata) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Description != that.Description {
		return false
	}
	if len(this.Links) != len(that.Links) {
		return false
	}
	for i, vx := range this.Links {
		vy := that.Links[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &common_config.Link{}
			}
			if q == nil {
				q = &common_config.Link{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if len(this.ServiceLinks) != len(that.ServiceLinks) {
		return false
	}
	for i, vx := range this.ServiceLinks {
		vy := that.ServiceLinks[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &common_config.Link{}
			}
			if q == nil {
				q = &common_config.Link{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ApplicationUserMetadata) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ApplicationUserMetadata)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
