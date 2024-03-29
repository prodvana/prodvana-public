// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: prodvana/blobs/blobs_manager.proto

package blobs

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

func (m *GetCasBlobReq) CloneVT() *GetCasBlobReq {
	if m == nil {
		return (*GetCasBlobReq)(nil)
	}
	r := new(GetCasBlobReq)
	r.Id = m.Id
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *GetCasBlobReq) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *GetCasBlobResp) CloneVT() *GetCasBlobResp {
	if m == nil {
		return (*GetCasBlobResp)(nil)
	}
	r := new(GetCasBlobResp)
	if rhs := m.Bytes; rhs != nil {
		tmpBytes := make([]byte, len(rhs))
		copy(tmpBytes, rhs)
		r.Bytes = tmpBytes
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *GetCasBlobResp) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *UploadCasBlobReq) CloneVT() *UploadCasBlobReq {
	if m == nil {
		return (*UploadCasBlobReq)(nil)
	}
	r := new(UploadCasBlobReq)
	if rhs := m.Bytes; rhs != nil {
		tmpBytes := make([]byte, len(rhs))
		copy(tmpBytes, rhs)
		r.Bytes = tmpBytes
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *UploadCasBlobReq) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *UploadCasBlobResp) CloneVT() *UploadCasBlobResp {
	if m == nil {
		return (*UploadCasBlobResp)(nil)
	}
	r := new(UploadCasBlobResp)
	r.Id = m.Id
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *UploadCasBlobResp) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *GetCasBlobReq) StableEqualVT(that *GetCasBlobReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Id != that.Id {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetCasBlobReq) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetCasBlobReq)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *GetCasBlobReq) EqualVT(that *GetCasBlobReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Id != that.Id {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetCasBlobReq) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetCasBlobReq)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *GetCasBlobResp) StableEqualVT(that *GetCasBlobResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if string(this.Bytes) != string(that.Bytes) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetCasBlobResp) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetCasBlobResp)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *GetCasBlobResp) EqualVT(that *GetCasBlobResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if string(this.Bytes) != string(that.Bytes) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetCasBlobResp) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetCasBlobResp)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *UploadCasBlobReq) StableEqualVT(that *UploadCasBlobReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if string(this.Bytes) != string(that.Bytes) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *UploadCasBlobReq) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*UploadCasBlobReq)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *UploadCasBlobReq) EqualVT(that *UploadCasBlobReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if string(this.Bytes) != string(that.Bytes) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *UploadCasBlobReq) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*UploadCasBlobReq)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *UploadCasBlobResp) StableEqualVT(that *UploadCasBlobResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Id != that.Id {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *UploadCasBlobResp) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*UploadCasBlobResp)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *UploadCasBlobResp) EqualVT(that *UploadCasBlobResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Id != that.Id {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *UploadCasBlobResp) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*UploadCasBlobResp)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
