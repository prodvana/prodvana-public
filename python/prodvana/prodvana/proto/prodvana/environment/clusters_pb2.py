# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/environment/clusters.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from prodvana.proto.validate import validate_pb2 as validate_dot_validate__pb2
from prodvana.proto.prodvana.common_config import env_pb2 as prodvana_dot_common__config_dot_env__pb2
from prodvana.proto.prodvana.common_config import task_pb2 as prodvana_dot_common__config_dot_task__pb2
from prodvana.proto.prodvana.common_config import kubernetes_config_pb2 as prodvana_dot_common__config_dot_kubernetes__config__pb2
from prodvana.proto.prodvana.common_config import parameters_pb2 as prodvana_dot_common__config_dot_parameters__pb2
from prodvana.proto.prodvana.common_config import rollback_pb2 as prodvana_dot_common__config_dot_rollback__pb2
from prodvana.proto.prodvana.labels import labels_pb2 as prodvana_dot_labels_dot_labels__pb2
from prodvana.proto.prodvana.runtimes.extensions import fetch_pb2 as prodvana_dot_runtimes_dot_extensions_dot_fetch__pb2
from prodvana.proto.prodvana.runtimes import runtimes_config_pb2 as prodvana_dot_runtimes_dot_runtimes__config__pb2
from prodvana.proto.prodvana.volumes import volumes_pb2 as prodvana_dot_volumes_dot_volumes__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n#prodvana/environment/clusters.proto\x12\x14prodvana.environment\x1a\x1egoogle/protobuf/duration.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x17validate/validate.proto\x1a prodvana/common_config/env.proto\x1a!prodvana/common_config/task.proto\x1a.prodvana/common_config/kubernetes_config.proto\x1a\'prodvana/common_config/parameters.proto\x1a%prodvana/common_config/rollback.proto\x1a\x1cprodvana/labels/labels.proto\x1a(prodvana/runtimes/extensions/fetch.proto\x1a\'prodvana/runtimes/runtimes_config.proto\x1a\x1eprodvana/volumes/volumes.proto\"\x95\x04\n\x0b\x43lusterAuth\x12\x38\n\x03\x65\x63s\x18\x05 \x01(\x0b\x32).prodvana.environment.ClusterAuth.ECSAuthH\x00\x12\x38\n\x03k8s\x18\x08 \x01(\x0b\x32).prodvana.environment.ClusterAuth.K8sAuthH\x00\x1ao\n\x07\x45\x43SAuth\x12\x12\n\naccess_key\x18\x01 \x01(\t\x12\x12\n\nsecret_key\x18\x02 \x01(\t\x12\x0e\n\x06region\x18\x03 \x01(\t\x12\x17\n\x0f\x61ssume_role_arn\x18\x04 \x01(\t\x12\x13\n\x0b\x63luster_arn\x18\x05 \x01(\t\x1a\xa8\x01\n\x07K8sAuth\x12J\n\tagent_env\x18\x01 \x03(\x0b\x32\x37.prodvana.environment.ClusterAuth.K8sAuth.AgentEnvEntry\x12 \n\x18\x61gent_externally_managed\x18\x02 \x01(\x08\x1a/\n\rAgentEnvEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x42\x0c\n\nauth_oneofJ\x04\x08\x01\x10\x02J\x04\x08\x02\x10\x03J\x04\x08\x03\x10\x04J\x04\x08\x04\x10\x05J\x04\x08\x06\x10\x07J\x04\x08\x07\x10\x08R\textensionR\x08\x65ndpointR\x07\x63\x61_certR\x05tokenR\x0fservice_accountR\x0ek8s_agent_auth\"\x8e\x02\n\x07\x43luster\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x12\n\ncluster_id\x18\x02 \x01(\t\x12/\n\x04\x61uth\x18\x04 \x01(\x0b\x32!.prodvana.environment.ClusterAuth\x12/\n\x04type\x18\x05 \x01(\x0e\x32!.prodvana.environment.ClusterType\x12\x33\n\x06\x63onfig\x18\x06 \x01(\x0b\x32#.prodvana.environment.ClusterConfig\x12<\n\x18last_heartbeat_timestamp\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.TimestampJ\x04\x08\x03\x10\x04R\x06origin\"\xc6\x01\n\x11\x46\x61keClusterConfig\x12Z\n\x11\x63rashing_programs\x18\x01 \x03(\x0b\x32?.prodvana.environment.FakeClusterConfig.CrashingProgramPatterns\x1aU\n\x17\x43rashingProgramPatterns\x12\x13\n\x0bimage_regex\x18\x01 \x01(\t\x12\x11\n\tcmd_regex\x18\x02 \x01(\t\x12\x12\n\nlog_output\x18\x03 \x01(\t\"\x84\x01\n\x0bRetryPolicy\x12:\n\rbase_interval\x18\x01 \x01(\x0b\x32\x19.google.protobuf.DurationB\x08\xfa\x42\x05\xaa\x01\x02*\x00\x12\x39\n\x0cmax_interval\x18\x02 \x01(\x0b\x32\x19.google.protobuf.DurationB\x08\xfa\x42\x05\xaa\x01\x02*\x00\"\xfa\x04\n\x15\x45xtensionFetchCommand\x12\x39\n\x0btask_config\x18\x01 \x01(\x0b\x32\".prodvana.common_config.TaskConfigH\x00\x12\x45\n\x11kubernetes_config\x18\x02 \x01(\x0b\x32(.prodvana.common_config.KubernetesConfigH\x00\x12\x30\n\rpoll_interval\x18\x03 \x01(\x0b\x32\x19.google.protobuf.Duration\x12=\n\x1asteady_state_poll_interval\x18\x06 \x01(\x0b\x32\x19.google.protobuf.Duration\x12*\n\x07timeout\x18\x04 \x01(\x0b\x32\x19.google.protobuf.Duration\x12\x37\n\x0cretry_policy\x18\x07 \x01(\x0b\x32!.prodvana.environment.RetryPolicy\x12j\n\x03\x65nv\x18\x05 \x03(\x0b\x32\x34.prodvana.environment.ExtensionFetchCommand.EnvEntryB\'\xfa\x42$\x9a\x01!\x18\x01\"\x1dr\x1b\x32\x19^[a-zA-Z_]+[a-zA-Z0-9_]*$\x12;\n\nfetch_mode\x18\x08 \x01(\x0e\x32\'.prodvana.runtimes.extensions.FetchMode\x1aL\n\x08\x45nvEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12/\n\x05value\x18\x02 \x01(\x0b\x32 .prodvana.common_config.EnvValue:\x02\x38\x01\x42\x12\n\x0b\x65xec_config\x12\x03\xf8\x42\x01\"\x84\x04\n\x17\x45xtensionGetInfoCommand\x12\x39\n\x0btask_config\x18\x01 \x01(\x0b\x32\".prodvana.common_config.TaskConfigH\x00\x12\x45\n\x11kubernetes_config\x18\x02 \x01(\x0b\x32(.prodvana.common_config.KubernetesConfigH\x00\x12\x30\n\rpoll_interval\x18\x03 \x01(\x0b\x32\x19.google.protobuf.Duration\x12\x37\n\x0cretry_policy\x18\x04 \x01(\x0b\x32!.prodvana.environment.RetryPolicy\x12l\n\x03\x65nv\x18\x05 \x03(\x0b\x32\x36.prodvana.environment.ExtensionGetInfoCommand.EnvEntryB\'\xfa\x42$\x9a\x01!\x18\x01\"\x1dr\x1b\x32\x19^[a-zA-Z_]+[a-zA-Z0-9_]*$\x12,\n$test_only_do_not_require_pvn_wrapper\x18\x06 \x01(\x08\x1aL\n\x08\x45nvEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12/\n\x05value\x18\x02 \x01(\x0b\x32 .prodvana.common_config.EnvValue:\x02\x38\x01\x42\x12\n\x0b\x65xec_config\x12\x03\xf8\x42\x01\"\xf6\x03\n\x15\x45xtensionApplyCommand\x12\x39\n\x0btask_config\x18\x01 \x01(\x0b\x32\".prodvana.common_config.TaskConfigH\x00\x12\x45\n\x11kubernetes_config\x18\x02 \x01(\x0b\x32(.prodvana.common_config.KubernetesConfigH\x00\x12*\n\x07timeout\x18\x04 \x01(\x0b\x32\x19.google.protobuf.Duration\x12j\n\x03\x65nv\x18\x05 \x03(\x0b\x32\x34.prodvana.environment.ExtensionApplyCommand.EnvEntryB\'\xfa\x42$\x9a\x01!\x18\x01\"\x1dr\x1b\x32\x19^[a-zA-Z_]+[a-zA-Z0-9_]*$\x12\x37\n\x0cretry_policy\x18\x07 \x01(\x0b\x32!.prodvana.environment.RetryPolicy\x12\x1c\n\x14retryable_exit_codes\x18\x08 \x03(\x05\x1aL\n\x08\x45nvEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12/\n\x05value\x18\x02 \x01(\x0b\x32 .prodvana.common_config.EnvValue:\x02\x38\x01\x42\x12\n\x0b\x65xec_config\x12\x03\xf8\x42\x01J\x04\x08\x03\x10\x04J\x04\x08\x06\x10\x07\"\xd6\x04\n\x16\x45xtensionClusterConfig\x12\x44\n\x05\x61pply\x18\x01 \x01(\x0b\x32+.prodvana.environment.ExtensionApplyCommandB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12:\n\x05\x66\x65tch\x18\x02 \x01(\x0b\x32+.prodvana.environment.ExtensionFetchCommand\x12<\n\x05\x64\x65\x62ug\x18\x08 \x01(\x0b\x32-.prodvana.environment.ExtensionGetInfoCommand\x12?\n\x08get_info\x18\t \x01(\x0b\x32-.prodvana.environment.ExtensionGetInfoCommand\x12N\n\nparameters\x18\x03 \x03(\x0b\x32+.prodvana.common_config.ParameterDefinitionB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\x12J\n\rproxy_runtime\x18\x04 \x01(\x0b\x32).prodvana.runtimes.RuntimeExecutionConfigB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12\x31\n\x04type\x18\x05 \x01(\x0e\x32#.prodvana.environment.ExtensionType\x12%\n\x1drequire_approval_before_apply\x18\x06 \x01(\x08\x12\x45\n\x18\x63onvergence_grace_period\x18\x07 \x01(\x0b\x32\x19.google.protobuf.DurationB\x08\xfa\x42\x05\xaa\x01\x02*\x00\"\xbc\x04\n\x18\x43ompiledExtensionCommand\x12\x13\n\x0bname_prefix\x18\x01 \x01(\t\x12\\\n\x07\x63ommand\x18\x02 \x01(\x0b\x32K.prodvana.environment.CompiledExtensionCommand.CompiledExtensionCommandExec\x12\x44\n\x11runtime_execution\x18\x03 \x01(\x0b\x32).prodvana.runtimes.RuntimeExecutionConfig\x12\x44\n\x03\x65nv\x18\x04 \x03(\x0b\x32\x37.prodvana.environment.CompiledExtensionCommand.EnvEntry\x1a\xd2\x01\n\x1c\x43ompiledExtensionCommandExec\x12\x39\n\x0btask_config\x18\x01 \x01(\x0b\x32\".prodvana.common_config.TaskConfigH\x00\x12\x45\n\x11kubernetes_config\x18\x02 \x01(\x0b\x32(.prodvana.common_config.KubernetesConfigH\x00\x42\x12\n\x0b\x65xec_config\x12\x03\xf8\x42\x01J\x04\x08\x03\x10\x04J\x04\x08\x04\x10\x05J\x04\x08\x05\x10\x06J\x04\x08\x06\x10\x07J\x04\x08\x07\x10\x08\x1aL\n\x08\x45nvEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12/\n\x05value\x18\x02 \x01(\x0b\x32 .prodvana.common_config.EnvValue:\x02\x38\x01\"(\n\x10IacRunnerCommand\x12\x14\n\x03\x63md\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\"\xd6\x05\n\x15TerraformRunnerConfig\x12J\n\rproxy_runtime\x18\x01 \x01(\x0b\x32).prodvana.runtimes.RuntimeExecutionConfigB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12\x41\n\x03\x65nv\x18\x02 \x03(\x0b\x32\x34.prodvana.environment.TerraformRunnerConfig.EnvEntry\x12)\n\x07volumes\x18\x03 \x03(\x0b\x32\x18.prodvana.volumes.Volume\x12\x37\n\x07pre_run\x18\x04 \x03(\x0b\x32&.prodvana.environment.IacRunnerCommand\x12\x30\n\rpoll_interval\x18\x05 \x01(\x0b\x32\x19.google.protobuf.Duration\x12=\n\x1asteady_state_poll_interval\x18\x07 \x01(\x0b\x32\x19.google.protobuf.Duration\x12%\n\x1drequire_approval_before_apply\x18\x06 \x01(\x08\x12\x45\n\x18\x63onvergence_grace_period\x18\x08 \x01(\x0b\x32\x19.google.protobuf.DurationB\x08\xfa\x42\x05\xaa\x01\x02*\x00\x12=\n\x12\x66\x65tch_retry_policy\x18\t \x01(\x0b\x32!.prodvana.environment.RetryPolicy\x12=\n\x12\x61pply_retry_policy\x18\n \x01(\x0b\x32!.prodvana.environment.RetryPolicy\x12\x1f\n\x17\x64isable_drift_detection\x18\x0b \x01(\x08\x1aL\n\x08\x45nvEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12/\n\x05value\x18\x02 \x01(\x0b\x32 .prodvana.common_config.EnvValue:\x02\x38\x01\"Z\n\x12GKEClusterMetadata\x12\x14\n\x0cproject_name\x18\x01 \x01(\t\x12\x18\n\x10\x63luster_location\x18\x02 \x01(\t\x12\x14\n\x0c\x63luster_name\x18\x03 \x01(\t\"^\n\x0f\x43lusterMetadata\x12\x37\n\x03gke\x18\x01 \x01(\x0b\x32(.prodvana.environment.GKEClusterMetadataH\x00\x42\x12\n\x10\x63luster_metadata\"\xb2\x03\n\x0c\x41wsEcsConfig\x12J\n\rproxy_runtime\x18\x01 \x01(\x0b\x32).prodvana.runtimes.RuntimeExecutionConfigB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12L\n\naccess_key\x18\x06 \x01(\x0b\x32,.prodvana.environment.AwsEcsConfig.AccessKeyB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01H\x00\x12\x17\n\x06region\x18\x04 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x1c\n\x0b\x65\x63s_cluster\x18\x05 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x1a\x86\x01\n\tAccessKey\x12\"\n\x11\x61ws_access_key_id\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12U\n\x15\x61ws_secret_access_key\x18\x02 \x01(\x0b\x32,.prodvana.common_config.SecretReferenceValueB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x42\x12\n\x0b\x63redentials\x12\x03\xf8\x42\x01J\x04\x08\x02\x10\x03J\x04\x08\x03\x10\x04R\x11\x61ws_access_key_idR\x15\x61ws_secret_access_key\"\x81\x02\n\x14GoogleCloudRunConfig\x12J\n\rproxy_runtime\x18\x01 \x01(\x0b\x32).prodvana.runtimes.RuntimeExecutionConfigB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12V\n\x14service_account_json\x18\x02 \x01(\x0b\x32,.prodvana.common_config.SecretReferenceValueB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01H\x00\x12\x18\n\x07project\x18\x03 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x17\n\x06region\x18\x04 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x42\x12\n\x0b\x63redentials\x12\x03\xf8\x42\x01\"\x98\x11\n\rClusterConfig\x12\x0c\n\x04name\x18\x0e \x01(\t\x12>\n\x08kubecost\x18\x03 \x01(\x0b\x32,.prodvana.environment.ClusterConfig.Kubecost\x12<\n\x07\x64\x61tadog\x18\x04 \x01(\x0b\x32+.prodvana.environment.ClusterConfig.Datadog\x12\x43\n\x0b\x61lb_ingress\x18\x06 \x01(\x0b\x32..prodvana.environment.ClusterConfig.ALBIngress\x12G\n\rargo_rollouts\x18\x07 \x01(\x0b\x32\x30.prodvana.environment.ClusterConfig.ArgoRollouts\x12\x43\n\x0bgke_ingress\x18\x08 \x01(\x0b\x32..prodvana.environment.ClusterConfig.GKEIngress\x12P\n\x18self_managed_gke_ingress\x18\x0b \x01(\x0b\x32..prodvana.environment.ClusterConfig.GKEIngress\x12I\n\x0e\x63loud_provider\x18\t \x01(\x0e\x32\x31.prodvana.environment.ClusterConfig.CloudProvider\x12H\n\x19\x64\x65tected_cluster_metadata\x18\x12 \x01(\x0b\x32%.prodvana.environment.ClusterMetadata\x12M\n\x1euser_supplied_cluster_metadata\x18\x13 \x01(\x0b\x32%.prodvana.environment.ClusterMetadata\x12\x37\n\x04\x66\x61ke\x18\n \x01(\x0b\x32\'.prodvana.environment.FakeClusterConfigH\x00\x12\x41\n\textension\x18\x0c \x01(\x0b\x32,.prodvana.environment.ExtensionClusterConfigH\x00\x12>\n\x06\x63ustom\x18\x17 \x01(\x0b\x32,.prodvana.environment.ExtensionClusterConfigH\x00\x12G\n\x10terraform_runner\x18\x0f \x01(\x0b\x32+.prodvana.environment.TerraformRunnerConfigH\x00\x12\x44\n\rpulumi_runner\x18\x10 \x01(\x0b\x32+.prodvana.environment.TerraformRunnerConfigH\x00\x12\x35\n\x07\x61ws_ecs\x18\x14 \x01(\x0b\x32\".prodvana.environment.AwsEcsConfigH\x00\x12\x46\n\x10google_cloud_run\x18\x15 \x01(\x0b\x32*.prodvana.environment.GoogleCloudRunConfigH\x00\x12\x30\n\x06labels\x18\x11 \x03(\x0b\x32 .prodvana.labels.LabelDefinition\x12\x41\n\rauto_rollback\x18\x16 \x01(\x0b\x32*.prodvana.common_config.AutoRollbackConfig\x1aH\n\x08Kubecost\x12\x0f\n\x07\x65nabled\x18\x01 \x01(\x08\x12\x1a\n\x12kubecost_namespace\x18\x03 \x01(\t\x12\x0f\n\x07managed\x18\x02 \x01(\x08\x1a\x65\n\x07\x44\x61tadog\x12\x0f\n\x07\x65nabled\x18\x01 \x01(\x08\x12\x19\n\x11\x64\x61tadog_namespace\x18\x03 \x01(\t\x12\x0f\n\x07\x61pi_key\x18\x04 \x01(\t\x12\x0c\n\x04site\x18\x05 \x01(\t\x12\x0f\n\x07managed\x18\x02 \x01(\x08\x1a\x62\n\nALBIngress\x12\x0f\n\x07\x65nabled\x18\x01 \x01(\x08\x12\x1e\n\ringress_class\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12#\n\x1b\x64\x65\x66\x61ult_balancer_attributes\x18\x03 \x03(\t\x1a\xba\x03\n\x0c\x41rgoRollouts\x12\x0f\n\x07\x65nabled\x18\x01 \x01(\x08\x12T\n\ttemplates\x18\x02 \x03(\x0b\x32\x41.prodvana.environment.ClusterConfig.ArgoRollouts.AnalysisTemplate\x1a\xc2\x02\n\x10\x41nalysisTemplate\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x66\n\x0b\x61rg_mapping\x18\x02 \x03(\x0b\x32Q.prodvana.environment.ClusterConfig.ArgoRollouts.AnalysisTemplate.ArgMappingEntry\x1a}\n\x0f\x41rgMappingEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12Y\n\x05value\x18\x02 \x01(\x0e\x32J.prodvana.environment.ClusterConfig.ArgoRollouts.AnalysisTemplate.ArgValue:\x02\x38\x01\"9\n\x08\x41rgValue\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x0b\n\x07SERVICE\x10\x01\x12\x13\n\x0fRELEASE_CHANNEL\x10\x02\x1a\x37\n\nGKEIngress\x12\x0f\n\x07\x65nabled\x18\x01 \x01(\x08\x12\x18\n\x10\x63ontainer_native\x18\x02 \x01(\x08\"T\n\rCloudProvider\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x07\n\x03\x45KS\x10\x01\x12\x07\n\x03GKE\x10\x02\x12\x07\n\x03\x41KS\x10\x03\x12\n\n\x06ONPREM\x10\x04\x12\x0f\n\x0bOTHER_CLOUD\x10\x05\x42\x0f\n\rcluster_oneofJ\x04\x08\x01\x10\x02J\x04\x08\x02\x10\x03J\x04\x08\x05\x10\x06R\x0f\x64isable_flaggerR\rdisable_istioR\x0f\x61ws_alb_ingress*\x91\x01\n\x0b\x43lusterType\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x07\n\x03K8S\x10\x01\x12\x07\n\x03\x45\x43S\x10\x02\x12\x08\n\x04\x46\x41KE\x10\x03\x12\r\n\tEXTENSION\x10\x04\x12\x14\n\x10TERRAFORM_RUNNER\x10\x05\x12\x11\n\rPULUMI_RUNNER\x10\x06\x12\x0b\n\x07\x41WS_ECS\x10\x07\x12\x14\n\x10GOOGLE_CLOUD_RUN\x10\x08*n\n\rExtensionType\x12\x0b\n\x07GENERIC\x10\x00\x12\r\n\tTERRAFORM\x10\x01\x12\n\n\x06PULUMI\x10\x02\x12\x15\n\x11\x41WS_ECS_EXTENSION\x10\x03\x12\x1e\n\x1aGOOGLE_CLOUD_RUN_EXTENSION\x10\x04\x42PZNgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/environmentb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'prodvana.environment.clusters_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZNgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/environment'
  _CLUSTERAUTH_K8SAUTH_AGENTENVENTRY._options = None
  _CLUSTERAUTH_K8SAUTH_AGENTENVENTRY._serialized_options = b'8\001'
  _RETRYPOLICY.fields_by_name['base_interval']._options = None
  _RETRYPOLICY.fields_by_name['base_interval']._serialized_options = b'\372B\005\252\001\002*\000'
  _RETRYPOLICY.fields_by_name['max_interval']._options = None
  _RETRYPOLICY.fields_by_name['max_interval']._serialized_options = b'\372B\005\252\001\002*\000'
  _EXTENSIONFETCHCOMMAND_ENVENTRY._options = None
  _EXTENSIONFETCHCOMMAND_ENVENTRY._serialized_options = b'8\001'
  _EXTENSIONFETCHCOMMAND.oneofs_by_name['exec_config']._options = None
  _EXTENSIONFETCHCOMMAND.oneofs_by_name['exec_config']._serialized_options = b'\370B\001'
  _EXTENSIONFETCHCOMMAND.fields_by_name['env']._options = None
  _EXTENSIONFETCHCOMMAND.fields_by_name['env']._serialized_options = b'\372B$\232\001!\030\001\"\035r\0332\031^[a-zA-Z_]+[a-zA-Z0-9_]*$'
  _EXTENSIONGETINFOCOMMAND_ENVENTRY._options = None
  _EXTENSIONGETINFOCOMMAND_ENVENTRY._serialized_options = b'8\001'
  _EXTENSIONGETINFOCOMMAND.oneofs_by_name['exec_config']._options = None
  _EXTENSIONGETINFOCOMMAND.oneofs_by_name['exec_config']._serialized_options = b'\370B\001'
  _EXTENSIONGETINFOCOMMAND.fields_by_name['env']._options = None
  _EXTENSIONGETINFOCOMMAND.fields_by_name['env']._serialized_options = b'\372B$\232\001!\030\001\"\035r\0332\031^[a-zA-Z_]+[a-zA-Z0-9_]*$'
  _EXTENSIONAPPLYCOMMAND_ENVENTRY._options = None
  _EXTENSIONAPPLYCOMMAND_ENVENTRY._serialized_options = b'8\001'
  _EXTENSIONAPPLYCOMMAND.oneofs_by_name['exec_config']._options = None
  _EXTENSIONAPPLYCOMMAND.oneofs_by_name['exec_config']._serialized_options = b'\370B\001'
  _EXTENSIONAPPLYCOMMAND.fields_by_name['env']._options = None
  _EXTENSIONAPPLYCOMMAND.fields_by_name['env']._serialized_options = b'\372B$\232\001!\030\001\"\035r\0332\031^[a-zA-Z_]+[a-zA-Z0-9_]*$'
  _EXTENSIONCLUSTERCONFIG.fields_by_name['apply']._options = None
  _EXTENSIONCLUSTERCONFIG.fields_by_name['apply']._serialized_options = b'\372B\005\212\001\002\020\001'
  _EXTENSIONCLUSTERCONFIG.fields_by_name['parameters']._options = None
  _EXTENSIONCLUSTERCONFIG.fields_by_name['parameters']._serialized_options = b'\372B\n\222\001\007\"\005\212\001\002\020\001'
  _EXTENSIONCLUSTERCONFIG.fields_by_name['proxy_runtime']._options = None
  _EXTENSIONCLUSTERCONFIG.fields_by_name['proxy_runtime']._serialized_options = b'\372B\005\212\001\002\020\001'
  _EXTENSIONCLUSTERCONFIG.fields_by_name['convergence_grace_period']._options = None
  _EXTENSIONCLUSTERCONFIG.fields_by_name['convergence_grace_period']._serialized_options = b'\372B\005\252\001\002*\000'
  _COMPILEDEXTENSIONCOMMAND_COMPILEDEXTENSIONCOMMANDEXEC.oneofs_by_name['exec_config']._options = None
  _COMPILEDEXTENSIONCOMMAND_COMPILEDEXTENSIONCOMMANDEXEC.oneofs_by_name['exec_config']._serialized_options = b'\370B\001'
  _COMPILEDEXTENSIONCOMMAND_ENVENTRY._options = None
  _COMPILEDEXTENSIONCOMMAND_ENVENTRY._serialized_options = b'8\001'
  _IACRUNNERCOMMAND.fields_by_name['cmd']._options = None
  _IACRUNNERCOMMAND.fields_by_name['cmd']._serialized_options = b'\372B\004r\002\020\001'
  _TERRAFORMRUNNERCONFIG_ENVENTRY._options = None
  _TERRAFORMRUNNERCONFIG_ENVENTRY._serialized_options = b'8\001'
  _TERRAFORMRUNNERCONFIG.fields_by_name['proxy_runtime']._options = None
  _TERRAFORMRUNNERCONFIG.fields_by_name['proxy_runtime']._serialized_options = b'\372B\005\212\001\002\020\001'
  _TERRAFORMRUNNERCONFIG.fields_by_name['convergence_grace_period']._options = None
  _TERRAFORMRUNNERCONFIG.fields_by_name['convergence_grace_period']._serialized_options = b'\372B\005\252\001\002*\000'
  _AWSECSCONFIG_ACCESSKEY.fields_by_name['aws_access_key_id']._options = None
  _AWSECSCONFIG_ACCESSKEY.fields_by_name['aws_access_key_id']._serialized_options = b'\372B\004r\002\020\001'
  _AWSECSCONFIG_ACCESSKEY.fields_by_name['aws_secret_access_key']._options = None
  _AWSECSCONFIG_ACCESSKEY.fields_by_name['aws_secret_access_key']._serialized_options = b'\372B\005\212\001\002\020\001'
  _AWSECSCONFIG.oneofs_by_name['credentials']._options = None
  _AWSECSCONFIG.oneofs_by_name['credentials']._serialized_options = b'\370B\001'
  _AWSECSCONFIG.fields_by_name['proxy_runtime']._options = None
  _AWSECSCONFIG.fields_by_name['proxy_runtime']._serialized_options = b'\372B\005\212\001\002\020\001'
  _AWSECSCONFIG.fields_by_name['access_key']._options = None
  _AWSECSCONFIG.fields_by_name['access_key']._serialized_options = b'\372B\005\212\001\002\020\001'
  _AWSECSCONFIG.fields_by_name['region']._options = None
  _AWSECSCONFIG.fields_by_name['region']._serialized_options = b'\372B\004r\002\020\001'
  _AWSECSCONFIG.fields_by_name['ecs_cluster']._options = None
  _AWSECSCONFIG.fields_by_name['ecs_cluster']._serialized_options = b'\372B\004r\002\020\001'
  _GOOGLECLOUDRUNCONFIG.oneofs_by_name['credentials']._options = None
  _GOOGLECLOUDRUNCONFIG.oneofs_by_name['credentials']._serialized_options = b'\370B\001'
  _GOOGLECLOUDRUNCONFIG.fields_by_name['proxy_runtime']._options = None
  _GOOGLECLOUDRUNCONFIG.fields_by_name['proxy_runtime']._serialized_options = b'\372B\005\212\001\002\020\001'
  _GOOGLECLOUDRUNCONFIG.fields_by_name['service_account_json']._options = None
  _GOOGLECLOUDRUNCONFIG.fields_by_name['service_account_json']._serialized_options = b'\372B\005\212\001\002\020\001'
  _GOOGLECLOUDRUNCONFIG.fields_by_name['project']._options = None
  _GOOGLECLOUDRUNCONFIG.fields_by_name['project']._serialized_options = b'\372B\004r\002\020\001'
  _GOOGLECLOUDRUNCONFIG.fields_by_name['region']._options = None
  _GOOGLECLOUDRUNCONFIG.fields_by_name['region']._serialized_options = b'\372B\004r\002\020\001'
  _CLUSTERCONFIG_ALBINGRESS.fields_by_name['ingress_class']._options = None
  _CLUSTERCONFIG_ALBINGRESS.fields_by_name['ingress_class']._serialized_options = b'\372B\004r\002\020\001'
  _CLUSTERCONFIG_ARGOROLLOUTS_ANALYSISTEMPLATE_ARGMAPPINGENTRY._options = None
  _CLUSTERCONFIG_ARGOROLLOUTS_ANALYSISTEMPLATE_ARGMAPPINGENTRY._serialized_options = b'8\001'
  _globals['_CLUSTERTYPE']._serialized_start=8335
  _globals['_CLUSTERTYPE']._serialized_end=8480
  _globals['_EXTENSIONTYPE']._serialized_start=8482
  _globals['_EXTENSIONTYPE']._serialized_end=8592
  _globals['_CLUSTERAUTH']._serialized_start=494
  _globals['_CLUSTERAUTH']._serialized_end=1027
  _globals['_CLUSTERAUTH_ECSAUTH']._serialized_start=625
  _globals['_CLUSTERAUTH_ECSAUTH']._serialized_end=736
  _globals['_CLUSTERAUTH_K8SAUTH']._serialized_start=739
  _globals['_CLUSTERAUTH_K8SAUTH']._serialized_end=907
  _globals['_CLUSTERAUTH_K8SAUTH_AGENTENVENTRY']._serialized_start=860
  _globals['_CLUSTERAUTH_K8SAUTH_AGENTENVENTRY']._serialized_end=907
  _globals['_CLUSTER']._serialized_start=1030
  _globals['_CLUSTER']._serialized_end=1300
  _globals['_FAKECLUSTERCONFIG']._serialized_start=1303
  _globals['_FAKECLUSTERCONFIG']._serialized_end=1501
  _globals['_FAKECLUSTERCONFIG_CRASHINGPROGRAMPATTERNS']._serialized_start=1416
  _globals['_FAKECLUSTERCONFIG_CRASHINGPROGRAMPATTERNS']._serialized_end=1501
  _globals['_RETRYPOLICY']._serialized_start=1504
  _globals['_RETRYPOLICY']._serialized_end=1636
  _globals['_EXTENSIONFETCHCOMMAND']._serialized_start=1639
  _globals['_EXTENSIONFETCHCOMMAND']._serialized_end=2273
  _globals['_EXTENSIONFETCHCOMMAND_ENVENTRY']._serialized_start=2177
  _globals['_EXTENSIONFETCHCOMMAND_ENVENTRY']._serialized_end=2253
  _globals['_EXTENSIONGETINFOCOMMAND']._serialized_start=2276
  _globals['_EXTENSIONGETINFOCOMMAND']._serialized_end=2792
  _globals['_EXTENSIONGETINFOCOMMAND_ENVENTRY']._serialized_start=2177
  _globals['_EXTENSIONGETINFOCOMMAND_ENVENTRY']._serialized_end=2253
  _globals['_EXTENSIONAPPLYCOMMAND']._serialized_start=2795
  _globals['_EXTENSIONAPPLYCOMMAND']._serialized_end=3297
  _globals['_EXTENSIONAPPLYCOMMAND_ENVENTRY']._serialized_start=2177
  _globals['_EXTENSIONAPPLYCOMMAND_ENVENTRY']._serialized_end=2253
  _globals['_EXTENSIONCLUSTERCONFIG']._serialized_start=3300
  _globals['_EXTENSIONCLUSTERCONFIG']._serialized_end=3898
  _globals['_COMPILEDEXTENSIONCOMMAND']._serialized_start=3901
  _globals['_COMPILEDEXTENSIONCOMMAND']._serialized_end=4473
  _globals['_COMPILEDEXTENSIONCOMMAND_COMPILEDEXTENSIONCOMMANDEXEC']._serialized_start=4185
  _globals['_COMPILEDEXTENSIONCOMMAND_COMPILEDEXTENSIONCOMMANDEXEC']._serialized_end=4395
  _globals['_COMPILEDEXTENSIONCOMMAND_ENVENTRY']._serialized_start=2177
  _globals['_COMPILEDEXTENSIONCOMMAND_ENVENTRY']._serialized_end=2253
  _globals['_IACRUNNERCOMMAND']._serialized_start=4475
  _globals['_IACRUNNERCOMMAND']._serialized_end=4515
  _globals['_TERRAFORMRUNNERCONFIG']._serialized_start=4518
  _globals['_TERRAFORMRUNNERCONFIG']._serialized_end=5244
  _globals['_TERRAFORMRUNNERCONFIG_ENVENTRY']._serialized_start=2177
  _globals['_TERRAFORMRUNNERCONFIG_ENVENTRY']._serialized_end=2253
  _globals['_GKECLUSTERMETADATA']._serialized_start=5246
  _globals['_GKECLUSTERMETADATA']._serialized_end=5336
  _globals['_CLUSTERMETADATA']._serialized_start=5338
  _globals['_CLUSTERMETADATA']._serialized_end=5432
  _globals['_AWSECSCONFIG']._serialized_start=5435
  _globals['_AWSECSCONFIG']._serialized_end=5869
  _globals['_AWSECSCONFIG_ACCESSKEY']._serialized_start=5661
  _globals['_AWSECSCONFIG_ACCESSKEY']._serialized_end=5795
  _globals['_GOOGLECLOUDRUNCONFIG']._serialized_start=5872
  _globals['_GOOGLECLOUDRUNCONFIG']._serialized_end=6129
  _globals['_CLUSTERCONFIG']._serialized_start=6132
  _globals['_CLUSTERCONFIG']._serialized_end=8332
  _globals['_CLUSTERCONFIG_KUBECOST']._serialized_start=7385
  _globals['_CLUSTERCONFIG_KUBECOST']._serialized_end=7457
  _globals['_CLUSTERCONFIG_DATADOG']._serialized_start=7459
  _globals['_CLUSTERCONFIG_DATADOG']._serialized_end=7560
  _globals['_CLUSTERCONFIG_ALBINGRESS']._serialized_start=7562
  _globals['_CLUSTERCONFIG_ALBINGRESS']._serialized_end=7660
  _globals['_CLUSTERCONFIG_ARGOROLLOUTS']._serialized_start=7663
  _globals['_CLUSTERCONFIG_ARGOROLLOUTS']._serialized_end=8105
  _globals['_CLUSTERCONFIG_ARGOROLLOUTS_ANALYSISTEMPLATE']._serialized_start=7783
  _globals['_CLUSTERCONFIG_ARGOROLLOUTS_ANALYSISTEMPLATE']._serialized_end=8105
  _globals['_CLUSTERCONFIG_ARGOROLLOUTS_ANALYSISTEMPLATE_ARGMAPPINGENTRY']._serialized_start=7921
  _globals['_CLUSTERCONFIG_ARGOROLLOUTS_ANALYSISTEMPLATE_ARGMAPPINGENTRY']._serialized_end=8046
  _globals['_CLUSTERCONFIG_ARGOROLLOUTS_ANALYSISTEMPLATE_ARGVALUE']._serialized_start=8048
  _globals['_CLUSTERCONFIG_ARGOROLLOUTS_ANALYSISTEMPLATE_ARGVALUE']._serialized_end=8105
  _globals['_CLUSTERCONFIG_GKEINGRESS']._serialized_start=8107
  _globals['_CLUSTERCONFIG_GKEINGRESS']._serialized_end=8162
  _globals['_CLUSTERCONFIG_CLOUDPROVIDER']._serialized_start=8164
  _globals['_CLUSTERCONFIG_CLOUDPROVIDER']._serialized_end=8248
# @@protoc_insertion_point(module_scope)
