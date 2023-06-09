# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/environment/environment_manager.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from prodvana.proto.prodvana.environment import clusters_pb2 as prodvana_dot_environment_dot_clusters__pb2
from prodvana.proto.prodvana.config_writeback import writeback_pb2 as prodvana_dot_config__writeback_dot_writeback__pb2
from prodvana.proto.prodvana.version import source_metadata_pb2 as prodvana_dot_version_dot_source__metadata__pb2
from prodvana.proto.validate import validate_pb2 as validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n.prodvana/environment/environment_manager.proto\x12\x14prodvana.environment\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a#prodvana/environment/clusters.proto\x1a)prodvana/config_writeback/writeback.proto\x1a&prodvana/version/source_metadata.proto\x1a\x17validate/validate.proto\"3\n\x1aGetClusterAgentApiTokenReq\x12\x15\n\x04name\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\"0\n\x1bGetClusterAgentApiTokenResp\x12\x11\n\tapi_token\x18\x01 \x01(\t\"\xaf\x02\n\x0eLinkClusterReq\x12\x0c\n\x04name\x18\x01 \x01(\t\x12/\n\x04\x61uth\x18\x02 \x01(\x0b\x32!.prodvana.environment.ClusterAuth\x12\x18\n\x10prodvana_managed\x18\x03 \x01(\x08\x12\x15\n\rdisable_istio\x18\x04 \x01(\x08\x12\x17\n\x0f\x64isable_flagger\x18\x05 \x01(\x08\x12/\n\x04type\x18\x06 \x01(\x0e\x32!.prodvana.environment.ClusterType\x12(\n\x06source\x18\x07 \x01(\x0e\x32\x18.prodvana.version.Source\x12\x39\n\x0fsource_metadata\x18\x08 \x01(\x0b\x32 .prodvana.version.SourceMetadata\"\xa8\x01\n\x0fLinkClusterResp\x12\x0f\n\x07success\x18\x01 \x01(\x08\x12\x0b\n\x03msg\x18\x02 \x01(\t\x12\x12\n\ncluster_id\x18\x03 \x01(\t\x12\x15\n\rk8s_agent_url\x18\x04 \x01(\t\x12\x17\n\x0fk8s_agent_image\x18\x05 \x01(\t\x12\x16\n\x0ek8s_agent_args\x18\x06 \x03(\t\x12\x1b\n\x13k8s_agent_api_token\x18\x07 \x01(\t\" \n\x10RemoveClusterReq\x12\x0c\n\x04name\x18\x01 \x01(\t\"\x13\n\x11RemoveClusterResp\"\'\n\x11GetClusterAuthReq\x12\x12\n\ncluster_id\x18\x01 \x01(\t\"E\n\x12GetClusterAuthResp\x12/\n\x04\x61uth\x18\x01 \x01(\x0b\x32!.prodvana.environment.ClusterAuth\"\x11\n\x0fListClustersReq\"\x85\x05\n\x10ListClustersResp\x12\x44\n\x08\x63lusters\x18\x01 \x03(\x0b\x32\x32.prodvana.environment.ListClustersResp.ClusterInfo\x1a\xaa\x04\n\x0b\x43lusterInfo\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\n\n\x02id\x18\x02 \x01(\t\x12\x34\n\x06origin\x18\x03 \x01(\x0e\x32$.prodvana.environment.Cluster.Origin\x12\x10\n\x08\x65ndpoint\x18\x04 \x01(\t\x12\x17\n\x0fservice_account\x18\x05 \x01(\t\x12H\n\x10writeback_config\x18\x06 \x01(\x0b\x32..prodvana.config_writeback.ConfigWritebackPath\x12/\n\x04type\x18\x07 \x01(\x0e\x32!.prodvana.environment.ClusterType\x12I\n\x03\x65\x63s\x18\x08 \x01(\x0b\x32:.prodvana.environment.ListClustersResp.ClusterInfo.EcsInfoH\x00\x12<\n\x18last_heartbeat_timestamp\x18\t \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x33\n\x06\x63onfig\x18\n \x01(\x0b\x32#.prodvana.environment.ClusterConfig\x12/\n\x04\x61uth\x18\x0b \x01(\x0b\x32!.prodvana.environment.ClusterAuth\x1a.\n\x07\x45\x63sInfo\x12\x0e\n\x06region\x18\x01 \x01(\t\x12\x13\n\x0b\x63luster_arn\x18\x02 \x01(\tB\x06\n\x04info\"6\n\rGetClusterReq\x12\x0f\n\x07runtime\x18\x01 \x01(\t\x12\x14\n\x0cinclude_auth\x18\x02 \x01(\x08\"U\n\x0eGetClusterResp\x12\x43\n\x07\x63luster\x18\x01 \x01(\x0b\x32\x32.prodvana.environment.ListClustersResp.ClusterInfo\"\xc5\x01\n\x13\x43onfigureClusterReq\x12\x14\n\x0cruntime_name\x18\x01 \x01(\t\x12\x33\n\x06\x63onfig\x18\x02 \x01(\x0b\x32#.prodvana.environment.ClusterConfig\x12(\n\x06source\x18\x03 \x01(\x0e\x32\x18.prodvana.version.Source\x12\x39\n\x0fsource_metadata\x18\x04 \x01(\x0b\x32 .prodvana.version.SourceMetadata\"\x16\n\x14\x43onfigureClusterResp\"\x1e\n\x1cValidateConfigureClusterResp\"+\n\x13GetClusterConfigReq\x12\x14\n\x0cruntime_name\x18\x01 \x01(\t\"K\n\x14GetClusterConfigResp\x12\x33\n\x06\x63onfig\x18\x02 \x01(\x0b\x32#.prodvana.environment.ClusterConfig\".\n\x16\x44\x65tectClusterConfigReq\x12\x14\n\x0cruntime_name\x18\x01 \x01(\t\"N\n\x17\x44\x65tectClusterConfigResp\x12\x33\n\x06\x63onfig\x18\x02 \x01(\x0b\x32#.prodvana.environment.ClusterConfig\")\n\x13GetClusterStatusReq\x12\x12\n\ncluster_id\x18\x01 \x01(\t\"T\n\x14GetClusterStatusResp\x12<\n\x18last_heartbeat_timestamp\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp2\xc3\x0c\n\x12\x45nvironmentManager\x12\xa3\x01\n\x17GetClusterAgentApiToken\x12\x30.prodvana.environment.GetClusterAgentApiTokenReq\x1a\x31.prodvana.environment.GetClusterAgentApiTokenResp\"#\x82\xd3\xe4\x93\x02\x1d\x12\x1b/v1/clusters/{name=*}/token\x12x\n\x0bLinkCluster\x12$.prodvana.environment.LinkClusterReq\x1a%.prodvana.environment.LinkClusterResp\"\x1c\x82\xd3\xe4\x93\x02\x16\"\x11/v1/clusters/link:\x01*\x12x\n\x0cListClusters\x12%.prodvana.environment.ListClustersReq\x1a&.prodvana.environment.ListClustersResp\"\x19\x82\xd3\xe4\x93\x02\x13\x12\x11/v1/clusters/list\x12}\n\nGetCluster\x12#.prodvana.environment.GetClusterReq\x1a$.prodvana.environment.GetClusterResp\"$\x82\xd3\xe4\x93\x02\x1e\x12\x1c/v1/clusters/{runtime=*}/get\x12\x80\x01\n\rRemoveCluster\x12&.prodvana.environment.RemoveClusterReq\x1a\'.prodvana.environment.RemoveClusterResp\"\x1e\x82\xd3\xe4\x93\x02\x18\"\x13/v1/clusters/remove:\x01*\x12\x82\x01\n\x0eGetClusterAuth\x12\'.prodvana.environment.GetClusterAuthReq\x1a(.prodvana.environment.GetClusterAuthResp\"\x1d\x82\xd3\xe4\x93\x02\x17\x12\x15/v1/clusters/get_auth\x12\x9a\x01\n\x10GetClusterConfig\x12).prodvana.environment.GetClusterConfigReq\x1a*.prodvana.environment.GetClusterConfigResp\"/\x82\xd3\xe4\x93\x02)\"$/v1/clusters/{runtime_name=*}/config:\x01*\x12\xaa\x01\n\x13\x44\x65tectClusterConfig\x12,.prodvana.environment.DetectClusterConfigReq\x1a-.prodvana.environment.DetectClusterConfigResp\"6\x82\xd3\xe4\x93\x02\x30\"+/v1/clusters/{runtime_name=*}/detect-config:\x01*\x12\x8c\x01\n\x10\x43onfigureCluster\x12).prodvana.environment.ConfigureClusterReq\x1a*.prodvana.environment.ConfigureClusterResp\"!\x82\xd3\xe4\x93\x02\x1b\"\x16/v1/clusters/configure:\x01*\x12\xa5\x01\n\x18ValidateConfigureCluster\x12).prodvana.environment.ConfigureClusterReq\x1a\x32.prodvana.environment.ValidateConfigureClusterResp\"*\x82\xd3\xe4\x93\x02$\"\x1f/v1/clusters/configure/validate:\x01*\x12\x8a\x01\n\x10GetClusterStatus\x12).prodvana.environment.GetClusterStatusReq\x1a*.prodvana.environment.GetClusterStatusResp\"\x1f\x82\xd3\xe4\x93\x02\x19\x12\x17/v1/clusters/get_statusBPZNgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/environmentb\x06proto3')



