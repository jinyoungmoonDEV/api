# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/identity/v2/package.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n&spaceone/api/identity/v2/package.proto\x12\x18spaceone.api.identity.v2\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"`\n\x14\x43reatePackageRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x02 \x01(\t\x12%\n\x04tags\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\"t\n\x14UpdatePackageRequest\x12\x12\n\npackage_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12%\n\x04tags\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\"$\n\x0ePackageRequest\x12\x12\n\npackage_id\x18\x01 \x01(\t\">\n\x19\x43hangePackageOrderRequest\x12\x12\n\npackage_id\x18\x01 \x01(\t\x12\r\n\x05order\x18\x02 \x01(\x05\"\xdf\x01\n\x0bPackageInfo\x12\x12\n\npackage_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12\r\n\x05order\x18\x04 \x01(\x05\x12\x12\n\nis_default\x18\x05 \x01(\x08\x12%\n\x04tags\x18\x0b \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x12\n\nupdated_at\x18  \x01(\t\"b\n\x12PackageSearchQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x12\n\npackage_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\"[\n\x0cPackagesInfo\x12\x36\n\x07results\x18\x01 \x03(\x0b\x32%.spaceone.api.identity.v2.PackageInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"H\n\x10PackageStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery2\xaf\x08\n\x07Package\x12\x87\x01\n\x06\x63reate\x12..spaceone.api.identity.v2.CreatePackageRequest\x1a%.spaceone.api.identity.v2.PackageInfo\"&\x82\xd3\xe4\x93\x02 \"\x1b/identity/v2/package/create:\x01*\x12\x87\x01\n\x06update\x12..spaceone.api.identity.v2.UpdatePackageRequest\x1a%.spaceone.api.identity.v2.PackageInfo\"&\x82\xd3\xe4\x93\x02 \"\x1b/identity/v2/package/update:\x01*\x12r\n\x06\x64\x65lete\x12(.spaceone.api.identity.v2.PackageRequest\x1a\x16.google.protobuf.Empty\"&\x82\xd3\xe4\x93\x02 \"\x1b/identity/v2/package/delete:\x01*\x12\x8b\x01\n\x0bset_default\x12(.spaceone.api.identity.v2.PackageRequest\x1a%.spaceone.api.identity.v2.PackageInfo\"+\x82\xd3\xe4\x93\x02%\" /identity/v2/package/set-default:\x01*\x12\x98\x01\n\x0c\x63hange_order\x12\x33.spaceone.api.identity.v2.ChangePackageOrderRequest\x1a%.spaceone.api.identity.v2.PackageInfo\",\x82\xd3\xe4\x93\x02&\"!/identity/v2/package/change-order:\x01*\x12{\n\x03get\x12(.spaceone.api.identity.v2.PackageRequest\x1a%.spaceone.api.identity.v2.PackageInfo\"#\x82\xd3\xe4\x93\x02\x1d\"\x18/identity/v2/package/get:\x01*\x12\x82\x01\n\x04list\x12,.spaceone.api.identity.v2.PackageSearchQuery\x1a&.spaceone.api.identity.v2.PackagesInfo\"$\x82\xd3\xe4\x93\x02\x1e\"\x19/identity/v2/package/list:\x01*\x12q\n\x04stat\x12*.spaceone.api.identity.v2.PackageStatQuery\x1a\x17.google.protobuf.Struct\"$\x82\xd3\xe4\x93\x02\x1e\"\x19/identity/v1/package/stat:\x01*B?Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v2b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.identity.v2.package_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v2'
  _globals['_PACKAGE'].methods_by_name['create']._loaded_options = None
  _globals['_PACKAGE'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002 \"\033/identity/v2/package/create:\001*'
  _globals['_PACKAGE'].methods_by_name['update']._loaded_options = None
  _globals['_PACKAGE'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002 \"\033/identity/v2/package/update:\001*'
  _globals['_PACKAGE'].methods_by_name['delete']._loaded_options = None
  _globals['_PACKAGE'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002 \"\033/identity/v2/package/delete:\001*'
  _globals['_PACKAGE'].methods_by_name['set_default']._loaded_options = None
  _globals['_PACKAGE'].methods_by_name['set_default']._serialized_options = b'\202\323\344\223\002%\" /identity/v2/package/set-default:\001*'
  _globals['_PACKAGE'].methods_by_name['change_order']._loaded_options = None
  _globals['_PACKAGE'].methods_by_name['change_order']._serialized_options = b'\202\323\344\223\002&\"!/identity/v2/package/change-order:\001*'
  _globals['_PACKAGE'].methods_by_name['get']._loaded_options = None
  _globals['_PACKAGE'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002\035\"\030/identity/v2/package/get:\001*'
  _globals['_PACKAGE'].methods_by_name['list']._loaded_options = None
  _globals['_PACKAGE'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002\036\"\031/identity/v2/package/list:\001*'
  _globals['_PACKAGE'].methods_by_name['stat']._loaded_options = None
  _globals['_PACKAGE'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002\036\"\031/identity/v1/package/stat:\001*'
  _globals['_CREATEPACKAGEREQUEST']._serialized_start=191
  _globals['_CREATEPACKAGEREQUEST']._serialized_end=287
  _globals['_UPDATEPACKAGEREQUEST']._serialized_start=289
  _globals['_UPDATEPACKAGEREQUEST']._serialized_end=405
  _globals['_PACKAGEREQUEST']._serialized_start=407
  _globals['_PACKAGEREQUEST']._serialized_end=443
  _globals['_CHANGEPACKAGEORDERREQUEST']._serialized_start=445
  _globals['_CHANGEPACKAGEORDERREQUEST']._serialized_end=507
  _globals['_PACKAGEINFO']._serialized_start=510
  _globals['_PACKAGEINFO']._serialized_end=733
  _globals['_PACKAGESEARCHQUERY']._serialized_start=735
  _globals['_PACKAGESEARCHQUERY']._serialized_end=833
  _globals['_PACKAGESINFO']._serialized_start=835
  _globals['_PACKAGESINFO']._serialized_end=926
  _globals['_PACKAGESTATQUERY']._serialized_start=928
  _globals['_PACKAGESTATQUERY']._serialized_end=1000
  _globals['_PACKAGE']._serialized_start=1003
  _globals['_PACKAGE']._serialized_end=2074
# @@protoc_insertion_point(module_scope)