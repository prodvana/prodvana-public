# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/common_config/constants.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from prodvana.proto.validate import validate_pb2 as validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n&prodvana/common_config/constants.proto\x12\x16prodvana.common_config\x1a\x17validate/validate.proto\"p\n\x08\x43onstant\x12\x15\n\x04name\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x38\n\x06string\x18\x03 \x01(\x0b\x32&.prodvana.common_config.StringConstantH\x00\x42\x13\n\x0c\x63onfig_oneof\x12\x03\xf8\x42\x01\"\x1f\n\x0eStringConstant\x12\r\n\x05value\x18\x01 \x01(\tBRZPgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_configb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'prodvana.common_config.constants_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZPgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config'
  _CONSTANT.oneofs_by_name['config_oneof']._options = None
  _CONSTANT.oneofs_by_name['config_oneof']._serialized_options = b'\370B\001'
  _CONSTANT.fields_by_name['name']._options = None
  _CONSTANT.fields_by_name['name']._serialized_options = b'\372B\004r\002\020\001'
  _globals['_CONSTANT']._serialized_start=91
  _globals['_CONSTANT']._serialized_end=203
  _globals['_STRINGCONSTANT']._serialized_start=205
  _globals['_STRINGCONSTANT']._serialized_end=236
# @@protoc_insertion_point(module_scope)