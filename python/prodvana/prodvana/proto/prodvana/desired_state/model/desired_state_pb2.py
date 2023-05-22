# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/desired_state/model/desired_state.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from prodvana.proto.validate import validate_pb2 as validate_dot_validate__pb2
from prodvana.proto.prodvana.common_config import program_pb2 as prodvana_dot_common__config_dot_program__pb2
from prodvana.proto.prodvana.common_config import retry_pb2 as prodvana_dot_common__config_dot_retry__pb2
from prodvana.proto.prodvana.protection import protection_reference_pb2 as prodvana_dot_protection_dot_protection__reference__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n0prodvana/desired_state/model/desired_state.proto\x12\x1cprodvana.desired_state.model\x1a\x1egoogle/protobuf/duration.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x17validate/validate.proto\x1a$prodvana/common_config/program.proto\x1a\"prodvana/common_config/retry.proto\x1a.prodvana/protection/protection_reference.proto\"p\n\x0eProtectionLink\x12;\n\tlifecycle\x18\x03 \x01(\x0b\x32(.prodvana.protection.ProtectionLifecycle\x12\x15\n\rattachment_id\x18\x04 \x01(\tJ\x04\x08\x01\x10\x02J\x04\x08\x02\x10\x03\"\x93\x06\n\tCondition\x12X\n\x07rc_cond\x18\x01 \x01(\x0b\x32\x45.prodvana.desired_state.model.Condition.ReleaseChannelStableConditionH\x00\x12Q\n\x0fmanual_approval\x18\x02 \x01(\x0b\x32\x36.prodvana.desired_state.model.Condition.ManualApprovalH\x00\x12\\\n\x0b\x63ustom_task\x18\x03 \x01(\x0b\x32\x45.prodvana.desired_state.model.Condition.CustomTaskSuccessfulConditionH\x00\x12\x18\n\x10\x64\x65sired_state_id\x18\x04 \x01(\t\x1a\xa7\x01\n\x1dReleaseChannelStableCondition\x12\x13\n\x0b\x61pplication\x18\x01 \x01(\t\x12\x0f\n\x07service\x18\x02 \x01(\t\x12\x12\n\nservice_id\x18\x03 \x01(\t\x12\x17\n\x0frelease_channel\x18\x04 \x01(\t\x12\x1a\n\x12release_channel_id\x18\x05 \x01(\t\x12\x17\n\x0fservice_version\x18\x06 \x01(\t\x1a\x1f\n\x0eManualApproval\x12\r\n\x05topic\x18\x01 \x01(\t\x1a\x88\x02\n\x1d\x43ustomTaskSuccessfulCondition\x12\x18\n\x10\x63ustom_task_name\x18\x01 \x01(\t\x12\x66\n\nprotection\x18\x02 \x01(\x0b\x32P.prodvana.desired_state.model.Condition.CustomTaskSuccessfulCondition.ProtectionH\x00\x1a[\n\nProtection\x12\x0c\n\x04name\x18\x01 \x01(\t\x12?\n\ttask_type\x18\x03 \x01(\x0e\x32,.prodvana.desired_state.model.CustomTaskTypeB\x08\n\x06sourceB\x0b\n\tcondition\"L\n\nIdentifier\x12\x30\n\x04type\x18\x01 \x01(\x0e\x32\".prodvana.desired_state.model.Type\x12\x0c\n\x04name\x18\x02 \x01(\t\"\xd3\x02\n\x08Metadata\x12>\n\rpreconditions\x18\x01 \x03(\x0b\x32\'.prodvana.desired_state.model.Condition\x12;\n\ninvariants\x18\x02 \x03(\x0b\x32\'.prodvana.desired_state.model.Condition\x12\x36\n\x04self\x18\x03 \x01(\x0b\x32(.prodvana.desired_state.model.Identifier\x12\x18\n\x10\x64\x65sired_state_id\x18\x04 \x01(\t\x12\x1d\n\x15root_desired_state_id\x18\x05 \x01(\t\x12\x46\n\x10protection_links\x18\x06 \x03(\x0b\x32,.prodvana.desired_state.model.ProtectionLinkJ\x04\x08\x07\x10\x08R\x0bprotections\"\xad\x01\n\x11StatusExplanation\x12\x39\n\x07subject\x18\x01 \x01(\x0b\x32(.prodvana.desired_state.model.Identifier\x12:\n\x06reason\x18\x02 \x01(\x0e\x32*.prodvana.desired_state.model.StatusReason\x12\x0f\n\x07message\x18\x03 \x01(\t\x12\x10\n\x08messages\x18\x04 \x03(\t\"\x8b\x01\n\x11\x41\x63tionExplanation\x12&\n\x02ts\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12=\n\x0b\x61\x63tion_type\x18\x02 \x01(\x0e\x32(.prodvana.desired_state.model.ActionType\x12\x0f\n\x07message\x18\x03 \x01(\t\"\xcf\x01\n\x07Version\x12\x0f\n\x07version\x18\x01 \x01(\t\x12\x10\n\x08replicas\x18\x02 \x01(\x05\x12\x1a\n\x12\x61vailable_replicas\x18\x08 \x01(\x05\x12\x32\n\x0epush_timestamp\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x0e\n\x06\x61\x63tive\x18\x05 \x01(\x08\x12\x17\n\x0ftarget_replicas\x18\x06 \x01(\x05J\x04\x08\x03\x10\x04J\x04\x08\x07\x10\x08R\tunhealthyR\x11unhealthy_reasons\"\x86\x03\n\x14ServiceInstanceState\x12\x34\n\x04meta\x18\x01 \x01(\x0b\x32&.prodvana.desired_state.model.Metadata\x12\x13\n\x0b\x61pplication\x18\x02 \x01(\t\x12\x0f\n\x07service\x18\x03 \x01(\t\x12\x17\n\x0frelease_channel\x18\x04 \x01(\t\x12\x12\n\nservice_id\x18\x08 \x01(\t\x12\x1a\n\x12release_channel_id\x18\t \x01(\t\x12\x37\n\x08versions\x18\x05 \x03(\x0b\x32%.prodvana.desired_state.model.Version\x12?\n\x10rollback_version\x18\x06 \x01(\x0b\x32%.prodvana.desired_state.model.Version\x12\x10\n\x08rollback\x18\x07 \x01(\x08\x12=\n\x08\x64\x65livery\x18\n \x01(\x0b\x32+.prodvana.desired_state.model.DeliveryState\"\x9d\x02\n\x0cServiceState\x12\x34\n\x04meta\x18\x01 \x01(\x0b\x32&.prodvana.desired_state.model.Metadata\x12\x13\n\x0b\x61pplication\x18\x02 \x01(\t\x12\x0f\n\x07service\x18\x03 \x01(\t\x12\x12\n\nservice_id\x18\x05 \x01(\t\x12L\n\x10release_channels\x18\x06 \x03(\x0b\x32\x32.prodvana.desired_state.model.ServiceInstanceState\x12\x43\n\x0c\x63ustom_tasks\x18\x07 \x03(\x0b\x32-.prodvana.desired_state.model.CustomTaskStateJ\x04\x08\x04\x10\x05J\x04\x08\x08\x10\t\"\xd2\x01\n\x11ServiceGroupState\x12\x34\n\x04meta\x18\x01 \x01(\x0b\x32&.prodvana.desired_state.model.Metadata\x12<\n\x08services\x18\x02 \x03(\x0b\x32*.prodvana.desired_state.model.ServiceState\x12\x43\n\x0c\x63ustom_tasks\x18\x03 \x03(\x0b\x32-.prodvana.desired_state.model.CustomTaskStateJ\x04\x08\x04\x10\x05\"\xa8\x02\n\x13\x43\x61naryProgressState\x12H\n\x06status\x18\x01 \x01(\x0e\x32\x38.prodvana.desired_state.model.CanaryProgressState.Status\x12 \n\rcanary_weight\x18\x02 \x01(\x05\x42\t\xfa\x42\x06\x1a\x04\x18\x64(\x00\x12+\n\x08\x64uration\x18\x03 \x01(\x0b\x32\x19.google.protobuf.Duration\x12\x39\n\x15pause_start_timestamp\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"=\n\x06Status\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x0b\n\x07PENDING\x10\x01\x12\n\n\x06PAUSED\x10\x02\x12\r\n\tCOMPLETED\x10\x03\"\x94\x03\n\rDeliveryState\x12\x18\n\x10\x64\x65sired_state_id\x18\x0c \x01(\t\x12\x42\n\x06status\x18\x08 \x01(\x0e\x32\x32.prodvana.desired_state.model.DeliveryState.Status\x12\x0f\n\x07message\x18\x06 \x01(\t\x12J\n\x0f\x63\x61nary_progress\x18\x0b \x03(\x0b\x32\x31.prodvana.desired_state.model.CanaryProgressState\x12\x11\n\tfirst_run\x18\r \x01(\x08\x12\x12\n\ngeneration\x18\x0e \x01(\t\"q\n\x06Status\x12\x12\n\x0eSTATUS_UNKNOWN\x10\x00\x12\x16\n\x12STATUS_PROGRESSING\x10\x01\x12\x11\n\rSTATUS_PAUSED\x10\x02\x12\x12\n\x0eSTATUS_HEALTHY\x10\x03\x12\x14\n\x10STATUS_UNHEALTHY\x10\x04J\x04\x08\x01\x10\x02J\x04\x08\x02\x10\x03J\x04\x08\x03\x10\x04J\x04\x08\x04\x10\x05J\x04\x08\x05\x10\x06J\x04\x08\x07\x10\x08J\x04\x08\t\x10\nJ\x04\x08\n\x10\x0b\"\xff\x03\n\rRuntimeObject\x12\x34\n\x04meta\x18\x01 \x01(\x0b\x32&.prodvana.desired_state.model.Metadata\x12\x13\n\x0bobject_type\x18\x02 \x01(\t\x12\x11\n\tnamespace\x18\x03 \x01(\t\x12\x0c\n\x04name\x18\x04 \x01(\t\x12\x37\n\x08versions\x18\x05 \x03(\x0b\x32%.prodvana.desired_state.model.Version\x12\x42\n\x06status\x18\x06 \x01(\x0e\x32\x32.prodvana.desired_state.model.RuntimeObject.Status\x12?\n\x10rollback_version\x18\x07 \x01(\x0b\x32%.prodvana.desired_state.model.Version\x12=\n\x08\x64\x65livery\x18\x08 \x01(\x0b\x32+.prodvana.desired_state.model.DeliveryState\x12\x18\n\x10version_agnostic\x18\n \x01(\x08\x12\x12\n\nstate_hash\x18\x0b \x01(\x0c\x12\x0f\n\x07message\x18\x0c \x01(\t\"0\n\x06Status\x12\x0b\n\x07PENDING\x10\x00\x12\r\n\tSUCCEEDED\x10\x01\x12\n\n\x06\x46\x41ILED\x10\x02J\x04\x08\t\x10\nR\x0eunhealthy_pods\"O\n\x0e\x43onditionState\x12=\n\x06status\x18\x01 \x01(\x0e\x32-.prodvana.desired_state.model.ConditionStatus\"\xde\x02\n\x0c\x43ontrolState\x12\x10\n\x08rollback\x18\x01 \x01(\x08\x12I\n\x13precondition_states\x18\x02 \x03(\x0b\x32,.prodvana.desired_state.model.ConditionState\x12\x46\n\x10invariant_states\x18\x03 \x03(\x0b\x32,.prodvana.desired_state.model.ConditionState\x12\x0e\n\x06paused\x18\x04 \x01(\x08\x12L\n\x13status_explanations\x18\x05 \x03(\x0b\x32/.prodvana.desired_state.model.StatusExplanation\x12K\n\x12\x61\x63tion_explanation\x18\x06 \x01(\x0b\x32/.prodvana.desired_state.model.ActionExplanation\"\x9e\x01\n\x13ManualApprovalState\x12\x34\n\x04meta\x18\x01 \x01(\x0b\x32&.prodvana.desired_state.model.Metadata\x12\x42\n\x06status\x18\x02 \x01(\x0e\x32\x32.prodvana.desired_state.model.ManualApprovalStatus\x12\r\n\x05topic\x18\x03 \x01(\t\"\xed\x04\n\x05State\x12=\n\x07service\x18\x01 \x01(\x0b\x32*.prodvana.desired_state.model.ServiceStateH\x00\x12N\n\x10service_instance\x18\x02 \x01(\x0b\x32\x32.prodvana.desired_state.model.ServiceInstanceStateH\x00\x12H\n\rservice_group\x18\x03 \x01(\x0b\x32/.prodvana.desired_state.model.ServiceGroupStateH\x00\x12\x45\n\x0eruntime_object\x18\x04 \x01(\x0b\x32+.prodvana.desired_state.model.RuntimeObjectH\x00\x12L\n\x0fmanual_approval\x18\x05 \x01(\x0b\x32\x31.prodvana.desired_state.model.ManualApprovalStateH\x00\x12\x44\n\x0b\x63ustom_task\x18\x06 \x01(\x0b\x32-.prodvana.desired_state.model.CustomTaskStateH\x00\x12S\n\x15protection_attachment\x18\x07 \x01(\x0b\x32\x32.prodvana.desired_state.model.ProtectionAttachmentH\x00\x12L\n\x0fprotection_link\x18\x08 \x01(\x0b\x32\x31.prodvana.desired_state.model.ProtectionLinkStateH\x00\x42\r\n\x0bstate_oneof\"\x82\x01\n\x0b\x41nnotations\x12I\n\x0b\x61nnotations\x18\x01 \x03(\x0b\x32\x34.prodvana.desired_state.model.Annotations.Annotation\x1a(\n\nAnnotation\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t\"\xae\x01\n\x18\x43ustomTaskExecutionState\x12>\n\x06status\x18\x01 \x01(\x0e\x32..prodvana.desired_state.model.CustomTaskStatus\x12\x10\n\x08\x61ttempts\x18\x02 \x01(\x03\x12@\n\x1clatest_attempt_end_timestamp\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\xce\x03\n\x0f\x43ustomTaskState\x12\x34\n\x04meta\x18\x01 \x01(\x0b\x32&.prodvana.desired_state.model.Metadata\x12\x15\n\x04name\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x1c\n\x0b\x64\x65scription\x18\x03 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x13\n\x0b\x61pplication\x18\x04 \x01(\t\x12\x16\n\x0e\x61pplication_id\x18\x05 \x01(\t\x12\x17\n\x0frelease_channel\x18\x06 \x01(\t\x12\x1a\n\x12release_channel_id\x18\x07 \x01(\t\x12@\n\x07program\x18\t \x01(\x0b\x32%.prodvana.common_config.ProgramConfigB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12J\n\ntask_state\x18\x0c \x01(\x0b\x32\x36.prodvana.desired_state.model.CustomTaskExecutionState\x12\x39\n\x0cretry_config\x18\r \x01(\x0b\x32#.prodvana.common_config.RetryConfig\x12\x13\n\x0bservice_ids\x18\x0e \x03(\tJ\x04\x08\x08\x10\tJ\x04\x08\n\x10\x0bJ\x04\x08\x0b\x10\x0c\"\xd4\x04\n\x13ProtectionLinkState\x12\x34\n\x04meta\x18\x01 \x01(\x0b\x32&.prodvana.desired_state.model.Metadata\x12=\n\x06status\x18\x02 \x01(\x0e\x32-.prodvana.desired_state.model.ConditionStatus\x12:\n\x04link\x18\x03 \x01(\x0b\x32,.prodvana.desired_state.model.ProtectionLink\x12\x35\n\x11started_timestamp\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x35\n\x11stopped_timestamp\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12T\n\x0estopped_reason\x18\x06 \x01(\x0e\x32<.prodvana.desired_state.model.ProtectionLinkState.StopReason\x12;\n\x17\x66irst_success_timestamp\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\x8a\x01\n\nStopReason\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x17\n\x13LIFECYCLE_COMPLETED\x10\x01\x12\x12\n\x0eSUCCEEDED_ONCE\x10\x02\x12\x1a\n\x16SUCCEEDED_FOR_DURATION\x10\x03\x12\r\n\tTIMED_OUT\x10\x04\x12\n\n\x06\x46\x41ILED\x10\x05\x12\x0b\n\x07\x44\x45LETED\x10\x06\"\x89\x04\n\x14ProtectionAttachment\x12\x34\n\x04meta\x18\x01 \x01(\x0b\x32&.prodvana.desired_state.model.Metadata\x12\x37\n\x08versions\x18\x02 \x03(\x0b\x32%.prodvana.desired_state.model.Version\x12\x46\n\x17last_completed_versions\x18\x05 \x03(\x0b\x32%.prodvana.desired_state.model.Version\x12<\n\x18last_completed_timestamp\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12I\n\x15last_completed_status\x18\x07 \x01(\x0e\x32*.prodvana.desired_state.model.SimpleStatus\x12[\n\"last_completed_status_explanations\x18\x08 \x03(\x0b\x32/.prodvana.desired_state.model.StatusExplanation\x12&\n\x1elast_completed_applied_version\x18\t \x01(\t\x12\x15\n\rprotection_id\x18\x03 \x01(\t\x12\x15\n\rattachment_id\x18\x04 \x01(\t\"\xde\x01\n\x06Signal\x12\x36\n\x04type\x18\x01 \x01(\x0e\x32(.prodvana.desired_state.model.SignalType\x12Z\n\x12\x64\x65livery_promotion\x18\x02 \x01(\x0b\x32<.prodvana.desired_state.model.Signal.DeliveryPromotionConfigH\x00\x1a\x36\n\x17\x44\x65liveryPromotionConfig\x12\r\n\x05stage\x18\x01 \x01(\x03\x12\x0c\n\x04\x66ull\x18\x02 \x01(\x08\x42\x08\n\x06\x63onfig\"?\n\x08\x44\x65\x62ugLog\x12&\n\x02ts\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x0b\n\x03log\x18\x02 \x01(\t*\xb3\x01\n\x04Type\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x0b\n\x07SERVICE\x10\x01\x12\x14\n\x10SERVICE_INSTANCE\x10\x02\x12\x11\n\rSERVICE_GROUP\x10\x03\x12\x12\n\x0eRUNTIME_OBJECT\x10\x04\x12\x13\n\x0fMANUAL_APPROVAL\x10\x05\x12\x0f\n\x0b\x43USTOM_TASK\x10\x06\x12\x19\n\x15PROTECTION_ATTACHMENT\x10\x07\x12\x13\n\x0fPROTECTION_LINK\x10\x08*w\n\tLifecycle\x12\x15\n\x11UNKNOWN_LIFECYCLE\x10\x00\x12\x15\n\x11\x43ONVERGENCE_START\x10\x01\x12\x10\n\x0cPRE_APPROVAL\x10\x02\x12\x11\n\rPOST_APPROVAL\x10\x03\x12\x08\n\x04PUSH\x10\x04\x12\r\n\tPOST_PUSH\x10\x05*\x8e\x01\n\x0e\x43ustomTaskType\x12\x1c\n\x18\x43USTOM_TASK_TYPE_UNKNOWN\x10\x00\x12\x15\n\x11PRE_APPROVAL_TASK\x10\x01\x12\x0c\n\x08\x41PPROVAL\x10\x02\x12\x16\n\x12POST_APPROVAL_TASK\x10\x03\x12\r\n\tPUSH_TASK\x10\x04\x12\x12\n\x0ePOST_PUSH_TASK\x10\x05*\xf0\x01\n\x06Status\x12\x12\n\x0eUNKNOWN_STATUS\x10\x00\x12\x0e\n\nCONVERGING\x10\x01\x12\r\n\tCONVERGED\x10\x02\x12\n\n\x06\x46\x41ILED\x10\x03\x12\x10\n\x0cROLLING_BACK\x10\x04\x12\x0f\n\x0bROLLED_BACK\x10\x05\x12\x13\n\x0f\x46\x41ILED_ROLLBACK\x10\x0c\x12\n\n\x06PAUSED\x10\x06\x12\x10\n\x0c\x43HILD_PAUSED\x10\x07\x12\x19\n\x15WAITING_PRECONDITIONS\x10\x08\x12\x0c\n\x08REPLACED\x10\t\x12\x1b\n\x17WAITING_MANUAL_APPROVAL\x10\n\x12\x0b\n\x07\x44\x45LETED\x10\x0b*R\n\x0cSimpleStatus\x12\x0e\n\nSS_UNKNOWN\x10\x00\x12\x11\n\rSS_CONVERGING\x10\x01\x12\x10\n\x0cSS_CONVERGED\x10\x02\x12\r\n\tSS_FAILED\x10\x03*\xec\x01\n\x0cStatusReason\x12\x12\n\x0eREASON_UNKNOWN\x10\x00\x12\x14\n\x10NO_CURRENT_STATE\x10\x01\x12\x10\n\x0c\x41PPLY_FAILED\x10\x02\x12\x12\n\x0eUNHEALTHY_PODS\x10\x03\x12\x11\n\rUPDATING_PODS\x10\x04\x12\x14\n\x10VERSION_MISMATCH\x10\x05\x12\x19\n\x15RUNTIME_OBJECT_FAILED\x10\x06\x12\x18\n\x14PRECONDITIONS_FAILED\x10\x07\x12\x1c\n\x18MANUAL_APPROVAL_REJECTED\x10\x08\x12\x10\n\x0cSTUCK_ENTITY\x10\t*r\n\nActionType\x12\x17\n\x13\x41\x43TION_TYPE_UNKNOWN\x10\x00\x12\x18\n\x14\x41\x43TION_TYPE_APPLYING\x10\x01\x12\x17\n\x13\x41\x43TION_TYPE_APPLIED\x10\x02\x12\x18\n\x14\x41\x43TION_TYPE_COMPLETE\x10\x03*\x96\x01\n\x0f\x43onditionStatus\x12\x1c\n\x18\x43ONDITION_UNKNOWN_STATUS\x10\x00\x12\x15\n\x11\x43ONDITION_PENDING\x10\x01\x12\x17\n\x13\x43ONDITION_SATISFIED\x10\x02\x12\x1f\n\x1b\x43ONDITION_MANUALLY_BYPASSED\x10\x03\x12\x14\n\x10\x43ONDITION_FAILED\x10\x04*?\n\x14ManualApprovalStatus\x12\x0b\n\x07PENDING\x10\x00\x12\x0c\n\x08\x41PPROVED\x10\x01\x12\x0c\n\x08REJECTED\x10\x02*\x85\x01\n\x10\x43ustomTaskStatus\x12\x17\n\x13\x43USTOM_TASK_PENDING\x10\x00\x12\x1a\n\x16\x43USTOM_TASK_SUCCESSFUL\x10\x01\x12!\n\x1d\x43USTOM_TASK_RETRIES_EXHAUSTED\x10\x02\x12\x19\n\x15\x43USTOM_TASK_TIMED_OUT\x10\x03*8\n\nSignalType\x12\x12\n\x0eSIGNAL_UNKNOWN\x10\x00\x12\x16\n\x12\x44\x45LIVERY_PROMOTION\x10\x01\x42XZVgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/modelb\x06proto3')

