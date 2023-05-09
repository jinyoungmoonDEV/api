// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.6.1
// source: spaceone/api/cost_analysis/v1/budget_usage.proto

package v1

import (
	v1 "github.com/cloudforet-io/api/dist/go/spaceone/api/core/v1"
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

//	{
//	   "query": {}
//	}
type BudgetUsageQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// +optional
	Query *v1.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	BudgetId string `protobuf:"bytes,2,opt,name=budget_id,json=budgetId,proto3" json:"budget_id,omitempty"`
	// +optional
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Date     string `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	DomainId string `protobuf:"bytes,11,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *BudgetUsageQuery) Reset() {
	*x = BudgetUsageQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_cost_analysis_v1_budget_usage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BudgetUsageQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BudgetUsageQuery) ProtoMessage() {}

func (x *BudgetUsageQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_budget_usage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BudgetUsageQuery.ProtoReflect.Descriptor instead.
func (*BudgetUsageQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDescGZIP(), []int{0}
}

func (x *BudgetUsageQuery) GetQuery() *v1.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *BudgetUsageQuery) GetBudgetId() string {
	if x != nil {
		return x.BudgetId
	}
	return ""
}

func (x *BudgetUsageQuery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BudgetUsageQuery) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *BudgetUsageQuery) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type BudgetUsageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BudgetId  string          `protobuf:"bytes,1,opt,name=budget_id,json=budgetId,proto3" json:"budget_id,omitempty"`
	Name      string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Date      string          `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	UsdCost   float32         `protobuf:"fixed32,5,opt,name=usd_cost,json=usdCost,proto3" json:"usd_cost,omitempty"`
	Limit     float32         `protobuf:"fixed32,6,opt,name=limit,proto3" json:"limit,omitempty"`
	CostTypes *_struct.Struct `protobuf:"bytes,7,opt,name=cost_types,json=costTypes,proto3" json:"cost_types,omitempty"`
	DomainId  string          `protobuf:"bytes,11,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	UpdatedAt string          `protobuf:"bytes,21,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *BudgetUsageInfo) Reset() {
	*x = BudgetUsageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_cost_analysis_v1_budget_usage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BudgetUsageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BudgetUsageInfo) ProtoMessage() {}

func (x *BudgetUsageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_budget_usage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BudgetUsageInfo.ProtoReflect.Descriptor instead.
func (*BudgetUsageInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDescGZIP(), []int{1}
}

func (x *BudgetUsageInfo) GetBudgetId() string {
	if x != nil {
		return x.BudgetId
	}
	return ""
}

func (x *BudgetUsageInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BudgetUsageInfo) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *BudgetUsageInfo) GetUsdCost() float32 {
	if x != nil {
		return x.UsdCost
	}
	return 0
}

func (x *BudgetUsageInfo) GetLimit() float32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *BudgetUsageInfo) GetCostTypes() *_struct.Struct {
	if x != nil {
		return x.CostTypes
	}
	return nil
}

func (x *BudgetUsageInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *BudgetUsageInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

//	{
//	       "results": [
//	           {
//	               "budget_id": "budget-abb377eb9e8b",
//	               "name": "Cloudforet-Budget3",
//	               "date": "2022-01",
//	               "usd_cost": 7671.164,
//	               "limit": 10000.0,
//	               "domain_id": "domain-58010aa2e451",
//	               "updated_at": "2022-07-19T04:26:08.099Z"
//	           },
//	           {
//	               "budget_id": "budget-abb377eb9e8b",
//	               "name": "Cloudforet-Budget3",
//	               "date": "2022-02",
//	               "usd_cost": 5931.771,
//	               "limit": 11000.0,
//	               "domain_id": "domain-58010aa2e451",
//	               "updated_at": "2022-07-19T04:26:08.105Z"
//	           }
//	       ],
//	       "total_count": 12
//	}
type BudgetUsagesInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results    []*BudgetUsageInfo `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount int32              `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *BudgetUsagesInfo) Reset() {
	*x = BudgetUsagesInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_cost_analysis_v1_budget_usage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BudgetUsagesInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BudgetUsagesInfo) ProtoMessage() {}

