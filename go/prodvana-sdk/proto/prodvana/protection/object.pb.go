// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.10
// source: prodvana/protection/object.proto

package protection

import (
	object "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/object"
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

type Protection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta   *object.ObjectMeta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Config *ProtectionConfig  `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"` // TODO: Add information about where this protection is currently used?
}

func (x *Protection) Reset() {
	*x = Protection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_protection_object_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Protection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Protection) ProtoMessage() {}

func (x *Protection) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_protection_object_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Protection.ProtoReflect.Descriptor instead.
func (*Protection) Descriptor() ([]byte, []int) {
	return file_prodvana_protection_object_proto_rawDescGZIP(), []int{0}
}

func (x *Protection) GetMeta() *object.ObjectMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Protection) GetConfig() *ProtectionConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

var File_prodvana_protection_object_proto protoreflect.FileDescriptor

var file_prodvana_protection_object_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e,
	0x61, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x7c, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f,
	0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12,
	0x3d, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x4f,
	0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2d, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e,
	0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64,
	0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prodvana_protection_object_proto_rawDescOnce sync.Once
	file_prodvana_protection_object_proto_rawDescData = file_prodvana_protection_object_proto_rawDesc
)

func file_prodvana_protection_object_proto_rawDescGZIP() []byte {
	file_prodvana_protection_object_proto_rawDescOnce.Do(func() {
		file_prodvana_protection_object_proto_rawDescData = protoimpl.X.CompressGZIP(file_prodvana_protection_object_proto_rawDescData)
	})
	return file_prodvana_protection_object_proto_rawDescData
}

var file_prodvana_protection_object_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_prodvana_protection_object_proto_goTypes = []interface{}{
	(*Protection)(nil),        // 0: prodvana.protection.Protection
	(*object.ObjectMeta)(nil), // 1: prodvana.object.ObjectMeta
	(*ProtectionConfig)(nil),  // 2: prodvana.protection.ProtectionConfig
}
var file_prodvana_protection_object_proto_depIdxs = []int32{
	1, // 0: prodvana.protection.Protection.meta:type_name -> prodvana.object.ObjectMeta
	2, // 1: prodvana.protection.Protection.config:type_name -> prodvana.protection.ProtectionConfig
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_prodvana_protection_object_proto_init() }
func file_prodvana_protection_object_proto_init() {
	if File_prodvana_protection_object_proto != nil {
		return
	}
	file_prodvana_protection_protection_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_prodvana_protection_object_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Protection); i {
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
			RawDescriptor: file_prodvana_protection_object_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prodvana_protection_object_proto_goTypes,
		DependencyIndexes: file_prodvana_protection_object_proto_depIdxs,
		MessageInfos:      file_prodvana_protection_object_proto_msgTypes,
	}.Build()
	File_prodvana_protection_object_proto = out.File
	file_prodvana_protection_object_proto_rawDesc = nil
	file_prodvana_protection_object_proto_goTypes = nil
	file_prodvana_protection_object_proto_depIdxs = nil
}
