# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/service/service_config.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from prodvana.proto.prodvana.capability import capability_pb2 as prodvana_dot_capability_dot_capability__pb2
from prodvana.proto.prodvana.common_config import env_pb2 as prodvana_dot_common__config_dot_env__pb2
from prodvana.proto.prodvana.common_config import kubernetes_config_pb2 as prodvana_dot_common__config_dot_kubernetes__config__pb2
from prodvana.proto.prodvana.common_config import helm_pb2 as prodvana_dot_common__config_dot_helm__pb2
from prodvana.proto.prodvana.common_config import maturity_pb2 as prodvana_dot_common__config_dot_maturity__pb2
from prodvana.proto.prodvana.common_config import parameters_pb2 as prodvana_dot_common__config_dot_parameters__pb2
from prodvana.proto.prodvana.common_config import program_pb2 as prodvana_dot_common__config_dot_program__pb2
from prodvana.proto.prodvana.common_config import ref_pb2 as prodvana_dot_common__config_dot_ref__pb2
from prodvana.proto.prodvana.common_config import retry_pb2 as prodvana_dot_common__config_dot_retry__pb2
from prodvana.proto.prodvana.common_config import task_pb2 as prodvana_dot_common__config_dot_task__pb2
from prodvana.proto.prodvana.delivery import delivery_config_pb2 as prodvana_dot_delivery_dot_delivery__config__pb2
from prodvana.proto.prodvana.delivery_extension import config_pb2 as prodvana_dot_delivery__extension_dot_config__pb2
from prodvana.proto.prodvana.protection import protection_reference_pb2 as prodvana_dot_protection_dot_protection__reference__pb2
from prodvana.proto.prodvana.release_channel import release_channel_config_pb2 as prodvana_dot_release__channel_dot_release__channel__config__pb2
from prodvana.proto.prodvana.runtimes import runtimes_config_pb2 as prodvana_dot_runtimes_dot_runtimes__config__pb2
from prodvana.proto.prodvana.workflow import integration_config_pb2 as prodvana_dot_workflow_dot_integration__config__pb2
from prodvana.proto.prodvana.service import parameters_pb2 as prodvana_dot_service_dot_parameters__pb2
from prodvana.proto.prodvana.volumes import volumes_pb2 as prodvana_dot_volumes_dot_volumes__pb2
from prodvana.proto.validate import validate_pb2 as validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n%prodvana/service/service_config.proto\x12\x10prodvana.service\x1a\x1egoogle/protobuf/duration.proto\x1a$prodvana/capability/capability.proto\x1a prodvana/common_config/env.proto\x1a.prodvana/common_config/kubernetes_config.proto\x1a!prodvana/common_config/helm.proto\x1a%prodvana/common_config/maturity.proto\x1a\'prodvana/common_config/parameters.proto\x1a$prodvana/common_config/program.proto\x1a prodvana/common_config/ref.proto\x1a\"prodvana/common_config/retry.proto\x1a!prodvana/common_config/task.proto\x1a\'prodvana/delivery/delivery_config.proto\x1a(prodvana/delivery_extension/config.proto\x1a.prodvana/protection/protection_reference.proto\x1a\x35prodvana/release_channel/release_channel_config.proto\x1a\'prodvana/runtimes/runtimes_config.proto\x1a*prodvana/workflow/integration_config.proto\x1a!prodvana/service/parameters.proto\x1a\x1eprodvana/volumes/volumes.proto\x1a\x17validate/validate.proto\"1\n\x0eReplicasConfig\x12\x0f\n\x05\x66ixed\x18\x01 \x01(\x05H\x00\x42\x0e\n\x0c\x63onfig_oneof\"\xaa\x02\n\x0eMetricAnalysis\x12J\n\x0csuccess_rate\x18\x02 \x01(\x0b\x32\x32.prodvana.service.MetricAnalysis.SuccessRateConfigH\x00\x12\x45\n\x0blatency_p95\x18\x03 \x01(\x0b\x32..prodvana.service.MetricAnalysis.LatencyConfigH\x00\x1a\x32\n\x11SuccessRateConfig\x12\x1d\n\x15min_threshold_percent\x18\x01 \x01(\x01\x1a?\n\rLatencyConfig\x12.\n\x0bmax_latency\x18\x01 \x01(\x0b\x32\x19.google.protobuf.DurationB\x10\n\x0e\x61nalysis_oneof\"\xb7\x02\n\x15ReleaseStrategyConfig\x12<\n\x19individual_stage_deadline\x18\x01 \x01(\x0b\x32\x19.google.protobuf.Duration\x12=\n\x1a\x61utomated_testing_duration\x18\x02 \x01(\x0b\x32\x19.google.protobuf.Duration\x12:\n\x10metrics_analysis\x18\x03 \x03(\x0b\x32 .prodvana.service.MetricAnalysis\x12\x17\n\x0fmanual_approval\x18\x04 \x01(\x08\x12\x31\n\x0e\x63heck_interval\x18\x05 \x01(\x0b\x32\x19.google.protobuf.Duration\x12\x19\n\x11\x66\x61ilure_threshold\x18\x06 \x01(\x05\"f\n\tTLSSecret\x12\x14\n\nraw_secret\x18\x03 \x01(\tH\x00\x12\x30\n\x06secret\x18\x04 \x01(\x0b\x32\x1e.prodvana.common_config.SecretH\x00\x42\x11\n\ntls_secret\x12\x03\xf8\x42\x01\"\x81\x01\n\x0eTLSCertificate\x12\x37\n\x08tls_cert\x18\x01 \x01(\x0b\x32\x1b.prodvana.service.TLSSecretB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12\x36\n\x07tls_key\x18\x02 \x01(\x0b\x32\x1b.prodvana.service.TLSSecretB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\"j\n\x0b\x43\x65rtificate\x12/\n\x03tls\x18\x01 \x01(\x0b\x32 .prodvana.service.TLSCertificateH\x00\x12\x16\n\x0c\x61ws_acm_cert\x18\x02 \x01(\tH\x00\x42\x12\n\x0b\x63\x65rtificate\x12\x03\xf8\x42\x01\"\xc4\x07\n\x17PerReleaseChannelConfig\x12 \n\x0frelease_channel\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12&\n\x10\x63ustom_hostnames\x18\x02 \x03(\tB\x0c\xfa\x42\t\x92\x01\x06\"\x04r\x02h\x01\x12W\n\x08programs\x18\x05 \x03(\x0b\x32\x36.prodvana.common_config.PerReleaseChannelProgramConfigB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\x12+\n\x04\x63\x65rt\x18\x06 \x01(\x0b\x32\x1d.prodvana.service.Certificate\x12:\n\x0f\x64\x65livery_config\x18\x07 \x01(\x0b\x32!.prodvana.delivery.DeliveryConfig\x12\x38\n\x07volumes\x18\x08 \x03(\x0b\x32\x18.prodvana.volumes.VolumeB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\x12\x32\n\x08replicas\x18\t \x01(\x0b\x32 .prodvana.service.ReplicasConfig\x12\x43\n\x0epre_push_tasks\x18\n \x03(\x0b\x32\x1c.prodvana.service.TaskConfigB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\x12U\n\x13\x64\x65livery_extensions\x18\x12 \x03(\x0b\x32).prodvana.service.DeliveryExtensionConfigB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\x12\x41\n\x10runtime_specific\x18\x0b \x01(\x0b\x32\'.prodvana.service.RuntimeSpecificConfig\x12\x1a\n\x12runtime_connection\x18\x11 \x01(\t\x12\x45\n\x11runtime_extension\x18\r \x01(\x0b\x32(.prodvana.service.RuntimeExtensionConfigH\x00\x12\x45\n\x11kubernetes_config\x18\x0e \x01(\x0b\x32(.prodvana.common_config.KubernetesConfigH\x00\x12\x43\n\x0f\x65xternal_config\x18\x10 \x01(\x0b\x32(.prodvana.common_config.KubernetesConfigH\x00\x12\x32\n\x04helm\x18\x0f \x01(\x0b\x32\".prodvana.common_config.HelmConfigH\x00\x42\x0e\n\x0c\x63onfig_oneofJ\x04\x08\x03\x10\x04J\x04\x08\x04\x10\x05J\x04\x08\x0c\x10\rR\x0bprotections\",\n\x13\x43\x61pabilityReference\x12\x15\n\x04name\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\"U\n\x18\x43ompiledCapabilityConfig\x12\x39\n\ncapability\x18\x01 \x01(\x0b\x32%.prodvana.capability.CapabilityConfig\".\n\x10ProgramReference\x12\x0e\n\x04name\x18\x01 \x01(\tH\x00\x42\n\n\x03ref\x12\x03\xf8\x42\x01\"6\n\rTaskReference\x12\x19\n\x0frelease_channel\x18\x01 \x01(\tH\x00\x42\n\n\x03ref\x12\x03\xf8\x42\x01\"\xf1\x01\n\nTaskConfig\x12@\n\x07program\x18\x01 \x01(\x0b\x32%.prodvana.common_config.ProgramConfigB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12\x38\n\x0c\x62\x61se_program\x18\x02 \x01(\x0b\x32\".prodvana.service.ProgramReference\x12\x39\n\x0cretry_config\x18\x03 \x01(\x0b\x32#.prodvana.common_config.RetryConfig\x12,\n\x03ref\x18\x04 \x01(\x0b\x32\x1f.prodvana.service.TaskReference\"p\n\x0eProtectionLink\x12;\n\tlifecycle\x18\x03 \x01(\x0b\x32(.prodvana.protection.ProtectionLifecycle\x12\x15\n\rattachment_id\x18\x04 \x01(\tJ\x04\x08\x01\x10\x02J\x04\x08\x02\x10\x03\"\xd3\x01\n\x17\x44\x65liveryExtensionConfig\x12G\n\x07inlined\x18\x01 \x01(\x0b\x32\x34.prodvana.delivery_extension.DeliveryExtensionConfigH\x00\x12\x16\n\x03ref\x18\x03 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01H\x00\x12\x44\n\tlifecycle\x18\x02 \x01(\x0e\x32%.prodvana.common_config.TaskLifecycleB\n\xfa\x42\x07\x82\x01\x04 \x00 \x01\x42\x11\n\ndefinition\x12\x03\xf8\x42\x01\"x\n\x19\x44\x65liveryExtensionInstance\x12\x44\n\x06\x63onfig\x18\x01 \x01(\x0b\x32\x34.prodvana.delivery_extension.DeliveryExtensionConfig\x12\x15\n\x04name\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\"\x9c\x02\n\x15RuntimeSpecificConfig\x12@\n\x03k8s\x18\x01 \x01(\x0b\x32\x31.prodvana.service.RuntimeSpecificConfig.K8SConfigH\x00\x1a\xae\x01\n\tK8SConfig\x12\x66\n\x13service_annotations\x18\x01 \x03(\x0b\x32I.prodvana.service.RuntimeSpecificConfig.K8SConfig.ServiceAnnotationsEntry\x1a\x39\n\x17ServiceAnnotationsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x42\x10\n\x0eruntime_config\"Z\n\x16RuntimeExtensionConfig\x12@\n\x10parameter_values\x18\x01 \x03(\x0b\x32&.prodvana.common_config.ParameterValue\"&\n\x12\x41utoRollbackConfig\x12\x10\n\x08\x64isabled\x18\x01 \x01(\x08\"\xec\x01\n\x1fProtectionConvergenceAttachment\x12:\n\x04name\x18\x01 \x01(\tB,\xfa\x42)r\'\x10\x00\x18?2!^[a-z]?([a-z0-9-]*[a-z0-9]){0,1}$\x12?\n\x03ref\x18\x02 \x01(\x0b\x32(.prodvana.protection.ProtectionReferenceB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12L\n\tlifecycle\x18\x03 \x03(\x0b\x32(.prodvana.protection.ProtectionLifecycleB\x0f\xfa\x42\x0c\x92\x01\t\x08\x01\"\x05\x8a\x01\x02\x10\x01\"\x9e\x0e\n\rServiceConfig\x12\x39\n\x04name\x18\x01 \x01(\tB+\xfa\x42(r&\x10\x01\x18?2 ^[a-z]([a-z0-9-]*[a-z0-9]){0,1}$\x12\x13\n\x0b\x61pplication\x18\x14 \x01(\t\x12\x46\n\x08programs\x18\x02 \x03(\x0b\x32%.prodvana.common_config.ProgramConfigB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\x12\x32\n\x08replicas\x18\x04 \x01(\x0b\x32 .prodvana.service.ReplicasConfig\x12\x41\n\x10release_strategy\x18\x05 \x01(\x0b\x32\'.prodvana.service.ReleaseStrategyConfig\x12U\n\x13per_release_channel\x18\t \x03(\x0b\x32).prodvana.service.PerReleaseChannelConfigB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\x12J\n\x0c\x63\x61pabilities\x18\n \x03(\x0b\x32%.prodvana.service.CapabilityReferenceB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\x12:\n\x0f\x64\x65livery_config\x18\x0b \x01(\x0b\x32!.prodvana.delivery.DeliveryConfig\x12\x38\n\x07volumes\x18\x0c \x03(\x0b\x32\x18.prodvana.volumes.VolumeB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\x12@\n\x12\x64\x65ploy_annotations\x18\r \x01(\x0b\x32$.prodvana.workflow.AnnotationsConfig\x12\x41\n\rbase_template\x18\x0e \x01(\x0b\x32*.prodvana.common_config.ServiceTemplateRef\x12\x34\n\x0epre_push_tasks\x18\x0f \x03(\x0b\x32\x1c.prodvana.service.TaskConfig\x12\x46\n\x13\x64\x65livery_extensions\x18\x1d \x03(\x0b\x32).prodvana.service.DeliveryExtensionConfig\x12Q\n\x1c\x64\x65livery_extension_instances\x18\x1e \x03(\x0b\x32+.prodvana.service.DeliveryExtensionInstance\x12\x41\n\x10runtime_specific\x18\x10 \x01(\x0b\x32\'.prodvana.service.RuntimeSpecificConfig\x12\x1a\n\x12runtime_connection\x18\x1c \x01(\t\x12N\n\nparameters\x18\x15 \x03(\x0b\x32+.prodvana.common_config.ParameterDefinitionB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\x12\x42\n\x10parameter_values\x18\x16 \x01(\x0b\x32(.prodvana.service.ServiceParameterValues\x12\x34\n\x11progress_deadline\x18\x18 \x01(\x0b\x32\x19.google.protobuf.Duration\x12\x45\n\x11runtime_extension\x18\x12 \x01(\x0b\x32(.prodvana.service.RuntimeExtensionConfigH\x00\x12\x45\n\x11kubernetes_config\x18\x13 \x01(\x0b\x32(.prodvana.common_config.KubernetesConfigH\x00\x12\x43\n\x0f\x65xternal_config\x18\x1b \x01(\x0b\x32(.prodvana.common_config.KubernetesConfigH\x00\x12\x32\n\x04helm\x18\x1a \x01(\x0b\x32\".prodvana.common_config.HelmConfigH\x00\x12M\n\x12parameters_autogen\x18\x17 \x01(\x0e\x32\x31.prodvana.service.ServiceConfig.ParametersAutogen\x12;\n\rauto_rollback\x18\x19 \x01(\x0b\x32$.prodvana.service.AutoRollbackConfig\x12\x1c\n\x14no_cleanup_on_delete\x18\x1f \x01(\x08\"M\n\x11ParametersAutogen\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x08\n\x04NONE\x10\x01\x12\t\n\x05IMAGE\x10\x02\x12\x16\n\x12IMAGE_AND_REPLICAS\x10\x03\x42\x0e\n\x0c\x63onfig_oneofJ\x04\x08\x06\x10\x07J\x04\x08\x07\x10\x08J\x04\x08\x08\x10\tJ\x04\x08\x11\x10\x12R\x08\x65xternalR\x10image_repositoryR\x15\x63ontainer_registry_idR\x0bprotections\"\x82\x0e\n\x1d\x43ompiledServiceInstanceConfig\x12\x0f\n\x07service\x18\x01 \x01(\t\x12\x13\n\x0b\x61pplication\x18\x12 \x01(\t\x12\x17\n\x0frelease_channel\x18\x02 \x01(\t\x12\x37\n\x08programs\x18\x03 \x03(\x0b\x32%.prodvana.common_config.ProgramConfig\x12\x32\n\x08replicas\x18\x04 \x01(\x0b\x32 .prodvana.service.ReplicasConfig\x12\x32\n\x08maturity\x18\x05 \x01(\x0e\x32 .prodvana.common_config.Maturity\x12\x41\n\x10release_strategy\x18\x06 \x01(\x0b\x32\'.prodvana.service.ReleaseStrategyConfig\x12\x18\n\x10\x63ustom_hostnames\x18\x07 \x03(\t\x12+\n\x04\x63\x65rt\x18\n \x01(\x0b\x32\x1d.prodvana.service.Certificate\x12\x46\n\x07runtime\x18\x0b \x01(\x0b\x32\x35.prodvana.release_channel.ReleaseChannelRuntimeConfig\x12\x44\n\x11runtime_execution\x18\x1b \x01(\x0b\x32).prodvana.runtimes.RuntimeExecutionConfig\x12@\n\x0c\x63\x61pabilities\x18\x0c \x03(\x0b\x32*.prodvana.service.CompiledCapabilityConfig\x12:\n\x0f\x64\x65livery_config\x18\r \x01(\x0b\x32!.prodvana.delivery.DeliveryConfig\x12\x38\n\x07volumes\x18\x0e \x03(\x0b\x32\x18.prodvana.volumes.VolumeB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\x12@\n\x12\x64\x65ploy_annotations\x18\x0f \x01(\x0b\x32$.prodvana.workflow.AnnotationsConfig\x12\x41\n\rbase_template\x18\x10 \x01(\x0b\x32*.prodvana.common_config.ServiceTemplateRef\x12\x34\n\x0epre_push_tasks\x18\x11 \x03(\x0b\x32\x1c.prodvana.service.TaskConfig\x12\x46\n\x13\x64\x65livery_extensions\x18\x1e \x03(\x0b\x32).prodvana.service.DeliveryExtensionConfig\x12Q\n\x1c\x64\x65livery_extension_instances\x18\x1f \x03(\x0b\x32+.prodvana.service.DeliveryExtensionInstance\x12\x41\n\x10runtime_specific\x18\x13 \x01(\x0b\x32\'.prodvana.service.RuntimeSpecificConfig\x12?\n\nparameters\x18\x18 \x03(\x0b\x32+.prodvana.common_config.ParameterDefinition\x12@\n\x10parameter_values\x18\x19 \x03(\x0b\x32&.prodvana.common_config.ParameterValue\x12\x34\n\x11progress_deadline\x18\x1a \x01(\x0b\x32\x19.google.protobuf.Duration\x12\x45\n\x11runtime_extension\x18\x16 \x01(\x0b\x32(.prodvana.service.RuntimeExtensionConfigH\x00\x12\x45\n\x11kubernetes_config\x18\x17 \x01(\x0b\x32(.prodvana.common_config.KubernetesConfigH\x00\x12\x32\n\x04helm\x18\x1d \x01(\x0b\x32\".prodvana.common_config.HelmConfigH\x00\x12n\n\x03\x65nv\x18\x1c \x03(\x0b\x32\x38.prodvana.service.CompiledServiceInstanceConfig.EnvEntryB\'\xfa\x42$\x9a\x01!\x18\x01\"\x1dr\x1b\x32\x19^[a-zA-Z_]+[a-zA-Z0-9_]*$\x12\x1c\n\x14no_cleanup_on_delete\x18  \x01(\x08\x1aL\n\x08\x45nvEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12/\n\x05value\x18\x02 \x01(\x0b\x32 .prodvana.common_config.EnvValue:\x02\x38\x01\x42\x0e\n\x0c\x63onfig_oneofJ\x04\x08\x08\x10\tJ\x04\x08\t\x10\nJ\x04\x08\x14\x10\x15J\x04\x08\x15\x10\x16R\x0bprotectionsR\x0cruntime_type\"\x88\x02\n\x11\x43ompiledJobConfig\x12\x13\n\x0bname_prefix\x18\x01 \x01(\t\x12\x17\n\x0frelease_channel\x18\x02 \x01(\t\x12\x37\n\x08programs\x18\x03 \x03(\x0b\x32%.prodvana.common_config.ProgramConfig\x12\x46\n\x07runtime\x18\x04 \x01(\x0b\x32\x35.prodvana.release_channel.ReleaseChannelRuntimeConfig\x12\x44\n\x11runtime_execution\x18\x05 \x01(\x0b\x32).prodvana.runtimes.RuntimeExecutionConfigBLZJgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/serviceb\x06proto3')



