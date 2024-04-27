package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/bradenaw/juniper/xslices"
	"github.com/pelletier/go-toml/v2"
	"github.com/pkg/errors"
	application_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/application"
	common_config_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config"
	ds_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state"
	ds_model_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/model"
	fly_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/fly"
	organization_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/organization"
	release_channel_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/release_channel"
	service_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/service"
	"github.com/spf13/cobra"
)

var deployFlags = struct {
	flyToml string
}{}

type appAndServiceCfg struct {
	appCfg        *application_pb.ApplicationConfig
	svcCfg        *service_pb.ServiceConfig
	paramValues   []*common_config_pb.ParameterValue
	requiresBuild bool
}

func getTier(tier *string) string {
	if tier == nil {
		return "production"
	}
	return *tier
}

func constructAppAndServiceCfg(tomlBytes []byte, tomlCfg *flyToml, defaultTagForBuild string) (*appAndServiceCfg, error) {
	patched, err := maybePatchFlyToml(tomlBytes)
	if err != nil {
		return nil, err
	}
	var registryIntegration *common_config_pb.ImageRegistryInfo
	var tag string
	if patched.requiresBuild {
		registryIntegration = &common_config_pb.ImageRegistryInfo{
			ContainerRegistry: configureFlags.flyRegistry,
			ImageRepository:   tomlCfg.App,
		}
		tag = defaultTagForBuild
	} else {
		var info *imageInfo
		info, tag, err = extractImageInfo(patched.origImage)
		if err != nil {
			return nil, err
		}
		if info != nil {
			registryIntegration = &common_config_pb.ImageRegistryInfo{
				ContainerRegistry: info.registryIntegration,
				ImageRepository:   info.repository,
			}
		}
	}
	var bundleName string
	var params []*common_config_pb.ParameterDefinition
	var values []*common_config_pb.ParameterValue
	if registryIntegration != nil {
		params = []*common_config_pb.ParameterDefinition{
			{
				Name:        "image",
				Description: "Docker image to deploy",
				Required:    true,
				ConfigOneof: &common_config_pb.ParameterDefinition_DockerImage{
					DockerImage: &common_config_pb.DockerImageParameterDefinition{
						ImageRegistryInfo: registryIntegration,
					},
				},
			},
		}
		values = []*common_config_pb.ParameterValue{
			{
				Name: "image",
				ValueOneof: &common_config_pb.ParameterValue_DockerImageTag{
					DockerImageTag: tag,
				},
			},
		}
		bundleName = "{{.Params.image.Tag}}"
	}
	if tomlCfg.Replicas == nil {
		appCfg := &application_pb.ApplicationConfig{
			Name: tomlCfg.App,
			ReleaseChannels: []*release_channel_pb.ReleaseChannelConfig{
				{
					Name: "all-regions",
					Runtimes: []*release_channel_pb.ReleaseChannelRuntimeConfig{
						{
							Runtime: configureFlags.flyRuntime,
						},
					},
				},
			},
		}
		svcCfg := &service_pb.ServiceConfig{
			Name:       tomlCfg.App,
			BundleName: bundleName,
			ConfigOneof: &service_pb.ServiceConfig_Fly{
				Fly: &fly_pb.FlyConfig{
					TomlOneof: &fly_pb.FlyConfig_Inlined{
						Inlined: string(patched.patched),
					},
				},
			},
			Parameters: params,
		}
		return &appAndServiceCfg{
			appCfg:        appCfg,
			svcCfg:        svcCfg,
			paramValues:   values,
			requiresBuild: patched.requiresBuild,
		}, nil
	}

	type rcInputConfig struct {
		region   string
		replicas int
		tier     string
	}
	rcInputs := []*rcInputConfig{
		{
			region:   tomlCfg.PrimaryRegion,
			replicas: *tomlCfg.Replicas,
			tier:     getTier(tomlCfg.Tier),
		},
	}
	for region, rc := range tomlCfg.Regions {
		if rc.Replicas == nil {
			return nil, errors.Errorf("replicas not specified for region %s", region)
		}
		rcInputs = append(rcInputs, &rcInputConfig{
			region:   region,
			replicas: *rc.Replicas,
			tier:     getTier(rc.Tier),
		})
	}
	hasMultiStage := slices.ContainsFunc(rcInputs, func(rc *rcInputConfig) bool {
		return rc.tier == "canary"
	})
	const regionConstantName = "region"
	releaseChannels := make([]*release_channel_pb.ReleaseChannelConfig, 0, len(rcInputs))
	for _, rcInput := range rcInputs {
		rcCfg := &release_channel_pb.ReleaseChannelConfig{
			Name:  rcInput.region,
			Group: rcInput.tier,
			Runtimes: []*release_channel_pb.ReleaseChannelRuntimeConfig{
				{
					Runtime: configureFlags.flyRuntime,
				},
			},
			Constants: []*common_config_pb.Constant{
				{
					Name: regionConstantName,
					ConfigOneof: &common_config_pb.Constant_String_{
						String_: &common_config_pb.StringConstant{
							Value: rcInput.region,
						},
					},
				},
			},
		}
		if rcInput.tier == "production" && hasMultiStage {
			rcCfg.Preconditions = []*release_channel_pb.Precondition{
				{
					Precondition: &release_channel_pb.Precondition_ReleaseChannelStable_{
						ReleaseChannelStable: &release_channel_pb.Precondition_ReleaseChannelStable{
							StableOneof: &release_channel_pb.Precondition_ReleaseChannelStable_Selector{
								Selector: "@group=canary",
							},
							AllowEmpty: true,
						},
					},
				},
				{
					Precondition: &release_channel_pb.Precondition_ManualApproval_{},
				},
			}
		}
		releaseChannels = append(releaseChannels, rcCfg)
	}
	appCfg := &application_pb.ApplicationConfig{
		Name:            tomlCfg.App,
		ReleaseChannels: releaseChannels,
	}
	svcCfg := &service_pb.ServiceConfig{
		Name:       tomlCfg.App,
		BundleName: bundleName,
		ConfigOneof: &service_pb.ServiceConfig_Fly{
			Fly: &fly_pb.FlyConfig{
				TomlOneof: &fly_pb.FlyConfig_Inlined{
					Inlined: string(patched.patched),
				},
				Regions: []string{"{{.Constants.region}}"},
			},
		},
		Parameters: params,
	}
	return &appAndServiceCfg{
		appCfg:        appCfg,
		svcCfg:        svcCfg,
		paramValues:   values,
		requiresBuild: patched.requiresBuild,
	}, nil
}

