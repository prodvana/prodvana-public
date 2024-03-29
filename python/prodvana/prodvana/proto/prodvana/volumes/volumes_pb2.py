# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/volumes/volumes.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from prodvana.proto.validate import validate_pb2 as validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1eprodvana/volumes/volumes.proto\x12\x10prodvana.volumes\x1a\x17validate/validate.proto\"\xd5\x02\n\x0cVolumeSource\x12=\n\x06secret\x18\x01 \x01(\x0b\x32+.prodvana.volumes.VolumeSource.SecretSourceH\x00\x12\x43\n\tephemeral\x18\x02 \x01(\x0b\x32..prodvana.volumes.VolumeSource.EphemeralSourceH\x00\x1a,\n\x0cSecretSource\x12\x1c\n\x0bsecret_name\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x1a\x88\x01\n\x0f\x45phemeralSource\x12\x45\n\x06medium\x18\x01 \x01(\x0e\x32\x35.prodvana.volumes.VolumeSource.EphemeralSource.Medium\".\n\x06Medium\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x0b\n\x07\x44\x45\x46\x41ULT\x10\x01\x12\n\n\x06MEMORY\x10\x02\x42\x08\n\x06source\"4\n\x0bVolumeMount\x12\x12\n\nmount_path\x18\x01 \x01(\t\x12\x11\n\tread_only\x18\x02 \x01(\x08\"}\n\x06Volume\x12\x15\n\x04name\x18\x01 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x01\x12.\n\x06source\x18\x02 \x01(\x0b\x32\x1e.prodvana.volumes.VolumeSource\x12,\n\x05mount\x18\x03 \x01(\x0b\x32\x1d.prodvana.volumes.VolumeMountBLZJgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/volumesb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'prodvana.volumes.volumes_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZJgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/volumes'
  _VOLUMESOURCE_SECRETSOURCE.fields_by_name['secret_name']._options = None
  _VOLUMESOURCE_SECRETSOURCE.fields_by_name['secret_name']._serialized_options = b'\372B\004r\002\020\001'
  _VOLUME.fields_by_name['name']._options = None
  _VOLUME.fields_by_name['name']._serialized_options = b'\372B\004r\002\020\001'
  _globals['_VOLUMESOURCE']._serialized_start=78
  _globals['_VOLUMESOURCE']._serialized_end=419
  _globals['_VOLUMESOURCE_SECRETSOURCE']._serialized_start=226
  _globals['_VOLUMESOURCE_SECRETSOURCE']._serialized_end=270
  _globals['_VOLUMESOURCE_EPHEMERALSOURCE']._serialized_start=273
  _globals['_VOLUMESOURCE_EPHEMERALSOURCE']._serialized_end=409
  _globals['_VOLUMESOURCE_EPHEMERALSOURCE_MEDIUM']._serialized_start=363
  _globals['_VOLUMESOURCE_EPHEMERALSOURCE_MEDIUM']._serialized_end=409
  _globals['_VOLUMEMOUNT']._serialized_start=421
  _globals['_VOLUMEMOUNT']._serialized_end=473
  _globals['_VOLUME']._serialized_start=475
  _globals['_VOLUME']._serialized_end=600
# @@protoc_insertion_point(module_scope)
