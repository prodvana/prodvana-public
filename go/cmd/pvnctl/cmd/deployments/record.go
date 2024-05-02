package deployments

import (
	"context"
	"fmt"
	"log"
	"prodvana/cmd/cmdutil"

	deployment_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/deployment"
	deployment_model_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/deployment/model"
	"github.com/spf13/cobra"
)

var recordFlags = struct {
	deploymentSystem string
	application      string
	service          string
	releaseChannel   string
	repository       string
	commit           string
	pending          bool
}{}

var recordCmd = &cobra.Command{
	Use:   "record",
	Short: "Record a deployment.",
	Long: `Record a deployment and print the resulting deployment id to stdout.

pvnctl deployments record --deployment-system=circleci
`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		if recordFlags.repository == "" || recordFlags.commit == "" {
			ciDefaults := cmdutil.DetectCIDefaults()
			if ciDefaults == nil {
				log.Fatal("Could not detect CI environment. Please pass --repository and --commit explicitly")
			}
			if recordFlags.repository == "" {
				recordFlags.repository = ciDefaults.Repository
			}
			if recordFlags.commit == "" {
				recordFlags.commit = ciDefaults.Commit
			}
		}
		resp, err := cmdutil.GetDeploymentManagerClient().RecordDeployment(ctx, &deployment_pb.RecordDeploymentReq{
			Config: &deployment_model_pb.DeploymentConfig{
				DeploymentSystem: recordFlags.deploymentSystem,
				Application:      recordFlags.application,
				Service:          recordFlags.service,
				ReleaseChannel:   recordFlags.releaseChannel,
				Repository:       recordFlags.repository,
				CommitId:         recordFlags.commit,
			},
		})
		if err != nil {
			log.Fatal("failed to record deployment", err)
		}
		fmt.Printf("%s\n", resp.Meta.Id)
	},
}

func init() {
	RootCmd.AddCommand(recordCmd)
	recordCmd.Flags().StringVarP(&recordFlags.deploymentSystem, "deployment-system", "d", "", "Deployment system doing this deployment")
	cmdutil.Must(recordCmd.MarkFlagRequired("deployment-system"))
	recordCmd.Flags().StringVarP(&recordFlags.application, "app", "a", "", "Application of the service being deployed")
	cmdutil.Must(recordCmd.MarkFlagRequired("app"))
	recordCmd.Flags().StringVarP(&recordFlags.service, "service", "s", "", "Service being deployed")
	cmdutil.Must(recordCmd.MarkFlagRequired("service"))
	recordCmd.Flags().StringVar(&recordFlags.releaseChannel, "release-channel", "", "Release Channel this deploy is happening in, e.g. staging, prod-eu")
	cmdutil.Must(recordCmd.MarkFlagRequired("release-channel"))
	recordCmd.Flags().StringVarP(&recordFlags.repository, "repository", "r", "", "Repository associated with this deployment. If not passed, will be inferred from CI environment variables. It is an error if this is not passed and the value cannot be inferred.")
	recordCmd.Flags().StringVarP(&recordFlags.commit, "commit", "c", "", "Commit associated with this deployment. If not passed, will be inferred from CI environment variables. It is an error if this is not passed and the value cannot be inferred.")
	recordCmd.Flags().BoolVar(&recordFlags.pending, "pending", false, "Create deployment with PENDING status instead of SUCCEEDED status, useful if the deployment status can be updated later.")
}
