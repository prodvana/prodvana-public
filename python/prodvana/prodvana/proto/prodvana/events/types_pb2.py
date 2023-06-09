# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/events/types.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from prodvana.proto.prodvana.desired_state.model import desired_state_pb2 as prodvana_dot_desired__state_dot_model_dot_desired__state__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1bprodvana/events/types.proto\x12\x0fprodvana.events\x1a\x30prodvana/desired_state/model/desired_state.proto\"\xd0\x01\n\x14SetDesiredStateEvent\x12\x34\n\x02id\x18\x01 \x01(\x0b\x32(.prodvana.desired_state.model.Identifier\x12\x18\n\x10\x64\x65sired_state_id\x18\x03 \x01(\t\x12\x1d\n\x15root_desired_state_id\x18\x04 \x01(\t\x12\x34\n\x07\x64\x65sired\x18\x02 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x13\n\x0bis_rollback\x18\x05 \x01(\x08\"\xc3\x03\n\x13SetTargetStateEvent\x12\x34\n\x02id\x18\x01 \x01(\x0b\x32(.prodvana.desired_state.model.Identifier\x12\x18\n\x10\x64\x65sired_state_id\x18\x08 \x01(\t\x12\x1d\n\x15root_desired_state_id\x18\t \x01(\t\x12\x33\n\x06target\x18\x02 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x34\n\x07\x64\x65sired\x18\x03 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x34\n\x07\x63urrent\x18\x04 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x18\n\x10is_auto_rollback\x18\x05 \x01(\x08\x12\x34\n\x06status\x18\x06 \x01(\x0e\x32$.prodvana.desired_state.model.Status\x12L\n\x13status_explanations\x18\x07 \x03(\x0b\x32/.prodvana.desired_state.model.StatusExplanation\"\x95\x03\n\x15\x41pplyTargetStateEvent\x12\x34\n\x02id\x18\x01 \x01(\x0b\x32(.prodvana.desired_state.model.Identifier\x12\x18\n\x10\x64\x65sired_state_id\x18\x06 \x01(\t\x12\x1d\n\x15root_desired_state_id\x18\x07 \x01(\t\x12\x33\n\x06target\x18\x02 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x34\n\x07\x63urrent\x18\x03 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x18\n\x10is_auto_rollback\x18\x04 \x01(\x08\x12\x42\n\x06result\x18\x05 \x01(\x0e\x32\x32.prodvana.events.ApplyTargetStateEvent.ApplyResult\x12\r\n\x05\x65rror\x18\x08 \x01(\t\"5\n\x0b\x41pplyResult\x12\x0b\n\x07UNKNOWN\x10\x00\x12\r\n\tSUCCEEDED\x10\x01\x12\n\n\x06\x46\x41ILED\x10\x02\"\xb0\x01\n\x10ProgramExitEvent\x12\x11\n\tnamespace\x18\x01 \x01(\t\x12\x0b\n\x03pod\x18\x02 \x01(\t\x12\x0f\n\x07program\x18\x03 \x01(\t\x12\x15\n\rrestart_count\x18\x04 \x01(\x05\x12\x0e\n\x06reason\x18\x05 \x01(\t\x12\x15\n\rkilled_reason\x18\t \x01(\t\x12\x11\n\texit_code\x18\x06 \x01(\x05\x12\x0e\n\x06signal\x18\x07 \x01(\x05J\x04\x08\x08\x10\tR\x04logs\"S\n\x15\x44\x65liveryProgressEvent\x12:\n\x05state\x18\x01 \x01(\x0b\x32+.prodvana.desired_state.model.DeliveryState\"w\n\x1c\x44\x65liveryManualPromotionEvent\x12:\n\x05state\x18\x01 \x01(\x0b\x32+.prodvana.desired_state.model.DeliveryState\x12\r\n\x05stage\x18\x02 \x01(\x03\x12\x0c\n\x04\x66ull\x18\x03 \x01(\x08\"h\n\x13ManualApprovalEvent\x12\x42\n\x06status\x18\x01 \x01(\x0e\x32\x32.prodvana.desired_state.model.ManualApprovalStatus\x12\r\n\x05topic\x18\x02 \x01(\t\"\\\n\x18\x43ustomTaskExecutionEvent\x12\x12\n\nsuccessful\x18\x01 \x01(\x08\x12\x10\n\x08\x61ttempts\x18\x02 \x01(\x05\x12\x0e\n\x06reason\x18\x03 \x01(\tJ\x04\x08\x04\x10\x05R\x04logs\"K\n\rRuntimeObject\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0c\n\x04kind\x18\x02 \x01(\t\x12\x11\n\tnamespace\x18\x03 \x01(\t\x12\x0b\n\x03url\x18\x04 \x01(\t\"\xe4\x02\n\x12RuntimeUpdateEvent\x12\x41\n\x06\x61\x63tion\x18\x01 \x01(\x0e\x32\x31.prodvana.events.RuntimeUpdateEvent.RuntimeAction\x12.\n\x06object\x18\x02 \x01(\x0b\x32\x1e.prodvana.events.RuntimeObject\x12\x41\n\x06status\x18\x03 \x01(\x0e\x32\x31.prodvana.events.RuntimeUpdateEvent.RuntimeStatus\"D\n\rRuntimeAction\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x14\n\x10\x43REATE_OR_UPDATE\x10\x01\x12\x10\n\x0cWAIT_HEALTHY\x10\x02\"R\n\rRuntimeStatus\x12\x1a\n\x16RUNTIME_STATUS_UNKNOWN\x10\x00\x12\x0b\n\x07PENDING\x10\x01\x12\x0b\n\x07SUCCESS\x10\x02\x12\x0b\n\x07\x46\x41ILURE\x10\x03\"\xab\x04\n\x1d\x44\x65siredStateStatusChangeEvent\x12\x34\n\x02id\x18\x01 \x01(\x0b\x32(.prodvana.desired_state.model.Identifier\x12\x18\n\x10\x64\x65sired_state_id\x18\x02 \x01(\t\x12\x1d\n\x15root_desired_state_id\x18\x03 \x01(\t\x12\x38\n\nold_status\x18\x04 \x01(\x0e\x32$.prodvana.desired_state.model.Status\x12\x38\n\nnew_status\x18\x05 \x01(\x0e\x32$.prodvana.desired_state.model.Status\x12\x34\n\x07\x64\x65sired\x18\x06 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x33\n\x06target\x18\x07 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x36\n\told_state\x18\x08 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x36\n\tnew_state\x18\t \x01(\x0b\x32#.prodvana.desired_state.model.State\x12L\n\x13status_explanations\x18\n \x03(\x0b\x32/.prodvana.desired_state.model.StatusExplanation\"\xc6\x04\n\x18KeyDeliveryDecisionEvent\x12\x34\n\x02id\x18\x01 \x01(\x0b\x32(.prodvana.desired_state.model.Identifier\x12\x18\n\x10\x64\x65sired_state_id\x18\x02 \x01(\t\x12\x1d\n\x15root_desired_state_id\x18\x03 \x01(\t\x12\x44\n\x08\x64\x65\x63ision\x18\x04 \x01(\x0e\x32\x32.prodvana.events.KeyDeliveryDecisionEvent.Decision\x12\x13\n\x0b\x65xplanation\x18\x05 \x01(\t\x12\x34\n\x06status\x18\x06 \x01(\x0e\x32$.prodvana.desired_state.model.Status\x12\x34\n\x07\x64\x65sired\x18\x07 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x33\n\x06target\x18\x08 \x01(\x0b\x32#.prodvana.desired_state.model.State\x12\x34\n\x07\x63urrent\x18\t \x01(\x0b\x32#.prodvana.desired_state.model.State\x12L\n\x13status_explanations\x18\n \x03(\x0b\x32/.prodvana.desired_state.model.StatusExplanation\";\n\x08\x44\x65\x63ision\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x16\n\x12ROLLBACK_INITIATED\x10\x01\x12\n\n\x06\x46\x41ILED\x10\x02\"\xa7\x06\n\x0c\x45ventDetails\x12\x42\n\x11set_desired_state\x18\x01 \x01(\x0b\x32%.prodvana.events.SetDesiredStateEventH\x00\x12@\n\x10set_target_state\x18\x02 \x01(\x0b\x32$.prodvana.events.SetTargetStateEventH\x00\x12\x39\n\x0cprogram_exit\x18\x03 \x01(\x0b\x32!.prodvana.events.ProgramExitEventH\x00\x12\x44\n\x12\x61pply_target_state\x18\x04 \x01(\x0b\x32&.prodvana.events.ApplyTargetStateEventH\x00\x12?\n\x0fmanual_approval\x18\x06 \x01(\x0b\x32$.prodvana.events.ManualApprovalEventH\x00\x12J\n\x15\x63ustom_task_execution\x18\x07 \x01(\x0b\x32).prodvana.events.CustomTaskExecutionEventH\x00\x12=\n\x0eruntime_update\x18\x08 \x01(\x0b\x32#.prodvana.events.RuntimeUpdateEventH\x00\x12\x43\n\x11\x64\x65livery_progress\x18\t \x01(\x0b\x32&.prodvana.events.DeliveryProgressEventH\x00\x12K\n\x12\x64\x65livery_promotion\x18\n \x01(\x0b\x32-.prodvana.events.DeliveryManualPromotionEventH\x00\x12U\n\x1b\x64\x65sired_state_status_change\x18\x0b \x01(\x0b\x32..prodvana.events.DesiredStateStatusChangeEventH\x00\x12J\n\x15key_delivery_decision\x18\x0c \x01(\x0b\x32).prodvana.events.KeyDeliveryDecisionEventH\x00\x42\t\n\x07\x64\x65tailsJ\x04\x08\x05\x10\x06*\xab\x02\n\tEventType\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x15\n\x11SET_DESIRED_STATE\x10\x01\x12\x14\n\x10SET_TARGET_STATE\x10\x02\x12\x10\n\x0cPROGRAM_EXIT\x10\x03\x12\x16\n\x12\x41PPLY_TARGET_STATE\x10\x04\x12\x13\n\x0fMANUAL_APPROVAL\x10\x06\x12\x19\n\x15\x43USTOM_TASK_EXECUTION\x10\x07\x12\x12\n\x0eRUNTIME_UPDATE\x10\x08\x12\x15\n\x11\x44\x45LIVERY_PROGRESS\x10\t\x12\x1d\n\x19\x44\x45LIVERY_MANUAL_PROMOTION\x10\n\x12\x1f\n\x1b\x44\x45SIRED_STATE_STATUS_CHANGE\x10\x0b\x12\x19\n\x15KEY_DELIVERY_DECISION\x10\x0c\"\x04\x08\x05\x10\x05\x42KZIgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/eventsb\x06proto3')