func configureAndDeploy(ctx context.Context, cfgs *appAndServiceCfg) error {
	validateResp, err := getApplicationManagerClient().ValidateConfigureApplication(ctx, &application_pb.ConfigureApplicationReq{
		ApplicationConfig: cfgs.appCfg,
	})
	if err != nil {
		return errors.Wrap(err, "failed to validate application")
	}

	_, err = getApplicationManagerClient().ConfigureApplication(ctx, &application_pb.ConfigureApplicationReq{
		ApplicationConfig: cfgs.appCfg,
		ApprovedDangerousActionIds: xslices.Map(validateResp.DangerousActions, func(action *common_config_pb.DangerousAction) string {
			return action.Id
		}),
	})
	if err != nil {
		return errors.Wrap(err, "failed to configure application")
	}

	svcConfigureResp, err := getServiceManagerClient().ConfigureService(ctx, &service_pb.ConfigureServiceReq{
		Application:   cfgs.appCfg.Name,
		ServiceConfig: cfgs.svcCfg,
	})
	if err != nil {
		return errors.Wrap(err, "failed to configure service")
	}
	applyResp, err := getServiceManagerClient().ApplyParameters(ctx, &service_pb.ApplyParametersReq{
		Oneof: &service_pb.ApplyParametersReq_ServiceConfigVersion{
			ServiceConfigVersion: &service_pb.ServiceConfigVersionReference{
				Application:          cfgs.appCfg.Name,
				Service:              cfgs.svcCfg.Name,
				ServiceConfigVersion: svcConfigureResp.ConfigVersion,
			},
		},
		Parameters:                cfgs.paramValues,
		HandleBundleNameDuplicate: true,
	})
	if err != nil {
		return errors.Wrap(err, "failed to apply parameters")
	}
	_, err = getDesiredStateManagerClient().SetDesiredState(ctx, &ds_pb.SetDesiredStateReq{
		DesiredState: &ds_model_pb.State{
			StateOneof: &ds_model_pb.State_Service{
				Service: &ds_model_pb.ServiceState{
					Application: cfgs.appCfg.Name,
					Service:     cfgs.svcCfg.Name,
					ReleaseChannelLabelSelectors: []*ds_model_pb.ServiceInstanceLabelSelector{
						{
							SelectorOneof: &ds_model_pb.ServiceInstanceLabelSelector_All{
								All: true,
							},
							Versions: []*ds_model_pb.Version{
								{
									Version: applyResp.Version,
								},
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		return errors.Wrap(err, "failed to set desired state")
	}
	return nil
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploy [<working directory>]",
	Long:  `Deploy a fly app`,
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		var workingDir string
		if len(args) > 0 {
			workingDir = args[0]
		} else {
			var err error
			workingDir, err = os.Getwd()
			if err != nil {
				return errors.Wrap(err, "failed to get working directory")
			}
		}
		flyTomlPath := deployFlags.flyToml
		if flyTomlPath == "" {
			flyTomlPath = filepath.Join(workingDir, "fly.toml")
		}
		tomlBytes, err := os.ReadFile(flyTomlPath)
		if err != nil {
			if os.IsNotExist(err) {
				return fmt.Errorf("fly configuration not found: %s", flyTomlPath)
			}
			return errors.Wrapf(err, "failed to read fly configuration at %s", flyTomlPath)
		}
		var tomlCfg flyToml
		if err := toml.Unmarshal(tomlBytes, &tomlCfg); err != nil {
			return errors.Wrapf(err, "failed to parse fly configuration at %s", flyTomlPath)
		}

		err = setupFlyIntegrations(ctx, configureFlags.flyOrg, configureFlags.flyRuntime, configureFlags.flyRegistry)
		if err != nil {
			return err
		}
		defaultTag := makeDefaultTag()

		cfgs, err := constructAppAndServiceCfg(tomlBytes, &tomlCfg, defaultTag)
		if err != nil {
			return err
		}
		if cfgs.requiresBuild {
			if err := buildApp(ctx, flyTomlPath, defaultTag, false); err != nil {
				return err
			}
		}
		orgResp, err := getOrganizationManagerClient().GetOrganization(ctx, &organization_pb.GetOrganizationReq{})
		if err != nil {
			return errors.Wrap(err, "failed to get organization")
		}
		if err := configureAndDeploy(ctx, cfgs); err != nil {
			return err
		}
		fmt.Printf("Deployed %[2]s/%[3]s. Follow along at: %[1]s/applications/%[2]s/services/%[3]s\n", orgResp.Organization.UiAddress, cfgs.appCfg.Name, cfgs.svcCfg.Name)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
	deployCmd.Flags().StringVar(&configureFlags.flyRuntime, "fly-runtime", "fly", "Name of the Fly Runtime on Prodvana")
	deployCmd.Flags().StringVar(&configureFlags.flyRegistry, "fly-registry", "fly-registry", "Name of the Fly registry integration on Prodvana")
	deployCmd.Flags().StringVar(&configureFlags.flyctlPath, "flyctl-path", "fly", "Path to flyctl binary")
	deployCmd.Flags().StringVar(&configureFlags.flyOrg, "fly-org", "", "Name of the Fly organization.")
	deployCmd.Flags().StringVarP(&deployFlags.flyToml, "config", "c", "", "Path to fly toml, defaults to fly.toml inside the working directory.")
	// TODO(naphat) we can infer this?
	must(deployCmd.MarkFlagRequired("fly-org"))
}
