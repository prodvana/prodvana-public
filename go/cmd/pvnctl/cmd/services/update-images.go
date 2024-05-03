package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/prodvana/prodvana-public/go/pkg/containerregistry/imageutils"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	"github.com/prodvana/prodvana-public/go/pkg/sets"

	"github.com/prodvana/prodvana-public/go/pkg/perrors"

	common_config_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config"
	"github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state"
	ds_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state"
	"github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/maestro"
	ds_model_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/model"
	release_channel_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/release_channel"
	service_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/service"
	"golang.org/x/exp/maps"
	"google.golang.org/grpc/codes"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	paramFlag                 []string
	paramByReleaseChannelFlag []string
	maestroOverwriteMode      bool
	successIfReplacedFlag     bool
)

type rcParamHolder struct {
	globalHolder *paramHolder
	paramsByRC   map[string]*paramHolder
}

func makeRcParamHolder(paramStrs []string) (map[string]*paramHolder, error) {
	if len(paramStrs) == 0 {
		return nil, fmt.Errorf("No parameters to update")
	}

	paramsByRC := make(map[string]*paramHolder)
	for _, imageStr := range paramStrs {
		components := strings.SplitN(imageStr, "=", 2)
		if len(components) != 2 {
			return nil, fmt.Errorf("--images-by-release-channel must be in <release_channel>=<image> format. Invalid format: %s", imageStr)
		}
		rc := components[0]
		holder, err := makeParamHolder(strings.Split(components[1], "|"))
		if err != nil {
			return nil, err
		}
		paramsByRC[rc] = holder
	}
	return paramsByRC, nil
}

type paramHolder struct {
	singleParamValue     string
	maybeSingleParamName string
	single               bool
	paramValues          map[string]string
}

func (i *paramHolder) GetParamValue(paramName string) (string, bool) {
	if i.single {
		return i.singleParamValue, true
	}

	image, ok := i.paramValues[paramName]
	return image, ok
}
func (i *paramHolder) GetParamValues() map[string]string {
	if i.single {
		return map[string]string{
			i.maybeSingleParamName: i.singleParamValue,
		}
	}
	return i.paramValues
}

func (i *paramHolder) Count() int {
	if i.single {
		return 1
	}
	return len(i.paramValues)
}

func makeParamHolder(paramStrs []string) (*paramHolder, error) {
	if len(paramStrs) == 0 {
		return nil, fmt.Errorf("No parameters to update")
	} else if len(paramStrs) == 1 {
		components := strings.SplitN(paramStrs[0], "=", 2)
		if len(components) == 2 {
			return &paramHolder{
				maybeSingleParamName: components[0],
				singleParamValue:     components[1],
				single:               true,
			}, nil
		}
		return &paramHolder{
			singleParamValue: paramStrs[0],
			single:           true,
		}, nil
	}

	imageMap := make(map[string]string)
	for _, imageStr := range paramStrs {
		components := strings.SplitN(imageStr, "=", 2)
		if len(components) != 2 {
			return nil, fmt.Errorf("--image must be in <program_name>=<image> format. Invalid format: %s", imageStr)
		}
		program := components[0]
		image := components[1]
		imageMap[program] = image
	}

	return &paramHolder{
		paramValues: imageMap,
	}, nil
}

