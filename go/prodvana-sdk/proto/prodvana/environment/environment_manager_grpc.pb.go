// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.10
// source: prodvana/environment/environment_manager.proto

package environment

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	EnvironmentManager_GetClusterAgentApiToken_FullMethodName  = "/prodvana.environment.EnvironmentManager/GetClusterAgentApiToken"
	EnvironmentManager_LinkCluster_FullMethodName              = "/prodvana.environment.EnvironmentManager/LinkCluster"
	EnvironmentManager_ListClusters_FullMethodName             = "/prodvana.environment.EnvironmentManager/ListClusters"
	EnvironmentManager_GetCluster_FullMethodName               = "/prodvana.environment.EnvironmentManager/GetCluster"
	EnvironmentManager_RemoveCluster_FullMethodName            = "/prodvana.environment.EnvironmentManager/RemoveCluster"
	EnvironmentManager_GetClusterAuth_FullMethodName           = "/prodvana.environment.EnvironmentManager/GetClusterAuth"
	EnvironmentManager_GetClusterConfig_FullMethodName         = "/prodvana.environment.EnvironmentManager/GetClusterConfig"
	EnvironmentManager_DetectClusterConfig_FullMethodName      = "/prodvana.environment.EnvironmentManager/DetectClusterConfig"
	EnvironmentManager_ConfigureCluster_FullMethodName         = "/prodvana.environment.EnvironmentManager/ConfigureCluster"
	EnvironmentManager_ValidateConfigureCluster_FullMethodName = "/prodvana.environment.EnvironmentManager/ValidateConfigureCluster"
	EnvironmentManager_GetClusterStatus_FullMethodName         = "/prodvana.environment.EnvironmentManager/GetClusterStatus"
)

// EnvironmentManagerClient is the client API for EnvironmentManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EnvironmentManagerClient interface {
	GetClusterAgentApiToken(ctx context.Context, in *GetClusterAgentApiTokenReq, opts ...grpc.CallOption) (*GetClusterAgentApiTokenResp, error)
	LinkCluster(ctx context.Context, in *LinkClusterReq, opts ...grpc.CallOption) (*LinkClusterResp, error)
	ListClusters(ctx context.Context, in *ListClustersReq, opts ...grpc.CallOption) (*ListClustersResp, error)
	GetCluster(ctx context.Context, in *GetClusterReq, opts ...grpc.CallOption) (*GetClusterResp, error)
	RemoveCluster(ctx context.Context, in *RemoveClusterReq, opts ...grpc.CallOption) (*RemoveClusterResp, error)
	// Deprecated.
	GetClusterAuth(ctx context.Context, in *GetClusterAuthReq, opts ...grpc.CallOption) (*GetClusterAuthResp, error)
	GetClusterConfig(ctx context.Context, in *GetClusterConfigReq, opts ...grpc.CallOption) (*GetClusterConfigResp, error)
	DetectClusterConfig(ctx context.Context, in *DetectClusterConfigReq, opts ...grpc.CallOption) (*DetectClusterConfigResp, error)
	ConfigureCluster(ctx context.Context, in *ConfigureClusterReq, opts ...grpc.CallOption) (*ConfigureClusterResp, error)
	ValidateConfigureCluster(ctx context.Context, in *ConfigureClusterReq, opts ...grpc.CallOption) (*ValidateConfigureClusterResp, error)
	GetClusterStatus(ctx context.Context, in *GetClusterStatusReq, opts ...grpc.CallOption) (*GetClusterStatusResp, error)
}

type environmentManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewEnvironmentManagerClient(cc grpc.ClientConnInterface) EnvironmentManagerClient {
	return &environmentManagerClient{cc}
}

