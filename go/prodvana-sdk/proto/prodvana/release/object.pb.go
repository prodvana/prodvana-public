// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.10
// source: prodvana/release/object.proto

package release

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	object "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/object"
	repo "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/repo"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReleaseStatus int32

const (
	ReleaseStatus_UNKNOWN_STATUS ReleaseStatus = 0
	ReleaseStatus_PENDING        ReleaseStatus = 1
	ReleaseStatus_SUCCEEDED      ReleaseStatus = 2
	ReleaseStatus_FAILED         ReleaseStatus = 3
)

// Enum value maps for ReleaseStatus.
var (
	ReleaseStatus_name = map[int32]string{
		0: "UNKNOWN_STATUS",
		1: "PENDING",
		2: "SUCCEEDED",
		3: "FAILED",
	}
	ReleaseStatus_value = map[string]int32{
		"UNKNOWN_STATUS": 0,
		"PENDING":        1,
		"SUCCEEDED":      2,
		"FAILED":         3,
	}
)

func (x ReleaseStatus) Enum() *ReleaseStatus {
	p := new(ReleaseStatus)
	*p = x
	return p
}

func (x ReleaseStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReleaseStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_prodvana_release_object_proto_enumTypes[0].Descriptor()
}

func (ReleaseStatus) Type() protoreflect.EnumType {
	return &file_prodvana_release_object_proto_enumTypes[0]
}

func (x ReleaseStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReleaseStatus.Descriptor instead.
func (ReleaseStatus) EnumDescriptor() ([]byte, []int) {
	return file_prodvana_release_object_proto_rawDescGZIP(), []int{0}
}

type ReleaseConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreationTimestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=creation_timestamp,json=creationTimestamp,proto3" json:"creation_timestamp,omitempty"` // must be unset on input
	DeploymentSystem  string                 `protobuf:"bytes,2,opt,name=deployment_system,json=deploymentSystem,proto3" json:"deployment_system,omitempty"`
	// required when recording releases from external systems
	Service string `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty"`
	// required when recording releases from external systems
	ReleaseChannel string `protobuf:"bytes,4,opt,name=release_channel,json=releaseChannel,proto3" json:"release_channel,omitempty"`
	// required when recording releases from external systems
	Application      string `protobuf:"bytes,8,opt,name=application,proto3" json:"application,omitempty"`
	Repository       string `protobuf:"bytes,5,opt,name=repository,proto3" json:"repository,omitempty"`             // e.g. github.com/foo/bar
	CommitId         string `protobuf:"bytes,6,opt,name=commit_id,json=commitId,proto3" json:"commit_id,omitempty"` // commit hash
	User             string `protobuf:"bytes,7,opt,name=user,proto3" json:"user,omitempty"`                         // if known
	ApplicationId    string `protobuf:"bytes,9,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	ServiceId        string `protobuf:"bytes,10,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	ReleaseChannelId string `protobuf:"bytes,11,opt,name=release_channel_id,json=releaseChannelId,proto3" json:"release_channel_id,omitempty"`
	ServiceVersion   string `protobuf:"bytes,12,opt,name=service_version,json=serviceVersion,proto3" json:"service_version,omitempty"` // next tag: 13
}

func (x *ReleaseConfig) Reset() {
	*x = ReleaseConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_release_object_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseConfig) ProtoMessage() {}

func (x *ReleaseConfig) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_release_object_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseConfig.ProtoReflect.Descriptor instead.
func (*ReleaseConfig) Descriptor() ([]byte, []int) {
	return file_prodvana_release_object_proto_rawDescGZIP(), []int{0}
}

func (x *ReleaseConfig) GetCreationTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.CreationTimestamp
	}
	return nil
}

func (x *ReleaseConfig) GetDeploymentSystem() string {
	if x != nil {
		return x.DeploymentSystem
	}
	return ""
}

func (x *ReleaseConfig) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *ReleaseConfig) GetReleaseChannel() string {
	if x != nil {
		return x.ReleaseChannel
	}
	return ""
}

func (x *ReleaseConfig) GetApplication() string {
	if x != nil {
		return x.Application
	}
	return ""
}

func (x *ReleaseConfig) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *ReleaseConfig) GetCommitId() string {
	if x != nil {
		return x.CommitId
	}
	return ""
}