_REPLICASCONFIG = DESCRIPTOR.message_types_by_name['ReplicasConfig']
_METRICANALYSIS = DESCRIPTOR.message_types_by_name['MetricAnalysis']
_METRICANALYSIS_SUCCESSRATECONFIG = _METRICANALYSIS.nested_types_by_name['SuccessRateConfig']
_METRICANALYSIS_LATENCYCONFIG = _METRICANALYSIS.nested_types_by_name['LatencyConfig']
_RELEASESTRATEGYCONFIG = DESCRIPTOR.message_types_by_name['ReleaseStrategyConfig']
_TLSSECRET = DESCRIPTOR.message_types_by_name['TLSSecret']
_TLSCERTIFICATE = DESCRIPTOR.message_types_by_name['TLSCertificate']
_CERTIFICATE = DESCRIPTOR.message_types_by_name['Certificate']
_PERRELEASECHANNELCONFIG = DESCRIPTOR.message_types_by_name['PerReleaseChannelConfig']
_CAPABILITYREFERENCE = DESCRIPTOR.message_types_by_name['CapabilityReference']
_COMPILEDCAPABILITYCONFIG = DESCRIPTOR.message_types_by_name['CompiledCapabilityConfig']
_PROGRAMREFERENCE = DESCRIPTOR.message_types_by_name['ProgramReference']
_TASKREFERENCE = DESCRIPTOR.message_types_by_name['TaskReference']
_TASKCONFIG = DESCRIPTOR.message_types_by_name['TaskConfig']
_PROTECTIONLINK = DESCRIPTOR.message_types_by_name['ProtectionLink']
_DELIVERYEXTENSIONCONFIG = DESCRIPTOR.message_types_by_name['DeliveryExtensionConfig']
_DELIVERYEXTENSIONINSTANCE = DESCRIPTOR.message_types_by_name['DeliveryExtensionInstance']
_RUNTIMESPECIFICCONFIG = DESCRIPTOR.message_types_by_name['RuntimeSpecificConfig']
_RUNTIMESPECIFICCONFIG_K8SCONFIG = _RUNTIMESPECIFICCONFIG.nested_types_by_name['K8SConfig']
_RUNTIMESPECIFICCONFIG_K8SCONFIG_SERVICEANNOTATIONSENTRY = _RUNTIMESPECIFICCONFIG_K8SCONFIG.nested_types_by_name['ServiceAnnotationsEntry']
_RUNTIMEEXTENSIONCONFIG = DESCRIPTOR.message_types_by_name['RuntimeExtensionConfig']
_AUTOROLLBACKCONFIG = DESCRIPTOR.message_types_by_name['AutoRollbackConfig']
_PROTECTIONCONVERGENCEATTACHMENT = DESCRIPTOR.message_types_by_name['ProtectionConvergenceAttachment']
_SERVICECONFIG = DESCRIPTOR.message_types_by_name['ServiceConfig']
_COMPILEDSERVICEINSTANCECONFIG = DESCRIPTOR.message_types_by_name['CompiledServiceInstanceConfig']
_COMPILEDSERVICEINSTANCECONFIG_ENVENTRY = _COMPILEDSERVICEINSTANCECONFIG.nested_types_by_name['EnvEntry']
_COMPILEDJOBCONFIG = DESCRIPTOR.message_types_by_name['CompiledJobConfig']
_SERVICECONFIG_PARAMETERSAUTOGEN = _SERVICECONFIG.enum_types_by_name['ParametersAutogen']
ReplicasConfig = _reflection.GeneratedProtocolMessageType('ReplicasConfig', (_message.Message,), {
  'DESCRIPTOR' : _REPLICASCONFIG,
  '__module__' : 'prodvana.service.service_config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.service.ReplicasConfig)
  })
