// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.10
// source: prodvana/version/source_metadata.proto

package version

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

type Source int32

const (
	Source_UNKNOWN_SOURCE     Source = 0
	Source_WEB                Source = 1
	Source_INTERACTIVE_PVNCTL Source = 2
	Source_CONFIG_FILE        Source = 3
	Source_REPO_FOLLOW        Source = 4
)

// Enum value maps for Source.
var (
	Source_name = map[int32]string{
		0: "UNKNOWN_SOURCE",
		1: "WEB",
		2: "INTERACTIVE_PVNCTL",
		3: "CONFIG_FILE",
		4: "REPO_FOLLOW",
	}
	Source_value = map[string]int32{
		"UNKNOWN_SOURCE":     0,
		"WEB":                1,
		"INTERACTIVE_PVNCTL": 2,
		"CONFIG_FILE":        3,
		"REPO_FOLLOW":        4,
	}
)

func (x Source) Enum() *Source {
	p := new(Source)
	*p = x
	return p
}

func (x Source) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Source) Descriptor() protoreflect.EnumDescriptor {
	return file_prodvana_version_source_metadata_proto_enumTypes[0].Descriptor()
}

func (Source) Type() protoreflect.EnumType {
	return &file_prodvana_version_source_metadata_proto_enumTypes[0]
}

func (x Source) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Source.Descriptor instead.
func (Source) EnumDescriptor() ([]byte, []int) {
	return file_prodvana_version_source_metadata_proto_rawDescGZIP(), []int{0}
}

// all of these fields are optional and only set if it makes sense for a given source.
type SourceMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// full repo URL, like git@github.com:foo/bar.git or https://github.com/foo/bar
	RepoUrl  string `protobuf:"bytes,1,opt,name=repo_url,json=repoUrl,proto3" json:"repo_url,omitempty"`
	FilePath string `protobuf:"bytes,2,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	Commit   string `protobuf:"bytes,3,opt,name=commit,proto3" json:"commit,omitempty"`
	// set internally, automatically, by Prodvana if the action was initiated by a user
	UserId string `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *SourceMetadata) Reset() {
	*x = SourceMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_version_source_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceMetadata) ProtoMessage() {}

func (x *SourceMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_version_source_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceMetadata.ProtoReflect.Descriptor instead.
func (*SourceMetadata) Descriptor() ([]byte, []int) {
	return file_prodvana_version_source_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *SourceMetadata) GetRepoUrl() string {
	if x != nil {
		return x.RepoUrl
	}
	return ""
}

func (x *SourceMetadata) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *SourceMetadata) GetCommit() string {
	if x != nil {
		return x.Commit
	}
	return ""
}

func (x *SourceMetadata) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_prodvana_version_source_metadata_proto protoreflect.FileDescriptor

var file_prodvana_version_source_metadata_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61,
	0x6e, 0x61, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x79, 0x0a, 0x0e, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08,
	0x72, 0x65, 0x70, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x72, 0x65, 0x70, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x2a, 0x5f, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43,
	0x45, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x57, 0x45, 0x42, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12,
	0x49, 0x4e, 0x54, 0x45, 0x52, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x50, 0x56, 0x4e, 0x43,
	0x54, 0x4c, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x46,
	0x49, 0x4c, 0x45, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x45, 0x50, 0x4f, 0x5f, 0x46, 0x4f,
	0x4c, 0x4c, 0x4f, 0x57, 0x10, 0x04, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x67, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prodvana_version_source_metadata_proto_rawDescOnce sync.Once
	file_prodvana_version_source_metadata_proto_rawDescData = file_prodvana_version_source_metadata_proto_rawDesc
)

func file_prodvana_version_source_metadata_proto_rawDescGZIP() []byte {
	file_prodvana_version_source_metadata_proto_rawDescOnce.Do(func() {
		file_prodvana_version_source_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_prodvana_version_source_metadata_proto_rawDescData)
	})
	return file_prodvana_version_source_metadata_proto_rawDescData
}

var file_prodvana_version_source_metadata_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_prodvana_version_source_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_prodvana_version_source_metadata_proto_goTypes = []interface{}{
	(Source)(0),            // 0: prodvana.version.Source
	(*SourceMetadata)(nil), // 1: prodvana.version.SourceMetadata
}
var file_prodvana_version_source_metadata_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_prodvana_version_source_metadata_proto_init() }
func file_prodvana_version_source_metadata_proto_init() {
	if File_prodvana_version_source_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prodvana_version_source_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceMetadata); i {
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
			RawDescriptor: file_prodvana_version_source_metadata_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prodvana_version_source_metadata_proto_goTypes,
		DependencyIndexes: file_prodvana_version_source_metadata_proto_depIdxs,
		EnumInfos:         file_prodvana_version_source_metadata_proto_enumTypes,
		MessageInfos:      file_prodvana_version_source_metadata_proto_msgTypes,
	}.Build()
	File_prodvana_version_source_metadata_proto = out.File
	file_prodvana_version_source_metadata_proto_rawDesc = nil
	file_prodvana_version_source_metadata_proto_goTypes = nil
	file_prodvana_version_source_metadata_proto_depIdxs = nil
}