func (x *BudgetUsagesInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_budget_usage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BudgetUsagesInfo.ProtoReflect.Descriptor instead.
func (*BudgetUsagesInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDescGZIP(), []int{2}
}

func (x *BudgetUsagesInfo) GetResults() []*BudgetUsageInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *BudgetUsagesInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type BudgetUsageAnalyzeQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query    *v1.TimeSeriesAnalyzeQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	DomainId string                     `protobuf:"bytes,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *BudgetUsageAnalyzeQuery) Reset() {
	*x = BudgetUsageAnalyzeQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_cost_analysis_v1_budget_usage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BudgetUsageAnalyzeQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BudgetUsageAnalyzeQuery) ProtoMessage() {}

func (x *BudgetUsageAnalyzeQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_budget_usage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BudgetUsageAnalyzeQuery.ProtoReflect.Descriptor instead.
func (*BudgetUsageAnalyzeQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDescGZIP(), []int{3}
}

func (x *BudgetUsageAnalyzeQuery) GetQuery() *v1.TimeSeriesAnalyzeQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *BudgetUsageAnalyzeQuery) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type BudgetUsageStatQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query    *v1.StatisticsQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	DomainId string              `protobuf:"bytes,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *BudgetUsageStatQuery) Reset() {
	*x = BudgetUsageStatQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_cost_analysis_v1_budget_usage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BudgetUsageStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BudgetUsageStatQuery) ProtoMessage() {}

func (x *BudgetUsageStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_budget_usage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BudgetUsageStatQuery.ProtoReflect.Descriptor instead.
func (*BudgetUsageStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDescGZIP(), []int{4}
}

func (x *BudgetUsageStatQuery) GetQuery() *v1.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *BudgetUsageStatQuery) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

var File_spaceone_api_cost_analysis_v1_budget_usage_proto protoreflect.FileDescriptor

var file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDesc = []byte{
	0x0a, 0x30, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1d, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa7, 0x01, 0x0a, 0x10, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x75, 0x64, 0x67,
	0x65, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0xfb, 0x01, 0x0a, 0x0f, 0x42, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a,
	0x09, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x64, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x75, 0x73, 0x64, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x36, 0x0a, 0x0a, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x7d, 0x0a, 0x10, 0x42, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x48, 0x0a, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74,
	0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x7a, 0x0a, 0x17, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x55, 0x73, 0x61, 0x67, 0x65, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x42, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x49, 0x64, 0x22, 0x70, 0x0a, 0x14, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x3b, 0x0a, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x49, 0x64, 0x32, 0xbf, 0x03, 0x0a, 0x0b, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x98, 0x01, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73,
	0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x2f,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x3a, 0x01, 0x2a, 0x22, 0x23, 0x2f, 0x63, 0x6f, 0x73,
	0x74, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x2d, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x8d, 0x01, 0x0a, 0x07, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x12, 0x36, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x64, 0x67,
	0x65, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x31, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x2b, 0x3a, 0x01, 0x2a, 0x22, 0x26, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2d, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x2d, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x12,
	0x84, 0x01, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x74, 0x12, 0x33, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x3a, 0x01,
	0x2a, 0x22, 0x23, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2d, 0x75, 0x73, 0x61, 0x67,
	0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d,
	0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x74,
	0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDescOnce sync.Once
	file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDescData = file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDesc
)

func file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDescGZIP() []byte {
	file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDescOnce.Do(func() {
		file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDescData)
	})
	return file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDescData
}