func waitSetDesiredStateResp(ctx context.Context, dsResp *ds_pb.SetDesiredStateResp, waitChannels []string) error {
	for {
		switch inner := dsResp.GetIdOneof().(type) {
		case *ds_pb.SetDesiredStateResp_DesiredStateId:
			resp, err := cmdutil.GetDesiredStateManagerClient().GetDesiredStateConvergenceSummary(ctx, &desired_state.GetDesiredStateConvergenceReq{
				DesiredStateId:         inner.DesiredStateId,
				FastNoDeprecatedFields: true,
			})
			if err != nil {
				return errors.Wrap(err, "failed to get desired state summary")
			}
			done, err := isConvergenceCompleted(resp.Summary, waitChannels)
			if err != nil {
				return err
			}
			if done {
				return nil
			}
		case *ds_pb.SetDesiredStateResp_ReleaseId:
			relResp, err := cmdutil.GetDesiredStateManagerClient().GetMaestroRelease(ctx, &ds_pb.GetMaestroReleaseReq{
				ReleaseId: inner.ReleaseId,
			})
			if err != nil {
				return errors.Wrap(err, "failed to get release")
			}
			requiredRc := sets.FromSlice(waitChannels)
			rcNameToId := make(map[string]string, len(waitChannels))
			for _, rc := range relResp.MaestroRelease.Config.CompiledDesiredState.GetService().ReleaseChannels {
				if requiredRc.Contains(rc.ReleaseChannel) {
					requiredRc.Remove(rc.ReleaseChannel)
					rcNameToId[rc.ReleaseChannel] = rc.ReleaseChannelId
				}
			}
			if requiredRc.Len() > 0 {
				return errors.Errorf("release %s does not have release channels %s", inner.ReleaseId, requiredRc.AsSlice())
			}
			allDone := true
			for rc, rcId := range rcNameToId {
				status := relResp.MaestroRelease.State.GetReleaseChannelStatuses()[rcId]
				switch status {
				case maestro.Status_SUCCEEDED, maestro.Status_SKIPPED: // do nothing
				case maestro.Status_FAILED:
					return errors.Errorf("release %s failed to promote to release channel %s", inner.ReleaseId, rc)
				default:
					allDone = false
				}
			}
			if allDone {
				return nil
			}
		default:
			return errors.Errorf("invalid SetDesiredState response: %T", inner)
		}
		time.Sleep(1 * time.Second)
	}
}

func isConvergenceCompleted(summary *desired_state.DesiredStateSummary, releaseChannels []string) (bool, error) {
	if summary.EntityGraph == nil {
		switch summary.Status {
		case ds_model_pb.Status_PENDING_SET_DESIRED_STATE:
			return false, nil
		case ds_model_pb.Status_FAILED:
			return false, errors.Errorf("pending desired state failed to apply. %s", summary.PendingSetDesiredState.GetTaskResult().GetLog())
		default:
			return false, errors.Errorf("internal error: entity graph not set")
		}
	}
	serviceDs := summary.DesiredState.GetService()
	if serviceDs == nil {
		return false, errors.Errorf("internal error: no service desired state")
	}
	releaseChannelsSet := sets.FromSlice(releaseChannels)
	completed := 0
	for _, node := range summary.EntityGraph.GetEntities() {
		switch node.GetId().GetType() {
		case ds_model_pb.Type_SERVICE_INSTANCE:
			rc := node.DesiredState.GetServiceInstance().GetReleaseChannel()
			if releaseChannelsSet.Contains(rc) {
				releaseChannelsSet.Remove(rc)
				switch node.Status {
				case ds_model_pb.Status_CONVERGED:
					completed++
				case ds_model_pb.Status_FAILED, ds_model_pb.Status_ROLLED_BACK:
					return true, errors.Errorf("Convergence for '%s' ended with status '%s'", rc, node.Status)
				case ds_model_pb.Status_REPLACED:
					if successIfReplacedFlag {
						log.Printf("Convergence for '%s' was replaced by another desired state, considering that a success", rc)
						completed++
					} else {
						return true, errors.Errorf("Convergence for '%s' was replaced by another desired state", rc)
					}
				}
			}
		}
	}
	return completed == len(releaseChannels), nil
}

func compatibleParamValue(def *common_config_pb.ParameterDefinition, value *common_config_pb.ParameterValue) bool {
	if def == nil {
		return false
	}
	switch def.ConfigOneof.(type) {
	case *common_config_pb.ParameterDefinition_DockerImage:
		switch value.ValueOneof.(type) {
		case *common_config_pb.ParameterValue_DockerImageTag:
			return true
		}
	case *common_config_pb.ParameterDefinition_Commit:
		switch value.ValueOneof.(type) {
		case *common_config_pb.ParameterValue_Commit:
			return true
		}
	case *common_config_pb.ParameterDefinition_Blob:
		switch value.ValueOneof.(type) {
		case *common_config_pb.ParameterValue_Blob:
			return true
		}
	case *common_config_pb.ParameterDefinition_Secret:
		switch value.ValueOneof.(type) {
		case *common_config_pb.ParameterValue_Secret:
			return true
		}
	case *common_config_pb.ParameterDefinition_Int:
		switch value.ValueOneof.(type) {
		case *common_config_pb.ParameterValue_Int:
			return true
		}
	case *common_config_pb.ParameterDefinition_String_:
		switch value.ValueOneof.(type) {
		case *common_config_pb.ParameterValue_String_:
			return true
		}
	}
	return false
}

