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
import prodvana.proto.prodvana.common_config.meta_pb2
import prodvana.proto.prodvana.common_config.version_push_pb2
import prodvana.proto.prodvana.object.meta_pb2
import prodvana.proto.prodvana.pipelines.pipelines_pb2
import sys

if sys.version_info >= (3, 8):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

class ServicePushParam(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    SERVICE_FIELD_NUMBER: builtins.int
    VERSION_PUSH_FIELD_NUMBER: builtins.int
    APPLICATION_FIELD_NUMBER: builtins.int
    service: builtins.str
    @property
    def version_push(self) -> prodvana.proto.prodvana.common_config.version_push_pb2.VersionPushOption: ...
    application: builtins.str
    def __init__(
        self,
        *,
        service: builtins.str = ...,
        version_push: prodvana.proto.prodvana.common_config.version_push_pb2.VersionPushOption | None = ...,
        application: builtins.str = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["version_push", b"version_push"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["application", b"application", "service", b"service", "version_push", b"version_push"]) -> None: ...

global___ServicePushParam = ServicePushParam

class ServiceInstancePush(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    APPLICATION_META_FIELD_NUMBER: builtins.int
    SERVICE_INSTANCE_META_FIELD_NUMBER: builtins.int
    TARGET_VERSION_FIELD_NUMBER: builtins.int
    @property
    def application_meta(self) -> prodvana.proto.prodvana.object.meta_pb2.ObjectMeta: ...
    @property
    def service_instance_meta(self) -> prodvana.proto.prodvana.common_config.meta_pb2.ServiceInstanceObjectMeta:
        """service starting version stored in service_instance_meta"""
    target_version: builtins.str
    """TODO(naphat/rohit) store rollback target information here
    string rollback_target_version = 4;
    """
    def __init__(
        self,
        *,
        application_meta: prodvana.proto.prodvana.object.meta_pb2.ObjectMeta | None = ...,
        service_instance_meta: prodvana.proto.prodvana.common_config.meta_pb2.ServiceInstanceObjectMeta | None = ...,
        target_version: builtins.str = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["application_meta", b"application_meta", "service_instance_meta", b"service_instance_meta"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["application_meta", b"application_meta", "service_instance_meta", b"service_instance_meta", "target_version", b"target_version"]) -> None: ...

global___ServiceInstancePush = ServiceInstancePush

class TaskStatus(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    ID_FIELD_NUMBER: builtins.int
    STATE_FIELD_NUMBER: builtins.int
    CREATION_TIMESTAMP_FIELD_NUMBER: builtins.int
    LAST_UPDATE_TIMESTAMP_FIELD_NUMBER: builtins.int
    id: builtins.int
    state: builtins.str
    @property
    def creation_timestamp(self) -> google.protobuf.timestamp_pb2.Timestamp: ...
    @property
    def last_update_timestamp(self) -> google.protobuf.timestamp_pb2.Timestamp: ...
    def __init__(
        self,
        *,
        id: builtins.int = ...,
        state: builtins.str = ...,
        creation_timestamp: google.protobuf.timestamp_pb2.Timestamp | None = ...,
        last_update_timestamp: google.protobuf.timestamp_pb2.Timestamp | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["creation_timestamp", b"creation_timestamp", "last_update_timestamp", b"last_update_timestamp"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["creation_timestamp", b"creation_timestamp", "id", b"id", "last_update_timestamp", b"last_update_timestamp", "state", b"state"]) -> None: ...

global___TaskStatus = TaskStatus

class PipelineState(google.protobuf.message.Message):
    """TODO(naphat) put stuff here"""

    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    def __init__(
        self,
    ) -> None: ...

global___PipelineState = PipelineState

class Pipeline(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    META_FIELD_NUMBER: builtins.int
    CONFIG_FIELD_NUMBER: builtins.int
    STATE_FIELD_NUMBER: builtins.int
    @property
    def meta(self) -> prodvana.proto.prodvana.object.meta_pb2.ObjectMeta: ...
    @property
    def config(self) -> prodvana.proto.prodvana.pipelines.pipelines_pb2.PipelineConfig: ...
    @property
    def state(self) -> global___PipelineState: ...
    def __init__(
        self,
        *,
        meta: prodvana.proto.prodvana.object.meta_pb2.ObjectMeta | None = ...,
        config: prodvana.proto.prodvana.pipelines.pipelines_pb2.PipelineConfig | None = ...,
        state: global___PipelineState | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["config", b"config", "meta", b"meta", "state", b"state"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["config", b"config", "meta", b"meta", "state", b"state"]) -> None: ...

global___Pipeline = Pipeline

class PipelineRunConfig(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    PIPELINE_CONFIG_FIELD_NUMBER: builtins.int
    SERVICE_PUSH_PARAMS_FIELD_NUMBER: builtins.int
    SKIP_ROLLBACK_VALIDATION_FIELD_NUMBER: builtins.int
    SERVICE_INSTANCE_PUSHES_FIELD_NUMBER: builtins.int
    SKIP_ALERT_CHECKS_FIELD_NUMBER: builtins.int
    @property
    def pipeline_config(self) -> prodvana.proto.prodvana.pipelines.pipelines_pb2.PipelineConfig: ...
    @property
    def service_push_params(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[global___ServicePushParam]: ...
    skip_rollback_validation: builtins.bool
    @property
    def service_instance_pushes(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[global___ServiceInstancePush]: ...
    skip_alert_checks: builtins.bool
    def __init__(
        self,
        *,
        pipeline_config: prodvana.proto.prodvana.pipelines.pipelines_pb2.PipelineConfig | None = ...,
        service_push_params: collections.abc.Iterable[global___ServicePushParam] | None = ...,
        skip_rollback_validation: builtins.bool = ...,
        service_instance_pushes: collections.abc.Iterable[global___ServiceInstancePush] | None = ...,
        skip_alert_checks: builtins.bool = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["pipeline_config", b"pipeline_config"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["pipeline_config", b"pipeline_config", "service_instance_pushes", b"service_instance_pushes", "service_push_params", b"service_push_params", "skip_alert_checks", b"skip_alert_checks", "skip_rollback_validation", b"skip_rollback_validation"]) -> None: ...

global___PipelineRunConfig = PipelineRunConfig

class PipelineRunState(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    CREATION_TIMESTAMP_FIELD_NUMBER: builtins.int
    LAST_UPDATE_TIMESTAMP_FIELD_NUMBER: builtins.int
    STATE_FIELD_NUMBER: builtins.int
    TERMINAL_FIELD_NUMBER: builtins.int
    TASK_STATUSES_FIELD_NUMBER: builtins.int
    LASTACTOR_FIELD_NUMBER: builtins.int
    INITIATOR_FIELD_NUMBER: builtins.int
    REASON_FIELD_NUMBER: builtins.int
    MESSAGE_FIELD_NUMBER: builtins.int
    EXTERNAL_LINK_FIELD_NUMBER: builtins.int
    @property
    def creation_timestamp(self) -> google.protobuf.timestamp_pb2.Timestamp: ...
    @property
    def last_update_timestamp(self) -> google.protobuf.timestamp_pb2.Timestamp: ...
    state: builtins.str
    terminal: builtins.bool
    """if pipeline run is in a terminal state"""
    @property
    def task_statuses(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[global___TaskStatus]:
        """This will only be set when querying specifically about this pipeline run"""
    lastActor: builtins.str
    initiator: builtins.str
    reason: builtins.str
    message: builtins.str
    external_link: builtins.str
    def __init__(
        self,
        *,
        creation_timestamp: google.protobuf.timestamp_pb2.Timestamp | None = ...,
        last_update_timestamp: google.protobuf.timestamp_pb2.Timestamp | None = ...,
        state: builtins.str = ...,
        terminal: builtins.bool = ...,
        task_statuses: collections.abc.Iterable[global___TaskStatus] | None = ...,
        lastActor: builtins.str = ...,
        initiator: builtins.str = ...,
        reason: builtins.str = ...,
        message: builtins.str = ...,
        external_link: builtins.str = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["creation_timestamp", b"creation_timestamp", "last_update_timestamp", b"last_update_timestamp"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["creation_timestamp", b"creation_timestamp", "external_link", b"external_link", "initiator", b"initiator", "lastActor", b"lastActor", "last_update_timestamp", b"last_update_timestamp", "message", b"message", "reason", b"reason", "state", b"state", "task_statuses", b"task_statuses", "terminal", b"terminal"]) -> None: ...

global___PipelineRunState = PipelineRunState

class PipelineRun(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    META_FIELD_NUMBER: builtins.int
    CONFIG_FIELD_NUMBER: builtins.int
    STATE_FIELD_NUMBER: builtins.int
    @property
    def meta(self) -> prodvana.proto.prodvana.common_config.meta_pb2.PipelineRunObjectMeta: ...
    @property
    def config(self) -> global___PipelineRunConfig: ...
    @property
    def state(self) -> global___PipelineRunState: ...
    def __init__(
        self,
        *,
        meta: prodvana.proto.prodvana.common_config.meta_pb2.PipelineRunObjectMeta | None = ...,
        config: global___PipelineRunConfig | None = ...,
        state: global___PipelineRunState | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["config", b"config", "meta", b"meta", "state", b"state"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["config", b"config", "meta", b"meta", "state", b"state"]) -> None: ...

global___PipelineRun = PipelineRun