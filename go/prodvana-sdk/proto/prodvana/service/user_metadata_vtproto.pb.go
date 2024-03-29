// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: prodvana/service/user_metadata.proto

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

func (m *ServiceUserMetadata) CloneVT() *ServiceUserMetadata {
	if m == nil {
		return (*ServiceUserMetadata)(nil)
	}
	r := new(ServiceUserMetadata)
	r.Description = m.Description
	r.FollowRepository = m.FollowRepository.CloneVT()
	if rhs := m.Links; rhs != nil {
		tmpContainer := make([]*common_config.Link, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Links = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ServiceUserMetadata) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *FollowContainerRepositorySettings) CloneVT() *FollowContainerRepositorySettings {
	if m == nil {
		return (*FollowContainerRepositorySettings)(nil)
	}
	r := new(FollowContainerRepositorySettings)
	r.Enabled = m.Enabled
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *FollowContainerRepositorySettings) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *ServiceUserMetadata) StableEqualVT(that *ServiceUserMetadata) bool {
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
	if !this.FollowRepository.StableEqualVT(that.FollowRepository) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ServiceUserMetadata) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ServiceUserMetadata)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *ServiceUserMetadata) EqualVT(that *ServiceUserMetadata) bool {
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
	if !this.FollowRepository.EqualVT(that.FollowRepository) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ServiceUserMetadata) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ServiceUserMetadata)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *FollowContainerRepositorySettings) StableEqualVT(that *FollowContainerRepositorySettings) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Enabled != that.Enabled {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *FollowContainerRepositorySettings) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*FollowContainerRepositorySettings)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *FollowContainerRepositorySettings) EqualVT(that *FollowContainerRepositorySettings) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Enabled != that.Enabled {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *FollowContainerRepositorySettings) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*FollowContainerRepositorySettings)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
