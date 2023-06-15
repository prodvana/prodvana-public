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
import prodvana.proto.prodvana.application.application_config_pb2
import prodvana.proto.prodvana.application.object_pb2
import prodvana.proto.prodvana.application.user_metadata_pb2
import prodvana.proto.prodvana.common_config.dangerous_action_pb2
import prodvana.proto.prodvana.insights.insights_pb2
import prodvana.proto.prodvana.metrics.metrics_pb2
import prodvana.proto.prodvana.object.meta_pb2
import prodvana.proto.prodvana.version.source_metadata_pb2
import sys

if sys.version_info >= (3, 8):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

class ListApplicationsReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    DETAILED_FIELD_NUMBER: builtins.int
    detailed: builtins.bool
    """if not set, only meta without version is returned"""
    def __init__(
        self,
        *,
        detailed: builtins.bool = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["detailed", b"detailed"]) -> None: ...

global___ListApplicationsReq = ListApplicationsReq

class ListApplicationsResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    APPLICATIONS_FIELD_NUMBER: builtins.int
    @property
    def applications(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[prodvana.proto.prodvana.application.object_pb2.Application]: ...
    def __init__(
        self,
        *,
        applications: collections.abc.Iterable[prodvana.proto.prodvana.application.object_pb2.Application] | None = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["applications", b"applications"]) -> None: ...

global___ListApplicationsResp = ListApplicationsResp

class GetApplicationReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    APPLICATION_FIELD_NUMBER: builtins.int
    application: builtins.str
    def __init__(
        self,
        *,
        application: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["application", b"application"]) -> None: ...

global___GetApplicationReq = GetApplicationReq

class GetApplicationResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    APPLICATION_FIELD_NUMBER: builtins.int
    @property
    def application(self) -> prodvana.proto.prodvana.application.object_pb2.Application: ...
    def __init__(
        self,
        *,
        application: prodvana.proto.prodvana.application.object_pb2.Application | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["application", b"application"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["application", b"application"]) -> None: ...

global___GetApplicationResp = GetApplicationResp

class ConfigureApplicationReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    APPLICATION_CONFIG_FIELD_NUMBER: builtins.int
    APPROVED_DANGEROUS_ACTION_IDS_FIELD_NUMBER: builtins.int
    SOURCE_FIELD_NUMBER: builtins.int
    SOURCE_METADATA_FIELD_NUMBER: builtins.int
    BASE_VERSION_FIELD_NUMBER: builtins.int
    @property
    def application_config(self) -> prodvana.proto.prodvana.application.application_config_pb2.ApplicationConfig: ...
    @property
    def approved_dangerous_action_ids(self) -> google.protobuf.internal.containers.RepeatedScalarFieldContainer[builtins.str]: ...
    source: prodvana.proto.prodvana.version.source_metadata_pb2.Source.ValueType
    @property
    def source_metadata(self) -> prodvana.proto.prodvana.version.source_metadata_pb2.SourceMetadata: ...
    base_version: builtins.str
    """optional, if this is not a new application, this is the existing version of 
    this application config that this configuration was based on.
    this can be used to avoid concurrent updates overwriting each other.
    NOTE: this will eventually be required for updates
    """
    def __init__(
        self,
        *,
        application_config: prodvana.proto.prodvana.application.application_config_pb2.ApplicationConfig | None = ...,
        approved_dangerous_action_ids: collections.abc.Iterable[builtins.str] | None = ...,
        source: prodvana.proto.prodvana.version.source_metadata_pb2.Source.ValueType = ...,
        source_metadata: prodvana.proto.prodvana.version.source_metadata_pb2.SourceMetadata | None = ...,
        base_version: builtins.str = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["application_config", b"application_config", "source_metadata", b"source_metadata"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["application_config", b"application_config", "approved_dangerous_action_ids", b"approved_dangerous_action_ids", "base_version", b"base_version", "source", b"source", "source_metadata", b"source_metadata"]) -> None: ...

global___ConfigureApplicationReq = ConfigureApplicationReq

class ConfigureApplicationResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    META_FIELD_NUMBER: builtins.int
    @property
    def meta(self) -> prodvana.proto.prodvana.object.meta_pb2.ObjectMeta: ...
    def __init__(
        self,
        *,
        meta: prodvana.proto.prodvana.object.meta_pb2.ObjectMeta | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["meta", b"meta"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["meta", b"meta"]) -> None: ...

global___ConfigureApplicationResp = ConfigureApplicationResp

class ValidateConfigureApplicationResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    CONFIG_FIELD_NUMBER: builtins.int
    COMPILED_CONFIG_FIELD_NUMBER: builtins.int
    DANGEROUS_ACTIONS_FIELD_NUMBER: builtins.int
    @property
    def config(self) -> prodvana.proto.prodvana.application.application_config_pb2.ApplicationConfig:
        """config as passed in by user"""
    @property
    def compiled_config(self) -> prodvana.proto.prodvana.application.application_config_pb2.ApplicationConfig:
        """config with defaults applied"""
    @property
    def dangerous_actions(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[prodvana.proto.prodvana.common_config.dangerous_action_pb2.DangerousAction]: ...
    def __init__(
        self,
        *,
        config: prodvana.proto.prodvana.application.application_config_pb2.ApplicationConfig | None = ...,
        compiled_config: prodvana.proto.prodvana.application.application_config_pb2.ApplicationConfig | None = ...,
        dangerous_actions: collections.abc.Iterable[prodvana.proto.prodvana.common_config.dangerous_action_pb2.DangerousAction] | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["compiled_config", b"compiled_config", "config", b"config"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["compiled_config", b"compiled_config", "config", b"config", "dangerous_actions", b"dangerous_actions"]) -> None: ...

global___ValidateConfigureApplicationResp = ValidateConfigureApplicationResp

class GetApplicationConfigReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    APPLICATION_FIELD_NUMBER: builtins.int
    VERSION_FIELD_NUMBER: builtins.int
    application: builtins.str
    version: builtins.str
    """omit to get latest version"""
    def __init__(
        self,
        *,
        application: builtins.str = ...,
        version: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["application", b"application", "version", b"version"]) -> None: ...

global___GetApplicationConfigReq = GetApplicationConfigReq

class GetApplicationConfigResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    CONFIG_FIELD_NUMBER: builtins.int
    VERSION_FIELD_NUMBER: builtins.int
    COMPILED_CONFIG_FIELD_NUMBER: builtins.int
    @property
    def config(self) -> prodvana.proto.prodvana.application.application_config_pb2.ApplicationConfig:
        """config as passed in by user"""
    version: builtins.str
    @property
    def compiled_config(self) -> prodvana.proto.prodvana.application.application_config_pb2.ApplicationConfig:
        """config with defaults applied"""
    def __init__(
        self,
        *,
        config: prodvana.proto.prodvana.application.application_config_pb2.ApplicationConfig | None = ...,
        version: builtins.str = ...,
        compiled_config: prodvana.proto.prodvana.application.application_config_pb2.ApplicationConfig | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["compiled_config", b"compiled_config", "config", b"config"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["compiled_config", b"compiled_config", "config", b"config", "version", b"version"]) -> None: ...

global___GetApplicationConfigResp = GetApplicationConfigResp

class DeleteApplicationReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    APPLICATION_FIELD_NUMBER: builtins.int
    application: builtins.str
    def __init__(
        self,
        *,
        application: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["application", b"application"]) -> None: ...

global___DeleteApplicationReq = DeleteApplicationReq

class DeleteApplicationResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    def __init__(
        self,
    ) -> None: ...

global___DeleteApplicationResp = DeleteApplicationResp

class GetApplicationMetricsReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    APPLICATION_FIELD_NUMBER: builtins.int
    START_TIMESTAMP_FIELD_NUMBER: builtins.int
    END_TIMESTAMP_FIELD_NUMBER: builtins.int
    application: builtins.str
    @property
    def start_timestamp(self) -> google.protobuf.timestamp_pb2.Timestamp: ...
    @property
    def end_timestamp(self) -> google.protobuf.timestamp_pb2.Timestamp: ...
    def __init__(
        self,
        *,
        application: builtins.str = ...,
        start_timestamp: google.protobuf.timestamp_pb2.Timestamp | None = ...,
        end_timestamp: google.protobuf.timestamp_pb2.Timestamp | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["end_timestamp", b"end_timestamp", "start_timestamp", b"start_timestamp"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["application", b"application", "end_timestamp", b"end_timestamp", "start_timestamp", b"start_timestamp"]) -> None: ...

global___GetApplicationMetricsReq = GetApplicationMetricsReq

class GetApplicationMetricsResp(google.protobuf.message.Message):
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

global___GetApplicationMetricsResp = GetApplicationMetricsResp

class GetApplicationInsightsReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    APPLICATION_FIELD_NUMBER: builtins.int
    application: builtins.str
    def __init__(
        self,
        *,
        application: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["application", b"application"]) -> None: ...

global___GetApplicationInsightsReq = GetApplicationInsightsReq

class GetApplicationInsightsResp(google.protobuf.message.Message):
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

global___GetApplicationInsightsResp = GetApplicationInsightsResp

class SnoozeApplicationInsightReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    APPLICATION_FIELD_NUMBER: builtins.int
    CLASS_FIELD_NUMBER: builtins.int
    DURATION_FIELD_NUMBER: builtins.int
    application: builtins.str
    @property
    def duration(self) -> google.protobuf.duration_pb2.Duration: ...
    def __init__(
        self,
        *,
        application: builtins.str = ...,
        duration: google.protobuf.duration_pb2.Duration | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["duration", b"duration"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["application", b"application", "class", b"class", "duration", b"duration"]) -> None: ...

global___SnoozeApplicationInsightReq = SnoozeApplicationInsightReq

class SnoozeApplicationInsightResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    def __init__(
        self,
    ) -> None: ...

global___SnoozeApplicationInsightResp = SnoozeApplicationInsightResp

class GetApplicationMetadataReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    APPLICATION_FIELD_NUMBER: builtins.int
    application: builtins.str
    def __init__(
        self,
        *,
        application: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["application", b"application"]) -> None: ...

global___GetApplicationMetadataReq = GetApplicationMetadataReq

class GetApplicationMetadataResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    METADATA_FIELD_NUMBER: builtins.int
    @property
    def metadata(self) -> prodvana.proto.prodvana.application.user_metadata_pb2.ApplicationUserMetadata:
        """metadata with no variables substitution, no modifications from parents"""
    def __init__(
        self,
        *,
        metadata: prodvana.proto.prodvana.application.user_metadata_pb2.ApplicationUserMetadata | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["metadata", b"metadata"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["metadata", b"metadata"]) -> None: ...

global___GetApplicationMetadataResp = GetApplicationMetadataResp

class SetApplicationMetadataReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    APPLICATION_FIELD_NUMBER: builtins.int
    METADATA_FIELD_NUMBER: builtins.int
    application: builtins.str
    @property
    def metadata(self) -> prodvana.proto.prodvana.application.user_metadata_pb2.ApplicationUserMetadata: ...
    def __init__(
        self,
        *,
        application: builtins.str = ...,
        metadata: prodvana.proto.prodvana.application.user_metadata_pb2.ApplicationUserMetadata | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["metadata", b"metadata"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["application", b"application", "metadata", b"metadata"]) -> None: ...

global___SetApplicationMetadataReq = SetApplicationMetadataReq

class SetApplicationMetadataResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    def __init__(
        self,
    ) -> None: ...

global___SetApplicationMetadataResp = SetApplicationMetadataResp
