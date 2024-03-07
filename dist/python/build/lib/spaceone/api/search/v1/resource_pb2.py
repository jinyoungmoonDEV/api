# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/search/v1/resource.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from spaceone.api.core.v2 import query_pb2 as spaceone_dot_api_dot_core_dot_v2_dot_query__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n%spaceone/api/search/v1/resource.proto\x12\x16spaceone.api.search.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"\x8e\x01\n\x15SearchResourceRequest\x12\x15\n\rresource_type\x18\x01 \x01(\t\x12\x0f\n\x07keyword\x18\x02 \x01(\t\x12\r\n\x05limit\x18\x03 \x01(\x05\x12\x12\n\nworkspaces\x18\x04 \x03(\t\x12\x16\n\x0e\x61ll_workspaces\x18\x05 \x01(\x08\x12\x12\n\nnext_token\x18\x06 \x01(\t\"\xaa\x01\n\x0cResourceInfo\x12\x13\n\x0bresource_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12%\n\x04tags\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\x12\x12\n\nproject_id\x18\x17 \x01(\t\"Z\n\rResourcesInfo\x12\x35\n\x07results\x18\x01 \x03(\x0b\x32$.spaceone.api.search.v1.ResourceInfo\x12\x12\n\nnext_token\x18\x02 \x01(\t2\x92\x01\n\x08Resource\x12\x85\x01\n\x06search\x12-.spaceone.api.search.v1.SearchResourceRequest\x1a%.spaceone.api.search.v1.ResourcesInfo\"%\x82\xd3\xe4\x93\x02\x1f\"\x1a/search/v1/resource/search:\x01*B=Z;github.com/cloudforet-io/api/dist/go/spaceone/api/search/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.search.v1.resource_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z;github.com/cloudforet-io/api/dist/go/spaceone/api/search/v1'
  _globals['_RESOURCE'].methods_by_name['search']._options = None
  _globals['_RESOURCE'].methods_by_name['search']._serialized_options = b'\202\323\344\223\002\037\"\032/search/v1/resource/search:\001*'
  _globals['_SEARCHRESOURCEREQUEST']._serialized_start=189
  _globals['_SEARCHRESOURCEREQUEST']._serialized_end=331
  _globals['_RESOURCEINFO']._serialized_start=334
  _globals['_RESOURCEINFO']._serialized_end=504
  _globals['_RESOURCESINFO']._serialized_start=506
  _globals['_RESOURCESINFO']._serialized_end=596
  _globals['_RESOURCE']._serialized_start=599
  _globals['_RESOURCE']._serialized_end=745
# @@protoc_insertion_point(module_scope)
