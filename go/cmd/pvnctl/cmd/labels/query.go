package labels

import (
	"context"
	"prodvana/cmd/cmdutil"
	"sort"

	"github.com/bradenaw/juniper/xmaps"
	object_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/object"
	"github.com/spf13/cobra"
)

var (
	typeToStringMap = map[object_pb.ObjectType]string{
		object_pb.ObjectType_RUNTIME:         "runtime",
		object_pb.ObjectType_APPLICATION:     "application",
		object_pb.ObjectType_SERVICE:         "service",
		object_pb.ObjectType_RELEASE_CHANNEL: "release-channel",
		object_pb.ObjectType_PROTECTION:      "protection",
	}
	stringToTypeMap = func() map[string]object_pb.ObjectType {
		reversed, ok := xmaps.ReverseSingle(typeToStringMap)
		if !ok {
			panic("invalid map")
		}
		return reversed
	}()
)

var queryCmd = &cobra.Command{
	Use:   "query <query>",
	Short: "Run a query to get a list of objects",
	Long: `Run a query to get a list of objects.

pvnctl labels query "@type=runtime"  # all runtimes
pvnctl labels query "@type=runtime env=staging"  # all runtimes with custom label env=staging
`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		query := args[0]
		ctx := context.Background()
		resp, err := cmdutil.GetObjectManagerClient().QueryObjects(ctx, &object_pb.QueryObjectReq{
			Query: query,
		})
		if err != nil {
			return err
		}
		objs := resp.Objects
		listing := cmdutil.NewOutputListing("TYPE", "ID", "NAME")
		sort.Slice(objs, func(i, j int) bool {
			if objs[i].Type == objs[j].Type {
				return objs[i].Name < objs[j].Name
			}
			return objs[i].Type < objs[j].Type
		})
		for _, obj := range objs {
			typeStr, ok := typeToStringMap[obj.Type]
			if !ok {
				typeStr = "unknown"
			}
			listing.AddRow(typeStr, obj.Id, obj.Name)
		}
		err = cmdutil.WriteStdout(listing)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(queryCmd)
}
