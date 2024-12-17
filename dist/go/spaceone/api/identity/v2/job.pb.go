// A Job is an act of collecting external cloud resources through plugins.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v3.6.1
// source: spaceone/api/identity/v2/job.proto

package v2

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

type JobsQuery_JobStatus int32

const (
	JobsQuery_NONE        JobsQuery_JobStatus = 0
	JobsQuery_PENDING     JobsQuery_JobStatus = 1
	JobsQuery_IN_PROGRESS JobsQuery_JobStatus = 2
	JobsQuery_FAILURE     JobsQuery_JobStatus = 3
	JobsQuery_SUCCESS     JobsQuery_JobStatus = 4
	JobsQuery_CANCELED    JobsQuery_JobStatus = 5
)

// Enum value maps for JobsQuery_JobStatus.
var (
	JobsQuery_JobStatus_name = map[int32]string{
		0: "NONE",
		1: "PENDING",
		2: "IN_PROGRESS",
		3: "FAILURE",
		4: "SUCCESS",
		5: "CANCELED",
	}
	JobsQuery_JobStatus_value = map[string]int32{
		"NONE":        0,
		"PENDING":     1,
		"IN_PROGRESS": 2,
		"FAILURE":     3,
		"SUCCESS":     4,
		"CANCELED":    5,
	}
)

func (x JobsQuery_JobStatus) Enum() *JobsQuery_JobStatus {
	p := new(JobsQuery_JobStatus)
	*p = x
	return p
}

func (x JobsQuery_JobStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobsQuery_JobStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_identity_v2_job_proto_enumTypes[0].Descriptor()
}

func (JobsQuery_JobStatus) Type() protoreflect.EnumType {
	return &file_spaceone_api_identity_v2_job_proto_enumTypes[0]
}

func (x JobsQuery_JobStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobsQuery_JobStatus.Descriptor instead.
func (JobsQuery_JobStatus) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_job_proto_rawDescGZIP(), []int{1, 0}
}

type JobInfo_Status int32

const (
	JobInfo_NONE        JobInfo_Status = 0
	JobInfo_PENDING     JobInfo_Status = 1
	JobInfo_IN_PROGRESS JobInfo_Status = 2
	JobInfo_FAILURE     JobInfo_Status = 3
	JobInfo_SUCCESS     JobInfo_Status = 4
	JobInfo_CANCELED    JobInfo_Status = 5
)

// Enum value maps for JobInfo_Status.
var (
	JobInfo_Status_name = map[int32]string{
		0: "NONE",
		1: "PENDING",
		2: "IN_PROGRESS",
		3: "FAILURE",
		4: "SUCCESS",
		5: "CANCELED",
	}
	JobInfo_Status_value = map[string]int32{
		"NONE":        0,
		"PENDING":     1,
		"IN_PROGRESS": 2,
		"FAILURE":     3,
		"SUCCESS":     4,
		"CANCELED":    5,
	}
)

func (x JobInfo_Status) Enum() *JobInfo_Status {
	p := new(JobInfo_Status)
	*p = x
	return p
}

func (x JobInfo_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobInfo_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_identity_v2_job_proto_enumTypes[1].Descriptor()
}

func (JobInfo_Status) Type() protoreflect.EnumType {
	return &file_spaceone_api_identity_v2_job_proto_enumTypes[1]
}

func (x JobInfo_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobInfo_Status.Descriptor instead.
func (JobInfo_Status) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_job_proto_rawDescGZIP(), []int{2, 0}
}

type JobInfo_ResourceGroup int32

const (
	JobInfo_RESOURCE_GROUP_NONE JobInfo_ResourceGroup = 0
	JobInfo_DOMAIN              JobInfo_ResourceGroup = 1
	JobInfo_WORKSPACE           JobInfo_ResourceGroup = 2
)

// Enum value maps for JobInfo_ResourceGroup.
var (
	JobInfo_ResourceGroup_name = map[int32]string{
		0: "RESOURCE_GROUP_NONE",
		1: "DOMAIN",
		2: "WORKSPACE",
	}
	JobInfo_ResourceGroup_value = map[string]int32{
		"RESOURCE_GROUP_NONE": 0,
		"DOMAIN":              1,
		"WORKSPACE":           2,
	}
)

func (x JobInfo_ResourceGroup) Enum() *JobInfo_ResourceGroup {
	p := new(JobInfo_ResourceGroup)
	*p = x
	return p
}

func (x JobInfo_ResourceGroup) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobInfo_ResourceGroup) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_identity_v2_job_proto_enumTypes[2].Descriptor()
}

func (JobInfo_ResourceGroup) Type() protoreflect.EnumType {
	return &file_spaceone_api_identity_v2_job_proto_enumTypes[2]
}

