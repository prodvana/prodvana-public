# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/delivery_module/manager.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from prodvana.proto.prodvana.delivery_module import object_pb2 as prodvana_dot_delivery__module_dot_object__pb2
from prodvana.proto.prodvana.delivery_module import config_pb2 as prodvana_dot_delivery__module_dot_config__pb2
from prodvana.proto.prodvana.version import source_metadata_pb2 as prodvana_dot_version_dot_source__metadata__pb2
from prodvana.proto.validate import validate_pb2 as validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n&prodvana/delivery_module/manager.proto\x12\x18prodvana.delivery_module\x1a\x1cgoogle/api/annotations.proto\x1a%prodvana/delivery_module/object.proto\x1a%prodvana/delivery_module/config.proto\x1a&prodvana/version/source_metadata.proto\x1a\x17validate/validate.proto\"\xdb\x01\n\x1a\x43onfigureDeliveryModuleReq\x12X\n\x16\x64\x65livery_module_config\x18\x01 \x01(\x0b\x32..prodvana.delivery_module.DeliveryModuleConfigB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12(\n\x06source\x18\x02 \x01(\x0e\x32\x18.prodvana.version.Source\x12\x39\n\x0fsource_metadata\x18\x03 \x01(\x0b\x32 .prodvana.version.SourceMetadata\"J\n\x1b\x43onfigureDeliveryModuleResp\x12\x1a\n\x12\x64\x65livery_module_id\x18\x01 \x01(\t\x12\x0f\n\x07version\x18\x02 \x01(\t\"%\n#ValidateConfigureDeliveryModuleResp\"?\n\x16ListDeliveryModulesReq\x12\x12\n\npage_token\x18\x01 \x01(\t\x12\x11\n\tpage_size\x18\x02 \x01(\x05\"v\n\x17ListDeliveryModulesResp\x12\x42\n\x10\x64\x65livery_modules\x18\x01 \x03(\x0b\x32(.prodvana.delivery_module.DeliveryModule\x12\x17\n\x0fnext_page_token\x18\x02 \x01(\t\"8\n\x14GetDeliveryModuleReq\x12 \n\x0f\x64\x65livery_module\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\"Z\n\x15GetDeliveryModuleResp\x12\x41\n\x0f\x64\x65livery_module\x18\x01 \x01(\x0b\x32(.prodvana.delivery_module.DeliveryModule\"O\n\x1aGetDeliveryModuleConfigReq\x12 \n\x0f\x64\x65livery_module\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x0f\n\x07version\x18\x02 \x01(\t\"n\n\x1bGetDeliveryModuleConfigResp\x12>\n\x06\x63onfig\x18\x01 \x01(\x0b\x32..prodvana.delivery_module.DeliveryModuleConfig\x12\x0f\n\x07version\x18\x02 \x01(\t2\x9c\x07\n\x15\x44\x65liveryModuleManager\x12\xb1\x01\n\x17\x43onfigureDeliveryModule\x12\x34.prodvana.delivery_module.ConfigureDeliveryModuleReq\x1a\x35.prodvana.delivery_module.ConfigureDeliveryModuleResp\")\x82\xd3\xe4\x93\x02#\"\x1e/v1/delivery_modules/configure:\x01*\x12\xca\x01\n\x1fValidateConfigureDeliveryModule\x12\x34.prodvana.delivery_module.ConfigureDeliveryModuleReq\x1a=.prodvana.delivery_module.ValidateConfigureDeliveryModuleResp\"2\x82\xd3\xe4\x93\x02,\"\'/v1/delivery_modules/configure/validate:\x01*\x12\x98\x01\n\x13ListDeliveryModules\x12\x30.prodvana.delivery_module.ListDeliveryModulesReq\x1a\x31.prodvana.delivery_module.ListDeliveryModulesResp\"\x1c\x82\xd3\xe4\x93\x02\x16\x12\x14/v1/delivery_modules\x12\xa5\x01\n\x11GetDeliveryModule\x12..prodvana.delivery_module.GetDeliveryModuleReq\x1a/.prodvana.delivery_module.GetDeliveryModuleResp\"/\x82\xd3\xe4\x93\x02)\x12\'/v1/delivery_module/{delivery_module=*}\x12\xbe\x01\n\x17GetDeliveryModuleConfig\x12\x34.prodvana.delivery_module.GetDeliveryModuleConfigReq\x1a\x35.prodvana.delivery_module.GetDeliveryModuleConfigResp\"6\x82\xd3\xe4\x93\x02\x30\x12./v1/delivery_module/{delivery_module=*}/configBTZRgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/delivery_moduleb\x06proto3')