var file_spaceone_api_cost_analysis_v1_budget_usage_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_spaceone_api_cost_analysis_v1_budget_usage_proto_goTypes = []interface{}{
	(*BudgetUsageQuery)(nil),          // 0: spaceone.api.cost_analysis.v1.BudgetUsageQuery
	(*BudgetUsageInfo)(nil),           // 1: spaceone.api.cost_analysis.v1.BudgetUsageInfo
	(*BudgetUsagesInfo)(nil),          // 2: spaceone.api.cost_analysis.v1.BudgetUsagesInfo
	(*BudgetUsageAnalyzeQuery)(nil),   // 3: spaceone.api.cost_analysis.v1.BudgetUsageAnalyzeQuery
	(*BudgetUsageStatQuery)(nil),      // 4: spaceone.api.cost_analysis.v1.BudgetUsageStatQuery
	(*v1.Query)(nil),                  // 5: spaceone.api.core.v1.Query
	(*_struct.Struct)(nil),            // 6: google.protobuf.Struct
	(*v1.TimeSeriesAnalyzeQuery)(nil), // 7: spaceone.api.core.v1.TimeSeriesAnalyzeQuery
	(*v1.StatisticsQuery)(nil),        // 8: spaceone.api.core.v1.StatisticsQuery
}
var file_spaceone_api_cost_analysis_v1_budget_usage_proto_depIdxs = []int32{
	5, // 0: spaceone.api.cost_analysis.v1.BudgetUsageQuery.query:type_name -> spaceone.api.core.v1.Query
	6, // 1: spaceone.api.cost_analysis.v1.BudgetUsageInfo.cost_types:type_name -> google.protobuf.Struct
	1, // 2: spaceone.api.cost_analysis.v1.BudgetUsagesInfo.results:type_name -> spaceone.api.cost_analysis.v1.BudgetUsageInfo
	7, // 3: spaceone.api.cost_analysis.v1.BudgetUsageAnalyzeQuery.query:type_name -> spaceone.api.core.v1.TimeSeriesAnalyzeQuery
	8, // 4: spaceone.api.cost_analysis.v1.BudgetUsageStatQuery.query:type_name -> spaceone.api.core.v1.StatisticsQuery
	0, // 5: spaceone.api.cost_analysis.v1.BudgetUsage.list:input_type -> spaceone.api.cost_analysis.v1.BudgetUsageQuery
	3, // 6: spaceone.api.cost_analysis.v1.BudgetUsage.analyze:input_type -> spaceone.api.cost_analysis.v1.BudgetUsageAnalyzeQuery
	4, // 7: spaceone.api.cost_analysis.v1.BudgetUsage.stat:input_type -> spaceone.api.cost_analysis.v1.BudgetUsageStatQuery
	2, // 8: spaceone.api.cost_analysis.v1.BudgetUsage.list:output_type -> spaceone.api.cost_analysis.v1.BudgetUsagesInfo
	6, // 9: spaceone.api.cost_analysis.v1.BudgetUsage.analyze:output_type -> google.protobuf.Struct
	6, // 10: spaceone.api.cost_analysis.v1.BudgetUsage.stat:output_type -> google.protobuf.Struct
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_spaceone_api_cost_analysis_v1_budget_usage_proto_init() }
func file_spaceone_api_cost_analysis_v1_budget_usage_proto_init() {
	if File_spaceone_api_cost_analysis_v1_budget_usage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spaceone_api_cost_analysis_v1_budget_usage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BudgetUsageQuery); i {
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
		file_spaceone_api_cost_analysis_v1_budget_usage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BudgetUsageInfo); i {
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
		file_spaceone_api_cost_analysis_v1_budget_usage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BudgetUsagesInfo); i {
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
		file_spaceone_api_cost_analysis_v1_budget_usage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BudgetUsageAnalyzeQuery); i {
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
		file_spaceone_api_cost_analysis_v1_budget_usage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BudgetUsageStatQuery); i {
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
			RawDescriptor: file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_cost_analysis_v1_budget_usage_proto_goTypes,
		DependencyIndexes: file_spaceone_api_cost_analysis_v1_budget_usage_proto_depIdxs,
		MessageInfos:      file_spaceone_api_cost_analysis_v1_budget_usage_proto_msgTypes,
	}.Build()
	File_spaceone_api_cost_analysis_v1_budget_usage_proto = out.File
	file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDesc = nil
	file_spaceone_api_cost_analysis_v1_budget_usage_proto_goTypes = nil
	file_spaceone_api_cost_analysis_v1_budget_usage_proto_depIdxs = nil
}
