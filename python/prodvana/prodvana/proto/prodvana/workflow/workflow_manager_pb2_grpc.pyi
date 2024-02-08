"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import abc
import grpc
import prodvana.proto.prodvana.workflow.workflow_manager_pb2

class WorkflowManagerStub:
    def __init__(self, channel: grpc.Channel) -> None: ...
    ListIntegrations: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.ListIntegrationsReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.ListIntegrationsResp,
    ]
    DeleteIntegration: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.DeleteIntegrationReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.DeleteIntegrationResp,
    ]
    CreateContainerRegistryIntegration: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.CreateContainerRegistryIntegrationReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.CreateContainerRegistryIntegrationRes,
    ]
    DeleteContainerRegistryIntegration: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.DeleteContainerRegistryIntegrationReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.DeleteContainerRegistryIntegrationResp,
    ]
    ListContainerRegistryIntegrations: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.ListContainerRegistryIntegrationsReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.ListContainerRegistryIntegrationsResp,
    ]
    GetContainerRegistryIntegration: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetContainerRegistryIntegrationReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetContainerRegistryIntegrationResp,
    ]
    GetServiceImageInfo: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetServiceImageInfoReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetServiceImageInfoResp,
    ]
    GetContainerRegistryImages: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetContainerRegistryImagesReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetContainerRegistryImagesRes,
    ]
    ListTrackedImageRepositories: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.ListTrackedImageRepositoriesReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.ListTrackedImageRepositoriesResp,
    ]
    GetTrackedImageRepository: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetTrackedImageRepositoryReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetTrackedImageRepositoryResp,
    ]
    TrackImageRepositories: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.TrackImageRepositoriesReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.TrackImageRepositoriesResp,
    ]
    StopTrackingImageRepository: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.StopTrackingImageRepositoryReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.StopTrackingImageRepositoryResp,
    ]
    GetProgramDefaults: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetProgramDefaultsReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetProgramDefaultsResp,
    ]
    GetImageCommitInfo: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetImageCommitInfoReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetImageCommitInfoResp,
    ]
    InstallSlack: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.InstallSlackReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.InstallSlackResp,
    ]
    UninstallSlack: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.UninstallSlackReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.UninstallSlackResp,
    ]
    GetInstallSlackUrl: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetInstallSlackUrlReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetInstallSlackUrlResp,
    ]
    InstallPagerDuty: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.InstallPagerDutyReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.InstallPagerDutyResp,
    ]
    GetInstallPagerDutyUrl: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetInstallPagerDutyUrlReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetInstallPagerDutyUrlResp,
    ]
    UninstallPagerDuty: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.UninstallPagerDutyReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.UninstallPagerDutyResp,
    ]
    GetGrafanaInstallation: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetGrafanaInstallationReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetGrafanaInstallationResp,
    ]
    InstallGrafana: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.InstallGrafanaReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.InstallGrafanaResp,
    ]
    UninstallGrafana: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.UninstallGrafanaReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.UninstallGrafanaResp,
    ]
    ListHoneycombEnvironments: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.ListHoneycombEnvironmentsReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.ListHoneycombEnvironmentsResp,
    ]
    AddHoneycombEnvironment: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.AddHoneycombEnvironmentReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.AddHoneycombEnvironmentResp,
    ]
    UpdateHoneycombEnvironment: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.UpdateHoneycombEnvironmentReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.UpdateHoneycombEnvironmentResp,
    ]
    DeleteHoneycombEnvironment: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.DeleteHoneycombEnvironmentReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.DeleteHoneycombEnvironmentResp,
    ]
    UninstallHoneycomb: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.UninstallHoneycombReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.UninstallHoneycombResp,
    ]
    GetInstallGitHubUrl: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetInstallGitHubUrlReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetInstallGitHubUrlResp,
    ]
    CreateGitHubApp: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.CreateGitHubAppReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.CreateGitHubAppResp,
    ]
    InstallGitHub: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.InstallGitHubReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.InstallGitHubResp,
    ]
    ListRepoCommits: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.ListRepoCommitsReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.ListRepoCommitsResp,
    ]
    GetCommitInfo: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetCommitInfoReq,
        prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetCommitInfoResp,
    ]

