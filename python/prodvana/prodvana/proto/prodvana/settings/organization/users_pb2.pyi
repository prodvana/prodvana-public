"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import builtins
import collections.abc
import google.protobuf.descriptor
import google.protobuf.internal.containers
import google.protobuf.message
import prodvana.proto.prodvana.users.users_pb2
import sys

if sys.version_info >= (3, 8):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

class SettingsUser(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    USER_FIELD_NUMBER: builtins.int
    ROLES_FIELD_NUMBER: builtins.int
    @property
    def user(self) -> prodvana.proto.prodvana.users.users_pb2.User: ...
    @property
    def roles(self) -> google.protobuf.internal.containers.RepeatedScalarFieldContainer[builtins.str]: ...
    def __init__(
        self,
        *,
        user: prodvana.proto.prodvana.users.users_pb2.User | None = ...,
        roles: collections.abc.Iterable[builtins.str] | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["user", b"user"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["roles", b"roles", "user", b"user"]) -> None: ...

global___SettingsUser = SettingsUser

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
    def user(self) -> global___SettingsUser: ...
    def __init__(
        self,
        *,
        user: global___SettingsUser | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["user", b"user"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["user", b"user"]) -> None: ...

global___GetUserResp = GetUserResp

class ListRolesReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    def __init__(
        self,
    ) -> None: ...

global___ListRolesReq = ListRolesReq

class RoleDefinition(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    NAME_FIELD_NUMBER: builtins.int
    DESCRIPTION_FIELD_NUMBER: builtins.int
    name: builtins.str
    description: builtins.str
    def __init__(
        self,
        *,
        name: builtins.str = ...,
        description: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["description", b"description", "name", b"name"]) -> None: ...

global___RoleDefinition = RoleDefinition

class ListRolesResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    ROLES_FIELD_NUMBER: builtins.int
    @property
    def roles(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[global___RoleDefinition]: ...
    def __init__(
        self,
        *,
        roles: collections.abc.Iterable[global___RoleDefinition] | None = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["roles", b"roles"]) -> None: ...

global___ListRolesResp = ListRolesResp

class ListUsersReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    PAGE_TOKEN_FIELD_NUMBER: builtins.int
    PAGE_SIZE_FIELD_NUMBER: builtins.int
    page_token: builtins.str
    page_size: builtins.int
    def __init__(
        self,
        *,
        page_token: builtins.str = ...,
        page_size: builtins.int = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["page_size", b"page_size", "page_token", b"page_token"]) -> None: ...

global___ListUsersReq = ListUsersReq

class ListUsersResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    USERS_FIELD_NUMBER: builtins.int
    NEXT_PAGE_TOKEN_FIELD_NUMBER: builtins.int
    @property
    def users(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[global___SettingsUser]: ...
    next_page_token: builtins.str
    def __init__(
        self,
        *,
        users: collections.abc.Iterable[global___SettingsUser] | None = ...,
        next_page_token: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["next_page_token", b"next_page_token", "users", b"users"]) -> None: ...

global___ListUsersResp = ListUsersResp

class SetRolesReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    USER_ID_FIELD_NUMBER: builtins.int
    ROLES_FIELD_NUMBER: builtins.int
    user_id: builtins.str
    @property
    def roles(self) -> google.protobuf.internal.containers.RepeatedScalarFieldContainer[builtins.str]: ...
    def __init__(
        self,
        *,
        user_id: builtins.str = ...,
        roles: collections.abc.Iterable[builtins.str] | None = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["roles", b"roles", "user_id", b"user_id"]) -> None: ...

global___SetRolesReq = SetRolesReq

class SetRolesResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    def __init__(
        self,
    ) -> None: ...

global___SetRolesResp = SetRolesResp

class UserInvite(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    EMAIL_FIELD_NUMBER: builtins.int
    ROLES_FIELD_NUMBER: builtins.int
    email: builtins.str
    @property
    def roles(self) -> google.protobuf.internal.containers.RepeatedScalarFieldContainer[builtins.str]: ...
    def __init__(
        self,
        *,
        email: builtins.str = ...,
        roles: collections.abc.Iterable[builtins.str] | None = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["email", b"email", "roles", b"roles"]) -> None: ...

global___UserInvite = UserInvite

class InviteUsersReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    USERS_FIELD_NUMBER: builtins.int
    @property
    def users(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[global___UserInvite]: ...
    def __init__(
        self,
        *,
        users: collections.abc.Iterable[global___UserInvite] | None = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["users", b"users"]) -> None: ...

global___InviteUsersReq = InviteUsersReq

class InviteUsersResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    def __init__(
        self,
    ) -> None: ...

global___InviteUsersResp = InviteUsersResp

class OrganizationSupportsInvitesReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    def __init__(
        self,
    ) -> None: ...

global___OrganizationSupportsInvitesReq = OrganizationSupportsInvitesReq

class OrganizationSupportsInvitesResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    SUPPORTS_INVITES_FIELD_NUMBER: builtins.int
    supports_invites: builtins.bool
    def __init__(
        self,
        *,
        supports_invites: builtins.bool = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["supports_invites", b"supports_invites"]) -> None: ...

global___OrganizationSupportsInvitesResp = OrganizationSupportsInvitesResp
