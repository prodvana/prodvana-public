package applications

import (
	"context"
	_ "embed"
	"log"
	"prodvana/cmd/cmdutil"

	"github.com/prodvana/prodvana-public/go/pkg/protohelpers"
	"github.com/prodvana/prodvana-public/go/pkg/protoyaml"

	application_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/application"
	common_config_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config"

	"github.com/spf13/cobra"
)

//go:embed templates/application_metadata.yaml
var defaultMetadataYaml string

func ParseApplicationMetadata(ctx context.Context, inputFileType, maybeInput, startingYaml string, noPrompt bool, options ...cmdutil.EditOptions[*application_pb.ApplicationUserMetadata]) (*application_pb.ApplicationUserMetadata, error) {
	var metadata application_pb.ApplicationUserMetadata
	_, err := cmdutil.EditOrReadConfig(ctx, inputFileType, maybeInput, startingYaml, &metadata, func(ctx context.Context, metadata *application_pb.ApplicationUserMetadata) ([]*common_config_pb.DangerousAction, error) {
		return nil, nil
	}, "Save changes?", noPrompt, options...)
	if err != nil {
		return nil, err
	}
	return &metadata, nil
}

var editMetadataCmd = &cobra.Command{
	Use:   "edit-metadata",
	Short: "Edit metadata of an application.",
	Long: `Edit metadata of an application.

pvnctl applications edit-metadata <app>  # will launch editor
pvnctl applications edit --input metadata.yaml <app>
`,
	Args: cobra.ExactArgs(1),
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
		appName := args[0]
		resp, err := cmdutil.GetApplicationManagerClient().GetApplicationMetadata(ctx, &application_pb.GetApplicationMetadataReq{
			Application: appName,
		})
		if err != nil {
			log.Fatal(err)
		}
		var metadata *application_pb.ApplicationUserMetadata
		if resp.Metadata == nil {
			metadata, err = ParseApplicationMetadata(ctx, inputFileType, input, defaultMetadataYaml, noPrompt)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			yamlBytes, err := protoyaml.Marshal(resp.Metadata)
			if err != nil {
				log.Fatal(err)
			}
			metadata, err = ParseApplicationMetadata(ctx, inputFileType, input, string(yamlBytes), noPrompt)
			if err != nil {
				log.Fatal(err)
			}
		}
		_, err = cmdutil.GetApplicationManagerClient().SetApplicationMetadata(ctx, &application_pb.SetApplicationMetadataReq{
			Application: appName,
			Metadata:    metadata,
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(editMetadataCmd)
	editMetadataCmd.Flags().String("input", "", "If set, take application metadata from a file instead of launching an editor. Supported formats: json, yaml, pbtxt")
	editMetadataCmd.Flags().String("input-file-format", protohelpers.FileTypeInfer, "Used in conjunction with --input. Specify input file type. Pass 'infer' to infer file type from file extension.")
	editMetadataCmd.Flags().BoolP("no-prompt", "y", false, "Whether to prompt for confirmation before saving.")
}
