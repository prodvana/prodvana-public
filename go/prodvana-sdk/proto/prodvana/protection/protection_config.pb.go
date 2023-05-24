// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.10
// source: prodvana/protection/protection_config.proto

package protection

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	common_config "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config"
	runtimes "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/runtimes"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProtectionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are assignable to ExecConfig:
	//
	//	*ProtectionConfig_TaskConfig
	//	*ProtectionConfig_KubernetesConfig
	ExecConfig isProtectionConfig_ExecConfig `protobuf_oneof:"exec_config"`
	// customize intervals instead of using Prodvana default
	// how often to run check even if it succeeds
	PollInterval *durationpb.Duration `protobuf:"bytes,4,opt,name=poll_interval,json=pollInterval,proto3" json:"poll_interval,omitempty"`
	// how long after a run has started to try another run if it has not completed yet
	Timeout    *durationpb.Duration                 `protobuf:"bytes,5,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Parameters []*common_config.ParameterDefinition `protobuf:"bytes,6,rep,name=parameters,proto3" json:"parameters,omitempty"`
}

func (x *ProtectionConfig) Reset() {
	*x = ProtectionConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_protection_protection_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtectionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtectionConfig) ProtoMessage() {}

func (x *ProtectionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_protection_protection_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtectionConfig.ProtoReflect.Descriptor instead.
func (*ProtectionConfig) Descriptor() ([]byte, []int) {
	return file_prodvana_protection_protection_config_proto_rawDescGZIP(), []int{0}
}

func (x *ProtectionConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *ProtectionConfig) GetExecConfig() isProtectionConfig_ExecConfig {
	if m != nil {
		return m.ExecConfig
	}
	return nil
}

func (x *ProtectionConfig) GetTaskConfig() *common_config.TaskConfig {
	if x, ok := x.GetExecConfig().(*ProtectionConfig_TaskConfig); ok {
		return x.TaskConfig
	}
	return nil
}

func (x *ProtectionConfig) GetKubernetesConfig() *common_config.KubernetesConfig {
	if x, ok := x.GetExecConfig().(*ProtectionConfig_KubernetesConfig); ok {
		return x.KubernetesConfig
	}
	return nil
}

func (x *ProtectionConfig) GetPollInterval() *durationpb.Duration {
	if x != nil {
		return x.PollInterval
	}
	return nil
}

func (x *ProtectionConfig) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *ProtectionConfig) GetParameters() []*common_config.ParameterDefinition {
	if x != nil {
		return x.Parameters
	}
	return nil
}

type isProtectionConfig_ExecConfig interface {
	isProtectionConfig_ExecConfig()
}

type ProtectionConfig_TaskConfig struct {
	// Inline task config with retry, template support.
	TaskConfig *common_config.TaskConfig `protobuf:"bytes,2,opt,name=task_config,json=taskConfig,proto3,oneof"`
}

type ProtectionConfig_KubernetesConfig struct {
	KubernetesConfig *common_config.KubernetesConfig `protobuf:"bytes,3,opt,name=kubernetes_config,json=kubernetesConfig,proto3,oneof"`
}

func (*ProtectionConfig_TaskConfig) isProtectionConfig_ExecConfig() {}

func (*ProtectionConfig_KubernetesConfig) isProtectionConfig_ExecConfig() {}

type CompiledProtectionAttachmentConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *ProtectionConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	// Protection source - where did this protection get attached from (service/app/org/...)?
	Attachment       *ProtectionAttachment            `protobuf:"bytes,2,opt,name=attachment,proto3" json:"attachment,omitempty"`
	RuntimeExecution *runtimes.RuntimeExecutionConfig `protobuf:"bytes,3,opt,name=runtime_execution,json=runtimeExecution,proto3" json:"runtime_execution,omitempty"`
	// The compiled environment for this attachment's context, e.g.  Release Channel.
	Env map[string]*common_config.EnvValue `protobuf:"bytes,4,rep,name=env,proto3" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// compiled parameter values
	ParameterValues []*common_config.ParameterValue `protobuf:"bytes,5,rep,name=parameter_values,json=parameterValues,proto3" json:"parameter_values,omitempty"`
}

func (x *CompiledProtectionAttachmentConfig) Reset() {
	*x = CompiledProtectionAttachmentConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_protection_protection_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompiledProtectionAttachmentConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompiledProtectionAttachmentConfig) ProtoMessage() {}

