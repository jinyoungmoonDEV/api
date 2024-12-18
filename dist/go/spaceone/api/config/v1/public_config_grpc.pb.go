// DomainConfig API which configure environments for domain

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: spaceone/api/config/v1/public_config.proto

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
	PublicConfig_Create_FullMethodName               = "/spaceone.api.config.v1.PublicConfig/create"
	PublicConfig_Update_FullMethodName               = "/spaceone.api.config.v1.PublicConfig/update"
	PublicConfig_Delete_FullMethodName               = "/spaceone.api.config.v1.PublicConfig/delete"
	PublicConfig_Get_FullMethodName                  = "/spaceone.api.config.v1.PublicConfig/get"
	PublicConfig_GetAccessibleConfigs_FullMethodName = "/spaceone.api.config.v1.PublicConfig/get_accessible_configs"
	PublicConfig_List_FullMethodName                 = "/spaceone.api.config.v1.PublicConfig/list"
	PublicConfig_Stat_FullMethodName                 = "/spaceone.api.config.v1.PublicConfig/stat"
)

// PublicConfigClient is the client API for PublicConfig service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PublicConfigClient interface {
	Create(ctx context.Context, in *CreatePublicConfigRequest, opts ...grpc.CallOption) (*PublicConfigInfo, error)
	Update(ctx context.Context, in *UpdatePublicConfigRequest, opts ...grpc.CallOption) (*PublicConfigInfo, error)
	Delete(ctx context.Context, in *PublicConfigRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Get(ctx context.Context, in *PublicConfigSearchQuery, opts ...grpc.CallOption) (*PublicConfigInfo, error)
	// This API for retrieving domain scoped configs that are accessible to users.
	GetAccessibleConfigs(ctx context.Context, in *PublicConfigSearchQuery, opts ...grpc.CallOption) (*PublicConfigsInfo, error)
	List(ctx context.Context, in *PublicConfigSearchQuery, opts ...grpc.CallOption) (*PublicConfigsInfo, error)
	Stat(ctx context.Context, in *PublicConfigStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type publicConfigClient struct {
	cc grpc.ClientConnInterface
}

func NewPublicConfigClient(cc grpc.ClientConnInterface) PublicConfigClient {
	return &publicConfigClient{cc}
}

func (c *publicConfigClient) Create(ctx context.Context, in *CreatePublicConfigRequest, opts ...grpc.CallOption) (*PublicConfigInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublicConfigInfo)
	err := c.cc.Invoke(ctx, PublicConfig_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicConfigClient) Update(ctx context.Context, in *UpdatePublicConfigRequest, opts ...grpc.CallOption) (*PublicConfigInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublicConfigInfo)
	err := c.cc.Invoke(ctx, PublicConfig_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicConfigClient) Delete(ctx context.Context, in *PublicConfigRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, PublicConfig_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicConfigClient) Get(ctx context.Context, in *PublicConfigSearchQuery, opts ...grpc.CallOption) (*PublicConfigInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublicConfigInfo)
	err := c.cc.Invoke(ctx, PublicConfig_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicConfigClient) GetAccessibleConfigs(ctx context.Context, in *PublicConfigSearchQuery, opts ...grpc.CallOption) (*PublicConfigsInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublicConfigsInfo)
	err := c.cc.Invoke(ctx, PublicConfig_GetAccessibleConfigs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicConfigClient) List(ctx context.Context, in *PublicConfigSearchQuery, opts ...grpc.CallOption) (*PublicConfigsInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublicConfigsInfo)
	err := c.cc.Invoke(ctx, PublicConfig_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicConfigClient) Stat(ctx context.Context, in *PublicConfigStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, PublicConfig_Stat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublicConfigServer is the server API for PublicConfig service.
// All implementations must embed UnimplementedPublicConfigServer
// for forward compatibility.
type PublicConfigServer interface {
	Create(context.Context, *CreatePublicConfigRequest) (*PublicConfigInfo, error)
	Update(context.Context, *UpdatePublicConfigRequest) (*PublicConfigInfo, error)
	Delete(context.Context, *PublicConfigRequest) (*empty.Empty, error)
	Get(context.Context, *PublicConfigSearchQuery) (*PublicConfigInfo, error)
	// This API for retrieving domain scoped configs that are accessible to users.
	GetAccessibleConfigs(context.Context, *PublicConfigSearchQuery) (*PublicConfigsInfo, error)
	List(context.Context, *PublicConfigSearchQuery) (*PublicConfigsInfo, error)
	Stat(context.Context, *PublicConfigStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedPublicConfigServer()
}

// UnimplementedPublicConfigServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPublicConfigServer struct{}

func (UnimplementedPublicConfigServer) Create(context.Context, *CreatePublicConfigRequest) (*PublicConfigInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPublicConfigServer) Update(context.Context, *UpdatePublicConfigRequest) (*PublicConfigInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPublicConfigServer) Delete(context.Context, *PublicConfigRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPublicConfigServer) Get(context.Context, *PublicConfigSearchQuery) (*PublicConfigInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPublicConfigServer) GetAccessibleConfigs(context.Context, *PublicConfigSearchQuery) (*PublicConfigsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccessibleConfigs not implemented")
}
func (UnimplementedPublicConfigServer) List(context.Context, *PublicConfigSearchQuery) (*PublicConfigsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPublicConfigServer) Stat(context.Context, *PublicConfigStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedPublicConfigServer) mustEmbedUnimplementedPublicConfigServer() {}
func (UnimplementedPublicConfigServer) testEmbeddedByValue()                      {}

// UnsafePublicConfigServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PublicConfigServer will
// result in compilation errors.
type UnsafePublicConfigServer interface {
	mustEmbedUnimplementedPublicConfigServer()
}

func RegisterPublicConfigServer(s grpc.ServiceRegistrar, srv PublicConfigServer) {
	// If the following call pancis, it indicates UnimplementedPublicConfigServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PublicConfig_ServiceDesc, srv)
}

func _PublicConfig_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePublicConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicConfigServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PublicConfig_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicConfigServer).Create(ctx, req.(*CreatePublicConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicConfig_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePublicConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicConfigServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PublicConfig_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicConfigServer).Update(ctx, req.(*UpdatePublicConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicConfig_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublicConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicConfigServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PublicConfig_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicConfigServer).Delete(ctx, req.(*PublicConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicConfig_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublicConfigSearchQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicConfigServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PublicConfig_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicConfigServer).Get(ctx, req.(*PublicConfigSearchQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicConfig_GetAccessibleConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublicConfigSearchQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicConfigServer).GetAccessibleConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PublicConfig_GetAccessibleConfigs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicConfigServer).GetAccessibleConfigs(ctx, req.(*PublicConfigSearchQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicConfig_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublicConfigSearchQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicConfigServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PublicConfig_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicConfigServer).List(ctx, req.(*PublicConfigSearchQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicConfig_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublicConfigStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicConfigServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PublicConfig_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicConfigServer).Stat(ctx, req.(*PublicConfigStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// PublicConfig_ServiceDesc is the grpc.ServiceDesc for PublicConfig service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PublicConfig_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.config.v1.PublicConfig",
	HandlerType: (*PublicConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _PublicConfig_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _PublicConfig_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _PublicConfig_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _PublicConfig_Get_Handler,
		},
		{
			MethodName: "get_accessible_configs",
			Handler:    _PublicConfig_GetAccessibleConfigs_Handler,
		},
		{
			MethodName: "list",
			Handler:    _PublicConfig_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _PublicConfig_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/config/v1/public_config.proto",
}
