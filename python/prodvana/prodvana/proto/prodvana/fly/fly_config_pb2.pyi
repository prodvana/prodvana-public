"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import builtins
import collections.abc
import google.protobuf.descriptor
import google.protobuf.internal.containers
import google.protobuf.message
import prodvana.proto.prodvana.common_config.kubernetes_config_pb2
import sys

if sys.version_info >= (3, 8):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

class FlyConfig(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    INLINED_FIELD_NUMBER: builtins.int
    LOCAL_FIELD_NUMBER: builtins.int
    REMOTE_FIELD_NUMBER: builtins.int
    REGIONS_FIELD_NUMBER: builtins.int
    inlined: builtins.str
    @property
    def local(self) -> prodvana.proto.prodvana.common_config.kubernetes_config_pb2.LocalConfig: ...
    @property
    def remote(self) -> prodvana.proto.prodvana.common_config.kubernetes_config_pb2.RemoteConfig: ...
    @property
    def regions(self) -> google.protobuf.internal.containers.RepeatedScalarFieldContainer[builtins.str]:
        """if set, only manage these regions"""
    def __init__(
        self,
        *,
        inlined: builtins.str = ...,
        local: prodvana.proto.prodvana.common_config.kubernetes_config_pb2.LocalConfig | None = ...,
        remote: prodvana.proto.prodvana.common_config.kubernetes_config_pb2.RemoteConfig | None = ...,
        regions: collections.abc.Iterable[builtins.str] | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["inlined", b"inlined", "local", b"local", "remote", b"remote", "toml_oneof", b"toml_oneof"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["inlined", b"inlined", "local", b"local", "regions", b"regions", "remote", b"remote", "toml_oneof", b"toml_oneof"]) -> None: ...
    def WhichOneof(self, oneof_group: typing_extensions.Literal["toml_oneof", b"toml_oneof"]) -> typing_extensions.Literal["inlined", "local", "remote"] | None: ...

global___FlyConfig = FlyConfig
