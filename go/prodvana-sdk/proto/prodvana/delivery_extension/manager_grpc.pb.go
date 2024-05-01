// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.10
// source: prodvana/delivery_extension/manager.proto

package delivery_extension

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
	DeliveryExtensionManager_ConfigureDeliveryExtension_FullMethodName         = "/prodvana.delivery_extension.DeliveryExtensionManager/ConfigureDeliveryExtension"
	DeliveryExtensionManager_ValidateConfigureDeliveryExtension_FullMethodName = "/prodvana.delivery_extension.DeliveryExtensionManager/ValidateConfigureDeliveryExtension"
	DeliveryExtensionManager_ListDeliveryExtensions_FullMethodName             = "/prodvana.delivery_extension.DeliveryExtensionManager/ListDeliveryExtensions"
	DeliveryExtensionManager_GetDeliveryExtension_FullMethodName               = "/prodvana.delivery_extension.DeliveryExtensionManager/GetDeliveryExtension"
	DeliveryExtensionManager_GetDeliveryExtensionConfig_FullMethodName         = "/prodvana.delivery_extension.DeliveryExtensionManager/GetDeliveryExtensionConfig"
	DeliveryExtensionManager_GetDeliveryExtensionInstanceConfig_FullMethodName = "/prodvana.delivery_extension.DeliveryExtensionManager/GetDeliveryExtensionInstanceConfig"
)

// DeliveryExtensionManagerClient is the client API for DeliveryExtensionManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeliveryExtensionManagerClient interface {
	ConfigureDeliveryExtension(ctx context.Context, in *ConfigureDeliveryExtensionReq, opts ...grpc.CallOption) (*ConfigureDeliveryExtensionResp, error)
	ValidateConfigureDeliveryExtension(ctx context.Context, in *ConfigureDeliveryExtensionReq, opts ...grpc.CallOption) (*ValidateConfigureDeliveryExtensionResp, error)
	ListDeliveryExtensions(ctx context.Context, in *ListDeliveryExtensionsReq, opts ...grpc.CallOption) (*ListDeliveryExtensionsResp, error)
	GetDeliveryExtension(ctx context.Context, in *GetDeliveryExtensionReq, opts ...grpc.CallOption) (*GetDeliveryExtensionResp, error)
	GetDeliveryExtensionConfig(ctx context.Context, in *GetDeliveryExtensionConfigReq, opts ...grpc.CallOption) (*GetDeliveryExtensionConfigResp, error)
	GetDeliveryExtensionInstanceConfig(ctx context.Context, in *GetDeliveryExtensionInstanceConfigReq, opts ...grpc.CallOption) (*GetDeliveryExtensionInstanceConfigResp, error)
}

type deliveryExtensionManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewDeliveryExtensionManagerClient(cc grpc.ClientConnInterface) DeliveryExtensionManagerClient {
	return &deliveryExtensionManagerClient{cc}
}

