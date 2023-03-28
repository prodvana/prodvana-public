// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.10
// source: prodvana/common_config/parameters.proto

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

type ProgramChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ProgramChange) Reset() {
	*x = ProgramChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_common_config_parameters_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProgramChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgramChange) ProtoMessage() {}

func (x *ProgramChange) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_common_config_parameters_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProgramChange.ProtoReflect.Descriptor instead.
func (*ProgramChange) Descriptor() ([]byte, []int) {
	return file_prodvana_common_config_parameters_proto_rawDescGZIP(), []int{0}
}

func (x *ProgramChange) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type StringParameterDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InitialValue string `protobuf:"bytes,1,opt,name=initial_value,json=initialValue,proto3" json:"initial_value,omitempty"`
}

func (x *StringParameterDefinition) Reset() {
	*x = StringParameterDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_common_config_parameters_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringParameterDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringParameterDefinition) ProtoMessage() {}

func (x *StringParameterDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_common_config_parameters_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringParameterDefinition.ProtoReflect.Descriptor instead.
func (*StringParameterDefinition) Descriptor() ([]byte, []int) {
	return file_prodvana_common_config_parameters_proto_rawDescGZIP(), []int{1}
}

func (x *StringParameterDefinition) GetInitialValue() string {
	if x != nil {
		return x.InitialValue
	}
	return ""
}

type DockerImageParameterDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// empty not a valid value
	InitialValue      string                                    `protobuf:"bytes,1,opt,name=initial_value,json=initialValue,proto3" json:"initial_value,omitempty"`
	ImageRegistryInfo *ImageRegistryInfo                        `protobuf:"bytes,2,opt,name=image_registry_info,json=imageRegistryInfo,proto3" json:"image_registry_info,omitempty"`
	Changes           []*DockerImageParameterDefinition_Changes `protobuf:"bytes,3,rep,name=changes,proto3" json:"changes,omitempty"`
}

func (x *DockerImageParameterDefinition) Reset() {
	*x = DockerImageParameterDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_common_config_parameters_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DockerImageParameterDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DockerImageParameterDefinition) ProtoMessage() {}

func (x *DockerImageParameterDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_common_config_parameters_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DockerImageParameterDefinition.ProtoReflect.Descriptor instead.
func (*DockerImageParameterDefinition) Descriptor() ([]byte, []int) {
	return file_prodvana_common_config_parameters_proto_rawDescGZIP(), []int{2}
}

func (x *DockerImageParameterDefinition) GetInitialValue() string {
	if x != nil {
		return x.InitialValue
	}
	return ""
}

func (x *DockerImageParameterDefinition) GetImageRegistryInfo() *ImageRegistryInfo {
	if x != nil {
		return x.ImageRegistryInfo
	}
	return nil
}

func (x *DockerImageParameterDefinition) GetChanges() []*DockerImageParameterDefinition_Changes {
	if x != nil {
		return x.Changes
	}
	return nil
}

type FixedReplicaChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FixedReplicaChange) Reset() {
	*x = FixedReplicaChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_common_config_parameters_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FixedReplicaChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FixedReplicaChange) ProtoMessage() {}

func (x *FixedReplicaChange) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_common_config_parameters_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FixedReplicaChange.ProtoReflect.Descriptor instead.
func (*FixedReplicaChange) Descriptor() ([]byte, []int) {
	return file_prodvana_common_config_parameters_proto_rawDescGZIP(), []int{3}
}

type IntParameterDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InitialValue int64                             `protobuf:"varint,1,opt,name=initial_value,json=initialValue,proto3" json:"initial_value,omitempty"`
	Changes      []*IntParameterDefinition_Changes `protobuf:"bytes,2,rep,name=changes,proto3" json:"changes,omitempty"`
}

func (x *IntParameterDefinition) Reset() {
	*x = IntParameterDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_common_config_parameters_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntParameterDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntParameterDefinition) ProtoMessage() {}

