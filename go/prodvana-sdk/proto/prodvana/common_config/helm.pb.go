// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.10
// source: prodvana/common_config/helm.proto

package common_config

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

type RemoteHelmChart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to RepoOneof:
	//
	//	*RemoteHelmChart_Repo
	RepoOneof    isRemoteHelmChart_RepoOneof `protobuf_oneof:"repo_oneof"`
	Chart        string                      `protobuf:"bytes,2,opt,name=chart,proto3" json:"chart,omitempty"`
	ChartVersion string                      `protobuf:"bytes,3,opt,name=chart_version,json=chartVersion,proto3" json:"chart_version,omitempty"`
}

func (x *RemoteHelmChart) Reset() {
	*x = RemoteHelmChart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_common_config_helm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteHelmChart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteHelmChart) ProtoMessage() {}

func (x *RemoteHelmChart) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_common_config_helm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteHelmChart.ProtoReflect.Descriptor instead.
func (*RemoteHelmChart) Descriptor() ([]byte, []int) {
	return file_prodvana_common_config_helm_proto_rawDescGZIP(), []int{0}
}

func (m *RemoteHelmChart) GetRepoOneof() isRemoteHelmChart_RepoOneof {
	if m != nil {
		return m.RepoOneof
	}
	return nil
}

func (x *RemoteHelmChart) GetRepo() string {
	if x, ok := x.GetRepoOneof().(*RemoteHelmChart_Repo); ok {
		return x.Repo
	}
	return ""
}

func (x *RemoteHelmChart) GetChart() string {
	if x != nil {
		return x.Chart
	}
	return ""
}

func (x *RemoteHelmChart) GetChartVersion() string {
	if x != nil {
		return x.ChartVersion
	}
	return ""
}

type isRemoteHelmChart_RepoOneof interface {
	isRemoteHelmChart_RepoOneof()
}

type RemoteHelmChart_Repo struct {
	Repo string `protobuf:"bytes,1,opt,name=repo,proto3,oneof"` // TODO(naphat) add integration support for private repositories
}

func (*RemoteHelmChart_Repo) isRemoteHelmChart_RepoOneof() {}

type HelmValuesOverrides struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to OverrideOneof:
	//
	//	*HelmValuesOverrides_Inlined
	//	*HelmValuesOverrides_Local
	OverrideOneof isHelmValuesOverrides_OverrideOneof `protobuf_oneof:"override_oneof"`
	// treat this as part of the above oneof, even though proto does not allow us to
	Map map[string]*EnvValue `protobuf:"bytes,3,rep,name=map,proto3" json:"map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *HelmValuesOverrides) Reset() {
	*x = HelmValuesOverrides{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_common_config_helm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelmValuesOverrides) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelmValuesOverrides) ProtoMessage() {}

func (x *HelmValuesOverrides) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_common_config_helm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelmValuesOverrides.ProtoReflect.Descriptor instead.
func (*HelmValuesOverrides) Descriptor() ([]byte, []int) {
	return file_prodvana_common_config_helm_proto_rawDescGZIP(), []int{1}
}

func (m *HelmValuesOverrides) GetOverrideOneof() isHelmValuesOverrides_OverrideOneof {
	if m != nil {
		return m.OverrideOneof
	}
	return nil
}

func (x *HelmValuesOverrides) GetInlined() string {
	if x, ok := x.GetOverrideOneof().(*HelmValuesOverrides_Inlined); ok {
		return x.Inlined
	}
	return ""
}

func (x *HelmValuesOverrides) GetLocal() *LocalConfig {
	if x, ok := x.GetOverrideOneof().(*HelmValuesOverrides_Local); ok {
		return x.Local
	}
	return nil
}

func (x *HelmValuesOverrides) GetMap() map[string]*EnvValue {
	if x != nil {
		return x.Map
	}
	return nil
}

type isHelmValuesOverrides_OverrideOneof interface {
	isHelmValuesOverrides_OverrideOneof()
}

type HelmValuesOverrides_Inlined struct {
	Inlined string `protobuf:"bytes,1,opt,name=inlined,proto3,oneof"`
}

type HelmValuesOverrides_Local struct {
	Local *LocalConfig `protobuf:"bytes,2,opt,name=local,proto3,oneof"`
}

func (*HelmValuesOverrides_Inlined) isHelmValuesOverrides_OverrideOneof() {}

func (*HelmValuesOverrides_Local) isHelmValuesOverrides_OverrideOneof() {}

type HelmConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ChartOneof:
	//
	//	*HelmConfig_Remote
	ChartOneof      isHelmConfig_ChartOneof `protobuf_oneof:"chart_oneof"`
	ValuesOverrides []*HelmValuesOverrides  `protobuf:"bytes,2,rep,name=values_overrides,json=valuesOverrides,proto3" json:"values_overrides,omitempty"`
	// optional release name, leave blank to have Prodvana generate one.
	// Mainly useful for migrating an existing helm release into Prodvana.
	ReleaseName string `protobuf:"bytes,3,opt,name=release_name,json=releaseName,proto3" json:"release_name,omitempty"`
	// used internally by Prodvana, do not set.
	Namespace string `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *HelmConfig) Reset() {
	*x = HelmConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_common_config_helm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelmConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelmConfig) ProtoMessage() {}

