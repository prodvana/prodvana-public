// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.10
// source: prodvana/common_config/release.proto

package common_config

import (
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

type ReleaseSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// customize what is bypassed in the fast releases (the default for rollbacks). Defaults to everything being bypassed.
	BypassSettings *ReleaseSettings_BypassSettings `protobuf:"bytes,1,opt,name=bypass_settings,json=bypassSettings,proto3" json:"bypass_settings,omitempty"`
}

func (x *ReleaseSettings) Reset() {
	*x = ReleaseSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_common_config_release_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseSettings) ProtoMessage() {}

func (x *ReleaseSettings) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_common_config_release_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseSettings.ProtoReflect.Descriptor instead.
func (*ReleaseSettings) Descriptor() ([]byte, []int) {
	return file_prodvana_common_config_release_proto_rawDescGZIP(), []int{0}
}

func (x *ReleaseSettings) GetBypassSettings() *ReleaseSettings_BypassSettings {
	if x != nil {
		return x.BypassSettings
	}
	return nil
}

type ReleaseSettings_BypassSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NoBypassProtections            bool `protobuf:"varint,1,opt,name=no_bypass_protections,json=noBypassProtections,proto3" json:"no_bypass_protections,omitempty"`
	NoBypassConvergenceExtensions  bool `protobuf:"varint,2,opt,name=no_bypass_convergence_extensions,json=noBypassConvergenceExtensions,proto3" json:"no_bypass_convergence_extensions,omitempty"`
	NoBypassApprovals              bool `protobuf:"varint,3,opt,name=no_bypass_approvals,json=noBypassApprovals,proto3" json:"no_bypass_approvals,omitempty"`
	NoBypassReleaseChannelOrdering bool `protobuf:"varint,4,opt,name=no_bypass_release_channel_ordering,json=noBypassReleaseChannelOrdering,proto3" json:"no_bypass_release_channel_ordering,omitempty"`
}

func (x *ReleaseSettings_BypassSettings) Reset() {
	*x = ReleaseSettings_BypassSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_common_config_release_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseSettings_BypassSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseSettings_BypassSettings) ProtoMessage() {}

func (x *ReleaseSettings_BypassSettings) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_common_config_release_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseSettings_BypassSettings.ProtoReflect.Descriptor instead.
func (*ReleaseSettings_BypassSettings) Descriptor() ([]byte, []int) {
	return file_prodvana_common_config_release_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ReleaseSettings_BypassSettings) GetNoBypassProtections() bool {
	if x != nil {
		return x.NoBypassProtections
	}
	return false
}

func (x *ReleaseSettings_BypassSettings) GetNoBypassConvergenceExtensions() bool {
	if x != nil {
		return x.NoBypassConvergenceExtensions
	}
	return false
}

func (x *ReleaseSettings_BypassSettings) GetNoBypassApprovals() bool {
	if x != nil {
		return x.NoBypassApprovals
	}
	return false
}

func (x *ReleaseSettings_BypassSettings) GetNoBypassReleaseChannelOrdering() bool {
	if x != nil {
		return x.NoBypassReleaseChannelOrdering
	}
	return false
}

var File_prodvana_common_config_release_proto protoreflect.FileDescriptor

var file_prodvana_common_config_release_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xfe,
	0x02, 0x0a, 0x0f, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x12, 0x5f, 0x0a, 0x0f, 0x62, 0x79, 0x70, 0x61, 0x73, 0x73, 0x5f, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x2e, 0x42, 0x79, 0x70, 0x61, 0x73, 0x73, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x52, 0x0e, 0x62, 0x79, 0x70, 0x61, 0x73, 0x73, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x1a, 0x89, 0x02, 0x0a, 0x0e, 0x42, 0x79, 0x70, 0x61, 0x73, 0x73, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x6e, 0x6f, 0x5f, 0x62, 0x79, 0x70,
	0x61, 0x73, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x6e, 0x6f, 0x42, 0x79, 0x70, 0x61, 0x73, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x47, 0x0a, 0x20, 0x6e, 0x6f,
	0x5f, 0x62, 0x79, 0x70, 0x61, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x67, 0x65,
	0x6e, 0x63, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x1d, 0x6e, 0x6f, 0x42, 0x79, 0x70, 0x61, 0x73, 0x73, 0x43, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x6e, 0x6f, 0x5f, 0x62, 0x79, 0x70, 0x61, 0x73, 0x73,
	0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x11, 0x6e, 0x6f, 0x42, 0x79, 0x70, 0x61, 0x73, 0x73, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x61, 0x6c, 0x73, 0x12, 0x4a, 0x0a, 0x22, 0x6e, 0x6f, 0x5f, 0x62, 0x79, 0x70, 0x61, 0x73, 0x73,
	0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x1e, 0x6e, 0x6f, 0x42, 0x79, 0x70, 0x61, 0x73, 0x73, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x42,
	0x52, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2d,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61,
	0x6e, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prodvana_common_config_release_proto_rawDescOnce sync.Once
	file_prodvana_common_config_release_proto_rawDescData = file_prodvana_common_config_release_proto_rawDesc
)

func file_prodvana_common_config_release_proto_rawDescGZIP() []byte {
	file_prodvana_common_config_release_proto_rawDescOnce.Do(func() {
		file_prodvana_common_config_release_proto_rawDescData = protoimpl.X.CompressGZIP(file_prodvana_common_config_release_proto_rawDescData)
	})
	return file_prodvana_common_config_release_proto_rawDescData
}

var file_prodvana_common_config_release_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_prodvana_common_config_release_proto_goTypes = []interface{}{
	(*ReleaseSettings)(nil),                // 0: prodvana.common_config.ReleaseSettings
	(*ReleaseSettings_BypassSettings)(nil), // 1: prodvana.common_config.ReleaseSettings.BypassSettings
}
var file_prodvana_common_config_release_proto_depIdxs = []int32{
	1, // 0: prodvana.common_config.ReleaseSettings.bypass_settings:type_name -> prodvana.common_config.ReleaseSettings.BypassSettings
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_prodvana_common_config_release_proto_init() }
func file_prodvana_common_config_release_proto_init() {
	if File_prodvana_common_config_release_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prodvana_common_config_release_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseSettings); i {
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
		file_prodvana_common_config_release_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseSettings_BypassSettings); i {
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
			RawDescriptor: file_prodvana_common_config_release_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prodvana_common_config_release_proto_goTypes,
		DependencyIndexes: file_prodvana_common_config_release_proto_depIdxs,
		MessageInfos:      file_prodvana_common_config_release_proto_msgTypes,
	}.Build()
	File_prodvana_common_config_release_proto = out.File
	file_prodvana_common_config_release_proto_rawDesc = nil
	file_prodvana_common_config_release_proto_goTypes = nil
	file_prodvana_common_config_release_proto_depIdxs = nil
}
