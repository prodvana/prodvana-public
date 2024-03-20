// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: prodvana/common_config/rollback.proto

package common_config

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

func (m *AutoRollbackConfig) CloneVT() *AutoRollbackConfig {
	if m == nil {
		return (*AutoRollbackConfig)(nil)
	}
	r := new(AutoRollbackConfig)
	r.Disabled = m.Disabled
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *AutoRollbackConfig) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *AutoRollbackConfig) EqualVT(that *AutoRollbackConfig) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Disabled != that.Disabled {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *AutoRollbackConfig) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*AutoRollbackConfig)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
