# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/common_config/links.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from prodvana.proto.validate import validate_pb2 as validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\"prodvana/common_config/links.proto\x12\x16prodvana.common_config\x1a\x17validate/validate.proto\";\n\x04Link\x12\x14\n\x03url\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x1d\n\x0c\x64isplay_name\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x42RZPgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_configb\x06proto3')



_LINK = DESCRIPTOR.message_types_by_name['Link']
Link = _reflection.GeneratedProtocolMessageType('Link', (_message.Message,), {
  'DESCRIPTOR' : _LINK,
  '__module__' : 'prodvana.common_config.links_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.common_config.Link)
  })
_sym_db.RegisterMessage(Link)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZPgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config'
  _LINK.fields_by_name['url']._options = None
  _LINK.fields_by_name['url']._serialized_options = b'\372B\004r\002\020\001'
  _LINK.fields_by_name['display_name']._options = None
  _LINK.fields_by_name['display_name']._serialized_options = b'\372B\004r\002\020\001'
  _LINK._serialized_start=87
  _LINK._serialized_end=146
# @@protoc_insertion_point(module_scope)