_EVENTTYPE = DESCRIPTOR.enum_types_by_name['EventType']
EventType = enum_type_wrapper.EnumTypeWrapper(_EVENTTYPE)
UNKNOWN = 0
SET_DESIRED_STATE = 1
SET_TARGET_STATE = 2
PROGRAM_EXIT = 3
APPLY_TARGET_STATE = 4
MANUAL_APPROVAL = 6
CUSTOM_TASK_EXECUTION = 7
RUNTIME_UPDATE = 8
DELIVERY_PROGRESS = 9
DELIVERY_MANUAL_PROMOTION = 10
DESIRED_STATE_STATUS_CHANGE = 11
KEY_DELIVERY_DECISION = 12


_SETDESIREDSTATEEVENT = DESCRIPTOR.message_types_by_name['SetDesiredStateEvent']
_SETTARGETSTATEEVENT = DESCRIPTOR.message_types_by_name['SetTargetStateEvent']
_APPLYTARGETSTATEEVENT = DESCRIPTOR.message_types_by_name['ApplyTargetStateEvent']
_PROGRAMEXITEVENT = DESCRIPTOR.message_types_by_name['ProgramExitEvent']
_DELIVERYPROGRESSEVENT = DESCRIPTOR.message_types_by_name['DeliveryProgressEvent']
_DELIVERYMANUALPROMOTIONEVENT = DESCRIPTOR.message_types_by_name['DeliveryManualPromotionEvent']
_MANUALAPPROVALEVENT = DESCRIPTOR.message_types_by_name['ManualApprovalEvent']
_CUSTOMTASKEXECUTIONEVENT = DESCRIPTOR.message_types_by_name['CustomTaskExecutionEvent']
_RUNTIMEOBJECT = DESCRIPTOR.message_types_by_name['RuntimeObject']
_RUNTIMEUPDATEEVENT = DESCRIPTOR.message_types_by_name['RuntimeUpdateEvent']
_DESIREDSTATESTATUSCHANGEEVENT = DESCRIPTOR.message_types_by_name['DesiredStateStatusChangeEvent']
_KEYDELIVERYDECISIONEVENT = DESCRIPTOR.message_types_by_name['KeyDeliveryDecisionEvent']
_EVENTDETAILS = DESCRIPTOR.message_types_by_name['EventDetails']
_APPLYTARGETSTATEEVENT_APPLYRESULT = _APPLYTARGETSTATEEVENT.enum_types_by_name['ApplyResult']
_RUNTIMEUPDATEEVENT_RUNTIMEACTION = _RUNTIMEUPDATEEVENT.enum_types_by_name['RuntimeAction']
_RUNTIMEUPDATEEVENT_RUNTIMESTATUS = _RUNTIMEUPDATEEVENT.enum_types_by_name['RuntimeStatus']
_KEYDELIVERYDECISIONEVENT_DECISION = _KEYDELIVERYDECISIONEVENT.enum_types_by_name['Decision']
SetDesiredStateEvent = _reflection.GeneratedProtocolMessageType('SetDesiredStateEvent', (_message.Message,), {
  'DESCRIPTOR' : _SETDESIREDSTATEEVENT,
  '__module__' : 'prodvana.events.types_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.events.SetDesiredStateEvent)
  })
