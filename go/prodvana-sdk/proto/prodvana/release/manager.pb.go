// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.10
// source: prodvana/release/manager.proto

package release

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	object "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/object"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type RecordReleaseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *ReleaseConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	// If true, create release with pending status that is meant to be updated later to either success or failure.
	// By default, releases are created with status SUCCEEDED.
	Pending bool `protobuf:"varint,2,opt,name=pending,proto3" json:"pending,omitempty"`
}

func (x *RecordReleaseReq) Reset() {
	*x = RecordReleaseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_release_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordReleaseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordReleaseReq) ProtoMessage() {}

func (x *RecordReleaseReq) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_release_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordReleaseReq.ProtoReflect.Descriptor instead.
func (*RecordReleaseReq) Descriptor() ([]byte, []int) {
	return file_prodvana_release_manager_proto_rawDescGZIP(), []int{0}
}

func (x *RecordReleaseReq) GetConfig() *ReleaseConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *RecordReleaseReq) GetPending() bool {
	if x != nil {
		return x.Pending
	}
	return false
}

type RecordReleaseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *object.ObjectMeta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *RecordReleaseResp) Reset() {
	*x = RecordReleaseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_release_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordReleaseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordReleaseResp) ProtoMessage() {}

func (x *RecordReleaseResp) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_release_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordReleaseResp.ProtoReflect.Descriptor instead.
func (*RecordReleaseResp) Descriptor() ([]byte, []int) {
	return file_prodvana_release_manager_proto_rawDescGZIP(), []int{1}
}

func (x *RecordReleaseResp) GetMeta() *object.ObjectMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type UpdateReleaseStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReleaseId string        `protobuf:"bytes,1,opt,name=release_id,json=releaseId,proto3" json:"release_id,omitempty"`
	Status    ReleaseStatus `protobuf:"varint,2,opt,name=status,proto3,enum=prodvana.release.ReleaseStatus" json:"status,omitempty"`
}

func (x *UpdateReleaseStatusReq) Reset() {
	*x = UpdateReleaseStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_release_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReleaseStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReleaseStatusReq) ProtoMessage() {}

func (x *UpdateReleaseStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_release_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReleaseStatusReq.ProtoReflect.Descriptor instead.
func (*UpdateReleaseStatusReq) Descriptor() ([]byte, []int) {
	return file_prodvana_release_manager_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateReleaseStatusReq) GetReleaseId() string {
	if x != nil {
		return x.ReleaseId
	}
	return ""
}

func (x *UpdateReleaseStatusReq) GetStatus() ReleaseStatus {
	if x != nil {
		return x.Status
	}
	return ReleaseStatus_UNKNOWN_STATUS
}

type UpdateReleaseStatusResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status ReleaseStatus `protobuf:"varint,1,opt,name=status,proto3,enum=prodvana.release.ReleaseStatus" json:"status,omitempty"`
}

func (x *UpdateReleaseStatusResp) Reset() {
	*x = UpdateReleaseStatusResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_release_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReleaseStatusResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReleaseStatusResp) ProtoMessage() {}

func (x *UpdateReleaseStatusResp) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_release_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReleaseStatusResp.ProtoReflect.Descriptor instead.
func (*UpdateReleaseStatusResp) Descriptor() ([]byte, []int) {
	return file_prodvana_release_manager_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateReleaseStatusResp) GetStatus() ReleaseStatus {
	if x != nil {
		return x.Status
	}
	return ReleaseStatus_UNKNOWN_STATUS
}

type ListReleasesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// filters for listing releases. Multiple filters are OR'ed together.
	Filters           []*ListReleasesReq_Filter `protobuf:"bytes,1,rep,name=filters,proto3" json:"filters,omitempty"`
	Filter            *ListReleasesReq_Filter   `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`                                                  // escape hatch to support openAPI, which cannot handle repeated list of messages on GET requests. This is joined to the filters list with an OR.
	StartingReleaseId string                    `protobuf:"bytes,3,opt,name=starting_release_id,json=startingReleaseId,proto3" json:"starting_release_id,omitempty"` // newer release, inclusive
	EndingReleaseId   string                    `protobuf:"bytes,4,opt,name=ending_release_id,json=endingReleaseId,proto3" json:"ending_release_id,omitempty"`       // older release, exclusive
	PageToken         string                    `protobuf:"bytes,5,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	PageSize          int32                     `protobuf:"varint,6,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListReleasesReq) Reset() {
	*x = ListReleasesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_release_manager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReleasesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReleasesReq) ProtoMessage() {}

