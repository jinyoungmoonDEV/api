# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/repository/v2/provider.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from spaceone.api.core.v1 import query_pb2 as spaceone_dot_api_dot_core_dot_v1_dot_query__pb2
from spaceone.api.repository.v2 import common_pb2 as spaceone_dot_api_dot_repository_dot_v2_dot_common__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n)spaceone/api/repository/v2/provider.proto\x12\x1aspaceone.api.repository.v2\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v1/query.proto\x1a\'spaceone/api/repository/v2/common.proto\"\xa2\x04\n\x15\x43reateProviderRequest\x12\x10\n\x08provider\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x37\n\tsync_mode\x18\x03 \x01(\x0e\x32$.spaceone.api.repository.v2.SyncMode\x12=\n\x0csync_options\x18\x04 \x01(\x0b\x32\'.spaceone.api.repository.v2.SyncOptions\x12<\n\x0b\x64\x65scription\x18\x05 \x03(\x0b\x32\'.spaceone.api.repository.v2.Description\x12:\n\x06schema\x18\x06 \x03(\x0b\x32*.spaceone.api.repository.v2.ProviderSchema\x12:\n\ncapability\x18\x07 \x01(\x0b\x32&.spaceone.api.repository.v2.Capability\x12\r\n\x05\x63olor\x18\x08 \x01(\t\x12\x0c\n\x04icon\x18\t \x01(\t\x12\x38\n\treference\x18\n \x03(\x0b\x32%.spaceone.api.repository.v2.Reference\x12*\n\x06labels\x18\x0b \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12%\n\x04tags\x18\x0c \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\"\xa2\x04\n\x15UpdateProviderRequest\x12\x10\n\x08provider\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x37\n\tsync_mode\x18\x03 \x01(\x0e\x32$.spaceone.api.repository.v2.SyncMode\x12=\n\x0csync_options\x18\x04 \x01(\x0b\x32\'.spaceone.api.repository.v2.SyncOptions\x12<\n\x0b\x64\x65scription\x18\x05 \x03(\x0b\x32\'.spaceone.api.repository.v2.Description\x12:\n\x06schema\x18\x06 \x03(\x0b\x32*.spaceone.api.repository.v2.ProviderSchema\x12:\n\ncapability\x18\x07 \x01(\x0b\x32&.spaceone.api.repository.v2.Capability\x12\r\n\x05\x63olor\x18\x08 \x01(\t\x12\x0c\n\x04icon\x18\t \x01(\t\x12\x38\n\treference\x18\n \x03(\x0b\x32%.spaceone.api.repository.v2.Reference\x12*\n\x06labels\x18\x0b \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12%\n\x04tags\x18\x0c \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\"6\n\x0fProviderRequest\x12\x10\n\x08provider\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x15 \x01(\t\"G\n\x12GetProviderRequest\x12\x10\n\x08provider\x18\x01 \x01(\t\x12\x0c\n\x04only\x18\x02 \x03(\t\x12\x11\n\tdomain_id\x18\x15 \x01(\t\"\xc7\x01\n\rProviderQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v1.Query\x12\x10\n\x08provider\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x37\n\tsync_mode\x18\x04 \x01(\x0e\x32$.spaceone.api.repository.v2.SyncMode\x12\x1e\n\x16remote_repository_name\x18\x05 \x01(\t\x12\x11\n\tdomain_id\x18\x15 \x01(\t\"\xe1\x04\n\x0cProviderInfo\x12\x10\n\x08provider\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x37\n\tsync_mode\x18\x03 \x01(\x0e\x32$.spaceone.api.repository.v2.SyncMode\x12=\n\x0csync_options\x18\x04 \x01(\x0b\x32\'.spaceone.api.repository.v2.SyncOptions\x12<\n\x0b\x64\x65scription\x18\x05 \x03(\x0b\x32\'.spaceone.api.repository.v2.Description\x12:\n\x06schema\x18\x06 \x03(\x0b\x32*.spaceone.api.repository.v2.ProviderSchema\x12:\n\ncapability\x18\x07 \x01(\x0b\x32&.spaceone.api.repository.v2.Capability\x12\r\n\x05\x63olor\x18\x08 \x01(\t\x12\x0c\n\x04icon\x18\t \x01(\t\x12\x38\n\treference\x18\n \x03(\x0b\x32%.spaceone.api.repository.v2.Reference\x12*\n\x06labels\x18\x0b \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12%\n\x04tags\x18\x0c \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x32\n\x11remote_repository\x18\r \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\"_\n\rProvidersInfo\x12\x39\n\x07results\x18\x01 \x03(\x0b\x32(.spaceone.api.repository.v2.ProviderInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"-\n\nCapability\x12\x1f\n\x17trusted_service_account\x18\x01 \x01(\t\"O\n\x0eProviderSchema\x12\x15\n\rresource_type\x18\x01 \x01(\t\x12\x13\n\x0bsecret_type\x18\x02 \x01(\t\x12\x11\n\tschema_id\x18\x03 \x01(\t\"K\n\x0b\x44\x65scription\x12\x15\n\rresource_type\x18\x01 \x01(\t\x12%\n\x04\x62ody\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\"I\n\tReference\x12\x15\n\rresource_type\x18\x01 \x01(\t\x12%\n\x04link\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct2\xb4\x06\n\x08Provider\x12\x8d\x01\n\x06\x63reate\x12\x31.spaceone.api.repository.v2.CreateProviderRequest\x1a(.spaceone.api.repository.v2.ProviderInfo\"&\x82\xd3\xe4\x93\x02 \"\x1e/repository/v2/provider/create\x12\x8d\x01\n\x06update\x12\x31.spaceone.api.repository.v2.UpdateProviderRequest\x1a(.spaceone.api.repository.v2.ProviderInfo\"&\x82\xd3\xe4\x93\x02 \"\x1e/repository/v2/provider/update\x12\x83\x01\n\x04sync\x12+.spaceone.api.repository.v2.ProviderRequest\x1a(.spaceone.api.repository.v2.ProviderInfo\"$\x82\xd3\xe4\x93\x02\x1e\"\x1c/repository/v2/provider/sync\x12u\n\x06\x64\x65lete\x12+.spaceone.api.repository.v2.ProviderRequest\x1a\x16.google.protobuf.Empty\"&\x82\xd3\xe4\x93\x02 \"\x1e/repository/v2/provider/delete\x12\x84\x01\n\x03get\x12..spaceone.api.repository.v2.GetProviderRequest\x1a(.spaceone.api.repository.v2.ProviderInfo\"#\x82\xd3\xe4\x93\x02\x1d\"\x1b/repository/v2/provider/get\x12\x83\x01\n\x04list\x12).spaceone.api.repository.v2.ProviderQuery\x1a).spaceone.api.repository.v2.ProvidersInfo\"%\x82\xd3\xe4\x93\x02\x1f\"\x1d/repository/v2/providers/listBAZ?github.com/cloudforet-io/api/dist/go/spaceone/api/repository/v2b\x06proto3')



