# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/object/meta.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from prodvana.proto.prodvana.version import source_metadata_pb2 as prodvana_dot_version_dot_source__metadata__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1aprodvana/object/meta.proto\x12\x0fprodvana.object\x1a&prodvana/version/source_metadata.proto\"\xb4\x01\n\nObjectMeta\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x0f\n\x07version\x18\x03 \x01(\t\x12\x16\n\x0e\x63onfig_version\x18\x04 \x01(\t\x12(\n\x06source\x18\x05 \x01(\x0e\x32\x18.prodvana.version.Source\x12\x39\n\x0fsource_metadata\x18\x06 \x01(\x0b\x32 .prodvana.version.SourceMetadataBKZIgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/objectb\x06proto3')



_OBJECTMETA = DESCRIPTOR.message_types_by_name['ObjectMeta']
ObjectMeta = _reflection.GeneratedProtocolMessageType('ObjectMeta', (_message.Message,), {
  'DESCRIPTOR' : _OBJECTMETA,
  '__module__' : 'prodvana.object.meta_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.object.ObjectMeta)
  })
_sym_db.RegisterMessage(ObjectMeta)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZIgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/object'
  _OBJECTMETA._serialized_start=88
  _OBJECTMETA._serialized_end=268
# @@protoc_insertion_point(module_scope)