_GETCLUSTERAGENTAPITOKENREQ = DESCRIPTOR.message_types_by_name['GetClusterAgentApiTokenReq']
_GETCLUSTERAGENTAPITOKENRESP = DESCRIPTOR.message_types_by_name['GetClusterAgentApiTokenResp']
_LINKCLUSTERREQ = DESCRIPTOR.message_types_by_name['LinkClusterReq']
_LINKCLUSTERRESP = DESCRIPTOR.message_types_by_name['LinkClusterResp']
_REMOVECLUSTERREQ = DESCRIPTOR.message_types_by_name['RemoveClusterReq']
_REMOVECLUSTERRESP = DESCRIPTOR.message_types_by_name['RemoveClusterResp']
_GETCLUSTERAUTHREQ = DESCRIPTOR.message_types_by_name['GetClusterAuthReq']
_GETCLUSTERAUTHRESP = DESCRIPTOR.message_types_by_name['GetClusterAuthResp']
_LISTCLUSTERSREQ = DESCRIPTOR.message_types_by_name['ListClustersReq']
_LISTCLUSTERSRESP = DESCRIPTOR.message_types_by_name['ListClustersResp']
_LISTCLUSTERSRESP_CLUSTERINFO = _LISTCLUSTERSRESP.nested_types_by_name['ClusterInfo']
_LISTCLUSTERSRESP_CLUSTERINFO_ECSINFO = _LISTCLUSTERSRESP_CLUSTERINFO.nested_types_by_name['EcsInfo']
_GETCLUSTERREQ = DESCRIPTOR.message_types_by_name['GetClusterReq']
_GETCLUSTERRESP = DESCRIPTOR.message_types_by_name['GetClusterResp']
_CONFIGURECLUSTERREQ = DESCRIPTOR.message_types_by_name['ConfigureClusterReq']
_CONFIGURECLUSTERRESP = DESCRIPTOR.message_types_by_name['ConfigureClusterResp']
_VALIDATECONFIGURECLUSTERRESP = DESCRIPTOR.message_types_by_name['ValidateConfigureClusterResp']
_GETCLUSTERCONFIGREQ = DESCRIPTOR.message_types_by_name['GetClusterConfigReq']
_GETCLUSTERCONFIGRESP = DESCRIPTOR.message_types_by_name['GetClusterConfigResp']
_DETECTCLUSTERCONFIGREQ = DESCRIPTOR.message_types_by_name['DetectClusterConfigReq']
_DETECTCLUSTERCONFIGRESP = DESCRIPTOR.message_types_by_name['DetectClusterConfigResp']
_GETCLUSTERSTATUSREQ = DESCRIPTOR.message_types_by_name['GetClusterStatusReq']
_GETCLUSTERSTATUSRESP = DESCRIPTOR.message_types_by_name['GetClusterStatusResp']
GetClusterAgentApiTokenReq = _reflection.GeneratedProtocolMessageType('GetClusterAgentApiTokenReq', (_message.Message,), {
  'DESCRIPTOR' : _GETCLUSTERAGENTAPITOKENREQ,
  '__module__' : 'prodvana.environment.environment_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.environment.GetClusterAgentApiTokenReq)
  })
