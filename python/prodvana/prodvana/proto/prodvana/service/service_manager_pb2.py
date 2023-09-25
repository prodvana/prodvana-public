# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/service/service_manager.proto
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
from prodvana.proto.prodvana.common_config import parameters_pb2 as prodvana_dot_common__config_dot_parameters__pb2
from prodvana.proto.prodvana.convergence import convergence_mode_pb2 as prodvana_dot_convergence_dot_convergence__mode__pb2
from prodvana.proto.prodvana.service import parameters_pb2 as prodvana_dot_service_dot_parameters__pb2
from prodvana.proto.prodvana.service import service_config_pb2 as prodvana_dot_service_dot_service__config__pb2
from prodvana.proto.prodvana.service import user_metadata_pb2 as prodvana_dot_service_dot_user__metadata__pb2
from prodvana.proto.prodvana.service import object_pb2 as prodvana_dot_service_dot_object__pb2
from prodvana.proto.prodvana.stat import efficiency_pb2 as prodvana_dot_stat_dot_efficiency__pb2
from prodvana.proto.prodvana.metrics import metrics_pb2 as prodvana_dot_metrics_dot_metrics__pb2
from prodvana.proto.prodvana.insights import insights_pb2 as prodvana_dot_insights_dot_insights__pb2
from prodvana.proto.prodvana.repo import repo_pb2 as prodvana_dot_repo_dot_repo__pb2
from prodvana.proto.prodvana.version import source_metadata_pb2 as prodvana_dot_version_dot_source__metadata__pb2
from prodvana.proto.validate import validate_pb2 as validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n&prodvana/service/service_manager.proto\x12\x10prodvana.service\x1a\x1cgoogle/api/annotations.proto\x1a\x1egoogle/protobuf/duration.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\'prodvana/common_config/parameters.proto\x1a+prodvana/convergence/convergence_mode.proto\x1a!prodvana/service/parameters.proto\x1a%prodvana/service/service_config.proto\x1a$prodvana/service/user_metadata.proto\x1a\x1dprodvana/service/object.proto\x1a\x1eprodvana/stat/efficiency.proto\x1a\x1eprodvana/metrics/metrics.proto\x1a prodvana/insights/insights.proto\x1a\x18prodvana/repo/repo.proto\x1a&prodvana/version/source_metadata.proto\x1a\x17validate/validate.proto\"w\n\x1dServiceConfigVersionReference\x12\x13\n\x0b\x61pplication\x18\x01 \x01(\t\x12\x18\n\x07service\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\'\n\x16service_config_version\x18\x03 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\"\xa8\x02\n\x16GenerateVersionNameReq\x12\x1c\n\x0b\x61pplication\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x18\n\x07service\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\'\n\x16service_config_version\x18\x03 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12I\n\nparameters\x18\x04 \x03(\x0b\x32&.prodvana.common_config.ParameterValueB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\x12\x62\n\x13per_release_channel\x18\x05 \x03(\x0b\x32\x36.prodvana.service.ApplyParametersReq.PerReleaseChannelB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\"*\n\x17GenerateVersionNameResp\x12\x0f\n\x07version\x18\x01 \x01(\t\"\xf4\x05\n\x12\x41pplyParametersReq\x12\x39\n\x0eservice_config\x18\x01 \x01(\x0b\x32\x1f.prodvana.service.ServiceConfigH\x00\x12Q\n\x16service_config_version\x18\x08 \x01(\x0b\x32/.prodvana.service.ServiceConfigVersionReferenceH\x00\x12I\n\nparameters\x18\t \x03(\x0b\x32&.prodvana.common_config.ParameterValueB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\x12\x62\n\x13per_release_channel\x18\n \x03(\x0b\x32\x36.prodvana.service.ApplyParametersReq.PerReleaseChannelB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\x12%\n\x1dtest_only_skip_registry_check\x18\x04 \x01(\x08\x12\x13\n\x0b\x61pplication\x18\x05 \x01(\t\x12(\n\x06source\x18\x06 \x01(\x0e\x32\x18.prodvana.version.Source\x12\x39\n\x0fsource_metadata\x18\x07 \x01(\x0b\x32 .prodvana.version.SourceMetadata\x12\x1c\n\x14\x62undle_name_override\x18\x0b \x01(\t\x12\x1f\n\x17skip_runtime_validation\x18\x0c \x01(\x08\x1a\x80\x01\n\x11PerReleaseChannel\x12 \n\x0frelease_channel\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12I\n\nparameters\x18\x02 \x03(\x0b\x32&.prodvana.common_config.ParameterValueB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\x42\x0c\n\x05oneof\x12\x03\xf8\x42\x01J\x04\x08\x02\x10\x03J\x04\x08\x03\x10\x04R\x10release_channelsR\x12\x63ompute_efficiency\"r\n\x13\x41pplyParametersResp\x12\x12\n\nservice_id\x18\x01 \x01(\t\x12\x0f\n\x07version\x18\x02 \x01(\t\x12\x36\n\x0f\x65\x66\x66iciency_stat\x18\x03 \x01(\x0b\x32\x1d.prodvana.stat.EfficiencyStat\"\xf9\x01\n\x1bValidateApplyParametersResp\x12/\n\x06\x63onfig\x18\x01 \x01(\x0b\x32\x1f.prodvana.service.ServiceConfig\x12\x38\n\x0f\x63ompiled_config\x18\x02 \x01(\x0b\x32\x1f.prodvana.service.ServiceConfig\x12\x13\n\x0b\x61pp_version\x18\x03 \x01(\t\x12Z\n!compiled_service_instance_configs\x18\x04 \x03(\x0b\x32/.prodvana.service.CompiledServiceInstanceConfig\"Z\n\x18GetMaterializedConfigReq\x12\x18\n\x07service\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x0f\n\x07version\x18\x02 \x01(\t\x12\x13\n\x0b\x61pplication\x18\x03 \x01(\t\"\xd3\x02\n\x19GetMaterializedConfigResp\x12/\n\x06\x63onfig\x18\x01 \x01(\x0b\x32\x1f.prodvana.service.ServiceConfig\x12\x0f\n\x07version\x18\x02 \x01(\t\x12\x38\n\x0f\x63ompiled_config\x18\x03 \x01(\x0b\x32\x1f.prodvana.service.ServiceConfig\x12Z\n!compiled_service_instance_configs\x18\x04 \x03(\x0b\x32/.prodvana.service.CompiledServiceInstanceConfig\x12^\n\x10version_metadata\x18\x05 \x01(\x0b\x32\x44.prodvana.service.ListMaterializedConfigVersionsResp.VersionMetadata\"A\n\x10\x44\x65leteServiceReq\x12\x18\n\x07service\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x13\n\x0b\x61pplication\x18\x02 \x01(\t\"\x13\n\x11\x44\x65leteServiceResp\"A\n\x0fListServicesReq\x12\x1c\n\x0b\x61pplication\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x10\n\x08\x64\x65tailed\x18\x02 \x01(\x08\"?\n\x10ListServicesResp\x12+\n\x08services\x18\x01 \x03(\x0b\x32\x19.prodvana.service.Service\"G\n\rGetServiceReq\x12\x1c\n\x0b\x61pplication\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x18\n\x07service\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\"<\n\x0eGetServiceResp\x12*\n\x07service\x18\x01 \x01(\x0b\x32\x19.prodvana.service.Service\"Q\n\x17ListServiceInstancesReq\x12\x1c\n\x0b\x61pplication\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x18\n\x07service\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\"X\n\x18ListServiceInstancesResp\x12<\n\x11service_instances\x18\x01 \x03(\x0b\x32!.prodvana.service.ServiceInstance\"q\n\x15GetServiceInstanceReq\x12\x1c\n\x0b\x61pplication\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x18\n\x07service\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12 \n\x0frelease_channel\x18\x03 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\"U\n\x16GetServiceInstanceResp\x12;\n\x10service_instance\x18\x01 \x01(\x0b\x32!.prodvana.service.ServiceInstance\"\xcc\x01\n\x14GetServiceMetricsReq\x12\x1c\n\x0b\x61pplication\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x18\n\x07service\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x33\n\x0fstart_timestamp\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x31\n\rend_timestamp\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x14\n\x0cinclude_cost\x18\x05 \x01(\x08\"~\n\x15GetServiceMetricsResp\x12?\n\x12\x64\x65ployment_metrics\x18\x01 \x01(\x0b\x32#.prodvana.metrics.DeploymentMetrics\x12$\n\x04\x63ost\x18\x02 \x01(\x0b\x32\x16.prodvana.metrics.Cost\"O\n\x15GetServiceInsightsReq\x12\x1c\n\x0b\x61pplication\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x18\n\x07service\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\"o\n\x0eListCommitsReq\x12\x1c\n\x0b\x61pplication\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x18\n\x07service\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x12\n\npage_token\x18\x03 \x01(\t\x12\x11\n\tpage_size\x18\x04 \x01(\x05\"R\n\x0fListCommitsResp\x12&\n\x07\x63ommits\x18\x01 \x03(\x0b\x32\x15.prodvana.repo.Commit\x12\x17\n\x0fnext_page_token\x18\x02 \x01(\t\"F\n\x16GetServiceInsightsResp\x12,\n\x08insights\x18\x01 \x03(\x0b\x32\x1a.prodvana.insights.Insight\"\xbd\x01\n\x17SnoozeServiceInsightReq\x12\x1c\n\x0b\x61pplication\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x18\n\x07service\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x31\n\x05\x63lass\x18\x03 \x01(\x0e\x32\x18.prodvana.insights.ClassB\x08\xfa\x42\x05\x82\x01\x02\x10\x01\x12\x37\n\x08\x64uration\x18\x04 \x01(\x0b\x32\x19.google.protobuf.DurationB\n\xfa\x42\x07\xaa\x01\x04\x08\x01*\x00\"\x1a\n\x18SnoozeServiceInsightResp\"O\n\x15GetServiceMetadataReq\x12\x1c\n\x0b\x61pplication\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x18\n\x07service\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\"Q\n\x16GetServiceMetadataResp\x12\x37\n\x08metadata\x18\x01 \x01(\x0b\x32%.prodvana.service.ServiceUserMetadata\"\x92\x01\n\x15SetServiceMetadataReq\x12\x1c\n\x0b\x61pplication\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x18\n\x07service\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x41\n\x08metadata\x18\x03 \x01(\x0b\x32%.prodvana.service.ServiceUserMetadataB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\"\x18\n\x16SetServiceMetadataResp\"\x97\x01\n\x1cSetServiceConvergenceModeReq\x12\x1c\n\x0b\x61pplication\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x18\n\x07service\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12?\n\x10\x63onvergence_mode\x18\x03 \x01(\x0e\x32%.prodvana.convergence.ConvergenceMode\"\x1f\n\x1dSetServiceConvergenceModeResp\"\x82\x01\n!ListMaterializedConfigVersionsReq\x12\x1c\n\x0b\x61pplication\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x18\n\x07service\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x12\n\npage_token\x18\x03 \x01(\t\x12\x11\n\tpage_size\x18\x04 \x01(\x05\"\xf4\x03\n\"ListMaterializedConfigVersionsResp\x12V\n\x08versions\x18\x01 \x03(\x0b\x32\x44.prodvana.service.ListMaterializedConfigVersionsResp.VersionMetadata\x12\x17\n\x0fnext_page_token\x18\x02 \x01(\t\x1a\xdc\x02\n\x0fVersionMetadata\x12\x0f\n\x07version\x18\x01 \x01(\t\x12\x36\n\x12\x63reation_timestamp\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x16\n\x0e\x63onfig_version\x18\x03 \x01(\t\x12?\n\nparameters\x18\x07 \x03(\x0b\x32+.prodvana.common_config.ParameterDefinition\x12\x42\n\x10parameter_values\x18\x04 \x01(\x0b\x32(.prodvana.service.ServiceParameterValues\x12(\n\x06source\x18\x05 \x01(\x0e\x32\x18.prodvana.version.Source\x12\x39\n\x0fsource_metadata\x18\x06 \x01(\x0b\x32 .prodvana.version.SourceMetadata\"\x9a\x02\n\x13\x43onfigureServiceReq\x12\x1c\n\x0b\x61pplication\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x41\n\x0eservice_config\x18\x02 \x01(\x0b\x32\x1f.prodvana.service.ServiceConfigB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12(\n\x06source\x18\x03 \x01(\x0e\x32\x18.prodvana.version.Source\x12\x39\n\x0fsource_metadata\x18\x04 \x01(\x0b\x32 .prodvana.version.SourceMetadata\x12\x1c\n\x14\x62undle_name_override\x18\x05 \x01(\t\x12\x1f\n\x17skip_runtime_validation\x18\x06 \x01(\x08\"B\n\x14\x43onfigureServiceResp\x12\x12\n\nservice_id\x18\x01 \x01(\t\x12\x16\n\x0e\x63onfig_version\x18\x02 \x01(\t\"\x8f\x01\n\x1cValidateConfigureServiceResp\x12\x35\n\x0cinput_config\x18\x01 \x01(\x0b\x32\x1f.prodvana.service.ServiceConfig\x12\x38\n\x0f\x63ompiled_config\x18\x02 \x01(\x0b\x32\x1f.prodvana.service.ServiceConfig\"}\n\x1cListServiceConfigVersionsReq\x12\x1c\n\x0b\x61pplication\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x18\n\x07service\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x12\n\npage_token\x18\x03 \x01(\t\x12\x11\n\tpage_size\x18\x04 \x01(\x05\"\xcd\x02\n\x1dListServiceConfigVersionsResp\x12Q\n\x08versions\x18\x01 \x03(\x0b\x32?.prodvana.service.ListServiceConfigVersionsResp.VersionMetadata\x12\x17\n\x0fnext_page_token\x18\x02 \x01(\t\x1a\xbf\x01\n\x0fVersionMetadata\x12\x0f\n\x07version\x18\x01 \x01(\t\x12\x36\n\x12\x63reation_timestamp\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12(\n\x06source\x18\x03 \x01(\x0e\x32\x18.prodvana.version.Source\x12\x39\n\x0fsource_metadata\x18\x04 \x01(\x0b\x32 .prodvana.version.SourceMetadata\"e\n\x13GetServiceConfigReq\x12\x1c\n\x0b\x61pplication\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x18\n\x07service\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x16\n\x0e\x63onfig_version\x18\x03 \x01(\t\"\xb2\x02\n\x14GetServiceConfigResp\x12/\n\x06\x63onfig\x18\x01 \x01(\x0b\x32\x1f.prodvana.service.ServiceConfig\x12\x35\n\x0cinput_config\x18\x04 \x01(\x0b\x32\x1f.prodvana.service.ServiceConfig\x12\x38\n\x0f\x63ompiled_config\x18\x05 \x01(\x0b\x32\x1f.prodvana.service.ServiceConfig\x12\x16\n\x0e\x63onfig_version\x18\x02 \x01(\t\x12`\n\x17\x63onfig_version_metadata\x18\x03 \x01(\x0b\x32?.prodvana.service.ListServiceConfigVersionsResp.VersionMetadata2\x82\x1d\n\x0eServiceManager\x12\x95\x01\n\x10\x43onfigureService\x12%.prodvana.service.ConfigureServiceReq\x1a&.prodvana.service.ConfigureServiceResp\"2\x82\xd3\xe4\x93\x02,\"\'/v1/{application=*}/services/configure2:\x01*\x12\xad\x01\n\x18ValidateConfigureService\x12%.prodvana.service.ConfigureServiceReq\x1a..prodvana.service.ValidateConfigureServiceResp\":\x82\xd3\xe4\x93\x02\x34\"//v1/{application=*}/services/configure/validate:\x01*\x12\xbe\x01\n\x19ListServiceConfigVersions\x12..prodvana.service.ListServiceConfigVersionsReq\x1a/.prodvana.service.ListServiceConfigVersionsResp\"@\x82\xd3\xe4\x93\x02:\x12\x38/v1/{application=*}/services/{service=*}/config/versions\x12\x9a\x01\n\x10GetServiceConfig\x12%.prodvana.service.GetServiceConfigReq\x1a&.prodvana.service.GetServiceConfigResp\"7\x82\xd3\xe4\x93\x02\x31\x12//v1/{application=*}/services/{service=*}/config\x12\xb2\x01\n\x13GenerateVersionName\x12(.prodvana.service.GenerateVersionNameReq\x1a).prodvana.service.GenerateVersionNameResp\"F\x82\xd3\xe4\x93\x02@\x12>/v1/{application=*}/services/{service=*}/generate-version-name\x12\x98\x01\n\x0f\x41pplyParameters\x12$.prodvana.service.ApplyParametersReq\x1a%.prodvana.service.ApplyParametersResp\"8\x82\xd3\xe4\x93\x02\x32\"-/v1/{application=*}/services/apply-parameters:\x01*\x12\xb1\x01\n\x17ValidateApplyParameters\x12$.prodvana.service.ApplyParametersReq\x1a-.prodvana.service.ValidateApplyParametersResp\"A\x82\xd3\xe4\x93\x02;\"6/v1/{application=*}/services/apply-parameters/validate:\x01*\x12\xb6\x01\n\x15GetMaterializedConfig\x12*.prodvana.service.GetMaterializedConfigReq\x1a+.prodvana.service.GetMaterializedConfigResp\"D\x82\xd3\xe4\x93\x02>\x12</v1/{application=*}/services/{service=*}/materialized/config\x12\xda\x01\n\x1eListMaterializedConfigVersions\x12\x33.prodvana.service.ListMaterializedConfigVersionsReq\x1a\x34.prodvana.service.ListMaterializedConfigVersionsResp\"M\x82\xd3\xe4\x93\x02G\x12\x45/v1/{application=*}/services/{service=*}/materialized/config/versions\x12\x91\x01\n\rDeleteService\x12\".prodvana.service.DeleteServiceReq\x1a#.prodvana.service.DeleteServiceResp\"7\x82\xd3\xe4\x93\x02\x31*//v1/{application=*}/services/{service=*}/delete\x12\x88\x01\n\x0cListServices\x12!.prodvana.service.ListServicesReq\x1a\".prodvana.service.ListServicesResp\"1\x82\xd3\xe4\x93\x02+\x12)/v1/applications/{application=*}/services\x12\x99\x01\n\x0bListCommits\x12 .prodvana.service.ListCommitsReq\x1a!.prodvana.service.ListCommitsResp\"E\x82\xd3\xe4\x93\x02?\x12=/v1/applications/{application=*}/services/{service=*}/commits\x12\x8e\x01\n\nGetService\x12\x1f.prodvana.service.GetServiceReq\x1a .prodvana.service.GetServiceResp\"=\x82\xd3\xe4\x93\x02\x37\x12\x35/v1/applications/{application=*}/services/{service=*}\x12\xbd\x01\n\x14ListServiceInstances\x12).prodvana.service.ListServiceInstancesReq\x1a*.prodvana.service.ListServiceInstancesResp\"N\x82\xd3\xe4\x93\x02H\x12\x46/v1/applications/{application=*}/services/{service=*}/release-channels\x12\xcb\x01\n\x12GetServiceInstance\x12\'.prodvana.service.GetServiceInstanceReq\x1a(.prodvana.service.GetServiceInstanceResp\"b\x82\xd3\xe4\x93\x02\\\x12Z/v1/applications/{application=*}/services/{service=*}/release-channels/{release_channel=*}\x12\xab\x01\n\x11GetServiceMetrics\x12&.prodvana.service.GetServiceMetricsReq\x1a\'.prodvana.service.GetServiceMetricsResp\"E\x82\xd3\xe4\x93\x02?\x12=/v1/applications/{application=*}/services/{service=*}/metrics\x12\xaf\x01\n\x12GetServiceInsights\x12\'.prodvana.service.GetServiceInsightsReq\x1a(.prodvana.service.GetServiceInsightsResp\"F\x82\xd3\xe4\x93\x02@\x12>/v1/applications/{application=*}/services/{service=*}/insights\x12\xbc\x01\n\x14SnoozeServiceInsight\x12).prodvana.service.SnoozeServiceInsightReq\x1a*.prodvana.service.SnoozeServiceInsightResp\"M\x82\xd3\xe4\x93\x02G\x1a\x45/v1/applications/{application=*}/services/{service=*}/insights/snooze\x12\xaf\x01\n\x12GetServiceMetadata\x12\'.prodvana.service.GetServiceMetadataReq\x1a(.prodvana.service.GetServiceMetadataResp\"F\x82\xd3\xe4\x93\x02@\x12>/v1/applications/{application=*}/services/{service=*}/metadata\x12\xb2\x01\n\x12SetServiceMetadata\x12\'.prodvana.service.SetServiceMetadataReq\x1a(.prodvana.service.SetServiceMetadataResp\"I\x82\xd3\xe4\x93\x02\x43\">/v1/applications/{application=*}/services/{service=*}/metadata:\x01*\x12\xcf\x01\n\x19SetServiceConvergenceMode\x12..prodvana.service.SetServiceConvergenceModeReq\x1a/.prodvana.service.SetServiceConvergenceModeResp\"Q\x82\xd3\xe4\x93\x02K\"F/v1/applications/{application=*}/services/{service=*}/convergence-mode:\x01*BLZJgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/serviceb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'prodvana.service.service_manager_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZJgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/service'
  _SERVICECONFIGVERSIONREFERENCE.fields_by_name['service']._options = None
  _SERVICECONFIGVERSIONREFERENCE.fields_by_name['service']._serialized_options = b'\372B\004r\002\020\001'
  _SERVICECONFIGVERSIONREFERENCE.fields_by_name['service_config_version']._options = None
  _SERVICECONFIGVERSIONREFERENCE.fields_by_name['service_config_version']._serialized_options = b'\372B\004r\002\020\001'
  _GENERATEVERSIONNAMEREQ.fields_by_name['application']._options = None
  _GENERATEVERSIONNAMEREQ.fields_by_name['application']._serialized_options = b'\372B\004r\002\020\001'
  _GENERATEVERSIONNAMEREQ.fields_by_name['service']._options = None
  _GENERATEVERSIONNAMEREQ.fields_by_name['service']._serialized_options = b'\372B\004r\002\020\001'
  _GENERATEVERSIONNAMEREQ.fields_by_name['service_config_version']._options = None
  _GENERATEVERSIONNAMEREQ.fields_by_name['service_config_version']._serialized_options = b'\372B\004r\002\020\001'
  _GENERATEVERSIONNAMEREQ.fields_by_name['parameters']._options = None
  _GENERATEVERSIONNAMEREQ.fields_by_name['parameters']._serialized_options = b'\372B\n\222\001\007\"\005\212\001\002\020\001'
  _GENERATEVERSIONNAMEREQ.fields_by_name['per_release_channel']._options = None
  _GENERATEVERSIONNAMEREQ.fields_by_name['per_release_channel']._serialized_options = b'\372B\n\222\001\007\"\005\212\001\002\020\001'
  _APPLYPARAMETERSREQ_PERRELEASECHANNEL.fields_by_name['release_channel']._options = None
  _APPLYPARAMETERSREQ_PERRELEASECHANNEL.fields_by_name['release_channel']._serialized_options = b'\372B\004r\002\020\001'
  _APPLYPARAMETERSREQ_PERRELEASECHANNEL.fields_by_name['parameters']._options = None
  _APPLYPARAMETERSREQ_PERRELEASECHANNEL.fields_by_name['parameters']._serialized_options = b'\372B\n\222\001\007\"\005\212\001\002\020\001'
  _APPLYPARAMETERSREQ.oneofs_by_name['oneof']._options = None
  _APPLYPARAMETERSREQ.oneofs_by_name['oneof']._serialized_options = b'\370B\001'
  _APPLYPARAMETERSREQ.fields_by_name['parameters']._options = None
  _APPLYPARAMETERSREQ.fields_by_name['parameters']._serialized_options = b'\372B\n\222\001\007\"\005\212\001\002\020\001'
  _APPLYPARAMETERSREQ.fields_by_name['per_release_channel']._options = None
  _APPLYPARAMETERSREQ.fields_by_name['per_release_channel']._serialized_options = b'\372B\n\222\001\007\"\005\212\001\002\020\001'
  _GETMATERIALIZEDCONFIGREQ.fields_by_name['service']._options = None
  _GETMATERIALIZEDCONFIGREQ.fields_by_name['service']._serialized_options = b'\372B\004r\002\020\001'
  _DELETESERVICEREQ.fields_by_name['service']._options = None
  _DELETESERVICEREQ.fields_by_name['service']._serialized_options = b'\372B\004r\002\020\001'
  _LISTSERVICESREQ.fields_by_name['application']._options = None
  _LISTSERVICESREQ.fields_by_name['application']._serialized_options = b'\372B\004r\002\020\001'
  _GETSERVICEREQ.fields_by_name['application']._options = None
  _GETSERVICEREQ.fields_by_name['application']._serialized_options = b'\372B\004r\002\020\001'
  _GETSERVICEREQ.fields_by_name['service']._options = None
  _GETSERVICEREQ.fields_by_name['service']._serialized_options = b'\372B\004r\002\020\001'
  _LISTSERVICEINSTANCESREQ.fields_by_name['application']._options = None
  _LISTSERVICEINSTANCESREQ.fields_by_name['application']._serialized_options = b'\372B\004r\002\020\001'
  _LISTSERVICEINSTANCESREQ.fields_by_name['service']._options = None
  _LISTSERVICEINSTANCESREQ.fields_by_name['service']._serialized_options = b'\372B\004r\002\020\001'
  _GETSERVICEINSTANCEREQ.fields_by_name['application']._options = None
  _GETSERVICEINSTANCEREQ.fields_by_name['application']._serialized_options = b'\372B\004r\002\020\001'
  _GETSERVICEINSTANCEREQ.fields_by_name['service']._options = None
  _GETSERVICEINSTANCEREQ.fields_by_name['service']._serialized_options = b'\372B\004r\002\020\001'
  _GETSERVICEINSTANCEREQ.fields_by_name['release_channel']._options = None
  _GETSERVICEINSTANCEREQ.fields_by_name['release_channel']._serialized_options = b'\372B\004r\002\020\001'
  _GETSERVICEMETRICSREQ.fields_by_name['application']._options = None
  _GETSERVICEMETRICSREQ.fields_by_name['application']._serialized_options = b'\372B\004r\002\020\001'
  _GETSERVICEMETRICSREQ.fields_by_name['service']._options = None
  _GETSERVICEMETRICSREQ.fields_by_name['service']._serialized_options = b'\372B\004r\002\020\001'
  _GETSERVICEINSIGHTSREQ.fields_by_name['application']._options = None
  _GETSERVICEINSIGHTSREQ.fields_by_name['application']._serialized_options = b'\372B\004r\002\020\001'
  _GETSERVICEINSIGHTSREQ.fields_by_name['service']._options = None
  _GETSERVICEINSIGHTSREQ.fields_by_name['service']._serialized_options = b'\372B\004r\002\020\001'
  _LISTCOMMITSREQ.fields_by_name['application']._options = None
  _LISTCOMMITSREQ.fields_by_name['application']._serialized_options = b'\372B\004r\002\020\001'
  _LISTCOMMITSREQ.fields_by_name['service']._options = None
  _LISTCOMMITSREQ.fields_by_name['service']._serialized_options = b'\372B\004r\002\020\001'
  _SNOOZESERVICEINSIGHTREQ.fields_by_name['application']._options = None
  _SNOOZESERVICEINSIGHTREQ.fields_by_name['application']._serialized_options = b'\372B\004r\002\020\001'
  _SNOOZESERVICEINSIGHTREQ.fields_by_name['service']._options = None
  _SNOOZESERVICEINSIGHTREQ.fields_by_name['service']._serialized_options = b'\372B\004r\002\020\001'
  _SNOOZESERVICEINSIGHTREQ.fields_by_name['class']._options = None
  _SNOOZESERVICEINSIGHTREQ.fields_by_name['class']._serialized_options = b'\372B\005\202\001\002\020\001'
  _SNOOZESERVICEINSIGHTREQ.fields_by_name['duration']._options = None
  _SNOOZESERVICEINSIGHTREQ.fields_by_name['duration']._serialized_options = b'\372B\007\252\001\004\010\001*\000'
  _GETSERVICEMETADATAREQ.fields_by_name['application']._options = None
  _GETSERVICEMETADATAREQ.fields_by_name['application']._serialized_options = b'\372B\004r\002\020\001'
  _GETSERVICEMETADATAREQ.fields_by_name['service']._options = None
  _GETSERVICEMETADATAREQ.fields_by_name['service']._serialized_options = b'\372B\004r\002\020\001'
  _SETSERVICEMETADATAREQ.fields_by_name['application']._options = None
  _SETSERVICEMETADATAREQ.fields_by_name['application']._serialized_options = b'\372B\004r\002\020\001'
  _SETSERVICEMETADATAREQ.fields_by_name['service']._options = None
  _SETSERVICEMETADATAREQ.fields_by_name['service']._serialized_options = b'\372B\004r\002\020\001'
  _SETSERVICEMETADATAREQ.fields_by_name['metadata']._options = None
  _SETSERVICEMETADATAREQ.fields_by_name['metadata']._serialized_options = b'\372B\005\212\001\002\020\001'
  _SETSERVICECONVERGENCEMODEREQ.fields_by_name['application']._options = None
  _SETSERVICECONVERGENCEMODEREQ.fields_by_name['application']._serialized_options = b'\372B\004r\002\020\001'
  _SETSERVICECONVERGENCEMODEREQ.fields_by_name['service']._options = None
  _SETSERVICECONVERGENCEMODEREQ.fields_by_name['service']._serialized_options = b'\372B\004r\002\020\001'
  _LISTMATERIALIZEDCONFIGVERSIONSREQ.fields_by_name['application']._options = None
  _LISTMATERIALIZEDCONFIGVERSIONSREQ.fields_by_name['application']._serialized_options = b'\372B\004r\002\020\001'
  _LISTMATERIALIZEDCONFIGVERSIONSREQ.fields_by_name['service']._options = None
  _LISTMATERIALIZEDCONFIGVERSIONSREQ.fields_by_name['service']._serialized_options = b'\372B\004r\002\020\001'
  _CONFIGURESERVICEREQ.fields_by_name['application']._options = None
  _CONFIGURESERVICEREQ.fields_by_name['application']._serialized_options = b'\372B\004r\002\020\001'
  _CONFIGURESERVICEREQ.fields_by_name['service_config']._options = None
  _CONFIGURESERVICEREQ.fields_by_name['service_config']._serialized_options = b'\372B\005\212\001\002\020\001'
  _LISTSERVICECONFIGVERSIONSREQ.fields_by_name['application']._options = None
  _LISTSERVICECONFIGVERSIONSREQ.fields_by_name['application']._serialized_options = b'\372B\004r\002\020\001'
  _LISTSERVICECONFIGVERSIONSREQ.fields_by_name['service']._options = None
  _LISTSERVICECONFIGVERSIONSREQ.fields_by_name['service']._serialized_options = b'\372B\004r\002\020\001'
  _GETSERVICECONFIGREQ.fields_by_name['application']._options = None
  _GETSERVICECONFIGREQ.fields_by_name['application']._serialized_options = b'\372B\004r\002\020\001'
  _GETSERVICECONFIGREQ.fields_by_name['service']._options = None
  _GETSERVICECONFIGREQ.fields_by_name['service']._serialized_options = b'\372B\004r\002\020\001'
  _SERVICEMANAGER.methods_by_name['ConfigureService']._options = None
  _SERVICEMANAGER.methods_by_name['ConfigureService']._serialized_options = b'\202\323\344\223\002,\"\'/v1/{application=*}/services/configure2:\001*'
  _SERVICEMANAGER.methods_by_name['ValidateConfigureService']._options = None
  _SERVICEMANAGER.methods_by_name['ValidateConfigureService']._serialized_options = b'\202\323\344\223\0024\"//v1/{application=*}/services/configure/validate:\001*'
  _SERVICEMANAGER.methods_by_name['ListServiceConfigVersions']._options = None
  _SERVICEMANAGER.methods_by_name['ListServiceConfigVersions']._serialized_options = b'\202\323\344\223\002:\0228/v1/{application=*}/services/{service=*}/config/versions'
  _SERVICEMANAGER.methods_by_name['GetServiceConfig']._options = None
  _SERVICEMANAGER.methods_by_name['GetServiceConfig']._serialized_options = b'\202\323\344\223\0021\022//v1/{application=*}/services/{service=*}/config'
  _SERVICEMANAGER.methods_by_name['GenerateVersionName']._options = None
  _SERVICEMANAGER.methods_by_name['GenerateVersionName']._serialized_options = b'\202\323\344\223\002@\022>/v1/{application=*}/services/{service=*}/generate-version-name'
  _SERVICEMANAGER.methods_by_name['ApplyParameters']._options = None
  _SERVICEMANAGER.methods_by_name['ApplyParameters']._serialized_options = b'\202\323\344\223\0022\"-/v1/{application=*}/services/apply-parameters:\001*'
  _SERVICEMANAGER.methods_by_name['ValidateApplyParameters']._options = None
  _SERVICEMANAGER.methods_by_name['ValidateApplyParameters']._serialized_options = b'\202\323\344\223\002;\"6/v1/{application=*}/services/apply-parameters/validate:\001*'
  _SERVICEMANAGER.methods_by_name['GetMaterializedConfig']._options = None
  _SERVICEMANAGER.methods_by_name['GetMaterializedConfig']._serialized_options = b'\202\323\344\223\002>\022</v1/{application=*}/services/{service=*}/materialized/config'
  _SERVICEMANAGER.methods_by_name['ListMaterializedConfigVersions']._options = None
  _SERVICEMANAGER.methods_by_name['ListMaterializedConfigVersions']._serialized_options = b'\202\323\344\223\002G\022E/v1/{application=*}/services/{service=*}/materialized/config/versions'
  _SERVICEMANAGER.methods_by_name['DeleteService']._options = None
  _SERVICEMANAGER.methods_by_name['DeleteService']._serialized_options = b'\202\323\344\223\0021*//v1/{application=*}/services/{service=*}/delete'
  _SERVICEMANAGER.methods_by_name['ListServices']._options = None
  _SERVICEMANAGER.methods_by_name['ListServices']._serialized_options = b'\202\323\344\223\002+\022)/v1/applications/{application=*}/services'
  _SERVICEMANAGER.methods_by_name['ListCommits']._options = None
  _SERVICEMANAGER.methods_by_name['ListCommits']._serialized_options = b'\202\323\344\223\002?\022=/v1/applications/{application=*}/services/{service=*}/commits'
  _SERVICEMANAGER.methods_by_name['GetService']._options = None
  _SERVICEMANAGER.methods_by_name['GetService']._serialized_options = b'\202\323\344\223\0027\0225/v1/applications/{application=*}/services/{service=*}'
  _SERVICEMANAGER.methods_by_name['ListServiceInstances']._options = None
  _SERVICEMANAGER.methods_by_name['ListServiceInstances']._serialized_options = b'\202\323\344\223\002H\022F/v1/applications/{application=*}/services/{service=*}/release-channels'
  _SERVICEMANAGER.methods_by_name['GetServiceInstance']._options = None
  _SERVICEMANAGER.methods_by_name['GetServiceInstance']._serialized_options = b'\202\323\344\223\002\\\022Z/v1/applications/{application=*}/services/{service=*}/release-channels/{release_channel=*}'
  _SERVICEMANAGER.methods_by_name['GetServiceMetrics']._options = None
  _SERVICEMANAGER.methods_by_name['GetServiceMetrics']._serialized_options = b'\202\323\344\223\002?\022=/v1/applications/{application=*}/services/{service=*}/metrics'
  _SERVICEMANAGER.methods_by_name['GetServiceInsights']._options = None
  _SERVICEMANAGER.methods_by_name['GetServiceInsights']._serialized_options = b'\202\323\344\223\002@\022>/v1/applications/{application=*}/services/{service=*}/insights'
  _SERVICEMANAGER.methods_by_name['SnoozeServiceInsight']._options = None
  _SERVICEMANAGER.methods_by_name['SnoozeServiceInsight']._serialized_options = b'\202\323\344\223\002G\032E/v1/applications/{application=*}/services/{service=*}/insights/snooze'
  _SERVICEMANAGER.methods_by_name['GetServiceMetadata']._options = None
  _SERVICEMANAGER.methods_by_name['GetServiceMetadata']._serialized_options = b'\202\323\344\223\002@\022>/v1/applications/{application=*}/services/{service=*}/metadata'
  _SERVICEMANAGER.methods_by_name['SetServiceMetadata']._options = None
  _SERVICEMANAGER.methods_by_name['SetServiceMetadata']._serialized_options = b'\202\323\344\223\002C\">/v1/applications/{application=*}/services/{service=*}/metadata:\001*'
  _SERVICEMANAGER.methods_by_name['SetServiceConvergenceMode']._options = None
  _SERVICEMANAGER.methods_by_name['SetServiceConvergenceMode']._serialized_options = b'\202\323\344\223\002K\"F/v1/applications/{application=*}/services/{service=*}/convergence-mode:\001*'
  _globals['_SERVICECONFIGVERSIONREFERENCE']._serialized_start=573
  _globals['_SERVICECONFIGVERSIONREFERENCE']._serialized_end=692
  _globals['_GENERATEVERSIONNAMEREQ']._serialized_start=695
  _globals['_GENERATEVERSIONNAMEREQ']._serialized_end=991
  _globals['_GENERATEVERSIONNAMERESP']._serialized_start=993
  _globals['_GENERATEVERSIONNAMERESP']._serialized_end=1035
  _globals['_APPLYPARAMETERSREQ']._serialized_start=1038
  _globals['_APPLYPARAMETERSREQ']._serialized_end=1794
  _globals['_APPLYPARAMETERSREQ_PERRELEASECHANNEL']._serialized_start=1602
  _globals['_APPLYPARAMETERSREQ_PERRELEASECHANNEL']._serialized_end=1730
  _globals['_APPLYPARAMETERSRESP']._serialized_start=1796
  _globals['_APPLYPARAMETERSRESP']._serialized_end=1910
  _globals['_VALIDATEAPPLYPARAMETERSRESP']._serialized_start=1913
  _globals['_VALIDATEAPPLYPARAMETERSRESP']._serialized_end=2162
  _globals['_GETMATERIALIZEDCONFIGREQ']._serialized_start=2164
  _globals['_GETMATERIALIZEDCONFIGREQ']._serialized_end=2254
  _globals['_GETMATERIALIZEDCONFIGRESP']._serialized_start=2257
  _globals['_GETMATERIALIZEDCONFIGRESP']._serialized_end=2596
  _globals['_DELETESERVICEREQ']._serialized_start=2598
  _globals['_DELETESERVICEREQ']._serialized_end=2663
  _globals['_DELETESERVICERESP']._serialized_start=2665
  _globals['_DELETESERVICERESP']._serialized_end=2684
  _globals['_LISTSERVICESREQ']._serialized_start=2686
  _globals['_LISTSERVICESREQ']._serialized_end=2751
  _globals['_LISTSERVICESRESP']._serialized_start=2753
  _globals['_LISTSERVICESRESP']._serialized_end=2816
  _globals['_GETSERVICEREQ']._serialized_start=2818
  _globals['_GETSERVICEREQ']._serialized_end=2889
  _globals['_GETSERVICERESP']._serialized_start=2891
  _globals['_GETSERVICERESP']._serialized_end=2951
  _globals['_LISTSERVICEINSTANCESREQ']._serialized_start=2953
  _globals['_LISTSERVICEINSTANCESREQ']._serialized_end=3034
  _globals['_LISTSERVICEINSTANCESRESP']._serialized_start=3036
  _globals['_LISTSERVICEINSTANCESRESP']._serialized_end=3124
  _globals['_GETSERVICEINSTANCEREQ']._serialized_start=3126
  _globals['_GETSERVICEINSTANCEREQ']._serialized_end=3239
  _globals['_GETSERVICEINSTANCERESP']._serialized_start=3241
  _globals['_GETSERVICEINSTANCERESP']._serialized_end=3326
  _globals['_GETSERVICEMETRICSREQ']._serialized_start=3329
  _globals['_GETSERVICEMETRICSREQ']._serialized_end=3533
  _globals['_GETSERVICEMETRICSRESP']._serialized_start=3535
  _globals['_GETSERVICEMETRICSRESP']._serialized_end=3661
  _globals['_GETSERVICEINSIGHTSREQ']._serialized_start=3663
  _globals['_GETSERVICEINSIGHTSREQ']._serialized_end=3742
  _globals['_LISTCOMMITSREQ']._serialized_start=3744
  _globals['_LISTCOMMITSREQ']._serialized_end=3855
  _globals['_LISTCOMMITSRESP']._serialized_start=3857
  _globals['_LISTCOMMITSRESP']._serialized_end=3939
  _globals['_GETSERVICEINSIGHTSRESP']._serialized_start=3941
  _globals['_GETSERVICEINSIGHTSRESP']._serialized_end=4011
  _globals['_SNOOZESERVICEINSIGHTREQ']._serialized_start=4014
  _globals['_SNOOZESERVICEINSIGHTREQ']._serialized_end=4203
  _globals['_SNOOZESERVICEINSIGHTRESP']._serialized_start=4205
  _globals['_SNOOZESERVICEINSIGHTRESP']._serialized_end=4231
  _globals['_GETSERVICEMETADATAREQ']._serialized_start=4233
  _globals['_GETSERVICEMETADATAREQ']._serialized_end=4312
  _globals['_GETSERVICEMETADATARESP']._serialized_start=4314
  _globals['_GETSERVICEMETADATARESP']._serialized_end=4395
  _globals['_SETSERVICEMETADATAREQ']._serialized_start=4398
  _globals['_SETSERVICEMETADATAREQ']._serialized_end=4544
  _globals['_SETSERVICEMETADATARESP']._serialized_start=4546
  _globals['_SETSERVICEMETADATARESP']._serialized_end=4570
  _globals['_SETSERVICECONVERGENCEMODEREQ']._serialized_start=4573
  _globals['_SETSERVICECONVERGENCEMODEREQ']._serialized_end=4724
  _globals['_SETSERVICECONVERGENCEMODERESP']._serialized_start=4726
  _globals['_SETSERVICECONVERGENCEMODERESP']._serialized_end=4757
  _globals['_LISTMATERIALIZEDCONFIGVERSIONSREQ']._serialized_start=4760
  _globals['_LISTMATERIALIZEDCONFIGVERSIONSREQ']._serialized_end=4890
  _globals['_LISTMATERIALIZEDCONFIGVERSIONSRESP']._serialized_start=4893
  _globals['_LISTMATERIALIZEDCONFIGVERSIONSRESP']._serialized_end=5393
  _globals['_LISTMATERIALIZEDCONFIGVERSIONSRESP_VERSIONMETADATA']._serialized_start=5045
  _globals['_LISTMATERIALIZEDCONFIGVERSIONSRESP_VERSIONMETADATA']._serialized_end=5393
  _globals['_CONFIGURESERVICEREQ']._serialized_start=5396
  _globals['_CONFIGURESERVICEREQ']._serialized_end=5678
  _globals['_CONFIGURESERVICERESP']._serialized_start=5680
  _globals['_CONFIGURESERVICERESP']._serialized_end=5746
  _globals['_VALIDATECONFIGURESERVICERESP']._serialized_start=5749
  _globals['_VALIDATECONFIGURESERVICERESP']._serialized_end=5892
  _globals['_LISTSERVICECONFIGVERSIONSREQ']._serialized_start=5894
  _globals['_LISTSERVICECONFIGVERSIONSREQ']._serialized_end=6019
  _globals['_LISTSERVICECONFIGVERSIONSRESP']._serialized_start=6022
  _globals['_LISTSERVICECONFIGVERSIONSRESP']._serialized_end=6355
  _globals['_LISTSERVICECONFIGVERSIONSRESP_VERSIONMETADATA']._serialized_start=6164
  _globals['_LISTSERVICECONFIGVERSIONSRESP_VERSIONMETADATA']._serialized_end=6355
  _globals['_GETSERVICECONFIGREQ']._serialized_start=6357
  _globals['_GETSERVICECONFIGREQ']._serialized_end=6458
  _globals['_GETSERVICECONFIGRESP']._serialized_start=6461
  _globals['_GETSERVICECONFIGRESP']._serialized_end=6767
  _globals['_SERVICEMANAGER']._serialized_start=6770
  _globals['_SERVICEMANAGER']._serialized_end=10484
# @@protoc_insertion_point(module_scope)
