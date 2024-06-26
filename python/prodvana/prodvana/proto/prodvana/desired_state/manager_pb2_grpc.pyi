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
    BypassConcurrencyLimit: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.desired_state.manager_pb2.BypassConcurrencyLimitReq,
        prodvana.proto.prodvana.desired_state.manager_pb2.BypassConcurrencyLimitResp,
    ]
    BypassDependencies: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.desired_state.manager_pb2.BypassDependenciesReq,
        prodvana.proto.prodvana.desired_state.manager_pb2.BypassDependenciesResp,
    ]
    ForceExecuteTask: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.desired_state.manager_pb2.ForceExecuteTaskReq,
        prodvana.proto.prodvana.desired_state.manager_pb2.ForceExecuteTaskResp,
    ]
    ListMaestroReleases: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.desired_state.manager_pb2.ListMaestroReleasesReq,
        prodvana.proto.prodvana.desired_state.manager_pb2.ListMaestroReleasesResp,
    ]
    GetMaestroRelease: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.desired_state.manager_pb2.GetMaestroReleaseReq,
        prodvana.proto.prodvana.desired_state.manager_pb2.GetMaestroReleaseResp,
    ]
    ListCombinedReleases: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.desired_state.manager_pb2.ListCombinedReleasesReq,
        prodvana.proto.prodvana.desired_state.manager_pb2.ListCombinedReleasesResp,
    ]
    """Get release history, where a release is either a Maestro Release or a Desired State from calling SetDesiredState"""
    ListServiceCombinedReleases: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.desired_state.manager_pb2.ListServiceCombinedReleasesReq,
        prodvana.proto.prodvana.desired_state.manager_pb2.ListCombinedReleasesResp,
    ]
    """Like ListCombinedReleases, but accepts an application/service names/ids instead of entity ID"""
    GetLatestCombinedReleaseDesiredState: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.desired_state.manager_pb2.GetLatestCombinedReleaseDesiredStateReq,
        prodvana.proto.prodvana.desired_state.manager_pb2.GetLatestCombinedReleaseDesiredStateResp,
    ]
    """get latest desired state that was set explicitly as part of a release, as defined by
    ListCombinedReleases. This is a shortcut for ListCombinedReleases(descending=True, page_size=1),
    then GetDesiredState or GetMaestroRelease depending on the type of release.
    """
    GetServiceLatestCombinedReleaseDesiredState: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.desired_state.manager_pb2.GetServiceLatestCombinedReleaseDesiredStateReq,
        prodvana.proto.prodvana.desired_state.manager_pb2.GetLatestCombinedReleaseDesiredStateResp,
    ]
    GetDebugState: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.desired_state.manager_pb2.GetDebugStateReq,
        prodvana.proto.prodvana.desired_state.manager_pb2.GetDebugStateResp,
    ]
    GetHistoricalEntityStats: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.desired_state.manager_pb2.GetHistoricalEntityStatsReq,
        prodvana.proto.prodvana.desired_state.manager_pb2.GetHistoricalEntityStatsResp,
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
    @abc.abstractmethod
    def BypassConcurrencyLimit(
        self,
        request: prodvana.proto.prodvana.desired_state.manager_pb2.BypassConcurrencyLimitReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.desired_state.manager_pb2.BypassConcurrencyLimitResp: ...
    @abc.abstractmethod
    def BypassDependencies(
        self,
        request: prodvana.proto.prodvana.desired_state.manager_pb2.BypassDependenciesReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.desired_state.manager_pb2.BypassDependenciesResp: ...
    @abc.abstractmethod
    def ForceExecuteTask(
        self,
        request: prodvana.proto.prodvana.desired_state.manager_pb2.ForceExecuteTaskReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.desired_state.manager_pb2.ForceExecuteTaskResp: ...
    @abc.abstractmethod
    def ListMaestroReleases(
        self,
        request: prodvana.proto.prodvana.desired_state.manager_pb2.ListMaestroReleasesReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.desired_state.manager_pb2.ListMaestroReleasesResp: ...
    @abc.abstractmethod
    def GetMaestroRelease(
        self,
        request: prodvana.proto.prodvana.desired_state.manager_pb2.GetMaestroReleaseReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.desired_state.manager_pb2.GetMaestroReleaseResp: ...
    @abc.abstractmethod
    def ListCombinedReleases(
        self,
        request: prodvana.proto.prodvana.desired_state.manager_pb2.ListCombinedReleasesReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.desired_state.manager_pb2.ListCombinedReleasesResp:
        """Get release history, where a release is either a Maestro Release or a Desired State from calling SetDesiredState"""
    @abc.abstractmethod
    def ListServiceCombinedReleases(
        self,
        request: prodvana.proto.prodvana.desired_state.manager_pb2.ListServiceCombinedReleasesReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.desired_state.manager_pb2.ListCombinedReleasesResp:
        """Like ListCombinedReleases, but accepts an application/service names/ids instead of entity ID"""
    @abc.abstractmethod
    def GetLatestCombinedReleaseDesiredState(
        self,
        request: prodvana.proto.prodvana.desired_state.manager_pb2.GetLatestCombinedReleaseDesiredStateReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.desired_state.manager_pb2.GetLatestCombinedReleaseDesiredStateResp:
        """get latest desired state that was set explicitly as part of a release, as defined by
        ListCombinedReleases. This is a shortcut for ListCombinedReleases(descending=True, page_size=1),
        then GetDesiredState or GetMaestroRelease depending on the type of release.
        """
    @abc.abstractmethod
    def GetServiceLatestCombinedReleaseDesiredState(
        self,
        request: prodvana.proto.prodvana.desired_state.manager_pb2.GetServiceLatestCombinedReleaseDesiredStateReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.desired_state.manager_pb2.GetLatestCombinedReleaseDesiredStateResp: ...
    @abc.abstractmethod
    def GetDebugState(
        self,
        request: prodvana.proto.prodvana.desired_state.manager_pb2.GetDebugStateReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.desired_state.manager_pb2.GetDebugStateResp: ...
    @abc.abstractmethod
    def GetHistoricalEntityStats(
        self,
        request: prodvana.proto.prodvana.desired_state.manager_pb2.GetHistoricalEntityStatsReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.desired_state.manager_pb2.GetHistoricalEntityStatsResp: ...

def add_DesiredStateManagerServicer_to_server(servicer: DesiredStateManagerServicer, server: grpc.Server) -> None: ...
