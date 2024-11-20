# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/opsflow/v1/task_type.proto
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
from spaceone.api.opsflow.v1 import task_category_pb2 as spaceone_dot_api_dot_opsflow_dot_v1_dot_task__category__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\'spaceone/api/opsflow/v1/task_type.proto\x12\x17spaceone.api.opsflow.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\x1a+spaceone/api/opsflow/v1/task_category.proto\"\xaa\x01\n\x15TaskTypeCreateRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x02 \x01(\t\x12\x32\n\x06\x66ields\x18\x03 \x03(\x0b\x32\".spaceone.api.opsflow.v1.TaskField\x12%\n\x04tags\x18\x0b \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x13\n\x0b\x63\x61tegory_id\x18\x15 \x01(\t\"\xc0\x01\n\x15TaskTypeUpdateRequest\x12\x14\n\x0ctask_type_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12\x32\n\x06\x66ields\x18\x04 \x03(\x0b\x32\".spaceone.api.opsflow.v1.TaskField\x12%\n\x04tags\x18\x0b \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x13\n\x0b\x63\x61tegory_id\x18\x15 \x01(\t\"\'\n\x0fTaskTypeRequest\x12\x14\n\x0ctask_type_id\x18\x01 \x01(\t\"e\n\x13TaskTypeSearchQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x14\n\x0ctask_type_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\"\xf2\x01\n\x0cTaskTypeInfo\x12\x14\n\x0ctask_type_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12\x32\n\x06\x66ields\x18\x04 \x03(\x0b\x32\".spaceone.api.opsflow.v1.TaskField\x12%\n\x04tags\x18\x0b \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x13\n\x0b\x63\x61tegory_id\x18\x16 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x12\n\nupdated_at\x18  \x01(\t\"\\\n\rTaskTypesInfo\x12\x36\n\x07results\x18\x01 \x03(\x0b\x32%.spaceone.api.opsflow.v1.TaskTypeInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"I\n\x11TaskTypeStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery2\x8d\x06\n\x08TaskType\x12\x88\x01\n\x06\x63reate\x12..spaceone.api.opsflow.v1.TaskTypeCreateRequest\x1a%.spaceone.api.opsflow.v1.TaskTypeInfo\"\'\x82\xd3\xe4\x93\x02!\"\x1c/opsflow/v1/task-type/create:\x01*\x12\x88\x01\n\x06update\x12..spaceone.api.opsflow.v1.TaskTypeUpdateRequest\x1a%.spaceone.api.opsflow.v1.TaskTypeInfo\"\'\x82\xd3\xe4\x93\x02!\"\x1c/opsflow/v1/task-type/update:\x01*\x12s\n\x06\x64\x65lete\x12(.spaceone.api.opsflow.v1.TaskTypeRequest\x1a\x16.google.protobuf.Empty\"\'\x82\xd3\xe4\x93\x02!\"\x1c/opsflow/v1/task-type/delete:\x01*\x12|\n\x03get\x12(.spaceone.api.opsflow.v1.TaskTypeRequest\x1a%.spaceone.api.opsflow.v1.TaskTypeInfo\"$\x82\xd3\xe4\x93\x02\x1e\"\x19/opsflow/v1/task-type/get:\x01*\x12\x83\x01\n\x04list\x12,.spaceone.api.opsflow.v1.TaskTypeSearchQuery\x1a&.spaceone.api.opsflow.v1.TaskTypesInfo\"%\x82\xd3\xe4\x93\x02\x1f\"\x1a/opsflow/v1/task-type/list:\x01*\x12r\n\x04stat\x12*.spaceone.api.opsflow.v1.TaskTypeStatQuery\x1a\x17.google.protobuf.Struct\"%\x82\xd3\xe4\x93\x02\x1f\"\x1a/opsflow/v1/task-type/stat:\x01*B>Z<github.com/cloudforet-io/api/dist/go/spaceone/api/opsflow/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.opsflow.v1.task_type_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z<github.com/cloudforet-io/api/dist/go/spaceone/api/opsflow/v1'
  _globals['_TASKTYPE'].methods_by_name['create']._loaded_options = None
  _globals['_TASKTYPE'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002!\"\034/opsflow/v1/task-type/create:\001*'
  _globals['_TASKTYPE'].methods_by_name['update']._loaded_options = None
  _globals['_TASKTYPE'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002!\"\034/opsflow/v1/task-type/update:\001*'
  _globals['_TASKTYPE'].methods_by_name['delete']._loaded_options = None
  _globals['_TASKTYPE'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002!\"\034/opsflow/v1/task-type/delete:\001*'
  _globals['_TASKTYPE'].methods_by_name['get']._loaded_options = None
  _globals['_TASKTYPE'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002\036\"\031/opsflow/v1/task-type/get:\001*'
  _globals['_TASKTYPE'].methods_by_name['list']._loaded_options = None
  _globals['_TASKTYPE'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002\037\"\032/opsflow/v1/task-type/list:\001*'
  _globals['_TASKTYPE'].methods_by_name['stat']._loaded_options = None
  _globals['_TASKTYPE'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002\037\"\032/opsflow/v1/task-type/stat:\001*'
  _globals['_TASKTYPECREATEREQUEST']._serialized_start=237
  _globals['_TASKTYPECREATEREQUEST']._serialized_end=407
  _globals['_TASKTYPEUPDATEREQUEST']._serialized_start=410
  _globals['_TASKTYPEUPDATEREQUEST']._serialized_end=602
  _globals['_TASKTYPEREQUEST']._serialized_start=604
  _globals['_TASKTYPEREQUEST']._serialized_end=643
  _globals['_TASKTYPESEARCHQUERY']._serialized_start=645
  _globals['_TASKTYPESEARCHQUERY']._serialized_end=746
  _globals['_TASKTYPEINFO']._serialized_start=749
  _globals['_TASKTYPEINFO']._serialized_end=991
  _globals['_TASKTYPESINFO']._serialized_start=993
  _globals['_TASKTYPESINFO']._serialized_end=1085
  _globals['_TASKTYPESTATQUERY']._serialized_start=1087
  _globals['_TASKTYPESTATQUERY']._serialized_end=1160
  _globals['_TASKTYPE']._serialized_start=1163
  _globals['_TASKTYPE']._serialized_end=1944
# @@protoc_insertion_point(module_scope)