_sym_db.RegisterMessage(SetDesiredStateEvent)

SetTargetStateEvent = _reflection.GeneratedProtocolMessageType('SetTargetStateEvent', (_message.Message,), {
  'DESCRIPTOR' : _SETTARGETSTATEEVENT,
  '__module__' : 'prodvana.events.types_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.events.SetTargetStateEvent)
  })
_sym_db.RegisterMessage(SetTargetStateEvent)

ApplyTargetStateEvent = _reflection.GeneratedProtocolMessageType('ApplyTargetStateEvent', (_message.Message,), {
  'DESCRIPTOR' : _APPLYTARGETSTATEEVENT,
  '__module__' : 'prodvana.events.types_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.events.ApplyTargetStateEvent)
  })
_sym_db.RegisterMessage(ApplyTargetStateEvent)

ProgramExitEvent = _reflection.GeneratedProtocolMessageType('ProgramExitEvent', (_message.Message,), {
  'DESCRIPTOR' : _PROGRAMEXITEVENT,
  '__module__' : 'prodvana.events.types_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.events.ProgramExitEvent)
  })
_sym_db.RegisterMessage(ProgramExitEvent)

DeliveryProgressEvent = _reflection.GeneratedProtocolMessageType('DeliveryProgressEvent', (_message.Message,), {
  'DESCRIPTOR' : _DELIVERYPROGRESSEVENT,
  '__module__' : 'prodvana.events.types_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.events.DeliveryProgressEvent)
  })
