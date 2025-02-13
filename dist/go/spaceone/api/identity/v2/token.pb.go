// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v3.6.1
// source: spaceone/api/identity/v2/token.proto

package v2

import (
	_ "github.com/golang/protobuf/ptypes/empty"
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

type IssueTokenRequest_AuthType int32

const (
	IssueTokenRequest_NONE     IssueTokenRequest_AuthType = 0
	IssueTokenRequest_LOCAL    IssueTokenRequest_AuthType = 1
	IssueTokenRequest_EXTERNAL IssueTokenRequest_AuthType = 2
	IssueTokenRequest_MFA      IssueTokenRequest_AuthType = 3
)

// Enum value maps for IssueTokenRequest_AuthType.
var (
	IssueTokenRequest_AuthType_name = map[int32]string{
		0: "NONE",
		1: "LOCAL",
		2: "EXTERNAL",
		3: "MFA",
	}
	IssueTokenRequest_AuthType_value = map[string]int32{
		"NONE":     0,
		"LOCAL":    1,
		"EXTERNAL": 2,
		"MFA":      3,
	}
)

func (x IssueTokenRequest_AuthType) Enum() *IssueTokenRequest_AuthType {
	p := new(IssueTokenRequest_AuthType)
	*p = x
	return p
}

func (x IssueTokenRequest_AuthType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IssueTokenRequest_AuthType) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_identity_v2_token_proto_enumTypes[0].Descriptor()
}

func (IssueTokenRequest_AuthType) Type() protoreflect.EnumType {
	return &file_spaceone_api_identity_v2_token_proto_enumTypes[0]
}

func (x IssueTokenRequest_AuthType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IssueTokenRequest_AuthType.Descriptor instead.
func (IssueTokenRequest_AuthType) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_token_proto_rawDescGZIP(), []int{0, 0}
}

type GrantTokenRequest_GrantType int32

const (
	GrantTokenRequest_GRANT_TYPE_NONE GrantTokenRequest_GrantType = 0
	GrantTokenRequest_REFRESH_TOKEN   GrantTokenRequest_GrantType = 1
	GrantTokenRequest_SYSTEM_TOKEN    GrantTokenRequest_GrantType = 2
)

// Enum value maps for GrantTokenRequest_GrantType.
var (
	GrantTokenRequest_GrantType_name = map[int32]string{
		0: "GRANT_TYPE_NONE",
		1: "REFRESH_TOKEN",
		2: "SYSTEM_TOKEN",
	}
	GrantTokenRequest_GrantType_value = map[string]int32{
		"GRANT_TYPE_NONE": 0,
		"REFRESH_TOKEN":   1,
		"SYSTEM_TOKEN":    2,
	}
)

func (x GrantTokenRequest_GrantType) Enum() *GrantTokenRequest_GrantType {
	p := new(GrantTokenRequest_GrantType)
	*p = x
	return p
}

func (x GrantTokenRequest_GrantType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GrantTokenRequest_GrantType) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_identity_v2_token_proto_enumTypes[1].Descriptor()
}

func (GrantTokenRequest_GrantType) Type() protoreflect.EnumType {
	return &file_spaceone_api_identity_v2_token_proto_enumTypes[1]
}

func (x GrantTokenRequest_GrantType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GrantTokenRequest_GrantType.Descriptor instead.
func (GrantTokenRequest_GrantType) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_token_proto_rawDescGZIP(), []int{2, 0}
}

type GrantTokenRequest_Scope int32

const (
	GrantTokenRequest_SCOPE_NONE GrantTokenRequest_Scope = 0
	GrantTokenRequest_SYSTEM     GrantTokenRequest_Scope = 1
	GrantTokenRequest_DOMAIN     GrantTokenRequest_Scope = 2
	GrantTokenRequest_WORKSPACE  GrantTokenRequest_Scope = 3
	GrantTokenRequest_PROJECT    GrantTokenRequest_Scope = 4
	GrantTokenRequest_USER       GrantTokenRequest_Scope = 5
)

// Enum value maps for GrantTokenRequest_Scope.
var (
	GrantTokenRequest_Scope_name = map[int32]string{
		0: "SCOPE_NONE",
		1: "SYSTEM",
		2: "DOMAIN",
		3: "WORKSPACE",
		4: "PROJECT",
		5: "USER",
	}
	GrantTokenRequest_Scope_value = map[string]int32{
		"SCOPE_NONE": 0,
		"SYSTEM":     1,
		"DOMAIN":     2,
		"WORKSPACE":  3,
		"PROJECT":    4,
		"USER":       5,
	}
)

func (x GrantTokenRequest_Scope) Enum() *GrantTokenRequest_Scope {
	p := new(GrantTokenRequest_Scope)
	*p = x
	return p
}

