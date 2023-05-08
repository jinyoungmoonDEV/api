// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: spaceone/api/statistics/v1/storage.proto

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
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Storage_Register_FullMethodName     = "/spaceone.api.statistics.v1.Storage/register"
	Storage_Update_FullMethodName       = "/spaceone.api.statistics.v1.Storage/update"
	Storage_UpdatePlugin_FullMethodName = "/spaceone.api.statistics.v1.Storage/update_plugin"
	Storage_VerifyPlugin_FullMethodName = "/spaceone.api.statistics.v1.Storage/verify_plugin"
	Storage_Enable_FullMethodName       = "/spaceone.api.statistics.v1.Storage/enable"
	Storage_Disable_FullMethodName      = "/spaceone.api.statistics.v1.Storage/disable"
	Storage_Deregister_FullMethodName   = "/spaceone.api.statistics.v1.Storage/deregister"
	Storage_Get_FullMethodName          = "/spaceone.api.statistics.v1.Storage/get"
	Storage_List_FullMethodName         = "/spaceone.api.statistics.v1.Storage/list"
	Storage_Stat_FullMethodName         = "/spaceone.api.statistics.v1.Storage/stat"
)

// StorageClient is the client API for Storage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageClient interface {
	Register(ctx context.Context, in *RegisterStorageRequest, opts ...grpc.CallOption) (*StorageInfo, error)
	Update(ctx context.Context, in *UpdateStorageRequest, opts ...grpc.CallOption) (*StorageInfo, error)
	UpdatePlugin(ctx context.Context, in *UpdateStoragePluginRequest, opts ...grpc.CallOption) (*StorageInfo, error)
	VerifyPlugin(ctx context.Context, in *StorageRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Enable(ctx context.Context, in *StorageRequest, opts ...grpc.CallOption) (*StorageInfo, error)
	Disable(ctx context.Context, in *StorageRequest, opts ...grpc.CallOption) (*StorageInfo, error)
	Deregister(ctx context.Context, in *StorageRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Get(ctx context.Context, in *GetStorageRequest, opts ...grpc.CallOption) (*StorageInfo, error)
	List(ctx context.Context, in *StorageQuery, opts ...grpc.CallOption) (*StoragesInfo, error)
	Stat(ctx context.Context, in *StorageStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type storageClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageClient(cc grpc.ClientConnInterface) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) Register(ctx context.Context, in *RegisterStorageRequest, opts ...grpc.CallOption) (*StorageInfo, error) {
	out := new(StorageInfo)
	err := c.cc.Invoke(ctx, Storage_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Update(ctx context.Context, in *UpdateStorageRequest, opts ...grpc.CallOption) (*StorageInfo, error) {
	out := new(StorageInfo)
	err := c.cc.Invoke(ctx, Storage_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) UpdatePlugin(ctx context.Context, in *UpdateStoragePluginRequest, opts ...grpc.CallOption) (*StorageInfo, error) {
	out := new(StorageInfo)
	err := c.cc.Invoke(ctx, Storage_UpdatePlugin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) VerifyPlugin(ctx context.Context, in *StorageRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Storage_VerifyPlugin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Enable(ctx context.Context, in *StorageRequest, opts ...grpc.CallOption) (*StorageInfo, error) {
	out := new(StorageInfo)
	err := c.cc.Invoke(ctx, Storage_Enable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Disable(ctx context.Context, in *StorageRequest, opts ...grpc.CallOption) (*StorageInfo, error) {
	out := new(StorageInfo)
	err := c.cc.Invoke(ctx, Storage_Disable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Deregister(ctx context.Context, in *StorageRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Storage_Deregister_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Get(ctx context.Context, in *GetStorageRequest, opts ...grpc.CallOption) (*StorageInfo, error) {
	out := new(StorageInfo)
	err := c.cc.Invoke(ctx, Storage_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) List(ctx context.Context, in *StorageQuery, opts ...grpc.CallOption) (*StoragesInfo, error) {
	out := new(StoragesInfo)
	err := c.cc.Invoke(ctx, Storage_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Stat(ctx context.Context, in *StorageStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, Storage_Stat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServer is the server API for Storage service.
// All implementations must embed UnimplementedStorageServer
// for forward compatibility
type StorageServer interface {
	Register(context.Context, *RegisterStorageRequest) (*StorageInfo, error)
	Update(context.Context, *UpdateStorageRequest) (*StorageInfo, error)
	UpdatePlugin(context.Context, *UpdateStoragePluginRequest) (*StorageInfo, error)
	VerifyPlugin(context.Context, *StorageRequest) (*empty.Empty, error)
	Enable(context.Context, *StorageRequest) (*StorageInfo, error)
	Disable(context.Context, *StorageRequest) (*StorageInfo, error)
	Deregister(context.Context, *StorageRequest) (*empty.Empty, error)
	Get(context.Context, *GetStorageRequest) (*StorageInfo, error)
	List(context.Context, *StorageQuery) (*StoragesInfo, error)
	Stat(context.Context, *StorageStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedStorageServer()
}

// UnimplementedStorageServer must be embedded to have forward compatible implementations.
type UnimplementedStorageServer struct {
}

func (UnimplementedStorageServer) Register(context.Context, *RegisterStorageRequest) (*StorageInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedStorageServer) Update(context.Context, *UpdateStorageRequest) (*StorageInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedStorageServer) UpdatePlugin(context.Context, *UpdateStoragePluginRequest) (*StorageInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlugin not implemented")
}
func (UnimplementedStorageServer) VerifyPlugin(context.Context, *StorageRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyPlugin not implemented")
}
func (UnimplementedStorageServer) Enable(context.Context, *StorageRequest) (*StorageInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enable not implemented")
}
func (UnimplementedStorageServer) Disable(context.Context, *StorageRequest) (*StorageInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disable not implemented")
}
func (UnimplementedStorageServer) Deregister(context.Context, *StorageRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deregister not implemented")
}
func (UnimplementedStorageServer) Get(context.Context, *GetStorageRequest) (*StorageInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedStorageServer) List(context.Context, *StorageQuery) (*StoragesInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedStorageServer) Stat(context.Context, *StorageStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedStorageServer) mustEmbedUnimplementedStorageServer() {}

// UnsafeStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServer will
// result in compilation errors.
type UnsafeStorageServer interface {
	mustEmbedUnimplementedStorageServer()
}

func RegisterStorageServer(s grpc.ServiceRegistrar, srv StorageServer) {
	s.RegisterService(&Storage_ServiceDesc, srv)
}

func _Storage_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Register(ctx, req.(*RegisterStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Update(ctx, req.(*UpdateStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_UpdatePlugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStoragePluginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).UpdatePlugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_UpdatePlugin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).UpdatePlugin(ctx, req.(*UpdateStoragePluginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_VerifyPlugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).VerifyPlugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_VerifyPlugin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).VerifyPlugin(ctx, req.(*StorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Enable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Enable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_Enable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Enable(ctx, req.(*StorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Disable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Disable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_Disable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Disable(ctx, req.(*StorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Deregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Deregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_Deregister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Deregister(ctx, req.(*StorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Get(ctx, req.(*GetStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).List(ctx, req.(*StorageQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Storage_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Stat(ctx, req.(*StorageStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// Storage_ServiceDesc is the grpc.ServiceDesc for Storage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Storage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.statistics.v1.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "register",
			Handler:    _Storage_Register_Handler,
		},
		{
			MethodName: "update",
			Handler:    _Storage_Update_Handler,
		},
		{
			MethodName: "update_plugin",
			Handler:    _Storage_UpdatePlugin_Handler,
		},
		{
			MethodName: "verify_plugin",
			Handler:    _Storage_VerifyPlugin_Handler,
		},
		{
			MethodName: "enable",
			Handler:    _Storage_Enable_Handler,
		},
		{
			MethodName: "disable",
			Handler:    _Storage_Disable_Handler,
		},
		{
			MethodName: "deregister",
			Handler:    _Storage_Deregister_Handler,
		},
		{
			MethodName: "get",
			Handler:    _Storage_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _Storage_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _Storage_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/statistics/v1/storage.proto",
}
