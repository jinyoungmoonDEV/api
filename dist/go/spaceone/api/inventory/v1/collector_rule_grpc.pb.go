//
//desc: A CollectorRule is a cloud service resource filtering the raw data from the Collector. The Cloud Service resource is created after the raw data is filtered by the CollectorRule.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: spaceone/api/inventory/v1/collector_rule.proto

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
	CollectorRule_Create_FullMethodName      = "/spaceone.api.inventory.v1.CollectorRule/create"
	CollectorRule_Update_FullMethodName      = "/spaceone.api.inventory.v1.CollectorRule/update"
	CollectorRule_ChangeOrder_FullMethodName = "/spaceone.api.inventory.v1.CollectorRule/change_order"
	CollectorRule_Delete_FullMethodName      = "/spaceone.api.inventory.v1.CollectorRule/delete"
	CollectorRule_Get_FullMethodName         = "/spaceone.api.inventory.v1.CollectorRule/get"
	CollectorRule_List_FullMethodName        = "/spaceone.api.inventory.v1.CollectorRule/list"
	CollectorRule_Stat_FullMethodName        = "/spaceone.api.inventory.v1.CollectorRule/stat"
)

// CollectorRuleClient is the client API for CollectorRule service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CollectorRuleClient interface {
	// desc: Creates a new CollectorRule. When creating the cloud service, this method can apply two types of conditions: mapping projects where the cloud service incurred to the Cloud Service, and mapping cloud service accounts to the Cloud Service. By adjusting the `condition_policy` parameter, the CollectorRule can be applied when all conditions are met, applied when any of the conditions are met, or always applied regardless of whether the conditions are met.
	// request_example: >-
	// {
	// "name": "match_service_account_test",
	// "conditions_policy": "ALWAYS",
	// "actions": {
	// "match_service_account": {"source": "account", "target": "data.project_id"}
	// },
	// "options": {"stop_processing": true},
	// "tags": {"b": "c", "a": "b"},
	// "collector_id": "collector-2e39891cbbb5"
	// }
	// response_example: >-
	// {
	// "collector_rule_id": "collector-rule-c8055231e212",
	// "name": "match_service_account_test",
	// "order": 2,
	// "conditions_policy": "ALWAYS",
	// "actions": {
	// "match_service_account": {
	// "source": "account",
	// "target": "data.project_id"
	// }
	// },
	// "options": {
	// "stop_processing": true
	// },
	// "tags": {
	// "a": "b",
	// "b": "c"
	// },
	// "collector_id": "collector-2e39891cbbb5",
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2022-07-19T10:13:28.335Z"
	// }
	Create(ctx context.Context, in *CreateCollectorRuleRequest, opts ...grpc.CallOption) (*CollectorRuleInfo, error)
	// desc: Updates a specific CollectorRule. You can make changes in CollectorRule settings, including filtering conditions.
	// request_example: >-
	// {
	// "collector_rule_id": "collector-rule-c8055231e212",
	// "name": "match_service_account_test",
	// "conditions_policy": "ALWAYS",
	// "actions": {
	// "match_service_account": {
	// "source": "account",
	// "target": "data.project_id"
	// }
	// },
	// "options": {
	// "stop_processing": true
	// },
	// "tags": {"b": "c", "a": "b"}
	// }
	// response_example: >-
	// {
	// "collector_rule_id": "collector-rule-c8055231e212",
	// "name": "match_service_account_test",
	// "order": 2,
	// "conditions_policy": "ALWAYS",
	// "actions": {
	// "match_service_account": {
	// "source": "account",
	// "target": "data.project_id"
	// }
	// },
	// "options": {
	// "stop_processing": true
	// },
	// "tags": {
	// "a": "b",
	// "b": "c"
	// },
	// "collector_id": "collector-2e39891cbbb5",
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2022-07-19T10:13:28.335Z"
	// }
	Update(ctx context.Context, in *UpdateCollectorRuleRequest, opts ...grpc.CallOption) (*CollectorRuleInfo, error)
	// desc: Changes the priority order of the CollectorRules to apply. If there are multiple CollectorRules applied in a specific service account, the priority order of the resources is required. This method changes the priority order to apply CollectorRules.
	// request_example: >-
	// {
	// "collector_rule_id": "collector-rule-c8055231e212",
	// "order": 2
	// }
	// response_example: >-
	// {
	// "collector_rule_id": "collector-rule-c8055231e212",
	// "name": "match_service_account_test",
	// "order": 2,
	// "conditions_policy": "ALWAYS",
	// "actions": {
	// "match_service_account": {
	// "source": "account",
	// "target": "data.project_id"
	// }
	// },
	// "options": {
	// "stop_processing": true
	// },
	// "tags": {
	// "a": "b",
	// "b": "c"
	// },
	// "collector_id": "collector-2e39891cbbb5",
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2022-07-19T10:13:28.335Z"
	// }
	ChangeOrder(ctx context.Context, in *ChangeCollectorRuleOrderRequest, opts ...grpc.CallOption) (*CollectorRuleInfo, error)
	// desc: Deletes a specific CollectorRule. You must specify the `collector_rule_id` of the CollectorRule to delete.
	// request_example: >-
	// {
	// "collector_rule_id": "collector-rule-c8055231e212",
	// }
	Delete(ctx context.Context, in *CollectorRuleRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// desc: Gets a specific CollectorRule. Prints detailed information about the CollectorRule, including  `conditions_policy` and conditions applied to Collector.
	// request_example: >-
	// {
	// "collector_rule_id": "collector-rule-c8055231e212",
	// }
	// response_example: >-
	// {
	// "collector_rule_id": "collector-rule-c8055231e212",
	// "name": "match_service_account",
	// "order": 1,
	// "conditions_policy": "ALWAYS",
	// "actions": {
	// "match_service_account": {
	// "source": "account",
	// "target": "data.project_id"
	// }
	// },
	// "options": {
	// "stop_processing": true
	// },
	// "tags": {},
	// "collector_id": "collector-2e39891cbbb5",
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2022-05-25T16:01:51.858Z"
	// }
	Get(ctx context.Context, in *GetCollectorRuleRequest, opts ...grpc.CallOption) (*CollectorRuleInfo, error)
	// desc: Gets a list of all CollectorRules. You can use a query to get a filtered list of CollectorRules.
	// request_example: >-
	// {
	// "query": {}
	// }
	// response_example: >-
	// {
	// "results": [
	// {
	// "collector_rule_id": "collector-rule-c8055231e212",
	// "name": "match_service_account",
	// "order": 1,
	// "conditions_policy": "ALWAYS",
	// "actions": {
	// "match_service_account": {
	// "source": "account",
	// "target": "data.project_id"
	// }
	// },
	// "options": {
	// "stop_processing": true
	// },
	// "tags": {},
	// "collector_id": "collector-2e39891cbbb5",
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2022-05-25T16:01:51.858Z"
	// },
	// {
	// "collector_rule_id": "collector-rule-t3345231e167",
	// "name": "match_service_account",
	// "order": 1,
	// "conditions_policy": "ALWAYS",
	// "actions": {
	// "match_service_account": {
	// "source": "account",
	// "target": "data.account_id"
	// }
	// },
	// "options": {
	// "stop_processing": true
	// },
	// "tags": {},
	// "collector_id": "collector-7163022d49a1",
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2022-06-03T16:00:54.099Z"
	// }
	// ],
	// "total_count": 2
	// }
	List(ctx context.Context, in *CollectorRuleQuery, opts ...grpc.CallOption) (*CollectorRulesInfo, error)
	Stat(ctx context.Context, in *CollectorRuleStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type collectorRuleClient struct {
	cc grpc.ClientConnInterface
}

func NewCollectorRuleClient(cc grpc.ClientConnInterface) CollectorRuleClient {
	return &collectorRuleClient{cc}
}

func (c *collectorRuleClient) Create(ctx context.Context, in *CreateCollectorRuleRequest, opts ...grpc.CallOption) (*CollectorRuleInfo, error) {
	out := new(CollectorRuleInfo)
	err := c.cc.Invoke(ctx, CollectorRule_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectorRuleClient) Update(ctx context.Context, in *UpdateCollectorRuleRequest, opts ...grpc.CallOption) (*CollectorRuleInfo, error) {
	out := new(CollectorRuleInfo)
	err := c.cc.Invoke(ctx, CollectorRule_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectorRuleClient) ChangeOrder(ctx context.Context, in *ChangeCollectorRuleOrderRequest, opts ...grpc.CallOption) (*CollectorRuleInfo, error) {
	out := new(CollectorRuleInfo)
	err := c.cc.Invoke(ctx, CollectorRule_ChangeOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectorRuleClient) Delete(ctx context.Context, in *CollectorRuleRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, CollectorRule_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectorRuleClient) Get(ctx context.Context, in *GetCollectorRuleRequest, opts ...grpc.CallOption) (*CollectorRuleInfo, error) {
	out := new(CollectorRuleInfo)
	err := c.cc.Invoke(ctx, CollectorRule_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectorRuleClient) List(ctx context.Context, in *CollectorRuleQuery, opts ...grpc.CallOption) (*CollectorRulesInfo, error) {
	out := new(CollectorRulesInfo)
	err := c.cc.Invoke(ctx, CollectorRule_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectorRuleClient) Stat(ctx context.Context, in *CollectorRuleStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, CollectorRule_Stat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollectorRuleServer is the server API for CollectorRule service.
// All implementations must embed UnimplementedCollectorRuleServer
// for forward compatibility
type CollectorRuleServer interface {
	// desc: Creates a new CollectorRule. When creating the cloud service, this method can apply two types of conditions: mapping projects where the cloud service incurred to the Cloud Service, and mapping cloud service accounts to the Cloud Service. By adjusting the `condition_policy` parameter, the CollectorRule can be applied when all conditions are met, applied when any of the conditions are met, or always applied regardless of whether the conditions are met.
	// request_example: >-
	// {
	// "name": "match_service_account_test",
	// "conditions_policy": "ALWAYS",
	// "actions": {
	// "match_service_account": {"source": "account", "target": "data.project_id"}
	// },
	// "options": {"stop_processing": true},
	// "tags": {"b": "c", "a": "b"},
	// "collector_id": "collector-2e39891cbbb5"
	// }
	// response_example: >-
	// {
	// "collector_rule_id": "collector-rule-c8055231e212",
	// "name": "match_service_account_test",
	// "order": 2,
	// "conditions_policy": "ALWAYS",
	// "actions": {
	// "match_service_account": {
	// "source": "account",
	// "target": "data.project_id"
	// }
	// },
	// "options": {
	// "stop_processing": true
	// },
	// "tags": {
	// "a": "b",
	// "b": "c"
	// },
	// "collector_id": "collector-2e39891cbbb5",
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2022-07-19T10:13:28.335Z"
	// }
	Create(context.Context, *CreateCollectorRuleRequest) (*CollectorRuleInfo, error)
	// desc: Updates a specific CollectorRule. You can make changes in CollectorRule settings, including filtering conditions.
	// request_example: >-
	// {
	// "collector_rule_id": "collector-rule-c8055231e212",
	// "name": "match_service_account_test",
	// "conditions_policy": "ALWAYS",
	// "actions": {
	// "match_service_account": {
	// "source": "account",
	// "target": "data.project_id"
	// }
	// },
	// "options": {
	// "stop_processing": true
	// },
	// "tags": {"b": "c", "a": "b"}
	// }
	// response_example: >-
	// {
	// "collector_rule_id": "collector-rule-c8055231e212",
	// "name": "match_service_account_test",
	// "order": 2,
	// "conditions_policy": "ALWAYS",
	// "actions": {
	// "match_service_account": {
	// "source": "account",
	// "target": "data.project_id"
	// }
	// },
	// "options": {
	// "stop_processing": true
	// },
	// "tags": {
	// "a": "b",
	// "b": "c"
	// },
	// "collector_id": "collector-2e39891cbbb5",
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2022-07-19T10:13:28.335Z"
	// }
	Update(context.Context, *UpdateCollectorRuleRequest) (*CollectorRuleInfo, error)
	// desc: Changes the priority order of the CollectorRules to apply. If there are multiple CollectorRules applied in a specific service account, the priority order of the resources is required. This method changes the priority order to apply CollectorRules.
	// request_example: >-
	// {
	// "collector_rule_id": "collector-rule-c8055231e212",
	// "order": 2
	// }
	// response_example: >-
	// {
	// "collector_rule_id": "collector-rule-c8055231e212",
	// "name": "match_service_account_test",
	// "order": 2,
	// "conditions_policy": "ALWAYS",
	// "actions": {
	// "match_service_account": {
	// "source": "account",
	// "target": "data.project_id"
	// }
	// },
	// "options": {
	// "stop_processing": true
	// },
	// "tags": {
	// "a": "b",
	// "b": "c"
	// },
	// "collector_id": "collector-2e39891cbbb5",
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2022-07-19T10:13:28.335Z"
	// }
	ChangeOrder(context.Context, *ChangeCollectorRuleOrderRequest) (*CollectorRuleInfo, error)
	// desc: Deletes a specific CollectorRule. You must specify the `collector_rule_id` of the CollectorRule to delete.
	// request_example: >-
	// {
	// "collector_rule_id": "collector-rule-c8055231e212",
	// }
	Delete(context.Context, *CollectorRuleRequest) (*empty.Empty, error)
	// desc: Gets a specific CollectorRule. Prints detailed information about the CollectorRule, including  `conditions_policy` and conditions applied to Collector.
	// request_example: >-
	// {
	// "collector_rule_id": "collector-rule-c8055231e212",
	// }
	// response_example: >-
	// {
	// "collector_rule_id": "collector-rule-c8055231e212",
	// "name": "match_service_account",
	// "order": 1,
	// "conditions_policy": "ALWAYS",
	// "actions": {
	// "match_service_account": {
	// "source": "account",
	// "target": "data.project_id"
	// }
	// },
	// "options": {
	// "stop_processing": true
	// },
	// "tags": {},
	// "collector_id": "collector-2e39891cbbb5",
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2022-05-25T16:01:51.858Z"
	// }
	Get(context.Context, *GetCollectorRuleRequest) (*CollectorRuleInfo, error)
	// desc: Gets a list of all CollectorRules. You can use a query to get a filtered list of CollectorRules.
	// request_example: >-
	// {
	// "query": {}
	// }
	// response_example: >-
	// {
	// "results": [
	// {
	// "collector_rule_id": "collector-rule-c8055231e212",
	// "name": "match_service_account",
	// "order": 1,
	// "conditions_policy": "ALWAYS",
	// "actions": {
	// "match_service_account": {
	// "source": "account",
	// "target": "data.project_id"
	// }
	// },
	// "options": {
	// "stop_processing": true
	// },
	// "tags": {},
	// "collector_id": "collector-2e39891cbbb5",
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2022-05-25T16:01:51.858Z"
	// },
	// {
	// "collector_rule_id": "collector-rule-t3345231e167",
	// "name": "match_service_account",
	// "order": 1,
	// "conditions_policy": "ALWAYS",
	// "actions": {
	// "match_service_account": {
	// "source": "account",
	// "target": "data.account_id"
	// }
	// },
	// "options": {
	// "stop_processing": true
	// },
	// "tags": {},
	// "collector_id": "collector-7163022d49a1",
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2022-06-03T16:00:54.099Z"
	// }
	// ],
	// "total_count": 2
	// }
	List(context.Context, *CollectorRuleQuery) (*CollectorRulesInfo, error)
	Stat(context.Context, *CollectorRuleStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedCollectorRuleServer()
}

// UnimplementedCollectorRuleServer must be embedded to have forward compatible implementations.
type UnimplementedCollectorRuleServer struct {
}

func (UnimplementedCollectorRuleServer) Create(context.Context, *CreateCollectorRuleRequest) (*CollectorRuleInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCollectorRuleServer) Update(context.Context, *UpdateCollectorRuleRequest) (*CollectorRuleInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCollectorRuleServer) ChangeOrder(context.Context, *ChangeCollectorRuleOrderRequest) (*CollectorRuleInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeOrder not implemented")
}
func (UnimplementedCollectorRuleServer) Delete(context.Context, *CollectorRuleRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCollectorRuleServer) Get(context.Context, *GetCollectorRuleRequest) (*CollectorRuleInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCollectorRuleServer) List(context.Context, *CollectorRuleQuery) (*CollectorRulesInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCollectorRuleServer) Stat(context.Context, *CollectorRuleStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedCollectorRuleServer) mustEmbedUnimplementedCollectorRuleServer() {}

// UnsafeCollectorRuleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CollectorRuleServer will
// result in compilation errors.
type UnsafeCollectorRuleServer interface {
	mustEmbedUnimplementedCollectorRuleServer()
}

func RegisterCollectorRuleServer(s grpc.ServiceRegistrar, srv CollectorRuleServer) {
	s.RegisterService(&CollectorRule_ServiceDesc, srv)
}

func _CollectorRule_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCollectorRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectorRuleServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectorRule_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectorRuleServer).Create(ctx, req.(*CreateCollectorRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectorRule_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCollectorRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectorRuleServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectorRule_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectorRuleServer).Update(ctx, req.(*UpdateCollectorRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectorRule_ChangeOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeCollectorRuleOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectorRuleServer).ChangeOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectorRule_ChangeOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectorRuleServer).ChangeOrder(ctx, req.(*ChangeCollectorRuleOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectorRule_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectorRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectorRuleServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectorRule_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectorRuleServer).Delete(ctx, req.(*CollectorRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectorRule_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCollectorRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectorRuleServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectorRule_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectorRuleServer).Get(ctx, req.(*GetCollectorRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectorRule_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectorRuleQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectorRuleServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectorRule_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectorRuleServer).List(ctx, req.(*CollectorRuleQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectorRule_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectorRuleStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectorRuleServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectorRule_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectorRuleServer).Stat(ctx, req.(*CollectorRuleStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// CollectorRule_ServiceDesc is the grpc.ServiceDesc for CollectorRule service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CollectorRule_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.inventory.v1.CollectorRule",
	HandlerType: (*CollectorRuleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _CollectorRule_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _CollectorRule_Update_Handler,
		},
		{
			MethodName: "change_order",
			Handler:    _CollectorRule_ChangeOrder_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _CollectorRule_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _CollectorRule_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _CollectorRule_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _CollectorRule_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/inventory/v1/collector_rule.proto",
}
