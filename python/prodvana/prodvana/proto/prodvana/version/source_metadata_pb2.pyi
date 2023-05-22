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

class _Source:
    ValueType = typing.NewType("ValueType", builtins.int)
    V: typing_extensions.TypeAlias = ValueType

class _SourceEnumTypeWrapper(google.protobuf.internal.enum_type_wrapper._EnumTypeWrapper[_Source.ValueType], builtins.type):  # noqa: F821
    DESCRIPTOR: google.protobuf.descriptor.EnumDescriptor
    UNKNOWN_SOURCE: _Source.ValueType  # 0
    WEB: _Source.ValueType  # 1
    INTERACTIVE_PVNCTL: _Source.ValueType  # 2
    CONFIG_FILE: _Source.ValueType  # 3
    REPO_FOLLOW: _Source.ValueType  # 4

class Source(_Source, metaclass=_SourceEnumTypeWrapper): ...

UNKNOWN_SOURCE: Source.ValueType  # 0
WEB: Source.ValueType  # 1
INTERACTIVE_PVNCTL: Source.ValueType  # 2
CONFIG_FILE: Source.ValueType  # 3
REPO_FOLLOW: Source.ValueType  # 4
global___Source = Source

class SourceMetadata(google.protobuf.message.Message):
    """all of these fields are optional and only set if it makes sense for a given source."""

    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    REPO_URL_FIELD_NUMBER: builtins.int
    FILE_PATH_FIELD_NUMBER: builtins.int
    COMMIT_FIELD_NUMBER: builtins.int
    USER_ID_FIELD_NUMBER: builtins.int
    repo_url: builtins.str
    """full repo URL, like git@github.com:foo/bar.git or https://github.com/foo/bar"""
    file_path: builtins.str
    commit: builtins.str
    user_id: builtins.str
    """set internally, automatically, by Prodvana if the action was initiated by a user"""
    def __init__(
        self,
        *,
        repo_url: builtins.str = ...,
        file_path: builtins.str = ...,
        commit: builtins.str = ...,
        user_id: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["commit", b"commit", "file_path", b"file_path", "repo_url", b"repo_url", "user_id", b"user_id"]) -> None: ...

global___SourceMetadata = SourceMetadata