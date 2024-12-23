// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: spaceone/api/opsflow/v1/comment.proto

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
	Comment_Create_FullMethodName = "/spaceone.api.opsflow.v1.Comment/create"
	Comment_Update_FullMethodName = "/spaceone.api.opsflow.v1.Comment/update"
	Comment_Delete_FullMethodName = "/spaceone.api.opsflow.v1.Comment/delete"
	Comment_Get_FullMethodName    = "/spaceone.api.opsflow.v1.Comment/get"
	Comment_List_FullMethodName   = "/spaceone.api.opsflow.v1.Comment/list"
	Comment_Stat_FullMethodName   = "/spaceone.api.opsflow.v1.Comment/stat"
)

// CommentClient is the client API for Comment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentClient interface {
	Create(ctx context.Context, in *CommentCreateRequest, opts ...grpc.CallOption) (*CommentInfo, error)
	Update(ctx context.Context, in *CommentUpdateRequest, opts ...grpc.CallOption) (*CommentInfo, error)
	Delete(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Get(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentInfo, error)
	List(ctx context.Context, in *CommentSearchQuery, opts ...grpc.CallOption) (*CommentsInfo, error)
	Stat(ctx context.Context, in *CommentStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type commentClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentClient(cc grpc.ClientConnInterface) CommentClient {
	return &commentClient{cc}
}

func (c *commentClient) Create(ctx context.Context, in *CommentCreateRequest, opts ...grpc.CallOption) (*CommentInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommentInfo)
	err := c.cc.Invoke(ctx, Comment_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) Update(ctx context.Context, in *CommentUpdateRequest, opts ...grpc.CallOption) (*CommentInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommentInfo)
	err := c.cc.Invoke(ctx, Comment_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) Delete(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Comment_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) Get(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommentInfo)
	err := c.cc.Invoke(ctx, Comment_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) List(ctx context.Context, in *CommentSearchQuery, opts ...grpc.CallOption) (*CommentsInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommentsInfo)
	err := c.cc.Invoke(ctx, Comment_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) Stat(ctx context.Context, in *CommentStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, Comment_Stat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServer is the server API for Comment service.
// All implementations must embed UnimplementedCommentServer
// for forward compatibility.
type CommentServer interface {
	Create(context.Context, *CommentCreateRequest) (*CommentInfo, error)
	Update(context.Context, *CommentUpdateRequest) (*CommentInfo, error)
	Delete(context.Context, *CommentRequest) (*empty.Empty, error)
	Get(context.Context, *CommentRequest) (*CommentInfo, error)
	List(context.Context, *CommentSearchQuery) (*CommentsInfo, error)
	Stat(context.Context, *CommentStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedCommentServer()
}

// UnimplementedCommentServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCommentServer struct{}

func (UnimplementedCommentServer) Create(context.Context, *CommentCreateRequest) (*CommentInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCommentServer) Update(context.Context, *CommentUpdateRequest) (*CommentInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCommentServer) Delete(context.Context, *CommentRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCommentServer) Get(context.Context, *CommentRequest) (*CommentInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCommentServer) List(context.Context, *CommentSearchQuery) (*CommentsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCommentServer) Stat(context.Context, *CommentStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedCommentServer) mustEmbedUnimplementedCommentServer() {}
func (UnimplementedCommentServer) testEmbeddedByValue()                 {}

// UnsafeCommentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentServer will
// result in compilation errors.
type UnsafeCommentServer interface {
	mustEmbedUnimplementedCommentServer()
}

func RegisterCommentServer(s grpc.ServiceRegistrar, srv CommentServer) {
	// If the following call pancis, it indicates UnimplementedCommentServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Comment_ServiceDesc, srv)
}

func _Comment_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Comment_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).Create(ctx, req.(*CommentCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Comment_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).Update(ctx, req.(*CommentUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Comment_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).Delete(ctx, req.(*CommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Comment_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).Get(ctx, req.(*CommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentSearchQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Comment_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).List(ctx, req.(*CommentSearchQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Comment_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).Stat(ctx, req.(*CommentStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// Comment_ServiceDesc is the grpc.ServiceDesc for Comment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Comment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.opsflow.v1.Comment",
	HandlerType: (*CommentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _Comment_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _Comment_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _Comment_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _Comment_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _Comment_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _Comment_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/opsflow/v1/comment.proto",
}