func (x *CompiledProtectionAttachmentConfig) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_protection_protection_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompiledProtectionAttachmentConfig.ProtoReflect.Descriptor instead.
func (*CompiledProtectionAttachmentConfig) Descriptor() ([]byte, []int) {
	return file_prodvana_protection_protection_config_proto_rawDescGZIP(), []int{1}
}

func (x *CompiledProtectionAttachmentConfig) GetConfig() *ProtectionConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *CompiledProtectionAttachmentConfig) GetAttachment() *ProtectionAttachment {
	if x != nil {
		return x.Attachment
	}
	return nil
}

func (x *CompiledProtectionAttachmentConfig) GetRuntimeExecution() *runtimes.RuntimeExecutionConfig {
	if x != nil {
		return x.RuntimeExecution
	}
	return nil
}

func (x *CompiledProtectionAttachmentConfig) GetEnv() map[string]*common_config.EnvValue {
	if x != nil {
		return x.Env
	}
	return nil
}

func (x *CompiledProtectionAttachmentConfig) GetParameterValues() []*common_config.ParameterValue {
	if x != nil {
		return x.ParameterValues
	}
	return nil
}

type ServiceInstanceAttachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service        string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	ReleaseChannel string `protobuf:"bytes,2,opt,name=release_channel,json=releaseChannel,proto3" json:"release_channel,omitempty"`
	Application    string `protobuf:"bytes,3,opt,name=application,proto3" json:"application,omitempty"`
}

func (x *ServiceInstanceAttachment) Reset() {
	*x = ServiceInstanceAttachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_protection_protection_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceInstanceAttachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceInstanceAttachment) ProtoMessage() {}

func (x *ServiceInstanceAttachment) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_protection_protection_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceInstanceAttachment.ProtoReflect.Descriptor instead.
func (*ServiceInstanceAttachment) Descriptor() ([]byte, []int) {
	return file_prodvana_protection_protection_config_proto_rawDescGZIP(), []int{2}
}

func (x *ServiceInstanceAttachment) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *ServiceInstanceAttachment) GetReleaseChannel() string {
	if x != nil {
		return x.ReleaseChannel
	}
	return ""
}

func (x *ServiceInstanceAttachment) GetApplication() string {
	if x != nil {
		return x.Application
	}
	return ""
}

type ReleaseChannelAttachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Application    string `protobuf:"bytes,1,opt,name=application,proto3" json:"application,omitempty"`
	ReleaseChannel string `protobuf:"bytes,2,opt,name=release_channel,json=releaseChannel,proto3" json:"release_channel,omitempty"`
}

func (x *ReleaseChannelAttachment) Reset() {
	*x = ReleaseChannelAttachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_protection_protection_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseChannelAttachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseChannelAttachment) ProtoMessage() {}

func (x *ReleaseChannelAttachment) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_protection_protection_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseChannelAttachment.ProtoReflect.Descriptor instead.
func (*ReleaseChannelAttachment) Descriptor() ([]byte, []int) {
	return file_prodvana_protection_protection_config_proto_rawDescGZIP(), []int{3}
}

func (x *ReleaseChannelAttachment) GetApplication() string {
	if x != nil {
		return x.Application
	}
	return ""
}

func (x *ReleaseChannelAttachment) GetReleaseChannel() string {
	if x != nil {
		return x.ReleaseChannel
	}
	return ""
}

type ProtectionAttachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Attachment:
	//
	//	*ProtectionAttachment_ServiceInstance
	//	*ProtectionAttachment_ReleaseChannel
	Attachment isProtectionAttachment_Attachment `protobuf_oneof:"attachment"`
}

func (x *ProtectionAttachment) Reset() {
	*x = ProtectionAttachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_protection_protection_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtectionAttachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtectionAttachment) ProtoMessage() {}

func (x *ProtectionAttachment) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_protection_protection_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtectionAttachment.ProtoReflect.Descriptor instead.
func (*ProtectionAttachment) Descriptor() ([]byte, []int) {
	return file_prodvana_protection_protection_config_proto_rawDescGZIP(), []int{4}
}

func (m *ProtectionAttachment) GetAttachment() isProtectionAttachment_Attachment {
	if m != nil {
		return m.Attachment
	}
	return nil
}

func (x *ProtectionAttachment) GetServiceInstance() *ServiceInstanceAttachment {
	if x, ok := x.GetAttachment().(*ProtectionAttachment_ServiceInstance); ok {
		return x.ServiceInstance
	}
	return nil
}

