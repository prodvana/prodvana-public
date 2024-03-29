# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from prodvana.proto.prodvana.application import application_manager_pb2 as prodvana_dot_application_dot_application__manager__pb2


class ApplicationManagerStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.ConfigureApplication = channel.unary_unary(
                '/prodvana.application.ApplicationManager/ConfigureApplication',
                request_serializer=prodvana_dot_application_dot_application__manager__pb2.ConfigureApplicationReq.SerializeToString,
                response_deserializer=prodvana_dot_application_dot_application__manager__pb2.ConfigureApplicationResp.FromString,
                )
        self.ValidateConfigureApplication = channel.unary_unary(
                '/prodvana.application.ApplicationManager/ValidateConfigureApplication',
                request_serializer=prodvana_dot_application_dot_application__manager__pb2.ConfigureApplicationReq.SerializeToString,
                response_deserializer=prodvana_dot_application_dot_application__manager__pb2.ValidateConfigureApplicationResp.FromString,
                )
        self.ListApplications = channel.unary_unary(
                '/prodvana.application.ApplicationManager/ListApplications',
                request_serializer=prodvana_dot_application_dot_application__manager__pb2.ListApplicationsReq.SerializeToString,
                response_deserializer=prodvana_dot_application_dot_application__manager__pb2.ListApplicationsResp.FromString,
                )
        self.GetApplicationConfig = channel.unary_unary(
                '/prodvana.application.ApplicationManager/GetApplicationConfig',
                request_serializer=prodvana_dot_application_dot_application__manager__pb2.GetApplicationConfigReq.SerializeToString,
                response_deserializer=prodvana_dot_application_dot_application__manager__pb2.GetApplicationConfigResp.FromString,
                )
        self.ListApplicationConfigVersions = channel.unary_unary(
                '/prodvana.application.ApplicationManager/ListApplicationConfigVersions',
                request_serializer=prodvana_dot_application_dot_application__manager__pb2.ListApplicationConfigVersionsReq.SerializeToString,
                response_deserializer=prodvana_dot_application_dot_application__manager__pb2.ListApplicationConfigVersionsResp.FromString,
                )
        self.GetApplication = channel.unary_unary(
                '/prodvana.application.ApplicationManager/GetApplication',
                request_serializer=prodvana_dot_application_dot_application__manager__pb2.GetApplicationReq.SerializeToString,
                response_deserializer=prodvana_dot_application_dot_application__manager__pb2.GetApplicationResp.FromString,
                )
        self.DeleteApplication = channel.unary_unary(
                '/prodvana.application.ApplicationManager/DeleteApplication',
                request_serializer=prodvana_dot_application_dot_application__manager__pb2.DeleteApplicationReq.SerializeToString,
                response_deserializer=prodvana_dot_application_dot_application__manager__pb2.DeleteApplicationResp.FromString,
                )
        self.GetApplicationMetrics = channel.unary_unary(
                '/prodvana.application.ApplicationManager/GetApplicationMetrics',
                request_serializer=prodvana_dot_application_dot_application__manager__pb2.GetApplicationMetricsReq.SerializeToString,
                response_deserializer=prodvana_dot_application_dot_application__manager__pb2.GetApplicationMetricsResp.FromString,
                )
        self.GetApplicationInsights = channel.unary_unary(
                '/prodvana.application.ApplicationManager/GetApplicationInsights',
                request_serializer=prodvana_dot_application_dot_application__manager__pb2.GetApplicationInsightsReq.SerializeToString,
                response_deserializer=prodvana_dot_application_dot_application__manager__pb2.GetApplicationInsightsResp.FromString,
                )
        self.SnoozeApplicationInsight = channel.unary_unary(
                '/prodvana.application.ApplicationManager/SnoozeApplicationInsight',
                request_serializer=prodvana_dot_application_dot_application__manager__pb2.SnoozeApplicationInsightReq.SerializeToString,
                response_deserializer=prodvana_dot_application_dot_application__manager__pb2.SnoozeApplicationInsightResp.FromString,
                )
        self.GetApplicationMetadata = channel.unary_unary(
                '/prodvana.application.ApplicationManager/GetApplicationMetadata',
                request_serializer=prodvana_dot_application_dot_application__manager__pb2.GetApplicationMetadataReq.SerializeToString,
                response_deserializer=prodvana_dot_application_dot_application__manager__pb2.GetApplicationMetadataResp.FromString,
                )
        self.SetApplicationMetadata = channel.unary_unary(
                '/prodvana.application.ApplicationManager/SetApplicationMetadata',
                request_serializer=prodvana_dot_application_dot_application__manager__pb2.SetApplicationMetadataReq.SerializeToString,
                response_deserializer=prodvana_dot_application_dot_application__manager__pb2.SetApplicationMetadataResp.FromString,
                )


