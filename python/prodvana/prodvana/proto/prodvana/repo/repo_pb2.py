# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/repo/repo.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from prodvana.proto.prodvana.impact_analysis import impact_analysis_pb2 as prodvana_dot_impact__analysis_dot_impact__analysis__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x18prodvana/repo/repo.proto\x12\rprodvana.repo\x1a\x1fgoogle/protobuf/timestamp.proto\x1a.prodvana/impact_analysis/impact_analysis.proto\"#\n\x04User\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\r\n\x05\x65mail\x18\x02 \x01(\t\"\xdd\x01\n\x06\x43ommit\x12\x11\n\tcommit_id\x18\x01 \x01(\t\x12\x0b\n\x03url\x18\x02 \x01(\t\x12\x0f\n\x07message\x18\x03 \x01(\t\x12#\n\x06\x61uthor\x18\x04 \x01(\x0b\x32\x13.prodvana.repo.User\x12G\n\x0fimpact_analysis\x18\x05 \x01(\x0b\x32..prodvana.impact_analysis.ImpactAnalysisResult\x12\x34\n\x10\x63ommit_timestamp\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"1\n\x0b\x46uzzyCommit\x12\x12\n\ncommit_ish\x18\x01 \x01(\t\x12\x0e\n\x06source\x18\x02 \x01(\tBIZGgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/repob\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'prodvana.repo.repo_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZGgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/repo'
  _globals['_USER']._serialized_start=124
  _globals['_USER']._serialized_end=159
  _globals['_COMMIT']._serialized_start=162
  _globals['_COMMIT']._serialized_end=383
  _globals['_FUZZYCOMMIT']._serialized_start=385
  _globals['_FUZZYCOMMIT']._serialized_end=434
# @@protoc_insertion_point(module_scope)