_sym_db.RegisterMessage(DeliveryProgressEvent)

DeliveryManualPromotionEvent = _reflection.GeneratedProtocolMessageType('DeliveryManualPromotionEvent', (_message.Message,), {
  'DESCRIPTOR' : _DELIVERYMANUALPROMOTIONEVENT,
  '__module__' : 'prodvana.events.types_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.events.DeliveryManualPromotionEvent)
  })
_sym_db.RegisterMessage(DeliveryManualPromotionEvent)

ManualApprovalEvent = _reflection.GeneratedProtocolMessageType('ManualApprovalEvent', (_message.Message,), {
  'DESCRIPTOR' : _MANUALAPPROVALEVENT,
  '__module__' : 'prodvana.events.types_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.events.ManualApprovalEvent)
  })
_sym_db.RegisterMessage(ManualApprovalEvent)

CustomTaskExecutionEvent = _reflection.GeneratedProtocolMessageType('CustomTaskExecutionEvent', (_message.Message,), {
  'DESCRIPTOR' : _CUSTOMTASKEXECUTIONEVENT,
  '__module__' : 'prodvana.events.types_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.events.CustomTaskExecutionEvent)
  })
_sym_db.RegisterMessage(CustomTaskExecutionEvent)

RuntimeObject = _reflection.GeneratedProtocolMessageType('RuntimeObject', (_message.Message,), {
  'DESCRIPTOR' : _RUNTIMEOBJECT,
  '__module__' : 'prodvana.events.types_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.events.RuntimeObject)
  })
_sym_db.RegisterMessage(RuntimeObject)

