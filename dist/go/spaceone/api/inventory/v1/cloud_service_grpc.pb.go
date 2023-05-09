// A CloudService is data of an `instance` of a `resource`. A CloudService follows the pre-created classification system of a CloudServiceType and indicates the property value of the `resource`.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: spaceone/api/inventory/v1/cloud_service.proto

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
	CloudService_Create_FullMethodName  = "/spaceone.api.inventory.v1.CloudService/create"
	CloudService_Update_FullMethodName  = "/spaceone.api.inventory.v1.CloudService/update"
	CloudService_Delete_FullMethodName  = "/spaceone.api.inventory.v1.CloudService/delete"
	CloudService_Get_FullMethodName     = "/spaceone.api.inventory.v1.CloudService/get"
	CloudService_List_FullMethodName    = "/spaceone.api.inventory.v1.CloudService/list"
	CloudService_Analyze_FullMethodName = "/spaceone.api.inventory.v1.CloudService/analyze"
	CloudService_Stat_FullMethodName    = "/spaceone.api.inventory.v1.CloudService/stat"
)

// CloudServiceClient is the client API for CloudService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CloudServiceClient interface {
	// Creates a new CloudService. A CloudService instance is created based on data including the `resource`'s attribute values. When creating, the classification information defined by CloudServiceType is also needed. The created CloudService has collection information which is the collection history for the `resource` by plugin.
	Create(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*CloudServiceInfo, error)
	// Updates a specific CloudService. You can make changes in CloudService settings, except for the classification system of CloudService and the information about the `resource` attribute value.
	Update(ctx context.Context, in *UpdateCloudServiceRequest, opts ...grpc.CallOption) (*CloudServiceInfo, error)
	// Deletes a specific CloudService. You must specify the `cloud_service_id` of the CloudService to delete.
	Delete(ctx context.Context, in *CloudServiceRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Gets a specific CloudService. Prints detailed information about the CloudService, including its state, classification information, and attribute values.
	Get(ctx context.Context, in *GetCloudServiceRequest, opts ...grpc.CallOption) (*CloudServiceInfo, error)
	// Gets a list of all CloudServices. You can use a query to get a filtered list of CloudServices.
	List(ctx context.Context, in *CloudServiceQuery, opts ...grpc.CallOption) (*CloudServicesInfo, error)
	Analyze(ctx context.Context, in *CloudServiceAnalyzeQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
	Stat(ctx context.Context, in *CloudServiceStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type cloudServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCloudServiceClient(cc grpc.ClientConnInterface) CloudServiceClient {
	return &cloudServiceClient{cc}
}

func (c *cloudServiceClient) Create(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*CloudServiceInfo, error) {
	out := new(CloudServiceInfo)
	err := c.cc.Invoke(ctx, CloudService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudServiceClient) Update(ctx context.Context, in *UpdateCloudServiceRequest, opts ...grpc.CallOption) (*CloudServiceInfo, error) {
	out := new(CloudServiceInfo)
	err := c.cc.Invoke(ctx, CloudService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudServiceClient) Delete(ctx context.Context, in *CloudServiceRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, CloudService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudServiceClient) Get(ctx context.Context, in *GetCloudServiceRequest, opts ...grpc.CallOption) (*CloudServiceInfo, error) {
	out := new(CloudServiceInfo)
	err := c.cc.Invoke(ctx, CloudService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudServiceClient) List(ctx context.Context, in *CloudServiceQuery, opts ...grpc.CallOption) (*CloudServicesInfo, error) {
	out := new(CloudServicesInfo)
	err := c.cc.Invoke(ctx, CloudService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudServiceClient) Analyze(ctx context.Context, in *CloudServiceAnalyzeQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, CloudService_Analyze_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudServiceClient) Stat(ctx context.Context, in *CloudServiceStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, CloudService_Stat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudServiceServer is the server API for CloudService service.
// All implementations must embed UnimplementedCloudServiceServer
// for forward compatibility
type CloudServiceServer interface {
	// Creates a new CloudService. A CloudService instance is created based on data including the `resource`'s attribute values. When creating, the classification information defined by CloudServiceType is also needed. The created CloudService has collection information which is the collection history for the `resource` by plugin.
	Create(context.Context, *CreateServiceRequest) (*CloudServiceInfo, error)
	// Updates a specific CloudService. You can make changes in CloudService settings, except for the classification system of CloudService and the information about the `resource` attribute value.
	Update(context.Context, *UpdateCloudServiceRequest) (*CloudServiceInfo, error)
	// Deletes a specific CloudService. You must specify the `cloud_service_id` of the CloudService to delete.
	Delete(context.Context, *CloudServiceRequest) (*empty.Empty, error)
	// Gets a specific CloudService. Prints detailed information about the CloudService, including its state, classification information, and attribute values.
	Get(context.Context, *GetCloudServiceRequest) (*CloudServiceInfo, error)
	// Gets a list of all CloudServices. You can use a query to get a filtered list of CloudServices.
	List(context.Context, *CloudServiceQuery) (*CloudServicesInfo, error)
	Analyze(context.Context, *CloudServiceAnalyzeQuery) (*_struct.Struct, error)
	Stat(context.Context, *CloudServiceStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedCloudServiceServer()
}

// UnimplementedCloudServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCloudServiceServer struct {
}

func (UnimplementedCloudServiceServer) Create(context.Context, *CreateServiceRequest) (*CloudServiceInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCloudServiceServer) Update(context.Context, *UpdateCloudServiceRequest) (*CloudServiceInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCloudServiceServer) Delete(context.Context, *CloudServiceRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCloudServiceServer) Get(context.Context, *GetCloudServiceRequest) (*CloudServiceInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCloudServiceServer) List(context.Context, *CloudServiceQuery) (*CloudServicesInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCloudServiceServer) Analyze(context.Context, *CloudServiceAnalyzeQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Analyze not implemented")
}
func (UnimplementedCloudServiceServer) Stat(context.Context, *CloudServiceStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedCloudServiceServer) mustEmbedUnimplementedCloudServiceServer() {}

// UnsafeCloudServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CloudServiceServer will
// result in compilation errors.
type UnsafeCloudServiceServer interface {
	mustEmbedUnimplementedCloudServiceServer()
}

func RegisterCloudServiceServer(s grpc.ServiceRegistrar, srv CloudServiceServer) {
	s.RegisterService(&CloudService_ServiceDesc, srv)
}

func _CloudService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudServiceServer).Create(ctx, req.(*CreateServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCloudServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudServiceServer).Update(ctx, req.(*UpdateCloudServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudServiceServer).Delete(ctx, req.(*CloudServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCloudServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudServiceServer).Get(ctx, req.(*GetCloudServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudServiceQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudServiceServer).List(ctx, req.(*CloudServiceQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudService_Analyze_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudServiceAnalyzeQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudServiceServer).Analyze(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudService_Analyze_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudServiceServer).Analyze(ctx, req.(*CloudServiceAnalyzeQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudService_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudServiceStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudServiceServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudService_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudServiceServer).Stat(ctx, req.(*CloudServiceStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// CloudService_ServiceDesc is the grpc.ServiceDesc for CloudService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CloudService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.inventory.v1.CloudService",
	HandlerType: (*CloudServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _CloudService_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _CloudService_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _CloudService_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _CloudService_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _CloudService_List_Handler,
		},
		{
			MethodName: "analyze",
			Handler:    _CloudService_Analyze_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _CloudService_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/inventory/v1/cloud_service.proto",
}
