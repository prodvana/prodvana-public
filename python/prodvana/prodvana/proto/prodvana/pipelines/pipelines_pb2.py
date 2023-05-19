# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/pipelines/pipelines.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from prodvana.proto.validate import validate_pb2 as validate_dot_validate__pb2
from prodvana.proto.prodvana.common_config import program_pb2 as prodvana_dot_common__config_dot_program__pb2
from prodvana.proto.prodvana.common_config import retry_pb2 as prodvana_dot_common__config_dot_retry__pb2
from prodvana.proto.prodvana.common_config import notification_pb2 as prodvana_dot_common__config_dot_notification__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\"prodvana/pipelines/pipelines.proto\x12\x12prodvana.pipelines\x1a\x1egoogle/protobuf/duration.proto\x1a\x17validate/validate.proto\x1a$prodvana/common_config/program.proto\x1a\"prodvana/common_config/retry.proto\x1a)prodvana/common_config/notification.proto\"\xa3\x01\n\x08PushTask\x12\x12\n\nservice_id\x18\x01 \x01(\t\x12\x0f\n\x07service\x18\x02 \x01(\t\x12\x1a\n\x12release_channel_id\x18\x03 \x01(\t\x12\x17\n\x0frelease_channel\x18\x04 \x01(\t\x12\x10\n\x08rollback\x18\x05 \x01(\x08\x12\x16\n\x0e\x61pplication_id\x18\x06 \x01(\t\x12\x13\n\x0b\x61pplication\x18\x07 \x01(\t\"C\n\x08WaitTask\x12\x37\n\x08\x64uration\x18\x01 \x01(\x0b\x32\x19.google.protobuf.DurationB\n\xfa\x42\x07\xaa\x01\x04\x08\x01*\x00\"\x80\x01\n\x0cParallelTask\x12:\n\x06tracks\x18\x01 \x03(\x0b\x32*.prodvana.pipelines.ParallelTask.TaskTrack\x1a\x34\n\tTaskTrack\x12\'\n\x05tasks\x18\x01 \x03(\x0b\x32\x18.prodvana.pipelines.Task\"\x14\n\x12ManualApprovalTask\"\xa9\x02\n\nCustomTask\x12\x1c\n\x0b\x64\x65scription\x18\x07 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12@\n\x07program\x18\x01 \x01(\x0b\x32%.prodvana.common_config.ProgramConfigB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12\x39\n\x0cretry_config\x18\x02 \x01(\x0b\x32#.prodvana.common_config.RetryConfig\x12\x13\n\x0b\x61pplication\x18\x03 \x01(\t\x12\x16\n\x0e\x61pplication_id\x18\x04 \x01(\t\x12\x17\n\x0frelease_channel\x18\x05 \x01(\t\x12\x1a\n\x12release_channel_id\x18\x06 \x01(\t\x12\x1e\n\x16\x66inal_compiled_program\x18\x08 \x01(\x08\"\x86\x03\n\x04Task\x12\x33\n\x08metadata\x18\x01 \x01(\x0b\x32!.prodvana.pipelines.Task.Metadata\x12\x31\n\tpush_task\x18\x02 \x01(\x0b\x32\x1c.prodvana.pipelines.PushTaskH\x00\x12\x31\n\twait_task\x18\x03 \x01(\x0b\x32\x1c.prodvana.pipelines.WaitTaskH\x00\x12\x39\n\rparallel_task\x18\x04 \x01(\x0b\x32 .prodvana.pipelines.ParallelTaskH\x00\x12\x46\n\x14manual_approval_task\x18\x05 \x01(\x0b\x32&.prodvana.pipelines.ManualApprovalTaskH\x00\x12\x35\n\x0b\x63ustom_task\x18\x06 \x01(\x0b\x32\x1e.prodvana.pipelines.CustomTaskH\x00\x1a\x16\n\x08Metadata\x12\n\n\x02id\x18\x01 \x01(\x05\x42\x11\n\ntask_oneof\x12\x03\xf8\x42\x01\"\xda\x01\n\x0ePipelineConfig\x12\x39\n\x04name\x18\x01 \x01(\tB+\xfa\x42(r&\x10\x01\x18?2 ^[a-z]([a-z0-9-]*[a-z0-9]){0,1}$\x12\x38\n\x05tasks\x18\x02 \x03(\x0b\x32\x18.prodvana.pipelines.TaskB\x0f\xfa\x42\x0c\x92\x01\t\x08\x01\"\x05\x8a\x01\x02\x10\x01\x12\x10\n\x08rollback\x18\x03 \x01(\x08\x12\x41\n\rnotifications\x18\x04 \x01(\x0b\x32*.prodvana.common_config.NotificationConfig\"\xde\x01\n\x10PipelineTemplate\x12;\n\x0bname_suffix\x18\x01 \x01(\tB&\xfa\x42#r!\x10\x01\x18?2\x1b^([a-z0-9-]*[a-z0-9]){0,1}$\x12\x38\n\x05tasks\x18\x02 \x03(\x0b\x32\x18.prodvana.pipelines.TaskB\x0f\xfa\x42\x0c\x92\x01\t\x08\x01\"\x05\x8a\x01\x02\x10\x01\x12\x10\n\x08rollback\x18\x03 \x01(\x08\x12\x41\n\rnotifications\x18\x04 \x01(\x0b\x32*.prodvana.common_config.NotificationConfigBNZLgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/pipelinesb\x06proto3')