_CONFIGUREDELIVERYMODULEREQ = DESCRIPTOR.message_types_by_name['ConfigureDeliveryModuleReq']
_CONFIGUREDELIVERYMODULERESP = DESCRIPTOR.message_types_by_name['ConfigureDeliveryModuleResp']
_VALIDATECONFIGUREDELIVERYMODULERESP = DESCRIPTOR.message_types_by_name['ValidateConfigureDeliveryModuleResp']
_LISTDELIVERYMODULESREQ = DESCRIPTOR.message_types_by_name['ListDeliveryModulesReq']
_LISTDELIVERYMODULESRESP = DESCRIPTOR.message_types_by_name['ListDeliveryModulesResp']
_GETDELIVERYMODULEREQ = DESCRIPTOR.message_types_by_name['GetDeliveryModuleReq']
_GETDELIVERYMODULERESP = DESCRIPTOR.message_types_by_name['GetDeliveryModuleResp']
_GETDELIVERYMODULECONFIGREQ = DESCRIPTOR.message_types_by_name['GetDeliveryModuleConfigReq']
_GETDELIVERYMODULECONFIGRESP = DESCRIPTOR.message_types_by_name['GetDeliveryModuleConfigResp']
ConfigureDeliveryModuleReq = _reflection.GeneratedProtocolMessageType('ConfigureDeliveryModuleReq', (_message.Message,), {
  'DESCRIPTOR' : _CONFIGUREDELIVERYMODULEREQ,
  '__module__' : 'prodvana.delivery_module.manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.delivery_module.ConfigureDeliveryModuleReq)
  })
_sym_db.RegisterMessage(ConfigureDeliveryModuleReq)

ConfigureDeliveryModuleResp = _reflection.GeneratedProtocolMessageType('ConfigureDeliveryModuleResp', (_message.Message,), {
  'DESCRIPTOR' : _CONFIGUREDELIVERYMODULERESP,
  '__module__' : 'prodvana.delivery_module.manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.delivery_module.ConfigureDeliveryModuleResp)
  })
_sym_db.RegisterMessage(ConfigureDeliveryModuleResp)

ValidateConfigureDeliveryModuleResp = _reflection.GeneratedProtocolMessageType('ValidateConfigureDeliveryModuleResp', (_message.Message,), {
  'DESCRIPTOR' : _VALIDATECONFIGUREDELIVERYMODULERESP,
  '__module__' : 'prodvana.delivery_module.manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.delivery_module.ValidateConfigureDeliveryModuleResp)
  })
_sym_db.RegisterMessage(ValidateConfigureDeliveryModuleResp)

ListDeliveryModulesReq = _reflection.GeneratedProtocolMessageType('ListDeliveryModulesReq', (_message.Message,), {
  'DESCRIPTOR' : _LISTDELIVERYMODULESREQ,
  '__module__' : 'prodvana.delivery_module.manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.delivery_module.ListDeliveryModulesReq)
  })
_sym_db.RegisterMessage(ListDeliveryModulesReq)

ListDeliveryModulesResp = _reflection.GeneratedProtocolMessageType('ListDeliveryModulesResp', (_message.Message,), {
  'DESCRIPTOR' : _LISTDELIVERYMODULESRESP,
  '__module__' : 'prodvana.delivery_module.manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.delivery_module.ListDeliveryModulesResp)
  })
_sym_db.RegisterMessage(ListDeliveryModulesResp)

GetDeliveryModuleReq = _reflection.GeneratedProtocolMessageType('GetDeliveryModuleReq', (_message.Message,), {
  'DESCRIPTOR' : _GETDELIVERYMODULEREQ,
  '__module__' : 'prodvana.delivery_module.manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.delivery_module.GetDeliveryModuleReq)
  })
_sym_db.RegisterMessage(GetDeliveryModuleReq)

GetDeliveryModuleResp = _reflection.GeneratedProtocolMessageType('GetDeliveryModuleResp', (_message.Message,), {
  'DESCRIPTOR' : _GETDELIVERYMODULERESP,
  '__module__' : 'prodvana.delivery_module.manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.delivery_module.GetDeliveryModuleResp)
  })
_sym_db.RegisterMessage(GetDeliveryModuleResp)

GetDeliveryModuleConfigReq = _reflection.GeneratedProtocolMessageType('GetDeliveryModuleConfigReq', (_message.Message,), {
  'DESCRIPTOR' : _GETDELIVERYMODULECONFIGREQ,
  '__module__' : 'prodvana.delivery_module.manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.delivery_module.GetDeliveryModuleConfigReq)
  })
