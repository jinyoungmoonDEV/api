// A Resource is a resource used for analysis on all microservices used in Cloudforet.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v3.6.1
// source: spaceone/api/statistics/v1/resource.proto

package v1

import (
	v2 "github.com/cloudforet-io/api/dist/go/spaceone/api/core/v2"
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

type StatAggregateJoin_JoinType int32

const (
	StatAggregateJoin_LEFT  StatAggregateJoin_JoinType = 0
	StatAggregateJoin_RIGHT StatAggregateJoin_JoinType = 1
	StatAggregateJoin_OUTER StatAggregateJoin_JoinType = 2
	StatAggregateJoin_INNER StatAggregateJoin_JoinType = 3
)

// Enum value maps for StatAggregateJoin_JoinType.
var (
	StatAggregateJoin_JoinType_name = map[int32]string{
		0: "LEFT",
		1: "RIGHT",
		2: "OUTER",
		3: "INNER",
	}
	StatAggregateJoin_JoinType_value = map[string]int32{
		"LEFT":  0,
		"RIGHT": 1,
		"OUTER": 2,
		"INNER": 3,
	}
)

func (x StatAggregateJoin_JoinType) Enum() *StatAggregateJoin_JoinType {
	p := new(StatAggregateJoin_JoinType)
	*p = x
	return p
}

func (x StatAggregateJoin_JoinType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatAggregateJoin_JoinType) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_statistics_v1_resource_proto_enumTypes[0].Descriptor()
}

func (StatAggregateJoin_JoinType) Type() protoreflect.EnumType {
	return &file_spaceone_api_statistics_v1_resource_proto_enumTypes[0]
}

func (x StatAggregateJoin_JoinType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatAggregateJoin_JoinType.Descriptor instead.
func (StatAggregateJoin_JoinType) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_statistics_v1_resource_proto_rawDescGZIP(), []int{1, 0}
}

type StatAggregateQuery struct {
	state        protoimpl.MessageState `protogen:"open.v1"`
	ResourceType string                 `protobuf:"bytes,1,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	Query        *v2.StatisticsQuery    `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	ExtendData    *_struct.Struct `protobuf:"bytes,3,opt,name=extend_data,json=extendData,proto3" json:"extend_data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StatAggregateQuery) Reset() {
	*x = StatAggregateQuery{}
	mi := &file_spaceone_api_statistics_v1_resource_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatAggregateQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatAggregateQuery) ProtoMessage() {}

func (x *StatAggregateQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_statistics_v1_resource_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatAggregateQuery.ProtoReflect.Descriptor instead.
func (*StatAggregateQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_statistics_v1_resource_proto_rawDescGZIP(), []int{0}
}

func (x *StatAggregateQuery) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *StatAggregateQuery) GetQuery() *v2.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *StatAggregateQuery) GetExtendData() *_struct.Struct {
	if x != nil {
		return x.ExtendData
	}
	return nil
}

type StatAggregateJoin struct {
	state        protoimpl.MessageState `protogen:"open.v1"`
	ResourceType string                 `protobuf:"bytes,1,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	Query        *v2.StatisticsQuery    `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	ExtendData *_struct.Struct `protobuf:"bytes,3,opt,name=extend_data,json=extendData,proto3" json:"extend_data,omitempty"`
	// +optional
	Type StatAggregateJoin_JoinType `protobuf:"varint,4,opt,name=type,proto3,enum=spaceone.api.statistics.v1.StatAggregateJoin_JoinType" json:"type,omitempty"`
	// +optional
	Keys          []string `protobuf:"bytes,5,rep,name=keys,proto3" json:"keys,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StatAggregateJoin) Reset() {
	*x = StatAggregateJoin{}
	mi := &file_spaceone_api_statistics_v1_resource_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatAggregateJoin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatAggregateJoin) ProtoMessage() {}

func (x *StatAggregateJoin) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_statistics_v1_resource_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatAggregateJoin.ProtoReflect.Descriptor instead.
func (*StatAggregateJoin) Descriptor() ([]byte, []int) {
	return file_spaceone_api_statistics_v1_resource_proto_rawDescGZIP(), []int{1}
}

func (x *StatAggregateJoin) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *StatAggregateJoin) GetQuery() *v2.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *StatAggregateJoin) GetExtendData() *_struct.Struct {
	if x != nil {
		return x.ExtendData
	}
	return nil
}

func (x *StatAggregateJoin) GetType() StatAggregateJoin_JoinType {
	if x != nil {
		return x.Type
	}
	return StatAggregateJoin_LEFT
}

func (x *StatAggregateJoin) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

type StatAggregateConcat struct {
	state        protoimpl.MessageState `protogen:"open.v1"`
	ResourceType string                 `protobuf:"bytes,1,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	Query        *v2.StatisticsQuery    `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	ExtendData    *_struct.Struct `protobuf:"bytes,3,opt,name=extend_data,json=extendData,proto3" json:"extend_data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StatAggregateConcat) Reset() {
	*x = StatAggregateConcat{}
	mi := &file_spaceone_api_statistics_v1_resource_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatAggregateConcat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatAggregateConcat) ProtoMessage() {}

