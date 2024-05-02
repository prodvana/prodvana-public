package services

import (
	"context"
	_ "embed"
	"log"
	"prodvana/cmd/cmdutil"

	"github.com/prodvana/prodvana-public/go/pkg/protohelpers"
	"github.com/prodvana/prodvana-public/go/pkg/protoyaml"

	common_config_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config"
	service_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/service"

	"github.com/spf13/cobra"
)

//go:embed templates/service_metadata.yaml
var defaultMetadataYaml string

func ParseServiceMetadata(ctx context.Context, inputFileType, maybeInput, startingYaml string, noPrompt bool, options ...cmdutil.EditOptions[*service_pb.ServiceUserMetadata]) (*service_pb.ServiceUserMetadata, error) {
	var metadata service_pb.ServiceUserMetadata
	_, err := cmdutil.EditOrReadConfig(ctx, inputFileType, maybeInput, startingYaml, &metadata, func(ctx context.Context, metadata *service_pb.ServiceUserMetadata) ([]*common_config_pb.DangerousAction, error) {
		return nil, nil
	}, "Save changes?", noPrompt, options...)
	if err != nil {
		return nil, err
	}
	return &metadata, nil
}

var editMetadataCmd = &cobra.Command{
	Use:   "edit-metadata",
	Short: "Edit metadata of an service.",
	Long: `Edit metadata of an service.

pvnctl services --app <app> edit-metadata <service>  # will launch editor
pvnctl services --app <app> edit --input metadata.yaml <service>
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
		serviceName := args[0]
		resp, err := cmdutil.GetServiceManagerClient().GetServiceMetadata(ctx, &service_pb.GetServiceMetadataReq{
			Application: cmdutil.App,
			Service:     serviceName,
		})
		if err != nil {
			log.Fatal(err)
		}
		var metadata *service_pb.ServiceUserMetadata
		if resp.Metadata == nil {
			metadata, err = ParseServiceMetadata(ctx, inputFileType, input, defaultMetadataYaml, noPrompt)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			yamlBytes, err := protoyaml.Marshal(resp.Metadata)
			if err != nil {
				log.Fatal(err)
			}
			metadata, err = ParseServiceMetadata(ctx, inputFileType, input, string(yamlBytes), noPrompt)
			if err != nil {
				log.Fatal(err)
			}
		}
		_, err = cmdutil.GetServiceManagerClient().SetServiceMetadata(ctx, &service_pb.SetServiceMetadataReq{
			Application: cmdutil.App,
			Service:     serviceName,
			Metadata:    metadata,
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(editMetadataCmd)
	editMetadataCmd.Flags().String("input", "", "If set, take service metadata from a file instead of launching an editor. Supported formats: json, yaml, pbtxt")
	editMetadataCmd.Flags().String("input-file-format", protohelpers.FileTypeInfer, "Used in conjunction with --input. Specify input file type. Pass 'infer' to infer file type from file extension.")
	editMetadataCmd.Flags().BoolP("no-prompt", "y", false, "Whether to prompt for confirmation before saving.")
}