_sym_db.RegisterMessage(GetClusterAgentApiTokenReq)

GetClusterAgentApiTokenResp = _reflection.GeneratedProtocolMessageType('GetClusterAgentApiTokenResp', (_message.Message,), {
  'DESCRIPTOR' : _GETCLUSTERAGENTAPITOKENRESP,
  '__module__' : 'prodvana.environment.environment_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.environment.GetClusterAgentApiTokenResp)
  })
_sym_db.RegisterMessage(GetClusterAgentApiTokenResp)

LinkClusterReq = _reflection.GeneratedProtocolMessageType('LinkClusterReq', (_message.Message,), {
  'DESCRIPTOR' : _LINKCLUSTERREQ,
  '__module__' : 'prodvana.environment.environment_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.environment.LinkClusterReq)
  })
_sym_db.RegisterMessage(LinkClusterReq)

LinkClusterResp = _reflection.GeneratedProtocolMessageType('LinkClusterResp', (_message.Message,), {
  'DESCRIPTOR' : _LINKCLUSTERRESP,
  '__module__' : 'prodvana.environment.environment_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.environment.LinkClusterResp)
  })
_sym_db.RegisterMessage(LinkClusterResp)

RemoveClusterReq = _reflection.GeneratedProtocolMessageType('RemoveClusterReq', (_message.Message,), {
  'DESCRIPTOR' : _REMOVECLUSTERREQ,
  '__module__' : 'prodvana.environment.environment_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.environment.RemoveClusterReq)
  })