_CREATEPROVIDERREQUEST = DESCRIPTOR.message_types_by_name['CreateProviderRequest']
_UPDATEPROVIDERREQUEST = DESCRIPTOR.message_types_by_name['UpdateProviderRequest']
_PROVIDERREQUEST = DESCRIPTOR.message_types_by_name['ProviderRequest']
_GETPROVIDERREQUEST = DESCRIPTOR.message_types_by_name['GetProviderRequest']
_PROVIDERQUERY = DESCRIPTOR.message_types_by_name['ProviderQuery']
_PROVIDERINFO = DESCRIPTOR.message_types_by_name['ProviderInfo']
_PROVIDERSINFO = DESCRIPTOR.message_types_by_name['ProvidersInfo']
_CAPABILITY = DESCRIPTOR.message_types_by_name['Capability']
_PROVIDERSCHEMA = DESCRIPTOR.message_types_by_name['ProviderSchema']
_DESCRIPTION = DESCRIPTOR.message_types_by_name['Description']
_REFERENCE = DESCRIPTOR.message_types_by_name['Reference']
CreateProviderRequest = _reflection.GeneratedProtocolMessageType('CreateProviderRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEPROVIDERREQUEST,
  '__module__' : 'spaceone.api.repository.v2.provider_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.repository.v2.CreateProviderRequest)
  })
_sym_db.RegisterMessage(CreateProviderRequest)

UpdateProviderRequest = _reflection.GeneratedProtocolMessageType('UpdateProviderRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPDATEPROVIDERREQUEST,
  '__module__' : 'spaceone.api.repository.v2.provider_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.repository.v2.UpdateProviderRequest)
  })
_sym_db.RegisterMessage(UpdateProviderRequest)

ProviderRequest = _reflection.GeneratedProtocolMessageType('ProviderRequest', (_message.Message,), {
  'DESCRIPTOR' : _PROVIDERREQUEST,
  '__module__' : 'spaceone.api.repository.v2.provider_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.repository.v2.ProviderRequest)
  })
_sym_db.RegisterMessage(ProviderRequest)

GetProviderRequest = _reflection.GeneratedProtocolMessageType('GetProviderRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETPROVIDERREQUEST,
  '__module__' : 'spaceone.api.repository.v2.provider_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.repository.v2.GetProviderRequest)
  })
_sym_db.RegisterMessage(GetProviderRequest)

ProviderQuery = _reflection.GeneratedProtocolMessageType('ProviderQuery', (_message.Message,), {
  'DESCRIPTOR' : _PROVIDERQUERY,
  '__module__' : 'spaceone.api.repository.v2.provider_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.repository.v2.ProviderQuery)
  })