func (x GrantTokenRequest_Scope) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GrantTokenRequest_Scope) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_identity_v2_token_proto_enumTypes[2].Descriptor()
}

func (GrantTokenRequest_Scope) Type() protoreflect.EnumType {
	return &file_spaceone_api_identity_v2_token_proto_enumTypes[2]
}

func (x GrantTokenRequest_Scope) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GrantTokenRequest_Scope.Descriptor instead.
func (GrantTokenRequest_Scope) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_token_proto_rawDescGZIP(), []int{2, 1}
}

type GrantTokenInfo_RoleType int32

const (
	GrantTokenInfo_ROLE_TYPE_NONE   GrantTokenInfo_RoleType = 0
	GrantTokenInfo_DOMAIN_ADMIN     GrantTokenInfo_RoleType = 1
	GrantTokenInfo_WORKSPACE_OWNER  GrantTokenInfo_RoleType = 2
	GrantTokenInfo_WORKSPACE_MEMBER GrantTokenInfo_RoleType = 3
	GrantTokenInfo_USER             GrantTokenInfo_RoleType = 4
)

// Enum value maps for GrantTokenInfo_RoleType.
var (
	GrantTokenInfo_RoleType_name = map[int32]string{
		0: "ROLE_TYPE_NONE",
		1: "DOMAIN_ADMIN",
		2: "WORKSPACE_OWNER",
		3: "WORKSPACE_MEMBER",
		4: "USER",
	}
	GrantTokenInfo_RoleType_value = map[string]int32{
		"ROLE_TYPE_NONE":   0,
		"DOMAIN_ADMIN":     1,
		"WORKSPACE_OWNER":  2,
		"WORKSPACE_MEMBER": 3,
		"USER":             4,
	}
)

func (x GrantTokenInfo_RoleType) Enum() *GrantTokenInfo_RoleType {
	p := new(GrantTokenInfo_RoleType)
	*p = x
	return p
}

func (x GrantTokenInfo_RoleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GrantTokenInfo_RoleType) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_identity_v2_token_proto_enumTypes[3].Descriptor()
}

func (GrantTokenInfo_RoleType) Type() protoreflect.EnumType {
	return &file_spaceone_api_identity_v2_token_proto_enumTypes[3]
}

func (x GrantTokenInfo_RoleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GrantTokenInfo_RoleType.Descriptor instead.
func (GrantTokenInfo_RoleType) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_token_proto_rawDescGZIP(), []int{3, 0}
}

//	{
//	 "credentials": {
//	     "user_id": "wonny@cloudforet.io",
//	     "password": "Password0*"
//	 },
//	 "auth_type": "LOCAL",
//	 "domain_id": "domain-123123a"
//	}
type IssueTokenRequest struct {
	state       protoimpl.MessageState `protogen:"open.v1"`
	Credentials *_struct.Struct        `protobuf:"bytes,1,opt,name=credentials,proto3" json:"credentials,omitempty"`
	// LOCAL, EXTERNAL
	// +optional
	AuthType IssueTokenRequest_AuthType `protobuf:"varint,2,opt,name=auth_type,json=authType,proto3,enum=spaceone.api.identity.v2.IssueTokenRequest_AuthType" json:"auth_type,omitempty"`
	// +optional
	Timeout int32 `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// if MFA is enabled, verify_code is required
	// +optional
	VerifyCode    string `protobuf:"bytes,4,opt,name=verify_code,json=verifyCode,proto3" json:"verify_code,omitempty"`
	DomainId      string `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IssueTokenRequest) Reset() {
	*x = IssueTokenRequest{}
	mi := &file_spaceone_api_identity_v2_token_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IssueTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueTokenRequest) ProtoMessage() {}

func (x *IssueTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_token_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueTokenRequest.ProtoReflect.Descriptor instead.
func (*IssueTokenRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_token_proto_rawDescGZIP(), []int{0}
}

func (x *IssueTokenRequest) GetCredentials() *_struct.Struct {
	if x != nil {
		return x.Credentials
	}
	return nil
}

func (x *IssueTokenRequest) GetAuthType() IssueTokenRequest_AuthType {
	if x != nil {
		return x.AuthType
	}
	return IssueTokenRequest_NONE
}

func (x *IssueTokenRequest) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *IssueTokenRequest) GetVerifyCode() string {
	if x != nil {
		return x.VerifyCode
	}
	return ""
}