_sym_db.RegisterMessage(RemoveClusterReq)

RemoveClusterResp = _reflection.GeneratedProtocolMessageType('RemoveClusterResp', (_message.Message,), {
  'DESCRIPTOR' : _REMOVECLUSTERRESP,
  '__module__' : 'prodvana.environment.environment_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.environment.RemoveClusterResp)
  })
_sym_db.RegisterMessage(RemoveClusterResp)

GetClusterAuthReq = _reflection.GeneratedProtocolMessageType('GetClusterAuthReq', (_message.Message,), {
  'DESCRIPTOR' : _GETCLUSTERAUTHREQ,
  '__module__' : 'prodvana.environment.environment_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.environment.GetClusterAuthReq)
  })
_sym_db.RegisterMessage(GetClusterAuthReq)

GetClusterAuthResp = _reflection.GeneratedProtocolMessageType('GetClusterAuthResp', (_message.Message,), {
  'DESCRIPTOR' : _GETCLUSTERAUTHRESP,
  '__module__' : 'prodvana.environment.environment_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.environment.GetClusterAuthResp)
  })
_sym_db.RegisterMessage(GetClusterAuthResp)

ListClustersReq = _reflection.GeneratedProtocolMessageType('ListClustersReq', (_message.Message,), {
  'DESCRIPTOR' : _LISTCLUSTERSREQ,
  '__module__' : 'prodvana.environment.environment_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.environment.ListClustersReq)
  })
_sym_db.RegisterMessage(ListClustersReq)

