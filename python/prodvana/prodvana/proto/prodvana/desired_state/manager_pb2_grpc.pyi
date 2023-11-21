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
    PreviewEntityGraph: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.desired_state.manager_pb2.SetDesiredStateReq,
        prodvana.proto.prodvana.desired_state.manager_pb2.PreviewEntityGraphResp,
    ]
    """validate a SetDesiredState call and return a preview entity graph
    TODO(naphat) delete ValidateDesiredState and replace with this
    """
    GetServiceDesiredStateConvergenceSummary: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.desired_state.manager_pb2.GetServiceDesiredStateConvergenceSummaryReq,
        prodvana.proto.prodvana.desired_state.manager_pb2.GetServiceDesiredStateConvergenceSummaryResp,
    ]
    GetServiceLatestDesiredStateId: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.desired_state.manager_pb2.GetServiceLatestDesiredStateIdReq,
        prodvana.proto.prodvana.desired_state.manager_pb2.GetServiceLatestDesiredStateIdResp,
    ]
    GetServiceDesiredStateIdHistory: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.desired_state.manager_pb2.GetServiceDesiredStateIdHistoryReq,
        prodvana.proto.prodvana.desired_state.manager_pb2.GetServiceDesiredStateIdHistoryResp,
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
    GetDesiredStateGraph: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.desired_state.manager_pb2.GetDesiredStateGraphReq,
        prodvana.proto.prodvana.desired_state.manager_pb2.GetDesiredStateGraphResp,
    ]
    GetDesiredStateConvergenceSummary: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.desired_state.manager_pb2.GetDesiredStateConvergenceReq,
        prodvana.proto.prodvana.desired_state.manager_pb2.GetDesiredStateConvergenceSummaryResp,
    ]
    """Get the convergence summary for a desired state id.
    If the desired state id is pending, the returned summary will not have the entity_graph set but will have pending_set_desired_state set.
    The status will be set to PENDING_SET_DESIRED_STATE.
    If the desired state id was pending and failed to be set, the returned summary will not have the entity_graph set but will have pending_set_desired_state set.
    The status will be set to FAILED.
    """
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

class DesiredStateManagerServicer(metaclass=abc.ABCMeta):
    @abc.abstractmethod
    def SetDesiredState(
        self,
        request: prodvana.proto.prodvana.desired_state.manager_pb2.SetDesiredStateReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.desired_state.manager_pb2.SetDesiredStateResp: ...
    @abc.abstractmethod
    def PreviewEntityGraph(
        self,
        request: prodvana.proto.prodvana.desired_state.manager_pb2.SetDesiredStateReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.desired_state.manager_pb2.PreviewEntityGraphResp:
        """validate a SetDesiredState call and return a preview entity graph
        TODO(naphat) delete ValidateDesiredState and replace with this
        """
    @abc.abstractmethod
    def GetServiceDesiredStateConvergenceSummary(
        self,
        request: prodvana.proto.prodvana.desired_state.manager_pb2.GetServiceDesiredStateConvergenceSummaryReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.desired_state.manager_pb2.GetServiceDesiredStateConvergenceSummaryResp: ...
    @abc.abstractmethod
    def GetServiceLatestDesiredStateId(
        self,
        request: prodvana.proto.prodvana.desired_state.manager_pb2.GetServiceLatestDesiredStateIdReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.desired_state.manager_pb2.GetServiceLatestDesiredStateIdResp: ...
    @abc.abstractmethod
    def GetServiceDesiredStateIdHistory(
        self,
        request: prodvana.proto.prodvana.desired_state.manager_pb2.GetServiceDesiredStateIdHistoryReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.desired_state.manager_pb2.GetServiceDesiredStateIdHistoryResp: ...
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
    def GetDesiredStateGraph(
        self,
        request: prodvana.proto.prodvana.desired_state.manager_pb2.GetDesiredStateGraphReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.desired_state.manager_pb2.GetDesiredStateGraphResp: ...
    @abc.abstractmethod
    def GetDesiredStateConvergenceSummary(
        self,
        request: prodvana.proto.prodvana.desired_state.manager_pb2.GetDesiredStateConvergenceReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.desired_state.manager_pb2.GetDesiredStateConvergenceSummaryResp:
        """Get the convergence summary for a desired state id.
        If the desired state id is pending, the returned summary will not have the entity_graph set but will have pending_set_desired_state set.
        The status will be set to PENDING_SET_DESIRED_STATE.
        If the desired state id was pending and failed to be set, the returned summary will not have the entity_graph set but will have pending_set_desired_state set.
        The status will be set to FAILED.
        """
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

def add_DesiredStateManagerServicer_to_server(servicer: DesiredStateManagerServicer, server: grpc.Server) -> None: ...