_sym_db.RegisterMessage(GetDeliveryModuleConfigReq)

GetDeliveryModuleConfigResp = _reflection.GeneratedProtocolMessageType('GetDeliveryModuleConfigResp', (_message.Message,), {
  'DESCRIPTOR' : _GETDELIVERYMODULECONFIGRESP,
  '__module__' : 'prodvana.delivery_module.manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.delivery_module.GetDeliveryModuleConfigResp)
  })
_sym_db.RegisterMessage(GetDeliveryModuleConfigResp)

_DELIVERYMODULEMANAGER = DESCRIPTOR.services_by_name['DeliveryModuleManager']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZRgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/delivery_module'
  _CONFIGUREDELIVERYMODULEREQ.fields_by_name['delivery_module_config']._options = None
  _CONFIGUREDELIVERYMODULEREQ.fields_by_name['delivery_module_config']._serialized_options = b'\372B\005\212\001\002\020\001'
  _GETDELIVERYMODULEREQ.fields_by_name['delivery_module']._options = None
  _GETDELIVERYMODULEREQ.fields_by_name['delivery_module']._serialized_options = b'\372B\004r\002\020\001'
  _GETDELIVERYMODULECONFIGREQ.fields_by_name['delivery_module']._options = None
  _GETDELIVERYMODULECONFIGREQ.fields_by_name['delivery_module']._serialized_options = b'\372B\004r\002\020\001'
  _DELIVERYMODULEMANAGER.methods_by_name['ConfigureDeliveryModule']._options = None
  _DELIVERYMODULEMANAGER.methods_by_name['ConfigureDeliveryModule']._serialized_options = b'\202\323\344\223\002#\"\036/v1/delivery_modules/configure:\001*'
  _DELIVERYMODULEMANAGER.methods_by_name['ValidateConfigureDeliveryModule']._options = None
  _DELIVERYMODULEMANAGER.methods_by_name['ValidateConfigureDeliveryModule']._serialized_options = b'\202\323\344\223\002,\"\'/v1/delivery_modules/configure/validate:\001*'
  _DELIVERYMODULEMANAGER.methods_by_name['ListDeliveryModules']._options = None
  _DELIVERYMODULEMANAGER.methods_by_name['ListDeliveryModules']._serialized_options = b'\202\323\344\223\002\026\022\024/v1/delivery_modules'
  _DELIVERYMODULEMANAGER.methods_by_name['GetDeliveryModule']._options = None
  _DELIVERYMODULEMANAGER.methods_by_name['GetDeliveryModule']._serialized_options = b'\202\323\344\223\002)\022\'/v1/delivery_module/{delivery_module=*}'
  _DELIVERYMODULEMANAGER.methods_by_name['GetDeliveryModuleConfig']._options = None
  _DELIVERYMODULEMANAGER.methods_by_name['GetDeliveryModuleConfig']._serialized_options = b'\202\323\344\223\0020\022./v1/delivery_module/{delivery_module=*}/config'
  _CONFIGUREDELIVERYMODULEREQ._serialized_start=242
  _CONFIGUREDELIVERYMODULEREQ._serialized_end=461
  _CONFIGUREDELIVERYMODULERESP._serialized_start=463
  _CONFIGUREDELIVERYMODULERESP._serialized_end=537
  _VALIDATECONFIGUREDELIVERYMODULERESP._serialized_start=539
  _VALIDATECONFIGUREDELIVERYMODULERESP._serialized_end=576
  _LISTDELIVERYMODULESREQ._serialized_start=578
  _LISTDELIVERYMODULESREQ._serialized_end=641
  _LISTDELIVERYMODULESRESP._serialized_start=643
  _LISTDELIVERYMODULESRESP._serialized_end=761
  _GETDELIVERYMODULEREQ._serialized_start=763
  _GETDELIVERYMODULEREQ._serialized_end=819
  _GETDELIVERYMODULERESP._serialized_start=821
  _GETDELIVERYMODULERESP._serialized_end=911
  _GETDELIVERYMODULECONFIGREQ._serialized_start=913
  _GETDELIVERYMODULECONFIGREQ._serialized_end=992
  _GETDELIVERYMODULECONFIGRESP._serialized_start=994
  _GETDELIVERYMODULECONFIGRESP._serialized_end=1104
  _DELIVERYMODULEMANAGER._serialized_start=1107
  _DELIVERYMODULEMANAGER._serialized_end=2031
# @@protoc_insertion_point(module_scope)
