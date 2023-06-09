// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.10
// source: prodvana/scm/types.proto

package scm

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

type ScmType int32

const (
	ScmType_UNKNOWN ScmType = 0
	ScmType_GITHUB  ScmType = 1
)

// Enum value maps for ScmType.
var (
	ScmType_name = map[int32]string{
		0: "UNKNOWN",
		1: "GITHUB",
	}
	ScmType_value = map[string]int32{
		"UNKNOWN": 0,
		"GITHUB":  1,
	}
)

func (x ScmType) Enum() *ScmType {
	p := new(ScmType)
	*p = x
	return p
}

func (x ScmType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ScmType) Descriptor() protoreflect.EnumDescriptor {
	return file_prodvana_scm_types_proto_enumTypes[0].Descriptor()
}

func (ScmType) Type() protoreflect.EnumType {
	return &file_prodvana_scm_types_proto_enumTypes[0]
}

func (x ScmType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ScmType.Descriptor instead.
func (ScmType) EnumDescriptor() ([]byte, []int) {
	return file_prodvana_scm_types_proto_rawDescGZIP(), []int{0}
}

var File_prodvana_scm_types_proto protoreflect.FileDescriptor

var file_prodvana_scm_types_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x73, 0x63, 0x6d, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x64,
	0x76, 0x61, 0x6e, 0x61, 0x2e, 0x73, 0x63, 0x6d, 0x2a, 0x22, 0x0a, 0x07, 0x53, 0x63, 0x6d, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x47, 0x49, 0x54, 0x48, 0x55, 0x42, 0x10, 0x01, 0x42, 0x48, 0x5a, 0x46,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76,
	0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2d, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61,
	0x6e, 0x61, 0x2f, 0x73, 0x63, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prodvana_scm_types_proto_rawDescOnce sync.Once
	file_prodvana_scm_types_proto_rawDescData = file_prodvana_scm_types_proto_rawDesc
)

func file_prodvana_scm_types_proto_rawDescGZIP() []byte {
	file_prodvana_scm_types_proto_rawDescOnce.Do(func() {
		file_prodvana_scm_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_prodvana_scm_types_proto_rawDescData)
	})
	return file_prodvana_scm_types_proto_rawDescData
}

var file_prodvana_scm_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_prodvana_scm_types_proto_goTypes = []interface{}{
	(ScmType)(0), // 0: prodvana.scm.ScmType
}
var file_prodvana_scm_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_prodvana_scm_types_proto_init() }
func file_prodvana_scm_types_proto_init() {
	if File_prodvana_scm_types_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_prodvana_scm_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prodvana_scm_types_proto_goTypes,
		DependencyIndexes: file_prodvana_scm_types_proto_depIdxs,
		EnumInfos:         file_prodvana_scm_types_proto_enumTypes,
	}.Build()
	File_prodvana_scm_types_proto = out.File
	file_prodvana_scm_types_proto_rawDesc = nil
	file_prodvana_scm_types_proto_goTypes = nil
	file_prodvana_scm_types_proto_depIdxs = nil
}