func (x *ReleaseConfig) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *ReleaseConfig) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

func (x *ReleaseConfig) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *ReleaseConfig) GetReleaseChannelId() string {
	if x != nil {
		return x.ReleaseChannelId
	}
	return ""
}

func (x *ReleaseConfig) GetServiceVersion() string {
	if x != nil {
		return x.ServiceVersion
	}
	return ""
}

type ReleaseState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status              ReleaseStatus          `protobuf:"varint,1,opt,name=status,proto3,enum=prodvana.release.ReleaseStatus" json:"status,omitempty"`
	LastUpdateTimestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_update_timestamp,json=lastUpdateTimestamp,proto3" json:"last_update_timestamp,omitempty"`
}

func (x *ReleaseState) Reset() {
	*x = ReleaseState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_release_object_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseState) ProtoMessage() {}

func (x *ReleaseState) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_release_object_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseState.ProtoReflect.Descriptor instead.
func (*ReleaseState) Descriptor() ([]byte, []int) {
	return file_prodvana_release_object_proto_rawDescGZIP(), []int{1}
}

func (x *ReleaseState) GetStatus() ReleaseStatus {
	if x != nil {
		return x.Status
	}
	return ReleaseStatus_UNKNOWN_STATUS
}

func (x *ReleaseState) GetLastUpdateTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdateTimestamp
	}
	return nil
}

type ImpactAnalysisComparison struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// commits likely to be impactful, prev_commit_id and new_commit_id have a merge base
	RelevantCommits   []*repo.Commit `protobuf:"bytes,1,rep,name=relevant_commits,json=relevantCommits,proto3" json:"relevant_commits,omitempty"`
	UnanalyzedCommits int64          `protobuf:"varint,2,opt,name=unanalyzed_commits,json=unanalyzedCommits,proto3" json:"unanalyzed_commits,omitempty"`
}

func (x *ImpactAnalysisComparison) Reset() {
	*x = ImpactAnalysisComparison{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_release_object_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImpactAnalysisComparison) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImpactAnalysisComparison) ProtoMessage() {}

func (x *ImpactAnalysisComparison) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_release_object_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImpactAnalysisComparison.ProtoReflect.Descriptor instead.
func (*ImpactAnalysisComparison) Descriptor() ([]byte, []int) {
	return file_prodvana_release_object_proto_rawDescGZIP(), []int{2}
}

func (x *ImpactAnalysisComparison) GetRelevantCommits() []*repo.Commit {
	if x != nil {
		return x.RelevantCommits
	}
	return nil
}

func (x *ImpactAnalysisComparison) GetUnanalyzedCommits() int64 {
	if x != nil {
		return x.UnanalyzedCommits
	}
	return 0
}

type ReleaseComparison struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prev           *object.ObjectMeta        `protobuf:"bytes,1,opt,name=prev,proto3" json:"prev,omitempty"`
	PrevRepository string                    `protobuf:"bytes,2,opt,name=prev_repository,json=prevRepository,proto3" json:"prev_repository,omitempty"`
	NewRepository  string                    `protobuf:"bytes,3,opt,name=new_repository,json=newRepository,proto3" json:"new_repository,omitempty"`
	PrevCommitId   string                    `protobuf:"bytes,4,opt,name=prev_commit_id,json=prevCommitId,proto3" json:"prev_commit_id,omitempty"`
	NewCommitId    string                    `protobuf:"bytes,5,opt,name=new_commit_id,json=newCommitId,proto3" json:"new_commit_id,omitempty"`
	ImpactAnalysis *ImpactAnalysisComparison `protobuf:"bytes,6,opt,name=impact_analysis,json=impactAnalysis,proto3" json:"impact_analysis,omitempty"`
	TotalCommits   *wrapperspb.Int64Value    `protobuf:"bytes,7,opt,name=total_commits,json=totalCommits,proto3" json:"total_commits,omitempty"` // only set if the previous commit is set and is from the same repo, and the repo is linked to prodvana
}

func (x *ReleaseComparison) Reset() {
	*x = ReleaseComparison{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_release_object_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseComparison) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseComparison) ProtoMessage() {}

