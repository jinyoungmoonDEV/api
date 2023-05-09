//A Post is a message published on a Board. It also provides notifications to Projects affected by the Post.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: spaceone/api/board/v1/post.proto

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
	Post_Create_FullMethodName           = "/spaceone.api.board.v1.Post/create"
	Post_Update_FullMethodName           = "/spaceone.api.board.v1.Post/update"
	Post_SendNotification_FullMethodName = "/spaceone.api.board.v1.Post/send_notification"
	Post_Delete_FullMethodName           = "/spaceone.api.board.v1.Post/delete"
	Post_Get_FullMethodName              = "/spaceone.api.board.v1.Post/get"
	Post_List_FullMethodName             = "/spaceone.api.board.v1.Post/list"
	Post_Stat_FullMethodName             = "/spaceone.api.board.v1.Post/stat"
)

// PostClient is the client API for Post service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostClient interface {
	// Creates a new Post under a specific Board. You must specify the `board_id`, `title`, and `contents`. The parameter `category` is not required but can be set in the scope of `categories` specified in the parent Board. You can make the new Post pinned or pop up by adjusting the parameters.
	Create(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*PostInfo, error)
	// Updates a specific Post. You can make changes in Post settings, except `board_id`, `post_id`, and `domain_id`.
	Update(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*PostInfo, error)
	// Not Implemented
	SendNotification(ctx context.Context, in *SendNotificationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Deletes a specific Post. You must specify the `post_id` of the Post to delete, and the `board_id` of the Board where the child Post to delete belongs.
	Delete(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Gets a specific Post. You must specify the `post_id` of the Post to get, and the `board_id` of the Board where the child Post to get belongs. Prints detailed information about the Post.
	Get(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*PostInfo, error)
	// Gets a list of all Posts. You can use a query to get a filtered list of Posts.
	List(ctx context.Context, in *PostQuery, opts ...grpc.CallOption) (*PostsInfo, error)
	Stat(ctx context.Context, in *PostStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type postClient struct {
	cc grpc.ClientConnInterface
}

func NewPostClient(cc grpc.ClientConnInterface) PostClient {
	return &postClient{cc}
}

func (c *postClient) Create(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*PostInfo, error) {
	out := new(PostInfo)
	err := c.cc.Invoke(ctx, Post_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postClient) Update(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*PostInfo, error) {
	out := new(PostInfo)
	err := c.cc.Invoke(ctx, Post_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postClient) SendNotification(ctx context.Context, in *SendNotificationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Post_SendNotification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postClient) Delete(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Post_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postClient) Get(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*PostInfo, error) {
	out := new(PostInfo)
	err := c.cc.Invoke(ctx, Post_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postClient) List(ctx context.Context, in *PostQuery, opts ...grpc.CallOption) (*PostsInfo, error) {
	out := new(PostsInfo)
	err := c.cc.Invoke(ctx, Post_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postClient) Stat(ctx context.Context, in *PostStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, Post_Stat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServer is the server API for Post service.
// All implementations must embed UnimplementedPostServer
// for forward compatibility
type PostServer interface {
	// Creates a new Post under a specific Board. You must specify the `board_id`, `title`, and `contents`. The parameter `category` is not required but can be set in the scope of `categories` specified in the parent Board. You can make the new Post pinned or pop up by adjusting the parameters.
	Create(context.Context, *CreatePostRequest) (*PostInfo, error)
	// Updates a specific Post. You can make changes in Post settings, except `board_id`, `post_id`, and `domain_id`.
	Update(context.Context, *UpdatePostRequest) (*PostInfo, error)
	// Not Implemented
	SendNotification(context.Context, *SendNotificationRequest) (*empty.Empty, error)
	// Deletes a specific Post. You must specify the `post_id` of the Post to delete, and the `board_id` of the Board where the child Post to delete belongs.
	Delete(context.Context, *PostRequest) (*empty.Empty, error)
	// Gets a specific Post. You must specify the `post_id` of the Post to get, and the `board_id` of the Board where the child Post to get belongs. Prints detailed information about the Post.
	Get(context.Context, *GetPostRequest) (*PostInfo, error)
	// Gets a list of all Posts. You can use a query to get a filtered list of Posts.
	List(context.Context, *PostQuery) (*PostsInfo, error)
	Stat(context.Context, *PostStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedPostServer()
}

// UnimplementedPostServer must be embedded to have forward compatible implementations.
type UnimplementedPostServer struct {
}

func (UnimplementedPostServer) Create(context.Context, *CreatePostRequest) (*PostInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPostServer) Update(context.Context, *UpdatePostRequest) (*PostInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPostServer) SendNotification(context.Context, *SendNotificationRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendNotification not implemented")
}
func (UnimplementedPostServer) Delete(context.Context, *PostRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPostServer) Get(context.Context, *GetPostRequest) (*PostInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPostServer) List(context.Context, *PostQuery) (*PostsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPostServer) Stat(context.Context, *PostStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedPostServer) mustEmbedUnimplementedPostServer() {}

// UnsafePostServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostServer will
// result in compilation errors.
type UnsafePostServer interface {
	mustEmbedUnimplementedPostServer()
}

func RegisterPostServer(s grpc.ServiceRegistrar, srv PostServer) {
	s.RegisterService(&Post_ServiceDesc, srv)
}

func _Post_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Post_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).Create(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Post_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Post_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).Update(ctx, req.(*UpdatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Post_SendNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).SendNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Post_SendNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).SendNotification(ctx, req.(*SendNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Post_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Post_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).Delete(ctx, req.(*PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Post_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Post_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).Get(ctx, req.(*GetPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Post_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Post_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).List(ctx, req.(*PostQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Post_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Post_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).Stat(ctx, req.(*PostStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// Post_ServiceDesc is the grpc.ServiceDesc for Post service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Post_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.board.v1.Post",
	HandlerType: (*PostServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _Post_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _Post_Update_Handler,
		},
		{
			MethodName: "send_notification",
			Handler:    _Post_SendNotification_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _Post_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _Post_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _Post_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _Post_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/board/v1/post.proto",
}