func (x *IssueTokenRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type TokenInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken  string                 `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TokenInfo) Reset() {
	*x = TokenInfo{}
	mi := &file_spaceone_api_identity_v2_token_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TokenInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenInfo) ProtoMessage() {}

func (x *TokenInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_token_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenInfo.ProtoReflect.Descriptor instead.
func (*TokenInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_token_proto_rawDescGZIP(), []int{1}
}

func (x *TokenInfo) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *TokenInfo) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

//	{
//	 "grant_type": "REFRESH_TOKEN",
//	 "token": "your_refresh_token_from_issue",
//	 "scope": "DOMAIN",
//	 "domain_id": "domain-123123a"
//	}
type GrantTokenRequest struct {
	state     protoimpl.MessageState      `protogen:"open.v1"`
	GrantType GrantTokenRequest_GrantType `protobuf:"varint,1,opt,name=grant_type,json=grantType,proto3,enum=spaceone.api.identity.v2.GrantTokenRequest_GrantType" json:"grant_type,omitempty"`
	Token     string                      `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Scope     GrantTokenRequest_Scope     `protobuf:"varint,3,opt,name=scope,proto3,enum=spaceone.api.identity.v2.GrantTokenRequest_Scope" json:"scope,omitempty"`
	// +optional
	Timeout int32 `protobuf:"varint,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// +optional
	Permissions []string `protobuf:"bytes,5,rep,name=permissions,proto3" json:"permissions,omitempty"`
	// +optional
	DomainId string `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	// +optional
	WorkspaceId   string `protobuf:"bytes,22,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GrantTokenRequest) Reset() {
	*x = GrantTokenRequest{}
	mi := &file_spaceone_api_identity_v2_token_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GrantTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrantTokenRequest) ProtoMessage() {}

func (x *GrantTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_token_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrantTokenRequest.ProtoReflect.Descriptor instead.
func (*GrantTokenRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_token_proto_rawDescGZIP(), []int{2}
}

func (x *GrantTokenRequest) GetGrantType() GrantTokenRequest_GrantType {
	if x != nil {
		return x.GrantType
	}
	return GrantTokenRequest_GRANT_TYPE_NONE
}

func (x *GrantTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GrantTokenRequest) GetScope() GrantTokenRequest_Scope {
	if x != nil {
		return x.Scope
	}
	return GrantTokenRequest_SCOPE_NONE
}

func (x *GrantTokenRequest) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *GrantTokenRequest) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *GrantTokenRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *GrantTokenRequest) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

type GrantTokenInfo struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	AccessToken   string                  `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RoleType      GrantTokenInfo_RoleType `protobuf:"varint,2,opt,name=role_type,json=roleType,proto3,enum=spaceone.api.identity.v2.GrantTokenInfo_RoleType" json:"role_type,omitempty"`
	DomainId      string                  `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	WorkspaceId   string                  `protobuf:"bytes,22,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	RoleId        string                  `protobuf:"bytes,23,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GrantTokenInfo) Reset() {
	*x = GrantTokenInfo{}
	mi := &file_spaceone_api_identity_v2_token_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GrantTokenInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrantTokenInfo) ProtoMessage() {}

func (x *GrantTokenInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_token_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrantTokenInfo.ProtoReflect.Descriptor instead.
func (*GrantTokenInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_token_proto_rawDescGZIP(), []int{3}
}

func (x *GrantTokenInfo) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *GrantTokenInfo) GetRoleType() GrantTokenInfo_RoleType {
	if x != nil {
		return x.RoleType
	}
	return GrantTokenInfo_ROLE_TYPE_NONE
}

func (x *GrantTokenInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *GrantTokenInfo) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *GrantTokenInfo) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

var File_spaceone_api_identity_v2_token_proto protoreflect.FileDescriptor

var file_spaceone_api_identity_v2_token_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x02, 0x0a, 0x11, 0x49, 0x73,
	0x73, 0x75, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x39, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0b, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x51, 0x0a, 0x09, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x36, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4c,
	0x4f, 0x43, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e,
	0x41, 0x4c, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x46, 0x41, 0x10, 0x03, 0x22, 0x53, 0x0a,
	0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a,
	0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0xe2, 0x03, 0x0a, 0x11, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x54, 0x0a, 0x0a, 0x67, 0x72, 0x61, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x09, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x47, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x47,
	0x72, 0x61, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x09, 0x47, 0x72, 0x61,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x47, 0x52, 0x41, 0x4e, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x52,
	0x45, 0x46, 0x52, 0x45, 0x53, 0x48, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x01, 0x12, 0x10,
	0x0a, 0x0c, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x02,
	0x22, 0x55, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x43, 0x4f,
	0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x59, 0x53,
	0x54, 0x45, 0x4d, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x10,
	0x02, 0x12, 0x0d, 0x0a, 0x09, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41, 0x43, 0x45, 0x10, 0x03,
	0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x04, 0x12, 0x08, 0x0a,
	0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x05, 0x22, 0xc3, 0x02, 0x0a, 0x0e, 0x47, 0x72, 0x61, 0x6e,
	0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x4e, 0x0a,
	0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x31, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x72, 0x61, 0x6e,
	0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x65, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e,
	0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x57, 0x4f, 0x52, 0x4b,
	0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x10, 0x02, 0x12, 0x14, 0x0a,
	0x10, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45,
	0x52, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x04, 0x32, 0x8d, 0x02,
	0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x7e, 0x0a, 0x05, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x12, 0x2b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x73, 0x73, 0x75,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x12, 0x83, 0x01, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x6e,
	0x74, 0x12, 0x2b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x72, 0x61,
	0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d,
	0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76,
	0x32, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x42, 0x3f, 0x5a,
	0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69,
	0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_identity_v2_token_proto_rawDescOnce sync.Once
	file_spaceone_api_identity_v2_token_proto_rawDescData = file_spaceone_api_identity_v2_token_proto_rawDesc
)

