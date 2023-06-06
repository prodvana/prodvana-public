# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/delivery_extension/config.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from prodvana.proto.validate import validate_pb2 as validate_dot_validate__pb2
from prodvana.proto.prodvana.common_config import task_pb2 as prodvana_dot_common__config_dot_task__pb2
from prodvana.proto.prodvana.common_config import env_pb2 as prodvana_dot_common__config_dot_env__pb2
from prodvana.proto.prodvana.common_config import kubernetes_config_pb2 as prodvana_dot_common__config_dot_kubernetes__config__pb2
from prodvana.proto.prodvana.common_config import parameters_pb2 as prodvana_dot_common__config_dot_parameters__pb2
from prodvana.proto.prodvana.runtimes import runtimes_config_pb2 as prodvana_dot_runtimes_dot_runtimes__config__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n(prodvana/delivery_extension/config.proto\x12\x1bprodvana.delivery_extension\x1a\x17validate/validate.proto\x1a!prodvana/common_config/task.proto\x1a prodvana/common_config/env.proto\x1a.prodvana/common_config/kubernetes_config.proto\x1a\'prodvana/common_config/parameters.proto\x1a\'prodvana/runtimes/runtimes_config.proto\"\xbb\x02\n\x17\x44\x65liveryExtensionConfig\x12:\n\x04name\x18\x01 \x01(\tB,\xfa\x42)r\'\x10\x00\x18?2!^[a-z]?([a-z0-9-]*[a-z0-9]){0,1}$\x12\x39\n\x0btask_config\x18\x02 \x01(\x0b\x32\".prodvana.common_config.TaskConfigH\x00\x12\x45\n\x11kubernetes_config\x18\x03 \x01(\x0b\x32(.prodvana.common_config.KubernetesConfigH\x00\x12N\n\nparameters\x18\x04 \x03(\x0b\x32+.prodvana.common_config.ParameterDefinitionB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\x42\x12\n\x0b\x65xec_config\x12\x03\xf8\x42\x01\"\x98\x01\n\x1f\x44\x65liveryExtensionInstanceConfig\x12\x39\n\x04name\x18\x01 \x01(\tB+\xfa\x42(r&\x10\x01\x18?2 ^[a-z]([a-z0-9-]*[a-z0-9]){0,1}$\x12:\n\nparameters\x18\x03 \x03(\x0b\x32&.prodvana.common_config.ParameterValue\"\xcf\x03\n\'CompiledDeliveryExtensionInstanceConfig\x12H\n\ndefinition\x18\x01 \x01(\x0b\x32\x34.prodvana.delivery_extension.DeliveryExtensionConfig\x12\x44\n\x11runtime_execution\x18\x02 \x01(\x0b\x32).prodvana.runtimes.RuntimeExecutionConfig\x12\x83\x01\n\x03\x65nv\x18\x03 \x03(\x0b\x32M.prodvana.delivery_extension.CompiledDeliveryExtensionInstanceConfig.EnvEntryB\'\xfa\x42$\x9a\x01!\x18\x01\"\x1dr\x1b\x32\x19^[a-zA-Z_]+[a-zA-Z0-9_]*$\x12@\n\x10parameter_values\x18\x04 \x03(\x0b\x32&.prodvana.common_config.ParameterValue\x1aL\n\x08\x45nvEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12/\n\x05value\x18\x02 \x01(\x0b\x32 .prodvana.common_config.EnvValue:\x02\x38\x01*@\n\x04Type\x12\x10\n\x0cUNKNOWN_TYPE\x10\x00\x12\x17\n\x13GLOBAL_USER_CREATED\x10\x01\x12\r\n\tEPHEMERAL\x10\x02\x42WZUgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/delivery_extensionb\x06proto3')

_TYPE = DESCRIPTOR.enum_types_by_name['Type']
Type = enum_type_wrapper.EnumTypeWrapper(_TYPE)
UNKNOWN_TYPE = 0
GLOBAL_USER_CREATED = 1
EPHEMERAL = 2


_DELIVERYEXTENSIONCONFIG = DESCRIPTOR.message_types_by_name['DeliveryExtensionConfig']
_DELIVERYEXTENSIONINSTANCECONFIG = DESCRIPTOR.message_types_by_name['DeliveryExtensionInstanceConfig']
_COMPILEDDELIVERYEXTENSIONINSTANCECONFIG = DESCRIPTOR.message_types_by_name['CompiledDeliveryExtensionInstanceConfig']
_COMPILEDDELIVERYEXTENSIONINSTANCECONFIG_ENVENTRY = _COMPILEDDELIVERYEXTENSIONINSTANCECONFIG.nested_types_by_name['EnvEntry']
DeliveryExtensionConfig = _reflection.GeneratedProtocolMessageType('DeliveryExtensionConfig', (_message.Message,), {
  'DESCRIPTOR' : _DELIVERYEXTENSIONCONFIG,
  '__module__' : 'prodvana.delivery_extension.config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.delivery_extension.DeliveryExtensionConfig)
  })
