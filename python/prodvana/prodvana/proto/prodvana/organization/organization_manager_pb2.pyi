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
import google.protobuf.timestamp_pb2
import prodvana.proto.prodvana.config_writeback.writeback_pb2
import prodvana.proto.prodvana.insights.insights_pb2
import prodvana.proto.prodvana.metrics.metrics_pb2
import prodvana.proto.prodvana.organization.user_metadata_pb2
import prodvana.proto.prodvana.users.users_pb2
import sys

if sys.version_info >= (3, 8):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

class OrganizationInfo(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    ID_FIELD_NUMBER: builtins.int
    DISPLAY_NAME_FIELD_NUMBER: builtins.int
    WRITEBACK_CONFIG_FIELD_NUMBER: builtins.int
    SLUG_FIELD_NUMBER: builtins.int
    USER_METADATA_FIELD_NUMBER: builtins.int
    id: builtins.str
    display_name: builtins.str
    @property
    def writeback_config(self) -> prodvana.proto.prodvana.config_writeback.writeback_pb2.ConfigWritebackPath: ...
    slug: builtins.str
    @property
    def user_metadata(self) -> prodvana.proto.prodvana.organization.user_metadata_pb2.OrganizationUserMetadata: ...
    def __init__(
        self,
        *,
        id: builtins.str = ...,
        display_name: builtins.str = ...,
        writeback_config: prodvana.proto.prodvana.config_writeback.writeback_pb2.ConfigWritebackPath | None = ...,
        slug: builtins.str = ...,
        user_metadata: prodvana.proto.prodvana.organization.user_metadata_pb2.OrganizationUserMetadata | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["user_metadata", b"user_metadata", "writeback_config", b"writeback_config"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["display_name", b"display_name", "id", b"id", "slug", b"slug", "user_metadata", b"user_metadata", "writeback_config", b"writeback_config"]) -> None: ...

global___OrganizationInfo = OrganizationInfo

class GetOrganizationReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    def __init__(
        self,
    ) -> None: ...

global___GetOrganizationReq = GetOrganizationReq

class GetOrganizationResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    ORGANIZATION_FIELD_NUMBER: builtins.int
    @property
    def organization(self) -> global___OrganizationInfo: ...
    def __init__(
        self,
        *,
        organization: global___OrganizationInfo | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["organization", b"organization"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["organization", b"organization"]) -> None: ...

global___GetOrganizationResp = GetOrganizationResp

class GetOrganizationMetricsReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    START_TIMESTAMP_FIELD_NUMBER: builtins.int
    END_TIMESTAMP_FIELD_NUMBER: builtins.int
    @property
    def start_timestamp(self) -> google.protobuf.timestamp_pb2.Timestamp: ...
    @property
    def end_timestamp(self) -> google.protobuf.timestamp_pb2.Timestamp: ...
    def __init__(
        self,
        *,
        start_timestamp: google.protobuf.timestamp_pb2.Timestamp | None = ...,
        end_timestamp: google.protobuf.timestamp_pb2.Timestamp | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["end_timestamp", b"end_timestamp", "start_timestamp", b"start_timestamp"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["end_timestamp", b"end_timestamp", "start_timestamp", b"start_timestamp"]) -> None: ...

global___GetOrganizationMetricsReq = GetOrganizationMetricsReq

class GetOrganizationMetricsResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    DEPLOYMENT_METRICS_FIELD_NUMBER: builtins.int
    @property
    def deployment_metrics(self) -> prodvana.proto.prodvana.metrics.metrics_pb2.DeploymentMetrics: ...
    def __init__(
        self,
        *,
        deployment_metrics: prodvana.proto.prodvana.metrics.metrics_pb2.DeploymentMetrics | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["deployment_metrics", b"deployment_metrics"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["deployment_metrics", b"deployment_metrics"]) -> None: ...

global___GetOrganizationMetricsResp = GetOrganizationMetricsResp

class GetOrganizationInsightsReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    def __init__(
        self,
    ) -> None: ...

global___GetOrganizationInsightsReq = GetOrganizationInsightsReq

class GetOrganizationInsightsResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    INSIGHTS_FIELD_NUMBER: builtins.int
    @property
    def insights(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[prodvana.proto.prodvana.insights.insights_pb2.Insight]: ...
    def __init__(
        self,
        *,
        insights: collections.abc.Iterable[prodvana.proto.prodvana.insights.insights_pb2.Insight] | None = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["insights", b"insights"]) -> None: ...

global___GetOrganizationInsightsResp = GetOrganizationInsightsResp

class SnoozeOrganizationInsightReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    CLASS_FIELD_NUMBER: builtins.int
    DURATION_FIELD_NUMBER: builtins.int
    @property
    def duration(self) -> google.protobuf.duration_pb2.Duration: ...
    def __init__(
        self,
        *,
        duration: google.protobuf.duration_pb2.Duration | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["duration", b"duration"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["class", b"class", "duration", b"duration"]) -> None: ...

global___SnoozeOrganizationInsightReq = SnoozeOrganizationInsightReq

class SnoozeOrganizationInsightResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    def __init__(
        self,
    ) -> None: ...

global___SnoozeOrganizationInsightResp = SnoozeOrganizationInsightResp

class GetOrganizationMetadataReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    def __init__(
        self,
    ) -> None: ...

global___GetOrganizationMetadataReq = GetOrganizationMetadataReq

class GetOrganizationMetadataResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    METADATA_FIELD_NUMBER: builtins.int
    @property
    def metadata(self) -> prodvana.proto.prodvana.organization.user_metadata_pb2.OrganizationUserMetadata:
        """metadata with no variables substitution, no modifications from parents"""
    def __init__(
        self,
        *,
        metadata: prodvana.proto.prodvana.organization.user_metadata_pb2.OrganizationUserMetadata | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["metadata", b"metadata"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["metadata", b"metadata"]) -> None: ...

global___GetOrganizationMetadataResp = GetOrganizationMetadataResp

class SetOrganizationMetadataReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    METADATA_FIELD_NUMBER: builtins.int
    @property
    def metadata(self) -> prodvana.proto.prodvana.organization.user_metadata_pb2.OrganizationUserMetadata: ...
    def __init__(
        self,
        *,
        metadata: prodvana.proto.prodvana.organization.user_metadata_pb2.OrganizationUserMetadata | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["metadata", b"metadata"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["metadata", b"metadata"]) -> None: ...

global___SetOrganizationMetadataReq = SetOrganizationMetadataReq

class SetOrganizationMetadataResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    def __init__(
        self,
    ) -> None: ...

global___SetOrganizationMetadataResp = SetOrganizationMetadataResp

class GetUserReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    USER_ID_FIELD_NUMBER: builtins.int
    user_id: builtins.str
    def __init__(
        self,
        *,
        user_id: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["user_id", b"user_id"]) -> None: ...

global___GetUserReq = GetUserReq

class GetUserResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    USER_FIELD_NUMBER: builtins.int
    @property
    def user(self) -> prodvana.proto.prodvana.users.users_pb2.User: ...
    def __init__(
        self,
        *,
        user: prodvana.proto.prodvana.users.users_pb2.User | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["user", b"user"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["user", b"user"]) -> None: ...

global___GetUserResp = GetUserResp