_TYPE = DESCRIPTOR.enum_types_by_name['Type']
Type = enum_type_wrapper.EnumTypeWrapper(_TYPE)
_LIFECYCLE = DESCRIPTOR.enum_types_by_name['Lifecycle']
Lifecycle = enum_type_wrapper.EnumTypeWrapper(_LIFECYCLE)
_CUSTOMTASKTYPE = DESCRIPTOR.enum_types_by_name['CustomTaskType']
CustomTaskType = enum_type_wrapper.EnumTypeWrapper(_CUSTOMTASKTYPE)
_STATUS = DESCRIPTOR.enum_types_by_name['Status']
Status = enum_type_wrapper.EnumTypeWrapper(_STATUS)
_SIMPLESTATUS = DESCRIPTOR.enum_types_by_name['SimpleStatus']
SimpleStatus = enum_type_wrapper.EnumTypeWrapper(_SIMPLESTATUS)
_STATUSREASON = DESCRIPTOR.enum_types_by_name['StatusReason']
StatusReason = enum_type_wrapper.EnumTypeWrapper(_STATUSREASON)
_ACTIONTYPE = DESCRIPTOR.enum_types_by_name['ActionType']
ActionType = enum_type_wrapper.EnumTypeWrapper(_ACTIONTYPE)
_CONDITIONSTATUS = DESCRIPTOR.enum_types_by_name['ConditionStatus']
ConditionStatus = enum_type_wrapper.EnumTypeWrapper(_CONDITIONSTATUS)
_MANUALAPPROVALSTATUS = DESCRIPTOR.enum_types_by_name['ManualApprovalStatus']
ManualApprovalStatus = enum_type_wrapper.EnumTypeWrapper(_MANUALAPPROVALSTATUS)
_CUSTOMTASKSTATUS = DESCRIPTOR.enum_types_by_name['CustomTaskStatus']
CustomTaskStatus = enum_type_wrapper.EnumTypeWrapper(_CUSTOMTASKSTATUS)
_SIGNALTYPE = DESCRIPTOR.enum_types_by_name['SignalType']
SignalType = enum_type_wrapper.EnumTypeWrapper(_SIGNALTYPE)
UNKNOWN = 0
SERVICE = 1
SERVICE_INSTANCE = 2
SERVICE_GROUP = 3
RUNTIME_OBJECT = 4
MANUAL_APPROVAL = 5
CUSTOM_TASK = 6
PROTECTION_ATTACHMENT = 7
PROTECTION_LINK = 8
UNKNOWN_LIFECYCLE = 0
CONVERGENCE_START = 1
PRE_APPROVAL = 2
POST_APPROVAL = 3
PUSH = 4
POST_PUSH = 5
CUSTOM_TASK_TYPE_UNKNOWN = 0
PRE_APPROVAL_TASK = 1
APPROVAL = 2
POST_APPROVAL_TASK = 3
PUSH_TASK = 4
POST_PUSH_TASK = 5
UNKNOWN_STATUS = 0
CONVERGING = 1
CONVERGED = 2
FAILED = 3
ROLLING_BACK = 4
ROLLED_BACK = 5
FAILED_ROLLBACK = 12
PAUSED = 6
CHILD_PAUSED = 7
WAITING_PRECONDITIONS = 8
REPLACED = 9
WAITING_MANUAL_APPROVAL = 10
DELETED = 11
SS_UNKNOWN = 0
SS_CONVERGING = 1
SS_CONVERGED = 2
SS_FAILED = 3
REASON_UNKNOWN = 0
NO_CURRENT_STATE = 1
APPLY_FAILED = 2
UNHEALTHY_PODS = 3
UPDATING_PODS = 4
VERSION_MISMATCH = 5
RUNTIME_OBJECT_FAILED = 6
PRECONDITIONS_FAILED = 7
MANUAL_APPROVAL_REJECTED = 8
STUCK_ENTITY = 9
ACTION_TYPE_UNKNOWN = 0
ACTION_TYPE_APPLYING = 1
ACTION_TYPE_APPLIED = 2
ACTION_TYPE_COMPLETE = 3
CONDITION_UNKNOWN_STATUS = 0
CONDITION_PENDING = 1
CONDITION_SATISFIED = 2
CONDITION_MANUALLY_BYPASSED = 3
CONDITION_FAILED = 4
PENDING = 0
APPROVED = 1
REJECTED = 2
CUSTOM_TASK_PENDING = 0
CUSTOM_TASK_SUCCESSFUL = 1
CUSTOM_TASK_RETRIES_EXHAUSTED = 2
CUSTOM_TASK_TIMED_OUT = 3
SIGNAL_UNKNOWN = 0
DELIVERY_PROMOTION = 1


