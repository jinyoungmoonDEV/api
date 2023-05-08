//
//desc: A CloudServiceType is a classification with hierarchical information of `CloudService`. A CloudServiceType provides information about which `group` a specific `Resource` belongs to and which `Services` are in it.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: spaceone/api/inventory/v1/cloud_service_type.proto

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
	CloudServiceType_Create_FullMethodName = "/spaceone.api.inventory.v1.CloudServiceType/create"
	CloudServiceType_Update_FullMethodName = "/spaceone.api.inventory.v1.CloudServiceType/update"
	CloudServiceType_Delete_FullMethodName = "/spaceone.api.inventory.v1.CloudServiceType/delete"
	CloudServiceType_Get_FullMethodName    = "/spaceone.api.inventory.v1.CloudServiceType/get"
	CloudServiceType_List_FullMethodName   = "/spaceone.api.inventory.v1.CloudServiceType/list"
	CloudServiceType_Stat_FullMethodName   = "/spaceone.api.inventory.v1.CloudServiceType/stat"
)

// CloudServiceTypeClient is the client API for CloudServiceType service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CloudServiceTypeClient interface {
	// desc: Creates a new CloudServiceType. You must specify the `name`, `provider`, and `group` parameters to create a CloudServiceType. One or several CloudServiceTypes exist in a specific `group`, and each CloudServiceType is identified by the `name` parameter.
	// request_example: >-
	// {
	// "name": "API-TEST",
	// "provider": "aws",
	// "group": "APIGateway",
	// "service_code": "AmazonApiGateway",
	// "is_primary": true,
	// "is_major": true,
	// "resource_type": "inventory.CloudService",
	// "metadata": {},
	// "labels": [
	// "Networking"
	// ],
	// "tags": {
	// "a": "b"
	// }
	// }
	// response_example: >-
	// {
	// "cloud_service_type_id": "cloud-svc-type-27dd73ac89f8",
	// "name": "API-TEST",
	// "provider": "aws",
	// "group": "APIGateway",
	// "cloud_service_type_key": "aws.APIGateway.API-TEST",
	// "service_code": "AmazonApiGateway",
	// "is_primary": true,
	// "is_major": true,
	// "resource_type": "inventory.CloudService",
	// "metadata": {},
	// "tags": {
	// "a": "b"
	// },
	// "labels": [
	// "Networking"
	// ],
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2022-06-22T01:38:16.301Z",
	// "updated_at": "2022-06-22T01:38:16.301Z"
	// }
	Create(ctx context.Context, in *CreateCloudServiceTypeRequest, opts ...grpc.CallOption) (*CloudServiceTypeInfo, error)
	// desc: Updates a specific CloudServiceType. You can make changes in CloudServiceType settings, except for `name`, `provider` and `group`. In particular, you can set the CloudServiceType's priority in a `group`.
	// request_example: >-
	// {
	// "cloud_service_type_id": "cloud-svc-type-27dd73ac89f8",
	// "service_code": "AmazonApi",
	// "metadata": {},
	// "labels": [
	// "Networking2"
	// ],
	// "tags": {
	// "b": "c"
	// }
	// }
	// response_example: >-
	// {
	// "cloud_service_type_id": "cloud-svc-type-27dd73ac89f8",
	// "name": "API-TEST",
	// "provider": "aws",
	// "group": "APIGateway",
	// "cloud_service_type_key": "aws.APIGateway.API-TEST",
	// "service_code": "AmazonApi",
	// "is_primary": true,
	// "is_major": true,
	// "resource_type": "inventory.CloudService",
	// "metadata": {},
	// "tags": {
	// "b": "c"
	// },
	// "labels": [
	// "Networking2"
	// ],
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2022-06-22T01:38:16.301Z",
	// "updated_at": "2022-06-22T02:12:11.184Z"
	// }
	Update(ctx context.Context, in *UpdateCloudServiceTypeRequest, opts ...grpc.CallOption) (*CloudServiceTypeInfo, error)
	// desc: Deletes a specific CloudServiceType. You must specify the `cloud_service_type_id` of the CloudServiceType to delete.
	// request_example: >-
	// {
	// "cloud_service_type_id": "cloud-svc-type-27dd73ac89f8"
	// }
	Delete(ctx context.Context, in *CloudServiceTypeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// desc: Gets a specific CloudServiceType. Prints detailed information about the CloudServiceType.
	// request_example: >-
	// {
	// "cloud_service_type_id": "cloud-svc-type-27dd73ac89f8"
	// }
	// response_example: >-
	// {
	// "cloud_service_type_id": "cloud-svc-type-27dd73ac89f8",
	// "name": "API-TEST",
	// "provider": "aws",
	// "group": "APIGateway",
	// "cloud_service_type_key": "aws.APIGateway.API-TEST",
	// "service_code": "AmazonApi",
	// "is_primary": true,
	// "is_major": true,
	// "resource_type": "inventory.CloudService",
	// "metadata": {},
	// "tags": {
	// "b": "c"
	// },
	// "labels": [
	// "Networking2"
	// ],
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2022-06-22T01:38:16.301Z",
	// "updated_at": "2022-06-22T02:12:11.184Z"
	// }
	Get(ctx context.Context, in *GetCloudServiceTypeRequest, opts ...grpc.CallOption) (*CloudServiceTypeInfo, error)
	// desc: Gets a list of all CloudServiceTypes. You can use a query to get a filtered list of CloudServiceTypes.
	// request_example: >-
	// {
	// "query": {
	// "filter": [
	// {
	// "key": "provider",
	// "value": "aws",
	// "operator": "eq"
	// }
	// ]
	// }
	// }
	// response_example: >-
	// {
	// "results": [
	// {
	// "cloud_service_type_id": "cloud-svc-type-7e1c113b39ff",
	// "name": "API",
	// "provider": "aws",
	// "group": "APIGateway",
	// "cloud_service_type_key": "aws.APIGateway.API",
	// "service_code": "AmazonApiGateway",
	// "is_primary": true,
	// "is_major": true,
	// "resource_type": "inventory.CloudService",
	// "metadata": {
	// },
	// "tags": {
	// "spaceone:icon": "https://spaceone.s3.ap-northeast-2.amazonaws.com/console-assets/icons/cloud-services/aws/Amazon-API-Gateway.svg"
	// },
	// "labels": [
	// "Networking"
	// ],
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2021-06-03T02:29:32.690Z",
	// "updated_at": "2022-06-22T00:04:45.477Z"
	// },
	// {
	// "cloud_service_type_id": "cloud-svc-type-64a0de601371",
	// "name": "Certificate",
	// "provider": "aws",
	// "group": "CertificateManager",
	// "cloud_service_type_key": "aws.CertificateManager.Certificate",
	// "service_code": "AWSCertificateManager",
	// "is_primary": true,
	// "resource_type": "inventory.CloudService",
	// "metadata": {
	// },
	// "tags": {
	// "spaceone:icon": "https://spaceone.s3.ap-northeast-2.amazonaws.com/console-assets/icons/cloud-services/aws/AWS-Certificate-Manager.svg"
	// },
	// "labels": [
	// "Security"
	// ],
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2021-06-03T02:29:53.052Z",
	// "updated_at": "2022-06-22T00:05:41.252Z"
	// }
	// ],
	// "total_count": 2
	// }
	List(ctx context.Context, in *CloudServiceTypeQuery, opts ...grpc.CallOption) (*CloudServiceTypesInfo, error)
	Stat(ctx context.Context, in *CloudServiceTypeStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type cloudServiceTypeClient struct {
	cc grpc.ClientConnInterface
}

func NewCloudServiceTypeClient(cc grpc.ClientConnInterface) CloudServiceTypeClient {
	return &cloudServiceTypeClient{cc}
}

func (c *cloudServiceTypeClient) Create(ctx context.Context, in *CreateCloudServiceTypeRequest, opts ...grpc.CallOption) (*CloudServiceTypeInfo, error) {
	out := new(CloudServiceTypeInfo)
	err := c.cc.Invoke(ctx, CloudServiceType_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudServiceTypeClient) Update(ctx context.Context, in *UpdateCloudServiceTypeRequest, opts ...grpc.CallOption) (*CloudServiceTypeInfo, error) {
	out := new(CloudServiceTypeInfo)
	err := c.cc.Invoke(ctx, CloudServiceType_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudServiceTypeClient) Delete(ctx context.Context, in *CloudServiceTypeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, CloudServiceType_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudServiceTypeClient) Get(ctx context.Context, in *GetCloudServiceTypeRequest, opts ...grpc.CallOption) (*CloudServiceTypeInfo, error) {
	out := new(CloudServiceTypeInfo)
	err := c.cc.Invoke(ctx, CloudServiceType_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudServiceTypeClient) List(ctx context.Context, in *CloudServiceTypeQuery, opts ...grpc.CallOption) (*CloudServiceTypesInfo, error) {
	out := new(CloudServiceTypesInfo)
	err := c.cc.Invoke(ctx, CloudServiceType_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudServiceTypeClient) Stat(ctx context.Context, in *CloudServiceTypeStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, CloudServiceType_Stat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudServiceTypeServer is the server API for CloudServiceType service.
// All implementations must embed UnimplementedCloudServiceTypeServer
// for forward compatibility
type CloudServiceTypeServer interface {
	// desc: Creates a new CloudServiceType. You must specify the `name`, `provider`, and `group` parameters to create a CloudServiceType. One or several CloudServiceTypes exist in a specific `group`, and each CloudServiceType is identified by the `name` parameter.
	// request_example: >-
	// {
	// "name": "API-TEST",
	// "provider": "aws",
	// "group": "APIGateway",
	// "service_code": "AmazonApiGateway",
	// "is_primary": true,
	// "is_major": true,
	// "resource_type": "inventory.CloudService",
	// "metadata": {},
	// "labels": [
	// "Networking"
	// ],
	// "tags": {
	// "a": "b"
	// }
	// }
	// response_example: >-
	// {
	// "cloud_service_type_id": "cloud-svc-type-27dd73ac89f8",
	// "name": "API-TEST",
	// "provider": "aws",
	// "group": "APIGateway",
	// "cloud_service_type_key": "aws.APIGateway.API-TEST",
	// "service_code": "AmazonApiGateway",
	// "is_primary": true,
	// "is_major": true,
	// "resource_type": "inventory.CloudService",
	// "metadata": {},
	// "tags": {
	// "a": "b"
	// },
	// "labels": [
	// "Networking"
	// ],
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2022-06-22T01:38:16.301Z",
	// "updated_at": "2022-06-22T01:38:16.301Z"
	// }
	Create(context.Context, *CreateCloudServiceTypeRequest) (*CloudServiceTypeInfo, error)
	// desc: Updates a specific CloudServiceType. You can make changes in CloudServiceType settings, except for `name`, `provider` and `group`. In particular, you can set the CloudServiceType's priority in a `group`.
	// request_example: >-
	// {
	// "cloud_service_type_id": "cloud-svc-type-27dd73ac89f8",
	// "service_code": "AmazonApi",
	// "metadata": {},
	// "labels": [
	// "Networking2"
	// ],
	// "tags": {
	// "b": "c"
	// }
	// }
	// response_example: >-
	// {
	// "cloud_service_type_id": "cloud-svc-type-27dd73ac89f8",
	// "name": "API-TEST",
	// "provider": "aws",
	// "group": "APIGateway",
	// "cloud_service_type_key": "aws.APIGateway.API-TEST",
	// "service_code": "AmazonApi",
	// "is_primary": true,
	// "is_major": true,
	// "resource_type": "inventory.CloudService",
	// "metadata": {},
	// "tags": {
	// "b": "c"
	// },
	// "labels": [
	// "Networking2"
	// ],
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2022-06-22T01:38:16.301Z",
	// "updated_at": "2022-06-22T02:12:11.184Z"
	// }
	Update(context.Context, *UpdateCloudServiceTypeRequest) (*CloudServiceTypeInfo, error)
	// desc: Deletes a specific CloudServiceType. You must specify the `cloud_service_type_id` of the CloudServiceType to delete.
	// request_example: >-
	// {
	// "cloud_service_type_id": "cloud-svc-type-27dd73ac89f8"
	// }
	Delete(context.Context, *CloudServiceTypeRequest) (*empty.Empty, error)
	// desc: Gets a specific CloudServiceType. Prints detailed information about the CloudServiceType.
	// request_example: >-
	// {
	// "cloud_service_type_id": "cloud-svc-type-27dd73ac89f8"
	// }
	// response_example: >-
	// {
	// "cloud_service_type_id": "cloud-svc-type-27dd73ac89f8",
	// "name": "API-TEST",
	// "provider": "aws",
	// "group": "APIGateway",
	// "cloud_service_type_key": "aws.APIGateway.API-TEST",
	// "service_code": "AmazonApi",
	// "is_primary": true,
	// "is_major": true,
	// "resource_type": "inventory.CloudService",
	// "metadata": {},
	// "tags": {
	// "b": "c"
	// },
	// "labels": [
	// "Networking2"
	// ],
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2022-06-22T01:38:16.301Z",
	// "updated_at": "2022-06-22T02:12:11.184Z"
	// }
	Get(context.Context, *GetCloudServiceTypeRequest) (*CloudServiceTypeInfo, error)
	// desc: Gets a list of all CloudServiceTypes. You can use a query to get a filtered list of CloudServiceTypes.
	// request_example: >-
	// {
	// "query": {
	// "filter": [
	// {
	// "key": "provider",
	// "value": "aws",
	// "operator": "eq"
	// }
	// ]
	// }
	// }
	// response_example: >-
	// {
	// "results": [
	// {
	// "cloud_service_type_id": "cloud-svc-type-7e1c113b39ff",
	// "name": "API",
	// "provider": "aws",
	// "group": "APIGateway",
	// "cloud_service_type_key": "aws.APIGateway.API",
	// "service_code": "AmazonApiGateway",
	// "is_primary": true,
	// "is_major": true,
	// "resource_type": "inventory.CloudService",
	// "metadata": {
	// },
	// "tags": {
	// "spaceone:icon": "https://spaceone.s3.ap-northeast-2.amazonaws.com/console-assets/icons/cloud-services/aws/Amazon-API-Gateway.svg"
	// },
	// "labels": [
	// "Networking"
	// ],
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2021-06-03T02:29:32.690Z",
	// "updated_at": "2022-06-22T00:04:45.477Z"
	// },
	// {
	// "cloud_service_type_id": "cloud-svc-type-64a0de601371",
	// "name": "Certificate",
	// "provider": "aws",
	// "group": "CertificateManager",
	// "cloud_service_type_key": "aws.CertificateManager.Certificate",
	// "service_code": "AWSCertificateManager",
	// "is_primary": true,
	// "resource_type": "inventory.CloudService",
	// "metadata": {
	// },
	// "tags": {
	// "spaceone:icon": "https://spaceone.s3.ap-northeast-2.amazonaws.com/console-assets/icons/cloud-services/aws/AWS-Certificate-Manager.svg"
	// },
	// "labels": [
	// "Security"
	// ],
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2021-06-03T02:29:53.052Z",
	// "updated_at": "2022-06-22T00:05:41.252Z"
	// }
	// ],
	// "total_count": 2
	// }
	List(context.Context, *CloudServiceTypeQuery) (*CloudServiceTypesInfo, error)
	Stat(context.Context, *CloudServiceTypeStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedCloudServiceTypeServer()
}

// UnimplementedCloudServiceTypeServer must be embedded to have forward compatible implementations.
type UnimplementedCloudServiceTypeServer struct {
}

func (UnimplementedCloudServiceTypeServer) Create(context.Context, *CreateCloudServiceTypeRequest) (*CloudServiceTypeInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCloudServiceTypeServer) Update(context.Context, *UpdateCloudServiceTypeRequest) (*CloudServiceTypeInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCloudServiceTypeServer) Delete(context.Context, *CloudServiceTypeRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCloudServiceTypeServer) Get(context.Context, *GetCloudServiceTypeRequest) (*CloudServiceTypeInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCloudServiceTypeServer) List(context.Context, *CloudServiceTypeQuery) (*CloudServiceTypesInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCloudServiceTypeServer) Stat(context.Context, *CloudServiceTypeStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedCloudServiceTypeServer) mustEmbedUnimplementedCloudServiceTypeServer() {}

// UnsafeCloudServiceTypeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CloudServiceTypeServer will
// result in compilation errors.
type UnsafeCloudServiceTypeServer interface {
	mustEmbedUnimplementedCloudServiceTypeServer()
}

func RegisterCloudServiceTypeServer(s grpc.ServiceRegistrar, srv CloudServiceTypeServer) {
	s.RegisterService(&CloudServiceType_ServiceDesc, srv)
}

func _CloudServiceType_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCloudServiceTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudServiceTypeServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudServiceType_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudServiceTypeServer).Create(ctx, req.(*CreateCloudServiceTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudServiceType_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCloudServiceTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudServiceTypeServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudServiceType_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudServiceTypeServer).Update(ctx, req.(*UpdateCloudServiceTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudServiceType_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudServiceTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudServiceTypeServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudServiceType_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudServiceTypeServer).Delete(ctx, req.(*CloudServiceTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudServiceType_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCloudServiceTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudServiceTypeServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudServiceType_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudServiceTypeServer).Get(ctx, req.(*GetCloudServiceTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudServiceType_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudServiceTypeQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudServiceTypeServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudServiceType_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudServiceTypeServer).List(ctx, req.(*CloudServiceTypeQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudServiceType_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudServiceTypeStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudServiceTypeServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudServiceType_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudServiceTypeServer).Stat(ctx, req.(*CloudServiceTypeStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// CloudServiceType_ServiceDesc is the grpc.ServiceDesc for CloudServiceType service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CloudServiceType_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.inventory.v1.CloudServiceType",
	HandlerType: (*CloudServiceTypeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _CloudServiceType_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _CloudServiceType_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _CloudServiceType_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _CloudServiceType_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _CloudServiceType_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _CloudServiceType_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/inventory/v1/cloud_service_type.proto",
}
