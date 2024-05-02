package services

import (
	"context"
	"log"
	"os"

	"prodvana/cmd/cmdutil"

	"github.com/spf13/cobra"

	convergence_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/convergence"
	service_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/service"
)

func observerConvergeMode(cmd *cobra.Command, args []string) {
	serviceName := args[0]
	ctx := context.Background()

	var convergenceMode convergence_pb.ConvergenceMode
	switch cmd.Name() {
	case "set-converge-mode":
		convergenceMode = convergence_pb.ConvergenceMode_CONVERGE
	case "set-observer-mode":
		convergenceMode = convergence_pb.ConvergenceMode_OBSERVER
	default:
		log.Fatalf("Unexpected command %s", cmd.Name())
	}

	_, err := cmdutil.GetServiceManagerClient().SetServiceConvergenceMode(ctx, &service_pb.SetServiceConvergenceModeReq{
		Application:     cmdutil.App,
		Service:         serviceName,
		ConvergenceMode: convergenceMode,
	})
	if err != nil {
		cmd.PrintErrf("Failed to set %s mode for service %s: %+v\n", convergenceMode.String(), serviceName, err)
		os.Exit(1)
	}
	cmd.Printf("Service %s set to %s mode\n", serviceName, convergenceMode.String())
}

var observerMode = &cobra.Command{
	Use:    "set-observer-mode <service-name>",
	Hidden: true,
	Short:  "Set service to observer mode",
	Long: `Disable convergence for service.

pvnctl services --app=<app> set-observer-mode <service-name>
`,
	Args: cobra.ExactArgs(1),
	Run:  observerConvergeMode,
}

var convergeMode = &cobra.Command{
	Use:    "set-converge-mode <service-name>",
	Hidden: true,
	Short:  "Set service to converge mode",
	Long: `Enable convergence for service.

pvnctl services --app=<app> set-converge-mode <service-name>
`,
	Args: cobra.ExactArgs(1),
	Run:  observerConvergeMode,
}

func init() {
	RootCmd.AddCommand(observerMode)
	RootCmd.AddCommand(convergeMode)
}