_PROTECTIONLINK = DESCRIPTOR.message_types_by_name['ProtectionLink']
_CONDITION = DESCRIPTOR.message_types_by_name['Condition']
_CONDITION_RELEASECHANNELSTABLECONDITION = _CONDITION.nested_types_by_name['ReleaseChannelStableCondition']
_CONDITION_MANUALAPPROVAL = _CONDITION.nested_types_by_name['ManualApproval']
_CONDITION_CUSTOMTASKSUCCESSFULCONDITION = _CONDITION.nested_types_by_name['CustomTaskSuccessfulCondition']
_CONDITION_CUSTOMTASKSUCCESSFULCONDITION_PROTECTION = _CONDITION_CUSTOMTASKSUCCESSFULCONDITION.nested_types_by_name['Protection']
_IDENTIFIER = DESCRIPTOR.message_types_by_name['Identifier']
_METADATA = DESCRIPTOR.message_types_by_name['Metadata']
_STATUSEXPLANATION = DESCRIPTOR.message_types_by_name['StatusExplanation']
_ACTIONEXPLANATION = DESCRIPTOR.message_types_by_name['ActionExplanation']
_VERSION = DESCRIPTOR.message_types_by_name['Version']
_SERVICEINSTANCESTATE = DESCRIPTOR.message_types_by_name['ServiceInstanceState']
_SERVICESTATE = DESCRIPTOR.message_types_by_name['ServiceState']
_SERVICEGROUPSTATE = DESCRIPTOR.message_types_by_name['ServiceGroupState']
_CANARYPROGRESSSTATE = DESCRIPTOR.message_types_by_name['CanaryProgressState']
_DELIVERYSTATE = DESCRIPTOR.message_types_by_name['DeliveryState']
_RUNTIMEOBJECT = DESCRIPTOR.message_types_by_name['RuntimeObject']
_CONDITIONSTATE = DESCRIPTOR.message_types_by_name['ConditionState']
_CONTROLSTATE = DESCRIPTOR.message_types_by_name['ControlState']
_MANUALAPPROVALSTATE = DESCRIPTOR.message_types_by_name['ManualApprovalState']
_STATE = DESCRIPTOR.message_types_by_name['State']
_ANNOTATIONS = DESCRIPTOR.message_types_by_name['Annotations']
_ANNOTATIONS_ANNOTATION = _ANNOTATIONS.nested_types_by_name['Annotation']
_CUSTOMTASKEXECUTIONSTATE = DESCRIPTOR.message_types_by_name['CustomTaskExecutionState']
_CUSTOMTASKSTATE = DESCRIPTOR.message_types_by_name['CustomTaskState']
_PROTECTIONLINKSTATE = DESCRIPTOR.message_types_by_name['ProtectionLinkState']
_PROTECTIONATTACHMENT = DESCRIPTOR.message_types_by_name['ProtectionAttachment']
_SIGNAL = DESCRIPTOR.message_types_by_name['Signal']
_SIGNAL_DELIVERYPROMOTIONCONFIG = _SIGNAL.nested_types_by_name['DeliveryPromotionConfig']
_DEBUGLOG = DESCRIPTOR.message_types_by_name['DebugLog']
_CANARYPROGRESSSTATE_STATUS = _CANARYPROGRESSSTATE.enum_types_by_name['Status']
_DELIVERYSTATE_STATUS = _DELIVERYSTATE.enum_types_by_name['Status']
_RUNTIMEOBJECT_STATUS = _RUNTIMEOBJECT.enum_types_by_name['Status']
_PROTECTIONLINKSTATE_STOPREASON = _PROTECTIONLINKSTATE.enum_types_by_name['StopReason']
ProtectionLink = _reflection.GeneratedProtocolMessageType('ProtectionLink', (_message.Message,), {
  'DESCRIPTOR' : _PROTECTIONLINK,
  '__module__' : 'prodvana.desired_state.model.desired_state_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.desired_state.model.ProtectionLink)
  })
