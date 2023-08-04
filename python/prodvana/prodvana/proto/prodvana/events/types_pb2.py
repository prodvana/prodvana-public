# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/events/types.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from prodvana.proto.prodvana.desired_state.model import desired_state_pb2 as prodvana_dot_desired__state_dot_model_dot_desired__state__pb2
from google.protobuf import any_pb2 as google_dot_protobuf_dot_any__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1bprodvana/events/types.proto\x12\x0fprodvana.events\x1a\x30prodvana/desired_state/model/desired_state.proto\x1a\x19google/protobuf/any.proto\"\xd0\x01\n\x14SetDesiredStateEvent\x12\x34\n\x02id\x18\x01 \x01(\x0b\x32(.prodvana.desired_state.model.Identifier\x12\x18\n\x10\x64\x65sired_state_id\x18\x03 \x01(\t\x12\x1d\n\x15root_desired_state_id\x18\x04 \x01(\t\x12\x34\n\x07\x64\x65sired\x18\x02 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x13\n\x0bis_rollback\x18\x05 \x01(\x08\"\xc3\x03\n\x13SetTargetStateEvent\x12\x34\n\x02id\x18\x01 \x01(\x0b\x32(.prodvana.desired_state.model.Identifier\x12\x18\n\x10\x64\x65sired_state_id\x18\x08 \x01(\t\x12\x1d\n\x15root_desired_state_id\x18\t \x01(\t\x12\x33\n\x06target\x18\x02 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x34\n\x07\x64\x65sired\x18\x03 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x34\n\x07\x63urrent\x18\x04 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x18\n\x10is_auto_rollback\x18\x05 \x01(\x08\x12\x34\n\x06status\x18\x06 \x01(\x0e\x32$.prodvana.desired_state.model.Status\x12L\n\x13status_explanations\x18\x07 \x03(\x0b\x32/.prodvana.desired_state.model.StatusExplanation\"\xac\x03\n\x15\x41pplyTargetStateEvent\x12\x34\n\x02id\x18\x01 \x01(\x0b\x32(.prodvana.desired_state.model.Identifier\x12\x18\n\x10\x64\x65sired_state_id\x18\x06 \x01(\t\x12\x1d\n\x15root_desired_state_id\x18\x07 \x01(\t\x12\x33\n\x06target\x18\x02 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x34\n\x07\x63urrent\x18\x03 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x18\n\x10is_auto_rollback\x18\x04 \x01(\x08\x12\x42\n\x06result\x18\x05 \x01(\x0e\x32\x32.prodvana.events.ApplyTargetStateEvent.ApplyResult\x12\r\n\x05\x65rror\x18\x08 \x01(\t\"L\n\x0b\x41pplyResult\x12\x0b\n\x07UNKNOWN\x10\x00\x12\r\n\tSUCCEEDED\x10\x01\x12\n\n\x06\x46\x41ILED\x10\x02\x12\x15\n\x11RETRIABLE_FAILURE\x10\x03\"\xb0\x01\n\x10ProgramExitEvent\x12\x11\n\tnamespace\x18\x01 \x01(\t\x12\x0b\n\x03pod\x18\x02 \x01(\t\x12\x0f\n\x07program\x18\x03 \x01(\t\x12\x15\n\rrestart_count\x18\x04 \x01(\x05\x12\x0e\n\x06reason\x18\x05 \x01(\t\x12\x15\n\rkilled_reason\x18\t \x01(\t\x12\x11\n\texit_code\x18\x06 \x01(\x05\x12\x0e\n\x06signal\x18\x07 \x01(\x05J\x04\x08\x08\x10\tR\x04logs\"S\n\x15\x44\x65liveryProgressEvent\x12:\n\x05state\x18\x01 \x01(\x0b\x32+.prodvana.desired_state.model.DeliveryState\"w\n\x1c\x44\x65liveryManualPromotionEvent\x12:\n\x05state\x18\x01 \x01(\x0b\x32+.prodvana.desired_state.model.DeliveryState\x12\r\n\x05stage\x18\x02 \x01(\x03\x12\x0c\n\x04\x66ull\x18\x03 \x01(\x08\"h\n\x13ManualApprovalEvent\x12\x42\n\x06status\x18\x01 \x01(\x0e\x32\x32.prodvana.desired_state.model.ManualApprovalStatus\x12\r\n\x05topic\x18\x02 \x01(\t\"\\\n\x18\x43ustomTaskExecutionEvent\x12\x12\n\nsuccessful\x18\x01 \x01(\x08\x12\x10\n\x08\x61ttempts\x18\x02 \x01(\x05\x12\x0e\n\x06reason\x18\x03 \x01(\tJ\x04\x08\x04\x10\x05R\x04logs\"K\n\rRuntimeObject\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0c\n\x04kind\x18\x02 \x01(\t\x12\x11\n\tnamespace\x18\x03 \x01(\t\x12\x0b\n\x03url\x18\x04 \x01(\t\"\xe4\x02\n\x12RuntimeUpdateEvent\x12\x41\n\x06\x61\x63tion\x18\x01 \x01(\x0e\x32\x31.prodvana.events.RuntimeUpdateEvent.RuntimeAction\x12.\n\x06object\x18\x02 \x01(\x0b\x32\x1e.prodvana.events.RuntimeObject\x12\x41\n\x06status\x18\x03 \x01(\x0e\x32\x31.prodvana.events.RuntimeUpdateEvent.RuntimeStatus\"D\n\rRuntimeAction\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x14\n\x10\x43REATE_OR_UPDATE\x10\x01\x12\x10\n\x0cWAIT_HEALTHY\x10\x02\"R\n\rRuntimeStatus\x12\x1a\n\x16RUNTIME_STATUS_UNKNOWN\x10\x00\x12\x0b\n\x07PENDING\x10\x01\x12\x0b\n\x07SUCCESS\x10\x02\x12\x0b\n\x07\x46\x41ILURE\x10\x03\"\xab\x04\n\x1d\x44\x65siredStateStatusChangeEvent\x12\x34\n\x02id\x18\x01 \x01(\x0b\x32(.prodvana.desired_state.model.Identifier\x12\x18\n\x10\x64\x65sired_state_id\x18\x02 \x01(\t\x12\x1d\n\x15root_desired_state_id\x18\x03 \x01(\t\x12\x38\n\nold_status\x18\x04 \x01(\x0e\x32$.prodvana.desired_state.model.Status\x12\x38\n\nnew_status\x18\x05 \x01(\x0e\x32$.prodvana.desired_state.model.Status\x12\x34\n\x07\x64\x65sired\x18\x06 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x33\n\x06target\x18\x07 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x36\n\told_state\x18\x08 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x36\n\tnew_state\x18\t \x01(\x0b\x32#.prodvana.desired_state.model.State\x12L\n\x13status_explanations\x18\n \x03(\x0b\x32/.prodvana.desired_state.model.StatusExplanation\"\xc6\x04\n\x18KeyDeliveryDecisionEvent\x12\x34\n\x02id\x18\x01 \x01(\x0b\x32(.prodvana.desired_state.model.Identifier\x12\x18\n\x10\x64\x65sired_state_id\x18\x02 \x01(\t\x12\x1d\n\x15root_desired_state_id\x18\x03 \x01(\t\x12\x44\n\x08\x64\x65\x63ision\x18\x04 \x01(\x0e\x32\x32.prodvana.events.KeyDeliveryDecisionEvent.Decision\x12\x13\n\x0b\x65xplanation\x18\x05 \x01(\t\x12\x34\n\x06status\x18\x06 \x01(\x0e\x32$.prodvana.desired_state.model.Status\x12\x34\n\x07\x64\x65sired\x18\x07 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x33\n\x06target\x18\x08 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x34\n\x07\x63urrent\x18\t \x01(\x0b\x32#.prodvana.desired_state.model.State\x12L\n\x13status_explanations\x18\n \x03(\x0b\x32/.prodvana.desired_state.model.StatusExplanation\";\n\x08\x44\x65\x63ision\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x16\n\x12ROLLBACK_INITIATED\x10\x01\x12\n\n\x06\x46\x41ILED\x10\x02\"\xe2\x01\n\x0cRpcCallEvent\x12\x30\n\x04type\x18\x01 \x01(\x0e\x32\".prodvana.events.RpcCallEvent.Type\x12\x13\n\x0brpc_service\x18\x02 \x01(\t\x12\x12\n\nrpc_method\x18\x03 \x01(\t\x12%\n\x07request\x18\x04 \x01(\x0b\x32\x14.google.protobuf.Any\x12&\n\x08response\x18\x05 \x01(\x0b\x32\x14.google.protobuf.Any\"(\n\x04Type\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x08\n\x04READ\x10\x01\x12\t\n\x05WRITE\x10\x02\"-\n\x11\x41pplicationHandle\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\n\n\x02id\x18\x02 \x01(\t\"i\n\x14ReleaseChannelHandle\x12\x37\n\x0b\x61pplication\x18\x01 \x01(\x0b\x32\".prodvana.events.ApplicationHandle\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\n\n\x02id\x18\x03 \x01(\t\"\xa2\x01\n\rServiceHandle\x12\x37\n\x0b\x61pplication\x18\x01 \x01(\x0b\x32\".prodvana.events.ApplicationHandle\x12>\n\x0frelease_channel\x18\x02 \x01(\x0b\x32%.prodvana.events.ReleaseChannelHandle\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\n\n\x02id\x18\x04 \x01(\t\")\n\rRuntimeHandle\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\n\n\x02id\x18\x02 \x01(\t\",\n\x10ProtectionHandle\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\n\n\x02id\x18\x02 \x01(\t\"\xa8\x03\n\x15PermissionDeniedEvent\x12=\n\x06\x61\x63tion\x18\x01 \x01(\x0e\x32-.prodvana.events.PermissionDeniedEvent.Action\x12\x39\n\x0b\x61pplication\x18\x02 \x01(\x0b\x32\".prodvana.events.ApplicationHandleH\x00\x12@\n\x0frelease_channel\x18\x03 \x01(\x0b\x32%.prodvana.events.ReleaseChannelHandleH\x00\x12\x31\n\x07service\x18\x04 \x01(\x0b\x32\x1e.prodvana.events.ServiceHandleH\x00\x12\x31\n\x07runtime\x18\x05 \x01(\x0b\x32\x1e.prodvana.events.RuntimeHandleH\x00\x12\x37\n\nprotection\x18\x06 \x01(\x0b\x32!.prodvana.events.ProtectionHandleH\x00\"*\n\x06\x41\x63tion\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x08\n\x04READ\x10\x01\x12\t\n\x05WRITE\x10\x02\x42\x08\n\x06target\"\xd2\x02\n\x12ObjectCreatedEvent\x12\x16\n\x0e\x63onfig_version\x18\x01 \x01(\t\x12\x39\n\x0b\x61pplication\x18\x02 \x01(\x0b\x32\".prodvana.events.ApplicationHandleH\x00\x12@\n\x0frelease_channel\x18\x03 \x01(\x0b\x32%.prodvana.events.ReleaseChannelHandleH\x00\x12\x31\n\x07service\x18\x04 \x01(\x0b\x32\x1e.prodvana.events.ServiceHandleH\x00\x12\x31\n\x07runtime\x18\x05 \x01(\x0b\x32\x1e.prodvana.events.RuntimeHandleH\x00\x12\x37\n\nprotection\x18\x06 \x01(\x0b\x32!.prodvana.events.ProtectionHandleH\x00\x42\x08\n\x06target\"\xd2\x02\n\x12ObjectDeletedEvent\x12\x16\n\x0e\x63onfig_version\x18\x01 \x01(\t\x12\x39\n\x0b\x61pplication\x18\x02 \x01(\x0b\x32\".prodvana.events.ApplicationHandleH\x00\x12@\n\x0frelease_channel\x18\x03 \x01(\x0b\x32%.prodvana.events.ReleaseChannelHandleH\x00\x12\x31\n\x07service\x18\x04 \x01(\x0b\x32\x1e.prodvana.events.ServiceHandleH\x00\x12\x31\n\x07runtime\x18\x05 \x01(\x0b\x32\x1e.prodvana.events.RuntimeHandleH\x00\x12\x37\n\nprotection\x18\x06 \x01(\x0b\x32!.prodvana.events.ProtectionHandleH\x00\x42\x08\n\x06target\"\xf3\x02\n\x13ObjectModifiedEvent\x12\x1a\n\x12old_config_version\x18\x01 \x01(\t\x12\x1a\n\x12new_config_version\x18\x02 \x01(\t\x12\x39\n\x0b\x61pplication\x18\x03 \x01(\x0b\x32\".prodvana.events.ApplicationHandleH\x00\x12@\n\x0frelease_channel\x18\x04 \x01(\x0b\x32%.prodvana.events.ReleaseChannelHandleH\x00\x12\x31\n\x07service\x18\x05 \x01(\x0b\x32\x1e.prodvana.events.ServiceHandleH\x00\x12\x31\n\x07runtime\x18\x06 \x01(\x0b\x32\x1e.prodvana.events.RuntimeHandleH\x00\x12\x37\n\nprotection\x18\x07 \x01(\x0b\x32!.prodvana.events.ProtectionHandleH\x00\x42\x08\n\x06target\"\xde\x08\n\x0c\x45ventDetails\x12\x42\n\x11set_desired_state\x18\x01 \x01(\x0b\x32%.prodvana.events.SetDesiredStateEventH\x00\x12@\n\x10set_target_state\x18\x02 \x01(\x0b\x32$.prodvana.events.SetTargetStateEventH\x00\x12\x39\n\x0cprogram_exit\x18\x03 \x01(\x0b\x32!.prodvana.events.ProgramExitEventH\x00\x12\x44\n\x12\x61pply_target_state\x18\x04 \x01(\x0b\x32&.prodvana.events.ApplyTargetStateEventH\x00\x12?\n\x0fmanual_approval\x18\x06 \x01(\x0b\x32$.prodvana.events.ManualApprovalEventH\x00\x12J\n\x15\x63ustom_task_execution\x18\x07 \x01(\x0b\x32).prodvana.events.CustomTaskExecutionEventH\x00\x12=\n\x0eruntime_update\x18\x08 \x01(\x0b\x32#.prodvana.events.RuntimeUpdateEventH\x00\x12\x43\n\x11\x64\x65livery_progress\x18\t \x01(\x0b\x32&.prodvana.events.DeliveryProgressEventH\x00\x12K\n\x12\x64\x65livery_promotion\x18\n \x01(\x0b\x32-.prodvana.events.DeliveryManualPromotionEventH\x00\x12U\n\x1b\x64\x65sired_state_status_change\x18\x0b \x01(\x0b\x32..prodvana.events.DesiredStateStatusChangeEventH\x00\x12J\n\x15key_delivery_decision\x18\x0c \x01(\x0b\x32).prodvana.events.KeyDeliveryDecisionEventH\x00\x12\x31\n\x08rpc_call\x18\r \x01(\x0b\x32\x1d.prodvana.events.RpcCallEventH\x00\x12\x43\n\x11permission_denied\x18\x0e \x01(\x0b\x32&.prodvana.events.PermissionDeniedEventH\x00\x12=\n\x0eobject_created\x18\x0f \x01(\x0b\x32#.prodvana.events.ObjectCreatedEventH\x00\x12=\n\x0eobject_deleted\x18\x10 \x01(\x0b\x32#.prodvana.events.ObjectDeletedEventH\x00\x12?\n\x0fobject_modified\x18\x11 \x01(\x0b\x32$.prodvana.events.ObjectModifiedEventH\x00\x42\t\n\x07\x64\x65tailsJ\x04\x08\x05\x10\x06*\x8c\x03\n\tEventType\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x15\n\x11SET_DESIRED_STATE\x10\x01\x12\x14\n\x10SET_TARGET_STATE\x10\x02\x12\x10\n\x0cPROGRAM_EXIT\x10\x03\x12\x16\n\x12\x41PPLY_TARGET_STATE\x10\x04\x12\x13\n\x0fMANUAL_APPROVAL\x10\x06\x12\x19\n\x15\x43USTOM_TASK_EXECUTION\x10\x07\x12\x12\n\x0eRUNTIME_UPDATE\x10\x08\x12\x15\n\x11\x44\x45LIVERY_PROGRESS\x10\t\x12\x1d\n\x19\x44\x45LIVERY_MANUAL_PROMOTION\x10\n\x12\x1f\n\x1b\x44\x45SIRED_STATE_STATUS_CHANGE\x10\x0b\x12\x19\n\x15KEY_DELIVERY_DECISION\x10\x0c\x12\x0c\n\x08RPC_CALL\x10\r\x12\x14\n\x10PERMISSON_DENIED\x10\x0e\x12\x12\n\x0eOBJECT_CREATED\x10\x0f\x12\x12\n\x0eOBJECT_DELETED\x10\x10\x12\x13\n\x0fOBJECT_MODIFIED\x10\x11\"\x04\x08\x05\x10\x05\x42KZIgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/eventsb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'prodvana.events.types_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZIgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/events'
  _globals['_EVENTTYPE']._serialized_start=6627
  _globals['_EVENTTYPE']._serialized_end=7023
  _globals['_SETDESIREDSTATEEVENT']._serialized_start=126
  _globals['_SETDESIREDSTATEEVENT']._serialized_end=334
  _globals['_SETTARGETSTATEEVENT']._serialized_start=337
  _globals['_SETTARGETSTATEEVENT']._serialized_end=788
  _globals['_APPLYTARGETSTATEEVENT']._serialized_start=791
  _globals['_APPLYTARGETSTATEEVENT']._serialized_end=1219
  _globals['_APPLYTARGETSTATEEVENT_APPLYRESULT']._serialized_start=1143
  _globals['_APPLYTARGETSTATEEVENT_APPLYRESULT']._serialized_end=1219
  _globals['_PROGRAMEXITEVENT']._serialized_start=1222
  _globals['_PROGRAMEXITEVENT']._serialized_end=1398
  _globals['_DELIVERYPROGRESSEVENT']._serialized_start=1400
  _globals['_DELIVERYPROGRESSEVENT']._serialized_end=1483
  _globals['_DELIVERYMANUALPROMOTIONEVENT']._serialized_start=1485
  _globals['_DELIVERYMANUALPROMOTIONEVENT']._serialized_end=1604
  _globals['_MANUALAPPROVALEVENT']._serialized_start=1606
  _globals['_MANUALAPPROVALEVENT']._serialized_end=1710
  _globals['_CUSTOMTASKEXECUTIONEVENT']._serialized_start=1712
  _globals['_CUSTOMTASKEXECUTIONEVENT']._serialized_end=1804
  _globals['_RUNTIMEOBJECT']._serialized_start=1806
  _globals['_RUNTIMEOBJECT']._serialized_end=1881
  _globals['_RUNTIMEUPDATEEVENT']._serialized_start=1884
  _globals['_RUNTIMEUPDATEEVENT']._serialized_end=2240
  _globals['_RUNTIMEUPDATEEVENT_RUNTIMEACTION']._serialized_start=2088
  _globals['_RUNTIMEUPDATEEVENT_RUNTIMEACTION']._serialized_end=2156
  _globals['_RUNTIMEUPDATEEVENT_RUNTIMESTATUS']._serialized_start=2158
  _globals['_RUNTIMEUPDATEEVENT_RUNTIMESTATUS']._serialized_end=2240
  _globals['_DESIREDSTATESTATUSCHANGEEVENT']._serialized_start=2243
  _globals['_DESIREDSTATESTATUSCHANGEEVENT']._serialized_end=2798
  _globals['_KEYDELIVERYDECISIONEVENT']._serialized_start=2801
  _globals['_KEYDELIVERYDECISIONEVENT']._serialized_end=3383
  _globals['_KEYDELIVERYDECISIONEVENT_DECISION']._serialized_start=3324
  _globals['_KEYDELIVERYDECISIONEVENT_DECISION']._serialized_end=3383
  _globals['_RPCCALLEVENT']._serialized_start=3386
  _globals['_RPCCALLEVENT']._serialized_end=3612
  _globals['_RPCCALLEVENT_TYPE']._serialized_start=3572
  _globals['_RPCCALLEVENT_TYPE']._serialized_end=3612
  _globals['_APPLICATIONHANDLE']._serialized_start=3614
  _globals['_APPLICATIONHANDLE']._serialized_end=3659
  _globals['_RELEASECHANNELHANDLE']._serialized_start=3661
  _globals['_RELEASECHANNELHANDLE']._serialized_end=3766
  _globals['_SERVICEHANDLE']._serialized_start=3769
  _globals['_SERVICEHANDLE']._serialized_end=3931
  _globals['_RUNTIMEHANDLE']._serialized_start=3933
  _globals['_RUNTIMEHANDLE']._serialized_end=3974
  _globals['_PROTECTIONHANDLE']._serialized_start=3976
  _globals['_PROTECTIONHANDLE']._serialized_end=4020
  _globals['_PERMISSIONDENIEDEVENT']._serialized_start=4023
  _globals['_PERMISSIONDENIEDEVENT']._serialized_end=4447
  _globals['_PERMISSIONDENIEDEVENT_ACTION']._serialized_start=4395
  _globals['_PERMISSIONDENIEDEVENT_ACTION']._serialized_end=4437
  _globals['_OBJECTCREATEDEVENT']._serialized_start=4450
  _globals['_OBJECTCREATEDEVENT']._serialized_end=4788
  _globals['_OBJECTDELETEDEVENT']._serialized_start=4791
  _globals['_OBJECTDELETEDEVENT']._serialized_end=5129
  _globals['_OBJECTMODIFIEDEVENT']._serialized_start=5132
  _globals['_OBJECTMODIFIEDEVENT']._serialized_end=5503
  _globals['_EVENTDETAILS']._serialized_start=5506
  _globals['_EVENTDETAILS']._serialized_end=6624
# @@protoc_insertion_point(module_scope)
