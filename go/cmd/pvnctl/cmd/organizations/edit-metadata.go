package organizations

import (
	"context"
	_ "embed"
	"log"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	"github.com/prodvana/prodvana-public/go/pkg/protohelpers"
	"github.com/prodvana/prodvana-public/go/pkg/protoyaml"

	common_config_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config"
	organization_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/organization"

	"github.com/spf13/cobra"
)

//go:embed templates/organization_metadata.yaml
var defaultMetadataYaml string

func handleOrganizationMetadata(ctx context.Context, inputFileType, maybeInput, startingYaml string, noPrompt bool) (*organization_pb.OrganizationUserMetadata, error) {
	var metadata organization_pb.OrganizationUserMetadata
	_, err := cmdutil.EditOrReadConfig(ctx, inputFileType, maybeInput, startingYaml, &metadata, func(ctx context.Context, metadata *organization_pb.OrganizationUserMetadata) ([]*common_config_pb.DangerousAction, error) {
		return nil, nil
	}, "Save changes?", noPrompt)
	if err != nil {
		return nil, err
	}
	return &metadata, nil
}

var editMetadataCmd = &cobra.Command{
	Use:   "edit-metadata",
	Short: "Edit metadata of an organization.",
	Long: `Edit metadata of an organization.

pvnctl organizations --app <app> edit-metadata <organization>  # will launch editor
pvnctl organizations --app <app> edit --input metadata.yaml <organization>
`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		input, err := cmd.Flags().GetString("input")
		if err != nil {
			log.Fatal(err)
		}
		inputFileType, err := cmd.Flags().GetString("input-file-format")
		if err != nil {
			log.Fatal(err)
		}
		noPrompt, err := cmd.Flags().GetBool("no-prompt")
		if err != nil {
			log.Fatal(err)
		}
		resp, err := cmdutil.GetOrganizationManagerClient().GetOrganizationMetadata(ctx, &organization_pb.GetOrganizationMetadataReq{})
		if err != nil {
			log.Fatal(err)
		}
		var metadata *organization_pb.OrganizationUserMetadata
		if resp.Metadata == nil {
			metadata, err = handleOrganizationMetadata(ctx, inputFileType, input, defaultMetadataYaml, noPrompt)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			yamlBytes, err := protoyaml.Marshal(resp.Metadata)
			if err != nil {
				log.Fatal(err)
			}
			metadata, err = handleOrganizationMetadata(ctx, inputFileType, input, string(yamlBytes), noPrompt)
			if err != nil {
				log.Fatal(err)
			}
		}
		_, err = cmdutil.GetOrganizationManagerClient().SetOrganizationMetadata(ctx, &organization_pb.SetOrganizationMetadataReq{
			Metadata: metadata,
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(editMetadataCmd)
	editMetadataCmd.Flags().String("input", "", "If set, take organization metadata from a file instead of launching an editor. Supported formats: json, yaml, pbtxt")
	editMetadataCmd.Flags().String("input-file-format", protohelpers.FileTypeInfer, "Used in conjunction with --input. Specify input file type. Pass 'infer' to infer file type from file extension.")
	editMetadataCmd.Flags().BoolP("no-prompt", "y", false, "Whether to prompt for confirmation before saving.")
}
