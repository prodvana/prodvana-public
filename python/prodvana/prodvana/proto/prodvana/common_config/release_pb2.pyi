"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import builtins
import google.protobuf.descriptor
import google.protobuf.message
import sys

if sys.version_info >= (3, 8):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

class ReleaseSettings(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    class BypassSettings(google.protobuf.message.Message):
        DESCRIPTOR: google.protobuf.descriptor.Descriptor

        NO_BYPASS_PROTECTIONS_FIELD_NUMBER: builtins.int
        NO_BYPASS_CONVERGENCE_EXTENSIONS_FIELD_NUMBER: builtins.int
        NO_BYPASS_APPROVALS_FIELD_NUMBER: builtins.int
        NO_BYPASS_RELEASE_CHANNEL_ORDERING_FIELD_NUMBER: builtins.int
        no_bypass_protections: builtins.bool
        no_bypass_convergence_extensions: builtins.bool
        no_bypass_approvals: builtins.bool
        no_bypass_release_channel_ordering: builtins.bool
        def __init__(
            self,
            *,
            no_bypass_protections: builtins.bool = ...,
            no_bypass_convergence_extensions: builtins.bool = ...,
            no_bypass_approvals: builtins.bool = ...,
            no_bypass_release_channel_ordering: builtins.bool = ...,
        ) -> None: ...
        def ClearField(self, field_name: typing_extensions.Literal["no_bypass_approvals", b"no_bypass_approvals", "no_bypass_convergence_extensions", b"no_bypass_convergence_extensions", "no_bypass_protections", b"no_bypass_protections", "no_bypass_release_channel_ordering", b"no_bypass_release_channel_ordering"]) -> None: ...

    BYPASS_SETTINGS_FIELD_NUMBER: builtins.int
    @property
    def bypass_settings(self) -> global___ReleaseSettings.BypassSettings:
        """customize what is bypassed in the fast releases (the default for rollbacks). Defaults to everything being bypassed."""
    def __init__(
        self,
        *,
        bypass_settings: global___ReleaseSettings.BypassSettings | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["bypass_settings", b"bypass_settings"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["bypass_settings", b"bypass_settings"]) -> None: ...

global___ReleaseSettings = ReleaseSettings
