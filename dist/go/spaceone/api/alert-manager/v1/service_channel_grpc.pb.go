// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: spaceone/api/alert_manager/v1/service_channel.proto

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
	ServiceChannel_Create_FullMethodName               = "/spaceone.api.alert_manager.v1.ServiceChannel/create"
	ServiceChannel_CreateForwardChannel_FullMethodName = "/spaceone.api.alert_manager.v1.ServiceChannel/create_forward_channel"
	ServiceChannel_Update_FullMethodName               = "/spaceone.api.alert_manager.v1.ServiceChannel/update"
	ServiceChannel_Enable_FullMethodName               = "/spaceone.api.alert_manager.v1.ServiceChannel/enable"
	ServiceChannel_Disable_FullMethodName              = "/spaceone.api.alert_manager.v1.ServiceChannel/disable"
	ServiceChannel_Delete_FullMethodName               = "/spaceone.api.alert_manager.v1.ServiceChannel/delete"
	ServiceChannel_Get_FullMethodName                  = "/spaceone.api.alert_manager.v1.ServiceChannel/get"
	ServiceChannel_List_FullMethodName                 = "/spaceone.api.alert_manager.v1.ServiceChannel/list"
	ServiceChannel_Stat_FullMethodName                 = "/spaceone.api.alert_manager.v1.ServiceChannel/stat"
)

