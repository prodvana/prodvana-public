# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/pvn_wrapper/job_manager.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from prodvana.proto.prodvana.pvn_wrapper import output_pb2 as prodvana_dot_pvn__wrapper_dot_output__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n&prodvana/pvn_wrapper/job_manager.proto\x12\x14prodvana.pvn_wrapper\x1a!prodvana/pvn_wrapper/output.proto\"R\n\x12ReportJobResultReq\x12\x0e\n\x06job_id\x18\x01 \x01(\t\x12,\n\x06output\x18\x02 \x01(\x0b\x32\x1c.prodvana.pvn_wrapper.Output\"\x15\n\x13ReportJobResultResp2t\n\nJobManager\x12\x66\n\x0fReportJobResult\x12(.prodvana.pvn_wrapper.ReportJobResultReq\x1a).prodvana.pvn_wrapper.ReportJobResultRespBPZNgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/pvn_wrapperb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'prodvana.pvn_wrapper.job_manager_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZNgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/pvn_wrapper'
  _globals['_REPORTJOBRESULTREQ']._serialized_start=99
  _globals['_REPORTJOBRESULTREQ']._serialized_end=181
  _globals['_REPORTJOBRESULTRESP']._serialized_start=183
  _globals['_REPORTJOBRESULTRESP']._serialized_end=204
  _globals['_JOBMANAGER']._serialized_start=206
  _globals['_JOBMANAGER']._serialized_end=322
# @@protoc_insertion_point(module_scope)
