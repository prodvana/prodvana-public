// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: prodvana/object/meta.proto

package object

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

func (m *ObjectMeta) CloneVT() *ObjectMeta {
	if m == nil {
		return (*ObjectMeta)(nil)
	}
	r := new(ObjectMeta)
	r.Id = m.Id
	r.Name = m.Name
	r.Version = m.Version
	r.ConfigVersion = m.ConfigVersion
	r.Source = m.Source
	r.SourceMetadata = m.SourceMetadata.CloneVT()
	r.Type = m.Type
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ObjectMeta) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *ObjectMeta) StableEqualVT(that *ObjectMeta) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Id != that.Id {
		return false
	}
	if this.Name != that.Name {
		return false
	}
	if this.Version != that.Version {
		return false
	}
	if this.ConfigVersion != that.ConfigVersion {
		return false
	}
	if this.Source != that.Source {
		return false
	}
	if !this.SourceMetadata.StableEqualVT(that.SourceMetadata) {
		return false
	}
	if this.Type != that.Type {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ObjectMeta) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ObjectMeta)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *ObjectMeta) EqualVT(that *ObjectMeta) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Id != that.Id {
		return false
	}
	if this.Name != that.Name {
		return false
	}
	if this.Version != that.Version {
		return false
	}
	if this.ConfigVersion != that.ConfigVersion {
		return false
	}
	if this.Source != that.Source {
		return false
	}
	if !this.SourceMetadata.EqualVT(that.SourceMetadata) {
		return false
	}
	if this.Type != that.Type {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ObjectMeta) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ObjectMeta)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
