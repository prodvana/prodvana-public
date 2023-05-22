"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import builtins
import google.protobuf.descriptor
import google.protobuf.message
import prodvana.proto.prodvana.version.source_metadata_pb2
import sys

if sys.version_info >= (3, 8):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

class ObjectMeta(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    ID_FIELD_NUMBER: builtins.int
    NAME_FIELD_NUMBER: builtins.int
    VERSION_FIELD_NUMBER: builtins.int
    CONFIG_VERSION_FIELD_NUMBER: builtins.int
    SOURCE_FIELD_NUMBER: builtins.int
    SOURCE_METADATA_FIELD_NUMBER: builtins.int
    id: builtins.str
    name: builtins.str
    version: builtins.str
    config_version: builtins.str
    """only set for objects who has configurations with parametrization support. At the time of this writing, only services."""
    source: prodvana.proto.prodvana.version.source_metadata_pb2.Source.ValueType
    @property
    def source_metadata(self) -> prodvana.proto.prodvana.version.source_metadata_pb2.SourceMetadata: ...
    def __init__(
        self,
        *,
        id: builtins.str = ...,
        name: builtins.str = ...,
        version: builtins.str = ...,
        config_version: builtins.str = ...,
        source: prodvana.proto.prodvana.version.source_metadata_pb2.Source.ValueType = ...,
        source_metadata: prodvana.proto.prodvana.version.source_metadata_pb2.SourceMetadata | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["source_metadata", b"source_metadata"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["config_version", b"config_version", "id", b"id", "name", b"name", "source", b"source", "source_metadata", b"source_metadata", "version", b"version"]) -> None: ...

global___ObjectMeta = ObjectMeta