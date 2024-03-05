# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/dashboard/v1/private_dashboard.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n1spaceone/api/dashboard/v1/private_dashboard.proto\x12\x19spaceone.api.dashboard.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"\xe6\x02\n\x1dPrivateCreateDashboardRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12+\n\x07layouts\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12*\n\tvariables\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12)\n\x08settings\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x31\n\x10variables_schema\x18\x07 \x01(\x0b\x32\x17.google.protobuf.Struct\x12-\n\x0c\x64isplay_info\x18\x08 \x01(\x0b\x32\x17.google.protobuf.Struct\x12*\n\x06labels\x18\t \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12%\n\x04tags\x18\n \x01(\x0b\x32\x17.google.protobuf.Struct\"\x84\x03\n\x1dPrivateUpdateDashboardRequest\x12\x1c\n\x14private_dashboard_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12+\n\x07layouts\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12*\n\tvariables\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12)\n\x08settings\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x31\n\x10variables_schema\x18\x07 \x01(\x0b\x32\x17.google.protobuf.Struct\x12-\n\x0c\x64isplay_info\x18\x08 \x01(\x0b\x32\x17.google.protobuf.Struct\x12*\n\x06labels\x18\t \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12%\n\x04tags\x18\n \x01(\x0b\x32\x17.google.protobuf.Struct\"7\n\x17PrivateDashboardRequest\x12\x1c\n\x14private_dashboard_id\x18\x01 \x01(\t\"O\n\x1ePrivateDashboardVersionRequest\x12\x1c\n\x14private_dashboard_id\x18\x01 \x01(\t\x12\x0f\n\x07version\x18\x04 \x01(\x05\"\x7f\n\"PrivateDashboardVersionSearchQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x1c\n\x14private_dashboard_id\x18\x02 \x01(\t\x12\x0f\n\x07version\x18\x03 \x01(\x05\"\x99\x01\n\x15PrivateDashboardQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x1c\n\x14private_dashboard_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x15 \x01(\t\x12\x12\n\nproject_id\x18\x16 \x01(\t\"\xee\x03\n\x14PrivateDashboardInfo\x12\x1c\n\x14private_dashboard_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x0f\n\x07version\x18\x04 \x01(\x05\x12+\n\x07layouts\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12*\n\tvariables\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\x12)\n\x08settings\x18\x07 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x31\n\x10variables_schema\x18\x08 \x01(\x0b\x32\x17.google.protobuf.Struct\x12-\n\x0c\x64isplay_info\x18\t \x01(\x0b\x32\x17.google.protobuf.Struct\x12*\n\x06labels\x18\n \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12%\n\x04tags\x18\x0b \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\x12\x0f\n\x07user_id\x18\x17 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x12\n\nupdated_at\x18  \x01(\t\"n\n\x15PrivateDashboardsInfo\x12@\n\x07results\x18\x01 \x03(\x0b\x32/.spaceone.api.dashboard.v1.PrivateDashboardInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"Q\n\x19PrivateDashboardStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery\"\xba\x02\n\x1bPrivateDashboardVersionInfo\x12\x1c\n\x14private_dashboard_id\x18\x01 \x01(\t\x12\x0f\n\x07version\x18\x02 \x01(\x05\x12\x0e\n\x06latest\x18\x03 \x01(\x08\x12+\n\x07layouts\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12*\n\tvariables\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12)\n\x08settings\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x31\n\x10variables_schema\x18\x07 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\"|\n\x1cPrivateDashboardVersionsInfo\x12G\n\x07results\x18\x01 \x03(\x0b\x32\x36.spaceone.api.dashboard.v1.PrivateDashboardVersionInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\x32\x8c\r\n\x10PrivateDashboard\x12\xa6\x01\n\x06\x63reate\x12\x38.spaceone.api.dashboard.v1.PrivateCreateDashboardRequest\x1a/.spaceone.api.dashboard.v1.PrivateDashboardInfo\"1\x82\xd3\xe4\x93\x02+\"&/dashboard/v1/private-dashboard/create:\x01*\x12\xa6\x01\n\x06update\x12\x38.spaceone.api.dashboard.v1.PrivateUpdateDashboardRequest\x1a/.spaceone.api.dashboard.v1.PrivateDashboardInfo\"1\x82\xd3\xe4\x93\x02+\"&/dashboard/v1/private-dashboard/update:\x01*\x12\x87\x01\n\x06\x64\x65lete\x12\x32.spaceone.api.dashboard.v1.PrivateDashboardRequest\x1a\x16.google.protobuf.Empty\"1\x82\xd3\xe4\x93\x02+\"&/dashboard/v1/private-dashboard/delete:\x01*\x12\x9a\x01\n\x03get\x12\x32.spaceone.api.dashboard.v1.PrivateDashboardRequest\x1a/.spaceone.api.dashboard.v1.PrivateDashboardInfo\".\x82\xd3\xe4\x93\x02(\"#/dashboard/v1/private-dashboard/get:\x01*\x12\x9e\x01\n\x0e\x64\x65lete_version\x12\x39.spaceone.api.dashboard.v1.PrivateDashboardVersionRequest\x1a\x16.google.protobuf.Empty\"9\x82\xd3\xe4\x93\x02\x33\"./dashboard/v1/private-dashboard/delete-version:\x01*\x12\xb7\x01\n\x0erevert_version\x12\x39.spaceone.api.dashboard.v1.PrivateDashboardVersionRequest\x1a/.spaceone.api.dashboard.v1.PrivateDashboardInfo\"9\x82\xd3\xe4\x93\x02\x33\"./dashboard/v1/private-dashboard/revert-version:\x01*\x12\xb8\x01\n\x0bget_version\x12\x39.spaceone.api.dashboard.v1.PrivateDashboardVersionRequest\x1a\x36.spaceone.api.dashboard.v1.PrivateDashboardVersionInfo\"6\x82\xd3\xe4\x93\x02\x30\"+/dashboard/v1/private-dashboard/get-version:\x01*\x12\xc1\x01\n\rlist_versions\x12=.spaceone.api.dashboard.v1.PrivateDashboardVersionSearchQuery\x1a\x37.spaceone.api.dashboard.v1.PrivateDashboardVersionsInfo\"8\x82\xd3\xe4\x93\x02\x32\"-/dashboard/v1/private-dashboard/list-versions:\x01*\x12\x9b\x01\n\x04list\x12\x30.spaceone.api.dashboard.v1.PrivateDashboardQuery\x1a\x30.spaceone.api.dashboard.v1.PrivateDashboardsInfo\"/\x82\xd3\xe4\x93\x02)\"$/dashboard/v1/private-dashboard/list:\x01*\x12\x86\x01\n\x04stat\x12\x34.spaceone.api.dashboard.v1.PrivateDashboardStatQuery\x1a\x17.google.protobuf.Struct\"/\x82\xd3\xe4\x93\x02)\"$/dashboard/v1/private-dashboard/stat:\x01*B@Z>github.com/cloudforet-io/api/dist/go/spaceone/api/dashboard/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.dashboard.v1.private_dashboard_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z>github.com/cloudforet-io/api/dist/go/spaceone/api/dashboard/v1'
  _globals['_PRIVATEDASHBOARD'].methods_by_name['create']._options = None
  _globals['_PRIVATEDASHBOARD'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002+\"&/dashboard/v1/private-dashboard/create:\001*'
  _globals['_PRIVATEDASHBOARD'].methods_by_name['update']._options = None
  _globals['_PRIVATEDASHBOARD'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002+\"&/dashboard/v1/private-dashboard/update:\001*'
  _globals['_PRIVATEDASHBOARD'].methods_by_name['delete']._options = None
  _globals['_PRIVATEDASHBOARD'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002+\"&/dashboard/v1/private-dashboard/delete:\001*'
  _globals['_PRIVATEDASHBOARD'].methods_by_name['get']._options = None
  _globals['_PRIVATEDASHBOARD'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002(\"#/dashboard/v1/private-dashboard/get:\001*'
  _globals['_PRIVATEDASHBOARD'].methods_by_name['delete_version']._options = None
  _globals['_PRIVATEDASHBOARD'].methods_by_name['delete_version']._serialized_options = b'\202\323\344\223\0023\"./dashboard/v1/private-dashboard/delete-version:\001*'
  _globals['_PRIVATEDASHBOARD'].methods_by_name['revert_version']._options = None
  _globals['_PRIVATEDASHBOARD'].methods_by_name['revert_version']._serialized_options = b'\202\323\344\223\0023\"./dashboard/v1/private-dashboard/revert-version:\001*'
  _globals['_PRIVATEDASHBOARD'].methods_by_name['get_version']._options = None
  _globals['_PRIVATEDASHBOARD'].methods_by_name['get_version']._serialized_options = b'\202\323\344\223\0020\"+/dashboard/v1/private-dashboard/get-version:\001*'
  _globals['_PRIVATEDASHBOARD'].methods_by_name['list_versions']._options = None
  _globals['_PRIVATEDASHBOARD'].methods_by_name['list_versions']._serialized_options = b'\202\323\344\223\0022\"-/dashboard/v1/private-dashboard/list-versions:\001*'
  _globals['_PRIVATEDASHBOARD'].methods_by_name['list']._options = None
  _globals['_PRIVATEDASHBOARD'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002)\"$/dashboard/v1/private-dashboard/list:\001*'
  _globals['_PRIVATEDASHBOARD'].methods_by_name['stat']._options = None
  _globals['_PRIVATEDASHBOARD'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002)\"$/dashboard/v1/private-dashboard/stat:\001*'
  _globals['_PRIVATECREATEDASHBOARDREQUEST']._serialized_start=204
  _globals['_PRIVATECREATEDASHBOARDREQUEST']._serialized_end=562
  _globals['_PRIVATEUPDATEDASHBOARDREQUEST']._serialized_start=565
  _globals['_PRIVATEUPDATEDASHBOARDREQUEST']._serialized_end=953
  _globals['_PRIVATEDASHBOARDREQUEST']._serialized_start=955
  _globals['_PRIVATEDASHBOARDREQUEST']._serialized_end=1010
  _globals['_PRIVATEDASHBOARDVERSIONREQUEST']._serialized_start=1012
  _globals['_PRIVATEDASHBOARDVERSIONREQUEST']._serialized_end=1091
  _globals['_PRIVATEDASHBOARDVERSIONSEARCHQUERY']._serialized_start=1093
  _globals['_PRIVATEDASHBOARDVERSIONSEARCHQUERY']._serialized_end=1220
  _globals['_PRIVATEDASHBOARDQUERY']._serialized_start=1223
  _globals['_PRIVATEDASHBOARDQUERY']._serialized_end=1376
  _globals['_PRIVATEDASHBOARDINFO']._serialized_start=1379
  _globals['_PRIVATEDASHBOARDINFO']._serialized_end=1873
  _globals['_PRIVATEDASHBOARDSINFO']._serialized_start=1875
  _globals['_PRIVATEDASHBOARDSINFO']._serialized_end=1985
  _globals['_PRIVATEDASHBOARDSTATQUERY']._serialized_start=1987
  _globals['_PRIVATEDASHBOARDSTATQUERY']._serialized_end=2068
  _globals['_PRIVATEDASHBOARDVERSIONINFO']._serialized_start=2071
  _globals['_PRIVATEDASHBOARDVERSIONINFO']._serialized_end=2385
  _globals['_PRIVATEDASHBOARDVERSIONSINFO']._serialized_start=2387
  _globals['_PRIVATEDASHBOARDVERSIONSINFO']._serialized_end=2511
  _globals['_PRIVATEDASHBOARD']._serialized_start=2514
  _globals['_PRIVATEDASHBOARD']._serialized_end=4190
# @@protoc_insertion_point(module_scope)
