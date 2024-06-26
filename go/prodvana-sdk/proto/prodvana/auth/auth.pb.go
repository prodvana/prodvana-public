// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.10
// source: prodvana/auth/auth.proto

package auth

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AuthToken) Reset() {
	*x = AuthToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_auth_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthToken) ProtoMessage() {}

func (x *AuthToken) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_auth_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthToken.ProtoReflect.Descriptor instead.
func (*AuthToken) Descriptor() ([]byte, []int) {
	return file_prodvana_auth_auth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ApiTokenInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name              string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description       string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ExpiresTimestamp  *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=expires_timestamp,json=expiresTimestamp,proto3" json:"expires_timestamp,omitempty"`
	CreationTimestamp *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=creation_timestamp,json=creationTimestamp,proto3" json:"creation_timestamp,omitempty"`
}

func (x *ApiTokenInfo) Reset() {
	*x = ApiTokenInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_auth_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiTokenInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiTokenInfo) ProtoMessage() {}

func (x *ApiTokenInfo) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_auth_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiTokenInfo.ProtoReflect.Descriptor instead.
func (*ApiTokenInfo) Descriptor() ([]byte, []int) {
	return file_prodvana_auth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *ApiTokenInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ApiTokenInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ApiTokenInfo) GetExpiresTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresTimestamp
	}
	return nil
}

func (x *ApiTokenInfo) GetCreationTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.CreationTimestamp
	}
	return nil
}

type AuthContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthToken *AuthToken `protobuf:"bytes,1,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"`
	ApiToken  string     `protobuf:"bytes,2,opt,name=api_token,json=apiToken,proto3" json:"api_token,omitempty"`
	Addr      string     `protobuf:"bytes,3,opt,name=addr,proto3" json:"addr,omitempty"`
	SkipTls   bool       `protobuf:"varint,4,opt,name=skip_tls,json=skipTls,proto3" json:"skip_tls,omitempty"`
}

func (x *AuthContext) Reset() {
	*x = AuthContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_auth_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthContext) ProtoMessage() {}

func (x *AuthContext) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_auth_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthContext.ProtoReflect.Descriptor instead.
func (*AuthContext) Descriptor() ([]byte, []int) {
	return file_prodvana_auth_auth_proto_rawDescGZIP(), []int{2}
}

func (x *AuthContext) GetAuthToken() *AuthToken {
	if x != nil {
		return x.AuthToken
	}
	return nil
}

func (x *AuthContext) GetApiToken() string {
	if x != nil {
		return x.ApiToken
	}
	return ""
}

func (x *AuthContext) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *AuthContext) GetSkipTls() bool {
	if x != nil {
		return x.SkipTls
	}
	return false
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contexts            map[string]*AuthContext `protobuf:"bytes,3,rep,name=contexts,proto3" json:"contexts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CurrentContext      string                  `protobuf:"bytes,4,opt,name=current_context,json=currentContext,proto3" json:"current_context,omitempty"`
	AdminContexts       map[string]*AuthContext `protobuf:"bytes,5,rep,name=admin_contexts,json=adminContexts,proto3" json:"admin_contexts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CurrentAdminContext string                  `protobuf:"bytes,6,opt,name=current_admin_context,json=currentAdminContext,proto3" json:"current_admin_context,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodvana_auth_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_prodvana_auth_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_prodvana_auth_auth_proto_rawDescGZIP(), []int{3}
}

func (x *Session) GetContexts() map[string]*AuthContext {
	if x != nil {
		return x.Contexts
	}
	return nil
}

func (x *Session) GetCurrentContext() string {
	if x != nil {
		return x.CurrentContext
	}
	return ""
}

func (x *Session) GetAdminContexts() map[string]*AuthContext {
	if x != nil {
		return x.AdminContexts
	}
	return nil
}

func (x *Session) GetCurrentAdminContext() string {
	if x != nil {
		return x.CurrentAdminContext
	}
	return ""
}

