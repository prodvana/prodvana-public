// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.10
// source: prodvana/common_config/kubernetes_config.proto

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

type KubernetesConfig_Type int32

const (
	KubernetesConfig_UNKNOWN    KubernetesConfig_Type = 0
	KubernetesConfig_KUBERNETES KubernetesConfig_Type = 1
	KubernetesConfig_KUSTOMIZE  KubernetesConfig_Type = 2
)

// Enum value maps for KubernetesConfig_Type.
var (
	KubernetesConfig_Type_name = map[int32]string{
		0: "UNKNOWN",
		1: "KUBERNETES",
		2: "KUSTOMIZE",
	}
	KubernetesConfig_Type_value = map[string]int32{
		"UNKNOWN":    0,
		"KUBERNETES": 1,
		"KUSTOMIZE":  2,
	}
)

func (x KubernetesConfig_Type) Enum() *KubernetesConfig_Type {
	p := new(KubernetesConfig_Type)
	*p = x
	return p
}

func (x KubernetesConfig_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KubernetesConfig_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_prodvana_common_config_kubernetes_config_proto_enumTypes[0].Descriptor()
}

func (KubernetesConfig_Type) Type() protoreflect.EnumType {
	return &file_prodvana_common_config_kubernetes_config_proto_enumTypes[0]
}

func (x KubernetesConfig_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KubernetesConfig_Type.Descriptor instead.
func (KubernetesConfig_Type) EnumDescriptor() ([]byte, []int) {
	return file_prodvana_common_config_kubernetes_config_proto_rawDescGZIP(), []int{1, 0}
}

type KubernetesConfig_EnvInjectionMode int32

const (
	KubernetesConfig_ENV_INJECT_UNKNOWN KubernetesConfig_EnvInjectionMode = 0
	// disables env injection entirely
	KubernetesConfig_ENV_INJECT_DISABLED KubernetesConfig_EnvInjectionMode = 1
	// inject non-secret env values from the Release Channel
	KubernetesConfig_ENV_INJECT_NON_SECRET_ENV KubernetesConfig_EnvInjectionMode = 2
)

// Enum value maps for KubernetesConfig_EnvInjectionMode.
var (
	KubernetesConfig_EnvInjectionMode_name = map[int32]string{
		0: "ENV_INJECT_UNKNOWN",
		1: "ENV_INJECT_DISABLED",
		2: "ENV_INJECT_NON_SECRET_ENV",
	}
	KubernetesConfig_EnvInjectionMode_value = map[string]int32{
		"ENV_INJECT_UNKNOWN":        0,
		"ENV_INJECT_DISABLED":       1,
		"ENV_INJECT_NON_SECRET_ENV": 2,
	}
)

func (x KubernetesConfig_EnvInjectionMode) Enum() *KubernetesConfig_EnvInjectionMode {
	p := new(KubernetesConfig_EnvInjectionMode)
	*p = x
	return p
}

func (x KubernetesConfig_EnvInjectionMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KubernetesConfig_EnvInjectionMode) Descriptor() protoreflect.EnumDescriptor {
	return file_prodvana_common_config_kubernetes_config_proto_enumTypes[1].Descriptor()
}

func (KubernetesConfig_EnvInjectionMode) Type() protoreflect.EnumType {
	return &file_prodvana_common_config_kubernetes_config_proto_enumTypes[1]
}

func (x KubernetesConfig_EnvInjectionMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KubernetesConfig_EnvInjectionMode.Descriptor instead.
func (KubernetesConfig_EnvInjectionMode) EnumDescriptor() ([]byte, []int) {
	return file_prodvana_common_config_kubernetes_config_proto_rawDescGZIP(), []int{1, 1}
}

type LocalConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specify a path to a local file or directory
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Specify multiple paths. They will get concatenated.
	Paths []string `protobuf:"bytes,2,rep,name=paths,proto3" json:"paths,omitempty"`
}

func (x *LocalConfig) Reset() {
	*x = LocalConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_common_config_kubernetes_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalConfig) ProtoMessage() {}

func (x *LocalConfig) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_common_config_kubernetes_config_proto_msgTypes[0]
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
	return file_prodvana_common_config_kubernetes_config_proto_rawDescGZIP(), []int{0}
}