func (x *ProtectionAttachment) GetReleaseChannel() *ReleaseChannelAttachment {
	if x, ok := x.GetAttachment().(*ProtectionAttachment_ReleaseChannel); ok {
		return x.ReleaseChannel
	}
	return nil
}

type isProtectionAttachment_Attachment interface {
	isProtectionAttachment_Attachment()
}

type ProtectionAttachment_ServiceInstance struct {
	ServiceInstance *ServiceInstanceAttachment `protobuf:"bytes,1,opt,name=service_instance,json=serviceInstance,proto3,oneof"`
}

type ProtectionAttachment_ReleaseChannel struct {
	ReleaseChannel *ReleaseChannelAttachment `protobuf:"bytes,2,opt,name=release_channel,json=releaseChannel,proto3,oneof"`
}

func (*ProtectionAttachment_ServiceInstance) isProtectionAttachment_Attachment() {}

func (*ProtectionAttachment_ReleaseChannel) isProtectionAttachment_Attachment() {}

var File_prodvana_protection_protection_config_proto protoreflect.FileDescriptor

var file_prodvana_protection_protection_config_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x70,
	0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x72, 0x6f,
	0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x27, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x70, 0x72, 0x6f, 0x64, 0x76,
	0x61, 0x6e, 0x61, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x2f, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xd8, 0x03, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3f, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0xfa, 0x42, 0x28, 0x72, 0x26, 0x10, 0x01, 0x18, 0x3f,
	0x32, 0x20, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x28, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39,
	0x2d, 0x5d, 0x2a, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x7b, 0x30, 0x2c, 0x31,
	0x7d, 0x24, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x0b, 0x74, 0x61, 0x73, 0x6b,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x48, 0x00, 0x52, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x57, 0x0a, 0x11, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x10, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3e, 0x0a, 0x0d, 0x70, 0x6f, 0x6c, 0x6c,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x70, 0x6f, 0x6c, 0x6c,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x5a, 0x0a,
	0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0d,
	0xfa, 0x42, 0x0a, 0x92, 0x01, 0x07, 0x22, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x42, 0x12, 0x0a, 0x0b, 0x65, 0x78, 0x65,
	0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0xb0, 0x04,
	0x0a, 0x22, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x3d, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x49, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61,
	0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x56,
	0x0a, 0x11, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x76, 0x61, 0x6e, 0x61, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x2e, 0x52, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x10, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x7b, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c,
	0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x45, 0x6e, 0x76,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x27, 0xfa, 0x42, 0x24, 0x9a, 0x01, 0x21, 0x18, 0x01, 0x22,
	0x1d, 0x72, 0x1b, 0x32, 0x19, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x5f, 0x5d, 0x2b,
	0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x2a, 0x24, 0x52, 0x03,
	0x65, 0x6e, 0x76, 0x12, 0x51, 0x0a, 0x10, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x58, 0x0a, 0x08, 0x45, 0x6e, 0x76, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x36, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x45, 0x6e, 0x76,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x80, 0x01, 0x0a, 0x19, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x65, 0x0a, 0x18, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0xe0, 0x01, 0x0a, 0x14, 0x50,
	0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x5b, 0x0a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52,
	0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x58, 0x0a, 0x0f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x76, 0x61, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x41, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0e, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x42, 0x11, 0x0a, 0x0a, 0x61, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x42, 0x4f, 0x5a,
	0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x64,
	0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2d, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76,
	0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prodvana_protection_protection_config_proto_rawDescOnce sync.Once
	file_prodvana_protection_protection_config_proto_rawDescData = file_prodvana_protection_protection_config_proto_rawDesc
)

func file_prodvana_protection_protection_config_proto_rawDescGZIP() []byte {
	file_prodvana_protection_protection_config_proto_rawDescOnce.Do(func() {
		file_prodvana_protection_protection_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_prodvana_protection_protection_config_proto_rawDescData)
	})
	return file_prodvana_protection_protection_config_proto_rawDescData
}

