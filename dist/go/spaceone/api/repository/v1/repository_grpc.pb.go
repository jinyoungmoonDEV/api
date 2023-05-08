//
//desc: A Repository is a repository storing data of deployable plugins.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: spaceone/api/repository/v1/repository.proto

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
	Repository_Register_FullMethodName   = "/spaceone.api.repository.v1.Repository/register"
	Repository_Update_FullMethodName     = "/spaceone.api.repository.v1.Repository/update"
	Repository_Deregister_FullMethodName = "/spaceone.api.repository.v1.Repository/deregister"
	Repository_Get_FullMethodName        = "/spaceone.api.repository.v1.Repository/get"
	Repository_List_FullMethodName       = "/spaceone.api.repository.v1.Repository/list"
	Repository_Stat_FullMethodName       = "/spaceone.api.repository.v1.Repository/stat"
)

// RepositoryClient is the client API for Repository service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RepositoryClient interface {
	// desc: Registers a Repository. The parameter `name` can only include alphabets, numbers, and hyphens(-). The parameter `repository_type` can be either `local` or `remote`. The parameter `endpoint` is needed if the `repository_type` is `remote`.
	// request_example: >-
	// {
	// "name": "Local",
	// "repository_type": "local",
	// "endpoint": "grpc+ssl://local-url:443"
	// }
	// response_example: >-
	// {
	// "repository_id": "repo-123456789012",
	// "name": "Local",
	// "repository_type": "local",
	// "endpoint": "grpc+ssl://local-url:443",
	// "created_at": "2022-01-01T02:27:02.924Z"
	// }
	Register(ctx context.Context, in *CreateRepositoryRequest, opts ...grpc.CallOption) (*RepositoryInfo, error)
	// desc: Updates a specific Repository registered. You must specify the `repository_id` of the Repository to update. You can make changes in Repository settings, including `name`.
	// request_example: >-
	// {
	// "repository_id": "repo-123456789012",
	// "name": "Local-repo"
	// }
	// response_example: >-
	// {
	// "repository_id": "repo-123456789012",
	// "name": "Local-repo",
	// "repository_type": "local",
	// "endpoint": "grpc+ssl://local-url:443",
	// "created_at": "2022-01-01T02:27:02.924Z"
	// }
	Update(ctx context.Context, in *UpdateRepositoryRequest, opts ...grpc.CallOption) (*RepositoryInfo, error)
	// desc: Deregisters and deletes a specific Repository. You must specify the `repository_id` of the Repository to deregister.
	// request_example: >-
	// {
	// "repository_id": "repo-123456789012"
	// }
	Deregister(ctx context.Context, in *RepositoryRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// desc: Gets a specific Repository. Prints detailed information about the Repository, including  `name`, `repository_type`, and `endpoint`.
	// request_example: >-
	// {
	// "repository_id": "repo-123456789012"
	// }
	// response_example: >-
	// {
	// "repository_id": "repo-123456789012",
	// "name": "Local-repo",
	// "repository_type": "local",
	// "endpoint": "grpc+ssl://local-url:443",
	// "created_at": "2022-01-01T02:26:29.081Z"
	// }
	Get(ctx context.Context, in *GetRepositoryRequest, opts ...grpc.CallOption) (*RepositoryInfo, error)
	// desc: Gets a list of all Repositories regardless of `domain`. You can use a query to get a filtered list of Repositories.
	// request_example: >-
	// {
	// "query": {},
	// "repository_id": "repo-123456789012"
	// }
	// response_example: >-
	// {
	// "results": [
	// {
	// "repository_id": "repo-123456789012",
	// "name": "Local-repo",
	// "repository_type": "local",
	// "endpoint": "grpc+ssl://local-url:443",
	// "created_at": "2022-01-01T02:26:29.081Z"
	// }
	// ],
	// "total_count": 1
	// }
	List(ctx context.Context, in *RepositoryQuery, opts ...grpc.CallOption) (*RepositoriesInfo, error)
	Stat(ctx context.Context, in *RepositoryStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type repositoryClient struct {
	cc grpc.ClientConnInterface
}

func NewRepositoryClient(cc grpc.ClientConnInterface) RepositoryClient {
	return &repositoryClient{cc}
}

func (c *repositoryClient) Register(ctx context.Context, in *CreateRepositoryRequest, opts ...grpc.CallOption) (*RepositoryInfo, error) {
	out := new(RepositoryInfo)
	err := c.cc.Invoke(ctx, Repository_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryClient) Update(ctx context.Context, in *UpdateRepositoryRequest, opts ...grpc.CallOption) (*RepositoryInfo, error) {
	out := new(RepositoryInfo)
	err := c.cc.Invoke(ctx, Repository_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryClient) Deregister(ctx context.Context, in *RepositoryRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Repository_Deregister_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryClient) Get(ctx context.Context, in *GetRepositoryRequest, opts ...grpc.CallOption) (*RepositoryInfo, error) {
	out := new(RepositoryInfo)
	err := c.cc.Invoke(ctx, Repository_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryClient) List(ctx context.Context, in *RepositoryQuery, opts ...grpc.CallOption) (*RepositoriesInfo, error) {
	out := new(RepositoriesInfo)
	err := c.cc.Invoke(ctx, Repository_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryClient) Stat(ctx context.Context, in *RepositoryStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, Repository_Stat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RepositoryServer is the server API for Repository service.
// All implementations must embed UnimplementedRepositoryServer
// for forward compatibility
type RepositoryServer interface {
	// desc: Registers a Repository. The parameter `name` can only include alphabets, numbers, and hyphens(-). The parameter `repository_type` can be either `local` or `remote`. The parameter `endpoint` is needed if the `repository_type` is `remote`.
	// request_example: >-
	// {
	// "name": "Local",
	// "repository_type": "local",
	// "endpoint": "grpc+ssl://local-url:443"
	// }
	// response_example: >-
	// {
	// "repository_id": "repo-123456789012",
	// "name": "Local",
	// "repository_type": "local",
	// "endpoint": "grpc+ssl://local-url:443",
	// "created_at": "2022-01-01T02:27:02.924Z"
	// }
	Register(context.Context, *CreateRepositoryRequest) (*RepositoryInfo, error)
	// desc: Updates a specific Repository registered. You must specify the `repository_id` of the Repository to update. You can make changes in Repository settings, including `name`.
	// request_example: >-
	// {
	// "repository_id": "repo-123456789012",
	// "name": "Local-repo"
	// }
	// response_example: >-
	// {
	// "repository_id": "repo-123456789012",
	// "name": "Local-repo",
	// "repository_type": "local",
	// "endpoint": "grpc+ssl://local-url:443",
	// "created_at": "2022-01-01T02:27:02.924Z"
	// }
	Update(context.Context, *UpdateRepositoryRequest) (*RepositoryInfo, error)
	// desc: Deregisters and deletes a specific Repository. You must specify the `repository_id` of the Repository to deregister.
	// request_example: >-
	// {
	// "repository_id": "repo-123456789012"
	// }
	Deregister(context.Context, *RepositoryRequest) (*empty.Empty, error)
	// desc: Gets a specific Repository. Prints detailed information about the Repository, including  `name`, `repository_type`, and `endpoint`.
	// request_example: >-
	// {
	// "repository_id": "repo-123456789012"
	// }
	// response_example: >-
	// {
	// "repository_id": "repo-123456789012",
	// "name": "Local-repo",
	// "repository_type": "local",
	// "endpoint": "grpc+ssl://local-url:443",
	// "created_at": "2022-01-01T02:26:29.081Z"
	// }
	Get(context.Context, *GetRepositoryRequest) (*RepositoryInfo, error)
	// desc: Gets a list of all Repositories regardless of `domain`. You can use a query to get a filtered list of Repositories.
	// request_example: >-
	// {
	// "query": {},
	// "repository_id": "repo-123456789012"
	// }
	// response_example: >-
	// {
	// "results": [
	// {
	// "repository_id": "repo-123456789012",
	// "name": "Local-repo",
	// "repository_type": "local",
	// "endpoint": "grpc+ssl://local-url:443",
	// "created_at": "2022-01-01T02:26:29.081Z"
	// }
	// ],
	// "total_count": 1
	// }
	List(context.Context, *RepositoryQuery) (*RepositoriesInfo, error)
	Stat(context.Context, *RepositoryStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedRepositoryServer()
}

// UnimplementedRepositoryServer must be embedded to have forward compatible implementations.
type UnimplementedRepositoryServer struct {
}

func (UnimplementedRepositoryServer) Register(context.Context, *CreateRepositoryRequest) (*RepositoryInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedRepositoryServer) Update(context.Context, *UpdateRepositoryRequest) (*RepositoryInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRepositoryServer) Deregister(context.Context, *RepositoryRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deregister not implemented")
}
func (UnimplementedRepositoryServer) Get(context.Context, *GetRepositoryRequest) (*RepositoryInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRepositoryServer) List(context.Context, *RepositoryQuery) (*RepositoriesInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRepositoryServer) Stat(context.Context, *RepositoryStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedRepositoryServer) mustEmbedUnimplementedRepositoryServer() {}

// UnsafeRepositoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RepositoryServer will
// result in compilation errors.
type UnsafeRepositoryServer interface {
	mustEmbedUnimplementedRepositoryServer()
}

func RegisterRepositoryServer(s grpc.ServiceRegistrar, srv RepositoryServer) {
	s.RegisterService(&Repository_ServiceDesc, srv)
}

func _Repository_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Repository_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServer).Register(ctx, req.(*CreateRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repository_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Repository_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServer).Update(ctx, req.(*UpdateRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repository_Deregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServer).Deregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Repository_Deregister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServer).Deregister(ctx, req.(*RepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repository_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Repository_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServer).Get(ctx, req.(*GetRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repository_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepositoryQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Repository_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServer).List(ctx, req.(*RepositoryQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repository_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepositoryStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Repository_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServer).Stat(ctx, req.(*RepositoryStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// Repository_ServiceDesc is the grpc.ServiceDesc for Repository service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Repository_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.repository.v1.Repository",
	HandlerType: (*RepositoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "register",
			Handler:    _Repository_Register_Handler,
		},
		{
			MethodName: "update",
			Handler:    _Repository_Update_Handler,
		},
		{
			MethodName: "deregister",
			Handler:    _Repository_Deregister_Handler,
		},
		{
			MethodName: "get",
			Handler:    _Repository_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _Repository_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _Repository_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/repository/v1/repository.proto",
}
