package deployments

import (
	"context"
	go_errors "errors"
	"fmt"
	"io"
	"prodvana/cmd/cmdutil"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/fatih/color"
	"github.com/pkg/errors"
	deployment_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/deployment"
	deployment_model_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/deployment/model"
	"github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/impact_analysis"
	"github.com/spf13/cobra"
)

var listFlags = struct {
	application     string
	services        []string
	releaseChannels []string
	desiredStateId  string
}{}

func formatCommit(commitId string) string {
	const limit = 7
	if len(commitId) < limit {
		return commitId
	}
	return commitId[:limit]
}

func withPeriod(message string) string {
	if strings.HasSuffix(message, ".") {
		return message
	}
	return message + "."
}

func impactExplanation(result *impact_analysis.ImpactAnalysisResult) string {
	explanation := withPeriod(result.GetImpactReasoning())
	if result.GetLikelihoodReasoning() != "" {
		explanation += " " + withPeriod(result.LikelihoodReasoning)
	}
	return strings.TrimSpace(explanation)
}

func formatDeployment(rel *deployment_model_pb.Deployment, w io.Writer) error {
	color.New(color.FgYellow).Fprintf(w, "deployment %s\n", rel.Meta.Id)
	headerTabW := tabwriter.NewWriter(w, 0, 0, 1, ' ', 0)
	fmt.Fprintf(headerTabW, "Created Date:\t%s\n", rel.Config.CreationTimestamp.AsTime().Format(time.RFC3339))
	if rel.State.LastUpdateTimestamp.AsTime() != rel.Config.CreationTimestamp.AsTime() {
		fmt.Fprintf(headerTabW, "Last Updated Date:\t%s\n", rel.State.LastUpdateTimestamp.AsTime().Format(time.RFC3339))
	}
	fmt.Fprintf(headerTabW, "Status:\t%v\n", rel.State.Status)
	err := headerTabW.Flush()
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "\n")
	bodyTabW := tabwriter.NewWriter(w, 0, 0, 1, ' ', 0)
	fmt.Fprintf(bodyTabW, "\t\tDeployment System:\t%s\n", rel.Config.DeploymentSystem)
	if rel.Config.User != "" {
		fmt.Fprintf(bodyTabW, "\t\tUser:\t%s\n", rel.Config.User)
	}
	fmt.Fprintf(bodyTabW, "\t\tApplication:\t%s\n", rel.Config.Application)
	fmt.Fprintf(bodyTabW, "\t\tService:\t%s\n", rel.Config.Service)
	fmt.Fprintf(bodyTabW, "\t\tRelease Channel:\t%s\n", rel.Config.ReleaseChannel)
	if rel.Config.Repository != "" && rel.Config.CommitId != "" {
		fmt.Fprintf(bodyTabW, "\t\tRepository:\t%s\n", rel.Config.Repository)
		if rel.Comparison != nil && rel.Comparison.PrevCommitId != "" {
			fmt.Fprintf(bodyTabW, "\t\tCommit Range:\t%s..%s", formatCommit(rel.Comparison.PrevCommitId), formatCommit(rel.Config.CommitId))
			commitAnalysis := rel.Comparison.CommitAnalysis
			if commitAnalysis != nil {
				fmt.Fprintf(bodyTabW, " (")
				first := true
				if commitAnalysis.CommitsAdded > 0 {
					fmt.Fprintf(bodyTabW, "%d commits added", commitAnalysis.CommitsAdded)
					first = false
				}
				if commitAnalysis.CommitsRemoved > 0 {
					if !first {
						fmt.Fprintf(bodyTabW, ", ")
					}
					fmt.Fprintf(bodyTabW, "%d commits removed", commitAnalysis.CommitsRemoved)
				}
				fmt.Fprintf(bodyTabW, ")\n")
			} else {
				fmt.Fprintf(bodyTabW, "\n")
			}
		} else {
			fmt.Fprintf(bodyTabW, "\t\tCommit:\t%s\n", formatCommit(rel.Config.CommitId))
		}
	}
	err = bodyTabW.Flush()
	if err != nil {
		return err
	}
	relevantAddedCommits := rel.Comparison.GetCommitAnalysis().GetImpactAnalysis().GetRelevantAddedCommits()
	if len(relevantAddedCommits) > 0 {
		commitTabW := tabwriter.NewWriter(w, 0, 0, 1, ' ', 0)
		fmt.Fprintf(commitTabW, "\n")
		if len(relevantAddedCommits) > 0 {
			fmt.Fprintf(commitTabW, "\t\tHigh Impact Commits Added:\n")
		}
		for _, commit := range relevantAddedCommits {
			fmt.Fprintf(commitTabW, "\t\t\t\t%s (%s):\t%s\t%s\n", formatCommit(commit.CommitId), commit.Author.GetEmail(), commit.ImpactAnalysis.GetCategory(), impactExplanation(commit.ImpactAnalysis))
		}
		if err := commitTabW.Flush(); err != nil {
			return err
		}
	}
	relevantRemovedCommits := rel.Comparison.GetCommitAnalysis().GetImpactAnalysis().GetRelevantRemovedCommits()
	if len(relevantRemovedCommits) > 0 {
		commitTabW := tabwriter.NewWriter(w, 0, 0, 1, ' ', 0)
		fmt.Fprintf(commitTabW, "\n")
		if len(relevantRemovedCommits) > 0 {
			fmt.Fprintf(commitTabW, "\t\tHigh Impact Commits Removed:\n")
		}
		for _, commit := range relevantRemovedCommits {
			fmt.Fprintf(commitTabW, "\t\t\t\t%s (%s):\t%s\t%s\n", formatCommit(commit.CommitId), commit.Author.GetEmail(), commit.ImpactAnalysis.GetCategory(), impactExplanation(commit.ImpactAnalysis))
		}
		if err := commitTabW.Flush(); err != nil {
			return err
		}
	}
	unanalyzedCommits := rel.Comparison.GetCommitAnalysis().GetImpactAnalysis().GetUnanalyzedCommits()
	if unanalyzedCommits > 0 {
		color.New(color.FgYellow).Fprintf(w, "\n  Still analyzing %d commits. Check back later.\n", unanalyzedCommits)
	}
	fmt.Fprintf(w, "\n")
	return nil
}

