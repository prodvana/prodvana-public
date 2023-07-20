"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import abc
import grpc
import prodvana.proto.prodvana.desired_state.manager_pb2

class DesiredStateManagerStub:
    def __init__(self, channel: grpc.Channel) -> None: ...
    SetDesiredState: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.desired_state.manager_pb2.SetDesiredStateReq,
        prodvana.proto.prodvana.desired_state.manager_pb2.SetDesiredStateResp,
    ]
    GetServiceDesiredStateConvergenceSummary: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.desired_state.manager_pb2.GetServiceDesiredStateConvergenceSummaryReq,
        prodvana.proto.prodvana.desired_state.manager_pb2.GetServiceDesiredStateConvergenceSummaryResp,
    ]
    GetServiceLastConvergedStates: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.desired_state.manager_pb2.GetServiceLastConvergedStateReq,
        prodvana.proto.prodvana.desired_state.manager_pb2.GetServiceLastConvergedStateResp,
    ]
    GetServiceDesiredStateHistory: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.desired_state.manager_pb2.GetServiceDesiredStateHistoryReq,
        prodvana.proto.prodvana.desired_state.manager_pb2.GetServiceDesiredStateHistoryResp,
    ]
    GetDesiredState: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.desired_state.manager_pb2.GetDesiredStateReq,
        prodvana.proto.prodvana.desired_state.manager_pb2.GetDesiredStateResp,
    ]
    GetDesiredStateConvergenceSummary: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.desired_state.manager_pb2.GetDesiredStateConvergenceReq,
        prodvana.proto.prodvana.desired_state.manager_pb2.GetDesiredStateConvergenceSummaryResp,
    ]
    ValidateDesiredState: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.desired_state.manager_pb2.ValidateDesiredStateReq,
        prodvana.proto.prodvana.desired_state.manager_pb2.ValidateDesiredStateResp,
    ]
    SetManualApproval: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.desired_state.manager_pb2.SetManualApprovalReq,
        prodvana.proto.prodvana.desired_state.manager_pb2.SetManualApprovalResp,
    ]
    PromoteDelivery: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.desired_state.manager_pb2.PromoteDeliveryReq,
        prodvana.proto.prodvana.desired_state.manager_pb2.PromoteDeliveryResp,
    ]
    BypassProtection: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.desired_state.manager_pb2.BypassProtectionReq,
        prodvana.proto.prodvana.desired_state.manager_pb2.BypassProtectionResp,
    ]
    ApproveRuntimeExtensionApply: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.desired_state.manager_pb2.ApproveRuntimeExtensionApplyReq,
        prodvana.proto.prodvana.desired_state.manager_pb2.ApproveRuntimeExtensionApplyResp,
    ]

class DesiredStateManagerServicer(metaclass=abc.ABCMeta):
    @abc.abstractmethod
    def SetDesiredState(
        self,
        request: prodvana.proto.prodvana.desired_state.manager_pb2.SetDesiredStateReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.desired_state.manager_pb2.SetDesiredStateResp: ...
    @abc.abstractmethod
    def GetServiceDesiredStateConvergenceSummary(
        self,
        request: prodvana.proto.prodvana.desired_state.manager_pb2.GetServiceDesiredStateConvergenceSummaryReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.desired_state.manager_pb2.GetServiceDesiredStateConvergenceSummaryResp: ...
    @abc.abstractmethod
    def GetServiceLastConvergedStates(
        self,
        request: prodvana.proto.prodvana.desired_state.manager_pb2.GetServiceLastConvergedStateReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.desired_state.manager_pb2.GetServiceLastConvergedStateResp: ...
    @abc.abstractmethod
    def GetServiceDesiredStateHistory(
        self,
        request: prodvana.proto.prodvana.desired_state.manager_pb2.GetServiceDesiredStateHistoryReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.desired_state.manager_pb2.GetServiceDesiredStateHistoryResp: ...
    @abc.abstractmethod
    def GetDesiredState(
        self,
        request: prodvana.proto.prodvana.desired_state.manager_pb2.GetDesiredStateReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.desired_state.manager_pb2.GetDesiredStateResp: ...
    @abc.abstractmethod
    def GetDesiredStateConvergenceSummary(
        self,
        request: prodvana.proto.prodvana.desired_state.manager_pb2.GetDesiredStateConvergenceReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.desired_state.manager_pb2.GetDesiredStateConvergenceSummaryResp: ...
    @abc.abstractmethod
    def ValidateDesiredState(
        self,
        request: prodvana.proto.prodvana.desired_state.manager_pb2.ValidateDesiredStateReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.desired_state.manager_pb2.ValidateDesiredStateResp: ...
    @abc.abstractmethod
    def SetManualApproval(
        self,
        request: prodvana.proto.prodvana.desired_state.manager_pb2.SetManualApprovalReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.desired_state.manager_pb2.SetManualApprovalResp: ...
    @abc.abstractmethod
    def PromoteDelivery(
        self,
        request: prodvana.proto.prodvana.desired_state.manager_pb2.PromoteDeliveryReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.desired_state.manager_pb2.PromoteDeliveryResp: ...
    @abc.abstractmethod
    def BypassProtection(
        self,
        request: prodvana.proto.prodvana.desired_state.manager_pb2.BypassProtectionReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.desired_state.manager_pb2.BypassProtectionResp: ...
    @abc.abstractmethod
    def ApproveRuntimeExtensionApply(
        self,
        request: prodvana.proto.prodvana.desired_state.manager_pb2.ApproveRuntimeExtensionApplyReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.desired_state.manager_pb2.ApproveRuntimeExtensionApplyResp: ...

def add_DesiredStateManagerServicer_to_server(servicer: DesiredStateManagerServicer, server: grpc.Server) -> None: ...