func buildParamValue(def *common_config_pb.ParameterDefinition, value string) (*common_config_pb.ParameterValue, error) {
	switch def.ConfigOneof.(type) {
	case *common_config_pb.ParameterDefinition_DockerImage:
		// check if this is a full image url
		ref, err := imageutils.ParseReference(value)
		if err == nil {
			// it is a full image url
			value = ref.Identifier.Name()
		} else {
			if errors.Is(err, imageutils.ErrNoIdentifier) {
				// was a full image url but no tag/digest present
				log.Fatalf("image url '%s' is missing a tag or digest", value)
			}
			// assume it's a tag, nothing to do
		}
		return &common_config_pb.ParameterValue{
			Name: def.Name,
			ValueOneof: &common_config_pb.ParameterValue_DockerImageTag{
				DockerImageTag: value,
			},
		}, nil
	case *common_config_pb.ParameterDefinition_Int:
		i, err := strconv.Atoi(value)
		if err != nil {
			return nil, errors.Wrapf(err, "%s is not a valid int", value)
		}
		return &common_config_pb.ParameterValue{
			Name: def.Name,
			ValueOneof: &common_config_pb.ParameterValue_Int{
				Int: int64(i),
			},
		}, nil
	case *common_config_pb.ParameterDefinition_Commit:
		return &common_config_pb.ParameterValue{
			Name: def.Name,
			ValueOneof: &common_config_pb.ParameterValue_Commit{
				Commit: value,
			},
		}, nil
	case *common_config_pb.ParameterDefinition_Blob:
		return &common_config_pb.ParameterValue{
			Name: def.Name,
			ValueOneof: &common_config_pb.ParameterValue_Blob{
				Blob: &common_config_pb.BlobParameterValue{
					BlobOneof: &common_config_pb.BlobParameterValue_Inlined{
						Inlined: value,
					},
				},
			},
		}, nil
	case *common_config_pb.ParameterDefinition_String_:
		return &common_config_pb.ParameterValue{
			Name: def.Name,
			ValueOneof: &common_config_pb.ParameterValue_String_{
				String_: value,
			},
		}, nil
	case *common_config_pb.ParameterDefinition_Secret:
		return nil, errors.Errorf("Secret paramter cannot be set via pvnctl")
	}
	return nil, errors.Errorf("Unsupported param type %T", def.ConfigOneof)
}

