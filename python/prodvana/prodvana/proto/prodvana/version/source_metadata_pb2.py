# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/version/source_metadata.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n&prodvana/version/source_metadata.proto\x12\x10prodvana.version\"V\n\x0eSourceMetadata\x12\x10\n\x08repo_url\x18\x01 \x01(\t\x12\x11\n\tfile_path\x18\x02 \x01(\t\x12\x0e\n\x06\x63ommit\x18\x03 \x01(\t\x12\x0f\n\x07user_id\x18\x04 \x01(\t*_\n\x06Source\x12\x12\n\x0eUNKNOWN_SOURCE\x10\x00\x12\x07\n\x03WEB\x10\x01\x12\x16\n\x12INTERACTIVE_PVNCTL\x10\x02\x12\x0f\n\x0b\x43ONFIG_FILE\x10\x03\x12\x0f\n\x0bREPO_FOLLOW\x10\x04\x42LZJgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/versionb\x06proto3')

_SOURCE = DESCRIPTOR.enum_types_by_name['Source']
Source = enum_type_wrapper.EnumTypeWrapper(_SOURCE)
UNKNOWN_SOURCE = 0
WEB = 1
INTERACTIVE_PVNCTL = 2
CONFIG_FILE = 3
REPO_FOLLOW = 4


_SOURCEMETADATA = DESCRIPTOR.message_types_by_name['SourceMetadata']
SourceMetadata = _reflection.GeneratedProtocolMessageType('SourceMetadata', (_message.Message,), {
  'DESCRIPTOR' : _SOURCEMETADATA,
  '__module__' : 'prodvana.version.source_metadata_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.version.SourceMetadata)
  })
_sym_db.RegisterMessage(SourceMetadata)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZJgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/version'
  _SOURCE._serialized_start=148
  _SOURCE._serialized_end=243
  _SOURCEMETADATA._serialized_start=60
  _SOURCEMETADATA._serialized_end=146
# @@protoc_insertion_point(module_scope)