func (x *IntParameterDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_common_config_parameters_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntParameterDefinition.ProtoReflect.Descriptor instead.
func (*IntParameterDefinition) Descriptor() ([]byte, []int) {
	return file_prodvana_common_config_parameters_proto_rawDescGZIP(), []int{4}
}

func (x *IntParameterDefinition) GetInitialValue() int64 {
	if x != nil {
		return x.InitialValue
	}
	return 0
}

func (x *IntParameterDefinition) GetChanges() []*IntParameterDefinition_Changes {
	if x != nil {
		return x.Changes
	}
	return nil
}

type ParameterDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// parameter name, used in substitutions
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// optional description for display purposes
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Types that are assignable to ConfigOneof:
	//
	//	*ParameterDefinition_String_
	//	*ParameterDefinition_DockerImage
	//	*ParameterDefinition_Int
	ConfigOneof isParameterDefinition_ConfigOneof `protobuf_oneof:"config_oneof"`
}

func (x *ParameterDefinition) Reset() {
	*x = ParameterDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_common_config_parameters_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParameterDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParameterDefinition) ProtoMessage() {}

func (x *ParameterDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_common_config_parameters_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParameterDefinition.ProtoReflect.Descriptor instead.
func (*ParameterDefinition) Descriptor() ([]byte, []int) {
	return file_prodvana_common_config_parameters_proto_rawDescGZIP(), []int{5}
}

func (x *ParameterDefinition) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ParameterDefinition) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (m *ParameterDefinition) GetConfigOneof() isParameterDefinition_ConfigOneof {
	if m != nil {
		return m.ConfigOneof
	}
	return nil
}

func (x *ParameterDefinition) GetString_() *StringParameterDefinition {
	if x, ok := x.GetConfigOneof().(*ParameterDefinition_String_); ok {
		return x.String_
	}
	return nil
}

func (x *ParameterDefinition) GetDockerImage() *DockerImageParameterDefinition {
	if x, ok := x.GetConfigOneof().(*ParameterDefinition_DockerImage); ok {
		return x.DockerImage
	}
	return nil
}

func (x *ParameterDefinition) GetInt() *IntParameterDefinition {
	if x, ok := x.GetConfigOneof().(*ParameterDefinition_Int); ok {
		return x.Int
	}
	return nil
}

type isParameterDefinition_ConfigOneof interface {
	isParameterDefinition_ConfigOneof()
}

type ParameterDefinition_String_ struct {
	String_ *StringParameterDefinition `protobuf:"bytes,3,opt,name=string,proto3,oneof"`
}

type ParameterDefinition_DockerImage struct {
	DockerImage *DockerImageParameterDefinition `protobuf:"bytes,4,opt,name=docker_image,json=dockerImage,proto3,oneof"`
}

type ParameterDefinition_Int struct {
	Int *IntParameterDefinition `protobuf:"bytes,5,opt,name=int,proto3,oneof"`
}

func (*ParameterDefinition_String_) isParameterDefinition_ConfigOneof() {}

func (*ParameterDefinition_DockerImage) isParameterDefinition_ConfigOneof() {}

func (*ParameterDefinition_Int) isParameterDefinition_ConfigOneof() {}

type ParameterValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are assignable to ValueOneof:
	//
	//	*ParameterValue_String_
	//	*ParameterValue_DockerImage
	//	*ParameterValue_Int
	ValueOneof isParameterValue_ValueOneof `protobuf_oneof:"value_oneof"`
}

func (x *ParameterValue) Reset() {
	*x = ParameterValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_common_config_parameters_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParameterValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParameterValue) ProtoMessage() {}

func (x *ParameterValue) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_common_config_parameters_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParameterValue.ProtoReflect.Descriptor instead.
func (*ParameterValue) Descriptor() ([]byte, []int) {
	return file_prodvana_common_config_parameters_proto_rawDescGZIP(), []int{6}
}

func (x *ParameterValue) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *ParameterValue) GetValueOneof() isParameterValue_ValueOneof {
	if m != nil {
		return m.ValueOneof
	}
	return nil
}

func (x *ParameterValue) GetString_() string {
	if x, ok := x.GetValueOneof().(*ParameterValue_String_); ok {
		return x.String_
	}
	return ""
}

func (x *ParameterValue) GetDockerImage() string {
	if x, ok := x.GetValueOneof().(*ParameterValue_DockerImage); ok {
		return x.DockerImage
	}
	return ""
}

func (x *ParameterValue) GetInt() int64 {
	if x, ok := x.GetValueOneof().(*ParameterValue_Int); ok {
		return x.Int
	}
	return 0
}

