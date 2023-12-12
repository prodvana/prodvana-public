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
import google.protobuf.timestamp_pb2
import prodvana.proto.prodvana.object.meta_pb2
import prodvana.proto.prodvana.repo.repo_pb2
import sys
import typing

if sys.version_info >= (3, 10):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

class _DeploymentStatus:
    ValueType = typing.NewType("ValueType", builtins.int)
    V: typing_extensions.TypeAlias = ValueType

class _DeploymentStatusEnumTypeWrapper(google.protobuf.internal.enum_type_wrapper._EnumTypeWrapper[_DeploymentStatus.ValueType], builtins.type):  # noqa: F821
    DESCRIPTOR: google.protobuf.descriptor.EnumDescriptor
    UNKNOWN_STATUS: _DeploymentStatus.ValueType  # 0
    PENDING: _DeploymentStatus.ValueType  # 1
    SUCCEEDED: _DeploymentStatus.ValueType  # 2
    FAILED: _DeploymentStatus.ValueType  # 3
    PREVIEW: _DeploymentStatus.ValueType  # 4

class DeploymentStatus(_DeploymentStatus, metaclass=_DeploymentStatusEnumTypeWrapper): ...

UNKNOWN_STATUS: DeploymentStatus.ValueType  # 0
PENDING: DeploymentStatus.ValueType  # 1
SUCCEEDED: DeploymentStatus.ValueType  # 2
FAILED: DeploymentStatus.ValueType  # 3
PREVIEW: DeploymentStatus.ValueType  # 4
global___DeploymentStatus = DeploymentStatus