// ServiceChannelClient is the client API for ServiceChannel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceChannelClient interface {
	Create(ctx context.Context, in *ServiceChannelCreateRequest, opts ...grpc.CallOption) (*ServiceChannelInfo, error)
	CreateForwardChannel(ctx context.Context, in *ServiceChannelCreateForwardChannelRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Update(ctx context.Context, in *ServiceChannelUpdateRequest, opts ...grpc.CallOption) (*ServiceChannelInfo, error)
	Enable(ctx context.Context, in *ServiceChannelRequest, opts ...grpc.CallOption) (*ServiceChannelInfo, error)
	Disable(ctx context.Context, in *ServiceChannelRequest, opts ...grpc.CallOption) (*ServiceChannelInfo, error)
	Delete(ctx context.Context, in *ServiceChannelRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Get(ctx context.Context, in *ServiceChannelRequest, opts ...grpc.CallOption) (*ServiceChannelInfo, error)
	List(ctx context.Context, in *ServiceChannelSearchQuery, opts ...grpc.CallOption) (*ServiceChannelsInfo, error)
	Stat(ctx context.Context, in *ServiceChannelStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type serviceChannelClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceChannelClient(cc grpc.ClientConnInterface) ServiceChannelClient {
	return &serviceChannelClient{cc}
}

func (c *serviceChannelClient) Create(ctx context.Context, in *ServiceChannelCreateRequest, opts ...grpc.CallOption) (*ServiceChannelInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ServiceChannelInfo)
	err := c.cc.Invoke(ctx, ServiceChannel_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceChannelClient) CreateForwardChannel(ctx context.Context, in *ServiceChannelCreateForwardChannelRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, ServiceChannel_CreateForwardChannel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceChannelClient) Update(ctx context.Context, in *ServiceChannelUpdateRequest, opts ...grpc.CallOption) (*ServiceChannelInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ServiceChannelInfo)
	err := c.cc.Invoke(ctx, ServiceChannel_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceChannelClient) Enable(ctx context.Context, in *ServiceChannelRequest, opts ...grpc.CallOption) (*ServiceChannelInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ServiceChannelInfo)
	err := c.cc.Invoke(ctx, ServiceChannel_Enable_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceChannelClient) Disable(ctx context.Context, in *ServiceChannelRequest, opts ...grpc.CallOption) (*ServiceChannelInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ServiceChannelInfo)
	err := c.cc.Invoke(ctx, ServiceChannel_Disable_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceChannelClient) Delete(ctx context.Context, in *ServiceChannelRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, ServiceChannel_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceChannelClient) Get(ctx context.Context, in *ServiceChannelRequest, opts ...grpc.CallOption) (*ServiceChannelInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ServiceChannelInfo)
	err := c.cc.Invoke(ctx, ServiceChannel_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceChannelClient) List(ctx context.Context, in *ServiceChannelSearchQuery, opts ...grpc.CallOption) (*ServiceChannelsInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ServiceChannelsInfo)
	err := c.cc.Invoke(ctx, ServiceChannel_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceChannelClient) Stat(ctx context.Context, in *ServiceChannelStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, ServiceChannel_Stat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceChannelServer is the server API for ServiceChannel service.
// All implementations must embed UnimplementedServiceChannelServer
// for forward compatibility.
type ServiceChannelServer interface {
	Create(context.Context, *ServiceChannelCreateRequest) (*ServiceChannelInfo, error)
	CreateForwardChannel(context.Context, *ServiceChannelCreateForwardChannelRequest) (*empty.Empty, error)
	Update(context.Context, *ServiceChannelUpdateRequest) (*ServiceChannelInfo, error)
	Enable(context.Context, *ServiceChannelRequest) (*ServiceChannelInfo, error)
	Disable(context.Context, *ServiceChannelRequest) (*ServiceChannelInfo, error)
	Delete(context.Context, *ServiceChannelRequest) (*empty.Empty, error)
	Get(context.Context, *ServiceChannelRequest) (*ServiceChannelInfo, error)
	List(context.Context, *ServiceChannelSearchQuery) (*ServiceChannelsInfo, error)
	Stat(context.Context, *ServiceChannelStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedServiceChannelServer()
}

// UnimplementedServiceChannelServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedServiceChannelServer struct{}

func (UnimplementedServiceChannelServer) Create(context.Context, *ServiceChannelCreateRequest) (*ServiceChannelInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedServiceChannelServer) CreateForwardChannel(context.Context, *ServiceChannelCreateForwardChannelRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateForwardChannel not implemented")
}
func (UnimplementedServiceChannelServer) Update(context.Context, *ServiceChannelUpdateRequest) (*ServiceChannelInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedServiceChannelServer) Enable(context.Context, *ServiceChannelRequest) (*ServiceChannelInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enable not implemented")
}
func (UnimplementedServiceChannelServer) Disable(context.Context, *ServiceChannelRequest) (*ServiceChannelInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disable not implemented")
}
func (UnimplementedServiceChannelServer) Delete(context.Context, *ServiceChannelRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedServiceChannelServer) Get(context.Context, *ServiceChannelRequest) (*ServiceChannelInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedServiceChannelServer) List(context.Context, *ServiceChannelSearchQuery) (*ServiceChannelsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedServiceChannelServer) Stat(context.Context, *ServiceChannelStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedServiceChannelServer) mustEmbedUnimplementedServiceChannelServer() {}
func (UnimplementedServiceChannelServer) testEmbeddedByValue()                        {}

// UnsafeServiceChannelServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceChannelServer will
// result in compilation errors.
type UnsafeServiceChannelServer interface {
	mustEmbedUnimplementedServiceChannelServer()
}

func RegisterServiceChannelServer(s grpc.ServiceRegistrar, srv ServiceChannelServer) {
	// If the following call pancis, it indicates UnimplementedServiceChannelServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ServiceChannel_ServiceDesc, srv)
}

func _ServiceChannel_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceChannelCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceChannelServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceChannel_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceChannelServer).Create(ctx, req.(*ServiceChannelCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceChannel_CreateForwardChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceChannelCreateForwardChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceChannelServer).CreateForwardChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceChannel_CreateForwardChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceChannelServer).CreateForwardChannel(ctx, req.(*ServiceChannelCreateForwardChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceChannel_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceChannelUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceChannelServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceChannel_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceChannelServer).Update(ctx, req.(*ServiceChannelUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceChannel_Enable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceChannelServer).Enable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceChannel_Enable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceChannelServer).Enable(ctx, req.(*ServiceChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceChannel_Disable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceChannelServer).Disable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceChannel_Disable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceChannelServer).Disable(ctx, req.(*ServiceChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceChannel_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceChannelServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceChannel_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceChannelServer).Delete(ctx, req.(*ServiceChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceChannel_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceChannelServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceChannel_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceChannelServer).Get(ctx, req.(*ServiceChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceChannel_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceChannelSearchQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceChannelServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceChannel_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceChannelServer).List(ctx, req.(*ServiceChannelSearchQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceChannel_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceChannelStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceChannelServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceChannel_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceChannelServer).Stat(ctx, req.(*ServiceChannelStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceChannel_ServiceDesc is the grpc.ServiceDesc for ServiceChannel service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceChannel_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.alert_manager.v1.ServiceChannel",
	HandlerType: (*ServiceChannelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _ServiceChannel_Create_Handler,
		},
		{
			MethodName: "create_forward_channel",
			Handler:    _ServiceChannel_CreateForwardChannel_Handler,
		},
		{
			MethodName: "update",
			Handler:    _ServiceChannel_Update_Handler,
		},
		{
			MethodName: "enable",
			Handler:    _ServiceChannel_Enable_Handler,
		},
		{
			MethodName: "disable",
			Handler:    _ServiceChannel_Disable_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _ServiceChannel_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _ServiceChannel_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _ServiceChannel_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _ServiceChannel_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/alert_manager/v1/service_channel.proto",
}