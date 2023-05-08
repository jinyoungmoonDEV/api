// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: spaceone/api/cost_analysis/v1/budget_usage.proto

package v1

import (
	context "context"
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
	BudgetUsage_List_FullMethodName    = "/spaceone.api.cost_analysis.v1.BudgetUsage/list"
	BudgetUsage_Analyze_FullMethodName = "/spaceone.api.cost_analysis.v1.BudgetUsage/analyze"
	BudgetUsage_Stat_FullMethodName    = "/spaceone.api.cost_analysis.v1.BudgetUsage/stat"
)

// BudgetUsageClient is the client API for BudgetUsage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BudgetUsageClient interface {
	// desc: Gets a list of all BudgetUsages. You can use a query to get a filtered list of BudgetUsages.
	// request_example: >-
	// {
	// "query": {}
	// }
	// response_example: >-
	// {
	// "results": [
	// {
	// "budget_id": "budget-abb377eb9e8b",
	// "name": "Cloudforet-Budget3",
	// "date": "2022-01",
	// "usd_cost": 7671.164,
	// "limit": 10000.0,
	// "domain_id": "domain-58010aa2e451",
	// "updated_at": "2022-07-19T04:26:08.099Z"
	// },
	// {
	// "budget_id": "budget-abb377eb9e8b",
	// "name": "Cloudforet-Budget3",
	// "date": "2022-02",
	// "usd_cost": 5931.771,
	// "limit": 11000.0,
	// "domain_id": "domain-58010aa2e451",
	// "updated_at": "2022-07-19T04:26:08.105Z"
	// }
	// ],
	// "total_count": 12
	// }
	List(ctx context.Context, in *BudgetUsageQuery, opts ...grpc.CallOption) (*BudgetUsagesInfo, error)
	Analyze(ctx context.Context, in *BudgetUsageAnalyzeQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
	Stat(ctx context.Context, in *BudgetUsageStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type budgetUsageClient struct {
	cc grpc.ClientConnInterface
}

func NewBudgetUsageClient(cc grpc.ClientConnInterface) BudgetUsageClient {
	return &budgetUsageClient{cc}
}

func (c *budgetUsageClient) List(ctx context.Context, in *BudgetUsageQuery, opts ...grpc.CallOption) (*BudgetUsagesInfo, error) {
	out := new(BudgetUsagesInfo)
	err := c.cc.Invoke(ctx, BudgetUsage_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *budgetUsageClient) Analyze(ctx context.Context, in *BudgetUsageAnalyzeQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, BudgetUsage_Analyze_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *budgetUsageClient) Stat(ctx context.Context, in *BudgetUsageStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, BudgetUsage_Stat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BudgetUsageServer is the server API for BudgetUsage service.
// All implementations must embed UnimplementedBudgetUsageServer
// for forward compatibility
type BudgetUsageServer interface {
	// desc: Gets a list of all BudgetUsages. You can use a query to get a filtered list of BudgetUsages.
	// request_example: >-
	// {
	// "query": {}
	// }
	// response_example: >-
	// {
	// "results": [
	// {
	// "budget_id": "budget-abb377eb9e8b",
	// "name": "Cloudforet-Budget3",
	// "date": "2022-01",
	// "usd_cost": 7671.164,
	// "limit": 10000.0,
	// "domain_id": "domain-58010aa2e451",
	// "updated_at": "2022-07-19T04:26:08.099Z"
	// },
	// {
	// "budget_id": "budget-abb377eb9e8b",
	// "name": "Cloudforet-Budget3",
	// "date": "2022-02",
	// "usd_cost": 5931.771,
	// "limit": 11000.0,
	// "domain_id": "domain-58010aa2e451",
	// "updated_at": "2022-07-19T04:26:08.105Z"
	// }
	// ],
	// "total_count": 12
	// }
	List(context.Context, *BudgetUsageQuery) (*BudgetUsagesInfo, error)
	Analyze(context.Context, *BudgetUsageAnalyzeQuery) (*_struct.Struct, error)
	Stat(context.Context, *BudgetUsageStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedBudgetUsageServer()
}

// UnimplementedBudgetUsageServer must be embedded to have forward compatible implementations.
type UnimplementedBudgetUsageServer struct {
}

func (UnimplementedBudgetUsageServer) List(context.Context, *BudgetUsageQuery) (*BudgetUsagesInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedBudgetUsageServer) Analyze(context.Context, *BudgetUsageAnalyzeQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Analyze not implemented")
}
func (UnimplementedBudgetUsageServer) Stat(context.Context, *BudgetUsageStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedBudgetUsageServer) mustEmbedUnimplementedBudgetUsageServer() {}

// UnsafeBudgetUsageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BudgetUsageServer will
// result in compilation errors.
type UnsafeBudgetUsageServer interface {
	mustEmbedUnimplementedBudgetUsageServer()
}

func RegisterBudgetUsageServer(s grpc.ServiceRegistrar, srv BudgetUsageServer) {
	s.RegisterService(&BudgetUsage_ServiceDesc, srv)
}

func _BudgetUsage_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BudgetUsageQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetUsageServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BudgetUsage_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetUsageServer).List(ctx, req.(*BudgetUsageQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _BudgetUsage_Analyze_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BudgetUsageAnalyzeQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetUsageServer).Analyze(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BudgetUsage_Analyze_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetUsageServer).Analyze(ctx, req.(*BudgetUsageAnalyzeQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _BudgetUsage_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BudgetUsageStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetUsageServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BudgetUsage_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetUsageServer).Stat(ctx, req.(*BudgetUsageStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// BudgetUsage_ServiceDesc is the grpc.ServiceDesc for BudgetUsage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BudgetUsage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.cost_analysis.v1.BudgetUsage",
	HandlerType: (*BudgetUsageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "list",
			Handler:    _BudgetUsage_List_Handler,
		},
		{
			MethodName: "analyze",
			Handler:    _BudgetUsage_Analyze_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _BudgetUsage_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/cost_analysis/v1/budget_usage.proto",
}