class DeploymentConfig(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    CREATION_TIMESTAMP_FIELD_NUMBER: builtins.int
    DEPLOYMENT_SYSTEM_FIELD_NUMBER: builtins.int
    SERVICE_FIELD_NUMBER: builtins.int
    RELEASE_CHANNEL_FIELD_NUMBER: builtins.int
    APPLICATION_FIELD_NUMBER: builtins.int
    REPOSITORY_FIELD_NUMBER: builtins.int
    COMMIT_ID_FIELD_NUMBER: builtins.int
    USER_FIELD_NUMBER: builtins.int
    APPLICATION_ID_FIELD_NUMBER: builtins.int
    SERVICE_ID_FIELD_NUMBER: builtins.int
    RELEASE_CHANNEL_ID_FIELD_NUMBER: builtins.int
    SERVICE_VERSION_FIELD_NUMBER: builtins.int
    DESIRED_STATE_ID_FIELD_NUMBER: builtins.int
    @property
    def creation_timestamp(self) -> google.protobuf.timestamp_pb2.Timestamp:
        """must be unset on input"""
    deployment_system: builtins.str
    service: builtins.str
    """required when recording deployments from external systems"""
    release_channel: builtins.str
    """required when recording deployments from external systems"""
    application: builtins.str
    """required when recording deployments from external systems"""
    repository: builtins.str
    """e.g. github.com/foo/bar"""
    commit_id: builtins.str
    """commit hash"""
    user: builtins.str
    """if known"""
    application_id: builtins.str
    """The following fields only make sense if managed by Prodvana."""
    service_id: builtins.str
    release_channel_id: builtins.str
    service_version: builtins.str
    desired_state_id: builtins.str
    """next tag: 14"""
    def __init__(
        self,
        *,
        creation_timestamp: google.protobuf.timestamp_pb2.Timestamp | None = ...,
        deployment_system: builtins.str = ...,
        service: builtins.str = ...,
        release_channel: builtins.str = ...,
        application: builtins.str = ...,
        repository: builtins.str = ...,
        commit_id: builtins.str = ...,
        user: builtins.str = ...,
        application_id: builtins.str = ...,
        service_id: builtins.str = ...,
        release_channel_id: builtins.str = ...,
        service_version: builtins.str = ...,
        desired_state_id: builtins.str = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["creation_timestamp", b"creation_timestamp"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["application", b"application", "application_id", b"application_id", "commit_id", b"commit_id", "creation_timestamp", b"creation_timestamp", "deployment_system", b"deployment_system", "desired_state_id", b"desired_state_id", "release_channel", b"release_channel", "release_channel_id", b"release_channel_id", "repository", b"repository", "service", b"service", "service_id", b"service_id", "service_version", b"service_version", "user", b"user"]) -> None: ...

global___DeploymentConfig = DeploymentConfig

class DeploymentState(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    STATUS_FIELD_NUMBER: builtins.int
    LAST_UPDATE_TIMESTAMP_FIELD_NUMBER: builtins.int
    status: global___DeploymentStatus.ValueType
    @property
    def last_update_timestamp(self) -> google.protobuf.timestamp_pb2.Timestamp: ...
    def __init__(
        self,
        *,
        status: global___DeploymentStatus.ValueType = ...,
        last_update_timestamp: google.protobuf.timestamp_pb2.Timestamp | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["last_update_timestamp", b"last_update_timestamp"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["last_update_timestamp", b"last_update_timestamp", "status", b"status"]) -> None: ...

global___DeploymentState = DeploymentState

class CommitAnalysis(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    COMMITS_ADDED_FIELD_NUMBER: builtins.int
    COMMITS_REMOVED_FIELD_NUMBER: builtins.int
    IMPACT_ANALYSIS_FIELD_NUMBER: builtins.int
    commits_added: builtins.int
    commits_removed: builtins.int
    @property
    def impact_analysis(self) -> global___ImpactAnalysisComparison: ...
    def __init__(
        self,
        *,
        commits_added: builtins.int = ...,
        commits_removed: builtins.int = ...,
        impact_analysis: global___ImpactAnalysisComparison | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["impact_analysis", b"impact_analysis"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["commits_added", b"commits_added", "commits_removed", b"commits_removed", "impact_analysis", b"impact_analysis"]) -> None: ...

global___CommitAnalysis = CommitAnalysis

class ImpactAnalysisComparison(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    RELEVANT_ADDED_COMMITS_FIELD_NUMBER: builtins.int
    RELEVANT_REMOVED_COMMITS_FIELD_NUMBER: builtins.int
    UNANALYZED_COMMITS_FIELD_NUMBER: builtins.int
    @property
    def relevant_added_commits(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[prodvana.proto.prodvana.repo.repo_pb2.Commit]:
        """commits likely to be impactful, prev_commit_id and new_commit_id have a merge base"""
    @property
    def relevant_removed_commits(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[prodvana.proto.prodvana.repo.repo_pb2.Commit]: ...
    unanalyzed_commits: builtins.int
    def __init__(
        self,
        *,
        relevant_added_commits: collections.abc.Iterable[prodvana.proto.prodvana.repo.repo_pb2.Commit] | None = ...,
        relevant_removed_commits: collections.abc.Iterable[prodvana.proto.prodvana.repo.repo_pb2.Commit] | None = ...,
        unanalyzed_commits: builtins.int = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["relevant_added_commits", b"relevant_added_commits", "relevant_removed_commits", b"relevant_removed_commits", "unanalyzed_commits", b"unanalyzed_commits"]) -> None: ...

global___ImpactAnalysisComparison = ImpactAnalysisComparison

class DeploymentComparison(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    PREV_FIELD_NUMBER: builtins.int
    PREV_REPOSITORY_FIELD_NUMBER: builtins.int
    NEW_REPOSITORY_FIELD_NUMBER: builtins.int
    PREV_COMMIT_ID_FIELD_NUMBER: builtins.int
    NEW_COMMIT_ID_FIELD_NUMBER: builtins.int
    COMMIT_ANALYSIS_FIELD_NUMBER: builtins.int
    PREV_SERVICE_ID_FIELD_NUMBER: builtins.int
    PREV_RELEASE_CHANNEL_ID_FIELD_NUMBER: builtins.int
    PREV_SERVICE_VERSION_FIELD_NUMBER: builtins.int
    NEW_SERVICE_ID_FIELD_NUMBER: builtins.int
    NEW_RELEASE_CHANNEL_ID_FIELD_NUMBER: builtins.int
    NEW_SERVICE_VERSION_FIELD_NUMBER: builtins.int
    @property
    def prev(self) -> prodvana.proto.prodvana.object.meta_pb2.ObjectMeta: ...
    prev_repository: builtins.str
    new_repository: builtins.str
    prev_commit_id: builtins.str
    new_commit_id: builtins.str
    @property
    def commit_analysis(self) -> global___CommitAnalysis:
        """only set if the previous commit is set and is from the same repo, and the repo is linked to prodvana"""
    prev_service_id: builtins.str
    """only set for Prodvana managed deployments"""
    prev_release_channel_id: builtins.str
    prev_service_version: builtins.str
    new_service_id: builtins.str
    new_release_channel_id: builtins.str
    new_service_version: builtins.str
    def __init__(
        self,
        *,
        prev: prodvana.proto.prodvana.object.meta_pb2.ObjectMeta | None = ...,
        prev_repository: builtins.str = ...,
        new_repository: builtins.str = ...,
        prev_commit_id: builtins.str = ...,
        new_commit_id: builtins.str = ...,
        commit_analysis: global___CommitAnalysis | None = ...,
        prev_service_id: builtins.str = ...,
        prev_release_channel_id: builtins.str = ...,
        prev_service_version: builtins.str = ...,
        new_service_id: builtins.str = ...,
        new_release_channel_id: builtins.str = ...,
        new_service_version: builtins.str = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["commit_analysis", b"commit_analysis", "prev", b"prev"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["commit_analysis", b"commit_analysis", "new_commit_id", b"new_commit_id", "new_release_channel_id", b"new_release_channel_id", "new_repository", b"new_repository", "new_service_id", b"new_service_id", "new_service_version", b"new_service_version", "prev", b"prev", "prev_commit_id", b"prev_commit_id", "prev_release_channel_id", b"prev_release_channel_id", "prev_repository", b"prev_repository", "prev_service_id", b"prev_service_id", "prev_service_version", b"prev_service_version"]) -> None: ...

global___DeploymentComparison = DeploymentComparison

class Deployment(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    META_FIELD_NUMBER: builtins.int
    CONFIG_FIELD_NUMBER: builtins.int
    COMPARISON_FIELD_NUMBER: builtins.int
    STATE_FIELD_NUMBER: builtins.int
    @property
    def meta(self) -> prodvana.proto.prodvana.object.meta_pb2.ObjectMeta: ...
    @property
    def config(self) -> global___DeploymentConfig: ...
    @property
    def comparison(self) -> global___DeploymentComparison:
        """TODO(naphat) should this really be part of the proto here, or should it be a separate endpoint so we can request arbitrary comparison?"""
    @property
    def state(self) -> global___DeploymentState: ...
    def __init__(
        self,
        *,
        meta: prodvana.proto.prodvana.object.meta_pb2.ObjectMeta | None = ...,
        config: global___DeploymentConfig | None = ...,
        comparison: global___DeploymentComparison | None = ...,
        state: global___DeploymentState | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["comparison", b"comparison", "config", b"config", "meta", b"meta", "state", b"state"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["comparison", b"comparison", "config", b"config", "meta", b"meta", "state", b"state"]) -> None: ...

global___Deployment = Deployment