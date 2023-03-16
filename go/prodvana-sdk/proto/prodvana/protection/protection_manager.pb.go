// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.10
// source: prodvana/protection/protection_manager.proto

package protection

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	version "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/version"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ConfigureProtectionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProtectionConfig *ProtectionConfig       `protobuf:"bytes,1,opt,name=protection_config,json=protectionConfig,proto3" json:"protection_config,omitempty"`
	Source           version.Source          `protobuf:"varint,2,opt,name=source,proto3,enum=prodvana.version.Source" json:"source,omitempty"`
	SourceMetadata   *version.SourceMetadata `protobuf:"bytes,3,opt,name=source_metadata,json=sourceMetadata,proto3" json:"source_metadata,omitempty"`
}

func (x *ConfigureProtectionReq) Reset() {
	*x = ConfigureProtectionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_protection_protection_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigureProtectionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigureProtectionReq) ProtoMessage() {}

func (x *ConfigureProtectionReq) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_protection_protection_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigureProtectionReq.ProtoReflect.Descriptor instead.
func (*ConfigureProtectionReq) Descriptor() ([]byte, []int) {
	return file_prodvana_protection_protection_manager_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigureProtectionReq) GetProtectionConfig() *ProtectionConfig {
	if x != nil {
		return x.ProtectionConfig
	}
	return nil
}

func (x *ConfigureProtectionReq) GetSource() version.Source {
	if x != nil {
		return x.Source
	}
	return version.Source(0)
}

func (x *ConfigureProtectionReq) GetSourceMetadata() *version.SourceMetadata {
	if x != nil {
		return x.SourceMetadata
	}
	return nil
}

type ConfigureProtectionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProtectionId string `protobuf:"bytes,1,opt,name=protection_id,json=protectionId,proto3" json:"protection_id,omitempty"`
	Version      string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ConfigureProtectionResp) Reset() {
	*x = ConfigureProtectionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_protection_protection_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigureProtectionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigureProtectionResp) ProtoMessage() {}

func (x *ConfigureProtectionResp) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_protection_protection_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigureProtectionResp.ProtoReflect.Descriptor instead.
func (*ConfigureProtectionResp) Descriptor() ([]byte, []int) {
	return file_prodvana_protection_protection_manager_proto_rawDescGZIP(), []int{1}
}

func (x *ConfigureProtectionResp) GetProtectionId() string {
	if x != nil {
		return x.ProtectionId
	}
	return ""
}

func (x *ConfigureProtectionResp) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type ValidateConfigureProtectionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ValidateConfigureProtectionResp) Reset() {
	*x = ValidateConfigureProtectionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_protection_protection_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateConfigureProtectionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateConfigureProtectionResp) ProtoMessage() {}

func (x *ValidateConfigureProtectionResp) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_protection_protection_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateConfigureProtectionResp.ProtoReflect.Descriptor instead.
func (*ValidateConfigureProtectionResp) Descriptor() ([]byte, []int) {
	return file_prodvana_protection_protection_manager_proto_rawDescGZIP(), []int{2}
}

type ListProtectionsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageToken string `protobuf:"bytes,1,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	PageSize  int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListProtectionsReq) Reset() {
	*x = ListProtectionsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_protection_protection_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProtectionsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProtectionsReq) ProtoMessage() {}

func (x *ListProtectionsReq) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_protection_protection_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProtectionsReq.ProtoReflect.Descriptor instead.
func (*ListProtectionsReq) Descriptor() ([]byte, []int) {
	return file_prodvana_protection_protection_manager_proto_rawDescGZIP(), []int{3}
}

func (x *ListProtectionsReq) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListProtectionsReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListProtectionsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protections   []*Protection `protobuf:"bytes,1,rep,name=protections,proto3" json:"protections,omitempty"`
	NextPageToken string        `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListProtectionsResp) Reset() {
	*x = ListProtectionsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_protection_protection_manager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProtectionsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProtectionsResp) ProtoMessage() {}

func (x *ListProtectionsResp) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_protection_protection_manager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProtectionsResp.ProtoReflect.Descriptor instead.
func (*ListProtectionsResp) Descriptor() ([]byte, []int) {
	return file_prodvana_protection_protection_manager_proto_rawDescGZIP(), []int{4}
}

func (x *ListProtectionsResp) GetProtections() []*Protection {
	if x != nil {
		return x.Protections
	}
	return nil
}

func (x *ListProtectionsResp) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type GetProtectionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protection string `protobuf:"bytes,1,opt,name=protection,proto3" json:"protection,omitempty"`
}

func (x *GetProtectionReq) Reset() {
	*x = GetProtectionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_protection_protection_manager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProtectionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProtectionReq) ProtoMessage() {}

