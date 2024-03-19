"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import builtins
import google.protobuf.descriptor
import google.protobuf.message
import prodvana.proto.prodvana.pvn_wrapper.output_pb2
import sys

if sys.version_info >= (3, 8):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

class ReportJobResultReq(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    JOB_ID_FIELD_NUMBER: builtins.int
    OUTPUT_FIELD_NUMBER: builtins.int
    job_id: builtins.str
    @property
    def output(self) -> prodvana.proto.prodvana.pvn_wrapper.output_pb2.Output: ...
    def __init__(
        self,
        *,
        job_id: builtins.str = ...,
        output: prodvana.proto.prodvana.pvn_wrapper.output_pb2.Output | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["output", b"output"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["job_id", b"job_id", "output", b"output"]) -> None: ...

global___ReportJobResultReq = ReportJobResultReq

class ReportJobResultResp(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    def __init__(
        self,
    ) -> None: ...

global___ReportJobResultResp = ReportJobResultResp
