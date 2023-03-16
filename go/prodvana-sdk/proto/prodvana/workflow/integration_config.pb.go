// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.10
// source: prodvana/workflow/integration_config.proto

package workflow

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

type AlertingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagerduty *AlertingConfig_PagerDuty `protobuf:"bytes,1,opt,name=pagerduty,proto3" json:"pagerduty,omitempty"`
}

func (x *AlertingConfig) Reset() {
	*x = AlertingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_workflow_integration_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertingConfig) ProtoMessage() {}

func (x *AlertingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_workflow_integration_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertingConfig.ProtoReflect.Descriptor instead.
func (*AlertingConfig) Descriptor() ([]byte, []int) {
	return file_prodvana_workflow_integration_config_proto_rawDescGZIP(), []int{0}
}

func (x *AlertingConfig) GetPagerduty() *AlertingConfig_PagerDuty {
	if x != nil {
		return x.Pagerduty
	}
	return nil
}

type AnnotationsConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Honeycomb *AnnotationsConfig_Honeycomb `protobuf:"bytes,1,opt,name=honeycomb,proto3" json:"honeycomb,omitempty"`
}

func (x *AnnotationsConfig) Reset() {
	*x = AnnotationsConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_workflow_integration_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnotationsConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnotationsConfig) ProtoMessage() {}

func (x *AnnotationsConfig) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_workflow_integration_config_proto_msgTypes[1]
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
	return file_prodvana_workflow_integration_config_proto_rawDescGZIP(), []int{1}
}

func (x *AnnotationsConfig) GetHoneycomb() *AnnotationsConfig_Honeycomb {
	if x != nil {
		return x.Honeycomb
	}
	return nil
}

type TokenConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenSecretKey     string `protobuf:"bytes,1,opt,name=token_secret_key,json=tokenSecretKey,proto3" json:"token_secret_key,omitempty"`
	TokenSecretVersion string `protobuf:"bytes,2,opt,name=token_secret_version,json=tokenSecretVersion,proto3" json:"token_secret_version,omitempty"`
}

func (x *TokenConfig) Reset() {
	*x = TokenConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_workflow_integration_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenConfig) ProtoMessage() {}

func (x *TokenConfig) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_workflow_integration_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenConfig.ProtoReflect.Descriptor instead.
func (*TokenConfig) Descriptor() ([]byte, []int) {
	return file_prodvana_workflow_integration_config_proto_rawDescGZIP(), []int{2}
}

func (x *TokenConfig) GetTokenSecretKey() string {
	if x != nil {
		return x.TokenSecretKey
	}
	return ""
}

func (x *TokenConfig) GetTokenSecretVersion() string {
	if x != nil {
		return x.TokenSecretVersion
	}
	return ""
}

type IntegrationConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ConfigOneof:
	//
	//	*IntegrationConfig_SlackConfig
	//	*IntegrationConfig_PagerdutyConfig
	ConfigOneof isIntegrationConfig_ConfigOneof `protobuf_oneof:"config_oneof"`
}

func (x *IntegrationConfig) Reset() {
	*x = IntegrationConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_workflow_integration_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegrationConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegrationConfig) ProtoMessage() {}

func (x *IntegrationConfig) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_workflow_integration_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegrationConfig.ProtoReflect.Descriptor instead.
func (*IntegrationConfig) Descriptor() ([]byte, []int) {
	return file_prodvana_workflow_integration_config_proto_rawDescGZIP(), []int{3}
}

func (m *IntegrationConfig) GetConfigOneof() isIntegrationConfig_ConfigOneof {
	if m != nil {
		return m.ConfigOneof
	}
	return nil
}

func (x *IntegrationConfig) GetSlackConfig() *TokenConfig {
	if x, ok := x.GetConfigOneof().(*IntegrationConfig_SlackConfig); ok {
		return x.SlackConfig
	}
	return nil
}

