"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import builtins
import google.protobuf.descriptor
import google.protobuf.message
import sys

if sys.version_info >= (3, 8):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

class LabelDefinition(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    LABEL_FIELD_NUMBER: builtins.int
    VALUE_FIELD_NUMBER: builtins.int
    label: builtins.str
    value: builtins.str
    def __init__(
        self,
        *,
        label: builtins.str = ...,
        value: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["label", b"label", "value", b"value"]) -> None: ...

global___LabelDefinition = LabelDefinition
