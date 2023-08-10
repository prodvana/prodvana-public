"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import abc
import collections.abc
import grpc
import prodvana.proto.prodvana.blobs.blobs_manager_pb2

class BlobsManagerStub:
    def __init__(self, channel: grpc.Channel) -> None: ...
    GetCasBlob: grpc.UnaryStreamMultiCallable[
        prodvana.proto.prodvana.blobs.blobs_manager_pb2.GetCasBlobReq,
        prodvana.proto.prodvana.blobs.blobs_manager_pb2.GetCasBlobResp,
    ]
    UploadCasBlob: grpc.StreamUnaryMultiCallable[
        prodvana.proto.prodvana.blobs.blobs_manager_pb2.UploadCasBlobReq,
        prodvana.proto.prodvana.blobs.blobs_manager_pb2.UploadCasBlobResp,
    ]

class BlobsManagerServicer(metaclass=abc.ABCMeta):
    @abc.abstractmethod
    def GetCasBlob(
        self,
        request: prodvana.proto.prodvana.blobs.blobs_manager_pb2.GetCasBlobReq,
        context: grpc.ServicerContext,
    ) -> collections.abc.Iterator[prodvana.proto.prodvana.blobs.blobs_manager_pb2.GetCasBlobResp]: ...
    @abc.abstractmethod
    def UploadCasBlob(
        self,
        request_iterator: collections.abc.Iterator[prodvana.proto.prodvana.blobs.blobs_manager_pb2.UploadCasBlobReq],
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.blobs.blobs_manager_pb2.UploadCasBlobResp: ...

def add_BlobsManagerServicer_to_server(servicer: BlobsManagerServicer, server: grpc.Server) -> None: ...