func streamDeployments(strm deployment_pb.DeploymentManager_ListDeploymentsStreamClient, w io.WriteCloser) error {
	defer func() { _ = w.Close() }()
	err := strm.CloseSend()
	if err != nil {
		return errors.Wrap(err, "failed to close send")
	}
	for {
		resp, err := strm.Recv()
		if err != nil {
			w.Close()
			if go_errors.Is(err, io.EOF) {
				return nil
			}
			return errors.Wrap(err, "failed to")
		}
		for _, rel := range resp.Deployments {
			err := formatDeployment(rel, w)
			if err != nil {
				return errors.Wrap(err, "failed to write deployment")
			}
		}
	}
}

var logCmd = &cobra.Command{
	Use:     "log",
	Aliases: []string{"list"},
	Short:   "List deployments.",
	Long: `List deployments. Follows a similar semantic to git log.

pvnctl deployments log  # list all deployments
pvnctl deployments log <rel-id>  # list all deployments beginning with release id
pvnctl deployments log <rel-id1>..<rel-id2>  # list all deployments from rel-id1 to rel-id2. rel-id1 must be before rel-id2.
pvnctl deployments log --service=<service> --service=<service2> # list all deployments for service and service2 in all release channels
pvnctl deployments log --service=<service> --release-channel=<channel> --release-channel=<channel2>  # list all deployments for service in channel and channel2
`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		var endDeploymentId string
		var startDeploymentId string
		if len(args) == 1 {
			if strings.Contains(args[0], "..") {
				components := strings.SplitN(args[0], "..", 2)
				startDeploymentId = components[0]
				endDeploymentId = components[1]
			} else {
				endDeploymentId = args[0]
			}
		}
		var filters []*deployment_pb.DeploymentFilter
		if len(listFlags.services) > 0 || len(listFlags.releaseChannels) > 0 || len(listFlags.application) > 0 || len(listFlags.desiredStateId) > 0 {
			filters = append(filters, &deployment_pb.DeploymentFilter{
				Application:     listFlags.application,
				Services:        listFlags.services,
				ReleaseChannels: listFlags.releaseChannels,
				DesiredStateId:  listFlags.desiredStateId,
			})
		}
		strm, err := cmdutil.GetDeploymentManagerClient().ListDeploymentsStream(ctx, &deployment_pb.ListDeploymentsReq{
			Filters:              filters,
			StartingDeploymentId: startDeploymentId,
			EndingDeploymentId:   endDeploymentId,
		})
		if err != nil {
			return errors.Wrap(err, "failed to list deployments")
		}
		r, w := io.Pipe()
		outChan := make(chan error, 1)
		go func() {
			defer close(outChan)
			defer func() { _ = r.Close() }()
			outChan <- cmdutil.MaybePage(r)
		}()
		strmErr := streamDeployments(strm, w)
		for outErr := range outChan {
			if outErr != nil {
				return outErr
			}
		}
		if strmErr != nil {
			return strmErr
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(logCmd)
	logCmd.Flags().StringVarP(&listFlags.application, "app", "a", "", "Filter to this application.")
	logCmd.Flags().StringSliceVarP(&listFlags.services, "service", "s", nil, "Filter to this service. Can be specified multiple times and will be joined using OR. Joined with --env using AND.")
	logCmd.Flags().StringSliceVar(&listFlags.releaseChannels, "release-channel", nil, "Filter to this release channel. Can be specified multiple times and will be joined using OR. Joined with --services using OR.")
	logCmd.Flags().StringVar(&listFlags.desiredStateId, "desired-state-id", "", "Filter to this desired state id.")
}
