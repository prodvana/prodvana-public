package runtimes

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"prodvana/cmd/cmdutil"

	"github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/environment"
	version_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/version"

	"github.com/spf13/cobra"
)

var (
	addK8sMagicOutputUrlOnly bool
)

var addK8sMagicCmd = &cobra.Command{
	Use:   "add-k8s-magic <name>",
	Short: "Add an existing K8s runtime to Prodvana via Magic URL.",
	Long: `Add an existing K8s runtime to Prodvana via Magic URL. Prints the ID of the new runtime to stdout on success.

pvnctl runtimes add-k8s-magic <name> [FLAGS]
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		ctx := context.Background()

		agentEnvVars, err := cmd.Flags().GetStringArray("agent-env")
		if err != nil {
			log.Fatal(err)
		}
		envVarMap := make(map[string]string)
		for _, env := range agentEnvVars {
			components := strings.SplitN(env, "=", 2)
			if len(components) != 2 {
				log.Fatalf("--agent-env must be in VAR=VALUE format. Invalid format: %s", env)
			}
			envVarMap[components[0]] = components[1]
		}

		resp, err := cmdutil.GetEnvironmentManagerClient().LinkCluster(ctx,
			&environment.LinkClusterReq{
				Name: name,
				Type: environment.ClusterType_K8S,
				Auth: &environment.ClusterAuth{
					AuthOneof: &environment.ClusterAuth_K8S{
						K8S: &environment.ClusterAuth_K8SAuth{
							AgentEnv: envVarMap,
						},
					},
				},
				Source: version_pb.Source_INTERACTIVE_PVNCTL,
			})
		if err != nil {
			cmd.PrintErrf("Failed to add K8s cluster: %+v\n", err)
			os.Exit(1)
		}
		if !resp.Success {
			cmd.PrintErrf("Failed to add cluster: %s\n", resp.Msg)
			os.Exit(1)
		}

		if addK8sMagicOutputUrlOnly {
			fmt.Printf("%s\n", resp.K8SAgentUrl)
		} else {
			fmt.Printf(`Runtime %s created. Run the following command to link the cluster
kubectl apply -f "%s"
`, resp.ClusterId, resp.K8SAgentUrl)
		}
	},
}

func init() {
	RootCmd.AddCommand(addK8sMagicCmd)
	addK8sMagicCmd.Flags().BoolVar(&addK8sMagicOutputUrlOnly, "output-url-only", false, "If set, on success, only the magic k8s URL will be printed to stdout, instead of the full set of instructions. This is useful for scripting.")
	addK8sMagicCmd.Flags().StringArray("agent-env", nil, "Env variables supplied to agent - use this for setting HTTPS_PROXY, NO_PROXY etc")
}
