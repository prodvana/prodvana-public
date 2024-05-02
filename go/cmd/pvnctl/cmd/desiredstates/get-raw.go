package desiredstates

import (
	"context"
	_ "embed"
	"log"
	"os"
	"prodvana/cmd/cmdutil"

	ds_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/spf13/cobra"
)

var getRaw = &cobra.Command{
	Use:    "get-raw",
	Hidden: true,
	Short:  "Get the raw desired state as a JSON",
	Long: `Get the raw desired state as a JSON.

pvnctl desired-states get-raw <desired state id>
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dsId := args[0]
		ctx := context.Background()
		resp, err := cmdutil.GetDesiredStateManagerClient().GetDesiredState(ctx, &ds_pb.GetDesiredStateReq{
			DesiredStateId: dsId,
		})
		if err != nil {
			log.Fatal(err)
		}

		jsonBytes, err := protojson.Marshal(resp)
		if err != nil {
			log.Fatal(err)
		}

		_, err = os.Stdout.Write(jsonBytes)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(getRaw)
}
