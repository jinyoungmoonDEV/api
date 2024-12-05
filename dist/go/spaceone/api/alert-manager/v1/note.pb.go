// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.6.1
// source: spaceone/api/alert_manager/v1/note.proto

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

type NoteCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlertId string `protobuf:"bytes,1,opt,name=alert_id,json=alertId,proto3" json:"alert_id,omitempty"`
	Note    string `protobuf:"bytes,2,opt,name=note,proto3" json:"note,omitempty"`
}

func (x *NoteCreateRequest) Reset() {
	*x = NoteCreateRequest{}
	mi := &file_spaceone_api_alert_manager_v1_note_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NoteCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteCreateRequest) ProtoMessage() {}

func (x *NoteCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_alert_manager_v1_note_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteCreateRequest.ProtoReflect.Descriptor instead.
func (*NoteCreateRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_alert_manager_v1_note_proto_rawDescGZIP(), []int{0}
}

func (x *NoteCreateRequest) GetAlertId() string {
	if x != nil {
		return x.AlertId
	}
	return ""
}

func (x *NoteCreateRequest) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

type NoteUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NoteId string `protobuf:"bytes,1,opt,name=note_id,json=noteId,proto3" json:"note_id,omitempty"`
	// +optional
	Note string `protobuf:"bytes,2,opt,name=note,proto3" json:"note,omitempty"`
}

func (x *NoteUpdateRequest) Reset() {
	*x = NoteUpdateRequest{}
	mi := &file_spaceone_api_alert_manager_v1_note_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NoteUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteUpdateRequest) ProtoMessage() {}

func (x *NoteUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_alert_manager_v1_note_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteUpdateRequest.ProtoReflect.Descriptor instead.
func (*NoteUpdateRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_alert_manager_v1_note_proto_rawDescGZIP(), []int{1}
}

func (x *NoteUpdateRequest) GetNoteId() string {
	if x != nil {
		return x.NoteId
	}
	return ""
}

func (x *NoteUpdateRequest) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

type NoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NoteId string `protobuf:"bytes,1,opt,name=note_id,json=noteId,proto3" json:"note_id,omitempty"`
}

func (x *NoteRequest) Reset() {
	*x = NoteRequest{}
	mi := &file_spaceone_api_alert_manager_v1_note_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteRequest) ProtoMessage() {}

func (x *NoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_alert_manager_v1_note_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteRequest.ProtoReflect.Descriptor instead.
func (*NoteRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_alert_manager_v1_note_proto_rawDescGZIP(), []int{2}
}

func (x *NoteRequest) GetNoteId() string {
	if x != nil {
		return x.NoteId
	}
	return ""
}

type NoteSearchQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// +optional
	Query *v2.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	WorkspaceId string `protobuf:"bytes,2,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	// +optional
	ServiceId string `protobuf:"bytes,3,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	// +optional
	AlertId string `protobuf:"bytes,4,opt,name=alert_id,json=alertId,proto3" json:"alert_id,omitempty"`
	// +optional
	NoteId string `protobuf:"bytes,5,opt,name=note_id,json=noteId,proto3" json:"note_id,omitempty"`
	// +optional
	CreatedBy string `protobuf:"bytes,31,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
}

func (x *NoteSearchQuery) Reset() {
	*x = NoteSearchQuery{}
	mi := &file_spaceone_api_alert_manager_v1_note_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NoteSearchQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteSearchQuery) ProtoMessage() {}

func (x *NoteSearchQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_alert_manager_v1_note_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteSearchQuery.ProtoReflect.Descriptor instead.
func (*NoteSearchQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_alert_manager_v1_note_proto_rawDescGZIP(), []int{3}
}

func (x *NoteSearchQuery) GetQuery() *v2.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *NoteSearchQuery) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *NoteSearchQuery) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *NoteSearchQuery) GetAlertId() string {
	if x != nil {
		return x.AlertId
	}
	return ""
}

func (x *NoteSearchQuery) GetNoteId() string {
	if x != nil {
		return x.NoteId
	}
	return ""
}

func (x *NoteSearchQuery) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

type NoteStatQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query *v2.StatisticsQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *NoteStatQuery) Reset() {
	*x = NoteStatQuery{}
	mi := &file_spaceone_api_alert_manager_v1_note_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NoteStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteStatQuery) ProtoMessage() {}

func (x *NoteStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_alert_manager_v1_note_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteStatQuery.ProtoReflect.Descriptor instead.
func (*NoteStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_alert_manager_v1_note_proto_rawDescGZIP(), []int{4}
}

func (x *NoteStatQuery) GetQuery() *v2.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

type NoteInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NoteId      string `protobuf:"bytes,1,opt,name=note_id,json=noteId,proto3" json:"note_id,omitempty"`
	Note        string `protobuf:"bytes,2,opt,name=note,proto3" json:"note,omitempty"`
	DomainId    string `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	WorkspaceId string `protobuf:"bytes,22,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	ServiceId   string `protobuf:"bytes,23,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	AlertId     string `protobuf:"bytes,24,opt,name=alert_id,json=alertId,proto3" json:"alert_id,omitempty"`
	CreatedAt   string `protobuf:"bytes,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy   string `protobuf:"bytes,32,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
}