_sym_db.RegisterMessage(ProtectionLink)

Condition = _reflection.GeneratedProtocolMessageType('Condition', (_message.Message,), {

  'ReleaseChannelStableCondition' : _reflection.GeneratedProtocolMessageType('ReleaseChannelStableCondition', (_message.Message,), {
    'DESCRIPTOR' : _CONDITION_RELEASECHANNELSTABLECONDITION,
    '__module__' : 'prodvana.desired_state.model.desired_state_pb2'
    # @@protoc_insertion_point(class_scope:prodvana.desired_state.model.Condition.ReleaseChannelStableCondition)
    })
  ,

  'ManualApproval' : _reflection.GeneratedProtocolMessageType('ManualApproval', (_message.Message,), {
    'DESCRIPTOR' : _CONDITION_MANUALAPPROVAL,
    '__module__' : 'prodvana.desired_state.model.desired_state_pb2'
    # @@protoc_insertion_point(class_scope:prodvana.desired_state.model.Condition.ManualApproval)
    })
  ,

  'CustomTaskSuccessfulCondition' : _reflection.GeneratedProtocolMessageType('CustomTaskSuccessfulCondition', (_message.Message,), {

    'Protection' : _reflection.GeneratedProtocolMessageType('Protection', (_message.Message,), {
      'DESCRIPTOR' : _CONDITION_CUSTOMTASKSUCCESSFULCONDITION_PROTECTION,
      '__module__' : 'prodvana.desired_state.model.desired_state_pb2'
      # @@protoc_insertion_point(class_scope:prodvana.desired_state.model.Condition.CustomTaskSuccessfulCondition.Protection)
      })
    ,
    'DESCRIPTOR' : _CONDITION_CUSTOMTASKSUCCESSFULCONDITION,
    '__module__' : 'prodvana.desired_state.model.desired_state_pb2'
    # @@protoc_insertion_point(class_scope:prodvana.desired_state.model.Condition.CustomTaskSuccessfulCondition)
    })
  ,
  'DESCRIPTOR' : _CONDITION,
  '__module__' : 'prodvana.desired_state.model.desired_state_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.desired_state.model.Condition)
  })