func (c *environmentManagerClient) GetClusterAgentApiToken(ctx context.Context, in *GetClusterAgentApiTokenReq, opts ...grpc.CallOption) (*GetClusterAgentApiTokenResp, error) {
	out := new(GetClusterAgentApiTokenResp)
	err := c.cc.Invoke(ctx, EnvironmentManager_GetClusterAgentApiToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentManagerClient) LinkCluster(ctx context.Context, in *LinkClusterReq, opts ...grpc.CallOption) (*LinkClusterResp, error) {
	out := new(LinkClusterResp)
	err := c.cc.Invoke(ctx, EnvironmentManager_LinkCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentManagerClient) ListClusters(ctx context.Context, in *ListClustersReq, opts ...grpc.CallOption) (*ListClustersResp, error) {
	out := new(ListClustersResp)
	err := c.cc.Invoke(ctx, EnvironmentManager_ListClusters_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentManagerClient) GetCluster(ctx context.Context, in *GetClusterReq, opts ...grpc.CallOption) (*GetClusterResp, error) {
	out := new(GetClusterResp)
	err := c.cc.Invoke(ctx, EnvironmentManager_GetCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentManagerClient) RemoveCluster(ctx context.Context, in *RemoveClusterReq, opts ...grpc.CallOption) (*RemoveClusterResp, error) {
	out := new(RemoveClusterResp)
	err := c.cc.Invoke(ctx, EnvironmentManager_RemoveCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentManagerClient) GetClusterAuth(ctx context.Context, in *GetClusterAuthReq, opts ...grpc.CallOption) (*GetClusterAuthResp, error) {
	out := new(GetClusterAuthResp)
	err := c.cc.Invoke(ctx, EnvironmentManager_GetClusterAuth_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentManagerClient) GetClusterConfig(ctx context.Context, in *GetClusterConfigReq, opts ...grpc.CallOption) (*GetClusterConfigResp, error) {
	out := new(GetClusterConfigResp)
	err := c.cc.Invoke(ctx, EnvironmentManager_GetClusterConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentManagerClient) DetectClusterConfig(ctx context.Context, in *DetectClusterConfigReq, opts ...grpc.CallOption) (*DetectClusterConfigResp, error) {
	out := new(DetectClusterConfigResp)
	err := c.cc.Invoke(ctx, EnvironmentManager_DetectClusterConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentManagerClient) ConfigureCluster(ctx context.Context, in *ConfigureClusterReq, opts ...grpc.CallOption) (*ConfigureClusterResp, error) {
	out := new(ConfigureClusterResp)
	err := c.cc.Invoke(ctx, EnvironmentManager_ConfigureCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentManagerClient) ValidateConfigureCluster(ctx context.Context, in *ConfigureClusterReq, opts ...grpc.CallOption) (*ValidateConfigureClusterResp, error) {
	out := new(ValidateConfigureClusterResp)
	err := c.cc.Invoke(ctx, EnvironmentManager_ValidateConfigureCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentManagerClient) GetClusterStatus(ctx context.Context, in *GetClusterStatusReq, opts ...grpc.CallOption) (*GetClusterStatusResp, error) {
	out := new(GetClusterStatusResp)
	err := c.cc.Invoke(ctx, EnvironmentManager_GetClusterStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnvironmentManagerServer is the server API for EnvironmentManager service.
// All implementations must embed UnimplementedEnvironmentManagerServer
// for forward compatibility
type EnvironmentManagerServer interface {
	GetClusterAgentApiToken(context.Context, *GetClusterAgentApiTokenReq) (*GetClusterAgentApiTokenResp, error)
	LinkCluster(context.Context, *LinkClusterReq) (*LinkClusterResp, error)
	ListClusters(context.Context, *ListClustersReq) (*ListClustersResp, error)
	GetCluster(context.Context, *GetClusterReq) (*GetClusterResp, error)
	RemoveCluster(context.Context, *RemoveClusterReq) (*RemoveClusterResp, error)
	// Deprecated.
	GetClusterAuth(context.Context, *GetClusterAuthReq) (*GetClusterAuthResp, error)
	GetClusterConfig(context.Context, *GetClusterConfigReq) (*GetClusterConfigResp, error)
	DetectClusterConfig(context.Context, *DetectClusterConfigReq) (*DetectClusterConfigResp, error)
	ConfigureCluster(context.Context, *ConfigureClusterReq) (*ConfigureClusterResp, error)
	ValidateConfigureCluster(context.Context, *ConfigureClusterReq) (*ValidateConfigureClusterResp, error)
	GetClusterStatus(context.Context, *GetClusterStatusReq) (*GetClusterStatusResp, error)
	mustEmbedUnimplementedEnvironmentManagerServer()
}

// UnimplementedEnvironmentManagerServer must be embedded to have forward compatible implementations.
type UnimplementedEnvironmentManagerServer struct {
}

func (UnimplementedEnvironmentManagerServer) GetClusterAgentApiToken(context.Context, *GetClusterAgentApiTokenReq) (*GetClusterAgentApiTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterAgentApiToken not implemented")
}
func (UnimplementedEnvironmentManagerServer) LinkCluster(context.Context, *LinkClusterReq) (*LinkClusterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LinkCluster not implemented")
}
func (UnimplementedEnvironmentManagerServer) ListClusters(context.Context, *ListClustersReq) (*ListClustersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClusters not implemented")
}
func (UnimplementedEnvironmentManagerServer) GetCluster(context.Context, *GetClusterReq) (*GetClusterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCluster not implemented")
}
func (UnimplementedEnvironmentManagerServer) RemoveCluster(context.Context, *RemoveClusterReq) (*RemoveClusterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveCluster not implemented")
}
func (UnimplementedEnvironmentManagerServer) GetClusterAuth(context.Context, *GetClusterAuthReq) (*GetClusterAuthResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterAuth not implemented")
}
func (UnimplementedEnvironmentManagerServer) GetClusterConfig(context.Context, *GetClusterConfigReq) (*GetClusterConfigResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterConfig not implemented")
}
func (UnimplementedEnvironmentManagerServer) DetectClusterConfig(context.Context, *DetectClusterConfigReq) (*DetectClusterConfigResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetectClusterConfig not implemented")
}
func (UnimplementedEnvironmentManagerServer) ConfigureCluster(context.Context, *ConfigureClusterReq) (*ConfigureClusterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureCluster not implemented")
}
func (UnimplementedEnvironmentManagerServer) ValidateConfigureCluster(context.Context, *ConfigureClusterReq) (*ValidateConfigureClusterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateConfigureCluster not implemented")
}
func (UnimplementedEnvironmentManagerServer) GetClusterStatus(context.Context, *GetClusterStatusReq) (*GetClusterStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterStatus not implemented")
}
func (UnimplementedEnvironmentManagerServer) mustEmbedUnimplementedEnvironmentManagerServer() {}

// UnsafeEnvironmentManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EnvironmentManagerServer will
// result in compilation errors.
type UnsafeEnvironmentManagerServer interface {
	mustEmbedUnimplementedEnvironmentManagerServer()
}

func RegisterEnvironmentManagerServer(s grpc.ServiceRegistrar, srv EnvironmentManagerServer) {
	s.RegisterService(&EnvironmentManager_ServiceDesc, srv)
}

func _EnvironmentManager_GetClusterAgentApiToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterAgentApiTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentManagerServer).GetClusterAgentApiToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentManager_GetClusterAgentApiToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentManagerServer).GetClusterAgentApiToken(ctx, req.(*GetClusterAgentApiTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentManager_LinkCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkClusterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentManagerServer).LinkCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentManager_LinkCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentManagerServer).LinkCluster(ctx, req.(*LinkClusterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentManager_ListClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClustersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentManagerServer).ListClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentManager_ListClusters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentManagerServer).ListClusters(ctx, req.(*ListClustersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentManager_GetCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentManagerServer).GetCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentManager_GetCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentManagerServer).GetCluster(ctx, req.(*GetClusterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentManager_RemoveCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveClusterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentManagerServer).RemoveCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentManager_RemoveCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentManagerServer).RemoveCluster(ctx, req.(*RemoveClusterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentManager_GetClusterAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentManagerServer).GetClusterAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentManager_GetClusterAuth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentManagerServer).GetClusterAuth(ctx, req.(*GetClusterAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentManager_GetClusterConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentManagerServer).GetClusterConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentManager_GetClusterConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentManagerServer).GetClusterConfig(ctx, req.(*GetClusterConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentManager_DetectClusterConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetectClusterConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentManagerServer).DetectClusterConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentManager_DetectClusterConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentManagerServer).DetectClusterConfig(ctx, req.(*DetectClusterConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentManager_ConfigureCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureClusterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentManagerServer).ConfigureCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentManager_ConfigureCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentManagerServer).ConfigureCluster(ctx, req.(*ConfigureClusterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentManager_ValidateConfigureCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureClusterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentManagerServer).ValidateConfigureCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentManager_ValidateConfigureCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentManagerServer).ValidateConfigureCluster(ctx, req.(*ConfigureClusterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentManager_GetClusterStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentManagerServer).GetClusterStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentManager_GetClusterStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentManagerServer).GetClusterStatus(ctx, req.(*GetClusterStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

// EnvironmentManager_ServiceDesc is the grpc.ServiceDesc for EnvironmentManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EnvironmentManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "prodvana.environment.EnvironmentManager",
	HandlerType: (*EnvironmentManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetClusterAgentApiToken",
			Handler:    _EnvironmentManager_GetClusterAgentApiToken_Handler,
		},
		{
			MethodName: "LinkCluster",
			Handler:    _EnvironmentManager_LinkCluster_Handler,
		},
		{
			MethodName: "ListClusters",
			Handler:    _EnvironmentManager_ListClusters_Handler,
		},
		{
			MethodName: "GetCluster",
			Handler:    _EnvironmentManager_GetCluster_Handler,
		},
		{
			MethodName: "RemoveCluster",
			Handler:    _EnvironmentManager_RemoveCluster_Handler,
		},
		{
			MethodName: "GetClusterAuth",
			Handler:    _EnvironmentManager_GetClusterAuth_Handler,
		},
		{
			MethodName: "GetClusterConfig",
			Handler:    _EnvironmentManager_GetClusterConfig_Handler,
		},
		{
			MethodName: "DetectClusterConfig",
			Handler:    _EnvironmentManager_DetectClusterConfig_Handler,
		},
		{
			MethodName: "ConfigureCluster",
			Handler:    _EnvironmentManager_ConfigureCluster_Handler,
		},
		{
			MethodName: "ValidateConfigureCluster",
			Handler:    _EnvironmentManager_ValidateConfigureCluster_Handler,
		},
		{
			MethodName: "GetClusterStatus",
			Handler:    _EnvironmentManager_GetClusterStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "prodvana/environment/environment_manager.proto",
}
