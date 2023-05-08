//
//desc: A Secret is an external data, encrypted by CloudForet.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: spaceone/api/secret/v1/secret.proto

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
	Secret_Create_FullMethodName     = "/spaceone.api.secret.v1.Secret/create"
	Secret_Update_FullMethodName     = "/spaceone.api.secret.v1.Secret/update"
	Secret_Delete_FullMethodName     = "/spaceone.api.secret.v1.Secret/delete"
	Secret_UpdateData_FullMethodName = "/spaceone.api.secret.v1.Secret/update_data"
	Secret_GetData_FullMethodName    = "/spaceone.api.secret.v1.Secret/get_data"
	Secret_Get_FullMethodName        = "/spaceone.api.secret.v1.Secret/get"
	Secret_List_FullMethodName       = "/spaceone.api.secret.v1.Secret/list"
	Secret_Stat_FullMethodName       = "/spaceone.api.secret.v1.Secret/stat"
)

// SecretClient is the client API for Secret service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecretClient interface {
	// desc: Creates a new Secret. When creating the resource, external `data` is encrypted, and a `secret_id` is issued for data access by other microservices.
	// request_example: >-
	// {
	// "name": "aws-dev",
	// "data": "********",
	// "secret_type": "CREDENTIALS",
	// "schema": "aws_access_key",
	// "service_account_id": "sa-123456789012",
	// "project_id": "project-123456789012",
	// "domain_id": "domain-123456789012"
	// }
	// response_example: >-
	// {
	// "secret_id": "secret-123456789012",
	// "name": "aws-dev",
	// "secret_type": "CREDENTIALS",
	// "tags": {},
	// "schema": "aws_access_key",
	// "provider": "aws",
	// "service_account_id": "sa-123456789012",
	// "project_id": "project-123456789012",
	// "domain_id": "domain-123456789012",
	// "created_at": "2022-01-01T06:10:14.851Z"
	// }
	Create(ctx context.Context, in *CreateSecretRequest, opts ...grpc.CallOption) (*SecretInfo, error)
	// desc: Updates a specific Secret. You can make changes in Secret settings, including `name` and`tags`.
	// request_example: >-
	// {
	// "secret_id": "secret-123456789012",
	// "name": "aws-dev2",
	// "tags": { "a": "b"},
	// "project_id": "project-123456789012",
	// "release_project": true,
	// "domain_id": "domain-123456789012"
	// }
	// response_example: >-
	// {
	// "secret_id": "secret-123456789012",
	// "name": "aws-dev2",
	// "secret_type": "CREDENTIALS",
	// "tags": { "a": "b"},
	// "schema": "aws_access_key",
	// "provider": "aws",
	// "service_account_id": "sa-123456789012",
	// "project_id": "",
	// "domain_id": "domain-123456789012",
	// "created_at": "2022-01-01T06:10:14.851Z"
	// }
	Update(ctx context.Context, in *UpdateSecretRequest, opts ...grpc.CallOption) (*SecretInfo, error)
	// desc: Deletes a specific Secret. You must specify the `secret_id` of the Secret to delete.
	// request_example: >-
	// {
	// "secret_id": "secret-123456789012",
	// "domain_id": "domain-123456789012"
	// }
	Delete(ctx context.Context, in *SecretRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// desc: Updates encrypted data of a specific Secret resource. For example, to change the parameter `data`, external data to encrypt, you can use `update_data` to create new encrypted data based on the changed `data` and store it in the Secret resource.
	// request_example: >-
	// {
	// "secret_id": "secret-123456789012",
	// "data": "********",
	// "domain_id": "domain-123456789012"
	// }
	// response_example: >-
	// {
	// "data": {
	// "encrypted_data": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	// },
	// "encrypted": true,
	// "encrypt_options": {
	// "encrypted_data_key": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
	// "encrypt_algorithm": "CloudForet_DEFAULT"
	// }
	// }
	UpdateData(ctx context.Context, in *UpdateSecretDataRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// desc: Gets a specific Secret. Prints detailed information about the Secret, including  `name`, `tags`, `schema`, and `provider`.
	// request_example: >-
	// {
	// "secret_id": "secret-123456789012",
	// "domain_id": "domain-123456789012"
	// }
	// response_example: >-
	// {
	// "data": {
	// "encrypted_data": "xxxxxxxxxxxxxxxxxx"
	// },
	// "encrypted": true,
	// "encrypt_options": {
	// "encrypt_algorithm": "SPACEONE_DEFAULT",
	// "encrypted_data_key": "xxxxxx"
	// }
	// }
	GetData(ctx context.Context, in *SecretRequest, opts ...grpc.CallOption) (*SecretDataInfo, error)
	// desc: Gets a specific Post. You must specify the `post_id` of the Post to get, and the `board_id` of the Board where the child Post to get belongs. Prints detailed information about the Post.
	// request_example: >-
	// {
	// "secret_id": "secret-123456789012",
	// "domain_id": "domain-123456789012"
	// }
	// response_example: >-
	// {
	// "secret_id": "secret-123456789012",
	// "name": "aws-dev",
	// "secret_type": "CREDENTIALS",
	// "tags": {},
	// "schema": "aws_access_key",
	// "provider": "aws",
	// "service_account_id": "sa-123456789012",
	// "project_id": "project-123456789012",
	// "domain_id": "domain-123456789012",
	// "created_at": "2022-01-01T06:10:14.851Z"
	// }
	Get(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*SecretInfo, error)
	// desc: Gets a list of all Posts. You can use a query to get a filtered list of Posts.
	// request_example: >-
	// {
	// "query": {},
	// "domain_id": "domain-123456789012"
	// }
	// response_example: >-
	// {
	// "results": [
	// {
	// "secret_id": "secret-123456789012",
	// "name": "aws-dev",
	// "secret_type": "CREDENTIALS",
	// "tags": {},
	// "schema": "aws_access_key",
	// "provider": "aws",
	// "service_account_id": "sa-123456789012",
	// "project_id": "project-123456789012",
	// "domain_id": "domain-123456789012",
	// "created_at": "2022-01-01T06:10:14.851Z"
	// },
	// {
	// "secret_id": "secret-987654321098",
	// "name": "plugin-credentials",
	// "secret_type": "CREDENTIALS",
	// "tags": {},
	// "domain_id": "domain-123456789012",
	// "created_at": "2022-01-01T02:31:01.709Z"
	// }
	// ],
	// "total_count": 2
	// }
	List(ctx context.Context, in *SecretQuery, opts ...grpc.CallOption) (*SecretsInfo, error)
	Stat(ctx context.Context, in *SecretStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type secretClient struct {
	cc grpc.ClientConnInterface
}

func NewSecretClient(cc grpc.ClientConnInterface) SecretClient {
	return &secretClient{cc}
}

func (c *secretClient) Create(ctx context.Context, in *CreateSecretRequest, opts ...grpc.CallOption) (*SecretInfo, error) {
	out := new(SecretInfo)
	err := c.cc.Invoke(ctx, Secret_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretClient) Update(ctx context.Context, in *UpdateSecretRequest, opts ...grpc.CallOption) (*SecretInfo, error) {
	out := new(SecretInfo)
	err := c.cc.Invoke(ctx, Secret_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretClient) Delete(ctx context.Context, in *SecretRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Secret_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretClient) UpdateData(ctx context.Context, in *UpdateSecretDataRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Secret_UpdateData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretClient) GetData(ctx context.Context, in *SecretRequest, opts ...grpc.CallOption) (*SecretDataInfo, error) {
	out := new(SecretDataInfo)
	err := c.cc.Invoke(ctx, Secret_GetData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretClient) Get(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*SecretInfo, error) {
	out := new(SecretInfo)
	err := c.cc.Invoke(ctx, Secret_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretClient) List(ctx context.Context, in *SecretQuery, opts ...grpc.CallOption) (*SecretsInfo, error) {
	out := new(SecretsInfo)
	err := c.cc.Invoke(ctx, Secret_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretClient) Stat(ctx context.Context, in *SecretStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, Secret_Stat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretServer is the server API for Secret service.
// All implementations must embed UnimplementedSecretServer
// for forward compatibility
type SecretServer interface {
	// desc: Creates a new Secret. When creating the resource, external `data` is encrypted, and a `secret_id` is issued for data access by other microservices.
	// request_example: >-
	// {
	// "name": "aws-dev",
	// "data": "********",
	// "secret_type": "CREDENTIALS",
	// "schema": "aws_access_key",
	// "service_account_id": "sa-123456789012",
	// "project_id": "project-123456789012",
	// "domain_id": "domain-123456789012"
	// }
	// response_example: >-
	// {
	// "secret_id": "secret-123456789012",
	// "name": "aws-dev",
	// "secret_type": "CREDENTIALS",
	// "tags": {},
	// "schema": "aws_access_key",
	// "provider": "aws",
	// "service_account_id": "sa-123456789012",
	// "project_id": "project-123456789012",
	// "domain_id": "domain-123456789012",
	// "created_at": "2022-01-01T06:10:14.851Z"
	// }
	Create(context.Context, *CreateSecretRequest) (*SecretInfo, error)
	// desc: Updates a specific Secret. You can make changes in Secret settings, including `name` and`tags`.
	// request_example: >-
	// {
	// "secret_id": "secret-123456789012",
	// "name": "aws-dev2",
	// "tags": { "a": "b"},
	// "project_id": "project-123456789012",
	// "release_project": true,
	// "domain_id": "domain-123456789012"
	// }
	// response_example: >-
	// {
	// "secret_id": "secret-123456789012",
	// "name": "aws-dev2",
	// "secret_type": "CREDENTIALS",
	// "tags": { "a": "b"},
	// "schema": "aws_access_key",
	// "provider": "aws",
	// "service_account_id": "sa-123456789012",
	// "project_id": "",
	// "domain_id": "domain-123456789012",
	// "created_at": "2022-01-01T06:10:14.851Z"
	// }
	Update(context.Context, *UpdateSecretRequest) (*SecretInfo, error)
	// desc: Deletes a specific Secret. You must specify the `secret_id` of the Secret to delete.
	// request_example: >-
	// {
	// "secret_id": "secret-123456789012",
	// "domain_id": "domain-123456789012"
	// }
	Delete(context.Context, *SecretRequest) (*empty.Empty, error)
	// desc: Updates encrypted data of a specific Secret resource. For example, to change the parameter `data`, external data to encrypt, you can use `update_data` to create new encrypted data based on the changed `data` and store it in the Secret resource.
	// request_example: >-
	// {
	// "secret_id": "secret-123456789012",
	// "data": "********",
	// "domain_id": "domain-123456789012"
	// }
	// response_example: >-
	// {
	// "data": {
	// "encrypted_data": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	// },
	// "encrypted": true,
	// "encrypt_options": {
	// "encrypted_data_key": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
	// "encrypt_algorithm": "CloudForet_DEFAULT"
	// }
	// }
	UpdateData(context.Context, *UpdateSecretDataRequest) (*empty.Empty, error)
	// desc: Gets a specific Secret. Prints detailed information about the Secret, including  `name`, `tags`, `schema`, and `provider`.
	// request_example: >-
	// {
	// "secret_id": "secret-123456789012",
	// "domain_id": "domain-123456789012"
	// }
	// response_example: >-
	// {
	// "data": {
	// "encrypted_data": "xxxxxxxxxxxxxxxxxx"
	// },
	// "encrypted": true,
	// "encrypt_options": {
	// "encrypt_algorithm": "SPACEONE_DEFAULT",
	// "encrypted_data_key": "xxxxxx"
	// }
	// }
	GetData(context.Context, *SecretRequest) (*SecretDataInfo, error)
	// desc: Gets a specific Post. You must specify the `post_id` of the Post to get, and the `board_id` of the Board where the child Post to get belongs. Prints detailed information about the Post.
	// request_example: >-
	// {
	// "secret_id": "secret-123456789012",
	// "domain_id": "domain-123456789012"
	// }
	// response_example: >-
	// {
	// "secret_id": "secret-123456789012",
	// "name": "aws-dev",
	// "secret_type": "CREDENTIALS",
	// "tags": {},
	// "schema": "aws_access_key",
	// "provider": "aws",
	// "service_account_id": "sa-123456789012",
	// "project_id": "project-123456789012",
	// "domain_id": "domain-123456789012",
	// "created_at": "2022-01-01T06:10:14.851Z"
	// }
	Get(context.Context, *GetSecretRequest) (*SecretInfo, error)
	// desc: Gets a list of all Posts. You can use a query to get a filtered list of Posts.
	// request_example: >-
	// {
	// "query": {},
	// "domain_id": "domain-123456789012"
	// }
	// response_example: >-
	// {
	// "results": [
	// {
	// "secret_id": "secret-123456789012",
	// "name": "aws-dev",
	// "secret_type": "CREDENTIALS",
	// "tags": {},
	// "schema": "aws_access_key",
	// "provider": "aws",
	// "service_account_id": "sa-123456789012",
	// "project_id": "project-123456789012",
	// "domain_id": "domain-123456789012",
	// "created_at": "2022-01-01T06:10:14.851Z"
	// },
	// {
	// "secret_id": "secret-987654321098",
	// "name": "plugin-credentials",
	// "secret_type": "CREDENTIALS",
	// "tags": {},
	// "domain_id": "domain-123456789012",
	// "created_at": "2022-01-01T02:31:01.709Z"
	// }
	// ],
	// "total_count": 2
	// }
	List(context.Context, *SecretQuery) (*SecretsInfo, error)
	Stat(context.Context, *SecretStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedSecretServer()
}

// UnimplementedSecretServer must be embedded to have forward compatible implementations.
type UnimplementedSecretServer struct {
}

func (UnimplementedSecretServer) Create(context.Context, *CreateSecretRequest) (*SecretInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSecretServer) Update(context.Context, *UpdateSecretRequest) (*SecretInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSecretServer) Delete(context.Context, *SecretRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSecretServer) UpdateData(context.Context, *UpdateSecretDataRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateData not implemented")
}
func (UnimplementedSecretServer) GetData(context.Context, *SecretRequest) (*SecretDataInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetData not implemented")
}
func (UnimplementedSecretServer) Get(context.Context, *GetSecretRequest) (*SecretInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSecretServer) List(context.Context, *SecretQuery) (*SecretsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedSecretServer) Stat(context.Context, *SecretStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedSecretServer) mustEmbedUnimplementedSecretServer() {}

// UnsafeSecretServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecretServer will
// result in compilation errors.
type UnsafeSecretServer interface {
	mustEmbedUnimplementedSecretServer()
}

func RegisterSecretServer(s grpc.ServiceRegistrar, srv SecretServer) {
	s.RegisterService(&Secret_ServiceDesc, srv)
}

func _Secret_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Secret_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServer).Create(ctx, req.(*CreateSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Secret_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Secret_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServer).Update(ctx, req.(*UpdateSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Secret_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Secret_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServer).Delete(ctx, req.(*SecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Secret_UpdateData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSecretDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServer).UpdateData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Secret_UpdateData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServer).UpdateData(ctx, req.(*UpdateSecretDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Secret_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Secret_GetData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServer).GetData(ctx, req.(*SecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Secret_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Secret_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServer).Get(ctx, req.(*GetSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Secret_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SecretQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Secret_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServer).List(ctx, req.(*SecretQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Secret_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SecretStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Secret_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServer).Stat(ctx, req.(*SecretStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// Secret_ServiceDesc is the grpc.ServiceDesc for Secret service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Secret_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.secret.v1.Secret",
	HandlerType: (*SecretServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _Secret_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _Secret_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _Secret_Delete_Handler,
		},
		{
			MethodName: "update_data",
			Handler:    _Secret_UpdateData_Handler,
		},
		{
			MethodName: "get_data",
			Handler:    _Secret_GetData_Handler,
		},
		{
			MethodName: "get",
			Handler:    _Secret_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _Secret_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _Secret_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/secret/v1/secret.proto",
}
