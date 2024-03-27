// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: spaceone/api/identity/v1/policy.proto

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
	Policy_Create_FullMethodName = "/spaceone.api.identity.v1.Policy/create"
	Policy_Update_FullMethodName = "/spaceone.api.identity.v1.Policy/update"
	Policy_Delete_FullMethodName = "/spaceone.api.identity.v1.Policy/delete"
	Policy_Get_FullMethodName    = "/spaceone.api.identity.v1.Policy/get"
	Policy_List_FullMethodName   = "/spaceone.api.identity.v1.Policy/list"
	Policy_Stat_FullMethodName   = "/spaceone.api.identity.v1.Policy/stat"
)

// PolicyClient is the client API for Policy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PolicyClient interface {
	Create(ctx context.Context, in *CreatePolicyRequest, opts ...grpc.CallOption) (*PolicyInfo, error)
	Update(ctx context.Context, in *UpdatePolicyRequest, opts ...grpc.CallOption) (*PolicyInfo, error)
	Delete(ctx context.Context, in *PolicyRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Get(ctx context.Context, in *GetPolicyRequest, opts ...grpc.CallOption) (*PolicyInfo, error)
	List(ctx context.Context, in *PolicyQuery, opts ...grpc.CallOption) (*PoliciesInfo, error)
	Stat(ctx context.Context, in *PolicyStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type policyClient struct {
	cc grpc.ClientConnInterface
}

func NewPolicyClient(cc grpc.ClientConnInterface) PolicyClient {
	return &policyClient{cc}
}

func (c *policyClient) Create(ctx context.Context, in *CreatePolicyRequest, opts ...grpc.CallOption) (*PolicyInfo, error) {
	out := new(PolicyInfo)
	err := c.cc.Invoke(ctx, Policy_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyClient) Update(ctx context.Context, in *UpdatePolicyRequest, opts ...grpc.CallOption) (*PolicyInfo, error) {
	out := new(PolicyInfo)
	err := c.cc.Invoke(ctx, Policy_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyClient) Delete(ctx context.Context, in *PolicyRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Policy_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyClient) Get(ctx context.Context, in *GetPolicyRequest, opts ...grpc.CallOption) (*PolicyInfo, error) {
	out := new(PolicyInfo)
	err := c.cc.Invoke(ctx, Policy_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyClient) List(ctx context.Context, in *PolicyQuery, opts ...grpc.CallOption) (*PoliciesInfo, error) {
	out := new(PoliciesInfo)
	err := c.cc.Invoke(ctx, Policy_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyClient) Stat(ctx context.Context, in *PolicyStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, Policy_Stat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolicyServer is the server API for Policy service.
// All implementations must embed UnimplementedPolicyServer
// for forward compatibility
type PolicyServer interface {
	Create(context.Context, *CreatePolicyRequest) (*PolicyInfo, error)
	Update(context.Context, *UpdatePolicyRequest) (*PolicyInfo, error)
	Delete(context.Context, *PolicyRequest) (*empty.Empty, error)
	Get(context.Context, *GetPolicyRequest) (*PolicyInfo, error)
	List(context.Context, *PolicyQuery) (*PoliciesInfo, error)
	Stat(context.Context, *PolicyStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedPolicyServer()
}

// UnimplementedPolicyServer must be embedded to have forward compatible implementations.
type UnimplementedPolicyServer struct {
}

func (UnimplementedPolicyServer) Create(context.Context, *CreatePolicyRequest) (*PolicyInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPolicyServer) Update(context.Context, *UpdatePolicyRequest) (*PolicyInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPolicyServer) Delete(context.Context, *PolicyRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPolicyServer) Get(context.Context, *GetPolicyRequest) (*PolicyInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPolicyServer) List(context.Context, *PolicyQuery) (*PoliciesInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPolicyServer) Stat(context.Context, *PolicyStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedPolicyServer) mustEmbedUnimplementedPolicyServer() {}

// UnsafePolicyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PolicyServer will
// result in compilation errors.
type UnsafePolicyServer interface {
	mustEmbedUnimplementedPolicyServer()
}

func RegisterPolicyServer(s grpc.ServiceRegistrar, srv PolicyServer) {
	s.RegisterService(&Policy_ServiceDesc, srv)
}

func _Policy_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Policy_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServer).Create(ctx, req.(*CreatePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Policy_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Policy_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServer).Update(ctx, req.(*UpdatePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Policy_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Policy_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServer).Delete(ctx, req.(*PolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Policy_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Policy_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServer).Get(ctx, req.(*GetPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Policy_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Policy_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServer).List(ctx, req.(*PolicyQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Policy_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Policy_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServer).Stat(ctx, req.(*PolicyStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// Policy_ServiceDesc is the grpc.ServiceDesc for Policy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Policy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.identity.v1.Policy",
	HandlerType: (*PolicyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _Policy_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _Policy_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _Policy_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _Policy_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _Policy_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _Policy_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/identity/v1/policy.proto",
}
