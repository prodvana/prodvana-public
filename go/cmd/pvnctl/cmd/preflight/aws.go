package preflight

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/prodvana/prodvana-public/go/pkg/cmdutil/tui"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/pkg/errors"
)

var eksArnContextNameRegex = regexp.MustCompile(`^arn:aws:eks:.*:cluster\/.*$`)

// eksctl adds context names like this: prodvana-support-ecs@mike-test-eksctl.us-west-2.eksctl.io
var eksctlSuffix = ".eksctl.io"

// eks context names by default are the eks arn eg: arn:aws:eks:us-east-2:718175777272:cluster/test
// this function will detect if the context name is an eks arn or not and return true if it is
func DetectEKSContextName(context string) bool {
	return eksArnContextNameRegex.MatchString(context) || strings.HasSuffix(context, eksctlSuffix)
}

var eksHostnameSuffix = ".eks.amazonaws.com"

func DetectEKSHostname(hostname string) bool {
	return strings.HasSuffix(hostname, eksHostnameSuffix)
}

// extracts the cluster name and the region from the eks arn
// example: arn:aws:eks:us-west-2:717408832317:cluster/naphat-testing-eksctl
// eksctl adds context names like this: prodvana-support-ecs@mike-test-eksctl.us-west-2.eksctl.io
func extractClusterDetailsFromContextName(contextName string) (string, string, error) {
	if strings.HasPrefix(contextName, "arn:aws:eks:") {
		parts := strings.Split(contextName, ":")
		if len(parts) != 6 {
			return "", "", fmt.Errorf("invalid eks ARN: %s", contextName)
		}
		region := parts[3]
		clusterName := strings.Split(parts[5], "/")[1]
		return clusterName, region, nil
	} else if strings.HasSuffix(contextName, eksctlSuffix) {
		parts := strings.Split(contextName, ".")
		if len(parts) != 4 {
			return "", "", fmt.Errorf("invalid eksctl context name: %s", contextName)
		}
		region := parts[1]
		clusterName := strings.Split(parts[0], "@")[1]
		return clusterName, region, nil
	}

	return "", "", fmt.Errorf("this context does not appear to be an EKS cluster: %s", contextName)
}

func extractIAMUserFromArn(arn string) (string, error) {
	re := regexp.MustCompile(`arn:aws:iam::\d+:user/(.+)`)
	matches := re.FindStringSubmatch(arn)
	if len(matches) != 2 {
		return "", fmt.Errorf("invalid iam user ARN: %s", arn)
	}
	return matches[1], nil
}

func DetectEKSCluster(context string) (bool, string, string) {
	if DetectEKSContextName(context) {
		clusterName, region, err := extractClusterDetailsFromContextName(context)
		if err != nil {
			return false, "", ""
		}
		return true, clusterName, region
	}
	return false, "", ""
}

func detectAWSCredentials() *sts.GetCallerIdentityOutput {
	// excluding EC2RoleProvider because it is very slow timing out and its very likely not relevant
	providers := []credentials.Provider{
		&credentials.EnvProvider{},
		&credentials.SharedCredentialsProvider{},
	}
	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			CredentialsChainVerboseErrors: aws.Bool(true),
			Credentials:                   credentials.NewChainCredentials(providers),
		},
	})
	if err != nil {
		return nil
	}

	stsClient := sts.New(sess)
	resp, err := stsClient.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err != nil {
		return nil
	}
	return resp
}

func createAWSSession(profile, region string) (*session.Session, error) {
	return session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           profile,
		Config: aws.Config{
			Region:                        aws.String(region),
			CredentialsChainVerboseErrors: aws.Bool(true),
		},
	})
}

type EKSClusterDetails struct {
	Cluster         *eks.Cluster
	NodeGroups      map[string]*eks.Nodegroup
	FargateProfiles map[string]*eks.FargateProfile
}

func fetchEKSClusterDetails(eksClient *eks.EKS, clusterName string) (*EKSClusterDetails, error) {
	describeClusterOutput, err := eksClient.DescribeCluster(&eks.DescribeClusterInput{
		Name: &clusterName,
	})
	if err != nil {
		return nil, err
	}

	nodeGroupsOutput, err := eksClient.ListNodegroups(&eks.ListNodegroupsInput{
		ClusterName: &clusterName,
	})
	if err != nil {
		return nil, err
	}
	nodeGroups := make(map[string]*eks.Nodegroup, len(nodeGroupsOutput.Nodegroups))
	for _, nodeGroup := range nodeGroupsOutput.Nodegroups {
		// get details about the node group
		nodeGroupDetails, err := eksClient.DescribeNodegroup(&eks.DescribeNodegroupInput{
			ClusterName:   &clusterName,
			NodegroupName: nodeGroup,
		})
		if err != nil {
			return nil, err
		}
		nodeGroups[*nodeGroup] = nodeGroupDetails.Nodegroup
	}

	if err != nil {
		return nil, err
	}

	fargateProfilesOutput, err := eksClient.ListFargateProfiles(&eks.ListFargateProfilesInput{
		ClusterName: &clusterName,
	})
	if err != nil {
		return nil, err
	}

	fargateProfiles := make(map[string]*eks.FargateProfile, len(fargateProfilesOutput.FargateProfileNames))
	for _, fargateProfileName := range fargateProfilesOutput.FargateProfileNames {
		fargateProfileDetails, err := eksClient.DescribeFargateProfile(&eks.DescribeFargateProfileInput{
			ClusterName:        &clusterName,
			FargateProfileName: fargateProfileName,
		})
		if err != nil {
			return nil, err
		}
		fargateProfiles[*fargateProfileName] = fargateProfileDetails.FargateProfile
	}
	return &EKSClusterDetails{
		Cluster:         describeClusterOutput.Cluster,
		NodeGroups:      nodeGroups,
		FargateProfiles: fargateProfiles,
	}, nil
}

func InteractiveGetRegion() (string, error) {
	var err error
	region, err := tui.QuestionInteractive("Which AWS region would you like to use?", "us-east-1")
	if err != nil {
		return "", errors.Wrap(err, "Error getting region selection")
	}
	if region == "" {
		region = "us-east-1"
	}
	return region, nil
}

func InteractiveGetAWSSession(inRegion string) (*session.Session, string, string, error) {
	profile, err := tui.QuestionInteractive("Which AWS profile would you like to use?", "default")
	if err != nil {
		return nil, "", "", errors.Wrap(err, "Error getting profile selection")
	}
	if profile == "" {
		profile = "default"
	}
	region := inRegion
	if region == "" {
		region, err = InteractiveGetRegion()
		if err != nil {
			return nil, "", "", err
		}
	}
	sess, err := createAWSSession(profile, region)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "Error creating AWS session")
	}
	return sess, profile, region, nil
}
