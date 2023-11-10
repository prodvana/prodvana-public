# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/release/manager.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from prodvana.proto.prodvana.object import meta_pb2 as prodvana_dot_object_dot_meta__pb2
from prodvana.proto.prodvana.release import object_pb2 as prodvana_dot_release_dot_object__pb2
from prodvana.proto.validate import validate_pb2 as validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1eprodvana/release/manager.proto\x12\x10prodvana.release\x1a\x1cgoogle/api/annotations.proto\x1a\x1aprodvana/object/meta.proto\x1a\x1dprodvana/release/object.proto\x1a\x17validate/validate.proto\"^\n\x10RecordReleaseReq\x12\x39\n\x06\x63onfig\x18\x01 \x01(\x0b\x32\x1f.prodvana.release.ReleaseConfigB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12\x0f\n\x07pending\x18\x02 \x01(\x08\">\n\x11RecordReleaseResp\x12)\n\x04meta\x18\x01 \x01(\x0b\x32\x1b.prodvana.object.ObjectMeta\"]\n\x16UpdateReleaseStatusReq\x12\x12\n\nrelease_id\x18\x01 \x01(\t\x12/\n\x06status\x18\x02 \x01(\x0e\x32\x1f.prodvana.release.ReleaseStatus\"J\n\x17UpdateReleaseStatusResp\x12/\n\x06status\x18\x01 \x01(\x0e\x32\x1f.prodvana.release.ReleaseStatus\"\xca\x02\n\x0fListReleasesReq\x12\x39\n\x07\x66ilters\x18\x01 \x03(\x0b\x32(.prodvana.release.ListReleasesReq.Filter\x12\x38\n\x06\x66ilter\x18\x02 \x01(\x0b\x32(.prodvana.release.ListReleasesReq.Filter\x12\x1b\n\x13starting_release_id\x18\x03 \x01(\t\x12\x19\n\x11\x65nding_release_id\x18\x04 \x01(\t\x12\x12\n\npage_token\x18\x05 \x01(\t\x12\x11\n\tpage_size\x18\x06 \x01(\x05\x1a\x63\n\x06\x46ilter\x12\x10\n\x08services\x18\x01 \x03(\t\x12\x18\n\x10release_channels\x18\x02 \x03(\t\x12\x13\n\x0b\x61pplication\x18\x03 \x01(\t\x12\x18\n\x10\x64\x65sired_state_id\x18\x04 \x01(\t\"X\n\x10ListReleasesResp\x12+\n\x08releases\x18\x01 \x03(\x0b\x32\x19.prodvana.release.Release\x12\x17\n\x0fnext_page_token\x18\x02 \x01(\t\"\x95\x02\n\x11\x43ompareReleaseReq\x12M\n\x0bnew_release\x18\x01 \x01(\x0b\x32..prodvana.release.CompareReleaseReq.ReleaseRefB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12N\n\x0cprev_release\x18\x02 \x01(\x0b\x32..prodvana.release.CompareReleaseReq.ReleaseRefB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x1a\x61\n\nReleaseRef\x12\x14\n\nrelease_id\x18\x01 \x01(\tH\x00\x12\x31\n\x06\x63onfig\x18\x02 \x01(\x0b\x32\x1f.prodvana.release.ReleaseConfigH\x00\x42\n\n\x03ref\x12\x03\xf8\x42\x01\"M\n\x12\x43ompareReleaseResp\x12\x37\n\ncomparison\x18\x01 \x01(\x0b\x32#.prodvana.release.ReleaseComparison2\xeb\x04\n\x0eReleaseManager\x12q\n\rRecordRelease\x12\".prodvana.release.RecordReleaseReq\x1a#.prodvana.release.RecordReleaseResp\"\x17\x82\xd3\xe4\x93\x02\x11\"\x0c/v1/releases:\x01*\x12\x99\x01\n\x13UpdateReleaseStatus\x12(.prodvana.release.UpdateReleaseStatusReq\x1a).prodvana.release.UpdateReleaseStatusResp\"-\x82\xd3\xe4\x93\x02\'\"\"/v1/releases/{release_id=*}/status:\x01*\x12k\n\x0cListReleases\x12!.prodvana.release.ListReleasesReq\x1a\".prodvana.release.ListReleasesResp\"\x14\x82\xd3\xe4\x93\x02\x0e\x12\x0c/v1/releases\x12_\n\x12ListReleasesStream\x12!.prodvana.release.ListReleasesReq\x1a\".prodvana.release.ListReleasesResp\"\x00\x30\x01\x12|\n\x0e\x43ompareRelease\x12#.prodvana.release.CompareReleaseReq\x1a$.prodvana.release.CompareReleaseResp\"\x1f\x82\xd3\xe4\x93\x02\x19\"\x14/v1/releases/compare:\x01*BLZJgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/releaseb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'prodvana.release.manager_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZJgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/release'
  _RECORDRELEASEREQ.fields_by_name['config']._options = None
  _RECORDRELEASEREQ.fields_by_name['config']._serialized_options = b'\372B\005\212\001\002\020\001'
  _COMPARERELEASEREQ_RELEASEREF.oneofs_by_name['ref']._options = None
  _COMPARERELEASEREQ_RELEASEREF.oneofs_by_name['ref']._serialized_options = b'\370B\001'
  _COMPARERELEASEREQ.fields_by_name['new_release']._options = None
  _COMPARERELEASEREQ.fields_by_name['new_release']._serialized_options = b'\372B\005\212\001\002\020\001'
  _COMPARERELEASEREQ.fields_by_name['prev_release']._options = None
  _COMPARERELEASEREQ.fields_by_name['prev_release']._serialized_options = b'\372B\005\212\001\002\020\001'
  _RELEASEMANAGER.methods_by_name['RecordRelease']._options = None
  _RELEASEMANAGER.methods_by_name['RecordRelease']._serialized_options = b'\202\323\344\223\002\021\"\014/v1/releases:\001*'
  _RELEASEMANAGER.methods_by_name['UpdateReleaseStatus']._options = None
  _RELEASEMANAGER.methods_by_name['UpdateReleaseStatus']._serialized_options = b'\202\323\344\223\002\'\"\"/v1/releases/{release_id=*}/status:\001*'
  _RELEASEMANAGER.methods_by_name['ListReleases']._options = None
  _RELEASEMANAGER.methods_by_name['ListReleases']._serialized_options = b'\202\323\344\223\002\016\022\014/v1/releases'
  _RELEASEMANAGER.methods_by_name['CompareRelease']._options = None
  _RELEASEMANAGER.methods_by_name['CompareRelease']._serialized_options = b'\202\323\344\223\002\031\"\024/v1/releases/compare:\001*'
  _globals['_RECORDRELEASEREQ']._serialized_start=166
  _globals['_RECORDRELEASEREQ']._serialized_end=260
  _globals['_RECORDRELEASERESP']._serialized_start=262
  _globals['_RECORDRELEASERESP']._serialized_end=324
  _globals['_UPDATERELEASESTATUSREQ']._serialized_start=326
  _globals['_UPDATERELEASESTATUSREQ']._serialized_end=419
  _globals['_UPDATERELEASESTATUSRESP']._serialized_start=421
  _globals['_UPDATERELEASESTATUSRESP']._serialized_end=495
  _globals['_LISTRELEASESREQ']._serialized_start=498
  _globals['_LISTRELEASESREQ']._serialized_end=828
  _globals['_LISTRELEASESREQ_FILTER']._serialized_start=729
  _globals['_LISTRELEASESREQ_FILTER']._serialized_end=828
  _globals['_LISTRELEASESRESP']._serialized_start=830
  _globals['_LISTRELEASESRESP']._serialized_end=918
  _globals['_COMPARERELEASEREQ']._serialized_start=921
  _globals['_COMPARERELEASEREQ']._serialized_end=1198
  _globals['_COMPARERELEASEREQ_RELEASEREF']._serialized_start=1101
  _globals['_COMPARERELEASEREQ_RELEASEREF']._serialized_end=1198
  _globals['_COMPARERELEASERESP']._serialized_start=1200
  _globals['_COMPARERELEASERESP']._serialized_end=1277
  _globals['_RELEASEMANAGER']._serialized_start=1280
  _globals['_RELEASEMANAGER']._serialized_end=1899
# @@protoc_insertion_point(module_scope)