var File_prodvana_auth_auth_proto protoreflect.FileDescriptor

var file_prodvana_auth_auth_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f, 0x64,
	0x76, 0x61, 0x6e, 0x61, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x09, 0x41, 0x75,
	0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xd8, 0x01,
	0x0a, 0x0c, 0x41, 0x70, 0x69, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x11, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x49, 0x0a,
	0x12, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x92, 0x01, 0x0a, 0x0b, 0x41, 0x75, 0x74,
	0x68, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x37, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x70, 0x69, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x69, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64,
	0x64, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x74, 0x6c, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x70, 0x54, 0x6c, 0x73, 0x22, 0xd4, 0x03,
	0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x08, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x50, 0x0a, 0x0e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x1a, 0x57, 0x0a, 0x0d, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x30, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x5c, 0x0a, 0x12, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x76, 0x61, 0x6e, 0x61, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x52, 0x0a, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x09, 0x61, 0x70, 0x69, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x64,
	0x76, 0x61, 0x6e, 0x61, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x76, 0x61, 0x6e, 0x61, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prodvana_auth_auth_proto_rawDescOnce sync.Once
	file_prodvana_auth_auth_proto_rawDescData = file_prodvana_auth_auth_proto_rawDesc
)

func file_prodvana_auth_auth_proto_rawDescGZIP() []byte {
	file_prodvana_auth_auth_proto_rawDescOnce.Do(func() {
		file_prodvana_auth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_prodvana_auth_auth_proto_rawDescData)
	})
	return file_prodvana_auth_auth_proto_rawDescData
}

var file_prodvana_auth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_prodvana_auth_auth_proto_goTypes = []interface{}{
	(*AuthToken)(nil),             // 0: prodvana.auth.AuthToken
	(*ApiTokenInfo)(nil),          // 1: prodvana.auth.ApiTokenInfo
	(*AuthContext)(nil),           // 2: prodvana.auth.AuthContext
	(*Session)(nil),               // 3: prodvana.auth.Session
	nil,                           // 4: prodvana.auth.Session.ContextsEntry
	nil,                           // 5: prodvana.auth.Session.AdminContextsEntry
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_prodvana_auth_auth_proto_depIdxs = []int32{
	6, // 0: prodvana.auth.ApiTokenInfo.expires_timestamp:type_name -> google.protobuf.Timestamp
	6, // 1: prodvana.auth.ApiTokenInfo.creation_timestamp:type_name -> google.protobuf.Timestamp
	0, // 2: prodvana.auth.AuthContext.auth_token:type_name -> prodvana.auth.AuthToken
	4, // 3: prodvana.auth.Session.contexts:type_name -> prodvana.auth.Session.ContextsEntry
	5, // 4: prodvana.auth.Session.admin_contexts:type_name -> prodvana.auth.Session.AdminContextsEntry
	2, // 5: prodvana.auth.Session.ContextsEntry.value:type_name -> prodvana.auth.AuthContext
	2, // 6: prodvana.auth.Session.AdminContextsEntry.value:type_name -> prodvana.auth.AuthContext
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_prodvana_auth_auth_proto_init() }
func file_prodvana_auth_auth_proto_init() {
	if File_prodvana_auth_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prodvana_auth_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthToken); i {
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
		file_prodvana_auth_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiTokenInfo); i {
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
		file_prodvana_auth_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthContext); i {
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
		file_prodvana_auth_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
			RawDescriptor: file_prodvana_auth_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prodvana_auth_auth_proto_goTypes,
		DependencyIndexes: file_prodvana_auth_auth_proto_depIdxs,
		MessageInfos:      file_prodvana_auth_auth_proto_msgTypes,
	}.Build()
	File_prodvana_auth_auth_proto = out.File
	file_prodvana_auth_auth_proto_rawDesc = nil
	file_prodvana_auth_auth_proto_goTypes = nil
	file_prodvana_auth_auth_proto_depIdxs = nil
}
