//
//desc: A Protocol defines the method to use when dispatching Notifications via a channel.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: spaceone/api/notification/v1/protocol.proto

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
	Protocol_Create_FullMethodName       = "/spaceone.api.notification.v1.Protocol/create"
	Protocol_Update_FullMethodName       = "/spaceone.api.notification.v1.Protocol/update"
	Protocol_UpdatePlugin_FullMethodName = "/spaceone.api.notification.v1.Protocol/update_plugin"
	Protocol_Enable_FullMethodName       = "/spaceone.api.notification.v1.Protocol/enable"
	Protocol_Disable_FullMethodName      = "/spaceone.api.notification.v1.Protocol/disable"
	Protocol_Delete_FullMethodName       = "/spaceone.api.notification.v1.Protocol/delete"
	Protocol_Get_FullMethodName          = "/spaceone.api.notification.v1.Protocol/get"
	Protocol_List_FullMethodName         = "/spaceone.api.notification.v1.Protocol/list"
	Protocol_Stat_FullMethodName         = "/spaceone.api.notification.v1.Protocol/stat"
)

// ProtocolClient is the client API for Protocol service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProtocolClient interface {
	// desc: Creates a new Protocol. When creating a protocol, you must specify the plugins provided from the repository, and you must also set the credentials to be set in the plugin if necessary.
	// request_example: >-
	// {
	// "name": "Email",
	// "plugin_info": {
	// "plugin_id": "plugin-email-noti-protocol",
	// "version": "1.0.1",
	// "options": {},
	// "secret_id": "secret-123546789012",
	// "metadata": {
	// "data": {
	// "schema": {
	// "properties": {
	// "email": {
	// "pattern": "^[\\W]*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4}[\\W]*,{1}[\\W]*)*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4})[\\W]*$",
	// "examples": [
	// "user1@test.com, user2@test.com"
	// ],
	// "minLength": 10.0,
	// "description": "Email address to receive notifications",
	// "type": "string",
	// "title": "Email Address"
	// }
	// },
	// "required": [
	// "email"
	// ],
	// "type": "object"
	// }
	// },
	// "data_type": "PLAIN_TEXT"
	// },
	// "upgrade_mode": "AUTO"
	// },
	// "tags": {}
	// }
	// response_example: >-
	// {
	// "protocol_id": "protocol-123546789012",
	// "name": "Email",
	// "state": "ENABLED",
	// "protocol_type": "EXTERNAL",
	// "capability": {
	// "supported_schema": [
	// "email_smtp"
	// ]
	// },
	// "plugin_info": {
	// "plugin_id": "plugin-email-noti-protocol",
	// "version": "1.0.1",
	// "options": {},
	// "secret_id": "secret-123546789012",
	// "metadata": {
	// "data": {
	// "schema": {
	// "properties": {
	// "email": {
	// "pattern": "^[\\W]*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4}[\\W]*,{1}[\\W]*)*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4})[\\W]*$",
	// "examples": [
	// "user1@test.com, user2@test.com"
	// ],
	// "minLength": 10.0,
	// "description": "Email address to receive notifications",
	// "type": "string",
	// "title": "Email Address"
	// }
	// },
	// "required": [
	// "email"
	// ],
	// "type": "object"
	// }
	// },
	// "data_type": "PLAIN_TEXT"
	// },
	// "upgrade_mode": "AUTO"
	// },
	// "tags": {},
	// "domain_id": "domain-123546789012",
	// "created_at": "2022-01-01T07:55:57.043Z"
	// }
	Create(ctx context.Context, in *CreateProtocolRequest, opts ...grpc.CallOption) (*ProtocolInfo, error)
	// desc: Updates a specific Protocol. The method `update` can update the name and tags only. If you want to update the plugin version or options, you can use `update_plugin` method.
	// request_example: >-
	// {
	// "protocol_id": "protocol-123456789012",
	// "name": "Email-test",
	// "tags": {
	// "type": "test"
	// }
	// }
	// response_example: >-
	// {
	// "protocol_id": "protocol-123546789012",
	// "name": "Email-test",
	// "state": "ENABLED",
	// "protocol_type": "EXTERNAL",
	// "capability": {
	// "supported_schema": [
	// "email_smtp"
	// ]
	// },
	// "plugin_info": {
	// "plugin_id": "plugin-email-noti-protocol",
	// "version": "1.0.1",
	// "options": {},
	// "secret_id": "secret-123546789012",
	// "metadata": {
	// "data": {
	// "schema": {
	// "properties": {
	// "email": {
	// "pattern": "^[\\W]*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4}[\\W]*,{1}[\\W]*)*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4})[\\W]*$",
	// "examples": [
	// "user1@test.com, user2@test.com"
	// ],
	// "minLength": 10.0,
	// "description": "Email address to receive notifications",
	// "type": "string",
	// "title": "Email Address"
	// }
	// },
	// "required": [
	// "email"
	// ],
	// "type": "object"
	// }
	// },
	// "data_type": "PLAIN_TEXT"
	// },
	// "upgrade_mode": "AUTO"
	// },
	// "tags": {
	// "type": "test"
	// },
	// "domain_id": "domain-123546789012",
	// "created_at": "2022-01-01T07:55:57.043Z"
	// }
	Update(ctx context.Context, in *UpdateProtocolRequest, opts ...grpc.CallOption) (*ProtocolInfo, error)
	// desc: Updates a plugin for a Protocol. It is usually used when redeploying a plugin to a new version.
	// request_example: >-
	// {
	// "protocol_id": "protocol-123456789012",
	// "version": "1.0.2",
	// "options": {}
	// }
	// response_example: >-
	// {
	// "protocol_id": "protocol-123546789012",
	// "name": "Email-test",
	// "state": "ENABLED",
	// "protocol_type": "EXTERNAL",
	// "capability": {
	// "supported_schema": [
	// "email_smtp"
	// ]
	// },
	// "plugin_info": {
	// "plugin_id": "plugin-email-noti-protocol",
	// "version": "1.0.2",
	// "options": {},
	// "secret_id": "secret-123546789012",
	// "metadata": {
	// "data": {
	// "schema": {
	// "properties": {
	// "email": {
	// "pattern": "^[\\W]*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4}[\\W]*,{1}[\\W]*)*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4})[\\W]*$",
	// "examples": [
	// "user1@test.com, user2@test.com"
	// ],
	// "minLength": 10.0,
	// "description": "Email address to receive notifications",
	// "type": "string",
	// "title": "Email Address"
	// }
	// },
	// "required": [
	// "email"
	// ],
	// "type": "object"
	// }
	// },
	// "data_type": "PLAIN_TEXT"
	// },
	// "upgrade_mode": "AUTO"
	// },
	// "tags": {
	// "type": "test"
	// },
	// "domain_id": "domain-123546789012",
	// "created_at": "2022-01-01T07:55:57.043Z"
	// }
	UpdatePlugin(ctx context.Context, in *UpdateProtocolPluginRequest, opts ...grpc.CallOption) (*ProtocolInfo, error)
	// desc: Enables a specific Protocol. If the Protocol is enabled, the Protocol can be used and the Notification can be dispatched.
	// request_example: >-
	// {
	// "protocol_id": "protocol-123456789012"
	// }
	// response_example: >-
	// {
	// "protocol_id": "protocol-123546789012",
	// "name": "Email-test",
	// "state": "ENABLED",
	// "protocol_type": "EXTERNAL",
	// "capability": {
	// "supported_schema": [
	// "email_smtp"
	// ]
	// },
	// "plugin_info": {
	// "plugin_id": "plugin-email-noti-protocol",
	// "version": "1.0.2",
	// "options": {},
	// "secret_id": "secret-123546789012",
	// "metadata": {
	// "data": {
	// "schema": {
	// "properties": {
	// "email": {
	// "pattern": "^[\\W]*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4}[\\W]*,{1}[\\W]*)*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4})[\\W]*$",
	// "examples": [
	// "user1@test.com, user2@test.com"
	// ],
	// "minLength": 10.0,
	// "description": "Email address to receive notifications",
	// "type": "string",
	// "title": "Email Address"
	// }
	// },
	// "required": [
	// "email"
	// ],
	// "type": "object"
	// }
	// },
	// "data_type": "PLAIN_TEXT"
	// },
	// "upgrade_mode": "AUTO"
	// },
	// "tags": {
	// "type": "test"
	// },
	// "domain_id": "domain-123546789012",
	// "created_at": "2022-01-01T07:55:57.043Z"
	// }
	Enable(ctx context.Context, in *ProtocolRequest, opts ...grpc.CallOption) (*ProtocolInfo, error)
	// desc: Disables a specific Protocol. If a Protocol is disabled, the Notification will not be dispatched, even if it is created.
	// request_example: >-
	// {
	// "protocol_id": "protocol-123456789012"
	// }
	// response_example: >-
	// {
	// "protocol_id": "protocol-123546789012",
	// "name": "Email-test",
	// "state": "DISABLED",
	// "protocol_type": "EXTERNAL",
	// "capability": {
	// "supported_schema": [
	// "email_smtp"
	// ]
	// },
	// "plugin_info": {
	// "plugin_id": "plugin-email-noti-protocol",
	// "version": "1.0.2",
	// "options": {},
	// "secret_id": "secret-123546789012",
	// "metadata": {
	// "data": {
	// "schema": {
	// "properties": {
	// "email": {
	// "pattern": "^[\\W]*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4}[\\W]*,{1}[\\W]*)*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4})[\\W]*$",
	// "examples": [
	// "user1@test.com, user2@test.com"
	// ],
	// "minLength": 10.0,
	// "description": "Email address to receive notifications",
	// "type": "string",
	// "title": "Email Address"
	// }
	// },
	// "required": [
	// "email"
	// ],
	// "type": "object"
	// }
	// },
	// "data_type": "PLAIN_TEXT"
	// },
	// "upgrade_mode": "AUTO"
	// },
	// "tags": {
	// "type": "test"
	// },
	// "domain_id": "domain-123546789012",
	// "created_at": "2022-01-01T07:55:57.043Z"
	// }
	Disable(ctx context.Context, in *ProtocolRequest, opts ...grpc.CallOption) (*ProtocolInfo, error)
	// desc: Deletes a specific Protocol. If there exists a channel using the Protocol, it cannot be deleted.
	// request_example: >-
	// {
	// "protocol_id": "protocol-123456789012"
	// }
	Delete(ctx context.Context, in *ProtocolRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// desc: Gets a specific Protocol. Prints detailed information about the Protocol.
	// request_example: >-
	// {
	// "protocol_id": "protocol-123456789012"
	// }
	// response_example: >-
	// {
	// "protocol_id": "protocol-123546789012",
	// "name": "Email-test",
	// "state": "DISABLED",
	// "protocol_type": "EXTERNAL",
	// "capability": {
	// "supported_schema": [
	// "email_smtp"
	// ]
	// },
	// "plugin_info": {
	// "plugin_id": "plugin-email-noti-protocol",
	// "version": "1.0.2",
	// "options": {},
	// "secret_id": "secret-123546789012",
	// "metadata": {
	// "data": {
	// "schema": {
	// "properties": {
	// "email": {
	// "pattern": "^[\\W]*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4}[\\W]*,{1}[\\W]*)*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4})[\\W]*$",
	// "examples": [
	// "user1@test.com, user2@test.com"
	// ],
	// "minLength": 10.0,
	// "description": "Email address to receive notifications",
	// "type": "string",
	// "title": "Email Address"
	// }
	// },
	// "required": [
	// "email"
	// ],
	// "type": "object"
	// }
	// },
	// "data_type": "PLAIN_TEXT"
	// },
	// "upgrade_mode": "AUTO"
	// },
	// "tags": {
	// "type": "test"
	// },
	// "domain_id": "domain-123546789012",
	// "created_at": "2022-01-01T07:55:57.043Z"
	// }
	Get(ctx context.Context, in *GetProtocolRequest, opts ...grpc.CallOption) (*ProtocolInfo, error)
	// desc: Gets a list of Protocols. You can use a query to get a filtered list of Protocols.
	// request_example: >-
	// {
	// "query": {}
	// }
	// response_example: >-
	// {
	// "results":[
	// {
	// "protocol_id":"protocol-123456789012",
	// "name":"Email",
	// "state":"ENABLED",
	// "protocol_type":"EXTERNAL",
	// "capability":{
	// "supported_schema":[
	// "email_smtp"
	// ]
	// },
	// "plugin_info":{
	// "plugin_id":"plugin-email-noti-protocol",
	// "version":"1.0.1",
	// "options":{
	//
	// },
	// "secret_id":"secret-123456789012",
	// "metadata":{
	// "data_type":"PLAIN_TEXT",
	// "data":{
	// "schema":{
	// "properties":{
	// "email":{
	// "pattern":"^[\\W]*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4}[\\W]*,{1}[\\W]*)*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4})[\\W]*$",
	// "examples":[
	// "user1@test.com, user2@test.com"
	// ],
	// "minLength":10.0,
	// "description":"Email address to receive notifications",
	// "type":"string",
	// "title":"Email Address"
	// }
	// },
	// "required":[
	// "email"
	// ],
	// "type":"object"
	// }
	// }
	// },
	// "upgrade_mode":"AUTO"
	// },
	// "tags":{
	//
	// },
	// "domain_id":"domain-123456789012",
	// "created_at":"2022-01-01T07:55:57.043Z"
	// }
	// ],
	// "total_count":1
	// }
	List(ctx context.Context, in *ProtocolQuery, opts ...grpc.CallOption) (*ProtocolsInfo, error)
	Stat(ctx context.Context, in *ProtocolStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type protocolClient struct {
	cc grpc.ClientConnInterface
}

func NewProtocolClient(cc grpc.ClientConnInterface) ProtocolClient {
	return &protocolClient{cc}
}

func (c *protocolClient) Create(ctx context.Context, in *CreateProtocolRequest, opts ...grpc.CallOption) (*ProtocolInfo, error) {
	out := new(ProtocolInfo)
	err := c.cc.Invoke(ctx, Protocol_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protocolClient) Update(ctx context.Context, in *UpdateProtocolRequest, opts ...grpc.CallOption) (*ProtocolInfo, error) {
	out := new(ProtocolInfo)
	err := c.cc.Invoke(ctx, Protocol_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protocolClient) UpdatePlugin(ctx context.Context, in *UpdateProtocolPluginRequest, opts ...grpc.CallOption) (*ProtocolInfo, error) {
	out := new(ProtocolInfo)
	err := c.cc.Invoke(ctx, Protocol_UpdatePlugin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protocolClient) Enable(ctx context.Context, in *ProtocolRequest, opts ...grpc.CallOption) (*ProtocolInfo, error) {
	out := new(ProtocolInfo)
	err := c.cc.Invoke(ctx, Protocol_Enable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protocolClient) Disable(ctx context.Context, in *ProtocolRequest, opts ...grpc.CallOption) (*ProtocolInfo, error) {
	out := new(ProtocolInfo)
	err := c.cc.Invoke(ctx, Protocol_Disable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protocolClient) Delete(ctx context.Context, in *ProtocolRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Protocol_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protocolClient) Get(ctx context.Context, in *GetProtocolRequest, opts ...grpc.CallOption) (*ProtocolInfo, error) {
	out := new(ProtocolInfo)
	err := c.cc.Invoke(ctx, Protocol_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protocolClient) List(ctx context.Context, in *ProtocolQuery, opts ...grpc.CallOption) (*ProtocolsInfo, error) {
	out := new(ProtocolsInfo)
	err := c.cc.Invoke(ctx, Protocol_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protocolClient) Stat(ctx context.Context, in *ProtocolStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, Protocol_Stat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProtocolServer is the server API for Protocol service.
// All implementations must embed UnimplementedProtocolServer
// for forward compatibility
type ProtocolServer interface {
	// desc: Creates a new Protocol. When creating a protocol, you must specify the plugins provided from the repository, and you must also set the credentials to be set in the plugin if necessary.
	// request_example: >-
	// {
	// "name": "Email",
	// "plugin_info": {
	// "plugin_id": "plugin-email-noti-protocol",
	// "version": "1.0.1",
	// "options": {},
	// "secret_id": "secret-123546789012",
	// "metadata": {
	// "data": {
	// "schema": {
	// "properties": {
	// "email": {
	// "pattern": "^[\\W]*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4}[\\W]*,{1}[\\W]*)*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4})[\\W]*$",
	// "examples": [
	// "user1@test.com, user2@test.com"
	// ],
	// "minLength": 10.0,
	// "description": "Email address to receive notifications",
	// "type": "string",
	// "title": "Email Address"
	// }
	// },
	// "required": [
	// "email"
	// ],
	// "type": "object"
	// }
	// },
	// "data_type": "PLAIN_TEXT"
	// },
	// "upgrade_mode": "AUTO"
	// },
	// "tags": {}
	// }
	// response_example: >-
	// {
	// "protocol_id": "protocol-123546789012",
	// "name": "Email",
	// "state": "ENABLED",
	// "protocol_type": "EXTERNAL",
	// "capability": {
	// "supported_schema": [
	// "email_smtp"
	// ]
	// },
	// "plugin_info": {
	// "plugin_id": "plugin-email-noti-protocol",
	// "version": "1.0.1",
	// "options": {},
	// "secret_id": "secret-123546789012",
	// "metadata": {
	// "data": {
	// "schema": {
	// "properties": {
	// "email": {
	// "pattern": "^[\\W]*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4}[\\W]*,{1}[\\W]*)*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4})[\\W]*$",
	// "examples": [
	// "user1@test.com, user2@test.com"
	// ],
	// "minLength": 10.0,
	// "description": "Email address to receive notifications",
	// "type": "string",
	// "title": "Email Address"
	// }
	// },
	// "required": [
	// "email"
	// ],
	// "type": "object"
	// }
	// },
	// "data_type": "PLAIN_TEXT"
	// },
	// "upgrade_mode": "AUTO"
	// },
	// "tags": {},
	// "domain_id": "domain-123546789012",
	// "created_at": "2022-01-01T07:55:57.043Z"
	// }
	Create(context.Context, *CreateProtocolRequest) (*ProtocolInfo, error)
	// desc: Updates a specific Protocol. The method `update` can update the name and tags only. If you want to update the plugin version or options, you can use `update_plugin` method.
	// request_example: >-
	// {
	// "protocol_id": "protocol-123456789012",
	// "name": "Email-test",
	// "tags": {
	// "type": "test"
	// }
	// }
	// response_example: >-
	// {
	// "protocol_id": "protocol-123546789012",
	// "name": "Email-test",
	// "state": "ENABLED",
	// "protocol_type": "EXTERNAL",
	// "capability": {
	// "supported_schema": [
	// "email_smtp"
	// ]
	// },
	// "plugin_info": {
	// "plugin_id": "plugin-email-noti-protocol",
	// "version": "1.0.1",
	// "options": {},
	// "secret_id": "secret-123546789012",
	// "metadata": {
	// "data": {
	// "schema": {
	// "properties": {
	// "email": {
	// "pattern": "^[\\W]*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4}[\\W]*,{1}[\\W]*)*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4})[\\W]*$",
	// "examples": [
	// "user1@test.com, user2@test.com"
	// ],
	// "minLength": 10.0,
	// "description": "Email address to receive notifications",
	// "type": "string",
	// "title": "Email Address"
	// }
	// },
	// "required": [
	// "email"
	// ],
	// "type": "object"
	// }
	// },
	// "data_type": "PLAIN_TEXT"
	// },
	// "upgrade_mode": "AUTO"
	// },
	// "tags": {
	// "type": "test"
	// },
	// "domain_id": "domain-123546789012",
	// "created_at": "2022-01-01T07:55:57.043Z"
	// }
	Update(context.Context, *UpdateProtocolRequest) (*ProtocolInfo, error)
	// desc: Updates a plugin for a Protocol. It is usually used when redeploying a plugin to a new version.
	// request_example: >-
	// {
	// "protocol_id": "protocol-123456789012",
	// "version": "1.0.2",
	// "options": {}
	// }
	// response_example: >-
	// {
	// "protocol_id": "protocol-123546789012",
	// "name": "Email-test",
	// "state": "ENABLED",
	// "protocol_type": "EXTERNAL",
	// "capability": {
	// "supported_schema": [
	// "email_smtp"
	// ]
	// },
	// "plugin_info": {
	// "plugin_id": "plugin-email-noti-protocol",
	// "version": "1.0.2",
	// "options": {},
	// "secret_id": "secret-123546789012",
	// "metadata": {
	// "data": {
	// "schema": {
	// "properties": {
	// "email": {
	// "pattern": "^[\\W]*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4}[\\W]*,{1}[\\W]*)*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4})[\\W]*$",
	// "examples": [
	// "user1@test.com, user2@test.com"
	// ],
	// "minLength": 10.0,
	// "description": "Email address to receive notifications",
	// "type": "string",
	// "title": "Email Address"
	// }
	// },
	// "required": [
	// "email"
	// ],
	// "type": "object"
	// }
	// },
	// "data_type": "PLAIN_TEXT"
	// },
	// "upgrade_mode": "AUTO"
	// },
	// "tags": {
	// "type": "test"
	// },
	// "domain_id": "domain-123546789012",
	// "created_at": "2022-01-01T07:55:57.043Z"
	// }
	UpdatePlugin(context.Context, *UpdateProtocolPluginRequest) (*ProtocolInfo, error)
	// desc: Enables a specific Protocol. If the Protocol is enabled, the Protocol can be used and the Notification can be dispatched.
	// request_example: >-
	// {
	// "protocol_id": "protocol-123456789012"
	// }
	// response_example: >-
	// {
	// "protocol_id": "protocol-123546789012",
	// "name": "Email-test",
	// "state": "ENABLED",
	// "protocol_type": "EXTERNAL",
	// "capability": {
	// "supported_schema": [
	// "email_smtp"
	// ]
	// },
	// "plugin_info": {
	// "plugin_id": "plugin-email-noti-protocol",
	// "version": "1.0.2",
	// "options": {},
	// "secret_id": "secret-123546789012",
	// "metadata": {
	// "data": {
	// "schema": {
	// "properties": {
	// "email": {
	// "pattern": "^[\\W]*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4}[\\W]*,{1}[\\W]*)*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4})[\\W]*$",
	// "examples": [
	// "user1@test.com, user2@test.com"
	// ],
	// "minLength": 10.0,
	// "description": "Email address to receive notifications",
	// "type": "string",
	// "title": "Email Address"
	// }
	// },
	// "required": [
	// "email"
	// ],
	// "type": "object"
	// }
	// },
	// "data_type": "PLAIN_TEXT"
	// },
	// "upgrade_mode": "AUTO"
	// },
	// "tags": {
	// "type": "test"
	// },
	// "domain_id": "domain-123546789012",
	// "created_at": "2022-01-01T07:55:57.043Z"
	// }
	Enable(context.Context, *ProtocolRequest) (*ProtocolInfo, error)
	// desc: Disables a specific Protocol. If a Protocol is disabled, the Notification will not be dispatched, even if it is created.
	// request_example: >-
	// {
	// "protocol_id": "protocol-123456789012"
	// }
	// response_example: >-
	// {
	// "protocol_id": "protocol-123546789012",
	// "name": "Email-test",
	// "state": "DISABLED",
	// "protocol_type": "EXTERNAL",
	// "capability": {
	// "supported_schema": [
	// "email_smtp"
	// ]
	// },
	// "plugin_info": {
	// "plugin_id": "plugin-email-noti-protocol",
	// "version": "1.0.2",
	// "options": {},
	// "secret_id": "secret-123546789012",
	// "metadata": {
	// "data": {
	// "schema": {
	// "properties": {
	// "email": {
	// "pattern": "^[\\W]*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4}[\\W]*,{1}[\\W]*)*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4})[\\W]*$",
	// "examples": [
	// "user1@test.com, user2@test.com"
	// ],
	// "minLength": 10.0,
	// "description": "Email address to receive notifications",
	// "type": "string",
	// "title": "Email Address"
	// }
	// },
	// "required": [
	// "email"
	// ],
	// "type": "object"
	// }
	// },
	// "data_type": "PLAIN_TEXT"
	// },
	// "upgrade_mode": "AUTO"
	// },
	// "tags": {
	// "type": "test"
	// },
	// "domain_id": "domain-123546789012",
	// "created_at": "2022-01-01T07:55:57.043Z"
	// }
	Disable(context.Context, *ProtocolRequest) (*ProtocolInfo, error)
	// desc: Deletes a specific Protocol. If there exists a channel using the Protocol, it cannot be deleted.
	// request_example: >-
	// {
	// "protocol_id": "protocol-123456789012"
	// }
	Delete(context.Context, *ProtocolRequest) (*empty.Empty, error)
	// desc: Gets a specific Protocol. Prints detailed information about the Protocol.
	// request_example: >-
	// {
	// "protocol_id": "protocol-123456789012"
	// }
	// response_example: >-
	// {
	// "protocol_id": "protocol-123546789012",
	// "name": "Email-test",
	// "state": "DISABLED",
	// "protocol_type": "EXTERNAL",
	// "capability": {
	// "supported_schema": [
	// "email_smtp"
	// ]
	// },
	// "plugin_info": {
	// "plugin_id": "plugin-email-noti-protocol",
	// "version": "1.0.2",
	// "options": {},
	// "secret_id": "secret-123546789012",
	// "metadata": {
	// "data": {
	// "schema": {
	// "properties": {
	// "email": {
	// "pattern": "^[\\W]*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4}[\\W]*,{1}[\\W]*)*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4})[\\W]*$",
	// "examples": [
	// "user1@test.com, user2@test.com"
	// ],
	// "minLength": 10.0,
	// "description": "Email address to receive notifications",
	// "type": "string",
	// "title": "Email Address"
	// }
	// },
	// "required": [
	// "email"
	// ],
	// "type": "object"
	// }
	// },
	// "data_type": "PLAIN_TEXT"
	// },
	// "upgrade_mode": "AUTO"
	// },
	// "tags": {
	// "type": "test"
	// },
	// "domain_id": "domain-123546789012",
	// "created_at": "2022-01-01T07:55:57.043Z"
	// }
	Get(context.Context, *GetProtocolRequest) (*ProtocolInfo, error)
	// desc: Gets a list of Protocols. You can use a query to get a filtered list of Protocols.
	// request_example: >-
	// {
	// "query": {}
	// }
	// response_example: >-
	// {
	// "results":[
	// {
	// "protocol_id":"protocol-123456789012",
	// "name":"Email",
	// "state":"ENABLED",
	// "protocol_type":"EXTERNAL",
	// "capability":{
	// "supported_schema":[
	// "email_smtp"
	// ]
	// },
	// "plugin_info":{
	// "plugin_id":"plugin-email-noti-protocol",
	// "version":"1.0.1",
	// "options":{
	//
	// },
	// "secret_id":"secret-123456789012",
	// "metadata":{
	// "data_type":"PLAIN_TEXT",
	// "data":{
	// "schema":{
	// "properties":{
	// "email":{
	// "pattern":"^[\\W]*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4}[\\W]*,{1}[\\W]*)*([\\w+\\-.%]+@[\\w\\-.]+\\.[A-Za-z]{2,4})[\\W]*$",
	// "examples":[
	// "user1@test.com, user2@test.com"
	// ],
	// "minLength":10.0,
	// "description":"Email address to receive notifications",
	// "type":"string",
	// "title":"Email Address"
	// }
	// },
	// "required":[
	// "email"
	// ],
	// "type":"object"
	// }
	// }
	// },
	// "upgrade_mode":"AUTO"
	// },
	// "tags":{
	//
	// },
	// "domain_id":"domain-123456789012",
	// "created_at":"2022-01-01T07:55:57.043Z"
	// }
	// ],
	// "total_count":1
	// }
	List(context.Context, *ProtocolQuery) (*ProtocolsInfo, error)
	Stat(context.Context, *ProtocolStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedProtocolServer()
}

// UnimplementedProtocolServer must be embedded to have forward compatible implementations.
type UnimplementedProtocolServer struct {
}

func (UnimplementedProtocolServer) Create(context.Context, *CreateProtocolRequest) (*ProtocolInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedProtocolServer) Update(context.Context, *UpdateProtocolRequest) (*ProtocolInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedProtocolServer) UpdatePlugin(context.Context, *UpdateProtocolPluginRequest) (*ProtocolInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlugin not implemented")
}
func (UnimplementedProtocolServer) Enable(context.Context, *ProtocolRequest) (*ProtocolInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enable not implemented")
}
func (UnimplementedProtocolServer) Disable(context.Context, *ProtocolRequest) (*ProtocolInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disable not implemented")
}
func (UnimplementedProtocolServer) Delete(context.Context, *ProtocolRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedProtocolServer) Get(context.Context, *GetProtocolRequest) (*ProtocolInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedProtocolServer) List(context.Context, *ProtocolQuery) (*ProtocolsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedProtocolServer) Stat(context.Context, *ProtocolStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedProtocolServer) mustEmbedUnimplementedProtocolServer() {}

// UnsafeProtocolServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProtocolServer will
// result in compilation errors.
type UnsafeProtocolServer interface {
	mustEmbedUnimplementedProtocolServer()
}

func RegisterProtocolServer(s grpc.ServiceRegistrar, srv ProtocolServer) {
	s.RegisterService(&Protocol_ServiceDesc, srv)
}

func _Protocol_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProtocolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Protocol_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServer).Create(ctx, req.(*CreateProtocolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Protocol_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProtocolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Protocol_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServer).Update(ctx, req.(*UpdateProtocolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Protocol_UpdatePlugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProtocolPluginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServer).UpdatePlugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Protocol_UpdatePlugin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServer).UpdatePlugin(ctx, req.(*UpdateProtocolPluginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Protocol_Enable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProtocolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServer).Enable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Protocol_Enable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServer).Enable(ctx, req.(*ProtocolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Protocol_Disable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProtocolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServer).Disable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Protocol_Disable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServer).Disable(ctx, req.(*ProtocolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Protocol_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProtocolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Protocol_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServer).Delete(ctx, req.(*ProtocolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Protocol_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProtocolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Protocol_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServer).Get(ctx, req.(*GetProtocolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Protocol_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProtocolQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Protocol_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServer).List(ctx, req.(*ProtocolQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Protocol_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProtocolStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Protocol_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServer).Stat(ctx, req.(*ProtocolStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// Protocol_ServiceDesc is the grpc.ServiceDesc for Protocol service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Protocol_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.notification.v1.Protocol",
	HandlerType: (*ProtocolServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _Protocol_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _Protocol_Update_Handler,
		},
		{
			MethodName: "update_plugin",
			Handler:    _Protocol_UpdatePlugin_Handler,
		},
		{
			MethodName: "enable",
			Handler:    _Protocol_Enable_Handler,
		},
		{
			MethodName: "disable",
			Handler:    _Protocol_Disable_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _Protocol_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _Protocol_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _Protocol_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _Protocol_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/notification/v1/protocol.proto",
}
