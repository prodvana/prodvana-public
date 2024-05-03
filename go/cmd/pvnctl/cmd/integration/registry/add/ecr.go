package add

import (
	"context"
	"fmt"
	"log"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	workflow_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/workflow"

	"github.com/spf13/cobra"
)

var ecrCmd = &cobra.Command{
	Use:   "ecr",
	Short: "Add or update an ECR integration",
	Long: `Add or update an ECR integration. Registry names are unique, so passing an existing name will update that integration.

	pvnctl integration registry add ecr <name> --access-key <key> --secret-key <secret> --region <region>
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		accessKey, err := cmd.Flags().GetString("access-key")
		if err != nil {
			log.Fatal(err)
		}
		secretKey, err := cmd.Flags().GetString("secret-key")
		if err != nil {
			log.Fatal(err)
		}
		region, err := cmd.Flags().GetString("region")
		if err != nil {
			log.Fatal(err)
		}
		roleArn, err := cmd.Flags().GetString("role-arn")
		if err != nil {
			log.Fatal(err)
		}

		req := &workflow_pb.CreateContainerRegistryIntegrationReq{
			Name: name,
			Type: workflow_pb.RegistryType_ECR,
			RegistryOptions: &workflow_pb.CreateContainerRegistryIntegrationReq_EcrOptions{
				EcrOptions: &workflow_pb.CreateContainerRegistryIntegrationReq_ECROptions{
					AccessKey: accessKey,
					SecretKey: secretKey,
					Region:    region,
					RoleArn:   roleArn,
				},
			},
		}
		resp, err := cmdutil.GetWorkflowManagerClient().CreateContainerRegistryIntegration(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(resp.IntegrationId)
	}}

func init() {
	RootCmd.AddCommand(ecrCmd)
	ecrCmd.Flags().String("access-key", "", "AWS Access Key")
	ecrCmd.Flags().String("secret-key", "", "AWS Secret Key")
	ecrCmd.Flags().String("region", "", "AWS Region")
	ecrCmd.Flags().String("role-arn", "", "AWS ARN of IAM role to assume while accessing Docker Registry. Useful to access registry in a different account from the current user.")
	cmdutil.Must(ecrCmd.MarkFlagRequired("access-key"))
	cmdutil.Must(ecrCmd.MarkFlagRequired("secret-key"))
	cmdutil.Must(ecrCmd.MarkFlagRequired("region"))
}
