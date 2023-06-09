// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.10
// source: prodvana/desired_state/manager.proto

package desired_state

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
	DesiredStateManager_SetDesiredState_FullMethodName                          = "/prodvana.desired_state.DesiredStateManager/SetDesiredState"
	DesiredStateManager_GetServiceDesiredStateConvergenceSummary_FullMethodName = "/prodvana.desired_state.DesiredStateManager/GetServiceDesiredStateConvergenceSummary"
	DesiredStateManager_GetServiceLastConvergedStates_FullMethodName            = "/prodvana.desired_state.DesiredStateManager/GetServiceLastConvergedStates"
	DesiredStateManager_GetServiceDesiredStateHistory_FullMethodName            = "/prodvana.desired_state.DesiredStateManager/GetServiceDesiredStateHistory"
	DesiredStateManager_GetDesiredState_FullMethodName                          = "/prodvana.desired_state.DesiredStateManager/GetDesiredState"
	DesiredStateManager_GetDesiredStateConvergenceSummary_FullMethodName        = "/prodvana.desired_state.DesiredStateManager/GetDesiredStateConvergenceSummary"
	DesiredStateManager_ValidateDesiredState_FullMethodName                     = "/prodvana.desired_state.DesiredStateManager/ValidateDesiredState"
	DesiredStateManager_SetManualApproval_FullMethodName                        = "/prodvana.desired_state.DesiredStateManager/SetManualApproval"
	DesiredStateManager_PromoteDelivery_FullMethodName                          = "/prodvana.desired_state.DesiredStateManager/PromoteDelivery"
	DesiredStateManager_BypassProtection_FullMethodName                         = "/prodvana.desired_state.DesiredStateManager/BypassProtection"
)

// DesiredStateManagerClient is the client API for DesiredStateManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DesiredStateManagerClient interface {
	SetDesiredState(ctx context.Context, in *SetDesiredStateReq, opts ...grpc.CallOption) (*SetDesiredStateResp, error)
	GetServiceDesiredStateConvergenceSummary(ctx context.Context, in *GetServiceDesiredStateConvergenceSummaryReq, opts ...grpc.CallOption) (*GetServiceDesiredStateConvergenceSummaryResp, error)
	GetServiceLastConvergedStates(ctx context.Context, in *GetServiceLastConvergedStateReq, opts ...grpc.CallOption) (*GetServiceLastConvergedStateResp, error)
	GetServiceDesiredStateHistory(ctx context.Context, in *GetServiceDesiredStateHistoryReq, opts ...grpc.CallOption) (*GetServiceDesiredStateHistoryResp, error)
	GetDesiredState(ctx context.Context, in *GetDesiredStateReq, opts ...grpc.CallOption) (*GetDesiredStateResp, error)
	GetDesiredStateConvergenceSummary(ctx context.Context, in *GetDesiredStateConvergenceReq, opts ...grpc.CallOption) (*GetDesiredStateConvergenceSummaryResp, error)
	ValidateDesiredState(ctx context.Context, in *ValidateDesiredStateReq, opts ...grpc.CallOption) (*ValidateDesiredStateResp, error)
	SetManualApproval(ctx context.Context, in *SetManualApprovalReq, opts ...grpc.CallOption) (*SetManualApprovalResp, error)
	PromoteDelivery(ctx context.Context, in *PromoteDeliveryReq, opts ...grpc.CallOption) (*PromoteDeliveryResp, error)
	BypassProtection(ctx context.Context, in *BypassProtectionReq, opts ...grpc.CallOption) (*BypassProtectionResp, error)
}

type desiredStateManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewDesiredStateManagerClient(cc grpc.ClientConnInterface) DesiredStateManagerClient {
	return &desiredStateManagerClient{cc}
}

