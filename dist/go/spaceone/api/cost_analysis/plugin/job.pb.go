// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v3.6.1
// source: spaceone/api/cost_analysis/plugin/job.proto

package plugin

import (
	_struct "github.com/golang/protobuf/ptypes/struct"
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

// {
//
// }
type GetTasksRequest struct {
	state      protoimpl.MessageState `protogen:"open.v1"`
	Options    *_struct.Struct        `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	SecretData *_struct.Struct        `protobuf:"bytes,2,opt,name=secret_data,json=secretData,proto3" json:"secret_data,omitempty"`
	// +optional
	Schema string `protobuf:"bytes,3,opt,name=schema,proto3" json:"schema,omitempty"`
	// +optional
	Start string `protobuf:"bytes,4,opt,name=start,proto3" json:"start,omitempty"`
	// +optional
	LastSynchronizedAt string            `protobuf:"bytes,5,opt,name=last_synchronized_at,json=lastSynchronizedAt,proto3" json:"last_synchronized_at,omitempty"`
	DomainId           string            `protobuf:"bytes,6,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	LinkedAccounts     []*_struct.Struct `protobuf:"bytes,7,rep,name=linked_accounts,json=linkedAccounts,proto3" json:"linked_accounts,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *GetTasksRequest) Reset() {
	*x = GetTasksRequest{}
	mi := &file_spaceone_api_cost_analysis_plugin_job_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTasksRequest) ProtoMessage() {}

func (x *GetTasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_plugin_job_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTasksRequest.ProtoReflect.Descriptor instead.
func (*GetTasksRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_plugin_job_proto_rawDescGZIP(), []int{0}
}

func (x *GetTasksRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *GetTasksRequest) GetSecretData() *_struct.Struct {
	if x != nil {
		return x.SecretData
	}
	return nil
}

func (x *GetTasksRequest) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *GetTasksRequest) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *GetTasksRequest) GetLastSynchronizedAt() string {
	if x != nil {
		return x.LastSynchronizedAt
	}
	return ""
}

func (x *GetTasksRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *GetTasksRequest) GetLinkedAccounts() []*_struct.Struct {
	if x != nil {
		return x.LinkedAccounts
	}
	return nil
}

type TaskInfo struct {
	state       protoimpl.MessageState `protogen:"open.v1"`
	TaskOptions *_struct.Struct        `protobuf:"bytes,1,opt,name=task_options,json=taskOptions,proto3" json:"task_options,omitempty"`
	// +optional
	TaskChanged   *ChangedInfo `protobuf:"bytes,2,opt,name=task_changed,json=taskChanged,proto3" json:"task_changed,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskInfo) Reset() {
	*x = TaskInfo{}
	mi := &file_spaceone_api_cost_analysis_plugin_job_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskInfo) ProtoMessage() {}

func (x *TaskInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_plugin_job_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskInfo.ProtoReflect.Descriptor instead.
func (*TaskInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_plugin_job_proto_rawDescGZIP(), []int{1}
}

func (x *TaskInfo) GetTaskOptions() *_struct.Struct {
	if x != nil {
		return x.TaskOptions
	}
	return nil
}

func (x *TaskInfo) GetTaskChanged() *ChangedInfo {
	if x != nil {
		return x.TaskChanged
	}
	return nil
}

type ChangedInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Start         string                 `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End           string                 `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	Filter        *_struct.Struct        `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChangedInfo) Reset() {
	*x = ChangedInfo{}
	mi := &file_spaceone_api_cost_analysis_plugin_job_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangedInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangedInfo) ProtoMessage() {}

func (x *ChangedInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_plugin_job_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangedInfo.ProtoReflect.Descriptor instead.
func (*ChangedInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_plugin_job_proto_rawDescGZIP(), []int{2}
}

func (x *ChangedInfo) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *ChangedInfo) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

func (x *ChangedInfo) GetFilter() *_struct.Struct {
	if x != nil {
		return x.Filter
	}
	return nil
}

type SyncedAccountInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccountId     string                 `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SyncedAccountInfo) Reset() {
	*x = SyncedAccountInfo{}
	mi := &file_spaceone_api_cost_analysis_plugin_job_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SyncedAccountInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncedAccountInfo) ProtoMessage() {}

func (x *SyncedAccountInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_plugin_job_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncedAccountInfo.ProtoReflect.Descriptor instead.
func (*SyncedAccountInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_plugin_job_proto_rawDescGZIP(), []int{3}
}

func (x *SyncedAccountInfo) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *SyncedAccountInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// {
//
// }
type TasksInfo struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Tasks          []*TaskInfo            `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	Changed        []*ChangedInfo         `protobuf:"bytes,2,rep,name=changed,proto3" json:"changed,omitempty"`
	SyncedAccounts []*SyncedAccountInfo   `protobuf:"bytes,3,rep,name=synced_accounts,json=syncedAccounts,proto3" json:"synced_accounts,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *TasksInfo) Reset() {
	*x = TasksInfo{}
	mi := &file_spaceone_api_cost_analysis_plugin_job_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TasksInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TasksInfo) ProtoMessage() {}

func (x *TasksInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_plugin_job_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TasksInfo.ProtoReflect.Descriptor instead.
func (*TasksInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_plugin_job_proto_rawDescGZIP(), []int{4}
}

func (x *TasksInfo) GetTasks() []*TaskInfo {
	if x != nil {
		return x.Tasks
	}
	return nil
}

func (x *TasksInfo) GetChanged() []*ChangedInfo {
	if x != nil {
		return x.Changed
	}
	return nil
}

func (x *TasksInfo) GetSyncedAccounts() []*SyncedAccountInfo {
	if x != nil {
		return x.SyncedAccounts
	}
	return nil
}

var File_spaceone_api_cost_analysis_plugin_job_proto protoreflect.FileDescriptor

var file_spaceone_api_cost_analysis_plugin_job_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74,
	0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd,
	0x02, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x38, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x30, 0x0a,
	0x14, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6c, 0x61, 0x73,
	0x74, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x0f,
	0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0e,
	0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0x99,
	0x01, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3a, 0x0a, 0x0c, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0b, 0x74, 0x61, 0x73, 0x6b,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x51, 0x0a, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x5f,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73,
	0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x74,
	0x61, 0x73, 0x6b, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x22, 0x66, 0x0a, 0x0b, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e,
	0x64, 0x12, 0x2f, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x22, 0x46, 0x0a, 0x11, 0x53, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xf7, 0x01, 0x0a, 0x09, 0x54,
	0x61, 0x73, 0x6b, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x41, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f,
	0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x48, 0x0a, 0x07, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74,
	0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x5d, 0x0a, 0x0f, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x32, 0x76, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x6f, 0x0a, 0x09, 0x67,
	0x65, 0x74, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x32, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74,
	0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x42, 0x48, 0x5a, 0x46,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73,
	0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_cost_analysis_plugin_job_proto_rawDescOnce sync.Once
	file_spaceone_api_cost_analysis_plugin_job_proto_rawDescData = file_spaceone_api_cost_analysis_plugin_job_proto_rawDesc
)

func file_spaceone_api_cost_analysis_plugin_job_proto_rawDescGZIP() []byte {
	file_spaceone_api_cost_analysis_plugin_job_proto_rawDescOnce.Do(func() {
		file_spaceone_api_cost_analysis_plugin_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_cost_analysis_plugin_job_proto_rawDescData)
	})
	return file_spaceone_api_cost_analysis_plugin_job_proto_rawDescData
}

var file_spaceone_api_cost_analysis_plugin_job_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_spaceone_api_cost_analysis_plugin_job_proto_goTypes = []any{
	(*GetTasksRequest)(nil),   // 0: spaceone.api.cost_analysis.plugin.GetTasksRequest
	(*TaskInfo)(nil),          // 1: spaceone.api.cost_analysis.plugin.TaskInfo
	(*ChangedInfo)(nil),       // 2: spaceone.api.cost_analysis.plugin.ChangedInfo
	(*SyncedAccountInfo)(nil), // 3: spaceone.api.cost_analysis.plugin.SyncedAccountInfo
	(*TasksInfo)(nil),         // 4: spaceone.api.cost_analysis.plugin.TasksInfo
	(*_struct.Struct)(nil),    // 5: google.protobuf.Struct
}
var file_spaceone_api_cost_analysis_plugin_job_proto_depIdxs = []int32{
	5,  // 0: spaceone.api.cost_analysis.plugin.GetTasksRequest.options:type_name -> google.protobuf.Struct
	5,  // 1: spaceone.api.cost_analysis.plugin.GetTasksRequest.secret_data:type_name -> google.protobuf.Struct
	5,  // 2: spaceone.api.cost_analysis.plugin.GetTasksRequest.linked_accounts:type_name -> google.protobuf.Struct
	5,  // 3: spaceone.api.cost_analysis.plugin.TaskInfo.task_options:type_name -> google.protobuf.Struct
	2,  // 4: spaceone.api.cost_analysis.plugin.TaskInfo.task_changed:type_name -> spaceone.api.cost_analysis.plugin.ChangedInfo
	5,  // 5: spaceone.api.cost_analysis.plugin.ChangedInfo.filter:type_name -> google.protobuf.Struct
	1,  // 6: spaceone.api.cost_analysis.plugin.TasksInfo.tasks:type_name -> spaceone.api.cost_analysis.plugin.TaskInfo
	2,  // 7: spaceone.api.cost_analysis.plugin.TasksInfo.changed:type_name -> spaceone.api.cost_analysis.plugin.ChangedInfo
	3,  // 8: spaceone.api.cost_analysis.plugin.TasksInfo.synced_accounts:type_name -> spaceone.api.cost_analysis.plugin.SyncedAccountInfo
	0,  // 9: spaceone.api.cost_analysis.plugin.Job.get_tasks:input_type -> spaceone.api.cost_analysis.plugin.GetTasksRequest
	4,  // 10: spaceone.api.cost_analysis.plugin.Job.get_tasks:output_type -> spaceone.api.cost_analysis.plugin.TasksInfo
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_spaceone_api_cost_analysis_plugin_job_proto_init() }
func file_spaceone_api_cost_analysis_plugin_job_proto_init() {
	if File_spaceone_api_cost_analysis_plugin_job_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spaceone_api_cost_analysis_plugin_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_cost_analysis_plugin_job_proto_goTypes,
		DependencyIndexes: file_spaceone_api_cost_analysis_plugin_job_proto_depIdxs,
		MessageInfos:      file_spaceone_api_cost_analysis_plugin_job_proto_msgTypes,
	}.Build()
	File_spaceone_api_cost_analysis_plugin_job_proto = out.File
	file_spaceone_api_cost_analysis_plugin_job_proto_rawDesc = nil
	file_spaceone_api_cost_analysis_plugin_job_proto_goTypes = nil
	file_spaceone_api_cost_analysis_plugin_job_proto_depIdxs = nil
}
