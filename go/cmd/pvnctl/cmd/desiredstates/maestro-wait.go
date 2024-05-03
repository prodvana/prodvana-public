package desiredstates

import (
	"context"
	go_errors "errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	ds "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state"
	maestro_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/maestro"
	"github.com/spf13/cobra"
)

var maestroWaitCmd = &cobra.Command{
	Use:    "maestro-wait <release id> <release channel id>",
	Short:  "Wait for maestro release to complete for a given release channel",
	Hidden: true,
	Args:   cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		releaseId := args[0]
		releaseChannelId := args[1]

		timeout, err := cmd.Flags().GetDuration("timeout")
		if err != nil {
			log.Fatal(err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		err = cmdutil.RenderWatch(ctx, func(ctx context.Context, writer io.Writer) (bool, error) {
			resp, err := cmdutil.GetDesiredStateManagerClient().GetMaestroRelease(ctx, &ds.GetMaestroReleaseReq{
				ReleaseId: releaseId,
			})
			if err != nil {
				return false, err
			}

			if resp.MaestroRelease == nil {
				fmt.Fprintf(writer, "No pending releases found for %s\n", releaseId)
				return false, nil
			}

			release := resp.MaestroRelease

			rcState, ok := release.State.ReleaseChannelStatuses[releaseChannelId]
			if !ok {
				return false, nil
			}

			return rcState == maestro_pb.Status_SUCCEEDED, nil
		})

		if err != nil {
			cmd.PrintErrf("%s\n", err)
			if ctx.Err() == context.DeadlineExceeded {
				os.Exit(3)
			}
			if go_errors.Is(err, errDesiredStateConvergenceFailed) {
				os.Exit(2)
			}
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(maestroWaitCmd)
	maestroWaitCmd.Flags().Duration("timeout", 5*time.Minute, "Timeout to wait for")
}
