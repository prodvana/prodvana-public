# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/insights/insights.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n prodvana/insights/insights.proto\x12\x11prodvana.insights\"\xba\x02\n\x07Insight\x12\r\n\x05title\x18\x01 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x02 \x01(\t\x12\'\n\x05\x63lass\x18\x06 \x01(\x0e\x32\x18.prodvana.insights.Class\x12\x35\n\x07service\x18\x03 \x01(\x0b\x32\".prodvana.insights.Insight.SubjectH\x00\x12\x39\n\x0b\x61pplication\x18\x04 \x01(\x0b\x32\".prodvana.insights.Insight.SubjectH\x00\x12:\n\x0corganization\x18\x05 \x01(\x0b\x32\".prodvana.insights.Insight.SubjectH\x00\x1a#\n\x07Subject\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\tB\x0f\n\rsubject_oneof*\xce\x01\n\x05\x43lass\x12\x0b\n\x07UNKNOWN\x10\x00\x12,\n(SUCCESSFUL_DEPLOYMENT_FREQUENCY_DECREASE\x10\x01\x12$\n DEPLOYMENT_FAILURE_RATE_INCREASE\x10\x02\x12\x32\n.MEDIAN_SUCCESSFUL_DEPLOYMENT_DURATION_INCREASE\x10\x03\x12\x30\n,MEDIAN_SUCCESSFUL_ROLLBACK_DURATION_INCREASE\x10\x04\x42MZKgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/insightsb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'prodvana.insights.insights_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZKgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/insights'
  _globals['_CLASS']._serialized_start=373
  _globals['_CLASS']._serialized_end=579
  _globals['_INSIGHT']._serialized_start=56
  _globals['_INSIGHT']._serialized_end=370
  _globals['_INSIGHT_SUBJECT']._serialized_start=318
  _globals['_INSIGHT_SUBJECT']._serialized_end=353
# @@protoc_insertion_point(module_scope)
