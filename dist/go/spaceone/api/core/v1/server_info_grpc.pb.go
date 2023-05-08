// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: spaceone/api/core/v1/server_info.proto

package v1

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ServerInfo_GetVersion_FullMethodName = "/spaceone.api.core.v1.ServerInfo/get_version"
)

// ServerInfoClient is the client API for ServerInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerInfoClient interface {
	GetVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VersionInfo, error)
}

type serverInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewServerInfoClient(cc grpc.ClientConnInterface) ServerInfoClient {
	return &serverInfoClient{cc}
}

func (c *serverInfoClient) GetVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VersionInfo, error) {
	out := new(VersionInfo)
	err := c.cc.Invoke(ctx, ServerInfo_GetVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerInfoServer is the server API for ServerInfo service.
// All implementations must embed UnimplementedServerInfoServer
// for forward compatibility
type ServerInfoServer interface {
	GetVersion(context.Context, *empty.Empty) (*VersionInfo, error)
	mustEmbedUnimplementedServerInfoServer()
}

// UnimplementedServerInfoServer must be embedded to have forward compatible implementations.
type UnimplementedServerInfoServer struct {
}

func (UnimplementedServerInfoServer) GetVersion(context.Context, *empty.Empty) (*VersionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (UnimplementedServerInfoServer) mustEmbedUnimplementedServerInfoServer() {}

// UnsafeServerInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerInfoServer will
// result in compilation errors.
type UnsafeServerInfoServer interface {
	mustEmbedUnimplementedServerInfoServer()
}

func RegisterServerInfoServer(s grpc.ServiceRegistrar, srv ServerInfoServer) {
	s.RegisterService(&ServerInfo_ServiceDesc, srv)
}

func _ServerInfo_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerInfoServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServerInfo_GetVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerInfoServer).GetVersion(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ServerInfo_ServiceDesc is the grpc.ServiceDesc for ServerInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.core.v1.ServerInfo",
	HandlerType: (*ServerInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get_version",
			Handler:    _ServerInfo_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/core/v1/server_info.proto",
}
