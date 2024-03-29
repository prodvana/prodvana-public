# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from prodvana.proto.prodvana.protection import protection_manager_pb2 as prodvana_dot_protection_dot_protection__manager__pb2


class ProtectionManagerStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.ConfigureProtection = channel.unary_unary(
                '/prodvana.protection.ProtectionManager/ConfigureProtection',
                request_serializer=prodvana_dot_protection_dot_protection__manager__pb2.ConfigureProtectionReq.SerializeToString,
                response_deserializer=prodvana_dot_protection_dot_protection__manager__pb2.ConfigureProtectionResp.FromString,
                )
        self.ValidateConfigureProtection = channel.unary_unary(
                '/prodvana.protection.ProtectionManager/ValidateConfigureProtection',
                request_serializer=prodvana_dot_protection_dot_protection__manager__pb2.ConfigureProtectionReq.SerializeToString,
                response_deserializer=prodvana_dot_protection_dot_protection__manager__pb2.ValidateConfigureProtectionResp.FromString,
                )
        self.ListProtections = channel.unary_unary(
                '/prodvana.protection.ProtectionManager/ListProtections',
                request_serializer=prodvana_dot_protection_dot_protection__manager__pb2.ListProtectionsReq.SerializeToString,
                response_deserializer=prodvana_dot_protection_dot_protection__manager__pb2.ListProtectionsResp.FromString,
                )
        self.GetProtection = channel.unary_unary(
                '/prodvana.protection.ProtectionManager/GetProtection',
                request_serializer=prodvana_dot_protection_dot_protection__manager__pb2.GetProtectionReq.SerializeToString,
                response_deserializer=prodvana_dot_protection_dot_protection__manager__pb2.GetProtectionResp.FromString,
                )
        self.GetProtectionConfig = channel.unary_unary(
                '/prodvana.protection.ProtectionManager/GetProtectionConfig',
                request_serializer=prodvana_dot_protection_dot_protection__manager__pb2.GetProtectionConfigReq.SerializeToString,
                response_deserializer=prodvana_dot_protection_dot_protection__manager__pb2.GetProtectionConfigResp.FromString,
                )
        self.GetProtectionAttachmentConfig = channel.unary_unary(
                '/prodvana.protection.ProtectionManager/GetProtectionAttachmentConfig',
                request_serializer=prodvana_dot_protection_dot_protection__manager__pb2.GetProtectionAttachmentConfigReq.SerializeToString,
                response_deserializer=prodvana_dot_protection_dot_protection__manager__pb2.GetProtectionAttachmentConfigResp.FromString,
                )


class ProtectionManagerServicer(object):
    """Missing associated documentation comment in .proto file."""

    def ConfigureProtection(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ValidateConfigureProtection(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListProtections(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetProtection(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetProtectionConfig(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetProtectionAttachmentConfig(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ProtectionManagerServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'ConfigureProtection': grpc.unary_unary_rpc_method_handler(
                    servicer.ConfigureProtection,
                    request_deserializer=prodvana_dot_protection_dot_protection__manager__pb2.ConfigureProtectionReq.FromString,
                    response_serializer=prodvana_dot_protection_dot_protection__manager__pb2.ConfigureProtectionResp.SerializeToString,
            ),
            'ValidateConfigureProtection': grpc.unary_unary_rpc_method_handler(
                    servicer.ValidateConfigureProtection,
                    request_deserializer=prodvana_dot_protection_dot_protection__manager__pb2.ConfigureProtectionReq.FromString,
                    response_serializer=prodvana_dot_protection_dot_protection__manager__pb2.ValidateConfigureProtectionResp.SerializeToString,
            ),
            'ListProtections': grpc.unary_unary_rpc_method_handler(
                    servicer.ListProtections,
                    request_deserializer=prodvana_dot_protection_dot_protection__manager__pb2.ListProtectionsReq.FromString,
                    response_serializer=prodvana_dot_protection_dot_protection__manager__pb2.ListProtectionsResp.SerializeToString,
            ),
            'GetProtection': grpc.unary_unary_rpc_method_handler(
                    servicer.GetProtection,
                    request_deserializer=prodvana_dot_protection_dot_protection__manager__pb2.GetProtectionReq.FromString,
                    response_serializer=prodvana_dot_protection_dot_protection__manager__pb2.GetProtectionResp.SerializeToString,
            ),
            'GetProtectionConfig': grpc.unary_unary_rpc_method_handler(
                    servicer.GetProtectionConfig,
                    request_deserializer=prodvana_dot_protection_dot_protection__manager__pb2.GetProtectionConfigReq.FromString,
                    response_serializer=prodvana_dot_protection_dot_protection__manager__pb2.GetProtectionConfigResp.SerializeToString,
            ),
            'GetProtectionAttachmentConfig': grpc.unary_unary_rpc_method_handler(
                    servicer.GetProtectionAttachmentConfig,
                    request_deserializer=prodvana_dot_protection_dot_protection__manager__pb2.GetProtectionAttachmentConfigReq.FromString,
                    response_serializer=prodvana_dot_protection_dot_protection__manager__pb2.GetProtectionAttachmentConfigResp.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'prodvana.protection.ProtectionManager', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ProtectionManager(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def ConfigureProtection(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.protection.ProtectionManager/ConfigureProtection',
            prodvana_dot_protection_dot_protection__manager__pb2.ConfigureProtectionReq.SerializeToString,
            prodvana_dot_protection_dot_protection__manager__pb2.ConfigureProtectionResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ValidateConfigureProtection(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.protection.ProtectionManager/ValidateConfigureProtection',
            prodvana_dot_protection_dot_protection__manager__pb2.ConfigureProtectionReq.SerializeToString,
            prodvana_dot_protection_dot_protection__manager__pb2.ValidateConfigureProtectionResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListProtections(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.protection.ProtectionManager/ListProtections',
            prodvana_dot_protection_dot_protection__manager__pb2.ListProtectionsReq.SerializeToString,
            prodvana_dot_protection_dot_protection__manager__pb2.ListProtectionsResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetProtection(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.protection.ProtectionManager/GetProtection',
            prodvana_dot_protection_dot_protection__manager__pb2.GetProtectionReq.SerializeToString,
            prodvana_dot_protection_dot_protection__manager__pb2.GetProtectionResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetProtectionConfig(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.protection.ProtectionManager/GetProtectionConfig',
            prodvana_dot_protection_dot_protection__manager__pb2.GetProtectionConfigReq.SerializeToString,
            prodvana_dot_protection_dot_protection__manager__pb2.GetProtectionConfigResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetProtectionAttachmentConfig(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.protection.ProtectionManager/GetProtectionAttachmentConfig',
            prodvana_dot_protection_dot_protection__manager__pb2.GetProtectionAttachmentConfigReq.SerializeToString,
            prodvana_dot_protection_dot_protection__manager__pb2.GetProtectionAttachmentConfigResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
