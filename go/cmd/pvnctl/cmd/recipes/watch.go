package recipes

import (
	"context"
	"io"
	"log"
	"sort"
	"sync"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"
	"github.com/prodvana/prodvana-public/go/pkg/sets"
	"golang.org/x/sync/errgroup"

	ds_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state"
	ds_model_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/model"
	recipe_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/recipe"

	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch status of a recipe",
	Long: `Watch status of a recipe

pvnctl recipes watch
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		recipeName := args[0]
		recipeResp, err := cmdutil.GetRecipeManagerClient().GetRecipe(ctx, &recipe_pb.GetRecipeReq{
			Recipe: recipeName,
		})
		if err != nil {
			log.Fatal(err)
		}

		var services []string

		for _, stage := range recipeResp.Recipe.Config.Stages {
			for _, target := range stage.Targets {
				switch typed := target.Target.(type) {
				case *recipe_pb.Target_Service:
					services = append(services, typed.Service.Service)
				}
			}
		}

		err = cmdutil.RenderWatch(ctx, func(ctx context.Context, writer io.Writer) (bool, error) {
			releaseChannels := sets.NewSet[string]()

			type rcStatusType struct {
				status ds_model_pb.Status
			}
			type svcStatusType struct {
				statuses map[string]rcStatusType
				name     string
			}
			var svcStatusesLock sync.Mutex
			svcStatuses := make(map[string]svcStatusType)

			eg, ctx := errgroup.WithContext(ctx)

			for _, stage := range recipeResp.Recipe.Config.Stages {
				for _, target := range stage.Targets {
					switch typed := target.Target.(type) {
					case *recipe_pb.Target_Service:
						eg.Go(func() error {
							svcStatus := svcStatusType{
								statuses: map[string]rcStatusType{},
								name:     typed.Service.Service,
							}
							dsIdResp, err := cmdutil.GetDesiredStateManagerClient().GetServiceLatestDesiredStateId(ctx, &ds_pb.GetServiceLatestDesiredStateIdReq{
								Application: typed.Service.Application,
								Service:     typed.Service.Service,
							})
							if err != nil {
								log.Fatal(err)
							}

							if dsIdResp.DesiredStateId == "" {
								log.Printf("No desired state found for service %s/%s\n", typed.Service.Application, typed.Service.Service)
								return nil
							}

							dsResp, err := cmdutil.GetDesiredStateManagerClient().GetDesiredStateGraph(ctx, &ds_pb.GetDesiredStateGraphReq{
								Query: &ds_pb.GetDesiredStateGraphReq_DesiredStateId{
									DesiredStateId: dsIdResp.DesiredStateId,
								},
								Types:                 []ds_model_pb.Type{ds_model_pb.Type_SERVICE_INSTANCE},
								ExcludeDependencyType: []ds_model_pb.DependencyType{ds_model_pb.DependencyType_DEPENDENCY_CHILD, ds_model_pb.DependencyType_DEPENDENCY_PRECONDITION, ds_model_pb.DependencyType_DEPENDENCY_SERVICE},
								Depth:                 0,
							})
							if err != nil {
								log.Fatal(err)
							}

							for _, entity := range dsResp.EntityGraph.Entities {
								if entity.Id.Type == ds_model_pb.Type_SERVICE_INSTANCE && entity.RootDesiredStateId == dsIdResp.DesiredStateId {
									rcId := entity.DesiredState.StateOneof.(*ds_model_pb.State_ServiceInstance).ServiceInstance.ReleaseChannel
									svcStatus.statuses[rcId] = rcStatusType{
										status: entity.Status,
									}
									svcStatusesLock.Lock()
									releaseChannels.Add(rcId)
									svcStatusesLock.Unlock()
								}
							}

							svcStatusesLock.Lock()
							svcStatuses[typed.Service.Service] = svcStatus
							svcStatusesLock.Unlock()

							return nil
						})
					}
				}
			}

			if err := eg.Wait(); err != nil {
				return false, err
			}
			rcs := releaseChannels.AsSlice()
			sort.Strings(rcs)

			outputListing := cmdutil.NewOutputListing(append([]string{"SERVICE"}, rcs...)...)
			for _, svc := range services {
				svcStatus := svcStatuses[svc]

				var row []interface{}
				row = append(row, svcStatus.name)
				for _, rc := range rcs {
					row = append(row, svcStatus.statuses[rc].status.String())
				}
				outputListing.AddRow(row...)
			}
			return false, outputListing.WriteTable(writer)
		})

		log.Fatal(err)
	},
}

func init() {
	RootCmd.AddCommand(watchCmd)
}
