// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.10
// source: prodvana/deployment/manager.proto

package deployment

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
	DeploymentManager_RecordDeployment_FullMethodName        = "/prodvana.deployment.DeploymentManager/RecordDeployment"
	DeploymentManager_UpdateDeploymentStatus_FullMethodName  = "/prodvana.deployment.DeploymentManager/UpdateDeploymentStatus"
	DeploymentManager_ListDeployments_FullMethodName         = "/prodvana.deployment.DeploymentManager/ListDeployments"
	DeploymentManager_ListDeploymentsStream_FullMethodName   = "/prodvana.deployment.DeploymentManager/ListDeploymentsStream"
	DeploymentManager_CompareDeployment_FullMethodName       = "/prodvana.deployment.DeploymentManager/CompareDeployment"
	DeploymentManager_PreviewDeployment_FullMethodName       = "/prodvana.deployment.DeploymentManager/PreviewDeployment"
	DeploymentManager_GetLatestDeployments_FullMethodName    = "/prodvana.deployment.DeploymentManager/GetLatestDeployments"
	DeploymentManager_CheckCommitInDeployment_FullMethodName = "/prodvana.deployment.DeploymentManager/CheckCommitInDeployment"
)

// DeploymentManagerClient is the client API for DeploymentManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeploymentManagerClient interface {
	RecordDeployment(ctx context.Context, in *RecordDeploymentReq, opts ...grpc.CallOption) (*RecordDeploymentResp, error)
	UpdateDeploymentStatus(ctx context.Context, in *UpdateDeploymentStatusReq, opts ...grpc.CallOption) (*UpdateDeploymentStatusResp, error)
	ListDeployments(ctx context.Context, in *ListDeploymentsReq, opts ...grpc.CallOption) (*ListDeploymentsResp, error)
	// page tokens arguments are ignored here
	ListDeploymentsStream(ctx context.Context, in *ListDeploymentsReq, opts ...grpc.CallOption) (DeploymentManager_ListDeploymentsStreamClient, error)
	CompareDeployment(ctx context.Context, in *CompareDeploymentReq, opts ...grpc.CallOption) (*CompareDeploymentResp, error)
	PreviewDeployment(ctx context.Context, in *PreviewDeploymentReq, opts ...grpc.CallOption) (*PreviewDeploymentResp, error)
	// returns the latest deployments for each (application, service, deployment channel) tuple.
	GetLatestDeployments(ctx context.Context, in *GetLatestDeploymentsReq, opts ...grpc.CallOption) (*GetLatestDeploymentsResp, error)
	CheckCommitInDeployment(ctx context.Context, in *CheckCommitInDeploymentReq, opts ...grpc.CallOption) (*CheckCommitInDeploymentResp, error)
}

type deploymentManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewDeploymentManagerClient(cc grpc.ClientConnInterface) DeploymentManagerClient {
	return &deploymentManagerClient{cc}
}