class ApplicationManagerServicer(object):
    """Missing associated documentation comment in .proto file."""

    def ConfigureApplication(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ValidateConfigureApplication(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListApplications(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetApplicationConfig(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListApplicationConfigVersions(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetApplication(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteApplication(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetApplicationMetrics(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetApplicationInsights(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SnoozeApplicationInsight(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetApplicationMetadata(self, request, context):
        """Get application metadata, useful for constructing edit workflows for metadata
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetApplicationMetadata(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ApplicationManagerServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'ConfigureApplication': grpc.unary_unary_rpc_method_handler(
                    servicer.ConfigureApplication,
                    request_deserializer=prodvana_dot_application_dot_application__manager__pb2.ConfigureApplicationReq.FromString,
                    response_serializer=prodvana_dot_application_dot_application__manager__pb2.ConfigureApplicationResp.SerializeToString,
            ),
            'ValidateConfigureApplication': grpc.unary_unary_rpc_method_handler(
                    servicer.ValidateConfigureApplication,
                    request_deserializer=prodvana_dot_application_dot_application__manager__pb2.ConfigureApplicationReq.FromString,
                    response_serializer=prodvana_dot_application_dot_application__manager__pb2.ValidateConfigureApplicationResp.SerializeToString,
            ),
            'ListApplications': grpc.unary_unary_rpc_method_handler(
                    servicer.ListApplications,
                    request_deserializer=prodvana_dot_application_dot_application__manager__pb2.ListApplicationsReq.FromString,
                    response_serializer=prodvana_dot_application_dot_application__manager__pb2.ListApplicationsResp.SerializeToString,
            ),
            'GetApplicationConfig': grpc.unary_unary_rpc_method_handler(
                    servicer.GetApplicationConfig,
                    request_deserializer=prodvana_dot_application_dot_application__manager__pb2.GetApplicationConfigReq.FromString,
                    response_serializer=prodvana_dot_application_dot_application__manager__pb2.GetApplicationConfigResp.SerializeToString,
            ),
            'ListApplicationConfigVersions': grpc.unary_unary_rpc_method_handler(
                    servicer.ListApplicationConfigVersions,
                    request_deserializer=prodvana_dot_application_dot_application__manager__pb2.ListApplicationConfigVersionsReq.FromString,
                    response_serializer=prodvana_dot_application_dot_application__manager__pb2.ListApplicationConfigVersionsResp.SerializeToString,
            ),
            'GetApplication': grpc.unary_unary_rpc_method_handler(
                    servicer.GetApplication,
                    request_deserializer=prodvana_dot_application_dot_application__manager__pb2.GetApplicationReq.FromString,
                    response_serializer=prodvana_dot_application_dot_application__manager__pb2.GetApplicationResp.SerializeToString,
            ),
            'DeleteApplication': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteApplication,
                    request_deserializer=prodvana_dot_application_dot_application__manager__pb2.DeleteApplicationReq.FromString,
                    response_serializer=prodvana_dot_application_dot_application__manager__pb2.DeleteApplicationResp.SerializeToString,
            ),
            'GetApplicationMetrics': grpc.unary_unary_rpc_method_handler(
                    servicer.GetApplicationMetrics,
                    request_deserializer=prodvana_dot_application_dot_application__manager__pb2.GetApplicationMetricsReq.FromString,
                    response_serializer=prodvana_dot_application_dot_application__manager__pb2.GetApplicationMetricsResp.SerializeToString,
            ),
            'GetApplicationInsights': grpc.unary_unary_rpc_method_handler(
                    servicer.GetApplicationInsights,
                    request_deserializer=prodvana_dot_application_dot_application__manager__pb2.GetApplicationInsightsReq.FromString,
                    response_serializer=prodvana_dot_application_dot_application__manager__pb2.GetApplicationInsightsResp.SerializeToString,
            ),
            'SnoozeApplicationInsight': grpc.unary_unary_rpc_method_handler(
                    servicer.SnoozeApplicationInsight,
                    request_deserializer=prodvana_dot_application_dot_application__manager__pb2.SnoozeApplicationInsightReq.FromString,
                    response_serializer=prodvana_dot_application_dot_application__manager__pb2.SnoozeApplicationInsightResp.SerializeToString,
            ),
            'GetApplicationMetadata': grpc.unary_unary_rpc_method_handler(
                    servicer.GetApplicationMetadata,
                    request_deserializer=prodvana_dot_application_dot_application__manager__pb2.GetApplicationMetadataReq.FromString,
                    response_serializer=prodvana_dot_application_dot_application__manager__pb2.GetApplicationMetadataResp.SerializeToString,
            ),
            'SetApplicationMetadata': grpc.unary_unary_rpc_method_handler(
                    servicer.SetApplicationMetadata,
                    request_deserializer=prodvana_dot_application_dot_application__manager__pb2.SetApplicationMetadataReq.FromString,
                    response_serializer=prodvana_dot_application_dot_application__manager__pb2.SetApplicationMetadataResp.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'prodvana.application.ApplicationManager', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ApplicationManager(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def ConfigureApplication(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.application.ApplicationManager/ConfigureApplication',
            prodvana_dot_application_dot_application__manager__pb2.ConfigureApplicationReq.SerializeToString,
            prodvana_dot_application_dot_application__manager__pb2.ConfigureApplicationResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ValidateConfigureApplication(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.application.ApplicationManager/ValidateConfigureApplication',
            prodvana_dot_application_dot_application__manager__pb2.ConfigureApplicationReq.SerializeToString,
            prodvana_dot_application_dot_application__manager__pb2.ValidateConfigureApplicationResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListApplications(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.application.ApplicationManager/ListApplications',
            prodvana_dot_application_dot_application__manager__pb2.ListApplicationsReq.SerializeToString,
            prodvana_dot_application_dot_application__manager__pb2.ListApplicationsResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetApplicationConfig(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.application.ApplicationManager/GetApplicationConfig',
            prodvana_dot_application_dot_application__manager__pb2.GetApplicationConfigReq.SerializeToString,
            prodvana_dot_application_dot_application__manager__pb2.GetApplicationConfigResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListApplicationConfigVersions(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.application.ApplicationManager/ListApplicationConfigVersions',
            prodvana_dot_application_dot_application__manager__pb2.ListApplicationConfigVersionsReq.SerializeToString,
            prodvana_dot_application_dot_application__manager__pb2.ListApplicationConfigVersionsResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetApplication(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.application.ApplicationManager/GetApplication',
            prodvana_dot_application_dot_application__manager__pb2.GetApplicationReq.SerializeToString,
            prodvana_dot_application_dot_application__manager__pb2.GetApplicationResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteApplication(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.application.ApplicationManager/DeleteApplication',
            prodvana_dot_application_dot_application__manager__pb2.DeleteApplicationReq.SerializeToString,
            prodvana_dot_application_dot_application__manager__pb2.DeleteApplicationResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetApplicationMetrics(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.application.ApplicationManager/GetApplicationMetrics',
            prodvana_dot_application_dot_application__manager__pb2.GetApplicationMetricsReq.SerializeToString,
            prodvana_dot_application_dot_application__manager__pb2.GetApplicationMetricsResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetApplicationInsights(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.application.ApplicationManager/GetApplicationInsights',
            prodvana_dot_application_dot_application__manager__pb2.GetApplicationInsightsReq.SerializeToString,
            prodvana_dot_application_dot_application__manager__pb2.GetApplicationInsightsResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SnoozeApplicationInsight(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.application.ApplicationManager/SnoozeApplicationInsight',
            prodvana_dot_application_dot_application__manager__pb2.SnoozeApplicationInsightReq.SerializeToString,
            prodvana_dot_application_dot_application__manager__pb2.SnoozeApplicationInsightResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetApplicationMetadata(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.application.ApplicationManager/GetApplicationMetadata',
            prodvana_dot_application_dot_application__manager__pb2.GetApplicationMetadataReq.SerializeToString,
            prodvana_dot_application_dot_application__manager__pb2.GetApplicationMetadataResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetApplicationMetadata(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.application.ApplicationManager/SetApplicationMetadata',
            prodvana_dot_application_dot_application__manager__pb2.SetApplicationMetadataReq.SerializeToString,
            prodvana_dot_application_dot_application__manager__pb2.SetApplicationMetadataResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
