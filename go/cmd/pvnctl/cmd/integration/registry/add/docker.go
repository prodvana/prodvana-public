package add

import (
	"context"
	"fmt"
	"log"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil"

	workflow_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/workflow"

	"github.com/spf13/cobra"
)

var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Add or update a docker registry integration",
	Long: `Add or update a docker registry integration. Registry names are unique, so passing an existing name will update that integration.

pvnctl integration registry add docker <name> --url <url> --username <username> --password <password>
pvnctl integration registry add docker <name> --url <url> --public
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		url, err := cmd.Flags().GetString("url")
		if err != nil {
			log.Fatal(err)
		}
		public, err := cmd.Flags().GetBool("public")
		if err != nil {
			log.Fatal(err)
		}
		var username, password string
		if !public {
			username, err = cmd.Flags().GetString("username")
			if err != nil {
				log.Fatal(err)
			}
			password, err = cmd.Flags().GetString("password")
			if err != nil {
				log.Fatal(err)
			}
			if username == "" || password == "" {
				log.Fatal("Pass --username and --password or --public")
			}
		}

		req := &workflow_pb.CreateContainerRegistryIntegrationReq{
			Url:      url,
			Username: username,
			Secret:   password,
			Name:     name,
		}
		if public {
			req.RegistryOptions = &workflow_pb.CreateContainerRegistryIntegrationReq_PublicRegistryOptions_{
				PublicRegistryOptions: &workflow_pb.CreateContainerRegistryIntegrationReq_PublicRegistryOptions{},
			}
		}
		resp, err := cmdutil.GetWorkflowManagerClient().CreateContainerRegistryIntegration(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(resp.IntegrationId)
	}}

func init() {
	RootCmd.AddCommand(dockerCmd)
	dockerCmd.Flags().String("url", "", "base url of the container registry")
	dockerCmd.Flags().String("username", "", "username for the container registry")
	dockerCmd.Flags().String("password", "", "password or personal access token for the container registry")
	dockerCmd.Flags().Bool("public", false, "add a public docker registry. If not true, username and password must be passed.")
	cmdutil.Must(dockerCmd.MarkFlagRequired("url"))
}
