from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class RecommendationRequest(_message.Message):
    __slots__ = ("user_id", "article_id")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    ARTICLE_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: str
    article_id: int
    def __init__(self, user_id: _Optional[str] = ..., article_id: _Optional[int] = ...) -> None: ...

class RecommendationResponse(_message.Message):
    __slots__ = ("product_ids",)
    PRODUCT_IDS_FIELD_NUMBER: _ClassVar[int]
    product_ids: _containers.RepeatedScalarFieldContainer[int]
    def __init__(self, product_ids: _Optional[_Iterable[int]] = ...) -> None: ...
