# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/impact_analysis/impact_analysis.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n.prodvana/impact_analysis/impact_analysis.proto\x12\x18prodvana.impact_analysis\"\x90\x02\n\x14ImpactAnalysisResult\x12:\n\x08\x63\x61tegory\x18\x01 \x01(\x0e\x32(.prodvana.impact_analysis.ImpactCategory\x12=\n\rincident_type\x18\x02 \x01(\x0e\x32&.prodvana.impact_analysis.IncidentType\x12\x14\n\x0cimpact_score\x18\x03 \x01(\x01\x12\x18\n\x10impact_reasoning\x18\x04 \x01(\t\x12\x18\n\x10likelihood_score\x18\x05 \x01(\x01\x12\x1c\n\x14likelihood_reasoning\x18\x06 \x01(\t\x12\x15\n\roverall_score\x18\x07 \x01(\x01*\x8a\x01\n\x0eImpactCategory\x12\x14\n\x10UNKNOWN_CATEGORY\x10\x00\x12\x0b\n\x07\x42\x41\x43KEND\x10\x01\x12\x0c\n\x08\x46RONTEND\x10\x02\x12\x12\n\x0e\x44\x41TA_INTEGRITY\x10\x03\x12\x0c\n\x08SECURITY\x10\x04\x12\x08\n\x04\x43OST\x10\x05\x12\x1b\n\x17PRODVANA_INTERNAL_ERROR\x10\x06*b\n\x0cIncidentType\x12\x14\n\x10UNKNOWN_INCIDENT\x10\x00\x12\n\n\x06OUTAGE\x10\x01\x12\x0b\n\x07\x45XPENSE\x10\x02\x12\x13\n\x0fSECURITY_BREACH\x10\x03\x12\x0e\n\nUSER_ISSUE\x10\x04\x42TZRgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/impact_analysisb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'prodvana.impact_analysis.impact_analysis_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZRgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/impact_analysis'
  _globals['_IMPACTCATEGORY']._serialized_start=352
  _globals['_IMPACTCATEGORY']._serialized_end=490
  _globals['_INCIDENTTYPE']._serialized_start=492
  _globals['_INCIDENTTYPE']._serialized_end=590
  _globals['_IMPACTANALYSISRESULT']._serialized_start=77
  _globals['_IMPACTANALYSISRESULT']._serialized_end=349
# @@protoc_insertion_point(module_scope)