var file_prodvana_protection_protection_config_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_prodvana_protection_protection_config_proto_goTypes = []interface{}{
	(*ProtectionConfig)(nil),                   // 0: prodvana.protection.ProtectionConfig
	(*CompiledProtectionAttachmentConfig)(nil), // 1: prodvana.protection.CompiledProtectionAttachmentConfig
	(*ServiceInstanceAttachment)(nil),          // 2: prodvana.protection.ServiceInstanceAttachment
	(*ReleaseChannelAttachment)(nil),           // 3: prodvana.protection.ReleaseChannelAttachment
	(*ProtectionAttachment)(nil),               // 4: prodvana.protection.ProtectionAttachment
	nil,                                        // 5: prodvana.protection.CompiledProtectionAttachmentConfig.EnvEntry
	(*common_config.TaskConfig)(nil),           // 6: prodvana.common_config.TaskConfig
	(*common_config.KubernetesConfig)(nil),     // 7: prodvana.common_config.KubernetesConfig
	(*durationpb.Duration)(nil),                // 8: google.protobuf.Duration
	(*common_config.ParameterDefinition)(nil),  // 9: prodvana.common_config.ParameterDefinition
	(*runtimes.RuntimeExecutionConfig)(nil),    // 10: prodvana.runtimes.RuntimeExecutionConfig
	(*common_config.ParameterValue)(nil),       // 11: prodvana.common_config.ParameterValue
	(*common_config.EnvValue)(nil),             // 12: prodvana.common_config.EnvValue
}
var file_prodvana_protection_protection_config_proto_depIdxs = []int32{
	6,  // 0: prodvana.protection.ProtectionConfig.task_config:type_name -> prodvana.common_config.TaskConfig
	7,  // 1: prodvana.protection.ProtectionConfig.kubernetes_config:type_name -> prodvana.common_config.KubernetesConfig
	8,  // 2: prodvana.protection.ProtectionConfig.poll_interval:type_name -> google.protobuf.Duration
	8,  // 3: prodvana.protection.ProtectionConfig.timeout:type_name -> google.protobuf.Duration
	9,  // 4: prodvana.protection.ProtectionConfig.parameters:type_name -> prodvana.common_config.ParameterDefinition
	0,  // 5: prodvana.protection.CompiledProtectionAttachmentConfig.config:type_name -> prodvana.protection.ProtectionConfig
	4,  // 6: prodvana.protection.CompiledProtectionAttachmentConfig.attachment:type_name -> prodvana.protection.ProtectionAttachment
	10, // 7: prodvana.protection.CompiledProtectionAttachmentConfig.runtime_execution:type_name -> prodvana.runtimes.RuntimeExecutionConfig
	5,  // 8: prodvana.protection.CompiledProtectionAttachmentConfig.env:type_name -> prodvana.protection.CompiledProtectionAttachmentConfig.EnvEntry
	11, // 9: prodvana.protection.CompiledProtectionAttachmentConfig.parameter_values:type_name -> prodvana.common_config.ParameterValue
	2,  // 10: prodvana.protection.ProtectionAttachment.service_instance:type_name -> prodvana.protection.ServiceInstanceAttachment
	3,  // 11: prodvana.protection.ProtectionAttachment.release_channel:type_name -> prodvana.protection.ReleaseChannelAttachment
	12, // 12: prodvana.protection.CompiledProtectionAttachmentConfig.EnvEntry.value:type_name -> prodvana.common_config.EnvValue
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_prodvana_protection_protection_config_proto_init() }
func file_prodvana_protection_protection_config_proto_init() {
	if File_prodvana_protection_protection_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prodvana_protection_protection_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtectionConfig); i {
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
		file_prodvana_protection_protection_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompiledProtectionAttachmentConfig); i {
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
		file_prodvana_protection_protection_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceInstanceAttachment); i {
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
		file_prodvana_protection_protection_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseChannelAttachment); i {
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
		file_prodvana_protection_protection_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtectionAttachment); i {
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
	file_prodvana_protection_protection_config_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ProtectionConfig_TaskConfig)(nil),
		(*ProtectionConfig_KubernetesConfig)(nil),
	}
	file_prodvana_protection_protection_config_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*ProtectionAttachment_ServiceInstance)(nil),
		(*ProtectionAttachment_ReleaseChannel)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_prodvana_protection_protection_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prodvana_protection_protection_config_proto_goTypes,
		DependencyIndexes: file_prodvana_protection_protection_config_proto_depIdxs,
		MessageInfos:      file_prodvana_protection_protection_config_proto_msgTypes,
	}.Build()
	File_prodvana_protection_protection_config_proto = out.File
	file_prodvana_protection_protection_config_proto_rawDesc = nil
	file_prodvana_protection_protection_config_proto_goTypes = nil
	file_prodvana_protection_protection_config_proto_depIdxs = nil
}
