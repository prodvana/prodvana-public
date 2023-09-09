# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/common_config/kubernetes_config.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from prodvana.proto.validate import validate_pb2 as validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n.prodvana/common_config/kubernetes_config.proto\x12\x16prodvana.common_config\x1a\x17validate/validate.proto\"g\n\x0bLocalConfig\x12\x0e\n\x04path\x18\x01 \x01(\tH\x00\x12\x19\n\x0ftarball_blob_id\x18\x04 \x01(\tH\x00\x12\r\n\x05paths\x18\x02 \x03(\t\x12\x10\n\x08sub_path\x18\x03 \x01(\tB\x0c\n\npath_oneof\"\xaf\x03\n\x10KubernetesConfig\x12\x45\n\x04type\x18\x01 \x01(\x0e\x32-.prodvana.common_config.KubernetesConfig.TypeB\x08\xfa\x42\x05\x82\x01\x02 \x00\x12\x1a\n\x07inlined\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01H\x00\x12\x34\n\x05local\x18\x03 \x01(\x0b\x32#.prodvana.common_config.LocalConfigH\x00\x12U\n\x12\x65nv_injection_mode\x18\x04 \x01(\x0e\x32\x39.prodvana.common_config.KubernetesConfig.EnvInjectionMode\"2\n\x04Type\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x0e\n\nKUBERNETES\x10\x01\x12\r\n\tKUSTOMIZE\x10\x02\"b\n\x10\x45nvInjectionMode\x12\x16\n\x12\x45NV_INJECT_UNKNOWN\x10\x00\x12\x17\n\x13\x45NV_INJECT_DISABLED\x10\x01\x12\x1d\n\x19\x45NV_INJECT_NON_SECRET_ENV\x10\x02\x42\x13\n\x0csource_oneof\x12\x03\xf8\x42\x01\x42RZPgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_configb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'prodvana.common_config.kubernetes_config_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZPgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config'
  _KUBERNETESCONFIG.oneofs_by_name['source_oneof']._options = None
  _KUBERNETESCONFIG.oneofs_by_name['source_oneof']._serialized_options = b'\370B\001'
  _KUBERNETESCONFIG.fields_by_name['type']._options = None
  _KUBERNETESCONFIG.fields_by_name['type']._serialized_options = b'\372B\005\202\001\002 \000'
  _KUBERNETESCONFIG.fields_by_name['inlined']._options = None
  _KUBERNETESCONFIG.fields_by_name['inlined']._serialized_options = b'\372B\004r\002\020\001'
  _globals['_LOCALCONFIG']._serialized_start=99
  _globals['_LOCALCONFIG']._serialized_end=202
  _globals['_KUBERNETESCONFIG']._serialized_start=205
  _globals['_KUBERNETESCONFIG']._serialized_end=636
  _globals['_KUBERNETESCONFIG_TYPE']._serialized_start=465
  _globals['_KUBERNETESCONFIG_TYPE']._serialized_end=515
  _globals['_KUBERNETESCONFIG_ENVINJECTIONMODE']._serialized_start=517
  _globals['_KUBERNETESCONFIG_ENVINJECTIONMODE']._serialized_end=615
# @@protoc_insertion_point(module_scope)
