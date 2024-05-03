package preflight

import (
	"fmt"
	"os"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil/tui"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/bradenaw/juniper/xslices"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	foundStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
	missingStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
)

var ecsCmd = &cobra.Command{
	Use:   "ecs",
	Short: "Run preflight checks to validate your ECS credentials are configured for Prodvana.",
	Long: `Run preflight checks to validate your ECS credentials are configured for Prodvana.

pvnctl preflight ecs
`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		caller := detectAWSCredentials()
		if caller == nil {
			cmd.PrintErr("No local AWS credentials detected. We will use your local AWS credentials to validate the ECS service account you will use to configure your Prodvana Runtime.\nPlease configure your local AWS credentials and try again.\n")
			os.Exit(1)
		}

		accepted, err := tui.PromptInteractive("We will ask for two credentials: a local credentials for validation only and the credentials you will be using to configure the ECS Runtime.\nThe local credentials will be used locally and discarded once validated. Continue? (Y/n)", true)
		if err != nil {
			cmd.PrintErrf("Error getting selection: %v\n", err)
			os.Exit(1)
		}
		if !accepted {
			cmd.PrintErr("Aborting preflight checks.\n")
			os.Exit(1)
		}

		var awsSess *session.Session
		correctCreds, err := tui.PromptInteractive(fmt.Sprintf("We found the following AWS Identity: %s. Is this the identity you would like to use to run verification locally (*not* the credentials you will use to configure the ECS Runtime)? (Y/n)", *caller.Arn), true)
		if err != nil {
			cmd.PrintErrf("Error getting selection: %v\n", err)
			os.Exit(1)
		}
		if correctCreds {
			sess, err := session.NewSessionWithOptions(session.Options{
				SharedConfigState: session.SharedConfigEnable,
			})
			if err != nil {
				cmd.PrintErrf("Error creating AWS session: %v\n", err)
				os.Exit(1)
			}
			if aws.StringValue(sess.Config.Region) == "" {
				region, err := InteractiveGetRegion()
				if err != nil {
					cmd.PrintErrf("Error getting region: %v\n", err)
					os.Exit(1)
				}
				sess.Config.Region = aws.String(region)
			}
			awsSess = sess
		} else {
			sess, _, _, err := InteractiveGetAWSSession("")
			if err != nil {
				cmd.PrintErrf("Error creating AWS session: %v\n", err)
				os.Exit(1)
			}
			awsSess = sess
		}

		cmd.Println()
		cmd.Println("We will now prompt for the AWS Access Key ID and Secret Access Key of the IAM user you will use to configure your Prodvana Runtime,\nso we can validate that the user has the required permissions.")

		accessKeyID, err := tui.QuestionInteractive("AWS Access Key ID:", "")
		if err != nil {
			cmd.PrintErrf("Error getting selection: %v\n", err)
			os.Exit(1)
		}
		if accessKeyID == "" {
			cmd.PrintErrln("No AWS Access Key ID provided, aborting preflight checks")
			os.Exit(1)
		}
		secretKey, err := tui.QuestionInteractiveSecret("AWS Secret Access Key:", "")
		if err != nil {
			cmd.PrintErrf("Error getting selection: %v\n", err)
			os.Exit(1)
		}
		if secretKey == "" {
			cmd.PrintErrln("No AWS Secret Access Key provided, aborting preflight checks")
			os.Exit(1)
		}

		testCredsSess, err := session.NewSession(&aws.Config{
			Credentials: credentials.NewStaticCredentials(accessKeyID, secretKey, ""),
		})
		if err != nil {
			cmd.PrintErrf("Error creating AWS session: %v\n", err)
			cmd.PrintErr("Double check that the AWS Access Key ID and Secret Access Key you provided are correct.\n")
			os.Exit(1)
		}

		stsClient := sts.New(testCredsSess)
		resp, err := stsClient.GetCallerIdentity(&sts.GetCallerIdentityInput{})
		if err != nil {
			cmd.PrintErrf("Failed to identify the provided credentials: %v\n", err)
			cmd.PrintErr("Double check that the AWS Access Key ID and Secret Access Key you provided are correct.\n")
			os.Exit(1)
		}
		iamUser, err := extractIAMUserFromArn(*resp.Arn)
		if err != nil {
			cmd.PrintErrf("Failed to identify the provided credentials: %v\n", err)
			cmd.PrintErr("Double check that the AWS Access Key ID and Secret Access Key you provided are correct.\n")
			os.Exit(1)
		}

		iamClient := iam.New(awsSess)
		input := &iam.ListAttachedUserPoliciesInput{UserName: &iamUser}
		result, err := iamClient.ListAttachedUserPolicies(input)
		if err != nil {
			if awsErr, ok := err.(awserr.Error); ok {
				// check for IAM Errors
				cmd.PrintErrln()
				switch awsErr.Code() {
				case "AccessDenied":
					cmd.PrintErrf("Access Denied Error: %s\n", awsErr.Message())
					cmd.PrintErrln("Please ensure the AWS credentials you selected have the required permissions to list attached policies.")
				case "UnauthorizedOperation":
					cmd.PrintErrf("Unauthorized Operation Error: %s\n", awsErr.Message())
					cmd.PrintErrln("Please ensure the AWS credentials you selected have the required permissions to list attached policies.")
				case "NoSuchEntity":
					cmd.PrintErr(missingStyle.Render("The AWS API returned a NoSuchEntity error. This could be caused by a few things:") + "\n")
					cmd.PrintErrf("  - Double check that '%s' is the correct IAM user name\n", iamUser)
					cmd.PrintErrln("  - Make sure the AWS credentials you're using have access to the same account as the IAM user you provided.")
					cmd.PrintErrf("AWS Error: %s\n", awsErr.Message())
				default:
					cmd.PrintErrf("Error listing attached policies: %s", err)
				}
			} else {
				cmd.PrintErrf("Error listing attached policies: %s", err)
			}
			os.Exit(1)
		}

		desiredPolicies := map[string]bool{
			"AmazonECS_FullAccess":                     false,
			"ResourceGroupsandTagEditorReadOnlyAccess": false,
		}

		for _, policy := range result.AttachedPolicies {
			if _, ok := desiredPolicies[*policy.PolicyName]; ok {
				desiredPolicies[*policy.PolicyName] = true
			}
		}

		cmd.Println()
		for policyName, found := range desiredPolicies {
			if found {
				cmd.Println(foundStyle.Render(fmt.Sprintf("[ ✓ ] IAM user %s has required policy: %s", iamUser, policyName)))
			} else {
				cmd.Println(missingStyle.Render(fmt.Sprintf("[ x ] IAM user %s is missing required policy: %s", iamUser, policyName)))
			}
		}

		// now find the specific cluster that the user will be onboarding
		ecsClient := ecs.New(awsSess)
		clusters, err := ecsClient.ListClusters(&ecs.ListClustersInput{})
		if err != nil {
			cmd.PrintErrf("Error listing ECS clusters: %s", err)
			os.Exit(1)
		}
		if len(clusters.ClusterArns) == 0 {
			cmd.Println(missingStyle.Render("[ x ] No ECS clusters found"))
		} else {
			selectedCluster, err := tui.GetInteractiveSelection("Select the ECS cluster you will be onboarding", xslices.Map(clusters.ClusterArns, aws.StringValue))
			if err != nil {
				cmd.PrintErrf("Error getting selection: %v\n", err)
				os.Exit(1)
			}
			_, err = ecsClient.DescribeClusters(&ecs.DescribeClustersInput{
				Clusters: []*string{
					aws.String(selectedCluster),
				},
			})
			if err != nil {
				cmd.PrintErrf("Error describing ECS cluster %s: %s", selectedCluster, err)
				cmd.PrintErrf("Double check that the AWS credentials (%s) to the ECS cluster %s.\n", *caller.Arn, selectedCluster)
				// TODO(naphat) soft error?
				os.Exit(1)
			}
			cmd.Println(foundStyle.Render(fmt.Sprintf("[ ✓ ] Cluster %s valid", selectedCluster)))
		}
	},
}

func init() {
	RootCmd.AddCommand(ecsCmd)
}
