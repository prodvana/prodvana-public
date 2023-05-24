// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.10
// source: prodvana/delivery_extension/config.proto

package delivery_extension

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	common_config "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config"
	runtimes "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/runtimes"
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

type DeliveryExtensionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are assignable to ExecConfig:
	//
	//	*DeliveryExtensionConfig_TaskConfig
	//	*DeliveryExtensionConfig_KubernetesConfig
	ExecConfig isDeliveryExtensionConfig_ExecConfig `protobuf_oneof:"exec_config"`
	Parameters []*common_config.ParameterDefinition `protobuf:"bytes,4,rep,name=parameters,proto3" json:"parameters,omitempty"`
}

func (x *DeliveryExtensionConfig) Reset() {
	*x = DeliveryExtensionConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_delivery_extension_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryExtensionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryExtensionConfig) ProtoMessage() {}

func (x *DeliveryExtensionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_delivery_extension_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryExtensionConfig.ProtoReflect.Descriptor instead.
func (*DeliveryExtensionConfig) Descriptor() ([]byte, []int) {
	return file_prodvana_delivery_extension_config_proto_rawDescGZIP(), []int{0}
}

func (x *DeliveryExtensionConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *DeliveryExtensionConfig) GetExecConfig() isDeliveryExtensionConfig_ExecConfig {
	if m != nil {
		return m.ExecConfig
	}
	return nil
}

func (x *DeliveryExtensionConfig) GetTaskConfig() *common_config.TaskConfig {
	if x, ok := x.GetExecConfig().(*DeliveryExtensionConfig_TaskConfig); ok {
		return x.TaskConfig
	}
	return nil
}

func (x *DeliveryExtensionConfig) GetKubernetesConfig() *common_config.KubernetesConfig {
	if x, ok := x.GetExecConfig().(*DeliveryExtensionConfig_KubernetesConfig); ok {
		return x.KubernetesConfig
	}
	return nil
}

func (x *DeliveryExtensionConfig) GetParameters() []*common_config.ParameterDefinition {
	if x != nil {
		return x.Parameters
	}
	return nil
}

type isDeliveryExtensionConfig_ExecConfig interface {
	isDeliveryExtensionConfig_ExecConfig()
}

type DeliveryExtensionConfig_TaskConfig struct {
	TaskConfig *common_config.TaskConfig `protobuf:"bytes,2,opt,name=task_config,json=taskConfig,proto3,oneof"`
}

type DeliveryExtensionConfig_KubernetesConfig struct {
	KubernetesConfig *common_config.KubernetesConfig `protobuf:"bytes,3,opt,name=kubernetes_config,json=kubernetesConfig,proto3,oneof"`
}

func (*DeliveryExtensionConfig_TaskConfig) isDeliveryExtensionConfig_ExecConfig() {}

func (*DeliveryExtensionConfig_KubernetesConfig) isDeliveryExtensionConfig_ExecConfig() {}

type DeliveryExtensionInstanceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string                          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Parameters []*common_config.ParameterValue `protobuf:"bytes,3,rep,name=parameters,proto3" json:"parameters,omitempty"`
}

func (x *DeliveryExtensionInstanceConfig) Reset() {
	*x = DeliveryExtensionInstanceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_delivery_extension_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryExtensionInstanceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryExtensionInstanceConfig) ProtoMessage() {}

func (x *DeliveryExtensionInstanceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_delivery_extension_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryExtensionInstanceConfig.ProtoReflect.Descriptor instead.
func (*DeliveryExtensionInstanceConfig) Descriptor() ([]byte, []int) {
	return file_prodvana_delivery_extension_config_proto_rawDescGZIP(), []int{1}
}

func (x *DeliveryExtensionInstanceConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeliveryExtensionInstanceConfig) GetParameters() []*common_config.ParameterValue {
	if x != nil {
		return x.Parameters
	}
	return nil
}

type CompiledDeliveryExtensionInstanceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Definition       *DeliveryExtensionConfig         `protobuf:"bytes,1,opt,name=definition,proto3" json:"definition,omitempty"`
	RuntimeExecution *runtimes.RuntimeExecutionConfig `protobuf:"bytes,2,opt,name=runtime_execution,json=runtimeExecution,proto3" json:"runtime_execution,omitempty"`
	// The compiled environment for this attachment's context, e.g.  Release Channel.
	Env map[string]*common_config.EnvValue `protobuf:"bytes,3,rep,name=env,proto3" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// compiled parameter values
	ParameterValues []*common_config.ParameterValue `protobuf:"bytes,4,rep,name=parameter_values,json=parameterValues,proto3" json:"parameter_values,omitempty"`
}

func (x *CompiledDeliveryExtensionInstanceConfig) Reset() {
	*x = CompiledDeliveryExtensionInstanceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_delivery_extension_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompiledDeliveryExtensionInstanceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompiledDeliveryExtensionInstanceConfig) ProtoMessage() {}

func (x *CompiledDeliveryExtensionInstanceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_delivery_extension_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompiledDeliveryExtensionInstanceConfig.ProtoReflect.Descriptor instead.
func (*CompiledDeliveryExtensionInstanceConfig) Descriptor() ([]byte, []int) {
	return file_prodvana_delivery_extension_config_proto_rawDescGZIP(), []int{2}
}

func (x *CompiledDeliveryExtensionInstanceConfig) GetDefinition() *DeliveryExtensionConfig {
	if x != nil {
		return x.Definition
	}
	return nil
}

func (x *CompiledDeliveryExtensionInstanceConfig) GetRuntimeExecution() *runtimes.RuntimeExecutionConfig {
	if x != nil {
		return x.RuntimeExecution
	}
	return nil
}

func (x *CompiledDeliveryExtensionInstanceConfig) GetEnv() map[string]*common_config.EnvValue {
	if x != nil {
		return x.Env
	}
	return nil
}

func (x *CompiledDeliveryExtensionInstanceConfig) GetParameterValues() []*common_config.ParameterValue {
	if x != nil {
		return x.ParameterValues
	}
	return nil
}

var File_prodvana_delivery_extension_config_proto protoreflect.FileDescriptor

var file_prodvana_delivery_extension_config_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x70, 0x72, 0x6f, 0x64,
	0x76, 0x61, 0x6e, 0x61, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x21, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e, 0x76, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27,
	0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea, 0x02, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x3f, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x2b, 0xfa, 0x42, 0x28, 0x72, 0x26, 0x10, 0x01, 0x18, 0x3f, 0x32, 0x20, 0x5e, 0x5b,
	0x61, 0x2d, 0x7a, 0x5d, 0x28, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2d, 0x5d, 0x2a, 0x5b,
	0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x7b, 0x30, 0x2c, 0x31, 0x7d, 0x24, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52,
	0x0a, 0x74, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x57, 0x0a, 0x11, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e,
	0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x48, 0x00, 0x52, 0x10, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x5a, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76,
	0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0d, 0xfa, 0x42, 0x0a, 0x92, 0x01, 0x07, 0x22, 0x05, 0x8a,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x42, 0x12, 0x0a, 0x0b, 0x65, 0x78, 0x65, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x03, 0xf8, 0x42, 0x01, 0x22, 0xaa, 0x01, 0x0a, 0x1f, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3f, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0xfa, 0x42, 0x28, 0x72, 0x26, 0x10, 0x01, 0x18,
	0x3f, 0x32, 0x20, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x28, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d,
	0x39, 0x2d, 0x5d, 0x2a, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x7b, 0x30, 0x2c,
	0x31, 0x7d, 0x24, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x0a, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x22, 0x8f, 0x04, 0x0a, 0x27, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x54, 0x0a,
	0x0a, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x64, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0a, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x56, 0x0a, 0x11, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x10, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x88, 0x01, 0x0a, 0x03,
	0x65, 0x6e, 0x76, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4d, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x76, 0x61, 0x6e, 0x61, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x45, 0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x27, 0xfa, 0x42, 0x24, 0x9a, 0x01, 0x21,
	0x18, 0x01, 0x22, 0x1d, 0x72, 0x1b, 0x32, 0x19, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a,
	0x5f, 0x5d, 0x2b, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x2a,
	0x24, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x51, 0x0a, 0x10, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x58, 0x0a, 0x08, 0x45, 0x6e, 0x76,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x36, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e,
	0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x45, 0x6e, 0x76, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x42, 0x57, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76,
	0x61, 0x6e, 0x61, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prodvana_delivery_extension_config_proto_rawDescOnce sync.Once
	file_prodvana_delivery_extension_config_proto_rawDescData = file_prodvana_delivery_extension_config_proto_rawDesc
)

func file_prodvana_delivery_extension_config_proto_rawDescGZIP() []byte {
	file_prodvana_delivery_extension_config_proto_rawDescOnce.Do(func() {
		file_prodvana_delivery_extension_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_prodvana_delivery_extension_config_proto_rawDescData)
	})
	return file_prodvana_delivery_extension_config_proto_rawDescData
}

var file_prodvana_delivery_extension_config_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_prodvana_delivery_extension_config_proto_goTypes = []interface{}{
	(*DeliveryExtensionConfig)(nil),                 // 0: prodvana.delivery_extension.DeliveryExtensionConfig
	(*DeliveryExtensionInstanceConfig)(nil),         // 1: prodvana.delivery_extension.DeliveryExtensionInstanceConfig
	(*CompiledDeliveryExtensionInstanceConfig)(nil), // 2: prodvana.delivery_extension.CompiledDeliveryExtensionInstanceConfig
	nil,                                    // 3: prodvana.delivery_extension.CompiledDeliveryExtensionInstanceConfig.EnvEntry
	(*common_config.TaskConfig)(nil),       // 4: prodvana.common_config.TaskConfig
	(*common_config.KubernetesConfig)(nil), // 5: prodvana.common_config.KubernetesConfig
	(*common_config.ParameterDefinition)(nil), // 6: prodvana.common_config.ParameterDefinition
	(*common_config.ParameterValue)(nil),      // 7: prodvana.common_config.ParameterValue
	(*runtimes.RuntimeExecutionConfig)(nil),   // 8: prodvana.runtimes.RuntimeExecutionConfig
	(*common_config.EnvValue)(nil),            // 9: prodvana.common_config.EnvValue
}
var file_prodvana_delivery_extension_config_proto_depIdxs = []int32{
	4, // 0: prodvana.delivery_extension.DeliveryExtensionConfig.task_config:type_name -> prodvana.common_config.TaskConfig
	5, // 1: prodvana.delivery_extension.DeliveryExtensionConfig.kubernetes_config:type_name -> prodvana.common_config.KubernetesConfig
	6, // 2: prodvana.delivery_extension.DeliveryExtensionConfig.parameters:type_name -> prodvana.common_config.ParameterDefinition
	7, // 3: prodvana.delivery_extension.DeliveryExtensionInstanceConfig.parameters:type_name -> prodvana.common_config.ParameterValue
	0, // 4: prodvana.delivery_extension.CompiledDeliveryExtensionInstanceConfig.definition:type_name -> prodvana.delivery_extension.DeliveryExtensionConfig
	8, // 5: prodvana.delivery_extension.CompiledDeliveryExtensionInstanceConfig.runtime_execution:type_name -> prodvana.runtimes.RuntimeExecutionConfig
	3, // 6: prodvana.delivery_extension.CompiledDeliveryExtensionInstanceConfig.env:type_name -> prodvana.delivery_extension.CompiledDeliveryExtensionInstanceConfig.EnvEntry
	7, // 7: prodvana.delivery_extension.CompiledDeliveryExtensionInstanceConfig.parameter_values:type_name -> prodvana.common_config.ParameterValue
	9, // 8: prodvana.delivery_extension.CompiledDeliveryExtensionInstanceConfig.EnvEntry.value:type_name -> prodvana.common_config.EnvValue
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_prodvana_delivery_extension_config_proto_init() }
func file_prodvana_delivery_extension_config_proto_init() {
	if File_prodvana_delivery_extension_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prodvana_delivery_extension_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryExtensionConfig); i {
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
		file_prodvana_delivery_extension_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryExtensionInstanceConfig); i {
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
		file_prodvana_delivery_extension_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompiledDeliveryExtensionInstanceConfig); i {
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
	file_prodvana_delivery_extension_config_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*DeliveryExtensionConfig_TaskConfig)(nil),
		(*DeliveryExtensionConfig_KubernetesConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_prodvana_delivery_extension_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prodvana_delivery_extension_config_proto_goTypes,
		DependencyIndexes: file_prodvana_delivery_extension_config_proto_depIdxs,
		MessageInfos:      file_prodvana_delivery_extension_config_proto_msgTypes,
	}.Build()
	File_prodvana_delivery_extension_config_proto = out.File
	file_prodvana_delivery_extension_config_proto_rawDesc = nil
	file_prodvana_delivery_extension_config_proto_goTypes = nil
	file_prodvana_delivery_extension_config_proto_depIdxs = nil
}
