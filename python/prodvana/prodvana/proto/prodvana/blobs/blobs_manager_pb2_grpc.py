# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from prodvana.proto.prodvana.blobs import blobs_manager_pb2 as prodvana_dot_blobs_dot_blobs__manager__pb2


class BlobsManagerStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetCasBlob = channel.unary_stream(
                '/prodvana.blobs.BlobsManager/GetCasBlob',
                request_serializer=prodvana_dot_blobs_dot_blobs__manager__pb2.GetCasBlobReq.SerializeToString,
                response_deserializer=prodvana_dot_blobs_dot_blobs__manager__pb2.GetCasBlobResp.FromString,
                )


class BlobsManagerServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetCasBlob(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_BlobsManagerServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetCasBlob': grpc.unary_stream_rpc_method_handler(
                    servicer.GetCasBlob,
                    request_deserializer=prodvana_dot_blobs_dot_blobs__manager__pb2.GetCasBlobReq.FromString,
                    response_serializer=prodvana_dot_blobs_dot_blobs__manager__pb2.GetCasBlobResp.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'prodvana.blobs.BlobsManager', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class BlobsManager(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetCasBlob(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/prodvana.blobs.BlobsManager/GetCasBlob',
            prodvana_dot_blobs_dot_blobs__manager__pb2.GetCasBlobReq.SerializeToString,
            prodvana_dot_blobs_dot_blobs__manager__pb2.GetCasBlobResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)