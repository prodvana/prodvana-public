// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.10
// source: prodvana/pvn_wrapper/output.proto

package pvn_wrapper

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

type Output struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExitCode         int32         `protobuf:"varint,1,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`   // Exit code of wrapped process. -1 if process failed to execute.
	ExecError        string        `protobuf:"bytes,2,opt,name=exec_error,json=execError,proto3" json:"exec_error,omitempty"` // Internal error when trying to execute wrapped process.
	StdoutBlobId     string        `protobuf:"bytes,3,opt,name=stdout_blob_id,json=stdoutBlobId,proto3" json:"stdout_blob_id,omitempty"`
	StderrBlobId     string        `protobuf:"bytes,4,opt,name=stderr_blob_id,json=stderrBlobId,proto3" json:"stderr_blob_id,omitempty"`
	Version          string        `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`                                              // Wrapper version.
	StartTimestampNs int64         `protobuf:"varint,6,opt,name=start_timestamp_ns,json=startTimestampNs,proto3" json:"start_timestamp_ns,omitempty"` // Timestamp when the process began executing, in ns.
	DurationNs       int64         `protobuf:"varint,7,opt,name=duration_ns,json=durationNs,proto3" json:"duration_ns,omitempty"`                     // Total execution duration of the process, in ns.
	Files            []*OutputFile `protobuf:"bytes,8,rep,name=files,proto3" json:"files,omitempty"`
	Hostname         string        `protobuf:"bytes,9,opt,name=hostname,proto3" json:"hostname,omitempty"`
}

func (x *Output) Reset() {
	*x = Output{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_pvn_wrapper_output_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Output) ProtoMessage() {}

func (x *Output) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_pvn_wrapper_output_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Output.ProtoReflect.Descriptor instead.
func (*Output) Descriptor() ([]byte, []int) {
	return file_prodvana_pvn_wrapper_output_proto_rawDescGZIP(), []int{0}
}

func (x *Output) GetExitCode() int32 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

func (x *Output) GetExecError() string {
	if x != nil {
		return x.ExecError
	}
	return ""
}

func (x *Output) GetStdoutBlobId() string {
	if x != nil {
		return x.StdoutBlobId
	}
	return ""
}

func (x *Output) GetStderrBlobId() string {
	if x != nil {
		return x.StderrBlobId
	}
	return ""
}

func (x *Output) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Output) GetStartTimestampNs() int64 {
	if x != nil {
		return x.StartTimestampNs
	}
	return 0
}

func (x *Output) GetDurationNs() int64 {
	if x != nil {
		return x.DurationNs
	}
	return 0
}

func (x *Output) GetFiles() []*OutputFile {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *Output) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

type OutputFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ContentBlobId string `protobuf:"bytes,2,opt,name=content_blob_id,json=contentBlobId,proto3" json:"content_blob_id,omitempty"`
}

func (x *OutputFile) Reset() {
	*x = OutputFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_pvn_wrapper_output_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutputFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutputFile) ProtoMessage() {}

func (x *OutputFile) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_pvn_wrapper_output_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutputFile.ProtoReflect.Descriptor instead.
func (*OutputFile) Descriptor() ([]byte, []int) {
	return file_prodvana_pvn_wrapper_output_proto_rawDescGZIP(), []int{1}
}

func (x *OutputFile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OutputFile) GetContentBlobId() string {
	if x != nil {
		return x.ContentBlobId
	}
	return ""
}

var File_prodvana_pvn_wrapper_output_proto protoreflect.FileDescriptor

var file_prodvana_pvn_wrapper_output_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x76, 0x6e, 0x5f, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x70, 0x76,
	0x6e, 0x5f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x22, 0xcd, 0x02, 0x0a, 0x06, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x65, 0x63, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x65, 0x63, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x24, 0x0a, 0x0e, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x62, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74,
	0x42, 0x6c, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72,
	0x5f, 0x62, 0x6c, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x42, 0x6c, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x10, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x4e, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x73, 0x12, 0x36, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e,
	0x70, 0x76, 0x6e, 0x5f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x0a, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x6f,
	0x62, 0x49, 0x64, 0x42, 0x50, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76,
	0x61, 0x6e, 0x61, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x76, 0x6e, 0x5f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prodvana_pvn_wrapper_output_proto_rawDescOnce sync.Once
	file_prodvana_pvn_wrapper_output_proto_rawDescData = file_prodvana_pvn_wrapper_output_proto_rawDesc
)

func file_prodvana_pvn_wrapper_output_proto_rawDescGZIP() []byte {
	file_prodvana_pvn_wrapper_output_proto_rawDescOnce.Do(func() {
		file_prodvana_pvn_wrapper_output_proto_rawDescData = protoimpl.X.CompressGZIP(file_prodvana_pvn_wrapper_output_proto_rawDescData)
	})
	return file_prodvana_pvn_wrapper_output_proto_rawDescData
}

var file_prodvana_pvn_wrapper_output_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_prodvana_pvn_wrapper_output_proto_goTypes = []interface{}{
	(*Output)(nil),     // 0: prodvana.pvn_wrapper.Output
	(*OutputFile)(nil), // 1: prodvana.pvn_wrapper.OutputFile
}
var file_prodvana_pvn_wrapper_output_proto_depIdxs = []int32{
	1, // 0: prodvana.pvn_wrapper.Output.files:type_name -> prodvana.pvn_wrapper.OutputFile
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_prodvana_pvn_wrapper_output_proto_init() }
func file_prodvana_pvn_wrapper_output_proto_init() {
	if File_prodvana_pvn_wrapper_output_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prodvana_pvn_wrapper_output_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Output); i {
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
		file_prodvana_pvn_wrapper_output_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutputFile); i {
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
			RawDescriptor: file_prodvana_pvn_wrapper_output_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prodvana_pvn_wrapper_output_proto_goTypes,
		DependencyIndexes: file_prodvana_pvn_wrapper_output_proto_depIdxs,
		MessageInfos:      file_prodvana_pvn_wrapper_output_proto_msgTypes,
	}.Build()
	File_prodvana_pvn_wrapper_output_proto = out.File
	file_prodvana_pvn_wrapper_output_proto_rawDesc = nil
	file_prodvana_pvn_wrapper_output_proto_goTypes = nil
	file_prodvana_pvn_wrapper_output_proto_depIdxs = nil
}
