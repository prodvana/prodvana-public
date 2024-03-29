"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import abc
import grpc
import prodvana.proto.prodvana.delivery_extension.manager_pb2

class DeliveryExtensionManagerStub:
    def __init__(self, channel: grpc.Channel) -> None: ...
    ConfigureDeliveryExtension: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.delivery_extension.manager_pb2.ConfigureDeliveryExtensionReq,
        prodvana.proto.prodvana.delivery_extension.manager_pb2.ConfigureDeliveryExtensionResp,
    ]
    ValidateConfigureDeliveryExtension: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.delivery_extension.manager_pb2.ConfigureDeliveryExtensionReq,
        prodvana.proto.prodvana.delivery_extension.manager_pb2.ValidateConfigureDeliveryExtensionResp,
    ]
    ListDeliveryExtensions: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.delivery_extension.manager_pb2.ListDeliveryExtensionsReq,
        prodvana.proto.prodvana.delivery_extension.manager_pb2.ListDeliveryExtensionsResp,
    ]
    GetDeliveryExtension: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.delivery_extension.manager_pb2.GetDeliveryExtensionReq,
        prodvana.proto.prodvana.delivery_extension.manager_pb2.GetDeliveryExtensionResp,
    ]
    GetDeliveryExtensionConfig: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.delivery_extension.manager_pb2.GetDeliveryExtensionConfigReq,
        prodvana.proto.prodvana.delivery_extension.manager_pb2.GetDeliveryExtensionConfigResp,
    ]
    GetDeliveryExtensionInstanceConfig: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.delivery_extension.manager_pb2.GetDeliveryExtensionInstanceConfigReq,
        prodvana.proto.prodvana.delivery_extension.manager_pb2.GetDeliveryExtensionInstanceConfigResp,
    ]

class DeliveryExtensionManagerServicer(metaclass=abc.ABCMeta):
    @abc.abstractmethod
    def ConfigureDeliveryExtension(
        self,
        request: prodvana.proto.prodvana.delivery_extension.manager_pb2.ConfigureDeliveryExtensionReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.delivery_extension.manager_pb2.ConfigureDeliveryExtensionResp: ...
    @abc.abstractmethod
    def ValidateConfigureDeliveryExtension(
        self,
        request: prodvana.proto.prodvana.delivery_extension.manager_pb2.ConfigureDeliveryExtensionReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.delivery_extension.manager_pb2.ValidateConfigureDeliveryExtensionResp: ...
    @abc.abstractmethod
    def ListDeliveryExtensions(
        self,
        request: prodvana.proto.prodvana.delivery_extension.manager_pb2.ListDeliveryExtensionsReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.delivery_extension.manager_pb2.ListDeliveryExtensionsResp: ...
    @abc.abstractmethod
    def GetDeliveryExtension(
        self,
        request: prodvana.proto.prodvana.delivery_extension.manager_pb2.GetDeliveryExtensionReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.delivery_extension.manager_pb2.GetDeliveryExtensionResp: ...
    @abc.abstractmethod
    def GetDeliveryExtensionConfig(
        self,
        request: prodvana.proto.prodvana.delivery_extension.manager_pb2.GetDeliveryExtensionConfigReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.delivery_extension.manager_pb2.GetDeliveryExtensionConfigResp: ...
    @abc.abstractmethod
    def GetDeliveryExtensionInstanceConfig(
        self,
        request: prodvana.proto.prodvana.delivery_extension.manager_pb2.GetDeliveryExtensionInstanceConfigReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.delivery_extension.manager_pb2.GetDeliveryExtensionInstanceConfigResp: ...

def add_DeliveryExtensionManagerServicer_to_server(servicer: DeliveryExtensionManagerServicer, server: grpc.Server) -> None: ...