func (x JobInfo_ResourceGroup) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobInfo_ResourceGroup.Descriptor instead.
func (JobInfo_ResourceGroup) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_job_proto_rawDescGZIP(), []int{2, 1}
}

//	{
//	   "job_id": "job-123456789012"
//	}
type JobRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	JobId         string                 `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobRequest) Reset() {
	*x = JobRequest{}
	mi := &file_spaceone_api_identity_v2_job_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobRequest) ProtoMessage() {}

func (x *JobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_job_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobRequest.ProtoReflect.Descriptor instead.
func (*JobRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_job_proto_rawDescGZIP(), []int{0}
}

func (x *JobRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

//	{
//	   "query": {
//	       "page": {
//	           "start": 1,
//	           "limit": 10
//	       }
//	   },
//	   "trusted_account_id": "ta-123456789012"
//	}
type JobsQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// +optional
	Query *v2.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	JobId string `protobuf:"bytes,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	// +optional
	Status JobsQuery_JobStatus `protobuf:"varint,3,opt,name=status,proto3,enum=spaceone.api.identity.v2.JobsQuery_JobStatus" json:"status,omitempty"`
	// +optional
	WorkspaceId string `protobuf:"bytes,21,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	// +optional
	TrustedAccountId string `protobuf:"bytes,22,opt,name=trusted_account_id,json=trustedAccountId,proto3" json:"trusted_account_id,omitempty"`
	// +optional
	PluginId      string `protobuf:"bytes,23,opt,name=plugin_id,json=pluginId,proto3" json:"plugin_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobsQuery) Reset() {
	*x = JobsQuery{}
	mi := &file_spaceone_api_identity_v2_job_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobsQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobsQuery) ProtoMessage() {}

func (x *JobsQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_job_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobsQuery.ProtoReflect.Descriptor instead.
func (*JobsQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_job_proto_rawDescGZIP(), []int{1}
}

func (x *JobsQuery) GetQuery() *v2.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *JobsQuery) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *JobsQuery) GetStatus() JobsQuery_JobStatus {
	if x != nil {
		return x.Status
	}
	return JobsQuery_NONE
}

func (x *JobsQuery) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *JobsQuery) GetTrustedAccountId() string {
	if x != nil {
		return x.TrustedAccountId
	}
	return ""
}

func (x *JobsQuery) GetPluginId() string {
	if x != nil {
		return x.PluginId
	}
	return ""
}

//	{
//	     "job_id": "job-123456789012",
//	     "status": "ERROR",
//	     "resource_group": "DOMAIN",
//	     "plugin_id": "plugin-aws-cloud-service-inven-collector",
//	     "trusted_account_id": "ta-123456789012",
//	     "domain_id": "domain-123456789012",
//	     "created_at": "2022-01-01T10:00:01.389Z",
//	     "updated_at": "2022-01-01T10:00:01.389Z",
//	     "finished_at": "2022-01-01T10:02:11.270Z"
//	}
type JobInfo struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	JobId            string                 `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Status           JobInfo_Status         `protobuf:"varint,2,opt,name=status,proto3,enum=spaceone.api.identity.v2.JobInfo_Status" json:"status,omitempty"`
	Options          *_struct.Struct        `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
	ResourceGroup    JobInfo_ResourceGroup  `protobuf:"varint,4,opt,name=resource_group,json=resourceGroup,proto3,enum=spaceone.api.identity.v2.JobInfo_ResourceGroup" json:"resource_group,omitempty"`
	ErrorMessage     string                 `protobuf:"bytes,5,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	DomainId         string                 `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	WorkspaceId      string                 `protobuf:"bytes,22,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	TrustedAccountId string                 `protobuf:"bytes,23,opt,name=trusted_account_id,json=trustedAccountId,proto3" json:"trusted_account_id,omitempty"`
	PluginId         string                 `protobuf:"bytes,24,opt,name=plugin_id,json=pluginId,proto3" json:"plugin_id,omitempty"`
	CreatedAt        string                 `protobuf:"bytes,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt        string                 `protobuf:"bytes,32,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	FinishedAt       string                 `protobuf:"bytes,33,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *JobInfo) Reset() {
	*x = JobInfo{}
	mi := &file_spaceone_api_identity_v2_job_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobInfo) ProtoMessage() {}

func (x *JobInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_job_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobInfo.ProtoReflect.Descriptor instead.
func (*JobInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_job_proto_rawDescGZIP(), []int{2}
}

func (x *JobInfo) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *JobInfo) GetStatus() JobInfo_Status {
	if x != nil {
		return x.Status
	}
	return JobInfo_NONE
}

func (x *JobInfo) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *JobInfo) GetResourceGroup() JobInfo_ResourceGroup {
	if x != nil {
		return x.ResourceGroup
	}
	return JobInfo_RESOURCE_GROUP_NONE
}

func (x *JobInfo) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *JobInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *JobInfo) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *JobInfo) GetTrustedAccountId() string {
	if x != nil {
		return x.TrustedAccountId
	}
	return ""
}

func (x *JobInfo) GetPluginId() string {
	if x != nil {
		return x.PluginId
	}
	return ""
}

func (x *JobInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *JobInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *JobInfo) GetFinishedAt() string {
	if x != nil {
		return x.FinishedAt
	}
	return ""
}

//	{
//	   "results": [
//	       {
//	           "job_id": "job-3b124006c2d2",
//	           "status": "SUCCESS",
//	           "resource_group": "DOMAIN",
//	           "plugin_id": "plugin-aws-cloud-service-inven-collector",
//	           "trusted_account_id": "ta-3b124006c2d2",
//	           "domain_id": "domain-58010aa2e451",
//	           "created_at": "2022-06-17T08:00:01.225Z",
//	           "updated_at": "2022-06-17T08:00:01.225Z",
//	           "finished_at": "2022-06-17T08:00:15.197Z"
//	       },
//	       {
//	           "job_id": "job-587a3d3b4db3",
//	           "status": "SUCCESS",
//	           "resource_group": "DOMAIN",
//	           "plugin_id": "plugin-aws-cloud-service-inven-collector",
//	           "trusted_account_id": "ta-587a3d3b4db3",
//	           "domain_id": "domain-58010aa2e451",
//	           "created_at": "2022-06-17T08:00:00.407Z",
//	           "updated_at": "2022-06-17T08:00:00.407Z",
//	           "finished_at": "2022-06-17T08:07:32.023Z"
//	       }
//	   ],
//	   "total_count": 2
//	}
type JobsInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Results       []*JobInfo             `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount    int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobsInfo) Reset() {
	*x = JobsInfo{}
	mi := &file_spaceone_api_identity_v2_job_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobsInfo) ProtoMessage() {}

