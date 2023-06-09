# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/protection/attachments.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from prodvana.proto.prodvana.protection import protection_reference_pb2 as prodvana_dot_protection_dot_protection__reference__pb2
from prodvana.proto.validate import validate_pb2 as validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n%prodvana/protection/attachments.proto\x12\x13prodvana.protection\x1a.prodvana/protection/protection_reference.proto\x1a\x17validate/validate.proto\"\xec\x01\n\x1fProtectionConvergenceAttachment\x12:\n\x04name\x18\x01 \x01(\tB,\xfa\x42)r\'\x10\x00\x18?2!^[a-z]?([a-z0-9-]*[a-z0-9]){0,1}$\x12?\n\x03ref\x18\x02 \x01(\x0b\x32(.prodvana.protection.ProtectionReferenceB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12L\n\tlifecycle\x18\x03 \x03(\x0b\x32(.prodvana.protection.ProtectionLifecycleB\x0f\xfa\x42\x0c\x92\x01\t\x08\x01\"\x05\x8a\x01\x02\x10\x01*Y\n\x0e\x41ttachmentType\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x13\n\x0fRELEASE_CHANNEL\x10\x01\x12\x14\n\x10SERVICE_INSTANCE\x10\x02\x12\x0f\n\x0b\x43ONVERGENCE\x10\x03\x42OZMgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/protectionb\x06proto3')

_ATTACHMENTTYPE = DESCRIPTOR.enum_types_by_name['AttachmentType']
AttachmentType = enum_type_wrapper.EnumTypeWrapper(_ATTACHMENTTYPE)
UNKNOWN = 0
RELEASE_CHANNEL = 1
SERVICE_INSTANCE = 2
CONVERGENCE = 3


_PROTECTIONCONVERGENCEATTACHMENT = DESCRIPTOR.message_types_by_name['ProtectionConvergenceAttachment']
ProtectionConvergenceAttachment = _reflection.GeneratedProtocolMessageType('ProtectionConvergenceAttachment', (_message.Message,), {
  'DESCRIPTOR' : _PROTECTIONCONVERGENCEATTACHMENT,
  '__module__' : 'prodvana.protection.attachments_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.protection.ProtectionConvergenceAttachment)
  })
_sym_db.RegisterMessage(ProtectionConvergenceAttachment)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZMgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/protection'
  _PROTECTIONCONVERGENCEATTACHMENT.fields_by_name['name']._options = None
  _PROTECTIONCONVERGENCEATTACHMENT.fields_by_name['name']._serialized_options = b'\372B)r\'\020\000\030?2!^[a-z]?([a-z0-9-]*[a-z0-9]){0,1}$'
  _PROTECTIONCONVERGENCEATTACHMENT.fields_by_name['ref']._options = None
  _PROTECTIONCONVERGENCEATTACHMENT.fields_by_name['ref']._serialized_options = b'\372B\005\212\001\002\020\001'
  _PROTECTIONCONVERGENCEATTACHMENT.fields_by_name['lifecycle']._options = None
  _PROTECTIONCONVERGENCEATTACHMENT.fields_by_name['lifecycle']._serialized_options = b'\372B\014\222\001\t\010\001\"\005\212\001\002\020\001'
  _ATTACHMENTTYPE._serialized_start=374
  _ATTACHMENTTYPE._serialized_end=463
  _PROTECTIONCONVERGENCEATTACHMENT._serialized_start=136
  _PROTECTIONCONVERGENCEATTACHMENT._serialized_end=372
# @@protoc_insertion_point(module_scope)
