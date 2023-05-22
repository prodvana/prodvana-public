# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/runtimes/runtimes_config.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from prodvana.proto.validate import validate_pb2 as validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\'prodvana/runtimes/runtimes_config.proto\x12\x11prodvana.runtimes\x1a\x17validate/validate.proto\"\xa6\x02\n\x1eK8SRuntimeInitializationConfig\x12\x13\n\x0b\x61gent_image\x18\x01 \x01(\t\x12\x12\n\nauth_token\x18\x03 \x01(\t\x12\x1c\n\x14use_resource_default\x18\x04 \x01(\x08\x12\"\n\x1ainteraction_server_address\x18\x05 \x01(\t\x12P\n\x08\x65nv_vars\x18\x06 \x03(\x0b\x32>.prodvana.runtimes.K8SRuntimeInitializationConfig.EnvVarsEntry\x1a.\n\x0c\x45nvVarsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01J\x04\x08\x02\x10\x03R\x11\x61piserver_address\"s\n\x1bRuntimeInitializationConfig\x12@\n\x03k8s\x18\x01 \x01(\x0b\x32\x31.prodvana.runtimes.K8SRuntimeInitializationConfigH\x00\x42\x12\n\x10runtime_specific\"\x91\x02\n\x1d\x43ontainerOrchestrationRuntime\x12\x43\n\x03k8s\x18\x01 \x01(\x0b\x32\x34.prodvana.runtimes.ContainerOrchestrationRuntime.K8sH\x00\x12\x43\n\x03\x65\x63s\x18\x02 \x01(\x0b\x32\x34.prodvana.runtimes.ContainerOrchestrationRuntime.ECSH\x00\x1a\x36\n\x03K8s\x12\x1a\n\tnamespace\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x13\n\x0bpvn_managed\x18\x02 \x01(\x08\x1a\x1e\n\x03\x45\x43S\x12\x17\n\x06prefix\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x42\x0e\n\x07\x62\x61\x63kend\x12\x03\xf8\x42\x01\"\xb9\x01\n\x16RuntimeExecutionConfig\x12<\n\x07runtime\x18\x01 \x01(\tB+\xfa\x42(r&\x10\x01\x18?2 ^[a-z]([a-z0-9-]*[a-z0-9]){0,1}$\x12S\n\x17\x63ontainer_orchestration\x18\x02 \x01(\x0b\x32\x30.prodvana.runtimes.ContainerOrchestrationRuntimeH\x00\x42\x0c\n\ntype_oneofBMZKgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/runtimesb\x06proto3')



_K8SRUNTIMEINITIALIZATIONCONFIG = DESCRIPTOR.message_types_by_name['K8SRuntimeInitializationConfig']
_K8SRUNTIMEINITIALIZATIONCONFIG_ENVVARSENTRY = _K8SRUNTIMEINITIALIZATIONCONFIG.nested_types_by_name['EnvVarsEntry']
_RUNTIMEINITIALIZATIONCONFIG = DESCRIPTOR.message_types_by_name['RuntimeInitializationConfig']
_CONTAINERORCHESTRATIONRUNTIME = DESCRIPTOR.message_types_by_name['ContainerOrchestrationRuntime']
_CONTAINERORCHESTRATIONRUNTIME_K8S = _CONTAINERORCHESTRATIONRUNTIME.nested_types_by_name['K8s']
_CONTAINERORCHESTRATIONRUNTIME_ECS = _CONTAINERORCHESTRATIONRUNTIME.nested_types_by_name['ECS']
_RUNTIMEEXECUTIONCONFIG = DESCRIPTOR.message_types_by_name['RuntimeExecutionConfig']
K8SRuntimeInitializationConfig = _reflection.GeneratedProtocolMessageType('K8SRuntimeInitializationConfig', (_message.Message,), {

  'EnvVarsEntry' : _reflection.GeneratedProtocolMessageType('EnvVarsEntry', (_message.Message,), {
    'DESCRIPTOR' : _K8SRUNTIMEINITIALIZATIONCONFIG_ENVVARSENTRY,
    '__module__' : 'prodvana.runtimes.runtimes_config_pb2'
    # @@protoc_insertion_point(class_scope:prodvana.runtimes.K8SRuntimeInitializationConfig.EnvVarsEntry)
    })
  ,
  'DESCRIPTOR' : _K8SRUNTIMEINITIALIZATIONCONFIG,
  '__module__' : 'prodvana.runtimes.runtimes_config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.runtimes.K8SRuntimeInitializationConfig)
  })
_sym_db.RegisterMessage(K8SRuntimeInitializationConfig)
_sym_db.RegisterMessage(K8SRuntimeInitializationConfig.EnvVarsEntry)