func (x *ReleaseComparison) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_release_object_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseComparison.ProtoReflect.Descriptor instead.
func (*ReleaseComparison) Descriptor() ([]byte, []int) {
	return file_prodvana_release_object_proto_rawDescGZIP(), []int{3}
}

func (x *ReleaseComparison) GetPrev() *object.ObjectMeta {
	if x != nil {
		return x.Prev
	}
	return nil
}

func (x *ReleaseComparison) GetPrevRepository() string {
	if x != nil {
		return x.PrevRepository
	}
	return ""
}

func (x *ReleaseComparison) GetNewRepository() string {
	if x != nil {
		return x.NewRepository
	}
	return ""
}

func (x *ReleaseComparison) GetPrevCommitId() string {
	if x != nil {
		return x.PrevCommitId
	}
	return ""
}

func (x *ReleaseComparison) GetNewCommitId() string {
	if x != nil {
		return x.NewCommitId
	}
	return ""
}

func (x *ReleaseComparison) GetImpactAnalysis() *ImpactAnalysisComparison {
	if x != nil {
		return x.ImpactAnalysis
	}
	return nil
}

func (x *ReleaseComparison) GetTotalCommits() *wrapperspb.Int64Value {
	if x != nil {
		return x.TotalCommits
	}
	return nil
}

type Release struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta   *object.ObjectMeta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Config *ReleaseConfig     `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	// TODO(naphat) should this really be part of the proto here, or should it be a separate endpoint so we can request arbitrary comparison?
	Comparison *ReleaseComparison `protobuf:"bytes,3,opt,name=comparison,proto3" json:"comparison,omitempty"`
	State      *ReleaseState      `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *Release) Reset() {
	*x = Release{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_release_object_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Release) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Release) ProtoMessage() {}

func (x *Release) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_release_object_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Release.ProtoReflect.Descriptor instead.
func (*Release) Descriptor() ([]byte, []int) {
	return file_prodvana_release_object_proto_rawDescGZIP(), []int{4}
}

func (x *Release) GetMeta() *object.ObjectMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Release) GetConfig() *ReleaseConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *Release) GetComparison() *ReleaseComparison {
	if x != nil {
		return x.Comparison
	}
	return nil
}

func (x *Release) GetState() *ReleaseState {
	if x != nil {
		return x.State
	}
	return nil
}

var File_prodvana_release_object_proto protoreflect.FileDescriptor

var file_prodvana_release_object_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18,
	0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x72, 0x65,
	0x70, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe3, 0x03, 0x0a, 0x0d, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x49, 0x0a, 0x12, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x11, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x34,
	0x0a, 0x11, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x10, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x27,
	0x0a, 0x0f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x27,
	0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x97, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76,
	0x61, 0x6e, 0x61, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x4e, 0x0a, 0x15, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x13, 0x6c, 0x61,
	0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x22, 0x8b, 0x01, 0x0a, 0x18, 0x49, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x41, 0x6e, 0x61, 0x6c,
	0x79, 0x73, 0x69, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x12, 0x40,
	0x0a, 0x10, 0x72, 0x65, 0x6c, 0x65, 0x76, 0x61, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76,
	0x61, 0x6e, 0x61, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52,
	0x0f, 0x72, 0x65, 0x6c, 0x65, 0x76, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73,
	0x12, 0x2d, 0x0a, 0x12, 0x75, 0x6e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x64, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x75, 0x6e,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x22,
	0xf5, 0x02, 0x0a, 0x11, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x72, 0x69, 0x73, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x04, 0x70, 0x72, 0x65, 0x76, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x52, 0x04, 0x70, 0x72, 0x65, 0x76, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x72,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x70, 0x72, 0x65, 0x76, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12,
	0x25, 0x0a, 0x0e, 0x6e, 0x65, 0x77, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x77, 0x52, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x70, 0x72, 0x65, 0x76, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d,
	0x6e, 0x65, 0x77, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x64,
	0x12, 0x53, 0x0a, 0x0f, 0x69, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x73, 0x69, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x76, 0x61, 0x6e, 0x61, 0x2e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x49, 0x6d, 0x70,
	0x61, 0x63, 0x74, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x72, 0x69, 0x73, 0x6f, 0x6e, 0x52, 0x0e, 0x69, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x41, 0x6e, 0x61,
	0x6c, 0x79, 0x73, 0x69, 0x73, 0x12, 0x40, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49,
	0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x22, 0xee, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x12, 0x37, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e,
	0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x43, 0x0a,
	0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x72, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73,
	0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x72, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2a, 0x4b, 0x0a, 0x0d, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55,
	0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49,
	0x4c, 0x45, 0x44, 0x10, 0x03, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x76, 0x61, 0x6e, 0x61, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x67, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prodvana_release_object_proto_rawDescOnce sync.Once
	file_prodvana_release_object_proto_rawDescData = file_prodvana_release_object_proto_rawDesc
)

