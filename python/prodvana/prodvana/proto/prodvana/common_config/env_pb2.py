# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/common_config/env.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from prodvana.proto.validate import validate_pb2 as validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n prodvana/common_config/env.proto\x12\x16prodvana.common_config\x1a\x17validate/validate.proto\"\xbe\x01\n\x08\x45nvValue\x12\x0f\n\x05value\x18\x01 \x01(\tH\x00\x12\x14\n\nraw_secret\x18\x02 \x01(\tH\x00\x12\x30\n\x06secret\x18\x03 \x01(\x0b\x32\x1e.prodvana.common_config.SecretH\x00\x12\x45\n\x11kubernetes_secret\x18\x04 \x01(\x0b\x32(.prodvana.common_config.KubernetesSecretH\x00\x42\x12\n\x0bvalue_oneof\x12\x03\xf8\x42\x01\"\xa3\x01\n\x14SecretReferenceValue\x12\x30\n\x06secret\x18\x01 \x01(\x0b\x32\x1e.prodvana.common_config.SecretH\x00\x12\x45\n\x11kubernetes_secret\x18\x02 \x01(\x0b\x32(.prodvana.common_config.KubernetesSecretH\x00\x42\x12\n\x0bvalue_oneof\x12\x03\xf8\x42\x01\"8\n\x06Secret\x12\x14\n\x03key\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x18\n\x07version\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\"F\n\x10KubernetesSecret\x12\x1c\n\x0bsecret_name\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x14\n\x03key\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x42RZPgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_configb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'prodvana.common_config.env_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZPgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config'
  _ENVVALUE.oneofs_by_name['value_oneof']._options = None
  _ENVVALUE.oneofs_by_name['value_oneof']._serialized_options = b'\370B\001'
  _SECRETREFERENCEVALUE.oneofs_by_name['value_oneof']._options = None
  _SECRETREFERENCEVALUE.oneofs_by_name['value_oneof']._serialized_options = b'\370B\001'
  _SECRET.fields_by_name['key']._options = None
  _SECRET.fields_by_name['key']._serialized_options = b'\372B\004r\002\020\001'
  _SECRET.fields_by_name['version']._options = None
  _SECRET.fields_by_name['version']._serialized_options = b'\372B\004r\002\020\001'
  _KUBERNETESSECRET.fields_by_name['secret_name']._options = None
  _KUBERNETESSECRET.fields_by_name['secret_name']._serialized_options = b'\372B\004r\002\020\001'
  _KUBERNETESSECRET.fields_by_name['key']._options = None
  _KUBERNETESSECRET.fields_by_name['key']._serialized_options = b'\372B\004r\002\020\001'
  _globals['_ENVVALUE']._serialized_start=86
  _globals['_ENVVALUE']._serialized_end=276
  _globals['_SECRETREFERENCEVALUE']._serialized_start=279
  _globals['_SECRETREFERENCEVALUE']._serialized_end=442
  _globals['_SECRET']._serialized_start=444
  _globals['_SECRET']._serialized_end=500
  _globals['_KUBERNETESSECRET']._serialized_start=502
  _globals['_KUBERNETESSECRET']._serialized_end=572
# @@protoc_insertion_point(module_scope)
