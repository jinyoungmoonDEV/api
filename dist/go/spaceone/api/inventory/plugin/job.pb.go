// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v3.6.1
// source: spaceone/api/inventory/plugin/job.proto

package plugin

import (
	_ "github.com/golang/protobuf/ptypes/empty"
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

type TaskFilter struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	Providers          []string               `protobuf:"bytes,1,rep,name=providers,proto3" json:"providers,omitempty"`
	CloudServiceGroups []string               `protobuf:"bytes,2,rep,name=cloud_service_groups,json=cloudServiceGroups,proto3" json:"cloud_service_groups,omitempty"`
	CloudServiceTypes  []string               `protobuf:"bytes,3,rep,name=cloud_service_types,json=cloudServiceTypes,proto3" json:"cloud_service_types,omitempty"`
	RegionCodes        []string               `protobuf:"bytes,4,rep,name=region_codes,json=regionCodes,proto3" json:"region_codes,omitempty"`
	Resources          []string               `protobuf:"bytes,5,rep,name=resources,proto3" json:"resources,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *TaskFilter) Reset() {
	*x = TaskFilter{}
	mi := &file_spaceone_api_inventory_plugin_job_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskFilter) ProtoMessage() {}

func (x *TaskFilter) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_plugin_job_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskFilter.ProtoReflect.Descriptor instead.
func (*TaskFilter) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_plugin_job_proto_rawDescGZIP(), []int{0}
}

func (x *TaskFilter) GetProviders() []string {
	if x != nil {
		return x.Providers
	}
	return nil
}

func (x *TaskFilter) GetCloudServiceGroups() []string {
	if x != nil {
		return x.CloudServiceGroups
	}
	return nil
}

func (x *TaskFilter) GetCloudServiceTypes() []string {
	if x != nil {
		return x.CloudServiceTypes
	}
	return nil
}

func (x *TaskFilter) GetRegionCodes() []string {
	if x != nil {
		return x.RegionCodes
	}
	return nil
}

func (x *TaskFilter) GetResources() []string {
	if x != nil {
		return x.Resources
	}
	return nil
}

type GetTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SecretData    *_struct.Struct        `protobuf:"bytes,1,opt,name=secret_data,json=secretData,proto3" json:"secret_data,omitempty"`
	Options       *_struct.Struct        `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	TaskFilter    *TaskFilter            `protobuf:"bytes,3,opt,name=task_filter,json=taskFilter,proto3" json:"task_filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTaskRequest) Reset() {
	*x = GetTaskRequest{}
	mi := &file_spaceone_api_inventory_plugin_job_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskRequest) ProtoMessage() {}

func (x *GetTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_plugin_job_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskRequest.ProtoReflect.Descriptor instead.
func (*GetTaskRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_plugin_job_proto_rawDescGZIP(), []int{1}
}

func (x *GetTaskRequest) GetSecretData() *_struct.Struct {
	if x != nil {
		return x.SecretData
	}
	return nil
}

func (x *GetTaskRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *GetTaskRequest) GetTaskFilter() *TaskFilter {
	if x != nil {
		return x.TaskFilter
	}
	return nil
}

type TaskInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskOptions   *_struct.Struct        `protobuf:"bytes,1,opt,name=task_options,json=taskOptions,proto3" json:"task_options,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskInfo) Reset() {
	*x = TaskInfo{}
	mi := &file_spaceone_api_inventory_plugin_job_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskInfo) ProtoMessage() {}

func (x *TaskInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_plugin_job_proto_msgTypes[2]
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
	return file_spaceone_api_inventory_plugin_job_proto_rawDescGZIP(), []int{2}
}

func (x *TaskInfo) GetTaskOptions() *_struct.Struct {
	if x != nil {
		return x.TaskOptions
	}
	return nil
}

type TasksInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tasks         []*TaskInfo            `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TasksInfo) Reset() {
	*x = TasksInfo{}
	mi := &file_spaceone_api_inventory_plugin_job_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TasksInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TasksInfo) ProtoMessage() {}

func (x *TasksInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_plugin_job_proto_msgTypes[3]
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
	return file_spaceone_api_inventory_plugin_job_proto_rawDescGZIP(), []int{3}
}

func (x *TasksInfo) GetTasks() []*TaskInfo {
	if x != nil {
		return x.Tasks
	}
	return nil
}

var File_spaceone_api_inventory_plugin_job_proto protoreflect.FileDescriptor

var file_spaceone_api_inventory_plugin_job_proto_rawDesc = []byte{
	0x0a, 0x27, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f,
	0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x01, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73,
	0x12, 0x30, 0x0a, 0x14, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x11, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x22, 0xc9, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x4a, 0x0a, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x52, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22,
	0x46, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3a, 0x0a, 0x0c, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0b, 0x74, 0x61, 0x73, 0x6b,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x4a, 0x0a, 0x09, 0x54, 0x61, 0x73, 0x6b, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3d, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x32, 0x6d, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x66, 0x0a, 0x09, 0x67, 0x65,
	0x74, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x2d, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f,
	0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x00, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_inventory_plugin_job_proto_rawDescOnce sync.Once
	file_spaceone_api_inventory_plugin_job_proto_rawDescData = file_spaceone_api_inventory_plugin_job_proto_rawDesc
)

func file_spaceone_api_inventory_plugin_job_proto_rawDescGZIP() []byte {
	file_spaceone_api_inventory_plugin_job_proto_rawDescOnce.Do(func() {
		file_spaceone_api_inventory_plugin_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_inventory_plugin_job_proto_rawDescData)
	})
	return file_spaceone_api_inventory_plugin_job_proto_rawDescData
}

var file_spaceone_api_inventory_plugin_job_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_spaceone_api_inventory_plugin_job_proto_goTypes = []any{
	(*TaskFilter)(nil),     // 0: spaceone.api.inventory.plugin.TaskFilter
	(*GetTaskRequest)(nil), // 1: spaceone.api.inventory.plugin.GetTaskRequest
	(*TaskInfo)(nil),       // 2: spaceone.api.inventory.plugin.TaskInfo
	(*TasksInfo)(nil),      // 3: spaceone.api.inventory.plugin.TasksInfo
	(*_struct.Struct)(nil), // 4: google.protobuf.Struct
}
var file_spaceone_api_inventory_plugin_job_proto_depIdxs = []int32{
	4, // 0: spaceone.api.inventory.plugin.GetTaskRequest.secret_data:type_name -> google.protobuf.Struct
	4, // 1: spaceone.api.inventory.plugin.GetTaskRequest.options:type_name -> google.protobuf.Struct
	0, // 2: spaceone.api.inventory.plugin.GetTaskRequest.task_filter:type_name -> spaceone.api.inventory.plugin.TaskFilter
	4, // 3: spaceone.api.inventory.plugin.TaskInfo.task_options:type_name -> google.protobuf.Struct
	2, // 4: spaceone.api.inventory.plugin.TasksInfo.tasks:type_name -> spaceone.api.inventory.plugin.TaskInfo
	1, // 5: spaceone.api.inventory.plugin.Job.get_tasks:input_type -> spaceone.api.inventory.plugin.GetTaskRequest
	3, // 6: spaceone.api.inventory.plugin.Job.get_tasks:output_type -> spaceone.api.inventory.plugin.TasksInfo
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_spaceone_api_inventory_plugin_job_proto_init() }
func file_spaceone_api_inventory_plugin_job_proto_init() {
	if File_spaceone_api_inventory_plugin_job_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spaceone_api_inventory_plugin_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_inventory_plugin_job_proto_goTypes,
		DependencyIndexes: file_spaceone_api_inventory_plugin_job_proto_depIdxs,
		MessageInfos:      file_spaceone_api_inventory_plugin_job_proto_msgTypes,
	}.Build()
	File_spaceone_api_inventory_plugin_job_proto = out.File
	file_spaceone_api_inventory_plugin_job_proto_rawDesc = nil
	file_spaceone_api_inventory_plugin_job_proto_goTypes = nil
	file_spaceone_api_inventory_plugin_job_proto_depIdxs = nil
}
