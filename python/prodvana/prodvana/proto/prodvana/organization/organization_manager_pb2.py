# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/organization/organization_manager.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from prodvana.proto.prodvana.config_writeback import writeback_pb2 as prodvana_dot_config__writeback_dot_writeback__pb2
from prodvana.proto.prodvana.insights import insights_pb2 as prodvana_dot_insights_dot_insights__pb2
from prodvana.proto.prodvana.metrics import metrics_pb2 as prodvana_dot_metrics_dot_metrics__pb2
from prodvana.proto.prodvana.organization import user_metadata_pb2 as prodvana_dot_organization_dot_user__metadata__pb2
from prodvana.proto.prodvana.users import users_pb2 as prodvana_dot_users_dot_users__pb2
from prodvana.proto.validate import validate_pb2 as validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n0prodvana/organization/organization_manager.proto\x12\x15prodvana.organization\x1a\x1cgoogle/api/annotations.proto\x1a\x1egoogle/protobuf/duration.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a)prodvana/config_writeback/writeback.proto\x1a prodvana/insights/insights.proto\x1a\x1eprodvana/metrics/metrics.proto\x1a)prodvana/organization/user_metadata.proto\x1a\x1aprodvana/users/users.proto\x1a\x17validate/validate.proto\"\xd4\x01\n\x10OrganizationInfo\x12\n\n\x02id\x18\x01 \x01(\t\x12\x14\n\x0c\x64isplay_name\x18\x02 \x01(\t\x12H\n\x10writeback_config\x18\x03 \x01(\x0b\x32..prodvana.config_writeback.ConfigWritebackPath\x12\x0c\n\x04slug\x18\x04 \x01(\t\x12\x46\n\ruser_metadata\x18\x05 \x01(\x0b\x32/.prodvana.organization.OrganizationUserMetadata\"\x14\n\x12GetOrganizationReq\"T\n\x13GetOrganizationResp\x12=\n\x0corganization\x18\x01 \x01(\x0b\x32\'.prodvana.organization.OrganizationInfo\"\x83\x01\n\x19GetOrganizationMetricsReq\x12\x33\n\x0fstart_timestamp\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x31\n\rend_timestamp\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"]\n\x1aGetOrganizationMetricsResp\x12?\n\x12\x64\x65ployment_metrics\x18\x01 \x01(\x0b\x32#.prodvana.metrics.DeploymentMetrics\"\x1c\n\x1aGetOrganizationInsightsReq\"K\n\x1bGetOrganizationInsightsResp\x12,\n\x08insights\x18\x01 \x03(\x0b\x32\x1a.prodvana.insights.Insight\"\x8a\x01\n\x1cSnoozeOrganizationInsightReq\x12\x31\n\x05\x63lass\x18\x01 \x01(\x0e\x32\x18.prodvana.insights.ClassB\x08\xfa\x42\x05\x82\x01\x02\x10\x01\x12\x37\n\x08\x64uration\x18\x02 \x01(\x0b\x32\x19.google.protobuf.DurationB\n\xfa\x42\x07\xaa\x01\x04\x08\x01*\x00\"\x1f\n\x1dSnoozeOrganizationInsightResp\"\x1c\n\x1aGetOrganizationMetadataReq\"`\n\x1bGetOrganizationMetadataResp\x12\x41\n\x08metadata\x18\x01 \x01(\x0b\x32/.prodvana.organization.OrganizationUserMetadata\"i\n\x1aSetOrganizationMetadataReq\x12K\n\x08metadata\x18\x01 \x01(\x0b\x32/.prodvana.organization.OrganizationUserMetadataB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\"\x1d\n\x1bSetOrganizationMetadataResp\"&\n\nGetUserReq\x12\x18\n\x07user_id\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\"1\n\x0bGetUserResp\x12\"\n\x04user\x18\x01 \x01(\x0b\x32\x14.prodvana.users.User2\xe1\x08\n\x13OrganizationManager\x12\x82\x01\n\x0fGetOrganization\x12).prodvana.organization.GetOrganizationReq\x1a*.prodvana.organization.GetOrganizationResp\"\x18\x82\xd3\xe4\x93\x02\x12\x12\x10/v1/organization\x12\x9f\x01\n\x16GetOrganizationMetrics\x12\x30.prodvana.organization.GetOrganizationMetricsReq\x1a\x31.prodvana.organization.GetOrganizationMetricsResp\" \x82\xd3\xe4\x93\x02\x1a\x12\x18/v1/organization/metrics\x12\xa3\x01\n\x17GetOrganizationInsights\x12\x31.prodvana.organization.GetOrganizationInsightsReq\x1a\x32.prodvana.organization.GetOrganizationInsightsResp\"!\x82\xd3\xe4\x93\x02\x1b\x12\x19/v1/organization/insights\x12\xb0\x01\n\x19SnoozeOrganizationInsight\x12\x33.prodvana.organization.SnoozeOrganizationInsightReq\x1a\x34.prodvana.organization.SnoozeOrganizationInsightResp\"(\x82\xd3\xe4\x93\x02\"\x1a /v1/organization/insights/snooze\x12\xa3\x01\n\x17GetOrganizationMetadata\x12\x31.prodvana.organization.GetOrganizationMetadataReq\x1a\x32.prodvana.organization.GetOrganizationMetadataResp\"!\x82\xd3\xe4\x93\x02\x1b\x12\x19/v1/organization/metadata\x12\xa6\x01\n\x17SetOrganizationMetadata\x12\x31.prodvana.organization.SetOrganizationMetadataReq\x1a\x32.prodvana.organization.SetOrganizationMetadataResp\"$\x82\xd3\xe4\x93\x02\x1e\"\x19/v1/organization/metadata:\x01*\x12{\n\x07GetUser\x12!.prodvana.organization.GetUserReq\x1a\".prodvana.organization.GetUserResp\")\x82\xd3\xe4\x93\x02#\x12!/v1/organization/user/{user_id=*}BQZOgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/organizationb\x06proto3')



_ORGANIZATIONINFO = DESCRIPTOR.message_types_by_name['OrganizationInfo']
_GETORGANIZATIONREQ = DESCRIPTOR.message_types_by_name['GetOrganizationReq']
_GETORGANIZATIONRESP = DESCRIPTOR.message_types_by_name['GetOrganizationResp']
_GETORGANIZATIONMETRICSREQ = DESCRIPTOR.message_types_by_name['GetOrganizationMetricsReq']
_GETORGANIZATIONMETRICSRESP = DESCRIPTOR.message_types_by_name['GetOrganizationMetricsResp']
_GETORGANIZATIONINSIGHTSREQ = DESCRIPTOR.message_types_by_name['GetOrganizationInsightsReq']
_GETORGANIZATIONINSIGHTSRESP = DESCRIPTOR.message_types_by_name['GetOrganizationInsightsResp']
_SNOOZEORGANIZATIONINSIGHTREQ = DESCRIPTOR.message_types_by_name['SnoozeOrganizationInsightReq']
_SNOOZEORGANIZATIONINSIGHTRESP = DESCRIPTOR.message_types_by_name['SnoozeOrganizationInsightResp']
_GETORGANIZATIONMETADATAREQ = DESCRIPTOR.message_types_by_name['GetOrganizationMetadataReq']
_GETORGANIZATIONMETADATARESP = DESCRIPTOR.message_types_by_name['GetOrganizationMetadataResp']
_SETORGANIZATIONMETADATAREQ = DESCRIPTOR.message_types_by_name['SetOrganizationMetadataReq']
_SETORGANIZATIONMETADATARESP = DESCRIPTOR.message_types_by_name['SetOrganizationMetadataResp']
_GETUSERREQ = DESCRIPTOR.message_types_by_name['GetUserReq']
_GETUSERRESP = DESCRIPTOR.message_types_by_name['GetUserResp']
OrganizationInfo = _reflection.GeneratedProtocolMessageType('OrganizationInfo', (_message.Message,), {
  'DESCRIPTOR' : _ORGANIZATIONINFO,
  '__module__' : 'prodvana.organization.organization_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.organization.OrganizationInfo)
  })
_sym_db.RegisterMessage(OrganizationInfo)

GetOrganizationReq = _reflection.GeneratedProtocolMessageType('GetOrganizationReq', (_message.Message,), {
  'DESCRIPTOR' : _GETORGANIZATIONREQ,
  '__module__' : 'prodvana.organization.organization_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.organization.GetOrganizationReq)
  })
_sym_db.RegisterMessage(GetOrganizationReq)

GetOrganizationResp = _reflection.GeneratedProtocolMessageType('GetOrganizationResp', (_message.Message,), {
  'DESCRIPTOR' : _GETORGANIZATIONRESP,
  '__module__' : 'prodvana.organization.organization_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.organization.GetOrganizationResp)
  })
_sym_db.RegisterMessage(GetOrganizationResp)

GetOrganizationMetricsReq = _reflection.GeneratedProtocolMessageType('GetOrganizationMetricsReq', (_message.Message,), {
  'DESCRIPTOR' : _GETORGANIZATIONMETRICSREQ,
  '__module__' : 'prodvana.organization.organization_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.organization.GetOrganizationMetricsReq)
  })
_sym_db.RegisterMessage(GetOrganizationMetricsReq)

GetOrganizationMetricsResp = _reflection.GeneratedProtocolMessageType('GetOrganizationMetricsResp', (_message.Message,), {
  'DESCRIPTOR' : _GETORGANIZATIONMETRICSRESP,
  '__module__' : 'prodvana.organization.organization_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.organization.GetOrganizationMetricsResp)
  })
_sym_db.RegisterMessage(GetOrganizationMetricsResp)

GetOrganizationInsightsReq = _reflection.GeneratedProtocolMessageType('GetOrganizationInsightsReq', (_message.Message,), {
  'DESCRIPTOR' : _GETORGANIZATIONINSIGHTSREQ,
  '__module__' : 'prodvana.organization.organization_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.organization.GetOrganizationInsightsReq)
  })
_sym_db.RegisterMessage(GetOrganizationInsightsReq)

GetOrganizationInsightsResp = _reflection.GeneratedProtocolMessageType('GetOrganizationInsightsResp', (_message.Message,), {
  'DESCRIPTOR' : _GETORGANIZATIONINSIGHTSRESP,
  '__module__' : 'prodvana.organization.organization_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.organization.GetOrganizationInsightsResp)
  })
_sym_db.RegisterMessage(GetOrganizationInsightsResp)

SnoozeOrganizationInsightReq = _reflection.GeneratedProtocolMessageType('SnoozeOrganizationInsightReq', (_message.Message,), {
  'DESCRIPTOR' : _SNOOZEORGANIZATIONINSIGHTREQ,
  '__module__' : 'prodvana.organization.organization_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.organization.SnoozeOrganizationInsightReq)
  })
_sym_db.RegisterMessage(SnoozeOrganizationInsightReq)

SnoozeOrganizationInsightResp = _reflection.GeneratedProtocolMessageType('SnoozeOrganizationInsightResp', (_message.Message,), {
  'DESCRIPTOR' : _SNOOZEORGANIZATIONINSIGHTRESP,
  '__module__' : 'prodvana.organization.organization_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.organization.SnoozeOrganizationInsightResp)
  })
_sym_db.RegisterMessage(SnoozeOrganizationInsightResp)

GetOrganizationMetadataReq = _reflection.GeneratedProtocolMessageType('GetOrganizationMetadataReq', (_message.Message,), {
  'DESCRIPTOR' : _GETORGANIZATIONMETADATAREQ,
  '__module__' : 'prodvana.organization.organization_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.organization.GetOrganizationMetadataReq)
  })
_sym_db.RegisterMessage(GetOrganizationMetadataReq)

GetOrganizationMetadataResp = _reflection.GeneratedProtocolMessageType('GetOrganizationMetadataResp', (_message.Message,), {
  'DESCRIPTOR' : _GETORGANIZATIONMETADATARESP,
  '__module__' : 'prodvana.organization.organization_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.organization.GetOrganizationMetadataResp)
  })
_sym_db.RegisterMessage(GetOrganizationMetadataResp)

SetOrganizationMetadataReq = _reflection.GeneratedProtocolMessageType('SetOrganizationMetadataReq', (_message.Message,), {
  'DESCRIPTOR' : _SETORGANIZATIONMETADATAREQ,
  '__module__' : 'prodvana.organization.organization_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.organization.SetOrganizationMetadataReq)
  })
_sym_db.RegisterMessage(SetOrganizationMetadataReq)

SetOrganizationMetadataResp = _reflection.GeneratedProtocolMessageType('SetOrganizationMetadataResp', (_message.Message,), {
  'DESCRIPTOR' : _SETORGANIZATIONMETADATARESP,
  '__module__' : 'prodvana.organization.organization_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.organization.SetOrganizationMetadataResp)
  })
_sym_db.RegisterMessage(SetOrganizationMetadataResp)

GetUserReq = _reflection.GeneratedProtocolMessageType('GetUserReq', (_message.Message,), {
  'DESCRIPTOR' : _GETUSERREQ,
  '__module__' : 'prodvana.organization.organization_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.organization.GetUserReq)
  })
_sym_db.RegisterMessage(GetUserReq)

GetUserResp = _reflection.GeneratedProtocolMessageType('GetUserResp', (_message.Message,), {
  'DESCRIPTOR' : _GETUSERRESP,
  '__module__' : 'prodvana.organization.organization_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.organization.GetUserResp)
  })
_sym_db.RegisterMessage(GetUserResp)

_ORGANIZATIONMANAGER = DESCRIPTOR.services_by_name['OrganizationManager']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZOgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/organization'
  _SNOOZEORGANIZATIONINSIGHTREQ.fields_by_name['class']._options = None
  _SNOOZEORGANIZATIONINSIGHTREQ.fields_by_name['class']._serialized_options = b'\372B\005\202\001\002\020\001'
  _SNOOZEORGANIZATIONINSIGHTREQ.fields_by_name['duration']._options = None
  _SNOOZEORGANIZATIONINSIGHTREQ.fields_by_name['duration']._serialized_options = b'\372B\007\252\001\004\010\001*\000'
  _SETORGANIZATIONMETADATAREQ.fields_by_name['metadata']._options = None
  _SETORGANIZATIONMETADATAREQ.fields_by_name['metadata']._serialized_options = b'\372B\005\212\001\002\020\001'
  _GETUSERREQ.fields_by_name['user_id']._options = None
  _GETUSERREQ.fields_by_name['user_id']._serialized_options = b'\372B\004r\002\020\001'
  _ORGANIZATIONMANAGER.methods_by_name['GetOrganization']._options = None
  _ORGANIZATIONMANAGER.methods_by_name['GetOrganization']._serialized_options = b'\202\323\344\223\002\022\022\020/v1/organization'
  _ORGANIZATIONMANAGER.methods_by_name['GetOrganizationMetrics']._options = None
  _ORGANIZATIONMANAGER.methods_by_name['GetOrganizationMetrics']._serialized_options = b'\202\323\344\223\002\032\022\030/v1/organization/metrics'
  _ORGANIZATIONMANAGER.methods_by_name['GetOrganizationInsights']._options = None
  _ORGANIZATIONMANAGER.methods_by_name['GetOrganizationInsights']._serialized_options = b'\202\323\344\223\002\033\022\031/v1/organization/insights'
  _ORGANIZATIONMANAGER.methods_by_name['SnoozeOrganizationInsight']._options = None
  _ORGANIZATIONMANAGER.methods_by_name['SnoozeOrganizationInsight']._serialized_options = b'\202\323\344\223\002\"\032 /v1/organization/insights/snooze'
  _ORGANIZATIONMANAGER.methods_by_name['GetOrganizationMetadata']._options = None
  _ORGANIZATIONMANAGER.methods_by_name['GetOrganizationMetadata']._serialized_options = b'\202\323\344\223\002\033\022\031/v1/organization/metadata'
  _ORGANIZATIONMANAGER.methods_by_name['SetOrganizationMetadata']._options = None
  _ORGANIZATIONMANAGER.methods_by_name['SetOrganizationMetadata']._serialized_options = b'\202\323\344\223\002\036\"\031/v1/organization/metadata:\001*'
  _ORGANIZATIONMANAGER.methods_by_name['GetUser']._options = None
  _ORGANIZATIONMANAGER.methods_by_name['GetUser']._serialized_options = b'\202\323\344\223\002#\022!/v1/organization/user/{user_id=*}'
  _ORGANIZATIONINFO._serialized_start=376
  _ORGANIZATIONINFO._serialized_end=588
  _GETORGANIZATIONREQ._serialized_start=590
  _GETORGANIZATIONREQ._serialized_end=610
  _GETORGANIZATIONRESP._serialized_start=612
  _GETORGANIZATIONRESP._serialized_end=696
  _GETORGANIZATIONMETRICSREQ._serialized_start=699
  _GETORGANIZATIONMETRICSREQ._serialized_end=830
  _GETORGANIZATIONMETRICSRESP._serialized_start=832
  _GETORGANIZATIONMETRICSRESP._serialized_end=925
  _GETORGANIZATIONINSIGHTSREQ._serialized_start=927
  _GETORGANIZATIONINSIGHTSREQ._serialized_end=955
  _GETORGANIZATIONINSIGHTSRESP._serialized_start=957
  _GETORGANIZATIONINSIGHTSRESP._serialized_end=1032
  _SNOOZEORGANIZATIONINSIGHTREQ._serialized_start=1035
  _SNOOZEORGANIZATIONINSIGHTREQ._serialized_end=1173
  _SNOOZEORGANIZATIONINSIGHTRESP._serialized_start=1175
  _SNOOZEORGANIZATIONINSIGHTRESP._serialized_end=1206
  _GETORGANIZATIONMETADATAREQ._serialized_start=1208
  _GETORGANIZATIONMETADATAREQ._serialized_end=1236
  _GETORGANIZATIONMETADATARESP._serialized_start=1238
  _GETORGANIZATIONMETADATARESP._serialized_end=1334
  _SETORGANIZATIONMETADATAREQ._serialized_start=1336
  _SETORGANIZATIONMETADATAREQ._serialized_end=1441
  _SETORGANIZATIONMETADATARESP._serialized_start=1443
  _SETORGANIZATIONMETADATARESP._serialized_end=1472
  _GETUSERREQ._serialized_start=1474
  _GETUSERREQ._serialized_end=1512
  _GETUSERRESP._serialized_start=1514
  _GETUSERRESP._serialized_end=1563
  _ORGANIZATIONMANAGER._serialized_start=1566
  _ORGANIZATIONMANAGER._serialized_end=2687
# @@protoc_insertion_point(module_scope)
