// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.10
// source: prodvana/common_config/task.proto

package common_config

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	volumes "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/volumes"
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

type TaskLifecycle int32

const (
	TaskLifecycle_UNKNOWN_TASK_LIFECYCLE TaskLifecycle = 0
	TaskLifecycle_CONVERGENCE_START      TaskLifecycle = 1
	// Runs before everything else. May wait for any dependent RCs to be stable.
	TaskLifecycle_PRE_APPROVAL TaskLifecycle = 2
	// Approval tasks, manual or automated. Must wait for all PRE_APPROVAL_TASK to be satisfied.
	TaskLifecycle_APPROVAL TaskLifecycle = 3
	// Runs after approval, just before service push.
	TaskLifecycle_POST_APPROVAL TaskLifecycle = 4
	// Runs as long as service push is in progress.
	TaskLifecycle_DEPLOYMENT TaskLifecycle = 5
	// Runs after service push succeeds (pods are replaced and healthy, ...), before declaring the service CONVERGED.
	TaskLifecycle_POST_DEPLOYMENT TaskLifecycle = 6
)

// Enum value maps for TaskLifecycle.
var (
	TaskLifecycle_name = map[int32]string{
		0: "UNKNOWN_TASK_LIFECYCLE",
		1: "CONVERGENCE_START",
		2: "PRE_APPROVAL",
		3: "APPROVAL",
		4: "POST_APPROVAL",
		5: "DEPLOYMENT",
		6: "POST_DEPLOYMENT",
	}
	TaskLifecycle_value = map[string]int32{
		"UNKNOWN_TASK_LIFECYCLE": 0,
		"CONVERGENCE_START":      1,
		"PRE_APPROVAL":           2,
		"APPROVAL":               3,
		"POST_APPROVAL":          4,
		"DEPLOYMENT":             5,
		"POST_DEPLOYMENT":        6,
	}
)

func (x TaskLifecycle) Enum() *TaskLifecycle {
	p := new(TaskLifecycle)
	*p = x
	return p
}

func (x TaskLifecycle) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskLifecycle) Descriptor() protoreflect.EnumDescriptor {
	return file_prodvana_common_config_task_proto_enumTypes[0].Descriptor()
}

func (TaskLifecycle) Type() protoreflect.EnumType {
	return &file_prodvana_common_config_task_proto_enumTypes[0]
}

func (x TaskLifecycle) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskLifecycle.Descriptor instead.
func (TaskLifecycle) EnumDescriptor() ([]byte, []int) {
	return file_prodvana_common_config_task_proto_rawDescGZIP(), []int{0}
}

type TaskConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Program *ProgramConfig    `protobuf:"bytes,1,opt,name=program,proto3" json:"program,omitempty"`
	Volumes []*volumes.Volume `protobuf:"bytes,3,rep,name=volumes,proto3" json:"volumes,omitempty"`
	// TTL after the job completes, defaults to 1 hour.
	Ttl *durationpb.Duration `protobuf:"bytes,4,opt,name=ttl,proto3" json:"ttl,omitempty"`
	// If not set, the task will not be retried once it starts executing once.
	RetryConfig *RetryConfig `protobuf:"bytes,2,opt,name=retry_config,json=retryConfig,proto3" json:"retry_config,omitempty"`
	// optional runtime specific configuration
	//
	// Types that are assignable to RuntimeSpecific:
	//
	//	*TaskConfig_K8S
	RuntimeSpecific isTaskConfig_RuntimeSpecific `protobuf_oneof:"runtime_specific"`
}

func (x *TaskConfig) Reset() {
	*x = TaskConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_common_config_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskConfig) ProtoMessage() {}

func (x *TaskConfig) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_common_config_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskConfig.ProtoReflect.Descriptor instead.
func (*TaskConfig) Descriptor() ([]byte, []int) {
	return file_prodvana_common_config_task_proto_rawDescGZIP(), []int{0}
}

func (x *TaskConfig) GetProgram() *ProgramConfig {
	if x != nil {
		return x.Program
	}
	return nil
}

func (x *TaskConfig) GetVolumes() []*volumes.Volume {
	if x != nil {
		return x.Volumes
	}
	return nil
}

func (x *TaskConfig) GetTtl() *durationpb.Duration {
	if x != nil {
		return x.Ttl
	}
	return nil
}

func (x *TaskConfig) GetRetryConfig() *RetryConfig {
	if x != nil {
		return x.RetryConfig
	}
	return nil
}

func (m *TaskConfig) GetRuntimeSpecific() isTaskConfig_RuntimeSpecific {
	if m != nil {
		return m.RuntimeSpecific
	}
	return nil
}

func (x *TaskConfig) GetK8S() *TaskConfig_KubernetesConfig {
	if x, ok := x.GetRuntimeSpecific().(*TaskConfig_K8S); ok {
		return x.K8S
	}
	return nil
}

type isTaskConfig_RuntimeSpecific interface {
	isTaskConfig_RuntimeSpecific()
}

type TaskConfig_K8S struct {
	K8S *TaskConfig_KubernetesConfig `protobuf:"bytes,5,opt,name=k8s,proto3,oneof"`
}

func (*TaskConfig_K8S) isTaskConfig_RuntimeSpecific() {}

type TaskConfig_KubernetesConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceAccount string `protobuf:"bytes,1,opt,name=service_account,json=serviceAccount,proto3" json:"service_account,omitempty"`
}