class WorkflowManagerServicer(metaclass=abc.ABCMeta):
    @abc.abstractmethod
    def ListIntegrations(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.ListIntegrationsReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.ListIntegrationsResp: ...
    @abc.abstractmethod
    def DeleteIntegration(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.DeleteIntegrationReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.DeleteIntegrationResp: ...
    @abc.abstractmethod
    def CreateContainerRegistryIntegration(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.CreateContainerRegistryIntegrationReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.CreateContainerRegistryIntegrationRes: ...
    @abc.abstractmethod
    def DeleteContainerRegistryIntegration(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.DeleteContainerRegistryIntegrationReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.DeleteContainerRegistryIntegrationResp: ...
    @abc.abstractmethod
    def ListContainerRegistryIntegrations(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.ListContainerRegistryIntegrationsReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.ListContainerRegistryIntegrationsResp: ...
    @abc.abstractmethod
    def GetContainerRegistryIntegration(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetContainerRegistryIntegrationReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetContainerRegistryIntegrationResp: ...
    @abc.abstractmethod
    def GetServiceImageInfo(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetServiceImageInfoReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetServiceImageInfoResp: ...
    @abc.abstractmethod
    def GetContainerRegistryImages(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetContainerRegistryImagesReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetContainerRegistryImagesRes: ...
    @abc.abstractmethod
    def ListTrackedImageRepositories(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.ListTrackedImageRepositoriesReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.ListTrackedImageRepositoriesResp: ...
    @abc.abstractmethod
    def GetTrackedImageRepository(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetTrackedImageRepositoryReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetTrackedImageRepositoryResp: ...
    @abc.abstractmethod
    def TrackImageRepositories(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.TrackImageRepositoriesReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.TrackImageRepositoriesResp: ...
    @abc.abstractmethod
    def StopTrackingImageRepository(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.StopTrackingImageRepositoryReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.StopTrackingImageRepositoryResp: ...
    @abc.abstractmethod
    def GetProgramDefaults(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetProgramDefaultsReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetProgramDefaultsResp: ...
    @abc.abstractmethod
    def GetImageCommitInfo(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetImageCommitInfoReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetImageCommitInfoResp: ...
    @abc.abstractmethod
    def InstallSlack(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.InstallSlackReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.InstallSlackResp: ...
    @abc.abstractmethod
    def UninstallSlack(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.UninstallSlackReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.UninstallSlackResp: ...
    @abc.abstractmethod
    def GetInstallSlackUrl(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetInstallSlackUrlReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetInstallSlackUrlResp: ...
    @abc.abstractmethod
    def InstallPagerDuty(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.InstallPagerDutyReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.InstallPagerDutyResp: ...
    @abc.abstractmethod
    def GetInstallPagerDutyUrl(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetInstallPagerDutyUrlReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetInstallPagerDutyUrlResp: ...
    @abc.abstractmethod
    def UninstallPagerDuty(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.UninstallPagerDutyReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.UninstallPagerDutyResp: ...
    @abc.abstractmethod
    def GetGrafanaInstallation(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetGrafanaInstallationReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetGrafanaInstallationResp: ...
    @abc.abstractmethod
    def InstallGrafana(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.InstallGrafanaReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.InstallGrafanaResp: ...
    @abc.abstractmethod
    def UninstallGrafana(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.UninstallGrafanaReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.UninstallGrafanaResp: ...
    @abc.abstractmethod
    def ListHoneycombEnvironments(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.ListHoneycombEnvironmentsReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.ListHoneycombEnvironmentsResp: ...
    @abc.abstractmethod
    def AddHoneycombEnvironment(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.AddHoneycombEnvironmentReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.AddHoneycombEnvironmentResp: ...
    @abc.abstractmethod
    def UpdateHoneycombEnvironment(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.UpdateHoneycombEnvironmentReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.UpdateHoneycombEnvironmentResp: ...
    @abc.abstractmethod
    def DeleteHoneycombEnvironment(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.DeleteHoneycombEnvironmentReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.DeleteHoneycombEnvironmentResp: ...
    @abc.abstractmethod
    def UninstallHoneycomb(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.UninstallHoneycombReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.UninstallHoneycombResp: ...
    @abc.abstractmethod
    def GetInstallGitHubUrl(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetInstallGitHubUrlReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetInstallGitHubUrlResp: ...
    @abc.abstractmethod
    def CreateGitHubApp(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.CreateGitHubAppReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.CreateGitHubAppResp: ...
    @abc.abstractmethod
    def InstallGitHub(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.InstallGitHubReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.InstallGitHubResp: ...
    @abc.abstractmethod
    def ListRepoCommits(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.ListRepoCommitsReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.ListRepoCommitsResp: ...
    @abc.abstractmethod
    def GetCommitInfo(
        self,
        request: prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetCommitInfoReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.workflow.workflow_manager_pb2.GetCommitInfoResp: ...

def add_WorkflowManagerServicer_to_server(servicer: WorkflowManagerServicer, server: grpc.Server) -> None: ...