_sym_db.RegisterMessage(ReplicasConfig)

MetricAnalysis = _reflection.GeneratedProtocolMessageType('MetricAnalysis', (_message.Message,), {

  'SuccessRateConfig' : _reflection.GeneratedProtocolMessageType('SuccessRateConfig', (_message.Message,), {
    'DESCRIPTOR' : _METRICANALYSIS_SUCCESSRATECONFIG,
    '__module__' : 'prodvana.service.service_config_pb2'
    # @@protoc_insertion_point(class_scope:prodvana.service.MetricAnalysis.SuccessRateConfig)
    })
  ,

  'LatencyConfig' : _reflection.GeneratedProtocolMessageType('LatencyConfig', (_message.Message,), {
    'DESCRIPTOR' : _METRICANALYSIS_LATENCYCONFIG,
    '__module__' : 'prodvana.service.service_config_pb2'
    # @@protoc_insertion_point(class_scope:prodvana.service.MetricAnalysis.LatencyConfig)
    })
  ,
  'DESCRIPTOR' : _METRICANALYSIS,
  '__module__' : 'prodvana.service.service_config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.service.MetricAnalysis)
  })
_sym_db.RegisterMessage(MetricAnalysis)
_sym_db.RegisterMessage(MetricAnalysis.SuccessRateConfig)
_sym_db.RegisterMessage(MetricAnalysis.LatencyConfig)

