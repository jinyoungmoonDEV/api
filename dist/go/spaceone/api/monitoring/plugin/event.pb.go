//
//desc: An Event is data created by an external monitoring system and collected by a Webhook plugin.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.6.1
// source: spaceone/api/monitoring/plugin/event.proto

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

type EventInfo_Severity int32

const (
	EventInfo_NONE          EventInfo_Severity = 0
	EventInfo_CRITICAL      EventInfo_Severity = 1
	EventInfo_ERROR         EventInfo_Severity = 2
	EventInfo_WARNING       EventInfo_Severity = 3
	EventInfo_INFO          EventInfo_Severity = 4
	EventInfo_NOT_AVAILABLE EventInfo_Severity = 5
)

// Enum value maps for EventInfo_Severity.
var (
	EventInfo_Severity_name = map[int32]string{
		0: "NONE",
		1: "CRITICAL",
		2: "ERROR",
		3: "WARNING",
		4: "INFO",
		5: "NOT_AVAILABLE",
	}
	EventInfo_Severity_value = map[string]int32{
		"NONE":          0,
		"CRITICAL":      1,
		"ERROR":         2,
		"WARNING":       3,
		"INFO":          4,
		"NOT_AVAILABLE": 5,
	}
)

func (x EventInfo_Severity) Enum() *EventInfo_Severity {
	p := new(EventInfo_Severity)
	*p = x
	return p
}

func (x EventInfo_Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventInfo_Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_monitoring_plugin_event_proto_enumTypes[0].Descriptor()
}

func (EventInfo_Severity) Type() protoreflect.EnumType {
	return &file_spaceone_api_monitoring_plugin_event_proto_enumTypes[0]
}

func (x EventInfo_Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventInfo_Severity.Descriptor instead.
func (EventInfo_Severity) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_monitoring_plugin_event_proto_rawDescGZIP(), []int{1, 0}
}

type ParseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// is_required: true
	Options *_struct.Struct `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	// is_required: true
	// desc: Unpredictable data format that comes from Webhook
	Data *_struct.Struct `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ParseRequest) Reset() {
	*x = ParseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_monitoring_plugin_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseRequest) ProtoMessage() {}

func (x *ParseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_monitoring_plugin_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseRequest.ProtoReflect.Descriptor instead.
func (*ParseRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_monitoring_plugin_event_proto_rawDescGZIP(), []int{0}
}

func (x *ParseRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *ParseRequest) GetData() *_struct.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

type EventInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventKey       string             `protobuf:"bytes,1,opt,name=event_key,json=eventKey,proto3" json:"event_key,omitempty"`
	EventType      string             `protobuf:"bytes,2,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	Title          string             `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description    string             `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Severity       EventInfo_Severity `protobuf:"varint,5,opt,name=severity,proto3,enum=spaceone.api.monitoring.plugin.EventInfo_Severity" json:"severity,omitempty"`
	Resource       *_struct.Struct    `protobuf:"bytes,6,opt,name=resource,proto3" json:"resource,omitempty"`
	Rule           string             `protobuf:"bytes,7,opt,name=rule,proto3" json:"rule,omitempty"`
	OccurredAt     string             `protobuf:"bytes,8,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
	AdditionalInfo *_struct.Struct    `protobuf:"bytes,9,opt,name=additional_info,json=additionalInfo,proto3" json:"additional_info,omitempty"`
	ImageUrl       string             `protobuf:"bytes,10,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	Provider       string             `protobuf:"bytes,11,opt,name=provider,proto3" json:"provider,omitempty"`
	Account        string             `protobuf:"bytes,12,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *EventInfo) Reset() {
	*x = EventInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_monitoring_plugin_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventInfo) ProtoMessage() {}

func (x *EventInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_monitoring_plugin_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventInfo.ProtoReflect.Descriptor instead.
func (*EventInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_monitoring_plugin_event_proto_rawDescGZIP(), []int{1}
}

func (x *EventInfo) GetEventKey() string {
	if x != nil {
		return x.EventKey
	}
	return ""
}

func (x *EventInfo) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *EventInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *EventInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *EventInfo) GetSeverity() EventInfo_Severity {
	if x != nil {
		return x.Severity
	}
	return EventInfo_NONE
}

func (x *EventInfo) GetResource() *_struct.Struct {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *EventInfo) GetRule() string {
	if x != nil {
		return x.Rule
	}
	return ""
}

func (x *EventInfo) GetOccurredAt() string {
	if x != nil {
		return x.OccurredAt
	}
	return ""
}

func (x *EventInfo) GetAdditionalInfo() *_struct.Struct {
	if x != nil {
		return x.AdditionalInfo
	}
	return nil
}

func (x *EventInfo) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *EventInfo) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *EventInfo) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

type EventsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*EventInfo `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *EventsInfo) Reset() {
	*x = EventsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_monitoring_plugin_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventsInfo) ProtoMessage() {}

func (x *EventsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_monitoring_plugin_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventsInfo.ProtoReflect.Descriptor instead.
func (*EventsInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_monitoring_plugin_event_proto_rawDescGZIP(), []int{2}
}

func (x *EventsInfo) GetResults() []*EventInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_spaceone_api_monitoring_plugin_event_proto protoreflect.FileDescriptor

var file_spaceone_api_monitoring_plugin_event_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x0c, 0x50, 0x61,
	0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2b, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xa7, 0x04, 0x0a, 0x09, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4e, 0x0a, 0x08,
	0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69,
	0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x33, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x72, 0x75, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x63, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x64, 0x41, 0x74, 0x12, 0x40, 0x0a, 0x0f, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0e, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x57, 0x0a, 0x08, 0x53,
	0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x49, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x01, 0x12,
	0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41,
	0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x46, 0x4f, 0x10,
	0x04, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42,
	0x4c, 0x45, 0x10, 0x05, 0x22, 0x51, 0x0a, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x43, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x32, 0x6c, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x63, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x73, 0x65, 0x12, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f,
	0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x00, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_monitoring_plugin_event_proto_rawDescOnce sync.Once
	file_spaceone_api_monitoring_plugin_event_proto_rawDescData = file_spaceone_api_monitoring_plugin_event_proto_rawDesc
)

func file_spaceone_api_monitoring_plugin_event_proto_rawDescGZIP() []byte {
	file_spaceone_api_monitoring_plugin_event_proto_rawDescOnce.Do(func() {
		file_spaceone_api_monitoring_plugin_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_monitoring_plugin_event_proto_rawDescData)
	})
	return file_spaceone_api_monitoring_plugin_event_proto_rawDescData
}

var file_spaceone_api_monitoring_plugin_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spaceone_api_monitoring_plugin_event_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_spaceone_api_monitoring_plugin_event_proto_goTypes = []interface{}{
	(EventInfo_Severity)(0), // 0: spaceone.api.monitoring.plugin.EventInfo.Severity
	(*ParseRequest)(nil),    // 1: spaceone.api.monitoring.plugin.ParseRequest
	(*EventInfo)(nil),       // 2: spaceone.api.monitoring.plugin.EventInfo
	(*EventsInfo)(nil),      // 3: spaceone.api.monitoring.plugin.EventsInfo
	(*_struct.Struct)(nil),  // 4: google.protobuf.Struct
}
var file_spaceone_api_monitoring_plugin_event_proto_depIdxs = []int32{
	4, // 0: spaceone.api.monitoring.plugin.ParseRequest.options:type_name -> google.protobuf.Struct
	4, // 1: spaceone.api.monitoring.plugin.ParseRequest.data:type_name -> google.protobuf.Struct
	0, // 2: spaceone.api.monitoring.plugin.EventInfo.severity:type_name -> spaceone.api.monitoring.plugin.EventInfo.Severity
	4, // 3: spaceone.api.monitoring.plugin.EventInfo.resource:type_name -> google.protobuf.Struct
	4, // 4: spaceone.api.monitoring.plugin.EventInfo.additional_info:type_name -> google.protobuf.Struct
	2, // 5: spaceone.api.monitoring.plugin.EventsInfo.results:type_name -> spaceone.api.monitoring.plugin.EventInfo
	1, // 6: spaceone.api.monitoring.plugin.Event.parse:input_type -> spaceone.api.monitoring.plugin.ParseRequest
	3, // 7: spaceone.api.monitoring.plugin.Event.parse:output_type -> spaceone.api.monitoring.plugin.EventsInfo
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_spaceone_api_monitoring_plugin_event_proto_init() }
func file_spaceone_api_monitoring_plugin_event_proto_init() {
	if File_spaceone_api_monitoring_plugin_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spaceone_api_monitoring_plugin_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParseRequest); i {
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
		file_spaceone_api_monitoring_plugin_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventInfo); i {
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
		file_spaceone_api_monitoring_plugin_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventsInfo); i {
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
			RawDescriptor: file_spaceone_api_monitoring_plugin_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_monitoring_plugin_event_proto_goTypes,
		DependencyIndexes: file_spaceone_api_monitoring_plugin_event_proto_depIdxs,
		EnumInfos:         file_spaceone_api_monitoring_plugin_event_proto_enumTypes,
		MessageInfos:      file_spaceone_api_monitoring_plugin_event_proto_msgTypes,
	}.Build()
	File_spaceone_api_monitoring_plugin_event_proto = out.File
	file_spaceone_api_monitoring_plugin_event_proto_rawDesc = nil
	file_spaceone_api_monitoring_plugin_event_proto_goTypes = nil
	file_spaceone_api_monitoring_plugin_event_proto_depIdxs = nil
}