func (x *JobsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_job_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobsInfo.ProtoReflect.Descriptor instead.
func (*JobsInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_job_proto_rawDescGZIP(), []int{3}
}

func (x *JobsInfo) GetResults() []*JobInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *JobsInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type JobStatQuery struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         *v2.StatisticsQuery    `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobStatQuery) Reset() {
	*x = JobStatQuery{}
	mi := &file_spaceone_api_identity_v2_job_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobStatQuery) ProtoMessage() {}

func (x *JobStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_job_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobStatQuery.ProtoReflect.Descriptor instead.
func (*JobStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_job_proto_rawDescGZIP(), []int{4}
}

func (x *JobStatQuery) GetQuery() *v2.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

var File_spaceone_api_identity_v2_job_proto protoreflect.FileDescriptor

var file_spaceone_api_identity_v2_job_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x0a, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x22, 0xe7,
	0x02, 0x0a, 0x09, 0x4a, 0x6f, 0x62, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x31, 0x0a, 0x05,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x32, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x45, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76,
	0x32, 0x2e, 0x4a, 0x6f, 0x62, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4a, 0x6f, 0x62, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a,
	0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x2c, 0x0a, 0x12, 0x74, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x72,
	0x75, 0x73, 0x74, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x5b, 0x0a, 0x09, 0x4a,
	0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12,
	0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x02,
	0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x03, 0x12, 0x0b, 0x0a,
	0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41,
	0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x05, 0x22, 0x9b, 0x05, 0x0a, 0x07, 0x4a, 0x6f, 0x62,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a,
	0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x56, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x76, 0x32, 0x2e, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2c, 0x0a,
	0x12, 0x74, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x72, 0x75, 0x73, 0x74,
	0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x21, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x22, 0x58, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50,
	0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x5f, 0x50,
	0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49,
	0x4c, 0x55, 0x52, 0x45, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10,
	0x05, 0x22, 0x43, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x47,
	0x52, 0x4f, 0x55, 0x50, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x44,
	0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x57, 0x4f, 0x52, 0x4b, 0x53,
	0x50, 0x41, 0x43, 0x45, 0x10, 0x02, 0x22, 0x68, 0x0a, 0x08, 0x4a, 0x6f, 0x62, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x3b, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x4a,
	0x6f, 0x62, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x4b, 0x0a, 0x0c, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x3b, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x32, 0xc0, 0x03,
	0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x6a, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x24, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x22, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x6f, 0x0a, 0x03, 0x67, 0x65, 0x74, 0x12, 0x24, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x76, 0x32, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x67,
	0x65, 0x74, 0x12, 0x71, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x23, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x4a, 0x6f, 0x62, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a,
	0x22, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x4a, 0x6f, 0x62, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15,
	0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x6a, 0x6f, 0x62,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x69, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x74, 0x12, 0x26, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x20,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76,
	0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_identity_v2_job_proto_rawDescOnce sync.Once
	file_spaceone_api_identity_v2_job_proto_rawDescData = file_spaceone_api_identity_v2_job_proto_rawDesc
)

func file_spaceone_api_identity_v2_job_proto_rawDescGZIP() []byte {
	file_spaceone_api_identity_v2_job_proto_rawDescOnce.Do(func() {
		file_spaceone_api_identity_v2_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_identity_v2_job_proto_rawDescData)
	})
	return file_spaceone_api_identity_v2_job_proto_rawDescData
}

var file_spaceone_api_identity_v2_job_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_spaceone_api_identity_v2_job_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_spaceone_api_identity_v2_job_proto_goTypes = []any{
	(JobsQuery_JobStatus)(0),   // 0: spaceone.api.identity.v2.JobsQuery.JobStatus
	(JobInfo_Status)(0),        // 1: spaceone.api.identity.v2.JobInfo.Status
	(JobInfo_ResourceGroup)(0), // 2: spaceone.api.identity.v2.JobInfo.ResourceGroup
	(*JobRequest)(nil),         // 3: spaceone.api.identity.v2.JobRequest
	(*JobsQuery)(nil),          // 4: spaceone.api.identity.v2.JobsQuery
	(*JobInfo)(nil),            // 5: spaceone.api.identity.v2.JobInfo
	(*JobsInfo)(nil),           // 6: spaceone.api.identity.v2.JobsInfo
	(*JobStatQuery)(nil),       // 7: spaceone.api.identity.v2.JobStatQuery
	(*v2.Query)(nil),           // 8: spaceone.api.core.v2.Query
	(*_struct.Struct)(nil),     // 9: google.protobuf.Struct
	(*v2.StatisticsQuery)(nil), // 10: spaceone.api.core.v2.StatisticsQuery
	(*empty.Empty)(nil),        // 11: google.protobuf.Empty
}
var file_spaceone_api_identity_v2_job_proto_depIdxs = []int32{
	8,  // 0: spaceone.api.identity.v2.JobsQuery.query:type_name -> spaceone.api.core.v2.Query
	0,  // 1: spaceone.api.identity.v2.JobsQuery.status:type_name -> spaceone.api.identity.v2.JobsQuery.JobStatus
	1,  // 2: spaceone.api.identity.v2.JobInfo.status:type_name -> spaceone.api.identity.v2.JobInfo.Status
	9,  // 3: spaceone.api.identity.v2.JobInfo.options:type_name -> google.protobuf.Struct
	2,  // 4: spaceone.api.identity.v2.JobInfo.resource_group:type_name -> spaceone.api.identity.v2.JobInfo.ResourceGroup
	5,  // 5: spaceone.api.identity.v2.JobsInfo.results:type_name -> spaceone.api.identity.v2.JobInfo
	10, // 6: spaceone.api.identity.v2.JobStatQuery.query:type_name -> spaceone.api.core.v2.StatisticsQuery
	3,  // 7: spaceone.api.identity.v2.Job.delete:input_type -> spaceone.api.identity.v2.JobRequest
	3,  // 8: spaceone.api.identity.v2.Job.get:input_type -> spaceone.api.identity.v2.JobRequest
	4,  // 9: spaceone.api.identity.v2.Job.list:input_type -> spaceone.api.identity.v2.JobsQuery
	7,  // 10: spaceone.api.identity.v2.Job.stat:input_type -> spaceone.api.identity.v2.JobStatQuery
	11, // 11: spaceone.api.identity.v2.Job.delete:output_type -> google.protobuf.Empty
	5,  // 12: spaceone.api.identity.v2.Job.get:output_type -> spaceone.api.identity.v2.JobInfo
	6,  // 13: spaceone.api.identity.v2.Job.list:output_type -> spaceone.api.identity.v2.JobsInfo
	9,  // 14: spaceone.api.identity.v2.Job.stat:output_type -> google.protobuf.Struct
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_spaceone_api_identity_v2_job_proto_init() }
func file_spaceone_api_identity_v2_job_proto_init() {
	if File_spaceone_api_identity_v2_job_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spaceone_api_identity_v2_job_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_identity_v2_job_proto_goTypes,
		DependencyIndexes: file_spaceone_api_identity_v2_job_proto_depIdxs,
		EnumInfos:         file_spaceone_api_identity_v2_job_proto_enumTypes,
		MessageInfos:      file_spaceone_api_identity_v2_job_proto_msgTypes,
	}.Build()
	File_spaceone_api_identity_v2_job_proto = out.File
	file_spaceone_api_identity_v2_job_proto_rawDesc = nil
	file_spaceone_api_identity_v2_job_proto_goTypes = nil
	file_spaceone_api_identity_v2_job_proto_depIdxs = nil
}
