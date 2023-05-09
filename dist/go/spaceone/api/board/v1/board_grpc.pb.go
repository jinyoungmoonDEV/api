//A Board is a bulletin-board-type resource for posting notices and announcements in Cloudforet.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: spaceone/api/board/v1/board.proto

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
	Board_Create_FullMethodName        = "/spaceone.api.board.v1.Board/create"
	Board_Update_FullMethodName        = "/spaceone.api.board.v1.Board/update"
	Board_SetCategories_FullMethodName = "/spaceone.api.board.v1.Board/set_categories"
	Board_Delete_FullMethodName        = "/spaceone.api.board.v1.Board/delete"
	Board_Get_FullMethodName           = "/spaceone.api.board.v1.Board/get"
	Board_List_FullMethodName          = "/spaceone.api.board.v1.Board/list"
	Board_Stat_FullMethodName          = "/spaceone.api.board.v1.Board/stat"
)

// BoardClient is the client API for Board service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BoardClient interface {
	// Creates a new Board with SYSTEM permission. The `name` of the board is only required. You can add one or more `categories` representing the Board's attributes.
	Create(ctx context.Context, in *CreateBoardRequest, opts ...grpc.CallOption) (*BoardInfo, error)
	// Updates a specific Board with SYSTEM permission. You can make changes in Board settings, including `name` and `tags`.
	Update(ctx context.Context, in *UpdateBoardRequest, opts ...grpc.CallOption) (*BoardInfo, error)
	// Adds or changes `categories` of a specific Board with SYSTEM permission. A change in `categories` of a Board does not affect the `category` of the child Posts.
	SetCategories(ctx context.Context, in *SetBoardCategoriesRequest, opts ...grpc.CallOption) (*BoardInfo, error)
	// Deletes a specific Board with `SYSTEM` permission. You can delete a Board regardless of the presence of Posts created under the Board.
	Delete(ctx context.Context, in *BoardRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Gets a specific Board. You must specify the `board_id` of the Board to get. Prints detailed information about the Board, including `name`, `categories`.
	Get(ctx context.Context, in *GetBoardRequest, opts ...grpc.CallOption) (*BoardInfo, error)
	// Gets a list of all Boards. You can use a query to get a filtered list of Boards.
	List(ctx context.Context, in *BoardQuery, opts ...grpc.CallOption) (*BoardsInfo, error)
	Stat(ctx context.Context, in *BoardStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type boardClient struct {
	cc grpc.ClientConnInterface
}

func NewBoardClient(cc grpc.ClientConnInterface) BoardClient {
	return &boardClient{cc}
}

func (c *boardClient) Create(ctx context.Context, in *CreateBoardRequest, opts ...grpc.CallOption) (*BoardInfo, error) {
	out := new(BoardInfo)
	err := c.cc.Invoke(ctx, Board_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardClient) Update(ctx context.Context, in *UpdateBoardRequest, opts ...grpc.CallOption) (*BoardInfo, error) {
	out := new(BoardInfo)
	err := c.cc.Invoke(ctx, Board_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardClient) SetCategories(ctx context.Context, in *SetBoardCategoriesRequest, opts ...grpc.CallOption) (*BoardInfo, error) {
	out := new(BoardInfo)
	err := c.cc.Invoke(ctx, Board_SetCategories_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardClient) Delete(ctx context.Context, in *BoardRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Board_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardClient) Get(ctx context.Context, in *GetBoardRequest, opts ...grpc.CallOption) (*BoardInfo, error) {
	out := new(BoardInfo)
	err := c.cc.Invoke(ctx, Board_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardClient) List(ctx context.Context, in *BoardQuery, opts ...grpc.CallOption) (*BoardsInfo, error) {
	out := new(BoardsInfo)
	err := c.cc.Invoke(ctx, Board_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardClient) Stat(ctx context.Context, in *BoardStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, Board_Stat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BoardServer is the server API for Board service.
// All implementations must embed UnimplementedBoardServer
// for forward compatibility
type BoardServer interface {
	// Creates a new Board with SYSTEM permission. The `name` of the board is only required. You can add one or more `categories` representing the Board's attributes.
	Create(context.Context, *CreateBoardRequest) (*BoardInfo, error)
	// Updates a specific Board with SYSTEM permission. You can make changes in Board settings, including `name` and `tags`.
	Update(context.Context, *UpdateBoardRequest) (*BoardInfo, error)
	// Adds or changes `categories` of a specific Board with SYSTEM permission. A change in `categories` of a Board does not affect the `category` of the child Posts.
	SetCategories(context.Context, *SetBoardCategoriesRequest) (*BoardInfo, error)
	// Deletes a specific Board with `SYSTEM` permission. You can delete a Board regardless of the presence of Posts created under the Board.
	Delete(context.Context, *BoardRequest) (*empty.Empty, error)
	// Gets a specific Board. You must specify the `board_id` of the Board to get. Prints detailed information about the Board, including `name`, `categories`.
	Get(context.Context, *GetBoardRequest) (*BoardInfo, error)
	// Gets a list of all Boards. You can use a query to get a filtered list of Boards.
	List(context.Context, *BoardQuery) (*BoardsInfo, error)
	Stat(context.Context, *BoardStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedBoardServer()
}

// UnimplementedBoardServer must be embedded to have forward compatible implementations.
type UnimplementedBoardServer struct {
}

func (UnimplementedBoardServer) Create(context.Context, *CreateBoardRequest) (*BoardInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedBoardServer) Update(context.Context, *UpdateBoardRequest) (*BoardInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedBoardServer) SetCategories(context.Context, *SetBoardCategoriesRequest) (*BoardInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCategories not implemented")
}
func (UnimplementedBoardServer) Delete(context.Context, *BoardRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedBoardServer) Get(context.Context, *GetBoardRequest) (*BoardInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedBoardServer) List(context.Context, *BoardQuery) (*BoardsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedBoardServer) Stat(context.Context, *BoardStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedBoardServer) mustEmbedUnimplementedBoardServer() {}

// UnsafeBoardServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BoardServer will
// result in compilation errors.
type UnsafeBoardServer interface {
	mustEmbedUnimplementedBoardServer()
}

func RegisterBoardServer(s grpc.ServiceRegistrar, srv BoardServer) {
	s.RegisterService(&Board_ServiceDesc, srv)
}

func _Board_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Board_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServer).Create(ctx, req.(*CreateBoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Board_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Board_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServer).Update(ctx, req.(*UpdateBoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Board_SetCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetBoardCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServer).SetCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Board_SetCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServer).SetCategories(ctx, req.(*SetBoardCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Board_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Board_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServer).Delete(ctx, req.(*BoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Board_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Board_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServer).Get(ctx, req.(*GetBoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Board_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BoardQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Board_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServer).List(ctx, req.(*BoardQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Board_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BoardStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Board_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServer).Stat(ctx, req.(*BoardStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// Board_ServiceDesc is the grpc.ServiceDesc for Board service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Board_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.board.v1.Board",
	HandlerType: (*BoardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _Board_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _Board_Update_Handler,
		},
		{
			MethodName: "set_categories",
			Handler:    _Board_SetCategories_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _Board_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _Board_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _Board_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _Board_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/board/v1/board.proto",
}