func (x *LocalConfig) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *LocalConfig) GetPaths() []string {
	if x != nil {
		return x.Paths
	}
	return nil
}

type KubernetesConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type KubernetesConfig_Type `protobuf:"varint,1,opt,name=type,proto3,enum=prodvana.common_config.KubernetesConfig_Type" json:"type,omitempty"`
	// Types that are assignable to SourceOneof:
	//
	//	*KubernetesConfig_Inlined
	//	*KubernetesConfig_Local
	SourceOneof isKubernetesConfig_SourceOneof `protobuf_oneof:"source_oneof"`
	// Defaults to ENV_INJECT_NON_SECRET_ENV
	EnvInjectionMode KubernetesConfig_EnvInjectionMode `protobuf:"varint,4,opt,name=env_injection_mode,json=envInjectionMode,proto3,enum=prodvana.common_config.KubernetesConfig_EnvInjectionMode" json:"env_injection_mode,omitempty"`
}

func (x *KubernetesConfig) Reset() {
	*x = KubernetesConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_common_config_kubernetes_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesConfig) ProtoMessage() {}

func (x *KubernetesConfig) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_common_config_kubernetes_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesConfig.ProtoReflect.Descriptor instead.
func (*KubernetesConfig) Descriptor() ([]byte, []int) {
	return file_prodvana_common_config_kubernetes_config_proto_rawDescGZIP(), []int{1}
}

func (x *KubernetesConfig) GetType() KubernetesConfig_Type {
	if x != nil {
		return x.Type
	}
	return KubernetesConfig_UNKNOWN
}

func (m *KubernetesConfig) GetSourceOneof() isKubernetesConfig_SourceOneof {
	if m != nil {
		return m.SourceOneof
	}
	return nil
}

func (x *KubernetesConfig) GetInlined() string {
	if x, ok := x.GetSourceOneof().(*KubernetesConfig_Inlined); ok {
		return x.Inlined
	}
	return ""
}

func (x *KubernetesConfig) GetLocal() *LocalConfig {
	if x, ok := x.GetSourceOneof().(*KubernetesConfig_Local); ok {
		return x.Local
	}
	return nil
}

func (x *KubernetesConfig) GetEnvInjectionMode() KubernetesConfig_EnvInjectionMode {
	if x != nil {
		return x.EnvInjectionMode
	}
	return KubernetesConfig_ENV_INJECT_UNKNOWN
}

type isKubernetesConfig_SourceOneof interface {
	isKubernetesConfig_SourceOneof()
}

type KubernetesConfig_Inlined struct {
	Inlined string `protobuf:"bytes,2,opt,name=inlined,proto3,oneof"`
}

type KubernetesConfig_Local struct {
	Local *LocalConfig `protobuf:"bytes,3,opt,name=local,proto3,oneof"`
}

func (*KubernetesConfig_Inlined) isKubernetesConfig_SourceOneof() {}

func (*KubernetesConfig_Local) isKubernetesConfig_SourceOneof() {}

var File_prodvana_common_config_kubernetes_config_proto protoreflect.FileDescriptor

