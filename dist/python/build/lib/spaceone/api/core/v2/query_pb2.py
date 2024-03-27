# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/core/v2/query.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n spaceone/api/core/v2/query.proto\x12\x14spaceone.api.core.v2\x1a\x1cgoogle/protobuf/struct.proto\"\xc1\x01\n\x06\x46ilter\x12\r\n\x03key\x18\x01 \x01(\tH\x00\x12\x0b\n\x01k\x18\x02 \x01(\tH\x00\x12\'\n\x05value\x18\x03 \x01(\x0b\x32\x16.google.protobuf.ValueH\x01\x12#\n\x01v\x18\x04 \x01(\x0b\x32\x16.google.protobuf.ValueH\x01\x12\x12\n\x08operator\x18\x05 \x01(\tH\x02\x12\x0b\n\x01o\x18\x06 \x01(\tH\x02\x42\x0b\n\tkey_aliasB\r\n\x0bvalue_aliasB\x10\n\x0eoperator_alias\"!\n\x04Sort\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\x0c\n\x04\x64\x65sc\x18\x02 \x01(\x08\"$\n\x04Page\x12\r\n\x05start\x18\x01 \x01(\r\x12\r\n\x05limit\x18\x02 \x01(\r\"D\n\x06Unwind\x12\x0c\n\x04path\x18\x01 \x01(\t\x12,\n\x06\x66ilter\x18\x02 \x03(\x0b\x32\x1c.spaceone.api.core.v2.Filter\"\xac\x02\n\x05Query\x12,\n\x06\x66ilter\x18\x01 \x03(\x0b\x32\x1c.spaceone.api.core.v2.Filter\x12/\n\tfilter_or\x18\x02 \x03(\x0b\x32\x1c.spaceone.api.core.v2.Filter\x12(\n\x04sort\x18\x03 \x03(\x0b\x32\x1a.spaceone.api.core.v2.Sort\x12(\n\x04page\x18\x04 \x01(\x0b\x32\x1a.spaceone.api.core.v2.Page\x12\x0f\n\x07minimal\x18\x05 \x01(\x08\x12\x12\n\ncount_only\x18\x06 \x01(\x08\x12\x0c\n\x04only\x18\x07 \x03(\t\x12\x0f\n\x07keyword\x18\x08 \x01(\t\x12,\n\x06unwind\x18\t \x01(\x0b\x32\x1c.spaceone.api.core.v2.Unwind\"|\n\x11\x41ggregateGroupKey\x12\r\n\x03key\x18\x01 \x01(\tH\x00\x12\x0b\n\x01k\x18\x02 \x01(\tH\x00\x12\x0e\n\x04name\x18\x03 \x01(\tH\x01\x12\x0b\n\x01n\x18\x04 \x01(\tH\x01\x12\x13\n\x0b\x64\x61te_format\x18\x05 \x01(\tB\x0b\n\tkey_aliasB\x0c\n\nname_alias\"l\n\x16\x41ggregateGroupSubField\x12\r\n\x03key\x18\x01 \x01(\tH\x00\x12\x0b\n\x01k\x18\x02 \x01(\tH\x00\x12\x0e\n\x04name\x18\x03 \x01(\tH\x01\x12\x0b\n\x01n\x18\x04 \x01(\tH\x01\x42\x0b\n\tkey_aliasB\x0c\n\nname_alias\"\xd0\x01\n\x15\x41ggregateSubCondition\x12\r\n\x03key\x18\x01 \x01(\tH\x00\x12\x0b\n\x01k\x18\x02 \x01(\tH\x00\x12\'\n\x05value\x18\x03 \x01(\x0b\x32\x16.google.protobuf.ValueH\x01\x12#\n\x01v\x18\x04 \x01(\x0b\x32\x16.google.protobuf.ValueH\x01\x12\x12\n\x08operator\x18\x05 \x01(\tH\x02\x12\x0b\n\x01o\x18\x06 \x01(\tH\x02\x42\x0b\n\tkey_aliasB\r\n\x0bvalue_aliasB\x10\n\x0eoperator_alias\"\x9b\x02\n\x13\x41ggregateGroupField\x12\r\n\x03key\x18\x01 \x01(\tH\x00\x12\x0b\n\x01k\x18\x02 \x01(\tH\x00\x12\x0e\n\x04name\x18\x03 \x01(\tH\x01\x12\x0b\n\x01n\x18\x04 \x01(\tH\x01\x12\x12\n\x08operator\x18\x05 \x01(\tH\x02\x12\x0b\n\x01o\x18\x06 \x01(\tH\x02\x12<\n\x06\x66ields\x18\x07 \x03(\x0b\x32,.spaceone.api.core.v2.AggregateGroupSubField\x12?\n\nconditions\x18\x08 \x03(\x0b\x32+.spaceone.api.core.v2.AggregateSubConditionB\x0b\n\tkey_aliasB\x0c\n\nname_aliasB\x10\n\x0eoperator_alias\"\x82\x01\n\x0e\x41ggregateGroup\x12\x35\n\x04keys\x18\x01 \x03(\x0b\x32\'.spaceone.api.core.v2.AggregateGroupKey\x12\x39\n\x06\x66ields\x18\x02 \x03(\x0b\x32).spaceone.api.core.v2.AggregateGroupField\"\x9e\x01\n\x15\x41ggregateProjectField\x12\r\n\x03key\x18\x01 \x01(\tH\x00\x12\x0b\n\x01k\x18\x02 \x01(\tH\x00\x12\x0e\n\x04name\x18\x03 \x01(\tH\x01\x12\x0b\n\x01n\x18\x04 \x01(\tH\x01\x12\x12\n\x08operator\x18\x05 \x01(\tH\x02\x12\x0b\n\x01o\x18\x06 \x01(\tH\x02\x42\x0b\n\tkey_aliasB\x0c\n\nname_aliasB\x10\n\x0eoperator_alias\"e\n\x10\x41ggregateProject\x12;\n\x06\x66ields\x18\x01 \x03(\x0b\x32+.spaceone.api.core.v2.AggregateProjectField\x12\x14\n\x0c\x65xclude_keys\x18\x02 \x01(\x08\"\x1f\n\x0f\x41ggregateUnwind\x12\x0c\n\x04path\x18\x01 \x01(\t\"\x1e\n\x0e\x41ggregateCount\x12\x0c\n\x04name\x18\x01 \x01(\t\"\xd7\x02\n\x13StatisticsAggregate\x12\x37\n\x06unwind\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.AggregateUnwindH\x00\x12\x35\n\x05group\x18\x02 \x01(\x0b\x32$.spaceone.api.core.v2.AggregateGroupH\x00\x12\x35\n\x05\x63ount\x18\x03 \x01(\x0b\x32$.spaceone.api.core.v2.AggregateCountH\x00\x12*\n\x04sort\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.ListValueH\x00\x12\x39\n\x07project\x18\x05 \x01(\x0b\x32&.spaceone.api.core.v2.AggregateProjectH\x00\x12\x0f\n\x05limit\x18\x06 \x01(\x05H\x00\x12\x0e\n\x04skip\x18\x07 \x01(\x05H\x00\x42\x11\n\x0f\x61ggregate_alias\"\xfb\x01\n\x0fStatisticsQuery\x12,\n\x06\x66ilter\x18\x01 \x03(\x0b\x32\x1c.spaceone.api.core.v2.Filter\x12/\n\tfilter_or\x18\x02 \x03(\x0b\x32\x1c.spaceone.api.core.v2.Filter\x12<\n\taggregate\x18\x03 \x03(\x0b\x32).spaceone.api.core.v2.StatisticsAggregate\x12(\n\x04page\x18\x04 \x01(\x0b\x32\x1a.spaceone.api.core.v2.Page\x12\x10\n\x08\x64istinct\x18\x05 \x01(\t\x12\x0f\n\x07keyword\x18\x06 \x01(\t\"\x89\x04\n\x16TimeSeriesAnalyzeQuery\x12M\n\x0bgranularity\x18\x01 \x01(\x0e\x32\x38.spaceone.api.core.v2.TimeSeriesAnalyzeQuery.Granularity\x12\r\n\x05start\x18\x02 \x01(\t\x12\x0b\n\x03\x65nd\x18\x03 \x01(\t\x12\x10\n\x08group_by\x18\x04 \x03(\t\x12\x13\n\x0b\x66ield_group\x18\x05 \x03(\t\x12,\n\x06\x66ilter\x18\x06 \x03(\x0b\x32\x1c.spaceone.api.core.v2.Filter\x12/\n\tfilter_or\x18\x07 \x03(\x0b\x32\x1c.spaceone.api.core.v2.Filter\x12(\n\x04page\x18\x08 \x01(\x0b\x32\x1a.spaceone.api.core.v2.Page\x12(\n\x04sort\x18\t \x03(\x0b\x32\x1a.spaceone.api.core.v2.Sort\x12\'\n\x06\x66ields\x18\n \x01(\x0b\x32\x17.google.protobuf.Struct\x12\'\n\x06select\x18\x0b \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x0f\n\x07keyword\x18\x0c \x01(\t\"G\n\x0bGranularity\x12\x14\n\x10GRANULARITY_NONE\x10\x00\x12\t\n\x05\x44\x41ILY\x10\x01\x12\x0b\n\x07MONTHLY\x10\x02\x12\n\n\x06YEARLY\x10\x03\"\xcb\x02\n\x0c\x41nalyzeQuery\x12\x10\n\x08group_by\x18\x01 \x03(\t\x12\x13\n\x0b\x66ield_group\x18\x02 \x03(\t\x12,\n\x06\x66ilter\x18\x03 \x03(\x0b\x32\x1c.spaceone.api.core.v2.Filter\x12/\n\tfilter_or\x18\x04 \x03(\x0b\x32\x1c.spaceone.api.core.v2.Filter\x12(\n\x04page\x18\x05 \x01(\x0b\x32\x1a.spaceone.api.core.v2.Page\x12(\n\x04sort\x18\x06 \x03(\x0b\x32\x1a.spaceone.api.core.v2.Sort\x12\'\n\x06\x66ields\x18\x07 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\'\n\x06select\x18\x08 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x0f\n\x07keyword\x18\t \x01(\t\"\xb1\x02\n\x11\x45xportSearchQuery\x12,\n\x06\x66ilter\x18\x01 \x03(\x0b\x32\x1c.spaceone.api.core.v2.Filter\x12/\n\tfilter_or\x18\x02 \x03(\x0b\x32\x1c.spaceone.api.core.v2.Filter\x12\x0f\n\x07keyword\x18\x03 \x01(\t\x12(\n\x04sort\x18\x04 \x03(\x0b\x32\x1a.spaceone.api.core.v2.Sort\x12*\n\x06\x66ields\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12,\n\x06unwind\x18\x06 \x01(\x0b\x32\x1c.spaceone.api.core.v2.Unwind\x12(\n\x04page\x18\x07 \x01(\x0b\x32\x1a.spaceone.api.core.v2.Page\"\xbc\x02\n\x12\x45xportAnalyzeQuery\x12,\n\x06\x66ilter\x18\x01 \x03(\x0b\x32\x1c.spaceone.api.core.v2.Filter\x12/\n\tfilter_or\x18\x02 \x03(\x0b\x32\x1c.spaceone.api.core.v2.Filter\x12\x0f\n\x07keyword\x18\x03 \x01(\t\x12(\n\x04sort\x18\x04 \x03(\x0b\x32\x1a.spaceone.api.core.v2.Sort\x12\x10\n\x08group_by\x18\x05 \x03(\t\x12\'\n\x06\x66ields\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\'\n\x06select\x18\x07 \x01(\x0b\x32\x17.google.protobuf.Struct\x12(\n\x04page\x18\x08 \x01(\x0b\x32\x1a.spaceone.api.core.v2.Page\"\xbc\x02\n\x0c\x45xportOption\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\r\n\x05title\x18\x02 \x01(\t\x12@\n\nquery_type\x18\x0b \x01(\x0e\x32,.spaceone.api.core.v2.ExportOption.QueryType\x12?\n\x0csearch_query\x18\x0c \x01(\x0b\x32\'.spaceone.api.core.v2.ExportSearchQueryH\x00\x12\x41\n\ranalyze_query\x18\r \x01(\x0b\x32(.spaceone.api.core.v2.ExportAnalyzeQueryH\x00\"9\n\tQueryType\x12\x13\n\x0fQUERY_TYPE_NONE\x10\x00\x12\n\n\x06SEARCH\x10\x01\x12\x0b\n\x07\x41NALYZE\x10\x02\x42\x0e\n\x0c\x65xport_queryB;Z9github.com/cloudforet-io/api/dist/go/spaceone/api/core/v2b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.core.v2.query_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z9github.com/cloudforet-io/api/dist/go/spaceone/api/core/v2'
  _globals['_FILTER']._serialized_start=89
  _globals['_FILTER']._serialized_end=282
  _globals['_SORT']._serialized_start=284
  _globals['_SORT']._serialized_end=317
  _globals['_PAGE']._serialized_start=319
  _globals['_PAGE']._serialized_end=355
  _globals['_UNWIND']._serialized_start=357
  _globals['_UNWIND']._serialized_end=425
  _globals['_QUERY']._serialized_start=428
  _globals['_QUERY']._serialized_end=728
  _globals['_AGGREGATEGROUPKEY']._serialized_start=730
  _globals['_AGGREGATEGROUPKEY']._serialized_end=854
  _globals['_AGGREGATEGROUPSUBFIELD']._serialized_start=856
  _globals['_AGGREGATEGROUPSUBFIELD']._serialized_end=964
  _globals['_AGGREGATESUBCONDITION']._serialized_start=967
  _globals['_AGGREGATESUBCONDITION']._serialized_end=1175
  _globals['_AGGREGATEGROUPFIELD']._serialized_start=1178
  _globals['_AGGREGATEGROUPFIELD']._serialized_end=1461
  _globals['_AGGREGATEGROUP']._serialized_start=1464
  _globals['_AGGREGATEGROUP']._serialized_end=1594
  _globals['_AGGREGATEPROJECTFIELD']._serialized_start=1597
  _globals['_AGGREGATEPROJECTFIELD']._serialized_end=1755
  _globals['_AGGREGATEPROJECT']._serialized_start=1757
  _globals['_AGGREGATEPROJECT']._serialized_end=1858
  _globals['_AGGREGATEUNWIND']._serialized_start=1860
  _globals['_AGGREGATEUNWIND']._serialized_end=1891
  _globals['_AGGREGATECOUNT']._serialized_start=1893
  _globals['_AGGREGATECOUNT']._serialized_end=1923
  _globals['_STATISTICSAGGREGATE']._serialized_start=1926
  _globals['_STATISTICSAGGREGATE']._serialized_end=2269
  _globals['_STATISTICSQUERY']._serialized_start=2272
  _globals['_STATISTICSQUERY']._serialized_end=2523
  _globals['_TIMESERIESANALYZEQUERY']._serialized_start=2526
  _globals['_TIMESERIESANALYZEQUERY']._serialized_end=3047
  _globals['_TIMESERIESANALYZEQUERY_GRANULARITY']._serialized_start=2976
  _globals['_TIMESERIESANALYZEQUERY_GRANULARITY']._serialized_end=3047
  _globals['_ANALYZEQUERY']._serialized_start=3050
  _globals['_ANALYZEQUERY']._serialized_end=3381
  _globals['_EXPORTSEARCHQUERY']._serialized_start=3384
  _globals['_EXPORTSEARCHQUERY']._serialized_end=3689
  _globals['_EXPORTANALYZEQUERY']._serialized_start=3692
  _globals['_EXPORTANALYZEQUERY']._serialized_end=4008
  _globals['_EXPORTOPTION']._serialized_start=4011
  _globals['_EXPORTOPTION']._serialized_end=4327
  _globals['_EXPORTOPTION_QUERYTYPE']._serialized_start=4254
  _globals['_EXPORTOPTION_QUERYTYPE']._serialized_end=4311
# @@protoc_insertion_point(module_scope)