func file_prodvana_release_object_proto_rawDescGZIP() []byte {
	file_prodvana_release_object_proto_rawDescOnce.Do(func() {
		file_prodvana_release_object_proto_rawDescData = protoimpl.X.CompressGZIP(file_prodvana_release_object_proto_rawDescData)
	})
	return file_prodvana_release_object_proto_rawDescData
}

var file_prodvana_release_object_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_prodvana_release_object_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_prodvana_release_object_proto_goTypes = []interface{}{
	(ReleaseStatus)(0),               // 0: prodvana.release.ReleaseStatus
	(*ReleaseConfig)(nil),            // 1: prodvana.release.ReleaseConfig
	(*ReleaseState)(nil),             // 2: prodvana.release.ReleaseState
	(*ImpactAnalysisComparison)(nil), // 3: prodvana.release.ImpactAnalysisComparison
	(*ReleaseComparison)(nil),        // 4: prodvana.release.ReleaseComparison
	(*Release)(nil),                  // 5: prodvana.release.Release
	(*timestamppb.Timestamp)(nil),    // 6: google.protobuf.Timestamp
	(*repo.Commit)(nil),              // 7: prodvana.repo.Commit
	(*object.ObjectMeta)(nil),        // 8: prodvana.object.ObjectMeta
	(*wrapperspb.Int64Value)(nil),    // 9: google.protobuf.Int64Value
}
var file_prodvana_release_object_proto_depIdxs = []int32{
	6,  // 0: prodvana.release.ReleaseConfig.creation_timestamp:type_name -> google.protobuf.Timestamp
	0,  // 1: prodvana.release.ReleaseState.status:type_name -> prodvana.release.ReleaseStatus
	6,  // 2: prodvana.release.ReleaseState.last_update_timestamp:type_name -> google.protobuf.Timestamp
	7,  // 3: prodvana.release.ImpactAnalysisComparison.relevant_commits:type_name -> prodvana.repo.Commit
	8,  // 4: prodvana.release.ReleaseComparison.prev:type_name -> prodvana.object.ObjectMeta
	3,  // 5: prodvana.release.ReleaseComparison.impact_analysis:type_name -> prodvana.release.ImpactAnalysisComparison
	9,  // 6: prodvana.release.ReleaseComparison.total_commits:type_name -> google.protobuf.Int64Value
	8,  // 7: prodvana.release.Release.meta:type_name -> prodvana.object.ObjectMeta
	1,  // 8: prodvana.release.Release.config:type_name -> prodvana.release.ReleaseConfig
	4,  // 9: prodvana.release.Release.comparison:type_name -> prodvana.release.ReleaseComparison
	2,  // 10: prodvana.release.Release.state:type_name -> prodvana.release.ReleaseState
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_prodvana_release_object_proto_init() }
func file_prodvana_release_object_proto_init() {
	if File_prodvana_release_object_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prodvana_release_object_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseConfig); i {
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
		file_prodvana_release_object_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseState); i {
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
		file_prodvana_release_object_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImpactAnalysisComparison); i {
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
		file_prodvana_release_object_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseComparison); i {
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
		file_prodvana_release_object_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Release); i {
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
			RawDescriptor: file_prodvana_release_object_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prodvana_release_object_proto_goTypes,
		DependencyIndexes: file_prodvana_release_object_proto_depIdxs,
		EnumInfos:         file_prodvana_release_object_proto_enumTypes,
		MessageInfos:      file_prodvana_release_object_proto_msgTypes,
	}.Build()
	File_prodvana_release_object_proto = out.File
	file_prodvana_release_object_proto_rawDesc = nil
	file_prodvana_release_object_proto_goTypes = nil
	file_prodvana_release_object_proto_depIdxs = nil
}
