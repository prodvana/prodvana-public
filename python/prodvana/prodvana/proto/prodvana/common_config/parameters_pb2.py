# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/common_config/parameters.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from prodvana.proto.prodvana.common_config import program_pb2 as prodvana_dot_common__config_dot_program__pb2
from prodvana.proto.prodvana.common_config import env_pb2 as prodvana_dot_common__config_dot_env__pb2
from prodvana.proto.validate import validate_pb2 as validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\'prodvana/common_config/parameters.proto\x12\x16prodvana.common_config\x1a$prodvana/common_config/program.proto\x1a prodvana/common_config/env.proto\x1a\x17validate/validate.proto\"&\n\rProgramChange\x12\x15\n\x04name\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\"2\n\x19StringParameterDefinition\x12\x15\n\rdefault_value\x18\x01 \x01(\t\"\xba\x02\n\x1e\x44ockerImageParameterDefinition\x12\x13\n\x0b\x64\x65\x66\x61ult_tag\x18\x01 \x01(\t\x12P\n\x13image_registry_info\x18\x02 \x01(\x0b\x32).prodvana.common_config.ImageRegistryInfoB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12^\n\x07\x63hanges\x18\x03 \x03(\x0b\x32>.prodvana.common_config.DockerImageParameterDefinition.ChangesB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\x1aQ\n\x07\x43hanges\x12\x38\n\x07program\x18\x01 \x01(\x0b\x32%.prodvana.common_config.ProgramChangeH\x00\x42\x0c\n\x05oneof\x12\x03\xf8\x42\x01\"\x14\n\x12\x46ixedReplicaChange\"\xe5\x01\n\x16IntParameterDefinition\x12\x15\n\rdefault_value\x18\x01 \x01(\x03\x12V\n\x07\x63hanges\x18\x02 \x03(\x0b\x32\x36.prodvana.common_config.IntParameterDefinition.ChangesB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\x1a\\\n\x07\x43hanges\x12\x43\n\rfixed_replica\x18\x01 \x01(\x0b\x32*.prodvana.common_config.FixedReplicaChangeH\x00\x42\x0c\n\x05oneof\x12\x03\xf8\x42\x01\"\x1b\n\x19SecretParameterDefinition\"H\n\x19\x43ommitParameterDefinition\x12\x1b\n\nrepository\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x0e\n\x06\x62ranch\x18\x02 \x01(\t\"0\n\x17\x42lobParameterDefinition\x12\x15\n\rdefault_value\x18\x01 \x01(\t\"\x87\x04\n\x13ParameterDefinition\x12\x15\n\x04name\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x13\n\x0b\x64\x65scription\x18\x02 \x01(\t\x12\x43\n\x06string\x18\x03 \x01(\x0b\x32\x31.prodvana.common_config.StringParameterDefinitionH\x00\x12N\n\x0c\x64ocker_image\x18\x04 \x01(\x0b\x32\x36.prodvana.common_config.DockerImageParameterDefinitionH\x00\x12=\n\x03int\x18\x05 \x01(\x0b\x32..prodvana.common_config.IntParameterDefinitionH\x00\x12\x43\n\x06secret\x18\x06 \x01(\x0b\x32\x31.prodvana.common_config.SecretParameterDefinitionH\x00\x12\x43\n\x06\x63ommit\x18\x08 \x01(\x0b\x32\x31.prodvana.common_config.CommitParameterDefinitionH\x00\x12?\n\x04\x62lob\x18\t \x01(\x0b\x32/.prodvana.common_config.BlobParameterDefinitionH\x00\x12\x10\n\x08required\x18\x07 \x01(\x08\x42\x13\n\x0c\x63onfig_oneof\x12\x03\xf8\x42\x01\"w\n\x14SecretParameterValue\x12\x14\n\nraw_secret\x18\x01 \x01(\tH\x00\x12\x34\n\nsecret_ref\x18\x02 \x01(\x0b\x32\x1e.prodvana.common_config.SecretH\x00\x42\x13\n\x0csecret_oneof\x12\x03\xf8\x42\x01\"S\n\x12\x42lobParameterValue\x12\x11\n\x07inlined\x18\x01 \x01(\tH\x00\x12\x17\n\rinlined_bytes\x18\x02 \x01(\x0cH\x00\x42\x11\n\nblob_oneof\x12\x03\xf8\x42\x01\"\xd6\x02\n\x0eParameterValue\x12\x15\n\x04name\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12\x10\n\x06string\x18\x02 \x01(\tH\x00\x12\r\n\x03int\x18\x04 \x01(\x03H\x00\x12V\n\x10\x64ocker_image_tag\x18\x05 \x01(\tB:\xfa\x42\x37r523^[A-Za-z0-9_][0-9A-Za-z_.-]*$|^sha256:[0-9a-f]{64}$H\x00\x12>\n\x06secret\x18\x06 \x01(\x0b\x32,.prodvana.common_config.SecretParameterValueH\x00\x12\x10\n\x06\x63ommit\x18\x07 \x01(\tH\x00\x12:\n\x04\x62lob\x18\x08 \x01(\x0b\x32*.prodvana.common_config.BlobParameterValueH\x00\x42\x12\n\x0bvalue_oneof\x12\x03\xf8\x42\x01J\x04\x08\x03\x10\x04R\x0c\x64ocker_image\"b\n\x10ParametersConfig\x12N\n\nparameters\x18\x01 \x03(\x0b\x32+.prodvana.common_config.ParameterDefinitionB\r\xfa\x42\n\x92\x01\x07\"\x05\x8a\x01\x02\x10\x01\x42RZPgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_configb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'prodvana.common_config.parameters_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZPgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config'
  _PROGRAMCHANGE.fields_by_name['name']._options = None
  _PROGRAMCHANGE.fields_by_name['name']._serialized_options = b'\372B\004r\002\020\001'
  _DOCKERIMAGEPARAMETERDEFINITION_CHANGES.oneofs_by_name['oneof']._options = None
  _DOCKERIMAGEPARAMETERDEFINITION_CHANGES.oneofs_by_name['oneof']._serialized_options = b'\370B\001'
  _DOCKERIMAGEPARAMETERDEFINITION.fields_by_name['image_registry_info']._options = None
  _DOCKERIMAGEPARAMETERDEFINITION.fields_by_name['image_registry_info']._serialized_options = b'\372B\005\212\001\002\020\001'
  _DOCKERIMAGEPARAMETERDEFINITION.fields_by_name['changes']._options = None
  _DOCKERIMAGEPARAMETERDEFINITION.fields_by_name['changes']._serialized_options = b'\372B\n\222\001\007\"\005\212\001\002\020\001'
  _INTPARAMETERDEFINITION_CHANGES.oneofs_by_name['oneof']._options = None
  _INTPARAMETERDEFINITION_CHANGES.oneofs_by_name['oneof']._serialized_options = b'\370B\001'
  _INTPARAMETERDEFINITION.fields_by_name['changes']._options = None
  _INTPARAMETERDEFINITION.fields_by_name['changes']._serialized_options = b'\372B\n\222\001\007\"\005\212\001\002\020\001'
  _COMMITPARAMETERDEFINITION.fields_by_name['repository']._options = None
  _COMMITPARAMETERDEFINITION.fields_by_name['repository']._serialized_options = b'\372B\004r\002\020\001'
  _PARAMETERDEFINITION.oneofs_by_name['config_oneof']._options = None
  _PARAMETERDEFINITION.oneofs_by_name['config_oneof']._serialized_options = b'\370B\001'
  _PARAMETERDEFINITION.fields_by_name['name']._options = None
  _PARAMETERDEFINITION.fields_by_name['name']._serialized_options = b'\372B\004r\002\020\001'
  _SECRETPARAMETERVALUE.oneofs_by_name['secret_oneof']._options = None
  _SECRETPARAMETERVALUE.oneofs_by_name['secret_oneof']._serialized_options = b'\370B\001'
  _BLOBPARAMETERVALUE.oneofs_by_name['blob_oneof']._options = None
  _BLOBPARAMETERVALUE.oneofs_by_name['blob_oneof']._serialized_options = b'\370B\001'
  _PARAMETERVALUE.oneofs_by_name['value_oneof']._options = None
  _PARAMETERVALUE.oneofs_by_name['value_oneof']._serialized_options = b'\370B\001'
  _PARAMETERVALUE.fields_by_name['name']._options = None
  _PARAMETERVALUE.fields_by_name['name']._serialized_options = b'\372B\004r\002\020\001'
  _PARAMETERVALUE.fields_by_name['docker_image_tag']._options = None
  _PARAMETERVALUE.fields_by_name['docker_image_tag']._serialized_options = b'\372B7r523^[A-Za-z0-9_][0-9A-Za-z_.-]*$|^sha256:[0-9a-f]{64}$'
  _PARAMETERSCONFIG.fields_by_name['parameters']._options = None
  _PARAMETERSCONFIG.fields_by_name['parameters']._serialized_options = b'\372B\n\222\001\007\"\005\212\001\002\020\001'
  _globals['_PROGRAMCHANGE']._serialized_start=164
  _globals['_PROGRAMCHANGE']._serialized_end=202
  _globals['_STRINGPARAMETERDEFINITION']._serialized_start=204
  _globals['_STRINGPARAMETERDEFINITION']._serialized_end=254
  _globals['_DOCKERIMAGEPARAMETERDEFINITION']._serialized_start=257
  _globals['_DOCKERIMAGEPARAMETERDEFINITION']._serialized_end=571
  _globals['_DOCKERIMAGEPARAMETERDEFINITION_CHANGES']._serialized_start=490
  _globals['_DOCKERIMAGEPARAMETERDEFINITION_CHANGES']._serialized_end=571
  _globals['_FIXEDREPLICACHANGE']._serialized_start=573
  _globals['_FIXEDREPLICACHANGE']._serialized_end=593
  _globals['_INTPARAMETERDEFINITION']._serialized_start=596
  _globals['_INTPARAMETERDEFINITION']._serialized_end=825
  _globals['_INTPARAMETERDEFINITION_CHANGES']._serialized_start=733
  _globals['_INTPARAMETERDEFINITION_CHANGES']._serialized_end=825
  _globals['_SECRETPARAMETERDEFINITION']._serialized_start=827
  _globals['_SECRETPARAMETERDEFINITION']._serialized_end=854
  _globals['_COMMITPARAMETERDEFINITION']._serialized_start=856
  _globals['_COMMITPARAMETERDEFINITION']._serialized_end=928
  _globals['_BLOBPARAMETERDEFINITION']._serialized_start=930
  _globals['_BLOBPARAMETERDEFINITION']._serialized_end=978
  _globals['_PARAMETERDEFINITION']._serialized_start=981
  _globals['_PARAMETERDEFINITION']._serialized_end=1500
  _globals['_SECRETPARAMETERVALUE']._serialized_start=1502
  _globals['_SECRETPARAMETERVALUE']._serialized_end=1621
  _globals['_BLOBPARAMETERVALUE']._serialized_start=1623
  _globals['_BLOBPARAMETERVALUE']._serialized_end=1706
  _globals['_PARAMETERVALUE']._serialized_start=1709
  _globals['_PARAMETERVALUE']._serialized_end=2051
  _globals['_PARAMETERSCONFIG']._serialized_start=2053
  _globals['_PARAMETERSCONFIG']._serialized_end=2151
# @@protoc_insertion_point(module_scope)
