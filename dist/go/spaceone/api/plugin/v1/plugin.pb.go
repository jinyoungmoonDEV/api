//
//desc: A Plugin is a resource managing endpoints of the plugin instances deployed. If there is a plugin instance that does not work properly, the Plugin requests the Supervisor to redeploy the instance.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.6.1
// source: spaceone/api/plugin/v1/plugin.proto

package v1

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PluginEndpointRequest_UpgradeMode int32

const (
	PluginEndpointRequest_MANUAL PluginEndpointRequest_UpgradeMode = 0
	PluginEndpointRequest_AUTO   PluginEndpointRequest_UpgradeMode = 1
)

// Enum value maps for PluginEndpointRequest_UpgradeMode.
var (
	PluginEndpointRequest_UpgradeMode_name = map[int32]string{
		0: "MANUAL",
		1: "AUTO",
	}
	PluginEndpointRequest_UpgradeMode_value = map[string]int32{
		"MANUAL": 0,
		"AUTO":   1,
	}
)

func (x PluginEndpointRequest_UpgradeMode) Enum() *PluginEndpointRequest_UpgradeMode {
	p := new(PluginEndpointRequest_UpgradeMode)
	*p = x
	return p
}

func (x PluginEndpointRequest_UpgradeMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PluginEndpointRequest_UpgradeMode) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_plugin_v1_plugin_proto_enumTypes[0].Descriptor()
}

func (PluginEndpointRequest_UpgradeMode) Type() protoreflect.EnumType {
	return &file_spaceone_api_plugin_v1_plugin_proto_enumTypes[0]
}

func (x PluginEndpointRequest_UpgradeMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PluginEndpointRequest_UpgradeMode.Descriptor instead.
func (PluginEndpointRequest_UpgradeMode) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_plugin_v1_plugin_proto_rawDescGZIP(), []int{0, 0}
}

type PluginEndpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// is_required: true
	PluginId string `protobuf:"bytes,1,opt,name=plugin_id,json=pluginId,proto3" json:"plugin_id,omitempty"`
	// is_required: false
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// is_required: false
	Labels *_struct.Struct `protobuf:"bytes,3,opt,name=labels,proto3" json:"labels,omitempty"`
	// is_required: true
	DomainId string `protobuf:"bytes,4,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	// is_required: false
	UpgradeMode PluginEndpointRequest_UpgradeMode `protobuf:"varint,5,opt,name=upgrade_mode,json=upgradeMode,proto3,enum=spaceone.api.plugin.v1.PluginEndpointRequest_UpgradeMode" json:"upgrade_mode,omitempty"`
}

func (x *PluginEndpointRequest) Reset() {
	*x = PluginEndpointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_plugin_v1_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginEndpointRequest) ProtoMessage() {}

func (x *PluginEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_plugin_v1_plugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginEndpointRequest.ProtoReflect.Descriptor instead.
func (*PluginEndpointRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_plugin_v1_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *PluginEndpointRequest) GetPluginId() string {
	if x != nil {
		return x.PluginId
	}
	return ""
}

func (x *PluginEndpointRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *PluginEndpointRequest) GetLabels() *_struct.Struct {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *PluginEndpointRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *PluginEndpointRequest) GetUpgradeMode() PluginEndpointRequest_UpgradeMode {
	if x != nil {
		return x.UpgradeMode
	}
	return PluginEndpointRequest_MANUAL
}

type PluginFailureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// is_required: true
	SupervisorId string `protobuf:"bytes,1,opt,name=supervisor_id,json=supervisorId,proto3" json:"supervisor_id,omitempty"`
	// is_required: true
	PluginId string `protobuf:"bytes,2,opt,name=plugin_id,json=pluginId,proto3" json:"plugin_id,omitempty"`
	// is_required: true
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// is_required: true
	DomainId string `protobuf:"bytes,4,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *PluginFailureRequest) Reset() {
	*x = PluginFailureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_plugin_v1_plugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginFailureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginFailureRequest) ProtoMessage() {}