func (x *StatAggregateConcat) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_statistics_v1_resource_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatAggregateConcat.ProtoReflect.Descriptor instead.
func (*StatAggregateConcat) Descriptor() ([]byte, []int) {
	return file_spaceone_api_statistics_v1_resource_proto_rawDescGZIP(), []int{2}
}

func (x *StatAggregateConcat) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *StatAggregateConcat) GetQuery() *v2.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *StatAggregateConcat) GetExtendData() *_struct.Struct {
	if x != nil {
		return x.ExtendData
	}
	return nil
}

type Sort struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Desc          bool                   `protobuf:"varint,2,opt,name=desc,proto3" json:"desc,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Sort) Reset() {
	*x = Sort{}
	mi := &file_spaceone_api_statistics_v1_resource_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Sort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sort) ProtoMessage() {}

func (x *Sort) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_statistics_v1_resource_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sort.ProtoReflect.Descriptor instead.
func (*Sort) Descriptor() ([]byte, []int) {
	return file_spaceone_api_statistics_v1_resource_proto_rawDescGZIP(), []int{3}
}

func (x *Sort) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Sort) GetDesc() bool {
	if x != nil {
		return x.Desc
	}
	return false
}

type StatAggregateFormula struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to FormulaAlias:
	//
	//	*StatAggregateFormula_Eval
	//	*StatAggregateFormula_Query
	FormulaAlias  isStatAggregateFormula_FormulaAlias `protobuf_oneof:"formula_alias"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StatAggregateFormula) Reset() {
	*x = StatAggregateFormula{}
	mi := &file_spaceone_api_statistics_v1_resource_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatAggregateFormula) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatAggregateFormula) ProtoMessage() {}

func (x *StatAggregateFormula) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_statistics_v1_resource_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatAggregateFormula.ProtoReflect.Descriptor instead.
func (*StatAggregateFormula) Descriptor() ([]byte, []int) {
	return file_spaceone_api_statistics_v1_resource_proto_rawDescGZIP(), []int{4}
}

func (x *StatAggregateFormula) GetFormulaAlias() isStatAggregateFormula_FormulaAlias {
	if x != nil {
		return x.FormulaAlias
	}
	return nil
}

func (x *StatAggregateFormula) GetEval() string {
	if x != nil {
		if x, ok := x.FormulaAlias.(*StatAggregateFormula_Eval); ok {
			return x.Eval
		}
	}
	return ""
}

func (x *StatAggregateFormula) GetQuery() string {
	if x != nil {
		if x, ok := x.FormulaAlias.(*StatAggregateFormula_Query); ok {
			return x.Query
		}
	}
	return ""
}

type isStatAggregateFormula_FormulaAlias interface {
	isStatAggregateFormula_FormulaAlias()
}

type StatAggregateFormula_Eval struct {
	Eval string `protobuf:"bytes,1,opt,name=eval,proto3,oneof"`
}