var push2Cmd = &cobra.Command{
	Use:   "push2",
	Short: "Push a service with parameters.",
	Long: `Push a service with parameters. Parameters not specified will be taken from latest value and then fall back to default value.

pvnctl services --app=<app> push2 --param=name=value --param-by-release-channel=release_channel=name=value <service>
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		rcHolder := &rcParamHolder{}
		if len(paramFlag) > 0 {
			paramHolder, err := makeParamHolder(paramFlag)
			if err != nil {
				log.Fatal(err)
			}
			rcHolder.globalHolder = paramHolder
		}
		if len(paramByReleaseChannelFlag) > 0 {
			holder, err := makeRcParamHolder(paramByReleaseChannelFlag)
			if err != nil {
				log.Fatal(err)
			}
			rcHolder.paramsByRC = holder
		}

		// Unlike update-images, require parameter names to be explicitly specified.
		if rcHolder.globalHolder != nil {
			for name := range rcHolder.globalHolder.paramValues {
				if name == "" {
					log.Fatal("Parameters must be of the format --param=name=value")
				}
			}
		}
		if rcHolder.paramsByRC != nil {
			for _, holder := range rcHolder.paramsByRC {
				for name := range holder.paramValues {
					if name == "" {
						log.Fatal("Per-release-channel parameters must be of the format --param-by-release-channel=release_channel=name=value,release_channel2=name=value2,...")
					}
				}
			}
		}

		waitChannels, err := cmd.Flags().GetStringSlice("wait-channel")
		if err != nil {
			log.Fatal(err)
		}
		serviceName := args[0]
		rcResp, err := cmdutil.GetReleaseChannelManagerClient().ListReleaseChannels(ctx, &release_channel_pb.ListReleaseChannelsReq{
			Application: cmdutil.App,
		})
		if err != nil {
			log.Fatal(err)
		}
		if rcHolder.paramsByRC != nil {
			for _, rc := range rcResp.ReleaseChannels {
				if _, found := rcHolder.paramsByRC[rc.Meta.Name]; !found {
					log.Fatalf("Missing parameter for ReleaseChannel %s", rc.Meta.Name)
				}
			}
		}
		maybeVersionResp, err := cmdutil.GetServiceManagerClient().GetMaterializedConfig(ctx, &service_pb.GetMaterializedConfigReq{
			Service:     serviceName,
			Application: cmdutil.App,
		})
		if err != nil {
			if perrors.GrpcErrCode(err) != codes.NotFound {
				log.Fatal(err)
			}
		}
		if err == nil && maybeVersionResp.GetVersionMetadata().GetConfigVersion() == "" {
			log.Fatal("Service is not using parameters flow")
		}
		err = push2CmdRun(ctx, serviceName, rcHolder, maybeVersionResp, rcResp, waitChannels)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func push2CmdRun(ctx context.Context, serviceName string, rcHolder *rcParamHolder, maybeVersionResp *service_pb.GetMaterializedConfigResp, rcResp *release_channel_pb.ListReleaseChannelsResp, waitChannels []string) error {
	validRcs := sets.NewSet[string]()
	for _, rc := range rcResp.ReleaseChannels {
		validRcs.Add(rc.Meta.Name)
	}
	// new parametrized flow
	latestConfigResp, err := cmdutil.GetServiceManagerClient().GetServiceConfig(ctx, &service_pb.GetServiceConfigReq{
		Application: cmdutil.App,
		Service:     serviceName,
	})
	if err != nil {
		return errors.Wrap(err, "failed to get config")
	}
	paramsByName := make(map[string]*common_config_pb.ParameterDefinition, len(latestConfigResp.Config.Parameters))
	for _, param := range latestConfigResp.Config.Parameters {
		paramsByName[param.Name] = param
	}
	usedParams := sets.NewSet[string]()
	var paramValues []*common_config_pb.ParameterValue
	paramValuesByRc := map[string]*service_pb.ApplyParametersReq_PerReleaseChannel{}
	if rcHolder.globalHolder != nil {
		paramHolder := rcHolder.globalHolder
		if paramHolder.Count() == 1 {
			if len(paramsByName) == 1 {
				for name := range paramsByName {
					paramHolder.maybeSingleParamName = name
				}
			} else {
				if paramHolder.maybeSingleParamName == "" {
					return errors.Errorf(
						"Service %s has %d params, please pass the param name you wish to update, e.g:\n\t--param <param_name>=%s",
						serviceName, len(paramsByName), paramHolder.singleParamValue,
					)
				}
			}

		}
		for paramName, rawValue := range paramHolder.GetParamValues() {
			usedParams.Add(paramName)
			def, ok := paramsByName[paramName]
			if !ok {
				return errors.Errorf("Service %s has no parameter %s", serviceName, paramName)
			}
			value, err := buildParamValue(def, rawValue)
			if err != nil {
				return err
			}
			paramValues = append(paramValues, value)
		}
	}
	if rcHolder.paramsByRC != nil {
		for rc, imgHolder := range rcHolder.paramsByRC {
			if imgHolder.Count() == 1 {
				if len(paramsByName) == 1 {
					for name := range paramsByName {
						imgHolder.maybeSingleParamName = name
					}
				} else {
					if imgHolder.maybeSingleParamName == "" {
						return errors.Errorf(
							"Service %s has %d params, please pass the param name you wish to update, e.g:\n\t--param-by-release-channel %s=<param_name>=%s",
							serviceName, len(paramsByName), rc, imgHolder.singleParamValue,
						)
					}
				}

			}
			perRc, ok := paramValuesByRc[rc]
			if !ok {
				paramValuesByRc[rc] = &service_pb.ApplyParametersReq_PerReleaseChannel{}
				perRc = paramValuesByRc[rc]
				perRc.ReleaseChannel = rc
			}
			for paramName, rawValue := range imgHolder.GetParamValues() {
				usedParams.Add(paramName)
				def, ok := paramsByName[paramName]
				if !ok {
					return errors.Errorf("Service %s has no parameter %s", serviceName, paramName)
				}
				value, err := buildParamValue(def, rawValue)
				if err != nil {
					log.Fatal(err)
				}
				perRc.Parameters = append(perRc.Parameters, value)
			}
		}
	}
	// maybeVersionResp may be nil here (if service is never pushed), so use the Get* APIs
	prevGlobalParams := maybeVersionResp.GetCompiledConfig().GetParameterValues().GetParameters()
	prevPerRcParams := maybeVersionResp.GetCompiledConfig().GetParameterValues().GetPerReleaseChannel()
	for _, param := range prevGlobalParams {
		if compatibleParamValue(paramsByName[param.Name], param) && !usedParams.Contains(param.Name) {
			paramValues = append(paramValues, param)
		}
	}
	for _, prevPerRc := range prevPerRcParams {
		if !validRcs.Contains(prevPerRc.ReleaseChannel) {
			// old rc, no longer relevant
			continue
		}
		perRc, ok := paramValuesByRc[prevPerRc.ReleaseChannel]
		if !ok {
			paramValuesByRc[prevPerRc.ReleaseChannel] = &service_pb.ApplyParametersReq_PerReleaseChannel{}
			perRc = paramValuesByRc[prevPerRc.ReleaseChannel]
			perRc.ReleaseChannel = prevPerRc.ReleaseChannel
		}
		for _, param := range prevPerRc.Parameters {
			if compatibleParamValue(paramsByName[param.Name], param) && !usedParams.Contains(param.Name) {
				perRc.Parameters = append(perRc.Parameters, param)
			}
		}
	}
	perRcs := maps.Values(paramValuesByRc)
	sort.Slice(perRcs, func(i, j int) bool {
		return perRcs[i].ReleaseChannel < perRcs[j].ReleaseChannel
	})
	applyResp, err := cmdutil.GetServiceManagerClient().ApplyParameters(ctx, &service_pb.ApplyParametersReq{
		Oneof: &service_pb.ApplyParametersReq_ServiceConfigVersion{
			ServiceConfigVersion: &service_pb.ServiceConfigVersionReference{
				Application:          cmdutil.App,
				Service:              serviceName,
				ServiceConfigVersion: latestConfigResp.ConfigVersion,
			},
		},
		Parameters:        paramValues,
		PerReleaseChannel: perRcs,
		Application:       cmdutil.App,
	})
	if err != nil {
		return errors.Wrap(err, "failed to apply parameters")
	}
	log.Printf("Created version: %s\n", applyResp.Version)
	pushResult, err := pushService(ctx, cmdutil.App, serviceName, applyResp.Version, maestroOverwriteMode)
	if err != nil {
		return errors.Wrap(err, "failed to push")
	}

	switch cmdutil.OutputFormat {
	case "json":
		outBytes, err := json.Marshal(pushResult.desiredStatePush.IdOneof)
		if err != nil {
			return errors.Wrap(err, "failed to marshal push result")
		}
		fmt.Printf("%s", outBytes)
	default:
		log.Printf("Created release: %+v\n", pushResult.desiredStatePush)
	}

	if len(waitChannels) == 0 {
		return nil
	}
	if pushResult.desiredStatePush != nil {
		return waitSetDesiredStateResp(ctx, pushResult.desiredStatePush, waitChannels)
	} else {
		return errors.Errorf("invalid push result")
	}
}

// For parametrized flow, this exists purely to maintain backwards compatibility and is the same as push2 with
// the exception of
var updateImagesCmd = &cobra.Command{
	Use:    "update-images",
	Short:  "Update container images for an existing service and push it.",
	Hidden: true, // deprecated and only exists for backwards compatibility with CI orbs
	Long: `Update container images for an existing service while retaining all other current config and push it.

pvnctl services --app=<app> update-images <service>
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		if len(paramFlag) > 0 && len(paramByReleaseChannelFlag) > 0 {
			log.Fatal("--param and --param-by-release-channel cannot be set at the same time")
		}

		var rcHolder *rcParamHolder
		if len(paramFlag) > 0 {
			paramHolder, err := makeParamHolder(paramFlag)
			if err != nil {
				log.Fatal(err)
			}
			rcHolder = &rcParamHolder{
				globalHolder: paramHolder,
			}
		} else if len(paramByReleaseChannelFlag) > 0 {
			holder, err := makeRcParamHolder(paramByReleaseChannelFlag)
			if err != nil {
				log.Fatal(err)
			}
			rcHolder = &rcParamHolder{
				paramsByRC: holder,
			}
		} else {
			log.Fatal("One of --param or --param-by-release-channel must be set")
		}

		waitChannels, err := cmd.Flags().GetStringSlice("wait-channel")
		if err != nil {
			log.Fatal(err)
		}
		serviceName := args[0]
		rcResp, err := cmdutil.GetReleaseChannelManagerClient().ListReleaseChannels(ctx, &release_channel_pb.ListReleaseChannelsReq{
			Application: cmdutil.App,
		})
		if err != nil {
			log.Fatal(err)
		}
		if rcHolder.paramsByRC != nil {
			for _, rc := range rcResp.ReleaseChannels {
				if _, found := rcHolder.paramsByRC[rc.Meta.Name]; !found {
					log.Fatalf("Missing parameter for ReleaseChannel %s", rc.Meta.Name)
				}
			}
		}
		paramsFlow := false
		maybeVersionResp, err := cmdutil.GetServiceManagerClient().GetMaterializedConfig(ctx, &service_pb.GetMaterializedConfigReq{
			Service:     serviceName,
			Application: cmdutil.App,
		})
		if err != nil {
			if perrors.GrpcErrCode(err) != codes.NotFound {
				log.Fatal(err)
			}
			paramsFlow = true
		}
		if paramsFlow || maybeVersionResp.GetVersionMetadata().GetConfigVersion() != "" {
			err := push2CmdRun(ctx, serviceName, rcHolder, maybeVersionResp, rcResp, waitChannels)
			if err != nil {
				log.Fatal(err)
			}
			return
		}

		if rcHolder.paramsByRC != nil {
			if maybeVersionResp.Config.PerReleaseChannel == nil {
				maybeVersionResp.Config.PerReleaseChannel = []*service_pb.PerReleaseChannelConfig{}
			}

			perRcConfigs := map[string]*service_pb.PerReleaseChannelConfig{}
			for _, perRcConfig := range maybeVersionResp.Config.PerReleaseChannel {
				perRcConfigs[perRcConfig.ReleaseChannel] = perRcConfig
			}

			topLevelPrograms := make(map[string]*common_config_pb.ProgramConfig)
			for _, program := range maybeVersionResp.Config.Programs {
				topLevelPrograms[program.Name] = program
			}

			updatedRcConfigs := make([]*service_pb.PerReleaseChannelConfig, 0, len(perRcConfigs))
			for rc, imgHolder := range rcHolder.paramsByRC {
				if imgHolder.Count() == 1 {
					if len(maybeVersionResp.Config.Programs) == 1 {
						imgHolder.maybeSingleParamName = maybeVersionResp.Config.Programs[0].Name
					} else {
						// if only one program image is being set, but there
						// are multiple programs, we can't infer which and need
						// the user to provide it
						if imgHolder.maybeSingleParamName == "" {
							log.Fatalf(
								"Service %s has %d programs, please pass the program name you wish to update, e.g:\n\t--images-by-release-channel %s=<program_name>=%s",
								serviceName, len(maybeVersionResp.Config.Programs), rc, imgHolder.singleParamValue,
							)
						}
					}
				}
				rcConfig, ok := perRcConfigs[rc]
				if !ok {
					rcConfig = &service_pb.PerReleaseChannelConfig{
						ReleaseChannel: rc,
					}
					perRcConfigs[rc] = rcConfig
				}
				programConfigs := map[string]*common_config_pb.PerReleaseChannelProgramConfig{}
				for _, cfg := range rcConfig.Programs {
					programConfigs[cfg.Name] = cfg
				}

				for program, image := range imgHolder.GetParamValues() {
					programConfig, ok := programConfigs[program]
					if !ok {
						programConfig = &common_config_pb.PerReleaseChannelProgramConfig{
							Name: program,
						}
						programConfigs[program] = programConfig
					}
					// check if this is a full image url
					ref, err := imageutils.ParseReference(image)
					if err == nil {
						// it is a full image url
						if (programConfig.ImageRegistryInfo != nil) || (topLevelPrograms[program] != nil &&
							topLevelPrograms[program].ImageRegistryInfo != nil) {
							programConfig.ImageTag = ref.Identifier.Name()
							programConfig.Image = ""
						} else {
							programConfig.ImageTag = ""
							programConfig.Image = image
						}

					} else {
						if errors.Is(err, imageutils.ErrNoIdentifier) {
							// was a full image url but no tag/digest present
							log.Fatalf("image url '%s' is missing a tag or digest", image)
						}
						// assume it's a tag
						programConfig.ImageTag = image
						programConfig.Image = ""
					}
				}

				updatedProgramConfigs := make([]*common_config_pb.PerReleaseChannelProgramConfig, 0, len(programConfigs))
				for _, cfg := range programConfigs {
					updatedProgramConfigs = append(updatedProgramConfigs, cfg)
				}
				rcConfig.Programs = updatedProgramConfigs
				updatedRcConfigs = append(updatedRcConfigs, rcConfig)
			}
			maybeVersionResp.Config.PerReleaseChannel = updatedRcConfigs
		} else {
			imgHolder := rcHolder.globalHolder

			if len(maybeVersionResp.Config.Programs) != imgHolder.Count() {
				log.Fatalf("Service %s has %d images to update - got %d updates", serviceName, len(maybeVersionResp.Config.Programs), imgHolder.Count())
			}

			for _, program := range maybeVersionResp.Config.Programs {
				if updatedImage, ok := imgHolder.GetParamValue(program.Name); !ok {
					log.Fatalf("No updated image found for program %s %+v", program.Name, program)
				} else {
					// check if this is a full image url
					ref, err := imageutils.ParseReference(updatedImage)
					if err == nil {
						// it is a full image url
						if program.ImageRegistryInfo != nil {
							program.ImageTag = ref.Identifier.Name()
							program.Image = ""
						} else {
							program.ImageTag = ""
							program.Image = updatedImage
						}

					} else {
						if errors.Is(err, imageutils.ErrNoIdentifier) {
							// was a full image url but no tag/digest present
							log.Fatalf("image url '%s' is missing a tag or digest", updatedImage)
						}
						// assume it's a tag
						program.ImageTag = updatedImage
						program.Image = ""
					}
				}
			}
		}

		pushResult, err := configureAndPushService(ctx, maybeVersionResp.Config, false)
		if err != nil {
			log.Fatal(err)
		}
		if len(waitChannels) == 0 {
			return
		}

		if pushResult.desiredStatePush != nil {
			err := waitSetDesiredStateResp(ctx, pushResult.desiredStatePush, waitChannels)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal("invalid push result")
		}
	},
}

