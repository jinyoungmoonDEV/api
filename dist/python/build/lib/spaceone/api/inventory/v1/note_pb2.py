# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/inventory/v1/note.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n$spaceone/api/inventory/v1/note.proto\x12\x19spaceone.api.inventory.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"4\n\x11\x43reateNoteRequest\x12\x11\n\trecord_id\x18\x01 \x01(\t\x12\x0c\n\x04note\x18\x02 \x01(\t\"2\n\x11UpdateNoteRequest\x12\x0f\n\x07note_id\x18\x01 \x01(\t\x12\x0c\n\x04note\x18\x02 \x01(\t\"\x1e\n\x0bNoteRequest\x12\x0f\n\x07note_id\x18\x01 \x01(\t\"\xb3\x01\n\tNoteQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x0f\n\x07note_id\x18\x02 \x01(\t\x12\x11\n\trecord_id\x18\x03 \x01(\t\x12\x18\n\x10\x63loud_service_id\x18\x04 \x01(\t\x12\x12\n\ncreated_by\x18\x05 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x15 \x01(\t\x12\x12\n\nproject_id\x18\x16 \x01(\t\"\xbb\x01\n\x08NoteInfo\x12\x0f\n\x07note_id\x18\x01 \x01(\t\x12\x11\n\trecord_id\x18\x02 \x01(\t\x12\x18\n\x10\x63loud_service_id\x18\x03 \x01(\t\x12\x0c\n\x04note\x18\x04 \x01(\t\x12\x12\n\ncreated_by\x18\x0b \x01(\t\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\x12\x12\n\nproject_id\x18\x17 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\"V\n\tNotesInfo\x12\x34\n\x07results\x18\x01 \x03(\x0b\x32#.spaceone.api.inventory.v1.NoteInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"E\n\rNoteStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery2\xdc\x05\n\x04Note\x12\x81\x01\n\x06\x63reate\x12,.spaceone.api.inventory.v1.CreateNoteRequest\x1a#.spaceone.api.inventory.v1.NoteInfo\"$\x82\xd3\xe4\x93\x02\x1e\"\x19/inventory/v1/note/create:\x01*\x12\x81\x01\n\x06update\x12,.spaceone.api.inventory.v1.UpdateNoteRequest\x1a#.spaceone.api.inventory.v1.NoteInfo\"$\x82\xd3\xe4\x93\x02\x1e\"\x19/inventory/v1/note/update:\x01*\x12n\n\x06\x64\x65lete\x12&.spaceone.api.inventory.v1.NoteRequest\x1a\x16.google.protobuf.Empty\"$\x82\xd3\xe4\x93\x02\x1e\"\x19/inventory/v1/note/delete:\x01*\x12u\n\x03get\x12&.spaceone.api.inventory.v1.NoteRequest\x1a#.spaceone.api.inventory.v1.NoteInfo\"!\x82\xd3\xe4\x93\x02\x1b\"\x16/inventory/v1/note/get:\x01*\x12v\n\x04list\x12$.spaceone.api.inventory.v1.NoteQuery\x1a$.spaceone.api.inventory.v1.NotesInfo\"\"\x82\xd3\xe4\x93\x02\x1c\"\x17/inventory/v1/note/list:\x01*\x12m\n\x04stat\x12(.spaceone.api.inventory.v1.NoteStatQuery\x1a\x17.google.protobuf.Struct\"\"\x82\xd3\xe4\x93\x02\x1c\"\x17/inventory/v1/note/stat:\x01*B@Z>github.com/cloudforet-io/api/dist/go/spaceone/api/inventory/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.inventory.v1.note_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z>github.com/cloudforet-io/api/dist/go/spaceone/api/inventory/v1'
  _globals['_NOTE'].methods_by_name['create']._options = None
  _globals['_NOTE'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002\036\"\031/inventory/v1/note/create:\001*'
  _globals['_NOTE'].methods_by_name['update']._options = None
  _globals['_NOTE'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002\036\"\031/inventory/v1/note/update:\001*'
  _globals['_NOTE'].methods_by_name['delete']._options = None
  _globals['_NOTE'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002\036\"\031/inventory/v1/note/delete:\001*'
  _globals['_NOTE'].methods_by_name['get']._options = None
  _globals['_NOTE'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002\033\"\026/inventory/v1/note/get:\001*'
  _globals['_NOTE'].methods_by_name['list']._options = None
  _globals['_NOTE'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002\034\"\027/inventory/v1/note/list:\001*'
  _globals['_NOTE'].methods_by_name['stat']._options = None
  _globals['_NOTE'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002\034\"\027/inventory/v1/note/stat:\001*'
  _globals['_CREATENOTEREQUEST']._serialized_start=190
  _globals['_CREATENOTEREQUEST']._serialized_end=242
  _globals['_UPDATENOTEREQUEST']._serialized_start=244
  _globals['_UPDATENOTEREQUEST']._serialized_end=294
  _globals['_NOTEREQUEST']._serialized_start=296
  _globals['_NOTEREQUEST']._serialized_end=326
  _globals['_NOTEQUERY']._serialized_start=329
  _globals['_NOTEQUERY']._serialized_end=508
  _globals['_NOTEINFO']._serialized_start=511
  _globals['_NOTEINFO']._serialized_end=698
  _globals['_NOTESINFO']._serialized_start=700
  _globals['_NOTESINFO']._serialized_end=786
  _globals['_NOTESTATQUERY']._serialized_start=788
  _globals['_NOTESTATQUERY']._serialized_end=857
  _globals['_NOTE']._serialized_start=860
  _globals['_NOTE']._serialized_end=1592
# @@protoc_insertion_point(module_scope)