func (c *deploymentManagerClient) RecordDeployment(ctx context.Context, in *RecordDeploymentReq, opts ...grpc.CallOption) (*RecordDeploymentResp, error) {
	out := new(RecordDeploymentResp)
	err := c.cc.Invoke(ctx, DeploymentManager_RecordDeployment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deploymentManagerClient) UpdateDeploymentStatus(ctx context.Context, in *UpdateDeploymentStatusReq, opts ...grpc.CallOption) (*UpdateDeploymentStatusResp, error) {
	out := new(UpdateDeploymentStatusResp)
	err := c.cc.Invoke(ctx, DeploymentManager_UpdateDeploymentStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deploymentManagerClient) ListDeployments(ctx context.Context, in *ListDeploymentsReq, opts ...grpc.CallOption) (*ListDeploymentsResp, error) {
	out := new(ListDeploymentsResp)
	err := c.cc.Invoke(ctx, DeploymentManager_ListDeployments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deploymentManagerClient) ListDeploymentsStream(ctx context.Context, in *ListDeploymentsReq, opts ...grpc.CallOption) (DeploymentManager_ListDeploymentsStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &DeploymentManager_ServiceDesc.Streams[0], DeploymentManager_ListDeploymentsStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &deploymentManagerListDeploymentsStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DeploymentManager_ListDeploymentsStreamClient interface {
	Recv() (*ListDeploymentsResp, error)
	grpc.ClientStream
}

type deploymentManagerListDeploymentsStreamClient struct {
	grpc.ClientStream
}

func (x *deploymentManagerListDeploymentsStreamClient) Recv() (*ListDeploymentsResp, error) {
	m := new(ListDeploymentsResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *deploymentManagerClient) CompareDeployment(ctx context.Context, in *CompareDeploymentReq, opts ...grpc.CallOption) (*CompareDeploymentResp, error) {
	out := new(CompareDeploymentResp)
	err := c.cc.Invoke(ctx, DeploymentManager_CompareDeployment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deploymentManagerClient) PreviewDeployment(ctx context.Context, in *PreviewDeploymentReq, opts ...grpc.CallOption) (*PreviewDeploymentResp, error) {
	out := new(PreviewDeploymentResp)
	err := c.cc.Invoke(ctx, DeploymentManager_PreviewDeployment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deploymentManagerClient) GetLatestDeployments(ctx context.Context, in *GetLatestDeploymentsReq, opts ...grpc.CallOption) (*GetLatestDeploymentsResp, error) {
	out := new(GetLatestDeploymentsResp)
	err := c.cc.Invoke(ctx, DeploymentManager_GetLatestDeployments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deploymentManagerClient) CheckCommitInDeployment(ctx context.Context, in *CheckCommitInDeploymentReq, opts ...grpc.CallOption) (*CheckCommitInDeploymentResp, error) {
	out := new(CheckCommitInDeploymentResp)
	err := c.cc.Invoke(ctx, DeploymentManager_CheckCommitInDeployment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeploymentManagerServer is the server API for DeploymentManager service.
// All implementations must embed UnimplementedDeploymentManagerServer
// for forward compatibility
type DeploymentManagerServer interface {
	RecordDeployment(context.Context, *RecordDeploymentReq) (*RecordDeploymentResp, error)
	UpdateDeploymentStatus(context.Context, *UpdateDeploymentStatusReq) (*UpdateDeploymentStatusResp, error)
	ListDeployments(context.Context, *ListDeploymentsReq) (*ListDeploymentsResp, error)
	// page tokens arguments are ignored here
	ListDeploymentsStream(*ListDeploymentsReq, DeploymentManager_ListDeploymentsStreamServer) error
	CompareDeployment(context.Context, *CompareDeploymentReq) (*CompareDeploymentResp, error)
	PreviewDeployment(context.Context, *PreviewDeploymentReq) (*PreviewDeploymentResp, error)
	// returns the latest deployments for each (application, service, deployment channel) tuple.
	GetLatestDeployments(context.Context, *GetLatestDeploymentsReq) (*GetLatestDeploymentsResp, error)
	CheckCommitInDeployment(context.Context, *CheckCommitInDeploymentReq) (*CheckCommitInDeploymentResp, error)
	mustEmbedUnimplementedDeploymentManagerServer()
}

// UnimplementedDeploymentManagerServer must be embedded to have forward compatible implementations.
type UnimplementedDeploymentManagerServer struct {
}

func (UnimplementedDeploymentManagerServer) RecordDeployment(context.Context, *RecordDeploymentReq) (*RecordDeploymentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordDeployment not implemented")
}
func (UnimplementedDeploymentManagerServer) UpdateDeploymentStatus(context.Context, *UpdateDeploymentStatusReq) (*UpdateDeploymentStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDeploymentStatus not implemented")
}
func (UnimplementedDeploymentManagerServer) ListDeployments(context.Context, *ListDeploymentsReq) (*ListDeploymentsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDeployments not implemented")
}
func (UnimplementedDeploymentManagerServer) ListDeploymentsStream(*ListDeploymentsReq, DeploymentManager_ListDeploymentsStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ListDeploymentsStream not implemented")
}
func (UnimplementedDeploymentManagerServer) CompareDeployment(context.Context, *CompareDeploymentReq) (*CompareDeploymentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompareDeployment not implemented")
}
func (UnimplementedDeploymentManagerServer) PreviewDeployment(context.Context, *PreviewDeploymentReq) (*PreviewDeploymentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreviewDeployment not implemented")
}
func (UnimplementedDeploymentManagerServer) GetLatestDeployments(context.Context, *GetLatestDeploymentsReq) (*GetLatestDeploymentsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestDeployments not implemented")
}
func (UnimplementedDeploymentManagerServer) CheckCommitInDeployment(context.Context, *CheckCommitInDeploymentReq) (*CheckCommitInDeploymentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckCommitInDeployment not implemented")
}
func (UnimplementedDeploymentManagerServer) mustEmbedUnimplementedDeploymentManagerServer() {}

// UnsafeDeploymentManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeploymentManagerServer will
// result in compilation errors.
type UnsafeDeploymentManagerServer interface {
	mustEmbedUnimplementedDeploymentManagerServer()
}

func RegisterDeploymentManagerServer(s grpc.ServiceRegistrar, srv DeploymentManagerServer) {
	s.RegisterService(&DeploymentManager_ServiceDesc, srv)
}

func _DeploymentManager_RecordDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordDeploymentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentManagerServer).RecordDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeploymentManager_RecordDeployment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentManagerServer).RecordDeployment(ctx, req.(*RecordDeploymentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeploymentManager_UpdateDeploymentStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeploymentStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentManagerServer).UpdateDeploymentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeploymentManager_UpdateDeploymentStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentManagerServer).UpdateDeploymentStatus(ctx, req.(*UpdateDeploymentStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeploymentManager_ListDeployments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDeploymentsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentManagerServer).ListDeployments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeploymentManager_ListDeployments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentManagerServer).ListDeployments(ctx, req.(*ListDeploymentsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeploymentManager_ListDeploymentsStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListDeploymentsReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DeploymentManagerServer).ListDeploymentsStream(m, &deploymentManagerListDeploymentsStreamServer{stream})
}

type DeploymentManager_ListDeploymentsStreamServer interface {
	Send(*ListDeploymentsResp) error
	grpc.ServerStream
}

type deploymentManagerListDeploymentsStreamServer struct {
	grpc.ServerStream
}

func (x *deploymentManagerListDeploymentsStreamServer) Send(m *ListDeploymentsResp) error {
	return x.ServerStream.SendMsg(m)
}

func _DeploymentManager_CompareDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompareDeploymentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentManagerServer).CompareDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeploymentManager_CompareDeployment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentManagerServer).CompareDeployment(ctx, req.(*CompareDeploymentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeploymentManager_PreviewDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreviewDeploymentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentManagerServer).PreviewDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeploymentManager_PreviewDeployment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentManagerServer).PreviewDeployment(ctx, req.(*PreviewDeploymentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeploymentManager_GetLatestDeployments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestDeploymentsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentManagerServer).GetLatestDeployments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeploymentManager_GetLatestDeployments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentManagerServer).GetLatestDeployments(ctx, req.(*GetLatestDeploymentsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeploymentManager_CheckCommitInDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckCommitInDeploymentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentManagerServer).CheckCommitInDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeploymentManager_CheckCommitInDeployment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentManagerServer).CheckCommitInDeployment(ctx, req.(*CheckCommitInDeploymentReq))
	}
	return interceptor(ctx, in, info, handler)
}

// DeploymentManager_ServiceDesc is the grpc.ServiceDesc for DeploymentManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeploymentManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "prodvana.deployment.DeploymentManager",
	HandlerType: (*DeploymentManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecordDeployment",
			Handler:    _DeploymentManager_RecordDeployment_Handler,
		},
		{
			MethodName: "UpdateDeploymentStatus",
			Handler:    _DeploymentManager_UpdateDeploymentStatus_Handler,
		},
		{
			MethodName: "ListDeployments",
			Handler:    _DeploymentManager_ListDeployments_Handler,
		},
		{
			MethodName: "CompareDeployment",
			Handler:    _DeploymentManager_CompareDeployment_Handler,
		},
		{
			MethodName: "PreviewDeployment",
			Handler:    _DeploymentManager_PreviewDeployment_Handler,
		},
		{
			MethodName: "GetLatestDeployments",
			Handler:    _DeploymentManager_GetLatestDeployments_Handler,
		},
		{
			MethodName: "CheckCommitInDeployment",
			Handler:    _DeploymentManager_CheckCommitInDeployment_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListDeploymentsStream",
			Handler:       _DeploymentManager_ListDeploymentsStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "prodvana/deployment/manager.proto",
}