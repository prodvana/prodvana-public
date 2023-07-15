"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import builtins
import collections.abc
import google.protobuf.descriptor
import google.protobuf.internal.containers
import google.protobuf.message
import prodvana.proto.prodvana.capability.capability_pb2
import prodvana.proto.prodvana.common_config.notification_pb2
import prodvana.proto.prodvana.release_channel.release_channel_config_pb2
import prodvana.proto.prodvana.workflow.integration_config_pb2
import sys

if sys.version_info >= (3, 8):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

class ApplicationConfig(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    NAME_FIELD_NUMBER: builtins.int
    RELEASE_CHANNELS_FIELD_NUMBER: builtins.int
    NOTIFICATIONS_FIELD_NUMBER: builtins.int
    ALERTS_FIELD_NUMBER: builtins.int
    CAPABILITIES_FIELD_NUMBER: builtins.int
    CAPABILITY_INSTANCES_FIELD_NUMBER: builtins.int
    name: builtins.str
    @property
    def release_channels(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[prodvana.proto.prodvana.release_channel.release_channel_config_pb2.ReleaseChannelConfig]: ...
    @property
    def notifications(self) -> prodvana.proto.prodvana.common_config.notification_pb2.NotificationConfig: ...
    @property
    def alerts(self) -> prodvana.proto.prodvana.workflow.integration_config_pb2.AlertingConfig: ...
    @property
    def capabilities(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[prodvana.proto.prodvana.capability.capability_pb2.CapabilityConfig]:
        """capabilities are dependencies that services in this applications can use"""
    @property
    def capability_instances(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[prodvana.proto.prodvana.capability.capability_pb2.CapabilityInstanceConfig]:
        """optional capability instances that can be referenced by `capabilities`, useful to deduplication.
        for example, you may choose to have two database capability instances, staging and prod,
        and use them across release channels staging, beta, and prod, where beta and prod use the prod db.
        """
    def __init__(
        self,
        *,
        name: builtins.str = ...,
        release_channels: collections.abc.Iterable[prodvana.proto.prodvana.release_channel.release_channel_config_pb2.ReleaseChannelConfig] | None = ...,
        notifications: prodvana.proto.prodvana.common_config.notification_pb2.NotificationConfig | None = ...,
        alerts: prodvana.proto.prodvana.workflow.integration_config_pb2.AlertingConfig | None = ...,
        capabilities: collections.abc.Iterable[prodvana.proto.prodvana.capability.capability_pb2.CapabilityConfig] | None = ...,
        capability_instances: collections.abc.Iterable[prodvana.proto.prodvana.capability.capability_pb2.CapabilityInstanceConfig] | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["alerts", b"alerts", "notifications", b"notifications"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["alerts", b"alerts", "capabilities", b"capabilities", "capability_instances", b"capability_instances", "name", b"name", "notifications", b"notifications", "release_channels", b"release_channels"]) -> None: ...

global___ApplicationConfig = ApplicationConfig
