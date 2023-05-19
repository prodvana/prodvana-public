"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import builtins
import collections.abc
import google.protobuf.descriptor
import google.protobuf.internal.containers
import google.protobuf.message
import google.protobuf.timestamp_pb2
import prodvana.proto.prodvana.config_writeback.writeback_pb2
import prodvana.proto.prodvana.environment.clusters_pb2
import prodvana.proto.prodvana.version.source_metadata_pb2
import sys

if sys.version_info >= (3, 8):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

class LinkClusterReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    NAME_FIELD_NUMBER: builtins.int
    AUTH_FIELD_NUMBER: builtins.int
    PRODVANA_MANAGED_FIELD_NUMBER: builtins.int
    DISABLE_ISTIO_FIELD_NUMBER: builtins.int
    DISABLE_FLAGGER_FIELD_NUMBER: builtins.int
    TYPE_FIELD_NUMBER: builtins.int
    SOURCE_FIELD_NUMBER: builtins.int
    SOURCE_METADATA_FIELD_NUMBER: builtins.int
    name: builtins.str
    @property
    def auth(self) -> prodvana.proto.prodvana.environment.clusters_pb2.ClusterAuth: ...
    prodvana_managed: builtins.bool
    """HACK(naphat) this should never be set by users. delete once pulumi runner no longer relies on this endpoint."""
    disable_istio: builtins.bool
    disable_flagger: builtins.bool
    type: prodvana.proto.prodvana.environment.clusters_pb2.ClusterType.ValueType
    source: prodvana.proto.prodvana.version.source_metadata_pb2.Source.ValueType
    @property
    def source_metadata(self) -> prodvana.proto.prodvana.version.source_metadata_pb2.SourceMetadata: ...
    def __init__(
        self,
        *,
        name: builtins.str = ...,
        auth: prodvana.proto.prodvana.environment.clusters_pb2.ClusterAuth | None = ...,
        prodvana_managed: builtins.bool = ...,
        disable_istio: builtins.bool = ...,
        disable_flagger: builtins.bool = ...,
        type: prodvana.proto.prodvana.environment.clusters_pb2.ClusterType.ValueType = ...,
        source: prodvana.proto.prodvana.version.source_metadata_pb2.Source.ValueType = ...,
        source_metadata: prodvana.proto.prodvana.version.source_metadata_pb2.SourceMetadata | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["auth", b"auth", "source_metadata", b"source_metadata"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["auth", b"auth", "disable_flagger", b"disable_flagger", "disable_istio", b"disable_istio", "name", b"name", "prodvana_managed", b"prodvana_managed", "source", b"source", "source_metadata", b"source_metadata", "type", b"type"]) -> None: ...

global___LinkClusterReq = LinkClusterReq

class LinkClusterResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    SUCCESS_FIELD_NUMBER: builtins.int
    MSG_FIELD_NUMBER: builtins.int
    CLUSTER_ID_FIELD_NUMBER: builtins.int
    K8S_AGENT_URL_FIELD_NUMBER: builtins.int
    success: builtins.bool
    """LinkCluster will attempt to talk to the cluster to validate the
    credentials. If it fails, want to communicate back what the failure was.
    """
    msg: builtins.str
    cluster_id: builtins.str
    k8s_agent_url: builtins.str
    """Magic URL to install agent"""
    def __init__(
        self,
        *,
        success: builtins.bool = ...,
        msg: builtins.str = ...,
        cluster_id: builtins.str = ...,
        k8s_agent_url: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["cluster_id", b"cluster_id", "k8s_agent_url", b"k8s_agent_url", "msg", b"msg", "success", b"success"]) -> None: ...

global___LinkClusterResp = LinkClusterResp

class RemoveClusterReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    NAME_FIELD_NUMBER: builtins.int
    name: builtins.str
    """TODO(mike): add an option to remove by cluster_id as well?"""
    def __init__(
        self,
        *,
        name: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["name", b"name"]) -> None: ...

global___RemoveClusterReq = RemoveClusterReq

class RemoveClusterResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    def __init__(
        self,
    ) -> None: ...

global___RemoveClusterResp = RemoveClusterResp

class GetClusterAuthReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    CLUSTER_ID_FIELD_NUMBER: builtins.int
    cluster_id: builtins.str
    def __init__(
        self,
        *,
        cluster_id: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["cluster_id", b"cluster_id"]) -> None: ...

global___GetClusterAuthReq = GetClusterAuthReq

class GetClusterAuthResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    AUTH_FIELD_NUMBER: builtins.int
    @property
    def auth(self) -> prodvana.proto.prodvana.environment.clusters_pb2.ClusterAuth: ...
    def __init__(
        self,
        *,
        auth: prodvana.proto.prodvana.environment.clusters_pb2.ClusterAuth | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["auth", b"auth"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["auth", b"auth"]) -> None: ...

global___GetClusterAuthResp = GetClusterAuthResp

class ListClustersReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    def __init__(
        self,
    ) -> None: ...

global___ListClustersReq = ListClustersReq

class ListClustersResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    class ClusterInfo(google.protobuf.message.Message):
        DESCRIPTOR: google.protobuf.descriptor.Descriptor

        class EcsInfo(google.protobuf.message.Message):
            DESCRIPTOR: google.protobuf.descriptor.Descriptor

            REGION_FIELD_NUMBER: builtins.int
            CLUSTER_ARN_FIELD_NUMBER: builtins.int
            region: builtins.str
            cluster_arn: builtins.str
            def __init__(
                self,
                *,
                region: builtins.str = ...,
                cluster_arn: builtins.str = ...,
            ) -> None: ...
            def ClearField(self, field_name: typing_extensions.Literal["cluster_arn", b"cluster_arn", "region", b"region"]) -> None: ...

        NAME_FIELD_NUMBER: builtins.int
        ID_FIELD_NUMBER: builtins.int
        ORIGIN_FIELD_NUMBER: builtins.int
        ENDPOINT_FIELD_NUMBER: builtins.int
        SERVICE_ACCOUNT_FIELD_NUMBER: builtins.int
        WRITEBACK_CONFIG_FIELD_NUMBER: builtins.int
        TYPE_FIELD_NUMBER: builtins.int
        ECS_FIELD_NUMBER: builtins.int
        LAST_HEARTBEAT_TIMESTAMP_FIELD_NUMBER: builtins.int
        name: builtins.str
        id: builtins.str
        origin: prodvana.proto.prodvana.environment.clusters_pb2.Cluster.Origin.ValueType
        endpoint: builtins.str
        service_account: builtins.str
        @property
        def writeback_config(self) -> prodvana.proto.prodvana.config_writeback.writeback_pb2.ConfigWritebackPath: ...
        type: prodvana.proto.prodvana.environment.clusters_pb2.ClusterType.ValueType
        @property
        def ecs(self) -> global___ListClustersResp.ClusterInfo.EcsInfo: ...
        @property
        def last_heartbeat_timestamp(self) -> google.protobuf.timestamp_pb2.Timestamp: ...
        def __init__(
            self,
            *,
            name: builtins.str = ...,
            id: builtins.str = ...,
            origin: prodvana.proto.prodvana.environment.clusters_pb2.Cluster.Origin.ValueType = ...,
            endpoint: builtins.str = ...,
            service_account: builtins.str = ...,
            writeback_config: prodvana.proto.prodvana.config_writeback.writeback_pb2.ConfigWritebackPath | None = ...,
            type: prodvana.proto.prodvana.environment.clusters_pb2.ClusterType.ValueType = ...,
            ecs: global___ListClustersResp.ClusterInfo.EcsInfo | None = ...,
            last_heartbeat_timestamp: google.protobuf.timestamp_pb2.Timestamp | None = ...,
        ) -> None: ...
        def HasField(self, field_name: typing_extensions.Literal["ecs", b"ecs", "info", b"info", "last_heartbeat_timestamp", b"last_heartbeat_timestamp", "writeback_config", b"writeback_config"]) -> builtins.bool: ...
        def ClearField(self, field_name: typing_extensions.Literal["ecs", b"ecs", "endpoint", b"endpoint", "id", b"id", "info", b"info", "last_heartbeat_timestamp", b"last_heartbeat_timestamp", "name", b"name", "origin", b"origin", "service_account", b"service_account", "type", b"type", "writeback_config", b"writeback_config"]) -> None: ...
        def WhichOneof(self, oneof_group: typing_extensions.Literal["info", b"info"]) -> typing_extensions.Literal["ecs"] | None: ...

    CLUSTERS_FIELD_NUMBER: builtins.int
    @property
    def clusters(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[global___ListClustersResp.ClusterInfo]: ...
    def __init__(
        self,
        *,
        clusters: collections.abc.Iterable[global___ListClustersResp.ClusterInfo] | None = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["clusters", b"clusters"]) -> None: ...

global___ListClustersResp = ListClustersResp

class ConfigureClusterReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    RUNTIME_NAME_FIELD_NUMBER: builtins.int
    CONFIG_FIELD_NUMBER: builtins.int
    SOURCE_FIELD_NUMBER: builtins.int
    SOURCE_METADATA_FIELD_NUMBER: builtins.int
    runtime_name: builtins.str
    @property
    def config(self) -> prodvana.proto.prodvana.environment.clusters_pb2.ClusterConfig: ...
    source: prodvana.proto.prodvana.version.source_metadata_pb2.Source.ValueType
    @property
    def source_metadata(self) -> prodvana.proto.prodvana.version.source_metadata_pb2.SourceMetadata: ...
    def __init__(
        self,
        *,
        runtime_name: builtins.str = ...,
        config: prodvana.proto.prodvana.environment.clusters_pb2.ClusterConfig | None = ...,
        source: prodvana.proto.prodvana.version.source_metadata_pb2.Source.ValueType = ...,
        source_metadata: prodvana.proto.prodvana.version.source_metadata_pb2.SourceMetadata | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["config", b"config", "source_metadata", b"source_metadata"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["config", b"config", "runtime_name", b"runtime_name", "source", b"source", "source_metadata", b"source_metadata"]) -> None: ...

global___ConfigureClusterReq = ConfigureClusterReq

class ConfigureClusterResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    def __init__(
        self,
    ) -> None: ...

global___ConfigureClusterResp = ConfigureClusterResp

class ValidateConfigureClusterResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    def __init__(
        self,
    ) -> None: ...

global___ValidateConfigureClusterResp = ValidateConfigureClusterResp

class GetClusterConfigReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    RUNTIME_NAME_FIELD_NUMBER: builtins.int
    runtime_name: builtins.str
    def __init__(
        self,
        *,
        runtime_name: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["runtime_name", b"runtime_name"]) -> None: ...

global___GetClusterConfigReq = GetClusterConfigReq

class GetClusterConfigResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    CONFIG_FIELD_NUMBER: builtins.int
    @property
    def config(self) -> prodvana.proto.prodvana.environment.clusters_pb2.ClusterConfig: ...
    def __init__(
        self,
        *,
        config: prodvana.proto.prodvana.environment.clusters_pb2.ClusterConfig | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["config", b"config"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["config", b"config"]) -> None: ...

global___GetClusterConfigResp = GetClusterConfigResp

class DetectClusterConfigReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    RUNTIME_NAME_FIELD_NUMBER: builtins.int
    runtime_name: builtins.str
    def __init__(
        self,
        *,
        runtime_name: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["runtime_name", b"runtime_name"]) -> None: ...

global___DetectClusterConfigReq = DetectClusterConfigReq

class DetectClusterConfigResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    CONFIG_FIELD_NUMBER: builtins.int
    @property
    def config(self) -> prodvana.proto.prodvana.environment.clusters_pb2.ClusterConfig: ...
    def __init__(
        self,
        *,
        config: prodvana.proto.prodvana.environment.clusters_pb2.ClusterConfig | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["config", b"config"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["config", b"config"]) -> None: ...

global___DetectClusterConfigResp = DetectClusterConfigResp

class GetClusterStatusReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    CLUSTER_ID_FIELD_NUMBER: builtins.int
    cluster_id: builtins.str
    def __init__(
        self,
        *,
        cluster_id: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["cluster_id", b"cluster_id"]) -> None: ...

global___GetClusterStatusReq = GetClusterStatusReq

class GetClusterStatusResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    LAST_HEARTBEAT_TIMESTAMP_FIELD_NUMBER: builtins.int
    @property
    def last_heartbeat_timestamp(self) -> google.protobuf.timestamp_pb2.Timestamp: ...
    def __init__(
        self,
        *,
        last_heartbeat_timestamp: google.protobuf.timestamp_pb2.Timestamp | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["last_heartbeat_timestamp", b"last_heartbeat_timestamp"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["last_heartbeat_timestamp", b"last_heartbeat_timestamp"]) -> None: ...

global___GetClusterStatusResp = GetClusterStatusResp
