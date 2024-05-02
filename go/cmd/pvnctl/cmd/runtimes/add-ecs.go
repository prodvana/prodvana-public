package runtimes

import (
	"context"
	"fmt"
	"os"

	"prodvana/cmd/cmdutil"

	"github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/environment"
	version_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/version"

	"github.com/spf13/cobra"
)

var addECSCmd = &cobra.Command{
	Use:    "add-ecs <name>",
	Hidden: true,
	Short:  "Add an existing ECS runtime to Prodvana.",
	Long: `Add an existing ECS runtime to Prodvana. Prints the ID of the new runtime to stdout on success.

pvnctl runtimes add-ecs <name> [FLAGS]
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		ctx := context.Background()
		resp, err := cmdutil.GetEnvironmentManagerClient().LinkCluster(ctx,
			&environment.LinkClusterReq{
				Name: name,
				Type: environment.ClusterType_ECS,
				Auth: &environment.ClusterAuth{
					AuthOneof: &environment.ClusterAuth_Ecs{
						Ecs: &environment.ClusterAuth_ECSAuth{
							AccessKey:     cmdutil.MustGetString(cmd, "access-key"),
							SecretKey:     cmdutil.MustGetString(cmd, "secret-key"),
							ClusterArn:    cmdutil.MustGetString(cmd, "cluster-arn"),
							Region:        cmdutil.MustGetString(cmd, "region"),
							AssumeRoleArn: cmdutil.MustGetString(cmd, "assume-role-arn"),
						},
					},
				},
				Source: version_pb.Source_INTERACTIVE_PVNCTL,
			})
		if err != nil {
			cmd.PrintErrf("Failed to add ECS cluster: %+v\n", err)
			os.Exit(1)
		}
		if !resp.Success {
			cmd.PrintErrf("Failed to add cluster: %s\n", resp.Msg)
			os.Exit(1)
		}

		fmt.Printf("%s\n", resp.ClusterId)
	},
}

func init() {
	RootCmd.AddCommand(addECSCmd)

	addECSCmd.Flags().String("cluster-arn", "", "ECS Cluster ARN")
	addECSCmd.Flags().String("assume-role-arn", "", "Role to assume")
	addECSCmd.Flags().String("access-key", "", "AWS Access Key")
	addECSCmd.Flags().String("secret-key", "", "AWS Secret Key")
	addECSCmd.Flags().String("region", "", "AWS Region")
	cmdutil.Must(addECSCmd.MarkFlagRequired("cluster-arn"))
	cmdutil.Must(addECSCmd.MarkFlagRequired("access-key"))
	cmdutil.Must(addECSCmd.MarkFlagRequired("secret-key"))
	cmdutil.Must(addECSCmd.MarkFlagRequired("region"))
}
