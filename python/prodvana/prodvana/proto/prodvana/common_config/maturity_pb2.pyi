"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import builtins
import google.protobuf.descriptor
import google.protobuf.internal.enum_type_wrapper
import sys
import typing

if sys.version_info >= (3, 10):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

class _Maturity:
    ValueType = typing.NewType("ValueType", builtins.int)
    V: typing_extensions.TypeAlias = ValueType

class _MaturityEnumTypeWrapper(google.protobuf.internal.enum_type_wrapper._EnumTypeWrapper[_Maturity.ValueType], builtins.type):  # noqa: F821
    DESCRIPTOR: google.protobuf.descriptor.EnumDescriptor
    FAST_ITERATION: _Maturity.ValueType  # 0
    PRELAUNCH: _Maturity.ValueType  # 1
    LIVE: _Maturity.ValueType  # 2

class Maturity(_Maturity, metaclass=_MaturityEnumTypeWrapper):
    """TODO(naphat) maturity likely should not be hardcoded but be flexible per organization"""

FAST_ITERATION: Maturity.ValueType  # 0
PRELAUNCH: Maturity.ValueType  # 1
LIVE: Maturity.ValueType  # 2
global___Maturity = Maturity
