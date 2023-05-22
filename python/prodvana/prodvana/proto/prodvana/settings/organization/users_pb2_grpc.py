# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from prodvana.proto.prodvana.settings.organization import users_pb2 as prodvana_dot_settings_dot_organization_dot_users__pb2


class UsersSettingsManagerStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetUser = channel.unary_unary(
                '/prodvana.settings.organization.UsersSettingsManager/GetUser',
                request_serializer=prodvana_dot_settings_dot_organization_dot_users__pb2.GetUserReq.SerializeToString,
                response_deserializer=prodvana_dot_settings_dot_organization_dot_users__pb2.GetUserResp.FromString,
                )
        self.ListUsers = channel.unary_unary(
                '/prodvana.settings.organization.UsersSettingsManager/ListUsers',
                request_serializer=prodvana_dot_settings_dot_organization_dot_users__pb2.ListUsersReq.SerializeToString,
                response_deserializer=prodvana_dot_settings_dot_organization_dot_users__pb2.ListUsersResp.FromString,
                )
        self.ListRoles = channel.unary_unary(
                '/prodvana.settings.organization.UsersSettingsManager/ListRoles',
                request_serializer=prodvana_dot_settings_dot_organization_dot_users__pb2.ListRolesReq.SerializeToString,
                response_deserializer=prodvana_dot_settings_dot_organization_dot_users__pb2.ListRolesResp.FromString,
                )
        self.SetRoles = channel.unary_unary(
                '/prodvana.settings.organization.UsersSettingsManager/SetRoles',
                request_serializer=prodvana_dot_settings_dot_organization_dot_users__pb2.SetRolesReq.SerializeToString,
                response_deserializer=prodvana_dot_settings_dot_organization_dot_users__pb2.SetRolesResp.FromString,
                )


class UsersSettingsManagerServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetUser(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListUsers(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListRoles(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetRoles(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_UsersSettingsManagerServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetUser': grpc.unary_unary_rpc_method_handler(
                    servicer.GetUser,
                    request_deserializer=prodvana_dot_settings_dot_organization_dot_users__pb2.GetUserReq.FromString,
                    response_serializer=prodvana_dot_settings_dot_organization_dot_users__pb2.GetUserResp.SerializeToString,
            ),
            'ListUsers': grpc.unary_unary_rpc_method_handler(
                    servicer.ListUsers,
                    request_deserializer=prodvana_dot_settings_dot_organization_dot_users__pb2.ListUsersReq.FromString,
                    response_serializer=prodvana_dot_settings_dot_organization_dot_users__pb2.ListUsersResp.SerializeToString,
            ),
            'ListRoles': grpc.unary_unary_rpc_method_handler(
                    servicer.ListRoles,
                    request_deserializer=prodvana_dot_settings_dot_organization_dot_users__pb2.ListRolesReq.FromString,
                    response_serializer=prodvana_dot_settings_dot_organization_dot_users__pb2.ListRolesResp.SerializeToString,
            ),
            'SetRoles': grpc.unary_unary_rpc_method_handler(
                    servicer.SetRoles,
                    request_deserializer=prodvana_dot_settings_dot_organization_dot_users__pb2.SetRolesReq.FromString,
                    response_serializer=prodvana_dot_settings_dot_organization_dot_users__pb2.SetRolesResp.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'prodvana.settings.organization.UsersSettingsManager', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class UsersSettingsManager(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetUser(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.settings.organization.UsersSettingsManager/GetUser',
            prodvana_dot_settings_dot_organization_dot_users__pb2.GetUserReq.SerializeToString,
            prodvana_dot_settings_dot_organization_dot_users__pb2.GetUserResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListUsers(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.settings.organization.UsersSettingsManager/ListUsers',
            prodvana_dot_settings_dot_organization_dot_users__pb2.ListUsersReq.SerializeToString,
            prodvana_dot_settings_dot_organization_dot_users__pb2.ListUsersResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListRoles(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.settings.organization.UsersSettingsManager/ListRoles',
            prodvana_dot_settings_dot_organization_dot_users__pb2.ListRolesReq.SerializeToString,
            prodvana_dot_settings_dot_organization_dot_users__pb2.ListRolesResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetRoles(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.settings.organization.UsersSettingsManager/SetRoles',
            prodvana_dot_settings_dot_organization_dot_users__pb2.SetRolesReq.SerializeToString,
            prodvana_dot_settings_dot_organization_dot_users__pb2.SetRolesResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)