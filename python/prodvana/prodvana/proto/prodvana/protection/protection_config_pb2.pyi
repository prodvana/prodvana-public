"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import builtins
import collections.abc
import google.protobuf.descriptor
import google.protobuf.duration_pb2
import google.protobuf.internal.containers
import google.protobuf.message
import prodvana.proto.prodvana.common_config.env_pb2
import prodvana.proto.prodvana.common_config.kubernetes_config_pb2
import prodvana.proto.prodvana.common_config.parameters_pb2
import prodvana.proto.prodvana.common_config.task_pb2
import prodvana.proto.prodvana.protection.builtins_pb2
import prodvana.proto.prodvana.runtimes.runtimes_config_pb2
import sys

if sys.version_info >= (3, 8):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

class BuiltinProtectionConfig(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    COMMIT_DENYLIST_FIELD_NUMBER: builtins.int
    ALLOWED_TIMES_FIELD_NUMBER: builtins.int
    @property
    def commit_denylist(self) -> prodvana.proto.prodvana.protection.builtins_pb2.CommitDenylistProtectionConfig: ...
    @property
    def allowed_times(self) -> prodvana.proto.prodvana.protection.builtins_pb2.AllowedTimesProtectionConfig: ...
    def __init__(
        self,
        *,
        commit_denylist: prodvana.proto.prodvana.protection.builtins_pb2.CommitDenylistProtectionConfig | None = ...,
        allowed_times: prodvana.proto.prodvana.protection.builtins_pb2.AllowedTimesProtectionConfig | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["allowed_times", b"allowed_times", "builtin_oneof", b"builtin_oneof", "commit_denylist", b"commit_denylist"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["allowed_times", b"allowed_times", "builtin_oneof", b"builtin_oneof", "commit_denylist", b"commit_denylist"]) -> None: ...
    def WhichOneof(self, oneof_group: typing_extensions.Literal["builtin_oneof", b"builtin_oneof"]) -> typing_extensions.Literal["commit_denylist", "allowed_times"] | None: ...

global___BuiltinProtectionConfig = BuiltinProtectionConfig

class ProtectionConfig(google.protobuf.message.Message):
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
    TASK_CONFIG_FIELD_NUMBER: builtins.int
    KUBERNETES_CONFIG_FIELD_NUMBER: builtins.int
    BUILTIN_FIELD_NUMBER: builtins.int
    POLL_INTERVAL_FIELD_NUMBER: builtins.int
    TIMEOUT_FIELD_NUMBER: builtins.int
    PARAMETERS_FIELD_NUMBER: builtins.int
    ENV_FIELD_NUMBER: builtins.int
    name: builtins.str
    @property
    def task_config(self) -> prodvana.proto.prodvana.common_config.task_pb2.TaskConfig:
        """Inline task config with retry, template support."""
    @property
    def kubernetes_config(self) -> prodvana.proto.prodvana.common_config.kubernetes_config_pb2.KubernetesConfig: ...
    @property
    def builtin(self) -> global___BuiltinProtectionConfig:
        """Other options here:
        - Ref to external repository of protection definitions.
        """
    @property
    def poll_interval(self) -> google.protobuf.duration_pb2.Duration:
        """customize intervals instead of using Prodvana default
        how often to run check even if it succeeds
        """
    @property
    def timeout(self) -> google.protobuf.duration_pb2.Duration:
        """deprecated, not used"""
    @property
    def parameters(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[prodvana.proto.prodvana.common_config.parameters_pb2.ParameterDefinition]: ...
    @property
    def env(self) -> google.protobuf.internal.containers.MessageMap[builtins.str, prodvana.proto.prodvana.common_config.env_pb2.EnvValue]:
        """optional env variables to inject and override from exec_config"""
    def __init__(
        self,
        *,
        name: builtins.str = ...,
        task_config: prodvana.proto.prodvana.common_config.task_pb2.TaskConfig | None = ...,
        kubernetes_config: prodvana.proto.prodvana.common_config.kubernetes_config_pb2.KubernetesConfig | None = ...,
        builtin: global___BuiltinProtectionConfig | None = ...,
        poll_interval: google.protobuf.duration_pb2.Duration | None = ...,
        timeout: google.protobuf.duration_pb2.Duration | None = ...,
        parameters: collections.abc.Iterable[prodvana.proto.prodvana.common_config.parameters_pb2.ParameterDefinition] | None = ...,
        env: collections.abc.Mapping[builtins.str, prodvana.proto.prodvana.common_config.env_pb2.EnvValue] | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["builtin", b"builtin", "exec_config", b"exec_config", "kubernetes_config", b"kubernetes_config", "poll_interval", b"poll_interval", "task_config", b"task_config", "timeout", b"timeout"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["builtin", b"builtin", "env", b"env", "exec_config", b"exec_config", "kubernetes_config", b"kubernetes_config", "name", b"name", "parameters", b"parameters", "poll_interval", b"poll_interval", "task_config", b"task_config", "timeout", b"timeout"]) -> None: ...
    def WhichOneof(self, oneof_group: typing_extensions.Literal["exec_config", b"exec_config"]) -> typing_extensions.Literal["task_config", "kubernetes_config", "builtin"] | None: ...

global___ProtectionConfig = ProtectionConfig

class CompiledProtectionAttachmentConfig(google.protobuf.message.Message):
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

    CONFIG_FIELD_NUMBER: builtins.int
    ATTACHMENT_FIELD_NUMBER: builtins.int
    RUNTIME_EXECUTION_FIELD_NUMBER: builtins.int
    ENV_FIELD_NUMBER: builtins.int
    PARAMETER_VALUES_FIELD_NUMBER: builtins.int
    @property
    def config(self) -> global___ProtectionConfig: ...
    @property
    def attachment(self) -> global___ProtectionAttachment:
        """Protection source - where did this protection get attached from (service/app/org/...)?"""
    @property
    def runtime_execution(self) -> prodvana.proto.prodvana.runtimes.runtimes_config_pb2.RuntimeExecutionConfig: ...
    @property
    def env(self) -> google.protobuf.internal.containers.MessageMap[builtins.str, prodvana.proto.prodvana.common_config.env_pb2.EnvValue]:
        """The compiled environment for this attachment's context, e.g.  Release Channel."""
    @property
    def parameter_values(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[prodvana.proto.prodvana.common_config.parameters_pb2.ParameterValue]:
        """compiled parameter values"""
    def __init__(
        self,
        *,
        config: global___ProtectionConfig | None = ...,
        attachment: global___ProtectionAttachment | None = ...,
        runtime_execution: prodvana.proto.prodvana.runtimes.runtimes_config_pb2.RuntimeExecutionConfig | None = ...,
        env: collections.abc.Mapping[builtins.str, prodvana.proto.prodvana.common_config.env_pb2.EnvValue] | None = ...,
        parameter_values: collections.abc.Iterable[prodvana.proto.prodvana.common_config.parameters_pb2.ParameterValue] | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["attachment", b"attachment", "config", b"config", "runtime_execution", b"runtime_execution"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["attachment", b"attachment", "config", b"config", "env", b"env", "parameter_values", b"parameter_values", "runtime_execution", b"runtime_execution"]) -> None: ...

global___CompiledProtectionAttachmentConfig = CompiledProtectionAttachmentConfig

class ServiceInstanceAttachment(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    SERVICE_FIELD_NUMBER: builtins.int
    RELEASE_CHANNEL_FIELD_NUMBER: builtins.int
    APPLICATION_FIELD_NUMBER: builtins.int
    service: builtins.str
    release_channel: builtins.str
    application: builtins.str
    def __init__(
        self,
        *,
        service: builtins.str = ...,
        release_channel: builtins.str = ...,
        application: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["application", b"application", "release_channel", b"release_channel", "service", b"service"]) -> None: ...

global___ServiceInstanceAttachment = ServiceInstanceAttachment

class ReleaseChannelAttachment(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    APPLICATION_FIELD_NUMBER: builtins.int
    RELEASE_CHANNEL_FIELD_NUMBER: builtins.int
    application: builtins.str
    release_channel: builtins.str
    def __init__(
        self,
        *,
        application: builtins.str = ...,
        release_channel: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["application", b"application", "release_channel", b"release_channel"]) -> None: ...

global___ReleaseChannelAttachment = ReleaseChannelAttachment

class ConvergenceAttachment(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    DESIRED_STATE_ID_FIELD_NUMBER: builtins.int
    desired_state_id: builtins.str
    def __init__(
        self,
        *,
        desired_state_id: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["desired_state_id", b"desired_state_id"]) -> None: ...

global___ConvergenceAttachment = ConvergenceAttachment

class ProtectionAttachment(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    SERVICE_INSTANCE_FIELD_NUMBER: builtins.int
    RELEASE_CHANNEL_FIELD_NUMBER: builtins.int
    CONVERGENCE_FIELD_NUMBER: builtins.int
    @property
    def service_instance(self) -> global___ServiceInstanceAttachment: ...
    @property
    def release_channel(self) -> global___ReleaseChannelAttachment: ...
    @property
    def convergence(self) -> global___ConvergenceAttachment: ...
    def __init__(
        self,
        *,
        service_instance: global___ServiceInstanceAttachment | None = ...,
        release_channel: global___ReleaseChannelAttachment | None = ...,
        convergence: global___ConvergenceAttachment | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["attachment", b"attachment", "convergence", b"convergence", "release_channel", b"release_channel", "service_instance", b"service_instance"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["attachment", b"attachment", "convergence", b"convergence", "release_channel", b"release_channel", "service_instance", b"service_instance"]) -> None: ...
    def WhichOneof(self, oneof_group: typing_extensions.Literal["attachment", b"attachment"]) -> typing_extensions.Literal["service_instance", "release_channel", "convergence"] | None: ...

global___ProtectionAttachment = ProtectionAttachment
