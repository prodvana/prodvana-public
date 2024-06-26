// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.10
// source: prodvana/recipe/manager.proto

package recipe

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
	RecipeManager_ConfigureRecipe_FullMethodName         = "/prodvana.recipe.RecipeManager/ConfigureRecipe"
	RecipeManager_ValidateConfigureRecipe_FullMethodName = "/prodvana.recipe.RecipeManager/ValidateConfigureRecipe"
	RecipeManager_ListRecipes_FullMethodName             = "/prodvana.recipe.RecipeManager/ListRecipes"
	RecipeManager_GetRecipe_FullMethodName               = "/prodvana.recipe.RecipeManager/GetRecipe"
	RecipeManager_GetRecipeConfig_FullMethodName         = "/prodvana.recipe.RecipeManager/GetRecipeConfig"
	RecipeManager_ApplyRecipeParameters_FullMethodName   = "/prodvana.recipe.RecipeManager/ApplyRecipeParameters"
)

// RecipeManagerClient is the client API for RecipeManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecipeManagerClient interface {
	ConfigureRecipe(ctx context.Context, in *ConfigureRecipeReq, opts ...grpc.CallOption) (*ConfigureRecipeResp, error)
	ValidateConfigureRecipe(ctx context.Context, in *ConfigureRecipeReq, opts ...grpc.CallOption) (*ValidateConfigureRecipeResp, error)
	ListRecipes(ctx context.Context, in *ListRecipesReq, opts ...grpc.CallOption) (*ListRecipesResp, error)
	GetRecipe(ctx context.Context, in *GetRecipeReq, opts ...grpc.CallOption) (*GetRecipeResp, error)
	GetRecipeConfig(ctx context.Context, in *GetRecipeConfigReq, opts ...grpc.CallOption) (*GetRecipeConfigResp, error)
	ApplyRecipeParameters(ctx context.Context, in *ApplyRecipeParametersReq, opts ...grpc.CallOption) (*ApplyRecipeParametersResp, error)
}

type recipeManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewRecipeManagerClient(cc grpc.ClientConnInterface) RecipeManagerClient {
	return &recipeManagerClient{cc}
}

