package releasechannels

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	"github.com/prodvana/prodvana-public/go/pkg/protohelpers"
	"github.com/prodvana/prodvana-public/go/pkg/protoyaml"

	common_config_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config"
	release_channel_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/release_channel"
	version_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/version"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

//go:embed templates/release_channel_config.yaml
var defaultConfigYaml string

func ParseReleaseChannelConfig(ctx context.Context, inputFileType, maybeInput, maybeRequiredName, startingYaml string, noPrompt bool, options ...cmdutil.EditOptions[*release_channel_pb.ReleaseChannelConfig]) (*release_channel_pb.ReleaseChannelConfig, error) {
	var config release_channel_pb.ReleaseChannelConfig
	_, err := cmdutil.EditOrReadConfig(ctx, inputFileType, maybeInput, startingYaml, &config, func(ctx context.Context, config *release_channel_pb.ReleaseChannelConfig) ([]*common_config_pb.DangerousAction, error) {
		if config.Name == "" {
			return nil, errors.Errorf("name must be set")
		}
		if maybeRequiredName != "" && config.Name != maybeRequiredName {
			return nil, errors.Errorf("name must be %s", maybeRequiredName)
		}
		return nil, nil
	}, "Save changes?", noPrompt, options...)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func configureReleaseChannel(ctx context.Context, application string, config *release_channel_pb.ReleaseChannelConfig) error {
	resp, err := cmdutil.GetReleaseChannelManagerClient().ConfigureReleaseChannel(ctx, &release_channel_pb.ConfigureReleaseChannelReq{
		ReleaseChannel: config,
		Source:         version_pb.Source_INTERACTIVE_PVNCTL,
		Application:    application,
	})
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", resp.Version)
	return nil
}

var configureCmd = &cobra.Command{
	Use:    "configure <name>",
	Short:  "Deprecated, please use 'pvnctl applications edit/create' instead.",
	Hidden: true,
	Long: `Deprecated, please use 'pvnctl applications edit/create' instead.
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.PrintErrf("Please run this command instead to add/edit release channels:\n\n# https://docs.prodvana.io/docs/release-channels#creatingediting-a-release-channel\npvnctl applications edit default\n")
		os.Exit(2)
	},
}

var editCmd = &cobra.Command{
	Use:   "edit <name>",
	Short: "Edit an existing release channel.",
	Long: `Edit an existing release channel.

pvnctl release-channels --app=<app> edit <app>  # will launch editor
pvnctl release-channels --app=<app> edit --input config.yaml <app>
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
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
		rcName := args[0]
		ctx := context.Background()
		resp, err := cmdutil.GetReleaseChannelManagerClient().GetReleaseChannelConfig(ctx, &release_channel_pb.GetReleaseChannelConfigReq{
			Application:    cmdutil.App,
			ReleaseChannel: rcName,
		})
		if err != nil {
			cmd.PrintErrf("Failed to fetch current release channel config: %+v\n", err)
			os.Exit(1)
		}
		yamlBytes, err := protoyaml.Marshal(resp.Config)
		if err != nil {
			log.Fatal(err)
		}
		config, err := ParseReleaseChannelConfig(ctx, inputFileType, input, rcName, string(yamlBytes), noPrompt)
		if err != nil {
			log.Fatal(err)
		}
		if err != nil {
			log.Fatal(err)
		}

		err = configureReleaseChannel(ctx, cmdutil.App, config)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Define a new release channel.",
	Long: `Define a new release channel.

pvnctl release-channels --app=<app> create  # will launch editor
pvnctl release-channels --app=<app> create --input config.yaml
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
		config, err := ParseReleaseChannelConfig(ctx, inputFileType, input, "", defaultConfigYaml, noPrompt)
		if err != nil {
			log.Fatal(err)
		}

		// TODO(naphat) validate that this release channel does not exist
		err = configureReleaseChannel(ctx, cmdutil.App, config)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(configureCmd)
	for _, cmd := range []*cobra.Command{createCmd, editCmd} {
		RootCmd.AddCommand(cmd)
		cmd.Flags().String("input", "", "If set, take release channel config from a file instead of launching an editor. Supported formats: json, yaml, pbtxt")
		cmd.Flags().String("input-file-format", protohelpers.FileTypeInfer, "Used in conjunction with --input. Specify input file type. Pass 'infer' to infer file type from file extension.")
		cmd.Flags().BoolP("no-prompt", "y", false, "Whether to prompt for confirmation before saving.")
	}
}
