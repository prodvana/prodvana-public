# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/common_config/version_push.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n)prodvana/common_config/version_push.proto\x12\x16prodvana.common_config\"\x9a\x01\n\x11VersionPushOption\x12<\n\x04type\x18\x01 \x01(\x0e\x32..prodvana.common_config.VersionPushOption.Type\x12\x0f\n\x07version\x18\x02 \x01(\t\"6\n\x04Type\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x0f\n\x0bPUSH_LATEST\x10\x01\x12\x10\n\x0cPUSH_VERSION\x10\x02\x42RZPgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_configb\x06proto3')



_VERSIONPUSHOPTION = DESCRIPTOR.message_types_by_name['VersionPushOption']
_VERSIONPUSHOPTION_TYPE = _VERSIONPUSHOPTION.enum_types_by_name['Type']
VersionPushOption = _reflection.GeneratedProtocolMessageType('VersionPushOption', (_message.Message,), {
  'DESCRIPTOR' : _VERSIONPUSHOPTION,
  '__module__' : 'prodvana.common_config.version_push_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.common_config.VersionPushOption)
  })
_sym_db.RegisterMessage(VersionPushOption)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZPgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config'
  _VERSIONPUSHOPTION._serialized_start=70
  _VERSIONPUSHOPTION._serialized_end=224
  _VERSIONPUSHOPTION_TYPE._serialized_start=170
  _VERSIONPUSHOPTION_TYPE._serialized_end=224
# @@protoc_insertion_point(module_scope)