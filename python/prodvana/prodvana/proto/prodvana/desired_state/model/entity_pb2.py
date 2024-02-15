# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/desired_state/model/entity.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from prodvana.proto.prodvana.common_config import task_pb2 as prodvana_dot_common__config_dot_task__pb2
from prodvana.proto.prodvana.desired_state.model import desired_state_pb2 as prodvana_dot_desired__state_dot_model_dot_desired__state__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n)prodvana/desired_state/model/entity.proto\x12\x1cprodvana.desired_state.model\x1a!prodvana/common_config/task.proto\x1a\x30prodvana/desired_state/model/desired_state.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\x9f\x08\n\rNotifications\x12U\n\x10program_failures\x18\x01 \x01(\x0b\x32;.prodvana.desired_state.model.Notifications.ProgramFailures\x12T\n\x10runtime_failures\x18\x02 \x03(\x0b\x32:.prodvana.desired_state.model.Notifications.RuntimeFailure\x12Y\n\x12protection_failure\x18\x03 \x03(\x0b\x32=.prodvana.desired_state.model.Notifications.ProtectionFailure\x12n\n\x1d\x63onvergence_extension_failure\x18\x04 \x03(\x0b\x32G.prodvana.desired_state.model.Notifications.ConvergenceExtensionFailure\x12[\n\x13\x64\x65layed_convergence\x18\x05 \x01(\x0b\x32>.prodvana.desired_state.model.Notifications.DelayedConvergence\x1a\x61\n\x0fProgramFailures\x12\x15\n\rfailure_count\x18\x01 \x01(\x05\x12\x37\n\x13most_recent_failure\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x1a\x8a\x02\n\x0eRuntimeFailure\x12\\\n\x0c\x66\x61ilure_type\x18\x01 \x01(\x0e\x32\x46.prodvana.desired_state.model.Notifications.RuntimeFailure.FailureType\"\x99\x01\n\x0b\x46\x61ilureType\x12\x0b\n\x07UNKNOWN\x10\x00\x12 \n\x1c\x45XTENSION_FETCH_INVOKE_ERROR\x10\x01\x12 \n\x1c\x45XTENSION_FETCH_RESULT_ERROR\x10\x02\x12\x17\n\x13RUNTIME_APPLY_ERROR\x10\x03\x12 \n\x1c\x45XTENSION_APPLY_INVOKE_ERROR\x10\x04\x1aT\n\x11ProtectionFailure\x12?\n\rprotection_id\x18\x01 \x01(\x0b\x32(.prodvana.desired_state.model.Identifier\x1a]\n\x1b\x43onvergenceExtensionFailure\x12>\n\x0c\x65xtension_id\x18\x01 \x01(\x0b\x32(.prodvana.desired_state.model.Identifier\x1a\x14\n\x12\x44\x65layedConvergence\"\xb9\n\n\x06\x45ntity\x12\x34\n\x02id\x18\x01 \x01(\x0b\x32(.prodvana.desired_state.model.Identifier\x12\x18\n\x10\x64\x65sired_state_id\x18\x02 \x01(\t\x12\x1d\n\x15root_desired_state_id\x18\x03 \x01(\t\x12>\n\x0b\x61nnotations\x18\x04 \x01(\x0b\x32).prodvana.desired_state.model.Annotations\x12\x34\n\x06status\x18\x05 \x01(\x0e\x32$.prodvana.desired_state.model.Status\x12\x41\n\rsimple_status\x18\x11 \x01(\x0e\x32*.prodvana.desired_state.model.SimpleStatus\x12;\n\x0estarting_state\x18\x06 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12<\n\x0flast_seen_state\x18\x07 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12:\n\rdesired_state\x18\x08 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x39\n\x0ctarget_state\x18\x13 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12K\n\x15precondition_statuses\x18\t \x03(\x0b\x32,.prodvana.desired_state.model.ConditionState\x12L\n\x13status_explanations\x18\n \x03(\x0b\x32/.prodvana.desired_state.model.StatusExplanation\x12\x34\n\x04logs\x18\x0b \x03(\x0b\x32&.prodvana.desired_state.model.DebugLog\x12K\n\x12\x61\x63tion_explanation\x18\x0c \x01(\x0b\x32/.prodvana.desired_state.model.ActionExplanation\x12\x39\n\x15last_update_timestamp\x18\r \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12:\n\x16last_fetched_timestamp\x18\x0e \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12:\n\x16last_applied_timestamp\x18\x0f \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12>\n\x0c\x64\x65pendencies\x18\x10 \x03(\x0b\x32(.prodvana.desired_state.model.Identifier\x12\x38\n\tlifecycle\x18\x12 \x01(\x0e\x32%.prodvana.common_config.TaskLifecycle\x12G\n\x10missing_approval\x18\x14 \x01(\x0b\x32-.prodvana.desired_state.model.MissingApproval\x12=\n\x0b\x61pply_error\x18\x15 \x01(\x0b\x32(.prodvana.desired_state.model.ApplyError\x12\x42\n\rnotifications\x18\x16 \x01(\x0b\x32+.prodvana.desired_state.model.Notifications\"}\n\x0b\x45ntityGraph\x12\x36\n\x04root\x18\x01 \x01(\x0b\x32(.prodvana.desired_state.model.Identifier\x12\x36\n\x08\x65ntities\x18\x02 \x03(\x0b\x32$.prodvana.desired_state.model.EntityBXZVgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/modelb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'prodvana.desired_state.model.entity_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZVgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/model'
  _globals['_NOTIFICATIONS']._serialized_start=194
  _globals['_NOTIFICATIONS']._serialized_end=1249
  _globals['_NOTIFICATIONS_PROGRAMFAILURES']._serialized_start=680
  _globals['_NOTIFICATIONS_PROGRAMFAILURES']._serialized_end=777
  _globals['_NOTIFICATIONS_RUNTIMEFAILURE']._serialized_start=780
  _globals['_NOTIFICATIONS_RUNTIMEFAILURE']._serialized_end=1046
  _globals['_NOTIFICATIONS_RUNTIMEFAILURE_FAILURETYPE']._serialized_start=893
  _globals['_NOTIFICATIONS_RUNTIMEFAILURE_FAILURETYPE']._serialized_end=1046
  _globals['_NOTIFICATIONS_PROTECTIONFAILURE']._serialized_start=1048
  _globals['_NOTIFICATIONS_PROTECTIONFAILURE']._serialized_end=1132
  _globals['_NOTIFICATIONS_CONVERGENCEEXTENSIONFAILURE']._serialized_start=1134
  _globals['_NOTIFICATIONS_CONVERGENCEEXTENSIONFAILURE']._serialized_end=1227
  _globals['_NOTIFICATIONS_DELAYEDCONVERGENCE']._serialized_start=1229
  _globals['_NOTIFICATIONS_DELAYEDCONVERGENCE']._serialized_end=1249
  _globals['_ENTITY']._serialized_start=1252
  _globals['_ENTITY']._serialized_end=2589
  _globals['_ENTITYGRAPH']._serialized_start=2591
  _globals['_ENTITYGRAPH']._serialized_end=2716
# @@protoc_insertion_point(module_scope)
