# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/desired_state/model/desired_state.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from prodvana.proto.validate import validate_pb2 as validate_dot_validate__pb2
from prodvana.proto.prodvana.common_config import parameters_pb2 as prodvana_dot_common__config_dot_parameters__pb2
from prodvana.proto.prodvana.common_config import program_pb2 as prodvana_dot_common__config_dot_program__pb2
from prodvana.proto.prodvana.common_config import retry_pb2 as prodvana_dot_common__config_dot_retry__pb2
from prodvana.proto.prodvana.common_config import task_pb2 as prodvana_dot_common__config_dot_task__pb2
from prodvana.proto.prodvana.environment import clusters_pb2 as prodvana_dot_environment_dot_clusters__pb2
from prodvana.proto.prodvana.protection import protection_reference_pb2 as prodvana_dot_protection_dot_protection__reference__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n0prodvana/desired_state/model/desired_state.proto\x12\x1cprodvana.desired_state.model\x1a\x1egoogle/protobuf/duration.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x17validate/validate.proto\x1a\'prodvana/common_config/parameters.proto\x1a$prodvana/common_config/program.proto\x1a\"prodvana/common_config/retry.proto\x1a!prodvana/common_config/task.proto\x1a#prodvana/environment/clusters.proto\x1a.prodvana/protection/protection_reference.proto\"p\n\x0eProtectionLink\x12;\n\tlifecycle\x18\x03 \x01(\x0b\x32(.prodvana.protection.ProtectionLifecycle\x12\x15\n\rattachment_id\x18\x04 \x01(\tJ\x04\x08\x01\x10\x02J\x04\x08\x02\x10\x03\"\xa8\x06\n\tCondition\x12X\n\x07rc_cond\x18\x01 \x01(\x0b\x32\x45.prodvana.desired_state.model.Condition.ReleaseChannelStableConditionH\x00\x12Q\n\x0fmanual_approval\x18\x02 \x01(\x0b\x32\x36.prodvana.desired_state.model.Condition.ManualApprovalH\x00\x12\\\n\x0b\x63ustom_task\x18\x03 \x01(\x0b\x32\x45.prodvana.desired_state.model.Condition.CustomTaskSuccessfulConditionH\x00\x12\x18\n\x10\x64\x65sired_state_id\x18\x04 \x01(\t\x1a\xa7\x01\n\x1dReleaseChannelStableCondition\x12\x13\n\x0b\x61pplication\x18\x01 \x01(\t\x12\x0f\n\x07service\x18\x02 \x01(\t\x12\x12\n\nservice_id\x18\x03 \x01(\t\x12\x17\n\x0frelease_channel\x18\x04 \x01(\t\x12\x1a\n\x12release_channel_id\x18\x05 \x01(\t\x12\x17\n\x0fservice_version\x18\x06 \x01(\t\x1a\x34\n\x0eManualApproval\x12\r\n\x05topic\x18\x01 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x02 \x01(\t\x1a\x88\x02\n\x1d\x43ustomTaskSuccessfulCondition\x12\x18\n\x10\x63ustom_task_name\x18\x01 \x01(\t\x12\x66\n\nprotection\x18\x02 \x01(\x0b\x32P.prodvana.desired_state.model.Condition.CustomTaskSuccessfulCondition.ProtectionH\x00\x1a[\n\nProtection\x12\x0c\n\x04name\x18\x01 \x01(\t\x12?\n\ttask_type\x18\x03 \x01(\x0e\x32,.prodvana.desired_state.model.CustomTaskTypeB\x08\n\x06sourceB\x0b\n\tcondition\"\xa0\x01\n\x11\x44\x65liveryExtension\x12\x13\n\x0binstance_id\x18\x01 \x01(\t\x12\x38\n\tlifecycle\x18\x02 \x01(\x0e\x32%.prodvana.common_config.TaskLifecycle\x12<\n\nreferences\x18\x03 \x03(\x0b\x32(.prodvana.desired_state.model.Identifier\"L\n\nIdentifier\x12\x30\n\x04type\x18\x01 \x01(\x0e\x32\".prodvana.desired_state.model.Type\x12\x0c\n\x04name\x18\x02 \x01(\t\"\xcb\x04\n\x08Metadata\x12>\n\rpreconditions\x18\x01 \x03(\x0b\x32\'.prodvana.desired_state.model.Condition\x12;\n\ninvariants\x18\x02 \x03(\x0b\x32\'.prodvana.desired_state.model.Condition\x12\x36\n\x04self\x18\x03 \x01(\x0b\x32(.prodvana.desired_state.model.Identifier\x12\x18\n\x10\x64\x65sired_state_id\x18\x04 \x01(\t\x12\x1d\n\x15root_desired_state_id\x18\x05 \x01(\t\x12\x46\n\x10protection_links\x18\x06 \x03(\x0b\x32,.prodvana.desired_state.model.ProtectionLink\x12L\n\x13\x64\x65livery_extensions\x18\x08 \x03(\x0b\x32/.prodvana.desired_state.model.DeliveryExtension\x12\"\n\x1atarget_state_set_by_parent\x18\t \x01(\x08\x12%\n\x1drequire_approval_before_apply\x18\n \x01(\x08\x12 \n\x18\x61pplies_in_observer_mode\x18\x0b \x01(\x08\x12;\n\x18\x63onvergence_grace_period\x18\x0c \x01(\x0b\x32\x19.google.protobuf.DurationJ\x04\x08\x07\x10\x08R\x0bprotections\"\xad\x01\n\x11StatusExplanation\x12\x39\n\x07subject\x18\x01 \x01(\x0b\x32(.prodvana.desired_state.model.Identifier\x12:\n\x06reason\x18\x02 \x01(\x0e\x32*.prodvana.desired_state.model.StatusReason\x12\x0f\n\x07message\x18\x03 \x01(\t\x12\x10\n\x08messages\x18\x04 \x03(\t\"\x8b\x01\n\x11\x41\x63tionExplanation\x12&\n\x02ts\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12=\n\x0b\x61\x63tion_type\x18\x02 \x01(\x0e\x32(.prodvana.desired_state.model.ActionType\x12\x0f\n\x07message\x18\x03 \x01(\t\"\xde\x01\n\x07Version\x12\x0f\n\x07version\x18\x01 \x01(\t\x12\x10\n\x08replicas\x18\x02 \x01(\x05\x12\x1a\n\x12\x61vailable_replicas\x18\x08 \x01(\x05\x12\x32\n\x0epush_timestamp\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x0e\n\x06\x61\x63tive\x18\x05 \x01(\x08\x12\x17\n\x0ftarget_replicas\x18\x06 \x01(\x05\x12\r\n\x05\x64irty\x18\t \x01(\x08J\x04\x08\x03\x10\x04J\x04\x08\x07\x10\x08R\tunhealthyR\x11unhealthy_reasons\"\xc2\x03\n\x14ServiceInstanceState\x12\x34\n\x04meta\x18\x01 \x01(\x0b\x32&.prodvana.desired_state.model.Metadata\x12\x13\n\x0b\x61pplication\x18\x02 \x01(\t\x12\x0f\n\x07service\x18\x03 \x01(\t\x12\x17\n\x0frelease_channel\x18\x04 \x01(\t\x12\x12\n\nservice_id\x18\x08 \x01(\t\x12\x1a\n\x12release_channel_id\x18\t \x01(\t\x12\x37\n\x08versions\x18\x05 \x03(\x0b\x32%.prodvana.desired_state.model.Version\x12\x41\n\x10rollback_version\x18\x06 \x01(\x0b\x32%.prodvana.desired_state.model.VersionH\x00\x12\"\n\x18\x63ompute_rollback_version\x18\x0b \x01(\x08H\x00\x12\x10\n\x08rollback\x18\x07 \x01(\x08\x12=\n\x08\x64\x65livery\x18\n \x01(\x0b\x32+.prodvana.desired_state.model.DeliveryStateB\x14\n\x12\x61utorollback_oneof\"\xd5\x03\n\x0cServiceState\x12\x34\n\x04meta\x18\x01 \x01(\x0b\x32&.prodvana.desired_state.model.Metadata\x12\x13\n\x0b\x61pplication\x18\x02 \x01(\t\x12\x0f\n\x07service\x18\x03 \x01(\t\x12\x12\n\nservice_id\x18\x05 \x01(\t\x12L\n\x10release_channels\x18\x06 \x03(\x0b\x32\x32.prodvana.desired_state.model.ServiceInstanceState\x12\x43\n\x0c\x63ustom_tasks\x18\x07 \x03(\x0b\x32-.prodvana.desired_state.model.CustomTaskState\x12Q\n\x13\x64\x65livery_extensions\x18\t \x03(\x0b\x32\x34.prodvana.desired_state.model.DeliveryExtensionState\x12\x63\n\x1frelease_channel_label_selectors\x18\n \x03(\x0b\x32:.prodvana.desired_state.model.ServiceInstanceLabelSelectorJ\x04\x08\x04\x10\x05J\x04\x08\x08\x10\t\"\xc5\x02\n\x1cServiceInstanceLabelSelector\x12\"\n\x18release_channel_selector\x18\x01 \x01(\tH\x00\x12\r\n\x03\x61ll\x18\x02 \x01(\x08H\x00\x12\x37\n\x08versions\x18\x03 \x03(\x0b\x32%.prodvana.desired_state.model.Version\x12\x41\n\x10rollback_version\x18\x04 \x01(\x0b\x32%.prodvana.desired_state.model.VersionH\x01\x12\"\n\x18\x63ompute_rollback_version\x18\x05 \x01(\x08H\x01\x12%\n\x1dmaterialized_release_channels\x18\x06 \x03(\tB\x15\n\x0eselector_oneof\x12\x03\xf8\x42\x01\x42\x14\n\x12\x61utorollback_oneof\"\xa5\x02\n\x11ServiceGroupState\x12\x34\n\x04meta\x18\x01 \x01(\x0b\x32&.prodvana.desired_state.model.Metadata\x12<\n\x08services\x18\x02 \x03(\x0b\x32*.prodvana.desired_state.model.ServiceState\x12\x43\n\x0c\x63ustom_tasks\x18\x03 \x03(\x0b\x32-.prodvana.desired_state.model.CustomTaskState\x12Q\n\x13\x64\x65livery_extensions\x18\x05 \x03(\x0b\x32\x34.prodvana.desired_state.model.DeliveryExtensionStateJ\x04\x08\x04\x10\x05\"\xa8\x02\n\x13\x43\x61naryProgressState\x12H\n\x06status\x18\x01 \x01(\x0e\x32\x38.prodvana.desired_state.model.CanaryProgressState.Status\x12 \n\rcanary_weight\x18\x02 \x01(\x05\x42\t\xfa\x42\x06\x1a\x04\x18\x64(\x00\x12+\n\x08\x64uration\x18\x03 \x01(\x0b\x32\x19.google.protobuf.Duration\x12\x39\n\x15pause_start_timestamp\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"=\n\x06Status\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x0b\n\x07PENDING\x10\x01\x12\n\n\x06PAUSED\x10\x02\x12\r\n\tCOMPLETED\x10\x03\"\x94\x03\n\rDeliveryState\x12\x18\n\x10\x64\x65sired_state_id\x18\x0c \x01(\t\x12\x42\n\x06status\x18\x08 \x01(\x0e\x32\x32.prodvana.desired_state.model.DeliveryState.Status\x12\x0f\n\x07message\x18\x06 \x01(\t\x12J\n\x0f\x63\x61nary_progress\x18\x0b \x03(\x0b\x32\x31.prodvana.desired_state.model.CanaryProgressState\x12\x11\n\tfirst_run\x18\r \x01(\x08\x12\x12\n\ngeneration\x18\x0e \x01(\t\"q\n\x06Status\x12\x12\n\x0eSTATUS_UNKNOWN\x10\x00\x12\x16\n\x12STATUS_PROGRESSING\x10\x01\x12\x11\n\rSTATUS_PAUSED\x10\x02\x12\x12\n\x0eSTATUS_HEALTHY\x10\x03\x12\x14\n\x10STATUS_UNHEALTHY\x10\x04J\x04\x08\x01\x10\x02J\x04\x08\x02\x10\x03J\x04\x08\x03\x10\x04J\x04\x08\x04\x10\x05J\x04\x08\x05\x10\x06J\x04\x08\x07\x10\x08J\x04\x08\t\x10\nJ\x04\x08\n\x10\x0b\"\x98\x03\n\x0c\x46\x65tchDetails\x12\x35\n\x11started_timestamp\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x37\n\x13\x63ompleted_timestamp\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x1a\n\x12\x66\x65tch_plan_blob_id\x18\x03 \x01(\t\x12&\n\x1e\x66\x65tch_plan_explanation_blob_id\x18\x04 \x01(\t\x12\x0f\n\x07version\x18\x05 \x01(\t\x12L\n\x0c\x66\x65tch_status\x18\x06 \x01(\x0e\x32\x36.prodvana.desired_state.model.FetchDetails.FetchStatus\x12 \n\x18\x66\x65tcher_desired_state_id\x18\x07 \x01(\t\x12\x0f\n\x07message\x18\x08 \x01(\t\"B\n\x0b\x46\x65tchStatus\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x0b\n\x07\x44RIFTED\x10\x01\x12\r\n\tCONVERGED\x10\x02\x12\n\n\x06\x46\x41ILED\x10\x03\"\x8f\x03\n\x1bRuntimeExtensionFetchOutput\x12\x41\n\rongoing_fetch\x18\x08 \x01(\x0b\x32*.prodvana.desired_state.model.FetchDetails\x12I\n\x15last_successful_fetch\x18\t \x01(\x0b\x32*.prodvana.desired_state.model.FetchDetails\x12\x45\n\x11last_failed_fetch\x18\n \x01(\x0b\x32*.prodvana.desired_state.model.FetchDetailsJ\x04\x08\x01\x10\x02J\x04\x08\x02\x10\x03J\x04\x08\x03\x10\x04J\x04\x08\x04\x10\x05J\x04\x08\x05\x10\x06J\x04\x08\x06\x10\x07J\x04\x08\x07\x10\x08R\x12\x66\x65tch_plan_blob_idR\x1e\x66\x65tch_plan_explanation_blob_idR\x07versionR\x0c\x66\x65tch_statusR\x11started_timestampR\x13\x63ompleted_timestamp\"@\n\x1fRuntimeExtensionUserDebugOutput\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0f\n\x07\x62lob_id\x18\x02 \x01(\t\"\x92\x03\n\x1cRuntimeExtensionDebugDetails\x12\x16\n\x0esystem_message\x18\x01 \x01(\t\x12T\n\rdebug_outputs\x18\x02 \x03(\x0b\x32=.prodvana.desired_state.model.RuntimeExtensionUserDebugOutput\x12Q\n\x06status\x18\x03 \x01(\x0e\x32\x41.prodvana.desired_state.model.RuntimeExtensionDebugDetails.Status\x12\x35\n\x11started_timestamp\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x37\n\x13\x63ompleted_timestamp\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x0f\n\x07version\x18\x06 \x01(\t\"0\n\x06Status\x12\x0b\n\x07UNKNOWN\x10\x00\x12\r\n\tCONVERGED\x10\x01\x12\n\n\x06\x46\x41ILED\x10\x02\"p\n\x1bRuntimeExtensionDebugOutput\x12Q\n\rdebug_details\x18\x01 \x01(\x0b\x32:.prodvana.desired_state.model.RuntimeExtensionDebugDetails\"\xe6\x02\n\x0c\x41pplyDetails\x12\x35\n\x11started_timestamp\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x37\n\x13\x63ompleted_timestamp\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x0f\n\x07version\x18\x03 \x01(\t\x12L\n\x0c\x61pply_status\x18\x04 \x01(\x0e\x32\x36.prodvana.desired_state.model.ApplyDetails.ApplyStatus\x12\x41\n\rfetch_details\x18\x05 \x01(\x0b\x32*.prodvana.desired_state.model.FetchDetails\"D\n\x0b\x41pplyStatus\x12\x0b\n\x07UNKNOWN\x10\x00\x12\r\n\tCONVERGED\x10\x01\x12\n\n\x06\x46\x41ILED\x10\x02\x12\r\n\tRETRYABLE\x10\x03\"\xa0\x01\n\x1bRuntimeExtensionApplyOutput\x12\x41\n\rongoing_apply\x18\x01 \x01(\x0b\x32*.prodvana.desired_state.model.ApplyDetails\x12>\n\nlast_apply\x18\x02 \x01(\x0b\x32*.prodvana.desired_state.model.ApplyDetails\"\xac\x01\n\x0c\x45xternalLink\x12\x41\n\x04type\x18\x01 \x01(\x0e\x32\x33.prodvana.desired_state.model.ExternalLink.LinkType\x12\x14\n\x03url\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x15\n\x04name\x18\x03 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\",\n\x08LinkType\x12\x0b\n\x07UNKNOWN\x10\x00\x12\n\n\x06\x44\x45TAIL\x10\x01\x12\x07\n\x03LOG\x10\x02\"\xf3\x11\n\rRuntimeObject\x12\x34\n\x04meta\x18\x01 \x01(\x0b\x32&.prodvana.desired_state.model.Metadata\x12\x13\n\x0bobject_type\x18\x02 \x01(\t\x12\x11\n\tnamespace\x18\x03 \x01(\t\x12\x0c\n\x04name\x18\x04 \x01(\t\x12\x15\n\rgenerate_name\x18\x18 \x01(\t\x12\x37\n\x08versions\x18\x05 \x03(\x0b\x32%.prodvana.desired_state.model.Version\x12\x42\n\x06status\x18\x06 \x01(\x0e\x32\x32.prodvana.desired_state.model.RuntimeObject.Status\x12?\n\x10rollback_version\x18\x07 \x01(\x0b\x32%.prodvana.desired_state.model.Version\x12=\n\x08\x64\x65livery\x18\x08 \x01(\x0b\x32+.prodvana.desired_state.model.DeliveryState\x12\x18\n\x10version_agnostic\x18\n \x01(\x08\x12\"\n\x1a\x64\x65sired_version_dirty_only\x18\x14 \x01(\x08\x12\x0f\n\x07message\x18\x0c \x01(\t\x12W\n\x11runtime_extension\x18\r \x01(\x0b\x32<.prodvana.desired_state.model.RuntimeObject.RuntimeExtension\x12+\n\x08interval\x18\x0f \x01(\x0b\x32\x19.google.protobuf.Duration\x12\x38\n\x15steady_state_interval\x18\x1a \x01(\x0b\x32\x19.google.protobuf.Duration\x12\x17\n\x0foutput_blob_ids\x18\x12 \x03(\t\x12\x12\n\nexit_codes\x18\x13 \x03(\x05\x12%\n\x1drequire_approval_before_apply\x18\x15 \x01(\x08\x12\x12\n\nraw_config\x18\x17 \x01(\t\x12W\n\x11management_status\x18\x19 \x01(\x0e\x32<.prodvana.desired_state.model.RuntimeObject.ManagementStatus\x12\x46\n\x17last_completed_task_run\x18\x1b \x01(\x0b\x32%.prodvana.desired_state.model.TaskRun\x12\x42\n\x0e\x65xternal_links\x18\x1c \x03(\x0b\x32*.prodvana.desired_state.model.ExternalLink\x1a\xc4\x08\n\x10RuntimeExtension\x12=\n\x05\x61pply\x18\x03 \x01(\x0b\x32..prodvana.environment.CompiledExtensionCommand\x12=\n\x05\x66\x65tch\x18\x01 \x01(\x0b\x32..prodvana.environment.CompiledExtensionCommand\x12=\n\x05\x64\x65\x62ug\x18\x11 \x01(\x0b\x32..prodvana.environment.CompiledExtensionCommand\x12\x31\n\x0e\x66\x65tch_interval\x18\x04 \x01(\x0b\x32\x19.google.protobuf.Duration\x12>\n\x1b\x66\x65tch_steady_state_interval\x18\x0b \x01(\x0b\x32\x19.google.protobuf.Duration\x12\x31\n\x0e\x64\x65\x62ug_interval\x18\x13 \x01(\x0b\x32\x19.google.protobuf.Duration\x12\x12\n\nservice_id\x18\x02 \x01(\t\x12\x1a\n\x12release_channel_id\x18\x06 \x01(\t\x12?\n\nparameters\x18\x07 \x03(\x0b\x32+.prodvana.common_config.ParameterDefinition\x12@\n\x10parameter_values\x18\x08 \x03(\x0b\x32&.prodvana.common_config.ParameterValue\x12\x31\n\x04type\x18\t \x01(\x0e\x32#.prodvana.environment.ExtensionType\x12O\n\x0c\x66\x65tch_output\x18\n \x01(\x0b\x32\x39.prodvana.desired_state.model.RuntimeExtensionFetchOutput\x12O\n\x0c\x64\x65\x62ug_output\x18\x10 \x01(\x0b\x32\x39.prodvana.desired_state.model.RuntimeExtensionDebugOutput\x12O\n\x0c\x61pply_output\x18\x0c \x01(\x0b\x32\x39.prodvana.desired_state.model.RuntimeExtensionApplyOutput\x12=\n\x12\x66\x65tch_retry_policy\x18\r \x01(\x0b\x32!.prodvana.environment.RetryPolicy\x12=\n\x12\x61pply_retry_policy\x18\x0e \x01(\x0b\x32!.prodvana.environment.RetryPolicy\x12=\n\x12\x64\x65\x62ug_retry_policy\x18\x12 \x01(\x0b\x32!.prodvana.environment.RetryPolicy\x12\"\n\x1a\x61pply_retryable_exit_codes\x18\x0f \x03(\x05J\x04\x08\x05\x10\x06R\rfetch_timeout\"0\n\x06Status\x12\x0b\n\x07PENDING\x10\x00\x12\r\n\tSUCCEEDED\x10\x01\x12\n\n\x06\x46\x41ILED\x10\x02\"2\n\x10ManagementStatus\x12\x0f\n\x0bPVN_MANAGED\x10\x00\x12\r\n\tUNMANAGED\x10\x01J\x04\x08\t\x10\nJ\x04\x08\x0b\x10\x0cJ\x04\x08\x0e\x10\x0fJ\x04\x08\x10\x10\x11J\x04\x08\x11\x10\x12J\x04\x08\x16\x10\x17R\x0eunhealthy_podsR\nstate_hashR\x10missing_approvalR\x07timeout\"O\n\x0e\x43onditionState\x12=\n\x06status\x18\x01 \x01(\x0e\x32-.prodvana.desired_state.model.ConditionStatus\"\xb6\x04\n\x0c\x43ontrolState\x12\x10\n\x08rollback\x18\x01 \x01(\x08\x12I\n\x13precondition_states\x18\x02 \x03(\x0b\x32,.prodvana.desired_state.model.ConditionState\x12\x46\n\x10invariant_states\x18\x03 \x03(\x0b\x32,.prodvana.desired_state.model.ConditionState\x12\x0e\n\x06paused\x18\x04 \x01(\x08\x12L\n\x13status_explanations\x18\x05 \x03(\x0b\x32/.prodvana.desired_state.model.StatusExplanation\x12K\n\x12\x61\x63tion_explanation\x18\x06 \x01(\x0b\x32/.prodvana.desired_state.model.ActionExplanation\x12:\n\x16last_fetched_timestamp\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12:\n\x16last_applied_timestamp\x18\x08 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12G\n\x10missing_approval\x18\t \x01(\x0b\x32-.prodvana.desired_state.model.MissingApproval\x12\x15\n\robserver_mode\x18\n \x01(\x08\"\xb3\x01\n\x13ManualApprovalState\x12\x34\n\x04meta\x18\x01 \x01(\x0b\x32&.prodvana.desired_state.model.Metadata\x12\x42\n\x06status\x18\x02 \x01(\x0e\x32\x32.prodvana.desired_state.model.ManualApprovalStatus\x12\r\n\x05topic\x18\x03 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x04 \x01(\t\"\xc1\x05\n\x05State\x12=\n\x07service\x18\x01 \x01(\x0b\x32*.prodvana.desired_state.model.ServiceStateH\x00\x12N\n\x10service_instance\x18\x02 \x01(\x0b\x32\x32.prodvana.desired_state.model.ServiceInstanceStateH\x00\x12H\n\rservice_group\x18\x03 \x01(\x0b\x32/.prodvana.desired_state.model.ServiceGroupStateH\x00\x12\x45\n\x0eruntime_object\x18\x04 \x01(\x0b\x32+.prodvana.desired_state.model.RuntimeObjectH\x00\x12L\n\x0fmanual_approval\x18\x05 \x01(\x0b\x32\x31.prodvana.desired_state.model.ManualApprovalStateH\x00\x12\x44\n\x0b\x63ustom_task\x18\x06 \x01(\x0b\x32-.prodvana.desired_state.model.CustomTaskStateH\x00\x12S\n\x15protection_attachment\x18\x07 \x01(\x0b\x32\x32.prodvana.desired_state.model.ProtectionAttachmentH\x00\x12L\n\x0fprotection_link\x18\x08 \x01(\x0b\x32\x31.prodvana.desired_state.model.ProtectionLinkStateH\x00\x12R\n\x12\x64\x65livery_extension\x18\t \x01(\x0b\x32\x34.prodvana.desired_state.model.DeliveryExtensionStateH\x00\x42\r\n\x0bstate_oneof\"\x82\x01\n\x0b\x41nnotations\x12I\n\x0b\x61nnotations\x18\x01 \x03(\x0b\x32\x34.prodvana.desired_state.model.Annotations.Annotation\x1a(\n\nAnnotation\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t\"\xae\x01\n\x18\x43ustomTaskExecutionState\x12>\n\x06status\x18\x01 \x01(\x0e\x32..prodvana.desired_state.model.CustomTaskStatus\x12\x10\n\x08\x61ttempts\x18\x02 \x01(\x03\x12@\n\x1clatest_attempt_end_timestamp\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\xce\x03\n\x0f\x43ustomTaskState\x12\x34\n\x04meta\x18\x01 \x01(\x0b\x32&.prodvana.desired_state.model.Metadata\x12\x15\n\x04name\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x1c\n\x0b\x64\x65scription\x18\x03 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x13\n\x0b\x61pplication\x18\x04 \x01(\t\x12\x16\n\x0e\x61pplication_id\x18\x05 \x01(\t\x12\x17\n\x0frelease_channel\x18\x06 \x01(\t\x12\x1a\n\x12release_channel_id\x18\x07 \x01(\t\x12@\n\x07program\x18\t \x01(\x0b\x32%.prodvana.common_config.ProgramConfigB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12J\n\ntask_state\x18\x0c \x01(\x0b\x32\x36.prodvana.desired_state.model.CustomTaskExecutionState\x12\x39\n\x0cretry_config\x18\r \x01(\x0b\x32#.prodvana.common_config.RetryConfig\x12\x13\n\x0bservice_ids\x18\x0e \x03(\tJ\x04\x08\x08\x10\tJ\x04\x08\n\x10\x0bJ\x04\x08\x0b\x10\x0c\"\xeb\x04\n\x13ProtectionLinkState\x12\x34\n\x04meta\x18\x01 \x01(\x0b\x32&.prodvana.desired_state.model.Metadata\x12=\n\x06status\x18\x02 \x01(\x0e\x32-.prodvana.desired_state.model.ConditionStatus\x12:\n\x04link\x18\x03 \x01(\x0b\x32,.prodvana.desired_state.model.ProtectionLink\x12\x35\n\x11started_timestamp\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x35\n\x11stopped_timestamp\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12T\n\x0estopped_reason\x18\x06 \x01(\x0e\x32<.prodvana.desired_state.model.ProtectionLinkState.StopReason\x12;\n\x17\x66irst_success_timestamp\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\xa1\x01\n\nStopReason\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x17\n\x13LIFECYCLE_COMPLETED\x10\x01\x12\x12\n\x0eSUCCEEDED_ONCE\x10\x02\x12\x1a\n\x16SUCCEEDED_FOR_DURATION\x10\x03\x12\r\n\tTIMED_OUT\x10\x04\x12\n\n\x06\x46\x41ILED\x10\x05\x12\x0b\n\x07\x44\x45LETED\x10\x06\x12\x15\n\x11MANUALLY_BYPASSED\x10\x07\"\x89\x04\n\x14ProtectionAttachment\x12\x34\n\x04meta\x18\x01 \x01(\x0b\x32&.prodvana.desired_state.model.Metadata\x12\x37\n\x08versions\x18\x02 \x03(\x0b\x32%.prodvana.desired_state.model.Version\x12\x46\n\x17last_completed_versions\x18\x05 \x03(\x0b\x32%.prodvana.desired_state.model.Version\x12<\n\x18last_completed_timestamp\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12I\n\x15last_completed_status\x18\x07 \x01(\x0e\x32*.prodvana.desired_state.model.SimpleStatus\x12[\n\"last_completed_status_explanations\x18\x08 \x03(\x0b\x32/.prodvana.desired_state.model.StatusExplanation\x12&\n\x1elast_completed_applied_version\x18\t \x01(\t\x12\x15\n\rprotection_id\x18\x03 \x01(\t\x12\x15\n\rattachment_id\x18\x04 \x01(\t\"\x81\x05\n\x16\x44\x65liveryExtensionState\x12\x34\n\x04meta\x18\x01 \x01(\x0b\x32&.prodvana.desired_state.model.Metadata\x12\x37\n\x08versions\x18\x02 \x03(\x0b\x32%.prodvana.desired_state.model.Version\x12\x14\n\x0c\x65xtension_id\x18\x03 \x01(\t\x12\x1d\n\x15\x65xtension_instance_id\x18\n \x01(\t\x12\x38\n\tlifecycle\x18\x0b \x01(\x0e\x32%.prodvana.common_config.TaskLifecycle\x12<\n\x18last_completed_timestamp\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12I\n\x15last_completed_status\x18\x07 \x01(\x0e\x32*.prodvana.desired_state.model.SimpleStatus\x12[\n\"last_completed_status_explanations\x18\x08 \x03(\x0b\x32/.prodvana.desired_state.model.StatusExplanation\x12&\n\x1elast_completed_applied_version\x18\t \x01(\t\x12<\n\nreferences\x18\x0c \x03(\x0b\x32(.prodvana.desired_state.model.Identifier\x12=\n\x06status\x18\r \x01(\x0e\x32-.prodvana.desired_state.model.ConditionStatus\"\x9c\x04\n\x06Signal\x12\x36\n\x04type\x18\x01 \x01(\x0e\x32(.prodvana.desired_state.model.SignalType\x12Z\n\x12\x64\x65livery_promotion\x18\x02 \x01(\x0b\x32<.prodvana.desired_state.model.Signal.DeliveryPromotionConfigH\x00\x12R\n\x11protection_bypass\x18\x03 \x01(\x0b\x32\x35.prodvana.desired_state.model.Signal.ProtectionBypassH\x00\x12\x63\n\x1aruntime_extension_approval\x18\x04 \x01(\x0b\x32=.prodvana.desired_state.model.Signal.RuntimeExtensionApprovalH\x00\x1a\x36\n\x17\x44\x65liveryPromotionConfig\x12\r\n\x05stage\x18\x01 \x01(\x03\x12\x0c\n\x04\x66ull\x18\x02 \x01(\x08\x1a\x12\n\x10ProtectionBypass\x1ao\n\x18RuntimeExtensionApproval\x12-\n\ttimestamp\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x0e\n\x06reject\x18\x02 \x01(\x08\x12\x14\n\x0cplan_blob_id\x18\x03 \x01(\tB\x08\n\x06\x63onfig\"?\n\x08\x44\x65\x62ugLog\x12&\n\x02ts\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x0b\n\x03log\x18\x02 \x01(\t\"\x85\x01\n\x18RuntimeExtensionMetadata\x12I\n\x06output\x18\x01 \x01(\x0b\x32\x39.prodvana.desired_state.model.RuntimeExtensionFetchOutput\x12\x1e\n\x16\x61pply_desired_state_id\x18\x02 \x01(\t\"u\n\x14ManualApprovalConfig\x12S\n\x11runtime_extension\x18\x01 \x01(\x0b\x32\x36.prodvana.desired_state.model.RuntimeExtensionMetadataH\x00\x42\x08\n\x06\x63onfig\"\xdd\x01\n\x0fMissingApproval\x12\x18\n\x10\x64\x65sired_state_id\x18\x01 \x01(\t\x12=\n\x0bsignal_type\x18\x02 \x01(\x0e\x32(.prodvana.desired_state.model.SignalType\x12\r\n\x05topic\x18\x03 \x01(\t\x12Q\n\x11runtime_extension\x18\x04 \x01(\x0b\x32\x36.prodvana.desired_state.model.RuntimeExtensionMetadata\x12\x0f\n\x07\x63urrent\x18\x05 \x01(\x08\"\x8e\x03\n\x19\x41pplyConditionUnsatisfied\x12k\n\x10missing_approval\x18\x01 \x01(\x0b\x32O.prodvana.desired_state.model.ApplyConditionUnsatisfied.InternalMissingApprovalH\x00\x1a\xf9\x01\n\x17InternalMissingApproval\x12\x18\n\x10\x64\x65sired_state_id\x18\x01 \x01(\t\x12=\n\x0bsignal_type\x18\x02 \x01(\x0e\x32(.prodvana.desired_state.model.SignalType\x12\r\n\x05topic\x18\x03 \x01(\t\x12Q\n\x11runtime_extension\x18\x04 \x01(\x0b\x32\x36.prodvana.desired_state.model.RuntimeExtensionMetadata\x12#\n\x1bgenerator_desired_state_ids\x18\x05 \x03(\tB\x08\n\x06reason\"\x97\x05\n\x07TaskRun\x12:\n\x06status\x18\x01 \x01(\x0e\x32*.prodvana.desired_state.model.SimpleStatus\x12L\n\x13status_explanations\x18\x02 \x03(\x0b\x32/.prodvana.desired_state.model.StatusExplanation\x12\x0f\n\x07version\x18\x03 \x01(\t\x12<\n\rseen_versions\x18\x04 \x03(\x0b\x32%.prodvana.desired_state.model.Version\x12\x35\n\x11started_timestamp\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x37\n\x13\x63ompleted_timestamp\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x17\n\x0foutput_blob_ids\x18\x07 \x03(\t\x12\x12\n\nexit_codes\x18\x08 \x03(\x05\x12?\n\rtask_entities\x18\t \x03(\x0b\x32(.prodvana.desired_state.model.Identifier\x12\x41\n\rfetch_details\x18\n \x01(\x0b\x32*.prodvana.desired_state.model.FetchDetails\x12\x1c\n\x14retryable_exit_codes\x18\x0b \x03(\x05\x12\x11\n\tretryable\x18\x0c \x01(\x08\x12:\n\x05phase\x18\r \x01(\x0e\x32+.prodvana.desired_state.model.TaskRun.Phase\"%\n\x05Phase\x12\x0b\n\x07RUNNING\x10\x00\x12\x0f\n\x0bNOT_STARTED\x10\x01\"\xd8\x01\n\x11TaskEntityContext\x12\x41\n\x12last_completed_run\x18\x01 \x01(\x0b\x32%.prodvana.desired_state.model.TaskRun\x12\x37\n\x08last_run\x18\x02 \x01(\x0b\x32%.prodvana.desired_state.model.TaskRun\x12\x11\n\tis_active\x18\x03 \x01(\x08J\x04\x08\x04\x10\x05J\x04\x08\x05\x10\x06R\x10missing_approvalR\x16is_not_in_grace_period*\xcb\x01\n\x04Type\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x0b\n\x07SERVICE\x10\x01\x12\x14\n\x10SERVICE_INSTANCE\x10\x02\x12\x11\n\rSERVICE_GROUP\x10\x03\x12\x12\n\x0eRUNTIME_OBJECT\x10\x04\x12\x13\n\x0fMANUAL_APPROVAL\x10\x05\x12\x0f\n\x0b\x43USTOM_TASK\x10\x06\x12\x19\n\x15PROTECTION_ATTACHMENT\x10\x07\x12\x13\n\x0fPROTECTION_LINK\x10\x08\x12\x16\n\x12\x44\x45LIVERY_EXTENSION\x10\t*\x9a\x01\n\x0e\x43ustomTaskType\x12\x1c\n\x18\x43USTOM_TASK_TYPE_UNKNOWN\x10\x00\x12\x15\n\x11PRE_APPROVAL_TASK\x10\x01\x12\x0c\n\x08\x41PPROVAL\x10\x02\x12\x16\n\x12POST_APPROVAL_TASK\x10\x03\x12\x13\n\x0f\x44\x45PLOYMENT_TASK\x10\x04\x12\x18\n\x14POST_DEPLOYMENT_TASK\x10\x05*\xfd\x01\n\x06Status\x12\x12\n\x0eUNKNOWN_STATUS\x10\x00\x12\x0e\n\nCONVERGING\x10\x01\x12\r\n\tCONVERGED\x10\x02\x12\n\n\x06\x46\x41ILED\x10\x03\x12\x10\n\x0cROLLING_BACK\x10\x04\x12\x0f\n\x0bROLLED_BACK\x10\x05\x12\x13\n\x0f\x46\x41ILED_ROLLBACK\x10\x0c\x12\n\n\x06PAUSED\x10\x06\x12\x10\n\x0c\x43HILD_PAUSED\x10\x07\x12\x19\n\x15WAITING_PRECONDITIONS\x10\x08\x12\x0c\n\x08REPLACED\x10\t\x12\x1b\n\x17WAITING_MANUAL_APPROVAL\x10\n\x12\x0b\n\x07\x44\x45LETED\x10\x0b\x12\x0b\n\x07PREVIEW\x10\r*o\n\x0cSimpleStatus\x12\x0e\n\nSS_UNKNOWN\x10\x00\x12\x11\n\rSS_CONVERGING\x10\x01\x12\x10\n\x0cSS_CONVERGED\x10\x02\x12\r\n\tSS_FAILED\x10\x03\x12\x1b\n\x17SS_WAITING_FOR_APPROVAL\x10\x04*\xff\x01\n\x0cStatusReason\x12\x12\n\x0eREASON_UNKNOWN\x10\x00\x12\x14\n\x10NO_CURRENT_STATE\x10\x01\x12\x10\n\x0c\x41PPLY_FAILED\x10\x02\x12\x12\n\x0eUNHEALTHY_PODS\x10\x03\x12\x11\n\rUPDATING_PODS\x10\x04\x12\x14\n\x10VERSION_MISMATCH\x10\x05\x12\x19\n\x15RUNTIME_OBJECT_FAILED\x10\x06\x12\x18\n\x14PRECONDITIONS_FAILED\x10\x07\x12\x1c\n\x18MANUAL_APPROVAL_REJECTED\x10\x08\x12\x10\n\x0cSTUCK_ENTITY\x10\t\x12\x11\n\rVERSION_DIRTY\x10\n*r\n\nActionType\x12\x17\n\x13\x41\x43TION_TYPE_UNKNOWN\x10\x00\x12\x18\n\x14\x41\x43TION_TYPE_APPLYING\x10\x01\x12\x17\n\x13\x41\x43TION_TYPE_APPLIED\x10\x02\x12\x18\n\x14\x41\x43TION_TYPE_COMPLETE\x10\x03*\x96\x01\n\x0f\x43onditionStatus\x12\x1c\n\x18\x43ONDITION_UNKNOWN_STATUS\x10\x00\x12\x15\n\x11\x43ONDITION_PENDING\x10\x01\x12\x17\n\x13\x43ONDITION_SATISFIED\x10\x02\x12\x1f\n\x1b\x43ONDITION_MANUALLY_BYPASSED\x10\x03\x12\x14\n\x10\x43ONDITION_FAILED\x10\x04*?\n\x14ManualApprovalStatus\x12\x0b\n\x07PENDING\x10\x00\x12\x0c\n\x08\x41PPROVED\x10\x01\x12\x0c\n\x08REJECTED\x10\x02*\x85\x01\n\x10\x43ustomTaskStatus\x12\x17\n\x13\x43USTOM_TASK_PENDING\x10\x00\x12\x1a\n\x16\x43USTOM_TASK_SUCCESSFUL\x10\x01\x12!\n\x1d\x43USTOM_TASK_RETRIES_EXHAUSTED\x10\x02\x12\x19\n\x15\x43USTOM_TASK_TIMED_OUT\x10\x03*\x8b\x01\n\nSignalType\x12\x12\n\x0eSIGNAL_UNKNOWN\x10\x00\x12\x16\n\x12\x44\x45LIVERY_PROMOTION\x10\x01\x12\x15\n\x11PROTECTION_BYPASS\x10\x02\x12\x1e\n\x1aRUNTIME_EXTENSION_APPROVAL\x10\x03\x12\x1a\n\x16SIGNAL_MANUAL_APPROVAL\x10\x04\x42XZVgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/modelb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'prodvana.desired_state.model.desired_state_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZVgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/model'
  _SERVICEINSTANCELABELSELECTOR.oneofs_by_name['selector_oneof']._options = None
  _SERVICEINSTANCELABELSELECTOR.oneofs_by_name['selector_oneof']._serialized_options = b'\370B\001'
  _CANARYPROGRESSSTATE.fields_by_name['canary_weight']._options = None
  _CANARYPROGRESSSTATE.fields_by_name['canary_weight']._serialized_options = b'\372B\006\032\004\030d(\000'
  _EXTERNALLINK.fields_by_name['url']._options = None
  _EXTERNALLINK.fields_by_name['url']._serialized_options = b'\372B\004r\002\020\001'
  _EXTERNALLINK.fields_by_name['name']._options = None
  _EXTERNALLINK.fields_by_name['name']._serialized_options = b'\372B\004r\002\020\001'
  _CUSTOMTASKSTATE.fields_by_name['name']._options = None
  _CUSTOMTASKSTATE.fields_by_name['name']._serialized_options = b'\372B\004r\002\020\001'
  _CUSTOMTASKSTATE.fields_by_name['description']._options = None
  _CUSTOMTASKSTATE.fields_by_name['description']._serialized_options = b'\372B\004r\002\020\001'
  _CUSTOMTASKSTATE.fields_by_name['program']._options = None
  _CUSTOMTASKSTATE.fields_by_name['program']._serialized_options = b'\372B\005\212\001\002\020\001'
  _globals['_TYPE']._serialized_start=15831
  _globals['_TYPE']._serialized_end=16034
  _globals['_CUSTOMTASKTYPE']._serialized_start=16037
  _globals['_CUSTOMTASKTYPE']._serialized_end=16191
  _globals['_STATUS']._serialized_start=16194
  _globals['_STATUS']._serialized_end=16447
  _globals['_SIMPLESTATUS']._serialized_start=16449
  _globals['_SIMPLESTATUS']._serialized_end=16560
  _globals['_STATUSREASON']._serialized_start=16563
  _globals['_STATUSREASON']._serialized_end=16818
  _globals['_ACTIONTYPE']._serialized_start=16820
  _globals['_ACTIONTYPE']._serialized_end=16934
  _globals['_CONDITIONSTATUS']._serialized_start=16937
  _globals['_CONDITIONSTATUS']._serialized_end=17087
  _globals['_MANUALAPPROVALSTATUS']._serialized_start=17089
  _globals['_MANUALAPPROVALSTATUS']._serialized_end=17152
  _globals['_CUSTOMTASKSTATUS']._serialized_start=17155
  _globals['_CUSTOMTASKSTATUS']._serialized_end=17288
  _globals['_SIGNALTYPE']._serialized_start=17291
  _globals['_SIGNALTYPE']._serialized_end=17430
  _globals['_PROTECTIONLINK']._serialized_start=407
  _globals['_PROTECTIONLINK']._serialized_end=519
  _globals['_CONDITION']._serialized_start=522
  _globals['_CONDITION']._serialized_end=1330
  _globals['_CONDITION_RELEASECHANNELSTABLECONDITION']._serialized_start=829
  _globals['_CONDITION_RELEASECHANNELSTABLECONDITION']._serialized_end=996
  _globals['_CONDITION_MANUALAPPROVAL']._serialized_start=998
  _globals['_CONDITION_MANUALAPPROVAL']._serialized_end=1050
  _globals['_CONDITION_CUSTOMTASKSUCCESSFULCONDITION']._serialized_start=1053
  _globals['_CONDITION_CUSTOMTASKSUCCESSFULCONDITION']._serialized_end=1317
  _globals['_CONDITION_CUSTOMTASKSUCCESSFULCONDITION_PROTECTION']._serialized_start=1216
  _globals['_CONDITION_CUSTOMTASKSUCCESSFULCONDITION_PROTECTION']._serialized_end=1307
  _globals['_DELIVERYEXTENSION']._serialized_start=1333
  _globals['_DELIVERYEXTENSION']._serialized_end=1493
  _globals['_IDENTIFIER']._serialized_start=1495
  _globals['_IDENTIFIER']._serialized_end=1571
  _globals['_METADATA']._serialized_start=1574
  _globals['_METADATA']._serialized_end=2161
  _globals['_STATUSEXPLANATION']._serialized_start=2164
  _globals['_STATUSEXPLANATION']._serialized_end=2337
  _globals['_ACTIONEXPLANATION']._serialized_start=2340
  _globals['_ACTIONEXPLANATION']._serialized_end=2479
  _globals['_VERSION']._serialized_start=2482
  _globals['_VERSION']._serialized_end=2704
  _globals['_SERVICEINSTANCESTATE']._serialized_start=2707
  _globals['_SERVICEINSTANCESTATE']._serialized_end=3157
  _globals['_SERVICESTATE']._serialized_start=3160
  _globals['_SERVICESTATE']._serialized_end=3629
  _globals['_SERVICEINSTANCELABELSELECTOR']._serialized_start=3632
  _globals['_SERVICEINSTANCELABELSELECTOR']._serialized_end=3957
  _globals['_SERVICEGROUPSTATE']._serialized_start=3960
  _globals['_SERVICEGROUPSTATE']._serialized_end=4253
  _globals['_CANARYPROGRESSSTATE']._serialized_start=4256
  _globals['_CANARYPROGRESSSTATE']._serialized_end=4552
  _globals['_CANARYPROGRESSSTATE_STATUS']._serialized_start=4491
  _globals['_CANARYPROGRESSSTATE_STATUS']._serialized_end=4552
  _globals['_DELIVERYSTATE']._serialized_start=4555
  _globals['_DELIVERYSTATE']._serialized_end=4959
  _globals['_DELIVERYSTATE_STATUS']._serialized_start=4798
  _globals['_DELIVERYSTATE_STATUS']._serialized_end=4911
  _globals['_FETCHDETAILS']._serialized_start=4962
  _globals['_FETCHDETAILS']._serialized_end=5370
  _globals['_FETCHDETAILS_FETCHSTATUS']._serialized_start=5304
  _globals['_FETCHDETAILS_FETCHSTATUS']._serialized_end=5370
  _globals['_RUNTIMEEXTENSIONFETCHOUTPUT']._serialized_start=5373
  _globals['_RUNTIMEEXTENSIONFETCHOUTPUT']._serialized_end=5772
  _globals['_RUNTIMEEXTENSIONUSERDEBUGOUTPUT']._serialized_start=5774
  _globals['_RUNTIMEEXTENSIONUSERDEBUGOUTPUT']._serialized_end=5838
  _globals['_RUNTIMEEXTENSIONDEBUGDETAILS']._serialized_start=5841
  _globals['_RUNTIMEEXTENSIONDEBUGDETAILS']._serialized_end=6243
  _globals['_RUNTIMEEXTENSIONDEBUGDETAILS_STATUS']._serialized_start=6195
  _globals['_RUNTIMEEXTENSIONDEBUGDETAILS_STATUS']._serialized_end=6243
  _globals['_RUNTIMEEXTENSIONDEBUGOUTPUT']._serialized_start=6245
  _globals['_RUNTIMEEXTENSIONDEBUGOUTPUT']._serialized_end=6357
  _globals['_APPLYDETAILS']._serialized_start=6360
  _globals['_APPLYDETAILS']._serialized_end=6718
  _globals['_APPLYDETAILS_APPLYSTATUS']._serialized_start=6650
  _globals['_APPLYDETAILS_APPLYSTATUS']._serialized_end=6718
  _globals['_RUNTIMEEXTENSIONAPPLYOUTPUT']._serialized_start=6721
  _globals['_RUNTIMEEXTENSIONAPPLYOUTPUT']._serialized_end=6881
  _globals['_EXTERNALLINK']._serialized_start=6884
  _globals['_EXTERNALLINK']._serialized_end=7056
  _globals['_EXTERNALLINK_LINKTYPE']._serialized_start=7012
  _globals['_EXTERNALLINK_LINKTYPE']._serialized_end=7056
  _globals['_RUNTIMEOBJECT']._serialized_start=7059
  _globals['_RUNTIMEOBJECT']._serialized_end=9350
  _globals['_RUNTIMEOBJECT_RUNTIMEEXTENSION']._serialized_start=8065
  _globals['_RUNTIMEOBJECT_RUNTIMEEXTENSION']._serialized_end=9157
  _globals['_RUNTIMEOBJECT_STATUS']._serialized_start=9159
  _globals['_RUNTIMEOBJECT_STATUS']._serialized_end=9207
  _globals['_RUNTIMEOBJECT_MANAGEMENTSTATUS']._serialized_start=9209
  _globals['_RUNTIMEOBJECT_MANAGEMENTSTATUS']._serialized_end=9259
  _globals['_CONDITIONSTATE']._serialized_start=9352
  _globals['_CONDITIONSTATE']._serialized_end=9431
  _globals['_CONTROLSTATE']._serialized_start=9434
  _globals['_CONTROLSTATE']._serialized_end=10000
  _globals['_MANUALAPPROVALSTATE']._serialized_start=10003
  _globals['_MANUALAPPROVALSTATE']._serialized_end=10182
  _globals['_STATE']._serialized_start=10185
  _globals['_STATE']._serialized_end=10890
  _globals['_ANNOTATIONS']._serialized_start=10893
  _globals['_ANNOTATIONS']._serialized_end=11023
  _globals['_ANNOTATIONS_ANNOTATION']._serialized_start=10983
  _globals['_ANNOTATIONS_ANNOTATION']._serialized_end=11023
  _globals['_CUSTOMTASKEXECUTIONSTATE']._serialized_start=11026
  _globals['_CUSTOMTASKEXECUTIONSTATE']._serialized_end=11200
  _globals['_CUSTOMTASKSTATE']._serialized_start=11203
  _globals['_CUSTOMTASKSTATE']._serialized_end=11665
  _globals['_PROTECTIONLINKSTATE']._serialized_start=11668
  _globals['_PROTECTIONLINKSTATE']._serialized_end=12287
  _globals['_PROTECTIONLINKSTATE_STOPREASON']._serialized_start=12126
  _globals['_PROTECTIONLINKSTATE_STOPREASON']._serialized_end=12287
  _globals['_PROTECTIONATTACHMENT']._serialized_start=12290
  _globals['_PROTECTIONATTACHMENT']._serialized_end=12811
  _globals['_DELIVERYEXTENSIONSTATE']._serialized_start=12814
  _globals['_DELIVERYEXTENSIONSTATE']._serialized_end=13455
  _globals['_SIGNAL']._serialized_start=13458
  _globals['_SIGNAL']._serialized_end=13998
  _globals['_SIGNAL_DELIVERYPROMOTIONCONFIG']._serialized_start=13801
  _globals['_SIGNAL_DELIVERYPROMOTIONCONFIG']._serialized_end=13855
  _globals['_SIGNAL_PROTECTIONBYPASS']._serialized_start=13857
  _globals['_SIGNAL_PROTECTIONBYPASS']._serialized_end=13875
  _globals['_SIGNAL_RUNTIMEEXTENSIONAPPROVAL']._serialized_start=13877
  _globals['_SIGNAL_RUNTIMEEXTENSIONAPPROVAL']._serialized_end=13988
  _globals['_DEBUGLOG']._serialized_start=14000
  _globals['_DEBUGLOG']._serialized_end=14063
  _globals['_RUNTIMEEXTENSIONMETADATA']._serialized_start=14066
  _globals['_RUNTIMEEXTENSIONMETADATA']._serialized_end=14199
  _globals['_MANUALAPPROVALCONFIG']._serialized_start=14201
  _globals['_MANUALAPPROVALCONFIG']._serialized_end=14318
  _globals['_MISSINGAPPROVAL']._serialized_start=14321
  _globals['_MISSINGAPPROVAL']._serialized_end=14542
  _globals['_APPLYCONDITIONUNSATISFIED']._serialized_start=14545
  _globals['_APPLYCONDITIONUNSATISFIED']._serialized_end=14943
  _globals['_APPLYCONDITIONUNSATISFIED_INTERNALMISSINGAPPROVAL']._serialized_start=14684
  _globals['_APPLYCONDITIONUNSATISFIED_INTERNALMISSINGAPPROVAL']._serialized_end=14933
  _globals['_TASKRUN']._serialized_start=14946
  _globals['_TASKRUN']._serialized_end=15609
  _globals['_TASKRUN_PHASE']._serialized_start=15572
  _globals['_TASKRUN_PHASE']._serialized_end=15609
  _globals['_TASKENTITYCONTEXT']._serialized_start=15612
  _globals['_TASKENTITYCONTEXT']._serialized_end=15828
# @@protoc_insertion_point(module_scope)
