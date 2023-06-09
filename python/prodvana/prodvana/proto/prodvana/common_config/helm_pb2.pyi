"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import builtins
import collections.abc
import google.protobuf.descriptor
import google.protobuf.internal.containers
import google.protobuf.message
import prodvana.proto.prodvana.common_config.env_pb2
import prodvana.proto.prodvana.common_config.kubernetes_config_pb2
import sys

if sys.version_info >= (3, 8):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

class RemoteHelmChart(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    REPO_FIELD_NUMBER: builtins.int
    CHART_FIELD_NUMBER: builtins.int
    CHART_VERSION_FIELD_NUMBER: builtins.int
    repo: builtins.str
    """TODO(naphat) add integration support for private repositories"""
    chart: builtins.str
    chart_version: builtins.str
    def __init__(
        self,
        *,
        repo: builtins.str = ...,
        chart: builtins.str = ...,
        chart_version: builtins.str = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["repo", b"repo", "repo_oneof", b"repo_oneof"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["chart", b"chart", "chart_version", b"chart_version", "repo", b"repo", "repo_oneof", b"repo_oneof"]) -> None: ...
    def WhichOneof(self, oneof_group: typing_extensions.Literal["repo_oneof", b"repo_oneof"]) -> typing_extensions.Literal["repo"] | None: ...

global___RemoteHelmChart = RemoteHelmChart

class HelmValuesOverrides(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    class MapEntry(google.protobuf.message.Message):
        DESCRIPTOR: google.protobuf.descriptor.Descriptor

        KEY_FIELD_NUMBER: builtins.int
        VALUE_FIELD_NUMBER: builtins.int
        key: builtins.str
        @property
        def value(self) -> prodvana.proto.prodvana.common_config.env_pb2.EnvValue: ...
        def __init__(
            self,
            *,
            key: builtins.str = ...,
            value: prodvana.proto.prodvana.common_config.env_pb2.EnvValue | None = ...,
        ) -> None: ...
        def HasField(self, field_name: typing_extensions.Literal["value", b"value"]) -> builtins.bool: ...
        def ClearField(self, field_name: typing_extensions.Literal["key", b"key", "value", b"value"]) -> None: ...

    INLINED_FIELD_NUMBER: builtins.int
    LOCAL_FIELD_NUMBER: builtins.int
    MAP_FIELD_NUMBER: builtins.int
    inlined: builtins.str
    @property
    def local(self) -> prodvana.proto.prodvana.common_config.kubernetes_config_pb2.LocalConfig: ...
    @property
    def map(self) -> google.protobuf.internal.containers.MessageMap[builtins.str, prodvana.proto.prodvana.common_config.env_pb2.EnvValue]:
        """treat this as part of the above oneof, even though proto does not allow us to"""
    def __init__(
        self,
        *,
        inlined: builtins.str = ...,
        local: prodvana.proto.prodvana.common_config.kubernetes_config_pb2.LocalConfig | None = ...,
        map: collections.abc.Mapping[builtins.str, prodvana.proto.prodvana.common_config.env_pb2.EnvValue] | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["inlined", b"inlined", "local", b"local", "override_oneof", b"override_oneof"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["inlined", b"inlined", "local", b"local", "map", b"map", "override_oneof", b"override_oneof"]) -> None: ...
    def WhichOneof(self, oneof_group: typing_extensions.Literal["override_oneof", b"override_oneof"]) -> typing_extensions.Literal["inlined", "local"] | None: ...

global___HelmValuesOverrides = HelmValuesOverrides

class HelmConfig(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    REMOTE_FIELD_NUMBER: builtins.int
    VALUES_OVERRIDES_FIELD_NUMBER: builtins.int
    RELEASE_NAME_FIELD_NUMBER: builtins.int
    NAMESPACE_FIELD_NUMBER: builtins.int
    @property
    def remote(self) -> global___RemoteHelmChart:
        """TODO(naphat) chart archive support, local charts"""
    @property
    def values_overrides(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[global___HelmValuesOverrides]: ...
    release_name: builtins.str
    """optional release name, leave blank to have Prodvana generate one.
    Mainly useful for migrating an existing helm release into Prodvana.
    """
    namespace: builtins.str
    """used internally by Prodvana, do not set."""
    def __init__(
        self,
        *,
        remote: global___RemoteHelmChart | None = ...,
        values_overrides: collections.abc.Iterable[global___HelmValuesOverrides] | None = ...,
        release_name: builtins.str = ...,
        namespace: builtins.str = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["chart_oneof", b"chart_oneof", "remote", b"remote"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["chart_oneof", b"chart_oneof", "namespace", b"namespace", "release_name", b"release_name", "remote", b"remote", "values_overrides", b"values_overrides"]) -> None: ...
    def WhichOneof(self, oneof_group: typing_extensions.Literal["chart_oneof", b"chart_oneof"]) -> typing_extensions.Literal["remote"] | None: ...

global___HelmConfig = HelmConfig