func (x *GetProtectionReq) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_protection_protection_manager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProtectionReq.ProtoReflect.Descriptor instead.
func (*GetProtectionReq) Descriptor() ([]byte, []int) {
	return file_prodvana_protection_protection_manager_proto_rawDescGZIP(), []int{5}
}

func (x *GetProtectionReq) GetProtection() string {
	if x != nil {
		return x.Protection
	}
	return ""
}

type GetProtectionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protection *Protection `protobuf:"bytes,1,opt,name=protection,proto3" json:"protection,omitempty"`
}

func (x *GetProtectionResp) Reset() {
	*x = GetProtectionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_protection_protection_manager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProtectionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProtectionResp) ProtoMessage() {}

func (x *GetProtectionResp) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_protection_protection_manager_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProtectionResp.ProtoReflect.Descriptor instead.
func (*GetProtectionResp) Descriptor() ([]byte, []int) {
	return file_prodvana_protection_protection_manager_proto_rawDescGZIP(), []int{6}
}

func (x *GetProtectionResp) GetProtection() *Protection {
	if x != nil {
		return x.Protection
	}
	return nil
}

type GetProtectionConfigReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protection string `protobuf:"bytes,1,opt,name=protection,proto3" json:"protection,omitempty"`
	Version    string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"` // omit to get latest version
}

func (x *GetProtectionConfigReq) Reset() {
	*x = GetProtectionConfigReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_protection_protection_manager_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProtectionConfigReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProtectionConfigReq) ProtoMessage() {}

func (x *GetProtectionConfigReq) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_protection_protection_manager_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProtectionConfigReq.ProtoReflect.Descriptor instead.
func (*GetProtectionConfigReq) Descriptor() ([]byte, []int) {
	return file_prodvana_protection_protection_manager_proto_rawDescGZIP(), []int{7}
}

func (x *GetProtectionConfigReq) GetProtection() string {
	if x != nil {
		return x.Protection
	}
	return ""
}

