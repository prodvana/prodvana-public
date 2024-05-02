package healthcheck

import (
	"context"
	"log"
	"prodvana/cmd/cmdutil"
	"prodvana/cmd/pvnctl/cmd/services/primitives"

	"github.com/spf13/cobra"
	"google.golang.org/grpc/health/grpc_health_v1"
)

var RootCmd = &cobra.Command{
	Use:   "healthcheck",
	Short: "Healthcheck connection to Prodvana.",
	Long: `Healthcheck connection to Prodvana.

pvnctl healthcheck
`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		resp, err := cmdutil.GetHealthClient().Check(ctx, &grpc_health_v1.HealthCheckRequest{})
		if err != nil {
			log.Fatalf("Unable to talk to Prodvana: %s", err)
		}
		if resp.Status != grpc_health_v1.HealthCheckResponse_SERVING {
			log.Fatalf("Internal Prodvana error. Server not running: %s", resp.Status)
		}
		cmd.Printf("Connection to Prodvana healthy\n")
	},
}

func init() {
	RootCmd.AddCommand(primitives.RootCmd)
}
