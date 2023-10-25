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
import prodvana.proto.prodvana.async_task.task_metadata_pb2
import prodvana.proto.prodvana.desired_state.model.desired_state_pb2
import prodvana.proto.prodvana.desired_state.model.entity_pb2
import prodvana.proto.prodvana.service.service_config_pb2
import prodvana.proto.prodvana.version.source_metadata_pb2
import sys

if sys.version_info >= (3, 8):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

class SetDesiredStateReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    DESIRED_STATE_FIELD_NUMBER: builtins.int
    ROLLBACK_FIELD_NUMBER: builtins.int
    SOURCE_FIELD_NUMBER: builtins.int
    SOURCE_METADATA_FIELD_NUMBER: builtins.int
    FORCE_ASYNC_SET_DESIRED_STATE_FIELD_NUMBER: builtins.int
    @property
    def desired_state(self) -> prodvana.proto.prodvana.desired_state.model.desired_state_pb2.State: ...
    rollback: builtins.bool
    """set if this is a rollback, which will generate a desired state with faster preconditions"""
    source: prodvana.proto.prodvana.version.source_metadata_pb2.Source.ValueType
    @property
    def source_metadata(self) -> prodvana.proto.prodvana.version.source_metadata_pb2.SourceMetadata: ...
    force_async_set_desired_state: builtins.bool
    """internal use only"""
    def __init__(
        self,
        *,
        desired_state: prodvana.proto.prodvana.desired_state.model.desired_state_pb2.State | None = ...,
        rollback: builtins.bool = ...,
        source: prodvana.proto.prodvana.version.source_metadata_pb2.Source.ValueType = ...,
        source_metadata: prodvana.proto.prodvana.version.source_metadata_pb2.SourceMetadata | None = ...,
        force_async_set_desired_state: builtins.bool = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["desired_state", b"desired_state", "source_metadata", b"source_metadata"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["desired_state", b"desired_state", "force_async_set_desired_state", b"force_async_set_desired_state", "rollback", b"rollback", "source", b"source", "source_metadata", b"source_metadata"]) -> None: ...

global___SetDesiredStateReq = SetDesiredStateReq

class ValidateDesiredStateReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    DESIRED_STATE_FIELD_NUMBER: builtins.int
    ROLLBACK_FIELD_NUMBER: builtins.int
    SERVICE_INSTANCE_CONFIGS_FIELD_NUMBER: builtins.int
    LABEL_EXPANSION_ONLY_FIELD_NUMBER: builtins.int
    DISREGARD_SERVICE_VERSIONS_FIELD_NUMBER: builtins.int
    @property
    def desired_state(self) -> prodvana.proto.prodvana.desired_state.model.desired_state_pb2.State: ...
    rollback: builtins.bool
    """set if this is a rollback, which will generate a desired state with faster preconditions"""
    @property
    def service_instance_configs(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[prodvana.proto.prodvana.service.service_config_pb2.CompiledServiceInstanceConfig]:
        """optional, if set, use these service configs instead of fetching from db. can be useful when doing a ValidateApplyParameters followed by a ValidateDesiredState"""
    label_expansion_only: builtins.bool
    """only do label expansion instead of the full compilation, can be useful for speedup if the only desired outcome is to look at the materialized release channels"""
    disregard_service_versions: builtins.bool
    """assume service versions in the request are invalid and do not try to fetch them."""
    def __init__(
        self,
        *,
        desired_state: prodvana.proto.prodvana.desired_state.model.desired_state_pb2.State | None = ...,
        rollback: builtins.bool = ...,
        service_instance_configs: collections.abc.Iterable[prodvana.proto.prodvana.service.service_config_pb2.CompiledServiceInstanceConfig] | None = ...,
        label_expansion_only: builtins.bool = ...,
        disregard_service_versions: builtins.bool = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["desired_state", b"desired_state"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["desired_state", b"desired_state", "disregard_service_versions", b"disregard_service_versions", "label_expansion_only", b"label_expansion_only", "rollback", b"rollback", "service_instance_configs", b"service_instance_configs"]) -> None: ...

global___ValidateDesiredStateReq = ValidateDesiredStateReq

class SetDesiredStateResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    DESIRED_STATE_ID_FIELD_NUMBER: builtins.int
    desired_state_id: builtins.str
    """unique identifier for the desired state that was just set"""
    def __init__(
        self,
        *,
        desired_state_id: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["desired_state_id", b"desired_state_id"]) -> None: ...

global___SetDesiredStateResp = SetDesiredStateResp

class PreviewEntityGraphResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    ENTITY_GRAPH_FIELD_NUMBER: builtins.int
    DESIRED_STATE_ID_FIELD_NUMBER: builtins.int
    @property
    def entity_graph(self) -> prodvana.proto.prodvana.desired_state.model.entity_pb2.EntityGraph: ...
    desired_state_id: builtins.str
    """unique identifier the preview entity graph that was just created. This same ID can be used across any endpoints that inspect desired states, e.g. GetDesiredState"""
    def __init__(
        self,
        *,
        entity_graph: prodvana.proto.prodvana.desired_state.model.entity_pb2.EntityGraph | None = ...,
        desired_state_id: builtins.str = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["entity_graph", b"entity_graph"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["desired_state_id", b"desired_state_id", "entity_graph", b"entity_graph"]) -> None: ...

global___PreviewEntityGraphResp = PreviewEntityGraphResp

class GetServiceDesiredStateConvergenceSummaryReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    APPLICATION_FIELD_NUMBER: builtins.int
    SERVICE_FIELD_NUMBER: builtins.int
    FAST_NO_DEPRECATED_FIELDS_FIELD_NUMBER: builtins.int
    application: builtins.str
    service: builtins.str
    fast_no_deprecated_fields: builtins.bool
    """unused"""
    def __init__(
        self,
        *,
        application: builtins.str = ...,
        service: builtins.str = ...,
        fast_no_deprecated_fields: builtins.bool = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["application", b"application", "fast_no_deprecated_fields", b"fast_no_deprecated_fields", "service", b"service"]) -> None: ...

global___GetServiceDesiredStateConvergenceSummaryReq = GetServiceDesiredStateConvergenceSummaryReq

class StatusExplanations(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    STATUS_EXPLANATIONS_FIELD_NUMBER: builtins.int
    @property
    def status_explanations(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[prodvana.proto.prodvana.desired_state.model.desired_state_pb2.StatusExplanation]: ...
    def __init__(
        self,
        *,
        status_explanations: collections.abc.Iterable[prodvana.proto.prodvana.desired_state.model.desired_state_pb2.StatusExplanation] | None = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["status_explanations", b"status_explanations"]) -> None: ...

global___StatusExplanations = StatusExplanations

class DebugLogs(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    DEBUG_LOGS_FIELD_NUMBER: builtins.int
    @property
    def debug_logs(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[prodvana.proto.prodvana.desired_state.model.desired_state_pb2.DebugLog]: ...
    def __init__(
        self,
        *,
        debug_logs: collections.abc.Iterable[prodvana.proto.prodvana.desired_state.model.desired_state_pb2.DebugLog] | None = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["debug_logs", b"debug_logs"]) -> None: ...

global___DebugLogs = DebugLogs

class PendingSetDesiredState(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    DESIRED_STATE_ID_FIELD_NUMBER: builtins.int
    COMPILED_DESIRED_STATE_FIELD_NUMBER: builtins.int
    TASK_STATUS_FIELD_NUMBER: builtins.int
    TASK_RESULT_FIELD_NUMBER: builtins.int
    desired_state_id: builtins.str
    @property
    def compiled_desired_state(self) -> prodvana.proto.prodvana.desired_state.model.desired_state_pb2.State: ...
    task_status: prodvana.proto.prodvana.async_task.task_metadata_pb2.TaskStatus.ValueType
    """will never contain SUCCESS, by definition, but may contain FAILED"""
    @property
    def task_result(self) -> prodvana.proto.prodvana.async_task.task_metadata_pb2.TaskResult:
        """will only be set for FAILED
        next tag: 5
        """
    def __init__(
        self,
        *,
        desired_state_id: builtins.str = ...,
        compiled_desired_state: prodvana.proto.prodvana.desired_state.model.desired_state_pb2.State | None = ...,
        task_status: prodvana.proto.prodvana.async_task.task_metadata_pb2.TaskStatus.ValueType = ...,
        task_result: prodvana.proto.prodvana.async_task.task_metadata_pb2.TaskResult | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["compiled_desired_state", b"compiled_desired_state", "task_result", b"task_result"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["compiled_desired_state", b"compiled_desired_state", "desired_state_id", b"desired_state_id", "task_result", b"task_result", "task_status", b"task_status"]) -> None: ...

global___PendingSetDesiredState = PendingSetDesiredState

class DesiredStateSummary(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    class StatusesEntry(google.protobuf.message.Message):
        DESCRIPTOR: google.protobuf.descriptor.Descriptor

        KEY_FIELD_NUMBER: builtins.int
        VALUE_FIELD_NUMBER: builtins.int
        key: builtins.str
        value: prodvana.proto.prodvana.desired_state.model.desired_state_pb2.Status.ValueType
        def __init__(
            self,
            *,
            key: builtins.str = ...,
            value: prodvana.proto.prodvana.desired_state.model.desired_state_pb2.Status.ValueType = ...,
        ) -> None: ...
        def ClearField(self, field_name: typing_extensions.Literal["key", b"key", "value", b"value"]) -> None: ...

    class PreconditionStatusesEntry(google.protobuf.message.Message):
        DESCRIPTOR: google.protobuf.descriptor.Descriptor

        KEY_FIELD_NUMBER: builtins.int
        VALUE_FIELD_NUMBER: builtins.int
        key: builtins.str
        @property
        def value(self) -> prodvana.proto.prodvana.desired_state.model.desired_state_pb2.ConditionState: ...
        def __init__(
            self,
            *,
            key: builtins.str = ...,
            value: prodvana.proto.prodvana.desired_state.model.desired_state_pb2.ConditionState | None = ...,
        ) -> None: ...
        def HasField(self, field_name: typing_extensions.Literal["value", b"value"]) -> builtins.bool: ...
        def ClearField(self, field_name: typing_extensions.Literal["key", b"key", "value", b"value"]) -> None: ...

    class StatusExplanationsEntry(google.protobuf.message.Message):
        DESCRIPTOR: google.protobuf.descriptor.Descriptor

        KEY_FIELD_NUMBER: builtins.int
        VALUE_FIELD_NUMBER: builtins.int
        key: builtins.str
        @property
        def value(self) -> global___StatusExplanations: ...
        def __init__(
            self,
            *,
            key: builtins.str = ...,
            value: global___StatusExplanations | None = ...,
        ) -> None: ...
        def HasField(self, field_name: typing_extensions.Literal["value", b"value"]) -> builtins.bool: ...
        def ClearField(self, field_name: typing_extensions.Literal["key", b"key", "value", b"value"]) -> None: ...

    class DebugLogsEntry(google.protobuf.message.Message):
        DESCRIPTOR: google.protobuf.descriptor.Descriptor

        KEY_FIELD_NUMBER: builtins.int
        VALUE_FIELD_NUMBER: builtins.int
        key: builtins.str
        @property
        def value(self) -> global___DebugLogs: ...
        def __init__(
            self,
            *,
            key: builtins.str = ...,
            value: global___DebugLogs | None = ...,
        ) -> None: ...
        def HasField(self, field_name: typing_extensions.Literal["value", b"value"]) -> builtins.bool: ...
        def ClearField(self, field_name: typing_extensions.Literal["key", b"key", "value", b"value"]) -> None: ...

    class ActionExplanationsEntry(google.protobuf.message.Message):
        DESCRIPTOR: google.protobuf.descriptor.Descriptor

        KEY_FIELD_NUMBER: builtins.int
        VALUE_FIELD_NUMBER: builtins.int
        key: builtins.str
        @property
        def value(self) -> prodvana.proto.prodvana.desired_state.model.desired_state_pb2.ActionExplanation: ...
        def __init__(
            self,
            *,
            key: builtins.str = ...,
            value: prodvana.proto.prodvana.desired_state.model.desired_state_pb2.ActionExplanation | None = ...,
        ) -> None: ...
        def HasField(self, field_name: typing_extensions.Literal["value", b"value"]) -> builtins.bool: ...
        def ClearField(self, field_name: typing_extensions.Literal["key", b"key", "value", b"value"]) -> None: ...

    class LastUpdateTimestampsEntry(google.protobuf.message.Message):
        DESCRIPTOR: google.protobuf.descriptor.Descriptor

        KEY_FIELD_NUMBER: builtins.int
        VALUE_FIELD_NUMBER: builtins.int
        key: builtins.str
        @property
        def value(self) -> google.protobuf.timestamp_pb2.Timestamp: ...
        def __init__(
            self,
            *,
            key: builtins.str = ...,
            value: google.protobuf.timestamp_pb2.Timestamp | None = ...,
        ) -> None: ...
        def HasField(self, field_name: typing_extensions.Literal["value", b"value"]) -> builtins.bool: ...
        def ClearField(self, field_name: typing_extensions.Literal["key", b"key", "value", b"value"]) -> None: ...

    class LastFetchedTimestampsEntry(google.protobuf.message.Message):
        DESCRIPTOR: google.protobuf.descriptor.Descriptor

        KEY_FIELD_NUMBER: builtins.int
        VALUE_FIELD_NUMBER: builtins.int
        key: builtins.str
        @property
        def value(self) -> google.protobuf.timestamp_pb2.Timestamp: ...
        def __init__(
            self,
            *,
            key: builtins.str = ...,
            value: google.protobuf.timestamp_pb2.Timestamp | None = ...,
        ) -> None: ...
        def HasField(self, field_name: typing_extensions.Literal["value", b"value"]) -> builtins.bool: ...
        def ClearField(self, field_name: typing_extensions.Literal["key", b"key", "value", b"value"]) -> None: ...

    class LastAppliedTimestampsEntry(google.protobuf.message.Message):
        DESCRIPTOR: google.protobuf.descriptor.Descriptor

        KEY_FIELD_NUMBER: builtins.int
        VALUE_FIELD_NUMBER: builtins.int
        key: builtins.str
        @property
        def value(self) -> google.protobuf.timestamp_pb2.Timestamp: ...
        def __init__(
            self,
            *,
            key: builtins.str = ...,
            value: google.protobuf.timestamp_pb2.Timestamp | None = ...,
        ) -> None: ...
        def HasField(self, field_name: typing_extensions.Literal["value", b"value"]) -> builtins.bool: ...
        def ClearField(self, field_name: typing_extensions.Literal["key", b"key", "value", b"value"]) -> None: ...

    ENTITY_GRAPH_FIELD_NUMBER: builtins.int
    CREATION_TIMESTAMP_FIELD_NUMBER: builtins.int
    LAST_UPDATE_TIMESTAMP_FIELD_NUMBER: builtins.int
    REPLACED_TIMESTAMP_FIELD_NUMBER: builtins.int
    INPUT_DESIRED_STATE_FIELD_NUMBER: builtins.int
    DESIRED_STATE_FIELD_NUMBER: builtins.int
    STARTING_STATE_FIELD_NUMBER: builtins.int
    LAST_SEEN_STATE_FIELD_NUMBER: builtins.int
    STATUS_FIELD_NUMBER: builtins.int
    SOURCE_FIELD_NUMBER: builtins.int
    SOURCE_METADATA_FIELD_NUMBER: builtins.int
    PENDING_SET_DESIRED_STATE_FIELD_NUMBER: builtins.int
    STATUSES_FIELD_NUMBER: builtins.int
    PRECONDITION_STATUSES_FIELD_NUMBER: builtins.int
    STATUS_EXPLANATIONS_FIELD_NUMBER: builtins.int
    DEBUG_LOGS_FIELD_NUMBER: builtins.int
    ACTION_EXPLANATIONS_FIELD_NUMBER: builtins.int
    LAST_UPDATE_TIMESTAMPS_FIELD_NUMBER: builtins.int
    LAST_FETCHED_TIMESTAMPS_FIELD_NUMBER: builtins.int
    LAST_APPLIED_TIMESTAMPS_FIELD_NUMBER: builtins.int
    @property
    def entity_graph(self) -> prodvana.proto.prodvana.desired_state.model.entity_pb2.EntityGraph: ...
    @property
    def creation_timestamp(self) -> google.protobuf.timestamp_pb2.Timestamp: ...
    @property
    def last_update_timestamp(self) -> google.protobuf.timestamp_pb2.Timestamp: ...
    @property
    def replaced_timestamp(self) -> google.protobuf.timestamp_pb2.Timestamp:
        """will only be set if desired state has been replaced"""
    @property
    def input_desired_state(self) -> prodvana.proto.prodvana.desired_state.model.desired_state_pb2.State:
        """uncompiled desired state originally passed as input to SetDesiredState"""
    @property
    def desired_state(self) -> prodvana.proto.prodvana.desired_state.model.desired_state_pb2.State:
        """compiled desired state"""
    @property
    def starting_state(self) -> prodvana.proto.prodvana.desired_state.model.desired_state_pb2.State: ...
    @property
    def last_seen_state(self) -> prodvana.proto.prodvana.desired_state.model.desired_state_pb2.State: ...
    status: prodvana.proto.prodvana.desired_state.model.desired_state_pb2.Status.ValueType
    source: prodvana.proto.prodvana.version.source_metadata_pb2.Source.ValueType
    """who set the desired state"""
    @property
    def source_metadata(self) -> prodvana.proto.prodvana.version.source_metadata_pb2.SourceMetadata: ...
    @property
    def pending_set_desired_state(self) -> global___PendingSetDesiredState:
        """the latest pending or failed set desired state request"""
    @property
    def statuses(self) -> google.protobuf.internal.containers.ScalarMap[builtins.str, prodvana.proto.prodvana.desired_state.model.desired_state_pb2.Status.ValueType]:
        """fields below are deprecated
        status of the desired state, all its descendents, and the root desired state.
        """
    @property
    def precondition_statuses(self) -> google.protobuf.internal.containers.MessageMap[builtins.str, prodvana.proto.prodvana.desired_state.model.desired_state_pb2.ConditionState]:
        """State of all preconditions involved in the desired state.
        Key format - "{desired_state_id}-{index}", e.g., "des-0752d6c76a6943abbfcf0515634584ba-0".
        """
    @property
    def status_explanations(self) -> google.protobuf.internal.containers.MessageMap[builtins.str, global___StatusExplanations]: ...
    @property
    def debug_logs(self) -> google.protobuf.internal.containers.MessageMap[builtins.str, global___DebugLogs]: ...
    @property
    def action_explanations(self) -> google.protobuf.internal.containers.MessageMap[builtins.str, prodvana.proto.prodvana.desired_state.model.desired_state_pb2.ActionExplanation]: ...
    @property
    def last_update_timestamps(self) -> google.protobuf.internal.containers.MessageMap[builtins.str, google.protobuf.timestamp_pb2.Timestamp]:
        """last_update_timestamp, last_fetched_timestamp, and last_applied_timestamp are only returned for the latest desired state"""
    @property
    def last_fetched_timestamps(self) -> google.protobuf.internal.containers.MessageMap[builtins.str, google.protobuf.timestamp_pb2.Timestamp]: ...
    @property
    def last_applied_timestamps(self) -> google.protobuf.internal.containers.MessageMap[builtins.str, google.protobuf.timestamp_pb2.Timestamp]: ...
    def __init__(
        self,
        *,
        entity_graph: prodvana.proto.prodvana.desired_state.model.entity_pb2.EntityGraph | None = ...,
        creation_timestamp: google.protobuf.timestamp_pb2.Timestamp | None = ...,
        last_update_timestamp: google.protobuf.timestamp_pb2.Timestamp | None = ...,
        replaced_timestamp: google.protobuf.timestamp_pb2.Timestamp | None = ...,
        input_desired_state: prodvana.proto.prodvana.desired_state.model.desired_state_pb2.State | None = ...,
        desired_state: prodvana.proto.prodvana.desired_state.model.desired_state_pb2.State | None = ...,
        starting_state: prodvana.proto.prodvana.desired_state.model.desired_state_pb2.State | None = ...,
        last_seen_state: prodvana.proto.prodvana.desired_state.model.desired_state_pb2.State | None = ...,
        status: prodvana.proto.prodvana.desired_state.model.desired_state_pb2.Status.ValueType = ...,
        source: prodvana.proto.prodvana.version.source_metadata_pb2.Source.ValueType = ...,
        source_metadata: prodvana.proto.prodvana.version.source_metadata_pb2.SourceMetadata | None = ...,
        pending_set_desired_state: global___PendingSetDesiredState | None = ...,
        statuses: collections.abc.Mapping[builtins.str, prodvana.proto.prodvana.desired_state.model.desired_state_pb2.Status.ValueType] | None = ...,
        precondition_statuses: collections.abc.Mapping[builtins.str, prodvana.proto.prodvana.desired_state.model.desired_state_pb2.ConditionState] | None = ...,
        status_explanations: collections.abc.Mapping[builtins.str, global___StatusExplanations] | None = ...,
        debug_logs: collections.abc.Mapping[builtins.str, global___DebugLogs] | None = ...,
        action_explanations: collections.abc.Mapping[builtins.str, prodvana.proto.prodvana.desired_state.model.desired_state_pb2.ActionExplanation] | None = ...,
        last_update_timestamps: collections.abc.Mapping[builtins.str, google.protobuf.timestamp_pb2.Timestamp] | None = ...,
        last_fetched_timestamps: collections.abc.Mapping[builtins.str, google.protobuf.timestamp_pb2.Timestamp] | None = ...,
        last_applied_timestamps: collections.abc.Mapping[builtins.str, google.protobuf.timestamp_pb2.Timestamp] | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["creation_timestamp", b"creation_timestamp", "desired_state", b"desired_state", "entity_graph", b"entity_graph", "input_desired_state", b"input_desired_state", "last_seen_state", b"last_seen_state", "last_update_timestamp", b"last_update_timestamp", "pending_set_desired_state", b"pending_set_desired_state", "replaced_timestamp", b"replaced_timestamp", "source_metadata", b"source_metadata", "starting_state", b"starting_state"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["action_explanations", b"action_explanations", "creation_timestamp", b"creation_timestamp", "debug_logs", b"debug_logs", "desired_state", b"desired_state", "entity_graph", b"entity_graph", "input_desired_state", b"input_desired_state", "last_applied_timestamps", b"last_applied_timestamps", "last_fetched_timestamps", b"last_fetched_timestamps", "last_seen_state", b"last_seen_state", "last_update_timestamp", b"last_update_timestamp", "last_update_timestamps", b"last_update_timestamps", "pending_set_desired_state", b"pending_set_desired_state", "precondition_statuses", b"precondition_statuses", "replaced_timestamp", b"replaced_timestamp", "source", b"source", "source_metadata", b"source_metadata", "starting_state", b"starting_state", "status", b"status", "status_explanations", b"status_explanations", "statuses", b"statuses"]) -> None: ...

global___DesiredStateSummary = DesiredStateSummary

class GetDesiredStateConvergenceReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    DESIRED_STATE_ID_FIELD_NUMBER: builtins.int
    FAST_NO_DEPRECATED_FIELDS_FIELD_NUMBER: builtins.int
    desired_state_id: builtins.str
    fast_no_deprecated_fields: builtins.bool
    """unused"""
    def __init__(
        self,
        *,
        desired_state_id: builtins.str = ...,
        fast_no_deprecated_fields: builtins.bool = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["desired_state_id", b"desired_state_id", "fast_no_deprecated_fields", b"fast_no_deprecated_fields"]) -> None: ...

global___GetDesiredStateConvergenceReq = GetDesiredStateConvergenceReq

class GetDesiredStateConvergenceSummaryResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    SUMMARY_FIELD_NUMBER: builtins.int
    @property
    def summary(self) -> global___DesiredStateSummary: ...
    def __init__(
        self,
        *,
        summary: global___DesiredStateSummary | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["summary", b"summary"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["summary", b"summary"]) -> None: ...

global___GetDesiredStateConvergenceSummaryResp = GetDesiredStateConvergenceSummaryResp

class GetServiceDesiredStateConvergenceSummaryResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    SUMMARY_FIELD_NUMBER: builtins.int
    @property
    def summary(self) -> global___DesiredStateSummary: ...
    def __init__(
        self,
        *,
        summary: global___DesiredStateSummary | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["summary", b"summary"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["summary", b"summary"]) -> None: ...

global___GetServiceDesiredStateConvergenceSummaryResp = GetServiceDesiredStateConvergenceSummaryResp

class GetServiceLastConvergedStateReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    APPLICATION_FIELD_NUMBER: builtins.int
    SERVICE_FIELD_NUMBER: builtins.int
    application: builtins.str
    service: builtins.str
    def __init__(
        self,
        *,
        application: builtins.str = ...,
        service: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["application", b"application", "service", b"service"]) -> None: ...

global___GetServiceLastConvergedStateReq = GetServiceLastConvergedStateReq

class GetServiceLastConvergedStateResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    SERVICE_INSTANCE_STATES_FIELD_NUMBER: builtins.int
    @property
    def service_instance_states(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[prodvana.proto.prodvana.desired_state.model.desired_state_pb2.ServiceInstanceState]: ...
    def __init__(
        self,
        *,
        service_instance_states: collections.abc.Iterable[prodvana.proto.prodvana.desired_state.model.desired_state_pb2.ServiceInstanceState] | None = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["service_instance_states", b"service_instance_states"]) -> None: ...

global___GetServiceLastConvergedStateResp = GetServiceLastConvergedStateResp

class GetServiceDesiredStateHistoryReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    APPLICATION_FIELD_NUMBER: builtins.int
    SERVICE_FIELD_NUMBER: builtins.int
    PAGE_TOKEN_FIELD_NUMBER: builtins.int
    PAGE_SIZE_FIELD_NUMBER: builtins.int
    FAST_NO_DEPRECATED_FIELDS_FIELD_NUMBER: builtins.int
    application: builtins.str
    service: builtins.str
    page_token: builtins.str
    page_size: builtins.int
    fast_no_deprecated_fields: builtins.bool
    """unused"""
    def __init__(
        self,
        *,
        application: builtins.str = ...,
        service: builtins.str = ...,
        page_token: builtins.str = ...,
        page_size: builtins.int = ...,
        fast_no_deprecated_fields: builtins.bool = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["application", b"application", "fast_no_deprecated_fields", b"fast_no_deprecated_fields", "page_size", b"page_size", "page_token", b"page_token", "service", b"service"]) -> None: ...

global___GetServiceDesiredStateHistoryReq = GetServiceDesiredStateHistoryReq

class GetServiceDesiredStateHistoryResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    DESIRED_STATES_FIELD_NUMBER: builtins.int
    NEXT_PAGE_TOKEN_FIELD_NUMBER: builtins.int
    @property
    def desired_states(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[global___DesiredStateSummary]: ...
    next_page_token: builtins.str
    def __init__(
        self,
        *,
        desired_states: collections.abc.Iterable[global___DesiredStateSummary] | None = ...,
        next_page_token: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["desired_states", b"desired_states", "next_page_token", b"next_page_token"]) -> None: ...

global___GetServiceDesiredStateHistoryResp = GetServiceDesiredStateHistoryResp

class GetDesiredStateReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    DESIRED_STATE_ID_FIELD_NUMBER: builtins.int
    desired_state_id: builtins.str
    def __init__(
        self,
        *,
        desired_state_id: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["desired_state_id", b"desired_state_id"]) -> None: ...

global___GetDesiredStateReq = GetDesiredStateReq

class GetDesiredStateResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    DESIRED_STATE_FIELD_NUMBER: builtins.int
    COMPILED_DESIRED_STATE_FIELD_NUMBER: builtins.int
    @property
    def desired_state(self) -> prodvana.proto.prodvana.desired_state.model.desired_state_pb2.State: ...
    @property
    def compiled_desired_state(self) -> prodvana.proto.prodvana.desired_state.model.desired_state_pb2.State: ...
    def __init__(
        self,
        *,
        desired_state: prodvana.proto.prodvana.desired_state.model.desired_state_pb2.State | None = ...,
        compiled_desired_state: prodvana.proto.prodvana.desired_state.model.desired_state_pb2.State | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["compiled_desired_state", b"compiled_desired_state", "desired_state", b"desired_state"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["compiled_desired_state", b"compiled_desired_state", "desired_state", b"desired_state"]) -> None: ...

global___GetDesiredStateResp = GetDesiredStateResp

class GetDesiredStateGraphReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    DESIRED_STATE_ID_FIELD_NUMBER: builtins.int
    TYPE_FIELD_NUMBER: builtins.int
    REQUIRED_ANNOTATIONS_FIELD_NUMBER: builtins.int
    DEPTH_FIELD_NUMBER: builtins.int
    desired_state_id: builtins.str
    """Root desired state id to get the graph for."""
    type: prodvana.proto.prodvana.desired_state.model.desired_state_pb2.Type.ValueType
    """Find interesting entities in the graph.
    Which entities are interesting is defined by the type and required_annotations fields.

    Which type of entities to find. Set to UNKNOWN to find all entities.
    """
    @property
    def required_annotations(self) -> prodvana.proto.prodvana.desired_state.model.desired_state_pb2.Annotations:
        """Which annotations are required for an entity to be considered interesting."""
    depth: builtins.int
    """For all interesting entities, also include all children up to the given depth. 0 means no children."""
    def __init__(
        self,
        *,
        desired_state_id: builtins.str = ...,
        type: prodvana.proto.prodvana.desired_state.model.desired_state_pb2.Type.ValueType = ...,
        required_annotations: prodvana.proto.prodvana.desired_state.model.desired_state_pb2.Annotations | None = ...,
        depth: builtins.int = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["required_annotations", b"required_annotations"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["depth", b"depth", "desired_state_id", b"desired_state_id", "required_annotations", b"required_annotations", "type", b"type"]) -> None: ...

global___GetDesiredStateGraphReq = GetDesiredStateGraphReq

class GetDesiredStateGraphResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    ENTITY_GRAPH_FIELD_NUMBER: builtins.int
    PENDING_SET_DESIRED_STATE_FIELD_NUMBER: builtins.int
    @property
    def entity_graph(self) -> prodvana.proto.prodvana.desired_state.model.entity_pb2.EntityGraph: ...
    @property
    def pending_set_desired_state(self) -> global___PendingSetDesiredState: ...
    def __init__(
        self,
        *,
        entity_graph: prodvana.proto.prodvana.desired_state.model.entity_pb2.EntityGraph | None = ...,
        pending_set_desired_state: global___PendingSetDesiredState | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["entity_graph", b"entity_graph", "pending_set_desired_state", b"pending_set_desired_state", "resp", b"resp"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["entity_graph", b"entity_graph", "pending_set_desired_state", b"pending_set_desired_state", "resp", b"resp"]) -> None: ...
    def WhichOneof(self, oneof_group: typing_extensions.Literal["resp", b"resp"]) -> typing_extensions.Literal["entity_graph", "pending_set_desired_state"] | None: ...

global___GetDesiredStateGraphResp = GetDesiredStateGraphResp

class ValidateDesiredStateResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    DESIRED_STATE_FIELD_NUMBER: builtins.int
    COMPILED_DESIRED_STATE_FIELD_NUMBER: builtins.int
    @property
    def desired_state(self) -> prodvana.proto.prodvana.desired_state.model.desired_state_pb2.State: ...
    @property
    def compiled_desired_state(self) -> prodvana.proto.prodvana.desired_state.model.desired_state_pb2.State: ...
    def __init__(
        self,
        *,
        desired_state: prodvana.proto.prodvana.desired_state.model.desired_state_pb2.State | None = ...,
        compiled_desired_state: prodvana.proto.prodvana.desired_state.model.desired_state_pb2.State | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["compiled_desired_state", b"compiled_desired_state", "desired_state", b"desired_state"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["compiled_desired_state", b"compiled_desired_state", "desired_state", b"desired_state"]) -> None: ...

global___ValidateDesiredStateResp = ValidateDesiredStateResp

class SetManualApprovalReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    DESIRED_STATE_ID_FIELD_NUMBER: builtins.int
    TOPIC_FIELD_NUMBER: builtins.int
    REJECT_FIELD_NUMBER: builtins.int
    SIGNAL_TYPE_FIELD_NUMBER: builtins.int
    desired_state_id: builtins.str
    topic: builtins.str
    """string application = 2 [(prodvana.proto.validate.rules).string.min_len = 1];
    string service = 3 [(prodvana.proto.validate.rules).string.min_len = 1];
    string release_channel = 4 [(prodvana.proto.validate.rules).string.min_len = 1];
    """
    reject: builtins.bool
    signal_type: builtins.str
    def __init__(
        self,
        *,
        desired_state_id: builtins.str = ...,
        topic: builtins.str = ...,
        reject: builtins.bool = ...,
        signal_type: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["desired_state_id", b"desired_state_id", "reject", b"reject", "signal_type", b"signal_type", "topic", b"topic"]) -> None: ...

global___SetManualApprovalReq = SetManualApprovalReq

class SetManualApprovalResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    def __init__(
        self,
    ) -> None: ...

global___SetManualApprovalResp = SetManualApprovalResp

class PromoteDeliveryReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    DESIRED_STATE_ID_FIELD_NUMBER: builtins.int
    STAGE_FIELD_NUMBER: builtins.int
    FULL_FIELD_NUMBER: builtins.int
    SOURCE_FIELD_NUMBER: builtins.int
    desired_state_id: builtins.str
    stage: builtins.int
    full: builtins.bool
    source: builtins.str
    def __init__(
        self,
        *,
        desired_state_id: builtins.str = ...,
        stage: builtins.int = ...,
        full: builtins.bool = ...,
        source: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["desired_state_id", b"desired_state_id", "full", b"full", "source", b"source", "stage", b"stage"]) -> None: ...

global___PromoteDeliveryReq = PromoteDeliveryReq

class PromoteDeliveryResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    def __init__(
        self,
    ) -> None: ...

global___PromoteDeliveryResp = PromoteDeliveryResp

class BypassProtectionReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    DESIRED_STATE_ID_FIELD_NUMBER: builtins.int
    SOURCE_FIELD_NUMBER: builtins.int
    desired_state_id: builtins.str
    source: builtins.str
    def __init__(
        self,
        *,
        desired_state_id: builtins.str = ...,
        source: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["desired_state_id", b"desired_state_id", "source", b"source"]) -> None: ...

global___BypassProtectionReq = BypassProtectionReq

class BypassProtectionResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    def __init__(
        self,
    ) -> None: ...

global___BypassProtectionResp = BypassProtectionResp
