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

class Output(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    EXIT_CODE_FIELD_NUMBER: builtins.int
    EXEC_ERROR_FIELD_NUMBER: builtins.int
    STDOUT_BLOB_ID_FIELD_NUMBER: builtins.int
    STDERR_BLOB_ID_FIELD_NUMBER: builtins.int
    VERSION_FIELD_NUMBER: builtins.int
    START_TIMESTAMP_NS_FIELD_NUMBER: builtins.int
    DURATION_NS_FIELD_NUMBER: builtins.int
    FILES_FIELD_NUMBER: builtins.int
    HOSTNAME_FIELD_NUMBER: builtins.int
    exit_code: builtins.int
    """Exit code of wrapped process. -1 if process failed to execute."""
    exec_error: builtins.str
    """Internal error when trying to execute wrapped process."""
    stdout_blob_id: builtins.str
    stderr_blob_id: builtins.str
    version: builtins.str
    """Wrapper version."""
    start_timestamp_ns: builtins.int
    """Timestamp when the process began executing, in ns."""
    duration_ns: builtins.int
    """Total execution duration of the process, in ns."""
    @property
    def files(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[global___OutputFile]: ...
    hostname: builtins.str
    def __init__(
        self,
        *,
        exit_code: builtins.int = ...,
        exec_error: builtins.str = ...,
        stdout_blob_id: builtins.str = ...,
        stderr_blob_id: builtins.str = ...,
        version: builtins.str = ...,
        start_timestamp_ns: builtins.int = ...,
        duration_ns: builtins.int = ...,
        files: collections.abc.Iterable[global___OutputFile] | None = ...,
        hostname: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["duration_ns", b"duration_ns", "exec_error", b"exec_error", "exit_code", b"exit_code", "files", b"files", "hostname", b"hostname", "start_timestamp_ns", b"start_timestamp_ns", "stderr_blob_id", b"stderr_blob_id", "stdout_blob_id", b"stdout_blob_id", "version", b"version"]) -> None: ...

global___Output = Output

class OutputFile(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    NAME_FIELD_NUMBER: builtins.int
    CONTENT_BLOB_ID_FIELD_NUMBER: builtins.int
    name: builtins.str
    content_blob_id: builtins.str
    def __init__(
        self,
        *,
        name: builtins.str = ...,
        content_blob_id: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["content_blob_id", b"content_blob_id", "name", b"name"]) -> None: ...

global___OutputFile = OutputFile
