# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from prodvana.proto.prodvana.secrets import secrets_manager_pb2 as prodvana_dot_secrets_dot_secrets__manager__pb2


class SecretsManagerStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.ListSecrets = channel.unary_unary(
                '/prodvana.environment.SecretsManager/ListSecrets',
                request_serializer=prodvana_dot_secrets_dot_secrets__manager__pb2.ListSecretsReq.SerializeToString,
                response_deserializer=prodvana_dot_secrets_dot_secrets__manager__pb2.ListSecretsResp.FromString,
                )
        self.ListSecretVersions = channel.unary_unary(
                '/prodvana.environment.SecretsManager/ListSecretVersions',
                request_serializer=prodvana_dot_secrets_dot_secrets__manager__pb2.ListSecretVersionsReq.SerializeToString,
                response_deserializer=prodvana_dot_secrets_dot_secrets__manager__pb2.ListSecretVersionsResp.FromString,
                )
        self.SetSecret = channel.unary_unary(
                '/prodvana.environment.SecretsManager/SetSecret',
                request_serializer=prodvana_dot_secrets_dot_secrets__manager__pb2.SetSecretReq.SerializeToString,
                response_deserializer=prodvana_dot_secrets_dot_secrets__manager__pb2.SetSecretResp.FromString,
                )
        self.DeleteSecret = channel.unary_unary(
                '/prodvana.environment.SecretsManager/DeleteSecret',
                request_serializer=prodvana_dot_secrets_dot_secrets__manager__pb2.DeleteSecretReq.SerializeToString,
                response_deserializer=prodvana_dot_secrets_dot_secrets__manager__pb2.DeleteSecretResp.FromString,
                )
        self.DeleteSecretVersion = channel.unary_unary(
                '/prodvana.environment.SecretsManager/DeleteSecretVersion',
                request_serializer=prodvana_dot_secrets_dot_secrets__manager__pb2.DeleteSecretVersionReq.SerializeToString,
                response_deserializer=prodvana_dot_secrets_dot_secrets__manager__pb2.DeleteSecretVersionResp.FromString,
                )


class SecretsManagerServicer(object):
    """Missing associated documentation comment in .proto file."""

    def ListSecrets(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListSecretVersions(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetSecret(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteSecret(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteSecretVersion(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_SecretsManagerServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'ListSecrets': grpc.unary_unary_rpc_method_handler(
                    servicer.ListSecrets,
                    request_deserializer=prodvana_dot_secrets_dot_secrets__manager__pb2.ListSecretsReq.FromString,
                    response_serializer=prodvana_dot_secrets_dot_secrets__manager__pb2.ListSecretsResp.SerializeToString,
            ),
            'ListSecretVersions': grpc.unary_unary_rpc_method_handler(
                    servicer.ListSecretVersions,
                    request_deserializer=prodvana_dot_secrets_dot_secrets__manager__pb2.ListSecretVersionsReq.FromString,
                    response_serializer=prodvana_dot_secrets_dot_secrets__manager__pb2.ListSecretVersionsResp.SerializeToString,
            ),
            'SetSecret': grpc.unary_unary_rpc_method_handler(
                    servicer.SetSecret,
                    request_deserializer=prodvana_dot_secrets_dot_secrets__manager__pb2.SetSecretReq.FromString,
                    response_serializer=prodvana_dot_secrets_dot_secrets__manager__pb2.SetSecretResp.SerializeToString,
            ),
            'DeleteSecret': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteSecret,
                    request_deserializer=prodvana_dot_secrets_dot_secrets__manager__pb2.DeleteSecretReq.FromString,
                    response_serializer=prodvana_dot_secrets_dot_secrets__manager__pb2.DeleteSecretResp.SerializeToString,
            ),
            'DeleteSecretVersion': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteSecretVersion,
                    request_deserializer=prodvana_dot_secrets_dot_secrets__manager__pb2.DeleteSecretVersionReq.FromString,
                    response_serializer=prodvana_dot_secrets_dot_secrets__manager__pb2.DeleteSecretVersionResp.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'prodvana.environment.SecretsManager', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class SecretsManager(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def ListSecrets(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.environment.SecretsManager/ListSecrets',
            prodvana_dot_secrets_dot_secrets__manager__pb2.ListSecretsReq.SerializeToString,
            prodvana_dot_secrets_dot_secrets__manager__pb2.ListSecretsResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListSecretVersions(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.environment.SecretsManager/ListSecretVersions',
            prodvana_dot_secrets_dot_secrets__manager__pb2.ListSecretVersionsReq.SerializeToString,
            prodvana_dot_secrets_dot_secrets__manager__pb2.ListSecretVersionsResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetSecret(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.environment.SecretsManager/SetSecret',
            prodvana_dot_secrets_dot_secrets__manager__pb2.SetSecretReq.SerializeToString,
            prodvana_dot_secrets_dot_secrets__manager__pb2.SetSecretResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteSecret(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.environment.SecretsManager/DeleteSecret',
            prodvana_dot_secrets_dot_secrets__manager__pb2.DeleteSecretReq.SerializeToString,
            prodvana_dot_secrets_dot_secrets__manager__pb2.DeleteSecretResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteSecretVersion(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.environment.SecretsManager/DeleteSecretVersion',
            prodvana_dot_secrets_dot_secrets__manager__pb2.DeleteSecretVersionReq.SerializeToString,
            prodvana_dot_secrets_dot_secrets__manager__pb2.DeleteSecretVersionResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
