# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/identity/v1/policy.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from spaceone.api.core.v1 import query_pb2 as spaceone_dot_api_dot_core_dot_v1_dot_query__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n%spaceone/api/identity/v1/policy.proto\x12\x18spaceone.api.identity.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v1/query.proto\"r\n\x13\x43reatePolicyRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x13\n\x0bpermissions\x18\x02 \x03(\t\x12%\n\x04tags\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x04 \x01(\t\"\x85\x01\n\x13UpdatePolicyRequest\x12\x11\n\tpolicy_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0bpermissions\x18\x03 \x03(\t\x12%\n\x04tags\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x05 \x01(\t\"5\n\rPolicyRequest\x12\x11\n\tpolicy_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x02 \x01(\t\"F\n\x10GetPolicyRequest\x12\x11\n\tpolicy_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x02 \x01(\t\x12\x0c\n\x04only\x18\x03 \x03(\t\"m\n\x0bPolicyQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v1.Query\x12\x11\n\tpolicy_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x11\n\tdomain_id\x18\x04 \x01(\t\"\x90\x01\n\nPolicyInfo\x12\x11\n\tpolicy_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0bpermissions\x18\x03 \x03(\t\x12\x11\n\tdomain_id\x18\x04 \x01(\t\x12%\n\x04tags\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x12\n\ncreated_at\x18\x06 \x01(\t\"Z\n\x0cPoliciesInfo\x12\x35\n\x07results\x18\x01 \x03(\x0b\x32$.spaceone.api.identity.v1.PolicyInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"Z\n\x0fPolicyStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v1.StatisticsQuery\x12\x11\n\tdomain_id\x18\x02 \x01(\t2\xf4\x05\n\x06Policy\x12\x84\x01\n\x06\x63reate\x12-.spaceone.api.identity.v1.CreatePolicyRequest\x1a$.spaceone.api.identity.v1.PolicyInfo\"%\x82\xd3\xe4\x93\x02\x1f\"\x1a/identity/v1/policy/create:\x01*\x12\x84\x01\n\x06update\x12-.spaceone.api.identity.v1.UpdatePolicyRequest\x1a$.spaceone.api.identity.v1.PolicyInfo\"%\x82\xd3\xe4\x93\x02\x1f\x1a\x1a/identity/v1/policy/update:\x01*\x12p\n\x06\x64\x65lete\x12\'.spaceone.api.identity.v1.PolicyRequest\x1a\x16.google.protobuf.Empty\"%\x82\xd3\xe4\x93\x02\x1f\"\x1a/identity/v1/policy/delete:\x01*\x12{\n\x03get\x12*.spaceone.api.identity.v1.GetPolicyRequest\x1a$.spaceone.api.identity.v1.PolicyInfo\"\"\x82\xd3\xe4\x93\x02\x1c\"\x17/identity/v1/policy/get:\x01*\x12z\n\x04list\x12%.spaceone.api.identity.v1.PolicyQuery\x1a&.spaceone.api.identity.v1.PoliciesInfo\"#\x82\xd3\xe4\x93\x02\x1d\"\x18/identity/v1/policy/list:\x01*\x12q\n\x04stat\x12).spaceone.api.identity.v1.PolicyStatQuery\x1a\x17.google.protobuf.Struct\"%\x82\xd3\xe4\x93\x02\x1f\"\x1a/identity/v1/policies/stat:\x01*B?Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v1b\x06proto3')



_CREATEPOLICYREQUEST = DESCRIPTOR.message_types_by_name['CreatePolicyRequest']
_UPDATEPOLICYREQUEST = DESCRIPTOR.message_types_by_name['UpdatePolicyRequest']
_POLICYREQUEST = DESCRIPTOR.message_types_by_name['PolicyRequest']
_GETPOLICYREQUEST = DESCRIPTOR.message_types_by_name['GetPolicyRequest']
_POLICYQUERY = DESCRIPTOR.message_types_by_name['PolicyQuery']
_POLICYINFO = DESCRIPTOR.message_types_by_name['PolicyInfo']
_POLICIESINFO = DESCRIPTOR.message_types_by_name['PoliciesInfo']
_POLICYSTATQUERY = DESCRIPTOR.message_types_by_name['PolicyStatQuery']
CreatePolicyRequest = _reflection.GeneratedProtocolMessageType('CreatePolicyRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEPOLICYREQUEST,
  '__module__' : 'spaceone.api.identity.v1.policy_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.identity.v1.CreatePolicyRequest)
  })
_sym_db.RegisterMessage(CreatePolicyRequest)

