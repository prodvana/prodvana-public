# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from prodvana.proto.prodvana.service import service_manager_pb2 as prodvana_dot_service_dot_service__manager__pb2


class ServiceManagerStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.ConfigureService = channel.unary_unary(
                '/prodvana.service.ServiceManager/ConfigureService',
                request_serializer=prodvana_dot_service_dot_service__manager__pb2.ConfigureServiceReq.SerializeToString,
                response_deserializer=prodvana_dot_service_dot_service__manager__pb2.ConfigureServiceResp.FromString,
                )
        self.ListServiceConfigVersions = channel.unary_unary(
                '/prodvana.service.ServiceManager/ListServiceConfigVersions',
                request_serializer=prodvana_dot_service_dot_service__manager__pb2.ListServiceConfigVersionsReq.SerializeToString,
                response_deserializer=prodvana_dot_service_dot_service__manager__pb2.ListServiceConfigVersionsResp.FromString,
                )
        self.GetServiceConfig = channel.unary_unary(
                '/prodvana.service.ServiceManager/GetServiceConfig',
                request_serializer=prodvana_dot_service_dot_service__manager__pb2.GetServiceConfigReq.SerializeToString,
                response_deserializer=prodvana_dot_service_dot_service__manager__pb2.GetServiceConfigResp.FromString,
                )
        self.ApplyParameters = channel.unary_unary(
                '/prodvana.service.ServiceManager/ApplyParameters',
                request_serializer=prodvana_dot_service_dot_service__manager__pb2.ApplyParametersReq.SerializeToString,
                response_deserializer=prodvana_dot_service_dot_service__manager__pb2.ApplyParametersResp.FromString,
                )
        self.ValidateApplyParameters = channel.unary_unary(
                '/prodvana.service.ServiceManager/ValidateApplyParameters',
                request_serializer=prodvana_dot_service_dot_service__manager__pb2.ApplyParametersReq.SerializeToString,
                response_deserializer=prodvana_dot_service_dot_service__manager__pb2.ValidateApplyParametersResp.FromString,
                )
        self.GetMaterializedConfig = channel.unary_unary(
                '/prodvana.service.ServiceManager/GetMaterializedConfig',
                request_serializer=prodvana_dot_service_dot_service__manager__pb2.GetMaterializedConfigReq.SerializeToString,
                response_deserializer=prodvana_dot_service_dot_service__manager__pb2.GetMaterializedConfigResp.FromString,
                )
        self.ListMaterializedConfigVersions = channel.unary_unary(
                '/prodvana.service.ServiceManager/ListMaterializedConfigVersions',
                request_serializer=prodvana_dot_service_dot_service__manager__pb2.ListMaterializedConfigVersionsReq.SerializeToString,
                response_deserializer=prodvana_dot_service_dot_service__manager__pb2.ListMaterializedConfigVersionsResp.FromString,
                )
        self.DeleteService = channel.unary_unary(
                '/prodvana.service.ServiceManager/DeleteService',
                request_serializer=prodvana_dot_service_dot_service__manager__pb2.DeleteServiceReq.SerializeToString,
                response_deserializer=prodvana_dot_service_dot_service__manager__pb2.DeleteServiceResp.FromString,
                )
        self.ListServices = channel.unary_unary(
                '/prodvana.service.ServiceManager/ListServices',
                request_serializer=prodvana_dot_service_dot_service__manager__pb2.ListServicesReq.SerializeToString,
                response_deserializer=prodvana_dot_service_dot_service__manager__pb2.ListServicesResp.FromString,
                )
        self.ListCommits = channel.unary_unary(
                '/prodvana.service.ServiceManager/ListCommits',
                request_serializer=prodvana_dot_service_dot_service__manager__pb2.ListCommitsReq.SerializeToString,
                response_deserializer=prodvana_dot_service_dot_service__manager__pb2.ListCommitsResp.FromString,
                )
        self.GetService = channel.unary_unary(
                '/prodvana.service.ServiceManager/GetService',
                request_serializer=prodvana_dot_service_dot_service__manager__pb2.GetServiceReq.SerializeToString,
                response_deserializer=prodvana_dot_service_dot_service__manager__pb2.GetServiceResp.FromString,
                )
        self.ListServiceInstances = channel.unary_unary(
                '/prodvana.service.ServiceManager/ListServiceInstances',
                request_serializer=prodvana_dot_service_dot_service__manager__pb2.ListServiceInstancesReq.SerializeToString,
                response_deserializer=prodvana_dot_service_dot_service__manager__pb2.ListServiceInstancesResp.FromString,
                )
        self.GetServiceInstance = channel.unary_unary(
                '/prodvana.service.ServiceManager/GetServiceInstance',
                request_serializer=prodvana_dot_service_dot_service__manager__pb2.GetServiceInstanceReq.SerializeToString,
                response_deserializer=prodvana_dot_service_dot_service__manager__pb2.GetServiceInstanceResp.FromString,
                )
        self.GetServiceMetrics = channel.unary_unary(
                '/prodvana.service.ServiceManager/GetServiceMetrics',
                request_serializer=prodvana_dot_service_dot_service__manager__pb2.GetServiceMetricsReq.SerializeToString,
                response_deserializer=prodvana_dot_service_dot_service__manager__pb2.GetServiceMetricsResp.FromString,
                )
        self.GetServiceInsights = channel.unary_unary(
                '/prodvana.service.ServiceManager/GetServiceInsights',
                request_serializer=prodvana_dot_service_dot_service__manager__pb2.GetServiceInsightsReq.SerializeToString,
                response_deserializer=prodvana_dot_service_dot_service__manager__pb2.GetServiceInsightsResp.FromString,
                )
        self.SnoozeServiceInsight = channel.unary_unary(
                '/prodvana.service.ServiceManager/SnoozeServiceInsight',
                request_serializer=prodvana_dot_service_dot_service__manager__pb2.SnoozeServiceInsightReq.SerializeToString,
                response_deserializer=prodvana_dot_service_dot_service__manager__pb2.SnoozeServiceInsightResp.FromString,
                )
        self.GetServiceMetadata = channel.unary_unary(
                '/prodvana.service.ServiceManager/GetServiceMetadata',
                request_serializer=prodvana_dot_service_dot_service__manager__pb2.GetServiceMetadataReq.SerializeToString,
                response_deserializer=prodvana_dot_service_dot_service__manager__pb2.GetServiceMetadataResp.FromString,
                )
        self.SetServiceMetadata = channel.unary_unary(
                '/prodvana.service.ServiceManager/SetServiceMetadata',
                request_serializer=prodvana_dot_service_dot_service__manager__pb2.SetServiceMetadataReq.SerializeToString,
                response_deserializer=prodvana_dot_service_dot_service__manager__pb2.SetServiceMetadataResp.FromString,
                )


