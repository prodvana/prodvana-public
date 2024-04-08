package main

import (
	"context"
	go_errors "errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bradenaw/juniper/xslices"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/hashicorp/go-multierror"
	"github.com/pelletier/go-toml/v2"
	"github.com/pkg/errors"
	application_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/application"
	common_config_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config"
	fly_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/fly"
	organization_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/organization"
	release_channel_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/release_channel"
	service_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/service"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
)

var configureFlags = struct {
	flyRuntime string
}{}

type flyProdvanaConfig struct {
	App      string `toml:"app"`
	Prodvana *struct {
		Application string `toml:"application"`
		Group       string `toml:"group"`
	} `toml:"prodvana"`
}

type prodvanaServiceAndApp struct {
	application string
	service     string
}

func configureOneApplication(ctx context.Context, appName string, rcs []*release_channel_pb.ReleaseChannelConfig, perRc []*service_pb.PerReleaseChannelConfig, imageInfos []imageInfo) (*prodvanaServiceAndApp, error) {
	if len(imageInfos) > 1 {
		return nil, errors.Errorf("Inconsistent image used. %+v", imageInfos)
	}
	appCfg := &application_pb.ApplicationConfig{
		Name:            appName,
		ReleaseChannels: rcs,
	}
	svcCfg := &service_pb.ServiceConfig{
		Name:              "fly-app",
		PerReleaseChannel: perRc,
	}
	if len(imageInfos) == 1 {
		svcCfg.Parameters = []*common_config_pb.ParameterDefinition{
			{
				Name:        "image",
				Description: "Docker image to deploy",
				Required:    true,
				ConfigOneof: &common_config_pb.ParameterDefinition_DockerImage{
					DockerImage: &common_config_pb.DockerImageParameterDefinition{
						ImageRegistryInfo: &common_config_pb.ImageRegistryInfo{
							ContainerRegistry: imageInfos[0].registryIntegration,
							ImageRepository:   imageInfos[0].repository,
						},
					},
				},
			},
		}
		svcCfg.BundleName = "{{.Params.image.Tag}}"
	}
	validateResp, err := getApplicationManagerClient().ValidateConfigureApplication(ctx, &application_pb.ConfigureApplicationReq{
		ApplicationConfig: appCfg,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to validate application")
	}

	_, err = getApplicationManagerClient().ConfigureApplication(ctx, &application_pb.ConfigureApplicationReq{
		ApplicationConfig: appCfg,
		ApprovedDangerousActionIds: xslices.Map(validateResp.DangerousActions, func(action *common_config_pb.DangerousAction) string {
			return action.Id
		}),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to configure application")
	}

	_, err = getServiceManagerClient().ConfigureService(ctx, &service_pb.ConfigureServiceReq{
		Application:   appName,
		ServiceConfig: svcCfg,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to configure service")
	}
	return &prodvanaServiceAndApp{
		application: appCfg.Name,
		service:     svcCfg.Name,
	}, nil
}

type patchResult struct {
	patched   []byte
	origImage string
}

type notPatchError struct {
	reason string
}

func (e *notPatchError) Error() string {
	return e.reason
}

func maybePatchFlyToml(bytes []byte) (*patchResult, error) {
	var cfg map[string]interface{}
	if err := toml.Unmarshal(bytes, &cfg); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal Fly configuration")
	}
	build, ok := cfg["build"].(map[string]interface{})
	if !ok {
		return nil, &notPatchError{"no build section"}
	}
	image, ok := build["image"].(string)
	if !ok {
		return nil, &notPatchError{"no build.image not set to a string"}
	}
	build["image"] = "{{.Params.image}}"
	bytes, err := toml.Marshal(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal Fly configuration")
	}
	return &patchResult{
		patched:   bytes,
		origImage: image,
	}, nil
}

// extractImageInfo extracts image registry information from a given image string.
// returns nil if the image info cannot be extracted
func extractImageInfo(image string) *imageInfo {
	parsed, err := name.ParseReference(image)
	if err != nil {
		return nil
	}
	repository := parsed.Context().RepositoryStr()
	if !strings.Contains(repository, "/") {
		repository = "library/" + repository
	}
	return &imageInfo{
		// TODO(naphat) do not hardcode dockerhub
		registryIntegration: "dockerhub-public",
		repository:          repository,
	}
}

type imageInfo struct {
	registryIntegration string
	repository          string
}

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "configure [<root directory>]",
	Long:  `Configure Prodvana by looking for Fly Apps in a given directory. Defaults to the root directory.`,
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		var dir string
		if len(args) == 0 {
			var err error
			dir, err = os.Getwd()
			if err != nil {
				return errors.Wrap(err, "failed to get current working directory")
			}
		} else {
			dir = args[0]
		}
		releaseChannels := map[string][]*release_channel_pb.ReleaseChannelConfig{}
		perReleaseChannelConfigs := map[string][]*service_pb.PerReleaseChannelConfig{}
		imageInfos := map[string]map[imageInfo]struct{}{}
		makeReleaseChannelConfig := func(group, name string) (*release_channel_pb.ReleaseChannelConfig, error) {
			cfg := &release_channel_pb.ReleaseChannelConfig{
				Name:  name,
				Group: group,
				Runtimes: []*release_channel_pb.ReleaseChannelRuntimeConfig{
					{
						Runtime: configureFlags.flyRuntime,
					},
				},
			}
			if group == "production" {
				cfg.Preconditions = []*release_channel_pb.Precondition{
					{
						Precondition: &release_channel_pb.Precondition_ReleaseChannelStable_{
							ReleaseChannelStable: &release_channel_pb.Precondition_ReleaseChannelStable{
								StableOneof: &release_channel_pb.Precondition_ReleaseChannelStable_Selector{
									Selector: "@group=staging",
								},
								AllowEmpty: true,
							},
						},
					},
					{
						Precondition: &release_channel_pb.Precondition_ManualApproval_{
							ManualApproval: &release_channel_pb.Precondition_ManualApproval{},
						},
					},
				}
			}
			return cfg, nil
		}
		err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			if filepath.Ext(path) == ".toml" {
				bytes, err := os.ReadFile(path)
				if err != nil {
					return errors.Wrap(err, "failed to read Fly configuration")
				}
				var cfg flyProdvanaConfig
				if err := toml.Unmarshal(bytes, &cfg); err != nil {
					return errors.Wrap(err, "failed to unmarshal Fly configuration")
				}
				if cfg.App == "" {
					log.Printf("skipping toml %s because it does not look like a fly toml", path)
					return nil
				}
				if cfg.Prodvana == nil {
					log.Printf("skipping fly toml %s because it does not have `prodvana` set", path)
					return nil
				}
				switch cfg.Prodvana.Group {
				case "staging", "production":
				default:
					return errors.Errorf("unknown group %s", cfg.Prodvana.Group)
				}
				patched, err := maybePatchFlyToml(bytes)
				var bytesForCfg []byte
				if err != nil {
					var notPatchError *notPatchError
					if !go_errors.As(err, &notPatchError) {
						return err
					}
					log.Printf("did not patch fly toml %s because %s, image parameters will not impact this Fly App", path, notPatchError.reason)
					bytesForCfg = bytes
				} else {
					bytesForCfg = patched.patched
					info := extractImageInfo(patched.origImage)
					if info != nil {
						if _, ok := imageInfos[cfg.Prodvana.Application]; !ok {
							imageInfos[cfg.Prodvana.Application] = map[imageInfo]struct{}{}
						}
						imageInfos[cfg.Prodvana.Application][*info] = struct{}{}
					}
				}
				rc, err := makeReleaseChannelConfig(cfg.Prodvana.Group, cfg.App)
				if err != nil {
					return err
				}
				releaseChannels[cfg.Prodvana.Application] = append(releaseChannels[cfg.Prodvana.Application], rc)
				perReleaseChannelConfigs[cfg.Prodvana.Application] = append(perReleaseChannelConfigs[cfg.Prodvana.Application], &service_pb.PerReleaseChannelConfig{
					ReleaseChannel: rc.Name,
					ConfigOneof: &service_pb.PerReleaseChannelConfig_Fly{
						Fly: &fly_pb.FlyConfig{
							TomlOneof: &fly_pb.FlyConfig_Inlined{
								Inlined: string(bytesForCfg),
							},
						},
					},
				})
			}
			return nil
		})
		if err != nil {
			return err
		}
		var overallErr error
		orgResp, err := getOrganizationManagerClient().GetOrganization(ctx, &organization_pb.GetOrganizationReq{})
		if err != nil {
			return errors.Wrap(err, "failed to get organization")
		}
		for appName, channels := range releaseChannels {
			configured, err := configureOneApplication(ctx, appName, channels, perReleaseChannelConfigs[appName], maps.Keys(imageInfos[appName]))
			if err != nil {
				overallErr = multierror.Append(overallErr, errors.Wrapf(err, "failed to process Prodvana Application %s", appName))
				continue
			}
			fmt.Printf("Configured Prodvana Application. Kick off a deployment at: %s/applications/%s/services/%s\n", orgResp.Organization.UiAddress, configured.application, configured.service)
		}
		return overallErr
	},
}

func init() {
	configureCmd.Flags().StringVar(&configureFlags.flyRuntime, "fly-runtime", "fly", "Name of the Fly Runtime on Prodvana")
}