type StatAggregateFormula_Query struct {
	Query string `protobuf:"bytes,2,opt,name=query,proto3,oneof"`
}

func (*StatAggregateFormula_Eval) isStatAggregateFormula_FormulaAlias() {}

func (*StatAggregateFormula_Query) isStatAggregateFormula_FormulaAlias() {}

type StatAggregateFillNA struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          *_struct.Struct        `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StatAggregateFillNA) Reset() {
	*x = StatAggregateFillNA{}
	mi := &file_spaceone_api_statistics_v1_resource_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatAggregateFillNA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatAggregateFillNA) ProtoMessage() {}

func (x *StatAggregateFillNA) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_statistics_v1_resource_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatAggregateFillNA.ProtoReflect.Descriptor instead.
func (*StatAggregateFillNA) Descriptor() ([]byte, []int) {
	return file_spaceone_api_statistics_v1_resource_proto_rawDescGZIP(), []int{5}
}

func (x *StatAggregateFillNA) GetData() *_struct.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

type StatAggregate struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to AggregateAlias:
	//
	//	*StatAggregate_Query
	//	*StatAggregate_Join
	//	*StatAggregate_Concat
	//	*StatAggregate_Sort
	//	*StatAggregate_Formula
	//	*StatAggregate_FillNa
	AggregateAlias isStatAggregate_AggregateAlias `protobuf_oneof:"aggregate_alias"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *StatAggregate) Reset() {
	*x = StatAggregate{}
	mi := &file_spaceone_api_statistics_v1_resource_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatAggregate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatAggregate) ProtoMessage() {}

func (x *StatAggregate) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_statistics_v1_resource_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatAggregate.ProtoReflect.Descriptor instead.
func (*StatAggregate) Descriptor() ([]byte, []int) {
	return file_spaceone_api_statistics_v1_resource_proto_rawDescGZIP(), []int{6}
}

func (x *StatAggregate) GetAggregateAlias() isStatAggregate_AggregateAlias {
	if x != nil {
		return x.AggregateAlias
	}
	return nil
}

func (x *StatAggregate) GetQuery() *StatAggregateQuery {
	if x != nil {
		if x, ok := x.AggregateAlias.(*StatAggregate_Query); ok {
			return x.Query
		}
	}
	return nil
}

func (x *StatAggregate) GetJoin() *StatAggregateJoin {
	if x != nil {
		if x, ok := x.AggregateAlias.(*StatAggregate_Join); ok {
			return x.Join
		}
	}
	return nil
}

func (x *StatAggregate) GetConcat() *StatAggregateConcat {
	if x != nil {
		if x, ok := x.AggregateAlias.(*StatAggregate_Concat); ok {
			return x.Concat
		}
	}
	return nil
}

func (x *StatAggregate) GetSort() *_struct.ListValue {
	if x != nil {
		if x, ok := x.AggregateAlias.(*StatAggregate_Sort); ok {
			return x.Sort
		}
	}
	return nil
}

func (x *StatAggregate) GetFormula() *StatAggregateFormula {
	if x != nil {
		if x, ok := x.AggregateAlias.(*StatAggregate_Formula); ok {
			return x.Formula
		}
	}
	return nil
}

func (x *StatAggregate) GetFillNa() *StatAggregateFillNA {
	if x != nil {
		if x, ok := x.AggregateAlias.(*StatAggregate_FillNa); ok {
			return x.FillNa
		}
	}
	return nil
}

type isStatAggregate_AggregateAlias interface {
	isStatAggregate_AggregateAlias()
}

type StatAggregate_Query struct {
	Query *StatAggregateQuery `protobuf:"bytes,1,opt,name=query,proto3,oneof"`
}

type StatAggregate_Join struct {
	Join *StatAggregateJoin `protobuf:"bytes,2,opt,name=join,proto3,oneof"`
}

type StatAggregate_Concat struct {
	Concat *StatAggregateConcat `protobuf:"bytes,3,opt,name=concat,proto3,oneof"`
}

type StatAggregate_Sort struct {
	Sort *_struct.ListValue `protobuf:"bytes,4,opt,name=sort,proto3,oneof"`
}