func (x *HelmConfig) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_common_config_helm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelmConfig.ProtoReflect.Descriptor instead.
func (*HelmConfig) Descriptor() ([]byte, []int) {
	return file_prodvana_common_config_helm_proto_rawDescGZIP(), []int{2}
}

func (m *HelmConfig) GetChartOneof() isHelmConfig_ChartOneof {
	if m != nil {
		return m.ChartOneof
	}
	return nil
}

func (x *HelmConfig) GetRemote() *RemoteHelmChart {
	if x, ok := x.GetChartOneof().(*HelmConfig_Remote); ok {
		return x.Remote
	}
	return nil
}

func (x *HelmConfig) GetValuesOverrides() []*HelmValuesOverrides {
	if x != nil {
		return x.ValuesOverrides
	}
	return nil
}

func (x *HelmConfig) GetReleaseName() string {
	if x != nil {
		return x.ReleaseName
	}
	return ""
}

func (x *HelmConfig) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type isHelmConfig_ChartOneof interface {
	isHelmConfig_ChartOneof()
}

type HelmConfig_Remote struct {
	Remote *RemoteHelmChart `protobuf:"bytes,1,opt,name=remote,proto3,oneof"` // TODO(naphat) chart archive support, local charts
}

func (*HelmConfig_Remote) isHelmConfig_ChartOneof() {}

var File_prodvana_common_config_helm_proto protoreflect.FileDescriptor

var file_prodvana_common_config_helm_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x68, 0x65, 0x6c, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x16, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e, 0x76,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x48, 0x65, 0x6c, 0x6d, 0x43, 0x68, 0x61, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x04, 0x72, 0x65,
	0x70, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f,
	0x12, 0x1d, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05, 0x63, 0x68, 0x61, 0x72, 0x74, 0x12,
	0x2c, 0x0a, 0x0d, 0x63, 0x68, 0x61, 0x72, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x0c, 0x63, 0x68, 0x61, 0x72, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x11, 0x0a,
	0x0a, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x12, 0x03, 0xf8, 0x42, 0x01,
	0x22, 0xb5, 0x02, 0x0a, 0x13, 0x48, 0x65, 0x6c, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x4f,
	0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x07, 0x69, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x48, 0x00, 0x52, 0x07, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x64, 0x12, 0x3b, 0x0a,
	0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x50, 0x0a, 0x03, 0x6d, 0x61,
	0x70, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61,
	0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x48, 0x65, 0x6c, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x4f, 0x76, 0x65, 0x72, 0x72,
	0x69, 0x64, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x9a, 0x01, 0x02, 0x18, 0x01, 0x52, 0x03, 0x6d, 0x61, 0x70, 0x1a, 0x58, 0x0a, 0x08,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x36, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69,
	0x64, 0x65, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x22, 0x85, 0x02, 0x0a, 0x0a, 0x48, 0x65, 0x6c,
	0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x41, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61,
	0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6d, 0x43, 0x68, 0x61, 0x72, 0x74,
	0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x56, 0x0a, 0x10, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x48, 0x65,
	0x6c, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65,
	0x73, 0x52, 0x0f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64,
	0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18,
	0x00, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x42, 0x12, 0x0a, 0x0b,
	0x63, 0x68, 0x61, 0x72, 0x74, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x12, 0x03, 0xf8, 0x42, 0x01,
	0x42, 0x52, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61,
	0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76,
	0x61, 0x6e, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prodvana_common_config_helm_proto_rawDescOnce sync.Once
	file_prodvana_common_config_helm_proto_rawDescData = file_prodvana_common_config_helm_proto_rawDesc
)