func (x *PluginFailureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_plugin_v1_plugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginFailureRequest.ProtoReflect.Descriptor instead.
func (*PluginFailureRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_plugin_v1_plugin_proto_rawDescGZIP(), []int{1}
}

func (x *PluginFailureRequest) GetSupervisorId() string {
	if x != nil {
		return x.SupervisorId
	}
	return ""
}

func (x *PluginFailureRequest) GetPluginId() string {
	if x != nil {
		return x.PluginId
	}
	return ""
}

func (x *PluginFailureRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *PluginFailureRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type PluginEndpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint       string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	AccessToken    string `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	UpdatedVersion string `protobuf:"bytes,3,opt,name=updated_version,json=updatedVersion,proto3" json:"updated_version,omitempty"`
}

func (x *PluginEndpoint) Reset() {
	*x = PluginEndpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_plugin_v1_plugin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginEndpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginEndpoint) ProtoMessage() {}

func (x *PluginEndpoint) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_plugin_v1_plugin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginEndpoint.ProtoReflect.Descriptor instead.
func (*PluginEndpoint) Descriptor() ([]byte, []int) {
	return file_spaceone_api_plugin_v1_plugin_proto_rawDescGZIP(), []int{2}
}

func (x *PluginEndpoint) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *PluginEndpoint) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *PluginEndpoint) GetUpdatedVersion() string {
	if x != nil {
		return x.UpdatedVersion
	}
	return ""
}

var File_spaceone_api_plugin_v1_plugin_proto protoreflect.FileDescriptor

var file_spaceone_api_plugin_v1_plugin_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x02, 0x0a, 0x15, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x5c, 0x0a, 0x0c, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x39, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0b, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x22, 0x23, 0x0a, 0x0b, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x4d, 0x6f,
	0x64, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x41, 0x4e, 0x55, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x41, 0x55, 0x54, 0x4f, 0x10, 0x01, 0x22, 0x8f, 0x01, 0x0a, 0x14, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76,
	0x69, 0x73, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a,
	0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x78, 0x0a, 0x0e, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x32, 0xaf, 0x02, 0x0a, 0x06, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12,
	0x9e, 0x01, 0x0a, 0x13, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2d, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f,
	0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x30,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x3a, 0x01, 0x2a, 0x22, 0x25, 0x2f, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x67, 0x65, 0x74,
	0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x83, 0x01, 0x0a, 0x0e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x66, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x12, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x25, 0x3a, 0x01, 0x2a, 0x22, 0x20, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2d, 0x66,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d,
	0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_plugin_v1_plugin_proto_rawDescOnce sync.Once
	file_spaceone_api_plugin_v1_plugin_proto_rawDescData = file_spaceone_api_plugin_v1_plugin_proto_rawDesc
)

func file_spaceone_api_plugin_v1_plugin_proto_rawDescGZIP() []byte {
	file_spaceone_api_plugin_v1_plugin_proto_rawDescOnce.Do(func() {
		file_spaceone_api_plugin_v1_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_plugin_v1_plugin_proto_rawDescData)
	})
	return file_spaceone_api_plugin_v1_plugin_proto_rawDescData
}

var file_spaceone_api_plugin_v1_plugin_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spaceone_api_plugin_v1_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_spaceone_api_plugin_v1_plugin_proto_goTypes = []interface{}{
	(PluginEndpointRequest_UpgradeMode)(0), // 0: spaceone.api.plugin.v1.PluginEndpointRequest.UpgradeMode
	(*PluginEndpointRequest)(nil),          // 1: spaceone.api.plugin.v1.PluginEndpointRequest
	(*PluginFailureRequest)(nil),           // 2: spaceone.api.plugin.v1.PluginFailureRequest
	(*PluginEndpoint)(nil),                 // 3: spaceone.api.plugin.v1.PluginEndpoint
	(*_struct.Struct)(nil),                 // 4: google.protobuf.Struct
	(*empty.Empty)(nil),                    // 5: google.protobuf.Empty
}
var file_spaceone_api_plugin_v1_plugin_proto_depIdxs = []int32{
	4, // 0: spaceone.api.plugin.v1.PluginEndpointRequest.labels:type_name -> google.protobuf.Struct
	0, // 1: spaceone.api.plugin.v1.PluginEndpointRequest.upgrade_mode:type_name -> spaceone.api.plugin.v1.PluginEndpointRequest.UpgradeMode
	1, // 2: spaceone.api.plugin.v1.Plugin.get_plugin_endpoint:input_type -> spaceone.api.plugin.v1.PluginEndpointRequest
	2, // 3: spaceone.api.plugin.v1.Plugin.notify_failure:input_type -> spaceone.api.plugin.v1.PluginFailureRequest
	3, // 4: spaceone.api.plugin.v1.Plugin.get_plugin_endpoint:output_type -> spaceone.api.plugin.v1.PluginEndpoint
	5, // 5: spaceone.api.plugin.v1.Plugin.notify_failure:output_type -> google.protobuf.Empty
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_spaceone_api_plugin_v1_plugin_proto_init() }
func file_spaceone_api_plugin_v1_plugin_proto_init() {
	if File_spaceone_api_plugin_v1_plugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spaceone_api_plugin_v1_plugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginEndpointRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spaceone_api_plugin_v1_plugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginFailureRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spaceone_api_plugin_v1_plugin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginEndpoint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spaceone_api_plugin_v1_plugin_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_plugin_v1_plugin_proto_goTypes,
		DependencyIndexes: file_spaceone_api_plugin_v1_plugin_proto_depIdxs,
		EnumInfos:         file_spaceone_api_plugin_v1_plugin_proto_enumTypes,
		MessageInfos:      file_spaceone_api_plugin_v1_plugin_proto_msgTypes,
	}.Build()
	File_spaceone_api_plugin_v1_plugin_proto = out.File
	file_spaceone_api_plugin_v1_plugin_proto_rawDesc = nil
	file_spaceone_api_plugin_v1_plugin_proto_goTypes = nil
	file_spaceone_api_plugin_v1_plugin_proto_depIdxs = nil
}
