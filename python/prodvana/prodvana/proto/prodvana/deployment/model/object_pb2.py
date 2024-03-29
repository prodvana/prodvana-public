# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/deployment/model/object.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from prodvana.proto.prodvana.object import meta_pb2 as prodvana_dot_object_dot_meta__pb2
from prodvana.proto.prodvana.repo import repo_pb2 as prodvana_dot_repo_dot_repo__pb2
from prodvana.proto.validate import validate_pb2 as validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n&prodvana/deployment/model/object.proto\x12\x19prodvana.deployment.model\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1aprodvana/object/meta.proto\x1a\x18prodvana/repo/repo.proto\x1a\x17validate/validate.proto\"\xf1\x02\n\x10\x44\x65ploymentConfig\x12\x36\n\x12\x63reation_timestamp\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\"\n\x11\x64\x65ployment_system\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x0f\n\x07service\x18\x03 \x01(\t\x12\x17\n\x0frelease_channel\x18\x04 \x01(\t\x12\x13\n\x0b\x61pplication\x18\x08 \x01(\t\x12\x12\n\nrepository\x18\x05 \x01(\t\x12\x11\n\tcommit_id\x18\x06 \x01(\t\x12\x0c\n\x04user\x18\x07 \x01(\t\x12\x16\n\x0e\x61pplication_id\x18\t \x01(\t\x12\x12\n\nservice_id\x18\n \x01(\t\x12\x1a\n\x12release_channel_id\x18\x0b \x01(\t\x12\x17\n\x0fservice_version\x18\x0c \x01(\t\x12\x18\n\x10\x64\x65sired_state_id\x18\r \x01(\t\x12\x12\n\nrelease_id\x18\x0e \x01(\t\"\x89\x01\n\x0f\x44\x65ploymentState\x12;\n\x06status\x18\x01 \x01(\x0e\x32+.prodvana.deployment.model.DeploymentStatus\x12\x39\n\x15last_update_timestamp\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\x8e\x01\n\x0e\x43ommitAnalysis\x12\x15\n\rcommits_added\x18\x01 \x01(\x03\x12\x17\n\x0f\x63ommits_removed\x18\x02 \x01(\x03\x12L\n\x0fimpact_analysis\x18\x03 \x01(\x0b\x32\x33.prodvana.deployment.model.ImpactAnalysisComparison\"\xa6\x01\n\x18ImpactAnalysisComparison\x12\x35\n\x16relevant_added_commits\x18\x01 \x03(\x0b\x32\x15.prodvana.repo.Commit\x12\x37\n\x18relevant_removed_commits\x18\x03 \x03(\x0b\x32\x15.prodvana.repo.Commit\x12\x1a\n\x12unanalyzed_commits\x18\x02 \x01(\x03\"\xbe\x03\n\x14\x44\x65ploymentComparison\x12)\n\x04prev\x18\x01 \x01(\x0b\x32\x1b.prodvana.object.ObjectMeta\x12\x17\n\x0fprev_repository\x18\x02 \x01(\t\x12\x16\n\x0enew_repository\x18\x03 \x01(\t\x12\x16\n\x0eprev_commit_id\x18\x04 \x01(\t\x12\x15\n\rnew_commit_id\x18\x05 \x01(\t\x12\x42\n\x0f\x63ommit_analysis\x18\x0e \x01(\x0b\x32).prodvana.deployment.model.CommitAnalysis\x12\x17\n\x0fprev_service_id\x18\x08 \x01(\t\x12\x1f\n\x17prev_release_channel_id\x18\t \x01(\t\x12\x1c\n\x14prev_service_version\x18\n \x01(\t\x12\x16\n\x0enew_service_id\x18\x0b \x01(\t\x12\x1e\n\x16new_release_channel_id\x18\x0c \x01(\t\x12\x1b\n\x13new_service_version\x18\r \x01(\tJ\x04\x08\x06\x10\x07J\x04\x08\x07\x10\x08R\x0fimpact_analysisR\rtotal_commits\"\xf4\x01\n\nDeployment\x12)\n\x04meta\x18\x01 \x01(\x0b\x32\x1b.prodvana.object.ObjectMeta\x12;\n\x06\x63onfig\x18\x02 \x01(\x0b\x32+.prodvana.deployment.model.DeploymentConfig\x12\x43\n\ncomparison\x18\x03 \x01(\x0b\x32/.prodvana.deployment.model.DeploymentComparison\x12\x39\n\x05state\x18\x04 \x01(\x0b\x32*.prodvana.deployment.model.DeploymentState*[\n\x10\x44\x65ploymentStatus\x12\x12\n\x0eUNKNOWN_STATUS\x10\x00\x12\x0b\n\x07PENDING\x10\x01\x12\r\n\tSUCCEEDED\x10\x02\x12\n\n\x06\x46\x41ILED\x10\x03\x12\x0b\n\x07PREVIEW\x10\x04\x42UZSgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/deployment/modelb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'prodvana.deployment.model.object_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZSgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/deployment/model'
  _DEPLOYMENTCONFIG.fields_by_name['deployment_system']._options = None
  _DEPLOYMENTCONFIG.fields_by_name['deployment_system']._serialized_options = b'\372B\004r\002\020\001'
  _globals['_DEPLOYMENTSTATUS']._serialized_start=1703
  _globals['_DEPLOYMENTSTATUS']._serialized_end=1794
  _globals['_DEPLOYMENTCONFIG']._serialized_start=182
  _globals['_DEPLOYMENTCONFIG']._serialized_end=551
  _globals['_DEPLOYMENTSTATE']._serialized_start=554
  _globals['_DEPLOYMENTSTATE']._serialized_end=691
  _globals['_COMMITANALYSIS']._serialized_start=694
  _globals['_COMMITANALYSIS']._serialized_end=836
  _globals['_IMPACTANALYSISCOMPARISON']._serialized_start=839
  _globals['_IMPACTANALYSISCOMPARISON']._serialized_end=1005
  _globals['_DEPLOYMENTCOMPARISON']._serialized_start=1008
  _globals['_DEPLOYMENTCOMPARISON']._serialized_end=1454
  _globals['_DEPLOYMENT']._serialized_start=1457
  _globals['_DEPLOYMENT']._serialized_end=1701
# @@protoc_insertion_point(module_scope)