RuntimeUpdateEvent = _reflection.GeneratedProtocolMessageType('RuntimeUpdateEvent', (_message.Message,), {
  'DESCRIPTOR' : _RUNTIMEUPDATEEVENT,
  '__module__' : 'prodvana.events.types_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.events.RuntimeUpdateEvent)
  })
_sym_db.RegisterMessage(RuntimeUpdateEvent)

DesiredStateStatusChangeEvent = _reflection.GeneratedProtocolMessageType('DesiredStateStatusChangeEvent', (_message.Message,), {
  'DESCRIPTOR' : _DESIREDSTATESTATUSCHANGEEVENT,
  '__module__' : 'prodvana.events.types_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.events.DesiredStateStatusChangeEvent)
  })
_sym_db.RegisterMessage(DesiredStateStatusChangeEvent)

KeyDeliveryDecisionEvent = _reflection.GeneratedProtocolMessageType('KeyDeliveryDecisionEvent', (_message.Message,), {
  'DESCRIPTOR' : _KEYDELIVERYDECISIONEVENT,
  '__module__' : 'prodvana.events.types_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.events.KeyDeliveryDecisionEvent)
  })
_sym_db.RegisterMessage(KeyDeliveryDecisionEvent)

EventDetails = _reflection.GeneratedProtocolMessageType('EventDetails', (_message.Message,), {
  'DESCRIPTOR' : _EVENTDETAILS,
  '__module__' : 'prodvana.events.types_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.events.EventDetails)
  })
_sym_db.RegisterMessage(EventDetails)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZIgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/events'
  _EVENTTYPE._serialized_start=4146
  _EVENTTYPE._serialized_end=4445
  _SETDESIREDSTATEEVENT._serialized_start=99
  _SETDESIREDSTATEEVENT._serialized_end=307
  _SETTARGETSTATEEVENT._serialized_start=310
  _SETTARGETSTATEEVENT._serialized_end=761
  _APPLYTARGETSTATEEVENT._serialized_start=764
  _APPLYTARGETSTATEEVENT._serialized_end=1169
  _APPLYTARGETSTATEEVENT_APPLYRESULT._serialized_start=1116
  _APPLYTARGETSTATEEVENT_APPLYRESULT._serialized_end=1169
  _PROGRAMEXITEVENT._serialized_start=1172
  _PROGRAMEXITEVENT._serialized_end=1348
  _DELIVERYPROGRESSEVENT._serialized_start=1350
  _DELIVERYPROGRESSEVENT._serialized_end=1433
  _DELIVERYMANUALPROMOTIONEVENT._serialized_start=1435
  _DELIVERYMANUALPROMOTIONEVENT._serialized_end=1554
  _MANUALAPPROVALEVENT._serialized_start=1556
  _MANUALAPPROVALEVENT._serialized_end=1660
  _CUSTOMTASKEXECUTIONEVENT._serialized_start=1662
  _CUSTOMTASKEXECUTIONEVENT._serialized_end=1754
  _RUNTIMEOBJECT._serialized_start=1756
  _RUNTIMEOBJECT._serialized_end=1831
  _RUNTIMEUPDATEEVENT._serialized_start=1834
  _RUNTIMEUPDATEEVENT._serialized_end=2190
  _RUNTIMEUPDATEEVENT_RUNTIMEACTION._serialized_start=2038
  _RUNTIMEUPDATEEVENT_RUNTIMEACTION._serialized_end=2106
  _RUNTIMEUPDATEEVENT_RUNTIMESTATUS._serialized_start=2108
  _RUNTIMEUPDATEEVENT_RUNTIMESTATUS._serialized_end=2190
  _DESIREDSTATESTATUSCHANGEEVENT._serialized_start=2193
  _DESIREDSTATESTATUSCHANGEEVENT._serialized_end=2748
  _KEYDELIVERYDECISIONEVENT._serialized_start=2751
  _KEYDELIVERYDECISIONEVENT._serialized_end=3333
  _KEYDELIVERYDECISIONEVENT_DECISION._serialized_start=3274
  _KEYDELIVERYDECISIONEVENT_DECISION._serialized_end=3333
  _EVENTDETAILS._serialized_start=3336
  _EVENTDETAILS._serialized_end=4143
# @@protoc_insertion_point(module_scope)
