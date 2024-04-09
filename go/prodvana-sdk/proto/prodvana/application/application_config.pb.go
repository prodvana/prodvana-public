// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.10
// source: prodvana/application/application_config.proto

package application

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	capability "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/capability"
	common_config "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config"
	release_channel "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/release_channel"
	workflow "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/workflow"
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

type AnnotationsConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Last9 *AnnotationsConfig_Last9 `protobuf:"bytes,1,opt,name=last9,proto3" json:"last9,omitempty"`
}

func (x *AnnotationsConfig) Reset() {
	*x = AnnotationsConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_application_application_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnotationsConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnotationsConfig) ProtoMessage() {}

func (x *AnnotationsConfig) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_application_application_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnotationsConfig.ProtoReflect.Descriptor instead.
func (*AnnotationsConfig) Descriptor() ([]byte, []int) {
	return file_prodvana_application_application_config_proto_rawDescGZIP(), []int{0}
}

func (x *AnnotationsConfig) GetLast9() *AnnotationsConfig_Last9 {
	if x != nil {
		return x.Last9
	}
	return nil
}

type ApplicationConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                 string                                                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ReleaseChannels      []*release_channel.ReleaseChannelConfig               `protobuf:"bytes,2,rep,name=release_channels,json=releaseChannels,proto3" json:"release_channels,omitempty"`
	ReleaseChannelGroups []*release_channel.ReleaseChannelGroupGeneratorConfig `protobuf:"bytes,11,rep,name=release_channel_groups,json=releaseChannelGroups,proto3" json:"release_channel_groups,omitempty"`
	Notifications        *common_config.NotificationConfig                     `protobuf:"bytes,4,opt,name=notifications,proto3" json:"notifications,omitempty"`
	Annotations          *AnnotationsConfig                                    `protobuf:"bytes,12,opt,name=annotations,proto3" json:"annotations,omitempty"`
	ReleaseSettings      *common_config.ReleaseSettings                        `protobuf:"bytes,13,opt,name=release_settings,json=releaseSettings,proto3" json:"release_settings,omitempty"`
	// deprecated
	Alerts *workflow.AlertingConfig `protobuf:"bytes,5,opt,name=alerts,proto3" json:"alerts,omitempty"`
	// deprecated
	Capabilities []*capability.CapabilityConfig `protobuf:"bytes,6,rep,name=capabilities,proto3" json:"capabilities,omitempty"`
	// deprecated
	CapabilityInstances []*capability.CapabilityInstanceConfig `protobuf:"bytes,7,rep,name=capability_instances,json=capabilityInstances,proto3" json:"capability_instances,omitempty"`
}

func (x *ApplicationConfig) Reset() {
	*x = ApplicationConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_application_application_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationConfig) ProtoMessage() {}

func (x *ApplicationConfig) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_application_application_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationConfig.ProtoReflect.Descriptor instead.
func (*ApplicationConfig) Descriptor() ([]byte, []int) {
	return file_prodvana_application_application_config_proto_rawDescGZIP(), []int{1}
}

func (x *ApplicationConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ApplicationConfig) GetReleaseChannels() []*release_channel.ReleaseChannelConfig {
	if x != nil {
		return x.ReleaseChannels
	}
	return nil
}

func (x *ApplicationConfig) GetReleaseChannelGroups() []*release_channel.ReleaseChannelGroupGeneratorConfig {
	if x != nil {
		return x.ReleaseChannelGroups
	}
	return nil
}

func (x *ApplicationConfig) GetNotifications() *common_config.NotificationConfig {
	if x != nil {
		return x.Notifications
	}
	return nil
}

func (x *ApplicationConfig) GetAnnotations() *AnnotationsConfig {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *ApplicationConfig) GetReleaseSettings() *common_config.ReleaseSettings {
	if x != nil {
		return x.ReleaseSettings
	}
	return nil
}

func (x *ApplicationConfig) GetAlerts() *workflow.AlertingConfig {
	if x != nil {
		return x.Alerts
	}
	return nil
}

func (x *ApplicationConfig) GetCapabilities() []*capability.CapabilityConfig {
	if x != nil {
		return x.Capabilities
	}
	return nil
}

func (x *ApplicationConfig) GetCapabilityInstances() []*capability.CapabilityInstanceConfig {
	if x != nil {
		return x.CapabilityInstances
	}
	return nil
}

type AnnotationsConfig_Last9 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataSource string `protobuf:"bytes,1,opt,name=data_source,json=dataSource,proto3" json:"data_source,omitempty"`
}

func (x *AnnotationsConfig_Last9) Reset() {
	*x = AnnotationsConfig_Last9{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_application_application_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnotationsConfig_Last9) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnotationsConfig_Last9) ProtoMessage() {}

func (x *AnnotationsConfig_Last9) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_application_application_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnotationsConfig_Last9.ProtoReflect.Descriptor instead.
func (*AnnotationsConfig_Last9) Descriptor() ([]byte, []int) {
	return file_prodvana_application_application_config_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AnnotationsConfig_Last9) GetDataSource() string {
	if x != nil {
		return x.DataSource
	}
	return ""
}

var File_prodvana_application_application_config_proto protoreflect.FileDescriptor

