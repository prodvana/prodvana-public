# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: prodvana/recipe/recipe_config.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from prodvana.proto.validate import validate_pb2 as validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n#prodvana/recipe/recipe_config.proto\x12\x0fprodvana.recipe\x1a\x17validate/validate.proto\"a\n\rServiceTarget\x12\x13\n\x0b\x61pplication\x18\x01 \x01(\t\x12\x16\n\x0e\x61pplication_id\x18\x02 \x01(\t\x12\x0f\n\x07service\x18\x03 \x01(\t\x12\x12\n\nservice_id\x18\x04 \x01(\t\"E\n\x06Target\x12\x31\n\x07service\x18\x01 \x01(\x0b\x32\x1e.prodvana.recipe.ServiceTargetH\x00\x42\x08\n\x06target\"7\n\x0bRecipeStage\x12(\n\x07targets\x18\x01 \x03(\x0b\x32\x17.prodvana.recipe.Target\"\x87\x02\n\x0cRecipeConfig\x12\x39\n\x04name\x18\x01 \x01(\tB+\xfa\x42(r&\x10\x01\x18?2 ^[a-z]([a-z0-9-]*[a-z0-9]){0,1}$\x12\x13\n\x0b\x61pplication\x18\x02 \x01(\t\x12,\n\x06stages\x18\x03 \x03(\x0b\x32\x1c.prodvana.recipe.RecipeStage\x12\x38\n\x08strategy\x18\x04 \x01(\x0e\x32&.prodvana.recipe.RecipeConfig.Strategy\"?\n\x08Strategy\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x0e\n\nBY_SERVICE\x10\x01\x12\x16\n\x12\x42Y_RELEASE_CHANNEL\x10\x02\x42KZIgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/recipeb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'prodvana.recipe.recipe_config_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZIgithub.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/recipe'
  _RECIPECONFIG.fields_by_name['name']._options = None
  _RECIPECONFIG.fields_by_name['name']._serialized_options = b'\372B(r&\020\001\030?2 ^[a-z]([a-z0-9-]*[a-z0-9]){0,1}$'
  _globals['_SERVICETARGET']._serialized_start=81
  _globals['_SERVICETARGET']._serialized_end=178
  _globals['_TARGET']._serialized_start=180
  _globals['_TARGET']._serialized_end=249
  _globals['_RECIPESTAGE']._serialized_start=251
  _globals['_RECIPESTAGE']._serialized_end=306
  _globals['_RECIPECONFIG']._serialized_start=309
  _globals['_RECIPECONFIG']._serialized_end=572
  _globals['_RECIPECONFIG_STRATEGY']._serialized_start=509
  _globals['_RECIPECONFIG_STRATEGY']._serialized_end=572
# @@protoc_insertion_point(module_scope)
