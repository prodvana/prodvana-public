# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/service/user_metadata.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from prodvana.proto.prodvana.common_config import links_pb2 as prodvana_dot_common__config_dot_links__pb2
from prodvana.proto.validate import validate_pb2 as validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n$prodvana/service/user_metadata.proto\x12\x10prodvana.service\x1a\"prodvana/common_config/links.proto\x1a\x17validate/validate.proto\"\xb6\x01\n\x13ServiceUserMetadata\x12\x13\n\x0b\x64\x65scription\x18\x01 \x01(\t\x12:\n\x05links\x18\x02 \x03(\x0b\x32\x1c.prodvana.common_config.LinkB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\x12N\n\x11\x66ollow_repository\x18\x03 \x01(\x0b\x32\x33.prodvana.service.FollowContainerRepositorySettings\"4\n!FollowContainerRepositorySettings\x12\x0f\n\x07\x65nabled\x18\x01 \x01(\x08\x42LZJgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/serviceb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'prodvana.service.user_metadata_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZJgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/service'
  _SERVICEUSERMETADATA.fields_by_name['links']._options = None
  _SERVICEUSERMETADATA.fields_by_name['links']._serialized_options = b'\372B\n\222\001\007\"\005\212\001\002\020\001'
  _globals['_SERVICEUSERMETADATA']._serialized_start=120
  _globals['_SERVICEUSERMETADATA']._serialized_end=302
  _globals['_FOLLOWCONTAINERREPOSITORYSETTINGS']._serialized_start=304
  _globals['_FOLLOWCONTAINERREPOSITORYSETTINGS']._serialized_end=356
# @@protoc_insertion_point(module_scope)