func (x *ListReleasesReq) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_release_manager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReleasesReq.ProtoReflect.Descriptor instead.
func (*ListReleasesReq) Descriptor() ([]byte, []int) {
	return file_prodvana_release_manager_proto_rawDescGZIP(), []int{4}
}

func (x *ListReleasesReq) GetFilters() []*ListReleasesReq_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *ListReleasesReq) GetFilter() *ListReleasesReq_Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *ListReleasesReq) GetStartingReleaseId() string {
	if x != nil {
		return x.StartingReleaseId
	}
	return ""
}

func (x *ListReleasesReq) GetEndingReleaseId() string {
	if x != nil {
		return x.EndingReleaseId
	}
	return ""
}

func (x *ListReleasesReq) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListReleasesReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListReleasesResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Releases      []*Release `protobuf:"bytes,1,rep,name=releases,proto3" json:"releases,omitempty"`
	NextPageToken string     `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListReleasesResp) Reset() {
	*x = ListReleasesResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_release_manager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReleasesResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReleasesResp) ProtoMessage() {}

func (x *ListReleasesResp) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_release_manager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReleasesResp.ProtoReflect.Descriptor instead.
func (*ListReleasesResp) Descriptor() ([]byte, []int) {
	return file_prodvana_release_manager_proto_rawDescGZIP(), []int{5}
}

func (x *ListReleasesResp) GetReleases() []*Release {
	if x != nil {
		return x.Releases
	}
	return nil
}

func (x *ListReleasesResp) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type ListReleasesReq_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// filters to releases for join(join(services, OR), join(release_channels, OR), AND)
	Services        []string `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	ReleaseChannels []string `protobuf:"bytes,2,rep,name=release_channels,json=releaseChannels,proto3" json:"release_channels,omitempty"`
}

func (x *ListReleasesReq_Filter) Reset() {
	*x = ListReleasesReq_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_release_manager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReleasesReq_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReleasesReq_Filter) ProtoMessage() {}

func (x *ListReleasesReq_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_release_manager_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReleasesReq_Filter.ProtoReflect.Descriptor instead.
func (*ListReleasesReq_Filter) Descriptor() ([]byte, []int) {
	return file_prodvana_release_manager_proto_rawDescGZIP(), []int{4, 0}
}

func (x *ListReleasesReq_Filter) GetServices() []string {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *ListReleasesReq_Filter) GetReleaseChannels() []string {
	if x != nil {
		return x.ReleaseChannels
	}
	return nil
}

var File_prodvana_release_manager_proto protoreflect.FileDescriptor

var file_prodvana_release_manager_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x72,
	0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2f, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x10, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x12, 0x41, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76,
	0x61, 0x6e, 0x61, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x44, 0x0a, 0x11, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2f, 0x0a, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76,
	0x61, 0x6e, 0x61, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x70, 0x0a, 0x16, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e,
	0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x52, 0x0a,
	0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76,
	0x61, 0x6e, 0x61, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x80, 0x03, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x12, 0x42, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e,
	0x61, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x40, 0x0a, 0x06, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x76, 0x61, 0x6e, 0x61, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x13, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x1a, 0x4f, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x73, 0x22, 0x71, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x35, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x12,
	0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xed, 0x03, 0x0a, 0x0e, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x71, 0x0a, 0x0d, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x22, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x23, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22,
	0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x12, 0x99, 0x01,
	0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61,
	0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x1a,
	0x29, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x27, 0x3a, 0x01, 0x2a, 0x22, 0x22, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x3d,
	0x2a, 0x7d, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x6b, 0x0a, 0x0c, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x76, 0x61, 0x6e, 0x61, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x12, 0x5f, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x21, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a,
	0x22, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x00, 0x30, 0x01, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70,
	0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x67,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prodvana_release_manager_proto_rawDescOnce sync.Once
	file_prodvana_release_manager_proto_rawDescData = file_prodvana_release_manager_proto_rawDesc
)

func file_prodvana_release_manager_proto_rawDescGZIP() []byte {
	file_prodvana_release_manager_proto_rawDescOnce.Do(func() {
		file_prodvana_release_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_prodvana_release_manager_proto_rawDescData)
	})
	return file_prodvana_release_manager_proto_rawDescData
}

