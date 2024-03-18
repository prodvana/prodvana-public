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

    INLINED_FIELD_NUMBER: builtins.int
    LOCAL_FIELD_NUMBER: builtins.int
    REMOTE_FIELD_NUMBER: builtins.int
    inlined: builtins.str
    @property
    def local(self) -> prodvana.proto.prodvana.common_config.kubernetes_config_pb2.LocalConfig: ...
    @property
    def remote(self) -> prodvana.proto.prodvana.common_config.kubernetes_config_pb2.RemoteConfig: ...
    def __init__(
        self,
        *,
        inlined: builtins.str = ...,
        local: prodvana.proto.prodvana.common_config.kubernetes_config_pb2.LocalConfig | None = ...,
        remote: prodvana.proto.prodvana.common_config.kubernetes_config_pb2.RemoteConfig | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["inlined", b"inlined", "local", b"local", "override_oneof", b"override_oneof", "remote", b"remote"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["inlined", b"inlined", "local", b"local", "override_oneof", b"override_oneof", "remote", b"remote"]) -> None: ...
    def WhichOneof(self, oneof_group: typing_extensions.Literal["override_oneof", b"override_oneof"]) -> typing_extensions.Literal["inlined", "local", "remote"] | None: ...

global___HelmValuesOverrides = HelmValuesOverrides

class HelmConfig(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    REMOTE_FIELD_NUMBER: builtins.int
    LOCAL_FIELD_NUMBER: builtins.int
    HELM_TARBALL_BLOB_ID_FIELD_NUMBER: builtins.int
    VALUES_OVERRIDES_FIELD_NUMBER: builtins.int
    RELEASE_NAME_FIELD_NUMBER: builtins.int
    NAMESPACE_FIELD_NUMBER: builtins.int
    FIXUP_OWNERSHIP_FIELD_NUMBER: builtins.int
    @property
    def remote(self) -> global___RemoteHelmChart: ...
    @property
    def local(self) -> prodvana.proto.prodvana.common_config.kubernetes_config_pb2.LocalConfig: ...
    helm_tarball_blob_id: builtins.str
    @property
    def values_overrides(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[global___HelmValuesOverrides]: ...
    release_name: builtins.str
    """optional release name, leave blank to have Prodvana generate one.
    Mainly useful for migrating an existing helm release into Prodvana.
    """
    namespace: builtins.str
    """used internally by Prodvana, do not set."""
    fixup_ownership: builtins.bool
    """Before running helm commands, first fixup labels and annotations of the Kubernetes objects
    to match the expected state. This is useful for migrating an existing Kubernetes object to be managed
    by Helm.
    """
    def __init__(
        self,
        *,
        remote: global___RemoteHelmChart | None = ...,
        local: prodvana.proto.prodvana.common_config.kubernetes_config_pb2.LocalConfig | None = ...,
        helm_tarball_blob_id: builtins.str = ...,
        values_overrides: collections.abc.Iterable[global___HelmValuesOverrides] | None = ...,
        release_name: builtins.str = ...,
        namespace: builtins.str = ...,
        fixup_ownership: builtins.bool = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["chart_oneof", b"chart_oneof", "helm_tarball_blob_id", b"helm_tarball_blob_id", "local", b"local", "remote", b"remote"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["chart_oneof", b"chart_oneof", "fixup_ownership", b"fixup_ownership", "helm_tarball_blob_id", b"helm_tarball_blob_id", "local", b"local", "namespace", b"namespace", "release_name", b"release_name", "remote", b"remote", "values_overrides", b"values_overrides"]) -> None: ...
    def WhichOneof(self, oneof_group: typing_extensions.Literal["chart_oneof", b"chart_oneof"]) -> typing_extensions.Literal["remote", "local", "helm_tarball_blob_id"] | None: ...

global___HelmConfig = HelmConfig
