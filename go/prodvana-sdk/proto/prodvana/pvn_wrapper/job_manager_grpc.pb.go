// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.10
// source: prodvana/pvn_wrapper/job_manager.proto

package pvn_wrapper

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
	JobManager_ReportJobResult_FullMethodName = "/prodvana.pvn_wrapper.JobManager/ReportJobResult"
)

// JobManagerClient is the client API for JobManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobManagerClient interface {
	ReportJobResult(ctx context.Context, in *ReportJobResultReq, opts ...grpc.CallOption) (*ReportJobResultResp, error)
}

type jobManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewJobManagerClient(cc grpc.ClientConnInterface) JobManagerClient {
	return &jobManagerClient{cc}
}

func (c *jobManagerClient) ReportJobResult(ctx context.Context, in *ReportJobResultReq, opts ...grpc.CallOption) (*ReportJobResultResp, error) {
	out := new(ReportJobResultResp)
	err := c.cc.Invoke(ctx, JobManager_ReportJobResult_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobManagerServer is the server API for JobManager service.
// All implementations must embed UnimplementedJobManagerServer
// for forward compatibility
type JobManagerServer interface {
	ReportJobResult(context.Context, *ReportJobResultReq) (*ReportJobResultResp, error)
	mustEmbedUnimplementedJobManagerServer()
}

// UnimplementedJobManagerServer must be embedded to have forward compatible implementations.
type UnimplementedJobManagerServer struct {
}

func (UnimplementedJobManagerServer) ReportJobResult(context.Context, *ReportJobResultReq) (*ReportJobResultResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportJobResult not implemented")
}
func (UnimplementedJobManagerServer) mustEmbedUnimplementedJobManagerServer() {}

// UnsafeJobManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobManagerServer will
// result in compilation errors.
type UnsafeJobManagerServer interface {
	mustEmbedUnimplementedJobManagerServer()
}

func RegisterJobManagerServer(s grpc.ServiceRegistrar, srv JobManagerServer) {
	s.RegisterService(&JobManager_ServiceDesc, srv)
}

func _JobManager_ReportJobResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportJobResultReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobManagerServer).ReportJobResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobManager_ReportJobResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobManagerServer).ReportJobResult(ctx, req.(*ReportJobResultReq))
	}
	return interceptor(ctx, in, info, handler)
}

// JobManager_ServiceDesc is the grpc.ServiceDesc for JobManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JobManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "prodvana.pvn_wrapper.JobManager",
	HandlerType: (*JobManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportJobResult",
			Handler:    _JobManager_ReportJobResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "prodvana/pvn_wrapper/job_manager.proto",
}
