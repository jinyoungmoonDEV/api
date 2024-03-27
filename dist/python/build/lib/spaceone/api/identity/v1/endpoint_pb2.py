# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/identity/v1/endpoint.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from spaceone.api.core.v1 import query_pb2 as spaceone_dot_api_dot_core_dot_v1_dot_query__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\'spaceone/api/identity/v1/endpoint.proto\x12\x18spaceone.api.identity.v1\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v1/query.proto\"\xca\x01\n\x0c\x45ndpointInfo\x12\x0f\n\x07service\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x10\n\x08\x65ndpoint\x18\x03 \x01(\t\x12\x43\n\x05state\x18\x04 \x01(\x0e\x32\x34.spaceone.api.identity.v1.EndpointInfo.EndpointState\x12\x0f\n\x07version\x18\x05 \x01(\t\"3\n\rEndpointState\x12\x08\n\x04NONE\x10\x00\x12\n\n\x06\x41\x43TIVE\x10\x01\x12\x0c\n\x08INACTIVE\x10\x02\"v\n\rEndpointQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v1.Query\x12\x0f\n\x07service\x18\x02 \x01(\t\x12\x15\n\rendpoint_type\x18\x03 \x01(\t\x12\x11\n\tdomain_id\x18\x06 \x01(\t\"]\n\rEndpointsInfo\x12\x37\n\x07results\x18\x01 \x03(\x0b\x32&.spaceone.api.identity.v1.EndpointInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\x32\x8b\x01\n\x08\x45ndpoint\x12\x7f\n\x04list\x12\'.spaceone.api.identity.v1.EndpointQuery\x1a\'.spaceone.api.identity.v1.EndpointsInfo\"%\x82\xd3\xe4\x93\x02\x1f\"\x1a/identity/v1/endpoint/list:\x01*B?Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.identity.v1.endpoint_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v1'
  _globals['_ENDPOINT'].methods_by_name['list']._options = None
  _globals['_ENDPOINT'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002\037\"\032/identity/v1/endpoint/list:\001*'
  _globals['_ENDPOINTINFO']._serialized_start=134
  _globals['_ENDPOINTINFO']._serialized_end=336
  _globals['_ENDPOINTINFO_ENDPOINTSTATE']._serialized_start=285
  _globals['_ENDPOINTINFO_ENDPOINTSTATE']._serialized_end=336
  _globals['_ENDPOINTQUERY']._serialized_start=338
  _globals['_ENDPOINTQUERY']._serialized_end=456
  _globals['_ENDPOINTSINFO']._serialized_start=458
  _globals['_ENDPOINTSINFO']._serialized_end=551
  _globals['_ENDPOINT']._serialized_start=554
  _globals['_ENDPOINT']._serialized_end=693
# @@protoc_insertion_point(module_scope)
