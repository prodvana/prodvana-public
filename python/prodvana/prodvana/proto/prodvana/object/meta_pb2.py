# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/object/meta.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from prodvana.proto.prodvana.version import source_metadata_pb2 as prodvana_dot_version_dot_source__metadata__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1aprodvana/object/meta.proto\x12\x0fprodvana.object\x1a&prodvana/version/source_metadata.proto\"\xdf\x01\n\nObjectMeta\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x0f\n\x07version\x18\x03 \x01(\t\x12\x16\n\x0e\x63onfig_version\x18\x04 \x01(\t\x12(\n\x06source\x18\x05 \x01(\x0e\x32\x18.prodvana.version.Source\x12\x39\n\x0fsource_metadata\x18\x06 \x01(\x0b\x32 .prodvana.version.SourceMetadata\x12)\n\x04type\x18\x07 \x01(\x0e\x32\x1b.prodvana.object.ObjectType*\x92\x01\n\nObjectType\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x0b\n\x07RUNTIME\x10\x01\x12\x0f\n\x0b\x41PPLICATION\x10\x02\x12\x0b\n\x07SERVICE\x10\x03\x12\x13\n\x0fRELEASE_CHANNEL\x10\x04\x12\x0e\n\nPROTECTION\x10\x05\x12\x0e\n\nDEPLOYMENT\x10\x06\x12\x0b\n\x07RELEASE\x10\x07\x12\n\n\x06RECIPE\x10\x08\x42KZIgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/objectb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'prodvana.object.meta_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZIgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/object'
  _globals['_OBJECTTYPE']._serialized_start=314
  _globals['_OBJECTTYPE']._serialized_end=460
  _globals['_OBJECTMETA']._serialized_start=88
  _globals['_OBJECTMETA']._serialized_end=311
# @@protoc_insertion_point(module_scope)
