// UserConfig API which configure environments for user

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v3.6.1
// source: spaceone/api/config/v1/user_config.proto

package v1

import (
	v2 "github.com/cloudforet-io/api/dist/go/spaceone/api/core/v2"
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

type SetUserConfigRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Name  string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data  *_struct.Struct        `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// +optional
	Tags          *_struct.Struct `protobuf:"bytes,3,opt,name=tags,proto3" json:"tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetUserConfigRequest) Reset() {
	*x = SetUserConfigRequest{}
	mi := &file_spaceone_api_config_v1_user_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetUserConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetUserConfigRequest) ProtoMessage() {}

func (x *SetUserConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_config_v1_user_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetUserConfigRequest.ProtoReflect.Descriptor instead.
func (*SetUserConfigRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_config_v1_user_config_proto_rawDescGZIP(), []int{0}
}

func (x *SetUserConfigRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SetUserConfigRequest) GetData() *_struct.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SetUserConfigRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

type UserConfigRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserConfigRequest) Reset() {
	*x = UserConfigRequest{}
	mi := &file_spaceone_api_config_v1_user_config_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserConfigRequest) ProtoMessage() {}

func (x *UserConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_config_v1_user_config_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserConfigRequest.ProtoReflect.Descriptor instead.
func (*UserConfigRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_config_v1_user_config_proto_rawDescGZIP(), []int{1}
}

func (x *UserConfigRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UserConfigQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// +optional
	Query *v2.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserConfigQuery) Reset() {
	*x = UserConfigQuery{}
	mi := &file_spaceone_api_config_v1_user_config_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserConfigQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserConfigQuery) ProtoMessage() {}

func (x *UserConfigQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_config_v1_user_config_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserConfigQuery.ProtoReflect.Descriptor instead.
func (*UserConfigQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_config_v1_user_config_proto_rawDescGZIP(), []int{2}
}

func (x *UserConfigQuery) GetQuery() *v2.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *UserConfigQuery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UserConfigInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data          *_struct.Struct        `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Tags          *_struct.Struct        `protobuf:"bytes,3,opt,name=tags,proto3" json:"tags,omitempty"`
	DomainId      string                 `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	UserId        string                 `protobuf:"bytes,22,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,32,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserConfigInfo) Reset() {
	*x = UserConfigInfo{}
	mi := &file_spaceone_api_config_v1_user_config_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserConfigInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserConfigInfo) ProtoMessage() {}

func (x *UserConfigInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_config_v1_user_config_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserConfigInfo.ProtoReflect.Descriptor instead.
func (*UserConfigInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_config_v1_user_config_proto_rawDescGZIP(), []int{3}
}

func (x *UserConfigInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserConfigInfo) GetData() *_struct.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *UserConfigInfo) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *UserConfigInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *UserConfigInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserConfigInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *UserConfigInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type UserConfigsInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Results       []*UserConfigInfo      `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount    int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserConfigsInfo) Reset() {
	*x = UserConfigsInfo{}
	mi := &file_spaceone_api_config_v1_user_config_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserConfigsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserConfigsInfo) ProtoMessage() {}

func (x *UserConfigsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_config_v1_user_config_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserConfigsInfo.ProtoReflect.Descriptor instead.
func (*UserConfigsInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_config_v1_user_config_proto_rawDescGZIP(), []int{4}
}

func (x *UserConfigsInfo) GetResults() []*UserConfigInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *UserConfigsInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type UserConfigStatQuery struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         *v2.StatisticsQuery    `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserConfigStatQuery) Reset() {
	*x = UserConfigStatQuery{}
	mi := &file_spaceone_api_config_v1_user_config_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserConfigStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserConfigStatQuery) ProtoMessage() {}

func (x *UserConfigStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_config_v1_user_config_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserConfigStatQuery.ProtoReflect.Descriptor instead.
func (*UserConfigStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_config_v1_user_config_proto_rawDescGZIP(), []int{5}
}

func (x *UserConfigStatQuery) GetQuery() *v2.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

var File_spaceone_api_config_v1_user_config_proto protoreflect.FileDescriptor

var file_spaceone_api_config_v1_user_config_proto_rawDesc = []byte{
	0x0a, 0x28, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x32, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01,
	0x0a, 0x14, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x22, 0x27, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x58, 0x0a,
	0x0f, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x31, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xf2, 0x01, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2b, 0x0a, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x1f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x20, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x74, 0x0a, 0x0f,
	0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x40, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x52, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x53, 0x74, 0x61, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x3b, 0x0a, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x32, 0x98, 0x07, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x88, 0x01, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x01,
	0x2a, 0x22, 0x1d, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x88, 0x01, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x01, 0x2a, 0x22, 0x1d, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x03,
	0x73, 0x65, 0x74, 0x12, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x26, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x65, 0x74,
	0x12, 0x75, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x29, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x28, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x01, 0x2a, 0x22, 0x1d, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x7f, 0x0a, 0x03, 0x67, 0x65, 0x74, 0x12, 0x29,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x80, 0x01, 0x0a, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x12, 0x27, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x27, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x74, 0x0a, 0x04, 0x73,
	0x74, 0x61, 0x74, 0x12, 0x2b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x74, 0x61, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f,
	0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_config_v1_user_config_proto_rawDescOnce sync.Once
	file_spaceone_api_config_v1_user_config_proto_rawDescData = file_spaceone_api_config_v1_user_config_proto_rawDesc
)

func file_spaceone_api_config_v1_user_config_proto_rawDescGZIP() []byte {
	file_spaceone_api_config_v1_user_config_proto_rawDescOnce.Do(func() {
		file_spaceone_api_config_v1_user_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_config_v1_user_config_proto_rawDescData)
	})
	return file_spaceone_api_config_v1_user_config_proto_rawDescData
}

var file_spaceone_api_config_v1_user_config_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_spaceone_api_config_v1_user_config_proto_goTypes = []any{
	(*SetUserConfigRequest)(nil), // 0: spaceone.api.config.v1.SetUserConfigRequest
	(*UserConfigRequest)(nil),    // 1: spaceone.api.config.v1.UserConfigRequest
	(*UserConfigQuery)(nil),      // 2: spaceone.api.config.v1.UserConfigQuery
	(*UserConfigInfo)(nil),       // 3: spaceone.api.config.v1.UserConfigInfo
	(*UserConfigsInfo)(nil),      // 4: spaceone.api.config.v1.UserConfigsInfo
	(*UserConfigStatQuery)(nil),  // 5: spaceone.api.config.v1.UserConfigStatQuery
	(*_struct.Struct)(nil),       // 6: google.protobuf.Struct
	(*v2.Query)(nil),             // 7: spaceone.api.core.v2.Query
	(*v2.StatisticsQuery)(nil),   // 8: spaceone.api.core.v2.StatisticsQuery
	(*empty.Empty)(nil),          // 9: google.protobuf.Empty
}
var file_spaceone_api_config_v1_user_config_proto_depIdxs = []int32{
	6,  // 0: spaceone.api.config.v1.SetUserConfigRequest.data:type_name -> google.protobuf.Struct
	6,  // 1: spaceone.api.config.v1.SetUserConfigRequest.tags:type_name -> google.protobuf.Struct
	7,  // 2: spaceone.api.config.v1.UserConfigQuery.query:type_name -> spaceone.api.core.v2.Query
	6,  // 3: spaceone.api.config.v1.UserConfigInfo.data:type_name -> google.protobuf.Struct
	6,  // 4: spaceone.api.config.v1.UserConfigInfo.tags:type_name -> google.protobuf.Struct
	3,  // 5: spaceone.api.config.v1.UserConfigsInfo.results:type_name -> spaceone.api.config.v1.UserConfigInfo
	8,  // 6: spaceone.api.config.v1.UserConfigStatQuery.query:type_name -> spaceone.api.core.v2.StatisticsQuery
	0,  // 7: spaceone.api.config.v1.UserConfig.create:input_type -> spaceone.api.config.v1.SetUserConfigRequest
	0,  // 8: spaceone.api.config.v1.UserConfig.update:input_type -> spaceone.api.config.v1.SetUserConfigRequest
	0,  // 9: spaceone.api.config.v1.UserConfig.set:input_type -> spaceone.api.config.v1.SetUserConfigRequest
	1,  // 10: spaceone.api.config.v1.UserConfig.delete:input_type -> spaceone.api.config.v1.UserConfigRequest
	1,  // 11: spaceone.api.config.v1.UserConfig.get:input_type -> spaceone.api.config.v1.UserConfigRequest
	2,  // 12: spaceone.api.config.v1.UserConfig.list:input_type -> spaceone.api.config.v1.UserConfigQuery
	5,  // 13: spaceone.api.config.v1.UserConfig.stat:input_type -> spaceone.api.config.v1.UserConfigStatQuery
	3,  // 14: spaceone.api.config.v1.UserConfig.create:output_type -> spaceone.api.config.v1.UserConfigInfo
	3,  // 15: spaceone.api.config.v1.UserConfig.update:output_type -> spaceone.api.config.v1.UserConfigInfo
	3,  // 16: spaceone.api.config.v1.UserConfig.set:output_type -> spaceone.api.config.v1.UserConfigInfo
	9,  // 17: spaceone.api.config.v1.UserConfig.delete:output_type -> google.protobuf.Empty
	3,  // 18: spaceone.api.config.v1.UserConfig.get:output_type -> spaceone.api.config.v1.UserConfigInfo
	4,  // 19: spaceone.api.config.v1.UserConfig.list:output_type -> spaceone.api.config.v1.UserConfigsInfo
	6,  // 20: spaceone.api.config.v1.UserConfig.stat:output_type -> google.protobuf.Struct
	14, // [14:21] is the sub-list for method output_type
	7,  // [7:14] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_spaceone_api_config_v1_user_config_proto_init() }
func file_spaceone_api_config_v1_user_config_proto_init() {
	if File_spaceone_api_config_v1_user_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spaceone_api_config_v1_user_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_config_v1_user_config_proto_goTypes,
		DependencyIndexes: file_spaceone_api_config_v1_user_config_proto_depIdxs,
		MessageInfos:      file_spaceone_api_config_v1_user_config_proto_msgTypes,
	}.Build()
	File_spaceone_api_config_v1_user_config_proto = out.File
	file_spaceone_api_config_v1_user_config_proto_rawDesc = nil
	file_spaceone_api_config_v1_user_config_proto_goTypes = nil
	file_spaceone_api_config_v1_user_config_proto_depIdxs = nil
}