func (x *NoteInfo) Reset() {
	*x = NoteInfo{}
	mi := &file_spaceone_api_alert_manager_v1_note_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NoteInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteInfo) ProtoMessage() {}

func (x *NoteInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_alert_manager_v1_note_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteInfo.ProtoReflect.Descriptor instead.
func (*NoteInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_alert_manager_v1_note_proto_rawDescGZIP(), []int{5}
}

func (x *NoteInfo) GetNoteId() string {
	if x != nil {
		return x.NoteId
	}
	return ""
}

func (x *NoteInfo) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *NoteInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *NoteInfo) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *NoteInfo) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *NoteInfo) GetAlertId() string {
	if x != nil {
		return x.AlertId
	}
	return ""
}

func (x *NoteInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *NoteInfo) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

type NotesInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results    []*NoteInfo `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount int32       `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *NotesInfo) Reset() {
	*x = NotesInfo{}
	mi := &file_spaceone_api_alert_manager_v1_note_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotesInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotesInfo) ProtoMessage() {}

func (x *NotesInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_alert_manager_v1_note_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotesInfo.ProtoReflect.Descriptor instead.
func (*NotesInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_alert_manager_v1_note_proto_rawDescGZIP(), []int{6}
}

func (x *NotesInfo) GetResults() []*NoteInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *NotesInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_spaceone_api_alert_manager_v1_note_proto protoreflect.FileDescriptor

var file_spaceone_api_alert_manager_v1_note_proto_rawDesc = []byte{
	0x0a, 0x28, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x11, 0x4e, 0x6f, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x22, 0x40, 0x0a, 0x11, 0x4e, 0x6f, 0x74, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6e, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x22, 0x26, 0x0a, 0x0b, 0x4e, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x74,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x65,
	0x49, 0x64, 0x22, 0xd9, 0x01, 0x0a, 0x0f, 0x4e, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x1f, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x22, 0x4c,
	0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x3b, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0xef, 0x01, 0x0a,
	0x08, 0x4e, 0x6f, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x74,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x65,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x1f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x20, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x22, 0x6f,
	0x0a, 0x09, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x41, 0x0a, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32,
	0xa4, 0x06, 0x0a, 0x04, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x8d, 0x01, 0x0a, 0x06, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x30, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x28,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x01, 0x2a, 0x22, 0x1d, 0x2f, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74,
	0x65, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x8d, 0x01, 0x0a, 0x06, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x30, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x28,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x01, 0x2a, 0x22, 0x1d, 0x2f, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74,
	0x65, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x76, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x2a, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x01,
	0x2a, 0x22, 0x1d, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x81, 0x01, 0x0a, 0x03, 0x67, 0x65, 0x74, 0x12, 0x2a, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x25, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x65,
	0x2f, 0x67, 0x65, 0x74, 0x12, 0x88, 0x01, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f,
	0x74, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x28, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f,
	0x74, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a,
	0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x75, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x74, 0x12, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f,
	0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x26,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74,
	0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d,
	0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_alert_manager_v1_note_proto_rawDescOnce sync.Once
	file_spaceone_api_alert_manager_v1_note_proto_rawDescData = file_spaceone_api_alert_manager_v1_note_proto_rawDesc
)

func file_spaceone_api_alert_manager_v1_note_proto_rawDescGZIP() []byte {
	file_spaceone_api_alert_manager_v1_note_proto_rawDescOnce.Do(func() {
		file_spaceone_api_alert_manager_v1_note_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_alert_manager_v1_note_proto_rawDescData)
	})
	return file_spaceone_api_alert_manager_v1_note_proto_rawDescData
}