RuntimeInitializationConfig = _reflection.GeneratedProtocolMessageType('RuntimeInitializationConfig', (_message.Message,), {
  'DESCRIPTOR' : _RUNTIMEINITIALIZATIONCONFIG,
  '__module__' : 'prodvana.runtimes.runtimes_config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.runtimes.RuntimeInitializationConfig)
  })
_sym_db.RegisterMessage(RuntimeInitializationConfig)

ContainerOrchestrationRuntime = _reflection.GeneratedProtocolMessageType('ContainerOrchestrationRuntime', (_message.Message,), {

  'K8s' : _reflection.GeneratedProtocolMessageType('K8s', (_message.Message,), {
    'DESCRIPTOR' : _CONTAINERORCHESTRATIONRUNTIME_K8S,
    '__module__' : 'prodvana.runtimes.runtimes_config_pb2'
    # @@protoc_insertion_point(class_scope:prodvana.runtimes.ContainerOrchestrationRuntime.K8s)
    })
  ,

  'ECS' : _reflection.GeneratedProtocolMessageType('ECS', (_message.Message,), {
    'DESCRIPTOR' : _CONTAINERORCHESTRATIONRUNTIME_ECS,
    '__module__' : 'prodvana.runtimes.runtimes_config_pb2'
    # @@protoc_insertion_point(class_scope:prodvana.runtimes.ContainerOrchestrationRuntime.ECS)
    })
  ,
  'DESCRIPTOR' : _CONTAINERORCHESTRATIONRUNTIME,
  '__module__' : 'prodvana.runtimes.runtimes_config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.runtimes.ContainerOrchestrationRuntime)
  })
_sym_db.RegisterMessage(ContainerOrchestrationRuntime)
_sym_db.RegisterMessage(ContainerOrchestrationRuntime.K8s)
_sym_db.RegisterMessage(ContainerOrchestrationRuntime.ECS)

RuntimeExecutionConfig = _reflection.GeneratedProtocolMessageType('RuntimeExecutionConfig', (_message.Message,), {
  'DESCRIPTOR' : _RUNTIMEEXECUTIONCONFIG,
  '__module__' : 'prodvana.runtimes.runtimes_config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.runtimes.RuntimeExecutionConfig)
  })
_sym_db.RegisterMessage(RuntimeExecutionConfig)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZKgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/runtimes'
  _K8SRUNTIMEINITIALIZATIONCONFIG_ENVVARSENTRY._options = None
  _K8SRUNTIMEINITIALIZATIONCONFIG_ENVVARSENTRY._serialized_options = b'8\001'
  _CONTAINERORCHESTRATIONRUNTIME_K8S.fields_by_name['namespace']._options = None
  _CONTAINERORCHESTRATIONRUNTIME_K8S.fields_by_name['namespace']._serialized_options = b'\372B\004r\002\020\001'
  _CONTAINERORCHESTRATIONRUNTIME_ECS.fields_by_name['prefix']._options = None
  _CONTAINERORCHESTRATIONRUNTIME_ECS.fields_by_name['prefix']._serialized_options = b'\372B\004r\002\020\001'
  _CONTAINERORCHESTRATIONRUNTIME.oneofs_by_name['backend']._options = None
  _CONTAINERORCHESTRATIONRUNTIME.oneofs_by_name['backend']._serialized_options = b'\370B\001'
  _RUNTIMEEXECUTIONCONFIG.fields_by_name['runtime']._options = None
  _RUNTIMEEXECUTIONCONFIG.fields_by_name['runtime']._serialized_options = b'\372B(r&\020\001\030?2 ^[a-z]([a-z0-9-]*[a-z0-9]){0,1}$'
  _K8SRUNTIMEINITIALIZATIONCONFIG._serialized_start=88
  _K8SRUNTIMEINITIALIZATIONCONFIG._serialized_end=382
  _K8SRUNTIMEINITIALIZATIONCONFIG_ENVVARSENTRY._serialized_start=311
  _K8SRUNTIMEINITIALIZATIONCONFIG_ENVVARSENTRY._serialized_end=357
  _RUNTIMEINITIALIZATIONCONFIG._serialized_start=384
  _RUNTIMEINITIALIZATIONCONFIG._serialized_end=499
  _CONTAINERORCHESTRATIONRUNTIME._serialized_start=502
  _CONTAINERORCHESTRATIONRUNTIME._serialized_end=775
  _CONTAINERORCHESTRATIONRUNTIME_K8S._serialized_start=673
  _CONTAINERORCHESTRATIONRUNTIME_K8S._serialized_end=727
  _CONTAINERORCHESTRATIONRUNTIME_ECS._serialized_start=729
  _CONTAINERORCHESTRATIONRUNTIME_ECS._serialized_end=759
  _RUNTIMEEXECUTIONCONFIG._serialized_start=778
  _RUNTIMEEXECUTIONCONFIG._serialized_end=963
# @@protoc_insertion_point(module_scope)