# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/desired_state/debug/debug.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from prodvana.proto.prodvana.desired_state.model import desired_state_pb2 as prodvana_dot_desired__state_dot_model_dot_desired__state__pb2
from prodvana.proto.prodvana.application import application_config_pb2 as prodvana_dot_application_dot_application__config__pb2
from prodvana.proto.prodvana.service import service_config_pb2 as prodvana_dot_service_dot_service__config__pb2
from google.protobuf import any_pb2 as google_dot_protobuf_dot_any__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n(prodvana/desired_state/debug/debug.proto\x12\x1cprodvana.desired_state.debug\x1a\x30prodvana/desired_state/model/desired_state.proto\x1a-prodvana/application/application_config.proto\x1a%prodvana/service/service_config.proto\x1a\x19google/protobuf/any.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xd8\x08\n\x0f\x45ntityDumpState\x12\x34\n\x02id\x18\x01 \x01(\x0b\x32(.prodvana.desired_state.model.Identifier\x12\x41\n\x1dlast_manager_update_timestamp\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12>\n\x0b\x61nnotations\x18\x03 \x01(\x0b\x32).prodvana.desired_state.model.Annotations\x12\x34\n\x06status\x18\x04 \x01(\x0e\x32$.prodvana.desired_state.model.Status\x12\x41\n\rsimple_status\x18\x0b \x01(\x0e\x32*.prodvana.desired_state.model.SimpleStatus\x12;\n\x07\x63ontrol\x18\x05 \x01(\x0b\x32*.prodvana.desired_state.model.ControlState\x12G\n\x08io_state\x18\x06 \x01(\x0b\x32\x35.prodvana.desired_state.debug.EntityDumpState.IOState\x12\x34\n\x04logs\x18\x07 \x03(\x0b\x32&.prodvana.desired_state.model.DebugLog\x12+\n\rdesired_state\x18\x08 \x01(\x0b\x32\x14.google.protobuf.Any\x12+\n\rfetched_state\x18\t \x01(\x0b\x32\x14.google.protobuf.Any\x12*\n\x0ctarget_state\x18\x0c \x01(\x0b\x32\x14.google.protobuf.Any\x12N\n\x0c\x63hild_states\x18\n \x03(\x0b\x32\x38.prodvana.desired_state.debug.EntityDumpState.ChildState\x12\x0e\n\x06\x61\x62sent\x18\r \x01(\x08\x12\x0f\n\x07\x64\x65leted\x18\x0e \x01(\x08\x12\x10\n\x08observer\x18\x0f \x01(\x08\x12\r\n\x05stale\x18\x10 \x01(\x08\x1a\xc4\x01\n\x07IOState\x12:\n\x16last_fetched_timestamp\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12:\n\x16last_applied_timestamp\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x41\n\x1d\x65xpected_next_apply_timestamp\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x1ax\n\nChildState\x12\x34\n\x02id\x18\x01 \x01(\x0b\x32(.prodvana.desired_state.model.Identifier\x12\x34\n\x06status\x18\x02 \x01(\x0e\x32$.prodvana.desired_state.model.Status\"\xda\x02\n\tDumpState\x12?\n\x08\x65ntities\x18\x01 \x03(\x0b\x32-.prodvana.desired_state.debug.EntityDumpState\x12\x43\n\x12\x61pplication_config\x18\x02 \x01(\x0b\x32\'.prodvana.application.ApplicationConfig\x12L\n\x1b\x63ompiled_application_config\x18\x03 \x01(\x0b\x32\'.prodvana.application.ApplicationConfig\x12\x37\n\x0eservice_config\x18\x04 \x01(\x0b\x32\x1f.prodvana.service.ServiceConfig\x12@\n\x17\x63ompiled_service_config\x18\x05 \x01(\x0b\x32\x1f.prodvana.service.ServiceConfigBXZVgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/debugb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'prodvana.desired_state.debug.debug_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZVgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/debug'
  _globals['_ENTITYDUMPSTATE']._serialized_start=271
  _globals['_ENTITYDUMPSTATE']._serialized_end=1383
  _globals['_ENTITYDUMPSTATE_IOSTATE']._serialized_start=1065
  _globals['_ENTITYDUMPSTATE_IOSTATE']._serialized_end=1261
  _globals['_ENTITYDUMPSTATE_CHILDSTATE']._serialized_start=1263
  _globals['_ENTITYDUMPSTATE_CHILDSTATE']._serialized_end=1383
  _globals['_DUMPSTATE']._serialized_start=1386
  _globals['_DUMPSTATE']._serialized_end=1732
# @@protoc_insertion_point(module_scope)
