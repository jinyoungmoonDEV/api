// An EventRule is a rule to transform the request data when an Event is generated.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: spaceone/api/monitoring/v1/event_rule.proto

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
	EventRule_Create_FullMethodName      = "/spaceone.api.monitoring.v1.EventRule/create"
	EventRule_Update_FullMethodName      = "/spaceone.api.monitoring.v1.EventRule/update"
	EventRule_ChangeOrder_FullMethodName = "/spaceone.api.monitoring.v1.EventRule/change_order"
	EventRule_Delete_FullMethodName      = "/spaceone.api.monitoring.v1.EventRule/delete"
	EventRule_Get_FullMethodName         = "/spaceone.api.monitoring.v1.EventRule/get"
	EventRule_List_FullMethodName        = "/spaceone.api.monitoring.v1.EventRule/list"
	EventRule_Stat_FullMethodName        = "/spaceone.api.monitoring.v1.EventRule/stat"
)

// EventRuleClient is the client API for EventRule service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventRuleClient interface {
	// Creates a new EventRule. You can filter the Events to apply the EventRule by setting the input parameter `Conditions`. The method can change the Events' assignee or Project.
	Create(ctx context.Context, in *CreateEventRuleRequest, opts ...grpc.CallOption) (*EventRuleInfo, error)
	// Changes a priority order between EventRules to apply. EventRules are filtered by the order configured.
	Update(ctx context.Context, in *UpdateEventRuleRequest, opts ...grpc.CallOption) (*EventRuleInfo, error)
	// Updates a specific EventRule. You can make changes in EventRule settings.
	ChangeOrder(ctx context.Context, in *ChangeEventRuleOrderRequest, opts ...grpc.CallOption) (*EventRuleInfo, error)
	// Deletes a specific EventRule. You must assign an EventRule resource to delete by specifying `event_rule_id`.
	Delete(ctx context.Context, in *EventRuleRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Gets a specific EventRule. Prints detailed information about the EventRule.
	Get(ctx context.Context, in *EventRuleRequest, opts ...grpc.CallOption) (*EventRuleInfo, error)
	// Gets a list of all EventRules. You can use a query to get a filtered list of EventRules. For example, you can adjust the scope of the list to a certain Project or Domain.
	List(ctx context.Context, in *EventRuleQuery, opts ...grpc.CallOption) (*EventRulesInfo, error)
	Stat(ctx context.Context, in *EventRuleStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type eventRuleClient struct {
	cc grpc.ClientConnInterface
}

func NewEventRuleClient(cc grpc.ClientConnInterface) EventRuleClient {
	return &eventRuleClient{cc}
}

func (c *eventRuleClient) Create(ctx context.Context, in *CreateEventRuleRequest, opts ...grpc.CallOption) (*EventRuleInfo, error) {
	out := new(EventRuleInfo)
	err := c.cc.Invoke(ctx, EventRule_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventRuleClient) Update(ctx context.Context, in *UpdateEventRuleRequest, opts ...grpc.CallOption) (*EventRuleInfo, error) {
	out := new(EventRuleInfo)
	err := c.cc.Invoke(ctx, EventRule_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventRuleClient) ChangeOrder(ctx context.Context, in *ChangeEventRuleOrderRequest, opts ...grpc.CallOption) (*EventRuleInfo, error) {
	out := new(EventRuleInfo)
	err := c.cc.Invoke(ctx, EventRule_ChangeOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventRuleClient) Delete(ctx context.Context, in *EventRuleRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, EventRule_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventRuleClient) Get(ctx context.Context, in *EventRuleRequest, opts ...grpc.CallOption) (*EventRuleInfo, error) {
	out := new(EventRuleInfo)
	err := c.cc.Invoke(ctx, EventRule_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventRuleClient) List(ctx context.Context, in *EventRuleQuery, opts ...grpc.CallOption) (*EventRulesInfo, error) {
	out := new(EventRulesInfo)
	err := c.cc.Invoke(ctx, EventRule_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventRuleClient) Stat(ctx context.Context, in *EventRuleStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, EventRule_Stat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventRuleServer is the server API for EventRule service.
// All implementations must embed UnimplementedEventRuleServer
// for forward compatibility
type EventRuleServer interface {
	// Creates a new EventRule. You can filter the Events to apply the EventRule by setting the input parameter `Conditions`. The method can change the Events' assignee or Project.
	Create(context.Context, *CreateEventRuleRequest) (*EventRuleInfo, error)
	// Changes a priority order between EventRules to apply. EventRules are filtered by the order configured.
	Update(context.Context, *UpdateEventRuleRequest) (*EventRuleInfo, error)
	// Updates a specific EventRule. You can make changes in EventRule settings.
	ChangeOrder(context.Context, *ChangeEventRuleOrderRequest) (*EventRuleInfo, error)
	// Deletes a specific EventRule. You must assign an EventRule resource to delete by specifying `event_rule_id`.
	Delete(context.Context, *EventRuleRequest) (*empty.Empty, error)
	// Gets a specific EventRule. Prints detailed information about the EventRule.
	Get(context.Context, *EventRuleRequest) (*EventRuleInfo, error)
	// Gets a list of all EventRules. You can use a query to get a filtered list of EventRules. For example, you can adjust the scope of the list to a certain Project or Domain.
	List(context.Context, *EventRuleQuery) (*EventRulesInfo, error)
	Stat(context.Context, *EventRuleStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedEventRuleServer()
}

// UnimplementedEventRuleServer must be embedded to have forward compatible implementations.
type UnimplementedEventRuleServer struct {
}

func (UnimplementedEventRuleServer) Create(context.Context, *CreateEventRuleRequest) (*EventRuleInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedEventRuleServer) Update(context.Context, *UpdateEventRuleRequest) (*EventRuleInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedEventRuleServer) ChangeOrder(context.Context, *ChangeEventRuleOrderRequest) (*EventRuleInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeOrder not implemented")
}
func (UnimplementedEventRuleServer) Delete(context.Context, *EventRuleRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedEventRuleServer) Get(context.Context, *EventRuleRequest) (*EventRuleInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedEventRuleServer) List(context.Context, *EventRuleQuery) (*EventRulesInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedEventRuleServer) Stat(context.Context, *EventRuleStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedEventRuleServer) mustEmbedUnimplementedEventRuleServer() {}

// UnsafeEventRuleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventRuleServer will
// result in compilation errors.
type UnsafeEventRuleServer interface {
	mustEmbedUnimplementedEventRuleServer()
}

func RegisterEventRuleServer(s grpc.ServiceRegistrar, srv EventRuleServer) {
	s.RegisterService(&EventRule_ServiceDesc, srv)
}

func _EventRule_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventRuleServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventRule_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventRuleServer).Create(ctx, req.(*CreateEventRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventRule_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEventRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventRuleServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventRule_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventRuleServer).Update(ctx, req.(*UpdateEventRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventRule_ChangeOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeEventRuleOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventRuleServer).ChangeOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventRule_ChangeOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventRuleServer).ChangeOrder(ctx, req.(*ChangeEventRuleOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventRule_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventRuleServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventRule_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventRuleServer).Delete(ctx, req.(*EventRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventRule_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventRuleServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventRule_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventRuleServer).Get(ctx, req.(*EventRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventRule_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRuleQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventRuleServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventRule_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventRuleServer).List(ctx, req.(*EventRuleQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventRule_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRuleStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventRuleServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventRule_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventRuleServer).Stat(ctx, req.(*EventRuleStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// EventRule_ServiceDesc is the grpc.ServiceDesc for EventRule service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventRule_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.monitoring.v1.EventRule",
	HandlerType: (*EventRuleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _EventRule_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _EventRule_Update_Handler,
		},
		{
			MethodName: "change_order",
			Handler:    _EventRule_ChangeOrder_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _EventRule_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _EventRule_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _EventRule_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _EventRule_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/monitoring/v1/event_rule.proto",
}
