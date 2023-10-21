// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.10
// source: prodvana/object/meta.proto

package object

import (
	version "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/version"
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

type ObjectType int32

const (
	ObjectType_UNKNOWN         ObjectType = 0
	ObjectType_RUNTIME         ObjectType = 1
	ObjectType_APPLICATION     ObjectType = 2
	ObjectType_SERVICE         ObjectType = 3
	ObjectType_RELEASE_CHANNEL ObjectType = 4
	ObjectType_PROTECTION      ObjectType = 5
	ObjectType_RELEASE         ObjectType = 6
)

// Enum value maps for ObjectType.
var (
	ObjectType_name = map[int32]string{
		0: "UNKNOWN",
		1: "RUNTIME",
		2: "APPLICATION",
		3: "SERVICE",
		4: "RELEASE_CHANNEL",
		5: "PROTECTION",
		6: "RELEASE",
	}
	ObjectType_value = map[string]int32{
		"UNKNOWN":         0,
		"RUNTIME":         1,
		"APPLICATION":     2,
		"SERVICE":         3,
		"RELEASE_CHANNEL": 4,
		"PROTECTION":      5,
		"RELEASE":         6,
	}
)

func (x ObjectType) Enum() *ObjectType {
	p := new(ObjectType)
	*p = x
	return p
}

func (x ObjectType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ObjectType) Descriptor() protoreflect.EnumDescriptor {
	return file_prodvana_object_meta_proto_enumTypes[0].Descriptor()
}

func (ObjectType) Type() protoreflect.EnumType {
	return &file_prodvana_object_meta_proto_enumTypes[0]
}

func (x ObjectType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ObjectType.Descriptor instead.
func (ObjectType) EnumDescriptor() ([]byte, []int) {
	return file_prodvana_object_meta_proto_rawDescGZIP(), []int{0}
}

type ObjectMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// only set for objects who has configurations with parametrization support. At the time of this writing, only services.
	ConfigVersion  string                  `protobuf:"bytes,4,opt,name=config_version,json=configVersion,proto3" json:"config_version,omitempty"`
	Source         version.Source          `protobuf:"varint,5,opt,name=source,proto3,enum=prodvana.version.Source" json:"source,omitempty"`
	SourceMetadata *version.SourceMetadata `protobuf:"bytes,6,opt,name=source_metadata,json=sourceMetadata,proto3" json:"source_metadata,omitempty"`
	Type           ObjectType              `protobuf:"varint,7,opt,name=type,proto3,enum=prodvana.object.ObjectType" json:"type,omitempty"`
}

func (x *ObjectMeta) Reset() {
	*x = ObjectMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_object_meta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectMeta) ProtoMessage() {}

func (x *ObjectMeta) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_object_meta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectMeta.ProtoReflect.Descriptor instead.
func (*ObjectMeta) Descriptor() ([]byte, []int) {
	return file_prodvana_object_meta_proto_rawDescGZIP(), []int{0}
}

func (x *ObjectMeta) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ObjectMeta) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ObjectMeta) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ObjectMeta) GetConfigVersion() string {
	if x != nil {
		return x.ConfigVersion
	}
	return ""
}

func (x *ObjectMeta) GetSource() version.Source {
	if x != nil {
		return x.Source
	}
	return version.Source(0)
}

func (x *ObjectMeta) GetSourceMetadata() *version.SourceMetadata {
	if x != nil {
		return x.SourceMetadata
	}
	return nil
}

func (x *ObjectMeta) GetType() ObjectType {
	if x != nil {
		return x.Type
	}
	return ObjectType_UNKNOWN
}

var File_prodvana_object_meta_proto protoreflect.FileDescriptor

var file_prodvana_object_meta_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x70, 0x72,
	0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x26, 0x70,
	0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2f,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x02, 0x0a, 0x0a, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x4d, 0x65, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x76, 0x61, 0x6e, 0x61, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0f, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x2a, 0x76, 0x0a, 0x0a, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x01, 0x12,
	0x0f, 0x0a, 0x0b, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x03, 0x12, 0x13, 0x0a,
	0x0f, 0x52, 0x45, 0x4c, 0x45, 0x41, 0x53, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c,
	0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x52, 0x4f, 0x54, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x4c, 0x45, 0x41, 0x53, 0x45, 0x10, 0x06, 0x42,
	0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2d,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61,
	0x6e, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prodvana_object_meta_proto_rawDescOnce sync.Once
	file_prodvana_object_meta_proto_rawDescData = file_prodvana_object_meta_proto_rawDesc
)

func file_prodvana_object_meta_proto_rawDescGZIP() []byte {
	file_prodvana_object_meta_proto_rawDescOnce.Do(func() {
		file_prodvana_object_meta_proto_rawDescData = protoimpl.X.CompressGZIP(file_prodvana_object_meta_proto_rawDescData)
	})
	return file_prodvana_object_meta_proto_rawDescData
}

var file_prodvana_object_meta_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_prodvana_object_meta_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_prodvana_object_meta_proto_goTypes = []interface{}{
	(ObjectType)(0),                // 0: prodvana.object.ObjectType
	(*ObjectMeta)(nil),             // 1: prodvana.object.ObjectMeta
	(version.Source)(0),            // 2: prodvana.version.Source
	(*version.SourceMetadata)(nil), // 3: prodvana.version.SourceMetadata
}
var file_prodvana_object_meta_proto_depIdxs = []int32{
	2, // 0: prodvana.object.ObjectMeta.source:type_name -> prodvana.version.Source
	3, // 1: prodvana.object.ObjectMeta.source_metadata:type_name -> prodvana.version.SourceMetadata
	0, // 2: prodvana.object.ObjectMeta.type:type_name -> prodvana.object.ObjectType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_prodvana_object_meta_proto_init() }
func file_prodvana_object_meta_proto_init() {
	if File_prodvana_object_meta_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prodvana_object_meta_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectMeta); i {
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
			RawDescriptor: file_prodvana_object_meta_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prodvana_object_meta_proto_goTypes,
		DependencyIndexes: file_prodvana_object_meta_proto_depIdxs,
		EnumInfos:         file_prodvana_object_meta_proto_enumTypes,
		MessageInfos:      file_prodvana_object_meta_proto_msgTypes,
	}.Build()
	File_prodvana_object_meta_proto = out.File
	file_prodvana_object_meta_proto_rawDesc = nil
	file_prodvana_object_meta_proto_goTypes = nil
	file_prodvana_object_meta_proto_depIdxs = nil
}