class ServiceManagerServicer(object):
    """Missing associated documentation comment in .proto file."""

    def ConfigureService(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListServiceConfigVersions(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetServiceConfig(self, request, context):
        """unparametrized configs
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ApplyParameters(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ValidateApplyParameters(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetMaterializedConfig(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListMaterializedConfigVersions(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteService(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListServices(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListCommits(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetService(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListServiceInstances(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetServiceInstance(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetServiceMetrics(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetServiceInsights(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SnoozeServiceInsight(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetServiceMetadata(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetServiceMetadata(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ServiceManagerServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'ConfigureService': grpc.unary_unary_rpc_method_handler(
                    servicer.ConfigureService,
                    request_deserializer=prodvana_dot_service_dot_service__manager__pb2.ConfigureServiceReq.FromString,
                    response_serializer=prodvana_dot_service_dot_service__manager__pb2.ConfigureServiceResp.SerializeToString,
            ),
            'ListServiceConfigVersions': grpc.unary_unary_rpc_method_handler(
                    servicer.ListServiceConfigVersions,
                    request_deserializer=prodvana_dot_service_dot_service__manager__pb2.ListServiceConfigVersionsReq.FromString,
                    response_serializer=prodvana_dot_service_dot_service__manager__pb2.ListServiceConfigVersionsResp.SerializeToString,
            ),
            'GetServiceConfig': grpc.unary_unary_rpc_method_handler(
                    servicer.GetServiceConfig,
                    request_deserializer=prodvana_dot_service_dot_service__manager__pb2.GetServiceConfigReq.FromString,
                    response_serializer=prodvana_dot_service_dot_service__manager__pb2.GetServiceConfigResp.SerializeToString,
            ),
            'ApplyParameters': grpc.unary_unary_rpc_method_handler(
                    servicer.ApplyParameters,
                    request_deserializer=prodvana_dot_service_dot_service__manager__pb2.ApplyParametersReq.FromString,
                    response_serializer=prodvana_dot_service_dot_service__manager__pb2.ApplyParametersResp.SerializeToString,
            ),
            'ValidateApplyParameters': grpc.unary_unary_rpc_method_handler(
                    servicer.ValidateApplyParameters,
                    request_deserializer=prodvana_dot_service_dot_service__manager__pb2.ApplyParametersReq.FromString,
                    response_serializer=prodvana_dot_service_dot_service__manager__pb2.ValidateApplyParametersResp.SerializeToString,
            ),
            'GetMaterializedConfig': grpc.unary_unary_rpc_method_handler(
                    servicer.GetMaterializedConfig,
                    request_deserializer=prodvana_dot_service_dot_service__manager__pb2.GetMaterializedConfigReq.FromString,
                    response_serializer=prodvana_dot_service_dot_service__manager__pb2.GetMaterializedConfigResp.SerializeToString,
            ),
            'ListMaterializedConfigVersions': grpc.unary_unary_rpc_method_handler(
                    servicer.ListMaterializedConfigVersions,
                    request_deserializer=prodvana_dot_service_dot_service__manager__pb2.ListMaterializedConfigVersionsReq.FromString,
                    response_serializer=prodvana_dot_service_dot_service__manager__pb2.ListMaterializedConfigVersionsResp.SerializeToString,
            ),
            'DeleteService': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteService,
                    request_deserializer=prodvana_dot_service_dot_service__manager__pb2.DeleteServiceReq.FromString,
                    response_serializer=prodvana_dot_service_dot_service__manager__pb2.DeleteServiceResp.SerializeToString,
            ),
            'ListServices': grpc.unary_unary_rpc_method_handler(
                    servicer.ListServices,
                    request_deserializer=prodvana_dot_service_dot_service__manager__pb2.ListServicesReq.FromString,
                    response_serializer=prodvana_dot_service_dot_service__manager__pb2.ListServicesResp.SerializeToString,
            ),
            'ListCommits': grpc.unary_unary_rpc_method_handler(
                    servicer.ListCommits,
                    request_deserializer=prodvana_dot_service_dot_service__manager__pb2.ListCommitsReq.FromString,
                    response_serializer=prodvana_dot_service_dot_service__manager__pb2.ListCommitsResp.SerializeToString,
            ),
            'GetService': grpc.unary_unary_rpc_method_handler(
                    servicer.GetService,
                    request_deserializer=prodvana_dot_service_dot_service__manager__pb2.GetServiceReq.FromString,
                    response_serializer=prodvana_dot_service_dot_service__manager__pb2.GetServiceResp.SerializeToString,
            ),
            'ListServiceInstances': grpc.unary_unary_rpc_method_handler(
                    servicer.ListServiceInstances,
                    request_deserializer=prodvana_dot_service_dot_service__manager__pb2.ListServiceInstancesReq.FromString,
                    response_serializer=prodvana_dot_service_dot_service__manager__pb2.ListServiceInstancesResp.SerializeToString,
            ),
            'GetServiceInstance': grpc.unary_unary_rpc_method_handler(
                    servicer.GetServiceInstance,
                    request_deserializer=prodvana_dot_service_dot_service__manager__pb2.GetServiceInstanceReq.FromString,
                    response_serializer=prodvana_dot_service_dot_service__manager__pb2.GetServiceInstanceResp.SerializeToString,
            ),
            'GetServiceMetrics': grpc.unary_unary_rpc_method_handler(
                    servicer.GetServiceMetrics,
                    request_deserializer=prodvana_dot_service_dot_service__manager__pb2.GetServiceMetricsReq.FromString,
                    response_serializer=prodvana_dot_service_dot_service__manager__pb2.GetServiceMetricsResp.SerializeToString,
            ),
            'GetServiceInsights': grpc.unary_unary_rpc_method_handler(
                    servicer.GetServiceInsights,
                    request_deserializer=prodvana_dot_service_dot_service__manager__pb2.GetServiceInsightsReq.FromString,
                    response_serializer=prodvana_dot_service_dot_service__manager__pb2.GetServiceInsightsResp.SerializeToString,
            ),
            'SnoozeServiceInsight': grpc.unary_unary_rpc_method_handler(
                    servicer.SnoozeServiceInsight,
                    request_deserializer=prodvana_dot_service_dot_service__manager__pb2.SnoozeServiceInsightReq.FromString,
                    response_serializer=prodvana_dot_service_dot_service__manager__pb2.SnoozeServiceInsightResp.SerializeToString,
            ),
            'GetServiceMetadata': grpc.unary_unary_rpc_method_handler(
                    servicer.GetServiceMetadata,
                    request_deserializer=prodvana_dot_service_dot_service__manager__pb2.GetServiceMetadataReq.FromString,
                    response_serializer=prodvana_dot_service_dot_service__manager__pb2.GetServiceMetadataResp.SerializeToString,
            ),
            'SetServiceMetadata': grpc.unary_unary_rpc_method_handler(
                    servicer.SetServiceMetadata,
                    request_deserializer=prodvana_dot_service_dot_service__manager__pb2.SetServiceMetadataReq.FromString,
                    response_serializer=prodvana_dot_service_dot_service__manager__pb2.SetServiceMetadataResp.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'prodvana.service.ServiceManager', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ServiceManager(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def ConfigureService(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.service.ServiceManager/ConfigureService',
            prodvana_dot_service_dot_service__manager__pb2.ConfigureServiceReq.SerializeToString,
            prodvana_dot_service_dot_service__manager__pb2.ConfigureServiceResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListServiceConfigVersions(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.service.ServiceManager/ListServiceConfigVersions',
            prodvana_dot_service_dot_service__manager__pb2.ListServiceConfigVersionsReq.SerializeToString,
            prodvana_dot_service_dot_service__manager__pb2.ListServiceConfigVersionsResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetServiceConfig(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.service.ServiceManager/GetServiceConfig',
            prodvana_dot_service_dot_service__manager__pb2.GetServiceConfigReq.SerializeToString,
            prodvana_dot_service_dot_service__manager__pb2.GetServiceConfigResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ApplyParameters(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.service.ServiceManager/ApplyParameters',
            prodvana_dot_service_dot_service__manager__pb2.ApplyParametersReq.SerializeToString,
            prodvana_dot_service_dot_service__manager__pb2.ApplyParametersResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ValidateApplyParameters(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.service.ServiceManager/ValidateApplyParameters',
            prodvana_dot_service_dot_service__manager__pb2.ApplyParametersReq.SerializeToString,
            prodvana_dot_service_dot_service__manager__pb2.ValidateApplyParametersResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetMaterializedConfig(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.service.ServiceManager/GetMaterializedConfig',
            prodvana_dot_service_dot_service__manager__pb2.GetMaterializedConfigReq.SerializeToString,
            prodvana_dot_service_dot_service__manager__pb2.GetMaterializedConfigResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListMaterializedConfigVersions(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.service.ServiceManager/ListMaterializedConfigVersions',
            prodvana_dot_service_dot_service__manager__pb2.ListMaterializedConfigVersionsReq.SerializeToString,
            prodvana_dot_service_dot_service__manager__pb2.ListMaterializedConfigVersionsResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteService(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.service.ServiceManager/DeleteService',
            prodvana_dot_service_dot_service__manager__pb2.DeleteServiceReq.SerializeToString,
            prodvana_dot_service_dot_service__manager__pb2.DeleteServiceResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListServices(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.service.ServiceManager/ListServices',
            prodvana_dot_service_dot_service__manager__pb2.ListServicesReq.SerializeToString,
            prodvana_dot_service_dot_service__manager__pb2.ListServicesResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListCommits(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.service.ServiceManager/ListCommits',
            prodvana_dot_service_dot_service__manager__pb2.ListCommitsReq.SerializeToString,
            prodvana_dot_service_dot_service__manager__pb2.ListCommitsResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetService(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.service.ServiceManager/GetService',
            prodvana_dot_service_dot_service__manager__pb2.GetServiceReq.SerializeToString,
            prodvana_dot_service_dot_service__manager__pb2.GetServiceResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListServiceInstances(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.service.ServiceManager/ListServiceInstances',
            prodvana_dot_service_dot_service__manager__pb2.ListServiceInstancesReq.SerializeToString,
            prodvana_dot_service_dot_service__manager__pb2.ListServiceInstancesResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetServiceInstance(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.service.ServiceManager/GetServiceInstance',
            prodvana_dot_service_dot_service__manager__pb2.GetServiceInstanceReq.SerializeToString,
            prodvana_dot_service_dot_service__manager__pb2.GetServiceInstanceResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetServiceMetrics(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.service.ServiceManager/GetServiceMetrics',
            prodvana_dot_service_dot_service__manager__pb2.GetServiceMetricsReq.SerializeToString,
            prodvana_dot_service_dot_service__manager__pb2.GetServiceMetricsResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetServiceInsights(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.service.ServiceManager/GetServiceInsights',
            prodvana_dot_service_dot_service__manager__pb2.GetServiceInsightsReq.SerializeToString,
            prodvana_dot_service_dot_service__manager__pb2.GetServiceInsightsResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SnoozeServiceInsight(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.service.ServiceManager/SnoozeServiceInsight',
            prodvana_dot_service_dot_service__manager__pb2.SnoozeServiceInsightReq.SerializeToString,
            prodvana_dot_service_dot_service__manager__pb2.SnoozeServiceInsightResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetServiceMetadata(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.service.ServiceManager/GetServiceMetadata',
            prodvana_dot_service_dot_service__manager__pb2.GetServiceMetadataReq.SerializeToString,
            prodvana_dot_service_dot_service__manager__pb2.GetServiceMetadataResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetServiceMetadata(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/prodvana.service.ServiceManager/SetServiceMetadata',
            prodvana_dot_service_dot_service__manager__pb2.SetServiceMetadataReq.SerializeToString,
            prodvana_dot_service_dot_service__manager__pb2.SetServiceMetadataResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)