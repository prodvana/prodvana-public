// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: prodvana/secrets/secrets_manager.proto

package secrets

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

func (m *ListSecretsReq) CloneVT() *ListSecretsReq {
	if m == nil {
		return (*ListSecretsReq)(nil)
	}
	r := new(ListSecretsReq)
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ListSecretsReq) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ListSecretsResp) CloneVT() *ListSecretsResp {
	if m == nil {
		return (*ListSecretsResp)(nil)
	}
	r := new(ListSecretsResp)
	if rhs := m.Secrets; rhs != nil {
		tmpContainer := make([]*common_config.Secret, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Secrets = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ListSecretsResp) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ListSecretVersionsReq) CloneVT() *ListSecretVersionsReq {
	if m == nil {
		return (*ListSecretVersionsReq)(nil)
	}
	r := new(ListSecretVersionsReq)
	r.Key = m.Key
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ListSecretVersionsReq) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ListSecretVersionsResp) CloneVT() *ListSecretVersionsResp {
	if m == nil {
		return (*ListSecretVersionsResp)(nil)
	}
	r := new(ListSecretVersionsResp)
	if rhs := m.Versions; rhs != nil {
		tmpContainer := make([]string, len(rhs))
		copy(tmpContainer, rhs)
		r.Versions = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ListSecretVersionsResp) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *SetSecretReq) CloneVT() *SetSecretReq {
	if m == nil {
		return (*SetSecretReq)(nil)
	}
	r := new(SetSecretReq)
	r.Key = m.Key
	r.Value = m.Value
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *SetSecretReq) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *SetSecretResp) CloneVT() *SetSecretResp {
	if m == nil {
		return (*SetSecretResp)(nil)
	}
	r := new(SetSecretResp)
	r.Version = m.Version
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *SetSecretResp) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *DeleteSecretReq) CloneVT() *DeleteSecretReq {
	if m == nil {
		return (*DeleteSecretReq)(nil)
	}
	r := new(DeleteSecretReq)
	r.Key = m.Key
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *DeleteSecretReq) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *DeleteSecretResp) CloneVT() *DeleteSecretResp {
	if m == nil {
		return (*DeleteSecretResp)(nil)
	}
	r := new(DeleteSecretResp)
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *DeleteSecretResp) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *DeleteSecretVersionReq) CloneVT() *DeleteSecretVersionReq {
	if m == nil {
		return (*DeleteSecretVersionReq)(nil)
	}
	r := new(DeleteSecretVersionReq)
	r.Key = m.Key
	r.Version = m.Version
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *DeleteSecretVersionReq) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *DeleteSecretVersionResp) CloneVT() *DeleteSecretVersionResp {
	if m == nil {
		return (*DeleteSecretVersionResp)(nil)
	}
	r := new(DeleteSecretVersionResp)
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *DeleteSecretVersionResp) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *ListSecretsReq) EqualVT(that *ListSecretsReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ListSecretsReq) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ListSecretsReq)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ListSecretsResp) EqualVT(that *ListSecretsResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.Secrets) != len(that.Secrets) {
		return false
	}
	for i, vx := range this.Secrets {
		vy := that.Secrets[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &common_config.Secret{}
			}
			if q == nil {
				q = &common_config.Secret{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ListSecretsResp) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ListSecretsResp)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ListSecretVersionsReq) EqualVT(that *ListSecretVersionsReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Key != that.Key {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ListSecretVersionsReq) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ListSecretVersionsReq)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ListSecretVersionsResp) EqualVT(that *ListSecretVersionsResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.Versions) != len(that.Versions) {
		return false
	}
	for i, vx := range this.Versions {
		vy := that.Versions[i]
		if vx != vy {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ListSecretVersionsResp) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ListSecretVersionsResp)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *SetSecretReq) EqualVT(that *SetSecretReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Key != that.Key {
		return false
	}
	if this.Value != that.Value {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *SetSecretReq) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*SetSecretReq)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *SetSecretResp) EqualVT(that *SetSecretResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Version != that.Version {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *SetSecretResp) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*SetSecretResp)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *DeleteSecretReq) EqualVT(that *DeleteSecretReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Key != that.Key {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *DeleteSecretReq) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*DeleteSecretReq)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *DeleteSecretResp) EqualVT(that *DeleteSecretResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *DeleteSecretResp) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*DeleteSecretResp)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *DeleteSecretVersionReq) EqualVT(that *DeleteSecretVersionReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Key != that.Key {
		return false
	}
	if this.Version != that.Version {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *DeleteSecretVersionReq) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*DeleteSecretVersionReq)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *DeleteSecretVersionResp) EqualVT(that *DeleteSecretVersionResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *DeleteSecretVersionResp) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*DeleteSecretVersionResp)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