_sym_db.RegisterMessage(Condition)
_sym_db.RegisterMessage(Condition.ReleaseChannelStableCondition)
_sym_db.RegisterMessage(Condition.ManualApproval)
_sym_db.RegisterMessage(Condition.CustomTaskSuccessfulCondition)
_sym_db.RegisterMessage(Condition.CustomTaskSuccessfulCondition.Protection)

Identifier = _reflection.GeneratedProtocolMessageType('Identifier', (_message.Message,), {
  'DESCRIPTOR' : _IDENTIFIER,
  '__module__' : 'prodvana.desired_state.model.desired_state_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.desired_state.model.Identifier)
  })
_sym_db.RegisterMessage(Identifier)

Metadata = _reflection.GeneratedProtocolMessageType('Metadata', (_message.Message,), {
  'DESCRIPTOR' : _METADATA,
  '__module__' : 'prodvana.desired_state.model.desired_state_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.desired_state.model.Metadata)
  })
_sym_db.RegisterMessage(Metadata)

StatusExplanation = _reflection.GeneratedProtocolMessageType('StatusExplanation', (_message.Message,), {
  'DESCRIPTOR' : _STATUSEXPLANATION,
  '__module__' : 'prodvana.desired_state.model.desired_state_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.desired_state.model.StatusExplanation)
  })
