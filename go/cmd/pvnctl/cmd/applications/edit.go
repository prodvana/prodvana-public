package applications

import (
	"context"
	_ "embed"
	"fmt"
	"log"

	"prodvana/cmd/cmdutil"

	"github.com/prodvana/prodvana-public/go/pkg/protohelpers"
	"github.com/prodvana/prodvana-public/go/pkg/protoyaml"

	application_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/application"
	common_config_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config"
	version_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/version"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

//go:embed templates/application_config.yaml
var defaultConfigYaml string

func ParseApplicationConfig(ctx context.Context, inputFileType, maybeInput, maybeRequiredName, startingYaml string, noPrompt bool, options ...cmdutil.EditOptions[*application_pb.ApplicationConfig]) (*application_pb.ApplicationConfig, []*common_config_pb.DangerousAction, error) {
	var config application_pb.ApplicationConfig
	actions, err := cmdutil.EditOrReadConfig(ctx, inputFileType, maybeInput, startingYaml, &config, func(ctx context.Context, config *application_pb.ApplicationConfig) ([]*common_config_pb.DangerousAction, error) {
		if config.Name == "" {
			return nil, errors.Errorf("name must be set")
		}
		if maybeRequiredName != "" && config.Name != maybeRequiredName {
			return nil, errors.Errorf("name must be %s", maybeRequiredName)
		}
		resp, err := cmdutil.GetApplicationManagerClient().ValidateConfigureApplication(ctx, &application_pb.ConfigureApplicationReq{
			ApplicationConfig: config,
		})
		if err != nil {
			return nil, errors.Wrapf(err, "validation failed")
		}
		return resp.DangerousActions, nil
	}, "Save changes?", noPrompt, options...)
	if err != nil {
		return nil, nil, err
	}
	return &config, actions, nil
}

func configureApplication(ctx context.Context, fetchedVersion string, config *application_pb.ApplicationConfig, approvedActions []*common_config_pb.DangerousAction) error {
	var approvedIds []string
	for _, action := range approvedActions {
		approvedIds = append(approvedIds, action.Id)
	}
	resp, err := cmdutil.GetApplicationManagerClient().ConfigureApplication(ctx, &application_pb.ConfigureApplicationReq{
		ApplicationConfig:          config,
		ApprovedDangerousActionIds: approvedIds,
		Source:                     version_pb.Source_INTERACTIVE_PVNCTL,
		BaseVersion:                fetchedVersion,
	})
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", resp.Meta.Version)
	return nil
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Define a new application",
	Long: `Define a new application

pvnctl applications create  # will launch editor
pvnctl applications create --input config.yaml
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
		config, actions, err := ParseApplicationConfig(ctx, inputFileType, input, "", defaultConfigYaml, noPrompt)
		if err != nil {
			log.Fatal(err)
		}

		// TODO(naphat) validate that this application does not exist
		err = configureApplication(ctx, "", config, actions)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit an existing application.",
	Long: `Edit an existing application.

pvnctl applications edit <app>  # will launch editor
pvnctl applications edit --input config.yaml <app>
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
		resp, err := cmdutil.GetApplicationManagerClient().GetApplicationConfig(ctx, &application_pb.GetApplicationConfigReq{
			Application: appName,
			Version:     baseVersion,
		})
		if err != nil {
			log.Fatal(err)
		}
		yamlBytes, err := protoyaml.Marshal(resp.Config)
		if err != nil {
			log.Fatal(err)
		}
		config, actions, err := ParseApplicationConfig(ctx, inputFileType, input, appName, string(yamlBytes), noPrompt)
		if err != nil {
			log.Fatal(err)
		}

		var fetchedVersion string
		// if the user explicitly asks to edit an older version of the application,
		// then we don't need to ensure we are overriding the latest config version
		// so only pass the baseVersion if we fetched the latest from GetApplicationConfig
		if baseVersion == "" {
			fetchedVersion = resp.Version
		}

		// TODO(naphat) validate that this application does not exist
		err = configureApplication(ctx, fetchedVersion, config, actions)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	for _, cmd := range []*cobra.Command{createCmd, editCmd} {
		RootCmd.AddCommand(cmd)
		cmd.Flags().String("input", "", "If set, take application config from a file instead of launching an editor. Supported formats: json, yaml, pbtxt")
		cmd.Flags().String("input-file-format", protohelpers.FileTypeInfer, "Used in conjunction with --input. Specify input file type. Pass 'infer' to infer file type from file extension.")
		cmd.Flags().BoolP("no-prompt", "y", false, "Whether to prompt for confirmation before saving.")
	}
	editCmd.Flags().String("base-version", "", "Base application version to edit from. Defaults to latest.")
}