func file_prodvana_common_config_helm_proto_rawDescGZIP() []byte {
	file_prodvana_common_config_helm_proto_rawDescOnce.Do(func() {
		file_prodvana_common_config_helm_proto_rawDescData = protoimpl.X.CompressGZIP(file_prodvana_common_config_helm_proto_rawDescData)
	})
	return file_prodvana_common_config_helm_proto_rawDescData
}

var file_prodvana_common_config_helm_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_prodvana_common_config_helm_proto_goTypes = []interface{}{
	(*RemoteHelmChart)(nil),     // 0: prodvana.common_config.RemoteHelmChart
	(*HelmValuesOverrides)(nil), // 1: prodvana.common_config.HelmValuesOverrides
	(*HelmConfig)(nil),          // 2: prodvana.common_config.HelmConfig
	nil,                         // 3: prodvana.common_config.HelmValuesOverrides.MapEntry
	(*LocalConfig)(nil),         // 4: prodvana.common_config.LocalConfig
	(*EnvValue)(nil),            // 5: prodvana.common_config.EnvValue
}
var file_prodvana_common_config_helm_proto_depIdxs = []int32{
	4, // 0: prodvana.common_config.HelmValuesOverrides.local:type_name -> prodvana.common_config.LocalConfig
	3, // 1: prodvana.common_config.HelmValuesOverrides.map:type_name -> prodvana.common_config.HelmValuesOverrides.MapEntry
	0, // 2: prodvana.common_config.HelmConfig.remote:type_name -> prodvana.common_config.RemoteHelmChart
	1, // 3: prodvana.common_config.HelmConfig.values_overrides:type_name -> prodvana.common_config.HelmValuesOverrides
	5, // 4: prodvana.common_config.HelmValuesOverrides.MapEntry.value:type_name -> prodvana.common_config.EnvValue
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_prodvana_common_config_helm_proto_init() }
func file_prodvana_common_config_helm_proto_init() {
	if File_prodvana_common_config_helm_proto != nil {
		return
	}
	file_prodvana_common_config_env_proto_init()
	file_prodvana_common_config_kubernetes_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_prodvana_common_config_helm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteHelmChart); i {
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
		file_prodvana_common_config_helm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelmValuesOverrides); i {
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
		file_prodvana_common_config_helm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelmConfig); i {
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
	file_prodvana_common_config_helm_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RemoteHelmChart_Repo)(nil),
	}
	file_prodvana_common_config_helm_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*HelmValuesOverrides_Inlined)(nil),
		(*HelmValuesOverrides_Local)(nil),
	}
	file_prodvana_common_config_helm_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*HelmConfig_Remote)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_prodvana_common_config_helm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prodvana_common_config_helm_proto_goTypes,
		DependencyIndexes: file_prodvana_common_config_helm_proto_depIdxs,
		MessageInfos:      file_prodvana_common_config_helm_proto_msgTypes,
	}.Build()
	File_prodvana_common_config_helm_proto = out.File
	file_prodvana_common_config_helm_proto_rawDesc = nil
	file_prodvana_common_config_helm_proto_goTypes = nil
	file_prodvana_common_config_helm_proto_depIdxs = nil
}