_PUSHTASK = DESCRIPTOR.message_types_by_name['PushTask']
_WAITTASK = DESCRIPTOR.message_types_by_name['WaitTask']
_PARALLELTASK = DESCRIPTOR.message_types_by_name['ParallelTask']
_PARALLELTASK_TASKTRACK = _PARALLELTASK.nested_types_by_name['TaskTrack']
_MANUALAPPROVALTASK = DESCRIPTOR.message_types_by_name['ManualApprovalTask']
_CUSTOMTASK = DESCRIPTOR.message_types_by_name['CustomTask']
_TASK = DESCRIPTOR.message_types_by_name['Task']
_TASK_METADATA = _TASK.nested_types_by_name['Metadata']
_PIPELINECONFIG = DESCRIPTOR.message_types_by_name['PipelineConfig']
_PIPELINETEMPLATE = DESCRIPTOR.message_types_by_name['PipelineTemplate']
PushTask = _reflection.GeneratedProtocolMessageType('PushTask', (_message.Message,), {
  'DESCRIPTOR' : _PUSHTASK,
  '__module__' : 'prodvana.pipelines.pipelines_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.pipelines.PushTask)
  })
_sym_db.RegisterMessage(PushTask)

WaitTask = _reflection.GeneratedProtocolMessageType('WaitTask', (_message.Message,), {
  'DESCRIPTOR' : _WAITTASK,
  '__module__' : 'prodvana.pipelines.pipelines_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.pipelines.WaitTask)
  })
_sym_db.RegisterMessage(WaitTask)

ParallelTask = _reflection.GeneratedProtocolMessageType('ParallelTask', (_message.Message,), {

  'TaskTrack' : _reflection.GeneratedProtocolMessageType('TaskTrack', (_message.Message,), {
    'DESCRIPTOR' : _PARALLELTASK_TASKTRACK,
    '__module__' : 'prodvana.pipelines.pipelines_pb2'
    # @@protoc_insertion_point(class_scope:prodvana.pipelines.ParallelTask.TaskTrack)
    })
  ,
  'DESCRIPTOR' : _PARALLELTASK,
  '__module__' : 'prodvana.pipelines.pipelines_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.pipelines.ParallelTask)
  })
_sym_db.RegisterMessage(ParallelTask)
_sym_db.RegisterMessage(ParallelTask.TaskTrack)

ManualApprovalTask = _reflection.GeneratedProtocolMessageType('ManualApprovalTask', (_message.Message,), {
  'DESCRIPTOR' : _MANUALAPPROVALTASK,
  '__module__' : 'prodvana.pipelines.pipelines_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.pipelines.ManualApprovalTask)
  })
_sym_db.RegisterMessage(ManualApprovalTask)

CustomTask = _reflection.GeneratedProtocolMessageType('CustomTask', (_message.Message,), {
  'DESCRIPTOR' : _CUSTOMTASK,
  '__module__' : 'prodvana.pipelines.pipelines_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.pipelines.CustomTask)
  })
_sym_db.RegisterMessage(CustomTask)