_sym_db.RegisterMessage(ProviderQuery)

ProviderInfo = _reflection.GeneratedProtocolMessageType('ProviderInfo', (_message.Message,), {
  'DESCRIPTOR' : _PROVIDERINFO,
  '__module__' : 'spaceone.api.repository.v2.provider_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.repository.v2.ProviderInfo)
  })
_sym_db.RegisterMessage(ProviderInfo)

ProvidersInfo = _reflection.GeneratedProtocolMessageType('ProvidersInfo', (_message.Message,), {
  'DESCRIPTOR' : _PROVIDERSINFO,
  '__module__' : 'spaceone.api.repository.v2.provider_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.repository.v2.ProvidersInfo)
  })
_sym_db.RegisterMessage(ProvidersInfo)

Capability = _reflection.GeneratedProtocolMessageType('Capability', (_message.Message,), {
  'DESCRIPTOR' : _CAPABILITY,
  '__module__' : 'spaceone.api.repository.v2.provider_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.repository.v2.Capability)
  })
_sym_db.RegisterMessage(Capability)

ProviderSchema = _reflection.GeneratedProtocolMessageType('ProviderSchema', (_message.Message,), {
  'DESCRIPTOR' : _PROVIDERSCHEMA,
  '__module__' : 'spaceone.api.repository.v2.provider_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.repository.v2.ProviderSchema)
  })
_sym_db.RegisterMessage(ProviderSchema)

Description = _reflection.GeneratedProtocolMessageType('Description', (_message.Message,), {
  'DESCRIPTOR' : _DESCRIPTION,
  '__module__' : 'spaceone.api.repository.v2.provider_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.repository.v2.Description)
  })
_sym_db.RegisterMessage(Description)

Reference = _reflection.GeneratedProtocolMessageType('Reference', (_message.Message,), {
  'DESCRIPTOR' : _REFERENCE,
  '__module__' : 'spaceone.api.repository.v2.provider_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.repository.v2.Reference)
  })
_sym_db.RegisterMessage(Reference)

_PROVIDER = DESCRIPTOR.services_by_name['Provider']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z?github.com/cloudforet-io/api/dist/go/spaceone/api/repository/v2'
  _PROVIDER.methods_by_name['create']._options = None
  _PROVIDER.methods_by_name['create']._serialized_options = b'\202\323\344\223\002 \"\036/repository/v2/provider/create'
  _PROVIDER.methods_by_name['update']._options = None
  _PROVIDER.methods_by_name['update']._serialized_options = b'\202\323\344\223\002 \"\036/repository/v2/provider/update'
  _PROVIDER.methods_by_name['sync']._options = None
  _PROVIDER.methods_by_name['sync']._serialized_options = b'\202\323\344\223\002\036\"\034/repository/v2/provider/sync'
  _PROVIDER.methods_by_name['delete']._options = None
  _PROVIDER.methods_by_name['delete']._serialized_options = b'\202\323\344\223\002 \"\036/repository/v2/provider/delete'
  _PROVIDER.methods_by_name['get']._options = None
  _PROVIDER.methods_by_name['get']._serialized_options = b'\202\323\344\223\002\035\"\033/repository/v2/provider/get'
  _PROVIDER.methods_by_name['list']._options = None
  _PROVIDER.methods_by_name['list']._serialized_options = b'\202\323\344\223\002\037\"\035/repository/v2/providers/list'
  _CREATEPROVIDERREQUEST._serialized_start=238
  _CREATEPROVIDERREQUEST._serialized_end=784
  _UPDATEPROVIDERREQUEST._serialized_start=787
  _UPDATEPROVIDERREQUEST._serialized_end=1333
  _PROVIDERREQUEST._serialized_start=1335
  _PROVIDERREQUEST._serialized_end=1389
  _GETPROVIDERREQUEST._serialized_start=1391
  _GETPROVIDERREQUEST._serialized_end=1462
  _PROVIDERQUERY._serialized_start=1465
  _PROVIDERQUERY._serialized_end=1664
  _PROVIDERINFO._serialized_start=1667
  _PROVIDERINFO._serialized_end=2276
  _PROVIDERSINFO._serialized_start=2278
  _PROVIDERSINFO._serialized_end=2373
  _CAPABILITY._serialized_start=2375
  _CAPABILITY._serialized_end=2420
  _PROVIDERSCHEMA._serialized_start=2422
  _PROVIDERSCHEMA._serialized_end=2501
  _DESCRIPTION._serialized_start=2503
  _DESCRIPTION._serialized_end=2578
  _REFERENCE._serialized_start=2580
  _REFERENCE._serialized_end=2653
  _PROVIDER._serialized_start=2656
  _PROVIDER._serialized_end=3476
# @@protoc_insertion_point(module_scope)