func (x *GetProtectionConfigReq) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type GetProtectionConfigResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config  *ProtectionConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	Version string            `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *GetProtectionConfigResp) Reset() {
	*x = GetProtectionConfigResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_protection_protection_manager_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProtectionConfigResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProtectionConfigResp) ProtoMessage() {}

func (x *GetProtectionConfigResp) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_protection_protection_manager_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProtectionConfigResp.ProtoReflect.Descriptor instead.
func (*GetProtectionConfigResp) Descriptor() ([]byte, []int) {
	return file_prodvana_protection_protection_manager_proto_rawDescGZIP(), []int{8}
}

func (x *GetProtectionConfigResp) GetConfig() *ProtectionConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *GetProtectionConfigResp) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_prodvana_protection_protection_manager_proto protoreflect.FileDescriptor

var file_prodvana_protection_protection_manager_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x26, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf3, 0x01, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x5c, 0x0a, 0x11,
	0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61,
	0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0f,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61,
	0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x58, 0x0a, 0x17, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0x21, 0x0a, 0x1f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x50, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x80, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x41,
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3b, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x27, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x74,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x54, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3f, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5b, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x12, 0x27, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x10, 0x01, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x72, 0x0a, 0x17, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x3d, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0x88, 0x06,
	0x0a, 0x11, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x12, 0x96, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x2c, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76,
	0x61, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01,
	0x2a, 0x22, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x12, 0xaf, 0x01, 0x0a,
	0x1b, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x76, 0x61, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x3a, 0x01, 0x2a, 0x22, 0x22, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x7d,
	0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x28, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x86, 0x01,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x25, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x26,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x3d, 0x2a, 0x7d, 0x12, 0x9f, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2b,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x2c, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x27, 0x12, 0x25, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3d, 0x2a,
	0x7d, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x4f, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f,
	0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f,
	0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_prodvana_protection_protection_manager_proto_rawDescOnce sync.Once
	file_prodvana_protection_protection_manager_proto_rawDescData = file_prodvana_protection_protection_manager_proto_rawDesc
)

func file_prodvana_protection_protection_manager_proto_rawDescGZIP() []byte {
	file_prodvana_protection_protection_manager_proto_rawDescOnce.Do(func() {
		file_prodvana_protection_protection_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_prodvana_protection_protection_manager_proto_rawDescData)
	})
	return file_prodvana_protection_protection_manager_proto_rawDescData
}

var file_prodvana_protection_protection_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_prodvana_protection_protection_manager_proto_goTypes = []interface{}{
	(*ConfigureProtectionReq)(nil),          // 0: prodvana.protection.ConfigureProtectionReq
	(*ConfigureProtectionResp)(nil),         // 1: prodvana.protection.ConfigureProtectionResp
	(*ValidateConfigureProtectionResp)(nil), // 2: prodvana.protection.ValidateConfigureProtectionResp
	(*ListProtectionsReq)(nil),              // 3: prodvana.protection.ListProtectionsReq
	(*ListProtectionsResp)(nil),             // 4: prodvana.protection.ListProtectionsResp
	(*GetProtectionReq)(nil),                // 5: prodvana.protection.GetProtectionReq
	(*GetProtectionResp)(nil),               // 6: prodvana.protection.GetProtectionResp
	(*GetProtectionConfigReq)(nil),          // 7: prodvana.protection.GetProtectionConfigReq
	(*GetProtectionConfigResp)(nil),         // 8: prodvana.protection.GetProtectionConfigResp
	(*ProtectionConfig)(nil),                // 9: prodvana.protection.ProtectionConfig
	(version.Source)(0),                     // 10: prodvana.version.Source
	(*version.SourceMetadata)(nil),          // 11: prodvana.version.SourceMetadata
	(*Protection)(nil),                      // 12: prodvana.protection.Protection
}
var file_prodvana_protection_protection_manager_proto_depIdxs = []int32{
	9,  // 0: prodvana.protection.ConfigureProtectionReq.protection_config:type_name -> prodvana.protection.ProtectionConfig
	10, // 1: prodvana.protection.ConfigureProtectionReq.source:type_name -> prodvana.version.Source
	11, // 2: prodvana.protection.ConfigureProtectionReq.source_metadata:type_name -> prodvana.version.SourceMetadata
	12, // 3: prodvana.protection.ListProtectionsResp.protections:type_name -> prodvana.protection.Protection
	12, // 4: prodvana.protection.GetProtectionResp.protection:type_name -> prodvana.protection.Protection
	9,  // 5: prodvana.protection.GetProtectionConfigResp.config:type_name -> prodvana.protection.ProtectionConfig
	0,  // 6: prodvana.protection.ProtectionManager.ConfigureProtection:input_type -> prodvana.protection.ConfigureProtectionReq
	0,  // 7: prodvana.protection.ProtectionManager.ValidateConfigureProtection:input_type -> prodvana.protection.ConfigureProtectionReq
	3,  // 8: prodvana.protection.ProtectionManager.ListProtections:input_type -> prodvana.protection.ListProtectionsReq
	5,  // 9: prodvana.protection.ProtectionManager.GetProtection:input_type -> prodvana.protection.GetProtectionReq
	7,  // 10: prodvana.protection.ProtectionManager.GetProtectionConfig:input_type -> prodvana.protection.GetProtectionConfigReq
	1,  // 11: prodvana.protection.ProtectionManager.ConfigureProtection:output_type -> prodvana.protection.ConfigureProtectionResp
	2,  // 12: prodvana.protection.ProtectionManager.ValidateConfigureProtection:output_type -> prodvana.protection.ValidateConfigureProtectionResp
	4,  // 13: prodvana.protection.ProtectionManager.ListProtections:output_type -> prodvana.protection.ListProtectionsResp
	6,  // 14: prodvana.protection.ProtectionManager.GetProtection:output_type -> prodvana.protection.GetProtectionResp
	8,  // 15: prodvana.protection.ProtectionManager.GetProtectionConfig:output_type -> prodvana.protection.GetProtectionConfigResp
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_prodvana_protection_protection_manager_proto_init() }
func file_prodvana_protection_protection_manager_proto_init() {
	if File_prodvana_protection_protection_manager_proto != nil {
		return
	}
	file_prodvana_protection_object_proto_init()
	file_prodvana_protection_protection_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_prodvana_protection_protection_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigureProtectionReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_prodvana_protection_protection_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigureProtectionResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_prodvana_protection_protection_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateConfigureProtectionResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_prodvana_protection_protection_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProtectionsReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_prodvana_protection_protection_manager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProtectionsResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_prodvana_protection_protection_manager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProtectionReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_prodvana_protection_protection_manager_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProtectionResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_prodvana_protection_protection_manager_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProtectionConfigReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_prodvana_protection_protection_manager_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProtectionConfigResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_prodvana_protection_protection_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_prodvana_protection_protection_manager_proto_goTypes,
		DependencyIndexes: file_prodvana_protection_protection_manager_proto_depIdxs,
		MessageInfos:      file_prodvana_protection_protection_manager_proto_msgTypes,
	}.Build()
	File_prodvana_protection_protection_manager_proto = out.File
	file_prodvana_protection_protection_manager_proto_rawDesc = nil
	file_prodvana_protection_protection_manager_proto_goTypes = nil
	file_prodvana_protection_protection_manager_proto_depIdxs = nil
}