var file_spaceone_api_alert_manager_v1_note_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_spaceone_api_alert_manager_v1_note_proto_goTypes = []any{
	(*NoteCreateRequest)(nil),  // 0: spaceone.api.alert_manager.v1.NoteCreateRequest
	(*NoteUpdateRequest)(nil),  // 1: spaceone.api.alert_manager.v1.NoteUpdateRequest
	(*NoteRequest)(nil),        // 2: spaceone.api.alert_manager.v1.NoteRequest
	(*NoteSearchQuery)(nil),    // 3: spaceone.api.alert_manager.v1.NoteSearchQuery
	(*NoteStatQuery)(nil),      // 4: spaceone.api.alert_manager.v1.NoteStatQuery
	(*NoteInfo)(nil),           // 5: spaceone.api.alert_manager.v1.NoteInfo
	(*NotesInfo)(nil),          // 6: spaceone.api.alert_manager.v1.NotesInfo
	(*v2.Query)(nil),           // 7: spaceone.api.core.v2.Query
	(*v2.StatisticsQuery)(nil), // 8: spaceone.api.core.v2.StatisticsQuery
	(*empty.Empty)(nil),        // 9: google.protobuf.Empty
	(*_struct.Struct)(nil),     // 10: google.protobuf.Struct
}
var file_spaceone_api_alert_manager_v1_note_proto_depIdxs = []int32{
	7,  // 0: spaceone.api.alert_manager.v1.NoteSearchQuery.query:type_name -> spaceone.api.core.v2.Query
	8,  // 1: spaceone.api.alert_manager.v1.NoteStatQuery.query:type_name -> spaceone.api.core.v2.StatisticsQuery
	5,  // 2: spaceone.api.alert_manager.v1.NotesInfo.results:type_name -> spaceone.api.alert_manager.v1.NoteInfo
	0,  // 3: spaceone.api.alert_manager.v1.Note.create:input_type -> spaceone.api.alert_manager.v1.NoteCreateRequest
	1,  // 4: spaceone.api.alert_manager.v1.Note.update:input_type -> spaceone.api.alert_manager.v1.NoteUpdateRequest
	2,  // 5: spaceone.api.alert_manager.v1.Note.delete:input_type -> spaceone.api.alert_manager.v1.NoteRequest
	2,  // 6: spaceone.api.alert_manager.v1.Note.get:input_type -> spaceone.api.alert_manager.v1.NoteRequest
	3,  // 7: spaceone.api.alert_manager.v1.Note.list:input_type -> spaceone.api.alert_manager.v1.NoteSearchQuery
	4,  // 8: spaceone.api.alert_manager.v1.Note.stat:input_type -> spaceone.api.alert_manager.v1.NoteStatQuery
	5,  // 9: spaceone.api.alert_manager.v1.Note.create:output_type -> spaceone.api.alert_manager.v1.NoteInfo
	5,  // 10: spaceone.api.alert_manager.v1.Note.update:output_type -> spaceone.api.alert_manager.v1.NoteInfo
	9,  // 11: spaceone.api.alert_manager.v1.Note.delete:output_type -> google.protobuf.Empty
	5,  // 12: spaceone.api.alert_manager.v1.Note.get:output_type -> spaceone.api.alert_manager.v1.NoteInfo
	6,  // 13: spaceone.api.alert_manager.v1.Note.list:output_type -> spaceone.api.alert_manager.v1.NotesInfo
	10, // 14: spaceone.api.alert_manager.v1.Note.stat:output_type -> google.protobuf.Struct
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_spaceone_api_alert_manager_v1_note_proto_init() }
func file_spaceone_api_alert_manager_v1_note_proto_init() {
	if File_spaceone_api_alert_manager_v1_note_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spaceone_api_alert_manager_v1_note_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_alert_manager_v1_note_proto_goTypes,
		DependencyIndexes: file_spaceone_api_alert_manager_v1_note_proto_depIdxs,
		MessageInfos:      file_spaceone_api_alert_manager_v1_note_proto_msgTypes,
	}.Build()
	File_spaceone_api_alert_manager_v1_note_proto = out.File
	file_spaceone_api_alert_manager_v1_note_proto_rawDesc = nil
	file_spaceone_api_alert_manager_v1_note_proto_goTypes = nil
	file_spaceone_api_alert_manager_v1_note_proto_depIdxs = nil
}
