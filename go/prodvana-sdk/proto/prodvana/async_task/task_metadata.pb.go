// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.10
// source: prodvana/async_task/task_metadata.proto

package async_task

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

type TaskStatus int32

const (
	TaskStatus_UNKNOWN TaskStatus = 0
	TaskStatus_PENDING TaskStatus = 1
	TaskStatus_RUNNING TaskStatus = 2
	TaskStatus_SUCCESS TaskStatus = 3
	TaskStatus_FAILED  TaskStatus = 4
)

// Enum value maps for TaskStatus.
var (
	TaskStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "PENDING",
		2: "RUNNING",
		3: "SUCCESS",
		4: "FAILED",
	}
	TaskStatus_value = map[string]int32{
		"UNKNOWN": 0,
		"PENDING": 1,
		"RUNNING": 2,
		"SUCCESS": 3,
		"FAILED":  4,
	}
)

func (x TaskStatus) Enum() *TaskStatus {
	p := new(TaskStatus)
	*p = x
	return p
}

func (x TaskStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_prodvana_async_task_task_metadata_proto_enumTypes[0].Descriptor()
}

func (TaskStatus) Type() protoreflect.EnumType {
	return &file_prodvana_async_task_task_metadata_proto_enumTypes[0]
}

func (x TaskStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskStatus.Descriptor instead.
func (TaskStatus) EnumDescriptor() ([]byte, []int) {
	return file_prodvana_async_task_task_metadata_proto_rawDescGZIP(), []int{0}
}

type TaskResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Will be set if task failed, otherwise may be empty.
	Log []byte `protobuf:"bytes,1,opt,name=log,proto3" json:"log,omitempty"`
}

func (x *TaskResult) Reset() {
	*x = TaskResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_async_task_task_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResult) ProtoMessage() {}

func (x *TaskResult) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_async_task_task_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResult.ProtoReflect.Descriptor instead.
func (*TaskResult) Descriptor() ([]byte, []int) {
	return file_prodvana_async_task_task_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *TaskResult) GetLog() []byte {
	if x != nil {
		return x.Log
	}
	return nil
}

var File_prodvana_async_task_task_metadata_proto protoreflect.FileDescriptor

var file_prodvana_async_task_task_metadata_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x61, 0x73, 0x79, 0x6e, 0x63,
	0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x76,
	0x61, 0x6e, 0x61, 0x2e, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x1e,
	0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x2a, 0x4c,
	0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e,
	0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e,
	0x47, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x03,
	0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x42, 0x4f, 0x5a, 0x4d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76,
	0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2d, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61,
	0x6e, 0x61, 0x2f, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prodvana_async_task_task_metadata_proto_rawDescOnce sync.Once
	file_prodvana_async_task_task_metadata_proto_rawDescData = file_prodvana_async_task_task_metadata_proto_rawDesc
)

func file_prodvana_async_task_task_metadata_proto_rawDescGZIP() []byte {
	file_prodvana_async_task_task_metadata_proto_rawDescOnce.Do(func() {
		file_prodvana_async_task_task_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_prodvana_async_task_task_metadata_proto_rawDescData)
	})
	return file_prodvana_async_task_task_metadata_proto_rawDescData
}

var file_prodvana_async_task_task_metadata_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_prodvana_async_task_task_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_prodvana_async_task_task_metadata_proto_goTypes = []interface{}{
	(TaskStatus)(0),    // 0: prodvana.async_task.TaskStatus
	(*TaskResult)(nil), // 1: prodvana.async_task.TaskResult
}
var file_prodvana_async_task_task_metadata_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_prodvana_async_task_task_metadata_proto_init() }
func file_prodvana_async_task_task_metadata_proto_init() {
	if File_prodvana_async_task_task_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prodvana_async_task_task_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskResult); i {
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
			RawDescriptor: file_prodvana_async_task_task_metadata_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prodvana_async_task_task_metadata_proto_goTypes,
		DependencyIndexes: file_prodvana_async_task_task_metadata_proto_depIdxs,
		EnumInfos:         file_prodvana_async_task_task_metadata_proto_enumTypes,
		MessageInfos:      file_prodvana_async_task_task_metadata_proto_msgTypes,
	}.Build()
	File_prodvana_async_task_task_metadata_proto = out.File
	file_prodvana_async_task_task_metadata_proto_rawDesc = nil
	file_prodvana_async_task_task_metadata_proto_goTypes = nil
	file_prodvana_async_task_task_metadata_proto_depIdxs = nil
}
