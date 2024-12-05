// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: spaceone/api/alert_manager/v1/user_channel.proto

package v1

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UserChannel_Create_FullMethodName  = "/spaceone.api.alert_manager.v1.UserChannel/create"
	UserChannel_Update_FullMethodName  = "/spaceone.api.alert_manager.v1.UserChannel/update"
	UserChannel_Enable_FullMethodName  = "/spaceone.api.alert_manager.v1.UserChannel/enable"
	UserChannel_Disable_FullMethodName = "/spaceone.api.alert_manager.v1.UserChannel/disable"
	UserChannel_Delete_FullMethodName  = "/spaceone.api.alert_manager.v1.UserChannel/delete"
	UserChannel_Get_FullMethodName     = "/spaceone.api.alert_manager.v1.UserChannel/get"
	UserChannel_List_FullMethodName    = "/spaceone.api.alert_manager.v1.UserChannel/list"
	UserChannel_Stat_FullMethodName    = "/spaceone.api.alert_manager.v1.UserChannel/stat"
)

// UserChannelClient is the client API for UserChannel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserChannelClient interface {
	Create(ctx context.Context, in *UserChannelCreateRequest, opts ...grpc.CallOption) (*UserChannelInfo, error)
	Update(ctx context.Context, in *UserChannelUpdateRequest, opts ...grpc.CallOption) (*UserChannelInfo, error)
	Enable(ctx context.Context, in *UserChannelRequest, opts ...grpc.CallOption) (*UserChannelInfo, error)
	Disable(ctx context.Context, in *UserChannelRequest, opts ...grpc.CallOption) (*UserChannelInfo, error)
	Delete(ctx context.Context, in *UserChannelRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Get(ctx context.Context, in *UserChannelRequest, opts ...grpc.CallOption) (*UserChannelInfo, error)
	List(ctx context.Context, in *UserChannelSearchQuery, opts ...grpc.CallOption) (*UserChannelsInfo, error)
	Stat(ctx context.Context, in *UserChannelStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type userChannelClient struct {
	cc grpc.ClientConnInterface
}

func NewUserChannelClient(cc grpc.ClientConnInterface) UserChannelClient {
	return &userChannelClient{cc}
}

func (c *userChannelClient) Create(ctx context.Context, in *UserChannelCreateRequest, opts ...grpc.CallOption) (*UserChannelInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserChannelInfo)
	err := c.cc.Invoke(ctx, UserChannel_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userChannelClient) Update(ctx context.Context, in *UserChannelUpdateRequest, opts ...grpc.CallOption) (*UserChannelInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserChannelInfo)
	err := c.cc.Invoke(ctx, UserChannel_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userChannelClient) Enable(ctx context.Context, in *UserChannelRequest, opts ...grpc.CallOption) (*UserChannelInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserChannelInfo)
	err := c.cc.Invoke(ctx, UserChannel_Enable_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userChannelClient) Disable(ctx context.Context, in *UserChannelRequest, opts ...grpc.CallOption) (*UserChannelInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserChannelInfo)
	err := c.cc.Invoke(ctx, UserChannel_Disable_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userChannelClient) Delete(ctx context.Context, in *UserChannelRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, UserChannel_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userChannelClient) Get(ctx context.Context, in *UserChannelRequest, opts ...grpc.CallOption) (*UserChannelInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserChannelInfo)
	err := c.cc.Invoke(ctx, UserChannel_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userChannelClient) List(ctx context.Context, in *UserChannelSearchQuery, opts ...grpc.CallOption) (*UserChannelsInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserChannelsInfo)
	err := c.cc.Invoke(ctx, UserChannel_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userChannelClient) Stat(ctx context.Context, in *UserChannelStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, UserChannel_Stat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserChannelServer is the server API for UserChannel service.
// All implementations must embed UnimplementedUserChannelServer
// for forward compatibility.
type UserChannelServer interface {
	Create(context.Context, *UserChannelCreateRequest) (*UserChannelInfo, error)
	Update(context.Context, *UserChannelUpdateRequest) (*UserChannelInfo, error)
	Enable(context.Context, *UserChannelRequest) (*UserChannelInfo, error)
	Disable(context.Context, *UserChannelRequest) (*UserChannelInfo, error)
	Delete(context.Context, *UserChannelRequest) (*empty.Empty, error)
	Get(context.Context, *UserChannelRequest) (*UserChannelInfo, error)
	List(context.Context, *UserChannelSearchQuery) (*UserChannelsInfo, error)
	Stat(context.Context, *UserChannelStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedUserChannelServer()
}

// UnimplementedUserChannelServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserChannelServer struct{}

func (UnimplementedUserChannelServer) Create(context.Context, *UserChannelCreateRequest) (*UserChannelInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserChannelServer) Update(context.Context, *UserChannelUpdateRequest) (*UserChannelInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUserChannelServer) Enable(context.Context, *UserChannelRequest) (*UserChannelInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enable not implemented")
}
func (UnimplementedUserChannelServer) Disable(context.Context, *UserChannelRequest) (*UserChannelInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disable not implemented")
}
func (UnimplementedUserChannelServer) Delete(context.Context, *UserChannelRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedUserChannelServer) Get(context.Context, *UserChannelRequest) (*UserChannelInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedUserChannelServer) List(context.Context, *UserChannelSearchQuery) (*UserChannelsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedUserChannelServer) Stat(context.Context, *UserChannelStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedUserChannelServer) mustEmbedUnimplementedUserChannelServer() {}
func (UnimplementedUserChannelServer) testEmbeddedByValue()                     {}

// UnsafeUserChannelServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserChannelServer will
// result in compilation errors.
type UnsafeUserChannelServer interface {
	mustEmbedUnimplementedUserChannelServer()
}

func RegisterUserChannelServer(s grpc.ServiceRegistrar, srv UserChannelServer) {
	// If the following call pancis, it indicates UnimplementedUserChannelServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserChannel_ServiceDesc, srv)
}

func _UserChannel_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserChannelCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserChannelServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserChannel_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserChannelServer).Create(ctx, req.(*UserChannelCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserChannel_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserChannelUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserChannelServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserChannel_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserChannelServer).Update(ctx, req.(*UserChannelUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserChannel_Enable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserChannelServer).Enable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserChannel_Enable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserChannelServer).Enable(ctx, req.(*UserChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserChannel_Disable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserChannelServer).Disable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserChannel_Disable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserChannelServer).Disable(ctx, req.(*UserChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserChannel_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserChannelServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserChannel_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserChannelServer).Delete(ctx, req.(*UserChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserChannel_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserChannelServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserChannel_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserChannelServer).Get(ctx, req.(*UserChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserChannel_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserChannelSearchQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserChannelServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserChannel_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserChannelServer).List(ctx, req.(*UserChannelSearchQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserChannel_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserChannelStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserChannelServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserChannel_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserChannelServer).Stat(ctx, req.(*UserChannelStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// UserChannel_ServiceDesc is the grpc.ServiceDesc for UserChannel service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserChannel_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.alert_manager.v1.UserChannel",
	HandlerType: (*UserChannelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _UserChannel_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _UserChannel_Update_Handler,
		},
		{
			MethodName: "enable",
			Handler:    _UserChannel_Enable_Handler,
		},
		{
			MethodName: "disable",
			Handler:    _UserChannel_Disable_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _UserChannel_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _UserChannel_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _UserChannel_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _UserChannel_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/alert_manager/v1/user_channel.proto",
}
