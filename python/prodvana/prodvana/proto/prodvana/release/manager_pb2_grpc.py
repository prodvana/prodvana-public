# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from prodvana.proto.prodvana.release import manager_pb2 as prodvana_dot_release_dot_manager__pb2


class ReleaseManagerStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.RecordRelease = channel.unary_unary(
                '/prodvana.release.ReleaseManager/RecordRelease',
                request_serializer=prodvana_dot_release_dot_manager__pb2.RecordReleaseReq.SerializeToString,
                response_deserializer=prodvana_dot_release_dot_manager__pb2.RecordReleaseResp.FromString,
                )
        self.UpdateRelease = channel.unary_unary(
                '/prodvana.release.ReleaseManager/UpdateRelease',
                request_serializer=prodvana_dot_release_dot_manager__pb2.UpdateReleaseStatusReq.SerializeToString,
                response_deserializer=prodvana_dot_release_dot_manager__pb2.UpdateReleaseStatusResp.FromString,
                )
        self.ListReleases = channel.unary_unary(
                '/prodvana.release.ReleaseManager/ListReleases',
                request_serializer=prodvana_dot_release_dot_manager__pb2.ListReleasesReq.SerializeToString,
                response_deserializer=prodvana_dot_release_dot_manager__pb2.ListReleasesResp.FromString,
                )
        self.ListReleasesStream = channel.unary_stream(
                '/prodvana.release.ReleaseManager/ListReleasesStream',
                request_serializer=prodvana_dot_release_dot_manager__pb2.ListReleasesReq.SerializeToString,
                response_deserializer=prodvana_dot_release_dot_manager__pb2.ListReleasesResp.FromString,
                )


class ReleaseManagerServicer(object):
    """Missing associated documentation comment in .proto file."""

    def RecordRelease(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateRelease(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListReleases(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListReleasesStream(self, request, context):
        """page tokens arguments are ignored here
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ReleaseManagerServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'RecordRelease': grpc.unary_unary_rpc_method_handler(
                    servicer.RecordRelease,
                    request_deserializer=prodvana_dot_release_dot_manager__pb2.RecordReleaseReq.FromString,
                    response_serializer=prodvana_dot_release_dot_manager__pb2.RecordReleaseResp.SerializeToString,
            ),
            'UpdateRelease': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateRelease,
                    request_deserializer=prodvana_dot_release_dot_manager__pb2.UpdateReleaseStatusReq.FromString,
                    response_serializer=prodvana_dot_release_dot_manager__pb2.UpdateReleaseStatusResp.SerializeToString,
            ),
            'ListReleases': grpc.unary_unary_rpc_method_handler(
                    servicer.ListReleases,
                    request_deserializer=prodvana_dot_release_dot_manager__pb2.ListReleasesReq.FromString,
                    response_serializer=prodvana_dot_release_dot_manager__pb2.ListReleasesResp.SerializeToString,
            ),
            'ListReleasesStream': grpc.unary_stream_rpc_method_handler(
                    servicer.ListReleasesStream,
                    request_deserializer=prodvana_dot_release_dot_manager__pb2.ListReleasesReq.FromString,
                    response_serializer=prodvana_dot_release_dot_manager__pb2.ListReleasesResp.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'prodvana.release.ReleaseManager', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ReleaseManager(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def RecordRelease(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.release.ReleaseManager/RecordRelease',
            prodvana_dot_release_dot_manager__pb2.RecordReleaseReq.SerializeToString,
            prodvana_dot_release_dot_manager__pb2.RecordReleaseResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateRelease(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.release.ReleaseManager/UpdateRelease',
            prodvana_dot_release_dot_manager__pb2.UpdateReleaseStatusReq.SerializeToString,
            prodvana_dot_release_dot_manager__pb2.UpdateReleaseStatusResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListReleases(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.release.ReleaseManager/ListReleases',
            prodvana_dot_release_dot_manager__pb2.ListReleasesReq.SerializeToString,
            prodvana_dot_release_dot_manager__pb2.ListReleasesResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListReleasesStream(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/prodvana.release.ReleaseManager/ListReleasesStream',
            prodvana_dot_release_dot_manager__pb2.ListReleasesReq.SerializeToString,
            prodvana_dot_release_dot_manager__pb2.ListReleasesResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