var file_prodvana_application_application_config_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x14, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x24, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f,
	0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x61, 0x70, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x70, 0x72, 0x6f,
	0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x70, 0x72,
	0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a, 0x11, 0x41, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x43,
	0x0a, 0x05, 0x6c, 0x61, 0x73, 0x74, 0x39, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c, 0x61, 0x73, 0x74, 0x39, 0x52, 0x05, 0x6c, 0x61,
	0x73, 0x74, 0x39, 0x1a, 0x31, 0x0a, 0x05, 0x4c, 0x61, 0x73, 0x74, 0x39, 0x12, 0x28, 0x0a, 0x0b,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0xa5, 0x07, 0x0a, 0x11, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3f, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0xfa, 0x42, 0x28, 0x72,
	0x26, 0x10, 0x01, 0x18, 0x3f, 0x32, 0x20, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x28, 0x5b, 0x61,
	0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2d, 0x5d, 0x2a, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d,
	0x29, 0x7b, 0x30, 0x2c, 0x31, 0x7d, 0x24, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x59, 0x0a,
	0x10, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61,
	0x6e, 0x61, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x12, 0x81, 0x01, 0x0a, 0x16, 0x72, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x76, 0x61, 0x6e, 0x61, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x0d, 0xfa, 0x42, 0x0a, 0x92, 0x01, 0x07, 0x22,
	0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x14, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x50, 0x0a, 0x0d,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x49,
	0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x52, 0x0a, 0x10, 0x72, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x0f, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x39, 0x0a,
	0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x12, 0x58, 0x0a, 0x0c, 0x63, 0x61, 0x70, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x0d, 0xfa, 0x42, 0x0a, 0x92, 0x01, 0x07, 0x22, 0x05, 0x8a,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x12, 0x6f, 0x0a, 0x14, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x61, 0x70, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42,
	0x0d, 0xfa, 0x42, 0x0a, 0x92, 0x01, 0x07, 0x22, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x13,
	0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x4a, 0x04, 0x08, 0x08, 0x10, 0x09, 0x4a,
	0x04, 0x08, 0x09, 0x10, 0x0a, 0x4a, 0x04, 0x08, 0x0a, 0x10, 0x0b, 0x52, 0x11, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x12,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x52, 0x14, 0x75, 0x73, 0x65, 0x5f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x24, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x5f, 0x64, 0x79,
	0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x42, 0x50,
	0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2d, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e,
	0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64,
	0x76, 0x61, 0x6e, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prodvana_application_application_config_proto_rawDescOnce sync.Once
	file_prodvana_application_application_config_proto_rawDescData = file_prodvana_application_application_config_proto_rawDesc
)

func file_prodvana_application_application_config_proto_rawDescGZIP() []byte {
	file_prodvana_application_application_config_proto_rawDescOnce.Do(func() {
		file_prodvana_application_application_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_prodvana_application_application_config_proto_rawDescData)
	})
	return file_prodvana_application_application_config_proto_rawDescData
}

var file_prodvana_application_application_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_prodvana_application_application_config_proto_goTypes = []interface{}{
	(*AnnotationsConfig)(nil),                                  // 0: prodvana.application.AnnotationsConfig
	(*ApplicationConfig)(nil),                                  // 1: prodvana.application.ApplicationConfig
	(*AnnotationsConfig_Last9)(nil),                            // 2: prodvana.application.AnnotationsConfig.Last9
	(*release_channel.ReleaseChannelConfig)(nil),               // 3: prodvana.release_channel.ReleaseChannelConfig
	(*release_channel.ReleaseChannelGroupGeneratorConfig)(nil), // 4: prodvana.release_channel.ReleaseChannelGroupGeneratorConfig
	(*common_config.NotificationConfig)(nil),                   // 5: prodvana.common_config.NotificationConfig
	(*common_config.ReleaseSettings)(nil),                      // 6: prodvana.common_config.ReleaseSettings
	(*workflow.AlertingConfig)(nil),                            // 7: prodvana.workflow.AlertingConfig
	(*capability.CapabilityConfig)(nil),                        // 8: prodvana.capability.CapabilityConfig
	(*capability.CapabilityInstanceConfig)(nil),                // 9: prodvana.capability.CapabilityInstanceConfig
}
var file_prodvana_application_application_config_proto_depIdxs = []int32{
	2, // 0: prodvana.application.AnnotationsConfig.last9:type_name -> prodvana.application.AnnotationsConfig.Last9
	3, // 1: prodvana.application.ApplicationConfig.release_channels:type_name -> prodvana.release_channel.ReleaseChannelConfig
	4, // 2: prodvana.application.ApplicationConfig.release_channel_groups:type_name -> prodvana.release_channel.ReleaseChannelGroupGeneratorConfig
	5, // 3: prodvana.application.ApplicationConfig.notifications:type_name -> prodvana.common_config.NotificationConfig
	0, // 4: prodvana.application.ApplicationConfig.annotations:type_name -> prodvana.application.AnnotationsConfig
	6, // 5: prodvana.application.ApplicationConfig.release_settings:type_name -> prodvana.common_config.ReleaseSettings
	7, // 6: prodvana.application.ApplicationConfig.alerts:type_name -> prodvana.workflow.AlertingConfig
	8, // 7: prodvana.application.ApplicationConfig.capabilities:type_name -> prodvana.capability.CapabilityConfig
	9, // 8: prodvana.application.ApplicationConfig.capability_instances:type_name -> prodvana.capability.CapabilityInstanceConfig
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_prodvana_application_application_config_proto_init() }
func file_prodvana_application_application_config_proto_init() {
	if File_prodvana_application_application_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prodvana_application_application_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnotationsConfig); i {
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
		file_prodvana_application_application_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationConfig); i {
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
		file_prodvana_application_application_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnotationsConfig_Last9); i {
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
			RawDescriptor: file_prodvana_application_application_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prodvana_application_application_config_proto_goTypes,
		DependencyIndexes: file_prodvana_application_application_config_proto_depIdxs,
		MessageInfos:      file_prodvana_application_application_config_proto_msgTypes,
	}.Build()
	File_prodvana_application_application_config_proto = out.File
	file_prodvana_application_application_config_proto_rawDesc = nil
	file_prodvana_application_application_config_proto_goTypes = nil
	file_prodvana_application_application_config_proto_depIdxs = nil
}
