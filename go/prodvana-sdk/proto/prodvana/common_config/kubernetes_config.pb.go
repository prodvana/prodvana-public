// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.10
// source: prodvana/common_config/kubernetes_config.proto

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

type KubernetesConfig_Type int32

const (
	KubernetesConfig_UNKNOWN    KubernetesConfig_Type = 0
	KubernetesConfig_KUBERNETES KubernetesConfig_Type = 1
	KubernetesConfig_KUSTOMIZE  KubernetesConfig_Type = 2
)

// Enum value maps for KubernetesConfig_Type.
var (
	KubernetesConfig_Type_name = map[int32]string{
		0: "UNKNOWN",
		1: "KUBERNETES",
		2: "KUSTOMIZE",
	}
	KubernetesConfig_Type_value = map[string]int32{
		"UNKNOWN":    0,
		"KUBERNETES": 1,
		"KUSTOMIZE":  2,
	}
)

func (x KubernetesConfig_Type) Enum() *KubernetesConfig_Type {
	p := new(KubernetesConfig_Type)
	*p = x
	return p
}

func (x KubernetesConfig_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KubernetesConfig_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_prodvana_common_config_kubernetes_config_proto_enumTypes[0].Descriptor()
}

func (KubernetesConfig_Type) Type() protoreflect.EnumType {
	return &file_prodvana_common_config_kubernetes_config_proto_enumTypes[0]
}

