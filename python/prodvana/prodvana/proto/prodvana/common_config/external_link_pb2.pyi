"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import builtins
import google.protobuf.descriptor
import google.protobuf.internal.enum_type_wrapper
import google.protobuf.message
import sys
import typing

if sys.version_info >= (3, 10):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

class ExternalLink(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    class _LinkType:
        ValueType = typing.NewType("ValueType", builtins.int)
        V: typing_extensions.TypeAlias = ValueType

    class _LinkTypeEnumTypeWrapper(google.protobuf.internal.enum_type_wrapper._EnumTypeWrapper[ExternalLink._LinkType.ValueType], builtins.type):  # noqa: F821
        DESCRIPTOR: google.protobuf.descriptor.EnumDescriptor
        UNKNOWN: ExternalLink._LinkType.ValueType  # 0
        DETAIL: ExternalLink._LinkType.ValueType  # 1
        LOG: ExternalLink._LinkType.ValueType  # 2

    class LinkType(_LinkType, metaclass=_LinkTypeEnumTypeWrapper): ...
    UNKNOWN: ExternalLink.LinkType.ValueType  # 0
    DETAIL: ExternalLink.LinkType.ValueType  # 1
    LOG: ExternalLink.LinkType.ValueType  # 2

    TYPE_FIELD_NUMBER: builtins.int
    URL_FIELD_NUMBER: builtins.int
    NAME_FIELD_NUMBER: builtins.int
    type: global___ExternalLink.LinkType.ValueType
    url: builtins.str
    name: builtins.str
    def __init__(
        self,
        *,
        type: global___ExternalLink.LinkType.ValueType = ...,
        url: builtins.str = ...,
        name: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["name", b"name", "type", b"type", "url", b"url"]) -> None: ...

global___ExternalLink = ExternalLink
