# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/alert_manager/plugin/protocol.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n0spaceone/api/alert_manager/plugin/protocol.proto\x12!spaceone.api.alert_manager.plugin\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\"J\n\x0bInitRequest\x12(\n\x07options\x18\x01 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\"z\n\rVerifyRequest\x12(\n\x07options\x18\x01 \x01(\x0b\x32\x17.google.protobuf.Struct\x12,\n\x0bsecret_data\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\"7\n\nPluginInfo\x12)\n\x08metadata\x18\x01 \x01(\x0b\x32\x17.google.protobuf.Struct2\xc9\x01\n\x08Protocol\x12g\n\x04init\x12..spaceone.api.alert_manager.plugin.InitRequest\x1a-.spaceone.api.alert_manager.plugin.PluginInfo\"\x00\x12T\n\x06verify\x12\x30.spaceone.api.alert_manager.plugin.VerifyRequest\x1a\x16.google.protobuf.Empty\"\x00\x42HZFgithub.com/cloudforet-io/api/dist/go/spaceone/api/alert-manager/pluginb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.alert_manager.plugin.protocol_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZFgithub.com/cloudforet-io/api/dist/go/spaceone/api/alert-manager/plugin'
  _globals['_INITREQUEST']._serialized_start=146
  _globals['_INITREQUEST']._serialized_end=220
  _globals['_VERIFYREQUEST']._serialized_start=222
  _globals['_VERIFYREQUEST']._serialized_end=344
  _globals['_PLUGININFO']._serialized_start=346
  _globals['_PLUGININFO']._serialized_end=401
  _globals['_PROTOCOL']._serialized_start=404
  _globals['_PROTOCOL']._serialized_end=605
# @@protoc_insertion_point(module_scope)
