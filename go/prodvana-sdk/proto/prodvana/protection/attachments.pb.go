// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.10
// source: prodvana/protection/attachments.proto

package protection

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

type AttachmentType int32

const (
	AttachmentType_UNKNOWN          AttachmentType = 0
	AttachmentType_RELEASE_CHANNEL  AttachmentType = 1
	AttachmentType_SERVICE_INSTANCE AttachmentType = 2
	AttachmentType_CONVERGENCE      AttachmentType = 3
)

// Enum value maps for AttachmentType.
var (
	AttachmentType_name = map[int32]string{
		0: "UNKNOWN",
		1: "RELEASE_CHANNEL",
		2: "SERVICE_INSTANCE",
		3: "CONVERGENCE",
	}
	AttachmentType_value = map[string]int32{
		"UNKNOWN":          0,
		"RELEASE_CHANNEL":  1,
		"SERVICE_INSTANCE": 2,
		"CONVERGENCE":      3,
	}
)

func (x AttachmentType) Enum() *AttachmentType {
	p := new(AttachmentType)
	*p = x
	return p
}

func (x AttachmentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AttachmentType) Descriptor() protoreflect.EnumDescriptor {
	return file_prodvana_protection_attachments_proto_enumTypes[0].Descriptor()
}

func (AttachmentType) Type() protoreflect.EnumType {
	return &file_prodvana_protection_attachments_proto_enumTypes[0]
}

func (x AttachmentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AttachmentType.Descriptor instead.
func (AttachmentType) EnumDescriptor() ([]byte, []int) {
	return file_prodvana_protection_attachments_proto_rawDescGZIP(), []int{0}
}

type ProtectionAttachmentConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // optional, default to protection name
	Ref  *ProtectionReference `protobuf:"bytes,2,opt,name=ref,proto3" json:"ref,omitempty"`
	// if set, the attachment is automatically used in all deployments for the subject of this attachment.
	// e.g. for release channels, all service instances in that release channel will use the attachment.
	Lifecycle []*ProtectionLifecycle `protobuf:"bytes,3,rep,name=lifecycle,proto3" json:"lifecycle,omitempty"`
}

func (x *ProtectionAttachmentConfig) Reset() {
	*x = ProtectionAttachmentConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_protection_attachments_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtectionAttachmentConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtectionAttachmentConfig) ProtoMessage() {}

func (x *ProtectionAttachmentConfig) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_protection_attachments_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtectionAttachmentConfig.ProtoReflect.Descriptor instead.
func (*ProtectionAttachmentConfig) Descriptor() ([]byte, []int) {
	return file_prodvana_protection_attachments_proto_rawDescGZIP(), []int{0}
}

func (x *ProtectionAttachmentConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProtectionAttachmentConfig) GetRef() *ProtectionReference {
	if x != nil {
		return x.Ref
	}
	return nil
}

func (x *ProtectionAttachmentConfig) GetLifecycle() []*ProtectionLifecycle {
	if x != nil {
		return x.Lifecycle
	}
	return nil
}

var File_prodvana_protection_attachments_proto protoreflect.FileDescriptor

var file_prodvana_protection_attachments_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x01, 0x0a, 0x1a, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x40, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x2c, 0xfa, 0x42, 0x29, 0x72, 0x27, 0x10, 0x00, 0x18, 0x3f, 0x32, 0x21, 0x5e,
	0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x3f, 0x28, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2d, 0x5d,
	0x2a, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x7b, 0x30, 0x2c, 0x31, 0x7d, 0x24,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x03, 0x72, 0x65, 0x66, 0x12, 0x57, 0x0a, 0x09,
	0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x42, 0x0f, 0xfa, 0x42, 0x0c, 0x92, 0x01,
	0x09, 0x08, 0x01, 0x22, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x09, 0x6c, 0x69, 0x66, 0x65,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x2a, 0x59, 0x0a, 0x0e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45, 0x4c, 0x45, 0x41, 0x53, 0x45, 0x5f,
	0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x45, 0x52,
	0x56, 0x49, 0x43, 0x45, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x02, 0x12,
	0x0f, 0x0a, 0x0b, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x47, 0x45, 0x4e, 0x43, 0x45, 0x10, 0x03,
	0x42, 0x4f, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61,
	0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76,
	0x61, 0x6e, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prodvana_protection_attachments_proto_rawDescOnce sync.Once
	file_prodvana_protection_attachments_proto_rawDescData = file_prodvana_protection_attachments_proto_rawDesc
)

func file_prodvana_protection_attachments_proto_rawDescGZIP() []byte {
	file_prodvana_protection_attachments_proto_rawDescOnce.Do(func() {
		file_prodvana_protection_attachments_proto_rawDescData = protoimpl.X.CompressGZIP(file_prodvana_protection_attachments_proto_rawDescData)
	})
	return file_prodvana_protection_attachments_proto_rawDescData
}

var file_prodvana_protection_attachments_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_prodvana_protection_attachments_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_prodvana_protection_attachments_proto_goTypes = []interface{}{
	(AttachmentType)(0),                // 0: prodvana.protection.AttachmentType
	(*ProtectionAttachmentConfig)(nil), // 1: prodvana.protection.ProtectionAttachmentConfig
	(*ProtectionReference)(nil),        // 2: prodvana.protection.ProtectionReference
	(*ProtectionLifecycle)(nil),        // 3: prodvana.protection.ProtectionLifecycle
}
var file_prodvana_protection_attachments_proto_depIdxs = []int32{
	2, // 0: prodvana.protection.ProtectionAttachmentConfig.ref:type_name -> prodvana.protection.ProtectionReference
	3, // 1: prodvana.protection.ProtectionAttachmentConfig.lifecycle:type_name -> prodvana.protection.ProtectionLifecycle
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_prodvana_protection_attachments_proto_init() }
func file_prodvana_protection_attachments_proto_init() {
	if File_prodvana_protection_attachments_proto != nil {
		return
	}
	file_prodvana_protection_protection_reference_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_prodvana_protection_attachments_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtectionAttachmentConfig); i {
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
			RawDescriptor: file_prodvana_protection_attachments_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prodvana_protection_attachments_proto_goTypes,
		DependencyIndexes: file_prodvana_protection_attachments_proto_depIdxs,
		EnumInfos:         file_prodvana_protection_attachments_proto_enumTypes,
		MessageInfos:      file_prodvana_protection_attachments_proto_msgTypes,
	}.Build()
	File_prodvana_protection_attachments_proto = out.File
	file_prodvana_protection_attachments_proto_rawDesc = nil
	file_prodvana_protection_attachments_proto_goTypes = nil
	file_prodvana_protection_attachments_proto_depIdxs = nil
}
