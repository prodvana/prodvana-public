"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import builtins
import collections.abc
import google.protobuf.descriptor
import google.protobuf.internal.containers
import google.protobuf.message
import sys

if sys.version_info >= (3, 8):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

class GetInfoOutput(google.protobuf.message.Message):
    """output format for the get_info command, in json format"""

    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    OUTPUTS_FIELD_NUMBER: builtins.int
    @property
    def outputs(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[global___OutputContent]: ...
    def __init__(
        self,
        *,
        outputs: collections.abc.Iterable[global___OutputContent] | None = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["outputs", b"outputs"]) -> None: ...

global___GetInfoOutput = GetInfoOutput

class OutputContent(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    NAME_FIELD_NUMBER: builtins.int
    TEXT_FIELD_NUMBER: builtins.int
    name: builtins.str
    text: builtins.str
    def __init__(
        self,
        *,
        name: builtins.str = ...,
        text: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["name", b"name", "text", b"text"]) -> None: ...

global___OutputContent = OutputContent