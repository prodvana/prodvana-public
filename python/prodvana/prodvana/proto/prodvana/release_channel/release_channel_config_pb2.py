# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/release_channel/release_channel_config.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from prodvana.proto.prodvana.common_config import env_pb2 as prodvana_dot_common__config_dot_env__pb2
from prodvana.proto.prodvana.common_config import maturity_pb2 as prodvana_dot_common__config_dot_maturity__pb2
from prodvana.proto.prodvana.pipelines import pipelines_pb2 as prodvana_dot_pipelines_dot_pipelines__pb2
from prodvana.proto.prodvana.protection import attachments_pb2 as prodvana_dot_protection_dot_attachments__pb2
from prodvana.proto.prodvana.protection import protection_reference_pb2 as prodvana_dot_protection_dot_protection__reference__pb2
from prodvana.proto.prodvana.runtimes import runtimes_config_pb2 as prodvana_dot_runtimes_dot_runtimes__config__pb2
from prodvana.proto.prodvana.workflow import integration_config_pb2 as prodvana_dot_workflow_dot_integration__config__pb2
from prodvana.proto.validate import validate_pb2 as validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n5prodvana/release_channel/release_channel_config.proto\x12\x18prodvana.release_channel\x1a\x1egoogle/protobuf/duration.proto\x1a prodvana/common_config/env.proto\x1a%prodvana/common_config/maturity.proto\x1a\"prodvana/pipelines/pipelines.proto\x1a%prodvana/protection/attachments.proto\x1a.prodvana/protection/protection_reference.proto\x1a\'prodvana/runtimes/runtimes_config.proto\x1a*prodvana/workflow/integration_config.proto\x1a\x17validate/validate.proto\"\xcd\x01\n\x06Policy\x12n\n\x0b\x64\x65\x66\x61ult_env\x18\x01 \x03(\x0b\x32\x30.prodvana.release_channel.Policy.DefaultEnvEntryB\'\xfa\x42$\x9a\x01!\x18\x01\"\x1dr\x1b\x32\x19^[a-zA-Z_]+[a-zA-Z0-9_]*$\x1aS\n\x0f\x44\x65\x66\x61ultEnvEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12/\n\x05value\x18\x02 \x01(\x0b\x32 .prodvana.common_config.EnvValue:\x02\x38\x01\"\xba\x04\n\x14ReleaseChannelConfig\x12\x39\n\x04name\x18\x01 \x01(\tB+\xfa\x42(r&\x10\x01\x18?2 ^[a-z]([a-z0-9-]*[a-z0-9]){0,1}$\x12\r\n\x05order\x18\x02 \x01(\x03\x12\x32\n\x08maturity\x18\x03 \x01(\x0e\x32 .prodvana.common_config.Maturity\x12\x30\n\x06policy\x18\x04 \x01(\x0b\x32 .prodvana.release_channel.Policy\x12G\n\x08runtimes\x18\x05 \x03(\x0b\x32\x35.prodvana.release_channel.ReleaseChannelRuntimeConfig\x12@\n\x12\x64\x65ploy_annotations\x18\x06 \x01(\x0b\x32$.prodvana.workflow.AnnotationsConfig\x12=\n\rpreconditions\x18\x07 \x03(\x0b\x32&.prodvana.release_channel.Precondition\x12Q\n\x0bprotections\x18\x08 \x03(\x0b\x32<.prodvana.release_channel.ProtectionReleaseChannelAttachment\x12U\n\x17\x63onvergence_protections\x18\t \x03(\x0b\x32\x34.prodvana.protection.ProtectionConvergenceAttachment\"\xed\x01\n\"ProtectionReleaseChannelAttachment\x12:\n\x04name\x18\x01 \x01(\tB,\xfa\x42)r\'\x10\x00\x18?2!^[a-z]?([a-z0-9-]*[a-z0-9]){0,1}$\x12?\n\x03ref\x18\x02 \x01(\x0b\x32(.prodvana.protection.ProtectionReferenceB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12J\n\tlifecycle\x18\x03 \x03(\x0b\x32(.prodvana.protection.ProtectionLifecycleB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\"\x94\x04\n\x0cPrecondition\x12]\n\x16release_channel_stable\x18\x01 \x01(\x0b\x32;.prodvana.release_channel.Precondition.ReleaseChannelStableH\x00\x12P\n\x0fmanual_approval\x18\x02 \x01(\x0b\x32\x35.prodvana.release_channel.Precondition.ManualApprovalH\x00\x12H\n\x0b\x63ustom_task\x18\x03 \x01(\x0b\x32\x31.prodvana.release_channel.Precondition.CustomTaskH\x00\x1a\\\n\x14ReleaseChannelStable\x12\x17\n\x0frelease_channel\x18\x01 \x01(\t\x12+\n\x08\x64uration\x18\x02 \x01(\x0b\x32\x19.google.protobuf.Duration\x1a\x33\n\x0eManualApproval\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x02 \x01(\t\x1aT\n\nCustomTask\x12\x11\n\ttask_name\x18\x01 \x01(\t\x12\x33\n\x0b\x63ustom_task\x18\x02 \x01(\x0b\x32\x1e.prodvana.pipelines.CustomTaskB\x0e\n\x0cpreconditionJ\x04\x08\x04\x10\x05R\nprotection\"\xb9\x02\n\x1bReleaseChannelRuntimeConfig\x12<\n\x07runtime\x18\x02 \x01(\tB+\xfa\x42(r&\x10\x01\x18?2 ^[a-z]([a-z0-9-]*[a-z0-9]){0,1}$\x12:\n\x04name\x18\x03 \x01(\tB,\xfa\x42)r\'\x10\x00\x18?2!^[a-z]?([a-z0-9-]*[a-z0-9]){0,1}$\x12S\n\x17\x63ontainer_orchestration\x18\x01 \x01(\x0b\x32\x30.prodvana.runtimes.ContainerOrchestrationRuntimeH\x00\x12=\n\x04type\x18\x04 \x01(\x0e\x32/.prodvana.release_channel.RuntimeConnectionTypeB\x0c\n\ncapability*V\n\x15RuntimeConnectionType\x12\x16\n\x12UNKNOWN_CONNECTION\x10\x00\x12\x16\n\x12LONG_LIVED_COMPUTE\x10\x01\x12\r\n\tEXTENSION\x10\x02\x42TZRgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/release_channelb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'prodvana.release_channel.release_channel_config_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZRgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/release_channel'
  _POLICY_DEFAULTENVENTRY._options = None
  _POLICY_DEFAULTENVENTRY._serialized_options = b'8\001'
  _POLICY.fields_by_name['default_env']._options = None
  _POLICY.fields_by_name['default_env']._serialized_options = b'\372B$\232\001!\030\001\"\035r\0332\031^[a-zA-Z_]+[a-zA-Z0-9_]*$'
  _RELEASECHANNELCONFIG.fields_by_name['name']._options = None
  _RELEASECHANNELCONFIG.fields_by_name['name']._serialized_options = b'\372B(r&\020\001\030?2 ^[a-z]([a-z0-9-]*[a-z0-9]){0,1}$'
  _PROTECTIONRELEASECHANNELATTACHMENT.fields_by_name['name']._options = None
  _PROTECTIONRELEASECHANNELATTACHMENT.fields_by_name['name']._serialized_options = b'\372B)r\'\020\000\030?2!^[a-z]?([a-z0-9-]*[a-z0-9]){0,1}$'
  _PROTECTIONRELEASECHANNELATTACHMENT.fields_by_name['ref']._options = None
  _PROTECTIONRELEASECHANNELATTACHMENT.fields_by_name['ref']._serialized_options = b'\372B\005\212\001\002\020\001'
  _PROTECTIONRELEASECHANNELATTACHMENT.fields_by_name['lifecycle']._options = None
  _PROTECTIONRELEASECHANNELATTACHMENT.fields_by_name['lifecycle']._serialized_options = b'\372B\n\222\001\007\"\005\212\001\002\020\001'
  _RELEASECHANNELRUNTIMECONFIG.fields_by_name['runtime']._options = None
  _RELEASECHANNELRUNTIMECONFIG.fields_by_name['runtime']._serialized_options = b'\372B(r&\020\001\030?2 ^[a-z]([a-z0-9-]*[a-z0-9]){0,1}$'
  _RELEASECHANNELRUNTIMECONFIG.fields_by_name['name']._options = None
  _RELEASECHANNELRUNTIMECONFIG.fields_by_name['name']._serialized_options = b'\372B)r\'\020\000\030?2!^[a-z]?([a-z0-9-]*[a-z0-9]){0,1}$'
  _globals['_RUNTIMECONNECTIONTYPE']._serialized_start=2293
  _globals['_RUNTIMECONNECTIONTYPE']._serialized_end=2379
  _globals['_POLICY']._serialized_start=422
  _globals['_POLICY']._serialized_end=627
  _globals['_POLICY_DEFAULTENVENTRY']._serialized_start=544
  _globals['_POLICY_DEFAULTENVENTRY']._serialized_end=627
  _globals['_RELEASECHANNELCONFIG']._serialized_start=630
  _globals['_RELEASECHANNELCONFIG']._serialized_end=1200
  _globals['_PROTECTIONRELEASECHANNELATTACHMENT']._serialized_start=1203
  _globals['_PROTECTIONRELEASECHANNELATTACHMENT']._serialized_end=1440
  _globals['_PRECONDITION']._serialized_start=1443
  _globals['_PRECONDITION']._serialized_end=1975
  _globals['_PRECONDITION_RELEASECHANNELSTABLE']._serialized_start=1710
  _globals['_PRECONDITION_RELEASECHANNELSTABLE']._serialized_end=1802
  _globals['_PRECONDITION_MANUALAPPROVAL']._serialized_start=1804
  _globals['_PRECONDITION_MANUALAPPROVAL']._serialized_end=1855
  _globals['_PRECONDITION_CUSTOMTASK']._serialized_start=1857
  _globals['_PRECONDITION_CUSTOMTASK']._serialized_end=1941
  _globals['_RELEASECHANNELRUNTIMECONFIG']._serialized_start=1978
  _globals['_RELEASECHANNELRUNTIMECONFIG']._serialized_end=2291
# @@protoc_insertion_point(module_scope)