func (x *IntegrationConfig) GetPagerdutyConfig() *TokenConfig {
	if x, ok := x.GetConfigOneof().(*IntegrationConfig_PagerdutyConfig); ok {
		return x.PagerdutyConfig
	}
	return nil
}

type isIntegrationConfig_ConfigOneof interface {
	isIntegrationConfig_ConfigOneof()
}

type IntegrationConfig_SlackConfig struct {
	SlackConfig *TokenConfig `protobuf:"bytes,1,opt,name=slack_config,json=slackConfig,proto3,oneof"`
}

type IntegrationConfig_PagerdutyConfig struct {
	PagerdutyConfig *TokenConfig `protobuf:"bytes,2,opt,name=pagerduty_config,json=pagerdutyConfig,proto3,oneof"`
}

func (*IntegrationConfig_SlackConfig) isIntegrationConfig_ConfigOneof() {}

func (*IntegrationConfig_PagerdutyConfig) isIntegrationConfig_ConfigOneof() {}

type AlertingConfig_PagerDuty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *AlertingConfig_PagerDuty) Reset() {
	*x = AlertingConfig_PagerDuty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_workflow_integration_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertingConfig_PagerDuty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertingConfig_PagerDuty) ProtoMessage() {}

func (x *AlertingConfig_PagerDuty) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_workflow_integration_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertingConfig_PagerDuty.ProtoReflect.Descriptor instead.
func (*AlertingConfig_PagerDuty) Descriptor() ([]byte, []int) {
	return file_prodvana_workflow_integration_config_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AlertingConfig_PagerDuty) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type AnnotationsConfig_Honeycomb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Environment string `protobuf:"bytes,1,opt,name=environment,proto3" json:"environment,omitempty"`
	Dataset     string `protobuf:"bytes,2,opt,name=dataset,proto3" json:"dataset,omitempty"`
}

func (x *AnnotationsConfig_Honeycomb) Reset() {
	*x = AnnotationsConfig_Honeycomb{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_workflow_integration_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnotationsConfig_Honeycomb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnotationsConfig_Honeycomb) ProtoMessage() {}

func (x *AnnotationsConfig_Honeycomb) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_workflow_integration_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnotationsConfig_Honeycomb.ProtoReflect.Descriptor instead.
func (*AnnotationsConfig_Honeycomb) Descriptor() ([]byte, []int) {
	return file_prodvana_workflow_integration_config_proto_rawDescGZIP(), []int{1, 0}
}

func (x *AnnotationsConfig_Honeycomb) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *AnnotationsConfig_Honeycomb) GetDataset() string {
	if x != nil {
		return x.Dataset
	}
	return ""
}

var File_prodvana_workflow_integration_config_proto protoreflect.FileDescriptor

var file_prodvana_workflow_integration_config_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x70, 0x72,
	0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a, 0x0e, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x49, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x72, 0x64, 0x75, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x44, 0x75, 0x74, 0x79, 0x52, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x72, 0x64, 0x75, 0x74, 0x79, 0x1a, 0x2e, 0x0a, 0x09, 0x50, 0x61, 0x67, 0x65, 0x72, 0x44,
	0x75, 0x74, 0x79, 0x12, 0x21, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xb3, 0x01, 0x0a, 0x11, 0x41, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4c, 0x0a, 0x09,
	0x68, 0x6f, 0x6e, 0x65, 0x79, 0x63, 0x6f, 0x6d, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x48, 0x6f, 0x6e, 0x65, 0x79, 0x63, 0x6f, 0x6d, 0x62, 0x52,
	0x09, 0x68, 0x6f, 0x6e, 0x65, 0x79, 0x63, 0x6f, 0x6d, 0x62, 0x1a, 0x50, 0x0a, 0x09, 0x48, 0x6f,
	0x6e, 0x65, 0x79, 0x63, 0x6f, 0x6d, 0x62, 0x12, 0x29, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x22, 0x69, 0x0a, 0x0b,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x28, 0x0a, 0x10, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x14, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x12, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xc0, 0x01, 0x0a, 0x11, 0x49, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x43, 0x0a,
	0x0c, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x4b, 0x0a, 0x10, 0x70, 0x61, 0x67, 0x65, 0x72, 0x64, 0x75, 0x74, 0x79, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x0f,
	0x70, 0x61, 0x67, 0x65, 0x72, 0x64, 0x75, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42,
	0x13, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x12,
	0x03, 0xf8, 0x42, 0x01, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e,
	0x61, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2d, 0x73, 0x64,
	0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61,
	0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_prodvana_workflow_integration_config_proto_rawDescOnce sync.Once
	file_prodvana_workflow_integration_config_proto_rawDescData = file_prodvana_workflow_integration_config_proto_rawDesc
)