func (c *recipeManagerClient) ConfigureRecipe(ctx context.Context, in *ConfigureRecipeReq, opts ...grpc.CallOption) (*ConfigureRecipeResp, error) {
	out := new(ConfigureRecipeResp)
	err := c.cc.Invoke(ctx, RecipeManager_ConfigureRecipe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeManagerClient) ValidateConfigureRecipe(ctx context.Context, in *ConfigureRecipeReq, opts ...grpc.CallOption) (*ValidateConfigureRecipeResp, error) {
	out := new(ValidateConfigureRecipeResp)
	err := c.cc.Invoke(ctx, RecipeManager_ValidateConfigureRecipe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeManagerClient) ListRecipes(ctx context.Context, in *ListRecipesReq, opts ...grpc.CallOption) (*ListRecipesResp, error) {
	out := new(ListRecipesResp)
	err := c.cc.Invoke(ctx, RecipeManager_ListRecipes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeManagerClient) GetRecipe(ctx context.Context, in *GetRecipeReq, opts ...grpc.CallOption) (*GetRecipeResp, error) {
	out := new(GetRecipeResp)
	err := c.cc.Invoke(ctx, RecipeManager_GetRecipe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeManagerClient) GetRecipeConfig(ctx context.Context, in *GetRecipeConfigReq, opts ...grpc.CallOption) (*GetRecipeConfigResp, error) {
	out := new(GetRecipeConfigResp)
	err := c.cc.Invoke(ctx, RecipeManager_GetRecipeConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeManagerClient) ApplyRecipeParameters(ctx context.Context, in *ApplyRecipeParametersReq, opts ...grpc.CallOption) (*ApplyRecipeParametersResp, error) {
	out := new(ApplyRecipeParametersResp)
	err := c.cc.Invoke(ctx, RecipeManager_ApplyRecipeParameters_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecipeManagerServer is the server API for RecipeManager service.
// All implementations must embed UnimplementedRecipeManagerServer
// for forward compatibility
type RecipeManagerServer interface {
	ConfigureRecipe(context.Context, *ConfigureRecipeReq) (*ConfigureRecipeResp, error)
	ValidateConfigureRecipe(context.Context, *ConfigureRecipeReq) (*ValidateConfigureRecipeResp, error)
	ListRecipes(context.Context, *ListRecipesReq) (*ListRecipesResp, error)
	GetRecipe(context.Context, *GetRecipeReq) (*GetRecipeResp, error)
	GetRecipeConfig(context.Context, *GetRecipeConfigReq) (*GetRecipeConfigResp, error)
	ApplyRecipeParameters(context.Context, *ApplyRecipeParametersReq) (*ApplyRecipeParametersResp, error)
	mustEmbedUnimplementedRecipeManagerServer()
}

// UnimplementedRecipeManagerServer must be embedded to have forward compatible implementations.
type UnimplementedRecipeManagerServer struct {
}

func (UnimplementedRecipeManagerServer) ConfigureRecipe(context.Context, *ConfigureRecipeReq) (*ConfigureRecipeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureRecipe not implemented")
}
func (UnimplementedRecipeManagerServer) ValidateConfigureRecipe(context.Context, *ConfigureRecipeReq) (*ValidateConfigureRecipeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateConfigureRecipe not implemented")
}
func (UnimplementedRecipeManagerServer) ListRecipes(context.Context, *ListRecipesReq) (*ListRecipesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRecipes not implemented")
}
func (UnimplementedRecipeManagerServer) GetRecipe(context.Context, *GetRecipeReq) (*GetRecipeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecipe not implemented")
}
func (UnimplementedRecipeManagerServer) GetRecipeConfig(context.Context, *GetRecipeConfigReq) (*GetRecipeConfigResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecipeConfig not implemented")
}
func (UnimplementedRecipeManagerServer) ApplyRecipeParameters(context.Context, *ApplyRecipeParametersReq) (*ApplyRecipeParametersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyRecipeParameters not implemented")
}
func (UnimplementedRecipeManagerServer) mustEmbedUnimplementedRecipeManagerServer() {}

// UnsafeRecipeManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecipeManagerServer will
// result in compilation errors.
type UnsafeRecipeManagerServer interface {
	mustEmbedUnimplementedRecipeManagerServer()
}

func RegisterRecipeManagerServer(s grpc.ServiceRegistrar, srv RecipeManagerServer) {
	s.RegisterService(&RecipeManager_ServiceDesc, srv)
}

func _RecipeManager_ConfigureRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureRecipeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeManagerServer).ConfigureRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecipeManager_ConfigureRecipe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeManagerServer).ConfigureRecipe(ctx, req.(*ConfigureRecipeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeManager_ValidateConfigureRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureRecipeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeManagerServer).ValidateConfigureRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecipeManager_ValidateConfigureRecipe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeManagerServer).ValidateConfigureRecipe(ctx, req.(*ConfigureRecipeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeManager_ListRecipes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRecipesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeManagerServer).ListRecipes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecipeManager_ListRecipes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeManagerServer).ListRecipes(ctx, req.(*ListRecipesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeManager_GetRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecipeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeManagerServer).GetRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecipeManager_GetRecipe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeManagerServer).GetRecipe(ctx, req.(*GetRecipeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeManager_GetRecipeConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecipeConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeManagerServer).GetRecipeConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecipeManager_GetRecipeConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeManagerServer).GetRecipeConfig(ctx, req.(*GetRecipeConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeManager_ApplyRecipeParameters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyRecipeParametersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeManagerServer).ApplyRecipeParameters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecipeManager_ApplyRecipeParameters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeManagerServer).ApplyRecipeParameters(ctx, req.(*ApplyRecipeParametersReq))
	}
	return interceptor(ctx, in, info, handler)
}

// RecipeManager_ServiceDesc is the grpc.ServiceDesc for RecipeManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecipeManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "prodvana.recipe.RecipeManager",
	HandlerType: (*RecipeManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConfigureRecipe",
			Handler:    _RecipeManager_ConfigureRecipe_Handler,
		},
		{
			MethodName: "ValidateConfigureRecipe",
			Handler:    _RecipeManager_ValidateConfigureRecipe_Handler,
		},
		{
			MethodName: "ListRecipes",
			Handler:    _RecipeManager_ListRecipes_Handler,
		},
		{
			MethodName: "GetRecipe",
			Handler:    _RecipeManager_GetRecipe_Handler,
		},
		{
			MethodName: "GetRecipeConfig",
			Handler:    _RecipeManager_GetRecipeConfig_Handler,
		},
		{
			MethodName: "ApplyRecipeParameters",
			Handler:    _RecipeManager_ApplyRecipeParameters_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "prodvana/recipe/manager.proto",
}
