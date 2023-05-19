"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import abc
import grpc
import prodvana.proto.prodvana.release_channel.release_channel_manager_pb2

class ReleaseChannelManagerStub:
    def __init__(self, channel: grpc.Channel) -> None: ...
    ConfigureReleaseChannel: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.release_channel.release_channel_manager_pb2.ConfigureReleaseChannelReq,
        prodvana.proto.prodvana.release_channel.release_channel_manager_pb2.ConfigureReleaseChannelResp,
    ]
    ListReleaseChannels: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.release_channel.release_channel_manager_pb2.ListReleaseChannelsReq,
        prodvana.proto.prodvana.release_channel.release_channel_manager_pb2.ListReleaseChannelsResp,
    ]
    DeleteReleaseChannel: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.release_channel.release_channel_manager_pb2.DeleteReleaseChannelReq,
        prodvana.proto.prodvana.release_channel.release_channel_manager_pb2.DeleteReleaseChannelResp,
    ]
    ListReleaseChannelsV2: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.release_channel.release_channel_manager_pb2.ListReleaseChannelsReq,
        prodvana.proto.prodvana.release_channel.release_channel_manager_pb2.ListReleaseChannelsResp,
    ]
    """identical to ListReleaseChannels, kept for backwards compatibility"""
    GetReleaseChannel: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.release_channel.release_channel_manager_pb2.GetReleaseChannelReq,
        prodvana.proto.prodvana.release_channel.release_channel_manager_pb2.GetReleaseChannelResp,
    ]
    GetReleaseChannelEvents: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.release_channel.release_channel_manager_pb2.GetReleaseChannelEventsReq,
        prodvana.proto.prodvana.release_channel.release_channel_manager_pb2.GetReleaseChannelEventsResp,
    ]

class ReleaseChannelManagerServicer(metaclass=abc.ABCMeta):
    @abc.abstractmethod
    def ConfigureReleaseChannel(
        self,
        request: prodvana.proto.prodvana.release_channel.release_channel_manager_pb2.ConfigureReleaseChannelReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.release_channel.release_channel_manager_pb2.ConfigureReleaseChannelResp: ...
    @abc.abstractmethod
    def ListReleaseChannels(
        self,
        request: prodvana.proto.prodvana.release_channel.release_channel_manager_pb2.ListReleaseChannelsReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.release_channel.release_channel_manager_pb2.ListReleaseChannelsResp: ...
    @abc.abstractmethod
    def DeleteReleaseChannel(
        self,
        request: prodvana.proto.prodvana.release_channel.release_channel_manager_pb2.DeleteReleaseChannelReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.release_channel.release_channel_manager_pb2.DeleteReleaseChannelResp: ...
    @abc.abstractmethod
    def ListReleaseChannelsV2(
        self,
        request: prodvana.proto.prodvana.release_channel.release_channel_manager_pb2.ListReleaseChannelsReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.release_channel.release_channel_manager_pb2.ListReleaseChannelsResp:
        """identical to ListReleaseChannels, kept for backwards compatibility"""
    @abc.abstractmethod
    def GetReleaseChannel(
        self,
        request: prodvana.proto.prodvana.release_channel.release_channel_manager_pb2.GetReleaseChannelReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.release_channel.release_channel_manager_pb2.GetReleaseChannelResp: ...
    @abc.abstractmethod
    def GetReleaseChannelEvents(
        self,
        request: prodvana.proto.prodvana.release_channel.release_channel_manager_pb2.GetReleaseChannelEventsReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.release_channel.release_channel_manager_pb2.GetReleaseChannelEventsResp: ...

def add_ReleaseChannelManagerServicer_to_server(servicer: ReleaseChannelManagerServicer, server: grpc.Server) -> None: ...