_sym_db.RegisterMessage(StatusExplanation)

ActionExplanation = _reflection.GeneratedProtocolMessageType('ActionExplanation', (_message.Message,), {
  'DESCRIPTOR' : _ACTIONEXPLANATION,
  '__module__' : 'prodvana.desired_state.model.desired_state_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.desired_state.model.ActionExplanation)
  })
_sym_db.RegisterMessage(ActionExplanation)

Version = _reflection.GeneratedProtocolMessageType('Version', (_message.Message,), {
  'DESCRIPTOR' : _VERSION,
  '__module__' : 'prodvana.desired_state.model.desired_state_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.desired_state.model.Version)
  })
_sym_db.RegisterMessage(Version)

ServiceInstanceState = _reflection.GeneratedProtocolMessageType('ServiceInstanceState', (_message.Message,), {
  'DESCRIPTOR' : _SERVICEINSTANCESTATE,
  '__module__' : 'prodvana.desired_state.model.desired_state_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.desired_state.model.ServiceInstanceState)
  })
_sym_db.RegisterMessage(ServiceInstanceState)

ServiceState = _reflection.GeneratedProtocolMessageType('ServiceState', (_message.Message,), {
  'DESCRIPTOR' : _SERVICESTATE,
  '__module__' : 'prodvana.desired_state.model.desired_state_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.desired_state.model.ServiceState)
  })
_sym_db.RegisterMessage(ServiceState)

ServiceGroupState = _reflection.GeneratedProtocolMessageType('ServiceGroupState', (_message.Message,), {
  'DESCRIPTOR' : _SERVICEGROUPSTATE,
  '__module__' : 'prodvana.desired_state.model.desired_state_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.desired_state.model.ServiceGroupState)
  })