func (x KubernetesConfig_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KubernetesConfig_Type.Descriptor instead.
func (KubernetesConfig_Type) EnumDescriptor() ([]byte, []int) {
	return file_prodvana_common_config_kubernetes_config_proto_rawDescGZIP(), []int{2, 0}
}

type KubernetesConfig_EnvInjectionMode int32

const (
	KubernetesConfig_ENV_INJECT_UNKNOWN KubernetesConfig_EnvInjectionMode = 0
	// disables env injection entirely
	KubernetesConfig_ENV_INJECT_DISABLED KubernetesConfig_EnvInjectionMode = 1
	// inject non-secret env values from the Release Channel
	KubernetesConfig_ENV_INJECT_NON_SECRET_ENV KubernetesConfig_EnvInjectionMode = 2
)

// Enum value maps for KubernetesConfig_EnvInjectionMode.
var (
	KubernetesConfig_EnvInjectionMode_name = map[int32]string{
		0: "ENV_INJECT_UNKNOWN",
		1: "ENV_INJECT_DISABLED",
		2: "ENV_INJECT_NON_SECRET_ENV",
	}
	KubernetesConfig_EnvInjectionMode_value = map[string]int32{
		"ENV_INJECT_UNKNOWN":        0,
		"ENV_INJECT_DISABLED":       1,
		"ENV_INJECT_NON_SECRET_ENV": 2,
	}
)

func (x KubernetesConfig_EnvInjectionMode) Enum() *KubernetesConfig_EnvInjectionMode {
	p := new(KubernetesConfig_EnvInjectionMode)
	*p = x
	return p
}

func (x KubernetesConfig_EnvInjectionMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KubernetesConfig_EnvInjectionMode) Descriptor() protoreflect.EnumDescriptor {
	return file_prodvana_common_config_kubernetes_config_proto_enumTypes[1].Descriptor()
}

func (KubernetesConfig_EnvInjectionMode) Type() protoreflect.EnumType {
	return &file_prodvana_common_config_kubernetes_config_proto_enumTypes[1]
}

func (x KubernetesConfig_EnvInjectionMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KubernetesConfig_EnvInjectionMode.Descriptor instead.
func (KubernetesConfig_EnvInjectionMode) EnumDescriptor() ([]byte, []int) {
	return file_prodvana_common_config_kubernetes_config_proto_rawDescGZIP(), []int{2, 1}
}

type LocalConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to PathOneof:
	//
	//	*LocalConfig_Path
	PathOneof isLocalConfig_PathOneof `protobuf_oneof:"path_oneof"`
	// cannot be set in conjunction with path or tarball_blob_id
	// Specify multiple paths. They will get concatenated.
	Paths []string `protobuf:"bytes,2,rep,name=paths,proto3" json:"paths,omitempty"`
	// A sub_path relative to path containing the actual config.
	// sub_path value can be templated, e.g. `{{.Builtins.ReleaseChannel.Name}}`.
	// if specified, path must also be specified and be a directory.
	SubPath string `protobuf:"bytes,3,opt,name=sub_path,json=subPath,proto3" json:"sub_path,omitempty"`
	// Used in conjunction with sub_path to exclude sub paths from being tarball'ed and uploaded
	// to Prodvana.
	// Follows the same format as gitignore.
	ExcludePatterns []string `protobuf:"bytes,5,rep,name=exclude_patterns,json=excludePatterns,proto3" json:"exclude_patterns,omitempty"`
	// If set, only include files that match this pattern. If not set, include all files not explicitly excluded.
	IncludePatterns []string `protobuf:"bytes,6,rep,name=include_patterns,json=includePatterns,proto3" json:"include_patterns,omitempty"`
}

func (x *LocalConfig) Reset() {
	*x = LocalConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_common_config_kubernetes_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalConfig) ProtoMessage() {}

func (x *LocalConfig) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_common_config_kubernetes_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalConfig.ProtoReflect.Descriptor instead.
func (*LocalConfig) Descriptor() ([]byte, []int) {
	return file_prodvana_common_config_kubernetes_config_proto_rawDescGZIP(), []int{0}
}

func (m *LocalConfig) GetPathOneof() isLocalConfig_PathOneof {
	if m != nil {
		return m.PathOneof
	}
	return nil
}

func (x *LocalConfig) GetPath() string {
	if x, ok := x.GetPathOneof().(*LocalConfig_Path); ok {
		return x.Path
	}
	return ""
}

func (x *LocalConfig) GetPaths() []string {
	if x != nil {
		return x.Paths
	}
	return nil
}

func (x *LocalConfig) GetSubPath() string {
	if x != nil {
		return x.SubPath
	}
	return ""
}

func (x *LocalConfig) GetExcludePatterns() []string {
	if x != nil {
		return x.ExcludePatterns
	}
	return nil
}

func (x *LocalConfig) GetIncludePatterns() []string {
	if x != nil {
		return x.IncludePatterns
	}
	return nil
}

type isLocalConfig_PathOneof interface {
	isLocalConfig_PathOneof()
}

type LocalConfig_Path struct {
	// Specify a path to a local file or directory
	Path string `protobuf:"bytes,1,opt,name=path,proto3,oneof"`
}

func (*LocalConfig_Path) isLocalConfig_PathOneof() {}

// a remote config, specified by a remote source and a subpath
type RemoteConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to RemoteOneof:
	//
	//	*RemoteConfig_TarballBlobId
	RemoteOneof isRemoteConfig_RemoteOneof `protobuf_oneof:"remote_oneof"`
	SubPath     string                     `protobuf:"bytes,2,opt,name=sub_path,json=subPath,proto3" json:"sub_path,omitempty"`
}

func (x *RemoteConfig) Reset() {
	*x = RemoteConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_common_config_kubernetes_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteConfig) ProtoMessage() {}

func (x *RemoteConfig) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_common_config_kubernetes_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteConfig.ProtoReflect.Descriptor instead.
func (*RemoteConfig) Descriptor() ([]byte, []int) {
	return file_prodvana_common_config_kubernetes_config_proto_rawDescGZIP(), []int{1}
}

func (m *RemoteConfig) GetRemoteOneof() isRemoteConfig_RemoteOneof {
	if m != nil {
		return m.RemoteOneof
	}
	return nil
}

func (x *RemoteConfig) GetTarballBlobId() string {
	if x, ok := x.GetRemoteOneof().(*RemoteConfig_TarballBlobId); ok {
		return x.TarballBlobId
	}
	return ""
}

func (x *RemoteConfig) GetSubPath() string {
	if x != nil {
		return x.SubPath
	}
	return ""
}

type isRemoteConfig_RemoteOneof interface {
	isRemoteConfig_RemoteOneof()
}

type RemoteConfig_TarballBlobId struct {
	// A directory tarball blob id uploaded to prodvana.
	TarballBlobId string `protobuf:"bytes,1,opt,name=tarball_blob_id,json=tarballBlobId,proto3,oneof"` // TODO(naphat) git repo support?
}

