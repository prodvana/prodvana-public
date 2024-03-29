// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: prodvana/protection/protection_manager.proto

package protection

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

func (m *ConfigureProtectionReq) CloneVT() *ConfigureProtectionReq {
	if m == nil {
		return (*ConfigureProtectionReq)(nil)
	}
	r := new(ConfigureProtectionReq)
	r.ProtectionConfig = m.ProtectionConfig.CloneVT()
	r.Source = m.Source
	r.SourceMetadata = m.SourceMetadata.CloneVT()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ConfigureProtectionReq) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ConfigureProtectionResp) CloneVT() *ConfigureProtectionResp {
	if m == nil {
		return (*ConfigureProtectionResp)(nil)
	}
	r := new(ConfigureProtectionResp)
	r.ProtectionId = m.ProtectionId
	r.Version = m.Version
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ConfigureProtectionResp) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ValidateConfigureProtectionResp) CloneVT() *ValidateConfigureProtectionResp {
	if m == nil {
		return (*ValidateConfigureProtectionResp)(nil)
	}
	r := new(ValidateConfigureProtectionResp)
	r.InputConfig = m.InputConfig.CloneVT()
	r.CompiledConfig = m.CompiledConfig.CloneVT()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ValidateConfigureProtectionResp) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ListProtectionsReq) CloneVT() *ListProtectionsReq {
	if m == nil {
		return (*ListProtectionsReq)(nil)
	}
	r := new(ListProtectionsReq)
	r.PageToken = m.PageToken
	r.PageSize = m.PageSize
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ListProtectionsReq) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *ListProtectionsResp) CloneVT() *ListProtectionsResp {
	if m == nil {
		return (*ListProtectionsResp)(nil)
	}
	r := new(ListProtectionsResp)
	r.NextPageToken = m.NextPageToken
	if rhs := m.Protections; rhs != nil {
		tmpContainer := make([]*Protection, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Protections = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *ListProtectionsResp) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *GetProtectionReq) CloneVT() *GetProtectionReq {
	if m == nil {
		return (*GetProtectionReq)(nil)
	}
	r := new(GetProtectionReq)
	r.Protection = m.Protection
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *GetProtectionReq) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *GetProtectionResp) CloneVT() *GetProtectionResp {
	if m == nil {
		return (*GetProtectionResp)(nil)
	}
	r := new(GetProtectionResp)
	r.Protection = m.Protection.CloneVT()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *GetProtectionResp) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *GetProtectionConfigReq) CloneVT() *GetProtectionConfigReq {
	if m == nil {
		return (*GetProtectionConfigReq)(nil)
	}
	r := new(GetProtectionConfigReq)
	r.Protection = m.Protection
	r.Version = m.Version
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *GetProtectionConfigReq) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *GetProtectionConfigResp) CloneVT() *GetProtectionConfigResp {
	if m == nil {
		return (*GetProtectionConfigResp)(nil)
	}
	r := new(GetProtectionConfigResp)
	r.InputConfig = m.InputConfig.CloneVT()
	r.CompiledConfig = m.CompiledConfig.CloneVT()
	r.Version = m.Version
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *GetProtectionConfigResp) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *GetProtectionAttachmentConfigReq) CloneVT() *GetProtectionAttachmentConfigReq {
	if m == nil {
		return (*GetProtectionAttachmentConfigReq)(nil)
	}
	r := new(GetProtectionAttachmentConfigReq)
	r.AttachmentId = m.AttachmentId
	r.Version = m.Version
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *GetProtectionAttachmentConfigReq) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *GetProtectionAttachmentConfigResp) CloneVT() *GetProtectionAttachmentConfigResp {
	if m == nil {
		return (*GetProtectionAttachmentConfigResp)(nil)
	}
	r := new(GetProtectionAttachmentConfigResp)
	r.Config = m.Config.CloneVT()
	r.Version = m.Version
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *GetProtectionAttachmentConfigResp) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *ConfigureProtectionReq) StableEqualVT(that *ConfigureProtectionReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.ProtectionConfig.StableEqualVT(that.ProtectionConfig) {
		return false
	}
	if this.Source != that.Source {
		return false
	}
	if !this.SourceMetadata.StableEqualVT(that.SourceMetadata) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ConfigureProtectionReq) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ConfigureProtectionReq)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *ConfigureProtectionReq) EqualVT(that *ConfigureProtectionReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.ProtectionConfig.EqualVT(that.ProtectionConfig) {
		return false
	}
	if this.Source != that.Source {
		return false
	}
	if !this.SourceMetadata.EqualVT(that.SourceMetadata) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ConfigureProtectionReq) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ConfigureProtectionReq)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ConfigureProtectionResp) StableEqualVT(that *ConfigureProtectionResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.ProtectionId != that.ProtectionId {
		return false
	}
	if this.Version != that.Version {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ConfigureProtectionResp) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ConfigureProtectionResp)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *ConfigureProtectionResp) EqualVT(that *ConfigureProtectionResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.ProtectionId != that.ProtectionId {
		return false
	}
	if this.Version != that.Version {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ConfigureProtectionResp) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ConfigureProtectionResp)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ValidateConfigureProtectionResp) StableEqualVT(that *ValidateConfigureProtectionResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.InputConfig.StableEqualVT(that.InputConfig) {
		return false
	}
	if !this.CompiledConfig.StableEqualVT(that.CompiledConfig) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ValidateConfigureProtectionResp) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ValidateConfigureProtectionResp)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *ValidateConfigureProtectionResp) EqualVT(that *ValidateConfigureProtectionResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.InputConfig.EqualVT(that.InputConfig) {
		return false
	}
	if !this.CompiledConfig.EqualVT(that.CompiledConfig) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ValidateConfigureProtectionResp) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ValidateConfigureProtectionResp)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ListProtectionsReq) StableEqualVT(that *ListProtectionsReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.PageToken != that.PageToken {
		return false
	}
	if this.PageSize != that.PageSize {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ListProtectionsReq) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ListProtectionsReq)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *ListProtectionsReq) EqualVT(that *ListProtectionsReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.PageToken != that.PageToken {
		return false
	}
	if this.PageSize != that.PageSize {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ListProtectionsReq) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ListProtectionsReq)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *ListProtectionsResp) StableEqualVT(that *ListProtectionsResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.Protections) != len(that.Protections) {
		return false
	}
	for i, vx := range this.Protections {
		vy := that.Protections[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Protection{}
			}
			if q == nil {
				q = &Protection{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	if this.NextPageToken != that.NextPageToken {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ListProtectionsResp) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ListProtectionsResp)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *ListProtectionsResp) EqualVT(that *ListProtectionsResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if len(this.Protections) != len(that.Protections) {
		return false
	}
	for i, vx := range this.Protections {
		vy := that.Protections[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Protection{}
			}
			if q == nil {
				q = &Protection{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if this.NextPageToken != that.NextPageToken {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ListProtectionsResp) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*ListProtectionsResp)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *GetProtectionReq) StableEqualVT(that *GetProtectionReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Protection != that.Protection {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetProtectionReq) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetProtectionReq)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *GetProtectionReq) EqualVT(that *GetProtectionReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Protection != that.Protection {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetProtectionReq) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetProtectionReq)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *GetProtectionResp) StableEqualVT(that *GetProtectionResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Protection.StableEqualVT(that.Protection) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetProtectionResp) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetProtectionResp)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *GetProtectionResp) EqualVT(that *GetProtectionResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Protection.EqualVT(that.Protection) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetProtectionResp) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetProtectionResp)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *GetProtectionConfigReq) StableEqualVT(that *GetProtectionConfigReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Protection != that.Protection {
		return false
	}
	if this.Version != that.Version {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetProtectionConfigReq) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetProtectionConfigReq)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *GetProtectionConfigReq) EqualVT(that *GetProtectionConfigReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Protection != that.Protection {
		return false
	}
	if this.Version != that.Version {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetProtectionConfigReq) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetProtectionConfigReq)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *GetProtectionConfigResp) StableEqualVT(that *GetProtectionConfigResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.CompiledConfig.StableEqualVT(that.CompiledConfig) {
		return false
	}
	if this.Version != that.Version {
		return false
	}
	if !this.InputConfig.StableEqualVT(that.InputConfig) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetProtectionConfigResp) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetProtectionConfigResp)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *GetProtectionConfigResp) EqualVT(that *GetProtectionConfigResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.CompiledConfig.EqualVT(that.CompiledConfig) {
		return false
	}
	if this.Version != that.Version {
		return false
	}
	if !this.InputConfig.EqualVT(that.InputConfig) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetProtectionConfigResp) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetProtectionConfigResp)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *GetProtectionAttachmentConfigReq) StableEqualVT(that *GetProtectionAttachmentConfigReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.AttachmentId != that.AttachmentId {
		return false
	}
	if this.Version != that.Version {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetProtectionAttachmentConfigReq) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetProtectionAttachmentConfigReq)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *GetProtectionAttachmentConfigReq) EqualVT(that *GetProtectionAttachmentConfigReq) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.AttachmentId != that.AttachmentId {
		return false
	}
	if this.Version != that.Version {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetProtectionAttachmentConfigReq) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetProtectionAttachmentConfigReq)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *GetProtectionAttachmentConfigResp) StableEqualVT(that *GetProtectionAttachmentConfigResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Config.StableEqualVT(that.Config) {
		return false
	}
	if this.Version != that.Version {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetProtectionAttachmentConfigResp) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetProtectionAttachmentConfigResp)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *GetProtectionAttachmentConfigResp) EqualVT(that *GetProtectionAttachmentConfigResp) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Config.EqualVT(that.Config) {
		return false
	}
	if this.Version != that.Version {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *GetProtectionAttachmentConfigResp) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*GetProtectionAttachmentConfigResp)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