func (c *desiredStateManagerClient) SetDesiredState(ctx context.Context, in *SetDesiredStateReq, opts ...grpc.CallOption) (*SetDesiredStateResp, error) {
	out := new(SetDesiredStateResp)
	err := c.cc.Invoke(ctx, DesiredStateManager_SetDesiredState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *desiredStateManagerClient) GetServiceDesiredStateConvergenceSummary(ctx context.Context, in *GetServiceDesiredStateConvergenceSummaryReq, opts ...grpc.CallOption) (*GetServiceDesiredStateConvergenceSummaryResp, error) {
	out := new(GetServiceDesiredStateConvergenceSummaryResp)
	err := c.cc.Invoke(ctx, DesiredStateManager_GetServiceDesiredStateConvergenceSummary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *desiredStateManagerClient) GetServiceLastConvergedStates(ctx context.Context, in *GetServiceLastConvergedStateReq, opts ...grpc.CallOption) (*GetServiceLastConvergedStateResp, error) {
	out := new(GetServiceLastConvergedStateResp)
	err := c.cc.Invoke(ctx, DesiredStateManager_GetServiceLastConvergedStates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *desiredStateManagerClient) GetServiceDesiredStateHistory(ctx context.Context, in *GetServiceDesiredStateHistoryReq, opts ...grpc.CallOption) (*GetServiceDesiredStateHistoryResp, error) {
	out := new(GetServiceDesiredStateHistoryResp)
	err := c.cc.Invoke(ctx, DesiredStateManager_GetServiceDesiredStateHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *desiredStateManagerClient) GetDesiredState(ctx context.Context, in *GetDesiredStateReq, opts ...grpc.CallOption) (*GetDesiredStateResp, error) {
	out := new(GetDesiredStateResp)
	err := c.cc.Invoke(ctx, DesiredStateManager_GetDesiredState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *desiredStateManagerClient) GetDesiredStateConvergenceSummary(ctx context.Context, in *GetDesiredStateConvergenceReq, opts ...grpc.CallOption) (*GetDesiredStateConvergenceSummaryResp, error) {
	out := new(GetDesiredStateConvergenceSummaryResp)
	err := c.cc.Invoke(ctx, DesiredStateManager_GetDesiredStateConvergenceSummary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *desiredStateManagerClient) ValidateDesiredState(ctx context.Context, in *ValidateDesiredStateReq, opts ...grpc.CallOption) (*ValidateDesiredStateResp, error) {
	out := new(ValidateDesiredStateResp)
	err := c.cc.Invoke(ctx, DesiredStateManager_ValidateDesiredState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *desiredStateManagerClient) SetManualApproval(ctx context.Context, in *SetManualApprovalReq, opts ...grpc.CallOption) (*SetManualApprovalResp, error) {
	out := new(SetManualApprovalResp)
	err := c.cc.Invoke(ctx, DesiredStateManager_SetManualApproval_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *desiredStateManagerClient) PromoteDelivery(ctx context.Context, in *PromoteDeliveryReq, opts ...grpc.CallOption) (*PromoteDeliveryResp, error) {
	out := new(PromoteDeliveryResp)
	err := c.cc.Invoke(ctx, DesiredStateManager_PromoteDelivery_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *desiredStateManagerClient) BypassProtection(ctx context.Context, in *BypassProtectionReq, opts ...grpc.CallOption) (*BypassProtectionResp, error) {
	out := new(BypassProtectionResp)
	err := c.cc.Invoke(ctx, DesiredStateManager_BypassProtection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DesiredStateManagerServer is the server API for DesiredStateManager service.
// All implementations must embed UnimplementedDesiredStateManagerServer
// for forward compatibility
type DesiredStateManagerServer interface {
	SetDesiredState(context.Context, *SetDesiredStateReq) (*SetDesiredStateResp, error)
	GetServiceDesiredStateConvergenceSummary(context.Context, *GetServiceDesiredStateConvergenceSummaryReq) (*GetServiceDesiredStateConvergenceSummaryResp, error)
	GetServiceLastConvergedStates(context.Context, *GetServiceLastConvergedStateReq) (*GetServiceLastConvergedStateResp, error)
	GetServiceDesiredStateHistory(context.Context, *GetServiceDesiredStateHistoryReq) (*GetServiceDesiredStateHistoryResp, error)
	GetDesiredState(context.Context, *GetDesiredStateReq) (*GetDesiredStateResp, error)
	GetDesiredStateConvergenceSummary(context.Context, *GetDesiredStateConvergenceReq) (*GetDesiredStateConvergenceSummaryResp, error)
	ValidateDesiredState(context.Context, *ValidateDesiredStateReq) (*ValidateDesiredStateResp, error)
	SetManualApproval(context.Context, *SetManualApprovalReq) (*SetManualApprovalResp, error)
	PromoteDelivery(context.Context, *PromoteDeliveryReq) (*PromoteDeliveryResp, error)
	BypassProtection(context.Context, *BypassProtectionReq) (*BypassProtectionResp, error)
	mustEmbedUnimplementedDesiredStateManagerServer()
}

// UnimplementedDesiredStateManagerServer must be embedded to have forward compatible implementations.
type UnimplementedDesiredStateManagerServer struct {
}

func (UnimplementedDesiredStateManagerServer) SetDesiredState(context.Context, *SetDesiredStateReq) (*SetDesiredStateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDesiredState not implemented")
}
func (UnimplementedDesiredStateManagerServer) GetServiceDesiredStateConvergenceSummary(context.Context, *GetServiceDesiredStateConvergenceSummaryReq) (*GetServiceDesiredStateConvergenceSummaryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServiceDesiredStateConvergenceSummary not implemented")
}
func (UnimplementedDesiredStateManagerServer) GetServiceLastConvergedStates(context.Context, *GetServiceLastConvergedStateReq) (*GetServiceLastConvergedStateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServiceLastConvergedStates not implemented")
}
func (UnimplementedDesiredStateManagerServer) GetServiceDesiredStateHistory(context.Context, *GetServiceDesiredStateHistoryReq) (*GetServiceDesiredStateHistoryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServiceDesiredStateHistory not implemented")
}
func (UnimplementedDesiredStateManagerServer) GetDesiredState(context.Context, *GetDesiredStateReq) (*GetDesiredStateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDesiredState not implemented")
}
func (UnimplementedDesiredStateManagerServer) GetDesiredStateConvergenceSummary(context.Context, *GetDesiredStateConvergenceReq) (*GetDesiredStateConvergenceSummaryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDesiredStateConvergenceSummary not implemented")
}
func (UnimplementedDesiredStateManagerServer) ValidateDesiredState(context.Context, *ValidateDesiredStateReq) (*ValidateDesiredStateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateDesiredState not implemented")
}
func (UnimplementedDesiredStateManagerServer) SetManualApproval(context.Context, *SetManualApprovalReq) (*SetManualApprovalResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetManualApproval not implemented")
}
func (UnimplementedDesiredStateManagerServer) PromoteDelivery(context.Context, *PromoteDeliveryReq) (*PromoteDeliveryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PromoteDelivery not implemented")
}
func (UnimplementedDesiredStateManagerServer) BypassProtection(context.Context, *BypassProtectionReq) (*BypassProtectionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BypassProtection not implemented")
}
func (UnimplementedDesiredStateManagerServer) mustEmbedUnimplementedDesiredStateManagerServer() {}

// UnsafeDesiredStateManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DesiredStateManagerServer will
// result in compilation errors.
type UnsafeDesiredStateManagerServer interface {
	mustEmbedUnimplementedDesiredStateManagerServer()
}

func RegisterDesiredStateManagerServer(s grpc.ServiceRegistrar, srv DesiredStateManagerServer) {
	s.RegisterService(&DesiredStateManager_ServiceDesc, srv)
}

func _DesiredStateManager_SetDesiredState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDesiredStateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DesiredStateManagerServer).SetDesiredState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DesiredStateManager_SetDesiredState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DesiredStateManagerServer).SetDesiredState(ctx, req.(*SetDesiredStateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DesiredStateManager_GetServiceDesiredStateConvergenceSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceDesiredStateConvergenceSummaryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DesiredStateManagerServer).GetServiceDesiredStateConvergenceSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DesiredStateManager_GetServiceDesiredStateConvergenceSummary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DesiredStateManagerServer).GetServiceDesiredStateConvergenceSummary(ctx, req.(*GetServiceDesiredStateConvergenceSummaryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DesiredStateManager_GetServiceLastConvergedStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceLastConvergedStateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DesiredStateManagerServer).GetServiceLastConvergedStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DesiredStateManager_GetServiceLastConvergedStates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DesiredStateManagerServer).GetServiceLastConvergedStates(ctx, req.(*GetServiceLastConvergedStateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DesiredStateManager_GetServiceDesiredStateHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceDesiredStateHistoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DesiredStateManagerServer).GetServiceDesiredStateHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DesiredStateManager_GetServiceDesiredStateHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DesiredStateManagerServer).GetServiceDesiredStateHistory(ctx, req.(*GetServiceDesiredStateHistoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DesiredStateManager_GetDesiredState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDesiredStateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DesiredStateManagerServer).GetDesiredState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DesiredStateManager_GetDesiredState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DesiredStateManagerServer).GetDesiredState(ctx, req.(*GetDesiredStateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DesiredStateManager_GetDesiredStateConvergenceSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDesiredStateConvergenceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DesiredStateManagerServer).GetDesiredStateConvergenceSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DesiredStateManager_GetDesiredStateConvergenceSummary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DesiredStateManagerServer).GetDesiredStateConvergenceSummary(ctx, req.(*GetDesiredStateConvergenceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DesiredStateManager_ValidateDesiredState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateDesiredStateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DesiredStateManagerServer).ValidateDesiredState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DesiredStateManager_ValidateDesiredState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DesiredStateManagerServer).ValidateDesiredState(ctx, req.(*ValidateDesiredStateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DesiredStateManager_SetManualApproval_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetManualApprovalReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DesiredStateManagerServer).SetManualApproval(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DesiredStateManager_SetManualApproval_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DesiredStateManagerServer).SetManualApproval(ctx, req.(*SetManualApprovalReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DesiredStateManager_PromoteDelivery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromoteDeliveryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DesiredStateManagerServer).PromoteDelivery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DesiredStateManager_PromoteDelivery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DesiredStateManagerServer).PromoteDelivery(ctx, req.(*PromoteDeliveryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DesiredStateManager_BypassProtection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BypassProtectionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DesiredStateManagerServer).BypassProtection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DesiredStateManager_BypassProtection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DesiredStateManagerServer).BypassProtection(ctx, req.(*BypassProtectionReq))
	}
	return interceptor(ctx, in, info, handler)
}

// DesiredStateManager_ServiceDesc is the grpc.ServiceDesc for DesiredStateManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DesiredStateManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "prodvana.desired_state.DesiredStateManager",
	HandlerType: (*DesiredStateManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetDesiredState",
			Handler:    _DesiredStateManager_SetDesiredState_Handler,
		},
		{
			MethodName: "GetServiceDesiredStateConvergenceSummary",
			Handler:    _DesiredStateManager_GetServiceDesiredStateConvergenceSummary_Handler,
		},
		{
			MethodName: "GetServiceLastConvergedStates",
			Handler:    _DesiredStateManager_GetServiceLastConvergedStates_Handler,
		},
		{
			MethodName: "GetServiceDesiredStateHistory",
			Handler:    _DesiredStateManager_GetServiceDesiredStateHistory_Handler,
		},
		{
			MethodName: "GetDesiredState",
			Handler:    _DesiredStateManager_GetDesiredState_Handler,
		},
		{
			MethodName: "GetDesiredStateConvergenceSummary",
			Handler:    _DesiredStateManager_GetDesiredStateConvergenceSummary_Handler,
		},
		{
			MethodName: "ValidateDesiredState",
			Handler:    _DesiredStateManager_ValidateDesiredState_Handler,
		},
		{
			MethodName: "SetManualApproval",
			Handler:    _DesiredStateManager_SetManualApproval_Handler,
		},
		{
			MethodName: "PromoteDelivery",
			Handler:    _DesiredStateManager_PromoteDelivery_Handler,
		},
		{
			MethodName: "BypassProtection",
			Handler:    _DesiredStateManager_BypassProtection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "prodvana/desired_state/manager.proto",
}
