// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.10
// source: prodvana/volumes/volumes.proto

package volumes

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type VolumeSource_EphemeralSource_Medium int32

const (
	VolumeSource_EphemeralSource_UNKNOWN VolumeSource_EphemeralSource_Medium = 0
	VolumeSource_EphemeralSource_DEFAULT VolumeSource_EphemeralSource_Medium = 1 // for kubernetes, this the the node's default storage medium
	VolumeSource_EphemeralSource_MEMORY  VolumeSource_EphemeralSource_Medium = 2
)

// Enum value maps for VolumeSource_EphemeralSource_Medium.
var (
	VolumeSource_EphemeralSource_Medium_name = map[int32]string{
		0: "UNKNOWN",
		1: "DEFAULT",
		2: "MEMORY",
	}
	VolumeSource_EphemeralSource_Medium_value = map[string]int32{
		"UNKNOWN": 0,
		"DEFAULT": 1,
		"MEMORY":  2,
	}
)

func (x VolumeSource_EphemeralSource_Medium) Enum() *VolumeSource_EphemeralSource_Medium {
	p := new(VolumeSource_EphemeralSource_Medium)
	*p = x
	return p
}

func (x VolumeSource_EphemeralSource_Medium) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VolumeSource_EphemeralSource_Medium) Descriptor() protoreflect.EnumDescriptor {
	return file_prodvana_volumes_volumes_proto_enumTypes[0].Descriptor()
}

func (VolumeSource_EphemeralSource_Medium) Type() protoreflect.EnumType {
	return &file_prodvana_volumes_volumes_proto_enumTypes[0]
}

func (x VolumeSource_EphemeralSource_Medium) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VolumeSource_EphemeralSource_Medium.Descriptor instead.
func (VolumeSource_EphemeralSource_Medium) EnumDescriptor() ([]byte, []int) {
	return file_prodvana_volumes_volumes_proto_rawDescGZIP(), []int{0, 1, 0}
}

type VolumeSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Source:
	//
	//	*VolumeSource_Secret
	//	*VolumeSource_Ephemeral
	Source isVolumeSource_Source `protobuf_oneof:"source"`
}

func (x *VolumeSource) Reset() {
	*x = VolumeSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_volumes_volumes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeSource) ProtoMessage() {}

func (x *VolumeSource) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_volumes_volumes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeSource.ProtoReflect.Descriptor instead.
func (*VolumeSource) Descriptor() ([]byte, []int) {
	return file_prodvana_volumes_volumes_proto_rawDescGZIP(), []int{0}
}

func (m *VolumeSource) GetSource() isVolumeSource_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *VolumeSource) GetSecret() *VolumeSource_SecretSource {
	if x, ok := x.GetSource().(*VolumeSource_Secret); ok {
		return x.Secret
	}
	return nil
}

func (x *VolumeSource) GetEphemeral() *VolumeSource_EphemeralSource {
	if x, ok := x.GetSource().(*VolumeSource_Ephemeral); ok {
		return x.Ephemeral
	}
	return nil
}

type isVolumeSource_Source interface {
	isVolumeSource_Source()
}

type VolumeSource_Secret struct {
	Secret *VolumeSource_SecretSource `protobuf:"bytes,1,opt,name=secret,proto3,oneof"`
}

type VolumeSource_Ephemeral struct {
	Ephemeral *VolumeSource_EphemeralSource `protobuf:"bytes,2,opt,name=ephemeral,proto3,oneof"`
}

func (*VolumeSource_Secret) isVolumeSource_Source() {}

func (*VolumeSource_Ephemeral) isVolumeSource_Source() {}

type VolumeMount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MountPath string `protobuf:"bytes,1,opt,name=mount_path,json=mountPath,proto3" json:"mount_path,omitempty"`
	ReadOnly  bool   `protobuf:"varint,2,opt,name=read_only,json=readOnly,proto3" json:"read_only,omitempty"`
}

