package cmdutil

import (
	"sync"

	"github.com/prodvana/prodvana-public/go/pkg/client"

	application_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/application"
	auth_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/auth"
	blobs_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/blobs"
	delivery_extension_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/delivery_extension"
	deployment_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/deployment"
	ds_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state"
	environment_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/environment"
	events_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/events"
	object_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/object"
	organization_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/organization"
	protection_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/protection"
	recipe_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/recipe"
	release_channel_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/release_channel"
	secrets_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/secrets"
	service_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/service"
	workflow_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/workflow"

	"google.golang.org/grpc/health/grpc_health_v1"
)

var (
	organizationManagerClient     organization_pb.OrganizationManagerClient
	organizationManagerClientOnce sync.Once

	applicationManagerClient     application_pb.ApplicationManagerClient
	applicationManagerClientOnce sync.Once

	serviceManagerClient     service_pb.ServiceManagerClient
	serviceManagerClientOnce sync.Once

	environmentManagerClient     environment_pb.EnvironmentManagerClient
	environmentManagerClientOnce sync.Once

	protectionManagerClient     protection_pb.ProtectionManagerClient
	protectionManagerClientOnce sync.Once

	deliveryTaskManagerClient     delivery_extension_pb.DeliveryExtensionManagerClient
	deliveryTaskManagerClientOnce sync.Once

	releaseChannelManagerClient     release_channel_pb.ReleaseChannelManagerClient
	releaseChannelManagerClientOnce sync.Once

	recipeManagerClient     recipe_pb.RecipeManagerClient
	recipeManagerClientOnce sync.Once

	workflowManagerClient     workflow_pb.WorkflowManagerClient
	workflowManagerClientOnce sync.Once

	apiTokenManagerClient     auth_pb.ApiTokenManagerClient
	apiTokenManagerClientOnce sync.Once

	desiredStateManagerClient     ds_pb.DesiredStateManagerClient
	desiredStateManagerClientOnce sync.Once

	blobsManagerClient     blobs_pb.BlobsManagerClient
	blobsManagerClientOnce sync.Once

	eventsManagerClient     events_pb.EventsManagerClient
	eventsManagerClientOnce sync.Once

	secretsManagerClient     secrets_pb.SecretsManagerClient
	secretsManagerClientOnce sync.Once

	objectManagerClient     object_pb.ObjectManagerClient
	objectManagerClientOnce sync.Once

	deploymentManagerClient     deployment_pb.DeploymentManagerClient
	deploymentManagerClientOnce sync.Once

	healthClient     grpc_health_v1.HealthClient
	healthClientOnce sync.Once
)

func GetOrganizationManagerClient() organization_pb.OrganizationManagerClient {
	organizationManagerClientOnce.Do(func() {
		organizationManagerClient = organization_pb.NewOrganizationManagerClient(client.GetApiserverConn())
	})
	return organizationManagerClient
}

func GetApplicationManagerClient() application_pb.ApplicationManagerClient {
	applicationManagerClientOnce.Do(func() {
		applicationManagerClient = application_pb.NewApplicationManagerClient(client.GetApiserverConn())
	})
	return applicationManagerClient
}

func GetServiceManagerClient() service_pb.ServiceManagerClient {
	serviceManagerClientOnce.Do(func() {
		serviceManagerClient = service_pb.NewServiceManagerClient(client.GetApiserverConn())
	})
	return serviceManagerClient
}

func GetEnvironmentManagerClient() environment_pb.EnvironmentManagerClient {
	environmentManagerClientOnce.Do(func() {
		environmentManagerClient = environment_pb.NewEnvironmentManagerClient(client.GetApiserverConn())
	})
	return environmentManagerClient
}

func GetReleaseChannelManagerClient() release_channel_pb.ReleaseChannelManagerClient {
	releaseChannelManagerClientOnce.Do(func() {
		releaseChannelManagerClient = release_channel_pb.NewReleaseChannelManagerClient(client.GetApiserverConn())
	})
	return releaseChannelManagerClient
}

func GetRecipeManagerClient() recipe_pb.RecipeManagerClient {
	recipeManagerClientOnce.Do(func() {
		recipeManagerClient = recipe_pb.NewRecipeManagerClient(client.GetApiserverConn())
	})
	return recipeManagerClient
}

func GetWorkflowManagerClient() workflow_pb.WorkflowManagerClient {
	workflowManagerClientOnce.Do(func() {
		workflowManagerClient = workflow_pb.NewWorkflowManagerClient(client.GetApiserverConn())
	})
	return workflowManagerClient
}

func GetHealthClient() grpc_health_v1.HealthClient {
	healthClientOnce.Do(func() {
		healthClient = grpc_health_v1.NewHealthClient(client.GetApiserverConn())
	})
	return healthClient
}

func GetApiTokenManagerClient() auth_pb.ApiTokenManagerClient {
	apiTokenManagerClientOnce.Do(func() {
		apiTokenManagerClient = auth_pb.NewApiTokenManagerClient(client.GetApiserverConn())
	})
	return apiTokenManagerClient
}

func GetDesiredStateManagerClient() ds_pb.DesiredStateManagerClient {
	desiredStateManagerClientOnce.Do(func() {
		desiredStateManagerClient = ds_pb.NewDesiredStateManagerClient(client.GetApiserverConn())
	})
	return desiredStateManagerClient
}

func GetProtectionManagerClient() protection_pb.ProtectionManagerClient {
	protectionManagerClientOnce.Do(func() {
		protectionManagerClient = protection_pb.NewProtectionManagerClient(client.GetApiserverConn())
	})
	return protectionManagerClient
}

func GetDeliveryExtensionManagerClient() delivery_extension_pb.DeliveryExtensionManagerClient {
	deliveryTaskManagerClientOnce.Do(func() {
		deliveryTaskManagerClient = delivery_extension_pb.NewDeliveryExtensionManagerClient(client.GetApiserverConn())
	})
	return deliveryTaskManagerClient
}

func GetBlobsManagerClient() blobs_pb.BlobsManagerClient {
	blobsManagerClientOnce.Do(func() {
		blobsManagerClient = blobs_pb.NewBlobsManagerClient(client.GetApiserverConn())
	})
	return blobsManagerClient
}

func GetEventsManagerClient() events_pb.EventsManagerClient {
	eventsManagerClientOnce.Do(func() {
		eventsManagerClient = events_pb.NewEventsManagerClient(client.GetApiserverConn())
	})
	return eventsManagerClient
}

func GetSecretsManagerClient() secrets_pb.SecretsManagerClient {
	secretsManagerClientOnce.Do(func() {
		secretsManagerClient = secrets_pb.NewSecretsManagerClient(client.GetApiserverConn())
	})
	return secretsManagerClient
}

func GetObjectManagerClient() object_pb.ObjectManagerClient {
	objectManagerClientOnce.Do(func() {
		objectManagerClient = object_pb.NewObjectManagerClient(client.GetApiserverConn())
	})
	return objectManagerClient
}

func GetDeploymentManagerClient() deployment_pb.DeploymentManagerClient {
	deploymentManagerClientOnce.Do(func() {
		deploymentManagerClient = deployment_pb.NewDeploymentManagerClient(client.GetApiserverConn())
	})
	return deploymentManagerClient
}