var file_prodvana_release_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_prodvana_release_manager_proto_goTypes = []interface{}{
	(*RecordReleaseReq)(nil),        // 0: prodvana.release.RecordReleaseReq
	(*RecordReleaseResp)(nil),       // 1: prodvana.release.RecordReleaseResp
	(*UpdateReleaseStatusReq)(nil),  // 2: prodvana.release.UpdateReleaseStatusReq
	(*UpdateReleaseStatusResp)(nil), // 3: prodvana.release.UpdateReleaseStatusResp
	(*ListReleasesReq)(nil),         // 4: prodvana.release.ListReleasesReq
	(*ListReleasesResp)(nil),        // 5: prodvana.release.ListReleasesResp
	(*ListReleasesReq_Filter)(nil),  // 6: prodvana.release.ListReleasesReq.Filter
	(*ReleaseConfig)(nil),           // 7: prodvana.release.ReleaseConfig
	(*object.ObjectMeta)(nil),       // 8: prodvana.object.ObjectMeta
	(ReleaseStatus)(0),              // 9: prodvana.release.ReleaseStatus
	(*Release)(nil),                 // 10: prodvana.release.Release
}
var file_prodvana_release_manager_proto_depIdxs = []int32{
	7,  // 0: prodvana.release.RecordReleaseReq.config:type_name -> prodvana.release.ReleaseConfig
	8,  // 1: prodvana.release.RecordReleaseResp.meta:type_name -> prodvana.object.ObjectMeta
	9,  // 2: prodvana.release.UpdateReleaseStatusReq.status:type_name -> prodvana.release.ReleaseStatus
	9,  // 3: prodvana.release.UpdateReleaseStatusResp.status:type_name -> prodvana.release.ReleaseStatus
	6,  // 4: prodvana.release.ListReleasesReq.filters:type_name -> prodvana.release.ListReleasesReq.Filter
	6,  // 5: prodvana.release.ListReleasesReq.filter:type_name -> prodvana.release.ListReleasesReq.Filter
	10, // 6: prodvana.release.ListReleasesResp.releases:type_name -> prodvana.release.Release
	0,  // 7: prodvana.release.ReleaseManager.RecordRelease:input_type -> prodvana.release.RecordReleaseReq
	2,  // 8: prodvana.release.ReleaseManager.UpdateReleaseStatus:input_type -> prodvana.release.UpdateReleaseStatusReq
	4,  // 9: prodvana.release.ReleaseManager.ListReleases:input_type -> prodvana.release.ListReleasesReq
	4,  // 10: prodvana.release.ReleaseManager.ListReleasesStream:input_type -> prodvana.release.ListReleasesReq
	1,  // 11: prodvana.release.ReleaseManager.RecordRelease:output_type -> prodvana.release.RecordReleaseResp
	3,  // 12: prodvana.release.ReleaseManager.UpdateReleaseStatus:output_type -> prodvana.release.UpdateReleaseStatusResp
	5,  // 13: prodvana.release.ReleaseManager.ListReleases:output_type -> prodvana.release.ListReleasesResp
	5,  // 14: prodvana.release.ReleaseManager.ListReleasesStream:output_type -> prodvana.release.ListReleasesResp
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_prodvana_release_manager_proto_init() }
func file_prodvana_release_manager_proto_init() {
	if File_prodvana_release_manager_proto != nil {
		return
	}
	file_prodvana_release_object_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_prodvana_release_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordReleaseReq); i {
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
		file_prodvana_release_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordReleaseResp); i {
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
		file_prodvana_release_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReleaseStatusReq); i {
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
		file_prodvana_release_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReleaseStatusResp); i {
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
		file_prodvana_release_manager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReleasesReq); i {
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
		file_prodvana_release_manager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReleasesResp); i {
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
		file_prodvana_release_manager_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReleasesReq_Filter); i {
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
			RawDescriptor: file_prodvana_release_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_prodvana_release_manager_proto_goTypes,
		DependencyIndexes: file_prodvana_release_manager_proto_depIdxs,
		MessageInfos:      file_prodvana_release_manager_proto_msgTypes,
	}.Build()
	File_prodvana_release_manager_proto = out.File
	file_prodvana_release_manager_proto_rawDesc = nil
	file_prodvana_release_manager_proto_goTypes = nil
	file_prodvana_release_manager_proto_depIdxs = nil
}