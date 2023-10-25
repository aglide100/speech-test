from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Audio(_message.Message):
    __slots__ = ["data"]
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: bytes
    def __init__(self, data: _Optional[bytes] = ...) -> None: ...

class Error(_message.Message):
    __slots__ = ["msg"]
    MSG_FIELD_NUMBER: _ClassVar[int]
    msg: str
    def __init__(self, msg: _Optional[str] = ...) -> None: ...

class Auth(_message.Message):
    __slots__ = ["token", "who"]
    TOKEN_FIELD_NUMBER: _ClassVar[int]
    WHO_FIELD_NUMBER: _ClassVar[int]
    token: str
    who: str
    def __init__(self, token: _Optional[str] = ..., who: _Optional[str] = ...) -> None: ...

class Job(_message.Message):
    __slots__ = ["content", "speaker", "id", "no"]
    CONTENT_FIELD_NUMBER: _ClassVar[int]
    SPEAKER_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    NO_FIELD_NUMBER: _ClassVar[int]
    content: str
    speaker: str
    id: str
    no: int
    def __init__(self, content: _Optional[str] = ..., speaker: _Optional[str] = ..., id: _Optional[str] = ..., no: _Optional[int] = ...) -> None: ...

class CheckingJobReq(_message.Message):
    __slots__ = ["auth"]
    AUTH_FIELD_NUMBER: _ClassVar[int]
    auth: Auth
    def __init__(self, auth: _Optional[_Union[Auth, _Mapping]] = ...) -> None: ...

class CheckingJobRes(_message.Message):
    __slots__ = ["job", "error"]
    JOB_FIELD_NUMBER: _ClassVar[int]
    ERROR_FIELD_NUMBER: _ClassVar[int]
    job: Job
    error: Error
    def __init__(self, job: _Optional[_Union[Job, _Mapping]] = ..., error: _Optional[_Union[Error, _Mapping]] = ...) -> None: ...

class SendingResultReq(_message.Message):
    __slots__ = ["audio", "auth", "job"]
    AUDIO_FIELD_NUMBER: _ClassVar[int]
    AUTH_FIELD_NUMBER: _ClassVar[int]
    JOB_FIELD_NUMBER: _ClassVar[int]
    audio: Audio
    auth: Auth
    job: Job
    def __init__(self, audio: _Optional[_Union[Audio, _Mapping]] = ..., auth: _Optional[_Union[Auth, _Mapping]] = ..., job: _Optional[_Union[Job, _Mapping]] = ...) -> None: ...

class MakingNewJobReq(_message.Message):
    __slots__ = ["auth", "content", "speaker", "path"]
    AUTH_FIELD_NUMBER: _ClassVar[int]
    CONTENT_FIELD_NUMBER: _ClassVar[int]
    SPEAKER_FIELD_NUMBER: _ClassVar[int]
    PATH_FIELD_NUMBER: _ClassVar[int]
    auth: Auth
    content: str
    speaker: str
    path: str
    def __init__(self, auth: _Optional[_Union[Auth, _Mapping]] = ..., content: _Optional[str] = ..., speaker: _Optional[str] = ..., path: _Optional[str] = ...) -> None: ...

class GetAudioReq(_message.Message):
    __slots__ = ["auth", "job"]
    AUTH_FIELD_NUMBER: _ClassVar[int]
    JOB_FIELD_NUMBER: _ClassVar[int]
    auth: Auth
    job: Job
    def __init__(self, auth: _Optional[_Union[Auth, _Mapping]] = ..., job: _Optional[_Union[Job, _Mapping]] = ...) -> None: ...
