# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/alert_manager/v1/event.proto
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
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from spaceone.api.core.v2 import query_pb2 as spaceone_dot_api_dot_core_dot_v2_dot_query__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n)spaceone/api/alert_manager/v1/event.proto\x12\x1dspaceone.api.alert_manager.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"c\n\x12\x45ventCreateRequest\x12\x12\n\nwebhook_id\x18\x01 \x01(\t\x12\x12\n\naccess_key\x18\x02 \x01(\t\x12%\n\x04\x64\x61ta\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\"\xe2\x04\n\tEventInfo\x12\x10\n\x08\x65vent_id\x18\x01 \x01(\t\x12\x11\n\tevent_key\x18\x02 \x01(\t\x12\x46\n\nevent_type\x18\x03 \x01(\x0e\x32\x32.spaceone.api.alert_manager.v1.EventInfo.EventType\x12\r\n\x05title\x18\x04 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x05 \x01(\t\x12>\n\x08severity\x18\x06 \x01(\x0e\x32,.spaceone.api.alert_manager.v1.EventSeverity\x12\x0c\n\x04rule\x18\x07 \x01(\t\x12\x11\n\timage_url\x18\x08 \x01(\t\x12\x11\n\tresources\x18\t \x03(\t\x12\x10\n\x08provider\x18\n \x01(\t\x12\x0f\n\x07\x61\x63\x63ount\x18\x0b \x01(\t\x12\x30\n\x0f\x61\x64\x64itional_info\x18\x0f \x01(\x0b\x32\x17.google.protobuf.Struct\x12)\n\x08raw_data\x18\x10 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\x12\x12\n\nservice_id\x18\x17 \x01(\t\x12\x12\n\nwebhook_id\x18\x18 \x01(\t\x12\x10\n\x08\x61lert_id\x18\x19 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x13\n\x0boccurred_at\x18  \x01(\t\"D\n\tEventType\x12\x13\n\x0f\x45VENT_TYPE_NONE\x10\x00\x12\t\n\x05\x41LERT\x10\x01\x12\x0c\n\x08RECOVERY\x10\x02\x12\t\n\x05\x45RROR\x10\x03*X\n\rEventSeverity\x12\x17\n\x13\x45VENT_SEVERITY_NONE\x10\x00\x12\x0c\n\x08\x43RITICAL\x10\x01\x12\t\n\x05\x45RROR\x10\x02\x12\x0b\n\x07WARNING\x10\x03\x12\x08\n\x04INFO\x10\x04\x32\x8a\x01\n\x05\x45vent\x12\x80\x01\n\x06\x63reate\x12\x31.spaceone.api.alert_manager.v1.EventCreateRequest\x1a\x16.google.protobuf.Empty\"+\x82\xd3\xe4\x93\x02%\" /alert-manager/v1/comment/create:\x01*BDZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/alert-manager/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.alert_manager.v1.event_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/alert-manager/v1'
  _globals['_EVENT'].methods_by_name['create']._loaded_options = None
  _globals['_EVENT'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002%\" /alert-manager/v1/comment/create:\001*'
  _globals['_EVENTSEVERITY']._serialized_start=913
  _globals['_EVENTSEVERITY']._serialized_end=1001
  _globals['_EVENTCREATEREQUEST']._serialized_start=199
  _globals['_EVENTCREATEREQUEST']._serialized_end=298
  _globals['_EVENTINFO']._serialized_start=301
  _globals['_EVENTINFO']._serialized_end=911
  _globals['_EVENTINFO_EVENTTYPE']._serialized_start=843
  _globals['_EVENTINFO_EVENTTYPE']._serialized_end=911
  _globals['_EVENT']._serialized_start=1004
  _globals['_EVENT']._serialized_end=1142
# @@protoc_insertion_point(module_scope)