# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/delivery_module/config.proto
"""Generated protocol buffer code."""
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
from prodvana.proto.prodvana.common_config import external_config_pb2 as prodvana_dot_common__config_dot_external__config__pb2
from prodvana.proto.prodvana.common_config import parameters_pb2 as prodvana_dot_common__config_dot_parameters__pb2
from prodvana.proto.prodvana.runtimes import runtimes_config_pb2 as prodvana_dot_runtimes_dot_runtimes__config__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n%prodvana/delivery_module/config.proto\x12\x18prodvana.delivery_module\x1a\x17validate/validate.proto\x1a!prodvana/common_config/task.proto\x1a prodvana/common_config/env.proto\x1a,prodvana/common_config/external_config.proto\x1a\'prodvana/common_config/parameters.proto\x1a\'prodvana/runtimes/runtimes_config.proto\"\xb3\x02\n\x14\x44\x65liveryModuleConfig\x12\x39\n\x04name\x18\x01 \x01(\tB+\xfa\x42(r&\x10\x01\x18?2 ^[a-z]([a-z0-9-]*[a-z0-9]){0,1}$\x12\x39\n\x0btask_config\x18\x02 \x01(\x0b\x32\".prodvana.common_config.TaskConfigH\x00\x12\x41\n\x0f\x65xternal_config\x18\x03 \x01(\x0b\x32&.prodvana.common_config.ExternalConfigH\x00\x12N\n\nparameters\x18\x04 \x03(\x0b\x32+.prodvana.common_config.ParameterDefinitionB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\x42\x12\n\x0b\x65xec_config\x12\x03\xf8\x42\x01\"\x95\x01\n\x1c\x44\x65liveryModuleInstanceConfig\x12\x39\n\x04name\x18\x01 \x01(\tB+\xfa\x42(r&\x10\x01\x18?2 ^[a-z]([a-z0-9-]*[a-z0-9]){0,1}$\x12:\n\nparameters\x18\x03 \x03(\x0b\x32&.prodvana.common_config.ParameterValue\"\xbf\x03\n$CompiledDeliveryModuleInstanceConfig\x12\x42\n\ndefinition\x18\x01 \x01(\x0b\x32..prodvana.delivery_module.DeliveryModuleConfig\x12\x44\n\x11runtime_execution\x18\x02 \x01(\x0b\x32).prodvana.runtimes.RuntimeExecutionConfig\x12}\n\x03\x65nv\x18\x03 \x03(\x0b\x32G.prodvana.delivery_module.CompiledDeliveryModuleInstanceConfig.EnvEntryB\'\xfa\x42$\x9a\x01!\x18\x01\"\x1dr\x1b\x32\x19^[a-zA-Z_]+[a-zA-Z0-9_]*$\x12@\n\x10parameter_values\x18\x04 \x03(\x0b\x32&.prodvana.common_config.ParameterValue\x1aL\n\x08\x45nvEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12/\n\x05value\x18\x02 \x01(\x0b\x32 .prodvana.common_config.EnvValue:\x02\x38\x01\x42TZRgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/delivery_moduleb\x06proto3')



_DELIVERYMODULECONFIG = DESCRIPTOR.message_types_by_name['DeliveryModuleConfig']
_DELIVERYMODULEINSTANCECONFIG = DESCRIPTOR.message_types_by_name['DeliveryModuleInstanceConfig']
_COMPILEDDELIVERYMODULEINSTANCECONFIG = DESCRIPTOR.message_types_by_name['CompiledDeliveryModuleInstanceConfig']
_COMPILEDDELIVERYMODULEINSTANCECONFIG_ENVENTRY = _COMPILEDDELIVERYMODULEINSTANCECONFIG.nested_types_by_name['EnvEntry']
DeliveryModuleConfig = _reflection.GeneratedProtocolMessageType('DeliveryModuleConfig', (_message.Message,), {
  'DESCRIPTOR' : _DELIVERYMODULECONFIG,
  '__module__' : 'prodvana.delivery_module.config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.delivery_module.DeliveryModuleConfig)
  })
