# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: pb/svc/audio/audio.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x18pb/svc/audio/audio.proto\">\n\x0bRequirement\x12\x0f\n\x07\x63ontent\x18\x01 \x01(\t\x12\x0f\n\x07speaker\x18\x02 \x01(\t\x12\r\n\x05token\x18\x03 \x01(\t\"$\n\x05\x41udio\x12\x0c\n\x04\x64\x61ta\x18\x01 \x01(\x0c\x12\r\n\x05\x65rror\x18\x02 \x01(\t\"N\n\x07Request\x12\x13\n\x04\x61uth\x18\x01 \x01(\x0b\x32\x05.Auth\x12\x0f\n\x07\x63ontent\x18\x02 \x01(\t\x12\x0f\n\x07speaker\x18\x03 \x01(\t\x12\x0c\n\x04path\x18\x04 \x01(\t\"#\n\x04\x41uth\x12\r\n\x05token\x18\x01 \x01(\t\x12\x0c\n\x04\x66rom\x18\x02 \x01(\t\"(\n\x03Job\x12\x0f\n\x07\x63ontent\x18\x01 \x01(\t\x12\x10\n\x08speacker\x18\x02 \x01(\t\"\x1f\n\x08\x43hecking\x12\x13\n\x04\x61uth\x18\x01 \x01(\x0b\x32\x05.Auth2\xa0\x01\n\x16\x41udioGenerationService\x12%\n\rGenerateAudio\x12\x0c.Requirement\x1a\x06.Audio\x12 \n\x0cMakingNewJob\x12\x08.Request\x1a\x06.Audio\x12\x1e\n\x0b\x43heckingJob\x12\t.Checking\x1a\x04.Job\x12\x1d\n\rSendingResult\x12\x06.Audio\x1a\x04.JobB7Z5github.com/aglide100/speech-test/cluster/pb/svc/audiob\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'pb.svc.audio.audio_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z5github.com/aglide100/speech-test/cluster/pb/svc/audio'
  _globals['_REQUIREMENT']._serialized_start=28
  _globals['_REQUIREMENT']._serialized_end=90
  _globals['_AUDIO']._serialized_start=92
  _globals['_AUDIO']._serialized_end=128
  _globals['_REQUEST']._serialized_start=130
  _globals['_REQUEST']._serialized_end=208
  _globals['_AUTH']._serialized_start=210
  _globals['_AUTH']._serialized_end=245
  _globals['_JOB']._serialized_start=247
  _globals['_JOB']._serialized_end=287
  _globals['_CHECKING']._serialized_start=289
  _globals['_CHECKING']._serialized_end=320
  _globals['_AUDIOGENERATIONSERVICE']._serialized_start=323
  _globals['_AUDIOGENERATIONSERVICE']._serialized_end=483
# @@protoc_insertion_point(module_scope)