type StatAggregate_Formula struct {
	Formula *StatAggregateFormula `protobuf:"bytes,5,opt,name=formula,proto3,oneof"`
}

type StatAggregate_FillNa struct {
	FillNa *StatAggregateFillNA `protobuf:"bytes,6,opt,name=fill_na,json=fillNa,proto3,oneof"`
}

func (*StatAggregate_Query) isStatAggregate_AggregateAlias() {}

func (*StatAggregate_Join) isStatAggregate_AggregateAlias() {}

func (*StatAggregate_Concat) isStatAggregate_AggregateAlias() {}

func (*StatAggregate_Sort) isStatAggregate_AggregateAlias() {}

func (*StatAggregate_Formula) isStatAggregate_AggregateAlias() {}

func (*StatAggregate_FillNa) isStatAggregate_AggregateAlias() {}

type StatPage struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// +optional
	Start         uint32 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	Limit         uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StatPage) Reset() {
	*x = StatPage{}
	mi := &file_spaceone_api_statistics_v1_resource_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatPage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatPage) ProtoMessage() {}

func (x *StatPage) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_statistics_v1_resource_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatPage.ProtoReflect.Descriptor instead.
func (*StatPage) Descriptor() ([]byte, []int) {
	return file_spaceone_api_statistics_v1_resource_proto_rawDescGZIP(), []int{7}
}