func (*RemoteConfig_TarballBlobId) isRemoteConfig_RemoteOneof() {}

type KubernetesConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type KubernetesConfig_Type `protobuf:"varint,1,opt,name=type,proto3,enum=prodvana.common_config.KubernetesConfig_Type" json:"type,omitempty"`
	// Types that are assignable to SourceOneof:
	//
	//	*KubernetesConfig_Inlined
	//	*KubernetesConfig_Local
	//	*KubernetesConfig_Remote
	SourceOneof isKubernetesConfig_SourceOneof `protobuf_oneof:"source_oneof"`
	// Defaults to ENV_INJECT_NON_SECRET_ENV
	EnvInjectionMode KubernetesConfig_EnvInjectionMode `protobuf:"varint,4,opt,name=env_injection_mode,json=envInjectionMode,proto3,enum=prodvana.common_config.KubernetesConfig_EnvInjectionMode" json:"env_injection_mode,omitempty"`
	Patches          []*KubernetesPatch                `protobuf:"bytes,6,rep,name=patches,proto3" json:"patches,omitempty"`
}

func (x *KubernetesConfig) Reset() {
	*x = KubernetesConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_common_config_kubernetes_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesConfig) ProtoMessage() {}

func (x *KubernetesConfig) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_common_config_kubernetes_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesConfig.ProtoReflect.Descriptor instead.
func (*KubernetesConfig) Descriptor() ([]byte, []int) {
	return file_prodvana_common_config_kubernetes_config_proto_rawDescGZIP(), []int{2}
}

func (x *KubernetesConfig) GetType() KubernetesConfig_Type {
	if x != nil {
		return x.Type
	}
	return KubernetesConfig_UNKNOWN
}

func (m *KubernetesConfig) GetSourceOneof() isKubernetesConfig_SourceOneof {
	if m != nil {
		return m.SourceOneof
	}
	return nil
}

func (x *KubernetesConfig) GetInlined() string {
	if x, ok := x.GetSourceOneof().(*KubernetesConfig_Inlined); ok {
		return x.Inlined
	}
	return ""
}

func (x *KubernetesConfig) GetLocal() *LocalConfig {
	if x, ok := x.GetSourceOneof().(*KubernetesConfig_Local); ok {
		return x.Local
	}
	return nil
}

func (x *KubernetesConfig) GetRemote() *RemoteConfig {
	if x, ok := x.GetSourceOneof().(*KubernetesConfig_Remote); ok {
		return x.Remote
	}
	return nil
}

func (x *KubernetesConfig) GetEnvInjectionMode() KubernetesConfig_EnvInjectionMode {
	if x != nil {
		return x.EnvInjectionMode
	}
	return KubernetesConfig_ENV_INJECT_UNKNOWN
}

func (x *KubernetesConfig) GetPatches() []*KubernetesPatch {
	if x != nil {
		return x.Patches
	}
	return nil
}

type isKubernetesConfig_SourceOneof interface {
	isKubernetesConfig_SourceOneof()
}

type KubernetesConfig_Inlined struct {
	Inlined string `protobuf:"bytes,2,opt,name=inlined,proto3,oneof"`
}

type KubernetesConfig_Local struct {
	Local *LocalConfig `protobuf:"bytes,3,opt,name=local,proto3,oneof"`
}

type KubernetesConfig_Remote struct {
	Remote *RemoteConfig `protobuf:"bytes,5,opt,name=remote,proto3,oneof"`
}

func (*KubernetesConfig_Inlined) isKubernetesConfig_SourceOneof() {}

func (*KubernetesConfig_Local) isKubernetesConfig_SourceOneof() {}

func (*KubernetesConfig_Remote) isKubernetesConfig_SourceOneof() {}

type KubernetesPatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target *KubernetesPatch_Target `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"` // filter to a specific Kubernetes object. Omit to apply to all objects.
	// Types that are assignable to PatchOneof:
	//
	//	*KubernetesPatch_Replace_
	PatchOneof isKubernetesPatch_PatchOneof `protobuf_oneof:"patch_oneof"`
}

func (x *KubernetesPatch) Reset() {
	*x = KubernetesPatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_common_config_kubernetes_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesPatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesPatch) ProtoMessage() {}

