// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.10
// source: prodvana/common_config/external_config.proto

package common_config

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

type ExternalConfig_Type int32

const (
	ExternalConfig_UNKNOWN    ExternalConfig_Type = 0
	ExternalConfig_KUBERNETES ExternalConfig_Type = 1
	ExternalConfig_KUSTOMIZE  ExternalConfig_Type = 2
	ExternalConfig_HELM       ExternalConfig_Type = 3
)

// Enum value maps for ExternalConfig_Type.
var (
	ExternalConfig_Type_name = map[int32]string{
		0: "UNKNOWN",
		1: "KUBERNETES",
		2: "KUSTOMIZE",
		3: "HELM",
	}
	ExternalConfig_Type_value = map[string]int32{
		"UNKNOWN":    0,
		"KUBERNETES": 1,
		"KUSTOMIZE":  2,
		"HELM":       3,
	}
)

func (x ExternalConfig_Type) Enum() *ExternalConfig_Type {
	p := new(ExternalConfig_Type)
	*p = x
	return p
}

func (x ExternalConfig_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExternalConfig_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_prodvana_common_config_external_config_proto_enumTypes[0].Descriptor()
}

func (ExternalConfig_Type) Type() protoreflect.EnumType {
	return &file_prodvana_common_config_external_config_proto_enumTypes[0]
}

func (x ExternalConfig_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExternalConfig_Type.Descriptor instead.
func (ExternalConfig_Type) EnumDescriptor() ([]byte, []int) {
	return file_prodvana_common_config_external_config_proto_rawDescGZIP(), []int{1, 0}
}

type LocalConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *LocalConfig) Reset() {
	*x = LocalConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_common_config_external_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalConfig) ProtoMessage() {}

func (x *LocalConfig) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_common_config_external_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalConfig.ProtoReflect.Descriptor instead.
func (*LocalConfig) Descriptor() ([]byte, []int) {
	return file_prodvana_common_config_external_config_proto_rawDescGZIP(), []int{0}
}

func (x *LocalConfig) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type ExternalConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type ExternalConfig_Type `protobuf:"varint,1,opt,name=type,proto3,enum=prodvana.common_config.ExternalConfig_Type" json:"type,omitempty"`
	// Types that are assignable to SourceOneof:
	//
	//	*ExternalConfig_Inlined
	//	*ExternalConfig_Local
	SourceOneof isExternalConfig_SourceOneof `protobuf_oneof:"source_oneof"`
}

func (x *ExternalConfig) Reset() {
	*x = ExternalConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_common_config_external_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalConfig) ProtoMessage() {}

func (x *ExternalConfig) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_common_config_external_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalConfig.ProtoReflect.Descriptor instead.
func (*ExternalConfig) Descriptor() ([]byte, []int) {
	return file_prodvana_common_config_external_config_proto_rawDescGZIP(), []int{1}
}

func (x *ExternalConfig) GetType() ExternalConfig_Type {
	if x != nil {
		return x.Type
	}
	return ExternalConfig_UNKNOWN
}

func (m *ExternalConfig) GetSourceOneof() isExternalConfig_SourceOneof {
	if m != nil {
		return m.SourceOneof
	}
	return nil
}

func (x *ExternalConfig) GetInlined() string {
	if x, ok := x.GetSourceOneof().(*ExternalConfig_Inlined); ok {
		return x.Inlined
	}
	return ""
}

func (x *ExternalConfig) GetLocal() *LocalConfig {
	if x, ok := x.GetSourceOneof().(*ExternalConfig_Local); ok {
		return x.Local
	}
	return nil
}

type isExternalConfig_SourceOneof interface {
	isExternalConfig_SourceOneof()
}

type ExternalConfig_Inlined struct {
	Inlined string `protobuf:"bytes,2,opt,name=inlined,proto3,oneof"`
}

type ExternalConfig_Local struct {
	Local *LocalConfig `protobuf:"bytes,3,opt,name=local,proto3,oneof"`
}

func (*ExternalConfig_Inlined) isExternalConfig_SourceOneof() {}

func (*ExternalConfig_Local) isExternalConfig_SourceOneof() {}

var File_prodvana_common_config_external_config_proto protoreflect.FileDescriptor

var file_prodvana_common_config_external_config_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x2a, 0x0a, 0x0b, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x86, 0x02, 0x0a, 0x0e,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3f,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x23, 0x0a, 0x07, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x07, 0x69, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x64, 0x12, 0x3b, 0x0a, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x22, 0x3c, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4b, 0x55, 0x42, 0x45, 0x52, 0x4e,
	0x45, 0x54, 0x45, 0x53, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x4b, 0x55, 0x53, 0x54, 0x4f, 0x4d,
	0x49, 0x5a, 0x45, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x45, 0x4c, 0x4d, 0x10, 0x03, 0x42,
	0x13, 0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x12,
	0x03, 0xf8, 0x42, 0x01, 0x42, 0x52, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x64,
	0x76, 0x61, 0x6e, 0x61, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prodvana_common_config_external_config_proto_rawDescOnce sync.Once
	file_prodvana_common_config_external_config_proto_rawDescData = file_prodvana_common_config_external_config_proto_rawDesc
)

func file_prodvana_common_config_external_config_proto_rawDescGZIP() []byte {
	file_prodvana_common_config_external_config_proto_rawDescOnce.Do(func() {
		file_prodvana_common_config_external_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_prodvana_common_config_external_config_proto_rawDescData)
	})
	return file_prodvana_common_config_external_config_proto_rawDescData
}

var file_prodvana_common_config_external_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_prodvana_common_config_external_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_prodvana_common_config_external_config_proto_goTypes = []interface{}{
	(ExternalConfig_Type)(0), // 0: prodvana.common_config.ExternalConfig.Type
	(*LocalConfig)(nil),      // 1: prodvana.common_config.LocalConfig
	(*ExternalConfig)(nil),   // 2: prodvana.common_config.ExternalConfig
}
var file_prodvana_common_config_external_config_proto_depIdxs = []int32{
	0, // 0: prodvana.common_config.ExternalConfig.type:type_name -> prodvana.common_config.ExternalConfig.Type
	1, // 1: prodvana.common_config.ExternalConfig.local:type_name -> prodvana.common_config.LocalConfig
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_prodvana_common_config_external_config_proto_init() }
func file_prodvana_common_config_external_config_proto_init() {
	if File_prodvana_common_config_external_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prodvana_common_config_external_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalConfig); i {
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
		file_prodvana_common_config_external_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalConfig); i {
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
	file_prodvana_common_config_external_config_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ExternalConfig_Inlined)(nil),
		(*ExternalConfig_Local)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_prodvana_common_config_external_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prodvana_common_config_external_config_proto_goTypes,
		DependencyIndexes: file_prodvana_common_config_external_config_proto_depIdxs,
		EnumInfos:         file_prodvana_common_config_external_config_proto_enumTypes,
		MessageInfos:      file_prodvana_common_config_external_config_proto_msgTypes,
	}.Build()
	File_prodvana_common_config_external_config_proto = out.File
	file_prodvana_common_config_external_config_proto_rawDesc = nil
	file_prodvana_common_config_external_config_proto_goTypes = nil
	file_prodvana_common_config_external_config_proto_depIdxs = nil
}