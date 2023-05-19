# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/common_config/external_config.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from prodvana.proto.validate import validate_pb2 as validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n,prodvana/common_config/external_config.proto\x12\x16prodvana.common_config\x1a\x17validate/validate.proto\"$\n\x0bLocalConfig\x12\x15\n\x04path\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\"\xa9\x03\n\x0e\x45xternalConfig\x12\x39\n\x04type\x18\x01 \x01(\x0e\x32+.prodvana.common_config.ExternalConfig.Type\x12\x1a\n\x07inlined\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01H\x00\x12\x34\n\x05local\x18\x03 \x01(\x0b\x32#.prodvana.common_config.LocalConfigH\x00\x12S\n\x12\x65nv_injection_mode\x18\x04 \x01(\x0e\x32\x37.prodvana.common_config.ExternalConfig.EnvInjectionMode\"<\n\x04Type\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x0e\n\nKUBERNETES\x10\x01\x12\r\n\tKUSTOMIZE\x10\x02\x12\x08\n\x04HELM\x10\x03\"b\n\x10\x45nvInjectionMode\x12\x16\n\x12\x45NV_INJECT_UNKNOWN\x10\x00\x12\x17\n\x13\x45NV_INJECT_DISABLED\x10\x01\x12\x1d\n\x19\x45NV_INJECT_NON_SECRET_ENV\x10\x02\x42\x13\n\x0csource_oneof\x12\x03\xf8\x42\x01\x42RZPgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_configb\x06proto3')



_LOCALCONFIG = DESCRIPTOR.message_types_by_name['LocalConfig']
_EXTERNALCONFIG = DESCRIPTOR.message_types_by_name['ExternalConfig']
_EXTERNALCONFIG_TYPE = _EXTERNALCONFIG.enum_types_by_name['Type']
_EXTERNALCONFIG_ENVINJECTIONMODE = _EXTERNALCONFIG.enum_types_by_name['EnvInjectionMode']
LocalConfig = _reflection.GeneratedProtocolMessageType('LocalConfig', (_message.Message,), {
  'DESCRIPTOR' : _LOCALCONFIG,
  '__module__' : 'prodvana.common_config.external_config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.common_config.LocalConfig)
  })
_sym_db.RegisterMessage(LocalConfig)

ExternalConfig = _reflection.GeneratedProtocolMessageType('ExternalConfig', (_message.Message,), {
  'DESCRIPTOR' : _EXTERNALCONFIG,
  '__module__' : 'prodvana.common_config.external_config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.common_config.ExternalConfig)
  })
_sym_db.RegisterMessage(ExternalConfig)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZPgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config'
  _LOCALCONFIG.fields_by_name['path']._options = None
  _LOCALCONFIG.fields_by_name['path']._serialized_options = b'\372B\004r\002\020\001'
  _EXTERNALCONFIG.oneofs_by_name['source_oneof']._options = None
  _EXTERNALCONFIG.oneofs_by_name['source_oneof']._serialized_options = b'\370B\001'
  _EXTERNALCONFIG.fields_by_name['inlined']._options = None
  _EXTERNALCONFIG.fields_by_name['inlined']._serialized_options = b'\372B\004r\002\020\001'
  _LOCALCONFIG._serialized_start=97
  _LOCALCONFIG._serialized_end=133
  _EXTERNALCONFIG._serialized_start=136
  _EXTERNALCONFIG._serialized_end=561
  _EXTERNALCONFIG_TYPE._serialized_start=380
  _EXTERNALCONFIG_TYPE._serialized_end=440
  _EXTERNALCONFIG_ENVINJECTIONMODE._serialized_start=442
  _EXTERNALCONFIG_ENVINJECTIONMODE._serialized_end=540
# @@protoc_insertion_point(module_scope)
