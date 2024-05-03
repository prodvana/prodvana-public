package runtimes

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	"github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/environment"
	version_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/version"

	"github.com/spf13/cobra"
)

var addK8sCmd = &cobra.Command{
	Use:   "add-k8s <name>",
	Short: "Add an existing K8s runtime to Prodvana.",
	Long: `Add an existing K8s runtime to Prodvana . Prints the ID and agent details for the new runtime to stdout on success.

pvnctl runtimes add-k8s <name> [FLAGS]
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		ctx := context.Background()

		agentEnvVars, err := cmd.Flags().GetStringArray("agent-env")
		if err != nil {
			log.Fatal(err)
		}
		agentManaged, err := cmd.Flags().GetBool("agent-externally-managed")
		if err != nil {
			log.Fatal(err)
		}
		outputFormat, err := cmd.Flags().GetString("format")
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
							AgentEnv:               envVarMap,
							AgentExternallyManaged: agentManaged,
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

		switch outputFormat {
		case "json":
			values := map[string]string{
				"cluster_id": resp.ClusterId,
				"image":      resp.K8SAgentImage,
				"args":       strings.Join(resp.K8SAgentArgs, " "),
				"token":      resp.K8SAgentApiToken,
				"magic_url":  resp.K8SAgentUrl,
			}
			jsn, err := json.Marshal(values)
			if err != nil {
				log.Fatal(err)
			}
			cmd.Printf("%s\n", jsn)
		default:
			cmd.Printf(`Runtime %s created. Run the following command to link the cluster
kubectl apply -f "%s"
`, resp.ClusterId, resp.K8SAgentUrl)
			cmd.Printf("\nAlternatively use the details below to construct a custom agent configuration:\n")
			cmd.Printf("Cluster ID: %s\n", resp.ClusterId)
			cmd.Printf("Image: %s\n", resp.K8SAgentImage)
			cmd.Printf("Agent Args: %s\n", strings.Join(resp.K8SAgentArgs, " "))
			cmd.Printf("Agent Token: %s\n", resp.K8SAgentApiToken)
		}

	},
}

func init() {
	RootCmd.AddCommand(addK8sCmd)
	addK8sCmd.Flags().Bool("agent-externally-managed", false, "If true, you are responsible for managing the prodvana agent lifecycle. If false, prodvana will manage the agent lifecycle, including updating the agent automatically.")
	addK8sCmd.Flags().StringArray("agent-env", nil, "Env variables supplied to agent - use this for setting HTTPS_PROXY, NO_PROXY etc")
}
