# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: audio.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0b\x61udio.proto\">\n\x0bRequirement\x12\x0f\n\x07\x63ontent\x18\x01 \x01(\t\x12\x0f\n\x07speaker\x18\x02 \x01(\t\x12\r\n\x05token\x18\x03 \x01(\t\"$\n\x05\x41udio\x12\x0c\n\x04\x64\x61ta\x18\x01 \x01(\x0c\x12\r\n\x05\x65rror\x18\x02 \x01(\t2?\n\x16\x41udioGenerationService\x12%\n\rGenerateAudio\x12\x0c.Requirement\x1a\x06.Audiob\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'audio_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  _globals['_REQUIREMENT']._serialized_start=15
  _globals['_REQUIREMENT']._serialized_end=77
  _globals['_AUDIO']._serialized_start=79
  _globals['_AUDIO']._serialized_end=115
  _globals['_AUDIOGENERATIONSERVICE']._serialized_start=117
  _globals['_AUDIOGENERATIONSERVICE']._serialized_end=180
# @@protoc_insertion_point(module_scope)