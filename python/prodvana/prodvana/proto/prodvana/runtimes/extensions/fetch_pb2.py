# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/runtimes/extensions/fetch.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from prodvana.proto.prodvana.common_config import external_link_pb2 as prodvana_dot_common__config_dot_external__link__pb2
from prodvana.proto.prodvana.runtimes import debug_event_pb2 as prodvana_dot_runtimes_dot_debug__event__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n(prodvana/runtimes/extensions/fetch.proto\x12\x1cprodvana.runtimes.extensions\x1a*prodvana/common_config/external_link.proto\x1a#prodvana/runtimes/debug_event.proto\"\x90\x01\n\x15\x45xternalObjectVersion\x12\x0f\n\x07version\x18\x01 \x01(\t\x12\x10\n\x08replicas\x18\x02 \x01(\x05\x12\x1a\n\x12\x61vailable_replicas\x18\x05 \x01(\x05\x12\x17\n\x0ftarget_replicas\x18\x06 \x01(\x05\x12\x0e\n\x06\x61\x63tive\x18\x03 \x01(\x08\x12\x0f\n\x07\x64rifted\x18\x04 \x01(\x08\"\xf5\x02\n\x0e\x45xternalObject\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x13\n\x0bobject_type\x18\x02 \x01(\t\x12\x45\n\x08versions\x18\x03 \x03(\x0b\x32\x33.prodvana.runtimes.extensions.ExternalObjectVersion\x12\x43\n\x06status\x18\x04 \x01(\x0e\x32\x33.prodvana.runtimes.extensions.ExternalObject.Status\x12<\n\x0e\x65xternal_links\x18\x05 \x03(\x0b\x32$.prodvana.common_config.ExternalLink\x12\x0f\n\x07message\x18\x06 \x01(\t\x12\x33\n\x0c\x64\x65\x62ug_events\x18\x07 \x03(\x0b\x32\x1d.prodvana.runtimes.DebugEvent\"0\n\x06Status\x12\x0b\n\x07PENDING\x10\x00\x12\r\n\tSUCCEEDED\x10\x01\x12\n\n\x06\x46\x41ILED\x10\x02\"L\n\x0b\x46\x65tchOutput\x12=\n\x07objects\x18\x01 \x03(\x0b\x32,.prodvana.runtimes.extensions.ExternalObject*I\n\tFetchMode\x12\x16\n\x12UNKNOWN_FETCH_MODE\x10\x00\x12\r\n\tEXIT_CODE\x10\x01\x12\x15\n\x11STRUCTURED_OUTPUT\x10\x02\x42XZVgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/runtimes/extensionsb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'prodvana.runtimes.extensions.fetch_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZVgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/runtimes/extensions'
  _globals['_FETCHMODE']._serialized_start=756
  _globals['_FETCHMODE']._serialized_end=829
  _globals['_EXTERNALOBJECTVERSION']._serialized_start=156
  _globals['_EXTERNALOBJECTVERSION']._serialized_end=300
  _globals['_EXTERNALOBJECT']._serialized_start=303
  _globals['_EXTERNALOBJECT']._serialized_end=676
  _globals['_EXTERNALOBJECT_STATUS']._serialized_start=628
  _globals['_EXTERNALOBJECT_STATUS']._serialized_end=676
  _globals['_FETCHOUTPUT']._serialized_start=678
  _globals['_FETCHOUTPUT']._serialized_end=754
# @@protoc_insertion_point(module_scope)
