# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/opsflow/v1/task_category.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n+spaceone/api/opsflow/v1/task_category.proto\x12\x17spaceone.api.opsflow.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"R\n\x0cStatusOption\x12\x11\n\tstatus_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\r\n\x05\x63olor\x18\x03 \x01(\t\x12\x12\n\nis_default\x18\x04 \x01(\x08\"\xba\x01\n\rStatusOptions\x12\x33\n\x04TODO\x18\x01 \x03(\x0b\x32%.spaceone.api.opsflow.v1.StatusOption\x12:\n\x0bIN_PROGRESS\x18\x02 \x03(\x0b\x32%.spaceone.api.opsflow.v1.StatusOption\x12\x38\n\tCOMPLETED\x18\x03 \x03(\x0b\x32%.spaceone.api.opsflow.v1.StatusOption\"\xad\x02\n\tTaskField\x12\x10\n\x08\x66ield_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12\x12\n\nfield_type\x18\x04 \x01(\t\x12(\n\x07options\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12H\n\x0eselection_type\x18\x06 \x01(\x0e\x32\x30.spaceone.api.opsflow.v1.TaskField.SelectionType\x12\x13\n\x0bis_required\x18\x07 \x01(\x08\x12\x12\n\nis_primary\x18\x08 \x01(\x08\":\n\rSelectionType\x12\x12\n\x0eSELECTION_NONE\x10\x00\x12\n\n\x06SINGLE\x10\x01\x12\t\n\x05MULTI\x10\x02\"\x94\x02\n\x19TaskCategoryCreateRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x02 \x01(\t\x12>\n\x0estatus_options\x18\x03 \x01(\x0b\x32&.spaceone.api.opsflow.v1.StatusOptions\x12\x32\n\x06\x66ields\x18\x04 \x03(\x0b\x32\".spaceone.api.opsflow.v1.TaskField\x12%\n\x04vars\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x0b \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x12\n\npackage_id\x18\x15 \x01(\t\"\x84\x02\n\x19TaskCategoryUpdateRequest\x12\x13\n\x0b\x63\x61tegory_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12>\n\x0estatus_options\x18\x04 \x01(\x0b\x32&.spaceone.api.opsflow.v1.StatusOptions\x12\r\n\x05\x66orce\x18\x05 \x01(\x08\x12%\n\x04vars\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x0b \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x12\n\npackage_id\x18\x15 \x01(\t\"y\n\x1fTaskCategoryUpdateFieldsRequest\x12\x13\n\x0b\x63\x61tegory_id\x18\x01 \x01(\t\x12\x32\n\x06\x66ields\x18\x02 \x03(\x0b\x32\".spaceone.api.opsflow.v1.TaskField\x12\r\n\x05\x66orce\x18\x03 \x01(\x08\"*\n\x13TaskCategoryRequest\x12\x13\n\x0b\x63\x61tegory_id\x18\x01 \x01(\t\"\x81\x01\n\x17TaskCategorySearchQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x13\n\x0b\x63\x61tegory_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x17\n\x0finclude_deleted\x18\x04 \x01(\x08\"\xfe\x02\n\x10TaskCategoryInfo\x12\x13\n\x0b\x63\x61tegory_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12>\n\x0estatus_options\x18\x04 \x01(\x0b\x32&.spaceone.api.opsflow.v1.StatusOptions\x12\x32\n\x06\x66ields\x18\x05 \x03(\x0b\x32\".spaceone.api.opsflow.v1.TaskField\x12\r\n\x05state\x18\x06 \x01(\t\x12%\n\x04vars\x18\x07 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x0b \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x12\n\npackage_id\x18\x16 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x12\n\nupdated_at\x18  \x01(\t\x12\x12\n\ndeleted_at\x18! \x01(\t\"e\n\x12TaskCategoriesInfo\x12:\n\x07results\x18\x01 \x03(\x0b\x32).spaceone.api.opsflow.v1.TaskCategoryInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"M\n\x15TaskCategoryStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery2\x92\x08\n\x0cTaskCategory\x12\x94\x01\n\x06\x63reate\x12\x32.spaceone.api.opsflow.v1.TaskCategoryCreateRequest\x1a).spaceone.api.opsflow.v1.TaskCategoryInfo\"+\x82\xd3\xe4\x93\x02%\" /opsflow/v1/task-category/create:\x01*\x12\x94\x01\n\x06update\x12\x32.spaceone.api.opsflow.v1.TaskCategoryUpdateRequest\x1a).spaceone.api.opsflow.v1.TaskCategoryInfo\"+\x82\xd3\xe4\x93\x02%\" /opsflow/v1/task-category/update:\x01*\x12\xa8\x01\n\rupdate_fields\x12\x38.spaceone.api.opsflow.v1.TaskCategoryUpdateFieldsRequest\x1a).spaceone.api.opsflow.v1.TaskCategoryInfo\"2\x82\xd3\xe4\x93\x02,\"\'/opsflow/v1/task-category/update_fields:\x01*\x12\x8e\x01\n\x06\x64\x65lete\x12,.spaceone.api.opsflow.v1.TaskCategoryRequest\x1a).spaceone.api.opsflow.v1.TaskCategoryInfo\"+\x82\xd3\xe4\x93\x02%\" /opsflow/v1/task-category/delete:\x01*\x12\x88\x01\n\x03get\x12,.spaceone.api.opsflow.v1.TaskCategoryRequest\x1a).spaceone.api.opsflow.v1.TaskCategoryInfo\"(\x82\xd3\xe4\x93\x02\"\"\x1d/opsflow/v1/task-category/get:\x01*\x12\x90\x01\n\x04list\x12\x30.spaceone.api.opsflow.v1.TaskCategorySearchQuery\x1a+.spaceone.api.opsflow.v1.TaskCategoriesInfo\")\x82\xd3\xe4\x93\x02#\"\x1e/opsflow/v1/task-category/list:\x01*\x12z\n\x04stat\x12..spaceone.api.opsflow.v1.TaskCategoryStatQuery\x1a\x17.google.protobuf.Struct\")\x82\xd3\xe4\x93\x02#\"\x1e/opsflow/v1/task-category/stat:\x01*B>Z<github.com/cloudforet-io/api/dist/go/spaceone/api/opsflow/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.opsflow.v1.task_category_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z<github.com/cloudforet-io/api/dist/go/spaceone/api/opsflow/v1'
  _globals['_TASKCATEGORY'].methods_by_name['create']._loaded_options = None
  _globals['_TASKCATEGORY'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002%\" /opsflow/v1/task-category/create:\001*'
  _globals['_TASKCATEGORY'].methods_by_name['update']._loaded_options = None
  _globals['_TASKCATEGORY'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002%\" /opsflow/v1/task-category/update:\001*'
  _globals['_TASKCATEGORY'].methods_by_name['update_fields']._loaded_options = None
  _globals['_TASKCATEGORY'].methods_by_name['update_fields']._serialized_options = b'\202\323\344\223\002,\"\'/opsflow/v1/task-category/update_fields:\001*'
  _globals['_TASKCATEGORY'].methods_by_name['delete']._loaded_options = None
  _globals['_TASKCATEGORY'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002%\" /opsflow/v1/task-category/delete:\001*'
  _globals['_TASKCATEGORY'].methods_by_name['get']._loaded_options = None
  _globals['_TASKCATEGORY'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002\"\"\035/opsflow/v1/task-category/get:\001*'
  _globals['_TASKCATEGORY'].methods_by_name['list']._loaded_options = None
  _globals['_TASKCATEGORY'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002#\"\036/opsflow/v1/task-category/list:\001*'
  _globals['_TASKCATEGORY'].methods_by_name['stat']._loaded_options = None
  _globals['_TASKCATEGORY'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002#\"\036/opsflow/v1/task-category/stat:\001*'
  _globals['_STATUSOPTION']._serialized_start=195
  _globals['_STATUSOPTION']._serialized_end=277
  _globals['_STATUSOPTIONS']._serialized_start=280
  _globals['_STATUSOPTIONS']._serialized_end=466
  _globals['_TASKFIELD']._serialized_start=469
  _globals['_TASKFIELD']._serialized_end=770
  _globals['_TASKFIELD_SELECTIONTYPE']._serialized_start=712
  _globals['_TASKFIELD_SELECTIONTYPE']._serialized_end=770
  _globals['_TASKCATEGORYCREATEREQUEST']._serialized_start=773
  _globals['_TASKCATEGORYCREATEREQUEST']._serialized_end=1049
  _globals['_TASKCATEGORYUPDATEREQUEST']._serialized_start=1052
  _globals['_TASKCATEGORYUPDATEREQUEST']._serialized_end=1312
  _globals['_TASKCATEGORYUPDATEFIELDSREQUEST']._serialized_start=1314
  _globals['_TASKCATEGORYUPDATEFIELDSREQUEST']._serialized_end=1435
  _globals['_TASKCATEGORYREQUEST']._serialized_start=1437
  _globals['_TASKCATEGORYREQUEST']._serialized_end=1479
  _globals['_TASKCATEGORYSEARCHQUERY']._serialized_start=1482
  _globals['_TASKCATEGORYSEARCHQUERY']._serialized_end=1611
  _globals['_TASKCATEGORYINFO']._serialized_start=1614
  _globals['_TASKCATEGORYINFO']._serialized_end=1996
  _globals['_TASKCATEGORIESINFO']._serialized_start=1998
  _globals['_TASKCATEGORIESINFO']._serialized_end=2099
  _globals['_TASKCATEGORYSTATQUERY']._serialized_start=2101
  _globals['_TASKCATEGORYSTATQUERY']._serialized_end=2178
  _globals['_TASKCATEGORY']._serialized_start=2181
  _globals['_TASKCATEGORY']._serialized_end=3223
# @@protoc_insertion_point(module_scope)
