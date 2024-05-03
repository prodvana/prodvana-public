package protections

import (
	"context"
	_ "embed"
	"fmt"
	"log"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	"github.com/prodvana/prodvana-public/go/pkg/protohelpers"
	"github.com/prodvana/prodvana-public/go/pkg/protoyaml"

	common_config_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config"
	protection_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/protection"
	version_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/version"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

//go:embed templates/protection_config.yaml
var defaultConfigYaml string

func ParseProtectionConfig(ctx context.Context, inputFileType, maybeInput, maybeRequiredName, startingYaml string, noPrompt bool, options ...cmdutil.EditOptions[*protection_pb.ProtectionConfig]) (*protection_pb.ProtectionConfig, []*common_config_pb.DangerousAction, error) {
	var config protection_pb.ProtectionConfig
	actions, err := cmdutil.EditOrReadConfig(ctx, inputFileType, maybeInput, startingYaml, &config, func(ctx context.Context, config *protection_pb.ProtectionConfig) ([]*common_config_pb.DangerousAction, error) {
		if config.Name == "" {
			return nil, errors.Errorf("name must be set")
		}
		if maybeRequiredName != "" && config.Name != maybeRequiredName {
			return nil, errors.Errorf("name must be %s", maybeRequiredName)
		}
		_, err := cmdutil.GetProtectionManagerClient().ValidateConfigureProtection(ctx, &protection_pb.ConfigureProtectionReq{
			ProtectionConfig: config,
		})
		if err != nil {
			return nil, errors.Wrapf(err, "validation failed")
		}
		return nil, nil
	}, "Save changes?", noPrompt, options...)
	if err != nil {
		return nil, nil, err
	}
	return &config, actions, nil
}

func configureProtection(ctx context.Context, config *protection_pb.ProtectionConfig, _ []*common_config_pb.DangerousAction) error {
	resp, err := cmdutil.GetProtectionManagerClient().ConfigureProtection(ctx, &protection_pb.ConfigureProtectionReq{
		ProtectionConfig: config,
		Source:           version_pb.Source_INTERACTIVE_PVNCTL,
	})
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", resp.Version)
	return nil
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Define a new protection",
	Long: `Define a new protection

pvnctl protections create  # will launch editor
pvnctl protections create --input config.yaml
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
		config, actions, err := ParseProtectionConfig(ctx, inputFileType, input, "", defaultConfigYaml, noPrompt)
		if err != nil {
			log.Fatal(err)
		}

		// TODO(naphat) validate that this protection does not exist
		err = configureProtection(ctx, config, actions)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit an existing protection.",
	Long: `Edit an existing protection.

pvnctl protections edit <app>  # will launch editor
pvnctl protections edit --input config.yaml <app>
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
		baseVersion, err := cmd.Flags().GetString("base-version")
		if err != nil {
			log.Fatal(err)
		}
		noPrompt, err := cmd.Flags().GetBool("no-prompt")
		if err != nil {
			log.Fatal(err)
		}
		appName := args[0]
		resp, err := cmdutil.GetProtectionManagerClient().GetProtectionConfig(ctx, &protection_pb.GetProtectionConfigReq{
			Protection: appName,
			Version:    baseVersion,
		})
		if err != nil {
			log.Fatal(err)
		}
		yamlBytes, err := protoyaml.Marshal(resp.InputConfig)
		if err != nil {
			log.Fatal(err)
		}
		config, actions, err := ParseProtectionConfig(ctx, inputFileType, input, appName, string(yamlBytes), noPrompt)
		if err != nil {
			log.Fatal(err)
		}

		// TODO(naphat) validate that this protection does not exist
		err = configureProtection(ctx, config, actions)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	for _, cmd := range []*cobra.Command{createCmd, editCmd} {
		RootCmd.AddCommand(cmd)
		cmd.Flags().String("input", "", "If set, take protection config from a file instead of launching an editor. Supported formats: json, yaml, pbtxt")
		cmd.Flags().String("input-file-format", protohelpers.FileTypeInfer, "Used in conjunction with --input. Specify input file type. Pass 'infer' to infer file type from file extension.")
		cmd.Flags().BoolP("no-prompt", "y", false, "Whether to prompt for confirmation before saving.")
	}
	editCmd.Flags().String("base-version", "", "Base protection version to edit from. Defaults to latest.")
}
