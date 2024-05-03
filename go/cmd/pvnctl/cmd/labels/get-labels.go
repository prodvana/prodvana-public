package labels

import (
	"context"
	"fmt"
	"sort"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	object_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/object"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
)

var getLabelsCmd = &cobra.Command{
	Use:   "get-labels <type> <id>",
	Short: "Get labels for an object",
	Long: `Get labels for an object.

pvnctl labels get-labels runtime rtm-...
`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		objTypeStr := args[0]
		id := args[1]
		objType, ok := stringToTypeMap[objTypeStr]
		if !ok {
			return fmt.Errorf("invalid type: %s. Valid values are: %+v", objTypeStr, maps.Keys(stringToTypeMap))
		}
		ctx := context.Background()
		resp, err := cmdutil.GetObjectManagerClient().GetLabels(ctx, &object_pb.GetLabelsReq{
			Type: objType,
			Id:   id,
		})
		if err != nil {
			return err
		}
		labels := resp.Labels
		listing := cmdutil.NewOutputListing("LABEL", "VALUE")
		sort.Slice(labels, func(i, j int) bool {
			if labels[i].Label == labels[j].Label {
				return labels[i].Value < labels[j].Value
			}
			return labels[i].Label < labels[j].Label
		})
		for _, label := range labels {
			listing.AddRow(label.Label, label.Value)
		}
		err = cmdutil.WriteStdout(listing)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(getLabelsCmd)
}
