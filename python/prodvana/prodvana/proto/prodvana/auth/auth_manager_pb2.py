# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/auth/auth_manager.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from prodvana.proto.prodvana.auth import auth_pb2 as prodvana_dot_auth_dot_auth__pb2
from prodvana.proto.prodvana.users import users_pb2 as prodvana_dot_users_dot_users__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n prodvana/auth/auth_manager.proto\x12\rprodvana.auth\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x18prodvana/auth/auth.proto\x1a\x1aprodvana/users/users.proto\"g\n\x0fGetTokenRequest\x12\x0c\n\x04\x63ode\x18\x01 \x01(\t\x12\x14\n\x0credirect_url\x18\x02 \x01(\t\x12\x0f\n\x07\x66or_cli\x18\x03 \x01(\x08\x12\x1f\n\x17prodvana_only_admin_org\x18\x04 \x01(\x08\"&\n\x12GetAuthUrlResponse\x12\x10\n\x08\x61uth_url\x18\x01 \x01(\t\"\x8a\x01\n\x11GetAuthUrlRequest\x12\r\n\x05state\x18\x02 \x01(\t\x12\x0f\n\x07\x66or_cli\x18\x03 \x01(\x08\x12\x12\n\ninvitation\x18\x04 \x01(\t\x12\x1f\n\x17prodvana_only_admin_org\x18\x06 \x01(\x08J\x04\x08\x01\x10\x02J\x04\x08\x05\x10\x06R\x0credirect_urlR\x06no_org\"\x0c\n\nCliAuthReq\"\xe2\x01\n\x0b\x43liAuthResp\x12<\n\x0buser_prompt\x18\x01 \x01(\x0b\x32%.prodvana.auth.CliAuthResp.UserPromptH\x00\x12)\n\x05token\x18\x02 \x01(\x0b\x32\x18.prodvana.auth.AuthTokenH\x00\x1a\\\n\nUserPrompt\x12\x11\n\tuser_code\x18\x01 \x01(\t\x12\x18\n\x10verification_uri\x18\x02 \x01(\t\x12!\n\x19verification_uri_complete\x18\x03 \x01(\tB\x0c\n\nresp_oneof\"l\n\x0fRefreshTokenReq\x12\'\n\x05token\x18\x01 \x01(\x0b\x32\x18.prodvana.auth.AuthToken\x12\x0f\n\x07\x66or_cli\x18\x02 \x01(\x08\x12\x1f\n\x17prodvana_only_admin_org\x18\x03 \x01(\x08\"_\n\x10GetAuthTokenResp\x12\'\n\x05token\x18\x01 \x01(\x0b\x32\x18.prodvana.auth.AuthToken\x12\"\n\x04user\x18\x02 \x01(\x0b\x32\x14.prodvana.users.User\"\x07\n\x05\x45mpty\"S\n\tLogoutReq\x12\x0f\n\x07\x66or_cli\x18\x01 \x01(\x08\x12\x14\n\x0credirect_url\x18\x02 \x01(\t\x12\x1f\n\x17prodvana_only_admin_org\x18\x03 \x01(\x08\" \n\nLogoutResp\x12\x12\n\nlogout_url\x18\x01 \x01(\t\"p\n\x14\x43reateOrgApiTokenReq\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x02 \x01(\t\x12\x35\n\x11\x65xpires_timestamp\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"U\n\x15\x43reateOrgApiTokenResp\x12\x11\n\tapi_token\x18\x01 \x01(\t\x12)\n\x04info\x18\x02 \x01(\x0b\x32\x1b.prodvana.auth.ApiTokenInfo\"$\n\x14\x44\x65leteOrgApiTokenReq\x12\x0c\n\x04name\x18\x01 \x01(\t\"\x17\n\x15\x44\x65leteOrgApiTokenResp\"\x15\n\x13ListOrgApiTokensReq\"C\n\x14ListOrgApiTokensResp\x12+\n\x06tokens\x18\x01 \x03(\x0b\x32\x1b.prodvana.auth.ApiTokenInfo2\xab\x03\n\x0b\x41uthManager\x12k\n\x0cGetAuthToken\x12\x1e.prodvana.auth.GetTokenRequest\x1a\x1f.prodvana.auth.GetAuthTokenResp\"\x1a\x82\xd3\xe4\x93\x02\x14\x12\x12/v1/auth/get_token\x12p\n\nGetAuthUrl\x12 .prodvana.auth.GetAuthUrlRequest\x1a!.prodvana.auth.GetAuthUrlResponse\"\x1d\x82\xd3\xe4\x93\x02\x17\x12\x15/v1/auth/get_auth_url\x12w\n\x0cRefreshToken\x12\x1e.prodvana.auth.RefreshTokenReq\x1a\x1f.prodvana.auth.GetAuthTokenResp\"&\x82\xd3\xe4\x93\x02 \"\x1b/v1/auth/refresh_auth_token:\x01*\x12\x44\n\x07\x43liAuth\x12\x19.prodvana.auth.CliAuthReq\x1a\x1a.prodvana.auth.CliAuthResp\"\x00\x30\x01\x32\xc9\x01\n\x12\x41uthSessionManager\x12S\n\x05\x43heck\x12\x14.prodvana.auth.Empty\x1a\x14.prodvana.auth.Empty\"\x1e\x82\xd3\xe4\x93\x02\x18\x12\x16/v1/auth/session/check\x12^\n\x06Logout\x12\x18.prodvana.auth.LogoutReq\x1a\x19.prodvana.auth.LogoutResp\"\x1f\x82\xd3\xe4\x93\x02\x19\"\x17/v1/auth/session/logout2\xa2\x04\n\x0f\x41piTokenManager\x12y\n\x11\x43reateOrgApiToken\x12\x14.prodvana.auth.Empty\x1a$.prodvana.auth.CreateOrgApiTokenResp\"(\x82\xd3\xe4\x93\x02\"\"\x1d/v1/auth/api/create_org_token:\x01*\x12\x8a\x01\n\x12\x43reateOrgApiToken2\x12#.prodvana.auth.CreateOrgApiTokenReq\x1a$.prodvana.auth.CreateOrgApiTokenResp\")\x82\xd3\xe4\x93\x02#\"\x1e/v1/auth/api/org_tokens/create:\x01*\x12\x88\x01\n\x11\x44\x65leteOrgApiToken\x12#.prodvana.auth.DeleteOrgApiTokenReq\x1a$.prodvana.auth.DeleteOrgApiTokenResp\"(\x82\xd3\xe4\x93\x02\"* /v1/auth/api/org_tokens/{name=*}\x12|\n\x10ListOrgApiTokens\x12\".prodvana.auth.ListOrgApiTokensReq\x1a#.prodvana.auth.ListOrgApiTokensResp\"\x1f\x82\xd3\xe4\x93\x02\x19\x12\x17/v1/auth/api/org_tokensBIZGgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/authb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'prodvana.auth.auth_manager_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZGgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/auth'
  _AUTHMANAGER.methods_by_name['GetAuthToken']._options = None
  _AUTHMANAGER.methods_by_name['GetAuthToken']._serialized_options = b'\202\323\344\223\002\024\022\022/v1/auth/get_token'
  _AUTHMANAGER.methods_by_name['GetAuthUrl']._options = None
  _AUTHMANAGER.methods_by_name['GetAuthUrl']._serialized_options = b'\202\323\344\223\002\027\022\025/v1/auth/get_auth_url'
  _AUTHMANAGER.methods_by_name['RefreshToken']._options = None
  _AUTHMANAGER.methods_by_name['RefreshToken']._serialized_options = b'\202\323\344\223\002 \"\033/v1/auth/refresh_auth_token:\001*'
  _AUTHSESSIONMANAGER.methods_by_name['Check']._options = None
  _AUTHSESSIONMANAGER.methods_by_name['Check']._serialized_options = b'\202\323\344\223\002\030\022\026/v1/auth/session/check'
  _AUTHSESSIONMANAGER.methods_by_name['Logout']._options = None
  _AUTHSESSIONMANAGER.methods_by_name['Logout']._serialized_options = b'\202\323\344\223\002\031\"\027/v1/auth/session/logout'
  _APITOKENMANAGER.methods_by_name['CreateOrgApiToken']._options = None
  _APITOKENMANAGER.methods_by_name['CreateOrgApiToken']._serialized_options = b'\202\323\344\223\002\"\"\035/v1/auth/api/create_org_token:\001*'
  _APITOKENMANAGER.methods_by_name['CreateOrgApiToken2']._options = None
  _APITOKENMANAGER.methods_by_name['CreateOrgApiToken2']._serialized_options = b'\202\323\344\223\002#\"\036/v1/auth/api/org_tokens/create:\001*'
  _APITOKENMANAGER.methods_by_name['DeleteOrgApiToken']._options = None
  _APITOKENMANAGER.methods_by_name['DeleteOrgApiToken']._serialized_options = b'\202\323\344\223\002\"* /v1/auth/api/org_tokens/{name=*}'
  _APITOKENMANAGER.methods_by_name['ListOrgApiTokens']._options = None
  _APITOKENMANAGER.methods_by_name['ListOrgApiTokens']._serialized_options = b'\202\323\344\223\002\031\022\027/v1/auth/api/org_tokens'
  _globals['_GETTOKENREQUEST']._serialized_start=168
  _globals['_GETTOKENREQUEST']._serialized_end=271
  _globals['_GETAUTHURLRESPONSE']._serialized_start=273
  _globals['_GETAUTHURLRESPONSE']._serialized_end=311
  _globals['_GETAUTHURLREQUEST']._serialized_start=314
  _globals['_GETAUTHURLREQUEST']._serialized_end=452
  _globals['_CLIAUTHREQ']._serialized_start=454
  _globals['_CLIAUTHREQ']._serialized_end=466
  _globals['_CLIAUTHRESP']._serialized_start=469
  _globals['_CLIAUTHRESP']._serialized_end=695
  _globals['_CLIAUTHRESP_USERPROMPT']._serialized_start=589
  _globals['_CLIAUTHRESP_USERPROMPT']._serialized_end=681
  _globals['_REFRESHTOKENREQ']._serialized_start=697
  _globals['_REFRESHTOKENREQ']._serialized_end=805
  _globals['_GETAUTHTOKENRESP']._serialized_start=807
  _globals['_GETAUTHTOKENRESP']._serialized_end=902
  _globals['_EMPTY']._serialized_start=904
  _globals['_EMPTY']._serialized_end=911
  _globals['_LOGOUTREQ']._serialized_start=913
  _globals['_LOGOUTREQ']._serialized_end=996
  _globals['_LOGOUTRESP']._serialized_start=998
  _globals['_LOGOUTRESP']._serialized_end=1030
  _globals['_CREATEORGAPITOKENREQ']._serialized_start=1032
  _globals['_CREATEORGAPITOKENREQ']._serialized_end=1144
  _globals['_CREATEORGAPITOKENRESP']._serialized_start=1146
  _globals['_CREATEORGAPITOKENRESP']._serialized_end=1231
  _globals['_DELETEORGAPITOKENREQ']._serialized_start=1233
  _globals['_DELETEORGAPITOKENREQ']._serialized_end=1269
  _globals['_DELETEORGAPITOKENRESP']._serialized_start=1271
  _globals['_DELETEORGAPITOKENRESP']._serialized_end=1294
  _globals['_LISTORGAPITOKENSREQ']._serialized_start=1296
  _globals['_LISTORGAPITOKENSREQ']._serialized_end=1317
  _globals['_LISTORGAPITOKENSRESP']._serialized_start=1319
  _globals['_LISTORGAPITOKENSRESP']._serialized_end=1386
  _globals['_AUTHMANAGER']._serialized_start=1389
  _globals['_AUTHMANAGER']._serialized_end=1816
  _globals['_AUTHSESSIONMANAGER']._serialized_start=1819
  _globals['_AUTHSESSIONMANAGER']._serialized_end=2020
  _globals['_APITOKENMANAGER']._serialized_start=2023
  _globals['_APITOKENMANAGER']._serialized_end=2569
# @@protoc_insertion_point(module_scope)