ListClustersResp = _reflection.GeneratedProtocolMessageType('ListClustersResp', (_message.Message,), {

  'ClusterInfo' : _reflection.GeneratedProtocolMessageType('ClusterInfo', (_message.Message,), {

    'EcsInfo' : _reflection.GeneratedProtocolMessageType('EcsInfo', (_message.Message,), {
      'DESCRIPTOR' : _LISTCLUSTERSRESP_CLUSTERINFO_ECSINFO,
      '__module__' : 'prodvana.environment.environment_manager_pb2'
      # @@protoc_insertion_point(class_scope:prodvana.environment.ListClustersResp.ClusterInfo.EcsInfo)
      })
    ,
    'DESCRIPTOR' : _LISTCLUSTERSRESP_CLUSTERINFO,
    '__module__' : 'prodvana.environment.environment_manager_pb2'
    # @@protoc_insertion_point(class_scope:prodvana.environment.ListClustersResp.ClusterInfo)
    })
  ,
  'DESCRIPTOR' : _LISTCLUSTERSRESP,
  '__module__' : 'prodvana.environment.environment_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.environment.ListClustersResp)
  })
_sym_db.RegisterMessage(ListClustersResp)
_sym_db.RegisterMessage(ListClustersResp.ClusterInfo)
_sym_db.RegisterMessage(ListClustersResp.ClusterInfo.EcsInfo)

GetClusterReq = _reflection.GeneratedProtocolMessageType('GetClusterReq', (_message.Message,), {
  'DESCRIPTOR' : _GETCLUSTERREQ,
  '__module__' : 'prodvana.environment.environment_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.environment.GetClusterReq)
  })
_sym_db.RegisterMessage(GetClusterReq)

GetClusterResp = _reflection.GeneratedProtocolMessageType('GetClusterResp', (_message.Message,), {
  'DESCRIPTOR' : _GETCLUSTERRESP,
  '__module__' : 'prodvana.environment.environment_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.environment.GetClusterResp)
  })
_sym_db.RegisterMessage(GetClusterResp)

ConfigureClusterReq = _reflection.GeneratedProtocolMessageType('ConfigureClusterReq', (_message.Message,), {
  'DESCRIPTOR' : _CONFIGURECLUSTERREQ,
  '__module__' : 'prodvana.environment.environment_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.environment.ConfigureClusterReq)
  })
_sym_db.RegisterMessage(ConfigureClusterReq)

ConfigureClusterResp = _reflection.GeneratedProtocolMessageType('ConfigureClusterResp', (_message.Message,), {
  'DESCRIPTOR' : _CONFIGURECLUSTERRESP,
  '__module__' : 'prodvana.environment.environment_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.environment.ConfigureClusterResp)
  })
_sym_db.RegisterMessage(ConfigureClusterResp)

ValidateConfigureClusterResp = _reflection.GeneratedProtocolMessageType('ValidateConfigureClusterResp', (_message.Message,), {
  'DESCRIPTOR' : _VALIDATECONFIGURECLUSTERRESP,
  '__module__' : 'prodvana.environment.environment_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.environment.ValidateConfigureClusterResp)
  })
_sym_db.RegisterMessage(ValidateConfigureClusterResp)

GetClusterConfigReq = _reflection.GeneratedProtocolMessageType('GetClusterConfigReq', (_message.Message,), {
  'DESCRIPTOR' : _GETCLUSTERCONFIGREQ,
  '__module__' : 'prodvana.environment.environment_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.environment.GetClusterConfigReq)
  })
_sym_db.RegisterMessage(GetClusterConfigReq)

GetClusterConfigResp = _reflection.GeneratedProtocolMessageType('GetClusterConfigResp', (_message.Message,), {
  'DESCRIPTOR' : _GETCLUSTERCONFIGRESP,
  '__module__' : 'prodvana.environment.environment_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.environment.GetClusterConfigResp)
  })
_sym_db.RegisterMessage(GetClusterConfigResp)

DetectClusterConfigReq = _reflection.GeneratedProtocolMessageType('DetectClusterConfigReq', (_message.Message,), {
  'DESCRIPTOR' : _DETECTCLUSTERCONFIGREQ,
  '__module__' : 'prodvana.environment.environment_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.environment.DetectClusterConfigReq)
  })
_sym_db.RegisterMessage(DetectClusterConfigReq)

