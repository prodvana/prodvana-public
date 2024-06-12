package recipes

import (
	"context"
	"log"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"
	ds_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state"
	maestro_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/maestro"
	ds_model_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/model"
	recipe_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/recipe"
	release_channel_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/release_channel"
	version_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/version"

	"github.com/spf13/cobra"
)

var (
	paramFlag []string
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply a recipe",
	Long: `Apply a recipe

pvnctl recipes apply <recipe-name> --param <name=value>
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		recipeName := args[0]
		applyResp, err := cmdutil.GetRecipeManagerClient().ApplyRecipeParameters(ctx, &recipe_pb.ApplyRecipeParametersReq{
			Recipe: recipeName,
		})
		if err != nil {
			log.Fatal(err)
		}

		recipeCfg, err := cmdutil.GetRecipeManagerClient().GetRecipe(ctx, &recipe_pb.GetRecipeReq{
			Recipe: recipeName,
		})
		if err != nil {
			log.Fatal(err)
		}

		svcIndex := make(map[string]*recipe_pb.ApplyRecipeParametersResp_ServiceVersion)
		for _, version := range applyResp.Versions {
			svcIndex[version.ServiceId] = version
		}

		var deps, nextDeps map[string][]*ds_model_pb.Condition
		for _, svc := range recipeCfg.Recipe.Config.Stages {
			deps = nextDeps
			nextDeps = make(map[string][]*ds_model_pb.Condition)

			for _, target := range svc.Targets {
				switch typed := target.Target.(type) {
				case *recipe_pb.Target_Service:
					version := svcIndex[typed.Service.ServiceId]

					appCfg, err := cmdutil.GetReleaseChannelManagerClient().ListReleaseChannelsV2(ctx, &release_channel_pb.ListReleaseChannelsReq{
						Application: version.Application,
					})
					if err != nil {
						log.Fatal(err)
					}
					var rcs []*ds_model_pb.ServiceInstanceState

					for _, rc := range appCfg.ReleaseChannels {
						rcs = append(rcs, &ds_model_pb.ServiceInstanceState{
							ReleaseChannel: rc.Meta.Name,
							Versions: []*ds_model_pb.Version{
								{
									Version: version.ServiceVersion,
								},
							},
							Meta: &ds_model_pb.Metadata{
								Preconditions: deps[rc.Meta.Name],
							},
						})

						nextDeps[rc.Meta.Name] = append(nextDeps[rc.Meta.Name], &ds_model_pb.Condition{
							Condition: &ds_model_pb.Condition_ServiceStable{
								ServiceStable: &ds_model_pb.Condition_ServiceStableCondition{
									Application:      version.Application,
									Service:          version.Service,
									ServiceId:        version.ServiceId,
									ServiceVersion:   version.ServiceVersion,
									ReleaseChannel:   rc.Meta.Name,
									ReleaseChannelId: rc.Meta.Id,
								},
							},
						})
					}

					if _, err := cmdutil.GetDesiredStateManagerClient().SetDesiredState(ctx, &ds_pb.SetDesiredStateReq{
						DesiredState: &ds_model_pb.State{
							StateOneof: &ds_model_pb.State_Service{
								Service: &ds_model_pb.ServiceState{
									Application:     version.Application,
									Service:         version.Service,
									ServiceId:       version.ServiceId,
									ReleaseChannels: rcs,
								},
							},
						},
						Source:                    version_pb.Source_INTERACTIVE_PVNCTL,
						ForceAsyncSetDesiredState: true,
						MaestroConfigOverride: &maestro_pb.MaestroConfig{
							Strategy: maestro_pb.Strategy_MAESTRO_DISABLED,
						},
					}); err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(applyCmd)
	applyCmd.Flags().StringSliceVar(&paramFlag, "param", nil, "Parameters in the name=value format.")
}