_sym_db.RegisterMessage(ServiceGroupState)

CanaryProgressState = _reflection.GeneratedProtocolMessageType('CanaryProgressState', (_message.Message,), {
  'DESCRIPTOR' : _CANARYPROGRESSSTATE,
  '__module__' : 'prodvana.desired_state.model.desired_state_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.desired_state.model.CanaryProgressState)
  })
_sym_db.RegisterMessage(CanaryProgressState)

DeliveryState = _reflection.GeneratedProtocolMessageType('DeliveryState', (_message.Message,), {
  'DESCRIPTOR' : _DELIVERYSTATE,
  '__module__' : 'prodvana.desired_state.model.desired_state_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.desired_state.model.DeliveryState)
  })
_sym_db.RegisterMessage(DeliveryState)

RuntimeObject = _reflection.GeneratedProtocolMessageType('RuntimeObject', (_message.Message,), {
  'DESCRIPTOR' : _RUNTIMEOBJECT,
  '__module__' : 'prodvana.desired_state.model.desired_state_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.desired_state.model.RuntimeObject)
  })
_sym_db.RegisterMessage(RuntimeObject)

ConditionState = _reflection.GeneratedProtocolMessageType('ConditionState', (_message.Message,), {
  'DESCRIPTOR' : _CONDITIONSTATE,
  '__module__' : 'prodvana.desired_state.model.desired_state_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.desired_state.model.ConditionState)
  })
_sym_db.RegisterMessage(ConditionState)

ControlState = _reflection.GeneratedProtocolMessageType('ControlState', (_message.Message,), {
  'DESCRIPTOR' : _CONTROLSTATE,
  '__module__' : 'prodvana.desired_state.model.desired_state_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.desired_state.model.ControlState)
  })
_sym_db.RegisterMessage(ControlState)

ManualApprovalState = _reflection.GeneratedProtocolMessageType('ManualApprovalState', (_message.Message,), {
  'DESCRIPTOR' : _MANUALAPPROVALSTATE,
  '__module__' : 'prodvana.desired_state.model.desired_state_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.desired_state.model.ManualApprovalState)
  })
_sym_db.RegisterMessage(ManualApprovalState)

State = _reflection.GeneratedProtocolMessageType('State', (_message.Message,), {
  'DESCRIPTOR' : _STATE,
  '__module__' : 'prodvana.desired_state.model.desired_state_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.desired_state.model.State)
  })
_sym_db.RegisterMessage(State)

Annotations = _reflection.GeneratedProtocolMessageType('Annotations', (_message.Message,), {

  'Annotation' : _reflection.GeneratedProtocolMessageType('Annotation', (_message.Message,), {
    'DESCRIPTOR' : _ANNOTATIONS_ANNOTATION,
    '__module__' : 'prodvana.desired_state.model.desired_state_pb2'
    # @@protoc_insertion_point(class_scope:prodvana.desired_state.model.Annotations.Annotation)
    })
  ,
  'DESCRIPTOR' : _ANNOTATIONS,
  '__module__' : 'prodvana.desired_state.model.desired_state_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.desired_state.model.Annotations)
  })
_sym_db.RegisterMessage(Annotations)
_sym_db.RegisterMessage(Annotations.Annotation)

CustomTaskExecutionState = _reflection.GeneratedProtocolMessageType('CustomTaskExecutionState', (_message.Message,), {
  'DESCRIPTOR' : _CUSTOMTASKEXECUTIONSTATE,
  '__module__' : 'prodvana.desired_state.model.desired_state_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.desired_state.model.CustomTaskExecutionState)
  })
_sym_db.RegisterMessage(CustomTaskExecutionState)

CustomTaskState = _reflection.GeneratedProtocolMessageType('CustomTaskState', (_message.Message,), {
  'DESCRIPTOR' : _CUSTOMTASKSTATE,
  '__module__' : 'prodvana.desired_state.model.desired_state_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.desired_state.model.CustomTaskState)
  })
_sym_db.RegisterMessage(CustomTaskState)

ProtectionLinkState = _reflection.GeneratedProtocolMessageType('ProtectionLinkState', (_message.Message,), {
  'DESCRIPTOR' : _PROTECTIONLINKSTATE,
  '__module__' : 'prodvana.desired_state.model.desired_state_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.desired_state.model.ProtectionLinkState)
  })
_sym_db.RegisterMessage(ProtectionLinkState)

ProtectionAttachment = _reflection.GeneratedProtocolMessageType('ProtectionAttachment', (_message.Message,), {
  'DESCRIPTOR' : _PROTECTIONATTACHMENT,
  '__module__' : 'prodvana.desired_state.model.desired_state_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.desired_state.model.ProtectionAttachment)
  })
_sym_db.RegisterMessage(ProtectionAttachment)

Signal = _reflection.GeneratedProtocolMessageType('Signal', (_message.Message,), {

  'DeliveryPromotionConfig' : _reflection.GeneratedProtocolMessageType('DeliveryPromotionConfig', (_message.Message,), {
    'DESCRIPTOR' : _SIGNAL_DELIVERYPROMOTIONCONFIG,
    '__module__' : 'prodvana.desired_state.model.desired_state_pb2'
    # @@protoc_insertion_point(class_scope:prodvana.desired_state.model.Signal.DeliveryPromotionConfig)
    })
  ,
  'DESCRIPTOR' : _SIGNAL,
  '__module__' : 'prodvana.desired_state.model.desired_state_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.desired_state.model.Signal)
  })
_sym_db.RegisterMessage(Signal)
_sym_db.RegisterMessage(Signal.DeliveryPromotionConfig)

DebugLog = _reflection.GeneratedProtocolMessageType('DebugLog', (_message.Message,), {
  'DESCRIPTOR' : _DEBUGLOG,
  '__module__' : 'prodvana.desired_state.model.desired_state_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.desired_state.model.DebugLog)
  })
