// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v3.6.1
// source: spaceone/api/alert_manager/v1/event.proto

package v1

import (
	_ "github.com/cloudforet-io/api/dist/go/spaceone/api/core/v2"
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

type EventSeverity int32

const (
	EventSeverity_EVENT_SEVERITY_NONE EventSeverity = 0
	EventSeverity_CRITICAL            EventSeverity = 1
	EventSeverity_ERROR               EventSeverity = 2
	EventSeverity_WARNING             EventSeverity = 3
	EventSeverity_INFO                EventSeverity = 4
)

// Enum value maps for EventSeverity.
var (
	EventSeverity_name = map[int32]string{
		0: "EVENT_SEVERITY_NONE",
		1: "CRITICAL",
		2: "ERROR",
		3: "WARNING",
		4: "INFO",
	}
	EventSeverity_value = map[string]int32{
		"EVENT_SEVERITY_NONE": 0,
		"CRITICAL":            1,
		"ERROR":               2,
		"WARNING":             3,
		"INFO":                4,
	}
)

func (x EventSeverity) Enum() *EventSeverity {
	p := new(EventSeverity)
	*p = x
	return p
}

func (x EventSeverity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventSeverity) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_alert_manager_v1_event_proto_enumTypes[0].Descriptor()
}

func (EventSeverity) Type() protoreflect.EnumType {
	return &file_spaceone_api_alert_manager_v1_event_proto_enumTypes[0]
}

func (x EventSeverity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventSeverity.Descriptor instead.
func (EventSeverity) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_alert_manager_v1_event_proto_rawDescGZIP(), []int{0}
}

type EventInfo_EventType int32

const (
	EventInfo_EVENT_TYPE_NONE EventInfo_EventType = 0
	EventInfo_ALERT           EventInfo_EventType = 1
	EventInfo_RECOVERY        EventInfo_EventType = 2
	EventInfo_ERROR           EventInfo_EventType = 3
)

// Enum value maps for EventInfo_EventType.
var (
	EventInfo_EventType_name = map[int32]string{
		0: "EVENT_TYPE_NONE",
		1: "ALERT",
		2: "RECOVERY",
		3: "ERROR",
	}
	EventInfo_EventType_value = map[string]int32{
		"EVENT_TYPE_NONE": 0,
		"ALERT":           1,
		"RECOVERY":        2,
		"ERROR":           3,
	}
)

func (x EventInfo_EventType) Enum() *EventInfo_EventType {
	p := new(EventInfo_EventType)
	*p = x
	return p
}

func (x EventInfo_EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventInfo_EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_alert_manager_v1_event_proto_enumTypes[1].Descriptor()
}

func (EventInfo_EventType) Type() protoreflect.EnumType {
	return &file_spaceone_api_alert_manager_v1_event_proto_enumTypes[1]
}

func (x EventInfo_EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventInfo_EventType.Descriptor instead.
func (EventInfo_EventType) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_alert_manager_v1_event_proto_rawDescGZIP(), []int{1, 0}
}

type EventCreateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WebhookId     string                 `protobuf:"bytes,1,opt,name=webhook_id,json=webhookId,proto3" json:"webhook_id,omitempty"`
	AccessKey     string                 `protobuf:"bytes,2,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	Data          *_struct.Struct        `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EventCreateRequest) Reset() {
	*x = EventCreateRequest{}
	mi := &file_spaceone_api_alert_manager_v1_event_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventCreateRequest) ProtoMessage() {}

func (x *EventCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_alert_manager_v1_event_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventCreateRequest.ProtoReflect.Descriptor instead.
func (*EventCreateRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_alert_manager_v1_event_proto_rawDescGZIP(), []int{0}
}

func (x *EventCreateRequest) GetWebhookId() string {
	if x != nil {
		return x.WebhookId
	}
	return ""
}

func (x *EventCreateRequest) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *EventCreateRequest) GetData() *_struct.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