func (x *KubernetesPatch) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_common_config_kubernetes_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesPatch.ProtoReflect.Descriptor instead.
func (*KubernetesPatch) Descriptor() ([]byte, []int) {
	return file_prodvana_common_config_kubernetes_config_proto_rawDescGZIP(), []int{3}
}

func (x *KubernetesPatch) GetTarget() *KubernetesPatch_Target {
	if x != nil {
		return x.Target
	}
	return nil
}

func (m *KubernetesPatch) GetPatchOneof() isKubernetesPatch_PatchOneof {
	if m != nil {
		return m.PatchOneof
	}
	return nil
}

func (x *KubernetesPatch) GetReplace() *KubernetesPatch_Replace {
	if x, ok := x.GetPatchOneof().(*KubernetesPatch_Replace_); ok {
		return x.Replace
	}
	return nil
}

type isKubernetesPatch_PatchOneof interface {
	isKubernetesPatch_PatchOneof()
}

type KubernetesPatch_Replace_ struct {
	Replace *KubernetesPatch_Replace `protobuf:"bytes,2,opt,name=replace,proto3,oneof"`
}

func (*KubernetesPatch_Replace_) isKubernetesPatch_PatchOneof() {}

type KubernetesPatch_Target struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// omit any of these fields to match all values
	Group     string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Version   string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Kind      string `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
	Name      string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,5,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *KubernetesPatch_Target) Reset() {
	*x = KubernetesPatch_Target{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_common_config_kubernetes_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesPatch_Target) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesPatch_Target) ProtoMessage() {}

func (x *KubernetesPatch_Target) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_common_config_kubernetes_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesPatch_Target.ProtoReflect.Descriptor instead.
func (*KubernetesPatch_Target) Descriptor() ([]byte, []int) {
	return file_prodvana_common_config_kubernetes_config_proto_rawDescGZIP(), []int{3, 0}
}

func (x *KubernetesPatch_Target) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *KubernetesPatch_Target) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *KubernetesPatch_Target) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *KubernetesPatch_Target) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KubernetesPatch_Target) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type KubernetesPatch_Replace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"` // json6902 path
	// Types that are assignable to ValueOneof:
	//
	//	*KubernetesPatch_Replace_String_
	//	*KubernetesPatch_Replace_IntAsString
	ValueOneof isKubernetesPatch_Replace_ValueOneof `protobuf_oneof:"value_oneof"`
}

func (x *KubernetesPatch_Replace) Reset() {
	*x = KubernetesPatch_Replace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_common_config_kubernetes_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesPatch_Replace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesPatch_Replace) ProtoMessage() {}

func (x *KubernetesPatch_Replace) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_common_config_kubernetes_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesPatch_Replace.ProtoReflect.Descriptor instead.
func (*KubernetesPatch_Replace) Descriptor() ([]byte, []int) {
	return file_prodvana_common_config_kubernetes_config_proto_rawDescGZIP(), []int{3, 1}
}

func (x *KubernetesPatch_Replace) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (m *KubernetesPatch_Replace) GetValueOneof() isKubernetesPatch_Replace_ValueOneof {
	if m != nil {
		return m.ValueOneof
	}
	return nil
}

func (x *KubernetesPatch_Replace) GetString_() string {
	if x, ok := x.GetValueOneof().(*KubernetesPatch_Replace_String_); ok {
		return x.String_
	}
	return ""
}

func (x *KubernetesPatch_Replace) GetIntAsString() string {
	if x, ok := x.GetValueOneof().(*KubernetesPatch_Replace_IntAsString); ok {
		return x.IntAsString
	}
	return ""
}

type isKubernetesPatch_Replace_ValueOneof interface {
	isKubernetesPatch_Replace_ValueOneof()
}

type KubernetesPatch_Replace_String_ struct {
	String_ string `protobuf:"bytes,2,opt,name=string,proto3,oneof"`
}

type KubernetesPatch_Replace_IntAsString struct {
	IntAsString string `protobuf:"bytes,3,opt,name=int_as_string,json=intAsString,proto3,oneof"`
}

func (*KubernetesPatch_Replace_String_) isKubernetesPatch_Replace_ValueOneof() {}

func (*KubernetesPatch_Replace_IntAsString) isKubernetesPatch_Replace_ValueOneof() {}