type isParameterValue_ValueOneof interface {
	isParameterValue_ValueOneof()
}

type ParameterValue_String_ struct {
	String_ string `protobuf:"bytes,2,opt,name=string,proto3,oneof"`
}

type ParameterValue_DockerImage struct {
	DockerImage string `protobuf:"bytes,3,opt,name=docker_image,json=dockerImage,proto3,oneof"`
}

type ParameterValue_Int struct {
	Int int64 `protobuf:"varint,4,opt,name=int,proto3,oneof"`
}

func (*ParameterValue_String_) isParameterValue_ValueOneof() {}

func (*ParameterValue_DockerImage) isParameterValue_ValueOneof() {}

func (*ParameterValue_Int) isParameterValue_ValueOneof() {}

type ParametersConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parameters []*ParameterDefinition `protobuf:"bytes,1,rep,name=parameters,proto3" json:"parameters,omitempty"`
}

func (x *ParametersConfig) Reset() {
	*x = ParametersConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_common_config_parameters_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParametersConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParametersConfig) ProtoMessage() {}

func (x *ParametersConfig) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_common_config_parameters_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParametersConfig.ProtoReflect.Descriptor instead.
func (*ParametersConfig) Descriptor() ([]byte, []int) {
	return file_prodvana_common_config_parameters_proto_rawDescGZIP(), []int{7}
}

func (x *ParametersConfig) GetParameters() []*ParameterDefinition {
	if x != nil {
		return x.Parameters
	}
	return nil
}

type DockerImageParameterDefinition_Changes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Oneof:
	//
	//	*DockerImageParameterDefinition_Changes_Program
	Oneof isDockerImageParameterDefinition_Changes_Oneof `protobuf_oneof:"oneof"`
}

func (x *DockerImageParameterDefinition_Changes) Reset() {
	*x = DockerImageParameterDefinition_Changes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_common_config_parameters_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DockerImageParameterDefinition_Changes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DockerImageParameterDefinition_Changes) ProtoMessage() {}

func (x *DockerImageParameterDefinition_Changes) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_common_config_parameters_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DockerImageParameterDefinition_Changes.ProtoReflect.Descriptor instead.
func (*DockerImageParameterDefinition_Changes) Descriptor() ([]byte, []int) {
	return file_prodvana_common_config_parameters_proto_rawDescGZIP(), []int{2, 0}
}

func (m *DockerImageParameterDefinition_Changes) GetOneof() isDockerImageParameterDefinition_Changes_Oneof {
	if m != nil {
		return m.Oneof
	}
	return nil
}

func (x *DockerImageParameterDefinition_Changes) GetProgram() *ProgramChange {
	if x, ok := x.GetOneof().(*DockerImageParameterDefinition_Changes_Program); ok {
		return x.Program
	}
	return nil
}

type isDockerImageParameterDefinition_Changes_Oneof interface {
	isDockerImageParameterDefinition_Changes_Oneof()
}

type DockerImageParameterDefinition_Changes_Program struct {
	Program *ProgramChange `protobuf:"bytes,1,opt,name=program,proto3,oneof"`
}

func (*DockerImageParameterDefinition_Changes_Program) isDockerImageParameterDefinition_Changes_Oneof() {
}

type IntParameterDefinition_Changes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Oneof:
	//
	//	*IntParameterDefinition_Changes_FixedReplica
	Oneof isIntParameterDefinition_Changes_Oneof `protobuf_oneof:"oneof"`
}

func (x *IntParameterDefinition_Changes) Reset() {
	*x = IntParameterDefinition_Changes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_common_config_parameters_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntParameterDefinition_Changes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntParameterDefinition_Changes) ProtoMessage() {}

func (x *IntParameterDefinition_Changes) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_common_config_parameters_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntParameterDefinition_Changes.ProtoReflect.Descriptor instead.
func (*IntParameterDefinition_Changes) Descriptor() ([]byte, []int) {
	return file_prodvana_common_config_parameters_proto_rawDescGZIP(), []int{4, 0}
}

func (m *IntParameterDefinition_Changes) GetOneof() isIntParameterDefinition_Changes_Oneof {
	if m != nil {
		return m.Oneof
	}
	return nil
}

