package registry

import (
	"context"
	"log"
	"os"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	workflow_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/workflow"

	"github.com/spf13/cobra"
)

var imageInfoCmd = &cobra.Command{
	Use:    "image-info",
	Short:  "Inspect container startup info.",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		repository, err := cmd.Flags().GetString("repository")
		if err != nil {
			log.Fatal(err)
		}
		integrationID, err := cmd.Flags().GetString("integration-id")
		if err != nil {
			log.Fatal(err)
		}
		tag, err := cmd.Flags().GetString("tag")
		if err != nil {
			log.Fatal(err)
		}
		resp, err := cmdutil.GetWorkflowManagerClient().GetProgramDefaults(context.Background(), &workflow_pb.GetProgramDefaultsReq{
			IntegrationId: integrationID,
			Repository:    repository,
			ImageIdOneof:  &workflow_pb.GetProgramDefaultsReq_Tag{Tag: tag},
		})
		if err != nil {
			log.Fatal(err)
		}

		imageInfo := cmdutil.NewOutputListing("CMD", "ENTRYPOINT", "ENV", "PORTS")
		imageInfo.AddRow(resp.ProgramDefaults.Cmd, resp.ProgramDefaults.Entrypoint, resp.ProgramDefaults.Env, resp.ProgramDefaults.Ports)

		err = cmdutil.WriteStdout(imageInfo)
		if err != nil {
			cmd.PrintErrf("Failed to dump image info: %+v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(imageInfoCmd)
	imageInfoCmd.Flags().String("integration-id", "", "Integration ID of registry to use")
	imageInfoCmd.Flags().String("repository", "", "Image repository of container to inspect")
	imageInfoCmd.Flags().String("tag", "", "Tag to inspect")
	cmdutil.Must(imageInfoCmd.MarkFlagRequired("integration-id"))
	cmdutil.Must(imageInfoCmd.MarkFlagRequired("repository"))
	cmdutil.Must(imageInfoCmd.MarkFlagRequired("tag"))
}
