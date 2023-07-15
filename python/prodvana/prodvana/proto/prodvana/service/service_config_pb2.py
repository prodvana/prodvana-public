# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/service/service_config.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from prodvana.proto.prodvana.capability import capability_pb2 as prodvana_dot_capability_dot_capability__pb2
from prodvana.proto.prodvana.common_config import constants_pb2 as prodvana_dot_common__config_dot_constants__pb2
from prodvana.proto.prodvana.common_config import env_pb2 as prodvana_dot_common__config_dot_env__pb2
from prodvana.proto.prodvana.common_config import kubernetes_config_pb2 as prodvana_dot_common__config_dot_kubernetes__config__pb2
from prodvana.proto.prodvana.common_config import helm_pb2 as prodvana_dot_common__config_dot_helm__pb2
from prodvana.proto.prodvana.common_config import maturity_pb2 as prodvana_dot_common__config_dot_maturity__pb2
from prodvana.proto.prodvana.common_config import parameters_pb2 as prodvana_dot_common__config_dot_parameters__pb2
from prodvana.proto.prodvana.common_config import program_pb2 as prodvana_dot_common__config_dot_program__pb2
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n%prodvana/service/service_config.proto\x12\x10prodvana.service\x1a\x1egoogle/protobuf/duration.proto\x1a$prodvana/capability/capability.proto\x1a&prodvana/common_config/constants.proto\x1a prodvana/common_config/env.proto\x1a.prodvana/common_config/kubernetes_config.proto\x1a!prodvana/common_config/helm.proto\x1a%prodvana/common_config/maturity.proto\x1a\'prodvana/common_config/parameters.proto\x1a$prodvana/common_config/program.proto\x1a\"prodvana/common_config/retry.proto\x1a!prodvana/common_config/task.proto\x1a\'prodvana/delivery/delivery_config.proto\x1a(prodvana/delivery_extension/config.proto\x1a.prodvana/protection/protection_reference.proto\x1a\x35prodvana/release_channel/release_channel_config.proto\x1a\'prodvana/runtimes/runtimes_config.proto\x1a*prodvana/workflow/integration_config.proto\x1a!prodvana/service/parameters.proto\x1a\x1eprodvana/volumes/volumes.proto\x1a\x17validate/validate.proto\"1\n\x0eReplicasConfig\x12\x0f\n\x05\x66ixed\x18\x01 \x01(\x05H\x00\x42\x0e\n\x0c\x63onfig_oneof\"\xaa\x02\n\x0eMetricAnalysis\x12J\n\x0csuccess_rate\x18\x02 \x01(\x0b\x32\x32.prodvana.service.MetricAnalysis.SuccessRateConfigH\x00\x12\x45\n\x0blatency_p95\x18\x03 \x01(\x0b\x32..prodvana.service.MetricAnalysis.LatencyConfigH\x00\x1a\x32\n\x11SuccessRateConfig\x12\x1d\n\x15min_threshold_percent\x18\x01 \x01(\x01\x1a?\n\rLatencyConfig\x12.\n\x0bmax_latency\x18\x01 \x01(\x0b\x32\x19.google.protobuf.DurationB\x10\n\x0e\x61nalysis_oneof\"\xb7\x02\n\x15ReleaseStrategyConfig\x12<\n\x19individual_stage_deadline\x18\x01 \x01(\x0b\x32\x19.google.protobuf.Duration\x12=\n\x1a\x61utomated_testing_duration\x18\x02 \x01(\x0b\x32\x19.google.protobuf.Duration\x12:\n\x10metrics_analysis\x18\x03 \x03(\x0b\x32 .prodvana.service.MetricAnalysis\x12\x17\n\x0fmanual_approval\x18\x04 \x01(\x08\x12\x31\n\x0e\x63heck_interval\x18\x05 \x01(\x0b\x32\x19.google.protobuf.Duration\x12\x19\n\x11\x66\x61ilure_threshold\x18\x06 \x01(\x05\"f\n\tTLSSecret\x12\x14\n\nraw_secret\x18\x03 \x01(\tH\x00\x12\x30\n\x06secret\x18\x04 \x01(\x0b\x32\x1e.prodvana.common_config.SecretH\x00\x42\x11\n\ntls_secret\x12\x03\xf8\x42\x01\"\x81\x01\n\x0eTLSCertificate\x12\x37\n\x08tls_cert\x18\x01 \x01(\x0b\x32\x1b.prodvana.service.TLSSecretB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12\x36\n\x07tls_key\x18\x02 \x01(\x0b\x32\x1b.prodvana.service.TLSSecretB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\"j\n\x0b\x43\x65rtificate\x12/\n\x03tls\x18\x01 \x01(\x0b\x32 .prodvana.service.TLSCertificateH\x00\x12\x16\n\x0c\x61ws_acm_cert\x18\x02 \x01(\tH\x00\x42\x12\n\x0b\x63\x65rtificate\x12\x03\xf8\x42\x01\"\xb1\t\n\x17PerReleaseChannelConfig\x12 \n\x0frelease_channel\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12&\n\x10\x63ustom_hostnames\x18\x02 \x03(\tB\x0c\xfa\x42\t\x92\x01\x06\"\x04r\x02h\x01\x12W\n\x08programs\x18\x05 \x03(\x0b\x32\x36.prodvana.common_config.PerReleaseChannelProgramConfigB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\x12+\n\x04\x63\x65rt\x18\x06 \x01(\x0b\x32\x1d.prodvana.service.Certificate\x12:\n\x0f\x64\x65livery_config\x18\x07 \x01(\x0b\x32!.prodvana.delivery.DeliveryConfig\x12\x38\n\x07volumes\x18\x08 \x03(\x0b\x32\x18.prodvana.volumes.VolumeB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\x12\x32\n\x08replicas\x18\t \x01(\x0b\x32 .prodvana.service.ReplicasConfig\x12\x43\n\x0epre_push_tasks\x18\n \x03(\x0b\x32\x1c.prodvana.service.TaskConfigB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\x12U\n\x13\x64\x65livery_extensions\x18\x12 \x03(\x0b\x32).prodvana.service.DeliveryExtensionConfigB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\x12\x41\n\x10runtime_specific\x18\x0b \x01(\x0b\x32\'.prodvana.service.RuntimeSpecificConfig\x12\x1a\n\x12runtime_connection\x18\x11 \x01(\t\x12\x45\n\x11runtime_extension\x18\r \x01(\x0b\x32(.prodvana.service.RuntimeExtensionConfigH\x00\x12\x45\n\x11kubernetes_config\x18\x0e \x01(\x0b\x32(.prodvana.common_config.KubernetesConfigH\x00\x12\x43\n\x0f\x65xternal_config\x18\x10 \x01(\x0b\x32(.prodvana.common_config.KubernetesConfigH\x00\x12\x32\n\x04helm\x18\x0f \x01(\x0b\x32\".prodvana.common_config.HelmConfigH\x00\x12h\n\x03\x65nv\x18\x13 \x03(\x0b\x32\x32.prodvana.service.PerReleaseChannelConfig.EnvEntryB\'\xfa\x42$\x9a\x01!\x18\x01\"\x1dr\x1b\x32\x19^[a-zA-Z_]+[a-zA-Z0-9_]*$\x12\x33\n\tconstants\x18\x14 \x03(\x0b\x32 .prodvana.common_config.Constant\x1aL\n\x08\x45nvEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12/\n\x05value\x18\x02 \x01(\x0b\x32 .prodvana.common_config.EnvValue:\x02\x38\x01\x42\x0e\n\x0c\x63onfig_oneofJ\x04\x08\x03\x10\x04J\x04\x08\x04\x10\x05J\x04\x08\x0c\x10\rR\x0bprotections\",\n\x13\x43\x61pabilityReference\x12\x15\n\x04name\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\"U\n\x18\x43ompiledCapabilityConfig\x12\x39\n\ncapability\x18\x01 \x01(\x0b\x32%.prodvana.capability.CapabilityConfig\".\n\x10ProgramReference\x12\x0e\n\x04name\x18\x01 \x01(\tH\x00\x42\n\n\x03ref\x12\x03\xf8\x42\x01\"6\n\rTaskReference\x12\x19\n\x0frelease_channel\x18\x01 \x01(\tH\x00\x42\n\n\x03ref\x12\x03\xf8\x42\x01\"\xf1\x01\n\nTaskConfig\x12@\n\x07program\x18\x01 \x01(\x0b\x32%.prodvana.common_config.ProgramConfigB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12\x38\n\x0c\x62\x61se_program\x18\x02 \x01(\x0b\x32\".prodvana.service.ProgramReference\x12\x39\n\x0cretry_config\x18\x03 \x01(\x0b\x32#.prodvana.common_config.RetryConfig\x12,\n\x03ref\x18\x04 \x01(\x0b\x32\x1f.prodvana.service.TaskReference\"p\n\x0eProtectionLink\x12;\n\tlifecycle\x18\x03 \x01(\x0b\x32(.prodvana.protection.ProtectionLifecycle\x12\x15\n\rattachment_id\x18\x04 \x01(\tJ\x04\x08\x01\x10\x02J\x04\x08\x02\x10\x03\"\xa2\x02\n\x17\x44\x65liveryExtensionConfig\x12G\n\x07inlined\x18\x01 \x01(\x0b\x32\x34.prodvana.delivery_extension.DeliveryExtensionConfigH\x00\x12\x1b\n\x08instance\x18\x03 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01H\x00\x12H\n\x03ref\x18\x04 \x01(\x0b\x32\x39.prodvana.delivery_extension.DeliveryExtensionInstanceRefH\x00\x12\x44\n\tlifecycle\x18\x02 \x01(\x0e\x32%.prodvana.common_config.TaskLifecycleB\n\xfa\x42\x07\x82\x01\x04 \x00 \x01\x42\x11\n\ndefinition\x12\x03\xf8\x42\x01\"\x9e\x02\n\x19\x44\x65liveryExtensionInstance\x12G\n\x07inlined\x18\x01 \x01(\x0b\x32\x34.prodvana.delivery_extension.DeliveryExtensionConfigH\x00\x12H\n\x03ref\x18\x03 \x01(\x0b\x32\x39.prodvana.delivery_extension.DeliveryExtensionInstanceRefH\x00\x12\x15\n\x04name\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x44\n\tlifecycle\x18\x04 \x01(\x0e\x32%.prodvana.common_config.TaskLifecycleB\n\xfa\x42\x07\x82\x01\x04 \x00 \x01\x42\x11\n\ndefinition\x12\x03\xf8\x42\x01\"\x9c\x02\n\x15RuntimeSpecificConfig\x12@\n\x03k8s\x18\x01 \x01(\x0b\x32\x31.prodvana.service.RuntimeSpecificConfig.K8SConfigH\x00\x1a\xae\x01\n\tK8SConfig\x12\x66\n\x13service_annotations\x18\x01 \x03(\x0b\x32I.prodvana.service.RuntimeSpecificConfig.K8SConfig.ServiceAnnotationsEntry\x1a\x39\n\x17ServiceAnnotationsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x42\x10\n\x0eruntime_config\"Z\n\x16RuntimeExtensionConfig\x12@\n\x10parameter_values\x18\x01 \x03(\x0b\x32&.prodvana.common_config.ParameterValue\"&\n\x12\x41utoRollbackConfig\x12\x10\n\x08\x64isabled\x18\x01 \x01(\x08\"\xec\x01\n\x1fProtectionConvergenceAttachment\x12:\n\x04name\x18\x01 \x01(\tB,\xfa\x42)r\'\x10\x00\x18?2!^[a-z]?([a-z0-9-]*[a-z0-9]){0,1}$\x12?\n\x03ref\x18\x02 \x01(\x0b\x32(.prodvana.protection.ProtectionReferenceB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12L\n\tlifecycle\x18\x03 \x03(\x0b\x32(.prodvana.protection.ProtectionLifecycleB\x0f\xfa\x42\x0c\x92\x01\t\x08\x01\"\x05\x8a\x01\x02\x10\x01\"\xd3\x0f\n\rServiceConfig\x12\x39\n\x04name\x18\x01 \x01(\tB+\xfa\x42(r&\x10\x01\x18?2 ^[a-z]([a-z0-9-]*[a-z0-9]){0,1}$\x12\x13\n\x0b\x61pplication\x18\x14 \x01(\t\x12\x46\n\x08programs\x18\x02 \x03(\x0b\x32%.prodvana.common_config.ProgramConfigB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\x12\x32\n\x08replicas\x18\x04 \x01(\x0b\x32 .prodvana.service.ReplicasConfig\x12\x41\n\x10release_strategy\x18\x05 \x01(\x0b\x32\'.prodvana.service.ReleaseStrategyConfig\x12U\n\x13per_release_channel\x18\t \x03(\x0b\x32).prodvana.service.PerReleaseChannelConfigB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\x12J\n\x0c\x63\x61pabilities\x18\n \x03(\x0b\x32%.prodvana.service.CapabilityReferenceB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\x12:\n\x0f\x64\x65livery_config\x18\x0b \x01(\x0b\x32!.prodvana.delivery.DeliveryConfig\x12\x38\n\x07volumes\x18\x0c \x03(\x0b\x32\x18.prodvana.volumes.VolumeB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\x12@\n\x12\x64\x65ploy_annotations\x18\r \x01(\x0b\x32$.prodvana.workflow.AnnotationsConfig\x12\x34\n\x0epre_push_tasks\x18\x0f \x03(\x0b\x32\x1c.prodvana.service.TaskConfig\x12\x46\n\x13\x64\x65livery_extensions\x18\x1d \x03(\x0b\x32).prodvana.service.DeliveryExtensionConfig\x12Q\n\x1c\x64\x65livery_extension_instances\x18\x1e \x03(\x0b\x32+.prodvana.service.DeliveryExtensionInstance\x12\x41\n\x10runtime_specific\x18\x10 \x01(\x0b\x32\'.prodvana.service.RuntimeSpecificConfig\x12\x1a\n\x12runtime_connection\x18\x1c \x01(\t\x12N\n\nparameters\x18\x15 \x03(\x0b\x32+.prodvana.common_config.ParameterDefinitionB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\x12\x42\n\x10parameter_values\x18\x16 \x01(\x0b\x32(.prodvana.service.ServiceParameterValues\x12\x33\n\tconstants\x18! \x03(\x0b\x32 .prodvana.common_config.Constant\x12\x34\n\x11progress_deadline\x18\x18 \x01(\x0b\x32\x19.google.protobuf.Duration\x12\x45\n\x11runtime_extension\x18\x12 \x01(\x0b\x32(.prodvana.service.RuntimeExtensionConfigH\x00\x12\x45\n\x11kubernetes_config\x18\x13 \x01(\x0b\x32(.prodvana.common_config.KubernetesConfigH\x00\x12\x43\n\x0f\x65xternal_config\x18\x1b \x01(\x0b\x32(.prodvana.common_config.KubernetesConfigH\x00\x12\x32\n\x04helm\x18\x1a \x01(\x0b\x32\".prodvana.common_config.HelmConfigH\x00\x12M\n\x12parameters_autogen\x18\x17 \x01(\x0e\x32\x31.prodvana.service.ServiceConfig.ParametersAutogen\x12;\n\rauto_rollback\x18\x19 \x01(\x0b\x32$.prodvana.service.AutoRollbackConfig\x12\x1c\n\x14no_cleanup_on_delete\x18\x1f \x01(\x08\x12^\n\x03\x65nv\x18  \x03(\x0b\x32(.prodvana.service.ServiceConfig.EnvEntryB\'\xfa\x42$\x9a\x01!\x18\x01\"\x1dr\x1b\x32\x19^[a-zA-Z_]+[a-zA-Z0-9_]*$\x1aL\n\x08\x45nvEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12/\n\x05value\x18\x02 \x01(\x0b\x32 .prodvana.common_config.EnvValue:\x02\x38\x01\"M\n\x11ParametersAutogen\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x08\n\x04NONE\x10\x01\x12\t\n\x05IMAGE\x10\x02\x12\x16\n\x12IMAGE_AND_REPLICAS\x10\x03\x42\x0e\n\x0c\x63onfig_oneofJ\x04\x08\x06\x10\x07J\x04\x08\x07\x10\x08J\x04\x08\x08\x10\tJ\x04\x08\x0e\x10\x0fJ\x04\x08\x11\x10\x12R\x08\x65xternalR\x10image_repositoryR\rbase_templateR\x15\x63ontainer_registry_idR\x0bprotections\"\x89\x0e\n\x1d\x43ompiledServiceInstanceConfig\x12\x0f\n\x07service\x18\x01 \x01(\t\x12\x13\n\x0b\x61pplication\x18\x12 \x01(\t\x12\x17\n\x0frelease_channel\x18\x02 \x01(\t\x12\x37\n\x08programs\x18\x03 \x03(\x0b\x32%.prodvana.common_config.ProgramConfig\x12\x32\n\x08replicas\x18\x04 \x01(\x0b\x32 .prodvana.service.ReplicasConfig\x12\x32\n\x08maturity\x18\x05 \x01(\x0e\x32 .prodvana.common_config.Maturity\x12\x41\n\x10release_strategy\x18\x06 \x01(\x0b\x32\'.prodvana.service.ReleaseStrategyConfig\x12\x18\n\x10\x63ustom_hostnames\x18\x07 \x03(\t\x12+\n\x04\x63\x65rt\x18\n \x01(\x0b\x32\x1d.prodvana.service.Certificate\x12\x46\n\x07runtime\x18\x0b \x01(\x0b\x32\x35.prodvana.release_channel.ReleaseChannelRuntimeConfig\x12\x44\n\x11runtime_execution\x18\x1b \x01(\x0b\x32).prodvana.runtimes.RuntimeExecutionConfig\x12@\n\x0c\x63\x61pabilities\x18\x0c \x03(\x0b\x32*.prodvana.service.CompiledCapabilityConfig\x12:\n\x0f\x64\x65livery_config\x18\r \x01(\x0b\x32!.prodvana.delivery.DeliveryConfig\x12\x38\n\x07volumes\x18\x0e \x03(\x0b\x32\x18.prodvana.volumes.VolumeB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\x12@\n\x12\x64\x65ploy_annotations\x18\x0f \x01(\x0b\x32$.prodvana.workflow.AnnotationsConfig\x12\x34\n\x0epre_push_tasks\x18\x11 \x03(\x0b\x32\x1c.prodvana.service.TaskConfig\x12\x46\n\x13\x64\x65livery_extensions\x18\x1e \x03(\x0b\x32).prodvana.service.DeliveryExtensionConfig\x12Q\n\x1c\x64\x65livery_extension_instances\x18\x1f \x03(\x0b\x32+.prodvana.service.DeliveryExtensionInstance\x12\x41\n\x10runtime_specific\x18\x13 \x01(\x0b\x32\'.prodvana.service.RuntimeSpecificConfig\x12?\n\nparameters\x18\x18 \x03(\x0b\x32+.prodvana.common_config.ParameterDefinition\x12@\n\x10parameter_values\x18\x19 \x03(\x0b\x32&.prodvana.common_config.ParameterValue\x12\x33\n\tconstants\x18! \x03(\x0b\x32 .prodvana.common_config.Constant\x12\x34\n\x11progress_deadline\x18\x1a \x01(\x0b\x32\x19.google.protobuf.Duration\x12\x45\n\x11runtime_extension\x18\x16 \x01(\x0b\x32(.prodvana.service.RuntimeExtensionConfigH\x00\x12\x45\n\x11kubernetes_config\x18\x17 \x01(\x0b\x32(.prodvana.common_config.KubernetesConfigH\x00\x12\x32\n\x04helm\x18\x1d \x01(\x0b\x32\".prodvana.common_config.HelmConfigH\x00\x12n\n\x03\x65nv\x18\x1c \x03(\x0b\x32\x38.prodvana.service.CompiledServiceInstanceConfig.EnvEntryB\'\xfa\x42$\x9a\x01!\x18\x01\"\x1dr\x1b\x32\x19^[a-zA-Z_]+[a-zA-Z0-9_]*$\x12\x1c\n\x14no_cleanup_on_delete\x18  \x01(\x08\x1aL\n\x08\x45nvEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12/\n\x05value\x18\x02 \x01(\x0b\x32 .prodvana.common_config.EnvValue:\x02\x38\x01\x42\x0e\n\x0c\x63onfig_oneofJ\x04\x08\x08\x10\tJ\x04\x08\t\x10\nJ\x04\x08\x10\x10\x11J\x04\x08\x14\x10\x15J\x04\x08\x15\x10\x16R\x0bprotectionsR\rbase_templateR\x0cruntime_type\"\x88\x02\n\x11\x43ompiledJobConfig\x12\x13\n\x0bname_prefix\x18\x01 \x01(\t\x12\x17\n\x0frelease_channel\x18\x02 \x01(\t\x12\x37\n\x08programs\x18\x03 \x03(\x0b\x32%.prodvana.common_config.ProgramConfig\x12\x46\n\x07runtime\x18\x04 \x01(\x0b\x32\x35.prodvana.release_channel.ReleaseChannelRuntimeConfig\x12\x44\n\x11runtime_execution\x18\x05 \x01(\x0b\x32).prodvana.runtimes.RuntimeExecutionConfigBLZJgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/serviceb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'prodvana.service.service_config_pb2', _globals)
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
  _PERRELEASECHANNELCONFIG_ENVENTRY._options = None
  _PERRELEASECHANNELCONFIG_ENVENTRY._serialized_options = b'8\001'
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
  _PERRELEASECHANNELCONFIG.fields_by_name['env']._options = None
  _PERRELEASECHANNELCONFIG.fields_by_name['env']._serialized_options = b'\372B$\232\001!\030\001\"\035r\0332\031^[a-zA-Z_]+[a-zA-Z0-9_]*$'
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
  _DELIVERYEXTENSIONCONFIG.fields_by_name['instance']._options = None
  _DELIVERYEXTENSIONCONFIG.fields_by_name['instance']._serialized_options = b'\372B\004r\002\020\001'
  _DELIVERYEXTENSIONCONFIG.fields_by_name['lifecycle']._options = None
  _DELIVERYEXTENSIONCONFIG.fields_by_name['lifecycle']._serialized_options = b'\372B\007\202\001\004 \000 \001'
  _DELIVERYEXTENSIONINSTANCE.oneofs_by_name['definition']._options = None
  _DELIVERYEXTENSIONINSTANCE.oneofs_by_name['definition']._serialized_options = b'\370B\001'
  _DELIVERYEXTENSIONINSTANCE.fields_by_name['name']._options = None
  _DELIVERYEXTENSIONINSTANCE.fields_by_name['name']._serialized_options = b'\372B\004r\002\020\001'
  _DELIVERYEXTENSIONINSTANCE.fields_by_name['lifecycle']._options = None
  _DELIVERYEXTENSIONINSTANCE.fields_by_name['lifecycle']._serialized_options = b'\372B\007\202\001\004 \000 \001'
  _RUNTIMESPECIFICCONFIG_K8SCONFIG_SERVICEANNOTATIONSENTRY._options = None
  _RUNTIMESPECIFICCONFIG_K8SCONFIG_SERVICEANNOTATIONSENTRY._serialized_options = b'8\001'
  _PROTECTIONCONVERGENCEATTACHMENT.fields_by_name['name']._options = None
  _PROTECTIONCONVERGENCEATTACHMENT.fields_by_name['name']._serialized_options = b'\372B)r\'\020\000\030?2!^[a-z]?([a-z0-9-]*[a-z0-9]){0,1}$'
  _PROTECTIONCONVERGENCEATTACHMENT.fields_by_name['ref']._options = None
  _PROTECTIONCONVERGENCEATTACHMENT.fields_by_name['ref']._serialized_options = b'\372B\005\212\001\002\020\001'
  _PROTECTIONCONVERGENCEATTACHMENT.fields_by_name['lifecycle']._options = None
  _PROTECTIONCONVERGENCEATTACHMENT.fields_by_name['lifecycle']._serialized_options = b'\372B\014\222\001\t\010\001\"\005\212\001\002\020\001'
  _SERVICECONFIG_ENVENTRY._options = None
  _SERVICECONFIG_ENVENTRY._serialized_options = b'8\001'
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
  _SERVICECONFIG.fields_by_name['env']._options = None
  _SERVICECONFIG.fields_by_name['env']._serialized_options = b'\372B$\232\001!\030\001\"\035r\0332\031^[a-zA-Z_]+[a-zA-Z0-9_]*$'
  _COMPILEDSERVICEINSTANCECONFIG_ENVENTRY._options = None
  _COMPILEDSERVICEINSTANCECONFIG_ENVENTRY._serialized_options = b'8\001'
  _COMPILEDSERVICEINSTANCECONFIG.fields_by_name['volumes']._options = None
  _COMPILEDSERVICEINSTANCECONFIG.fields_by_name['volumes']._serialized_options = b'\372B\n\222\001\007\"\005\212\001\002\020\001'
  _COMPILEDSERVICEINSTANCECONFIG.fields_by_name['env']._options = None
  _COMPILEDSERVICEINSTANCECONFIG.fields_by_name['env']._serialized_options = b'\372B$\232\001!\030\001\"\035r\0332\031^[a-zA-Z_]+[a-zA-Z0-9_]*$'
  _globals['_REPLICASCONFIG']._serialized_start=838
  _globals['_REPLICASCONFIG']._serialized_end=887
  _globals['_METRICANALYSIS']._serialized_start=890
  _globals['_METRICANALYSIS']._serialized_end=1188
  _globals['_METRICANALYSIS_SUCCESSRATECONFIG']._serialized_start=1055
  _globals['_METRICANALYSIS_SUCCESSRATECONFIG']._serialized_end=1105
  _globals['_METRICANALYSIS_LATENCYCONFIG']._serialized_start=1107
  _globals['_METRICANALYSIS_LATENCYCONFIG']._serialized_end=1170
  _globals['_RELEASESTRATEGYCONFIG']._serialized_start=1191
  _globals['_RELEASESTRATEGYCONFIG']._serialized_end=1502
  _globals['_TLSSECRET']._serialized_start=1504
  _globals['_TLSSECRET']._serialized_end=1606
  _globals['_TLSCERTIFICATE']._serialized_start=1609
  _globals['_TLSCERTIFICATE']._serialized_end=1738
  _globals['_CERTIFICATE']._serialized_start=1740
  _globals['_CERTIFICATE']._serialized_end=1846
  _globals['_PERRELEASECHANNELCONFIG']._serialized_start=1849
  _globals['_PERRELEASECHANNELCONFIG']._serialized_end=3050
  _globals['_PERRELEASECHANNELCONFIG_ENVENTRY']._serialized_start=2927
  _globals['_PERRELEASECHANNELCONFIG_ENVENTRY']._serialized_end=3003
  _globals['_CAPABILITYREFERENCE']._serialized_start=3052
  _globals['_CAPABILITYREFERENCE']._serialized_end=3096
  _globals['_COMPILEDCAPABILITYCONFIG']._serialized_start=3098
  _globals['_COMPILEDCAPABILITYCONFIG']._serialized_end=3183
  _globals['_PROGRAMREFERENCE']._serialized_start=3185
  _globals['_PROGRAMREFERENCE']._serialized_end=3231
  _globals['_TASKREFERENCE']._serialized_start=3233
  _globals['_TASKREFERENCE']._serialized_end=3287
  _globals['_TASKCONFIG']._serialized_start=3290
  _globals['_TASKCONFIG']._serialized_end=3531
  _globals['_PROTECTIONLINK']._serialized_start=3533
  _globals['_PROTECTIONLINK']._serialized_end=3645
  _globals['_DELIVERYEXTENSIONCONFIG']._serialized_start=3648
  _globals['_DELIVERYEXTENSIONCONFIG']._serialized_end=3938
  _globals['_DELIVERYEXTENSIONINSTANCE']._serialized_start=3941
  _globals['_DELIVERYEXTENSIONINSTANCE']._serialized_end=4227
  _globals['_RUNTIMESPECIFICCONFIG']._serialized_start=4230
  _globals['_RUNTIMESPECIFICCONFIG']._serialized_end=4514
  _globals['_RUNTIMESPECIFICCONFIG_K8SCONFIG']._serialized_start=4322
  _globals['_RUNTIMESPECIFICCONFIG_K8SCONFIG']._serialized_end=4496
  _globals['_RUNTIMESPECIFICCONFIG_K8SCONFIG_SERVICEANNOTATIONSENTRY']._serialized_start=4439
  _globals['_RUNTIMESPECIFICCONFIG_K8SCONFIG_SERVICEANNOTATIONSENTRY']._serialized_end=4496
  _globals['_RUNTIMEEXTENSIONCONFIG']._serialized_start=4516
  _globals['_RUNTIMEEXTENSIONCONFIG']._serialized_end=4606
  _globals['_AUTOROLLBACKCONFIG']._serialized_start=4608
  _globals['_AUTOROLLBACKCONFIG']._serialized_end=4646
  _globals['_PROTECTIONCONVERGENCEATTACHMENT']._serialized_start=4649
  _globals['_PROTECTIONCONVERGENCEATTACHMENT']._serialized_end=4885
  _globals['_SERVICECONFIG']._serialized_start=4888
  _globals['_SERVICECONFIG']._serialized_end=6891
  _globals['_SERVICECONFIG_ENVENTRY']._serialized_start=2927
  _globals['_SERVICECONFIG_ENVENTRY']._serialized_end=3003
  _globals['_SERVICECONFIG_PARAMETERSAUTOGEN']._serialized_start=6689
  _globals['_SERVICECONFIG_PARAMETERSAUTOGEN']._serialized_end=6766
  _globals['_COMPILEDSERVICEINSTANCECONFIG']._serialized_start=6894
  _globals['_COMPILEDSERVICEINSTANCECONFIG']._serialized_end=8695
  _globals['_COMPILEDSERVICEINSTANCECONFIG_ENVENTRY']._serialized_start=2927
  _globals['_COMPILEDSERVICEINSTANCECONFIG_ENVENTRY']._serialized_end=3003
  _globals['_COMPILEDJOBCONFIG']._serialized_start=8698
  _globals['_COMPILEDJOBCONFIG']._serialized_end=8962
# @@protoc_insertion_point(module_scope)
