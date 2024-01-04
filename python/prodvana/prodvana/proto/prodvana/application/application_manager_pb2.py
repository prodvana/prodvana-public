# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/application/application_manager.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from prodvana.proto.prodvana.application import application_config_pb2 as prodvana_dot_application_dot_application__config__pb2
from prodvana.proto.prodvana.application import object_pb2 as prodvana_dot_application_dot_object__pb2
from prodvana.proto.prodvana.application import user_metadata_pb2 as prodvana_dot_application_dot_user__metadata__pb2
from prodvana.proto.prodvana.common_config import dangerous_action_pb2 as prodvana_dot_common__config_dot_dangerous__action__pb2
from prodvana.proto.prodvana.metrics import metrics_pb2 as prodvana_dot_metrics_dot_metrics__pb2
from prodvana.proto.prodvana.insights import insights_pb2 as prodvana_dot_insights_dot_insights__pb2
from prodvana.proto.prodvana.object import meta_pb2 as prodvana_dot_object_dot_meta__pb2
from prodvana.proto.prodvana.version import source_metadata_pb2 as prodvana_dot_version_dot_source__metadata__pb2
from prodvana.proto.validate import validate_pb2 as validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n.prodvana/application/application_manager.proto\x12\x14prodvana.application\x1a\x1cgoogle/api/annotations.proto\x1a\x1egoogle/protobuf/duration.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a-prodvana/application/application_config.proto\x1a!prodvana/application/object.proto\x1a(prodvana/application/user_metadata.proto\x1a-prodvana/common_config/dangerous_action.proto\x1a\x1eprodvana/metrics/metrics.proto\x1a prodvana/insights/insights.proto\x1a\x1aprodvana/object/meta.proto\x1a&prodvana/version/source_metadata.proto\x1a\x17validate/validate.proto\"\'\n\x13ListApplicationsReq\x12\x10\n\x08\x64\x65tailed\x18\x01 \x01(\x08\"O\n\x14ListApplicationsResp\x12\x37\n\x0c\x61pplications\x18\x01 \x03(\x0b\x32!.prodvana.application.Application\"1\n\x11GetApplicationReq\x12\x1c\n\x0b\x61pplication\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\"L\n\x12GetApplicationResp\x12\x36\n\x0b\x61pplication\x18\x01 \x01(\x0b\x32!.prodvana.application.Application\"\x8a\x02\n\x17\x43onfigureApplicationReq\x12M\n\x12\x61pplication_config\x18\x01 \x01(\x0b\x32\'.prodvana.application.ApplicationConfigB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12%\n\x1d\x61pproved_dangerous_action_ids\x18\x02 \x03(\t\x12(\n\x06source\x18\x03 \x01(\x0e\x32\x18.prodvana.version.Source\x12\x39\n\x0fsource_metadata\x18\x04 \x01(\x0b\x32 .prodvana.version.SourceMetadata\x12\x14\n\x0c\x62\x61se_version\x18\x05 \x01(\t\"b\n\x18\x43onfigureApplicationResp\x12)\n\x04meta\x18\x01 \x01(\x0b\x32\x1b.prodvana.object.ObjectMeta\x12\x1b\n\x13\x63reated_new_version\x18\x02 \x01(\x08\"\xe1\x01\n ValidateConfigureApplicationResp\x12\x37\n\x06\x63onfig\x18\x01 \x01(\x0b\x32\'.prodvana.application.ApplicationConfig\x12@\n\x0f\x63ompiled_config\x18\x02 \x01(\x0b\x32\'.prodvana.application.ApplicationConfig\x12\x42\n\x11\x64\x61ngerous_actions\x18\x03 \x03(\x0b\x32\'.prodvana.common_config.DangerousAction\"H\n\x17GetApplicationConfigReq\x12\x1c\n\x0b\x61pplication\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x0f\n\x07version\x18\x02 \x01(\t\"\xa6\x01\n\x18GetApplicationConfigResp\x12\x37\n\x06\x63onfig\x18\x01 \x01(\x0b\x32\'.prodvana.application.ApplicationConfig\x12\x0f\n\x07version\x18\x02 \x01(\t\x12@\n\x0f\x63ompiled_config\x18\x03 \x01(\x0b\x32\'.prodvana.application.ApplicationConfig\"g\n ListApplicationConfigVersionsReq\x12\x1c\n\x0b\x61pplication\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x12\n\npage_token\x18\x02 \x01(\t\x12\x11\n\tpage_size\x18\x03 \x01(\x05\"\xd9\x02\n!ListApplicationConfigVersionsResp\x12Y\n\x08versions\x18\x01 \x03(\x0b\x32G.prodvana.application.ListApplicationConfigVersionsResp.VersionMetadata\x12\x17\n\x0fnext_page_token\x18\x02 \x01(\t\x1a\xbf\x01\n\x0fVersionMetadata\x12\x0f\n\x07version\x18\x01 \x01(\t\x12\x36\n\x12\x63reation_timestamp\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12(\n\x06source\x18\x03 \x01(\x0e\x32\x18.prodvana.version.Source\x12\x39\n\x0fsource_metadata\x18\x04 \x01(\x0b\x32 .prodvana.version.SourceMetadata\"4\n\x14\x44\x65leteApplicationReq\x12\x1c\n\x0b\x61pplication\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\"\x17\n\x15\x44\x65leteApplicationResp\"\xa0\x01\n\x18GetApplicationMetricsReq\x12\x1c\n\x0b\x61pplication\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x33\n\x0fstart_timestamp\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x31\n\rend_timestamp\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\\\n\x19GetApplicationMetricsResp\x12?\n\x12\x64\x65ployment_metrics\x18\x01 \x01(\x0b\x32#.prodvana.metrics.DeploymentMetrics\"9\n\x19GetApplicationInsightsReq\x12\x1c\n\x0b\x61pplication\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\"J\n\x1aGetApplicationInsightsResp\x12,\n\x08insights\x18\x01 \x03(\x0b\x32\x1a.prodvana.insights.Insight\"\xa7\x01\n\x1bSnoozeApplicationInsightReq\x12\x1c\n\x0b\x61pplication\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x31\n\x05\x63lass\x18\x02 \x01(\x0e\x32\x18.prodvana.insights.ClassB\x08\xfa\x42\x05\x82\x01\x02\x10\x01\x12\x37\n\x08\x64uration\x18\x03 \x01(\x0b\x32\x19.google.protobuf.DurationB\n\xfa\x42\x07\xaa\x01\x04\x08\x01*\x00\"\x1e\n\x1cSnoozeApplicationInsightResp\"9\n\x19GetApplicationMetadataReq\x12\x1c\n\x0b\x61pplication\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\"]\n\x1aGetApplicationMetadataResp\x12?\n\x08metadata\x18\x01 \x01(\x0b\x32-.prodvana.application.ApplicationUserMetadata\"\x84\x01\n\x19SetApplicationMetadataReq\x12\x1c\n\x0b\x61pplication\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12I\n\x08metadata\x18\x02 \x01(\x0b\x32-.prodvana.application.ApplicationUserMetadataB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\"\x1c\n\x1aSetApplicationMetadataResp2\x8b\x10\n\x12\x41pplicationManager\x12\x9c\x01\n\x14\x43onfigureApplication\x12-.prodvana.application.ConfigureApplicationReq\x1a..prodvana.application.ConfigureApplicationResp\"%\x82\xd3\xe4\x93\x02\x1f\"\x1a/v1/applications/configure:\x01*\x12\xb5\x01\n\x1cValidateConfigureApplication\x12-.prodvana.application.ConfigureApplicationReq\x1a\x36.prodvana.application.ValidateConfigureApplicationResp\".\x82\xd3\xe4\x93\x02(\"#/v1/applications/configure/validate:\x01*\x12\x83\x01\n\x10ListApplications\x12).prodvana.application.ListApplicationsReq\x1a*.prodvana.application.ListApplicationsResp\"\x18\x82\xd3\xe4\x93\x02\x12\x12\x10/v1/applications\x12\xa6\x01\n\x14GetApplicationConfig\x12-.prodvana.application.GetApplicationConfigReq\x1a..prodvana.application.GetApplicationConfigResp\"/\x82\xd3\xe4\x93\x02)\x12\'/v1/applications/{application=*}/config\x12\xca\x01\n\x1dListApplicationConfigVersions\x12\x36.prodvana.application.ListApplicationConfigVersionsReq\x1a\x37.prodvana.application.ListApplicationConfigVersionsResp\"8\x82\xd3\xe4\x93\x02\x32\x12\x30/v1/applications/{application=*}/config/versions\x12\x8d\x01\n\x0eGetApplication\x12\'.prodvana.application.GetApplicationReq\x1a(.prodvana.application.GetApplicationResp\"(\x82\xd3\xe4\x93\x02\"\x12 /v1/applications/{application=*}\x12\x90\x01\n\x11\x44\x65leteApplication\x12*.prodvana.application.DeleteApplicationReq\x1a+.prodvana.application.DeleteApplicationResp\"\"\x82\xd3\xe4\x93\x02\x1c*\x1a/v1/{application=*}/delete\x12\xaa\x01\n\x15GetApplicationMetrics\x12..prodvana.application.GetApplicationMetricsReq\x1a/.prodvana.application.GetApplicationMetricsResp\"0\x82\xd3\xe4\x93\x02*\x12(/v1/applications/{application=*}/metrics\x12\xae\x01\n\x16GetApplicationInsights\x12/.prodvana.application.GetApplicationInsightsReq\x1a\x30.prodvana.application.GetApplicationInsightsResp\"1\x82\xd3\xe4\x93\x02+\x12)/v1/applications/{application=*}/insights\x12\xbb\x01\n\x18SnoozeApplicationInsight\x12\x31.prodvana.application.SnoozeApplicationInsightReq\x1a\x32.prodvana.application.SnoozeApplicationInsightResp\"8\x82\xd3\xe4\x93\x02\x32\x1a\x30/v1/applications/{application=*}/insights/snooze\x12\xae\x01\n\x16GetApplicationMetadata\x12/.prodvana.application.GetApplicationMetadataReq\x1a\x30.prodvana.application.GetApplicationMetadataResp\"1\x82\xd3\xe4\x93\x02+\x12)/v1/applications/{application=*}/metadata\x12\xb1\x01\n\x16SetApplicationMetadata\x12/.prodvana.application.SetApplicationMetadataReq\x1a\x30.prodvana.application.SetApplicationMetadataResp\"4\x82\xd3\xe4\x93\x02.\")/v1/applications/{application=*}/metadata:\x01*BPZNgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/applicationb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'prodvana.application.application_manager_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZNgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/application'
  _GETAPPLICATIONREQ.fields_by_name['application']._options = None
  _GETAPPLICATIONREQ.fields_by_name['application']._serialized_options = b'\372B\004r\002\020\001'
  _CONFIGUREAPPLICATIONREQ.fields_by_name['application_config']._options = None
  _CONFIGUREAPPLICATIONREQ.fields_by_name['application_config']._serialized_options = b'\372B\005\212\001\002\020\001'
  _GETAPPLICATIONCONFIGREQ.fields_by_name['application']._options = None
  _GETAPPLICATIONCONFIGREQ.fields_by_name['application']._serialized_options = b'\372B\004r\002\020\001'
  _LISTAPPLICATIONCONFIGVERSIONSREQ.fields_by_name['application']._options = None
  _LISTAPPLICATIONCONFIGVERSIONSREQ.fields_by_name['application']._serialized_options = b'\372B\004r\002\020\001'
  _DELETEAPPLICATIONREQ.fields_by_name['application']._options = None
  _DELETEAPPLICATIONREQ.fields_by_name['application']._serialized_options = b'\372B\004r\002\020\001'
  _GETAPPLICATIONMETRICSREQ.fields_by_name['application']._options = None
  _GETAPPLICATIONMETRICSREQ.fields_by_name['application']._serialized_options = b'\372B\004r\002\020\001'
  _GETAPPLICATIONINSIGHTSREQ.fields_by_name['application']._options = None
  _GETAPPLICATIONINSIGHTSREQ.fields_by_name['application']._serialized_options = b'\372B\004r\002\020\001'
  _SNOOZEAPPLICATIONINSIGHTREQ.fields_by_name['application']._options = None
  _SNOOZEAPPLICATIONINSIGHTREQ.fields_by_name['application']._serialized_options = b'\372B\004r\002\020\001'
  _SNOOZEAPPLICATIONINSIGHTREQ.fields_by_name['class']._options = None
  _SNOOZEAPPLICATIONINSIGHTREQ.fields_by_name['class']._serialized_options = b'\372B\005\202\001\002\020\001'
  _SNOOZEAPPLICATIONINSIGHTREQ.fields_by_name['duration']._options = None
  _SNOOZEAPPLICATIONINSIGHTREQ.fields_by_name['duration']._serialized_options = b'\372B\007\252\001\004\010\001*\000'
  _GETAPPLICATIONMETADATAREQ.fields_by_name['application']._options = None
  _GETAPPLICATIONMETADATAREQ.fields_by_name['application']._serialized_options = b'\372B\004r\002\020\001'
  _SETAPPLICATIONMETADATAREQ.fields_by_name['application']._options = None
  _SETAPPLICATIONMETADATAREQ.fields_by_name['application']._serialized_options = b'\372B\004r\002\020\001'
  _SETAPPLICATIONMETADATAREQ.fields_by_name['metadata']._options = None
  _SETAPPLICATIONMETADATAREQ.fields_by_name['metadata']._serialized_options = b'\372B\005\212\001\002\020\001'
  _APPLICATIONMANAGER.methods_by_name['ConfigureApplication']._options = None
  _APPLICATIONMANAGER.methods_by_name['ConfigureApplication']._serialized_options = b'\202\323\344\223\002\037\"\032/v1/applications/configure:\001*'
  _APPLICATIONMANAGER.methods_by_name['ValidateConfigureApplication']._options = None
  _APPLICATIONMANAGER.methods_by_name['ValidateConfigureApplication']._serialized_options = b'\202\323\344\223\002(\"#/v1/applications/configure/validate:\001*'
  _APPLICATIONMANAGER.methods_by_name['ListApplications']._options = None
  _APPLICATIONMANAGER.methods_by_name['ListApplications']._serialized_options = b'\202\323\344\223\002\022\022\020/v1/applications'
  _APPLICATIONMANAGER.methods_by_name['GetApplicationConfig']._options = None
  _APPLICATIONMANAGER.methods_by_name['GetApplicationConfig']._serialized_options = b'\202\323\344\223\002)\022\'/v1/applications/{application=*}/config'
  _APPLICATIONMANAGER.methods_by_name['ListApplicationConfigVersions']._options = None
  _APPLICATIONMANAGER.methods_by_name['ListApplicationConfigVersions']._serialized_options = b'\202\323\344\223\0022\0220/v1/applications/{application=*}/config/versions'
  _APPLICATIONMANAGER.methods_by_name['GetApplication']._options = None
  _APPLICATIONMANAGER.methods_by_name['GetApplication']._serialized_options = b'\202\323\344\223\002\"\022 /v1/applications/{application=*}'
  _APPLICATIONMANAGER.methods_by_name['DeleteApplication']._options = None
  _APPLICATIONMANAGER.methods_by_name['DeleteApplication']._serialized_options = b'\202\323\344\223\002\034*\032/v1/{application=*}/delete'
  _APPLICATIONMANAGER.methods_by_name['GetApplicationMetrics']._options = None
  _APPLICATIONMANAGER.methods_by_name['GetApplicationMetrics']._serialized_options = b'\202\323\344\223\002*\022(/v1/applications/{application=*}/metrics'
  _APPLICATIONMANAGER.methods_by_name['GetApplicationInsights']._options = None
  _APPLICATIONMANAGER.methods_by_name['GetApplicationInsights']._serialized_options = b'\202\323\344\223\002+\022)/v1/applications/{application=*}/insights'
  _APPLICATIONMANAGER.methods_by_name['SnoozeApplicationInsight']._options = None
  _APPLICATIONMANAGER.methods_by_name['SnoozeApplicationInsight']._serialized_options = b'\202\323\344\223\0022\0320/v1/applications/{application=*}/insights/snooze'
  _APPLICATIONMANAGER.methods_by_name['GetApplicationMetadata']._options = None
  _APPLICATIONMANAGER.methods_by_name['GetApplicationMetadata']._serialized_options = b'\202\323\344\223\002+\022)/v1/applications/{application=*}/metadata'
  _APPLICATIONMANAGER.methods_by_name['SetApplicationMetadata']._options = None
  _APPLICATIONMANAGER.methods_by_name['SetApplicationMetadata']._serialized_options = b'\202\323\344\223\002.\")/v1/applications/{application=*}/metadata:\001*'
  _globals['_LISTAPPLICATIONSREQ']._serialized_start=497
  _globals['_LISTAPPLICATIONSREQ']._serialized_end=536
  _globals['_LISTAPPLICATIONSRESP']._serialized_start=538
  _globals['_LISTAPPLICATIONSRESP']._serialized_end=617
  _globals['_GETAPPLICATIONREQ']._serialized_start=619
  _globals['_GETAPPLICATIONREQ']._serialized_end=668
  _globals['_GETAPPLICATIONRESP']._serialized_start=670
  _globals['_GETAPPLICATIONRESP']._serialized_end=746
  _globals['_CONFIGUREAPPLICATIONREQ']._serialized_start=749
  _globals['_CONFIGUREAPPLICATIONREQ']._serialized_end=1015
  _globals['_CONFIGUREAPPLICATIONRESP']._serialized_start=1017
  _globals['_CONFIGUREAPPLICATIONRESP']._serialized_end=1115
  _globals['_VALIDATECONFIGUREAPPLICATIONRESP']._serialized_start=1118
  _globals['_VALIDATECONFIGUREAPPLICATIONRESP']._serialized_end=1343
  _globals['_GETAPPLICATIONCONFIGREQ']._serialized_start=1345
  _globals['_GETAPPLICATIONCONFIGREQ']._serialized_end=1417
  _globals['_GETAPPLICATIONCONFIGRESP']._serialized_start=1420
  _globals['_GETAPPLICATIONCONFIGRESP']._serialized_end=1586
  _globals['_LISTAPPLICATIONCONFIGVERSIONSREQ']._serialized_start=1588
  _globals['_LISTAPPLICATIONCONFIGVERSIONSREQ']._serialized_end=1691
  _globals['_LISTAPPLICATIONCONFIGVERSIONSRESP']._serialized_start=1694
  _globals['_LISTAPPLICATIONCONFIGVERSIONSRESP']._serialized_end=2039
  _globals['_LISTAPPLICATIONCONFIGVERSIONSRESP_VERSIONMETADATA']._serialized_start=1848
  _globals['_LISTAPPLICATIONCONFIGVERSIONSRESP_VERSIONMETADATA']._serialized_end=2039
  _globals['_DELETEAPPLICATIONREQ']._serialized_start=2041
  _globals['_DELETEAPPLICATIONREQ']._serialized_end=2093
  _globals['_DELETEAPPLICATIONRESP']._serialized_start=2095
  _globals['_DELETEAPPLICATIONRESP']._serialized_end=2118
  _globals['_GETAPPLICATIONMETRICSREQ']._serialized_start=2121
  _globals['_GETAPPLICATIONMETRICSREQ']._serialized_end=2281
  _globals['_GETAPPLICATIONMETRICSRESP']._serialized_start=2283
  _globals['_GETAPPLICATIONMETRICSRESP']._serialized_end=2375
  _globals['_GETAPPLICATIONINSIGHTSREQ']._serialized_start=2377
  _globals['_GETAPPLICATIONINSIGHTSREQ']._serialized_end=2434
  _globals['_GETAPPLICATIONINSIGHTSRESP']._serialized_start=2436
  _globals['_GETAPPLICATIONINSIGHTSRESP']._serialized_end=2510
  _globals['_SNOOZEAPPLICATIONINSIGHTREQ']._serialized_start=2513
  _globals['_SNOOZEAPPLICATIONINSIGHTREQ']._serialized_end=2680
  _globals['_SNOOZEAPPLICATIONINSIGHTRESP']._serialized_start=2682
  _globals['_SNOOZEAPPLICATIONINSIGHTRESP']._serialized_end=2712
  _globals['_GETAPPLICATIONMETADATAREQ']._serialized_start=2714
  _globals['_GETAPPLICATIONMETADATAREQ']._serialized_end=2771
  _globals['_GETAPPLICATIONMETADATARESP']._serialized_start=2773
  _globals['_GETAPPLICATIONMETADATARESP']._serialized_end=2866
  _globals['_SETAPPLICATIONMETADATAREQ']._serialized_start=2869
  _globals['_SETAPPLICATIONMETADATAREQ']._serialized_end=3001
  _globals['_SETAPPLICATIONMETADATARESP']._serialized_start=3003
  _globals['_SETAPPLICATIONMETADATARESP']._serialized_end=3031
  _globals['_APPLICATIONMANAGER']._serialized_start=3034
  _globals['_APPLICATIONMANAGER']._serialized_end=5093
# @@protoc_insertion_point(module_scope)
