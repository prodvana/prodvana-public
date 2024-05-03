package blobs

import (
	"context"
	_ "embed"
	"io"
	"log"
	"os"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	"github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/blobs"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get blob content by ID.",
	Long: `Get blob content by ID.

pvnctl blobs get <id>
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		blobId := args[0]
		ctx := context.Background()
		blobResp, err := cmdutil.GetBlobsManagerClient().GetCasBlob(ctx, &blobs.GetCasBlobReq{Id: blobId})
		if err != nil {
			log.Fatalf("failed to get blob: %s", err)
		}
		defer func() { _ = blobResp.CloseSend() }()
		for {
			chunk, err := blobResp.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatalf("Failed to get logs chunk: %s", err)
			}
			_, err = os.Stdout.Write(chunk.Bytes)
			if err != nil {
				log.Fatalf("failed to write: %s", err)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