func (x *IntParameterDefinition_Changes) GetFixedReplica() *FixedReplicaChange {
	if x, ok := x.GetOneof().(*IntParameterDefinition_Changes_FixedReplica); ok {
		return x.FixedReplica
	}
	return nil
}

type isIntParameterDefinition_Changes_Oneof interface {
	isIntParameterDefinition_Changes_Oneof()
}

type IntParameterDefinition_Changes_FixedReplica struct {
	FixedReplica *FixedReplicaChange `protobuf:"bytes,1,opt,name=fixed_replica,json=fixedReplica,proto3,oneof"`
}

func (*IntParameterDefinition_Changes_FixedReplica) isIntParameterDefinition_Changes_Oneof() {}

var File_prodvana_common_config_parameters_proto protoreflect.FileDescriptor

var file_prodvana_common_config_parameters_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x70, 0x72, 0x6f, 0x64, 0x76,
	0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x1a, 0x24, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x2c, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x40,
	0x0a, 0x19, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0xf8, 0x02, 0x0a, 0x1e, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x0d, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x10, 0x01, 0x52, 0x0c, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x63, 0x0a, 0x13, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x11, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x67, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61,
	0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x42, 0x0d, 0xfa, 0x42, 0x0a, 0x92, 0x01, 0x07, 0x22,
	0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x1a,
	0x5a, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x48, 0x00, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x42, 0x0c, 0x0a,
	0x05, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0x14, 0x0a, 0x12, 0x46,
	0x69, 0x78, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x22, 0x8a, 0x02, 0x0a, 0x16, 0x49, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x5f, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x36, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x49, 0x6e, 0x74, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x42, 0x0d, 0xfa, 0x42, 0x0a, 0x92,
	0x01, 0x07, 0x22, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x73, 0x1a, 0x6a, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x51, 0x0a,
	0x0d, 0x66, 0x69, 0x78, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x46, 0x69,
	0x78, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x48, 0x00, 0x52, 0x0c, 0x66, 0x69, 0x78, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x42, 0x0c, 0x0a, 0x05, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0xd7,
	0x02, 0x0a, 0x13, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x44, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x12, 0x5b, 0x0a, 0x0c, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76,
	0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x00, 0x52, 0x0b, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x42, 0x0a, 0x03, 0x69, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x49, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x03,
	0x69, 0x6e, 0x74, 0x42, 0x13, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6f, 0x6e,
	0x65, 0x6f, 0x66, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0x94, 0x01, 0x0a, 0x0e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x12, 0x23, 0x0a, 0x0c, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x03, 0x69, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x03, 0x69, 0x6e, 0x74, 0x42, 0x12, 0x0a, 0x0b, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22,
	0x6e, 0x0a, 0x10, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x5a, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61,
	0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0d, 0xfa, 0x42, 0x0a, 0x92, 0x01, 0x07, 0x22, 0x05, 0x8a, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x42,
	0x52, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2d,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61,
	0x6e, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prodvana_common_config_parameters_proto_rawDescOnce sync.Once
	file_prodvana_common_config_parameters_proto_rawDescData = file_prodvana_common_config_parameters_proto_rawDesc
)

func file_prodvana_common_config_parameters_proto_rawDescGZIP() []byte {
	file_prodvana_common_config_parameters_proto_rawDescOnce.Do(func() {
		file_prodvana_common_config_parameters_proto_rawDescData = protoimpl.X.CompressGZIP(file_prodvana_common_config_parameters_proto_rawDescData)
	})
	return file_prodvana_common_config_parameters_proto_rawDescData
}

