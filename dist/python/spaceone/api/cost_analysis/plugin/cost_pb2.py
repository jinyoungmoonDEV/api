# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/cost_analysis/plugin/cost.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n,spaceone/api/cost_analysis/plugin/cost.proto\x12!spaceone.api.cost_analysis.plugin\x1a\x1cgoogle/protobuf/struct.proto\"\xba\x01\n\x0eGetDataRequest\x12(\n\x07options\x18\x01 \x01(\x0b\x32\x17.google.protobuf.Struct\x12,\n\x0bsecret_data\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x0e\n\x06schema\x18\x03 \x01(\t\x12-\n\x0ctask_options\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x05 \x01(\t\"\xb7\x02\n\x08\x43ostInfo\x12\x0c\n\x04\x63ost\x18\x01 \x01(\x01\x12\x16\n\x0eusage_quantity\x18\x02 \x01(\x01\x12\x12\n\nusage_unit\x18\x03 \x01(\t\x12\x10\n\x08provider\x18\x04 \x01(\t\x12\x13\n\x0bregion_code\x18\x05 \x01(\t\x12\x0f\n\x07product\x18\x06 \x01(\t\x12\x12\n\nusage_type\x18\x07 \x01(\t\x12\x10\n\x08resource\x18\x08 \x01(\t\x12%\n\x04tags\x18\x0c \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x30\n\x0f\x61\x64\x64itional_info\x18\r \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04\x64\x61ta\x18\x0e \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x13\n\x0b\x62illed_date\x18\x15 \x01(\t\"I\n\tCostsInfo\x12<\n\x07results\x18\x01 \x03(\x0b\x32+.spaceone.api.cost_analysis.plugin.CostInfo2w\n\x04\x43ost\x12o\n\x08get_data\x12\x31.spaceone.api.cost_analysis.plugin.GetDataRequest\x1a,.spaceone.api.cost_analysis.plugin.CostsInfo\"\x00\x30\x01\x42HZFgithub.com/cloudforet-io/api/dist/go/spaceone/api/cost_analysis/pluginb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.cost_analysis.plugin.cost_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZFgithub.com/cloudforet-io/api/dist/go/spaceone/api/cost_analysis/plugin'
  _globals['_GETDATAREQUEST']._serialized_start=114
  _globals['_GETDATAREQUEST']._serialized_end=300
  _globals['_COSTINFO']._serialized_start=303
  _globals['_COSTINFO']._serialized_end=614
  _globals['_COSTSINFO']._serialized_start=616
  _globals['_COSTSINFO']._serialized_end=689
  _globals['_COST']._serialized_start=691
  _globals['_COST']._serialized_end=810
# @@protoc_insertion_point(module_scope)
