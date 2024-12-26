// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: spaceone/api/config/v1/shared_config.proto

package v1

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SharedConfig_Create_FullMethodName = "/spaceone.api.config.v1.SharedConfig/create"
	SharedConfig_Update_FullMethodName = "/spaceone.api.config.v1.SharedConfig/update"
	SharedConfig_Delete_FullMethodName = "/spaceone.api.config.v1.SharedConfig/delete"
	SharedConfig_Get_FullMethodName    = "/spaceone.api.config.v1.SharedConfig/get"
	SharedConfig_List_FullMethodName   = "/spaceone.api.config.v1.SharedConfig/list"
)

// SharedConfigClient is the client API for SharedConfig service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SharedConfigClient interface {
	Create(ctx context.Context, in *CreateSharedConfigRequest, opts ...grpc.CallOption) (*SharedConfigInfo, error)
	Update(ctx context.Context, in *UpdateSharedConfigRequest, opts ...grpc.CallOption) (*SharedConfigInfo, error)
	Delete(ctx context.Context, in *SharedConfigRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Get(ctx context.Context, in *SharedConfigSearchQuery, opts ...grpc.CallOption) (*SharedConfigInfo, error)
	List(ctx context.Context, in *SharedConfigSearchQuery, opts ...grpc.CallOption) (*SharedConfigsInfo, error)
}

type sharedConfigClient struct {
	cc grpc.ClientConnInterface
}

func NewSharedConfigClient(cc grpc.ClientConnInterface) SharedConfigClient {
	return &sharedConfigClient{cc}
}

func (c *sharedConfigClient) Create(ctx context.Context, in *CreateSharedConfigRequest, opts ...grpc.CallOption) (*SharedConfigInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SharedConfigInfo)
	err := c.cc.Invoke(ctx, SharedConfig_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sharedConfigClient) Update(ctx context.Context, in *UpdateSharedConfigRequest, opts ...grpc.CallOption) (*SharedConfigInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SharedConfigInfo)
	err := c.cc.Invoke(ctx, SharedConfig_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sharedConfigClient) Delete(ctx context.Context, in *SharedConfigRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, SharedConfig_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sharedConfigClient) Get(ctx context.Context, in *SharedConfigSearchQuery, opts ...grpc.CallOption) (*SharedConfigInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SharedConfigInfo)
	err := c.cc.Invoke(ctx, SharedConfig_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sharedConfigClient) List(ctx context.Context, in *SharedConfigSearchQuery, opts ...grpc.CallOption) (*SharedConfigsInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SharedConfigsInfo)
	err := c.cc.Invoke(ctx, SharedConfig_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SharedConfigServer is the server API for SharedConfig service.
// All implementations must embed UnimplementedSharedConfigServer
// for forward compatibility.
type SharedConfigServer interface {
	Create(context.Context, *CreateSharedConfigRequest) (*SharedConfigInfo, error)
	Update(context.Context, *UpdateSharedConfigRequest) (*SharedConfigInfo, error)
	Delete(context.Context, *SharedConfigRequest) (*empty.Empty, error)
	Get(context.Context, *SharedConfigSearchQuery) (*SharedConfigInfo, error)
	List(context.Context, *SharedConfigSearchQuery) (*SharedConfigsInfo, error)
	mustEmbedUnimplementedSharedConfigServer()
}

// UnimplementedSharedConfigServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSharedConfigServer struct{}

func (UnimplementedSharedConfigServer) Create(context.Context, *CreateSharedConfigRequest) (*SharedConfigInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSharedConfigServer) Update(context.Context, *UpdateSharedConfigRequest) (*SharedConfigInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSharedConfigServer) Delete(context.Context, *SharedConfigRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSharedConfigServer) Get(context.Context, *SharedConfigSearchQuery) (*SharedConfigInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSharedConfigServer) List(context.Context, *SharedConfigSearchQuery) (*SharedConfigsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedSharedConfigServer) mustEmbedUnimplementedSharedConfigServer() {}
func (UnimplementedSharedConfigServer) testEmbeddedByValue()                      {}

// UnsafeSharedConfigServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SharedConfigServer will
// result in compilation errors.
type UnsafeSharedConfigServer interface {
	mustEmbedUnimplementedSharedConfigServer()
}

func RegisterSharedConfigServer(s grpc.ServiceRegistrar, srv SharedConfigServer) {
	// If the following call pancis, it indicates UnimplementedSharedConfigServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SharedConfig_ServiceDesc, srv)
}

func _SharedConfig_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSharedConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharedConfigServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SharedConfig_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharedConfigServer).Create(ctx, req.(*CreateSharedConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SharedConfig_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSharedConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharedConfigServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SharedConfig_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharedConfigServer).Update(ctx, req.(*UpdateSharedConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SharedConfig_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SharedConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharedConfigServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SharedConfig_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharedConfigServer).Delete(ctx, req.(*SharedConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SharedConfig_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SharedConfigSearchQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharedConfigServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SharedConfig_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharedConfigServer).Get(ctx, req.(*SharedConfigSearchQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _SharedConfig_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SharedConfigSearchQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharedConfigServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SharedConfig_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharedConfigServer).List(ctx, req.(*SharedConfigSearchQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// SharedConfig_ServiceDesc is the grpc.ServiceDesc for SharedConfig service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SharedConfig_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.config.v1.SharedConfig",
	HandlerType: (*SharedConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _SharedConfig_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _SharedConfig_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _SharedConfig_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _SharedConfig_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _SharedConfig_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/config/v1/shared_config.proto",
}