func file_spaceone_api_identity_v2_token_proto_rawDescGZIP() []byte {
	file_spaceone_api_identity_v2_token_proto_rawDescOnce.Do(func() {
		file_spaceone_api_identity_v2_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_identity_v2_token_proto_rawDescData)
	})
	return file_spaceone_api_identity_v2_token_proto_rawDescData
}

var file_spaceone_api_identity_v2_token_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_spaceone_api_identity_v2_token_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_spaceone_api_identity_v2_token_proto_goTypes = []any{
	(IssueTokenRequest_AuthType)(0),  // 0: spaceone.api.identity.v2.IssueTokenRequest.AuthType
	(GrantTokenRequest_GrantType)(0), // 1: spaceone.api.identity.v2.GrantTokenRequest.GrantType
	(GrantTokenRequest_Scope)(0),     // 2: spaceone.api.identity.v2.GrantTokenRequest.Scope
	(GrantTokenInfo_RoleType)(0),     // 3: spaceone.api.identity.v2.GrantTokenInfo.RoleType
	(*IssueTokenRequest)(nil),        // 4: spaceone.api.identity.v2.IssueTokenRequest
	(*TokenInfo)(nil),                // 5: spaceone.api.identity.v2.TokenInfo
	(*GrantTokenRequest)(nil),        // 6: spaceone.api.identity.v2.GrantTokenRequest
	(*GrantTokenInfo)(nil),           // 7: spaceone.api.identity.v2.GrantTokenInfo
	(*_struct.Struct)(nil),           // 8: google.protobuf.Struct
}
var file_spaceone_api_identity_v2_token_proto_depIdxs = []int32{
	8, // 0: spaceone.api.identity.v2.IssueTokenRequest.credentials:type_name -> google.protobuf.Struct
	0, // 1: spaceone.api.identity.v2.IssueTokenRequest.auth_type:type_name -> spaceone.api.identity.v2.IssueTokenRequest.AuthType
	1, // 2: spaceone.api.identity.v2.GrantTokenRequest.grant_type:type_name -> spaceone.api.identity.v2.GrantTokenRequest.GrantType
	2, // 3: spaceone.api.identity.v2.GrantTokenRequest.scope:type_name -> spaceone.api.identity.v2.GrantTokenRequest.Scope
	3, // 4: spaceone.api.identity.v2.GrantTokenInfo.role_type:type_name -> spaceone.api.identity.v2.GrantTokenInfo.RoleType
	4, // 5: spaceone.api.identity.v2.Token.issue:input_type -> spaceone.api.identity.v2.IssueTokenRequest
	6, // 6: spaceone.api.identity.v2.Token.grant:input_type -> spaceone.api.identity.v2.GrantTokenRequest
	5, // 7: spaceone.api.identity.v2.Token.issue:output_type -> spaceone.api.identity.v2.TokenInfo
	7, // 8: spaceone.api.identity.v2.Token.grant:output_type -> spaceone.api.identity.v2.GrantTokenInfo
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_spaceone_api_identity_v2_token_proto_init() }
func file_spaceone_api_identity_v2_token_proto_init() {
	if File_spaceone_api_identity_v2_token_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spaceone_api_identity_v2_token_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_identity_v2_token_proto_goTypes,
		DependencyIndexes: file_spaceone_api_identity_v2_token_proto_depIdxs,
		EnumInfos:         file_spaceone_api_identity_v2_token_proto_enumTypes,
		MessageInfos:      file_spaceone_api_identity_v2_token_proto_msgTypes,
	}.Build()
	File_spaceone_api_identity_v2_token_proto = out.File
	file_spaceone_api_identity_v2_token_proto_rawDesc = nil
	file_spaceone_api_identity_v2_token_proto_goTypes = nil
	file_spaceone_api_identity_v2_token_proto_depIdxs = nil
}
