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

class _TaskStatus:
    ValueType = typing.NewType("ValueType", builtins.int)
    V: typing_extensions.TypeAlias = ValueType

class _TaskStatusEnumTypeWrapper(google.protobuf.internal.enum_type_wrapper._EnumTypeWrapper[_TaskStatus.ValueType], builtins.type):  # noqa: F821
    DESCRIPTOR: google.protobuf.descriptor.EnumDescriptor
    UNKNOWN: _TaskStatus.ValueType  # 0
    PENDING: _TaskStatus.ValueType  # 1
    RUNNING: _TaskStatus.ValueType  # 2
    SUCCESS: _TaskStatus.ValueType  # 3
    FAILED: _TaskStatus.ValueType  # 4

class TaskStatus(_TaskStatus, metaclass=_TaskStatusEnumTypeWrapper): ...

UNKNOWN: TaskStatus.ValueType  # 0
PENDING: TaskStatus.ValueType  # 1
RUNNING: TaskStatus.ValueType  # 2
SUCCESS: TaskStatus.ValueType  # 3
FAILED: TaskStatus.ValueType  # 4
global___TaskStatus = TaskStatus

class TaskResult(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    LOG_FIELD_NUMBER: builtins.int
    log: builtins.bytes
    """Will be set if task failed, otherwise may be empty."""
    def __init__(
        self,
        *,
        log: builtins.bytes = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["log", b"log"]) -> None: ...

global___TaskResult = TaskResult