Task = _reflection.GeneratedProtocolMessageType('Task', (_message.Message,), {

  'Metadata' : _reflection.GeneratedProtocolMessageType('Metadata', (_message.Message,), {
    'DESCRIPTOR' : _TASK_METADATA,
    '__module__' : 'prodvana.pipelines.pipelines_pb2'
    # @@protoc_insertion_point(class_scope:prodvana.pipelines.Task.Metadata)
    })
  ,
  'DESCRIPTOR' : _TASK,
  '__module__' : 'prodvana.pipelines.pipelines_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.pipelines.Task)
  })
_sym_db.RegisterMessage(Task)
_sym_db.RegisterMessage(Task.Metadata)

PipelineConfig = _reflection.GeneratedProtocolMessageType('PipelineConfig', (_message.Message,), {
  'DESCRIPTOR' : _PIPELINECONFIG,
  '__module__' : 'prodvana.pipelines.pipelines_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.pipelines.PipelineConfig)
  })
_sym_db.RegisterMessage(PipelineConfig)

PipelineTemplate = _reflection.GeneratedProtocolMessageType('PipelineTemplate', (_message.Message,), {
  'DESCRIPTOR' : _PIPELINETEMPLATE,
  '__module__' : 'prodvana.pipelines.pipelines_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.pipelines.PipelineTemplate)
  })
_sym_db.RegisterMessage(PipelineTemplate)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZLgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/pipelines'
  _WAITTASK.fields_by_name['duration']._options = None
  _WAITTASK.fields_by_name['duration']._serialized_options = b'\372B\007\252\001\004\010\001*\000'
  _CUSTOMTASK.fields_by_name['description']._options = None
  _CUSTOMTASK.fields_by_name['description']._serialized_options = b'\372B\004r\002\020\001'
  _CUSTOMTASK.fields_by_name['program']._options = None
  _CUSTOMTASK.fields_by_name['program']._serialized_options = b'\372B\005\212\001\002\020\001'
  _TASK.oneofs_by_name['task_oneof']._options = None
  _TASK.oneofs_by_name['task_oneof']._serialized_options = b'\370B\001'
  _PIPELINECONFIG.fields_by_name['name']._options = None
  _PIPELINECONFIG.fields_by_name['name']._serialized_options = b'\372B(r&\020\001\030?2 ^[a-z]([a-z0-9-]*[a-z0-9]){0,1}$'
  _PIPELINECONFIG.fields_by_name['tasks']._options = None
  _PIPELINECONFIG.fields_by_name['tasks']._serialized_options = b'\372B\014\222\001\t\010\001\"\005\212\001\002\020\001'
  _PIPELINETEMPLATE.fields_by_name['name_suffix']._options = None
  _PIPELINETEMPLATE.fields_by_name['name_suffix']._serialized_options = b'\372B#r!\020\001\030?2\033^([a-z0-9-]*[a-z0-9]){0,1}$'
  _PIPELINETEMPLATE.fields_by_name['tasks']._options = None
  _PIPELINETEMPLATE.fields_by_name['tasks']._serialized_options = b'\372B\014\222\001\t\010\001\"\005\212\001\002\020\001'
  _PUSHTASK._serialized_start=233
  _PUSHTASK._serialized_end=396
  _WAITTASK._serialized_start=398
  _WAITTASK._serialized_end=465
  _PARALLELTASK._serialized_start=468
  _PARALLELTASK._serialized_end=596
  _PARALLELTASK_TASKTRACK._serialized_start=544
  _PARALLELTASK_TASKTRACK._serialized_end=596
  _MANUALAPPROVALTASK._serialized_start=598
  _MANUALAPPROVALTASK._serialized_end=618
  _CUSTOMTASK._serialized_start=621
  _CUSTOMTASK._serialized_end=918
  _TASK._serialized_start=921
  _TASK._serialized_end=1311
  _TASK_METADATA._serialized_start=1270
  _TASK_METADATA._serialized_end=1292
  _PIPELINECONFIG._serialized_start=1314
  _PIPELINECONFIG._serialized_end=1532
  _PIPELINETEMPLATE._serialized_start=1535
  _PIPELINETEMPLATE._serialized_end=1757
# @@protoc_insertion_point(module_scope)