var File_prodvana_common_config_kubernetes_config_proto protoreflect.FileDescriptor

var file_prodvana_common_config_kubernetes_config_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x16, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xcf, 0x01, 0x0a, 0x0b, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x14, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x75, 0x62, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x75, 0x62, 0x50, 0x61, 0x74, 0x68, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x78, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x50, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x6e, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x70,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x69,
	0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x42, 0x0c,
	0x0a, 0x0a, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x4a, 0x04, 0x08, 0x04,
	0x10, 0x05, 0x52, 0x0f, 0x74, 0x61, 0x72, 0x62, 0x61, 0x6c, 0x6c, 0x5f, 0x62, 0x6c, 0x6f, 0x62,
	0x5f, 0x69, 0x64, 0x22, 0x63, 0x0a, 0x0c, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x28, 0x0a, 0x0f, 0x74, 0x61, 0x72, 0x62, 0x61, 0x6c, 0x6c, 0x5f, 0x62,
	0x6c, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d,
	0x74, 0x61, 0x72, 0x62, 0x61, 0x6c, 0x6c, 0x42, 0x6c, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x75, 0x62, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x75, 0x62, 0x50, 0x61, 0x74, 0x68, 0x42, 0x0e, 0x0a, 0x0c, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x22, 0xe9, 0x04, 0x0a, 0x10, 0x4b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4b, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82,
	0x01, 0x02, 0x20, 0x00, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x07, 0x69, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x07, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x64, 0x12,
	0x3b, 0x0a, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x3e, 0x0a, 0x06,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x67, 0x0a, 0x12,
	0x65, 0x6e, 0x76, 0x5f, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x39, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76,
	0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x45, 0x6e, 0x76, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x6f, 0x64, 0x65, 0x52, 0x10, 0x65, 0x6e, 0x76, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x50, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e,
	0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x50, 0x61, 0x74, 0x63, 0x68, 0x42,
	0x0d, 0xfa, 0x42, 0x0a, 0x92, 0x01, 0x07, 0x22, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x07,
	0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x22, 0x32, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a,
	0x4b, 0x55, 0x42, 0x45, 0x52, 0x4e, 0x45, 0x54, 0x45, 0x53, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09,
	0x4b, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x49, 0x5a, 0x45, 0x10, 0x02, 0x22, 0x62, 0x0a, 0x10, 0x45,
	0x6e, 0x76, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12,
	0x16, 0x0a, 0x12, 0x45, 0x4e, 0x56, 0x5f, 0x49, 0x4e, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x4e, 0x56, 0x5f, 0x49,
	0x4e, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01,
	0x12, 0x1d, 0x0a, 0x19, 0x45, 0x4e, 0x56, 0x5f, 0x49, 0x4e, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x4e,
	0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54, 0x5f, 0x45, 0x4e, 0x56, 0x10, 0x02, 0x42,
	0x13, 0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x12,
	0x03, 0xf8, 0x42, 0x01, 0x22, 0xb7, 0x03, 0x0a, 0x0f, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x50, 0x61, 0x74, 0x63, 0x68, 0x12, 0x46, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76,
	0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x50, 0x61, 0x74, 0x63,
	0x68, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x12, 0x55, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x50, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x07,
	0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x1a, 0x7e, 0x0a, 0x06, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x1a, 0x71, 0x0a, 0x07, 0x52, 0x65, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x12, 0x24, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x5f, 0x61, 0x73, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x41, 0x73,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x12, 0x0a, 0x0b, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f,
	0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x42, 0x12, 0x0a, 0x0b, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x42, 0x52,
	0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2d, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e,
	0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64,
	0x76, 0x61, 0x6e, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prodvana_common_config_kubernetes_config_proto_rawDescOnce sync.Once
	file_prodvana_common_config_kubernetes_config_proto_rawDescData = file_prodvana_common_config_kubernetes_config_proto_rawDesc
)

func file_prodvana_common_config_kubernetes_config_proto_rawDescGZIP() []byte {
	file_prodvana_common_config_kubernetes_config_proto_rawDescOnce.Do(func() {
		file_prodvana_common_config_kubernetes_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_prodvana_common_config_kubernetes_config_proto_rawDescData)
	})
	return file_prodvana_common_config_kubernetes_config_proto_rawDescData
}

