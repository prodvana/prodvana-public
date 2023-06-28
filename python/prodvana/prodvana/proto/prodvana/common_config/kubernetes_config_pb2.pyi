"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import builtins
import collections.abc
import google.protobuf.descriptor
import google.protobuf.internal.containers
import google.protobuf.internal.enum_type_wrapper
import google.protobuf.message
import sys
import typing

if sys.version_info >= (3, 10):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

class LocalConfig(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    PATH_FIELD_NUMBER: builtins.int
    PATHS_FIELD_NUMBER: builtins.int
    path: builtins.str
    """Specify a path to a local file or directory"""
    @property
    def paths(self) -> google.protobuf.internal.containers.RepeatedScalarFieldContainer[builtins.str]:
        """Specify multiple paths. They will get concatenated."""
    def __init__(
        self,
        *,
        path: builtins.str = ...,
        paths: collections.abc.Iterable[builtins.str] | None = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["path", b"path", "paths", b"paths"]) -> None: ...

global___LocalConfig = LocalConfig

class KubernetesConfig(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    class _Type:
        ValueType = typing.NewType("ValueType", builtins.int)
        V: typing_extensions.TypeAlias = ValueType

    class _TypeEnumTypeWrapper(google.protobuf.internal.enum_type_wrapper._EnumTypeWrapper[KubernetesConfig._Type.ValueType], builtins.type):  # noqa: F821
        DESCRIPTOR: google.protobuf.descriptor.EnumDescriptor
        UNKNOWN: KubernetesConfig._Type.ValueType  # 0
        KUBERNETES: KubernetesConfig._Type.ValueType  # 1
        KUSTOMIZE: KubernetesConfig._Type.ValueType  # 2

    class Type(_Type, metaclass=_TypeEnumTypeWrapper): ...
    UNKNOWN: KubernetesConfig.Type.ValueType  # 0
    KUBERNETES: KubernetesConfig.Type.ValueType  # 1
    KUSTOMIZE: KubernetesConfig.Type.ValueType  # 2

    class _EnvInjectionMode:
        ValueType = typing.NewType("ValueType", builtins.int)
        V: typing_extensions.TypeAlias = ValueType

    class _EnvInjectionModeEnumTypeWrapper(google.protobuf.internal.enum_type_wrapper._EnumTypeWrapper[KubernetesConfig._EnvInjectionMode.ValueType], builtins.type):  # noqa: F821
        DESCRIPTOR: google.protobuf.descriptor.EnumDescriptor
        ENV_INJECT_UNKNOWN: KubernetesConfig._EnvInjectionMode.ValueType  # 0
        ENV_INJECT_DISABLED: KubernetesConfig._EnvInjectionMode.ValueType  # 1
        """disables env injection entirely"""
        ENV_INJECT_NON_SECRET_ENV: KubernetesConfig._EnvInjectionMode.ValueType  # 2
        """inject non-secret env values from the Release Channel"""

    class EnvInjectionMode(_EnvInjectionMode, metaclass=_EnvInjectionModeEnumTypeWrapper): ...
    ENV_INJECT_UNKNOWN: KubernetesConfig.EnvInjectionMode.ValueType  # 0
    ENV_INJECT_DISABLED: KubernetesConfig.EnvInjectionMode.ValueType  # 1
    """disables env injection entirely"""
    ENV_INJECT_NON_SECRET_ENV: KubernetesConfig.EnvInjectionMode.ValueType  # 2
    """inject non-secret env values from the Release Channel"""

    TYPE_FIELD_NUMBER: builtins.int
    INLINED_FIELD_NUMBER: builtins.int
    LOCAL_FIELD_NUMBER: builtins.int
    ENV_INJECTION_MODE_FIELD_NUMBER: builtins.int
    type: global___KubernetesConfig.Type.ValueType
    inlined: builtins.str
    @property
    def local(self) -> global___LocalConfig: ...
    env_injection_mode: global___KubernetesConfig.EnvInjectionMode.ValueType
    """Defaults to ENV_INJECT_NON_SECRET_ENV"""
    def __init__(
        self,
        *,
        type: global___KubernetesConfig.Type.ValueType = ...,
        inlined: builtins.str = ...,
        local: global___LocalConfig | None = ...,
        env_injection_mode: global___KubernetesConfig.EnvInjectionMode.ValueType = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["inlined", b"inlined", "local", b"local", "source_oneof", b"source_oneof"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["env_injection_mode", b"env_injection_mode", "inlined", b"inlined", "local", b"local", "source_oneof", b"source_oneof", "type", b"type"]) -> None: ...
    def WhichOneof(self, oneof_group: typing_extensions.Literal["source_oneof", b"source_oneof"]) -> typing_extensions.Literal["inlined", "local"] | None: ...

global___KubernetesConfig = KubernetesConfig
