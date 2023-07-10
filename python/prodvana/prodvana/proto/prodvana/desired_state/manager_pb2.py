# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/desired_state/manager.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from prodvana.proto.prodvana.desired_state.model import desired_state_pb2 as prodvana_dot_desired__state_dot_model_dot_desired__state__pb2
from prodvana.proto.prodvana.desired_state.model import entity_pb2 as prodvana_dot_desired__state_dot_model_dot_entity__pb2
from prodvana.proto.prodvana.service import service_config_pb2 as prodvana_dot_service_dot_service__config__pb2
from prodvana.proto.validate import validate_pb2 as validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n$prodvana/desired_state/manager.proto\x12\x16prodvana.desired_state\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x30prodvana/desired_state/model/desired_state.proto\x1a)prodvana/desired_state/model/entity.proto\x1a%prodvana/service/service_config.proto\x1a\x17validate/validate.proto\"l\n\x12SetDesiredStateReq\x12\x44\n\rdesired_state\x18\x01 \x01(\x0b\x32#.prodvana.desired_state.model.StateB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12\x10\n\x08rollback\x18\x02 \x01(\x08\"\xc4\x01\n\x17ValidateDesiredStateReq\x12\x44\n\rdesired_state\x18\x01 \x01(\x0b\x32#.prodvana.desired_state.model.StateB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12\x10\n\x08rollback\x18\x02 \x01(\x08\x12Q\n\x18service_instance_configs\x18\x03 \x03(\x0b\x32/.prodvana.service.CompiledServiceInstanceConfig\"/\n\x13SetDesiredStateResp\x12\x18\n\x10\x64\x65sired_state_id\x18\x01 \x01(\t\"e\n+GetServiceDesiredStateConvergenceSummaryReq\x12\x1c\n\x0b\x61pplication\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x18\n\x07service\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\"b\n\x12StatusExplanations\x12L\n\x13status_explanations\x18\x01 \x03(\x0b\x32/.prodvana.desired_state.model.StatusExplanation\"G\n\tDebugLogs\x12:\n\ndebug_logs\x18\x01 \x03(\x0b\x32&.prodvana.desired_state.model.DebugLog\"\xaf\x0f\n\x13\x44\x65siredStateSummary\x12?\n\x0c\x65ntity_graph\x18\x0f \x01(\x0b\x32).prodvana.desired_state.model.EntityGraph\x12\x36\n\x12\x63reation_timestamp\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x36\n\x12replaced_timestamp\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12;\n\x0estarting_state\x18\x01 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12<\n\x0flast_seen_state\x18\x02 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12:\n\rdesired_state\x18\x03 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12K\n\x08statuses\x18\x04 \x03(\x0b\x32\x39.prodvana.desired_state.DesiredStateSummary.StatusesEntry\x12\x39\n\x15last_update_timestamp\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x64\n\x15precondition_statuses\x18\x08 \x03(\x0b\x32\x45.prodvana.desired_state.DesiredStateSummary.PreconditionStatusesEntry\x12`\n\x13status_explanations\x18\t \x03(\x0b\x32\x43.prodvana.desired_state.DesiredStateSummary.StatusExplanationsEntry\x12N\n\ndebug_logs\x18\n \x03(\x0b\x32:.prodvana.desired_state.DesiredStateSummary.DebugLogsEntry\x12`\n\x13\x61\x63tion_explanations\x18\x0b \x03(\x0b\x32\x43.prodvana.desired_state.DesiredStateSummary.ActionExplanationsEntry\x12\x65\n\x16last_update_timestamps\x18\x0c \x03(\x0b\x32\x45.prodvana.desired_state.DesiredStateSummary.LastUpdateTimestampsEntry\x12g\n\x17last_fetched_timestamps\x18\r \x03(\x0b\x32\x46.prodvana.desired_state.DesiredStateSummary.LastFetchedTimestampsEntry\x12g\n\x17last_applied_timestamps\x18\x0e \x03(\x0b\x32\x46.prodvana.desired_state.DesiredStateSummary.LastAppliedTimestampsEntry\x1aU\n\rStatusesEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\x33\n\x05value\x18\x02 \x01(\x0e\x32$.prodvana.desired_state.model.Status:\x02\x38\x01\x1ai\n\x19PreconditionStatusesEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12;\n\x05value\x18\x02 \x01(\x0b\x32,.prodvana.desired_state.model.ConditionState:\x02\x38\x01\x1a\x65\n\x17StatusExplanationsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\x39\n\x05value\x18\x02 \x01(\x0b\x32*.prodvana.desired_state.StatusExplanations:\x02\x38\x01\x1aS\n\x0e\x44\x65\x62ugLogsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\x30\n\x05value\x18\x02 \x01(\x0b\x32!.prodvana.desired_state.DebugLogs:\x02\x38\x01\x1aj\n\x17\x41\x63tionExplanationsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12>\n\x05value\x18\x02 \x01(\x0b\x32/.prodvana.desired_state.model.ActionExplanation:\x02\x38\x01\x1aW\n\x19LastUpdateTimestampsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12)\n\x05value\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp:\x02\x38\x01\x1aX\n\x1aLastFetchedTimestampsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12)\n\x05value\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp:\x02\x38\x01\x1aX\n\x1aLastAppliedTimestampsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12)\n\x05value\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp:\x02\x38\x01\"B\n\x1dGetDesiredStateConvergenceReq\x12!\n\x10\x64\x65sired_state_id\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\"e\n%GetDesiredStateConvergenceSummaryResp\x12<\n\x07summary\x18\x01 \x01(\x0b\x32+.prodvana.desired_state.DesiredStateSummary\"l\n,GetServiceDesiredStateConvergenceSummaryResp\x12<\n\x07summary\x18\x01 \x01(\x0b\x32+.prodvana.desired_state.DesiredStateSummary\"Y\n\x1fGetServiceLastConvergedStateReq\x12\x1c\n\x0b\x61pplication\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x18\n\x07service\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\"w\n GetServiceLastConvergedStateResp\x12S\n\x17service_instance_states\x18\x01 \x03(\x0b\x32\x32.prodvana.desired_state.model.ServiceInstanceState\"\x81\x01\n GetServiceDesiredStateHistoryReq\x12\x1c\n\x0b\x61pplication\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x18\n\x07service\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x12\n\npage_token\x18\x03 \x01(\t\x12\x11\n\tpage_size\x18\x04 \x01(\x05\"\x81\x01\n!GetServiceDesiredStateHistoryResp\x12\x43\n\x0e\x64\x65sired_states\x18\x01 \x03(\x0b\x32+.prodvana.desired_state.DesiredStateSummary\x12\x17\n\x0fnext_page_token\x18\x02 \x01(\t\"7\n\x12GetDesiredStateReq\x12!\n\x10\x64\x65sired_state_id\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\"\x96\x01\n\x13GetDesiredStateResp\x12:\n\rdesired_state\x18\x01 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x43\n\x16\x63ompiled_desired_state\x18\x02 \x01(\x0b\x32#.prodvana.desired_state.model.State\"\x9b\x01\n\x18ValidateDesiredStateResp\x12:\n\rdesired_state\x18\x01 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x43\n\x16\x63ompiled_desired_state\x18\x02 \x01(\x0b\x32#.prodvana.desired_state.model.State\"a\n\x14SetManualApprovalReq\x12!\n\x10\x64\x65sired_state_id\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x16\n\x05topic\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x0e\n\x06reject\x18\x03 \x01(\x08\"\x17\n\x15SetManualApprovalResp\"m\n\x12PromoteDeliveryReq\x12!\n\x10\x64\x65sired_state_id\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\r\n\x05stage\x18\x02 \x01(\x03\x12\x0c\n\x04\x66ull\x18\x03 \x01(\x08\x12\x17\n\x06source\x18\x04 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\"\x15\n\x13PromoteDeliveryResp\"Q\n\x13\x42ypassProtectionReq\x12!\n\x10\x64\x65sired_state_id\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x17\n\x06source\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\"\x16\n\x14\x42ypassProtectionResp2\xf5\x0e\n\x13\x44\x65siredStateManager\x12\x89\x01\n\x0fSetDesiredState\x12*.prodvana.desired_state.SetDesiredStateReq\x1a+.prodvana.desired_state.SetDesiredStateResp\"\x1d\x82\xd3\xe4\x93\x02\x17\"\x12/v1/desired_states:\x01*\x12\x89\x02\n(GetServiceDesiredStateConvergenceSummary\x12\x43.prodvana.desired_state.GetServiceDesiredStateConvergenceSummaryReq\x1a\x44.prodvana.desired_state.GetServiceDesiredStateConvergenceSummaryResp\"R\x82\xd3\xe4\x93\x02L\x12J/v1/applications/{application=*}/services/{service=*}/latest_desired_state\x12\xe0\x01\n\x1dGetServiceLastConvergedStates\x12\x37.prodvana.desired_state.GetServiceLastConvergedStateReq\x1a\x38.prodvana.desired_state.GetServiceLastConvergedStateResp\"L\x82\xd3\xe4\x93\x02\x46\x12\x44/v1/applications/{application=*}/services/{service=*}/last_converged\x12\xe2\x01\n\x1dGetServiceDesiredStateHistory\x12\x38.prodvana.desired_state.GetServiceDesiredStateHistoryReq\x1a\x39.prodvana.desired_state.GetServiceDesiredStateHistoryResp\"L\x82\xd3\xe4\x93\x02\x46\x12\x44/v1/applications/{application=*}/services/{service=*}/desired_states\x12\x9b\x01\n\x0fGetDesiredState\x12*.prodvana.desired_state.GetDesiredStateReq\x1a+.prodvana.desired_state.GetDesiredStateResp\"/\x82\xd3\xe4\x93\x02)\x12\'/v1/desired_states/{desired_state_id=*}\x12\xd2\x01\n!GetDesiredStateConvergenceSummary\x12\x35.prodvana.desired_state.GetDesiredStateConvergenceReq\x1a=.prodvana.desired_state.GetDesiredStateConvergenceSummaryResp\"7\x82\xd3\xe4\x93\x02\x31\x12//v1/desired_states/{desired_state_id=*}/summary\x12\xa1\x01\n\x14ValidateDesiredState\x12/.prodvana.desired_state.ValidateDesiredStateReq\x1a\x30.prodvana.desired_state.ValidateDesiredStateResp\"&\x82\xd3\xe4\x93\x02 \"\x1b/v1/desired_states/validate:\x01*\x12\xa8\x01\n\x11SetManualApproval\x12,.prodvana.desired_state.SetManualApprovalReq\x1a-.prodvana.desired_state.SetManualApprovalResp\"6\x82\xd3\xe4\x93\x02\x30\"+/v1/desired_states/approve_manual_condition:\x01*\x12\x9a\x01\n\x0fPromoteDelivery\x12*.prodvana.desired_state.PromoteDeliveryReq\x1a+.prodvana.desired_state.PromoteDeliveryResp\".\x82\xd3\xe4\x93\x02(\"#/v1/desired_states/promote_delivery:\x01*\x12\x9e\x01\n\x10\x42ypassProtection\x12+.prodvana.desired_state.BypassProtectionReq\x1a,.prodvana.desired_state.BypassProtectionResp\"/\x82\xd3\xe4\x93\x02)\"$/v1/desired_states/bypass_protection:\x01*BRZPgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_stateb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'prodvana.desired_state.manager_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZPgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state'
  _SETDESIREDSTATEREQ.fields_by_name['desired_state']._options = None
  _SETDESIREDSTATEREQ.fields_by_name['desired_state']._serialized_options = b'\372B\005\212\001\002\020\001'
  _VALIDATEDESIREDSTATEREQ.fields_by_name['desired_state']._options = None
  _VALIDATEDESIREDSTATEREQ.fields_by_name['desired_state']._serialized_options = b'\372B\005\212\001\002\020\001'
  _GETSERVICEDESIREDSTATECONVERGENCESUMMARYREQ.fields_by_name['application']._options = None
  _GETSERVICEDESIREDSTATECONVERGENCESUMMARYREQ.fields_by_name['application']._serialized_options = b'\372B\004r\002\020\001'
  _GETSERVICEDESIREDSTATECONVERGENCESUMMARYREQ.fields_by_name['service']._options = None
  _GETSERVICEDESIREDSTATECONVERGENCESUMMARYREQ.fields_by_name['service']._serialized_options = b'\372B\004r\002\020\001'
  _DESIREDSTATESUMMARY_STATUSESENTRY._options = None
  _DESIREDSTATESUMMARY_STATUSESENTRY._serialized_options = b'8\001'
  _DESIREDSTATESUMMARY_PRECONDITIONSTATUSESENTRY._options = None
  _DESIREDSTATESUMMARY_PRECONDITIONSTATUSESENTRY._serialized_options = b'8\001'
  _DESIREDSTATESUMMARY_STATUSEXPLANATIONSENTRY._options = None
  _DESIREDSTATESUMMARY_STATUSEXPLANATIONSENTRY._serialized_options = b'8\001'
  _DESIREDSTATESUMMARY_DEBUGLOGSENTRY._options = None
  _DESIREDSTATESUMMARY_DEBUGLOGSENTRY._serialized_options = b'8\001'
  _DESIREDSTATESUMMARY_ACTIONEXPLANATIONSENTRY._options = None
  _DESIREDSTATESUMMARY_ACTIONEXPLANATIONSENTRY._serialized_options = b'8\001'
  _DESIREDSTATESUMMARY_LASTUPDATETIMESTAMPSENTRY._options = None
  _DESIREDSTATESUMMARY_LASTUPDATETIMESTAMPSENTRY._serialized_options = b'8\001'
  _DESIREDSTATESUMMARY_LASTFETCHEDTIMESTAMPSENTRY._options = None
  _DESIREDSTATESUMMARY_LASTFETCHEDTIMESTAMPSENTRY._serialized_options = b'8\001'
  _DESIREDSTATESUMMARY_LASTAPPLIEDTIMESTAMPSENTRY._options = None
  _DESIREDSTATESUMMARY_LASTAPPLIEDTIMESTAMPSENTRY._serialized_options = b'8\001'
  _GETDESIREDSTATECONVERGENCEREQ.fields_by_name['desired_state_id']._options = None
  _GETDESIREDSTATECONVERGENCEREQ.fields_by_name['desired_state_id']._serialized_options = b'\372B\004r\002\020\001'
  _GETSERVICELASTCONVERGEDSTATEREQ.fields_by_name['application']._options = None
  _GETSERVICELASTCONVERGEDSTATEREQ.fields_by_name['application']._serialized_options = b'\372B\004r\002\020\001'
  _GETSERVICELASTCONVERGEDSTATEREQ.fields_by_name['service']._options = None
  _GETSERVICELASTCONVERGEDSTATEREQ.fields_by_name['service']._serialized_options = b'\372B\004r\002\020\001'
  _GETSERVICEDESIREDSTATEHISTORYREQ.fields_by_name['application']._options = None
  _GETSERVICEDESIREDSTATEHISTORYREQ.fields_by_name['application']._serialized_options = b'\372B\004r\002\020\001'
  _GETSERVICEDESIREDSTATEHISTORYREQ.fields_by_name['service']._options = None
  _GETSERVICEDESIREDSTATEHISTORYREQ.fields_by_name['service']._serialized_options = b'\372B\004r\002\020\001'
  _GETDESIREDSTATEREQ.fields_by_name['desired_state_id']._options = None
  _GETDESIREDSTATEREQ.fields_by_name['desired_state_id']._serialized_options = b'\372B\004r\002\020\001'
  _SETMANUALAPPROVALREQ.fields_by_name['desired_state_id']._options = None
  _SETMANUALAPPROVALREQ.fields_by_name['desired_state_id']._serialized_options = b'\372B\004r\002\020\001'
  _SETMANUALAPPROVALREQ.fields_by_name['topic']._options = None
  _SETMANUALAPPROVALREQ.fields_by_name['topic']._serialized_options = b'\372B\004r\002\020\001'
  _PROMOTEDELIVERYREQ.fields_by_name['desired_state_id']._options = None
  _PROMOTEDELIVERYREQ.fields_by_name['desired_state_id']._serialized_options = b'\372B\004r\002\020\001'
  _PROMOTEDELIVERYREQ.fields_by_name['source']._options = None
  _PROMOTEDELIVERYREQ.fields_by_name['source']._serialized_options = b'\372B\004r\002\020\001'
  _BYPASSPROTECTIONREQ.fields_by_name['desired_state_id']._options = None
  _BYPASSPROTECTIONREQ.fields_by_name['desired_state_id']._serialized_options = b'\372B\004r\002\020\001'
  _BYPASSPROTECTIONREQ.fields_by_name['source']._options = None
  _BYPASSPROTECTIONREQ.fields_by_name['source']._serialized_options = b'\372B\004r\002\020\001'
  _DESIREDSTATEMANAGER.methods_by_name['SetDesiredState']._options = None
  _DESIREDSTATEMANAGER.methods_by_name['SetDesiredState']._serialized_options = b'\202\323\344\223\002\027\"\022/v1/desired_states:\001*'
  _DESIREDSTATEMANAGER.methods_by_name['GetServiceDesiredStateConvergenceSummary']._options = None
  _DESIREDSTATEMANAGER.methods_by_name['GetServiceDesiredStateConvergenceSummary']._serialized_options = b'\202\323\344\223\002L\022J/v1/applications/{application=*}/services/{service=*}/latest_desired_state'
  _DESIREDSTATEMANAGER.methods_by_name['GetServiceLastConvergedStates']._options = None
  _DESIREDSTATEMANAGER.methods_by_name['GetServiceLastConvergedStates']._serialized_options = b'\202\323\344\223\002F\022D/v1/applications/{application=*}/services/{service=*}/last_converged'
  _DESIREDSTATEMANAGER.methods_by_name['GetServiceDesiredStateHistory']._options = None
  _DESIREDSTATEMANAGER.methods_by_name['GetServiceDesiredStateHistory']._serialized_options = b'\202\323\344\223\002F\022D/v1/applications/{application=*}/services/{service=*}/desired_states'
  _DESIREDSTATEMANAGER.methods_by_name['GetDesiredState']._options = None
  _DESIREDSTATEMANAGER.methods_by_name['GetDesiredState']._serialized_options = b'\202\323\344\223\002)\022\'/v1/desired_states/{desired_state_id=*}'
  _DESIREDSTATEMANAGER.methods_by_name['GetDesiredStateConvergenceSummary']._options = None
  _DESIREDSTATEMANAGER.methods_by_name['GetDesiredStateConvergenceSummary']._serialized_options = b'\202\323\344\223\0021\022//v1/desired_states/{desired_state_id=*}/summary'
  _DESIREDSTATEMANAGER.methods_by_name['ValidateDesiredState']._options = None
  _DESIREDSTATEMANAGER.methods_by_name['ValidateDesiredState']._serialized_options = b'\202\323\344\223\002 \"\033/v1/desired_states/validate:\001*'
  _DESIREDSTATEMANAGER.methods_by_name['SetManualApproval']._options = None
  _DESIREDSTATEMANAGER.methods_by_name['SetManualApproval']._serialized_options = b'\202\323\344\223\0020\"+/v1/desired_states/approve_manual_condition:\001*'
  _DESIREDSTATEMANAGER.methods_by_name['PromoteDelivery']._options = None
  _DESIREDSTATEMANAGER.methods_by_name['PromoteDelivery']._serialized_options = b'\202\323\344\223\002(\"#/v1/desired_states/promote_delivery:\001*'
  _DESIREDSTATEMANAGER.methods_by_name['BypassProtection']._options = None
  _DESIREDSTATEMANAGER.methods_by_name['BypassProtection']._serialized_options = b'\202\323\344\223\002)\"$/v1/desired_states/bypass_protection:\001*'
  _globals['_SETDESIREDSTATEREQ']._serialized_start=284
  _globals['_SETDESIREDSTATEREQ']._serialized_end=392
  _globals['_VALIDATEDESIREDSTATEREQ']._serialized_start=395
  _globals['_VALIDATEDESIREDSTATEREQ']._serialized_end=591
  _globals['_SETDESIREDSTATERESP']._serialized_start=593
  _globals['_SETDESIREDSTATERESP']._serialized_end=640
  _globals['_GETSERVICEDESIREDSTATECONVERGENCESUMMARYREQ']._serialized_start=642
  _globals['_GETSERVICEDESIREDSTATECONVERGENCESUMMARYREQ']._serialized_end=743
  _globals['_STATUSEXPLANATIONS']._serialized_start=745
  _globals['_STATUSEXPLANATIONS']._serialized_end=843
  _globals['_DEBUGLOGS']._serialized_start=845
  _globals['_DEBUGLOGS']._serialized_end=916
  _globals['_DESIREDSTATESUMMARY']._serialized_start=919
  _globals['_DESIREDSTATESUMMARY']._serialized_end=2886
  _globals['_DESIREDSTATESUMMARY_STATUSESENTRY']._serialized_start=2129
  _globals['_DESIREDSTATESUMMARY_STATUSESENTRY']._serialized_end=2214
  _globals['_DESIREDSTATESUMMARY_PRECONDITIONSTATUSESENTRY']._serialized_start=2216
  _globals['_DESIREDSTATESUMMARY_PRECONDITIONSTATUSESENTRY']._serialized_end=2321
  _globals['_DESIREDSTATESUMMARY_STATUSEXPLANATIONSENTRY']._serialized_start=2323
  _globals['_DESIREDSTATESUMMARY_STATUSEXPLANATIONSENTRY']._serialized_end=2424
  _globals['_DESIREDSTATESUMMARY_DEBUGLOGSENTRY']._serialized_start=2426
  _globals['_DESIREDSTATESUMMARY_DEBUGLOGSENTRY']._serialized_end=2509
  _globals['_DESIREDSTATESUMMARY_ACTIONEXPLANATIONSENTRY']._serialized_start=2511
  _globals['_DESIREDSTATESUMMARY_ACTIONEXPLANATIONSENTRY']._serialized_end=2617
  _globals['_DESIREDSTATESUMMARY_LASTUPDATETIMESTAMPSENTRY']._serialized_start=2619
  _globals['_DESIREDSTATESUMMARY_LASTUPDATETIMESTAMPSENTRY']._serialized_end=2706
  _globals['_DESIREDSTATESUMMARY_LASTFETCHEDTIMESTAMPSENTRY']._serialized_start=2708
  _globals['_DESIREDSTATESUMMARY_LASTFETCHEDTIMESTAMPSENTRY']._serialized_end=2796
  _globals['_DESIREDSTATESUMMARY_LASTAPPLIEDTIMESTAMPSENTRY']._serialized_start=2798
  _globals['_DESIREDSTATESUMMARY_LASTAPPLIEDTIMESTAMPSENTRY']._serialized_end=2886
  _globals['_GETDESIREDSTATECONVERGENCEREQ']._serialized_start=2888
  _globals['_GETDESIREDSTATECONVERGENCEREQ']._serialized_end=2954
  _globals['_GETDESIREDSTATECONVERGENCESUMMARYRESP']._serialized_start=2956
  _globals['_GETDESIREDSTATECONVERGENCESUMMARYRESP']._serialized_end=3057
  _globals['_GETSERVICEDESIREDSTATECONVERGENCESUMMARYRESP']._serialized_start=3059
  _globals['_GETSERVICEDESIREDSTATECONVERGENCESUMMARYRESP']._serialized_end=3167
  _globals['_GETSERVICELASTCONVERGEDSTATEREQ']._serialized_start=3169
  _globals['_GETSERVICELASTCONVERGEDSTATEREQ']._serialized_end=3258
  _globals['_GETSERVICELASTCONVERGEDSTATERESP']._serialized_start=3260
  _globals['_GETSERVICELASTCONVERGEDSTATERESP']._serialized_end=3379
  _globals['_GETSERVICEDESIREDSTATEHISTORYREQ']._serialized_start=3382
  _globals['_GETSERVICEDESIREDSTATEHISTORYREQ']._serialized_end=3511
  _globals['_GETSERVICEDESIREDSTATEHISTORYRESP']._serialized_start=3514
  _globals['_GETSERVICEDESIREDSTATEHISTORYRESP']._serialized_end=3643
  _globals['_GETDESIREDSTATEREQ']._serialized_start=3645
  _globals['_GETDESIREDSTATEREQ']._serialized_end=3700
  _globals['_GETDESIREDSTATERESP']._serialized_start=3703
  _globals['_GETDESIREDSTATERESP']._serialized_end=3853
  _globals['_VALIDATEDESIREDSTATERESP']._serialized_start=3856
  _globals['_VALIDATEDESIREDSTATERESP']._serialized_end=4011
  _globals['_SETMANUALAPPROVALREQ']._serialized_start=4013
  _globals['_SETMANUALAPPROVALREQ']._serialized_end=4110
  _globals['_SETMANUALAPPROVALRESP']._serialized_start=4112
  _globals['_SETMANUALAPPROVALRESP']._serialized_end=4135
  _globals['_PROMOTEDELIVERYREQ']._serialized_start=4137
  _globals['_PROMOTEDELIVERYREQ']._serialized_end=4246
  _globals['_PROMOTEDELIVERYRESP']._serialized_start=4248
  _globals['_PROMOTEDELIVERYRESP']._serialized_end=4269
  _globals['_BYPASSPROTECTIONREQ']._serialized_start=4271
  _globals['_BYPASSPROTECTIONREQ']._serialized_end=4352
  _globals['_BYPASSPROTECTIONRESP']._serialized_start=4354
  _globals['_BYPASSPROTECTIONRESP']._serialized_end=4376
  _globals['_DESIREDSTATEMANAGER']._serialized_start=4379
  _globals['_DESIREDSTATEMANAGER']._serialized_end=6288
# @@protoc_insertion_point(module_scope)
