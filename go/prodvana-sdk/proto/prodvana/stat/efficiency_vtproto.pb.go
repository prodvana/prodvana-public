// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: prodvana/stat/efficiency.proto

package stat

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

func (m *EfficiencyStat) CloneVT() *EfficiencyStat {
	if m == nil {
		return (*EfficiencyStat)(nil)
	}
	r := new(EfficiencyStat)
	r.SavedLines = m.SavedLines
	r.MaterializedLines = m.MaterializedLines
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *EfficiencyStat) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *EfficiencyStat) EqualVT(that *EfficiencyStat) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.SavedLines != that.SavedLines {
		return false
	}
	if this.MaterializedLines != that.MaterializedLines {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *EfficiencyStat) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*EfficiencyStat)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