var file_prodvana_common_config_parameters_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_prodvana_common_config_parameters_proto_goTypes = []interface{}{
	(*ProgramChange)(nil),                          // 0: prodvana.common_config.ProgramChange
	(*StringParameterDefinition)(nil),              // 1: prodvana.common_config.StringParameterDefinition
	(*DockerImageParameterDefinition)(nil),         // 2: prodvana.common_config.DockerImageParameterDefinition
	(*FixedReplicaChange)(nil),                     // 3: prodvana.common_config.FixedReplicaChange
	(*IntParameterDefinition)(nil),                 // 4: prodvana.common_config.IntParameterDefinition
	(*ParameterDefinition)(nil),                    // 5: prodvana.common_config.ParameterDefinition
	(*ParameterValue)(nil),                         // 6: prodvana.common_config.ParameterValue
	(*ParametersConfig)(nil),                       // 7: prodvana.common_config.ParametersConfig
	(*DockerImageParameterDefinition_Changes)(nil), // 8: prodvana.common_config.DockerImageParameterDefinition.Changes
	(*IntParameterDefinition_Changes)(nil),         // 9: prodvana.common_config.IntParameterDefinition.Changes
	(*ImageRegistryInfo)(nil),                      // 10: prodvana.common_config.ImageRegistryInfo
}
var file_prodvana_common_config_parameters_proto_depIdxs = []int32{
	10, // 0: prodvana.common_config.DockerImageParameterDefinition.image_registry_info:type_name -> prodvana.common_config.ImageRegistryInfo
	8,  // 1: prodvana.common_config.DockerImageParameterDefinition.changes:type_name -> prodvana.common_config.DockerImageParameterDefinition.Changes
	9,  // 2: prodvana.common_config.IntParameterDefinition.changes:type_name -> prodvana.common_config.IntParameterDefinition.Changes
	1,  // 3: prodvana.common_config.ParameterDefinition.string:type_name -> prodvana.common_config.StringParameterDefinition
	2,  // 4: prodvana.common_config.ParameterDefinition.docker_image:type_name -> prodvana.common_config.DockerImageParameterDefinition
	4,  // 5: prodvana.common_config.ParameterDefinition.int:type_name -> prodvana.common_config.IntParameterDefinition
	5,  // 6: prodvana.common_config.ParametersConfig.parameters:type_name -> prodvana.common_config.ParameterDefinition
	0,  // 7: prodvana.common_config.DockerImageParameterDefinition.Changes.program:type_name -> prodvana.common_config.ProgramChange
	3,  // 8: prodvana.common_config.IntParameterDefinition.Changes.fixed_replica:type_name -> prodvana.common_config.FixedReplicaChange
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_prodvana_common_config_parameters_proto_init() }
func file_prodvana_common_config_parameters_proto_init() {
	if File_prodvana_common_config_parameters_proto != nil {
		return
	}
	file_prodvana_common_config_program_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_prodvana_common_config_parameters_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProgramChange); i {
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
		file_prodvana_common_config_parameters_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringParameterDefinition); i {
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
		file_prodvana_common_config_parameters_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DockerImageParameterDefinition); i {
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
		file_prodvana_common_config_parameters_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FixedReplicaChange); i {
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
		file_prodvana_common_config_parameters_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntParameterDefinition); i {
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
		file_prodvana_common_config_parameters_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParameterDefinition); i {
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
		file_prodvana_common_config_parameters_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParameterValue); i {
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
		file_prodvana_common_config_parameters_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParametersConfig); i {
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
		file_prodvana_common_config_parameters_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DockerImageParameterDefinition_Changes); i {
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
		file_prodvana_common_config_parameters_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntParameterDefinition_Changes); i {
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
	file_prodvana_common_config_parameters_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*ParameterDefinition_String_)(nil),
		(*ParameterDefinition_DockerImage)(nil),
		(*ParameterDefinition_Int)(nil),
	}
	file_prodvana_common_config_parameters_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*ParameterValue_String_)(nil),
		(*ParameterValue_DockerImage)(nil),
		(*ParameterValue_Int)(nil),
	}
	file_prodvana_common_config_parameters_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*DockerImageParameterDefinition_Changes_Program)(nil),
	}
	file_prodvana_common_config_parameters_proto_msgTypes[9].OneofWrappers = []interface{}{
		(*IntParameterDefinition_Changes_FixedReplica)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_prodvana_common_config_parameters_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prodvana_common_config_parameters_proto_goTypes,
		DependencyIndexes: file_prodvana_common_config_parameters_proto_depIdxs,
		MessageInfos:      file_prodvana_common_config_parameters_proto_msgTypes,
	}.Build()
	File_prodvana_common_config_parameters_proto = out.File
	file_prodvana_common_config_parameters_proto_rawDesc = nil
	file_prodvana_common_config_parameters_proto_goTypes = nil
	file_prodvana_common_config_parameters_proto_depIdxs = nil
}