// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.10
// source: prodvana/release_channel/object.proto

package release_channel

import (
	object "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/object"
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

type ReleaseChannelProtectionAttachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protection     string `protobuf:"bytes,1,opt,name=protection,proto3" json:"protection,omitempty"`
	Attachment     string `protobuf:"bytes,2,opt,name=attachment,proto3" json:"attachment,omitempty"`
	DesiredStateId string `protobuf:"bytes,3,opt,name=desired_state_id,json=desiredStateId,proto3" json:"desired_state_id,omitempty"`
}

func (x *ReleaseChannelProtectionAttachment) Reset() {
	*x = ReleaseChannelProtectionAttachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_release_channel_object_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseChannelProtectionAttachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseChannelProtectionAttachment) ProtoMessage() {}

func (x *ReleaseChannelProtectionAttachment) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_release_channel_object_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseChannelProtectionAttachment.ProtoReflect.Descriptor instead.
func (*ReleaseChannelProtectionAttachment) Descriptor() ([]byte, []int) {
	return file_prodvana_release_channel_object_proto_rawDescGZIP(), []int{0}
}

func (x *ReleaseChannelProtectionAttachment) GetProtection() string {
	if x != nil {
		return x.Protection
	}
	return ""
}

func (x *ReleaseChannelProtectionAttachment) GetAttachment() string {
	if x != nil {
		return x.Attachment
	}
	return ""
}

func (x *ReleaseChannelProtectionAttachment) GetDesiredStateId() string {
	if x != nil {
		return x.DesiredStateId
	}
	return ""
}

type ReleaseChannelState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProtectionAttachments []*ReleaseChannelProtectionAttachment `protobuf:"bytes,1,rep,name=protection_attachments,json=protectionAttachments,proto3" json:"protection_attachments,omitempty"`
}

func (x *ReleaseChannelState) Reset() {
	*x = ReleaseChannelState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_release_channel_object_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseChannelState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseChannelState) ProtoMessage() {}

func (x *ReleaseChannelState) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_release_channel_object_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseChannelState.ProtoReflect.Descriptor instead.
func (*ReleaseChannelState) Descriptor() ([]byte, []int) {
	return file_prodvana_release_channel_object_proto_rawDescGZIP(), []int{1}
}

func (x *ReleaseChannelState) GetProtectionAttachments() []*ReleaseChannelProtectionAttachment {
	if x != nil {
		return x.ProtectionAttachments
	}
	return nil
}

type ReleaseChannel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta   *object.ObjectMeta    `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Config *ReleaseChannelConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	State  *ReleaseChannelState  `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *ReleaseChannel) Reset() {
	*x = ReleaseChannel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_release_channel_object_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseChannel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseChannel) ProtoMessage() {}

func (x *ReleaseChannel) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_release_channel_object_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseChannel.ProtoReflect.Descriptor instead.
func (*ReleaseChannel) Descriptor() ([]byte, []int) {
	return file_prodvana_release_channel_object_proto_rawDescGZIP(), []int{2}
}

func (x *ReleaseChannel) GetMeta() *object.ObjectMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *ReleaseChannel) GetConfig() *ReleaseChannelConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *ReleaseChannel) GetState() *ReleaseChannelState {
	if x != nil {
		return x.State
	}
	return nil
}

var File_prodvana_release_channel_object_proto protoreflect.FileDescriptor

var file_prodvana_release_channel_object_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e,
	0x61, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x70,
	0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x22, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x64,
	0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x49, 0x64, 0x22, 0x8a, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x73, 0x0a,
	0x16, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x15, 0x70, 0x72, 0x6f,
	0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0xce, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x2f, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x46, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e,
	0x61, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x43,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x42, 0x54, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76,
	0x61, 0x6e, 0x61, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_prodvana_release_channel_object_proto_rawDescOnce sync.Once
	file_prodvana_release_channel_object_proto_rawDescData = file_prodvana_release_channel_object_proto_rawDesc
)

func file_prodvana_release_channel_object_proto_rawDescGZIP() []byte {
	file_prodvana_release_channel_object_proto_rawDescOnce.Do(func() {
		file_prodvana_release_channel_object_proto_rawDescData = protoimpl.X.CompressGZIP(file_prodvana_release_channel_object_proto_rawDescData)
	})
	return file_prodvana_release_channel_object_proto_rawDescData
}

var file_prodvana_release_channel_object_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_prodvana_release_channel_object_proto_goTypes = []interface{}{
	(*ReleaseChannelProtectionAttachment)(nil), // 0: prodvana.release_channel.ReleaseChannelProtectionAttachment
	(*ReleaseChannelState)(nil),                // 1: prodvana.release_channel.ReleaseChannelState
	(*ReleaseChannel)(nil),                     // 2: prodvana.release_channel.ReleaseChannel
	(*object.ObjectMeta)(nil),                  // 3: prodvana.object.ObjectMeta
	(*ReleaseChannelConfig)(nil),               // 4: prodvana.release_channel.ReleaseChannelConfig
}
var file_prodvana_release_channel_object_proto_depIdxs = []int32{
	0, // 0: prodvana.release_channel.ReleaseChannelState.protection_attachments:type_name -> prodvana.release_channel.ReleaseChannelProtectionAttachment
	3, // 1: prodvana.release_channel.ReleaseChannel.meta:type_name -> prodvana.object.ObjectMeta
	4, // 2: prodvana.release_channel.ReleaseChannel.config:type_name -> prodvana.release_channel.ReleaseChannelConfig
	1, // 3: prodvana.release_channel.ReleaseChannel.state:type_name -> prodvana.release_channel.ReleaseChannelState
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_prodvana_release_channel_object_proto_init() }
func file_prodvana_release_channel_object_proto_init() {
	if File_prodvana_release_channel_object_proto != nil {
		return
	}
	file_prodvana_release_channel_release_channel_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_prodvana_release_channel_object_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseChannelProtectionAttachment); i {
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
		file_prodvana_release_channel_object_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseChannelState); i {
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
		file_prodvana_release_channel_object_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseChannel); i {
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
			RawDescriptor: file_prodvana_release_channel_object_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prodvana_release_channel_object_proto_goTypes,
		DependencyIndexes: file_prodvana_release_channel_object_proto_depIdxs,
		MessageInfos:      file_prodvana_release_channel_object_proto_msgTypes,
	}.Build()
	File_prodvana_release_channel_object_proto = out.File
	file_prodvana_release_channel_object_proto_rawDesc = nil
	file_prodvana_release_channel_object_proto_goTypes = nil
	file_prodvana_release_channel_object_proto_depIdxs = nil
}
