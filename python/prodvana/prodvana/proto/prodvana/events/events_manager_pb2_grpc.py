# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from prodvana.proto.prodvana.events import events_manager_pb2 as prodvana_dot_events_dot_events__manager__pb2


class EventsManagerStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetEvents = channel.unary_unary(
                '/prodvana.events.EventsManager/GetEvents',
                request_serializer=prodvana_dot_events_dot_events__manager__pb2.GetEventsReq.SerializeToString,
                response_deserializer=prodvana_dot_events_dot_events__manager__pb2.GetEventsResp.FromString,
                )


class EventsManagerServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetEvents(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_EventsManagerServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetEvents': grpc.unary_unary_rpc_method_handler(
                    servicer.GetEvents,
                    request_deserializer=prodvana_dot_events_dot_events__manager__pb2.GetEventsReq.FromString,
                    response_serializer=prodvana_dot_events_dot_events__manager__pb2.GetEventsResp.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'prodvana.events.EventsManager', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class EventsManager(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetEvents(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.events.EventsManager/GetEvents',
            prodvana_dot_events_dot_events__manager__pb2.GetEventsReq.SerializeToString,
            prodvana_dot_events_dot_events__manager__pb2.GetEventsResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
