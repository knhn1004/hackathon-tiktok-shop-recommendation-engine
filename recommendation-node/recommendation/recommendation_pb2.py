# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: recommendation.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x14recommendation.proto\x12\x0erecommendation\"<\n\x15RecommendationRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x12\n\narticle_id\x18\x02 \x01(\x04\"-\n\x16RecommendationResponse\x12\x13\n\x0bproduct_ids\x18\x01 \x03(\x04\x32|\n\x15RecommendationService\x12\x63\n\x12GetRecommendations\x12%.recommendation.RecommendationRequest\x1a&.recommendation.RecommendationResponseb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'recommendation_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  DESCRIPTOR._loaded_options = None
  _globals['_RECOMMENDATIONREQUEST']._serialized_start=40
  _globals['_RECOMMENDATIONREQUEST']._serialized_end=100
  _globals['_RECOMMENDATIONRESPONSE']._serialized_start=102
  _globals['_RECOMMENDATIONRESPONSE']._serialized_end=147
  _globals['_RECOMMENDATIONSERVICE']._serialized_start=149
  _globals['_RECOMMENDATIONSERVICE']._serialized_end=273
# @@protoc_insertion_point(module_scope)