func (x *StatPage) GetStart() uint32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *StatPage) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ResourceStatRequest struct {
	state     protoimpl.MessageState `protogen:"open.v1"`
	Aggregate []*StatAggregate       `protobuf:"bytes,1,rep,name=aggregate,proto3" json:"aggregate,omitempty"`
	// +optional
	Page          *StatPage `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResourceStatRequest) Reset() {
	*x = ResourceStatRequest{}
	mi := &file_spaceone_api_statistics_v1_resource_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceStatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceStatRequest) ProtoMessage() {}

func (x *ResourceStatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_statistics_v1_resource_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceStatRequest.ProtoReflect.Descriptor instead.
func (*ResourceStatRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_statistics_v1_resource_proto_rawDescGZIP(), []int{8}
}

func (x *ResourceStatRequest) GetAggregate() []*StatAggregate {
	if x != nil {
		return x.Aggregate
	}
	return nil
}

func (x *ResourceStatRequest) GetPage() *StatPage {
	if x != nil {
		return x.Page
	}
	return nil
}

var File_spaceone_api_statistics_v1_resource_proto protoreflect.FileDescriptor

var file_spaceone_api_statistics_v1_resource_proto_rawDesc = []byte{
	0x0a, 0x29, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a, 0x12, 0x53, 0x74, 0x61, 0x74, 0x41, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x3b, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x38,
	0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x22, 0xc6, 0x02, 0x0a, 0x11, 0x53, 0x74, 0x61,
	0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x23,
	0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x3b, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x38, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x4a, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x4a, 0x6f, 0x69, 0x6e, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x22, 0x35, 0x0a, 0x08, 0x4a, 0x6f,
	0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x45, 0x46, 0x54, 0x10, 0x00,
	0x12, 0x09, 0x0a, 0x05, 0x52, 0x49, 0x47, 0x48, 0x54, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4f,
	0x55, 0x54, 0x45, 0x52, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4e, 0x4e, 0x45, 0x52, 0x10,
	0x03, 0x22, 0xb1, 0x01, 0x0a, 0x13, 0x53, 0x74, 0x61, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x63, 0x61, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3b,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x38, 0x0a, 0x0b, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x64, 0x44, 0x61, 0x74, 0x61, 0x22, 0x2c, 0x0a, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64,
	0x65, 0x73, 0x63, 0x22, 0x55, 0x0a, 0x14, 0x53, 0x74, 0x61, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x12, 0x14, 0x0a, 0x04, 0x65,
	0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x65, 0x76, 0x61,
	0x6c, 0x12, 0x16, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x42, 0x0f, 0x0a, 0x0d, 0x66, 0x6f, 0x72,
	0x6d, 0x75, 0x6c, 0x61, 0x5f, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x22, 0x42, 0x0a, 0x13, 0x53, 0x74,
	0x61, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x6c, 0x4e,
	0x41, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xc6,
	0x03, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x12, 0x46, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48,
	0x00, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x43, 0x0a, 0x04, 0x6a, 0x6f, 0x69, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x4a, 0x6f, 0x69, 0x6e, 0x48, 0x00, 0x52, 0x04, 0x6a, 0x6f, 0x69, 0x6e, 0x12, 0x49, 0x0a,
	0x06, 0x63, 0x6f, 0x6e, 0x63, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x41,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x63, 0x61, 0x74, 0x48, 0x00,
	0x52, 0x06, 0x63, 0x6f, 0x6e, 0x63, 0x61, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x48, 0x00, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x4c, 0x0a, 0x07, 0x66, 0x6f,
	0x72, 0x6d, 0x75, 0x6c, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x41, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x48, 0x00, 0x52,
	0x07, 0x66, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x12, 0x4a, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x6c,
	0x5f, 0x6e, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x6c, 0x4e, 0x41, 0x48, 0x00, 0x52, 0x06, 0x66, 0x69,
	0x6c, 0x6c, 0x4e, 0x61, 0x42, 0x11, 0x0a, 0x0f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x5f, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x22, 0x36, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22,
	0x98, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x09, 0x61, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x41, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x09, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x12, 0x38, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x32, 0x85, 0x01, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x79, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x74, 0x12,
	0x2f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x21, 0x3a, 0x01, 0x2a, 0x22, 0x1c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_statistics_v1_resource_proto_rawDescOnce sync.Once
	file_spaceone_api_statistics_v1_resource_proto_rawDescData = file_spaceone_api_statistics_v1_resource_proto_rawDesc
)

func file_spaceone_api_statistics_v1_resource_proto_rawDescGZIP() []byte {
	file_spaceone_api_statistics_v1_resource_proto_rawDescOnce.Do(func() {
		file_spaceone_api_statistics_v1_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_statistics_v1_resource_proto_rawDescData)
	})
	return file_spaceone_api_statistics_v1_resource_proto_rawDescData
}

var file_spaceone_api_statistics_v1_resource_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spaceone_api_statistics_v1_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_spaceone_api_statistics_v1_resource_proto_goTypes = []any{
	(StatAggregateJoin_JoinType)(0), // 0: spaceone.api.statistics.v1.StatAggregateJoin.JoinType
	(*StatAggregateQuery)(nil),      // 1: spaceone.api.statistics.v1.StatAggregateQuery
	(*StatAggregateJoin)(nil),       // 2: spaceone.api.statistics.v1.StatAggregateJoin
	(*StatAggregateConcat)(nil),     // 3: spaceone.api.statistics.v1.StatAggregateConcat
	(*Sort)(nil),                    // 4: spaceone.api.statistics.v1.Sort
	(*StatAggregateFormula)(nil),    // 5: spaceone.api.statistics.v1.StatAggregateFormula
	(*StatAggregateFillNA)(nil),     // 6: spaceone.api.statistics.v1.StatAggregateFillNA
	(*StatAggregate)(nil),           // 7: spaceone.api.statistics.v1.StatAggregate
	(*StatPage)(nil),                // 8: spaceone.api.statistics.v1.StatPage
	(*ResourceStatRequest)(nil),     // 9: spaceone.api.statistics.v1.ResourceStatRequest
	(*v2.StatisticsQuery)(nil),      // 10: spaceone.api.core.v2.StatisticsQuery
	(*_struct.Struct)(nil),          // 11: google.protobuf.Struct
	(*_struct.ListValue)(nil),       // 12: google.protobuf.ListValue
}
var file_spaceone_api_statistics_v1_resource_proto_depIdxs = []int32{
	10, // 0: spaceone.api.statistics.v1.StatAggregateQuery.query:type_name -> spaceone.api.core.v2.StatisticsQuery
	11, // 1: spaceone.api.statistics.v1.StatAggregateQuery.extend_data:type_name -> google.protobuf.Struct
	10, // 2: spaceone.api.statistics.v1.StatAggregateJoin.query:type_name -> spaceone.api.core.v2.StatisticsQuery
	11, // 3: spaceone.api.statistics.v1.StatAggregateJoin.extend_data:type_name -> google.protobuf.Struct
	0,  // 4: spaceone.api.statistics.v1.StatAggregateJoin.type:type_name -> spaceone.api.statistics.v1.StatAggregateJoin.JoinType
	10, // 5: spaceone.api.statistics.v1.StatAggregateConcat.query:type_name -> spaceone.api.core.v2.StatisticsQuery
	11, // 6: spaceone.api.statistics.v1.StatAggregateConcat.extend_data:type_name -> google.protobuf.Struct
	11, // 7: spaceone.api.statistics.v1.StatAggregateFillNA.data:type_name -> google.protobuf.Struct
	1,  // 8: spaceone.api.statistics.v1.StatAggregate.query:type_name -> spaceone.api.statistics.v1.StatAggregateQuery
	2,  // 9: spaceone.api.statistics.v1.StatAggregate.join:type_name -> spaceone.api.statistics.v1.StatAggregateJoin
	3,  // 10: spaceone.api.statistics.v1.StatAggregate.concat:type_name -> spaceone.api.statistics.v1.StatAggregateConcat
	12, // 11: spaceone.api.statistics.v1.StatAggregate.sort:type_name -> google.protobuf.ListValue
	5,  // 12: spaceone.api.statistics.v1.StatAggregate.formula:type_name -> spaceone.api.statistics.v1.StatAggregateFormula
	6,  // 13: spaceone.api.statistics.v1.StatAggregate.fill_na:type_name -> spaceone.api.statistics.v1.StatAggregateFillNA
	7,  // 14: spaceone.api.statistics.v1.ResourceStatRequest.aggregate:type_name -> spaceone.api.statistics.v1.StatAggregate
	8,  // 15: spaceone.api.statistics.v1.ResourceStatRequest.page:type_name -> spaceone.api.statistics.v1.StatPage
	9,  // 16: spaceone.api.statistics.v1.Resource.stat:input_type -> spaceone.api.statistics.v1.ResourceStatRequest
	11, // 17: spaceone.api.statistics.v1.Resource.stat:output_type -> google.protobuf.Struct
	17, // [17:18] is the sub-list for method output_type
	16, // [16:17] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_spaceone_api_statistics_v1_resource_proto_init() }
func file_spaceone_api_statistics_v1_resource_proto_init() {
	if File_spaceone_api_statistics_v1_resource_proto != nil {
		return
	}
	file_spaceone_api_statistics_v1_resource_proto_msgTypes[4].OneofWrappers = []any{
		(*StatAggregateFormula_Eval)(nil),
		(*StatAggregateFormula_Query)(nil),
	}
	file_spaceone_api_statistics_v1_resource_proto_msgTypes[6].OneofWrappers = []any{
		(*StatAggregate_Query)(nil),
		(*StatAggregate_Join)(nil),
		(*StatAggregate_Concat)(nil),
		(*StatAggregate_Sort)(nil),
		(*StatAggregate_Formula)(nil),
		(*StatAggregate_FillNa)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spaceone_api_statistics_v1_resource_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_statistics_v1_resource_proto_goTypes,
		DependencyIndexes: file_spaceone_api_statistics_v1_resource_proto_depIdxs,
		EnumInfos:         file_spaceone_api_statistics_v1_resource_proto_enumTypes,
		MessageInfos:      file_spaceone_api_statistics_v1_resource_proto_msgTypes,
	}.Build()
	File_spaceone_api_statistics_v1_resource_proto = out.File
	file_spaceone_api_statistics_v1_resource_proto_rawDesc = nil
	file_spaceone_api_statistics_v1_resource_proto_goTypes = nil
	file_spaceone_api_statistics_v1_resource_proto_depIdxs = nil
}
