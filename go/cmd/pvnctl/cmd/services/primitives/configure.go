package primitives

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	common_config_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config"
	service_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/service"
	version_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/version"

	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/durationpb"
)

var configureCmd = &cobra.Command{
	Use:   "configure <name>",
	Short: "Configure a service to prepare for push.",
	Long: `Configure a service to be pushed later.

pvnctl services --app=<app> primitives configure <name> [FLAGS]  # will output version string
pvnctl services --app=<app> primitives push <name> --version=<version from configure command>
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		replicas, err := cmd.Flags().GetInt("replicas")
		if err != nil {
			log.Fatal(err)
		}
		image, err := cmd.Flags().GetString("image")
		if err != nil {
			log.Fatal(err)
		}
		imageTag, err := cmd.Flags().GetString("image-tag")
		if err != nil {
			log.Fatal(err)
		}
		registryId, err := cmd.Flags().GetString("container-registry-id")
		if err != nil {
			log.Fatal(err)
		}
		registryName, err := cmd.Flags().GetString("container-registry")
		if err != nil {
			log.Fatal(err)
		}
		imageRepository, err := cmd.Flags().GetString("image-repository")
		if err != nil {
			log.Fatal(err)
		}
		if (registryId != "" || registryName != "") && imageRepository == "" {
			log.Fatal("pass -image-repository")
		}
		skipRegistryCheck, err := cmd.Flags().GetBool("skip-registry-check")
		if err != nil {
			log.Fatal(err)
		}
		cmdStr, err := cmd.Flags().GetString("cmd")
		if err != nil {
			log.Fatal(err)
		}
		envVarsRaw, err := cmd.Flags().GetStringArray("env")
		if err != nil {
			log.Fatal(err)
		}
		envVarsSecretsRaw, err := cmd.Flags().GetStringArray("env-secrets")
		if err != nil {
			log.Fatal(err)
		}
		envVars, err := cmdutil.ProcessEnvVars(envVarsRaw, envVarsSecretsRaw)
		if err != nil {
			log.Fatal(err)
		}
		capabilities, err := cmd.Flags().GetStringArray("capability")
		if err != nil {
			log.Fatal(err)
		}
		useRuntimeExtension, err := cmd.Flags().GetBool("use-runtime-extension")
		if err != nil {
			log.Fatal(err)
		}
		progressDeadline, err := cmd.Flags().GetDuration("progress-deadline")
		if err != nil {
			log.Fatal(err)
		}
		bundleNameOverride, err := cmd.Flags().GetString("bundle-name-override")
		if err != nil {
			log.Fatal(err)
		}
		capRef := make([]*service_pb.CapabilityReference, len(capabilities))
		for idx, cap := range capabilities {
			capRef[idx] = &service_pb.CapabilityReference{Name: cap}
		}
		ports, err := cmd.Flags().GetIntSlice("port")
		if err != nil {
			log.Fatal(err)
		}
		portConfigs := make([]*common_config_pb.PortConfig, len(ports))
		for idx, p := range ports {
			portConfigs[idx] = &common_config_pb.PortConfig{Port: int32(p)}
		}

		resources, err := processResourceRequirements(cmd)
		if err != nil {
			log.Fatal(err)
		}

		healthCheck, err := processHealthChecks(cmd)
		if err != nil {
			log.Fatal(err)
		}

		releaseStrategy, err := processReleaseStrategy(cmd)
		if err != nil {
			log.Fatal(err)
		}

		var registryInfo *common_config_pb.ImageRegistryInfo
		if registryId != "" || registryName != "" {
			registryInfo = &common_config_pb.ImageRegistryInfo{
				ContainerRegistryId: registryId,
				ContainerRegistry:   registryName,
				ImageRepository:     imageRepository,
			}
		}

		var genericRuntime *service_pb.ServiceConfig_RuntimeExtension
		if useRuntimeExtension {
			genericRuntime = &service_pb.ServiceConfig_RuntimeExtension{
				RuntimeExtension: &service_pb.RuntimeExtensionConfig{},
			}
		}

		ctx := context.Background()
		svcConfig := &service_pb.ServiceConfig{
			Name: name,
			Programs: []*common_config_pb.ProgramConfig{
				{
					Name:              "default",
					Image:             image,
					ImageTag:          imageTag,
					Cmd:               strings.Fields(cmdStr),
					Resources:         resources,
					HealthCheck:       healthCheck,
					Env:               envVars,
					Ports:             portConfigs,
					ImageRegistryInfo: registryInfo,
				},
			},
			Replicas: &service_pb.ReplicasConfig{
				ConfigOneof: &service_pb.ReplicasConfig_Fixed{
					Fixed: int32(replicas),
				},
			},
			ReleaseStrategy: releaseStrategy,
			Capabilities:    capRef,
			ConfigOneof:     genericRuntime,
		}
		ApplyParameters := &service_pb.ApplyParametersReq{
			Oneof: &service_pb.ApplyParametersReq_ServiceConfig{
				ServiceConfig: svcConfig,
			},
			TestOnlySkipRegistryCheck: skipRegistryCheck,
			Application:               cmdutil.App,
			Source:                    version_pb.Source_INTERACTIVE_PVNCTL,
			BundleNameOverride:        bundleNameOverride,
		}
		if progressDeadline > 0 {
			svcConfig.ProgressDeadline = durationpb.New(progressDeadline)
		}
		resp, err := cmdutil.GetServiceManagerClient().ApplyParameters(ctx, ApplyParameters)
		if err != nil {
			cmd.PrintErrf("Failed to configure service: %+v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", resp.Version)
	},
}

func processReleaseStrategy(cmd *cobra.Command) (*service_pb.ReleaseStrategyConfig, error) {
	individualStageDeadline, err := cmd.Flags().GetDuration("individual-stage-deadline")
	if err != nil {
		log.Fatal(err)
	}
	releaseDuration, err := cmd.Flags().GetDuration("automated-testing-duration")
	if err != nil {
		log.Fatal(err)
	}
	checkInterval, err := cmd.Flags().GetDuration("check-interval")
	if err != nil {
		log.Fatal(err)
	}
	failureThreshold, err := cmd.Flags().GetInt32("failure-threshold")
	if err != nil {
		log.Fatal(err)
	}
	manualApproval, err := cmd.Flags().GetBool("manual-approval")
	if err != nil {
		log.Fatal(err)
	}

	releaseStrategy := &service_pb.ReleaseStrategyConfig{
		AutomatedTestingDuration: durationpb.New(releaseDuration),
		MetricsAnalysis: []*service_pb.MetricAnalysis{
			{
				AnalysisOneof: &service_pb.MetricAnalysis_SuccessRate{
					SuccessRate: &service_pb.MetricAnalysis_SuccessRateConfig{
						// TODO(naphat) allow customization?
						MinThresholdPercent: 99.5,
					},
				},
			},
			{
				AnalysisOneof: &service_pb.MetricAnalysis_LatencyP95{
					LatencyP95: &service_pb.MetricAnalysis_LatencyConfig{
						// TODO(naphat) allow customization?
						MaxLatency: durationpb.New(500 * time.Millisecond),
					},
				},
			},
		},
		CheckInterval:    durationpb.New(checkInterval),
		FailureThreshold: failureThreshold,
		ManualApproval:   manualApproval,
	}

	if individualStageDeadline != 0 {
		releaseStrategy.IndividualStageDeadline = durationpb.New(individualStageDeadline)
	}
	return releaseStrategy, nil
}

func processResourceRequirements(cmd *cobra.Command) (*common_config_pb.ResourceRequirements, error) {
	requestMemory, err := cmd.Flags().GetString("memory")
	if err != nil {
		return nil, err
	}
	requestCpu, err := cmd.Flags().GetString("cpu")
	if err != nil {
		return nil, err
	}

	var resources *common_config_pb.ResourceRequirements

	if requestCpu != "" && requestMemory != "" {
		cmd.Println("setting cpu and memory ", requestCpu, requestMemory)
		resources = &common_config_pb.ResourceRequirements{
			Requests: &common_config_pb.ResourceList{
				Memory: requestMemory,
				Cpu:    requestCpu,
			},
		}
	}

	return resources, nil
}

func processHealthChecks(cmd *cobra.Command) (*common_config_pb.HealthCheck, error) {
	healthCheckHttpPath, err := cmd.Flags().GetString("healthcheck-http-path")
	if err != nil {
		return nil, err
	}

	healthCheckHttpPort, err := cmd.Flags().GetInt32("healthcheck-http-port")
	if err != nil {
		return nil, err
	}

	healthCheckDelay, err := cmd.Flags().GetDuration("healthcheck-delay")
	if err != nil {
		return nil, err
	}

	healthCheckPeriod, err := cmd.Flags().GetDuration("healthcheck-period")
	if err != nil {
		return nil, err
	}

	healthCheckCmd, err := cmd.Flags().GetString("healthcheck-cmd")
	if err != nil {
		return nil, err
	}

	healthCheckTcpPort, err := cmd.Flags().GetInt32("healthcheck-tcp-port")
	if err != nil {
		return nil, err
	}

	healthCheckTcpHost, err := cmd.Flags().GetString("healthcheck-tcp-host")
	if err != nil {
		return nil, err
	}

	var healthCheck *common_config_pb.HealthCheck

	if healthCheckHttpPath != "" {
		healthCheck = &common_config_pb.HealthCheck{
			ProbeConfig: &common_config_pb.HealthCheck_Http{
				Http: &common_config_pb.HttpProbe{
					Path: healthCheckHttpPath,
					Port: healthCheckHttpPort,
				},
			},
		}
	} else if healthCheckCmd != "" {
		healthCheck = &common_config_pb.HealthCheck{
			ProbeConfig: &common_config_pb.HealthCheck_Cmd{
				Cmd: &common_config_pb.CmdProbe{
					Command: strings.Fields(healthCheckCmd),
				},
			},
		}
	} else if healthCheckTcpPort != 0 {
		healthCheck = &common_config_pb.HealthCheck{
			ProbeConfig: &common_config_pb.HealthCheck_Tcp{
				Tcp: &common_config_pb.TcpProbe{
					Port: healthCheckTcpPort,
					Host: healthCheckTcpHost,
				},
			},
		}
	}

	if healthCheckDelay > 0 {
		healthCheck.Delay = &durationpb.Duration{
			Seconds: int64(healthCheckDelay.Seconds()),
		}
	}
	if healthCheckPeriod > 0 {
		healthCheck.Period = &durationpb.Duration{
			Seconds: int64(healthCheckPeriod.Seconds()),
		}
	}

	return healthCheck, nil
}

func init() {
	RootCmd.AddCommand(configureCmd)
	configureCmd.Flags().StringP("image", "i", "", "Image to use for the service.")
	// TODO(naphat) support prompting for these
	configureCmd.Flags().StringP("cmd", "c", "", "Command to run for the service.")
	configureCmd.Flags().IntP("replicas", "r", 3, "Number of replicas for the service.")
	configureCmd.Flags().IntSliceP("port", "p", nil, "Ports to expose.")

	configureCmd.Flags().String("container-registry-id", "", "Optionally use a container registry integration.")
	configureCmd.Flags().String("container-registry", "", "Optionally use a container registry integration.")
	configureCmd.Flags().String("image-repository", "", "Specify image repository for container integration integration. Required if -container-registry[-id] is passed.")
	configureCmd.Flags().Bool("skip-registry-check", false, "Skip any kind of attempts to contact container registry, e.g. to check if image exists or get program defaults.")
	cmdutil.Must(configureCmd.Flags().MarkHidden("skip-registry-check"))
	configureCmd.Flags().StringP("image-tag", "t", "", "Image tag to use for the service.")
	configureCmd.MarkFlagsMutuallyExclusive("image", "image-tag")
	configureCmd.MarkFlagsRequiredTogether("image-tag", "image-repository")

	// Release Strategy
	configureCmd.Flags().Bool("manual-approval", false, "Whether or not the service requires a manual approval on push.")
	configureCmd.Flags().Duration("individual-stage-deadline", 0, "How long in each individual stage of release analysis before rolling back. Includes the time it takes for pods to be healthy. Defaults to 10 minutes.")
	configureCmd.Flags().DurationP("automated-testing-duration", "d", 10*time.Minute, "How long to run automated release analysis for.")
	configureCmd.Flags().Duration("check-interval", 0, "Interval to check for automated release analysis.")
	configureCmd.Flags().Int32("failure-threshold", 0, "How many failures before automated rollback happens.")
	configureCmd.Flags().Duration("progress-deadline", 0, "How long to wait before marking deployment as failing if it does not make progress.")

	// Resource Limits
	configureCmd.Flags().String("memory", "", "Resource memory limit.")
	configureCmd.Flags().String("cpu", "", "Resource cput limit.")

	// Health Checks
	configureCmd.Flags().Duration("healthcheck-delay", 0, "Initial delay to start the health check.")
	configureCmd.Flags().Duration("healthcheck-period", 0, "Interval to repeat the health check.")
	configureCmd.Flags().String("healthcheck-http-path", "", "Path for the http health check.")
	configureCmd.Flags().Int32("healthcheck-http-port", 0, "Port for the http health check.")
	configureCmd.Flags().String("healthcheck-cmd", "", "Command for health check")
	configureCmd.Flags().Int32("healthcheck-tcp-port", 0, "TCP port for health check")
	configureCmd.Flags().String("healthcheck-tcp-host", "", "TCP host for health check")

	// Environment Variables
	configureCmd.Flags().StringArrayP("env", "e", nil, "Env variables in VAR=VALUE format, can be specified multiple times. If duplicate variables are specified, the last one takes precedence.")
	configureCmd.Flags().StringArrayP("env-secrets", "s", nil, "Env variables set from current value of secrets in VAR=SECRET_NAME format, can be specified multiple times. If duplicate variables are specified, the last one takes precedence.")

	// Capabilities
	configureCmd.Flags().StringArray("capability", nil, "Capabilities to add to the service, can be specified multiple times.")

	// Generic runtime
	configureCmd.Flags().Bool("use-runtime-extension", false, "Use generic runtime for this service")

	configureCmd.Flags().String("bundle-name-override", "", "Bundle/version name to use instead of autogenerated name")
}