type EventInfo struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	EventId        string                 `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	EventKey       string                 `protobuf:"bytes,2,opt,name=event_key,json=eventKey,proto3" json:"event_key,omitempty"`
	EventType      EventInfo_EventType    `protobuf:"varint,3,opt,name=event_type,json=eventType,proto3,enum=spaceone.api.alert_manager.v1.EventInfo_EventType" json:"event_type,omitempty"`
	Title          string                 `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Description    string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Severity       EventSeverity          `protobuf:"varint,6,opt,name=severity,proto3,enum=spaceone.api.alert_manager.v1.EventSeverity" json:"severity,omitempty"`
	Rule           string                 `protobuf:"bytes,7,opt,name=rule,proto3" json:"rule,omitempty"`
	ImageUrl       string                 `protobuf:"bytes,8,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	Resources      []string               `protobuf:"bytes,9,rep,name=resources,proto3" json:"resources,omitempty"`
	Provider       string                 `protobuf:"bytes,10,opt,name=provider,proto3" json:"provider,omitempty"`
	Account        string                 `protobuf:"bytes,11,opt,name=account,proto3" json:"account,omitempty"`
	AdditionalInfo *_struct.Struct        `protobuf:"bytes,15,opt,name=additional_info,json=additionalInfo,proto3" json:"additional_info,omitempty"`
	RawData        *_struct.Struct        `protobuf:"bytes,16,opt,name=raw_data,json=rawData,proto3" json:"raw_data,omitempty"`
	DomainId       string                 `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	WorkspaceId    string                 `protobuf:"bytes,22,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	ServiceId      string                 `protobuf:"bytes,23,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	WebhookId      string                 `protobuf:"bytes,24,opt,name=webhook_id,json=webhookId,proto3" json:"webhook_id,omitempty"`
	AlertId        string                 `protobuf:"bytes,25,opt,name=alert_id,json=alertId,proto3" json:"alert_id,omitempty"`
	CreatedAt      string                 `protobuf:"bytes,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	OccurredAt     string                 `protobuf:"bytes,32,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *EventInfo) Reset() {
	*x = EventInfo{}
	mi := &file_spaceone_api_alert_manager_v1_event_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventInfo) ProtoMessage() {}

func (x *EventInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_alert_manager_v1_event_proto_msgTypes[1]
	if x != nil {
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
	return file_spaceone_api_alert_manager_v1_event_proto_rawDescGZIP(), []int{1}
}

func (x *EventInfo) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *EventInfo) GetEventKey() string {
	if x != nil {
		return x.EventKey
	}
	return ""
}

func (x *EventInfo) GetEventType() EventInfo_EventType {
	if x != nil {
		return x.EventType
	}
	return EventInfo_EVENT_TYPE_NONE
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

func (x *EventInfo) GetSeverity() EventSeverity {
	if x != nil {
		return x.Severity
	}
	return EventSeverity_EVENT_SEVERITY_NONE
}

func (x *EventInfo) GetRule() string {
	if x != nil {
		return x.Rule
	}
	return ""
}

func (x *EventInfo) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *EventInfo) GetResources() []string {
	if x != nil {
		return x.Resources
	}
	return nil
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

func (x *EventInfo) GetAdditionalInfo() *_struct.Struct {
	if x != nil {
		return x.AdditionalInfo
	}
	return nil
}

func (x *EventInfo) GetRawData() *_struct.Struct {
	if x != nil {
		return x.RawData
	}
	return nil
}

func (x *EventInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *EventInfo) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *EventInfo) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *EventInfo) GetWebhookId() string {
	if x != nil {
		return x.WebhookId
	}
	return ""
}

func (x *EventInfo) GetAlertId() string {
	if x != nil {
		return x.AlertId
	}
	return ""
}

func (x *EventInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *EventInfo) GetOccurredAt() string {
	if x != nil {
		return x.OccurredAt
	}
	return ""
}

var File_spaceone_api_alert_manager_v1_event_proto protoreflect.FileDescriptor