_sym_db.RegisterMessage(DeliveryExtensionConfig)

DeliveryExtensionInstanceConfig = _reflection.GeneratedProtocolMessageType('DeliveryExtensionInstanceConfig', (_message.Message,), {
  'DESCRIPTOR' : _DELIVERYEXTENSIONINSTANCECONFIG,
  '__module__' : 'prodvana.delivery_extension.config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.delivery_extension.DeliveryExtensionInstanceConfig)
  })
_sym_db.RegisterMessage(DeliveryExtensionInstanceConfig)

CompiledDeliveryExtensionInstanceConfig = _reflection.GeneratedProtocolMessageType('CompiledDeliveryExtensionInstanceConfig', (_message.Message,), {

  'EnvEntry' : _reflection.GeneratedProtocolMessageType('EnvEntry', (_message.Message,), {
    'DESCRIPTOR' : _COMPILEDDELIVERYEXTENSIONINSTANCECONFIG_ENVENTRY,
    '__module__' : 'prodvana.delivery_extension.config_pb2'
    # @@protoc_insertion_point(class_scope:prodvana.delivery_extension.CompiledDeliveryExtensionInstanceConfig.EnvEntry)
    })
  ,
  'DESCRIPTOR' : _COMPILEDDELIVERYEXTENSIONINSTANCECONFIG,
  '__module__' : 'prodvana.delivery_extension.config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.delivery_extension.CompiledDeliveryExtensionInstanceConfig)
  })
_sym_db.RegisterMessage(CompiledDeliveryExtensionInstanceConfig)
_sym_db.RegisterMessage(CompiledDeliveryExtensionInstanceConfig.EnvEntry)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZUgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/delivery_extension'
  _DELIVERYEXTENSIONCONFIG.oneofs_by_name['exec_config']._options = None
  _DELIVERYEXTENSIONCONFIG.oneofs_by_name['exec_config']._serialized_options = b'\370B\001'
  _DELIVERYEXTENSIONCONFIG.fields_by_name['name']._options = None
  _DELIVERYEXTENSIONCONFIG.fields_by_name['name']._serialized_options = b'\372B)r\'\020\000\030?2!^[a-z]?([a-z0-9-]*[a-z0-9]){0,1}$'
  _DELIVERYEXTENSIONCONFIG.fields_by_name['parameters']._options = None
  _DELIVERYEXTENSIONCONFIG.fields_by_name['parameters']._serialized_options = b'\372B\n\222\001\007\"\005\212\001\002\020\001'
  _DELIVERYEXTENSIONINSTANCECONFIG.fields_by_name['name']._options = None
  _DELIVERYEXTENSIONINSTANCECONFIG.fields_by_name['name']._serialized_options = b'\372B(r&\020\001\030?2 ^[a-z]([a-z0-9-]*[a-z0-9]){0,1}$'
  _COMPILEDDELIVERYEXTENSIONINSTANCECONFIG_ENVENTRY._options = None
  _COMPILEDDELIVERYEXTENSIONINSTANCECONFIG_ENVENTRY._serialized_options = b'8\001'
  _COMPILEDDELIVERYEXTENSIONINSTANCECONFIG.fields_by_name['env']._options = None
  _COMPILEDDELIVERYEXTENSIONINSTANCECONFIG.fields_by_name['env']._serialized_options = b'\372B$\232\001!\030\001\"\035r\0332\031^[a-zA-Z_]+[a-zA-Z0-9_]*$'
  _TYPE._serialized_start=1236
  _TYPE._serialized_end=1300
  _DELIVERYEXTENSIONCONFIG._serialized_start=298
  _DELIVERYEXTENSIONCONFIG._serialized_end=613
  _DELIVERYEXTENSIONINSTANCECONFIG._serialized_start=616
  _DELIVERYEXTENSIONINSTANCECONFIG._serialized_end=768
  _COMPILEDDELIVERYEXTENSIONINSTANCECONFIG._serialized_start=771
  _COMPILEDDELIVERYEXTENSIONINSTANCECONFIG._serialized_end=1234
  _COMPILEDDELIVERYEXTENSIONINSTANCECONFIG_ENVENTRY._serialized_start=1158
  _COMPILEDDELIVERYEXTENSIONINSTANCECONFIG_ENVENTRY._serialized_end=1234
# @@protoc_insertion_point(module_scope)
