package services

import (
	"context"
	_ "embed"
	"fmt"
	"log"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	"github.com/prodvana/prodvana-public/go/pkg/protohelpers"
	"github.com/prodvana/prodvana-public/go/pkg/protoyaml"

	"github.com/prodvana/prodvana-public/go/pkg/perrors"

	common_config_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config"
	ds_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state"
	maestro_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/maestro"
	ds_model_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/model"
	service_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/service"
	version_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/version"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

//go:embed templates/service_config.yaml
var defaultServiceConfigYaml string

func ParseServiceConfig(ctx context.Context, inputFileType, maybeInput, maybeRequiredName, startingYaml string, noPrompt bool, options ...cmdutil.EditOptions[*service_pb.ServiceConfig]) (*service_pb.ServiceConfig, error) {
	var config service_pb.ServiceConfig
	_, err := cmdutil.EditOrReadConfig(ctx, inputFileType, maybeInput, startingYaml, &config, func(ctx context.Context, config *service_pb.ServiceConfig) ([]*common_config_pb.DangerousAction, error) {
		if config.Name == "" {
			return nil, errors.Errorf("name must be set")
		}
		if maybeRequiredName != "" && config.Name != maybeRequiredName {
			return nil, errors.Errorf("name must be %s", maybeRequiredName)
		}
		config = proto.Clone(config).(*service_pb.ServiceConfig)
		if cmdutil.App != "" {
			config.Application = cmdutil.App
		}
		_, err := cmdutil.GetServiceManagerClient().ValidateApplyParameters(ctx, &service_pb.ApplyParametersReq{
			Oneof: &service_pb.ApplyParametersReq_ServiceConfig{
				ServiceConfig: config,
			},
		})
		if err != nil {
			return nil, errors.Wrapf(err, "validation failed")
		}
		return nil, nil
	}, "Save changes and push?", noPrompt, options...)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

type pushResp struct {
	desiredStatePush *ds_pb.SetDesiredStateResp
}

func pushService(ctx context.Context, app, service, version string, overwriteMaestroMode bool) (*pushResp, error) {
	var currentDs *ds_model_pb.ServiceState
	latest, err := cmdutil.GetDesiredStateManagerClient().GetServiceLatestCombinedReleaseDesiredState(ctx, &ds_pb.GetServiceLatestCombinedReleaseDesiredStateReq{
		Application: app,
		Service:     service,
	})
	if err != nil {
		if perrors.GrpcErrCode(err) != codes.FailedPrecondition {
			return nil, errors.Wrap(err, "failed to query for latest release")
		}
		currentDs = &ds_model_pb.ServiceState{
			Application: app,
			Service:     service,
		}
	} else {
		currentDs = latest.InputDesiredState.GetService()
	}
	currentDs.ReleaseChannels = nil
	if len(currentDs.ReleaseChannelLabelSelectors) == 0 {
		currentDs.ReleaseChannelLabelSelectors = append(currentDs.ReleaseChannelLabelSelectors, &ds_model_pb.ServiceInstanceLabelSelector{
			SelectorOneof: &ds_model_pb.ServiceInstanceLabelSelector_All{
				All: true,
			},
			Versions: []*ds_model_pb.Version{
				{
					Version: version,
				},
			},
			AutorollbackOneof: &ds_model_pb.ServiceInstanceLabelSelector_ComputeRollbackVersion{
				ComputeRollbackVersion: true,
			},
		})
	} else {
		currentDs.ReleaseChannelLabelSelectors[0].Versions = []*ds_model_pb.Version{
			{
				Version: version,
			},
		}
	}
	dsReq := &ds_pb.SetDesiredStateReq{
		DesiredState: &ds_model_pb.State{
			StateOneof: &ds_model_pb.State_Service{
				Service: currentDs,
			},
		},
		Source:                    version_pb.Source_INTERACTIVE_PVNCTL,
		ForceAsyncSetDesiredState: cmdutil.ForceAsyncSetDesiredState,
	}
	if overwriteMaestroMode {
		dsReq.MaestroConfigOverride = &maestro_pb.MaestroConfig{
			Strategy: maestro_pb.Strategy_IMMEDIATE,
		}
	}

	dsResp, err := cmdutil.GetDesiredStateManagerClient().SetDesiredState(ctx, dsReq)
	if err != nil {
		return nil, errors.Wrap(err, "failed to set desired state")
	}
	return &pushResp{
		desiredStatePush: dsResp,
	}, nil
}

func PushService(ctx context.Context, app, service, version string) error {
	_, err := pushService(ctx, app, service, version, false)
	return err
}

func configureAndPushService(ctx context.Context, config *service_pb.ServiceConfig, noPush bool) (*pushResp, error) {
	resp, err := cmdutil.GetServiceManagerClient().ApplyParameters(ctx, &service_pb.ApplyParametersReq{
		Oneof: &service_pb.ApplyParametersReq_ServiceConfig{
			ServiceConfig: config,
		},
		Application: cmdutil.App,
		Source:      version_pb.Source_INTERACTIVE_PVNCTL,
	})
	if err != nil {
		return nil, err
	}

	if noPush {
		// HACK(naphat) this output is doubled as a scripting output
		fmt.Printf("%s\n", resp.Version)
		return nil, nil
	}
	pushResp, err := pushService(ctx, cmdutil.App, config.Name, resp.Version, false)
	if err != nil {
		return nil, err
	}
	switch inner := pushResp.desiredStatePush.IdOneof.(type) {
	case *ds_pb.SetDesiredStateResp_DesiredStateId:
		fmt.Printf("Converging desired state %s. To follow along, run:\n", inner.DesiredStateId)
		fmt.Printf("pvnctl services --app %s wait-desired-state %s\n", cmdutil.App, config.Name)
	case *ds_pb.SetDesiredStateResp_ReleaseId:
		// TODO(naphat) implement the equivalent of wait-desired-state for releases
		fmt.Printf("Enqueued Release %s.\n", inner.ReleaseId)
	default:
		return nil, errors.Errorf("unexpected desired state id type %T", inner)
	}
	return pushResp, nil
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Define a new service",
	Long: `Define a new service

pvnctl services --app=<app> create  # will launch editor
pvnctl services --app=<app> create --input service-config.yaml
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
		noPush, err := cmd.Flags().GetBool("no-push")
		if err != nil {
			log.Fatal(err)
		}
		noPrompt, err := cmd.Flags().GetBool("no-prompt")
		if err != nil {
			log.Fatal(err)
		}
		config, err := ParseServiceConfig(ctx, inputFileType, input, "", defaultServiceConfigYaml, noPrompt)
		if err != nil {
			log.Fatal(err)
		}

		// TODO(naphat) validate that this service does not exist
		_, err = configureAndPushService(ctx, config, noPush)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push an existing service.",
	Long: `Push an existing service.

pvnctl services --app=<app> push <service>  # will launch editor
pvnctl services --app=<app> push --input service-config.yaml <service>
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
		noPush, err := cmd.Flags().GetBool("no-push")
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
		serviceName := args[0]
		resp, err := cmdutil.GetServiceManagerClient().GetMaterializedConfig(ctx, &service_pb.GetMaterializedConfigReq{
			Service:     serviceName,
			Version:     baseVersion,
			Application: cmdutil.App,
		})
		if err != nil {
			log.Fatal(err)
		}
		yamlBytes, err := protoyaml.Marshal(resp.Config)
		if err != nil {
			log.Fatal(err)
		}
		config, err := ParseServiceConfig(ctx, inputFileType, input, serviceName, string(yamlBytes), noPrompt)
		if err != nil {
			log.Fatal(err)
		}
		_, err = configureAndPushService(ctx, config, noPush)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	for _, cmd := range []*cobra.Command{createCmd, pushCmd} {
		RootCmd.AddCommand(cmd)
		cmd.Flags().String("input", "", "If set, take service config from a file instead of launching an editor. Supported formats: json, yaml, pbtxt")
		cmd.Flags().String("input-file-format", protohelpers.FileTypeInfer, "Used in conjunction with --input. Specify input file type. Pass 'infer' to infer file type from file extension.")
		cmd.Flags().Bool("no-push", false, "Configure service only, do not push.")
		cmd.Flags().BoolP("no-prompt", "y", false, "Whether to prompt for confirmation before pushing.")
	}
	pushCmd.Flags().String("base-version", "", "Base service version to edit from. Defaults to latest.")
}
