# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/template/service.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from prodvana.proto.prodvana.object import meta_pb2 as prodvana_dot_object_dot_meta__pb2
from prodvana.proto.prodvana.service import service_config_pb2 as prodvana_dot_service_dot_service__config__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from prodvana.proto.validate import validate_pb2 as validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1fprodvana/template/service.proto\x12\x11prodvana.template\x1a\x1aprodvana/object/meta.proto\x1a%prodvana/service/service_config.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x17validate/validate.proto\"\xba\x01\n\x0fServiceTemplate\x12)\n\x04meta\x18\x01 \x01(\x0b\x32\x1b.prodvana.object.ObjectMeta\x12\x39\n\x15last_update_timestamp\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x41\n\x0eservice_config\x18\x03 \x01(\x0b\x32\x1f.prodvana.service.ServiceConfigB\x08\xfa\x42\x05\x8a\x01\x02\x08\x01\x42MZKgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/templateb\x06proto3')



_SERVICETEMPLATE = DESCRIPTOR.message_types_by_name['ServiceTemplate']
ServiceTemplate = _reflection.GeneratedProtocolMessageType('ServiceTemplate', (_message.Message,), {
  'DESCRIPTOR' : _SERVICETEMPLATE,
  '__module__' : 'prodvana.template.service_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.template.ServiceTemplate)
  })
_sym_db.RegisterMessage(ServiceTemplate)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZKgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/template'
  _SERVICETEMPLATE.fields_by_name['service_config']._options = None
  _SERVICETEMPLATE.fields_by_name['service_config']._serialized_options = b'\372B\005\212\001\002\010\001'
  _SERVICETEMPLATE._serialized_start=180
  _SERVICETEMPLATE._serialized_end=366
# @@protoc_insertion_point(module_scope)