func init() {
	RootCmd.AddCommand(updateImagesCmd)
	updateImagesCmd.Flags().StringSliceVar(&paramFlag, "image", nil, "Image to update to. If there are multiple programs in this service, use <program_name>=<image> format.")
	updateImagesCmd.Flags().StringSliceVar(&paramByReleaseChannelFlag, "images-by-release-channel", nil, "Image(s) to update to, per release channel. The format is a comma-separated list of <release_channel>=<image>. If you have multiple programs you can use <release_channel>=<program>=<image>|<program>=<image>,<release_channel>=....")
	updateImagesCmd.Flags().StringSlice("wait-channel", nil, "A list of release channel names. When provided, this command will not exit until the push to each channel is finished (either succeeded or failed.)")
	updateImagesCmd.Flags().BoolVar(&successIfReplacedFlag, "success-on-replaced", true, "When wait-channel is enabled, if the desired state is replaced, consider the push successful.")
	updateImagesCmd.Flags().BoolVar(&maestroOverwriteMode, "maestro-overwrite", false, "Enable maestro overwrite mode for this push.")
	cmdutil.Must(updateImagesCmd.Flags().MarkHidden("maestro-overwrite")) // Keep this outside public documentation till it is stabilized

	RootCmd.AddCommand(push2Cmd)
	push2Cmd.Flags().StringSliceVar(&paramFlag, "param", nil, "Parameters in the name=value format.")
	push2Cmd.Flags().StringSliceVar(&paramByReleaseChannelFlag, "param-by-release-channel", nil, "Parameters per release channel. The format is a comma-separated list of <release_channel>=<name>=<value>.")
	push2Cmd.Flags().StringSlice("wait-channel", nil, "A list of release channel names. When provided, this command will not exit until the push to each channel is finished (either succeeded or failed.)")
	push2Cmd.Flags().BoolVar(&successIfReplacedFlag, "success-on-replaced", true, "When wait-channel is enabled, if the desired state is replaced, consider the push successful.")
	push2Cmd.Flags().BoolVar(&maestroOverwriteMode, "maestro-overwrite", false, "Enable maestro overwrite mode for this push.")
	cmdutil.Must(push2Cmd.Flags().MarkHidden("maestro-overwrite")) // Keep this outside public documentation till it is stabilized
}
