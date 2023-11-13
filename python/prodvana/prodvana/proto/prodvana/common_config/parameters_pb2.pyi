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
import prodvana.proto.prodvana.common_config.program_pb2
import sys

if sys.version_info >= (3, 8):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

class ProgramChange(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    NAME_FIELD_NUMBER: builtins.int
    name: builtins.str
    def __init__(
        self,
        *,
        name: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["name", b"name"]) -> None: ...

global___ProgramChange = ProgramChange

class StringParameterDefinition(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    DEFAULT_VALUE_FIELD_NUMBER: builtins.int
    default_value: builtins.str
    def __init__(
        self,
        *,
        default_value: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["default_value", b"default_value"]) -> None: ...

global___StringParameterDefinition = StringParameterDefinition

class DockerImageParameterDefinition(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    class Changes(google.protobuf.message.Message):
        DESCRIPTOR: google.protobuf.descriptor.Descriptor

        PROGRAM_FIELD_NUMBER: builtins.int
        @property
        def program(self) -> global___ProgramChange: ...
        def __init__(
            self,
            *,
            program: global___ProgramChange | None = ...,
        ) -> None: ...
        def HasField(self, field_name: typing_extensions.Literal["oneof", b"oneof", "program", b"program"]) -> builtins.bool: ...
        def ClearField(self, field_name: typing_extensions.Literal["oneof", b"oneof", "program", b"program"]) -> None: ...
        def WhichOneof(self, oneof_group: typing_extensions.Literal["oneof", b"oneof"]) -> typing_extensions.Literal["program"] | None: ...

    DEFAULT_TAG_FIELD_NUMBER: builtins.int
    IMAGE_REGISTRY_INFO_FIELD_NUMBER: builtins.int
    CHANGES_FIELD_NUMBER: builtins.int
    default_tag: builtins.str
    """can be omitted if parameter is required"""
    @property
    def image_registry_info(self) -> prodvana.proto.prodvana.common_config.program_pb2.ImageRegistryInfo: ...
    @property
    def changes(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[global___DockerImageParameterDefinition.Changes]: ...
    def __init__(
        self,
        *,
        default_tag: builtins.str = ...,
        image_registry_info: prodvana.proto.prodvana.common_config.program_pb2.ImageRegistryInfo | None = ...,
        changes: collections.abc.Iterable[global___DockerImageParameterDefinition.Changes] | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["image_registry_info", b"image_registry_info"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["changes", b"changes", "default_tag", b"default_tag", "image_registry_info", b"image_registry_info"]) -> None: ...

global___DockerImageParameterDefinition = DockerImageParameterDefinition

class FixedReplicaChange(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    def __init__(
        self,
    ) -> None: ...

global___FixedReplicaChange = FixedReplicaChange

class IntParameterDefinition(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    class Changes(google.protobuf.message.Message):
        DESCRIPTOR: google.protobuf.descriptor.Descriptor

        FIXED_REPLICA_FIELD_NUMBER: builtins.int
        @property
        def fixed_replica(self) -> global___FixedReplicaChange: ...
        def __init__(
            self,
            *,
            fixed_replica: global___FixedReplicaChange | None = ...,
        ) -> None: ...
        def HasField(self, field_name: typing_extensions.Literal["fixed_replica", b"fixed_replica", "oneof", b"oneof"]) -> builtins.bool: ...
        def ClearField(self, field_name: typing_extensions.Literal["fixed_replica", b"fixed_replica", "oneof", b"oneof"]) -> None: ...
        def WhichOneof(self, oneof_group: typing_extensions.Literal["oneof", b"oneof"]) -> typing_extensions.Literal["fixed_replica"] | None: ...

    DEFAULT_VALUE_FIELD_NUMBER: builtins.int
    CHANGES_FIELD_NUMBER: builtins.int
    default_value: builtins.int
    @property
    def changes(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[global___IntParameterDefinition.Changes]: ...
    def __init__(
        self,
        *,
        default_value: builtins.int = ...,
        changes: collections.abc.Iterable[global___IntParameterDefinition.Changes] | None = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["changes", b"changes", "default_value", b"default_value"]) -> None: ...

global___IntParameterDefinition = IntParameterDefinition

class SecretParameterDefinition(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    def __init__(
        self,
    ) -> None: ...

global___SecretParameterDefinition = SecretParameterDefinition

class CommitParameterDefinition(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    REPOSITORY_FIELD_NUMBER: builtins.int
    BRANCH_FIELD_NUMBER: builtins.int
    repository: builtins.str
    branch: builtins.str
    """branch to pull commits from, defaults to `main`"""
    def __init__(
        self,
        *,
        repository: builtins.str = ...,
        branch: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["branch", b"branch", "repository", b"repository"]) -> None: ...

global___CommitParameterDefinition = CommitParameterDefinition

class ParameterDefinition(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    NAME_FIELD_NUMBER: builtins.int
    DESCRIPTION_FIELD_NUMBER: builtins.int
    STRING_FIELD_NUMBER: builtins.int
    DOCKER_IMAGE_FIELD_NUMBER: builtins.int
    INT_FIELD_NUMBER: builtins.int
    SECRET_FIELD_NUMBER: builtins.int
    COMMIT_FIELD_NUMBER: builtins.int
    REQUIRED_FIELD_NUMBER: builtins.int
    name: builtins.str
    """parameter name, used in substitutions"""
    description: builtins.str
    """optional description for display purposes"""
    @property
    def string(self) -> global___StringParameterDefinition: ...
    @property
    def docker_image(self) -> global___DockerImageParameterDefinition: ...
    @property
    def int(self) -> global___IntParameterDefinition: ...
    @property
    def secret(self) -> global___SecretParameterDefinition: ...
    @property
    def commit(self) -> global___CommitParameterDefinition: ...
    required: builtins.bool
    """next: 9"""
    def __init__(
        self,
        *,
        name: builtins.str = ...,
        description: builtins.str = ...,
        string: global___StringParameterDefinition | None = ...,
        docker_image: global___DockerImageParameterDefinition | None = ...,
        int: global___IntParameterDefinition | None = ...,
        secret: global___SecretParameterDefinition | None = ...,
        commit: global___CommitParameterDefinition | None = ...,
        required: builtins.bool = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["commit", b"commit", "config_oneof", b"config_oneof", "docker_image", b"docker_image", "int", b"int", "secret", b"secret", "string", b"string"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["commit", b"commit", "config_oneof", b"config_oneof", "description", b"description", "docker_image", b"docker_image", "int", b"int", "name", b"name", "required", b"required", "secret", b"secret", "string", b"string"]) -> None: ...
    def WhichOneof(self, oneof_group: typing_extensions.Literal["config_oneof", b"config_oneof"]) -> typing_extensions.Literal["string", "docker_image", "int", "secret", "commit"] | None: ...

global___ParameterDefinition = ParameterDefinition

class SecretParameterValue(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    RAW_SECRET_FIELD_NUMBER: builtins.int
    SECRET_REF_FIELD_NUMBER: builtins.int
    raw_secret: builtins.str
    """Raw secret value"""
    @property
    def secret_ref(self) -> prodvana.proto.prodvana.common_config.env_pb2.Secret:
        """Existing secret reference."""
    def __init__(
        self,
        *,
        raw_secret: builtins.str = ...,
        secret_ref: prodvana.proto.prodvana.common_config.env_pb2.Secret | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["raw_secret", b"raw_secret", "secret_oneof", b"secret_oneof", "secret_ref", b"secret_ref"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["raw_secret", b"raw_secret", "secret_oneof", b"secret_oneof", "secret_ref", b"secret_ref"]) -> None: ...
    def WhichOneof(self, oneof_group: typing_extensions.Literal["secret_oneof", b"secret_oneof"]) -> typing_extensions.Literal["raw_secret", "secret_ref"] | None: ...

global___SecretParameterValue = SecretParameterValue

class ParameterValue(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    NAME_FIELD_NUMBER: builtins.int
    STRING_FIELD_NUMBER: builtins.int
    INT_FIELD_NUMBER: builtins.int
    DOCKER_IMAGE_TAG_FIELD_NUMBER: builtins.int
    SECRET_FIELD_NUMBER: builtins.int
    COMMIT_FIELD_NUMBER: builtins.int
    name: builtins.str
    string: builtins.str
    int: builtins.int
    docker_image_tag: builtins.str
    @property
    def secret(self) -> global___SecretParameterValue: ...
    commit: builtins.str
    def __init__(
        self,
        *,
        name: builtins.str = ...,
        string: builtins.str = ...,
        int: builtins.int = ...,
        docker_image_tag: builtins.str = ...,
        secret: global___SecretParameterValue | None = ...,
        commit: builtins.str = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["commit", b"commit", "docker_image_tag", b"docker_image_tag", "int", b"int", "secret", b"secret", "string", b"string", "value_oneof", b"value_oneof"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["commit", b"commit", "docker_image_tag", b"docker_image_tag", "int", b"int", "name", b"name", "secret", b"secret", "string", b"string", "value_oneof", b"value_oneof"]) -> None: ...
    def WhichOneof(self, oneof_group: typing_extensions.Literal["value_oneof", b"value_oneof"]) -> typing_extensions.Literal["string", "int", "docker_image_tag", "secret", "commit"] | None: ...

global___ParameterValue = ParameterValue

class ParametersConfig(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    PARAMETERS_FIELD_NUMBER: builtins.int
    @property
    def parameters(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[global___ParameterDefinition]: ...
    def __init__(
        self,
        *,
        parameters: collections.abc.Iterable[global___ParameterDefinition] | None = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["parameters", b"parameters"]) -> None: ...

global___ParametersConfig = ParametersConfig
