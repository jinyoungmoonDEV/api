# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/cost_analysis/plugin/data_source.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n3spaceone/api/cost_analysis/plugin/data_source.proto\x12!spaceone.api.cost_analysis.plugin\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\"J\n\x0bInitRequest\x12(\n\x07options\x18\x01 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x02 \x01(\t\"\x90\x01\n\x13PluginVerifyRequest\x12(\n\x07options\x18\x01 \x01(\x0b\x32\x17.google.protobuf.Struct\x12,\n\x0bsecret_data\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x0e\n\x06schema\x18\x03 \x01(\t\x12\x11\n\tdomain_id\x18\x04 \x01(\t\"7\n\nPluginInfo\x12)\n\x08metadata\x18\x01 \x01(\x0b\x32\x17.google.protobuf.Struct2\xd1\x01\n\nDataSource\x12g\n\x04init\x12..spaceone.api.cost_analysis.plugin.InitRequest\x1a-.spaceone.api.cost_analysis.plugin.PluginInfo\"\x00\x12Z\n\x06verify\x12\x36.spaceone.api.cost_analysis.plugin.PluginVerifyRequest\x1a\x16.google.protobuf.Empty\"\x00\x42HZFgithub.com/cloudforet-io/api/dist/go/spaceone/api/cost_analysis/pluginb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.cost_analysis.plugin.data_source_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZFgithub.com/cloudforet-io/api/dist/go/spaceone/api/cost_analysis/plugin'
  _globals['_INITREQUEST']._serialized_start=149
  _globals['_INITREQUEST']._serialized_end=223
  _globals['_PLUGINVERIFYREQUEST']._serialized_start=226
  _globals['_PLUGINVERIFYREQUEST']._serialized_end=370
  _globals['_PLUGININFO']._serialized_start=372
  _globals['_PLUGININFO']._serialized_end=427
  _globals['_DATASOURCE']._serialized_start=430
  _globals['_DATASOURCE']._serialized_end=639
# @@protoc_insertion_point(module_scope)
