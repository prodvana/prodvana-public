# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/users/users.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1aprodvana/users/users.proto\x12\x0eprodvana.users\"\\\n\x04User\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x12\n\nfirst_name\x18\x03 \x01(\t\x12\x11\n\tlast_name\x18\x04 \x01(\t\x12\x0e\n\x06\x65mails\x18\x05 \x03(\tBJZHgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/usersb\x06proto3')



_USER = DESCRIPTOR.message_types_by_name['User']
User = _reflection.GeneratedProtocolMessageType('User', (_message.Message,), {
  'DESCRIPTOR' : _USER,
  '__module__' : 'prodvana.users.users_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.users.User)
  })
_sym_db.RegisterMessage(User)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZHgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/users'
  _USER._serialized_start=46
  _USER._serialized_end=138
# @@protoc_insertion_point(module_scope)
