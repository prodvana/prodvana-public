"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import builtins
import collections.abc
import google.protobuf.descriptor
import google.protobuf.internal.containers
import google.protobuf.message
import prodvana.proto.prodvana.object.meta_pb2
import prodvana.proto.prodvana.release.object_pb2
import sys

if sys.version_info >= (3, 8):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

class RecordReleaseReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    CONFIG_FIELD_NUMBER: builtins.int
    PENDING_FIELD_NUMBER: builtins.int
    @property
    def config(self) -> prodvana.proto.prodvana.release.object_pb2.ReleaseConfig: ...
    pending: builtins.bool
    """If true, create release with pending status that is meant to be updated later to either success or failure.
    By default, releases are created with status SUCCEEDED.
    """
    def __init__(
        self,
        *,
        config: prodvana.proto.prodvana.release.object_pb2.ReleaseConfig | None = ...,
        pending: builtins.bool = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["config", b"config"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["config", b"config", "pending", b"pending"]) -> None: ...

global___RecordReleaseReq = RecordReleaseReq

class RecordReleaseResp(google.protobuf.message.Message):
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

global___RecordReleaseResp = RecordReleaseResp

class UpdateReleaseStatusReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    RELEASE_ID_FIELD_NUMBER: builtins.int
    STATUS_FIELD_NUMBER: builtins.int
    release_id: builtins.str
    status: prodvana.proto.prodvana.release.object_pb2.ReleaseStatus.ValueType
    def __init__(
        self,
        *,
        release_id: builtins.str = ...,
        status: prodvana.proto.prodvana.release.object_pb2.ReleaseStatus.ValueType = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["release_id", b"release_id", "status", b"status"]) -> None: ...

global___UpdateReleaseStatusReq = UpdateReleaseStatusReq

class UpdateReleaseStatusResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    STATUS_FIELD_NUMBER: builtins.int
    status: prodvana.proto.prodvana.release.object_pb2.ReleaseStatus.ValueType
    def __init__(
        self,
        *,
        status: prodvana.proto.prodvana.release.object_pb2.ReleaseStatus.ValueType = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["status", b"status"]) -> None: ...

global___UpdateReleaseStatusResp = UpdateReleaseStatusResp

class ListReleasesReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    class Filter(google.protobuf.message.Message):
        DESCRIPTOR: google.protobuf.descriptor.Descriptor

        SERVICES_FIELD_NUMBER: builtins.int
        RELEASE_CHANNELS_FIELD_NUMBER: builtins.int
        APPLICATION_FIELD_NUMBER: builtins.int
        DESIRED_STATE_ID_FIELD_NUMBER: builtins.int
        @property
        def services(self) -> google.protobuf.internal.containers.RepeatedScalarFieldContainer[builtins.str]:
            """filters to releases for join(join(services, OR), join(release_channels, OR), AND)"""
        @property
        def release_channels(self) -> google.protobuf.internal.containers.RepeatedScalarFieldContainer[builtins.str]: ...
        application: builtins.str
        """if set, the filters on services and release channels above will be limited to this app.
        Otherwise, there is no app filter (so e.g. release_channels=['staging'] will select
        staging release channels across all apps.)
        """
        desired_state_id: builtins.str
        """desired_state_id filter is AND'ed with everything else in the same filter object"""
        def __init__(
            self,
            *,
            services: collections.abc.Iterable[builtins.str] | None = ...,
            release_channels: collections.abc.Iterable[builtins.str] | None = ...,
            application: builtins.str = ...,
            desired_state_id: builtins.str = ...,
        ) -> None: ...
        def ClearField(self, field_name: typing_extensions.Literal["application", b"application", "desired_state_id", b"desired_state_id", "release_channels", b"release_channels", "services", b"services"]) -> None: ...

    FILTERS_FIELD_NUMBER: builtins.int
    FILTER_FIELD_NUMBER: builtins.int
    STARTING_RELEASE_ID_FIELD_NUMBER: builtins.int
    ENDING_RELEASE_ID_FIELD_NUMBER: builtins.int
    PAGE_TOKEN_FIELD_NUMBER: builtins.int
    PAGE_SIZE_FIELD_NUMBER: builtins.int
    @property
    def filters(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[global___ListReleasesReq.Filter]:
        """filters for listing releases. Multiple filters are OR'ed together."""
    @property
    def filter(self) -> global___ListReleasesReq.Filter:
        """escape hatch to support openAPI, which cannot handle repeated list of messages on GET requests. This is joined to the filters list with an OR."""
    starting_release_id: builtins.str
    """newer release, inclusive"""
    ending_release_id: builtins.str
    """older release, exclusive"""
    page_token: builtins.str
    page_size: builtins.int
    def __init__(
        self,
        *,
        filters: collections.abc.Iterable[global___ListReleasesReq.Filter] | None = ...,
        filter: global___ListReleasesReq.Filter | None = ...,
        starting_release_id: builtins.str = ...,
        ending_release_id: builtins.str = ...,
        page_token: builtins.str = ...,
        page_size: builtins.int = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["filter", b"filter"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["ending_release_id", b"ending_release_id", "filter", b"filter", "filters", b"filters", "page_size", b"page_size", "page_token", b"page_token", "starting_release_id", b"starting_release_id"]) -> None: ...

global___ListReleasesReq = ListReleasesReq

class ListReleasesResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    RELEASES_FIELD_NUMBER: builtins.int
    NEXT_PAGE_TOKEN_FIELD_NUMBER: builtins.int
    @property
    def releases(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[prodvana.proto.prodvana.release.object_pb2.Release]: ...
    next_page_token: builtins.str
    def __init__(
        self,
        *,
        releases: collections.abc.Iterable[prodvana.proto.prodvana.release.object_pb2.Release] | None = ...,
        next_page_token: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["next_page_token", b"next_page_token", "releases", b"releases"]) -> None: ...

global___ListReleasesResp = ListReleasesResp

class CompareReleaseReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    class ReleaseRef(google.protobuf.message.Message):
        DESCRIPTOR: google.protobuf.descriptor.Descriptor

        RELEASE_ID_FIELD_NUMBER: builtins.int
        CONFIG_FIELD_NUMBER: builtins.int
        release_id: builtins.str
        @property
        def config(self) -> prodvana.proto.prodvana.release.object_pb2.ReleaseConfig: ...
        def __init__(
            self,
            *,
            release_id: builtins.str = ...,
            config: prodvana.proto.prodvana.release.object_pb2.ReleaseConfig | None = ...,
        ) -> None: ...
        def HasField(self, field_name: typing_extensions.Literal["config", b"config", "ref", b"ref", "release_id", b"release_id"]) -> builtins.bool: ...
        def ClearField(self, field_name: typing_extensions.Literal["config", b"config", "ref", b"ref", "release_id", b"release_id"]) -> None: ...
        def WhichOneof(self, oneof_group: typing_extensions.Literal["ref", b"ref"]) -> typing_extensions.Literal["release_id", "config"] | None: ...

    NEW_RELEASE_FIELD_NUMBER: builtins.int
    PREV_RELEASE_FIELD_NUMBER: builtins.int
    @property
    def new_release(self) -> global___CompareReleaseReq.ReleaseRef: ...
    @property
    def prev_release(self) -> global___CompareReleaseReq.ReleaseRef: ...
    def __init__(
        self,
        *,
        new_release: global___CompareReleaseReq.ReleaseRef | None = ...,
        prev_release: global___CompareReleaseReq.ReleaseRef | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["new_release", b"new_release", "prev_release", b"prev_release"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["new_release", b"new_release", "prev_release", b"prev_release"]) -> None: ...

global___CompareReleaseReq = CompareReleaseReq

class CompareReleaseResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    COMPARISON_FIELD_NUMBER: builtins.int
    @property
    def comparison(self) -> prodvana.proto.prodvana.release.object_pb2.ReleaseComparison: ...
    def __init__(
        self,
        *,
        comparison: prodvana.proto.prodvana.release.object_pb2.ReleaseComparison | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["comparison", b"comparison"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["comparison", b"comparison"]) -> None: ...

global___CompareReleaseResp = CompareReleaseResp