func (c *deliveryExtensionManagerClient) ConfigureDeliveryExtension(ctx context.Context, in *ConfigureDeliveryExtensionReq, opts ...grpc.CallOption) (*ConfigureDeliveryExtensionResp, error) {
	out := new(ConfigureDeliveryExtensionResp)
	err := c.cc.Invoke(ctx, DeliveryExtensionManager_ConfigureDeliveryExtension_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveryExtensionManagerClient) ValidateConfigureDeliveryExtension(ctx context.Context, in *ConfigureDeliveryExtensionReq, opts ...grpc.CallOption) (*ValidateConfigureDeliveryExtensionResp, error) {
	out := new(ValidateConfigureDeliveryExtensionResp)
	err := c.cc.Invoke(ctx, DeliveryExtensionManager_ValidateConfigureDeliveryExtension_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveryExtensionManagerClient) ListDeliveryExtensions(ctx context.Context, in *ListDeliveryExtensionsReq, opts ...grpc.CallOption) (*ListDeliveryExtensionsResp, error) {
	out := new(ListDeliveryExtensionsResp)
	err := c.cc.Invoke(ctx, DeliveryExtensionManager_ListDeliveryExtensions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveryExtensionManagerClient) GetDeliveryExtension(ctx context.Context, in *GetDeliveryExtensionReq, opts ...grpc.CallOption) (*GetDeliveryExtensionResp, error) {
	out := new(GetDeliveryExtensionResp)
	err := c.cc.Invoke(ctx, DeliveryExtensionManager_GetDeliveryExtension_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveryExtensionManagerClient) GetDeliveryExtensionConfig(ctx context.Context, in *GetDeliveryExtensionConfigReq, opts ...grpc.CallOption) (*GetDeliveryExtensionConfigResp, error) {
	out := new(GetDeliveryExtensionConfigResp)
	err := c.cc.Invoke(ctx, DeliveryExtensionManager_GetDeliveryExtensionConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveryExtensionManagerClient) GetDeliveryExtensionInstanceConfig(ctx context.Context, in *GetDeliveryExtensionInstanceConfigReq, opts ...grpc.CallOption) (*GetDeliveryExtensionInstanceConfigResp, error) {
	out := new(GetDeliveryExtensionInstanceConfigResp)
	err := c.cc.Invoke(ctx, DeliveryExtensionManager_GetDeliveryExtensionInstanceConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeliveryExtensionManagerServer is the server API for DeliveryExtensionManager service.
// All implementations must embed UnimplementedDeliveryExtensionManagerServer
// for forward compatibility
type DeliveryExtensionManagerServer interface {
	ConfigureDeliveryExtension(context.Context, *ConfigureDeliveryExtensionReq) (*ConfigureDeliveryExtensionResp, error)
	ValidateConfigureDeliveryExtension(context.Context, *ConfigureDeliveryExtensionReq) (*ValidateConfigureDeliveryExtensionResp, error)
	ListDeliveryExtensions(context.Context, *ListDeliveryExtensionsReq) (*ListDeliveryExtensionsResp, error)
	GetDeliveryExtension(context.Context, *GetDeliveryExtensionReq) (*GetDeliveryExtensionResp, error)
	GetDeliveryExtensionConfig(context.Context, *GetDeliveryExtensionConfigReq) (*GetDeliveryExtensionConfigResp, error)
	GetDeliveryExtensionInstanceConfig(context.Context, *GetDeliveryExtensionInstanceConfigReq) (*GetDeliveryExtensionInstanceConfigResp, error)
	mustEmbedUnimplementedDeliveryExtensionManagerServer()
}

// UnimplementedDeliveryExtensionManagerServer must be embedded to have forward compatible implementations.
type UnimplementedDeliveryExtensionManagerServer struct {
}

func (UnimplementedDeliveryExtensionManagerServer) ConfigureDeliveryExtension(context.Context, *ConfigureDeliveryExtensionReq) (*ConfigureDeliveryExtensionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureDeliveryExtension not implemented")
}
func (UnimplementedDeliveryExtensionManagerServer) ValidateConfigureDeliveryExtension(context.Context, *ConfigureDeliveryExtensionReq) (*ValidateConfigureDeliveryExtensionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateConfigureDeliveryExtension not implemented")
}
func (UnimplementedDeliveryExtensionManagerServer) ListDeliveryExtensions(context.Context, *ListDeliveryExtensionsReq) (*ListDeliveryExtensionsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDeliveryExtensions not implemented")
}
func (UnimplementedDeliveryExtensionManagerServer) GetDeliveryExtension(context.Context, *GetDeliveryExtensionReq) (*GetDeliveryExtensionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeliveryExtension not implemented")
}
func (UnimplementedDeliveryExtensionManagerServer) GetDeliveryExtensionConfig(context.Context, *GetDeliveryExtensionConfigReq) (*GetDeliveryExtensionConfigResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeliveryExtensionConfig not implemented")
}
func (UnimplementedDeliveryExtensionManagerServer) GetDeliveryExtensionInstanceConfig(context.Context, *GetDeliveryExtensionInstanceConfigReq) (*GetDeliveryExtensionInstanceConfigResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeliveryExtensionInstanceConfig not implemented")
}
func (UnimplementedDeliveryExtensionManagerServer) mustEmbedUnimplementedDeliveryExtensionManagerServer() {
}

// UnsafeDeliveryExtensionManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeliveryExtensionManagerServer will
// result in compilation errors.
type UnsafeDeliveryExtensionManagerServer interface {
	mustEmbedUnimplementedDeliveryExtensionManagerServer()
}

func RegisterDeliveryExtensionManagerServer(s grpc.ServiceRegistrar, srv DeliveryExtensionManagerServer) {
	s.RegisterService(&DeliveryExtensionManager_ServiceDesc, srv)
}

func _DeliveryExtensionManager_ConfigureDeliveryExtension_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureDeliveryExtensionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryExtensionManagerServer).ConfigureDeliveryExtension(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeliveryExtensionManager_ConfigureDeliveryExtension_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryExtensionManagerServer).ConfigureDeliveryExtension(ctx, req.(*ConfigureDeliveryExtensionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeliveryExtensionManager_ValidateConfigureDeliveryExtension_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureDeliveryExtensionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryExtensionManagerServer).ValidateConfigureDeliveryExtension(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeliveryExtensionManager_ValidateConfigureDeliveryExtension_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryExtensionManagerServer).ValidateConfigureDeliveryExtension(ctx, req.(*ConfigureDeliveryExtensionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeliveryExtensionManager_ListDeliveryExtensions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDeliveryExtensionsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryExtensionManagerServer).ListDeliveryExtensions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeliveryExtensionManager_ListDeliveryExtensions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryExtensionManagerServer).ListDeliveryExtensions(ctx, req.(*ListDeliveryExtensionsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeliveryExtensionManager_GetDeliveryExtension_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeliveryExtensionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryExtensionManagerServer).GetDeliveryExtension(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeliveryExtensionManager_GetDeliveryExtension_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryExtensionManagerServer).GetDeliveryExtension(ctx, req.(*GetDeliveryExtensionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeliveryExtensionManager_GetDeliveryExtensionConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeliveryExtensionConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryExtensionManagerServer).GetDeliveryExtensionConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeliveryExtensionManager_GetDeliveryExtensionConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryExtensionManagerServer).GetDeliveryExtensionConfig(ctx, req.(*GetDeliveryExtensionConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeliveryExtensionManager_GetDeliveryExtensionInstanceConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeliveryExtensionInstanceConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryExtensionManagerServer).GetDeliveryExtensionInstanceConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeliveryExtensionManager_GetDeliveryExtensionInstanceConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryExtensionManagerServer).GetDeliveryExtensionInstanceConfig(ctx, req.(*GetDeliveryExtensionInstanceConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

// DeliveryExtensionManager_ServiceDesc is the grpc.ServiceDesc for DeliveryExtensionManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeliveryExtensionManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "prodvana.delivery_extension.DeliveryExtensionManager",
	HandlerType: (*DeliveryExtensionManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConfigureDeliveryExtension",
			Handler:    _DeliveryExtensionManager_ConfigureDeliveryExtension_Handler,
		},
		{
			MethodName: "ValidateConfigureDeliveryExtension",
			Handler:    _DeliveryExtensionManager_ValidateConfigureDeliveryExtension_Handler,
		},
		{
			MethodName: "ListDeliveryExtensions",
			Handler:    _DeliveryExtensionManager_ListDeliveryExtensions_Handler,
		},
		{
			MethodName: "GetDeliveryExtension",
			Handler:    _DeliveryExtensionManager_GetDeliveryExtension_Handler,
		},
		{
			MethodName: "GetDeliveryExtensionConfig",
			Handler:    _DeliveryExtensionManager_GetDeliveryExtensionConfig_Handler,
		},
		{
			MethodName: "GetDeliveryExtensionInstanceConfig",
			Handler:    _DeliveryExtensionManager_GetDeliveryExtensionInstanceConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "prodvana/delivery_extension/manager.proto",
}
