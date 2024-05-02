package desiredstates

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"prodvana/cmd/cmdutil"

	"github.com/prodvana/prodvana-public/go/pkg/protohelpers"

	common_config_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config"
	ds_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state"
	ds_model_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/model"
	version_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/version"

	"github.com/spf13/cobra"
)

//go:embed templates/desired_state.yaml
var defaultDesiredStateYaml string

func handleDesiredState(ctx context.Context, inputFileType, maybeInput, startingYaml string, noPrompt bool) (*ds_model_pb.State, error) {
	var ds ds_model_pb.State
	_, err := cmdutil.EditOrReadConfig(ctx, inputFileType, maybeInput, startingYaml, &ds, func(ctx context.Context, ds *ds_model_pb.State) ([]*common_config_pb.DangerousAction, error) {
		_, err := cmdutil.GetDesiredStateManagerClient().ValidateDesiredState(ctx, &ds_pb.ValidateDesiredStateReq{
			DesiredState: ds,
		})
		return nil, err
	}, "Save changes?", noPrompt)
	if err != nil {
		return nil, err
	}
	return &ds, nil
}

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure a new desired state.",
	Long: `Configure a new desired state.

pvnctl desired-states configure --input ds.yaml
`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		input, err := cmd.Flags().GetString("input")
		if err != nil {
			log.Fatal(err)
		}
		inputFileType, err := cmd.Flags().GetString("input-file-format")
		if err != nil {
			log.Fatal(err)
		}
		noPrompt, err := cmd.Flags().GetBool("no-prompt")
		if err != nil {
			log.Fatal(err)
		}
		ds, err := handleDesiredState(ctx, inputFileType, input, defaultDesiredStateYaml, noPrompt)
		if err != nil {
			log.Fatal(err)
		}
		setDesiredStateResp, err := cmdutil.GetDesiredStateManagerClient().SetDesiredState(ctx, &ds_pb.SetDesiredStateReq{
			DesiredState:              ds,
			Source:                    version_pb.Source_INTERACTIVE_PVNCTL,
			ForceAsyncSetDesiredState: cmdutil.ForceAsyncSetDesiredState,
		})
		if err != nil {
			log.Fatal(err)
		}

		switch cmdutil.OutputFormat {
		case "json":
			outBytes, err := json.Marshal(setDesiredStateResp.IdOneof)
			if err != nil {
				log.Fatalf("failed to marshal SetDesiredState result: %+v", err)
			}
			fmt.Printf("%s", outBytes)
		default:
			log.Printf("Created release: %+v\n", setDesiredStateResp)
		}
	},
}

func init() {
	RootCmd.AddCommand(configureCmd)
	configureCmd.Flags().String("input", "", "If set, take desired state from a file instead of launching an editor. Supported formats: json, yaml, pbtxt")
	configureCmd.Flags().String("input-file-format", protohelpers.FileTypeInfer, "Used in conjunction with --input. Specify input file type. Pass 'infer' to infer file type from file extension.")
	configureCmd.Flags().BoolP("no-prompt", "y", false, "Whether to prompt for confirmation before saving.")
}
