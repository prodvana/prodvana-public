// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.10
// source: prodvana/release/manager.proto

package release

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
	ReleaseManager_RecordRelease_FullMethodName       = "/prodvana.release.ReleaseManager/RecordRelease"
	ReleaseManager_UpdateReleaseStatus_FullMethodName = "/prodvana.release.ReleaseManager/UpdateReleaseStatus"
	ReleaseManager_ListReleases_FullMethodName        = "/prodvana.release.ReleaseManager/ListReleases"
	ReleaseManager_ListReleasesStream_FullMethodName  = "/prodvana.release.ReleaseManager/ListReleasesStream"
)

// ReleaseManagerClient is the client API for ReleaseManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReleaseManagerClient interface {
	RecordRelease(ctx context.Context, in *RecordReleaseReq, opts ...grpc.CallOption) (*RecordReleaseResp, error)
	UpdateReleaseStatus(ctx context.Context, in *UpdateReleaseStatusReq, opts ...grpc.CallOption) (*UpdateReleaseStatusResp, error)
	ListReleases(ctx context.Context, in *ListReleasesReq, opts ...grpc.CallOption) (*ListReleasesResp, error)
	// page tokens arguments are ignored here
	ListReleasesStream(ctx context.Context, in *ListReleasesReq, opts ...grpc.CallOption) (ReleaseManager_ListReleasesStreamClient, error)
}

type releaseManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewReleaseManagerClient(cc grpc.ClientConnInterface) ReleaseManagerClient {
	return &releaseManagerClient{cc}
}

func (c *releaseManagerClient) RecordRelease(ctx context.Context, in *RecordReleaseReq, opts ...grpc.CallOption) (*RecordReleaseResp, error) {
	out := new(RecordReleaseResp)
	err := c.cc.Invoke(ctx, ReleaseManager_RecordRelease_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseManagerClient) UpdateReleaseStatus(ctx context.Context, in *UpdateReleaseStatusReq, opts ...grpc.CallOption) (*UpdateReleaseStatusResp, error) {
	out := new(UpdateReleaseStatusResp)
	err := c.cc.Invoke(ctx, ReleaseManager_UpdateReleaseStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseManagerClient) ListReleases(ctx context.Context, in *ListReleasesReq, opts ...grpc.CallOption) (*ListReleasesResp, error) {
	out := new(ListReleasesResp)
	err := c.cc.Invoke(ctx, ReleaseManager_ListReleases_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseManagerClient) ListReleasesStream(ctx context.Context, in *ListReleasesReq, opts ...grpc.CallOption) (ReleaseManager_ListReleasesStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ReleaseManager_ServiceDesc.Streams[0], ReleaseManager_ListReleasesStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &releaseManagerListReleasesStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ReleaseManager_ListReleasesStreamClient interface {
	Recv() (*ListReleasesResp, error)
	grpc.ClientStream
}

type releaseManagerListReleasesStreamClient struct {
	grpc.ClientStream
}

func (x *releaseManagerListReleasesStreamClient) Recv() (*ListReleasesResp, error) {
	m := new(ListReleasesResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ReleaseManagerServer is the server API for ReleaseManager service.
// All implementations must embed UnimplementedReleaseManagerServer
// for forward compatibility
type ReleaseManagerServer interface {
	RecordRelease(context.Context, *RecordReleaseReq) (*RecordReleaseResp, error)
	UpdateReleaseStatus(context.Context, *UpdateReleaseStatusReq) (*UpdateReleaseStatusResp, error)
	ListReleases(context.Context, *ListReleasesReq) (*ListReleasesResp, error)
	// page tokens arguments are ignored here
	ListReleasesStream(*ListReleasesReq, ReleaseManager_ListReleasesStreamServer) error
	mustEmbedUnimplementedReleaseManagerServer()
}

// UnimplementedReleaseManagerServer must be embedded to have forward compatible implementations.
type UnimplementedReleaseManagerServer struct {
}

func (UnimplementedReleaseManagerServer) RecordRelease(context.Context, *RecordReleaseReq) (*RecordReleaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordRelease not implemented")
}
func (UnimplementedReleaseManagerServer) UpdateReleaseStatus(context.Context, *UpdateReleaseStatusReq) (*UpdateReleaseStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReleaseStatus not implemented")
}
func (UnimplementedReleaseManagerServer) ListReleases(context.Context, *ListReleasesReq) (*ListReleasesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReleases not implemented")
}
func (UnimplementedReleaseManagerServer) ListReleasesStream(*ListReleasesReq, ReleaseManager_ListReleasesStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ListReleasesStream not implemented")
}
func (UnimplementedReleaseManagerServer) mustEmbedUnimplementedReleaseManagerServer() {}

// UnsafeReleaseManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReleaseManagerServer will
// result in compilation errors.
type UnsafeReleaseManagerServer interface {
	mustEmbedUnimplementedReleaseManagerServer()
}

func RegisterReleaseManagerServer(s grpc.ServiceRegistrar, srv ReleaseManagerServer) {
	s.RegisterService(&ReleaseManager_ServiceDesc, srv)
}

func _ReleaseManager_RecordRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordReleaseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseManagerServer).RecordRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReleaseManager_RecordRelease_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseManagerServer).RecordRelease(ctx, req.(*RecordReleaseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseManager_UpdateReleaseStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReleaseStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseManagerServer).UpdateReleaseStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReleaseManager_UpdateReleaseStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseManagerServer).UpdateReleaseStatus(ctx, req.(*UpdateReleaseStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseManager_ListReleases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReleasesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseManagerServer).ListReleases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReleaseManager_ListReleases_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseManagerServer).ListReleases(ctx, req.(*ListReleasesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseManager_ListReleasesStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListReleasesReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReleaseManagerServer).ListReleasesStream(m, &releaseManagerListReleasesStreamServer{stream})
}

type ReleaseManager_ListReleasesStreamServer interface {
	Send(*ListReleasesResp) error
	grpc.ServerStream
}

type releaseManagerListReleasesStreamServer struct {
	grpc.ServerStream
}

func (x *releaseManagerListReleasesStreamServer) Send(m *ListReleasesResp) error {
	return x.ServerStream.SendMsg(m)
}

// ReleaseManager_ServiceDesc is the grpc.ServiceDesc for ReleaseManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReleaseManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "prodvana.release.ReleaseManager",
	HandlerType: (*ReleaseManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecordRelease",
			Handler:    _ReleaseManager_RecordRelease_Handler,
		},
		{
			MethodName: "UpdateReleaseStatus",
			Handler:    _ReleaseManager_UpdateReleaseStatus_Handler,
		},
		{
			MethodName: "ListReleases",
			Handler:    _ReleaseManager_ListReleases_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListReleasesStream",
			Handler:       _ReleaseManager_ListReleasesStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "prodvana/release/manager.proto",
}