_sym_db.RegisterMessage(DebugLog)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZVgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/model'
  _CANARYPROGRESSSTATE.fields_by_name['canary_weight']._options = None
  _CANARYPROGRESSSTATE.fields_by_name['canary_weight']._serialized_options = b'\372B\006\032\004\030d(\000'
  _CUSTOMTASKSTATE.fields_by_name['name']._options = None
  _CUSTOMTASKSTATE.fields_by_name['name']._serialized_options = b'\372B\004r\002\020\001'
  _CUSTOMTASKSTATE.fields_by_name['description']._options = None
  _CUSTOMTASKSTATE.fields_by_name['description']._serialized_options = b'\372B\004r\002\020\001'
  _CUSTOMTASKSTATE.fields_by_name['program']._options = None
  _CUSTOMTASKSTATE.fields_by_name['program']._serialized_options = b'\372B\005\212\001\002\020\001'
  _TYPE._serialized_start=7668
  _TYPE._serialized_end=7847
  _LIFECYCLE._serialized_start=7849
  _LIFECYCLE._serialized_end=7968
  _CUSTOMTASKTYPE._serialized_start=7971
  _CUSTOMTASKTYPE._serialized_end=8113
  _STATUS._serialized_start=8116
  _STATUS._serialized_end=8356
  _SIMPLESTATUS._serialized_start=8358
  _SIMPLESTATUS._serialized_end=8440
  _STATUSREASON._serialized_start=8443
  _STATUSREASON._serialized_end=8679
  _ACTIONTYPE._serialized_start=8681
  _ACTIONTYPE._serialized_end=8795
  _CONDITIONSTATUS._serialized_start=8798
  _CONDITIONSTATUS._serialized_end=8948
  _MANUALAPPROVALSTATUS._serialized_start=8950
  _MANUALAPPROVALSTATUS._serialized_end=9013
  _CUSTOMTASKSTATUS._serialized_start=9016
  _CUSTOMTASKSTATUS._serialized_end=9149
  _SIGNALTYPE._serialized_start=9151
  _SIGNALTYPE._serialized_end=9207
  _PROTECTIONLINK._serialized_start=294
  _PROTECTIONLINK._serialized_end=406
  _CONDITION._serialized_start=409
  _CONDITION._serialized_end=1196
  _CONDITION_RELEASECHANNELSTABLECONDITION._serialized_start=716
  _CONDITION_RELEASECHANNELSTABLECONDITION._serialized_end=883
  _CONDITION_MANUALAPPROVAL._serialized_start=885
  _CONDITION_MANUALAPPROVAL._serialized_end=916
  _CONDITION_CUSTOMTASKSUCCESSFULCONDITION._serialized_start=919
  _CONDITION_CUSTOMTASKSUCCESSFULCONDITION._serialized_end=1183
  _CONDITION_CUSTOMTASKSUCCESSFULCONDITION_PROTECTION._serialized_start=1082
  _CONDITION_CUSTOMTASKSUCCESSFULCONDITION_PROTECTION._serialized_end=1173
  _IDENTIFIER._serialized_start=1198
  _IDENTIFIER._serialized_end=1274
  _METADATA._serialized_start=1277
  _METADATA._serialized_end=1616
  _STATUSEXPLANATION._serialized_start=1619
  _STATUSEXPLANATION._serialized_end=1792
  _ACTIONEXPLANATION._serialized_start=1795
  _ACTIONEXPLANATION._serialized_end=1934
  _VERSION._serialized_start=1937
  _VERSION._serialized_end=2144
  _SERVICEINSTANCESTATE._serialized_start=2147
  _SERVICEINSTANCESTATE._serialized_end=2537
  _SERVICESTATE._serialized_start=2540
  _SERVICESTATE._serialized_end=2825
  _SERVICEGROUPSTATE._serialized_start=2828
  _SERVICEGROUPSTATE._serialized_end=3038
  _CANARYPROGRESSSTATE._serialized_start=3041
  _CANARYPROGRESSSTATE._serialized_end=3337
  _CANARYPROGRESSSTATE_STATUS._serialized_start=3276
  _CANARYPROGRESSSTATE_STATUS._serialized_end=3337
  _DELIVERYSTATE._serialized_start=3340
  _DELIVERYSTATE._serialized_end=3744
  _DELIVERYSTATE_STATUS._serialized_start=3583
  _DELIVERYSTATE_STATUS._serialized_end=3696
  _RUNTIMEOBJECT._serialized_start=3747
  _RUNTIMEOBJECT._serialized_end=4258
  _RUNTIMEOBJECT_STATUS._serialized_start=4188
  _RUNTIMEOBJECT_STATUS._serialized_end=4236
  _CONDITIONSTATE._serialized_start=4260
  _CONDITIONSTATE._serialized_end=4339
  _CONTROLSTATE._serialized_start=4342
  _CONTROLSTATE._serialized_end=4692
  _MANUALAPPROVALSTATE._serialized_start=4695
  _MANUALAPPROVALSTATE._serialized_end=4853
  _STATE._serialized_start=4856
  _STATE._serialized_end=5477
  _ANNOTATIONS._serialized_start=5480
  _ANNOTATIONS._serialized_end=5610
  _ANNOTATIONS_ANNOTATION._serialized_start=5570
  _ANNOTATIONS_ANNOTATION._serialized_end=5610
  _CUSTOMTASKEXECUTIONSTATE._serialized_start=5613
  _CUSTOMTASKEXECUTIONSTATE._serialized_end=5787
  _CUSTOMTASKSTATE._serialized_start=5790
  _CUSTOMTASKSTATE._serialized_end=6252
  _PROTECTIONLINKSTATE._serialized_start=6255
  _PROTECTIONLINKSTATE._serialized_end=6851
  _PROTECTIONLINKSTATE_STOPREASON._serialized_start=6713
  _PROTECTIONLINKSTATE_STOPREASON._serialized_end=6851
  _PROTECTIONATTACHMENT._serialized_start=6854
  _PROTECTIONATTACHMENT._serialized_end=7375
  _SIGNAL._serialized_start=7378
  _SIGNAL._serialized_end=7600
  _SIGNAL_DELIVERYPROMOTIONCONFIG._serialized_start=7536
  _SIGNAL_DELIVERYPROMOTIONCONFIG._serialized_end=7590
  _DEBUGLOG._serialized_start=7602
  _DEBUGLOG._serialized_end=7665
# @@protoc_insertion_point(module_scope)