var file_prodvana_common_config_kubernetes_config_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x16, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x37, 0x0a, 0x0b, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x22, 0xd7, 0x03, 0x0a, 0x10, 0x4b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x4b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x82, 0x01, 0x02, 0x20, 0x00, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x07,
	0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x07, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x64, 0x12, 0x3b, 0x0a, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x67,
	0x0a, 0x12, 0x65, 0x6e, 0x76, 0x5f, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x39, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x45, 0x6e, 0x76, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x10, 0x65, 0x6e, 0x76, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x22, 0x32, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a,
	0x4b, 0x55, 0x42, 0x45, 0x52, 0x4e, 0x45, 0x54, 0x45, 0x53, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09,
	0x4b, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x49, 0x5a, 0x45, 0x10, 0x02, 0x22, 0x62, 0x0a, 0x10, 0x45,
	0x6e, 0x76, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12,
	0x16, 0x0a, 0x12, 0x45, 0x4e, 0x56, 0x5f, 0x49, 0x4e, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x4e, 0x56, 0x5f, 0x49,
	0x4e, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01,
	0x12, 0x1d, 0x0a, 0x19, 0x45, 0x4e, 0x56, 0x5f, 0x49, 0x4e, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x4e,
	0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54, 0x5f, 0x45, 0x4e, 0x56, 0x10, 0x02, 0x42,
	0x13, 0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x12,
	0x03, 0xf8, 0x42, 0x01, 0x42, 0x52, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x64,
	0x76, 0x61, 0x6e, 0x61, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prodvana_common_config_kubernetes_config_proto_rawDescOnce sync.Once
	file_prodvana_common_config_kubernetes_config_proto_rawDescData = file_prodvana_common_config_kubernetes_config_proto_rawDesc
)

func file_prodvana_common_config_kubernetes_config_proto_rawDescGZIP() []byte {
	file_prodvana_common_config_kubernetes_config_proto_rawDescOnce.Do(func() {
		file_prodvana_common_config_kubernetes_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_prodvana_common_config_kubernetes_config_proto_rawDescData)
	})
	return file_prodvana_common_config_kubernetes_config_proto_rawDescData
}

var file_prodvana_common_config_kubernetes_config_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_prodvana_common_config_kubernetes_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_prodvana_common_config_kubernetes_config_proto_goTypes = []interface{}{
	(KubernetesConfig_Type)(0),             // 0: prodvana.common_config.KubernetesConfig.Type
	(KubernetesConfig_EnvInjectionMode)(0), // 1: prodvana.common_config.KubernetesConfig.EnvInjectionMode
	(*LocalConfig)(nil),                    // 2: prodvana.common_config.LocalConfig
	(*KubernetesConfig)(nil),               // 3: prodvana.common_config.KubernetesConfig
}
var file_prodvana_common_config_kubernetes_config_proto_depIdxs = []int32{
	0, // 0: prodvana.common_config.KubernetesConfig.type:type_name -> prodvana.common_config.KubernetesConfig.Type
	2, // 1: prodvana.common_config.KubernetesConfig.local:type_name -> prodvana.common_config.LocalConfig
	1, // 2: prodvana.common_config.KubernetesConfig.env_injection_mode:type_name -> prodvana.common_config.KubernetesConfig.EnvInjectionMode
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_prodvana_common_config_kubernetes_config_proto_init() }
func file_prodvana_common_config_kubernetes_config_proto_init() {
	if File_prodvana_common_config_kubernetes_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prodvana_common_config_kubernetes_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_prodvana_common_config_kubernetes_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesConfig); i {
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
	file_prodvana_common_config_kubernetes_config_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*KubernetesConfig_Inlined)(nil),
		(*KubernetesConfig_Local)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_prodvana_common_config_kubernetes_config_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prodvana_common_config_kubernetes_config_proto_goTypes,
		DependencyIndexes: file_prodvana_common_config_kubernetes_config_proto_depIdxs,
		EnumInfos:         file_prodvana_common_config_kubernetes_config_proto_enumTypes,
		MessageInfos:      file_prodvana_common_config_kubernetes_config_proto_msgTypes,
	}.Build()
	File_prodvana_common_config_kubernetes_config_proto = out.File
	file_prodvana_common_config_kubernetes_config_proto_rawDesc = nil
	file_prodvana_common_config_kubernetes_config_proto_goTypes = nil
	file_prodvana_common_config_kubernetes_config_proto_depIdxs = nil
}
