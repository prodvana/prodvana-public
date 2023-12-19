// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.10
// source: prodvana/settings/organization/roles.proto

package organization

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

type Role int32

const (
	Role_UNSPECIFIED   Role = 0
	Role_ORG_ADMIN     Role = 1
	Role_CLOUD_ADMIN   Role = 2
	Role_SERVICE_OWNER Role = 3
	Role_VIEWER        Role = 4
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "ORG_ADMIN",
		2: "CLOUD_ADMIN",
		3: "SERVICE_OWNER",
		4: "VIEWER",
	}
	Role_value = map[string]int32{
		"UNSPECIFIED":   0,
		"ORG_ADMIN":     1,
		"CLOUD_ADMIN":   2,
		"SERVICE_OWNER": 3,
		"VIEWER":        4,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_prodvana_settings_organization_roles_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_prodvana_settings_organization_roles_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_prodvana_settings_organization_roles_proto_rawDescGZIP(), []int{0}
}

var File_prodvana_settings_organization_roles_proto protoreflect.FileDescriptor

var file_prodvana_settings_organization_roles_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x70, 0x72,
	0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x56, 0x0a, 0x04,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4f, 0x52, 0x47, 0x5f, 0x41, 0x44, 0x4d,
	0x49, 0x4e, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x41, 0x44,
	0x4d, 0x49, 0x4e, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45,
	0x5f, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x56, 0x49, 0x45, 0x57,
	0x45, 0x52, 0x10, 0x04, 0x42, 0x5a, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x64,
	0x76, 0x61, 0x6e, 0x61, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prodvana_settings_organization_roles_proto_rawDescOnce sync.Once
	file_prodvana_settings_organization_roles_proto_rawDescData = file_prodvana_settings_organization_roles_proto_rawDesc
)

func file_prodvana_settings_organization_roles_proto_rawDescGZIP() []byte {
	file_prodvana_settings_organization_roles_proto_rawDescOnce.Do(func() {
		file_prodvana_settings_organization_roles_proto_rawDescData = protoimpl.X.CompressGZIP(file_prodvana_settings_organization_roles_proto_rawDescData)
	})
	return file_prodvana_settings_organization_roles_proto_rawDescData
}

var file_prodvana_settings_organization_roles_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_prodvana_settings_organization_roles_proto_goTypes = []interface{}{
	(Role)(0), // 0: prodvana.settings.organization.Role
}
var file_prodvana_settings_organization_roles_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_prodvana_settings_organization_roles_proto_init() }
func file_prodvana_settings_organization_roles_proto_init() {
	if File_prodvana_settings_organization_roles_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_prodvana_settings_organization_roles_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prodvana_settings_organization_roles_proto_goTypes,
		DependencyIndexes: file_prodvana_settings_organization_roles_proto_depIdxs,
		EnumInfos:         file_prodvana_settings_organization_roles_proto_enumTypes,
	}.Build()
	File_prodvana_settings_organization_roles_proto = out.File
	file_prodvana_settings_organization_roles_proto_rawDesc = nil
	file_prodvana_settings_organization_roles_proto_goTypes = nil
	file_prodvana_settings_organization_roles_proto_depIdxs = nil
}
