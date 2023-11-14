# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/identity/v2/workspace.proto
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
from spaceone.api.core.v1 import query_pb2 as spaceone_dot_api_dot_core_dot_v1_dot_query__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n(spaceone/api/identity/v2/workspace.proto\x12\x18spaceone.api.identity.v2\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v1/query.proto\"`\n\x16\x43reateWorkSpaceRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12%\n\x04tags\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\"v\n\x16UpdateWorkSpaceRequest\x12\x14\n\x0cworkspace_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12%\n\x04tags\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\";\n\x10WorkspaceRequest\x12\x14\n\x0cworkspace_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x15 \x01(\t\"y\n\x14WorkspaceSearchQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v1.Query\x12\x14\n\x0cworkspace_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x11\n\tdomain_id\x18\x15 \x01(\t\"\xed\x01\n\rWorkspaceInfo\x12\x14\n\x0cworkspace_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12<\n\x05state\x18\x03 \x01(\x0e\x32-.spaceone.api.identity.v2.WorkspaceInfo.State\x12%\n\x04tags\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\",\n\x05State\x12\x08\n\x04NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\"_\n\x0eWorkspacesInfo\x12\x38\n\x07results\x18\x01 \x03(\x0b\x32\'.spaceone.api.identity.v2.WorkspaceInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"]\n\x12WorkspaceStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v1.StatisticsQuery\x12\x11\n\tdomain_id\x18\x02 \x01(\t2\xbf\x08\n\tWorkspace\x12\x8d\x01\n\x06\x63reate\x12\x30.spaceone.api.identity.v2.CreateWorkSpaceRequest\x1a\'.spaceone.api.identity.v2.WorkspaceInfo\"(\x82\xd3\xe4\x93\x02\"\"\x1d/identity/v2/workspace/create:\x01*\x12\x8d\x01\n\x06update\x12\x30.spaceone.api.identity.v2.UpdateWorkSpaceRequest\x1a\'.spaceone.api.identity.v2.WorkspaceInfo\"(\x82\xd3\xe4\x93\x02\"\"\x1d/identity/v2/workspace/update:\x01*\x12v\n\x06\x64\x65lete\x12*.spaceone.api.identity.v2.WorkspaceRequest\x1a\x16.google.protobuf.Empty\"(\x82\xd3\xe4\x93\x02\"\"\x1d/identity/v2/workspace/delete:\x01*\x12\x87\x01\n\x06\x65nable\x12*.spaceone.api.identity.v2.WorkspaceRequest\x1a\'.spaceone.api.identity.v2.WorkspaceInfo\"(\x82\xd3\xe4\x93\x02\"\"\x1d/identity/v2/workspace/enable:\x01*\x12\x89\x01\n\x07\x64isable\x12*.spaceone.api.identity.v2.WorkspaceRequest\x1a\'.spaceone.api.identity.v2.WorkspaceInfo\")\x82\xd3\xe4\x93\x02#\"\x1e/identity/v2/workspace/disable:\x01*\x12\x81\x01\n\x03get\x12*.spaceone.api.identity.v2.WorkspaceRequest\x1a\'.spaceone.api.identity.v2.WorkspaceInfo\"%\x82\xd3\xe4\x93\x02\x1f\"\x1a/identity/v2/workspace/get:\x01*\x12\x88\x01\n\x04list\x12..spaceone.api.identity.v2.WorkspaceSearchQuery\x1a(.spaceone.api.identity.v2.WorkspacesInfo\"&\x82\xd3\xe4\x93\x02 \"\x1b/identity/v2/workspace/list:\x01*\x12u\n\x04stat\x12,.spaceone.api.identity.v2.WorkspaceStatQuery\x1a\x17.google.protobuf.Struct\"&\x82\xd3\xe4\x93\x02 \"\x1b/identity/v2/workspace/stat:\x01*B?Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v2b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.identity.v2.workspace_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v2'
  _WORKSPACE.methods_by_name['create']._options = None
  _WORKSPACE.methods_by_name['create']._serialized_options = b'\202\323\344\223\002\"\"\035/identity/v2/workspace/create:\001*'
  _WORKSPACE.methods_by_name['update']._options = None
  _WORKSPACE.methods_by_name['update']._serialized_options = b'\202\323\344\223\002\"\"\035/identity/v2/workspace/update:\001*'
  _WORKSPACE.methods_by_name['delete']._options = None
  _WORKSPACE.methods_by_name['delete']._serialized_options = b'\202\323\344\223\002\"\"\035/identity/v2/workspace/delete:\001*'
  _WORKSPACE.methods_by_name['enable']._options = None
  _WORKSPACE.methods_by_name['enable']._serialized_options = b'\202\323\344\223\002\"\"\035/identity/v2/workspace/enable:\001*'
  _WORKSPACE.methods_by_name['disable']._options = None
  _WORKSPACE.methods_by_name['disable']._serialized_options = b'\202\323\344\223\002#\"\036/identity/v2/workspace/disable:\001*'
  _WORKSPACE.methods_by_name['get']._options = None
  _WORKSPACE.methods_by_name['get']._serialized_options = b'\202\323\344\223\002\037\"\032/identity/v2/workspace/get:\001*'
  _WORKSPACE.methods_by_name['list']._options = None
  _WORKSPACE.methods_by_name['list']._serialized_options = b'\202\323\344\223\002 \"\033/identity/v2/workspace/list:\001*'
  _WORKSPACE.methods_by_name['stat']._options = None
  _WORKSPACE.methods_by_name['stat']._serialized_options = b'\202\323\344\223\002 \"\033/identity/v2/workspace/stat:\001*'
  _globals['_CREATEWORKSPACEREQUEST']._serialized_start=193
  _globals['_CREATEWORKSPACEREQUEST']._serialized_end=289
  _globals['_UPDATEWORKSPACEREQUEST']._serialized_start=291
  _globals['_UPDATEWORKSPACEREQUEST']._serialized_end=409
  _globals['_WORKSPACEREQUEST']._serialized_start=411
  _globals['_WORKSPACEREQUEST']._serialized_end=470
  _globals['_WORKSPACESEARCHQUERY']._serialized_start=472
  _globals['_WORKSPACESEARCHQUERY']._serialized_end=593
  _globals['_WORKSPACEINFO']._serialized_start=596
  _globals['_WORKSPACEINFO']._serialized_end=833
  _globals['_WORKSPACEINFO_STATE']._serialized_start=789
  _globals['_WORKSPACEINFO_STATE']._serialized_end=833
  _globals['_WORKSPACESINFO']._serialized_start=835
  _globals['_WORKSPACESINFO']._serialized_end=930
  _globals['_WORKSPACESTATQUERY']._serialized_start=932
  _globals['_WORKSPACESTATQUERY']._serialized_end=1025
  _globals['_WORKSPACE']._serialized_start=1028
  _globals['_WORKSPACE']._serialized_end=2115
# @@protoc_insertion_point(module_scope)