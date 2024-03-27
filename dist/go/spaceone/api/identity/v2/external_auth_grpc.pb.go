// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: spaceone/api/identity/v2/external_auth.proto

package v2

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
	ExternalAuth_Set_FullMethodName   = "/spaceone.api.identity.v2.ExternalAuth/set"
	ExternalAuth_Unset_FullMethodName = "/spaceone.api.identity.v2.ExternalAuth/unset"
	ExternalAuth_Get_FullMethodName   = "/spaceone.api.identity.v2.ExternalAuth/get"
)

// ExternalAuthClient is the client API for ExternalAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExternalAuthClient interface {
	Set(ctx context.Context, in *SetExternalAuthRequest, opts ...grpc.CallOption) (*ExternalAuthInfo, error)
	Unset(ctx context.Context, in *ExternalAuthRequest, opts ...grpc.CallOption) (*ExternalAuthInfo, error)
	Get(ctx context.Context, in *ExternalAuthRequest, opts ...grpc.CallOption) (*ExternalAuthInfo, error)
}

type externalAuthClient struct {
	cc grpc.ClientConnInterface
}

func NewExternalAuthClient(cc grpc.ClientConnInterface) ExternalAuthClient {
	return &externalAuthClient{cc}
}

func (c *externalAuthClient) Set(ctx context.Context, in *SetExternalAuthRequest, opts ...grpc.CallOption) (*ExternalAuthInfo, error) {
	out := new(ExternalAuthInfo)
	err := c.cc.Invoke(ctx, ExternalAuth_Set_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalAuthClient) Unset(ctx context.Context, in *ExternalAuthRequest, opts ...grpc.CallOption) (*ExternalAuthInfo, error) {
	out := new(ExternalAuthInfo)
	err := c.cc.Invoke(ctx, ExternalAuth_Unset_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalAuthClient) Get(ctx context.Context, in *ExternalAuthRequest, opts ...grpc.CallOption) (*ExternalAuthInfo, error) {
	out := new(ExternalAuthInfo)
	err := c.cc.Invoke(ctx, ExternalAuth_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExternalAuthServer is the server API for ExternalAuth service.
// All implementations must embed UnimplementedExternalAuthServer
// for forward compatibility
type ExternalAuthServer interface {
	Set(context.Context, *SetExternalAuthRequest) (*ExternalAuthInfo, error)
	Unset(context.Context, *ExternalAuthRequest) (*ExternalAuthInfo, error)
	Get(context.Context, *ExternalAuthRequest) (*ExternalAuthInfo, error)
	mustEmbedUnimplementedExternalAuthServer()
}

// UnimplementedExternalAuthServer must be embedded to have forward compatible implementations.
type UnimplementedExternalAuthServer struct {
}

func (UnimplementedExternalAuthServer) Set(context.Context, *SetExternalAuthRequest) (*ExternalAuthInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedExternalAuthServer) Unset(context.Context, *ExternalAuthRequest) (*ExternalAuthInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unset not implemented")
}
func (UnimplementedExternalAuthServer) Get(context.Context, *ExternalAuthRequest) (*ExternalAuthInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedExternalAuthServer) mustEmbedUnimplementedExternalAuthServer() {}

// UnsafeExternalAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExternalAuthServer will
// result in compilation errors.
type UnsafeExternalAuthServer interface {
	mustEmbedUnimplementedExternalAuthServer()
}

func RegisterExternalAuthServer(s grpc.ServiceRegistrar, srv ExternalAuthServer) {
	s.RegisterService(&ExternalAuth_ServiceDesc, srv)
}

func _ExternalAuth_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetExternalAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalAuthServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalAuth_Set_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalAuthServer).Set(ctx, req.(*SetExternalAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalAuth_Unset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExternalAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalAuthServer).Unset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalAuth_Unset_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalAuthServer).Unset(ctx, req.(*ExternalAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalAuth_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExternalAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalAuthServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalAuth_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalAuthServer).Get(ctx, req.(*ExternalAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExternalAuth_ServiceDesc is the grpc.ServiceDesc for ExternalAuth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExternalAuth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.identity.v2.ExternalAuth",
	HandlerType: (*ExternalAuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "set",
			Handler:    _ExternalAuth_Set_Handler,
		},
		{
			MethodName: "unset",
			Handler:    _ExternalAuth_Unset_Handler,
		},
		{
			MethodName: "get",
			Handler:    _ExternalAuth_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/identity/v2/external_auth.proto",
}