ReleaseStrategyConfig = _reflection.GeneratedProtocolMessageType('ReleaseStrategyConfig', (_message.Message,), {
  'DESCRIPTOR' : _RELEASESTRATEGYCONFIG,
  '__module__' : 'prodvana.service.service_config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.service.ReleaseStrategyConfig)
  })
_sym_db.RegisterMessage(ReleaseStrategyConfig)

TLSSecret = _reflection.GeneratedProtocolMessageType('TLSSecret', (_message.Message,), {
  'DESCRIPTOR' : _TLSSECRET,
  '__module__' : 'prodvana.service.service_config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.service.TLSSecret)
  })
_sym_db.RegisterMessage(TLSSecret)

TLSCertificate = _reflection.GeneratedProtocolMessageType('TLSCertificate', (_message.Message,), {
  'DESCRIPTOR' : _TLSCERTIFICATE,
  '__module__' : 'prodvana.service.service_config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.service.TLSCertificate)
  })
_sym_db.RegisterMessage(TLSCertificate)

Certificate = _reflection.GeneratedProtocolMessageType('Certificate', (_message.Message,), {
  'DESCRIPTOR' : _CERTIFICATE,
  '__module__' : 'prodvana.service.service_config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.service.Certificate)
  })
_sym_db.RegisterMessage(Certificate)

