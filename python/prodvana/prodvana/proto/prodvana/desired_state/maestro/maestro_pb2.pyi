"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import builtins
import google.protobuf.descriptor
import google.protobuf.internal.enum_type_wrapper
import google.protobuf.message
import google.protobuf.timestamp_pb2
import prodvana.proto.prodvana.desired_state.model.desired_state_pb2
import prodvana.proto.prodvana.object.meta_pb2
import sys
import typing

if sys.version_info >= (3, 10):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

class _Status:
    ValueType = typing.NewType("ValueType", builtins.int)
    V: typing_extensions.TypeAlias = ValueType

class _StatusEnumTypeWrapper(google.protobuf.internal.enum_type_wrapper._EnumTypeWrapper[_Status.ValueType], builtins.type):  # noqa: F821
    DESCRIPTOR: google.protobuf.descriptor.EnumDescriptor
    UNKNOWN_STATUS: _Status.ValueType  # 0
    PENDING: _Status.ValueType  # 1
    """Pending means this train entry has never been started and is waiting for dependencies to be met."""
    IN_PROGRESS: _Status.ValueType  # 2
    """Dependencies have been met and this train entry has been promoted."""
    SUCCEEDED: _Status.ValueType  # 3
    """Successfully converged."""
    FAILED: _Status.ValueType  # 4
    """Promoted but failed to converge."""
    SKIPPED: _Status.ValueType  # 5
    """This train entry was skipped - any dependencies on this entry can proceed."""

class Status(_Status, metaclass=_StatusEnumTypeWrapper): ...

UNKNOWN_STATUS: Status.ValueType  # 0
PENDING: Status.ValueType  # 1
"""Pending means this train entry has never been started and is waiting for dependencies to be met."""
IN_PROGRESS: Status.ValueType  # 2
"""Dependencies have been met and this train entry has been promoted."""
SUCCEEDED: Status.ValueType  # 3
"""Successfully converged."""
FAILED: Status.ValueType  # 4
"""Promoted but failed to converge."""
SKIPPED: Status.ValueType  # 5
"""This train entry was skipped - any dependencies on this entry can proceed."""
global___Status = Status

class _Strategy:
    ValueType = typing.NewType("ValueType", builtins.int)
    V: typing_extensions.TypeAlias = ValueType

class _StrategyEnumTypeWrapper(google.protobuf.internal.enum_type_wrapper._EnumTypeWrapper[_Strategy.ValueType], builtins.type):  # noqa: F821
    DESCRIPTOR: google.protobuf.descriptor.EnumDescriptor
    UNKNOWN_STRATEGY: _Strategy.ValueType  # 0
    IMMEDIATE: _Strategy.ValueType  # 1
    ON_STABLE: _Strategy.ValueType  # 2

class Strategy(_Strategy, metaclass=_StrategyEnumTypeWrapper): ...

UNKNOWN_STRATEGY: Strategy.ValueType  # 0
IMMEDIATE: Strategy.ValueType  # 1
ON_STABLE: Strategy.ValueType  # 2
global___Strategy = Strategy

class MaestroConfig(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    STRATEGY_FIELD_NUMBER: builtins.int
    strategy: global___Strategy.ValueType
    def __init__(
        self,
        *,
        strategy: global___Strategy.ValueType = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["strategy", b"strategy"]) -> None: ...

global___MaestroConfig = MaestroConfig

class MaestroReleaseConfig(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    ENTITY_ID_FIELD_NUMBER: builtins.int
    MAESTRO_CONFIG_FIELD_NUMBER: builtins.int
    DESIRED_STATE_FIELD_NUMBER: builtins.int
    CREATION_TIMESTAMP_FIELD_NUMBER: builtins.int
    @property
    def entity_id(self) -> prodvana.proto.prodvana.desired_state.model.desired_state_pb2.Identifier: ...
    @property
    def maestro_config(self) -> global___MaestroConfig: ...
    @property
    def desired_state(self) -> prodvana.proto.prodvana.desired_state.model.desired_state_pb2.State: ...
    @property
    def creation_timestamp(self) -> google.protobuf.timestamp_pb2.Timestamp: ...
    def __init__(
        self,
        *,
        entity_id: prodvana.proto.prodvana.desired_state.model.desired_state_pb2.Identifier | None = ...,
        maestro_config: global___MaestroConfig | None = ...,
        desired_state: prodvana.proto.prodvana.desired_state.model.desired_state_pb2.State | None = ...,
        creation_timestamp: google.protobuf.timestamp_pb2.Timestamp | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["creation_timestamp", b"creation_timestamp", "desired_state", b"desired_state", "entity_id", b"entity_id", "maestro_config", b"maestro_config"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["creation_timestamp", b"creation_timestamp", "desired_state", b"desired_state", "entity_id", b"entity_id", "maestro_config", b"maestro_config"]) -> None: ...

global___MaestroReleaseConfig = MaestroReleaseConfig

class MaestroReleaseState(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    STATUS_FIELD_NUMBER: builtins.int
    LAST_UPDATE_TIMESTAMP_FIELD_NUMBER: builtins.int
    status: global___Status.ValueType
    @property
    def last_update_timestamp(self) -> google.protobuf.timestamp_pb2.Timestamp: ...
    def __init__(
        self,
        *,
        status: global___Status.ValueType = ...,
        last_update_timestamp: google.protobuf.timestamp_pb2.Timestamp | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["last_update_timestamp", b"last_update_timestamp"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["last_update_timestamp", b"last_update_timestamp", "status", b"status"]) -> None: ...

global___MaestroReleaseState = MaestroReleaseState

class MaestroRelease(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    META_FIELD_NUMBER: builtins.int
    CONFIG_FIELD_NUMBER: builtins.int
    STATE_FIELD_NUMBER: builtins.int
    @property
    def meta(self) -> prodvana.proto.prodvana.object.meta_pb2.ObjectMeta: ...
    @property
    def config(self) -> global___MaestroReleaseConfig: ...
    @property
    def state(self) -> global___MaestroReleaseState: ...
    def __init__(
        self,
        *,
        meta: prodvana.proto.prodvana.object.meta_pb2.ObjectMeta | None = ...,
        config: global___MaestroReleaseConfig | None = ...,
        state: global___MaestroReleaseState | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["config", b"config", "meta", b"meta", "state", b"state"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["config", b"config", "meta", b"meta", "state", b"state"]) -> None: ...

global___MaestroRelease = MaestroRelease
