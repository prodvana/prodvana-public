package runtimes

import (
	"context"
	"log"

	"prodvana/cmd/cmdutil"

	"github.com/prodvana/prodvana-public/go/pkg/protohelpers"
	"github.com/prodvana/prodvana-public/go/pkg/protoyaml"

	common_config_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config"
	environment_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/environment"
	version_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/version"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func ParseRuntimeConfig(ctx context.Context, runtimeId, inputFileType, maybeInput, startingYaml string, noPrompt bool, options ...cmdutil.EditOptions[*environment_pb.ClusterConfig]) (*environment_pb.ClusterConfig, error) {
	var config environment_pb.ClusterConfig
	_, err := cmdutil.EditOrReadConfig(ctx, inputFileType, maybeInput, startingYaml, &config, func(ctx context.Context, config *environment_pb.ClusterConfig) ([]*common_config_pb.DangerousAction, error) {
		_, err := cmdutil.GetEnvironmentManagerClient().ValidateConfigureCluster(ctx, &environment_pb.ConfigureClusterReq{
			RuntimeName: runtimeId,
			Config:      config,
		})
		if err != nil {
			return nil, errors.Wrapf(err, "validation failed")
		}
		return nil, nil
	}, "Save changes?", noPrompt, options...)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure a runtime.",
	Long: `Configure a runtime.

pvnctl runtimes configure [FLAGS]
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
		runtime := args[0]
		resp, err := cmdutil.GetEnvironmentManagerClient().GetClusterConfig(ctx, &environment_pb.GetClusterConfigReq{
			RuntimeName: runtime,
		})
		if err != nil {
			log.Fatal(err)
		}
		yamlBytes, err := protoyaml.Marshal(resp.Config)
		if err != nil {
			log.Fatal(err)
		}
		config, err := ParseRuntimeConfig(ctx, runtime, inputFileType, input, string(yamlBytes), noPrompt)
		if err != nil {
			log.Fatal(err)
		}
		_, err = cmdutil.GetEnvironmentManagerClient().ConfigureCluster(ctx,
			&environment_pb.ConfigureClusterReq{
				RuntimeName: runtime,
				Config:      config,
				Source:      version_pb.Source_INTERACTIVE_PVNCTL,
			},
		)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(configureCmd)

	configureCmd.Flags().String("input", "", "If set, take application config from a file instead of launching an editor. Supported formats: json, yaml, pbtxt")
	configureCmd.Flags().String("input-file-format", protohelpers.FileTypeInfer, "Used in conjunction with --input. Specify input file type. Pass 'infer' to infer file type from file extension.")
	configureCmd.Flags().BoolP("no-prompt", "y", false, "Whether to prompt for confirmation before saving.")
}