PerReleaseChannelConfig = _reflection.GeneratedProtocolMessageType('PerReleaseChannelConfig', (_message.Message,), {
  'DESCRIPTOR' : _PERRELEASECHANNELCONFIG,
  '__module__' : 'prodvana.service.service_config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.service.PerReleaseChannelConfig)
  })
_sym_db.RegisterMessage(PerReleaseChannelConfig)

CapabilityReference = _reflection.GeneratedProtocolMessageType('CapabilityReference', (_message.Message,), {
  'DESCRIPTOR' : _CAPABILITYREFERENCE,
  '__module__' : 'prodvana.service.service_config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.service.CapabilityReference)
  })
_sym_db.RegisterMessage(CapabilityReference)

CompiledCapabilityConfig = _reflection.GeneratedProtocolMessageType('CompiledCapabilityConfig', (_message.Message,), {
  'DESCRIPTOR' : _COMPILEDCAPABILITYCONFIG,
  '__module__' : 'prodvana.service.service_config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.service.CompiledCapabilityConfig)
  })
_sym_db.RegisterMessage(CompiledCapabilityConfig)

ProgramReference = _reflection.GeneratedProtocolMessageType('ProgramReference', (_message.Message,), {
  'DESCRIPTOR' : _PROGRAMREFERENCE,
  '__module__' : 'prodvana.service.service_config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.service.ProgramReference)
  })
_sym_db.RegisterMessage(ProgramReference)