func file_prodvana_workflow_integration_config_proto_rawDescGZIP() []byte {
	file_prodvana_workflow_integration_config_proto_rawDescOnce.Do(func() {
		file_prodvana_workflow_integration_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_prodvana_workflow_integration_config_proto_rawDescData)
	})
	return file_prodvana_workflow_integration_config_proto_rawDescData
}

var file_prodvana_workflow_integration_config_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_prodvana_workflow_integration_config_proto_goTypes = []interface{}{
	(*AlertingConfig)(nil),              // 0: prodvana.workflow.AlertingConfig
	(*AnnotationsConfig)(nil),           // 1: prodvana.workflow.AnnotationsConfig
	(*TokenConfig)(nil),                 // 2: prodvana.workflow.TokenConfig
	(*IntegrationConfig)(nil),           // 3: prodvana.workflow.IntegrationConfig
	(*AlertingConfig_PagerDuty)(nil),    // 4: prodvana.workflow.AlertingConfig.PagerDuty
	(*AnnotationsConfig_Honeycomb)(nil), // 5: prodvana.workflow.AnnotationsConfig.Honeycomb
}
var file_prodvana_workflow_integration_config_proto_depIdxs = []int32{
	4, // 0: prodvana.workflow.AlertingConfig.pagerduty:type_name -> prodvana.workflow.AlertingConfig.PagerDuty
	5, // 1: prodvana.workflow.AnnotationsConfig.honeycomb:type_name -> prodvana.workflow.AnnotationsConfig.Honeycomb
	2, // 2: prodvana.workflow.IntegrationConfig.slack_config:type_name -> prodvana.workflow.TokenConfig
	2, // 3: prodvana.workflow.IntegrationConfig.pagerduty_config:type_name -> prodvana.workflow.TokenConfig
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_prodvana_workflow_integration_config_proto_init() }
func file_prodvana_workflow_integration_config_proto_init() {
	if File_prodvana_workflow_integration_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prodvana_workflow_integration_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertingConfig); i {
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
		file_prodvana_workflow_integration_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_prodvana_workflow_integration_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenConfig); i {
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
		file_prodvana_workflow_integration_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntegrationConfig); i {
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
		file_prodvana_workflow_integration_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertingConfig_PagerDuty); i {
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
		file_prodvana_workflow_integration_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnotationsConfig_Honeycomb); i {
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
	file_prodvana_workflow_integration_config_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*IntegrationConfig_SlackConfig)(nil),
		(*IntegrationConfig_PagerdutyConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_prodvana_workflow_integration_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prodvana_workflow_integration_config_proto_goTypes,
		DependencyIndexes: file_prodvana_workflow_integration_config_proto_depIdxs,
		MessageInfos:      file_prodvana_workflow_integration_config_proto_msgTypes,
	}.Build()
	File_prodvana_workflow_integration_config_proto = out.File
	file_prodvana_workflow_integration_config_proto_rawDesc = nil
	file_prodvana_workflow_integration_config_proto_goTypes = nil
	file_prodvana_workflow_integration_config_proto_depIdxs = nil
}