func (x *VolumeMount) Reset() {
	*x = VolumeMount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_volumes_volumes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeMount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeMount) ProtoMessage() {}

func (x *VolumeMount) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_volumes_volumes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeMount.ProtoReflect.Descriptor instead.
func (*VolumeMount) Descriptor() ([]byte, []int) {
	return file_prodvana_volumes_volumes_proto_rawDescGZIP(), []int{1}
}

func (x *VolumeMount) GetMountPath() string {
	if x != nil {
		return x.MountPath
	}
	return ""
}

func (x *VolumeMount) GetReadOnly() bool {
	if x != nil {
		return x.ReadOnly
	}
	return false
}

type Volume struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Source *VolumeSource `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Mount  *VolumeMount  `protobuf:"bytes,3,opt,name=mount,proto3" json:"mount,omitempty"`
}

func (x *Volume) Reset() {
	*x = Volume{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_volumes_volumes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Volume) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Volume) ProtoMessage() {}

func (x *Volume) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_volumes_volumes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Volume.ProtoReflect.Descriptor instead.
func (*Volume) Descriptor() ([]byte, []int) {
	return file_prodvana_volumes_volumes_proto_rawDescGZIP(), []int{2}
}

func (x *Volume) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Volume) GetSource() *VolumeSource {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *Volume) GetMount() *VolumeMount {
	if x != nil {
		return x.Mount
	}
	return nil
}

type VolumeSource_SecretSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecretName string `protobuf:"bytes,1,opt,name=secret_name,json=secretName,proto3" json:"secret_name,omitempty"`
}

func (x *VolumeSource_SecretSource) Reset() {
	*x = VolumeSource_SecretSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_volumes_volumes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeSource_SecretSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeSource_SecretSource) ProtoMessage() {}

func (x *VolumeSource_SecretSource) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_volumes_volumes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeSource_SecretSource.ProtoReflect.Descriptor instead.
func (*VolumeSource_SecretSource) Descriptor() ([]byte, []int) {
	return file_prodvana_volumes_volumes_proto_rawDescGZIP(), []int{0, 0}
}

func (x *VolumeSource_SecretSource) GetSecretName() string {
	if x != nil {
		return x.SecretName
	}
	return ""
}

type VolumeSource_EphemeralSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Medium VolumeSource_EphemeralSource_Medium `protobuf:"varint,1,opt,name=medium,proto3,enum=prodvana.volumes.VolumeSource_EphemeralSource_Medium" json:"medium,omitempty"`
}

func (x *VolumeSource_EphemeralSource) Reset() {
	*x = VolumeSource_EphemeralSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_volumes_volumes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeSource_EphemeralSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeSource_EphemeralSource) ProtoMessage() {}

func (x *VolumeSource_EphemeralSource) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_volumes_volumes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeSource_EphemeralSource.ProtoReflect.Descriptor instead.
func (*VolumeSource_EphemeralSource) Descriptor() ([]byte, []int) {
	return file_prodvana_volumes_volumes_proto_rawDescGZIP(), []int{0, 1}
}

func (x *VolumeSource_EphemeralSource) GetMedium() VolumeSource_EphemeralSource_Medium {
	if x != nil {
		return x.Medium
	}
	return VolumeSource_EphemeralSource_UNKNOWN
}

var File_prodvana_volumes_volumes_proto protoreflect.FileDescriptor

var file_prodvana_volumes_volumes_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x73, 0x2f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x73, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x02, 0x0a, 0x0c,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x06,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x2e,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x06, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x12, 0x4e, 0x0a, 0x09, 0x65, 0x70, 0x68, 0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e,
	0x61, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x45, 0x70, 0x68, 0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x09, 0x65, 0x70, 0x68, 0x65, 0x6d, 0x65,
	0x72, 0x61, 0x6c, 0x1a, 0x38, 0x0a, 0x0c, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x90, 0x01,
	0x0a, 0x0f, 0x45, 0x70, 0x68, 0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x4d, 0x0a, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x35, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x73, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x45, 0x70, 0x68, 0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x52, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d,
	0x22, 0x2e, 0x0a, 0x06, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x46, 0x41, 0x55,
	0x4c, 0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x45, 0x4d, 0x4f, 0x52, 0x59, 0x10, 0x02,
	0x42, 0x08, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x49, 0x0a, 0x0b, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64,
	0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x61,
	0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0x92, 0x01, 0x0a, 0x06, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73,
	0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x4d, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x05, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e,
	0x61, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2d, 0x73, 0x64,
	0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61,
	0x2f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prodvana_volumes_volumes_proto_rawDescOnce sync.Once
	file_prodvana_volumes_volumes_proto_rawDescData = file_prodvana_volumes_volumes_proto_rawDesc
)

func file_prodvana_volumes_volumes_proto_rawDescGZIP() []byte {
	file_prodvana_volumes_volumes_proto_rawDescOnce.Do(func() {
		file_prodvana_volumes_volumes_proto_rawDescData = protoimpl.X.CompressGZIP(file_prodvana_volumes_volumes_proto_rawDescData)
	})
	return file_prodvana_volumes_volumes_proto_rawDescData
}

var file_prodvana_volumes_volumes_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_prodvana_volumes_volumes_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_prodvana_volumes_volumes_proto_goTypes = []interface{}{
	(VolumeSource_EphemeralSource_Medium)(0), // 0: prodvana.volumes.VolumeSource.EphemeralSource.Medium
	(*VolumeSource)(nil),                     // 1: prodvana.volumes.VolumeSource
	(*VolumeMount)(nil),                      // 2: prodvana.volumes.VolumeMount
	(*Volume)(nil),                           // 3: prodvana.volumes.Volume
	(*VolumeSource_SecretSource)(nil),        // 4: prodvana.volumes.VolumeSource.SecretSource
	(*VolumeSource_EphemeralSource)(nil),     // 5: prodvana.volumes.VolumeSource.EphemeralSource
}
var file_prodvana_volumes_volumes_proto_depIdxs = []int32{
	4, // 0: prodvana.volumes.VolumeSource.secret:type_name -> prodvana.volumes.VolumeSource.SecretSource
	5, // 1: prodvana.volumes.VolumeSource.ephemeral:type_name -> prodvana.volumes.VolumeSource.EphemeralSource
	1, // 2: prodvana.volumes.Volume.source:type_name -> prodvana.volumes.VolumeSource
	2, // 3: prodvana.volumes.Volume.mount:type_name -> prodvana.volumes.VolumeMount
	0, // 4: prodvana.volumes.VolumeSource.EphemeralSource.medium:type_name -> prodvana.volumes.VolumeSource.EphemeralSource.Medium
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_prodvana_volumes_volumes_proto_init() }
func file_prodvana_volumes_volumes_proto_init() {
	if File_prodvana_volumes_volumes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prodvana_volumes_volumes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeSource); i {
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
		file_prodvana_volumes_volumes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeMount); i {
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
		file_prodvana_volumes_volumes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Volume); i {
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
		file_prodvana_volumes_volumes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeSource_SecretSource); i {
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
		file_prodvana_volumes_volumes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeSource_EphemeralSource); i {
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
	file_prodvana_volumes_volumes_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*VolumeSource_Secret)(nil),
		(*VolumeSource_Ephemeral)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_prodvana_volumes_volumes_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prodvana_volumes_volumes_proto_goTypes,
		DependencyIndexes: file_prodvana_volumes_volumes_proto_depIdxs,
		EnumInfos:         file_prodvana_volumes_volumes_proto_enumTypes,
		MessageInfos:      file_prodvana_volumes_volumes_proto_msgTypes,
	}.Build()
	File_prodvana_volumes_volumes_proto = out.File
	file_prodvana_volumes_volumes_proto_rawDesc = nil
	file_prodvana_volumes_volumes_proto_goTypes = nil
	file_prodvana_volumes_volumes_proto_depIdxs = nil
}