TaskReference = _reflection.GeneratedProtocolMessageType('TaskReference', (_message.Message,), {
  'DESCRIPTOR' : _TASKREFERENCE,
  '__module__' : 'prodvana.service.service_config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.service.TaskReference)
  })
_sym_db.RegisterMessage(TaskReference)

TaskConfig = _reflection.GeneratedProtocolMessageType('TaskConfig', (_message.Message,), {
  'DESCRIPTOR' : _TASKCONFIG,
  '__module__' : 'prodvana.service.service_config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.service.TaskConfig)
  })
_sym_db.RegisterMessage(TaskConfig)

ProtectionLink = _reflection.GeneratedProtocolMessageType('ProtectionLink', (_message.Message,), {
  'DESCRIPTOR' : _PROTECTIONLINK,
  '__module__' : 'prodvana.service.service_config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.service.ProtectionLink)
  })
_sym_db.RegisterMessage(ProtectionLink)

DeliveryExtensionConfig = _reflection.GeneratedProtocolMessageType('DeliveryExtensionConfig', (_message.Message,), {
  'DESCRIPTOR' : _DELIVERYEXTENSIONCONFIG,
  '__module__' : 'prodvana.service.service_config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.service.DeliveryExtensionConfig)
  })
_sym_db.RegisterMessage(DeliveryExtensionConfig)

DeliveryExtensionInstance = _reflection.GeneratedProtocolMessageType('DeliveryExtensionInstance', (_message.Message,), {
  'DESCRIPTOR' : _DELIVERYEXTENSIONINSTANCE,
  '__module__' : 'prodvana.service.service_config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.service.DeliveryExtensionInstance)
  })
_sym_db.RegisterMessage(DeliveryExtensionInstance)

RuntimeSpecificConfig = _reflection.GeneratedProtocolMessageType('RuntimeSpecificConfig', (_message.Message,), {

  'K8SConfig' : _reflection.GeneratedProtocolMessageType('K8SConfig', (_message.Message,), {

    'ServiceAnnotationsEntry' : _reflection.GeneratedProtocolMessageType('ServiceAnnotationsEntry', (_message.Message,), {
      'DESCRIPTOR' : _RUNTIMESPECIFICCONFIG_K8SCONFIG_SERVICEANNOTATIONSENTRY,
      '__module__' : 'prodvana.service.service_config_pb2'
      # @@protoc_insertion_point(class_scope:prodvana.service.RuntimeSpecificConfig.K8SConfig.ServiceAnnotationsEntry)
      })
    ,
    'DESCRIPTOR' : _RUNTIMESPECIFICCONFIG_K8SCONFIG,
    '__module__' : 'prodvana.service.service_config_pb2'
    # @@protoc_insertion_point(class_scope:prodvana.service.RuntimeSpecificConfig.K8SConfig)
    })
  ,
  'DESCRIPTOR' : _RUNTIMESPECIFICCONFIG,
  '__module__' : 'prodvana.service.service_config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.service.RuntimeSpecificConfig)
  })
_sym_db.RegisterMessage(RuntimeSpecificConfig)
_sym_db.RegisterMessage(RuntimeSpecificConfig.K8SConfig)
_sym_db.RegisterMessage(RuntimeSpecificConfig.K8SConfig.ServiceAnnotationsEntry)

RuntimeExtensionConfig = _reflection.GeneratedProtocolMessageType('RuntimeExtensionConfig', (_message.Message,), {
  'DESCRIPTOR' : _RUNTIMEEXTENSIONCONFIG,
  '__module__' : 'prodvana.service.service_config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.service.RuntimeExtensionConfig)
  })
_sym_db.RegisterMessage(RuntimeExtensionConfig)

AutoRollbackConfig = _reflection.GeneratedProtocolMessageType('AutoRollbackConfig', (_message.Message,), {
  'DESCRIPTOR' : _AUTOROLLBACKCONFIG,
  '__module__' : 'prodvana.service.service_config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.service.AutoRollbackConfig)
  })
_sym_db.RegisterMessage(AutoRollbackConfig)

ProtectionConvergenceAttachment = _reflection.GeneratedProtocolMessageType('ProtectionConvergenceAttachment', (_message.Message,), {
  'DESCRIPTOR' : _PROTECTIONCONVERGENCEATTACHMENT,
  '__module__' : 'prodvana.service.service_config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.service.ProtectionConvergenceAttachment)
  })
_sym_db.RegisterMessage(ProtectionConvergenceAttachment)

ServiceConfig = _reflection.GeneratedProtocolMessageType('ServiceConfig', (_message.Message,), {
  'DESCRIPTOR' : _SERVICECONFIG,
  '__module__' : 'prodvana.service.service_config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.service.ServiceConfig)
  })
_sym_db.RegisterMessage(ServiceConfig)

CompiledServiceInstanceConfig = _reflection.GeneratedProtocolMessageType('CompiledServiceInstanceConfig', (_message.Message,), {

  'EnvEntry' : _reflection.GeneratedProtocolMessageType('EnvEntry', (_message.Message,), {
    'DESCRIPTOR' : _COMPILEDSERVICEINSTANCECONFIG_ENVENTRY,
    '__module__' : 'prodvana.service.service_config_pb2'
    # @@protoc_insertion_point(class_scope:prodvana.service.CompiledServiceInstanceConfig.EnvEntry)
    })
  ,
  'DESCRIPTOR' : _COMPILEDSERVICEINSTANCECONFIG,
  '__module__' : 'prodvana.service.service_config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.service.CompiledServiceInstanceConfig)
  })
_sym_db.RegisterMessage(CompiledServiceInstanceConfig)
_sym_db.RegisterMessage(CompiledServiceInstanceConfig.EnvEntry)