DetectClusterConfigResp = _reflection.GeneratedProtocolMessageType('DetectClusterConfigResp', (_message.Message,), {
  'DESCRIPTOR' : _DETECTCLUSTERCONFIGRESP,
  '__module__' : 'prodvana.environment.environment_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.environment.DetectClusterConfigResp)
  })
_sym_db.RegisterMessage(DetectClusterConfigResp)

GetClusterStatusReq = _reflection.GeneratedProtocolMessageType('GetClusterStatusReq', (_message.Message,), {
  'DESCRIPTOR' : _GETCLUSTERSTATUSREQ,
  '__module__' : 'prodvana.environment.environment_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.environment.GetClusterStatusReq)
  })
_sym_db.RegisterMessage(GetClusterStatusReq)

GetClusterStatusResp = _reflection.GeneratedProtocolMessageType('GetClusterStatusResp', (_message.Message,), {
  'DESCRIPTOR' : _GETCLUSTERSTATUSRESP,
  '__module__' : 'prodvana.environment.environment_manager_pb2'
  # @@protoc_insertion_point(class_scope:prodvana.environment.GetClusterStatusResp)
  })
_sym_db.RegisterMessage(GetClusterStatusResp)

_ENVIRONMENTMANAGER = DESCRIPTOR.services_by_name['EnvironmentManager']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZNgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/environment'
  _GETCLUSTERAGENTAPITOKENREQ.fields_by_name['name']._options = None
  _GETCLUSTERAGENTAPITOKENREQ.fields_by_name['name']._serialized_options = b'\372B\004r\002\020\001'
  _ENVIRONMENTMANAGER.methods_by_name['GetClusterAgentApiToken']._options = None
  _ENVIRONMENTMANAGER.methods_by_name['GetClusterAgentApiToken']._serialized_options = b'\202\323\344\223\002\035\022\033/v1/clusters/{name=*}/token'
  _ENVIRONMENTMANAGER.methods_by_name['LinkCluster']._options = None
  _ENVIRONMENTMANAGER.methods_by_name['LinkCluster']._serialized_options = b'\202\323\344\223\002\026\"\021/v1/clusters/link:\001*'
  _ENVIRONMENTMANAGER.methods_by_name['ListClusters']._options = None
  _ENVIRONMENTMANAGER.methods_by_name['ListClusters']._serialized_options = b'\202\323\344\223\002\023\022\021/v1/clusters/list'
  _ENVIRONMENTMANAGER.methods_by_name['GetCluster']._options = None
  _ENVIRONMENTMANAGER.methods_by_name['GetCluster']._serialized_options = b'\202\323\344\223\002\036\022\034/v1/clusters/{runtime=*}/get'
  _ENVIRONMENTMANAGER.methods_by_name['RemoveCluster']._options = None
  _ENVIRONMENTMANAGER.methods_by_name['RemoveCluster']._serialized_options = b'\202\323\344\223\002\030\"\023/v1/clusters/remove:\001*'
  _ENVIRONMENTMANAGER.methods_by_name['GetClusterAuth']._options = None
  _ENVIRONMENTMANAGER.methods_by_name['GetClusterAuth']._serialized_options = b'\202\323\344\223\002\027\022\025/v1/clusters/get_auth'
  _ENVIRONMENTMANAGER.methods_by_name['GetClusterConfig']._options = None
  _ENVIRONMENTMANAGER.methods_by_name['GetClusterConfig']._serialized_options = b'\202\323\344\223\002)\"$/v1/clusters/{runtime_name=*}/config:\001*'
  _ENVIRONMENTMANAGER.methods_by_name['DetectClusterConfig']._options = None
  _ENVIRONMENTMANAGER.methods_by_name['DetectClusterConfig']._serialized_options = b'\202\323\344\223\0020\"+/v1/clusters/{runtime_name=*}/detect-config:\001*'
  _ENVIRONMENTMANAGER.methods_by_name['ConfigureCluster']._options = None
  _ENVIRONMENTMANAGER.methods_by_name['ConfigureCluster']._serialized_options = b'\202\323\344\223\002\033\"\026/v1/clusters/configure:\001*'
  _ENVIRONMENTMANAGER.methods_by_name['ValidateConfigureCluster']._options = None
  _ENVIRONMENTMANAGER.methods_by_name['ValidateConfigureCluster']._serialized_options = b'\202\323\344\223\002$\"\037/v1/clusters/configure/validate:\001*'
  _ENVIRONMENTMANAGER.methods_by_name['GetClusterStatus']._options = None
  _ENVIRONMENTMANAGER.methods_by_name['GetClusterStatus']._serialized_options = b'\202\323\344\223\002\031\022\027/v1/clusters/get_status'
  _GETCLUSTERAGENTAPITOKENREQ._serialized_start=280
  _GETCLUSTERAGENTAPITOKENREQ._serialized_end=331
  _GETCLUSTERAGENTAPITOKENRESP._serialized_start=333
  _GETCLUSTERAGENTAPITOKENRESP._serialized_end=381
  _LINKCLUSTERREQ._serialized_start=384
  _LINKCLUSTERREQ._serialized_end=687
  _LINKCLUSTERRESP._serialized_start=690
  _LINKCLUSTERRESP._serialized_end=858
  _REMOVECLUSTERREQ._serialized_start=860
  _REMOVECLUSTERREQ._serialized_end=892
  _REMOVECLUSTERRESP._serialized_start=894
  _REMOVECLUSTERRESP._serialized_end=913
  _GETCLUSTERAUTHREQ._serialized_start=915
  _GETCLUSTERAUTHREQ._serialized_end=954
  _GETCLUSTERAUTHRESP._serialized_start=956
  _GETCLUSTERAUTHRESP._serialized_end=1025
  _LISTCLUSTERSREQ._serialized_start=1027
  _LISTCLUSTERSREQ._serialized_end=1044
  _LISTCLUSTERSRESP._serialized_start=1047
  _LISTCLUSTERSRESP._serialized_end=1692
  _LISTCLUSTERSRESP_CLUSTERINFO._serialized_start=1138
  _LISTCLUSTERSRESP_CLUSTERINFO._serialized_end=1692
  _LISTCLUSTERSRESP_CLUSTERINFO_ECSINFO._serialized_start=1638
  _LISTCLUSTERSRESP_CLUSTERINFO_ECSINFO._serialized_end=1684
  _GETCLUSTERREQ._serialized_start=1694
  _GETCLUSTERREQ._serialized_end=1748
  _GETCLUSTERRESP._serialized_start=1750
  _GETCLUSTERRESP._serialized_end=1835
  _CONFIGURECLUSTERREQ._serialized_start=1838
  _CONFIGURECLUSTERREQ._serialized_end=2035
  _CONFIGURECLUSTERRESP._serialized_start=2037
  _CONFIGURECLUSTERRESP._serialized_end=2059
  _VALIDATECONFIGURECLUSTERRESP._serialized_start=2061
  _VALIDATECONFIGURECLUSTERRESP._serialized_end=2091
  _GETCLUSTERCONFIGREQ._serialized_start=2093
  _GETCLUSTERCONFIGREQ._serialized_end=2136
  _GETCLUSTERCONFIGRESP._serialized_start=2138
  _GETCLUSTERCONFIGRESP._serialized_end=2213
  _DETECTCLUSTERCONFIGREQ._serialized_start=2215
  _DETECTCLUSTERCONFIGREQ._serialized_end=2261
  _DETECTCLUSTERCONFIGRESP._serialized_start=2263
  _DETECTCLUSTERCONFIGRESP._serialized_end=2341
  _GETCLUSTERSTATUSREQ._serialized_start=2343
  _GETCLUSTERSTATUSREQ._serialized_end=2384
  _GETCLUSTERSTATUSRESP._serialized_start=2386
  _GETCLUSTERSTATUSRESP._serialized_end=2470
  _ENVIRONMENTMANAGER._serialized_start=2473
  _ENVIRONMENTMANAGER._serialized_end=4076
# @@protoc_insertion_point(module_scope)
