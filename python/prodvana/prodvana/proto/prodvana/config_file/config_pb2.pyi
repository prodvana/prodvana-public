"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import builtins
import google.protobuf.descriptor
import google.protobuf.message
import prodvana.proto.prodvana.application.application_config_pb2
import prodvana.proto.prodvana.application.user_metadata_pb2
import prodvana.proto.prodvana.delivery_extension.config_pb2
import prodvana.proto.prodvana.environment.clusters_pb2
import prodvana.proto.prodvana.protection.protection_config_pb2
import prodvana.proto.prodvana.release_channel.release_channel_config_pb2
import prodvana.proto.prodvana.service.service_config_pb2
import prodvana.proto.prodvana.service.user_metadata_pb2
import sys
import typing

if sys.version_info >= (3, 8):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

class ProdvanaConfig(google.protobuf.message.Message):
    """schema for config.pvn.yaml"""

    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    APPLICATION_FIELD_NUMBER: builtins.int
    SERVICE_FIELD_NUMBER: builtins.int
    PROTECTION_FIELD_NUMBER: builtins.int
    RUNTIME_FIELD_NUMBER: builtins.int
    DELIVERY_EXTENSION_FIELD_NUMBER: builtins.int
    RELEASE_CHANNEL_FIELD_NUMBER: builtins.int
    APPLICATION_METADATA_FIELD_NUMBER: builtins.int
    SERVICE_METADATA_FIELD_NUMBER: builtins.int
    @property
    def application(self) -> prodvana.proto.prodvana.application.application_config_pb2.ApplicationConfig: ...
    @property
    def service(self) -> prodvana.proto.prodvana.service.service_config_pb2.ServiceConfig: ...
    @property
    def protection(self) -> prodvana.proto.prodvana.protection.protection_config_pb2.ProtectionConfig: ...
    @property
    def runtime(self) -> prodvana.proto.prodvana.environment.clusters_pb2.ClusterConfig: ...
    @property
    def delivery_extension(self) -> prodvana.proto.prodvana.delivery_extension.config_pb2.DeliveryExtensionConfig: ...
    @property
    def release_channel(self) -> prodvana.proto.prodvana.release_channel.release_channel_config_pb2.ReleaseChannelConfig: ...
    @property
    def application_metadata(self) -> prodvana.proto.prodvana.application.user_metadata_pb2.ApplicationUserMetadata: ...
    @property
    def service_metadata(self) -> prodvana.proto.prodvana.service.user_metadata_pb2.ServiceUserMetadata: ...
    def __init__(
        self,
        *,
        application: prodvana.proto.prodvana.application.application_config_pb2.ApplicationConfig | None = ...,
        service: prodvana.proto.prodvana.service.service_config_pb2.ServiceConfig | None = ...,
        protection: prodvana.proto.prodvana.protection.protection_config_pb2.ProtectionConfig | None = ...,
        runtime: prodvana.proto.prodvana.environment.clusters_pb2.ClusterConfig | None = ...,
        delivery_extension: prodvana.proto.prodvana.delivery_extension.config_pb2.DeliveryExtensionConfig | None = ...,
        release_channel: prodvana.proto.prodvana.release_channel.release_channel_config_pb2.ReleaseChannelConfig | None = ...,
        application_metadata: prodvana.proto.prodvana.application.user_metadata_pb2.ApplicationUserMetadata | None = ...,
        service_metadata: prodvana.proto.prodvana.service.user_metadata_pb2.ServiceUserMetadata | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["application", b"application", "application_metadata", b"application_metadata", "config_oneof", b"config_oneof", "delivery_extension", b"delivery_extension", "metadata_oneof", b"metadata_oneof", "protection", b"protection", "release_channel", b"release_channel", "runtime", b"runtime", "service", b"service", "service_metadata", b"service_metadata"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["application", b"application", "application_metadata", b"application_metadata", "config_oneof", b"config_oneof", "delivery_extension", b"delivery_extension", "metadata_oneof", b"metadata_oneof", "protection", b"protection", "release_channel", b"release_channel", "runtime", b"runtime", "service", b"service", "service_metadata", b"service_metadata"]) -> None: ...
    @typing.overload
    def WhichOneof(self, oneof_group: typing_extensions.Literal["config_oneof", b"config_oneof"]) -> typing_extensions.Literal["application", "service", "protection", "runtime", "delivery_extension", "release_channel"] | None: ...
    @typing.overload
    def WhichOneof(self, oneof_group: typing_extensions.Literal["metadata_oneof", b"metadata_oneof"]) -> typing_extensions.Literal["application_metadata", "service_metadata"] | None: ...

global___ProdvanaConfig = ProdvanaConfig