CompiledJobConfig = _reflection.GeneratedProtocolMessageType('CompiledJobConfig', (_message.Message,), {
  'DESCRIPTOR' : _COMPILEDJOBCONFIG,
  '__module__' : 'prodvana.service.service_config_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.service.CompiledJobConfig)
  })
_sym_db.RegisterMessage(CompiledJobConfig)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZJgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/service'
  _TLSSECRET.oneofs_by_name['tls_secret']._options = None
  _TLSSECRET.oneofs_by_name['tls_secret']._serialized_options = b'\370B\001'
  _TLSCERTIFICATE.fields_by_name['tls_cert']._options = None
  _TLSCERTIFICATE.fields_by_name['tls_cert']._serialized_options = b'\372B\005\212\001\002\020\001'
  _TLSCERTIFICATE.fields_by_name['tls_key']._options = None
  _TLSCERTIFICATE.fields_by_name['tls_key']._serialized_options = b'\372B\005\212\001\002\020\001'
  _CERTIFICATE.oneofs_by_name['certificate']._options = None
  _CERTIFICATE.oneofs_by_name['certificate']._serialized_options = b'\370B\001'
  _PERRELEASECHANNELCONFIG.fields_by_name['release_channel']._options = None
  _PERRELEASECHANNELCONFIG.fields_by_name['release_channel']._serialized_options = b'\372B\004r\002\020\001'
  _PERRELEASECHANNELCONFIG.fields_by_name['custom_hostnames']._options = None
  _PERRELEASECHANNELCONFIG.fields_by_name['custom_hostnames']._serialized_options = b'\372B\t\222\001\006\"\004r\002h\001'
  _PERRELEASECHANNELCONFIG.fields_by_name['programs']._options = None
  _PERRELEASECHANNELCONFIG.fields_by_name['programs']._serialized_options = b'\372B\n\222\001\007\"\005\212\001\002\020\001'
  _PERRELEASECHANNELCONFIG.fields_by_name['volumes']._options = None
  _PERRELEASECHANNELCONFIG.fields_by_name['volumes']._serialized_options = b'\372B\n\222\001\007\"\005\212\001\002\020\001'
  _PERRELEASECHANNELCONFIG.fields_by_name['pre_push_tasks']._options = None
  _PERRELEASECHANNELCONFIG.fields_by_name['pre_push_tasks']._serialized_options = b'\372B\n\222\001\007\"\005\212\001\002\020\001'
  _PERRELEASECHANNELCONFIG.fields_by_name['delivery_extensions']._options = None
  _PERRELEASECHANNELCONFIG.fields_by_name['delivery_extensions']._serialized_options = b'\372B\n\222\001\007\"\005\212\001\002\020\001'
  _CAPABILITYREFERENCE.fields_by_name['name']._options = None
  _CAPABILITYREFERENCE.fields_by_name['name']._serialized_options = b'\372B\004r\002\020\001'
  _PROGRAMREFERENCE.oneofs_by_name['ref']._options = None
  _PROGRAMREFERENCE.oneofs_by_name['ref']._serialized_options = b'\370B\001'
  _TASKREFERENCE.oneofs_by_name['ref']._options = None
  _TASKREFERENCE.oneofs_by_name['ref']._serialized_options = b'\370B\001'
  _TASKCONFIG.fields_by_name['program']._options = None
  _TASKCONFIG.fields_by_name['program']._serialized_options = b'\372B\005\212\001\002\020\001'
  _DELIVERYEXTENSIONCONFIG.oneofs_by_name['definition']._options = None
  _DELIVERYEXTENSIONCONFIG.oneofs_by_name['definition']._serialized_options = b'\370B\001'
  _DELIVERYEXTENSIONCONFIG.fields_by_name['ref']._options = None
  _DELIVERYEXTENSIONCONFIG.fields_by_name['ref']._serialized_options = b'\372B\004r\002\020\001'
  _DELIVERYEXTENSIONCONFIG.fields_by_name['lifecycle']._options = None
  _DELIVERYEXTENSIONCONFIG.fields_by_name['lifecycle']._serialized_options = b'\372B\007\202\001\004 \000 \001'
  _DELIVERYEXTENSIONINSTANCE.fields_by_name['name']._options = None
  _DELIVERYEXTENSIONINSTANCE.fields_by_name['name']._serialized_options = b'\372B\004r\002\020\001'
  _RUNTIMESPECIFICCONFIG_K8SCONFIG_SERVICEANNOTATIONSENTRY._options = None
  _RUNTIMESPECIFICCONFIG_K8SCONFIG_SERVICEANNOTATIONSENTRY._serialized_options = b'8\001'
  _PROTECTIONCONVERGENCEATTACHMENT.fields_by_name['name']._options = None
  _PROTECTIONCONVERGENCEATTACHMENT.fields_by_name['name']._serialized_options = b'\372B)r\'\020\000\030?2!^[a-z]?([a-z0-9-]*[a-z0-9]){0,1}$'
  _PROTECTIONCONVERGENCEATTACHMENT.fields_by_name['ref']._options = None
  _PROTECTIONCONVERGENCEATTACHMENT.fields_by_name['ref']._serialized_options = b'\372B\005\212\001\002\020\001'
  _PROTECTIONCONVERGENCEATTACHMENT.fields_by_name['lifecycle']._options = None
  _PROTECTIONCONVERGENCEATTACHMENT.fields_by_name['lifecycle']._serialized_options = b'\372B\014\222\001\t\010\001\"\005\212\001\002\020\001'
  _SERVICECONFIG.fields_by_name['name']._options = None
  _SERVICECONFIG.fields_by_name['name']._serialized_options = b'\372B(r&\020\001\030?2 ^[a-z]([a-z0-9-]*[a-z0-9]){0,1}$'
  _SERVICECONFIG.fields_by_name['programs']._options = None
  _SERVICECONFIG.fields_by_name['programs']._serialized_options = b'\372B\n\222\001\007\"\005\212\001\002\020\001'
  _SERVICECONFIG.fields_by_name['per_release_channel']._options = None
  _SERVICECONFIG.fields_by_name['per_release_channel']._serialized_options = b'\372B\n\222\001\007\"\005\212\001\002\020\001'
  _SERVICECONFIG.fields_by_name['capabilities']._options = None
  _SERVICECONFIG.fields_by_name['capabilities']._serialized_options = b'\372B\n\222\001\007\"\005\212\001\002\020\001'
  _SERVICECONFIG.fields_by_name['volumes']._options = None
  _SERVICECONFIG.fields_by_name['volumes']._serialized_options = b'\372B\n\222\001\007\"\005\212\001\002\020\001'
  _SERVICECONFIG.fields_by_name['parameters']._options = None
  _SERVICECONFIG.fields_by_name['parameters']._serialized_options = b'\372B\n\222\001\007\"\005\212\001\002\020\001'
  _COMPILEDSERVICEINSTANCECONFIG_ENVENTRY._options = None
  _COMPILEDSERVICEINSTANCECONFIG_ENVENTRY._serialized_options = b'8\001'
  _COMPILEDSERVICEINSTANCECONFIG.fields_by_name['volumes']._options = None
  _COMPILEDSERVICEINSTANCECONFIG.fields_by_name['volumes']._serialized_options = b'\372B\n\222\001\007\"\005\212\001\002\020\001'
  _COMPILEDSERVICEINSTANCECONFIG.fields_by_name['env']._options = None
  _COMPILEDSERVICEINSTANCECONFIG.fields_by_name['env']._serialized_options = b'\372B$\232\001!\030\001\"\035r\0332\031^[a-zA-Z_]+[a-zA-Z0-9_]*$'
  _REPLICASCONFIG._serialized_start=832
  _REPLICASCONFIG._serialized_end=881
  _METRICANALYSIS._serialized_start=884
  _METRICANALYSIS._serialized_end=1182
  _METRICANALYSIS_SUCCESSRATECONFIG._serialized_start=1049
  _METRICANALYSIS_SUCCESSRATECONFIG._serialized_end=1099
  _METRICANALYSIS_LATENCYCONFIG._serialized_start=1101
  _METRICANALYSIS_LATENCYCONFIG._serialized_end=1164
  _RELEASESTRATEGYCONFIG._serialized_start=1185
  _RELEASESTRATEGYCONFIG._serialized_end=1496
  _TLSSECRET._serialized_start=1498
  _TLSSECRET._serialized_end=1600
  _TLSCERTIFICATE._serialized_start=1603
  _TLSCERTIFICATE._serialized_end=1732
  _CERTIFICATE._serialized_start=1734
  _CERTIFICATE._serialized_end=1840
  _PERRELEASECHANNELCONFIG._serialized_start=1843
  _PERRELEASECHANNELCONFIG._serialized_end=2807
  _CAPABILITYREFERENCE._serialized_start=2809
  _CAPABILITYREFERENCE._serialized_end=2853
  _COMPILEDCAPABILITYCONFIG._serialized_start=2855
  _COMPILEDCAPABILITYCONFIG._serialized_end=2940
  _PROGRAMREFERENCE._serialized_start=2942
  _PROGRAMREFERENCE._serialized_end=2988
  _TASKREFERENCE._serialized_start=2990
  _TASKREFERENCE._serialized_end=3044
  _TASKCONFIG._serialized_start=3047
  _TASKCONFIG._serialized_end=3288
  _PROTECTIONLINK._serialized_start=3290
  _PROTECTIONLINK._serialized_end=3402
  _DELIVERYEXTENSIONCONFIG._serialized_start=3405
  _DELIVERYEXTENSIONCONFIG._serialized_end=3616
  _DELIVERYEXTENSIONINSTANCE._serialized_start=3618
  _DELIVERYEXTENSIONINSTANCE._serialized_end=3738
  _RUNTIMESPECIFICCONFIG._serialized_start=3741
  _RUNTIMESPECIFICCONFIG._serialized_end=4025
  _RUNTIMESPECIFICCONFIG_K8SCONFIG._serialized_start=3833
  _RUNTIMESPECIFICCONFIG_K8SCONFIG._serialized_end=4007
  _RUNTIMESPECIFICCONFIG_K8SCONFIG_SERVICEANNOTATIONSENTRY._serialized_start=3950
  _RUNTIMESPECIFICCONFIG_K8SCONFIG_SERVICEANNOTATIONSENTRY._serialized_end=4007
  _RUNTIMEEXTENSIONCONFIG._serialized_start=4027
  _RUNTIMEEXTENSIONCONFIG._serialized_end=4117
  _AUTOROLLBACKCONFIG._serialized_start=4119
  _AUTOROLLBACKCONFIG._serialized_end=4157
  _PROTECTIONCONVERGENCEATTACHMENT._serialized_start=4160
  _PROTECTIONCONVERGENCEATTACHMENT._serialized_end=4396
  _SERVICECONFIG._serialized_start=4399
  _SERVICECONFIG._serialized_end=6221
  _SERVICECONFIG_PARAMETERSAUTOGEN._serialized_start=6040
  _SERVICECONFIG_PARAMETERSAUTOGEN._serialized_end=6117
  _COMPILEDSERVICEINSTANCECONFIG._serialized_start=6224
  _COMPILEDSERVICEINSTANCECONFIG._serialized_end=8018
  _COMPILEDSERVICEINSTANCECONFIG_ENVENTRY._serialized_start=7875
  _COMPILEDSERVICEINSTANCECONFIG_ENVENTRY._serialized_end=7951
  _COMPILEDJOBCONFIG._serialized_start=8021
  _COMPILEDJOBCONFIG._serialized_end=8285
# @@protoc_insertion_point(module_scope)