var file_prodvana_common_config_kubernetes_config_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_prodvana_common_config_kubernetes_config_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_prodvana_common_config_kubernetes_config_proto_goTypes = []interface{}{
	(KubernetesConfig_Type)(0),             // 0: prodvana.common_config.KubernetesConfig.Type
	(KubernetesConfig_EnvInjectionMode)(0), // 1: prodvana.common_config.KubernetesConfig.EnvInjectionMode
	(*LocalConfig)(nil),                    // 2: prodvana.common_config.LocalConfig
	(*RemoteConfig)(nil),                   // 3: prodvana.common_config.RemoteConfig
	(*KubernetesConfig)(nil),               // 4: prodvana.common_config.KubernetesConfig
	(*KubernetesPatch)(nil),                // 5: prodvana.common_config.KubernetesPatch
	(*KubernetesPatch_Target)(nil),         // 6: prodvana.common_config.KubernetesPatch.Target
	(*KubernetesPatch_Replace)(nil),        // 7: prodvana.common_config.KubernetesPatch.Replace
}
var file_prodvana_common_config_kubernetes_config_proto_depIdxs = []int32{
	0, // 0: prodvana.common_config.KubernetesConfig.type:type_name -> prodvana.common_config.KubernetesConfig.Type
	2, // 1: prodvana.common_config.KubernetesConfig.local:type_name -> prodvana.common_config.LocalConfig
	3, // 2: prodvana.common_config.KubernetesConfig.remote:type_name -> prodvana.common_config.RemoteConfig
	1, // 3: prodvana.common_config.KubernetesConfig.env_injection_mode:type_name -> prodvana.common_config.KubernetesConfig.EnvInjectionMode
	5, // 4: prodvana.common_config.KubernetesConfig.patches:type_name -> prodvana.common_config.KubernetesPatch
	6, // 5: prodvana.common_config.KubernetesPatch.target:type_name -> prodvana.common_config.KubernetesPatch.Target
	7, // 6: prodvana.common_config.KubernetesPatch.replace:type_name -> prodvana.common_config.KubernetesPatch.Replace
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_prodvana_common_config_kubernetes_config_proto_init() }
func file_prodvana_common_config_kubernetes_config_proto_init() {
	if File_prodvana_common_config_kubernetes_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prodvana_common_config_kubernetes_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalConfig); i {
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
		file_prodvana_common_config_kubernetes_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteConfig); i {
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
		file_prodvana_common_config_kubernetes_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesConfig); i {
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
		file_prodvana_common_config_kubernetes_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesPatch); i {
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
		file_prodvana_common_config_kubernetes_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesPatch_Target); i {
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
		file_prodvana_common_config_kubernetes_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesPatch_Replace); i {
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
	file_prodvana_common_config_kubernetes_config_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*LocalConfig_Path)(nil),
	}
	file_prodvana_common_config_kubernetes_config_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*RemoteConfig_TarballBlobId)(nil),
	}
	file_prodvana_common_config_kubernetes_config_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*KubernetesConfig_Inlined)(nil),
		(*KubernetesConfig_Local)(nil),
		(*KubernetesConfig_Remote)(nil),
	}
	file_prodvana_common_config_kubernetes_config_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*KubernetesPatch_Replace_)(nil),
	}
	file_prodvana_common_config_kubernetes_config_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*KubernetesPatch_Replace_String_)(nil),
		(*KubernetesPatch_Replace_IntAsString)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_prodvana_common_config_kubernetes_config_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prodvana_common_config_kubernetes_config_proto_goTypes,
		DependencyIndexes: file_prodvana_common_config_kubernetes_config_proto_depIdxs,
		EnumInfos:         file_prodvana_common_config_kubernetes_config_proto_enumTypes,
		MessageInfos:      file_prodvana_common_config_kubernetes_config_proto_msgTypes,
	}.Build()
	File_prodvana_common_config_kubernetes_config_proto = out.File
	file_prodvana_common_config_kubernetes_config_proto_rawDesc = nil
	file_prodvana_common_config_kubernetes_config_proto_goTypes = nil
	file_prodvana_common_config_kubernetes_config_proto_depIdxs = nil
}