func (x *TaskConfig_KubernetesConfig) Reset() {
	*x = TaskConfig_KubernetesConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_common_config_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskConfig_KubernetesConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskConfig_KubernetesConfig) ProtoMessage() {}

func (x *TaskConfig_KubernetesConfig) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_common_config_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskConfig_KubernetesConfig.ProtoReflect.Descriptor instead.
func (*TaskConfig_KubernetesConfig) Descriptor() ([]byte, []int) {
	return file_prodvana_common_config_task_proto_rawDescGZIP(), []int{0, 0}
}

func (x *TaskConfig_KubernetesConfig) GetServiceAccount() string {
	if x != nil {
		return x.ServiceAccount
	}
	return ""
}

var File_prodvana_common_config_task_proto protoreflect.FileDescriptor

var file_prodvana_common_config_task_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x16, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x72, 0x65, 0x74,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61,
	0x6e, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73,
	0x2f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2,
	0x03, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x49, 0x0a,
	0x07, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x41, 0x0a, 0x07, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x76, 0x61, 0x6e, 0x61, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x2e, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x42, 0x0d, 0xfa, 0x42, 0x0a, 0x92, 0x01, 0x07, 0x22, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x07, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x03, 0x74,
	0x74, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x46, 0x0a, 0x0c, 0x72, 0x65, 0x74, 0x72,
	0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x0b, 0x72, 0x65, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x47, 0x0a, 0x03, 0x6b, 0x38, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x48, 0x00, 0x52, 0x03, 0x6b, 0x38, 0x73, 0x1a, 0x44, 0x0a, 0x10, 0x4b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x30, 0x0a,
	0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x12, 0x0a, 0x10, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x63, 0x2a, 0x9a, 0x01, 0x0a, 0x0d, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x66, 0x65,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x5f, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x4c, 0x49, 0x46, 0x45, 0x43, 0x59, 0x43, 0x4c, 0x45, 0x10,
	0x00, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x47, 0x45, 0x4e, 0x43, 0x45,
	0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x52, 0x45, 0x5f,
	0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x50,
	0x50, 0x52, 0x4f, 0x56, 0x41, 0x4c, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x4f, 0x53, 0x54,
	0x5f, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x41, 0x4c, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x44,
	0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x50,
	0x4f, 0x53, 0x54, 0x5f, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x06,
	0x42, 0x52, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61,
	0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76,
	0x61, 0x6e, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prodvana_common_config_task_proto_rawDescOnce sync.Once
	file_prodvana_common_config_task_proto_rawDescData = file_prodvana_common_config_task_proto_rawDesc
)

func file_prodvana_common_config_task_proto_rawDescGZIP() []byte {
	file_prodvana_common_config_task_proto_rawDescOnce.Do(func() {
		file_prodvana_common_config_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_prodvana_common_config_task_proto_rawDescData)
	})
	return file_prodvana_common_config_task_proto_rawDescData
}

var file_prodvana_common_config_task_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_prodvana_common_config_task_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_prodvana_common_config_task_proto_goTypes = []interface{}{
	(TaskLifecycle)(0),                  // 0: prodvana.common_config.TaskLifecycle
	(*TaskConfig)(nil),                  // 1: prodvana.common_config.TaskConfig
	(*TaskConfig_KubernetesConfig)(nil), // 2: prodvana.common_config.TaskConfig.KubernetesConfig
	(*ProgramConfig)(nil),               // 3: prodvana.common_config.ProgramConfig
	(*volumes.Volume)(nil),              // 4: prodvana.volumes.Volume
	(*durationpb.Duration)(nil),         // 5: google.protobuf.Duration
	(*RetryConfig)(nil),                 // 6: prodvana.common_config.RetryConfig
}
var file_prodvana_common_config_task_proto_depIdxs = []int32{
	3, // 0: prodvana.common_config.TaskConfig.program:type_name -> prodvana.common_config.ProgramConfig
	4, // 1: prodvana.common_config.TaskConfig.volumes:type_name -> prodvana.volumes.Volume
	5, // 2: prodvana.common_config.TaskConfig.ttl:type_name -> google.protobuf.Duration
	6, // 3: prodvana.common_config.TaskConfig.retry_config:type_name -> prodvana.common_config.RetryConfig
	2, // 4: prodvana.common_config.TaskConfig.k8s:type_name -> prodvana.common_config.TaskConfig.KubernetesConfig
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_prodvana_common_config_task_proto_init() }
func file_prodvana_common_config_task_proto_init() {
	if File_prodvana_common_config_task_proto != nil {
		return
	}
	file_prodvana_common_config_retry_proto_init()
	file_prodvana_common_config_program_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_prodvana_common_config_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskConfig); i {
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
		file_prodvana_common_config_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskConfig_KubernetesConfig); i {
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
	file_prodvana_common_config_task_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TaskConfig_K8S)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_prodvana_common_config_task_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prodvana_common_config_task_proto_goTypes,
		DependencyIndexes: file_prodvana_common_config_task_proto_depIdxs,
		EnumInfos:         file_prodvana_common_config_task_proto_enumTypes,
		MessageInfos:      file_prodvana_common_config_task_proto_msgTypes,
	}.Build()
	File_prodvana_common_config_task_proto = out.File
	file_prodvana_common_config_task_proto_rawDesc = nil
	file_prodvana_common_config_task_proto_goTypes = nil
	file_prodvana_common_config_task_proto_depIdxs = nil
}
