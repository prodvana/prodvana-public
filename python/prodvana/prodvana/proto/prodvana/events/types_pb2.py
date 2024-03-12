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


from prodvana.proto.prodvana.pvn_wrapper import output_pb2 as prodvana_dot_pvn__wrapper_dot_output__pb2
from prodvana.proto.prodvana.deployment.model import object_pb2 as prodvana_dot_deployment_dot_model_dot_object__pb2
from prodvana.proto.prodvana.desired_state.model import desired_state_pb2 as prodvana_dot_desired__state_dot_model_dot_desired__state__pb2
from google.protobuf import any_pb2 as google_dot_protobuf_dot_any__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1bprodvana/events/types.proto\x12\x0fprodvana.events\x1a!prodvana/pvn_wrapper/output.proto\x1a&prodvana/deployment/model/object.proto\x1a\x30prodvana/desired_state/model/desired_state.proto\x1a\x19google/protobuf/any.proto\"\xd0\x01\n\x14SetDesiredStateEvent\x12\x34\n\x02id\x18\x01 \x01(\x0b\x32(.prodvana.desired_state.model.Identifier\x12\x18\n\x10\x64\x65sired_state_id\x18\x03 \x01(\t\x12\x1d\n\x15root_desired_state_id\x18\x04 \x01(\t\x12\x34\n\x07\x64\x65sired\x18\x02 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x13\n\x0bis_rollback\x18\x05 \x01(\x08\"\xc3\x03\n\x13SetTargetStateEvent\x12\x34\n\x02id\x18\x01 \x01(\x0b\x32(.prodvana.desired_state.model.Identifier\x12\x18\n\x10\x64\x65sired_state_id\x18\x08 \x01(\t\x12\x1d\n\x15root_desired_state_id\x18\t \x01(\t\x12\x33\n\x06target\x18\x02 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x34\n\x07\x64\x65sired\x18\x03 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x34\n\x07\x63urrent\x18\x04 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x18\n\x10is_auto_rollback\x18\x05 \x01(\x08\x12\x34\n\x06status\x18\x06 \x01(\x0e\x32$.prodvana.desired_state.model.Status\x12L\n\x13status_explanations\x18\x07 \x03(\x0b\x32/.prodvana.desired_state.model.StatusExplanation\"\xac\x03\n\x15\x41pplyTargetStateEvent\x12\x34\n\x02id\x18\x01 \x01(\x0b\x32(.prodvana.desired_state.model.Identifier\x12\x18\n\x10\x64\x65sired_state_id\x18\x06 \x01(\t\x12\x1d\n\x15root_desired_state_id\x18\x07 \x01(\t\x12\x33\n\x06target\x18\x02 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x34\n\x07\x63urrent\x18\x03 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x18\n\x10is_auto_rollback\x18\x04 \x01(\x08\x12\x42\n\x06result\x18\x05 \x01(\x0e\x32\x32.prodvana.events.ApplyTargetStateEvent.ApplyResult\x12\r\n\x05\x65rror\x18\x08 \x01(\t\"L\n\x0b\x41pplyResult\x12\x0b\n\x07UNKNOWN\x10\x00\x12\r\n\tSUCCEEDED\x10\x01\x12\n\n\x06\x46\x41ILED\x10\x02\x12\x15\n\x11RETRIABLE_FAILURE\x10\x03\"\xe9\x01\n\x10ProgramExitEvent\x12\x11\n\tnamespace\x18\x01 \x01(\t\x12\x0b\n\x03pod\x18\x02 \x01(\t\x12\x0f\n\x07program\x18\x03 \x01(\t\x12\x15\n\rrestart_count\x18\x04 \x01(\x05\x12\x0e\n\x06reason\x18\x05 \x01(\t\x12\x15\n\rkilled_reason\x18\t \x01(\t\x12\x11\n\texit_code\x18\x06 \x01(\x05\x12\x0e\n\x06signal\x18\x07 \x01(\x05\x12\x37\n\x11structured_output\x18\n \x01(\x0b\x32\x1c.prodvana.pvn_wrapper.OutputJ\x04\x08\x08\x10\tR\x04logs\"S\n\x15\x44\x65liveryProgressEvent\x12:\n\x05state\x18\x01 \x01(\x0b\x32+.prodvana.desired_state.model.DeliveryState\"w\n\x1c\x44\x65liveryManualPromotionEvent\x12:\n\x05state\x18\x01 \x01(\x0b\x32+.prodvana.desired_state.model.DeliveryState\x12\r\n\x05stage\x18\x02 \x01(\x03\x12\x0c\n\x04\x66ull\x18\x03 \x01(\x08\"h\n\x13ManualApprovalEvent\x12\x42\n\x06status\x18\x01 \x01(\x0e\x32\x32.prodvana.desired_state.model.ManualApprovalStatus\x12\r\n\x05topic\x18\x02 \x01(\t\"\\\n\x18\x43ustomTaskExecutionEvent\x12\x12\n\nsuccessful\x18\x01 \x01(\x08\x12\x10\n\x08\x61ttempts\x18\x02 \x01(\x05\x12\x0e\n\x06reason\x18\x03 \x01(\tJ\x04\x08\x04\x10\x05R\x04logs\"K\n\rRuntimeObject\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0c\n\x04kind\x18\x02 \x01(\t\x12\x11\n\tnamespace\x18\x03 \x01(\t\x12\x0b\n\x03url\x18\x04 \x01(\t\"\xe4\x02\n\x12RuntimeUpdateEvent\x12\x41\n\x06\x61\x63tion\x18\x01 \x01(\x0e\x32\x31.prodvana.events.RuntimeUpdateEvent.RuntimeAction\x12.\n\x06object\x18\x02 \x01(\x0b\x32\x1e.prodvana.events.RuntimeObject\x12\x41\n\x06status\x18\x03 \x01(\x0e\x32\x31.prodvana.events.RuntimeUpdateEvent.RuntimeStatus\"D\n\rRuntimeAction\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x14\n\x10\x43REATE_OR_UPDATE\x10\x01\x12\x10\n\x0cWAIT_HEALTHY\x10\x02\"R\n\rRuntimeStatus\x12\x1a\n\x16RUNTIME_STATUS_UNKNOWN\x10\x00\x12\x0b\n\x07PENDING\x10\x01\x12\x0b\n\x07SUCCESS\x10\x02\x12\x0b\n\x07\x46\x41ILURE\x10\x03\"\xab\x04\n\x1d\x44\x65siredStateStatusChangeEvent\x12\x34\n\x02id\x18\x01 \x01(\x0b\x32(.prodvana.desired_state.model.Identifier\x12\x18\n\x10\x64\x65sired_state_id\x18\x02 \x01(\t\x12\x1d\n\x15root_desired_state_id\x18\x03 \x01(\t\x12\x38\n\nold_status\x18\x04 \x01(\x0e\x32$.prodvana.desired_state.model.Status\x12\x38\n\nnew_status\x18\x05 \x01(\x0e\x32$.prodvana.desired_state.model.Status\x12\x34\n\x07\x64\x65sired\x18\x06 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x33\n\x06target\x18\x07 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x36\n\told_state\x18\x08 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x36\n\tnew_state\x18\t \x01(\x0b\x32#.prodvana.desired_state.model.State\x12L\n\x13status_explanations\x18\n \x03(\x0b\x32/.prodvana.desired_state.model.StatusExplanation\"\xc6\x04\n\x18KeyDeliveryDecisionEvent\x12\x34\n\x02id\x18\x01 \x01(\x0b\x32(.prodvana.desired_state.model.Identifier\x12\x18\n\x10\x64\x65sired_state_id\x18\x02 \x01(\t\x12\x1d\n\x15root_desired_state_id\x18\x03 \x01(\t\x12\x44\n\x08\x64\x65\x63ision\x18\x04 \x01(\x0e\x32\x32.prodvana.events.KeyDeliveryDecisionEvent.Decision\x12\x13\n\x0b\x65xplanation\x18\x05 \x01(\t\x12\x34\n\x06status\x18\x06 \x01(\x0e\x32$.prodvana.desired_state.model.Status\x12\x34\n\x07\x64\x65sired\x18\x07 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x33\n\x06target\x18\x08 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x34\n\x07\x63urrent\x18\t \x01(\x0b\x32#.prodvana.desired_state.model.State\x12L\n\x13status_explanations\x18\n \x03(\x0b\x32/.prodvana.desired_state.model.StatusExplanation\";\n\x08\x44\x65\x63ision\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x16\n\x12ROLLBACK_INITIATED\x10\x01\x12\n\n\x06\x46\x41ILED\x10\x02\"\xe2\x01\n\x0cRpcCallEvent\x12\x30\n\x04type\x18\x01 \x01(\x0e\x32\".prodvana.events.RpcCallEvent.Type\x12\x13\n\x0brpc_service\x18\x02 \x01(\t\x12\x12\n\nrpc_method\x18\x03 \x01(\t\x12%\n\x07request\x18\x04 \x01(\x0b\x32\x14.google.protobuf.Any\x12&\n\x08response\x18\x05 \x01(\x0b\x32\x14.google.protobuf.Any\"(\n\x04Type\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x08\n\x04READ\x10\x01\x12\t\n\x05WRITE\x10\x02\"-\n\x11\x41pplicationHandle\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\n\n\x02id\x18\x02 \x01(\t\"i\n\x14ReleaseChannelHandle\x12\x37\n\x0b\x61pplication\x18\x01 \x01(\x0b\x32\".prodvana.events.ApplicationHandle\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\n\n\x02id\x18\x03 \x01(\t\"\xa2\x01\n\rServiceHandle\x12\x37\n\x0b\x61pplication\x18\x01 \x01(\x0b\x32\".prodvana.events.ApplicationHandle\x12>\n\x0frelease_channel\x18\x02 \x01(\x0b\x32%.prodvana.events.ReleaseChannelHandle\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\n\n\x02id\x18\x04 \x01(\t\")\n\rRuntimeHandle\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\n\n\x02id\x18\x02 \x01(\t\",\n\x10ProtectionHandle\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\n\n\x02id\x18\x02 \x01(\t\"\xa8\x03\n\x15PermissionDeniedEvent\x12=\n\x06\x61\x63tion\x18\x01 \x01(\x0e\x32-.prodvana.events.PermissionDeniedEvent.Action\x12\x39\n\x0b\x61pplication\x18\x02 \x01(\x0b\x32\".prodvana.events.ApplicationHandleH\x00\x12@\n\x0frelease_channel\x18\x03 \x01(\x0b\x32%.prodvana.events.ReleaseChannelHandleH\x00\x12\x31\n\x07service\x18\x04 \x01(\x0b\x32\x1e.prodvana.events.ServiceHandleH\x00\x12\x31\n\x07runtime\x18\x05 \x01(\x0b\x32\x1e.prodvana.events.RuntimeHandleH\x00\x12\x37\n\nprotection\x18\x06 \x01(\x0b\x32!.prodvana.events.ProtectionHandleH\x00\"*\n\x06\x41\x63tion\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x08\n\x04READ\x10\x01\x12\t\n\x05WRITE\x10\x02\x42\x08\n\x06target\"\xd2\x02\n\x12ObjectCreatedEvent\x12\x16\n\x0e\x63onfig_version\x18\x01 \x01(\t\x12\x39\n\x0b\x61pplication\x18\x02 \x01(\x0b\x32\".prodvana.events.ApplicationHandleH\x00\x12@\n\x0frelease_channel\x18\x03 \x01(\x0b\x32%.prodvana.events.ReleaseChannelHandleH\x00\x12\x31\n\x07service\x18\x04 \x01(\x0b\x32\x1e.prodvana.events.ServiceHandleH\x00\x12\x31\n\x07runtime\x18\x05 \x01(\x0b\x32\x1e.prodvana.events.RuntimeHandleH\x00\x12\x37\n\nprotection\x18\x06 \x01(\x0b\x32!.prodvana.events.ProtectionHandleH\x00\x42\x08\n\x06target\"\xd2\x02\n\x12ObjectDeletedEvent\x12\x16\n\x0e\x63onfig_version\x18\x01 \x01(\t\x12\x39\n\x0b\x61pplication\x18\x02 \x01(\x0b\x32\".prodvana.events.ApplicationHandleH\x00\x12@\n\x0frelease_channel\x18\x03 \x01(\x0b\x32%.prodvana.events.ReleaseChannelHandleH\x00\x12\x31\n\x07service\x18\x04 \x01(\x0b\x32\x1e.prodvana.events.ServiceHandleH\x00\x12\x31\n\x07runtime\x18\x05 \x01(\x0b\x32\x1e.prodvana.events.RuntimeHandleH\x00\x12\x37\n\nprotection\x18\x06 \x01(\x0b\x32!.prodvana.events.ProtectionHandleH\x00\x42\x08\n\x06target\"\xf3\x02\n\x13ObjectModifiedEvent\x12\x1a\n\x12old_config_version\x18\x01 \x01(\t\x12\x1a\n\x12new_config_version\x18\x02 \x01(\t\x12\x39\n\x0b\x61pplication\x18\x03 \x01(\x0b\x32\".prodvana.events.ApplicationHandleH\x00\x12@\n\x0frelease_channel\x18\x04 \x01(\x0b\x32%.prodvana.events.ReleaseChannelHandleH\x00\x12\x31\n\x07service\x18\x05 \x01(\x0b\x32\x1e.prodvana.events.ServiceHandleH\x00\x12\x31\n\x07runtime\x18\x06 \x01(\x0b\x32\x1e.prodvana.events.RuntimeHandleH\x00\x12\x37\n\nprotection\x18\x07 \x01(\x0b\x32!.prodvana.events.ProtectionHandleH\x00\x42\x08\n\x06target\"S\n\x16\x44\x65ploymentCreatedEvent\x12\x39\n\ndeployment\x18\x01 \x01(\x0b\x32%.prodvana.deployment.model.Deployment\"\x94\x01\n\x16\x44\x65ploymentUpdatedEvent\x12\x39\n\ndeployment\x18\x01 \x01(\x0b\x32%.prodvana.deployment.model.Deployment\x12?\n\nold_status\x18\x02 \x01(\x0e\x32+.prodvana.deployment.model.DeploymentStatus\"C\n\x0fKubectlCmdEvent\x12\x0f\n\x07runtime\x18\x01 \x01(\t\x12\x0c\n\x04\x61rgs\x18\x02 \x03(\t\x12\x11\n\texit_code\x18\x03 \x01(\x05\"\x87\x01\n\x12\x41uditLogDebugEvent\x12\x41\n\x07\x64\x65tails\x18\x01 \x03(\x0b\x32\x30.prodvana.events.AuditLogDebugEvent.DetailsEntry\x1a.\n\x0c\x44\x65tailsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\xe5\n\n\x0c\x45ventDetails\x12\x42\n\x11set_desired_state\x18\x01 \x01(\x0b\x32%.prodvana.events.SetDesiredStateEventH\x00\x12@\n\x10set_target_state\x18\x02 \x01(\x0b\x32$.prodvana.events.SetTargetStateEventH\x00\x12\x39\n\x0cprogram_exit\x18\x03 \x01(\x0b\x32!.prodvana.events.ProgramExitEventH\x00\x12\x44\n\x12\x61pply_target_state\x18\x04 \x01(\x0b\x32&.prodvana.events.ApplyTargetStateEventH\x00\x12?\n\x0fmanual_approval\x18\x06 \x01(\x0b\x32$.prodvana.events.ManualApprovalEventH\x00\x12J\n\x15\x63ustom_task_execution\x18\x07 \x01(\x0b\x32).prodvana.events.CustomTaskExecutionEventH\x00\x12=\n\x0eruntime_update\x18\x08 \x01(\x0b\x32#.prodvana.events.RuntimeUpdateEventH\x00\x12\x43\n\x11\x64\x65livery_progress\x18\t \x01(\x0b\x32&.prodvana.events.DeliveryProgressEventH\x00\x12K\n\x12\x64\x65livery_promotion\x18\n \x01(\x0b\x32-.prodvana.events.DeliveryManualPromotionEventH\x00\x12U\n\x1b\x64\x65sired_state_status_change\x18\x0b \x01(\x0b\x32..prodvana.events.DesiredStateStatusChangeEventH\x00\x12J\n\x15key_delivery_decision\x18\x0c \x01(\x0b\x32).prodvana.events.KeyDeliveryDecisionEventH\x00\x12\x31\n\x08rpc_call\x18\r \x01(\x0b\x32\x1d.prodvana.events.RpcCallEventH\x00\x12\x43\n\x11permission_denied\x18\x0e \x01(\x0b\x32&.prodvana.events.PermissionDeniedEventH\x00\x12=\n\x0eobject_created\x18\x0f \x01(\x0b\x32#.prodvana.events.ObjectCreatedEventH\x00\x12=\n\x0eobject_deleted\x18\x10 \x01(\x0b\x32#.prodvana.events.ObjectDeletedEventH\x00\x12?\n\x0fobject_modified\x18\x11 \x01(\x0b\x32$.prodvana.events.ObjectModifiedEventH\x00\x12\x37\n\x0bkubectl_cmd\x18\x12 \x01(\x0b\x32 .prodvana.events.KubectlCmdEventH\x00\x12>\n\x0f\x61udit_log_debug\x18\x13 \x01(\x0b\x32#.prodvana.events.AuditLogDebugEventH\x00\x12\x45\n\x12\x64\x65ployment_created\x18\x14 \x01(\x0b\x32\'.prodvana.events.DeploymentCreatedEventH\x00\x12\x45\n\x12\x64\x65ployment_updated\x18\x15 \x01(\x0b\x32\'.prodvana.events.DeploymentUpdatedEventH\x00\x42\t\n\x07\x64\x65tailsJ\x04\x08\x05\x10\x06*\xe2\x03\n\tEventType\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x15\n\x11SET_DESIRED_STATE\x10\x01\x12\x14\n\x10SET_TARGET_STATE\x10\x02\x12\x10\n\x0cPROGRAM_EXIT\x10\x03\x12\x16\n\x12\x41PPLY_TARGET_STATE\x10\x04\x12\x13\n\x0fMANUAL_APPROVAL\x10\x06\x12\x19\n\x15\x43USTOM_TASK_EXECUTION\x10\x07\x12\x12\n\x0eRUNTIME_UPDATE\x10\x08\x12\x15\n\x11\x44\x45LIVERY_PROGRESS\x10\t\x12\x1d\n\x19\x44\x45LIVERY_MANUAL_PROMOTION\x10\n\x12\x1f\n\x1b\x44\x45SIRED_STATE_STATUS_CHANGE\x10\x0b\x12\x19\n\x15KEY_DELIVERY_DECISION\x10\x0c\x12\x0c\n\x08RPC_CALL\x10\r\x12\x14\n\x10PERMISSON_DENIED\x10\x0e\x12\x12\n\x0eOBJECT_CREATED\x10\x0f\x12\x12\n\x0eOBJECT_DELETED\x10\x10\x12\x13\n\x0fOBJECT_MODIFIED\x10\x11\x12\x0f\n\x0bKUBECTL_CMD\x10\x12\x12\x13\n\x0f\x41UDIT_LOG_DEBUG\x10\x13\x12\x16\n\x12\x44\x45PLOYMENT_CREATED\x10\x14\x12\x16\n\x12\x44\x45PLOYMENT_UPDATED\x10\x15\"\x04\x08\x05\x10\x05\x42KZIgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/eventsb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'prodvana.events.types_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZIgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/events'
  _AUDITLOGDEBUGEVENT_DETAILSENTRY._options = None
  _AUDITLOGDEBUGEVENT_DETAILSENTRY._serialized_options = b'8\001'
  _globals['_EVENTTYPE']._serialized_start=7465
  _globals['_EVENTTYPE']._serialized_end=7947
  _globals['_SETDESIREDSTATEEVENT']._serialized_start=201
  _globals['_SETDESIREDSTATEEVENT']._serialized_end=409
  _globals['_SETTARGETSTATEEVENT']._serialized_start=412
  _globals['_SETTARGETSTATEEVENT']._serialized_end=863
  _globals['_APPLYTARGETSTATEEVENT']._serialized_start=866
  _globals['_APPLYTARGETSTATEEVENT']._serialized_end=1294
  _globals['_APPLYTARGETSTATEEVENT_APPLYRESULT']._serialized_start=1218
  _globals['_APPLYTARGETSTATEEVENT_APPLYRESULT']._serialized_end=1294
  _globals['_PROGRAMEXITEVENT']._serialized_start=1297
  _globals['_PROGRAMEXITEVENT']._serialized_end=1530
  _globals['_DELIVERYPROGRESSEVENT']._serialized_start=1532
  _globals['_DELIVERYPROGRESSEVENT']._serialized_end=1615
  _globals['_DELIVERYMANUALPROMOTIONEVENT']._serialized_start=1617
  _globals['_DELIVERYMANUALPROMOTIONEVENT']._serialized_end=1736
  _globals['_MANUALAPPROVALEVENT']._serialized_start=1738
  _globals['_MANUALAPPROVALEVENT']._serialized_end=1842
  _globals['_CUSTOMTASKEXECUTIONEVENT']._serialized_start=1844
  _globals['_CUSTOMTASKEXECUTIONEVENT']._serialized_end=1936
  _globals['_RUNTIMEOBJECT']._serialized_start=1938
  _globals['_RUNTIMEOBJECT']._serialized_end=2013
  _globals['_RUNTIMEUPDATEEVENT']._serialized_start=2016
  _globals['_RUNTIMEUPDATEEVENT']._serialized_end=2372
  _globals['_RUNTIMEUPDATEEVENT_RUNTIMEACTION']._serialized_start=2220
  _globals['_RUNTIMEUPDATEEVENT_RUNTIMEACTION']._serialized_end=2288
  _globals['_RUNTIMEUPDATEEVENT_RUNTIMESTATUS']._serialized_start=2290
  _globals['_RUNTIMEUPDATEEVENT_RUNTIMESTATUS']._serialized_end=2372
  _globals['_DESIREDSTATESTATUSCHANGEEVENT']._serialized_start=2375
  _globals['_DESIREDSTATESTATUSCHANGEEVENT']._serialized_end=2930
  _globals['_KEYDELIVERYDECISIONEVENT']._serialized_start=2933
  _globals['_KEYDELIVERYDECISIONEVENT']._serialized_end=3515
  _globals['_KEYDELIVERYDECISIONEVENT_DECISION']._serialized_start=3456
  _globals['_KEYDELIVERYDECISIONEVENT_DECISION']._serialized_end=3515
  _globals['_RPCCALLEVENT']._serialized_start=3518
  _globals['_RPCCALLEVENT']._serialized_end=3744
  _globals['_RPCCALLEVENT_TYPE']._serialized_start=3704
  _globals['_RPCCALLEVENT_TYPE']._serialized_end=3744
  _globals['_APPLICATIONHANDLE']._serialized_start=3746
  _globals['_APPLICATIONHANDLE']._serialized_end=3791
  _globals['_RELEASECHANNELHANDLE']._serialized_start=3793
  _globals['_RELEASECHANNELHANDLE']._serialized_end=3898
  _globals['_SERVICEHANDLE']._serialized_start=3901
  _globals['_SERVICEHANDLE']._serialized_end=4063
  _globals['_RUNTIMEHANDLE']._serialized_start=4065
  _globals['_RUNTIMEHANDLE']._serialized_end=4106
  _globals['_PROTECTIONHANDLE']._serialized_start=4108
  _globals['_PROTECTIONHANDLE']._serialized_end=4152
  _globals['_PERMISSIONDENIEDEVENT']._serialized_start=4155
  _globals['_PERMISSIONDENIEDEVENT']._serialized_end=4579
  _globals['_PERMISSIONDENIEDEVENT_ACTION']._serialized_start=4527
  _globals['_PERMISSIONDENIEDEVENT_ACTION']._serialized_end=4569
  _globals['_OBJECTCREATEDEVENT']._serialized_start=4582
  _globals['_OBJECTCREATEDEVENT']._serialized_end=4920
  _globals['_OBJECTDELETEDEVENT']._serialized_start=4923
  _globals['_OBJECTDELETEDEVENT']._serialized_end=5261
  _globals['_OBJECTMODIFIEDEVENT']._serialized_start=5264
  _globals['_OBJECTMODIFIEDEVENT']._serialized_end=5635
  _globals['_DEPLOYMENTCREATEDEVENT']._serialized_start=5637
  _globals['_DEPLOYMENTCREATEDEVENT']._serialized_end=5720
  _globals['_DEPLOYMENTUPDATEDEVENT']._serialized_start=5723
  _globals['_DEPLOYMENTUPDATEDEVENT']._serialized_end=5871
  _globals['_KUBECTLCMDEVENT']._serialized_start=5873
  _globals['_KUBECTLCMDEVENT']._serialized_end=5940
  _globals['_AUDITLOGDEBUGEVENT']._serialized_start=5943
  _globals['_AUDITLOGDEBUGEVENT']._serialized_end=6078
  _globals['_AUDITLOGDEBUGEVENT_DETAILSENTRY']._serialized_start=6032
  _globals['_AUDITLOGDEBUGEVENT_DETAILSENTRY']._serialized_end=6078
  _globals['_EVENTDETAILS']._serialized_start=6081
  _globals['_EVENTDETAILS']._serialized_end=7462
# @@protoc_insertion_point(module_scope)
