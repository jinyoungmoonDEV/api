# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/identity/v2/workspace_group.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n.spaceone/api/identity/v2/workspace_group.proto\x12\x18spaceone.api.identity.v2\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"R\n\x1b\x43reateWorkspaceGroupRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12%\n\x04tags\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\"n\n\x1bUpdateWorkspaceGroupRequest\x12\x1a\n\x12workspace_group_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12%\n\x04tags\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\"Q\n\x1fWorkspacesWorkspaceGroupRequest\x12\x1a\n\x12workspace_group_id\x18\x01 \x01(\t\x12\x12\n\nworkspaces\x18\x02 \x03(\t\"G\n\x1aUsersWorkspaceGroupRequest\x12\x1a\n\x12workspace_group_id\x18\x01 \x01(\t\x12\r\n\x05users\x18\x02 \x03(\t\"3\n\x15WorkspaceGroupRequest\x12\x1a\n\x12workspace_group_id\x18\x01 \x01(\t\"\x8f\x02\n\x12WorkspaceGroupInfo\x12\x1a\n\x12workspace_group_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x12\n\nworkspaces\x18\x03 \x03(\t\x12\x31\n\x05users\x18\x04 \x03(\x0b\x32\".spaceone.api.identity.v2.UserInfo\x12%\n\x04tags\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x12\n\ncreated_by\x18\x06 \x01(\t\x12\x12\n\nupdated_by\x18\x07 \x01(\t\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x12\n\nupdated_at\x18  \x01(\t\"\x99\x01\n\x19WorkspaceGroupSearchQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x1a\n\x12workspace_group_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x12\n\ncreated_by\x18\x04 \x01(\t\x12\x12\n\nupdated_by\x18\x05 \x01(\t\"?\n\x08UserInfo\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x0f\n\x07role_id\x18\x02 \x01(\t\x12\x11\n\trole_type\x18\x03 \x01(\t\"i\n\x13WorkspaceGroupsInfo\x12=\n\x07results\x18\x01 \x03(\x0b\x32,.spaceone.api.identity.v2.WorkspaceGroupInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"O\n\x17WorkspaceGroupStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery2\xc4\x0c\n\x0eWorkspaceGroup\x12\x9d\x01\n\x06\x63reate\x12\x35.spaceone.api.identity.v2.CreateWorkspaceGroupRequest\x1a,.spaceone.api.identity.v2.WorkspaceGroupInfo\".\x82\xd3\xe4\x93\x02(\"#/identity/v2/workspace-group/create:\x01*\x12\x9d\x01\n\x06update\x12\x35.spaceone.api.identity.v2.UpdateWorkspaceGroupRequest\x1a,.spaceone.api.identity.v2.WorkspaceGroupInfo\".\x82\xd3\xe4\x93\x02(\"#/identity/v2/workspace-group/update:\x01*\x12\x81\x01\n\x06\x64\x65lete\x12/.spaceone.api.identity.v2.WorkspaceGroupRequest\x1a\x16.google.protobuf.Empty\".\x82\xd3\xe4\x93\x02(\"#/identity/v2/workspace-group/delete:\x01*\x12\xb1\x01\n\x0e\x61\x64\x64_workspaces\x12\x39.spaceone.api.identity.v2.WorkspacesWorkspaceGroupRequest\x1a,.spaceone.api.identity.v2.WorkspaceGroupInfo\"6\x82\xd3\xe4\x93\x02\x30\"+/identity/v2/workspace-group/add-workspaces:\x01*\x12\xb7\x01\n\x11remove_workspaces\x12\x39.spaceone.api.identity.v2.WorkspacesWorkspaceGroupRequest\x1a,.spaceone.api.identity.v2.WorkspaceGroupInfo\"9\x82\xd3\xe4\x93\x02\x33\"./identity/v2/workspace-group/remove-workspaces:\x01*\x12\xa2\x01\n\tadd_users\x12\x34.spaceone.api.identity.v2.UsersWorkspaceGroupRequest\x1a,.spaceone.api.identity.v2.WorkspaceGroupInfo\"1\x82\xd3\xe4\x93\x02+\"&/identity/v2/workspace-group/add-users:\x01*\x12\xa8\x01\n\x0cremove_users\x12\x34.spaceone.api.identity.v2.UsersWorkspaceGroupRequest\x1a,.spaceone.api.identity.v2.WorkspaceGroupInfo\"4\x82\xd3\xe4\x93\x02.\")/identity/v2/workspace-group/remove-users:\x01*\x12\x91\x01\n\x03get\x12/.spaceone.api.identity.v2.WorkspaceGroupRequest\x1a,.spaceone.api.identity.v2.WorkspaceGroupInfo\"+\x82\xd3\xe4\x93\x02%\" /identity/v2/workspace-group/get:\x01*\x12\x98\x01\n\x04list\x12\x33.spaceone.api.identity.v2.WorkspaceGroupSearchQuery\x1a-.spaceone.api.identity.v2.WorkspaceGroupsInfo\",\x82\xd3\xe4\x93\x02&\"!/identity/v2/workspace-group/list:\x01*\x12\x80\x01\n\x04stat\x12\x31.spaceone.api.identity.v2.WorkspaceGroupStatQuery\x1a\x17.google.protobuf.Struct\",\x82\xd3\xe4\x93\x02&\"!/identity/v2/workspace-group/stat:\x01*B?Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v2b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.identity.v2.workspace_group_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v2'
  _globals['_WORKSPACEGROUP'].methods_by_name['create']._loaded_options = None
  _globals['_WORKSPACEGROUP'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002(\"#/identity/v2/workspace-group/create:\001*'
  _globals['_WORKSPACEGROUP'].methods_by_name['update']._loaded_options = None
  _globals['_WORKSPACEGROUP'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002(\"#/identity/v2/workspace-group/update:\001*'
  _globals['_WORKSPACEGROUP'].methods_by_name['delete']._loaded_options = None
  _globals['_WORKSPACEGROUP'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002(\"#/identity/v2/workspace-group/delete:\001*'
  _globals['_WORKSPACEGROUP'].methods_by_name['add_workspaces']._loaded_options = None
  _globals['_WORKSPACEGROUP'].methods_by_name['add_workspaces']._serialized_options = b'\202\323\344\223\0020\"+/identity/v2/workspace-group/add-workspaces:\001*'
  _globals['_WORKSPACEGROUP'].methods_by_name['remove_workspaces']._loaded_options = None
  _globals['_WORKSPACEGROUP'].methods_by_name['remove_workspaces']._serialized_options = b'\202\323\344\223\0023\"./identity/v2/workspace-group/remove-workspaces:\001*'
  _globals['_WORKSPACEGROUP'].methods_by_name['add_users']._loaded_options = None
  _globals['_WORKSPACEGROUP'].methods_by_name['add_users']._serialized_options = b'\202\323\344\223\002+\"&/identity/v2/workspace-group/add-users:\001*'
  _globals['_WORKSPACEGROUP'].methods_by_name['remove_users']._loaded_options = None
  _globals['_WORKSPACEGROUP'].methods_by_name['remove_users']._serialized_options = b'\202\323\344\223\002.\")/identity/v2/workspace-group/remove-users:\001*'
  _globals['_WORKSPACEGROUP'].methods_by_name['get']._loaded_options = None
  _globals['_WORKSPACEGROUP'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002%\" /identity/v2/workspace-group/get:\001*'
  _globals['_WORKSPACEGROUP'].methods_by_name['list']._loaded_options = None
  _globals['_WORKSPACEGROUP'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002&\"!/identity/v2/workspace-group/list:\001*'
  _globals['_WORKSPACEGROUP'].methods_by_name['stat']._loaded_options = None
  _globals['_WORKSPACEGROUP'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002&\"!/identity/v2/workspace-group/stat:\001*'
  _globals['_CREATEWORKSPACEGROUPREQUEST']._serialized_start=199
  _globals['_CREATEWORKSPACEGROUPREQUEST']._serialized_end=281
  _globals['_UPDATEWORKSPACEGROUPREQUEST']._serialized_start=283
  _globals['_UPDATEWORKSPACEGROUPREQUEST']._serialized_end=393
  _globals['_WORKSPACESWORKSPACEGROUPREQUEST']._serialized_start=395
  _globals['_WORKSPACESWORKSPACEGROUPREQUEST']._serialized_end=476
  _globals['_USERSWORKSPACEGROUPREQUEST']._serialized_start=478
  _globals['_USERSWORKSPACEGROUPREQUEST']._serialized_end=549
  _globals['_WORKSPACEGROUPREQUEST']._serialized_start=551
  _globals['_WORKSPACEGROUPREQUEST']._serialized_end=602
  _globals['_WORKSPACEGROUPINFO']._serialized_start=605
  _globals['_WORKSPACEGROUPINFO']._serialized_end=876
  _globals['_WORKSPACEGROUPSEARCHQUERY']._serialized_start=879
  _globals['_WORKSPACEGROUPSEARCHQUERY']._serialized_end=1032
  _globals['_USERINFO']._serialized_start=1034
  _globals['_USERINFO']._serialized_end=1097
  _globals['_WORKSPACEGROUPSINFO']._serialized_start=1099
  _globals['_WORKSPACEGROUPSINFO']._serialized_end=1204
  _globals['_WORKSPACEGROUPSTATQUERY']._serialized_start=1206
  _globals['_WORKSPACEGROUPSTATQUERY']._serialized_end=1285
  _globals['_WORKSPACEGROUP']._serialized_start=1288
  _globals['_WORKSPACEGROUP']._serialized_end=2892
# @@protoc_insertion_point(module_scope)
