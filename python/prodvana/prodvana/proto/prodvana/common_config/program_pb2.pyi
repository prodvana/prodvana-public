"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import builtins
import collections.abc
import google.protobuf.descriptor
import google.protobuf.duration_pb2
import google.protobuf.internal.containers
import google.protobuf.internal.enum_type_wrapper
import google.protobuf.message
import prodvana.proto.prodvana.common_config.env_pb2
import sys
import typing

if sys.version_info >= (3, 10):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

class PortConfig(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    class _Protocol:
        ValueType = typing.NewType("ValueType", builtins.int)
        V: typing_extensions.TypeAlias = ValueType

    class _ProtocolEnumTypeWrapper(google.protobuf.internal.enum_type_wrapper._EnumTypeWrapper[PortConfig._Protocol.ValueType], builtins.type):  # noqa: F821
        DESCRIPTOR: google.protobuf.descriptor.EnumDescriptor
        UNKNOWN: PortConfig._Protocol.ValueType  # 0
        HTTP: PortConfig._Protocol.ValueType  # 1
        HTTP2: PortConfig._Protocol.ValueType  # 2
        GRPC: PortConfig._Protocol.ValueType  # 3
        TCP: PortConfig._Protocol.ValueType  # 4

    class Protocol(_Protocol, metaclass=_ProtocolEnumTypeWrapper): ...
    UNKNOWN: PortConfig.Protocol.ValueType  # 0
    HTTP: PortConfig.Protocol.ValueType  # 1
    HTTP2: PortConfig.Protocol.ValueType  # 2
    GRPC: PortConfig.Protocol.ValueType  # 3
    TCP: PortConfig.Protocol.ValueType  # 4

    PORT_FIELD_NUMBER: builtins.int
    TARGET_PORT_FIELD_NUMBER: builtins.int
    EXTERNAL_FIELD_NUMBER: builtins.int
    PROTOCOL_FIELD_NUMBER: builtins.int
    TLS_FIELD_NUMBER: builtins.int
    port: builtins.int
    target_port: builtins.int
    """optional, default to the same value as `port`"""
    external: builtins.bool
    """if this port should be exposed to the public internet"""
    protocol: global___PortConfig.Protocol.ValueType
    tls: builtins.bool
    """if the connection should use TLS"""
    def __init__(
        self,
        *,
        port: builtins.int = ...,
        target_port: builtins.int = ...,
        external: builtins.bool = ...,
        protocol: global___PortConfig.Protocol.ValueType = ...,
        tls: builtins.bool = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["external", b"external", "port", b"port", "protocol", b"protocol", "target_port", b"target_port", "tls", b"tls"]) -> None: ...

global___PortConfig = PortConfig

class ResourceList(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    MEMORY_FIELD_NUMBER: builtins.int
    CPU_FIELD_NUMBER: builtins.int
    memory: builtins.str
    cpu: builtins.str
    def __init__(
        self,
        *,
        memory: builtins.str = ...,
        cpu: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["cpu", b"cpu", "memory", b"memory"]) -> None: ...

global___ResourceList = ResourceList

class ResourceRequirements(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    REQUESTS_FIELD_NUMBER: builtins.int
    @property
    def requests(self) -> global___ResourceList: ...
    def __init__(
        self,
        *,
        requests: global___ResourceList | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["requests", b"requests"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["requests", b"requests"]) -> None: ...

global___ResourceRequirements = ResourceRequirements

class HttpProbe(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    PATH_FIELD_NUMBER: builtins.int
    PORT_FIELD_NUMBER: builtins.int
    TLS_FIELD_NUMBER: builtins.int
    path: builtins.str
    port: builtins.int
    tls: builtins.bool
    def __init__(
        self,
        *,
        path: builtins.str = ...,
        port: builtins.int = ...,
        tls: builtins.bool = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["path", b"path", "port", b"port", "tls", b"tls"]) -> None: ...

global___HttpProbe = HttpProbe

class CmdProbe(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    COMMAND_FIELD_NUMBER: builtins.int
    @property
    def command(self) -> google.protobuf.internal.containers.RepeatedScalarFieldContainer[builtins.str]: ...
    def __init__(
        self,
        *,
        command: collections.abc.Iterable[builtins.str] | None = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["command", b"command"]) -> None: ...

global___CmdProbe = CmdProbe

class TcpProbe(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    PORT_FIELD_NUMBER: builtins.int
    HOST_FIELD_NUMBER: builtins.int
    port: builtins.int
    host: builtins.str
    def __init__(
        self,
        *,
        port: builtins.int = ...,
        host: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["host", b"host", "port", b"port"]) -> None: ...

global___TcpProbe = TcpProbe

class HealthCheck(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    HTTP_FIELD_NUMBER: builtins.int
    CMD_FIELD_NUMBER: builtins.int
    TCP_FIELD_NUMBER: builtins.int
    DELAY_FIELD_NUMBER: builtins.int
    PERIOD_FIELD_NUMBER: builtins.int
    @property
    def http(self) -> global___HttpProbe: ...
    @property
    def cmd(self) -> global___CmdProbe: ...
    @property
    def tcp(self) -> global___TcpProbe: ...
    @property
    def delay(self) -> google.protobuf.duration_pb2.Duration: ...
    @property
    def period(self) -> google.protobuf.duration_pb2.Duration: ...
    def __init__(
        self,
        *,
        http: global___HttpProbe | None = ...,
        cmd: global___CmdProbe | None = ...,
        tcp: global___TcpProbe | None = ...,
        delay: google.protobuf.duration_pb2.Duration | None = ...,
        period: google.protobuf.duration_pb2.Duration | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["cmd", b"cmd", "delay", b"delay", "http", b"http", "period", b"period", "probe_config", b"probe_config", "tcp", b"tcp"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["cmd", b"cmd", "delay", b"delay", "http", b"http", "period", b"period", "probe_config", b"probe_config", "tcp", b"tcp"]) -> None: ...
    def WhichOneof(self, oneof_group: typing_extensions.Literal["probe_config", b"probe_config"]) -> typing_extensions.Literal["http", "cmd", "tcp"] | None: ...

global___HealthCheck = HealthCheck

class ImageRegistryInfo(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    CONTAINER_REGISTRY_ID_FIELD_NUMBER: builtins.int
    CONTAINER_REGISTRY_FIELD_NUMBER: builtins.int
    IMAGE_REPOSITORY_FIELD_NUMBER: builtins.int
    container_registry_id: builtins.str
    """both container_registry_id and container_registry are supported, but only set one not both"""
    container_registry: builtins.str
    image_repository: builtins.str
    """Not the URL, everything after: username/reponame
    leaving it as once field because the spec considers it 1 field
    in theory you can have no username, or nest the repo names
    """
    def __init__(
        self,
        *,
        container_registry_id: builtins.str = ...,
        container_registry: builtins.str = ...,
        image_repository: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["container_registry", b"container_registry", "container_registry_id", b"container_registry_id", "image_repository", b"image_repository"]) -> None: ...

global___ImageRegistryInfo = ImageRegistryInfo

class ImageDetails(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    REGISTRY_INFO_FIELD_NUMBER: builtins.int
    IDENTIFIER_FIELD_NUMBER: builtins.int
    @property
    def registry_info(self) -> global___ImageRegistryInfo: ...
    identifier: builtins.str
    """image tag or digest"""
    def __init__(
        self,
        *,
        registry_info: global___ImageRegistryInfo | None = ...,
        identifier: builtins.str = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["registry_info", b"registry_info"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["identifier", b"identifier", "registry_info", b"registry_info"]) -> None: ...

global___ImageDetails = ImageDetails

class ProgramConfig(google.protobuf.message.Message):
    """This ProgramConfig is used in multiple places, including service configuration and custom pipeline tasks.
    The validations expressed here represent the lowest common denominator of the use cases.
    """

    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    class EnvEntry(google.protobuf.message.Message):
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

    NAME_FIELD_NUMBER: builtins.int
    IMAGE_FIELD_NUMBER: builtins.int
    IMAGE_TAG_FIELD_NUMBER: builtins.int
    IMAGE_REGISTRY_INFO_FIELD_NUMBER: builtins.int
    CMD_FIELD_NUMBER: builtins.int
    ENTRYPOINT_FIELD_NUMBER: builtins.int
    ENV_FIELD_NUMBER: builtins.int
    RESOURCES_FIELD_NUMBER: builtins.int
    HEALTH_CHECK_FIELD_NUMBER: builtins.int
    PORTS_FIELD_NUMBER: builtins.int
    TEMPLATE_COMPLETE_FIELD_NUMBER: builtins.int
    WORKING_DIRECTORY_FIELD_NUMBER: builtins.int
    name: builtins.str
    image: builtins.str
    image_tag: builtins.str
    @property
    def image_registry_info(self) -> global___ImageRegistryInfo:
        """optional, not guaranteed to be compatible with `image` (e.g. if user decides to paste in a public image string)"""
    @property
    def cmd(self) -> google.protobuf.internal.containers.RepeatedScalarFieldContainer[builtins.str]: ...
    @property
    def entrypoint(self) -> google.protobuf.internal.containers.RepeatedScalarFieldContainer[builtins.str]: ...
    @property
    def env(self) -> google.protobuf.internal.containers.MessageMap[builtins.str, prodvana.proto.prodvana.common_config.env_pb2.EnvValue]: ...
    @property
    def resources(self) -> global___ResourceRequirements: ...
    @property
    def health_check(self) -> global___HealthCheck: ...
    @property
    def ports(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[global___PortConfig]:
        """empty list is ok, default from docker image will be used"""
    template_complete: builtins.bool
    """Only set when this ServiceConfig represents a ServiceTemplate
    When set, this program will be added in whole to the target service's
    programs in full, and will not partially apply to a program with the
    same name in the target service's config.
    """
    working_directory: builtins.str
    """working directory, defaults to runtime's implementation (usually the default working directory in the docker image)"""
    def __init__(
        self,
        *,
        name: builtins.str = ...,
        image: builtins.str = ...,
        image_tag: builtins.str = ...,
        image_registry_info: global___ImageRegistryInfo | None = ...,
        cmd: collections.abc.Iterable[builtins.str] | None = ...,
        entrypoint: collections.abc.Iterable[builtins.str] | None = ...,
        env: collections.abc.Mapping[builtins.str, prodvana.proto.prodvana.common_config.env_pb2.EnvValue] | None = ...,
        resources: global___ResourceRequirements | None = ...,
        health_check: global___HealthCheck | None = ...,
        ports: collections.abc.Iterable[global___PortConfig] | None = ...,
        template_complete: builtins.bool = ...,
        working_directory: builtins.str = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["health_check", b"health_check", "image_registry_info", b"image_registry_info", "resources", b"resources"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["cmd", b"cmd", "entrypoint", b"entrypoint", "env", b"env", "health_check", b"health_check", "image", b"image", "image_registry_info", b"image_registry_info", "image_tag", b"image_tag", "name", b"name", "ports", b"ports", "resources", b"resources", "template_complete", b"template_complete", "working_directory", b"working_directory"]) -> None: ...

global___ProgramConfig = ProgramConfig

class PerReleaseChannelProgramConfig(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    class EnvEntry(google.protobuf.message.Message):
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

    NAME_FIELD_NUMBER: builtins.int
    ENV_FIELD_NUMBER: builtins.int
    IMAGE_FIELD_NUMBER: builtins.int
    IMAGE_REGISTRY_INFO_FIELD_NUMBER: builtins.int
    IMAGE_TAG_FIELD_NUMBER: builtins.int
    name: builtins.str
    """must match a program in top-level program config"""
    @property
    def env(self) -> google.protobuf.internal.containers.MessageMap[builtins.str, prodvana.proto.prodvana.common_config.env_pb2.EnvValue]: ...
    image: builtins.str
    """If present, overrides main program config.
    TODO: If users need to enforce that different RCs require different images, add requires_separate_images option to ProgramConfig.
    This will prevent accidentally undoing separate images via the UI or autopush.
    """
    @property
    def image_registry_info(self) -> global___ImageRegistryInfo:
        """optional, not guaranteed to be compatible with `image` (e.g. if user decides to paste in a public image string)"""
    image_tag: builtins.str
    def __init__(
        self,
        *,
        name: builtins.str = ...,
        env: collections.abc.Mapping[builtins.str, prodvana.proto.prodvana.common_config.env_pb2.EnvValue] | None = ...,
        image: builtins.str = ...,
        image_registry_info: global___ImageRegistryInfo | None = ...,
        image_tag: builtins.str = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["image_registry_info", b"image_registry_info"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["env", b"env", "image", b"image", "image_registry_info", b"image_registry_info", "image_tag", b"image_tag", "name", b"name"]) -> None: ...

global___PerReleaseChannelProgramConfig = PerReleaseChannelProgramConfig
