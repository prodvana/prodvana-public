# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/deployment/manager.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from prodvana.proto.prodvana.object import meta_pb2 as prodvana_dot_object_dot_meta__pb2
from prodvana.proto.prodvana.deployment import object_pb2 as prodvana_dot_deployment_dot_object__pb2
from prodvana.proto.validate import validate_pb2 as validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n!prodvana/deployment/manager.proto\x12\x13prodvana.deployment\x1a\x1cgoogle/api/annotations.proto\x1a\x1aprodvana/object/meta.proto\x1a prodvana/deployment/object.proto\x1a\x17validate/validate.proto\"g\n\x13RecordDeploymentReq\x12?\n\x06\x63onfig\x18\x01 \x01(\x0b\x32%.prodvana.deployment.DeploymentConfigB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12\x0f\n\x07pending\x18\x02 \x01(\x08\"A\n\x14RecordDeploymentResp\x12)\n\x04meta\x18\x01 \x01(\x0b\x32\x1b.prodvana.object.ObjectMeta\"i\n\x19UpdateDeploymentStatusReq\x12\x15\n\rdeployment_id\x18\x01 \x01(\t\x12\x35\n\x06status\x18\x02 \x01(\x0e\x32%.prodvana.deployment.DeploymentStatus\"S\n\x1aUpdateDeploymentStatusResp\x12\x35\n\x06status\x18\x01 \x01(\x0e\x32%.prodvana.deployment.DeploymentStatus\"m\n\x10\x44\x65ploymentFilter\x12\x10\n\x08services\x18\x01 \x03(\t\x12\x18\n\x10release_channels\x18\x02 \x03(\t\x12\x13\n\x0b\x61pplication\x18\x03 \x01(\t\x12\x18\n\x10\x64\x65sired_state_id\x18\x04 \x01(\t\"\xe8\x01\n\x12ListDeploymentsReq\x12\x36\n\x07\x66ilters\x18\x01 \x03(\x0b\x32%.prodvana.deployment.DeploymentFilter\x12\x35\n\x06\x66ilter\x18\x02 \x01(\x0b\x32%.prodvana.deployment.DeploymentFilter\x12\x1e\n\x16starting_deployment_id\x18\x03 \x01(\t\x12\x1c\n\x14\x65nding_deployment_id\x18\x04 \x01(\t\x12\x12\n\npage_token\x18\x05 \x01(\t\x12\x11\n\tpage_size\x18\x06 \x01(\x05\"d\n\x13ListDeploymentsResp\x12\x34\n\x0b\x64\x65ployments\x18\x01 \x03(\x0b\x32\x1f.prodvana.deployment.Deployment\x12\x17\n\x0fnext_page_token\x18\x02 \x01(\t\"m\n\rDeploymentRef\x12\x17\n\rdeployment_id\x18\x01 \x01(\tH\x00\x12\x37\n\x06\x63onfig\x18\x02 \x01(\x0b\x32%.prodvana.deployment.DeploymentConfigH\x00\x42\n\n\x03ref\x12\x03\xf8\x42\x01\"\x99\x01\n\x14\x43ompareDeploymentReq\x12\x44\n\x0enew_deployment\x18\x01 \x01(\x0b\x32\".prodvana.deployment.DeploymentRefB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12;\n\x0fprev_deployment\x18\x02 \x01(\x0b\x32\".prodvana.deployment.DeploymentRef\"V\n\x15\x43ompareDeploymentResp\x12=\n\ncomparison\x18\x01 \x01(\x0b\x32).prodvana.deployment.DeploymentComparison\"\x94\x01\n\x14PreviewDeploymentReq\x12?\n\x06\x63onfig\x18\x01 \x01(\x0b\x32%.prodvana.deployment.DeploymentConfigB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12;\n\x0fprev_deployment\x18\x02 \x01(\x0b\x32\".prodvana.deployment.DeploymentRef\"L\n\x15PreviewDeploymentResp\x12\x33\n\ndeployment\x18\x01 \x01(\x0b\x32\x1f.prodvana.deployment.Deployment\"\xe6\x01\n\x17GetLatestDeploymentsReq\x12\x36\n\x07\x66ilters\x18\x01 \x03(\x0b\x32%.prodvana.deployment.DeploymentFilter\x12\x35\n\x06\x66ilter\x18\x02 \x01(\x0b\x32%.prodvana.deployment.DeploymentFilter\x12\x35\n\x06status\x18\x03 \x01(\x0e\x32%.prodvana.deployment.DeploymentStatus\x12\x12\n\npage_token\x18\x04 \x01(\t\x12\x11\n\tpage_size\x18\x05 \x01(\x05\"i\n\x18GetLatestDeploymentsResp\x12\x34\n\x0b\x64\x65ployments\x18\x01 \x03(\x0b\x32\x1f.prodvana.deployment.Deployment\x12\x17\n\x0fnext_page_token\x18\x02 \x01(\t\"u\n\x19\x44\x65ploymentServiceInstance\x12\x1c\n\x0b\x61pplication\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x18\n\x07service\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12 \n\x0frelease_channel\x18\x03 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\"\xee\x01\n\x1a\x43heckCommitInDeploymentReq\x12 \n\rdeployment_id\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01H\x00\x12_\n\x1b\x64\x65ployment_service_instance\x18\x02 \x01(\x0b\x32..prodvana.deployment.DeploymentServiceInstanceB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01H\x00\x12\x1b\n\nrepository\x18\x03 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x17\n\x06\x63ommit\x18\x04 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x42\x17\n\x10\x64\x65ployment_oneof\x12\x03\xf8\x42\x01\"\xae\x01\n\x1b\x43heckCommitInDeploymentResp\x12G\n\x06result\x18\x01 \x01(\x0e\x32\x37.prodvana.deployment.CheckCommitInDeploymentResp.Result\"F\n\x06Result\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x0c\n\x08INCLUDED\x10\x01\x12\x0f\n\x0bNO_RELATION\x10\x02\x12\x10\n\x0cNOT_INCLUDED\x10\x03\x32\x96\t\n\x11\x44\x65ploymentManager\x12\x83\x01\n\x10RecordDeployment\x12(.prodvana.deployment.RecordDeploymentReq\x1a).prodvana.deployment.RecordDeploymentResp\"\x1a\x82\xd3\xe4\x93\x02\x14\"\x0f/v1/deployments:\x01*\x12\xae\x01\n\x16UpdateDeploymentStatus\x12..prodvana.deployment.UpdateDeploymentStatusReq\x1a/.prodvana.deployment.UpdateDeploymentStatusResp\"3\x82\xd3\xe4\x93\x02-\"(/v1/deployments/{deployment_id=*}/status:\x01*\x12}\n\x0fListDeployments\x12\'.prodvana.deployment.ListDeploymentsReq\x1a(.prodvana.deployment.ListDeploymentsResp\"\x17\x82\xd3\xe4\x93\x02\x11\x12\x0f/v1/deployments\x12n\n\x15ListDeploymentsStream\x12\'.prodvana.deployment.ListDeploymentsReq\x1a(.prodvana.deployment.ListDeploymentsResp\"\x00\x30\x01\x12\x8e\x01\n\x11\x43ompareDeployment\x12).prodvana.deployment.CompareDeploymentReq\x1a*.prodvana.deployment.CompareDeploymentResp\"\"\x82\xd3\xe4\x93\x02\x1c\"\x17/v1/deployments/compare:\x01*\x12\x8e\x01\n\x11PreviewDeployment\x12).prodvana.deployment.PreviewDeploymentReq\x1a*.prodvana.deployment.PreviewDeploymentResp\"\"\x82\xd3\xe4\x93\x02\x1c\"\x17/v1/deployments/preview:\x01*\x12\x93\x01\n\x14GetLatestDeployments\x12,.prodvana.deployment.GetLatestDeploymentsReq\x1a-.prodvana.deployment.GetLatestDeploymentsResp\"\x1e\x82\xd3\xe4\x93\x02\x18\x12\x16/v1/deployments/latest\x12\xa2\x01\n\x17\x43heckCommitInDeployment\x12/.prodvana.deployment.CheckCommitInDeploymentReq\x1a\x30.prodvana.deployment.CheckCommitInDeploymentResp\"$\x82\xd3\xe4\x93\x02\x1e\x12\x1c/v1/deployments/check_commitBOZMgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/deploymentb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'prodvana.deployment.manager_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZMgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/deployment'
  _RECORDDEPLOYMENTREQ.fields_by_name['config']._options = None
  _RECORDDEPLOYMENTREQ.fields_by_name['config']._serialized_options = b'\372B\005\212\001\002\020\001'
  _DEPLOYMENTREF.oneofs_by_name['ref']._options = None
  _DEPLOYMENTREF.oneofs_by_name['ref']._serialized_options = b'\370B\001'
  _COMPAREDEPLOYMENTREQ.fields_by_name['new_deployment']._options = None
  _COMPAREDEPLOYMENTREQ.fields_by_name['new_deployment']._serialized_options = b'\372B\005\212\001\002\020\001'
  _PREVIEWDEPLOYMENTREQ.fields_by_name['config']._options = None
  _PREVIEWDEPLOYMENTREQ.fields_by_name['config']._serialized_options = b'\372B\005\212\001\002\020\001'
  _DEPLOYMENTSERVICEINSTANCE.fields_by_name['application']._options = None
  _DEPLOYMENTSERVICEINSTANCE.fields_by_name['application']._serialized_options = b'\372B\004r\002\020\001'
  _DEPLOYMENTSERVICEINSTANCE.fields_by_name['service']._options = None
  _DEPLOYMENTSERVICEINSTANCE.fields_by_name['service']._serialized_options = b'\372B\004r\002\020\001'
  _DEPLOYMENTSERVICEINSTANCE.fields_by_name['release_channel']._options = None
  _DEPLOYMENTSERVICEINSTANCE.fields_by_name['release_channel']._serialized_options = b'\372B\004r\002\020\001'
  _CHECKCOMMITINDEPLOYMENTREQ.oneofs_by_name['deployment_oneof']._options = None
  _CHECKCOMMITINDEPLOYMENTREQ.oneofs_by_name['deployment_oneof']._serialized_options = b'\370B\001'
  _CHECKCOMMITINDEPLOYMENTREQ.fields_by_name['deployment_id']._options = None
  _CHECKCOMMITINDEPLOYMENTREQ.fields_by_name['deployment_id']._serialized_options = b'\372B\004r\002\020\001'
  _CHECKCOMMITINDEPLOYMENTREQ.fields_by_name['deployment_service_instance']._options = None
  _CHECKCOMMITINDEPLOYMENTREQ.fields_by_name['deployment_service_instance']._serialized_options = b'\372B\005\212\001\002\020\001'
  _CHECKCOMMITINDEPLOYMENTREQ.fields_by_name['repository']._options = None
  _CHECKCOMMITINDEPLOYMENTREQ.fields_by_name['repository']._serialized_options = b'\372B\004r\002\020\001'
  _CHECKCOMMITINDEPLOYMENTREQ.fields_by_name['commit']._options = None
  _CHECKCOMMITINDEPLOYMENTREQ.fields_by_name['commit']._serialized_options = b'\372B\004r\002\020\001'
  _DEPLOYMENTMANAGER.methods_by_name['RecordDeployment']._options = None
  _DEPLOYMENTMANAGER.methods_by_name['RecordDeployment']._serialized_options = b'\202\323\344\223\002\024\"\017/v1/deployments:\001*'
  _DEPLOYMENTMANAGER.methods_by_name['UpdateDeploymentStatus']._options = None
  _DEPLOYMENTMANAGER.methods_by_name['UpdateDeploymentStatus']._serialized_options = b'\202\323\344\223\002-\"(/v1/deployments/{deployment_id=*}/status:\001*'
  _DEPLOYMENTMANAGER.methods_by_name['ListDeployments']._options = None
  _DEPLOYMENTMANAGER.methods_by_name['ListDeployments']._serialized_options = b'\202\323\344\223\002\021\022\017/v1/deployments'
  _DEPLOYMENTMANAGER.methods_by_name['CompareDeployment']._options = None
  _DEPLOYMENTMANAGER.methods_by_name['CompareDeployment']._serialized_options = b'\202\323\344\223\002\034\"\027/v1/deployments/compare:\001*'
  _DEPLOYMENTMANAGER.methods_by_name['PreviewDeployment']._options = None
  _DEPLOYMENTMANAGER.methods_by_name['PreviewDeployment']._serialized_options = b'\202\323\344\223\002\034\"\027/v1/deployments/preview:\001*'
  _DEPLOYMENTMANAGER.methods_by_name['GetLatestDeployments']._options = None
  _DEPLOYMENTMANAGER.methods_by_name['GetLatestDeployments']._serialized_options = b'\202\323\344\223\002\030\022\026/v1/deployments/latest'
  _DEPLOYMENTMANAGER.methods_by_name['CheckCommitInDeployment']._options = None
  _DEPLOYMENTMANAGER.methods_by_name['CheckCommitInDeployment']._serialized_options = b'\202\323\344\223\002\036\022\034/v1/deployments/check_commit'
  _globals['_RECORDDEPLOYMENTREQ']._serialized_start=175
  _globals['_RECORDDEPLOYMENTREQ']._serialized_end=278
  _globals['_RECORDDEPLOYMENTRESP']._serialized_start=280
  _globals['_RECORDDEPLOYMENTRESP']._serialized_end=345
  _globals['_UPDATEDEPLOYMENTSTATUSREQ']._serialized_start=347
  _globals['_UPDATEDEPLOYMENTSTATUSREQ']._serialized_end=452
  _globals['_UPDATEDEPLOYMENTSTATUSRESP']._serialized_start=454
  _globals['_UPDATEDEPLOYMENTSTATUSRESP']._serialized_end=537
  _globals['_DEPLOYMENTFILTER']._serialized_start=539
  _globals['_DEPLOYMENTFILTER']._serialized_end=648
  _globals['_LISTDEPLOYMENTSREQ']._serialized_start=651
  _globals['_LISTDEPLOYMENTSREQ']._serialized_end=883
  _globals['_LISTDEPLOYMENTSRESP']._serialized_start=885
  _globals['_LISTDEPLOYMENTSRESP']._serialized_end=985
  _globals['_DEPLOYMENTREF']._serialized_start=987
  _globals['_DEPLOYMENTREF']._serialized_end=1096
  _globals['_COMPAREDEPLOYMENTREQ']._serialized_start=1099
  _globals['_COMPAREDEPLOYMENTREQ']._serialized_end=1252
  _globals['_COMPAREDEPLOYMENTRESP']._serialized_start=1254
  _globals['_COMPAREDEPLOYMENTRESP']._serialized_end=1340
  _globals['_PREVIEWDEPLOYMENTREQ']._serialized_start=1343
  _globals['_PREVIEWDEPLOYMENTREQ']._serialized_end=1491
  _globals['_PREVIEWDEPLOYMENTRESP']._serialized_start=1493
  _globals['_PREVIEWDEPLOYMENTRESP']._serialized_end=1569
  _globals['_GETLATESTDEPLOYMENTSREQ']._serialized_start=1572
  _globals['_GETLATESTDEPLOYMENTSREQ']._serialized_end=1802
  _globals['_GETLATESTDEPLOYMENTSRESP']._serialized_start=1804
  _globals['_GETLATESTDEPLOYMENTSRESP']._serialized_end=1909
  _globals['_DEPLOYMENTSERVICEINSTANCE']._serialized_start=1911
  _globals['_DEPLOYMENTSERVICEINSTANCE']._serialized_end=2028
  _globals['_CHECKCOMMITINDEPLOYMENTREQ']._serialized_start=2031
  _globals['_CHECKCOMMITINDEPLOYMENTREQ']._serialized_end=2269
  _globals['_CHECKCOMMITINDEPLOYMENTRESP']._serialized_start=2272
  _globals['_CHECKCOMMITINDEPLOYMENTRESP']._serialized_end=2446
  _globals['_CHECKCOMMITINDEPLOYMENTRESP_RESULT']._serialized_start=2376
  _globals['_CHECKCOMMITINDEPLOYMENTRESP_RESULT']._serialized_end=2446
  _globals['_DEPLOYMENTMANAGER']._serialized_start=2449
  _globals['_DEPLOYMENTMANAGER']._serialized_end=3623
# @@protoc_insertion_point(module_scope)