var file_spaceone_api_alert_manager_v1_event_proto_rawDesc = []byte{
	0x0a, 0x29, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x12, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x77,
	0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xb2, 0x06, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x51, 0x0a, 0x0a,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x32, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x48, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72,
	0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x72, 0x75, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55,
	0x72, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x40, 0x0a, 0x0f, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0e, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x32, 0x0a, 0x08, 0x72, 0x61, 0x77, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x07, 0x72, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x77,
	0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x63, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x64, 0x41, 0x74, 0x22, 0x44, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4c, 0x45, 0x52, 0x54,
	0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x59, 0x10, 0x02,
	0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x2a, 0x58, 0x0a, 0x0d, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x17, 0x0a, 0x13,
	0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x49, 0x54, 0x49, 0x43, 0x41,
	0x4c, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x0b,
	0x0a, 0x07, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x49,
	0x4e, 0x46, 0x4f, 0x10, 0x04, 0x32, 0x8a, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x80, 0x01, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x31, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x3a, 0x01, 0x2a,
	0x22, 0x20, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2d, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_alert_manager_v1_event_proto_rawDescOnce sync.Once
	file_spaceone_api_alert_manager_v1_event_proto_rawDescData = file_spaceone_api_alert_manager_v1_event_proto_rawDesc
)

func file_spaceone_api_alert_manager_v1_event_proto_rawDescGZIP() []byte {
	file_spaceone_api_alert_manager_v1_event_proto_rawDescOnce.Do(func() {
		file_spaceone_api_alert_manager_v1_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_alert_manager_v1_event_proto_rawDescData)
	})
	return file_spaceone_api_alert_manager_v1_event_proto_rawDescData
}

var file_spaceone_api_alert_manager_v1_event_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_spaceone_api_alert_manager_v1_event_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_spaceone_api_alert_manager_v1_event_proto_goTypes = []any{
	(EventSeverity)(0),         // 0: spaceone.api.alert_manager.v1.EventSeverity
	(EventInfo_EventType)(0),   // 1: spaceone.api.alert_manager.v1.EventInfo.EventType
	(*EventCreateRequest)(nil), // 2: spaceone.api.alert_manager.v1.EventCreateRequest
	(*EventInfo)(nil),          // 3: spaceone.api.alert_manager.v1.EventInfo
	(*_struct.Struct)(nil),     // 4: google.protobuf.Struct
	(*empty.Empty)(nil),        // 5: google.protobuf.Empty
}
var file_spaceone_api_alert_manager_v1_event_proto_depIdxs = []int32{
	4, // 0: spaceone.api.alert_manager.v1.EventCreateRequest.data:type_name -> google.protobuf.Struct
	1, // 1: spaceone.api.alert_manager.v1.EventInfo.event_type:type_name -> spaceone.api.alert_manager.v1.EventInfo.EventType
	0, // 2: spaceone.api.alert_manager.v1.EventInfo.severity:type_name -> spaceone.api.alert_manager.v1.EventSeverity
	4, // 3: spaceone.api.alert_manager.v1.EventInfo.additional_info:type_name -> google.protobuf.Struct
	4, // 4: spaceone.api.alert_manager.v1.EventInfo.raw_data:type_name -> google.protobuf.Struct
	2, // 5: spaceone.api.alert_manager.v1.Event.create:input_type -> spaceone.api.alert_manager.v1.EventCreateRequest
	5, // 6: spaceone.api.alert_manager.v1.Event.create:output_type -> google.protobuf.Empty
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_spaceone_api_alert_manager_v1_event_proto_init() }
func file_spaceone_api_alert_manager_v1_event_proto_init() {
	if File_spaceone_api_alert_manager_v1_event_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spaceone_api_alert_manager_v1_event_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_alert_manager_v1_event_proto_goTypes,
		DependencyIndexes: file_spaceone_api_alert_manager_v1_event_proto_depIdxs,
		EnumInfos:         file_spaceone_api_alert_manager_v1_event_proto_enumTypes,
		MessageInfos:      file_spaceone_api_alert_manager_v1_event_proto_msgTypes,
	}.Build()
	File_spaceone_api_alert_manager_v1_event_proto = out.File
	file_spaceone_api_alert_manager_v1_event_proto_rawDesc = nil
	file_spaceone_api_alert_manager_v1_event_proto_goTypes = nil
	file_spaceone_api_alert_manager_v1_event_proto_depIdxs = nil
}