UpdatePolicyRequest = _reflection.GeneratedProtocolMessageType('UpdatePolicyRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPDATEPOLICYREQUEST,
  '__module__' : 'spaceone.api.identity.v1.policy_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.identity.v1.UpdatePolicyRequest)
  })
_sym_db.RegisterMessage(UpdatePolicyRequest)

PolicyRequest = _reflection.GeneratedProtocolMessageType('PolicyRequest', (_message.Message,), {
  'DESCRIPTOR' : _POLICYREQUEST,
  '__module__' : 'spaceone.api.identity.v1.policy_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.identity.v1.PolicyRequest)
  })
_sym_db.RegisterMessage(PolicyRequest)

GetPolicyRequest = _reflection.GeneratedProtocolMessageType('GetPolicyRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETPOLICYREQUEST,
  '__module__' : 'spaceone.api.identity.v1.policy_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.identity.v1.GetPolicyRequest)
  })
_sym_db.RegisterMessage(GetPolicyRequest)

PolicyQuery = _reflection.GeneratedProtocolMessageType('PolicyQuery', (_message.Message,), {
  'DESCRIPTOR' : _POLICYQUERY,
  '__module__' : 'spaceone.api.identity.v1.policy_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.identity.v1.PolicyQuery)
  })
_sym_db.RegisterMessage(PolicyQuery)

PolicyInfo = _reflection.GeneratedProtocolMessageType('PolicyInfo', (_message.Message,), {
  'DESCRIPTOR' : _POLICYINFO,
  '__module__' : 'spaceone.api.identity.v1.policy_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.identity.v1.PolicyInfo)
  })
_sym_db.RegisterMessage(PolicyInfo)

PoliciesInfo = _reflection.GeneratedProtocolMessageType('PoliciesInfo', (_message.Message,), {
  'DESCRIPTOR' : _POLICIESINFO,
  '__module__' : 'spaceone.api.identity.v1.policy_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.identity.v1.PoliciesInfo)
  })
_sym_db.RegisterMessage(PoliciesInfo)

PolicyStatQuery = _reflection.GeneratedProtocolMessageType('PolicyStatQuery', (_message.Message,), {
  'DESCRIPTOR' : _POLICYSTATQUERY,
  '__module__' : 'spaceone.api.identity.v1.policy_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.identity.v1.PolicyStatQuery)
  })
_sym_db.RegisterMessage(PolicyStatQuery)

_POLICY = DESCRIPTOR.services_by_name['Policy']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v1'
  _POLICY.methods_by_name['create']._options = None
  _POLICY.methods_by_name['create']._serialized_options = b'\202\323\344\223\002\037\"\032/identity/v1/policy/create:\001*'
  _POLICY.methods_by_name['update']._options = None
  _POLICY.methods_by_name['update']._serialized_options = b'\202\323\344\223\002\037\032\032/identity/v1/policy/update:\001*'
  _POLICY.methods_by_name['delete']._options = None
  _POLICY.methods_by_name['delete']._serialized_options = b'\202\323\344\223\002\037\"\032/identity/v1/policy/delete:\001*'
  _POLICY.methods_by_name['get']._options = None
  _POLICY.methods_by_name['get']._serialized_options = b'\202\323\344\223\002\034\"\027/identity/v1/policy/get:\001*'
  _POLICY.methods_by_name['list']._options = None
  _POLICY.methods_by_name['list']._serialized_options = b'\202\323\344\223\002\035\"\030/identity/v1/policy/list:\001*'
  _POLICY.methods_by_name['stat']._options = None
  _POLICY.methods_by_name['stat']._serialized_options = b'\202\323\344\223\002\037\"\032/identity/v1/policies/stat:\001*'
  _CREATEPOLICYREQUEST._serialized_start=190
  _CREATEPOLICYREQUEST._serialized_end=304
  _UPDATEPOLICYREQUEST._serialized_start=307
  _UPDATEPOLICYREQUEST._serialized_end=440
  _POLICYREQUEST._serialized_start=442
  _POLICYREQUEST._serialized_end=495
  _GETPOLICYREQUEST._serialized_start=497
  _GETPOLICYREQUEST._serialized_end=567
  _POLICYQUERY._serialized_start=569
  _POLICYQUERY._serialized_end=678
  _POLICYINFO._serialized_start=681
  _POLICYINFO._serialized_end=825
  _POLICIESINFO._serialized_start=827
  _POLICIESINFO._serialized_end=917
  _POLICYSTATQUERY._serialized_start=919
  _POLICYSTATQUERY._serialized_end=1009
  _POLICY._serialized_start=1012
  _POLICY._serialized_end=1768
# @@protoc_insertion_point(module_scope)