_sym_db.RegisterMessage(DeliveryModuleConfig)

DeliveryModuleInstanceConfig = _reflection.GeneratedProtocolMessageType('DeliveryModuleInstanceConfig', (_message.Message,), {
  'DESCRIPTOR' : _DELIVERYMODULEINSTANCECONFIG,
  '__module__' : 'prodvana.delivery_module.config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.delivery_module.DeliveryModuleInstanceConfig)
  })
_sym_db.RegisterMessage(DeliveryModuleInstanceConfig)

CompiledDeliveryModuleInstanceConfig = _reflection.GeneratedProtocolMessageType('CompiledDeliveryModuleInstanceConfig', (_message.Message,), {

  'EnvEntry' : _reflection.GeneratedProtocolMessageType('EnvEntry', (_message.Message,), {
    'DESCRIPTOR' : _COMPILEDDELIVERYMODULEINSTANCECONFIG_ENVENTRY,
    '__module__' : 'prodvana.delivery_module.config_pb2'
    # @@protoc_insertion_point(class_scope:prodvana.delivery_module.CompiledDeliveryModuleInstanceConfig.EnvEntry)
    })
  ,
  'DESCRIPTOR' : _COMPILEDDELIVERYMODULEINSTANCECONFIG,
  '__module__' : 'prodvana.delivery_module.config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.delivery_module.CompiledDeliveryModuleInstanceConfig)
  })
_sym_db.RegisterMessage(CompiledDeliveryModuleInstanceConfig)
_sym_db.RegisterMessage(CompiledDeliveryModuleInstanceConfig.EnvEntry)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZRgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/delivery_module'
  _DELIVERYMODULECONFIG.oneofs_by_name['exec_config']._options = None
  _DELIVERYMODULECONFIG.oneofs_by_name['exec_config']._serialized_options = b'\370B\001'
  _DELIVERYMODULECONFIG.fields_by_name['name']._options = None
  _DELIVERYMODULECONFIG.fields_by_name['name']._serialized_options = b'\372B(r&\020\001\030?2 ^[a-z]([a-z0-9-]*[a-z0-9]){0,1}$'
  _DELIVERYMODULECONFIG.fields_by_name['parameters']._options = None
  _DELIVERYMODULECONFIG.fields_by_name['parameters']._serialized_options = b'\372B\n\222\001\007\"\005\212\001\002\020\001'
  _DELIVERYMODULEINSTANCECONFIG.fields_by_name['name']._options = None
  _DELIVERYMODULEINSTANCECONFIG.fields_by_name['name']._serialized_options = b'\372B(r&\020\001\030?2 ^[a-z]([a-z0-9-]*[a-z0-9]){0,1}$'
  _COMPILEDDELIVERYMODULEINSTANCECONFIG_ENVENTRY._options = None
  _COMPILEDDELIVERYMODULEINSTANCECONFIG_ENVENTRY._serialized_options = b'8\001'
  _COMPILEDDELIVERYMODULEINSTANCECONFIG.fields_by_name['env']._options = None
  _COMPILEDDELIVERYMODULEINSTANCECONFIG.fields_by_name['env']._serialized_options = b'\372B$\232\001!\030\001\"\035r\0332\031^[a-zA-Z_]+[a-zA-Z0-9_]*$'
  _DELIVERYMODULECONFIG._serialized_start=290
  _DELIVERYMODULECONFIG._serialized_end=597
  _DELIVERYMODULEINSTANCECONFIG._serialized_start=600
  _DELIVERYMODULEINSTANCECONFIG._serialized_end=749
  _COMPILEDDELIVERYMODULEINSTANCECONFIG._serialized_start=752
  _COMPILEDDELIVERYMODULEINSTANCECONFIG._serialized_end=1199
  _COMPILEDDELIVERYMODULEINSTANCECONFIG_ENVENTRY._serialized_start=1123
  _COMPILEDDELIVERYMODULEINSTANCECONFIG_ENVENTRY._serialized_end=1199
# @@protoc_insertion_point(module_scope)
