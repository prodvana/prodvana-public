package registry

import (
	"context"
	"log"
	"os"

	"prodvana/cmd/cmdutil"

	workflow_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/workflow"

	"github.com/spf13/cobra"
)

var listImagesCmd = &cobra.Command{
	Use:   "list-images",
	Short: "List container images for a given repository.",
	Run: func(cmd *cobra.Command, args []string) {
		repository, err := cmd.Flags().GetString("repository")
		if err != nil {
			log.Fatal(err)
		}
		integrationID, err := cmd.Flags().GetString("integration-id")
		if err != nil {
			log.Fatal(err)
		}
		skipRegistryCache, err := cmd.Flags().GetBool("skip-registry-cache")
		if err != nil {
			log.Fatal(err)
		}

		var token string
		for {
			resp, err := cmdutil.GetWorkflowManagerClient().GetContainerRegistryImages(context.Background(), &workflow_pb.GetContainerRegistryImagesReq{
				IntegrationId:     integrationID,
				ImageRepository:   repository,
				PageToken:         token,
				SkipRegistryCache: skipRegistryCache,
			})
			if err != nil {
				log.Fatal(err)
			}

			// TODO: May be a good idea to omit COMMIT if none of the images have commit info.
			listing := cmdutil.NewOutputListing("URL", "CREATED", "COMMIT")

			for _, img := range resp.Images {
				commit := ""
				if img.Commit != nil {
					commit = img.Commit.CommitId
				}
				listing.AddRow(img.Url, img.Created.AsTime(), commit)
			}

			err = cmdutil.WriteStdout(listing)
			if err != nil {
				cmd.PrintErrf("Failed to list images: %+v\n", err)
				os.Exit(1)
			}

			if resp.NextPageToken == "" {
				break
			}

			token = resp.NextPageToken
		}

	},
}

func init() {
	RootCmd.AddCommand(listImagesCmd)
	listImagesCmd.Flags().String("integration-id", "", "Integration ID of registry to use")
	listImagesCmd.Flags().String("repository", "", "Image repository to list")
	listImagesCmd.Flags().Bool("skip-registry-cache", false, "Skip registry cache and fetch images from registry directly")
	cmdutil.Must(listImagesCmd.MarkFlagRequired("integration-id"))
	cmdutil.Must(listImagesCmd.MarkFlagRequired("repository"))
}
